package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ligurio/recidive/formats"
	"github.com/matoous/go-nanoid"
)

// 2MB
const maxSize = 2 * 1024 * 1024

func viewHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP request %s -> %s %s\n", r.RemoteAddr, r.Method, r.URL)

	id := mux.Vars(r)["id"]

	db := initDb(dbpath)
	defer db.Close()

	report := new(formats.Report)
	if err := db.Where("uid = ?", id).First(&report).Error; err != nil {
		log.Println(err.Error())
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	db.Preload("Suites.Tests").Find(&report)
	db.Preload("Suites").Preload("Suites.Tests").Find(&report)

	report.Hits = report.Hits + 1
	db.Save(&report)

	err := templates.ExecuteTemplate(w, "view", report)
	if err != nil {
		log.Println(err)
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func opensearchHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP request %s -> %s %s\n", r.RemoteAddr, r.Method, r.URL)

	var tmpl = template.Must(template.ParseFiles("static/templates/opensearch.xml"))
	err := tmpl.Execute(w, r.Host)
	if err != nil {
		log.Println("DEBUG: ", err)
		errorHandler(w, r, http.StatusInternalServerError)
	}
}

func ProcessQuery(r *http.Request) ([]formats.Report, error) {
	/*
		Process HTTP query and return a list of reports
	*/

	duration := r.URL.Query().Get("duration")
	if duration != "" {
		log.Println("URL parameter 'duration' is: ", string(duration))
	}

	period := 0
	period, _ = strconv.Atoi(duration)

	db := initDb(dbpath)
	defer db.Close()

	var reports []formats.Report
	if period > 0 {
		now := time.Now()
		then := now.AddDate(0, 0, 0-period)
		db.Where("created_at BETWEEN ? AND ?", then, now).Find(&reports)
	} else {
		db.Find(&reports)
	}

	return reports, nil
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP request %s -> %s %s\n", r.RemoteAddr, r.Method, r.URL)
	log.Println("GET params were:", r.URL.Query())

	// TODO: pass r.URL.RequestURI() to chart URL in HTML template
	reports, err := ProcessQuery(r)
	err = templates.ExecuteTemplate(w, "list", reports)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP request %s -> %s %s\n", r.RemoteAddr, r.Method, r.URL)
	if r.Method == "GET" {
		err := templates.ExecuteTemplate(w, "upload", nil)
		if err != nil {
			log.Println("DEBUG: ", err)
			errorHandler(w, r, http.StatusInternalServerError)
		}
		return
	}

	if err := r.ParseMultipartForm(maxSize); err != nil {
		log.Printf("DEBUG: Max size is exceeded in ParseMultipartForm: %s\n", err)
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	db := initDb(dbpath)
	defer db.Close()

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, err := fileHeader.Open()
			if err != nil {
				log.Println("Failed to open file")
			}
			log.Printf("DEBUG: File: %s\n", fileHeader.Filename)
			report, err := formats.ReadReport(file, fileHeader.Filename)

			report.UID = makeID()
			log.Println("Report ID is", report.UID)

			if err == nil && report != nil {
				log.Println("DEBUG: successful upload")
				db.Debug().Create(&report)
				errors := db.GetErrors()
				if len(errors) != 0 {
					for err := range errors {
						log.Println("DEBUG: Insert failed", err)
					}
					errorHandler(w, r, http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(report.UID + "\n"))
				}
			} else {
				log.Println("DEBUG: ReadReport failed", err)
				errorHandler(w, r, http.StatusInternalServerError)
			}
		}
	}
}

func makeID() string {

	id, err := gonanoid.Nanoid(10)
	if err != nil {
		panic(err)
	}

	return id
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	switch status {
	case http.StatusNotFound:
		renderTemplate(w, "notfound")
	case http.StatusInternalServerError:
		renderTemplate(w, "notfound")
	}
}

func StartServer(listenAddr string, staticDir *string) error {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", listHandler).Methods("GET")
	r.HandleFunc("/view/{id}/", viewHandler).Methods("GET")
	r.HandleFunc("/upload", uploadHandler).Methods("GET", "POST")
	r.HandleFunc("/chart/pie", drawPieChart).Methods("GET")
	r.HandleFunc("/chart/bar", drawBarChart).Methods("GET")
	r.HandleFunc("/chart/stackedbar", drawStackedBarChart).Methods("GET")
	r.HandleFunc("/chart/twoaxes", drawTwoAxesChart).Methods("GET")
	r.HandleFunc("/opensearch.xml", opensearchHandler).Methods("GET")

	s := http.StripPrefix("/static/", http.FileServer(http.Dir(*staticDir)))
	r.PathPrefix("/static/").Handler(s)
	http.Handle("/", r)

	log.Printf("Start on %s\n", *httpAddr)
	return http.ListenAndServe(*httpAddr, r)

}

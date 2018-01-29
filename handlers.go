package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// 2MB
const maxSize = 2 * 1024 * 1024

func viewHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP request %s -> %s %s\n", r.RemoteAddr, r.Method, r.URL)

	id := mux.Vars(r)["id"]

	db := initDb(dbpath)
	defer db.Close()

	report := new(Report)
	report.ReportId = id
	db.Find(&report, "report_id = ?", report.ReportId)
	if db.GetErrors() != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	report.Hits = report.Hits + 1
	db.Save(&report)

	err := templates.ExecuteTemplate(w, "view", report)
	if err != nil {
		log.Println(err)
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP request %s -> %s %s\n", r.RemoteAddr, r.Method, r.URL)

	err := templates.ExecuteTemplate(w, "about", r.Host)
	if err != nil {
		log.Println("DEBUG: ", err)
		errorHandler(w, r, http.StatusInternalServerError)
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

func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP request %s -> %s %s\n", r.RemoteAddr, r.Method, r.URL)

	db := initDb(dbpath)
	defer db.Close()

	var reports []Report
	db.Find(&reports)
	if db.GetErrors() != nil {
		errorHandler(w, r, http.StatusInternalServerError)
	}

	err := templates.ExecuteTemplate(w, "list", reports)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP request %s -> %s %s\n", r.RemoteAddr, r.Method, r.URL)
	if r.Method == "GET" {
		ctime := time.Now().Unix()

		hasher := sha1.New()
		hasher.Write([]byte(strconv.FormatInt(ctime, 10)))
		token := fmt.Sprintf("%x", hasher.Sum(nil))

		err := templates.ExecuteTemplate(w, "upload", token)
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

	token := r.Form.Get("token")

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, err := fileHeader.Open()
			if err != nil {
				log.Println("Failed to open file")
			}
			log.Printf("DEBUG: File: %s\n", fileHeader.Filename)
			report, err := ReadReport(file, fileHeader.Filename)
			if err == nil && report != nil {
				log.Println("DEBUG: successful upload")

				db.Debug().NewRecord(report)
				db.Debug().Create(&report)
				db.Debug().NewRecord(report)

				if db.GetErrors() != nil {
					log.Println("DEBUG: Insert failed", err)
					errorHandler(w, r, http.StatusInternalServerError)
				} else {
					if token != "" {
						// FIXME: page without token
						renderTemplate(w, "upload")
					} else {
						w.WriteHeader(http.StatusOK)
						w.Write([]byte(report.ReportId + "\n"))
					}
				}
			} else {
				log.Println("DEBUG: ReadReport failed", err)
				errorHandler(w, r, http.StatusInternalServerError)
			}
		}
	}
}

func makeid() string {
	return time.Now().Format("20060102-150405")
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
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/view/{id}/", viewHandler).Methods("GET")
	r.HandleFunc("/upload", uploadHandler).Methods("GET", "POST")
	r.HandleFunc("/opensearch.xml", opensearchHandler).Methods("GET")

	s := http.StripPrefix("/static/", http.FileServer(http.Dir(*staticDir)))
	r.PathPrefix("/static/").Handler(s)
	http.Handle("/", r)

	log.Printf("Start on %s\n", *httpAddr)
	return http.ListenAndServe(*httpAddr, r)

}

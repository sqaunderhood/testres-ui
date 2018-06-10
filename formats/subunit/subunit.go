// Package subunit provides a reader of the SubUnit protocol.

package subunit

import (
	"io"
)

func NewParser(r io.Reader) (*Testsuite, error) {

	if ts, err := Parser_v2(r); err != nil {
		return nil, err
	} else {
		return ts, nil
	}

	//return &ts, nil
}

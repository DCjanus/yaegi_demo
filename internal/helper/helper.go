// Package helper provides helper functions for the application.
// This package is for showing how to call host functions in interpreter.
package helper

import (
	"net/http"
)

var UselessHeader = "useless-helper"

func UselessHelper(w http.ResponseWriter, name string) {
	w.Header().Set("Via", name)
}

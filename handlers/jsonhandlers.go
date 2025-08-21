package handlers

import (
	cd "LoftSQLite3/capaDatos"
	"encoding/json"
	"fmt"
	"net/http"
)

func Wilder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	wilder := cd.Wilder()
	byteJson, e := json.Marshal(wilder)
	if e != nil {
		fmt.Print(e.Error())
	}
	fmt.Fprintf(w, string(byteJson))
}

package handlers

import (
	"encoding/json"
	"fmt"
	cd "github.com/angorita/loft/capaDatos"
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

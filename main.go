package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	bind := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hai from test"))
	})

	http.ListenAndServe(bind, nil)
}

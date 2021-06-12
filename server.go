package main

import (
	"log"
	"net/http"
)

func main() {
	// Handler points to available directories
	http.Handle("/web/html", http.StripPrefix("/web/html", http.FileServer(http.Dir("web/html"))))
	http.Handle("/web/scripts/", http.StripPrefix("/web/scripts/", http.FileServer(http.Dir("web/scripts"))))
	http.Handle("/web/css/", http.StripPrefix("/web/css/", http.FileServer(http.Dir("web/css"))))
	http.Handle("/web/media/", http.StripPrefix("/web/media/", http.FileServer(http.Dir("web/media"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path[1:] == "" {
			http.ServeFile(w, r, "web/html/index.html")
		}
	})

	//Serves local webpage for testing
	if true {
		errhttp := http.ListenAndServe(":8080", nil)
		if errhttp != nil {
			log.Fatal("Web server (HTTP): ", errhttp)
		}
	} else {
		//Serves the webpage to the internet
		errhttps := http.ListenAndServeTLS(":443", "certs/cert.pem", "certs/key.pem", nil)
		if errhttps != nil {
			log.Fatal("Web server (HTTPS): ", errhttps)
		}
	}

}

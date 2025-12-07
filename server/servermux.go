package server

import "net/http"

func CreateServerWithMux() {
	mux := http.NewServeMux()
	mux.Handle("/google", http.RedirectHandler("https://google.com", 301))
	mux.HandleFunc("/", ServeHttp)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

package main

import (
	"net/http"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	http := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		ReadTimeout:       time.Second * 60,
		ReadHeaderTimeout: time.Second * 60,
	}
	err := http.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

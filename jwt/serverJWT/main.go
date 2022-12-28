package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/welcomeThroughCookie", WelcomeThroughCookie)
	http.HandleFunc("/welcomeThroughBearer", WelcomeThroughBearer)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

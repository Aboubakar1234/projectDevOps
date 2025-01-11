package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type whoami struct {
	Name  string
	Title string
	State string
}

func main() {
	request1()
}

func whoAmI(response http.ResponseWriter, r *http.Request) {
	who := []whoami{
		whoami{Name: "Aboubakar mohamed abdoulaye",
			Title: "DevOps and Continous Deployment",
			State: "FR",
		},
	}

	json.NewEncoder(response).Encode(who)

	fmt.Println("Endpoint Hit: /whoami")
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the Web API!")
	fmt.Println("Endpoint Hit: homePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "Aboubakar mohamed abdoulaye"

	fmt.Fprintf(response, "A little bit about me...")
	fmt.Println("Endpoint Hit: ", who)
}

func healthCheck(response http.ResponseWriter, r *http.Request) {
    response.WriteHeader(http.StatusOK) // Renvoie un code HTTP 200
    fmt.Fprintf(response, "OK")        // Message simple pour indiquer que tout va bien
    fmt.Println("Endpoint Hit: /health")
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/whoami", whoAmI)
	http.HandleFunc("/health", healthCheck)
	
	log.Fatal(http.ListenAndServe(":8085", nil))
}

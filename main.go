
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    twilio "github.com/twilio/twilio-go"
)

var twilioClient *twilio.RestClient

func main() {
    // Initialize database connection
    InitDB("new_user:new_password@tcp(127.0.0.1:3306)/voip")

    // Initialize Twilio client
    accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
    authToken := os.Getenv("TWILIO_AUTH_TOKEN")
    twilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: accountSid,
        Password: authToken,
    })

    router := mux.NewRouter()
    router.HandleFunc("/register", RegisterHandler).Methods("POST")
    router.HandleFunc("/login", LoginHandler).Methods("POST")
    router.HandleFunc("/call", CallHandler).Methods("POST")

    fmt.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

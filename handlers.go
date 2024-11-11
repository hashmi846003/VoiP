package main

import (
    "encoding/json"
    "net/http"
    "fmt"

    openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Register a new user
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var user struct {
        Username string `json:"username"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
    if err != nil {
        http.Error(w, "Error registering user", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "User registered successfully")
}

// Authenticate user
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var user struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var storedPassword string
    err := DB.QueryRow("SELECT password FROM users WHERE username = ?", user.Username).Scan(&storedPassword)
    if err != nil || storedPassword != user.Password {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    fmt.Fprintf(w, "User logged in successfully")
}

// Initiate a call using Twilio
func CallHandler(w http.ResponseWriter, r *http.Request) {
    var call struct {
        To   string `json:"to"`   // Mobile number to call
        From string `json:"from"` // Your Twilio number
    }
    if err := json.NewDecoder(r.Body).Decode(&call); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    params := &openapi.CreateCallParams{}
    params.SetTo(call.To)
    params.SetFrom(call.From)
    params.SetUrl("http://demo.twilio.com/docs/voice.xml")

    _, err := twilioClient.Api.CreateCall(params)
    if err != nil {
        http.Error(w, "Failed to initiate call", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Call initiated successfully")
}


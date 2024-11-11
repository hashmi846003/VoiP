/*package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
	//handlers\handlers.go
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    router := mux.NewRouter()

    // Register handlers
    router.HandleFunc("/register", RegisterHandler).Methods("POST")
    router.HandleFunc("/login", LoginHandler).Methods("POST")
    router.HandleFunc("/call", CallHandler).Methods("POST")

    // Connect to the database
    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/voip")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Check the connection
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
*/
/*package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB

func main() {
    var err error
    // Update the database connection string with your credentials
    db, err := sql.Open("mysql", "new_user:new_password@tcp(127.0.0.1:3306)/voip")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Check the connection
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    router := mux.NewRouter()

    // Register handlers
    router.HandleFunc("/register", RegisterHandler).Methods("POST")
    router.HandleFunc("/login", LoginHandler).Methods("POST")
    router.HandleFunc("/call", CallHandler).Methods("POST")

    fmt.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
*/
/*
package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
    twilio "github.com/twilio/twilio-go"
   // openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

//var db *sql.DB
//var twilioClient *twilio.RestClient

func main() {
    var err error
    // Connect to MySQL
    db, err := sql.Open("mysql", "new_user:new_password@tcp(127.0.0.1:3306)/voip")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

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
*/
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

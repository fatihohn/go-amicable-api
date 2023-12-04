package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

type DbConfig struct {
    Host        string
    Dbname      string
    User        string
    Password    string
    Port        int
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Golang Dev Environment!")
}

func loadDbConfig() (DbConfig) {
    // Load environment variables from the .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Access environment variables
    host := os.Getenv("POSTGRES_HOST")
    dbname := os.Getenv("POSTGRED_DB")
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))

    return DbConfig{
        Host:       host,
        Dbname:     dbname,
        User:       user,
        Password:   password,
        Port:       port,
    }
}

func loadDB(config *DbConfig) {
	// Connect to the PostgreSQL database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")

	// Perform database operations (e.g., CRUD) here
	// ...
    // Example: Query data from a table
    // rows, err := db.Query("SELECT id, name FROM your_table")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer rows.Close()

    // for rows.Next() {
    //     var id int
    //     var name string
    //     err := rows.Scan(&id, &name)
    //     if err != nil {
    //         log.Fatal(err)
    //     }
    //     fmt.Println(id, name)
    // }

    // // Example: Insert data into a table
    // _, err = db.Exec("INSERT INTO your_table(name) VALUES($1)", "John Doe")
    // if err != nil {
    //     log.Fatal(err)
    // }

}

func loadServer() {
    http.HandleFunc("/", handler)
    fmt.Println("Server is running on :8080...")
    http.ListenAndServe(":8080", nil)
}

func main() {
    config := loadDbConfig()
    loadDB(&config)
    loadServer()
}


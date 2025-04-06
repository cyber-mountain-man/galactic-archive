package main

import (
	"encoding/json" // For decoding JSON data from entries.json
	"html/template" // For rendering HTML templates using Go's html/template engine
	"log"           // For logging server messages and errors
	"net/http"      // For building the web server and handling HTTP requests
	"os"            // For reading files from the filesystem
)

// Entry represents a single archival record in the Galactic Archive.
// It maps directly to a JSON object in entries.json.
type Entry struct {
    ID          string `json:"id"`          // Unique identifier for the entry
    Name        string `json:"name"`        // Display name/title of the entry
    Type        string `json:"type"`        // Category: planet, artifact, creed, event, etc.
    Faction     string `json:"faction"`     // Affiliated group or origin
    Description string `json:"description"` // The body text or lore description
}

// templates loads all HTML template files from the "templates" directory.
// template.Must will panic if parsing fails — useful during development.
var templates = template.Must(template.ParseGlob("templates/*.html"))

// GetEntriesByType reads entries from entries.json and returns only those
// matching the specified type (e.g. "creed", "event").
func GetEntriesByType(entryType string) ([]Entry, error) {
    // Read the raw JSON file into memory
    data, err := os.ReadFile("data/entries.json")
    if err != nil {
        return nil, err // Return an error if file read fails
    }

    var entries []Entry

    // Unmarshal JSON data into the entries slice
    if err := json.Unmarshal(data, &entries); err != nil {
        return nil, err // Return an error if JSON parsing fails
    }

    // Filter entries by the requested type
    var filtered []Entry
    for _, entry := range entries {
        if entry.Type == entryType {
            filtered = append(filtered, entry)
        }
    }

    return filtered, nil
}

func main() {
    // Serve static assets like CSS and JS from the "static" folder
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Route: Home page handler (GET /)
    // Renders index.html (which includes htmx buttons and UI)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("🛸 Reached / route: attempting to render index.html")
        err := templates.ExecuteTemplate(w, "layout.html", nil) 
        if err != nil {
            log.Println("❌ Error rendering template:", err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })
    

    // Route: /entries — this endpoint is triggered by htmx when a user
    // clicks a button to filter entries by type
    http.HandleFunc("/entries", func(w http.ResponseWriter, r *http.Request) {
        // Grab the entry type from query params (?type=creed)
        entryType := r.URL.Query().Get("type")

        // Retrieve filtered entries from data source (currently JSON)
        entries, err := GetEntriesByType(entryType)
        if err != nil {
            http.Error(w, "Failed to load entries", http.StatusInternalServerError)
            return
        }

        // Render entries.html, passing in the filtered entries
        err = templates.ExecuteTemplate(w, "entries.html", entries)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    // Log startup message and begin listening on port 8080
    log.Println("🌌 Galactic Archive listening at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server or exit on failure
}

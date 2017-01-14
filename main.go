package main

import (
    "html/template"
    "net/http"
    "fmt"
    "os"
)

var templates *template.Template
var homeTemplate *template.Template

func main() {
    PopulateTemplate()
    http.Handle("/static/", http.FileServer(http.Dir("public")))
    http.HandleFunc("/login/", LoginFunc)
    http.HandleFunc("/", HomeFunc)
    fmt.Println("Server running on 8081")
    http.ListenAndServe("0.0.0.0:8081", nil)
}

// PopulateTemplate reads the ./templates folder and parse all the html files inside it
// and  it stores it in the templates variable which will be looked up by other variables
func PopulateTemplate() {
    templates, err := template.ParseGlob("./templates/*.html")

    if err != nil {
        fmt.Println("main.go: PopulateTemplate: ", err)
        os.Exit(1)
    }

    homeTemplate = templates.Lookup("task.html")
}

// HomeFunc handles the / URL and asks the name of the user in German
func HomeFunc(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        type Task struct {
            Title string
            Content string
            Created string
        }

        type Context struct {
            Tasks []Task
        }

        task1 := Task {
            Title: "Programming",
            Content: "Golang",
            Created: "Mehre",
        }

        task2 := Task {
            Title: "Entertaining",
            Content: "Watching twitch.tv",
            Created: "Mehre 2",
        }

        task3 := Task {
            Title: "Sleeping",
            Content: "Sleep all day",
            Created: "Mehre 3",
        }

        var tasks []Task
        tasks = append(tasks, task1)
        tasks = append(tasks, task2)
        tasks = append(tasks, task3)

        context := Context{Tasks: tasks}

        homeTemplate.Execute(w, context)
    }
}

// LoginFunc handles the /login URL and shows the profile page of the logged in user on a GET request
// handles the Login process on the post request
func LoginFunc(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You are on the profile page")
}

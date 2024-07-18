package main

import (
    "html/template"
    "net/http"
    "sync"
)

type Todo struct {
    Title string
    Done  bool
}

type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}

var tmpl = template.Must(template.ParseFiles("layout.html"))
var todos = []Todo{
    {Title: "Task 1", Done: false},
    {Title: "Task 2", Done: true},
    {Title: "Task 3", Done: true},
}
var mutex = &sync.Mutex{}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos:     todos,
        }
        tmpl.Execute(w, data)
    })

    http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            title := r.FormValue("title")
            if title != "" {
                mutex.Lock()
                todos = append(todos, Todo{Title: title, Done: false})
                mutex.Unlock()
            }
        }
        http.Redirect(w, r, "/", http.StatusFound)
    })

    http.ListenAndServe(":8080", nil)
}

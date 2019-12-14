package main

import ("fmt"
        "net/http")

func hello_world( w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Go Web Hello World!\n")
}

func main() {
        http.HandleFunc("/", hello_world)
        http.ListenAndServe(":8081", nil)
}

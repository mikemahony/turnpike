package turnpike_test

import (
    "github.com/AnimationMentor/turnpike"
    "net/http"
)

func ExampleServer_NewServer() {
    s := turnpike.NewServer()

    http.Handle("/ws", s.Handler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}

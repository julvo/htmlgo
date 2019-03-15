package main

import (
    "fmt"
    "net/http"
    "log"

    h "github.com/julvo/htmlgo"
    a "github.com/julvo/htmlgo/attributes"
)

func main() {
    doc := h.DoctypeHtml5 +
           h.Html(a.Attr(),
              h.Head(a.Attr()),
              h.Body(a.Attr(),
                h.Div(a.Attr(a.Class("columns")),
                  h.Div(a.Attr(a.Class("column", "is-narrow"), a.Data_("url", "https://hello.world")),
                        h.Text("Hello World")),
                  h.Div(a.Attr(a.Class("column")),
                        h.Text(`<script>alert("Text is escaped")</script>`)),
                  h.Div(a.Attr(a.Class("column"))))))
    fmt.Println(doc)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, doc)
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}

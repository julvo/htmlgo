package main

import (
	"log"
	"net/http"

	. "github.com/julvo/htmlgo"
)

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fruit := []string{"Apple", "Banana", "Orange"}

	fruitListItems := Nodes()
	for _, f := range fruit {
		fruitListItems.Append(Li_(Text(f)))
	}

	content :=
		Nodes(
			navbar(false),
			Ul_(fruitListItems),
			footer())

	page("home", content).RenderTo(w)
}

func page(title string, content Node) Node {
	p :=
		Document(
			Doctype("html"),
			Html_(
				Head_(
					Title_(Text(title)),
					Meta(Charset("utf-8")),
					Meta(Name("viewport").Content("width=device-width").InitialScale("1"))),
				Body_(
					Div(Class("is-size-6").Style_("background:grey;")),
					content)))

	return p
}

func navbar(isLoggedIn bool) Node {
	navItems := Nodes()
	if !isLoggedIn {
		navItems.Append(A(Href("/login"), Text("Login")))
	}

	nav :=
		Nav_(
			Div_(navItems),
			Hr_())

	return nav
}

func footer() Node {
	return Footer_(
		Hr_(),
		Text(HTML("&copy Acme Ltd, 2019")))
}

package main

import (
	"log"
	"net/http"

	. "github.com/julvo/htmlgo"
	a "github.com/julvo/htmlgo/attributes"
)

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fruit := []string{"Apple", "Banana", "Orange"}

	fruitListItems := HTML("")
	for _, f := range fruit {
		fruitListItems += Li_(Text(f))
	}

	content :=
		navbar(false) +
			Ul_(fruitListItems) +
			footer()

	WriteTo(w, page("Home", content))
}

func page(title string, content HTML) HTML {
	p :=
		Html5_(
			Head_(
				Title_(Text(title)),
				Meta(Attr(a.Charset_("utf-8"))),
				Meta(Attr(a.Name_("viewport"), a.Content_("width=device-width"), a.InitialScale_("1"))),
				Link(Attr(a.Rel_("stylesheet"), a.Href_("/static/css/main.min.css")))),
			Body_(
				content,
				Script(Attr(a.Src_("/static/js/main.min.js")), JS{})))

	return p
}

func navbar(isLoggedIn bool) HTML {
	var navItems HTML
	if !isLoggedIn {
		navItems = A(Attr(a.Href_("/login")), Text_("Login"))
	}

	nav :=
		Nav_(
			Div_(navItems),
			Hr_())

	return nav
}

func footer() HTML {
	return Footer_(
		Hr_(),
		Text_("&copy Acme Ltd, 2019"))
}

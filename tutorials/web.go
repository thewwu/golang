package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is index page")

	// http format
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
	fmt.Fprintf(w, "<p>Go is fast!</p>")
	fmt.Fprintf(w, "<p>...and simple!</p>")
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong>")

	// Single line
	fmt.Fprintf(w, `
		<h6>You can even do ...</h6>
		<h5>multiple lines ...</h5>
		<h4>in one %s</h4>`, "formatted print")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about page")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	t, err := template.ParseFiles("tutorials/basictemplating.html")
	fmt.Println(err)
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}

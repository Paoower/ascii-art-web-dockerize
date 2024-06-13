package main

import (
	asciiart "ascii-art-web/ascii-art"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var result string

const RED = "\033[31;1m"
const GREEN = "\033[32;1m"
const YELLOW = "\033[33;1m"
const NONE = "\033[0m"

type PageData struct {
	Lines string
}

type ErrorData struct {
	Error error
}

func main() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("ascii-web-server:")

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii", postHandler)
	http.HandleFunc("/download", download)

	log.Println(GREEN, "Server started at http://localhost:8080", NONE)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Check correct GET method
	if r.Method != "GET" {
		log.Printf("%v Bad request %v on %v page%v\n", RED, r.Method, r.URL.Path, NONE)
		badRequestHandler(w)
		return
	}

	if r.URL.Path != "/" {
		log.Printf("%v Tried to access unexistant route %v%v\n", YELLOW, r.URL.Path, NONE)
		notFoundHandler(w)
		return
	}

	// Parse the template for the home page
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("%v Error parsing home template: %v%v", RED, err, NONE)
		internalServerErrorHandler(w, err)
		return
	}

	data := PageData{
		Lines: result,
	}

	// Execute the template with the provided data
	err = t.Execute(w, data)
	if err != nil {
		log.Printf("%v Error executing home template: %v%v", RED, err, NONE)
		internalServerErrorHandler(w, err)
		return
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	//Check correct POST method
	if r.Method != "POST" {
		log.Printf("%v Bad request %v on %v page%v\n", RED, r.Method, r.URL.Path, NONE)
		badRequestHandler(w)
		return
	}

	// Parse form values
	err := r.ParseForm()
	if err != nil {
		log.Printf("%v Error parsing data form: %v%v", RED, err, NONE)
		internalServerErrorHandler(w, err)
		return
	}

	// Retrieve input and banner style from the form
	input := r.FormValue("input")
	style := r.FormValue("banner")

	input = strings.Replace(input, "\r\n", "\n", -1)

	if style == "" {
		log.Printf("%v No banner provided: style: %s%v\n", RED, style, NONE)
		internalServerErrorHandler(w, err)
		return
	}

	// Generate ASCII art based on input and style
	output, err := asciiart.GetAscii(input, style)
	if err != nil {
		internalServerErrorHandler(w, err)
		return
	}
	result = "\n" + output

	log.Printf("%v POST request on /ascii successful %v", GREEN, NONE)

	// Redirect to the home page after successful processing
	http.Redirect(w, r, "/", http.StatusFound)
}

func download(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%v Bad request %v on %v page%v\n", RED, r.Method, r.URL.Path, NONE)
		badRequestHandler(w)
		return
	}

	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Disposition", "attachment; filename=output.txt")
	w.Header().Set("Content-Length", strconv.Itoa(len(result)))
	w.Write([]byte(strings.Replace(result, "&nbsp;", " ", -1)))
	result = ""
}

func notFoundHandler(w http.ResponseWriter) {
	// Send 404 code
	w.WriteHeader(http.StatusNotFound)

	// Parse the 404 template
	t, err := template.ParseFiles("templates/404.html")
	if err != nil {
		log.Printf("%v Error executing 404 template: %v%v", RED, err, NONE)
		internalServerErrorHandler(w, err)
		return
	}

	// Execute the 404 template
	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("%v Error executing 404 template: %v%v", RED, err, NONE)
		internalServerErrorHandler(w, err)
		return
	}
}

func badRequestHandler(w http.ResponseWriter) {
	// Send 400 code
	w.WriteHeader(http.StatusBadRequest)

	// Parse the 400 template
	t, err := template.ParseFiles("templates/400.html")
	if err != nil {
		log.Printf("%v Error executing 400 template: %v%v", RED, err, NONE)
		internalServerErrorHandler(w, err)
		return
	}

	// Execute the 400 template
	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("%v Error executing 400 template: %v%v", RED, err, NONE)
		internalServerErrorHandler(w, err)
		return
	}
}

func internalServerErrorHandler(w http.ResponseWriter, erro error) {
	// Send 500 code
	w.WriteHeader(http.StatusInternalServerError)

	// Parse the 500 template
	t, err := template.ParseFiles("templates/500.html")
	if err != nil {
		log.Printf("%v Error executing 500 template: %v%v", RED, err, NONE)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	errorData := ErrorData{
		Error: erro,
	}

	// Execute the 500 template
	err = t.Execute(w, errorData)
	if err != nil {
		log.Printf("Error executing 500 template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

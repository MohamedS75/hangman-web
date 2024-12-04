package main

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
)

type Game struct {
	Level    string
	Historic string
}

var currentGame = Game{}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Printf("Error: handler for %s not found\n", html.EscapeString(r.URL.Path))
		return
	}

	fmt.Printf("got / request\n")

	tmpl := template.Must(template.ParseFiles("../template/index.html"))
	var i = struct {
		Word   string
		ToFind string
	}{
		Word:   "O________",
		ToFind: "Ordinateur",
	}
	r.ParseForm()
	L := r.Form.Get("input")
	fmt.Printf("Received input: %s\n", L)
	err := tmpl.Execute(w, i)
	if err != nil {
		panic(err)
	}
}

func hangman(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("level") != "" {
		currentGame = Game{Level: r.URL.Query().Get("level"), Historic: ""}
	}

	wordToGuess := "Ordinateur"

	r.ParseForm()
	letter := r.Form.Get("letter")
	if letter != "" {
		currentGame.Historic += letter + " "
	}

	maskedWord := maskWord(wordToGuess, currentGame.Historic)

	tmpl := template.Must(template.ParseFiles("../template/hangman.html"))
	var i = struct {
		Level    string
		Word     string
		Historic string
	}{
		Level:    currentGame.Level,
		Word:     maskedWord,
		Historic: currentGame.Historic,
	}

	err := tmpl.Execute(w, i)
	if err != nil {
		panic(err)
	}
}

func maskWord(word, historic string) string {
	masked := ""
	for _, char := range word {
		if containsRune(historic, char) {
			masked += string(char)
		} else {
			masked += "_"
		}
	}
	return masked
}

func containsRune(historic string, r rune) bool {
	for _, char := range historic {
		if char == r {
			return true
		}
	}
	return false
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("../template/css"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/hangman", hangman)

	http.ListenAndServe(":8080", nil)
}

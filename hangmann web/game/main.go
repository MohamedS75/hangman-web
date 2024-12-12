package main

import (
	"bufio"
	"fmt"
	"html"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Game struct {
	Level       string
	Historic    string
	WordToGuess string
	Attempts    int
	MaxAttempts int
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
		ToFind: "Mot à trouver",
	}
	r.ParseForm()
	L := r.Form.Get("input")
	fmt.Printf("Received input: %s\n", L)
	err := tmpl.Execute(w, i)
	if err != nil {
		panic(err)
	}
}

func loadWords(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

func getRandomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func hangman(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Get("NewGame") != "" {
		var words = []string{}
		if r.URL.Query().Get("NewGame") == "facile" {
			words, _ = loadWords("words.txt")
		} else if r.URL.Query().Get("NewGame") == "moyen" {
			words, _ = loadWords("words2.txt")
		} else if r.URL.Query().Get("NewGame") == "difficile" {
			words, _ = loadWords("words3.txt")

		}

		currentGame = Game{
			WordToGuess: getRandomWord(words),
			Attempts:    0,
			MaxAttempts: 10,
		}
	}

	if currentGame.WordToGuess == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()
	letter := r.Form.Get("letter")
	if letter != "" {
		currentGame.Historic += letter + " "
		currentGame.Attempts++
	}

	maskedWord := maskWord(currentGame.WordToGuess, currentGame.Historic)

	if maskedWord == currentGame.WordToGuess {
		http.Redirect(w, r, "/gameover?result=win", http.StatusFound)
		return
	}

	if currentGame.Attempts >= currentGame.MaxAttempts {
		http.Redirect(w, r, "/gameover?result=lose", http.StatusFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("../template/hangman.html"))
	var data = struct {
		Level       string
		Word        string
		Historic    string
		Attempts    int
		MaxAttempts int
		WordToGuess string
		Url 		  string
	}{
		Level:       currentGame.Level,
		Word:        maskedWord,
		Historic:    currentGame.Historic,
		Attempts:    currentGame.Attempts,
		MaxAttempts: currentGame.MaxAttempts,
		WordToGuess: currentGame.WordToGuess,
	}
	if data.Attempts == 0 {
		data.Url = "image/0.png"
	 } else if data.Attempts == 1 {
		data.Url = "image/1.png"
	 } else if data.Attempts == 2 {
		data.Url = "image/2.png"
	 } else if data.Attempts == 3 {
		data.Url = "image/3.png"
	 } else if data.Attempts == 4 {
		data.Url = "image/4.png"
	 } else if data.Attempts == 5 {
		data.Url = "image/5.png"
	 } else if data.Attempts == 6 {
		data.Url = "image/6.png"
	 } else if data.Attempts == 7 {
		data.Url = "image/7.png"
	 } else if data.Attempts == 8 {
		data.Url = "image/8.png"
	 } else if data.Attempts == 9 {
		data.Url = "image/9.png"
	 } else if data.Attempts == 10 {
		data.Url = "image/10.png"
	 }
	 err := tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func gameOver(w http.ResponseWriter, r *http.Request) {
	result := r.URL.Query().Get("result")
	message := ""

	if result == "win" {
		message = "Félicitations ! Vous avez gagné !"
	} else if result == "lose" {
		message = "Dommage ! Vous avez perdu !"
	}

	tmpl := template.Must(template.ParseFiles("../template/gameover.html"))
	err := tmpl.Execute(w, struct {
		Message     string
		WordToGuess string
	}{
		Message:     message,
		WordToGuess: currentGame.WordToGuess,
	})
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
	http.Handle("/image/", http.StripPrefix("/image", http.FileServer(http.Dir("../template/image"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/hangman", hangman)
	http.HandleFunc("/gameover", gameOver)

	http.ListenAndServe(":8080", nil)
}

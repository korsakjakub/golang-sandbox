package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

var tiles = [9]string{"", "", "", "", "", "", "", "", ""}
var turn = "X"

func reset() {
	for i := 0; i < 8; i++ {
		tiles[i] = ""
	}
	turn = "X"
}

func toggleTurn(turn string) string {
	if turn == "X" {
		return "O"
	}
	return "X"
}

func main() {
	title := "Tic Tac Toe"

	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(base(title, tiles, turn)))

	mux.Handle("/reset", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}
		reset()
		w.WriteHeader(http.StatusOK)
		base(title, tiles, turn).Render(context.Background(), w)
	}))

	mux.Handle("/move/{position}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}
		position := r.PathValue("position")
		fmt.Println("position: " + position)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(position))
	}))

	log.Fatal(http.ListenAndServe(":4000", mux))
}

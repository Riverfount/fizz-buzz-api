package server

import "github.com/Riverfount/fizz-buzz-api/internal/handler"

func (a *App) RegisterRoutes() {
	a.Mux.HandleFunc("/hello", handler.HelloHandler)
	a.Mux.HandleFunc("/fizzbuzz", handler.FizzBuzzHandler)
}

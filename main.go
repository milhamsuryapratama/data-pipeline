package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello-world", func(writer http.ResponseWriter, request *http.Request) {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

		log.Debug().
			Str("Scale", "833 cents").
			Float64("Interval", 833.09).
			Msg("Fibonacci is everywhere")
	})

	err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), mux)
	if err != nil {
		panic(err)
	}
}

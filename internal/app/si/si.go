package si

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// API start a web server to expose one route for State Integration.
// See more here: https://developer.valvesoftware.com/wiki/Counter-Strike:_Global_Offensive_Game_State_Integration
func API(port, token string) {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := mux.NewRouter()

	// Logging.
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		})
	},
	)

	// Auth.
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// TODO: fix this. Dont work very well.
			// a := auth{}
			// err := json.NewDecoder(r.Body).Decode(&a)
			// if err != nil {
			// 	log.Println(err.Error())
			// 	w.WriteHeader(http.StatusInternalServerError)
			// 	return
			// }

			// if a.Auth.Token != token {
			// 	w.WriteHeader(http.StatusUnauthorized)
			// 	return
			// }

			next.ServeHTTP(w, r)
		})
	},
	)

	// Service route.
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}

		ob := observer{}
		err = json.Unmarshal(data, &ob)
		if err != nil && !errors.Is(err, io.EOF) {
			log.Println(err.Error())
		}

		//fmt.Println(os.Player.Name)
		log.Printf("data: %s", string(data))

	}).Methods("POST")

	srv := &http.Server{
		Addr: "0.0.0.0:" + port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf(srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

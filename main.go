package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	var listenAddr string

	flag.StringVar(&listenAddr, "listen-addr", ":5000", "server listen address")
	flag.Parse()

	db := sqlx.MustConnect("mysql", "root:secret@tcp(localhost:3306)/erp")

	userMgr := UserMgr{
		db: db,
	}

	frontend := Frontend{
		userMgr: userMgr,
	}

	admin := Admin{
		userMgr: userMgr,
	}

	router := http.NewServeMux()
	router.HandleFunc("/", frontend.index)
	router.HandleFunc("/timeline", frontend.timeline)
	router.HandleFunc("/register", frontend.register)
	router.HandleFunc("/login", frontend.login)

	router.HandleFunc("/admin", admin.index)
	router.HandleFunc("/admin/employee/new", admin.employeeNew)
	router.HandleFunc("/admin/employee/list", admin.employeeList)

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	log.Println("Server is ready to handle requests at", listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	log.Println("Server stopped")
}

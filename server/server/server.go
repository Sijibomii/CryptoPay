package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anthdm/hollywood/actor"
	"github.com/gorilla/mux"
	"github.com/sijibomii/cryptopay/config"
	"github.com/sijibomii/cryptopay/core/db"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/core/utils"
	"github.com/sijibomii/cryptopay/server/controllers"
	"github.com/sijibomii/cryptopay/server/mailer"
	"github.com/sijibomii/cryptopay/server/util"
)

// posgress connection string should come in config
func Run(config config.Config) {
	r := mux.NewRouter()

	pg := *initPool(config.Postgres, 10)

	pg.DB.AutoMigrate(&models.User{}, &models.Store{}, &models.ClientToken{}, &models.Session{})

	// define app state

	// spawn up db client
	e := actor.NewEngine()
	pid := e.Spawn(newDbClient(pg), "dbClient")

	mailerPid := e.Spawn(newMailerClient(&config.Mailer), "mailer")

	appState := &util.AppState{
		Postgres:   pid,
		PgExecutor: pg,
		Config:     &config,
		Engine:     e,
		Mailer:     mailerPid,
	}

	// routes register
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.IndexHandler(w, r, appState)
	}).Methods("GET")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.LoginHandler(w, r, appState)
	}).Methods("POST")

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		controllers.RegisterHandler(w, r, appState)
	}).Methods("POST")

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler: r,
	}

	log.Printf("Server listening on %s:%d\n", config.Server.Host, config.Server.Port)
	log.Fatal(server.ListenAndServe())
}

func newMailerClient(m *config.MailerConfig) actor.Producer {
	return func() actor.Receiver {
		return &mailer.Mailer{
			SmtpHost:     m.SmtpHost,
			SmtpPort:     m.SmtpPort,
			SmtpUsername: m.SmtpUsername,
			SmtpPassword: m.SmtpPassword,
		}
	}
}

func newDbClient(DB utils.PgExecutor) actor.Producer {
	return func() actor.Receiver {
		return &db.DBClient{
			PgExecutor: DB,
		}
	}
}

func initPool(connection string, pool_size int) *utils.PgExecutor {
	// Initialize the connection pool
	pool, err := utils.InitPool(connection, pool_size)
	if err != nil {
		fmt.Println(connection)
		panic(err)
	}

	// Create an instance of PgExecutor with the connection pool
	executor := &utils.PgExecutor{DB: pool}

	return executor
	// // Use the executor to perform database operations
	// // For example:
	// // executor.GetDB().Create(&YourModel{Field1: "Value1", Field2: "Value2"})

	// // Close the connection pool when done
	// defer DB.Close()
}

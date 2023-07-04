package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anthdm/hollywood/actor"
	"github.com/gorilla/mux"
	btc_processor "github.com/sijibomii/cryptopay/block_processor/bitcoin"
	"github.com/sijibomii/cryptopay/blockchain_client/bitcoin"
	coinclient "github.com/sijibomii/cryptopay/coin_client"
	"github.com/sijibomii/cryptopay/config"
	"github.com/sijibomii/cryptopay/core/db"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/core/utils"
	"github.com/sijibomii/cryptopay/server/controllers"
	"github.com/sijibomii/cryptopay/server/mailer"
	"github.com/sijibomii/cryptopay/server/middleware"
	"github.com/sijibomii/cryptopay/server/util"
)

// posgress connection string should come in config
func Run(config config.Config) {
	r := mux.NewRouter()

	pg := *initPool(config.Postgres, 10)

	pg.DB.AutoMigrate(&models.User{}, &models.Store{}, &models.ClientToken{}, &models.Session{}, models.Payment{}, models.Payout{}, models.BtcBlockChainStatus{})

	e := actor.NewEngine()

	pid := e.Spawn(newDbClient(pg), "dbClient")

	mailerPid := e.Spawn(newMailerClient(&config.Mailer), "mailer")

	// currencyClient
	value := os.Getenv("COIN_API_KEY")
	client_url := os.Getenv("BITCOIN_CHAIN_URL")
	coinClientPid := e.Spawn(newCoinClient(value), "coinClient")

	btcChainClient := e.Spawn(newBtcChainClient(client_url), "btcChainClient")

	processorClient := e.Spawn(newProcessorClient("testnet", pid, e, btcChainClient), "processorClient")

	pollerClient := e.Spawn(newPollerClient("testnet", pid, btcChainClient, processorClient), "pollerClient")

	// send message
	e.Send(pollerClient, btc_processor.StartPollingMessage{
		Ignore_previous_blocks: true,
	})

	pendingPollerClient := e.Spawn(newPendingPollerClient("testnet", pid, btcChainClient, processorClient), "pollerClient")
	e.Send(pendingPollerClient, btc_processor.StartPBPollingMessage{})

	appState := &util.AppState{
		Postgres:        pid,
		PgExecutor:      pg,
		Config:          &config,
		Engine:          e,
		Mailer:          mailerPid,
		CoinClient:      coinClientPid,
		BtcClient:       btcChainClient,
		ProcessorClient: processorClient,
		PollerClient:    pollerClient,
		PBPollerClient:  pendingPollerClient,
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

	// activate

	r.HandleFunc("/reset_password", func(w http.ResponseWriter, r *http.Request) {
		controllers.ResetPasswordHandler(w, r, appState)
	}).Methods("POST")

	r.HandleFunc("/change_password", func(w http.ResponseWriter, r *http.Request) {
		controllers.ChangePasswordHandler(w, r, appState)
	}).Methods("POST")

	// protected routes
	secureRoutes := r.PathPrefix("/").Subrouter()
	secureRoutes.Use(middleware.AuthMiddleware(appState))

	secureRoutes.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		controllers.ProfileHandler(w, r, appState)
	}).Methods("GET")

	secureRoutes.HandleFunc("/client_tokens", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllClientTokensHandler(w, r, appState)
	}).Methods("GET")

	secureRoutes.HandleFunc("/client_tokens", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateClientTokensHandler(w, r, appState)
	}).Methods("POST")

	secureRoutes.HandleFunc("/client_tokens/:id", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetClientTokenByIdHandler(w, r, appState)
	}).Methods("GET")

	secureRoutes.HandleFunc("/stores", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetStoresList(w, r, appState)
	}).Methods("GET")

	secureRoutes.HandleFunc("/stores", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateStores(w, r, appState)
	}).Methods("POST")

	// /stores/{id}
	secureRoutes.HandleFunc("/stores/:id", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetStoresById(w, r, appState)
	}).Methods("GET")

	secureRoutes.HandleFunc("/stores/:id", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateStoresById(w, r, appState)
	}).Methods("PATCH")

	secureRoutes.HandleFunc("/stores/:id", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteStoreById(w, r, appState)
	}).Methods("DELETE")

	// add separate route for addition of address

	paymentRoutes := r.PathPrefix("/payments").Subrouter()

	paymentRoutes.Use(middleware.PaymentMiddleware(appState))

	paymentRoutes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreatePayment(w, r, appState)
	}).Methods("POST")

	paymentRoutes.HandleFunc("/payments/:id/status", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetPaymentStatus(w, r, appState)
	}).Methods("GET")

	// /vouchers

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler: r,
	}

	log.Printf("Server listening on %s:%d\n", config.Server.Host, config.Server.Port)
	log.Fatal(server.ListenAndServe())
}

func newPendingPollerClient(network string, postgresClient, btcClient, processorClient *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &btc_processor.PBPoller{
			PostgresClient: postgresClient,
			Network:        network,
			BlockProcessor: processorClient,
			BtcClient:      btcClient,
		}
	}
}

func newPollerClient(network string, postgresClient, btcClient, processorClient *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &btc_processor.Poller{
			PostgresClient: postgresClient,
			Network:        network,
			BlockProcessor: processorClient,
			BtcClient:      btcClient,
		}
	}
}

func newProcessorClient(network string, postgresClient *actor.PID, engine *actor.Engine, btcClient *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &btc_processor.Processor{
			PostgresClient: postgresClient,
			Network:        network,
			Engine:         engine,
			BtcClient:      btcClient,
		}
	}
}

func newBtcChainClient(url_string string) actor.Producer {
	return func() actor.Receiver {
		return &bitcoin.BlockchainClient{
			BSUrl: "https://blockstream.info/api",
		}
	}
}

func newCoinClient(api_key string) actor.Producer {
	return func() actor.Receiver {
		return &coinclient.CoinClient{
			Key: api_key,
		}
	}
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

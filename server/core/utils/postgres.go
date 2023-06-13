package utils

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PgExecutorAddr = actor.PID

type Client struct {
	Conn      *gorm.DB
	ServerPID *actor.PID
}

var (
	DB *gorm.DB
)

type PgExecutor struct {
	DB *gorm.DB
}

func InitDB(url string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitPool(url string, maxConnections int) (*gorm.DB, error) {
	db, err := InitDB(url)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxOpenConns(maxConnections)
	db.LogMode(true)

	DB = db
	return DB, nil
}

func (p *PgExecutor) GetDB() *gorm.DB {
	return p.DB
}

// func main() {
// 	// Initialize the connection pool
// 	pool, err := InitPool("your-connection-string", 10)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Create an instance of PgExecutor with the connection pool
// 	executor := &PgExecutor{DB: pool}

// 	// Use the executor to perform database operations
// 	// For example:
// 	// executor.GetDB().Create(&YourModel{Field1: "Value1", Field2: "Value2"})

// 	// Close the connection pool when done
// 	defer DB.Close()
// }

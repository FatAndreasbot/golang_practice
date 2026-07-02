package t1

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	connection  *sql.DB
	initialyzer sync.Once
}

func (db *Database) GetConnection() *sql.DB {
	var err error
	db.initialyzer.Do(func() {
		db.connection, err = sql.Open("sqlite3", "../../concurrency/t_2/data.db")
		// printng just to see if the conn recreares
		fmt.Println("connecting")
	})

	if err != nil {
		// panics in production is a bad habit, but i feel it makes sense in server init
		panic(err)
	}
	return db.connection
}

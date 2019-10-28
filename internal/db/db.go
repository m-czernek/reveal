package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/reveal/internal/config"
)

var db *sql.DB
var doOnce sync.Once

func initDbPool() {
	var err error
	db, err = sql.Open("mysql", config.ReadConfigFromEnv().GetFullDatabaseUrl())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func GetDatabase() *sql.DB {
	doOnce.Do(initDbPool)
	return db
}

// type Salary struct {
// 	ID           int    `json:"id"`
// 	YearlySalary int    `json:"yearly_salary"`
// 	Currency     string `json:"currency"`
// }

// func (s Salary) String() string {
// 	return fmt.Sprintf("{ ID: %d, yearlySalary: %d, currency: %s }", s.ID, s.YearlySalary, s.Currency)
// }

// func connectToDatabase() {
// 	db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/backend")
// 	if err != nil {
// 		log.Fatal(err)
// 		panic(err.Error())
// 	}

// 	defer db.Close()

// 	insert, err := db.Query("INSERT INTO salary(yearly_salary, currency) VALUES ( 100000, 'USD' )")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer insert.Close()
// 	// log.Println("Inserted: " + insert)

// 	results, err := db.Query("SELECT id, yearly_salary, currency FROM salary")
// 	if err != nil {
// 		panic(err.Error()) // proper error handling instead of panic in your app
// 	}

// 	for results.Next() {
// 		var salary Salary
// 		// for each row, scan the result into our tag composite object
// 		err = results.Scan(&salary.ID, &salary.YearlySalary, &salary.Currency)
// 		if err != nil {
// 			panic(err.Error()) // proper error handling instead of panic in your app
// 		}
// 		// and then print out the tag's Name attribute
// 		log.Printf(salary.String())
// 	}
// }

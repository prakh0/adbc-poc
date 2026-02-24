package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/arrow-adbc/go/adbc/drivermgr"
)

func main() {
	var drv drivermgr.Driver

	db, err := drv.NewDatabase(map[string]string{
		"driver": "mysql",
		"uri":    "root:MySQL_pass01@tcp(34.28.25.89:3306)/ettaflow",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	conn, err := db.Open(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	stmt, err := conn.NewStatement()
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.SetSqlQuery("SELECT * FROM airportdb.booking")
	stmt.SetOption("adbc.statement.batch_size", "250000")
	startTime := time.Now()
	result, _, err := stmt.ExecuteQuery(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer result.Release()
	batchNum := 0
	for result.Next() {
		batchNum++
		batch := result.RecordBatch()
		fmt.Printf("Batch #%d: %d rows\n", batchNum, batch.NumRows())
		batch.Release()
	}
	fmt.Printf("Total time for reading data: %v\n", time.Since(startTime))
	fmt.Println("Completed successfully!")
}

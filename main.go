package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/arrow-adbc/go/adbc"
	"github.com/apache/arrow-adbc/go/adbc/drivermgr"
)

func main() {
	var drv drivermgr.Driver
	uri := os.Getenv("DB_URI")
	if uri == "" {
		log.Fatal("DB_URI environment variable is not set")
	}

	db, err := drv.NewDatabase(map[string]string{
		"driver": "mysql",
		"uri":    uri,
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

	ingestConn, err := db.Open(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer ingestConn.Close()

	rowInserted, err := adbc.IngestStream(
		context.Background(),
		ingestConn,
		result,
		"booking_copy",
		"adbc.ingest.mode.replace",
		adbc.IngestStreamOptions{
			Extra: map[string]string{
				"adbc.statement.ingest.batch_size": "12000",
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of rows inserted: %d\n", rowInserted)
	fmt.Printf("Total time for reading data: %v\n", time.Since(startTime))
	fmt.Println("Completed successfully!")
}

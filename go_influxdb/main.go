package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/influxdata/influxdb-client-go"
)

func main() {
	client := influxdb2.NewClient("http://54.219.21.24:8086", "eh2GDwqhbF-gbCEmkrAQROc3lG5efcXGCkx5XmCYm4hvpnYSk5N-8t9Y5cT6PVNWJREcR3La-f2qw8uaXiOZjg==")
	defer client.Close()

	queryAPI := client.QueryAPI("coinsky")
	query := `from(bucket: "nft_mint_tracker")
  |> range(start: -1h)
  |> filter(fn: (r) => r["_measurement"] == "eth")
  |> filter(fn: (r) => r["_field"] == "mints")
  |> count()
  |> yield(name: "count")`

	// get QueryTableResult
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		result, err := queryAPI.Query(context.Background(), query)
		if err != nil {
			panic(err)
		}

		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %\n", result.Err().Error())
		}

		defer wg.Done()
	}()
	wg.Wait()
}

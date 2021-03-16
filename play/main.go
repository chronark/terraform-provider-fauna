package main

import (
	"log"
	"os"
	// "strconv"
	f "github.com/fauna/faunadb-go/v3/faunadb"
)

func main() {

	client := f.NewFaunaClient(os.Getenv("FAUNA_KEY"))

	res, err := client.Query(f.Update(f.Database("db2"), f.Obj{
		"name": "db2",
	}))

	if err != nil {
		panic(err)
	}

	log.Println(res)

}

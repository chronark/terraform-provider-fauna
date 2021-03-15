package main

import (
	f "github.com/fauna/faunadb-go/v3/faunadb"
	"log"
)

func main() {
	client := f.NewFaunaClient("")
	res, err := client.Query(f.CreateCollection(f.Obj{"name": "x"}))
	if err != nil {
		panic(err)
	}
	log.Println(res)

}

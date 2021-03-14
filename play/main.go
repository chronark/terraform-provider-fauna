package main

import (
	"fmt"
	f "github.com/fauna/faunadb-go/v3/faunadb"
)

func main() {
	client := f.NewFaunaClient("fnAEESEtKbACB_YVroJX7uHMsJlzUUVXrQIkq0x8")
	res, err := client.Query(f.CreateCollection(f.Obj{"name": "x"}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}

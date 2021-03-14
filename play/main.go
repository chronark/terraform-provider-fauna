package main

import (
	"fmt"
	f "github.com/fauna/faunadb-go/v3/faunadb"
)

type Database struct {
	Ref f.RefV `fauna:"RefV"`
}

type Collection struct {
	Name        string `fauna:"name"`
	Ts          int64  `fauna:"ts"`
	HistoryDays int64  `fauna:"history_days"`
	TTLDays     int64  `fauna:"ttl_days"`
}

func main() {
	client := f.NewFaunaClient("fnAEESEtKbACB_YVroJX7uHMsJlzUUVXrQIkq0x8")
	res, err := client.Query(f.CreateCollection(f.Obj{"name": "x"}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}

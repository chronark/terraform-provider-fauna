package main

import (
	"fmt"
	f "github.com/fauna/faunadb-go/v3/faunadb"
)

type Database struct {
	Ref f.RefV `fauna:"RefV"`
}

func main() {
	client := f.NewFaunaClient("fnAEEL4dXVACBy4dTtIQpSfO5U4t6rccY2gLO7Da")
	res, err := client.Query(f.Paginate(f.Databases()))
	if err != nil {
		panic(err)
	}
	var refs []f.RefV
	err = res.At(f.ObjKey("data")).Get(&refs)
	if err != nil {
		panic(err)
	}
	databases := make([]string, 0)
	for _, db := range refs {
		databases = append(databases, db.ID)
	}
	fmt.Println(databases)
}

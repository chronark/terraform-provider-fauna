package main

import (
	"log"
	"os"
	// "strconv"
	f "github.com/fauna/faunadb-go/v3/faunadb"
)

func getName(raw map[string]interface{}) string {
	data := raw["data"]
	if data == nil {
		return ""
	}

	name, err := raw["data"].(f.ObjectV).At(f.ObjKey("name")).GetValue()
	if err != nil {
		return ""
	}
	return name.String()

}

func flattenKey(raw map[string]interface{}) (map[string]interface{}, error) {
	flatKey := make(map[string]interface{})

	flatKey["id"] = raw["ref"].(f.RefV).ID
	flatKey["name"] = getName(raw)
	flatKey["ts"] = raw["ts"]
	flatKey["role"] = raw["role"]
	flatKey["hashed_secret"] = raw["hashed_secret"]

	return flatKey, nil
}

func main() {
	client := f.NewFaunaClient(os.Getenv("FAUNA_KEY"))
	res, err := client.Query(f.Map(f.Paginate(f.Keys(), f.Size(10000)), f.Lambda("key", f.Get(f.Var("key")))))
	if err != nil {
		panic(err)
	}
	keys := make([]map[string]interface{}, 0)
	err = res.At(f.ObjKey("data")).Get(&keys)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(keys); i++ {
		flatKey, err := flattenKey(keys[i])
		if err != nil {
			panic(err)
		}
		keys[i] = flatKey

	}
	log.Println(keys)

}

// [map[
// 	data:Obj{
// 		"name": "tf"
// 		}
// 	hashed_secret:"$2a$05$qu8BT.oc1CH3pIfq8vNwB.uslwf2OLDDrypQiLwebbPdNF2dSIaf."
// 	ref: RefV{ID: "293128014041973255", Collection: &RefV{ID: "keys"}}
// 	role:"admin"
// 	ts:1615807508447000]
// map[data:Obj{"name": "tmp"} hashed_secret:"$2a$05$xxIssZ1k7FqCURcuV8perO/RY6ONnt3KazKwdCdRfZO4zNAIDCf1W" ref:RefV{ID: "293146609145872901", Collection: &RefV{ID: "keys"}} role:"admin" ts:1615825242110000]]

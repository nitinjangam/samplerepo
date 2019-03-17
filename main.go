package main

import (
	"fmt"

	"gopkg.in/couchbase/gocb.v1"
)

type user struct {
	name      string   `json: "uid"`
	email     string   `json:"email"`
	interests []string `json:"interests"`
}

type person struct {
	fName string `json: "firstName"`
	lName string `json:"lastName"`
	cNo   int    `json:"contactNo"`
	eId   string `json:"emailId"`
	add   string `json:"address"`
	pCode int    `json:"pinCode"`
}

func main() {
	cluster, err := gocb.Connect("couchbase://localhost")
	fmt.Println(err)
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "Chmod777nit",
	})
	bucket, _ := cluster.OpenBucket("newDb2Bucket", "")
	bucket.Manager("", "").CreatePrimaryIndex("", true, false)
	bucket.Upsert("nitinjangam1",
		user{
			name:      "nitin",
			email:     "njangam169@gmail.com",
			interests: []string{"nitin", "jangam"},
		}, 0)
	var inuser interface{}
	bucket.Get("nitinjangam", &inuser)
	fmt.Printf("User: %v\n", inuser)
	newpc := 400706
	query := gocb.NewN1qlQuery("SELECT * FROM newDb2Bucket WHERE $1 = pinCode")
	rows, _ := bucket.ExecuteN1qlQuery(query,
		[]interface{}{newpc})
	var row interface{}
	for rows.Next(&row) {
		fmt.Printf("Row: %v\n", row)
	}
}

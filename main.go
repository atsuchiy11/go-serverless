package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{
		Region:   aws.String("ap-northeast-1"),
		Endpoint: aws.String("http://localhost:8000"),
	})
	table := db.Table("sls-user-table")

	// Put item
	u := User{ID: 1, Name: "John", Email: "john@example.com"}
	err := table.Put(u).Run()

	// Get item
	var result User
	err = table.Get("ID", u.ID).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

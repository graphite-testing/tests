package main

import (
	"github.com/surrealdb/surrealdb.go"
)

type User struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
	Bff  string `json:"bff"`
}

type Car struct {
	Id    string `json:"id,omitempty"`
	Owner string `json:"owner"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

func main() {
	db := connect()

	CreateUser(db)
	GetUsers(db)
	CreateCar(db, "cars:volvo", Car{
		Owner: "users:bob",
		Brand: "Volvo",
		Model: "V70",
	})
}

func connect() *surrealdb.DB {
	db, err := surrealdb.New("ws://localhost:8000/rpc")
	if err != nil {
		panic(err)
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": "jocke",
		"pass": "test",
	}); err != nil {
		panic(err)
	}

	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}

	return db
}

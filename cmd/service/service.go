package main

import (
	"DB_APPS/pkg/storage"
	"DB_APPS/pkg/storage/membd"
	"DB_APPS/pkg/storage/postgres"
	"fmt"
	"log"
	"os"
)

var db storage.Interface

func main() {
	var err error
	pwd := os.Getenv("postgres")
	if pwd == "" {
		os.Exit(1)
	}
	constr :=
		"postgresql://" + pwd + "@" + "192.168.1.165" + "/tasks"
	db, err := postgres.New(constr)
	if err != nil {
		log.Fatal(err)
	}
	db = membd.DB{}
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)
}

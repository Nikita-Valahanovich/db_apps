package main

import (
	"DB_APPS/pkg/storage"
	"DB_APPS/pkg/storage/membd"
	"DB_APPS/pkg/storage/postgres"
	"fmt"
	"log"
	"os"
	"time"
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
	db, err = postgres.New(constr)
	if err != nil {
		log.Fatal(err)
	}
	db = membd.DB{}
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)

	task := postgres.Task{
		Opened:     time.Now().Unix(),
		Closed:     0,
		AuthorID:   1,
		AssignedID: 2,
		Title:      "Новая задача",
		Content:    "Описание задачи",
	}

	taskID, err := db.NewTask(task)
	if err != nil {
		log.Fatalf("Ошибка при создании задачи: %v", err)
	}
	fmt.Printf("Задача успешно создана с ID: %d\n", taskID)

	var number int
	fmt.Print("Введите id")
	id, err := fmt.Scanln(&number)
	if err != nil {
		log.Fatal(err)
	}
	task = postgres.Task{
		AuthorID:   1,
		AssignedID: 2,
		Title:      "Изменение новой задачи",
		Content:    "Изменение описания новой задачи",
	}
	taskID, err = db.UpdateTask(id, task)
	if err != nil {
		log.Fatalf("Ошибка при изменении задачи: %v", err)
	}
	fmt.Printf("Задача с ID: %d успешно изменена\n", taskID)

	taskID, err = db.DeleteTask(id)
	if err != nil {
		log.Fatalf("Ошибка при удалении задачи: %v", err)
	}
	fmt.Printf("Задача с ID: %d успешно удалена\n", taskID)
}

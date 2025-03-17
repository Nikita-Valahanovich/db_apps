package storage

import (
	"DB_APPS/pkg/storage/postgres"
)

type Interface interface {
	Tasks(int, int) ([]postgres.Task, error)
	NewTask(postgres.Task) (int, error)
	UpdateTask(id int, task postgres.Task) (int, error)
	DeleteTask(id int) (int, error)
}

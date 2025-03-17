package storage

import "DB_APPS/pkg/storage/postgres"

type Interface interface {
	Tasks(int, int) ([]postgres.Task, error)
	NewTask(postgres.Task) (int, error)
}

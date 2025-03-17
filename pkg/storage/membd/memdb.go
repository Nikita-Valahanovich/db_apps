package membd

import "DB_APPS/pkg/storage/postgres"

type DB []postgres.Task

func (db DB) Tasks(int, int) ([]postgres.Task, error) {
	return db, nil
}
func (db DB) NewTask(postgres.Task) (int, error) {
	return 0, nil
}

func (db DB) UpdateTask(id int, task postgres.Task) (int, error) {
	return id, nil
}

func (db DB) DeleteTask(id int) (int, error) {
	return id, nil
}

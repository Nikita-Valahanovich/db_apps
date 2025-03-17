package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// NewStorage создаёт новый экземпляр Storage
func NewStorage(db *pgxpool.Pool) *Storage {
	return &Storage{db: db}
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// Задача.
type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

// Tasks возвращает список задач из БД.
func (s *Storage) Tasks(taskID, authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE
			($1 = 0 OR id = $1) AND
			($2 = 0 OR author_id = $2)
		ORDER BY id;
	`,
		taskID,
		authorID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

// NewTask создаёт новую задачу и возвращает её id.
func (s *Storage) NewTask(task Task) (int, error) {
	var taskID int

	query := `
		INSERT INTO tasks (opened, closed, author_id, assigned_id, title, content)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	err := s.db.QueryRow(context.Background(), query,
		task.Opened, task.Closed, task.AuthorID, task.AssignedID, task.Title, task.Content,
	).Scan(&taskID)

	if err != nil {
		return 0, err
	}

	return taskID, nil
}

// DeleteTask удаляет задачу по id
func (s *Storage) DeleteTask(id int) (int, error) {
	_, err := s.db.Exec(context.Background(), `
		DELETE FROM tasks (id)
		WHERE id = $1;
	`,
		id,
	)

	if err != nil {
		return id, err
	}

	return id, nil
}

// UpdateTask обновляет задачу по id
func (s *Storage) UpdateTask(id int, task Task) (int, error) {
	_, err := s.db.Exec(context.Background(), `
		UPDATE tasks (author_id, assigned_id, title, content)
		SET author_id = $2, assigned_id = $3, title = $4, content = $5
		WHERE id = $1;
	`,
		id, task.AuthorID, task.AssignedID, task.Title, task.Content,
	)

	if err != nil {
		return id, err
	}

	return id, nil
}

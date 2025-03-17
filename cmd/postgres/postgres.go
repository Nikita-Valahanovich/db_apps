package postgres

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	ID   int
	Name string
}

type Book struct {
	Title string
	Year  int
}

func main() {
	// Создаём контекст
	ctx := context.Background()

	// Получаем пароль из переменных окружения
	pwd := os.Getenv("dbpass")

	// Подключаемся к базе данных через pgxpool
	db, err := pgxpool.Connect(ctx, "postgres://postgres:"+pwd+"@192.168.1.165/tasks")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	// Вставка пользователя
	_, err = db.Exec(ctx, `INSERT INTO users (name) VALUES ($1);`, "test_user")
	if err != nil {
		log.Fatalf("Insert error: %v\n", err)
	}

	// Запрос пользователей
	rows, err := db.Query(ctx, `SELECT id, name FROM users ORDER BY id;`)
	if err != nil {
		log.Fatalf("Select error: %v\n", err)
	}
	defer rows.Close()

	// Выводим пользователей
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			log.Fatalf("Scan error: %v\n", err)
		}
		log.Printf("User: ID=%d, Name=%s", u.ID, u.Name)
	}

	// Добавление книг
	books := []Book{
		{"Go Programming", 2020},
		{"Advanced Go", 2022},
	}

	if err := addBooks(ctx, db, books); err != nil {
		log.Fatalf("Failed to add books: %v\n", err)
	}
}

// Добавляет книги в БД одной транзакцией
func addBooks(ctx context.Context, db *pgxpool.Pool, books []Book) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Создаём batch
	batch := &pgx.Batch{}
	for _, book := range books {
		batch.Queue(`INSERT INTO books(title, year) VALUES ($1, $2)`, book.Title, book.Year)
	}

	// Выполняем batch-запрос
	res := tx.SendBatch(ctx, batch)
	if err := res.Close(); err != nil {
		return err
	}

	// Подтверждаем транзакцию
	return tx.Commit(ctx)
}

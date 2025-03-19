DROP TABLE IF EXISTS tasks_labels, tasks, labels, users;

-- пользователи системы
CREATE IF NOT EXISTS TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- метки задач
CREATE IF NOT EXISTS TABLE labels (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- задачи
CREATE IF NOT EXISTS TABLE tasks (
    id SERIAL PRIMARY KEY,
    opened BIGINT NOT NULL DEFAULT extract(epoch from now()), -- время создания задачи
    closed BIGINT DEFAULT 0, -- время выполнения задачи
    author_id INTEGER REFERENCES users(id) DEFAULT 0, - автор задачи
    assigned_id INTEGER REFERENCES users(id) DEFAULT 0, -- ответственный
    title TEXT, -- название задачи
    content TEXT -- задачи
    );

-- связь многие - ко- многим между задачами и метками
CREATE IF NOT EXISTS TABLE tasks_labels (
    task_id INTEGER REFERENCES tasks(id),
    label_id INTEGER REFERENCES labels(id)
);
-- наполнение БД начальными данными
INSERT INTO users (id, name) VALUES (0, 'default');
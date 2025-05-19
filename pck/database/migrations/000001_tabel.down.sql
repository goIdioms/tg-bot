DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS user_states;


-- Применить миграции:
-- migrate -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" -path migrations up

-- Откатить последнюю миграцию:
-- migrate -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" -path migrations down

-- Откатить все миграции:
-- migrate -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" -path migrations drop

-- Перейти к конкретной версии миграции:
-- migrate -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" -path migrations goto 2

-- Проверить статус миграций:
-- migrate -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" -path migrations version
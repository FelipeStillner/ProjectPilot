all:
	docker-compose down -v
	docker-compose up -d
	python3 scripts/setup-db.py

task-manager:
	cd services/task-manager && go run cmd/main.go
	
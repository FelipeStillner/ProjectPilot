all:
	docker-compose down -v
	docker-compose up -d
	python3 scripts/setup-db.py

task-manager:
	cd services/task-manager && go run cmd/main.go

text-interpreter:
	cd services/text-interpreter && pip install -r requirements.txt > /dev/null 2>&1 && python3 src/main.py
	
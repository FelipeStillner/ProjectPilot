all:
	docker-compose down -v
	docker-compose up -d
	python3 scripts/setup-db.py

task-manager:
	cp .env services/task-manager/.env
	python3 -m grpc_tools.protoc -I. --python_out=services/text-interpreter --pyi_out=services/text-interpreter --grpc_python_out=services/text-interpreter proto/task-manager.proto
	protoc --go_out=services/task-manager --go-grpc_out=services/task-manager proto/task-manager.proto
	cd services/task-manager && go run cmd/main.go

text-interpreter:
	cp .env services/text-interpreter/.env
	python3 -m grpc_tools.protoc -I. --python_out=services/text-interpreter --pyi_out=services/text-interpreter --grpc_python_out=services/text-interpreter proto/text-interpreter.proto
	cd services/text-interpreter && pip install -r requirements.txt > /dev/null 2>&1 && python3 main.py

access-manager:
	cp .env services/access-manager/.env
	protoc --go_out=services/access-manager --go-grpc_out=services/access-manager proto/access-manager.proto
	cd services/access-manager && go run cmd/main.go

calendar-manager:
	cp .env services/calendar-manager/.env
	protoc --go_out=services/calendar-manager --go-grpc_out=services/calendar-manager proto/calendar-manager.proto
	cd services/calendar-manager && go run cmd/main.go
	
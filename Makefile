all:
	docker-compose down -v
	docker-compose up -d
	python3 scripts/setup-db.py

task-manager:
	python3 -m grpc_tools.protoc -I. --python_out=services/text-interpreter --pyi_out=services/text-interpreter --grpc_python_out=services/text-interpreter proto/task-manager.proto
	protoc --go_out=services/task-manager --go-grpc_out=services/task-manager proto/task-manager.proto
	cd services/task-manager && go run cmd/main.go

text-interpreter:
	python3 -m grpc_tools.protoc -I. --python_out=services/text-interpreter --pyi_out=services/text-interpreter --grpc_python_out=services/text-interpreter proto/text-interpreter.proto
	cd services/text-interpreter && pip install -r requirements.txt > /dev/null 2>&1 && python3 main.py
	
import grpc
from concurrent import futures
import proto.task_manager_pb2
import proto.task_manager_pb2_grpc
import proto.text_interpreter_pb2
import proto.text_interpreter_pb2_grpc
from openai import OpenAI
import os
from dotenv import load_dotenv
import json


# gRPC Client for TaskManager
class TaskManagerClient:
    def __init__(self):
        host = os.getenv("HOST_TASK_MANAGER")
        port = os.getenv("PORT_GRPC_TASK_MANAGER")
        self.channel = grpc.insecure_channel(host + ":" + port)
        self.stub = proto.task_manager_pb2_grpc.TaskManagerStub(self.channel)

    def create_task(self, name, description, priority, assignee, status):
        request = proto.task_manager_pb2.CreateTaskRequest(
            name=name,
            description=description,
            priority=priority,
            assignee=assignee,
            status=status,
        )
        return self.stub.CreateTask(request)


# gRPC Server Implementation
class TextInterpreterServicer(proto.text_interpreter_pb2_grpc.TextInterpreterServicer):
    def __init__(self):
        self.task_manager_client = TaskManagerClient()

    def InterpretText(self, request, context):
        text = request.text

        # Call ChatGPT API to extract tasks
        tasks = self.extract_tasks_from_text(text)

        # Process extracted tasks
        for task in tasks:
            self.task_manager_client.create_task(
                name=task["name"],
                description=task["description"],
                priority=task["priority"],
                assignee=task["assignee"],
                status=task["status"],
            )

        return proto.text_interpreter_pb2.InterpretTextResponse(
            success=True, message="Tasks processed successfully"
        )

    def extract_tasks_from_text(self, text):
        client = OpenAI(api_key=os.getenv("OPENAI_API_KEY"))

        completion = client.chat.completions.create(
            model="gpt-4o-mini",
            store=True,
            messages=[
                {
                    "role": "user",
                    "content": """Parse the next message text and extract tasks following the exact format for your response:
                    [
            {
                "name": "Task 1",
                "description": "Task 1 description",
                "priority": "HIGH",
                "assignee": 1,
                "status": "TODO",
            },
        ]""",
                },
                {
                    "role": "user",
                    "content": text,
                }
            ],
        )

        tasksJsonString = completion.choices[0].message.content
        tasks = json.loads(tasksJsonString)
        return tasks


# Start gRPC server
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    proto.text_interpreter_pb2_grpc.add_TextInterpreterServicer_to_server(
        TextInterpreterServicer(), server
    )
    host = os.getenv("HOST_TEXT_INTERPRETER")
    port = os.getenv("PORT_GRPC_TEXT_INTERPRETER")
    server.add_insecure_port("[::]:" + port)
    server.start()
    print("TextInterpreter gRPC Server started on port " + port)
    server.wait_for_termination()


if __name__ == "__main__":
    load_dotenv()
    serve()

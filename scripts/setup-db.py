from time import sleep
import psycopg2
import os
from dotenv import load_dotenv

load_dotenv()

# Database connection configuration
HOST_POSTGRES = os.getenv('HOST_POSTGRES')
PORT_POSTGRES = os.getenv('PORT_POSTGRES')
USER_POSTGRES = os.getenv('USER_POSTGRES')
PASSWORD_POSTGRES = os.getenv('PASSWORD_POSTGRES')
DATABASE_POSTGRES = os.getenv('DATABASE_POSTGRES')

# Connect to the PostgreSQL server
try:
    sleep(1)
    connection = psycopg2.connect(
        host=HOST_POSTGRES,
        port=PORT_POSTGRES,
        user=USER_POSTGRES,
        password=USER_POSTGRES,
        dbname=DATABASE_POSTGRES,
    )
    cursor = connection.cursor()

    # SQL commands to create tables and set up the database
    create_table_sql = """
    CREATE TABLE IF NOT EXISTS tasks (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT,
        priority TEXT NOT NULL,
        Assignee TEXT,
        status TEXT NOT NULL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    );
    """

    # Execute the SQL commands
    cursor.execute(create_table_sql)

    # Commit changes
    connection.commit()

    print("Database setup complete!")

    # Close communication with the database
    if connection:
        cursor.close()
        connection.close()
except Exception as error:
    print(f"Error while setting up the database: {error}")
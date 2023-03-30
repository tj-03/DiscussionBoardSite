## Setting Up Backend Server Project

### Prerequisites
- Go installed on your machine
- MongoDB installed on your machine
- mongod program is in your PATH 
### Starting the Database 
1. Once you have mongo installed and have the mongod binary in your path or have a mongod process running on your machine, run the database initialization program to set the databases up.
To do this, cd to the DB_Init folder and run
    ```
    go run .
    ```

2. In the backend folder run 
    ``` 
    go run .
    ``` 
    to start a regular server and run 
    ``` 
    go run . -local
    ``` 
    to start a mock server if you want to run the server with no authentication and static mock data.

    If you are running the server with authenication enabled, make sure you the firebaseKeys.json file in the backend directory. 


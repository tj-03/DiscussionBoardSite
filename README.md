# ReadMe

## Semester Project for Intro to Software Engineering at the University of Florida

### Project Name 
 - Discussion Board Website

### Project Description
 - Website where users can discuss various topics, marketed towards software engineers and similar disciplines.

### Project Model

#### Front-end
- Programmed in Angular Typescript.

#### Back-end
- Programmed in Golang and implementing Firebase for various applications of databases.

### Contributors

#### Front-end
- Alexander Opyrchal
- Ilya Medvedev

#### Back-end
- Tristan Joseph
- Anthony Rebello

### Setup and Initialization

#### Required Downloads
- Go installed on machine
- MongoDB installed on machine
- Mongo program is in PATH
- Angular installed on machine
- Cypress installed on machine (for testing)

#### Initializing the Database
- Start MongoDB service (on mac the command is 'brew services start mongodb-community<version>')
- Initialize the database by navigating to DB_Init folder and running the command 'go run .'

#### Starting the Server
- Navigate to the backend folder and run the command 'go run .' for test, and 'go run . -local' for regular server. Test does not require authentication.
- The server, by default, listens on localhost:8080

#### Starting Program
- To ensure Angular is properly installed, run 'ng add @angular/material'
- Run 'ng serve' in the 'frontend' folder to start the program
- Connect by opening a browser of your choice and navigating to the URL provided by the command prompt

#### Cypress Testing
- Run 'npx cypress open' in a different command prompt instance in the same folder to open the testing program
- Select 'E2E Testing'. All tests should be available here.
- Select any test to begin testing

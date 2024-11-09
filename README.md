![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white) ![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E) ![CSS3](https://img.shields.io/badge/css3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)

# Medical Portal

A full-stack Golang web application with a PostgreSQL database for managing patient records in a medical setting. The application features two main portals: a Receptionist Portal and a Doctor Portal.

## Features

- User Authentication with JWT
- Role-Based Access Control (RBAC)
- Receptionist Portal:
  - Register new patients
  - Perform CRUD operations on patient records
- Doctor Portal:
  - View patient details
  - Update patient information

## Tech Stack

- Backend: Golang with net/http
- Database: PostgreSQL
- Frontend: HTML, CSS (Tailwind CSS), and JavaScript

## Project Structure

```
medical-portal/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── auth/
│   │   ├── auth.go
│   │   └── auth_test.go
│   ├── handlers/
│   │   ├── doctor.go
│   │   ├── doctor_test.go
│   │   ├── receptionist.go
│   │   └── receptionist_test.go
│   ├── middleware/
│   │   ├── rbac.go
│   │   └── rbac_test.go
│   ├── models/
│   │   ├── patient.go
│   │   └── user.go
│   └── repository/
│       ├── patient.go
│       ├── patient_test.go
│       ├── user.go
│       └── user_test.go
├── migrations/
│   ├── 000001_create_users_table.up.sql
│   ├── 000001_create_users_table.down.sql
│   ├── 000002_create_patients_table.up.sql
│   └── 000002_create_patients_table.down.sql
├── static/
│   ├── index.html
│   ├── login.html
│   ├── doctor.html
│   └── receptionist.html
├── .env
├── go.mod
├── go.sum
└── README.md
```


## Setup and Installation

1. Clone the repository
2. Install Go (version 1.17 or later) and PostgreSQL
3. Set up the database and update the `.env` file with your PostgreSQL credentials
4. Run `go mod tidy` to install dependencies
5. Apply database migrations
6. Build and run the server: `go run cmd/server/main.go`

## Usage

1. Open a web browser and navigate to `http://localhost:8080`
2. Log in as either a doctor or receptionist
3. Use the respective portal to manage patient records

## API Documentation

A Postman collection is provided in the repository for testing the API endpoints.
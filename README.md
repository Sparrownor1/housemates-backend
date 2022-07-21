# housemates-backend
an all-round app to handle everything house-related

## Setup
Requirements:
- Go
- PostgreSQL

Database Setup:
- Make sure you have a database created for Housemates
- In a `.env` file, initialize the following keys:
  - `DB_USER`
  - `DB_PASSWORD`
  - `DB_NAME`
  - `PORT` (optional)
 
Running the API:
- `go run main.go`
- The API should be available at `localhost:5000` or at the port specified

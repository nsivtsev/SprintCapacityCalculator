Let's start with the backend part of the project. We will use GoLang for the backend and MySQL for the database. We will use the `gin-gonic/gin` package for creating the REST API and `go-sql-driver/mysql` for connecting to the MySQL database.

backend/main.go
# Install dependencies
cd backend
go mod init
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/mysql

cd ../frontend
npm install

# Run all necessary parts of the codebase
# Assuming Docker and Docker Compose are already installed
docker-compose up --build
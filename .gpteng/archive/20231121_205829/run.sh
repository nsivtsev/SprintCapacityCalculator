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

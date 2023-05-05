# Go programming language - Gin Web Framework

## Packages

- gin [https://pkg.go.dev/github.com/gin-gonic/gin](https://pkg.go.dev/github.com/gin-gonic/gin)
- godotenv [https://pkg.go.dev/github.com/joho/godotenv](https://pkg.go.dev/github.com/joho/godotenv)
- gorm [https://pkg.go.dev/gorm.io/gorm](https://pkg.go.dev/gorm.io/gorm)
- gorm PostgreSQL driver [https://pkg.go.dev/gorm.io/driver/postgres](https://pkg.go.dev/gorm.io/driver/postgres)

``` bash
# Install gin package
go get github.com/gin-gonic/gin

# Install godotenv package
go get github.com/joho/godotenv

# Install gorm package
go get gorm.io/gorm

# Install gorm PostgreSQL driver package
go get gorm.io/driver/postgres
```

## Start server and application

``` bash
docker-compose up -d --build
```

## Test

after start server

Open [http://localhost:3000](http://localhost:3000) with your browser to test service and application.

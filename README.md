# JWT authentication in Go

Install packages
```sh
# for web server and routing - Gin
go get -u github.com/gin-gonic/gin

# for loading environment variables
go get -u github.com/joho/godotenv

# for DB
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

# for hashing the password
go get -u golang.org/x/crypto/bcrypt

# for JWT
go get -u github.com/golang-jwt/jwt/v5

# TODO: for hot reloading
# go get github.com/githubnemo/CompileDaemon
# go install github.com/githubnemo/CompileDaemon
```

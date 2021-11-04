# Simple CRUD API Golang Postgres

CRUD API Golang with postgresql
ERD : 
![Screenshot](ERD.png)

## how-to
- download this repo
- prepare the database, just run migration script in steps below
- Change .env and user.env as your setting
- run the migration script to create database
```
go run gorm-api\script\migration\create_migration_script.go
```
- run the program
```
go run gorm-api\app\main.go
```
register user
```
curl -X POST -F "username=user" -F "password=test" -F "admin=true" http://localhost:8080/create-user
```
login ( this will generate access token for 45 minutes)
```
curl -X POST localhost:8080/login -d "username=user&password=test"
```
test admin (with token example)
```
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTQyODAyMzk5LCJuYW1lIjoiSm9uIERvZSJ9.K-uDZN_XX0J9PaUJXeRGjHmtfDwLr9StGZG1FOIa5Hc" localhost:8080/admin
```

# Golang - Books CRUD API

A simple REST API to CRUD books on a mysql database.

## Tecnologies
- Golang Mux Framework
- Golang
- Mysql driver

## Launch
- `$ git clone git@github.com:arpesam/golang-books-api.git`
- `go run main.go`

>:warning: you need to run a docker or local MySQL client at localhost:3306. The mux framework will automatically create a table on your MySql database.

To run a docker with mysql, just run `docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7`

## Endpoints
#### Create
```
curl -X POST http://localhost:8080/books/ \
 -H 'Content-Type: application/json' \
 -d '{ "name": "Happy Book Store", "author": "Elon Musk", "publication": "Nature" }'
 ```
#### Read
```
curl http://localhost:8080/books/{bookId} \
 -H 'Content-Type: application/json'
 ```
#### Update
```
curl -X PUT http://localhost:8080/books/{bookId} \
 -H 'Content-Type: application/json' \
 -d '{ "name": "Happy Book Store", "author": "Elon Musk", "publication": "Nature" }'
 ```
#### Delete
```
curl -X DELETE http://localhost:8080/books/{bookId}
 ```
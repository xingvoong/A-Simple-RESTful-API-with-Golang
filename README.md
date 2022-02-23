# A-Simple-RESTful-API-with-Golang
a simple RESTful API with Golang with Gin web framework that add and retrieve books


## Author
- [@xingvoong](https://github.com/xingvoong)

## API docs

### Get all books
* GET `books`
* Retrieves a list of all books

**Usage**
```
curl http://localhost:8080/books
```

**Returns:** JSON
```
[
    {
        "id": "1",
        "name": "How to Be an Antiracist",
        "author": "Ibram X.Kendi",
        "genre": "Nonfiction"
    },
    {
        "id": "2",
        "name": "Clean Code: A Handbook of Agile Software Craftmanship",
        "author": "Robert C.Martin Series",
        "genre": "Computer Science"
    },
    {
        "id": "3",
        "name": "Geek Love",
        "author": "Katherine Dunn",
        "genre": "Fiction"
    }
]

```

**Success Status Code:** `200 OK`

### Get a book by name
* GET `books/book_name`
* Retrieves a book with the given name

**Query Parameters**:
| Parameter   | Type        |Description|
| ----------- | ----------- |-----------|
| book_name   | string      |Specifies the name for which to retrieve the book.|

**Success Status Code:** `200 OK`

**Usage**:
```
curl http://localhost:8080/books/"Geek%20Love"

```
**Returns**: JSON
```
{
    "id": "3",
    "name": "Geek Love",
    "author": "Katherine Dunn",
    "genre": "Fiction"
}
```

### Add a book
* POST `books`
* Add a book with given ID, title, author and genre

**Success Status Code:** `201 OK`

**Usage**:
```
curl http://localhost:8080/books
  --include \
  --header "Content-Type: application/json" \
  --request "POST" \
  --data '{"id": "4","name": "Ultralearning","author": "Scott H.Young","genre": "Productivity"}'

```

**Return**:
```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Wed, 23 Feb 2022 00:17:45 GMT
Content-Length: 106

{
    "id": "4",
    "name": "Ultralearning",
    "author": "Scott H.Young",
    "genre": "Productivity"
}
```
## Run Locally
Check required tech is installed (see below).

Clone the project
```bash
https://github.com/xingvoong/A-Simple-RESTful-API-with-Golang
```
Go to the project directory
```bash
cd A-Simple-RESTful-API-with-Golang
```

Start the server
```bash
go run .
```
Now the app is ready at http://localhost:8080

## Tech Stack
**Server**: Golang, Gin

## Requirements
- Web browser
- Golang, gin
- Linux, macOS, or Windows

## Acknowledgement
- [go.dev](https://go.dev/doc/tutorial/web-service-gin)

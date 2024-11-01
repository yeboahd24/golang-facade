# Bookstore API

This is a simple API that allows you to create, read, update, and delete books.


This restructured version provides several benefits:

- It uses a Facade pattern to separate the business logic from the HTTP routes.
- It uses a Repository pattern to separate the data access logic from the business logic.
- It uses a Service pattern to separate the business logic from the data access layer.
- It uses a Dependency Injection pattern to separate the data access layer from the business logic.

## Endpoints

### Create a book

`POST /books`

Example request body:

```json
{
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "price": 12.99
  }

  ```

Example response:

```json
{
  "id": 1,
  "title": "The Great Gatsby",
  "author": "F.
  Scott Fitzgerald",
  "price": 12.99
}
```

### Get all books

`GET /books`

Example response:

```json
[
  {
    "id": 1,
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "price": 12.99
  },
  {
    "id": 2,
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "price": 9.99
  }
]
```

### Get a book

`GET /books/:id`

Example response:

```json
{
  "id": 1,
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "price": 12.99
}
```

### Update a book

`PUT /books/:id`

Example request body:

```json
{
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "price": 12.99
}
```

Example response:

```json
{
  "id": 1,
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "price": 12.99
}
```

### Delete a book

`DELETE /books/:id`

Example response:

```json
{
  "message": "Book deleted successfully"
}
```

## Running the application

To run the application, you need to have Go installed on your machine. You can then run the following command:

```bash
go run cmd/main.go
```

This will start the application on port 8080.

## Testing the application

To test the application, you can use the following command:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title": "The Great Gatsby", "author": "F. Scott Fitzgerald", "price": 12.99}' http://localhost:8080/books
```

This will create a new book with the given title, author, and price.

You can then use the following command to get all books:

```bash
curl http://localhost:8080/books
```

This will return all books in the database.

You can then use the following command to get a specific book:

```bash
curl http://localhost:8080/books/1
```

This will return the book with the ID 1.

You can then use the following command to update a book:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"title": "The Great Gatsby", "author": "F. Scott Fitzgerald", "price": 12.99}' http://localhost:8080/books/1
```

This will update the book with the ID 1 with the given title, author, and price.

You can then use the following command to delete a book:

```bash
curl -X DELETE http://localhost:8080/books/1
```

This will delete the book with the ID 1.

## Conclusion

In this tutorial, we learned how to create a simple API using Go and the Gin framework. We also learned how to use the Facade pattern, Repository pattern, Service pattern, and Dependency Injection pattern to separate the business logic from the HTTP routes.

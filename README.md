# Go Gram

Instagram Clone RESTful API with Go

Final project for Scalable Web Service with Golang Course - DTS Kominfo X hacktiv8id

- [Demo](https://go-gram-production.up.railway.app)
- [API Documentation](https://go-gram-production.up.railway.app/swagger/index.html)

## Instalation & Run Locally

```bash
# Download this project
go get github.com/husfuu/go-gram
```

Before running the API server, you must set up the database configuration with yours by creating an `.env` file. You can see the environment variables used in this project in the `.env.example` file in this repository.

Then follow this steps:

```bash
# Build and Run
cd go-gram
go build
./go-gram
```

## ERD
![drawSQL-export-2022-10-25_08_27](https://user-images.githubusercontent.com/70875733/197654868-fb9b6279-944e-45cb-aa61-49d7b198e2a6.png)

## Structure

```js
go-gram
├── auth
├── config // database connection and migration
├── dto // request and response schema and put some validation
│   ├── comment-dto.go
│   ├── photo-dto.go
│   └── ...
├── entity // schemas of database table
│   ├── comment
│   ├── photo
│   └── ...
├── handler // payload process from frontend
│   ├── commentHandler
│   │   └── comment-handler.go
│   ├── photoHandler
│   │   └── photo-handler.go
│   └── ...
├── helper
│   ├── constant.go
│   └── ...
├── middleware
│   └── authorization.go
├── repository // query process to interact to database
│   ├── commentRepository
│   │   └── comment-repository.go
│   └── ...
├── server
│   ├── route.go
│   └── server.go
├── service // implement business logic
│   ├── commentService
│   │   └── comment-service.go
│   └── ...
└── validation
    ├── comment-validation.go
    ├── photo-validation.go
    └── ...
```

## TODO

- [ ] Database seeding
- [ ] Http Test
- [x] API Doc
- [ ] Implement soft delete

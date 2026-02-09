# Todo API (Go)

A simple backend application written in Go that accepts HTTP requests and performs basic CRUD operations for Todo items using in-memory storage.

## Features
- GET /todos        -> list all todos
- POST /todos       -> create a new todo
- DELETE /todos/{id} -> delete a todo by id

## Tech
- Go (net/http)
- JSON encode/decode
- In-memory storage (slice)

## Run
```bash
go run .
```
**Or**

```bash
go run main.go handlers.go models.go
```

## Server runs at

http://localhost:8080/todos


## API Examples (PowerShell)

### Get Todo

```bash
Invoke-RestMethod http://localhost:8080/todos
```

<img width="498" height="309" alt="image" src="https://github.com/user-attachments/assets/1c22cdc9-293c-4e99-a371-14ba7f365b12" />


### Create Todo

```bash
Invoke-RestMethod -Uri http://localhost:8080/todos -Method POST -ContentType "application/json" -Body '{"title":"Learn Go","done":false}'
```


<img width="921" height="682" alt="image" src="https://github.com/user-attachments/assets/d83932a6-61b6-429e-ae38-a4dce9ae11b2" />

<img width="907" height="253" alt="image" src="https://github.com/user-attachments/assets/f61bd01f-a946-48f8-aa19-0d4ead211b7c" />

<img width="764" height="192" alt="image" src="https://github.com/user-attachments/assets/14d6384a-0fbf-46cd-963f-79041b5de97f" />

### Delete Todo (example id = 1)

```bash
Invoke-RestMethod -Uri http://localhost:8080/todos/1 -Method DELETE
```

<img width="889" height="274" alt="image" src="https://github.com/user-attachments/assets/6beead28-cfc2-48dc-8fe1-3b35b27e9979" />

<img width="852" height="267" alt="image" src="https://github.com/user-attachments/assets/77ace26e-254e-4667-81a8-5477bb1afd29" />




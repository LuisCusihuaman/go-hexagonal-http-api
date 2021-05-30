<!-- TITLE -->
<h1 align="center">Go Hexagonal RESTful API</h1>

<!-- SUMMARY -->

_Hexagonal architecture, cqrs, event modeling and testing in golang_

Based on the projects of the [API HTTP en Go aplicando Arquitectura Hexagonal
](https://pro.codely.tv/library/api-go-hexagonal/199314/about/)
<br />

## 🔨 Requirements

- [Go v1.15+](https://golang.org/dl/)
- [Docker](https://docs.docker.com/engine/install/#server)
  and [Docker Compose](https://docs.docker.com/compose/install/#install-compose)

## 🚀 Environment setup

```bash
# download project
git clone https://github.com/LuisCusihuaman/go-hexagonal-http-api

cd go-hexagonal-http-api

# run api with docker compose
docker-compose up -d --build
```

## 🏎️ Getting Started

At this time, you have a RESTful API server running at `http://localhost:8080/`. It provides the following endpoints:

* `GET /health`: a healthcheck service provided for health checking purpose (needed when implementing a server cluster)
* `POST /courses`: creates a new course

Try the URL `http://localhost:8080/health` in a browser, and you should see something like `"everything is ok!"`
displayed.

If you have `cURL` or some API client tools (e.g. [Postman](https://www.postman.com/downloads/)), you may try the
following more complex scenarios:

```bash
# healthcheck service provided by: GET /health
curl -X GET --location "http://localhost:8080/health" \
    -H "Accept: application/json"
# should return a healthcheck text: 'everything is ok!'

# creates a new course via: POST /courses
curl -sw '%{http_code}' -X POST --location "http://localhost:8080/courses" \
    -H "Content-Type: application/json" \
    -d "{
          \"id\": \"11d365ab-062f-4dff-9277-8474c50d8e12\",
          \"name\": \"Golang course\",
          \"duration\": \"1 month\"
        }"
# should return a http status code: 201
```

### Connect to database

```bash
docker-compose exec mysql mysql -u user -ppassword -D backoffice

SELECT * FROM courses;
```

### Show logs

```bash
docker-compose logs -f mooc-api
```

### Tests

To execute all tests, just run:

```bash
go test ./... 
```
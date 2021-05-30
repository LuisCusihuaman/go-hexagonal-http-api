FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/LuisCusihuaman/go-hexagonal_http_api-course
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/luiscusihuaman-mooc-api cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/luiscusihuaman-mooc-api /go/bin/luiscusihuaman-mooc-api
ENTRYPOINT ["/go/bin/luiscusihuaman-mooc-api"]

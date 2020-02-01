FROM golang:1.13.5-alpine3.11 AS build

WORKDIR /
COPY . /go/src/github.com/easy-to-study/go_study/practice/toDoListBackend
RUN apk update \
  && apk add --no-cache git \
  && go get github.com/go-sql-driver/mysql \
  && go get github.com/google/uuid \
  && go get github.com/gorilla/mux
RUN cd /go/src/github.com/easy-to-study/go_study/practice/toDoListBackend && go build -o bin/todolist main.go

FROM alpine:3.8
COPY --from=build /go/src/github.com/easy-to-study/go_study/practice/toDoListBackend/bin/todolist /usr/local/bin/
CMD ["todolist"]

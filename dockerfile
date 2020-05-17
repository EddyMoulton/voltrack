FROM golang:latest

WORKDIR /app
ADD VERSION .

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build ./cmd/stocktracker" --command="./stocktracker"
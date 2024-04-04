FROM golang:alpine
WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8083

ENTRYPOINT [ "go", "run", "main.go" ]
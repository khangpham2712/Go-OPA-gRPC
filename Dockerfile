FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod tidy
# COPY *.go .
COPY . .
EXPOSE 50051
ENTRYPOINT [ "go", "run", "cmd/main.go" ]

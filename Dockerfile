FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out ./cmd/main/main.go
EXPOSE 9090:9090
CMD ./out
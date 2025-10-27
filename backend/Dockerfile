FROM golang:latest

WORKDIR /app/server

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o server .

EXPOSE 8080

ENV PORT 8080

CMD ["/app/server/server"]
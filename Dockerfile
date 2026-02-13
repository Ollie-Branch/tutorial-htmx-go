FROM golang:latest AS base

# Set working directory
WORKDIR /app

COPY . .

RUN go mod download && go build -o htmx-backend

EXPOSE 9001

CMD ["/app/htmx-backend"]


FROM golang:1.17-alpine3.14 as builder

WORKDIR /app/backend

RUN apk update && apk upgrade

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build

FROM alpine:latest

WORKDIR /app/backend/

COPY --from=builder /app/backend/backend .
COPY --from=builder /app/backend/.env .

CMD ["./backend"]

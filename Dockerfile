FROM golang:1.21.6 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -trimpath -o app .

FROM alpine:latest

RUN adduser -D appuser

RUN apk --no-cache add ca-certificates

WORKDIR /home/appuser

COPY --from=build /app/app .

USER appuser

CMD ["./app"]
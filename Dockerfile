FROM golang:1.12-alpine as builder

RUN apk add --no-cache git

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ./out/fizzbuzzserver .

FROM scratch

COPY --from=builder /build/out/fizzbuzzserver /app/fizzbuzzserver

EXPOSE 8080

CMD ["./app/fizzbuzzserver"]
# My dockerfile
FROM golang as builder

ENV GO111MODULE=on
WORKDIR /app

ADD . .

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-assignment .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/go-assignment .

CMD ["./go-assignment"]

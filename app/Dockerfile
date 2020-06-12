FROM golang:1.14.4 as builder
WORKDIR /app/
RUN go get "github.com/sirupsen/logrus"
COPY main.go . 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.12.0
WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main .
CMD ["./main"]
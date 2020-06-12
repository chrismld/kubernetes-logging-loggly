FROM golang:1.14.4
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get "github.com/sirupsen/logrus"
RUN go build -o main . 

FROM alpine:3.12.0
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./main"]
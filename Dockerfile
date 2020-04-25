FROM golang:1.14.2-alpine

RUN mkdir /app 

ADD . /app/

WORKDIR /app 

RUN go build -o main .

ENV PORT 9000

CMD ["./main"]
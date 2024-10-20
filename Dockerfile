FROM golang:1.22.4

WORKDIR /app

COPY ./ /app

RUN go build main.go

CMD /app/main

EXPOSE 8080
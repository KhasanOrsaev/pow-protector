FROM golang:1.17

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o /usr/bin/pow-server main.go

CMD ["pow-server", "8001"]

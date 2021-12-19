FROM golang:1.17

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o /usr/bin/pow-client main.go

CMD ["sleep","3600"]

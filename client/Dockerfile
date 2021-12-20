FROM golang:1.17

RUN mkdir /app
COPY . /app/
WORKDIR /app
RUN go build -o pow-client main.go

FROM ubuntu

COPY --from=0 /app/pow-client /usr/bin

CMD ["sleep","3600"]

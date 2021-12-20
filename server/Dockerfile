FROM golang:1.17

RUN mkdir /app
COPY . /app/
WORKDIR /app
RUN go build -o pow-server main.go

FROM ubuntu

COPY --from=0 /app/pow-server /usr/bin

CMD ["pow-server", "8001"]

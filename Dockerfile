
FROM  golang:1.19rc2-alpine3.15


WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY main.go ./

RUN go build main.go

EXPOSE 8080

CMD [ "/app/main" ]
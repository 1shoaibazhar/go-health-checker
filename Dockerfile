FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 4000

RUN go build -o app

ENTRYPOINT [ "./app" ]

FROM golang:1.15-alpine

WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download

COPY . /app

RUN go build -o party-app

EXPOSE 9000

CMD [ "/app/party-app" ]
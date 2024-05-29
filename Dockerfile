FROM golang:1.19-alpine

WORKDIR /usr/scr/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./cmd/web/*.go

EXPOSE 8080
ENTRYPOINT ["app"] 
FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /notesyapi

EXPOSE 8080

CMD [ "/notesyapi" ]
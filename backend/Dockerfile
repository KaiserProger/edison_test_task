FROM golang:1.20-alpine

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /application

EXPOSE 8080

CMD [ "/application" ]
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

ARG CORS_IP
ENV CORS_IP $CORS_IP

RUN CGO_ENABLED=0 GOOS=linux go build -o /application

EXPOSE 8080

CMD [ "/application" ]
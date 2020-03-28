FROM golang:1.13.0-alpine3.10

WORKDIR /app

COPY . .
RUN go mod download

# build app binary
RUN go build

EXPOSE 8080
CMD [ "./social-dashboard" ]


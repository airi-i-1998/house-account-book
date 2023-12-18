FROM golang:1.20-alpine3.16
ENV ROOT /app
WORKDIR ${ROOT}
RUN apk update && apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY source dest
EXPOSE 8080
CMD ["go", "run", "main.go"]

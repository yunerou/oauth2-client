FROM golang:1.20-bullseye AS base
RUN apt-get update -y && apt-get install -y openssl git curl
WORKDIR /app
RUN go env -w GOPROXY=https://goproxy.io,https://proxy.golang.org,direct
COPY . /app
RUN go get -d
RUN CGO_ENABLED=0 go build . && rm -rf /app/*


# Dev enviroment
FROM base as dev
COPY . /app
RUN go get -u
RUN go mod tidy
CMD ["go", "run", "."]
EXPOSE 8080

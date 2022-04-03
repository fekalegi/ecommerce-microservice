FROM golang:alpine AS build-env

WORKDIR /product

COPY go.mod ./
COPY go.sum ./

COPY . .

RUN apk add --no-cache bash ca-certificates git
ENV GOPRIVATE=github.com/fekalegi/*
ENV GONOSUMDB=github.com/fekalegi/*
RUN echo "machine github.com login fekalegi password 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" > ~/.netrc
RUN go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o goapp .

COPY *.go ./

RUN go build -o /ecommerce-microservice/product

EXPOSE 5001

CMD ["/ecommerce-microservice/product"]

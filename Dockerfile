FROM golang:1.12-alpine AS build_base
RUN apk add --no-cache git
WORKDIR /tmp/go-sample-app
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
#RUN go build -o /go/bin/myapp main.go
RUN go build -o ./out/go-sample-app .
# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates
COPY --from=build_base /tmp/go-sample-app/out/go-sample-app /app/go-sample-app
EXPOSE 4000
CMD ["/app/go-sample-app"]
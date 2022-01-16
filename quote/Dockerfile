FROM golang:1.16.5-alpine as builder

EXPOSE 8080
RUN apk update && apk add git

WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
COPY ./quote .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app .

FROM scratch
WORKDIR /bin
COPY --from=builder /go/bin/app /app

ENTRYPOINT ["/app"]
FROM golang:1.18-alpine AS builder

RUN apk --no-cache add git

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOSUMDB=off

WORKDIR /app

COPY . .

RUN go mod vendor
RUN go build -ldflags "-s -w" -o main

FROM gcr.io/distroless/static

COPY --from=builder /app/main /app/main

WORKDIR /app

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["./main"]

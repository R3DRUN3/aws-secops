#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./...

#final stage
FROM alpine:latest
#add non-root user
RUN addgroup -S dockerappgroup && adduser -S aws-secops-user -G dockerappgroup 
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
USER aws-secops-user
ENTRYPOINT /app
LABEL Name=awssecops Version=0.0.1

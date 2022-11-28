FROM golang AS builder
WORKDIR /app
ADD . .
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -o /app/app
ENTRYPOINT /app/app

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/app /app
ENTRYPOINT [ "/app/app" ]

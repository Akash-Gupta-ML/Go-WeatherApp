FROM golang:1.20 AS project

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o server .

FROM alpine:latest

RUN apk update && apk add --no-cache libc6-compat

RUN mkdir -p /app/templates /app/static

COPY --from=project /app/server /server

COPY --from=project /app/templates /app/templates
COPY --from=project /app/static /app/static

EXPOSE 8080

CMD ["/server"]

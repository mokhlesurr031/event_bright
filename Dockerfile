FROM golang:1.19-alpine3.16 AS build
WORKDIR /app
COPY . .
RUN go build -o main .


FROM alpine:3.14
WORKDIR /app
COPY --from=build /app .
ENTRYPOINT ["/app/main"]
CMD ["serve"]
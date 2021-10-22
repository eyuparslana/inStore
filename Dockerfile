# syntax=docker/dockerfile:1
FROM golang:alpine

ENV GO111MODULE=on
ENV API_PORT=8080
ENV EXPORT_FILE_PATH=/app/output
WORKDIR /app
COPY . ./
RUN go build -o /in-store

CMD [ "/in-store" ]
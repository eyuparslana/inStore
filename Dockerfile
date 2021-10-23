# syntax=docker/dockerfile:1
FROM golang:alpine

ENV GO111MODULE=on
ENV API_PORT=8080
ENV EXPORT_FILE_PATH=/./output
ENV RECORD_FREQ=10
WORKDIR /app
COPY . ./
RUN go build -o /in-store

CMD [ "/in-store" ]
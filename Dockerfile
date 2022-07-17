FROM golang:1.17-alpine AS GO_BUILD
RUN apk update && apk add --no-cache musl-dev gcc build-base
COPY . /bookmicroservice
COPY ./config/*.yaml /go/bin/bookmicroservice/config/
WORKDIR /bookmicroservice
RUN go build -o /go/bin/bookmicroservice

FROM alpine:latest
# RUN apk --no-cache add curl
WORKDIR /app
COPY --from=GO_BUILD /go/bin/bookmicroservice/ ./
EXPOSE 8080
CMD ./bookmicroservice
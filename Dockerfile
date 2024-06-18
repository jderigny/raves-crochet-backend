FROM golang:1.22 AS build-stage

WORKDIR /go/src/github.com/jderigny/raves-crochet-backend

COPY . .

WORKDIR /go/src/github.com/jderigny/raves-crochet-backend/cmd/raves-crochet-backend

RUN CGO_ENABLED=0 GOOS=linux go build -o /raves-crochet-backend

FROM debian:12.5-slim AS build-release-stage

WORKDIR /

COPY --from=build-stage /raves-crochet-backend /raves-crochet-backend

EXPOSE 80

ENTRYPOINT ["/raves-crochet-backend"]
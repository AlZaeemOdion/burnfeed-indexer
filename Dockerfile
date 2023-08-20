FROM golang:1.21-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git make

WORKDIR /indexer
COPY . .
RUN make build

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /indexer/bin/indexer /usr/local/bin/

ENTRYPOINT ["indexer"]

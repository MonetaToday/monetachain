FROM --platform=linux golang:1.21-alpine AS builder

RUN apk add --update --no-cache make git

WORKDIR /root
COPY . ./monetachain

WORKDIR /root/monetachain
RUN go build -o ./build/monetachaind ./cmd/monetachaind/main.go

FROM --platform=linux alpine

RUN apk add --update --no-cache curl jq

COPY --from=builder /root/monetachain/build/monetachaind /usr/local/bin/monetachaind
COPY ./bootstrap /usr/local/bootstrap

ENTRYPOINT ["/usr/local/bootstrap/start.sh"]
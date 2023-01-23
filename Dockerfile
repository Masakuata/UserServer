FROM golang:1.19.2 AS builder

COPY . /xgus

WORKDIR /xgus

RUN go get && CGO_ENABLED=0 GOOS=linux go build -o xgus .

FROM scratch

COPY --from=builder /xgus/xgus /xgus/xgus

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 42100

WORKDIR /xgss

ENTRYPOINT ["./xgus"]
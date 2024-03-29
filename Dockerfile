FROM golang:1.22 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build  -installsuffix cgo -ldflags="-s -w" -o /radix-log-api .

RUN useradd -M --uid 1000 radix-log-api

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /radix-log-api /usr/local/bin/radix-log-api
COPY --from=builder /etc/passwd /etc/passwd
USER 1000

ENTRYPOINT ["/usr/local/bin/radix-log-api"]

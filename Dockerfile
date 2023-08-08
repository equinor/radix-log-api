FROM golang:1.20 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
COPY Makefile ./

COPY . .

RUN make test
RUN make staticcheck
RUN make

RUN useradd -M --uid 1000 radix-log-api

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/bin/radix-log-api /usr/local/bin/radix-log-api
COPY --from=builder /etc/passwd /etc/passwd
USER 1000

ENTRYPOINT ["/usr/local/bin/radix-log-api"]
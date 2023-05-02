FROM golang:1.20-alpine as builder

RUN apk update && \
    apk add ca-certificates && \
    apk add --no-cache gcc musl-dev 

RUN go install honnef.co/go/tools/cmd/staticcheck@v0.4.3

WORKDIR /go/src/github.com/equinor/radix-log-api/

# get dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy api code
COPY . .

# lint and unit tests
RUN staticcheck ./... && \
    go vet ./... && \
    CGO_ENABLED=0 GOOS=linux go test ./...

# Build radix vulnerability scanner API go project
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o /usr/local/bin/radix-log-api

RUN addgroup -S -g 1000 radix-log-api
RUN adduser -S -u 1000 -G radix-log-api radix-log-api

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/radix-log-api /usr/local/bin/radix-log-api
COPY --from=builder /etc/passwd /etc/passwd
USER 1000
EXPOSE 3003
ENTRYPOINT ["/usr/local/bin/radix-log-api"]
# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /out/qgate ./cmd/qgate
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /out/webhook ./cmd/webhook

# Runtime image for the qgate CLI
FROM gcr.io/distroless/static-debian12:nonroot AS qgate
WORKDIR /
COPY --from=builder /out/qgate /usr/local/bin/qgate
USER 65532:65532
ENTRYPOINT ["/usr/local/bin/qgate"]

# Runtime image for the Kubernetes Validating Webhook
FROM gcr.io/distroless/static-debian12:nonroot AS webhook
WORKDIR /
COPY --from=builder /out/webhook /usr/local/bin/webhook
USER 65532:65532
EXPOSE 8443
ENTRYPOINT ["/usr/local/bin/webhook"]

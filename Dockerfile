FROM golang:1.24-bookworm AS builder

RUN apt update
RUN apt upgrade -y

WORKDIR /app

COPY . .

RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts
RUN git config --global url."ssh://git@github.com/".insteadOf "https://github.com/"
RUN --mount=type=ssh go mod download

RUN GOOS=linux go build -o api ./cmd/api/main.go

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/api /

CMD [ "/api" ]

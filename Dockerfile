FROM golang:1.24-bookworm AS builder

RUN apt update
RUN apt upgrade -y

WORKDIR /app

COPY . .

RUN  go mod download && go mod verify

RUN mkdir -p /bin
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v -o /bin/api ./cmd/api/main.go

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# force rebuild for this part
ARG BUILD_DATE=$(date +%s)
LABEL rebuild_trigger=$BUILD_DATE
COPY --from=builder /bin/api /api

EXPOSE 8080 

CMD [ "/api" ]

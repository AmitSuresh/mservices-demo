FROM golang:latest

RUN apt-get update && \
    apt-get install -y ca-certificates grep && \
    rm -rf /var/lib/apt/lists/*

RUN addgroup --system app && \
    adduser --system --ingroup app app

WORKDIR /app

COPY . .

RUN go install github.com/grpc-ecosystem/grpc-health-probe@latest

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -o server ./main.go

RUN chown -R app:app /app

USER app

EXPOSE 9082

ENTRYPOINT ["/app/server"]


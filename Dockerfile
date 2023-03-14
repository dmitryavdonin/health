
FROM golang:1.18-alpine AS builder
WORKDIR /app 
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o health .



FROM scratch
COPY --from=builder /app/health /usr/bin/
EXPOSE 8000
CMD ["/usr/bin/health"]

# syntax=docker/dockerfile:1
FROM golang:1.20
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /workspace/internal/config/hello-service.json ./internal/config/hello-service.json
COPY --from=0 /workspace/app ./
CMD ["./app"]  
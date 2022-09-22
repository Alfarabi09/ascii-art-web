FROM golang:1.17-alpine AS builder
WORKDIR /app 
COPY . .
 
RUN  go build -o main cmd/main.go

FROM alpine:3.6
WORKDIR /app
LABEL  captain="Alfarabi09" teammate="Mike7" port:="7777" project="ASCII-ART-WEB|DOCKERIZE"
COPY --from=builder /app/main .
COPY --from=builder /app/pkg/ascii/files  /app/pkg/ascii/files
COPY --from=builder /app/template /app/template



CMD [ "/app/main" ]
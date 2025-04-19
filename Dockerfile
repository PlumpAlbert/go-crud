FROM golang:alpine AS builder
WORKDIR /build 
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o gocrud ./cmd/gocrud

FROM scratch AS production
WORKDIR /app
COPY .env.example .env
COPY ./templates/ ./templates
COPY --from=builder /build/gocrud .
EXPOSE 8080
CMD ["/app/gocrud"]

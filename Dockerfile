FROM golang:1.18.10-bullseye

# build phase
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build ./cmd/main.go

CMD ["./main"]



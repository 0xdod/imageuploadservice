FROM golang:1.18-bullseye AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download -x

COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.release=`git rev-parse --short=8 HEAD`'" -o /bin/image-uploader ./cmd/image-uploader

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY .config.yaml ./

COPY --from=builder /bin/image-uploader ./

CMD ["./image-uploader"]
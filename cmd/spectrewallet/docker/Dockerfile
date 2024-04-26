# -- multistage docker build: stage #1: build stage
FROM golang:1.18-alpine AS build

RUN mkdir -p /go/src/github.com/spectre-project/spectred/

WORKDIR /go/src/github.com/spectre-project/spectred/

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /go/src/github.com/spectre-project/spectred/cmd/spectrewallet
RUN GOOS=linux go build -a -installsuffix cgo -o spectrewallet .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/spectre-project/spectred/cmd/spectrewallet/spectrewallet /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]

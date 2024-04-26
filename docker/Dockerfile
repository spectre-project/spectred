# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/spectre-project/spectred/

WORKDIR /go/src/github.com/spectre-project/spectred/

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache spectred dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o spectred .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/spectre-project/spectred/spectred /app/
COPY --from=build /go/src/github.com/spectre-project/spectred/infrastructure/config/sample-spectred.conf /app/

USER root
ENTRYPOINT [ "/sbin/tini", "--" ]

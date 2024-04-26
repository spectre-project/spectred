ARG SPECTRED_IMAGE
ARG SPECTREMINER_IMAGE

FROM ${SPECTRED_IMAGE} as spectred
FROM ${SPECTREMINER_IMAGE} as spectreminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/spectre-project/spectred/

WORKDIR /go/src/github.com/spectre-project/spectred/

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=spectred /app/ /app/
COPY --from=spectreminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/spectre-project/spectred/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]

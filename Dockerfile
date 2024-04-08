########################################
### BASE IMAGE #########################
########################################
ARG GO_VERSION=1.22.2
ARG ALPINE_VERSION=3.19

# Choosing golang alpine due to small footprint
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS base

ENV CGO_ENABLED=0 \
    GIN_MODE=release

# Go packages setup
WORKDIR $GOPATH/
COPY go.* ./
RUN go install ./...

# Build
COPY . ./
RUN gofmt -w -s . && \
    CGO_ENABLED=${CGO_ENABLED} GIN_MODE=${GIN_MODE} go build -o docker-example-app cmd/main.go

########################################
### PRODUCTION #########################
########################################
ARG USER_UID=1002

FROM alpine:${ALPINE_VERSION} AS production

# Declare variables to be used in docker
ARG USER_GID=${USER_UID}

# Update the docker packages
#RUN apk -U upgrade

COPY --from=base /go/dev.env/ ../
COPY --from=base /go/docker-example-app .

# Required. Run service as non-root.
USER ${USER_UID}

#HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
#  CMD wget -Y off -O /dev/null http://localhost:8080/docker-example/ping || exit 1

CMD ["./docker-example-app"]

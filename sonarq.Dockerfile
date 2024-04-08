########################################
### BUILDER IMAGE ######################
########################################
ARG GO_VERSION=1.22.2
ARG ALPINE_VERSION=3.19

# Choosing golang alpine due to small footprint
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

# Go packages setup
WORKDIR $GOPATH

ENV CGO_ENABLED=0 \
    GIN_MODE=release

# copy go.sum and go.mod
COPY go.* ./

# Build
COPY . ./

RUN go test ./... && gofmt -w -s . && \
    CGO_ENABLED=${CGO_ENABLED} GIN_MODE=${GIN_MODE} go build -o hello-app cmd/main.go


########################################
### SONAR-SCANNER IMAGE ################
########################################
# Etapa de análisis estático
FROM sonarsource/sonar-scanner-cli AS sonar-scanner
WORKDIR $GOPATH

# Copiar el código fuente de tu aplicación desde la etapa de compilación
COPY --from=builder /go/ .

# Ejecutar el análisis estático utilizando SonarScanner
RUN sonar-scanner \
        -Dsonar.projectKey=docker-example \
        -Dsonar.sources=. \
        -Dsonar.host.url=http://localhost:9003 \
        -Dsonar.token=sqp_bfbcbb1bc5b735e96f6e940dfbe1dc8944390a52

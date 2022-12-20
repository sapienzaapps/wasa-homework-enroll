# This file is used by Docker "build" or "buildah" to create a container image for this Go project
# The build is done using "multi-stage" approach, where a temporary container ("builder") is used to build the Go
# executable, and the final image is from scratch (empty container) for both security and performance reasons.

# DO NOT MODIFY UNLESS IT'S STRICTLY NECESSARY

ARG DOCKER_PREFIX
FROM ${DOCKER_PREFIX}node:lts AS uibuilder
COPY webui webui
WORKDIR webui
RUN npm config set update-notifier false && npm install && npm run build-embed

ARG DOCKER_PREFIX
FROM ${DOCKER_PREFIX}enrico204/golang:1.19.4-6 AS builder

# Disable Go proxy and public checksum for private repositories (Go 1.13+)
ENV GOPRIVATE git.sapienzaapps.it

### Copy Go code
COPY . .
COPY --from=uibuilder webui webui

### Set some build variables
ARG APP_VERSION
ARG BUILD_DATE
ARG REPO_HASH

RUN go generate -mod=vendor ./...

### Build executables, strip debug symbols and compress with UPX
WORKDIR /src/cmd/
RUN /bin/bash -euo pipefail -c "for ex in \$(ls); do pushd \$ex; go build -tags webui,openapi -mod=vendor -ldflags \"-extldflags \\\"-static\\\" -X main.AppVersion=${APP_VERSION} -X main.BuildDate=${BUILD_DATE}\" -a -installsuffix cgo -o /app/\$ex .; popd; done"
RUN cd /app/ && strip * && upx -9 *

### Create final container from scratch
FROM scratch

### Inform Docker about which port is used
EXPOSE 3000 4000

### Populate scratch with CA certificates and Timezone infos from the builder image
ENV ZONEINFO /zoneinfo.zip
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /zoneinfo.zip /
COPY --from=builder /etc/passwd /etc/passwd

### Copy the build executable from the builder image
WORKDIR /app/
COPY --from=builder /app/* ./

### Set some build variables
ARG APP_VERSION
ARG BUILD_DATE
ARG PROJECT_NAME
ARG GROUP_NAME

### Downgrade to user level (from root)
USER appuser

### Executable command
CMD ["/app/webapi"]

### OpenContainers tags
LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="${GROUP_NAME} - ${PROJECT_NAME}" \
      org.opencontainers.image.authors="SapienzaApps <sapienzaapps@gmail.com>" \
      org.opencontainers.image.source="https://git.sapienzaapps.it/${GROUP_NAME}/${PROJECT_NAME}" \
      org.opencontainers.image.revision="${REPO_HASH}" \
      org.opencontainers.image.vendor="SapienzaApps" \
	  org.opencontainers.image.version="${APP_VERSION}"

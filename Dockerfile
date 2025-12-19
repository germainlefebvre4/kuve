FROM alpine:3 AS build

ARG DISTRIBUTION=linux
ARG CPU_ARCH=amd64
ARG JQ_VERSION=1.7
ARG KUVE_VERSION=0.1.0

WORKDIR /app

# SHELL ["/bin/ash", "-o", "pipefail", "-c"]
RUN apk update && \
    apk add --no-cache \
        curl
RUN curl --output jq-linux64 https://github.com/stedolan/jq/releases/download/jq-${JQ_VERSION}/jq-linux64 && \
    mv jq-linux64 /usr/local/bin/jq && \
    chmod +x /usr/local/bin/jq

RUN VERSION=$(curl -s "https://api.github.com/repos/germainlefebvre4/kuve/releases/tags/v${KUVE_VERSION}" | jq -r '.tag_name') && \
    curl -L --output /app/kuve "https://github.com/germainlefebvre4/kuve/releases/download/${VERSION}/kuve_${DISTRIBUTION}_${CPU_ARCH}" && \
    chmod +x /app/kuve


FROM alpine:3

COPY --from=build /app/kuve /usr/local/bin/kuve

WORKDIR /cv

ENTRYPOINT ["kuve"]
CMD ["--help"]

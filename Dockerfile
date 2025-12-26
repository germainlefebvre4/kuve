FROM alpine:3 AS build

ARG DISTRIBUTION=linux
ARG CPU_ARCH=amd64
ARG KUVE_VERSION=0.2.1

WORKDIR /app

RUN apk update && \
    apk add --no-cache curl ca-certificates

RUN curl -L --retry 5 --retry-delay 2 --retry-all-errors \
    --output kuve.tar.gz \
    "https://github.com/germainlefebvre4/kuve/releases/download/v${KUVE_VERSION}/kuve_${DISTRIBUTION}_${CPU_ARCH}.tar.gz" && \
    tar -zxvf kuve.tar.gz && \
    chmod +x kuve


FROM alpine:3

COPY --from=build /app/kuve /usr/local/bin/kuve

WORKDIR /app

ENTRYPOINT ["kuve"]
CMD ["--help"]

FROM alpine:3.4

RUN apk --no-cache add \
    apache2-utils

ENTRYPOINT ["tail", "-f", "/dev/null"]
FROM alpine:3.4

RUN apk --no-cache add \
    apache2-utils \
    jq \
    curl

WORKDIR /
COPY ./push_record.sh /
ENTRYPOINT ["tail", "-f", "/dev/null"]
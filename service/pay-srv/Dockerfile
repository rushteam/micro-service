FROM alpine

RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*

WORKDIR /usr/bin/
COPY ./dist/ /usr/bin/
COPY ./service/pay-srv/config.yaml /usr/bin/config.yaml
# CMD pay-srv --config_path=config.yaml
CMD pay-srv
# CMD ["pay-srv","--config_path=config.yaml"]
ENTRYPOINT [ "/pay-srv" ]
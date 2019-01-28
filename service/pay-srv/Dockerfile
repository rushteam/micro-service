FROM alpine

COPY ./dist/ /usr/bin/
COPY ./service/pay-srv/config.yaml /usr/bin/config.yaml
CMD ["pay-srv","--config_path" "config.yaml"]
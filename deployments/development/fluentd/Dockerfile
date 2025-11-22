FROM fluent/fluentd:edge-debian

USER root

# Elasticsearch 8.x uyumlu gem
RUN gem install elasticsearch --no-document --version 8.19.0

# Fluentd Elasticsearch plugin
RUN gem install fluent-plugin-elasticsearch --no-document --version 5.4.3

# Config dosyasını kopyala
COPY ./conf/fluentd.conf /fluentd/etc/fluentd.conf

# Log dizini
RUN mkdir -p /fluentd/log

USER fluent

EXPOSE 24224
EXPOSE 24224/udp

CMD ["fluentd", "-c", "/fluentd/etc/fluentd.conf", "-v"]
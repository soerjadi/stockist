FROM golang:1.23-alpine3.20 AS builder

# Build
COPY . /src/app

WORKDIR /src/app

RUN apk update \
  && apk upgrade \
  && apk add \
  git build-base ca-certificates && \
  update-ca-certificates 2>/dev/null || true

RUN mkdir -p /src/app/bin && \
  mkdir -p /src/app/conf && \
  mkdir -p /var/log/stockist && \
  mkdir -p /etc/stockist
  
# RUN  mkdir supervisor-src && \
#   cd supervisor-src && \
#   git clone https://github.com/ochinchina/supervisord.git . && \
#   go generate && \
#   GOOS=linux go build -tags release -a -ldflags "-linkmode external -extldflags -static" -o /src/app/bin/supervisord

# ARG SUPERVISOR_CONF

# COPY ./script/${SUPERVISOR_CONF} /src/app/conf/supervisor.conf

COPY ./files/config.development.ini /etc/stockist/config.ini

ARG CACHEBUILD

# to prevent cache on building our app
RUN cd cmd/rest && \
  GOOS=linux go build -tags release -a -ldflags "-linkmode external -extldflags -static" -o /src/app/bin/stockist && \
  echo $CACHEBUILD


# Serve
FROM busybox

RUN mkdir -p /etc/stockist/

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/app/bin/stockist /usr/app/stockist
COPY --from=builder /etc/stockist/config.ini /etc/stockist/config.ini


RUN ls /usr/app/

CMD [ "/usr/app/stockist" ]
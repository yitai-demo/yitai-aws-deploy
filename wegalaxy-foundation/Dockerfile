ARG PRJNAME="wegalaxy-foundation"
ARG PRJPATH="github.com/degalaxy/wegalaxy-service/${PRJNAME}"
ARG PRJROOT="github.com/degalaxy/wegalaxy-service"

FROM golang:1.17-alpine AS builder
ARG GITHUB_TOKEN

RUN apk update
RUN apk add git

ARG PRJNAME
ARG PRJPATH
ARG PRJROOT

ADD . /go/src/${PRJROOT}
WORKDIR /go/src/${PRJPATH}
RUN git config --global url."https://${GITHUB_TOKEN}:@github.com/".insteadOf "https://github.com/"
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /${PRJNAME} /go/src/${PRJPATH}

FROM alpine:3.15 as base

ARG PRJNAME
ARG PRJPATH

ENV BUILD_DEPS="gettext"  \
    RUNTIME_DEPS="libintl"

RUN set -x && \
    apk add --update $RUNTIME_DEPS && \
    apk add --virtual build_deps $BUILD_DEPS &&  \
    cp /usr/bin/envsubst /usr/local/bin/envsubst && \
    apk del build_deps
    
COPY --from=builder /go/src/${PRJPATH}/config/config.template.yaml  /config/config.template.yaml
COPY --from=builder /${PRJNAME} /${PRJNAME}
COPY --from=builder /go/src/${PRJPATH}/start.sh /start.sh
RUN  chmod +x /start.sh
ENTRYPOINT ["/start.sh"]

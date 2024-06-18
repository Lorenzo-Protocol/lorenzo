FROM golang:1.20-alpine AS build

RUN apk add build-base git linux-headers

WORKDIR /work
COPY go.mod go.sum /work/
RUN go mod download

COPY ./ /work
RUN LEDGER_ENABLED=false make clean build


FROM alpine AS run
RUN apk add bash curl jq
EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["start", "--log_format", "plain"]
STOPSIGNAL SIGTERM
VOLUME /lorenzod
WORKDIR /lorenzod

COPY contrib/images/lorenzod-env/wrapper.sh /usr/bin/wrapper.sh
COPY --from=build /work/build/lorenzod /lorenzod/
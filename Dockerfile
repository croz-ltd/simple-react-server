FROM golang:alpine AS builder

LABEL maintainer="mjovanovic@croz.net"

ENV USER=srs
ENV UID=10001

WORKDIR /workspace
COPY . .

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

RUN go get -d -v && \
    CGO_ENABLED=0 GOOS=linux go build -a -o simple-react-server && \
    chmod a+x simple-react-server

##
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /workspace/simple-react-server /go/bin/simple-react-server

USER srs:srs

ENTRYPOINT ["/go/bin/simple-react-server"]
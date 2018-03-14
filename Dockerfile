FROM alpine:3.7

RUN apk add --no-cache \
    ca-certificates \
    git \
    go \
    musl-dev \
    tzdata

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

ENV CFG_REDISPORT "6379"
ENV CFG_PORT "80"

RUN mkdir -p "$GOPATH/src/github.com/jlyon1/terraformchat" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/github.com/jlyon1/terraformchat

RUN git clone https://github.com/jlyon1/terraformchat.git ./

RUN go get -v
RUN go build -o terraformchat

EXPOSE 80

CMD ["./terraformchat"]

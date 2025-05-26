FROM golang:1.20-alpine as buildbase

RUN apk add build-base git



WORKDIR /go/src/github.com/hyle-team/bridgeless-core/v12

ENV GO111MODULE="on"
ENV CGO_ENABLED=1
ENV GOOS="linux"
ENV GOPRIVATE=github.com/*
ENV GONOSUMDB=github.com/*
ENV GONOPROXY=github.com/*

COPY ./go.mod ./go.sum ./
# Read the CI_ACCESS_TOKEN from the .env file
ARG CI_ACCESS_TOKEN
RUN git config --global url."https://olegfomenkodev:${CI_ACCESS_TOKEN}@github.com/".insteadOf "https://github.com/"
RUN go mod download

COPY . .

RUN go mod vendor
RUN git clone https://github.com/hyle-team/cosmos-sdk.git && \
    cd cosmos-sdk/cosmovisor && \
    git checkout 7e87d2103e03dca96073cb519bf3cbe09551c014 && \
    go install ./cmd/cosmovisor

RUN cp $GOPATH/bin/cosmovisor /usr/local/bin/cosmovisor

RUN go build  -mod=mod  -o /usr/local/bin/bridgeless-core github.com/hyle-team/bridgeless-core/v12/cmd/bridgeless-cored



###

FROM alpine:3.9

RUN apk add --no-cache ca-certificates
COPY ./config/genesis.json /config/genesis.json
COPY --from=buildbase /usr/local/bin/bridgeless-core /usr/local/bin/bridgeless-core
COPY --from=buildbase /usr/local/bin/cosmovisor /usr/local/bin/cosmovisor
ENTRYPOINT ["bridgeless-core"]
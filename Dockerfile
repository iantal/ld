FROM golang:alpine as builder

ENV GO111MODULE="" \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN apk add --no-cache git

ARG DOCKER_NETRC
RUN echo "${DOCKER_NETRC}" > ~/.netrc
RUN go mod download

COPY . .
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .


FROM golang:alpine as deploy
COPY --from=builder /dist .

ENV VERSION="7.9.0"

RUN apk --update --no-cache add \
    ruby \
    ruby-rugged \
    ruby-charlock_holmes \
    ruby-json \
  && apk --update --no-cache add -t build-dependencies \
    build-base \
    cmake \
    icu-dev \
    ruby-dev \
    unzip \
    zip \
  && apk add git \
  && gem install --no-document github-linguist -v ${VERSION} \
  && apk del build-dependencies \
  && rm -rf /var/cache/apk/*


ENV BASE_PATH="/opt/data"
VOLUME [ "/opt/data" ]
RUN git config --global user.email "test@rm.com" && git config --global user.name "testrm"
EXPOSE 8003
CMD ["./main"]
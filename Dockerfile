FROM golang:1.13-alpine

ARG ARG_YT_API_KEY
ARG ARG_USERNAME_DB
ARG ARG_PASSWORD_DB
ARG ARG_VERSION
ARG ARG_HOST_SUB

ENV YT_API_KEY=$ARG_YT_API_KEY
ENV USERNAME_DB=$ARG_USERNAME_DB
ENV PASSWORD_DB=$ARG_PASSWORD_DB
ENV VERSION=$ARG_VERSION
ENV HOST_SUB=$ARG_HOST_SUB

ENV CORS_ORIGIN=http://dadard.fr

RUN apk add --update git gcc libc-dev wget

RUN wget https://yt-dl.org/downloads/latest/youtube-dl -O /usr/local/bin/youtube-dl
RUN chmod a+rx /usr/local/bin/youtube-dl

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
## BUILDER
FROM golang:1.22-alpine as builder

ARG GH_PAT

WORKDIR /src

COPY . .

RUN echo "machine github.com\n\tlogin bot\n\tpassword ${GH_PAT}" >> ~/.netrc && \
    go mod download && \
    rm -f ~/.netrc

RUN go build -o myaction ./cmd/myaction


## DEPLOY
FROM alpine:3

RUN apk update

WORKDIR /cmd

COPY --from=builder /src/myaction /cmd/myaction

CMD ["/cmd/action"]
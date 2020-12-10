# ベースとなるDockerイメージ指定
FROM golang:alpine
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/app
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/app

RUN apk add --no-cache \
  alpine-sdk \
  git \
  && go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
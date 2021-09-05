FROM golang:1.12 as build

WORKDIR /kokoko
# ADDよりCOPYが推奨されている
COPY . /kokoko

RUN apt-get update && apt-get install -y \
        sudo \
        curl \
        git \
        make \
        less \
        vim \
  && go build -o main

# freshのインストール
RUN go get github.com/pilu/fresh
RUN go get github.com/go-delve/delve/cmd/dlv@v1.3.0

CMD fresh

# exec と　バイナリのファイル場所
# CMD ["dlv", "exec", "./tmp/runner-build", "--listen=:12345", "--headless=true", "--api-version=2", "--accept-multiclient", "--continue"]

# COPY ./entrypoint.sh /sh/entrypoint.sh
# RUN chmod 755 /sh/entrypoint.sh
# ENTRYPOINT ["/sh/entrypoint.sh"]

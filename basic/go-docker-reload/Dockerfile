# 元となるイメージを指定
FROM golang:1.13

# 作業ディレクトリを指定
WORKDIR /go/src/go-echo-docker-example
# 左側(ローカル)のディレクトリをイメージの作業ディレクトリ(ここではWORKDIRで指定した/go/src/go-echo-docker-example)にコピー
COPY . .

# イメージのビルド時に実行するコマンド
# freshのインストール
RUN go get github.com/pilu/fresh
# イメージからコンテナを作成する際に実行
CMD ["fresh"]

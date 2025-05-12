# 開発用のイメージ
FROM golang:1.24.2 AS dev

# 作業ディレクトリの設定
WORKDIR /app

# `air` のインストール
RUN go install github.com/air-verse/air@latest

# go.mod, go.sum をコピーして依存関係を解決
COPY go.mod go.sum ./
RUN go mod download

# 残りのソースコードをコピー
COPY . .

# ポートを開放（Echo のデフォルトポート）
EXPOSE 8080

# `air` を実行してホットリロード
CMD ["air"]
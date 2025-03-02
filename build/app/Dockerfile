# ベースイメージ: Go 1.22 (Alpine)
FROM golang:1.22-alpine

# govulncheck のインストール
RUN go install golang.org/x/vuln/cmd/govulncheck@latest

# 作業ディレクトリの設定
WORKDIR /app

# ローカルのgo.modとgo.sumをコピーし、依存関係を取得
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションのビルド
RUN go build -o app ./cmd

# コンテナ起動時は、govulncheckで脆弱性診断を実行（GithubActionsで実行する場合は別途ワークフローを用意）
CMD ["govulncheck", "./..."]
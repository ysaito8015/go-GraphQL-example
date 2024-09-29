## gqlgen のインストール


```shell
$ go install github.com/99designs/gqlgen@latest
$ gqlgen version
```


## パッケージの導入


```shell
$ go mod init mygraphql
$ go get -u github.com/99designs/gqlgen
```


## gqlgen の初期化


```shell
$ gqlgen init
```


## 生成されたファイルツリー


```shell
.
├─ graph
│   ├─ generated.go         # 生成されたリゾルバ（特定のフィールドのデータを返す関数）をサーバで稼働させるロジック
│   ├─ model
│   │   └─ models_gen.go    # GraphQL のスキーマ (`schema.graphql`) に対応する Go の構造体型が定義されている
│   ├─ resolver.go          # リゾルバ構造体の定義
│   ├─ schema.graphqls      # GraphQL のスキーマ定義
│   └─ schema.resolvers.go  # リゾルバ構造体のメソッド定義
├─ gqlgen.yml               # gqlgen で生成する際の設定記述
├─ server.go                # サーバエントリポイント
├─ go.mod
└─ go.sum
```


## GraphQL サーバの起動

```shell
$ go mod tidy
$ go run ./server.go
```

## ハマったとこ
- `*.graphqls` 拡張子は複数形にしないと `gqlgen generate` で空のファイルが生成される
- `./internal/server/main.go` にサーバのエントリポイントを作る
- `package main` じゃないファイルに `go run` すると `package command-line-arguments is not a main package` が表示される


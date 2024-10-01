## メモ
- `./bin/setup.sh` の `TIMESTAMP` を `DATETIME` に置き換える
    - https://github.com/saki-engineering/graphql-sample/issues/1
- `"message": "diredtive isAuthenticated is not implemented"` と帰ってくる
    - directive については、この章には書かれていない


## 流れ
1. DB の準備
2. ORM コードの生成
3. DB 操作を直接扱わないように services を導入
4. services を利用して `schema.resorver.go` に実装
5. `resolver.go` の構造体 `Resorver` に services を inject する
6. `internal/server/main.go` に DB への接続、service の生成、Resover への service セットを実装

## メモ
- GraphQL クエリで取得するフィールドはすべてスカラ型
- すでに用意されている 5 つのスカラ型以外は、カスタムスカラ型として定義する
- カスタムスカラの作り方
    1. カスタムスカラ対応型を使う
        - `github.com/99designs/gqlgen/graphql` に用意されている
    2. 独自の型を作り、その型に `MarshalGQL/UnmarshalGQL` メソッドを実装する
    3. 直接扱えるように `MarshalXxx/UnmarshalXxx` 関数を実装する

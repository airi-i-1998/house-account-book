```mermaid
sequenceDiagram
participant A as 認証サーバー
participant B as ウェブサーバー
participant C as アプリ
Note over A, C: ユーザー登録
C->>A: ユーザー登録リクエストAPI
A-->>C: サインインレスポンス (Token: json)
Note left of C: Tokenはアプリ内に保存

Note over A, C: WebView アクセス
C->>B: HTTPリクエスト (Token: HTTP Header)
B->>A: ユーザー登録確認 (Token: json)
A-->>B: 新規ユーザーページ
B-->>C: HTTPレスポンス
````

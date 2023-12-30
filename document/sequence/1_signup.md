```mermaid
sequenceDiagram
participant A as ブラウザ
participant B as Webサーバー
Note over A, B: ユーザー登録
A->>B: HTTP POST要求<br>(ユーザー登録リクエスト)
B-->>A: HTTP POST応答

````

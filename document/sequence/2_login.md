```mermaid
sequenceDiagram
participant A as ブラウザ
participant B as Webサーバー
Note over A, B: ログイン
A->>B: HTTP POST要求<br>(ログインリクエスト)
B-->>A: HTTP POST応答<br>(ホームページへ遷移)

````

## 家計簿アプリ

### 概要

- 概要の説明
- デモ画面のスクリーンショット

### 機能のPRDと技術仕様書の一覧

| 機能名称 | PRDリンク | 技術仕様書のリンク |
| :--- | :--- | :--- |
| ユーザー登録機能 | [/document/prd/1_signup.md](/document/prd/1_signup.md) | [/document/techspec/1_signup.md](/document/techspec/1_signup.md) |
| ログイン機能 | XXXXXXXX | XXXXXXXX |
| ログアウト機能 | XXXXXXXX | XXXXXXXX |
| 支出追加機能 | XXXXXXXX | XXXXXXXX |
| 支出更新機能 | XXXXXXXX | XXXXXXXX |
| 支出削除機能 | XXXXXXXX | XXXXXXXX |
| 収入追加機能 | XXXXXXXX | XXXXXXXX |
| 収入更新機能 | XXXXXXXX | XXXXXXXX |
| 収入削除機能 | XXXXXXXX | XXXXXXXX |

- [PRDのフォーマット例](/document/prd/1_login.md)
- [技術仕様書のフォーマット例  ](/document/techspec/1_login.md)


### 技術選定とその背景

- バックエンド
  - Go
    - 記述方法がシンプルでわかりやすい
    - 高速処理が可能
    - 並行処理・並列処理が可能

- フロントエンド
  - React
    - 宣言的なViewによってコードの可読性が向上する
    - コンポーネントベースのため様々なデータの取り出し、再利用がしやすい
    - コンポーネント化することで機能の変更、追加がしやすいため

### インフラ設計図

- 後で作成

### API一覧
- [APIスキーマのリンク](/document/api-schema/openapi.yml)

| 機能名称 | APIのエンドポイント |
| :--- | :--- |
| ユーザー登録API | [localhost:3000/signup](localhost:3000/signup) |
| ログインAPI | XXXXXXXX |
| 残高確認API | XXXXXXXX |

- スキーマ作成時の参照先
  - [Swagger公式doc](https://swagger.io/docs/specification/api-host-and-base-path/)
  - [Swagger拡張機能](https://marketplace.visualstudio.com/items?itemName=Arjun.swagger-viewer)

### テーブル一覧

- ER図の画像
  - XXXXXX
- [ER図のリンク](/document/db/db_er.dio)
- ER図作成時の参照先
  - [Draw.io VSCode拡張機能](https://marketplace.visualstudio.com/items?itemName=hediet.vscode-drawio)

### ディレクトリ構成の説明

- XXX
- XXX
  - XXX

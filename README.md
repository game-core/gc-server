# gc-server
ゲームサーバー
![GitHub repo size](https://img.shields.io/github/repo-size/game-core/gc-server)
<img src="https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic">
<img src="https://img.shields.io/badge/-Nuxt.js-00C58E.svg?logo=nuxt.js&style=plastic">
<img src="https://img.shields.io/badge/-Typescript-007ACC.svg?logo=typescript&style=plastic">
<img src="https://img.shields.io/badge/-Mysql-4479A1.svg?logo=mysql&style=plastic">
<img src="https://img.shields.io/badge/-Redis-D82C20.svg?logo=redis&style=plastic">
<img src="https://img.shields.io/badge/-Kubernetes-326CE5.svg?logo=kubernetes&style=plastic">

ゲームサーバー基盤(gRPC)

## アーキテクチャ
![architecture drawio](https://github.com/game-core/gc-server/assets/71867595/d43b608d-89eb-4d32-8bb1-b2fac7fa9815)

## 機能
- API
    - [Account：アカウント](https://github.com/game-core/gc-server/blob/main/docs/md/function/api/account.md)
    - [LoginBonus：ログインボーナス](https://github.com/game-core/gc-server/blob/main/docs/md/function/api/loginBonus.md)
- Service
    - [Account：アカウント](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/account.md)
    - [Action：ユーザー行動管理](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/action.md)
    - [Event：イベント](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/event.md)
    - [Item：アイテム](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/item.md)
    - [LoginBonus：ログインボーナス](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/loginBonus.md)
    - [Shard：シャード管理](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/shard.md)
    - [Transaction：トランザクション管理](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/transaction.md)
## リンク
- [環境構築](./docs/md/environment.md)
- [自動生成](./docs/md/generator.md)

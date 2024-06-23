# 環境構築

## Server

- 起動
```
make docker_up
```
- マイグレーション
  - `./docs/sql`配下のsqlファイルが実行される
```
make gen_migration
```
- マスターデータインポート
  - 事前に[GASの設定]()を行う
```
make gen_master
```

## Test
- テスト実行
```
make gen_test
```

# Item
アイテム関連関連。  
[model](https://github.com/game-core/gc-server/tree/main/pkg/domain/model/item)

- [Receive](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/item.md#Receive)

## Create
アイテムを受け取る。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| req | *ItemReceiveRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *ItemReceiveResponse | レスポンス |
| err | error | エラー |

# LoginBonus
ログインボーナス関連。  
[model](https://github.com/game-core/gc-server/tree/main/pkg/domain/model/loginBonus)

- [GetUser](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/loginBonus.md#GetUser)
- [GetMaster](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/loginBonus.md#GetMaster)
- [Receive](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/loginBonus.md#Receive)

## GetUser
ログインボーナスのユーザーステータスを取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| req | *LoginBonusGetUserRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *LoginBonusGetUserResponse | レスポンス |
| err | error | エラー |

## GetMaster
ログインボーナスのマスターデータを取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| req | *LoginBonusGetMasterRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *LoginBonusGetMasterResponse | レスポンス |
| err | error | エラー |

## Receive
ログインボーナスを受けとる。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| now | time.Time | 現在時刻 |
| req | *LoginBonusReceiveRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *LoginBonusReceiveResponse | レスポンス |
| err | error | エラー |

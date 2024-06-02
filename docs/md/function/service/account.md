# Account
アカウント関連。  
[model](https://github.com/game-core/gc-server/tree/main/pkg/domain/model/account)

- [FindByUserId](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/account.md#FindByUserId)
- [Create](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/account.md#Create)
- [Login](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/account.md#Login)
- [Get](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/account.md#Get)
- [CreateUserId](https://github.com/game-core/gc-server/blob/main/docs/md/function/service/account.md#CreateUserId)

## FindByUserId
ユーザーIDから取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| userId | string | ユーザーID |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *userAccount.UserAccount | ユーザーアカウントモデル |
| err | error | エラー |


## Create
アカウントを作成する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| req | *AccountCreateRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *AccountCreateResponse | レスポンス |
| err | error | エラー |

## Login
アカウントをログインする。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| req | *AccountLoginRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *AccountLoginResponse | レスポンス |
| err | error | エラー |

## Get
アカウントを取得する。
- request

| Name | Type               | Description |
| :--- |:-------------------| :--- |
| ctx | context.Context    | コンテキスト |
| req | *AccountGetRequest | リクエスト |

- response

| Name | Type                | Description |
| :--- |:--------------------| :--- |
| res | *AccountGetResponse | レスポンス |
| err | error               | エラー |

## CreateUserId
ユーザーIDを作成する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| userId | string | ユーザーID |
| err | error | エラー |

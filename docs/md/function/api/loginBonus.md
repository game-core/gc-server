# LoginBonus
ログインボーナス関連。  
[protobuf](https://github.com/game-core/gc-server/tree/main/docs/proto/api/game/loginBonus)  

- [Receive](https://github.com/game-core/gc-server/blob/main/docs/md/function/api/loginBonus.md#Receive)

## Receive
ログインボーナスを受けとる。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |

```json
{
    "user_id": "0:aP9UvDkOqvP5iW4YRSd6",
    "master_idle_bonus_id": 1
}
```
- response
```json
{
    "user_login_bonus": {
        "user_id": "0:ZJJrANH5F8gbNbusyH-9",
        "master_login_bonus_id": "1",
        "received_at": {
            "seconds": "1710605245",
            "nanos": 587063179
        }
    },
}
```

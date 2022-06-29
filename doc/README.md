# intro

透過分佈式算法`raft`，架構一套`slot`系統

# flow
login
```
C_S_LoginServer -> S_C_LoginServer -> S_C_LoginGame -> C_S_LoginGame
{
    C_S_MatchGame : 開始匹配遊戲
    S_C_MatchGame (取消匹配 C_S_CancelMatchGame) : 匹配中/取消匹配
    C_S_GetTableInfo : 獲取遊戲資訊
    S_C_SendTableInfo : 推送遊戲結果
}
S_C_SendLeave


# 主流程
    ### 進入遊戲
        S_C_LoginServer -> S_C_LoginServer: 玩家身份登入驗證
        S_C_LoginGame -> C_S_LoginGame : 登入新的遊戲廳房
    ### 開始玩遊戲
        C_S_MatchGame : 開始配對玩家
        S_C_MatchGame : 返回配對結果
        C_S_CancelMatchGame : 取消配對結果
    ### 遊戲結算
        C_S_GetTableInfo : 獲取遊戲內容
        S_C_SendTableInfo : 推送遊戲結果
        S_C_SendLeave : 送出玩家退出訊息

# 通用協議
    ### 心跳
        C_S_PING : Client Ping
        S_C_PONG : Server Pong

    ### 玩家
        C_S_UserTransfer : 獲取玩家目前金額資訊
        C_S_UserRecords : 玩家過去遊玩紀錄

    ### 系統
        S_C_BroadCast : 廣播公告
        S_C_KICKOUT : 踢玩家

```



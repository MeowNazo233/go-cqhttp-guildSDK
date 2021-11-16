# go-cqhttp-guildSDK
获取go-cqhttp的QQ频道消息

go-cqhttp设置为
-ws-reverse:
    disabled: false
    universal: ws://localhost:7790/ws/
    api: 
    event: 
    reconnect-interval: 3000
    middlewares:
    <<: *default 

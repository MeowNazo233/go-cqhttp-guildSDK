# go-cqhttp-guildSDK
获取go-cqhttp的QQ频道消息

go-cqhttp设置为
  - ws-reverse:
      # 是否禁用当前反向WS服务
      disabled: false
      # 反向WS Universal 地址
      # 注意 设置了此项地址后下面两项将会被忽略
      universal: ws://localhost:7790/ws/
      # 反向WS API 地址
      api: 
      # 反向WS Event 地址
      event: 
      # 重连间隔 单位毫秒
      reconnect-interval: 3000
      middlewares:
        <<: *default # 引用默认中间件

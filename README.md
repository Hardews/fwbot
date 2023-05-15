# fwbot

## 🐱‍🏍功能（待完善）

- [x] 私聊调用 GPT 的接口进行回复
- [x] 调用百度的 api 查看天气
- [x] 点歌
- [ ] 简单上下文？
- [ ] 好友请求处理
- [ ] 群聊



## 🚀快速开始（服务器需可膜法上网）

### fwbot

按注释填写 `config.yaml`，运行程序即可。

### go-cqhttp

移步官方文档：https://docs.go-cqhttp.org/guide/config.html

## 🍳配置

### 部分配置文件

```yaml
  # 反向WS设置
  - ws-reverse:
      # 反向WS Universal 地址
      # 注意 设置了此项地址后下面两项将会被忽略
      universal: ws://127.0.0.1:8077
      # 反向WS API 地址
      api: ws://your_websocket_api.server
      # 反向WS Event 地址
      event: ws://your_websocket_event.server
      # 重连间隔 单位毫秒
      reconnect-interval: 30000
      middlewares:
        <<: *default # 引用默认中间件
```

### 天气数据（不需要可以不管）

根据 `dump.sql` 还原天气所需要的城市信息

### 点歌功能（不需要可以不管）

歌曲 id 获取可使用项目：[Binaryify/NeteaseCloudMusicApi: 网易云音乐 Node.js API service (github.com)](https://github.com/Binaryify/NeteaseCloudMusicApi)

然后填写部署后的 url 

## 🔧设计说明

用反向ws进行通信。读取信息时用的是 ws 读取。

router 层主要是用来起服务，让 bot 可以成功连接。

成功连接之后，调用 service 层的 start 函数，然后开始监听信息并且发送信息的管道。

# WebSocket 帧同步联机 Demo

这是一个基于 WebSocket 实现的轻量级帧同步联机演示项目，适用于浏览器游戏开发学习和测试。

#### 功能特性

• 基于 WebSocket 的实时通信

• 帧同步机制实现

• 多客户端同步显示

• 简单易用的演示界面

## 快速开始

本地部署

1. 克隆仓库：

```bash
git clone https://github.com/beijian128/framesync.git
cd framesync
```

2. 运行服务端：

```bash
go run main.go
```

3. 在浏览器中访问：

```
http://localhost:8080/web/
```

可以打开多个浏览器窗口或标签页来模拟多个客户端联机效果。

## 在线体验

作者已在云端部署了一个服务，可直接访问体验：
[https://beijian99.top/framesync/](https://beijian99.top/framesync/)

## 技术栈

• 后端：Go

• 前端：HTML5 + JavaScript

• 通信协议：WebSocket


**核心代码见** services/gate/framesync.go (服务端)    和 web/index.html （客户端）

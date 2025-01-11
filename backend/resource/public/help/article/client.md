# 客户端使用教程


## OBS 和 websocket 插件


1. 下载并安装 OBS

    - 访问 [`https://obsproject.com/`](https://obsproject.com/) 官网
    - 选择对应操作系统版本下载

2. 下载并安装 OBS WebSocket 插件
    - 访问 [`https://github.com/obsproject/obs-websocket/releases`](https://github.com/obsproject/obs-websocket/releases)
    - 选择对应操作系统版本下载
        - 安装版本 直接安装
        - 绿色版本(ZIP) 直接解压到 OBS 安装目录 如 `C:\Program Files\obs-studio`

3. 启动 OBS

4. 顶部菜单栏依次点击 `工具`- `WebSocket服务器设置`, 弹出对话框
    - 插件设置
        - 勾选  [`√`] 开启WebSocket服务器
    - 服务器端口
        - 默认 4455
    - 服务器密码
        - 关闭 [` `] 开启身份认证
    - 确定并关闭对话框

5. 完成配置


6. 添加场景, 并配置来源

## 配置VPN

<!-- TODO -->

## 直播助手


1. 下载并安装

2. 启动和激活
    - 自动连接 OBS WebSocket, 并同步场景
3. 配置
    - 配置代理服务
    - 配置推流地址
    - 配置默认轮播节目
    - 配置触发动作和节目
4. 开始直播
# frp_android

`frp_android` 是一个面向 Android 使用场景的 frpc Go 封装层。项目基于 `github.com/fatedier/frp`，对外提供了较直接的配置对象和生命周期方法，便于通过 `gomobile bind` 生成 AAR 后在 Android 应用中调用。

当前仓库重点提供：

- `FrpcConfig`：客户端公共配置
- `FrpcProxyConfig`：代理配置
- `FrpcVisitorConfig`：访问者配置
- `Start` / `Stop` / `Reload`：启动、停止、重载 frpc

本文档中的示例参考了 Android `Service` 场景，但已去除实际业务域名、令牌、设备标识规则等隐私信息。

## 使用方法

### 1. 在 Android 项目中引入

将 AAR 放入 Android 项目的 `libs` 目录，然后在 Gradle 中声明：

```gradle
dependencies {
    implementation(files("libs/frpandroid.aar"))
}
```

### 2. 创建并启动 frpc

下面是一个脱敏后的 Kotlin 示例，适合放在 Android `Service` 或其他后台组件中调用：

```kotlin
package com.example.app

import android.app.Service
import android.content.Intent
import android.os.IBinder
import frpandroid.Frpandroid

class FrpcService : Service() {
    override fun onCreate() {
        super.onCreate()
        startFrpc()
    }

    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        return START_STICKY
    }

    override fun onBind(intent: Intent?): IBinder? = null

    override fun onDestroy() {
        super.onDestroy()
        Frpandroid.stop()
    }

    private fun startFrpc() {
        try {
            val cfg = Frpandroid.newFrpcConfig()
            cfg.setServerAddr("your-frps.example.com")
            cfg.setServerPort(7000)
            cfg.setAuthToken("replace-with-your-token")
            cfg.setLoginFailExit(false)
            cfg.setLogTo("/sdcard/Android/data/your.package.name/files/frpc.log")
            cfg.setLogMaxDays(3)

            val proxy = Frpandroid.newFrpcProxyConfig("stcp")
            proxy.setName("adb-device-001")
            proxy.setSecretKey("replace-with-strong-secret")
            proxy.setLocalIP("127.0.0.1")
            proxy.setLocalPort(5555)
            proxy.setUseEncryption(true)
            proxy.setUseCompression(true)

            cfg.addProxy(proxy)
            Frpandroid.start(cfg)

            //reload
            proxy.setSecretKey("replace-with-strong-secret")
            Frpandroid.reload(cfg)
        } catch (t: Throwable) {
            t.printStackTrace()
        }
    }
}
```

### 3. 常用配置项

`FrpcConfig` 常见配置：

- `setServerAddr(String)`：frps 地址
- `setServerPort(Int)`：frps 端口
- `setAuthToken(String)`：token 鉴权
- `setLoginFailExit(Boolean)`：登录失败后是否直接退出
- `setLogTo(String)`：日志输出路径
- `setLogLevel(String)`：日志级别
- `setLogMaxDays(Long)`：日志保留天数
- `addProxy(FrpcProxyConfig)`：添加代理
- `addVisitor(FrpcVisitorConfig)`：添加访问者

`FrpcProxyConfig` 常见配置：

- `newFrpcProxyConfig("tcp" | "udp" | "http" | "https" | "stcp" | "xtcp" | "sudp")`
- `setName(String)`：代理名称
- `setLocalIP(String)`：本地服务地址
- `setLocalPort(Int)`：本地服务端口
- `setUseEncryption(Boolean)`：是否启用加密
- `setUseCompression(Boolean)`：是否启用压缩
- `setSecretKey(String)`：`stcp` / `xtcp` 等场景使用的密钥

`FrpcVisitorConfig` 常用于内网穿透访问端场景，例如：

- `newFrpcVisitorConfig("stcp")`
- `setName(String)`
- `setServerName(String)`
- `setSecretKey(String)`
- `setBindAddr(String)`
- `setBindPort(Int)`

### 4. 生命周期建议

推荐做法：

- 在 `Service.onCreate()` 或明确的启动入口中调用 `Frpandroid.start(cfg)`
- 在组件销毁时调用 `Frpandroid.stop()`
- 配置变更后使用 `Frpandroid.reload(cfg)` 重载

注意事项：

- 重复调用 `start` 会返回错误，因为内部已做运行状态保护
- `stop` 在未运行时会直接返回，不会抛错
- `reload` 的行为是先 `stop` 再 `start`

## 构建指南

### 环境要求

建议准备以下环境：

- Go
- Android SDK
- `gomobile`
- 可正常编译 `golang.org/x/mobile` 的 Android 工具链

### 安装 gomobile

```powershell
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init
```

### 生成 Android AAR

本仓库当前没有内置自动化脚本，通常直接在仓库根目录执行 `gomobile bind`：

```powershell
gomobile bind -target=android -o frpandroid.aar
```

如果本地工具链配置正常，执行后会生成：

- `frpandroid.aar`
- 可能附带源码包或中间产物，具体取决于本地 gomobile 版本和构建方式

### 构建建议

- 优先先在本仓库中验证 `gomobile bind` 成功，再接入 Android 项目
- Android 侧建议把日志目录放在应用私有目录，避免直接写公共存储
- 生产环境不要把服务地址、token、secret 写死在源码中

## 常见问题

### 1. `start` 调用失败，提示配置错误

优先检查以下项目：

- `ServerAddr` 是否为空
- `ServerPort` 是否大于 0
- 是否至少添加了一个 `proxy` 或 `visitor`
- `newFrpcProxyConfig(...)` / `newFrpcVisitorConfig(...)` 传入的类型字符串是否正确

### 2. 重复启动报错 `frpc is already running`

这是预期行为。当前实现内部做了单例控制，同一时刻只允许一个 frpc 实例运行。

可选处理方式：

- 启动前自行维护状态，避免重复调用
- 配置变化时使用 `Frpandroid.reload(cfg)`
- 在适当时机先调用 `Frpandroid.stop()`

### 3. Android 后台运行一段时间后断开

这通常不是本库独有问题，常见原因包括：

- 系统后台限制
- 设备厂商省电策略
- 前台服务未正确配置
- 网络切换或系统休眠

建议：

- 对长期连接场景使用前台服务
- 根据设备系统版本补齐后台运行权限和通知配置
- 配合日志输出分析断线原因

### 4. 日志文件无法写入

常见原因：

- 日志目录不存在
- 应用无对应目录写权限
- Android 高版本对外部存储访问有限制

建议优先使用应用私有目录，例如：

```text
/sdcard/Android/data/<your.package.name>/files/frpc.log
```

或者直接使用应用内部文件目录。

### 5. `stcp` / `xtcp` 无法连通

优先检查：

- 服务端和访问端的 `secretKey` 是否一致
- `serverName` / `proxy name` 是否匹配
- 本地服务地址和端口是否真实可访问
- frps 服务端是否已正确开启对应能力

### 6. 是否可以直接照搬示例里的设备命名方式

不建议。设备标识、MAC、序列号、真实域名、认证口令都属于敏感信息或高风险信息。

建议改为：

- 使用你自己的设备编号体系
- 使用后端下发的短期凭证
- 避免把唯一设备标识直接暴露在代理名中

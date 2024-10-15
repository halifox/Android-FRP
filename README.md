# Android AAR for frp Client

该项目提供了一个 [fatedier/frp](https://github.com/fatedier/frp) 客户端的 Android AAR 库封装，使您可以在 Android 应用中直接使用 `frp` 功能。

## 使用方法

将 `frp.aar` 集成到您的 Android 项目中：

1. **将 AAR 文件导入项目**：
   将生成的 `frp.aar` 文件放入 Android 项目的 `libs` 目录下。

2. **在 `build.gradle` 文件中添加库依赖**：
   打开项目的 `build.gradle` 文件（通常在 `app` 目录下），在 `dependencies` 块中添加以下内容：
    ```kotlin
    implementation(files("libs/frp.aar"))
    ```

3. **在代码中调用 `frp` 客户端**：
   导入并使用 `Frp` 类中的方法，以便在 Android 应用中调用 `frp` 的功能。例如：
    ```kotlin
    import frp.Frp

    // 启动服务端
    Frp.runServer(configFilePath)

    // 启动客户端
    Frp.runClient(configFilePath)
    ```
   请将 `configFilePath` 替换为 `frp` 配置文件的路径，以正确初始化 `frp`。

以上步骤完成后，您便可以在 Android 应用中使用 `frp` 提供的功能。

## 构建指南

按以下步骤构建 AAR 库：

1. **安装 `gomobile`**：
    ```bash
    go install golang.org/x/mobile/cmd/gomobile@latest
    ```
   该命令会安装 `gomobile` 工具到 `$GOPATH/bin` 中。

2. **初始化 `gomobile`**：
    ```bash
    gomobile init
    ```
   此步骤下载并配置 Android 和 iOS 的 SDK 依赖。如果您没有配置 Android SDK 和 NDK，请按照之前的说明进行操作。

3. **获取 `gomobile bind` 包**：
    ```bash
    go get golang.org/x/mobile/bind
    ```
   该命令获取 `bind` 包，以便 `gomobile` 能够将 Go 包绑定为移动平台库。

4. **构建 AAR**：
    ```bash
    gomobile bind -androidapi 21 .
    ```
    - 此命令将在当前目录下生成 AAR 文件（包含 `frp` 客户端的 Android 库）。
    - 如果您希望生成 iOS Framework，可以使用 `-target=ios` 参数。
    - 构建完成后，目录下会生成：
        - `frp.aar`：包含 frp 客户端的 Android AAR 库。
        - `frp-sources.jar`：源代码 JAR，用于参考。

## gomobile 简介

`gomobile` 是 Go 语言的一个移动开发工具集，旨在帮助开发者将 Go 编写的代码编译为 Android 和 iOS 可用的库。这允许开发者在移动应用中使用 Go 实现的功能。`gomobile` 包含两个核心命令：`gomobile init` 和
`gomobile bind`。

### 基本功能

- **跨平台支持**：允许开发者在 Android 和 iOS 上复用用 Go 编写的代码。
- **AAR 和 Framework 打包**：将 Go 包编译为 Android 的 AAR 和 iOS 的 Framework，使它们在各自平台上作为原生库使用。
- **API 导出**：将 Go 包中的公共 API 自动转换为 Java（Android）或 Objective-C（iOS）接口，从而在移动端应用中调用 Go 代码。

### 安装和初始化

1. **安装 `gomobile`**：
   首先，确保您安装了最新版的 Go。运行以下命令安装 `gomobile`：
    ```bash
    go install golang.org/x/mobile/cmd/gomobile@latest
    ```
   `gomobile` 将自动下载并安装到 `$GOPATH/bin` 目录下。

2. **初始化 `gomobile`**：
   `gomobile init` 命令会安装所需的 Android 和 iOS 工具链，并配置环境，以便 `gomobile` 使用。在命令行中运行：
    ```bash
    gomobile init
    ```
   **注意**：确保您的系统安装了 Android SDK 和 Android NDK，尤其是在 Android 平台上构建时。

    - **Android SDK**：请从 [Android 开发者网站](https://developer.android.com/studio)下载安装，并确保 `sdkmanager` 和 `ndk-bundle` 路径已添加到环境变量中。
    - **Android NDK**：Android NDK 必须版本不低于 r19。

3. **获取 `gomobile bind` 包**：
   `gomobile bind` 是 `gomobile` 工具集中的一个命令，用于生成 Android AAR 和 iOS Framework。运行以下命令安装该包：
    ```bash
    go get golang.org/x/mobile/bind
    ```
   安装完成后，您可以使用 `gomobile bind` 将 Go 包编译为 Android AAR 或 iOS Framework。

### gomobile bind 命令

`gomobile bind` 将 Go 包中的公共方法编译为移动平台可用的库。以下是常用的参数说明：

- `-target`：指定目标平台，如 `android`、`ios` 或 `android/ios`。
- `-androidapi`：设置 Android 的最低 API 级别（默认为 15）。例如，`-androidapi 21` 表示构建的库需要 Android 5.0 (API 21) 以上的设备。
- `-javapkg`：自定义生成的 Java 包名，默认与 Go 包名称一致。

示例：

```bash
gomobile bind -target=android -androidapi 21 .
```

该命令会将当前目录中的 Go 包编译为 Android 库，并设置最低 API 级别为 21。

### 常见问题

- **无法初始化**：如果 `gomobile init` 失败，确保您的系统已安装 Go 和 Java SDK（用于 Android 开发）以及 Android NDK。
- **Go 版本不支持**：`gomobile` 可能要求 Go 的较新版本。建议始终使用最新的 Go 和 `gomobile` 版本。
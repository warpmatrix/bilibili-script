# Cookie

> 如果使用 github actions 方式运行，请勿将 cookie 存储到该文件夹中或者其他公开存储的方式，避免造成个人信息的泄漏

同一个镜像可以建立多个容器，因此该项目可以支持多个用户运行脚本。只需要在文件夹中添加对应用户的 cookie 信息，并在 `docker-compose.yml` 进行相应的设置即可。

> b 站的 cookie 是长期有效的，一台主机上的浏览器可以存储一个用户的 cookie 信息，只要在该浏览器不进行登出操作 cookie 不会失效。

需要的 cookie 字段有：`SESSDATA`、`bili_jct`、`DedeUserID`。主要的操作步骤如下：

1. 对于一个用户，建立一个文件记录其 cookie 信息，如：`touch user.cookie`

2. 在创建的文件中以 `key=value` 的形式添加对应的 cookie 信息，示例如下：

    ```plaintext
    SESSDATA=sessdata_val
    BILI_JCT=bilijct_val
    DEDEUSERID=dedeuserid_val
    ```

3. 在 `docker-compose.yml` 的 `services` 条目下添加对应的用户信息，如：

    ```yml
    user:
      build: .
      env_file: cookie/user.cookie
    ```

# Cookie

同一个镜像可以建立多个容器，因此该项目可以支持多个用户运行脚本。只需要在文件夹中添加对应用户的 cookie 信息，并在 `docker-compose.yml` 进行相应的设置即可。

> b 站的 cookie 是长期有效的，一台主机上的浏览器可以存储一个用户的 cookie 信息，只要在该浏览器不进行登出操作 cookie 不会失效。

需要的 cookie 字段有：`SESSDATA`、`bili_jct`、`DedeUserID`。主要的操作步骤如下：

1. 对于一个用户，建立一个文件记录其 cookie 信息，如：`touch user.cookie`
2. 在创建的文件中以 `key=value` 的形式添加对应的 cookie 信息，示例如下：

    ```plaintext
    SESSDATA=sessdata_val
    bili_jct=bilijct_val
    DedeUserID=dedeuserid_val
    ```

3. 在 `docker-compose.yml` 的 `services` 条目下添加对应的用户信息，如：

    ```yml
    user:
      image: bilibili-script
      env_file: cookie/user.cookie
    ```

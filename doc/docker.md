# Run Script by Docker

## Quick Start

1. 在 cookie 文件夹中创建存储 cookie 信息的文件，如：`touch user.cookie`
2. 将 bili_jct、DedeUserID、SESSDATA 对应的 cookie 信息，以 `key=value`、键名全部大写的形式存储在新创建的文件中。
3. 执行 `make` 命令，就会在服务器上构建对应的容器。默认启动进程为 bash，通过 `docker attach` 即可进入容器，查看执行情况。

## Advanced Setting

- 自定义配置文件：在 config 文件夹中可以找到配置文件的模板，根据其中的内容修改并创建新的配置文件。然后，在 [docker-compose.yml](../docker-compose.yml) 中，参考其中被注释的 `services.user2`，在 `args` 字段中添加 `cfgFile` 字段，其值设为配置文件的路径。
- 脚本的运行时间：在 [docker-compose.yml](../docker-compose.yml) 的 `services.<service-id>.build.args.cron` 中可以设置脚本的运行时间。默认时间为 UTC 的 0:20，即北京时间的 8:20，修改其中的值即可修改脚本运行时间，具体的时间格式为：min hour day mon week。具体添加的字段可以参考 [docker-compose.yml](../docker-compose.yml) 中被注释的 `services.user2` 的 `args.cron` 字段。
- 多账户的支持：要想多账户运行脚本可以在 [docker-compose.yml](../docker-compose.yml) 中添加新的 services，并且在 `env_file` 中设置为对应用户的 cookie 文件路径。可以参考 [docker-compose.yml](../docker-compose.yml) 中被注释的 `services.user2` 作为模板，`build.args` 中传递的参数。

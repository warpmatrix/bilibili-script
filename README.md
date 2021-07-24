# Bilibili Script

一个简单的 bilibili 脚本，只需对应的 cookie 信息，就可以自动完成一些基本的日常操作如自动签到。具体需要的 cookie 信息为：bili_jct、DedeUserID、SESSDATA，查看 cookie 信息的方式可以自行搜索。用户可以提供配置文件对脚本的行为进行设置，项目根据具体的部署方式，支持对多个账户进行操作，并且可以根据不同的账户设置不同的配置文件。目前项目维护两种部署方式：

- [使用 github 提供的 github actions](doc/github-action.md)
- [在一台服务器上使用 docker 创建容器实现环境隔离并自动运行](doc/docker.md)

<!-- TODO: 更多部署方式如阿里云、腾讯云等其他 serverless 服务 -->

## Advanced Settings

用户可以自定义**脚本行为**，在 `config` 文件夹中提供了具体的模板。用户创建新的 yaml 文件后，并且根据具体的部署方式重新映射配置文件即可。同时，用户也可以自己配置脚本**运行时间**以及**多账户**脚本执行。具体的配置方式参见 [github action](doc/github-action.md) 和 [docker](doc/docker.md) 的部署文档。关于配置文件内容的详细说明可以参照配置文件中的注释。

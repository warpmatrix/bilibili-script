# Todo

- [x] 将 main 的执行文件放入 docker 中运行
- [x] 添加环境变量
- [x] 完成基本的 get 操作
- [x] 在 docker 中设置 client，进行伪装
- [x] 设置 cookie
- [x] 引入 json，转换得到的数据
- [x] 完成 domain model
- [x] 引入用户自定义的 config (yaml)
- [x] 自定义 logger 完成日志的记录，还是考虑引入第三方库？
- [x] logger 的测试和结构构思
- [x] 完成 fetcher 的其他方法，重构整理 client 结构
- [x] 引入等待机制，模拟浏览器行为
- [x] `todo` 和 `.gitignore` 的迁移，实现 master 和 dev 的分离
- [x] docker 的 cron
- [x] 分离 dockerfile 和 docker-compose（docker 对多行变量的支持？）
- [x] 多账户的支持，通过 docker-compose 进行具体配置
- [x] 重构 status.code 部分，部分请求可能以 400 状态码的形式返回结构
- [x] github action 对多用户的支持，实现 user_cookie 环境变量
- [ ] 完成日常任务：manga、live、daily
- [ ] 默认配置的支持
- [ ] 使用 docker secrect 管理 cookie 作为系统变量进行传递
- [ ] log 的持久化
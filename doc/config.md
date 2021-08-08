# Configuration Specification

> 所有任务都可以通过直接删除或注释相关配置剔除任务。所有的任务均为可选运行，并且提供默认参数，方便的同时留给用户最大的定制空间。

视频相关任务 (video)：

| 字段名 | 字段内容 | 示例 | 说明 |
| :-: | :-: | :-: | :-: |
| daily | 由 watch 和 share 组成的序列 | `daily: [watch, share]` | 观看和分享均为模拟网页端的行为，不会进行真正的分享
| coin | 一个整数或者设置 `num` 和 `like` 字段 | `coin: 2` | `coin` 的整数或 `num` 表示使用的硬币数量，只能是 0-5 的一个整数；`like` 表示投币的同时是否点赞。

漫画相关任务 (manga)：直接填写漫画的平台（android 或 ios），示例：`manga: android`

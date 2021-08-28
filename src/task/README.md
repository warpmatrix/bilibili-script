# Task in Bilibili

可以通过直接增加删除 go 文件修改脚本可以执行的任务。`task.go` 中提供了任务的模板 `task` 结构体，其余文件定义具体的实例和方法实现模板。

其中，`task` 结构体的 `name`、`impl`、`result` 为必用字段，必须在具体的实现文件中定义。

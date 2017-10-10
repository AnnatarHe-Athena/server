## server

对 webapp, Android, iOS, Mac 客户端的 api server，提供数据支持。

采用 GraphQL 结构，数据库使用 PostgresQL.

## usage

1. 请首先 clone 此仓库。

2. 若在 Unix 环境中，请安装相应的 make tool 以及 golang，然后执行 `make deps`

3. 开发环境请首先安装 docker, 并 clone [server-dockerfile](https://github.com/AnnatarHe-Athena/server-dockerfile) 项目，然后根据其中的 README.md 修改对应的目录文件。并按照其中的指引启用开发环境和部署环境

4. 生产环境打包：请**务必**在 Unix 环境中执行 `make build`， 然后进入 `/tmp` 目录上传至服务器。再由服务器解压缩，然后交给 server-dockerfile 这个项目中的生产环境启动即可。

## 其他

关于为何在生产环境打包一定要在 Unix 环境中进行是有原因的。因为 Windows 和 Unix 对于目录的分隔符处理不一致，导致在 windows 下打包产生的目录并不是 Unix 部署时候所期望的目录。会形成类似于 `src\\hello\\world\\main.go` 这样的文件，而不是层级结构。


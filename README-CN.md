# Solo CI

[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/solo-ci/Lobby)

### Description

一个轻量级的Golang CI/CD工具，全自动clone代码，构建，部署，只需要几行配置即可

```json
{
  "get_list":[
    "github.com/asaskevich/govalidator"
  ],
  "zip_list":[
    "conf"
  ],
  "after_script":"echo hello",
  "before_script":"pwd"
}
```

### Features

- 完美集成Gitlab，Github（application/json）
- 配置要求远低于主流CI工具（Jenkins etc.）内存占用低，可以运行在任何配置Linux主机中
- 一键开启，只需要Golang环境和Git环境，程序会自动获取自己所需要的环境
- 配制简单，只有四个配置项
- 一键clone，build，test，打包成tar
- 支持自定义脚本，构建前构建后触发均可自定义，脚本执行目录为当前构建目录
- REST API支持，可以集成进任何系统
- 可以保存任意数量的构建，不丢任何构建

### Use

1. go get github.com/astaxie/beego  go get github.com/mattn/go-sqlite3  go get github.com/satori/go.uuid
2. 配置好主机的GOPATH，GOROOT，GIT环境
3. 下载编译好的二进制
4. 使用REST API新建项目
5. 在你的项目中写个简单的solo.json，并且在代码管理中配置webhook (配置地址请看REST API)
6. push！触发CI
7. 构建好的程序会存在workspace文件夹中
8. go get github.com/mattn/go-sqlite3需要预安装gcc, 必要时使用"sudo apt-get install build-essential"命令安装

### REST API

| Method | Url                                      | Params                                   | Description |
| :----: | :--------------------------------------- | :--------------------------------------- | :---------- |
|  POST  | http://your-ip:13233/v1/solohook/:project_id | - project_id(path)                       | 触发Webhook   |
|  POST  | http://your-ip:13233/v1/project          | - name(form)                             | 创建项目        |
|        |                                          | - type(form, gitlab or github or bitbucket) |             |
|        |                                          | - url(form)                              |             |
|        |                                          | - path(form，the position of solo-ci.json) |             |
|        |                                          | - branch(form)                           |             |
|        |                                          | - secret_token(form,not necessary)       |             |
|        |                                          | - main_path(form,the position of main.go) |             |
| DELETE | http://your-ip:13233/v1/project/:project_id | - project_id(path)                       | 删除项目        |
|  PUT   | http://your-ip:13233/v1/project/:project_id | - project_id(path)                       | 更新项目        |
|        |                                          | - name(form)                             |             |
|        |                                          | - type(form, gitlab or github) |             |
|        |                                          | - url(form)                              |             |
|        |                                          | - path(form，the position of solo-ci.json) |             |
|        |                                          | - branch(form)                           |             |
|        |                                          | - secret_token(form,not necessary)       |             |
|  GET   | http://your-ip:13233/v1/project/:project_id | - project_id(path)                       | 获取项目信息      |
|  GET   | http://your-ip:13233/v1/project          | - project_id(path)                       | 获取项目列表      |
|        |                                          | - page (default 0)                       |             |
|        |                                          | - pageSize(default 20)                   |             |

### 配置文件solo-ci.json

- get_list：需要下载的Go包
- zip_list：构建完成需要打包进项目的文件或者目录
- before_script：构建之前执行的脚本
- after_script：构建之后执行的脚本

所有的选项都不是必须存在的，及时你什么都不写也可以，下面是一个空的配置文件例子

```json
{

}
```

### Next

- Web GUI 支持

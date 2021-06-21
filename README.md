#整体框架
```
本服务通过gin框架来实现主要功能:
(1)启动时读取app.ini 配置文件，监听配置中的http端口号，用户请求过来时，返回有用的士兵配置。
(2)启动时用github.com/spf13/pflag 来解析命令行参数传入的 config.army.model.json文件路径，
整理格式且只保留有用的数据，生成新的json文件。启动时解析该文件。
```

#目录结构
```
├── config                              #操作配置文件
│   ├── ini.go
│   ├── init.go
│   └── jsonConfig.go
├── configFile                          #配置文件文件夹
│   ├── app.ini
│   └── config.army.model.json
├── ctrl                                #处理router的请求
│   └── routerCtrl.go
├── handle                              #处理具体业务逻辑
│   ├── handle.go
├── data.json                           #生成的json文件
├── go.mod
├── main.go                             #代码入口
├── pressureTest                        #压力测试文件夹
│   ├── pressure.py
│   └── report.html
├── router                              #路由文件夹
│   └── router.go
├── service                             #处理通用逻辑
│   └── jsonDeal.go
├── soldierInfo                         #士兵数据结构
│   └── soldierInfo.go
├── soldier_test.go                     #测试代码
└── util                                #通用工具方法
    └── jsonRW.go
```

#代码逻辑分层  soldier
| 层     | 文件夹|主要职责 |调用关系|
| :----: | :----|:---- | :-----|
|router  | /router|路由转发 |调用ctrl|
|ctrl   | /ctrl|处理路由 |调用handle |
|handle   | /handle|处理具体业务逻辑 |调用service module|
|service | /service|处理通用逻辑|调用module|
|module  | /soldierInfo|数据模型 |被handles service调用|

#存储设计
| 内容 | 类型 |  说明 |
| :----:| :---- | :----|
|Id  | string | 存储ID数据|
|Name |string |存储士兵名字数据|
|UnlockArena| string |士兵解锁度|
|Rarity   |  string |士兵稀有度|
|CombatPoints| string |士兵战力|
|Desc     |    string |士兵阶段|


#接口设计

##1.输入士兵id获取稀有度
##接口地址
```
/idGetRarity
```
## 请求方式
GET
## 请求示例
```
http://127.0.0.1:8000/idGetRarity?id=10101
```
## 参数  说明

``` 
id  string类型 输入需要查询的士兵id
```

```
成功示例 
{
    "Rarity": "1",
    "condition": "success",
    "id": "10101"
}
#错误示列 
{
    "Rarity": "error",
    "condition": "error",
    "id": ""
}
```

##2.输入士兵id获取战力
```
接口地址 
/idGetCombatPoints
```
## 请求方式
GET
## 请求示例
```
http://127.0.0.1:8000/idGetCombatPoints?id=10101
```
## 参数  说明
``` 
id  string类型 输入需要查询的士兵id
```
```
成功示例 
{
    "CombatPoints": "167",
    "condition": "success",
    "id": "10101"
}
错误示列 
{
    "CombatPoints": "error",
    "condition": "not_exist",
    "id": "101"
}
```
##3.获取每个阶段解锁相应士兵的json数据

##接口地址
```
/getSoldiersJson
```
## 请求方式
GET
## 请求示例
```
http://127.0.0.1:8000/getSoldiersJson?stage=army_desc_10101
```
## 参数  说明
``` 
id  string类型 输入需要查询的阶段
```
```
成功示例
{
    "condition": "success",
    "data": [
        {
            "id": "10102",
            "name": "Swordsman",
            "UnlockArena": "0",
            "rarity": "1",
            "combatPoints": "342",
            "desc": "army_desc_10101"
        },
        {
            "id": "10103",
            "name": "Swordsman",
            "UnlockArena": "0",
            "rarity": "1",
            "combatPoints": "691",
            "desc": "army_desc_10101"
        }
    ],
    "stage": "army_desc_10101"
}
错误示例
{
    "condition": "not_exist",
    "data": "",
    "stage": "army_desc_"
}
```

#第三方库
```
github.com/spf13/pflag      从命令行获取需要的参数
github.com/gin-gonic/gin     gin开发框架
github.com/clod-moon/goconf  解析配置文件
```

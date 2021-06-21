#整体框架
本服务通过gin框架来实现主要功能

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
│   ├── handle.go
│   └── routerCtrl.go
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
|ctrl   | /ctrl|处理路由 |调用service module|
|service | /service|处理通用逻辑| 被ctrl调用 调用module|
|module  | /soldierInfo|数据模型 |被ctrl service调用|

#存储设计
| 内容 | 类型 |  说明 |
| :----:| :---- | :----|
|Id  | string | 存储ID数据|
|Name |string |存储士兵名字数据|
|UnlockArena| string |士兵解锁度|
|Rarity   |  string |士兵稀有度|
|CombatPoints| string |士兵战力|
|Desc     |    string |士兵阶段|





# ecod

## 代码生成工具

### 概述

```
此项目为go的db操作代码生成工具，目前使用db引擎为go原生database/sql。程序自动生成结构体和对应的基本操作函数。便于项目移植。
```


### 说明

```
可执行程序在bin/下。(可自行将可执行程序添加环境变量中）
ecod (linux)
ecod.exe (win)
ecod.app (mac)
```



### 安装方法（linux环境)

```
下拉：
    git clone https://github.com/et-zone/ecod.git

程序添加环境遍历
	sudo cp ecod/bin/ecod /usr/local/bin   
	·
```

### 使用方法


```
 如linux环境下，配置好配置json文件，
 执行命令 ecod -n xx.json
```

### 配置文件案例


```

{
    "tname": "teach", -- 表名
    "sName": "teacher", -- 结构体名称
    "data": [
        {
            "fname": "id", -- 字段db名称
            "ftype": "int64" -- 字段go类型
        },
        {
            "fname": "name",
            "ftype": "string"
        },
        {
            "fname": "v_fname",
            "ftype": "string"
        }
    ]
}
```



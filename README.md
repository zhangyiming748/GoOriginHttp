# 查询天气接口文档

---
title: go原生http v1.0.0
language_tabs:
- shell: Shell
- http: HTTP
- javascript: JavaScript
- ruby: Ruby
- python: Python
- php: PHP
- java: Java
- go: Go
  toc_footers: []
  includes: []
  search: true
  code_clipboard: true
  highlight_theme: darkula
  headingLevel: 2
  generator: "@tarslib/widdershins v4.0.17"

---

# go原生http

> v1.0.0

Base URLs:

# Default

## GET 获取预报天气

GET /v1/GetWeather

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|City|query|string| 否 |城市的中文名|
|extensions|query|string| 否 |none|

> 返回示例

> 成功

```json
{
  "resCode": "200",
  "resMessage": "success",
  "retData": "{\"status\":\"1\",\"count\":\"1\",\"info\":\"OK\",\"infocode\":\"10000\",\"forecasts\":[{\"city\":\"石景山区\",\"adcode\":\"110107\",\"province\":\"北京\",\"reporttime\":\"2023-05-06 10:09:13\",\"casts\":[{\"date\":\"2023-05-06\",\"week\":\"6\",\"dayweather\":\"多云\",\"nightweather\":\"晴\",\"daytemp\":\"23\",\"nighttemp\":\"11\",\"daywind\":\"北\",\"nightwind\":\"北\",\"daypower\":\"≤3\",\"nightpower\":\"≤3\",\"daytemp_float\":\"23.0\",\"nighttemp_float\":\"11.0\"},{\"date\":\"2023-05-07\",\"week\":\"7\",\"dayweather\":\"晴\",\"nightweather\":\"晴\",\"daytemp\":\"26\",\"nighttemp\":\"11\",\"daywind\":\"北\",\"nightwind\":\"北\",\"daypower\":\"≤3\",\"nightpower\":\"≤3\",\"daytemp_float\":\"26.0\",\"nighttemp_float\":\"11.0\"},{\"date\":\"2023-05-08\",\"week\":\"1\",\"dayweather\":\"晴\",\"nightweather\":\"多云\",\"daytemp\":\"27\",\"nighttemp\":\"13\",\"daywind\":\"南\",\"nightwind\":\"南\",\"daypower\":\"≤3\",\"nightpower\":\"≤3\",\"daytemp_float\":\"27.0\",\"nighttemp_float\":\"13.0\"},{\"date\":\"2023-05-09\",\"week\":\"2\",\"dayweather\":\"多云\",\"nightweather\":\"多云\",\"daytemp\":\"27\",\"nighttemp\":\"15\",\"daywind\":\"南\",\"nightwind\":\"南\",\"daypower\":\"≤3\",\"nightpower\":\"≤3\",\"daytemp_float\":\"27.0\",\"nighttemp_float\":\"15.0\"}]}]}"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

## GET 删除全部实时天气

GET /v1/DeleteAllLive

> 返回示例

> 成功

```json
{
  "resCode": "200",
  "resMessage": "success",
  "retData": null
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

## GET 获取城市代码列表

GET /v1/GetCity

> 返回示例

> 成功

```json
{
  "resCode": "200",
  "resMessage": "成功",
  "retData": {
    "琼海市": "469002",
    "大庆市": "230600",
    "大庆市市辖区": "230601",
    "石景山区": "110107",
    "东城区": "110101",
    "北京市": "110000",
    "北京市市辖区": "110100",
    "海淀区": "110108"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|


# 需要放在根目录的配置文件内容如下

```ini
[redis]
connect = 127.0.0.1:6379
pass = 888888
db = 0
[mysql]
user = zen
passwd = 123456
database = mydb
ip = 127.0.0.1
port = 33060
[log]
level = Debug
[weather]
key = 高德开放平台32位数字
```
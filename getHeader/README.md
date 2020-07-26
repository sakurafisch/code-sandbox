# 获取 get 请求头的参数

## 启动服务器

```bash
go run main.go
```

## 发送 get 请求

```bash
curl "http://127.0.0.1:8090/param?name=winnerwinter&age=18"
```

返回数据如下：

```
map[Accept:[*/*] User-Agent:[curl/7.55.1]]
map[age:[18] name:[winnerwinter]]
[winnerwinter]
winnerwinter
```
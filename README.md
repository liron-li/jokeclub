# jokeclub

基于 `gin` 的go web 框架

# 安装

1.

```
dep ensure -v
```

2.

```
go install
```

3. 编译

```
go build -ldflags "-s -w"
```

4. 生成api文档

```
   apidoc -i app/controllers/api -o apidoc/
```

5.交叉编译
```
CGO_ENABLED=0 GOOS=linux go build  -a -installsuffix cgo -o jokeclub .
```
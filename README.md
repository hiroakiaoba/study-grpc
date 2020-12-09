# study-grpc

grpcのべんきょ

TODOリストサービスを作成する。

## データ

- user
  - id
  - name
  - age
- project
  - id
  - title
  - users
- todo
  - id
  - content
  - status: waiting | doing | done
  - createdAt
  - (projectId)

## コマンド

protoからコード生成

```shell
protoc \
  -Iproto \
  --go_out=plugins=grpc:api \ proto/*.proto
```

grpcurl

```shell
grpcurl -plaintext localhost:50051 todoService.TodoService/GetAll
```


## 参考

https://github.com/gami/grpc_example

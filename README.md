# study-grpc

grpcのべんきょ

TODOリストサービスを作成する。

## データ

- user
  - id
  - login_name
  - password
- project
  - id
  - title
  - users
  - userId
  - created_at
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

```shell
grpcurl -d '{"loginName":"hoge", "password":"password"}' -plaintext localhost:50051 todoService.TodoService/GetAll
```

```shell
grpcurl -d '{"loginName":"hoge", "password":"password"}' -plaintext localhost:50051 user.UserService/SignUp
```


## 参考

https://github.com/gami/grpc_example

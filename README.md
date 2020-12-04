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

```shell
protoc \
  -Iproto \
  --go_out=plugins=grpc:api \ proto/*.proto
```

## 参考

https://github.com/gami/grpc_example

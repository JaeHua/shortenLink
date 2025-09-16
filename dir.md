shortenLink
│  go.mod
│  go.sum
│  README
│
│
├─api
│  ├─convert
│  │      convert.api
│  │
│  └─show
│          show.api
│
├─cmd
│  ├─convert-api
│  │  │  convert.go
│  │  │
│  │  ├─etc
│  │  │      convert-api.yaml
│  │  │
│  │  ├─internal
│  │  │  ├─config
│  │  │  │      config.go
│  │  │  │
│  │  │  ├─errorx
│  │  │  │      baseerror.go
│  │  │  │
│  │  │  ├─handler
│  │  │  │      converthandler.go
│  │  │  │      routes.go
│  │  │  │
│  │  │  ├─logic
│  │  │  │      convertlogic.go
│  │  │  │
│  │  │  ├─svc
│  │  │  │      servicecontext.go
│  │  │  │
│  │  │  └─types
│  │  │          types.go
│  │  │
│  │  └─logs
│  │      └─convert-api
│  │              access.log
│  │              error.log
│  │              severe.log
│  │              slow.log
│  │              stat.log
│  │
│  ├─convert-rpc
│  │  │  convert.go
│  │  │
│  │  ├─convertclient
│  │  │      convert.go
│  │  │
│  │  ├─etc
│  │  │      convert.yaml
│  │  │
│  │  ├─internal
│  │  │  ├─config
│  │  │  │      config.go
│  │  │  │
│  │  │  ├─logic
│  │  │  │      convertlogic.go
│  │  │  │
│  │  │  ├─server
│  │  │  │      convertserver.go
│  │  │  │
│  │  │  └─svc
│  │  │          servicecontext.go
│  │  │
│  │  ├─logs
│  │  │  └─convert-rpc
│  │  │          access.log
│  │  │          error.log
│  │  │          severe.log
│  │  │          slow.log
│  │  │          stat.log
│  │  │
│  │  └─shortenLink
│  │      └─rpc
│  │          └─convert
│  │                  convert.pb.go
│  │                  convert_grpc.pb.go
│  │
│  ├─sequence-rpc
│  │  │  sequence.go
│  │  │
│  │  ├─etc
│  │  │      sequence.yaml
│  │  │
│  │  ├─internal
│  │  │  ├─config
│  │  │  │      config.go
│  │  │  │
│  │  │  ├─logic
│  │  │  │      nextlogic.go
│  │  │  │
│  │  │  ├─server
│  │  │  │      sequenceserver.go
│  │  │  │
│  │  │  └─svc
│  │  │          servicecontext.go
│  │  │
│  │  ├─logs
│  │  │  └─sequence-rpc
│  │  │          access.log
│  │  │          error.log
│  │  │          severe.log
│  │  │          slow.log
│  │  │          stat.log
│  │  │
│  │  ├─sequenceclient
│  │  │      sequence.go
│  │  │
│  │  └─shortenLink
│  │      └─rpc
│  │          └─sequence
│  │                  sequence.pb.go
│  │                  sequence_grpc.pb.go
│  │
│  ├─show-api
│  │  │  show.go
│  │  │
│  │  ├─etc
│  │  │      show-api.yaml
│  │  │
│  │  ├─internal
│  │  │  ├─config
│  │  │  │      config.go
│  │  │  │
│  │  │  ├─errorx
│  │  │  │      baseerror.go
│  │  │  │
│  │  │  ├─handler
│  │  │  │      routes.go
│  │  │  │      showhandler.go
│  │  │  │
│  │  │  ├─logic
│  │  │  │      showlogic.go
│  │  │  │
│  │  │  ├─svc
│  │  │  │      servicecontext.go
│  │  │  │
│  │  │  └─types
│  │  │          types.go
│  │  │
│  │  └─logs
│  │      └─show-api
│  │              access.log
│  │              error.log
│  │              severe.log
│  │              slow.log
│  │              stat.log
│  │
│  └─show-rpc
│      │  show.go
│      │
│      ├─etc
│      │      show.yaml
│      │
│      ├─internal
│      │  ├─config
│      │  │      config.go
│      │  │
│      │  ├─logic
│      │  │      showlogic.go
│      │  │
│      │  ├─server
│      │  │      showserver.go
│      │  │
│      │  └─svc
│      │          servicecontext.go
│      │
│      ├─logs
│      │  └─show-rpc
│      │          access.log
│      │          error.log
│      │          severe.log
│      │          slow.log
│      │          stat.log
│      │
│      ├─shortenLink
│      │  └─rpc
│      │      └─show
│      │              show.pb.go
│      │              show_grpc.pb.go
│      │
│      └─showclient
│              show.go
│
├─db
│  └─migrations
│          sequence.sql
│          short_url_map.sql
│
├─model
│      sequencemodel.go
│      sequencemodel_gen.go
│      shorturlmapmodel.go
│      shorturlmapmodel_gen.go
│      vars.go
│
├─pkg
│  ├─base62
│  │      base62.go
│  │      base62_test.go
│  │
│  ├─connect
│  │      connect.go
│  │      connect_test.go
│  │
│  ├─md5
│  │      md5.go
│  │      md5_test.go
│  │
│  └─urltool
│          urltool.go
│          urltool_test.go
│
└─rpc
    ├─convert
    │      convert.proto
    │
    ├─sequence
    │      sequence.proto
    │
    └─show
            show.proto
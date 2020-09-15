module github.com/airingone/air-grpc

go 1.13

require (
	github.com/airingone/air-etcd v1.0.2
	github.com/airingone/config v1.0.7
	github.com/airingone/log v0.0.0-20200831072359-6cec598b97cf
	github.com/airingone/pro_proto v1.0.1
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/spf13/afero v1.3.5 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sys v0.0.0-20200909081042-eff7692f9009 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20200806022845-90696ccdc692 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.61.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)

replace google.golang.org/grpc v1.32.0 => google.golang.org/grpc v1.26.0

replace google.golang.org/api v0.15.1 => google.golang.org/api v0.14.0

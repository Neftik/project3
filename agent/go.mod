module github.com/Neftik/project3/agent

go 1.23.1

require (
	github.com/fatih/color v1.16.0
	github.com/neandrson/protos v0.0.0-20250502105814-f8a135693888
	google.golang.org/grpc v1.63.2
)

require (
	github.com/Neftik/project3/shunting-yard v0.0.0-20250508102915-132fd6610a7d
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/Neftik/project3/shunting-yard => ../shunting-yard

require github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/protobuf v1.36.5
)

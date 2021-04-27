module github.com/labiraus/gomud-api

go 1.16

require (
	github.com/labiraus/gomud-api/api v0.0.0-00010101000000-000000000000
	github.com/labiraus/gomud-common v0.0.4
	golang.org/x/net v0.0.0-20210423184538-5f58ad60dda6 // indirect
	golang.org/x/sys v0.0.0-20210426230700-d19ff857e887 // indirect
	google.golang.org/genproto v0.0.0-20210426193834-eac7f76ac494 // indirect
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/labiraus/gomud-api/api => ./api

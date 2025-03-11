module github.com/vishalanarase/guestbook/services/auth

go 1.24.1

replace github.com/vishalanarase/guestbook/clients/auth => ../../clients/auth

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3
	github.com/vishalanarase/guestbook/clients/auth v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.33.0
	google.golang.org/grpc v1.71.0
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.5-20250307204501-0409229c3780.1 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250303144028-a0af3efb3deb // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250303144028-a0af3efb3deb // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

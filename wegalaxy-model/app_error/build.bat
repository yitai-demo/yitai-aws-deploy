@del /f /s /q stub\*.go >nul 2>nul
set GOOGLEAPI_PATH=/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis
set DESCRIPTOR_PATH=/include
protoc -I. -I%GOPATH%%DESCRIPTOR_PATH%  -I%GOPATH%%GOOGLEAPI_PATH% ^
     --go_out=./  --go_opt=paths=source_relative ^
     --go-grpc_out=./ --go-grpc_opt=paths=source_relative ^
     *.proto
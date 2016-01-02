.PHONY:: test beat-arm beat-amd64 build

GO_BUILD_FLAGS=-v -tags netgo -installsuffix netgo -ldflags '-s -w'

test::
	@go test ./...

beat-arm::
	@GOOS=linux GOARCH=arm go build $(GO_BUILD_FLAGS) -o dist/beat_arm beat/main.go

beat-amd64::
	@GOARCH=amd64 go build $(GO_BUILD_FLAGS) -o dist/beat_amd64 beat/main.go

build::
	@go build $(GO_BUILD_FLAGS) -o dist/heartbeat main.go
	@strip -s dist/heartbeat

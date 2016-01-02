.PHONY:: build-beat-arm build-beat-amd64

GO_BUILD_FLAGS=-v -tags netgo -installsuffix netgo -ldflags '-s -w'

build-beat-arm::
	@GOOS=linux GOARCH=arm go build $(GO_BUILD_FLAGS) -o dist/beat_arm beat/main.go

build-beat-amd64::
	@GOOS=linux GOARCH=amd64 go build $(GO_BUILD_FLAGS) -o dist/beat_amd64 beat/main.go

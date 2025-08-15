.PHONY: build help

# 根据不同平台编译生成不同格式的二进制包
ifeq ($(shell uname),Darwin)
 OS_PLATFORM="darwin"
else
 ifeq ($(OS),Windows_NT)
  OS_PLATFORM="windows"
 else
  OS_PLATFORM="linux"
 endif
endif

BINARY="rag_interview"
PLATFORM ?= ${OS_PLATFORM}

help:
	@echo "make build - 编译 Go 代码, 生成二进制文件"

build:
	@echo "PLATFORM:${PLATFORM}, GOARCH:amd64"
	CGO_ENABLED=0 GOOS=${PLATFORM} GOARCH=amd64 go build -o ${BINARY}

build-test:
	@echo "PLATFORM:${PLATFORM}, GOARCH:amd64"
	CGO_ENABLED=0 GOOS=${PLATFORM} GOARCH=amd64 go build -gcflags=all="-N -l" -o ${BINARY}

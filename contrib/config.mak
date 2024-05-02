ifeq ($(TARGET_GOOS),windows)
GOLDFLAGS+=-H=windowsgui
endif
export CGO_ENABLED=0

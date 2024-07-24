SOURCES=Makefile main.go main_release.go main_debug.go config.go config_release.go config_template.go 
GARBLE_BIN = $(shell go env GOPATH)/bin/garble

all: socks5-ssh-proxy socks5-ssh-proxy.release
test: socks5-ssh-proxy
	cp socks5-ssh-proxy ~/.ssh; cd ~/.ssh; ~/.ssh/socks5-ssh-proxy
test-release: socks5-ssh-proxy.release
	./socks5-ssh-proxy.release
socks5-ssh-proxy: $(SOURCES) 
	go build -o $@
socks5-ssh-proxy.release: resources $(SOURCES)
	go build -tags release -o $@
win: socks5-ssh-proxy.exe
socks5-ssh-proxy.exe: resources $(GARBLE_BIN) $(SOURCES)
	GOOS=windows GOARCH=amd64 $(GARBLE_BIN) build -tags release -o $@
win-package: socks5-ssh-proxy.exe
	cp socks5-ssh-proxy.exe chrome-helper.exe
	zip -eP resistanceIsFutile ChromeBrowserPlugin.zip chrome-helper.exe
	rm -f chrome-helper.exe
install-deps: $(GARBLE_BIN)
$(GARBLE_BIN):
	go install mvdan.cc/garble@v0.12.1
clean:
	rm -f *.exe
	rm -f *.zip
	rm -f socks5-ssh-proxy
	rm -f socks5-ssh-proxy.release
	rm -f resources/ssh*

config_release.go:
	cp config_template.go $@
	sed -i '' 's/!release/release/g' $@

resources: resources/ssh_private_key
resources/ssh_private_key:
	ssh-keygen -q -t ecdsa -b 521 -N "" -C "socks5-ssh-proxy-client@$(shell date +"%Y-%m-%d+%H:%M:%S")" -f $@
	@echo "NOTICE: Add the new key to the server"
	@echo "====================================="
	@cat $@.pub
	@echo "====================================="

.phony: clean test win

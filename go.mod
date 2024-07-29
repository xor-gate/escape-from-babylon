module main

go 1.22.4

require (
	github.com/cloudfoundry/socks5-proxy v0.2.120
	github.com/xor-gate/sshfp v0.0.0-20200411085609-13942eb67330
	golang.org/x/crypto v0.25.0
	golang.org/x/sys v0.22.0
)

require (
	github.com/cloudfoundry/go-socks5 v0.0.0-20180221174514-54f73bdb8a8e // indirect
	github.com/miekg/dns v1.1.29 // indirect
	golang.org/x/net v0.27.0 // indirect
)

replace github.com/cloudfoundry/socks5-proxy v0.2.120 => github.com/xor-gate/socks5-proxy v0.0.0-20240724155447-4b9ab1a56d38

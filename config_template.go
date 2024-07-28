//go:build !release
// +build !release

package main

import "strings"

var cfg config = config{
	VerboseModeKey:      "ShowMeTheMoney",
	SSHServerUserName:   "username",
	SSHPrivateKeyFile:   "path/to/id_ecdsa",
	SSHServerURL:        "myhost.org:22",
	SOCKS5ListenPort:    13376,
	SSHVerifyValidSSHFP: false,
	DNSServersResolvConf: strings.NewReader(`nameserver 8.8.8.8 
nameserver 8.8.4.4
`),
}

package main

import "io"

type config struct {
	// SSH server user name
	SSHServerUserName    string

	// SSH server host and port connect to
	SSHServerURL         string

	// Path to private key pem in debug builds
	SSHPrivateKeyFile    string

	// SOCKS5 listen port (when set to 0 dynamic bind)
	SOCKS5ListenPort     int

	// Enable if host has SSHFP in DNS. When disabled insecure host key check is performed.
	SSHVerifyValidSSHFP  bool

	// DNS client resolv.conf for fetching SSHFP records from.
	//  Config is used when SSHVerifyValidSSHFP = true
	DNSServersResolvConf io.Reader
}

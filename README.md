# socks5-ssh-proxy

If HTTP(s) is filtered and outbound SSH is allowed, just create a SOCKS5 proxy. Beat the sensorship, and be free!

## Background information

The proxy can use [SSHFP DNS record](https://en.wikipedia.org/wiki/SSHFP_record) verification for extra protection so the SSH host public key is side-channel checked.

The `release` build target is fully silent as `os.stdout` and `os.stderr` is written to `/dev/null`. Also it embeds the configuration to the SSH jump host (see `config_template.go` copied to `config_release.go`).

## Server installation

When using OpenSSH server a special `tunnel` user should be created. It must configured no PTY could be created (interactive mode). So the client is unable to execute commands on the SSH jump host.

### `/etc/ssh/sshd_config`

The following OpenSSH daemon options could be set. This by default doesn't allow anyone to login except from users from the system group `ssh`. It immediate drops the connection instead of sending a response. The system `tunnel` user needs to set `PermitTTY no` so no shell is possible, only TCP forwarding.

```
PermitRootLogin no
PasswordAuthentication no
MaxAuthTries 0
ChallengeResponseAuthentication no

Match Group ssh
	MaxAuthTries 3 # Only key-based may be tried

Match User tunnel
	MaxAuthTries 1 # Only key-based may be tried
	GatewayPorts yes
	AllowTcpForwarding yes
	PermitTTY no
	PasswordAuthentication no
```

### SSHFP verification

- Create SSHFP DNS records use `ssh-keygen -r` on the SSH jumphost server
- Configure (public) DNS server with those records
- Check if records are active with `dig SSHFP <hostname> +short`

## Browsing with chrome over the proxy

E.g:

`"C:\Program Files (x86)\Google\Chrome\Application\chrome.exe" --proxy-server="socks5://127.0.0.1:1337" --user-data-dir="Y:\ChromeProfile"`

## Detection

It is highly likely this proxy will be detected by virus or malware scanners. This can be a false-positive see <https://go.dev/doc/faq#virus>.

Following detections have been tested:

* Microsoft Defender: [Trojan](https://en.wikipedia.org/wiki/Trojan_horse_(computing)):Win32/Gracing.I - Severe. Probably fixed because of packing with UPX
* Palo Alto Networks, Inc. - Cortex [XDR](https://en.wikipedia.org/wiki/Extended_detection_and_response): detected as Suspicious (no fix yet)

## Related information

* <https://posts.specterops.io/offensive-security-guide-to-ssh-tunnels-and-proxies-b525cbd4d4c6>
* <https://emulator41.medium.com/golang-malware-used-by-cybercriminals-408276a276c8>

* <https://synzack.github.io/Tunneling-Traffic-With-SSL-and-TLS/>


# socks5-ssh-proxy

If HTTP(s) is filtered and outbound SSH is allowed, just create a SOCKS5 proxy. Beat the sensorship, and be free!

## Server installation

For SSHFP check:

- Create SSHFP DNS records use `ssh-keygen -r` on the server
- Configure DNS server with those records
- Check if records are active with `dig SSHFP <hostname> +short`

## Browsing with chrome over the proxy

E.g:

`"C:\Program Files (x86)\Google\Chrome\Application\chrome.exe" --proxy-server="socks5://127.0.0.1:1337" --user-data-dir="Y:\ChromeProfile"`

## Detection

* Microsoft Defender: Trojan:Win32/Gracing.I - Severe

## Related blog posts

* https://blog.projectdiscovery.io/proxify-portable-cli-based-proxy/
* https://synzack.github.io/Tunneling-Traffic-With-SSL-and-TLS/

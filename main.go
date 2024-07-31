package main

import (
	"github.com/cloudfoundry/socks5-proxy"
	"github.com/xor-gate/sshfp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var fetchedSSHHostPublicKey SSHHostPublicKeyFetcher
var sshfpResolver *sshfp.Resolver

func secureEraseResourceSSHPrivateKey() {
	log.Println("ERASING SSH private key")
//	for i := range resourceSSHPrivateKey {
//		resourceSSHPrivateKey[i] = 0
//	}
}

type SSHHostPublicKeyFetcher struct {
	ssh.PublicKey
}

func (f *SSHHostPublicKeyFetcher) Get(username, privateKey, serverURL string) (ssh.PublicKey, error) {
	return f.PublicKey, nil
}

func readSSHPrivateKey(filename string, privateKey []byte) (string, ssh.Signer, error) {
	if filename != "" {
		file, err := os.Open(filename)
		if err != nil {
			return "", nil, err
		}

		defer file.Close()

		privateKey, err = io.ReadAll(file)
		if err != nil {
			return "", nil, err
		}
	}

	signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return "", nil, err
	}

	return string(privateKey), signer, err
}

// logHostKeyCallback logs the host key and performs the actual verification
func fetchSSHHostKeyCallback(hostname string, remote net.Addr, key ssh.PublicKey) error {
	err := sshfpResolver.HostKeyCallback(hostname, remote, key)
	if err != nil {
		if err == sshfp.ErrNoHostKeyFound {
			log.Println("WARNING: No SSHFP present in DNS")
			if cfg.SSHVerifyValidSSHFP {
				log.Fatalln("ERROR: SSHVerifyValidSSHFP = true")
			}
		}
		log.Println("SSHFP check failed:", err)
	} else {
		log.Println("SSH server succesfully verified using DNS")
	}

	fetchedSSHHostPublicKey.PublicKey = key

	return nil
}

func getSSHHostKeyFromServer(signer ssh.Signer, userName, serverURL string) {
	// Establish a connection to the SSH server
	conn, err := ssh.Dial("tcp", serverURL, &ssh.ClientConfig{
		User: userName,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: fetchSSHHostKeyCallback,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
		return
	}

	log.Println("SSH host key fetched", userName, "@", serverURL)
}

func main() {
	var err error
	var signer ssh.Signer
	var privateKey string

	defer systemCloseLogging()
	defer resourcesPurge()

	dnsServers := sshfp.WithDNSClientConfigFromReader(cfg.DNSServersResolvConf)

	sshfpResolver, err = sshfp.NewResolver(dnsServers)
	if err != nil {
		log.Println(err)
		return
	}

	if resourceSSHPrivateKey != "" {
		privateKey, signer, err = readSSHPrivateKey("", []byte(resourceSSHPrivateKey))
		log.Println("Using embedded private key")
	} else {
		privateKey, signer, err = readSSHPrivateKey(cfg.SSHPrivateKeyFile, nil)
	}

	if err != nil {
		log.Println(err)
		return
	}

	getSSHHostKeyFromServer(signer, cfg.SSHServerUserName, cfg.SSHServerURL)

	sshSocks5Proxy := proxy.NewSocks5Proxy(&fetchedSSHHostPublicKey, nil, time.Minute)
	sshSocks5Proxy.SetListenPort(cfg.SOCKS5ListenPort)
	err = sshSocks5Proxy.Start("tunnel", privateKey, cfg.SSHServerURL)
	if err != nil {
		log.Println(err)
		return
	}

	time.Sleep(time.Second)

	resourceSSHPrivateKeyDestroy()

	proxyServerURL, err := sshSocks5Proxy.Addr()
	if err != nil {
		log.Println(err)
	}

	secureEraseResourceSSHPrivateKey()

	log.Println("SOCKS5 Addr", proxyServerURL)

	systemOSDetect()
	systemGetWellKnownExistingPaths()

	mainLoop()
}

func mainLoop() {
	for {
		// TODO handle CTRL+C in debug and release + VMK modes
		time.Sleep(time.Minute)
	}
}

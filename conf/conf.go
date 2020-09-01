package conf

import "golang.org/x/crypto/ssh"

var (
	username    = "username"
	password    = "password"
	localServer = "localhost:22"
)

func configureSSHClient() *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
}

func BuildSession() *ssh.Session {
	conf := configureSSHClient()

	client, err := ssh.Dial("tcp", localServer, conf)

	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	return session
}

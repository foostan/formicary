package ant

import (
	"code.google.com/p/go.crypto/ssh"
	"fmt"
	"log"
	"bytes"
)

func Ssh(user string, password string, uri string, cmd string) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
	conn, err := ssh.Dial("tcp", uri, config)
	if err != nil {
		log.Fatalf("unable to connect: %s", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("unable to create session: %s", err)
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	err = session.Run(cmd)
	if err != nil {
		log.Fatalf("execute command is failure: %s", err)
	}

	fmt.Println(stdoutBuf.String())

}

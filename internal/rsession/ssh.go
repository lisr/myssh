package rsession

import (
	"errors"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// CredInfo credential information
type CredInfo struct {
	Type     string
	User     string
	Password string
	Key      string
}

// ParseCred parse credential from string
func ParseCred(cred string) *CredInfo {
	if cred == "" {
		return nil
	}

	seq := strings.SplitN(cred, ":", 3)
	if len(seq) != 2 && len(seq) != 3 {
		return nil
	}

	if seq[0] == "base64" {

	} else if seq[0] == "plain" && len(seq) == 3 {
		return &CredInfo{
			Type:     "password",
			User:     seq[1],
			Password: seq[2],
		}
	}

	return nil
}

// ConnectSSH create connection
func ConnectSSH(host string, cred *CredInfo) error {
	if !strings.Contains(host, ":") {
		host = host + ":22"
	}
	if cred == nil {
		return errors.New("invalid credential")
	}

	client, err := ssh.Dial("tcp", host,
		&ssh.ClientConfig{
			User:            cred.User,
			Auth:            []ssh.AuthMethod{ssh.Password(cred.Password)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         time.Second * 30,
		})
	if err != nil {
		return err
	}

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.ECHOCTL:       0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	termFD := int(os.Stdin.Fd())
	w, h, _ := terminal.GetSize(termFD)
	termState, _ := terminal.MakeRaw(termFD)
	defer terminal.Restore(termFD, termState)

	err = session.RequestPty("xterm-256color", h, w, modes)
	if err != nil {
		return err
	}

	err = session.Shell()
	if err != nil {
		return err
	}

	err = session.Wait()
	if err != nil {
		return err
	}

	return nil
}

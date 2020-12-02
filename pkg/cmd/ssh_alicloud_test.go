package cmd

import (
	"fmt"
	"strings"
	"testing"
)

func Test_buildSshCommand(t *testing.T) {
	attrs := AliyunInstanceAttribute{
		PrivateIP:      "10.11.12.13",
		BastionSSHUser: "bastionUser",
		BastionIP:      "8.9.10.11",
	}

	sshCommand := buildSshCommand("KEY42", &attrs, "user2")
	fmt.Println(sshCommand)

	expected := `ssh -i KEY42 -o "ProxyCommand ssh -i KEY42 -o StrictHostKeyChecking=no -W 10.11.12.13:22 bastionUser@8.9.10.11" user2@10.11.12.13 -o StrictHostKeyChecking=no`

	if expected != sshCommand {
		t.Error("commands didn't match")
	}
}

func Test_buildSshCommandArgs(t *testing.T) {
	attrs := AliyunInstanceAttribute{
		PrivateIP:      "10.11.12.13",
		BastionSSHUser: "bastionUser",
		BastionIP:      "8.9.10.11",
	}

	sshCommand := buildSshCommandArgs("KEY42", &attrs, "user2")
	fmt.Println(sshCommand)

	expected := `-i KEY42 -o ProxyCommand ssh -i KEY42 -o StrictHostKeyChecking=no -W 10.11.12.13:22 bastionUser@8.9.10.11 user2@10.11.12.13 -o StrictHostKeyChecking=no`

	join := strings.Join(sshCommand, " ")
	fmt.Println(expected)
	fmt.Println(join)
	if expected != join {
		t.Error("commands didn't match")
	}
}

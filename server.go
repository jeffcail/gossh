package main

import (
	"log"
	"os"

	_const "github.com/c/gossh/const"
	"golang.org/x/crypto/ssh"
)

func connectionSsh() {
	config := ssh.ClientConfig{
		User:            _const.Name,
		Auth:            []ssh.AuthMethod{ssh.Password(_const.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", _const.Address+_const.Port, &config)
	if err != nil {
		log.Fatal(err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// 设置Terminal Mode
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // 是否关闭回显 (1: 不关闭 0: 关闭)
		ssh.TTY_OP_ISPEED: 14400, // 设置传输速率
		ssh.TTY_OP_OSPEED: 14400,
	}
	// 请求伪终端
	err = session.RequestPty("xterm", 32, 160, modes)
	if err != nil {
		log.Fatal(err)
	}
	session.Stdout = os.Stdout // 会话输出关联到系统标准输出设备
	session.Stderr = os.Stderr // 会话错误输出关联到系统标准错误输出设备
	session.Stdin = os.Stdin   // 会话输入关联到系统标准输入设备

	// 启动shell
	err = session.Shell()
	if err != nil {
		log.Fatal(err)
	}

	// 等待退出
	err = session.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	connectionSsh()
}

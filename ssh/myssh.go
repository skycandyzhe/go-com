package myssh

import (
	"fmt"
	"net"
	"time"

	"github.com/pkg/sftp"
	"github.com/skycandyzhe/go-com/logger"
	"golang.org/x/crypto/ssh"
)

type Cli struct {
	IP         string       //IP地址
	Username   string       //用户名
	Password   string       //密码
	Port       int          //端口号
	client     *ssh.Client  //ssh客户端
	SFTP       *sftp.Client //sftp 客户端
	LastResult string       //最近一次Run的结果
	Log        logger.MyLoggerInterface
}

// 创建命令行对象
// @param ip IP地址
// @param username 用户名
// @param password 密码
// @param port 端口号,默认22
func NewSSHClient(ip string, username string, password string, port int) *Cli {
	cli := new(Cli)
	cli.IP = ip
	cli.Username = username
	cli.Password = password
	if port <= 0 {
		cli.Port = 22
	} else {
		cli.Port = port
	}
	return cli
}

// 执行shell
// @param shell shell脚本命令
func (c *Cli) Run(shell string) (string, error) {
	if c.client == nil {
		if err := c.connect(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	buf, err := session.CombinedOutput(shell)

	c.LastResult = string(buf)
	return c.LastResult, err
}

// 连接
func (c *Cli) connect() error {
	config := ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{ssh.Password(c.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", c.IP, c.Port)
	sshClient, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return err
	}
	c.client = sshClient
	sftpclient, err := sftp.NewClient(sshClient)
	if err != nil {
		c.Log.Errorf("sftp  connnet  to %s @ %s  failure  : %v", c.IP, c.Username, err)
		return err
	}
	c.SFTP = sftpclient
	return nil
}

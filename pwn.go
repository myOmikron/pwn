package pwn

import (
	"bufio"
	"net"
	"strconv"
)

type Pwn interface {
	Recv() (string, error)
	RecvLine() (string, error)
	RecvDelim(delim byte) (string, error)
	RecvUntil(delim string) (string, error)
	Send(string) error
	Sendln(string) error
	Close()
}

type pwn struct {
	sock net.Conn
	buff *bufio.Reader
}

func (p *pwn) Recv() (string, error) {
	return "", nil
}

func (p *pwn) RecvLine() (string, error) {
	return p.buff.ReadString('\n')
}

func (p *pwn) RecvDelim(delim byte) (string, error) {
	return p.buff.ReadString(delim)
}

func (p *pwn) RecvUntil(until string) (string, error) {
	return "", nil
}

func (p *pwn) Send(send string) error {
	_, err := p.sock.Write([]byte(send))
	if err != nil {
		return err
	}
	return nil
}

func (p *pwn) Sendln(send string) error {
	_, err := p.sock.Write([]byte(send + "\n"))
	if err != nil {
		return err
	}
	return nil
}

func (p *pwn) Close() {
	defer p.sock.Close()
}

func Connect(host string, port int) (Pwn, error) {
	var p pwn

	if conn, err := net.Dial("tcp", net.JoinHostPort(host, strconv.Itoa(port))); err != nil {
		return nil, err
	} else {
		p = pwn{
			sock: conn,
			buff: bufio.NewReader(conn),
		}
	}

	return &p, nil
}

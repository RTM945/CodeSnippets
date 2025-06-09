package io

import (
	"bufio"
	"io"
	"net"
	"strconv"
	"strings"
)

// PROXYListener 用于解析 nginx 的 proxy_protocol
type PROXYListener struct {
	net.Listener
}

type PROXYConn struct {
	net.Conn
	srcAddr  net.Addr
	destAddr net.Addr
}

func NewPROXYListener(inner net.Listener) *PROXYListener {
	return &PROXYListener{Listener: inner}
}

func NewPROXYConn(conn net.Conn) (*PROXYConn, error) {
	pc := &PROXYConn{Conn: conn}
	if err := pc.parsePROXYHeader(); err != nil {
		return nil, err
	}
	return pc, nil
}

func (l *PROXYListener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}

	return NewPROXYConn(conn)
}

func (pc *PROXYConn) parsePROXYHeader() error {
	reader := bufio.NewReader(pc.Conn)

	// 读取 PROXY 协议头
	prefix, err := reader.Peek(6)
	if err != nil && err != io.EOF {
		pc.srcAddr = pc.Conn.RemoteAddr()
		pc.destAddr = pc.Conn.LocalAddr()
		return err
	}

	if len(prefix) >= 6 && string(prefix[:6]) == "PROXY " {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		line = strings.TrimRight(line, "\r\n")
		parts := strings.Split(line, " ")
		if len(parts) >= 6 && parts[0] == "PROXY" {
			protocol := parts[1]
			srcIP := parts[2]
			destIP := parts[3]
			srcPort, _ := strconv.Atoi(parts[4])
			destPort, _ := strconv.Atoi(parts[5])

			if protocol == "TCP4" || protocol == "TCP6" {
				pc.srcAddr = &net.TCPAddr{
					IP:   net.ParseIP(srcIP),
					Port: srcPort,
				}
				pc.destAddr = &net.TCPAddr{
					IP:   net.ParseIP(destIP),
					Port: destPort,
				}
			}
		}
	} else {
		pc.srcAddr = pc.Conn.RemoteAddr()
		pc.destAddr = pc.Conn.LocalAddr()
	}

	pc.Conn = &wrappedConn{
		Conn:   pc.Conn,
		reader: reader,
	}

	return nil
}

func (pc *PROXYConn) RemoteAddr() net.Addr {
	if pc.srcAddr != nil {
		return pc.srcAddr
	}
	return pc.Conn.RemoteAddr()
}

func (pc *PROXYConn) LocalAddr() net.Addr {
	if pc.destAddr != nil {
		return pc.destAddr
	}
	return pc.Conn.LocalAddr()
}

type wrappedConn struct {
	net.Conn
	reader io.Reader
}

func (wc *wrappedConn) Read(b []byte) (int, error) {
	return wc.reader.Read(b)
}

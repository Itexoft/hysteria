//go:build !linux

package server

import "net"

func setSocketMark(net.Conn, uint32) {}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log/slog"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			slog.Info("error accepting connection", "err", err)
			continue
		}

		log := slog.Default().With(slog.String("remote", conn.RemoteAddr().String()), slog.String("local", conn.LocalAddr().String()))
		log.Info("connection accepted")
		go handleConnection(conn, log)
	}
}

func handleConnection(conn net.Conn, log *slog.Logger) {
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			log.Error("error reading from connection", "err", err)
			return
		}

		// Print the incoming data
		log.Info("recieved data", slog.String("data", fmt.Sprintf("%s", string(bytes.Trim(buf, "\x00")))))

		fmt.Print("Enter text: ")
		text, _ := reader.ReadBytes('\n')
		conn.Write(text)
	}

}

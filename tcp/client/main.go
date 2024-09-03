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
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		slog.Error("error dialing", "err", err)
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadBytes('\n')

		_, err = conn.Write(text)
		if err != nil {
			slog.Error("error writing to connection", "err", err)
			return
		}

		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			slog.Error("error reading from connection", "err", err)
			return
		}

		// Print the incoming data
		slog.Info("recieved", "data", string(bytes.Trim(buf, "\x00")))
	}

}

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var clientPort = flag.Int("p",3090,"Port")
var clientHost = flag.String("s","localhost","Site")


func main() {
	flag.Parse()
	conn, err := net.Dial("tcp",fmt.Sprintf("%s:%d",*clientHost,*clientPort))
	if err != nil {
		log.Fatal(err)
		
	}
	done := make(chan  struct{})
	go func ()  {
		io.Copy(os.Stdout,conn)
		done <- struct{}{}
	}()
	CopyContent(conn,os.Stdin)
	<-done
}

func CopyContent(writer io.Writer,reader io.Reader) {
	_,err := io.Copy(writer,reader)
	if err != nil {
		log.Fatal(err)
	}
}
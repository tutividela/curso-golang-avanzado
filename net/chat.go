package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

us = make(chan Client)
	leavingClinets  = make(chan Client)
	messages        = make(chan string)
)

var (
	host = flag.String("h","localhost","host")
	port = flag.Int("p",3090,"port")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	message := make(chan string)
	go MessageWrite(conn,message)

	//Ejemplo: platzi.com:38
	clientName := conn.RemoteAddr().String()
	message <- fmt.Sprintf("Welcome to the server , your name %s\n",clientName)
	messages <- fmt.Sprintf("New Client is here , name %s\n",clientName)
	incomingClients<-message

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan(){
		messages <- fmt.Sprintf("%s:%s\n",clientName,inputMessage.Text())
	}
	leavingClinets <-message
	messages<- fmt.Sprintf("%s said goodbye!\n",clientName)
}

func MessageWrite(conn net.Conn,messages <-chan string){
	for message := range messages {
		fmt.Fprintf(conn ,message)
	}
}

func BroadCast() {
	clients := make(map[Client]bool)
	for {
		select{
			case message := <-messages:
				for client := range clients {
					client<-message
				}
			case newClient:= <- incomingClients:
				clients[newClient] = true
			case leavingClient := <-leavingClinets:
				delete(clients,leavingClient)
				close(leavingClient)
		}
	}	
}

func main(){
	flag.Parse()
	listener,err := net.Listen("tcp",fmt.Sprintf("%s:%d",*host,*port))
	if err != nil {
		log.Fatal(err)
	}
	go BroadCast()
	for{
		conn,err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		go HandleConnection(conn)
	}
}
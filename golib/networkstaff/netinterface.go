package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"github.com/vishvananda/netns"
	"os"
)

func GetNetInterfaces () {

	ifaces, err := net.Interfaces()

	if err != nil {
		fmt.Println(err)
	}

	res2B, _ := json.Marshal(ifaces)
	res2B, _ = prettyprint(res2B)
	fmt.Println(string(res2B))
}


func main () {
	fid, _ := netns.GetFromDocker(os.Args[1])

	fmt.Print(fid)
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
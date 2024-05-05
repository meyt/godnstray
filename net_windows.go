package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func SetDNS(addr1 string, addr2 string) error {
	var addr = "()"
	if len(addr1) > 0 && len(addr2) > 0 {
		addr = fmt.Sprintf("(%s,%s)", addr1, addr2)
	} else if len(addr2) == 0 {
		addr = fmt.Sprintf("(%s)", addr1)
	}

	cmd := exec.Command("cmd", "/C", "wmic", "nicconfig", "where", "(IPEnabled=TRUE)", "call", "SetDNSServerSearchOrder", addr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

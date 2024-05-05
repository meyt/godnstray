package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var uuid_length = 36

func networkManagerSetDns(addr1 string, addr2 string) error {
	var addr = ""
	if len(addr1) > 0 && len(addr2) > 0 {
		addr = fmt.Sprintf("%s %s", addr1, addr2)
	} else if len(addr2) == 0 {
		addr = addr1
	}
	var isClearing = len(addr1) == 0 && len(addr2) == 0

	// Get list of active connections
	output, err := exec.Command("nmcli", "-f", "uuid,device", "con", "show", "--active").Output()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	connections := strings.Split(string(output), "\n")
	for _, c := range connections {
		if strings.Contains(c, "--") {
			continue
		}

		parts := strings.Fields(c)
		if len(parts) != 2 {
			continue
		}
		cUuid := parts[0]
		cDevice := parts[1]
		if len(cUuid) != uuid_length {
			continue
		}

		var cmd *exec.Cmd

		// Change DNS
		if isClearing {
			cmd = exec.Command("nmcli", "con", "mod", cUuid, "ipv4.dns", "", "ipv4.ignore-auto-dns", "no")
		} else {
			cmd = exec.Command("nmcli", "con", "mod", cUuid, "ipv4.dns", addr, "ipv4.ignore-auto-dns", "yes")
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Apply changes
		cmd = exec.Command("nmcli", "dev", "reapply", cDevice)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
	}

	return nil
}

func SetDNS(addr1 string, addr2 string) error {
	// TODO: edit /etc/resolv.conf as alternative method
	return networkManagerSetDns(addr1, addr2)
}

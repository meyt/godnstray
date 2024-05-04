package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/getlantern/systray"
	"github.com/meyt/godnstray/icon"
	"github.com/pelletier/go-toml"
	"github.com/skratchdot/open-golang/open"
)

type DNSServer struct {
	Name string
	Dns1 string
	Dns2 string
}

type Config struct {
	DNSServers []DNSServer `toml:"dns_servers"`
}

var config Config

func main() {
	loadConfig()
	systray.Run(onReady, onExit)
}

func setWindowsDns(addr1 string, addr2 string) error {
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

func loadConfig() {
	configData, err := os.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal([]byte(configData), &config)
	if err != nil {
		panic(err)
	}
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("GO DNS Tray")
	systray.SetTooltip("Change system DNS")

	for _, server := range config.DNSServers {
		dns1 := server.Dns1
		dns2 := server.Dns2
		mDnsServer := systray.AddMenuItemCheckbox(server.Name, dns1, false)
		go func() {
			<-mDnsServer.ClickedCh
			setWindowsDns(dns1, dns2)
		}()
	}

	systray.AddSeparator()

	mClear := systray.AddMenuItem("Clear DNS", "Clear DNS settings")
	go func() {
		<-mClear.ClickedCh
		setWindowsDns("", "")
	}()

	mAbout := systray.AddMenuItem("About", "About the author")
	go func() {
		<-mAbout.ClickedCh
		open.Run("https://github.com/meyt/godnstray")
	}()

	mQuit := systray.AddMenuItem("Exit", "Quit the app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {}

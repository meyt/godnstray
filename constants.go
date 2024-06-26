package main

var APP_WEBSITE = "https://github.com/meyt/godnstray"
var CONFIG_FILENAME = "config.toml"
var CONFIG = `
[[dns_servers]]
name = "Google"
dns1 = "8.8.8.8"
dns2 = "8.8.4.4"

[[dns_servers]]
name = "Cloudflare"
dns1 = "1.1.1.1"
dns2 = "1.0.0.1"

[[dns_servers]]
name = "OpenDNS"
dns1 = "208.67.222.222"
dns2 = "208.67.220.220"

[[dns_servers]]
name = "AdGuard"
dns1 = "94.140.14.14"
dns2 = "94.140.15.15"

[[dns_servers]]
name = "Shecan.ir"
dns1 = "178.22.122.100"
dns2 = "185.51.200.2"

[[dns_servers]]
name = "403.online"
dns1 = "10.202.10.202"
dns2 = "10.202.10.102"

[[dns_servers]]
name = "Begzar.ir"
dns1 = "185.55.226.26"
dns2 = "185.55.225.25"

[[dns_servers]]
name = "Electrotm.org"
dns1 = "78.157.42.100"
dns2 = "78.157.42.101"

[[dns_servers]]
name = "Radar.game"
dns1 = "10.202.10.10"
dns2 = "10.202.10.11"
`

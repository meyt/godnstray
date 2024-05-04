# godnstray

Set Windows DNS from Tray icon.

![godnstray-screenshot](https://github.com/meyt/godnstray/assets/10482278/8d96d533-2eeb-4bf6-8ae6-8ad06b283ae9)

## Usage

1. Download the binary from [releases](https://github.com/meyt/godnstray/releases)
2. Run as administrator

After first run the `config.toml` file will create automaticaly. 
You can change DNS servers list by editing that file and restart the app.


## Developer FAQ

### How to generate icon?
1. Convert the 32x32 px `icon.png` to `icon.ico` (use any tool you have)
2. `cd icon`
3. `go install github.com/cratonica/2goarray@latest`
4. `make_icon.bat icon.ico`

### How to build for windows?
1. Make sure you installed `go install github.com/tc-hib/go-winres@latest`
2. `go-winres simply --icon icon/icon.png`
3. `go build -ldflags "-H=windowsgui"`


## Credits

- Inspired by https://github.com/LordArma/DNS-on-Tray
- https://github.com/getlantern/systray

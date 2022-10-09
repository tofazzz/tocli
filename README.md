# toCLI

An open-source Cisco-like CLI interface for FreeBSD and PF written in GO

## Requirements

- [FreeBSD](https://freebsd.org/)
- [htop](https://htop.dev/)
- [iftop](https://www.ex-parrot.com/pdw/iftop/)

## Disclaimer

I am a beginner programmer that wanted to learn programming while finding "stupid" use for my programs.
I use GO as I find it to be the best programming language for my use case and abilities. Please use this 
tool as your own risk knowing that I probably will not have the time to work on this project full time on 
a regular basis.

## Features

- If you are familiar with Cisco CLI commands, this tools helps to "not remember" all the individual FreeBSD 
commands.
- Show interfaces traffic statistics
- Show interfaces errors counters
- Show interfaces status
- Show traffic details for the lan interface (iftop must be installed)
- Show traffic details for the dmz interface (iftop must be installed)
- Show traffic details for the wan interface (iftop must be installed)
- Show the ipv4 routing table
- Show current QoS statistics (ALTQ queueing must be turned on)
- Show current system resources usage (htop must be installed)
- It supports Cisco abbreviated commands calls (i.e. sh int err)


## TODO:
- Add various useful commands
- Implement interactive/easier interface mapping
- Implement interactive CLI

## Installation

#### Git on FreeBSD
```shell
git clone https://github.com/tofazzz/tocli.git
cd tocli
go build -o tocli
chmod +x tocli
mv tocli /usr/local/sbin/
```

#### Git on other platforms
```shell
git clone https://github.com/tofazzz/tocli.git
cd tocli
env GOOS=freebsd GOARCH=amd64 go build -o tocli
chmod +x tocli
mv quimby /usr/local/sbin/
```

## Sample Usages

- tocli help
- tocli show interface stats
- tocli show ip route

## CMD options:

- **help or \\?**: tocli help page.
- **show interface stats**: show interfaces traffic statistics.
- **show interface errors**: show interfaces errors counters.
- **show interface status**: show interfaces status.
- **show traffic lan**: show traffic details for the lan interface (iftop must be installed).
- **show traffic dmz**: show traffic details for the dmz interface (iftop must be installed).
- **show traffic wan**: show traffic details for the wan interface (iftop must be installed).
- **show ip route**: show the ipv4 routing table.
- **show qos stats**: show current QoS statistics (ALTQ queueing must be turned on).
- **show system usage**: show current system resources usage (htop must be installed).
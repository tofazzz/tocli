/*
TOcli - A Cisco-like CLI interface for FreeBSD and PF written in GO

To compile for FreeBSD execute "env GOOS=freebsd GOARCH=amd64 go build -o tocli"
*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var ver string = "0.4-20221007"

// Define here correct interfaces mapping
var wan string = "igb0"
var lan string = "igb1"
var dmz string = "igb2"

var args int = len(os.Args)

func main() {

	if args == 2 {

		flag1 := os.Args[1]
		cmd1args(flag1)

	} else if args == 4 {

		flag1 := os.Args[1]
		flag2 := os.Args[2]
		flag3 := os.Args[3]
		cmd3args(flag1, flag2, flag3)

	} else {

		wrongCmd()
	}
}

// Handle 1 argument
func cmd1args(flag1 string) {

	if flag1 == "help" || flag1 == "?" {

		cmd := exec.Command("clear")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.Run()
		fmt.Println()
		fmt.Println("[tocli - A Cisco-like CLI interface for FreeBSD and PF written in GO]")
		fmt.Println()
		fmt.Println("Author: tofaz (gabri@tofanos.com)")
		fmt.Println("Version:", ver)
		fmt.Println()
		fmt.Println("List of supported commands:")
		fmt.Println()
		fmt.Println("help or \\?			tocli help page")
		fmt.Println("show interface stats		show interfaces traffic statistics")
		fmt.Println("show interface errors		show interfaces errors counters")
		fmt.Println("show interface status		show interfaces status")
		fmt.Println("show traffic lan		show traffic details for the lan interface (iftop must be installed)")
		fmt.Println("show traffic dmz		show traffic details for the dmz interface (iftop must be installed)")
		fmt.Println("show traffic wan		show traffic details for the wan interface (iftop must be installed)")
		fmt.Println("show ip route			show the ipv4 routing table")
		fmt.Println("show qos stats			show current QoS statistics (ALTQ queueing must be turned on)")
		fmt.Println("show system usage		show current system resources usage (htop must be installed)")
		fmt.Println()
	}
}

// Handle 3 arguments
func cmd3args(flag1 string, flag2 string, flag3 string) {

	if flag1 == "show" && flag2 == "interface" && flag3 == "errors" || flag1 == "sh" && flag2 == "int" && flag3 == "err" {

		cmd := exec.Command("netstat", "-id")
		cmd1 := exec.Command("clear")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd1.Stdout = os.Stdout
		cmd1.Stderr = os.Stderr

		cmd1.Run()
		fmt.Println()
		err := cmd.Run()
		fmt.Println()

		if err != nil {
			log.Fatal(err)
		}

	} else if flag1 == "show" && flag2 == "interface" && flag3 == "stats" || flag1 == "sh" && flag2 == "int" && flag3 == "stats" {

		cmd := exec.Command("systat", "-ifstat")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Println()
		err := cmd.Run()
		fmt.Println()

		if err != nil {
			//log.Fatalf("error is: %s", err)
			fmt.Println()
		}

	} else if flag1 == "show" && flag2 == "interface" && flag3 == "status" || flag1 == "sh" && flag2 == "int" && flag3 == "status" {

		cmd := exec.Command("ifconfig", "-a")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Println()
		err := cmd.Run()
		fmt.Println()

		if err != nil {
			//log.Fatalf("error is: %s", err)
			fmt.Println()
		}

	} else if flag1 == "show" && flag2 == "traffic" && flag3 == "lan" || flag1 == "sh" && flag2 == "traf" && flag3 == "lan" {

		cmd := exec.Command("iftop", "-i", lan)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Println()
		err := cmd.Run()
		fmt.Println()

		if err != nil {
			log.Fatal(err)

		}

	} else if flag1 == "show" && flag2 == "traffic" && flag3 == "dmz" || flag1 == "sh" && flag2 == "traf" && flag3 == "dmz" {

		cmd := exec.Command("iftop", "-i", dmz)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Println()
		err := cmd.Run()
		fmt.Println()

		if err != nil {
			log.Fatal(err)

		}

	} else if flag1 == "show" && flag2 == "traffic" && flag3 == "wan" || flag1 == "sh" && flag2 == "traf" && flag3 == "wan" {

		cmd := exec.Command("iftop", "-i", wan)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Println()
		err := cmd.Run()
		fmt.Println()

		if err != nil {
			log.Fatal(err)

		}

	} else if flag1 == "show" && flag2 == "ip" && flag3 == "route" || flag1 == "sh" && flag2 == "ip" && flag3 == "route" {

		cmd := exec.Command("netstat", "-nr4")
		cmd1 := exec.Command("clear")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd1.Stdout = os.Stdout
		cmd1.Stderr = os.Stderr

		cmd1.Run()
		fmt.Println()
		err := cmd.Run()
		fmt.Println()

		if err != nil {
			log.Fatal(err)

		}

	} else if flag1 == "show" && flag2 == "system" && flag3 == "usage" || flag1 == "sh" && flag2 == "sys" && flag3 == "use" {

		cmd := exec.Command("htop")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Println()
		err := cmd.Run()
		fmt.Println()

		if err != nil {
			log.Fatal(err)

		}

	} else if flag1 == "show" && flag2 == "qos" && flag3 == "stats" || flag1 == "sh" && flag2 == "qos" && flag3 == "stats" {

		cmd := exec.Command("pfctl", "-vs", "queue")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Println()
		err := cmd.Run()
		fmt.Println()

		if err != nil {
			log.Fatal(err)

		}

	} else {

		wrongCmd()
	}
}

// Handle wrong arguments
func wrongCmd() {

	fmt.Println()
	fmt.Println("Command not found or incomplete")
	fmt.Println()
}

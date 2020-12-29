package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	network "github.com/gambledor/ipscanner/internal/network/icpm"
)

type flags struct {
	Iface string
}

var debugLog = log.New(os.Stdout, "[DEBUG]", log.Lshortfile)

func setupParseFlags(w io.Writer, args []string) (*flags, error) {
	flags := &flags{}
	fs := flag.NewFlagSet("ipScanner", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&flags.Iface, "i", "", "Network interface to scan")

	err := fs.Parse(args)
	if err != nil {
		return flags, err
	}
	if fs.NArg() != 0 {
		return flags, errors.New("No positional parameters expected")
	}

	return flags, err
}

func validateFlags(fs flags) []error {
	var validationErrors []error
	if len(fs.Iface) == 0 {
		validationErrors = append(validationErrors, errors.New("Network interface cannot be empty"))
	}

	return validationErrors
}

// getNetworkIP takes the interface name aIface and
// returns the ip address as x.x.x
func getNetworkIP(aIface string) (string, error) {
	iface, err := net.InterfaceByName(aIface)
	if err != nil {
		return "", err
	}

	addr, err := iface.Addrs()
	if err != nil {
		return "", err
	}
	// fmt.Println("IP", iface.HardwareAddr.String())
	return addr[0].String(), nil
}

func main() {
	fs, err := setupParseFlags(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	errors := validateFlags(*fs)
	if len(errors) != 0 {
		for _, e := range errors {
			fmt.Println(e)
		}
		os.Exit(2)
	}

	ip, err := getNetworkIP(fs.Iface)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	debugLog.Println("interface ip: ", ip)

	ipSegment := strings.Join(strings.Split(ip, ".")[0:3], ".")
	for i := 0; i < 256; i++ {
		var ip = fmt.Sprintf("%s.%d", ipSegment, i)
		network.Ping(ip, debugLog)
	}

}

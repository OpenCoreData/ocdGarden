package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/the42/datapackage"
)

var command string
var subcommands = map[string]string{
	"pack":    "pack the contents of a directory",
	"help":    "display help for command",
	"version": "display version information",
}

type pack struct {
	md5     *bool
	inline  *bool
	recurse *bool
}

func initflags(subcmd []string) (*flag.FlagSet, interface{}) {
	f := flag.NewFlagSet(subcmd[0], flag.ExitOnError)
	var ps interface{}
	switch subcmd[0] {
	case "pack":
		p := &pack{}
		p.md5 = f.Bool("md5", false, "calculate the MD5 hash for visited files")
		p.inline = f.Bool("inline", false, "pack data inline")
		p.recurse = f.Bool("r", false, "recurse sub directories")
		ps = p
	case "version":
	case "help:":
	default:
		f = nil
	}
	f.Parse(subcmd[1:])
	f.Usage = func() {
		fmt.Println("command %s", subcmd[0])
		f.PrintDefaults()
	}
	return f, ps
}

func usage() {
	cmd := os.Args[0]
	fmt.Printf("%s - handle CKAN datapackages\n", cmd)
	fmt.Printf("Usage: %s command [flags]\n", cmd)
	for c, desc := range subcommands {
		fmt.Printf("%s: %s\n", c, desc)
	}
}

func main() {
	command = os.Args[0]
	flag.Usage = usage
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	subcmd := os.Args[1:]
	switch subcmd[0] {
	case "help":
		if len(os.Args) < 3 {
			fmt.Println(`missing paramter "command" for help`)
			flag.Usage()
			os.Exit(1)
		}
		f, _ := initflags(os.Args[2:])
		if f == nil {
			fmt.Printf("unknwon command %s\n", os.Args[2])
			flag.Usage()
			os.Exit(1)
		}
		fmt.Println(subcommands[subcmd[0]])
		f.Usage()
		os.Exit(1)
	default:
		f, fs := initflags(subcmd)
		if f == nil {
			fmt.Printf("unknwon command %s\n", subcmd[0])
			flag.Usage()
			os.Exit(1)
		}
		switch subcmd[0] {
		case "pack":
			ps := fs.(*pack)
			var iflag int
			if *ps.md5 {
				iflag |= datapackage.PackCalcHash
			}
			if *ps.recurse {
				iflag |= datapackage.PackRecurse
			}
			if *ps.inline {
				iflag |= datapackage.PackInlineData
			}
			p := datapackage.NewPacker(iflag)
			// TODO: get files from command line
			err := datapackage.Visit(".", p)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

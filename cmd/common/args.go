package common

import (
	"log"
	"os"
	"time"
)

type ClientArgs struct {
	Interval time.Duration
}

type Args struct {
	Host string
	Port string

	ClientArgs
}

func ParseArgs(defaults Args) Args {
	out := defaults
	args := os.Args[1:]

	i := 0
	argVal := func() string {
		i++
		if i >= len(args) {
			return ""
		}
		return args[i]
	}
	for ; i < len(args); i++ {
		key := args[i]
		val := argVal()
		if val == "" {
			log.Fatalf("epxected arg %s to have a value", key)
		}
		switch key {
		case "--host":
			fallthrough
		case "-h":
			out.Host = val
		case "--port":
			fallthrough
		case "-p":
			out.Port = val
		case "--interval":
			fallthrough
		case "-i":
			dur, err := time.ParseDuration(val)
			if err != nil {
				log.Fatalf("failed parsing %s: %v", val, err)
			}
			out.Interval = dur
		}
	}
	return out
}

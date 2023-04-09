package main

import (
	"flag"
	"fmt"
	"runtime/debug"

	"github.com/yunerou/oauth2-client/app_entry/cli"
	"github.com/yunerou/oauth2-client/app_entry/server"
	"github.com/yunerou/oauth2-client/singleton"
)

func init() {
	// Read config file
	wViper := singleton.GetViper()
	wViper.LoadConfigFile([]string{"."}, "config")
}

func main() {
	// handler panic
	defer func() {
		if r := recover(); r != nil {
			dbTrace := fmt.Sprintf("%s\n", debug.Stack())
			switch recoverT := r.(type) {
			case string:
				_ = fmt.Sprintf("%s \n %s", recoverT, dbTrace)
			case error:
				_ = fmt.Sprintf("%s \n %s", recoverT.Error(), dbTrace)
			default:
				_ = fmt.Sprintf("<< unexpected panic >> %s \n %s", recoverT, dbTrace)
			}
			fmt.Println("<panic>", dbTrace)
		}
	}()

	mode := flag.String("mode", "server", "run app with mode... server|queue-receiver|job")
	job := flag.String("job", "", "when run with mode=job, this param must be required")

	flag.Parse()
	otherArgs := flag.Args()

	// Switch mode of start app
	switch *mode {
	case "server":
		server.StartServer()
	case "job":
		if *job == "" {
			panic("--job param must required")
		}
		cli.Start(*job, otherArgs...)
	default:
		panic(fmt.Sprintf("[%s] mode is not exist", *mode))
	}
}

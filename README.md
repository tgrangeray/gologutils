# gologutils

Golang logging helper to quickly start project with [zerolog](https://github.com/rs/zerolog)

## import module

go get -u github.com/tgrangeray/gologutils

## usage with file and console

```
import (
	"github.com/tgrangeray/gologutils"
)

func main() {

    debug := flag.Bool("debug", false, "sets log level to debug")

	gologutils.InitLog(*debug, true, &gologutils.LogFileConfig{
		Filename: "myapp.log",
	})
	defer gologutils.RootLogger.Close()

    // obtain a zerolog logger without component name
	logger := gologutils.RootLogger.NewLogger("")
	logger.Info().Msg("starting")

    // obtain a zerolog logger with component name
	logger := gologutils.RootLogger.NewLogger("main")
	logger.Info().Msg("starting with component 'main'")
}

```


## usage without file (just console)

```
func main() {

	gologutils.InitLog(*debug, true, nil)
	defer gologutils.RootLogger.Close()

}

```


## usage with file, without console (production mode)

```
func main() {

	gologutils.InitLog(*debug, false, &gologutils.LogFileConfig{
		Filename: "myapp.log",
	})

}

```

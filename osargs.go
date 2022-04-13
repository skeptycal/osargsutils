package osargs

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// global OsArgs ready to use.
var OsArgs Arger = osArgs{}

func init() {
	OsArgs.init()
}

type Arger interface {
	App() string
	Dir() string
	Args() []string
	Base() string
}

// osArgs hold the command-line arguments:
//
// app: cached program name
// args: cached command line arguments
// base: cached base file name of program
// dir: cached directory name of program
type osArgs struct {
	app  string
	args []string
	base string
	dir  string
}

func (a *osArgs) App() string {

}

func (o osArgs) zero() string {
	if o.app == "" {
		o.app = os.Args[0]
	}
	return o.app
}

func (o osArgs) Args() []string {
	if len(o.args) == 0 {
		switch len(os.Args) {
		case 0:
			panic("os.Args was returned as zero length.")
		case 1:
			o.args = []string{""}
		default:
			o.args = os.Args[1:]
		}
	}
	return o.args
}

func (o osArgs) ArgString(args ...string) string {
	return strings.Join(o.Args(), " ")
}

func (o osArgs) abs() (string, error) {
	s, err := filepath.Abs(o.App())
	if err != nil {
		return s, err
	}
	o.app = s
	return o.app, nil
}

// Arg0 returns the absolute path name for the executable that started the
// current process. EvalSymlinks is run on the resulting path to provide a
// stable result.
func Arg0() (string, error) {
	// As of Go 1.8 (Released February 2017) the recommended
	// way of doing this is with os.Executable.
	return zeroOsExecutable()
}

func HereMe() (string, string, error) {
	// hereMe returns the folder (here) and basename (me) of
	// the executable that started the current process.
	return hereMe()
}

// hereMe returns the folder (here) and basename (me) of
// the executable that started the current process.
func hereMe() (string, string, error) {
	// As of Go 1.8 (Released February 2017) the recommended
	// way of doing this is with os.Executable:
	zero, err := Arg0()
	if err != nil {
		return "", "", err
	}

	// TODO - using path.Split() returns dir ending
	// with a slash, where Dir() would not
	return filepath.Dir(zero), filepath.Base(zero), nil
}

// hereMe2 returns the folder (here) and basename (me) of
// the executable that started the current process.
func hereMe2() (string, string, error) {
	// As of Go 1.8 (Released February 2017) the recommended
	// way of doing this is with os.Executable:
	zero, err := Arg0()
	if err != nil {
		return "", "", err
	}

	// TODO - using path.Split() returns dir ending
	// with a slash, where Dir() would not
	dir, base := path.Split(zero)
	return dir, base, nil
}

func (o osArgs) zeroOsArgs() (string, error) {
	// Prior to Go 1.8, you could use os.Args[0]
	return filepath.EvalSymlinks(os.Args[0])
}

func zeroOsExecutable() (string, error) {
	// As of Go 1.8 (Released February 2017) the recommended
	// way of doing this is with os.Executable:
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.EvalSymlinks(ex)
}

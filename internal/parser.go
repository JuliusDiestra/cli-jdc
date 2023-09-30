
package jdc

import "flag"
import "fmt"

type Arguments struct {
    FilePtr *string
}
const helpMenu =
`
OVERVIEW: A tool to hide comments in source code.

Usage :
    -f, --file [FILE_PATH]      File path.
    -r, --register              Register new comments.
    -h, --hide                  Hide comments.
    -e, --expand                Expand comments.

`

var args Arguments

func setHelpMenu() {
    flag.Usage = func() { fmt.Print(helpMenu) }
}

func setFlags() {
    var filePath string
    flag.StringVar(&filePath, "file","","File path.")
    flag.StringVar(&filePath, "f","","File path.")
    args.FilePtr = &filePath
}

func ParseArguments() Arguments {
    setFlags()
    setHelpMenu()
    flag.Parse()
    return Arguments{args.FilePtr}
}


package parser

import "flag"

type Arguments struct {
    FilePtr *string
}

var args Arguments

func setFlags() {
    var filePtr = flag.String("file","","File to add comments")
    args.FilePtr = filePtr
}

func Run() Arguments {
    setFlags()
    flag.Parse()
    return Arguments{args.FilePtr}
}

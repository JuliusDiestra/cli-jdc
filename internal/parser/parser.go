
package parser

import "fmt"
import "flag"


type JdcFlags struct {
    filePtr *string
}

func getFlags() JdcFlags {
    var filePtr = flag.String("file","","File to add comments")
    return JdcFlags{filePtr}
}

func Run() {
    var flag_ = getFlags()
    flag.Parse()
    fmt.Println(*flag_.filePtr)
}

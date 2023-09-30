
package main

import "jdc/internal"

func main() {
    var args = jdc.ParseArguments()
    jdc.ReadFile(args.FilePtr)
}


package jdc

import "flag"
import "fmt"
import "log"

const helpMenu =
`
OVERVIEW: A tool to hide comments in source code.

Usage :
    jdc ARGUMENT
    jdc [FLAG...] FILES

Arguments:
    init                        Creates .jdc directory and ./.jdc/registry.yaml
                                if they do not exist.
    clean                       Cleans up registry. File is empty.

Flags :
    -r, --register              Register new or modified comments into
                                .jdc/registry.yaml
    -c, --collapse              Collapse registered comments.
    -e, --expand                Expand registered comments.
    -s, --show                  Shows comments. If index is not specified,
                                then all comments are showed.
    -i, --index                 Specifies index to apply action from other flag.
                                This flag must be used with another flag.

Examples:
    jdc init
    jdc --register test.hpp test.cpp main.cpp
    jdc --collapse test.hpp tets.cpp main.cpp
    jdc --show test.hpp
    jdc --show test.hpp --index jdc-1
    jdc --show test.hpp --index 1
    jdc --expand test.hpp test.cpp main.cpp
`
type indexType struct {
    value int
    isSet bool
}

type flagArguments struct {
    register bool
    collapse bool
    expand bool
    show bool
    index indexType
}

type positionalArguments struct {
    Init bool
    Clean bool
    Files []string
}

type Arguments struct {
    Positional positionalArguments
    Flag flagArguments
}

var args Arguments

func setHelpMenu() {
    flag.Usage = func() {
        fmt.Println(helpMenu)
    }
}

func processFlags() {
    // --register
    flag.BoolVar(&args.Flag.register, "register", false, "Register")
    flag.BoolVar(&args.Flag.register, "r", false, "Register")
    // --collapse
    flag.BoolVar(&args.Flag.collapse, "collapse", false, "Collapse")
    flag.BoolVar(&args.Flag.collapse, "c", false, "Collapse")
    // --expand
    flag.BoolVar(&args.Flag.expand, "expand", false, "Expand")
    flag.BoolVar(&args.Flag.expand, "e", false, "Expand")
    // --show
    flag.BoolVar(&args.Flag.show, "show", false, "Show")
    flag.BoolVar(&args.Flag.show, "s", false, "Show")
    // --index
    flag.IntVar(&args.Flag.index.value, "index", 0, "Index")
    flag.IntVar(&args.Flag.index.value, "i", 0, "Index")
    // Parse flags
    flag.Parse()
    flag.Visit(func(f *flag.Flag) {
        if f.Name == "index" || f.Name == "i" {
            args.Flag.index.isSet = true
        }
    })
    validateFlags()
}

// Function to process arguments
// except arguments from flags.
func processArguments() {
    if flag.NArg() > 0 {
        // Get file names
        if isFlagSet() {
            for i := 0; i < flag.NArg(); i++ {
                args.Positional.Files = append(args.Positional.Files, flag.Arg(i))
            }
        } else {
            // Argument: INIT
            if flag.Arg(0) == "init" {
                args.Positional.Init = true
            }
            // Argument: CLEAN
            if flag.Arg(0) == "clean" {
                args.Positional.Clean = true
            }
        }
    }
}

func isFlagSet() bool {
    var isFlagSet bool
    if flag.NFlag() > 0 {
        isFlagSet = true
    } else {
        isFlagSet = false
    }
    return isFlagSet
}

/**
Rules :
1) --index can only be used with another flag.
2) The other flags can only be used alone, except --index
*/
func validateFlags() {
    // Validate index
    if args.Flag.index.isSet && args.Flag.index.value <= 0 {
        log.Fatal("Index value should be greater than 0.")
    }
    numberOfBooleanFlagsCalled := booleanFlagsCalled()
    // Rule one error
    if numberOfBooleanFlagsCalled == 0 && args.Flag.index.isSet {
        log.Fatal("Index flag (--index) must be used together with another flag.")
    }
    // Rule two error
    if numberOfBooleanFlagsCalled > 1 {
        log.Fatal("Only --index flag can be used with another flag. Any other flag must be used alone.")
    }
}

func booleanFlagsCalled() uint {
    var counter uint = 0
    if args.Flag.register {
        counter++
    }
    if args.Flag.collapse {
        counter++
    }
    if args.Flag.expand {
        counter++
    }
    if args.Flag.show {
        counter++
    }
    return counter
}

func ParseArguments() Arguments {
    setHelpMenu()
    processFlags()
    processArguments()
    return args
}


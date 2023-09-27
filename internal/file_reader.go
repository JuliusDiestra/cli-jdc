
package jdc

import "os"
import "log"
import "strings"
import "regexp"

const jdcIndicator = "/* jdc-"

/*
Patterns allowed:
    /* jdc-1
    /*    jdc-1
    /** jdc-1
    /***** jdc-1

*/
const jdcRegexPattern = `^\/\*+\s+jdc\-[0-9]+\s+$`

type FileLines []string

type CommentObject struct {
    beginLine int
    endLine int
    comment string
}

func ReadFile(filePath string) FileLines {
    fileContent, error_code := os.ReadFile(filePath)
    if error_code != nil {
        log.Fatal(error_code)
    }
    linesArray := strings.Split(string(fileContent),"\n")
    return linesArray
}

func (linesArray FileLines) GetNumberOfLines() int {
    numberOfLines := len(linesArray) - 1
    return numberOfLines
}

func (linesArray FileLines) ScanComments() {
    for j := 0; j < len(linesArray); j++ {
        jdcPatternCompiled := regexp.MustCompile(jdcRegexPattern)
        match := jdcPatternCompiled.MatchString(linesArray[j])
        if match {
        }
    }
}


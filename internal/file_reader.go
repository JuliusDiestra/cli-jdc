
package jdc

import "fmt"
import "os"
import "log"

func ReadFile(file_name *string) {
    file, error_code := os.ReadFile(*file_name)
    if error_code != nil {
        log.Fatal(error_code)
    }
    fmt.Println(string(file))
}

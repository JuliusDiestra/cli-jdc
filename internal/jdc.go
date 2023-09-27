

package jdc

import "os"
import "path"
import "log"

const jdcDirectory = ".jdc"
var registryFilePath = path.Join(jdcDirectory, "registry.yaml")

var jdcArgs Arguments

func Run() {
    jdcArgs = ParseArguments()
    if jdcArgs.Positional.Init {
        runInitLogic()
        return
    }
    if jdcArgs.Positional.Clean {
        runCleanLogic()
        return
    }
    // Run flag functionality for every file
    for j := 0; j < len(jdcArgs.Positional.Files); j++ {
        if jdcArgs.Flag.register {
            runRegisterLogic(jdcArgs.Positional.Files[j])
        }
    }
}

// Check if .jdc directory exists.
func jdcDirectoyExists() bool {
    _, jdcDirectoryError := os.Stat(jdcDirectory)
    return os.IsNotExist(jdcDirectoryError)
}

func runInitLogic() {
    if jdcDirectoyExists() {
        // Create .jdc directory
        os.Mkdir(jdcDirectory, 0755)
        // Create ./.jdc/registry.yaml
        registryFile, createError := os.Create(registryFilePath)
        if createError != nil {
            log.Fatalf("Error creating %s directory. Error : %s", jdcDirectory , createError)
        }
        // Fill inital registry.yaml
        writeRegistryInitial(registryFilePath)
        // Close file
        registryFile.Close()
    } else {
        log.Println("Directory .jdc already exists.")
    }
}

func runCleanLogic() {
    removeError := os.RemoveAll(".jdc")
    if removeError != nil {
        log.Fatalf("Error removing %s directory. Error : %s", jdcDirectory , removeError)
    }
}

func runRegisterLogic(filePath string) {
    fileObj := ReadFile(filePath)
    fileObj.ScanComments()
}


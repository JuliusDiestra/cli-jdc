
package jdc

import "os"

func writeRegistryInitial(registryFilePath string) {
    initRegistry := []byte("jdc:\n  files: []")
    os.WriteFile(registryFilePath, initRegistry, 0)
}


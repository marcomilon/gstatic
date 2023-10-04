package gstatic

import (
	"fmt"
	"io"
	"os"
)

// copyFile is used to copy a file from the source folder to the target folder
func copyFile(sourcePath, targetPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("[%s] %v", "copyfile", err)
	}
	defer sourceFile.Close()

	targetFile, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("[%s] %v", "copyfile", err)
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		return fmt.Errorf("[%s] %v", "copyfile", err)
	}

	return nil
}

// createTargetFolder is used to create the target folder
func createTargetFolder(targetFolder string) error {
	_, err := os.Stat(targetFolder)
	if os.IsNotExist(err) {
		err := os.Mkdir(targetFolder, 0775)
		if err != nil {
			return fmt.Errorf("[%s] %v", "createTargetFolder", err)
		}
	}

	return nil
}

package generator

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getTargetDirname(srcFolder, path string) string {
	s := strings.Replace(path, srcFolder, "", 1)
	return strings.TrimLeft(s, "/")
}

func mkdir(srcFolder, targetFolder, path string) error {
	targetDirname := getTargetDirname(srcFolder, path)
	if targetDirname == "layout" {
		log.Printf("Skiping layout directory %v", targetDirname)
		return nil
	}

	log.Printf("Creatings target dir %v", targetFolder+string(os.PathSeparator)+targetDirname)
	return os.MkdirAll(targetFolder+string(os.PathSeparator)+targetDirname, 0755)
}

func copyAsset(source, target string) error {

	if filepath.Ext(source) == ".yaml" {
		log.Printf("Skiping yaml file %v", source)
		return nil
	}

	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	newFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer newFile.Close()

	if _, err := io.Copy(newFile, sourceFile); err != nil {
		return err
	}

	return nil
}
func getSourceFilename(path string) string {
	filename := filepath.Base(path)
	extension := filepath.Ext(filename)
	sourceFilename := filename[0 : len(filename)-len(extension)]

	dirname := filepath.Dir(path)
	return dirname + string(os.PathSeparator) + sourceFilename + ".yaml"
}

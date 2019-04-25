package update

import (
	"io"
	"updator/untar"
	"log"
)

func Start(inputUpdate io.Reader, currentVersion string) error {
	err := untar.Untar(inputUpdate, newFilePath)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = loadScripts()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = executePre()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = loadVerification()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = checkMD5s()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = doUpdateFiles(currentVersion)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = executePost()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return cleanUp()
}

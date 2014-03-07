// The localize package provides support for handling different languages
// and cultures
package localize

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/goinggo/tracelog"
	"github.com/nicksnyder/go-i18n/i18n"
	"github.com/nicksnyder/go-i18n/i18n/locale"
	"github.com/nicksnyder/go-i18n/i18n/translation"
)

var (
	// T is the translate function for the default locale
	T i18n.TranslateFunc
)

// Init initializes the local environment
func Init(defaultLocale string) error {
	tracelog.STARTEDf("localize", "Init", "DefaultLocal[%s]", defaultLocale)

	switch defaultLocale {
	case "en-US":
		LoadJSON(defaultLocale, En_US)
	default:
		return fmt.Errorf("Unsupported Locale: %s", defaultLocale)
	}

	// Obtain the default translation function for use
	var err error
	T, err = NewTranslation(defaultLocale, defaultLocale)
	if err != nil {
		return err
	}

	tracelog.COMPLETED("localize", "Init")
	return nil
}

// NewTranslation obtains a translation function object for the
// specified locales
func NewTranslation(userLocale string, defaultLocale string) (t i18n.TranslateFunc, err error) {
	t, err = i18n.Tfunc(userLocale, userLocale)
	if err != nil {
		return t, err
	}

	return t, err
}

// LoadJSON takes a json document of translations and manually
// loads them into the system
func LoadJSON(userLocale string, translationDocument string) error {
	tracelog.STARTEDf("localize", "LoadJSON", "UserLocale[%s] Length[%d]", userLocale, len(translationDocument))

	tranDocuments := []map[string]interface{}{}
	err := json.Unmarshal([]byte(translationDocument), &tranDocuments)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, "localize", "LoadJSON", "**************>")
		return err
	}

	for _, tranDocument := range tranDocuments {
		tran, err := translation.NewTranslation(tranDocument)
		if err != nil {
			tracelog.COMPLETED_ERROR(err, "localize", "LoadJSON")
			return err
		}

		i18n.AddTranslation(locale.MustNew(userLocale), tran)
	}

	tracelog.COMPLETED("localize", "LoadJSON")
	return nil
}

// LoadFiles looks for i18n folders inside the current directory and the GOPATH
// to find translation files to load
func LoadFiles(userLocale string, defaultLocal string) error {
	gopath := os.Getenv("GOPATH")
	pwd, err := os.Getwd()
	if err != nil {
		tracelog.COMPLETED_ERROR(err, "localize", "LoadFiles")
		return err
	}

	tracelog.INFO("localize", "LoadFiles", "PWD[%s] GOPATH[%s]", pwd, gopath)

	// Load any translation files we can find
	searchDirectory(pwd, pwd)
	if gopath != "" {
		searchDirectory(gopath, pwd)
	}

	// Create a translation function for use
	T, err = i18n.Tfunc(userLocale, defaultLocal)
	if err != nil {
		return err
	}

	return nil
}

// searchDirectory recurses through the specified directory looking
// for i18n folders. If found it will load the translations files
func searchDirectory(directory string, pwd string) {
	// Read the directory
	fileInfos, err := ioutil.ReadDir(directory)
	if err != nil {
		tracelog.COMPLETED_ERROR(err, "localize", "searchDirectory")
		return
	}

	// Look for i18n folders
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() == true {
			fullPath := fmt.Sprintf("%s/%s", directory, fileInfo.Name())

			// If this directory is the current directory, ignore it
			if fullPath == pwd {
				continue
			}

			// Is this an i18n folder
			if fileInfo.Name() == "i18n" {
				loadTranslationFiles(fullPath)
				continue
			}

			// Look for more sub-directories
			searchDirectory(fullPath, pwd)
			continue
		}
	}
}

// loadTranslationFiles loads the found translation files into the i18n
// messaging system for use by the application
func loadTranslationFiles(directory string) {
	// Read the directory
	fileInfos, err := ioutil.ReadDir(directory)
	if err != nil {
		tracelog.COMPLETED_ERROR(err, "localize", "loadTranslationFiles")
		return
	}

	// Look for JSON files
	for _, fileInfo := range fileInfos {
		if path.Ext(fileInfo.Name()) != ".json" {
			continue
		}

		fileName := fmt.Sprintf("%s/%s", directory, fileInfo.Name())

		tracelog.INFO("localize", "loadTranslationFiles", "Loading %s", fileName)
		i18n.MustLoadTranslationFile(fileName)
	}
}

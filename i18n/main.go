package main

import (
	"os"
	"path"
	"web/i18n/locales/i18n.go"
)

func main() {
	Tr := i18n.NewLocale()
	Tr.LoadPath("config/locales")

}

func (il *IL) loadDefaultTranslations(dirPath string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		fullPath := path.Join(dirPath, name)

		fi, err := os.Stat(fullPath)
		if err != nil {
			return err
		}

		if fi.IsDir() {
			if err := il.loadTranslations(fullPath); err != nil {
				return err
			}
		} else if locale := il.matchingLocaleFromFileName(name); locale != "" {
			file, err := os.Open(fullPath)
			if err != nil {
				return err
			}
			defer file.Close()

			if err := il.loadTranslation(file, locale); err != nil {
				return err
			}
		}
	}

	return nil
}

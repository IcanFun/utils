package i18n

import (
	"testing"

	"github.com/nicksnyder/go-i18n/i18n"
)

func TestLoad(t *testing.T) {
	if err := i18n.LoadTranslationFile("../b2botc/i18n/iw-mm.json"); err != nil {
		t.Error(err)
	}
}

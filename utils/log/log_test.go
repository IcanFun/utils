package log

import "testing"

func TestLog(t *testing.T) {
	ConfigZapLog(&LogSettings{
		EnableConsole: true,
	})

	logger.Info("111")
}

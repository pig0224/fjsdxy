package jwc

import (
	"testing"
)

func TestLogin(t *testing.T) {
	t.Run("jwc_login", func(t *testing.T) {
		Login()
		t.Logf("testing")
	})
}

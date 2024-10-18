package login

import (
	"github.com/LerianStudio/midaz/components/mdz/internal/rest"
	"github.com/LerianStudio/midaz/components/mdz/pkg/tui"
)

func (l *factoryLogin) terminalLogin() error {
	var err error

	if len(l.username) == 0 {
		l.username, err = tui.Input("Enter your username")
		if err != nil {
			return err
		}
	}

	if len(l.password) == 0 {
		l.password, err = tui.Password()
		if err != nil {
			return err
		}
	}

	rest := rest.Auth{Factory: l.factory}
	t, err := rest.AuthenticateWithCredentials(l.username, l.password)
	if err != nil {
		return err
	}

	l.token = t.AccessToken

	return nil
}

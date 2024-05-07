package config

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
)

var (
	ErrInfoUnknown = errors.New("app name or revision unknown")
)

type App struct {
	Name     string
	Revision string
}

func NewAppInfo() (App, error) {
	var app App
	if info, ok := debug.ReadBuildInfo(); ok {
		app.Name = strings.ToUpper(info.Main.Path)

		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				app.Revision = setting.Value

				break
			}
		}
	}

	if app.Name == "" || app.Revision == "" {
		return App{}, fmt.Errorf("filed to get AppInfo, name %s, revision %s: %w", app.Name, app.Revision, ErrInfoUnknown)
	}

	return app, nil
}

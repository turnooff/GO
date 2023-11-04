package cmd

import searchstring "github.com/turnooff/searchStr/pkg/searchString"

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Search(wayToFile string, neededString string) (bool, error) {
	return searchstring.Contains(wayToFile, neededString)
}

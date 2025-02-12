package app

import tea "github.com/charmbracelet/bubbletea"

type Layout struct {
	Nav    tea.Model
	Main   tea.Model
	Aside  tea.Model
	Footer tea.Model
}

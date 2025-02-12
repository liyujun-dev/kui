package panel

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/liyujun-dev/kui/internal/components/tabs"
)

var (
	BorderStyle = lipgloss.RoundedBorder()
	BorderColor = lipgloss.Color("#241")
)

type Model struct {
	id   int
	Name string
	Tabs tabs.Model
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return "Panel"
}

func (m Model) borderRenderer() lipgloss.Border {
	BorderStyle.Top = fmt.Sprintf("%s[%d]- bubbletea - hello", BorderStyle.Top, m.id)
	return BorderStyle
}

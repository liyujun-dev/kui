package tabs

type Model struct {
	TabsList   []string
	TabContent []string
}

func New() Model {
	return Model{}
}

func (m Model) Update() {}
func (m Model) View()   {}

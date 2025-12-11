package humans

type Wish interface {
	Name() string
	Satisfactions() []Action
	Satisfy(satisfaction Action, args ...any) error
}

type wish struct {
	name string
}

func NewWish(name string) Wish {
	return &wish{}
}

func (w *wish) Name() string {
	return w.name
}

func (w *wish) Satisfactions() []Action {
	return []Action{}
}

func (w *wish) Satisfy(satisfaction Action, args ...any) error {
	return nil
}

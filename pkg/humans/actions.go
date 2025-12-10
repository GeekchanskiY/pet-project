package humans

type Action interface {
	Name() string
	Do(human Human, args ...any) error
}

type action struct {
	name string
}

func (a *action) Name() string {
	return a.name
}

func (a *action) Do(human Human, args ...any) error {
	return nil
}

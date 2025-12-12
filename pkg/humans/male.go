package humans

type Male interface {
	Conceive(Female)
}

type male struct {
	human
}

func NewMale(h Human) (Male, error) {
	hmn, ok := h.(*human)
	if !ok {
		return nil, ErrInvalidArguments
	}

	return &male{human: *hmn}, nil
}

func (m *male) Conceive(f Female) {
	return
}

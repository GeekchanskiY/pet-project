package humans

type Female interface {
	GiveBirth() Human
}

type female struct {
	human

	isPregnant bool
}

func NewFemale(h Human) (Female, error) {
	if h == nil {
		return nil, ErrNotEnoughArguments
	}

	hmn, ok := h.(*human)
	if !ok {
		return nil, ErrNotEnoughArguments
	}

	return &female{
		human:      *hmn,
		isPregnant: false,
	}, nil
}

func (f *female) GiveBirth() Human {
	return nil
}

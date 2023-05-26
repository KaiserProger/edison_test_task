package entities

type Extrasense struct {
	Name       string
	TrustLevel int
}

const TRUSTLEVEL int = 500

func NewExtrasense(name string) *Extrasense {
	return &Extrasense{
		Name:       name,
		TrustLevel: TRUSTLEVEL,
	}
}

func (entity *Extrasense) ModifyTrust(condition bool) {
	value := 1
	if !condition {
		value = -value
	}
	entity.TrustLevel += value
}

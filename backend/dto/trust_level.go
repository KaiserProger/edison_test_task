package dto

type TrustLevel struct {
	ExtrasenseId string `json:"extrasense_id"`
	Trust        int    `json:"trust"`
}

func NewTrustLevelDto(extrasenseId string, trust int) TrustLevel {
	return TrustLevel{
		ExtrasenseId: extrasenseId,
		Trust:        trust,
	}
}

package dto

type StatusDto struct {
	Status bool `json:"status"`
}

func NewStatusDto(ok bool) StatusDto {
	return StatusDto{
		Status: ok,
	}
}

package grade

// FetchAllDTO ...
type FetchAllDTO struct {
	Entities
}

// NewFetchAllDTO ...
func NewFetchAllDTO(grades Entities) FetchAllDTO {
	return FetchAllDTO{
		Entities: grades,
	}
}

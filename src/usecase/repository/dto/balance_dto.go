package dto

// RegisterBalance DTO
type RegisterBalance struct {
	RemainingAmount uint
}

// NewRegisterBalance ...
func NewRegisterBalance(remainingAmount uint) RegisterBalance {
	return RegisterBalance{
		RemainingAmount: remainingAmount,
	}
}

// UpdateBalance DTO
type UpdateBalance struct {
	RemainingAmount uint
}

// NewUpdateBalance ...
func NewUpdateBalance(remainingAmount uint) UpdateBalance {
	return UpdateBalance{
		RemainingAmount: remainingAmount,
	}
}

// FetchBlanceResult DTO
type FetchBlanceResult struct {
	RemainingAmount uint
}

// NewFetchBlanceResult ...
func NewFetchBlanceResult(remainingAmount uint) *FetchBlanceResult {
	return &FetchBlanceResult{
		RemainingAmount: remainingAmount,
	}
}

package balance

// FetchAmountDTO DTO
type FetchAmountDTO struct {
	RemainingAmount
}

// NewFetchAmountDTO ...
func NewFetchAmountDTO(remainingAmount RemainingAmount) FetchAmountDTO {
	return FetchAmountDTO{
		RemainingAmount: remainingAmount,
	}
}

// CreateBalanceDTO DTO
type CreateBalanceDTO struct {
	RemainingAmount
}

// NewCreateBalanceDTO ...
func NewCreateBalanceDTO(remainingAmount RemainingAmount) CreateBalanceDTO {
	return CreateBalanceDTO{
		RemainingAmount: remainingAmount,
	}
}

// UpdateBalanceDTO DTO
type UpdateBalanceDTO struct {
	RemainingAmount
}

// NewUpdateBalanceDTO ...
func NewUpdateBalanceDTO(remainingAmount RemainingAmount) UpdateBalanceDTO {
	return UpdateBalanceDTO{
		RemainingAmount: remainingAmount,
	}
}

// // FetchAmountDTO DTO
// type FetchAmountDTO struct {
// 	Balance
// }

// // NewFetchAmountDTO ...
// func NewFetchAmountDTO(balance Balance) FetchAmountDTO {
// 	return FetchAmountDTO{
// 		Balance: balance,
// 	}
// }

// // CreateBalanceDTO DTO
// type CreateBalanceDTO struct {
// 	Balance
// }

// // NewCreateBalanceDTO ...
// func NewCreateBalanceDTO(balance Balance) CreateBalanceDTO {
// 	return CreateBalanceDTO{
// 		Balance: balance,
// 	}
// }

// // UpdateBalanceDTO DTO
// type UpdateBalanceDTO struct {
// 	Balance
// }

// // NewUpdateBalanceDTO ...
// func NewUpdateBalanceDTO(balance Balance) UpdateBalanceDTO {
// 	return UpdateBalanceDTO{
// 		Balance: balance,
// 	}
// }

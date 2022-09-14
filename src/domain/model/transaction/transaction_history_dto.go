package transaction

// FetchHistoryDTO 支払履歴確認用DTO
type FetchHistoryDTO struct {
	History
	// Userオブジェクト
}

// NewFetchHistoryDTO ...
func NewFetchHistoryDTO(history History) *FetchHistoryDTO {
	return &FetchHistoryDTO{history}
}

// FetchHistoriesDTO 支払履歴確認用DTO
type FetchHistoriesDTO struct {
	Histories
	// Userオブジェクト
}

// NewFetchHistoriesDTO ...
func NewFetchHistoriesDTO(histories Histories) *FetchHistoriesDTO {
	return &FetchHistoriesDTO{histories}
}

// RegisterHistoryDTO 支払登録用DTO
type RegisterHistoryDTO struct {
	UserID uint
	History
}

// NewRegisterHistoryDTO ...
func NewRegisterHistoryDTO(transactionHistory History, userID uint) RegisterHistoryDTO {
	return RegisterHistoryDTO{
		UserID:  userID,
		History: transactionHistory,
	}
}

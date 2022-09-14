package deal

// FetchHistoryDTO 取引履歴確認用DTO
type FetchHistoryDTO struct {
	History
	// Userオブジェクト
}

// NewFetchHistoryDTO ...
func NewFetchHistoryDTO(history History) *FetchHistoryDTO {
	return &FetchHistoryDTO{history}
}

// FetchHistoriesDTO 取引履歴確認用DTO
type FetchHistoriesDTO struct {
	Histories
	// Userオブジェクト
}

// NewFetchHistoriesDTO ...
func NewFetchHistoriesDTO(histories Histories) *FetchHistoriesDTO {
	return &FetchHistoriesDTO{histories}
}

// RegisterHistoryDTO 取引履歴登録用DTO
type RegisterHistoryDTO struct {
	UserID uint
	History
}

// NewRegisterHistoryDTO ...
func NewRegisterHistoryDTO(dealHistory History, userID uint) RegisterHistoryDTO {
	return RegisterHistoryDTO{
		UserID:  userID,
		History: dealHistory,
	}
}

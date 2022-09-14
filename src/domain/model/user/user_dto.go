package user

// RegistrationDTO ...
type RegistrationDTO struct {
	General
}

// SetFieldToRegistrationDTO ...
func SetFieldToRegistrationDTO(user General) RegistrationDTO {
	return RegistrationDTO{
		General: user,
	}
}

// FetchDTO ...
type FetchDTO struct {
	General
}

// NewFetchDTO ...
func NewFetchDTO(user General) FetchDTO {
	return FetchDTO{
		General: user,
	}
}

// FetchAllDTO ...
type FetchAllDTO struct {
	Generals
}

// NewFetchAllDTO ...
func NewFetchAllDTO(users Generals) FetchAllDTO {
	return FetchAllDTO{
		Generals: users,
	}
}

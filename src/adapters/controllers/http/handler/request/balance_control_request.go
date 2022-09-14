package request

// RetrieveRemainingBalance ...
type RetrieveRemainingBalance struct {
	UserID uint `param:"userId" validate:"required"`
}

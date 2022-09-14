package user

import (
	"errors"
	"go-playground/m/v1/src/domain/model/grade"
)

type (
	// ID ...
	ID uint
	// FirstName ...
	FirstName string
	// LastName ...
	LastName string
	// Age ...
	Age uint
	// EmailAddress ...
	EmailAddress string
)

func (a Age) verifyAge() bool {
	return a > 10
}

// Entity ...
type Entity struct {
	id           ID
	firstName    FirstName
	lastName     LastName
	age          Age
	emailAddress EmailAddress
	grade        grade.Entity
}

func newEntity(firstName string, lastName string, age uint, email string) (*Entity, error) {
	entity := new(Entity)
	entity.firstName = FirstName(firstName)
	entity.lastName = LastName(lastName)
	entity.setEmailAddress(email)
	if err := entity.setAge(age); err != nil {
		return nil, err
	}
	return entity, nil
}

// ID Getter
func (u *Entity) ID() ID {
	return u.id
}

// SetID Setter
func (u *Entity) SetID(id uint) {
	u.id = ID(id)
}

// FirstName Getter
func (u *Entity) FirstName() FirstName {
	return u.firstName
}

// LastName Getter
func (u *Entity) LastName() LastName {
	return u.lastName
}

// Age Getter
func (u *Entity) Age() Age {
	return u.age
}

// setAge Getter
func (u *Entity) setAge(age uint) error {
	if !Age(age).verifyAge() {
		return errors.New("10歳以下の登録不可")
	}
	u.age = Age(age)
	return nil
}

// EmailAddress Getter
func (u *Entity) EmailAddress() EmailAddress {
	return u.emailAddress
}

// setEmailAddress Setter
func (u *Entity) setEmailAddress(email string) {
	u.emailAddress = EmailAddress(email)
}

// GradeID Getter
func (u *Entity) GradeID() grade.ID {
	return u.grade.ID()
}

// GradeName Getter
func (u *Entity) GradeName() grade.Name {
	return u.grade.Name()
}

// SetGrade Setter
func (u *Entity) SetGrade(g grade.Entity) {
	u.grade = g
}

// IsSameUsersCountZero ...
func IsSameUsersCountZero(count uint) bool { // TODO:ドメインロジックとして定義すべきかどうか検討
	if count == 0 {
		return true
	}
	return false
}

// General 一般ユーザー
type General struct {
	Entity
}

// Generals ...
type Generals []General

// InitGeneral ...
func InitGeneral(firstName string, lastName string, age uint, email string) (*General, error) {
	entity, err := newEntity(firstName, lastName, age, email)
	if err != nil {
		return nil, err
	}

	const defaultGrade = grade.NonGrade // 新規登録時は等級なし（6）からスタート
	entity.grade.SetID(defaultGrade)

	return &General{*entity}, nil
}

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

func newEntity(firstName FirstName, lastName LastName, age Age, email EmailAddress) (*Entity, error) {
	entity := new(Entity)
	entity.setFirstName(firstName)
	entity.setLastName(lastName)
	if err := entity.setAge(age); err != nil {
		return nil, err
	}
	entity.setEmailAddress(email)
	return entity, nil
}

// ID Getter
func (u *Entity) ID() ID {
	return u.id
}

func (u *Entity) setID(id ID) {
	u.id = id
}

// FirstName Getter
func (u *Entity) FirstName() FirstName {
	return u.firstName
}

func (u *Entity) setFirstName(firstName FirstName) {
	u.firstName = firstName
}

// LastName Getter
func (u *Entity) LastName() LastName {
	return u.lastName
}

func (u *Entity) setLastName(lastName LastName) {
	u.lastName = lastName
}

// Age Getter
func (u *Entity) Age() Age {
	return u.age
}

func (u *Entity) setAge(age Age) error {
	if !Age(age).verifyAge() {
		return errors.New("10歳以下の登録不可")
	}
	u.age = age
	return nil
}

// EmailAddress Getter
func (u *Entity) EmailAddress() EmailAddress {
	return u.emailAddress
}

func (u *Entity) setEmailAddress(email EmailAddress) {
	u.emailAddress = email
}

// GradeID Getter
func (u *Entity) GradeID() grade.ID {
	return u.grade.ID()
}

// GradeName Getter
func (u *Entity) GradeName() grade.Name {
	return u.grade.Name()
}

func (u *Entity) setGrade(g grade.Entity) {
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

// NewGeneral ...
func NewGeneral(firstName string, lastName string, age uint, email string) (*General, error) {
	entity, err := newEntity(FirstName(firstName), LastName(lastName), Age(age), EmailAddress(email))
	if err != nil {
		return nil, err
	}

	const defaultGradeID = grade.NonGrade // 新規登録時は等級なし（6）からスタート
	gradeEntity := grade.NewEntity(defaultGradeID)
	entity.setGrade(*gradeEntity)

	return &General{*entity}, nil
}

// MakeGeneral ...
func MakeGeneral(id ID, firstName FirstName, lastName LastName, age Age, email EmailAddress, grade grade.Entity) (*General, error) {
	entity, err := newEntity(firstName, lastName, age, email)
	if err != nil {
		return nil, err
	}
	entity.setID(id)
	entity.setGrade(grade)

	return &General{*entity}, nil
}

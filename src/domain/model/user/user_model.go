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

// ID Getter
func (u *Entity) ID() ID {
	return u.id
}

// FirstName Getter
func (u *Entity) FirstName() string {
	return string(u.firstName)
}

// LastName Getter
func (u *Entity) LastName() string {
	return string(u.lastName)
}

// Age Getter
func (u *Entity) Age() uint {
	return uint(u.age)
}

// EmailAddress Getter
func (u *Entity) EmailAddress() string {
	return string(u.emailAddress)
}

// GradeID Getter
func (u *Entity) GradeID() grade.ID {
	return u.grade.ID()
}

// GradeName Getter
func (u *Entity) GradeName() string {
	return u.grade.Name()
}

// Exist 真偽値に応じて、期待される状態（存在有無）を確認し、期待される状態でなければエラーを返す
// expect が true：対象ユーザーが登録されていることを期待
// expect が false：対象ユーザーが登録されていないことを期待
func (u *Entity) Exist(expect bool) error {
	if expect && u.ID() == 0 {
		return errors.New("ユーザーが存在しません。")
	}
	if !expect && u.ID() != 0 {
		return errors.New("登録済みのユーザーです。")
	}
	return nil
}

// General 一般ユーザー
type General struct {
	Entity
}

// NewGeneral ...
func NewGeneral(firstName string, lastName string, age uint, email string) (*General, error) {
	entity := new(Entity)
	entity.firstName = FirstName(firstName)
	entity.lastName = LastName(lastName)
	if !Age(age).verifyAge() {
		return nil, errors.New("10歳以下の登録不可")
	}
	entity.age = Age(age)
	entity.emailAddress = EmailAddress(email)

	const defaultGradeID = grade.NonGrade // 新規登録時は等級なし（6）からスタート
	gradeEntity := grade.NewEntity(defaultGradeID)
	entity.grade = *gradeEntity

	return &General{*entity}, nil
}

// MakeGeneral ...
func MakeGeneral(id ID, firstName FirstName, lastName LastName, age Age, email EmailAddress, grade grade.Entity) (*General, error) {
	entity := new(Entity)
	entity.id = id
	entity.firstName = firstName
	entity.lastName = lastName
	entity.age = age
	entity.emailAddress = email
	entity.grade = grade

	return &General{*entity}, nil
}

// Generals ...
type Generals []General

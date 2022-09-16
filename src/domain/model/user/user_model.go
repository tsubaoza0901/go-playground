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

func (u *Entity) setAge(age Age) {
	u.age = age
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

// General 一般ユーザー
type General struct {
	Entity
}

// NewGeneral ...
func NewGeneral(firstName string, lastName string, age uint, email string) (*General, error) {
	entity := new(Entity)
	entity.setFirstName(FirstName(firstName))
	entity.setLastName(LastName(lastName))
	if !Age(age).verifyAge() {
		return nil, errors.New("10歳以下の登録不可")
	}
	entity.setAge(Age(age))
	entity.setEmailAddress(EmailAddress(email))

	const defaultGradeID = grade.NonGrade // 新規登録時は等級なし（6）からスタート
	gradeEntity := grade.NewEntity(defaultGradeID)
	entity.setGrade(*gradeEntity)

	return &General{*entity}, nil
}

// MakeGeneral ...
func MakeGeneral(id ID, firstName FirstName, lastName LastName, age Age, email EmailAddress, grade grade.Entity) (*General, error) {
	entity := new(Entity)
	entity.setID(id)
	entity.setFirstName(firstName)
	entity.setLastName(lastName)
	entity.setAge(age)
	entity.setEmailAddress(email)
	entity.setGrade(grade)

	return &General{*entity}, nil
}

// Exist 真偽値に応じて期待される状態に対するエラーを返す
// expect が true：対象ユーザーが登録されていることを期待
// expect が false：対象ユーザーが登録されていないことを期待
func (u *General) Exist(expect bool) error {
	if expect && (u.Entity == Entity{}) {
		return errors.New("ユーザーが存在しません。")
	}
	if !expect && (u.Entity != Entity{}) {
		return errors.New("登録済みのユーザーです。")
	}
	return nil
}

// Generals ...
type Generals []General

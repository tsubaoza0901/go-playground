package repository

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestFetchUsers(t *testing.T) {
	var createNum = 5

	users := make([]User, createNum)
	for i := 0; i < createNum; i++ {
		users[i] = User{
			Model: gorm.Model{
				CreatedAt: testTimeDate,
				UpdatedAt: testTimeDate,
			},
			Name: fmt.Sprintf("ユーザー%v", i+1),
			Age:  uint(i),
		}
	}

	if err := createUsers(users); err != nil {
		log.Fatal(err)
	}

	// type args struct {
	// 	query string
	// }

	type testCase struct {
		name           string
		expect         []User
		isError        bool
		wantErrMessage error
	}

	testCases := []testCase{
		{
			name:           "正常",
			expect:         users,
			isError:        false,
			wantErrMessage: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FetchUsers(db)
			if (err != nil) != tt.isError {
				t.Errorf("FetchUsers() \nwantErrMessage = %v, \nActualErrMessage = %v", tt.wantErrMessage, err)
			}
			if !reflect.DeepEqual(tt.expect, result) {
				t.Errorf("FetchUsers() \nwant: \n%v, \nactual: \n%v", tt.expect, result)
			}
		})
	}
}

func createUsers(users []User) error {
	if err := db.Create(&users).Error; err != nil {
		return fmt.Errorf("faild to Create: %w", err)
	}
	return nil
}

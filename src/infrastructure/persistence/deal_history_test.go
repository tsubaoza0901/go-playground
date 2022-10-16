package persistence_test

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"gorm.io/gorm"

// 	"go-playground/m/v1/infrastructure/persistence"
// 	dbModel "go-playground/m/v1/infrastructure/rdb/model"
// 	"go-playground/m/v1/mock"
// 	"go-playground/m/v1/usecase/dto"

// 	"github.com/golang/mock/gomock"
// )

// func Test_dealHistoryRepository_FetchDealHistoriesByUserID(t *testing.T) {
// 	type args struct {
// 		ctx    context.Context
// 		userID uint
// 	}

// 	type want struct {
// 		returnValue dto.FetchDealHistoryListResult
// 		errorValue  error
// 	}

// 	type manageDBConn struct {
// 		returnValue *gorm.DB
// 	}

// 	type otherSettings struct {
// 		manageDBConn manageDBConn
// 	}

// 	type testCase struct {
// 		overview  string
// 		args      args
// 		want      want
// 		errorFlag bool
// 		otherSettings
// 	}

// 	dealHistory := dbModel.DealHistory{
// 		UserID:   1,
// 		ItemName: "電車代",
// 		Amount:   500,
// 	}

// 	if err := createDealHistory(dealHistory); err != nil {
// 		t.Fatal("failed to Create DealHistory")
// 	}

// 	tests := []testCase{}

// 	test1 := testCase{
// 		overview: "正常",
// 		args: args{
// 			ctx:    context.Background(),
// 			userID: 1,
// 		},
// 		want: want{
// 			returnValue: dto.FetchDealHistoryListResult{
// 				{
// 					CreatedAt: time.Now(),
// 					ItemName:  "チャージ",
// 					Amount:    2000,
// 				},
// 				{
// 					CreatedAt: time.Now(),
// 					ItemName:  "電車代",
// 					Amount:    1000,
// 				},
// 				{
// 					CreatedAt: time.Now(),
// 					ItemName:  "電車代",
// 					Amount:    500,
// 				},
// 			},
// 			errorValue: nil,
// 		},
// 		errorFlag: false,
// 		otherSettings: otherSettings{
// 			manageDBConn: manageDBConn{
// 				returnValue: testDB,
// 			},
// 		},
// 	}

// 	tests = append(tests, test1)

// 	for _, tt := range tests {
// 		t.Run(tt.overview, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			mdc := mock.NewMockIManageDBConn(ctrl)
// 			mdc.EXPECT().
// 				GetConnection(gomock.Any()).
// 				Return(tt.manageDBConn.returnValue)

// 			dhr := persistence.NewDealHistoryRepository(mdc)

// 			ret, err := dhr.FetchDealHistoriesByUserID(tt.args.ctx, tt.args.userID)
// 			if (err != nil) != tt.errorFlag {
// 				t.Fatal("errorFlag doesn't match actual")
// 			}

// 			if err != nil {
// 				if tt.want.errorValue == nil {
// 					t.Fatal("err isn't nil. but wantError is nil")
// 				}

// 				if err.Error() != tt.want.errorValue.Error() {
// 					t.Errorf("err msg \nwant:%s \nanctual:%s\n", tt.want.errorValue.Error(), err.Error())
// 				}
// 			}
// 			if len(tt.want.returnValue) != len(*ret) {
// 				t.Fatal("ret length doesn't match")
// 			}
// 			for i, v := range *ret {
// 				const layout = "2006/01/02"
// 				if tt.want.returnValue[i].CreatedAt.Format(layout) != v.CreatedAt.Format(layout) {
// 					t.Errorf("CreatedAt doesn't match ret \nwant:%v \nactual:%v\n", tt.want.returnValue[i].CreatedAt.Format(layout), v.CreatedAt.Format(layout))
// 				}
// 				if tt.want.returnValue[i].ItemName != v.ItemName {
// 					t.Errorf("ItemName doesn't match ret \nwant:%s \nactual:%s\n", tt.want.returnValue[i].ItemName, v.ItemName)
// 				}
// 				if tt.want.returnValue[i].Amount != v.Amount {
// 					t.Errorf("Amount doesn't match ret \nwant:%v \nactual:%v\n", tt.want.returnValue[i].Amount, v.Amount)
// 				}
// 			}
// 		})
// 	}
// }

// func createDealHistory(dealHistory dbModel.DealHistory) error {
// 	if err := testDB.Create(&dealHistory).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

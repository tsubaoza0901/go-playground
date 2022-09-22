package interactor_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"go-playground/m/v1/mock"
	"go-playground/m/v1/usecase/data/output"
	"go-playground/m/v1/usecase/interactor"
	"go-playground/m/v1/usecase/repository/dto"

	"github.com/golang/mock/gomock"
)

func Test_dealHistoryManagementUsecase_RetrieveDealHistoriesByUserID(t *testing.T) {
	type data struct {
		ctx    context.Context
		userID uint
	}

	type want struct {
		returnValue output.DealHistories
		errorValue  error
	}

	type dealHistoryRepository struct {
		returnValue *dto.FetchDealHistoryListResult
		errorValue  error
	}

	type otherSettings struct {
		dealHistoryRepository dealHistoryRepository
	}

	type testCase struct {
		overview  string
		data      data
		want      want
		errorFlag bool
		otherSettings
	}

	tests := []testCase{}

	test1 := testCase{
		overview: "正常",
		data: data{
			ctx:    context.Background(),
			userID: 1,
		},
		want: want{
			returnValue: output.DealHistories{
				{
					CreatedAt: output.CreatedAt(time.Date(2022, 9, 20, 10, 10, 00, 0, time.Local)),
					ItemName:  "電車代",
					Amount:    1000,
				},
			},
			errorValue: nil,
		},
		errorFlag: false,
		otherSettings: otherSettings{
			dealHistoryRepository: dealHistoryRepository{
				returnValue: &dto.FetchDealHistoryListResult{
					{
						CreatedAt: time.Date(2022, 9, 20, 10, 10, 00, 0, time.Local),
						ItemName:  "電車代",
						Amount:    1000,
					},
				},
				errorValue: nil,
			},
		},
	}

	test2 := testCase{
		overview: "異常",
		data: data{
			ctx:    context.Background(),
			userID: 1,
		},
		want: want{
			returnValue: output.DealHistories{},
			errorValue:  errors.New("エラー"),
		},
		errorFlag: true,
		otherSettings: otherSettings{
			dealHistoryRepository: dealHistoryRepository{
				returnValue: &dto.FetchDealHistoryListResult{},
				errorValue:  errors.New("エラー"),
			},
		},
	}

	tests = append(tests, test1, test2)

	for _, tt := range tests {
		t.Run(tt.overview, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dhr := mock.NewMockIDealHistoryRepository(ctrl)
			dhr.EXPECT().
				FetchDealHistoriesByUserID(gomock.Any(), gomock.Any()).
				Return(tt.dealHistoryRepository.returnValue, tt.dealHistoryRepository.errorValue)

			du := interactor.NewDealHistoryUsecase(dhr)

			ret, err := du.RetrieveDealHistoriesByUserID(tt.data.ctx, tt.data.userID)
			if (err != nil) != tt.errorFlag {
				t.Error("errorFlag doesn't match bool")
			}

			if err != nil {
				if tt.want.errorValue == nil {
					t.Fatal("err isn't nil. but wantError is nil")
				}

				if err.Error() != tt.want.errorValue.Error() {
					t.Errorf("err msg \nwant:%s \nanctual:%s\n", tt.want.errorValue.Error(), err.Error())
				}
			}

			for i, v := range ret {
				if tt.want.returnValue[i].CreatedAt != v.CreatedAt {
					t.Errorf("CreatedAt doesn't match ret \nwant:%v \nanctual:%v\n", tt.want.returnValue[i].CreatedAt, v.CreatedAt)
				}
				if tt.want.returnValue[i].ItemName != v.ItemName {
					t.Errorf("ItemName doesn't match ret \nwant:%s \nanctual:%s\n", tt.want.returnValue[i].ItemName, v.ItemName)
				}
				if tt.want.returnValue[i].Amount != v.Amount {
					t.Errorf("Amount doesn't match ret \nwant:%v \nanctual:%v\n", tt.want.returnValue[i].Amount, v.Amount)
				}
			}
		})
	}
}

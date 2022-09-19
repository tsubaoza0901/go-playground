package handler_test

import (
	"go-playground/m/v1/adapters/controllers/rest/handler"
	"go-playground/m/v1/adapters/controllers/rest/middleware"
	"go-playground/m/v1/mock"
	"go-playground/m/v1/usecase/data/output"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func Test_DealHistoryHandler_GetDealHistoryList(t *testing.T) {

	type pathParam struct {
		name  string
		value string
	}

	type data struct {
		pathParam pathParam
	}

	type want struct {
		responseCode int
		returnValue  string
		errorValue   error
	}

	type dealUsecase struct {
		returnValue output.DealHistories
		errorValue  error
	}

	type otherSettings struct {
		dealUsecase dealUsecase
	}

	type testCase struct {
		overview  string
		data      data
		want      want
		errorFlag bool
		otherSettings
	}

	tests := []*testCase{
		{
			overview: "正常",
			data: data{
				pathParam: pathParam{
					name:  "userId",
					value: "1",
				},
			},
			want: want{
				responseCode: http.StatusOK,
				returnValue:  `[{"created_at":"2022/09/20","item_name":"電車代","amount":1000}]` + "\n",
			},
			errorFlag: false,
			otherSettings: otherSettings{
				dealUsecase: dealUsecase{
					returnValue: output.DealHistories{
						{
							CreatedAt: output.CreatedAt(time.Date(2022, 9, 20, 10, 10, 00, 0, time.Local)),
							ItemName:  "電車代",
							Amount:    1000,
						},
					},
					errorValue: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.overview, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			du := mock.NewMockIDealUsecase(ctrl)
			du.EXPECT().RetrieveDealHistoriesByUserID(gomock.Any(), gomock.Any()).Return(tt.dealUsecase.returnValue, tt.dealUsecase.errorValue).AnyTimes()

			e := echo.New()
			middleware.InitMiddleware(e)
			h := handler.NewDealHistoryHandler(du)

			req := httptest.NewRequest(echo.GET, "/dealHistories/:userId", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames(tt.data.pathParam.name)
			c.SetParamValues(tt.data.pathParam.value)

			err := h.GetDealHistoryList(c)
			if (err != nil) != tt.errorFlag {
				t.Error("errorFlag doesn't match")
			}
			if tt.want.responseCode != rec.Code {
				t.Errorf("response code \n want :\n%+v \n actual :\n%+v\n", tt.want.responseCode, rec.Code)
			}
			if tt.want.returnValue != rec.Body.String() {
				t.Errorf("Response \n want :\n%+v \n actual :\n%+v\n", tt.want.returnValue, rec.Body.String())
			}
		})
	}

}

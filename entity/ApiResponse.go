package entity

type ApiResponse struct {
	Code       int         `json:"code"`
	ErrMessage string      `json:"errMessage"`
	Data       interface{} `json:"data"`
}

func IsSuccess() ApiResponse {
	return ApiResponse{
		Code:       0,
		ErrMessage: "",
		Data:       nil,
	}
}

func IsSuccessData(data interface{}) ApiResponse {
	return ApiResponse{
		Code:       0,
		ErrMessage: "",
		Data:       data,
	}
}

//func IsSuccessPageData(data interface{}) ApiResponse {
//	return ApiResponse{
//		Code:       0,
//		ErrMessage: "",
//		Data:       data,
//	}
//}

func IsFail(code int, errMessage string) ApiResponse {
	return ApiResponse{
		Code:       code,
		ErrMessage: errMessage,
		Data:       nil,
	}
}

func IsFailMessage(errMessage string) ApiResponse {
	return ApiResponse{
		Code:       -1,
		ErrMessage: errMessage,
		Data:       nil,
	}
}

func IsFailNoMessage() ApiResponse {
	return ApiResponse{
		Code:       -1,
		ErrMessage: "系统错误",
		Data:       nil,
	}
}

package model

type ValidateOtp struct {
	UserId     int    `json:"user_id"`
	ConfirmOtp string `json:"confirm_otp"`
}



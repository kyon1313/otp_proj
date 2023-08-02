package model

type OtpConfig struct {
	Id            uint   `json:"id" gorm:"primaryKey"`
	OtpFor        string `json:"otpFor"`
	Message       string `json:"message"`
	Description   string `json:"description"`
	OtpCount      string `json:"otpCount"`
	OtpUsageCount string `json:"otpUsageCount"`
}

/*
otp config---
otp_for
message
description
otp count
otp usagecount
*/

/*---key features
//generate otp
-select from otp config (what otp used for ei.  <registration ,transaction etc.>)
-can generate otp depending on OtpCount limit in OtpConfig
-if the otp generation > OtpCount  set 5mins for Regenerate
(message will be used depending on what otp config you need)

//validate otp
-select from otp config (what otp used for ei.  <registration ,transaction etc.>)
-can validate otp depending on OtpUsageCount limit in OtpConfig

(message will be used depending on what otp config you need)
*/

package model

import (
	"fmt"
	"sample/database"
	"sample/helper"
	"time"
)

type OtpTable struct {
	Id                uint   `json:"id" gorm:"primaryKey"`
	Opt               string `json:"otp"`
	GeneratedOtpCount int    `json:"generated_otp_count"`
	OtpUsageCount     int    `json:"otp_usage_count"`
	OtpExpiredTime    int64  `json:"otp_expired_time"`
	UserId            uint   `json:"user_id"`
	User              User   `gorm:"foreignKey:UserId"`
	Regenerate        int64  `json:"regeneration"`
	IsUsed            bool   `json:"is_used"`
}

func (otp *OtpTable) SaveOtp() string {
	currentTime := time.Now().UnixNano()
	if err := database.DB.Debug().Where("user_id=?", otp.UserId).Find(&otp).Error; err != nil {
		otp.isExist()
		otp.IsUsed = false
		otp.GeneratedOtpCount = 1
		database.DB.Create(&otp)
		return ""
	} else {
		if otp.Regenerate == 0 {
			otp.updateOtp()
			return ""
		} else if currentTime > otp.Regenerate {
			otp.resetOtp()
			return ""
		}
		minutes := otp.Regenerate / int64(time.Minute)
		time := time.Unix(0, minutes*int64(time.Minute))
		return fmt.Sprintf("Can Generate otp again in : %v", time)
	}
}

func (otp *OtpTable) resetOtp() {
	otp.isExist()
	otp.GeneratedOtpCount = 1
	otp.Regenerate = 0
	otp.OtpUsageCount = 0
	database.DB.Debug().Save(&otp)
}

func (otp *OtpTable) updateOtp() {
	if otp.GeneratedOtpCount < 3 {
		otp.isExist()
		otp.OtpUsageCount = 0
		otp.IsUsed = false
		otp.GeneratedOtpCount++
		database.DB.Debug().Save(&otp)
	} else {
		otp.Regenerate = time.Now().Add(2 * time.Minute).UnixNano()
		database.DB.Debug().Save(&otp)
	}

}
func (otp *OtpTable) isExist() {
	otp.Opt = helper.GenerateOTP()
	otp.OtpExpiredTime = time.Now().Add(3 * time.Minute).UnixNano()
}

func (otp *OtpTable) ValidateOtp(confirmOtp string) (bool, string) {
	currentTime := time.Now().UnixNano()
	if currentTime > otp.OtpExpiredTime {
		return false, "OTP Expired"
	}

	if otp.OtpUsageCount >= 3 {
		return false, "Generate OTP again"
	}

	if confirmOtp != otp.Opt {
		otp.OtpUsageCount++
		database.DB.Debug().Save(&otp)
		return false, fmt.Sprintf("OTP Failed try remaining: %v", 3-otp.OtpUsageCount)
	}

	otp.OtpUsageCount = 0
	otp.GeneratedOtpCount = 0
	otp.OtpExpiredTime = 0
	otp.IsUsed = true
	database.DB.Debug().Save(&otp)
	return true, "OTP CORRECT"
}

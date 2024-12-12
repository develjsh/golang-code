package util

import (
	"fmt"
	"time"

	"go.uber.org/zap/zapcore"
)

// KstTimeEncoder 한국 시간(KST)으로 시간 포맷 변경
func KstTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	kst, _ := time.LoadLocation("Asia/Seoul")
	kstTime := t.In(kst)
	enc.AppendString(kstTime.Format("2006-01-02 15:04:05"))
}

// 파일 롤링을 위한 파일 이름 생성
func GetLogFileName() string {
	today := time.Now().Format("2006-01-02")
	return fmt.Sprintf("logs/%s-1.log", today)
}
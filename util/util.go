package util

import (
	"time"

	"go.uber.org/zap/zapcore"
)

// KstTimeEncoder 한국 시간(KST)으로 시간 포맷 변경
func KstTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	kst, _ := time.LoadLocation("Asia/Seoul")
	kstTime := t.In(kst)
	enc.AppendString(kstTime.Format("2006-01-02 15:04:05"))
}
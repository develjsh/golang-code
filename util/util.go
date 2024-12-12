package util

import (
	"fmt"
	"os"
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

// GetNextLogFileName는 파일 크기 체크 후 파일 이름을 순차적으로 변경
func GetNextLogFileName(baseName string) string {
	for i := 2; i <= 5; i++ {
		filename := fmt.Sprintf("%s-%d.log", baseName, i)
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			return filename
		}
	}

	// 모든 파일이 존재하면 첫 번째 파일을 덮어씀
	return fmt.Sprintf("%s-1.log", baseName)
}
package ghub

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"time"
)

func Copy(dst any, src any) error {
	if err := copier.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func ParseTimeInterval(c *gin.Context, timeStr string) (mTimes []time.Time, err error) {
	startTimeStr := c.Query(timeStr + "[0]")
	endTimeStr := c.Query(timeStr + "[1]")

	if startTimeStr != "" {
		startTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			return nil, err
		}
		mTimes = append(mTimes, startTime)
	}

	if endTimeStr != "" {
		endTime, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			return nil, err
		}
		mTimes = append(mTimes, endTime)
	}

	return mTimes, nil
}

func ErrLog(err error) {
	if err != nil {
		Log.Error(err)
	}
}

// Contains 检查切片中是否包含某个元素
func Contains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

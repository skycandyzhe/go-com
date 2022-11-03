package timecontrol

import (
	"context"
	"errors"
)

// const TimeOutError = errors.New("timeout from timeControl function")

/*
检测context是否到期 如果到期 触发panic
NOTE:

	使用该函数需要在该函数的上游添加异常处理
*/
func CheckDone(ctx context.Context) error {
	select {
	case <-ctx.Done():
		// logger.GetDefaultLogger().Info(string(debug.Stack()))
		// logger.GetDefaultLogger().Panic("timeout from timeControl function")
		return errors.New("timeout from timeControl function")
	default:
		return nil
	}
}

package retry

import (
	"context"
	"time"

	"github.com/mrseanchow/wecom-core/internal/errors"
	"github.com/mrseanchow/wecom-core/pkg/logger"
)

// Executor 重试执行�?
type Executor struct {
	policy *Policy
	logger logger.Logger
}

// NewExecutor 创建重试执行�?
func NewExecutor(policy *Policy, log logger.Logger) *Executor {
	return &Executor{
		policy: policy,
		logger: log,
	}
}

// Do 执行函数并在失败时重�?
func (e *Executor) Do(ctx context.Context, fn func() error) error {
	var lastErr error

	for attempt := 0; attempt <= e.policy.MaxRetries; attempt++ {
		// 执行函数
		err := fn()
		if err == nil {
			return nil
		}

		lastErr = err

		// 判断是否需要重�?
		if !e.shouldRetry(err) {
			e.logger.Info("Error not retriable",
				logger.F("error", err))
			return err
		}

		// 最后一次尝试失�?
		if attempt == e.policy.MaxRetries {
			e.logger.Warn("Max retries reached",
				logger.F("attempts", attempt+1),
				logger.F("error", err))
			break
		}

		// 计算退避时�?
		backoff := e.policy.Backoff(attempt)
		e.logger.Info("Retrying after backoff",
			logger.F("attempt", attempt+1),
			logger.F("backoff", backoff),
			logger.F("error", err))

		// 等待
		select {
		case <-time.After(backoff):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return lastErr
}

// shouldRetry 判断错误是否需要重�?
func (e *Executor) shouldRetry(err error) bool {
	// 使用 errors 包中的判断函�?
	return errors.IsRetriable(err)
}


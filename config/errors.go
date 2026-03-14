package config

import (
	"errors"
	"fmt"
)

var (
	// ErrMissingCorpID 缺少企业ID
	ErrMissingCorpID = errors.New("corpID is required")

	// ErrMissingCorpSecret 缺少应用凭证密钥
	ErrMissingCorpSecret = errors.New("agents configuration is required")

	// ErrInvalidTimeout 无效的超时时�?
	ErrInvalidTimeout = errors.New("timeout must be greater than 0")

	// ErrInvalidMaxRetries 无效的重试次�?
	ErrInvalidMaxRetries = errors.New("maxRetries must be greater than or equal to 0")
)

// ErrInvalidAgentConfig 无效的应用配�?
type ErrInvalidAgentConfig struct {
	AgentKey string
	Reason   string
}

func (e *ErrInvalidAgentConfig) Error() string {
	return fmt.Sprintf("invalid agent config for '%s': %s", e.AgentKey, e.Reason)
}


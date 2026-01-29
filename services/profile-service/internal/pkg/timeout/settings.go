package timeout

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

var clientTimeoutSettings atomic.Pointer[ServiceClientSettings]

type ServiceClientSettings struct {
	DefaultSettings ServiceTimeout            `yaml:"default"`
	Services        map[string]ServiceTimeout `yaml:"services"`
}

// ServiceTimeout Настройки на сервис
type ServiceTimeout struct {
	Timeout  time.Duration              `yaml:"timeout"`
	Handlers map[string]HandlerSettings `yaml:"handlers"`
}

// HandlerSettings Настройки на ручки
type HandlerSettings struct {
	Timeout time.Duration `yaml:"timeout"`
}

func getServiceClientTimeoutSettings() *ServiceClientSettings {
	settings := clientTimeoutSettings.Load()
	if settings == nil {
		return &ServiceClientSettings{}
	}

	return settings
}

func SetServiceClientTimeoutSettings(settings ServiceClientSettings) {
	clientTimeoutSettings.Store(&settings)
}

func newCtxWithTimeout(ctx context.Context, service, method string, settings *ServiceClientSettings) (context.Context, context.CancelFunc) {
	t, err := target(service, method, settings)

	if err != nil {
		return ctx, func() {}
	}

	return context.WithTimeout(ctx, t)
}

func target(targetServiceName, handler string, settings *ServiceClientSettings) (time.Duration, error) {
	if serviceTimeout, ok := settings.Services[targetServiceName]; ok {
		if handlerTimeout, handlerOk := serviceTimeout.Handlers[handler]; handlerOk {
			return handlerTimeout.Timeout, nil
		}

		// default path
		if serviceTimeout.Timeout != 0 {
			return serviceTimeout.Timeout, nil
		}
	}

	if settings.DefaultSettings.Timeout != 0 {
		return settings.DefaultSettings.Timeout, nil
	}

	return time.Duration(0), fmt.Errorf("timeout for service and handler not found")
}

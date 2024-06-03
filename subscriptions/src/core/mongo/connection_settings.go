package mongo

import (
	"fmt"
	"strings"
	"subscriptions/core/helpers"
	"time"
)

type ConnectionSettings struct {
	connectionString string
	timeout          time.Duration
	retryCount       int
	// sleepBetweenRetry The time to wait between retries in seconds
	sleepBetweenRetry time.Duration
	// retrySleepTimeout the maximum time in seconds allowed for each retry step
	retrySleepTimeout time.Duration
}

func (settings ConnectionSettings) GetConnectionString() string {

	var computed = settings.connectionString

	if strings.Contains(computed, "?") {
		computed += fmt.Sprintf("&connectTimeoutMS=%d", settings.timeout.Milliseconds())
	} else {
		if !helpers.EndsWith(computed, "/") {
			computed += "/"
		}
		computed += fmt.Sprintf("?connectTimeoutMS=%d", settings.timeout.Milliseconds())
	}

	if !strings.Contains(computed, "keepAlive") {
		computed += "&keepAlive=true"
	}
	if !strings.Contains(computed, "autoReconnect") {
		computed += "&autoReconnect=true"
	}
	if !strings.Contains(computed, "socketTimeoutMS") {
		computed += fmt.Sprintf("&socketTimeoutMS=%d", settings.timeout.Milliseconds())
	}

	return computed
}

func NewSettings(connectionString string) ConnectionSettings {
	return ConnectionSettings{
		connectionString: connectionString,
	}
}

func (settings ConnectionSettings) WithTimeout(duration time.Duration) ConnectionSettings {
	settings.timeout = duration
	return settings
}

func (settings ConnectionSettings) WithRetry(
	retryCount int,
	sleepBetweenRetry time.Duration,
	retrySleepTimeout time.Duration) ConnectionSettings {
	settings.retryCount = retryCount
	settings.sleepBetweenRetry = sleepBetweenRetry
	settings.retrySleepTimeout = retrySleepTimeout
	return settings
}

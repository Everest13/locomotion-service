package config

type configKey string

const (
	KafkaQuickstartEventsTopicName = "quickstart-events"
)

func GetValue(ctx context.Context, key configKey) realtimeconfig.Value {
	return getValue(config.Client(), ctx, key)
}

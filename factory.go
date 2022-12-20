package tracecomponent

import (
	"context"
	"strconv"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	typeStr         = "tracecomponent"
	defaultInterval = 1
)

func CreateDefaultConfig() component.Config {

	return &Config{
		ReceiverSettings: config.NewReceiverSettings(component.NewID(typeStr)),
		Interval:         strconv.Itoa(defaultInterval),
	}
}

// func CreateTracesReceiver(_ context.Context, params component.ReceiverCreateSettings, bcfg Config,
//
//	nextConsumer Component.consumer.Traces) (component.TracesReceiver, error){
//		return nil, nil
//	}
func createTracesReceiver(_ context.Context, p receiver.CreateSettings, cfg component.Config, nextconsumer consumer.Traces) (receiver.Traces, error) {
	if nextconsumer == nil{
		return nil, component.ErrNilNextConsumer
	}
	logger := p.Logger
	tracecomponentCfg := cfg.(*Config)

	traceRcvr := &tracecomponentReceiver{
		logger: logger,
		nextConsumer: nextconsumer,
		config: tracecomponentCfg,
	}

	return traceRcvr, nil
}

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		CreateDefaultConfig,
		receiver.WithTraces(createTracesReceiver, component.StabilityLevelStable),
	)

}
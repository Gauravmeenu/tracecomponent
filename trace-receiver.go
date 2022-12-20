package tracecomponent

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.uber.org/zap"
)

type tracecomponentReceiver struct{
	host component.Host
	cancel context.CancelFunc
	logger *zap.Logger
	nextConsumer consumer.Traces
	config *Config

}

func (tracecomponentRcvr *tracecomponentReceiver)Start(ctx context.Context, host component.Host)error{
	tracecomponentRcvr.host=host
	ctx=context.Background()
	ctx, tracecomponentRcvr.cancel=context.WithCancel(ctx)
	interval, _ := time.ParseDuration(tracecomponentRcvr.config.Interval)
	go func(){
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for{
			select{
			case <-ticker.C:
				tracecomponentRcvr.logger.Info("I should start processing now!!")
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (tracecomponentRcvr *tracecomponentReceiver)Shutdown(ctx context.Context)error{
	tracecomponentRcvr.cancel()
	return nil
}
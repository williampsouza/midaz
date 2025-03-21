package rabbitmq

import (
	"context"
	"github.com/LerianStudio/midaz/pkg"
	"github.com/LerianStudio/midaz/pkg/mlog"
	"github.com/LerianStudio/midaz/pkg/mopentelemetry"
	"github.com/LerianStudio/midaz/pkg/mrabbitmq"
	"github.com/LerianStudio/midaz/pkg/net/http"
)

const numWorkers = 5

// ConsumerRepository provides an interface for Consumer related to rabbitmq.
//
//go:generate mockgen --destination=consumer.mock.go --package=rabbitmq . ConsumerRepository
type ConsumerRepository interface {
	Register(queueName string, handler QueueHandlerFunc)
	RunConsumers() error
}

// QueueHandlerFunc is a function that process a specific queue.
type QueueHandlerFunc func(ctx context.Context, body []byte) error

// ConsumerRoutes struct
type ConsumerRoutes struct {
	conn   *mrabbitmq.RabbitMQConnection
	routes map[string]QueueHandlerFunc
	mlog.Logger
	mopentelemetry.Telemetry
}

// NewConsumerRoutes creates a new instance of ConsumerRoutes.
func NewConsumerRoutes(conn *mrabbitmq.RabbitMQConnection, logger mlog.Logger, telemetry *mopentelemetry.Telemetry) *ConsumerRoutes {
	cr := &ConsumerRoutes{
		conn:      conn,
		routes:    make(map[string]QueueHandlerFunc),
		Logger:    logger,
		Telemetry: *telemetry,
	}

	_, err := conn.GetNewConnect()
	if err != nil {
		panic("Failed to connect rabbitmq")
	}

	return cr
}

// Register add a new queue to handler.
func (cr *ConsumerRoutes) Register(queueName string, handler QueueHandlerFunc) {
	cr.routes[queueName] = handler
}

// RunConsumers  init consume for all registry queues.
func (cr *ConsumerRoutes) RunConsumers() error {
	for queueName, handler := range cr.routes {
		cr.Logger.Infof("Initializing consumer for queue: %s", queueName)

		err := cr.conn.Channel.Qos(
			10,
			0,
			false,
		)
		if err != nil {
			return err
		}

		messages, err := cr.conn.Channel.Consume(
			queueName,
			"",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}

		for i := 0; i < numWorkers; i++ {
			go func(workerID int, queue string, handlerFunc QueueHandlerFunc) {
				for msg := range messages {
					midazID, found := msg.Headers[http.HeaderMidazID]
					if !found {
						midazID = pkg.GenerateUUIDv7().String()
					}

					log := cr.Logger.WithFields(
						http.HeaderMidazID, midazID.(string),
					).WithDefaultMessageTemplate(midazID.(string) + " | ")

					ctx := pkg.ContextWithLogger(
						pkg.ContextWithMidazID(context.Background(), midazID.(string)),
						log,
					)

					err := handlerFunc(ctx, msg.Body)
					if err != nil {
						cr.Logger.Errorf("Worker %d: Error processing message from queue %s: %v", workerID, queue, err)

						_ = msg.Nack(false, true)

						continue
					}

					_ = msg.Ack(false)
				}
			}(i, queueName, handler)
		}
	}

	return nil
}

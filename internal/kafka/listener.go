/*
Copyright © 2022 António Inácio

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kafka

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

func NewListener(config ListenerConfig, consumer Consumer) *Listener {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Net.TLS.Enable = true
	saramaConfig.Net.SASL.Enable = true
	saramaConfig.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	saramaConfig.Net.SASL.User = config.Username
	saramaConfig.Net.SASL.Password = config.Password
	saramaConfig.ClientID = config.GroupId

	return &Listener{l: *zap.S(), saramaConfig: saramaConfig, consumer: &consumer, config: config}
}

func (l *Listener) Listen() {
	var log = l.l

	// We can use this to have more information but zap is not compatible with
	// Sarama's logger interface:
	// sarama.Logger = ...

	/**
	 * Set up a new Sarama consumerGroup group
	 */
	consumerGroup := ConsumerGroup{
		ready:    make(chan bool),
		consumer: *l.consumer,
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(l.config.BootstrapServer, ","), l.config.GroupId, l.saramaConfig)
	if err != nil {
		log.Error("Error creating consumerGroup group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumerGroup session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, l.config.Topics, &consumerGroup); err != nil {
				log.Error("Error from consumerGroup: %v", err)
			}
			// check if context was cancelled, signaling that the consumerGroup should stop
			if ctx.Err() != nil {
				return
			}
		}
	}()

	consumerGroup.WaitUntilReady()

	log.Info("Consumer up and running", "groupId", l.config.GroupId, "topics", l.config.Topics)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Info("Consumer terminating: context cancelled")
	case <-sigterm:
		log.Info("Consumer terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Error("Error closing client: %v", err)
	}
}

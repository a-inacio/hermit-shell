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
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

// Listener allows a Consumer handling kafka messages without much hassle
type Listener struct {
	l            zap.SugaredLogger
	saramaConfig *sarama.Config
	consumer     *Consumer
	config       ListenerConfig
}

// Consumer can freely consume messages, return Error if something goes wrong, nil otherwise.
type Consumer interface {
	Consume(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error
}

// ListenerConfig defines the settings for a Listener to be able to connect to a Kafka Broker.
type ListenerConfig struct {
	BootstrapServer string
	Topics          []string
	GroupId         string
	Username        string
	Password        string
}

// ConsumerGroup represents a Sarama consumer group consumer
type ConsumerGroup struct {
	ready    chan bool
	consumer Consumer
}

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
	"errors"
)

func NewConsumerChannel() *ConsumerChannel {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	return &ConsumerChannel{
		message: make(chan []byte, 1),
		ack:     make(chan bool, 1),
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (cc *ConsumerChannel) ConsumerWaitRead() ([]byte, error) {
	select {
	case <-cc.ctx.Done():
		return nil, errors.New("Operation completed")
	default:
		return <-cc.message, nil
	}
}

func (cc *ConsumerChannel) ConsumerAcknowledge() {
	cc.ack <- true
}

func (cc *ConsumerChannel) ConsumerCancel() {
	cc.cancel()
}

func (cc *ConsumerChannel) ProducerWrite(message []byte) error {
	select {
	case <-cc.ctx.Done():
		return errors.New("Operation completed")
	default:
		cc.message <- message
	}

	return nil
}

func (cc *ConsumerChannel) ProducerWaitAcknowledge() (bool, error) {
	select {
	case <-cc.ctx.Done():
		return false, errors.New("Operation completed")
	default:
		return <-cc.ack, nil
	}
}

func (cc *ConsumerChannel) WaitUntilFinished() {
	<-cc.ctx.Done()
	close(cc.message)
	close(cc.ack)
}

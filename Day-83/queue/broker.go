package queue

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type BrokerMessage struct {
	Offset    int64
	Topic     string
	Partition int
	Payload   string
}

// PartitionLog holds a single log stream partition.
type PartitionLog struct {
	mu         sync.RWMutex
	nextOffset int64
	messages   []BrokerMessage
}

// MessageBroker simulates a Kafka-style partitioned broker.
type MessageBroker struct {
	mu         sync.RWMutex
	partitions map[string]map[int]*PartitionLog
	offsets    map[string]int64 // consumerGroup -> offset
}

// NewMessageBroker initializes broker.
func NewMessageBroker() *MessageBroker {
	return &MessageBroker{
		partitions: make(map[string]map[int]*PartitionLog),
		offsets:    make(map[string]int64),
	}
}

// Publish appends a message to a specific topic and partition.
func (b *MessageBroker) Publish(topic string, partition int, payload string) int64 {
	b.mu.Lock()
	if _, exists := b.partitions[topic]; !exists {
		b.partitions[topic] = make(map[int]*PartitionLog)
	}

	pLog, exists := b.partitions[topic][partition]
	if !exists {
		pLog = &PartitionLog{messages: make([]BrokerMessage, 0)}
		b.partitions[topic][partition] = pLog
	}
	b.mu.Unlock()

	pLog.mu.Lock()
	defer pLog.mu.Unlock()

	offset := atomic.AddInt64(&pLog.nextOffset, 1) - 1
	msg := BrokerMessage{
		Offset:    offset,
		Topic:     topic,
		Partition: partition,
		Payload:   payload,
	}

	pLog.messages = append(pLog.messages, msg)
	return offset
}

// ConsumeGroup reads messages from a consumer group's saved offset position.
func (b *MessageBroker) ConsumeGroup(group string, topic string, partition int, limit int) []BrokerMessage {
	b.mu.Lock()
	startOffset := b.offsets[fmt.Sprintf("%s:%s:%d", group, topic, partition)]
	pLog, exists := b.partitions[topic][partition]
	b.mu.Unlock()

	if !exists {
		return nil
	}

	pLog.mu.RLock()
	defer pLog.mu.RUnlock()

	consumed := make([]BrokerMessage, 0)
	var lastOffset int64 = startOffset

	for _, msg := range pLog.messages {
		if msg.Offset >= startOffset {
			consumed = append(consumed, msg)
			lastOffset = msg.Offset + 1
			if len(consumed) >= limit {
				break
			}
		}
	}

	// Commit consumer group offset
	b.mu.Lock()
	b.offsets[fmt.Sprintf("%s:%s:%d", group, topic, partition)] = lastOffset
	b.mu.Unlock()

	return consumed
}

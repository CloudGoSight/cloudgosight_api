package mq

import (
	"encoding/gob"
	"github.com/CloudGoSight/cloudgosight_api/biz/aria2"
	"strconv"
	"sync"
	"time"
)

// Message 消息事件正文
type Message struct {
	// 消息触发者
	TriggeredBy string

	// 事件标识
	Event string

	// 消息正文
	Content interface{}
}

type CallbackFunc func(Message)

// MQ 消息队列
type MQ interface {
	aria2.Notifier

	// 发布一个消息
	Publish(string, Message)

	// 订阅一个消息主题
	Subscribe(string, int) <-chan Message

	// 订阅一个消息主题，注册触发回调函数
	SubscribeCallback(string, CallbackFunc)

	// 取消订阅一个消息主题
	Unsubscribe(string, <-chan Message)
}

type inMemoryMQ struct {
	// topic是key, value是chan数组，对应多个消费者
	topics    map[string][]chan Message
	callbacks map[string][]CallbackFunc
	sync.RWMutex
}

var GlobalMQ = NewMQ()

func NewMQ() MQ {
	return &inMemoryMQ{
		topics:    make(map[string][]chan Message),
		callbacks: make(map[string][]CallbackFunc),
	}
}

func init() {
	// gob 编解码
	// todo 改成sonic
	gob.Register(Message{})
	gob.Register([]aria2.Event{})
}

func (i *inMemoryMQ) Publish(topic string, message Message) {
	i.RLock()
	subscribersChan, okChan := i.topics[topic]
	subscribersCallback, okCallback := i.callbacks[topic]
	i.RUnlock()

	if okChan {
		// select example: https://gobyexample-cn.github.io/timeouts
		// select 默认处理第一个已准备好的接收操作
		go func(subscribersChan []chan Message) {
			for i := 0; i < len(subscribersChan); i++ {
				select {
				case subscribersChan[i] <- message:
				case <-time.After(time.Millisecond * 500):
				}
			}
		}(subscribersChan)

	}

	if okCallback {
		for i := 0; i < len(subscribersCallback); i++ {
			go subscribersCallback[i](message)
		}
	}
}

func (i *inMemoryMQ) Subscribe(topic string, buffer int) <-chan Message {
	ch := make(chan Message, buffer)
	i.Lock()
	i.topics[topic] = append(i.topics[topic], ch)
	i.Unlock()
	return ch
}

func (i *inMemoryMQ) SubscribeCallback(topic string, callbackFunc CallbackFunc) {
	i.Lock()
	i.callbacks[topic] = append(i.callbacks[topic], callbackFunc)
	i.Unlock()
}

// 将sub从消费者中移除，保留其他消费者
func (i *inMemoryMQ) Unsubscribe(topic string, sub <-chan Message) {
	i.Lock()
	defer i.Unlock()

	subscribers, ok := i.topics[topic]
	if !ok {
		return
	}

	var newSubs []chan Message
	for _, subscriber := range subscribers {
		if subscriber == sub {
			continue
		}
		newSubs = append(newSubs, subscriber)
	}

	i.topics[topic] = newSubs
}

func (i *inMemoryMQ) Aria2Notify(events []aria2.Event, status int) {
	for _, event := range events {
		i.Publish(event.Gid, Message{
			TriggeredBy: event.Gid,
			Event:       strconv.FormatInt(int64(status), 10),
			Content:     events,
		})
	}
}

// OnDownloadStart 下载开始
func (i *inMemoryMQ) OnDownloadStart(events []aria2.Event) {
	i.Aria2Notify(events, aria2.Downloading)
}

// OnDownloadPause 下载暂停
func (i *inMemoryMQ) OnDownloadPause(events []aria2.Event) {
	i.Aria2Notify(events, aria2.Paused)
}

// OnDownloadStop 下载停止
func (i *inMemoryMQ) OnDownloadStop(events []aria2.Event) {
	i.Aria2Notify(events, aria2.Canceled)
}

// OnDownloadComplete 下载完成
func (i *inMemoryMQ) OnDownloadComplete(events []aria2.Event) {
	i.Aria2Notify(events, aria2.Complete)
}

// OnDownloadError 下载出错
func (i *inMemoryMQ) OnDownloadError(events []aria2.Event) {
	i.Aria2Notify(events, aria2.Error)
}

// OnBtDownloadComplete BT下载完成
func (i *inMemoryMQ) OnBtDownloadComplete(events []aria2.Event) {
	i.Aria2Notify(events, aria2.Complete)
}

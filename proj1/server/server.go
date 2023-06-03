package server

import (
	"encoding/json"
	"fmt"
	"log"
	"proj1/feed"
	"proj1/queue"
	"proj1/semaphore"
	"runtime"
	"sync"
)

type CRUD_Response struct {
	Success bool `json:"success"`
	Id      int  `json:"id"`
}

type Get_Feed_Response struct {
	Id       int              `json:"id"`
	FeedList []*feed.JsonFeed `json:"feed"`
}

type Config struct {
	Encoder *json.Encoder // Represents the buffer to encode Responses
	Decoder *json.Decoder // Represents the buffer to decode Requests
	Mode    string        // Represents whether the server should execute
	// sequentially or in parallel
	// If Mode == "s"  then run the sequential version
	// If Mode == "p"  then run the parallel version
	// These are the only values for Version
	ConsumersCount int // Represents the number of consumers to spawn
}

func producer(msg_queue *queue.LockFreeQueue, config *Config, sem *semaphore.Semaphore, isDone *bool) {
	var msg queue.Request
	if err := config.Decoder.Decode(&msg); err != nil {
		log.Println((err))
	}

	if config.Mode == "p" {
		msg_queue.Enqueue(&msg)
		sem.Up()
	} else {
		msg_queue.Enqueue(&msg)
		// fmt.Println(msg)
	}

	if msg.Command == "DONE" {
		*isDone = true
	}
}

func consumer(msg_queue *queue.LockFreeQueue, user_feed feed.Feed, config *Config, sem *semaphore.Semaphore, wg *sync.WaitGroup, isDone *bool) {

	var msg *queue.Request

	for {
		if config.Mode == "p" {
			sem.Down()
			msg = msg_queue.Dequeue()
		} else {
			msg = msg_queue.Dequeue()
		}

		if msg == nil && config.Mode == "s" {
			break
		}

		if msg == nil && *isDone {
			break
		}

		if msg != nil {
			switch cmd := msg.Command; cmd {
			case "ADD":
				user_feed.Add(msg.Body, float64(msg.TimeStamp))
				msg_response := CRUD_Response{Success: true, Id: msg.Id}
				if err := config.Encoder.Encode(&(msg_response)); err != nil {
					log.Println(err)
				}
			case "REMOVE":
				if user_feed.Remove(float64(msg.TimeStamp)) {

					msg_response := CRUD_Response{Success: true, Id: msg.Id}
					if err := config.Encoder.Encode(&(msg_response)); err != nil {
						log.Println(err)
					}
				} else {

					msg_response := CRUD_Response{Success: false, Id: msg.Id}
					if err := config.Encoder.Encode(&(msg_response)); err != nil {
						log.Println(err)
					}
				}
			case "FEED":
				lst := user_feed.GetFeedList()
				msg_response := Get_Feed_Response{FeedList: lst, Id: msg.Id}

				if err := config.Encoder.Encode(&msg_response); err != nil {
					log.Println(err)
				}
			case "CONTAINS":
				if user_feed.Contains(float64(msg.TimeStamp)) {
					msg_response := CRUD_Response{Success: true, Id: msg.Id}

					if err := config.Encoder.Encode(&(msg_response)); err != nil {
						log.Println(err)
					}
				} else {
					msg_response := CRUD_Response{Success: false, Id: msg.Id}

					if err := config.Encoder.Encode(&(msg_response)); err != nil {
						log.Println(err)
					}
				}
			case "DONE":
				break
			default:
				fmt.Println("default")
			}
		}

	}

	if config.Mode == "p" {
		wg.Done()
	}

}

//Run starts up the twitter server based on the configuration
//information provided and only returns when the server is fully
// shutdown.
func Run(config Config) {
	feed := feed.NewFeed()
	msg_queue := queue.NewLockFreeQueue()
	sem := semaphore.NewSemaphore(0)
	var wg sync.WaitGroup
	isDone := false
	if config.Mode == "s" {
		for !isDone {
			producer(msg_queue, &config, sem, &isDone)
			consumer(msg_queue, feed, &config, sem, &wg, &isDone)
		}
	} else {

		for i := 0; i < config.ConsumersCount; i++ {
			wg.Add(1)
			go consumer(msg_queue, feed, &config, sem, &wg, &isDone)
		}

		for !isDone {
			producer(msg_queue, &config, sem, &isDone)
		}

		for runtime.NumGoroutine() > 1 {
			sem.Up()
		}
	}
}

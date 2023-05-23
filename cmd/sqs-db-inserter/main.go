package main

import (
	"log"
	"os"
	"time"

	"github.com/uvalib/virgo4-sqs-sdk/awssqs"
)

// main entry point
func main() {

	log.Printf("===> %s service starting up (version: %s) <===", os.Args[0], Version())

	// Get config params and use them to init service context. Any issues are fatal
	cfg := LoadConfiguration()

	log.Printf("[main] initializing SQS...")
	// load our AWS_SQS helper object
	sqs, err := awssqs.NewAwsSqs(awssqs.AwsSqsConfig{MessageBucketName: cfg.MessageBucketName})
	fatalIfError(err)

	log.Printf("[main] getting queue handle...")
	// get the queue handles from the queue name
	inQueueHandle, err := sqs.QueueHandle(cfg.InQueueName)
	fatalIfError(err)

	// access to the DB
	log.Printf("[main] creating database proxy...")
	dbProxy := NewDbProxy(*cfg)

	// create the record channel
	inboundMessageChan := make(chan awssqs.Message, cfg.WorkerQueueSize)

	// create counter object
	counter := Counter{}

	// start workers here
	log.Printf("[main] starting workers...")
	for w := 1; w <= cfg.Workers; w++ {
		go worker(w, dbProxy, sqs, inQueueHandle, inboundMessageChan, &counter)
	}

	log.Printf("[main] starting main polling loop...")
	for {

		// wait for a batch of messages
		messages, err := sqs.BatchMessageGet(inQueueHandle, awssqs.MAX_SQS_BLOCK_COUNT, time.Duration(cfg.PollTimeOut)*time.Second)
		if err != nil {
			log.Printf("ERROR: during message get (%s), sleeping and retrying", err.Error())

			// sleep for a while
			time.Sleep(1 * time.Second)

			// and try again
			continue
		}

		// did we receive any?
		sz := len(messages)
		if sz != 0 {

			for _, m := range messages {
				inboundMessageChan <- m
			}

		} else {
			log.Printf("[main] no new messages available")
			s, e := counter.Get()
			log.Printf("[main] since startup, processed %d messages (success: %d, error: %d)",
				s+e, s, e)
		}
	}
}

//
// end of file
//

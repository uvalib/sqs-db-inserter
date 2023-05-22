package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/uvalib/virgo4-sqs-sdk/awssqs"
)

// time to wait for inbound messages before doing something else
var waitTimeout = 5 * time.Second

func worker(workerId int, dbProxy *DbProxy, aws awssqs.AWS_SQS, queue awssqs.QueueHandle, inbound <-chan awssqs.Message, counter *Counter) {

	// keep a list of the messages processed
	queued := make([]awssqs.Message, 0, 1)
	var message awssqs.Message

	for {

		arrived := false

		// process a message or wait...
		select {
		case message = <-inbound:
			arrived = true

		case <-time.After(waitTimeout):
		}

		// we have an inbound message to process
		if arrived == true {

			// time the process
			start := time.Now()

			// add it to the queued list
			queued = append(queued, message)

			// decode the message and process it
			var pl map[string]interface{}
			err := json.Unmarshal(message.Payload, &pl)
			if err == nil {

				// should we ignore this message
				ignore := dbProxy.WillIgnore(pl)

				// un-ignored messages go to the database
				if ignore == false {
					// insert into the database
					err = dbProxy.Insert(pl)
				}

				if err == nil {
					// delete it from the inbound queue
					err = blockDelete(workerId, aws, queue, queued)
					fatalIfError(err)

					if ignore == false {
						duration := time.Since(start)
						log.Printf("INFO: worker %d: processed message in %d milliseconds", workerId, duration.Milliseconds())
						counter.AddSuccess(1)
					} else {
						log.Printf("INFO: worker %d: ignored message", workerId)
						counter.AddIgnore(1)
					}
				} else {
					log.Printf("worker %d: ERROR message failed to insert (%s) (%s)", workerId, err.Error(), string(message.Payload))
					counter.AddError(1)
				}
			} else {
				log.Printf("worker %d: ERROR message failed to unmarshal (%s) (%s)", workerId, err.Error(), string(message.Payload))
				counter.AddError(1)
			}

			// clear the queue
			queued = queued[:0]
		}
	}
}

func blockDelete(workerId int, aws awssqs.AWS_SQS, queue awssqs.QueueHandle, messages []awssqs.Message) error {

	// delete the block
	opStatus, err := aws.BatchMessageDelete(queue, messages)
	if err != nil {
		if err != awssqs.ErrOneOrMoreOperationsUnsuccessful {
			return err
		}
	}

	// did we fail
	if err == awssqs.ErrOneOrMoreOperationsUnsuccessful {
		for ix, op := range opStatus {
			if op == false {
				log.Printf("worker %d: WARNING message %d failed to delete", workerId, ix)
			}
		}
	}

	return nil
}

//
// end of file
//

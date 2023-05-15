package main

import (
	"log"
	"time"

	"github.com/uvalib/virgo4-sqs-sdk/awssqs"
)

// time to wait for inbound messages before doing something else
var waitTimeout = 5 * time.Second

func worker(workerId int, dbproxy *DbProxy, aws awssqs.AWS_SQS, queue awssqs.QueueHandle, inbound <-chan awssqs.Message) {

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

			// get the message identifier
			//id, found := message.GetAttribute(awssqs.AttributeKeyRecordId)
			//if found == false {
			//	id = "unknown"
			//	log.Printf("WARNING: cannot locate document id, using default")
			//}

			// add it to the queued list
			queued = append(queued, message)

			// process it

			// delete it from the inbound queue
			err := blockDelete(workerId, aws, queue, queued)
			fatalIfError(err)

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
				log.Printf("worker %d: ERROR message %d failed to delete", workerId, ix)
			}
		}
	}

	return nil
}

//
// end of file
//

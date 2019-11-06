package event

import (
	"log"
	"os"
)

type Control struct{}

type Event string
type Handler func()

type EventHandler struct {
	eventStream   chan Event
	controlStream chan Control
	handlers      map[Event][]Handler
	logger        *log.Logger
}

const SystemStarted = Event("system_started")
const Connected = Event("connected")

func (rcv *EventHandler) RegisterHandler(e Event, h Handler) {
	rcv.logger.Printf("register \"%s\" handler...\n", e)
	rcv.handlers[e] = append(rcv.handlers[e], h)
}

func (rcv *EventHandler) Send(e Event) {
	rcv.logger.Printf("send event: %s\n", e)
	rcv.eventStream <- e
}

func (rcv *EventHandler) Start() {
	rcv.logger.Println("start event handler...")
	go rcv.eventProcessing()
}

func (rcv *EventHandler) Stop() {
	rcv.logger.Println("stop event handler...")
	close(rcv.eventStream)
	<-rcv.controlStream
}

func (rcv *EventHandler) eventProcessing() {
	rcv.logger.Println("start event processing...")

	for e := range rcv.eventStream {
		for _, h := range rcv.handlers[e] {
			h()
		}
	}

	rcv.logger.Println("stop event processing")
	rcv.controlStream <- Control{}
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		make(chan Event),
		make(chan Control),
		make(map[Event][]Handler),
		log.New(os.Stdout, "[log] ", 0),
	}
}

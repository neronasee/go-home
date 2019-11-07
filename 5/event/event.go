package event

import (
	"log"
	"os"
)

type Event string
type Handler func()

type registerParams struct {
	e Event
	h Handler
}

type Control struct {
	t      string
	params registerParams
}

type EventHandler struct {
	eventStream   chan Event
	controlStream chan Control
	handlers      map[Event]Handler

	logger *log.Logger
}

const (
	SystemStarted = Event("system_started")
	Connected     = Event("connected")

	registerSignal = "register"
	stopSignal     = "stop"
)

func (rcv *EventHandler) RegisterHandler(e Event, h Handler) {
	rcv.logger.Printf("register \"%s\" handler...\n", e)
	rcv.controlStream <- Control{
		registerSignal,
		registerParams{e, h},
	}
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
	rcv.controlStream <- Control{t: stopSignal}
}

func (rcv *EventHandler) eventProcessing() {
	rcv.logger.Println("start event processing...")

	for {
		select {
		case c := <-rcv.controlStream:
			switch c.t {
			case registerSignal:
				rcv.handlers[c.params.e] = c.params.h
			case stopSignal:
				rcv.logger.Println("stop event processing")
				return
			default:
				rcv.logger.Println("Unknwon Control type")
			}
		case e := <-rcv.eventStream:
			handler, ok := rcv.handlers[e]

			if ok {
				handler()
			} else {
				rcv.logger.Printf("Error with %s", e)
			}
		}
	}
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		make(chan Event),
		make(chan Control),
		make(map[Event]Handler),

		log.New(os.Stdout, "[log] ", 0),
	}
}

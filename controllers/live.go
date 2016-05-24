package controllers

import (
	"container/list"
	"github.com/gorilla/websocket"
	"time"
	"encoding/json"
	"net/http"
	"github.com/astaxie/beego"
	_ "github.com/davecgh/go-spew/spew"
	"fmt"
	"strings"
)

type LiveController struct {
	MainController
}

/**
 * Register Event Types
 */
type EventData interface{}
type Event struct {
	event string
	Name 		string
	Email 		string
	Fired 		int
	Data 		[]EventData
}

/**
 * Register subscription and subscriber types
 */
type Subscriber struct {
	Email 		string
	Conn  		*websocket.Conn
}
type Subscription struct {
	Archive 	[]Event      // All the events from the archive.
	New     	<-chan Event // New events coming in.
}

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
	archiveSize = 20
)

/**
 * Create necessary global WS channels
 */
var (
	subscribe = make(chan Subscriber, 10)
	unsubscribe = make(chan string, 10)
	publish = make(chan Event, 10)
	waitingList = list.New()
	subscribers = list.New()
	archive = list.New()
)

/**
 * Basic Archive & Event functions
 */
func NewEvent(name string, email string, data []EventData) Event {
	return Event{"push", name, email, int(time.Now().Unix()), data}
}
func NewArchive(event Event) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}
func GetEvents(lastReceived int) []Event { // GetEvents returns all events after lastReceived.
	events := make([]Event, 0, archive.Len())
	for event := archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.Fired > int(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}

/**
 * Create join and leave funcs
 */
func Join(email string, ws *websocket.Conn) {
	fmt.Printf("User joined: %v", email)
	subscribe <- Subscriber{Email: email, Conn: ws}
}
func Leave(email string) {
	fmt.Printf("User left: %v", email)
	unsubscribe <- email
}
func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Email == user {
			return true
		}
	}
	return false
}

/**
 * Define a function to process all
 * incoming and outgoing requests from
 * our web sockets.
 */
func liveSockets(){
	for {
		select {
		case sub := <-subscribe:
			if !isUserExist(subscribers, sub.Email) {
				subscribers.PushBack(sub)
//				NewEvent := NewEvent("CLIENT.JOIN", sub.Email, []EventData{})
//				spew.Dump(NewEvent)
//				publish <- NewEvent
			}
		case event := <-publish:
			for ch := waitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				waitingList.Remove(ch)
			}
			BroadcastEvent(event)
			NewArchive(event)
		case unsub := <-unsubscribe:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Email == unsub {
					// Close connection.
					subscribers.Remove(sub)
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
					}
					break
				}
			}
		}
	}
}

func BroadcastEvent(event Event) {
	data, err := json.Marshal(event)
	if err != nil {
		return
	}
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		subEmail := strings.Trim(strings.ToLower(sub.Value.(Subscriber).Email), " ")
		eventEmail := strings.Trim(strings.ToLower(event.Email), " ")
		if subEmail == eventEmail{
			ws := sub.Value.(Subscriber).Conn
			if ws != nil {
				if ws.WriteMessage(websocket.TextMessage, data) != nil {
					unsubscribe <- sub.Value.(Subscriber).Email
				}
			}
		}
	}
}


/**
 * Define init function to run our liveSockets in a goroutine
 */
func init(){
	go liveSockets()
}

/**
 * Define the funcs that respond to particular routes.
 */
func (this *LiveController) Join(){
	// Verify user is logged in, get session data.
	sess := this.GetSession("go.mbitson.com")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}
	m := sess.(map[string]interface{})

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	// Join chat room.
	Join(m["username"].(string), ws)
	defer Leave(m["username"].(string))

	// Receive requests from clients.
	for {
		_, _, err := ws.ReadMessage() // SECOND VARIABLE _ IS REPLY
		if err != nil {
			return
		}
	}
}
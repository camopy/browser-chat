package websocket

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	service "github.com/camopy/browser-chat/app/application/service"
	submitmessage "github.com/camopy/browser-chat/app/application/usecase/submitMessage"
	"github.com/camopy/browser-chat/app/domain/entity"
	"github.com/camopy/browser-chat/app/domain/repository"
	"github.com/camopy/browser-chat/config"
	"github.com/gorilla/websocket"
)

type Websocket struct {
	conf        *config.Conf
	db          repository.ChatMessageRepository
	mediator    service.Mediator
	broadcaster service.Broadcaster
	upgrader    websocket.Upgrader
	clients     map[*websocket.Conn]bool
	mu          sync.Mutex
}

func New(db repository.ChatMessageRepository, mediator service.Mediator, broadcaster service.Broadcaster, conf *config.Conf) *Websocket {
	clients := make(map[*websocket.Conn]bool)
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &Websocket{
		conf:        conf,
		db:          db,
		mediator:    mediator,
		broadcaster: broadcaster,
		clients:     clients,
		upgrader:    upgrader,
	}
}

func (ws *Websocket) removeClient(client *websocket.Conn) {
	ws.mu.Lock()
	delete(ws.clients, client)
	ws.mu.Unlock()
}

func (websocket *Websocket) addClient(client *websocket.Conn) {
	websocket.mu.Lock()
	websocket.clients[client] = true
	websocket.mu.Unlock()
}

func (ws *Websocket) Start() error {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/websocket", ws.HandleConnections)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", ws.conf.Server.Port),
		Handler:      http.DefaultServeMux,
		ReadTimeout:  ws.conf.Server.TimeoutRead,
		WriteTimeout: ws.conf.Server.TimeoutWrite,
		IdleTimeout:  ws.conf.Server.TimeoutIdle,
	}

	fmt.Printf("Starting websocket at %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start websocket: %v", err)
	}
	return nil
}

func (websocket *Websocket) Stop() error {
	return nil
}

func (websocket *Websocket) HandleConnections(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling new connection from %s\n", r.RemoteAddr)
	ws, err := websocket.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	websocket.addClient(ws)

	for {
		var input submitmessage.Input
		err := ws.ReadJSON(&input)
		if err != nil {
			fmt.Println("error:", err)
			websocket.removeClient(ws)
			break
		}

		if input.Message != "" && input.UserName != "" {
			input.Time = time.Now()
			submit := submitmessage.New(websocket.mediator)
			submit.Execute(input)
		}
	}
}

func (websocket *Websocket) HandleMessages(broadcaster service.Broadcaster) {
	for {
		msg := broadcaster.Receive()
		for client := range websocket.clients {
			websocket.sendMessageToClient(client, msg)
		}
	}
}

func (websocket *Websocket) sendMessageToClient(client *websocket.Conn, msg *entity.ChatMessage) {
	err := client.WriteJSON(msg)
	if err != nil && unsafeError(err) {
		log.Printf("error: %v", err)
		client.Close()
		websocket.removeClient(client)
	}
}

func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}

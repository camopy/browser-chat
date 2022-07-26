package websocket

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	service "github.com/camopy/browser-chat/app/application/service"
	submitmessage "github.com/camopy/browser-chat/app/application/usecase/submitMessage"
	usersignup "github.com/camopy/browser-chat/app/application/usecase/userSignup"
	"github.com/camopy/browser-chat/app/domain/entity"
	"github.com/camopy/browser-chat/app/domain/repository"
	"github.com/camopy/browser-chat/app/infra/websocket/middleware"
	"github.com/camopy/browser-chat/app/util/validator"
	"github.com/camopy/browser-chat/config"
	"github.com/go-chi/chi"
	goval "github.com/go-playground/validator"
	"github.com/gorilla/websocket"
)

type Websocket struct {
	conf        config.ServerConf
	chatRepo    repository.ChatMessageRepository
	userRepo    repository.UserRepository
	mediator    service.Mediator
	broadcaster service.Broadcaster
	validator   *goval.Validate
	upgrader    websocket.Upgrader
	clients     map[*websocket.Conn]bool
	mu          sync.Mutex
}

func New(chatRepo repository.ChatMessageRepository, userRepo repository.UserRepository, mediator service.Mediator, broadcaster service.Broadcaster, conf config.ServerConf, validator *goval.Validate) *Websocket {
	clients := make(map[*websocket.Conn]bool)
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &Websocket{
		conf:        conf,
		chatRepo:    chatRepo,
		userRepo:    userRepo,
		mediator:    mediator,
		broadcaster: broadcaster,
		validator:   validator,
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
	r := chi.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("public")))
	r.HandleFunc("/websocket", ws.HandleConnections)
	r.With(middleware.ContentTypeJson).Post("/signup", ws.HandleSignUp)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", ws.conf.Port),
		Handler:      r,
		ReadTimeout:  ws.conf.TimeoutRead,
		WriteTimeout: ws.conf.TimeoutWrite,
		IdleTimeout:  ws.conf.TimeoutIdle,
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

func (websocket *Websocket) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	input := usersignup.Input{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Printf("error decoding form: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "%v"}`, err)
		return
	}

	if err := websocket.validator.Struct(input); err != nil {
		sendInvalidForm(w, err)
		return
	}

	u := usersignup.NewUserSignup(websocket.userRepo)
	output, err := u.Execute(input)
	if err != nil {
		log.Printf("error on sign up use case: %v", err)
		if err == repository.ErrUserNameAlreadyExists {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "%v"}`, err)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, err)
		return
	}

	if err := json.NewEncoder(w).Encode(output); err != nil {
		log.Printf("error encoding response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, err)
		return
	}
}

func sendInvalidForm(w http.ResponseWriter, err error) {
	log.Printf("error decoding form: %v", err)

	resp := validator.ToErrResponse(err)
	if resp == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, "form error response failure")
		return
	}

	respBody, err := json.Marshal(resp)
	if err != nil {
		log.Printf("error encoding response: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, "json creation failure")
		return
	}
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(respBody)
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

package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"sync/atomic"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/notnil/chess"
)

const (
	ColorWhite = "white"
	ColorBlack = "black"
)

type ChessHub struct {
	ChanNewUser         chan string
	ChanWaitingRoom     map[string]chan bool // key is UserID
	ChanRoomsUserJoined chan *ChessRoom
	ChanUserJoined      map[string]chan bool     // key is RoomID
	CountUserJoined     map[string]*atomic.Int32 // key is RoomID
	Rooms               map[string]*ChessRoom    // key is RoomID
	Clients             map[string]*ChessClient  // key is UserID
}

func NewChessHub() *ChessHub {
	return &ChessHub{
		ChanNewUser:         make(chan string),
		ChanWaitingRoom:     make(map[string]chan bool),
		ChanRoomsUserJoined: make(chan *ChessRoom),
		ChanUserJoined:      make(map[string]chan bool),
		CountUserJoined:     make(map[string]*atomic.Int32),
		Rooms:               make(map[string]*ChessRoom),
		Clients:             make(map[string]*ChessClient),
	}
}

func (c *ChessHub) RunWorkerNewUser() {
	queue := make([]string, 0)

	for userID := range c.ChanNewUser {
		queue = append(queue, userID)
		if len(queue) >= 2 {
			var (
				userID1    = queue[0]
				userID2    = queue[1]
				roomID     = uuid.NewString()
				user1Color = rand.Intn(2) == 0
				user2Color = !user1Color
			)
			queue = queue[2:]
			c.NewRoom(roomID)
			c.NotifyUserForPickedRoom(roomID, userID1, user1Color)
			c.NotifyUserForPickedRoom(roomID, userID2, user2Color)
		}
	}
}

func (c *ChessHub) RunWorkerUserJoined() {
	var wg sync.WaitGroup

	runWorkerUserJoined := func(roomID string, ch chan bool) {
		for range ch {
			count := c.CountUserJoined[roomID]
			count.Add(1)
			if count.Load() == 2 {
				c.Rooms[roomID].Game = chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))
				c.NotifyUsers(roomID)
				delete(c.CountUserJoined, roomID)
				return
			}
		}
		close(ch)
		wg.Done()
	}

	for room := range c.ChanRoomsUserJoined {
		wg.Add(1)
		go runWorkerUserJoined(room.ID, c.ChanUserJoined[room.ID])
	}

	wg.Wait()
}

func (c *ChessHub) NotifyUsers(roomID string) {
	for _, user := range c.Rooms[roomID].Clients {
		user.ChanNotifyWhenReady <- true
	}
}

func (c *ChessHub) NewRoom(roomID string) {
	c.Rooms[roomID] = &ChessRoom{
		ID:      roomID,
		Clients: make(map[string]*ChessClient),
	}
	c.ChanUserJoined[roomID] = make(chan bool)
	c.CountUserJoined[roomID] = &atomic.Int32{}

	c.ChanRoomsUserJoined <- c.Rooms[roomID]
}

func (c *ChessHub) NewUser(userID string) {
	c.ChanWaitingRoom[userID] = make(chan bool)
	c.ChanNewUser <- userID
}

func (c *ChessHub) UserJoined(userID string, conn *websocket.Conn) error {
	client, ok := c.Clients[userID]
	if !ok {
		// TODO
		// return err
		return nil
	}

	if client.ActiveConn != nil {
		// TODO
		// return err
		return nil
	}

	client.ActiveConn = conn
	client.ChanNotifyWhenReady = make(chan bool)
	c.ChanUserJoined[client.RoomID] <- true
	return nil
}

func (c *ChessHub) NotifyUserForPickedRoom(roomID, userID string, color bool) {
	c.Clients[userID] = &ChessClient{
		ID:     userID,
		RoomID: roomID,
	}
	c.Rooms[roomID].Clients[userID] = c.Clients[userID]

	switch color {
	case true:
		c.Clients[userID].Color = ColorWhite
	case false:
		c.Clients[userID].Color = ColorBlack
	}

	c.ChanWaitingRoom[userID] <- true
}

func (c *ChessHub) WaitForRoom(userID string) {
	if ch, ok := c.ChanWaitingRoom[userID]; ok {
		<-ch
		delete(c.ChanWaitingRoom, userID)
	}
}

func (c *ChessHub) WaitForOthers(userID string) {
	if client, ok := c.Clients[userID]; ok {
		<-client.ChanNotifyWhenReady
	}
}

func (c *ChessHub) StartGame(userID string) error {
	var (
		client     = c.Clients[userID]
		roomID     = client.RoomID
		movesOrder = []string{ColorWhite, ColorBlack}
		game       = c.Rooms[roomID].Game
		index      = 0
	)

	players, err := c.GetPlayers(roomID)
	if err != nil {
		return err
	}

	client.ActiveConn.WriteMessage(websocket.TextMessage, []byte(client.Color))

	for game.Outcome() == chess.NoOutcome {
		var (
			color         = movesOrder[index%2]
			oppositeColor = movesOrder[(index+1)%2]
		)

		mt, message, err := players[color].ActiveConn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			break
		}

		if err := game.MoveStr(string(message)); err != nil {
			players[color].ActiveConn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			continue
		}

		if players[oppositeColor].ActiveConn == nil {
			break
		}

		players[oppositeColor].ActiveConn.WriteMessage(websocket.TextMessage, message)
		index++
	}

	client.ActiveConn.WriteMessage(websocket.TextMessage, []byte(game.Outcome()))
	client.ActiveConn.WriteMessage(websocket.TextMessage, []byte(game.Method().String()))

	return nil
}

func (c *ChessHub) GetPlayers(roomID string) (map[string]*ChessClient, error) {
	players := make(map[string]*ChessClient)
	for _, v := range c.Rooms[roomID].Clients {
		players[v.Color] = v
	}
	return players, nil
}

type ChessRoom struct {
	ID      string
	Clients map[string]*ChessClient
	Game    *chess.Game
}

type ChessClient struct {
	ID                  string
	Color               string
	RoomID              string
	ActiveConn          *websocket.Conn
	ChanNotifyWhenReady chan bool
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	ChessHub *ChessHub
}

func NewServer(chessHub *ChessHub) *Server {
	return &Server{
		ChessHub: chessHub,
	}
}

func (s *Server) PickRoom(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	userID := uuid.NewString()

	s.ChessHub.NewUser(userID)
	s.ChessHub.WaitForRoom(userID)

	conn.WriteMessage(websocket.TextMessage, []byte(userID))
}

func (s *Server) JoinRoom(w http.ResponseWriter, r *http.Request) {
	var (
		clientID = mux.Vars(r)["client_id"]
	)

	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	defer func() {
		client := s.ChessHub.Clients[clientID]
		client.ActiveConn = nil
		delete(s.ChessHub.Clients, clientID)
		delete(s.ChessHub.Rooms, client.RoomID)
	}()

	s.ChessHub.UserJoined(clientID, conn)
	s.ChessHub.WaitForOthers(clientID)
	s.ChessHub.StartGame(clientID)
}

func main() {
	var (
		chessHub = NewChessHub()
		server   = NewServer(chessHub)
		router   = mux.NewRouter()

		port = getenv("PORT", "8080")
	)

	go chessHub.RunWorkerNewUser()
	go chessHub.RunWorkerUserJoined()

	router.HandleFunc("/rooms", server.PickRoom)
	router.HandleFunc("/rooms/{client_id}", server.JoinRoom)
	http.Handle("/", router)

	log.Printf("running chess server on port :%s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println(err)
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

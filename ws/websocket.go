package ws

// // Check gorilla's docs
// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// // Upgrades http to websocket conn
// func Upgrade(w gin.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
// 	// For checking client (origin of websocket conn)
// 	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
// 	conn, err := upgrader.Upgrade(w, r, nil)

// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}

// 	return conn, nil
// }

// // Creates a new pool, and initializes all the channels and maps.
// 	pool := websocket.NewPool()

// 	// ServeMux - Router
// 	// Socket functions are usually executed in go routines.
// 	go pool.Start()

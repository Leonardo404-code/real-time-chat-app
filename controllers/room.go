package controllers

// Room maintains the set of active users and broadcasts messages to the users.
type Room struct {
	users map[*User]bool

	broadcast chan []byte

	register chan *User

	unregister chan *User
}

func NewRoom() *Room {
	return &Room{
		broadcast:  make(chan []byte),
		register:   make(chan *User),
		unregister: make(chan *User),
		users:      make(map[*User]bool),
	}
}

func (r *Room) Run() {
	for {
		select {
		case user := <-r.register:
			r.users[user] = true
		case user := <-r.unregister:
			if _, ok := r.users[user]; ok {
				delete(r.users, user)
				close(user.send)
			}
		case message := <-r.broadcast:
			for user := range r.users {
				select {
				case user.send <- message:
				default:
					close(user.send)
					delete(r.users, user)
				}
			}
		}
	}
}

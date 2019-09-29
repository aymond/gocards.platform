package data

import "time"

// Game is the game state
type Game struct {
	ID        int
	UUID      string
	GameName  string
	UserID    int
	CreatedAt time.Time
}

// CreatedAtDate format the CreatedAt date to display nicely on the screen
func (game *Game) CreatedAtDate() string {
	return game.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// CreateGame creates a new game
func (user *User) CreateGame(gameName string) (game Game, err error) {
	statement := "insert into games (uuid, game_name, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, game_name, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), gameName, user.ID, time.Now()).Scan(&game.ID, &game.UUID, &game.GameName, &game.UserID, &game.CreatedAt)
	return
}

// Games returns all Games in the database
func Games() (games []Game, err error) {
	rows, err := Db.Query("SELECT id, uuid, game_name, user_id, created_at FROM games ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		game := Game{}
		if err = rows.Scan(&game.ID, &game.UUID, &game.GameName, &game.UserID, &game.CreatedAt); err != nil {
			return
		}
		games = append(games, game)
	}
	rows.Close()
	return
}

// GameByUUID Get a Game by the UUID
func GameByUUID(uuid string) (game Game, err error) {
	game = Game{}
	err = Db.QueryRow("SELECT id, uuid, game_name, user_id, created_at FROM threads WHERE uuid = $1", uuid).
		Scan(&game.ID, &game.UUID, &game.GameName, &game.UserID, &game.CreatedAt)
	return
}

// User gets the user who started this thread
func (game *Game) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", game.UserID).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return
}

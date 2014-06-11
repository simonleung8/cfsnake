package game

import (
	"strconv"
	"strings"
	"time"
)

const stepSpeed time.Duration = 7 //in 1/10th second

type Game struct {
	instanceID  string //ID to identify different intance in redis
	players     []player
	playersData []playerData
	glbMap      []string //map for players controlled by other instance
	started     bool     //game started?
}

type player struct {
	Name     string
	Snake    []string
	GameOver bool
}

type playerData struct {
	Direction int    //1(east),2,3,4(north)
	Token     string //ID for the player
}

func (g *Game) New() {
	g = &Game{"", []player{}, []playerData{}, []string{}, false}
}

func (g *Game) NewPlayer() string {
	tmpData := []string{"15,11", "14,11", "13,11", "12,11", "11,11", "10,11"}

	//TODO: block joining if game has started

	token := generateToken()
	g.players = append(g.players, player{"simon", tmpData, false})
	g.playersData = append(g.playersData, playerData{2, token})
	return token
}

func (g *Game) Start() {
	//TODO: any instances can start game, coordinate with redis.start(func game.start)
	g.started = true

	go g.StartMoving()
}

func (g *Game) StartMoving() {
	ticker := time.NewTicker(stepSpeed * (time.Second / 10))
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			//TODO: check redis, populate glbMap
			g.calcualteSteps() //move all snakes forward in this instance
			g.checkCollision()
			//TODO: combine local and global maps, send to redis
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func (g *Game) GetPlayersMap() []player {
	return g.players
}

func (g *Game) Direction(token string) int {
	index := g.findToken(token)
	if index == -1 {
		return -1
	} else {
		return g.playersData[index].Direction
	}
}

func (g *Game) SetDirection(token string, direction int) {
	index := g.findToken(token)
	if index != -1 {
		g.playersData[index].Direction = direction
	}
}

func (g *Game) calcualteSteps() {
	var head []string
	var x, y int64
	for pIndex, player := range g.players {

		if player.GameOver {
			continue //dead player can't move
		}

		for i := len(player.Snake) - 1; i > 0; i-- {
			player.Snake[i] = player.Snake[i-1]
		}
		head = strings.Split(player.Snake[0], ",")
		x, _ = strconv.ParseInt(head[0], 10, 0)
		y, _ = strconv.ParseInt(head[1], 10, 0)

		//move forward
		if g.playersData[pIndex].Direction == 1 {
			x++
			player.Snake[0] = strconv.FormatInt(x, 10) + "," + strconv.FormatInt(y, 10)
		} else if g.playersData[pIndex].Direction == 2 {
			y++
			player.Snake[0] = strconv.FormatInt(x, 10) + "," + strconv.FormatInt(y, 10)
		} else if g.playersData[pIndex].Direction == 3 {
			x--
			player.Snake[0] = strconv.FormatInt(x, 10) + "," + strconv.FormatInt(y, 10)
		} else if g.playersData[pIndex].Direction == 4 {
			y--
			player.Snake[0] = strconv.FormatInt(x, 10) + "," + strconv.FormatInt(y, 10)
		}
	}

}

func (g *Game) checkCollision() {
	var heads []string
	for _, player := range g.players {
		heads = append(heads, player.Snake[0])
	}

	// checking local map
	for i, player := range g.players { //loop through all players
		for _, body := range player.Snake { //scan through snake body of the player
			if index := containString(heads, body, i); index != -1 {
				g.players[index].GameOver = true
			}
		}
	}

	//TODO: check global map collision
}

func (g *Game) findToken(token string) int {
	for i := 0; i < len(g.playersData); i++ {
		if g.playersData[i].Token == token {
			return i
		}
	}
	return -1
}

func generateToken() string {
	return time.Now().String()
}

func containString(ary []string, str string, currentIndex int) int {
	for i, v := range ary {
		if i != currentIndex && v == str {
			return i
		}
	}
	return -1
}

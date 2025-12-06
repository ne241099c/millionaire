package game

import "fmt"

type Player struct {
	ID   string
	Hand []Card
	Rank int
}

type Game struct {
	Players []*Player // 参加プレイヤーのリスト

	TableCards   []Card // 現在場に出ているカード
	LastPlayerID string // 最後にカードを出した人のID

	TurnIndex    int  // 何番目の人のターンか
	IsRevolution bool // 革命中かどうか
	PassCount    int  // 連続パスの数
}

func NewGame() *Game {
	return &Game{
		Players:    make([]*Player, 0),
		TableCards: make([]Card, 0),
		TurnIndex:  0,
	}
}

func (g *Game) Join(playerID string) {
	for _, p := range g.Players {
		if p.ID == playerID {
			return
		}
	}

	newPlayer := &Player{
		ID:   playerID,
		Hand: make([]Card, 0),
	}
	g.Players = append(g.Players, newPlayer)
	fmt.Printf("ゲーム参加: %s (現在 %d 人)\n", playerID, len(g.Players))
}

func (g *Game) Start() {
	if len(g.Players) < 2 {
		fmt.Println("エラー: プレイヤーが足りません")
		return
	}

	deck := NewDeck(1)
	deck.Shuffle()

	playerIdx := 0
	for _, card := range deck {
		g.Players[playerIdx].Hand = append(g.Players[playerIdx].Hand, card)

		playerIdx++
		if playerIdx >= len(g.Players) {
			playerIdx = 0
		}
	}

	g.TurnIndex = 0
	g.IsRevolution = false
	g.TableCards = nil
	g.PassCount = 0

	fmt.Println("★ゲームスタート！カードを配りました")
}

func (g *Game) DebugPrint() {
	fmt.Println("=== 現在の状況 ===")
	if g.IsRevolution {
		fmt.Println("【 革命中！ 】")
	}
	fmt.Printf("現在のターン: %s さん\n", g.Players[g.TurnIndex].ID)
	fmt.Printf("場のカード: %v\n", g.TableCards)

	for _, p := range g.Players {
		fmt.Printf("- %s の手札 (%d枚): %v\n", p.ID, len(p.Hand), p.Hand)
	}
	fmt.Println("==================")
}

func (p *Player) hasCards(targetCards []Card) bool {
	// 手札の枚数チェック
	handMap := make(map[string]int)
	for _, card := range p.Hand {
		handMap[card.String()]++
	}

	// 出そうとしているカードがマップにあるか確認
	for _, c := range targetCards {
		if handMap[c.String()] > 0 {
			handMap[c.String()]--
		} else {
			return false // 持ってない
		}
	}
	return true
}

func (p *Player) removeCards(targetCards []Card) {
	var newHand []Card

	toRemove := make(map[string]int)
	for _, c := range targetCards {
		toRemove[c.String()]++
	}

	// // 手札を走査し、削除リストにないものだけ残す
	for _, c := range p.Hand {
		if count, ok := toRemove[c.String()]; ok && count > 0 {
			toRemove[c.String()]--
		} else {
			newHand = append(newHand, c)
		}
	}
	p.Hand = newHand
}

func (g *Game) PlayCard(playerID string, cards []Card) error {
	currentPlayer := g.Players[g.TurnIndex]
	if currentPlayer.ID != playerID {
		return fmt.Errorf("あなたのターンではありません")
	}

	if !currentPlayer.hasCards(cards) {
		return fmt.Errorf("持っていないカードが含まれています")
	}

	// ルール判定

	// 手札から削除
	currentPlayer.removeCards(cards)

	// 場に出す
	g.TableCards = cards
	g.LastPlayerID = playerID

	// パスカウントのセット
	g.PassCount = 0

	// 次のターンへ
	g.advanceTurn()

	fmt.Printf("★処理成功: %s が %v を出しました\n", playerID, cards)
	return nil
}

func (g *Game) advanceTurn() {
	g.TurnIndex++
	if g.TurnIndex >= len(g.Players) {
		g.TurnIndex = 0 // 一周したら最初の人へ
	}
}

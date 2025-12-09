package game

import "fmt"

type Player struct {
	ID   string
	Hand []Card
	Name string
	Rank int
}

type Game struct {
	Players         []*Player // 参加プレイヤーのリスト
	FinishedPlayers []*Player // 順位が決まったプレイヤーのリスト
	TableCards      []Card    // 現在場に出ているカード
	LastPlayerID    string    // 最後にカードを出した人のID
	TurnIndex       int       // 何番目の人のターンか
	IsRevolution    bool      // 革命中かどうか
	Is11Back        bool      // 11バック中かどうか
	PassCount       int       // 連続パスの数
	IsActive        bool      // ゲームがアクティブかどうか
}

type HandType int

const (
	HandTypeInvalid  HandType = iota // 無効
	HandTypeSingle                   // 単騎
	HandTypePair                     // ペア
	HandTypeSequence                 // 階段
)

func NewGame() *Game {
	return &Game{
		Players:         make([]*Player, 0),
		FinishedPlayers: make([]*Player, 0),
		TableCards:      make([]Card, 0),
		TurnIndex:       0,
	}
}

func (g *Game) Join(playerID string, name string) {
	for _, p := range g.Players {
		if p.ID == playerID {
			return
		}
	}

	newPlayer := &Player{
		ID:   playerID,
		Name: name,
		Hand: make([]Card, 0),
	}
	g.Players = append(g.Players, newPlayer)
	fmt.Printf("ゲーム参加: %s(%s) (現在 %d 人)\n", name, playerID, len(g.Players))
}

func (g *Game) Start() {
	if g.IsActive {
		fmt.Println("⚠️ ゲームは既に進行中です")
		return
	}

	if len(g.Players) < 2 {
		fmt.Println("エラー: プレイヤーが足りません")
		return
	}

	g.IsActive = true

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
	_, err := g.validatePlay(cards)
	if err != nil {
		return err
	}

	// 手札から削除
	currentPlayer.removeCards(cards)

	// 場に出す
	g.TableCards = cards
	g.LastPlayerID = playerID

	// パスカウントのセット
	g.PassCount = 0

	// 革命チェック
	if len(cards) >= 4 {
		g.IsRevolution = !g.IsRevolution
		fmt.Printf("★革命が起きました！ (Revolution: %v)\n", g.IsRevolution)
	}

	// 11バックチェック
	if isJack(cards) {
		g.Is11Back = true // Jを出すたびに通常と革命を行ったり来たりさせたい場合はg.Is11Back = !g.Is11Backにする
		fmt.Println("🃏 11バック! (強さが反転します)")
	}

	// あがり判定
	if len(currentPlayer.Hand) == 0 {
		// 順位リストに追加
		g.FinishedPlayers = append(g.FinishedPlayers, currentPlayer)
		// ランクセット
		currentPlayer.Rank = len(g.FinishedPlayers)

		fmt.Printf("🎉 おめでとう！ %s さんが %d 位で抜けました！\n", playerID, currentPlayer.Rank)

		//終了チェック
		if len(g.FinishedPlayers) >= len(g.Players)-1 {
			fmt.Println("🏁 ゲーム終了！")
		}
	}

	if isEight(cards) {
		fmt.Println("✂️ 8切り発生!場を流して、もう一度あなたの番です")
		g.clearTable()
		fmt.Printf("★処理成功: %s が %v を出しました(8切り)\n", playerID, cards)
		return nil
	}

	// 次のターンへ
	g.advanceTurn()

	fmt.Printf("★処理成功: %s が %v を出しました\n", playerID, cards)
	return nil
}

func (g *Game) validatePlay(cards []Card) (int, error) {
	// 役分析
	myType, myStr, err := g.analyzeHand(cards)
	if err != nil {
		return 0, err
	}

	// 場にカードがない場合
	if len(g.TableCards) == 0 {
		return myStr, nil
	}

	// 場にカードがある場合
	tableType, tableStr, _ := g.analyzeHand(g.TableCards)

	// 枚数チェック
	if len(cards) != len(g.TableCards) {
		return 0, fmt.Errorf("枚数が違います (場:%d枚 vs 出:%d枚)", len(g.TableCards), len(cards))
	}

	// 役チェック
	if myType != tableType {
		return 0, fmt.Errorf("役の種類が違います (場:%v vs 出:%v)", tableType, myType)
	}

	// スペ3チェック
	isMySpe3 := (len(cards) == 1 && cards[0].Suit == Spade && cards[0].Rank == Three)
	isTableJoker := (len(g.TableCards) == 1 && g.TableCards[0].Suit == Joker)
	if isMySpe3 && isTableJoker {
		fmt.Println("🗡️ スペ3発動!ジョーカーを討ち取りました")
		return 999, nil
	}

	// 強さチェック
	if myStr <= tableStr {
		return 0, fmt.Errorf("場のカードより弱いです")
	}

	return myStr, nil
}

func (g *Game) analyzeHand(cards []Card) (HandType, int, error) {
	effectiveRev := (g.IsRevolution != g.Is11Back)

	count := len(cards)
	if count == 0 {
		return HandTypeInvalid, 0, fmt.Errorf("カードがありません")
	}

	if count == 1 {
		return HandTypeSingle, GetStrength(cards[0], effectiveRev), nil
	}

	// 階段
	if IsSequence(cards) {
		maxStr := -999

		for _, c := range cards {
			if c.Suit == Joker {
				continue
			}
			s := GetStrength(c, effectiveRev)
			if s > maxStr {
				maxStr = s
			}
		}

		if maxStr == -999 {
			maxStr = GetStrength(cards[0], effectiveRev)
		}

		return HandTypeSequence, maxStr, nil
	}

	// ペア
	if IsPair(cards) {
		baseStr := -999
		for _, c := range cards {
			if c.Suit != Joker {
				baseStr = GetStrength(c, effectiveRev)
				break
			}
		}

		if baseStr == -999 {
			baseStr = GetStrength(cards[0], effectiveRev)
		}

		return HandTypePair, baseStr, nil
	}
	return HandTypeInvalid, 0, fmt.Errorf("役になっていません")
}

func (g *Game) advanceTurn() {
	for i := 0; i < len(g.Players); i++ {
		g.TurnIndex++
		if g.TurnIndex >= len(g.Players) {
			g.TurnIndex = 0
		}

		if len(g.Players[g.TurnIndex].Hand) > 0 {
			return
		}
	}
}

func (g *Game) Pass(playerID string) error {
	if g.Players[g.TurnIndex].ID != playerID {
		return fmt.Errorf("あなたのターンではありません")
	}

	g.PassCount++
	fmt.Printf("★パス: %s (連続パス %d 回)\n", playerID, g.PassCount)

	activePlayerCount := len(g.Players) - len(g.FinishedPlayers)

	// 全員パスしたかチェック
	if g.PassCount >= activePlayerCount-1 {
		fmt.Println("★場が流れました！次の親は最後にカードを出した人です")
		g.clearTable()

		g.setTurnToID(g.LastPlayerID)

		if len(g.Players[g.TurnIndex].Hand) == 0 {
			g.advanceTurn()
		}

	} else {
		g.advanceTurn()
	}

	return nil
}

func (g *Game) clearTable() {
	g.TableCards = nil
	g.PassCount = 0
	g.Is11Back = false
}

func (g *Game) setTurnToID(targetID string) {
	for i, p := range g.Players {
		if p.ID == targetID {
			g.TurnIndex = i
			return
		}
	}
}

func isEight(cards []Card) bool {
	for _, c := range cards {
		if c.Rank == Eight {
			return true
		}
	}

	return false
}

func isJack(cards []Card) bool {
	for _, c := range cards {
		if c.Rank == Jack {
			return true
		}
	}

	return false
}

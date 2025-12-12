import styles from './GameRoom.module.css';
import { Card } from '../../components/Card/Card';
import { useState } from 'react';

export const GameScreen = ({ gameState, roomID, username, onStart, onPlay, onPass, logout }) => {
    const [selectedCards, setSelectedCards] = useState([]);
    const [isDragOver, setIsDragOver] = useState(false);

    if (!gameState) {
        return <div className={styles.container}>データ待機中...</div>;
    }

    const hand = gameState.hand || [];
    const tableCards = gameState.table_cards || [];
    const isActive = gameState.is_active;
    const isMyTurn = gameState.is_my_turn;
    const winnerName = gameState.winner_name;
    const allPlayers = gameState.all_players || []

    const toggleCard = (card) => {
        setSelectedCards(prev => {
            const isSelected = prev.some(c => c.Suit === card.Suit && c.Rank === card.Rank);
            if (isSelected) {
                return prev.filter(c => !(c.Suit === card.Suit && c.Rank === card.Rank));
            } else {
                return [...prev, card];
            }
        });
    };

    const isSelected = (card) => {
        return selectedCards.some(c => c.Suit === card.Suit && c.Rank === card.Rank);
    };

    const handleDragStart = (e, card) => {
        if (!isSelected(card)) {
            setSelectedCards([card]);
        }
    };

    const handleDragOver = (e) => {
        e.preventDefault();
        setIsDragOver(true); // 見た目を変える
    };

    const handleDragLeave = () => {
        setIsDragOver(false);
    };

    const handleDrop = (e) => {
        e.preventDefault();
        setIsDragOver(false);

        // 何も選択していなければ無視
        if (selectedCards.length === 0) return;

        // カードを出す
        onPlay(selectedCards);
        setSelectedCards([]); // 選択解除
    };

    if (!isActive && winnerName) {
        return (
            <div className={styles.container}>
                <div className={styles.gameSet}>
                    <h1 style={{ color: '#E91E63', fontSize: '3rem' }}>🏆 GAME SET!</h1>
                    <h2>勝者: {winnerName}</h2>
                    <br />
                    <button
                        className={styles.button}
                        onClick={onStart}
                        style={{ fontSize: '1.2em', padding: '15px 30px' }}
                    >
                        もう一度遊ぶ
                    </button>
                    <br /><br />
                    <button onClick={logout}>退出する</button>
                </div>
            </div>
        );
    }

    return (
        <div className={styles.container}>
            <header className={styles.header}>
                <h1>Room: {roomID}</h1>
                <p>Player: {username}</p>

                <button
                    onClick={logout}
                    className={styles.logoutButton}
                >
                    退出する
                </button>

                {!isActive && (
                    <div style={{ margin: '10px 0' }}>
                        <button className={styles.button} onClick={onStart}>
                            ▶ ゲーム開始
                        </button>
                    </div>
                )}

                {isActive && (
                    <div style={{ margin: '10px 0' }}>
                        <button
                            onClick={onPass}
                            disabled={!isMyTurn}
                            className={styles.passButton}>
                            🛑 パス
                        </button>
                    </div>
                )}
            </header>

            <main>
                <div className={styles.handInfo}>
                    {allPlayers
                        .filter(p => p.name !== username) // 自分は除外
                        .map((p, i) => (
                            <div key={i} className={styles.handInfoItem}>
                                {/* 名前 */}
                                <div style={{fontWeight: 'bold', fontSize: '0.9em'}}>{p.name}</div>
                                
                                {/* 残り枚数アイコン */}
                                <div style={{fontSize: '2em'}}>🂠 {p.hand_count}</div>
                                
                                {/* 順位がついている場合 */}
                                {p.rank > 0 && (
                                    <div className={styles.handInfoRank}>
                                        {p.rank}位
                                    </div>
                                )}
                            </div>
                        ))}
                </div>
                <h3>テーブル</h3>
                <div
                    className={`${styles.tableArea} ${isDragOver ? styles.tableAreaActive : ''}`}
                    onDragOver={handleDragOver}
                    onDragLeave={handleDragLeave}
                    onDrop={handleDrop}
                >
                    {/* 場に出ているカードを表示 */}
                    {tableCards.length > 0 ? (
                        tableCards.map((card, i) => (
                            <Card
                                key={`table-${i}`}
                                card={card}
                                isSelected={false}
                            />
                        ))
                    ) : (
                        <span style={{ color: '#ddd', opacity: 0.5 }}>No Cards</span>
                    )}
                </div>

                <h3>あなたの手札 ({hand.length}枚)</h3>

                <div className={styles.handArea}>
                    {hand.length > 0 ? (
                        hand.map((card, index) => (
                            <Card
                                key={`${card.Suit}-${card.Rank}`}
                                card={card}
                                onClick={() => toggleCard(card)}
                                isSelected={isSelected(card)}
                                onDragStart={(e) => handleDragStart(e, card)}
                            />
                        ))
                    ) : (
                        <p className={styles.message}>手札がありません</p>
                    )}
                </div>

                {/* デバッグ用 */}
                <details className={styles.debug}>
                    <summary>内部データを見る</summary>
                    <pre>{JSON.stringify(gameState, null, 2)}</pre>
                </details>
            </main>
        </div>
    );
};
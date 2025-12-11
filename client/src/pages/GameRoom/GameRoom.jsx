import styles from './GameRoom.module.css';
import { Card } from '../../components/Card/Card';
import { useState } from 'react';

export const GameScreen = ({ gameState, roomID, username, onStart, onPlay, logout }) => {
    const [selectedCards, setSelectedCards] = useState([]);
    const [isDragOver, setIsDragOver] = useState(false);

    if (!gameState) {
        return <div className={styles.container}>データ待機中...</div>;
    }

    const hand = gameState.hand || [];
    const tableCards = gameState.table_cards || [];

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

    return (
        <div className={styles.container}>
            <header className={styles.header}>
                <h1>Room: {roomID}</h1>
                <p>Player: {username}</p>

                <button
                    onClick={logout}
                    style={{ marginLeft: '10px', padding: '5px 10px', fontSize: '12px', background: '#666', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer' }}
                >
                    退出する
                </button>

                <div style={{ margin: '10px 0' }}>
                    <button
                        className={styles.button}
                        onClick={onStart}
                    >
                        ▶ ゲーム開始
                    </button>
                </div>
            </header>

            <main>
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
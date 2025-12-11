import styles from './GameRoom.module.css';
import { Card } from '../../components/Card/Card';

export const GameScreen = ({ gameState, roomID, username }) => {
    if (!gameState) {
        return <div className={styles.container}>データ待機中...</div>;
    }

    const hand = gameState.hand || [];

    return (
        <div className={styles.container}>
            <header className={styles.header}>
                <h1>Room: {roomID}</h1>
                <p>Player: {username}</p>
            </header>

            <main>
                <h3>あなたの手札 ({hand.length}枚)</h3>

                <div className={styles.handArea}>
                    {hand.length > 0 ? (
                        hand.map((card, index) => (
                            <Card
                                key={`${card.Suit}-${card.Rank}`}
                                card={card}
                                // まだ機能はないのでログだけ出す
                                onClick={() => console.log("クリック:", card)}
                                isSelected={false}
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
import styles from './HandArea.module.css';
import { Card } from '../../../components/Card/Card';

export const HandArea = ({ hand, username, isMyTurn, toggleCard, isSelected, onDragStart }) => {
    return (
        <div>
            <div className={styles.playerBadge}>
                <span className={isMyTurn ? styles.nameTagActive : styles.nameTag}>
                    あなた: {username}
                </span>
                <span style={{ marginLeft: '10px' }}>手札: {hand.length}枚</span>
            </div>

            <div className={`${styles.handArea} ${!isMyTurn ? styles.handAreaDisabled : ''}`}>
                {hand.length > 0 ? (
                    hand.map((card) => (
                        <Card
                            key={`${card.Suit}-${card.Rank}`}
                            card={card}
                            onClick={() => toggleCard(card)}
                            isSelected={isSelected(card)}
                            onDragStart={(e) => onDragStart(e, card)}
                        />
                    ))
                ) : (
                    <p className={styles.message}>手札がありません</p>
                )}
            </div>
        </div>
    );
};
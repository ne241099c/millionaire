import React from 'react';
import styles from './OpponentArea.module.css';

export const OpponentArea = ({ allPlayers, currentTurnID, username, isActive }) => {
    const opponents = allPlayers.filter(p => p.name !== username);
    return (
        <div className={styles.container}>
            {opponents.map((p, i) => {
                const isTurn = isActive && (p.id === currentTurnID);

                const cardClass = `
                    ${styles.opponentCard} 
                    ${isTurn ? styles.activeTurn : ''}
                    ${p.rank > 0 ? styles.finished : ''}
                `;

                return (
                    <div key={i} className={cardClass}>
                        <div style={{ textAlign: 'center', marginTop: '5px' }}>
                            <div className={styles.handCount}>
                                🂠 {p.hand_count}
                            </div>
                            <div className={isTurn ? styles.nameActive : styles.name}>
                                {p.name}
                            </div>
                        </div>

                        {isTurn && (
                            <div className={styles.thinking}>Thinking...</div>
                        )}

                        {p.rank > 0 && (
                            <div className={styles.rankBadge}>
                                {p.rank}位
                            </div>
                        )}
                    </div>
                );
            })}
        </div>
    );
};
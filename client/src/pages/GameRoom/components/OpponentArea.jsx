import React from 'react';
import styles from './OpponentArea.module.css';

export const OpponentArea = ({ allPlayers, currentTurnID, username, isActive }) => {
    const players = allPlayers;
    return (
        <div className={styles.container}>
            {players.map((p, i) => {
                const isTurn = isActive && (p.id === currentTurnID);
                const isMe = p.name === username;

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
                                {isMe ? `You (${p.name})` : p.name}
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
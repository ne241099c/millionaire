import React from 'react';
import styles from './GameResult.module.css';

export const GameResult = ({ allPlayers, onReset, logout }) => {
    const ranking = [...(allPlayers || [])].sort((a, b) => {
        if (a.rank === 0) return 1;
        if (b.rank === 0) return -1;
        return a.rank - b.rank;
    });

    return (
        <div className={styles.container}>
            <div className={styles.gameSet}>
                <h1 style={{ color: '#E91E63', fontSize: '3rem' }}>🏆 GAME SET!</h1>
                <div style={{ margin: '20px 0', textAlign: 'left' }}>
                    {ranking.map((p, i) => {
                        const rankDisplay = p.rank === 0 ? '-' : `${p.rank}位`;
                        return (
                            <div key={i} style={styles.playerResult}>
                                <span>{rankDisplay}</span>
                                <span style={{ fontWeight: 'bold' }}>{p.name}</span>
                            </div>
                        )
                    })}
                </div>
                <br />
                <button
                    className={styles.button}
                    onClick={onReset}
                    style={{ fontSize: '1.2em', padding: '15px 30px' }}
                >
                    ロビーへ戻る
                </button>
                <br /><br />
                <button onClick={logout}>退出する</button>
            </div>
        </div>
    );
};

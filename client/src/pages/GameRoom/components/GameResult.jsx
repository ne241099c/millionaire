import React from 'react';
import styles from './GameResult.module.css';

export const GameResult = ({ winnerName, onStart, logout }) => {
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
};

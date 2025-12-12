import React from 'react';
import styles from './GameHeader.module.css';

export const GameHeader = ({ roomID, username, isActive, isMyTurn, isRevolution, onStart, onPass, logout }) => {
    return (
        <header className={styles.header}>
            <div className={styles.roomInfo}>
                <h1>Room: {roomID}</h1>
                <span className={styles.playerInfo}>Player: {username}</span>
            </div>

            {isRevolution && (
                <div className={styles.revolutionBadge}>
                    ⚠️ 革命中 ⚠️
                </div>
            )}

            <div className={styles.actionArea}>
                {/* 開始ボタン */}
                {!isActive && (
                    <button
                        className={`${styles.button} ${styles.startButton}`}
                        onClick={onStart}
                    >
                        ▶ ゲーム開始
                    </button>
                )}

                {/* パスボタン */}
                {isActive && (
                    <button
                        onClick={onPass}
                        disabled={!isMyTurn}
                        className={`
                            ${styles.button} 
                            ${isMyTurn ? styles.passButton : styles.passButtonDisabled}
                        `}
                    >
                        パス
                    </button>
                )}

                {/* 退出ボタン */}
                <button
                    className={`${styles.button} ${styles.logoutButton}`}
                    onClick={logout}
                >
                    退出
                </button>
            </div>
        </header>
    );
};
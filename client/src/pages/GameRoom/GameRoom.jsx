import { useState } from 'react';
import styles from './GameRoom.module.css';

import { GameHeader } from './components/GameHeader';
import { GameResult } from './components/GameResult';
import { OpponentArea } from './components/OpponentArea';
import { TableArea } from './components/TableArea';
import { HandArea } from './components/HandArea';

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
    const allPlayers = gameState.all_players || [];
    const currentTurnID = gameState.current_turn_id;
    const isRevolution = isActive && gameState.is_revolution;;
    const effectiveMyTurn = isActive && isMyTurn;

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
            <GameResult
                winnerName={winnerName}
                onStart={onStart}
                logout={logout}
            />
        );
    }

    const containerClass = `
        ${styles.container} 
        ${isRevolution ? styles.revolution : ''}
        ${effectiveMyTurn ? styles.myTurn : ''}
    `;

    return (
        <div className={containerClass}>
            {/* ヘッダー: 部屋情報、ボタン類 */}
            <GameHeader
                roomID={roomID}
                username={username}
                isActive={isActive}
                isMyTurn={effectiveMyTurn}
                isRevolution={isRevolution}
                onStart={onStart}
                onPass={onPass}
                logout={logout}
            />
            <main>
                {/* 相手エリア: 画面上部に並ぶ */}
                <OpponentArea
                    allPlayers={allPlayers}
                    currentTurnID={currentTurnID}
                    username={username}
                    isActive={isActive}
                />

                {/* テーブルエリア: カードを出す場所 */}
                <TableArea
                    tableCards={tableCards}
                    isDragOver={isDragOver}
                    onDragOver={handleDragOver}
                    onDragLeave={handleDragLeave}
                    onDrop={handleDrop}
                />

                {/* 手札エリア: 自分のカード */}
                <HandArea
                    hand={hand}
                    username={username}
                    isMyTurn={effectiveMyTurn}
                    toggleCard={toggleCard}
                    isSelected={isSelected}
                    onDragStart={handleDragStart}
                />

                {/* デバッグ用 */}
                <details className={styles.debug}>
                    <summary>内部データを見る</summary>
                    <pre>{JSON.stringify(gameState, null, 2)}</pre>
                </details>
            </main>
        </div>
    );
};
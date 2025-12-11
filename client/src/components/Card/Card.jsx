import styles from './Card.module.css';

export const Card = ({ card, onClick, isSelected }) => {
    const suits = ["♠", "♥", "♦", "♣", "Joker"];
    const ranks = ["", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"];

    const isRed = card.Suit === 1 || card.Suit === 2; // Hearts or Diamonds

    const suitStr = suits[card.suit];
    const rankStr = card.Suit === 4 ? "" : ranks[card.Rank];

    const classList = [
        styles.card,                        // 基本スタイル
        isRed ? styles.red : styles.black,  // 赤か黒か
        isSelected ? styles.selected : ''   // 選択中か
    ].join(' ');

    return (
        <div
            onClick={onClick}
            className={classList}
        >
            <div>{suitStr}</div>
            <div>{rankStr}</div>
        </div>
    );
};
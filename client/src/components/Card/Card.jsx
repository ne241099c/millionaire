export const Card = ({ card, onClick, isSelected }) => {
    const suits = ["♠", "♥", "♦", "♣", "Joker"];
    const ranks = ["", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"];

    const isRed = card.suit === 1 || card.suit === 2; // Hearts or Diamonds

    const suitStr = suits[card.suit];
    const rankStr = rankStr = card.Suit === 4 ? "" : ranks[card.Rank];

    const classList = [
        styles.card,                        // 基本スタイル
        isRed ? styles.red : styles.black,  // 赤か黒か
        isSelected ? styles.selected : ''   // 選択中か
    ].join(' ');

    return (
        <div
            onClick={onClick}
            className="red-text"
        >
            <div>{suitStr}</div>
            <div>{rankStr}</div>
        </div>
    );
};
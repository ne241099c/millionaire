export default function DebugRoom() {
    return (
        <>
            <h1>♠ 大富豪 デバッグルーム ♦</h1>

            <div styleClass="margin-bottom: 10px;">
                <button id="btn-start" >▶ ゲーム開始</button>
                
                <button id="btn-pass"  styleClass="background-color: #ffccbc; display: none;">🛑 パス</button>
                
                <button >🃏 カードを出す</button>
            </div>

            <div id="status-bar" styleClass="padding: 10px; margin-bottom: 10px; font-weight: bold; font-size: 1.2em; border: 2px solid #ccc; background: #fff;">
                待機中...
            </div>

            <div id="game-area">
                <h3>場 (Table): <span id="table-cards">なし</span></h3>
                <hr/>
                <h3>あなたの手札 (Hand): <span id="hand-count">0</span>枚</h3>
                <div id="hand-container">
                    (ゲームを開始してください)
                </div>
            </div>

            <div id="log"></div>
        </>
    );
}
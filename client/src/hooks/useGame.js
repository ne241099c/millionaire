import DebugRoom from "../pages/DebugRoom"
import { useState, useEffect, useRef, use } from "react";

export const useGame = () => {
    // 接続状態
    const [isConnected, setIsConnected] = useState(false)

    // ゲーム状態
    const [gameState, setGameState] = useState(null)

    // 入室チェック中かどうか
    const [isEntry, setIsEntry] = useState(true);

    // WebSocket接続
    const socketRef = useRef(null)

    useEffect(() => {
        const savedName = sessionStorage.getItem("poker_name");
        const savedRoom = sessionStorage.getItem("poker_room");

        if (savedName && savedRoom && !socketRef.current) {
            console.log("🔄 前回のセッションから復帰します...");
            setTimeout(() => {
                connect(savedName, savedRoom);
                setIsEntry(false); // チェック完了
            }, 500);
            connect(savedName, savedRoom);
        } else {
            setIsEntry(false); // チェック完了
        }
    }, []);

    const connect = (name, roomID) => {
        if (!name) {
            alert("名前を入力してください")
            return
        }

        if (socketRef.current && socketRef.current.readyState === WebSocket.OPEN) {
            console.log("すでに接続済みです");
            return;
        }

        sessionStorage.setItem("poker_name", name);
        sessionStorage.setItem("poker_room", roomID);

        const baseUrl = import.meta.env.VITE_WS_URL
        if (!baseUrl) {
            console.error("設定エラー: VITE_WS_URL が見つかりません")
            return
        }

        const wsUrl = `${baseUrl}?room=${roomID}&name=${encodeURIComponent(name)}`
        console.log("接続開始:", wsUrl)

        const ws = new WebSocket(wsUrl)
        socketRef.current = ws;

        // 接続成功時の処理
        ws.onopen = () => {
            console.log("✅ サーバーに繋がりました")
            setIsConnected(true) // 画面をゲームモードに切り替え
        };

        ws.onmessage = (event) => {
            // JSON文字データを、JSのオブジェクトに変換
            const msg = JSON.parse(event.data)

            // ゲームの状態データなら保存する
            if (msg.type === "game_status") {
                setGameState(msg.payload)
            }
        };

        // 切断されたときの処理
        ws.onclose = () => {
            if (socketRef.current !== ws) {
                console.log("古い接続の切断を無視しました");
                return;
            }

            console.log("❌ 切断されました")
            setIsConnected(false) // ログイン画面に戻す
            setGameState(null)
            socketRef.current = null;
        };

        socketRef.current = ws
    };

    const startGame = () => {
        if (!socketRef.current) return;

        const msg = {
            type: "start_game", // Go側の MsgStartGame に対応
            payload: {}
        };

        socketRef.current.send(JSON.stringify(msg));
        console.log("📤 ゲーム開始リクエストを送信しました");
    };

    const playCards = (cards) => {
        if (!socketRef.current) return;

        const payload = {
            cards: cards
        };

        const msg = {
            type: "play_card",
            payload: payload
        };

        socketRef.current.send(JSON.stringify(msg));
        console.log("📤 カードを送信:", cards);
    };

    const logout = () => {
        // セッション情報を削除
        sessionStorage.removeItem("poker_name");
        sessionStorage.removeItem("poker_room");
        
        // ソケットを切断
        if (socketRef.current) {
            socketRef.current.close();
        }
        // 画面をリセット
        setIsConnected(false);
        setGameState(null);
    };

    // 片付け
    useEffect(() => {
        return () => {
            if (socketRef.current) {
                socketRef.current.close()
            }
        }
    }, [])

    return {
        isConnected,
        gameState,
        isEntry,
        connect,
        startGame,
        playCards,
    };
};
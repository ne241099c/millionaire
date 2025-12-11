import DebugRoom from "../pages/DebugRoom"
import { useState, useEffect, useRef } from "react";

export const useGame = () => {
  // 接続状態
  const [isConnected, setIsConnected] = useState(false)

  // ゲーム状態
  const [gameState, setGameState] = useState(null)

  // WebSocket接続
  const socketRef = useRef(null)

  const connect = (name, roomID) => {
    if (!name) {
      alert("名前を入力してください")
      return
    }

    const baseUrl = import.meta.env.VITE_WS_URL
    if (!baseUrl) {
      console.error("設定エラー: VITE_WS_URL が見つかりません")
      return
    }

    const wsUrl = `${baseUrl}?room=${roomID}&name=${encodeURIComponent(name)}`
    console.log("接続開始:", wsUrl)

    const ws = new WebSocket(wsUrl)

  // 接続成功時の処理
    ws.onopen = () => {
      console.log("✅ サーバーに繋がりました")
      setIsConnected(true) // 画面をゲームモードに切り替え
    }

    ws.onmessage = (event) => {
      // JSON文字データを、JSのオブジェクトに変換
      const msg = JSON.parse(event.data)
      
      // ゲームの状態データなら保存する
      if (msg.type === "game_status") {
        setGameState(msg.payload)
      }
    }

  // 切断されたときの処理
  ws.onclose = () => {
    console.log("❌ 切断されました")
    setIsConnected(false) // ログイン画面に戻す
    setGameState(null)
  }
  
  socketRef.current = ws
  }

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
    connect,
  };
};
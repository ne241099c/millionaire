import { useState } from 'react';

// Appから接続関数を受け取る
export const LoginScreen = ({ onJoin }) => {
  const [username, setUsername] = useState('');
  const [roomID, setRoomID] = useState('default');

  const handleJoin = () => {
    if (!username) {
      alert("名前を入力してください");
      return;
    }
    onJoin(username, roomID);
  };

  return (
    <div className="container">
      <h1>🃏 大富豪 Online</h1>
      <div style={{ display: 'flex', flexDirection: 'column', gap: '10px' }}>
        <input 
          type="text" 
          placeholder="ユーザー名 (必須)" 
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <input 
          type="text" 
          placeholder="部屋ID (default)" 
          value={roomID}
          onChange={(e) => setRoomID(e.target.value)}
        />
        <button onClick={handleJoin}>ゲームに参加</button>
      </div>
    </div>
  );
};
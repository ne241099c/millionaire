import { useRef, useState } from 'react'
import { useGame } from './hooks/useGame'
import { LoginScreen } from './pages/Login'
import { GameScreen } from './pages/GameRoom/GameRoom'
import './App.css'
import DebugRoom from './pages/DebugRoom'

function App() {
  const { isConnected, gameState, isEntry, connect, startGame, playCards, logout} = useGame();
  const [currentUser, setCurrentUser] = useState({ name: '', room: '' });

  const handleJoin = (name, room) => {
    setCurrentUser({ name, room });
    connect(name, room);
  };

  if (isEntry) {
    return (
      <div className="container" style={{marginTop: '50px'}}>
        <h2>🔄 復帰中...</h2>
        <p>サーバーと通信しています</p>
      </div>
    );
  }
  
  if (!isConnected) {
    return <LoginScreen onJoin={handleJoin} />;
  }

  return (
    <GameScreen
      gameState={gameState}
      username={currentUser.name}
      roomID={currentUser.room}
      onStart={startGame}
      onPlay={playCards}
      logout={logout}
    />
  );

  // return (
  //   <>
  //     <DebugRoom></DebugRoom>
  //   </>
  // )
}

export default App

import { useRef, useState } from 'react'
import { useGame } from './hooks/useGame'
import { LoginScreen } from './pages/Login'
import { GameScreen } from './pages/GameRoom/GameRoom'
import './App.css'
import DebugRoom from './pages/DebugRoom'

function App() {
  const { isConnected, gameState, connect, startGame, playCards} = useGame();
  const [currentUser, setCurrentUser] = useState({ name: '', room: '' });

  const handleJoin = (name, room) => {
    setCurrentUser({ name, room });
    connect(name, room);
  };

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
    />
  );

  // return (
  //   <>
  //     <DebugRoom></DebugRoom>
  //   </>
  // )
}

export default App

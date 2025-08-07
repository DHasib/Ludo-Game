import React, { useEffect, useState } from 'react';

export default function GameBoard({ token }) {
  const [log, setLog] = useState([]);

  useEffect(() => {
    const ws = new WebSocket(`ws://${window.location.host}/ws`);
    ws.onmessage = e => setLog(l => [...l, e.data]);
    return () => ws.close();
  }, []);

  return (
    <div>
      <h2>Ludo Game</h2>
      <ul>
        {log.map((l, i) => <li key={i}>{l}</li>)}
      </ul>
    </div>
  );
}

import WaterJugInputProps from './types/waterJugInput';
import Form from './components/Form';
import { useState } from 'react';

function App() {

  const [input, setInput] = useState<WaterJugInputProps | null>(null);

  const solveChallenge = async () => {
    const response = await fetch('http://localhost:8080/', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({x_cap: 2, y_cap: 10, target: 4})
    }).then(function(response) {
      return response.json();
    })
    .then(function(myJson) {
      console.log(JSON.stringify(myJson));
    //data = myJson;
    });
}

  return (
    <>
      <div>
       <h1>Water Jug Challenge</h1>
          <Form solveChallenge={solveChallenge}/>
      </div>
    </>
  )
}

export default App

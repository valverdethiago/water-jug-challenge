import { SolutionStep } from './types/solutionStep';
import { ErrorMessage } from './types/error';
import Form from './components/Form';
import StepList from './components/StepList';
import { useState } from 'react';
import classes from './App.module.css';

function App() {
  const [steps, setSteps] = useState<SolutionStep[]>([]);
  console.log(import.meta.env.VITE_BACKEND_URL);

  
  const solveChallenge = async (x_cap:number, y_cap:number, target: number) => {
    await fetch(import.meta.env.VITE_BACKEND_URL, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({x_cap: x_cap, y_cap: y_cap, target: target})
    }).then(function(response) {
      if(response.ok) {
        return response.json();
      } else {
        return response.json().then(text => { 
          const errorMessage = JSON.parse(JSON.stringify(text)) as ErrorMessage
          throw new Error(errorMessage.Message); 
        })
      }
    })
    .then(function(json) {
      const parsedSteps = JSON.parse(JSON.stringify(json)) as SolutionStep[];
      setSteps(parsedSteps);
    })
    .catch(function(error) {
      console.log(error.message);
    });
}

  return (
    <>
      <div className={classes.app}>
       <h1>Water Jug Challenge</h1>
          <Form solveChallenge={solveChallenge}/>
          <StepList steps={steps} />
      </div>
    </>
  )
}

export default App

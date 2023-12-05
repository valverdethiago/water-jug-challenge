import { SolutionStep } from './types/solutionStep';
import Form from './components/Form';
import StepList from './components/StepList';
import { useState } from 'react';
import classes from './App.module.css';

function App() {
  const [steps, setSteps] = useState<SolutionStep[]>([]);

  
  const solveChallenge = async (x_cap:number, y_cap:number, target: number) => {
    const response = await fetch('http://localhost:8080/', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({x_cap: x_cap, y_cap: y_cap, target: target})
    }).then(function(response) {
      return response.json();
    })
    .then(function(json) {
      console.log(json);
      console.log(JSON.stringify(json));
      const parsedSteps = JSON.parse(JSON.stringify(json)) as SolutionStep[];
      setSteps(parsedSteps);
      console.log(parsedSteps);
    //data = myJson;
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

import { useState } from "react"
import classes from './Form.module.css';

type FormProps = { 
    solveChallenge: (x_cap:number, y_cap:number, target: number) => Promise<void>;
}


const Form = ( {solveChallenge}: FormProps) => {

    const [x_cap, setX_cap] = useState<number>(-1);
    const [y_cap, setY_cap] = useState<number>(-1);
    const [target, setTarget] = useState<number>(-1);
    return (
        <div className={classes.challenge_form}>
            <h2>Please inform the values to do the calculation</h2>
            <div className={classes.challenge_form_container}>
                <input 
                    type="text"
                    placeholder="X Capacity"
                    onChange={(e) => setX_cap(+e.target.value)} />
                <input 
                    type="text"
                    placeholder="Y Capacity"
                    onChange={(e) => setY_cap(+e.target.value)} />
                <input 
                    type="text"
                    placeholder="Target"
                    onChange={(e) => setTarget(+e.target.value)} />
            </div>
            <button onClick={() => solveChallenge(x_cap, y_cap, target)}>Solve IT</button>
        </div>
    )
}

export default Form;
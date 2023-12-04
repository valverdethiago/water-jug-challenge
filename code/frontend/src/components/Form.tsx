import { useState } from "react"
type FormProps = { 
    solveChallenge: (x_cap:number, y_cap:number, target: number) => Promise<void>;
}


const Form = ( {solveChallenge}: FormProps) => {

    const [x_cap, setX_cap] = useState<number>(2);
    const [y_cap, setY_cap] = useState<number>(10);
    const [target, setTarget] = useState<number>(4);
    return (
        <div>
            <h2>Search</h2>
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
            <button onClick={() => solveChallenge(x_cap, y_cap, target)}>Solve IT</button>
        </div>
    )
}

export default Form;
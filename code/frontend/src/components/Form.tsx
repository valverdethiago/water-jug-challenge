type FormProps = { 
    solveChallenge: () => Promise<void>;
}

const Form = ( {solveChallenge}: FormProps) => {
    return (
        <div>
            <h2>Search</h2>
            <button onClick={() => solveChallenge()}>Solve IT</button>
        </div>
    )
}

export default Form
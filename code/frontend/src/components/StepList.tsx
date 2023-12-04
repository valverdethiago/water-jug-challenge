import { SolutionStep } from "../types/solutionStep"

type StepListProps = {
    steps : SolutionStep[];
}

const StepList = ( {steps}: StepListProps) => {
    return(
        <div>
            <ul>
                {steps.map( (step, index)  => (
                    <li key={index}>
                        {step.BucketX} {step.BucketY} {step.Explanation}
                    </li>
                ))} 
            </ul>
        </div>
    )
}

export default StepList;
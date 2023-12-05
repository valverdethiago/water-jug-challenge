import { SolutionStep } from "../types/solutionStep"
import classes from './StepList.module.css'

type StepListProps = {
    steps : SolutionStep[];
}

const StepList = ( {steps}: StepListProps) => {
    return(
        steps.length > 0 ?
        <div className={classes.steplist_container}>
            <table>
                <thead>
                    <tr className={classes.steplist_header_row}>
                        <td>Bucket X</td>
                        <td>Bucket Y</td>
                        <td>Explanation</td>
                    </tr>
                </thead>
                <tbody>
                    {steps.map( (step, key)  => (
                        <tr  role="row" key={key}>
                            <td>{step.BucketX}</td>
                            <td>{step.BucketY}</td>
                            <td>{step.Explanation}</td>
                        </tr>
                    ))} 
                </tbody>
            </table>
        </div>
        : <p></p>
    )
}

export default StepList;
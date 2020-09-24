import React from "react"

class VariableTable extends React.Component {
    
    render() {
        const { actions, configuration, step, routines } = this.props
        const steps = this.props.steps.slice(0, step)
        const keys = configuration.observedVariables

        const newRoutineIndices = routines.reduce((result, routine, i) =>
            result.concat((routines[i-1] ? routines[i-1].steps.length : 0) + (result[i-1] ? result[i-1] : 0)), 
            []
        )

        return (
            <table>
                <thead>
                    DPCCP
                </thead>
                <tbody>
                    {
                    steps.map((step, i) => 
                    <>
                        {
                            newRoutineIndices.includes(i) &&
                            <>
                                <tr style={{background: "gray", width: "100%"}}>
                                    <td colSpan={keys.length + 1}>
                                        {step.subroutineStack.join(" ") + " " + routines[newRoutineIndices.indexOf(i)].name}
                                    </td>
                                </tr>
                                <tr>
                                    <th>Substep</th>
                                    {routines[newRoutineIndices.indexOf(i)].observedVariables.map(key => <th key={key}>{key}</th>)}
                                </tr>
                            </>
                        }
                        <tr>
                            <td></td>
                            {
                            keys.map((key, j) => 
                                <td key={key + i + j}>
                                    {'{'}{step.variables[key]?.join(", ")}{'}'}
                                </td>)
                            }
                        </tr>
                    </>
                )}
                </tbody>
            </table>
        )
    }
}

export default VariableTable
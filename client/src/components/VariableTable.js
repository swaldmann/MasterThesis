import React from "react"

class VariableTable extends React.Component {
    
    render() {
        const { configuration, step, routines } = this.props
        const steps = this.props.steps.slice(0, step)

        const newRoutineIndices = routines.reduce((result, routine, i) =>
            result.concat((routines[i-1] ? routines[i-1].steps.length : 0) + (result[i-1] ? result[i-1] : 0)), 
            []
        )

        const keys = configuration.observedVariables
        const keyLengths = routines.map(r => r.observedVariables.length)//configuration.observedVariables
        const maxKeyLength = Math.max(...keyLengths)
        let routineIndex = -1

        return (
            <table>
                <thead>
                    <tr>
                        <td colSpan={maxKeyLength}>{this.props.algorithm.label}</td>
                    </tr>
                </thead>
                <tbody>
                    {
                    steps.map((step, i) => {
                        const includesI = newRoutineIndices.includes(i)
                        if (includesI) routineIndex = routineIndex + 1;
                        return (
                            <>
                                {
                                    includesI &&
                                    <>
                                        <tr style={{background: "gray", width: "100%"}}>
                                            <td colSpan={keyLengths.length + 1}>
                                                {step.subroutineStack.join(" ") + " " + routines[newRoutineIndices.indexOf(i)].name}
                                            </td>
                                        </tr>
                                        <tr>
                                            {routines[newRoutineIndices.indexOf(i)].observedVariables.map(key => <th key={key}>{key}</th>)}
                                        </tr>
                                    </>
                                }
                                <tr>
                                    {
                                    routines[routineIndex].observedVariables.map((key, j) => 
                                        <td key={key + i + j}>
                                            {'{'}{step.variables[key]?.join(", ")}{'}'}
                                        </td>)
                                    }
                                </tr>
                            </>
                        )})}
                </tbody>
            </table>
        )
    }
}

export default VariableTable
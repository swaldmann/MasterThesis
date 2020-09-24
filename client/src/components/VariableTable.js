import React from "react"

class VariableTable extends React.Component {
    
    render() {
        const { configuration, step, routines } = this.props
        const steps = this.props.steps.slice(0, step)

        const newRoutineIndices = routines.reduce((result, routine, i) =>
            result.concat((routines[i-1] ? routines[i-1].steps.length : 0) + (result[i-1] ? result[i-1] : 0)), 
            []
        )

        const keyLengths = routines.map(r => r.obeservedRelations.length)//configuration.obeservedRelations
        const maxKeyLength = Math.max(...keyLengths)
        let routineIndex = -1
        console.log(routines)

        return (
            <div>
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
                                                {routines[newRoutineIndices.indexOf(i)].obeservedRelations.map(variable => variable.identifier).map(key => <th key={key}>{key}</th>)}
                                            </tr>
                                        </>
                                    }
                                    <tr>
                                        {
                                        routines[routineIndex].obeservedRelations.map((variable, j) => 
                                            <td key={variable.identifier + i + j}>
                                                {'{'}{step.variables[variable.identifier]?.join(", ")}{'}'}
                                            </td>)
                                        }
                                    </tr>
                                </>
                            )})}
                    </tbody>
                </table>
            </div>
        )
    }
}

export default VariableTable
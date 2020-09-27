import React from "react"

class VariableTable extends React.Component {

    render() {
        const { actions, configuration, step, routines } = this.props
        const steps = this.props.steps.slice(0, step + 1)

        const newRoutineIndices = routines.reduce((result, routine, i) =>
            result.concat((routines[i - 1] ? routines[i - 1].steps.length : 0) + (result[i - 1] ? result[i - 1] : 0)),
            []
        )
        const keyLengths = routines.map(r => r.observedRelations.length)//configuration.observedRelations
        const maxKeyLength = Math.max(...keyLengths)
        let routineIndex = -1

        if (routines.length === 0) return (<></>)

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
                                const newRoutineIndex = newRoutineIndices.indexOf(i)
                                const includesI = newRoutineIndex != -1
                                if (includesI) {
                                    routineIndex = routineIndex + 1
                                    console.log(routines[routineIndex].observedRelations)
                                    actions.updateCurrentRoutine(routines[routineIndex].observedRelations)
                                }
                                return (
                                    <>
                                        {
                                            includesI &&
                                            <>
                                                <tr style={{ background: "gray", width: "100%" }}>
                                                    <td colSpan={keyLengths.length + 1}>
                                                        {step.subroutineStack.join(" ") + " " + routines[newRoutineIndex].name}
                                                    </td>
                                                </tr>
                                                <tr>
                                                    {
                                                        routines[newRoutineIndex].observedRelations.map(r =>
                                                            <th key={r.identifier} style={{color: "rgba(" + r.color.R + "," + r.color.G + "," + r.color.B + "," + r.color.A + ")"}}>{r.identifier}</th>
                                                        )
                                                    }
                                                </tr>
                                            </>
                                        }
                                        <tr>
                                            {
                                                routines[routineIndex].observedRelations.map((variable, j) =>
                                                    <td key={variable.identifier + i + j}>
                                                        {'{'}{step.variables[variable.identifier]?.join(", ")}{'}'}
                                                    </td>)
                                            }
                                        </tr>
                                    </>
                                )
                            })}
                    </tbody>
                </table>
            </div>
        )
    }
}

export default VariableTable
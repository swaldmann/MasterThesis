import React from "react"

class VariableTable extends React.Component {
    
    render() {
        const { actions, configuration, step } = this.props
        const steps = this.props.steps.slice(0, step)
        console.log("OOYOF")
        console.log(actions);
        console.log(steps);
        console.log(steps[0])
        console.log(configuration)
        const keys = configuration.observedVariables

        return (
            <table>
                <thead>
                    <tr>
                        {keys.map(key => <th key={key}>{key}</th>)}
                    </tr>
                </thead>
                <tbody>
                    {
                    steps.map((step, i) => 
                        <tr>
                            {
                            keys.map((key, j) => 
                            <td key={key + i + j}>
                                {'{'}{step.variables[key]?.join(", ")}{'}'}
                            </td>)
                            }
                        </tr>)
                    }
                </tbody>
            </table>
        )
    }
}

export default VariableTable
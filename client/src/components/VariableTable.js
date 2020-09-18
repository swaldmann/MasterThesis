import React from "react"

class VariableTable extends React.Component {
    
    render() {
        const { actions, configuration, step } = this.props
        const steps = this.props.steps.slice(0, step)
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
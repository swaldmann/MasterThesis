import React from "react"
import VariableTableEntry from './VariableTableEntry'
class VariableTable extends React.Component {

    render() {
        const { routines } = this.props

        return (
            <div>
                <table>
                    <thead>
                        <tr>
                            <td></td>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            routines.map(routine => 
                                <VariableTableEntry step={routine} level={0} />
                            )
                        }
                    </tbody>
                </table>
            </div>
        )
    }
}

export default VariableTable
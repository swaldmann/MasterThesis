import React from "react"
import VisibleVariableTableEntry from '../containers/VisibleVariableTableEntry'
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
                            routines.map((routine, i) => 
                                <VisibleVariableTableEntry key={i} step={routine} level={0} />
                            )
                        }
                    </tbody>
                </table>
            </div>
        )
    }
}

export default VariableTable
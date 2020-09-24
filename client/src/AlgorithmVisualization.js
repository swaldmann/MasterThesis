import React from "react"
import VisibleJoinProblemSettings from "./containers/VisibleJoinProblemSettings"
import VisibleAlgorithmCanvas from "./containers/VisibleAlgorithmCanvas"
import VisibleVariableTable from "./containers/VisibleVariableTable"

class AlgorithmVisualization extends React.Component {
    render() {
        return (
            <div>
                <VisibleJoinProblemSettings />
                <div style={{display: "flex"}}>
                    <VisibleAlgorithmCanvas />
                    <VisibleVariableTable />
                </div>
            </div>
        )
    }
}

export default AlgorithmVisualization

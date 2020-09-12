import React from "react"
import VisibleJoinProblemSettings from "./containers/VisibleJoinProblemSettings"
import VisibleAlgorithmCanvas from "./containers/VisibleAlgorithmCanvas"
import VisibleVariableTable from "./containers/VisibleVariableTable"

class AlgorithmVisualization extends React.Component {
    render() {
        return (
            <div style={{display:"flex"}}>
                <div style={{width: "400px", flex:"0 0 400px", justifyContent:"space-between", marginRight:"80px"}}>
                    <VisibleJoinProblemSettings />
                    <VisibleAlgorithmCanvas />
                </div>
                <div style={{width: "1 1 100%"}}>
                    <VisibleVariableTable />
                </div>
            </div>
        )
    }
}

export default AlgorithmVisualization

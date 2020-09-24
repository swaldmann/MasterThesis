import React from "react"
import VisibleJoinProblemSettings from "./containers/VisibleJoinProblemSettings"
import VisibleAlgorithmCanvas from "./containers/VisibleAlgorithmCanvas"
import VisibleVariableTable from "./containers/VisibleVariableTable"

class AlgorithmVisualization extends React.Component {
    render() {
        return (
            <>
                <VisibleJoinProblemSettings />
                <main className="flexibleColumn half">
                    <VisibleAlgorithmCanvas />
                    <VisibleVariableTable />
                </main>
            </>
        )
    }
}

export default AlgorithmVisualization

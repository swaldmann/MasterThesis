import React from "react"
import VisibleQueryGraphCanvas from "./containers/VisibleQueryGraphCanvas"
import VisibleAlgorithmCanvas from "./containers/VisibleAlgorithmCanvas"

class AlgorithmVisualization extends React.Component {
    render() {
        return (
            <div style={{display:"flex"}}>
                <div style={{width: "400px", flex:"0 0 400px", justifyContent:"space-between", marginRight:"80px"}}>
                <VisibleQueryGraphCanvas />
                </div>
                <div style={{width: "1 1 100%"}}>
                    <VisibleAlgorithmCanvas />
                </div>
            </div>
        )
    }
}

export default AlgorithmVisualization

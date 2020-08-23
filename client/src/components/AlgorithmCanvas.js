import React from "react"
import Select from "react-select"
import AlgorithmGraph from "../math/AlgorithmGraph"
import QueryGraph from "../math/QueryGraph"
import { ALGORITHMS, QUERY_GRAPH_OPTIONS } from "../constants/AlgorithmConstants"

const ENDPOINT = "http://localhost:8080/api"

class AlgorithmCanvas extends React.Component {
    state = {
        nodeColors: [],
        step: 0
    }

    constructor(props) {
        super(props)
        this.algorithmCanvasRef = React.createRef()
    }

    componentDidUpdate() {
        this.redrawGraph()
    }

    async componentDidMount() {
        this.updateAlgorithm()
    }

    async updateAlgorithm() {
        const response = await fetch(ENDPOINT + "/algorithm/" + this.props.algorithm.value)
        const json = await response.json()
        console.log("JSON");
        console.log(json);
        //const object = JSON.parse(json)
        console.log(json.graphStates)
        console.log(json.graphStates)

        /*this.setState({
            nodeColors: json.graphStates[0].nodeColors
        })*/
    }

    handleAlgorithmChange = algorithm => {
        this.setState({ 
            algorithm: algorithm
        }, () => {
            this.updateAlgorithm()
        })     
    }

    handleNextStep = step => {
        this.setState({
            step: step + 1
        })
    } 

    handlePreviousStep = step => {
        this.setState({
            step: step === 0 ? step : step - 1
        })
    }

    redrawGraph() {        
        const canvas = this.algorithmCanvasRef.current
        const ctx = canvas.getContext('2d')
        const { numberOfRelations, graphTypeIndex } = this.props
        const queryGraph = new QueryGraph(numberOfRelations)
        const graphType = QUERY_GRAPH_OPTIONS[graphTypeIndex]
        const { nodeColors } = this.state

        queryGraph.draw(graphType.value, nodeColors, ctx)
    }

    render() {
        const { algorithm, step } = this.state
        return (
            <div>
                <h3>Algorithm</h3>
                <Select name="color" 
                   className="select" 
                defaultValue={ALGORITHMS[0]} 
                 placeholder="Query Graph" 
                       value={algorithm}
                    onChange={this.handleAlgorithmChange} 
                     options={ALGORITHMS} />
                <button onClick={() => this.handlePreviousStep(step)}>Previous Step</button>
                <button onClick={() => this.handleNextStep(step)}>Next Step</button>
                <p>Step {step}</p>
                <canvas ref={this.algorithmCanvasRef} width="800px" height="800px" style={{width:"800px", height:"800px"}}></canvas>
            </div>
        )
    }
}

export default AlgorithmCanvas

import React from "react"
import Select from "react-select"
import QueryGraph from "../math/QueryGraph"
import { ALGORITHMS, QUERY_GRAPH_OPTIONS } from "../constants/AlgorithmConstants"
import Hotkeys from 'react-hot-keys'

const ENDPOINT = "https://dbs-visualization.ew.r.appspot.com/api"

class AlgorithmCanvas extends React.Component {
    state = {
        initialGraphState: [],
        graphState: [],
        counters: [],
        diffs: []
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
        const { actions, numberOfRelations, graphTypeIndex } = this.props
        const graphType = QUERY_GRAPH_OPTIONS[graphTypeIndex]
        const response = await fetch(ENDPOINT + "/algorithm/" + this.props.algorithm.value
                                              + "/relations/" + numberOfRelations
                                              + "/graphType/" + graphType.value)
        const json = await response.json()
        actions.updateConfiguration(json.configuration)
        actions.updateSteps(json.diffs)

        this.setState({
            initialGraphState: JSON.parse(JSON.stringify(json.begin.graphState)),
            graphState: json.diffs[0].graphState,
            counters: json.begin.counters,
            diffs: json.diffs,
            step: 0
        })
    }

    handleAlgorithmChange = algorithm => {
        this.setState({ 
            algorithm: algorithm
        }, () => {
            this.updateAlgorithm()
        })     
    }

    handleNextStep = step => {
        const { graphState, diffs } = this.state
        const diff = diffs[step + 1]

        const { actions } = this.props
        actions.increaseStep(1)

        this.setState({
            graphState: Object.assign(graphState, diff.graphState)
        })
    } 

    handlePreviousStep = step => {
        const { graphState, diffs } = this.state
        const diff = diffs[step - 1]
        const newGraphState = Object.assign(graphState, diff.graphState)
        const { actions } = this.props
        actions.decreaseStep(1)

        this.setState({
            graphState: newGraphState
        })
    }

    redrawGraph() {        
        const canvas = this.algorithmCanvasRef.current
        const ctx = canvas.getContext('2d')
        const { numberOfRelations, graphTypeIndex } = this.props
        const queryGraph = new QueryGraph(numberOfRelations)
        const graphType = QUERY_GRAPH_OPTIONS[graphTypeIndex]
        const { graphState } = this.state

        queryGraph.draw(graphType.value, graphState.nodeColors, ctx)
    }

    onKeyDown(keyName, e, handle) {
        const { step } = this.props
        const { diffs } = this.state
        if (keyName === "d" && step < diffs.length - 1) {
            this.handleNextStep(step)
        } else if (keyName === "a" && step > 0) {
            this.handlePreviousStep(step)
        } else if (keyName === "r") {
            this.updateAlgorithm()
        }
    }

    render() {
        const { algorithm, diffs } = this.state
        const { step } = this.props
        const isFirstStep = step === 0
        const isLastStep = step === diffs.length - 1

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
                <Hotkeys keyName="a,d" onKeyDown={this.onKeyDown.bind(this)} allowRepeat={true} />
                <button onClick={() => this.handleAlgorithmChange(algorithm)}>Recalculate Algorithm (r)</button>
                <div style={{marginTop: "50px"}}>
                    <canvas ref={this.algorithmCanvasRef} width="400px" height="400px" style={{width:"400px", height:"400px"}}></canvas>
                    <button onClick={() => this.handlePreviousStep(step)} disabled={isFirstStep}>Previous Step (a)</button>
                    <button onClick={() => this.handleNextStep(step)} disabled={isLastStep}>Next Step (d)</button>
                    <p>Step {step + 1} of {diffs.length}</p>
                </div>
            </div>
        )
    }
}

export default AlgorithmCanvas

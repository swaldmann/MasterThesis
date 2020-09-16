import React from "react"
import Select from "react-select"
import QueryGraph from "../math/QueryGraph"
import { ALGORITHMS, QUERY_GRAPH_OPTIONS } from "../constants/AlgorithmConstants"

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
        console.log("Res")
        console.log(response)
        console.log(json)

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

        console.log("Next step");
        console.log(diff);
        console.log(this.state)

        const { actions } = this.props
        actions.increaseStep(1)

        this.setState({
            graphState: Object.assign(graphState, diff.graphState)
        })
    } 

    handlePreviousStep = step => {
        const { graphState, diffs } = this.state
        const diff = diffs[step - 1]

        //let newGraphState
        /*if (step === 1) {
            newGraphState = JSON.parse(JSON.stringify(this.state.initialGraphState))
        } else {
            
        }*/
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

    render() {
        const { algorithm, diffs } = this.state
        const { step } = this.props

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
                <button onClick={() => this.handlePreviousStep(step)} disabled={step === 0}>Previous Step</button>
                <button onClick={() => this.handleNextStep(step)} disabled={step === diffs.length - 1}>Next Step</button>
                <p>Step {step + 1} of {diffs.length}</p>
                <canvas ref={this.algorithmCanvasRef} width="400px" height="400px" style={{width:"400px", height:"400px"}}></canvas>
            </div>
        )
    }
}

export default AlgorithmCanvas

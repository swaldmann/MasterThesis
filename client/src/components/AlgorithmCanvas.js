import React from "react"
import Select from "react-select"
import QueryGraph from "../math/QueryGraph"
import { ALGORITHMS, QUERY_GRAPH_OPTIONS } from "../constants/AlgorithmConstants"
import Hotkeys from 'react-hot-keys'

//const ENDPOINT = "https://dbs-visualization.ew.r.appspot.com/api"
const ENDPOINT = "http://localhost:8080/api"
class AlgorithmCanvas extends React.Component {
    state = {
        graphState: [],
        counters: [],
        steps: []
    }

    constructor(props) {
        super(props)
        this.algorithmCanvasRef = React.createRef()
    }

    async componentDidMount() {
        this.updateAlgorithm()
    }

    async updateAlgorithm() {
        const { actions, settingNumberOfRelations, settingGraphTypeValue } = this.props
        const response = await fetch(ENDPOINT + "/algorithm/" + this.props.algorithm.value
                                              + "/relations/" + settingNumberOfRelations
                                              + "/graphType/" + settingGraphTypeValue)
        const json = await response.json()
        actions.changeQueryGraph(json.queryGraph)
        actions.updateConfiguration(json.configuration)
        actions.resetSteps()
        const concatSteps = (result, routine) => result.concat(routine.steps)
        const steps = json.routines.reduce(concatSteps, [])
        actions.updateSteps(steps)

        console.log("YO")
        console.log(json.routines);
        actions.updateRoutines(json.routines)

        this.setState({
            graphState: steps[0].graphState,
            steps: steps,
            routineStepLengths: json.routines.map(r => r.steps.length),
            step: 0
        })

        this.redrawGraph()
    }

    handleAlgorithmChange = algorithm => {
        this.setState({ 
            algorithm: algorithm
        }, () => {
            this.updateAlgorithm()
        })     
    }

    handleNextStep = step => {
        const { steps } = this.state
        const nextStep = steps[step + 1]

        const { actions } = this.props
        actions.increaseStep(1)

        this.setState({
            graphState: nextStep.graphState
        })
        this.redrawGraph()
    } 

    handlePreviousStep = step => {
        const { steps } = this.state
        const previousStep = steps[step - 1]
        const { actions } = this.props
        actions.decreaseStep(1)

        this.setState({
            graphState: previousStep.graphState
        })
        this.redrawGraph()
    }

    redrawGraph() {
        const { graphState } = this.state
        if (!graphState.nodeColors) return
        const canvas = this.algorithmCanvasRef.current
        const ctx = canvas.getContext('2d')
        const { queryGraph, settingGraphTypeValue } = this.props
        console.log(queryGraph);
        const drawnQueryGraph = new QueryGraph(queryGraph)
        drawnQueryGraph.draw(settingGraphTypeValue, graphState.nodeColors, ctx)
    }

    onKeyDown(keyName, e, handle) {
        const { step } = this.props
        const { steps } = this.state
        if (keyName === "d" && step < steps.length - 1) {
            this.handleNextStep(step)
        } else if (keyName === "a" && step > 0) {
            this.handlePreviousStep(step)
        } else if (keyName === "r") {
            this.updateAlgorithm()
        }
    }

    render() {
        const { algorithm, steps } = this.state
        const { step } = this.props
        const isFirstStep = step === 0
        const isLastStep = step === steps.length - 1

        return (
            <div>
                <h3>Algorithm</h3>
                <Select name="color" 
                   className="select" 
                defaultValue={ALGORITHMS[0]} 
                 placeholder="Algorithm" 
                       value={algorithm}
                    onChange={this.handleAlgorithmChange} 
                     options={ALGORITHMS} />
                <Hotkeys keyName="a,d,r" onKeyDown={this.onKeyDown.bind(this)} allowRepeat={true} />
                <button className="emphasized" onClick={() => this.handleAlgorithmChange(algorithm)}>Recalculate Algorithm (r)</button>
                <div style={{marginTop: "50px"}}>
                    <canvas ref={this.algorithmCanvasRef} width="500px" height="500px" style={{width:"500px", height:"500px"}}></canvas>
                    <button onClick={() => this.handlePreviousStep(step)} disabled={isFirstStep}>Previous Step (a)</button>
                    <button onClick={() => this.handleNextStep(step)} disabled={isLastStep}>Next Step (d)</button>
                    <p>Step {step + 1} of {steps.length}</p>
                </div>
            </div>
        )
    }
}

export default AlgorithmCanvas

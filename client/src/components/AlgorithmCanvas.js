import React from "react"
import QueryGraph from "../math/QueryGraph"
import Hotkeys from 'react-hot-keys'

class AlgorithmCanvas extends React.Component {
    state = {
        graphState: [],
        counters: []
    }

    constructor(props) {
        super(props)
        this.algorithmCanvasRef = React.createRef()
    }

    resize = () => {
        const canvas = this.algorithmCanvasRef.current
        const rect = canvas.parentNode.getBoundingClientRect()
        canvas.width = rect.width
        canvas.height = rect.width
        this.redrawGraph()
    }

    componentDidMount() {
        window.addEventListener('resize', this.resize)
    }

    componentWillUnmount() {
        window.removeEventListener('resize', this.resize)
    }

    componentDidUpdate(prevProps) {
        if (prevProps.graphState !== this.props.graphState) {
            this.redrawGraph()
        }
    }

    handleNextStep = step => {
        const { steps, actions } = this.props
        const nextStep = steps[step + 1]
        actions.updateCurrentStepUUID(nextStep.uuid)
        actions.increaseStep(1)
        actions.updateGraphState(nextStep.graphState)
        //window.scrollTo(0, document.body.scrollHeight);
    }

    handlePreviousStep = step => {
        const { actions, steps } = this.props
        const previousStep = steps[step - 1]
        actions.updateCurrentStepUUID(previousStep.uuid)
        actions.decreaseStep(1)
        actions.updateGraphState(previousStep.graphState)
        //window.scrollTo(0, document.body.scrollHeight);
    }

    redrawGraph() {
        const graphState = this.props.graphState
        const { queryGraph, settingGraphTypeValue } = this.props
        if (!graphState.nodeColors) return
        const canvas = this.algorithmCanvasRef.current
        const ctx = canvas.getContext('2d')
        const drawnQueryGraph = new QueryGraph(queryGraph)
        drawnQueryGraph.draw(settingGraphTypeValue, graphState.nodeColors, ctx)
    }

    onKeyDown(keyName, e, handle) {
        const { step, steps } = this.props
        if (keyName === "d" && step < steps.length - 1) {
            this.handleNextStep(step)
        } else if (keyName === "a" && step > 0) {
            this.handlePreviousStep(step)
        }
    }

    render() {
        const { step, steps } = this.props
        const isFirstStep = step === 0
        const isLastStep = step === steps.length - 1

        return (
            <>
                <Hotkeys keyName="a,d" onKeyDown={this.onKeyDown.bind(this)} allowRepeat={true} />
                <div className="fixed">
                    <canvas ref={this.algorithmCanvasRef} width="500px" height="500px" style={{ width: "100%", height: "50vw" }}></canvas>
                    <button onClick={() => this.handlePreviousStep(step)} disabled={isFirstStep}>Previous Step<span className="shortcut">A</span></button>
                    <button onClick={() => this.handleNextStep(step)} disabled={isLastStep}>Next Step<span className="shortcut">D</span></button>
                    <p>Step {step + 1} of {steps.length}</p>
                </div>

            </>
        )
    }
}

export default AlgorithmCanvas

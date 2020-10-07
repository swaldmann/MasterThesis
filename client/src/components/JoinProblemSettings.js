import React from "react"
import Select from "react-select"
import Slider from 'rc-slider'
import Hotkeys from 'react-hot-keys'
import '../styles/App.css'
import "rc-slider/assets/index.css";
import { QUERY_GRAPH_OPTIONS } from "../constants/AlgorithmConstants"

//const ENDPOINT = "https://dbs-visualization.ew.r.appspot.com/api" // Production
const ENDPOINT = "http://localhost:8080/api" // Local development

// Create an array with values from 3 to 8 and store them in an object
// For some reason this doesn't work, despite giving the same result as ``marksStatic``
/*const marks = [...Array(6).keys()]
              .map(x => x + 3)
              .reduce((object, number) => Object.defineProperty(object, number, { value: number.toString() }), {})*/

// Fallback for the above code
const marksStatic = {
    3: "3",
    4: "4",
    5: "5",
    6: "6",
    7: "7",
    8: "8",
    9: "9",
   10: "10",
}

class JoinProblemSettings extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            selectedQueryGraphIndex: 4
        }
    }
    
    async componentWillMount() {
        await this.fetchAvailableAlgorithms()
        await this.updateAlgorithm()
    }
    
    async fetchAvailableAlgorithms() {
        const response = await fetch(ENDPOINT + "/algorithms")
        const json = await response.json()
        const { actions } = this.props
        actions.updateAlgorithms(json.algorithms)
        actions.changeAlgorithm(json.algorithms[0])
    }

    async updateAlgorithm() {
        const { 
            actions, 
            settingNumberOfRelations, 
            settingGraphTypeValue,
            algorithms,
            algorithm 
        } = this.props
        const response = await fetch(ENDPOINT + "/algorithm/" + algorithm.value
                                              + "/relations/" + settingNumberOfRelations
                                              + "/graphType/" + settingGraphTypeValue)
        const json = await response.json()
        actions.changeQueryGraph(json.queryGraph)
        actions.updateConfiguration(json.configuration)
        actions.resetSteps()

        const concatSteps = (result, routine) => result.concat(routine.steps)
        const steps = json.routines.reduce(concatSteps, []).flatMap(function loop(step) {
            if (step && step.steps && !step.uuid)
                return step.steps.flatMap(loop)
            else
                return [step]
        }).filter(step => step && step.hasOwnProperty("uuid"))
        console.log(steps);
        actions.updateSteps(steps)
        actions.updateRoutines(json.routines)
        const firstStep = steps[0]
        if (firstStep) {
            actions.updateCurrentStepUUID(steps[0].uuid)
            actions.updateGraphState(steps[0].graphState)
        }
    }
    
    handleNumberOfRelationsChange = numberOfRelations => {
        const actions = this.props.actions
        actions.changeQueryNumberOfRelations(numberOfRelations)
    }

    handleGraphTypeOptionChange = graphTypeOption => {
        const actions = this.props.actions
        actions.changeQueryGraphTypeOption(graphTypeOption.value)
    }

    handleAlgorithmChange = algorithm => {
        const { actions } = this.props
        actions.changeAlgorithm(algorithm)
    }
    
    onKeyDown(keyName, e, handle) {
        if (keyName === "r") {
            this.updateAlgorithm()
        }
    }

    render() {
        const { graphTypeOptionValue, numberOfRelations, algorithms, algorithm } = this.props
        if (!algorithms) { 
            return <div />
        }
        const graphTypeOption = QUERY_GRAPH_OPTIONS.find(o => o.value === graphTypeOptionValue)
        console.log("Render");

        return (
            <header className="grid" style={{ background: "#1f2329" }}>
                <div>
                    <h5>Number of relations</h5>
                    <Slider className="slider"
                                marks={marksStatic}
                          handleStyle={{background:"white", border: 0, height:"26px", width:"26px", marginTop:"-9px"}}
                           trackStyle={{background:"white", height: "4px", borderRadius:"0px"}} 
                            railStyle={{background:"white", height: "4px", borderRadius:"0px"}}
                             dotStyle={{height:"20px", transform:"translate(2px, 6px)", border:"none", borderRadius:"0px", width:"3px"}}
                                style={{width:"100%"}}    
                                 dots={true}
                             onChange={this.handleNumberOfRelationsChange}
                                  min={3}
                         defaultValue={numberOfRelations}
                                  max={10} />
                </div>
                <div>
                    <h5>Graph Type</h5>
                    <Select className="select" 
                                 name="color" 
                                style={{width:"100%"}}
                         defaultValue={QUERY_GRAPH_OPTIONS[4]} 
                          placeholder="Query Graph" 
                                value={graphTypeOption} 
                             onChange={this.handleGraphTypeOptionChange} 
                              options={QUERY_GRAPH_OPTIONS} />
                    {
                        graphTypeOption.value === "tree" && 
                        <div className="info">Only complete binary trees are supported.</div>
                    }
                </div>
                <div>
                    <h5>Algorithm</h5>
                    {algorithm && <Select 
                            name="color" 
                       className="select"
                     placeholder="Algorithm"
                           value={algorithm}
                        onChange={this.handleAlgorithmChange} 
                         options={algorithms} />}
                </div>
                <div>
                    <h5>Calculation</h5>
                    <Hotkeys keyName="r" onKeyDown={this.onKeyDown.bind(this)} allowRepeat={true} />
                    <button className="emphasized" 
                              onClick={() => this.updateAlgorithm()}>
                        Recalculate Algorithm<span className="shortcut">R</span>
                    </button>
                </div>
            </header>
        )
    }
}

export default JoinProblemSettings

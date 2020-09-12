import React from "react"
import Select from "react-select"
import Slider from 'rc-slider';
import { QUERY_GRAPH_OPTIONS } from "../constants/AlgorithmConstants"
import '../styles/App.css'
import "rc-slider/assets/index.css";

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
    
    handleNumberOfRelationsChange = numberOfRelations => {
        const actions = this.props.actions
        actions.changeQueryNumberOfRelations(numberOfRelations)
    }

    handleGraphTypeChange = graphType => {
        const actions = this.props.actions
        const graphTypeIndex = QUERY_GRAPH_OPTIONS.findIndex(type => type.value === graphType.value)
        actions.changeQueryGraphType(graphTypeIndex)
    }

    render() {
        const { graphTypeIndex } = this.props
        const graphType = QUERY_GRAPH_OPTIONS[graphTypeIndex]
        
        return (
            <div>
                <h3>Graph</h3>
                <div>
                    <h5>Number of relations</h5>
                    <Slider className="slider"
                                marks={marksStatic}
                          handleStyle={{background:"white", border: 0, height:"26px",width:"26px",marginTop:"-9px"}}
                           trackStyle={{background:"white", height: "4px", borderRadius:"0px"}} 
                            railStyle={{background:"white", height: "4px", borderRadius:"0px"}}
                             dotStyle={{height:"20px", transform:"translate(2px, 6px)", border:"none", borderRadius:"0px",width:"3px"}}
                                style={{width:"100%"}}    
                                 dots={true}
                             onChange={this.handleNumberOfRelationsChange}
                                  min={3}
                         defaultValue={5}
                                  max={10} />
                    <h5>Graph Type</h5>
                    <Select className="select" 
                                 name="color" 
                                style={{width:"100%"}}
                         defaultValue={QUERY_GRAPH_OPTIONS[4]} 
                          placeholder="Query Graph" 
                                value={graphType} 
                             onChange={this.handleGraphTypeChange} 
                              options={QUERY_GRAPH_OPTIONS} />
                    {graphType.value === "tree" && <div className="info">Only complete binary trees are supported.</div>}
                </div>
            </div>
        )
    }
}

export default JoinProblemSettings

import React from "react"
import Select from "react-select"
import Slider from 'rc-slider';
import QueryGraph from "../math/QueryGraph"
import { QUERY_GRAPH_OPTIONS } from "../constants/AlgorithmConstants"
import '../styles/App.css'
import "rc-slider/assets/index.css";


// Create an array with values from 3 to 8 and store them in an object
// For some reason this doesn't work, despite giving the same result as ``marks2``
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

class QueryGraphCanvas extends React.Component {
    constructor(props) {
        super(props)
        console.log(props);
        this.graphCanvasRef = React.createRef()
    }

    componentDidUpdate() {
        this.redrawGraph()
    }

    handleNumberOfRelationsChange = numberOfRelations => {
        const actions = this.props.actions
        actions.changeQueryNumberOfRelations(numberOfRelations)
        console.log("R");
        /*this.setState({
            numberOfRelations
        }, () => {
            this.redrawGraph()
        })*/
    }

    handleGraphTypeChange = graphType => {
        const actions = this.props.actions
        const graphTypeIndex = QUERY_GRAPH_OPTIONS.findIndex(type => type.value === graphType.value)
        console.log(this);
        actions.changeQueryGraphType(graphTypeIndex)
        this.redrawGraph()
        /*this.setState({ 
            graphType
        }, () => {
            this.redrawGraph()
        })*/
    }

    handleCreateGraph = () => {
        console.log("Button")
    }

    redrawGraph() {        
        const canvas = this.graphCanvasRef.current
        const ctx = canvas.getContext('2d')
        ctx.strokeStyle = 'rgb(250, 250, 250)'
        ctx.fillStyle = 'rgb(250, 250, 250)'
        const { numberOfRelations } = this.props 
        const queryGraph = new QueryGraph(numberOfRelations)
        const graphType = QUERY_GRAPH_OPTIONS[this.props.graphTypeIndex]
        queryGraph.draw(graphType.value, [], ctx)
    }

    componentDidMount() {
        this.redrawGraph()
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
                         defaultValue={QUERY_GRAPH_OPTIONS[1]} 
                          placeholder="Query Graph" 
                                value={graphType} 
                             onChange={this.handleGraphTypeChange} 
                              options={QUERY_GRAPH_OPTIONS} />
                    {graphType.value === "tree" && <div className="info">Only complete binary trees are supported.</div>}
                    {/*<button style={{marginTop:"100px"}} onClick={this.handleCreateGraph}>Create Graph</button>*/}
                    <h5>Graph</h5>
                    <canvas ref={this.graphCanvasRef} 
                         height="400px"
                          width="400px" 
                          style={{width: "400px", height: "400px"}} />
                </div>
            </div>
        )
    }
}

export default QueryGraphCanvas

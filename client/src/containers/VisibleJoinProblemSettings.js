import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import JoinProblemSettings from '../components/JoinProblemSettings'

import * as Actions from '../actions'

import { ALGORITHMS } from "../constants/AlgorithmConstants"

const mapStateToProps = state => {
    return {        
        // Required to keep track of the currently selected 
        // settings without immedeately triggering a server
        // request
        numberOfRelations: state.settingNumberOfRelations,
        graphTypeOptionValue: state.settingGraphTypeValue,
        algorithm: ALGORITHMS[state.algorithmIndex],

        steps: state.steps,
        settingNumberOfRelations: state.settingNumberOfRelations,
        settingGraphTypeValue: state.settingGraphTypeValue,
    }
}

const mapDispatchToProps = dispatch => ({
    actions: bindActionCreators(Actions, dispatch)
})

const VisbleJoinProblemSettings = connect(
    mapStateToProps,
    mapDispatchToProps
)(JoinProblemSettings)

export default VisbleJoinProblemSettings
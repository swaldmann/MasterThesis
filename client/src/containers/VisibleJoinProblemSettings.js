import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import JoinProblemSettings from '../components/JoinProblemSettings'

import * as Actions from '../actions'

const mapStateToProps = state => {
    return {        
        // Required to keep track of the currently selected 
        // settings without immedeately triggering a server
        // request
        numberOfRelations: state.settingNumberOfRelations,
        graphTypeOptionValue: state.settingGraphTypeValue
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
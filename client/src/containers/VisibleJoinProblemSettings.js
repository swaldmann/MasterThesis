import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import JoinProblemSettings from '../components/JoinProblemSettings'

import * as Actions from '../actions'

const mapStateToProps = state => {
    return {
        numberOfRelations: state.queryGraph.numberOfRelations,
        graphTypeIndex: state.queryGraph.graphTypeIndex
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
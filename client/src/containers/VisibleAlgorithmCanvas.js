import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { ALGORITHMS } from '../constants/AlgorithmConstants'
import AlgorithmCanvas from '../components/AlgorithmCanvas'

import * as Actions from '../actions'

const mapStateToProps = state => {
    return {
        algorithm: ALGORITHMS[state.algorithmIndex],
        settingNumberOfRelations: state.settingNumberOfRelations,
        settingGraphTypeValue: state.settingGraphTypeValue,

        routine: state.currentRoutine,
        configuration: state.configuration,
        graphState: state.graphState,
        queryGraph: state.queryGraph,
        steps: state.steps,
        step: state.step
    }
}

const mapDispatchToProps = dispatch => ({
    actions: bindActionCreators(Actions, dispatch)
})

const VisbleAlgorithmCanvas = connect(
    mapStateToProps,
    mapDispatchToProps
)(AlgorithmCanvas)

export default VisbleAlgorithmCanvas
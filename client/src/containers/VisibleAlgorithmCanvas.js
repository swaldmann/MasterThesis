import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import AlgorithmCanvas from '../components/AlgorithmCanvas'

import * as Actions from '../actions'

const mapStateToProps = state => {
    return {
        algorithm: state.algorithm,
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
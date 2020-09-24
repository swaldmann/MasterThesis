import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import VariableTable from '../components/VariableTable'
import { ALGORITHMS } from '../constants/AlgorithmConstants'

import * as Actions from '../actions'

const mapStateToProps = state => {
    return {
        routines: state.routines,
        steps: state.steps,
        step: state.step,
        configuration: state.configuration,
        algorithm: ALGORITHMS[state.algorithmIndex]
    }
}

const mapDispatchToProps = dispatch => ({
    actions: bindActionCreators(Actions, dispatch)
})

const VisbleVariableTable = connect(
    mapStateToProps,
    mapDispatchToProps
)(VariableTable)



export default VisbleVariableTable
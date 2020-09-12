import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import VariableTable from '../components/VariableTable'

import * as Actions from '../actions'

const mapStateToProps = state => {
    return {
        steps: state.steps,
        step: state.step,
        configuration: state.configuration
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
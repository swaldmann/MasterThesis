import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import * as Actions from '../actions'
import VariableTableEntry from '../components/VariableTableEntry'

const mapStateToProps = state => {
    return {
        currentMaxStep: state.step
    }
}

const mapDispatchToProps = dispatch => ({
    actions: bindActionCreators(Actions, dispatch)
})

const VisbleVariableTableEntry = connect(
    mapStateToProps,
    mapDispatchToProps
)(VariableTableEntry)

export default VisbleVariableTableEntry

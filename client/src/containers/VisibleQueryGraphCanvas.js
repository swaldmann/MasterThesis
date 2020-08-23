import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import QueryGraphCanvas from '../components/QueryGraphCanvas'

import * as Actions from '../actions'

const mapStateToProps = state => {
    console.log(state);
    return {
        numberOfRelations: state.queryGraph.numberOfRelations,
        graphTypeIndex: state.queryGraph.graphTypeIndex
    }
}

const mapDispatchToProps = dispatch => ({
    actions: bindActionCreators(Actions, dispatch)
})

const VisbleQueryGraphCanvas = connect(
    mapStateToProps,
    mapDispatchToProps
)(QueryGraphCanvas)

export default VisbleQueryGraphCanvas
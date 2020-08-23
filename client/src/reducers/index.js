import {
    combineReducers
} from 'redux'

import {
    algorithmIndex
} from './algorithm'
import {
    queryGraph
} from './queryGraph'

const appReducer = combineReducers({
    queryGraph,
    algorithmIndex
})

const rootReducer = (state, action) => {
    if (action.type === 'CLEAR_STATE') {
        state = undefined
    }
    return appReducer(state, action)
}

export default rootReducer
import { combineReducers } from 'redux'
import { algorithmIndex } from './algorithm'
import { queryGraph, settingGraphTypeValue, settingNumberOfRelations } from './queryGraph'
import { variables } from './variables'
import { configuration } from './configuration'
import { steps } from './steps'
import { step } from './step'

const appReducer = combineReducers({
    queryGraph,
    settingGraphTypeValue,
    settingNumberOfRelations,
    variables,
    algorithmIndex,
    steps,
    step,
    configuration
})

const rootReducer = (state, action) => {
    if (action.type === 'CLEAR_STATE') {
        state = undefined
    }
    return appReducer(state, action)
}

export default rootReducer
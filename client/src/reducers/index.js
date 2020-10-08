import { combineReducers } from 'redux'
import { algorithms, algorithm } from './algorithm'
import { queryGraph, settingGraphTypeValue, settingNumberOfRelations } from './queryGraph'
import { configuration } from './configuration'
import { routines } from './routines'
import { steps } from './steps'
import { step, stepUUID } from './step'
import { graphState } from './graphState'

const appReducer = combineReducers({
    queryGraph,
    settingGraphTypeValue,
    settingNumberOfRelations,
    graphState,
    algorithms,
    algorithm,
    routines,
    steps,
    step,
    stepUUID,
    configuration
})

const rootReducer = (state, action) => {
    if (action.type === 'CLEAR_STATE') {
        state = undefined
    }
    return appReducer(state, action)
}

export default rootReducer
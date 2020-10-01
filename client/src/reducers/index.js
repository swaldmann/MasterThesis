import { combineReducers } from 'redux'
import { algorithms, algorithm } from './algorithm'
import { queryGraph, settingGraphTypeValue, settingNumberOfRelations } from './queryGraph'
import { variables } from './variables'
import { configuration } from './configuration'
import { routines, currentRoutine } from './routines'
import { steps } from './steps'
import { step, stepUUID, renderedStep } from './step'
import { graphState } from './graphState'

const appReducer = combineReducers({
    queryGraph,
    settingGraphTypeValue,
    settingNumberOfRelations,
    variables,
    graphState,
    algorithms,
    algorithm,
    routines,
    currentRoutine,
    steps,
    step,
    stepUUID,
    renderedStep,
    configuration
})

const rootReducer = (state, action) => {
    if (action.type === 'CLEAR_STATE') {
        state = undefined
    }
    return appReducer(state, action)
}

export default rootReducer
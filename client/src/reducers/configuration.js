import {
    UPDATE_CONFIGURATION
} from '../constants/ActionTypes'

const DEFAULT_CONFIGURATION = {
    observedVariables: []
}

export function configuration(state = DEFAULT_CONFIGURATION, action) {
    switch (action.type) {
        case UPDATE_CONFIGURATION:
            return action.configuration
        default:
            return state
    }
}

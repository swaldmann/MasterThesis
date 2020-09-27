import {
    UPDATE_CURRENT_ROUTINE,
    UPDATE_ROUTINES
} from '../constants/ActionTypes'

export function routines(state = [], action) {
    switch (action.type) {
        case UPDATE_ROUTINES:
            return action.routines
        default:
            return state
    }
}

export function currentRoutine(state = {}, action) {
    switch (action.type) {
        case UPDATE_CURRENT_ROUTINE:
            return action.routine
        default:
            return state
    }
}
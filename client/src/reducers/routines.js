import {
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

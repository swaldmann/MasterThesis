import {
    CHANGE_ALGORITHM
} from '../constants/ActionTypes'

export function algorithmIndex(state = 0, action) {
    switch (action.type) {
        case CHANGE_ALGORITHM:
            return action.algorithmIndex
        default:
            return state
    }
}

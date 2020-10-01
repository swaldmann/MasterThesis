import {
    CHANGE_ALGORITHM,
    UPDATE_ALGORITHMS
} from '../constants/ActionTypes'

export function algorithms(state = [], action) {
    switch (action.type) {
        case UPDATE_ALGORITHMS:
            return action.algorithms
        default:
            return state
    }
}

export function algorithm(state = {}, action) {
    switch (action.type) {
        case CHANGE_ALGORITHM:
            return action.algorithm
        default:
            return state
    }
}

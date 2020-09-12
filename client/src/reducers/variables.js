import {
    ADD_VARIABLE_ROW,
    POP_VARIABLE_ROW
} from '../constants/ActionTypes'

export function variables(state = [], action) {
    switch (action.type) {
        case ADD_VARIABLE_ROW:
            return [...state, action.variableRow]
        case POP_VARIABLE_ROW:
            return state.slice(0, -1)
        default:
            return state
    }
}

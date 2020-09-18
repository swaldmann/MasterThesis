import {
    INCREASE_STEP,
    DECREASE_STEP,
    RESET_STEPS
} from '../constants/ActionTypes'

export function step(state = 0, action) {
    switch (action.type) {
        case INCREASE_STEP:
            return state + action.increase
        case DECREASE_STEP:
            return state - action.decrease
        case RESET_STEPS:
            return 0
        default:
            return state
    }
}

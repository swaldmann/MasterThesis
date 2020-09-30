import {
    INCREASE_STEP,
    DECREASE_STEP,
    RESET_STEPS,
    INCREASE_RENDERED_STEP,
    DECREASE_RENDERED_STEP,
    RESET_RENDERED_STEP
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

export function renderedStep(state = 0, action) {
    switch(action.type) {
        case INCREASE_RENDERED_STEP:
            return state + action.increase
        case DECREASE_RENDERED_STEP:
            return state - action.decrease
        case RESET_RENDERED_STEP:
            return 0
        default:
            return state
    }
}
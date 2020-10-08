import {
    INCREASE_STEP,
    DECREASE_STEP,
    RESET_STEPS,
    UPDATE_STEP_UUID
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

export function stepUUID(state = "", action) {
    switch(action.type) {
        case UPDATE_STEP_UUID:
            return action.uuid
        default:
            return state
    }
}
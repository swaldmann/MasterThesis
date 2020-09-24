import {
    UPDATE_GRAPH_STATE
} from '../constants/ActionTypes'

export function graphState(state = {}, action) {
    switch (action.type) {
        case UPDATE_GRAPH_STATE:
            return action.graphState
        default:
            return state
    }
}

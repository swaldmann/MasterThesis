import {
    CHANGE_QUERY_GRAPH_TYPE,
    CHANGE_QUERY_NUMBER_RELATIONS
    
} from '../constants/ActionTypes'

const QUERY_GRAPH_DEFAULT = {
    numberOfRelations: 5,
    graphTypeIndex: 4
}

export function queryGraph(state = QUERY_GRAPH_DEFAULT, action) {
    console.log(state);
    switch (action.type) {
        case CHANGE_QUERY_GRAPH_TYPE:
            return {...state, graphTypeIndex: action.graphTypeIndex}
        case CHANGE_QUERY_NUMBER_RELATIONS:
            return {...state, numberOfRelations: action.numberOfRelations}
        default:
            return state
    }
}

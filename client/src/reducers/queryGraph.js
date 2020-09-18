import {
    CHANGE_QUERY_GRAPH_TYPE,
    CHANGE_QUERY_NUMBER_RELATIONS,
    CHANGE_QUERY_GRAPH    
} from '../constants/ActionTypes'

const QUERY_GRAPH_DEFAULT = {
    problemID: 0,
    problemType: "chain",
    problemTypeFormatted: "Chain",
    problemNeighbors: {
        0: 1,
        1: 0
    },
    problemNumberOfRelations: 2,
    problemRelations: [
        {
            relationCardinality: 135.1106833796348,
            relationName: "<unknown>",
            relationPID: 0,
            relationRID: 0
        },
        {
            relationCardinality: 7.70954008593561,
            relationName: "<unknown>",
            relationPID: 0,
            relationRID: 1
        }
    ],
    problemSelectivities: {
        "0,1": 0.12570116771021658
    }
}

export function queryGraph(state = QUERY_GRAPH_DEFAULT, action) {
    switch(action.type) {
        case CHANGE_QUERY_GRAPH:
            return action.queryGraph
        default:
            return state
    }
}

export function settingGraphTypeValue(state = "moerkotte", action) {
    switch (action.type) {
        case CHANGE_QUERY_GRAPH_TYPE:
            return action.graphTypeOptionValue
        default:
            return state
    }
}

export function settingNumberOfRelations(state = 5, action) {
    switch (action.type) {
        case CHANGE_QUERY_NUMBER_RELATIONS:
            return action.numberOfRelations
        default:
            return state
    }
}
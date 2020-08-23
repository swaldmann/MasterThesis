import * as types from '../constants/ActionTypes'

// Query Graph
export const changeQueryNumberOfRelations = numberOfRelations => ({ type: types.CHANGE_QUERY_NUMBER_RELATIONS, numberOfRelations })
export const changeQueryGraphType = graphTypeIndex => ({ type: types.CHANGE_QUERY_GRAPH_TYPE, graphTypeIndex })

// Algorithm
export const changeAlgorithm = algorithmIndex => ({ type: types.CHANGE_ALGORITHM, algorithmIndex })

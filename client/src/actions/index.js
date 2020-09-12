import * as types from '../constants/ActionTypes'

// Query Graph
export const changeQueryNumberOfRelations = numberOfRelations => ({ type: types.CHANGE_QUERY_NUMBER_RELATIONS, numberOfRelations })
export const changeQueryGraphType = graphTypeIndex => ({ type: types.CHANGE_QUERY_GRAPH_TYPE, graphTypeIndex })

// Algorithm
export const changeAlgorithm = algorithmIndex => ({ type: types.CHANGE_ALGORITHM, algorithmIndex })

// Variables
export const addVariableRow = variableRow => ({ type: types.ADD_VARIABLE_ROW, variableRow })
export const popVariableRow = () => ({ type: types.POP_VARIABLE_ROW })

// Steps
export const updateSteps = steps => ({ type: types.UPDATE_STEPS, steps })

// Update Configuration
export const updateConfiguration = configuration => ({ type: types.UPDATE_CONFIGURATION, configuration })

// Step
export const increaseStep = increase => ({ type: types.INCREASE_STEP, increase })
export const decreaseStep = decrease => ({ type: types.DECREASE_STEP, decrease })
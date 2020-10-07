import * as types from '../constants/ActionTypes'

// Algorithm
export const updateAlgorithms = algorithms => ({ type: types.UPDATE_ALGORITHMS, algorithms })

// Options
export const changeOptionNumberOfRelations = numberOfRelations => ({ type: types.CHANGE_QUERY_NUMBER_RELATIONS, numberOfRelations })
export const changeOptionQueryGraphType = graphType => ({ type: types.CHANGE_QUERY_GRAPH_TYPE, graphType })
export const changeOptionAlgorithm = algorithm => ({ type: types.CHANGE_ALGORITHM, algorithm })

// Query Graph
export const changeQueryGraph = queryGraph => ({ type: types.CHANGE_QUERY_GRAPH, queryGraph })
export const updateGraphState = graphState => ({ type: types.UPDATE_GRAPH_STATE, graphState })

// Routines
export const updateRoutines = routines => ({ type: types.UPDATE_ROUTINES, routines })

// Steps
export const increaseStep = increase => ({ type: types.INCREASE_STEP, increase })
export const decreaseStep = decrease => ({ type: types.DECREASE_STEP, decrease })
export const resetSteps = () => ({ type: types.RESET_STEPS })
export const updateSteps = steps => ({ type: types.UPDATE_STEPS, steps })
export const updateCurrentStepUUID = uuid => ({ type: types.UPDATE_STEP_UUID, uuid })

// Configuration (To be removed)
export const updateConfiguration = configuration => ({ type: types.UPDATE_CONFIGURATION, configuration })

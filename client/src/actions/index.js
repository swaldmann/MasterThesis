import * as types from '../constants/ActionTypes'

// Query Graph
export const changeQueryNumberOfRelations = numberOfRelations => ({ type: types.CHANGE_QUERY_NUMBER_RELATIONS, numberOfRelations })
export const changeQueryGraphTypeOption = graphTypeOptionValue => ({ type: types.CHANGE_QUERY_GRAPH_TYPE, graphTypeOptionValue })
export const changeQueryGraph = queryGraph => ({ type: types.CHANGE_QUERY_GRAPH, queryGraph })

// Algorithm
export const changeAlgorithm = algorithm => ({ type: types.CHANGE_ALGORITHM, algorithm })
export const updateAlgorithms = algorithms => ({ type: types.UPDATE_ALGORITHMS, algorithms })

// Variables
export const addVariableRow = variableRow => ({ type: types.ADD_VARIABLE_ROW, variableRow })
export const popVariableRow = () => ({ type: types.POP_VARIABLE_ROW })

// Configuration
export const updateConfiguration = configuration => ({ type: types.UPDATE_CONFIGURATION, configuration })

// Steps
export const increaseStep = increase => ({ type: types.INCREASE_STEP, increase })
export const decreaseStep = decrease => ({ type: types.DECREASE_STEP, decrease })
export const resetSteps = () => ({ type: types.RESET_STEPS })
export const updateSteps = steps => ({ type: types.UPDATE_STEPS, steps })
export const updateStepUUID = uuid => ({ type: types.UPDATE_STEP_UUID, uuid })

export const increaseRenderedStep = increase => ({ type: types.INCREASE_RENDERED_STEP, increase })
export const decreaseRenderedStep = decrease => ({ type: types.DECREASE_RENDERED_STEP, decrease })
export const resetRenderedStep = () => ({ type: types.RESET_RENDERED_STEP })

export const updateRoutines = routines => ({ type: types.UPDATE_ROUTINES, routines })
export const updateCurrentRoutine = routine => ({ type: types.UPDATE_CURRENT_ROUTINE, routine })

export const updateGraphState = graphState => ({ type: types.UPDATE_GRAPH_STATE, graphState })
package main

import (
	"image/color"

	"github.com/google/uuid"
)

// Visualizable Type conformance for visualizing join ordering/query graph algorithms
type Visualizable func(QG QueryGraph, JTC JoinTreeCreator) *Tree

// VisualizationOn Boolean indicating whether the visualization code should be executed
var VisualizationOn = false

var (
	// BlueColor Color constant for blue nodes
	BlueColor = color.RGBA{90, 165, 255, 1}
	// GreenColor constant for green nodes
	GreenColor = color.RGBA{85, 165, 34, 1}
	// GrayColor Color constant for gray nodes
	GrayColor = color.RGBA{120, 120, 120, 1}
	// WhiteColor Color constant for white nodes
	WhiteColor = color.RGBA{255, 255, 255, 1}
	// RedColor Color constant for red nodes
	RedColor = color.RGBA{255, 0, 0, 1}
	// OrangeColor Color constant for orange nodes
	OrangeColor = color.RGBA{235, 165, 50, 1}
)

var routines = []*VisualizationRoutine{}
var stack = []*VisualizationRoutine{}

func resetRoutines() {
	routines = []*VisualizationRoutine{}
}

func popStack() {
	stack = stack[:len(stack)-1]
}

// Visualize Wrapper function for calling functions in the visualization context
func Visualize(visualization Visualizable, QG QueryGraph, JTC JoinTreeCreator) []*VisualizationRoutine {
	resetRoutines()
	oldVisualizationOn := VisualizationOn
	VisualizationOn = true
	visualization(QG, JTC)
	VisualizationOn = oldVisualizationOn
	defer resetRoutines()
	return routines
}

// StartVisualizationRoutine Begin a new visualization routine
func StartVisualizationRoutine(routine *VisualizationRoutine) {
	stack = append(stack, routine)

	if len(stack) > 1 {
		currentStackIndex := len(stack) - 2
		currentRoutine := stack[currentStackIndex]
		var v interface{} = routine
		currentRoutine.Steps = append(currentRoutine.Steps, &v)
	} else {
		routines = append(routines, routine)
	}
}

// EndVisualizationRoutine End a visualization routine and optionally pass a result
func EndVisualizationRoutine(result ...*VisualizationRoutineResult) {
	//popStack()
	//currentRoutineIndex := len(routines) - 1
	//var v interface{}
	//v = result
	//routines[currentRoutineIndex].Steps = append(routines[currentRoutineIndex].Steps, &v)
}

// AddVisualizationStep Add a new atomic visualization step to the current routine
func AddVisualizationStep(QG QueryGraph, relations VariableTableRow) {
	n := uint(len(QG.R))

	nodeColors := []NodeColor{}

	// Color each node explicitly, not just changes.
	// This could possibly also be done for changes only,
	// but it's not easy to achieve modularity and
	// it's way harder to debug, both in the server and
	// client/visualization.
	currentStackIndex := len(stack) - 1

	// Create graph state
	observedRelations := stack[currentStackIndex].ObservedRelations
	for i := n - 1; int(i-1) >= -1; i-- {
		for _, relation := range observedRelations {
			relationIndexes := relations[relation.Identifier]
			if contains(relationIndexes, i) {
				color := relation.Color
				nodeColor := NodeColor{NodeIndex: i, Color: color}
				nodeColors = append(nodeColors, nodeColor)
			}
		}
	}
	graphState := GraphState{NodeColors: nodeColors}
	uuid := uuid.New().String()
	step := &VisualizationStep{GraphState: graphState, Variables: relations, UUID: uuid}

	currentRoutine := stack[currentStackIndex]
	var v interface{} = step
	currentRoutine.Steps = append(currentRoutine.Steps, &v)
}

// VariableTableEntry Entry in the variable table.
type VariableTableEntry []uint

// VariableTableRow Table holding values of variables over time.
type VariableTableRow map[string]VariableTableEntry

// Algorithm Defines an algorithm.
type Algorithm struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// GraphState Defines the state of a graph.
type GraphState struct {
	NodeColors []NodeColor `json:"nodeColors"`
}

// NodeColor Defines the color of a graph node.
type NodeColor struct {
	NodeIndex uint       `json:"nodeIndex"`
	Color     color.RGBA `json:"color"`
}

// ObservedRelation Represents an observed relation and its configuration.
type ObservedRelation struct {
	Identifier string     `json:"identifier"`
	Color      color.RGBA `json:"color"`
	// TODO: Description string `json:"description"`
}

// VisualizationRoutineResult A description of the visualization routine's return value
type VisualizationRoutineResult struct {
	Description string `json:"description"`
}

// VisualizationRoutine A top-level routine executed by the algorithm.
type VisualizationRoutine struct {
	Name              string             `json:"name"`
	ObservedRelations []ObservedRelation `json:"observedRelations"`
	Steps             []*interface{}     `json:"steps"`
}

// VisualizationStep An atomic visualization step that transfers the visualization to a new state.
type VisualizationStep struct {
	GraphState GraphState    `json:"graphState"`
	Variables  VariableTableRow `json:"variables"`
	UUID       string        `json:"uuid"`
}

/* ------------- */
/* To be removed */
/* ------------- */

// Configuration Specifies a visualization configuration
type Configuration struct {
	// No use case for that any more. Can probably be removed.
}

// AlgorithmCounter Defines a counter in an algorithm.
type AlgorithmCounter struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

package main

import "image/color"

// Visualizable Type conformance for visualizing join ordering/query graph algorithms
type Visualizable func(QG QueryGraph, JTC JoinTreeCreator) *Tree

var visualizationOn = false

var (
	blueColor   = color.RGBA{90, 165, 255, 1}
	greenColor  = color.RGBA{85, 165, 34, 1}
	grayColor   = color.RGBA{120, 120, 120, 1}
	whiteColor  = color.RGBA{255, 255, 255, 1}
	redColor    = color.RGBA{255, 0, 0, 1}
	orangeColor = color.RGBA{235, 165, 50, 1}
)

var routines = []VisualizationRoutine{}
var changes = []VisualizationStep{}
var currentRoutine VisualizationRoutine
var stack = []string{}

func resetChanges() {
	changes = []VisualizationStep{}
}

func resetRoutines() {
	routines = []VisualizationRoutine{}
}

func visualize(visualization Visualizable, QG QueryGraph, JTC JoinTreeCreator) []VisualizationRoutine {
	oldVisualizationOn := visualizationOn
	visualizationOn = true
	visualization(QG, JTC)
	visualizationOn = oldVisualizationOn
	defer resetChanges()
	defer resetRoutines()
	return routines
}

func visualizeRelations(QG QueryGraph, relations VariableTable, stack SubroutineStack) {
	n := uint(len(QG.R))

	nodeColors := []NodeColor{}

	// Color each node explicitly, not just changes.
	// This could possibly also be done for changes only,
	// but it's not easy to achieve modularity and
	// it's way harder to debug, both in the server and
	// client/visualization.
	observedRelations := currentRoutine.ObservedRelations
	for j := n - 1; int(j-1) >= -1; j-- {
		for _, relation := range observedRelations {
			relationIndexes := relations[relation.Identifier]
			if contains(relationIndexes, j) {
				nodeColor := relation.Color
				nodeConfiguration := NodeColor{NodeIndex: j, Color: nodeColor}
				nodeColors = append(nodeColors, nodeConfiguration)
			}
		}
	}
	graphState := GraphState{NodeColors: nodeColors}
	change := VisualizationStep{GraphState: graphState, Variables: relations, SubroutineStack: stack}
	changes = append(changes, change)
}

// SubroutineStack Description of the current (recursive) call stack.
type SubroutineStack []string

// VariableTableEntry Entry in the variable table.
type VariableTableEntry []uint

// VariableTable Table holding values of variables over time.
type VariableTable map[string]VariableTableEntry

// Algorithm Defines an algorithm.
type Algorithm struct {
	Name       string       `json:"name"`
	GraphState []GraphState `json:"graphState"`
	//Counters   []AlgorithmCounter `json:"counters"`
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

// VisualizationRoutine A top-level routine executed by the algorithm.
type VisualizationRoutine struct {
	Name              string              `json:"name"`
	ObservedRelations []ObservedRelation  `json:"observedRelations"`
	Steps             []VisualizationStep `json:"steps"`
}

// VisualizationStep An atomic visualization step that transfers the visualization to a new state.
type VisualizationStep struct {
	GraphState      GraphState      `json:"graphState"`
	SubroutineStack SubroutineStack `json:"subroutineStack"`
	Variables       VariableTable   `json:"variables"`
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

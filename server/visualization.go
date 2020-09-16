package main

import "image/color"

// Visualizable Protocol for visualizing query graph algorithms
type Visualizable func(QG QueryGraph, JTC JoinTreeCreator) *Tree

var visualizationOn = false

var blueColor = color.RGBA{0, 0, 255, 1}
var greenColor = color.RGBA{85, 165, 34, 1}
var grayColor = color.RGBA{120, 120, 120, 1}
var whiteColor = color.RGBA{255, 255, 255, 1}
var redColor = color.RGBA{255, 0, 0, 1}

var changes = []interface{}{}

func resetChanges() {
	changes = []interface{}{}
}

func visualize(visualization Visualizable, QG QueryGraph, JTC JoinTreeCreator) {
	oldVisualizationOn := visualizationOn
	visualizationOn = true
	visualization(QG, JTC)
	visualizationOn = oldVisualizationOn
}

// VariableTableEntry Entry in the debug variable table
type VariableTableEntry []uint

// VariableTable Table holding values of variables over time
type VariableTable map[string]VariableTableEntry

// AlgorithmCounter Defines a counter in an algorithm.
type AlgorithmCounter struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// Algorithm Defines an algorithm.
type Algorithm struct {
	Name       string             `json:"name"`
	Counters   []AlgorithmCounter `json:"counters"`
	GraphState []GraphState       `json:"graphState"`
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

// Configuration Specifies a visualization configuration
type Configuration struct {
	ObserverdVariables []string `json:"observedVariables"`
}

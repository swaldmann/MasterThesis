package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"reflect"

	rainbow "github.com/fatih/color"
)

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

var steps = []interface{}{}
var routines = []*VisualizationRoutine{}
var stack = []*VisualizationRoutine{}

func resetAllSteps() {
	steps = []interface{}{}
}

func resetSteps(routineKey string) {
	//delete(steps, routineKey)
}

func resetRoutines() {
	routines = []*VisualizationRoutine{}
}

func popStack() {
	stack = stack[:len(stack)-1]
}

func visualize(visualization Visualizable, QG QueryGraph, JTC JoinTreeCreator) []*VisualizationRoutine {
	oldVisualizationOn := visualizationOn
	visualizationOn = true
	visualization(QG, JTC)
	visualizationOn = oldVisualizationOn
	defer resetAllSteps()
	defer resetRoutines()
	return routines
}

func startVisualizeRoutine(routine *VisualizationRoutine) {
	stack = append(stack, routine)

	if len(stack) > 1 {
		//currentStackIndex := len(stack) - 1
		currentRoutineIndex := len(routines) - 1
		var v interface{}
		v = routine
		routines[currentRoutineIndex].Steps = append(routines[currentRoutineIndex].Steps, &v)
	} else {
		routines = append(routines, routine)
	}
}

func recursivelyAppendToSteps(steps *[]*interface{}, step *interface{}) {
	*steps = append(*steps, step)
}

func addVisualizationStep(QG QueryGraph, relations VariableTable) {
	n := uint(len(QG.R))

	nodeColors := []NodeColor{}

	// Color each node explicitly, not just changes.
	// This could possibly also be done for changes only,
	// but it's not easy to achieve modularity and
	// it's way harder to debug, both in the server and
	// client/visualization.
	currentStackIndex := len(stack) - 1
	//currentRoutineIndex := len(routines) - 1

	// Create graph state
	observedRelations := stack[currentStackIndex].ObservedRelations
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
	step := &VisualizationStep{GraphState: graphState, Variables: relations}

	currentRoutine := stack[currentStackIndex]
	var v interface{}
	v = step

	fmt.Println(&routines[0])
	fmt.Println(&currentRoutine)
	fmt.Println(currentRoutine)
	fmt.Println(step)
	//recursivelyAppendToSteps(&currentRoutine.Steps, &v)
	currentRoutine.Steps = append(currentRoutine.Steps, &v)

	fmt.Println(routines)
	bolB, _ := json.Marshal(routines)
	fmt.Println(string(bolB))
	rainbow.Blue("=========")
	/*rainbow.Blue("BEGIN ======")
	fmt.Println(currentRoutine)
	fmt.Println(&currentRoutine)
	fmt.Println(relations)
	fmt.Println(currentRoutine.Name)
	fmt.Println(currentStackIndex)
	for i := 0; i < currentStackIndex+1; i++ {

		stepLength := len(currentRoutine.Steps)
		if i == currentStackIndex || stepLength == 0 {
			rainbow.Green("Append to " + currentRoutine.Name)
			fmt.Println(step)
			fmt.Println(&currentRoutine)
			rainbow.Green("Before")
			fmt.Println(currentRoutine.Steps)
			//recursivelyAppendToSteps(currentRoutine.Steps, *step)
			var v interface{}
			v = step
			currentRoutine.Steps = append(currentRoutine.Steps, &v)
			rainbow.Green("After")
			fmt.Println(&currentRoutine.Steps)
			rainbow.Yellow("All routines")
			fmt.Println(routines)
			bolB, _ := json.Marshal(routines)
			fmt.Println(string(bolB))
			break
		}

		fmt.Println("---> " + currentRoutine.Name)
		//fmt.Println(currentRoutine)

		steps := currentRoutine.Steps
		l := *steps[stepLength-1]
		lastStep := to_struct_ptr(*steps[stepLength-1])
		fmt.Println(steps)
		fmt.Println(l)
		fmt.Println(reflect.TypeOf(l))
		fmt.Println(&lastStep)
		fmt.Println(lastStep)
		fmt.Println(reflect.TypeOf(lastStep))
		if v, ok := (lastStep).(*VisualizationRoutine); ok {
			rainbow.Red("Before")
			//fmt.Println(currentRoutine)
			fmt.Println(currentRoutine)
			//fmt.Println(*currentRoutine)
			//fmt.Println(currentRoutine.Name)
			currentRoutine = v
			rainbow.Red("New")
			//fmt.Println(v)
			fmt.Println(&v)
			rainbow.Red("After")
			fmt.Println(currentRoutine)
			fmt.Println(&currentRoutine)
			fmt.Println(*currentRoutine)
			//fmt.Println(currentRoutine.Name)
		}
	}
	rainbow.Blue("END ========")

	//stack[currentStackIndex].Steps = append(stack[currentStackIndex].Steps, step)
	//routines[currentRoutineIndex].Steps = append(routines[currentRoutineIndex].Steps, stack[currentStackIndex].Steps...)*/
}

//
// Return a pointer to the supplied struct via interface{}
//
func to_struct_ptr(obj interface{}) interface{} {
	vp := reflect.New(reflect.TypeOf(obj))
	vp.Elem().Set(reflect.ValueOf(obj))
	return vp.Interface()
}

// SubroutineStack Description of the current (recursive) call stack.
//type SubroutineStack []string

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
	Name              string             `json:"name"`
	ObservedRelations []ObservedRelation `json:"observedRelations"`
	Steps             []*interface{}     `json:"steps"`
}

// VisualizationStep An atomic visualization step that transfers the visualization to a new state.
type VisualizationStep struct {
	GraphState GraphState `json:"graphState"`
	//SubroutineStack []string      `json:"subroutineStack"`
	Variables VariableTable `json:"variables"`
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

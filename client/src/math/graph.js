class Graph {
    constructor(numberOfVertices) {
        this.numberOfVertices = numberOfVertices
        this.adjacencyList = new Map()
    }

    addVertex(v) {
        this.adjacencyList.set(v, [])    
    }

    addEdge(v, w) {
        this.adjacencyList.get(v).push(w)
        this.adjacencyList.get(w).push(v)
    }

    toString() {
        this.adjacencyList.toString()
    }
}

export default Graph
const π = Math.PI
const r_node = 20
const margin = 12
class QueryGraph {

    _nodeColors = []
    _numberOfNodes = 0

    constructor(numberOfNodes) {
        this._numberOfNodes = numberOfNodes
        this._nodeColors = Array(numberOfNodes).fill("white")  
    }
    
    draw(type, nodeColors, ctx) {
        console.log("Yo");
        console.log(nodeColors);

        nodeColors.forEach(nodeColor => {
            this._nodeColors[nodeColor.nodeIndex] = "rgb(" + nodeColor.color.R + "," + nodeColor.color.G + "," + nodeColor.color.B + ")"
        })

        const dpr = window.devicePixelRatio || 1
        const rect = ctx.canvas.getBoundingClientRect()
        ctx.canvas.width = rect.width * dpr
        ctx.canvas.height = rect.height * dpr
        ctx.clearRect(0, 0, ctx.canvas.clientWidth, ctx.canvas.clientHeight)
        ctx.scale(dpr, dpr)

        ctx.strokeStyle = 'rgb(250, 250, 250)'
        ctx.fillStyle = 'rgb(250, 250, 250)'

        const numberOfNodes = this._numberOfNodes
        switch(type) {
            case "chain":
                this._drawChainQuery(numberOfNodes, ctx)
                break
            case "star":
                this._drawStarQuery(numberOfNodes, ctx)
                break
            case "tree":
                this._drawTreeQuery(numberOfNodes, ctx)
                break
            case "cyclic":
                this._drawCyclicQuery(numberOfNodes, ctx)
                break
            case "cycle":
                this._drawCycleQuery(numberOfNodes, ctx)
                break    
            case "grid":
                this._drawGridQuery(numberOfNodes, ctx)
                break
            case "clique":
                this._drawCliqueQuery(numberOfNodes, ctx)
                break
            default:
                break
          } 
    }

    _drawChainQuery(numberOfNodes, ctx) {
        for (let i = 0; i < numberOfNodes; i++) {
            const drawableWidth = ctx.canvas.clientWidth - margin * 2 - r_node * 2
            const x = margin + r_node + i * drawableWidth/(numberOfNodes - 1)
            const y = 200
            this._drawNode(i, x, y, "white", ctx)

            // Connect the lines
            if (i !== 0) {
                const x_previous = margin + r_node + (i - 1) * drawableWidth/(numberOfNodes - 1)
                const y_previous = y
                this._drawLine(x, y, x_previous, y_previous, ctx)
            }
        }
    }

    _drawStarQuery(numberOfNodes, ctx) {
        for (let i = 0; i < numberOfNodes; i++) {
            const x_center = ctx.canvas.clientWidth/2
            const y_center = ctx.canvas.clientHeight/2                

            if (i === 0) { // Center node
                this._drawNode(i, x_center, y_center, "white", ctx)
                continue
            }
            const θ = 2 * π/(numberOfNodes - 1) * i
            const r = ctx.canvas.clientWidth/2 - r_node - margin
            const x = r * Math.cos(θ) + r + r_node + margin
            const y = r * Math.sin(θ) + r + r_node + margin
            this._drawLine(x, y, x_center, y_center, ctx)
            this._drawNode(i, x, y, "white", ctx)
        }
    }

    _drawTreeQuery(numberOfNodes, ctx) {
        const drawableWidth = ctx.canvas.clientWidth - margin * 2 - r_node * 2
        const drawableHeight = ctx.canvas.clientWidth - margin * 2 - r_node * 2
        const x_center = ctx.canvas.clientWidth/2
        const y_offset = drawableHeight/Math.floor(Math.log2(numberOfNodes))

        for (let i = 0; i < numberOfNodes; i++) {
            if (i === 0) { // Root node
                this._drawNode(i, x_center, margin + r_node, "white", ctx)
                continue
            }
            const calculateRow = index => Math.floor(Math.log2(index + 1))
            const row = calculateRow(i)
            const calculateNumberOfColumnsInRow = row => 2 ** row
            const numberOfColumnsInRow = calculateNumberOfColumnsInRow(row)

            const calculateColumn = (index, row) => index - 2 ** row + 1
            const calculateX = (column, numberOfColumnsInRow) => (column + 0.5) * drawableWidth/numberOfColumnsInRow + margin + r_node
            const calculateY = row => r_node + margin + y_offset * row
            const column = calculateColumn(i, row)
            const x = calculateX(column, numberOfColumnsInRow)
            const y = calculateY(row)
            this._drawNode(i, x, y, "white", ctx)

            const parentIndex = Math.floor((i - 1)/2)
            const row_previous = calculateRow(parentIndex)
            const numberOfColumns_previous = calculateNumberOfColumnsInRow(row_previous)
            const column_previous = calculateColumn(parentIndex, row_previous)
            const x_previous = calculateX(column_previous, numberOfColumns_previous)
            const y_previous = calculateY(row_previous)
            this._drawLine(x, y, x_previous, y_previous, ctx)
        }
    }

    _drawCyclicQuery(numberOfNodes, ctx) {

    }

    _drawCycleQuery(numberOfNodes, ctx) {
        for (let i = 0; i < numberOfNodes; i++) {
            const θ = 2 * π/numberOfNodes * i
            
            const r = ctx.canvas.clientWidth/2 - r_node - margin
            const x = r * Math.cos(θ) + r + r_node + margin
            const y = r * Math.sin(θ) + r + r_node + margin
            this._drawNode(i, x, y, "white", ctx)

            const θ_previous = 2 * π/numberOfNodes * (i - 1)
            const x_previous = r * Math.cos(θ_previous) + r + r_node + margin
            const y_previous = r * Math.sin(θ_previous) + r + r_node + margin
            this._drawLine(x, y, x_previous, y_previous, ctx)
        }
    }

    _drawGridQuery(numberOfNodes, ctx) {

    }

    _drawCliqueQuery(numberOfNodes, ctx) {

    }

    _drawNode(i, x, y, color, ctx) {
        ctx.fillStyle = this._nodeColors[i]
        ctx.beginPath()
        ctx.arc(x, y, r_node, 0, 2 * π)
        ctx.fill()

        ctx.font = "20px sans-serif";
        ctx.fillStyle = "rgb(30, 33, 39)"
        ctx.textBaseline = "middle"
        ctx.textAlign = "center"
        ctx.fillText("R" + i, x, y)
    }

    _drawLine(x_source, y_source, x_dest, y_dest, ctx) {
        const compositeOperationBefore = ctx.globalCompositeOperation
        ctx.globalCompositeOperation = 'destination-over'
        ctx.beginPath()
        ctx.moveTo(x_source, y_source)
        ctx.lineTo(x_dest, y_dest)
        ctx.stroke()
        ctx.globalCompositeOperation = compositeOperationBefore
    }
}

export default QueryGraph
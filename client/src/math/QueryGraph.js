const π = Math.PI
const r_node = 20
const margin = 24

CanvasRenderingContext2D.prototype.roundRect = function (x, y, w, h, r) {
    if (w < 2 * r) r = w / 2;
    if (h < 2 * r) r = h / 2;
    this.beginPath();
    this.moveTo(x+r, y);
    this.arcTo(x+w, y,   x+w, y+h, r);
    this.arcTo(x+w, y+h, x,   y+h, r);
    this.arcTo(x,   y+h, x,   y,   r);
    this.arcTo(x,   y,   x+w, y,   r);
    this.closePath();
    return this;
}
class QueryGraph {

    _queryGraph = {}
    _nodeColors = []
    _numberOfNodes = 0

    constructor(queryGraph) {
        const numberOfNodes = queryGraph.relationCardinalities.length
        this._queryGraph = queryGraph
        this._numberOfNodes = numberOfNodes
        this._nodeColors = Array(numberOfNodes).fill("white")  
    }
    
    draw(type, nodeColors, ctx) {
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
            case "moerkotte":
                this._drawMoerkotteQuery(ctx)
                break
            default:
                break
          } 
    }

    _drawChainQuery(numberOfNodes, ctx) {
        for (let i = 0; i < numberOfNodes; i++) {
            const cardinality = this._queryGraph.relationCardinalities[i]
            const drawableWidth = ctx.canvas.clientWidth - margin * 2 - r_node * 2
            const x = margin + r_node + i * drawableWidth/(numberOfNodes - 1)
            const y = 200
            this._drawNode(i, x, y, "white", ctx, cardinality)

            // Connect the lines
            if (i !== 0) {
                const x_previous = margin + r_node + (i - 1) * drawableWidth/(numberOfNodes - 1)
                const y_previous = y
                const selectivityKey = (1 << (i-1)) | (1 << i)
                const selectivity = this._queryGraph.selectivities[selectivityKey]
                this._drawEdge(x, y, x_previous, y_previous, ctx, selectivity)
            }
        }
    }

    _drawStarQuery(numberOfNodes, ctx) {
        for (let i = 0; i < numberOfNodes; i++) {
            const cardinality = this._queryGraph.relationCardinalities[i]
            const x_center = ctx.canvas.clientWidth/2
            const y_center = ctx.canvas.clientHeight/2

            if (i === 0) { // Center node
                this._drawNode(i, x_center, y_center, "white", ctx, cardinality)
                continue
            }
            const θ = 2 * π/(numberOfNodes - 1) * i
            const r = ctx.canvas.clientWidth/2 - r_node - margin
            const x = r * Math.cos(θ) + r + r_node + margin
            const y = r * Math.sin(θ) + r + r_node + margin

            const selectivityKey = 1 | (1 << i)
            const selectivity = this._queryGraph.selectivities[selectivityKey]
            this._drawEdge(x, y, x_center, y_center, ctx, selectivity)
            this._drawNode(i, x, y, "white", ctx, cardinality)
        }
    }

    _drawTreeQuery(numberOfNodes, ctx) {
        const drawableWidth = ctx.canvas.clientWidth - margin * 2 - r_node * 2
        const drawableHeight = ctx.canvas.clientWidth - margin * 2 - r_node * 2
        const x_center = ctx.canvas.clientWidth/2
        const y_offset = drawableHeight/Math.floor(Math.log2(numberOfNodes))

        for (let i = 0; i < numberOfNodes; i++) {
            const cardinality = this._queryGraph.relationCardinalities[i]
            if (i === 0) { // Root node
                this._drawNode(i, x_center, margin + r_node, "white", ctx, cardinality)
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
            this._drawNode(i, x, y, "white", ctx, cardinality)

            const parentIndex = Math.floor((i - 1)/2)
            const row_previous = calculateRow(parentIndex)
            const numberOfColumns_previous = calculateNumberOfColumnsInRow(row_previous)
            const column_previous = calculateColumn(parentIndex, row_previous)
            const x_previous = calculateX(column_previous, numberOfColumns_previous)
            const y_previous = calculateY(row_previous)

            const selectivityKey = (1 << parentIndex) | (1 << i)
            const selectivity = this._queryGraph.selectivities[selectivityKey]
            this._drawEdge(x, y, x_previous, y_previous, ctx, selectivity)
        }
    }

    // Hard-coded special case of a cyclic query
    _drawMoerkotteQuery(ctx) {
        const width = ctx.canvas.clientWidth
        const height = ctx.canvas.clientHeight
        const centerX = ctx.canvas.clientWidth/2
        const centerY = ctx.canvas.clientHeight/2
        const offset = margin + r_node

        const x0 = centerX, y0 = offset
        const x1 = offset, y1 = centerY
        const x2 = centerX, y2 = centerY
        const x3 = width - offset, y3 = centerY
        const x4 = centerX, y4 = height - offset

        this._drawNode(0, x0, y0, "white", ctx, this._queryGraph.relationCardinalities[0])
        this._drawNode(1, x1, y1, "white", ctx, this._queryGraph.relationCardinalities[1])
        this._drawNode(2, x2, y2, "white", ctx, this._queryGraph.relationCardinalities[2])
        this._drawNode(3, x3, y3, "white", ctx, this._queryGraph.relationCardinalities[3])
        this._drawNode(4, x4, y4, "white", ctx, this._queryGraph.relationCardinalities[4])
        
        this._drawEdge(x0, y0, x1, y1, ctx, this._queryGraph.selectivities[(1 << 0) | (1 << 1)])
        this._drawEdge(x0, y0, x2, y2, ctx, this._queryGraph.selectivities[(1 << 0) | (1 << 2)])
        this._drawEdge(x0, y0, x3, y3, ctx, this._queryGraph.selectivities[(1 << 0) | (1 << 3)])
        this._drawEdge(x1, y1, x4, y4, ctx, this._queryGraph.selectivities[(1 << 1) | (1 << 4)])
        this._drawEdge(x2, y2, x3, y3, ctx, this._queryGraph.selectivities[(1 << 2) | (1 << 3)])
        this._drawEdge(x2, y2, x4, y4, ctx, this._queryGraph.selectivities[(1 << 2) | (1 << 4)])
        this._drawEdge(x3, y3, x4, y4, ctx, this._queryGraph.selectivities[(1 << 3) | (1 << 4)])
    }

    _drawCycleQuery(numberOfNodes, ctx) {
        for (let i = 0; i < numberOfNodes; i++) {
            const cardinality = this._queryGraph.relationCardinalities[i]
            const θ = 2 * π/numberOfNodes * i
            
            const r = ctx.canvas.clientWidth/2 - r_node - margin
            const x = r * Math.cos(θ) + r + r_node + margin
            const y = r * Math.sin(θ) + r + r_node + margin
            this._drawNode(i, x, y, "white", ctx, cardinality)

            const θ_next = 2 * π/numberOfNodes * (i + 1)
            const x_next = r * Math.cos(θ_next) + r + r_node + margin
            const y_next = r * Math.sin(θ_next) + r + r_node + margin

            const selectivityKey = (1 << ((i+1) % numberOfNodes)) | (1 << i)
            const selectivity = this._queryGraph.selectivities[selectivityKey]
            this._drawEdge(x, y, x_next, y_next, ctx, selectivity)
        }
    }

    // The number of nodes is unfortunately not
    // sufficient to draw those. Might need more 
    // parameters.
    _drawCyclicQuery(numberOfNodes, ctx) { }

    _drawGridQuery(numberOfNodes, ctx) { }

    _drawCliqueQuery(numberOfNodes, ctx) { }

    _drawNode(i, x, y, color, ctx, label) {
        ctx.fillStyle = this._nodeColors[i]
        ctx.beginPath()
        ctx.arc(x, y, r_node, 0, 2 * π)
        ctx.fill()

        ctx.font = "20px sans-serif"
        ctx.fillStyle = "rgb(30, 33, 39)"
        ctx.textBaseline = "middle"
        ctx.textAlign = "center"
        ctx.fillText("R" + i, x, y)

        const oldFillColor = ctx.fillColor
        ctx.fillStyle = "#555"
        ctx.font = "12px sans-serif"
        const textWidth = ctx.measureText(label).width
        ctx.roundRect(x - textWidth/2 - 3 + 10, y + 12.4, textWidth + 6, 18, 3).fill()
        ctx.fillStyle = "white"
        ctx.fillText(label, x + 10, y + 22)
        ctx.fillStyle = oldFillColor
    }

    _drawEdge(x_source, y_source, x_dest, y_dest, ctx, weight = null) {
        const compositeOperationBefore = ctx.globalCompositeOperation
        ctx.globalCompositeOperation = 'destination-over'
        ctx.beginPath()
        ctx.moveTo(x_source, y_source)
        ctx.lineTo(x_dest, y_dest)
        ctx.stroke()
        ctx.globalCompositeOperation = compositeOperationBefore
        if (weight) {
            const weightLabelCenterX = Math.min(x_dest, x_source) + Math.abs(x_dest - x_source)/2
            const weightLabelCenterY = Math.min(y_dest, y_source) + Math.abs(y_dest - y_source)/2
            const oldFillStyle = ctx.fillStyle

            const labelText = [...weight.toString()].slice(0,7).join("") // Limit to a maximum of 5 characters
            const labelWidth = ctx.measureText(labelText).width + 10
            const labelHeight = 18

            ctx.fillStyle = "white"
            ctx.fillRect(weightLabelCenterX - labelWidth/2, weightLabelCenterY - labelHeight/2, labelWidth, labelHeight)
            ctx.fillStyle = oldFillStyle

            ctx.font = "14px sans-serif"
            ctx.fillStyle = "#222"
            ctx.fillText(labelText, weightLabelCenterX, weightLabelCenterY)
            ctx.moveTo(weightLabelCenterX, weightLabelCenterY)
        }
    }

    _indexesOfSetBits(S) {
        const reversedBitArray = [...S.toString(2)].reverse()
        const concatIfOne = (result, c, i) => c === "1" ? result.concat(i) : result
        return reversedBitArray.reduce(concatIfOne, [])
    }
}

export default QueryGraph
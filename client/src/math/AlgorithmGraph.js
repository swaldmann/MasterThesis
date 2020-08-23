const π = Math.PI

class AlgorithmGraph {
    draw(ctx) {
        const dpr = window.devicePixelRatio || 1
        const rect = ctx.canvas.getBoundingClientRect()
        ctx.canvas.width = rect.width * dpr
        ctx.canvas.height = rect.height * dpr
        ctx.clearRect(0, 0, ctx.canvas.clientWidth, ctx.canvas.clientHeight)
        ctx.scale(dpr, dpr)

        ctx.strokeStyle = 'rgb(250, 250, 250)'
        ctx.fillStyle = 'rgb(250, 250, 250)'
        
        ctx.beginPath()
        ctx.moveTo(20, 20)
        ctx.lineTo(20, 100)
        ctx.stroke()
    }
}

export default AlgorithmGraph
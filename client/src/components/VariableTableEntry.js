import React from "react"

class VariableTableEntry extends React.Component {
    render() {
        const { routine, parent, step, level } = this.props

        const rBody = 40
        const gBody = 44
        const bBody = 52
        const rParent = rBody - 8 * (level - 1)
        const gParent = gBody - 8 * (level - 1)
        const bParent = bBody - 8 * (level - 1)
        const parentStyle = { backgroundColor: "rgba(" + rParent + "," + gParent + "," + bParent + ", 1)" }
        const r = rBody - 8 * level
        const g = gBody - 8 * level
        const b = bBody - 8 * level
        const style = { backgroundColor: "rgba(" + r + "," + g + "," + b + ", 1)" }
        const borderLeft = 30 * level + "px solid #282c34"
        const marginLeft = "-" + 30 * level + "px"
        const fakeInsetStyle = { borderLeft: borderLeft, marginLeft: marginLeft }

        // Routine
        if (step.name) {
            return (
                <>
                    <tr style={style}>
                        <td style={{ fakeInsetStyle }} colSpan="100%">{parent && parent.name === step.name && Array(level - 1).fill("→").join("") + " "}{step.name}</td>
                    </tr>
                    <tr style={style}>
                        {
                            step.observedRelations.map((r, i) => {
                                const color = "rgba(" + r.color.R + "," + r.color.G + "," + r.color.B + "," + r.color.A + ")"
                                const borderLeft = i === 0 ? 30 * level + "px solid #282c34" : "0px"
                                return (
                                    <td style={{ color: color/*, borderLeft: borderLeft, marginLeft: marginLeft*/ }}>
                                        {r.identifier}
                                    </td>
                                )
                            })}
                    </tr>
                    {
                        step.steps && step.steps.map(s =>
                            <VariableTableEntry parent={step} step={s} level={level + 1} />
                        )
                    }
                </>
            )
        }
        else if (step.steps) {
            return step.steps.map(s =>
                <VariableTableEntry parent={step} step={s} level={level + 1} />
            )
        }
        else if (step.variables) {
            return (<tr style={parentStyle}>{parent.observedRelations.map(r => <td key={r.identifier} >{step.variables[r.identifier] ? step.variables[r.identifier].join(",") : ""}</td>)}</tr>)
        }
        else if (step.description) {
            return (<tr><td>Result for {parent.name}: {step.description}</td></tr>)
        }
        return (<div className="error">There has been an error in the produced JSON output.</div>) // Should never happen if a valid JSON is produced
    }
}

export default VariableTableEntry
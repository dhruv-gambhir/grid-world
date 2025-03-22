package main

import (
    "fmt"

    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/plotutil"
    "gonum.org/v1/plot/vg"
)

// PlotValueIterationPerState draws one line per cell's utility
// across all Value Iteration iterations, assigning each line a color.
func PlotValueIterationPerState() {
    p := plot.New()
    p.Title.Text = "Value Iteration: Utility per State"
    p.X.Label.Text = "Iteration"
    p.Y.Label.Text = "Utility"

    lineIndex := 0
    // For each cell, build a line series of utility across iterations
    for r := 0; r < numRows; r++ {
        for c := 0; c < numCols; c++ {
            if Walls[[2]int{r, c}] {
                continue // skip walls
            }

            // Build XY points for this (r,c)
            pts := make(plotter.XYs, len(valueIterationHistory))
            for i := 0; i < len(valueIterationHistory); i++ {
                pts[i].X = float64(i)
                pts[i].Y = valueIterationHistory[i][r][c]
            }

            line, err := plotter.NewLine(pts)
            if err != nil {
                fmt.Println("Error creating line:", err)
                return
            }

            // Assign a unique color from the built-in palette
            line.Color = plotutil.Color(lineIndex)

            // Label in the legend with (row,col)
            label := fmt.Sprintf("(%d,%d)", r, c)
            p.Add(line)
            p.Legend.Add(label, line)

            lineIndex++
        }
    }

    // Adjust legend position if you prefer:
    p.Legend.Top = true
    p.Legend.Left = true

    // Save to PNG
    if err := p.Save(8*vg.Inch, 5*vg.Inch, "vi_states.png"); err != nil {
        fmt.Println("Error saving VI plot:", err)
    } else {
        fmt.Println("Saved vi_states.png")
    }
}

// PlotPolicyIterationPerState draws one line per cell’s utility
// across all Policy Iteration iterations, assigning each line a color.
func PlotPolicyIterationPerState() {
    p := plot.New()
    p.Title.Text = "Policy Iteration: Utility per State"
    p.X.Label.Text = "Iteration"
    p.Y.Label.Text = "Utility"

    lineIndex := 0
    for r := 0; r < numRows; r++ {
        for c := 0; c < numCols; c++ {
            if Walls[[2]int{r, c}] {
                continue
            }

            pts := make(plotter.XYs, len(policyIterationHistory))
            for i := 0; i < len(policyIterationHistory); i++ {
                pts[i].X = float64(i)
                pts[i].Y = policyIterationHistory[i][r][c]
            }

            line, err := plotter.NewLine(pts)
            if err != nil {
                fmt.Println("Error creating line:", err)
                return
            }

            // Assign a color
            line.Color = plotutil.Color(lineIndex)

            label := fmt.Sprintf("(%d,%d)", r, c)
            p.Add(line)
            p.Legend.Add(label, line)

            lineIndex++
        }
    }

    p.Legend.Top = true
    p.Legend.Left = true

    if err := p.Save(8*vg.Inch, 5*vg.Inch, "pi_states.png"); err != nil {
        fmt.Println("Error saving PI plot:", err)
    } else {
        fmt.Println("Saved pi_states.png")
    }
}


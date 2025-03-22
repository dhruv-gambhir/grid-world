package main

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Plot function
func PlotValueIteration() {
	p := plot.New()
	p.Title.Text = "Value Iteration"
	p.X.Label.Text = "Iterations"
	p.Y.Label.Text = "Utility Estimates"

	lineIndex := 0
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if Walls[[2]int{r, c}] {
				continue // skipping walls
			}

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

			line.Color = plotutil.Color(lineIndex)

			label := fmt.Sprintf("(%d,%d)", r, c)
			p.Add(line)
			p.Legend.Add(label, line)

			lineIndex++
		}
	}

	p.Legend.Top = true
	p.Legend.Left = true

	if err := p.Save(8*vg.Inch, 5*vg.Inch, "value_iteration.png"); err != nil {
		fmt.Println("Error saving value_iteration plot:", err)
	} else {
		fmt.Println("Saved value_iteration.png")
	}
    fmt.Println()
}

// Plot Function
func PlotPolicyIteration() {
	p := plot.New()
	p.Title.Text = "Policy Iteration"
	p.X.Label.Text = "Iterations"
	p.Y.Label.Text = "Utility Estimates"

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

			line.Color = plotutil.Color(lineIndex)

			label := fmt.Sprintf("(%d,%d)", r, c)
			p.Add(line)
			p.Legend.Add(label, line)

			lineIndex++
		}
	}

	p.Legend.Top = true
	p.Legend.Left = true

	if err := p.Save(8*vg.Inch, 5*vg.Inch, "policy_iteration.png"); err != nil {
		fmt.Println("Error saving policy_iteration plot:", err)
	} else {
		fmt.Println("saved policy_iteration.png")
        fmt.Println()
	}
}

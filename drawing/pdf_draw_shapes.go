/*
 * Draw multiple shapes in a new PDF file.
 *
 * Run as: go run pdf_draw_shapes.go
 */

package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/contentstream/draw"
	"github.com/unidoc/unipdf/v3/creator"
)

const licenseKey = `
-----BEGIN UNIDOC LICENSE KEY-----
Free trial license keys are available at: https://unidoc.io/
-----END UNIDOC LICENSE KEY-----
`

func init() {
	// Enable debug-level logging.
	// unicommon.SetLogger(unicommon.NewConsoleLogger(unicommon.LogLevelDebug))

	err := license.SetLicenseKey(licenseKey, `Company Name`)
	if err != nil {
		panic(err)
	}
}

func main() {
	// New creator with default properties (pagesize letter default).
	c := creator.New()
	c.NewPage()

	err := drawLine(c)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = drawPolyLine(c)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = drawRectangle(c)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = drawEllipse(c)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = drawCurve(c)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = drawPolyBezierCurve(c)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = drawCurvePolygon(c)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = drawPolygon(c)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	c.WriteToFile("output.pdf")

	fmt.Println("Completed, output saved to output.pdf file")
}

func drawHeader(c *creator.Creator, x, y float64, title string) error {
	p := c.NewStyledParagraph()
	p.SetPos(x, y)

	text := p.SetText(title)
	text.Style.FontSize = 20

	return c.Draw(p)
}

// Draw a single line
func drawLine(c *creator.Creator) error {
	drawHeader(c, 30, 10, "-- Line --")

	line := c.NewLine(50, 50, 100, 130)
	line.SetLineWidth(1.5)
	// Draw a red line, use hex color util to get r,g,b codes from html hex color.
	red := creator.ColorRGBFromHex("#ff0000")
	line.SetColor(red)

	return c.Draw(line)
}

// Draw line polygon.
func drawPolyLine(c *creator.Creator) error {
	drawHeader(c, 180, 10, "-- Polyline --")

	poly := c.NewPolyline([]draw.Point{
		draw.NewPoint(220, 50),
		draw.NewPoint(240, 90),
		draw.NewPoint(260, 100),
		draw.NewPoint(280, 130),
	})

	// Set line color and width
	poly.SetLineColor(creator.ColorBlue)
	poly.SetLineWidth(2)

	return c.Draw(poly)
}

// Draw a rectangle
func drawRectangle(c *creator.Creator) error {
	drawHeader(c, 420, 10, "-- Rectangle --")

	rect := c.NewRectangle(420, 50, 100, 100)

	// Set fill color.
	rect.SetFillColor(creator.ColorGreen)

	// Set border color.
	rect.SetBorderColor(creator.ColorRed)

	// Set border width
	rect.SetBorderWidth(2)

	return c.Draw(rect)
}

func drawEllipse(c *creator.Creator) error {
	drawHeader(c, 30, 200, "-- Ellipse --")

	ellipse := c.NewEllipse(80, 260, 100, 50)

	// Set border and fill color.
	ellipse.SetBorderColor(creator.ColorBlack)
	ellipse.SetFillColor(creator.ColorYellow)

	return c.Draw(ellipse)
}

// Draw a simple curve.
func drawCurve(c *creator.Creator) error {
	drawHeader(c, 180, 200, "-- Curve --")

	curve := c.NewCurve(220, 240, 200, 280, 300, 320)
	curve.SetColor(creator.ColorRed)
	curve.SetWidth(2)

	return c.Draw(curve)
}

// Draw using bezier curve.
func drawPolyBezierCurve(c *creator.Creator) error {
	drawHeader(c, 380, 200, "-- Poly-Bezier Curve -- ")

	curve := c.NewPolyBezierCurve([]draw.CubicBezierCurve{
		draw.NewCubicBezierCurve(450, 300, 478, 284, 505, 310, 500, 340), // top right
		draw.NewCubicBezierCurve(500, 340, 500, 380, 479, 420, 450, 400), // bottom right
		draw.NewCubicBezierCurve(450, 400, 421, 420, 400, 380, 400, 340), // bottom left
		draw.NewCubicBezierCurve(400, 340, 395, 310, 422, 284, 450, 300), // top left
		draw.NewCubicBezierCurve(450, 300, 446, 288, 442, 260, 450, 250), // leaf
	})

	curve.SetBorderColor(creator.ColorBlue)
	curve.SetFillColor(creator.ColorGreen)
	curve.SetBorderWidth(2)

	return c.Draw(curve)
}

// Draw using curve polygon.
func drawCurvePolygon(c *creator.Creator) error {
	drawHeader(c, 30, 500, "-- Curve Polygon --")

	curvePolygon := c.NewCurvePolygon([][]draw.CubicBezierCurve{
		{
			draw.NewCubicBezierCurve(10, 750, 10, 750, 10, 650, 10, 650),
			draw.NewCubicBezierCurve(10, 650, 20, 550, 200, 550, 210, 650),
			draw.NewCubicBezierCurve(210, 650, 210, 650, 210, 750, 210, 750),
			draw.NewCubicBezierCurve(210, 750, 10, 750, 10, 750, 10, 750),
		},
		{
			draw.NewCubicBezierCurve(60, 675, 60, 675, 85, 675, 85, 675),
			draw.NewCubicBezierCurve(85, 675, 85, 675, 85, 650, 85, 650),
			draw.NewCubicBezierCurve(85, 650, 85, 650, 60, 650, 60, 650),
			draw.NewCubicBezierCurve(60, 650, 60, 650, 60, 675, 60, 675),
		},
		{
			draw.NewCubicBezierCurve(110, 675, 110, 675, 135, 675, 135, 675),
			draw.NewCubicBezierCurve(135, 675, 125, 630, 115, 630, 110, 675),
		},
	})
	curvePolygon.SetBorderColor(creator.ColorRGBFromHex("#00FF00"))
	curvePolygon.SetBorderWidth(3)
	curvePolygon.SetBorderOpacity(1)
	curvePolygon.SetFillColor(creator.ColorRGBFromHex("#0000FF"))
	curvePolygon.SetFillOpacity(0.5)

	return c.Draw(curvePolygon)
}

// Draw simple polygon.
func drawPolygon(c *creator.Creator) error {
	drawHeader(c, 100, 850, "-- Polygon -- ")

	// Draw polygon.
	polygon := c.NewPolygon([][]draw.Point{{
		{X: 50, Y: 950},
		{X: 100, Y: 900},
		{X: 200.0, Y: 900.0},
		{X: 250.0, Y: 950.0},
		{X: 250.0, Y: 1100.0},
		{X: 200.0, Y: 1150.0},
		{X: 100.0, Y: 1150.0},
		{X: 50.0, Y: 1100.0},
		{X: 50.0, Y: 950.0},
	}})
	polygon.SetFillColor(creator.ColorYellow)
	polygon.SetBorderColor(creator.ColorRed)
	polygon.SetBorderWidth(3)
	polygon.SetFillOpacity(0.5)
	polygon.SetBorderOpacity(0.9)

	return c.Draw(polygon)
}

package ui

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/johanhenriksson/goworld/assets"
	"github.com/johanhenriksson/goworld/geometry"
	"github.com/johanhenriksson/goworld/render"
)

type Rect struct {
	*Element
	Color render.Color
	quad  *geometry.Quad
}

func (m *Manager) NewRect(color render.Color, x, y, w, h, z float32) *Rect {
	// UI Manager should provide access to some resource manager thingy
	// mat := m.Resources.GetMaterial("assets/materials/ui_color.json")
	mat := assets.GetMaterial("ui_color")

	el := m.NewElement("Rect", x, y, w, h, z)
	r := &Rect{
		Element: el,
		Color:   color,
		quad:    geometry.NewQuad(mat, color, w, h, z),
	}
	//r.quad.SetBorderWidth(5)
	return r
}

func (r *Rect) Draw(args render.DrawArgs) {
	// this is sort of ugly. we dont really want to duplicate the transform
	// multiplication to every element. on the other hand, most elements
	// will need to apply the transform before they draw themselves

	/* compute local transform */
	local := args
	local.Transform = r.Element.Transform.Matrix.Mul4(args.Transform)

	/* draw rect */
	// TODO set color
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	r.quad.Draw(local)

	/* call parent - draw children etc */
	r.Element.Draw(args)
}

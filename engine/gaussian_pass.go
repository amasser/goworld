package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/johanhenriksson/goworld/render"
)

// GaussianPass represents a gaussian blur pass.
type GaussianPass struct {
	fbo      *render.FrameBuffer
	material *render.Material
	quad     *render.Quad
	Output   *render.Texture
}

// NewGaussianPass creates a new Gaussian Blur pass.
func NewGaussianPass(input *render.Texture) *GaussianPass {
	fbo := render.CreateFrameBuffer(input.Width, input.Height)
	fbo.ClearColor = render.Color4(1, 1, 1, 1)
	texture := fbo.AttachBuffer(gl.COLOR_ATTACHMENT0, gl.RED, gl.RGB, gl.FLOAT)

	mat := render.CreateMaterial(render.CompileVFShader("/assets/shaders/gaussian"))
	mat.AddDescriptors(render.F32_XYZUV)
	mat.AddTexture("tex_input", input)

	quad := render.NewQuad(mat)

	return &GaussianPass{
		fbo:      fbo,
		material: mat,
		quad:     quad,
		Output:   texture,
	}
}

// DrawPass draws the gaussian blurred output to the frame buffer.
func (p *GaussianPass) DrawPass(scene *Scene) {
	p.fbo.Bind()
	p.fbo.Clear()

	p.quad.Draw()

	p.fbo.Unbind()
}

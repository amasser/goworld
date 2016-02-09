package engine

import (
    "github.com/johanhenriksson/goworld/render"
    "github.com/johanhenriksson/goworld/assets"

    "github.com/go-gl/gl/v4.1-core/gl"
    mgl "github.com/go-gl/mathgl/mgl32"
)

type ShadowPass struct {
    Output    *render.Texture
    Material  *render.Material
    Width     int
    Height    int
    shadowmap *render.FrameBuffer
}

func NewShadowPass(input *render.GeometryBuffer) *ShadowPass {
    mat := assets.GetMaterial("color_geometry")

    shadow_map := render.CreateFrameBuffer(500,500)
    shadow_map.ClearColor = render.Color4(1,1,1,1)
    shadow_texture := shadow_map.AddBuffer(gl.DEPTH_ATTACHMENT, gl.DEPTH_COMPONENT24, gl.DEPTH_COMPONENT, gl.FLOAT)

    p := &ShadowPass {
        Material: mat,
        shadowmap: shadow_map,
        Output: shadow_texture,
        Width: 300,
        Height: 300,
    }
    return p
}

func (sp *ShadowPass) DrawPass(scene *Scene, light *Light) {

    if light.Type != DirectionalLight {
        return
    }

    sp.shadowmap.Bind()
    sp.shadowmap.Clear()

    /* use light pass shader */
    sp.Material.Use()
    //shader := sp.Material.Shader

    /* clear */


    p := mgl.Ortho(float32(-sp.Width / 2), float32(sp.Width / 2), float32(-sp.Height / 2), float32(sp.Height / 2), -150, 150)
    v := mgl.LookAtV(mgl.Vec3{0,0,0}, light.Position.Normalize(), mgl.Vec3{0,1,0})
    vp := p.Mul4(v)

    args := render.DrawArgs {
        Projection: p,
        View: v,
        VP: vp,
        MVP: vp,
        Transform: mgl.Ident4(),

        Pass: "geometry",
        Shader: sp.Material.Shader,
    }
    scene.DrawCall(args)
    //scene.Draw("geometry", sp.Material.Shader)

    sp.shadowmap.Unbind()
    gl.Viewport(0, 0, int32(scene.Camera.Width), int32(scene.Camera.Height))
}
package ui;

import (
    "github.com/johanhenriksson/goworld/engine"
    "github.com/johanhenriksson/goworld/render"
    mgl "github.com/go-gl/mathgl/mgl32"
)

/** Main UI manager. Handles routing of events and drawing the UI. */
type Manager struct {
    Window      *engine.Window
    Viewport    mgl.Mat4

    Children    []render.Drawable

    /** Projection matrix - orthographic */
    /* Shaders */
}

func NewManager(wnd *engine.Window) *Manager {
    m := &Manager {
        Window: wnd,
        Viewport: mgl.Ortho(0, float32(wnd.Width), 0, float32(wnd.Height), 1000, -1000),
        Children: []render.Drawable{},
    }
    return m
}

func (m *Manager) Append(child render.Drawable) {
    m.Children = append(m.Children, child)
}

func (m *Manager) Draw() {
    args := render.DrawArgs {
        Viewport: m.Viewport,
        Transform: mgl.Ident4(),
    }
    for _, el := range m.Children {
        el.Draw(args)
    }
}


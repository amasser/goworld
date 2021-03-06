package render

import (
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

/* Rename to something that reflects that it's an attachment to a frame buffer */
type DrawBuffer struct {
	Target  uint32 // GL attachment enum (DEPTH_ATTACHMENT, COLOR_ATTACHMENT etc)
	Texture *Texture
}

/** Represents an OpenGL frame buffer object */
type FrameBuffer struct {
	Buffers    []DrawBuffer
	ClearColor Color
	Width      int32
	Height     int32
	id         uint32
	mipLvl     int32
}

var ScreenBuffer = FrameBuffer{
	Buffers:    []DrawBuffer{},
	ClearColor: Color{0, 0, 0, 1},
	Width:      0,
	Height:     0,
	id:         0,
	mipLvl:     0,
}

/**
 * Create a new frame buffer texture and attach it to the given target.
 * Returns a pointer to the created texture object. FBO must be bound first.
 */
func (f *FrameBuffer) AttachBuffer(target, internalFormat, format, datatype uint32) *Texture {
	// Create texture object
	texture := CreateTexture(f.Width, f.Height)
	texture.Format = format
	texture.InternalFormat = internalFormat
	texture.DataType = datatype
	texture.Clear()

	// Set texture as frame buffer target
	texture.FrameBufferTarget(target)

	if target != gl.DEPTH_ATTACHMENT {
		// Attach to frame buffer
		f.Buffers = append(f.Buffers, DrawBuffer{
			Target:  target,
			Texture: texture,
		})
	}

	return texture
}

func CreateFrameBuffer(width, height int32) *FrameBuffer {
	f := &FrameBuffer{
		Width:      width,
		Height:     height,
		Buffers:    []DrawBuffer{},
		ClearColor: Color4(0, 0, 0, 1),
	}
	gl.GenFramebuffers(1, &f.id)
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.id)
	return f
}

func (f *FrameBuffer) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, 0) // why?

	// bind this frame buffer
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.id)

	// set viewport size equal to buffer size
	gl.Viewport(0, 0, f.Width, f.Height)
}

func (f *FrameBuffer) Unbind() {
	// finish drawing
	gl.Flush()

	// unbind
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

/* Clear the frame buffer. Make sure its bound first */
func (f *FrameBuffer) Clear() {
	gl.ClearColor(f.ClearColor.R, f.ClearColor.G, f.ClearColor.B, f.ClearColor.A)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

/* Delete frame buffer object */
func (f *FrameBuffer) Delete() {
	if f.id == 0 {
		panic("Cant delete framebuffer 0")
	}
	gl.DeleteFramebuffers(1, &f.id)
	f.id = 0
}

func (f *FrameBuffer) Sample(target uint32, x, y int) Color {
	pixel := make([]float32, 4)
	gl.ReadBuffer(target)
	gl.ReadPixels(int32(x), int32(y), 1, 1, gl.RGBA, gl.FLOAT, unsafe.Pointer(&pixel[0]))
	return Color4(pixel[0], pixel[1], pixel[2], pixel[3])
}

func (f *FrameBuffer) SampleDepth(x, y int) float32 {
	float := float32(0)
	gl.ReadPixels(int32(x), int32(y), 1, 1, gl.DEPTH_COMPONENT, gl.FLOAT, unsafe.Pointer(&float))
	return float
}

// DrawBuffers sets up all the attached buffers for drawing
func (f *FrameBuffer) DrawBuffers() {
	buff := []uint32{}
	for _, buffer := range f.Buffers {
		buff = append(buff, buffer.Target)
	}
	gl.DrawBuffers(int32(len(buff)), &buff[0])
}

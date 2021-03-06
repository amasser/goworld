package render

import (
	"image"
	"image/draw"
	_ "image/png" // png support
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/johanhenriksson/goworld/util"
)

type TextureFilter int32
type WrapMode int32

const (
	LinearFilter  = TextureFilter(gl.LINEAR)
	NearestFilter = TextureFilter(gl.NEAREST)
)

const (
	ClampWrap = WrapMode(gl.CLAMP_TO_EDGE)
)

// Texture represents an OpenGL 2D texture object
type Texture struct {
	ID             uint32
	Width          int32
	Height         int32
	Format         uint32
	InternalFormat uint32
	DataType       uint32
	MipLevel       int32

	filter TextureFilter
	wrap   WrapMode
}

// CreateTexture creates a new 2D texture and sets some sane defaults
func CreateTexture(width, height int32) *Texture {
	var id uint32
	gl.GenTextures(1, &id)

	tx := &Texture{
		ID:             id,
		Width:          width,
		Height:         height,
		Format:         gl.RGBA,
		InternalFormat: gl.RGBA,
		DataType:       gl.UNSIGNED_BYTE,
	}
	tx.Bind()

	/* Texture parameters - pass as parameters? */
	tx.SetFilter(LinearFilter)
	tx.SetWrapMode(ClampWrap)

	return tx
}

func (tx *Texture) SetFilter(filter TextureFilter) {
	tx.filter = filter
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int32(filter))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int32(filter))
}

func (tx *Texture) SetWrapMode(mode WrapMode) {
	tx.wrap = mode
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, int32(mode))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, int32(mode))
}

// Use binds this texture to the given texture slot */
func (tx *Texture) Use(slot uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + slot)
	gl.Enable(gl.TEXTURE_2D)
	tx.Bind()
}

// Bind texture to the currently active texture slot
func (tx *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, tx.ID)
}

// FrameBufferTarget attaches this texture to the current frame buffer object
func (tx *Texture) FrameBufferTarget(attachment uint32) {
	gl.FramebufferTexture(gl.FRAMEBUFFER, attachment, tx.ID, tx.MipLevel)
}

// Clear the texture
func (tx *Texture) Clear() {
	tx.Bind()
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		int32(tx.InternalFormat), // gl.RGBA,
		tx.Width, tx.Height,
		0,
		tx.Format,   //gl.RGBA,
		tx.DataType, // gl.UNSIGNED_BYTE,
		nil)         // null ptr
}

// Buffer buffers texture data from an image object
func (tx *Texture) Buffer(img *image.RGBA) {
	/* Buffer image data */
	tx.Bind()
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		int32(tx.InternalFormat),
		tx.Width, tx.Height,
		0,
		tx.Format, tx.DataType,
		gl.Ptr(img.Pix))
}

// BufferFloats buffers texture data from a float array
func (tx *Texture) BufferFloats(img []float32) {
	/* Buffer image data */
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		int32(tx.InternalFormat),
		tx.Width, tx.Height,
		0,
		tx.Format, tx.DataType,
		gl.Ptr(&img[0]))
}

// TextureFromImage is a helper method to create an OpenGL texture from an image object */
func TextureFromImage(img *image.RGBA) *Texture {
	width := int32(img.Rect.Size().X)
	height := int32(img.Rect.Size().Y)
	tx := CreateTexture(width, height)
	tx.Buffer(img)
	return tx
}

// TextureFromFile loads a texture from file */
func TextureFromFile(file string) (*Texture, error) {
	img, err := ImageFromFile(file)
	if err != nil {
		return nil, err
	}
	return TextureFromImage(img), nil
}

// ImageFromFile loads an image from a file. Returns an RGBA image object */
func ImageFromFile(file string) (*image.RGBA, error) {
	// todo: http support?

	imgFile, err := os.Open(util.ExePath + file)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	return rgba, nil
}

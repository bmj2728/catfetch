package catpic

import (
	"image"
	"sync"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type CatPic struct {
	img       image.Image
	mu        sync.Mutex
	isLoading bool
}

func NewCatImage(img image.Image) *CatPic {
	return &CatPic{
		img: img,
	}
}

func (p *CatPic) IsLoading() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.isLoading
}

func (p *CatPic) GetImage() image.Image {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.img
}

func (p *CatPic) SetImage(img image.Image) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.img = img
}

func (p *CatPic) SetLoading() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.isLoading = true
}

func (p *CatPic) ClearLoading() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.isLoading = false
}

func (p *CatPic) Draw(gtx layout.Context) layout.Dimensions {
	img := p.GetImage()
	if img == nil {
		// Could render a placeholder or loading state here
		return layout.Dimensions{Size: gtx.Constraints.Min}
	}

	bounds := img.Bounds()
	imgW, imgH := float32(bounds.Dx()), float32(bounds.Dy())
	maxW, maxH := float32(gtx.Constraints.Max.X), float32(gtx.Constraints.Max.Y)

	scale := min(maxW/imgW, maxH/imgH)
	finalW, finalH := int(imgW*scale), int(imgH*scale)

	// Clip to the scaled bounds
	defer clip.Rect{Max: image.Pt(finalW, finalH)}.Push(gtx.Ops).Pop()

	imgOp := paint.NewImageOp(img)
	imgOp.Filter = paint.FilterLinear
	imgOp.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return layout.Dimensions{Size: image.Pt(finalW, finalH)}
}

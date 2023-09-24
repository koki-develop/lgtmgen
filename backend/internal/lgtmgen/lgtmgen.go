package lgtmgen

import (
	"math"

	"github.com/cockroachdb/errors"
	"gopkg.in/gographics/imagick.v3/imagick"
)

const (
	maxSideLength float64 = 425
	font          string  = "assets/fonts/Archivo_Black/ArchivoBlack-Regular.ttf"
)

func Generate(src []byte) ([]byte, error) {
	imagick.Initialize()
	defer imagick.Terminate()

	srcmw := imagick.NewMagickWand()
	defer srcmw.Destroy()
	if err := srcmw.ReadImageBlob(src); err != nil {
		return nil, errors.Wrap(err, "failed to read image")
	}
	w := srcmw.GetImageWidth()
	h := srcmw.GetImageHeight()
	dw, dh := calcImageSize(float64(w), float64(h))
	ttlfs, txtfs := calcFontSize(dw, dh)

	ttldw := imagick.NewDrawingWand()
	defer ttldw.Destroy()
	txtdw := imagick.NewDrawingWand()
	defer txtdw.Destroy()

	if err := ttldw.SetFont(font); err != nil {
		return nil, errors.Wrap(err, "failed to set font to title")
	}
	if err := txtdw.SetFont(font); err != nil {
		return nil, errors.Wrap(err, "failed to set font to text")
	}

	fgpw := imagick.NewPixelWand()
	if ok := fgpw.SetColor("#ffffff"); !ok {
		return nil, errors.New("invalid color")
	}
	bgpw := imagick.NewPixelWand()
	if ok := bgpw.SetColor("#000000"); !ok {
		return nil, errors.New("invalid color")
	}

	ttldw.SetStrokeColor(bgpw)
	txtdw.SetStrokeColor(bgpw)
	ttldw.SetStrokeWidth(1)
	txtdw.SetStrokeWidth(0.8)
	ttldw.SetFillColor(fgpw)
	txtdw.SetFillColor(fgpw)
	ttldw.SetFontSize(ttlfs)
	txtdw.SetFontSize(txtfs)
	ttldw.SetGravity(imagick.GRAVITY_CENTER)
	txtdw.SetGravity(imagick.GRAVITY_CENTER)
	ttldw.Annotation(0, 0, "L G T M")
	txtdw.Annotation(0, ttlfs/1.5, "L o o k s   G o o d   T o   M e")

	cimw := srcmw.CoalesceImages()
	defer cimw.Destroy()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	_ = mw.SetImageDelay(srcmw.GetImageDelay())

	for i := 0; i < int(cimw.GetNumberImages()); i++ {
		if ok := cimw.SetIteratorIndex(i); !ok {
			return nil, errors.New("invalid index")
		}

		img := cimw.GetImage()
		defer img.Destroy()

		if err := img.AdaptiveResizeImage(uint(dw), uint(dh)); err != nil {
			return nil, errors.Wrap(err, "failed to resize image")
		}
		if err := img.DrawImage(ttldw); err != nil {
			return nil, errors.Wrap(err, "failed to draw title")
		}
		if err := img.DrawImage(txtdw); err != nil {
			return nil, errors.Wrap(err, "failed to draw text")
		}
		if err := mw.AddImage(img); err != nil {
			return nil, errors.Wrap(err, "failed to add image")
		}
	}

	return mw.GetImagesBlob(), nil
}

func calcImageSize(w, h float64) (float64, float64) {
	if w > h {
		return maxSideLength, maxSideLength * h / w
	} else {
		return maxSideLength * w / h, maxSideLength
	}
}

func calcFontSize(w, h float64) (float64, float64) {
	return math.Min(h/2, w/6), math.Min(h/9, w/27)
}

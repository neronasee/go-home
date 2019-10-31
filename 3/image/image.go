package image

const DefaultWidth = 38
const DefaultHeight = 18

type Point struct{}
type ColoredPoint struct {
	point Point
	color uint8
}
type Image struct {
	width  int
	height int
	points []ColoredPoint
}

func (rcv *Image) Load(rawData []uint8) {
	newPoints := make([]ColoredPoint, len(rawData))

	for i, color := range rawData {
		newPoints[i] = ColoredPoint{
			Point{},
			color,
		}
	}

	rcv.points = newPoints
}

func NewImage() *Image {
	return &Image{
		width:  DefaultWidth,
		height: DefaultHeight,
		points: make([]ColoredPoint, DefaultWidth*DefaultHeight),
	}
}

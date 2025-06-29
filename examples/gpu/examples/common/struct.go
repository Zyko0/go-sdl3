package common

type PositionColorVertex struct {
	X, Y, Z    float32
	R, G, B, A uint8
}

func NewPosColorVert(x, y, z float32, r, g, b, a uint8) PositionColorVertex {
	return PositionColorVertex{
		X: x, Y: y, Z: z,
		R: r, G: g, B: b, A: a,
	}
}

type PositionTextureVertex struct {
	X, Y, Z float32
	U, V    float32
}

func NewPosTexVert(x, y, z float32, u, v float32) PositionTextureVertex {
	return PositionTextureVertex{
		X: x, Y: y, Z: z,
		U: u, V: v,
	}
}

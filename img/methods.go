package img

// Animation

// IMG_FreeAnimation - Dispose of an IMG_Animation and free its resources.
// (https://wiki.libsdl.org/SDL3_img/IMG_FreeAnimation)
func (anim *Animation) Free() {
	iFreeAnimation(anim)
}

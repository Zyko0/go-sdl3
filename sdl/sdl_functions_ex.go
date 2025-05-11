package sdl

var (
	//puregogen:library path:windows=sdl.dll path:unix=sdl.so alias=sdl
	//puregogen:function symbol=SDL_ShowMessageBox
	iShowMessageBox func(data *messageBoxData, buttonid *int32) bool
)

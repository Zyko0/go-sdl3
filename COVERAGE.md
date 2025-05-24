# API Coverage


This file tracks the functions that have been wrapped.<br>
The following emojis mean (they are clickable and should link to the code implementation):
- :heavy_check_mark: = implemented
- :x: = not implemented yet
- :question: = not planned / don't know about integrating it or not

## SDL

### Init

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_Init](https://wiki.libsdl.org/SDL3/SDL_Init) | [:heavy_check_mark:](sdl/functions.go#L13) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13899) |
| [SDL_InitSubSystem](https://wiki.libsdl.org/SDL3/SDL_InitSubSystem) | [:heavy_check_mark:](sdl/functions.go#L23) | [:x:](sdl/sdl_functions_js.go#L13909) |
| [SDL_QuitSubSystem](https://wiki.libsdl.org/SDL3/SDL_QuitSubSystem) | [:heavy_check_mark:](sdl/functions.go#L33) | [:x:](sdl/sdl_functions_js.go#L13922) |
| [SDL_WasInit](https://wiki.libsdl.org/SDL3/SDL_WasInit) | [:heavy_check_mark:](sdl/functions.go#L39) | [:x:](sdl/sdl_functions_js.go#L13933) |
| [SDL_Quit](https://wiki.libsdl.org/SDL3/SDL_Quit) | [:heavy_check_mark:](sdl/functions.go#L45) | [:heavy_check_mark:]() |
| [SDL_IsMainThread](https://wiki.libsdl.org/SDL3/SDL_IsMainThread) | [:heavy_check_mark:](sdl/functions.go#L51) | [:x:](sdl/sdl_functions_js.go#L13950) |
| [SDL_RunOnMainThread](https://wiki.libsdl.org/SDL3/SDL_RunOnMainThread) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13961) |
| [SDL_SetAppMetadata](https://wiki.libsdl.org/SDL3/SDL_SetAppMetadata) | [:heavy_check_mark:](sdl/functions.go#L57) | [:x:](sdl/sdl_functions_js.go#L13978) |
| [SDL_SetAppMetadataProperty](https://wiki.libsdl.org/SDL3/SDL_SetAppMetadataProperty) | [:heavy_check_mark:](sdl/functions.go#L67) | [:x:](sdl/sdl_functions_js.go#L13995) |
| [SDL_GetAppMetadataProperty](https://wiki.libsdl.org/SDL3/SDL_GetAppMetadataProperty) | [:heavy_check_mark:](sdl/functions.go#L77) | [:x:](sdl/sdl_functions_js.go#L14010) |
### Hints

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetHintWithPriority](https://wiki.libsdl.org/SDL3/SDL_SetHintWithPriority) | [:heavy_check_mark:](sdl/functions.go#L85) | [:x:](sdl/sdl_functions_js.go#L13785) |
| [SDL_SetHint](https://wiki.libsdl.org/SDL3/SDL_SetHint) | [:heavy_check_mark:](sdl/functions.go#L95) | [:x:](sdl/sdl_functions_js.go#L13802) |
| [SDL_ResetHint](https://wiki.libsdl.org/SDL3/SDL_ResetHint) | [:heavy_check_mark:](sdl/functions.go#L105) | [:x:](sdl/sdl_functions_js.go#L13817) |
| [SDL_ResetHints](https://wiki.libsdl.org/SDL3/SDL_ResetHints) | [:heavy_check_mark:](sdl/functions.go#L115) | [:x:](sdl/sdl_functions_js.go#L13830) |
| [SDL_GetHint](https://wiki.libsdl.org/SDL3/SDL_GetHint) | [:heavy_check_mark:](sdl/functions.go#L121) | [:x:](sdl/sdl_functions_js.go#L13839) |
| [SDL_GetHintBoolean](https://wiki.libsdl.org/SDL3/SDL_GetHintBoolean) | [:heavy_check_mark:](sdl/functions.go#L127) | [:x:](sdl/sdl_functions_js.go#L13852) |
| [SDL_AddHintCallback](https://wiki.libsdl.org/SDL3/SDL_AddHintCallback) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13867) |
| [SDL_RemoveHintCallback](https://wiki.libsdl.org/SDL3/SDL_RemoveHintCallback) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13884) |
### Error

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetError](https://wiki.libsdl.org/SDL3/SDL_SetError) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L488) |
| [SDL_SetErrorV](https://wiki.libsdl.org/SDL3/SDL_SetErrorV) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L501) |
| [SDL_OutOfMemory](https://wiki.libsdl.org/SDL3/SDL_OutOfMemory) | [:heavy_check_mark:](sdl/functions.go#L138) | [:x:](sdl/sdl_functions_js.go#L516) |
| [SDL_GetError](https://wiki.libsdl.org/SDL3/SDL_GetError) | [:heavy_check_mark:](sdl/init_notjs.go#L32) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L527) |
| [SDL_ClearError](https://wiki.libsdl.org/SDL3/SDL_ClearError) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L535) |
### Properties

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetGlobalProperties](https://wiki.libsdl.org/SDL3/SDL_GetGlobalProperties) | [:heavy_check_mark:](sdl/functions.go#L146) | [:x:](sdl/sdl_functions_js.go#L546) |
| [SDL_CreateProperties](https://wiki.libsdl.org/SDL3/SDL_CreateProperties) | [:heavy_check_mark:](sdl/functions.go#L157) | [:x:](sdl/sdl_functions_js.go#L557) |
| [SDL_CopyProperties](https://wiki.libsdl.org/SDL3/SDL_CopyProperties) | [:heavy_check_mark:](sdl/functions.go#L168) | [:x:](sdl/sdl_functions_js.go#L568) |
| [SDL_LockProperties](https://wiki.libsdl.org/SDL3/SDL_LockProperties) | [:heavy_check_mark:](sdl/methods.go#L5689) | [:x:](sdl/sdl_functions_js.go#L583) |
| [SDL_UnlockProperties](https://wiki.libsdl.org/SDL3/SDL_UnlockProperties) | [:heavy_check_mark:](sdl/methods.go#L5699) | [:x:](sdl/sdl_functions_js.go#L596) |
| [SDL_SetPointerPropertyWithCleanup](https://wiki.libsdl.org/SDL3/SDL_SetPointerPropertyWithCleanup) | [:x:](sdl/methods.go#L5705) | [:x:](sdl/sdl_functions_js.go#L607) |
| [SDL_SetPointerProperty](https://wiki.libsdl.org/SDL3/SDL_SetPointerProperty) | [:x:](sdl/methods.go#L5712) | [:x:](sdl/sdl_functions_js.go#L628) |
| [SDL_SetStringProperty](https://wiki.libsdl.org/SDL3/SDL_SetStringProperty) | [:heavy_check_mark:](sdl/methods.go#L5719) | [:x:](sdl/sdl_functions_js.go#L645) |
| [SDL_SetNumberProperty](https://wiki.libsdl.org/SDL3/SDL_SetNumberProperty) | [:heavy_check_mark:](sdl/methods.go#L5729) | [:x:](sdl/sdl_functions_js.go#L662) |
| [SDL_SetFloatProperty](https://wiki.libsdl.org/SDL3/SDL_SetFloatProperty) | [:heavy_check_mark:](sdl/methods.go#L5739) | [:x:](sdl/sdl_functions_js.go#L679) |
| [SDL_SetBooleanProperty](https://wiki.libsdl.org/SDL3/SDL_SetBooleanProperty) | [:heavy_check_mark:](sdl/methods.go#L5749) | [:x:](sdl/sdl_functions_js.go#L696) |
| [SDL_HasProperty](https://wiki.libsdl.org/SDL3/SDL_HasProperty) | [:heavy_check_mark:](sdl/methods.go#L5759) | [:x:](sdl/sdl_functions_js.go#L713) |
| [SDL_GetPropertyType](https://wiki.libsdl.org/SDL3/SDL_GetPropertyType) | [:heavy_check_mark:](sdl/methods.go#L5765) | [:x:](sdl/sdl_functions_js.go#L728) |
| [SDL_GetPointerProperty](https://wiki.libsdl.org/SDL3/SDL_GetPointerProperty) | [:x:](sdl/methods.go#L5771) | [:x:](sdl/sdl_functions_js.go#L743) |
| [SDL_GetStringProperty](https://wiki.libsdl.org/SDL3/SDL_GetStringProperty) | [:heavy_check_mark:](sdl/methods.go#L5778) | [:x:](sdl/sdl_functions_js.go#L760) |
| [SDL_GetNumberProperty](https://wiki.libsdl.org/SDL3/SDL_GetNumberProperty) | [:heavy_check_mark:](sdl/methods.go#L5784) | [:x:](sdl/sdl_functions_js.go#L777) |
| [SDL_GetFloatProperty](https://wiki.libsdl.org/SDL3/SDL_GetFloatProperty) | [:heavy_check_mark:](sdl/methods.go#L5790) | [:x:](sdl/sdl_functions_js.go#L794) |
| [SDL_GetBooleanProperty](https://wiki.libsdl.org/SDL3/SDL_GetBooleanProperty) | [:heavy_check_mark:](sdl/methods.go#L5796) | [:x:](sdl/sdl_functions_js.go#L811) |
| [SDL_ClearProperty](https://wiki.libsdl.org/SDL3/SDL_ClearProperty) | [:heavy_check_mark:](sdl/methods.go#L5802) | [:x:](sdl/sdl_functions_js.go#L828) |
| [SDL_EnumerateProperties](https://wiki.libsdl.org/SDL3/SDL_EnumerateProperties) | [:x:](sdl/methods.go#L5812) | [:x:](sdl/sdl_functions_js.go#L843) |
| [SDL_DestroyProperties](https://wiki.libsdl.org/SDL3/SDL_DestroyProperties) | [:heavy_check_mark:](sdl/methods.go#L5819) | [:x:](sdl/sdl_functions_js.go#L860) |
### Log

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetLogPriorities](https://wiki.libsdl.org/SDL3/SDL_SetLogPriorities) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14087) |
| [SDL_SetLogPriority](https://wiki.libsdl.org/SDL3/SDL_SetLogPriority) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14098) |
| [SDL_GetLogPriority](https://wiki.libsdl.org/SDL3/SDL_GetLogPriority) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14111) |
| [SDL_ResetLogPriorities](https://wiki.libsdl.org/SDL3/SDL_ResetLogPriorities) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14124) |
| [SDL_SetLogPriorityPrefix](https://wiki.libsdl.org/SDL3/SDL_SetLogPriorityPrefix) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14133) |
| [SDL_Log](https://wiki.libsdl.org/SDL3/SDL_Log) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14148) |
| [SDL_LogTrace](https://wiki.libsdl.org/SDL3/SDL_LogTrace) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14159) |
| [SDL_LogVerbose](https://wiki.libsdl.org/SDL3/SDL_LogVerbose) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14172) |
| [SDL_LogDebug](https://wiki.libsdl.org/SDL3/SDL_LogDebug) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14185) |
| [SDL_LogInfo](https://wiki.libsdl.org/SDL3/SDL_LogInfo) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14198) |
| [SDL_LogWarn](https://wiki.libsdl.org/SDL3/SDL_LogWarn) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14211) |
| [SDL_LogError](https://wiki.libsdl.org/SDL3/SDL_LogError) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14224) |
| [SDL_LogCritical](https://wiki.libsdl.org/SDL3/SDL_LogCritical) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14237) |
| [SDL_LogMessage](https://wiki.libsdl.org/SDL3/SDL_LogMessage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14250) |
| [SDL_LogMessageV](https://wiki.libsdl.org/SDL3/SDL_LogMessageV) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14265) |
| [SDL_GetDefaultLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_GetDefaultLogOutputFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14282) |
| [SDL_GetLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_GetLogOutputFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14293) |
| [SDL_SetLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_SetLogOutputFunction) | [:x:](sdl/methods.go#L3998) | [:x:](sdl/sdl_functions_js.go#L14312) |
### Video

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumVideoDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumVideoDrivers) | [:heavy_check_mark:](sdl/functions.go#L490) | [:x:](sdl/sdl_functions_js.go#L5828) |
| [SDL_GetVideoDriver](https://wiki.libsdl.org/SDL3/SDL_GetVideoDriver) | [:heavy_check_mark:](sdl/functions.go#L496) | [:x:](sdl/sdl_functions_js.go#L5839) |
| [SDL_GetCurrentVideoDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentVideoDriver) | [:heavy_check_mark:](sdl/functions.go#L502) | [:x:](sdl/sdl_functions_js.go#L5852) |
| [SDL_GetSystemTheme](https://wiki.libsdl.org/SDL3/SDL_GetSystemTheme) | [:heavy_check_mark:](sdl/functions.go#L508) | [:x:](sdl/sdl_functions_js.go#L5863) |
| [SDL_GetDisplays](https://wiki.libsdl.org/SDL3/SDL_GetDisplays) | [:heavy_check_mark:](sdl/functions.go#L514) | [:x:](sdl/sdl_functions_js.go#L5874) |
| [SDL_GetPrimaryDisplay](https://wiki.libsdl.org/SDL3/SDL_GetPrimaryDisplay) | [:heavy_check_mark:](sdl/functions.go#L528) | [:x:](sdl/sdl_functions_js.go#L5890) |
| [SDL_GetDisplayProperties](https://wiki.libsdl.org/SDL3/SDL_GetDisplayProperties) | [:heavy_check_mark:](sdl/methods.go#L3550) | [:x:](sdl/sdl_functions_js.go#L5901) |
| [SDL_GetDisplayName](https://wiki.libsdl.org/SDL3/SDL_GetDisplayName) | [:heavy_check_mark:](sdl/methods.go#L3561) | [:x:](sdl/sdl_functions_js.go#L5914) |
| [SDL_GetDisplayBounds](https://wiki.libsdl.org/SDL3/SDL_GetDisplayBounds) | [:heavy_check_mark:](sdl/methods.go#L3572) | [:x:](sdl/sdl_functions_js.go#L5927) |
| [SDL_GetDisplayUsableBounds](https://wiki.libsdl.org/SDL3/SDL_GetDisplayUsableBounds) | [:heavy_check_mark:](sdl/methods.go#L3584) | [:x:](sdl/sdl_functions_js.go#L5945) |
| [SDL_GetNaturalDisplayOrientation](https://wiki.libsdl.org/SDL3/SDL_GetNaturalDisplayOrientation) | [:heavy_check_mark:](sdl/methods.go#L3596) | [:x:](sdl/sdl_functions_js.go#L5963) |
| [SDL_GetCurrentDisplayOrientation](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayOrientation) | [:heavy_check_mark:](sdl/methods.go#L3602) | [:x:](sdl/sdl_functions_js.go#L5976) |
| [SDL_GetDisplayContentScale](https://wiki.libsdl.org/SDL3/SDL_GetDisplayContentScale) | [:heavy_check_mark:](sdl/methods.go#L3608) | [:x:](sdl/sdl_functions_js.go#L5989) |
| [SDL_GetFullscreenDisplayModes](https://wiki.libsdl.org/SDL3/SDL_GetFullscreenDisplayModes) | [:heavy_check_mark:](sdl/methods.go#L3619) | [:x:](sdl/sdl_functions_js.go#L6002) |
| [SDL_GetClosestFullscreenDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetClosestFullscreenDisplayMode) | [:heavy_check_mark:](sdl/methods.go#L3633) | [:x:](sdl/sdl_functions_js.go#L6020) |
| [SDL_GetDesktopDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetDesktopDisplayMode) | [:heavy_check_mark:](sdl/methods.go#L3645) | [:x:](sdl/sdl_functions_js.go#L6046) |
| [SDL_GetCurrentDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayMode) | [:heavy_check_mark:](sdl/methods.go#L3656) | [:x:](sdl/sdl_functions_js.go#L6062) |
| [SDL_GetDisplayForPoint](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForPoint) | [:heavy_check_mark:](sdl/functions.go#L534) | [:x:](sdl/sdl_functions_js.go#L6078) |
| [SDL_GetDisplayForRect](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForRect) | [:heavy_check_mark:](sdl/functions.go#L540) | [:x:](sdl/sdl_functions_js.go#L6094) |
| [SDL_GetDisplayForWindow](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForWindow) | [:heavy_check_mark:](sdl/functions.go#L546) | [:x:](sdl/sdl_functions_js.go#L6110) |
| [SDL_GetWindowPixelDensity](https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelDensity) | [:heavy_check_mark:](sdl/methods.go#L4044) | [:x:](sdl/sdl_functions_js.go#L6126) |
| [SDL_GetWindowDisplayScale](https://wiki.libsdl.org/SDL3/SDL_GetWindowDisplayScale) | [:heavy_check_mark:](sdl/methods.go#L4055) | [:x:](sdl/sdl_functions_js.go#L6142) |
| [SDL_SetWindowFullscreenMode](https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreenMode) | [:heavy_check_mark:](sdl/methods.go#L4066) | [:x:](sdl/sdl_functions_js.go#L6158) |
| [SDL_GetWindowFullscreenMode](https://wiki.libsdl.org/SDL3/SDL_GetWindowFullscreenMode) | [:heavy_check_mark:](sdl/methods.go#L4076) | [:x:](sdl/sdl_functions_js.go#L6179) |
| [SDL_GetWindowICCProfile](https://wiki.libsdl.org/SDL3/SDL_GetWindowICCProfile) | [:heavy_check_mark:](sdl/methods.go#L4082) | [:x:](sdl/sdl_functions_js.go#L6198) |
| [SDL_GetWindowPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelFormat) | [:heavy_check_mark:](sdl/methods.go#L4096) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6219) |
| [SDL_GetWindows](https://wiki.libsdl.org/SDL3/SDL_GetWindows) | [:heavy_check_mark:](sdl/functions.go#L552) | [:x:](sdl/sdl_functions_js.go#L6232) |
| [SDL_CreateWindow](https://wiki.libsdl.org/SDL3/SDL_CreateWindow) | [:heavy_check_mark:](sdl/functions.go#L566) | [:x:](sdl/sdl_functions_js.go#L6248) |
| [SDL_CreatePopupWindow](https://wiki.libsdl.org/SDL3/SDL_CreatePopupWindow) | [:heavy_check_mark:](sdl/methods.go#L4107) | [:x:](sdl/sdl_functions_js.go#L6270) |
| [SDL_CreateWindowWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateWindowWithProperties) | [:heavy_check_mark:](sdl/functions.go#L577) | [:x:](sdl/sdl_functions_js.go#L6299) |
| [SDL_GetWindowID](https://wiki.libsdl.org/SDL3/SDL_GetWindowID) | [:heavy_check_mark:](sdl/methods.go#L4118) | [:x:](sdl/sdl_functions_js.go#L6315) |
| [SDL_GetWindowFromID](https://wiki.libsdl.org/SDL3/SDL_GetWindowFromID) | [:heavy_check_mark:](sdl/methods.go#L50) | [:x:](sdl/sdl_functions_js.go#L6331) |
| [SDL_GetWindowParent](https://wiki.libsdl.org/SDL3/SDL_GetWindowParent) | [:heavy_check_mark:](sdl/methods.go#L4129) | [:x:](sdl/sdl_functions_js.go#L6347) |
| [SDL_GetWindowProperties](https://wiki.libsdl.org/SDL3/SDL_GetWindowProperties) | [:heavy_check_mark:](sdl/methods.go#L4135) | [:x:](sdl/sdl_functions_js.go#L6366) |
| [SDL_GetWindowFlags](https://wiki.libsdl.org/SDL3/SDL_GetWindowFlags) | [:heavy_check_mark:](sdl/methods.go#L4146) | [:x:](sdl/sdl_functions_js.go#L6382) |
| [SDL_SetWindowTitle](https://wiki.libsdl.org/SDL3/SDL_SetWindowTitle) | [:heavy_check_mark:](sdl/methods.go#L4152) | [:x:](sdl/sdl_functions_js.go#L6398) |
| [SDL_GetWindowTitle](https://wiki.libsdl.org/SDL3/SDL_GetWindowTitle) | [:heavy_check_mark:](sdl/methods.go#L4162) | [:x:](sdl/sdl_functions_js.go#L6416) |
| [SDL_SetWindowIcon](https://wiki.libsdl.org/SDL3/SDL_SetWindowIcon) | [:heavy_check_mark:](sdl/methods.go#L4168) | [:x:](sdl/sdl_functions_js.go#L6432) |
| [SDL_SetWindowPosition](https://wiki.libsdl.org/SDL3/SDL_SetWindowPosition) | [:heavy_check_mark:](sdl/methods.go#L4178) | [:x:](sdl/sdl_functions_js.go#L6453) |
| [SDL_GetWindowPosition](https://wiki.libsdl.org/SDL3/SDL_GetWindowPosition) | [:heavy_check_mark:](sdl/methods.go#L4188) | [:x:](sdl/sdl_functions_js.go#L6473) |
| [SDL_SetWindowSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowSize) | [:heavy_check_mark:](sdl/methods.go#L4200) | [:x:](sdl/sdl_functions_js.go#L6499) |
| [SDL_GetWindowSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowSize) | [:heavy_check_mark:](sdl/methods.go#L4210) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6519) |
| [SDL_GetWindowSafeArea](https://wiki.libsdl.org/SDL3/SDL_GetWindowSafeArea) | [:heavy_check_mark:](sdl/methods.go#L4222) | [:x:](sdl/sdl_functions_js.go#L6541) |
| [SDL_SetWindowAspectRatio](https://wiki.libsdl.org/SDL3/SDL_SetWindowAspectRatio) | [:heavy_check_mark:](sdl/methods.go#L4234) | [:x:](sdl/sdl_functions_js.go#L6562) |
| [SDL_GetWindowAspectRatio](https://wiki.libsdl.org/SDL3/SDL_GetWindowAspectRatio) | [:heavy_check_mark:](sdl/methods.go#L4244) | [:x:](sdl/sdl_functions_js.go#L6582) |
| [SDL_GetWindowBordersSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowBordersSize) | [:heavy_check_mark:](sdl/methods.go#L4256) | [:x:](sdl/sdl_functions_js.go#L6608) |
| [SDL_GetWindowSizeInPixels](https://wiki.libsdl.org/SDL3/SDL_GetWindowSizeInPixels) | [:heavy_check_mark:](sdl/methods.go#L4268) | [:x:](sdl/sdl_functions_js.go#L6644) |
| [SDL_SetWindowMinimumSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowMinimumSize) | [:heavy_check_mark:](sdl/methods.go#L4280) | [:x:](sdl/sdl_functions_js.go#L6670) |
| [SDL_GetWindowMinimumSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowMinimumSize) | [:heavy_check_mark:](sdl/methods.go#L4290) | [:x:](sdl/sdl_functions_js.go#L6690) |
| [SDL_SetWindowMaximumSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowMaximumSize) | [:heavy_check_mark:](sdl/methods.go#L4302) | [:x:](sdl/sdl_functions_js.go#L6716) |
| [SDL_GetWindowMaximumSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowMaximumSize) | [:heavy_check_mark:](sdl/methods.go#L4312) | [:x:](sdl/sdl_functions_js.go#L6736) |
| [SDL_SetWindowBordered](https://wiki.libsdl.org/SDL3/SDL_SetWindowBordered) | [:heavy_check_mark:](sdl/methods.go#L4324) | [:x:](sdl/sdl_functions_js.go#L6762) |
| [SDL_SetWindowResizable](https://wiki.libsdl.org/SDL3/SDL_SetWindowResizable) | [:heavy_check_mark:](sdl/methods.go#L4334) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6780) |
| [SDL_SetWindowAlwaysOnTop](https://wiki.libsdl.org/SDL3/SDL_SetWindowAlwaysOnTop) | [:heavy_check_mark:](sdl/methods.go#L4344) | [:x:](sdl/sdl_functions_js.go#L6795) |
| [SDL_ShowWindow](https://wiki.libsdl.org/SDL3/SDL_ShowWindow) | [:heavy_check_mark:](sdl/methods.go#L4354) | [:x:](sdl/sdl_functions_js.go#L6813) |
| [SDL_HideWindow](https://wiki.libsdl.org/SDL3/SDL_HideWindow) | [:heavy_check_mark:](sdl/methods.go#L4364) | [:x:](sdl/sdl_functions_js.go#L6829) |
| [SDL_RaiseWindow](https://wiki.libsdl.org/SDL3/SDL_RaiseWindow) | [:heavy_check_mark:](sdl/methods.go#L4374) | [:x:](sdl/sdl_functions_js.go#L6845) |
| [SDL_MaximizeWindow](https://wiki.libsdl.org/SDL3/SDL_MaximizeWindow) | [:heavy_check_mark:](sdl/methods.go#L4384) | [:x:](sdl/sdl_functions_js.go#L6861) |
| [SDL_MinimizeWindow](https://wiki.libsdl.org/SDL3/SDL_MinimizeWindow) | [:heavy_check_mark:](sdl/methods.go#L4394) | [:x:](sdl/sdl_functions_js.go#L6877) |
| [SDL_RestoreWindow](https://wiki.libsdl.org/SDL3/SDL_RestoreWindow) | [:heavy_check_mark:](sdl/methods.go#L4404) | [:x:](sdl/sdl_functions_js.go#L6893) |
| [SDL_SetWindowFullscreen](https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreen) | [:heavy_check_mark:](sdl/methods.go#L4414) | [:x:](sdl/sdl_functions_js.go#L6909) |
| [SDL_SyncWindow](https://wiki.libsdl.org/SDL3/SDL_SyncWindow) | [:heavy_check_mark:](sdl/methods.go#L4424) | [:x:](sdl/sdl_functions_js.go#L6927) |
| [SDL_WindowHasSurface](https://wiki.libsdl.org/SDL3/SDL_WindowHasSurface) | [:heavy_check_mark:](sdl/methods.go#L4434) | [:x:](sdl/sdl_functions_js.go#L6943) |
| [SDL_GetWindowSurface](https://wiki.libsdl.org/SDL3/SDL_GetWindowSurface) | [:heavy_check_mark:](sdl/methods.go#L4440) | [:x:](sdl/sdl_functions_js.go#L6959) |
| [SDL_SetWindowSurfaceVSync](https://wiki.libsdl.org/SDL3/SDL_SetWindowSurfaceVSync) | [:heavy_check_mark:](sdl/methods.go#L4451) | [:x:](sdl/sdl_functions_js.go#L6978) |
| [SDL_GetWindowSurfaceVSync](https://wiki.libsdl.org/SDL3/SDL_GetWindowSurfaceVSync) | [:heavy_check_mark:](sdl/methods.go#L4461) | [:x:](sdl/sdl_functions_js.go#L6996) |
| [SDL_UpdateWindowSurface](https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurface) | [:heavy_check_mark:](sdl/methods.go#L4473) | [:x:](sdl/sdl_functions_js.go#L7017) |
| [SDL_UpdateWindowSurfaceRects](https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurfaceRects) | [:heavy_check_mark:](sdl/methods.go#L4483) | [:x:](sdl/sdl_functions_js.go#L7033) |
| [SDL_DestroyWindowSurface](https://wiki.libsdl.org/SDL3/SDL_DestroyWindowSurface) | [:heavy_check_mark:](sdl/methods.go#L4493) | [:x:](sdl/sdl_functions_js.go#L7056) |
| [SDL_SetWindowKeyboardGrab](https://wiki.libsdl.org/SDL3/SDL_SetWindowKeyboardGrab) | [:heavy_check_mark:](sdl/methods.go#L4503) | [:x:](sdl/sdl_functions_js.go#L7072) |
| [SDL_SetWindowMouseGrab](https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseGrab) | [:heavy_check_mark:](sdl/methods.go#L4513) | [:x:](sdl/sdl_functions_js.go#L7090) |
| [SDL_GetWindowKeyboardGrab](https://wiki.libsdl.org/SDL3/SDL_GetWindowKeyboardGrab) | [:heavy_check_mark:](sdl/methods.go#L4523) | [:x:](sdl/sdl_functions_js.go#L7108) |
| [SDL_GetWindowMouseGrab](https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseGrab) | [:heavy_check_mark:](sdl/methods.go#L4529) | [:x:](sdl/sdl_functions_js.go#L7124) |
| [SDL_GetGrabbedWindow](https://wiki.libsdl.org/SDL3/SDL_GetGrabbedWindow) | [:heavy_check_mark:](sdl/functions.go#L588) | [:x:](sdl/sdl_functions_js.go#L7140) |
| [SDL_SetWindowMouseRect](https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseRect) | [:heavy_check_mark:](sdl/methods.go#L4535) | [:x:](sdl/sdl_functions_js.go#L7154) |
| [SDL_GetWindowMouseRect](https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseRect) | [:heavy_check_mark:](sdl/methods.go#L4545) | [:x:](sdl/sdl_functions_js.go#L7175) |
| [SDL_SetWindowOpacity](https://wiki.libsdl.org/SDL3/SDL_SetWindowOpacity) | [:heavy_check_mark:](sdl/methods.go#L4551) | [:x:](sdl/sdl_functions_js.go#L7194) |
| [SDL_GetWindowOpacity](https://wiki.libsdl.org/SDL3/SDL_GetWindowOpacity) | [:heavy_check_mark:](sdl/methods.go#L4561) | [:x:](sdl/sdl_functions_js.go#L7212) |
| [SDL_SetWindowParent](https://wiki.libsdl.org/SDL3/SDL_SetWindowParent) | [:heavy_check_mark:](sdl/methods.go#L4567) | [:x:](sdl/sdl_functions_js.go#L7228) |
| [SDL_SetWindowModal](https://wiki.libsdl.org/SDL3/SDL_SetWindowModal) | [:heavy_check_mark:](sdl/methods.go#L4577) | [:x:](sdl/sdl_functions_js.go#L7249) |
| [SDL_SetWindowFocusable](https://wiki.libsdl.org/SDL3/SDL_SetWindowFocusable) | [:heavy_check_mark:](sdl/methods.go#L4587) | [:x:](sdl/sdl_functions_js.go#L7267) |
| [SDL_ShowWindowSystemMenu](https://wiki.libsdl.org/SDL3/SDL_ShowWindowSystemMenu) | [:heavy_check_mark:](sdl/methods.go#L4597) | [:x:](sdl/sdl_functions_js.go#L7285) |
| [SDL_SetWindowHitTest](https://wiki.libsdl.org/SDL3/SDL_SetWindowHitTest) | [:x:](sdl/methods.go#L4607) | [:x:](sdl/sdl_functions_js.go#L7305) |
| [SDL_SetWindowShape](https://wiki.libsdl.org/SDL3/SDL_SetWindowShape) | [:heavy_check_mark:](sdl/methods.go#L4614) | [:x:](sdl/sdl_functions_js.go#L7325) |
| [SDL_FlashWindow](https://wiki.libsdl.org/SDL3/SDL_FlashWindow) | [:heavy_check_mark:](sdl/methods.go#L4624) | [:x:](sdl/sdl_functions_js.go#L7346) |
| [SDL_SetWindowProgressState](https://wiki.libsdl.org/SDL3/SDL_SetWindowProgressState) | [:question:]() | [:question:]() |
| [SDL_GetWindowProgressState](https://wiki.libsdl.org/SDL3/SDL_GetWindowProgressState) | [:question:]() | [:question:]() |
| [SDL_SetWindowProgressValue](https://wiki.libsdl.org/SDL3/SDL_SetWindowProgressValue) | [:question:]() | [:question:]() |
| [SDL_GetWindowProgressValue](https://wiki.libsdl.org/SDL3/SDL_GetWindowProgressValue) | [:question:]() | [:question:]() |
| [SDL_DestroyWindow](https://wiki.libsdl.org/SDL3/SDL_DestroyWindow) | [:heavy_check_mark:](sdl/methods.go#L4634) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7364) |
| [SDL_ScreenSaverEnabled](https://wiki.libsdl.org/SDL3/SDL_ScreenSaverEnabled) | [:heavy_check_mark:](sdl/functions.go#L594) | [:x:](sdl/sdl_functions_js.go#L7376) |
| [SDL_EnableScreenSaver](https://wiki.libsdl.org/SDL3/SDL_EnableScreenSaver) | [:heavy_check_mark:](sdl/functions.go#L600) | [:x:](sdl/sdl_functions_js.go#L7387) |
| [SDL_DisableScreenSaver](https://wiki.libsdl.org/SDL3/SDL_DisableScreenSaver) | [:heavy_check_mark:](sdl/functions.go#L610) | [:x:](sdl/sdl_functions_js.go#L7398) |
| [SDL_GL_LoadLibrary](https://wiki.libsdl.org/SDL3/SDL_GL_LoadLibrary) | [:heavy_check_mark:](sdl/functions.go#L620) | [:x:](sdl/sdl_functions_js.go#L7409) |
| [SDL_GL_GetProcAddress](https://wiki.libsdl.org/SDL3/SDL_GL_GetProcAddress) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7422) |
| [SDL_EGL_GetProcAddress](https://wiki.libsdl.org/SDL3/SDL_EGL_GetProcAddress) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7435) |
| [SDL_GL_UnloadLibrary](https://wiki.libsdl.org/SDL3/SDL_GL_UnloadLibrary) | [:heavy_check_mark:](sdl/functions.go#L636) | [:x:](sdl/sdl_functions_js.go#L7448) |
| [SDL_GL_ExtensionSupported](https://wiki.libsdl.org/SDL3/SDL_GL_ExtensionSupported) | [:heavy_check_mark:](sdl/functions.go#L630) | [:x:](sdl/sdl_functions_js.go#L7457) |
| [SDL_GL_ResetAttributes](https://wiki.libsdl.org/SDL3/SDL_GL_ResetAttributes) | [:heavy_check_mark:](sdl/functions.go#L685) | [:x:](sdl/sdl_functions_js.go#L7470) |
| [SDL_GL_SetAttribute](https://wiki.libsdl.org/SDL3/SDL_GL_SetAttribute) | [:heavy_check_mark:](sdl/functions.go#L691) | [:x:](sdl/sdl_functions_js.go#L7479) |
| [SDL_GL_GetAttribute](https://wiki.libsdl.org/SDL3/SDL_GL_GetAttribute) | [:heavy_check_mark:](sdl/functions.go#L701) | [:x:](sdl/sdl_functions_js.go#L7494) |
| [SDL_GL_CreateContext](https://wiki.libsdl.org/SDL3/SDL_GL_CreateContext) | [:heavy_check_mark:](sdl/functions.go#L642) | [:x:](sdl/sdl_functions_js.go#L7512) |
| [SDL_GL_MakeCurrent](https://wiki.libsdl.org/SDL3/SDL_GL_MakeCurrent) | [:heavy_check_mark:](sdl/functions.go#L653) | [:x:](sdl/sdl_functions_js.go#L7528) |
| [SDL_GL_GetCurrentWindow](https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentWindow) | [:heavy_check_mark:](sdl/functions.go#L713) | [:x:](sdl/sdl_functions_js.go#L7546) |
| [SDL_GL_GetCurrentContext](https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentContext) | [:heavy_check_mark:](sdl/functions.go#L724) | [:x:](sdl/sdl_functions_js.go#L7560) |
| [SDL_EGL_GetCurrentDisplay](https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentDisplay) | [:heavy_check_mark:](sdl/functions.go#L735) | [:x:](sdl/sdl_functions_js.go#L7571) |
| [SDL_EGL_GetCurrentConfig](https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentConfig) | [:heavy_check_mark:](sdl/functions.go#L746) | [:x:](sdl/sdl_functions_js.go#L7582) |
| [SDL_EGL_GetWindowSurface](https://wiki.libsdl.org/SDL3/SDL_EGL_GetWindowSurface) | [:heavy_check_mark:](sdl/functions.go#L663) | [:x:](sdl/sdl_functions_js.go#L7593) |
| [SDL_EGL_SetAttributeCallbacks](https://wiki.libsdl.org/SDL3/SDL_EGL_SetAttributeCallbacks) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7609) |
| [SDL_GL_SetSwapInterval](https://wiki.libsdl.org/SDL3/SDL_GL_SetSwapInterval) | [:heavy_check_mark:](sdl/functions.go#L757) | [:x:](sdl/sdl_functions_js.go#L7626) |
| [SDL_GL_GetSwapInterval](https://wiki.libsdl.org/SDL3/SDL_GL_GetSwapInterval) | [:heavy_check_mark:](sdl/functions.go#L767) | [:x:](sdl/sdl_functions_js.go#L7639) |
| [SDL_GL_SwapWindow](https://wiki.libsdl.org/SDL3/SDL_GL_SwapWindow) | [:heavy_check_mark:](sdl/functions.go#L669) | [:x:](sdl/sdl_functions_js.go#L7655) |
| [SDL_GL_DestroyContext](https://wiki.libsdl.org/SDL3/SDL_GL_DestroyContext) | [:heavy_check_mark:](sdl/functions.go#L679) | [:x:](sdl/sdl_functions_js.go#L7671) |
### Events

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_PumpEvents](https://wiki.libsdl.org/SDL3/SDL_PumpEvents) | [:heavy_check_mark:](sdl/functions.go#L184) | [:x:](sdl/sdl_functions_js.go#L10929) |
| [SDL_PeepEvents](https://wiki.libsdl.org/SDL3/SDL_PeepEvents) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10938) |
| [SDL_HasEvent](https://wiki.libsdl.org/SDL3/SDL_HasEvent) | [:heavy_check_mark:](sdl/functions.go#L192) | [:x:](sdl/sdl_functions_js.go#L10962) |
| [SDL_HasEvents](https://wiki.libsdl.org/SDL3/SDL_HasEvents) | [:heavy_check_mark:](sdl/functions.go#L198) | [:x:](sdl/sdl_functions_js.go#L10975) |
| [SDL_FlushEvent](https://wiki.libsdl.org/SDL3/SDL_FlushEvent) | [:heavy_check_mark:](sdl/functions.go#L204) | [:x:](sdl/sdl_functions_js.go#L10990) |
| [SDL_FlushEvents](https://wiki.libsdl.org/SDL3/SDL_FlushEvents) | [:heavy_check_mark:](sdl/functions.go#L210) | [:x:](sdl/sdl_functions_js.go#L11001) |
| [SDL_PollEvent](https://wiki.libsdl.org/SDL3/SDL_PollEvent) | [:heavy_check_mark:](sdl/functions.go#L216) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11014) |
| [SDL_WaitEvent](https://wiki.libsdl.org/SDL3/SDL_WaitEvent) | [:heavy_check_mark:](sdl/functions.go#L222) | [:x:](sdl/sdl_functions_js.go#L11027) |
| [SDL_WaitEventTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitEventTimeout) | [:heavy_check_mark:](sdl/functions.go#L232) | [:x:](sdl/sdl_functions_js.go#L11043) |
| [SDL_PushEvent](https://wiki.libsdl.org/SDL3/SDL_PushEvent) | [:heavy_check_mark:](sdl/functions.go#L238) | [:x:](sdl/sdl_functions_js.go#L11061) |
| [SDL_SetEventFilter](https://wiki.libsdl.org/SDL3/SDL_SetEventFilter) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11077) |
| [SDL_GetEventFilter](https://wiki.libsdl.org/SDL3/SDL_GetEventFilter) | [:x:](sdl/methods.go#L3695) | [:x:](sdl/sdl_functions_js.go#L11090) |
| [SDL_AddEventWatch](https://wiki.libsdl.org/SDL3/SDL_AddEventWatch) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11111) |
| [SDL_RemoveEventWatch](https://wiki.libsdl.org/SDL3/SDL_RemoveEventWatch) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11126) |
| [SDL_FilterEvents](https://wiki.libsdl.org/SDL3/SDL_FilterEvents) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11139) |
| [SDL_SetEventEnabled](https://wiki.libsdl.org/SDL3/SDL_SetEventEnabled) | [:heavy_check_mark:](sdl/functions.go#L254) | [:x:](sdl/sdl_functions_js.go#L11152) |
| [SDL_EventEnabled](https://wiki.libsdl.org/SDL3/SDL_EventEnabled) | [:heavy_check_mark:](sdl/functions.go#L260) | [:x:](sdl/sdl_functions_js.go#L11165) |
| [SDL_RegisterEvents](https://wiki.libsdl.org/SDL3/SDL_RegisterEvents) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11178) |
| [SDL_GetWindowFromEvent](https://wiki.libsdl.org/SDL3/SDL_GetWindowFromEvent) | [:heavy_check_mark:](sdl/methods.go#L1875) | [:x:](sdl/sdl_functions_js.go#L11191) |
| [SDL_GetEventDescription](https://wiki.libsdl.org/SDL3/SDL_GetEventDescription) | [:question:]() | [:question:]() |
### Keyboard

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_HasKeyboard](https://wiki.libsdl.org/SDL3/SDL_HasKeyboard) | [:heavy_check_mark:](sdl/functions.go#L1123) | [:x:](sdl/sdl_functions_js.go#L10155) |
| [SDL_GetKeyboards](https://wiki.libsdl.org/SDL3/SDL_GetKeyboards) | [:heavy_check_mark:](sdl/functions.go#L1129) | [:x:](sdl/sdl_functions_js.go#L10166) |
| [SDL_GetKeyboardNameForID](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardNameForID) | [:heavy_check_mark:](sdl/methods.go#L3669) | [:x:](sdl/sdl_functions_js.go#L10182) |
| [SDL_GetKeyboardFocus](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardFocus) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10195) |
| [SDL_GetKeyboardState](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardState) | [:heavy_check_mark:](sdl/functions.go#L1143) | [:x:](sdl/sdl_functions_js.go#L10209) |
| [SDL_ResetKeyboard](https://wiki.libsdl.org/SDL3/SDL_ResetKeyboard) | [:heavy_check_mark:](sdl/functions.go#L1153) | [:x:](sdl/sdl_functions_js.go#L10227) |
| [SDL_GetModState](https://wiki.libsdl.org/SDL3/SDL_GetModState) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10236) |
| [SDL_SetModState](https://wiki.libsdl.org/SDL3/SDL_SetModState) | [:x:](sdl/methods.go#L2759) | [:x:](sdl/sdl_functions_js.go#L10247) |
| [SDL_GetKeyFromScancode](https://wiki.libsdl.org/SDL3/SDL_GetKeyFromScancode) | [:heavy_check_mark:](sdl/methods.go#L4760) | [:x:](sdl/sdl_functions_js.go#L10258) |
| [SDL_GetScancodeFromKey](https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromKey) | [:heavy_check_mark:](sdl/methods.go#L5912) | [:x:](sdl/sdl_functions_js.go#L10275) |
| [SDL_SetScancodeName](https://wiki.libsdl.org/SDL3/SDL_SetScancodeName) | [:heavy_check_mark:](sdl/methods.go#L4766) | [:x:](sdl/sdl_functions_js.go#L10293) |
| [SDL_GetScancodeName](https://wiki.libsdl.org/SDL3/SDL_GetScancodeName) | [:heavy_check_mark:](sdl/methods.go#L4776) | [:x:](sdl/sdl_functions_js.go#L10308) |
| [SDL_GetScancodeFromName](https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromName) | [:heavy_check_mark:](sdl/functions.go#L1159) | [:x:](sdl/sdl_functions_js.go#L10321) |
| [SDL_GetKeyName](https://wiki.libsdl.org/SDL3/SDL_GetKeyName) | [:heavy_check_mark:](sdl/methods.go#L5918) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10334) |
| [SDL_GetKeyFromName](https://wiki.libsdl.org/SDL3/SDL_GetKeyFromName) | [:heavy_check_mark:](sdl/functions.go#L1165) | [:x:](sdl/sdl_functions_js.go#L10344) |
| [SDL_StartTextInput](https://wiki.libsdl.org/SDL3/SDL_StartTextInput) | [:heavy_check_mark:](sdl/methods.go#L4640) | [:x:](sdl/sdl_functions_js.go#L10357) |
| [SDL_StartTextInputWithProperties](https://wiki.libsdl.org/SDL3/SDL_StartTextInputWithProperties) | [:heavy_check_mark:](sdl/methods.go#L4650) | [:x:](sdl/sdl_functions_js.go#L10373) |
| [SDL_TextInputActive](https://wiki.libsdl.org/SDL3/SDL_TextInputActive) | [:heavy_check_mark:](sdl/methods.go#L4660) | [:x:](sdl/sdl_functions_js.go#L10391) |
| [SDL_StopTextInput](https://wiki.libsdl.org/SDL3/SDL_StopTextInput) | [:heavy_check_mark:](sdl/methods.go#L4666) | [:x:](sdl/sdl_functions_js.go#L10407) |
| [SDL_ClearComposition](https://wiki.libsdl.org/SDL3/SDL_ClearComposition) | [:x:](sdl/methods.go#L4676) | [:x:](sdl/sdl_functions_js.go#L10423) |
| [SDL_SetTextInputArea](https://wiki.libsdl.org/SDL3/SDL_SetTextInputArea) | [:heavy_check_mark:](sdl/methods.go#L4683) | [:x:](sdl/sdl_functions_js.go#L10439) |
| [SDL_GetTextInputArea](https://wiki.libsdl.org/SDL3/SDL_GetTextInputArea) | [:heavy_check_mark:](sdl/methods.go#L4693) | [:x:](sdl/sdl_functions_js.go#L10462) |
| [SDL_HasScreenKeyboardSupport](https://wiki.libsdl.org/SDL3/SDL_HasScreenKeyboardSupport) | [:heavy_check_mark:](sdl/functions.go#L1171) | [:x:](sdl/sdl_functions_js.go#L10488) |
| [SDL_ScreenKeyboardShown](https://wiki.libsdl.org/SDL3/SDL_ScreenKeyboardShown) | [:heavy_check_mark:](sdl/methods.go#L4706) | [:x:](sdl/sdl_functions_js.go#L10499) |
### Mouse

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_HasMouse](https://wiki.libsdl.org/SDL3/SDL_HasMouse) | [:heavy_check_mark:](sdl/functions.go#L1179) | [:x:](sdl/sdl_functions_js.go#L10515) |
| [SDL_GetMice](https://wiki.libsdl.org/SDL3/SDL_GetMice) | [:heavy_check_mark:](sdl/functions.go#L1185) | [:x:](sdl/sdl_functions_js.go#L10526) |
| [SDL_GetMouseNameForID](https://wiki.libsdl.org/SDL3/SDL_GetMouseNameForID) | [:heavy_check_mark:](sdl/methods.go#L3682) | [:x:](sdl/sdl_functions_js.go#L10542) |
| [SDL_GetMouseFocus](https://wiki.libsdl.org/SDL3/SDL_GetMouseFocus) | [:heavy_check_mark:](sdl/functions.go#L1199) | [:x:](sdl/sdl_functions_js.go#L10555) |
| [SDL_GetMouseState](https://wiki.libsdl.org/SDL3/SDL_GetMouseState) | [:heavy_check_mark:](sdl/functions.go#L1205) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10569) |
| [SDL_GetGlobalMouseState](https://wiki.libsdl.org/SDL3/SDL_GetGlobalMouseState) | [:heavy_check_mark:](sdl/functions.go#L1215) | [:x:](sdl/sdl_functions_js.go#L10585) |
| [SDL_GetRelativeMouseState](https://wiki.libsdl.org/SDL3/SDL_GetRelativeMouseState) | [:heavy_check_mark:](sdl/functions.go#L1225) | [:x:](sdl/sdl_functions_js.go#L10606) |
| [SDL_WarpMouseInWindow](https://wiki.libsdl.org/SDL3/SDL_WarpMouseInWindow) | [:heavy_check_mark:](sdl/methods.go#L4712) | [:x:](sdl/sdl_functions_js.go#L10627) |
| [SDL_WarpMouseGlobal](https://wiki.libsdl.org/SDL3/SDL_WarpMouseGlobal) | [:heavy_check_mark:](sdl/functions.go#L1235) | [:x:](sdl/sdl_functions_js.go#L10645) |
| [SDL_SetRelativeMouseTransform](https://wiki.libsdl.org/SDL3/SDL_SetRelativeMouseTransform) | [:question:]() | [:question:]() |
| [SDL_SetWindowRelativeMouseMode](https://wiki.libsdl.org/SDL3/SDL_SetWindowRelativeMouseMode) | [:heavy_check_mark:](sdl/methods.go#L4718) | [:x:](sdl/sdl_functions_js.go#L10660) |
| [SDL_GetWindowRelativeMouseMode](https://wiki.libsdl.org/SDL3/SDL_GetWindowRelativeMouseMode) | [:heavy_check_mark:](sdl/methods.go#L4728) | [:x:](sdl/sdl_functions_js.go#L10678) |
| [SDL_CaptureMouse](https://wiki.libsdl.org/SDL3/SDL_CaptureMouse) | [:heavy_check_mark:](sdl/functions.go#L1245) | [:x:](sdl/sdl_functions_js.go#L10694) |
| [SDL_CreateCursor](https://wiki.libsdl.org/SDL3/SDL_CreateCursor) | [:heavy_check_mark:](sdl/functions.go#L1255) | [:x:](sdl/sdl_functions_js.go#L10707) |
| [SDL_CreateColorCursor](https://wiki.libsdl.org/SDL3/SDL_CreateColorCursor) | [:heavy_check_mark:](sdl/methods.go#L1851) | [:x:](sdl/sdl_functions_js.go#L10739) |
| [SDL_CreateSystemCursor](https://wiki.libsdl.org/SDL3/SDL_CreateSystemCursor) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10762) |
| [SDL_SetCursor](https://wiki.libsdl.org/SDL3/SDL_SetCursor) | [:heavy_check_mark:](sdl/functions.go#L1270) | [:x:](sdl/sdl_functions_js.go#L10778) |
| [SDL_GetCursor](https://wiki.libsdl.org/SDL3/SDL_GetCursor) | [:heavy_check_mark:](sdl/functions.go#L1280) | [:x:](sdl/sdl_functions_js.go#L10794) |
| [SDL_GetDefaultCursor](https://wiki.libsdl.org/SDL3/SDL_GetDefaultCursor) | [:heavy_check_mark:](sdl/functions.go#L1286) | [:x:](sdl/sdl_functions_js.go#L10808) |
| [SDL_DestroyCursor](https://wiki.libsdl.org/SDL3/SDL_DestroyCursor) | [:heavy_check_mark:](sdl/methods.go#L593) | [:x:](sdl/sdl_functions_js.go#L10822) |
| [SDL_ShowCursor](https://wiki.libsdl.org/SDL3/SDL_ShowCursor) | [:heavy_check_mark:](sdl/functions.go#L1297) | [:x:](sdl/sdl_functions_js.go#L10836) |
| [SDL_HideCursor](https://wiki.libsdl.org/SDL3/SDL_HideCursor) | [:heavy_check_mark:](sdl/functions.go#L1307) | [:x:](sdl/sdl_functions_js.go#L10847) |
| [SDL_CursorVisible](https://wiki.libsdl.org/SDL3/SDL_CursorVisible) | [:heavy_check_mark:](sdl/functions.go#L1317) | [:x:](sdl/sdl_functions_js.go#L10858) |
### Touch

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetTouchDevices](https://wiki.libsdl.org/SDL3/SDL_GetTouchDevices) | [:heavy_check_mark:](sdl/functions.go#L1325) | [:x:](sdl/sdl_functions_js.go#L10869) |
| [SDL_GetTouchDeviceName](https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceName) | [:heavy_check_mark:](sdl/methods.go#L63) | [:x:](sdl/sdl_functions_js.go#L10885) |
| [SDL_GetTouchDeviceType](https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceType) | [:heavy_check_mark:](sdl/methods.go#L74) | [:x:](sdl/sdl_functions_js.go#L10898) |
| [SDL_GetTouchFingers](https://wiki.libsdl.org/SDL3/SDL_GetTouchFingers) | [:heavy_check_mark:](sdl/methods.go#L80) | [:x:](sdl/sdl_functions_js.go#L10911) |
### Gamepad

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AddGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMapping) | [:heavy_check_mark:](sdl/functions.go#L1341) | [:x:](sdl/sdl_functions_js.go#L8980) |
| [SDL_AddGamepadMappingsFromIO](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromIO) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8993) |
| [SDL_AddGamepadMappingsFromFile](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromFile) | [:heavy_check_mark:](sdl/functions.go#L1351) | [:x:](sdl/sdl_functions_js.go#L9011) |
| [SDL_ReloadGamepadMappings](https://wiki.libsdl.org/SDL3/SDL_ReloadGamepadMappings) | [:heavy_check_mark:](sdl/functions.go#L1361) | [:x:](sdl/sdl_functions_js.go#L9024) |
| [SDL_GetGamepadMappings](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappings) | [:heavy_check_mark:](sdl/functions.go#L1371) | [:x:](sdl/sdl_functions_js.go#L9035) |
| [SDL_GetGamepadMappingForGUID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForGUID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9051) |
| [SDL_GetGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMapping) | [:heavy_check_mark:](sdl/methods.go#L2449) | [:x:](sdl/sdl_functions_js.go#L9064) |
| [SDL_SetGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_SetGamepadMapping) | [:x:](sdl/methods.go#L755) | [:x:](sdl/sdl_functions_js.go#L9080) |
| [SDL_HasGamepad](https://wiki.libsdl.org/SDL3/SDL_HasGamepad) | [:heavy_check_mark:](sdl/functions.go#L1385) | [:x:](sdl/sdl_functions_js.go#L9095) |
| [SDL_GetGamepads](https://wiki.libsdl.org/SDL3/SDL_GetGamepads) | [:heavy_check_mark:](sdl/functions.go#L1391) | [:x:](sdl/sdl_functions_js.go#L9106) |
| [SDL_IsGamepad](https://wiki.libsdl.org/SDL3/SDL_IsGamepad) | [:heavy_check_mark:](sdl/methods.go#L762) | [:x:](sdl/sdl_functions_js.go#L9122) |
| [SDL_GetGamepadNameForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadNameForID) | [:heavy_check_mark:](sdl/methods.go#L768) | [:x:](sdl/sdl_functions_js.go#L9135) |
| [SDL_GetGamepadPathForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPathForID) | [:heavy_check_mark:](sdl/methods.go#L779) | [:x:](sdl/sdl_functions_js.go#L9148) |
| [SDL_GetGamepadPlayerIndexForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndexForID) | [:heavy_check_mark:](sdl/methods.go#L790) | [:x:](sdl/sdl_functions_js.go#L9161) |
| [SDL_GetGamepadGUIDForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadGUIDForID) | [:x:](sdl/methods.go#L796) | [:x:](sdl/sdl_functions_js.go#L9174) |
| [SDL_GetGamepadVendorForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendorForID) | [:heavy_check_mark:](sdl/methods.go#L803) | [:x:](sdl/sdl_functions_js.go#L9187) |
| [SDL_GetGamepadProductForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductForID) | [:heavy_check_mark:](sdl/methods.go#L809) | [:x:](sdl/sdl_functions_js.go#L9200) |
| [SDL_GetGamepadProductVersionForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersionForID) | [:heavy_check_mark:](sdl/methods.go#L815) | [:x:](sdl/sdl_functions_js.go#L9213) |
| [SDL_GetGamepadTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeForID) | [:heavy_check_mark:](sdl/methods.go#L821) | [:x:](sdl/sdl_functions_js.go#L9226) |
| [SDL_GetRealGamepadTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadTypeForID) | [:heavy_check_mark:](sdl/methods.go#L827) | [:x:](sdl/sdl_functions_js.go#L9239) |
| [SDL_GetGamepadMappingForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForID) | [:heavy_check_mark:](sdl/methods.go#L833) | [:x:](sdl/sdl_functions_js.go#L9252) |
| [SDL_OpenGamepad](https://wiki.libsdl.org/SDL3/SDL_OpenGamepad) | [:heavy_check_mark:](sdl/methods.go#L845) | [:x:](sdl/sdl_functions_js.go#L9265) |
| [SDL_GetGamepadFromID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromID) | [:heavy_check_mark:](sdl/methods.go#L856) | [:x:](sdl/sdl_functions_js.go#L9281) |
| [SDL_GetGamepadFromPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromPlayerIndex) | [:heavy_check_mark:](sdl/functions.go#L1405) | [:x:](sdl/sdl_functions_js.go#L9297) |
| [SDL_GetGamepadProperties](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProperties) | [:heavy_check_mark:](sdl/methods.go#L2461) | [:x:](sdl/sdl_functions_js.go#L9313) |
| [SDL_GetGamepadID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadID) | [:heavy_check_mark:](sdl/methods.go#L2472) | [:x:](sdl/sdl_functions_js.go#L9329) |
| [SDL_GetGamepadName](https://wiki.libsdl.org/SDL3/SDL_GetGamepadName) | [:heavy_check_mark:](sdl/methods.go#L2483) | [:x:](sdl/sdl_functions_js.go#L9345) |
| [SDL_GetGamepadPath](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPath) | [:heavy_check_mark:](sdl/methods.go#L2489) | [:x:](sdl/sdl_functions_js.go#L9361) |
| [SDL_GetGamepadType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadType) | [:heavy_check_mark:](sdl/methods.go#L2495) | [:x:](sdl/sdl_functions_js.go#L9377) |
| [SDL_GetRealGamepadType](https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadType) | [:heavy_check_mark:](sdl/methods.go#L2501) | [:x:](sdl/sdl_functions_js.go#L9393) |
| [SDL_GetGamepadPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L2507) | [:x:](sdl/sdl_functions_js.go#L9409) |
| [SDL_SetGamepadPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_SetGamepadPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L2513) | [:x:](sdl/sdl_functions_js.go#L9425) |
| [SDL_GetGamepadVendor](https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendor) | [:heavy_check_mark:](sdl/methods.go#L2523) | [:x:](sdl/sdl_functions_js.go#L9443) |
| [SDL_GetGamepadProduct](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProduct) | [:heavy_check_mark:](sdl/methods.go#L2529) | [:x:](sdl/sdl_functions_js.go#L9459) |
| [SDL_GetGamepadProductVersion](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersion) | [:heavy_check_mark:](sdl/methods.go#L2535) | [:x:](sdl/sdl_functions_js.go#L9475) |
| [SDL_GetGamepadFirmwareVersion](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFirmwareVersion) | [:heavy_check_mark:](sdl/methods.go#L2541) | [:x:](sdl/sdl_functions_js.go#L9491) |
| [SDL_GetGamepadSerial](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSerial) | [:heavy_check_mark:](sdl/methods.go#L2547) | [:x:](sdl/sdl_functions_js.go#L9507) |
| [SDL_GetGamepadSteamHandle](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSteamHandle) | [:heavy_check_mark:](sdl/methods.go#L2553) | [:x:](sdl/sdl_functions_js.go#L9523) |
| [SDL_GetGamepadConnectionState](https://wiki.libsdl.org/SDL3/SDL_GetGamepadConnectionState) | [:heavy_check_mark:](sdl/methods.go#L2559) | [:x:](sdl/sdl_functions_js.go#L9539) |
| [SDL_GetGamepadPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPowerInfo) | [:heavy_check_mark:](sdl/methods.go#L2571) | [:x:](sdl/sdl_functions_js.go#L9555) |
| [SDL_GamepadConnected](https://wiki.libsdl.org/SDL3/SDL_GamepadConnected) | [:heavy_check_mark:](sdl/methods.go#L2581) | [:x:](sdl/sdl_functions_js.go#L9576) |
| [SDL_GetGamepadJoystick](https://wiki.libsdl.org/SDL3/SDL_GetGamepadJoystick) | [:heavy_check_mark:](sdl/methods.go#L2587) | [:x:](sdl/sdl_functions_js.go#L9592) |
| [SDL_SetGamepadEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_SetGamepadEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L1411) | [:x:](sdl/sdl_functions_js.go#L9611) |
| [SDL_GamepadEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_GamepadEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L1417) | [:x:](sdl/sdl_functions_js.go#L9622) |
| [SDL_GetGamepadBindings](https://wiki.libsdl.org/SDL3/SDL_GetGamepadBindings) | [:heavy_check_mark:](sdl/methods.go#L2598) | [:x:](sdl/sdl_functions_js.go#L9633) |
| [SDL_UpdateGamepads](https://wiki.libsdl.org/SDL3/SDL_UpdateGamepads) | [:heavy_check_mark:](sdl/functions.go#L1423) | [:x:](sdl/sdl_functions_js.go#L9654) |
| [SDL_GetGamepadTypeFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeFromString) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9663) |
| [SDL_GetGamepadStringForType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForType) | [:heavy_check_mark:](sdl/methods.go#L5601) | [:x:](sdl/sdl_functions_js.go#L9676) |
| [SDL_GetGamepadAxisFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxisFromString) | [:heavy_check_mark:](sdl/functions.go#L1429) | [:x:](sdl/sdl_functions_js.go#L9689) |
| [SDL_GetGamepadStringForAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForAxis) | [:heavy_check_mark:](sdl/methods.go#L585) | [:x:](sdl/sdl_functions_js.go#L9702) |
| [SDL_GamepadHasAxis](https://wiki.libsdl.org/SDL3/SDL_GamepadHasAxis) | [:heavy_check_mark:](sdl/methods.go#L2612) | [:x:](sdl/sdl_functions_js.go#L9715) |
| [SDL_GetGamepadAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxis) | [:heavy_check_mark:](sdl/methods.go#L2618) | [:x:](sdl/sdl_functions_js.go#L9733) |
| [SDL_GetGamepadButtonFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonFromString) | [:heavy_check_mark:](sdl/functions.go#L1435) | [:x:](sdl/sdl_functions_js.go#L9751) |
| [SDL_GetGamepadStringForButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForButton) | [:heavy_check_mark:](sdl/methods.go#L443) | [:x:](sdl/sdl_functions_js.go#L9764) |
| [SDL_GamepadHasButton](https://wiki.libsdl.org/SDL3/SDL_GamepadHasButton) | [:heavy_check_mark:](sdl/methods.go#L2624) | [:x:](sdl/sdl_functions_js.go#L9777) |
| [SDL_GetGamepadButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButton) | [:heavy_check_mark:](sdl/methods.go#L2630) | [:x:](sdl/sdl_functions_js.go#L9795) |
| [SDL_GetGamepadButtonLabelForType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabelForType) | [:heavy_check_mark:](sdl/methods.go#L5607) | [:x:](sdl/sdl_functions_js.go#L9813) |
| [SDL_GetGamepadButtonLabel](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabel) | [:heavy_check_mark:](sdl/methods.go#L2636) | [:x:](sdl/sdl_functions_js.go#L9828) |
| [SDL_GetNumGamepadTouchpads](https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpads) | [:heavy_check_mark:](sdl/methods.go#L2642) | [:x:](sdl/sdl_functions_js.go#L9846) |
| [SDL_GetNumGamepadTouchpadFingers](https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpadFingers) | [:heavy_check_mark:](sdl/methods.go#L2648) | [:x:](sdl/sdl_functions_js.go#L9862) |
| [SDL_GetGamepadTouchpadFinger](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTouchpadFinger) | [:x:](sdl/methods.go#L2654) | [:x:](sdl/sdl_functions_js.go#L9880) |
| [SDL_GamepadHasSensor](https://wiki.libsdl.org/SDL3/SDL_GamepadHasSensor) | [:heavy_check_mark:](sdl/methods.go#L2661) | [:x:](sdl/sdl_functions_js.go#L9920) |
| [SDL_SetGamepadSensorEnabled](https://wiki.libsdl.org/SDL3/SDL_SetGamepadSensorEnabled) | [:heavy_check_mark:](sdl/methods.go#L2667) | [:x:](sdl/sdl_functions_js.go#L9938) |
| [SDL_GamepadSensorEnabled](https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEnabled) | [:heavy_check_mark:](sdl/methods.go#L2677) | [:x:](sdl/sdl_functions_js.go#L9958) |
| [SDL_GetGamepadSensorDataRate](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorDataRate) | [:heavy_check_mark:](sdl/methods.go#L2683) | [:x:](sdl/sdl_functions_js.go#L9976) |
| [SDL_GetGamepadSensorData](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorData) | [:x:](sdl/methods.go#L2689) | [:x:](sdl/sdl_functions_js.go#L9994) |
| [SDL_RumbleGamepad](https://wiki.libsdl.org/SDL3/SDL_RumbleGamepad) | [:heavy_check_mark:](sdl/methods.go#L2696) | [:x:](sdl/sdl_functions_js.go#L10019) |
| [SDL_RumbleGamepadTriggers](https://wiki.libsdl.org/SDL3/SDL_RumbleGamepadTriggers) | [:heavy_check_mark:](sdl/methods.go#L2706) | [:x:](sdl/sdl_functions_js.go#L10041) |
| [SDL_SetGamepadLED](https://wiki.libsdl.org/SDL3/SDL_SetGamepadLED) | [:heavy_check_mark:](sdl/methods.go#L2716) | [:x:](sdl/sdl_functions_js.go#L10063) |
| [SDL_SendGamepadEffect](https://wiki.libsdl.org/SDL3/SDL_SendGamepadEffect) | [:heavy_check_mark:](sdl/methods.go#L2726) | [:x:](sdl/sdl_functions_js.go#L10085) |
| [SDL_CloseGamepad](https://wiki.libsdl.org/SDL3/SDL_CloseGamepad) | [:heavy_check_mark:](sdl/methods.go#L2737) | [:x:](sdl/sdl_functions_js.go#L10105) |
| [SDL_GetGamepadAppleSFSymbolsNameForButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForButton) | [:x:](sdl/methods.go#L2743) | [:x:](sdl/sdl_functions_js.go#L10119) |
| [SDL_GetGamepadAppleSFSymbolsNameForAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForAxis) | [:x:](sdl/methods.go#L2750) | [:x:](sdl/sdl_functions_js.go#L10137) |
### Joystick

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_LockJoysticks](https://wiki.libsdl.org/SDL3/SDL_LockJoysticks) | [:heavy_check_mark:](sdl/functions.go#L1443) | [:x:](sdl/sdl_functions_js.go#L8041) |
| [SDL_UnlockJoysticks](https://wiki.libsdl.org/SDL3/SDL_UnlockJoysticks) | [:heavy_check_mark:](sdl/functions.go#L1449) | [:x:](sdl/sdl_functions_js.go#L8050) |
| [SDL_HasJoystick](https://wiki.libsdl.org/SDL3/SDL_HasJoystick) | [:heavy_check_mark:](sdl/functions.go#L1455) | [:x:](sdl/sdl_functions_js.go#L8059) |
| [SDL_GetJoysticks](https://wiki.libsdl.org/SDL3/SDL_GetJoysticks) | [:heavy_check_mark:](sdl/functions.go#L1461) | [:x:](sdl/sdl_functions_js.go#L8070) |
| [SDL_GetJoystickNameForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickNameForID) | [:heavy_check_mark:](sdl/methods.go#L658) | [:x:](sdl/sdl_functions_js.go#L8086) |
| [SDL_GetJoystickPathForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPathForID) | [:heavy_check_mark:](sdl/methods.go#L669) | [:x:](sdl/sdl_functions_js.go#L8099) |
| [SDL_GetJoystickPlayerIndexForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndexForID) | [:heavy_check_mark:](sdl/methods.go#L680) | [:x:](sdl/sdl_functions_js.go#L8112) |
| [SDL_GetJoystickGUIDForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDForID) | [:x:](sdl/methods.go#L686) | [:x:](sdl/sdl_functions_js.go#L8125) |
| [SDL_GetJoystickVendorForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendorForID) | [:heavy_check_mark:](sdl/methods.go#L693) | [:x:](sdl/sdl_functions_js.go#L8138) |
| [SDL_GetJoystickProductForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductForID) | [:heavy_check_mark:](sdl/methods.go#L699) | [:x:](sdl/sdl_functions_js.go#L8151) |
| [SDL_GetJoystickProductVersionForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersionForID) | [:heavy_check_mark:](sdl/methods.go#L705) | [:x:](sdl/sdl_functions_js.go#L8164) |
| [SDL_GetJoystickTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickTypeForID) | [:heavy_check_mark:](sdl/methods.go#L711) | [:x:](sdl/sdl_functions_js.go#L8177) |
| [SDL_OpenJoystick](https://wiki.libsdl.org/SDL3/SDL_OpenJoystick) | [:heavy_check_mark:](sdl/methods.go#L717) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8190) |
| [SDL_GetJoystickFromID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromID) | [:heavy_check_mark:](sdl/methods.go#L728) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8202) |
| [SDL_GetJoystickFromPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromPlayerIndex) | [:heavy_check_mark:](sdl/functions.go#L1475) | [:x:](sdl/sdl_functions_js.go#L8216) |
| [SDL_AttachVirtualJoystick](https://wiki.libsdl.org/SDL3/SDL_AttachVirtualJoystick) | [:heavy_check_mark:](sdl/functions.go#L1481) | [:x:](sdl/sdl_functions_js.go#L8232) |
| [SDL_DetachVirtualJoystick](https://wiki.libsdl.org/SDL3/SDL_DetachVirtualJoystick) | [:heavy_check_mark:](sdl/methods.go#L739) | [:x:](sdl/sdl_functions_js.go#L8248) |
| [SDL_IsJoystickVirtual](https://wiki.libsdl.org/SDL3/SDL_IsJoystickVirtual) | [:heavy_check_mark:](sdl/methods.go#L749) | [:x:](sdl/sdl_functions_js.go#L8261) |
| [SDL_SetJoystickVirtualAxis](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualAxis) | [:heavy_check_mark:](sdl/methods.go#L5259) | [:x:](sdl/sdl_functions_js.go#L8274) |
| [SDL_SetJoystickVirtualBall](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualBall) | [:heavy_check_mark:](sdl/methods.go#L5269) | [:x:](sdl/sdl_functions_js.go#L8294) |
| [SDL_SetJoystickVirtualButton](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualButton) | [:heavy_check_mark:](sdl/methods.go#L5279) | [:x:](sdl/sdl_functions_js.go#L8316) |
| [SDL_SetJoystickVirtualHat](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualHat) | [:heavy_check_mark:](sdl/methods.go#L5289) | [:x:](sdl/sdl_functions_js.go#L8336) |
| [SDL_SetJoystickVirtualTouchpad](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualTouchpad) | [:heavy_check_mark:](sdl/methods.go#L5299) | [:x:](sdl/sdl_functions_js.go#L8356) |
| [SDL_SendJoystickVirtualSensorData](https://wiki.libsdl.org/SDL3/SDL_SendJoystickVirtualSensorData) | [:heavy_check_mark:](sdl/methods.go#L5309) | [:x:](sdl/sdl_functions_js.go#L8384) |
| [SDL_GetJoystickProperties](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProperties) | [:heavy_check_mark:](sdl/methods.go#L5319) | [:x:](sdl/sdl_functions_js.go#L8411) |
| [SDL_GetJoystickName](https://wiki.libsdl.org/SDL3/SDL_GetJoystickName) | [:heavy_check_mark:](sdl/methods.go#L5330) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8427) |
| [SDL_GetJoystickPath](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPath) | [:heavy_check_mark:](sdl/methods.go#L5341) | [:x:](sdl/sdl_functions_js.go#L8440) |
| [SDL_GetJoystickPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L5352) | [:x:](sdl/sdl_functions_js.go#L8456) |
| [SDL_SetJoystickPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_SetJoystickPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L5358) | [:x:](sdl/sdl_functions_js.go#L8472) |
| [SDL_GetJoystickGUID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUID) | [:x:](sdl/methods.go#L5368) | [:x:](sdl/sdl_functions_js.go#L8490) |
| [SDL_GetJoystickVendor](https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendor) | [:heavy_check_mark:](sdl/methods.go#L5375) | [:x:](sdl/sdl_functions_js.go#L8506) |
| [SDL_GetJoystickProduct](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProduct) | [:heavy_check_mark:](sdl/methods.go#L5381) | [:x:](sdl/sdl_functions_js.go#L8522) |
| [SDL_GetJoystickProductVersion](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersion) | [:heavy_check_mark:](sdl/methods.go#L5387) | [:x:](sdl/sdl_functions_js.go#L8538) |
| [SDL_GetJoystickFirmwareVersion](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFirmwareVersion) | [:heavy_check_mark:](sdl/methods.go#L5393) | [:x:](sdl/sdl_functions_js.go#L8554) |
| [SDL_GetJoystickSerial](https://wiki.libsdl.org/SDL3/SDL_GetJoystickSerial) | [:heavy_check_mark:](sdl/methods.go#L5399) | [:x:](sdl/sdl_functions_js.go#L8570) |
| [SDL_GetJoystickType](https://wiki.libsdl.org/SDL3/SDL_GetJoystickType) | [:heavy_check_mark:](sdl/methods.go#L5405) | [:x:](sdl/sdl_functions_js.go#L8586) |
| [SDL_GetJoystickGUIDInfo](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDInfo) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8602) |
| [SDL_JoystickConnected](https://wiki.libsdl.org/SDL3/SDL_JoystickConnected) | [:heavy_check_mark:](sdl/methods.go#L5411) | [:x:](sdl/sdl_functions_js.go#L8633) |
| [SDL_GetJoystickID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickID) | [:heavy_check_mark:](sdl/methods.go#L5417) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8649) |
| [SDL_GetNumJoystickAxes](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickAxes) | [:heavy_check_mark:](sdl/methods.go#L5428) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8662) |
| [SDL_GetNumJoystickBalls](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickBalls) | [:heavy_check_mark:](sdl/methods.go#L5439) | [:x:](sdl/sdl_functions_js.go#L8675) |
| [SDL_GetNumJoystickHats](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickHats) | [:heavy_check_mark:](sdl/methods.go#L5450) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8691) |
| [SDL_GetNumJoystickButtons](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickButtons) | [:heavy_check_mark:](sdl/methods.go#L5461) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8704) |
| [SDL_SetJoystickEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_SetJoystickEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L1487) | [:x:](sdl/sdl_functions_js.go#L8717) |
| [SDL_JoystickEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_JoystickEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L1493) | [:x:](sdl/sdl_functions_js.go#L8728) |
| [SDL_UpdateJoysticks](https://wiki.libsdl.org/SDL3/SDL_UpdateJoysticks) | [:heavy_check_mark:](sdl/functions.go#L1499) | [:x:](sdl/sdl_functions_js.go#L8739) |
| [SDL_GetJoystickAxis](https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxis) | [:heavy_check_mark:](sdl/methods.go#L5472) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8748) |
| [SDL_GetJoystickAxisInitialState](https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxisInitialState) | [:heavy_check_mark:](sdl/methods.go#L5483) | [:x:](sdl/sdl_functions_js.go#L8763) |
| [SDL_GetJoystickBall](https://wiki.libsdl.org/SDL3/SDL_GetJoystickBall) | [:heavy_check_mark:](sdl/methods.go#L5492) | [:x:](sdl/sdl_functions_js.go#L8786) |
| [SDL_GetJoystickHat](https://wiki.libsdl.org/SDL3/SDL_GetJoystickHat) | [:heavy_check_mark:](sdl/methods.go#L5504) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8814) |
| [SDL_GetJoystickButton](https://wiki.libsdl.org/SDL3/SDL_GetJoystickButton) | [:heavy_check_mark:](sdl/methods.go#L5510) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8829) |
| [SDL_RumbleJoystick](https://wiki.libsdl.org/SDL3/SDL_RumbleJoystick) | [:heavy_check_mark:](sdl/methods.go#L5516) | [:x:](sdl/sdl_functions_js.go#L8844) |
| [SDL_RumbleJoystickTriggers](https://wiki.libsdl.org/SDL3/SDL_RumbleJoystickTriggers) | [:heavy_check_mark:](sdl/methods.go#L5522) | [:x:](sdl/sdl_functions_js.go#L8866) |
| [SDL_SetJoystickLED](https://wiki.libsdl.org/SDL3/SDL_SetJoystickLED) | [:heavy_check_mark:](sdl/methods.go#L5532) | [:x:](sdl/sdl_functions_js.go#L8888) |
| [SDL_SendJoystickEffect](https://wiki.libsdl.org/SDL3/SDL_SendJoystickEffect) | [:heavy_check_mark:](sdl/methods.go#L5542) | [:x:](sdl/sdl_functions_js.go#L8910) |
| [SDL_CloseJoystick](https://wiki.libsdl.org/SDL3/SDL_CloseJoystick) | [:heavy_check_mark:](sdl/methods.go#L5552) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8930) |
| [SDL_GetJoystickConnectionState](https://wiki.libsdl.org/SDL3/SDL_GetJoystickConnectionState) | [:heavy_check_mark:](sdl/methods.go#L5558) | [:x:](sdl/sdl_functions_js.go#L8943) |
| [SDL_GetJoystickPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPowerInfo) | [:heavy_check_mark:](sdl/methods.go#L5569) | [:x:](sdl/sdl_functions_js.go#L8959) |
### Haptic

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetHaptics](https://wiki.libsdl.org/SDL3/SDL_GetHaptics) | [:heavy_check_mark:](sdl/functions.go#L1507) | [:x:](sdl/sdl_functions_js.go#L13265) |
| [SDL_GetHapticNameForID](https://wiki.libsdl.org/SDL3/SDL_GetHapticNameForID) | [:heavy_check_mark:](sdl/methods.go#L2819) | [:x:](sdl/sdl_functions_js.go#L13281) |
| [SDL_OpenHaptic](https://wiki.libsdl.org/SDL3/SDL_OpenHaptic) | [:heavy_check_mark:](sdl/methods.go#L2830) | [:x:](sdl/sdl_functions_js.go#L13294) |
| [SDL_GetHapticFromID](https://wiki.libsdl.org/SDL3/SDL_GetHapticFromID) | [:heavy_check_mark:](sdl/methods.go#L2841) | [:x:](sdl/sdl_functions_js.go#L13310) |
| [SDL_GetHapticID](https://wiki.libsdl.org/SDL3/SDL_GetHapticID) | [:heavy_check_mark:](sdl/methods.go#L2221) | [:x:](sdl/sdl_functions_js.go#L13326) |
| [SDL_GetHapticName](https://wiki.libsdl.org/SDL3/SDL_GetHapticName) | [:heavy_check_mark:](sdl/methods.go#L2232) | [:x:](sdl/sdl_functions_js.go#L13342) |
| [SDL_IsMouseHaptic](https://wiki.libsdl.org/SDL3/SDL_IsMouseHaptic) | [:heavy_check_mark:](sdl/functions.go#L1521) | [:x:](sdl/sdl_functions_js.go#L13358) |
| [SDL_OpenHapticFromMouse](https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromMouse) | [:heavy_check_mark:](sdl/functions.go#L1527) | [:x:](sdl/sdl_functions_js.go#L13369) |
| [SDL_IsJoystickHaptic](https://wiki.libsdl.org/SDL3/SDL_IsJoystickHaptic) | [:heavy_check_mark:](sdl/methods.go#L5582) | [:x:](sdl/sdl_functions_js.go#L13383) |
| [SDL_OpenHapticFromJoystick](https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromJoystick) | [:heavy_check_mark:](sdl/methods.go#L5588) | [:x:](sdl/sdl_functions_js.go#L13399) |
| [SDL_CloseHaptic](https://wiki.libsdl.org/SDL3/SDL_CloseHaptic) | [:heavy_check_mark:](sdl/methods.go#L2243) | [:x:](sdl/sdl_functions_js.go#L13418) |
| [SDL_GetMaxHapticEffects](https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffects) | [:heavy_check_mark:](sdl/methods.go#L2249) | [:x:](sdl/sdl_functions_js.go#L13432) |
| [SDL_GetMaxHapticEffectsPlaying](https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffectsPlaying) | [:heavy_check_mark:](sdl/methods.go#L2260) | [:x:](sdl/sdl_functions_js.go#L13448) |
| [SDL_GetHapticFeatures](https://wiki.libsdl.org/SDL3/SDL_GetHapticFeatures) | [:heavy_check_mark:](sdl/methods.go#L2271) | [:x:](sdl/sdl_functions_js.go#L13464) |
| [SDL_GetNumHapticAxes](https://wiki.libsdl.org/SDL3/SDL_GetNumHapticAxes) | [:heavy_check_mark:](sdl/methods.go#L2282) | [:x:](sdl/sdl_functions_js.go#L13480) |
| [SDL_HapticEffectSupported](https://wiki.libsdl.org/SDL3/SDL_HapticEffectSupported) | [:heavy_check_mark:](sdl/methods.go#L2293) | [:x:](sdl/sdl_functions_js.go#L13496) |
| [SDL_CreateHapticEffect](https://wiki.libsdl.org/SDL3/SDL_CreateHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2299) | [:x:](sdl/sdl_functions_js.go#L13517) |
| [SDL_UpdateHapticEffect](https://wiki.libsdl.org/SDL3/SDL_UpdateHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2310) | [:x:](sdl/sdl_functions_js.go#L13538) |
| [SDL_RunHapticEffect](https://wiki.libsdl.org/SDL3/SDL_RunHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2320) | [:x:](sdl/sdl_functions_js.go#L13561) |
| [SDL_StopHapticEffect](https://wiki.libsdl.org/SDL3/SDL_StopHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2330) | [:x:](sdl/sdl_functions_js.go#L13581) |
| [SDL_DestroyHapticEffect](https://wiki.libsdl.org/SDL3/SDL_DestroyHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2340) | [:x:](sdl/sdl_functions_js.go#L13599) |
| [SDL_GetHapticEffectStatus](https://wiki.libsdl.org/SDL3/SDL_GetHapticEffectStatus) | [:heavy_check_mark:](sdl/methods.go#L2346) | [:x:](sdl/sdl_functions_js.go#L13615) |
| [SDL_SetHapticGain](https://wiki.libsdl.org/SDL3/SDL_SetHapticGain) | [:heavy_check_mark:](sdl/methods.go#L2352) | [:x:](sdl/sdl_functions_js.go#L13633) |
| [SDL_SetHapticAutocenter](https://wiki.libsdl.org/SDL3/SDL_SetHapticAutocenter) | [:heavy_check_mark:](sdl/methods.go#L2362) | [:x:](sdl/sdl_functions_js.go#L13651) |
| [SDL_PauseHaptic](https://wiki.libsdl.org/SDL3/SDL_PauseHaptic) | [:heavy_check_mark:](sdl/methods.go#L2372) | [:x:](sdl/sdl_functions_js.go#L13669) |
| [SDL_ResumeHaptic](https://wiki.libsdl.org/SDL3/SDL_ResumeHaptic) | [:heavy_check_mark:](sdl/methods.go#L2382) | [:x:](sdl/sdl_functions_js.go#L13685) |
| [SDL_StopHapticEffects](https://wiki.libsdl.org/SDL3/SDL_StopHapticEffects) | [:heavy_check_mark:](sdl/methods.go#L2392) | [:x:](sdl/sdl_functions_js.go#L13701) |
| [SDL_HapticRumbleSupported](https://wiki.libsdl.org/SDL3/SDL_HapticRumbleSupported) | [:heavy_check_mark:](sdl/methods.go#L2402) | [:x:](sdl/sdl_functions_js.go#L13717) |
| [SDL_InitHapticRumble](https://wiki.libsdl.org/SDL3/SDL_InitHapticRumble) | [:heavy_check_mark:](sdl/methods.go#L2408) | [:x:](sdl/sdl_functions_js.go#L13733) |
| [SDL_PlayHapticRumble](https://wiki.libsdl.org/SDL3/SDL_PlayHapticRumble) | [:heavy_check_mark:](sdl/methods.go#L2418) | [:x:](sdl/sdl_functions_js.go#L13749) |
| [SDL_StopHapticRumble](https://wiki.libsdl.org/SDL3/SDL_StopHapticRumble) | [:heavy_check_mark:](sdl/methods.go#L2428) | [:x:](sdl/sdl_functions_js.go#L13769) |
### Audio

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumAudioDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumAudioDrivers) | [:heavy_check_mark:](sdl/functions.go#L780) | [:x:](sdl/sdl_functions_js.go#L2357) |
| [SDL_GetAudioDriver](https://wiki.libsdl.org/SDL3/SDL_GetAudioDriver) | [:heavy_check_mark:](sdl/functions.go#L786) | [:x:](sdl/sdl_functions_js.go#L2368) |
| [SDL_GetCurrentAudioDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentAudioDriver) | [:heavy_check_mark:](sdl/functions.go#L792) | [:x:](sdl/sdl_functions_js.go#L2381) |
| [SDL_GetAudioPlaybackDevices](https://wiki.libsdl.org/SDL3/SDL_GetAudioPlaybackDevices) | [:heavy_check_mark:](sdl/functions.go#L798) | [:x:](sdl/sdl_functions_js.go#L2392) |
| [SDL_GetAudioRecordingDevices](https://wiki.libsdl.org/SDL3/SDL_GetAudioRecordingDevices) | [:heavy_check_mark:](sdl/functions.go#L812) | [:x:](sdl/sdl_functions_js.go#L2408) |
| [SDL_GetAudioDeviceName](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceName) | [:heavy_check_mark:](sdl/methods.go#L226) | [:x:](sdl/sdl_functions_js.go#L2424) |
| [SDL_GetAudioDeviceFormat](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceFormat) | [:heavy_check_mark:](sdl/methods.go#L237) | [:x:](sdl/sdl_functions_js.go#L2437) |
| [SDL_GetAudioDeviceChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceChannelMap) | [:heavy_check_mark:](sdl/methods.go#L250) | [:x:](sdl/sdl_functions_js.go#L2460) |
| [SDL_OpenAudioDevice](https://wiki.libsdl.org/SDL3/SDL_OpenAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L264) | [:x:](sdl/sdl_functions_js.go#L2478) |
| [SDL_IsAudioDevicePhysical](https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePhysical) | [:heavy_check_mark:](sdl/methods.go#L275) | [:x:](sdl/sdl_functions_js.go#L2496) |
| [SDL_IsAudioDevicePlayback](https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePlayback) | [:heavy_check_mark:](sdl/methods.go#L281) | [:x:](sdl/sdl_functions_js.go#L2509) |
| [SDL_PauseAudioDevice](https://wiki.libsdl.org/SDL3/SDL_PauseAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L287) | [:x:](sdl/sdl_functions_js.go#L2522) |
| [SDL_ResumeAudioDevice](https://wiki.libsdl.org/SDL3/SDL_ResumeAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L297) | [:x:](sdl/sdl_functions_js.go#L2535) |
| [SDL_AudioDevicePaused](https://wiki.libsdl.org/SDL3/SDL_AudioDevicePaused) | [:heavy_check_mark:](sdl/methods.go#L307) | [:x:](sdl/sdl_functions_js.go#L2548) |
| [SDL_GetAudioDeviceGain](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceGain) | [:heavy_check_mark:](sdl/methods.go#L313) | [:x:](sdl/sdl_functions_js.go#L2561) |
| [SDL_SetAudioDeviceGain](https://wiki.libsdl.org/SDL3/SDL_SetAudioDeviceGain) | [:heavy_check_mark:](sdl/methods.go#L324) | [:x:](sdl/sdl_functions_js.go#L2574) |
| [SDL_CloseAudioDevice](https://wiki.libsdl.org/SDL3/SDL_CloseAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L334) | [:x:](sdl/sdl_functions_js.go#L2589) |
| [SDL_BindAudioStreams](https://wiki.libsdl.org/SDL3/SDL_BindAudioStreams) | [:heavy_check_mark:](sdl/methods.go#L340) | [:x:](sdl/sdl_functions_js.go#L2600) |
| [SDL_BindAudioStream](https://wiki.libsdl.org/SDL3/SDL_BindAudioStream) | [:heavy_check_mark:](sdl/methods.go#L350) | [:x:](sdl/sdl_functions_js.go#L2620) |
| [SDL_UnbindAudioStreams](https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStreams) | [:heavy_check_mark:](sdl/functions.go#L826) | [:x:](sdl/sdl_functions_js.go#L2638) |
| [SDL_UnbindAudioStream](https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStream) | [:heavy_check_mark:](sdl/methods.go#L3704) | [:x:](sdl/sdl_functions_js.go#L2654) |
| [SDL_GetAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamDevice) | [:heavy_check_mark:](sdl/methods.go#L3710) | [:x:](sdl/sdl_functions_js.go#L2668) |
| [SDL_CreateAudioStream](https://wiki.libsdl.org/SDL3/SDL_CreateAudioStream) | [:heavy_check_mark:](sdl/functions.go#L832) | [:x:](sdl/sdl_functions_js.go#L2684) |
| [SDL_GetAudioStreamProperties](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamProperties) | [:heavy_check_mark:](sdl/methods.go#L3716) | [:x:](sdl/sdl_functions_js.go#L2708) |
| [SDL_GetAudioStreamFormat](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFormat) | [:heavy_check_mark:](sdl/methods.go#L3727) | [:x:](sdl/sdl_functions_js.go#L2724) |
| [SDL_SetAudioStreamFormat](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFormat) | [:heavy_check_mark:](sdl/methods.go#L3737) | [:x:](sdl/sdl_functions_js.go#L2750) |
| [SDL_GetAudioStreamFrequencyRatio](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFrequencyRatio) | [:heavy_check_mark:](sdl/methods.go#L3747) | [:x:](sdl/sdl_functions_js.go#L2776) |
| [SDL_SetAudioStreamFrequencyRatio](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFrequencyRatio) | [:heavy_check_mark:](sdl/methods.go#L3758) | [:x:](sdl/sdl_functions_js.go#L2792) |
| [SDL_GetAudioStreamGain](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamGain) | [:heavy_check_mark:](sdl/methods.go#L3768) | [:x:](sdl/sdl_functions_js.go#L2810) |
| [SDL_SetAudioStreamGain](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGain) | [:heavy_check_mark:](sdl/methods.go#L3779) | [:x:](sdl/sdl_functions_js.go#L2826) |
| [SDL_GetAudioStreamInputChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamInputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3789) | [:x:](sdl/sdl_functions_js.go#L2844) |
| [SDL_GetAudioStreamOutputChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamOutputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3803) | [:x:](sdl/sdl_functions_js.go#L2865) |
| [SDL_SetAudioStreamInputChannelMap](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamInputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3817) | [:x:](sdl/sdl_functions_js.go#L2886) |
| [SDL_SetAudioStreamOutputChannelMap](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamOutputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3827) | [:x:](sdl/sdl_functions_js.go#L2909) |
| [SDL_PutAudioStreamData](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData) | [:heavy_check_mark:](sdl/methods.go#L3837) | [:x:](sdl/sdl_functions_js.go#L2932) |
| [SDL_PutAudioStreamPlanarData](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamPlanarData) | [:question:]() | [:question:]() |
| [SDL_GetAudioStreamData](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData) | [:heavy_check_mark:](sdl/methods.go#L3847) | [:x:](sdl/sdl_functions_js.go#L2952) |
| [SDL_GetAudioStreamAvailable](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamAvailable) | [:heavy_check_mark:](sdl/methods.go#L3858) | [:x:](sdl/sdl_functions_js.go#L2972) |
| [SDL_GetAudioStreamQueued](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamQueued) | [:heavy_check_mark:](sdl/methods.go#L3869) | [:x:](sdl/sdl_functions_js.go#L2988) |
| [SDL_FlushAudioStream](https://wiki.libsdl.org/SDL3/SDL_FlushAudioStream) | [:heavy_check_mark:](sdl/methods.go#L3880) | [:x:](sdl/sdl_functions_js.go#L3004) |
| [SDL_ClearAudioStream](https://wiki.libsdl.org/SDL3/SDL_ClearAudioStream) | [:heavy_check_mark:](sdl/methods.go#L3890) | [:x:](sdl/sdl_functions_js.go#L3020) |
| [SDL_PauseAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_PauseAudioStreamDevice) | [:heavy_check_mark:](sdl/methods.go#L3900) | [:x:](sdl/sdl_functions_js.go#L3036) |
| [SDL_ResumeAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_ResumeAudioStreamDevice) | [:heavy_check_mark:](sdl/methods.go#L3910) | [:x:](sdl/sdl_functions_js.go#L3052) |
| [SDL_AudioStreamDevicePaused](https://wiki.libsdl.org/SDL3/SDL_AudioStreamDevicePaused) | [:heavy_check_mark:](sdl/methods.go#L3920) | [:x:](sdl/sdl_functions_js.go#L3068) |
| [SDL_LockAudioStream](https://wiki.libsdl.org/SDL3/SDL_LockAudioStream) | [:heavy_check_mark:](sdl/methods.go#L3926) | [:x:](sdl/sdl_functions_js.go#L3084) |
| [SDL_UnlockAudioStream](https://wiki.libsdl.org/SDL3/SDL_UnlockAudioStream) | [:heavy_check_mark:](sdl/methods.go#L3936) | [:x:](sdl/sdl_functions_js.go#L3100) |
| [SDL_SetAudioStreamGetCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGetCallback) | [:x:](sdl/methods.go#L3946) | [:x:](sdl/sdl_functions_js.go#L3116) |
| [SDL_SetAudioStreamPutCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamPutCallback) | [:x:](sdl/methods.go#L3953) | [:x:](sdl/sdl_functions_js.go#L3136) |
| [SDL_DestroyAudioStream](https://wiki.libsdl.org/SDL3/SDL_DestroyAudioStream) | [:heavy_check_mark:](sdl/methods.go#L3960) | [:x:](sdl/sdl_functions_js.go#L3156) |
| [SDL_OpenAudioDeviceStream](https://wiki.libsdl.org/SDL3/SDL_OpenAudioDeviceStream) | [:x:](sdl/methods.go#L360) | [:x:](sdl/sdl_functions_js.go#L3170) |
| [SDL_SetAudioIterationCallbacks](https://wiki.libsdl.org/SDL3/SDL_SetAudioIterationCallbacks) | [:question:]() | [:question:]() |
| [SDL_SetAudioPostmixCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioPostmixCallback) | [:x:](sdl/methods.go#L367) | [:x:](sdl/sdl_functions_js.go#L3194) |
| [SDL_LoadWAV_IO](https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO) | [:heavy_check_mark:](sdl/functions.go#L844) | [:x:](sdl/sdl_functions_js.go#L3211) |
| [SDL_LoadWAV](https://wiki.libsdl.org/SDL3/SDL_LoadWAV) | [:heavy_check_mark:](sdl/functions.go#L858) | [:x:](sdl/sdl_functions_js.go#L3244) |
| [SDL_MixAudio](https://wiki.libsdl.org/SDL3/SDL_MixAudio) | [:heavy_check_mark:](sdl/functions.go#L872) | [:x:](sdl/sdl_functions_js.go#L3272) |
| [SDL_ConvertAudioSamples](https://wiki.libsdl.org/SDL3/SDL_ConvertAudioSamples) | [:heavy_check_mark:](sdl/functions.go#L883) | [:x:](sdl/sdl_functions_js.go#L3299) |
| [SDL_GetAudioFormatName](https://wiki.libsdl.org/SDL3/SDL_GetAudioFormatName) | [:heavy_check_mark:](sdl/methods.go#L1203) | [:x:](sdl/sdl_functions_js.go#L3337) |
| [SDL_GetSilenceValueForFormat](https://wiki.libsdl.org/SDL3/SDL_GetSilenceValueForFormat) | [:heavy_check_mark:](sdl/methods.go#L1209) | [:x:](sdl/sdl_functions_js.go#L3350) |
### Time

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetDateTimeLocalePreferences](https://wiki.libsdl.org/SDL3/SDL_GetDateTimeLocalePreferences) | [:x:](sdl/methods.go#L3458) | [:x:](sdl/sdl_functions_js.go#L16962) |
| [SDL_GetCurrentTime](https://wiki.libsdl.org/SDL3/SDL_GetCurrentTime) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16983) |
| [SDL_TimeToDateTime](https://wiki.libsdl.org/SDL3/SDL_TimeToDateTime) | [:x:](sdl/methods.go#L5650) | [:x:](sdl/sdl_functions_js.go#L16999) |
| [SDL_DateTimeToTime](https://wiki.libsdl.org/SDL3/SDL_DateTimeToTime) | [:heavy_check_mark:](sdl/methods.go#L5898) | [:x:](sdl/sdl_functions_js.go#L17019) |
| [SDL_TimeToWindows](https://wiki.libsdl.org/SDL3/SDL_TimeToWindows) | [:x:](sdl/methods.go#L5657) | [:x:](sdl/sdl_functions_js.go#L17040) |
| [SDL_TimeFromWindows](https://wiki.libsdl.org/SDL3/SDL_TimeFromWindows) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L17061) |
| [SDL_GetDaysInMonth](https://wiki.libsdl.org/SDL3/SDL_GetDaysInMonth) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L17076) |
| [SDL_GetDayOfYear](https://wiki.libsdl.org/SDL3/SDL_GetDayOfYear) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L17091) |
| [SDL_GetDayOfWeek](https://wiki.libsdl.org/SDL3/SDL_GetDayOfWeek) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L17108) |
### Timer

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetTicks](https://wiki.libsdl.org/SDL3/SDL_GetTicks) | [:heavy_check_mark:](sdl/functions.go#L905) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L17125) |
| [SDL_GetTicksNS](https://wiki.libsdl.org/SDL3/SDL_GetTicksNS) | [:heavy_check_mark:](sdl/functions.go#L911) | [:x:](sdl/sdl_functions_js.go#L17133) |
| [SDL_GetPerformanceCounter](https://wiki.libsdl.org/SDL3/SDL_GetPerformanceCounter) | [:heavy_check_mark:](sdl/functions.go#L917) | [:x:](sdl/sdl_functions_js.go#L17144) |
| [SDL_GetPerformanceFrequency](https://wiki.libsdl.org/SDL3/SDL_GetPerformanceFrequency) | [:heavy_check_mark:](sdl/functions.go#L923) | [:x:](sdl/sdl_functions_js.go#L17155) |
| [SDL_Delay](https://wiki.libsdl.org/SDL3/SDL_Delay) | [:heavy_check_mark:](sdl/functions.go#L929) | [:x:](sdl/sdl_functions_js.go#L17166) |
| [SDL_DelayNS](https://wiki.libsdl.org/SDL3/SDL_DelayNS) | [:heavy_check_mark:](sdl/functions.go#L935) | [:x:](sdl/sdl_functions_js.go#L17177) |
| [SDL_DelayPrecise](https://wiki.libsdl.org/SDL3/SDL_DelayPrecise) | [:heavy_check_mark:](sdl/functions.go#L941) | [:x:](sdl/sdl_functions_js.go#L17188) |
| [SDL_AddTimer](https://wiki.libsdl.org/SDL3/SDL_AddTimer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L17199) |
| [SDL_AddTimerNS](https://wiki.libsdl.org/SDL3/SDL_AddTimerNS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L17216) |
| [SDL_RemoveTimer](https://wiki.libsdl.org/SDL3/SDL_RemoveTimer) | [:x:](sdl/methods.go#L4800) | [:x:](sdl/sdl_functions_js.go#L17233) |
### Render

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumRenderDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumRenderDrivers) | [:heavy_check_mark:](sdl/functions.go#L316) | [:x:](sdl/sdl_functions_js.go#L14586) |
| [SDL_GetRenderDriver](https://wiki.libsdl.org/SDL3/SDL_GetRenderDriver) | [:heavy_check_mark:](sdl/functions.go#L322) | [:x:](sdl/sdl_functions_js.go#L14597) |
| [SDL_CreateWindowAndRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateWindowAndRenderer) | [:heavy_check_mark:](sdl/functions.go#L328) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L14610) |
| [SDL_CreateRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateRenderer) | [:heavy_check_mark:](sdl/methods.go#L4741) | [:x:](sdl/sdl_functions_js.go#L14638) |
| [SDL_CreateRendererWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateRendererWithProperties) | [:heavy_check_mark:](sdl/functions.go#L341) | [:x:](sdl/sdl_functions_js.go#L14659) |
| [SDL_CreateGPURenderer](https://wiki.libsdl.org/SDL3/SDL_CreateGPURenderer) | [:question:]() | [:question:]() |
| [SDL_CreateSoftwareRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateSoftwareRenderer) | [:heavy_check_mark:](sdl/methods.go#L1862) | [:x:](sdl/sdl_functions_js.go#L14675) |
| [SDL_GetRenderer](https://wiki.libsdl.org/SDL3/SDL_GetRenderer) | [:heavy_check_mark:](sdl/methods.go#L4752) | [:x:](sdl/sdl_functions_js.go#L14694) |
| [SDL_GetRenderWindow](https://wiki.libsdl.org/SDL3/SDL_GetRenderWindow) | [:heavy_check_mark:](sdl/methods.go#L2854) | [:x:](sdl/sdl_functions_js.go#L14713) |
| [SDL_GetRendererName](https://wiki.libsdl.org/SDL3/SDL_GetRendererName) | [:heavy_check_mark:](sdl/methods.go#L2865) | [:x:](sdl/sdl_functions_js.go#L14732) |
| [SDL_GetRendererProperties](https://wiki.libsdl.org/SDL3/SDL_GetRendererProperties) | [:heavy_check_mark:](sdl/methods.go#L2876) | [:x:](sdl/sdl_functions_js.go#L14748) |
| [SDL_GetRenderOutputSize](https://wiki.libsdl.org/SDL3/SDL_GetRenderOutputSize) | [:heavy_check_mark:](sdl/methods.go#L2882) | [:x:](sdl/sdl_functions_js.go#L14764) |
| [SDL_GetCurrentRenderOutputSize](https://wiki.libsdl.org/SDL3/SDL_GetCurrentRenderOutputSize) | [:heavy_check_mark:](sdl/methods.go#L2893) | [:x:](sdl/sdl_functions_js.go#L14790) |
| [SDL_CreateTexture](https://wiki.libsdl.org/SDL3/SDL_CreateTexture) | [:heavy_check_mark:](sdl/methods.go#L2904) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L14816) |
| [SDL_CreateTextureFromSurface](https://wiki.libsdl.org/SDL3/SDL_CreateTextureFromSurface) | [:heavy_check_mark:](sdl/methods.go#L2915) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L14839) |
| [SDL_CreateTextureWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateTextureWithProperties) | [:heavy_check_mark:](sdl/methods.go#L2926) | [:x:](sdl/sdl_functions_js.go#L14859) |
| [SDL_GetTextureProperties](https://wiki.libsdl.org/SDL3/SDL_GetTextureProperties) | [:heavy_check_mark:](sdl/methods.go#L939) | [:x:](sdl/sdl_functions_js.go#L14880) |
| [SDL_GetRendererFromTexture](https://wiki.libsdl.org/SDL3/SDL_GetRendererFromTexture) | [:heavy_check_mark:](sdl/methods.go#L950) | [:x:](sdl/sdl_functions_js.go#L14896) |
| [SDL_GetTextureSize](https://wiki.libsdl.org/SDL3/SDL_GetTextureSize) | [:heavy_check_mark:](sdl/methods.go#L961) | [:x:](sdl/sdl_functions_js.go#L14915) |
| [SDL_SetTextureColorMod](https://wiki.libsdl.org/SDL3/SDL_SetTextureColorMod) | [:heavy_check_mark:](sdl/methods.go#L972) | [:x:](sdl/sdl_functions_js.go#L14941) |
| [SDL_SetTextureColorModFloat](https://wiki.libsdl.org/SDL3/SDL_SetTextureColorModFloat) | [:heavy_check_mark:](sdl/methods.go#L982) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L14963) |
| [SDL_GetTextureColorMod](https://wiki.libsdl.org/SDL3/SDL_GetTextureColorMod) | [:heavy_check_mark:](sdl/methods.go#L992) | [:x:](sdl/sdl_functions_js.go#L14979) |
| [SDL_GetTextureColorModFloat](https://wiki.libsdl.org/SDL3/SDL_GetTextureColorModFloat) | [:heavy_check_mark:](sdl/methods.go#L1003) | [:x:](sdl/sdl_functions_js.go#L15010) |
| [SDL_SetTextureAlphaMod](https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1014) | [:x:](sdl/sdl_functions_js.go#L15041) |
| [SDL_SetTextureAlphaModFloat](https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaModFloat) | [:heavy_check_mark:](sdl/methods.go#L1024) | [:x:](sdl/sdl_functions_js.go#L15059) |
| [SDL_GetTextureAlphaMod](https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1034) | [:x:](sdl/sdl_functions_js.go#L15077) |
| [SDL_GetTextureAlphaModFloat](https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaModFloat) | [:heavy_check_mark:](sdl/methods.go#L1045) | [:x:](sdl/sdl_functions_js.go#L15098) |
| [SDL_SetTextureBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetTextureBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1056) | [:x:](sdl/sdl_functions_js.go#L15119) |
| [SDL_GetTextureBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetTextureBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1066) | [:x:](sdl/sdl_functions_js.go#L15137) |
| [SDL_SetTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_SetTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L1077) | [:x:](sdl/sdl_functions_js.go#L15158) |
| [SDL_GetTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_GetTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L1087) | [:x:](sdl/sdl_functions_js.go#L15176) |
| [SDL_UpdateTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateTexture) | [:heavy_check_mark:](sdl/methods.go#L1098) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15197) |
| [SDL_UpdateYUVTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateYUVTexture) | [:x:](sdl/methods.go#L1110) | [:x:](sdl/sdl_functions_js.go#L15221) |
| [SDL_UpdateNVTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateNVTexture) | [:x:](sdl/methods.go#L1117) | [:x:](sdl/sdl_functions_js.go#L15263) |
| [SDL_LockTexture](https://wiki.libsdl.org/SDL3/SDL_LockTexture) | [:heavy_check_mark:](sdl/methods.go#L1125) | [:x:](sdl/sdl_functions_js.go#L15298) |
| [SDL_LockTextureToSurface](https://wiki.libsdl.org/SDL3/SDL_LockTextureToSurface) | [:heavy_check_mark:](sdl/methods.go#L1138) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15329) |
| [SDL_UnlockTexture](https://wiki.libsdl.org/SDL3/SDL_UnlockTexture) | [:heavy_check_mark:](sdl/methods.go#L1148) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15351) |
| [SDL_SetRenderTarget](https://wiki.libsdl.org/SDL3/SDL_SetRenderTarget) | [:heavy_check_mark:](sdl/methods.go#L2937) | [:x:](sdl/sdl_functions_js.go#L15362) |
| [SDL_GetRenderTarget](https://wiki.libsdl.org/SDL3/SDL_GetRenderTarget) | [:heavy_check_mark:](sdl/methods.go#L2947) | [:x:](sdl/sdl_functions_js.go#L15383) |
| [SDL_SetRenderLogicalPresentation](https://wiki.libsdl.org/SDL3/SDL_SetRenderLogicalPresentation) | [:heavy_check_mark:](sdl/methods.go#L2953) | [:x:](sdl/sdl_functions_js.go#L15402) |
| [SDL_GetRenderLogicalPresentation](https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentation) | [:heavy_check_mark:](sdl/methods.go#L2963) | [:x:](sdl/sdl_functions_js.go#L15424) |
| [SDL_GetRenderLogicalPresentationRect](https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentationRect) | [:heavy_check_mark:](sdl/methods.go#L2975) | [:x:](sdl/sdl_functions_js.go#L15455) |
| [SDL_RenderCoordinatesFromWindow](https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesFromWindow) | [:x:](sdl/methods.go#L2987) | [:x:](sdl/sdl_functions_js.go#L15476) |
| [SDL_RenderCoordinatesToWindow](https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesToWindow) | [:x:](sdl/methods.go#L2994) | [:x:](sdl/sdl_functions_js.go#L15506) |
| [SDL_ConvertEventToRenderCoordinates](https://wiki.libsdl.org/SDL3/SDL_ConvertEventToRenderCoordinates) | [:x:](sdl/methods.go#L3001) | [:x:](sdl/sdl_functions_js.go#L15536) |
| [SDL_SetRenderViewport](https://wiki.libsdl.org/SDL3/SDL_SetRenderViewport) | [:heavy_check_mark:](sdl/methods.go#L3008) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15557) |
| [SDL_GetRenderViewport](https://wiki.libsdl.org/SDL3/SDL_GetRenderViewport) | [:heavy_check_mark:](sdl/methods.go#L3018) | [:x:](sdl/sdl_functions_js.go#L15574) |
| [SDL_RenderViewportSet](https://wiki.libsdl.org/SDL3/SDL_RenderViewportSet) | [:heavy_check_mark:](sdl/methods.go#L3030) | [:x:](sdl/sdl_functions_js.go#L15595) |
| [SDL_GetRenderSafeArea](https://wiki.libsdl.org/SDL3/SDL_GetRenderSafeArea) | [:heavy_check_mark:](sdl/methods.go#L3036) | [:x:](sdl/sdl_functions_js.go#L15611) |
| [SDL_SetRenderClipRect](https://wiki.libsdl.org/SDL3/SDL_SetRenderClipRect) | [:heavy_check_mark:](sdl/methods.go#L3048) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15632) |
| [SDL_GetRenderClipRect](https://wiki.libsdl.org/SDL3/SDL_GetRenderClipRect) | [:heavy_check_mark:](sdl/methods.go#L3058) | [:x:](sdl/sdl_functions_js.go#L15649) |
| [SDL_RenderClipEnabled](https://wiki.libsdl.org/SDL3/SDL_RenderClipEnabled) | [:heavy_check_mark:](sdl/methods.go#L3070) | [:x:](sdl/sdl_functions_js.go#L15670) |
| [SDL_SetRenderScale](https://wiki.libsdl.org/SDL3/SDL_SetRenderScale) | [:heavy_check_mark:](sdl/methods.go#L3080) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15686) |
| [SDL_GetRenderScale](https://wiki.libsdl.org/SDL3/SDL_GetRenderScale) | [:heavy_check_mark:](sdl/methods.go#L3090) | [:x:](sdl/sdl_functions_js.go#L15703) |
| [SDL_SetRenderDrawColor](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColor) | [:heavy_check_mark:](sdl/methods.go#L3102) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15729) |
| [SDL_SetRenderDrawColorFloat](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColorFloat) | [:heavy_check_mark:](sdl/methods.go#L3112) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15751) |
| [SDL_GetRenderDrawColor](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColor) | [:heavy_check_mark:](sdl/methods.go#L3122) | [:x:](sdl/sdl_functions_js.go#L15768) |
| [SDL_GetRenderDrawColorFloat](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColorFloat) | [:heavy_check_mark:](sdl/methods.go#L3134) | [:x:](sdl/sdl_functions_js.go#L15804) |
| [SDL_SetRenderColorScale](https://wiki.libsdl.org/SDL3/SDL_SetRenderColorScale) | [:heavy_check_mark:](sdl/methods.go#L3146) | [:x:](sdl/sdl_functions_js.go#L15840) |
| [SDL_GetRenderColorScale](https://wiki.libsdl.org/SDL3/SDL_GetRenderColorScale) | [:heavy_check_mark:](sdl/methods.go#L3156) | [:x:](sdl/sdl_functions_js.go#L15858) |
| [SDL_SetRenderDrawBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawBlendMode) | [:heavy_check_mark:](sdl/methods.go#L3168) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15879) |
| [SDL_GetRenderDrawBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawBlendMode) | [:heavy_check_mark:](sdl/methods.go#L3178) | [:x:](sdl/sdl_functions_js.go#L15894) |
| [SDL_RenderClear](https://wiki.libsdl.org/SDL3/SDL_RenderClear) | [:heavy_check_mark:](sdl/methods.go#L3190) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15915) |
| [SDL_RenderPoint](https://wiki.libsdl.org/SDL3/SDL_RenderPoint) | [:heavy_check_mark:](sdl/methods.go#L3199) | [:x:](sdl/sdl_functions_js.go#L15928) |
| [SDL_RenderPoints](https://wiki.libsdl.org/SDL3/SDL_RenderPoints) | [:heavy_check_mark:](sdl/methods.go#L3209) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15948) |
| [SDL_RenderLine](https://wiki.libsdl.org/SDL3/SDL_RenderLine) | [:heavy_check_mark:](sdl/methods.go#L3219) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15966) |
| [SDL_RenderLines](https://wiki.libsdl.org/SDL3/SDL_RenderLines) | [:heavy_check_mark:](sdl/methods.go#L3229) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15983) |
| [SDL_RenderRect](https://wiki.libsdl.org/SDL3/SDL_RenderRect) | [:heavy_check_mark:](sdl/methods.go#L3239) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16001) |
| [SDL_RenderRects](https://wiki.libsdl.org/SDL3/SDL_RenderRects) | [:heavy_check_mark:](sdl/methods.go#L3249) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16018) |
| [SDL_RenderFillRect](https://wiki.libsdl.org/SDL3/SDL_RenderFillRect) | [:heavy_check_mark:](sdl/methods.go#L3259) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16036) |
| [SDL_RenderFillRects](https://wiki.libsdl.org/SDL3/SDL_RenderFillRects) | [:heavy_check_mark:](sdl/methods.go#L3269) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16053) |
| [SDL_RenderTexture](https://wiki.libsdl.org/SDL3/SDL_RenderTexture) | [:heavy_check_mark:](sdl/methods.go#L3279) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16071) |
| [SDL_RenderTextureRotated](https://wiki.libsdl.org/SDL3/SDL_RenderTextureRotated) | [:heavy_check_mark:](sdl/methods.go#L3289) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16095) |
| [SDL_RenderTextureAffine](https://wiki.libsdl.org/SDL3/SDL_RenderTextureAffine) | [:heavy_check_mark:](sdl/methods.go#L3299) | [:x:](sdl/sdl_functions_js.go#L16125) |
| [SDL_RenderTextureTiled](https://wiki.libsdl.org/SDL3/SDL_RenderTextureTiled) | [:heavy_check_mark:](sdl/methods.go#L3309) | [:x:](sdl/sdl_functions_js.go#L16166) |
| [SDL_RenderTexture9Grid](https://wiki.libsdl.org/SDL3/SDL_RenderTexture9Grid) | [:heavy_check_mark:](sdl/methods.go#L3319) | [:x:](sdl/sdl_functions_js.go#L16199) |
| [SDL_RenderTexture9GridTiled](https://wiki.libsdl.org/SDL3/SDL_RenderTexture9GridTiled) | [:question:]() | [:question:]() |
| [SDL_RenderGeometry](https://wiki.libsdl.org/SDL3/SDL_RenderGeometry) | [:heavy_check_mark:](sdl/methods.go#L3329) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16240) |
| [SDL_RenderGeometryRaw](https://wiki.libsdl.org/SDL3/SDL_RenderGeometryRaw) | [:heavy_check_mark:](sdl/methods.go#L3339) | [:x:](sdl/sdl_functions_js.go#L16267) |
| [SDL_SetRenderTextureAddressMode](https://wiki.libsdl.org/SDL3/SDL_SetRenderTextureAddressMode) | [:question:]() | [:question:]() |
| [SDL_GetRenderTextureAddressMode](https://wiki.libsdl.org/SDL3/SDL_GetRenderTextureAddressMode) | [:question:]() | [:question:]() |
| [SDL_RenderReadPixels](https://wiki.libsdl.org/SDL3/SDL_RenderReadPixels) | [:heavy_check_mark:](sdl/methods.go#L3360) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16317) |
| [SDL_RenderPresent](https://wiki.libsdl.org/SDL3/SDL_RenderPresent) | [:heavy_check_mark:](sdl/methods.go#L3371) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16336) |
| [SDL_DestroyTexture](https://wiki.libsdl.org/SDL3/SDL_DestroyTexture) | [:heavy_check_mark:](sdl/methods.go#L1154) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16349) |
| [SDL_DestroyRenderer](https://wiki.libsdl.org/SDL3/SDL_DestroyRenderer) | [:heavy_check_mark:](sdl/methods.go#L3381) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16361) |
| [SDL_FlushRenderer](https://wiki.libsdl.org/SDL3/SDL_FlushRenderer) | [:heavy_check_mark:](sdl/methods.go#L3387) | [:x:](sdl/sdl_functions_js.go#L16373) |
| [SDL_GetRenderMetalLayer](https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalLayer) | [:x:](sdl/methods.go#L3397) | [:x:](sdl/sdl_functions_js.go#L16389) |
| [SDL_GetRenderMetalCommandEncoder](https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalCommandEncoder) | [:x:](sdl/methods.go#L3404) | [:x:](sdl/sdl_functions_js.go#L16405) |
| [SDL_AddVulkanRenderSemaphores](https://wiki.libsdl.org/SDL3/SDL_AddVulkanRenderSemaphores) | [:x:](sdl/methods.go#L3411) | [:x:](sdl/sdl_functions_js.go#L16421) |
| [SDL_SetRenderVSync](https://wiki.libsdl.org/SDL3/SDL_SetRenderVSync) | [:heavy_check_mark:](sdl/methods.go#L3418) | [:x:](sdl/sdl_functions_js.go#L16443) |
| [SDL_GetRenderVSync](https://wiki.libsdl.org/SDL3/SDL_GetRenderVSync) | [:heavy_check_mark:](sdl/methods.go#L3428) | [:x:](sdl/sdl_functions_js.go#L16461) |
| [SDL_RenderDebugText](https://wiki.libsdl.org/SDL3/SDL_RenderDebugText) | [:heavy_check_mark:](sdl/methods.go#L3450) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16482) |
| [SDL_RenderDebugTextFormat](https://wiki.libsdl.org/SDL3/SDL_RenderDebugTextFormat) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16502) |
| [SDL_SetDefaultTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_SetDefaultTextureScaleMode) | [:question:]() | [:question:]() |
| [SDL_GetDefaultTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_GetDefaultTextureScaleMode) | [:question:]() | [:question:]() |
| [SDL_CreateGPURenderState](https://wiki.libsdl.org/SDL3/SDL_CreateGPURenderState) | [:question:]() | [:question:]() |
| [SDL_SetGPURenderStateFragmentUniforms](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderStateFragmentUniforms) | [:question:]() | [:question:]() |
| [SDL_SetRenderGPUState](https://wiki.libsdl.org/SDL3/SDL_SetRenderGPUState) | [:question:]() | [:question:]() |
| [SDL_DestroyGPURenderState](https://wiki.libsdl.org/SDL3/SDL_DestroyGPURenderState) | [:question:]() | [:question:]() |
### SharedObject

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_LoadObject](https://wiki.libsdl.org/SDL3/SDL_LoadObject) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14023) |
| [SDL_LoadFunction](https://wiki.libsdl.org/SDL3/SDL_LoadFunction) | [:x:](sdl/methods.go#L4784) | [:x:](sdl/sdl_functions_js.go#L14039) |
| [SDL_UnloadObject](https://wiki.libsdl.org/SDL3/SDL_UnloadObject) | [:x:](sdl/methods.go#L4791) | [:x:](sdl/sdl_functions_js.go#L14057) |
### Thread

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateThread](https://wiki.libsdl.org/SDL3/SDL_CreateThread) | [:question:]() | [:question:]() |
| [SDL_CreateThreadWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateThreadWithProperties) | [:question:]() | [:question:]() |
| [SDL_GetThreadName](https://wiki.libsdl.org/SDL3/SDL_GetThreadName) | [:x:](sdl/methods.go#L485) | [:x:](sdl/sdl_functions_js.go#L871) |
| [SDL_GetCurrentThreadID](https://wiki.libsdl.org/SDL3/SDL_GetCurrentThreadID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L887) |
| [SDL_GetThreadID](https://wiki.libsdl.org/SDL3/SDL_GetThreadID) | [:x:](sdl/methods.go#L492) | [:x:](sdl/sdl_functions_js.go#L898) |
| [SDL_SetCurrentThreadPriority](https://wiki.libsdl.org/SDL3/SDL_SetCurrentThreadPriority) | [:x:](sdl/methods.go#L23) | [:x:](sdl/sdl_functions_js.go#L914) |
| [SDL_WaitThread](https://wiki.libsdl.org/SDL3/SDL_WaitThread) | [:x:](sdl/methods.go#L499) | [:x:](sdl/sdl_functions_js.go#L927) |
| [SDL_GetThreadState](https://wiki.libsdl.org/SDL3/SDL_GetThreadState) | [:x:](sdl/methods.go#L506) | [:x:](sdl/sdl_functions_js.go#L946) |
| [SDL_DetachThread](https://wiki.libsdl.org/SDL3/SDL_DetachThread) | [:x:](sdl/methods.go#L513) | [:x:](sdl/sdl_functions_js.go#L962) |
| [SDL_GetTLS](https://wiki.libsdl.org/SDL3/SDL_GetTLS) | [:x:](sdl/methods.go#L610) | [:x:](sdl/sdl_functions_js.go#L976) |
| [SDL_SetTLS](https://wiki.libsdl.org/SDL3/SDL_SetTLS) | [:x:](sdl/methods.go#L617) | [:x:](sdl/sdl_functions_js.go#L992) |
| [SDL_CleanupTLS](https://wiki.libsdl.org/SDL3/SDL_CleanupTLS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1012) |
### Mutex

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateMutex](https://wiki.libsdl.org/SDL3/SDL_CreateMutex) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1021) |
| [SDL_LockMutex](https://wiki.libsdl.org/SDL3/SDL_LockMutex) | [:heavy_check_mark:](sdl/methods.go#L5624) | [:x:](sdl/sdl_functions_js.go#L1035) |
| [SDL_TryLockMutex](https://wiki.libsdl.org/SDL3/SDL_TryLockMutex) | [:heavy_check_mark:](sdl/methods.go#L5630) | [:x:](sdl/sdl_functions_js.go#L1049) |
| [SDL_UnlockMutex](https://wiki.libsdl.org/SDL3/SDL_UnlockMutex) | [:heavy_check_mark:](sdl/methods.go#L5636) | [:x:](sdl/sdl_functions_js.go#L1065) |
| [SDL_DestroyMutex](https://wiki.libsdl.org/SDL3/SDL_DestroyMutex) | [:heavy_check_mark:](sdl/methods.go#L5642) | [:x:](sdl/sdl_functions_js.go#L1079) |
| [SDL_CreateRWLock](https://wiki.libsdl.org/SDL3/SDL_CreateRWLock) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1093) |
| [SDL_LockRWLockForReading](https://wiki.libsdl.org/SDL3/SDL_LockRWLockForReading) | [:heavy_check_mark:](sdl/methods.go#L1165) | [:x:](sdl/sdl_functions_js.go#L1107) |
| [SDL_LockRWLockForWriting](https://wiki.libsdl.org/SDL3/SDL_LockRWLockForWriting) | [:heavy_check_mark:](sdl/methods.go#L1171) | [:x:](sdl/sdl_functions_js.go#L1121) |
| [SDL_TryLockRWLockForReading](https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForReading) | [:heavy_check_mark:](sdl/methods.go#L1177) | [:x:](sdl/sdl_functions_js.go#L1135) |
| [SDL_TryLockRWLockForWriting](https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForWriting) | [:heavy_check_mark:](sdl/methods.go#L1183) | [:x:](sdl/sdl_functions_js.go#L1151) |
| [SDL_UnlockRWLock](https://wiki.libsdl.org/SDL3/SDL_UnlockRWLock) | [:heavy_check_mark:](sdl/methods.go#L1189) | [:x:](sdl/sdl_functions_js.go#L1167) |
| [SDL_DestroyRWLock](https://wiki.libsdl.org/SDL3/SDL_DestroyRWLock) | [:heavy_check_mark:](sdl/methods.go#L1195) | [:x:](sdl/sdl_functions_js.go#L1181) |
| [SDL_CreateSemaphore](https://wiki.libsdl.org/SDL3/SDL_CreateSemaphore) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1195) |
| [SDL_DestroySemaphore](https://wiki.libsdl.org/SDL3/SDL_DestroySemaphore) | [:heavy_check_mark:](sdl/methods.go#L5833) | [:x:](sdl/sdl_functions_js.go#L1211) |
| [SDL_WaitSemaphore](https://wiki.libsdl.org/SDL3/SDL_WaitSemaphore) | [:heavy_check_mark:](sdl/methods.go#L5839) | [:x:](sdl/sdl_functions_js.go#L1225) |
| [SDL_TryWaitSemaphore](https://wiki.libsdl.org/SDL3/SDL_TryWaitSemaphore) | [:heavy_check_mark:](sdl/methods.go#L5845) | [:x:](sdl/sdl_functions_js.go#L1239) |
| [SDL_WaitSemaphoreTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitSemaphoreTimeout) | [:heavy_check_mark:](sdl/methods.go#L5851) | [:x:](sdl/sdl_functions_js.go#L1255) |
| [SDL_SignalSemaphore](https://wiki.libsdl.org/SDL3/SDL_SignalSemaphore) | [:heavy_check_mark:](sdl/methods.go#L5857) | [:x:](sdl/sdl_functions_js.go#L1273) |
| [SDL_GetSemaphoreValue](https://wiki.libsdl.org/SDL3/SDL_GetSemaphoreValue) | [:heavy_check_mark:](sdl/methods.go#L5863) | [:x:](sdl/sdl_functions_js.go#L1287) |
| [SDL_CreateCondition](https://wiki.libsdl.org/SDL3/SDL_CreateCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1303) |
| [SDL_DestroyCondition](https://wiki.libsdl.org/SDL3/SDL_DestroyCondition) | [:x:](sdl/methods.go#L4007) | [:x:](sdl/sdl_functions_js.go#L1317) |
| [SDL_SignalCondition](https://wiki.libsdl.org/SDL3/SDL_SignalCondition) | [:x:](sdl/methods.go#L4014) | [:x:](sdl/sdl_functions_js.go#L1331) |
| [SDL_BroadcastCondition](https://wiki.libsdl.org/SDL3/SDL_BroadcastCondition) | [:x:](sdl/methods.go#L4021) | [:x:](sdl/sdl_functions_js.go#L1345) |
| [SDL_WaitCondition](https://wiki.libsdl.org/SDL3/SDL_WaitCondition) | [:x:](sdl/methods.go#L4028) | [:x:](sdl/sdl_functions_js.go#L1359) |
| [SDL_WaitConditionTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitConditionTimeout) | [:x:](sdl/methods.go#L4035) | [:x:](sdl/sdl_functions_js.go#L1378) |
| [SDL_ShouldInit](https://wiki.libsdl.org/SDL3/SDL_ShouldInit) | [:x:](sdl/methods.go#L3497) | [:x:](sdl/sdl_functions_js.go#L1401) |
| [SDL_ShouldQuit](https://wiki.libsdl.org/SDL3/SDL_ShouldQuit) | [:x:](sdl/methods.go#L3504) | [:x:](sdl/sdl_functions_js.go#L1417) |
| [SDL_SetInitialized](https://wiki.libsdl.org/SDL3/SDL_SetInitialized) | [:x:](sdl/methods.go#L3511) | [:x:](sdl/sdl_functions_js.go#L1433) |
### Atomic

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_TryLockSpinlock](https://wiki.libsdl.org/SDL3/SDL_TryLockSpinlock) | [:heavy_check_mark:](sdl/methods.go#L465) | [:x:](sdl/sdl_functions_js.go#L246) |
| [SDL_LockSpinlock](https://wiki.libsdl.org/SDL3/SDL_LockSpinlock) | [:heavy_check_mark:](sdl/methods.go#L471) | [:x:](sdl/sdl_functions_js.go#L262) |
| [SDL_UnlockSpinlock](https://wiki.libsdl.org/SDL3/SDL_UnlockSpinlock) | [:heavy_check_mark:](sdl/methods.go#L477) | [:x:](sdl/sdl_functions_js.go#L276) |
| [SDL_MemoryBarrierReleaseFunction](https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierReleaseFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L290) |
| [SDL_MemoryBarrierAcquireFunction](https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierAcquireFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L299) |
| [SDL_CompareAndSwapAtomicInt](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L869) | [:x:](sdl/sdl_functions_js.go#L308) |
| [SDL_SetAtomicInt](https://wiki.libsdl.org/SDL3/SDL_SetAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L875) | [:x:](sdl/sdl_functions_js.go#L328) |
| [SDL_GetAtomicInt](https://wiki.libsdl.org/SDL3/SDL_GetAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L881) | [:x:](sdl/sdl_functions_js.go#L346) |
| [SDL_AddAtomicInt](https://wiki.libsdl.org/SDL3/SDL_AddAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L887) | [:x:](sdl/sdl_functions_js.go#L362) |
| [SDL_CompareAndSwapAtomicU32](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicU32) | [:x:](sdl/methods.go#L5666) | [:x:](sdl/sdl_functions_js.go#L380) |
| [SDL_SetAtomicU32](https://wiki.libsdl.org/SDL3/SDL_SetAtomicU32) | [:x:](sdl/methods.go#L5673) | [:x:](sdl/sdl_functions_js.go#L400) |
| [SDL_GetAtomicU32](https://wiki.libsdl.org/SDL3/SDL_GetAtomicU32) | [:x:](sdl/methods.go#L5680) | [:x:](sdl/sdl_functions_js.go#L418) |
| [SDL_CompareAndSwapAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicPointer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L434) |
| [SDL_SetAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_SetAtomicPointer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L454) |
| [SDL_GetAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_GetAtomicPointer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L472) |
### Filesystem

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetBasePath](https://wiki.libsdl.org/SDL3/SDL_GetBasePath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11210) |
| [SDL_GetPrefPath](https://wiki.libsdl.org/SDL3/SDL_GetPrefPath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11221) |
| [SDL_GetUserFolder](https://wiki.libsdl.org/SDL3/SDL_GetUserFolder) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11236) |
| [SDL_CreateDirectory](https://wiki.libsdl.org/SDL3/SDL_CreateDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11249) |
| [SDL_EnumerateDirectory](https://wiki.libsdl.org/SDL3/SDL_EnumerateDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11262) |
| [SDL_RemovePath](https://wiki.libsdl.org/SDL3/SDL_RemovePath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11279) |
| [SDL_RenamePath](https://wiki.libsdl.org/SDL3/SDL_RenamePath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11292) |
| [SDL_CopyFile](https://wiki.libsdl.org/SDL3/SDL_CopyFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11307) |
| [SDL_GetPathInfo](https://wiki.libsdl.org/SDL3/SDL_GetPathInfo) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11322) |
| [SDL_GlobDirectory](https://wiki.libsdl.org/SDL3/SDL_GlobDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11340) |
| [SDL_GetCurrentDirectory](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11362) |
### IOStream

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_IOFromFile](https://wiki.libsdl.org/SDL3/SDL_IOFromFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1449) |
| [SDL_IOFromMem](https://wiki.libsdl.org/SDL3/SDL_IOFromMem) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1467) |
| [SDL_IOFromConstMem](https://wiki.libsdl.org/SDL3/SDL_IOFromConstMem) | [:heavy_check_mark:](sdl/functions.go#L288) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1485) |
| [SDL_IOFromDynamicMem](https://wiki.libsdl.org/SDL3/SDL_IOFromDynamicMem) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1505) |
| [SDL_OpenIO](https://wiki.libsdl.org/SDL3/SDL_OpenIO) | [:x:](sdl/methods.go#L32) | [:x:](sdl/sdl_functions_js.go#L1519) |
| [SDL_CloseIO](https://wiki.libsdl.org/SDL3/SDL_CloseIO) | [:heavy_check_mark:](sdl/methods.go#L4809) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1540) |
| [SDL_GetIOProperties](https://wiki.libsdl.org/SDL3/SDL_GetIOProperties) | [:heavy_check_mark:](sdl/methods.go#L4819) | [:x:](sdl/sdl_functions_js.go#L1553) |
| [SDL_GetIOStatus](https://wiki.libsdl.org/SDL3/SDL_GetIOStatus) | [:heavy_check_mark:](sdl/methods.go#L4830) | [:x:](sdl/sdl_functions_js.go#L1569) |
| [SDL_GetIOSize](https://wiki.libsdl.org/SDL3/SDL_GetIOSize) | [:heavy_check_mark:](sdl/methods.go#L4836) | [:x:](sdl/sdl_functions_js.go#L1585) |
| [SDL_SeekIO](https://wiki.libsdl.org/SDL3/SDL_SeekIO) | [:heavy_check_mark:](sdl/methods.go#L4847) | [:x:](sdl/sdl_functions_js.go#L1601) |
| [SDL_TellIO](https://wiki.libsdl.org/SDL3/SDL_TellIO) | [:heavy_check_mark:](sdl/methods.go#L4858) | [:x:](sdl/sdl_functions_js.go#L1621) |
| [SDL_ReadIO](https://wiki.libsdl.org/SDL3/SDL_ReadIO) | [:heavy_check_mark:](sdl/methods.go#L4864) | [:x:](sdl/sdl_functions_js.go#L1637) |
| [SDL_WriteIO](https://wiki.libsdl.org/SDL3/SDL_WriteIO) | [:heavy_check_mark:](sdl/methods.go#L4875) | [:x:](sdl/sdl_functions_js.go#L1657) |
| [SDL_IOprintf](https://wiki.libsdl.org/SDL3/SDL_IOprintf) | [:heavy_check_mark:](sdl/methods.go#L4886) | [:x:](sdl/sdl_functions_js.go#L1677) |
| [SDL_IOvprintf](https://wiki.libsdl.org/SDL3/SDL_IOvprintf) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1695) |
| [SDL_FlushIO](https://wiki.libsdl.org/SDL3/SDL_FlushIO) | [:heavy_check_mark:](sdl/methods.go#L4897) | [:x:](sdl/sdl_functions_js.go#L1715) |
| [SDL_LoadFile_IO](https://wiki.libsdl.org/SDL3/SDL_LoadFile_IO) | [:heavy_check_mark:](sdl/methods.go#L4907) | [:x:](sdl/sdl_functions_js.go#L1731) |
| [SDL_LoadFile](https://wiki.libsdl.org/SDL3/SDL_LoadFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1754) |
| [SDL_SaveFile_IO](https://wiki.libsdl.org/SDL3/SDL_SaveFile_IO) | [:heavy_check_mark:](sdl/methods.go#L4920) | [:x:](sdl/sdl_functions_js.go#L1772) |
| [SDL_SaveFile](https://wiki.libsdl.org/SDL3/SDL_SaveFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1794) |
| [SDL_ReadU8](https://wiki.libsdl.org/SDL3/SDL_ReadU8) | [:heavy_check_mark:](sdl/methods.go#L4931) | [:x:](sdl/sdl_functions_js.go#L1811) |
| [SDL_ReadS8](https://wiki.libsdl.org/SDL3/SDL_ReadS8) | [:heavy_check_mark:](sdl/methods.go#L4943) | [:x:](sdl/sdl_functions_js.go#L1832) |
| [SDL_ReadU16LE](https://wiki.libsdl.org/SDL3/SDL_ReadU16LE) | [:heavy_check_mark:](sdl/methods.go#L4955) | [:x:](sdl/sdl_functions_js.go#L1853) |
| [SDL_ReadS16LE](https://wiki.libsdl.org/SDL3/SDL_ReadS16LE) | [:heavy_check_mark:](sdl/methods.go#L4967) | [:x:](sdl/sdl_functions_js.go#L1874) |
| [SDL_ReadU16BE](https://wiki.libsdl.org/SDL3/SDL_ReadU16BE) | [:heavy_check_mark:](sdl/methods.go#L4979) | [:x:](sdl/sdl_functions_js.go#L1895) |
| [SDL_ReadS16BE](https://wiki.libsdl.org/SDL3/SDL_ReadS16BE) | [:heavy_check_mark:](sdl/methods.go#L4991) | [:x:](sdl/sdl_functions_js.go#L1916) |
| [SDL_ReadU32LE](https://wiki.libsdl.org/SDL3/SDL_ReadU32LE) | [:heavy_check_mark:](sdl/methods.go#L5003) | [:x:](sdl/sdl_functions_js.go#L1937) |
| [SDL_ReadS32LE](https://wiki.libsdl.org/SDL3/SDL_ReadS32LE) | [:heavy_check_mark:](sdl/methods.go#L5015) | [:x:](sdl/sdl_functions_js.go#L1958) |
| [SDL_ReadU32BE](https://wiki.libsdl.org/SDL3/SDL_ReadU32BE) | [:heavy_check_mark:](sdl/methods.go#L5027) | [:x:](sdl/sdl_functions_js.go#L1979) |
| [SDL_ReadS32BE](https://wiki.libsdl.org/SDL3/SDL_ReadS32BE) | [:heavy_check_mark:](sdl/methods.go#L5039) | [:x:](sdl/sdl_functions_js.go#L2000) |
| [SDL_ReadU64LE](https://wiki.libsdl.org/SDL3/SDL_ReadU64LE) | [:heavy_check_mark:](sdl/methods.go#L5051) | [:x:](sdl/sdl_functions_js.go#L2021) |
| [SDL_ReadS64LE](https://wiki.libsdl.org/SDL3/SDL_ReadS64LE) | [:heavy_check_mark:](sdl/methods.go#L5063) | [:x:](sdl/sdl_functions_js.go#L2042) |
| [SDL_ReadU64BE](https://wiki.libsdl.org/SDL3/SDL_ReadU64BE) | [:heavy_check_mark:](sdl/methods.go#L5075) | [:x:](sdl/sdl_functions_js.go#L2063) |
| [SDL_ReadS64BE](https://wiki.libsdl.org/SDL3/SDL_ReadS64BE) | [:heavy_check_mark:](sdl/methods.go#L5087) | [:x:](sdl/sdl_functions_js.go#L2084) |
| [SDL_WriteU8](https://wiki.libsdl.org/SDL3/SDL_WriteU8) | [:heavy_check_mark:](sdl/methods.go#L5099) | [:x:](sdl/sdl_functions_js.go#L2105) |
| [SDL_WriteS8](https://wiki.libsdl.org/SDL3/SDL_WriteS8) | [:heavy_check_mark:](sdl/methods.go#L5109) | [:x:](sdl/sdl_functions_js.go#L2123) |
| [SDL_WriteU16LE](https://wiki.libsdl.org/SDL3/SDL_WriteU16LE) | [:heavy_check_mark:](sdl/methods.go#L5119) | [:x:](sdl/sdl_functions_js.go#L2141) |
| [SDL_WriteS16LE](https://wiki.libsdl.org/SDL3/SDL_WriteS16LE) | [:heavy_check_mark:](sdl/methods.go#L5129) | [:x:](sdl/sdl_functions_js.go#L2159) |
| [SDL_WriteU16BE](https://wiki.libsdl.org/SDL3/SDL_WriteU16BE) | [:heavy_check_mark:](sdl/methods.go#L5139) | [:x:](sdl/sdl_functions_js.go#L2177) |
| [SDL_WriteS16BE](https://wiki.libsdl.org/SDL3/SDL_WriteS16BE) | [:heavy_check_mark:](sdl/methods.go#L5149) | [:x:](sdl/sdl_functions_js.go#L2195) |
| [SDL_WriteU32LE](https://wiki.libsdl.org/SDL3/SDL_WriteU32LE) | [:heavy_check_mark:](sdl/methods.go#L5159) | [:x:](sdl/sdl_functions_js.go#L2213) |
| [SDL_WriteS32LE](https://wiki.libsdl.org/SDL3/SDL_WriteS32LE) | [:heavy_check_mark:](sdl/methods.go#L5169) | [:x:](sdl/sdl_functions_js.go#L2231) |
| [SDL_WriteU32BE](https://wiki.libsdl.org/SDL3/SDL_WriteU32BE) | [:heavy_check_mark:](sdl/methods.go#L5179) | [:x:](sdl/sdl_functions_js.go#L2249) |
| [SDL_WriteS32BE](https://wiki.libsdl.org/SDL3/SDL_WriteS32BE) | [:heavy_check_mark:](sdl/methods.go#L5189) | [:x:](sdl/sdl_functions_js.go#L2267) |
| [SDL_WriteU64LE](https://wiki.libsdl.org/SDL3/SDL_WriteU64LE) | [:heavy_check_mark:](sdl/methods.go#L5199) | [:x:](sdl/sdl_functions_js.go#L2285) |
| [SDL_WriteS64LE](https://wiki.libsdl.org/SDL3/SDL_WriteS64LE) | [:heavy_check_mark:](sdl/methods.go#L5209) | [:x:](sdl/sdl_functions_js.go#L2303) |
| [SDL_WriteU64BE](https://wiki.libsdl.org/SDL3/SDL_WriteU64BE) | [:heavy_check_mark:](sdl/methods.go#L5219) | [:x:](sdl/sdl_functions_js.go#L2321) |
| [SDL_WriteS64BE](https://wiki.libsdl.org/SDL3/SDL_WriteS64BE) | [:heavy_check_mark:](sdl/methods.go#L5229) | [:x:](sdl/sdl_functions_js.go#L2339) |
### AsyncIO

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AsyncIOFromFile](https://wiki.libsdl.org/SDL3/SDL_AsyncIOFromFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L23) |
| [SDL_GetAsyncIOSize](https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOSize) | [:x:](sdl/methods.go#L3467) | [:x:](sdl/sdl_functions_js.go#L41) |
| [SDL_ReadAsyncIO](https://wiki.libsdl.org/SDL3/SDL_ReadAsyncIO) | [:x:](sdl/methods.go#L3474) | [:x:](sdl/sdl_functions_js.go#L57) |
| [SDL_WriteAsyncIO](https://wiki.libsdl.org/SDL3/SDL_WriteAsyncIO) | [:x:](sdl/methods.go#L3481) | [:x:](sdl/sdl_functions_js.go#L86) |
| [SDL_CloseAsyncIO](https://wiki.libsdl.org/SDL3/SDL_CloseAsyncIO) | [:x:](sdl/methods.go#L3488) | [:x:](sdl/sdl_functions_js.go#L115) |
| [SDL_CreateAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_CreateAsyncIOQueue) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L140) |
| [SDL_DestroyAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_DestroyAsyncIOQueue) | [:heavy_check_mark:](sdl/methods.go#L1321) | [:x:](sdl/sdl_functions_js.go#L154) |
| [SDL_GetAsyncIOResult](https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOResult) | [:heavy_check_mark:](sdl/methods.go#L1327) | [:x:](sdl/sdl_functions_js.go#L168) |
| [SDL_WaitAsyncIOResult](https://wiki.libsdl.org/SDL3/SDL_WaitAsyncIOResult) | [:heavy_check_mark:](sdl/methods.go#L1337) | [:x:](sdl/sdl_functions_js.go#L189) |
| [SDL_SignalAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_SignalAsyncIOQueue) | [:heavy_check_mark:](sdl/methods.go#L1347) | [:x:](sdl/sdl_functions_js.go#L212) |
| [SDL_LoadFileAsync](https://wiki.libsdl.org/SDL3/SDL_LoadFileAsync) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L226) |
### Storage

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_OpenTitleStorage](https://wiki.libsdl.org/SDL3/SDL_OpenTitleStorage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16524) |
| [SDL_OpenUserStorage](https://wiki.libsdl.org/SDL3/SDL_OpenUserStorage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16542) |
| [SDL_OpenFileStorage](https://wiki.libsdl.org/SDL3/SDL_OpenFileStorage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16562) |
| [SDL_OpenStorage](https://wiki.libsdl.org/SDL3/SDL_OpenStorage) | [:x:](sdl/methods.go#L5615) | [:x:](sdl/sdl_functions_js.go#L16578) |
| [SDL_CloseStorage](https://wiki.libsdl.org/SDL3/SDL_CloseStorage) | [:heavy_check_mark:](sdl/methods.go#L96) | [:x:](sdl/sdl_functions_js.go#L16599) |
| [SDL_StorageReady](https://wiki.libsdl.org/SDL3/SDL_StorageReady) | [:heavy_check_mark:](sdl/methods.go#L106) | [:x:](sdl/sdl_functions_js.go#L16615) |
| [SDL_GetStorageFileSize](https://wiki.libsdl.org/SDL3/SDL_GetStorageFileSize) | [:heavy_check_mark:](sdl/methods.go#L112) | [:x:](sdl/sdl_functions_js.go#L16631) |
| [SDL_ReadStorageFile](https://wiki.libsdl.org/SDL3/SDL_ReadStorageFile) | [:heavy_check_mark:](sdl/methods.go#L124) | [:x:](sdl/sdl_functions_js.go#L16654) |
| [SDL_WriteStorageFile](https://wiki.libsdl.org/SDL3/SDL_WriteStorageFile) | [:heavy_check_mark:](sdl/methods.go#L136) | [:x:](sdl/sdl_functions_js.go#L16676) |
| [SDL_CreateStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_CreateStorageDirectory) | [:heavy_check_mark:](sdl/methods.go#L146) | [:x:](sdl/sdl_functions_js.go#L16698) |
| [SDL_EnumerateStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_EnumerateStorageDirectory) | [:x:](sdl/methods.go#L156) | [:x:](sdl/sdl_functions_js.go#L16716) |
| [SDL_RemoveStoragePath](https://wiki.libsdl.org/SDL3/SDL_RemoveStoragePath) | [:heavy_check_mark:](sdl/methods.go#L163) | [:x:](sdl/sdl_functions_js.go#L16738) |
| [SDL_RenameStoragePath](https://wiki.libsdl.org/SDL3/SDL_RenameStoragePath) | [:heavy_check_mark:](sdl/methods.go#L173) | [:x:](sdl/sdl_functions_js.go#L16756) |
| [SDL_CopyStorageFile](https://wiki.libsdl.org/SDL3/SDL_CopyStorageFile) | [:heavy_check_mark:](sdl/methods.go#L183) | [:x:](sdl/sdl_functions_js.go#L16776) |
| [SDL_GetStoragePathInfo](https://wiki.libsdl.org/SDL3/SDL_GetStoragePathInfo) | [:heavy_check_mark:](sdl/methods.go#L193) | [:x:](sdl/sdl_functions_js.go#L16796) |
| [SDL_GetStorageSpaceRemaining](https://wiki.libsdl.org/SDL3/SDL_GetStorageSpaceRemaining) | [:heavy_check_mark:](sdl/methods.go#L205) | [:x:](sdl/sdl_functions_js.go#L16819) |
| [SDL_GlobStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_GlobStorageDirectory) | [:heavy_check_mark:](sdl/methods.go#L211) | [:x:](sdl/sdl_functions_js.go#L16835) |
### Pixels

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPixelFormatName](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatName) | [:heavy_check_mark:](sdl/methods.go#L5871) | [:x:](sdl/sdl_functions_js.go#L3386) |
| [SDL_GetMasksForPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetMasksForPixelFormat) | [:x:](sdl/methods.go#L5877) | [:x:](sdl/sdl_functions_js.go#L3399) |
| [SDL_GetPixelFormatForMasks](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatForMasks) | [:heavy_check_mark:](sdl/functions.go#L354) | [:x:](sdl/sdl_functions_js.go#L3437) |
| [SDL_GetPixelFormatDetails](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatDetails) | [:heavy_check_mark:](sdl/methods.go#L5885) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3458) |
| [SDL_CreatePalette](https://wiki.libsdl.org/SDL3/SDL_CreatePalette) | [:heavy_check_mark:](sdl/functions.go#L360) | [:x:](sdl/sdl_functions_js.go#L3470) |
| [SDL_SetPaletteColors](https://wiki.libsdl.org/SDL3/SDL_SetPaletteColors) | [:heavy_check_mark:](sdl/methods.go#L5241) | [:x:](sdl/sdl_functions_js.go#L3486) |
| [SDL_DestroyPalette](https://wiki.libsdl.org/SDL3/SDL_DestroyPalette) | [:heavy_check_mark:](sdl/methods.go#L5251) | [:x:](sdl/sdl_functions_js.go#L3511) |
| [SDL_MapRGB](https://wiki.libsdl.org/SDL3/SDL_MapRGB) | [:heavy_check_mark:](sdl/methods.go#L1355) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3525) |
| [SDL_MapRGBA](https://wiki.libsdl.org/SDL3/SDL_MapRGBA) | [:heavy_check_mark:](sdl/methods.go#L1361) | [:x:](sdl/sdl_functions_js.go#L3552) |
| [SDL_GetRGB](https://wiki.libsdl.org/SDL3/SDL_GetRGB) | [:heavy_check_mark:](sdl/functions.go#L383) | [:x:](sdl/sdl_functions_js.go#L3581) |
| [SDL_GetRGBA](https://wiki.libsdl.org/SDL3/SDL_GetRGBA) | [:heavy_check_mark:](sdl/functions.go#L391) | [:x:](sdl/sdl_functions_js.go#L3617) |
### Surface

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateSurface](https://wiki.libsdl.org/SDL3/SDL_CreateSurface) | [:heavy_check_mark:](sdl/functions.go#L401) | [:x:](sdl/sdl_functions_js.go#L3932) |
| [SDL_CreateSurfaceFrom](https://wiki.libsdl.org/SDL3/SDL_CreateSurfaceFrom) | [:heavy_check_mark:](sdl/functions.go#L412) | [:x:](sdl/sdl_functions_js.go#L3952) |
| [SDL_DestroySurface](https://wiki.libsdl.org/SDL3/SDL_DestroySurface) | [:heavy_check_mark:](sdl/methods.go#L1369) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3976) |
| [SDL_GetSurfaceProperties](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceProperties) | [:heavy_check_mark:](sdl/methods.go#L1378) | [:x:](sdl/sdl_functions_js.go#L3988) |
| [SDL_SetSurfaceColorspace](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorspace) | [:heavy_check_mark:](sdl/methods.go#L1389) | [:x:](sdl/sdl_functions_js.go#L4004) |
| [SDL_GetSurfaceColorspace](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorspace) | [:heavy_check_mark:](sdl/methods.go#L1399) | [:x:](sdl/sdl_functions_js.go#L4022) |
| [SDL_CreateSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_CreateSurfacePalette) | [:heavy_check_mark:](sdl/methods.go#L1405) | [:x:](sdl/sdl_functions_js.go#L4038) |
| [SDL_SetSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_SetSurfacePalette) | [:heavy_check_mark:](sdl/methods.go#L1416) | [:x:](sdl/sdl_functions_js.go#L4057) |
| [SDL_GetSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_GetSurfacePalette) | [:heavy_check_mark:](sdl/methods.go#L1426) | [:x:](sdl/sdl_functions_js.go#L4078) |
| [SDL_AddSurfaceAlternateImage](https://wiki.libsdl.org/SDL3/SDL_AddSurfaceAlternateImage) | [:heavy_check_mark:](sdl/methods.go#L1432) | [:x:](sdl/sdl_functions_js.go#L4097) |
| [SDL_SurfaceHasAlternateImages](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasAlternateImages) | [:heavy_check_mark:](sdl/methods.go#L1442) | [:x:](sdl/sdl_functions_js.go#L4118) |
| [SDL_GetSurfaceImages](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceImages) | [:heavy_check_mark:](sdl/methods.go#L1448) | [:x:](sdl/sdl_functions_js.go#L4134) |
| [SDL_RemoveSurfaceAlternateImages](https://wiki.libsdl.org/SDL3/SDL_RemoveSurfaceAlternateImages) | [:heavy_check_mark:](sdl/methods.go#L1462) | [:x:](sdl/sdl_functions_js.go#L4155) |
| [SDL_LockSurface](https://wiki.libsdl.org/SDL3/SDL_LockSurface) | [:heavy_check_mark:](sdl/methods.go#L1468) | [:x:](sdl/sdl_functions_js.go#L4169) |
| [SDL_UnlockSurface](https://wiki.libsdl.org/SDL3/SDL_UnlockSurface) | [:heavy_check_mark:](sdl/methods.go#L1478) | [:x:](sdl/sdl_functions_js.go#L4185) |
| [SDL_LoadBMP_IO](https://wiki.libsdl.org/SDL3/SDL_LoadBMP_IO) | [:heavy_check_mark:](sdl/functions.go#L424) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4199) |
| [SDL_LoadBMP](https://wiki.libsdl.org/SDL3/SDL_LoadBMP) | [:heavy_check_mark:](sdl/functions.go#L435) | [:x:](sdl/sdl_functions_js.go#L4219) |
| [SDL_SaveBMP_IO](https://wiki.libsdl.org/SDL3/SDL_SaveBMP_IO) | [:heavy_check_mark:](sdl/methods.go#L1484) | [:x:](sdl/sdl_functions_js.go#L4235) |
| [SDL_SaveBMP](https://wiki.libsdl.org/SDL3/SDL_SaveBMP) | [:heavy_check_mark:](sdl/methods.go#L1494) | [:x:](sdl/sdl_functions_js.go#L4258) |
| [SDL_SetSurfaceRLE](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceRLE) | [:heavy_check_mark:](sdl/methods.go#L1504) | [:x:](sdl/sdl_functions_js.go#L4276) |
| [SDL_SurfaceHasRLE](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasRLE) | [:heavy_check_mark:](sdl/methods.go#L1514) | [:x:](sdl/sdl_functions_js.go#L4294) |
| [SDL_SetSurfaceColorKey](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorKey) | [:heavy_check_mark:](sdl/methods.go#L1520) | [:x:](sdl/sdl_functions_js.go#L4310) |
| [SDL_SurfaceHasColorKey](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasColorKey) | [:heavy_check_mark:](sdl/methods.go#L1530) | [:x:](sdl/sdl_functions_js.go#L4330) |
| [SDL_GetSurfaceColorKey](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorKey) | [:heavy_check_mark:](sdl/methods.go#L1536) | [:x:](sdl/sdl_functions_js.go#L4346) |
| [SDL_SetSurfaceColorMod](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorMod) | [:heavy_check_mark:](sdl/methods.go#L1547) | [:x:](sdl/sdl_functions_js.go#L4367) |
| [SDL_GetSurfaceColorMod](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorMod) | [:heavy_check_mark:](sdl/methods.go#L1557) | [:x:](sdl/sdl_functions_js.go#L4389) |
| [SDL_SetSurfaceAlphaMod](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1569) | [:x:](sdl/sdl_functions_js.go#L4420) |
| [SDL_GetSurfaceAlphaMod](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1579) | [:x:](sdl/sdl_functions_js.go#L4438) |
| [SDL_SetSurfaceBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1591) | [:x:](sdl/sdl_functions_js.go#L4459) |
| [SDL_GetSurfaceBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1601) | [:x:](sdl/sdl_functions_js.go#L4477) |
| [SDL_SetSurfaceClipRect](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceClipRect) | [:heavy_check_mark:](sdl/methods.go#L1613) | [:x:](sdl/sdl_functions_js.go#L4498) |
| [SDL_GetSurfaceClipRect](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceClipRect) | [:heavy_check_mark:](sdl/methods.go#L1619) | [:x:](sdl/sdl_functions_js.go#L4519) |
| [SDL_FlipSurface](https://wiki.libsdl.org/SDL3/SDL_FlipSurface) | [:heavy_check_mark:](sdl/methods.go#L1631) | [:x:](sdl/sdl_functions_js.go#L4540) |
| [SDL_DuplicateSurface](https://wiki.libsdl.org/SDL3/SDL_DuplicateSurface) | [:heavy_check_mark:](sdl/methods.go#L1641) | [:x:](sdl/sdl_functions_js.go#L4558) |
| [SDL_ScaleSurface](https://wiki.libsdl.org/SDL3/SDL_ScaleSurface) | [:heavy_check_mark:](sdl/methods.go#L1652) | [:x:](sdl/sdl_functions_js.go#L4577) |
| [SDL_ConvertSurface](https://wiki.libsdl.org/SDL3/SDL_ConvertSurface) | [:heavy_check_mark:](sdl/methods.go#L1663) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4602) |
| [SDL_ConvertSurfaceAndColorspace](https://wiki.libsdl.org/SDL3/SDL_ConvertSurfaceAndColorspace) | [:heavy_check_mark:](sdl/methods.go#L1674) | [:x:](sdl/sdl_functions_js.go#L4619) |
| [SDL_ConvertPixels](https://wiki.libsdl.org/SDL3/SDL_ConvertPixels) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4649) |
| [SDL_ConvertPixelsAndColorspace](https://wiki.libsdl.org/SDL3/SDL_ConvertPixelsAndColorspace) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4676) |
| [SDL_PremultiplyAlpha](https://wiki.libsdl.org/SDL3/SDL_PremultiplyAlpha) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4711) |
| [SDL_PremultiplySurfaceAlpha](https://wiki.libsdl.org/SDL3/SDL_PremultiplySurfaceAlpha) | [:heavy_check_mark:](sdl/methods.go#L1685) | [:x:](sdl/sdl_functions_js.go#L4740) |
| [SDL_ClearSurface](https://wiki.libsdl.org/SDL3/SDL_ClearSurface) | [:heavy_check_mark:](sdl/methods.go#L1695) | [:x:](sdl/sdl_functions_js.go#L4758) |
| [SDL_FillSurfaceRect](https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRect) | [:heavy_check_mark:](sdl/methods.go#L1705) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4782) |
| [SDL_FillSurfaceRects](https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRects) | [:heavy_check_mark:](sdl/methods.go#L1715) | [:x:](sdl/sdl_functions_js.go#L4801) |
| [SDL_BlitSurface](https://wiki.libsdl.org/SDL3/SDL_BlitSurface) | [:heavy_check_mark:](sdl/methods.go#L1725) | [:x:](sdl/sdl_functions_js.go#L4826) |
| [SDL_BlitSurfaceUnchecked](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUnchecked) | [:heavy_check_mark:](sdl/methods.go#L1735) | [:x:](sdl/sdl_functions_js.go#L4857) |
| [SDL_BlitSurfaceScaled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceScaled) | [:heavy_check_mark:](sdl/methods.go#L1745) | [:x:](sdl/sdl_functions_js.go#L4888) |
| [SDL_BlitSurfaceUncheckedScaled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUncheckedScaled) | [:heavy_check_mark:](sdl/methods.go#L1755) | [:x:](sdl/sdl_functions_js.go#L4921) |
| [SDL_StretchSurface](https://wiki.libsdl.org/SDL3/SDL_StretchSurface) | [:question:]() | [:question:]() |
| [SDL_BlitSurfaceTiled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiled) | [:heavy_check_mark:](sdl/methods.go#L1765) | [:x:](sdl/sdl_functions_js.go#L4954) |
| [SDL_BlitSurfaceTiledWithScale](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiledWithScale) | [:heavy_check_mark:](sdl/methods.go#L1775) | [:x:](sdl/sdl_functions_js.go#L4985) |
| [SDL_BlitSurface9Grid](https://wiki.libsdl.org/SDL3/SDL_BlitSurface9Grid) | [:heavy_check_mark:](sdl/methods.go#L1785) | [:x:](sdl/sdl_functions_js.go#L5020) |
| [SDL_MapSurfaceRGB](https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGB) | [:heavy_check_mark:](sdl/methods.go#L1795) | [:x:](sdl/sdl_functions_js.go#L5063) |
| [SDL_MapSurfaceRGBA](https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGBA) | [:heavy_check_mark:](sdl/methods.go#L1801) | [:x:](sdl/sdl_functions_js.go#L5085) |
| [SDL_ReadSurfacePixel](https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixel) | [:heavy_check_mark:](sdl/methods.go#L1807) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5109) |
| [SDL_ReadSurfacePixelFloat](https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixelFloat) | [:heavy_check_mark:](sdl/methods.go#L1819) | [:x:](sdl/sdl_functions_js.go#L5140) |
| [SDL_WriteSurfacePixel](https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixel) | [:heavy_check_mark:](sdl/methods.go#L1831) | [:x:](sdl/sdl_functions_js.go#L5180) |
| [SDL_WriteSurfacePixelFloat](https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixelFloat) | [:heavy_check_mark:](sdl/methods.go#L1841) | [:x:](sdl/sdl_functions_js.go#L5208) |
### BlendMode

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ComposeCustomBlendMode](https://wiki.libsdl.org/SDL3/SDL_ComposeCustomBlendMode) | [:heavy_check_mark:](sdl/functions.go#L452) | [:x:](sdl/sdl_functions_js.go#L3363) |
### Rect

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_RectToFRect](https://wiki.libsdl.org/SDL3/SDL_RectToFRect) | [:question:]() | [:question:]() |
| [SDL_PointInRect](https://wiki.libsdl.org/SDL3/SDL_PointInRect) | [:question:]() | [:question:]() |
| [SDL_RectEmpty](https://wiki.libsdl.org/SDL3/SDL_RectEmpty) | [:question:]() | [:question:]() |
| [SDL_RectsEqual](https://wiki.libsdl.org/SDL3/SDL_RectsEqual) | [:question:]() | [:question:]() |
| [SDL_HasRectIntersection](https://wiki.libsdl.org/SDL3/SDL_HasRectIntersection) | [:heavy_check_mark:](sdl/methods.go#L626) | [:x:](sdl/sdl_functions_js.go#L3658) |
| [SDL_GetRectIntersection](https://wiki.libsdl.org/SDL3/SDL_GetRectIntersection) | [:heavy_check_mark:](sdl/methods.go#L632) | [:x:](sdl/sdl_functions_js.go#L3679) |
| [SDL_GetRectUnion](https://wiki.libsdl.org/SDL3/SDL_GetRectUnion) | [:heavy_check_mark:](sdl/methods.go#L644) | [:x:](sdl/sdl_functions_js.go#L3705) |
| [SDL_GetRectEnclosingPoints](https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPoints) | [:heavy_check_mark:](sdl/functions.go#L1540) | [:x:](sdl/sdl_functions_js.go#L3731) |
| [SDL_GetRectAndLineIntersection](https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersection) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3759) |
| [SDL_PointInRectFloat](https://wiki.libsdl.org/SDL3/SDL_PointInRectFloat) | [:question:]() | [:question:]() |
| [SDL_RectEmptyFloat](https://wiki.libsdl.org/SDL3/SDL_RectEmptyFloat) | [:question:]() | [:question:]() |
| [SDL_RectsEqualEpsilon](https://wiki.libsdl.org/SDL3/SDL_RectsEqualEpsilon) | [:question:]() | [:question:]() |
| [SDL_RectsEqualFloat](https://wiki.libsdl.org/SDL3/SDL_RectsEqualFloat) | [:question:]() | [:question:]() |
| [SDL_HasRectIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_HasRectIntersectionFloat) | [:x:](sdl/methods.go#L3968) | [:x:](sdl/sdl_functions_js.go#L3795) |
| [SDL_GetRectIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectIntersectionFloat) | [:x:](sdl/methods.go#L3975) | [:x:](sdl/sdl_functions_js.go#L3816) |
| [SDL_GetRectUnionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectUnionFloat) | [:x:](sdl/methods.go#L3982) | [:x:](sdl/sdl_functions_js.go#L3842) |
| [SDL_GetRectEnclosingPointsFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPointsFloat) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3868) |
| [SDL_GetRectAndLineIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersectionFloat) | [:x:](sdl/methods.go#L3989) | [:x:](sdl/sdl_functions_js.go#L3896) |
### Camera

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumCameraDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumCameraDrivers) | [:heavy_check_mark:](sdl/functions.go#L1552) | [:x:](sdl/sdl_functions_js.go#L5236) |
| [SDL_GetCameraDriver](https://wiki.libsdl.org/SDL3/SDL_GetCameraDriver) | [:heavy_check_mark:](sdl/functions.go#L1558) | [:x:](sdl/sdl_functions_js.go#L5247) |
| [SDL_GetCurrentCameraDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentCameraDriver) | [:heavy_check_mark:](sdl/functions.go#L1564) | [:x:](sdl/sdl_functions_js.go#L5260) |
| [SDL_GetCameras](https://wiki.libsdl.org/SDL3/SDL_GetCameras) | [:heavy_check_mark:](sdl/functions.go#L1570) | [:x:](sdl/sdl_functions_js.go#L5271) |
| [SDL_GetCameraSupportedFormats](https://wiki.libsdl.org/SDL3/SDL_GetCameraSupportedFormats) | [:x:](sdl/methods.go#L3520) | [:x:](sdl/sdl_functions_js.go#L5287) |
| [SDL_GetCameraName](https://wiki.libsdl.org/SDL3/SDL_GetCameraName) | [:x:](sdl/methods.go#L3527) | [:x:](sdl/sdl_functions_js.go#L5305) |
| [SDL_GetCameraPosition](https://wiki.libsdl.org/SDL3/SDL_GetCameraPosition) | [:x:](sdl/methods.go#L3534) | [:x:](sdl/sdl_functions_js.go#L5318) |
| [SDL_OpenCamera](https://wiki.libsdl.org/SDL3/SDL_OpenCamera) | [:x:](sdl/methods.go#L3541) | [:x:](sdl/sdl_functions_js.go#L5331) |
| [SDL_GetCameraPermissionState](https://wiki.libsdl.org/SDL3/SDL_GetCameraPermissionState) | [:heavy_check_mark:](sdl/methods.go#L376) | [:x:](sdl/sdl_functions_js.go#L5352) |
| [SDL_GetCameraID](https://wiki.libsdl.org/SDL3/SDL_GetCameraID) | [:heavy_check_mark:](sdl/methods.go#L382) | [:x:](sdl/sdl_functions_js.go#L5368) |
| [SDL_GetCameraProperties](https://wiki.libsdl.org/SDL3/SDL_GetCameraProperties) | [:heavy_check_mark:](sdl/methods.go#L393) | [:x:](sdl/sdl_functions_js.go#L5384) |
| [SDL_GetCameraFormat](https://wiki.libsdl.org/SDL3/SDL_GetCameraFormat) | [:heavy_check_mark:](sdl/methods.go#L404) | [:x:](sdl/sdl_functions_js.go#L5400) |
| [SDL_AcquireCameraFrame](https://wiki.libsdl.org/SDL3/SDL_AcquireCameraFrame) | [:heavy_check_mark:](sdl/methods.go#L416) | [:x:](sdl/sdl_functions_js.go#L5421) |
| [SDL_ReleaseCameraFrame](https://wiki.libsdl.org/SDL3/SDL_ReleaseCameraFrame) | [:heavy_check_mark:](sdl/methods.go#L429) | [:x:](sdl/sdl_functions_js.go#L5445) |
| [SDL_CloseCamera](https://wiki.libsdl.org/SDL3/SDL_CloseCamera) | [:heavy_check_mark:](sdl/methods.go#L435) | [:x:](sdl/sdl_functions_js.go#L5464) |
### Clipboard

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetClipboardText](https://wiki.libsdl.org/SDL3/SDL_SetClipboardText) | [:heavy_check_mark:](sdl/functions.go#L1586) | [:x:](sdl/sdl_functions_js.go#L5478) |
| [SDL_GetClipboardText](https://wiki.libsdl.org/SDL3/SDL_GetClipboardText) | [:heavy_check_mark:](sdl/functions.go#L1596) | [:x:](sdl/sdl_functions_js.go#L5491) |
| [SDL_HasClipboardText](https://wiki.libsdl.org/SDL3/SDL_HasClipboardText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5502) |
| [SDL_SetPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_SetPrimarySelectionText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5513) |
| [SDL_GetPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_GetPrimarySelectionText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5526) |
| [SDL_HasPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_HasPrimarySelectionText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5537) |
| [SDL_SetClipboardData](https://wiki.libsdl.org/SDL3/SDL_SetClipboardData) | [:x:](sdl/methods.go#L41) | [:x:](sdl/sdl_functions_js.go#L5548) |
| [SDL_ClearClipboardData](https://wiki.libsdl.org/SDL3/SDL_ClearClipboardData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5572) |
| [SDL_GetClipboardData](https://wiki.libsdl.org/SDL3/SDL_GetClipboardData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5583) |
| [SDL_HasClipboardData](https://wiki.libsdl.org/SDL3/SDL_HasClipboardData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5601) |
| [SDL_GetClipboardMimeTypes](https://wiki.libsdl.org/SDL3/SDL_GetClipboardMimeTypes) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5614) |
### Dialog

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ShowOpenFileDialog](https://wiki.libsdl.org/SDL3/SDL_ShowOpenFileDialog) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7684) |
| [SDL_ShowSaveFileDialog](https://wiki.libsdl.org/SDL3/SDL_ShowSaveFileDialog) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7713) |
| [SDL_ShowOpenFolderDialog](https://wiki.libsdl.org/SDL3/SDL_ShowOpenFolderDialog) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7740) |
| [SDL_ShowFileDialogWithProperties](https://wiki.libsdl.org/SDL3/SDL_ShowFileDialogWithProperties) | [:x:](sdl/methods.go#L2440) | [:x:](sdl/sdl_functions_js.go#L7762) |
### GPU

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GPUSupportsShaderFormats](https://wiki.libsdl.org/SDL3/SDL_GPUSupportsShaderFormats) | [:heavy_check_mark:](sdl/functions.go#L460) | [:x:](sdl/sdl_functions_js.go#L11373) |
| [SDL_GPUSupportsProperties](https://wiki.libsdl.org/SDL3/SDL_GPUSupportsProperties) | [:heavy_check_mark:](sdl/methods.go#L5825) | [:x:](sdl/sdl_functions_js.go#L11388) |
| [SDL_CreateGPUDevice](https://wiki.libsdl.org/SDL3/SDL_CreateGPUDevice) | [:heavy_check_mark:](sdl/functions.go#L466) | [:x:](sdl/sdl_functions_js.go#L11401) |
| [SDL_CreateGPUDeviceWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateGPUDeviceWithProperties) | [:heavy_check_mark:](sdl/functions.go#L477) | [:x:](sdl/sdl_functions_js.go#L11421) |
| [SDL_DestroyGPUDevice](https://wiki.libsdl.org/SDL3/SDL_DestroyGPUDevice) | [:heavy_check_mark:](sdl/methods.go#L1883) | [:x:](sdl/sdl_functions_js.go#L11437) |
| [SDL_GetNumGPUDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumGPUDrivers) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11451) |
| [SDL_GetGPUDriver](https://wiki.libsdl.org/SDL3/SDL_GetGPUDriver) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11462) |
| [SDL_GetGPUDeviceDriver](https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceDriver) | [:heavy_check_mark:](sdl/methods.go#L1889) | [:x:](sdl/sdl_functions_js.go#L11475) |
| [SDL_GetGPUShaderFormats](https://wiki.libsdl.org/SDL3/SDL_GetGPUShaderFormats) | [:heavy_check_mark:](sdl/methods.go#L1895) | [:x:](sdl/sdl_functions_js.go#L11491) |
| [SDL_GetGPUDeviceProperties](https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceProperties) | [:question:]() | [:question:]() |
| [SDL_CreateGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_CreateGPUComputePipeline) | [:heavy_check_mark:](sdl/methods.go#L1918) | [:x:](sdl/sdl_functions_js.go#L11507) |
| [SDL_CreateGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_CreateGPUGraphicsPipeline) | [:heavy_check_mark:](sdl/methods.go#L1944) | [:x:](sdl/sdl_functions_js.go#L11531) |
| [SDL_CreateGPUSampler](https://wiki.libsdl.org/SDL3/SDL_CreateGPUSampler) | [:heavy_check_mark:](sdl/methods.go#L1955) | [:x:](sdl/sdl_functions_js.go#L11555) |
| [SDL_CreateGPUShader](https://wiki.libsdl.org/SDL3/SDL_CreateGPUShader) | [:heavy_check_mark:](sdl/methods.go#L1974) | [:x:](sdl/sdl_functions_js.go#L11579) |
| [SDL_CreateGPUTexture](https://wiki.libsdl.org/SDL3/SDL_CreateGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L1996) | [:x:](sdl/sdl_functions_js.go#L11603) |
| [SDL_CreateGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_CreateGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2007) | [:x:](sdl/sdl_functions_js.go#L11627) |
| [SDL_CreateGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_CreateGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2018) | [:x:](sdl/sdl_functions_js.go#L11651) |
| [SDL_SetGPUBufferName](https://wiki.libsdl.org/SDL3/SDL_SetGPUBufferName) | [:heavy_check_mark:](sdl/methods.go#L2029) | [:x:](sdl/sdl_functions_js.go#L11675) |
| [SDL_SetGPUTextureName](https://wiki.libsdl.org/SDL3/SDL_SetGPUTextureName) | [:heavy_check_mark:](sdl/methods.go#L2035) | [:x:](sdl/sdl_functions_js.go#L11696) |
| [SDL_InsertGPUDebugLabel](https://wiki.libsdl.org/SDL3/SDL_InsertGPUDebugLabel) | [:heavy_check_mark:](sdl/methods.go#L5926) | [:x:](sdl/sdl_functions_js.go#L11717) |
| [SDL_PushGPUDebugGroup](https://wiki.libsdl.org/SDL3/SDL_PushGPUDebugGroup) | [:heavy_check_mark:](sdl/methods.go#L5932) | [:x:](sdl/sdl_functions_js.go#L11733) |
| [SDL_PopGPUDebugGroup](https://wiki.libsdl.org/SDL3/SDL_PopGPUDebugGroup) | [:heavy_check_mark:](sdl/methods.go#L5938) | [:x:](sdl/sdl_functions_js.go#L11749) |
| [SDL_ReleaseGPUTexture](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2041) | [:x:](sdl/sdl_functions_js.go#L11763) |
| [SDL_ReleaseGPUSampler](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUSampler) | [:heavy_check_mark:](sdl/methods.go#L2047) | [:x:](sdl/sdl_functions_js.go#L11782) |
| [SDL_ReleaseGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2053) | [:x:](sdl/sdl_functions_js.go#L11801) |
| [SDL_ReleaseGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2059) | [:x:](sdl/sdl_functions_js.go#L11820) |
| [SDL_ReleaseGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUComputePipeline) | [:heavy_check_mark:](sdl/methods.go#L2065) | [:x:](sdl/sdl_functions_js.go#L11839) |
| [SDL_ReleaseGPUShader](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUShader) | [:heavy_check_mark:](sdl/methods.go#L2071) | [:x:](sdl/sdl_functions_js.go#L11858) |
| [SDL_ReleaseGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUGraphicsPipeline) | [:heavy_check_mark:](sdl/methods.go#L2077) | [:x:](sdl/sdl_functions_js.go#L11877) |
| [SDL_AcquireGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_AcquireGPUCommandBuffer) | [:heavy_check_mark:](sdl/methods.go#L2083) | [:x:](sdl/sdl_functions_js.go#L11896) |
| [SDL_PushGPUVertexUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUVertexUniformData) | [:heavy_check_mark:](sdl/methods.go#L5944) | [:x:](sdl/sdl_functions_js.go#L11915) |
| [SDL_PushGPUFragmentUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUFragmentUniformData) | [:heavy_check_mark:](sdl/methods.go#L5950) | [:x:](sdl/sdl_functions_js.go#L11935) |
| [SDL_PushGPUComputeUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUComputeUniformData) | [:heavy_check_mark:](sdl/methods.go#L5956) | [:x:](sdl/sdl_functions_js.go#L11955) |
| [SDL_BeginGPURenderPass](https://wiki.libsdl.org/SDL3/SDL_BeginGPURenderPass) | [:heavy_check_mark:](sdl/methods.go#L5962) | [:x:](sdl/sdl_functions_js.go#L11975) |
| [SDL_BindGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_BindGPUGraphicsPipeline) | [:heavy_check_mark:](sdl/methods.go#L1217) | [:x:](sdl/sdl_functions_js.go#L12006) |
| [SDL_SetGPUViewport](https://wiki.libsdl.org/SDL3/SDL_SetGPUViewport) | [:heavy_check_mark:](sdl/methods.go#L1223) | [:x:](sdl/sdl_functions_js.go#L12025) |
| [SDL_SetGPUScissor](https://wiki.libsdl.org/SDL3/SDL_SetGPUScissor) | [:heavy_check_mark:](sdl/methods.go#L1229) | [:x:](sdl/sdl_functions_js.go#L12044) |
| [SDL_SetGPUBlendConstants](https://wiki.libsdl.org/SDL3/SDL_SetGPUBlendConstants) | [:question:]() | [:question:]() |
| [SDL_SetGPUStencilReference](https://wiki.libsdl.org/SDL3/SDL_SetGPUStencilReference) | [:heavy_check_mark:](sdl/methods.go#L1235) | [:x:](sdl/sdl_functions_js.go#L12063) |
| [SDL_BindGPUVertexBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexBuffers) | [:heavy_check_mark:](sdl/methods.go#L1241) | [:x:](sdl/sdl_functions_js.go#L12079) |
| [SDL_BindGPUIndexBuffer](https://wiki.libsdl.org/SDL3/SDL_BindGPUIndexBuffer) | [:heavy_check_mark:](sdl/methods.go#L1247) | [:x:](sdl/sdl_functions_js.go#L12102) |
| [SDL_BindGPUVertexSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexSamplers) | [:heavy_check_mark:](sdl/methods.go#L1253) | [:x:](sdl/sdl_functions_js.go#L12123) |
| [SDL_BindGPUVertexStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageTextures) | [:heavy_check_mark:](sdl/methods.go#L1259) | [:x:](sdl/sdl_functions_js.go#L12146) |
| [SDL_BindGPUVertexStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageBuffers) | [:heavy_check_mark:](sdl/methods.go#L1265) | [:x:](sdl/sdl_functions_js.go#L12169) |
| [SDL_BindGPUFragmentSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentSamplers) | [:heavy_check_mark:](sdl/methods.go#L1271) | [:x:](sdl/sdl_functions_js.go#L12192) |
| [SDL_BindGPUFragmentStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageTextures) | [:heavy_check_mark:](sdl/methods.go#L1277) | [:x:](sdl/sdl_functions_js.go#L12215) |
| [SDL_BindGPUFragmentStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageBuffers) | [:heavy_check_mark:](sdl/methods.go#L1283) | [:x:](sdl/sdl_functions_js.go#L12238) |
| [SDL_DrawGPUIndexedPrimitives](https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitives) | [:heavy_check_mark:](sdl/methods.go#L1289) | [:x:](sdl/sdl_functions_js.go#L12261) |
| [SDL_DrawGPUPrimitives](https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitives) | [:heavy_check_mark:](sdl/methods.go#L1295) | [:x:](sdl/sdl_functions_js.go#L12285) |
| [SDL_DrawGPUPrimitivesIndirect](https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitivesIndirect) | [:heavy_check_mark:](sdl/methods.go#L1301) | [:x:](sdl/sdl_functions_js.go#L12307) |
| [SDL_DrawGPUIndexedPrimitivesIndirect](https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitivesIndirect) | [:heavy_check_mark:](sdl/methods.go#L1307) | [:x:](sdl/sdl_functions_js.go#L12330) |
| [SDL_EndGPURenderPass](https://wiki.libsdl.org/SDL3/SDL_EndGPURenderPass) | [:heavy_check_mark:](sdl/methods.go#L1313) | [:x:](sdl/sdl_functions_js.go#L12353) |
| [SDL_BeginGPUComputePass](https://wiki.libsdl.org/SDL3/SDL_BeginGPUComputePass) | [:heavy_check_mark:](sdl/methods.go#L5968) | [:x:](sdl/sdl_functions_js.go#L12367) |
| [SDL_BindGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputePipeline) | [:heavy_check_mark:](sdl/methods.go#L895) | [:x:](sdl/sdl_functions_js.go#L12400) |
| [SDL_BindGPUComputeSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeSamplers) | [:heavy_check_mark:](sdl/methods.go#L901) | [:x:](sdl/sdl_functions_js.go#L12419) |
| [SDL_BindGPUComputeStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageTextures) | [:heavy_check_mark:](sdl/methods.go#L907) | [:x:](sdl/sdl_functions_js.go#L12442) |
| [SDL_BindGPUComputeStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageBuffers) | [:heavy_check_mark:](sdl/methods.go#L913) | [:x:](sdl/sdl_functions_js.go#L12465) |
| [SDL_DispatchGPUCompute](https://wiki.libsdl.org/SDL3/SDL_DispatchGPUCompute) | [:heavy_check_mark:](sdl/methods.go#L919) | [:x:](sdl/sdl_functions_js.go#L12488) |
| [SDL_DispatchGPUComputeIndirect](https://wiki.libsdl.org/SDL3/SDL_DispatchGPUComputeIndirect) | [:heavy_check_mark:](sdl/methods.go#L925) | [:x:](sdl/sdl_functions_js.go#L12508) |
| [SDL_EndGPUComputePass](https://wiki.libsdl.org/SDL3/SDL_EndGPUComputePass) | [:heavy_check_mark:](sdl/methods.go#L931) | [:x:](sdl/sdl_functions_js.go#L12529) |
| [SDL_MapGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_MapGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2094) | [:x:](sdl/sdl_functions_js.go#L12543) |
| [SDL_UnmapGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_UnmapGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2105) | [:x:](sdl/sdl_functions_js.go#L12566) |
| [SDL_BeginGPUCopyPass](https://wiki.libsdl.org/SDL3/SDL_BeginGPUCopyPass) | [:heavy_check_mark:](sdl/methods.go#L5974) | [:x:](sdl/sdl_functions_js.go#L12585) |
| [SDL_UploadToGPUTexture](https://wiki.libsdl.org/SDL3/SDL_UploadToGPUTexture) | [:x:](sdl/methods.go#L2768) | [:x:](sdl/sdl_functions_js.go#L12604) |
| [SDL_UploadToGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_UploadToGPUBuffer) | [:x:](sdl/methods.go#L2775) | [:x:](sdl/sdl_functions_js.go#L12630) |
| [SDL_CopyGPUTextureToTexture](https://wiki.libsdl.org/SDL3/SDL_CopyGPUTextureToTexture) | [:x:](sdl/methods.go#L2782) | [:x:](sdl/sdl_functions_js.go#L12656) |
| [SDL_CopyGPUBufferToBuffer](https://wiki.libsdl.org/SDL3/SDL_CopyGPUBufferToBuffer) | [:x:](sdl/methods.go#L2789) | [:x:](sdl/sdl_functions_js.go#L12688) |
| [SDL_DownloadFromGPUTexture](https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUTexture) | [:x:](sdl/methods.go#L2796) | [:x:](sdl/sdl_functions_js.go#L12716) |
| [SDL_DownloadFromGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUBuffer) | [:x:](sdl/methods.go#L2803) | [:x:](sdl/sdl_functions_js.go#L12740) |
| [SDL_EndGPUCopyPass](https://wiki.libsdl.org/SDL3/SDL_EndGPUCopyPass) | [:x:](sdl/methods.go#L2810) | [:x:](sdl/sdl_functions_js.go#L12764) |
| [SDL_GenerateMipmapsForGPUTexture](https://wiki.libsdl.org/SDL3/SDL_GenerateMipmapsForGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L5980) | [:x:](sdl/sdl_functions_js.go#L12778) |
| [SDL_BlitGPUTexture](https://wiki.libsdl.org/SDL3/SDL_BlitGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L5986) | [:x:](sdl/sdl_functions_js.go#L12797) |
| [SDL_WindowSupportsGPUSwapchainComposition](https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUSwapchainComposition) | [:heavy_check_mark:](sdl/methods.go#L2111) | [:x:](sdl/sdl_functions_js.go#L12816) |
| [SDL_WindowSupportsGPUPresentMode](https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUPresentMode) | [:heavy_check_mark:](sdl/methods.go#L2117) | [:x:](sdl/sdl_functions_js.go#L12839) |
| [SDL_ClaimWindowForGPUDevice](https://wiki.libsdl.org/SDL3/SDL_ClaimWindowForGPUDevice) | [:heavy_check_mark:](sdl/methods.go#L2123) | [:x:](sdl/sdl_functions_js.go#L12862) |
| [SDL_ReleaseWindowFromGPUDevice](https://wiki.libsdl.org/SDL3/SDL_ReleaseWindowFromGPUDevice) | [:heavy_check_mark:](sdl/methods.go#L2133) | [:x:](sdl/sdl_functions_js.go#L12883) |
| [SDL_SetGPUSwapchainParameters](https://wiki.libsdl.org/SDL3/SDL_SetGPUSwapchainParameters) | [:heavy_check_mark:](sdl/methods.go#L2139) | [:x:](sdl/sdl_functions_js.go#L12902) |
| [SDL_SetGPUAllowedFramesInFlight](https://wiki.libsdl.org/SDL3/SDL_SetGPUAllowedFramesInFlight) | [:heavy_check_mark:](sdl/methods.go#L2149) | [:x:](sdl/sdl_functions_js.go#L12927) |
| [SDL_GetGPUSwapchainTextureFormat](https://wiki.libsdl.org/SDL3/SDL_GetGPUSwapchainTextureFormat) | [:heavy_check_mark:](sdl/methods.go#L2159) | [:x:](sdl/sdl_functions_js.go#L12945) |
| [SDL_AcquireGPUSwapchainTexture](https://wiki.libsdl.org/SDL3/SDL_AcquireGPUSwapchainTexture) | [:heavy_check_mark:](sdl/methods.go#L5992) | [:x:](sdl/sdl_functions_js.go#L12966) |
| [SDL_WaitForGPUSwapchain](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUSwapchain) | [:heavy_check_mark:](sdl/methods.go#L2165) | [:x:](sdl/sdl_functions_js.go#L13002) |
| [SDL_WaitAndAcquireGPUSwapchainTexture](https://wiki.libsdl.org/SDL3/SDL_WaitAndAcquireGPUSwapchainTexture) | [:heavy_check_mark:](sdl/methods.go#L6004) | [:x:](sdl/sdl_functions_js.go#L13023) |
| [SDL_SubmitGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBuffer) | [:heavy_check_mark:](sdl/methods.go#L6016) | [:x:](sdl/sdl_functions_js.go#L13059) |
| [SDL_SubmitGPUCommandBufferAndAcquireFence](https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBufferAndAcquireFence) | [:heavy_check_mark:](sdl/methods.go#L6026) | [:x:](sdl/sdl_functions_js.go#L13075) |
| [SDL_CancelGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_CancelGPUCommandBuffer) | [:heavy_check_mark:](sdl/methods.go#L6037) | [:x:](sdl/sdl_functions_js.go#L13094) |
| [SDL_WaitForGPUIdle](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUIdle) | [:heavy_check_mark:](sdl/methods.go#L2175) | [:x:](sdl/sdl_functions_js.go#L13110) |
| [SDL_WaitForGPUFences](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUFences) | [:heavy_check_mark:](sdl/methods.go#L2185) | [:x:](sdl/sdl_functions_js.go#L13126) |
| [SDL_QueryGPUFence](https://wiki.libsdl.org/SDL3/SDL_QueryGPUFence) | [:heavy_check_mark:](sdl/methods.go#L2195) | [:x:](sdl/sdl_functions_js.go#L13151) |
| [SDL_ReleaseGPUFence](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUFence) | [:heavy_check_mark:](sdl/methods.go#L2201) | [:x:](sdl/sdl_functions_js.go#L13172) |
| [SDL_GPUTextureFormatTexelBlockSize](https://wiki.libsdl.org/SDL3/SDL_GPUTextureFormatTexelBlockSize) | [:heavy_check_mark:](sdl/methods.go#L451) | [:x:](sdl/sdl_functions_js.go#L13191) |
| [SDL_GPUTextureSupportsFormat](https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsFormat) | [:heavy_check_mark:](sdl/methods.go#L2207) | [:x:](sdl/sdl_functions_js.go#L13204) |
| [SDL_GPUTextureSupportsSampleCount](https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsSampleCount) | [:heavy_check_mark:](sdl/methods.go#L2213) | [:x:](sdl/sdl_functions_js.go#L13226) |
| [SDL_CalculateGPUTextureFormatSize](https://wiki.libsdl.org/SDL3/SDL_CalculateGPUTextureFormatSize) | [:heavy_check_mark:](sdl/methods.go#L457) | [:x:](sdl/sdl_functions_js.go#L13246) |
| [SDL_GDKSuspendGPU](https://wiki.libsdl.org/SDL3/SDL_GDKSuspendGPU) | [:question:]() | [:question:]() |
| [SDL_GDKResumeGPU](https://wiki.libsdl.org/SDL3/SDL_GDKResumeGPU) | [:question:]() | [:question:]() |
### MessageBox

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ShowMessageBox](https://wiki.libsdl.org/SDL3/SDL_ShowMessageBox) | [:heavy_check_mark:](sdl/functions.go#L980) | [:x:](sdl/sdl_functions_js.go#L14325) |
| [SDL_ShowSimpleMessageBox](https://wiki.libsdl.org/SDL3/SDL_ShowSimpleMessageBox) | [:heavy_check_mark:](sdl/functions.go#L1013) | [:x:](sdl/sdl_functions_js.go#L14340) |
### Vulkan

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_Vulkan_LoadLibrary](https://wiki.libsdl.org/SDL3/SDL_Vulkan_LoadLibrary) | [:question:]() | [:question:]() |
| [SDL_Vulkan_GetVkGetInstanceProcAddr](https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetVkGetInstanceProcAddr) | [:question:]() | [:question:]() |
| [SDL_Vulkan_UnloadLibrary](https://wiki.libsdl.org/SDL3/SDL_Vulkan_UnloadLibrary) | [:question:]() | [:question:]() |
| [SDL_Vulkan_GetInstanceExtensions](https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetInstanceExtensions) | [:question:]() | [:question:]() |
| [SDL_Vulkan_CreateSurface](https://wiki.libsdl.org/SDL3/SDL_Vulkan_CreateSurface) | [:question:]() | [:question:]() |
| [SDL_Vulkan_DestroySurface](https://wiki.libsdl.org/SDL3/SDL_Vulkan_DestroySurface) | [:question:]() | [:question:]() |
| [SDL_Vulkan_GetPresentationSupport](https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetPresentationSupport) | [:question:]() | [:question:]() |
### Metal

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_Metal_CreateView](https://wiki.libsdl.org/SDL3/SDL_Metal_CreateView) | [:x:](sdl/methods.go#L4734) | [:x:](sdl/sdl_functions_js.go#L14362) |
| [SDL_Metal_DestroyView](https://wiki.libsdl.org/SDL3/SDL_Metal_DestroyView) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14378) |
| [SDL_Metal_GetLayer](https://wiki.libsdl.org/SDL3/SDL_Metal_GetLayer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14389) |
### Platform

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetPowerInfo) | [:heavy_check_mark:](sdl/functions.go#L1031) | [:x:](sdl/sdl_functions_js.go#L7807) |
### Power

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetSensors](https://wiki.libsdl.org/SDL3/SDL_GetSensors) | [:heavy_check_mark:](sdl/functions.go#L1046) | [:x:](sdl/sdl_functions_js.go#L7828) |
| [SDL_GetSensorNameForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorNameForID) | [:heavy_check_mark:](sdl/methods.go#L6128) | [:x:](sdl/sdl_functions_js.go#L7844) |
| [SDL_GetSensorTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorTypeForID) | [:heavy_check_mark:](sdl/methods.go#L6134) | [:x:](sdl/sdl_functions_js.go#L7857) |
| [SDL_GetSensorNonPortableTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableTypeForID) | [:heavy_check_mark:](sdl/methods.go#L6140) | [:x:](sdl/sdl_functions_js.go#L7870) |
| [SDL_OpenSensor](https://wiki.libsdl.org/SDL3/SDL_OpenSensor) | [:heavy_check_mark:](sdl/methods.go#L6146) | [:x:](sdl/sdl_functions_js.go#L7883) |
| [SDL_GetSensorFromID](https://wiki.libsdl.org/SDL3/SDL_GetSensorFromID) | [:heavy_check_mark:](sdl/methods.go#L6157) | [:x:](sdl/sdl_functions_js.go#L7899) |
| [SDL_GetSensorProperties](https://wiki.libsdl.org/SDL3/SDL_GetSensorProperties) | [:heavy_check_mark:](sdl/methods.go#L522) | [:x:](sdl/sdl_functions_js.go#L7915) |
| [SDL_GetSensorName](https://wiki.libsdl.org/SDL3/SDL_GetSensorName) | [:heavy_check_mark:](sdl/methods.go#L533) | [:x:](sdl/sdl_functions_js.go#L7931) |
| [SDL_GetSensorType](https://wiki.libsdl.org/SDL3/SDL_GetSensorType) | [:heavy_check_mark:](sdl/methods.go#L544) | [:x:](sdl/sdl_functions_js.go#L7947) |
| [SDL_GetSensorNonPortableType](https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableType) | [:heavy_check_mark:](sdl/methods.go#L550) | [:x:](sdl/sdl_functions_js.go#L7963) |
| [SDL_GetSensorID](https://wiki.libsdl.org/SDL3/SDL_GetSensorID) | [:heavy_check_mark:](sdl/methods.go#L556) | [:x:](sdl/sdl_functions_js.go#L7979) |
| [SDL_GetSensorData](https://wiki.libsdl.org/SDL3/SDL_GetSensorData) | [:heavy_check_mark:](sdl/methods.go#L567) | [:x:](sdl/sdl_functions_js.go#L7995) |
| [SDL_CloseSensor](https://wiki.libsdl.org/SDL3/SDL_CloseSensor) | [:heavy_check_mark:](sdl/methods.go#L577) | [:x:](sdl/sdl_functions_js.go#L8018) |
| [SDL_UpdateSensors](https://wiki.libsdl.org/SDL3/SDL_UpdateSensors) | [:heavy_check_mark:](sdl/functions.go#L1060) | [:x:](sdl/sdl_functions_js.go#L8032) |
### Sensor

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateProcess](https://wiki.libsdl.org/SDL3/SDL_CreateProcess) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14415) |
| [SDL_CreateProcessWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateProcessWithProperties) | [:heavy_check_mark:](sdl/functions.go#L1068) | [:x:](sdl/sdl_functions_js.go#L14435) |
| [SDL_GetProcessProperties](https://wiki.libsdl.org/SDL3/SDL_GetProcessProperties) | [:heavy_check_mark:](sdl/methods.go#L6049) | [:x:](sdl/sdl_functions_js.go#L14451) |
| [SDL_ReadProcess](https://wiki.libsdl.org/SDL3/SDL_ReadProcess) | [:heavy_check_mark:](sdl/methods.go#L6060) | [:x:](sdl/sdl_functions_js.go#L14467) |
| [SDL_GetProcessInput](https://wiki.libsdl.org/SDL3/SDL_GetProcessInput) | [:heavy_check_mark:](sdl/methods.go#L6078) | [:x:](sdl/sdl_functions_js.go#L14493) |
| [SDL_GetProcessOutput](https://wiki.libsdl.org/SDL3/SDL_GetProcessOutput) | [:heavy_check_mark:](sdl/methods.go#L6089) | [:x:](sdl/sdl_functions_js.go#L14512) |
| [SDL_KillProcess](https://wiki.libsdl.org/SDL3/SDL_KillProcess) | [:heavy_check_mark:](sdl/methods.go#L6100) | [:x:](sdl/sdl_functions_js.go#L14531) |
| [SDL_WaitProcess](https://wiki.libsdl.org/SDL3/SDL_WaitProcess) | [:heavy_check_mark:](sdl/methods.go#L6110) | [:x:](sdl/sdl_functions_js.go#L14549) |
| [SDL_DestroyProcess](https://wiki.libsdl.org/SDL3/SDL_DestroyProcess) | [:heavy_check_mark:](sdl/methods.go#L6120) | [:x:](sdl/sdl_functions_js.go#L14572) |
### Process

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_MostSignificantBitIndex32](https://wiki.libsdl.org/SDL3/SDL_MostSignificantBitIndex32) | [:question:]() | [:question:]() |
| [SDL_HasExactlyOneBitSet32](https://wiki.libsdl.org/SDL3/SDL_HasExactlyOneBitSet32) | [:question:]() | [:question:]() |
### Bits

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SwapFloat](https://wiki.libsdl.org/SDL3/SDL_SwapFloat) | [:question:]() | [:question:]() |
| [SDL_Swap16](https://wiki.libsdl.org/SDL3/SDL_Swap16) | [:question:]() | [:question:]() |
| [SDL_Swap32](https://wiki.libsdl.org/SDL3/SDL_Swap32) | [:question:]() | [:question:]() |
| [SDL_Swap64](https://wiki.libsdl.org/SDL3/SDL_Swap64) | [:question:]() | [:question:]() |
### Endian

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ReportAssertion](https://wiki.libsdl.org/SDL3/SDL_ReportAssertion) | [:question:]() | [:question:]() |
| [SDL_SetAssertionHandler](https://wiki.libsdl.org/SDL3/SDL_SetAssertionHandler) | [:question:]() | [:question:]() |
| [SDL_GetDefaultAssertionHandler](https://wiki.libsdl.org/SDL3/SDL_GetDefaultAssertionHandler) | [:question:]() | [:question:]() |
| [SDL_GetAssertionHandler](https://wiki.libsdl.org/SDL3/SDL_GetAssertionHandler) | [:question:]() | [:question:]() |
| [SDL_GetAssertionReport](https://wiki.libsdl.org/SDL3/SDL_GetAssertionReport) | [:question:]() | [:question:]() |
| [SDL_ResetAssertionReport](https://wiki.libsdl.org/SDL3/SDL_ResetAssertionReport) | [:question:]() | [:question:]() |
### Assert

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumLogicalCPUCores](https://wiki.libsdl.org/SDL3/SDL_GetNumLogicalCPUCores) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5630) |
| [SDL_GetCPUCacheLineSize](https://wiki.libsdl.org/SDL3/SDL_GetCPUCacheLineSize) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5641) |
| [SDL_HasAltiVec](https://wiki.libsdl.org/SDL3/SDL_HasAltiVec) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5652) |
| [SDL_HasMMX](https://wiki.libsdl.org/SDL3/SDL_HasMMX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5663) |
| [SDL_HasSSE](https://wiki.libsdl.org/SDL3/SDL_HasSSE) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5674) |
| [SDL_HasSSE2](https://wiki.libsdl.org/SDL3/SDL_HasSSE2) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5685) |
| [SDL_HasSSE3](https://wiki.libsdl.org/SDL3/SDL_HasSSE3) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5696) |
| [SDL_HasSSE41](https://wiki.libsdl.org/SDL3/SDL_HasSSE41) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5707) |
| [SDL_HasSSE42](https://wiki.libsdl.org/SDL3/SDL_HasSSE42) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5718) |
| [SDL_HasAVX](https://wiki.libsdl.org/SDL3/SDL_HasAVX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5729) |
| [SDL_HasAVX2](https://wiki.libsdl.org/SDL3/SDL_HasAVX2) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5740) |
| [SDL_HasAVX512F](https://wiki.libsdl.org/SDL3/SDL_HasAVX512F) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5751) |
| [SDL_HasARMSIMD](https://wiki.libsdl.org/SDL3/SDL_HasARMSIMD) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5762) |
| [SDL_HasNEON](https://wiki.libsdl.org/SDL3/SDL_HasNEON) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5773) |
| [SDL_HasLSX](https://wiki.libsdl.org/SDL3/SDL_HasLSX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5784) |
| [SDL_HasLASX](https://wiki.libsdl.org/SDL3/SDL_HasLASX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5795) |
| [SDL_GetSystemRAM](https://wiki.libsdl.org/SDL3/SDL_GetSystemRAM) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5806) |
| [SDL_GetSIMDAlignment](https://wiki.libsdl.org/SDL3/SDL_GetSIMDAlignment) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5817) |
### CPUInfo

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPreferredLocales](https://wiki.libsdl.org/SDL3/SDL_GetPreferredLocales) | [:heavy_check_mark:](sdl/functions.go#L1094) | [:x:](sdl/sdl_functions_js.go#L14071) |
### Intrinsics

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetWindowsMessageHook](https://wiki.libsdl.org/SDL3/SDL_SetWindowsMessageHook) | [:question:]() | [:question:]() |
| [SDL_GetDirect3D9AdapterIndex](https://wiki.libsdl.org/SDL3/SDL_GetDirect3D9AdapterIndex) | [:question:]() | [:question:]() |
| [SDL_GetDXGIOutputInfo](https://wiki.libsdl.org/SDL3/SDL_GetDXGIOutputInfo) | [:question:]() | [:question:]() |
| [SDL_SetX11EventHook](https://wiki.libsdl.org/SDL3/SDL_SetX11EventHook) | [:x:](sdl/methods.go#L601) | [:x:](sdl/sdl_functions_js.go#L16862) |
| [SDL_SetLinuxThreadPriority](https://wiki.libsdl.org/SDL3/SDL_SetLinuxThreadPriority) | [:question:]() | [:question:]() |
| [SDL_SetLinuxThreadPriorityAndPolicy](https://wiki.libsdl.org/SDL3/SDL_SetLinuxThreadPriorityAndPolicy) | [:question:]() | [:question:]() |
| [SDL_SetiOSAnimationCallback](https://wiki.libsdl.org/SDL3/SDL_SetiOSAnimationCallback) | [:question:]() | [:question:]() |
| [SDL_SetiOSEventPump](https://wiki.libsdl.org/SDL3/SDL_SetiOSEventPump) | [:question:]() | [:question:]() |
| [SDL_GetAndroidJNIEnv](https://wiki.libsdl.org/SDL3/SDL_GetAndroidJNIEnv) | [:question:]() | [:question:]() |
| [SDL_GetAndroidActivity](https://wiki.libsdl.org/SDL3/SDL_GetAndroidActivity) | [:question:]() | [:question:]() |
| [SDL_GetAndroidSDKVersion](https://wiki.libsdl.org/SDL3/SDL_GetAndroidSDKVersion) | [:question:]() | [:question:]() |
| [SDL_IsChromebook](https://wiki.libsdl.org/SDL3/SDL_IsChromebook) | [:question:]() | [:question:]() |
| [SDL_IsDeXMode](https://wiki.libsdl.org/SDL3/SDL_IsDeXMode) | [:question:]() | [:question:]() |
| [SDL_SendAndroidBackButton](https://wiki.libsdl.org/SDL3/SDL_SendAndroidBackButton) | [:question:]() | [:question:]() |
| [SDL_GetAndroidInternalStoragePath](https://wiki.libsdl.org/SDL3/SDL_GetAndroidInternalStoragePath) | [:question:]() | [:question:]() |
| [SDL_GetAndroidExternalStorageState](https://wiki.libsdl.org/SDL3/SDL_GetAndroidExternalStorageState) | [:question:]() | [:question:]() |
| [SDL_GetAndroidExternalStoragePath](https://wiki.libsdl.org/SDL3/SDL_GetAndroidExternalStoragePath) | [:question:]() | [:question:]() |
| [SDL_GetAndroidCachePath](https://wiki.libsdl.org/SDL3/SDL_GetAndroidCachePath) | [:question:]() | [:question:]() |
| [SDL_RequestAndroidPermission](https://wiki.libsdl.org/SDL3/SDL_RequestAndroidPermission) | [:question:]() | [:question:]() |
| [SDL_ShowAndroidToast](https://wiki.libsdl.org/SDL3/SDL_ShowAndroidToast) | [:question:]() | [:question:]() |
| [SDL_SendAndroidMessage](https://wiki.libsdl.org/SDL3/SDL_SendAndroidMessage) | [:question:]() | [:question:]() |
| [SDL_IsTablet](https://wiki.libsdl.org/SDL3/SDL_IsTablet) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16875) |
| [SDL_IsTV](https://wiki.libsdl.org/SDL3/SDL_IsTV) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16886) |
| [SDL_GetSandbox](https://wiki.libsdl.org/SDL3/SDL_GetSandbox) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16897) |
| [SDL_OnApplicationWillTerminate](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillTerminate) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16908) |
| [SDL_OnApplicationDidReceiveMemoryWarning](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidReceiveMemoryWarning) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16917) |
| [SDL_OnApplicationWillEnterBackground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillEnterBackground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16926) |
| [SDL_OnApplicationDidEnterBackground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidEnterBackground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16935) |
| [SDL_OnApplicationWillEnterForeground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillEnterForeground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16944) |
| [SDL_OnApplicationDidEnterForeground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidEnterForeground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16953) |
| [SDL_OnApplicationDidChangeStatusBarOrientation](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidChangeStatusBarOrientation) | [:question:]() | [:question:]() |
| [SDL_GetGDKTaskQueue](https://wiki.libsdl.org/SDL3/SDL_GetGDKTaskQueue) | [:question:]() | [:question:]() |
| [SDL_GetGDKDefaultUser](https://wiki.libsdl.org/SDL3/SDL_GetGDKDefaultUser) | [:question:]() | [:question:]() |
### Locale

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_OpenURL](https://wiki.libsdl.org/SDL3/SDL_OpenURL) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14402) |
### System

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GUIDToString](https://wiki.libsdl.org/SDL3/SDL_GUIDToString) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7779) |
| [SDL_StringToGUID](https://wiki.libsdl.org/SDL3/SDL_StringToGUID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7794) |
### Misc

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AppInit](https://wiki.libsdl.org/SDL3/SDL_AppInit) | [:question:]() | [:question:]() |
| [SDL_AppIterate](https://wiki.libsdl.org/SDL3/SDL_AppIterate) | [:question:]() | [:question:]() |
| [SDL_AppEvent](https://wiki.libsdl.org/SDL3/SDL_AppEvent) | [:question:]() | [:question:]() |
| [SDL_AppQuit](https://wiki.libsdl.org/SDL3/SDL_AppQuit) | [:question:]() | [:question:]() |
| [SDL_main](https://wiki.libsdl.org/SDL3/SDL_main) | [:question:]() | [:question:]() |
| [SDL_SetMainReady](https://wiki.libsdl.org/SDL3/SDL_SetMainReady) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L17246) |
| [SDL_RunApp](https://wiki.libsdl.org/SDL3/SDL_RunApp) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L17255) |
| [SDL_EnterAppMainCallbacks](https://wiki.libsdl.org/SDL3/SDL_EnterAppMainCallbacks) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L17277) |
| [SDL_RegisterApp](https://wiki.libsdl.org/SDL3/SDL_RegisterApp) | [:question:]() | [:question:]() |
| [SDL_UnregisterApp](https://wiki.libsdl.org/SDL3/SDL_UnregisterApp) | [:question:]() | [:question:]() |
| [SDL_GDKSuspendComplete](https://wiki.libsdl.org/SDL3/SDL_GDKSuspendComplete) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L17303) |
### GUID

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_malloc](https://wiki.libsdl.org/SDL3/SDL_malloc) | [:question:]() | [:question:]() |
| [SDL_calloc](https://wiki.libsdl.org/SDL3/SDL_calloc) | [:question:]() | [:question:]() |
| [SDL_realloc](https://wiki.libsdl.org/SDL3/SDL_realloc) | [:question:]() | [:question:]() |
| [SDL_free](https://wiki.libsdl.org/SDL3/SDL_free) | [:question:]() | [:question:]() |
| [SDL_GetOriginalMemoryFunctions](https://wiki.libsdl.org/SDL3/SDL_GetOriginalMemoryFunctions) | [:question:]() | [:question:]() |
| [SDL_GetMemoryFunctions](https://wiki.libsdl.org/SDL3/SDL_GetMemoryFunctions) | [:question:]() | [:question:]() |
| [SDL_SetMemoryFunctions](https://wiki.libsdl.org/SDL3/SDL_SetMemoryFunctions) | [:question:]() | [:question:]() |
| [SDL_aligned_alloc](https://wiki.libsdl.org/SDL3/SDL_aligned_alloc) | [:question:]() | [:question:]() |
| [SDL_aligned_free](https://wiki.libsdl.org/SDL3/SDL_aligned_free) | [:question:]() | [:question:]() |
| [SDL_GetNumAllocations](https://wiki.libsdl.org/SDL3/SDL_GetNumAllocations) | [:question:]() | [:question:]() |
| [SDL_GetEnvironment](https://wiki.libsdl.org/SDL3/SDL_GetEnvironment) | [:question:]() | [:question:]() |
| [SDL_CreateEnvironment](https://wiki.libsdl.org/SDL3/SDL_CreateEnvironment) | [:question:]() | [:question:]() |
| [SDL_GetEnvironmentVariable](https://wiki.libsdl.org/SDL3/SDL_GetEnvironmentVariable) | [:question:]() | [:question:]() |
| [SDL_GetEnvironmentVariables](https://wiki.libsdl.org/SDL3/SDL_GetEnvironmentVariables) | [:question:]() | [:question:]() |
| [SDL_SetEnvironmentVariable](https://wiki.libsdl.org/SDL3/SDL_SetEnvironmentVariable) | [:question:]() | [:question:]() |
| [SDL_UnsetEnvironmentVariable](https://wiki.libsdl.org/SDL3/SDL_UnsetEnvironmentVariable) | [:question:]() | [:question:]() |
| [SDL_DestroyEnvironment](https://wiki.libsdl.org/SDL3/SDL_DestroyEnvironment) | [:question:]() | [:question:]() |
| [SDL_getenv](https://wiki.libsdl.org/SDL3/SDL_getenv) | [:question:]() | [:question:]() |
| [SDL_getenv_unsafe](https://wiki.libsdl.org/SDL3/SDL_getenv_unsafe) | [:question:]() | [:question:]() |
| [SDL_setenv_unsafe](https://wiki.libsdl.org/SDL3/SDL_setenv_unsafe) | [:question:]() | [:question:]() |
| [SDL_unsetenv_unsafe](https://wiki.libsdl.org/SDL3/SDL_unsetenv_unsafe) | [:question:]() | [:question:]() |
| [SDL_qsort](https://wiki.libsdl.org/SDL3/SDL_qsort) | [:question:]() | [:question:]() |
| [SDL_bsearch](https://wiki.libsdl.org/SDL3/SDL_bsearch) | [:question:]() | [:question:]() |
| [SDL_qsort_r](https://wiki.libsdl.org/SDL3/SDL_qsort_r) | [:question:]() | [:question:]() |
| [SDL_bsearch_r](https://wiki.libsdl.org/SDL3/SDL_bsearch_r) | [:question:]() | [:question:]() |
| [SDL_abs](https://wiki.libsdl.org/SDL3/SDL_abs) | [:question:]() | [:question:]() |
| [SDL_isalpha](https://wiki.libsdl.org/SDL3/SDL_isalpha) | [:question:]() | [:question:]() |
| [SDL_isalnum](https://wiki.libsdl.org/SDL3/SDL_isalnum) | [:question:]() | [:question:]() |
| [SDL_isblank](https://wiki.libsdl.org/SDL3/SDL_isblank) | [:question:]() | [:question:]() |
| [SDL_iscntrl](https://wiki.libsdl.org/SDL3/SDL_iscntrl) | [:question:]() | [:question:]() |
| [SDL_isdigit](https://wiki.libsdl.org/SDL3/SDL_isdigit) | [:question:]() | [:question:]() |
| [SDL_isxdigit](https://wiki.libsdl.org/SDL3/SDL_isxdigit) | [:question:]() | [:question:]() |
| [SDL_ispunct](https://wiki.libsdl.org/SDL3/SDL_ispunct) | [:question:]() | [:question:]() |
| [SDL_isspace](https://wiki.libsdl.org/SDL3/SDL_isspace) | [:question:]() | [:question:]() |
| [SDL_isupper](https://wiki.libsdl.org/SDL3/SDL_isupper) | [:question:]() | [:question:]() |
| [SDL_islower](https://wiki.libsdl.org/SDL3/SDL_islower) | [:question:]() | [:question:]() |
| [SDL_isprint](https://wiki.libsdl.org/SDL3/SDL_isprint) | [:question:]() | [:question:]() |
| [SDL_isgraph](https://wiki.libsdl.org/SDL3/SDL_isgraph) | [:question:]() | [:question:]() |
| [SDL_toupper](https://wiki.libsdl.org/SDL3/SDL_toupper) | [:question:]() | [:question:]() |
| [SDL_tolower](https://wiki.libsdl.org/SDL3/SDL_tolower) | [:question:]() | [:question:]() |
| [SDL_crc16](https://wiki.libsdl.org/SDL3/SDL_crc16) | [:question:]() | [:question:]() |
| [SDL_crc32](https://wiki.libsdl.org/SDL3/SDL_crc32) | [:question:]() | [:question:]() |
| [SDL_murmur3_32](https://wiki.libsdl.org/SDL3/SDL_murmur3_32) | [:question:]() | [:question:]() |
| [SDL_memcpy](https://wiki.libsdl.org/SDL3/SDL_memcpy) | [:question:]() | [:question:]() |
| [SDL_memmove](https://wiki.libsdl.org/SDL3/SDL_memmove) | [:question:]() | [:question:]() |
| [SDL_memset](https://wiki.libsdl.org/SDL3/SDL_memset) | [:question:]() | [:question:]() |
| [SDL_memset4](https://wiki.libsdl.org/SDL3/SDL_memset4) | [:question:]() | [:question:]() |
| [SDL_memcmp](https://wiki.libsdl.org/SDL3/SDL_memcmp) | [:question:]() | [:question:]() |
| [SDL_wcslen](https://wiki.libsdl.org/SDL3/SDL_wcslen) | [:question:]() | [:question:]() |
| [SDL_wcsnlen](https://wiki.libsdl.org/SDL3/SDL_wcsnlen) | [:question:]() | [:question:]() |
| [SDL_wcslcpy](https://wiki.libsdl.org/SDL3/SDL_wcslcpy) | [:question:]() | [:question:]() |
| [SDL_wcslcat](https://wiki.libsdl.org/SDL3/SDL_wcslcat) | [:question:]() | [:question:]() |
| [SDL_wcsdup](https://wiki.libsdl.org/SDL3/SDL_wcsdup) | [:question:]() | [:question:]() |
| [SDL_wcsstr](https://wiki.libsdl.org/SDL3/SDL_wcsstr) | [:question:]() | [:question:]() |
| [SDL_wcsnstr](https://wiki.libsdl.org/SDL3/SDL_wcsnstr) | [:question:]() | [:question:]() |
| [SDL_wcscmp](https://wiki.libsdl.org/SDL3/SDL_wcscmp) | [:question:]() | [:question:]() |
| [SDL_wcsncmp](https://wiki.libsdl.org/SDL3/SDL_wcsncmp) | [:question:]() | [:question:]() |
| [SDL_wcscasecmp](https://wiki.libsdl.org/SDL3/SDL_wcscasecmp) | [:question:]() | [:question:]() |
| [SDL_wcsncasecmp](https://wiki.libsdl.org/SDL3/SDL_wcsncasecmp) | [:question:]() | [:question:]() |
| [SDL_wcstol](https://wiki.libsdl.org/SDL3/SDL_wcstol) | [:question:]() | [:question:]() |
| [SDL_strlen](https://wiki.libsdl.org/SDL3/SDL_strlen) | [:question:]() | [:question:]() |
| [SDL_strnlen](https://wiki.libsdl.org/SDL3/SDL_strnlen) | [:question:]() | [:question:]() |
| [SDL_strlcpy](https://wiki.libsdl.org/SDL3/SDL_strlcpy) | [:question:]() | [:question:]() |
| [SDL_utf8strlcpy](https://wiki.libsdl.org/SDL3/SDL_utf8strlcpy) | [:question:]() | [:question:]() |
| [SDL_strlcat](https://wiki.libsdl.org/SDL3/SDL_strlcat) | [:question:]() | [:question:]() |
| [SDL_strdup](https://wiki.libsdl.org/SDL3/SDL_strdup) | [:question:]() | [:question:]() |
| [SDL_strndup](https://wiki.libsdl.org/SDL3/SDL_strndup) | [:question:]() | [:question:]() |
| [SDL_strrev](https://wiki.libsdl.org/SDL3/SDL_strrev) | [:question:]() | [:question:]() |
| [SDL_strupr](https://wiki.libsdl.org/SDL3/SDL_strupr) | [:question:]() | [:question:]() |
| [SDL_strlwr](https://wiki.libsdl.org/SDL3/SDL_strlwr) | [:question:]() | [:question:]() |
| [SDL_strchr](https://wiki.libsdl.org/SDL3/SDL_strchr) | [:question:]() | [:question:]() |
| [SDL_strrchr](https://wiki.libsdl.org/SDL3/SDL_strrchr) | [:question:]() | [:question:]() |
| [SDL_strstr](https://wiki.libsdl.org/SDL3/SDL_strstr) | [:question:]() | [:question:]() |
| [SDL_strnstr](https://wiki.libsdl.org/SDL3/SDL_strnstr) | [:question:]() | [:question:]() |
| [SDL_strcasestr](https://wiki.libsdl.org/SDL3/SDL_strcasestr) | [:question:]() | [:question:]() |
| [SDL_strtok_r](https://wiki.libsdl.org/SDL3/SDL_strtok_r) | [:question:]() | [:question:]() |
| [SDL_utf8strlen](https://wiki.libsdl.org/SDL3/SDL_utf8strlen) | [:question:]() | [:question:]() |
| [SDL_utf8strnlen](https://wiki.libsdl.org/SDL3/SDL_utf8strnlen) | [:question:]() | [:question:]() |
| [SDL_itoa](https://wiki.libsdl.org/SDL3/SDL_itoa) | [:question:]() | [:question:]() |
| [SDL_uitoa](https://wiki.libsdl.org/SDL3/SDL_uitoa) | [:question:]() | [:question:]() |
| [SDL_ltoa](https://wiki.libsdl.org/SDL3/SDL_ltoa) | [:question:]() | [:question:]() |
| [SDL_ultoa](https://wiki.libsdl.org/SDL3/SDL_ultoa) | [:question:]() | [:question:]() |
| [SDL_lltoa](https://wiki.libsdl.org/SDL3/SDL_lltoa) | [:question:]() | [:question:]() |
| [SDL_ulltoa](https://wiki.libsdl.org/SDL3/SDL_ulltoa) | [:question:]() | [:question:]() |
| [SDL_atoi](https://wiki.libsdl.org/SDL3/SDL_atoi) | [:question:]() | [:question:]() |
| [SDL_atof](https://wiki.libsdl.org/SDL3/SDL_atof) | [:question:]() | [:question:]() |
| [SDL_strtol](https://wiki.libsdl.org/SDL3/SDL_strtol) | [:question:]() | [:question:]() |
| [SDL_strtoul](https://wiki.libsdl.org/SDL3/SDL_strtoul) | [:question:]() | [:question:]() |
| [SDL_strtoll](https://wiki.libsdl.org/SDL3/SDL_strtoll) | [:question:]() | [:question:]() |
| [SDL_strtoull](https://wiki.libsdl.org/SDL3/SDL_strtoull) | [:question:]() | [:question:]() |
| [SDL_strtod](https://wiki.libsdl.org/SDL3/SDL_strtod) | [:question:]() | [:question:]() |
| [SDL_strcmp](https://wiki.libsdl.org/SDL3/SDL_strcmp) | [:question:]() | [:question:]() |
| [SDL_strncmp](https://wiki.libsdl.org/SDL3/SDL_strncmp) | [:question:]() | [:question:]() |
| [SDL_strcasecmp](https://wiki.libsdl.org/SDL3/SDL_strcasecmp) | [:question:]() | [:question:]() |
| [SDL_strncasecmp](https://wiki.libsdl.org/SDL3/SDL_strncasecmp) | [:question:]() | [:question:]() |
| [SDL_strpbrk](https://wiki.libsdl.org/SDL3/SDL_strpbrk) | [:question:]() | [:question:]() |
| [SDL_StepUTF8](https://wiki.libsdl.org/SDL3/SDL_StepUTF8) | [:question:]() | [:question:]() |
| [SDL_StepBackUTF8](https://wiki.libsdl.org/SDL3/SDL_StepBackUTF8) | [:question:]() | [:question:]() |
| [SDL_UCS4ToUTF8](https://wiki.libsdl.org/SDL3/SDL_UCS4ToUTF8) | [:question:]() | [:question:]() |
| [SDL_sscanf](https://wiki.libsdl.org/SDL3/SDL_sscanf) | [:question:]() | [:question:]() |
| [SDL_vsscanf](https://wiki.libsdl.org/SDL3/SDL_vsscanf) | [:question:]() | [:question:]() |
| [SDL_snprintf](https://wiki.libsdl.org/SDL3/SDL_snprintf) | [:question:]() | [:question:]() |
| [SDL_swprintf](https://wiki.libsdl.org/SDL3/SDL_swprintf) | [:question:]() | [:question:]() |
| [SDL_vsnprintf](https://wiki.libsdl.org/SDL3/SDL_vsnprintf) | [:question:]() | [:question:]() |
| [SDL_vswprintf](https://wiki.libsdl.org/SDL3/SDL_vswprintf) | [:question:]() | [:question:]() |
| [SDL_asprintf](https://wiki.libsdl.org/SDL3/SDL_asprintf) | [:question:]() | [:question:]() |
| [SDL_vasprintf](https://wiki.libsdl.org/SDL3/SDL_vasprintf) | [:question:]() | [:question:]() |
| [SDL_srand](https://wiki.libsdl.org/SDL3/SDL_srand) | [:question:]() | [:question:]() |
| [SDL_rand](https://wiki.libsdl.org/SDL3/SDL_rand) | [:question:]() | [:question:]() |
| [SDL_randf](https://wiki.libsdl.org/SDL3/SDL_randf) | [:question:]() | [:question:]() |
| [SDL_rand_bits](https://wiki.libsdl.org/SDL3/SDL_rand_bits) | [:question:]() | [:question:]() |
| [SDL_rand_r](https://wiki.libsdl.org/SDL3/SDL_rand_r) | [:question:]() | [:question:]() |
| [SDL_randf_r](https://wiki.libsdl.org/SDL3/SDL_randf_r) | [:question:]() | [:question:]() |
| [SDL_rand_bits_r](https://wiki.libsdl.org/SDL3/SDL_rand_bits_r) | [:question:]() | [:question:]() |
| [SDL_acos](https://wiki.libsdl.org/SDL3/SDL_acos) | [:question:]() | [:question:]() |
| [SDL_acosf](https://wiki.libsdl.org/SDL3/SDL_acosf) | [:question:]() | [:question:]() |
| [SDL_asin](https://wiki.libsdl.org/SDL3/SDL_asin) | [:question:]() | [:question:]() |
| [SDL_asinf](https://wiki.libsdl.org/SDL3/SDL_asinf) | [:question:]() | [:question:]() |
| [SDL_atan](https://wiki.libsdl.org/SDL3/SDL_atan) | [:question:]() | [:question:]() |
| [SDL_atanf](https://wiki.libsdl.org/SDL3/SDL_atanf) | [:question:]() | [:question:]() |
| [SDL_atan2](https://wiki.libsdl.org/SDL3/SDL_atan2) | [:question:]() | [:question:]() |
| [SDL_atan2f](https://wiki.libsdl.org/SDL3/SDL_atan2f) | [:question:]() | [:question:]() |
| [SDL_ceil](https://wiki.libsdl.org/SDL3/SDL_ceil) | [:question:]() | [:question:]() |
| [SDL_ceilf](https://wiki.libsdl.org/SDL3/SDL_ceilf) | [:question:]() | [:question:]() |
| [SDL_copysign](https://wiki.libsdl.org/SDL3/SDL_copysign) | [:question:]() | [:question:]() |
| [SDL_copysignf](https://wiki.libsdl.org/SDL3/SDL_copysignf) | [:question:]() | [:question:]() |
| [SDL_cos](https://wiki.libsdl.org/SDL3/SDL_cos) | [:question:]() | [:question:]() |
| [SDL_cosf](https://wiki.libsdl.org/SDL3/SDL_cosf) | [:question:]() | [:question:]() |
| [SDL_exp](https://wiki.libsdl.org/SDL3/SDL_exp) | [:question:]() | [:question:]() |
| [SDL_expf](https://wiki.libsdl.org/SDL3/SDL_expf) | [:question:]() | [:question:]() |
| [SDL_fabs](https://wiki.libsdl.org/SDL3/SDL_fabs) | [:question:]() | [:question:]() |
| [SDL_fabsf](https://wiki.libsdl.org/SDL3/SDL_fabsf) | [:question:]() | [:question:]() |
| [SDL_floor](https://wiki.libsdl.org/SDL3/SDL_floor) | [:question:]() | [:question:]() |
| [SDL_floorf](https://wiki.libsdl.org/SDL3/SDL_floorf) | [:question:]() | [:question:]() |
| [SDL_trunc](https://wiki.libsdl.org/SDL3/SDL_trunc) | [:question:]() | [:question:]() |
| [SDL_truncf](https://wiki.libsdl.org/SDL3/SDL_truncf) | [:question:]() | [:question:]() |
| [SDL_fmod](https://wiki.libsdl.org/SDL3/SDL_fmod) | [:question:]() | [:question:]() |
| [SDL_fmodf](https://wiki.libsdl.org/SDL3/SDL_fmodf) | [:question:]() | [:question:]() |
| [SDL_isinf](https://wiki.libsdl.org/SDL3/SDL_isinf) | [:question:]() | [:question:]() |
| [SDL_isinff](https://wiki.libsdl.org/SDL3/SDL_isinff) | [:question:]() | [:question:]() |
| [SDL_isnan](https://wiki.libsdl.org/SDL3/SDL_isnan) | [:question:]() | [:question:]() |
| [SDL_isnanf](https://wiki.libsdl.org/SDL3/SDL_isnanf) | [:question:]() | [:question:]() |
| [SDL_log](https://wiki.libsdl.org/SDL3/SDL_log) | [:question:]() | [:question:]() |
| [SDL_logf](https://wiki.libsdl.org/SDL3/SDL_logf) | [:question:]() | [:question:]() |
| [SDL_log10](https://wiki.libsdl.org/SDL3/SDL_log10) | [:question:]() | [:question:]() |
| [SDL_log10f](https://wiki.libsdl.org/SDL3/SDL_log10f) | [:question:]() | [:question:]() |
| [SDL_modf](https://wiki.libsdl.org/SDL3/SDL_modf) | [:question:]() | [:question:]() |
| [SDL_modff](https://wiki.libsdl.org/SDL3/SDL_modff) | [:question:]() | [:question:]() |
| [SDL_pow](https://wiki.libsdl.org/SDL3/SDL_pow) | [:question:]() | [:question:]() |
| [SDL_powf](https://wiki.libsdl.org/SDL3/SDL_powf) | [:question:]() | [:question:]() |
| [SDL_round](https://wiki.libsdl.org/SDL3/SDL_round) | [:question:]() | [:question:]() |
| [SDL_roundf](https://wiki.libsdl.org/SDL3/SDL_roundf) | [:question:]() | [:question:]() |
| [SDL_lround](https://wiki.libsdl.org/SDL3/SDL_lround) | [:question:]() | [:question:]() |
| [SDL_lroundf](https://wiki.libsdl.org/SDL3/SDL_lroundf) | [:question:]() | [:question:]() |
| [SDL_scalbn](https://wiki.libsdl.org/SDL3/SDL_scalbn) | [:question:]() | [:question:]() |
| [SDL_scalbnf](https://wiki.libsdl.org/SDL3/SDL_scalbnf) | [:question:]() | [:question:]() |
| [SDL_sin](https://wiki.libsdl.org/SDL3/SDL_sin) | [:question:]() | [:question:]() |
| [SDL_sinf](https://wiki.libsdl.org/SDL3/SDL_sinf) | [:question:]() | [:question:]() |
| [SDL_sqrt](https://wiki.libsdl.org/SDL3/SDL_sqrt) | [:question:]() | [:question:]() |
| [SDL_sqrtf](https://wiki.libsdl.org/SDL3/SDL_sqrtf) | [:question:]() | [:question:]() |
| [SDL_tan](https://wiki.libsdl.org/SDL3/SDL_tan) | [:question:]() | [:question:]() |
| [SDL_tanf](https://wiki.libsdl.org/SDL3/SDL_tanf) | [:question:]() | [:question:]() |
| [SDL_iconv_open](https://wiki.libsdl.org/SDL3/SDL_iconv_open) | [:question:]() | [:question:]() |
| [SDL_iconv_close](https://wiki.libsdl.org/SDL3/SDL_iconv_close) | [:question:]() | [:question:]() |
| [SDL_iconv](https://wiki.libsdl.org/SDL3/SDL_iconv) | [:question:]() | [:question:]() |
| [SDL_iconv_string](https://wiki.libsdl.org/SDL3/SDL_iconv_string) | [:question:]() | [:question:]() |
| [SDL_size_mul_check_overflow](https://wiki.libsdl.org/SDL3/SDL_size_mul_check_overflow) | [:question:]() | [:question:]() |
| [SDL_size_add_check_overflow](https://wiki.libsdl.org/SDL3/SDL_size_add_check_overflow) | [:question:]() | [:question:]() |

## TTF

### TTF

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [TTF_Version](https://wiki.libsdl.org/SDL3_ttf/TTF_Version) | [:heavy_check_mark:](ttf/functions.go#L11) | [:x:](ttf/ttf_functions_js.go#L13) |
| [TTF_GetFreeTypeVersion](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFreeTypeVersion) | [:heavy_check_mark:](ttf/functions.go#L17) | [:x:](ttf/ttf_functions_js.go#L24) |
| [TTF_GetHarfBuzzVersion](https://wiki.libsdl.org/SDL3_ttf/TTF_GetHarfBuzzVersion) | [:heavy_check_mark:](ttf/functions.go#L27) | [:x:](ttf/ttf_functions_js.go#L48) |
| [TTF_Init](https://wiki.libsdl.org/SDL3_ttf/TTF_Init) | [:heavy_check_mark:](ttf/functions.go#L37) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L72) |
| [TTF_OpenFont](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFont) | [:heavy_check_mark:](ttf/functions.go#L47) | [:x:](ttf/ttf_functions_js.go#L80) |
| [TTF_OpenFontIO](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontIO) | [:heavy_check_mark:](ttf/functions.go#L58) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L96) |
| [TTF_OpenFontWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontWithProperties) | [:heavy_check_mark:](ttf/functions.go#L69) | [:x:](ttf/ttf_functions_js.go#L114) |
| [TTF_CopyFont](https://wiki.libsdl.org/SDL3_ttf/TTF_CopyFont) | [:heavy_check_mark:](ttf/methods.go#L409) | [:x:](ttf/ttf_functions_js.go#L131) |
| [TTF_GetFontProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontProperties) | [:heavy_check_mark:](ttf/methods.go#L420) | [:x:](ttf/ttf_functions_js.go#L148) |
| [TTF_GetFontGeneration](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontGeneration) | [:heavy_check_mark:](ttf/methods.go#L426) | [:x:](ttf/ttf_functions_js.go#L164) |
| [TTF_AddFallbackFont](https://wiki.libsdl.org/SDL3_ttf/TTF_AddFallbackFont) | [:heavy_check_mark:](ttf/methods.go#L432) | [:x:](ttf/ttf_functions_js.go#L180) |
| [TTF_RemoveFallbackFont](https://wiki.libsdl.org/SDL3_ttf/TTF_RemoveFallbackFont) | [:heavy_check_mark:](ttf/methods.go#L442) | [:x:](ttf/ttf_functions_js.go#L201) |
| [TTF_ClearFallbackFonts](https://wiki.libsdl.org/SDL3_ttf/TTF_ClearFallbackFonts) | [:heavy_check_mark:](ttf/methods.go#L448) | [:x:](ttf/ttf_functions_js.go#L220) |
| [TTF_SetFontSize](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSize) | [:heavy_check_mark:](ttf/methods.go#L454) | [:x:](ttf/ttf_functions_js.go#L234) |
| [TTF_SetFontSizeDPI](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSizeDPI) | [:heavy_check_mark:](ttf/methods.go#L464) | [:x:](ttf/ttf_functions_js.go#L252) |
| [TTF_GetFontSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSize) | [:heavy_check_mark:](ttf/methods.go#L474) | [:x:](ttf/ttf_functions_js.go#L274) |
| [TTF_GetFontDPI](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDPI) | [:heavy_check_mark:](ttf/methods.go#L485) | [:x:](ttf/ttf_functions_js.go#L290) |
| [TTF_SetFontStyle](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontStyle) | [:heavy_check_mark:](ttf/methods.go#L497) | [:x:](ttf/ttf_functions_js.go#L316) |
| [TTF_GetFontStyle](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyle) | [:heavy_check_mark:](ttf/methods.go#L503) | [:x:](ttf/ttf_functions_js.go#L332) |
| [TTF_SetFontOutline](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontOutline) | [:heavy_check_mark:](ttf/methods.go#L509) | [:x:](ttf/ttf_functions_js.go#L348) |
| [TTF_GetFontOutline](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontOutline) | [:heavy_check_mark:](ttf/methods.go#L519) | [:x:](ttf/ttf_functions_js.go#L366) |
| [TTF_SetFontHinting](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontHinting) | [:heavy_check_mark:](ttf/methods.go#L525) | [:x:](ttf/ttf_functions_js.go#L382) |
| [TTF_GetNumFontFaces](https://wiki.libsdl.org/SDL3_ttf/TTF_GetNumFontFaces) | [:heavy_check_mark:](ttf/methods.go#L531) | [:x:](ttf/ttf_functions_js.go#L398) |
| [TTF_GetFontHinting](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHinting) | [:heavy_check_mark:](ttf/methods.go#L537) | [:x:](ttf/ttf_functions_js.go#L414) |
| [TTF_SetFontSDF](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSDF) | [:heavy_check_mark:](ttf/methods.go#L543) | [:x:](ttf/ttf_functions_js.go#L430) |
| [TTF_GetFontSDF](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSDF) | [:heavy_check_mark:](ttf/methods.go#L553) | [:x:](ttf/ttf_functions_js.go#L448) |
| [TTF_GetFontWeight](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontWeight) | [:question:]() | [:question:]() |
| [TTF_SetFontWrapAlignment](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontWrapAlignment) | [:heavy_check_mark:](ttf/methods.go#L559) | [:x:](ttf/ttf_functions_js.go#L464) |
| [TTF_GetFontWrapAlignment](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontWrapAlignment) | [:heavy_check_mark:](ttf/methods.go#L565) | [:x:](ttf/ttf_functions_js.go#L480) |
| [TTF_GetFontHeight](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHeight) | [:heavy_check_mark:](ttf/methods.go#L571) | [:x:](ttf/ttf_functions_js.go#L496) |
| [TTF_GetFontAscent](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontAscent) | [:heavy_check_mark:](ttf/methods.go#L577) | [:x:](ttf/ttf_functions_js.go#L512) |
| [TTF_GetFontDescent](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDescent) | [:heavy_check_mark:](ttf/methods.go#L583) | [:x:](ttf/ttf_functions_js.go#L528) |
| [TTF_SetFontLineSkip](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLineSkip) | [:heavy_check_mark:](ttf/methods.go#L589) | [:x:](ttf/ttf_functions_js.go#L544) |
| [TTF_GetFontLineSkip](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontLineSkip) | [:heavy_check_mark:](ttf/methods.go#L595) | [:x:](ttf/ttf_functions_js.go#L560) |
| [TTF_SetFontKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontKerning) | [:heavy_check_mark:](ttf/methods.go#L601) | [:x:](ttf/ttf_functions_js.go#L576) |
| [TTF_GetFontKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontKerning) | [:heavy_check_mark:](ttf/methods.go#L607) | [:x:](ttf/ttf_functions_js.go#L592) |
| [TTF_FontIsFixedWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsFixedWidth) | [:heavy_check_mark:](ttf/methods.go#L613) | [:x:](ttf/ttf_functions_js.go#L608) |
| [TTF_FontIsScalable](https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsScalable) | [:heavy_check_mark:](ttf/methods.go#L619) | [:x:](ttf/ttf_functions_js.go#L624) |
| [TTF_GetFontFamilyName](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontFamilyName) | [:heavy_check_mark:](ttf/methods.go#L625) | [:x:](ttf/ttf_functions_js.go#L640) |
| [TTF_GetFontStyleName](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyleName) | [:heavy_check_mark:](ttf/methods.go#L631) | [:x:](ttf/ttf_functions_js.go#L656) |
| [TTF_SetFontDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontDirection) | [:heavy_check_mark:](ttf/methods.go#L637) | [:x:](ttf/ttf_functions_js.go#L672) |
| [TTF_GetFontDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDirection) | [:heavy_check_mark:](ttf/methods.go#L647) | [:x:](ttf/ttf_functions_js.go#L690) |
| [TTF_StringToTag](https://wiki.libsdl.org/SDL3_ttf/TTF_StringToTag) | [:heavy_check_mark:](ttf/functions.go#L80) | [:x:](ttf/ttf_functions_js.go#L706) |
| [TTF_TagToString](https://wiki.libsdl.org/SDL3_ttf/TTF_TagToString) | [:heavy_check_mark:](ttf/functions.go#L86) | [:x:](ttf/ttf_functions_js.go#L1958) |
| [TTF_SetFontScript](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontScript) | [:heavy_check_mark:](ttf/methods.go#L653) | [:x:](ttf/ttf_functions_js.go#L719) |
| [TTF_GetFontScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontScript) | [:heavy_check_mark:](ttf/methods.go#L663) | [:x:](ttf/ttf_functions_js.go#L737) |
| [TTF_GetGlyphScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphScript) | [:heavy_check_mark:](ttf/functions.go#L100) | [:x:](ttf/ttf_functions_js.go#L753) |
| [TTF_SetFontLanguage](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLanguage) | [:heavy_check_mark:](ttf/methods.go#L669) | [:x:](ttf/ttf_functions_js.go#L766) |
| [TTF_FontHasGlyph](https://wiki.libsdl.org/SDL3_ttf/TTF_FontHasGlyph) | [:heavy_check_mark:](ttf/methods.go#L679) | [:x:](ttf/ttf_functions_js.go#L784) |
| [TTF_GetGlyphImage](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImage) | [:heavy_check_mark:](ttf/methods.go#L685) | [:x:](ttf/ttf_functions_js.go#L802) |
| [TTF_GetGlyphImageForIndex](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImageForIndex) | [:heavy_check_mark:](ttf/methods.go#L698) | [:x:](ttf/ttf_functions_js.go#L826) |
| [TTF_GetGlyphMetrics](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphMetrics) | [:heavy_check_mark:](ttf/methods.go#L711) | [:x:](ttf/ttf_functions_js.go#L850) |
| [TTF_GetGlyphKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphKerning) | [:heavy_check_mark:](ttf/methods.go#L723) | [:x:](ttf/ttf_functions_js.go#L893) |
| [TTF_GetStringSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSize) | [:heavy_check_mark:](ttf/methods.go#L735) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L918) |
| [TTF_GetStringSizeWrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSizeWrapped) | [:heavy_check_mark:](ttf/methods.go#L747) | [:x:](ttf/ttf_functions_js.go#L943) |
| [TTF_MeasureString](https://wiki.libsdl.org/SDL3_ttf/TTF_MeasureString) | [:heavy_check_mark:](ttf/methods.go#L760) | [:x:](ttf/ttf_functions_js.go#L975) |
| [TTF_RenderText_Solid](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid) | [:heavy_check_mark:](ttf/methods.go#L773) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1976) |
| [TTF_RenderText_Solid_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid_Wrapped) | [:heavy_check_mark:](ttf/methods.go#L784) | [:x:](ttf/ttf_functions_js.go#L1998) |
| [TTF_RenderGlyph_Solid](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Solid) | [:heavy_check_mark:](ttf/methods.go#L795) | [:x:](ttf/ttf_functions_js.go#L2023) |
| [TTF_RenderText_Shaded](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded) | [:heavy_check_mark:](ttf/methods.go#L806) | [:x:](ttf/ttf_functions_js.go#L2042) |
| [TTF_RenderText_Shaded_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded_Wrapped) | [:heavy_check_mark:](ttf/methods.go#L817) | [:x:](ttf/ttf_functions_js.go#L2065) |
| [TTF_RenderGlyph_Shaded](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Shaded) | [:heavy_check_mark:](ttf/methods.go#L828) | [:x:](ttf/ttf_functions_js.go#L2090) |
| [TTF_RenderText_Blended](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended) | [:heavy_check_mark:](ttf/methods.go#L839) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L2111) |
| [TTF_RenderText_Blended_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended_Wrapped) | [:heavy_check_mark:](ttf/methods.go#L850) | [:x:](ttf/ttf_functions_js.go#L2133) |
| [TTF_RenderGlyph_Blended](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Blended) | [:heavy_check_mark:](ttf/methods.go#L861) | [:x:](ttf/ttf_functions_js.go#L2158) |
| [TTF_RenderText_LCD](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD) | [:heavy_check_mark:](ttf/methods.go#L872) | [:x:](ttf/ttf_functions_js.go#L2177) |
| [TTF_RenderText_LCD_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD_Wrapped) | [:heavy_check_mark:](ttf/methods.go#L883) | [:x:](ttf/ttf_functions_js.go#L2200) |
| [TTF_RenderGlyph_LCD](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_LCD) | [:heavy_check_mark:](ttf/methods.go#L894) | [:x:](ttf/ttf_functions_js.go#L2225) |
| [TTF_CreateSurfaceTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateSurfaceTextEngine) | [:heavy_check_mark:](ttf/functions.go#L111) | [:x:](ttf/ttf_functions_js.go#L1007) |
| [TTF_DrawSurfaceText](https://wiki.libsdl.org/SDL3_ttf/TTF_DrawSurfaceText) | [:heavy_check_mark:](ttf/methods.go#L13) | [:x:](ttf/ttf_functions_js.go#L1019) |
| [TTF_DestroySurfaceTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroySurfaceTextEngine) | [:heavy_check_mark:](ttf/methods.go#L366) | [:x:](ttf/ttf_functions_js.go#L1044) |
| [TTF_CreateRendererTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngine) | [:heavy_check_mark:](ttf/functions.go#L122) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1058) |
| [TTF_CreateRendererTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngineWithProperties) | [:heavy_check_mark:](ttf/functions.go#L133) | [:x:](ttf/ttf_functions_js.go#L1072) |
| [TTF_DrawRendererText](https://wiki.libsdl.org/SDL3_ttf/TTF_DrawRendererText) | [:heavy_check_mark:](ttf/methods.go#L23) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1089) |
| [TTF_DestroyRendererTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyRendererTextEngine) | [:heavy_check_mark:](ttf/methods.go#L372) | [:x:](ttf/ttf_functions_js.go#L1106) |
| [TTF_CreateGPUTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngine) | [:question:]() | [:question:](ttf/ttf_functions_js.go#L1120) |
| [TTF_CreateGPUTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngineWithProperties) | [:question:]() | [:question:](ttf/ttf_functions_js.go#L1137) |
| [TTF_GetGPUTextDrawData](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextDrawData) | [:heavy_check_mark:](ttf/methods.go#L33) | [:x:](ttf/ttf_functions_js.go#L1154) |
| [TTF_DestroyGPUTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyGPUTextEngine) | [:heavy_check_mark:](ttf/methods.go#L378) | [:x:](ttf/ttf_functions_js.go#L1171) |
| [TTF_SetGPUTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_SetGPUTextEngineWinding) | [:heavy_check_mark:](ttf/methods.go#L384) | [:x:](ttf/ttf_functions_js.go#L1185) |
| [TTF_GetGPUTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextEngineWinding) | [:heavy_check_mark:](ttf/methods.go#L390) | [:x:](ttf/ttf_functions_js.go#L1201) |
| [TTF_CreateText](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateText) | [:heavy_check_mark:](ttf/methods.go#L396) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1217) |
| [TTF_GetTextProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextProperties) | [:heavy_check_mark:](ttf/methods.go#L44) | [:x:](ttf/ttf_functions_js.go#L1243) |
| [TTF_SetTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextEngine) | [:heavy_check_mark:](ttf/methods.go#L50) | [:x:](ttf/ttf_functions_js.go#L1259) |
| [TTF_GetTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextEngine) | [:heavy_check_mark:](ttf/methods.go#L60) | [:x:](ttf/ttf_functions_js.go#L1280) |
| [TTF_SetTextFont](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextFont) | [:heavy_check_mark:](ttf/methods.go#L71) | [:x:](ttf/ttf_functions_js.go#L1297) |
| [TTF_GetTextFont](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextFont) | [:heavy_check_mark:](ttf/methods.go#L81) | [:x:](ttf/ttf_functions_js.go#L1318) |
| [TTF_SetTextDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextDirection) | [:heavy_check_mark:](ttf/methods.go#L92) | [:x:](ttf/ttf_functions_js.go#L1335) |
| [TTF_GetTextDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextDirection) | [:heavy_check_mark:](ttf/methods.go#L102) | [:x:](ttf/ttf_functions_js.go#L1353) |
| [TTF_SetTextScript](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextScript) | [:heavy_check_mark:](ttf/methods.go#L108) | [:x:](ttf/ttf_functions_js.go#L1369) |
| [TTF_GetTextScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextScript) | [:heavy_check_mark:](ttf/methods.go#L118) | [:x:](ttf/ttf_functions_js.go#L1387) |
| [TTF_SetTextColor](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColor) | [:heavy_check_mark:](ttf/methods.go#L124) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1403) |
| [TTF_SetTextColorFloat](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColorFloat) | [:heavy_check_mark:](ttf/methods.go#L134) | [:x:](ttf/ttf_functions_js.go#L1424) |
| [TTF_GetTextColor](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColor) | [:heavy_check_mark:](ttf/methods.go#L144) | [:x:](ttf/ttf_functions_js.go#L1448) |
| [TTF_GetTextColorFloat](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColorFloat) | [:heavy_check_mark:](ttf/methods.go#L156) | [:x:](ttf/ttf_functions_js.go#L1484) |
| [TTF_SetTextPosition](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextPosition) | [:heavy_check_mark:](ttf/methods.go#L168) | [:x:](ttf/ttf_functions_js.go#L1520) |
| [TTF_GetTextPosition](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextPosition) | [:heavy_check_mark:](ttf/methods.go#L174) | [:x:](ttf/ttf_functions_js.go#L1540) |
| [TTF_SetTextWrapWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWidth) | [:heavy_check_mark:](ttf/methods.go#L184) | [:x:](ttf/ttf_functions_js.go#L1566) |
| [TTF_GetTextWrapWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextWrapWidth) | [:heavy_check_mark:](ttf/methods.go#L194) | [:x:](ttf/ttf_functions_js.go#L1584) |
| [TTF_SetTextWrapWhitespaceVisible](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWhitespaceVisible) | [:heavy_check_mark:](ttf/methods.go#L206) | [:x:](ttf/ttf_functions_js.go#L1605) |
| [TTF_TextWrapWhitespaceVisible](https://wiki.libsdl.org/SDL3_ttf/TTF_TextWrapWhitespaceVisible) | [:heavy_check_mark:](ttf/methods.go#L216) | [:x:](ttf/ttf_functions_js.go#L1623) |
| [TTF_SetTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextString) | [:heavy_check_mark:](ttf/methods.go#L222) | [:x:](ttf/ttf_functions_js.go#L1639) |
| [TTF_InsertTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_InsertTextString) | [:heavy_check_mark:](ttf/methods.go#L232) | [:x:](ttf/ttf_functions_js.go#L1659) |
| [TTF_AppendTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_AppendTextString) | [:heavy_check_mark:](ttf/methods.go#L242) | [:x:](ttf/ttf_functions_js.go#L1681) |
| [TTF_DeleteTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_DeleteTextString) | [:heavy_check_mark:](ttf/methods.go#L252) | [:x:](ttf/ttf_functions_js.go#L1701) |
| [TTF_GetTextSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSize) | [:heavy_check_mark:](ttf/methods.go#L262) | [:x:](ttf/ttf_functions_js.go#L1721) |
| [TTF_GetTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubString) | [:heavy_check_mark:](ttf/methods.go#L274) | [:x:](ttf/ttf_functions_js.go#L1747) |
| [TTF_GetTextSubStringForLine](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForLine) | [:heavy_check_mark:](ttf/methods.go#L286) | [:x:](ttf/ttf_functions_js.go#L1770) |
| [TTF_GetTextSubStringsForRange](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringsForRange) | [:heavy_check_mark:](ttf/methods.go#L298) | [:x:](ttf/ttf_functions_js.go#L1793) |
| [TTF_GetTextSubStringForPoint](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForPoint) | [:heavy_check_mark:](ttf/methods.go#L312) | [:x:](ttf/ttf_functions_js.go#L1819) |
| [TTF_GetPreviousTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetPreviousTextSubString) | [:heavy_check_mark:](ttf/methods.go#L324) | [:x:](ttf/ttf_functions_js.go#L1844) |
| [TTF_GetNextTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetNextTextSubString) | [:heavy_check_mark:](ttf/methods.go#L336) | [:x:](ttf/ttf_functions_js.go#L1870) |
| [TTF_UpdateText](https://wiki.libsdl.org/SDL3_ttf/TTF_UpdateText) | [:heavy_check_mark:](ttf/methods.go#L348) | [:x:](ttf/ttf_functions_js.go#L1896) |
| [TTF_DestroyText](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyText) | [:heavy_check_mark:](ttf/methods.go#L358) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1912) |
| [TTF_CloseFont](https://wiki.libsdl.org/SDL3_ttf/TTF_CloseFont) | [:heavy_check_mark:](ttf/methods.go#L905) | [:x:](ttf/ttf_functions_js.go#L1924) |
| [TTF_Quit](https://wiki.libsdl.org/SDL3_ttf/TTF_Quit) | [:heavy_check_mark:](ttf/functions.go#L144) | [:x:](ttf/ttf_functions_js.go#L1938) |
| [TTF_WasInit](https://wiki.libsdl.org/SDL3_ttf/TTF_WasInit) | [:heavy_check_mark:](ttf/functions.go#L150) | [:x:](ttf/ttf_functions_js.go#L1947) |

## IMG

### Image

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [IMG_Version](https://wiki.libsdl.org/SDL3_image/IMG_Version) | [:heavy_check_mark:](img/functions.go#L11) | [:x:](img/img_functions_js.go#L13) |
| [IMG_LoadTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTyped_IO) | [:heavy_check_mark:](img/functions.go#L17) | [:x:](img/img_functions_js.go#L24) |
| [IMG_Load](https://wiki.libsdl.org/SDL3_image/IMG_Load) | [:heavy_check_mark:](img/functions.go#L28) | [:x:](img/img_functions_js.go#L45) |
| [IMG_Load_IO](https://wiki.libsdl.org/SDL3_image/IMG_Load_IO) | [:heavy_check_mark:](img/functions.go#L39) | [:x:](img/img_functions_js.go#L59) |
| [IMG_LoadTexture](https://wiki.libsdl.org/SDL3_image/IMG_LoadTexture) | [:heavy_check_mark:](img/functions.go#L50) | [:x:](img/img_functions_js.go#L78) |
| [IMG_LoadTexture_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTexture_IO) | [:heavy_check_mark:](img/functions.go#L61) | [:x:](img/img_functions_js.go#L97) |
| [IMG_LoadTextureTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTextureTyped_IO) | [:heavy_check_mark:](img/functions.go#L72) | [:x:](img/img_functions_js.go#L121) |
| [IMG_isAVIF](https://wiki.libsdl.org/SDL3_image/IMG_isAVIF) | [:question:]() | [:question:]() |
| [IMG_isICO](https://wiki.libsdl.org/SDL3_image/IMG_isICO) | [:question:]() | [:question:]() |
| [IMG_isCUR](https://wiki.libsdl.org/SDL3_image/IMG_isCUR) | [:question:]() | [:question:]() |
| [IMG_isBMP](https://wiki.libsdl.org/SDL3_image/IMG_isBMP) | [:question:]() | [:question:]() |
| [IMG_isGIF](https://wiki.libsdl.org/SDL3_image/IMG_isGIF) | [:question:]() | [:question:]() |
| [IMG_isJPG](https://wiki.libsdl.org/SDL3_image/IMG_isJPG) | [:question:]() | [:question:]() |
| [IMG_isJXL](https://wiki.libsdl.org/SDL3_image/IMG_isJXL) | [:question:]() | [:question:]() |
| [IMG_isLBM](https://wiki.libsdl.org/SDL3_image/IMG_isLBM) | [:question:]() | [:question:]() |
| [IMG_isPCX](https://wiki.libsdl.org/SDL3_image/IMG_isPCX) | [:question:]() | [:question:]() |
| [IMG_isPNG](https://wiki.libsdl.org/SDL3_image/IMG_isPNG) | [:question:]() | [:question:]() |
| [IMG_isPNM](https://wiki.libsdl.org/SDL3_image/IMG_isPNM) | [:question:]() | [:question:]() |
| [IMG_isSVG](https://wiki.libsdl.org/SDL3_image/IMG_isSVG) | [:question:]() | [:question:]() |
| [IMG_isQOI](https://wiki.libsdl.org/SDL3_image/IMG_isQOI) | [:question:]() | [:question:]() |
| [IMG_isTIF](https://wiki.libsdl.org/SDL3_image/IMG_isTIF) | [:question:]() | [:question:]() |
| [IMG_isXCF](https://wiki.libsdl.org/SDL3_image/IMG_isXCF) | [:question:]() | [:question:]() |
| [IMG_isXPM](https://wiki.libsdl.org/SDL3_image/IMG_isXPM) | [:question:]() | [:question:]() |
| [IMG_isXV](https://wiki.libsdl.org/SDL3_image/IMG_isXV) | [:question:]() | [:question:]() |
| [IMG_isWEBP](https://wiki.libsdl.org/SDL3_image/IMG_isWEBP) | [:question:]() | [:question:]() |
| [IMG_LoadAVIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAVIF_IO) | [:heavy_check_mark:](img/functions.go#L191) | [:x:](img/img_functions_js.go#L147) |
| [IMG_LoadICO_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadICO_IO) | [:heavy_check_mark:](img/functions.go#L202) | [:x:](img/img_functions_js.go#L164) |
| [IMG_LoadCUR_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadCUR_IO) | [:heavy_check_mark:](img/functions.go#L213) | [:x:](img/img_functions_js.go#L181) |
| [IMG_LoadBMP_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadBMP_IO) | [:heavy_check_mark:](img/functions.go#L224) | [:x:](img/img_functions_js.go#L198) |
| [IMG_LoadGIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadGIF_IO) | [:heavy_check_mark:](img/functions.go#L235) | [:x:](img/img_functions_js.go#L215) |
| [IMG_LoadJPG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadJPG_IO) | [:heavy_check_mark:](img/functions.go#L246) | [:x:](img/img_functions_js.go#L232) |
| [IMG_LoadJXL_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadJXL_IO) | [:heavy_check_mark:](img/functions.go#L257) | [:x:](img/img_functions_js.go#L249) |
| [IMG_LoadLBM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadLBM_IO) | [:heavy_check_mark:](img/functions.go#L268) | [:x:](img/img_functions_js.go#L266) |
| [IMG_LoadPCX_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPCX_IO) | [:heavy_check_mark:](img/functions.go#L279) | [:x:](img/img_functions_js.go#L283) |
| [IMG_LoadPNG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPNG_IO) | [:heavy_check_mark:](img/functions.go#L290) | [:x:](img/img_functions_js.go#L300) |
| [IMG_LoadPNM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPNM_IO) | [:heavy_check_mark:](img/functions.go#L301) | [:x:](img/img_functions_js.go#L317) |
| [IMG_LoadSVG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadSVG_IO) | [:heavy_check_mark:](img/functions.go#L312) | [:x:](img/img_functions_js.go#L334) |
| [IMG_LoadQOI_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadQOI_IO) | [:heavy_check_mark:](img/functions.go#L323) | [:x:](img/img_functions_js.go#L351) |
| [IMG_LoadTGA_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTGA_IO) | [:heavy_check_mark:](img/functions.go#L334) | [:x:](img/img_functions_js.go#L368) |
| [IMG_LoadTIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTIF_IO) | [:heavy_check_mark:](img/functions.go#L345) | [:x:](img/img_functions_js.go#L385) |
| [IMG_LoadXCF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXCF_IO) | [:heavy_check_mark:](img/functions.go#L356) | [:x:](img/img_functions_js.go#L402) |
| [IMG_LoadXPM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXPM_IO) | [:heavy_check_mark:](img/functions.go#L367) | [:x:](img/img_functions_js.go#L419) |
| [IMG_LoadXV_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXV_IO) | [:heavy_check_mark:](img/functions.go#L378) | [:x:](img/img_functions_js.go#L436) |
| [IMG_LoadWEBP_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadWEBP_IO) | [:heavy_check_mark:](img/functions.go#L389) | [:x:](img/img_functions_js.go#L453) |
| [IMG_LoadSizedSVG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadSizedSVG_IO) | [:heavy_check_mark:](img/functions.go#L400) | [:x:](img/img_functions_js.go#L470) |
| [IMG_ReadXPMFromArray](https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArray) | [:heavy_check_mark:](img/functions.go#L411) | [:x:](img/img_functions_js.go#L491) |
| [IMG_ReadXPMFromArrayToRGB888](https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArrayToRGB888) | [:heavy_check_mark:](img/functions.go#L423) | [:x:](img/img_functions_js.go#L508) |
| [IMG_SaveAVIF](https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF) | [:heavy_check_mark:](img/functions.go#L435) | [:x:](img/img_functions_js.go#L525) |
| [IMG_SaveAVIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF_IO) | [:heavy_check_mark:](img/functions.go#L445) | [:x:](img/img_functions_js.go#L545) |
| [IMG_SavePNG](https://wiki.libsdl.org/SDL3_image/IMG_SavePNG) | [:heavy_check_mark:](img/functions.go#L455) | [:x:](img/img_functions_js.go#L570) |
| [IMG_SavePNG_IO](https://wiki.libsdl.org/SDL3_image/IMG_SavePNG_IO) | [:heavy_check_mark:](img/functions.go#L465) | [:x:](img/img_functions_js.go#L588) |
| [IMG_SaveJPG](https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG) | [:heavy_check_mark:](img/functions.go#L475) | [:x:](img/img_functions_js.go#L611) |
| [IMG_SaveJPG_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG_IO) | [:heavy_check_mark:](img/functions.go#L485) | [:x:](img/img_functions_js.go#L631) |
| [IMG_LoadAnimation](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation) | [:heavy_check_mark:](img/functions.go#L495) | [:x:](img/img_functions_js.go#L656) |
| [IMG_LoadAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation_IO) | [:heavy_check_mark:](img/functions.go#L506) | [:x:](img/img_functions_js.go#L670) |
| [IMG_LoadAnimationTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimationTyped_IO) | [:heavy_check_mark:](img/functions.go#L517) | [:x:](img/img_functions_js.go#L689) |
| [IMG_FreeAnimation](https://wiki.libsdl.org/SDL3_image/IMG_FreeAnimation) | [:heavy_check_mark:](img/methods.go#L6) | [:x:](img/img_functions_js.go#L710) |
| [IMG_LoadGIFAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadGIFAnimation_IO) | [:heavy_check_mark:](img/functions.go#L528) | [:x:](img/img_functions_js.go#L724) |
| [IMG_LoadWEBPAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadWEBPAnimation_IO) | [:heavy_check_mark:](img/functions.go#L539) | [:x:](img/img_functions_js.go#L741) |

## MIXER

### Mixer

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [Mix_Version](https://wiki.libsdl.org/SDL3_mixer/Mix_Version) | [:heavy_check_mark:](mixer/functions.go#L11) | [:x:](mixer/mixer_functions_js.go#L13) |
| [Mix_Init](https://wiki.libsdl.org/SDL3_mixer/Mix_Init) | [:x:](mixer/methods.go#L35) | [:x:](mixer/mixer_functions_js.go#L24) |
| [Mix_Quit](https://wiki.libsdl.org/SDL3_mixer/Mix_Quit) | [:heavy_check_mark:](mixer/functions.go#L23) | [:x:](mixer/mixer_functions_js.go#L37) |
| [Mix_OpenAudio](https://wiki.libsdl.org/SDL3_mixer/Mix_OpenAudio) | [:heavy_check_mark:](mixer/functions.go#L29) | [:x:](mixer/mixer_functions_js.go#L46) |
| [Mix_PauseAudio](https://wiki.libsdl.org/SDL3_mixer/Mix_PauseAudio) | [:heavy_check_mark:](mixer/functions.go#L44) | [:x:](mixer/mixer_functions_js.go#L67) |
| [Mix_QuerySpec](https://wiki.libsdl.org/SDL3_mixer/Mix_QuerySpec) | [:heavy_check_mark:](mixer/functions.go#L50) | [:x:](mixer/mixer_functions_js.go#L78) |
| [Mix_AllocateChannels](https://wiki.libsdl.org/SDL3_mixer/Mix_AllocateChannels) | [:heavy_check_mark:](mixer/functions.go#L62) | [:x:](mixer/mixer_functions_js.go#L104) |
| [Mix_LoadWAV_IO](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadWAV_IO) | [:heavy_check_mark:](mixer/functions.go#L68) | [:x:](mixer/mixer_functions_js.go#L117) |
| [Mix_LoadWAV](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadWAV) | [:heavy_check_mark:](mixer/functions.go#L79) | [:x:](mixer/mixer_functions_js.go#L136) |
| [Mix_LoadMUS](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadMUS) | [:heavy_check_mark:](mixer/functions.go#L101) | [:x:](mixer/mixer_functions_js.go#L150) |
| [Mix_LoadMUS_IO](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadMUS_IO) | [:heavy_check_mark:](mixer/functions.go#L90) | [:x:](mixer/mixer_functions_js.go#L164) |
| [Mix_LoadMUSType_IO](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadMUSType_IO) | [:heavy_check_mark:](mixer/functions.go#L112) | [:x:](mixer/mixer_functions_js.go#L183) |
| [Mix_QuickLoad_WAV](https://wiki.libsdl.org/SDL3_mixer/Mix_QuickLoad_WAV) | [:heavy_check_mark:](mixer/functions.go#L123) | [:x:](mixer/mixer_functions_js.go#L204) |
| [Mix_QuickLoad_RAW](https://wiki.libsdl.org/SDL3_mixer/Mix_QuickLoad_RAW) | [:heavy_check_mark:](mixer/functions.go#L134) | [:x:](mixer/mixer_functions_js.go#L221) |
| [Mix_FreeChunk](https://wiki.libsdl.org/SDL3_mixer/Mix_FreeChunk) | [:heavy_check_mark:](mixer/methods.go#L44) | [:x:](mixer/mixer_functions_js.go#L240) |
| [Mix_FreeMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_FreeMusic) | [:heavy_check_mark:](mixer/methods.go#L57) | [:x:](mixer/mixer_functions_js.go#L254) |
| [Mix_GetNumChunkDecoders](https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumChunkDecoders) | [:heavy_check_mark:](mixer/functions.go#L145) | [:x:](mixer/mixer_functions_js.go#L268) |
| [Mix_GetChunkDecoder](https://wiki.libsdl.org/SDL3_mixer/Mix_GetChunkDecoder) | [:heavy_check_mark:](mixer/functions.go#L151) | [:x:](mixer/mixer_functions_js.go#L279) |
| [Mix_HasChunkDecoder](https://wiki.libsdl.org/SDL3_mixer/Mix_HasChunkDecoder) | [:heavy_check_mark:](mixer/functions.go#L157) | [:x:](mixer/mixer_functions_js.go#L292) |
| [Mix_GetNumMusicDecoders](https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumMusicDecoders) | [:heavy_check_mark:](mixer/functions.go#L163) | [:x:](mixer/mixer_functions_js.go#L305) |
| [Mix_GetMusicDecoder](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicDecoder) | [:heavy_check_mark:](mixer/functions.go#L169) | [:x:](mixer/mixer_functions_js.go#L316) |
| [Mix_HasMusicDecoder](https://wiki.libsdl.org/SDL3_mixer/Mix_HasMusicDecoder) | [:heavy_check_mark:](mixer/functions.go#L175) | [:x:](mixer/mixer_functions_js.go#L329) |
| [Mix_GetMusicType](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicType) | [:x:](mixer/methods.go#L63) | [:x:](mixer/mixer_functions_js.go#L342) |
| [Mix_GetMusicTitle](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicTitle) | [:x:](mixer/methods.go#L70) | [:x:](mixer/mixer_functions_js.go#L358) |
| [Mix_GetMusicTitleTag](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicTitleTag) | [:x:](mixer/methods.go#L77) | [:x:](mixer/mixer_functions_js.go#L374) |
| [Mix_GetMusicArtistTag](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicArtistTag) | [:x:](mixer/methods.go#L84) | [:x:](mixer/mixer_functions_js.go#L390) |
| [Mix_GetMusicAlbumTag](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicAlbumTag) | [:x:](mixer/methods.go#L91) | [:x:](mixer/mixer_functions_js.go#L406) |
| [Mix_GetMusicCopyrightTag](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicCopyrightTag) | [:x:](mixer/methods.go#L98) | [:x:](mixer/mixer_functions_js.go#L422) |
| [Mix_SetPostMix](https://wiki.libsdl.org/SDL3_mixer/Mix_SetPostMix) | [:x:](mixer/methods.go#L187) | [:x:](mixer/mixer_functions_js.go#L438) |
| [Mix_HookMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_HookMusic) | [:x:](mixer/methods.go#L194) | [:x:](mixer/mixer_functions_js.go#L451) |
| [Mix_HookMusicFinished](https://wiki.libsdl.org/SDL3_mixer/Mix_HookMusicFinished) | [:x:](mixer/methods.go#L8) | [:x:](mixer/mixer_functions_js.go#L464) |
| [Mix_GetMusicHookData](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicHookData) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L475) |
| [Mix_ChannelFinished](https://wiki.libsdl.org/SDL3_mixer/Mix_ChannelFinished) | [:x:](mixer/methods.go#L17) | [:x:](mixer/mixer_functions_js.go#L486) |
| [Mix_RegisterEffect](https://wiki.libsdl.org/SDL3_mixer/Mix_RegisterEffect) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L497) |
| [Mix_UnregisterEffect](https://wiki.libsdl.org/SDL3_mixer/Mix_UnregisterEffect) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L516) |
| [Mix_UnregisterAllEffects](https://wiki.libsdl.org/SDL3_mixer/Mix_UnregisterAllEffects) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L531) |
| [Mix_SetPanning](https://wiki.libsdl.org/SDL3_mixer/Mix_SetPanning) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L544) |
| [Mix_SetPosition](https://wiki.libsdl.org/SDL3_mixer/Mix_SetPosition) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L561) |
| [Mix_SetDistance](https://wiki.libsdl.org/SDL3_mixer/Mix_SetDistance) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L578) |
| [Mix_SetReverseStereo](https://wiki.libsdl.org/SDL3_mixer/Mix_SetReverseStereo) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L593) |
| [Mix_ReserveChannels](https://wiki.libsdl.org/SDL3_mixer/Mix_ReserveChannels) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L608) |
| [Mix_GroupChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupChannel) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L621) |
| [Mix_GroupChannels](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupChannels) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L636) |
| [Mix_GroupAvailable](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupAvailable) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L653) |
| [Mix_GroupCount](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupCount) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L666) |
| [Mix_GroupOldest](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupOldest) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L679) |
| [Mix_GroupNewer](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupNewer) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L692) |
| [Mix_PlayChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_PlayChannel) | [:heavy_check_mark:](mixer/functions.go#L181) | [:x:](mixer/mixer_functions_js.go#L705) |
| [Mix_PlayChannelTimed](https://wiki.libsdl.org/SDL3_mixer/Mix_PlayChannelTimed) | [:heavy_check_mark:](mixer/functions.go#L192) | [:x:](mixer/mixer_functions_js.go#L725) |
| [Mix_PlayMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_PlayMusic) | [:heavy_check_mark:](mixer/methods.go#L105) | [:x:](mixer/mixer_functions_js.go#L747) |
| [Mix_FadeInMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInMusic) | [:x:](mixer/methods.go#L115) | [:x:](mixer/mixer_functions_js.go#L765) |
| [Mix_FadeInMusicPos](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInMusicPos) | [:x:](mixer/methods.go#L122) | [:x:](mixer/mixer_functions_js.go#L785) |
| [Mix_FadeInChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInChannel) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L807) |
| [Mix_FadeInChannelTimed](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInChannelTimed) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L829) |
| [Mix_Volume](https://wiki.libsdl.org/SDL3_mixer/Mix_Volume) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L853) |
| [Mix_VolumeChunk](https://wiki.libsdl.org/SDL3_mixer/Mix_VolumeChunk) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L868) |
| [Mix_VolumeMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_VolumeMusic) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L886) |
| [Mix_GetMusicVolume](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicVolume) | [:x:](mixer/methods.go#L129) | [:x:](mixer/mixer_functions_js.go#L899) |
| [Mix_MasterVolume](https://wiki.libsdl.org/SDL3_mixer/Mix_MasterVolume) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L915) |
| [Mix_HaltChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_HaltChannel) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L928) |
| [Mix_HaltGroup](https://wiki.libsdl.org/SDL3_mixer/Mix_HaltGroup) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L939) |
| [Mix_HaltMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_HaltMusic) | [:heavy_check_mark:](mixer/functions.go#L227) | [:x:](mixer/mixer_functions_js.go#L950) |
| [Mix_ExpireChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_ExpireChannel) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L959) |
| [Mix_FadeOutChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeOutChannel) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L974) |
| [Mix_FadeOutGroup](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeOutGroup) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L989) |
| [Mix_FadeOutMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeOutMusic) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1004) |
| [Mix_FadingMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_FadingMusic) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1017) |
| [Mix_FadingChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_FadingChannel) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1028) |
| [Mix_Pause](https://wiki.libsdl.org/SDL3_mixer/Mix_Pause) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1041) |
| [Mix_PauseGroup](https://wiki.libsdl.org/SDL3_mixer/Mix_PauseGroup) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1052) |
| [Mix_Resume](https://wiki.libsdl.org/SDL3_mixer/Mix_Resume) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1063) |
| [Mix_ResumeGroup](https://wiki.libsdl.org/SDL3_mixer/Mix_ResumeGroup) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1074) |
| [Mix_Paused](https://wiki.libsdl.org/SDL3_mixer/Mix_Paused) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1085) |
| [Mix_PauseMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_PauseMusic) | [:heavy_check_mark:](mixer/functions.go#L215) | [:x:](mixer/mixer_functions_js.go#L1098) |
| [Mix_ResumeMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_ResumeMusic) | [:heavy_check_mark:](mixer/functions.go#L221) | [:x:](mixer/mixer_functions_js.go#L1107) |
| [Mix_RewindMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_RewindMusic) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1116) |
| [Mix_PausedMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_PausedMusic) | [:heavy_check_mark:](mixer/functions.go#L209) | [:x:](mixer/mixer_functions_js.go#L1125) |
| [Mix_ModMusicJumpToOrder](https://wiki.libsdl.org/SDL3_mixer/Mix_ModMusicJumpToOrder) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1136) |
| [Mix_StartTrack](https://wiki.libsdl.org/SDL3_mixer/Mix_StartTrack) | [:x:](mixer/methods.go#L136) | [:x:](mixer/mixer_functions_js.go#L1149) |
| [Mix_GetNumTracks](https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumTracks) | [:x:](mixer/methods.go#L143) | [:x:](mixer/mixer_functions_js.go#L1167) |
| [Mix_SetMusicPosition](https://wiki.libsdl.org/SDL3_mixer/Mix_SetMusicPosition) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1183) |
| [Mix_GetMusicPosition](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicPosition) | [:x:](mixer/methods.go#L150) | [:x:](mixer/mixer_functions_js.go#L1196) |
| [Mix_MusicDuration](https://wiki.libsdl.org/SDL3_mixer/Mix_MusicDuration) | [:x:](mixer/methods.go#L157) | [:x:](mixer/mixer_functions_js.go#L1212) |
| [Mix_GetMusicLoopStartTime](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopStartTime) | [:x:](mixer/methods.go#L164) | [:x:](mixer/mixer_functions_js.go#L1228) |
| [Mix_GetMusicLoopEndTime](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopEndTime) | [:x:](mixer/methods.go#L171) | [:x:](mixer/mixer_functions_js.go#L1244) |
| [Mix_GetMusicLoopLengthTime](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopLengthTime) | [:x:](mixer/methods.go#L178) | [:x:](mixer/mixer_functions_js.go#L1260) |
| [Mix_Playing](https://wiki.libsdl.org/SDL3_mixer/Mix_Playing) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1276) |
| [Mix_PlayingMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_PlayingMusic) | [:heavy_check_mark:](mixer/functions.go#L203) | [:x:](mixer/mixer_functions_js.go#L1289) |
| [Mix_SetSoundFonts](https://wiki.libsdl.org/SDL3_mixer/Mix_SetSoundFonts) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1300) |
| [Mix_GetSoundFonts](https://wiki.libsdl.org/SDL3_mixer/Mix_GetSoundFonts) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1313) |
| [Mix_EachSoundFont](https://wiki.libsdl.org/SDL3_mixer/Mix_EachSoundFont) | [:x:](mixer/methods.go#L26) | [:x:](mixer/mixer_functions_js.go#L1324) |
| [Mix_SetTimidityCfg](https://wiki.libsdl.org/SDL3_mixer/Mix_SetTimidityCfg) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1339) |
| [Mix_GetTimidityCfg](https://wiki.libsdl.org/SDL3_mixer/Mix_GetTimidityCfg) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1352) |
| [Mix_GetChunk](https://wiki.libsdl.org/SDL3_mixer/Mix_GetChunk) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1363) |
| [Mix_CloseAudio](https://wiki.libsdl.org/SDL3_mixer/Mix_CloseAudio) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1377) |

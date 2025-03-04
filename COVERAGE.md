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
| [SDL_Init](https://wiki.libsdl.org/SDL3/SDL_Init) | [:heavy_check_mark:](./functions.go#L13) | [:heavy_check_mark:](./sdl_functions_js.go#L13910) |
| [SDL_InitSubSystem](https://wiki.libsdl.org/SDL3/SDL_InitSubSystem) | [:heavy_check_mark:](./functions.go#L23) | [:x:](./sdl_functions_js.go#L13920) |
| [SDL_QuitSubSystem](https://wiki.libsdl.org/SDL3/SDL_QuitSubSystem) | [:heavy_check_mark:](./functions.go#L33) | [:x:](./sdl_functions_js.go#L13933) |
| [SDL_WasInit](https://wiki.libsdl.org/SDL3/SDL_WasInit) | [:heavy_check_mark:](./functions.go#L39) | [:x:](./sdl_functions_js.go#L13944) |
| [SDL_Quit](https://wiki.libsdl.org/SDL3/SDL_Quit) | [:heavy_check_mark:](./functions.go#L45) | [:heavy_check_mark:]() |
| [SDL_IsMainThread](https://wiki.libsdl.org/SDL3/SDL_IsMainThread) | [:question:]() | [:question:](./sdl_functions_js.go#L13961) |
| [SDL_RunOnMainThread](https://wiki.libsdl.org/SDL3/SDL_RunOnMainThread) | [:question:]() | [:question:](./sdl_functions_js.go#L13972) |
| [SDL_SetAppMetadata](https://wiki.libsdl.org/SDL3/SDL_SetAppMetadata) | [:question:]() | [:question:](./sdl_functions_js.go#L13989) |
| [SDL_SetAppMetadataProperty](https://wiki.libsdl.org/SDL3/SDL_SetAppMetadataProperty) | [:question:]() | [:question:](./sdl_functions_js.go#L14006) |
| [SDL_GetAppMetadataProperty](https://wiki.libsdl.org/SDL3/SDL_GetAppMetadataProperty) | [:question:]() | [:question:](./sdl_functions_js.go#L14021) |
### Hints

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetHintWithPriority](https://wiki.libsdl.org/SDL3/SDL_SetHintWithPriority) | [:heavy_check_mark:](./functions.go#L59) | [:x:](./sdl_functions_js.go#L13796) |
| [SDL_SetHint](https://wiki.libsdl.org/SDL3/SDL_SetHint) | [:heavy_check_mark:](./functions.go#L69) | [:x:](./sdl_functions_js.go#L13813) |
| [SDL_ResetHint](https://wiki.libsdl.org/SDL3/SDL_ResetHint) | [:heavy_check_mark:](./functions.go#L79) | [:x:](./sdl_functions_js.go#L13828) |
| [SDL_ResetHints](https://wiki.libsdl.org/SDL3/SDL_ResetHints) | [:heavy_check_mark:](./functions.go#L89) | [:x:](./sdl_functions_js.go#L13841) |
| [SDL_GetHint](https://wiki.libsdl.org/SDL3/SDL_GetHint) | [:heavy_check_mark:](./functions.go#L95) | [:x:](./sdl_functions_js.go#L13850) |
| [SDL_GetHintBoolean](https://wiki.libsdl.org/SDL3/SDL_GetHintBoolean) | [:heavy_check_mark:](./functions.go#L101) | [:x:](./sdl_functions_js.go#L13863) |
| [SDL_AddHintCallback](https://wiki.libsdl.org/SDL3/SDL_AddHintCallback) | [:question:]() | [:question:](./sdl_functions_js.go#L13878) |
| [SDL_RemoveHintCallback](https://wiki.libsdl.org/SDL3/SDL_RemoveHintCallback) | [:question:]() | [:question:](./sdl_functions_js.go#L13895) |
### Error

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetError](https://wiki.libsdl.org/SDL3/SDL_SetError) | [:question:]() | [:question:](./sdl_functions_js.go#L488) |
| [SDL_SetErrorV](https://wiki.libsdl.org/SDL3/SDL_SetErrorV) | [:question:]() | [:question:](./sdl_functions_js.go#L501) |
| [SDL_OutOfMemory](https://wiki.libsdl.org/SDL3/SDL_OutOfMemory) | [:question:]() | [:question:](./sdl_functions_js.go#L516) |
| [SDL_GetError](https://wiki.libsdl.org/SDL3/SDL_GetError) | [:heavy_check_mark:](./library_notjs.go#L13) | [:heavy_check_mark:](./sdl_functions_js.go#L527) |
| [SDL_ClearError](https://wiki.libsdl.org/SDL3/SDL_ClearError) | [:question:]() | [:question:](./sdl_functions_js.go#L535) |
### Properties

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetGlobalProperties](https://wiki.libsdl.org/SDL3/SDL_GetGlobalProperties) | [:heavy_check_mark:](./functions.go#L116) | [:x:](./sdl_functions_js.go#L546) |
| [SDL_CreateProperties](https://wiki.libsdl.org/SDL3/SDL_CreateProperties) | [:heavy_check_mark:](./functions.go#L127) | [:x:](./sdl_functions_js.go#L557) |
| [SDL_CopyProperties](https://wiki.libsdl.org/SDL3/SDL_CopyProperties) | [:x:](./methods.go#L5302) | [:x:](./sdl_functions_js.go#L568) |
| [SDL_LockProperties](https://wiki.libsdl.org/SDL3/SDL_LockProperties) | [:x:](./methods.go#L5309) | [:x:](./sdl_functions_js.go#L583) |
| [SDL_UnlockProperties](https://wiki.libsdl.org/SDL3/SDL_UnlockProperties) | [:x:](./methods.go#L5316) | [:x:](./sdl_functions_js.go#L596) |
| [SDL_SetPointerPropertyWithCleanup](https://wiki.libsdl.org/SDL3/SDL_SetPointerPropertyWithCleanup) | [:x:](./methods.go#L5323) | [:x:](./sdl_functions_js.go#L607) |
| [SDL_SetPointerProperty](https://wiki.libsdl.org/SDL3/SDL_SetPointerProperty) | [:x:](./methods.go#L5330) | [:x:](./sdl_functions_js.go#L628) |
| [SDL_SetStringProperty](https://wiki.libsdl.org/SDL3/SDL_SetStringProperty) | [:x:](./methods.go#L5337) | [:x:](./sdl_functions_js.go#L645) |
| [SDL_SetNumberProperty](https://wiki.libsdl.org/SDL3/SDL_SetNumberProperty) | [:x:](./methods.go#L5344) | [:x:](./sdl_functions_js.go#L662) |
| [SDL_SetFloatProperty](https://wiki.libsdl.org/SDL3/SDL_SetFloatProperty) | [:x:](./methods.go#L5351) | [:x:](./sdl_functions_js.go#L679) |
| [SDL_SetBooleanProperty](https://wiki.libsdl.org/SDL3/SDL_SetBooleanProperty) | [:x:](./methods.go#L5358) | [:x:](./sdl_functions_js.go#L696) |
| [SDL_HasProperty](https://wiki.libsdl.org/SDL3/SDL_HasProperty) | [:x:](./methods.go#L5365) | [:x:](./sdl_functions_js.go#L713) |
| [SDL_GetPropertyType](https://wiki.libsdl.org/SDL3/SDL_GetPropertyType) | [:x:](./methods.go#L5372) | [:x:](./sdl_functions_js.go#L728) |
| [SDL_GetPointerProperty](https://wiki.libsdl.org/SDL3/SDL_GetPointerProperty) | [:x:](./methods.go#L5379) | [:x:](./sdl_functions_js.go#L743) |
| [SDL_GetStringProperty](https://wiki.libsdl.org/SDL3/SDL_GetStringProperty) | [:x:](./methods.go#L5386) | [:x:](./sdl_functions_js.go#L760) |
| [SDL_GetNumberProperty](https://wiki.libsdl.org/SDL3/SDL_GetNumberProperty) | [:x:](./methods.go#L5393) | [:x:](./sdl_functions_js.go#L777) |
| [SDL_GetFloatProperty](https://wiki.libsdl.org/SDL3/SDL_GetFloatProperty) | [:x:](./methods.go#L5400) | [:x:](./sdl_functions_js.go#L794) |
| [SDL_GetBooleanProperty](https://wiki.libsdl.org/SDL3/SDL_GetBooleanProperty) | [:x:](./methods.go#L5407) | [:x:](./sdl_functions_js.go#L811) |
| [SDL_ClearProperty](https://wiki.libsdl.org/SDL3/SDL_ClearProperty) | [:x:](./methods.go#L5414) | [:x:](./sdl_functions_js.go#L828) |
| [SDL_EnumerateProperties](https://wiki.libsdl.org/SDL3/SDL_EnumerateProperties) | [:x:](./methods.go#L5421) | [:x:](./sdl_functions_js.go#L843) |
| [SDL_DestroyProperties](https://wiki.libsdl.org/SDL3/SDL_DestroyProperties) | [:x:](./methods.go#L5428) | [:x:](./sdl_functions_js.go#L860) |
### Log

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetLogPriorities](https://wiki.libsdl.org/SDL3/SDL_SetLogPriorities) | [:question:]() | [:question:](./sdl_functions_js.go#L14098) |
| [SDL_SetLogPriority](https://wiki.libsdl.org/SDL3/SDL_SetLogPriority) | [:question:]() | [:question:](./sdl_functions_js.go#L14109) |
| [SDL_GetLogPriority](https://wiki.libsdl.org/SDL3/SDL_GetLogPriority) | [:question:]() | [:question:](./sdl_functions_js.go#L14122) |
| [SDL_ResetLogPriorities](https://wiki.libsdl.org/SDL3/SDL_ResetLogPriorities) | [:question:]() | [:question:](./sdl_functions_js.go#L14135) |
| [SDL_SetLogPriorityPrefix](https://wiki.libsdl.org/SDL3/SDL_SetLogPriorityPrefix) | [:question:]() | [:question:](./sdl_functions_js.go#L14144) |
| [SDL_Log](https://wiki.libsdl.org/SDL3/SDL_Log) | [:question:]() | [:question:](./sdl_functions_js.go#L14159) |
| [SDL_LogTrace](https://wiki.libsdl.org/SDL3/SDL_LogTrace) | [:question:]() | [:question:](./sdl_functions_js.go#L14170) |
| [SDL_LogVerbose](https://wiki.libsdl.org/SDL3/SDL_LogVerbose) | [:question:]() | [:question:](./sdl_functions_js.go#L14183) |
| [SDL_LogDebug](https://wiki.libsdl.org/SDL3/SDL_LogDebug) | [:question:]() | [:question:](./sdl_functions_js.go#L14196) |
| [SDL_LogInfo](https://wiki.libsdl.org/SDL3/SDL_LogInfo) | [:question:]() | [:question:](./sdl_functions_js.go#L14209) |
| [SDL_LogWarn](https://wiki.libsdl.org/SDL3/SDL_LogWarn) | [:question:]() | [:question:](./sdl_functions_js.go#L14222) |
| [SDL_LogError](https://wiki.libsdl.org/SDL3/SDL_LogError) | [:question:]() | [:question:](./sdl_functions_js.go#L14235) |
| [SDL_LogCritical](https://wiki.libsdl.org/SDL3/SDL_LogCritical) | [:question:]() | [:question:](./sdl_functions_js.go#L14248) |
| [SDL_LogMessage](https://wiki.libsdl.org/SDL3/SDL_LogMessage) | [:question:]() | [:question:](./sdl_functions_js.go#L14261) |
| [SDL_LogMessageV](https://wiki.libsdl.org/SDL3/SDL_LogMessageV) | [:question:]() | [:question:](./sdl_functions_js.go#L14276) |
| [SDL_GetDefaultLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_GetDefaultLogOutputFunction) | [:question:]() | [:question:](./sdl_functions_js.go#L14293) |
| [SDL_GetLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_GetLogOutputFunction) | [:question:]() | [:question:](./sdl_functions_js.go#L14304) |
| [SDL_SetLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_SetLogOutputFunction) | [:x:](./methods.go#L3577) | [:x:](./sdl_functions_js.go#L14323) |
### Video

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumVideoDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumVideoDrivers) | [:heavy_check_mark:](./functions.go#L438) | [:x:](./sdl_functions_js.go#L5828) |
| [SDL_GetVideoDriver](https://wiki.libsdl.org/SDL3/SDL_GetVideoDriver) | [:heavy_check_mark:](./functions.go#L444) | [:x:](./sdl_functions_js.go#L5839) |
| [SDL_GetCurrentVideoDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentVideoDriver) | [:heavy_check_mark:](./functions.go#L450) | [:x:](./sdl_functions_js.go#L5852) |
| [SDL_GetSystemTheme](https://wiki.libsdl.org/SDL3/SDL_GetSystemTheme) | [:heavy_check_mark:](./functions.go#L456) | [:x:](./sdl_functions_js.go#L5863) |
| [SDL_GetDisplays](https://wiki.libsdl.org/SDL3/SDL_GetDisplays) | [:heavy_check_mark:](./functions.go#L462) | [:x:](./sdl_functions_js.go#L5874) |
| [SDL_GetPrimaryDisplay](https://wiki.libsdl.org/SDL3/SDL_GetPrimaryDisplay) | [:heavy_check_mark:](./functions.go#L476) | [:x:](./sdl_functions_js.go#L5890) |
| [SDL_GetDisplayProperties](https://wiki.libsdl.org/SDL3/SDL_GetDisplayProperties) | [:heavy_check_mark:](./methods.go#L3134) | [:x:](./sdl_functions_js.go#L5901) |
| [SDL_GetDisplayName](https://wiki.libsdl.org/SDL3/SDL_GetDisplayName) | [:heavy_check_mark:](./methods.go#L3145) | [:x:](./sdl_functions_js.go#L5914) |
| [SDL_GetDisplayBounds](https://wiki.libsdl.org/SDL3/SDL_GetDisplayBounds) | [:heavy_check_mark:](./methods.go#L3156) | [:x:](./sdl_functions_js.go#L5927) |
| [SDL_GetDisplayUsableBounds](https://wiki.libsdl.org/SDL3/SDL_GetDisplayUsableBounds) | [:heavy_check_mark:](./methods.go#L3168) | [:x:](./sdl_functions_js.go#L5945) |
| [SDL_GetNaturalDisplayOrientation](https://wiki.libsdl.org/SDL3/SDL_GetNaturalDisplayOrientation) | [:heavy_check_mark:](./methods.go#L3180) | [:x:](./sdl_functions_js.go#L5963) |
| [SDL_GetCurrentDisplayOrientation](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayOrientation) | [:heavy_check_mark:](./methods.go#L3186) | [:x:](./sdl_functions_js.go#L5976) |
| [SDL_GetDisplayContentScale](https://wiki.libsdl.org/SDL3/SDL_GetDisplayContentScale) | [:heavy_check_mark:](./methods.go#L3192) | [:x:](./sdl_functions_js.go#L5989) |
| [SDL_GetFullscreenDisplayModes](https://wiki.libsdl.org/SDL3/SDL_GetFullscreenDisplayModes) | [:heavy_check_mark:](./methods.go#L3203) | [:x:](./sdl_functions_js.go#L6002) |
| [SDL_GetClosestFullscreenDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetClosestFullscreenDisplayMode) | [:heavy_check_mark:](./methods.go#L3217) | [:x:](./sdl_functions_js.go#L6020) |
| [SDL_GetDesktopDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetDesktopDisplayMode) | [:heavy_check_mark:](./methods.go#L3229) | [:x:](./sdl_functions_js.go#L6046) |
| [SDL_GetCurrentDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayMode) | [:heavy_check_mark:](./methods.go#L3240) | [:x:](./sdl_functions_js.go#L6062) |
| [SDL_GetDisplayForPoint](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForPoint) | [:heavy_check_mark:](./functions.go#L482) | [:x:](./sdl_functions_js.go#L6078) |
| [SDL_GetDisplayForRect](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForRect) | [:heavy_check_mark:](./functions.go#L488) | [:x:](./sdl_functions_js.go#L6094) |
| [SDL_GetDisplayForWindow](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForWindow) | [:heavy_check_mark:](./functions.go#L494) | [:x:](./sdl_functions_js.go#L6110) |
| [SDL_GetWindowPixelDensity](https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelDensity) | [:heavy_check_mark:](./methods.go#L3623) | [:x:](./sdl_functions_js.go#L6126) |
| [SDL_GetWindowDisplayScale](https://wiki.libsdl.org/SDL3/SDL_GetWindowDisplayScale) | [:heavy_check_mark:](./methods.go#L3634) | [:x:](./sdl_functions_js.go#L6142) |
| [SDL_SetWindowFullscreenMode](https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreenMode) | [:heavy_check_mark:](./methods.go#L3645) | [:x:](./sdl_functions_js.go#L6158) |
| [SDL_GetWindowFullscreenMode](https://wiki.libsdl.org/SDL3/SDL_GetWindowFullscreenMode) | [:heavy_check_mark:](./methods.go#L3655) | [:x:](./sdl_functions_js.go#L6179) |
| [SDL_GetWindowICCProfile](https://wiki.libsdl.org/SDL3/SDL_GetWindowICCProfile) | [:heavy_check_mark:](./methods.go#L3661) | [:x:](./sdl_functions_js.go#L6198) |
| [SDL_GetWindowPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelFormat) | [:heavy_check_mark:](./methods.go#L3675) | [:x:](./sdl_functions_js.go#L6219) |
| [SDL_GetWindows](https://wiki.libsdl.org/SDL3/SDL_GetWindows) | [:heavy_check_mark:](./functions.go#L500) | [:x:](./sdl_functions_js.go#L6235) |
| [SDL_CreateWindow](https://wiki.libsdl.org/SDL3/SDL_CreateWindow) | [:heavy_check_mark:](./functions.go#L514) | [:x:](./sdl_functions_js.go#L6251) |
| [SDL_CreatePopupWindow](https://wiki.libsdl.org/SDL3/SDL_CreatePopupWindow) | [:heavy_check_mark:](./methods.go#L3686) | [:x:](./sdl_functions_js.go#L6273) |
| [SDL_CreateWindowWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateWindowWithProperties) | [:heavy_check_mark:](./functions.go#L525) | [:x:](./sdl_functions_js.go#L6302) |
| [SDL_GetWindowID](https://wiki.libsdl.org/SDL3/SDL_GetWindowID) | [:heavy_check_mark:](./methods.go#L3697) | [:x:](./sdl_functions_js.go#L6318) |
| [SDL_GetWindowFromID](https://wiki.libsdl.org/SDL3/SDL_GetWindowFromID) | [:heavy_check_mark:](./methods.go#L50) | [:x:](./sdl_functions_js.go#L6334) |
| [SDL_GetWindowParent](https://wiki.libsdl.org/SDL3/SDL_GetWindowParent) | [:heavy_check_mark:](./methods.go#L3708) | [:x:](./sdl_functions_js.go#L6350) |
| [SDL_GetWindowProperties](https://wiki.libsdl.org/SDL3/SDL_GetWindowProperties) | [:heavy_check_mark:](./methods.go#L3714) | [:x:](./sdl_functions_js.go#L6369) |
| [SDL_GetWindowFlags](https://wiki.libsdl.org/SDL3/SDL_GetWindowFlags) | [:heavy_check_mark:](./methods.go#L3725) | [:x:](./sdl_functions_js.go#L6385) |
| [SDL_SetWindowTitle](https://wiki.libsdl.org/SDL3/SDL_SetWindowTitle) | [:heavy_check_mark:](./methods.go#L3731) | [:x:](./sdl_functions_js.go#L6401) |
| [SDL_GetWindowTitle](https://wiki.libsdl.org/SDL3/SDL_GetWindowTitle) | [:heavy_check_mark:](./methods.go#L3741) | [:x:](./sdl_functions_js.go#L6419) |
| [SDL_SetWindowIcon](https://wiki.libsdl.org/SDL3/SDL_SetWindowIcon) | [:heavy_check_mark:](./methods.go#L3747) | [:x:](./sdl_functions_js.go#L6435) |
| [SDL_SetWindowPosition](https://wiki.libsdl.org/SDL3/SDL_SetWindowPosition) | [:heavy_check_mark:](./methods.go#L3757) | [:x:](./sdl_functions_js.go#L6456) |
| [SDL_GetWindowPosition](https://wiki.libsdl.org/SDL3/SDL_GetWindowPosition) | [:heavy_check_mark:](./methods.go#L3767) | [:x:](./sdl_functions_js.go#L6476) |
| [SDL_SetWindowSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowSize) | [:heavy_check_mark:](./methods.go#L3779) | [:x:](./sdl_functions_js.go#L6502) |
| [SDL_GetWindowSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowSize) | [:heavy_check_mark:](./methods.go#L3789) | [:heavy_check_mark:](./sdl_functions_js.go#L6522) |
| [SDL_GetWindowSafeArea](https://wiki.libsdl.org/SDL3/SDL_GetWindowSafeArea) | [:heavy_check_mark:](./methods.go#L3801) | [:x:](./sdl_functions_js.go#L6544) |
| [SDL_SetWindowAspectRatio](https://wiki.libsdl.org/SDL3/SDL_SetWindowAspectRatio) | [:heavy_check_mark:](./methods.go#L3813) | [:x:](./sdl_functions_js.go#L6565) |
| [SDL_GetWindowAspectRatio](https://wiki.libsdl.org/SDL3/SDL_GetWindowAspectRatio) | [:heavy_check_mark:](./methods.go#L3823) | [:x:](./sdl_functions_js.go#L6585) |
| [SDL_GetWindowBordersSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowBordersSize) | [:heavy_check_mark:](./methods.go#L3835) | [:x:](./sdl_functions_js.go#L6611) |
| [SDL_GetWindowSizeInPixels](https://wiki.libsdl.org/SDL3/SDL_GetWindowSizeInPixels) | [:heavy_check_mark:](./methods.go#L3847) | [:x:](./sdl_functions_js.go#L6647) |
| [SDL_SetWindowMinimumSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowMinimumSize) | [:heavy_check_mark:](./methods.go#L3859) | [:x:](./sdl_functions_js.go#L6673) |
| [SDL_GetWindowMinimumSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowMinimumSize) | [:heavy_check_mark:](./methods.go#L3869) | [:x:](./sdl_functions_js.go#L6693) |
| [SDL_SetWindowMaximumSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowMaximumSize) | [:heavy_check_mark:](./methods.go#L3881) | [:x:](./sdl_functions_js.go#L6719) |
| [SDL_GetWindowMaximumSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowMaximumSize) | [:heavy_check_mark:](./methods.go#L3891) | [:x:](./sdl_functions_js.go#L6739) |
| [SDL_SetWindowBordered](https://wiki.libsdl.org/SDL3/SDL_SetWindowBordered) | [:heavy_check_mark:](./methods.go#L3903) | [:x:](./sdl_functions_js.go#L6765) |
| [SDL_SetWindowResizable](https://wiki.libsdl.org/SDL3/SDL_SetWindowResizable) | [:heavy_check_mark:](./methods.go#L3913) | [:x:](./sdl_functions_js.go#L6783) |
| [SDL_SetWindowAlwaysOnTop](https://wiki.libsdl.org/SDL3/SDL_SetWindowAlwaysOnTop) | [:heavy_check_mark:](./methods.go#L3923) | [:x:](./sdl_functions_js.go#L6801) |
| [SDL_ShowWindow](https://wiki.libsdl.org/SDL3/SDL_ShowWindow) | [:heavy_check_mark:](./methods.go#L3933) | [:x:](./sdl_functions_js.go#L6819) |
| [SDL_HideWindow](https://wiki.libsdl.org/SDL3/SDL_HideWindow) | [:heavy_check_mark:](./methods.go#L3943) | [:x:](./sdl_functions_js.go#L6835) |
| [SDL_RaiseWindow](https://wiki.libsdl.org/SDL3/SDL_RaiseWindow) | [:heavy_check_mark:](./methods.go#L3953) | [:x:](./sdl_functions_js.go#L6851) |
| [SDL_MaximizeWindow](https://wiki.libsdl.org/SDL3/SDL_MaximizeWindow) | [:heavy_check_mark:](./methods.go#L3963) | [:x:](./sdl_functions_js.go#L6867) |
| [SDL_MinimizeWindow](https://wiki.libsdl.org/SDL3/SDL_MinimizeWindow) | [:heavy_check_mark:](./methods.go#L3973) | [:x:](./sdl_functions_js.go#L6883) |
| [SDL_RestoreWindow](https://wiki.libsdl.org/SDL3/SDL_RestoreWindow) | [:heavy_check_mark:](./methods.go#L3983) | [:x:](./sdl_functions_js.go#L6899) |
| [SDL_SetWindowFullscreen](https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreen) | [:heavy_check_mark:](./methods.go#L3993) | [:x:](./sdl_functions_js.go#L6915) |
| [SDL_SyncWindow](https://wiki.libsdl.org/SDL3/SDL_SyncWindow) | [:heavy_check_mark:](./methods.go#L4003) | [:x:](./sdl_functions_js.go#L6933) |
| [SDL_WindowHasSurface](https://wiki.libsdl.org/SDL3/SDL_WindowHasSurface) | [:heavy_check_mark:](./methods.go#L4013) | [:x:](./sdl_functions_js.go#L6949) |
| [SDL_GetWindowSurface](https://wiki.libsdl.org/SDL3/SDL_GetWindowSurface) | [:heavy_check_mark:](./methods.go#L4019) | [:x:](./sdl_functions_js.go#L6965) |
| [SDL_SetWindowSurfaceVSync](https://wiki.libsdl.org/SDL3/SDL_SetWindowSurfaceVSync) | [:heavy_check_mark:](./methods.go#L4030) | [:x:](./sdl_functions_js.go#L6984) |
| [SDL_GetWindowSurfaceVSync](https://wiki.libsdl.org/SDL3/SDL_GetWindowSurfaceVSync) | [:heavy_check_mark:](./methods.go#L4040) | [:x:](./sdl_functions_js.go#L7002) |
| [SDL_UpdateWindowSurface](https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurface) | [:heavy_check_mark:](./methods.go#L4052) | [:x:](./sdl_functions_js.go#L7023) |
| [SDL_UpdateWindowSurfaceRects](https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurfaceRects) | [:heavy_check_mark:](./methods.go#L4062) | [:x:](./sdl_functions_js.go#L7039) |
| [SDL_DestroyWindowSurface](https://wiki.libsdl.org/SDL3/SDL_DestroyWindowSurface) | [:heavy_check_mark:](./methods.go#L4072) | [:x:](./sdl_functions_js.go#L7062) |
| [SDL_SetWindowKeyboardGrab](https://wiki.libsdl.org/SDL3/SDL_SetWindowKeyboardGrab) | [:heavy_check_mark:](./methods.go#L4082) | [:x:](./sdl_functions_js.go#L7078) |
| [SDL_SetWindowMouseGrab](https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseGrab) | [:heavy_check_mark:](./methods.go#L4092) | [:x:](./sdl_functions_js.go#L7096) |
| [SDL_GetWindowKeyboardGrab](https://wiki.libsdl.org/SDL3/SDL_GetWindowKeyboardGrab) | [:heavy_check_mark:](./methods.go#L4102) | [:x:](./sdl_functions_js.go#L7114) |
| [SDL_GetWindowMouseGrab](https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseGrab) | [:heavy_check_mark:](./methods.go#L4108) | [:x:](./sdl_functions_js.go#L7130) |
| [SDL_GetGrabbedWindow](https://wiki.libsdl.org/SDL3/SDL_GetGrabbedWindow) | [:heavy_check_mark:](./functions.go#L536) | [:x:](./sdl_functions_js.go#L7146) |
| [SDL_SetWindowMouseRect](https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseRect) | [:heavy_check_mark:](./methods.go#L4114) | [:x:](./sdl_functions_js.go#L7160) |
| [SDL_GetWindowMouseRect](https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseRect) | [:heavy_check_mark:](./methods.go#L4124) | [:x:](./sdl_functions_js.go#L7181) |
| [SDL_SetWindowOpacity](https://wiki.libsdl.org/SDL3/SDL_SetWindowOpacity) | [:heavy_check_mark:](./methods.go#L4130) | [:x:](./sdl_functions_js.go#L7200) |
| [SDL_GetWindowOpacity](https://wiki.libsdl.org/SDL3/SDL_GetWindowOpacity) | [:heavy_check_mark:](./methods.go#L4140) | [:x:](./sdl_functions_js.go#L7218) |
| [SDL_SetWindowParent](https://wiki.libsdl.org/SDL3/SDL_SetWindowParent) | [:heavy_check_mark:](./methods.go#L4146) | [:x:](./sdl_functions_js.go#L7234) |
| [SDL_SetWindowModal](https://wiki.libsdl.org/SDL3/SDL_SetWindowModal) | [:heavy_check_mark:](./methods.go#L4156) | [:x:](./sdl_functions_js.go#L7255) |
| [SDL_SetWindowFocusable](https://wiki.libsdl.org/SDL3/SDL_SetWindowFocusable) | [:heavy_check_mark:](./methods.go#L4166) | [:x:](./sdl_functions_js.go#L7273) |
| [SDL_ShowWindowSystemMenu](https://wiki.libsdl.org/SDL3/SDL_ShowWindowSystemMenu) | [:heavy_check_mark:](./methods.go#L4176) | [:x:](./sdl_functions_js.go#L7291) |
| [SDL_SetWindowHitTest](https://wiki.libsdl.org/SDL3/SDL_SetWindowHitTest) | [:x:](./methods.go#L4186) | [:x:](./sdl_functions_js.go#L7311) |
| [SDL_SetWindowShape](https://wiki.libsdl.org/SDL3/SDL_SetWindowShape) | [:heavy_check_mark:](./methods.go#L4193) | [:x:](./sdl_functions_js.go#L7331) |
| [SDL_FlashWindow](https://wiki.libsdl.org/SDL3/SDL_FlashWindow) | [:heavy_check_mark:](./methods.go#L4203) | [:x:](./sdl_functions_js.go#L7352) |
| [SDL_DestroyWindow](https://wiki.libsdl.org/SDL3/SDL_DestroyWindow) | [:heavy_check_mark:](./methods.go#L4213) | [:heavy_check_mark:](./sdl_functions_js.go#L7370) |
| [SDL_ScreenSaverEnabled](https://wiki.libsdl.org/SDL3/SDL_ScreenSaverEnabled) | [:heavy_check_mark:](./functions.go#L542) | [:x:](./sdl_functions_js.go#L7382) |
| [SDL_EnableScreenSaver](https://wiki.libsdl.org/SDL3/SDL_EnableScreenSaver) | [:heavy_check_mark:](./functions.go#L548) | [:x:](./sdl_functions_js.go#L7393) |
| [SDL_DisableScreenSaver](https://wiki.libsdl.org/SDL3/SDL_DisableScreenSaver) | [:heavy_check_mark:](./functions.go#L558) | [:x:](./sdl_functions_js.go#L7404) |
| [SDL_GL_LoadLibrary](https://wiki.libsdl.org/SDL3/SDL_GL_LoadLibrary) | [:question:]() | [:question:](./sdl_functions_js.go#L7415) |
| [SDL_GL_GetProcAddress](https://wiki.libsdl.org/SDL3/SDL_GL_GetProcAddress) | [:question:]() | [:question:](./sdl_functions_js.go#L7428) |
| [SDL_EGL_GetProcAddress](https://wiki.libsdl.org/SDL3/SDL_EGL_GetProcAddress) | [:question:]() | [:question:](./sdl_functions_js.go#L7441) |
| [SDL_GL_UnloadLibrary](https://wiki.libsdl.org/SDL3/SDL_GL_UnloadLibrary) | [:question:]() | [:question:](./sdl_functions_js.go#L7454) |
| [SDL_GL_ExtensionSupported](https://wiki.libsdl.org/SDL3/SDL_GL_ExtensionSupported) | [:question:]() | [:question:](./sdl_functions_js.go#L7463) |
| [SDL_GL_ResetAttributes](https://wiki.libsdl.org/SDL3/SDL_GL_ResetAttributes) | [:question:]() | [:question:](./sdl_functions_js.go#L7476) |
| [SDL_GL_SetAttribute](https://wiki.libsdl.org/SDL3/SDL_GL_SetAttribute) | [:x:](./methods.go#L4367) | [:x:](./sdl_functions_js.go#L7485) |
| [SDL_GL_GetAttribute](https://wiki.libsdl.org/SDL3/SDL_GL_GetAttribute) | [:x:](./methods.go#L4374) | [:x:](./sdl_functions_js.go#L7500) |
| [SDL_GL_CreateContext](https://wiki.libsdl.org/SDL3/SDL_GL_CreateContext) | [:x:](./methods.go#L4219) | [:x:](./sdl_functions_js.go#L7518) |
| [SDL_GL_MakeCurrent](https://wiki.libsdl.org/SDL3/SDL_GL_MakeCurrent) | [:x:](./methods.go#L4226) | [:x:](./sdl_functions_js.go#L7534) |
| [SDL_GL_GetCurrentWindow](https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentWindow) | [:question:]() | [:question:](./sdl_functions_js.go#L7552) |
| [SDL_GL_GetCurrentContext](https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentContext) | [:question:]() | [:question:](./sdl_functions_js.go#L7566) |
| [SDL_EGL_GetCurrentDisplay](https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentDisplay) | [:question:]() | [:question:](./sdl_functions_js.go#L7577) |
| [SDL_EGL_GetCurrentConfig](https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentConfig) | [:question:]() | [:question:](./sdl_functions_js.go#L7588) |
| [SDL_EGL_GetWindowSurface](https://wiki.libsdl.org/SDL3/SDL_EGL_GetWindowSurface) | [:x:](./methods.go#L4233) | [:x:](./sdl_functions_js.go#L7599) |
| [SDL_EGL_SetAttributeCallbacks](https://wiki.libsdl.org/SDL3/SDL_EGL_SetAttributeCallbacks) | [:question:]() | [:question:](./sdl_functions_js.go#L7615) |
| [SDL_GL_SetSwapInterval](https://wiki.libsdl.org/SDL3/SDL_GL_SetSwapInterval) | [:question:]() | [:question:](./sdl_functions_js.go#L7632) |
| [SDL_GL_GetSwapInterval](https://wiki.libsdl.org/SDL3/SDL_GL_GetSwapInterval) | [:question:]() | [:question:](./sdl_functions_js.go#L7645) |
| [SDL_GL_SwapWindow](https://wiki.libsdl.org/SDL3/SDL_GL_SwapWindow) | [:x:](./methods.go#L4240) | [:x:](./sdl_functions_js.go#L7661) |
| [SDL_GL_DestroyContext](https://wiki.libsdl.org/SDL3/SDL_GL_DestroyContext) | [:question:]() | [:question:](./sdl_functions_js.go#L7677) |
### Events

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_PumpEvents](https://wiki.libsdl.org/SDL3/SDL_PumpEvents) | [:heavy_check_mark:](./functions.go#L144) | [:x:](./sdl_functions_js.go#L10940) |
| [SDL_PeepEvents](https://wiki.libsdl.org/SDL3/SDL_PeepEvents) | [:question:]() | [:question:](./sdl_functions_js.go#L10949) |
| [SDL_HasEvent](https://wiki.libsdl.org/SDL3/SDL_HasEvent) | [:heavy_check_mark:](./functions.go#L152) | [:x:](./sdl_functions_js.go#L10973) |
| [SDL_HasEvents](https://wiki.libsdl.org/SDL3/SDL_HasEvents) | [:heavy_check_mark:](./functions.go#L158) | [:x:](./sdl_functions_js.go#L10986) |
| [SDL_FlushEvent](https://wiki.libsdl.org/SDL3/SDL_FlushEvent) | [:heavy_check_mark:](./functions.go#L164) | [:x:](./sdl_functions_js.go#L11001) |
| [SDL_FlushEvents](https://wiki.libsdl.org/SDL3/SDL_FlushEvents) | [:heavy_check_mark:](./functions.go#L170) | [:x:](./sdl_functions_js.go#L11012) |
| [SDL_PollEvent](https://wiki.libsdl.org/SDL3/SDL_PollEvent) | [:heavy_check_mark:](./functions.go#L176) | [:heavy_check_mark:](./sdl_functions_js.go#L11025) |
| [SDL_WaitEvent](https://wiki.libsdl.org/SDL3/SDL_WaitEvent) | [:heavy_check_mark:](./functions.go#L182) | [:x:](./sdl_functions_js.go#L11038) |
| [SDL_WaitEventTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitEventTimeout) | [:heavy_check_mark:](./functions.go#L192) | [:x:](./sdl_functions_js.go#L11054) |
| [SDL_PushEvent](https://wiki.libsdl.org/SDL3/SDL_PushEvent) | [:heavy_check_mark:](./functions.go#L198) | [:x:](./sdl_functions_js.go#L11072) |
| [SDL_SetEventFilter](https://wiki.libsdl.org/SDL3/SDL_SetEventFilter) | [:question:]() | [:question:](./sdl_functions_js.go#L11088) |
| [SDL_GetEventFilter](https://wiki.libsdl.org/SDL3/SDL_GetEventFilter) | [:x:](./methods.go#L3279) | [:x:](./sdl_functions_js.go#L11101) |
| [SDL_AddEventWatch](https://wiki.libsdl.org/SDL3/SDL_AddEventWatch) | [:question:]() | [:question:](./sdl_functions_js.go#L11122) |
| [SDL_RemoveEventWatch](https://wiki.libsdl.org/SDL3/SDL_RemoveEventWatch) | [:question:]() | [:question:](./sdl_functions_js.go#L11137) |
| [SDL_FilterEvents](https://wiki.libsdl.org/SDL3/SDL_FilterEvents) | [:question:]() | [:question:](./sdl_functions_js.go#L11150) |
| [SDL_SetEventEnabled](https://wiki.libsdl.org/SDL3/SDL_SetEventEnabled) | [:heavy_check_mark:](./functions.go#L214) | [:x:](./sdl_functions_js.go#L11163) |
| [SDL_EventEnabled](https://wiki.libsdl.org/SDL3/SDL_EventEnabled) | [:heavy_check_mark:](./functions.go#L220) | [:x:](./sdl_functions_js.go#L11176) |
| [SDL_RegisterEvents](https://wiki.libsdl.org/SDL3/SDL_RegisterEvents) | [:question:]() | [:question:](./sdl_functions_js.go#L11189) |
| [SDL_GetWindowFromEvent](https://wiki.libsdl.org/SDL3/SDL_GetWindowFromEvent) | [:x:](./methods.go#L1695) | [:x:](./sdl_functions_js.go#L11202) |
### Keyboard

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_HasKeyboard](https://wiki.libsdl.org/SDL3/SDL_HasKeyboard) | [:heavy_check_mark:](./functions.go#L813) | [:x:](./sdl_functions_js.go#L10161) |
| [SDL_GetKeyboards](https://wiki.libsdl.org/SDL3/SDL_GetKeyboards) | [:heavy_check_mark:](./functions.go#L819) | [:x:](./sdl_functions_js.go#L10172) |
| [SDL_GetKeyboardNameForID](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardNameForID) | [:heavy_check_mark:](./methods.go#L3253) | [:x:](./sdl_functions_js.go#L10188) |
| [SDL_GetKeyboardFocus](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardFocus) | [:question:]() | [:question:](./sdl_functions_js.go#L10201) |
| [SDL_GetKeyboardState](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardState) | [:heavy_check_mark:](./functions.go#L833) | [:x:](./sdl_functions_js.go#L10215) |
| [SDL_ResetKeyboard](https://wiki.libsdl.org/SDL3/SDL_ResetKeyboard) | [:heavy_check_mark:](./functions.go#L843) | [:x:](./sdl_functions_js.go#L10233) |
| [SDL_GetModState](https://wiki.libsdl.org/SDL3/SDL_GetModState) | [:question:]() | [:question:](./sdl_functions_js.go#L10242) |
| [SDL_SetModState](https://wiki.libsdl.org/SDL3/SDL_SetModState) | [:x:](./methods.go#L2419) | [:x:](./sdl_functions_js.go#L10253) |
| [SDL_GetKeyFromScancode](https://wiki.libsdl.org/SDL3/SDL_GetKeyFromScancode) | [:heavy_check_mark:](./methods.go#L4383) | [:x:](./sdl_functions_js.go#L10264) |
| [SDL_GetScancodeFromKey](https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromKey) | [:heavy_check_mark:](./methods.go#L5543) | [:x:](./sdl_functions_js.go#L10281) |
| [SDL_SetScancodeName](https://wiki.libsdl.org/SDL3/SDL_SetScancodeName) | [:heavy_check_mark:](./methods.go#L4389) | [:x:](./sdl_functions_js.go#L10299) |
| [SDL_GetScancodeName](https://wiki.libsdl.org/SDL3/SDL_GetScancodeName) | [:heavy_check_mark:](./methods.go#L4399) | [:x:](./sdl_functions_js.go#L10314) |
| [SDL_GetScancodeFromName](https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromName) | [:heavy_check_mark:](./functions.go#L849) | [:x:](./sdl_functions_js.go#L10327) |
| [SDL_GetKeyName](https://wiki.libsdl.org/SDL3/SDL_GetKeyName) | [:heavy_check_mark:](./methods.go#L5549) | [:heavy_check_mark:](./sdl_functions_js.go#L10340) |
| [SDL_GetKeyFromName](https://wiki.libsdl.org/SDL3/SDL_GetKeyFromName) | [:heavy_check_mark:](./functions.go#L855) | [:x:](./sdl_functions_js.go#L10350) |
| [SDL_StartTextInput](https://wiki.libsdl.org/SDL3/SDL_StartTextInput) | [:heavy_check_mark:](./methods.go#L4247) | [:x:](./sdl_functions_js.go#L10363) |
| [SDL_StartTextInputWithProperties](https://wiki.libsdl.org/SDL3/SDL_StartTextInputWithProperties) | [:heavy_check_mark:](./methods.go#L4257) | [:x:](./sdl_functions_js.go#L10379) |
| [SDL_TextInputActive](https://wiki.libsdl.org/SDL3/SDL_TextInputActive) | [:heavy_check_mark:](./methods.go#L4267) | [:x:](./sdl_functions_js.go#L10397) |
| [SDL_StopTextInput](https://wiki.libsdl.org/SDL3/SDL_StopTextInput) | [:heavy_check_mark:](./methods.go#L4273) | [:x:](./sdl_functions_js.go#L10413) |
| [SDL_ClearComposition](https://wiki.libsdl.org/SDL3/SDL_ClearComposition) | [:x:](./methods.go#L4283) | [:x:](./sdl_functions_js.go#L10429) |
| [SDL_SetTextInputArea](https://wiki.libsdl.org/SDL3/SDL_SetTextInputArea) | [:heavy_check_mark:](./methods.go#L4290) | [:x:](./sdl_functions_js.go#L10445) |
| [SDL_GetTextInputArea](https://wiki.libsdl.org/SDL3/SDL_GetTextInputArea) | [:heavy_check_mark:](./methods.go#L4300) | [:x:](./sdl_functions_js.go#L10468) |
| [SDL_HasScreenKeyboardSupport](https://wiki.libsdl.org/SDL3/SDL_HasScreenKeyboardSupport) | [:heavy_check_mark:](./functions.go#L861) | [:x:](./sdl_functions_js.go#L10494) |
| [SDL_ScreenKeyboardShown](https://wiki.libsdl.org/SDL3/SDL_ScreenKeyboardShown) | [:heavy_check_mark:](./methods.go#L4313) | [:x:](./sdl_functions_js.go#L10505) |
### Mouse

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_HasMouse](https://wiki.libsdl.org/SDL3/SDL_HasMouse) | [:heavy_check_mark:](./functions.go#L869) | [:x:](./sdl_functions_js.go#L10521) |
| [SDL_GetMice](https://wiki.libsdl.org/SDL3/SDL_GetMice) | [:heavy_check_mark:](./functions.go#L875) | [:x:](./sdl_functions_js.go#L10532) |
| [SDL_GetMouseNameForID](https://wiki.libsdl.org/SDL3/SDL_GetMouseNameForID) | [:heavy_check_mark:](./methods.go#L3266) | [:x:](./sdl_functions_js.go#L10548) |
| [SDL_GetMouseFocus](https://wiki.libsdl.org/SDL3/SDL_GetMouseFocus) | [:heavy_check_mark:](./functions.go#L889) | [:x:](./sdl_functions_js.go#L10561) |
| [SDL_GetMouseState](https://wiki.libsdl.org/SDL3/SDL_GetMouseState) | [:heavy_check_mark:](./functions.go#L895) | [:x:](./sdl_functions_js.go#L10575) |
| [SDL_GetGlobalMouseState](https://wiki.libsdl.org/SDL3/SDL_GetGlobalMouseState) | [:heavy_check_mark:](./functions.go#L905) | [:x:](./sdl_functions_js.go#L10596) |
| [SDL_GetRelativeMouseState](https://wiki.libsdl.org/SDL3/SDL_GetRelativeMouseState) | [:heavy_check_mark:](./functions.go#L915) | [:x:](./sdl_functions_js.go#L10617) |
| [SDL_WarpMouseInWindow](https://wiki.libsdl.org/SDL3/SDL_WarpMouseInWindow) | [:heavy_check_mark:](./methods.go#L4319) | [:x:](./sdl_functions_js.go#L10638) |
| [SDL_WarpMouseGlobal](https://wiki.libsdl.org/SDL3/SDL_WarpMouseGlobal) | [:heavy_check_mark:](./functions.go#L925) | [:x:](./sdl_functions_js.go#L10656) |
| [SDL_SetRelativeMouseTransform](https://wiki.libsdl.org/SDL3/SDL_SetRelativeMouseTransform) | [:question:]() | [:question:]() |
| [SDL_SetWindowRelativeMouseMode](https://wiki.libsdl.org/SDL3/SDL_SetWindowRelativeMouseMode) | [:heavy_check_mark:](./methods.go#L4325) | [:x:](./sdl_functions_js.go#L10671) |
| [SDL_GetWindowRelativeMouseMode](https://wiki.libsdl.org/SDL3/SDL_GetWindowRelativeMouseMode) | [:heavy_check_mark:](./methods.go#L4335) | [:x:](./sdl_functions_js.go#L10689) |
| [SDL_CaptureMouse](https://wiki.libsdl.org/SDL3/SDL_CaptureMouse) | [:heavy_check_mark:](./functions.go#L935) | [:x:](./sdl_functions_js.go#L10705) |
| [SDL_CreateCursor](https://wiki.libsdl.org/SDL3/SDL_CreateCursor) | [:heavy_check_mark:](./functions.go#L945) | [:x:](./sdl_functions_js.go#L10718) |
| [SDL_CreateColorCursor](https://wiki.libsdl.org/SDL3/SDL_CreateColorCursor) | [:x:](./methods.go#L1679) | [:x:](./sdl_functions_js.go#L10750) |
| [SDL_CreateSystemCursor](https://wiki.libsdl.org/SDL3/SDL_CreateSystemCursor) | [:question:]() | [:question:](./sdl_functions_js.go#L10773) |
| [SDL_SetCursor](https://wiki.libsdl.org/SDL3/SDL_SetCursor) | [:heavy_check_mark:](./functions.go#L960) | [:x:](./sdl_functions_js.go#L10789) |
| [SDL_GetCursor](https://wiki.libsdl.org/SDL3/SDL_GetCursor) | [:heavy_check_mark:](./functions.go#L970) | [:x:](./sdl_functions_js.go#L10805) |
| [SDL_GetDefaultCursor](https://wiki.libsdl.org/SDL3/SDL_GetDefaultCursor) | [:heavy_check_mark:](./functions.go#L976) | [:x:](./sdl_functions_js.go#L10819) |
| [SDL_DestroyCursor](https://wiki.libsdl.org/SDL3/SDL_DestroyCursor) | [:x:](./methods.go#L579) | [:x:](./sdl_functions_js.go#L10833) |
| [SDL_ShowCursor](https://wiki.libsdl.org/SDL3/SDL_ShowCursor) | [:heavy_check_mark:](./functions.go#L987) | [:x:](./sdl_functions_js.go#L10847) |
| [SDL_HideCursor](https://wiki.libsdl.org/SDL3/SDL_HideCursor) | [:heavy_check_mark:](./functions.go#L997) | [:x:](./sdl_functions_js.go#L10858) |
| [SDL_CursorVisible](https://wiki.libsdl.org/SDL3/SDL_CursorVisible) | [:heavy_check_mark:](./functions.go#L1007) | [:x:](./sdl_functions_js.go#L10869) |
### Touch

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetTouchDevices](https://wiki.libsdl.org/SDL3/SDL_GetTouchDevices) | [:heavy_check_mark:](./functions.go#L1015) | [:x:](./sdl_functions_js.go#L10880) |
| [SDL_GetTouchDeviceName](https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceName) | [:heavy_check_mark:](./methods.go#L63) | [:x:](./sdl_functions_js.go#L10896) |
| [SDL_GetTouchDeviceType](https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceType) | [:heavy_check_mark:](./methods.go#L74) | [:x:](./sdl_functions_js.go#L10909) |
| [SDL_GetTouchFingers](https://wiki.libsdl.org/SDL3/SDL_GetTouchFingers) | [:heavy_check_mark:](./methods.go#L80) | [:x:](./sdl_functions_js.go#L10922) |
### Gamepad

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AddGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMapping) | [:heavy_check_mark:](./functions.go#L1031) | [:x:](./sdl_functions_js.go#L8986) |
| [SDL_AddGamepadMappingsFromIO](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromIO) | [:question:]() | [:question:](./sdl_functions_js.go#L8999) |
| [SDL_AddGamepadMappingsFromFile](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromFile) | [:heavy_check_mark:](./functions.go#L1041) | [:x:](./sdl_functions_js.go#L9017) |
| [SDL_ReloadGamepadMappings](https://wiki.libsdl.org/SDL3/SDL_ReloadGamepadMappings) | [:heavy_check_mark:](./functions.go#L1051) | [:x:](./sdl_functions_js.go#L9030) |
| [SDL_GetGamepadMappings](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappings) | [:heavy_check_mark:](./functions.go#L1061) | [:x:](./sdl_functions_js.go#L9041) |
| [SDL_GetGamepadMappingForGUID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForGUID) | [:question:]() | [:question:](./sdl_functions_js.go#L9057) |
| [SDL_GetGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMapping) | [:x:](./methods.go#L2130) | [:x:](./sdl_functions_js.go#L9070) |
| [SDL_SetGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_SetGamepadMapping) | [:x:](./methods.go#L735) | [:x:](./sdl_functions_js.go#L9086) |
| [SDL_HasGamepad](https://wiki.libsdl.org/SDL3/SDL_HasGamepad) | [:heavy_check_mark:](./functions.go#L1075) | [:x:](./sdl_functions_js.go#L9101) |
| [SDL_GetGamepads](https://wiki.libsdl.org/SDL3/SDL_GetGamepads) | [:heavy_check_mark:](./functions.go#L1081) | [:x:](./sdl_functions_js.go#L9112) |
| [SDL_IsGamepad](https://wiki.libsdl.org/SDL3/SDL_IsGamepad) | [:x:](./methods.go#L742) | [:x:](./sdl_functions_js.go#L9128) |
| [SDL_GetGamepadNameForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadNameForID) | [:x:](./methods.go#L749) | [:x:](./sdl_functions_js.go#L9141) |
| [SDL_GetGamepadPathForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPathForID) | [:x:](./methods.go#L756) | [:x:](./sdl_functions_js.go#L9154) |
| [SDL_GetGamepadPlayerIndexForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndexForID) | [:x:](./methods.go#L763) | [:x:](./sdl_functions_js.go#L9167) |
| [SDL_GetGamepadGUIDForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadGUIDForID) | [:x:](./methods.go#L770) | [:x:](./sdl_functions_js.go#L9180) |
| [SDL_GetGamepadVendorForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendorForID) | [:x:](./methods.go#L777) | [:x:](./sdl_functions_js.go#L9193) |
| [SDL_GetGamepadProductForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductForID) | [:x:](./methods.go#L784) | [:x:](./sdl_functions_js.go#L9206) |
| [SDL_GetGamepadProductVersionForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersionForID) | [:x:](./methods.go#L791) | [:x:](./sdl_functions_js.go#L9219) |
| [SDL_GetGamepadTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeForID) | [:x:](./methods.go#L798) | [:x:](./sdl_functions_js.go#L9232) |
| [SDL_GetRealGamepadTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadTypeForID) | [:x:](./methods.go#L805) | [:x:](./sdl_functions_js.go#L9245) |
| [SDL_GetGamepadMappingForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForID) | [:x:](./methods.go#L812) | [:x:](./sdl_functions_js.go#L9258) |
| [SDL_OpenGamepad](https://wiki.libsdl.org/SDL3/SDL_OpenGamepad) | [:x:](./methods.go#L819) | [:x:](./sdl_functions_js.go#L9271) |
| [SDL_GetGamepadFromID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromID) | [:x:](./methods.go#L826) | [:x:](./sdl_functions_js.go#L9287) |
| [SDL_GetGamepadFromPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromPlayerIndex) | [:heavy_check_mark:](./functions.go#L1095) | [:x:](./sdl_functions_js.go#L9303) |
| [SDL_GetGamepadProperties](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProperties) | [:x:](./methods.go#L2137) | [:x:](./sdl_functions_js.go#L9319) |
| [SDL_GetGamepadID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadID) | [:x:](./methods.go#L2144) | [:x:](./sdl_functions_js.go#L9335) |
| [SDL_GetGamepadName](https://wiki.libsdl.org/SDL3/SDL_GetGamepadName) | [:x:](./methods.go#L2151) | [:x:](./sdl_functions_js.go#L9351) |
| [SDL_GetGamepadPath](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPath) | [:x:](./methods.go#L2158) | [:x:](./sdl_functions_js.go#L9367) |
| [SDL_GetGamepadType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadType) | [:x:](./methods.go#L2165) | [:x:](./sdl_functions_js.go#L9383) |
| [SDL_GetRealGamepadType](https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadType) | [:x:](./methods.go#L2172) | [:x:](./sdl_functions_js.go#L9399) |
| [SDL_GetGamepadPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndex) | [:x:](./methods.go#L2179) | [:x:](./sdl_functions_js.go#L9415) |
| [SDL_SetGamepadPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_SetGamepadPlayerIndex) | [:x:](./methods.go#L2186) | [:x:](./sdl_functions_js.go#L9431) |
| [SDL_GetGamepadVendor](https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendor) | [:x:](./methods.go#L2193) | [:x:](./sdl_functions_js.go#L9449) |
| [SDL_GetGamepadProduct](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProduct) | [:x:](./methods.go#L2200) | [:x:](./sdl_functions_js.go#L9465) |
| [SDL_GetGamepadProductVersion](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersion) | [:x:](./methods.go#L2207) | [:x:](./sdl_functions_js.go#L9481) |
| [SDL_GetGamepadFirmwareVersion](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFirmwareVersion) | [:x:](./methods.go#L2214) | [:x:](./sdl_functions_js.go#L9497) |
| [SDL_GetGamepadSerial](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSerial) | [:x:](./methods.go#L2221) | [:x:](./sdl_functions_js.go#L9513) |
| [SDL_GetGamepadSteamHandle](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSteamHandle) | [:x:](./methods.go#L2228) | [:x:](./sdl_functions_js.go#L9529) |
| [SDL_GetGamepadConnectionState](https://wiki.libsdl.org/SDL3/SDL_GetGamepadConnectionState) | [:x:](./methods.go#L2235) | [:x:](./sdl_functions_js.go#L9545) |
| [SDL_GetGamepadPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPowerInfo) | [:x:](./methods.go#L2242) | [:x:](./sdl_functions_js.go#L9561) |
| [SDL_GamepadConnected](https://wiki.libsdl.org/SDL3/SDL_GamepadConnected) | [:x:](./methods.go#L2249) | [:x:](./sdl_functions_js.go#L9582) |
| [SDL_GetGamepadJoystick](https://wiki.libsdl.org/SDL3/SDL_GetGamepadJoystick) | [:x:](./methods.go#L2256) | [:x:](./sdl_functions_js.go#L9598) |
| [SDL_SetGamepadEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_SetGamepadEventsEnabled) | [:heavy_check_mark:](./functions.go#L1101) | [:x:](./sdl_functions_js.go#L9617) |
| [SDL_GamepadEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_GamepadEventsEnabled) | [:heavy_check_mark:](./functions.go#L1107) | [:x:](./sdl_functions_js.go#L9628) |
| [SDL_GetGamepadBindings](https://wiki.libsdl.org/SDL3/SDL_GetGamepadBindings) | [:heavy_check_mark:](./methods.go#L2263) | [:x:](./sdl_functions_js.go#L9639) |
| [SDL_UpdateGamepads](https://wiki.libsdl.org/SDL3/SDL_UpdateGamepads) | [:heavy_check_mark:](./functions.go#L1113) | [:x:](./sdl_functions_js.go#L9660) |
| [SDL_GetGamepadTypeFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeFromString) | [:question:]() | [:question:](./sdl_functions_js.go#L9669) |
| [SDL_GetGamepadStringForType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForType) | [:heavy_check_mark:](./methods.go#L5214) | [:x:](./sdl_functions_js.go#L9682) |
| [SDL_GetGamepadAxisFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxisFromString) | [:heavy_check_mark:](./functions.go#L1119) | [:x:](./sdl_functions_js.go#L9695) |
| [SDL_GetGamepadStringForAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForAxis) | [:x:](./methods.go#L570) | [:x:](./sdl_functions_js.go#L9708) |
| [SDL_GamepadHasAxis](https://wiki.libsdl.org/SDL3/SDL_GamepadHasAxis) | [:x:](./methods.go#L2277) | [:x:](./sdl_functions_js.go#L9721) |
| [SDL_GetGamepadAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxis) | [:x:](./methods.go#L2284) | [:x:](./sdl_functions_js.go#L9739) |
| [SDL_GetGamepadButtonFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonFromString) | [:heavy_check_mark:](./functions.go#L1125) | [:x:](./sdl_functions_js.go#L9757) |
| [SDL_GetGamepadStringForButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForButton) | [:heavy_check_mark:](./methods.go#L437) | [:x:](./sdl_functions_js.go#L9770) |
| [SDL_GamepadHasButton](https://wiki.libsdl.org/SDL3/SDL_GamepadHasButton) | [:x:](./methods.go#L2291) | [:x:](./sdl_functions_js.go#L9783) |
| [SDL_GetGamepadButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButton) | [:x:](./methods.go#L2298) | [:x:](./sdl_functions_js.go#L9801) |
| [SDL_GetGamepadButtonLabelForType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabelForType) | [:heavy_check_mark:](./methods.go#L5220) | [:x:](./sdl_functions_js.go#L9819) |
| [SDL_GetGamepadButtonLabel](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabel) | [:x:](./methods.go#L2305) | [:x:](./sdl_functions_js.go#L9834) |
| [SDL_GetNumGamepadTouchpads](https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpads) | [:x:](./methods.go#L2312) | [:x:](./sdl_functions_js.go#L9852) |
| [SDL_GetNumGamepadTouchpadFingers](https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpadFingers) | [:x:](./methods.go#L2319) | [:x:](./sdl_functions_js.go#L9868) |
| [SDL_GetGamepadTouchpadFinger](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTouchpadFinger) | [:x:](./methods.go#L2326) | [:x:](./sdl_functions_js.go#L9886) |
| [SDL_GamepadHasSensor](https://wiki.libsdl.org/SDL3/SDL_GamepadHasSensor) | [:x:](./methods.go#L2333) | [:x:](./sdl_functions_js.go#L9926) |
| [SDL_SetGamepadSensorEnabled](https://wiki.libsdl.org/SDL3/SDL_SetGamepadSensorEnabled) | [:x:](./methods.go#L2340) | [:x:](./sdl_functions_js.go#L9944) |
| [SDL_GamepadSensorEnabled](https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEnabled) | [:x:](./methods.go#L2347) | [:x:](./sdl_functions_js.go#L9964) |
| [SDL_GetGamepadSensorDataRate](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorDataRate) | [:x:](./methods.go#L2354) | [:x:](./sdl_functions_js.go#L9982) |
| [SDL_GetGamepadSensorData](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorData) | [:x:](./methods.go#L2361) | [:x:](./sdl_functions_js.go#L10000) |
| [SDL_RumbleGamepad](https://wiki.libsdl.org/SDL3/SDL_RumbleGamepad) | [:x:](./methods.go#L2368) | [:x:](./sdl_functions_js.go#L10025) |
| [SDL_RumbleGamepadTriggers](https://wiki.libsdl.org/SDL3/SDL_RumbleGamepadTriggers) | [:x:](./methods.go#L2375) | [:x:](./sdl_functions_js.go#L10047) |
| [SDL_SetGamepadLED](https://wiki.libsdl.org/SDL3/SDL_SetGamepadLED) | [:x:](./methods.go#L2382) | [:x:](./sdl_functions_js.go#L10069) |
| [SDL_SendGamepadEffect](https://wiki.libsdl.org/SDL3/SDL_SendGamepadEffect) | [:x:](./methods.go#L2389) | [:x:](./sdl_functions_js.go#L10091) |
| [SDL_CloseGamepad](https://wiki.libsdl.org/SDL3/SDL_CloseGamepad) | [:x:](./methods.go#L2396) | [:x:](./sdl_functions_js.go#L10111) |
| [SDL_GetGamepadAppleSFSymbolsNameForButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForButton) | [:x:](./methods.go#L2403) | [:x:](./sdl_functions_js.go#L10125) |
| [SDL_GetGamepadAppleSFSymbolsNameForAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForAxis) | [:x:](./methods.go#L2410) | [:x:](./sdl_functions_js.go#L10143) |
### Joystick

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_LockJoysticks](https://wiki.libsdl.org/SDL3/SDL_LockJoysticks) | [:heavy_check_mark:](./functions.go#L1133) | [:x:](./sdl_functions_js.go#L8047) |
| [SDL_UnlockJoysticks](https://wiki.libsdl.org/SDL3/SDL_UnlockJoysticks) | [:heavy_check_mark:](./functions.go#L1139) | [:x:](./sdl_functions_js.go#L8056) |
| [SDL_HasJoystick](https://wiki.libsdl.org/SDL3/SDL_HasJoystick) | [:heavy_check_mark:](./functions.go#L1145) | [:x:](./sdl_functions_js.go#L8065) |
| [SDL_GetJoysticks](https://wiki.libsdl.org/SDL3/SDL_GetJoysticks) | [:heavy_check_mark:](./functions.go#L1151) | [:x:](./sdl_functions_js.go#L8076) |
| [SDL_GetJoystickNameForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickNameForID) | [:x:](./methods.go#L643) | [:x:](./sdl_functions_js.go#L8092) |
| [SDL_GetJoystickPathForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPathForID) | [:x:](./methods.go#L650) | [:x:](./sdl_functions_js.go#L8105) |
| [SDL_GetJoystickPlayerIndexForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndexForID) | [:x:](./methods.go#L657) | [:x:](./sdl_functions_js.go#L8118) |
| [SDL_GetJoystickGUIDForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDForID) | [:x:](./methods.go#L664) | [:x:](./sdl_functions_js.go#L8131) |
| [SDL_GetJoystickVendorForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendorForID) | [:x:](./methods.go#L671) | [:x:](./sdl_functions_js.go#L8144) |
| [SDL_GetJoystickProductForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductForID) | [:x:](./methods.go#L678) | [:x:](./sdl_functions_js.go#L8157) |
| [SDL_GetJoystickProductVersionForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersionForID) | [:x:](./methods.go#L685) | [:x:](./sdl_functions_js.go#L8170) |
| [SDL_GetJoystickTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickTypeForID) | [:x:](./methods.go#L692) | [:x:](./sdl_functions_js.go#L8183) |
| [SDL_OpenJoystick](https://wiki.libsdl.org/SDL3/SDL_OpenJoystick) | [:heavy_check_mark:](./methods.go#L699) | [:heavy_check_mark:](./sdl_functions_js.go#L8196) |
| [SDL_GetJoystickFromID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromID) | [:heavy_check_mark:](./methods.go#L710) | [:heavy_check_mark:](./sdl_functions_js.go#L8208) |
| [SDL_GetJoystickFromPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromPlayerIndex) | [:heavy_check_mark:](./functions.go#L1165) | [:x:](./sdl_functions_js.go#L8222) |
| [SDL_AttachVirtualJoystick](https://wiki.libsdl.org/SDL3/SDL_AttachVirtualJoystick) | [:heavy_check_mark:](./functions.go#L1171) | [:x:](./sdl_functions_js.go#L8238) |
| [SDL_DetachVirtualJoystick](https://wiki.libsdl.org/SDL3/SDL_DetachVirtualJoystick) | [:x:](./methods.go#L721) | [:x:](./sdl_functions_js.go#L8254) |
| [SDL_IsJoystickVirtual](https://wiki.libsdl.org/SDL3/SDL_IsJoystickVirtual) | [:x:](./methods.go#L728) | [:x:](./sdl_functions_js.go#L8267) |
| [SDL_SetJoystickVirtualAxis](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualAxis) | [:heavy_check_mark:](./methods.go#L4872) | [:x:](./sdl_functions_js.go#L8280) |
| [SDL_SetJoystickVirtualBall](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualBall) | [:heavy_check_mark:](./methods.go#L4882) | [:x:](./sdl_functions_js.go#L8300) |
| [SDL_SetJoystickVirtualButton](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualButton) | [:heavy_check_mark:](./methods.go#L4892) | [:x:](./sdl_functions_js.go#L8322) |
| [SDL_SetJoystickVirtualHat](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualHat) | [:heavy_check_mark:](./methods.go#L4902) | [:x:](./sdl_functions_js.go#L8342) |
| [SDL_SetJoystickVirtualTouchpad](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualTouchpad) | [:heavy_check_mark:](./methods.go#L4912) | [:x:](./sdl_functions_js.go#L8362) |
| [SDL_SendJoystickVirtualSensorData](https://wiki.libsdl.org/SDL3/SDL_SendJoystickVirtualSensorData) | [:heavy_check_mark:](./methods.go#L4922) | [:x:](./sdl_functions_js.go#L8390) |
| [SDL_GetJoystickProperties](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProperties) | [:heavy_check_mark:](./methods.go#L4932) | [:x:](./sdl_functions_js.go#L8417) |
| [SDL_GetJoystickName](https://wiki.libsdl.org/SDL3/SDL_GetJoystickName) | [:heavy_check_mark:](./methods.go#L4943) | [:heavy_check_mark:](./sdl_functions_js.go#L8433) |
| [SDL_GetJoystickPath](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPath) | [:heavy_check_mark:](./methods.go#L4954) | [:x:](./sdl_functions_js.go#L8446) |
| [SDL_GetJoystickPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndex) | [:heavy_check_mark:](./methods.go#L4965) | [:x:](./sdl_functions_js.go#L8462) |
| [SDL_SetJoystickPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_SetJoystickPlayerIndex) | [:heavy_check_mark:](./methods.go#L4971) | [:x:](./sdl_functions_js.go#L8478) |
| [SDL_GetJoystickGUID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUID) | [:x:](./methods.go#L4981) | [:x:](./sdl_functions_js.go#L8496) |
| [SDL_GetJoystickVendor](https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendor) | [:heavy_check_mark:](./methods.go#L4988) | [:x:](./sdl_functions_js.go#L8512) |
| [SDL_GetJoystickProduct](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProduct) | [:heavy_check_mark:](./methods.go#L4994) | [:x:](./sdl_functions_js.go#L8528) |
| [SDL_GetJoystickProductVersion](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersion) | [:heavy_check_mark:](./methods.go#L5000) | [:x:](./sdl_functions_js.go#L8544) |
| [SDL_GetJoystickFirmwareVersion](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFirmwareVersion) | [:heavy_check_mark:](./methods.go#L5006) | [:x:](./sdl_functions_js.go#L8560) |
| [SDL_GetJoystickSerial](https://wiki.libsdl.org/SDL3/SDL_GetJoystickSerial) | [:heavy_check_mark:](./methods.go#L5012) | [:x:](./sdl_functions_js.go#L8576) |
| [SDL_GetJoystickType](https://wiki.libsdl.org/SDL3/SDL_GetJoystickType) | [:heavy_check_mark:](./methods.go#L5018) | [:x:](./sdl_functions_js.go#L8592) |
| [SDL_GetJoystickGUIDInfo](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDInfo) | [:question:]() | [:question:](./sdl_functions_js.go#L8608) |
| [SDL_JoystickConnected](https://wiki.libsdl.org/SDL3/SDL_JoystickConnected) | [:heavy_check_mark:](./methods.go#L5024) | [:x:](./sdl_functions_js.go#L8639) |
| [SDL_GetJoystickID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickID) | [:heavy_check_mark:](./methods.go#L5030) | [:heavy_check_mark:](./sdl_functions_js.go#L8655) |
| [SDL_GetNumJoystickAxes](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickAxes) | [:heavy_check_mark:](./methods.go#L5041) | [:heavy_check_mark:](./sdl_functions_js.go#L8668) |
| [SDL_GetNumJoystickBalls](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickBalls) | [:heavy_check_mark:](./methods.go#L5052) | [:x:](./sdl_functions_js.go#L8681) |
| [SDL_GetNumJoystickHats](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickHats) | [:heavy_check_mark:](./methods.go#L5063) | [:heavy_check_mark:](./sdl_functions_js.go#L8697) |
| [SDL_GetNumJoystickButtons](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickButtons) | [:heavy_check_mark:](./methods.go#L5074) | [:heavy_check_mark:](./sdl_functions_js.go#L8710) |
| [SDL_SetJoystickEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_SetJoystickEventsEnabled) | [:heavy_check_mark:](./functions.go#L1177) | [:x:](./sdl_functions_js.go#L8723) |
| [SDL_JoystickEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_JoystickEventsEnabled) | [:heavy_check_mark:](./functions.go#L1183) | [:x:](./sdl_functions_js.go#L8734) |
| [SDL_UpdateJoysticks](https://wiki.libsdl.org/SDL3/SDL_UpdateJoysticks) | [:heavy_check_mark:](./functions.go#L1189) | [:x:](./sdl_functions_js.go#L8745) |
| [SDL_GetJoystickAxis](https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxis) | [:heavy_check_mark:](./methods.go#L5085) | [:heavy_check_mark:](./sdl_functions_js.go#L8754) |
| [SDL_GetJoystickAxisInitialState](https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxisInitialState) | [:heavy_check_mark:](./methods.go#L5096) | [:x:](./sdl_functions_js.go#L8769) |
| [SDL_GetJoystickBall](https://wiki.libsdl.org/SDL3/SDL_GetJoystickBall) | [:heavy_check_mark:](./methods.go#L5105) | [:x:](./sdl_functions_js.go#L8792) |
| [SDL_GetJoystickHat](https://wiki.libsdl.org/SDL3/SDL_GetJoystickHat) | [:heavy_check_mark:](./methods.go#L5117) | [:heavy_check_mark:](./sdl_functions_js.go#L8820) |
| [SDL_GetJoystickButton](https://wiki.libsdl.org/SDL3/SDL_GetJoystickButton) | [:heavy_check_mark:](./methods.go#L5123) | [:heavy_check_mark:](./sdl_functions_js.go#L8835) |
| [SDL_RumbleJoystick](https://wiki.libsdl.org/SDL3/SDL_RumbleJoystick) | [:heavy_check_mark:](./methods.go#L5129) | [:x:](./sdl_functions_js.go#L8850) |
| [SDL_RumbleJoystickTriggers](https://wiki.libsdl.org/SDL3/SDL_RumbleJoystickTriggers) | [:heavy_check_mark:](./methods.go#L5135) | [:x:](./sdl_functions_js.go#L8872) |
| [SDL_SetJoystickLED](https://wiki.libsdl.org/SDL3/SDL_SetJoystickLED) | [:heavy_check_mark:](./methods.go#L5145) | [:x:](./sdl_functions_js.go#L8894) |
| [SDL_SendJoystickEffect](https://wiki.libsdl.org/SDL3/SDL_SendJoystickEffect) | [:heavy_check_mark:](./methods.go#L5155) | [:x:](./sdl_functions_js.go#L8916) |
| [SDL_CloseJoystick](https://wiki.libsdl.org/SDL3/SDL_CloseJoystick) | [:heavy_check_mark:](./methods.go#L5165) | [:heavy_check_mark:](./sdl_functions_js.go#L8936) |
| [SDL_GetJoystickConnectionState](https://wiki.libsdl.org/SDL3/SDL_GetJoystickConnectionState) | [:heavy_check_mark:](./methods.go#L5171) | [:x:](./sdl_functions_js.go#L8949) |
| [SDL_GetJoystickPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPowerInfo) | [:heavy_check_mark:](./methods.go#L5182) | [:x:](./sdl_functions_js.go#L8965) |
### Haptic

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetHaptics](https://wiki.libsdl.org/SDL3/SDL_GetHaptics) | [:question:]() | [:question:](./sdl_functions_js.go#L13276) |
| [SDL_GetHapticNameForID](https://wiki.libsdl.org/SDL3/SDL_GetHapticNameForID) | [:x:](./methods.go#L2479) | [:x:](./sdl_functions_js.go#L13292) |
| [SDL_OpenHaptic](https://wiki.libsdl.org/SDL3/SDL_OpenHaptic) | [:x:](./methods.go#L2486) | [:x:](./sdl_functions_js.go#L13305) |
| [SDL_GetHapticFromID](https://wiki.libsdl.org/SDL3/SDL_GetHapticFromID) | [:x:](./methods.go#L2493) | [:x:](./sdl_functions_js.go#L13321) |
| [SDL_GetHapticID](https://wiki.libsdl.org/SDL3/SDL_GetHapticID) | [:x:](./methods.go#L1958) | [:x:](./sdl_functions_js.go#L13337) |
| [SDL_GetHapticName](https://wiki.libsdl.org/SDL3/SDL_GetHapticName) | [:x:](./methods.go#L1965) | [:x:](./sdl_functions_js.go#L13353) |
| [SDL_IsMouseHaptic](https://wiki.libsdl.org/SDL3/SDL_IsMouseHaptic) | [:question:]() | [:question:](./sdl_functions_js.go#L13369) |
| [SDL_OpenHapticFromMouse](https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromMouse) | [:question:]() | [:question:](./sdl_functions_js.go#L13380) |
| [SDL_IsJoystickHaptic](https://wiki.libsdl.org/SDL3/SDL_IsJoystickHaptic) | [:heavy_check_mark:](./methods.go#L5195) | [:x:](./sdl_functions_js.go#L13394) |
| [SDL_OpenHapticFromJoystick](https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromJoystick) | [:heavy_check_mark:](./methods.go#L5201) | [:x:](./sdl_functions_js.go#L13410) |
| [SDL_CloseHaptic](https://wiki.libsdl.org/SDL3/SDL_CloseHaptic) | [:x:](./methods.go#L1972) | [:x:](./sdl_functions_js.go#L13429) |
| [SDL_GetMaxHapticEffects](https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffects) | [:x:](./methods.go#L1979) | [:x:](./sdl_functions_js.go#L13443) |
| [SDL_GetMaxHapticEffectsPlaying](https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffectsPlaying) | [:x:](./methods.go#L1986) | [:x:](./sdl_functions_js.go#L13459) |
| [SDL_GetHapticFeatures](https://wiki.libsdl.org/SDL3/SDL_GetHapticFeatures) | [:x:](./methods.go#L1993) | [:x:](./sdl_functions_js.go#L13475) |
| [SDL_GetNumHapticAxes](https://wiki.libsdl.org/SDL3/SDL_GetNumHapticAxes) | [:x:](./methods.go#L2000) | [:x:](./sdl_functions_js.go#L13491) |
| [SDL_HapticEffectSupported](https://wiki.libsdl.org/SDL3/SDL_HapticEffectSupported) | [:x:](./methods.go#L2007) | [:x:](./sdl_functions_js.go#L13507) |
| [SDL_CreateHapticEffect](https://wiki.libsdl.org/SDL3/SDL_CreateHapticEffect) | [:x:](./methods.go#L2014) | [:x:](./sdl_functions_js.go#L13528) |
| [SDL_UpdateHapticEffect](https://wiki.libsdl.org/SDL3/SDL_UpdateHapticEffect) | [:x:](./methods.go#L2021) | [:x:](./sdl_functions_js.go#L13549) |
| [SDL_RunHapticEffect](https://wiki.libsdl.org/SDL3/SDL_RunHapticEffect) | [:x:](./methods.go#L2028) | [:x:](./sdl_functions_js.go#L13572) |
| [SDL_StopHapticEffect](https://wiki.libsdl.org/SDL3/SDL_StopHapticEffect) | [:x:](./methods.go#L2035) | [:x:](./sdl_functions_js.go#L13592) |
| [SDL_DestroyHapticEffect](https://wiki.libsdl.org/SDL3/SDL_DestroyHapticEffect) | [:x:](./methods.go#L2042) | [:x:](./sdl_functions_js.go#L13610) |
| [SDL_GetHapticEffectStatus](https://wiki.libsdl.org/SDL3/SDL_GetHapticEffectStatus) | [:x:](./methods.go#L2049) | [:x:](./sdl_functions_js.go#L13626) |
| [SDL_SetHapticGain](https://wiki.libsdl.org/SDL3/SDL_SetHapticGain) | [:x:](./methods.go#L2056) | [:x:](./sdl_functions_js.go#L13644) |
| [SDL_SetHapticAutocenter](https://wiki.libsdl.org/SDL3/SDL_SetHapticAutocenter) | [:x:](./methods.go#L2063) | [:x:](./sdl_functions_js.go#L13662) |
| [SDL_PauseHaptic](https://wiki.libsdl.org/SDL3/SDL_PauseHaptic) | [:x:](./methods.go#L2070) | [:x:](./sdl_functions_js.go#L13680) |
| [SDL_ResumeHaptic](https://wiki.libsdl.org/SDL3/SDL_ResumeHaptic) | [:x:](./methods.go#L2077) | [:x:](./sdl_functions_js.go#L13696) |
| [SDL_StopHapticEffects](https://wiki.libsdl.org/SDL3/SDL_StopHapticEffects) | [:x:](./methods.go#L2084) | [:x:](./sdl_functions_js.go#L13712) |
| [SDL_HapticRumbleSupported](https://wiki.libsdl.org/SDL3/SDL_HapticRumbleSupported) | [:x:](./methods.go#L2091) | [:x:](./sdl_functions_js.go#L13728) |
| [SDL_InitHapticRumble](https://wiki.libsdl.org/SDL3/SDL_InitHapticRumble) | [:x:](./methods.go#L2098) | [:x:](./sdl_functions_js.go#L13744) |
| [SDL_PlayHapticRumble](https://wiki.libsdl.org/SDL3/SDL_PlayHapticRumble) | [:x:](./methods.go#L2105) | [:x:](./sdl_functions_js.go#L13760) |
| [SDL_StopHapticRumble](https://wiki.libsdl.org/SDL3/SDL_StopHapticRumble) | [:x:](./methods.go#L2112) | [:x:](./sdl_functions_js.go#L13780) |
### Audio

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumAudioDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumAudioDrivers) | [:heavy_check_mark:](./functions.go#L572) | [:x:](./sdl_functions_js.go#L2357) |
| [SDL_GetAudioDriver](https://wiki.libsdl.org/SDL3/SDL_GetAudioDriver) | [:heavy_check_mark:](./functions.go#L578) | [:x:](./sdl_functions_js.go#L2368) |
| [SDL_GetCurrentAudioDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentAudioDriver) | [:question:]() | [:question:](./sdl_functions_js.go#L2381) |
| [SDL_GetAudioPlaybackDevices](https://wiki.libsdl.org/SDL3/SDL_GetAudioPlaybackDevices) | [:heavy_check_mark:](./functions.go#L584) | [:x:](./sdl_functions_js.go#L2392) |
| [SDL_GetAudioRecordingDevices](https://wiki.libsdl.org/SDL3/SDL_GetAudioRecordingDevices) | [:heavy_check_mark:](./functions.go#L598) | [:x:](./sdl_functions_js.go#L2408) |
| [SDL_GetAudioDeviceName](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceName) | [:heavy_check_mark:](./methods.go#L220) | [:x:](./sdl_functions_js.go#L2424) |
| [SDL_GetAudioDeviceFormat](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceFormat) | [:heavy_check_mark:](./methods.go#L231) | [:x:](./sdl_functions_js.go#L2437) |
| [SDL_GetAudioDeviceChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceChannelMap) | [:heavy_check_mark:](./methods.go#L244) | [:x:](./sdl_functions_js.go#L2460) |
| [SDL_OpenAudioDevice](https://wiki.libsdl.org/SDL3/SDL_OpenAudioDevice) | [:heavy_check_mark:](./methods.go#L258) | [:x:](./sdl_functions_js.go#L2478) |
| [SDL_IsAudioDevicePhysical](https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePhysical) | [:heavy_check_mark:](./methods.go#L269) | [:x:](./sdl_functions_js.go#L2496) |
| [SDL_IsAudioDevicePlayback](https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePlayback) | [:heavy_check_mark:](./methods.go#L275) | [:x:](./sdl_functions_js.go#L2509) |
| [SDL_PauseAudioDevice](https://wiki.libsdl.org/SDL3/SDL_PauseAudioDevice) | [:heavy_check_mark:](./methods.go#L281) | [:x:](./sdl_functions_js.go#L2522) |
| [SDL_ResumeAudioDevice](https://wiki.libsdl.org/SDL3/SDL_ResumeAudioDevice) | [:heavy_check_mark:](./methods.go#L291) | [:x:](./sdl_functions_js.go#L2535) |
| [SDL_AudioDevicePaused](https://wiki.libsdl.org/SDL3/SDL_AudioDevicePaused) | [:heavy_check_mark:](./methods.go#L301) | [:x:](./sdl_functions_js.go#L2548) |
| [SDL_GetAudioDeviceGain](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceGain) | [:heavy_check_mark:](./methods.go#L307) | [:x:](./sdl_functions_js.go#L2561) |
| [SDL_SetAudioDeviceGain](https://wiki.libsdl.org/SDL3/SDL_SetAudioDeviceGain) | [:heavy_check_mark:](./methods.go#L318) | [:x:](./sdl_functions_js.go#L2574) |
| [SDL_CloseAudioDevice](https://wiki.libsdl.org/SDL3/SDL_CloseAudioDevice) | [:heavy_check_mark:](./methods.go#L328) | [:x:](./sdl_functions_js.go#L2589) |
| [SDL_BindAudioStreams](https://wiki.libsdl.org/SDL3/SDL_BindAudioStreams) | [:heavy_check_mark:](./methods.go#L334) | [:x:](./sdl_functions_js.go#L2600) |
| [SDL_BindAudioStream](https://wiki.libsdl.org/SDL3/SDL_BindAudioStream) | [:heavy_check_mark:](./methods.go#L344) | [:x:](./sdl_functions_js.go#L2620) |
| [SDL_UnbindAudioStreams](https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStreams) | [:heavy_check_mark:](./functions.go#L612) | [:x:](./sdl_functions_js.go#L2638) |
| [SDL_UnbindAudioStream](https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStream) | [:heavy_check_mark:](./methods.go#L3288) | [:x:](./sdl_functions_js.go#L2654) |
| [SDL_GetAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamDevice) | [:heavy_check_mark:](./methods.go#L3294) | [:x:](./sdl_functions_js.go#L2668) |
| [SDL_CreateAudioStream](https://wiki.libsdl.org/SDL3/SDL_CreateAudioStream) | [:heavy_check_mark:](./functions.go#L618) | [:x:](./sdl_functions_js.go#L2684) |
| [SDL_GetAudioStreamProperties](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamProperties) | [:heavy_check_mark:](./methods.go#L3300) | [:x:](./sdl_functions_js.go#L2708) |
| [SDL_GetAudioStreamFormat](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFormat) | [:x:](./methods.go#L3311) | [:x:](./sdl_functions_js.go#L2724) |
| [SDL_SetAudioStreamFormat](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFormat) | [:x:](./methods.go#L3318) | [:x:](./sdl_functions_js.go#L2750) |
| [SDL_GetAudioStreamFrequencyRatio](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFrequencyRatio) | [:heavy_check_mark:](./methods.go#L3325) | [:x:](./sdl_functions_js.go#L2776) |
| [SDL_SetAudioStreamFrequencyRatio](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFrequencyRatio) | [:heavy_check_mark:](./methods.go#L3336) | [:x:](./sdl_functions_js.go#L2792) |
| [SDL_GetAudioStreamGain](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamGain) | [:heavy_check_mark:](./methods.go#L3346) | [:x:](./sdl_functions_js.go#L2810) |
| [SDL_SetAudioStreamGain](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGain) | [:heavy_check_mark:](./methods.go#L3357) | [:x:](./sdl_functions_js.go#L2826) |
| [SDL_GetAudioStreamInputChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamInputChannelMap) | [:heavy_check_mark:](./methods.go#L3367) | [:x:](./sdl_functions_js.go#L2844) |
| [SDL_GetAudioStreamOutputChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamOutputChannelMap) | [:heavy_check_mark:](./methods.go#L3381) | [:x:](./sdl_functions_js.go#L2865) |
| [SDL_SetAudioStreamInputChannelMap](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamInputChannelMap) | [:heavy_check_mark:](./methods.go#L3395) | [:x:](./sdl_functions_js.go#L2886) |
| [SDL_SetAudioStreamOutputChannelMap](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamOutputChannelMap) | [:heavy_check_mark:](./methods.go#L3405) | [:x:](./sdl_functions_js.go#L2909) |
| [SDL_PutAudioStreamData](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData) | [:heavy_check_mark:](./methods.go#L3415) | [:x:](./sdl_functions_js.go#L2932) |
| [SDL_GetAudioStreamData](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData) | [:heavy_check_mark:](./methods.go#L3425) | [:x:](./sdl_functions_js.go#L2952) |
| [SDL_GetAudioStreamAvailable](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamAvailable) | [:heavy_check_mark:](./methods.go#L3436) | [:x:](./sdl_functions_js.go#L2972) |
| [SDL_GetAudioStreamQueued](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamQueued) | [:heavy_check_mark:](./methods.go#L3447) | [:x:](./sdl_functions_js.go#L2988) |
| [SDL_FlushAudioStream](https://wiki.libsdl.org/SDL3/SDL_FlushAudioStream) | [:heavy_check_mark:](./methods.go#L3458) | [:x:](./sdl_functions_js.go#L3004) |
| [SDL_ClearAudioStream](https://wiki.libsdl.org/SDL3/SDL_ClearAudioStream) | [:heavy_check_mark:](./methods.go#L3468) | [:x:](./sdl_functions_js.go#L3020) |
| [SDL_PauseAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_PauseAudioStreamDevice) | [:heavy_check_mark:](./methods.go#L3478) | [:x:](./sdl_functions_js.go#L3036) |
| [SDL_ResumeAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_ResumeAudioStreamDevice) | [:heavy_check_mark:](./methods.go#L3488) | [:x:](./sdl_functions_js.go#L3052) |
| [SDL_AudioStreamDevicePaused](https://wiki.libsdl.org/SDL3/SDL_AudioStreamDevicePaused) | [:heavy_check_mark:](./methods.go#L3498) | [:x:](./sdl_functions_js.go#L3068) |
| [SDL_LockAudioStream](https://wiki.libsdl.org/SDL3/SDL_LockAudioStream) | [:heavy_check_mark:](./methods.go#L3504) | [:x:](./sdl_functions_js.go#L3084) |
| [SDL_UnlockAudioStream](https://wiki.libsdl.org/SDL3/SDL_UnlockAudioStream) | [:heavy_check_mark:](./methods.go#L3514) | [:x:](./sdl_functions_js.go#L3100) |
| [SDL_SetAudioStreamGetCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGetCallback) | [:x:](./methods.go#L3524) | [:x:](./sdl_functions_js.go#L3116) |
| [SDL_SetAudioStreamPutCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamPutCallback) | [:x:](./methods.go#L3531) | [:x:](./sdl_functions_js.go#L3136) |
| [SDL_DestroyAudioStream](https://wiki.libsdl.org/SDL3/SDL_DestroyAudioStream) | [:x:](./methods.go#L3538) | [:x:](./sdl_functions_js.go#L3156) |
| [SDL_OpenAudioDeviceStream](https://wiki.libsdl.org/SDL3/SDL_OpenAudioDeviceStream) | [:x:](./methods.go#L354) | [:x:](./sdl_functions_js.go#L3170) |
| [SDL_SetAudioPostmixCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioPostmixCallback) | [:x:](./methods.go#L361) | [:x:](./sdl_functions_js.go#L3194) |
| [SDL_LoadWAV_IO](https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO) | [:heavy_check_mark:](./functions.go#L630) | [:x:](./sdl_functions_js.go#L3211) |
| [SDL_LoadWAV](https://wiki.libsdl.org/SDL3/SDL_LoadWAV) | [:heavy_check_mark:](./functions.go#L644) | [:x:](./sdl_functions_js.go#L3244) |
| [SDL_MixAudio](https://wiki.libsdl.org/SDL3/SDL_MixAudio) | [:heavy_check_mark:](./functions.go#L658) | [:x:](./sdl_functions_js.go#L3272) |
| [SDL_ConvertAudioSamples](https://wiki.libsdl.org/SDL3/SDL_ConvertAudioSamples) | [:heavy_check_mark:](./functions.go#L669) | [:x:](./sdl_functions_js.go#L3299) |
| [SDL_GetAudioFormatName](https://wiki.libsdl.org/SDL3/SDL_GetAudioFormatName) | [:x:](./methods.go#L1128) | [:x:](./sdl_functions_js.go#L3337) |
| [SDL_GetSilenceValueForFormat](https://wiki.libsdl.org/SDL3/SDL_GetSilenceValueForFormat) | [:x:](./methods.go#L1135) | [:x:](./sdl_functions_js.go#L3350) |
### Time

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetDateTimeLocalePreferences](https://wiki.libsdl.org/SDL3/SDL_GetDateTimeLocalePreferences) | [:x:](./methods.go#L3033) | [:x:](./sdl_functions_js.go#L16982) |
| [SDL_GetCurrentTime](https://wiki.libsdl.org/SDL3/SDL_GetCurrentTime) | [:x:](./methods.go#L3042) | [:x:](./sdl_functions_js.go#L17003) |
| [SDL_TimeToDateTime](https://wiki.libsdl.org/SDL3/SDL_TimeToDateTime) | [:x:](./methods.go#L5263) | [:x:](./sdl_functions_js.go#L17019) |
| [SDL_DateTimeToTime](https://wiki.libsdl.org/SDL3/SDL_DateTimeToTime) | [:heavy_check_mark:](./methods.go#L5529) | [:x:](./sdl_functions_js.go#L17039) |
| [SDL_TimeToWindows](https://wiki.libsdl.org/SDL3/SDL_TimeToWindows) | [:x:](./methods.go#L5270) | [:x:](./sdl_functions_js.go#L17060) |
| [SDL_TimeFromWindows](https://wiki.libsdl.org/SDL3/SDL_TimeFromWindows) | [:question:]() | [:question:](./sdl_functions_js.go#L17081) |
| [SDL_GetDaysInMonth](https://wiki.libsdl.org/SDL3/SDL_GetDaysInMonth) | [:question:]() | [:question:](./sdl_functions_js.go#L17096) |
| [SDL_GetDayOfYear](https://wiki.libsdl.org/SDL3/SDL_GetDayOfYear) | [:question:]() | [:question:](./sdl_functions_js.go#L17111) |
| [SDL_GetDayOfWeek](https://wiki.libsdl.org/SDL3/SDL_GetDayOfWeek) | [:question:]() | [:question:](./sdl_functions_js.go#L17128) |
### Timer

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetTicks](https://wiki.libsdl.org/SDL3/SDL_GetTicks) | [:heavy_check_mark:](./functions.go#L691) | [:heavy_check_mark:](./sdl_functions_js.go#L17145) |
| [SDL_GetTicksNS](https://wiki.libsdl.org/SDL3/SDL_GetTicksNS) | [:question:]() | [:question:](./sdl_functions_js.go#L17153) |
| [SDL_GetPerformanceCounter](https://wiki.libsdl.org/SDL3/SDL_GetPerformanceCounter) | [:question:]() | [:question:](./sdl_functions_js.go#L17164) |
| [SDL_GetPerformanceFrequency](https://wiki.libsdl.org/SDL3/SDL_GetPerformanceFrequency) | [:question:]() | [:question:](./sdl_functions_js.go#L17175) |
| [SDL_Delay](https://wiki.libsdl.org/SDL3/SDL_Delay) | [:question:]() | [:question:](./sdl_functions_js.go#L17186) |
| [SDL_DelayNS](https://wiki.libsdl.org/SDL3/SDL_DelayNS) | [:question:]() | [:question:](./sdl_functions_js.go#L17197) |
| [SDL_DelayPrecise](https://wiki.libsdl.org/SDL3/SDL_DelayPrecise) | [:question:]() | [:question:](./sdl_functions_js.go#L17208) |
| [SDL_AddTimer](https://wiki.libsdl.org/SDL3/SDL_AddTimer) | [:question:]() | [:question:](./sdl_functions_js.go#L17219) |
| [SDL_AddTimerNS](https://wiki.libsdl.org/SDL3/SDL_AddTimerNS) | [:question:]() | [:question:](./sdl_functions_js.go#L17236) |
| [SDL_RemoveTimer](https://wiki.libsdl.org/SDL3/SDL_RemoveTimer) | [:x:](./methods.go#L4423) | [:x:](./sdl_functions_js.go#L17253) |
### Render

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumRenderDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumRenderDrivers) | [:heavy_check_mark:](./functions.go#L276) | [:x:](./sdl_functions_js.go#L14603) |
| [SDL_GetRenderDriver](https://wiki.libsdl.org/SDL3/SDL_GetRenderDriver) | [:heavy_check_mark:](./functions.go#L282) | [:x:](./sdl_functions_js.go#L14614) |
| [SDL_CreateWindowAndRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateWindowAndRenderer) | [:heavy_check_mark:](./functions.go#L288) | [:heavy_check_mark:](./sdl_functions_js.go#L14627) |
| [SDL_CreateRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateRenderer) | [:heavy_check_mark:](./methods.go#L4348) | [:x:](./sdl_functions_js.go#L14655) |
| [SDL_CreateRendererWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateRendererWithProperties) | [:heavy_check_mark:](./functions.go#L301) | [:x:](./sdl_functions_js.go#L14676) |
| [SDL_CreateSoftwareRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateSoftwareRenderer) | [:x:](./methods.go#L1686) | [:x:](./sdl_functions_js.go#L14692) |
| [SDL_GetRenderer](https://wiki.libsdl.org/SDL3/SDL_GetRenderer) | [:heavy_check_mark:](./methods.go#L4359) | [:x:](./sdl_functions_js.go#L14711) |
| [SDL_GetRenderWindow](https://wiki.libsdl.org/SDL3/SDL_GetRenderWindow) | [:x:](./methods.go#L2502) | [:x:](./sdl_functions_js.go#L14730) |
| [SDL_GetRendererName](https://wiki.libsdl.org/SDL3/SDL_GetRendererName) | [:heavy_check_mark:](./methods.go#L2509) | [:x:](./sdl_functions_js.go#L14749) |
| [SDL_GetRendererProperties](https://wiki.libsdl.org/SDL3/SDL_GetRendererProperties) | [:heavy_check_mark:](./methods.go#L2520) | [:x:](./sdl_functions_js.go#L14765) |
| [SDL_GetRenderOutputSize](https://wiki.libsdl.org/SDL3/SDL_GetRenderOutputSize) | [:x:](./methods.go#L2526) | [:x:](./sdl_functions_js.go#L14781) |
| [SDL_GetCurrentRenderOutputSize](https://wiki.libsdl.org/SDL3/SDL_GetCurrentRenderOutputSize) | [:x:](./methods.go#L2533) | [:x:](./sdl_functions_js.go#L14807) |
| [SDL_CreateTexture](https://wiki.libsdl.org/SDL3/SDL_CreateTexture) | [:heavy_check_mark:](./methods.go#L2540) | [:heavy_check_mark:](./sdl_functions_js.go#L14833) |
| [SDL_CreateTextureFromSurface](https://wiki.libsdl.org/SDL3/SDL_CreateTextureFromSurface) | [:heavy_check_mark:](./methods.go#L2551) | [:heavy_check_mark:](./sdl_functions_js.go#L14856) |
| [SDL_CreateTextureWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateTextureWithProperties) | [:x:](./methods.go#L2562) | [:x:](./sdl_functions_js.go#L14876) |
| [SDL_GetTextureProperties](https://wiki.libsdl.org/SDL3/SDL_GetTextureProperties) | [:x:](./methods.go#L916) | [:x:](./sdl_functions_js.go#L14897) |
| [SDL_GetRendererFromTexture](https://wiki.libsdl.org/SDL3/SDL_GetRendererFromTexture) | [:x:](./methods.go#L923) | [:x:](./sdl_functions_js.go#L14913) |
| [SDL_GetTextureSize](https://wiki.libsdl.org/SDL3/SDL_GetTextureSize) | [:x:](./methods.go#L930) | [:x:](./sdl_functions_js.go#L14932) |
| [SDL_SetTextureColorMod](https://wiki.libsdl.org/SDL3/SDL_SetTextureColorMod) | [:x:](./methods.go#L937) | [:x:](./sdl_functions_js.go#L14958) |
| [SDL_SetTextureColorModFloat](https://wiki.libsdl.org/SDL3/SDL_SetTextureColorModFloat) | [:heavy_check_mark:](./methods.go#L944) | [:heavy_check_mark:](./sdl_functions_js.go#L14980) |
| [SDL_GetTextureColorMod](https://wiki.libsdl.org/SDL3/SDL_GetTextureColorMod) | [:x:](./methods.go#L954) | [:x:](./sdl_functions_js.go#L14996) |
| [SDL_GetTextureColorModFloat](https://wiki.libsdl.org/SDL3/SDL_GetTextureColorModFloat) | [:x:](./methods.go#L961) | [:x:](./sdl_functions_js.go#L15027) |
| [SDL_SetTextureAlphaMod](https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaMod) | [:x:](./methods.go#L968) | [:x:](./sdl_functions_js.go#L15058) |
| [SDL_SetTextureAlphaModFloat](https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaModFloat) | [:x:](./methods.go#L975) | [:x:](./sdl_functions_js.go#L15076) |
| [SDL_GetTextureAlphaMod](https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaMod) | [:x:](./methods.go#L982) | [:x:](./sdl_functions_js.go#L15094) |
| [SDL_GetTextureAlphaModFloat](https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaModFloat) | [:x:](./methods.go#L989) | [:x:](./sdl_functions_js.go#L15115) |
| [SDL_SetTextureBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetTextureBlendMode) | [:x:](./methods.go#L996) | [:x:](./sdl_functions_js.go#L15136) |
| [SDL_GetTextureBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetTextureBlendMode) | [:x:](./methods.go#L1003) | [:x:](./sdl_functions_js.go#L15154) |
| [SDL_SetTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_SetTextureScaleMode) | [:x:](./methods.go#L1010) | [:x:](./sdl_functions_js.go#L15175) |
| [SDL_GetTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_GetTextureScaleMode) | [:x:](./methods.go#L1017) | [:x:](./sdl_functions_js.go#L15193) |
| [SDL_UpdateTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateTexture) | [:heavy_check_mark:](./methods.go#L1024) | [:heavy_check_mark:](./sdl_functions_js.go#L15214) |
| [SDL_UpdateYUVTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateYUVTexture) | [:x:](./methods.go#L1036) | [:x:](./sdl_functions_js.go#L15238) |
| [SDL_UpdateNVTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateNVTexture) | [:x:](./methods.go#L1043) | [:x:](./sdl_functions_js.go#L15280) |
| [SDL_LockTexture](https://wiki.libsdl.org/SDL3/SDL_LockTexture) | [:x:](./methods.go#L1050) | [:x:](./sdl_functions_js.go#L15315) |
| [SDL_LockTextureToSurface](https://wiki.libsdl.org/SDL3/SDL_LockTextureToSurface) | [:heavy_check_mark:](./methods.go#L1057) | [:heavy_check_mark:](./sdl_functions_js.go#L15346) |
| [SDL_UnlockTexture](https://wiki.libsdl.org/SDL3/SDL_UnlockTexture) | [:heavy_check_mark:](./methods.go#L1067) | [:heavy_check_mark:](./sdl_functions_js.go#L15368) |
| [SDL_SetRenderTarget](https://wiki.libsdl.org/SDL3/SDL_SetRenderTarget) | [:x:](./methods.go#L2569) | [:x:](./sdl_functions_js.go#L15379) |
| [SDL_GetRenderTarget](https://wiki.libsdl.org/SDL3/SDL_GetRenderTarget) | [:x:](./methods.go#L2576) | [:x:](./sdl_functions_js.go#L15400) |
| [SDL_SetRenderLogicalPresentation](https://wiki.libsdl.org/SDL3/SDL_SetRenderLogicalPresentation) | [:x:](./methods.go#L2583) | [:x:](./sdl_functions_js.go#L15419) |
| [SDL_GetRenderLogicalPresentation](https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentation) | [:x:](./methods.go#L2590) | [:x:](./sdl_functions_js.go#L15441) |
| [SDL_GetRenderLogicalPresentationRect](https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentationRect) | [:x:](./methods.go#L2597) | [:x:](./sdl_functions_js.go#L15472) |
| [SDL_RenderCoordinatesFromWindow](https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesFromWindow) | [:x:](./methods.go#L2604) | [:x:](./sdl_functions_js.go#L15493) |
| [SDL_RenderCoordinatesToWindow](https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesToWindow) | [:x:](./methods.go#L2611) | [:x:](./sdl_functions_js.go#L15523) |
| [SDL_ConvertEventToRenderCoordinates](https://wiki.libsdl.org/SDL3/SDL_ConvertEventToRenderCoordinates) | [:x:](./methods.go#L2618) | [:x:](./sdl_functions_js.go#L15553) |
| [SDL_SetRenderViewport](https://wiki.libsdl.org/SDL3/SDL_SetRenderViewport) | [:heavy_check_mark:](./methods.go#L2625) | [:heavy_check_mark:](./sdl_functions_js.go#L15574) |
| [SDL_GetRenderViewport](https://wiki.libsdl.org/SDL3/SDL_GetRenderViewport) | [:heavy_check_mark:](./methods.go#L2635) | [:x:](./sdl_functions_js.go#L15591) |
| [SDL_RenderViewportSet](https://wiki.libsdl.org/SDL3/SDL_RenderViewportSet) | [:heavy_check_mark:](./methods.go#L2647) | [:x:](./sdl_functions_js.go#L15612) |
| [SDL_GetRenderSafeArea](https://wiki.libsdl.org/SDL3/SDL_GetRenderSafeArea) | [:heavy_check_mark:](./methods.go#L2653) | [:x:](./sdl_functions_js.go#L15628) |
| [SDL_SetRenderClipRect](https://wiki.libsdl.org/SDL3/SDL_SetRenderClipRect) | [:heavy_check_mark:](./methods.go#L2665) | [:heavy_check_mark:](./sdl_functions_js.go#L15649) |
| [SDL_GetRenderClipRect](https://wiki.libsdl.org/SDL3/SDL_GetRenderClipRect) | [:heavy_check_mark:](./methods.go#L2675) | [:x:](./sdl_functions_js.go#L15666) |
| [SDL_RenderClipEnabled](https://wiki.libsdl.org/SDL3/SDL_RenderClipEnabled) | [:heavy_check_mark:](./methods.go#L2687) | [:x:](./sdl_functions_js.go#L15687) |
| [SDL_SetRenderScale](https://wiki.libsdl.org/SDL3/SDL_SetRenderScale) | [:heavy_check_mark:](./methods.go#L2697) | [:heavy_check_mark:](./sdl_functions_js.go#L15703) |
| [SDL_GetRenderScale](https://wiki.libsdl.org/SDL3/SDL_GetRenderScale) | [:heavy_check_mark:](./methods.go#L2707) | [:x:](./sdl_functions_js.go#L15720) |
| [SDL_SetRenderDrawColor](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColor) | [:heavy_check_mark:](./methods.go#L2719) | [:heavy_check_mark:](./sdl_functions_js.go#L15746) |
| [SDL_SetRenderDrawColorFloat](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColorFloat) | [:heavy_check_mark:](./methods.go#L2729) | [:heavy_check_mark:](./sdl_functions_js.go#L15768) |
| [SDL_GetRenderDrawColor](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColor) | [:heavy_check_mark:](./methods.go#L2739) | [:x:](./sdl_functions_js.go#L15785) |
| [SDL_GetRenderDrawColorFloat](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColorFloat) | [:heavy_check_mark:](./methods.go#L2751) | [:x:](./sdl_functions_js.go#L15821) |
| [SDL_SetRenderColorScale](https://wiki.libsdl.org/SDL3/SDL_SetRenderColorScale) | [:heavy_check_mark:](./methods.go#L2763) | [:x:](./sdl_functions_js.go#L15857) |
| [SDL_GetRenderColorScale](https://wiki.libsdl.org/SDL3/SDL_GetRenderColorScale) | [:heavy_check_mark:](./methods.go#L2773) | [:x:](./sdl_functions_js.go#L15875) |
| [SDL_SetRenderDrawBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawBlendMode) | [:x:](./methods.go#L2785) | [:x:](./sdl_functions_js.go#L15896) |
| [SDL_GetRenderDrawBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawBlendMode) | [:x:](./methods.go#L2792) | [:x:](./sdl_functions_js.go#L15914) |
| [SDL_RenderClear](https://wiki.libsdl.org/SDL3/SDL_RenderClear) | [:heavy_check_mark:](./methods.go#L2799) | [:heavy_check_mark:](./sdl_functions_js.go#L15935) |
| [SDL_RenderPoint](https://wiki.libsdl.org/SDL3/SDL_RenderPoint) | [:x:](./methods.go#L2808) | [:x:](./sdl_functions_js.go#L15948) |
| [SDL_RenderPoints](https://wiki.libsdl.org/SDL3/SDL_RenderPoints) | [:heavy_check_mark:](./methods.go#L2815) | [:heavy_check_mark:](./sdl_functions_js.go#L15968) |
| [SDL_RenderLine](https://wiki.libsdl.org/SDL3/SDL_RenderLine) | [:heavy_check_mark:](./methods.go#L2825) | [:heavy_check_mark:](./sdl_functions_js.go#L15986) |
| [SDL_RenderLines](https://wiki.libsdl.org/SDL3/SDL_RenderLines) | [:heavy_check_mark:](./methods.go#L2835) | [:heavy_check_mark:](./sdl_functions_js.go#L16003) |
| [SDL_RenderRect](https://wiki.libsdl.org/SDL3/SDL_RenderRect) | [:heavy_check_mark:](./methods.go#L2845) | [:heavy_check_mark:](./sdl_functions_js.go#L16021) |
| [SDL_RenderRects](https://wiki.libsdl.org/SDL3/SDL_RenderRects) | [:heavy_check_mark:](./methods.go#L2855) | [:heavy_check_mark:](./sdl_functions_js.go#L16038) |
| [SDL_RenderFillRect](https://wiki.libsdl.org/SDL3/SDL_RenderFillRect) | [:heavy_check_mark:](./methods.go#L2865) | [:heavy_check_mark:](./sdl_functions_js.go#L16056) |
| [SDL_RenderFillRects](https://wiki.libsdl.org/SDL3/SDL_RenderFillRects) | [:heavy_check_mark:](./methods.go#L2875) | [:heavy_check_mark:](./sdl_functions_js.go#L16073) |
| [SDL_RenderTexture](https://wiki.libsdl.org/SDL3/SDL_RenderTexture) | [:heavy_check_mark:](./methods.go#L2885) | [:heavy_check_mark:](./sdl_functions_js.go#L16091) |
| [SDL_RenderTextureRotated](https://wiki.libsdl.org/SDL3/SDL_RenderTextureRotated) | [:heavy_check_mark:](./methods.go#L2895) | [:heavy_check_mark:](./sdl_functions_js.go#L16115) |
| [SDL_RenderTextureAffine](https://wiki.libsdl.org/SDL3/SDL_RenderTextureAffine) | [:x:](./methods.go#L2905) | [:x:](./sdl_functions_js.go#L16145) |
| [SDL_RenderTextureTiled](https://wiki.libsdl.org/SDL3/SDL_RenderTextureTiled) | [:x:](./methods.go#L2912) | [:x:](./sdl_functions_js.go#L16186) |
| [SDL_RenderTexture9Grid](https://wiki.libsdl.org/SDL3/SDL_RenderTexture9Grid) | [:x:](./methods.go#L2919) | [:x:](./sdl_functions_js.go#L16219) |
| [SDL_RenderTexture9GridTiled](https://wiki.libsdl.org/SDL3/SDL_RenderTexture9GridTiled) | [:question:]() | [:question:]() |
| [SDL_RenderGeometry](https://wiki.libsdl.org/SDL3/SDL_RenderGeometry) | [:heavy_check_mark:](./methods.go#L2926) | [:heavy_check_mark:](./sdl_functions_js.go#L16260) |
| [SDL_RenderGeometryRaw](https://wiki.libsdl.org/SDL3/SDL_RenderGeometryRaw) | [:x:](./methods.go#L2936) | [:x:](./sdl_functions_js.go#L16287) |
| [SDL_RenderReadPixels](https://wiki.libsdl.org/SDL3/SDL_RenderReadPixels) | [:heavy_check_mark:](./methods.go#L2943) | [:heavy_check_mark:](./sdl_functions_js.go#L16337) |
| [SDL_RenderPresent](https://wiki.libsdl.org/SDL3/SDL_RenderPresent) | [:heavy_check_mark:](./methods.go#L2954) | [:heavy_check_mark:](./sdl_functions_js.go#L16356) |
| [SDL_DestroyTexture](https://wiki.libsdl.org/SDL3/SDL_DestroyTexture) | [:heavy_check_mark:](./methods.go#L1073) | [:heavy_check_mark:](./sdl_functions_js.go#L16369) |
| [SDL_DestroyRenderer](https://wiki.libsdl.org/SDL3/SDL_DestroyRenderer) | [:heavy_check_mark:](./methods.go#L2964) | [:heavy_check_mark:](./sdl_functions_js.go#L16381) |
| [SDL_FlushRenderer](https://wiki.libsdl.org/SDL3/SDL_FlushRenderer) | [:heavy_check_mark:](./methods.go#L2970) | [:x:](./sdl_functions_js.go#L16393) |
| [SDL_GetRenderMetalLayer](https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalLayer) | [:x:](./methods.go#L2980) | [:x:](./sdl_functions_js.go#L16409) |
| [SDL_GetRenderMetalCommandEncoder](https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalCommandEncoder) | [:x:](./methods.go#L2987) | [:x:](./sdl_functions_js.go#L16425) |
| [SDL_AddVulkanRenderSemaphores](https://wiki.libsdl.org/SDL3/SDL_AddVulkanRenderSemaphores) | [:x:](./methods.go#L2994) | [:x:](./sdl_functions_js.go#L16441) |
| [SDL_SetRenderVSync](https://wiki.libsdl.org/SDL3/SDL_SetRenderVSync) | [:x:](./methods.go#L3001) | [:x:](./sdl_functions_js.go#L16463) |
| [SDL_GetRenderVSync](https://wiki.libsdl.org/SDL3/SDL_GetRenderVSync) | [:x:](./methods.go#L3008) | [:x:](./sdl_functions_js.go#L16481) |
| [SDL_RenderDebugText](https://wiki.libsdl.org/SDL3/SDL_RenderDebugText) | [:heavy_check_mark:](./methods.go#L3025) | [:heavy_check_mark:](./sdl_functions_js.go#L16502) |
| [SDL_RenderDebugTextFormat](https://wiki.libsdl.org/SDL3/SDL_RenderDebugTextFormat) | [:question:]() | [:question:](./sdl_functions_js.go#L16522) |
| [SDL_SetDefaultTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_SetDefaultTextureScaleMode) | [:question:]() | [:question:]() |
| [SDL_GetDefaultTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_GetDefaultTextureScaleMode) | [:question:]() | [:question:]() |
### SharedObject

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_LoadObject](https://wiki.libsdl.org/SDL3/SDL_LoadObject) | [:question:]() | [:question:](./sdl_functions_js.go#L14034) |
| [SDL_LoadFunction](https://wiki.libsdl.org/SDL3/SDL_LoadFunction) | [:x:](./methods.go#L4407) | [:x:](./sdl_functions_js.go#L14050) |
| [SDL_UnloadObject](https://wiki.libsdl.org/SDL3/SDL_UnloadObject) | [:x:](./methods.go#L4414) | [:x:](./sdl_functions_js.go#L14068) |
### Thread

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateThread](https://wiki.libsdl.org/SDL3/SDL_CreateThread) | [:question:]() | [:question:]() |
| [SDL_CreateThreadWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateThreadWithProperties) | [:question:]() | [:question:]() |
| [SDL_GetThreadName](https://wiki.libsdl.org/SDL3/SDL_GetThreadName) | [:x:](./methods.go#L482) | [:x:](./sdl_functions_js.go#L871) |
| [SDL_GetCurrentThreadID](https://wiki.libsdl.org/SDL3/SDL_GetCurrentThreadID) | [:question:]() | [:question:](./sdl_functions_js.go#L887) |
| [SDL_GetThreadID](https://wiki.libsdl.org/SDL3/SDL_GetThreadID) | [:x:](./methods.go#L489) | [:x:](./sdl_functions_js.go#L898) |
| [SDL_SetCurrentThreadPriority](https://wiki.libsdl.org/SDL3/SDL_SetCurrentThreadPriority) | [:x:](./methods.go#L23) | [:x:](./sdl_functions_js.go#L914) |
| [SDL_WaitThread](https://wiki.libsdl.org/SDL3/SDL_WaitThread) | [:x:](./methods.go#L496) | [:x:](./sdl_functions_js.go#L927) |
| [SDL_GetThreadState](https://wiki.libsdl.org/SDL3/SDL_GetThreadState) | [:x:](./methods.go#L503) | [:x:](./sdl_functions_js.go#L946) |
| [SDL_DetachThread](https://wiki.libsdl.org/SDL3/SDL_DetachThread) | [:x:](./methods.go#L510) | [:x:](./sdl_functions_js.go#L962) |
| [SDL_GetTLS](https://wiki.libsdl.org/SDL3/SDL_GetTLS) | [:x:](./methods.go#L597) | [:x:](./sdl_functions_js.go#L976) |
| [SDL_SetTLS](https://wiki.libsdl.org/SDL3/SDL_SetTLS) | [:x:](./methods.go#L604) | [:x:](./sdl_functions_js.go#L992) |
| [SDL_CleanupTLS](https://wiki.libsdl.org/SDL3/SDL_CleanupTLS) | [:question:]() | [:question:](./sdl_functions_js.go#L1012) |
### Mutex

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateMutex](https://wiki.libsdl.org/SDL3/SDL_CreateMutex) | [:question:]() | [:question:](./sdl_functions_js.go#L1021) |
| [SDL_LockMutex](https://wiki.libsdl.org/SDL3/SDL_LockMutex) | [:heavy_check_mark:](./methods.go#L5237) | [:x:](./sdl_functions_js.go#L1035) |
| [SDL_TryLockMutex](https://wiki.libsdl.org/SDL3/SDL_TryLockMutex) | [:heavy_check_mark:](./methods.go#L5243) | [:x:](./sdl_functions_js.go#L1049) |
| [SDL_UnlockMutex](https://wiki.libsdl.org/SDL3/SDL_UnlockMutex) | [:heavy_check_mark:](./methods.go#L5249) | [:x:](./sdl_functions_js.go#L1065) |
| [SDL_DestroyMutex](https://wiki.libsdl.org/SDL3/SDL_DestroyMutex) | [:heavy_check_mark:](./methods.go#L5255) | [:x:](./sdl_functions_js.go#L1079) |
| [SDL_CreateRWLock](https://wiki.libsdl.org/SDL3/SDL_CreateRWLock) | [:question:]() | [:question:](./sdl_functions_js.go#L1093) |
| [SDL_LockRWLockForReading](https://wiki.libsdl.org/SDL3/SDL_LockRWLockForReading) | [:x:](./methods.go#L1084) | [:x:](./sdl_functions_js.go#L1107) |
| [SDL_LockRWLockForWriting](https://wiki.libsdl.org/SDL3/SDL_LockRWLockForWriting) | [:x:](./methods.go#L1091) | [:x:](./sdl_functions_js.go#L1121) |
| [SDL_TryLockRWLockForReading](https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForReading) | [:x:](./methods.go#L1098) | [:x:](./sdl_functions_js.go#L1135) |
| [SDL_TryLockRWLockForWriting](https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForWriting) | [:x:](./methods.go#L1105) | [:x:](./sdl_functions_js.go#L1151) |
| [SDL_UnlockRWLock](https://wiki.libsdl.org/SDL3/SDL_UnlockRWLock) | [:x:](./methods.go#L1112) | [:x:](./sdl_functions_js.go#L1167) |
| [SDL_DestroyRWLock](https://wiki.libsdl.org/SDL3/SDL_DestroyRWLock) | [:x:](./methods.go#L1119) | [:x:](./sdl_functions_js.go#L1181) |
| [SDL_CreateSemaphore](https://wiki.libsdl.org/SDL3/SDL_CreateSemaphore) | [:question:]() | [:question:](./sdl_functions_js.go#L1195) |
| [SDL_DestroySemaphore](https://wiki.libsdl.org/SDL3/SDL_DestroySemaphore) | [:x:](./methods.go#L5458) | [:x:](./sdl_functions_js.go#L1211) |
| [SDL_WaitSemaphore](https://wiki.libsdl.org/SDL3/SDL_WaitSemaphore) | [:x:](./methods.go#L5465) | [:x:](./sdl_functions_js.go#L1225) |
| [SDL_TryWaitSemaphore](https://wiki.libsdl.org/SDL3/SDL_TryWaitSemaphore) | [:x:](./methods.go#L5472) | [:x:](./sdl_functions_js.go#L1239) |
| [SDL_WaitSemaphoreTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitSemaphoreTimeout) | [:x:](./methods.go#L5479) | [:x:](./sdl_functions_js.go#L1255) |
| [SDL_SignalSemaphore](https://wiki.libsdl.org/SDL3/SDL_SignalSemaphore) | [:x:](./methods.go#L5486) | [:x:](./sdl_functions_js.go#L1273) |
| [SDL_GetSemaphoreValue](https://wiki.libsdl.org/SDL3/SDL_GetSemaphoreValue) | [:x:](./methods.go#L5493) | [:x:](./sdl_functions_js.go#L1287) |
| [SDL_CreateCondition](https://wiki.libsdl.org/SDL3/SDL_CreateCondition) | [:question:]() | [:question:](./sdl_functions_js.go#L1303) |
| [SDL_DestroyCondition](https://wiki.libsdl.org/SDL3/SDL_DestroyCondition) | [:x:](./methods.go#L3586) | [:x:](./sdl_functions_js.go#L1317) |
| [SDL_SignalCondition](https://wiki.libsdl.org/SDL3/SDL_SignalCondition) | [:x:](./methods.go#L3593) | [:x:](./sdl_functions_js.go#L1331) |
| [SDL_BroadcastCondition](https://wiki.libsdl.org/SDL3/SDL_BroadcastCondition) | [:x:](./methods.go#L3600) | [:x:](./sdl_functions_js.go#L1345) |
| [SDL_WaitCondition](https://wiki.libsdl.org/SDL3/SDL_WaitCondition) | [:x:](./methods.go#L3607) | [:x:](./sdl_functions_js.go#L1359) |
| [SDL_WaitConditionTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitConditionTimeout) | [:x:](./methods.go#L3614) | [:x:](./sdl_functions_js.go#L1378) |
| [SDL_ShouldInit](https://wiki.libsdl.org/SDL3/SDL_ShouldInit) | [:x:](./methods.go#L3081) | [:x:](./sdl_functions_js.go#L1401) |
| [SDL_ShouldQuit](https://wiki.libsdl.org/SDL3/SDL_ShouldQuit) | [:x:](./methods.go#L3088) | [:x:](./sdl_functions_js.go#L1417) |
| [SDL_SetInitialized](https://wiki.libsdl.org/SDL3/SDL_SetInitialized) | [:x:](./methods.go#L3095) | [:x:](./sdl_functions_js.go#L1433) |
### Atomic

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_TryLockSpinlock](https://wiki.libsdl.org/SDL3/SDL_TryLockSpinlock) | [:x:](./methods.go#L459) | [:x:](./sdl_functions_js.go#L246) |
| [SDL_LockSpinlock](https://wiki.libsdl.org/SDL3/SDL_LockSpinlock) | [:x:](./methods.go#L466) | [:x:](./sdl_functions_js.go#L262) |
| [SDL_UnlockSpinlock](https://wiki.libsdl.org/SDL3/SDL_UnlockSpinlock) | [:x:](./methods.go#L473) | [:x:](./sdl_functions_js.go#L276) |
| [SDL_MemoryBarrierReleaseFunction](https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierReleaseFunction) | [:question:]() | [:question:](./sdl_functions_js.go#L290) |
| [SDL_MemoryBarrierAcquireFunction](https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierAcquireFunction) | [:question:]() | [:question:](./sdl_functions_js.go#L299) |
| [SDL_CompareAndSwapAtomicInt](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicInt) | [:x:](./methods.go#L835) | [:x:](./sdl_functions_js.go#L308) |
| [SDL_SetAtomicInt](https://wiki.libsdl.org/SDL3/SDL_SetAtomicInt) | [:x:](./methods.go#L842) | [:x:](./sdl_functions_js.go#L328) |
| [SDL_GetAtomicInt](https://wiki.libsdl.org/SDL3/SDL_GetAtomicInt) | [:x:](./methods.go#L849) | [:x:](./sdl_functions_js.go#L346) |
| [SDL_AddAtomicInt](https://wiki.libsdl.org/SDL3/SDL_AddAtomicInt) | [:x:](./methods.go#L856) | [:x:](./sdl_functions_js.go#L362) |
| [SDL_CompareAndSwapAtomicU32](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicU32) | [:x:](./methods.go#L5279) | [:x:](./sdl_functions_js.go#L380) |
| [SDL_SetAtomicU32](https://wiki.libsdl.org/SDL3/SDL_SetAtomicU32) | [:x:](./methods.go#L5286) | [:x:](./sdl_functions_js.go#L400) |
| [SDL_GetAtomicU32](https://wiki.libsdl.org/SDL3/SDL_GetAtomicU32) | [:x:](./methods.go#L5293) | [:x:](./sdl_functions_js.go#L418) |
| [SDL_CompareAndSwapAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicPointer) | [:question:]() | [:question:](./sdl_functions_js.go#L434) |
| [SDL_SetAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_SetAtomicPointer) | [:question:]() | [:question:](./sdl_functions_js.go#L454) |
| [SDL_GetAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_GetAtomicPointer) | [:question:]() | [:question:](./sdl_functions_js.go#L472) |
### Filesystem

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetBasePath](https://wiki.libsdl.org/SDL3/SDL_GetBasePath) | [:question:]() | [:question:](./sdl_functions_js.go#L11221) |
| [SDL_GetPrefPath](https://wiki.libsdl.org/SDL3/SDL_GetPrefPath) | [:question:]() | [:question:](./sdl_functions_js.go#L11232) |
| [SDL_GetUserFolder](https://wiki.libsdl.org/SDL3/SDL_GetUserFolder) | [:question:]() | [:question:](./sdl_functions_js.go#L11247) |
| [SDL_CreateDirectory](https://wiki.libsdl.org/SDL3/SDL_CreateDirectory) | [:question:]() | [:question:](./sdl_functions_js.go#L11260) |
| [SDL_EnumerateDirectory](https://wiki.libsdl.org/SDL3/SDL_EnumerateDirectory) | [:question:]() | [:question:](./sdl_functions_js.go#L11273) |
| [SDL_RemovePath](https://wiki.libsdl.org/SDL3/SDL_RemovePath) | [:question:]() | [:question:](./sdl_functions_js.go#L11290) |
| [SDL_RenamePath](https://wiki.libsdl.org/SDL3/SDL_RenamePath) | [:question:]() | [:question:](./sdl_functions_js.go#L11303) |
| [SDL_CopyFile](https://wiki.libsdl.org/SDL3/SDL_CopyFile) | [:question:]() | [:question:](./sdl_functions_js.go#L11318) |
| [SDL_GetPathInfo](https://wiki.libsdl.org/SDL3/SDL_GetPathInfo) | [:question:]() | [:question:](./sdl_functions_js.go#L11333) |
| [SDL_GlobDirectory](https://wiki.libsdl.org/SDL3/SDL_GlobDirectory) | [:question:]() | [:question:](./sdl_functions_js.go#L11351) |
| [SDL_GetCurrentDirectory](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDirectory) | [:question:]() | [:question:](./sdl_functions_js.go#L11373) |
### IOStream

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_IOFromFile](https://wiki.libsdl.org/SDL3/SDL_IOFromFile) | [:question:]() | [:question:](./sdl_functions_js.go#L1449) |
| [SDL_IOFromMem](https://wiki.libsdl.org/SDL3/SDL_IOFromMem) | [:question:]() | [:question:](./sdl_functions_js.go#L1467) |
| [SDL_IOFromConstMem](https://wiki.libsdl.org/SDL3/SDL_IOFromConstMem) | [:heavy_check_mark:](./functions.go#L248) | [:heavy_check_mark:](./sdl_functions_js.go#L1485) |
| [SDL_IOFromDynamicMem](https://wiki.libsdl.org/SDL3/SDL_IOFromDynamicMem) | [:question:]() | [:question:](./sdl_functions_js.go#L1505) |
| [SDL_OpenIO](https://wiki.libsdl.org/SDL3/SDL_OpenIO) | [:x:](./methods.go#L32) | [:x:](./sdl_functions_js.go#L1519) |
| [SDL_CloseIO](https://wiki.libsdl.org/SDL3/SDL_CloseIO) | [:heavy_check_mark:](./methods.go#L4432) | [:heavy_check_mark:](./sdl_functions_js.go#L1540) |
| [SDL_GetIOProperties](https://wiki.libsdl.org/SDL3/SDL_GetIOProperties) | [:heavy_check_mark:](./methods.go#L4442) | [:x:](./sdl_functions_js.go#L1553) |
| [SDL_GetIOStatus](https://wiki.libsdl.org/SDL3/SDL_GetIOStatus) | [:heavy_check_mark:](./methods.go#L4453) | [:x:](./sdl_functions_js.go#L1569) |
| [SDL_GetIOSize](https://wiki.libsdl.org/SDL3/SDL_GetIOSize) | [:heavy_check_mark:](./methods.go#L4459) | [:x:](./sdl_functions_js.go#L1585) |
| [SDL_SeekIO](https://wiki.libsdl.org/SDL3/SDL_SeekIO) | [:heavy_check_mark:](./methods.go#L4470) | [:x:](./sdl_functions_js.go#L1601) |
| [SDL_TellIO](https://wiki.libsdl.org/SDL3/SDL_TellIO) | [:heavy_check_mark:](./methods.go#L4481) | [:x:](./sdl_functions_js.go#L1621) |
| [SDL_ReadIO](https://wiki.libsdl.org/SDL3/SDL_ReadIO) | [:heavy_check_mark:](./methods.go#L4487) | [:x:](./sdl_functions_js.go#L1637) |
| [SDL_WriteIO](https://wiki.libsdl.org/SDL3/SDL_WriteIO) | [:heavy_check_mark:](./methods.go#L4498) | [:x:](./sdl_functions_js.go#L1657) |
| [SDL_IOprintf](https://wiki.libsdl.org/SDL3/SDL_IOprintf) | [:heavy_check_mark:](./methods.go#L4509) | [:x:](./sdl_functions_js.go#L1677) |
| [SDL_IOvprintf](https://wiki.libsdl.org/SDL3/SDL_IOvprintf) | [:question:]() | [:question:](./sdl_functions_js.go#L1695) |
| [SDL_FlushIO](https://wiki.libsdl.org/SDL3/SDL_FlushIO) | [:heavy_check_mark:](./methods.go#L4520) | [:x:](./sdl_functions_js.go#L1715) |
| [SDL_LoadFile_IO](https://wiki.libsdl.org/SDL3/SDL_LoadFile_IO) | [:x:](./methods.go#L4530) | [:x:](./sdl_functions_js.go#L1731) |
| [SDL_LoadFile](https://wiki.libsdl.org/SDL3/SDL_LoadFile) | [:question:]() | [:question:](./sdl_functions_js.go#L1754) |
| [SDL_SaveFile_IO](https://wiki.libsdl.org/SDL3/SDL_SaveFile_IO) | [:x:](./methods.go#L4537) | [:x:](./sdl_functions_js.go#L1772) |
| [SDL_SaveFile](https://wiki.libsdl.org/SDL3/SDL_SaveFile) | [:question:]() | [:question:](./sdl_functions_js.go#L1794) |
| [SDL_ReadU8](https://wiki.libsdl.org/SDL3/SDL_ReadU8) | [:heavy_check_mark:](./methods.go#L4544) | [:x:](./sdl_functions_js.go#L1811) |
| [SDL_ReadS8](https://wiki.libsdl.org/SDL3/SDL_ReadS8) | [:heavy_check_mark:](./methods.go#L4556) | [:x:](./sdl_functions_js.go#L1832) |
| [SDL_ReadU16LE](https://wiki.libsdl.org/SDL3/SDL_ReadU16LE) | [:heavy_check_mark:](./methods.go#L4568) | [:x:](./sdl_functions_js.go#L1853) |
| [SDL_ReadS16LE](https://wiki.libsdl.org/SDL3/SDL_ReadS16LE) | [:heavy_check_mark:](./methods.go#L4580) | [:x:](./sdl_functions_js.go#L1874) |
| [SDL_ReadU16BE](https://wiki.libsdl.org/SDL3/SDL_ReadU16BE) | [:heavy_check_mark:](./methods.go#L4592) | [:x:](./sdl_functions_js.go#L1895) |
| [SDL_ReadS16BE](https://wiki.libsdl.org/SDL3/SDL_ReadS16BE) | [:heavy_check_mark:](./methods.go#L4604) | [:x:](./sdl_functions_js.go#L1916) |
| [SDL_ReadU32LE](https://wiki.libsdl.org/SDL3/SDL_ReadU32LE) | [:heavy_check_mark:](./methods.go#L4616) | [:x:](./sdl_functions_js.go#L1937) |
| [SDL_ReadS32LE](https://wiki.libsdl.org/SDL3/SDL_ReadS32LE) | [:heavy_check_mark:](./methods.go#L4628) | [:x:](./sdl_functions_js.go#L1958) |
| [SDL_ReadU32BE](https://wiki.libsdl.org/SDL3/SDL_ReadU32BE) | [:heavy_check_mark:](./methods.go#L4640) | [:x:](./sdl_functions_js.go#L1979) |
| [SDL_ReadS32BE](https://wiki.libsdl.org/SDL3/SDL_ReadS32BE) | [:heavy_check_mark:](./methods.go#L4652) | [:x:](./sdl_functions_js.go#L2000) |
| [SDL_ReadU64LE](https://wiki.libsdl.org/SDL3/SDL_ReadU64LE) | [:heavy_check_mark:](./methods.go#L4664) | [:x:](./sdl_functions_js.go#L2021) |
| [SDL_ReadS64LE](https://wiki.libsdl.org/SDL3/SDL_ReadS64LE) | [:heavy_check_mark:](./methods.go#L4676) | [:x:](./sdl_functions_js.go#L2042) |
| [SDL_ReadU64BE](https://wiki.libsdl.org/SDL3/SDL_ReadU64BE) | [:heavy_check_mark:](./methods.go#L4688) | [:x:](./sdl_functions_js.go#L2063) |
| [SDL_ReadS64BE](https://wiki.libsdl.org/SDL3/SDL_ReadS64BE) | [:heavy_check_mark:](./methods.go#L4700) | [:x:](./sdl_functions_js.go#L2084) |
| [SDL_WriteU8](https://wiki.libsdl.org/SDL3/SDL_WriteU8) | [:heavy_check_mark:](./methods.go#L4712) | [:x:](./sdl_functions_js.go#L2105) |
| [SDL_WriteS8](https://wiki.libsdl.org/SDL3/SDL_WriteS8) | [:heavy_check_mark:](./methods.go#L4722) | [:x:](./sdl_functions_js.go#L2123) |
| [SDL_WriteU16LE](https://wiki.libsdl.org/SDL3/SDL_WriteU16LE) | [:heavy_check_mark:](./methods.go#L4732) | [:x:](./sdl_functions_js.go#L2141) |
| [SDL_WriteS16LE](https://wiki.libsdl.org/SDL3/SDL_WriteS16LE) | [:heavy_check_mark:](./methods.go#L4742) | [:x:](./sdl_functions_js.go#L2159) |
| [SDL_WriteU16BE](https://wiki.libsdl.org/SDL3/SDL_WriteU16BE) | [:heavy_check_mark:](./methods.go#L4752) | [:x:](./sdl_functions_js.go#L2177) |
| [SDL_WriteS16BE](https://wiki.libsdl.org/SDL3/SDL_WriteS16BE) | [:heavy_check_mark:](./methods.go#L4762) | [:x:](./sdl_functions_js.go#L2195) |
| [SDL_WriteU32LE](https://wiki.libsdl.org/SDL3/SDL_WriteU32LE) | [:heavy_check_mark:](./methods.go#L4772) | [:x:](./sdl_functions_js.go#L2213) |
| [SDL_WriteS32LE](https://wiki.libsdl.org/SDL3/SDL_WriteS32LE) | [:heavy_check_mark:](./methods.go#L4782) | [:x:](./sdl_functions_js.go#L2231) |
| [SDL_WriteU32BE](https://wiki.libsdl.org/SDL3/SDL_WriteU32BE) | [:heavy_check_mark:](./methods.go#L4792) | [:x:](./sdl_functions_js.go#L2249) |
| [SDL_WriteS32BE](https://wiki.libsdl.org/SDL3/SDL_WriteS32BE) | [:heavy_check_mark:](./methods.go#L4802) | [:x:](./sdl_functions_js.go#L2267) |
| [SDL_WriteU64LE](https://wiki.libsdl.org/SDL3/SDL_WriteU64LE) | [:heavy_check_mark:](./methods.go#L4812) | [:x:](./sdl_functions_js.go#L2285) |
| [SDL_WriteS64LE](https://wiki.libsdl.org/SDL3/SDL_WriteS64LE) | [:heavy_check_mark:](./methods.go#L4822) | [:x:](./sdl_functions_js.go#L2303) |
| [SDL_WriteU64BE](https://wiki.libsdl.org/SDL3/SDL_WriteU64BE) | [:heavy_check_mark:](./methods.go#L4832) | [:x:](./sdl_functions_js.go#L2321) |
| [SDL_WriteS64BE](https://wiki.libsdl.org/SDL3/SDL_WriteS64BE) | [:heavy_check_mark:](./methods.go#L4842) | [:x:](./sdl_functions_js.go#L2339) |
### AsyncIO

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AsyncIOFromFile](https://wiki.libsdl.org/SDL3/SDL_AsyncIOFromFile) | [:question:]() | [:question:](./sdl_functions_js.go#L23) |
| [SDL_GetAsyncIOSize](https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOSize) | [:x:](./methods.go#L3051) | [:x:](./sdl_functions_js.go#L41) |
| [SDL_ReadAsyncIO](https://wiki.libsdl.org/SDL3/SDL_ReadAsyncIO) | [:x:](./methods.go#L3058) | [:x:](./sdl_functions_js.go#L57) |
| [SDL_WriteAsyncIO](https://wiki.libsdl.org/SDL3/SDL_WriteAsyncIO) | [:x:](./methods.go#L3065) | [:x:](./sdl_functions_js.go#L86) |
| [SDL_CloseAsyncIO](https://wiki.libsdl.org/SDL3/SDL_CloseAsyncIO) | [:x:](./methods.go#L3072) | [:x:](./sdl_functions_js.go#L115) |
| [SDL_CreateAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_CreateAsyncIOQueue) | [:question:]() | [:question:](./sdl_functions_js.go#L140) |
| [SDL_DestroyAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_DestroyAsyncIOQueue) | [:x:](./methods.go#L1274) | [:x:](./sdl_functions_js.go#L154) |
| [SDL_GetAsyncIOResult](https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOResult) | [:x:](./methods.go#L1281) | [:x:](./sdl_functions_js.go#L168) |
| [SDL_WaitAsyncIOResult](https://wiki.libsdl.org/SDL3/SDL_WaitAsyncIOResult) | [:x:](./methods.go#L1288) | [:x:](./sdl_functions_js.go#L189) |
| [SDL_SignalAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_SignalAsyncIOQueue) | [:x:](./methods.go#L1295) | [:x:](./sdl_functions_js.go#L212) |
| [SDL_LoadFileAsync](https://wiki.libsdl.org/SDL3/SDL_LoadFileAsync) | [:question:]() | [:question:](./sdl_functions_js.go#L226) |
### Storage

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_OpenTitleStorage](https://wiki.libsdl.org/SDL3/SDL_OpenTitleStorage) | [:question:]() | [:question:](./sdl_functions_js.go#L16544) |
| [SDL_OpenUserStorage](https://wiki.libsdl.org/SDL3/SDL_OpenUserStorage) | [:question:]() | [:question:](./sdl_functions_js.go#L16562) |
| [SDL_OpenFileStorage](https://wiki.libsdl.org/SDL3/SDL_OpenFileStorage) | [:question:]() | [:question:](./sdl_functions_js.go#L16582) |
| [SDL_OpenStorage](https://wiki.libsdl.org/SDL3/SDL_OpenStorage) | [:x:](./methods.go#L5228) | [:x:](./sdl_functions_js.go#L16598) |
| [SDL_CloseStorage](https://wiki.libsdl.org/SDL3/SDL_CloseStorage) | [:heavy_check_mark:](./methods.go#L96) | [:x:](./sdl_functions_js.go#L16619) |
| [SDL_StorageReady](https://wiki.libsdl.org/SDL3/SDL_StorageReady) | [:heavy_check_mark:](./methods.go#L106) | [:x:](./sdl_functions_js.go#L16635) |
| [SDL_GetStorageFileSize](https://wiki.libsdl.org/SDL3/SDL_GetStorageFileSize) | [:heavy_check_mark:](./methods.go#L112) | [:x:](./sdl_functions_js.go#L16651) |
| [SDL_ReadStorageFile](https://wiki.libsdl.org/SDL3/SDL_ReadStorageFile) | [:heavy_check_mark:](./methods.go#L124) | [:x:](./sdl_functions_js.go#L16674) |
| [SDL_WriteStorageFile](https://wiki.libsdl.org/SDL3/SDL_WriteStorageFile) | [:heavy_check_mark:](./methods.go#L136) | [:x:](./sdl_functions_js.go#L16696) |
| [SDL_CreateStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_CreateStorageDirectory) | [:heavy_check_mark:](./methods.go#L146) | [:x:](./sdl_functions_js.go#L16718) |
| [SDL_EnumerateStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_EnumerateStorageDirectory) | [:x:](./methods.go#L156) | [:x:](./sdl_functions_js.go#L16736) |
| [SDL_RemoveStoragePath](https://wiki.libsdl.org/SDL3/SDL_RemoveStoragePath) | [:heavy_check_mark:](./methods.go#L163) | [:x:](./sdl_functions_js.go#L16758) |
| [SDL_RenameStoragePath](https://wiki.libsdl.org/SDL3/SDL_RenameStoragePath) | [:heavy_check_mark:](./methods.go#L173) | [:x:](./sdl_functions_js.go#L16776) |
| [SDL_CopyStorageFile](https://wiki.libsdl.org/SDL3/SDL_CopyStorageFile) | [:heavy_check_mark:](./methods.go#L183) | [:x:](./sdl_functions_js.go#L16796) |
| [SDL_GetStoragePathInfo](https://wiki.libsdl.org/SDL3/SDL_GetStoragePathInfo) | [:heavy_check_mark:](./methods.go#L193) | [:x:](./sdl_functions_js.go#L16816) |
| [SDL_GetStorageSpaceRemaining](https://wiki.libsdl.org/SDL3/SDL_GetStorageSpaceRemaining) | [:heavy_check_mark:](./methods.go#L205) | [:x:](./sdl_functions_js.go#L16839) |
| [SDL_GlobStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_GlobStorageDirectory) | [:x:](./methods.go#L211) | [:x:](./sdl_functions_js.go#L16855) |
### Pixels

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPixelFormatName](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatName) | [:heavy_check_mark:](./methods.go#L5502) | [:x:](./sdl_functions_js.go#L3386) |
| [SDL_GetMasksForPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetMasksForPixelFormat) | [:x:](./methods.go#L5508) | [:x:](./sdl_functions_js.go#L3399) |
| [SDL_GetPixelFormatForMasks](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatForMasks) | [:heavy_check_mark:](./functions.go#L314) | [:x:](./sdl_functions_js.go#L3437) |
| [SDL_GetPixelFormatDetails](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatDetails) | [:heavy_check_mark:](./methods.go#L5516) | [:heavy_check_mark:](./sdl_functions_js.go#L3458) |
| [SDL_CreatePalette](https://wiki.libsdl.org/SDL3/SDL_CreatePalette) | [:heavy_check_mark:](./functions.go#L320) | [:x:](./sdl_functions_js.go#L3470) |
| [SDL_SetPaletteColors](https://wiki.libsdl.org/SDL3/SDL_SetPaletteColors) | [:heavy_check_mark:](./methods.go#L4854) | [:x:](./sdl_functions_js.go#L3486) |
| [SDL_DestroyPalette](https://wiki.libsdl.org/SDL3/SDL_DestroyPalette) | [:heavy_check_mark:](./methods.go#L4864) | [:x:](./sdl_functions_js.go#L3511) |
| [SDL_MapRGB](https://wiki.libsdl.org/SDL3/SDL_MapRGB) | [:x:](./methods.go#L1304) | [:x:](./sdl_functions_js.go#L3525) |
| [SDL_MapRGBA](https://wiki.libsdl.org/SDL3/SDL_MapRGBA) | [:x:](./methods.go#L1311) | [:x:](./sdl_functions_js.go#L3552) |
| [SDL_GetRGB](https://wiki.libsdl.org/SDL3/SDL_GetRGB) | [:heavy_check_mark:](./functions.go#L343) | [:x:](./sdl_functions_js.go#L3581) |
| [SDL_GetRGBA](https://wiki.libsdl.org/SDL3/SDL_GetRGBA) | [:heavy_check_mark:](./functions.go#L351) | [:x:](./sdl_functions_js.go#L3617) |
### Surface

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateSurface](https://wiki.libsdl.org/SDL3/SDL_CreateSurface) | [:heavy_check_mark:](./functions.go#L372) | [:x:](./sdl_functions_js.go#L3932) |
| [SDL_CreateSurfaceFrom](https://wiki.libsdl.org/SDL3/SDL_CreateSurfaceFrom) | [:question:]() | [:question:](./sdl_functions_js.go#L3952) |
| [SDL_DestroySurface](https://wiki.libsdl.org/SDL3/SDL_DestroySurface) | [:heavy_check_mark:](./methods.go#L1320) | [:heavy_check_mark:](./sdl_functions_js.go#L3976) |
| [SDL_GetSurfaceProperties](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceProperties) | [:x:](./methods.go#L1329) | [:x:](./sdl_functions_js.go#L3988) |
| [SDL_SetSurfaceColorspace](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorspace) | [:x:](./methods.go#L1336) | [:x:](./sdl_functions_js.go#L4004) |
| [SDL_GetSurfaceColorspace](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorspace) | [:x:](./methods.go#L1343) | [:x:](./sdl_functions_js.go#L4022) |
| [SDL_CreateSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_CreateSurfacePalette) | [:x:](./methods.go#L1350) | [:x:](./sdl_functions_js.go#L4038) |
| [SDL_SetSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_SetSurfacePalette) | [:x:](./methods.go#L1357) | [:x:](./sdl_functions_js.go#L4057) |
| [SDL_GetSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_GetSurfacePalette) | [:x:](./methods.go#L1364) | [:x:](./sdl_functions_js.go#L4078) |
| [SDL_AddSurfaceAlternateImage](https://wiki.libsdl.org/SDL3/SDL_AddSurfaceAlternateImage) | [:x:](./methods.go#L1371) | [:x:](./sdl_functions_js.go#L4097) |
| [SDL_SurfaceHasAlternateImages](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasAlternateImages) | [:x:](./methods.go#L1378) | [:x:](./sdl_functions_js.go#L4118) |
| [SDL_GetSurfaceImages](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceImages) | [:x:](./methods.go#L1385) | [:x:](./sdl_functions_js.go#L4134) |
| [SDL_RemoveSurfaceAlternateImages](https://wiki.libsdl.org/SDL3/SDL_RemoveSurfaceAlternateImages) | [:x:](./methods.go#L1392) | [:x:](./sdl_functions_js.go#L4155) |
| [SDL_LockSurface](https://wiki.libsdl.org/SDL3/SDL_LockSurface) | [:x:](./methods.go#L1399) | [:x:](./sdl_functions_js.go#L4169) |
| [SDL_UnlockSurface](https://wiki.libsdl.org/SDL3/SDL_UnlockSurface) | [:x:](./methods.go#L1406) | [:x:](./sdl_functions_js.go#L4185) |
| [SDL_LoadBMP_IO](https://wiki.libsdl.org/SDL3/SDL_LoadBMP_IO) | [:heavy_check_mark:](./functions.go#L383) | [:heavy_check_mark:](./sdl_functions_js.go#L4199) |
| [SDL_LoadBMP](https://wiki.libsdl.org/SDL3/SDL_LoadBMP) | [:heavy_check_mark:](./functions.go#L394) | [:x:](./sdl_functions_js.go#L4219) |
| [SDL_SaveBMP_IO](https://wiki.libsdl.org/SDL3/SDL_SaveBMP_IO) | [:x:](./methods.go#L1413) | [:x:](./sdl_functions_js.go#L4235) |
| [SDL_SaveBMP](https://wiki.libsdl.org/SDL3/SDL_SaveBMP) | [:x:](./methods.go#L1420) | [:x:](./sdl_functions_js.go#L4258) |
| [SDL_SetSurfaceRLE](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceRLE) | [:x:](./methods.go#L1427) | [:x:](./sdl_functions_js.go#L4276) |
| [SDL_SurfaceHasRLE](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasRLE) | [:x:](./methods.go#L1434) | [:x:](./sdl_functions_js.go#L4294) |
| [SDL_SetSurfaceColorKey](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorKey) | [:x:](./methods.go#L1441) | [:x:](./sdl_functions_js.go#L4310) |
| [SDL_SurfaceHasColorKey](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasColorKey) | [:x:](./methods.go#L1448) | [:x:](./sdl_functions_js.go#L4330) |
| [SDL_GetSurfaceColorKey](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorKey) | [:x:](./methods.go#L1455) | [:x:](./sdl_functions_js.go#L4346) |
| [SDL_SetSurfaceColorMod](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorMod) | [:x:](./methods.go#L1462) | [:x:](./sdl_functions_js.go#L4367) |
| [SDL_GetSurfaceColorMod](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorMod) | [:x:](./methods.go#L1469) | [:x:](./sdl_functions_js.go#L4389) |
| [SDL_SetSurfaceAlphaMod](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceAlphaMod) | [:x:](./methods.go#L1476) | [:x:](./sdl_functions_js.go#L4420) |
| [SDL_GetSurfaceAlphaMod](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceAlphaMod) | [:x:](./methods.go#L1483) | [:x:](./sdl_functions_js.go#L4438) |
| [SDL_SetSurfaceBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceBlendMode) | [:x:](./methods.go#L1490) | [:x:](./sdl_functions_js.go#L4459) |
| [SDL_GetSurfaceBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceBlendMode) | [:x:](./methods.go#L1497) | [:x:](./sdl_functions_js.go#L4477) |
| [SDL_SetSurfaceClipRect](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceClipRect) | [:x:](./methods.go#L1504) | [:x:](./sdl_functions_js.go#L4498) |
| [SDL_GetSurfaceClipRect](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceClipRect) | [:x:](./methods.go#L1511) | [:x:](./sdl_functions_js.go#L4519) |
| [SDL_FlipSurface](https://wiki.libsdl.org/SDL3/SDL_FlipSurface) | [:x:](./methods.go#L1518) | [:x:](./sdl_functions_js.go#L4540) |
| [SDL_DuplicateSurface](https://wiki.libsdl.org/SDL3/SDL_DuplicateSurface) | [:x:](./methods.go#L1525) | [:x:](./sdl_functions_js.go#L4558) |
| [SDL_ScaleSurface](https://wiki.libsdl.org/SDL3/SDL_ScaleSurface) | [:x:](./methods.go#L1532) | [:x:](./sdl_functions_js.go#L4577) |
| [SDL_ConvertSurface](https://wiki.libsdl.org/SDL3/SDL_ConvertSurface) | [:heavy_check_mark:](./methods.go#L1539) | [:heavy_check_mark:](./sdl_functions_js.go#L4602) |
| [SDL_ConvertSurfaceAndColorspace](https://wiki.libsdl.org/SDL3/SDL_ConvertSurfaceAndColorspace) | [:x:](./methods.go#L1550) | [:x:](./sdl_functions_js.go#L4619) |
| [SDL_ConvertPixels](https://wiki.libsdl.org/SDL3/SDL_ConvertPixels) | [:question:]() | [:question:](./sdl_functions_js.go#L4649) |
| [SDL_ConvertPixelsAndColorspace](https://wiki.libsdl.org/SDL3/SDL_ConvertPixelsAndColorspace) | [:question:]() | [:question:](./sdl_functions_js.go#L4676) |
| [SDL_PremultiplyAlpha](https://wiki.libsdl.org/SDL3/SDL_PremultiplyAlpha) | [:question:]() | [:question:](./sdl_functions_js.go#L4711) |
| [SDL_PremultiplySurfaceAlpha](https://wiki.libsdl.org/SDL3/SDL_PremultiplySurfaceAlpha) | [:x:](./methods.go#L1557) | [:x:](./sdl_functions_js.go#L4740) |
| [SDL_ClearSurface](https://wiki.libsdl.org/SDL3/SDL_ClearSurface) | [:x:](./methods.go#L1564) | [:x:](./sdl_functions_js.go#L4758) |
| [SDL_FillSurfaceRect](https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRect) | [:heavy_check_mark:](./methods.go#L1571) | [:heavy_check_mark:](./sdl_functions_js.go#L4782) |
| [SDL_FillSurfaceRects](https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRects) | [:x:](./methods.go#L1581) | [:x:](./sdl_functions_js.go#L4801) |
| [SDL_BlitSurface](https://wiki.libsdl.org/SDL3/SDL_BlitSurface) | [:x:](./methods.go#L1588) | [:x:](./sdl_functions_js.go#L4826) |
| [SDL_BlitSurfaceUnchecked](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUnchecked) | [:x:](./methods.go#L1595) | [:x:](./sdl_functions_js.go#L4857) |
| [SDL_BlitSurfaceScaled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceScaled) | [:x:](./methods.go#L1602) | [:x:](./sdl_functions_js.go#L4888) |
| [SDL_BlitSurfaceUncheckedScaled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUncheckedScaled) | [:x:](./methods.go#L1609) | [:x:](./sdl_functions_js.go#L4921) |
| [SDL_StretchSurface](https://wiki.libsdl.org/SDL3/SDL_StretchSurface) | [:question:]() | [:question:]() |
| [SDL_BlitSurfaceTiled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiled) | [:x:](./methods.go#L1616) | [:x:](./sdl_functions_js.go#L4954) |
| [SDL_BlitSurfaceTiledWithScale](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiledWithScale) | [:x:](./methods.go#L1623) | [:x:](./sdl_functions_js.go#L4985) |
| [SDL_BlitSurface9Grid](https://wiki.libsdl.org/SDL3/SDL_BlitSurface9Grid) | [:x:](./methods.go#L1630) | [:x:](./sdl_functions_js.go#L5020) |
| [SDL_MapSurfaceRGB](https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGB) | [:x:](./methods.go#L1637) | [:x:](./sdl_functions_js.go#L5063) |
| [SDL_MapSurfaceRGBA](https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGBA) | [:x:](./methods.go#L1644) | [:x:](./sdl_functions_js.go#L5085) |
| [SDL_ReadSurfacePixel](https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixel) | [:x:](./methods.go#L1651) | [:x:](./sdl_functions_js.go#L5109) |
| [SDL_ReadSurfacePixelFloat](https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixelFloat) | [:x:](./methods.go#L1658) | [:x:](./sdl_functions_js.go#L5140) |
| [SDL_WriteSurfacePixel](https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixel) | [:x:](./methods.go#L1665) | [:x:](./sdl_functions_js.go#L5180) |
| [SDL_WriteSurfacePixelFloat](https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixelFloat) | [:x:](./methods.go#L1672) | [:x:](./sdl_functions_js.go#L5208) |
### BlendMode

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ComposeCustomBlendMode](https://wiki.libsdl.org/SDL3/SDL_ComposeCustomBlendMode) | [:heavy_check_mark:](./functions.go#L411) | [:x:](./sdl_functions_js.go#L3363) |
### Rect

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_RectToFRect](https://wiki.libsdl.org/SDL3/SDL_RectToFRect) | [:question:]() | [:question:]() |
| [SDL_PointInRect](https://wiki.libsdl.org/SDL3/SDL_PointInRect) | [:question:]() | [:question:]() |
| [SDL_RectEmpty](https://wiki.libsdl.org/SDL3/SDL_RectEmpty) | [:question:]() | [:question:]() |
| [SDL_RectsEqual](https://wiki.libsdl.org/SDL3/SDL_RectsEqual) | [:question:]() | [:question:]() |
| [SDL_HasRectIntersection](https://wiki.libsdl.org/SDL3/SDL_HasRectIntersection) | [:x:](./methods.go#L613) | [:x:](./sdl_functions_js.go#L3658) |
| [SDL_GetRectIntersection](https://wiki.libsdl.org/SDL3/SDL_GetRectIntersection) | [:x:](./methods.go#L620) | [:x:](./sdl_functions_js.go#L3679) |
| [SDL_GetRectUnion](https://wiki.libsdl.org/SDL3/SDL_GetRectUnion) | [:x:](./methods.go#L627) | [:x:](./sdl_functions_js.go#L3705) |
| [SDL_GetRectEnclosingPoints](https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPoints) | [:x:](./methods.go#L1144) | [:x:](./sdl_functions_js.go#L3731) |
| [SDL_GetRectAndLineIntersection](https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersection) | [:x:](./methods.go#L634) | [:x:](./sdl_functions_js.go#L3759) |
| [SDL_PointInRectFloat](https://wiki.libsdl.org/SDL3/SDL_PointInRectFloat) | [:question:]() | [:question:]() |
| [SDL_RectEmptyFloat](https://wiki.libsdl.org/SDL3/SDL_RectEmptyFloat) | [:question:]() | [:question:]() |
| [SDL_RectsEqualEpsilon](https://wiki.libsdl.org/SDL3/SDL_RectsEqualEpsilon) | [:question:]() | [:question:]() |
| [SDL_RectsEqualFloat](https://wiki.libsdl.org/SDL3/SDL_RectsEqualFloat) | [:question:]() | [:question:]() |
| [SDL_HasRectIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_HasRectIntersectionFloat) | [:x:](./methods.go#L3547) | [:x:](./sdl_functions_js.go#L3795) |
| [SDL_GetRectIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectIntersectionFloat) | [:x:](./methods.go#L3554) | [:x:](./sdl_functions_js.go#L3816) |
| [SDL_GetRectUnionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectUnionFloat) | [:x:](./methods.go#L3561) | [:x:](./sdl_functions_js.go#L3842) |
| [SDL_GetRectEnclosingPointsFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPointsFloat) | [:question:]() | [:question:](./sdl_functions_js.go#L3868) |
| [SDL_GetRectAndLineIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersectionFloat) | [:x:](./methods.go#L3568) | [:x:](./sdl_functions_js.go#L3896) |
### Camera

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumCameraDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumCameraDrivers) | [:heavy_check_mark:](./functions.go#L1201) | [:x:](./sdl_functions_js.go#L5236) |
| [SDL_GetCameraDriver](https://wiki.libsdl.org/SDL3/SDL_GetCameraDriver) | [:heavy_check_mark:](./functions.go#L1207) | [:x:](./sdl_functions_js.go#L5247) |
| [SDL_GetCurrentCameraDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentCameraDriver) | [:heavy_check_mark:](./functions.go#L1213) | [:x:](./sdl_functions_js.go#L5260) |
| [SDL_GetCameras](https://wiki.libsdl.org/SDL3/SDL_GetCameras) | [:heavy_check_mark:](./functions.go#L1219) | [:x:](./sdl_functions_js.go#L5271) |
| [SDL_GetCameraSupportedFormats](https://wiki.libsdl.org/SDL3/SDL_GetCameraSupportedFormats) | [:x:](./methods.go#L3104) | [:x:](./sdl_functions_js.go#L5287) |
| [SDL_GetCameraName](https://wiki.libsdl.org/SDL3/SDL_GetCameraName) | [:x:](./methods.go#L3111) | [:x:](./sdl_functions_js.go#L5305) |
| [SDL_GetCameraPosition](https://wiki.libsdl.org/SDL3/SDL_GetCameraPosition) | [:x:](./methods.go#L3118) | [:x:](./sdl_functions_js.go#L5318) |
| [SDL_OpenCamera](https://wiki.libsdl.org/SDL3/SDL_OpenCamera) | [:x:](./methods.go#L3125) | [:x:](./sdl_functions_js.go#L5331) |
| [SDL_GetCameraPermissionState](https://wiki.libsdl.org/SDL3/SDL_GetCameraPermissionState) | [:heavy_check_mark:](./methods.go#L370) | [:x:](./sdl_functions_js.go#L5352) |
| [SDL_GetCameraID](https://wiki.libsdl.org/SDL3/SDL_GetCameraID) | [:heavy_check_mark:](./methods.go#L376) | [:x:](./sdl_functions_js.go#L5368) |
| [SDL_GetCameraProperties](https://wiki.libsdl.org/SDL3/SDL_GetCameraProperties) | [:heavy_check_mark:](./methods.go#L387) | [:x:](./sdl_functions_js.go#L5384) |
| [SDL_GetCameraFormat](https://wiki.libsdl.org/SDL3/SDL_GetCameraFormat) | [:heavy_check_mark:](./methods.go#L398) | [:x:](./sdl_functions_js.go#L5400) |
| [SDL_AcquireCameraFrame](https://wiki.libsdl.org/SDL3/SDL_AcquireCameraFrame) | [:heavy_check_mark:](./methods.go#L410) | [:x:](./sdl_functions_js.go#L5421) |
| [SDL_ReleaseCameraFrame](https://wiki.libsdl.org/SDL3/SDL_ReleaseCameraFrame) | [:heavy_check_mark:](./methods.go#L423) | [:x:](./sdl_functions_js.go#L5445) |
| [SDL_CloseCamera](https://wiki.libsdl.org/SDL3/SDL_CloseCamera) | [:heavy_check_mark:](./methods.go#L429) | [:x:](./sdl_functions_js.go#L5464) |
### Clipboard

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetClipboardText](https://wiki.libsdl.org/SDL3/SDL_SetClipboardText) | [:heavy_check_mark:](./functions.go#L1235) | [:x:](./sdl_functions_js.go#L5478) |
| [SDL_GetClipboardText](https://wiki.libsdl.org/SDL3/SDL_GetClipboardText) | [:heavy_check_mark:](./functions.go#L1245) | [:x:](./sdl_functions_js.go#L5491) |
| [SDL_HasClipboardText](https://wiki.libsdl.org/SDL3/SDL_HasClipboardText) | [:question:]() | [:question:](./sdl_functions_js.go#L5502) |
| [SDL_SetPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_SetPrimarySelectionText) | [:question:]() | [:question:](./sdl_functions_js.go#L5513) |
| [SDL_GetPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_GetPrimarySelectionText) | [:question:]() | [:question:](./sdl_functions_js.go#L5526) |
| [SDL_HasPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_HasPrimarySelectionText) | [:question:]() | [:question:](./sdl_functions_js.go#L5537) |
| [SDL_SetClipboardData](https://wiki.libsdl.org/SDL3/SDL_SetClipboardData) | [:x:](./methods.go#L41) | [:x:](./sdl_functions_js.go#L5548) |
| [SDL_ClearClipboardData](https://wiki.libsdl.org/SDL3/SDL_ClearClipboardData) | [:question:]() | [:question:](./sdl_functions_js.go#L5572) |
| [SDL_GetClipboardData](https://wiki.libsdl.org/SDL3/SDL_GetClipboardData) | [:question:]() | [:question:](./sdl_functions_js.go#L5583) |
| [SDL_HasClipboardData](https://wiki.libsdl.org/SDL3/SDL_HasClipboardData) | [:question:]() | [:question:](./sdl_functions_js.go#L5601) |
| [SDL_GetClipboardMimeTypes](https://wiki.libsdl.org/SDL3/SDL_GetClipboardMimeTypes) | [:question:]() | [:question:](./sdl_functions_js.go#L5614) |
### Dialog

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ShowOpenFileDialog](https://wiki.libsdl.org/SDL3/SDL_ShowOpenFileDialog) | [:question:]() | [:question:](./sdl_functions_js.go#L7690) |
| [SDL_ShowSaveFileDialog](https://wiki.libsdl.org/SDL3/SDL_ShowSaveFileDialog) | [:question:]() | [:question:](./sdl_functions_js.go#L7719) |
| [SDL_ShowOpenFolderDialog](https://wiki.libsdl.org/SDL3/SDL_ShowOpenFolderDialog) | [:question:]() | [:question:](./sdl_functions_js.go#L7746) |
| [SDL_ShowFileDialogWithProperties](https://wiki.libsdl.org/SDL3/SDL_ShowFileDialogWithProperties) | [:x:](./methods.go#L2121) | [:x:](./sdl_functions_js.go#L7768) |
### GPU

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GPUSupportsShaderFormats](https://wiki.libsdl.org/SDL3/SDL_GPUSupportsShaderFormats) | [:heavy_check_mark:](./functions.go#L419) | [:x:](./sdl_functions_js.go#L11384) |
| [SDL_GPUSupportsProperties](https://wiki.libsdl.org/SDL3/SDL_GPUSupportsProperties) | [:x:](./methods.go#L5435) | [:x:](./sdl_functions_js.go#L11399) |
| [SDL_CreateGPUDevice](https://wiki.libsdl.org/SDL3/SDL_CreateGPUDevice) | [:heavy_check_mark:](./functions.go#L425) | [:x:](./sdl_functions_js.go#L11412) |
| [SDL_CreateGPUDeviceWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateGPUDeviceWithProperties) | [:x:](./methods.go#L5442) | [:x:](./sdl_functions_js.go#L11432) |
| [SDL_DestroyGPUDevice](https://wiki.libsdl.org/SDL3/SDL_DestroyGPUDevice) | [:x:](./methods.go#L1704) | [:x:](./sdl_functions_js.go#L11448) |
| [SDL_GetNumGPUDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumGPUDrivers) | [:question:]() | [:question:](./sdl_functions_js.go#L11462) |
| [SDL_GetGPUDriver](https://wiki.libsdl.org/SDL3/SDL_GetGPUDriver) | [:question:]() | [:question:](./sdl_functions_js.go#L11473) |
| [SDL_GetGPUDeviceDriver](https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceDriver) | [:x:](./methods.go#L1711) | [:x:](./sdl_functions_js.go#L11486) |
| [SDL_GetGPUShaderFormats](https://wiki.libsdl.org/SDL3/SDL_GetGPUShaderFormats) | [:x:](./methods.go#L1718) | [:x:](./sdl_functions_js.go#L11502) |
| [SDL_CreateGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_CreateGPUComputePipeline) | [:x:](./methods.go#L1725) | [:x:](./sdl_functions_js.go#L11518) |
| [SDL_CreateGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_CreateGPUGraphicsPipeline) | [:x:](./methods.go#L1732) | [:x:](./sdl_functions_js.go#L11542) |
| [SDL_CreateGPUSampler](https://wiki.libsdl.org/SDL3/SDL_CreateGPUSampler) | [:x:](./methods.go#L1739) | [:x:](./sdl_functions_js.go#L11566) |
| [SDL_CreateGPUShader](https://wiki.libsdl.org/SDL3/SDL_CreateGPUShader) | [:x:](./methods.go#L1746) | [:x:](./sdl_functions_js.go#L11590) |
| [SDL_CreateGPUTexture](https://wiki.libsdl.org/SDL3/SDL_CreateGPUTexture) | [:x:](./methods.go#L1753) | [:x:](./sdl_functions_js.go#L11614) |
| [SDL_CreateGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_CreateGPUBuffer) | [:x:](./methods.go#L1760) | [:x:](./sdl_functions_js.go#L11638) |
| [SDL_CreateGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_CreateGPUTransferBuffer) | [:x:](./methods.go#L1767) | [:x:](./sdl_functions_js.go#L11662) |
| [SDL_SetGPUBufferName](https://wiki.libsdl.org/SDL3/SDL_SetGPUBufferName) | [:x:](./methods.go#L1774) | [:x:](./sdl_functions_js.go#L11686) |
| [SDL_SetGPUTextureName](https://wiki.libsdl.org/SDL3/SDL_SetGPUTextureName) | [:x:](./methods.go#L1781) | [:x:](./sdl_functions_js.go#L11707) |
| [SDL_InsertGPUDebugLabel](https://wiki.libsdl.org/SDL3/SDL_InsertGPUDebugLabel) | [:heavy_check_mark:](./methods.go#L5557) | [:x:](./sdl_functions_js.go#L11728) |
| [SDL_PushGPUDebugGroup](https://wiki.libsdl.org/SDL3/SDL_PushGPUDebugGroup) | [:heavy_check_mark:](./methods.go#L5563) | [:x:](./sdl_functions_js.go#L11744) |
| [SDL_PopGPUDebugGroup](https://wiki.libsdl.org/SDL3/SDL_PopGPUDebugGroup) | [:heavy_check_mark:](./methods.go#L5569) | [:x:](./sdl_functions_js.go#L11760) |
| [SDL_ReleaseGPUTexture](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTexture) | [:x:](./methods.go#L1788) | [:x:](./sdl_functions_js.go#L11774) |
| [SDL_ReleaseGPUSampler](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUSampler) | [:x:](./methods.go#L1795) | [:x:](./sdl_functions_js.go#L11793) |
| [SDL_ReleaseGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUBuffer) | [:x:](./methods.go#L1802) | [:x:](./sdl_functions_js.go#L11812) |
| [SDL_ReleaseGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTransferBuffer) | [:x:](./methods.go#L1809) | [:x:](./sdl_functions_js.go#L11831) |
| [SDL_ReleaseGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUComputePipeline) | [:x:](./methods.go#L1816) | [:x:](./sdl_functions_js.go#L11850) |
| [SDL_ReleaseGPUShader](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUShader) | [:x:](./methods.go#L1823) | [:x:](./sdl_functions_js.go#L11869) |
| [SDL_ReleaseGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUGraphicsPipeline) | [:x:](./methods.go#L1830) | [:x:](./sdl_functions_js.go#L11888) |
| [SDL_AcquireGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_AcquireGPUCommandBuffer) | [:x:](./methods.go#L1837) | [:x:](./sdl_functions_js.go#L11907) |
| [SDL_PushGPUVertexUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUVertexUniformData) | [:heavy_check_mark:](./methods.go#L5575) | [:x:](./sdl_functions_js.go#L11926) |
| [SDL_PushGPUFragmentUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUFragmentUniformData) | [:heavy_check_mark:](./methods.go#L5581) | [:x:](./sdl_functions_js.go#L11946) |
| [SDL_PushGPUComputeUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUComputeUniformData) | [:heavy_check_mark:](./methods.go#L5587) | [:x:](./sdl_functions_js.go#L11966) |
| [SDL_BeginGPURenderPass](https://wiki.libsdl.org/SDL3/SDL_BeginGPURenderPass) | [:heavy_check_mark:](./methods.go#L5593) | [:x:](./sdl_functions_js.go#L11986) |
| [SDL_BindGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_BindGPUGraphicsPipeline) | [:x:](./methods.go#L1153) | [:x:](./sdl_functions_js.go#L12017) |
| [SDL_SetGPUViewport](https://wiki.libsdl.org/SDL3/SDL_SetGPUViewport) | [:x:](./methods.go#L1160) | [:x:](./sdl_functions_js.go#L12036) |
| [SDL_SetGPUScissor](https://wiki.libsdl.org/SDL3/SDL_SetGPUScissor) | [:x:](./methods.go#L1167) | [:x:](./sdl_functions_js.go#L12055) |
| [SDL_SetGPUBlendConstants](https://wiki.libsdl.org/SDL3/SDL_SetGPUBlendConstants) | [:question:]() | [:question:]() |
| [SDL_SetGPUStencilReference](https://wiki.libsdl.org/SDL3/SDL_SetGPUStencilReference) | [:x:](./methods.go#L1174) | [:x:](./sdl_functions_js.go#L12074) |
| [SDL_BindGPUVertexBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexBuffers) | [:x:](./methods.go#L1181) | [:x:](./sdl_functions_js.go#L12090) |
| [SDL_BindGPUIndexBuffer](https://wiki.libsdl.org/SDL3/SDL_BindGPUIndexBuffer) | [:x:](./methods.go#L1188) | [:x:](./sdl_functions_js.go#L12113) |
| [SDL_BindGPUVertexSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexSamplers) | [:x:](./methods.go#L1195) | [:x:](./sdl_functions_js.go#L12134) |
| [SDL_BindGPUVertexStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageTextures) | [:x:](./methods.go#L1202) | [:x:](./sdl_functions_js.go#L12157) |
| [SDL_BindGPUVertexStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageBuffers) | [:x:](./methods.go#L1209) | [:x:](./sdl_functions_js.go#L12180) |
| [SDL_BindGPUFragmentSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentSamplers) | [:x:](./methods.go#L1216) | [:x:](./sdl_functions_js.go#L12203) |
| [SDL_BindGPUFragmentStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageTextures) | [:x:](./methods.go#L1223) | [:x:](./sdl_functions_js.go#L12226) |
| [SDL_BindGPUFragmentStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageBuffers) | [:x:](./methods.go#L1230) | [:x:](./sdl_functions_js.go#L12249) |
| [SDL_DrawGPUIndexedPrimitives](https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitives) | [:x:](./methods.go#L1237) | [:x:](./sdl_functions_js.go#L12272) |
| [SDL_DrawGPUPrimitives](https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitives) | [:x:](./methods.go#L1244) | [:x:](./sdl_functions_js.go#L12296) |
| [SDL_DrawGPUPrimitivesIndirect](https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitivesIndirect) | [:x:](./methods.go#L1251) | [:x:](./sdl_functions_js.go#L12318) |
| [SDL_DrawGPUIndexedPrimitivesIndirect](https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitivesIndirect) | [:x:](./methods.go#L1258) | [:x:](./sdl_functions_js.go#L12341) |
| [SDL_EndGPURenderPass](https://wiki.libsdl.org/SDL3/SDL_EndGPURenderPass) | [:x:](./methods.go#L1265) | [:x:](./sdl_functions_js.go#L12364) |
| [SDL_BeginGPUComputePass](https://wiki.libsdl.org/SDL3/SDL_BeginGPUComputePass) | [:heavy_check_mark:](./methods.go#L5599) | [:x:](./sdl_functions_js.go#L12378) |
| [SDL_BindGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputePipeline) | [:x:](./methods.go#L865) | [:x:](./sdl_functions_js.go#L12411) |
| [SDL_BindGPUComputeSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeSamplers) | [:x:](./methods.go#L872) | [:x:](./sdl_functions_js.go#L12430) |
| [SDL_BindGPUComputeStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageTextures) | [:x:](./methods.go#L879) | [:x:](./sdl_functions_js.go#L12453) |
| [SDL_BindGPUComputeStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageBuffers) | [:x:](./methods.go#L886) | [:x:](./sdl_functions_js.go#L12476) |
| [SDL_DispatchGPUCompute](https://wiki.libsdl.org/SDL3/SDL_DispatchGPUCompute) | [:x:](./methods.go#L893) | [:x:](./sdl_functions_js.go#L12499) |
| [SDL_DispatchGPUComputeIndirect](https://wiki.libsdl.org/SDL3/SDL_DispatchGPUComputeIndirect) | [:x:](./methods.go#L900) | [:x:](./sdl_functions_js.go#L12519) |
| [SDL_EndGPUComputePass](https://wiki.libsdl.org/SDL3/SDL_EndGPUComputePass) | [:x:](./methods.go#L907) | [:x:](./sdl_functions_js.go#L12540) |
| [SDL_MapGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_MapGPUTransferBuffer) | [:x:](./methods.go#L1844) | [:x:](./sdl_functions_js.go#L12554) |
| [SDL_UnmapGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_UnmapGPUTransferBuffer) | [:x:](./methods.go#L1851) | [:x:](./sdl_functions_js.go#L12577) |
| [SDL_BeginGPUCopyPass](https://wiki.libsdl.org/SDL3/SDL_BeginGPUCopyPass) | [:heavy_check_mark:](./methods.go#L5605) | [:x:](./sdl_functions_js.go#L12596) |
| [SDL_UploadToGPUTexture](https://wiki.libsdl.org/SDL3/SDL_UploadToGPUTexture) | [:x:](./methods.go#L2428) | [:x:](./sdl_functions_js.go#L12615) |
| [SDL_UploadToGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_UploadToGPUBuffer) | [:x:](./methods.go#L2435) | [:x:](./sdl_functions_js.go#L12641) |
| [SDL_CopyGPUTextureToTexture](https://wiki.libsdl.org/SDL3/SDL_CopyGPUTextureToTexture) | [:x:](./methods.go#L2442) | [:x:](./sdl_functions_js.go#L12667) |
| [SDL_CopyGPUBufferToBuffer](https://wiki.libsdl.org/SDL3/SDL_CopyGPUBufferToBuffer) | [:x:](./methods.go#L2449) | [:x:](./sdl_functions_js.go#L12699) |
| [SDL_DownloadFromGPUTexture](https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUTexture) | [:x:](./methods.go#L2456) | [:x:](./sdl_functions_js.go#L12727) |
| [SDL_DownloadFromGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUBuffer) | [:x:](./methods.go#L2463) | [:x:](./sdl_functions_js.go#L12751) |
| [SDL_EndGPUCopyPass](https://wiki.libsdl.org/SDL3/SDL_EndGPUCopyPass) | [:x:](./methods.go#L2470) | [:x:](./sdl_functions_js.go#L12775) |
| [SDL_GenerateMipmapsForGPUTexture](https://wiki.libsdl.org/SDL3/SDL_GenerateMipmapsForGPUTexture) | [:heavy_check_mark:](./methods.go#L5611) | [:x:](./sdl_functions_js.go#L12789) |
| [SDL_BlitGPUTexture](https://wiki.libsdl.org/SDL3/SDL_BlitGPUTexture) | [:heavy_check_mark:](./methods.go#L5617) | [:x:](./sdl_functions_js.go#L12808) |
| [SDL_WindowSupportsGPUSwapchainComposition](https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUSwapchainComposition) | [:x:](./methods.go#L1858) | [:x:](./sdl_functions_js.go#L12827) |
| [SDL_WindowSupportsGPUPresentMode](https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUPresentMode) | [:x:](./methods.go#L1865) | [:x:](./sdl_functions_js.go#L12850) |
| [SDL_ClaimWindowForGPUDevice](https://wiki.libsdl.org/SDL3/SDL_ClaimWindowForGPUDevice) | [:x:](./methods.go#L1872) | [:x:](./sdl_functions_js.go#L12873) |
| [SDL_ReleaseWindowFromGPUDevice](https://wiki.libsdl.org/SDL3/SDL_ReleaseWindowFromGPUDevice) | [:x:](./methods.go#L1879) | [:x:](./sdl_functions_js.go#L12894) |
| [SDL_SetGPUSwapchainParameters](https://wiki.libsdl.org/SDL3/SDL_SetGPUSwapchainParameters) | [:x:](./methods.go#L1886) | [:x:](./sdl_functions_js.go#L12913) |
| [SDL_SetGPUAllowedFramesInFlight](https://wiki.libsdl.org/SDL3/SDL_SetGPUAllowedFramesInFlight) | [:x:](./methods.go#L1893) | [:x:](./sdl_functions_js.go#L12938) |
| [SDL_GetGPUSwapchainTextureFormat](https://wiki.libsdl.org/SDL3/SDL_GetGPUSwapchainTextureFormat) | [:x:](./methods.go#L1900) | [:x:](./sdl_functions_js.go#L12956) |
| [SDL_AcquireGPUSwapchainTexture](https://wiki.libsdl.org/SDL3/SDL_AcquireGPUSwapchainTexture) | [:heavy_check_mark:](./methods.go#L5623) | [:x:](./sdl_functions_js.go#L12977) |
| [SDL_WaitForGPUSwapchain](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUSwapchain) | [:x:](./methods.go#L1907) | [:x:](./sdl_functions_js.go#L13013) |
| [SDL_WaitAndAcquireGPUSwapchainTexture](https://wiki.libsdl.org/SDL3/SDL_WaitAndAcquireGPUSwapchainTexture) | [:heavy_check_mark:](./methods.go#L5635) | [:x:](./sdl_functions_js.go#L13034) |
| [SDL_SubmitGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBuffer) | [:heavy_check_mark:](./methods.go#L5647) | [:x:](./sdl_functions_js.go#L13070) |
| [SDL_SubmitGPUCommandBufferAndAcquireFence](https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBufferAndAcquireFence) | [:heavy_check_mark:](./methods.go#L5657) | [:x:](./sdl_functions_js.go#L13086) |
| [SDL_CancelGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_CancelGPUCommandBuffer) | [:heavy_check_mark:](./methods.go#L5668) | [:x:](./sdl_functions_js.go#L13105) |
| [SDL_WaitForGPUIdle](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUIdle) | [:x:](./methods.go#L1914) | [:x:](./sdl_functions_js.go#L13121) |
| [SDL_WaitForGPUFences](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUFences) | [:x:](./methods.go#L1921) | [:x:](./sdl_functions_js.go#L13137) |
| [SDL_QueryGPUFence](https://wiki.libsdl.org/SDL3/SDL_QueryGPUFence) | [:x:](./methods.go#L1928) | [:x:](./sdl_functions_js.go#L13162) |
| [SDL_ReleaseGPUFence](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUFence) | [:x:](./methods.go#L1935) | [:x:](./sdl_functions_js.go#L13183) |
| [SDL_GPUTextureFormatTexelBlockSize](https://wiki.libsdl.org/SDL3/SDL_GPUTextureFormatTexelBlockSize) | [:heavy_check_mark:](./methods.go#L445) | [:x:](./sdl_functions_js.go#L13202) |
| [SDL_GPUTextureSupportsFormat](https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsFormat) | [:x:](./methods.go#L1942) | [:x:](./sdl_functions_js.go#L13215) |
| [SDL_GPUTextureSupportsSampleCount](https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsSampleCount) | [:x:](./methods.go#L1949) | [:x:](./sdl_functions_js.go#L13237) |
| [SDL_CalculateGPUTextureFormatSize](https://wiki.libsdl.org/SDL3/SDL_CalculateGPUTextureFormatSize) | [:heavy_check_mark:](./methods.go#L451) | [:x:](./sdl_functions_js.go#L13257) |
| [SDL_GDKSuspendGPU](https://wiki.libsdl.org/SDL3/SDL_GDKSuspendGPU) | [:question:]() | [:question:]() |
| [SDL_GDKResumeGPU](https://wiki.libsdl.org/SDL3/SDL_GDKResumeGPU) | [:question:]() | [:question:]() |
### MessageBox

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ShowMessageBox](https://wiki.libsdl.org/SDL3/SDL_ShowMessageBox) | [:question:]() | [:question:](./sdl_functions_js.go#L14336) |
| [SDL_ShowSimpleMessageBox](https://wiki.libsdl.org/SDL3/SDL_ShowSimpleMessageBox) | [:heavy_check_mark:](./functions.go#L718) | [:x:](./sdl_functions_js.go#L14357) |
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
| [SDL_Metal_CreateView](https://wiki.libsdl.org/SDL3/SDL_Metal_CreateView) | [:x:](./methods.go#L4341) | [:x:](./sdl_functions_js.go#L14379) |
| [SDL_Metal_DestroyView](https://wiki.libsdl.org/SDL3/SDL_Metal_DestroyView) | [:question:]() | [:question:](./sdl_functions_js.go#L14395) |
| [SDL_Metal_GetLayer](https://wiki.libsdl.org/SDL3/SDL_Metal_GetLayer) | [:question:]() | [:question:](./sdl_functions_js.go#L14406) |
### Platform

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetPowerInfo) | [:heavy_check_mark:](./functions.go#L736) | [:x:](./sdl_functions_js.go#L7813) |
### Power

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetSensors](https://wiki.libsdl.org/SDL3/SDL_GetSensors) | [:heavy_check_mark:](./functions.go#L751) | [:x:](./sdl_functions_js.go#L7834) |
| [SDL_GetSensorNameForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorNameForID) | [:heavy_check_mark:](./methods.go#L5759) | [:x:](./sdl_functions_js.go#L7850) |
| [SDL_GetSensorTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorTypeForID) | [:heavy_check_mark:](./methods.go#L5765) | [:x:](./sdl_functions_js.go#L7863) |
| [SDL_GetSensorNonPortableTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableTypeForID) | [:heavy_check_mark:](./methods.go#L5771) | [:x:](./sdl_functions_js.go#L7876) |
| [SDL_OpenSensor](https://wiki.libsdl.org/SDL3/SDL_OpenSensor) | [:heavy_check_mark:](./methods.go#L5777) | [:x:](./sdl_functions_js.go#L7889) |
| [SDL_GetSensorFromID](https://wiki.libsdl.org/SDL3/SDL_GetSensorFromID) | [:heavy_check_mark:](./methods.go#L5788) | [:x:](./sdl_functions_js.go#L7905) |
| [SDL_GetSensorProperties](https://wiki.libsdl.org/SDL3/SDL_GetSensorProperties) | [:x:](./methods.go#L519) | [:x:](./sdl_functions_js.go#L7921) |
| [SDL_GetSensorName](https://wiki.libsdl.org/SDL3/SDL_GetSensorName) | [:x:](./methods.go#L526) | [:x:](./sdl_functions_js.go#L7937) |
| [SDL_GetSensorType](https://wiki.libsdl.org/SDL3/SDL_GetSensorType) | [:x:](./methods.go#L533) | [:x:](./sdl_functions_js.go#L7953) |
| [SDL_GetSensorNonPortableType](https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableType) | [:x:](./methods.go#L540) | [:x:](./sdl_functions_js.go#L7969) |
| [SDL_GetSensorID](https://wiki.libsdl.org/SDL3/SDL_GetSensorID) | [:x:](./methods.go#L547) | [:x:](./sdl_functions_js.go#L7985) |
| [SDL_GetSensorData](https://wiki.libsdl.org/SDL3/SDL_GetSensorData) | [:x:](./methods.go#L554) | [:x:](./sdl_functions_js.go#L8001) |
| [SDL_CloseSensor](https://wiki.libsdl.org/SDL3/SDL_CloseSensor) | [:x:](./methods.go#L561) | [:x:](./sdl_functions_js.go#L8024) |
| [SDL_UpdateSensors](https://wiki.libsdl.org/SDL3/SDL_UpdateSensors) | [:heavy_check_mark:](./functions.go#L765) | [:x:](./sdl_functions_js.go#L8038) |
### Sensor

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateProcess](https://wiki.libsdl.org/SDL3/SDL_CreateProcess) | [:question:]() | [:question:](./sdl_functions_js.go#L14432) |
| [SDL_CreateProcessWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateProcessWithProperties) | [:x:](./methods.go#L5449) | [:x:](./sdl_functions_js.go#L14452) |
| [SDL_GetProcessProperties](https://wiki.libsdl.org/SDL3/SDL_GetProcessProperties) | [:heavy_check_mark:](./methods.go#L5680) | [:x:](./sdl_functions_js.go#L14468) |
| [SDL_ReadProcess](https://wiki.libsdl.org/SDL3/SDL_ReadProcess) | [:heavy_check_mark:](./methods.go#L5691) | [:x:](./sdl_functions_js.go#L14484) |
| [SDL_GetProcessInput](https://wiki.libsdl.org/SDL3/SDL_GetProcessInput) | [:heavy_check_mark:](./methods.go#L5709) | [:x:](./sdl_functions_js.go#L14510) |
| [SDL_GetProcessOutput](https://wiki.libsdl.org/SDL3/SDL_GetProcessOutput) | [:heavy_check_mark:](./methods.go#L5720) | [:x:](./sdl_functions_js.go#L14529) |
| [SDL_KillProcess](https://wiki.libsdl.org/SDL3/SDL_KillProcess) | [:heavy_check_mark:](./methods.go#L5731) | [:x:](./sdl_functions_js.go#L14548) |
| [SDL_WaitProcess](https://wiki.libsdl.org/SDL3/SDL_WaitProcess) | [:heavy_check_mark:](./methods.go#L5741) | [:x:](./sdl_functions_js.go#L14566) |
| [SDL_DestroyProcess](https://wiki.libsdl.org/SDL3/SDL_DestroyProcess) | [:heavy_check_mark:](./methods.go#L5751) | [:x:](./sdl_functions_js.go#L14589) |
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
| [SDL_GetNumLogicalCPUCores](https://wiki.libsdl.org/SDL3/SDL_GetNumLogicalCPUCores) | [:question:]() | [:question:](./sdl_functions_js.go#L5630) |
| [SDL_GetCPUCacheLineSize](https://wiki.libsdl.org/SDL3/SDL_GetCPUCacheLineSize) | [:question:]() | [:question:](./sdl_functions_js.go#L5641) |
| [SDL_HasAltiVec](https://wiki.libsdl.org/SDL3/SDL_HasAltiVec) | [:question:]() | [:question:](./sdl_functions_js.go#L5652) |
| [SDL_HasMMX](https://wiki.libsdl.org/SDL3/SDL_HasMMX) | [:question:]() | [:question:](./sdl_functions_js.go#L5663) |
| [SDL_HasSSE](https://wiki.libsdl.org/SDL3/SDL_HasSSE) | [:question:]() | [:question:](./sdl_functions_js.go#L5674) |
| [SDL_HasSSE2](https://wiki.libsdl.org/SDL3/SDL_HasSSE2) | [:question:]() | [:question:](./sdl_functions_js.go#L5685) |
| [SDL_HasSSE3](https://wiki.libsdl.org/SDL3/SDL_HasSSE3) | [:question:]() | [:question:](./sdl_functions_js.go#L5696) |
| [SDL_HasSSE41](https://wiki.libsdl.org/SDL3/SDL_HasSSE41) | [:question:]() | [:question:](./sdl_functions_js.go#L5707) |
| [SDL_HasSSE42](https://wiki.libsdl.org/SDL3/SDL_HasSSE42) | [:question:]() | [:question:](./sdl_functions_js.go#L5718) |
| [SDL_HasAVX](https://wiki.libsdl.org/SDL3/SDL_HasAVX) | [:question:]() | [:question:](./sdl_functions_js.go#L5729) |
| [SDL_HasAVX2](https://wiki.libsdl.org/SDL3/SDL_HasAVX2) | [:question:]() | [:question:](./sdl_functions_js.go#L5740) |
| [SDL_HasAVX512F](https://wiki.libsdl.org/SDL3/SDL_HasAVX512F) | [:question:]() | [:question:](./sdl_functions_js.go#L5751) |
| [SDL_HasARMSIMD](https://wiki.libsdl.org/SDL3/SDL_HasARMSIMD) | [:question:]() | [:question:](./sdl_functions_js.go#L5762) |
| [SDL_HasNEON](https://wiki.libsdl.org/SDL3/SDL_HasNEON) | [:question:]() | [:question:](./sdl_functions_js.go#L5773) |
| [SDL_HasLSX](https://wiki.libsdl.org/SDL3/SDL_HasLSX) | [:question:]() | [:question:](./sdl_functions_js.go#L5784) |
| [SDL_HasLASX](https://wiki.libsdl.org/SDL3/SDL_HasLASX) | [:question:]() | [:question:](./sdl_functions_js.go#L5795) |
| [SDL_GetSystemRAM](https://wiki.libsdl.org/SDL3/SDL_GetSystemRAM) | [:question:]() | [:question:](./sdl_functions_js.go#L5806) |
| [SDL_GetSIMDAlignment](https://wiki.libsdl.org/SDL3/SDL_GetSIMDAlignment) | [:question:]() | [:question:](./sdl_functions_js.go#L5817) |
### CPUInfo

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPreferredLocales](https://wiki.libsdl.org/SDL3/SDL_GetPreferredLocales) | [:heavy_check_mark:](./functions.go#L793) | [:x:](./sdl_functions_js.go#L14082) |
### Intrinsics

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetWindowsMessageHook](https://wiki.libsdl.org/SDL3/SDL_SetWindowsMessageHook) | [:question:]() | [:question:]() |
| [SDL_GetDirect3D9AdapterIndex](https://wiki.libsdl.org/SDL3/SDL_GetDirect3D9AdapterIndex) | [:question:]() | [:question:]() |
| [SDL_GetDXGIOutputInfo](https://wiki.libsdl.org/SDL3/SDL_GetDXGIOutputInfo) | [:question:]() | [:question:]() |
| [SDL_SetX11EventHook](https://wiki.libsdl.org/SDL3/SDL_SetX11EventHook) | [:x:](./methods.go#L588) | [:x:](./sdl_functions_js.go#L16882) |
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
| [SDL_IsTablet](https://wiki.libsdl.org/SDL3/SDL_IsTablet) | [:question:]() | [:question:](./sdl_functions_js.go#L16895) |
| [SDL_IsTV](https://wiki.libsdl.org/SDL3/SDL_IsTV) | [:question:]() | [:question:](./sdl_functions_js.go#L16906) |
| [SDL_GetSandbox](https://wiki.libsdl.org/SDL3/SDL_GetSandbox) | [:question:]() | [:question:](./sdl_functions_js.go#L16917) |
| [SDL_OnApplicationWillTerminate](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillTerminate) | [:question:]() | [:question:](./sdl_functions_js.go#L16928) |
| [SDL_OnApplicationDidReceiveMemoryWarning](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidReceiveMemoryWarning) | [:question:]() | [:question:](./sdl_functions_js.go#L16937) |
| [SDL_OnApplicationWillEnterBackground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillEnterBackground) | [:question:]() | [:question:](./sdl_functions_js.go#L16946) |
| [SDL_OnApplicationDidEnterBackground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidEnterBackground) | [:question:]() | [:question:](./sdl_functions_js.go#L16955) |
| [SDL_OnApplicationWillEnterForeground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillEnterForeground) | [:question:]() | [:question:](./sdl_functions_js.go#L16964) |
| [SDL_OnApplicationDidEnterForeground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidEnterForeground) | [:question:]() | [:question:](./sdl_functions_js.go#L16973) |
| [SDL_OnApplicationDidChangeStatusBarOrientation](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidChangeStatusBarOrientation) | [:question:]() | [:question:]() |
| [SDL_GetGDKTaskQueue](https://wiki.libsdl.org/SDL3/SDL_GetGDKTaskQueue) | [:question:]() | [:question:]() |
| [SDL_GetGDKDefaultUser](https://wiki.libsdl.org/SDL3/SDL_GetGDKDefaultUser) | [:question:]() | [:question:]() |
### Locale

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_OpenURL](https://wiki.libsdl.org/SDL3/SDL_OpenURL) | [:question:]() | [:question:](./sdl_functions_js.go#L14419) |
### System

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GUIDToString](https://wiki.libsdl.org/SDL3/SDL_GUIDToString) | [:question:]() | [:question:](./sdl_functions_js.go#L7785) |
| [SDL_StringToGUID](https://wiki.libsdl.org/SDL3/SDL_StringToGUID) | [:question:]() | [:question:](./sdl_functions_js.go#L7800) |
### Misc

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AppInit](https://wiki.libsdl.org/SDL3/SDL_AppInit) | [:question:]() | [:question:]() |
| [SDL_AppIterate](https://wiki.libsdl.org/SDL3/SDL_AppIterate) | [:question:]() | [:question:]() |
| [SDL_AppEvent](https://wiki.libsdl.org/SDL3/SDL_AppEvent) | [:question:]() | [:question:]() |
| [SDL_AppQuit](https://wiki.libsdl.org/SDL3/SDL_AppQuit) | [:question:]() | [:question:]() |
| [SDL_main](https://wiki.libsdl.org/SDL3/SDL_main) | [:question:]() | [:question:]() |
| [SDL_SetMainReady](https://wiki.libsdl.org/SDL3/SDL_SetMainReady) | [:question:]() | [:question:](./sdl_functions_js.go#L17266) |
| [SDL_RunApp](https://wiki.libsdl.org/SDL3/SDL_RunApp) | [:question:]() | [:question:](./sdl_functions_js.go#L17275) |
| [SDL_EnterAppMainCallbacks](https://wiki.libsdl.org/SDL3/SDL_EnterAppMainCallbacks) | [:question:]() | [:question:](./sdl_functions_js.go#L17297) |
| [SDL_RegisterApp](https://wiki.libsdl.org/SDL3/SDL_RegisterApp) | [:question:]() | [:question:]() |
| [SDL_UnregisterApp](https://wiki.libsdl.org/SDL3/SDL_UnregisterApp) | [:question:]() | [:question:]() |
| [SDL_GDKSuspendComplete](https://wiki.libsdl.org/SDL3/SDL_GDKSuspendComplete) | [:question:]() | [:question:](./sdl_functions_js.go#L17323) |
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
| [TTF_Version](https://wiki.libsdl.org/SDL3_ttf/TTF_Version) | [:question:]() | [:question:]() |
| [TTF_GetFreeTypeVersion](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFreeTypeVersion) | [:question:]() | [:question:]() |
| [TTF_GetHarfBuzzVersion](https://wiki.libsdl.org/SDL3_ttf/TTF_GetHarfBuzzVersion) | [:question:]() | [:question:]() |
| [TTF_Init](https://wiki.libsdl.org/SDL3_ttf/TTF_Init) | [:question:]() | [:question:]() |
| [TTF_OpenFont](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFont) | [:question:]() | [:question:]() |
| [TTF_OpenFontIO](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontIO) | [:question:]() | [:question:]() |
| [TTF_OpenFontWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontWithProperties) | [:question:]() | [:question:]() |
| [TTF_CopyFont](https://wiki.libsdl.org/SDL3_ttf/TTF_CopyFont) | [:heavy_check_mark:](./ttf/methods.go#L409) | [:heavy_check_mark:]() |
| [TTF_GetFontProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontProperties) | [:heavy_check_mark:](./ttf/methods.go#L420) | [:heavy_check_mark:]() |
| [TTF_GetFontGeneration](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontGeneration) | [:heavy_check_mark:](./ttf/methods.go#L426) | [:heavy_check_mark:]() |
| [TTF_AddFallbackFont](https://wiki.libsdl.org/SDL3_ttf/TTF_AddFallbackFont) | [:heavy_check_mark:](./ttf/methods.go#L432) | [:heavy_check_mark:]() |
| [TTF_RemoveFallbackFont](https://wiki.libsdl.org/SDL3_ttf/TTF_RemoveFallbackFont) | [:heavy_check_mark:](./ttf/methods.go#L442) | [:heavy_check_mark:]() |
| [TTF_ClearFallbackFonts](https://wiki.libsdl.org/SDL3_ttf/TTF_ClearFallbackFonts) | [:heavy_check_mark:](./ttf/methods.go#L448) | [:heavy_check_mark:]() |
| [TTF_SetFontSize](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSize) | [:heavy_check_mark:](./ttf/methods.go#L454) | [:heavy_check_mark:]() |
| [TTF_SetFontSizeDPI](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSizeDPI) | [:heavy_check_mark:](./ttf/methods.go#L464) | [:heavy_check_mark:]() |
| [TTF_GetFontSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSize) | [:heavy_check_mark:](./ttf/methods.go#L474) | [:heavy_check_mark:]() |
| [TTF_GetFontDPI](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDPI) | [:heavy_check_mark:](./ttf/methods.go#L485) | [:heavy_check_mark:]() |
| [TTF_SetFontStyle](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontStyle) | [:heavy_check_mark:](./ttf/methods.go#L497) | [:heavy_check_mark:]() |
| [TTF_GetFontStyle](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyle) | [:heavy_check_mark:](./ttf/methods.go#L503) | [:heavy_check_mark:]() |
| [TTF_SetFontOutline](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontOutline) | [:heavy_check_mark:](./ttf/methods.go#L509) | [:heavy_check_mark:]() |
| [TTF_GetFontOutline](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontOutline) | [:heavy_check_mark:](./ttf/methods.go#L519) | [:heavy_check_mark:]() |
| [TTF_SetFontHinting](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontHinting) | [:heavy_check_mark:](./ttf/methods.go#L525) | [:heavy_check_mark:]() |
| [TTF_GetNumFontFaces](https://wiki.libsdl.org/SDL3_ttf/TTF_GetNumFontFaces) | [:heavy_check_mark:](./ttf/methods.go#L531) | [:heavy_check_mark:]() |
| [TTF_GetFontHinting](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHinting) | [:heavy_check_mark:](./ttf/methods.go#L537) | [:heavy_check_mark:]() |
| [TTF_SetFontSDF](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSDF) | [:question:]() | [:question:]() |
| [TTF_GetFontSDF](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSDF) | [:question:]() | [:question:]() |
| [TTF_SetFontWrapAlignment](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontWrapAlignment) | [:heavy_check_mark:](./ttf/methods.go#L543) | [:heavy_check_mark:]() |
| [TTF_GetFontWrapAlignment](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontWrapAlignment) | [:heavy_check_mark:](./ttf/methods.go#L549) | [:heavy_check_mark:]() |
| [TTF_GetFontHeight](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHeight) | [:heavy_check_mark:](./ttf/methods.go#L555) | [:heavy_check_mark:]() |
| [TTF_GetFontAscent](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontAscent) | [:heavy_check_mark:](./ttf/methods.go#L561) | [:heavy_check_mark:]() |
| [TTF_GetFontDescent](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDescent) | [:heavy_check_mark:](./ttf/methods.go#L567) | [:heavy_check_mark:]() |
| [TTF_SetFontLineSkip](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLineSkip) | [:heavy_check_mark:](./ttf/methods.go#L573) | [:heavy_check_mark:]() |
| [TTF_GetFontLineSkip](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontLineSkip) | [:heavy_check_mark:](./ttf/methods.go#L579) | [:heavy_check_mark:]() |
| [TTF_SetFontKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontKerning) | [:heavy_check_mark:](./ttf/methods.go#L585) | [:heavy_check_mark:]() |
| [TTF_GetFontKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontKerning) | [:heavy_check_mark:](./ttf/methods.go#L591) | [:heavy_check_mark:]() |
| [TTF_FontIsFixedWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsFixedWidth) | [:heavy_check_mark:](./ttf/methods.go#L597) | [:heavy_check_mark:]() |
| [TTF_FontIsScalable](https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsScalable) | [:question:]() | [:question:]() |
| [TTF_GetFontFamilyName](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontFamilyName) | [:heavy_check_mark:](./ttf/methods.go#L603) | [:heavy_check_mark:]() |
| [TTF_GetFontStyleName](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyleName) | [:heavy_check_mark:](./ttf/methods.go#L609) | [:heavy_check_mark:]() |
| [TTF_SetFontDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontDirection) | [:heavy_check_mark:](./ttf/methods.go#L615) | [:heavy_check_mark:]() |
| [TTF_GetFontDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDirection) | [:heavy_check_mark:](./ttf/methods.go#L625) | [:heavy_check_mark:]() |
| [TTF_StringToTag](https://wiki.libsdl.org/SDL3_ttf/TTF_StringToTag) | [:question:]() | [:question:]() |
| [TTF_TagToString](https://wiki.libsdl.org/SDL3_ttf/TTF_TagToString) | [:question:]() | [:question:]() |
| [TTF_SetFontScript](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontScript) | [:heavy_check_mark:](./ttf/methods.go#L631) | [:heavy_check_mark:]() |
| [TTF_GetFontScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontScript) | [:heavy_check_mark:](./ttf/methods.go#L641) | [:heavy_check_mark:]() |
| [TTF_GetGlyphScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphScript) | [:question:]() | [:question:]() |
| [TTF_SetFontLanguage](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLanguage) | [:question:]() | [:question:]() |
| [TTF_FontHasGlyph](https://wiki.libsdl.org/SDL3_ttf/TTF_FontHasGlyph) | [:heavy_check_mark:](./ttf/methods.go#L647) | [:heavy_check_mark:]() |
| [TTF_GetGlyphImage](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImage) | [:heavy_check_mark:](./ttf/methods.go#L653) | [:heavy_check_mark:]() |
| [TTF_GetGlyphImageForIndex](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImageForIndex) | [:heavy_check_mark:](./ttf/methods.go#L666) | [:heavy_check_mark:]() |
| [TTF_GetGlyphMetrics](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphMetrics) | [:heavy_check_mark:](./ttf/methods.go#L679) | [:heavy_check_mark:]() |
| [TTF_GetGlyphKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphKerning) | [:question:]() | [:question:]() |
| [TTF_GetStringSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSize) | [:heavy_check_mark:](./ttf/methods.go#L691) | [:heavy_check_mark:]() |
| [TTF_GetStringSizeWrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSizeWrapped) | [:heavy_check_mark:](./ttf/methods.go#L703) | [:heavy_check_mark:]() |
| [TTF_MeasureString](https://wiki.libsdl.org/SDL3_ttf/TTF_MeasureString) | [:heavy_check_mark:](./ttf/methods.go#L716) | [:heavy_check_mark:]() |
| [TTF_RenderText_Solid](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid) | [:question:]() | [:question:]() |
| [TTF_RenderText_Solid_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid_Wrapped) | [:question:]() | [:question:]() |
| [TTF_RenderGlyph_Solid](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Solid) | [:question:]() | [:question:]() |
| [TTF_RenderText_Shaded](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded) | [:question:]() | [:question:]() |
| [TTF_RenderText_Shaded_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded_Wrapped) | [:question:]() | [:question:]() |
| [TTF_RenderGlyph_Shaded](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Shaded) | [:question:]() | [:question:]() |
| [TTF_RenderText_Blended](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended) | [:question:]() | [:question:]() |
| [TTF_RenderText_Blended_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended_Wrapped) | [:question:]() | [:question:]() |
| [TTF_RenderGlyph_Blended](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Blended) | [:question:]() | [:question:]() |
| [TTF_RenderText_LCD](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD) | [:question:]() | [:question:]() |
| [TTF_RenderText_LCD_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD_Wrapped) | [:question:]() | [:question:]() |
| [TTF_RenderGlyph_LCD](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_LCD) | [:question:]() | [:question:]() |
| [TTF_CreateSurfaceTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateSurfaceTextEngine) | [:question:]() | [:question:]() |
| [TTF_DrawSurfaceText](https://wiki.libsdl.org/SDL3_ttf/TTF_DrawSurfaceText) | [:heavy_check_mark:](./ttf/methods.go#L13) | [:heavy_check_mark:]() |
| [TTF_DestroySurfaceTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroySurfaceTextEngine) | [:heavy_check_mark:](./ttf/methods.go#L366) | [:heavy_check_mark:]() |
| [TTF_CreateRendererTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngine) | [:question:]() | [:question:]() |
| [TTF_CreateRendererTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngineWithProperties) | [:question:]() | [:question:]() |
| [TTF_DrawRendererText](https://wiki.libsdl.org/SDL3_ttf/TTF_DrawRendererText) | [:heavy_check_mark:](./ttf/methods.go#L23) | [:heavy_check_mark:]() |
| [TTF_DestroyRendererTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyRendererTextEngine) | [:heavy_check_mark:](./ttf/methods.go#L372) | [:heavy_check_mark:]() |
| [TTF_CreateGPUTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngine) | [:question:]() | [:question:]() |
| [TTF_CreateGPUTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngineWithProperties) | [:question:]() | [:question:]() |
| [TTF_GetGPUTextDrawData](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextDrawData) | [:heavy_check_mark:](./ttf/methods.go#L33) | [:heavy_check_mark:]() |
| [TTF_DestroyGPUTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyGPUTextEngine) | [:heavy_check_mark:](./ttf/methods.go#L378) | [:heavy_check_mark:]() |
| [TTF_SetGPUTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_SetGPUTextEngineWinding) | [:heavy_check_mark:](./ttf/methods.go#L384) | [:heavy_check_mark:]() |
| [TTF_GetGPUTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextEngineWinding) | [:heavy_check_mark:](./ttf/methods.go#L390) | [:heavy_check_mark:]() |
| [TTF_CreateText](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateText) | [:heavy_check_mark:](./ttf/methods.go#L396) | [:heavy_check_mark:]() |
| [TTF_GetTextProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextProperties) | [:heavy_check_mark:](./ttf/methods.go#L44) | [:heavy_check_mark:]() |
| [TTF_SetTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextEngine) | [:heavy_check_mark:](./ttf/methods.go#L50) | [:heavy_check_mark:]() |
| [TTF_GetTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextEngine) | [:heavy_check_mark:](./ttf/methods.go#L60) | [:heavy_check_mark:]() |
| [TTF_SetTextFont](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextFont) | [:heavy_check_mark:](./ttf/methods.go#L71) | [:heavy_check_mark:]() |
| [TTF_GetTextFont](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextFont) | [:heavy_check_mark:](./ttf/methods.go#L81) | [:heavy_check_mark:]() |
| [TTF_SetTextDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextDirection) | [:heavy_check_mark:](./ttf/methods.go#L92) | [:heavy_check_mark:]() |
| [TTF_GetTextDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextDirection) | [:heavy_check_mark:](./ttf/methods.go#L102) | [:heavy_check_mark:]() |
| [TTF_SetTextScript](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextScript) | [:heavy_check_mark:](./ttf/methods.go#L108) | [:heavy_check_mark:]() |
| [TTF_GetTextScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextScript) | [:heavy_check_mark:](./ttf/methods.go#L118) | [:heavy_check_mark:]() |
| [TTF_SetTextColor](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColor) | [:heavy_check_mark:](./ttf/methods.go#L124) | [:heavy_check_mark:]() |
| [TTF_SetTextColorFloat](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColorFloat) | [:heavy_check_mark:](./ttf/methods.go#L134) | [:heavy_check_mark:]() |
| [TTF_GetTextColor](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColor) | [:heavy_check_mark:](./ttf/methods.go#L144) | [:heavy_check_mark:]() |
| [TTF_GetTextColorFloat](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColorFloat) | [:heavy_check_mark:](./ttf/methods.go#L156) | [:heavy_check_mark:]() |
| [TTF_SetTextPosition](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextPosition) | [:heavy_check_mark:](./ttf/methods.go#L168) | [:heavy_check_mark:]() |
| [TTF_GetTextPosition](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextPosition) | [:heavy_check_mark:](./ttf/methods.go#L174) | [:heavy_check_mark:]() |
| [TTF_SetTextWrapWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWidth) | [:heavy_check_mark:](./ttf/methods.go#L184) | [:heavy_check_mark:]() |
| [TTF_GetTextWrapWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextWrapWidth) | [:heavy_check_mark:](./ttf/methods.go#L194) | [:heavy_check_mark:]() |
| [TTF_SetTextWrapWhitespaceVisible](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWhitespaceVisible) | [:heavy_check_mark:](./ttf/methods.go#L206) | [:heavy_check_mark:]() |
| [TTF_TextWrapWhitespaceVisible](https://wiki.libsdl.org/SDL3_ttf/TTF_TextWrapWhitespaceVisible) | [:heavy_check_mark:](./ttf/methods.go#L216) | [:heavy_check_mark:]() |
| [TTF_SetTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextString) | [:heavy_check_mark:](./ttf/methods.go#L222) | [:heavy_check_mark:]() |
| [TTF_InsertTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_InsertTextString) | [:heavy_check_mark:](./ttf/methods.go#L232) | [:heavy_check_mark:]() |
| [TTF_AppendTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_AppendTextString) | [:heavy_check_mark:](./ttf/methods.go#L242) | [:heavy_check_mark:]() |
| [TTF_DeleteTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_DeleteTextString) | [:heavy_check_mark:](./ttf/methods.go#L252) | [:heavy_check_mark:]() |
| [TTF_GetTextSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSize) | [:heavy_check_mark:](./ttf/methods.go#L262) | [:heavy_check_mark:]() |
| [TTF_GetTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubString) | [:heavy_check_mark:](./ttf/methods.go#L274) | [:heavy_check_mark:]() |
| [TTF_GetTextSubStringForLine](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForLine) | [:heavy_check_mark:](./ttf/methods.go#L286) | [:heavy_check_mark:]() |
| [TTF_GetTextSubStringsForRange](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringsForRange) | [:heavy_check_mark:](./ttf/methods.go#L298) | [:heavy_check_mark:]() |
| [TTF_GetTextSubStringForPoint](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForPoint) | [:heavy_check_mark:](./ttf/methods.go#L312) | [:heavy_check_mark:]() |
| [TTF_GetPreviousTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetPreviousTextSubString) | [:heavy_check_mark:](./ttf/methods.go#L336) | [:heavy_check_mark:]() |
| [TTF_GetNextTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetNextTextSubString) | [:question:]() | [:question:]() |
| [TTF_UpdateText](https://wiki.libsdl.org/SDL3_ttf/TTF_UpdateText) | [:heavy_check_mark:](./ttf/methods.go#L348) | [:heavy_check_mark:]() |
| [TTF_DestroyText](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyText) | [:heavy_check_mark:](./ttf/methods.go#L358) | [:heavy_check_mark:]() |
| [TTF_CloseFont](https://wiki.libsdl.org/SDL3_ttf/TTF_CloseFont) | [:heavy_check_mark:](./ttf/methods.go#L729) | [:heavy_check_mark:]() |
| [TTF_Quit](https://wiki.libsdl.org/SDL3_ttf/TTF_Quit) | [:question:]() | [:question:]() |
| [TTF_WasInit](https://wiki.libsdl.org/SDL3_ttf/TTF_WasInit) | [:question:]() | [:question:]() |

## IMG

### Image

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [IMG_Version](https://wiki.libsdl.org/SDL3_image/IMG_Version) | [:heavy_check_mark:](./img/image.go#L11) | [:heavy_check_mark:]() |
| [IMG_LoadTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTyped_IO) | [:heavy_check_mark:](./img/image.go#L17) | [:heavy_check_mark:]() |
| [IMG_Load](https://wiki.libsdl.org/SDL3_image/IMG_Load) | [:heavy_check_mark:](./img/image.go#L28) | [:heavy_check_mark:]() |
| [IMG_Load_IO](https://wiki.libsdl.org/SDL3_image/IMG_Load_IO) | [:heavy_check_mark:](./img/image.go#L39) | [:heavy_check_mark:]() |
| [IMG_LoadTexture](https://wiki.libsdl.org/SDL3_image/IMG_LoadTexture) | [:heavy_check_mark:](./img/image.go#L50) | [:heavy_check_mark:]() |
| [IMG_LoadTexture_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTexture_IO) | [:heavy_check_mark:](./img/image.go#L61) | [:heavy_check_mark:]() |
| [IMG_LoadTextureTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTextureTyped_IO) | [:heavy_check_mark:](./img/image.go#L72) | [:heavy_check_mark:]() |
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
| [IMG_LoadAVIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAVIF_IO) | [:heavy_check_mark:](./img/image.go#L83) | [:heavy_check_mark:]() |
| [IMG_LoadICO_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadICO_IO) | [:heavy_check_mark:](./img/image.go#L94) | [:heavy_check_mark:]() |
| [IMG_LoadCUR_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadCUR_IO) | [:heavy_check_mark:](./img/image.go#L105) | [:heavy_check_mark:]() |
| [IMG_LoadBMP_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadBMP_IO) | [:heavy_check_mark:](./img/image.go#L116) | [:heavy_check_mark:]() |
| [IMG_LoadGIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadGIF_IO) | [:heavy_check_mark:](./img/image.go#L127) | [:heavy_check_mark:]() |
| [IMG_LoadJPG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadJPG_IO) | [:heavy_check_mark:](./img/image.go#L138) | [:heavy_check_mark:]() |
| [IMG_LoadJXL_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadJXL_IO) | [:heavy_check_mark:](./img/image.go#L149) | [:heavy_check_mark:]() |
| [IMG_LoadLBM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadLBM_IO) | [:heavy_check_mark:](./img/image.go#L160) | [:heavy_check_mark:]() |
| [IMG_LoadPCX_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPCX_IO) | [:heavy_check_mark:](./img/image.go#L171) | [:heavy_check_mark:]() |
| [IMG_LoadPNG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPNG_IO) | [:heavy_check_mark:](./img/image.go#L182) | [:heavy_check_mark:]() |
| [IMG_LoadPNM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPNM_IO) | [:heavy_check_mark:](./img/image.go#L193) | [:heavy_check_mark:]() |
| [IMG_LoadSVG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadSVG_IO) | [:heavy_check_mark:](./img/image.go#L204) | [:heavy_check_mark:]() |
| [IMG_LoadQOI_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadQOI_IO) | [:heavy_check_mark:](./img/image.go#L215) | [:heavy_check_mark:]() |
| [IMG_LoadTGA_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTGA_IO) | [:heavy_check_mark:](./img/image.go#L226) | [:heavy_check_mark:]() |
| [IMG_LoadTIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTIF_IO) | [:heavy_check_mark:](./img/image.go#L237) | [:heavy_check_mark:]() |
| [IMG_LoadXCF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXCF_IO) | [:heavy_check_mark:](./img/image.go#L248) | [:heavy_check_mark:]() |
| [IMG_LoadXPM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXPM_IO) | [:heavy_check_mark:](./img/image.go#L259) | [:heavy_check_mark:]() |
| [IMG_LoadXV_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXV_IO) | [:heavy_check_mark:](./img/image.go#L270) | [:heavy_check_mark:]() |
| [IMG_LoadWEBP_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadWEBP_IO) | [:heavy_check_mark:](./img/image.go#L281) | [:heavy_check_mark:]() |
| [IMG_LoadSizedSVG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadSizedSVG_IO) | [:heavy_check_mark:](./img/image.go#L292) | [:heavy_check_mark:]() |
| [IMG_ReadXPMFromArray](https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArray) | [:heavy_check_mark:](./img/image.go#L303) | [:heavy_check_mark:]() |
| [IMG_ReadXPMFromArrayToRGB888](https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArrayToRGB888) | [:heavy_check_mark:](./img/image.go#L315) | [:heavy_check_mark:]() |
| [IMG_SaveAVIF](https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF) | [:heavy_check_mark:](./img/image.go#L327) | [:heavy_check_mark:]() |
| [IMG_SaveAVIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF_IO) | [:heavy_check_mark:](./img/image.go#L337) | [:heavy_check_mark:]() |
| [IMG_SavePNG](https://wiki.libsdl.org/SDL3_image/IMG_SavePNG) | [:heavy_check_mark:](./img/image.go#L347) | [:heavy_check_mark:]() |
| [IMG_SavePNG_IO](https://wiki.libsdl.org/SDL3_image/IMG_SavePNG_IO) | [:heavy_check_mark:](./img/image.go#L357) | [:heavy_check_mark:]() |
| [IMG_SaveJPG](https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG) | [:heavy_check_mark:](./img/image.go#L367) | [:heavy_check_mark:]() |
| [IMG_SaveJPG_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG_IO) | [:heavy_check_mark:](./img/image.go#L377) | [:heavy_check_mark:]() |
| [IMG_LoadAnimation](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation) | [:heavy_check_mark:](./img/image.go#L387) | [:heavy_check_mark:]() |
| [IMG_LoadAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation_IO) | [:heavy_check_mark:](./img/image.go#L398) | [:heavy_check_mark:]() |
| [IMG_LoadAnimationTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimationTyped_IO) | [:heavy_check_mark:](./img/image.go#L409) | [:heavy_check_mark:]() |
| [IMG_FreeAnimation](https://wiki.libsdl.org/SDL3_image/IMG_FreeAnimation) | [:heavy_check_mark:](./img/methods.go#L6) | [:heavy_check_mark:]() |
| [IMG_LoadGIFAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadGIFAnimation_IO) | [:heavy_check_mark:](./img/image.go#L420) | [:heavy_check_mark:]() |
| [IMG_LoadWEBPAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadWEBPAnimation_IO) | [:heavy_check_mark:](./img/image.go#L431) | [:heavy_check_mark:]() |

## MIXER

### Mixer

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [Mix_Version](https://wiki.libsdl.org/SDL3_mixer/Mix_Version) | [:question:]() | [:question:]() |
| [Mix_Init](https://wiki.libsdl.org/SDL3_mixer/Mix_Init) | [:x:](./mixer/methods.go#L33) | [:x:]() |
| [Mix_Quit](https://wiki.libsdl.org/SDL3_mixer/Mix_Quit) | [:question:]() | [:question:]() |
| [Mix_OpenAudio](https://wiki.libsdl.org/SDL3_mixer/Mix_OpenAudio) | [:question:]() | [:question:]() |
| [Mix_PauseAudio](https://wiki.libsdl.org/SDL3_mixer/Mix_PauseAudio) | [:question:]() | [:question:]() |
| [Mix_QuerySpec](https://wiki.libsdl.org/SDL3_mixer/Mix_QuerySpec) | [:question:]() | [:question:]() |
| [Mix_AllocateChannels](https://wiki.libsdl.org/SDL3_mixer/Mix_AllocateChannels) | [:question:]() | [:question:]() |
| [Mix_LoadWAV_IO](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadWAV_IO) | [:question:]() | [:question:]() |
| [Mix_LoadWAV](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadWAV) | [:question:]() | [:question:]() |
| [Mix_LoadMUS](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadMUS) | [:question:]() | [:question:]() |
| [Mix_LoadMUS_IO](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadMUS_IO) | [:question:]() | [:question:]() |
| [Mix_LoadMUSType_IO](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadMUSType_IO) | [:question:]() | [:question:]() |
| [Mix_QuickLoad_WAV](https://wiki.libsdl.org/SDL3_mixer/Mix_QuickLoad_WAV) | [:question:]() | [:question:]() |
| [Mix_QuickLoad_RAW](https://wiki.libsdl.org/SDL3_mixer/Mix_QuickLoad_RAW) | [:question:]() | [:question:]() |
| [Mix_FreeChunk](https://wiki.libsdl.org/SDL3_mixer/Mix_FreeChunk) | [:x:](./mixer/methods.go#L42) | [:x:]() |
| [Mix_FreeMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_FreeMusic) | [:x:](./mixer/methods.go#L56) | [:x:]() |
| [Mix_GetNumChunkDecoders](https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumChunkDecoders) | [:question:]() | [:question:]() |
| [Mix_GetChunkDecoder](https://wiki.libsdl.org/SDL3_mixer/Mix_GetChunkDecoder) | [:question:]() | [:question:]() |
| [Mix_HasChunkDecoder](https://wiki.libsdl.org/SDL3_mixer/Mix_HasChunkDecoder) | [:question:]() | [:question:]() |
| [Mix_GetNumMusicDecoders](https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumMusicDecoders) | [:question:]() | [:question:]() |
| [Mix_GetMusicDecoder](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicDecoder) | [:question:]() | [:question:]() |
| [Mix_HasMusicDecoder](https://wiki.libsdl.org/SDL3_mixer/Mix_HasMusicDecoder) | [:question:]() | [:question:]() |
| [Mix_GetMusicType](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicType) | [:x:](./mixer/methods.go#L63) | [:x:]() |
| [Mix_GetMusicTitle](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicTitle) | [:x:](./mixer/methods.go#L70) | [:x:]() |
| [Mix_GetMusicTitleTag](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicTitleTag) | [:x:](./mixer/methods.go#L77) | [:x:]() |
| [Mix_GetMusicArtistTag](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicArtistTag) | [:x:](./mixer/methods.go#L84) | [:x:]() |
| [Mix_GetMusicAlbumTag](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicAlbumTag) | [:x:](./mixer/methods.go#L91) | [:x:]() |
| [Mix_GetMusicCopyrightTag](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicCopyrightTag) | [:x:](./mixer/methods.go#L98) | [:x:]() |
| [Mix_SetPostMix](https://wiki.libsdl.org/SDL3_mixer/Mix_SetPostMix) | [:x:](./mixer/methods.go#L184) | [:x:]() |
| [Mix_HookMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_HookMusic) | [:x:](./mixer/methods.go#L191) | [:x:]() |
| [Mix_HookMusicFinished](https://wiki.libsdl.org/SDL3_mixer/Mix_HookMusicFinished) | [:x:](./mixer/methods.go#L6) | [:x:]() |
| [Mix_GetMusicHookData](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicHookData) | [:question:]() | [:question:]() |
| [Mix_ChannelFinished](https://wiki.libsdl.org/SDL3_mixer/Mix_ChannelFinished) | [:x:](./mixer/methods.go#L15) | [:x:]() |
| [Mix_RegisterEffect](https://wiki.libsdl.org/SDL3_mixer/Mix_RegisterEffect) | [:question:]() | [:question:]() |
| [Mix_UnregisterEffect](https://wiki.libsdl.org/SDL3_mixer/Mix_UnregisterEffect) | [:question:]() | [:question:]() |
| [Mix_UnregisterAllEffects](https://wiki.libsdl.org/SDL3_mixer/Mix_UnregisterAllEffects) | [:question:]() | [:question:]() |
| [Mix_SetPanning](https://wiki.libsdl.org/SDL3_mixer/Mix_SetPanning) | [:question:]() | [:question:]() |
| [Mix_SetPosition](https://wiki.libsdl.org/SDL3_mixer/Mix_SetPosition) | [:question:]() | [:question:]() |
| [Mix_SetDistance](https://wiki.libsdl.org/SDL3_mixer/Mix_SetDistance) | [:question:]() | [:question:]() |
| [Mix_SetReverseStereo](https://wiki.libsdl.org/SDL3_mixer/Mix_SetReverseStereo) | [:question:]() | [:question:]() |
| [Mix_ReserveChannels](https://wiki.libsdl.org/SDL3_mixer/Mix_ReserveChannels) | [:question:]() | [:question:]() |
| [Mix_GroupChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupChannel) | [:question:]() | [:question:]() |
| [Mix_GroupChannels](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupChannels) | [:question:]() | [:question:]() |
| [Mix_GroupAvailable](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupAvailable) | [:question:]() | [:question:]() |
| [Mix_GroupCount](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupCount) | [:question:]() | [:question:]() |
| [Mix_GroupOldest](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupOldest) | [:question:]() | [:question:]() |
| [Mix_GroupNewer](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupNewer) | [:question:]() | [:question:]() |
| [Mix_PlayChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_PlayChannel) | [:question:]() | [:question:]() |
| [Mix_PlayChannelTimed](https://wiki.libsdl.org/SDL3_mixer/Mix_PlayChannelTimed) | [:question:]() | [:question:]() |
| [Mix_PlayMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_PlayMusic) | [:x:](./mixer/methods.go#L105) | [:x:]() |
| [Mix_FadeInMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInMusic) | [:x:](./mixer/methods.go#L112) | [:x:]() |
| [Mix_FadeInMusicPos](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInMusicPos) | [:x:](./mixer/methods.go#L119) | [:x:]() |
| [Mix_FadeInChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInChannel) | [:question:]() | [:question:]() |
| [Mix_FadeInChannelTimed](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInChannelTimed) | [:question:]() | [:question:]() |
| [Mix_Volume](https://wiki.libsdl.org/SDL3_mixer/Mix_Volume) | [:question:]() | [:question:]() |
| [Mix_VolumeChunk](https://wiki.libsdl.org/SDL3_mixer/Mix_VolumeChunk) | [:question:]() | [:question:]() |
| [Mix_VolumeMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_VolumeMusic) | [:question:]() | [:question:]() |
| [Mix_GetMusicVolume](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicVolume) | [:x:](./mixer/methods.go#L126) | [:x:]() |
| [Mix_MasterVolume](https://wiki.libsdl.org/SDL3_mixer/Mix_MasterVolume) | [:question:]() | [:question:]() |
| [Mix_HaltChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_HaltChannel) | [:question:]() | [:question:]() |
| [Mix_HaltGroup](https://wiki.libsdl.org/SDL3_mixer/Mix_HaltGroup) | [:question:]() | [:question:]() |
| [Mix_HaltMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_HaltMusic) | [:question:]() | [:question:]() |
| [Mix_ExpireChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_ExpireChannel) | [:question:]() | [:question:]() |
| [Mix_FadeOutChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeOutChannel) | [:question:]() | [:question:]() |
| [Mix_FadeOutGroup](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeOutGroup) | [:question:]() | [:question:]() |
| [Mix_FadeOutMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeOutMusic) | [:question:]() | [:question:]() |
| [Mix_FadingMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_FadingMusic) | [:question:]() | [:question:]() |
| [Mix_FadingChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_FadingChannel) | [:question:]() | [:question:]() |
| [Mix_Pause](https://wiki.libsdl.org/SDL3_mixer/Mix_Pause) | [:question:]() | [:question:]() |
| [Mix_PauseGroup](https://wiki.libsdl.org/SDL3_mixer/Mix_PauseGroup) | [:question:]() | [:question:]() |
| [Mix_Resume](https://wiki.libsdl.org/SDL3_mixer/Mix_Resume) | [:question:]() | [:question:]() |
| [Mix_ResumeGroup](https://wiki.libsdl.org/SDL3_mixer/Mix_ResumeGroup) | [:question:]() | [:question:]() |
| [Mix_Paused](https://wiki.libsdl.org/SDL3_mixer/Mix_Paused) | [:question:]() | [:question:]() |
| [Mix_PauseMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_PauseMusic) | [:question:]() | [:question:]() |
| [Mix_ResumeMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_ResumeMusic) | [:question:]() | [:question:]() |
| [Mix_RewindMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_RewindMusic) | [:question:]() | [:question:]() |
| [Mix_PausedMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_PausedMusic) | [:question:]() | [:question:]() |
| [Mix_ModMusicJumpToOrder](https://wiki.libsdl.org/SDL3_mixer/Mix_ModMusicJumpToOrder) | [:question:]() | [:question:]() |
| [Mix_StartTrack](https://wiki.libsdl.org/SDL3_mixer/Mix_StartTrack) | [:x:](./mixer/methods.go#L133) | [:x:]() |
| [Mix_GetNumTracks](https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumTracks) | [:x:](./mixer/methods.go#L140) | [:x:]() |
| [Mix_SetMusicPosition](https://wiki.libsdl.org/SDL3_mixer/Mix_SetMusicPosition) | [:question:]() | [:question:]() |
| [Mix_GetMusicPosition](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicPosition) | [:x:](./mixer/methods.go#L147) | [:x:]() |
| [Mix_MusicDuration](https://wiki.libsdl.org/SDL3_mixer/Mix_MusicDuration) | [:x:](./mixer/methods.go#L154) | [:x:]() |
| [Mix_GetMusicLoopStartTime](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopStartTime) | [:x:](./mixer/methods.go#L161) | [:x:]() |
| [Mix_GetMusicLoopEndTime](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopEndTime) | [:x:](./mixer/methods.go#L168) | [:x:]() |
| [Mix_GetMusicLoopLengthTime](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopLengthTime) | [:x:](./mixer/methods.go#L175) | [:x:]() |
| [Mix_Playing](https://wiki.libsdl.org/SDL3_mixer/Mix_Playing) | [:question:]() | [:question:]() |
| [Mix_PlayingMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_PlayingMusic) | [:question:]() | [:question:]() |
| [Mix_SetSoundFonts](https://wiki.libsdl.org/SDL3_mixer/Mix_SetSoundFonts) | [:question:]() | [:question:]() |
| [Mix_GetSoundFonts](https://wiki.libsdl.org/SDL3_mixer/Mix_GetSoundFonts) | [:question:]() | [:question:]() |
| [Mix_EachSoundFont](https://wiki.libsdl.org/SDL3_mixer/Mix_EachSoundFont) | [:x:](./mixer/methods.go#L24) | [:x:]() |
| [Mix_SetTimidityCfg](https://wiki.libsdl.org/SDL3_mixer/Mix_SetTimidityCfg) | [:question:]() | [:question:]() |
| [Mix_GetTimidityCfg](https://wiki.libsdl.org/SDL3_mixer/Mix_GetTimidityCfg) | [:question:]() | [:question:]() |
| [Mix_GetChunk](https://wiki.libsdl.org/SDL3_mixer/Mix_GetChunk) | [:question:]() | [:question:]() |
| [Mix_CloseAudio](https://wiki.libsdl.org/SDL3_mixer/Mix_CloseAudio) | [:question:]() | [:question:]() |

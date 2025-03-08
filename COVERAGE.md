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
| [SDL_SetAppMetadata](https://wiki.libsdl.org/SDL3/SDL_SetAppMetadata) | [:heavy_check_mark:](./functions.go#L54) | [:x:](./sdl_functions_js.go#L13989) |
| [SDL_SetAppMetadataProperty](https://wiki.libsdl.org/SDL3/SDL_SetAppMetadataProperty) | [:heavy_check_mark:](./functions.go#L64) | [:x:](./sdl_functions_js.go#L14006) |
| [SDL_GetAppMetadataProperty](https://wiki.libsdl.org/SDL3/SDL_GetAppMetadataProperty) | [:heavy_check_mark:](./functions.go#L74) | [:x:](./sdl_functions_js.go#L14021) |
### Hints

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetHintWithPriority](https://wiki.libsdl.org/SDL3/SDL_SetHintWithPriority) | [:heavy_check_mark:](./functions.go#L82) | [:x:](./sdl_functions_js.go#L13796) |
| [SDL_SetHint](https://wiki.libsdl.org/SDL3/SDL_SetHint) | [:heavy_check_mark:](./functions.go#L92) | [:x:](./sdl_functions_js.go#L13813) |
| [SDL_ResetHint](https://wiki.libsdl.org/SDL3/SDL_ResetHint) | [:heavy_check_mark:](./functions.go#L102) | [:x:](./sdl_functions_js.go#L13828) |
| [SDL_ResetHints](https://wiki.libsdl.org/SDL3/SDL_ResetHints) | [:heavy_check_mark:](./functions.go#L112) | [:x:](./sdl_functions_js.go#L13841) |
| [SDL_GetHint](https://wiki.libsdl.org/SDL3/SDL_GetHint) | [:heavy_check_mark:](./functions.go#L118) | [:x:](./sdl_functions_js.go#L13850) |
| [SDL_GetHintBoolean](https://wiki.libsdl.org/SDL3/SDL_GetHintBoolean) | [:heavy_check_mark:](./functions.go#L124) | [:x:](./sdl_functions_js.go#L13863) |
| [SDL_AddHintCallback](https://wiki.libsdl.org/SDL3/SDL_AddHintCallback) | [:question:]() | [:question:](./sdl_functions_js.go#L13878) |
| [SDL_RemoveHintCallback](https://wiki.libsdl.org/SDL3/SDL_RemoveHintCallback) | [:question:]() | [:question:](./sdl_functions_js.go#L13895) |
### Error

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetError](https://wiki.libsdl.org/SDL3/SDL_SetError) | [:question:]() | [:question:](./sdl_functions_js.go#L488) |
| [SDL_SetErrorV](https://wiki.libsdl.org/SDL3/SDL_SetErrorV) | [:question:]() | [:question:](./sdl_functions_js.go#L501) |
| [SDL_OutOfMemory](https://wiki.libsdl.org/SDL3/SDL_OutOfMemory) | [:heavy_check_mark:](./functions.go#L135) | [:x:](./sdl_functions_js.go#L516) |
| [SDL_GetError](https://wiki.libsdl.org/SDL3/SDL_GetError) | [:heavy_check_mark:](./init_notjs.go#L32) | [:heavy_check_mark:](./sdl_functions_js.go#L527) |
| [SDL_ClearError](https://wiki.libsdl.org/SDL3/SDL_ClearError) | [:question:]() | [:question:](./sdl_functions_js.go#L535) |
### Properties

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetGlobalProperties](https://wiki.libsdl.org/SDL3/SDL_GetGlobalProperties) | [:heavy_check_mark:](./functions.go#L143) | [:x:](./sdl_functions_js.go#L546) |
| [SDL_CreateProperties](https://wiki.libsdl.org/SDL3/SDL_CreateProperties) | [:heavy_check_mark:](./functions.go#L154) | [:x:](./sdl_functions_js.go#L557) |
| [SDL_CopyProperties](https://wiki.libsdl.org/SDL3/SDL_CopyProperties) | [:heavy_check_mark:](./functions.go#L165) | [:x:](./sdl_functions_js.go#L568) |
| [SDL_LockProperties](https://wiki.libsdl.org/SDL3/SDL_LockProperties) | [:heavy_check_mark:](./methods.go#L5623) | [:x:](./sdl_functions_js.go#L583) |
| [SDL_UnlockProperties](https://wiki.libsdl.org/SDL3/SDL_UnlockProperties) | [:heavy_check_mark:](./methods.go#L5633) | [:x:](./sdl_functions_js.go#L596) |
| [SDL_SetPointerPropertyWithCleanup](https://wiki.libsdl.org/SDL3/SDL_SetPointerPropertyWithCleanup) | [:x:](./methods.go#L5639) | [:x:](./sdl_functions_js.go#L607) |
| [SDL_SetPointerProperty](https://wiki.libsdl.org/SDL3/SDL_SetPointerProperty) | [:x:](./methods.go#L5646) | [:x:](./sdl_functions_js.go#L628) |
| [SDL_SetStringProperty](https://wiki.libsdl.org/SDL3/SDL_SetStringProperty) | [:heavy_check_mark:](./methods.go#L5653) | [:x:](./sdl_functions_js.go#L645) |
| [SDL_SetNumberProperty](https://wiki.libsdl.org/SDL3/SDL_SetNumberProperty) | [:heavy_check_mark:](./methods.go#L5663) | [:x:](./sdl_functions_js.go#L662) |
| [SDL_SetFloatProperty](https://wiki.libsdl.org/SDL3/SDL_SetFloatProperty) | [:heavy_check_mark:](./methods.go#L5673) | [:x:](./sdl_functions_js.go#L679) |
| [SDL_SetBooleanProperty](https://wiki.libsdl.org/SDL3/SDL_SetBooleanProperty) | [:heavy_check_mark:](./methods.go#L5683) | [:x:](./sdl_functions_js.go#L696) |
| [SDL_HasProperty](https://wiki.libsdl.org/SDL3/SDL_HasProperty) | [:heavy_check_mark:](./methods.go#L5693) | [:x:](./sdl_functions_js.go#L713) |
| [SDL_GetPropertyType](https://wiki.libsdl.org/SDL3/SDL_GetPropertyType) | [:heavy_check_mark:](./methods.go#L5699) | [:x:](./sdl_functions_js.go#L728) |
| [SDL_GetPointerProperty](https://wiki.libsdl.org/SDL3/SDL_GetPointerProperty) | [:x:](./methods.go#L5705) | [:x:](./sdl_functions_js.go#L743) |
| [SDL_GetStringProperty](https://wiki.libsdl.org/SDL3/SDL_GetStringProperty) | [:heavy_check_mark:](./methods.go#L5712) | [:x:](./sdl_functions_js.go#L760) |
| [SDL_GetNumberProperty](https://wiki.libsdl.org/SDL3/SDL_GetNumberProperty) | [:heavy_check_mark:](./methods.go#L5718) | [:x:](./sdl_functions_js.go#L777) |
| [SDL_GetFloatProperty](https://wiki.libsdl.org/SDL3/SDL_GetFloatProperty) | [:heavy_check_mark:](./methods.go#L5724) | [:x:](./sdl_functions_js.go#L794) |
| [SDL_GetBooleanProperty](https://wiki.libsdl.org/SDL3/SDL_GetBooleanProperty) | [:heavy_check_mark:](./methods.go#L5730) | [:x:](./sdl_functions_js.go#L811) |
| [SDL_ClearProperty](https://wiki.libsdl.org/SDL3/SDL_ClearProperty) | [:heavy_check_mark:](./methods.go#L5736) | [:x:](./sdl_functions_js.go#L828) |
| [SDL_EnumerateProperties](https://wiki.libsdl.org/SDL3/SDL_EnumerateProperties) | [:x:](./methods.go#L5746) | [:x:](./sdl_functions_js.go#L843) |
| [SDL_DestroyProperties](https://wiki.libsdl.org/SDL3/SDL_DestroyProperties) | [:heavy_check_mark:](./methods.go#L5753) | [:x:](./sdl_functions_js.go#L860) |
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
| [SDL_SetLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_SetLogOutputFunction) | [:x:](./methods.go#L3904) | [:x:](./sdl_functions_js.go#L14323) |
### Video

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumVideoDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumVideoDrivers) | [:heavy_check_mark:](./functions.go#L486) | [:x:](./sdl_functions_js.go#L5828) |
| [SDL_GetVideoDriver](https://wiki.libsdl.org/SDL3/SDL_GetVideoDriver) | [:heavy_check_mark:](./functions.go#L492) | [:x:](./sdl_functions_js.go#L5839) |
| [SDL_GetCurrentVideoDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentVideoDriver) | [:heavy_check_mark:](./functions.go#L498) | [:x:](./sdl_functions_js.go#L5852) |
| [SDL_GetSystemTheme](https://wiki.libsdl.org/SDL3/SDL_GetSystemTheme) | [:heavy_check_mark:](./functions.go#L504) | [:x:](./sdl_functions_js.go#L5863) |
| [SDL_GetDisplays](https://wiki.libsdl.org/SDL3/SDL_GetDisplays) | [:heavy_check_mark:](./functions.go#L510) | [:x:](./sdl_functions_js.go#L5874) |
| [SDL_GetPrimaryDisplay](https://wiki.libsdl.org/SDL3/SDL_GetPrimaryDisplay) | [:heavy_check_mark:](./functions.go#L524) | [:x:](./sdl_functions_js.go#L5890) |
| [SDL_GetDisplayProperties](https://wiki.libsdl.org/SDL3/SDL_GetDisplayProperties) | [:heavy_check_mark:](./methods.go#L3456) | [:x:](./sdl_functions_js.go#L5901) |
| [SDL_GetDisplayName](https://wiki.libsdl.org/SDL3/SDL_GetDisplayName) | [:heavy_check_mark:](./methods.go#L3467) | [:x:](./sdl_functions_js.go#L5914) |
| [SDL_GetDisplayBounds](https://wiki.libsdl.org/SDL3/SDL_GetDisplayBounds) | [:heavy_check_mark:](./methods.go#L3478) | [:x:](./sdl_functions_js.go#L5927) |
| [SDL_GetDisplayUsableBounds](https://wiki.libsdl.org/SDL3/SDL_GetDisplayUsableBounds) | [:heavy_check_mark:](./methods.go#L3490) | [:x:](./sdl_functions_js.go#L5945) |
| [SDL_GetNaturalDisplayOrientation](https://wiki.libsdl.org/SDL3/SDL_GetNaturalDisplayOrientation) | [:heavy_check_mark:](./methods.go#L3502) | [:x:](./sdl_functions_js.go#L5963) |
| [SDL_GetCurrentDisplayOrientation](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayOrientation) | [:heavy_check_mark:](./methods.go#L3508) | [:x:](./sdl_functions_js.go#L5976) |
| [SDL_GetDisplayContentScale](https://wiki.libsdl.org/SDL3/SDL_GetDisplayContentScale) | [:heavy_check_mark:](./methods.go#L3514) | [:x:](./sdl_functions_js.go#L5989) |
| [SDL_GetFullscreenDisplayModes](https://wiki.libsdl.org/SDL3/SDL_GetFullscreenDisplayModes) | [:heavy_check_mark:](./methods.go#L3525) | [:x:](./sdl_functions_js.go#L6002) |
| [SDL_GetClosestFullscreenDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetClosestFullscreenDisplayMode) | [:heavy_check_mark:](./methods.go#L3539) | [:x:](./sdl_functions_js.go#L6020) |
| [SDL_GetDesktopDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetDesktopDisplayMode) | [:heavy_check_mark:](./methods.go#L3551) | [:x:](./sdl_functions_js.go#L6046) |
| [SDL_GetCurrentDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayMode) | [:heavy_check_mark:](./methods.go#L3562) | [:x:](./sdl_functions_js.go#L6062) |
| [SDL_GetDisplayForPoint](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForPoint) | [:heavy_check_mark:](./functions.go#L530) | [:x:](./sdl_functions_js.go#L6078) |
| [SDL_GetDisplayForRect](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForRect) | [:heavy_check_mark:](./functions.go#L536) | [:x:](./sdl_functions_js.go#L6094) |
| [SDL_GetDisplayForWindow](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForWindow) | [:heavy_check_mark:](./functions.go#L542) | [:x:](./sdl_functions_js.go#L6110) |
| [SDL_GetWindowPixelDensity](https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelDensity) | [:heavy_check_mark:](./methods.go#L3950) | [:x:](./sdl_functions_js.go#L6126) |
| [SDL_GetWindowDisplayScale](https://wiki.libsdl.org/SDL3/SDL_GetWindowDisplayScale) | [:heavy_check_mark:](./methods.go#L3961) | [:x:](./sdl_functions_js.go#L6142) |
| [SDL_SetWindowFullscreenMode](https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreenMode) | [:heavy_check_mark:](./methods.go#L3972) | [:x:](./sdl_functions_js.go#L6158) |
| [SDL_GetWindowFullscreenMode](https://wiki.libsdl.org/SDL3/SDL_GetWindowFullscreenMode) | [:heavy_check_mark:](./methods.go#L3982) | [:x:](./sdl_functions_js.go#L6179) |
| [SDL_GetWindowICCProfile](https://wiki.libsdl.org/SDL3/SDL_GetWindowICCProfile) | [:heavy_check_mark:](./methods.go#L3988) | [:x:](./sdl_functions_js.go#L6198) |
| [SDL_GetWindowPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelFormat) | [:heavy_check_mark:](./methods.go#L4002) | [:x:](./sdl_functions_js.go#L6219) |
| [SDL_GetWindows](https://wiki.libsdl.org/SDL3/SDL_GetWindows) | [:heavy_check_mark:](./functions.go#L548) | [:x:](./sdl_functions_js.go#L6235) |
| [SDL_CreateWindow](https://wiki.libsdl.org/SDL3/SDL_CreateWindow) | [:heavy_check_mark:](./functions.go#L562) | [:x:](./sdl_functions_js.go#L6251) |
| [SDL_CreatePopupWindow](https://wiki.libsdl.org/SDL3/SDL_CreatePopupWindow) | [:heavy_check_mark:](./methods.go#L4013) | [:x:](./sdl_functions_js.go#L6273) |
| [SDL_CreateWindowWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateWindowWithProperties) | [:heavy_check_mark:](./functions.go#L573) | [:x:](./sdl_functions_js.go#L6302) |
| [SDL_GetWindowID](https://wiki.libsdl.org/SDL3/SDL_GetWindowID) | [:heavy_check_mark:](./methods.go#L4024) | [:x:](./sdl_functions_js.go#L6318) |
| [SDL_GetWindowFromID](https://wiki.libsdl.org/SDL3/SDL_GetWindowFromID) | [:heavy_check_mark:](./methods.go#L50) | [:x:](./sdl_functions_js.go#L6334) |
| [SDL_GetWindowParent](https://wiki.libsdl.org/SDL3/SDL_GetWindowParent) | [:heavy_check_mark:](./methods.go#L4035) | [:x:](./sdl_functions_js.go#L6350) |
| [SDL_GetWindowProperties](https://wiki.libsdl.org/SDL3/SDL_GetWindowProperties) | [:heavy_check_mark:](./methods.go#L4041) | [:x:](./sdl_functions_js.go#L6369) |
| [SDL_GetWindowFlags](https://wiki.libsdl.org/SDL3/SDL_GetWindowFlags) | [:heavy_check_mark:](./methods.go#L4052) | [:x:](./sdl_functions_js.go#L6385) |
| [SDL_SetWindowTitle](https://wiki.libsdl.org/SDL3/SDL_SetWindowTitle) | [:heavy_check_mark:](./methods.go#L4058) | [:x:](./sdl_functions_js.go#L6401) |
| [SDL_GetWindowTitle](https://wiki.libsdl.org/SDL3/SDL_GetWindowTitle) | [:heavy_check_mark:](./methods.go#L4068) | [:x:](./sdl_functions_js.go#L6419) |
| [SDL_SetWindowIcon](https://wiki.libsdl.org/SDL3/SDL_SetWindowIcon) | [:heavy_check_mark:](./methods.go#L4074) | [:x:](./sdl_functions_js.go#L6435) |
| [SDL_SetWindowPosition](https://wiki.libsdl.org/SDL3/SDL_SetWindowPosition) | [:heavy_check_mark:](./methods.go#L4084) | [:x:](./sdl_functions_js.go#L6456) |
| [SDL_GetWindowPosition](https://wiki.libsdl.org/SDL3/SDL_GetWindowPosition) | [:heavy_check_mark:](./methods.go#L4094) | [:x:](./sdl_functions_js.go#L6476) |
| [SDL_SetWindowSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowSize) | [:heavy_check_mark:](./methods.go#L4106) | [:x:](./sdl_functions_js.go#L6502) |
| [SDL_GetWindowSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowSize) | [:heavy_check_mark:](./methods.go#L4116) | [:heavy_check_mark:](./sdl_functions_js.go#L6522) |
| [SDL_GetWindowSafeArea](https://wiki.libsdl.org/SDL3/SDL_GetWindowSafeArea) | [:heavy_check_mark:](./methods.go#L4128) | [:x:](./sdl_functions_js.go#L6544) |
| [SDL_SetWindowAspectRatio](https://wiki.libsdl.org/SDL3/SDL_SetWindowAspectRatio) | [:heavy_check_mark:](./methods.go#L4140) | [:x:](./sdl_functions_js.go#L6565) |
| [SDL_GetWindowAspectRatio](https://wiki.libsdl.org/SDL3/SDL_GetWindowAspectRatio) | [:heavy_check_mark:](./methods.go#L4150) | [:x:](./sdl_functions_js.go#L6585) |
| [SDL_GetWindowBordersSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowBordersSize) | [:heavy_check_mark:](./methods.go#L4162) | [:x:](./sdl_functions_js.go#L6611) |
| [SDL_GetWindowSizeInPixels](https://wiki.libsdl.org/SDL3/SDL_GetWindowSizeInPixels) | [:heavy_check_mark:](./methods.go#L4174) | [:x:](./sdl_functions_js.go#L6647) |
| [SDL_SetWindowMinimumSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowMinimumSize) | [:heavy_check_mark:](./methods.go#L4186) | [:x:](./sdl_functions_js.go#L6673) |
| [SDL_GetWindowMinimumSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowMinimumSize) | [:heavy_check_mark:](./methods.go#L4196) | [:x:](./sdl_functions_js.go#L6693) |
| [SDL_SetWindowMaximumSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowMaximumSize) | [:heavy_check_mark:](./methods.go#L4208) | [:x:](./sdl_functions_js.go#L6719) |
| [SDL_GetWindowMaximumSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowMaximumSize) | [:heavy_check_mark:](./methods.go#L4218) | [:x:](./sdl_functions_js.go#L6739) |
| [SDL_SetWindowBordered](https://wiki.libsdl.org/SDL3/SDL_SetWindowBordered) | [:heavy_check_mark:](./methods.go#L4230) | [:x:](./sdl_functions_js.go#L6765) |
| [SDL_SetWindowResizable](https://wiki.libsdl.org/SDL3/SDL_SetWindowResizable) | [:heavy_check_mark:](./methods.go#L4240) | [:x:](./sdl_functions_js.go#L6783) |
| [SDL_SetWindowAlwaysOnTop](https://wiki.libsdl.org/SDL3/SDL_SetWindowAlwaysOnTop) | [:heavy_check_mark:](./methods.go#L4250) | [:x:](./sdl_functions_js.go#L6801) |
| [SDL_ShowWindow](https://wiki.libsdl.org/SDL3/SDL_ShowWindow) | [:heavy_check_mark:](./methods.go#L4260) | [:x:](./sdl_functions_js.go#L6819) |
| [SDL_HideWindow](https://wiki.libsdl.org/SDL3/SDL_HideWindow) | [:heavy_check_mark:](./methods.go#L4270) | [:x:](./sdl_functions_js.go#L6835) |
| [SDL_RaiseWindow](https://wiki.libsdl.org/SDL3/SDL_RaiseWindow) | [:heavy_check_mark:](./methods.go#L4280) | [:x:](./sdl_functions_js.go#L6851) |
| [SDL_MaximizeWindow](https://wiki.libsdl.org/SDL3/SDL_MaximizeWindow) | [:heavy_check_mark:](./methods.go#L4290) | [:x:](./sdl_functions_js.go#L6867) |
| [SDL_MinimizeWindow](https://wiki.libsdl.org/SDL3/SDL_MinimizeWindow) | [:heavy_check_mark:](./methods.go#L4300) | [:x:](./sdl_functions_js.go#L6883) |
| [SDL_RestoreWindow](https://wiki.libsdl.org/SDL3/SDL_RestoreWindow) | [:heavy_check_mark:](./methods.go#L4310) | [:x:](./sdl_functions_js.go#L6899) |
| [SDL_SetWindowFullscreen](https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreen) | [:heavy_check_mark:](./methods.go#L4320) | [:x:](./sdl_functions_js.go#L6915) |
| [SDL_SyncWindow](https://wiki.libsdl.org/SDL3/SDL_SyncWindow) | [:heavy_check_mark:](./methods.go#L4330) | [:x:](./sdl_functions_js.go#L6933) |
| [SDL_WindowHasSurface](https://wiki.libsdl.org/SDL3/SDL_WindowHasSurface) | [:heavy_check_mark:](./methods.go#L4340) | [:x:](./sdl_functions_js.go#L6949) |
| [SDL_GetWindowSurface](https://wiki.libsdl.org/SDL3/SDL_GetWindowSurface) | [:heavy_check_mark:](./methods.go#L4346) | [:x:](./sdl_functions_js.go#L6965) |
| [SDL_SetWindowSurfaceVSync](https://wiki.libsdl.org/SDL3/SDL_SetWindowSurfaceVSync) | [:heavy_check_mark:](./methods.go#L4357) | [:x:](./sdl_functions_js.go#L6984) |
| [SDL_GetWindowSurfaceVSync](https://wiki.libsdl.org/SDL3/SDL_GetWindowSurfaceVSync) | [:heavy_check_mark:](./methods.go#L4367) | [:x:](./sdl_functions_js.go#L7002) |
| [SDL_UpdateWindowSurface](https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurface) | [:heavy_check_mark:](./methods.go#L4379) | [:x:](./sdl_functions_js.go#L7023) |
| [SDL_UpdateWindowSurfaceRects](https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurfaceRects) | [:heavy_check_mark:](./methods.go#L4389) | [:x:](./sdl_functions_js.go#L7039) |
| [SDL_DestroyWindowSurface](https://wiki.libsdl.org/SDL3/SDL_DestroyWindowSurface) | [:heavy_check_mark:](./methods.go#L4399) | [:x:](./sdl_functions_js.go#L7062) |
| [SDL_SetWindowKeyboardGrab](https://wiki.libsdl.org/SDL3/SDL_SetWindowKeyboardGrab) | [:heavy_check_mark:](./methods.go#L4409) | [:x:](./sdl_functions_js.go#L7078) |
| [SDL_SetWindowMouseGrab](https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseGrab) | [:heavy_check_mark:](./methods.go#L4419) | [:x:](./sdl_functions_js.go#L7096) |
| [SDL_GetWindowKeyboardGrab](https://wiki.libsdl.org/SDL3/SDL_GetWindowKeyboardGrab) | [:heavy_check_mark:](./methods.go#L4429) | [:x:](./sdl_functions_js.go#L7114) |
| [SDL_GetWindowMouseGrab](https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseGrab) | [:heavy_check_mark:](./methods.go#L4435) | [:x:](./sdl_functions_js.go#L7130) |
| [SDL_GetGrabbedWindow](https://wiki.libsdl.org/SDL3/SDL_GetGrabbedWindow) | [:heavy_check_mark:](./functions.go#L584) | [:x:](./sdl_functions_js.go#L7146) |
| [SDL_SetWindowMouseRect](https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseRect) | [:heavy_check_mark:](./methods.go#L4441) | [:x:](./sdl_functions_js.go#L7160) |
| [SDL_GetWindowMouseRect](https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseRect) | [:heavy_check_mark:](./methods.go#L4451) | [:x:](./sdl_functions_js.go#L7181) |
| [SDL_SetWindowOpacity](https://wiki.libsdl.org/SDL3/SDL_SetWindowOpacity) | [:heavy_check_mark:](./methods.go#L4457) | [:x:](./sdl_functions_js.go#L7200) |
| [SDL_GetWindowOpacity](https://wiki.libsdl.org/SDL3/SDL_GetWindowOpacity) | [:heavy_check_mark:](./methods.go#L4467) | [:x:](./sdl_functions_js.go#L7218) |
| [SDL_SetWindowParent](https://wiki.libsdl.org/SDL3/SDL_SetWindowParent) | [:heavy_check_mark:](./methods.go#L4473) | [:x:](./sdl_functions_js.go#L7234) |
| [SDL_SetWindowModal](https://wiki.libsdl.org/SDL3/SDL_SetWindowModal) | [:heavy_check_mark:](./methods.go#L4483) | [:x:](./sdl_functions_js.go#L7255) |
| [SDL_SetWindowFocusable](https://wiki.libsdl.org/SDL3/SDL_SetWindowFocusable) | [:heavy_check_mark:](./methods.go#L4493) | [:x:](./sdl_functions_js.go#L7273) |
| [SDL_ShowWindowSystemMenu](https://wiki.libsdl.org/SDL3/SDL_ShowWindowSystemMenu) | [:heavy_check_mark:](./methods.go#L4503) | [:x:](./sdl_functions_js.go#L7291) |
| [SDL_SetWindowHitTest](https://wiki.libsdl.org/SDL3/SDL_SetWindowHitTest) | [:x:](./methods.go#L4513) | [:x:](./sdl_functions_js.go#L7311) |
| [SDL_SetWindowShape](https://wiki.libsdl.org/SDL3/SDL_SetWindowShape) | [:heavy_check_mark:](./methods.go#L4520) | [:x:](./sdl_functions_js.go#L7331) |
| [SDL_FlashWindow](https://wiki.libsdl.org/SDL3/SDL_FlashWindow) | [:heavy_check_mark:](./methods.go#L4530) | [:x:](./sdl_functions_js.go#L7352) |
| [SDL_DestroyWindow](https://wiki.libsdl.org/SDL3/SDL_DestroyWindow) | [:heavy_check_mark:](./methods.go#L4540) | [:heavy_check_mark:](./sdl_functions_js.go#L7370) |
| [SDL_ScreenSaverEnabled](https://wiki.libsdl.org/SDL3/SDL_ScreenSaverEnabled) | [:heavy_check_mark:](./functions.go#L590) | [:x:](./sdl_functions_js.go#L7382) |
| [SDL_EnableScreenSaver](https://wiki.libsdl.org/SDL3/SDL_EnableScreenSaver) | [:heavy_check_mark:](./functions.go#L596) | [:x:](./sdl_functions_js.go#L7393) |
| [SDL_DisableScreenSaver](https://wiki.libsdl.org/SDL3/SDL_DisableScreenSaver) | [:heavy_check_mark:](./functions.go#L606) | [:x:](./sdl_functions_js.go#L7404) |
| [SDL_GL_LoadLibrary](https://wiki.libsdl.org/SDL3/SDL_GL_LoadLibrary) | [:question:]() | [:question:](./sdl_functions_js.go#L7415) |
| [SDL_GL_GetProcAddress](https://wiki.libsdl.org/SDL3/SDL_GL_GetProcAddress) | [:question:]() | [:question:](./sdl_functions_js.go#L7428) |
| [SDL_EGL_GetProcAddress](https://wiki.libsdl.org/SDL3/SDL_EGL_GetProcAddress) | [:question:]() | [:question:](./sdl_functions_js.go#L7441) |
| [SDL_GL_UnloadLibrary](https://wiki.libsdl.org/SDL3/SDL_GL_UnloadLibrary) | [:question:]() | [:question:](./sdl_functions_js.go#L7454) |
| [SDL_GL_ExtensionSupported](https://wiki.libsdl.org/SDL3/SDL_GL_ExtensionSupported) | [:heavy_check_mark:](./functions.go#L616) | [:x:](./sdl_functions_js.go#L7463) |
| [SDL_GL_ResetAttributes](https://wiki.libsdl.org/SDL3/SDL_GL_ResetAttributes) | [:heavy_check_mark:](./functions.go#L622) | [:x:](./sdl_functions_js.go#L7476) |
| [SDL_GL_SetAttribute](https://wiki.libsdl.org/SDL3/SDL_GL_SetAttribute) | [:heavy_check_mark:](./functions.go#L628) | [:x:](./sdl_functions_js.go#L7485) |
| [SDL_GL_GetAttribute](https://wiki.libsdl.org/SDL3/SDL_GL_GetAttribute) | [:heavy_check_mark:](./functions.go#L638) | [:x:](./sdl_functions_js.go#L7500) |
| [SDL_GL_CreateContext](https://wiki.libsdl.org/SDL3/SDL_GL_CreateContext) | [:x:](./methods.go#L4546) | [:x:](./sdl_functions_js.go#L7518) |
| [SDL_GL_MakeCurrent](https://wiki.libsdl.org/SDL3/SDL_GL_MakeCurrent) | [:x:](./methods.go#L4553) | [:x:](./sdl_functions_js.go#L7534) |
| [SDL_GL_GetCurrentWindow](https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentWindow) | [:heavy_check_mark:](./functions.go#L650) | [:x:](./sdl_functions_js.go#L7552) |
| [SDL_GL_GetCurrentContext](https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentContext) | [:heavy_check_mark:](./functions.go#L661) | [:x:](./sdl_functions_js.go#L7566) |
| [SDL_EGL_GetCurrentDisplay](https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentDisplay) | [:heavy_check_mark:](./functions.go#L672) | [:x:](./sdl_functions_js.go#L7577) |
| [SDL_EGL_GetCurrentConfig](https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentConfig) | [:heavy_check_mark:](./functions.go#L683) | [:x:](./sdl_functions_js.go#L7588) |
| [SDL_EGL_GetWindowSurface](https://wiki.libsdl.org/SDL3/SDL_EGL_GetWindowSurface) | [:x:](./methods.go#L4560) | [:x:](./sdl_functions_js.go#L7599) |
| [SDL_EGL_SetAttributeCallbacks](https://wiki.libsdl.org/SDL3/SDL_EGL_SetAttributeCallbacks) | [:question:]() | [:question:](./sdl_functions_js.go#L7615) |
| [SDL_GL_SetSwapInterval](https://wiki.libsdl.org/SDL3/SDL_GL_SetSwapInterval) | [:heavy_check_mark:](./functions.go#L694) | [:x:](./sdl_functions_js.go#L7632) |
| [SDL_GL_GetSwapInterval](https://wiki.libsdl.org/SDL3/SDL_GL_GetSwapInterval) | [:heavy_check_mark:](./functions.go#L704) | [:x:](./sdl_functions_js.go#L7645) |
| [SDL_GL_SwapWindow](https://wiki.libsdl.org/SDL3/SDL_GL_SwapWindow) | [:x:](./methods.go#L4567) | [:x:](./sdl_functions_js.go#L7661) |
| [SDL_GL_DestroyContext](https://wiki.libsdl.org/SDL3/SDL_GL_DestroyContext) | [:question:]() | [:question:](./sdl_functions_js.go#L7677) |
### Events

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_PumpEvents](https://wiki.libsdl.org/SDL3/SDL_PumpEvents) | [:heavy_check_mark:](./functions.go#L181) | [:x:](./sdl_functions_js.go#L10940) |
| [SDL_PeepEvents](https://wiki.libsdl.org/SDL3/SDL_PeepEvents) | [:question:]() | [:question:](./sdl_functions_js.go#L10949) |
| [SDL_HasEvent](https://wiki.libsdl.org/SDL3/SDL_HasEvent) | [:heavy_check_mark:](./functions.go#L189) | [:x:](./sdl_functions_js.go#L10973) |
| [SDL_HasEvents](https://wiki.libsdl.org/SDL3/SDL_HasEvents) | [:heavy_check_mark:](./functions.go#L195) | [:x:](./sdl_functions_js.go#L10986) |
| [SDL_FlushEvent](https://wiki.libsdl.org/SDL3/SDL_FlushEvent) | [:heavy_check_mark:](./functions.go#L201) | [:x:](./sdl_functions_js.go#L11001) |
| [SDL_FlushEvents](https://wiki.libsdl.org/SDL3/SDL_FlushEvents) | [:heavy_check_mark:](./functions.go#L207) | [:x:](./sdl_functions_js.go#L11012) |
| [SDL_PollEvent](https://wiki.libsdl.org/SDL3/SDL_PollEvent) | [:heavy_check_mark:](./functions.go#L213) | [:heavy_check_mark:](./sdl_functions_js.go#L11025) |
| [SDL_WaitEvent](https://wiki.libsdl.org/SDL3/SDL_WaitEvent) | [:heavy_check_mark:](./functions.go#L219) | [:x:](./sdl_functions_js.go#L11038) |
| [SDL_WaitEventTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitEventTimeout) | [:heavy_check_mark:](./functions.go#L229) | [:x:](./sdl_functions_js.go#L11054) |
| [SDL_PushEvent](https://wiki.libsdl.org/SDL3/SDL_PushEvent) | [:heavy_check_mark:](./functions.go#L235) | [:x:](./sdl_functions_js.go#L11072) |
| [SDL_SetEventFilter](https://wiki.libsdl.org/SDL3/SDL_SetEventFilter) | [:question:]() | [:question:](./sdl_functions_js.go#L11088) |
| [SDL_GetEventFilter](https://wiki.libsdl.org/SDL3/SDL_GetEventFilter) | [:x:](./methods.go#L3601) | [:x:](./sdl_functions_js.go#L11101) |
| [SDL_AddEventWatch](https://wiki.libsdl.org/SDL3/SDL_AddEventWatch) | [:question:]() | [:question:](./sdl_functions_js.go#L11122) |
| [SDL_RemoveEventWatch](https://wiki.libsdl.org/SDL3/SDL_RemoveEventWatch) | [:question:]() | [:question:](./sdl_functions_js.go#L11137) |
| [SDL_FilterEvents](https://wiki.libsdl.org/SDL3/SDL_FilterEvents) | [:question:]() | [:question:](./sdl_functions_js.go#L11150) |
| [SDL_SetEventEnabled](https://wiki.libsdl.org/SDL3/SDL_SetEventEnabled) | [:heavy_check_mark:](./functions.go#L251) | [:x:](./sdl_functions_js.go#L11163) |
| [SDL_EventEnabled](https://wiki.libsdl.org/SDL3/SDL_EventEnabled) | [:heavy_check_mark:](./functions.go#L257) | [:x:](./sdl_functions_js.go#L11176) |
| [SDL_RegisterEvents](https://wiki.libsdl.org/SDL3/SDL_RegisterEvents) | [:question:]() | [:question:](./sdl_functions_js.go#L11189) |
| [SDL_GetWindowFromEvent](https://wiki.libsdl.org/SDL3/SDL_GetWindowFromEvent) | [:heavy_check_mark:](./methods.go#L1875) | [:x:](./sdl_functions_js.go#L11202) |
### Keyboard

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_HasKeyboard](https://wiki.libsdl.org/SDL3/SDL_HasKeyboard) | [:heavy_check_mark:](./functions.go#L1006) | [:x:](./sdl_functions_js.go#L10161) |
| [SDL_GetKeyboards](https://wiki.libsdl.org/SDL3/SDL_GetKeyboards) | [:heavy_check_mark:](./functions.go#L1012) | [:x:](./sdl_functions_js.go#L10172) |
| [SDL_GetKeyboardNameForID](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardNameForID) | [:heavy_check_mark:](./methods.go#L3575) | [:x:](./sdl_functions_js.go#L10188) |
| [SDL_GetKeyboardFocus](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardFocus) | [:question:]() | [:question:](./sdl_functions_js.go#L10201) |
| [SDL_GetKeyboardState](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardState) | [:heavy_check_mark:](./functions.go#L1026) | [:x:](./sdl_functions_js.go#L10215) |
| [SDL_ResetKeyboard](https://wiki.libsdl.org/SDL3/SDL_ResetKeyboard) | [:heavy_check_mark:](./functions.go#L1036) | [:x:](./sdl_functions_js.go#L10233) |
| [SDL_GetModState](https://wiki.libsdl.org/SDL3/SDL_GetModState) | [:question:]() | [:question:](./sdl_functions_js.go#L10242) |
| [SDL_SetModState](https://wiki.libsdl.org/SDL3/SDL_SetModState) | [:x:](./methods.go#L2703) | [:x:](./sdl_functions_js.go#L10253) |
| [SDL_GetKeyFromScancode](https://wiki.libsdl.org/SDL3/SDL_GetKeyFromScancode) | [:heavy_check_mark:](./methods.go#L4694) | [:x:](./sdl_functions_js.go#L10264) |
| [SDL_GetScancodeFromKey](https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromKey) | [:heavy_check_mark:](./methods.go#L5846) | [:x:](./sdl_functions_js.go#L10281) |
| [SDL_SetScancodeName](https://wiki.libsdl.org/SDL3/SDL_SetScancodeName) | [:heavy_check_mark:](./methods.go#L4700) | [:x:](./sdl_functions_js.go#L10299) |
| [SDL_GetScancodeName](https://wiki.libsdl.org/SDL3/SDL_GetScancodeName) | [:heavy_check_mark:](./methods.go#L4710) | [:x:](./sdl_functions_js.go#L10314) |
| [SDL_GetScancodeFromName](https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromName) | [:heavy_check_mark:](./functions.go#L1042) | [:x:](./sdl_functions_js.go#L10327) |
| [SDL_GetKeyName](https://wiki.libsdl.org/SDL3/SDL_GetKeyName) | [:heavy_check_mark:](./methods.go#L5852) | [:heavy_check_mark:](./sdl_functions_js.go#L10340) |
| [SDL_GetKeyFromName](https://wiki.libsdl.org/SDL3/SDL_GetKeyFromName) | [:heavy_check_mark:](./functions.go#L1048) | [:x:](./sdl_functions_js.go#L10350) |
| [SDL_StartTextInput](https://wiki.libsdl.org/SDL3/SDL_StartTextInput) | [:heavy_check_mark:](./methods.go#L4574) | [:x:](./sdl_functions_js.go#L10363) |
| [SDL_StartTextInputWithProperties](https://wiki.libsdl.org/SDL3/SDL_StartTextInputWithProperties) | [:heavy_check_mark:](./methods.go#L4584) | [:x:](./sdl_functions_js.go#L10379) |
| [SDL_TextInputActive](https://wiki.libsdl.org/SDL3/SDL_TextInputActive) | [:heavy_check_mark:](./methods.go#L4594) | [:x:](./sdl_functions_js.go#L10397) |
| [SDL_StopTextInput](https://wiki.libsdl.org/SDL3/SDL_StopTextInput) | [:heavy_check_mark:](./methods.go#L4600) | [:x:](./sdl_functions_js.go#L10413) |
| [SDL_ClearComposition](https://wiki.libsdl.org/SDL3/SDL_ClearComposition) | [:x:](./methods.go#L4610) | [:x:](./sdl_functions_js.go#L10429) |
| [SDL_SetTextInputArea](https://wiki.libsdl.org/SDL3/SDL_SetTextInputArea) | [:heavy_check_mark:](./methods.go#L4617) | [:x:](./sdl_functions_js.go#L10445) |
| [SDL_GetTextInputArea](https://wiki.libsdl.org/SDL3/SDL_GetTextInputArea) | [:heavy_check_mark:](./methods.go#L4627) | [:x:](./sdl_functions_js.go#L10468) |
| [SDL_HasScreenKeyboardSupport](https://wiki.libsdl.org/SDL3/SDL_HasScreenKeyboardSupport) | [:heavy_check_mark:](./functions.go#L1054) | [:x:](./sdl_functions_js.go#L10494) |
| [SDL_ScreenKeyboardShown](https://wiki.libsdl.org/SDL3/SDL_ScreenKeyboardShown) | [:heavy_check_mark:](./methods.go#L4640) | [:x:](./sdl_functions_js.go#L10505) |
### Mouse

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_HasMouse](https://wiki.libsdl.org/SDL3/SDL_HasMouse) | [:heavy_check_mark:](./functions.go#L1062) | [:x:](./sdl_functions_js.go#L10521) |
| [SDL_GetMice](https://wiki.libsdl.org/SDL3/SDL_GetMice) | [:heavy_check_mark:](./functions.go#L1068) | [:x:](./sdl_functions_js.go#L10532) |
| [SDL_GetMouseNameForID](https://wiki.libsdl.org/SDL3/SDL_GetMouseNameForID) | [:heavy_check_mark:](./methods.go#L3588) | [:x:](./sdl_functions_js.go#L10548) |
| [SDL_GetMouseFocus](https://wiki.libsdl.org/SDL3/SDL_GetMouseFocus) | [:heavy_check_mark:](./functions.go#L1082) | [:x:](./sdl_functions_js.go#L10561) |
| [SDL_GetMouseState](https://wiki.libsdl.org/SDL3/SDL_GetMouseState) | [:heavy_check_mark:](./functions.go#L1088) | [:x:](./sdl_functions_js.go#L10575) |
| [SDL_GetGlobalMouseState](https://wiki.libsdl.org/SDL3/SDL_GetGlobalMouseState) | [:heavy_check_mark:](./functions.go#L1098) | [:x:](./sdl_functions_js.go#L10596) |
| [SDL_GetRelativeMouseState](https://wiki.libsdl.org/SDL3/SDL_GetRelativeMouseState) | [:heavy_check_mark:](./functions.go#L1108) | [:x:](./sdl_functions_js.go#L10617) |
| [SDL_WarpMouseInWindow](https://wiki.libsdl.org/SDL3/SDL_WarpMouseInWindow) | [:heavy_check_mark:](./methods.go#L4646) | [:x:](./sdl_functions_js.go#L10638) |
| [SDL_WarpMouseGlobal](https://wiki.libsdl.org/SDL3/SDL_WarpMouseGlobal) | [:heavy_check_mark:](./functions.go#L1118) | [:x:](./sdl_functions_js.go#L10656) |
| [SDL_SetRelativeMouseTransform](https://wiki.libsdl.org/SDL3/SDL_SetRelativeMouseTransform) | [:question:]() | [:question:]() |
| [SDL_SetWindowRelativeMouseMode](https://wiki.libsdl.org/SDL3/SDL_SetWindowRelativeMouseMode) | [:heavy_check_mark:](./methods.go#L4652) | [:x:](./sdl_functions_js.go#L10671) |
| [SDL_GetWindowRelativeMouseMode](https://wiki.libsdl.org/SDL3/SDL_GetWindowRelativeMouseMode) | [:heavy_check_mark:](./methods.go#L4662) | [:x:](./sdl_functions_js.go#L10689) |
| [SDL_CaptureMouse](https://wiki.libsdl.org/SDL3/SDL_CaptureMouse) | [:heavy_check_mark:](./functions.go#L1128) | [:x:](./sdl_functions_js.go#L10705) |
| [SDL_CreateCursor](https://wiki.libsdl.org/SDL3/SDL_CreateCursor) | [:heavy_check_mark:](./functions.go#L1138) | [:x:](./sdl_functions_js.go#L10718) |
| [SDL_CreateColorCursor](https://wiki.libsdl.org/SDL3/SDL_CreateColorCursor) | [:heavy_check_mark:](./methods.go#L1851) | [:x:](./sdl_functions_js.go#L10750) |
| [SDL_CreateSystemCursor](https://wiki.libsdl.org/SDL3/SDL_CreateSystemCursor) | [:question:]() | [:question:](./sdl_functions_js.go#L10773) |
| [SDL_SetCursor](https://wiki.libsdl.org/SDL3/SDL_SetCursor) | [:heavy_check_mark:](./functions.go#L1153) | [:x:](./sdl_functions_js.go#L10789) |
| [SDL_GetCursor](https://wiki.libsdl.org/SDL3/SDL_GetCursor) | [:heavy_check_mark:](./functions.go#L1163) | [:x:](./sdl_functions_js.go#L10805) |
| [SDL_GetDefaultCursor](https://wiki.libsdl.org/SDL3/SDL_GetDefaultCursor) | [:heavy_check_mark:](./functions.go#L1169) | [:x:](./sdl_functions_js.go#L10819) |
| [SDL_DestroyCursor](https://wiki.libsdl.org/SDL3/SDL_DestroyCursor) | [:heavy_check_mark:](./methods.go#L593) | [:x:](./sdl_functions_js.go#L10833) |
| [SDL_ShowCursor](https://wiki.libsdl.org/SDL3/SDL_ShowCursor) | [:heavy_check_mark:](./functions.go#L1180) | [:x:](./sdl_functions_js.go#L10847) |
| [SDL_HideCursor](https://wiki.libsdl.org/SDL3/SDL_HideCursor) | [:heavy_check_mark:](./functions.go#L1190) | [:x:](./sdl_functions_js.go#L10858) |
| [SDL_CursorVisible](https://wiki.libsdl.org/SDL3/SDL_CursorVisible) | [:heavy_check_mark:](./functions.go#L1200) | [:x:](./sdl_functions_js.go#L10869) |
### Touch

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetTouchDevices](https://wiki.libsdl.org/SDL3/SDL_GetTouchDevices) | [:heavy_check_mark:](./functions.go#L1208) | [:x:](./sdl_functions_js.go#L10880) |
| [SDL_GetTouchDeviceName](https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceName) | [:heavy_check_mark:](./methods.go#L63) | [:x:](./sdl_functions_js.go#L10896) |
| [SDL_GetTouchDeviceType](https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceType) | [:heavy_check_mark:](./methods.go#L74) | [:x:](./sdl_functions_js.go#L10909) |
| [SDL_GetTouchFingers](https://wiki.libsdl.org/SDL3/SDL_GetTouchFingers) | [:heavy_check_mark:](./methods.go#L80) | [:x:](./sdl_functions_js.go#L10922) |
### Gamepad

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AddGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMapping) | [:heavy_check_mark:](./functions.go#L1224) | [:x:](./sdl_functions_js.go#L8986) |
| [SDL_AddGamepadMappingsFromIO](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromIO) | [:question:]() | [:question:](./sdl_functions_js.go#L8999) |
| [SDL_AddGamepadMappingsFromFile](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromFile) | [:heavy_check_mark:](./functions.go#L1234) | [:x:](./sdl_functions_js.go#L9017) |
| [SDL_ReloadGamepadMappings](https://wiki.libsdl.org/SDL3/SDL_ReloadGamepadMappings) | [:heavy_check_mark:](./functions.go#L1244) | [:x:](./sdl_functions_js.go#L9030) |
| [SDL_GetGamepadMappings](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappings) | [:heavy_check_mark:](./functions.go#L1254) | [:x:](./sdl_functions_js.go#L9041) |
| [SDL_GetGamepadMappingForGUID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForGUID) | [:question:]() | [:question:](./sdl_functions_js.go#L9057) |
| [SDL_GetGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMapping) | [:heavy_check_mark:](./methods.go#L2393) | [:x:](./sdl_functions_js.go#L9070) |
| [SDL_SetGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_SetGamepadMapping) | [:x:](./methods.go#L755) | [:x:](./sdl_functions_js.go#L9086) |
| [SDL_HasGamepad](https://wiki.libsdl.org/SDL3/SDL_HasGamepad) | [:heavy_check_mark:](./functions.go#L1268) | [:x:](./sdl_functions_js.go#L9101) |
| [SDL_GetGamepads](https://wiki.libsdl.org/SDL3/SDL_GetGamepads) | [:heavy_check_mark:](./functions.go#L1274) | [:x:](./sdl_functions_js.go#L9112) |
| [SDL_IsGamepad](https://wiki.libsdl.org/SDL3/SDL_IsGamepad) | [:heavy_check_mark:](./methods.go#L762) | [:x:](./sdl_functions_js.go#L9128) |
| [SDL_GetGamepadNameForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadNameForID) | [:heavy_check_mark:](./methods.go#L768) | [:x:](./sdl_functions_js.go#L9141) |
| [SDL_GetGamepadPathForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPathForID) | [:heavy_check_mark:](./methods.go#L779) | [:x:](./sdl_functions_js.go#L9154) |
| [SDL_GetGamepadPlayerIndexForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndexForID) | [:heavy_check_mark:](./methods.go#L790) | [:x:](./sdl_functions_js.go#L9167) |
| [SDL_GetGamepadGUIDForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadGUIDForID) | [:x:](./methods.go#L796) | [:x:](./sdl_functions_js.go#L9180) |
| [SDL_GetGamepadVendorForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendorForID) | [:heavy_check_mark:](./methods.go#L803) | [:x:](./sdl_functions_js.go#L9193) |
| [SDL_GetGamepadProductForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductForID) | [:heavy_check_mark:](./methods.go#L809) | [:x:](./sdl_functions_js.go#L9206) |
| [SDL_GetGamepadProductVersionForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersionForID) | [:heavy_check_mark:](./methods.go#L815) | [:x:](./sdl_functions_js.go#L9219) |
| [SDL_GetGamepadTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeForID) | [:heavy_check_mark:](./methods.go#L821) | [:x:](./sdl_functions_js.go#L9232) |
| [SDL_GetRealGamepadTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadTypeForID) | [:heavy_check_mark:](./methods.go#L827) | [:x:](./sdl_functions_js.go#L9245) |
| [SDL_GetGamepadMappingForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForID) | [:heavy_check_mark:](./methods.go#L833) | [:x:](./sdl_functions_js.go#L9258) |
| [SDL_OpenGamepad](https://wiki.libsdl.org/SDL3/SDL_OpenGamepad) | [:heavy_check_mark:](./methods.go#L845) | [:x:](./sdl_functions_js.go#L9271) |
| [SDL_GetGamepadFromID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromID) | [:heavy_check_mark:](./methods.go#L856) | [:x:](./sdl_functions_js.go#L9287) |
| [SDL_GetGamepadFromPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromPlayerIndex) | [:heavy_check_mark:](./functions.go#L1288) | [:x:](./sdl_functions_js.go#L9303) |
| [SDL_GetGamepadProperties](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProperties) | [:heavy_check_mark:](./methods.go#L2405) | [:x:](./sdl_functions_js.go#L9319) |
| [SDL_GetGamepadID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadID) | [:heavy_check_mark:](./methods.go#L2416) | [:x:](./sdl_functions_js.go#L9335) |
| [SDL_GetGamepadName](https://wiki.libsdl.org/SDL3/SDL_GetGamepadName) | [:heavy_check_mark:](./methods.go#L2427) | [:x:](./sdl_functions_js.go#L9351) |
| [SDL_GetGamepadPath](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPath) | [:heavy_check_mark:](./methods.go#L2433) | [:x:](./sdl_functions_js.go#L9367) |
| [SDL_GetGamepadType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadType) | [:heavy_check_mark:](./methods.go#L2439) | [:x:](./sdl_functions_js.go#L9383) |
| [SDL_GetRealGamepadType](https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadType) | [:heavy_check_mark:](./methods.go#L2445) | [:x:](./sdl_functions_js.go#L9399) |
| [SDL_GetGamepadPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndex) | [:heavy_check_mark:](./methods.go#L2451) | [:x:](./sdl_functions_js.go#L9415) |
| [SDL_SetGamepadPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_SetGamepadPlayerIndex) | [:heavy_check_mark:](./methods.go#L2457) | [:x:](./sdl_functions_js.go#L9431) |
| [SDL_GetGamepadVendor](https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendor) | [:heavy_check_mark:](./methods.go#L2467) | [:x:](./sdl_functions_js.go#L9449) |
| [SDL_GetGamepadProduct](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProduct) | [:heavy_check_mark:](./methods.go#L2473) | [:x:](./sdl_functions_js.go#L9465) |
| [SDL_GetGamepadProductVersion](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersion) | [:heavy_check_mark:](./methods.go#L2479) | [:x:](./sdl_functions_js.go#L9481) |
| [SDL_GetGamepadFirmwareVersion](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFirmwareVersion) | [:heavy_check_mark:](./methods.go#L2485) | [:x:](./sdl_functions_js.go#L9497) |
| [SDL_GetGamepadSerial](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSerial) | [:heavy_check_mark:](./methods.go#L2491) | [:x:](./sdl_functions_js.go#L9513) |
| [SDL_GetGamepadSteamHandle](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSteamHandle) | [:heavy_check_mark:](./methods.go#L2497) | [:x:](./sdl_functions_js.go#L9529) |
| [SDL_GetGamepadConnectionState](https://wiki.libsdl.org/SDL3/SDL_GetGamepadConnectionState) | [:heavy_check_mark:](./methods.go#L2503) | [:x:](./sdl_functions_js.go#L9545) |
| [SDL_GetGamepadPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPowerInfo) | [:heavy_check_mark:](./methods.go#L2515) | [:x:](./sdl_functions_js.go#L9561) |
| [SDL_GamepadConnected](https://wiki.libsdl.org/SDL3/SDL_GamepadConnected) | [:heavy_check_mark:](./methods.go#L2525) | [:x:](./sdl_functions_js.go#L9582) |
| [SDL_GetGamepadJoystick](https://wiki.libsdl.org/SDL3/SDL_GetGamepadJoystick) | [:heavy_check_mark:](./methods.go#L2531) | [:x:](./sdl_functions_js.go#L9598) |
| [SDL_SetGamepadEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_SetGamepadEventsEnabled) | [:heavy_check_mark:](./functions.go#L1294) | [:x:](./sdl_functions_js.go#L9617) |
| [SDL_GamepadEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_GamepadEventsEnabled) | [:heavy_check_mark:](./functions.go#L1300) | [:x:](./sdl_functions_js.go#L9628) |
| [SDL_GetGamepadBindings](https://wiki.libsdl.org/SDL3/SDL_GetGamepadBindings) | [:heavy_check_mark:](./methods.go#L2542) | [:x:](./sdl_functions_js.go#L9639) |
| [SDL_UpdateGamepads](https://wiki.libsdl.org/SDL3/SDL_UpdateGamepads) | [:heavy_check_mark:](./functions.go#L1306) | [:x:](./sdl_functions_js.go#L9660) |
| [SDL_GetGamepadTypeFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeFromString) | [:question:]() | [:question:](./sdl_functions_js.go#L9669) |
| [SDL_GetGamepadStringForType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForType) | [:heavy_check_mark:](./methods.go#L5535) | [:x:](./sdl_functions_js.go#L9682) |
| [SDL_GetGamepadAxisFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxisFromString) | [:heavy_check_mark:](./functions.go#L1312) | [:x:](./sdl_functions_js.go#L9695) |
| [SDL_GetGamepadStringForAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForAxis) | [:heavy_check_mark:](./methods.go#L585) | [:x:](./sdl_functions_js.go#L9708) |
| [SDL_GamepadHasAxis](https://wiki.libsdl.org/SDL3/SDL_GamepadHasAxis) | [:heavy_check_mark:](./methods.go#L2556) | [:x:](./sdl_functions_js.go#L9721) |
| [SDL_GetGamepadAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxis) | [:heavy_check_mark:](./methods.go#L2562) | [:x:](./sdl_functions_js.go#L9739) |
| [SDL_GetGamepadButtonFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonFromString) | [:heavy_check_mark:](./functions.go#L1318) | [:x:](./sdl_functions_js.go#L9757) |
| [SDL_GetGamepadStringForButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForButton) | [:heavy_check_mark:](./methods.go#L443) | [:x:](./sdl_functions_js.go#L9770) |
| [SDL_GamepadHasButton](https://wiki.libsdl.org/SDL3/SDL_GamepadHasButton) | [:heavy_check_mark:](./methods.go#L2568) | [:x:](./sdl_functions_js.go#L9783) |
| [SDL_GetGamepadButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButton) | [:heavy_check_mark:](./methods.go#L2574) | [:x:](./sdl_functions_js.go#L9801) |
| [SDL_GetGamepadButtonLabelForType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabelForType) | [:heavy_check_mark:](./methods.go#L5541) | [:x:](./sdl_functions_js.go#L9819) |
| [SDL_GetGamepadButtonLabel](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabel) | [:heavy_check_mark:](./methods.go#L2580) | [:x:](./sdl_functions_js.go#L9834) |
| [SDL_GetNumGamepadTouchpads](https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpads) | [:heavy_check_mark:](./methods.go#L2586) | [:x:](./sdl_functions_js.go#L9852) |
| [SDL_GetNumGamepadTouchpadFingers](https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpadFingers) | [:heavy_check_mark:](./methods.go#L2592) | [:x:](./sdl_functions_js.go#L9868) |
| [SDL_GetGamepadTouchpadFinger](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTouchpadFinger) | [:x:](./methods.go#L2598) | [:x:](./sdl_functions_js.go#L9886) |
| [SDL_GamepadHasSensor](https://wiki.libsdl.org/SDL3/SDL_GamepadHasSensor) | [:heavy_check_mark:](./methods.go#L2605) | [:x:](./sdl_functions_js.go#L9926) |
| [SDL_SetGamepadSensorEnabled](https://wiki.libsdl.org/SDL3/SDL_SetGamepadSensorEnabled) | [:heavy_check_mark:](./methods.go#L2611) | [:x:](./sdl_functions_js.go#L9944) |
| [SDL_GamepadSensorEnabled](https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEnabled) | [:heavy_check_mark:](./methods.go#L2621) | [:x:](./sdl_functions_js.go#L9964) |
| [SDL_GetGamepadSensorDataRate](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorDataRate) | [:heavy_check_mark:](./methods.go#L2627) | [:x:](./sdl_functions_js.go#L9982) |
| [SDL_GetGamepadSensorData](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorData) | [:x:](./methods.go#L2633) | [:x:](./sdl_functions_js.go#L10000) |
| [SDL_RumbleGamepad](https://wiki.libsdl.org/SDL3/SDL_RumbleGamepad) | [:heavy_check_mark:](./methods.go#L2640) | [:x:](./sdl_functions_js.go#L10025) |
| [SDL_RumbleGamepadTriggers](https://wiki.libsdl.org/SDL3/SDL_RumbleGamepadTriggers) | [:heavy_check_mark:](./methods.go#L2650) | [:x:](./sdl_functions_js.go#L10047) |
| [SDL_SetGamepadLED](https://wiki.libsdl.org/SDL3/SDL_SetGamepadLED) | [:heavy_check_mark:](./methods.go#L2660) | [:x:](./sdl_functions_js.go#L10069) |
| [SDL_SendGamepadEffect](https://wiki.libsdl.org/SDL3/SDL_SendGamepadEffect) | [:heavy_check_mark:](./methods.go#L2670) | [:x:](./sdl_functions_js.go#L10091) |
| [SDL_CloseGamepad](https://wiki.libsdl.org/SDL3/SDL_CloseGamepad) | [:heavy_check_mark:](./methods.go#L2681) | [:x:](./sdl_functions_js.go#L10111) |
| [SDL_GetGamepadAppleSFSymbolsNameForButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForButton) | [:x:](./methods.go#L2687) | [:x:](./sdl_functions_js.go#L10125) |
| [SDL_GetGamepadAppleSFSymbolsNameForAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForAxis) | [:x:](./methods.go#L2694) | [:x:](./sdl_functions_js.go#L10143) |
### Joystick

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_LockJoysticks](https://wiki.libsdl.org/SDL3/SDL_LockJoysticks) | [:heavy_check_mark:](./functions.go#L1326) | [:x:](./sdl_functions_js.go#L8047) |
| [SDL_UnlockJoysticks](https://wiki.libsdl.org/SDL3/SDL_UnlockJoysticks) | [:heavy_check_mark:](./functions.go#L1332) | [:x:](./sdl_functions_js.go#L8056) |
| [SDL_HasJoystick](https://wiki.libsdl.org/SDL3/SDL_HasJoystick) | [:heavy_check_mark:](./functions.go#L1338) | [:x:](./sdl_functions_js.go#L8065) |
| [SDL_GetJoysticks](https://wiki.libsdl.org/SDL3/SDL_GetJoysticks) | [:heavy_check_mark:](./functions.go#L1344) | [:x:](./sdl_functions_js.go#L8076) |
| [SDL_GetJoystickNameForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickNameForID) | [:heavy_check_mark:](./methods.go#L658) | [:x:](./sdl_functions_js.go#L8092) |
| [SDL_GetJoystickPathForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPathForID) | [:heavy_check_mark:](./methods.go#L669) | [:x:](./sdl_functions_js.go#L8105) |
| [SDL_GetJoystickPlayerIndexForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndexForID) | [:heavy_check_mark:](./methods.go#L680) | [:x:](./sdl_functions_js.go#L8118) |
| [SDL_GetJoystickGUIDForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDForID) | [:x:](./methods.go#L686) | [:x:](./sdl_functions_js.go#L8131) |
| [SDL_GetJoystickVendorForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendorForID) | [:heavy_check_mark:](./methods.go#L693) | [:x:](./sdl_functions_js.go#L8144) |
| [SDL_GetJoystickProductForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductForID) | [:heavy_check_mark:](./methods.go#L699) | [:x:](./sdl_functions_js.go#L8157) |
| [SDL_GetJoystickProductVersionForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersionForID) | [:heavy_check_mark:](./methods.go#L705) | [:x:](./sdl_functions_js.go#L8170) |
| [SDL_GetJoystickTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickTypeForID) | [:heavy_check_mark:](./methods.go#L711) | [:x:](./sdl_functions_js.go#L8183) |
| [SDL_OpenJoystick](https://wiki.libsdl.org/SDL3/SDL_OpenJoystick) | [:heavy_check_mark:](./methods.go#L717) | [:heavy_check_mark:](./sdl_functions_js.go#L8196) |
| [SDL_GetJoystickFromID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromID) | [:heavy_check_mark:](./methods.go#L728) | [:heavy_check_mark:](./sdl_functions_js.go#L8208) |
| [SDL_GetJoystickFromPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromPlayerIndex) | [:heavy_check_mark:](./functions.go#L1358) | [:x:](./sdl_functions_js.go#L8222) |
| [SDL_AttachVirtualJoystick](https://wiki.libsdl.org/SDL3/SDL_AttachVirtualJoystick) | [:heavy_check_mark:](./functions.go#L1364) | [:x:](./sdl_functions_js.go#L8238) |
| [SDL_DetachVirtualJoystick](https://wiki.libsdl.org/SDL3/SDL_DetachVirtualJoystick) | [:heavy_check_mark:](./methods.go#L739) | [:x:](./sdl_functions_js.go#L8254) |
| [SDL_IsJoystickVirtual](https://wiki.libsdl.org/SDL3/SDL_IsJoystickVirtual) | [:heavy_check_mark:](./methods.go#L749) | [:x:](./sdl_functions_js.go#L8267) |
| [SDL_SetJoystickVirtualAxis](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualAxis) | [:heavy_check_mark:](./methods.go#L5193) | [:x:](./sdl_functions_js.go#L8280) |
| [SDL_SetJoystickVirtualBall](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualBall) | [:heavy_check_mark:](./methods.go#L5203) | [:x:](./sdl_functions_js.go#L8300) |
| [SDL_SetJoystickVirtualButton](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualButton) | [:heavy_check_mark:](./methods.go#L5213) | [:x:](./sdl_functions_js.go#L8322) |
| [SDL_SetJoystickVirtualHat](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualHat) | [:heavy_check_mark:](./methods.go#L5223) | [:x:](./sdl_functions_js.go#L8342) |
| [SDL_SetJoystickVirtualTouchpad](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualTouchpad) | [:heavy_check_mark:](./methods.go#L5233) | [:x:](./sdl_functions_js.go#L8362) |
| [SDL_SendJoystickVirtualSensorData](https://wiki.libsdl.org/SDL3/SDL_SendJoystickVirtualSensorData) | [:heavy_check_mark:](./methods.go#L5243) | [:x:](./sdl_functions_js.go#L8390) |
| [SDL_GetJoystickProperties](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProperties) | [:heavy_check_mark:](./methods.go#L5253) | [:x:](./sdl_functions_js.go#L8417) |
| [SDL_GetJoystickName](https://wiki.libsdl.org/SDL3/SDL_GetJoystickName) | [:heavy_check_mark:](./methods.go#L5264) | [:heavy_check_mark:](./sdl_functions_js.go#L8433) |
| [SDL_GetJoystickPath](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPath) | [:heavy_check_mark:](./methods.go#L5275) | [:x:](./sdl_functions_js.go#L8446) |
| [SDL_GetJoystickPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndex) | [:heavy_check_mark:](./methods.go#L5286) | [:x:](./sdl_functions_js.go#L8462) |
| [SDL_SetJoystickPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_SetJoystickPlayerIndex) | [:heavy_check_mark:](./methods.go#L5292) | [:x:](./sdl_functions_js.go#L8478) |
| [SDL_GetJoystickGUID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUID) | [:x:](./methods.go#L5302) | [:x:](./sdl_functions_js.go#L8496) |
| [SDL_GetJoystickVendor](https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendor) | [:heavy_check_mark:](./methods.go#L5309) | [:x:](./sdl_functions_js.go#L8512) |
| [SDL_GetJoystickProduct](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProduct) | [:heavy_check_mark:](./methods.go#L5315) | [:x:](./sdl_functions_js.go#L8528) |
| [SDL_GetJoystickProductVersion](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersion) | [:heavy_check_mark:](./methods.go#L5321) | [:x:](./sdl_functions_js.go#L8544) |
| [SDL_GetJoystickFirmwareVersion](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFirmwareVersion) | [:heavy_check_mark:](./methods.go#L5327) | [:x:](./sdl_functions_js.go#L8560) |
| [SDL_GetJoystickSerial](https://wiki.libsdl.org/SDL3/SDL_GetJoystickSerial) | [:heavy_check_mark:](./methods.go#L5333) | [:x:](./sdl_functions_js.go#L8576) |
| [SDL_GetJoystickType](https://wiki.libsdl.org/SDL3/SDL_GetJoystickType) | [:heavy_check_mark:](./methods.go#L5339) | [:x:](./sdl_functions_js.go#L8592) |
| [SDL_GetJoystickGUIDInfo](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDInfo) | [:question:]() | [:question:](./sdl_functions_js.go#L8608) |
| [SDL_JoystickConnected](https://wiki.libsdl.org/SDL3/SDL_JoystickConnected) | [:heavy_check_mark:](./methods.go#L5345) | [:x:](./sdl_functions_js.go#L8639) |
| [SDL_GetJoystickID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickID) | [:heavy_check_mark:](./methods.go#L5351) | [:heavy_check_mark:](./sdl_functions_js.go#L8655) |
| [SDL_GetNumJoystickAxes](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickAxes) | [:heavy_check_mark:](./methods.go#L5362) | [:heavy_check_mark:](./sdl_functions_js.go#L8668) |
| [SDL_GetNumJoystickBalls](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickBalls) | [:heavy_check_mark:](./methods.go#L5373) | [:x:](./sdl_functions_js.go#L8681) |
| [SDL_GetNumJoystickHats](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickHats) | [:heavy_check_mark:](./methods.go#L5384) | [:heavy_check_mark:](./sdl_functions_js.go#L8697) |
| [SDL_GetNumJoystickButtons](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickButtons) | [:heavy_check_mark:](./methods.go#L5395) | [:heavy_check_mark:](./sdl_functions_js.go#L8710) |
| [SDL_SetJoystickEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_SetJoystickEventsEnabled) | [:heavy_check_mark:](./functions.go#L1370) | [:x:](./sdl_functions_js.go#L8723) |
| [SDL_JoystickEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_JoystickEventsEnabled) | [:heavy_check_mark:](./functions.go#L1376) | [:x:](./sdl_functions_js.go#L8734) |
| [SDL_UpdateJoysticks](https://wiki.libsdl.org/SDL3/SDL_UpdateJoysticks) | [:heavy_check_mark:](./functions.go#L1382) | [:x:](./sdl_functions_js.go#L8745) |
| [SDL_GetJoystickAxis](https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxis) | [:heavy_check_mark:](./methods.go#L5406) | [:heavy_check_mark:](./sdl_functions_js.go#L8754) |
| [SDL_GetJoystickAxisInitialState](https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxisInitialState) | [:heavy_check_mark:](./methods.go#L5417) | [:x:](./sdl_functions_js.go#L8769) |
| [SDL_GetJoystickBall](https://wiki.libsdl.org/SDL3/SDL_GetJoystickBall) | [:heavy_check_mark:](./methods.go#L5426) | [:x:](./sdl_functions_js.go#L8792) |
| [SDL_GetJoystickHat](https://wiki.libsdl.org/SDL3/SDL_GetJoystickHat) | [:heavy_check_mark:](./methods.go#L5438) | [:heavy_check_mark:](./sdl_functions_js.go#L8820) |
| [SDL_GetJoystickButton](https://wiki.libsdl.org/SDL3/SDL_GetJoystickButton) | [:heavy_check_mark:](./methods.go#L5444) | [:heavy_check_mark:](./sdl_functions_js.go#L8835) |
| [SDL_RumbleJoystick](https://wiki.libsdl.org/SDL3/SDL_RumbleJoystick) | [:heavy_check_mark:](./methods.go#L5450) | [:x:](./sdl_functions_js.go#L8850) |
| [SDL_RumbleJoystickTriggers](https://wiki.libsdl.org/SDL3/SDL_RumbleJoystickTriggers) | [:heavy_check_mark:](./methods.go#L5456) | [:x:](./sdl_functions_js.go#L8872) |
| [SDL_SetJoystickLED](https://wiki.libsdl.org/SDL3/SDL_SetJoystickLED) | [:heavy_check_mark:](./methods.go#L5466) | [:x:](./sdl_functions_js.go#L8894) |
| [SDL_SendJoystickEffect](https://wiki.libsdl.org/SDL3/SDL_SendJoystickEffect) | [:heavy_check_mark:](./methods.go#L5476) | [:x:](./sdl_functions_js.go#L8916) |
| [SDL_CloseJoystick](https://wiki.libsdl.org/SDL3/SDL_CloseJoystick) | [:heavy_check_mark:](./methods.go#L5486) | [:heavy_check_mark:](./sdl_functions_js.go#L8936) |
| [SDL_GetJoystickConnectionState](https://wiki.libsdl.org/SDL3/SDL_GetJoystickConnectionState) | [:heavy_check_mark:](./methods.go#L5492) | [:x:](./sdl_functions_js.go#L8949) |
| [SDL_GetJoystickPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPowerInfo) | [:heavy_check_mark:](./methods.go#L5503) | [:x:](./sdl_functions_js.go#L8965) |
### Haptic

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetHaptics](https://wiki.libsdl.org/SDL3/SDL_GetHaptics) | [:question:]() | [:question:](./sdl_functions_js.go#L13276) |
| [SDL_GetHapticNameForID](https://wiki.libsdl.org/SDL3/SDL_GetHapticNameForID) | [:x:](./methods.go#L2763) | [:x:](./sdl_functions_js.go#L13292) |
| [SDL_OpenHaptic](https://wiki.libsdl.org/SDL3/SDL_OpenHaptic) | [:x:](./methods.go#L2770) | [:x:](./sdl_functions_js.go#L13305) |
| [SDL_GetHapticFromID](https://wiki.libsdl.org/SDL3/SDL_GetHapticFromID) | [:x:](./methods.go#L2777) | [:x:](./sdl_functions_js.go#L13321) |
| [SDL_GetHapticID](https://wiki.libsdl.org/SDL3/SDL_GetHapticID) | [:heavy_check_mark:](./methods.go#L2165) | [:x:](./sdl_functions_js.go#L13337) |
| [SDL_GetHapticName](https://wiki.libsdl.org/SDL3/SDL_GetHapticName) | [:heavy_check_mark:](./methods.go#L2176) | [:x:](./sdl_functions_js.go#L13353) |
| [SDL_IsMouseHaptic](https://wiki.libsdl.org/SDL3/SDL_IsMouseHaptic) | [:question:]() | [:question:](./sdl_functions_js.go#L13369) |
| [SDL_OpenHapticFromMouse](https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromMouse) | [:question:]() | [:question:](./sdl_functions_js.go#L13380) |
| [SDL_IsJoystickHaptic](https://wiki.libsdl.org/SDL3/SDL_IsJoystickHaptic) | [:heavy_check_mark:](./methods.go#L5516) | [:x:](./sdl_functions_js.go#L13394) |
| [SDL_OpenHapticFromJoystick](https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromJoystick) | [:heavy_check_mark:](./methods.go#L5522) | [:x:](./sdl_functions_js.go#L13410) |
| [SDL_CloseHaptic](https://wiki.libsdl.org/SDL3/SDL_CloseHaptic) | [:heavy_check_mark:](./methods.go#L2187) | [:x:](./sdl_functions_js.go#L13429) |
| [SDL_GetMaxHapticEffects](https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffects) | [:heavy_check_mark:](./methods.go#L2193) | [:x:](./sdl_functions_js.go#L13443) |
| [SDL_GetMaxHapticEffectsPlaying](https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffectsPlaying) | [:heavy_check_mark:](./methods.go#L2204) | [:x:](./sdl_functions_js.go#L13459) |
| [SDL_GetHapticFeatures](https://wiki.libsdl.org/SDL3/SDL_GetHapticFeatures) | [:heavy_check_mark:](./methods.go#L2215) | [:x:](./sdl_functions_js.go#L13475) |
| [SDL_GetNumHapticAxes](https://wiki.libsdl.org/SDL3/SDL_GetNumHapticAxes) | [:heavy_check_mark:](./methods.go#L2226) | [:x:](./sdl_functions_js.go#L13491) |
| [SDL_HapticEffectSupported](https://wiki.libsdl.org/SDL3/SDL_HapticEffectSupported) | [:heavy_check_mark:](./methods.go#L2237) | [:x:](./sdl_functions_js.go#L13507) |
| [SDL_CreateHapticEffect](https://wiki.libsdl.org/SDL3/SDL_CreateHapticEffect) | [:heavy_check_mark:](./methods.go#L2243) | [:x:](./sdl_functions_js.go#L13528) |
| [SDL_UpdateHapticEffect](https://wiki.libsdl.org/SDL3/SDL_UpdateHapticEffect) | [:heavy_check_mark:](./methods.go#L2254) | [:x:](./sdl_functions_js.go#L13549) |
| [SDL_RunHapticEffect](https://wiki.libsdl.org/SDL3/SDL_RunHapticEffect) | [:heavy_check_mark:](./methods.go#L2264) | [:x:](./sdl_functions_js.go#L13572) |
| [SDL_StopHapticEffect](https://wiki.libsdl.org/SDL3/SDL_StopHapticEffect) | [:heavy_check_mark:](./methods.go#L2274) | [:x:](./sdl_functions_js.go#L13592) |
| [SDL_DestroyHapticEffect](https://wiki.libsdl.org/SDL3/SDL_DestroyHapticEffect) | [:heavy_check_mark:](./methods.go#L2284) | [:x:](./sdl_functions_js.go#L13610) |
| [SDL_GetHapticEffectStatus](https://wiki.libsdl.org/SDL3/SDL_GetHapticEffectStatus) | [:heavy_check_mark:](./methods.go#L2290) | [:x:](./sdl_functions_js.go#L13626) |
| [SDL_SetHapticGain](https://wiki.libsdl.org/SDL3/SDL_SetHapticGain) | [:heavy_check_mark:](./methods.go#L2296) | [:x:](./sdl_functions_js.go#L13644) |
| [SDL_SetHapticAutocenter](https://wiki.libsdl.org/SDL3/SDL_SetHapticAutocenter) | [:heavy_check_mark:](./methods.go#L2306) | [:x:](./sdl_functions_js.go#L13662) |
| [SDL_PauseHaptic](https://wiki.libsdl.org/SDL3/SDL_PauseHaptic) | [:heavy_check_mark:](./methods.go#L2316) | [:x:](./sdl_functions_js.go#L13680) |
| [SDL_ResumeHaptic](https://wiki.libsdl.org/SDL3/SDL_ResumeHaptic) | [:heavy_check_mark:](./methods.go#L2326) | [:x:](./sdl_functions_js.go#L13696) |
| [SDL_StopHapticEffects](https://wiki.libsdl.org/SDL3/SDL_StopHapticEffects) | [:heavy_check_mark:](./methods.go#L2336) | [:x:](./sdl_functions_js.go#L13712) |
| [SDL_HapticRumbleSupported](https://wiki.libsdl.org/SDL3/SDL_HapticRumbleSupported) | [:heavy_check_mark:](./methods.go#L2346) | [:x:](./sdl_functions_js.go#L13728) |
| [SDL_InitHapticRumble](https://wiki.libsdl.org/SDL3/SDL_InitHapticRumble) | [:heavy_check_mark:](./methods.go#L2352) | [:x:](./sdl_functions_js.go#L13744) |
| [SDL_PlayHapticRumble](https://wiki.libsdl.org/SDL3/SDL_PlayHapticRumble) | [:heavy_check_mark:](./methods.go#L2362) | [:x:](./sdl_functions_js.go#L13760) |
| [SDL_StopHapticRumble](https://wiki.libsdl.org/SDL3/SDL_StopHapticRumble) | [:heavy_check_mark:](./methods.go#L2372) | [:x:](./sdl_functions_js.go#L13780) |
### Audio

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumAudioDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumAudioDrivers) | [:heavy_check_mark:](./functions.go#L717) | [:x:](./sdl_functions_js.go#L2357) |
| [SDL_GetAudioDriver](https://wiki.libsdl.org/SDL3/SDL_GetAudioDriver) | [:heavy_check_mark:](./functions.go#L723) | [:x:](./sdl_functions_js.go#L2368) |
| [SDL_GetCurrentAudioDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentAudioDriver) | [:heavy_check_mark:](./functions.go#L729) | [:x:](./sdl_functions_js.go#L2381) |
| [SDL_GetAudioPlaybackDevices](https://wiki.libsdl.org/SDL3/SDL_GetAudioPlaybackDevices) | [:heavy_check_mark:](./functions.go#L735) | [:x:](./sdl_functions_js.go#L2392) |
| [SDL_GetAudioRecordingDevices](https://wiki.libsdl.org/SDL3/SDL_GetAudioRecordingDevices) | [:heavy_check_mark:](./functions.go#L749) | [:x:](./sdl_functions_js.go#L2408) |
| [SDL_GetAudioDeviceName](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceName) | [:heavy_check_mark:](./methods.go#L226) | [:x:](./sdl_functions_js.go#L2424) |
| [SDL_GetAudioDeviceFormat](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceFormat) | [:heavy_check_mark:](./methods.go#L237) | [:x:](./sdl_functions_js.go#L2437) |
| [SDL_GetAudioDeviceChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceChannelMap) | [:heavy_check_mark:](./methods.go#L250) | [:x:](./sdl_functions_js.go#L2460) |
| [SDL_OpenAudioDevice](https://wiki.libsdl.org/SDL3/SDL_OpenAudioDevice) | [:heavy_check_mark:](./methods.go#L264) | [:x:](./sdl_functions_js.go#L2478) |
| [SDL_IsAudioDevicePhysical](https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePhysical) | [:heavy_check_mark:](./methods.go#L275) | [:x:](./sdl_functions_js.go#L2496) |
| [SDL_IsAudioDevicePlayback](https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePlayback) | [:heavy_check_mark:](./methods.go#L281) | [:x:](./sdl_functions_js.go#L2509) |
| [SDL_PauseAudioDevice](https://wiki.libsdl.org/SDL3/SDL_PauseAudioDevice) | [:heavy_check_mark:](./methods.go#L287) | [:x:](./sdl_functions_js.go#L2522) |
| [SDL_ResumeAudioDevice](https://wiki.libsdl.org/SDL3/SDL_ResumeAudioDevice) | [:heavy_check_mark:](./methods.go#L297) | [:x:](./sdl_functions_js.go#L2535) |
| [SDL_AudioDevicePaused](https://wiki.libsdl.org/SDL3/SDL_AudioDevicePaused) | [:heavy_check_mark:](./methods.go#L307) | [:x:](./sdl_functions_js.go#L2548) |
| [SDL_GetAudioDeviceGain](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceGain) | [:heavy_check_mark:](./methods.go#L313) | [:x:](./sdl_functions_js.go#L2561) |
| [SDL_SetAudioDeviceGain](https://wiki.libsdl.org/SDL3/SDL_SetAudioDeviceGain) | [:heavy_check_mark:](./methods.go#L324) | [:x:](./sdl_functions_js.go#L2574) |
| [SDL_CloseAudioDevice](https://wiki.libsdl.org/SDL3/SDL_CloseAudioDevice) | [:heavy_check_mark:](./methods.go#L334) | [:x:](./sdl_functions_js.go#L2589) |
| [SDL_BindAudioStreams](https://wiki.libsdl.org/SDL3/SDL_BindAudioStreams) | [:heavy_check_mark:](./methods.go#L340) | [:x:](./sdl_functions_js.go#L2600) |
| [SDL_BindAudioStream](https://wiki.libsdl.org/SDL3/SDL_BindAudioStream) | [:heavy_check_mark:](./methods.go#L350) | [:x:](./sdl_functions_js.go#L2620) |
| [SDL_UnbindAudioStreams](https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStreams) | [:heavy_check_mark:](./functions.go#L763) | [:x:](./sdl_functions_js.go#L2638) |
| [SDL_UnbindAudioStream](https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStream) | [:heavy_check_mark:](./methods.go#L3610) | [:x:](./sdl_functions_js.go#L2654) |
| [SDL_GetAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamDevice) | [:heavy_check_mark:](./methods.go#L3616) | [:x:](./sdl_functions_js.go#L2668) |
| [SDL_CreateAudioStream](https://wiki.libsdl.org/SDL3/SDL_CreateAudioStream) | [:heavy_check_mark:](./functions.go#L769) | [:x:](./sdl_functions_js.go#L2684) |
| [SDL_GetAudioStreamProperties](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamProperties) | [:heavy_check_mark:](./methods.go#L3622) | [:x:](./sdl_functions_js.go#L2708) |
| [SDL_GetAudioStreamFormat](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFormat) | [:heavy_check_mark:](./methods.go#L3633) | [:x:](./sdl_functions_js.go#L2724) |
| [SDL_SetAudioStreamFormat](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFormat) | [:heavy_check_mark:](./methods.go#L3643) | [:x:](./sdl_functions_js.go#L2750) |
| [SDL_GetAudioStreamFrequencyRatio](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFrequencyRatio) | [:heavy_check_mark:](./methods.go#L3653) | [:x:](./sdl_functions_js.go#L2776) |
| [SDL_SetAudioStreamFrequencyRatio](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFrequencyRatio) | [:heavy_check_mark:](./methods.go#L3664) | [:x:](./sdl_functions_js.go#L2792) |
| [SDL_GetAudioStreamGain](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamGain) | [:heavy_check_mark:](./methods.go#L3674) | [:x:](./sdl_functions_js.go#L2810) |
| [SDL_SetAudioStreamGain](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGain) | [:heavy_check_mark:](./methods.go#L3685) | [:x:](./sdl_functions_js.go#L2826) |
| [SDL_GetAudioStreamInputChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamInputChannelMap) | [:heavy_check_mark:](./methods.go#L3695) | [:x:](./sdl_functions_js.go#L2844) |
| [SDL_GetAudioStreamOutputChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamOutputChannelMap) | [:heavy_check_mark:](./methods.go#L3709) | [:x:](./sdl_functions_js.go#L2865) |
| [SDL_SetAudioStreamInputChannelMap](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamInputChannelMap) | [:heavy_check_mark:](./methods.go#L3723) | [:x:](./sdl_functions_js.go#L2886) |
| [SDL_SetAudioStreamOutputChannelMap](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamOutputChannelMap) | [:heavy_check_mark:](./methods.go#L3733) | [:x:](./sdl_functions_js.go#L2909) |
| [SDL_PutAudioStreamData](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData) | [:heavy_check_mark:](./methods.go#L3743) | [:x:](./sdl_functions_js.go#L2932) |
| [SDL_GetAudioStreamData](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData) | [:heavy_check_mark:](./methods.go#L3753) | [:x:](./sdl_functions_js.go#L2952) |
| [SDL_GetAudioStreamAvailable](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamAvailable) | [:heavy_check_mark:](./methods.go#L3764) | [:x:](./sdl_functions_js.go#L2972) |
| [SDL_GetAudioStreamQueued](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamQueued) | [:heavy_check_mark:](./methods.go#L3775) | [:x:](./sdl_functions_js.go#L2988) |
| [SDL_FlushAudioStream](https://wiki.libsdl.org/SDL3/SDL_FlushAudioStream) | [:heavy_check_mark:](./methods.go#L3786) | [:x:](./sdl_functions_js.go#L3004) |
| [SDL_ClearAudioStream](https://wiki.libsdl.org/SDL3/SDL_ClearAudioStream) | [:heavy_check_mark:](./methods.go#L3796) | [:x:](./sdl_functions_js.go#L3020) |
| [SDL_PauseAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_PauseAudioStreamDevice) | [:heavy_check_mark:](./methods.go#L3806) | [:x:](./sdl_functions_js.go#L3036) |
| [SDL_ResumeAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_ResumeAudioStreamDevice) | [:heavy_check_mark:](./methods.go#L3816) | [:x:](./sdl_functions_js.go#L3052) |
| [SDL_AudioStreamDevicePaused](https://wiki.libsdl.org/SDL3/SDL_AudioStreamDevicePaused) | [:heavy_check_mark:](./methods.go#L3826) | [:x:](./sdl_functions_js.go#L3068) |
| [SDL_LockAudioStream](https://wiki.libsdl.org/SDL3/SDL_LockAudioStream) | [:heavy_check_mark:](./methods.go#L3832) | [:x:](./sdl_functions_js.go#L3084) |
| [SDL_UnlockAudioStream](https://wiki.libsdl.org/SDL3/SDL_UnlockAudioStream) | [:heavy_check_mark:](./methods.go#L3842) | [:x:](./sdl_functions_js.go#L3100) |
| [SDL_SetAudioStreamGetCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGetCallback) | [:x:](./methods.go#L3852) | [:x:](./sdl_functions_js.go#L3116) |
| [SDL_SetAudioStreamPutCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamPutCallback) | [:x:](./methods.go#L3859) | [:x:](./sdl_functions_js.go#L3136) |
| [SDL_DestroyAudioStream](https://wiki.libsdl.org/SDL3/SDL_DestroyAudioStream) | [:heavy_check_mark:](./methods.go#L3866) | [:x:](./sdl_functions_js.go#L3156) |
| [SDL_OpenAudioDeviceStream](https://wiki.libsdl.org/SDL3/SDL_OpenAudioDeviceStream) | [:x:](./methods.go#L360) | [:x:](./sdl_functions_js.go#L3170) |
| [SDL_SetAudioPostmixCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioPostmixCallback) | [:x:](./methods.go#L367) | [:x:](./sdl_functions_js.go#L3194) |
| [SDL_LoadWAV_IO](https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO) | [:heavy_check_mark:](./functions.go#L781) | [:x:](./sdl_functions_js.go#L3211) |
| [SDL_LoadWAV](https://wiki.libsdl.org/SDL3/SDL_LoadWAV) | [:heavy_check_mark:](./functions.go#L795) | [:x:](./sdl_functions_js.go#L3244) |
| [SDL_MixAudio](https://wiki.libsdl.org/SDL3/SDL_MixAudio) | [:heavy_check_mark:](./functions.go#L809) | [:x:](./sdl_functions_js.go#L3272) |
| [SDL_ConvertAudioSamples](https://wiki.libsdl.org/SDL3/SDL_ConvertAudioSamples) | [:heavy_check_mark:](./functions.go#L820) | [:x:](./sdl_functions_js.go#L3299) |
| [SDL_GetAudioFormatName](https://wiki.libsdl.org/SDL3/SDL_GetAudioFormatName) | [:heavy_check_mark:](./methods.go#L1203) | [:x:](./sdl_functions_js.go#L3337) |
| [SDL_GetSilenceValueForFormat](https://wiki.libsdl.org/SDL3/SDL_GetSilenceValueForFormat) | [:heavy_check_mark:](./methods.go#L1209) | [:x:](./sdl_functions_js.go#L3350) |
### Time

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetDateTimeLocalePreferences](https://wiki.libsdl.org/SDL3/SDL_GetDateTimeLocalePreferences) | [:x:](./methods.go#L3364) | [:x:](./sdl_functions_js.go#L16982) |
| [SDL_GetCurrentTime](https://wiki.libsdl.org/SDL3/SDL_GetCurrentTime) | [:question:]() | [:question:](./sdl_functions_js.go#L17003) |
| [SDL_TimeToDateTime](https://wiki.libsdl.org/SDL3/SDL_TimeToDateTime) | [:x:](./methods.go#L5584) | [:x:](./sdl_functions_js.go#L17019) |
| [SDL_DateTimeToTime](https://wiki.libsdl.org/SDL3/SDL_DateTimeToTime) | [:heavy_check_mark:](./methods.go#L5832) | [:x:](./sdl_functions_js.go#L17039) |
| [SDL_TimeToWindows](https://wiki.libsdl.org/SDL3/SDL_TimeToWindows) | [:x:](./methods.go#L5591) | [:x:](./sdl_functions_js.go#L17060) |
| [SDL_TimeFromWindows](https://wiki.libsdl.org/SDL3/SDL_TimeFromWindows) | [:question:]() | [:question:](./sdl_functions_js.go#L17081) |
| [SDL_GetDaysInMonth](https://wiki.libsdl.org/SDL3/SDL_GetDaysInMonth) | [:question:]() | [:question:](./sdl_functions_js.go#L17096) |
| [SDL_GetDayOfYear](https://wiki.libsdl.org/SDL3/SDL_GetDayOfYear) | [:question:]() | [:question:](./sdl_functions_js.go#L17111) |
| [SDL_GetDayOfWeek](https://wiki.libsdl.org/SDL3/SDL_GetDayOfWeek) | [:question:]() | [:question:](./sdl_functions_js.go#L17128) |
### Timer

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetTicks](https://wiki.libsdl.org/SDL3/SDL_GetTicks) | [:heavy_check_mark:](./functions.go#L842) | [:heavy_check_mark:](./sdl_functions_js.go#L17145) |
| [SDL_GetTicksNS](https://wiki.libsdl.org/SDL3/SDL_GetTicksNS) | [:heavy_check_mark:](./functions.go#L848) | [:x:](./sdl_functions_js.go#L17153) |
| [SDL_GetPerformanceCounter](https://wiki.libsdl.org/SDL3/SDL_GetPerformanceCounter) | [:heavy_check_mark:](./functions.go#L854) | [:x:](./sdl_functions_js.go#L17164) |
| [SDL_GetPerformanceFrequency](https://wiki.libsdl.org/SDL3/SDL_GetPerformanceFrequency) | [:heavy_check_mark:](./functions.go#L860) | [:x:](./sdl_functions_js.go#L17175) |
| [SDL_Delay](https://wiki.libsdl.org/SDL3/SDL_Delay) | [:heavy_check_mark:](./functions.go#L866) | [:x:](./sdl_functions_js.go#L17186) |
| [SDL_DelayNS](https://wiki.libsdl.org/SDL3/SDL_DelayNS) | [:heavy_check_mark:](./functions.go#L872) | [:x:](./sdl_functions_js.go#L17197) |
| [SDL_DelayPrecise](https://wiki.libsdl.org/SDL3/SDL_DelayPrecise) | [:heavy_check_mark:](./functions.go#L878) | [:x:](./sdl_functions_js.go#L17208) |
| [SDL_AddTimer](https://wiki.libsdl.org/SDL3/SDL_AddTimer) | [:question:]() | [:question:](./sdl_functions_js.go#L17219) |
| [SDL_AddTimerNS](https://wiki.libsdl.org/SDL3/SDL_AddTimerNS) | [:question:]() | [:question:](./sdl_functions_js.go#L17236) |
| [SDL_RemoveTimer](https://wiki.libsdl.org/SDL3/SDL_RemoveTimer) | [:x:](./methods.go#L4734) | [:x:](./sdl_functions_js.go#L17253) |
### Render

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumRenderDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumRenderDrivers) | [:heavy_check_mark:](./functions.go#L313) | [:x:](./sdl_functions_js.go#L14603) |
| [SDL_GetRenderDriver](https://wiki.libsdl.org/SDL3/SDL_GetRenderDriver) | [:heavy_check_mark:](./functions.go#L319) | [:x:](./sdl_functions_js.go#L14614) |
| [SDL_CreateWindowAndRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateWindowAndRenderer) | [:heavy_check_mark:](./functions.go#L325) | [:heavy_check_mark:](./sdl_functions_js.go#L14627) |
| [SDL_CreateRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateRenderer) | [:heavy_check_mark:](./methods.go#L4675) | [:x:](./sdl_functions_js.go#L14655) |
| [SDL_CreateRendererWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateRendererWithProperties) | [:heavy_check_mark:](./functions.go#L338) | [:x:](./sdl_functions_js.go#L14676) |
| [SDL_CreateSoftwareRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateSoftwareRenderer) | [:heavy_check_mark:](./methods.go#L1862) | [:x:](./sdl_functions_js.go#L14692) |
| [SDL_GetRenderer](https://wiki.libsdl.org/SDL3/SDL_GetRenderer) | [:heavy_check_mark:](./methods.go#L4686) | [:x:](./sdl_functions_js.go#L14711) |
| [SDL_GetRenderWindow](https://wiki.libsdl.org/SDL3/SDL_GetRenderWindow) | [:heavy_check_mark:](./methods.go#L2786) | [:x:](./sdl_functions_js.go#L14730) |
| [SDL_GetRendererName](https://wiki.libsdl.org/SDL3/SDL_GetRendererName) | [:heavy_check_mark:](./methods.go#L2797) | [:x:](./sdl_functions_js.go#L14749) |
| [SDL_GetRendererProperties](https://wiki.libsdl.org/SDL3/SDL_GetRendererProperties) | [:heavy_check_mark:](./methods.go#L2808) | [:x:](./sdl_functions_js.go#L14765) |
| [SDL_GetRenderOutputSize](https://wiki.libsdl.org/SDL3/SDL_GetRenderOutputSize) | [:heavy_check_mark:](./methods.go#L2814) | [:x:](./sdl_functions_js.go#L14781) |
| [SDL_GetCurrentRenderOutputSize](https://wiki.libsdl.org/SDL3/SDL_GetCurrentRenderOutputSize) | [:heavy_check_mark:](./methods.go#L2825) | [:x:](./sdl_functions_js.go#L14807) |
| [SDL_CreateTexture](https://wiki.libsdl.org/SDL3/SDL_CreateTexture) | [:heavy_check_mark:](./methods.go#L2836) | [:heavy_check_mark:](./sdl_functions_js.go#L14833) |
| [SDL_CreateTextureFromSurface](https://wiki.libsdl.org/SDL3/SDL_CreateTextureFromSurface) | [:heavy_check_mark:](./methods.go#L2847) | [:heavy_check_mark:](./sdl_functions_js.go#L14856) |
| [SDL_CreateTextureWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateTextureWithProperties) | [:heavy_check_mark:](./methods.go#L2858) | [:x:](./sdl_functions_js.go#L14876) |
| [SDL_GetTextureProperties](https://wiki.libsdl.org/SDL3/SDL_GetTextureProperties) | [:heavy_check_mark:](./methods.go#L939) | [:x:](./sdl_functions_js.go#L14897) |
| [SDL_GetRendererFromTexture](https://wiki.libsdl.org/SDL3/SDL_GetRendererFromTexture) | [:heavy_check_mark:](./methods.go#L950) | [:x:](./sdl_functions_js.go#L14913) |
| [SDL_GetTextureSize](https://wiki.libsdl.org/SDL3/SDL_GetTextureSize) | [:heavy_check_mark:](./methods.go#L961) | [:x:](./sdl_functions_js.go#L14932) |
| [SDL_SetTextureColorMod](https://wiki.libsdl.org/SDL3/SDL_SetTextureColorMod) | [:heavy_check_mark:](./methods.go#L972) | [:x:](./sdl_functions_js.go#L14958) |
| [SDL_SetTextureColorModFloat](https://wiki.libsdl.org/SDL3/SDL_SetTextureColorModFloat) | [:heavy_check_mark:](./methods.go#L982) | [:heavy_check_mark:](./sdl_functions_js.go#L14980) |
| [SDL_GetTextureColorMod](https://wiki.libsdl.org/SDL3/SDL_GetTextureColorMod) | [:heavy_check_mark:](./methods.go#L992) | [:x:](./sdl_functions_js.go#L14996) |
| [SDL_GetTextureColorModFloat](https://wiki.libsdl.org/SDL3/SDL_GetTextureColorModFloat) | [:heavy_check_mark:](./methods.go#L1003) | [:x:](./sdl_functions_js.go#L15027) |
| [SDL_SetTextureAlphaMod](https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaMod) | [:heavy_check_mark:](./methods.go#L1014) | [:x:](./sdl_functions_js.go#L15058) |
| [SDL_SetTextureAlphaModFloat](https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaModFloat) | [:heavy_check_mark:](./methods.go#L1024) | [:x:](./sdl_functions_js.go#L15076) |
| [SDL_GetTextureAlphaMod](https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaMod) | [:heavy_check_mark:](./methods.go#L1034) | [:x:](./sdl_functions_js.go#L15094) |
| [SDL_GetTextureAlphaModFloat](https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaModFloat) | [:heavy_check_mark:](./methods.go#L1045) | [:x:](./sdl_functions_js.go#L15115) |
| [SDL_SetTextureBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetTextureBlendMode) | [:heavy_check_mark:](./methods.go#L1056) | [:x:](./sdl_functions_js.go#L15136) |
| [SDL_GetTextureBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetTextureBlendMode) | [:heavy_check_mark:](./methods.go#L1066) | [:x:](./sdl_functions_js.go#L15154) |
| [SDL_SetTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_SetTextureScaleMode) | [:heavy_check_mark:](./methods.go#L1077) | [:x:](./sdl_functions_js.go#L15175) |
| [SDL_GetTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_GetTextureScaleMode) | [:heavy_check_mark:](./methods.go#L1087) | [:x:](./sdl_functions_js.go#L15193) |
| [SDL_UpdateTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateTexture) | [:heavy_check_mark:](./methods.go#L1098) | [:heavy_check_mark:](./sdl_functions_js.go#L15214) |
| [SDL_UpdateYUVTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateYUVTexture) | [:x:](./methods.go#L1110) | [:x:](./sdl_functions_js.go#L15238) |
| [SDL_UpdateNVTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateNVTexture) | [:x:](./methods.go#L1117) | [:x:](./sdl_functions_js.go#L15280) |
| [SDL_LockTexture](https://wiki.libsdl.org/SDL3/SDL_LockTexture) | [:heavy_check_mark:](./methods.go#L1125) | [:x:](./sdl_functions_js.go#L15315) |
| [SDL_LockTextureToSurface](https://wiki.libsdl.org/SDL3/SDL_LockTextureToSurface) | [:heavy_check_mark:](./methods.go#L1138) | [:heavy_check_mark:](./sdl_functions_js.go#L15346) |
| [SDL_UnlockTexture](https://wiki.libsdl.org/SDL3/SDL_UnlockTexture) | [:heavy_check_mark:](./methods.go#L1148) | [:heavy_check_mark:](./sdl_functions_js.go#L15368) |
| [SDL_SetRenderTarget](https://wiki.libsdl.org/SDL3/SDL_SetRenderTarget) | [:heavy_check_mark:](./methods.go#L2869) | [:x:](./sdl_functions_js.go#L15379) |
| [SDL_GetRenderTarget](https://wiki.libsdl.org/SDL3/SDL_GetRenderTarget) | [:heavy_check_mark:](./methods.go#L2879) | [:x:](./sdl_functions_js.go#L15400) |
| [SDL_SetRenderLogicalPresentation](https://wiki.libsdl.org/SDL3/SDL_SetRenderLogicalPresentation) | [:heavy_check_mark:](./methods.go#L2885) | [:x:](./sdl_functions_js.go#L15419) |
| [SDL_GetRenderLogicalPresentation](https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentation) | [:heavy_check_mark:](./methods.go#L2895) | [:x:](./sdl_functions_js.go#L15441) |
| [SDL_GetRenderLogicalPresentationRect](https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentationRect) | [:heavy_check_mark:](./methods.go#L2907) | [:x:](./sdl_functions_js.go#L15472) |
| [SDL_RenderCoordinatesFromWindow](https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesFromWindow) | [:x:](./methods.go#L2919) | [:x:](./sdl_functions_js.go#L15493) |
| [SDL_RenderCoordinatesToWindow](https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesToWindow) | [:x:](./methods.go#L2926) | [:x:](./sdl_functions_js.go#L15523) |
| [SDL_ConvertEventToRenderCoordinates](https://wiki.libsdl.org/SDL3/SDL_ConvertEventToRenderCoordinates) | [:x:](./methods.go#L2933) | [:x:](./sdl_functions_js.go#L15553) |
| [SDL_SetRenderViewport](https://wiki.libsdl.org/SDL3/SDL_SetRenderViewport) | [:heavy_check_mark:](./methods.go#L2940) | [:heavy_check_mark:](./sdl_functions_js.go#L15574) |
| [SDL_GetRenderViewport](https://wiki.libsdl.org/SDL3/SDL_GetRenderViewport) | [:heavy_check_mark:](./methods.go#L2950) | [:x:](./sdl_functions_js.go#L15591) |
| [SDL_RenderViewportSet](https://wiki.libsdl.org/SDL3/SDL_RenderViewportSet) | [:heavy_check_mark:](./methods.go#L2962) | [:x:](./sdl_functions_js.go#L15612) |
| [SDL_GetRenderSafeArea](https://wiki.libsdl.org/SDL3/SDL_GetRenderSafeArea) | [:heavy_check_mark:](./methods.go#L2968) | [:x:](./sdl_functions_js.go#L15628) |
| [SDL_SetRenderClipRect](https://wiki.libsdl.org/SDL3/SDL_SetRenderClipRect) | [:heavy_check_mark:](./methods.go#L2980) | [:heavy_check_mark:](./sdl_functions_js.go#L15649) |
| [SDL_GetRenderClipRect](https://wiki.libsdl.org/SDL3/SDL_GetRenderClipRect) | [:heavy_check_mark:](./methods.go#L2990) | [:x:](./sdl_functions_js.go#L15666) |
| [SDL_RenderClipEnabled](https://wiki.libsdl.org/SDL3/SDL_RenderClipEnabled) | [:heavy_check_mark:](./methods.go#L3002) | [:x:](./sdl_functions_js.go#L15687) |
| [SDL_SetRenderScale](https://wiki.libsdl.org/SDL3/SDL_SetRenderScale) | [:heavy_check_mark:](./methods.go#L3012) | [:heavy_check_mark:](./sdl_functions_js.go#L15703) |
| [SDL_GetRenderScale](https://wiki.libsdl.org/SDL3/SDL_GetRenderScale) | [:heavy_check_mark:](./methods.go#L3022) | [:x:](./sdl_functions_js.go#L15720) |
| [SDL_SetRenderDrawColor](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColor) | [:heavy_check_mark:](./methods.go#L3034) | [:heavy_check_mark:](./sdl_functions_js.go#L15746) |
| [SDL_SetRenderDrawColorFloat](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColorFloat) | [:heavy_check_mark:](./methods.go#L3044) | [:heavy_check_mark:](./sdl_functions_js.go#L15768) |
| [SDL_GetRenderDrawColor](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColor) | [:heavy_check_mark:](./methods.go#L3054) | [:x:](./sdl_functions_js.go#L15785) |
| [SDL_GetRenderDrawColorFloat](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColorFloat) | [:heavy_check_mark:](./methods.go#L3066) | [:x:](./sdl_functions_js.go#L15821) |
| [SDL_SetRenderColorScale](https://wiki.libsdl.org/SDL3/SDL_SetRenderColorScale) | [:heavy_check_mark:](./methods.go#L3078) | [:x:](./sdl_functions_js.go#L15857) |
| [SDL_GetRenderColorScale](https://wiki.libsdl.org/SDL3/SDL_GetRenderColorScale) | [:heavy_check_mark:](./methods.go#L3088) | [:x:](./sdl_functions_js.go#L15875) |
| [SDL_SetRenderDrawBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawBlendMode) | [:heavy_check_mark:](./methods.go#L3100) | [:x:](./sdl_functions_js.go#L15896) |
| [SDL_GetRenderDrawBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawBlendMode) | [:heavy_check_mark:](./methods.go#L3110) | [:x:](./sdl_functions_js.go#L15914) |
| [SDL_RenderClear](https://wiki.libsdl.org/SDL3/SDL_RenderClear) | [:heavy_check_mark:](./methods.go#L3122) | [:heavy_check_mark:](./sdl_functions_js.go#L15935) |
| [SDL_RenderPoint](https://wiki.libsdl.org/SDL3/SDL_RenderPoint) | [:x:](./methods.go#L3131) | [:x:](./sdl_functions_js.go#L15948) |
| [SDL_RenderPoints](https://wiki.libsdl.org/SDL3/SDL_RenderPoints) | [:heavy_check_mark:](./methods.go#L3138) | [:heavy_check_mark:](./sdl_functions_js.go#L15968) |
| [SDL_RenderLine](https://wiki.libsdl.org/SDL3/SDL_RenderLine) | [:heavy_check_mark:](./methods.go#L3148) | [:heavy_check_mark:](./sdl_functions_js.go#L15986) |
| [SDL_RenderLines](https://wiki.libsdl.org/SDL3/SDL_RenderLines) | [:heavy_check_mark:](./methods.go#L3158) | [:heavy_check_mark:](./sdl_functions_js.go#L16003) |
| [SDL_RenderRect](https://wiki.libsdl.org/SDL3/SDL_RenderRect) | [:heavy_check_mark:](./methods.go#L3168) | [:heavy_check_mark:](./sdl_functions_js.go#L16021) |
| [SDL_RenderRects](https://wiki.libsdl.org/SDL3/SDL_RenderRects) | [:heavy_check_mark:](./methods.go#L3178) | [:heavy_check_mark:](./sdl_functions_js.go#L16038) |
| [SDL_RenderFillRect](https://wiki.libsdl.org/SDL3/SDL_RenderFillRect) | [:heavy_check_mark:](./methods.go#L3188) | [:heavy_check_mark:](./sdl_functions_js.go#L16056) |
| [SDL_RenderFillRects](https://wiki.libsdl.org/SDL3/SDL_RenderFillRects) | [:heavy_check_mark:](./methods.go#L3198) | [:heavy_check_mark:](./sdl_functions_js.go#L16073) |
| [SDL_RenderTexture](https://wiki.libsdl.org/SDL3/SDL_RenderTexture) | [:heavy_check_mark:](./methods.go#L3208) | [:heavy_check_mark:](./sdl_functions_js.go#L16091) |
| [SDL_RenderTextureRotated](https://wiki.libsdl.org/SDL3/SDL_RenderTextureRotated) | [:heavy_check_mark:](./methods.go#L3218) | [:heavy_check_mark:](./sdl_functions_js.go#L16115) |
| [SDL_RenderTextureAffine](https://wiki.libsdl.org/SDL3/SDL_RenderTextureAffine) | [:x:](./methods.go#L3228) | [:x:](./sdl_functions_js.go#L16145) |
| [SDL_RenderTextureTiled](https://wiki.libsdl.org/SDL3/SDL_RenderTextureTiled) | [:x:](./methods.go#L3235) | [:x:](./sdl_functions_js.go#L16186) |
| [SDL_RenderTexture9Grid](https://wiki.libsdl.org/SDL3/SDL_RenderTexture9Grid) | [:x:](./methods.go#L3242) | [:x:](./sdl_functions_js.go#L16219) |
| [SDL_RenderTexture9GridTiled](https://wiki.libsdl.org/SDL3/SDL_RenderTexture9GridTiled) | [:question:]() | [:question:]() |
| [SDL_RenderGeometry](https://wiki.libsdl.org/SDL3/SDL_RenderGeometry) | [:heavy_check_mark:](./methods.go#L3249) | [:heavy_check_mark:](./sdl_functions_js.go#L16260) |
| [SDL_RenderGeometryRaw](https://wiki.libsdl.org/SDL3/SDL_RenderGeometryRaw) | [:x:](./methods.go#L3259) | [:x:](./sdl_functions_js.go#L16287) |
| [SDL_RenderReadPixels](https://wiki.libsdl.org/SDL3/SDL_RenderReadPixels) | [:heavy_check_mark:](./methods.go#L3266) | [:heavy_check_mark:](./sdl_functions_js.go#L16337) |
| [SDL_RenderPresent](https://wiki.libsdl.org/SDL3/SDL_RenderPresent) | [:heavy_check_mark:](./methods.go#L3277) | [:heavy_check_mark:](./sdl_functions_js.go#L16356) |
| [SDL_DestroyTexture](https://wiki.libsdl.org/SDL3/SDL_DestroyTexture) | [:heavy_check_mark:](./methods.go#L1154) | [:heavy_check_mark:](./sdl_functions_js.go#L16369) |
| [SDL_DestroyRenderer](https://wiki.libsdl.org/SDL3/SDL_DestroyRenderer) | [:heavy_check_mark:](./methods.go#L3287) | [:heavy_check_mark:](./sdl_functions_js.go#L16381) |
| [SDL_FlushRenderer](https://wiki.libsdl.org/SDL3/SDL_FlushRenderer) | [:heavy_check_mark:](./methods.go#L3293) | [:x:](./sdl_functions_js.go#L16393) |
| [SDL_GetRenderMetalLayer](https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalLayer) | [:x:](./methods.go#L3303) | [:x:](./sdl_functions_js.go#L16409) |
| [SDL_GetRenderMetalCommandEncoder](https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalCommandEncoder) | [:x:](./methods.go#L3310) | [:x:](./sdl_functions_js.go#L16425) |
| [SDL_AddVulkanRenderSemaphores](https://wiki.libsdl.org/SDL3/SDL_AddVulkanRenderSemaphores) | [:x:](./methods.go#L3317) | [:x:](./sdl_functions_js.go#L16441) |
| [SDL_SetRenderVSync](https://wiki.libsdl.org/SDL3/SDL_SetRenderVSync) | [:heavy_check_mark:](./methods.go#L3324) | [:x:](./sdl_functions_js.go#L16463) |
| [SDL_GetRenderVSync](https://wiki.libsdl.org/SDL3/SDL_GetRenderVSync) | [:heavy_check_mark:](./methods.go#L3334) | [:x:](./sdl_functions_js.go#L16481) |
| [SDL_RenderDebugText](https://wiki.libsdl.org/SDL3/SDL_RenderDebugText) | [:heavy_check_mark:](./methods.go#L3356) | [:heavy_check_mark:](./sdl_functions_js.go#L16502) |
| [SDL_RenderDebugTextFormat](https://wiki.libsdl.org/SDL3/SDL_RenderDebugTextFormat) | [:question:]() | [:question:](./sdl_functions_js.go#L16522) |
| [SDL_SetDefaultTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_SetDefaultTextureScaleMode) | [:question:]() | [:question:]() |
| [SDL_GetDefaultTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_GetDefaultTextureScaleMode) | [:question:]() | [:question:]() |
### SharedObject

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_LoadObject](https://wiki.libsdl.org/SDL3/SDL_LoadObject) | [:question:]() | [:question:](./sdl_functions_js.go#L14034) |
| [SDL_LoadFunction](https://wiki.libsdl.org/SDL3/SDL_LoadFunction) | [:x:](./methods.go#L4718) | [:x:](./sdl_functions_js.go#L14050) |
| [SDL_UnloadObject](https://wiki.libsdl.org/SDL3/SDL_UnloadObject) | [:x:](./methods.go#L4725) | [:x:](./sdl_functions_js.go#L14068) |
### Thread

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateThread](https://wiki.libsdl.org/SDL3/SDL_CreateThread) | [:question:]() | [:question:]() |
| [SDL_CreateThreadWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateThreadWithProperties) | [:question:]() | [:question:]() |
| [SDL_GetThreadName](https://wiki.libsdl.org/SDL3/SDL_GetThreadName) | [:x:](./methods.go#L485) | [:x:](./sdl_functions_js.go#L871) |
| [SDL_GetCurrentThreadID](https://wiki.libsdl.org/SDL3/SDL_GetCurrentThreadID) | [:question:]() | [:question:](./sdl_functions_js.go#L887) |
| [SDL_GetThreadID](https://wiki.libsdl.org/SDL3/SDL_GetThreadID) | [:x:](./methods.go#L492) | [:x:](./sdl_functions_js.go#L898) |
| [SDL_SetCurrentThreadPriority](https://wiki.libsdl.org/SDL3/SDL_SetCurrentThreadPriority) | [:x:](./methods.go#L23) | [:x:](./sdl_functions_js.go#L914) |
| [SDL_WaitThread](https://wiki.libsdl.org/SDL3/SDL_WaitThread) | [:x:](./methods.go#L499) | [:x:](./sdl_functions_js.go#L927) |
| [SDL_GetThreadState](https://wiki.libsdl.org/SDL3/SDL_GetThreadState) | [:x:](./methods.go#L506) | [:x:](./sdl_functions_js.go#L946) |
| [SDL_DetachThread](https://wiki.libsdl.org/SDL3/SDL_DetachThread) | [:x:](./methods.go#L513) | [:x:](./sdl_functions_js.go#L962) |
| [SDL_GetTLS](https://wiki.libsdl.org/SDL3/SDL_GetTLS) | [:x:](./methods.go#L610) | [:x:](./sdl_functions_js.go#L976) |
| [SDL_SetTLS](https://wiki.libsdl.org/SDL3/SDL_SetTLS) | [:x:](./methods.go#L617) | [:x:](./sdl_functions_js.go#L992) |
| [SDL_CleanupTLS](https://wiki.libsdl.org/SDL3/SDL_CleanupTLS) | [:question:]() | [:question:](./sdl_functions_js.go#L1012) |
### Mutex

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateMutex](https://wiki.libsdl.org/SDL3/SDL_CreateMutex) | [:question:]() | [:question:](./sdl_functions_js.go#L1021) |
| [SDL_LockMutex](https://wiki.libsdl.org/SDL3/SDL_LockMutex) | [:heavy_check_mark:](./methods.go#L5558) | [:x:](./sdl_functions_js.go#L1035) |
| [SDL_TryLockMutex](https://wiki.libsdl.org/SDL3/SDL_TryLockMutex) | [:heavy_check_mark:](./methods.go#L5564) | [:x:](./sdl_functions_js.go#L1049) |
| [SDL_UnlockMutex](https://wiki.libsdl.org/SDL3/SDL_UnlockMutex) | [:heavy_check_mark:](./methods.go#L5570) | [:x:](./sdl_functions_js.go#L1065) |
| [SDL_DestroyMutex](https://wiki.libsdl.org/SDL3/SDL_DestroyMutex) | [:heavy_check_mark:](./methods.go#L5576) | [:x:](./sdl_functions_js.go#L1079) |
| [SDL_CreateRWLock](https://wiki.libsdl.org/SDL3/SDL_CreateRWLock) | [:question:]() | [:question:](./sdl_functions_js.go#L1093) |
| [SDL_LockRWLockForReading](https://wiki.libsdl.org/SDL3/SDL_LockRWLockForReading) | [:heavy_check_mark:](./methods.go#L1165) | [:x:](./sdl_functions_js.go#L1107) |
| [SDL_LockRWLockForWriting](https://wiki.libsdl.org/SDL3/SDL_LockRWLockForWriting) | [:heavy_check_mark:](./methods.go#L1171) | [:x:](./sdl_functions_js.go#L1121) |
| [SDL_TryLockRWLockForReading](https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForReading) | [:heavy_check_mark:](./methods.go#L1177) | [:x:](./sdl_functions_js.go#L1135) |
| [SDL_TryLockRWLockForWriting](https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForWriting) | [:heavy_check_mark:](./methods.go#L1183) | [:x:](./sdl_functions_js.go#L1151) |
| [SDL_UnlockRWLock](https://wiki.libsdl.org/SDL3/SDL_UnlockRWLock) | [:heavy_check_mark:](./methods.go#L1189) | [:x:](./sdl_functions_js.go#L1167) |
| [SDL_DestroyRWLock](https://wiki.libsdl.org/SDL3/SDL_DestroyRWLock) | [:heavy_check_mark:](./methods.go#L1195) | [:x:](./sdl_functions_js.go#L1181) |
| [SDL_CreateSemaphore](https://wiki.libsdl.org/SDL3/SDL_CreateSemaphore) | [:question:]() | [:question:](./sdl_functions_js.go#L1195) |
| [SDL_DestroySemaphore](https://wiki.libsdl.org/SDL3/SDL_DestroySemaphore) | [:heavy_check_mark:](./methods.go#L5767) | [:x:](./sdl_functions_js.go#L1211) |
| [SDL_WaitSemaphore](https://wiki.libsdl.org/SDL3/SDL_WaitSemaphore) | [:heavy_check_mark:](./methods.go#L5773) | [:x:](./sdl_functions_js.go#L1225) |
| [SDL_TryWaitSemaphore](https://wiki.libsdl.org/SDL3/SDL_TryWaitSemaphore) | [:heavy_check_mark:](./methods.go#L5779) | [:x:](./sdl_functions_js.go#L1239) |
| [SDL_WaitSemaphoreTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitSemaphoreTimeout) | [:heavy_check_mark:](./methods.go#L5785) | [:x:](./sdl_functions_js.go#L1255) |
| [SDL_SignalSemaphore](https://wiki.libsdl.org/SDL3/SDL_SignalSemaphore) | [:heavy_check_mark:](./methods.go#L5791) | [:x:](./sdl_functions_js.go#L1273) |
| [SDL_GetSemaphoreValue](https://wiki.libsdl.org/SDL3/SDL_GetSemaphoreValue) | [:heavy_check_mark:](./methods.go#L5797) | [:x:](./sdl_functions_js.go#L1287) |
| [SDL_CreateCondition](https://wiki.libsdl.org/SDL3/SDL_CreateCondition) | [:question:]() | [:question:](./sdl_functions_js.go#L1303) |
| [SDL_DestroyCondition](https://wiki.libsdl.org/SDL3/SDL_DestroyCondition) | [:x:](./methods.go#L3913) | [:x:](./sdl_functions_js.go#L1317) |
| [SDL_SignalCondition](https://wiki.libsdl.org/SDL3/SDL_SignalCondition) | [:x:](./methods.go#L3920) | [:x:](./sdl_functions_js.go#L1331) |
| [SDL_BroadcastCondition](https://wiki.libsdl.org/SDL3/SDL_BroadcastCondition) | [:x:](./methods.go#L3927) | [:x:](./sdl_functions_js.go#L1345) |
| [SDL_WaitCondition](https://wiki.libsdl.org/SDL3/SDL_WaitCondition) | [:x:](./methods.go#L3934) | [:x:](./sdl_functions_js.go#L1359) |
| [SDL_WaitConditionTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitConditionTimeout) | [:x:](./methods.go#L3941) | [:x:](./sdl_functions_js.go#L1378) |
| [SDL_ShouldInit](https://wiki.libsdl.org/SDL3/SDL_ShouldInit) | [:x:](./methods.go#L3403) | [:x:](./sdl_functions_js.go#L1401) |
| [SDL_ShouldQuit](https://wiki.libsdl.org/SDL3/SDL_ShouldQuit) | [:x:](./methods.go#L3410) | [:x:](./sdl_functions_js.go#L1417) |
| [SDL_SetInitialized](https://wiki.libsdl.org/SDL3/SDL_SetInitialized) | [:x:](./methods.go#L3417) | [:x:](./sdl_functions_js.go#L1433) |
### Atomic

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_TryLockSpinlock](https://wiki.libsdl.org/SDL3/SDL_TryLockSpinlock) | [:heavy_check_mark:](./methods.go#L465) | [:x:](./sdl_functions_js.go#L246) |
| [SDL_LockSpinlock](https://wiki.libsdl.org/SDL3/SDL_LockSpinlock) | [:heavy_check_mark:](./methods.go#L471) | [:x:](./sdl_functions_js.go#L262) |
| [SDL_UnlockSpinlock](https://wiki.libsdl.org/SDL3/SDL_UnlockSpinlock) | [:heavy_check_mark:](./methods.go#L477) | [:x:](./sdl_functions_js.go#L276) |
| [SDL_MemoryBarrierReleaseFunction](https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierReleaseFunction) | [:question:]() | [:question:](./sdl_functions_js.go#L290) |
| [SDL_MemoryBarrierAcquireFunction](https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierAcquireFunction) | [:question:]() | [:question:](./sdl_functions_js.go#L299) |
| [SDL_CompareAndSwapAtomicInt](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicInt) | [:heavy_check_mark:](./methods.go#L869) | [:x:](./sdl_functions_js.go#L308) |
| [SDL_SetAtomicInt](https://wiki.libsdl.org/SDL3/SDL_SetAtomicInt) | [:heavy_check_mark:](./methods.go#L875) | [:x:](./sdl_functions_js.go#L328) |
| [SDL_GetAtomicInt](https://wiki.libsdl.org/SDL3/SDL_GetAtomicInt) | [:heavy_check_mark:](./methods.go#L881) | [:x:](./sdl_functions_js.go#L346) |
| [SDL_AddAtomicInt](https://wiki.libsdl.org/SDL3/SDL_AddAtomicInt) | [:heavy_check_mark:](./methods.go#L887) | [:x:](./sdl_functions_js.go#L362) |
| [SDL_CompareAndSwapAtomicU32](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicU32) | [:x:](./methods.go#L5600) | [:x:](./sdl_functions_js.go#L380) |
| [SDL_SetAtomicU32](https://wiki.libsdl.org/SDL3/SDL_SetAtomicU32) | [:x:](./methods.go#L5607) | [:x:](./sdl_functions_js.go#L400) |
| [SDL_GetAtomicU32](https://wiki.libsdl.org/SDL3/SDL_GetAtomicU32) | [:x:](./methods.go#L5614) | [:x:](./sdl_functions_js.go#L418) |
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
| [SDL_IOFromConstMem](https://wiki.libsdl.org/SDL3/SDL_IOFromConstMem) | [:heavy_check_mark:](./functions.go#L285) | [:heavy_check_mark:](./sdl_functions_js.go#L1485) |
| [SDL_IOFromDynamicMem](https://wiki.libsdl.org/SDL3/SDL_IOFromDynamicMem) | [:question:]() | [:question:](./sdl_functions_js.go#L1505) |
| [SDL_OpenIO](https://wiki.libsdl.org/SDL3/SDL_OpenIO) | [:x:](./methods.go#L32) | [:x:](./sdl_functions_js.go#L1519) |
| [SDL_CloseIO](https://wiki.libsdl.org/SDL3/SDL_CloseIO) | [:heavy_check_mark:](./methods.go#L4743) | [:heavy_check_mark:](./sdl_functions_js.go#L1540) |
| [SDL_GetIOProperties](https://wiki.libsdl.org/SDL3/SDL_GetIOProperties) | [:heavy_check_mark:](./methods.go#L4753) | [:x:](./sdl_functions_js.go#L1553) |
| [SDL_GetIOStatus](https://wiki.libsdl.org/SDL3/SDL_GetIOStatus) | [:heavy_check_mark:](./methods.go#L4764) | [:x:](./sdl_functions_js.go#L1569) |
| [SDL_GetIOSize](https://wiki.libsdl.org/SDL3/SDL_GetIOSize) | [:heavy_check_mark:](./methods.go#L4770) | [:x:](./sdl_functions_js.go#L1585) |
| [SDL_SeekIO](https://wiki.libsdl.org/SDL3/SDL_SeekIO) | [:heavy_check_mark:](./methods.go#L4781) | [:x:](./sdl_functions_js.go#L1601) |
| [SDL_TellIO](https://wiki.libsdl.org/SDL3/SDL_TellIO) | [:heavy_check_mark:](./methods.go#L4792) | [:x:](./sdl_functions_js.go#L1621) |
| [SDL_ReadIO](https://wiki.libsdl.org/SDL3/SDL_ReadIO) | [:heavy_check_mark:](./methods.go#L4798) | [:x:](./sdl_functions_js.go#L1637) |
| [SDL_WriteIO](https://wiki.libsdl.org/SDL3/SDL_WriteIO) | [:heavy_check_mark:](./methods.go#L4809) | [:x:](./sdl_functions_js.go#L1657) |
| [SDL_IOprintf](https://wiki.libsdl.org/SDL3/SDL_IOprintf) | [:heavy_check_mark:](./methods.go#L4820) | [:x:](./sdl_functions_js.go#L1677) |
| [SDL_IOvprintf](https://wiki.libsdl.org/SDL3/SDL_IOvprintf) | [:question:]() | [:question:](./sdl_functions_js.go#L1695) |
| [SDL_FlushIO](https://wiki.libsdl.org/SDL3/SDL_FlushIO) | [:heavy_check_mark:](./methods.go#L4831) | [:x:](./sdl_functions_js.go#L1715) |
| [SDL_LoadFile_IO](https://wiki.libsdl.org/SDL3/SDL_LoadFile_IO) | [:heavy_check_mark:](./methods.go#L4841) | [:x:](./sdl_functions_js.go#L1731) |
| [SDL_LoadFile](https://wiki.libsdl.org/SDL3/SDL_LoadFile) | [:question:]() | [:question:](./sdl_functions_js.go#L1754) |
| [SDL_SaveFile_IO](https://wiki.libsdl.org/SDL3/SDL_SaveFile_IO) | [:heavy_check_mark:](./methods.go#L4854) | [:x:](./sdl_functions_js.go#L1772) |
| [SDL_SaveFile](https://wiki.libsdl.org/SDL3/SDL_SaveFile) | [:question:]() | [:question:](./sdl_functions_js.go#L1794) |
| [SDL_ReadU8](https://wiki.libsdl.org/SDL3/SDL_ReadU8) | [:heavy_check_mark:](./methods.go#L4865) | [:x:](./sdl_functions_js.go#L1811) |
| [SDL_ReadS8](https://wiki.libsdl.org/SDL3/SDL_ReadS8) | [:heavy_check_mark:](./methods.go#L4877) | [:x:](./sdl_functions_js.go#L1832) |
| [SDL_ReadU16LE](https://wiki.libsdl.org/SDL3/SDL_ReadU16LE) | [:heavy_check_mark:](./methods.go#L4889) | [:x:](./sdl_functions_js.go#L1853) |
| [SDL_ReadS16LE](https://wiki.libsdl.org/SDL3/SDL_ReadS16LE) | [:heavy_check_mark:](./methods.go#L4901) | [:x:](./sdl_functions_js.go#L1874) |
| [SDL_ReadU16BE](https://wiki.libsdl.org/SDL3/SDL_ReadU16BE) | [:heavy_check_mark:](./methods.go#L4913) | [:x:](./sdl_functions_js.go#L1895) |
| [SDL_ReadS16BE](https://wiki.libsdl.org/SDL3/SDL_ReadS16BE) | [:heavy_check_mark:](./methods.go#L4925) | [:x:](./sdl_functions_js.go#L1916) |
| [SDL_ReadU32LE](https://wiki.libsdl.org/SDL3/SDL_ReadU32LE) | [:heavy_check_mark:](./methods.go#L4937) | [:x:](./sdl_functions_js.go#L1937) |
| [SDL_ReadS32LE](https://wiki.libsdl.org/SDL3/SDL_ReadS32LE) | [:heavy_check_mark:](./methods.go#L4949) | [:x:](./sdl_functions_js.go#L1958) |
| [SDL_ReadU32BE](https://wiki.libsdl.org/SDL3/SDL_ReadU32BE) | [:heavy_check_mark:](./methods.go#L4961) | [:x:](./sdl_functions_js.go#L1979) |
| [SDL_ReadS32BE](https://wiki.libsdl.org/SDL3/SDL_ReadS32BE) | [:heavy_check_mark:](./methods.go#L4973) | [:x:](./sdl_functions_js.go#L2000) |
| [SDL_ReadU64LE](https://wiki.libsdl.org/SDL3/SDL_ReadU64LE) | [:heavy_check_mark:](./methods.go#L4985) | [:x:](./sdl_functions_js.go#L2021) |
| [SDL_ReadS64LE](https://wiki.libsdl.org/SDL3/SDL_ReadS64LE) | [:heavy_check_mark:](./methods.go#L4997) | [:x:](./sdl_functions_js.go#L2042) |
| [SDL_ReadU64BE](https://wiki.libsdl.org/SDL3/SDL_ReadU64BE) | [:heavy_check_mark:](./methods.go#L5009) | [:x:](./sdl_functions_js.go#L2063) |
| [SDL_ReadS64BE](https://wiki.libsdl.org/SDL3/SDL_ReadS64BE) | [:heavy_check_mark:](./methods.go#L5021) | [:x:](./sdl_functions_js.go#L2084) |
| [SDL_WriteU8](https://wiki.libsdl.org/SDL3/SDL_WriteU8) | [:heavy_check_mark:](./methods.go#L5033) | [:x:](./sdl_functions_js.go#L2105) |
| [SDL_WriteS8](https://wiki.libsdl.org/SDL3/SDL_WriteS8) | [:heavy_check_mark:](./methods.go#L5043) | [:x:](./sdl_functions_js.go#L2123) |
| [SDL_WriteU16LE](https://wiki.libsdl.org/SDL3/SDL_WriteU16LE) | [:heavy_check_mark:](./methods.go#L5053) | [:x:](./sdl_functions_js.go#L2141) |
| [SDL_WriteS16LE](https://wiki.libsdl.org/SDL3/SDL_WriteS16LE) | [:heavy_check_mark:](./methods.go#L5063) | [:x:](./sdl_functions_js.go#L2159) |
| [SDL_WriteU16BE](https://wiki.libsdl.org/SDL3/SDL_WriteU16BE) | [:heavy_check_mark:](./methods.go#L5073) | [:x:](./sdl_functions_js.go#L2177) |
| [SDL_WriteS16BE](https://wiki.libsdl.org/SDL3/SDL_WriteS16BE) | [:heavy_check_mark:](./methods.go#L5083) | [:x:](./sdl_functions_js.go#L2195) |
| [SDL_WriteU32LE](https://wiki.libsdl.org/SDL3/SDL_WriteU32LE) | [:heavy_check_mark:](./methods.go#L5093) | [:x:](./sdl_functions_js.go#L2213) |
| [SDL_WriteS32LE](https://wiki.libsdl.org/SDL3/SDL_WriteS32LE) | [:heavy_check_mark:](./methods.go#L5103) | [:x:](./sdl_functions_js.go#L2231) |
| [SDL_WriteU32BE](https://wiki.libsdl.org/SDL3/SDL_WriteU32BE) | [:heavy_check_mark:](./methods.go#L5113) | [:x:](./sdl_functions_js.go#L2249) |
| [SDL_WriteS32BE](https://wiki.libsdl.org/SDL3/SDL_WriteS32BE) | [:heavy_check_mark:](./methods.go#L5123) | [:x:](./sdl_functions_js.go#L2267) |
| [SDL_WriteU64LE](https://wiki.libsdl.org/SDL3/SDL_WriteU64LE) | [:heavy_check_mark:](./methods.go#L5133) | [:x:](./sdl_functions_js.go#L2285) |
| [SDL_WriteS64LE](https://wiki.libsdl.org/SDL3/SDL_WriteS64LE) | [:heavy_check_mark:](./methods.go#L5143) | [:x:](./sdl_functions_js.go#L2303) |
| [SDL_WriteU64BE](https://wiki.libsdl.org/SDL3/SDL_WriteU64BE) | [:heavy_check_mark:](./methods.go#L5153) | [:x:](./sdl_functions_js.go#L2321) |
| [SDL_WriteS64BE](https://wiki.libsdl.org/SDL3/SDL_WriteS64BE) | [:heavy_check_mark:](./methods.go#L5163) | [:x:](./sdl_functions_js.go#L2339) |
### AsyncIO

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AsyncIOFromFile](https://wiki.libsdl.org/SDL3/SDL_AsyncIOFromFile) | [:question:]() | [:question:](./sdl_functions_js.go#L23) |
| [SDL_GetAsyncIOSize](https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOSize) | [:x:](./methods.go#L3373) | [:x:](./sdl_functions_js.go#L41) |
| [SDL_ReadAsyncIO](https://wiki.libsdl.org/SDL3/SDL_ReadAsyncIO) | [:x:](./methods.go#L3380) | [:x:](./sdl_functions_js.go#L57) |
| [SDL_WriteAsyncIO](https://wiki.libsdl.org/SDL3/SDL_WriteAsyncIO) | [:x:](./methods.go#L3387) | [:x:](./sdl_functions_js.go#L86) |
| [SDL_CloseAsyncIO](https://wiki.libsdl.org/SDL3/SDL_CloseAsyncIO) | [:x:](./methods.go#L3394) | [:x:](./sdl_functions_js.go#L115) |
| [SDL_CreateAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_CreateAsyncIOQueue) | [:question:]() | [:question:](./sdl_functions_js.go#L140) |
| [SDL_DestroyAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_DestroyAsyncIOQueue) | [:heavy_check_mark:](./methods.go#L1321) | [:x:](./sdl_functions_js.go#L154) |
| [SDL_GetAsyncIOResult](https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOResult) | [:heavy_check_mark:](./methods.go#L1327) | [:x:](./sdl_functions_js.go#L168) |
| [SDL_WaitAsyncIOResult](https://wiki.libsdl.org/SDL3/SDL_WaitAsyncIOResult) | [:heavy_check_mark:](./methods.go#L1337) | [:x:](./sdl_functions_js.go#L189) |
| [SDL_SignalAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_SignalAsyncIOQueue) | [:heavy_check_mark:](./methods.go#L1347) | [:x:](./sdl_functions_js.go#L212) |
| [SDL_LoadFileAsync](https://wiki.libsdl.org/SDL3/SDL_LoadFileAsync) | [:question:]() | [:question:](./sdl_functions_js.go#L226) |
### Storage

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_OpenTitleStorage](https://wiki.libsdl.org/SDL3/SDL_OpenTitleStorage) | [:question:]() | [:question:](./sdl_functions_js.go#L16544) |
| [SDL_OpenUserStorage](https://wiki.libsdl.org/SDL3/SDL_OpenUserStorage) | [:question:]() | [:question:](./sdl_functions_js.go#L16562) |
| [SDL_OpenFileStorage](https://wiki.libsdl.org/SDL3/SDL_OpenFileStorage) | [:question:]() | [:question:](./sdl_functions_js.go#L16582) |
| [SDL_OpenStorage](https://wiki.libsdl.org/SDL3/SDL_OpenStorage) | [:x:](./methods.go#L5549) | [:x:](./sdl_functions_js.go#L16598) |
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
| [SDL_GlobStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_GlobStorageDirectory) | [:heavy_check_mark:](./methods.go#L211) | [:x:](./sdl_functions_js.go#L16855) |
### Pixels

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPixelFormatName](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatName) | [:heavy_check_mark:](./methods.go#L5805) | [:x:](./sdl_functions_js.go#L3386) |
| [SDL_GetMasksForPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetMasksForPixelFormat) | [:x:](./methods.go#L5811) | [:x:](./sdl_functions_js.go#L3399) |
| [SDL_GetPixelFormatForMasks](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatForMasks) | [:heavy_check_mark:](./functions.go#L351) | [:x:](./sdl_functions_js.go#L3437) |
| [SDL_GetPixelFormatDetails](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatDetails) | [:heavy_check_mark:](./methods.go#L5819) | [:heavy_check_mark:](./sdl_functions_js.go#L3458) |
| [SDL_CreatePalette](https://wiki.libsdl.org/SDL3/SDL_CreatePalette) | [:heavy_check_mark:](./functions.go#L357) | [:x:](./sdl_functions_js.go#L3470) |
| [SDL_SetPaletteColors](https://wiki.libsdl.org/SDL3/SDL_SetPaletteColors) | [:heavy_check_mark:](./methods.go#L5175) | [:x:](./sdl_functions_js.go#L3486) |
| [SDL_DestroyPalette](https://wiki.libsdl.org/SDL3/SDL_DestroyPalette) | [:heavy_check_mark:](./methods.go#L5185) | [:x:](./sdl_functions_js.go#L3511) |
| [SDL_MapRGB](https://wiki.libsdl.org/SDL3/SDL_MapRGB) | [:heavy_check_mark:](./methods.go#L1355) | [:heavy_check_mark:](./sdl_functions_js.go#L3525) |
| [SDL_MapRGBA](https://wiki.libsdl.org/SDL3/SDL_MapRGBA) | [:heavy_check_mark:](./methods.go#L1361) | [:x:](./sdl_functions_js.go#L3552) |
| [SDL_GetRGB](https://wiki.libsdl.org/SDL3/SDL_GetRGB) | [:heavy_check_mark:](./functions.go#L380) | [:x:](./sdl_functions_js.go#L3581) |
| [SDL_GetRGBA](https://wiki.libsdl.org/SDL3/SDL_GetRGBA) | [:heavy_check_mark:](./functions.go#L388) | [:x:](./sdl_functions_js.go#L3617) |
### Surface

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateSurface](https://wiki.libsdl.org/SDL3/SDL_CreateSurface) | [:heavy_check_mark:](./functions.go#L409) | [:x:](./sdl_functions_js.go#L3932) |
| [SDL_CreateSurfaceFrom](https://wiki.libsdl.org/SDL3/SDL_CreateSurfaceFrom) | [:question:]() | [:question:](./sdl_functions_js.go#L3952) |
| [SDL_DestroySurface](https://wiki.libsdl.org/SDL3/SDL_DestroySurface) | [:heavy_check_mark:](./methods.go#L1369) | [:heavy_check_mark:](./sdl_functions_js.go#L3976) |
| [SDL_GetSurfaceProperties](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceProperties) | [:heavy_check_mark:](./methods.go#L1378) | [:x:](./sdl_functions_js.go#L3988) |
| [SDL_SetSurfaceColorspace](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorspace) | [:heavy_check_mark:](./methods.go#L1389) | [:x:](./sdl_functions_js.go#L4004) |
| [SDL_GetSurfaceColorspace](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorspace) | [:heavy_check_mark:](./methods.go#L1399) | [:x:](./sdl_functions_js.go#L4022) |
| [SDL_CreateSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_CreateSurfacePalette) | [:heavy_check_mark:](./methods.go#L1405) | [:x:](./sdl_functions_js.go#L4038) |
| [SDL_SetSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_SetSurfacePalette) | [:heavy_check_mark:](./methods.go#L1416) | [:x:](./sdl_functions_js.go#L4057) |
| [SDL_GetSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_GetSurfacePalette) | [:heavy_check_mark:](./methods.go#L1426) | [:x:](./sdl_functions_js.go#L4078) |
| [SDL_AddSurfaceAlternateImage](https://wiki.libsdl.org/SDL3/SDL_AddSurfaceAlternateImage) | [:heavy_check_mark:](./methods.go#L1432) | [:x:](./sdl_functions_js.go#L4097) |
| [SDL_SurfaceHasAlternateImages](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasAlternateImages) | [:heavy_check_mark:](./methods.go#L1442) | [:x:](./sdl_functions_js.go#L4118) |
| [SDL_GetSurfaceImages](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceImages) | [:heavy_check_mark:](./methods.go#L1448) | [:x:](./sdl_functions_js.go#L4134) |
| [SDL_RemoveSurfaceAlternateImages](https://wiki.libsdl.org/SDL3/SDL_RemoveSurfaceAlternateImages) | [:heavy_check_mark:](./methods.go#L1462) | [:x:](./sdl_functions_js.go#L4155) |
| [SDL_LockSurface](https://wiki.libsdl.org/SDL3/SDL_LockSurface) | [:heavy_check_mark:](./methods.go#L1468) | [:x:](./sdl_functions_js.go#L4169) |
| [SDL_UnlockSurface](https://wiki.libsdl.org/SDL3/SDL_UnlockSurface) | [:heavy_check_mark:](./methods.go#L1478) | [:x:](./sdl_functions_js.go#L4185) |
| [SDL_LoadBMP_IO](https://wiki.libsdl.org/SDL3/SDL_LoadBMP_IO) | [:heavy_check_mark:](./functions.go#L420) | [:heavy_check_mark:](./sdl_functions_js.go#L4199) |
| [SDL_LoadBMP](https://wiki.libsdl.org/SDL3/SDL_LoadBMP) | [:heavy_check_mark:](./functions.go#L431) | [:x:](./sdl_functions_js.go#L4219) |
| [SDL_SaveBMP_IO](https://wiki.libsdl.org/SDL3/SDL_SaveBMP_IO) | [:heavy_check_mark:](./methods.go#L1484) | [:x:](./sdl_functions_js.go#L4235) |
| [SDL_SaveBMP](https://wiki.libsdl.org/SDL3/SDL_SaveBMP) | [:heavy_check_mark:](./methods.go#L1494) | [:x:](./sdl_functions_js.go#L4258) |
| [SDL_SetSurfaceRLE](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceRLE) | [:heavy_check_mark:](./methods.go#L1504) | [:x:](./sdl_functions_js.go#L4276) |
| [SDL_SurfaceHasRLE](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasRLE) | [:heavy_check_mark:](./methods.go#L1514) | [:x:](./sdl_functions_js.go#L4294) |
| [SDL_SetSurfaceColorKey](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorKey) | [:heavy_check_mark:](./methods.go#L1520) | [:x:](./sdl_functions_js.go#L4310) |
| [SDL_SurfaceHasColorKey](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasColorKey) | [:heavy_check_mark:](./methods.go#L1530) | [:x:](./sdl_functions_js.go#L4330) |
| [SDL_GetSurfaceColorKey](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorKey) | [:heavy_check_mark:](./methods.go#L1536) | [:x:](./sdl_functions_js.go#L4346) |
| [SDL_SetSurfaceColorMod](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorMod) | [:heavy_check_mark:](./methods.go#L1547) | [:x:](./sdl_functions_js.go#L4367) |
| [SDL_GetSurfaceColorMod](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorMod) | [:heavy_check_mark:](./methods.go#L1557) | [:x:](./sdl_functions_js.go#L4389) |
| [SDL_SetSurfaceAlphaMod](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceAlphaMod) | [:heavy_check_mark:](./methods.go#L1569) | [:x:](./sdl_functions_js.go#L4420) |
| [SDL_GetSurfaceAlphaMod](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceAlphaMod) | [:heavy_check_mark:](./methods.go#L1579) | [:x:](./sdl_functions_js.go#L4438) |
| [SDL_SetSurfaceBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceBlendMode) | [:heavy_check_mark:](./methods.go#L1591) | [:x:](./sdl_functions_js.go#L4459) |
| [SDL_GetSurfaceBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceBlendMode) | [:heavy_check_mark:](./methods.go#L1601) | [:x:](./sdl_functions_js.go#L4477) |
| [SDL_SetSurfaceClipRect](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceClipRect) | [:heavy_check_mark:](./methods.go#L1613) | [:x:](./sdl_functions_js.go#L4498) |
| [SDL_GetSurfaceClipRect](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceClipRect) | [:heavy_check_mark:](./methods.go#L1619) | [:x:](./sdl_functions_js.go#L4519) |
| [SDL_FlipSurface](https://wiki.libsdl.org/SDL3/SDL_FlipSurface) | [:heavy_check_mark:](./methods.go#L1631) | [:x:](./sdl_functions_js.go#L4540) |
| [SDL_DuplicateSurface](https://wiki.libsdl.org/SDL3/SDL_DuplicateSurface) | [:heavy_check_mark:](./methods.go#L1641) | [:x:](./sdl_functions_js.go#L4558) |
| [SDL_ScaleSurface](https://wiki.libsdl.org/SDL3/SDL_ScaleSurface) | [:heavy_check_mark:](./methods.go#L1652) | [:x:](./sdl_functions_js.go#L4577) |
| [SDL_ConvertSurface](https://wiki.libsdl.org/SDL3/SDL_ConvertSurface) | [:heavy_check_mark:](./methods.go#L1663) | [:heavy_check_mark:](./sdl_functions_js.go#L4602) |
| [SDL_ConvertSurfaceAndColorspace](https://wiki.libsdl.org/SDL3/SDL_ConvertSurfaceAndColorspace) | [:heavy_check_mark:](./methods.go#L1674) | [:x:](./sdl_functions_js.go#L4619) |
| [SDL_ConvertPixels](https://wiki.libsdl.org/SDL3/SDL_ConvertPixels) | [:question:]() | [:question:](./sdl_functions_js.go#L4649) |
| [SDL_ConvertPixelsAndColorspace](https://wiki.libsdl.org/SDL3/SDL_ConvertPixelsAndColorspace) | [:question:]() | [:question:](./sdl_functions_js.go#L4676) |
| [SDL_PremultiplyAlpha](https://wiki.libsdl.org/SDL3/SDL_PremultiplyAlpha) | [:question:]() | [:question:](./sdl_functions_js.go#L4711) |
| [SDL_PremultiplySurfaceAlpha](https://wiki.libsdl.org/SDL3/SDL_PremultiplySurfaceAlpha) | [:heavy_check_mark:](./methods.go#L1685) | [:x:](./sdl_functions_js.go#L4740) |
| [SDL_ClearSurface](https://wiki.libsdl.org/SDL3/SDL_ClearSurface) | [:heavy_check_mark:](./methods.go#L1695) | [:x:](./sdl_functions_js.go#L4758) |
| [SDL_FillSurfaceRect](https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRect) | [:heavy_check_mark:](./methods.go#L1705) | [:heavy_check_mark:](./sdl_functions_js.go#L4782) |
| [SDL_FillSurfaceRects](https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRects) | [:heavy_check_mark:](./methods.go#L1715) | [:x:](./sdl_functions_js.go#L4801) |
| [SDL_BlitSurface](https://wiki.libsdl.org/SDL3/SDL_BlitSurface) | [:heavy_check_mark:](./methods.go#L1725) | [:x:](./sdl_functions_js.go#L4826) |
| [SDL_BlitSurfaceUnchecked](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUnchecked) | [:heavy_check_mark:](./methods.go#L1735) | [:x:](./sdl_functions_js.go#L4857) |
| [SDL_BlitSurfaceScaled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceScaled) | [:heavy_check_mark:](./methods.go#L1745) | [:x:](./sdl_functions_js.go#L4888) |
| [SDL_BlitSurfaceUncheckedScaled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUncheckedScaled) | [:heavy_check_mark:](./methods.go#L1755) | [:x:](./sdl_functions_js.go#L4921) |
| [SDL_StretchSurface](https://wiki.libsdl.org/SDL3/SDL_StretchSurface) | [:question:]() | [:question:]() |
| [SDL_BlitSurfaceTiled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiled) | [:heavy_check_mark:](./methods.go#L1765) | [:x:](./sdl_functions_js.go#L4954) |
| [SDL_BlitSurfaceTiledWithScale](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiledWithScale) | [:heavy_check_mark:](./methods.go#L1775) | [:x:](./sdl_functions_js.go#L4985) |
| [SDL_BlitSurface9Grid](https://wiki.libsdl.org/SDL3/SDL_BlitSurface9Grid) | [:heavy_check_mark:](./methods.go#L1785) | [:x:](./sdl_functions_js.go#L5020) |
| [SDL_MapSurfaceRGB](https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGB) | [:heavy_check_mark:](./methods.go#L1795) | [:x:](./sdl_functions_js.go#L5063) |
| [SDL_MapSurfaceRGBA](https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGBA) | [:heavy_check_mark:](./methods.go#L1801) | [:x:](./sdl_functions_js.go#L5085) |
| [SDL_ReadSurfacePixel](https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixel) | [:heavy_check_mark:](./methods.go#L1807) | [:heavy_check_mark:](./sdl_functions_js.go#L5109) |
| [SDL_ReadSurfacePixelFloat](https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixelFloat) | [:heavy_check_mark:](./methods.go#L1819) | [:x:](./sdl_functions_js.go#L5140) |
| [SDL_WriteSurfacePixel](https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixel) | [:heavy_check_mark:](./methods.go#L1831) | [:x:](./sdl_functions_js.go#L5180) |
| [SDL_WriteSurfacePixelFloat](https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixelFloat) | [:heavy_check_mark:](./methods.go#L1841) | [:x:](./sdl_functions_js.go#L5208) |
### BlendMode

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ComposeCustomBlendMode](https://wiki.libsdl.org/SDL3/SDL_ComposeCustomBlendMode) | [:heavy_check_mark:](./functions.go#L448) | [:x:](./sdl_functions_js.go#L3363) |
### Rect

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_RectToFRect](https://wiki.libsdl.org/SDL3/SDL_RectToFRect) | [:question:]() | [:question:]() |
| [SDL_PointInRect](https://wiki.libsdl.org/SDL3/SDL_PointInRect) | [:question:]() | [:question:]() |
| [SDL_RectEmpty](https://wiki.libsdl.org/SDL3/SDL_RectEmpty) | [:question:]() | [:question:]() |
| [SDL_RectsEqual](https://wiki.libsdl.org/SDL3/SDL_RectsEqual) | [:question:]() | [:question:]() |
| [SDL_HasRectIntersection](https://wiki.libsdl.org/SDL3/SDL_HasRectIntersection) | [:heavy_check_mark:](./methods.go#L626) | [:x:](./sdl_functions_js.go#L3658) |
| [SDL_GetRectIntersection](https://wiki.libsdl.org/SDL3/SDL_GetRectIntersection) | [:heavy_check_mark:](./methods.go#L632) | [:x:](./sdl_functions_js.go#L3679) |
| [SDL_GetRectUnion](https://wiki.libsdl.org/SDL3/SDL_GetRectUnion) | [:heavy_check_mark:](./methods.go#L644) | [:x:](./sdl_functions_js.go#L3705) |
| [SDL_GetRectEnclosingPoints](https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPoints) | [:heavy_check_mark:](./functions.go#L1394) | [:x:](./sdl_functions_js.go#L3731) |
| [SDL_GetRectAndLineIntersection](https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersection) | [:question:]() | [:question:](./sdl_functions_js.go#L3759) |
| [SDL_PointInRectFloat](https://wiki.libsdl.org/SDL3/SDL_PointInRectFloat) | [:question:]() | [:question:]() |
| [SDL_RectEmptyFloat](https://wiki.libsdl.org/SDL3/SDL_RectEmptyFloat) | [:question:]() | [:question:]() |
| [SDL_RectsEqualEpsilon](https://wiki.libsdl.org/SDL3/SDL_RectsEqualEpsilon) | [:question:]() | [:question:]() |
| [SDL_RectsEqualFloat](https://wiki.libsdl.org/SDL3/SDL_RectsEqualFloat) | [:question:]() | [:question:]() |
| [SDL_HasRectIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_HasRectIntersectionFloat) | [:x:](./methods.go#L3874) | [:x:](./sdl_functions_js.go#L3795) |
| [SDL_GetRectIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectIntersectionFloat) | [:x:](./methods.go#L3881) | [:x:](./sdl_functions_js.go#L3816) |
| [SDL_GetRectUnionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectUnionFloat) | [:x:](./methods.go#L3888) | [:x:](./sdl_functions_js.go#L3842) |
| [SDL_GetRectEnclosingPointsFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPointsFloat) | [:question:]() | [:question:](./sdl_functions_js.go#L3868) |
| [SDL_GetRectAndLineIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersectionFloat) | [:x:](./methods.go#L3895) | [:x:](./sdl_functions_js.go#L3896) |
### Camera

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumCameraDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumCameraDrivers) | [:heavy_check_mark:](./functions.go#L1406) | [:x:](./sdl_functions_js.go#L5236) |
| [SDL_GetCameraDriver](https://wiki.libsdl.org/SDL3/SDL_GetCameraDriver) | [:heavy_check_mark:](./functions.go#L1412) | [:x:](./sdl_functions_js.go#L5247) |
| [SDL_GetCurrentCameraDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentCameraDriver) | [:heavy_check_mark:](./functions.go#L1418) | [:x:](./sdl_functions_js.go#L5260) |
| [SDL_GetCameras](https://wiki.libsdl.org/SDL3/SDL_GetCameras) | [:heavy_check_mark:](./functions.go#L1424) | [:x:](./sdl_functions_js.go#L5271) |
| [SDL_GetCameraSupportedFormats](https://wiki.libsdl.org/SDL3/SDL_GetCameraSupportedFormats) | [:x:](./methods.go#L3426) | [:x:](./sdl_functions_js.go#L5287) |
| [SDL_GetCameraName](https://wiki.libsdl.org/SDL3/SDL_GetCameraName) | [:x:](./methods.go#L3433) | [:x:](./sdl_functions_js.go#L5305) |
| [SDL_GetCameraPosition](https://wiki.libsdl.org/SDL3/SDL_GetCameraPosition) | [:x:](./methods.go#L3440) | [:x:](./sdl_functions_js.go#L5318) |
| [SDL_OpenCamera](https://wiki.libsdl.org/SDL3/SDL_OpenCamera) | [:x:](./methods.go#L3447) | [:x:](./sdl_functions_js.go#L5331) |
| [SDL_GetCameraPermissionState](https://wiki.libsdl.org/SDL3/SDL_GetCameraPermissionState) | [:heavy_check_mark:](./methods.go#L376) | [:x:](./sdl_functions_js.go#L5352) |
| [SDL_GetCameraID](https://wiki.libsdl.org/SDL3/SDL_GetCameraID) | [:heavy_check_mark:](./methods.go#L382) | [:x:](./sdl_functions_js.go#L5368) |
| [SDL_GetCameraProperties](https://wiki.libsdl.org/SDL3/SDL_GetCameraProperties) | [:heavy_check_mark:](./methods.go#L393) | [:x:](./sdl_functions_js.go#L5384) |
| [SDL_GetCameraFormat](https://wiki.libsdl.org/SDL3/SDL_GetCameraFormat) | [:heavy_check_mark:](./methods.go#L404) | [:x:](./sdl_functions_js.go#L5400) |
| [SDL_AcquireCameraFrame](https://wiki.libsdl.org/SDL3/SDL_AcquireCameraFrame) | [:heavy_check_mark:](./methods.go#L416) | [:x:](./sdl_functions_js.go#L5421) |
| [SDL_ReleaseCameraFrame](https://wiki.libsdl.org/SDL3/SDL_ReleaseCameraFrame) | [:heavy_check_mark:](./methods.go#L429) | [:x:](./sdl_functions_js.go#L5445) |
| [SDL_CloseCamera](https://wiki.libsdl.org/SDL3/SDL_CloseCamera) | [:heavy_check_mark:](./methods.go#L435) | [:x:](./sdl_functions_js.go#L5464) |
### Clipboard

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetClipboardText](https://wiki.libsdl.org/SDL3/SDL_SetClipboardText) | [:heavy_check_mark:](./functions.go#L1440) | [:x:](./sdl_functions_js.go#L5478) |
| [SDL_GetClipboardText](https://wiki.libsdl.org/SDL3/SDL_GetClipboardText) | [:heavy_check_mark:](./functions.go#L1450) | [:x:](./sdl_functions_js.go#L5491) |
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
| [SDL_ShowFileDialogWithProperties](https://wiki.libsdl.org/SDL3/SDL_ShowFileDialogWithProperties) | [:x:](./methods.go#L2384) | [:x:](./sdl_functions_js.go#L7768) |
### GPU

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GPUSupportsShaderFormats](https://wiki.libsdl.org/SDL3/SDL_GPUSupportsShaderFormats) | [:heavy_check_mark:](./functions.go#L456) | [:x:](./sdl_functions_js.go#L11384) |
| [SDL_GPUSupportsProperties](https://wiki.libsdl.org/SDL3/SDL_GPUSupportsProperties) | [:heavy_check_mark:](./methods.go#L5759) | [:x:](./sdl_functions_js.go#L11399) |
| [SDL_CreateGPUDevice](https://wiki.libsdl.org/SDL3/SDL_CreateGPUDevice) | [:heavy_check_mark:](./functions.go#L462) | [:x:](./sdl_functions_js.go#L11412) |
| [SDL_CreateGPUDeviceWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateGPUDeviceWithProperties) | [:heavy_check_mark:](./functions.go#L473) | [:x:](./sdl_functions_js.go#L11432) |
| [SDL_DestroyGPUDevice](https://wiki.libsdl.org/SDL3/SDL_DestroyGPUDevice) | [:heavy_check_mark:](./methods.go#L1883) | [:x:](./sdl_functions_js.go#L11448) |
| [SDL_GetNumGPUDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumGPUDrivers) | [:question:]() | [:question:](./sdl_functions_js.go#L11462) |
| [SDL_GetGPUDriver](https://wiki.libsdl.org/SDL3/SDL_GetGPUDriver) | [:question:]() | [:question:](./sdl_functions_js.go#L11473) |
| [SDL_GetGPUDeviceDriver](https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceDriver) | [:heavy_check_mark:](./methods.go#L1889) | [:x:](./sdl_functions_js.go#L11486) |
| [SDL_GetGPUShaderFormats](https://wiki.libsdl.org/SDL3/SDL_GetGPUShaderFormats) | [:heavy_check_mark:](./methods.go#L1895) | [:x:](./sdl_functions_js.go#L11502) |
| [SDL_CreateGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_CreateGPUComputePipeline) | [:heavy_check_mark:](./methods.go#L1901) | [:x:](./sdl_functions_js.go#L11518) |
| [SDL_CreateGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_CreateGPUGraphicsPipeline) | [:heavy_check_mark:](./methods.go#L1912) | [:x:](./sdl_functions_js.go#L11542) |
| [SDL_CreateGPUSampler](https://wiki.libsdl.org/SDL3/SDL_CreateGPUSampler) | [:heavy_check_mark:](./methods.go#L1923) | [:x:](./sdl_functions_js.go#L11566) |
| [SDL_CreateGPUShader](https://wiki.libsdl.org/SDL3/SDL_CreateGPUShader) | [:heavy_check_mark:](./methods.go#L1929) | [:x:](./sdl_functions_js.go#L11590) |
| [SDL_CreateGPUTexture](https://wiki.libsdl.org/SDL3/SDL_CreateGPUTexture) | [:heavy_check_mark:](./methods.go#L1940) | [:x:](./sdl_functions_js.go#L11614) |
| [SDL_CreateGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_CreateGPUBuffer) | [:heavy_check_mark:](./methods.go#L1951) | [:x:](./sdl_functions_js.go#L11638) |
| [SDL_CreateGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_CreateGPUTransferBuffer) | [:heavy_check_mark:](./methods.go#L1962) | [:x:](./sdl_functions_js.go#L11662) |
| [SDL_SetGPUBufferName](https://wiki.libsdl.org/SDL3/SDL_SetGPUBufferName) | [:heavy_check_mark:](./methods.go#L1973) | [:x:](./sdl_functions_js.go#L11686) |
| [SDL_SetGPUTextureName](https://wiki.libsdl.org/SDL3/SDL_SetGPUTextureName) | [:heavy_check_mark:](./methods.go#L1979) | [:x:](./sdl_functions_js.go#L11707) |
| [SDL_InsertGPUDebugLabel](https://wiki.libsdl.org/SDL3/SDL_InsertGPUDebugLabel) | [:heavy_check_mark:](./methods.go#L5860) | [:x:](./sdl_functions_js.go#L11728) |
| [SDL_PushGPUDebugGroup](https://wiki.libsdl.org/SDL3/SDL_PushGPUDebugGroup) | [:heavy_check_mark:](./methods.go#L5866) | [:x:](./sdl_functions_js.go#L11744) |
| [SDL_PopGPUDebugGroup](https://wiki.libsdl.org/SDL3/SDL_PopGPUDebugGroup) | [:heavy_check_mark:](./methods.go#L5872) | [:x:](./sdl_functions_js.go#L11760) |
| [SDL_ReleaseGPUTexture](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTexture) | [:heavy_check_mark:](./methods.go#L1985) | [:x:](./sdl_functions_js.go#L11774) |
| [SDL_ReleaseGPUSampler](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUSampler) | [:heavy_check_mark:](./methods.go#L1991) | [:x:](./sdl_functions_js.go#L11793) |
| [SDL_ReleaseGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUBuffer) | [:heavy_check_mark:](./methods.go#L1997) | [:x:](./sdl_functions_js.go#L11812) |
| [SDL_ReleaseGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTransferBuffer) | [:heavy_check_mark:](./methods.go#L2003) | [:x:](./sdl_functions_js.go#L11831) |
| [SDL_ReleaseGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUComputePipeline) | [:heavy_check_mark:](./methods.go#L2009) | [:x:](./sdl_functions_js.go#L11850) |
| [SDL_ReleaseGPUShader](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUShader) | [:heavy_check_mark:](./methods.go#L2015) | [:x:](./sdl_functions_js.go#L11869) |
| [SDL_ReleaseGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUGraphicsPipeline) | [:heavy_check_mark:](./methods.go#L2021) | [:x:](./sdl_functions_js.go#L11888) |
| [SDL_AcquireGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_AcquireGPUCommandBuffer) | [:heavy_check_mark:](./methods.go#L2027) | [:x:](./sdl_functions_js.go#L11907) |
| [SDL_PushGPUVertexUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUVertexUniformData) | [:heavy_check_mark:](./methods.go#L5878) | [:x:](./sdl_functions_js.go#L11926) |
| [SDL_PushGPUFragmentUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUFragmentUniformData) | [:heavy_check_mark:](./methods.go#L5884) | [:x:](./sdl_functions_js.go#L11946) |
| [SDL_PushGPUComputeUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUComputeUniformData) | [:heavy_check_mark:](./methods.go#L5890) | [:x:](./sdl_functions_js.go#L11966) |
| [SDL_BeginGPURenderPass](https://wiki.libsdl.org/SDL3/SDL_BeginGPURenderPass) | [:heavy_check_mark:](./methods.go#L5896) | [:x:](./sdl_functions_js.go#L11986) |
| [SDL_BindGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_BindGPUGraphicsPipeline) | [:heavy_check_mark:](./methods.go#L1217) | [:x:](./sdl_functions_js.go#L12017) |
| [SDL_SetGPUViewport](https://wiki.libsdl.org/SDL3/SDL_SetGPUViewport) | [:heavy_check_mark:](./methods.go#L1223) | [:x:](./sdl_functions_js.go#L12036) |
| [SDL_SetGPUScissor](https://wiki.libsdl.org/SDL3/SDL_SetGPUScissor) | [:heavy_check_mark:](./methods.go#L1229) | [:x:](./sdl_functions_js.go#L12055) |
| [SDL_SetGPUBlendConstants](https://wiki.libsdl.org/SDL3/SDL_SetGPUBlendConstants) | [:question:]() | [:question:]() |
| [SDL_SetGPUStencilReference](https://wiki.libsdl.org/SDL3/SDL_SetGPUStencilReference) | [:heavy_check_mark:](./methods.go#L1235) | [:x:](./sdl_functions_js.go#L12074) |
| [SDL_BindGPUVertexBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexBuffers) | [:heavy_check_mark:](./methods.go#L1241) | [:x:](./sdl_functions_js.go#L12090) |
| [SDL_BindGPUIndexBuffer](https://wiki.libsdl.org/SDL3/SDL_BindGPUIndexBuffer) | [:heavy_check_mark:](./methods.go#L1247) | [:x:](./sdl_functions_js.go#L12113) |
| [SDL_BindGPUVertexSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexSamplers) | [:heavy_check_mark:](./methods.go#L1253) | [:x:](./sdl_functions_js.go#L12134) |
| [SDL_BindGPUVertexStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageTextures) | [:heavy_check_mark:](./methods.go#L1259) | [:x:](./sdl_functions_js.go#L12157) |
| [SDL_BindGPUVertexStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageBuffers) | [:heavy_check_mark:](./methods.go#L1265) | [:x:](./sdl_functions_js.go#L12180) |
| [SDL_BindGPUFragmentSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentSamplers) | [:heavy_check_mark:](./methods.go#L1271) | [:x:](./sdl_functions_js.go#L12203) |
| [SDL_BindGPUFragmentStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageTextures) | [:heavy_check_mark:](./methods.go#L1277) | [:x:](./sdl_functions_js.go#L12226) |
| [SDL_BindGPUFragmentStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageBuffers) | [:heavy_check_mark:](./methods.go#L1283) | [:x:](./sdl_functions_js.go#L12249) |
| [SDL_DrawGPUIndexedPrimitives](https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitives) | [:heavy_check_mark:](./methods.go#L1289) | [:x:](./sdl_functions_js.go#L12272) |
| [SDL_DrawGPUPrimitives](https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitives) | [:heavy_check_mark:](./methods.go#L1295) | [:x:](./sdl_functions_js.go#L12296) |
| [SDL_DrawGPUPrimitivesIndirect](https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitivesIndirect) | [:heavy_check_mark:](./methods.go#L1301) | [:x:](./sdl_functions_js.go#L12318) |
| [SDL_DrawGPUIndexedPrimitivesIndirect](https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitivesIndirect) | [:heavy_check_mark:](./methods.go#L1307) | [:x:](./sdl_functions_js.go#L12341) |
| [SDL_EndGPURenderPass](https://wiki.libsdl.org/SDL3/SDL_EndGPURenderPass) | [:heavy_check_mark:](./methods.go#L1313) | [:x:](./sdl_functions_js.go#L12364) |
| [SDL_BeginGPUComputePass](https://wiki.libsdl.org/SDL3/SDL_BeginGPUComputePass) | [:heavy_check_mark:](./methods.go#L5902) | [:x:](./sdl_functions_js.go#L12378) |
| [SDL_BindGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputePipeline) | [:heavy_check_mark:](./methods.go#L895) | [:x:](./sdl_functions_js.go#L12411) |
| [SDL_BindGPUComputeSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeSamplers) | [:heavy_check_mark:](./methods.go#L901) | [:x:](./sdl_functions_js.go#L12430) |
| [SDL_BindGPUComputeStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageTextures) | [:heavy_check_mark:](./methods.go#L907) | [:x:](./sdl_functions_js.go#L12453) |
| [SDL_BindGPUComputeStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageBuffers) | [:heavy_check_mark:](./methods.go#L913) | [:x:](./sdl_functions_js.go#L12476) |
| [SDL_DispatchGPUCompute](https://wiki.libsdl.org/SDL3/SDL_DispatchGPUCompute) | [:heavy_check_mark:](./methods.go#L919) | [:x:](./sdl_functions_js.go#L12499) |
| [SDL_DispatchGPUComputeIndirect](https://wiki.libsdl.org/SDL3/SDL_DispatchGPUComputeIndirect) | [:heavy_check_mark:](./methods.go#L925) | [:x:](./sdl_functions_js.go#L12519) |
| [SDL_EndGPUComputePass](https://wiki.libsdl.org/SDL3/SDL_EndGPUComputePass) | [:heavy_check_mark:](./methods.go#L931) | [:x:](./sdl_functions_js.go#L12540) |
| [SDL_MapGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_MapGPUTransferBuffer) | [:heavy_check_mark:](./methods.go#L2038) | [:x:](./sdl_functions_js.go#L12554) |
| [SDL_UnmapGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_UnmapGPUTransferBuffer) | [:heavy_check_mark:](./methods.go#L2049) | [:x:](./sdl_functions_js.go#L12577) |
| [SDL_BeginGPUCopyPass](https://wiki.libsdl.org/SDL3/SDL_BeginGPUCopyPass) | [:heavy_check_mark:](./methods.go#L5908) | [:x:](./sdl_functions_js.go#L12596) |
| [SDL_UploadToGPUTexture](https://wiki.libsdl.org/SDL3/SDL_UploadToGPUTexture) | [:x:](./methods.go#L2712) | [:x:](./sdl_functions_js.go#L12615) |
| [SDL_UploadToGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_UploadToGPUBuffer) | [:x:](./methods.go#L2719) | [:x:](./sdl_functions_js.go#L12641) |
| [SDL_CopyGPUTextureToTexture](https://wiki.libsdl.org/SDL3/SDL_CopyGPUTextureToTexture) | [:x:](./methods.go#L2726) | [:x:](./sdl_functions_js.go#L12667) |
| [SDL_CopyGPUBufferToBuffer](https://wiki.libsdl.org/SDL3/SDL_CopyGPUBufferToBuffer) | [:x:](./methods.go#L2733) | [:x:](./sdl_functions_js.go#L12699) |
| [SDL_DownloadFromGPUTexture](https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUTexture) | [:x:](./methods.go#L2740) | [:x:](./sdl_functions_js.go#L12727) |
| [SDL_DownloadFromGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUBuffer) | [:x:](./methods.go#L2747) | [:x:](./sdl_functions_js.go#L12751) |
| [SDL_EndGPUCopyPass](https://wiki.libsdl.org/SDL3/SDL_EndGPUCopyPass) | [:x:](./methods.go#L2754) | [:x:](./sdl_functions_js.go#L12775) |
| [SDL_GenerateMipmapsForGPUTexture](https://wiki.libsdl.org/SDL3/SDL_GenerateMipmapsForGPUTexture) | [:heavy_check_mark:](./methods.go#L5914) | [:x:](./sdl_functions_js.go#L12789) |
| [SDL_BlitGPUTexture](https://wiki.libsdl.org/SDL3/SDL_BlitGPUTexture) | [:heavy_check_mark:](./methods.go#L5920) | [:x:](./sdl_functions_js.go#L12808) |
| [SDL_WindowSupportsGPUSwapchainComposition](https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUSwapchainComposition) | [:heavy_check_mark:](./methods.go#L2055) | [:x:](./sdl_functions_js.go#L12827) |
| [SDL_WindowSupportsGPUPresentMode](https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUPresentMode) | [:heavy_check_mark:](./methods.go#L2061) | [:x:](./sdl_functions_js.go#L12850) |
| [SDL_ClaimWindowForGPUDevice](https://wiki.libsdl.org/SDL3/SDL_ClaimWindowForGPUDevice) | [:heavy_check_mark:](./methods.go#L2067) | [:x:](./sdl_functions_js.go#L12873) |
| [SDL_ReleaseWindowFromGPUDevice](https://wiki.libsdl.org/SDL3/SDL_ReleaseWindowFromGPUDevice) | [:heavy_check_mark:](./methods.go#L2077) | [:x:](./sdl_functions_js.go#L12894) |
| [SDL_SetGPUSwapchainParameters](https://wiki.libsdl.org/SDL3/SDL_SetGPUSwapchainParameters) | [:heavy_check_mark:](./methods.go#L2083) | [:x:](./sdl_functions_js.go#L12913) |
| [SDL_SetGPUAllowedFramesInFlight](https://wiki.libsdl.org/SDL3/SDL_SetGPUAllowedFramesInFlight) | [:heavy_check_mark:](./methods.go#L2093) | [:x:](./sdl_functions_js.go#L12938) |
| [SDL_GetGPUSwapchainTextureFormat](https://wiki.libsdl.org/SDL3/SDL_GetGPUSwapchainTextureFormat) | [:heavy_check_mark:](./methods.go#L2103) | [:x:](./sdl_functions_js.go#L12956) |
| [SDL_AcquireGPUSwapchainTexture](https://wiki.libsdl.org/SDL3/SDL_AcquireGPUSwapchainTexture) | [:heavy_check_mark:](./methods.go#L5926) | [:x:](./sdl_functions_js.go#L12977) |
| [SDL_WaitForGPUSwapchain](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUSwapchain) | [:heavy_check_mark:](./methods.go#L2109) | [:x:](./sdl_functions_js.go#L13013) |
| [SDL_WaitAndAcquireGPUSwapchainTexture](https://wiki.libsdl.org/SDL3/SDL_WaitAndAcquireGPUSwapchainTexture) | [:heavy_check_mark:](./methods.go#L5938) | [:x:](./sdl_functions_js.go#L13034) |
| [SDL_SubmitGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBuffer) | [:heavy_check_mark:](./methods.go#L5950) | [:x:](./sdl_functions_js.go#L13070) |
| [SDL_SubmitGPUCommandBufferAndAcquireFence](https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBufferAndAcquireFence) | [:heavy_check_mark:](./methods.go#L5960) | [:x:](./sdl_functions_js.go#L13086) |
| [SDL_CancelGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_CancelGPUCommandBuffer) | [:heavy_check_mark:](./methods.go#L5971) | [:x:](./sdl_functions_js.go#L13105) |
| [SDL_WaitForGPUIdle](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUIdle) | [:heavy_check_mark:](./methods.go#L2119) | [:x:](./sdl_functions_js.go#L13121) |
| [SDL_WaitForGPUFences](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUFences) | [:heavy_check_mark:](./methods.go#L2129) | [:x:](./sdl_functions_js.go#L13137) |
| [SDL_QueryGPUFence](https://wiki.libsdl.org/SDL3/SDL_QueryGPUFence) | [:heavy_check_mark:](./methods.go#L2139) | [:x:](./sdl_functions_js.go#L13162) |
| [SDL_ReleaseGPUFence](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUFence) | [:heavy_check_mark:](./methods.go#L2145) | [:x:](./sdl_functions_js.go#L13183) |
| [SDL_GPUTextureFormatTexelBlockSize](https://wiki.libsdl.org/SDL3/SDL_GPUTextureFormatTexelBlockSize) | [:heavy_check_mark:](./methods.go#L451) | [:x:](./sdl_functions_js.go#L13202) |
| [SDL_GPUTextureSupportsFormat](https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsFormat) | [:heavy_check_mark:](./methods.go#L2151) | [:x:](./sdl_functions_js.go#L13215) |
| [SDL_GPUTextureSupportsSampleCount](https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsSampleCount) | [:heavy_check_mark:](./methods.go#L2157) | [:x:](./sdl_functions_js.go#L13237) |
| [SDL_CalculateGPUTextureFormatSize](https://wiki.libsdl.org/SDL3/SDL_CalculateGPUTextureFormatSize) | [:heavy_check_mark:](./methods.go#L457) | [:x:](./sdl_functions_js.go#L13257) |
| [SDL_GDKSuspendGPU](https://wiki.libsdl.org/SDL3/SDL_GDKSuspendGPU) | [:question:]() | [:question:]() |
| [SDL_GDKResumeGPU](https://wiki.libsdl.org/SDL3/SDL_GDKResumeGPU) | [:question:]() | [:question:]() |
### MessageBox

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ShowMessageBox](https://wiki.libsdl.org/SDL3/SDL_ShowMessageBox) | [:question:]() | [:question:](./sdl_functions_js.go#L14336) |
| [SDL_ShowSimpleMessageBox](https://wiki.libsdl.org/SDL3/SDL_ShowSimpleMessageBox) | [:heavy_check_mark:](./functions.go#L905) | [:x:](./sdl_functions_js.go#L14357) |
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
| [SDL_Metal_CreateView](https://wiki.libsdl.org/SDL3/SDL_Metal_CreateView) | [:x:](./methods.go#L4668) | [:x:](./sdl_functions_js.go#L14379) |
| [SDL_Metal_DestroyView](https://wiki.libsdl.org/SDL3/SDL_Metal_DestroyView) | [:question:]() | [:question:](./sdl_functions_js.go#L14395) |
| [SDL_Metal_GetLayer](https://wiki.libsdl.org/SDL3/SDL_Metal_GetLayer) | [:question:]() | [:question:](./sdl_functions_js.go#L14406) |
### Platform

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetPowerInfo) | [:heavy_check_mark:](./functions.go#L923) | [:x:](./sdl_functions_js.go#L7813) |
### Power

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetSensors](https://wiki.libsdl.org/SDL3/SDL_GetSensors) | [:heavy_check_mark:](./functions.go#L938) | [:x:](./sdl_functions_js.go#L7834) |
| [SDL_GetSensorNameForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorNameForID) | [:heavy_check_mark:](./methods.go#L6062) | [:x:](./sdl_functions_js.go#L7850) |
| [SDL_GetSensorTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorTypeForID) | [:heavy_check_mark:](./methods.go#L6068) | [:x:](./sdl_functions_js.go#L7863) |
| [SDL_GetSensorNonPortableTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableTypeForID) | [:heavy_check_mark:](./methods.go#L6074) | [:x:](./sdl_functions_js.go#L7876) |
| [SDL_OpenSensor](https://wiki.libsdl.org/SDL3/SDL_OpenSensor) | [:heavy_check_mark:](./methods.go#L6080) | [:x:](./sdl_functions_js.go#L7889) |
| [SDL_GetSensorFromID](https://wiki.libsdl.org/SDL3/SDL_GetSensorFromID) | [:heavy_check_mark:](./methods.go#L6091) | [:x:](./sdl_functions_js.go#L7905) |
| [SDL_GetSensorProperties](https://wiki.libsdl.org/SDL3/SDL_GetSensorProperties) | [:heavy_check_mark:](./methods.go#L522) | [:x:](./sdl_functions_js.go#L7921) |
| [SDL_GetSensorName](https://wiki.libsdl.org/SDL3/SDL_GetSensorName) | [:heavy_check_mark:](./methods.go#L533) | [:x:](./sdl_functions_js.go#L7937) |
| [SDL_GetSensorType](https://wiki.libsdl.org/SDL3/SDL_GetSensorType) | [:heavy_check_mark:](./methods.go#L544) | [:x:](./sdl_functions_js.go#L7953) |
| [SDL_GetSensorNonPortableType](https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableType) | [:heavy_check_mark:](./methods.go#L550) | [:x:](./sdl_functions_js.go#L7969) |
| [SDL_GetSensorID](https://wiki.libsdl.org/SDL3/SDL_GetSensorID) | [:heavy_check_mark:](./methods.go#L556) | [:x:](./sdl_functions_js.go#L7985) |
| [SDL_GetSensorData](https://wiki.libsdl.org/SDL3/SDL_GetSensorData) | [:heavy_check_mark:](./methods.go#L567) | [:x:](./sdl_functions_js.go#L8001) |
| [SDL_CloseSensor](https://wiki.libsdl.org/SDL3/SDL_CloseSensor) | [:heavy_check_mark:](./methods.go#L577) | [:x:](./sdl_functions_js.go#L8024) |
| [SDL_UpdateSensors](https://wiki.libsdl.org/SDL3/SDL_UpdateSensors) | [:heavy_check_mark:](./functions.go#L952) | [:x:](./sdl_functions_js.go#L8038) |
### Sensor

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateProcess](https://wiki.libsdl.org/SDL3/SDL_CreateProcess) | [:question:]() | [:question:](./sdl_functions_js.go#L14432) |
| [SDL_CreateProcessWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateProcessWithProperties) | [:heavy_check_mark:](./functions.go#L960) | [:x:](./sdl_functions_js.go#L14452) |
| [SDL_GetProcessProperties](https://wiki.libsdl.org/SDL3/SDL_GetProcessProperties) | [:heavy_check_mark:](./methods.go#L5983) | [:x:](./sdl_functions_js.go#L14468) |
| [SDL_ReadProcess](https://wiki.libsdl.org/SDL3/SDL_ReadProcess) | [:heavy_check_mark:](./methods.go#L5994) | [:x:](./sdl_functions_js.go#L14484) |
| [SDL_GetProcessInput](https://wiki.libsdl.org/SDL3/SDL_GetProcessInput) | [:heavy_check_mark:](./methods.go#L6012) | [:x:](./sdl_functions_js.go#L14510) |
| [SDL_GetProcessOutput](https://wiki.libsdl.org/SDL3/SDL_GetProcessOutput) | [:heavy_check_mark:](./methods.go#L6023) | [:x:](./sdl_functions_js.go#L14529) |
| [SDL_KillProcess](https://wiki.libsdl.org/SDL3/SDL_KillProcess) | [:heavy_check_mark:](./methods.go#L6034) | [:x:](./sdl_functions_js.go#L14548) |
| [SDL_WaitProcess](https://wiki.libsdl.org/SDL3/SDL_WaitProcess) | [:heavy_check_mark:](./methods.go#L6044) | [:x:](./sdl_functions_js.go#L14566) |
| [SDL_DestroyProcess](https://wiki.libsdl.org/SDL3/SDL_DestroyProcess) | [:heavy_check_mark:](./methods.go#L6054) | [:x:](./sdl_functions_js.go#L14589) |
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
| [SDL_GetPreferredLocales](https://wiki.libsdl.org/SDL3/SDL_GetPreferredLocales) | [:heavy_check_mark:](./functions.go#L986) | [:x:](./sdl_functions_js.go#L14082) |
### Intrinsics

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetWindowsMessageHook](https://wiki.libsdl.org/SDL3/SDL_SetWindowsMessageHook) | [:question:]() | [:question:]() |
| [SDL_GetDirect3D9AdapterIndex](https://wiki.libsdl.org/SDL3/SDL_GetDirect3D9AdapterIndex) | [:question:]() | [:question:]() |
| [SDL_GetDXGIOutputInfo](https://wiki.libsdl.org/SDL3/SDL_GetDXGIOutputInfo) | [:question:]() | [:question:]() |
| [SDL_SetX11EventHook](https://wiki.libsdl.org/SDL3/SDL_SetX11EventHook) | [:x:](./methods.go#L601) | [:x:](./sdl_functions_js.go#L16882) |
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
| [TTF_Version](https://wiki.libsdl.org/SDL3_ttf/TTF_Version) | [:heavy_check_mark:](./ttf/functions.go#L11) | [:x:](./ttf/ttf_functions_js.go#L13) |
| [TTF_GetFreeTypeVersion](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFreeTypeVersion) | [:heavy_check_mark:](./ttf/functions.go#L17) | [:x:](./ttf/ttf_functions_js.go#L24) |
| [TTF_GetHarfBuzzVersion](https://wiki.libsdl.org/SDL3_ttf/TTF_GetHarfBuzzVersion) | [:heavy_check_mark:](./ttf/functions.go#L27) | [:x:](./ttf/ttf_functions_js.go#L48) |
| [TTF_Init](https://wiki.libsdl.org/SDL3_ttf/TTF_Init) | [:heavy_check_mark:](./ttf/functions.go#L37) | [:x:](./ttf/ttf_functions_js.go#L72) |
| [TTF_OpenFont](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFont) | [:heavy_check_mark:](./ttf/functions.go#L47) | [:x:](./ttf/ttf_functions_js.go#L83) |
| [TTF_OpenFontIO](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontIO) | [:heavy_check_mark:](./ttf/functions.go#L58) | [:x:](./ttf/ttf_functions_js.go#L99) |
| [TTF_OpenFontWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontWithProperties) | [:heavy_check_mark:](./ttf/functions.go#L69) | [:x:](./ttf/ttf_functions_js.go#L120) |
| [TTF_CopyFont](https://wiki.libsdl.org/SDL3_ttf/TTF_CopyFont) | [:heavy_check_mark:](./ttf/methods.go#L409) | [:x:](./ttf/ttf_functions_js.go#L137) |
| [TTF_GetFontProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontProperties) | [:heavy_check_mark:](./ttf/methods.go#L420) | [:x:](./ttf/ttf_functions_js.go#L154) |
| [TTF_GetFontGeneration](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontGeneration) | [:heavy_check_mark:](./ttf/methods.go#L426) | [:x:](./ttf/ttf_functions_js.go#L170) |
| [TTF_AddFallbackFont](https://wiki.libsdl.org/SDL3_ttf/TTF_AddFallbackFont) | [:heavy_check_mark:](./ttf/methods.go#L432) | [:x:](./ttf/ttf_functions_js.go#L186) |
| [TTF_RemoveFallbackFont](https://wiki.libsdl.org/SDL3_ttf/TTF_RemoveFallbackFont) | [:heavy_check_mark:](./ttf/methods.go#L442) | [:x:](./ttf/ttf_functions_js.go#L207) |
| [TTF_ClearFallbackFonts](https://wiki.libsdl.org/SDL3_ttf/TTF_ClearFallbackFonts) | [:heavy_check_mark:](./ttf/methods.go#L448) | [:x:](./ttf/ttf_functions_js.go#L226) |
| [TTF_SetFontSize](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSize) | [:heavy_check_mark:](./ttf/methods.go#L454) | [:x:](./ttf/ttf_functions_js.go#L240) |
| [TTF_SetFontSizeDPI](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSizeDPI) | [:heavy_check_mark:](./ttf/methods.go#L464) | [:x:](./ttf/ttf_functions_js.go#L258) |
| [TTF_GetFontSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSize) | [:heavy_check_mark:](./ttf/methods.go#L474) | [:x:](./ttf/ttf_functions_js.go#L280) |
| [TTF_GetFontDPI](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDPI) | [:heavy_check_mark:](./ttf/methods.go#L485) | [:x:](./ttf/ttf_functions_js.go#L296) |
| [TTF_SetFontStyle](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontStyle) | [:heavy_check_mark:](./ttf/methods.go#L497) | [:x:](./ttf/ttf_functions_js.go#L322) |
| [TTF_GetFontStyle](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyle) | [:heavy_check_mark:](./ttf/methods.go#L503) | [:x:](./ttf/ttf_functions_js.go#L338) |
| [TTF_SetFontOutline](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontOutline) | [:heavy_check_mark:](./ttf/methods.go#L509) | [:x:](./ttf/ttf_functions_js.go#L354) |
| [TTF_GetFontOutline](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontOutline) | [:heavy_check_mark:](./ttf/methods.go#L519) | [:x:](./ttf/ttf_functions_js.go#L372) |
| [TTF_SetFontHinting](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontHinting) | [:heavy_check_mark:](./ttf/methods.go#L525) | [:x:](./ttf/ttf_functions_js.go#L388) |
| [TTF_GetNumFontFaces](https://wiki.libsdl.org/SDL3_ttf/TTF_GetNumFontFaces) | [:heavy_check_mark:](./ttf/methods.go#L531) | [:x:](./ttf/ttf_functions_js.go#L404) |
| [TTF_GetFontHinting](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHinting) | [:heavy_check_mark:](./ttf/methods.go#L537) | [:x:](./ttf/ttf_functions_js.go#L420) |
| [TTF_SetFontSDF](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSDF) | [:heavy_check_mark:](./ttf/methods.go#L543) | [:x:](./ttf/ttf_functions_js.go#L436) |
| [TTF_GetFontSDF](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSDF) | [:heavy_check_mark:](./ttf/methods.go#L553) | [:x:](./ttf/ttf_functions_js.go#L454) |
| [TTF_SetFontWrapAlignment](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontWrapAlignment) | [:heavy_check_mark:](./ttf/methods.go#L559) | [:x:](./ttf/ttf_functions_js.go#L470) |
| [TTF_GetFontWrapAlignment](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontWrapAlignment) | [:heavy_check_mark:](./ttf/methods.go#L565) | [:x:](./ttf/ttf_functions_js.go#L486) |
| [TTF_GetFontHeight](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHeight) | [:heavy_check_mark:](./ttf/methods.go#L571) | [:x:](./ttf/ttf_functions_js.go#L502) |
| [TTF_GetFontAscent](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontAscent) | [:heavy_check_mark:](./ttf/methods.go#L577) | [:x:](./ttf/ttf_functions_js.go#L518) |
| [TTF_GetFontDescent](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDescent) | [:heavy_check_mark:](./ttf/methods.go#L583) | [:x:](./ttf/ttf_functions_js.go#L534) |
| [TTF_SetFontLineSkip](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLineSkip) | [:heavy_check_mark:](./ttf/methods.go#L589) | [:x:](./ttf/ttf_functions_js.go#L550) |
| [TTF_GetFontLineSkip](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontLineSkip) | [:heavy_check_mark:](./ttf/methods.go#L595) | [:x:](./ttf/ttf_functions_js.go#L566) |
| [TTF_SetFontKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontKerning) | [:heavy_check_mark:](./ttf/methods.go#L601) | [:x:](./ttf/ttf_functions_js.go#L582) |
| [TTF_GetFontKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontKerning) | [:heavy_check_mark:](./ttf/methods.go#L607) | [:x:](./ttf/ttf_functions_js.go#L598) |
| [TTF_FontIsFixedWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsFixedWidth) | [:heavy_check_mark:](./ttf/methods.go#L613) | [:x:](./ttf/ttf_functions_js.go#L614) |
| [TTF_FontIsScalable](https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsScalable) | [:heavy_check_mark:](./ttf/methods.go#L619) | [:x:](./ttf/ttf_functions_js.go#L630) |
| [TTF_GetFontFamilyName](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontFamilyName) | [:heavy_check_mark:](./ttf/methods.go#L625) | [:x:](./ttf/ttf_functions_js.go#L646) |
| [TTF_GetFontStyleName](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyleName) | [:heavy_check_mark:](./ttf/methods.go#L631) | [:x:](./ttf/ttf_functions_js.go#L662) |
| [TTF_SetFontDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontDirection) | [:heavy_check_mark:](./ttf/methods.go#L637) | [:x:](./ttf/ttf_functions_js.go#L678) |
| [TTF_GetFontDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDirection) | [:heavy_check_mark:](./ttf/methods.go#L647) | [:x:](./ttf/ttf_functions_js.go#L696) |
| [TTF_StringToTag](https://wiki.libsdl.org/SDL3_ttf/TTF_StringToTag) | [:heavy_check_mark:](./ttf/functions.go#L80) | [:x:](./ttf/ttf_functions_js.go#L712) |
| [TTF_TagToString](https://wiki.libsdl.org/SDL3_ttf/TTF_TagToString) | [:heavy_check_mark:](./ttf/functions.go#L86) | [:x:](./ttf/ttf_functions_js.go#L1980) |
| [TTF_SetFontScript](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontScript) | [:heavy_check_mark:](./ttf/methods.go#L653) | [:x:](./ttf/ttf_functions_js.go#L725) |
| [TTF_GetFontScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontScript) | [:heavy_check_mark:](./ttf/methods.go#L663) | [:x:](./ttf/ttf_functions_js.go#L743) |
| [TTF_GetGlyphScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphScript) | [:heavy_check_mark:](./ttf/functions.go#L100) | [:x:](./ttf/ttf_functions_js.go#L759) |
| [TTF_SetFontLanguage](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLanguage) | [:heavy_check_mark:](./ttf/methods.go#L669) | [:x:](./ttf/ttf_functions_js.go#L772) |
| [TTF_FontHasGlyph](https://wiki.libsdl.org/SDL3_ttf/TTF_FontHasGlyph) | [:heavy_check_mark:](./ttf/methods.go#L679) | [:x:](./ttf/ttf_functions_js.go#L790) |
| [TTF_GetGlyphImage](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImage) | [:heavy_check_mark:](./ttf/methods.go#L685) | [:x:](./ttf/ttf_functions_js.go#L808) |
| [TTF_GetGlyphImageForIndex](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImageForIndex) | [:heavy_check_mark:](./ttf/methods.go#L698) | [:x:](./ttf/ttf_functions_js.go#L832) |
| [TTF_GetGlyphMetrics](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphMetrics) | [:heavy_check_mark:](./ttf/methods.go#L711) | [:x:](./ttf/ttf_functions_js.go#L856) |
| [TTF_GetGlyphKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphKerning) | [:heavy_check_mark:](./ttf/methods.go#L723) | [:x:](./ttf/ttf_functions_js.go#L899) |
| [TTF_GetStringSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSize) | [:heavy_check_mark:](./ttf/methods.go#L735) | [:x:](./ttf/ttf_functions_js.go#L924) |
| [TTF_GetStringSizeWrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSizeWrapped) | [:heavy_check_mark:](./ttf/methods.go#L747) | [:x:](./ttf/ttf_functions_js.go#L954) |
| [TTF_MeasureString](https://wiki.libsdl.org/SDL3_ttf/TTF_MeasureString) | [:heavy_check_mark:](./ttf/methods.go#L760) | [:x:](./ttf/ttf_functions_js.go#L986) |
| [TTF_RenderText_Solid](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid) | [:heavy_check_mark:](./ttf/methods.go#L773) | [:x:](./ttf/ttf_functions_js.go#L1998) |
| [TTF_RenderText_Solid_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid_Wrapped) | [:heavy_check_mark:](./ttf/methods.go#L784) | [:x:](./ttf/ttf_functions_js.go#L2021) |
| [TTF_RenderGlyph_Solid](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Solid) | [:heavy_check_mark:](./ttf/methods.go#L795) | [:x:](./ttf/ttf_functions_js.go#L2046) |
| [TTF_RenderText_Shaded](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded) | [:heavy_check_mark:](./ttf/methods.go#L806) | [:x:](./ttf/ttf_functions_js.go#L2065) |
| [TTF_RenderText_Shaded_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded_Wrapped) | [:heavy_check_mark:](./ttf/methods.go#L817) | [:x:](./ttf/ttf_functions_js.go#L2088) |
| [TTF_RenderGlyph_Shaded](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Shaded) | [:heavy_check_mark:](./ttf/methods.go#L828) | [:x:](./ttf/ttf_functions_js.go#L2113) |
| [TTF_RenderText_Blended](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended) | [:heavy_check_mark:](./ttf/methods.go#L839) | [:x:](./ttf/ttf_functions_js.go#L2134) |
| [TTF_RenderText_Blended_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended_Wrapped) | [:heavy_check_mark:](./ttf/methods.go#L850) | [:x:](./ttf/ttf_functions_js.go#L2157) |
| [TTF_RenderGlyph_Blended](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Blended) | [:heavy_check_mark:](./ttf/methods.go#L861) | [:x:](./ttf/ttf_functions_js.go#L2182) |
| [TTF_RenderText_LCD](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD) | [:heavy_check_mark:](./ttf/methods.go#L872) | [:x:](./ttf/ttf_functions_js.go#L2201) |
| [TTF_RenderText_LCD_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD_Wrapped) | [:heavy_check_mark:](./ttf/methods.go#L883) | [:x:](./ttf/ttf_functions_js.go#L2224) |
| [TTF_RenderGlyph_LCD](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_LCD) | [:heavy_check_mark:](./ttf/methods.go#L894) | [:x:](./ttf/ttf_functions_js.go#L2249) |
| [TTF_CreateSurfaceTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateSurfaceTextEngine) | [:heavy_check_mark:](./ttf/functions.go#L111) | [:x:](./ttf/ttf_functions_js.go#L1018) |
| [TTF_DrawSurfaceText](https://wiki.libsdl.org/SDL3_ttf/TTF_DrawSurfaceText) | [:heavy_check_mark:](./ttf/methods.go#L13) | [:x:](./ttf/ttf_functions_js.go#L1030) |
| [TTF_DestroySurfaceTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroySurfaceTextEngine) | [:heavy_check_mark:](./ttf/methods.go#L366) | [:x:](./ttf/ttf_functions_js.go#L1055) |
| [TTF_CreateRendererTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngine) | [:heavy_check_mark:](./ttf/functions.go#L122) | [:x:](./ttf/ttf_functions_js.go#L1069) |
| [TTF_CreateRendererTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngineWithProperties) | [:heavy_check_mark:](./ttf/functions.go#L133) | [:x:](./ttf/ttf_functions_js.go#L1086) |
| [TTF_DrawRendererText](https://wiki.libsdl.org/SDL3_ttf/TTF_DrawRendererText) | [:heavy_check_mark:](./ttf/methods.go#L23) | [:x:](./ttf/ttf_functions_js.go#L1103) |
| [TTF_DestroyRendererTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyRendererTextEngine) | [:heavy_check_mark:](./ttf/methods.go#L372) | [:x:](./ttf/ttf_functions_js.go#L1123) |
| [TTF_CreateGPUTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngine) | [:question:]() | [:question:](./ttf/ttf_functions_js.go#L1137) |
| [TTF_CreateGPUTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngineWithProperties) | [:question:]() | [:question:](./ttf/ttf_functions_js.go#L1154) |
| [TTF_GetGPUTextDrawData](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextDrawData) | [:heavy_check_mark:](./ttf/methods.go#L33) | [:x:](./ttf/ttf_functions_js.go#L1171) |
| [TTF_DestroyGPUTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyGPUTextEngine) | [:heavy_check_mark:](./ttf/methods.go#L378) | [:x:](./ttf/ttf_functions_js.go#L1188) |
| [TTF_SetGPUTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_SetGPUTextEngineWinding) | [:heavy_check_mark:](./ttf/methods.go#L384) | [:x:](./ttf/ttf_functions_js.go#L1202) |
| [TTF_GetGPUTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextEngineWinding) | [:heavy_check_mark:](./ttf/methods.go#L390) | [:x:](./ttf/ttf_functions_js.go#L1218) |
| [TTF_CreateText](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateText) | [:heavy_check_mark:](./ttf/methods.go#L396) | [:x:](./ttf/ttf_functions_js.go#L1234) |
| [TTF_GetTextProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextProperties) | [:heavy_check_mark:](./ttf/methods.go#L44) | [:x:](./ttf/ttf_functions_js.go#L1260) |
| [TTF_SetTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextEngine) | [:heavy_check_mark:](./ttf/methods.go#L50) | [:x:](./ttf/ttf_functions_js.go#L1276) |
| [TTF_GetTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextEngine) | [:heavy_check_mark:](./ttf/methods.go#L60) | [:x:](./ttf/ttf_functions_js.go#L1297) |
| [TTF_SetTextFont](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextFont) | [:heavy_check_mark:](./ttf/methods.go#L71) | [:x:](./ttf/ttf_functions_js.go#L1314) |
| [TTF_GetTextFont](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextFont) | [:heavy_check_mark:](./ttf/methods.go#L81) | [:x:](./ttf/ttf_functions_js.go#L1335) |
| [TTF_SetTextDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextDirection) | [:heavy_check_mark:](./ttf/methods.go#L92) | [:x:](./ttf/ttf_functions_js.go#L1352) |
| [TTF_GetTextDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextDirection) | [:heavy_check_mark:](./ttf/methods.go#L102) | [:x:](./ttf/ttf_functions_js.go#L1370) |
| [TTF_SetTextScript](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextScript) | [:heavy_check_mark:](./ttf/methods.go#L108) | [:x:](./ttf/ttf_functions_js.go#L1386) |
| [TTF_GetTextScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextScript) | [:heavy_check_mark:](./ttf/methods.go#L118) | [:x:](./ttf/ttf_functions_js.go#L1404) |
| [TTF_SetTextColor](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColor) | [:heavy_check_mark:](./ttf/methods.go#L124) | [:x:](./ttf/ttf_functions_js.go#L1420) |
| [TTF_SetTextColorFloat](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColorFloat) | [:heavy_check_mark:](./ttf/methods.go#L134) | [:x:](./ttf/ttf_functions_js.go#L1444) |
| [TTF_GetTextColor](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColor) | [:heavy_check_mark:](./ttf/methods.go#L144) | [:x:](./ttf/ttf_functions_js.go#L1468) |
| [TTF_GetTextColorFloat](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColorFloat) | [:heavy_check_mark:](./ttf/methods.go#L156) | [:x:](./ttf/ttf_functions_js.go#L1504) |
| [TTF_SetTextPosition](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextPosition) | [:heavy_check_mark:](./ttf/methods.go#L168) | [:x:](./ttf/ttf_functions_js.go#L1540) |
| [TTF_GetTextPosition](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextPosition) | [:heavy_check_mark:](./ttf/methods.go#L174) | [:x:](./ttf/ttf_functions_js.go#L1560) |
| [TTF_SetTextWrapWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWidth) | [:heavy_check_mark:](./ttf/methods.go#L184) | [:x:](./ttf/ttf_functions_js.go#L1586) |
| [TTF_GetTextWrapWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextWrapWidth) | [:heavy_check_mark:](./ttf/methods.go#L194) | [:x:](./ttf/ttf_functions_js.go#L1604) |
| [TTF_SetTextWrapWhitespaceVisible](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWhitespaceVisible) | [:heavy_check_mark:](./ttf/methods.go#L206) | [:x:](./ttf/ttf_functions_js.go#L1625) |
| [TTF_TextWrapWhitespaceVisible](https://wiki.libsdl.org/SDL3_ttf/TTF_TextWrapWhitespaceVisible) | [:heavy_check_mark:](./ttf/methods.go#L216) | [:x:](./ttf/ttf_functions_js.go#L1643) |
| [TTF_SetTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextString) | [:heavy_check_mark:](./ttf/methods.go#L222) | [:x:](./ttf/ttf_functions_js.go#L1659) |
| [TTF_InsertTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_InsertTextString) | [:heavy_check_mark:](./ttf/methods.go#L232) | [:x:](./ttf/ttf_functions_js.go#L1679) |
| [TTF_AppendTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_AppendTextString) | [:heavy_check_mark:](./ttf/methods.go#L242) | [:x:](./ttf/ttf_functions_js.go#L1701) |
| [TTF_DeleteTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_DeleteTextString) | [:heavy_check_mark:](./ttf/methods.go#L252) | [:x:](./ttf/ttf_functions_js.go#L1721) |
| [TTF_GetTextSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSize) | [:heavy_check_mark:](./ttf/methods.go#L262) | [:x:](./ttf/ttf_functions_js.go#L1741) |
| [TTF_GetTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubString) | [:heavy_check_mark:](./ttf/methods.go#L274) | [:x:](./ttf/ttf_functions_js.go#L1767) |
| [TTF_GetTextSubStringForLine](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForLine) | [:heavy_check_mark:](./ttf/methods.go#L286) | [:x:](./ttf/ttf_functions_js.go#L1790) |
| [TTF_GetTextSubStringsForRange](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringsForRange) | [:heavy_check_mark:](./ttf/methods.go#L298) | [:x:](./ttf/ttf_functions_js.go#L1813) |
| [TTF_GetTextSubStringForPoint](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForPoint) | [:heavy_check_mark:](./ttf/methods.go#L312) | [:x:](./ttf/ttf_functions_js.go#L1839) |
| [TTF_GetPreviousTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetPreviousTextSubString) | [:heavy_check_mark:](./ttf/methods.go#L324) | [:x:](./ttf/ttf_functions_js.go#L1864) |
| [TTF_GetNextTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetNextTextSubString) | [:heavy_check_mark:](./ttf/methods.go#L336) | [:x:](./ttf/ttf_functions_js.go#L1890) |
| [TTF_UpdateText](https://wiki.libsdl.org/SDL3_ttf/TTF_UpdateText) | [:heavy_check_mark:](./ttf/methods.go#L348) | [:x:](./ttf/ttf_functions_js.go#L1916) |
| [TTF_DestroyText](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyText) | [:heavy_check_mark:](./ttf/methods.go#L358) | [:x:](./ttf/ttf_functions_js.go#L1932) |
| [TTF_CloseFont](https://wiki.libsdl.org/SDL3_ttf/TTF_CloseFont) | [:heavy_check_mark:](./ttf/methods.go#L905) | [:x:](./ttf/ttf_functions_js.go#L1946) |
| [TTF_Quit](https://wiki.libsdl.org/SDL3_ttf/TTF_Quit) | [:heavy_check_mark:](./ttf/functions.go#L144) | [:x:](./ttf/ttf_functions_js.go#L1960) |
| [TTF_WasInit](https://wiki.libsdl.org/SDL3_ttf/TTF_WasInit) | [:heavy_check_mark:](./ttf/functions.go#L150) | [:x:](./ttf/ttf_functions_js.go#L1969) |

## IMG

### Image

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [IMG_Version](https://wiki.libsdl.org/SDL3_image/IMG_Version) | [:heavy_check_mark:](./img/functions.go#L11) | [:x:](./img/img_functions_js.go#L13) |
| [IMG_LoadTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTyped_IO) | [:heavy_check_mark:](./img/functions.go#L17) | [:x:](./img/img_functions_js.go#L24) |
| [IMG_Load](https://wiki.libsdl.org/SDL3_image/IMG_Load) | [:heavy_check_mark:](./img/functions.go#L28) | [:x:](./img/img_functions_js.go#L45) |
| [IMG_Load_IO](https://wiki.libsdl.org/SDL3_image/IMG_Load_IO) | [:heavy_check_mark:](./img/functions.go#L39) | [:x:](./img/img_functions_js.go#L59) |
| [IMG_LoadTexture](https://wiki.libsdl.org/SDL3_image/IMG_LoadTexture) | [:heavy_check_mark:](./img/functions.go#L50) | [:x:](./img/img_functions_js.go#L78) |
| [IMG_LoadTexture_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTexture_IO) | [:heavy_check_mark:](./img/functions.go#L61) | [:x:](./img/img_functions_js.go#L97) |
| [IMG_LoadTextureTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTextureTyped_IO) | [:heavy_check_mark:](./img/functions.go#L72) | [:x:](./img/img_functions_js.go#L121) |
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
| [IMG_LoadAVIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAVIF_IO) | [:heavy_check_mark:](./img/functions.go#L191) | [:x:](./img/img_functions_js.go#L147) |
| [IMG_LoadICO_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadICO_IO) | [:heavy_check_mark:](./img/functions.go#L202) | [:x:](./img/img_functions_js.go#L164) |
| [IMG_LoadCUR_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadCUR_IO) | [:heavy_check_mark:](./img/functions.go#L213) | [:x:](./img/img_functions_js.go#L181) |
| [IMG_LoadBMP_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadBMP_IO) | [:heavy_check_mark:](./img/functions.go#L224) | [:x:](./img/img_functions_js.go#L198) |
| [IMG_LoadGIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadGIF_IO) | [:heavy_check_mark:](./img/functions.go#L235) | [:x:](./img/img_functions_js.go#L215) |
| [IMG_LoadJPG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadJPG_IO) | [:heavy_check_mark:](./img/functions.go#L246) | [:x:](./img/img_functions_js.go#L232) |
| [IMG_LoadJXL_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadJXL_IO) | [:heavy_check_mark:](./img/functions.go#L257) | [:x:](./img/img_functions_js.go#L249) |
| [IMG_LoadLBM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadLBM_IO) | [:heavy_check_mark:](./img/functions.go#L268) | [:x:](./img/img_functions_js.go#L266) |
| [IMG_LoadPCX_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPCX_IO) | [:heavy_check_mark:](./img/functions.go#L279) | [:x:](./img/img_functions_js.go#L283) |
| [IMG_LoadPNG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPNG_IO) | [:heavy_check_mark:](./img/functions.go#L290) | [:x:](./img/img_functions_js.go#L300) |
| [IMG_LoadPNM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPNM_IO) | [:heavy_check_mark:](./img/functions.go#L301) | [:x:](./img/img_functions_js.go#L317) |
| [IMG_LoadSVG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadSVG_IO) | [:heavy_check_mark:](./img/functions.go#L312) | [:x:](./img/img_functions_js.go#L334) |
| [IMG_LoadQOI_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadQOI_IO) | [:heavy_check_mark:](./img/functions.go#L323) | [:x:](./img/img_functions_js.go#L351) |
| [IMG_LoadTGA_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTGA_IO) | [:heavy_check_mark:](./img/functions.go#L334) | [:x:](./img/img_functions_js.go#L368) |
| [IMG_LoadTIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTIF_IO) | [:heavy_check_mark:](./img/functions.go#L345) | [:x:](./img/img_functions_js.go#L385) |
| [IMG_LoadXCF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXCF_IO) | [:heavy_check_mark:](./img/functions.go#L356) | [:x:](./img/img_functions_js.go#L402) |
| [IMG_LoadXPM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXPM_IO) | [:heavy_check_mark:](./img/functions.go#L367) | [:x:](./img/img_functions_js.go#L419) |
| [IMG_LoadXV_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXV_IO) | [:heavy_check_mark:](./img/functions.go#L378) | [:x:](./img/img_functions_js.go#L436) |
| [IMG_LoadWEBP_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadWEBP_IO) | [:heavy_check_mark:](./img/functions.go#L389) | [:x:](./img/img_functions_js.go#L453) |
| [IMG_LoadSizedSVG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadSizedSVG_IO) | [:heavy_check_mark:](./img/functions.go#L400) | [:x:](./img/img_functions_js.go#L470) |
| [IMG_ReadXPMFromArray](https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArray) | [:heavy_check_mark:](./img/functions.go#L411) | [:x:](./img/img_functions_js.go#L491) |
| [IMG_ReadXPMFromArrayToRGB888](https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArrayToRGB888) | [:heavy_check_mark:](./img/functions.go#L423) | [:x:](./img/img_functions_js.go#L508) |
| [IMG_SaveAVIF](https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF) | [:heavy_check_mark:](./img/functions.go#L435) | [:x:](./img/img_functions_js.go#L525) |
| [IMG_SaveAVIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF_IO) | [:heavy_check_mark:](./img/functions.go#L445) | [:x:](./img/img_functions_js.go#L545) |
| [IMG_SavePNG](https://wiki.libsdl.org/SDL3_image/IMG_SavePNG) | [:heavy_check_mark:](./img/functions.go#L455) | [:x:](./img/img_functions_js.go#L570) |
| [IMG_SavePNG_IO](https://wiki.libsdl.org/SDL3_image/IMG_SavePNG_IO) | [:heavy_check_mark:](./img/functions.go#L465) | [:x:](./img/img_functions_js.go#L588) |
| [IMG_SaveJPG](https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG) | [:heavy_check_mark:](./img/functions.go#L475) | [:x:](./img/img_functions_js.go#L611) |
| [IMG_SaveJPG_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG_IO) | [:heavy_check_mark:](./img/functions.go#L485) | [:x:](./img/img_functions_js.go#L631) |
| [IMG_LoadAnimation](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation) | [:heavy_check_mark:](./img/functions.go#L495) | [:x:](./img/img_functions_js.go#L656) |
| [IMG_LoadAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation_IO) | [:heavy_check_mark:](./img/functions.go#L506) | [:x:](./img/img_functions_js.go#L670) |
| [IMG_LoadAnimationTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimationTyped_IO) | [:heavy_check_mark:](./img/functions.go#L517) | [:x:](./img/img_functions_js.go#L689) |
| [IMG_FreeAnimation](https://wiki.libsdl.org/SDL3_image/IMG_FreeAnimation) | [:heavy_check_mark:](./img/methods.go#L6) | [:x:](./img/img_functions_js.go#L710) |
| [IMG_LoadGIFAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadGIFAnimation_IO) | [:heavy_check_mark:](./img/functions.go#L528) | [:x:](./img/img_functions_js.go#L724) |
| [IMG_LoadWEBPAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadWEBPAnimation_IO) | [:heavy_check_mark:](./img/functions.go#L539) | [:x:](./img/img_functions_js.go#L741) |

## MIXER

### Mixer

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [Mix_Version](https://wiki.libsdl.org/SDL3_mixer/Mix_Version) | [:heavy_check_mark:](./mixer/functions.go#L11) | [:x:](./mixer/mixer_functions_js.go#L13) |
| [Mix_Init](https://wiki.libsdl.org/SDL3_mixer/Mix_Init) | [:x:](./mixer/methods.go#L33) | [:x:](./mixer/mixer_functions_js.go#L24) |
| [Mix_Quit](https://wiki.libsdl.org/SDL3_mixer/Mix_Quit) | [:heavy_check_mark:](./mixer/functions.go#L23) | [:x:](./mixer/mixer_functions_js.go#L37) |
| [Mix_OpenAudio](https://wiki.libsdl.org/SDL3_mixer/Mix_OpenAudio) | [:heavy_check_mark:](./mixer/functions.go#L29) | [:x:](./mixer/mixer_functions_js.go#L46) |
| [Mix_PauseAudio](https://wiki.libsdl.org/SDL3_mixer/Mix_PauseAudio) | [:heavy_check_mark:](./mixer/functions.go#L44) | [:x:](./mixer/mixer_functions_js.go#L67) |
| [Mix_QuerySpec](https://wiki.libsdl.org/SDL3_mixer/Mix_QuerySpec) | [:heavy_check_mark:](./mixer/functions.go#L50) | [:x:](./mixer/mixer_functions_js.go#L78) |
| [Mix_AllocateChannels](https://wiki.libsdl.org/SDL3_mixer/Mix_AllocateChannels) | [:heavy_check_mark:](./mixer/functions.go#L62) | [:x:](./mixer/mixer_functions_js.go#L104) |
| [Mix_LoadWAV_IO](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadWAV_IO) | [:heavy_check_mark:](./mixer/functions.go#L68) | [:x:](./mixer/mixer_functions_js.go#L117) |
| [Mix_LoadWAV](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadWAV) | [:heavy_check_mark:](./mixer/functions.go#L79) | [:x:](./mixer/mixer_functions_js.go#L136) |
| [Mix_LoadMUS](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadMUS) | [:heavy_check_mark:](./mixer/functions.go#L101) | [:x:](./mixer/mixer_functions_js.go#L150) |
| [Mix_LoadMUS_IO](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadMUS_IO) | [:heavy_check_mark:](./mixer/functions.go#L90) | [:x:](./mixer/mixer_functions_js.go#L164) |
| [Mix_LoadMUSType_IO](https://wiki.libsdl.org/SDL3_mixer/Mix_LoadMUSType_IO) | [:heavy_check_mark:](./mixer/functions.go#L112) | [:x:](./mixer/mixer_functions_js.go#L183) |
| [Mix_QuickLoad_WAV](https://wiki.libsdl.org/SDL3_mixer/Mix_QuickLoad_WAV) | [:heavy_check_mark:](./mixer/functions.go#L123) | [:x:](./mixer/mixer_functions_js.go#L204) |
| [Mix_QuickLoad_RAW](https://wiki.libsdl.org/SDL3_mixer/Mix_QuickLoad_RAW) | [:heavy_check_mark:](./mixer/functions.go#L134) | [:x:](./mixer/mixer_functions_js.go#L221) |
| [Mix_FreeChunk](https://wiki.libsdl.org/SDL3_mixer/Mix_FreeChunk) | [:x:](./mixer/methods.go#L42) | [:x:](./mixer/mixer_functions_js.go#L240) |
| [Mix_FreeMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_FreeMusic) | [:x:](./mixer/methods.go#L56) | [:x:](./mixer/mixer_functions_js.go#L254) |
| [Mix_GetNumChunkDecoders](https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumChunkDecoders) | [:heavy_check_mark:](./mixer/functions.go#L145) | [:x:](./mixer/mixer_functions_js.go#L268) |
| [Mix_GetChunkDecoder](https://wiki.libsdl.org/SDL3_mixer/Mix_GetChunkDecoder) | [:heavy_check_mark:](./mixer/functions.go#L151) | [:x:](./mixer/mixer_functions_js.go#L279) |
| [Mix_HasChunkDecoder](https://wiki.libsdl.org/SDL3_mixer/Mix_HasChunkDecoder) | [:heavy_check_mark:](./mixer/functions.go#L157) | [:x:](./mixer/mixer_functions_js.go#L292) |
| [Mix_GetNumMusicDecoders](https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumMusicDecoders) | [:heavy_check_mark:](./mixer/functions.go#L163) | [:x:](./mixer/mixer_functions_js.go#L305) |
| [Mix_GetMusicDecoder](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicDecoder) | [:heavy_check_mark:](./mixer/functions.go#L169) | [:x:](./mixer/mixer_functions_js.go#L316) |
| [Mix_HasMusicDecoder](https://wiki.libsdl.org/SDL3_mixer/Mix_HasMusicDecoder) | [:heavy_check_mark:](./mixer/functions.go#L175) | [:x:](./mixer/mixer_functions_js.go#L329) |
| [Mix_GetMusicType](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicType) | [:x:](./mixer/methods.go#L63) | [:x:](./mixer/mixer_functions_js.go#L342) |
| [Mix_GetMusicTitle](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicTitle) | [:x:](./mixer/methods.go#L70) | [:x:](./mixer/mixer_functions_js.go#L358) |
| [Mix_GetMusicTitleTag](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicTitleTag) | [:x:](./mixer/methods.go#L77) | [:x:](./mixer/mixer_functions_js.go#L374) |
| [Mix_GetMusicArtistTag](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicArtistTag) | [:x:](./mixer/methods.go#L84) | [:x:](./mixer/mixer_functions_js.go#L390) |
| [Mix_GetMusicAlbumTag](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicAlbumTag) | [:x:](./mixer/methods.go#L91) | [:x:](./mixer/mixer_functions_js.go#L406) |
| [Mix_GetMusicCopyrightTag](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicCopyrightTag) | [:x:](./mixer/methods.go#L98) | [:x:](./mixer/mixer_functions_js.go#L422) |
| [Mix_SetPostMix](https://wiki.libsdl.org/SDL3_mixer/Mix_SetPostMix) | [:x:](./mixer/methods.go#L184) | [:x:](./mixer/mixer_functions_js.go#L438) |
| [Mix_HookMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_HookMusic) | [:x:](./mixer/methods.go#L191) | [:x:](./mixer/mixer_functions_js.go#L451) |
| [Mix_HookMusicFinished](https://wiki.libsdl.org/SDL3_mixer/Mix_HookMusicFinished) | [:x:](./mixer/methods.go#L6) | [:x:](./mixer/mixer_functions_js.go#L464) |
| [Mix_GetMusicHookData](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicHookData) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L475) |
| [Mix_ChannelFinished](https://wiki.libsdl.org/SDL3_mixer/Mix_ChannelFinished) | [:x:](./mixer/methods.go#L15) | [:x:](./mixer/mixer_functions_js.go#L486) |
| [Mix_RegisterEffect](https://wiki.libsdl.org/SDL3_mixer/Mix_RegisterEffect) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L497) |
| [Mix_UnregisterEffect](https://wiki.libsdl.org/SDL3_mixer/Mix_UnregisterEffect) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L516) |
| [Mix_UnregisterAllEffects](https://wiki.libsdl.org/SDL3_mixer/Mix_UnregisterAllEffects) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L531) |
| [Mix_SetPanning](https://wiki.libsdl.org/SDL3_mixer/Mix_SetPanning) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L544) |
| [Mix_SetPosition](https://wiki.libsdl.org/SDL3_mixer/Mix_SetPosition) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L561) |
| [Mix_SetDistance](https://wiki.libsdl.org/SDL3_mixer/Mix_SetDistance) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L578) |
| [Mix_SetReverseStereo](https://wiki.libsdl.org/SDL3_mixer/Mix_SetReverseStereo) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L593) |
| [Mix_ReserveChannels](https://wiki.libsdl.org/SDL3_mixer/Mix_ReserveChannels) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L608) |
| [Mix_GroupChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupChannel) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L621) |
| [Mix_GroupChannels](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupChannels) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L636) |
| [Mix_GroupAvailable](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupAvailable) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L653) |
| [Mix_GroupCount](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupCount) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L666) |
| [Mix_GroupOldest](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupOldest) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L679) |
| [Mix_GroupNewer](https://wiki.libsdl.org/SDL3_mixer/Mix_GroupNewer) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L692) |
| [Mix_PlayChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_PlayChannel) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L705) |
| [Mix_PlayChannelTimed](https://wiki.libsdl.org/SDL3_mixer/Mix_PlayChannelTimed) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L725) |
| [Mix_PlayMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_PlayMusic) | [:x:](./mixer/methods.go#L105) | [:x:](./mixer/mixer_functions_js.go#L747) |
| [Mix_FadeInMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInMusic) | [:x:](./mixer/methods.go#L112) | [:x:](./mixer/mixer_functions_js.go#L765) |
| [Mix_FadeInMusicPos](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInMusicPos) | [:x:](./mixer/methods.go#L119) | [:x:](./mixer/mixer_functions_js.go#L785) |
| [Mix_FadeInChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInChannel) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L807) |
| [Mix_FadeInChannelTimed](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInChannelTimed) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L829) |
| [Mix_Volume](https://wiki.libsdl.org/SDL3_mixer/Mix_Volume) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L853) |
| [Mix_VolumeChunk](https://wiki.libsdl.org/SDL3_mixer/Mix_VolumeChunk) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L868) |
| [Mix_VolumeMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_VolumeMusic) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L886) |
| [Mix_GetMusicVolume](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicVolume) | [:x:](./mixer/methods.go#L126) | [:x:](./mixer/mixer_functions_js.go#L899) |
| [Mix_MasterVolume](https://wiki.libsdl.org/SDL3_mixer/Mix_MasterVolume) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L915) |
| [Mix_HaltChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_HaltChannel) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L928) |
| [Mix_HaltGroup](https://wiki.libsdl.org/SDL3_mixer/Mix_HaltGroup) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L939) |
| [Mix_HaltMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_HaltMusic) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L950) |
| [Mix_ExpireChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_ExpireChannel) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L959) |
| [Mix_FadeOutChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeOutChannel) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L974) |
| [Mix_FadeOutGroup](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeOutGroup) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L989) |
| [Mix_FadeOutMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_FadeOutMusic) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1004) |
| [Mix_FadingMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_FadingMusic) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1017) |
| [Mix_FadingChannel](https://wiki.libsdl.org/SDL3_mixer/Mix_FadingChannel) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1028) |
| [Mix_Pause](https://wiki.libsdl.org/SDL3_mixer/Mix_Pause) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1041) |
| [Mix_PauseGroup](https://wiki.libsdl.org/SDL3_mixer/Mix_PauseGroup) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1052) |
| [Mix_Resume](https://wiki.libsdl.org/SDL3_mixer/Mix_Resume) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1063) |
| [Mix_ResumeGroup](https://wiki.libsdl.org/SDL3_mixer/Mix_ResumeGroup) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1074) |
| [Mix_Paused](https://wiki.libsdl.org/SDL3_mixer/Mix_Paused) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1085) |
| [Mix_PauseMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_PauseMusic) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1098) |
| [Mix_ResumeMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_ResumeMusic) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1107) |
| [Mix_RewindMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_RewindMusic) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1116) |
| [Mix_PausedMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_PausedMusic) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1125) |
| [Mix_ModMusicJumpToOrder](https://wiki.libsdl.org/SDL3_mixer/Mix_ModMusicJumpToOrder) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1136) |
| [Mix_StartTrack](https://wiki.libsdl.org/SDL3_mixer/Mix_StartTrack) | [:x:](./mixer/methods.go#L133) | [:x:](./mixer/mixer_functions_js.go#L1149) |
| [Mix_GetNumTracks](https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumTracks) | [:x:](./mixer/methods.go#L140) | [:x:](./mixer/mixer_functions_js.go#L1167) |
| [Mix_SetMusicPosition](https://wiki.libsdl.org/SDL3_mixer/Mix_SetMusicPosition) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1183) |
| [Mix_GetMusicPosition](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicPosition) | [:x:](./mixer/methods.go#L147) | [:x:](./mixer/mixer_functions_js.go#L1196) |
| [Mix_MusicDuration](https://wiki.libsdl.org/SDL3_mixer/Mix_MusicDuration) | [:x:](./mixer/methods.go#L154) | [:x:](./mixer/mixer_functions_js.go#L1212) |
| [Mix_GetMusicLoopStartTime](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopStartTime) | [:x:](./mixer/methods.go#L161) | [:x:](./mixer/mixer_functions_js.go#L1228) |
| [Mix_GetMusicLoopEndTime](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopEndTime) | [:x:](./mixer/methods.go#L168) | [:x:](./mixer/mixer_functions_js.go#L1244) |
| [Mix_GetMusicLoopLengthTime](https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopLengthTime) | [:x:](./mixer/methods.go#L175) | [:x:](./mixer/mixer_functions_js.go#L1260) |
| [Mix_Playing](https://wiki.libsdl.org/SDL3_mixer/Mix_Playing) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1276) |
| [Mix_PlayingMusic](https://wiki.libsdl.org/SDL3_mixer/Mix_PlayingMusic) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1289) |
| [Mix_SetSoundFonts](https://wiki.libsdl.org/SDL3_mixer/Mix_SetSoundFonts) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1300) |
| [Mix_GetSoundFonts](https://wiki.libsdl.org/SDL3_mixer/Mix_GetSoundFonts) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1313) |
| [Mix_EachSoundFont](https://wiki.libsdl.org/SDL3_mixer/Mix_EachSoundFont) | [:x:](./mixer/methods.go#L24) | [:x:](./mixer/mixer_functions_js.go#L1324) |
| [Mix_SetTimidityCfg](https://wiki.libsdl.org/SDL3_mixer/Mix_SetTimidityCfg) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1339) |
| [Mix_GetTimidityCfg](https://wiki.libsdl.org/SDL3_mixer/Mix_GetTimidityCfg) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1352) |
| [Mix_GetChunk](https://wiki.libsdl.org/SDL3_mixer/Mix_GetChunk) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1363) |
| [Mix_CloseAudio](https://wiki.libsdl.org/SDL3_mixer/Mix_CloseAudio) | [:question:]() | [:question:](./mixer/mixer_functions_js.go#L1377) |

# API Coverage


This file tracks the functions that have been wrapped.<br>
The following emojis mean (they are clickable and should link to the code implementation):
- :heavy_check_mark: = implemented
- :x: = not implemented yet
- :question: = not planned / don't know about integrating it or not
<details open>
<summary><h2>SDL</h2></summary>
<details open>
<summary><h3>Init</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_Init](https://wiki.libsdl.org/SDL3/SDL_Init) | [:heavy_check_mark:](sdl/functions.go#L14) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9659) |
| [SDL_InitSubSystem](https://wiki.libsdl.org/SDL3/SDL_InitSubSystem) | [:heavy_check_mark:](sdl/functions.go#L24) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9669) |
| [SDL_QuitSubSystem](https://wiki.libsdl.org/SDL3/SDL_QuitSubSystem) | [:heavy_check_mark:](sdl/functions.go#L34) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9678) |
| [SDL_WasInit](https://wiki.libsdl.org/SDL3/SDL_WasInit) | [:heavy_check_mark:](sdl/functions.go#L40) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9686) |
| [SDL_Quit](https://wiki.libsdl.org/SDL3/SDL_Quit) | [:heavy_check_mark:](sdl/functions.go#L46) | [:question:]() |
| [SDL_IsMainThread](https://wiki.libsdl.org/SDL3/SDL_IsMainThread) | [:heavy_check_mark:](sdl/functions.go#L52) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9699) |
| [SDL_RunOnMainThread](https://wiki.libsdl.org/SDL3/SDL_RunOnMainThread) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12952) |
| [SDL_SetAppMetadata](https://wiki.libsdl.org/SDL3/SDL_SetAppMetadata) | [:heavy_check_mark:](sdl/functions.go#L58) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9723) |
| [SDL_SetAppMetadataProperty](https://wiki.libsdl.org/SDL3/SDL_SetAppMetadataProperty) | [:heavy_check_mark:](sdl/functions.go#L68) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9738) |
| [SDL_GetAppMetadataProperty](https://wiki.libsdl.org/SDL3/SDL_GetAppMetadataProperty) | [:heavy_check_mark:](sdl/functions.go#L78) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9751) |
</details>
<details open>
<summary><h3>Hints</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetHintWithPriority](https://wiki.libsdl.org/SDL3/SDL_SetHintWithPriority) | [:heavy_check_mark:](sdl/functions.go#L86) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9558) |
| [SDL_SetHint](https://wiki.libsdl.org/SDL3/SDL_SetHint) | [:heavy_check_mark:](sdl/functions.go#L96) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9573) |
| [SDL_ResetHint](https://wiki.libsdl.org/SDL3/SDL_ResetHint) | [:heavy_check_mark:](sdl/functions.go#L106) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9586) |
| [SDL_ResetHints](https://wiki.libsdl.org/SDL3/SDL_ResetHints) | [:heavy_check_mark:](sdl/functions.go#L116) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9597) |
| [SDL_GetHint](https://wiki.libsdl.org/SDL3/SDL_GetHint) | [:heavy_check_mark:](sdl/functions.go#L122) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9603) |
| [SDL_GetHintBoolean](https://wiki.libsdl.org/SDL3/SDL_GetHintBoolean) | [:heavy_check_mark:](sdl/functions.go#L128) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9614) |
| [SDL_AddHintCallback](https://wiki.libsdl.org/SDL3/SDL_AddHintCallback) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12921) |
| [SDL_RemoveHintCallback](https://wiki.libsdl.org/SDL3/SDL_RemoveHintCallback) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12937) |
</details>
<details>
<summary><h3>Error</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetError](https://wiki.libsdl.org/SDL3/SDL_SetError) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L365) |
| [SDL_SetErrorV](https://wiki.libsdl.org/SDL3/SDL_SetErrorV) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L376) |
| [SDL_OutOfMemory](https://wiki.libsdl.org/SDL3/SDL_OutOfMemory) | [:heavy_check_mark:](sdl/functions.go#L139) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L390) |
| [SDL_GetError](https://wiki.libsdl.org/SDL3/SDL_GetError) | [:heavy_check_mark:](sdl/init_notjs.go#L32) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L397) |
| [SDL_ClearError](https://wiki.libsdl.org/SDL3/SDL_ClearError) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L405) |
</details>
<details>
<summary><h3>Version</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetVersion](https://wiki.libsdl.org/SDL3/SDL_GetVersion) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L23) |
| [SDL_GetRevision](https://wiki.libsdl.org/SDL3/SDL_GetRevision) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13466) |
</details>
<details open>
<summary><h3>Properties</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetGlobalProperties](https://wiki.libsdl.org/SDL3/SDL_GetGlobalProperties) | [:heavy_check_mark:](sdl/functions.go#L147) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L412) |
| [SDL_CreateProperties](https://wiki.libsdl.org/SDL3/SDL_CreateProperties) | [:heavy_check_mark:](sdl/functions.go#L158) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L419) |
| [SDL_CopyProperties](https://wiki.libsdl.org/SDL3/SDL_CopyProperties) | [:heavy_check_mark:](sdl/functions.go#L169) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L427) |
| [SDL_LockProperties](https://wiki.libsdl.org/SDL3/SDL_LockProperties) | [:heavy_check_mark:](sdl/methods.go#L5785) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L438) |
| [SDL_UnlockProperties](https://wiki.libsdl.org/SDL3/SDL_UnlockProperties) | [:heavy_check_mark:](sdl/methods.go#L5795) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L447) |
| [SDL_SetPointerPropertyWithCleanup](https://wiki.libsdl.org/SDL3/SDL_SetPointerPropertyWithCleanup) | [:x:](sdl/methods.go#L5801) | [:x:](sdl/sdl_functions_js.go#L12279) |
| [SDL_SetPointerProperty](https://wiki.libsdl.org/SDL3/SDL_SetPointerProperty) | [:heavy_check_mark:](sdl/methods.go#L5808) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L476) |
| [SDL_SetStringProperty](https://wiki.libsdl.org/SDL3/SDL_SetStringProperty) | [:heavy_check_mark:](sdl/methods.go#L5818) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L492) |
| [SDL_SetNumberProperty](https://wiki.libsdl.org/SDL3/SDL_SetNumberProperty) | [:heavy_check_mark:](sdl/methods.go#L5828) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L508) |
| [SDL_SetFloatProperty](https://wiki.libsdl.org/SDL3/SDL_SetFloatProperty) | [:heavy_check_mark:](sdl/methods.go#L5838) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L524) |
| [SDL_SetBooleanProperty](https://wiki.libsdl.org/SDL3/SDL_SetBooleanProperty) | [:heavy_check_mark:](sdl/methods.go#L5848) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L539) |
| [SDL_HasProperty](https://wiki.libsdl.org/SDL3/SDL_HasProperty) | [:heavy_check_mark:](sdl/methods.go#L5858) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L554) |
| [SDL_GetPropertyType](https://wiki.libsdl.org/SDL3/SDL_GetPropertyType) | [:heavy_check_mark:](sdl/methods.go#L5864) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L567) |
| [SDL_GetPointerProperty](https://wiki.libsdl.org/SDL3/SDL_GetPointerProperty) | [:heavy_check_mark:](sdl/methods.go#L5870) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L580) |
| [SDL_GetStringProperty](https://wiki.libsdl.org/SDL3/SDL_GetStringProperty) | [:heavy_check_mark:](sdl/methods.go#L5876) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L596) |
| [SDL_GetNumberProperty](https://wiki.libsdl.org/SDL3/SDL_GetNumberProperty) | [:heavy_check_mark:](sdl/methods.go#L5882) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L611) |
| [SDL_GetFloatProperty](https://wiki.libsdl.org/SDL3/SDL_GetFloatProperty) | [:heavy_check_mark:](sdl/methods.go#L5888) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L626) |
| [SDL_GetBooleanProperty](https://wiki.libsdl.org/SDL3/SDL_GetBooleanProperty) | [:heavy_check_mark:](sdl/methods.go#L5894) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L641) |
| [SDL_ClearProperty](https://wiki.libsdl.org/SDL3/SDL_ClearProperty) | [:heavy_check_mark:](sdl/methods.go#L5900) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L656) |
| [SDL_GetNumProperties](https://wiki.libsdl.org/SDL3/SDL_GetNumProperties) | [:question:]() | [:question:]() |
| [SDL_EnumerateProperties](https://wiki.libsdl.org/SDL3/SDL_EnumerateProperties) | [:heavy_check_mark:](sdl/methods.go#L5910) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L669) |
| [SDL_DestroyProperties](https://wiki.libsdl.org/SDL3/SDL_DestroyProperties) | [:heavy_check_mark:](sdl/methods.go#L5916) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L684) |
</details>
<details>
<summary><h3>Log</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetLogPriorities](https://wiki.libsdl.org/SDL3/SDL_SetLogPriorities) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9811) |
| [SDL_SetLogPriority](https://wiki.libsdl.org/SDL3/SDL_SetLogPriority) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9819) |
| [SDL_GetLogPriority](https://wiki.libsdl.org/SDL3/SDL_GetLogPriority) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9829) |
| [SDL_ResetLogPriorities](https://wiki.libsdl.org/SDL3/SDL_ResetLogPriorities) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9838) |
| [SDL_SetLogPriorityPrefix](https://wiki.libsdl.org/SDL3/SDL_SetLogPriorityPrefix) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9844) |
| [SDL_Log](https://wiki.libsdl.org/SDL3/SDL_Log) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9857) |
| [SDL_LogTrace](https://wiki.libsdl.org/SDL3/SDL_LogTrace) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9867) |
| [SDL_LogVerbose](https://wiki.libsdl.org/SDL3/SDL_LogVerbose) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9879) |
| [SDL_LogDebug](https://wiki.libsdl.org/SDL3/SDL_LogDebug) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9891) |
| [SDL_LogInfo](https://wiki.libsdl.org/SDL3/SDL_LogInfo) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9903) |
| [SDL_LogWarn](https://wiki.libsdl.org/SDL3/SDL_LogWarn) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9915) |
| [SDL_LogError](https://wiki.libsdl.org/SDL3/SDL_LogError) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9927) |
| [SDL_LogCritical](https://wiki.libsdl.org/SDL3/SDL_LogCritical) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9939) |
| [SDL_LogMessage](https://wiki.libsdl.org/SDL3/SDL_LogMessage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9951) |
| [SDL_LogMessageV](https://wiki.libsdl.org/SDL3/SDL_LogMessageV) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9965) |
| [SDL_GetDefaultLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_GetDefaultLogOutputFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12980) |
| [SDL_GetLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_GetLogOutputFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9993) |
| [SDL_SetLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_SetLogOutputFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12988) |
</details>
<details open>
<summary><h3>Video</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumVideoDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumVideoDrivers) | [:heavy_check_mark:](sdl/functions.go#L614) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4179) |
| [SDL_GetVideoDriver](https://wiki.libsdl.org/SDL3/SDL_GetVideoDriver) | [:heavy_check_mark:](sdl/functions.go#L620) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4186) |
| [SDL_GetCurrentVideoDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentVideoDriver) | [:heavy_check_mark:](sdl/functions.go#L626) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4195) |
| [SDL_GetSystemTheme](https://wiki.libsdl.org/SDL3/SDL_GetSystemTheme) | [:heavy_check_mark:](sdl/functions.go#L632) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4202) |
| [SDL_GetDisplays](https://wiki.libsdl.org/SDL3/SDL_GetDisplays) | [:heavy_check_mark:](sdl/functions.go#L638) | [:x:](sdl/sdl_functions_js.go#L4209) |
| [SDL_GetPrimaryDisplay](https://wiki.libsdl.org/SDL3/SDL_GetPrimaryDisplay) | [:heavy_check_mark:](sdl/functions.go#L652) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4221) |
| [SDL_GetDisplayProperties](https://wiki.libsdl.org/SDL3/SDL_GetDisplayProperties) | [:heavy_check_mark:](sdl/methods.go#L3702) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4228) |
| [SDL_GetDisplayName](https://wiki.libsdl.org/SDL3/SDL_GetDisplayName) | [:heavy_check_mark:](sdl/methods.go#L3713) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4237) |
| [SDL_GetDisplayBounds](https://wiki.libsdl.org/SDL3/SDL_GetDisplayBounds) | [:heavy_check_mark:](sdl/methods.go#L3724) | [:x:](sdl/sdl_functions_js.go#L4246) |
| [SDL_GetDisplayUsableBounds](https://wiki.libsdl.org/SDL3/SDL_GetDisplayUsableBounds) | [:heavy_check_mark:](sdl/methods.go#L3736) | [:x:](sdl/sdl_functions_js.go#L4260) |
| [SDL_GetNaturalDisplayOrientation](https://wiki.libsdl.org/SDL3/SDL_GetNaturalDisplayOrientation) | [:heavy_check_mark:](sdl/methods.go#L3748) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4274) |
| [SDL_GetCurrentDisplayOrientation](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayOrientation) | [:heavy_check_mark:](sdl/methods.go#L3754) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4283) |
| [SDL_GetDisplayContentScale](https://wiki.libsdl.org/SDL3/SDL_GetDisplayContentScale) | [:heavy_check_mark:](sdl/methods.go#L3760) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4292) |
| [SDL_GetFullscreenDisplayModes](https://wiki.libsdl.org/SDL3/SDL_GetFullscreenDisplayModes) | [:heavy_check_mark:](sdl/methods.go#L3771) | [:x:](sdl/sdl_functions_js.go#L4301) |
| [SDL_GetClosestFullscreenDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetClosestFullscreenDisplayMode) | [:heavy_check_mark:](sdl/methods.go#L3785) | [:x:](sdl/sdl_functions_js.go#L4315) |
| [SDL_GetDesktopDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetDesktopDisplayMode) | [:heavy_check_mark:](sdl/methods.go#L3797) | [:x:](sdl/sdl_functions_js.go#L4337) |
| [SDL_GetCurrentDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayMode) | [:heavy_check_mark:](sdl/methods.go#L3808) | [:x:](sdl/sdl_functions_js.go#L4347) |
| [SDL_GetDisplayForPoint](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForPoint) | [:heavy_check_mark:](sdl/functions.go#L658) | [:x:](sdl/sdl_functions_js.go#L4357) |
| [SDL_GetDisplayForRect](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForRect) | [:heavy_check_mark:](sdl/functions.go#L664) | [:x:](sdl/sdl_functions_js.go#L4369) |
| [SDL_GetDisplayForWindow](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForWindow) | [:heavy_check_mark:](sdl/functions.go#L670) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4381) |
| [SDL_GetWindowPixelDensity](https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelDensity) | [:heavy_check_mark:](sdl/methods.go#L4142) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4390) |
| [SDL_GetWindowDisplayScale](https://wiki.libsdl.org/SDL3/SDL_GetWindowDisplayScale) | [:heavy_check_mark:](sdl/methods.go#L4153) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4399) |
| [SDL_SetWindowFullscreenMode](https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreenMode) | [:heavy_check_mark:](sdl/methods.go#L4164) | [:x:](sdl/sdl_functions_js.go#L4408) |
| [SDL_GetWindowFullscreenMode](https://wiki.libsdl.org/SDL3/SDL_GetWindowFullscreenMode) | [:heavy_check_mark:](sdl/methods.go#L4174) | [:x:](sdl/sdl_functions_js.go#L4422) |
| [SDL_GetWindowICCProfile](https://wiki.libsdl.org/SDL3/SDL_GetWindowICCProfile) | [:heavy_check_mark:](sdl/methods.go#L4180) | [:x:](sdl/sdl_functions_js.go#L4432) |
| [SDL_GetWindowPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelFormat) | [:heavy_check_mark:](sdl/methods.go#L4194) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4446) |
| [SDL_GetWindows](https://wiki.libsdl.org/SDL3/SDL_GetWindows) | [:heavy_check_mark:](sdl/functions.go#L676) | [:x:](sdl/sdl_functions_js.go#L4459) |
| [SDL_CreateWindow](https://wiki.libsdl.org/SDL3/SDL_CreateWindow) | [:heavy_check_mark:](sdl/functions.go#L690) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4471) |
| [SDL_CreatePopupWindow](https://wiki.libsdl.org/SDL3/SDL_CreatePopupWindow) | [:heavy_check_mark:](sdl/methods.go#L4205) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4488) |
| [SDL_CreateWindowWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateWindowWithProperties) | [:heavy_check_mark:](sdl/functions.go#L701) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4507) |
| [SDL_GetWindowID](https://wiki.libsdl.org/SDL3/SDL_GetWindowID) | [:heavy_check_mark:](sdl/methods.go#L4216) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4516) |
| [SDL_GetWindowFromID](https://wiki.libsdl.org/SDL3/SDL_GetWindowFromID) | [:heavy_check_mark:](sdl/methods.go#L32) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4525) |
| [SDL_GetWindowParent](https://wiki.libsdl.org/SDL3/SDL_GetWindowParent) | [:heavy_check_mark:](sdl/methods.go#L4227) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4534) |
| [SDL_GetWindowProperties](https://wiki.libsdl.org/SDL3/SDL_GetWindowProperties) | [:heavy_check_mark:](sdl/methods.go#L4233) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4543) |
| [SDL_GetWindowFlags](https://wiki.libsdl.org/SDL3/SDL_GetWindowFlags) | [:heavy_check_mark:](sdl/methods.go#L4244) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4552) |
| [SDL_SetWindowTitle](https://wiki.libsdl.org/SDL3/SDL_SetWindowTitle) | [:heavy_check_mark:](sdl/methods.go#L4250) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4561) |
| [SDL_GetWindowTitle](https://wiki.libsdl.org/SDL3/SDL_GetWindowTitle) | [:heavy_check_mark:](sdl/methods.go#L4260) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4574) |
| [SDL_SetWindowIcon](https://wiki.libsdl.org/SDL3/SDL_SetWindowIcon) | [:heavy_check_mark:](sdl/methods.go#L4266) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4583) |
| [SDL_SetWindowPosition](https://wiki.libsdl.org/SDL3/SDL_SetWindowPosition) | [:heavy_check_mark:](sdl/methods.go#L4276) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4594) |
| [SDL_GetWindowPosition](https://wiki.libsdl.org/SDL3/SDL_GetWindowPosition) | [:heavy_check_mark:](sdl/methods.go#L4286) | [:x:](sdl/sdl_functions_js.go#L4607) |
| [SDL_SetWindowSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowSize) | [:heavy_check_mark:](sdl/methods.go#L4298) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4623) |
| [SDL_GetWindowSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowSize) | [:heavy_check_mark:](sdl/methods.go#L4308) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4636) |
| [SDL_GetWindowSafeArea](https://wiki.libsdl.org/SDL3/SDL_GetWindowSafeArea) | [:heavy_check_mark:](sdl/methods.go#L4320) | [:x:](sdl/sdl_functions_js.go#L4658) |
| [SDL_SetWindowAspectRatio](https://wiki.libsdl.org/SDL3/SDL_SetWindowAspectRatio) | [:heavy_check_mark:](sdl/methods.go#L4332) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4672) |
| [SDL_GetWindowAspectRatio](https://wiki.libsdl.org/SDL3/SDL_GetWindowAspectRatio) | [:heavy_check_mark:](sdl/methods.go#L4342) | [:x:](sdl/sdl_functions_js.go#L4685) |
| [SDL_GetWindowBordersSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowBordersSize) | [:heavy_check_mark:](sdl/methods.go#L4354) | [:x:](sdl/sdl_functions_js.go#L4701) |
| [SDL_GetWindowSizeInPixels](https://wiki.libsdl.org/SDL3/SDL_GetWindowSizeInPixels) | [:heavy_check_mark:](sdl/methods.go#L4366) | [:x:](sdl/sdl_functions_js.go#L4721) |
| [SDL_SetWindowMinimumSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowMinimumSize) | [:heavy_check_mark:](sdl/methods.go#L4378) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4737) |
| [SDL_GetWindowMinimumSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowMinimumSize) | [:heavy_check_mark:](sdl/methods.go#L4388) | [:x:](sdl/sdl_functions_js.go#L4750) |
| [SDL_SetWindowMaximumSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowMaximumSize) | [:heavy_check_mark:](sdl/methods.go#L4400) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4766) |
| [SDL_GetWindowMaximumSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowMaximumSize) | [:heavy_check_mark:](sdl/methods.go#L4410) | [:x:](sdl/sdl_functions_js.go#L4779) |
| [SDL_SetWindowBordered](https://wiki.libsdl.org/SDL3/SDL_SetWindowBordered) | [:heavy_check_mark:](sdl/methods.go#L4422) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4795) |
| [SDL_SetWindowResizable](https://wiki.libsdl.org/SDL3/SDL_SetWindowResizable) | [:heavy_check_mark:](sdl/methods.go#L4432) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4806) |
| [SDL_SetWindowAlwaysOnTop](https://wiki.libsdl.org/SDL3/SDL_SetWindowAlwaysOnTop) | [:heavy_check_mark:](sdl/methods.go#L4442) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4821) |
| [SDL_SetWindowFillDocument](https://wiki.libsdl.org/SDL3/SDL_SetWindowFillDocument) | [:heavy_check_mark:](sdl/methods.go#L4452) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12519) |
| [SDL_ShowWindow](https://wiki.libsdl.org/SDL3/SDL_ShowWindow) | [:heavy_check_mark:](sdl/methods.go#L4462) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4832) |
| [SDL_HideWindow](https://wiki.libsdl.org/SDL3/SDL_HideWindow) | [:heavy_check_mark:](sdl/methods.go#L4472) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4841) |
| [SDL_RaiseWindow](https://wiki.libsdl.org/SDL3/SDL_RaiseWindow) | [:heavy_check_mark:](sdl/methods.go#L4482) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4850) |
| [SDL_MaximizeWindow](https://wiki.libsdl.org/SDL3/SDL_MaximizeWindow) | [:heavy_check_mark:](sdl/methods.go#L4492) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4859) |
| [SDL_MinimizeWindow](https://wiki.libsdl.org/SDL3/SDL_MinimizeWindow) | [:heavy_check_mark:](sdl/methods.go#L4502) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4868) |
| [SDL_RestoreWindow](https://wiki.libsdl.org/SDL3/SDL_RestoreWindow) | [:heavy_check_mark:](sdl/methods.go#L4512) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4877) |
| [SDL_SetWindowFullscreen](https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreen) | [:heavy_check_mark:](sdl/methods.go#L4522) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4886) |
| [SDL_SyncWindow](https://wiki.libsdl.org/SDL3/SDL_SyncWindow) | [:heavy_check_mark:](sdl/methods.go#L4532) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4897) |
| [SDL_WindowHasSurface](https://wiki.libsdl.org/SDL3/SDL_WindowHasSurface) | [:heavy_check_mark:](sdl/methods.go#L4542) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4906) |
| [SDL_GetWindowSurface](https://wiki.libsdl.org/SDL3/SDL_GetWindowSurface) | [:heavy_check_mark:](sdl/methods.go#L4548) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4915) |
| [SDL_SetWindowSurfaceVSync](https://wiki.libsdl.org/SDL3/SDL_SetWindowSurfaceVSync) | [:heavy_check_mark:](sdl/methods.go#L4559) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4924) |
| [SDL_GetWindowSurfaceVSync](https://wiki.libsdl.org/SDL3/SDL_GetWindowSurfaceVSync) | [:heavy_check_mark:](sdl/methods.go#L4569) | [:x:](sdl/sdl_functions_js.go#L4935) |
| [SDL_UpdateWindowSurface](https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurface) | [:heavy_check_mark:](sdl/methods.go#L4581) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4949) |
| [SDL_UpdateWindowSurfaceRects](https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurfaceRects) | [:heavy_check_mark:](sdl/methods.go#L4591) | [:x:](sdl/sdl_functions_js.go#L4958) |
| [SDL_DestroyWindowSurface](https://wiki.libsdl.org/SDL3/SDL_DestroyWindowSurface) | [:heavy_check_mark:](sdl/methods.go#L4601) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4974) |
| [SDL_SetWindowKeyboardGrab](https://wiki.libsdl.org/SDL3/SDL_SetWindowKeyboardGrab) | [:heavy_check_mark:](sdl/methods.go#L4611) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4984) |
| [SDL_SetWindowMouseGrab](https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseGrab) | [:heavy_check_mark:](sdl/methods.go#L4621) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4995) |
| [SDL_GetWindowKeyboardGrab](https://wiki.libsdl.org/SDL3/SDL_GetWindowKeyboardGrab) | [:heavy_check_mark:](sdl/methods.go#L4631) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5006) |
| [SDL_GetWindowMouseGrab](https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseGrab) | [:heavy_check_mark:](sdl/methods.go#L4637) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5015) |
| [SDL_GetGrabbedWindow](https://wiki.libsdl.org/SDL3/SDL_GetGrabbedWindow) | [:heavy_check_mark:](sdl/functions.go#L712) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5024) |
| [SDL_SetWindowMouseRect](https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseRect) | [:heavy_check_mark:](sdl/methods.go#L4643) | [:x:](sdl/sdl_functions_js.go#L5031) |
| [SDL_GetWindowMouseRect](https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseRect) | [:heavy_check_mark:](sdl/methods.go#L4653) | [:x:](sdl/sdl_functions_js.go#L5045) |
| [SDL_SetWindowOpacity](https://wiki.libsdl.org/SDL3/SDL_SetWindowOpacity) | [:heavy_check_mark:](sdl/methods.go#L4659) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5055) |
| [SDL_GetWindowOpacity](https://wiki.libsdl.org/SDL3/SDL_GetWindowOpacity) | [:heavy_check_mark:](sdl/methods.go#L4669) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5066) |
| [SDL_SetWindowParent](https://wiki.libsdl.org/SDL3/SDL_SetWindowParent) | [:heavy_check_mark:](sdl/methods.go#L4675) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5075) |
| [SDL_SetWindowModal](https://wiki.libsdl.org/SDL3/SDL_SetWindowModal) | [:heavy_check_mark:](sdl/methods.go#L4685) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5086) |
| [SDL_SetWindowFocusable](https://wiki.libsdl.org/SDL3/SDL_SetWindowFocusable) | [:heavy_check_mark:](sdl/methods.go#L4695) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5097) |
| [SDL_ShowWindowSystemMenu](https://wiki.libsdl.org/SDL3/SDL_ShowWindowSystemMenu) | [:heavy_check_mark:](sdl/methods.go#L4705) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5108) |
| [SDL_SetWindowHitTest](https://wiki.libsdl.org/SDL3/SDL_SetWindowHitTest) | [:heavy_check_mark:](sdl/methods.go#L4715) | [:x:](sdl/sdl_functions_js.go#L5121) |
| [SDL_SetWindowShape](https://wiki.libsdl.org/SDL3/SDL_SetWindowShape) | [:heavy_check_mark:](sdl/methods.go#L4725) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5135) |
| [SDL_FlashWindow](https://wiki.libsdl.org/SDL3/SDL_FlashWindow) | [:heavy_check_mark:](sdl/methods.go#L4735) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5146) |
| [SDL_SetWindowProgressState](https://wiki.libsdl.org/SDL3/SDL_SetWindowProgressState) | [:heavy_check_mark:](sdl/methods.go#L4745) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12530) |
| [SDL_GetWindowProgressState](https://wiki.libsdl.org/SDL3/SDL_GetWindowProgressState) | [:heavy_check_mark:](sdl/methods.go#L4755) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12541) |
| [SDL_SetWindowProgressValue](https://wiki.libsdl.org/SDL3/SDL_SetWindowProgressValue) | [:heavy_check_mark:](sdl/methods.go#L4766) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12550) |
| [SDL_GetWindowProgressValue](https://wiki.libsdl.org/SDL3/SDL_GetWindowProgressValue) | [:heavy_check_mark:](sdl/methods.go#L4776) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12561) |
| [SDL_DestroyWindow](https://wiki.libsdl.org/SDL3/SDL_DestroyWindow) | [:heavy_check_mark:](sdl/methods.go#L4787) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5157) |
| [SDL_ScreenSaverEnabled](https://wiki.libsdl.org/SDL3/SDL_ScreenSaverEnabled) | [:heavy_check_mark:](sdl/functions.go#L718) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5169) |
| [SDL_EnableScreenSaver](https://wiki.libsdl.org/SDL3/SDL_EnableScreenSaver) | [:heavy_check_mark:](sdl/functions.go#L724) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5176) |
| [SDL_DisableScreenSaver](https://wiki.libsdl.org/SDL3/SDL_DisableScreenSaver) | [:heavy_check_mark:](sdl/functions.go#L734) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5183) |
| [SDL_GL_LoadLibrary](https://wiki.libsdl.org/SDL3/SDL_GL_LoadLibrary) | [:heavy_check_mark:](sdl/functions.go#L744) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5190) |
| [SDL_GL_GetProcAddress](https://wiki.libsdl.org/SDL3/SDL_GL_GetProcAddress) | [:heavy_check_mark:](sdl/functions.go#L754) | [:x:](sdl/sdl_functions_js.go#L12570) |
| [SDL_EGL_GetProcAddress](https://wiki.libsdl.org/SDL3/SDL_EGL_GetProcAddress) | [:heavy_check_mark:](sdl/functions.go#L760) | [:x:](sdl/sdl_functions_js.go#L12582) |
| [SDL_GL_UnloadLibrary](https://wiki.libsdl.org/SDL3/SDL_GL_UnloadLibrary) | [:heavy_check_mark:](sdl/functions.go#L772) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5227) |
| [SDL_GL_ExtensionSupported](https://wiki.libsdl.org/SDL3/SDL_GL_ExtensionSupported) | [:heavy_check_mark:](sdl/functions.go#L766) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5233) |
| [SDL_GL_ResetAttributes](https://wiki.libsdl.org/SDL3/SDL_GL_ResetAttributes) | [:heavy_check_mark:](sdl/functions.go#L821) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5244) |
| [SDL_GL_SetAttribute](https://wiki.libsdl.org/SDL3/SDL_GL_SetAttribute) | [:heavy_check_mark:](sdl/functions.go#L827) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5250) |
| [SDL_GL_GetAttribute](https://wiki.libsdl.org/SDL3/SDL_GL_GetAttribute) | [:heavy_check_mark:](sdl/functions.go#L837) | [:x:](sdl/sdl_functions_js.go#L5261) |
| [SDL_GL_CreateContext](https://wiki.libsdl.org/SDL3/SDL_GL_CreateContext) | [:heavy_check_mark:](sdl/functions.go#L778) | [:x:](sdl/sdl_functions_js.go#L12594) |
| [SDL_GL_MakeCurrent](https://wiki.libsdl.org/SDL3/SDL_GL_MakeCurrent) | [:heavy_check_mark:](sdl/functions.go#L789) | [:x:](sdl/sdl_functions_js.go#L12605) |
| [SDL_GL_GetCurrentWindow](https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentWindow) | [:heavy_check_mark:](sdl/functions.go#L849) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5309) |
| [SDL_GL_GetCurrentContext](https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentContext) | [:heavy_check_mark:](sdl/functions.go#L860) | [:x:](sdl/sdl_functions_js.go#L12617) |
| [SDL_EGL_GetCurrentDisplay](https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentDisplay) | [:heavy_check_mark:](sdl/functions.go#L871) | [:x:](sdl/sdl_functions_js.go#L5327) |
| [SDL_EGL_GetCurrentConfig](https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentConfig) | [:heavy_check_mark:](sdl/functions.go#L882) | [:x:](sdl/sdl_functions_js.go#L5335) |
| [SDL_EGL_GetWindowSurface](https://wiki.libsdl.org/SDL3/SDL_EGL_GetWindowSurface) | [:heavy_check_mark:](sdl/functions.go#L799) | [:x:](sdl/sdl_functions_js.go#L5343) |
| [SDL_EGL_SetAttributeCallbacks](https://wiki.libsdl.org/SDL3/SDL_EGL_SetAttributeCallbacks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L12626) |
| [SDL_GL_SetSwapInterval](https://wiki.libsdl.org/SDL3/SDL_GL_SetSwapInterval) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5370) |
| [SDL_GL_GetSwapInterval](https://wiki.libsdl.org/SDL3/SDL_GL_GetSwapInterval) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L5379) |
| [SDL_GL_SwapWindow](https://wiki.libsdl.org/SDL3/SDL_GL_SwapWindow) | [:heavy_check_mark:](sdl/functions.go#L805) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5391) |
| [SDL_GL_DestroyContext](https://wiki.libsdl.org/SDL3/SDL_GL_DestroyContext) | [:heavy_check_mark:](sdl/functions.go#L815) | [:x:](sdl/sdl_functions_js.go#L12641) |
</details>
<details open>
<summary><h3>Events</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_PumpEvents](https://wiki.libsdl.org/SDL3/SDL_PumpEvents) | [:heavy_check_mark:](sdl/functions.go#L185) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7687) |
| [SDL_PeepEvents](https://wiki.libsdl.org/SDL3/SDL_PeepEvents) | [:heavy_check_mark:](sdl/functions.go#L213) | [:x:](sdl/sdl_functions_js.go#L7693) |
| [SDL_HasEvent](https://wiki.libsdl.org/SDL3/SDL_HasEvent) | [:heavy_check_mark:](sdl/functions.go#L239) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7713) |
| [SDL_HasEvents](https://wiki.libsdl.org/SDL3/SDL_HasEvents) | [:heavy_check_mark:](sdl/functions.go#L245) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7722) |
| [SDL_FlushEvent](https://wiki.libsdl.org/SDL3/SDL_FlushEvent) | [:heavy_check_mark:](sdl/functions.go#L251) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7733) |
| [SDL_FlushEvents](https://wiki.libsdl.org/SDL3/SDL_FlushEvents) | [:heavy_check_mark:](sdl/functions.go#L257) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7741) |
| [SDL_PollEvent](https://wiki.libsdl.org/SDL3/SDL_PollEvent) | [:heavy_check_mark:](sdl/functions.go#L263) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7751) |
| [SDL_WaitEvent](https://wiki.libsdl.org/SDL3/SDL_WaitEvent) | [:heavy_check_mark:](sdl/functions.go#L269) | [:x:](sdl/sdl_functions_js.go#L7764) |
| [SDL_WaitEventTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitEventTimeout) | [:heavy_check_mark:](sdl/functions.go#L279) | [:x:](sdl/sdl_functions_js.go#L7776) |
| [SDL_PushEvent](https://wiki.libsdl.org/SDL3/SDL_PushEvent) | [:heavy_check_mark:](sdl/functions.go#L285) | [:x:](sdl/sdl_functions_js.go#L7790) |
| [SDL_SetEventFilter](https://wiki.libsdl.org/SDL3/SDL_SetEventFilter) | [:heavy_check_mark:](sdl/functions.go#L295) | [:x:](sdl/sdl_functions_js.go#L12817) |
| [SDL_GetEventFilter](https://wiki.libsdl.org/SDL3/SDL_GetEventFilter) | [:heavy_check_mark:](sdl/functions.go#L301) | [:x:](sdl/sdl_functions_js.go#L7815) |
| [SDL_AddEventWatch](https://wiki.libsdl.org/SDL3/SDL_AddEventWatch) | [:heavy_check_mark:](sdl/functions.go#L314) | [:x:](sdl/sdl_functions_js.go#L12828) |
| [SDL_RemoveEventWatch](https://wiki.libsdl.org/SDL3/SDL_RemoveEventWatch) | [:heavy_check_mark:](sdl/functions.go#L324) | [:x:](sdl/sdl_functions_js.go#L12840) |
| [SDL_FilterEvents](https://wiki.libsdl.org/SDL3/SDL_FilterEvents) | [:heavy_check_mark:](sdl/functions.go#L330) | [:x:](sdl/sdl_functions_js.go#L12851) |
| [SDL_SetEventEnabled](https://wiki.libsdl.org/SDL3/SDL_SetEventEnabled) | [:heavy_check_mark:](sdl/functions.go#L336) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7870) |
| [SDL_EventEnabled](https://wiki.libsdl.org/SDL3/SDL_EventEnabled) | [:heavy_check_mark:](sdl/functions.go#L342) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7880) |
| [SDL_RegisterEvents](https://wiki.libsdl.org/SDL3/SDL_RegisterEvents) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7889) |
| [SDL_GetWindowFromEvent](https://wiki.libsdl.org/SDL3/SDL_GetWindowFromEvent) | [:heavy_check_mark:](sdl/methods.go#L1964) | [:x:](sdl/sdl_functions_js.go#L7898) |
| [SDL_GetEventDescription](https://wiki.libsdl.org/SDL3/SDL_GetEventDescription) | [:heavy_check_mark:](sdl/methods.go#L1970) | [:x:](sdl/sdl_functions_js.go#L12862) |
</details>
<details open>
<summary><h3>Keyboard</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_HasKeyboard](https://wiki.libsdl.org/SDL3/SDL_HasKeyboard) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7135) |
| [SDL_GetKeyboards](https://wiki.libsdl.org/SDL3/SDL_GetKeyboards) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7143) |
| [SDL_GetKeyboardNameForID](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardNameForID) | [:heavy_check_mark:](sdl/methods.go#L3821) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7155) |
| [SDL_GetKeyboardFocus](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardFocus) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7164) |
| [SDL_GetKeyboardState](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L12775) |
| [SDL_ResetKeyboard](https://wiki.libsdl.org/SDL3/SDL_ResetKeyboard) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7190) |
| [SDL_GetModState](https://wiki.libsdl.org/SDL3/SDL_GetModState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7196) |
| [SDL_SetModState](https://wiki.libsdl.org/SDL3/SDL_SetModState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7204) |
| [SDL_GetKeyFromScancode](https://wiki.libsdl.org/SDL3/SDL_GetKeyFromScancode) | [:heavy_check_mark:](sdl/methods.go#L4916) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7211) |
| [SDL_GetScancodeFromKey](https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromKey) | [:heavy_check_mark:](sdl/methods.go#L6009) | [:x:](sdl/sdl_functions_js.go#L7224) |
| [SDL_SetScancodeName](https://wiki.libsdl.org/SDL3/SDL_SetScancodeName) | [:heavy_check_mark:](sdl/methods.go#L4922) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7238) |
| [SDL_GetScancodeName](https://wiki.libsdl.org/SDL3/SDL_GetScancodeName) | [:heavy_check_mark:](sdl/methods.go#L4932) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7251) |
| [SDL_GetScancodeFromName](https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromName) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7260) |
| [SDL_GetKeyName](https://wiki.libsdl.org/SDL3/SDL_GetKeyName) | [:heavy_check_mark:](sdl/methods.go#L6015) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7272) |
| [SDL_GetKeyFromName](https://wiki.libsdl.org/SDL3/SDL_GetKeyFromName) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7282) |
| [SDL_StartTextInput](https://wiki.libsdl.org/SDL3/SDL_StartTextInput) | [:heavy_check_mark:](sdl/methods.go#L4793) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7294) |
| [SDL_StartTextInputWithProperties](https://wiki.libsdl.org/SDL3/SDL_StartTextInputWithProperties) | [:heavy_check_mark:](sdl/methods.go#L4803) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7307) |
| [SDL_TextInputActive](https://wiki.libsdl.org/SDL3/SDL_TextInputActive) | [:heavy_check_mark:](sdl/methods.go#L4813) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7321) |
| [SDL_StopTextInput](https://wiki.libsdl.org/SDL3/SDL_StopTextInput) | [:heavy_check_mark:](sdl/methods.go#L4819) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7334) |
| [SDL_ClearComposition](https://wiki.libsdl.org/SDL3/SDL_ClearComposition) | [:heavy_check_mark:](sdl/methods.go#L4829) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7347) |
| [SDL_SetTextInputArea](https://wiki.libsdl.org/SDL3/SDL_SetTextInputArea) | [:heavy_check_mark:](sdl/methods.go#L4839) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7360) |
| [SDL_GetTextInputArea](https://wiki.libsdl.org/SDL3/SDL_GetTextInputArea) | [:heavy_check_mark:](sdl/methods.go#L4849) | [:x:](sdl/sdl_functions_js.go#L7378) |
| [SDL_HasScreenKeyboardSupport](https://wiki.libsdl.org/SDL3/SDL_HasScreenKeyboardSupport) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7394) |
| [SDL_ScreenKeyboardShown](https://wiki.libsdl.org/SDL3/SDL_ScreenKeyboardShown) | [:heavy_check_mark:](sdl/methods.go#L4862) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7402) |
</details>
<details open>
<summary><h3>Mouse</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_HasMouse](https://wiki.libsdl.org/SDL3/SDL_HasMouse) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7415) |
| [SDL_GetMice](https://wiki.libsdl.org/SDL3/SDL_GetMice) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7422) |
| [SDL_GetMouseNameForID](https://wiki.libsdl.org/SDL3/SDL_GetMouseNameForID) | [:heavy_check_mark:](sdl/methods.go#L3834) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7434) |
| [SDL_GetMouseFocus](https://wiki.libsdl.org/SDL3/SDL_GetMouseFocus) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7443) |
| [SDL_GetMouseState](https://wiki.libsdl.org/SDL3/SDL_GetMouseState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7450) |
| [SDL_GetGlobalMouseState](https://wiki.libsdl.org/SDL3/SDL_GetGlobalMouseState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7466) |
| [SDL_GetRelativeMouseState](https://wiki.libsdl.org/SDL3/SDL_GetRelativeMouseState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7480) |
| [SDL_WarpMouseInWindow](https://wiki.libsdl.org/SDL3/SDL_WarpMouseInWindow) | [:heavy_check_mark:](sdl/methods.go#L4868) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7494) |
| [SDL_WarpMouseGlobal](https://wiki.libsdl.org/SDL3/SDL_WarpMouseGlobal) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7506) |
| [SDL_SetRelativeMouseTransform](https://wiki.libsdl.org/SDL3/SDL_SetRelativeMouseTransform) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L12787) |
| [SDL_SetWindowRelativeMouseMode](https://wiki.libsdl.org/SDL3/SDL_SetWindowRelativeMouseMode) | [:heavy_check_mark:](sdl/methods.go#L4874) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7517) |
| [SDL_GetWindowRelativeMouseMode](https://wiki.libsdl.org/SDL3/SDL_GetWindowRelativeMouseMode) | [:heavy_check_mark:](sdl/methods.go#L4884) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7528) |
| [SDL_CaptureMouse](https://wiki.libsdl.org/SDL3/SDL_CaptureMouse) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7537) |
| [SDL_CreateCursor](https://wiki.libsdl.org/SDL3/SDL_CreateCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7546) |
| [SDL_CreateColorCursor](https://wiki.libsdl.org/SDL3/SDL_CreateColorCursor) | [:heavy_check_mark:](sdl/methods.go#L1940) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7568) |
| [SDL_CreateAnimatedCursor](https://wiki.libsdl.org/SDL3/SDL_CreateAnimatedCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L12799) |
| [SDL_CreateSystemCursor](https://wiki.libsdl.org/SDL3/SDL_CreateSystemCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7581) |
| [SDL_SetCursor](https://wiki.libsdl.org/SDL3/SDL_SetCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7590) |
| [SDL_GetCursor](https://wiki.libsdl.org/SDL3/SDL_GetCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7599) |
| [SDL_GetDefaultCursor](https://wiki.libsdl.org/SDL3/SDL_GetDefaultCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7606) |
| [SDL_DestroyCursor](https://wiki.libsdl.org/SDL3/SDL_DestroyCursor) | [:heavy_check_mark:](sdl/methods.go#L663) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7613) |
| [SDL_ShowCursor](https://wiki.libsdl.org/SDL3/SDL_ShowCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7622) |
| [SDL_HideCursor](https://wiki.libsdl.org/SDL3/SDL_HideCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7629) |
| [SDL_CursorVisible](https://wiki.libsdl.org/SDL3/SDL_CursorVisible) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7636) |
</details>
<details open>
<summary><h3>Touch</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetTouchDevices](https://wiki.libsdl.org/SDL3/SDL_GetTouchDevices) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7643) |
| [SDL_GetTouchDeviceName](https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceName) | [:heavy_check_mark:](sdl/methods.go#L45) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7655) |
| [SDL_GetTouchDeviceType](https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceType) | [:heavy_check_mark:](sdl/methods.go#L56) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7664) |
| [SDL_GetTouchFingers](https://wiki.libsdl.org/SDL3/SDL_GetTouchFingers) | [:heavy_check_mark:](sdl/methods.go#L62) | [:x:](sdl/sdl_functions_js.go#L7673) |
</details>
<details open>
<summary><h3>Gamepad</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AddGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMapping) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6368) |
| [SDL_AddGamepadMappingsFromIO](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromIO) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L6379) |
| [SDL_AddGamepadMappingsFromFile](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromFile) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6390) |
| [SDL_ReloadGamepadMappings](https://wiki.libsdl.org/SDL3/SDL_ReloadGamepadMappings) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6401) |
| [SDL_GetGamepadMappings](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappings) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L6408) |
| [SDL_GetGamepadMappingForGUID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForGUID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12754) |
| [SDL_GetGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMapping) | [:heavy_check_mark:](sdl/methods.go#L2509) | [:x:](sdl/sdl_functions_js.go#L6433) |
| [SDL_SetGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_SetGamepadMapping) | [:heavy_check_mark:](sdl/methods.go#L800) | [:x:](sdl/sdl_functions_js.go#L6443) |
| [SDL_HasGamepad](https://wiki.libsdl.org/SDL3/SDL_HasGamepad) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6457) |
| [SDL_GetGamepads](https://wiki.libsdl.org/SDL3/SDL_GetGamepads) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L6464) |
| [SDL_IsGamepad](https://wiki.libsdl.org/SDL3/SDL_IsGamepad) | [:heavy_check_mark:](sdl/methods.go#L812) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6476) |
| [SDL_GetGamepadNameForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadNameForID) | [:heavy_check_mark:](sdl/methods.go#L818) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6485) |
| [SDL_GetGamepadPathForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPathForID) | [:heavy_check_mark:](sdl/methods.go#L829) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6494) |
| [SDL_GetGamepadPlayerIndexForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndexForID) | [:heavy_check_mark:](sdl/methods.go#L840) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6503) |
| [SDL_GetGamepadGUIDForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadGUIDForID) | [:x:](sdl/methods.go#L846) | [:x:](sdl/sdl_functions_js.go#L12764) |
| [SDL_GetGamepadVendorForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendorForID) | [:heavy_check_mark:](sdl/methods.go#L853) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6525) |
| [SDL_GetGamepadProductForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductForID) | [:heavy_check_mark:](sdl/methods.go#L859) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6534) |
| [SDL_GetGamepadProductVersionForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersionForID) | [:heavy_check_mark:](sdl/methods.go#L865) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6543) |
| [SDL_GetGamepadTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeForID) | [:heavy_check_mark:](sdl/methods.go#L871) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6552) |
| [SDL_GetRealGamepadTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadTypeForID) | [:heavy_check_mark:](sdl/methods.go#L877) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6561) |
| [SDL_GetGamepadMappingForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForID) | [:heavy_check_mark:](sdl/methods.go#L883) | [:x:](sdl/sdl_functions_js.go#L6570) |
| [SDL_OpenGamepad](https://wiki.libsdl.org/SDL3/SDL_OpenGamepad) | [:heavy_check_mark:](sdl/methods.go#L895) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6580) |
| [SDL_GetGamepadFromID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromID) | [:heavy_check_mark:](sdl/methods.go#L906) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6589) |
| [SDL_GetGamepadFromPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromPlayerIndex) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6598) |
| [SDL_GetGamepadProperties](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProperties) | [:heavy_check_mark:](sdl/methods.go#L2521) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6607) |
| [SDL_GetGamepadID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadID) | [:heavy_check_mark:](sdl/methods.go#L2532) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6616) |
| [SDL_GetGamepadName](https://wiki.libsdl.org/SDL3/SDL_GetGamepadName) | [:heavy_check_mark:](sdl/methods.go#L2543) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6625) |
| [SDL_GetGamepadPath](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPath) | [:heavy_check_mark:](sdl/methods.go#L2549) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6634) |
| [SDL_GetGamepadType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadType) | [:heavy_check_mark:](sdl/methods.go#L2555) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6643) |
| [SDL_GetRealGamepadType](https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadType) | [:heavy_check_mark:](sdl/methods.go#L2561) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6652) |
| [SDL_GetGamepadPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L2567) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6661) |
| [SDL_SetGamepadPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_SetGamepadPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L2573) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6670) |
| [SDL_GetGamepadVendor](https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendor) | [:heavy_check_mark:](sdl/methods.go#L2583) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6681) |
| [SDL_GetGamepadProduct](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProduct) | [:heavy_check_mark:](sdl/methods.go#L2589) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6690) |
| [SDL_GetGamepadProductVersion](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersion) | [:heavy_check_mark:](sdl/methods.go#L2595) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6699) |
| [SDL_GetGamepadFirmwareVersion](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFirmwareVersion) | [:heavy_check_mark:](sdl/methods.go#L2601) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6708) |
| [SDL_GetGamepadSerial](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSerial) | [:heavy_check_mark:](sdl/methods.go#L2607) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6717) |
| [SDL_GetGamepadSteamHandle](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSteamHandle) | [:heavy_check_mark:](sdl/methods.go#L2613) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6726) |
| [SDL_GetGamepadConnectionState](https://wiki.libsdl.org/SDL3/SDL_GetGamepadConnectionState) | [:heavy_check_mark:](sdl/methods.go#L2619) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6735) |
| [SDL_GetGamepadPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPowerInfo) | [:heavy_check_mark:](sdl/methods.go#L2631) | [:x:](sdl/sdl_functions_js.go#L6744) |
| [SDL_GamepadConnected](https://wiki.libsdl.org/SDL3/SDL_GamepadConnected) | [:heavy_check_mark:](sdl/methods.go#L2641) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6758) |
| [SDL_GetGamepadJoystick](https://wiki.libsdl.org/SDL3/SDL_GetGamepadJoystick) | [:heavy_check_mark:](sdl/methods.go#L2647) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6767) |
| [SDL_SetGamepadEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_SetGamepadEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6776) |
| [SDL_GamepadEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_GamepadEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6784) |
| [SDL_GetGamepadBindings](https://wiki.libsdl.org/SDL3/SDL_GetGamepadBindings) | [:heavy_check_mark:](sdl/methods.go#L2658) | [:x:](sdl/sdl_functions_js.go#L6791) |
| [SDL_UpdateGamepads](https://wiki.libsdl.org/SDL3/SDL_UpdateGamepads) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6805) |
| [SDL_GetGamepadTypeFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeFromString) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6811) |
| [SDL_GetGamepadStringForType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForType) | [:heavy_check_mark:](sdl/methods.go#L5736) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6822) |
| [SDL_GetGamepadAxisFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxisFromString) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6831) |
| [SDL_GetGamepadStringForAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForAxis) | [:heavy_check_mark:](sdl/methods.go#L655) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6842) |
| [SDL_GamepadHasAxis](https://wiki.libsdl.org/SDL3/SDL_GamepadHasAxis) | [:heavy_check_mark:](sdl/methods.go#L2672) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6851) |
| [SDL_GetGamepadAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxis) | [:heavy_check_mark:](sdl/methods.go#L2678) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6862) |
| [SDL_GetGamepadButtonFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonFromString) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6873) |
| [SDL_GetGamepadStringForButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForButton) | [:heavy_check_mark:](sdl/methods.go#L550) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6884) |
| [SDL_GamepadHasButton](https://wiki.libsdl.org/SDL3/SDL_GamepadHasButton) | [:heavy_check_mark:](sdl/methods.go#L2684) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6893) |
| [SDL_GetGamepadButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButton) | [:heavy_check_mark:](sdl/methods.go#L2690) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6904) |
| [SDL_GetGamepadButtonLabelForType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabelForType) | [:heavy_check_mark:](sdl/methods.go#L5742) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6915) |
| [SDL_GetGamepadButtonLabel](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabel) | [:heavy_check_mark:](sdl/methods.go#L2696) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6926) |
| [SDL_GetNumGamepadTouchpads](https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpads) | [:heavy_check_mark:](sdl/methods.go#L2702) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6937) |
| [SDL_GetNumGamepadTouchpadFingers](https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpadFingers) | [:heavy_check_mark:](sdl/methods.go#L2708) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6946) |
| [SDL_GetGamepadTouchpadFinger](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTouchpadFinger) | [:x:](sdl/methods.go#L2714) | [:x:](sdl/sdl_functions_js.go#L6957) |
| [SDL_GamepadHasSensor](https://wiki.libsdl.org/SDL3/SDL_GamepadHasSensor) | [:heavy_check_mark:](sdl/methods.go#L2721) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6981) |
| [SDL_SetGamepadSensorEnabled](https://wiki.libsdl.org/SDL3/SDL_SetGamepadSensorEnabled) | [:heavy_check_mark:](sdl/methods.go#L2727) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6992) |
| [SDL_GamepadSensorEnabled](https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEnabled) | [:heavy_check_mark:](sdl/methods.go#L2737) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7005) |
| [SDL_GetGamepadSensorDataRate](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorDataRate) | [:heavy_check_mark:](sdl/methods.go#L2743) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7016) |
| [SDL_GetGamepadSensorData](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorData) | [:x:](sdl/methods.go#L2749) | [:x:](sdl/sdl_functions_js.go#L7027) |
| [SDL_GamepadHasCapSense](https://wiki.libsdl.org/SDL3/SDL_GamepadHasCapSense) | [:question:]() | [:question:]() |
| [SDL_GetGamepadCapSense](https://wiki.libsdl.org/SDL3/SDL_GetGamepadCapSense) | [:question:]() | [:question:]() |
| [SDL_RumbleGamepad](https://wiki.libsdl.org/SDL3/SDL_RumbleGamepad) | [:heavy_check_mark:](sdl/methods.go#L2756) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7045) |
| [SDL_RumbleGamepadTriggers](https://wiki.libsdl.org/SDL3/SDL_RumbleGamepadTriggers) | [:heavy_check_mark:](sdl/methods.go#L2766) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7060) |
| [SDL_SetGamepadLED](https://wiki.libsdl.org/SDL3/SDL_SetGamepadLED) | [:heavy_check_mark:](sdl/methods.go#L2776) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7075) |
| [SDL_SendGamepadEffect](https://wiki.libsdl.org/SDL3/SDL_SendGamepadEffect) | [:heavy_check_mark:](sdl/methods.go#L2786) | [:x:](sdl/sdl_functions_js.go#L7090) |
| [SDL_CloseGamepad](https://wiki.libsdl.org/SDL3/SDL_CloseGamepad) | [:heavy_check_mark:](sdl/methods.go#L2797) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7104) |
| [SDL_GetGamepadAppleSFSymbolsNameForButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForButton) | [:heavy_check_mark:](sdl/methods.go#L2803) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7113) |
| [SDL_GetGamepadAppleSFSymbolsNameForAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForAxis) | [:heavy_check_mark:](sdl/methods.go#L2809) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7124) |
</details>
<details open>
<summary><h3>Joystick</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_LockJoysticks](https://wiki.libsdl.org/SDL3/SDL_LockJoysticks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5683) |
| [SDL_TryLockJoysticks](https://wiki.libsdl.org/SDL3/SDL_TryLockJoysticks) | [:question:]() | [:question:]() |
| [SDL_UnlockJoysticks](https://wiki.libsdl.org/SDL3/SDL_UnlockJoysticks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5689) |
| [SDL_HasJoystick](https://wiki.libsdl.org/SDL3/SDL_HasJoystick) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5695) |
| [SDL_GetJoysticks](https://wiki.libsdl.org/SDL3/SDL_GetJoysticks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L5702) |
| [SDL_GetJoystickNameForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickNameForID) | [:heavy_check_mark:](sdl/methods.go#L703) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5714) |
| [SDL_GetJoystickPathForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPathForID) | [:heavy_check_mark:](sdl/methods.go#L714) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5723) |
| [SDL_GetJoystickPlayerIndexForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndexForID) | [:heavy_check_mark:](sdl/methods.go#L725) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5732) |
| [SDL_GetJoystickGUIDForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDForID) | [:x:](sdl/methods.go#L731) | [:x:](sdl/sdl_functions_js.go#L12713) |
| [SDL_GetJoystickVendorForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendorForID) | [:heavy_check_mark:](sdl/methods.go#L738) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5754) |
| [SDL_GetJoystickProductForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductForID) | [:heavy_check_mark:](sdl/methods.go#L744) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5763) |
| [SDL_GetJoystickProductVersionForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersionForID) | [:heavy_check_mark:](sdl/methods.go#L750) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5772) |
| [SDL_GetJoystickTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickTypeForID) | [:heavy_check_mark:](sdl/methods.go#L756) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5781) |
| [SDL_OpenJoystick](https://wiki.libsdl.org/SDL3/SDL_OpenJoystick) | [:heavy_check_mark:](sdl/methods.go#L762) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5790) |
| [SDL_GetJoystickFromID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromID) | [:heavy_check_mark:](sdl/methods.go#L773) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5802) |
| [SDL_GetJoystickFromPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromPlayerIndex) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5816) |
| [SDL_AttachVirtualJoystick](https://wiki.libsdl.org/SDL3/SDL_AttachVirtualJoystick) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L5825) |
| [SDL_DetachVirtualJoystick](https://wiki.libsdl.org/SDL3/SDL_DetachVirtualJoystick) | [:heavy_check_mark:](sdl/methods.go#L784) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5838) |
| [SDL_IsJoystickVirtual](https://wiki.libsdl.org/SDL3/SDL_IsJoystickVirtual) | [:heavy_check_mark:](sdl/methods.go#L794) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5847) |
| [SDL_SetJoystickVirtualAxis](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualAxis) | [:heavy_check_mark:](sdl/methods.go#L5394) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5856) |
| [SDL_SetJoystickVirtualBall](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualBall) | [:heavy_check_mark:](sdl/methods.go#L5404) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5869) |
| [SDL_SetJoystickVirtualButton](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualButton) | [:heavy_check_mark:](sdl/methods.go#L5414) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5884) |
| [SDL_SetJoystickVirtualHat](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualHat) | [:heavy_check_mark:](sdl/methods.go#L5424) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5897) |
| [SDL_SetJoystickVirtualTouchpad](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualTouchpad) | [:heavy_check_mark:](sdl/methods.go#L5434) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5910) |
| [SDL_SendJoystickVirtualSensorData](https://wiki.libsdl.org/SDL3/SDL_SendJoystickVirtualSensorData) | [:heavy_check_mark:](sdl/methods.go#L5444) | [:x:](sdl/sdl_functions_js.go#L5931) |
| [SDL_GetJoystickProperties](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProperties) | [:heavy_check_mark:](sdl/methods.go#L5454) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5951) |
| [SDL_GetJoystickName](https://wiki.libsdl.org/SDL3/SDL_GetJoystickName) | [:heavy_check_mark:](sdl/methods.go#L5465) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5960) |
| [SDL_GetJoystickPath](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPath) | [:heavy_check_mark:](sdl/methods.go#L5476) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5973) |
| [SDL_GetJoystickPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L5487) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5982) |
| [SDL_SetJoystickPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_SetJoystickPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L5493) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5991) |
| [SDL_GetJoystickGUID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUID) | [:x:](sdl/methods.go#L5503) | [:x:](sdl/sdl_functions_js.go#L12724) |
| [SDL_GetJoystickVendor](https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendor) | [:heavy_check_mark:](sdl/methods.go#L5510) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6018) |
| [SDL_GetJoystickProduct](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProduct) | [:heavy_check_mark:](sdl/methods.go#L5516) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6027) |
| [SDL_GetJoystickProductVersion](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersion) | [:heavy_check_mark:](sdl/methods.go#L5522) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6036) |
| [SDL_GetJoystickFirmwareVersion](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFirmwareVersion) | [:heavy_check_mark:](sdl/methods.go#L5528) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6045) |
| [SDL_GetJoystickSerial](https://wiki.libsdl.org/SDL3/SDL_GetJoystickSerial) | [:heavy_check_mark:](sdl/methods.go#L5534) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6054) |
| [SDL_GetJoystickType](https://wiki.libsdl.org/SDL3/SDL_GetJoystickType) | [:heavy_check_mark:](sdl/methods.go#L5540) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6063) |
| [SDL_GetJoystickGUIDInfo](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDInfo) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12735) |
| [SDL_JoystickConnected](https://wiki.libsdl.org/SDL3/SDL_JoystickConnected) | [:heavy_check_mark:](sdl/methods.go#L5546) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6103) |
| [SDL_GetJoystickID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickID) | [:heavy_check_mark:](sdl/methods.go#L5552) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6112) |
| [SDL_GetNumJoystickAxes](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickAxes) | [:heavy_check_mark:](sdl/methods.go#L5563) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6125) |
| [SDL_GetNumJoystickBalls](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickBalls) | [:heavy_check_mark:](sdl/methods.go#L5574) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6138) |
| [SDL_GetNumJoystickHats](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickHats) | [:heavy_check_mark:](sdl/methods.go#L5585) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6147) |
| [SDL_GetNumJoystickButtons](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickButtons) | [:heavy_check_mark:](sdl/methods.go#L5596) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6160) |
| [SDL_SetJoystickEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_SetJoystickEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6173) |
| [SDL_JoystickEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_JoystickEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6181) |
| [SDL_UpdateJoysticks](https://wiki.libsdl.org/SDL3/SDL_UpdateJoysticks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6188) |
| [SDL_GetJoystickAxis](https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxis) | [:heavy_check_mark:](sdl/methods.go#L5607) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6194) |
| [SDL_GetJoystickAxisInitialState](https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxisInitialState) | [:heavy_check_mark:](sdl/methods.go#L5618) | [:x:](sdl/sdl_functions_js.go#L6209) |
| [SDL_GetJoystickBall](https://wiki.libsdl.org/SDL3/SDL_GetJoystickBall) | [:heavy_check_mark:](sdl/methods.go#L5627) | [:x:](sdl/sdl_functions_js.go#L6225) |
| [SDL_GetJoystickHat](https://wiki.libsdl.org/SDL3/SDL_GetJoystickHat) | [:heavy_check_mark:](sdl/methods.go#L5639) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6243) |
| [SDL_GetJoystickButton](https://wiki.libsdl.org/SDL3/SDL_GetJoystickButton) | [:heavy_check_mark:](sdl/methods.go#L5645) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6258) |
| [SDL_RumbleJoystick](https://wiki.libsdl.org/SDL3/SDL_RumbleJoystick) | [:heavy_check_mark:](sdl/methods.go#L5651) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6273) |
| [SDL_RumbleJoystickTriggers](https://wiki.libsdl.org/SDL3/SDL_RumbleJoystickTriggers) | [:heavy_check_mark:](sdl/methods.go#L5657) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6288) |
| [SDL_SetJoystickLED](https://wiki.libsdl.org/SDL3/SDL_SetJoystickLED) | [:heavy_check_mark:](sdl/methods.go#L5667) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6303) |
| [SDL_SendJoystickEffect](https://wiki.libsdl.org/SDL3/SDL_SendJoystickEffect) | [:heavy_check_mark:](sdl/methods.go#L5677) | [:x:](sdl/sdl_functions_js.go#L6318) |
| [SDL_CloseJoystick](https://wiki.libsdl.org/SDL3/SDL_CloseJoystick) | [:heavy_check_mark:](sdl/methods.go#L5687) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6332) |
| [SDL_GetJoystickConnectionState](https://wiki.libsdl.org/SDL3/SDL_GetJoystickConnectionState) | [:heavy_check_mark:](sdl/methods.go#L5693) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6345) |
| [SDL_GetJoystickPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPowerInfo) | [:heavy_check_mark:](sdl/methods.go#L5704) | [:x:](sdl/sdl_functions_js.go#L6354) |
</details>
<details open>
<summary><h3>Haptic</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetHaptics](https://wiki.libsdl.org/SDL3/SDL_GetHaptics) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9245) |
| [SDL_GetHapticNameForID](https://wiki.libsdl.org/SDL3/SDL_GetHapticNameForID) | [:heavy_check_mark:](sdl/methods.go#L2861) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9257) |
| [SDL_OpenHaptic](https://wiki.libsdl.org/SDL3/SDL_OpenHaptic) | [:heavy_check_mark:](sdl/methods.go#L2872) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9266) |
| [SDL_GetHapticFromID](https://wiki.libsdl.org/SDL3/SDL_GetHapticFromID) | [:heavy_check_mark:](sdl/methods.go#L2883) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9275) |
| [SDL_GetHapticID](https://wiki.libsdl.org/SDL3/SDL_GetHapticID) | [:heavy_check_mark:](sdl/methods.go#L2290) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9284) |
| [SDL_GetHapticName](https://wiki.libsdl.org/SDL3/SDL_GetHapticName) | [:heavy_check_mark:](sdl/methods.go#L2301) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9293) |
| [SDL_IsMouseHaptic](https://wiki.libsdl.org/SDL3/SDL_IsMouseHaptic) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9302) |
| [SDL_OpenHapticFromMouse](https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromMouse) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9309) |
| [SDL_IsJoystickHaptic](https://wiki.libsdl.org/SDL3/SDL_IsJoystickHaptic) | [:heavy_check_mark:](sdl/methods.go#L5717) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9316) |
| [SDL_OpenHapticFromJoystick](https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromJoystick) | [:heavy_check_mark:](sdl/methods.go#L5723) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9325) |
| [SDL_CloseHaptic](https://wiki.libsdl.org/SDL3/SDL_CloseHaptic) | [:heavy_check_mark:](sdl/methods.go#L2312) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9334) |
| [SDL_GetMaxHapticEffects](https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffects) | [:heavy_check_mark:](sdl/methods.go#L2318) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9343) |
| [SDL_GetMaxHapticEffectsPlaying](https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffectsPlaying) | [:heavy_check_mark:](sdl/methods.go#L2329) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9352) |
| [SDL_GetHapticFeatures](https://wiki.libsdl.org/SDL3/SDL_GetHapticFeatures) | [:heavy_check_mark:](sdl/methods.go#L2340) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9361) |
| [SDL_GetNumHapticAxes](https://wiki.libsdl.org/SDL3/SDL_GetNumHapticAxes) | [:heavy_check_mark:](sdl/methods.go#L2351) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9370) |
| [SDL_HapticEffectSupported](https://wiki.libsdl.org/SDL3/SDL_HapticEffectSupported) | [:heavy_check_mark:](sdl/methods.go#L2362) | [:x:](sdl/sdl_functions_js.go#L9379) |
| [SDL_CreateHapticEffect](https://wiki.libsdl.org/SDL3/SDL_CreateHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2368) | [:x:](sdl/sdl_functions_js.go#L9393) |
| [SDL_UpdateHapticEffect](https://wiki.libsdl.org/SDL3/SDL_UpdateHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2379) | [:x:](sdl/sdl_functions_js.go#L9407) |
| [SDL_RunHapticEffect](https://wiki.libsdl.org/SDL3/SDL_RunHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2389) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9423) |
| [SDL_StopHapticEffect](https://wiki.libsdl.org/SDL3/SDL_StopHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2399) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9436) |
| [SDL_DestroyHapticEffect](https://wiki.libsdl.org/SDL3/SDL_DestroyHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2409) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9447) |
| [SDL_GetHapticEffectStatus](https://wiki.libsdl.org/SDL3/SDL_GetHapticEffectStatus) | [:heavy_check_mark:](sdl/methods.go#L2415) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9458) |
| [SDL_SetHapticGain](https://wiki.libsdl.org/SDL3/SDL_SetHapticGain) | [:heavy_check_mark:](sdl/methods.go#L2421) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9469) |
| [SDL_SetHapticAutocenter](https://wiki.libsdl.org/SDL3/SDL_SetHapticAutocenter) | [:heavy_check_mark:](sdl/methods.go#L2431) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9480) |
| [SDL_PauseHaptic](https://wiki.libsdl.org/SDL3/SDL_PauseHaptic) | [:heavy_check_mark:](sdl/methods.go#L2441) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9491) |
| [SDL_ResumeHaptic](https://wiki.libsdl.org/SDL3/SDL_ResumeHaptic) | [:heavy_check_mark:](sdl/methods.go#L2451) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9500) |
| [SDL_StopHapticEffects](https://wiki.libsdl.org/SDL3/SDL_StopHapticEffects) | [:heavy_check_mark:](sdl/methods.go#L2461) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9509) |
| [SDL_HapticRumbleSupported](https://wiki.libsdl.org/SDL3/SDL_HapticRumbleSupported) | [:heavy_check_mark:](sdl/methods.go#L2471) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9518) |
| [SDL_InitHapticRumble](https://wiki.libsdl.org/SDL3/SDL_InitHapticRumble) | [:heavy_check_mark:](sdl/methods.go#L2477) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9527) |
| [SDL_PlayHapticRumble](https://wiki.libsdl.org/SDL3/SDL_PlayHapticRumble) | [:heavy_check_mark:](sdl/methods.go#L2487) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9536) |
| [SDL_StopHapticRumble](https://wiki.libsdl.org/SDL3/SDL_StopHapticRumble) | [:heavy_check_mark:](sdl/methods.go#L2497) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9549) |
</details>
<details open>
<summary><h3>Audio</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumAudioDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumAudioDrivers) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1656) |
| [SDL_GetAudioDriver](https://wiki.libsdl.org/SDL3/SDL_GetAudioDriver) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1663) |
| [SDL_GetCurrentAudioDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentAudioDriver) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1672) |
| [SDL_GetAudioPlaybackDevices](https://wiki.libsdl.org/SDL3/SDL_GetAudioPlaybackDevices) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L1679) |
| [SDL_GetAudioRecordingDevices](https://wiki.libsdl.org/SDL3/SDL_GetAudioRecordingDevices) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L1691) |
| [SDL_GetAudioDeviceName](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceName) | [:heavy_check_mark:](sdl/methods.go#L211) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1703) |
| [SDL_GetAudioDeviceFormat](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceFormat) | [:heavy_check_mark:](sdl/methods.go#L222) | [:x:](sdl/sdl_functions_js.go#L1712) |
| [SDL_GetAudioDeviceChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceChannelMap) | [:heavy_check_mark:](sdl/methods.go#L235) | [:x:](sdl/sdl_functions_js.go#L1728) |
| [SDL_OpenAudioDevice](https://wiki.libsdl.org/SDL3/SDL_OpenAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L249) | [:x:](sdl/sdl_functions_js.go#L1742) |
| [SDL_IsAudioDevicePhysical](https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePhysical) | [:heavy_check_mark:](sdl/methods.go#L260) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1756) |
| [SDL_IsAudioDevicePlayback](https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePlayback) | [:heavy_check_mark:](sdl/methods.go#L266) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1765) |
| [SDL_PauseAudioDevice](https://wiki.libsdl.org/SDL3/SDL_PauseAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L272) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1774) |
| [SDL_ResumeAudioDevice](https://wiki.libsdl.org/SDL3/SDL_ResumeAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L282) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1783) |
| [SDL_AudioDevicePaused](https://wiki.libsdl.org/SDL3/SDL_AudioDevicePaused) | [:heavy_check_mark:](sdl/methods.go#L292) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1792) |
| [SDL_GetAudioDeviceGain](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceGain) | [:heavy_check_mark:](sdl/methods.go#L298) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1801) |
| [SDL_SetAudioDeviceGain](https://wiki.libsdl.org/SDL3/SDL_SetAudioDeviceGain) | [:heavy_check_mark:](sdl/methods.go#L309) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1810) |
| [SDL_CloseAudioDevice](https://wiki.libsdl.org/SDL3/SDL_CloseAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L319) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1821) |
| [SDL_BindAudioStreams](https://wiki.libsdl.org/SDL3/SDL_BindAudioStreams) | [:heavy_check_mark:](sdl/methods.go#L325) | [:x:](sdl/sdl_functions_js.go#L1829) |
| [SDL_BindAudioStream](https://wiki.libsdl.org/SDL3/SDL_BindAudioStream) | [:heavy_check_mark:](sdl/methods.go#L335) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1845) |
| [SDL_UnbindAudioStreams](https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStreams) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L1856) |
| [SDL_UnbindAudioStream](https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStream) | [:heavy_check_mark:](sdl/methods.go#L3849) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1869) |
| [SDL_GetAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamDevice) | [:heavy_check_mark:](sdl/methods.go#L3855) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1877) |
| [SDL_CreateAudioStream](https://wiki.libsdl.org/SDL3/SDL_CreateAudioStream) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L1886) |
| [SDL_GetAudioStreamProperties](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamProperties) | [:heavy_check_mark:](sdl/methods.go#L3861) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1900) |
| [SDL_GetAudioStreamFormat](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFormat) | [:heavy_check_mark:](sdl/methods.go#L3872) | [:x:](sdl/sdl_functions_js.go#L1909) |
| [SDL_SetAudioStreamFormat](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFormat) | [:heavy_check_mark:](sdl/methods.go#L3882) | [:x:](sdl/sdl_functions_js.go#L1925) |
| [SDL_GetAudioStreamFrequencyRatio](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFrequencyRatio) | [:heavy_check_mark:](sdl/methods.go#L3892) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1941) |
| [SDL_SetAudioStreamFrequencyRatio](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFrequencyRatio) | [:heavy_check_mark:](sdl/methods.go#L3903) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1950) |
| [SDL_GetAudioStreamGain](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamGain) | [:heavy_check_mark:](sdl/methods.go#L3913) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1961) |
| [SDL_SetAudioStreamGain](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGain) | [:heavy_check_mark:](sdl/methods.go#L3924) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1970) |
| [SDL_GetAudioStreamInputChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamInputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3934) | [:x:](sdl/sdl_functions_js.go#L1981) |
| [SDL_GetAudioStreamOutputChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamOutputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3948) | [:x:](sdl/sdl_functions_js.go#L1995) |
| [SDL_SetAudioStreamInputChannelMap](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamInputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3962) | [:x:](sdl/sdl_functions_js.go#L2009) |
| [SDL_SetAudioStreamOutputChannelMap](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamOutputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3972) | [:x:](sdl/sdl_functions_js.go#L2025) |
| [SDL_PutAudioStreamData](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData) | [:heavy_check_mark:](sdl/methods.go#L3982) | [:x:](sdl/sdl_functions_js.go#L2041) |
| [SDL_PutAudioStreamDataNoCopy](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamDataNoCopy) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12315) |
| [SDL_PutAudioStreamPlanarData](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamPlanarData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12333) |
| [SDL_GetAudioStreamData](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData) | [:heavy_check_mark:](sdl/methods.go#L3994) | [:x:](sdl/sdl_functions_js.go#L2055) |
| [SDL_GetAudioStreamAvailable](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamAvailable) | [:heavy_check_mark:](sdl/methods.go#L4005) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2069) |
| [SDL_GetAudioStreamQueued](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamQueued) | [:heavy_check_mark:](sdl/methods.go#L4016) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2078) |
| [SDL_FlushAudioStream](https://wiki.libsdl.org/SDL3/SDL_FlushAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4027) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2087) |
| [SDL_ClearAudioStream](https://wiki.libsdl.org/SDL3/SDL_ClearAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4037) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2096) |
| [SDL_PauseAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_PauseAudioStreamDevice) | [:heavy_check_mark:](sdl/methods.go#L4047) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2105) |
| [SDL_ResumeAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_ResumeAudioStreamDevice) | [:heavy_check_mark:](sdl/methods.go#L4057) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2114) |
| [SDL_AudioStreamDevicePaused](https://wiki.libsdl.org/SDL3/SDL_AudioStreamDevicePaused) | [:heavy_check_mark:](sdl/methods.go#L4067) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2123) |
| [SDL_LockAudioStream](https://wiki.libsdl.org/SDL3/SDL_LockAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4073) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2132) |
| [SDL_UnlockAudioStream](https://wiki.libsdl.org/SDL3/SDL_UnlockAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4083) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2141) |
| [SDL_SetAudioStreamGetCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGetCallback) | [:heavy_check_mark:](sdl/methods.go#L4093) | [:x:](sdl/sdl_functions_js.go#L12351) |
| [SDL_SetAudioStreamPutCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamPutCallback) | [:heavy_check_mark:](sdl/methods.go#L4099) | [:x:](sdl/sdl_functions_js.go#L12365) |
| [SDL_DestroyAudioStream](https://wiki.libsdl.org/SDL3/SDL_DestroyAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4105) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2190) |
| [SDL_OpenAudioDeviceStream](https://wiki.libsdl.org/SDL3/SDL_OpenAudioDeviceStream) | [:heavy_check_mark:](sdl/methods.go#L345) | [:x:](sdl/sdl_functions_js.go#L12379) |
| [SDL_SetAudioPostmixCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioPostmixCallback) | [:heavy_check_mark:](sdl/methods.go#L351) | [:x:](sdl/sdl_functions_js.go#L12397) |
| [SDL_LoadWAV_IO](https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L2240) |
| [SDL_LoadWAV](https://wiki.libsdl.org/SDL3/SDL_LoadWAV) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L2260) |
| [SDL_MixAudio](https://wiki.libsdl.org/SDL3/SDL_MixAudio) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L2278) |
| [SDL_ConvertAudioSamples](https://wiki.libsdl.org/SDL3/SDL_ConvertAudioSamples) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L2298) |
| [SDL_GetAudioFormatName](https://wiki.libsdl.org/SDL3/SDL_GetAudioFormatName) | [:heavy_check_mark:](sdl/methods.go#L1285) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2320) |
| [SDL_GetSilenceValueForFormat](https://wiki.libsdl.org/SDL3/SDL_GetSilenceValueForFormat) | [:heavy_check_mark:](sdl/methods.go#L1291) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2329) |
</details>
<details>
<summary><h3>Time</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetDateTimeLocalePreferences](https://wiki.libsdl.org/SDL3/SDL_GetDateTimeLocalePreferences) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11920) |
| [SDL_GetCurrentTime](https://wiki.libsdl.org/SDL3/SDL_GetCurrentTime) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11934) |
| [SDL_TimeToDateTime](https://wiki.libsdl.org/SDL3/SDL_TimeToDateTime) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11946) |
| [SDL_DateTimeToTime](https://wiki.libsdl.org/SDL3/SDL_DateTimeToTime) | [:heavy_check_mark:](sdl/methods.go#L5995) | [:x:](sdl/sdl_functions_js.go#L11962) |
| [SDL_TimeToWindows](https://wiki.libsdl.org/SDL3/SDL_TimeToWindows) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11976) |
| [SDL_TimeFromWindows](https://wiki.libsdl.org/SDL3/SDL_TimeFromWindows) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11991) |
| [SDL_GetDaysInMonth](https://wiki.libsdl.org/SDL3/SDL_GetDaysInMonth) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12002) |
| [SDL_GetDayOfYear](https://wiki.libsdl.org/SDL3/SDL_GetDayOfYear) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12013) |
| [SDL_GetDayOfWeek](https://wiki.libsdl.org/SDL3/SDL_GetDayOfWeek) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12026) |
</details>
<details open>
<summary><h3>Timer</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetTicks](https://wiki.libsdl.org/SDL3/SDL_GetTicks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12039) |
| [SDL_GetTicksNS](https://wiki.libsdl.org/SDL3/SDL_GetTicksNS) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12047) |
| [SDL_GetPerformanceCounter](https://wiki.libsdl.org/SDL3/SDL_GetPerformanceCounter) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12054) |
| [SDL_GetPerformanceFrequency](https://wiki.libsdl.org/SDL3/SDL_GetPerformanceFrequency) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12061) |
| [SDL_Delay](https://wiki.libsdl.org/SDL3/SDL_Delay) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12068) |
| [SDL_DelayNS](https://wiki.libsdl.org/SDL3/SDL_DelayNS) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12076) |
| [SDL_DelayPrecise](https://wiki.libsdl.org/SDL3/SDL_DelayPrecise) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12084) |
| [SDL_AddTimer](https://wiki.libsdl.org/SDL3/SDL_AddTimer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13202) |
| [SDL_AddTimerNS](https://wiki.libsdl.org/SDL3/SDL_AddTimerNS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13216) |
| [SDL_RemoveTimer](https://wiki.libsdl.org/SDL3/SDL_RemoveTimer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12126) |
</details>
<details open>
<summary><h3>Render</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumRenderDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumRenderDrivers) | [:heavy_check_mark:](sdl/functions.go#L420) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10199) |
| [SDL_GetRenderDriver](https://wiki.libsdl.org/SDL3/SDL_GetRenderDriver) | [:heavy_check_mark:](sdl/functions.go#L426) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10206) |
| [SDL_CreateWindowAndRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateWindowAndRenderer) | [:heavy_check_mark:](sdl/functions.go#L432) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10215) |
| [SDL_CreateRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateRenderer) | [:heavy_check_mark:](sdl/methods.go#L4897) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10243) |
| [SDL_CreateRendererWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateRendererWithProperties) | [:heavy_check_mark:](sdl/functions.go#L445) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10256) |
| [SDL_CreateGPURenderer](https://wiki.libsdl.org/SDL3/SDL_CreateGPURenderer) | [:heavy_check_mark:](sdl/methods.go#L2008) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13013) |
| [SDL_GetGPURendererDevice](https://wiki.libsdl.org/SDL3/SDL_GetGPURendererDevice) | [:heavy_check_mark:](sdl/methods.go#L2907) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13024) |
| [SDL_CreateSoftwareRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateSoftwareRenderer) | [:heavy_check_mark:](sdl/methods.go#L1951) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10265) |
| [SDL_GetRenderer](https://wiki.libsdl.org/SDL3/SDL_GetRenderer) | [:heavy_check_mark:](sdl/methods.go#L4908) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10274) |
| [SDL_GetRenderWindow](https://wiki.libsdl.org/SDL3/SDL_GetRenderWindow) | [:heavy_check_mark:](sdl/methods.go#L2896) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10283) |
| [SDL_GetRendererName](https://wiki.libsdl.org/SDL3/SDL_GetRendererName) | [:heavy_check_mark:](sdl/methods.go#L2918) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10292) |
| [SDL_GetRendererProperties](https://wiki.libsdl.org/SDL3/SDL_GetRendererProperties) | [:heavy_check_mark:](sdl/methods.go#L2929) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10301) |
| [SDL_GetRenderOutputSize](https://wiki.libsdl.org/SDL3/SDL_GetRenderOutputSize) | [:heavy_check_mark:](sdl/methods.go#L2935) | [:x:](sdl/sdl_functions_js.go#L10310) |
| [SDL_GetCurrentRenderOutputSize](https://wiki.libsdl.org/SDL3/SDL_GetCurrentRenderOutputSize) | [:heavy_check_mark:](sdl/methods.go#L2946) | [:x:](sdl/sdl_functions_js.go#L10326) |
| [SDL_CreateTexture](https://wiki.libsdl.org/SDL3/SDL_CreateTexture) | [:heavy_check_mark:](sdl/methods.go#L2957) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10342) |
| [SDL_CreateTextureFromSurface](https://wiki.libsdl.org/SDL3/SDL_CreateTextureFromSurface) | [:heavy_check_mark:](sdl/methods.go#L2968) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10365) |
| [SDL_CreateTextureWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateTextureWithProperties) | [:heavy_check_mark:](sdl/methods.go#L2979) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10385) |
| [SDL_GetTextureProperties](https://wiki.libsdl.org/SDL3/SDL_GetTextureProperties) | [:heavy_check_mark:](sdl/methods.go#L992) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10396) |
| [SDL_GetRendererFromTexture](https://wiki.libsdl.org/SDL3/SDL_GetRendererFromTexture) | [:heavy_check_mark:](sdl/methods.go#L1003) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10405) |
| [SDL_GetTextureSize](https://wiki.libsdl.org/SDL3/SDL_GetTextureSize) | [:heavy_check_mark:](sdl/methods.go#L1014) | [:x:](sdl/sdl_functions_js.go#L10414) |
| [SDL_SetTexturePalette](https://wiki.libsdl.org/SDL3/SDL_SetTexturePalette) | [:heavy_check_mark:](sdl/methods.go#L1025) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13033) |
| [SDL_GetTexturePalette](https://wiki.libsdl.org/SDL3/SDL_GetTexturePalette) | [:heavy_check_mark:](sdl/methods.go#L1035) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13044) |
| [SDL_SetTextureColorMod](https://wiki.libsdl.org/SDL3/SDL_SetTextureColorMod) | [:heavy_check_mark:](sdl/methods.go#L1041) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10430) |
| [SDL_SetTextureColorModFloat](https://wiki.libsdl.org/SDL3/SDL_SetTextureColorModFloat) | [:heavy_check_mark:](sdl/methods.go#L1051) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10445) |
| [SDL_GetTextureColorMod](https://wiki.libsdl.org/SDL3/SDL_GetTextureColorMod) | [:heavy_check_mark:](sdl/methods.go#L1061) | [:x:](sdl/sdl_functions_js.go#L10461) |
| [SDL_GetTextureColorModFloat](https://wiki.libsdl.org/SDL3/SDL_GetTextureColorModFloat) | [:heavy_check_mark:](sdl/methods.go#L1072) | [:x:](sdl/sdl_functions_js.go#L10479) |
| [SDL_SetTextureAlphaMod](https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1083) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10497) |
| [SDL_SetTextureAlphaModFloat](https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaModFloat) | [:heavy_check_mark:](sdl/methods.go#L1093) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10508) |
| [SDL_GetTextureAlphaMod](https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1103) | [:x:](sdl/sdl_functions_js.go#L10519) |
| [SDL_GetTextureAlphaModFloat](https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaModFloat) | [:heavy_check_mark:](sdl/methods.go#L1114) | [:x:](sdl/sdl_functions_js.go#L10533) |
| [SDL_SetTextureBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetTextureBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1125) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10547) |
| [SDL_GetTextureBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetTextureBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1135) | [:x:](sdl/sdl_functions_js.go#L10558) |
| [SDL_SetTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_SetTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L1146) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10572) |
| [SDL_GetTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_GetTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L1156) | [:x:](sdl/sdl_functions_js.go#L10583) |
| [SDL_UpdateTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateTexture) | [:heavy_check_mark:](sdl/methods.go#L1167) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10597) |
| [SDL_UpdateYUVTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateYUVTexture) | [:heavy_check_mark:](sdl/methods.go#L1179) | [:x:](sdl/sdl_functions_js.go#L10621) |
| [SDL_UpdateNVTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateNVTexture) | [:heavy_check_mark:](sdl/methods.go#L1193) | [:x:](sdl/sdl_functions_js.go#L10647) |
| [SDL_LockTexture](https://wiki.libsdl.org/SDL3/SDL_LockTexture) | [:heavy_check_mark:](sdl/methods.go#L1207) | [:x:](sdl/sdl_functions_js.go#L10669) |
| [SDL_LockTextureToSurface](https://wiki.libsdl.org/SDL3/SDL_LockTextureToSurface) | [:heavy_check_mark:](sdl/methods.go#L1220) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10687) |
| [SDL_UnlockTexture](https://wiki.libsdl.org/SDL3/SDL_UnlockTexture) | [:heavy_check_mark:](sdl/methods.go#L1230) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10709) |
| [SDL_SetRenderTarget](https://wiki.libsdl.org/SDL3/SDL_SetRenderTarget) | [:heavy_check_mark:](sdl/methods.go#L2990) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10720) |
| [SDL_GetRenderTarget](https://wiki.libsdl.org/SDL3/SDL_GetRenderTarget) | [:heavy_check_mark:](sdl/methods.go#L3000) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10731) |
| [SDL_SetRenderLogicalPresentation](https://wiki.libsdl.org/SDL3/SDL_SetRenderLogicalPresentation) | [:heavy_check_mark:](sdl/methods.go#L3006) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10740) |
| [SDL_GetRenderLogicalPresentation](https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentation) | [:heavy_check_mark:](sdl/methods.go#L3016) | [:x:](sdl/sdl_functions_js.go#L10755) |
| [SDL_GetRenderLogicalPresentationRect](https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentationRect) | [:heavy_check_mark:](sdl/methods.go#L3028) | [:x:](sdl/sdl_functions_js.go#L10773) |
| [SDL_RenderCoordinatesFromWindow](https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesFromWindow) | [:heavy_check_mark:](sdl/methods.go#L3040) | [:x:](sdl/sdl_functions_js.go#L10787) |
| [SDL_RenderCoordinatesToWindow](https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesToWindow) | [:heavy_check_mark:](sdl/methods.go#L3052) | [:x:](sdl/sdl_functions_js.go#L10807) |
| [SDL_ConvertEventToRenderCoordinates](https://wiki.libsdl.org/SDL3/SDL_ConvertEventToRenderCoordinates) | [:heavy_check_mark:](sdl/methods.go#L3064) | [:x:](sdl/sdl_functions_js.go#L10827) |
| [SDL_SetRenderViewport](https://wiki.libsdl.org/SDL3/SDL_SetRenderViewport) | [:heavy_check_mark:](sdl/methods.go#L3074) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10841) |
| [SDL_GetRenderViewport](https://wiki.libsdl.org/SDL3/SDL_GetRenderViewport) | [:heavy_check_mark:](sdl/methods.go#L3084) | [:x:](sdl/sdl_functions_js.go#L10858) |
| [SDL_RenderViewportSet](https://wiki.libsdl.org/SDL3/SDL_RenderViewportSet) | [:heavy_check_mark:](sdl/methods.go#L3096) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10872) |
| [SDL_GetRenderSafeArea](https://wiki.libsdl.org/SDL3/SDL_GetRenderSafeArea) | [:heavy_check_mark:](sdl/methods.go#L3102) | [:x:](sdl/sdl_functions_js.go#L10881) |
| [SDL_SetRenderClipRect](https://wiki.libsdl.org/SDL3/SDL_SetRenderClipRect) | [:heavy_check_mark:](sdl/methods.go#L3114) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10895) |
| [SDL_GetRenderClipRect](https://wiki.libsdl.org/SDL3/SDL_GetRenderClipRect) | [:heavy_check_mark:](sdl/methods.go#L3124) | [:x:](sdl/sdl_functions_js.go#L10912) |
| [SDL_RenderClipEnabled](https://wiki.libsdl.org/SDL3/SDL_RenderClipEnabled) | [:heavy_check_mark:](sdl/methods.go#L3136) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10926) |
| [SDL_SetRenderScale](https://wiki.libsdl.org/SDL3/SDL_SetRenderScale) | [:heavy_check_mark:](sdl/methods.go#L3146) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10935) |
| [SDL_GetRenderScale](https://wiki.libsdl.org/SDL3/SDL_GetRenderScale) | [:heavy_check_mark:](sdl/methods.go#L3156) | [:x:](sdl/sdl_functions_js.go#L10952) |
| [SDL_SetRenderDrawColor](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColor) | [:heavy_check_mark:](sdl/methods.go#L3168) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10968) |
| [SDL_SetRenderDrawColorFloat](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColorFloat) | [:heavy_check_mark:](sdl/methods.go#L3178) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10990) |
| [SDL_GetRenderDrawColor](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColor) | [:heavy_check_mark:](sdl/methods.go#L3188) | [:x:](sdl/sdl_functions_js.go#L11007) |
| [SDL_GetRenderDrawColorFloat](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColorFloat) | [:heavy_check_mark:](sdl/methods.go#L3200) | [:x:](sdl/sdl_functions_js.go#L11027) |
| [SDL_SetRenderColorScale](https://wiki.libsdl.org/SDL3/SDL_SetRenderColorScale) | [:heavy_check_mark:](sdl/methods.go#L3212) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11047) |
| [SDL_GetRenderColorScale](https://wiki.libsdl.org/SDL3/SDL_GetRenderColorScale) | [:heavy_check_mark:](sdl/methods.go#L3222) | [:x:](sdl/sdl_functions_js.go#L11058) |
| [SDL_SetRenderDrawBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawBlendMode) | [:heavy_check_mark:](sdl/methods.go#L3234) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11072) |
| [SDL_GetRenderDrawBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawBlendMode) | [:heavy_check_mark:](sdl/methods.go#L3244) | [:x:](sdl/sdl_functions_js.go#L11087) |
| [SDL_RenderClear](https://wiki.libsdl.org/SDL3/SDL_RenderClear) | [:heavy_check_mark:](sdl/methods.go#L3256) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11101) |
| [SDL_RenderPoint](https://wiki.libsdl.org/SDL3/SDL_RenderPoint) | [:heavy_check_mark:](sdl/methods.go#L3265) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11114) |
| [SDL_RenderPoints](https://wiki.libsdl.org/SDL3/SDL_RenderPoints) | [:heavy_check_mark:](sdl/methods.go#L3275) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11127) |
| [SDL_RenderLine](https://wiki.libsdl.org/SDL3/SDL_RenderLine) | [:heavy_check_mark:](sdl/methods.go#L3285) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11145) |
| [SDL_RenderLines](https://wiki.libsdl.org/SDL3/SDL_RenderLines) | [:heavy_check_mark:](sdl/methods.go#L3295) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11162) |
| [SDL_RenderRect](https://wiki.libsdl.org/SDL3/SDL_RenderRect) | [:heavy_check_mark:](sdl/methods.go#L3305) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11180) |
| [SDL_RenderRects](https://wiki.libsdl.org/SDL3/SDL_RenderRects) | [:heavy_check_mark:](sdl/methods.go#L3315) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11197) |
| [SDL_RenderFillRect](https://wiki.libsdl.org/SDL3/SDL_RenderFillRect) | [:heavy_check_mark:](sdl/methods.go#L3325) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11215) |
| [SDL_RenderFillRects](https://wiki.libsdl.org/SDL3/SDL_RenderFillRects) | [:heavy_check_mark:](sdl/methods.go#L3335) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11232) |
| [SDL_RenderTexture](https://wiki.libsdl.org/SDL3/SDL_RenderTexture) | [:heavy_check_mark:](sdl/methods.go#L3345) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11250) |
| [SDL_RenderTextureRotated](https://wiki.libsdl.org/SDL3/SDL_RenderTextureRotated) | [:heavy_check_mark:](sdl/methods.go#L3355) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11274) |
| [SDL_RenderTextureAffine](https://wiki.libsdl.org/SDL3/SDL_RenderTextureAffine) | [:heavy_check_mark:](sdl/methods.go#L3365) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11304) |
| [SDL_RenderTextureTiled](https://wiki.libsdl.org/SDL3/SDL_RenderTextureTiled) | [:heavy_check_mark:](sdl/methods.go#L3375) | [:x:](sdl/sdl_functions_js.go#L11332) |
| [SDL_RenderTexture9Grid](https://wiki.libsdl.org/SDL3/SDL_RenderTexture9Grid) | [:heavy_check_mark:](sdl/methods.go#L3385) | [:x:](sdl/sdl_functions_js.go#L11352) |
| [SDL_RenderTexture9GridTiled](https://wiki.libsdl.org/SDL3/SDL_RenderTexture9GridTiled) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13053) |
| [SDL_RenderGeometry](https://wiki.libsdl.org/SDL3/SDL_RenderGeometry) | [:heavy_check_mark:](sdl/methods.go#L3395) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11380) |
| [SDL_RenderGeometryRaw](https://wiki.libsdl.org/SDL3/SDL_RenderGeometryRaw) | [:heavy_check_mark:](sdl/methods.go#L3405) | [:x:](sdl/sdl_functions_js.go#L11407) |
| [SDL_SetRenderTextureAddressMode](https://wiki.libsdl.org/SDL3/SDL_SetRenderTextureAddressMode) | [:heavy_check_mark:](sdl/methods.go#L3426) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13083) |
| [SDL_GetRenderTextureAddressMode](https://wiki.libsdl.org/SDL3/SDL_GetRenderTextureAddressMode) | [:heavy_check_mark:](sdl/methods.go#L3436) | [:x:](sdl/sdl_functions_js.go#L13096) |
| [SDL_RenderReadPixels](https://wiki.libsdl.org/SDL3/SDL_RenderReadPixels) | [:heavy_check_mark:](sdl/methods.go#L3473) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11441) |
| [SDL_RenderPresent](https://wiki.libsdl.org/SDL3/SDL_RenderPresent) | [:heavy_check_mark:](sdl/methods.go#L3484) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11460) |
| [SDL_DestroyTexture](https://wiki.libsdl.org/SDL3/SDL_DestroyTexture) | [:heavy_check_mark:](sdl/methods.go#L1236) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11473) |
| [SDL_DestroyRenderer](https://wiki.libsdl.org/SDL3/SDL_DestroyRenderer) | [:heavy_check_mark:](sdl/methods.go#L3494) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11485) |
| [SDL_FlushRenderer](https://wiki.libsdl.org/SDL3/SDL_FlushRenderer) | [:heavy_check_mark:](sdl/methods.go#L3500) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11497) |
| [SDL_GetRenderMetalLayer](https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalLayer) | [:x:](sdl/methods.go#L3510) | [:x:](sdl/sdl_functions_js.go#L11506) |
| [SDL_GetRenderMetalCommandEncoder](https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalCommandEncoder) | [:x:](sdl/methods.go#L3517) | [:x:](sdl/sdl_functions_js.go#L11516) |
| [SDL_AddVulkanRenderSemaphores](https://wiki.libsdl.org/SDL3/SDL_AddVulkanRenderSemaphores) | [:heavy_check_mark:](sdl/methods.go#L3524) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11526) |
| [SDL_SetRenderVSync](https://wiki.libsdl.org/SDL3/SDL_SetRenderVSync) | [:heavy_check_mark:](sdl/methods.go#L3534) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11541) |
| [SDL_GetRenderVSync](https://wiki.libsdl.org/SDL3/SDL_GetRenderVSync) | [:heavy_check_mark:](sdl/methods.go#L3544) | [:x:](sdl/sdl_functions_js.go#L11552) |
| [SDL_RenderDebugText](https://wiki.libsdl.org/SDL3/SDL_RenderDebugText) | [:heavy_check_mark:](sdl/methods.go#L3566) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11566) |
| [SDL_RenderDebugTextFormat](https://wiki.libsdl.org/SDL3/SDL_RenderDebugTextFormat) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11586) |
| [SDL_SetDefaultTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_SetDefaultTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L3572) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13112) |
| [SDL_GetDefaultTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_GetDefaultTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L3582) | [:x:](sdl/sdl_functions_js.go#L13123) |
| [SDL_CreateGPURenderState](https://wiki.libsdl.org/SDL3/SDL_CreateGPURenderState) | [:heavy_check_mark:](sdl/methods.go#L3448) | [:x:](sdl/sdl_functions_js.go#L13555) |
| [SDL_SetGPURenderStateSamplerBindings](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderStateSamplerBindings) | [:question:]() | [:question:]() |
| [SDL_SetGPURenderStateStorageTextures](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderStateStorageTextures) | [:question:]() | [:question:]() |
| [SDL_SetGPURenderStateStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderStateStorageBuffers) | [:question:]() | [:question:]() |
| [SDL_SetGPURenderStateFragmentUniforms](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderStateFragmentUniforms) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13137) |
| [SDL_SetGPURenderState](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderState) | [:heavy_check_mark:](sdl/methods.go#L3461) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13153) |
| [SDL_DestroyGPURenderState](https://wiki.libsdl.org/SDL3/SDL_DestroyGPURenderState) | [:heavy_check_mark:](sdl/methods.go#L3596) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13164) |
| [SDL_GDKSuspendRenderer](https://wiki.libsdl.org/SDL3/SDL_GDKSuspendRenderer) | [:question:]() | [:question:]() |
| [SDL_GDKResumeRenderer](https://wiki.libsdl.org/SDL3/SDL_GDKResumeRenderer) | [:question:]() | [:question:]() |
</details>
<details>
<summary><h3>SharedObject</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_LoadObject](https://wiki.libsdl.org/SDL3/SDL_LoadObject) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9762) |
| [SDL_LoadFunction](https://wiki.libsdl.org/SDL3/SDL_LoadFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12966) |
| [SDL_UnloadObject](https://wiki.libsdl.org/SDL3/SDL_UnloadObject) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L9791) |
</details>
<details>
<summary><h3>Thread</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateThread](https://wiki.libsdl.org/SDL3/SDL_CreateThread) | [:question:]() | [:question:]() |
| [SDL_CreateThreadWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateThreadWithProperties) | [:question:]() | [:question:]() |
| [SDL_GetThreadName](https://wiki.libsdl.org/SDL3/SDL_GetThreadName) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L692) |
| [SDL_GetCurrentThreadID](https://wiki.libsdl.org/SDL3/SDL_GetCurrentThreadID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L701) |
| [SDL_GetThreadID](https://wiki.libsdl.org/SDL3/SDL_GetThreadID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L708) |
| [SDL_SetCurrentThreadPriority](https://wiki.libsdl.org/SDL3/SDL_SetCurrentThreadPriority) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L717) |
| [SDL_WaitThread](https://wiki.libsdl.org/SDL3/SDL_WaitThread) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L726) |
| [SDL_GetThreadState](https://wiki.libsdl.org/SDL3/SDL_GetThreadState) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L739) |
| [SDL_DetachThread](https://wiki.libsdl.org/SDL3/SDL_DetachThread) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L748) |
| [SDL_GetTLS](https://wiki.libsdl.org/SDL3/SDL_GetTLS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L756) |
| [SDL_SetTLS](https://wiki.libsdl.org/SDL3/SDL_SetTLS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12299) |
| [SDL_CleanupTLS](https://wiki.libsdl.org/SDL3/SDL_CleanupTLS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L788) |
</details>
<details>
<summary><h3>Mutex</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateMutex](https://wiki.libsdl.org/SDL3/SDL_CreateMutex) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L794) |
| [SDL_LockMutex](https://wiki.libsdl.org/SDL3/SDL_LockMutex) | [:heavy_check_mark:](sdl/methods.go#L5759) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L801) |
| [SDL_TryLockMutex](https://wiki.libsdl.org/SDL3/SDL_TryLockMutex) | [:heavy_check_mark:](sdl/methods.go#L5765) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L809) |
| [SDL_UnlockMutex](https://wiki.libsdl.org/SDL3/SDL_UnlockMutex) | [:heavy_check_mark:](sdl/methods.go#L5771) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L818) |
| [SDL_DestroyMutex](https://wiki.libsdl.org/SDL3/SDL_DestroyMutex) | [:heavy_check_mark:](sdl/methods.go#L5777) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L826) |
| [SDL_CreateRWLock](https://wiki.libsdl.org/SDL3/SDL_CreateRWLock) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L835) |
| [SDL_LockRWLockForReading](https://wiki.libsdl.org/SDL3/SDL_LockRWLockForReading) | [:heavy_check_mark:](sdl/methods.go#L1247) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L842) |
| [SDL_LockRWLockForWriting](https://wiki.libsdl.org/SDL3/SDL_LockRWLockForWriting) | [:heavy_check_mark:](sdl/methods.go#L1253) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L850) |
| [SDL_TryLockRWLockForReading](https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForReading) | [:heavy_check_mark:](sdl/methods.go#L1259) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L858) |
| [SDL_TryLockRWLockForWriting](https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForWriting) | [:heavy_check_mark:](sdl/methods.go#L1265) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L867) |
| [SDL_UnlockRWLock](https://wiki.libsdl.org/SDL3/SDL_UnlockRWLock) | [:heavy_check_mark:](sdl/methods.go#L1271) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L876) |
| [SDL_DestroyRWLock](https://wiki.libsdl.org/SDL3/SDL_DestroyRWLock) | [:heavy_check_mark:](sdl/methods.go#L1277) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L884) |
| [SDL_CreateSemaphore](https://wiki.libsdl.org/SDL3/SDL_CreateSemaphore) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L893) |
| [SDL_DestroySemaphore](https://wiki.libsdl.org/SDL3/SDL_DestroySemaphore) | [:heavy_check_mark:](sdl/methods.go#L5930) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L902) |
| [SDL_WaitSemaphore](https://wiki.libsdl.org/SDL3/SDL_WaitSemaphore) | [:heavy_check_mark:](sdl/methods.go#L5936) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L911) |
| [SDL_TryWaitSemaphore](https://wiki.libsdl.org/SDL3/SDL_TryWaitSemaphore) | [:heavy_check_mark:](sdl/methods.go#L5942) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L919) |
| [SDL_WaitSemaphoreTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitSemaphoreTimeout) | [:heavy_check_mark:](sdl/methods.go#L5948) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L928) |
| [SDL_SignalSemaphore](https://wiki.libsdl.org/SDL3/SDL_SignalSemaphore) | [:heavy_check_mark:](sdl/methods.go#L5954) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L939) |
| [SDL_GetSemaphoreValue](https://wiki.libsdl.org/SDL3/SDL_GetSemaphoreValue) | [:heavy_check_mark:](sdl/methods.go#L5960) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L947) |
| [SDL_CreateCondition](https://wiki.libsdl.org/SDL3/SDL_CreateCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L956) |
| [SDL_DestroyCondition](https://wiki.libsdl.org/SDL3/SDL_DestroyCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L963) |
| [SDL_SignalCondition](https://wiki.libsdl.org/SDL3/SDL_SignalCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L972) |
| [SDL_BroadcastCondition](https://wiki.libsdl.org/SDL3/SDL_BroadcastCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L980) |
| [SDL_WaitCondition](https://wiki.libsdl.org/SDL3/SDL_WaitCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L988) |
| [SDL_WaitConditionTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitConditionTimeout) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L998) |
| [SDL_ShouldInit](https://wiki.libsdl.org/SDL3/SDL_ShouldInit) | [:heavy_check_mark:](sdl/methods.go#L3641) | [:x:](sdl/sdl_functions_js.go#L1011) |
| [SDL_ShouldQuit](https://wiki.libsdl.org/SDL3/SDL_ShouldQuit) | [:heavy_check_mark:](sdl/methods.go#L3647) | [:x:](sdl/sdl_functions_js.go#L1023) |
| [SDL_SetInitialized](https://wiki.libsdl.org/SDL3/SDL_SetInitialized) | [:heavy_check_mark:](sdl/methods.go#L3653) | [:x:](sdl/sdl_functions_js.go#L1035) |
</details>
<details>
<summary><h3>Atomic</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_TryLockSpinlock](https://wiki.libsdl.org/SDL3/SDL_TryLockSpinlock) | [:heavy_check_mark:](sdl/methods.go#L572) | [:x:](sdl/sdl_functions_js.go#L179) |
| [SDL_LockSpinlock](https://wiki.libsdl.org/SDL3/SDL_LockSpinlock) | [:heavy_check_mark:](sdl/methods.go#L578) | [:x:](sdl/sdl_functions_js.go#L191) |
| [SDL_UnlockSpinlock](https://wiki.libsdl.org/SDL3/SDL_UnlockSpinlock) | [:heavy_check_mark:](sdl/methods.go#L584) | [:x:](sdl/sdl_functions_js.go#L202) |
| [SDL_MemoryBarrierReleaseFunction](https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierReleaseFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L213) |
| [SDL_MemoryBarrierAcquireFunction](https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierAcquireFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L219) |
| [SDL_CompareAndSwapAtomicInt](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L919) | [:x:](sdl/sdl_functions_js.go#L225) |
| [SDL_SetAtomicInt](https://wiki.libsdl.org/SDL3/SDL_SetAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L925) | [:x:](sdl/sdl_functions_js.go#L241) |
| [SDL_GetAtomicInt](https://wiki.libsdl.org/SDL3/SDL_GetAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L931) | [:x:](sdl/sdl_functions_js.go#L255) |
| [SDL_AddAtomicInt](https://wiki.libsdl.org/SDL3/SDL_AddAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L937) | [:x:](sdl/sdl_functions_js.go#L267) |
| [SDL_CompareAndSwapAtomicU32](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicU32) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L281) |
| [SDL_SetAtomicU32](https://wiki.libsdl.org/SDL3/SDL_SetAtomicU32) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L297) |
| [SDL_GetAtomicU32](https://wiki.libsdl.org/SDL3/SDL_GetAtomicU32) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L311) |
| [SDL_AddAtomicU32](https://wiki.libsdl.org/SDL3/SDL_AddAtomicU32) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12265) |
| [SDL_CompareAndSwapAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicPointer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L323) |
| [SDL_SetAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_SetAtomicPointer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L339) |
| [SDL_GetAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_GetAtomicPointer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L353) |
</details>
<details>
<summary><h3>Filesystem</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetBasePath](https://wiki.libsdl.org/SDL3/SDL_GetBasePath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7910) |
| [SDL_GetPrefPath](https://wiki.libsdl.org/SDL3/SDL_GetPrefPath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7917) |
| [SDL_GetUserFolder](https://wiki.libsdl.org/SDL3/SDL_GetUserFolder) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7931) |
| [SDL_CreateDirectory](https://wiki.libsdl.org/SDL3/SDL_CreateDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7940) |
| [SDL_EnumerateDirectory](https://wiki.libsdl.org/SDL3/SDL_EnumerateDirectory) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L12878) |
| [SDL_RemovePath](https://wiki.libsdl.org/SDL3/SDL_RemovePath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7968) |
| [SDL_RenamePath](https://wiki.libsdl.org/SDL3/SDL_RenamePath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7979) |
| [SDL_CopyFile](https://wiki.libsdl.org/SDL3/SDL_CopyFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7992) |
| [SDL_GetPathInfo](https://wiki.libsdl.org/SDL3/SDL_GetPathInfo) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8005) |
| [SDL_GlobDirectory](https://wiki.libsdl.org/SDL3/SDL_GlobDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8019) |
| [SDL_GetCurrentDirectory](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8037) |
</details>
<details open>
<summary><h3>IOStream</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_IOFromFile](https://wiki.libsdl.org/SDL3/SDL_IOFromFile) | [:heavy_check_mark:](sdl/functions.go#L370) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1048) |
| [SDL_IOFromMem](https://wiki.libsdl.org/SDL3/SDL_IOFromMem) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1061) |
| [SDL_IOFromConstMem](https://wiki.libsdl.org/SDL3/SDL_IOFromConstMem) | [:heavy_check_mark:](sdl/functions.go#L386) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1073) |
| [SDL_IOFromDynamicMem](https://wiki.libsdl.org/SDL3/SDL_IOFromDynamicMem) | [:heavy_check_mark:](sdl/functions.go#L402) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1093) |
| [SDL_OpenIO](https://wiki.libsdl.org/SDL3/SDL_OpenIO) | [:x:](sdl/methods.go#L23) | [:x:](sdl/sdl_functions_js.go#L1100) |
| [SDL_CloseIO](https://wiki.libsdl.org/SDL3/SDL_CloseIO) | [:heavy_check_mark:](sdl/methods.go#L4944) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1114) |
| [SDL_GetIOProperties](https://wiki.libsdl.org/SDL3/SDL_GetIOProperties) | [:heavy_check_mark:](sdl/methods.go#L4954) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1128) |
| [SDL_GetIOStatus](https://wiki.libsdl.org/SDL3/SDL_GetIOStatus) | [:heavy_check_mark:](sdl/methods.go#L4965) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1137) |
| [SDL_GetIOSize](https://wiki.libsdl.org/SDL3/SDL_GetIOSize) | [:heavy_check_mark:](sdl/methods.go#L4971) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1146) |
| [SDL_SeekIO](https://wiki.libsdl.org/SDL3/SDL_SeekIO) | [:heavy_check_mark:](sdl/methods.go#L4982) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1155) |
| [SDL_TellIO](https://wiki.libsdl.org/SDL3/SDL_TellIO) | [:heavy_check_mark:](sdl/methods.go#L4993) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1168) |
| [SDL_ReadIO](https://wiki.libsdl.org/SDL3/SDL_ReadIO) | [:heavy_check_mark:](sdl/methods.go#L4999) | [:x:](sdl/sdl_functions_js.go#L1177) |
| [SDL_WriteIO](https://wiki.libsdl.org/SDL3/SDL_WriteIO) | [:heavy_check_mark:](sdl/methods.go#L5010) | [:x:](sdl/sdl_functions_js.go#L1191) |
| [SDL_IOprintf](https://wiki.libsdl.org/SDL3/SDL_IOprintf) | [:heavy_check_mark:](sdl/methods.go#L5021) | [:x:](sdl/sdl_functions_js.go#L1205) |
| [SDL_IOvprintf](https://wiki.libsdl.org/SDL3/SDL_IOvprintf) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1219) |
| [SDL_FlushIO](https://wiki.libsdl.org/SDL3/SDL_FlushIO) | [:heavy_check_mark:](sdl/methods.go#L5032) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1235) |
| [SDL_LoadFile_IO](https://wiki.libsdl.org/SDL3/SDL_LoadFile_IO) | [:heavy_check_mark:](sdl/methods.go#L5042) | [:x:](sdl/sdl_functions_js.go#L1244) |
| [SDL_LoadFile](https://wiki.libsdl.org/SDL3/SDL_LoadFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1260) |
| [SDL_SaveFile_IO](https://wiki.libsdl.org/SDL3/SDL_SaveFile_IO) | [:heavy_check_mark:](sdl/methods.go#L5055) | [:x:](sdl/sdl_functions_js.go#L1274) |
| [SDL_SaveFile](https://wiki.libsdl.org/SDL3/SDL_SaveFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1290) |
| [SDL_ReadU8](https://wiki.libsdl.org/SDL3/SDL_ReadU8) | [:heavy_check_mark:](sdl/methods.go#L5066) | [:x:](sdl/sdl_functions_js.go#L1306) |
| [SDL_ReadS8](https://wiki.libsdl.org/SDL3/SDL_ReadS8) | [:heavy_check_mark:](sdl/methods.go#L5078) | [:x:](sdl/sdl_functions_js.go#L1320) |
| [SDL_ReadU16LE](https://wiki.libsdl.org/SDL3/SDL_ReadU16LE) | [:heavy_check_mark:](sdl/methods.go#L5090) | [:x:](sdl/sdl_functions_js.go#L1334) |
| [SDL_ReadS16LE](https://wiki.libsdl.org/SDL3/SDL_ReadS16LE) | [:heavy_check_mark:](sdl/methods.go#L5102) | [:x:](sdl/sdl_functions_js.go#L1348) |
| [SDL_ReadU16BE](https://wiki.libsdl.org/SDL3/SDL_ReadU16BE) | [:heavy_check_mark:](sdl/methods.go#L5114) | [:x:](sdl/sdl_functions_js.go#L1362) |
| [SDL_ReadS16BE](https://wiki.libsdl.org/SDL3/SDL_ReadS16BE) | [:heavy_check_mark:](sdl/methods.go#L5126) | [:x:](sdl/sdl_functions_js.go#L1376) |
| [SDL_ReadU32LE](https://wiki.libsdl.org/SDL3/SDL_ReadU32LE) | [:heavy_check_mark:](sdl/methods.go#L5138) | [:x:](sdl/sdl_functions_js.go#L1390) |
| [SDL_ReadS32LE](https://wiki.libsdl.org/SDL3/SDL_ReadS32LE) | [:heavy_check_mark:](sdl/methods.go#L5150) | [:x:](sdl/sdl_functions_js.go#L1404) |
| [SDL_ReadU32BE](https://wiki.libsdl.org/SDL3/SDL_ReadU32BE) | [:heavy_check_mark:](sdl/methods.go#L5162) | [:x:](sdl/sdl_functions_js.go#L1418) |
| [SDL_ReadS32BE](https://wiki.libsdl.org/SDL3/SDL_ReadS32BE) | [:heavy_check_mark:](sdl/methods.go#L5174) | [:x:](sdl/sdl_functions_js.go#L1432) |
| [SDL_ReadU64LE](https://wiki.libsdl.org/SDL3/SDL_ReadU64LE) | [:heavy_check_mark:](sdl/methods.go#L5186) | [:x:](sdl/sdl_functions_js.go#L1446) |
| [SDL_ReadS64LE](https://wiki.libsdl.org/SDL3/SDL_ReadS64LE) | [:heavy_check_mark:](sdl/methods.go#L5198) | [:x:](sdl/sdl_functions_js.go#L1460) |
| [SDL_ReadU64BE](https://wiki.libsdl.org/SDL3/SDL_ReadU64BE) | [:heavy_check_mark:](sdl/methods.go#L5210) | [:x:](sdl/sdl_functions_js.go#L1474) |
| [SDL_ReadS64BE](https://wiki.libsdl.org/SDL3/SDL_ReadS64BE) | [:heavy_check_mark:](sdl/methods.go#L5222) | [:x:](sdl/sdl_functions_js.go#L1488) |
| [SDL_WriteU8](https://wiki.libsdl.org/SDL3/SDL_WriteU8) | [:heavy_check_mark:](sdl/methods.go#L5234) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1502) |
| [SDL_WriteS8](https://wiki.libsdl.org/SDL3/SDL_WriteS8) | [:heavy_check_mark:](sdl/methods.go#L5244) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1513) |
| [SDL_WriteU16LE](https://wiki.libsdl.org/SDL3/SDL_WriteU16LE) | [:heavy_check_mark:](sdl/methods.go#L5254) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1524) |
| [SDL_WriteS16LE](https://wiki.libsdl.org/SDL3/SDL_WriteS16LE) | [:heavy_check_mark:](sdl/methods.go#L5264) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1535) |
| [SDL_WriteU16BE](https://wiki.libsdl.org/SDL3/SDL_WriteU16BE) | [:heavy_check_mark:](sdl/methods.go#L5274) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1546) |
| [SDL_WriteS16BE](https://wiki.libsdl.org/SDL3/SDL_WriteS16BE) | [:heavy_check_mark:](sdl/methods.go#L5284) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1557) |
| [SDL_WriteU32LE](https://wiki.libsdl.org/SDL3/SDL_WriteU32LE) | [:heavy_check_mark:](sdl/methods.go#L5294) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1568) |
| [SDL_WriteS32LE](https://wiki.libsdl.org/SDL3/SDL_WriteS32LE) | [:heavy_check_mark:](sdl/methods.go#L5304) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1579) |
| [SDL_WriteU32BE](https://wiki.libsdl.org/SDL3/SDL_WriteU32BE) | [:heavy_check_mark:](sdl/methods.go#L5314) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1590) |
| [SDL_WriteS32BE](https://wiki.libsdl.org/SDL3/SDL_WriteS32BE) | [:heavy_check_mark:](sdl/methods.go#L5324) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1601) |
| [SDL_WriteU64LE](https://wiki.libsdl.org/SDL3/SDL_WriteU64LE) | [:heavy_check_mark:](sdl/methods.go#L5334) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1612) |
| [SDL_WriteS64LE](https://wiki.libsdl.org/SDL3/SDL_WriteS64LE) | [:heavy_check_mark:](sdl/methods.go#L5344) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1623) |
| [SDL_WriteU64BE](https://wiki.libsdl.org/SDL3/SDL_WriteU64BE) | [:heavy_check_mark:](sdl/methods.go#L5354) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1634) |
| [SDL_WriteS64BE](https://wiki.libsdl.org/SDL3/SDL_WriteS64BE) | [:heavy_check_mark:](sdl/methods.go#L5364) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1645) |
</details>
<details open>
<summary><h3>AsyncIO</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AsyncIOFromFile](https://wiki.libsdl.org/SDL3/SDL_AsyncIOFromFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L30) |
| [SDL_GetAsyncIOSize](https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOSize) | [:heavy_check_mark:](sdl/methods.go#L3604) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L43) |
| [SDL_ReadAsyncIO](https://wiki.libsdl.org/SDL3/SDL_ReadAsyncIO) | [:x:](sdl/methods.go#L3615) | [:x:](sdl/sdl_functions_js.go#L52) |
| [SDL_WriteAsyncIO](https://wiki.libsdl.org/SDL3/SDL_WriteAsyncIO) | [:x:](sdl/methods.go#L3622) | [:x:](sdl/sdl_functions_js.go#L72) |
| [SDL_CloseAsyncIO](https://wiki.libsdl.org/SDL3/SDL_CloseAsyncIO) | [:heavy_check_mark:](sdl/methods.go#L3629) | [:x:](sdl/sdl_functions_js.go#L92) |
| [SDL_CreateAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_CreateAsyncIOQueue) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L109) |
| [SDL_DestroyAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_DestroyAsyncIOQueue) | [:heavy_check_mark:](sdl/methods.go#L1410) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L116) |
| [SDL_GetAsyncIOResult](https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOResult) | [:heavy_check_mark:](sdl/methods.go#L1416) | [:x:](sdl/sdl_functions_js.go#L125) |
| [SDL_WaitAsyncIOResult](https://wiki.libsdl.org/SDL3/SDL_WaitAsyncIOResult) | [:heavy_check_mark:](sdl/methods.go#L1426) | [:x:](sdl/sdl_functions_js.go#L139) |
| [SDL_SignalAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_SignalAsyncIOQueue) | [:heavy_check_mark:](sdl/methods.go#L1436) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L155) |
| [SDL_LoadFileAsync](https://wiki.libsdl.org/SDL3/SDL_LoadFileAsync) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L163) |
</details>
<details open>
<summary><h3>Storage</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_OpenTitleStorage](https://wiki.libsdl.org/SDL3/SDL_OpenTitleStorage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11603) |
| [SDL_OpenUserStorage](https://wiki.libsdl.org/SDL3/SDL_OpenUserStorage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11616) |
| [SDL_OpenFileStorage](https://wiki.libsdl.org/SDL3/SDL_OpenFileStorage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11631) |
| [SDL_OpenStorage](https://wiki.libsdl.org/SDL3/SDL_OpenStorage) | [:x:](sdl/methods.go#L5750) | [:x:](sdl/sdl_functions_js.go#L11642) |
| [SDL_CloseStorage](https://wiki.libsdl.org/SDL3/SDL_CloseStorage) | [:heavy_check_mark:](sdl/methods.go#L78) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11656) |
| [SDL_StorageReady](https://wiki.libsdl.org/SDL3/SDL_StorageReady) | [:heavy_check_mark:](sdl/methods.go#L88) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11666) |
| [SDL_GetStorageFileSize](https://wiki.libsdl.org/SDL3/SDL_GetStorageFileSize) | [:heavy_check_mark:](sdl/methods.go#L94) | [:x:](sdl/sdl_functions_js.go#L11675) |
| [SDL_ReadStorageFile](https://wiki.libsdl.org/SDL3/SDL_ReadStorageFile) | [:heavy_check_mark:](sdl/methods.go#L106) | [:x:](sdl/sdl_functions_js.go#L11691) |
| [SDL_WriteStorageFile](https://wiki.libsdl.org/SDL3/SDL_WriteStorageFile) | [:heavy_check_mark:](sdl/methods.go#L118) | [:x:](sdl/sdl_functions_js.go#L11709) |
| [SDL_CreateStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_CreateStorageDirectory) | [:heavy_check_mark:](sdl/methods.go#L128) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11727) |
| [SDL_EnumerateStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_EnumerateStorageDirectory) | [:heavy_check_mark:](sdl/methods.go#L138) | [:x:](sdl/sdl_functions_js.go#L13173) |
| [SDL_RemoveStoragePath](https://wiki.libsdl.org/SDL3/SDL_RemoveStoragePath) | [:heavy_check_mark:](sdl/methods.go#L148) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11762) |
| [SDL_RenameStoragePath](https://wiki.libsdl.org/SDL3/SDL_RenameStoragePath) | [:heavy_check_mark:](sdl/methods.go#L158) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11775) |
| [SDL_CopyStorageFile](https://wiki.libsdl.org/SDL3/SDL_CopyStorageFile) | [:heavy_check_mark:](sdl/methods.go#L168) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11790) |
| [SDL_GetStoragePathInfo](https://wiki.libsdl.org/SDL3/SDL_GetStoragePathInfo) | [:heavy_check_mark:](sdl/methods.go#L178) | [:x:](sdl/sdl_functions_js.go#L11805) |
| [SDL_GetStorageSpaceRemaining](https://wiki.libsdl.org/SDL3/SDL_GetStorageSpaceRemaining) | [:heavy_check_mark:](sdl/methods.go#L190) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11821) |
| [SDL_GlobStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_GlobStorageDirectory) | [:heavy_check_mark:](sdl/methods.go#L196) | [:x:](sdl/sdl_functions_js.go#L11830) |
</details>
<details open>
<summary><h3>Pixels</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPixelFormatName](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatName) | [:heavy_check_mark:](sdl/methods.go#L5968) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2357) |
| [SDL_GetMasksForPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetMasksForPixelFormat) | [:x:](sdl/methods.go#L5974) | [:x:](sdl/sdl_functions_js.go#L2366) |
| [SDL_GetPixelFormatForMasks](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatForMasks) | [:heavy_check_mark:](sdl/functions.go#L458) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2388) |
| [SDL_GetPixelFormatDetails](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatDetails) | [:heavy_check_mark:](sdl/methods.go#L5982) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2401) |
| [SDL_CreatePalette](https://wiki.libsdl.org/SDL3/SDL_CreatePalette) | [:heavy_check_mark:](sdl/functions.go#L464) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2413) |
| [SDL_SetPaletteColors](https://wiki.libsdl.org/SDL3/SDL_SetPaletteColors) | [:heavy_check_mark:](sdl/methods.go#L5376) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2422) |
| [SDL_DestroyPalette](https://wiki.libsdl.org/SDL3/SDL_DestroyPalette) | [:heavy_check_mark:](sdl/methods.go#L5386) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2442) |
| [SDL_MapRGB](https://wiki.libsdl.org/SDL3/SDL_MapRGB) | [:heavy_check_mark:](sdl/methods.go#L1444) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2454) |
| [SDL_MapRGBA](https://wiki.libsdl.org/SDL3/SDL_MapRGBA) | [:heavy_check_mark:](sdl/methods.go#L1450) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2481) |
| [SDL_GetRGB](https://wiki.libsdl.org/SDL3/SDL_GetRGB) | [:heavy_check_mark:](sdl/functions.go#L487) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2502) |
| [SDL_GetRGBA](https://wiki.libsdl.org/SDL3/SDL_GetRGBA) | [:heavy_check_mark:](sdl/functions.go#L495) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2533) |
</details>
<details open>
<summary><h3>Surface</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateSurface](https://wiki.libsdl.org/SDL3/SDL_CreateSurface) | [:heavy_check_mark:](sdl/functions.go#L505) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2748) |
| [SDL_CreateSurfaceFrom](https://wiki.libsdl.org/SDL3/SDL_CreateSurfaceFrom) | [:heavy_check_mark:](sdl/functions.go#L516) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2748) |
| [SDL_DestroySurface](https://wiki.libsdl.org/SDL3/SDL_DestroySurface) | [:heavy_check_mark:](sdl/methods.go#L1458) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2781) |
| [SDL_GetSurfaceProperties](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceProperties) | [:heavy_check_mark:](sdl/methods.go#L1467) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2793) |
| [SDL_SetSurfaceColorspace](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorspace) | [:heavy_check_mark:](sdl/methods.go#L1478) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2806) |
| [SDL_GetSurfaceColorspace](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorspace) | [:heavy_check_mark:](sdl/methods.go#L1488) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2817) |
| [SDL_CreateSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_CreateSurfacePalette) | [:heavy_check_mark:](sdl/methods.go#L1494) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2826) |
| [SDL_SetSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_SetSurfacePalette) | [:heavy_check_mark:](sdl/methods.go#L1505) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2839) |
| [SDL_GetSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_GetSurfacePalette) | [:heavy_check_mark:](sdl/methods.go#L1515) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2850) |
| [SDL_AddSurfaceAlternateImage](https://wiki.libsdl.org/SDL3/SDL_AddSurfaceAlternateImage) | [:heavy_check_mark:](sdl/methods.go#L1521) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2863) |
| [SDL_SurfaceHasAlternateImages](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasAlternateImages) | [:heavy_check_mark:](sdl/methods.go#L1531) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2874) |
| [SDL_GetSurfaceImages](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceImages) | [:heavy_check_mark:](sdl/methods.go#L1537) | [:x:](sdl/sdl_functions_js.go#L2887) |
| [SDL_RemoveSurfaceAlternateImages](https://wiki.libsdl.org/SDL3/SDL_RemoveSurfaceAlternateImages) | [:heavy_check_mark:](sdl/methods.go#L1551) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2901) |
| [SDL_LockSurface](https://wiki.libsdl.org/SDL3/SDL_LockSurface) | [:heavy_check_mark:](sdl/methods.go#L1557) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2912) |
| [SDL_UnlockSurface](https://wiki.libsdl.org/SDL3/SDL_UnlockSurface) | [:heavy_check_mark:](sdl/methods.go#L1567) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2925) |
| [SDL_LoadSurface_IO](https://wiki.libsdl.org/SDL3/SDL_LoadSurface_IO) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12411) |
| [SDL_LoadSurface](https://wiki.libsdl.org/SDL3/SDL_LoadSurface) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12422) |
| [SDL_LoadBMP_IO](https://wiki.libsdl.org/SDL3/SDL_LoadBMP_IO) | [:heavy_check_mark:](sdl/functions.go#L528) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2936) |
| [SDL_LoadBMP](https://wiki.libsdl.org/SDL3/SDL_LoadBMP) | [:heavy_check_mark:](sdl/functions.go#L539) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2956) |
| [SDL_SaveBMP_IO](https://wiki.libsdl.org/SDL3/SDL_SaveBMP_IO) | [:heavy_check_mark:](sdl/methods.go#L1573) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2968) |
| [SDL_SaveBMP](https://wiki.libsdl.org/SDL3/SDL_SaveBMP) | [:heavy_check_mark:](sdl/methods.go#L1583) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2981) |
| [SDL_LoadPNG_IO](https://wiki.libsdl.org/SDL3/SDL_LoadPNG_IO) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12433) |
| [SDL_LoadPNG](https://wiki.libsdl.org/SDL3/SDL_LoadPNG) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12444) |
| [SDL_SavePNG_IO](https://wiki.libsdl.org/SDL3/SDL_SavePNG_IO) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12455) |
| [SDL_SavePNG](https://wiki.libsdl.org/SDL3/SDL_SavePNG) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12468) |
| [SDL_LoadJPG_IO](https://wiki.libsdl.org/SDL3/SDL_LoadJPG_IO) | [:question:]() | [:question:]() |
| [SDL_LoadJPG](https://wiki.libsdl.org/SDL3/SDL_LoadJPG) | [:question:]() | [:question:]() |
| [SDL_SetSurfaceRLE](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceRLE) | [:heavy_check_mark:](sdl/methods.go#L1593) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2998) |
| [SDL_SurfaceHasRLE](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasRLE) | [:heavy_check_mark:](sdl/methods.go#L1603) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3009) |
| [SDL_SetSurfaceColorKey](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorKey) | [:heavy_check_mark:](sdl/methods.go#L1609) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3022) |
| [SDL_SurfaceHasColorKey](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasColorKey) | [:heavy_check_mark:](sdl/methods.go#L1619) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3038) |
| [SDL_GetSurfaceColorKey](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorKey) | [:heavy_check_mark:](sdl/methods.go#L1625) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3047) |
| [SDL_SetSurfaceColorMod](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorMod) | [:heavy_check_mark:](sdl/methods.go#L1636) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3065) |
| [SDL_GetSurfaceColorMod](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorMod) | [:heavy_check_mark:](sdl/methods.go#L1646) | [:x:](sdl/sdl_functions_js.go#L3081) |
| [SDL_SetSurfaceAlphaMod](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1658) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3099) |
| [SDL_GetSurfaceAlphaMod](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1668) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3113) |
| [SDL_SetSurfaceBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1680) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3131) |
| [SDL_GetSurfaceBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1690) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3145) |
| [SDL_SetSurfaceClipRect](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceClipRect) | [:heavy_check_mark:](sdl/methods.go#L1702) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3163) |
| [SDL_GetSurfaceClipRect](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceClipRect) | [:heavy_check_mark:](sdl/methods.go#L1708) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3180) |
| [SDL_FlipSurface](https://wiki.libsdl.org/SDL3/SDL_FlipSurface) | [:heavy_check_mark:](sdl/methods.go#L1720) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3198) |
| [SDL_RotateSurface](https://wiki.libsdl.org/SDL3/SDL_RotateSurface) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12481) |
| [SDL_DuplicateSurface](https://wiki.libsdl.org/SDL3/SDL_DuplicateSurface) | [:heavy_check_mark:](sdl/methods.go#L1730) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3209) |
| [SDL_ScaleSurface](https://wiki.libsdl.org/SDL3/SDL_ScaleSurface) | [:heavy_check_mark:](sdl/methods.go#L1741) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3222) |
| [SDL_ConvertSurface](https://wiki.libsdl.org/SDL3/SDL_ConvertSurface) | [:heavy_check_mark:](sdl/methods.go#L1752) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3238) |
| [SDL_ConvertSurfaceAndColorspace](https://wiki.libsdl.org/SDL3/SDL_ConvertSurfaceAndColorspace) | [:heavy_check_mark:](sdl/methods.go#L1763) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3255) |
| [SDL_ConvertPixels](https://wiki.libsdl.org/SDL3/SDL_ConvertPixels) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3272) |
| [SDL_ConvertPixelsAndColorspace](https://wiki.libsdl.org/SDL3/SDL_ConvertPixelsAndColorspace) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3296) |
| [SDL_PremultiplyAlpha](https://wiki.libsdl.org/SDL3/SDL_PremultiplyAlpha) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3328) |
| [SDL_PremultiplySurfaceAlpha](https://wiki.libsdl.org/SDL3/SDL_PremultiplySurfaceAlpha) | [:heavy_check_mark:](sdl/methods.go#L1774) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3354) |
| [SDL_ClearSurface](https://wiki.libsdl.org/SDL3/SDL_ClearSurface) | [:heavy_check_mark:](sdl/methods.go#L1784) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3371) |
| [SDL_FillSurfaceRect](https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRect) | [:heavy_check_mark:](sdl/methods.go#L1794) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3388) |
| [SDL_FillSurfaceRects](https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRects) | [:heavy_check_mark:](sdl/methods.go#L1804) | [:x:](sdl/sdl_functions_js.go#L3407) |
| [SDL_BlitSurface](https://wiki.libsdl.org/SDL3/SDL_BlitSurface) | [:heavy_check_mark:](sdl/methods.go#L1814) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3425) |
| [SDL_BlitSurfaceUnchecked](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUnchecked) | [:heavy_check_mark:](sdl/methods.go#L1824) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3449) |
| [SDL_BlitSurfaceScaled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceScaled) | [:heavy_check_mark:](sdl/methods.go#L1834) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3473) |
| [SDL_BlitSurfaceUncheckedScaled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUncheckedScaled) | [:heavy_check_mark:](sdl/methods.go#L1844) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3498) |
| [SDL_StretchSurface](https://wiki.libsdl.org/SDL3/SDL_StretchSurface) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12492) |
| [SDL_BlitSurfaceTiled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiled) | [:heavy_check_mark:](sdl/methods.go#L1854) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3523) |
| [SDL_BlitSurfaceTiledWithScale](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiledWithScale) | [:heavy_check_mark:](sdl/methods.go#L1864) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3547) |
| [SDL_BlitSurface9Grid](https://wiki.libsdl.org/SDL3/SDL_BlitSurface9Grid) | [:heavy_check_mark:](sdl/methods.go#L1874) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3575) |
| [SDL_MapSurfaceRGB](https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGB) | [:heavy_check_mark:](sdl/methods.go#L1884) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3611) |
| [SDL_MapSurfaceRGBA](https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGBA) | [:heavy_check_mark:](sdl/methods.go#L1890) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3626) |
| [SDL_ReadSurfacePixel](https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixel) | [:heavy_check_mark:](sdl/methods.go#L1896) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3643) |
| [SDL_ReadSurfacePixelFloat](https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixelFloat) | [:heavy_check_mark:](sdl/methods.go#L1908) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3674) |
| [SDL_WriteSurfacePixel](https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixel) | [:heavy_check_mark:](sdl/methods.go#L1920) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3703) |
| [SDL_WriteSurfacePixelFloat](https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixelFloat) | [:heavy_check_mark:](sdl/methods.go#L1930) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3724) |
</details>
<details open>
<summary><h3>BlendMode</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ComposeCustomBlendMode](https://wiki.libsdl.org/SDL3/SDL_ComposeCustomBlendMode) | [:heavy_check_mark:](sdl/functions.go#L556) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2338) |
</details>
<details open>
<summary><h3>Rect</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_RectToFRect](https://wiki.libsdl.org/SDL3/SDL_RectToFRect) | [:question:]() | [:question:]() |
| [SDL_PointInRect](https://wiki.libsdl.org/SDL3/SDL_PointInRect) | [:question:]() | [:question:]() |
| [SDL_RectEmpty](https://wiki.libsdl.org/SDL3/SDL_RectEmpty) | [:question:]() | [:question:]() |
| [SDL_RectsEqual](https://wiki.libsdl.org/SDL3/SDL_RectsEqual) | [:question:]() | [:question:]() |
| [SDL_HasRectIntersection](https://wiki.libsdl.org/SDL3/SDL_HasRectIntersection) | [:heavy_check_mark:](sdl/methods.go#L671) | [:x:](sdl/sdl_functions_js.go#L2569) |
| [SDL_GetRectIntersection](https://wiki.libsdl.org/SDL3/SDL_GetRectIntersection) | [:heavy_check_mark:](sdl/methods.go#L677) | [:x:](sdl/sdl_functions_js.go#L2583) |
| [SDL_GetRectUnion](https://wiki.libsdl.org/SDL3/SDL_GetRectUnion) | [:heavy_check_mark:](sdl/methods.go#L689) | [:x:](sdl/sdl_functions_js.go#L2599) |
| [SDL_GetRectEnclosingPoints](https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPoints) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L2615) |
| [SDL_GetRectAndLineIntersection](https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersection) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L2633) |
| [SDL_PointInRectFloat](https://wiki.libsdl.org/SDL3/SDL_PointInRectFloat) | [:question:]() | [:question:]() |
| [SDL_RectEmptyFloat](https://wiki.libsdl.org/SDL3/SDL_RectEmptyFloat) | [:question:]() | [:question:]() |
| [SDL_RectsEqualEpsilon](https://wiki.libsdl.org/SDL3/SDL_RectsEqualEpsilon) | [:question:]() | [:question:]() |
| [SDL_RectsEqualFloat](https://wiki.libsdl.org/SDL3/SDL_RectsEqualFloat) | [:question:]() | [:question:]() |
| [SDL_HasRectIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_HasRectIntersectionFloat) | [:heavy_check_mark:](sdl/methods.go#L4113) | [:x:](sdl/sdl_functions_js.go#L2653) |
| [SDL_GetRectIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectIntersectionFloat) | [:x:](sdl/methods.go#L4119) | [:x:](sdl/sdl_functions_js.go#L2667) |
| [SDL_GetRectUnionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectUnionFloat) | [:x:](sdl/methods.go#L4126) | [:x:](sdl/sdl_functions_js.go#L2683) |
| [SDL_GetRectEnclosingPointsFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPointsFloat) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L2699) |
| [SDL_GetRectAndLineIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersectionFloat) | [:x:](sdl/methods.go#L4133) | [:x:](sdl/sdl_functions_js.go#L2717) |
</details>
<details open>
<summary><h3>Camera</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumCameraDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumCameraDrivers) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3743) |
| [SDL_GetCameraDriver](https://wiki.libsdl.org/SDL3/SDL_GetCameraDriver) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3751) |
| [SDL_GetCurrentCameraDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentCameraDriver) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3760) |
| [SDL_GetCameras](https://wiki.libsdl.org/SDL3/SDL_GetCameras) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L3768) |
| [SDL_GetCameraSupportedFormats](https://wiki.libsdl.org/SDL3/SDL_GetCameraSupportedFormats) | [:heavy_check_mark:](sdl/methods.go#L3661) | [:x:](sdl/sdl_functions_js.go#L3780) |
| [SDL_GetCameraName](https://wiki.libsdl.org/SDL3/SDL_GetCameraName) | [:x:](sdl/methods.go#L3676) | [:x:](sdl/sdl_functions_js.go#L3794) |
| [SDL_GetCameraPosition](https://wiki.libsdl.org/SDL3/SDL_GetCameraPosition) | [:heavy_check_mark:](sdl/methods.go#L3683) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3803) |
| [SDL_OpenCamera](https://wiki.libsdl.org/SDL3/SDL_OpenCamera) | [:heavy_check_mark:](sdl/methods.go#L3689) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3812) |
| [SDL_GetCameraPermissionState](https://wiki.libsdl.org/SDL3/SDL_GetCameraPermissionState) | [:heavy_check_mark:](sdl/methods.go#L359) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3826) |
| [SDL_GetCameraID](https://wiki.libsdl.org/SDL3/SDL_GetCameraID) | [:heavy_check_mark:](sdl/methods.go#L365) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3839) |
| [SDL_GetCameraProperties](https://wiki.libsdl.org/SDL3/SDL_GetCameraProperties) | [:heavy_check_mark:](sdl/methods.go#L376) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3852) |
| [SDL_GetCameraFormat](https://wiki.libsdl.org/SDL3/SDL_GetCameraFormat) | [:heavy_check_mark:](sdl/methods.go#L387) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3865) |
| [SDL_AcquireCameraFrame](https://wiki.libsdl.org/SDL3/SDL_AcquireCameraFrame) | [:heavy_check_mark:](sdl/methods.go#L399) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3883) |
| [SDL_ReleaseCameraFrame](https://wiki.libsdl.org/SDL3/SDL_ReleaseCameraFrame) | [:heavy_check_mark:](sdl/methods.go#L412) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3903) |
| [SDL_CloseCamera](https://wiki.libsdl.org/SDL3/SDL_CloseCamera) | [:heavy_check_mark:](sdl/methods.go#L418) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3919) |
</details>
<details open>
<summary><h3>MessageBox</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ShowMessageBox](https://wiki.libsdl.org/SDL3/SDL_ShowMessageBox) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10019) |
| [SDL_ShowSimpleMessageBox](https://wiki.libsdl.org/SDL3/SDL_ShowSimpleMessageBox) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10034) |
</details>
<details open>
<summary><h3>Clipboard</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetClipboardText](https://wiki.libsdl.org/SDL3/SDL_SetClipboardText) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3931) |
| [SDL_GetClipboardText](https://wiki.libsdl.org/SDL3/SDL_GetClipboardText) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L3943) |
| [SDL_HasClipboardText](https://wiki.libsdl.org/SDL3/SDL_HasClipboardText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3951) |
| [SDL_SetPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_SetPrimarySelectionText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3958) |
| [SDL_GetPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_GetPrimarySelectionText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3969) |
| [SDL_HasPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_HasPrimarySelectionText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3977) |
| [SDL_SetClipboardData](https://wiki.libsdl.org/SDL3/SDL_SetClipboardData) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L13488) |
| [SDL_ClearClipboardData](https://wiki.libsdl.org/SDL3/SDL_ClearClipboardData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4008) |
| [SDL_GetClipboardData](https://wiki.libsdl.org/SDL3/SDL_GetClipboardData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4015) |
| [SDL_HasClipboardData](https://wiki.libsdl.org/SDL3/SDL_HasClipboardData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4029) |
| [SDL_GetClipboardMimeTypes](https://wiki.libsdl.org/SDL3/SDL_GetClipboardMimeTypes) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4040) |
</details>
<details open>
<summary><h3>Dialog</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ShowOpenFileDialog](https://wiki.libsdl.org/SDL3/SDL_ShowOpenFileDialog) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L13509) |
| [SDL_ShowSaveFileDialog](https://wiki.libsdl.org/SDL3/SDL_ShowSaveFileDialog) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L13533) |
| [SDL_ShowOpenFolderDialog](https://wiki.libsdl.org/SDL3/SDL_ShowOpenFolderDialog) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L12651) |
| [SDL_ShowFileDialogWithProperties](https://wiki.libsdl.org/SDL3/SDL_ShowFileDialogWithProperties) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L12670) |
</details>
<details open>
<summary><h3>Tray</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateTray](https://wiki.libsdl.org/SDL3/SDL_CreateTray) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L13230) |
| [SDL_CreateTrayWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateTrayWithProperties) | [:question:]() | [:question:]() |
| [SDL_SetTrayIcon](https://wiki.libsdl.org/SDL3/SDL_SetTrayIcon) | [:heavy_check_mark:](sdl/methods.go#L426) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13244) |
| [SDL_SetTrayTooltip](https://wiki.libsdl.org/SDL3/SDL_SetTrayTooltip) | [:heavy_check_mark:](sdl/methods.go#L432) | [:x:](sdl/sdl_functions_js.go#L13254) |
| [SDL_CreateTrayMenu](https://wiki.libsdl.org/SDL3/SDL_CreateTrayMenu) | [:heavy_check_mark:](sdl/methods.go#L444) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13267) |
| [SDL_CreateTraySubmenu](https://wiki.libsdl.org/SDL3/SDL_CreateTraySubmenu) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13276) |
| [SDL_GetTrayMenu](https://wiki.libsdl.org/SDL3/SDL_GetTrayMenu) | [:heavy_check_mark:](sdl/methods.go#L438) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13285) |
| [SDL_GetTraySubmenu](https://wiki.libsdl.org/SDL3/SDL_GetTraySubmenu) | [:heavy_check_mark:](sdl/methods.go#L488) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13294) |
| [SDL_GetTrayEntries](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntries) | [:heavy_check_mark:](sdl/methods.go#L458) | [:x:](sdl/sdl_functions_js.go#L13303) |
| [SDL_RemoveTrayEntry](https://wiki.libsdl.org/SDL3/SDL_RemoveTrayEntry) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13317) |
| [SDL_InsertTrayEntryAt](https://wiki.libsdl.org/SDL3/SDL_InsertTrayEntryAt) | [:heavy_check_mark:](sdl/methods.go#L468) | [:x:](sdl/sdl_functions_js.go#L13325) |
| [SDL_SetTrayEntryLabel](https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryLabel) | [:heavy_check_mark:](sdl/methods.go#L494) | [:x:](sdl/sdl_functions_js.go#L13343) |
| [SDL_GetTrayEntryLabel](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryLabel) | [:heavy_check_mark:](sdl/methods.go#L500) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13356) |
| [SDL_SetTrayEntryChecked](https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryChecked) | [:heavy_check_mark:](sdl/methods.go#L506) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13365) |
| [SDL_GetTrayEntryChecked](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryChecked) | [:heavy_check_mark:](sdl/methods.go#L512) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13375) |
| [SDL_SetTrayEntryEnabled](https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryEnabled) | [:heavy_check_mark:](sdl/methods.go#L518) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13384) |
| [SDL_GetTrayEntryEnabled](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryEnabled) | [:heavy_check_mark:](sdl/methods.go#L524) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13394) |
| [SDL_SetTrayEntryCallback](https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryCallback) | [:heavy_check_mark:](sdl/methods.go#L530) | [:x:](sdl/sdl_functions_js.go#L13403) |
| [SDL_ClickTrayEntry](https://wiki.libsdl.org/SDL3/SDL_ClickTrayEntry) | [:heavy_check_mark:](sdl/methods.go#L536) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13416) |
| [SDL_DestroyTray](https://wiki.libsdl.org/SDL3/SDL_DestroyTray) | [:heavy_check_mark:](sdl/methods.go#L450) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13424) |
| [SDL_GetTrayEntryParent](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryParent) | [:heavy_check_mark:](sdl/methods.go#L542) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13433) |
| [SDL_GetTrayMenuParentEntry](https://wiki.libsdl.org/SDL3/SDL_GetTrayMenuParentEntry) | [:heavy_check_mark:](sdl/methods.go#L474) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13442) |
| [SDL_GetTrayMenuParentTray](https://wiki.libsdl.org/SDL3/SDL_GetTrayMenuParentTray) | [:heavy_check_mark:](sdl/methods.go#L480) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13451) |
| [SDL_UpdateTrays](https://wiki.libsdl.org/SDL3/SDL_UpdateTrays) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13460) |
</details>
<details open>
<summary><h3>Notification</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_RequestNotificationPermission](https://wiki.libsdl.org/SDL3/SDL_RequestNotificationPermission) | [:question:]() | [:question:]() |
| [SDL_ShowNotificationWithProperties](https://wiki.libsdl.org/SDL3/SDL_ShowNotificationWithProperties) | [:question:]() | [:question:]() |
| [SDL_ShowNotification](https://wiki.libsdl.org/SDL3/SDL_ShowNotification) | [:question:]() | [:question:]() |
| [SDL_RemoveNotification](https://wiki.libsdl.org/SDL3/SDL_RemoveNotification) | [:question:]() | [:question:]() |
</details>
<details open>
<summary><h3>GPU</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GPUSupportsShaderFormats](https://wiki.libsdl.org/SDL3/SDL_GPUSupportsShaderFormats) | [:heavy_check_mark:](sdl/functions.go#L564) | [:x:](sdl/sdl_functions_js.go#L8045) |
| [SDL_GPUSupportsProperties](https://wiki.libsdl.org/SDL3/SDL_GPUSupportsProperties) | [:heavy_check_mark:](sdl/methods.go#L5922) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8059) |
| [SDL_CreateGPUDevice](https://wiki.libsdl.org/SDL3/SDL_CreateGPUDevice) | [:heavy_check_mark:](sdl/functions.go#L574) | [:x:](sdl/sdl_functions_js.go#L8068) |
| [SDL_CreateGPUDeviceWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateGPUDeviceWithProperties) | [:heavy_check_mark:](sdl/functions.go#L589) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8084) |
| [SDL_DestroyGPUDevice](https://wiki.libsdl.org/SDL3/SDL_DestroyGPUDevice) | [:heavy_check_mark:](sdl/methods.go#L1984) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8093) |
| [SDL_GetNumGPUDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumGPUDrivers) | [:heavy_check_mark:](sdl/functions.go#L600) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8102) |
| [SDL_GetGPUDriver](https://wiki.libsdl.org/SDL3/SDL_GetGPUDriver) | [:heavy_check_mark:](sdl/functions.go#L606) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8109) |
| [SDL_GetGPUDeviceDriver](https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceDriver) | [:heavy_check_mark:](sdl/methods.go#L1990) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8118) |
| [SDL_GetGPUShaderFormats](https://wiki.libsdl.org/SDL3/SDL_GetGPUShaderFormats) | [:heavy_check_mark:](sdl/methods.go#L2002) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8127) |
| [SDL_GetGPUDeviceProperties](https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceProperties) | [:heavy_check_mark:](sdl/methods.go#L1996) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12894) |
| [SDL_CreateGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_CreateGPUComputePipeline) | [:heavy_check_mark:](sdl/methods.go#L2019) | [:x:](sdl/sdl_functions_js.go#L8136) |
| [SDL_CreateGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_CreateGPUGraphicsPipeline) | [:heavy_check_mark:](sdl/methods.go#L2031) | [:x:](sdl/sdl_functions_js.go#L8151) |
| [SDL_CreateGPUSampler](https://wiki.libsdl.org/SDL3/SDL_CreateGPUSampler) | [:heavy_check_mark:](sdl/methods.go#L2042) | [:x:](sdl/sdl_functions_js.go#L8166) |
| [SDL_CreateGPUShader](https://wiki.libsdl.org/SDL3/SDL_CreateGPUShader) | [:heavy_check_mark:](sdl/methods.go#L2053) | [:x:](sdl/sdl_functions_js.go#L13473) |
| [SDL_CreateGPUTexture](https://wiki.libsdl.org/SDL3/SDL_CreateGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2065) | [:x:](sdl/sdl_functions_js.go#L8204) |
| [SDL_CreateGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_CreateGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2076) | [:x:](sdl/sdl_functions_js.go#L8218) |
| [SDL_CreateGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_CreateGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2087) | [:x:](sdl/sdl_functions_js.go#L8232) |
| [SDL_SetGPUBufferName](https://wiki.libsdl.org/SDL3/SDL_SetGPUBufferName) | [:heavy_check_mark:](sdl/methods.go#L2098) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8246) |
| [SDL_SetGPUTextureName](https://wiki.libsdl.org/SDL3/SDL_SetGPUTextureName) | [:heavy_check_mark:](sdl/methods.go#L2104) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8260) |
| [SDL_InsertGPUDebugLabel](https://wiki.libsdl.org/SDL3/SDL_InsertGPUDebugLabel) | [:heavy_check_mark:](sdl/methods.go#L6023) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8274) |
| [SDL_PushGPUDebugGroup](https://wiki.libsdl.org/SDL3/SDL_PushGPUDebugGroup) | [:heavy_check_mark:](sdl/methods.go#L6029) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8286) |
| [SDL_PopGPUDebugGroup](https://wiki.libsdl.org/SDL3/SDL_PopGPUDebugGroup) | [:heavy_check_mark:](sdl/methods.go#L6035) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8298) |
| [SDL_ReleaseGPUTexture](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2110) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8306) |
| [SDL_ReleaseGPUSampler](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUSampler) | [:heavy_check_mark:](sdl/methods.go#L2116) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8316) |
| [SDL_ReleaseGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2122) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8326) |
| [SDL_ReleaseGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2128) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8336) |
| [SDL_ReleaseGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUComputePipeline) | [:heavy_check_mark:](sdl/methods.go#L2134) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8346) |
| [SDL_ReleaseGPUShader](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUShader) | [:heavy_check_mark:](sdl/methods.go#L2140) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8356) |
| [SDL_ReleaseGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUGraphicsPipeline) | [:heavy_check_mark:](sdl/methods.go#L2146) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8366) |
| [SDL_AcquireGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_AcquireGPUCommandBuffer) | [:heavy_check_mark:](sdl/methods.go#L2152) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8376) |
| [SDL_PushGPUVertexUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUVertexUniformData) | [:heavy_check_mark:](sdl/methods.go#L6041) | [:x:](sdl/sdl_functions_js.go#L8385) |
| [SDL_PushGPUFragmentUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUFragmentUniformData) | [:heavy_check_mark:](sdl/methods.go#L6047) | [:x:](sdl/sdl_functions_js.go#L8400) |
| [SDL_PushGPUComputeUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUComputeUniformData) | [:heavy_check_mark:](sdl/methods.go#L6053) | [:x:](sdl/sdl_functions_js.go#L8415) |
| [SDL_BeginGPURenderPass](https://wiki.libsdl.org/SDL3/SDL_BeginGPURenderPass) | [:heavy_check_mark:](sdl/methods.go#L6059) | [:x:](sdl/sdl_functions_js.go#L8430) |
| [SDL_BindGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_BindGPUGraphicsPipeline) | [:heavy_check_mark:](sdl/methods.go#L1299) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8448) |
| [SDL_SetGPUViewport](https://wiki.libsdl.org/SDL3/SDL_SetGPUViewport) | [:heavy_check_mark:](sdl/methods.go#L1305) | [:x:](sdl/sdl_functions_js.go#L8458) |
| [SDL_SetGPUScissor](https://wiki.libsdl.org/SDL3/SDL_SetGPUScissor) | [:heavy_check_mark:](sdl/methods.go#L1311) | [:x:](sdl/sdl_functions_js.go#L8471) |
| [SDL_SetGPUBlendConstants](https://wiki.libsdl.org/SDL3/SDL_SetGPUBlendConstants) | [:question:]() | [:question:]() |
| [SDL_SetGPUStencilReference](https://wiki.libsdl.org/SDL3/SDL_SetGPUStencilReference) | [:heavy_check_mark:](sdl/methods.go#L1317) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8484) |
| [SDL_BindGPUVertexBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexBuffers) | [:heavy_check_mark:](sdl/methods.go#L1323) | [:x:](sdl/sdl_functions_js.go#L8494) |
| [SDL_BindGPUIndexBuffer](https://wiki.libsdl.org/SDL3/SDL_BindGPUIndexBuffer) | [:heavy_check_mark:](sdl/methods.go#L1330) | [:x:](sdl/sdl_functions_js.go#L8511) |
| [SDL_BindGPUVertexSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexSamplers) | [:heavy_check_mark:](sdl/methods.go#L1336) | [:x:](sdl/sdl_functions_js.go#L8526) |
| [SDL_BindGPUVertexStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageTextures) | [:heavy_check_mark:](sdl/methods.go#L1343) | [:x:](sdl/sdl_functions_js.go#L8543) |
| [SDL_BindGPUVertexStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageBuffers) | [:heavy_check_mark:](sdl/methods.go#L1350) | [:x:](sdl/sdl_functions_js.go#L8560) |
| [SDL_BindGPUFragmentSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentSamplers) | [:heavy_check_mark:](sdl/methods.go#L1357) | [:x:](sdl/sdl_functions_js.go#L8577) |
| [SDL_BindGPUFragmentStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageTextures) | [:heavy_check_mark:](sdl/methods.go#L1364) | [:x:](sdl/sdl_functions_js.go#L8594) |
| [SDL_BindGPUFragmentStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageBuffers) | [:heavy_check_mark:](sdl/methods.go#L1371) | [:x:](sdl/sdl_functions_js.go#L8611) |
| [SDL_DrawGPUIndexedPrimitives](https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitives) | [:heavy_check_mark:](sdl/methods.go#L1378) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8628) |
| [SDL_DrawGPUPrimitives](https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitives) | [:heavy_check_mark:](sdl/methods.go#L1384) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8646) |
| [SDL_DrawGPUPrimitivesIndirect](https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitivesIndirect) | [:heavy_check_mark:](sdl/methods.go#L1390) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8662) |
| [SDL_DrawGPUIndexedPrimitivesIndirect](https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitivesIndirect) | [:heavy_check_mark:](sdl/methods.go#L1396) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8676) |
| [SDL_EndGPURenderPass](https://wiki.libsdl.org/SDL3/SDL_EndGPURenderPass) | [:heavy_check_mark:](sdl/methods.go#L1402) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8690) |
| [SDL_BeginGPUComputePass](https://wiki.libsdl.org/SDL3/SDL_BeginGPUComputePass) | [:heavy_check_mark:](sdl/methods.go#L6065) | [:x:](sdl/sdl_functions_js.go#L8698) |
| [SDL_BindGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputePipeline) | [:heavy_check_mark:](sdl/methods.go#L945) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8718) |
| [SDL_BindGPUComputeSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeSamplers) | [:heavy_check_mark:](sdl/methods.go#L951) | [:x:](sdl/sdl_functions_js.go#L8728) |
| [SDL_BindGPUComputeStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageTextures) | [:heavy_check_mark:](sdl/methods.go#L958) | [:x:](sdl/sdl_functions_js.go#L8745) |
| [SDL_BindGPUComputeStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageBuffers) | [:heavy_check_mark:](sdl/methods.go#L965) | [:x:](sdl/sdl_functions_js.go#L8762) |
| [SDL_DispatchGPUCompute](https://wiki.libsdl.org/SDL3/SDL_DispatchGPUCompute) | [:heavy_check_mark:](sdl/methods.go#L972) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8779) |
| [SDL_DispatchGPUComputeIndirect](https://wiki.libsdl.org/SDL3/SDL_DispatchGPUComputeIndirect) | [:heavy_check_mark:](sdl/methods.go#L978) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8793) |
| [SDL_EndGPUComputePass](https://wiki.libsdl.org/SDL3/SDL_EndGPUComputePass) | [:heavy_check_mark:](sdl/methods.go#L984) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8805) |
| [SDL_MapGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_MapGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2163) | [:x:](sdl/sdl_functions_js.go#L8813) |
| [SDL_UnmapGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_UnmapGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2174) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8827) |
| [SDL_BeginGPUCopyPass](https://wiki.libsdl.org/SDL3/SDL_BeginGPUCopyPass) | [:heavy_check_mark:](sdl/methods.go#L6071) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8837) |
| [SDL_UploadToGPUTexture](https://wiki.libsdl.org/SDL3/SDL_UploadToGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2817) | [:x:](sdl/sdl_functions_js.go#L8846) |
| [SDL_UploadToGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_UploadToGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2823) | [:x:](sdl/sdl_functions_js.go#L8863) |
| [SDL_CopyGPUTextureToTexture](https://wiki.libsdl.org/SDL3/SDL_CopyGPUTextureToTexture) | [:heavy_check_mark:](sdl/methods.go#L2829) | [:x:](sdl/sdl_functions_js.go#L8880) |
| [SDL_CopyGPUBufferToBuffer](https://wiki.libsdl.org/SDL3/SDL_CopyGPUBufferToBuffer) | [:heavy_check_mark:](sdl/methods.go#L2835) | [:x:](sdl/sdl_functions_js.go#L8903) |
| [SDL_DownloadFromGPUTexture](https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2841) | [:x:](sdl/sdl_functions_js.go#L8922) |
| [SDL_DownloadFromGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2847) | [:x:](sdl/sdl_functions_js.go#L8937) |
| [SDL_EndGPUCopyPass](https://wiki.libsdl.org/SDL3/SDL_EndGPUCopyPass) | [:heavy_check_mark:](sdl/methods.go#L2853) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8952) |
| [SDL_GenerateMipmapsForGPUTexture](https://wiki.libsdl.org/SDL3/SDL_GenerateMipmapsForGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L6077) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8960) |
| [SDL_BlitGPUTexture](https://wiki.libsdl.org/SDL3/SDL_BlitGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L6083) | [:x:](sdl/sdl_functions_js.go#L8970) |
| [SDL_WindowSupportsGPUSwapchainComposition](https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUSwapchainComposition) | [:heavy_check_mark:](sdl/methods.go#L2180) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8983) |
| [SDL_WindowSupportsGPUPresentMode](https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUPresentMode) | [:heavy_check_mark:](sdl/methods.go#L2186) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8996) |
| [SDL_ClaimWindowForGPUDevice](https://wiki.libsdl.org/SDL3/SDL_ClaimWindowForGPUDevice) | [:heavy_check_mark:](sdl/methods.go#L2192) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9009) |
| [SDL_ReleaseWindowFromGPUDevice](https://wiki.libsdl.org/SDL3/SDL_ReleaseWindowFromGPUDevice) | [:heavy_check_mark:](sdl/methods.go#L2202) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9020) |
| [SDL_SetGPUSwapchainParameters](https://wiki.libsdl.org/SDL3/SDL_SetGPUSwapchainParameters) | [:heavy_check_mark:](sdl/methods.go#L2208) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9030) |
| [SDL_SetGPUAllowedFramesInFlight](https://wiki.libsdl.org/SDL3/SDL_SetGPUAllowedFramesInFlight) | [:heavy_check_mark:](sdl/methods.go#L2218) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9045) |
| [SDL_GetGPUSwapchainTextureFormat](https://wiki.libsdl.org/SDL3/SDL_GetGPUSwapchainTextureFormat) | [:heavy_check_mark:](sdl/methods.go#L2228) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9056) |
| [SDL_AcquireGPUSwapchainTexture](https://wiki.libsdl.org/SDL3/SDL_AcquireGPUSwapchainTexture) | [:heavy_check_mark:](sdl/methods.go#L6089) | [:x:](sdl/sdl_functions_js.go#L9067) |
| [SDL_WaitForGPUSwapchain](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUSwapchain) | [:heavy_check_mark:](sdl/methods.go#L2234) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9087) |
| [SDL_WaitAndAcquireGPUSwapchainTexture](https://wiki.libsdl.org/SDL3/SDL_WaitAndAcquireGPUSwapchainTexture) | [:heavy_check_mark:](sdl/methods.go#L6101) | [:x:](sdl/sdl_functions_js.go#L9098) |
| [SDL_SubmitGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBuffer) | [:heavy_check_mark:](sdl/methods.go#L6113) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9118) |
| [SDL_SubmitGPUCommandBufferAndAcquireFence](https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBufferAndAcquireFence) | [:heavy_check_mark:](sdl/methods.go#L6123) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9127) |
| [SDL_CancelGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_CancelGPUCommandBuffer) | [:heavy_check_mark:](sdl/methods.go#L6134) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9136) |
| [SDL_WaitForGPUIdle](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUIdle) | [:heavy_check_mark:](sdl/methods.go#L2244) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9145) |
| [SDL_WaitForGPUFences](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUFences) | [:heavy_check_mark:](sdl/methods.go#L2254) | [:x:](sdl/sdl_functions_js.go#L9154) |
| [SDL_QueryGPUFence](https://wiki.libsdl.org/SDL3/SDL_QueryGPUFence) | [:heavy_check_mark:](sdl/methods.go#L2264) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9172) |
| [SDL_ReleaseGPUFence](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUFence) | [:heavy_check_mark:](sdl/methods.go#L2270) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9183) |
| [SDL_GPUTextureFormatTexelBlockSize](https://wiki.libsdl.org/SDL3/SDL_GPUTextureFormatTexelBlockSize) | [:heavy_check_mark:](sdl/methods.go#L558) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9193) |
| [SDL_GPUTextureSupportsFormat](https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsFormat) | [:heavy_check_mark:](sdl/methods.go#L2276) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9202) |
| [SDL_GPUTextureSupportsSampleCount](https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsSampleCount) | [:heavy_check_mark:](sdl/methods.go#L2282) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9217) |
| [SDL_CalculateGPUTextureFormatSize](https://wiki.libsdl.org/SDL3/SDL_CalculateGPUTextureFormatSize) | [:heavy_check_mark:](sdl/methods.go#L564) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9230) |
| [SDL_GetPixelFormatFromGPUTextureFormat](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatFromGPUTextureFormat) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12903) |
| [SDL_GetGPUTextureFormatFromPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetGPUTextureFormatFromPixelFormat) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12912) |
| [SDL_GDKSuspendGPU](https://wiki.libsdl.org/SDL3/SDL_GDKSuspendGPU) | [:question:]() | [:question:]() |
| [SDL_GDKResumeGPU](https://wiki.libsdl.org/SDL3/SDL_GDKResumeGPU) | [:question:]() | [:question:]() |
</details>
<details>
<summary><h3>Vulkan</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_Vulkan_LoadLibrary](https://wiki.libsdl.org/SDL3/SDL_Vulkan_LoadLibrary) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12183) |
| [SDL_Vulkan_GetVkGetInstanceProcAddr](https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetVkGetInstanceProcAddr) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L12194) |
| [SDL_Vulkan_UnloadLibrary](https://wiki.libsdl.org/SDL3/SDL_Vulkan_UnloadLibrary) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L12202) |
| [SDL_Vulkan_GetInstanceExtensions](https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetInstanceExtensions) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L12208) |
| [SDL_Vulkan_CreateSurface](https://wiki.libsdl.org/SDL3/SDL_Vulkan_CreateSurface) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L12220) |
| [SDL_Vulkan_DestroySurface](https://wiki.libsdl.org/SDL3/SDL_Vulkan_DestroySurface) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L12238) |
| [SDL_Vulkan_GetPresentationSupport](https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetPresentationSupport) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L12251) |
</details>
<details>
<summary><h3>Metal</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_Metal_CreateView](https://wiki.libsdl.org/SDL3/SDL_Metal_CreateView) | [:x:](sdl/methods.go#L4890) | [:x:](sdl/sdl_functions_js.go#L10051) |
| [SDL_Metal_DestroyView](https://wiki.libsdl.org/SDL3/SDL_Metal_DestroyView) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10061) |
| [SDL_Metal_GetLayer](https://wiki.libsdl.org/SDL3/SDL_Metal_GetLayer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10070) |
</details>
<details open>
<summary><h3>Power</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetPowerInfo) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L5536) |
</details>
<details open>
<summary><h3>Sensor</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetSensors](https://wiki.libsdl.org/SDL3/SDL_GetSensors) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L5550) |
| [SDL_GetSensorNameForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorNameForID) | [:heavy_check_mark:](sdl/methods.go#L6225) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5562) |
| [SDL_GetSensorTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorTypeForID) | [:heavy_check_mark:](sdl/methods.go#L6231) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5571) |
| [SDL_GetSensorNonPortableTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableTypeForID) | [:heavy_check_mark:](sdl/methods.go#L6237) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5580) |
| [SDL_OpenSensor](https://wiki.libsdl.org/SDL3/SDL_OpenSensor) | [:heavy_check_mark:](sdl/methods.go#L6243) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5589) |
| [SDL_GetSensorFromID](https://wiki.libsdl.org/SDL3/SDL_GetSensorFromID) | [:heavy_check_mark:](sdl/methods.go#L6254) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5598) |
| [SDL_GetSensorProperties](https://wiki.libsdl.org/SDL3/SDL_GetSensorProperties) | [:heavy_check_mark:](sdl/methods.go#L592) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5607) |
| [SDL_GetSensorName](https://wiki.libsdl.org/SDL3/SDL_GetSensorName) | [:heavy_check_mark:](sdl/methods.go#L603) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5616) |
| [SDL_GetSensorType](https://wiki.libsdl.org/SDL3/SDL_GetSensorType) | [:heavy_check_mark:](sdl/methods.go#L614) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5625) |
| [SDL_GetSensorNonPortableType](https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableType) | [:heavy_check_mark:](sdl/methods.go#L620) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5634) |
| [SDL_GetSensorID](https://wiki.libsdl.org/SDL3/SDL_GetSensorID) | [:heavy_check_mark:](sdl/methods.go#L626) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5643) |
| [SDL_GetSensorData](https://wiki.libsdl.org/SDL3/SDL_GetSensorData) | [:heavy_check_mark:](sdl/methods.go#L637) | [:x:](sdl/sdl_functions_js.go#L5652) |
| [SDL_CloseSensor](https://wiki.libsdl.org/SDL3/SDL_CloseSensor) | [:heavy_check_mark:](sdl/methods.go#L647) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5668) |
| [SDL_UpdateSensors](https://wiki.libsdl.org/SDL3/SDL_UpdateSensors) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5677) |
</details>
<details>
<summary><h3>Process</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateProcess](https://wiki.libsdl.org/SDL3/SDL_CreateProcess) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12999) |
| [SDL_CreateProcessWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateProcessWithProperties) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10111) |
| [SDL_GetProcessProperties](https://wiki.libsdl.org/SDL3/SDL_GetProcessProperties) | [:heavy_check_mark:](sdl/methods.go#L6146) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10120) |
| [SDL_ReadProcess](https://wiki.libsdl.org/SDL3/SDL_ReadProcess) | [:heavy_check_mark:](sdl/methods.go#L6157) | [:x:](sdl/sdl_functions_js.go#L10129) |
| [SDL_GetProcessInput](https://wiki.libsdl.org/SDL3/SDL_GetProcessInput) | [:heavy_check_mark:](sdl/methods.go#L6175) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10145) |
| [SDL_GetProcessOutput](https://wiki.libsdl.org/SDL3/SDL_GetProcessOutput) | [:heavy_check_mark:](sdl/methods.go#L6186) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10154) |
| [SDL_KillProcess](https://wiki.libsdl.org/SDL3/SDL_KillProcess) | [:heavy_check_mark:](sdl/methods.go#L6197) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10163) |
| [SDL_WaitProcess](https://wiki.libsdl.org/SDL3/SDL_WaitProcess) | [:heavy_check_mark:](sdl/methods.go#L6207) | [:x:](sdl/sdl_functions_js.go#L10174) |
| [SDL_DestroyProcess](https://wiki.libsdl.org/SDL3/SDL_DestroyProcess) | [:heavy_check_mark:](sdl/methods.go#L6217) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10190) |
</details>
<details>
<summary><h3>Bits</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_MostSignificantBitIndex32](https://wiki.libsdl.org/SDL3/SDL_MostSignificantBitIndex32) | [:question:]() | [:question:]() |
| [SDL_HasExactlyOneBitSet32](https://wiki.libsdl.org/SDL3/SDL_HasExactlyOneBitSet32) | [:question:]() | [:question:]() |
</details>
<details>
<summary><h3>Endian</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SwapFloat](https://wiki.libsdl.org/SDL3/SDL_SwapFloat) | [:question:]() | [:question:]() |
| [SDL_Swap16](https://wiki.libsdl.org/SDL3/SDL_Swap16) | [:question:]() | [:question:]() |
| [SDL_Swap32](https://wiki.libsdl.org/SDL3/SDL_Swap32) | [:question:]() | [:question:]() |
| [SDL_Swap64](https://wiki.libsdl.org/SDL3/SDL_Swap64) | [:question:]() | [:question:]() |
</details>
<details>
<summary><h3>Assert</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ReportAssertion](https://wiki.libsdl.org/SDL3/SDL_ReportAssertion) | [:question:]() | [:question:]() |
| [SDL_SetAssertionHandler](https://wiki.libsdl.org/SDL3/SDL_SetAssertionHandler) | [:question:]() | [:question:]() |
| [SDL_GetDefaultAssertionHandler](https://wiki.libsdl.org/SDL3/SDL_GetDefaultAssertionHandler) | [:question:]() | [:question:]() |
| [SDL_GetAssertionHandler](https://wiki.libsdl.org/SDL3/SDL_GetAssertionHandler) | [:question:]() | [:question:]() |
| [SDL_GetAssertionReport](https://wiki.libsdl.org/SDL3/SDL_GetAssertionReport) | [:question:]() | [:question:]() |
| [SDL_ResetAssertionReport](https://wiki.libsdl.org/SDL3/SDL_ResetAssertionReport) | [:question:]() | [:question:]() |
</details>
<details>
<summary><h3>CPUInfo</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumLogicalCPUCores](https://wiki.libsdl.org/SDL3/SDL_GetNumLogicalCPUCores) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4052) |
| [SDL_GetCPUCacheLineSize](https://wiki.libsdl.org/SDL3/SDL_GetCPUCacheLineSize) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4059) |
| [SDL_HasAltiVec](https://wiki.libsdl.org/SDL3/SDL_HasAltiVec) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4066) |
| [SDL_HasMMX](https://wiki.libsdl.org/SDL3/SDL_HasMMX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4073) |
| [SDL_HasSSE](https://wiki.libsdl.org/SDL3/SDL_HasSSE) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4080) |
| [SDL_HasSSE2](https://wiki.libsdl.org/SDL3/SDL_HasSSE2) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4087) |
| [SDL_HasSSE3](https://wiki.libsdl.org/SDL3/SDL_HasSSE3) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4094) |
| [SDL_HasSSE41](https://wiki.libsdl.org/SDL3/SDL_HasSSE41) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4101) |
| [SDL_HasSSE42](https://wiki.libsdl.org/SDL3/SDL_HasSSE42) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4108) |
| [SDL_HasAVX](https://wiki.libsdl.org/SDL3/SDL_HasAVX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4115) |
| [SDL_HasAVX2](https://wiki.libsdl.org/SDL3/SDL_HasAVX2) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4122) |
| [SDL_HasAVX512F](https://wiki.libsdl.org/SDL3/SDL_HasAVX512F) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4129) |
| [SDL_HasARMSIMD](https://wiki.libsdl.org/SDL3/SDL_HasARMSIMD) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4136) |
| [SDL_HasNEON](https://wiki.libsdl.org/SDL3/SDL_HasNEON) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4143) |
| [SDL_HasSVE2](https://wiki.libsdl.org/SDL3/SDL_HasSVE2) | [:question:]() | [:question:]() |
| [SDL_HasLSX](https://wiki.libsdl.org/SDL3/SDL_HasLSX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4150) |
| [SDL_HasLASX](https://wiki.libsdl.org/SDL3/SDL_HasLASX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4157) |
| [SDL_GetSystemRAM](https://wiki.libsdl.org/SDL3/SDL_GetSystemRAM) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4164) |
| [SDL_GetSIMDAlignment](https://wiki.libsdl.org/SDL3/SDL_GetSIMDAlignment) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4171) |
| [SDL_GetSystemPageSize](https://wiki.libsdl.org/SDL3/SDL_GetSystemPageSize) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12512) |
</details>
<details>
<summary><h3>Locale</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPreferredLocales](https://wiki.libsdl.org/SDL3/SDL_GetPreferredLocales) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9799) |
</details>
<details>
<summary><h3>System</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetWindowsMessageHook](https://wiki.libsdl.org/SDL3/SDL_SetWindowsMessageHook) | [:question:]() | [:question:]() |
| [SDL_GetDirect3D9AdapterIndex](https://wiki.libsdl.org/SDL3/SDL_GetDirect3D9AdapterIndex) | [:question:]() | [:question:]() |
| [SDL_GetDXGIOutputInfo](https://wiki.libsdl.org/SDL3/SDL_GetDXGIOutputInfo) | [:question:]() | [:question:]() |
| [SDL_SetX11EventHook](https://wiki.libsdl.org/SDL3/SDL_SetX11EventHook) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13191) |
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
| [SDL_IsPhone](https://wiki.libsdl.org/SDL3/SDL_IsPhone) | [:question:]() | [:question:]() |
| [SDL_IsTablet](https://wiki.libsdl.org/SDL3/SDL_IsTablet) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11863) |
| [SDL_IsTV](https://wiki.libsdl.org/SDL3/SDL_IsTV) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11870) |
| [SDL_GetDeviceFormFactor](https://wiki.libsdl.org/SDL3/SDL_GetDeviceFormFactor) | [:question:]() | [:question:]() |
| [SDL_GetDeviceFormFactorName](https://wiki.libsdl.org/SDL3/SDL_GetDeviceFormFactorName) | [:question:]() | [:question:]() |
| [SDL_GetSandbox](https://wiki.libsdl.org/SDL3/SDL_GetSandbox) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11877) |
| [SDL_OnApplicationWillTerminate](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillTerminate) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11884) |
| [SDL_OnApplicationDidReceiveMemoryWarning](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidReceiveMemoryWarning) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11890) |
| [SDL_OnApplicationWillEnterBackground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillEnterBackground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11896) |
| [SDL_OnApplicationDidEnterBackground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidEnterBackground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11902) |
| [SDL_OnApplicationWillEnterForeground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillEnterForeground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11908) |
| [SDL_OnApplicationDidEnterForeground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidEnterForeground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11914) |
| [SDL_OnApplicationDidChangeStatusBarOrientation](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidChangeStatusBarOrientation) | [:question:]() | [:question:]() |
| [SDL_GetGDKTaskQueue](https://wiki.libsdl.org/SDL3/SDL_GetGDKTaskQueue) | [:question:]() | [:question:]() |
| [SDL_GetGDKDefaultUser](https://wiki.libsdl.org/SDL3/SDL_GetGDKDefaultUser) | [:question:]() | [:question:]() |
| [SDL_IsUbuntuTouch](https://wiki.libsdl.org/SDL3/SDL_IsUbuntuTouch) | [:question:]() | [:question:]() |
</details>
<details>
<summary><h3>Misc</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_OpenURL](https://wiki.libsdl.org/SDL3/SDL_OpenURL) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10080) |
</details>
<details>
<summary><h3>GUID</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GUIDToString](https://wiki.libsdl.org/SDL3/SDL_GUIDToString) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12685) |
| [SDL_StringToGUID](https://wiki.libsdl.org/SDL3/SDL_StringToGUID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L12700) |
</details>
<details>
<summary><h3>Stdinc</h3></summary>

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
| [SDL_aligned_alloc_zero](https://wiki.libsdl.org/SDL3/SDL_aligned_alloc_zero) | [:question:]() | [:question:]() |
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
| [SDL_wcstoul](https://wiki.libsdl.org/SDL3/SDL_wcstoul) | [:question:]() | [:question:]() |
| [SDL_wcstoll](https://wiki.libsdl.org/SDL3/SDL_wcstoll) | [:question:]() | [:question:]() |
| [SDL_wcstoull](https://wiki.libsdl.org/SDL3/SDL_wcstoull) | [:question:]() | [:question:]() |
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
</details>
</details>
<details open>
<summary><h2>TTF</h2></summary>
<details open>
<summary><h3>TTF</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [TTF_Version](https://wiki.libsdl.org/SDL3_ttf/TTF_Version) | [:heavy_check_mark:](ttf/functions.go#L11) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L13) |
| [TTF_GetFreeTypeVersion](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFreeTypeVersion) | [:heavy_check_mark:](ttf/functions.go#L17) | [:x:](ttf/ttf_functions_js.go#L20) |
| [TTF_GetHarfBuzzVersion](https://wiki.libsdl.org/SDL3_ttf/TTF_GetHarfBuzzVersion) | [:heavy_check_mark:](ttf/functions.go#L27) | [:x:](ttf/ttf_functions_js.go#L35) |
| [TTF_Init](https://wiki.libsdl.org/SDL3_ttf/TTF_Init) | [:heavy_check_mark:](ttf/functions.go#L37) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L50) |
| [TTF_OpenFont](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFont) | [:heavy_check_mark:](ttf/functions.go#L47) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L58) |
| [TTF_OpenFontIO](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontIO) | [:heavy_check_mark:](ttf/functions.go#L58) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L71) |
| [TTF_OpenFontWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontWithProperties) | [:heavy_check_mark:](ttf/functions.go#L69) | [:x:](ttf/ttf_functions_js.go#L1591) |
| [TTF_CopyFont](https://wiki.libsdl.org/SDL3_ttf/TTF_CopyFont) | [:heavy_check_mark:](ttf/methods.go#L409) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L106) |
| [TTF_GetFontProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontProperties) | [:heavy_check_mark:](ttf/methods.go#L420) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L115) |
| [TTF_GetFontGeneration](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontGeneration) | [:heavy_check_mark:](ttf/methods.go#L426) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L124) |
| [TTF_AddFallbackFont](https://wiki.libsdl.org/SDL3_ttf/TTF_AddFallbackFont) | [:heavy_check_mark:](ttf/methods.go#L432) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L133) |
| [TTF_RemoveFallbackFont](https://wiki.libsdl.org/SDL3_ttf/TTF_RemoveFallbackFont) | [:heavy_check_mark:](ttf/methods.go#L442) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L144) |
| [TTF_ClearFallbackFonts](https://wiki.libsdl.org/SDL3_ttf/TTF_ClearFallbackFonts) | [:heavy_check_mark:](ttf/methods.go#L448) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L154) |
| [TTF_SetFontSize](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSize) | [:heavy_check_mark:](ttf/methods.go#L454) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L162) |
| [TTF_SetFontSizeDPI](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSizeDPI) | [:heavy_check_mark:](ttf/methods.go#L464) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L173) |
| [TTF_GetFontSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSize) | [:heavy_check_mark:](ttf/methods.go#L474) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L188) |
| [TTF_GetFontDPI](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDPI) | [:heavy_check_mark:](ttf/methods.go#L485) | [:x:](ttf/ttf_functions_js.go#L197) |
| [TTF_SetFontStyle](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontStyle) | [:heavy_check_mark:](ttf/methods.go#L497) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L213) |
| [TTF_GetFontStyle](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyle) | [:heavy_check_mark:](ttf/methods.go#L503) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L223) |
| [TTF_SetFontOutline](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontOutline) | [:heavy_check_mark:](ttf/methods.go#L509) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L232) |
| [TTF_GetFontOutline](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontOutline) | [:heavy_check_mark:](ttf/methods.go#L519) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L243) |
| [TTF_SetFontHinting](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontHinting) | [:heavy_check_mark:](ttf/methods.go#L525) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L252) |
| [TTF_GetNumFontFaces](https://wiki.libsdl.org/SDL3_ttf/TTF_GetNumFontFaces) | [:heavy_check_mark:](ttf/methods.go#L531) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L262) |
| [TTF_GetFontHinting](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHinting) | [:heavy_check_mark:](ttf/methods.go#L537) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L271) |
| [TTF_SetFontSDF](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSDF) | [:heavy_check_mark:](ttf/methods.go#L543) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L280) |
| [TTF_GetFontSDF](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSDF) | [:heavy_check_mark:](ttf/methods.go#L553) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L291) |
| [TTF_GetFontWeight](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontWeight) | [:question:]() | [:question:](ttf/ttf_functions_js.go#L1600) |
| [TTF_SetFontWrapAlignment](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontWrapAlignment) | [:heavy_check_mark:](ttf/methods.go#L559) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L300) |
| [TTF_GetFontWrapAlignment](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontWrapAlignment) | [:heavy_check_mark:](ttf/methods.go#L565) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L310) |
| [TTF_GetFontHeight](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHeight) | [:heavy_check_mark:](ttf/methods.go#L571) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L319) |
| [TTF_GetFontAscent](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontAscent) | [:heavy_check_mark:](ttf/methods.go#L577) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L328) |
| [TTF_GetFontDescent](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDescent) | [:heavy_check_mark:](ttf/methods.go#L583) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L337) |
| [TTF_SetFontLineSkip](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLineSkip) | [:heavy_check_mark:](ttf/methods.go#L589) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L346) |
| [TTF_GetFontLineSkip](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontLineSkip) | [:heavy_check_mark:](ttf/methods.go#L595) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L356) |
| [TTF_SetFontKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontKerning) | [:heavy_check_mark:](ttf/methods.go#L601) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L365) |
| [TTF_GetFontKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontKerning) | [:heavy_check_mark:](ttf/methods.go#L607) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L375) |
| [TTF_FontIsFixedWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsFixedWidth) | [:heavy_check_mark:](ttf/methods.go#L613) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L384) |
| [TTF_FontIsScalable](https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsScalable) | [:heavy_check_mark:](ttf/methods.go#L619) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L393) |
| [TTF_GetFontFamilyName](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontFamilyName) | [:heavy_check_mark:](ttf/methods.go#L625) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L402) |
| [TTF_GetFontStyleName](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyleName) | [:heavy_check_mark:](ttf/methods.go#L631) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L411) |
| [TTF_SetFontDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontDirection) | [:heavy_check_mark:](ttf/methods.go#L637) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L420) |
| [TTF_GetFontDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDirection) | [:heavy_check_mark:](ttf/methods.go#L647) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L431) |
| [TTF_SetFontCharSpacing](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontCharSpacing) | [:question:]() | [:question:]() |
| [TTF_GetFontCharSpacing](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontCharSpacing) | [:question:]() | [:question:]() |
| [TTF_StringToTag](https://wiki.libsdl.org/SDL3_ttf/TTF_StringToTag) | [:heavy_check_mark:](ttf/functions.go#L80) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L440) |
| [TTF_TagToString](https://wiki.libsdl.org/SDL3_ttf/TTF_TagToString) | [:heavy_check_mark:](ttf/functions.go#L86) | [:x:](ttf/ttf_functions_js.go#L1314) |
| [TTF_SetFontScript](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontScript) | [:heavy_check_mark:](ttf/methods.go#L653) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L451) |
| [TTF_GetFontScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontScript) | [:heavy_check_mark:](ttf/methods.go#L663) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L462) |
| [TTF_GetGlyphScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphScript) | [:heavy_check_mark:](ttf/functions.go#L100) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L471) |
| [TTF_SetFontLanguage](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLanguage) | [:heavy_check_mark:](ttf/methods.go#L669) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L480) |
| [TTF_FontHasGlyph](https://wiki.libsdl.org/SDL3_ttf/TTF_FontHasGlyph) | [:heavy_check_mark:](ttf/methods.go#L679) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L493) |
| [TTF_GetGlyphImage](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImage) | [:heavy_check_mark:](ttf/methods.go#L685) | [:x:](ttf/ttf_functions_js.go#L504) |
| [TTF_GetGlyphImageForIndex](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImageForIndex) | [:heavy_check_mark:](ttf/methods.go#L698) | [:x:](ttf/ttf_functions_js.go#L520) |
| [TTF_GetGlyphMetrics](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphMetrics) | [:heavy_check_mark:](ttf/methods.go#L711) | [:x:](ttf/ttf_functions_js.go#L536) |
| [TTF_GetGlyphKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphKerning) | [:heavy_check_mark:](ttf/methods.go#L723) | [:x:](ttf/ttf_functions_js.go#L560) |
| [TTF_GetStringSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSize) | [:heavy_check_mark:](ttf/methods.go#L735) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L578) |
| [TTF_GetStringSizeWrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSizeWrapped) | [:heavy_check_mark:](ttf/methods.go#L747) | [:x:](ttf/ttf_functions_js.go#L603) |
| [TTF_MeasureString](https://wiki.libsdl.org/SDL3_ttf/TTF_MeasureString) | [:heavy_check_mark:](ttf/methods.go#L760) | [:x:](ttf/ttf_functions_js.go#L625) |
| [TTF_RenderText_Solid](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid) | [:heavy_check_mark:](ttf/methods.go#L773) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1330) |
| [TTF_RenderText_Solid_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid_Wrapped) | [:heavy_check_mark:](ttf/methods.go#L784) | [:x:](ttf/ttf_functions_js.go#L1353) |
| [TTF_RenderGlyph_Solid](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Solid) | [:heavy_check_mark:](ttf/methods.go#L795) | [:x:](ttf/ttf_functions_js.go#L1647) |
| [TTF_RenderText_Shaded](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded) | [:heavy_check_mark:](ttf/methods.go#L806) | [:x:](ttf/ttf_functions_js.go#L1662) |
| [TTF_RenderText_Shaded_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded_Wrapped) | [:heavy_check_mark:](ttf/methods.go#L817) | [:x:](ttf/ttf_functions_js.go#L1683) |
| [TTF_RenderGlyph_Shaded](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Shaded) | [:heavy_check_mark:](ttf/methods.go#L828) | [:x:](ttf/ttf_functions_js.go#L1706) |
| [TTF_RenderText_Blended](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended) | [:heavy_check_mark:](ttf/methods.go#L839) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1462) |
| [TTF_RenderText_Blended_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended_Wrapped) | [:heavy_check_mark:](ttf/methods.go#L850) | [:x:](ttf/ttf_functions_js.go#L1484) |
| [TTF_RenderGlyph_Blended](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Blended) | [:heavy_check_mark:](ttf/methods.go#L861) | [:x:](ttf/ttf_functions_js.go#L1723) |
| [TTF_RenderText_LCD](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD) | [:heavy_check_mark:](ttf/methods.go#L872) | [:x:](ttf/ttf_functions_js.go#L1738) |
| [TTF_RenderText_LCD_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD_Wrapped) | [:heavy_check_mark:](ttf/methods.go#L883) | [:x:](ttf/ttf_functions_js.go#L1759) |
| [TTF_RenderGlyph_LCD](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_LCD) | [:heavy_check_mark:](ttf/methods.go#L894) | [:x:](ttf/ttf_functions_js.go#L1782) |
| [TTF_CreateSurfaceTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateSurfaceTextEngine) | [:heavy_check_mark:](ttf/functions.go#L111) | [:x:](ttf/ttf_functions_js.go#L647) |
| [TTF_DrawSurfaceText](https://wiki.libsdl.org/SDL3_ttf/TTF_DrawSurfaceText) | [:heavy_check_mark:](ttf/methods.go#L13) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L655) |
| [TTF_DestroySurfaceTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroySurfaceTextEngine) | [:heavy_check_mark:](ttf/methods.go#L366) | [:x:](ttf/ttf_functions_js.go#L670) |
| [TTF_CreateRendererTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngine) | [:heavy_check_mark:](ttf/functions.go#L122) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L681) |
| [TTF_CreateRendererTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngineWithProperties) | [:heavy_check_mark:](ttf/functions.go#L133) | [:x:](ttf/ttf_functions_js.go#L1609) |
| [TTF_DrawRendererText](https://wiki.libsdl.org/SDL3_ttf/TTF_DrawRendererText) | [:heavy_check_mark:](ttf/methods.go#L23) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L712) |
| [TTF_DestroyRendererTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyRendererTextEngine) | [:heavy_check_mark:](ttf/methods.go#L372) | [:x:](ttf/ttf_functions_js.go#L729) |
| [TTF_CreateGPUTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngine) | [:heavy_check_mark:](ttf/functions.go#L144) | [:x:](ttf/ttf_functions_js.go#L740) |
| [TTF_CreateGPUTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngineWithProperties) | [:heavy_check_mark:](ttf/functions.go#L154) | [:x:](ttf/ttf_functions_js.go#L1619) |
| [TTF_GetGPUTextDrawData](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextDrawData) | [:heavy_check_mark:](ttf/methods.go#L33) | [:x:](ttf/ttf_functions_js.go#L767) |
| [TTF_DestroyGPUTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyGPUTextEngine) | [:heavy_check_mark:](ttf/methods.go#L378) | [:x:](ttf/ttf_functions_js.go#L777) |
| [TTF_SetGPUTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_SetGPUTextEngineWinding) | [:heavy_check_mark:](ttf/methods.go#L384) | [:x:](ttf/ttf_functions_js.go#L788) |
| [TTF_GetGPUTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextEngineWinding) | [:heavy_check_mark:](ttf/methods.go#L390) | [:x:](ttf/ttf_functions_js.go#L801) |
| [TTF_CreateGLTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGLTextEngine) | [:question:]() | [:question:]() |
| [TTF_CreateGLTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGLTextEngineWithProperties) | [:question:]() | [:question:]() |
| [TTF_GetGLTextDrawData](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGLTextDrawData) | [:question:]() | [:question:]() |
| [TTF_DestroyGLTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyGLTextEngine) | [:question:]() | [:question:]() |
| [TTF_SetGLTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_SetGLTextEngineWinding) | [:question:]() | [:question:]() |
| [TTF_GetGLTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGLTextEngineWinding) | [:question:]() | [:question:]() |
| [TTF_CreateText](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateText) | [:heavy_check_mark:](ttf/methods.go#L396) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L813) |
| [TTF_GetTextProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextProperties) | [:heavy_check_mark:](ttf/methods.go#L44) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L839) |
| [TTF_SetTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextEngine) | [:heavy_check_mark:](ttf/methods.go#L50) | [:x:](ttf/ttf_functions_js.go#L848) |
| [TTF_GetTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextEngine) | [:heavy_check_mark:](ttf/methods.go#L60) | [:x:](ttf/ttf_functions_js.go#L862) |
| [TTF_SetTextFont](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextFont) | [:heavy_check_mark:](ttf/methods.go#L71) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L872) |
| [TTF_GetTextFont](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextFont) | [:heavy_check_mark:](ttf/methods.go#L81) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L883) |
| [TTF_SetTextDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextDirection) | [:heavy_check_mark:](ttf/methods.go#L92) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L892) |
| [TTF_GetTextDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextDirection) | [:heavy_check_mark:](ttf/methods.go#L102) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L903) |
| [TTF_SetTextScript](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextScript) | [:heavy_check_mark:](ttf/methods.go#L108) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L912) |
| [TTF_GetTextScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextScript) | [:heavy_check_mark:](ttf/methods.go#L118) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L923) |
| [TTF_SetTextColor](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColor) | [:heavy_check_mark:](ttf/methods.go#L124) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L932) |
| [TTF_SetTextColorFloat](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColorFloat) | [:heavy_check_mark:](ttf/methods.go#L134) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L953) |
| [TTF_GetTextColor](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColor) | [:heavy_check_mark:](ttf/methods.go#L144) | [:x:](ttf/ttf_functions_js.go#L970) |
| [TTF_GetTextColorFloat](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColorFloat) | [:heavy_check_mark:](ttf/methods.go#L156) | [:x:](ttf/ttf_functions_js.go#L990) |
| [TTF_SetTextPosition](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextPosition) | [:heavy_check_mark:](ttf/methods.go#L168) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1010) |
| [TTF_GetTextPosition](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextPosition) | [:heavy_check_mark:](ttf/methods.go#L174) | [:x:](ttf/ttf_functions_js.go#L1023) |
| [TTF_SetTextWrapWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWidth) | [:heavy_check_mark:](ttf/methods.go#L184) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1039) |
| [TTF_GetTextWrapWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextWrapWidth) | [:heavy_check_mark:](ttf/methods.go#L194) | [:x:](ttf/ttf_functions_js.go#L1050) |
| [TTF_SetTextWrapWhitespaceVisible](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWhitespaceVisible) | [:heavy_check_mark:](ttf/methods.go#L206) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1064) |
| [TTF_TextWrapWhitespaceVisible](https://wiki.libsdl.org/SDL3_ttf/TTF_TextWrapWhitespaceVisible) | [:heavy_check_mark:](ttf/methods.go#L216) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1075) |
| [TTF_SetTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextString) | [:heavy_check_mark:](ttf/methods.go#L222) | [:x:](ttf/ttf_functions_js.go#L1084) |
| [TTF_InsertTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_InsertTextString) | [:heavy_check_mark:](ttf/methods.go#L232) | [:x:](ttf/ttf_functions_js.go#L1100) |
| [TTF_AppendTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_AppendTextString) | [:heavy_check_mark:](ttf/methods.go#L242) | [:x:](ttf/ttf_functions_js.go#L1118) |
| [TTF_DeleteTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_DeleteTextString) | [:heavy_check_mark:](ttf/methods.go#L252) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1134) |
| [TTF_GetTextSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSize) | [:heavy_check_mark:](ttf/methods.go#L262) | [:x:](ttf/ttf_functions_js.go#L1147) |
| [TTF_GetTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubString) | [:heavy_check_mark:](ttf/methods.go#L274) | [:x:](ttf/ttf_functions_js.go#L1163) |
| [TTF_GetTextSubStringForLine](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForLine) | [:heavy_check_mark:](ttf/methods.go#L286) | [:x:](ttf/ttf_functions_js.go#L1179) |
| [TTF_GetTextSubStringsForRange](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringsForRange) | [:heavy_check_mark:](ttf/methods.go#L298) | [:x:](ttf/ttf_functions_js.go#L1629) |
| [TTF_GetTextSubStringForPoint](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForPoint) | [:heavy_check_mark:](ttf/methods.go#L312) | [:x:](ttf/ttf_functions_js.go#L1221) |
| [TTF_GetPreviousTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetPreviousTextSubString) | [:heavy_check_mark:](ttf/methods.go#L324) | [:x:](ttf/ttf_functions_js.go#L1239) |
| [TTF_GetNextTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetNextTextSubString) | [:heavy_check_mark:](ttf/methods.go#L336) | [:x:](ttf/ttf_functions_js.go#L1255) |
| [TTF_UpdateText](https://wiki.libsdl.org/SDL3_ttf/TTF_UpdateText) | [:heavy_check_mark:](ttf/methods.go#L348) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1271) |
| [TTF_DestroyText](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyText) | [:heavy_check_mark:](ttf/methods.go#L358) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1280) |
| [TTF_CloseFont](https://wiki.libsdl.org/SDL3_ttf/TTF_CloseFont) | [:heavy_check_mark:](ttf/methods.go#L905) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1292) |
| [TTF_Quit](https://wiki.libsdl.org/SDL3_ttf/TTF_Quit) | [:heavy_check_mark:](ttf/functions.go#L164) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1301) |
| [TTF_WasInit](https://wiki.libsdl.org/SDL3_ttf/TTF_WasInit) | [:heavy_check_mark:](ttf/functions.go#L170) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1307) |
</details>
</details>
<details open>
<summary><h2>IMG</h2></summary>
<details open>
<summary><h3>Image</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [IMG_Version](https://wiki.libsdl.org/SDL3_image/IMG_Version) | [:heavy_check_mark:](img/functions.go#L11) | [:heavy_check_mark:](img/img_functions_js.go#L12) |
| [IMG_Load](https://wiki.libsdl.org/SDL3_image/IMG_Load) | [:heavy_check_mark:](img/functions.go#L28) | [:heavy_check_mark:](img/img_functions_js.go#L34) |
| [IMG_Load_IO](https://wiki.libsdl.org/SDL3_image/IMG_Load_IO) | [:heavy_check_mark:](img/functions.go#L39) | [:heavy_check_mark:](img/img_functions_js.go#L45) |
| [IMG_LoadTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTyped_IO) | [:heavy_check_mark:](img/functions.go#L17) | [:heavy_check_mark:](img/img_functions_js.go#L19) |
| [IMG_LoadTexture](https://wiki.libsdl.org/SDL3_image/IMG_LoadTexture) | [:heavy_check_mark:](img/functions.go#L50) | [:heavy_check_mark:](img/img_functions_js.go#L56) |
| [IMG_LoadTexture_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTexture_IO) | [:heavy_check_mark:](img/functions.go#L61) | [:heavy_check_mark:](img/img_functions_js.go#L69) |
| [IMG_LoadTextureTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTextureTyped_IO) | [:heavy_check_mark:](img/functions.go#L72) | [:heavy_check_mark:](img/img_functions_js.go#L82) |
| [IMG_LoadGPUTexture](https://wiki.libsdl.org/SDL3_image/IMG_LoadGPUTexture) | [:question:]() | [:question:]() |
| [IMG_LoadGPUTexture_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadGPUTexture_IO) | [:question:]() | [:question:]() |
| [IMG_LoadGPUTextureTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadGPUTextureTyped_IO) | [:question:]() | [:question:]() |
| [IMG_GetClipboardImage](https://wiki.libsdl.org/SDL3_image/IMG_GetClipboardImage) | [:question:]() | [:question:]() |
| [IMG_isANI](https://wiki.libsdl.org/SDL3_image/IMG_isANI) | [:question:]() | [:question:]() |
| [IMG_isAVIF](https://wiki.libsdl.org/SDL3_image/IMG_isAVIF) | [:question:]() | [:question:]() |
| [IMG_isCUR](https://wiki.libsdl.org/SDL3_image/IMG_isCUR) | [:question:]() | [:question:]() |
| [IMG_isBMP](https://wiki.libsdl.org/SDL3_image/IMG_isBMP) | [:question:]() | [:question:]() |
| [IMG_isGIF](https://wiki.libsdl.org/SDL3_image/IMG_isGIF) | [:question:]() | [:question:]() |
| [IMG_isICO](https://wiki.libsdl.org/SDL3_image/IMG_isICO) | [:question:]() | [:question:]() |
| [IMG_isJPG](https://wiki.libsdl.org/SDL3_image/IMG_isJPG) | [:question:]() | [:question:]() |
| [IMG_isJXL](https://wiki.libsdl.org/SDL3_image/IMG_isJXL) | [:question:]() | [:question:]() |
| [IMG_isLBM](https://wiki.libsdl.org/SDL3_image/IMG_isLBM) | [:question:]() | [:question:]() |
| [IMG_isPCX](https://wiki.libsdl.org/SDL3_image/IMG_isPCX) | [:question:]() | [:question:]() |
| [IMG_isPNG](https://wiki.libsdl.org/SDL3_image/IMG_isPNG) | [:question:]() | [:question:]() |
| [IMG_isPNM](https://wiki.libsdl.org/SDL3_image/IMG_isPNM) | [:question:]() | [:question:]() |
| [IMG_isQOI](https://wiki.libsdl.org/SDL3_image/IMG_isQOI) | [:question:]() | [:question:]() |
| [IMG_isSVG](https://wiki.libsdl.org/SDL3_image/IMG_isSVG) | [:question:]() | [:question:]() |
| [IMG_isTIF](https://wiki.libsdl.org/SDL3_image/IMG_isTIF) | [:question:]() | [:question:]() |
| [IMG_isWEBP](https://wiki.libsdl.org/SDL3_image/IMG_isWEBP) | [:question:]() | [:question:]() |
| [IMG_isXCF](https://wiki.libsdl.org/SDL3_image/IMG_isXCF) | [:question:]() | [:question:]() |
| [IMG_isXPM](https://wiki.libsdl.org/SDL3_image/IMG_isXPM) | [:question:]() | [:question:]() |
| [IMG_isXV](https://wiki.libsdl.org/SDL3_image/IMG_isXV) | [:question:]() | [:question:]() |
| [IMG_LoadAVIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAVIF_IO) | [:heavy_check_mark:](img/functions.go#L191) | [:heavy_check_mark:](img/img_functions_js.go#L99) |
| [IMG_LoadBMP_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadBMP_IO) | [:heavy_check_mark:](img/functions.go#L224) | [:heavy_check_mark:](img/img_functions_js.go#L126) |
| [IMG_LoadCUR_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadCUR_IO) | [:heavy_check_mark:](img/functions.go#L213) | [:heavy_check_mark:](img/img_functions_js.go#L117) |
| [IMG_LoadGIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadGIF_IO) | [:heavy_check_mark:](img/functions.go#L235) | [:heavy_check_mark:](img/img_functions_js.go#L135) |
| [IMG_LoadICO_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadICO_IO) | [:heavy_check_mark:](img/functions.go#L202) | [:heavy_check_mark:](img/img_functions_js.go#L108) |
| [IMG_LoadJPG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadJPG_IO) | [:heavy_check_mark:](img/functions.go#L246) | [:heavy_check_mark:](img/img_functions_js.go#L144) |
| [IMG_LoadJXL_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadJXL_IO) | [:heavy_check_mark:](img/functions.go#L257) | [:heavy_check_mark:](img/img_functions_js.go#L153) |
| [IMG_LoadLBM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadLBM_IO) | [:heavy_check_mark:](img/functions.go#L268) | [:heavy_check_mark:](img/img_functions_js.go#L162) |
| [IMG_LoadPCX_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPCX_IO) | [:heavy_check_mark:](img/functions.go#L279) | [:heavy_check_mark:](img/img_functions_js.go#L171) |
| [IMG_LoadPNG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPNG_IO) | [:heavy_check_mark:](img/functions.go#L290) | [:heavy_check_mark:](img/img_functions_js.go#L180) |
| [IMG_LoadPNM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPNM_IO) | [:heavy_check_mark:](img/functions.go#L301) | [:heavy_check_mark:](img/img_functions_js.go#L189) |
| [IMG_LoadSVG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadSVG_IO) | [:heavy_check_mark:](img/functions.go#L312) | [:heavy_check_mark:](img/img_functions_js.go#L198) |
| [IMG_LoadSizedSVG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadSizedSVG_IO) | [:heavy_check_mark:](img/functions.go#L400) | [:heavy_check_mark:](img/img_functions_js.go#L270) |
| [IMG_LoadQOI_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadQOI_IO) | [:heavy_check_mark:](img/functions.go#L323) | [:heavy_check_mark:](img/img_functions_js.go#L207) |
| [IMG_LoadTGA_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTGA_IO) | [:heavy_check_mark:](img/functions.go#L334) | [:heavy_check_mark:](img/img_functions_js.go#L216) |
| [IMG_LoadTIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTIF_IO) | [:heavy_check_mark:](img/functions.go#L345) | [:heavy_check_mark:](img/img_functions_js.go#L225) |
| [IMG_LoadWEBP_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadWEBP_IO) | [:heavy_check_mark:](img/functions.go#L389) | [:heavy_check_mark:](img/img_functions_js.go#L261) |
| [IMG_LoadXCF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXCF_IO) | [:heavy_check_mark:](img/functions.go#L356) | [:heavy_check_mark:](img/img_functions_js.go#L234) |
| [IMG_LoadXPM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXPM_IO) | [:heavy_check_mark:](img/functions.go#L367) | [:heavy_check_mark:](img/img_functions_js.go#L243) |
| [IMG_LoadXV_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXV_IO) | [:heavy_check_mark:](img/functions.go#L378) | [:heavy_check_mark:](img/img_functions_js.go#L252) |
| [IMG_ReadXPMFromArray](https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArray) | [:heavy_check_mark:](img/functions.go#L411) | [:x:](img/img_functions_js.go#L628) |
| [IMG_ReadXPMFromArrayToRGB888](https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArrayToRGB888) | [:heavy_check_mark:](img/functions.go#L423) | [:x:](img/img_functions_js.go#L640) |
| [IMG_Save](https://wiki.libsdl.org/SDL3_image/IMG_Save) | [:question:]() | [:question:]() |
| [IMG_SaveTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveTyped_IO) | [:question:]() | [:question:]() |
| [IMG_SaveAVIF](https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF) | [:heavy_check_mark:](img/functions.go#L435) | [:heavy_check_mark:](img/img_functions_js.go#L317) |
| [IMG_SaveAVIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF_IO) | [:heavy_check_mark:](img/functions.go#L445) | [:heavy_check_mark:](img/img_functions_js.go#L332) |
| [IMG_SaveBMP](https://wiki.libsdl.org/SDL3_image/IMG_SaveBMP) | [:question:]() | [:question:]() |
| [IMG_SaveBMP_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveBMP_IO) | [:question:]() | [:question:]() |
| [IMG_SaveCUR](https://wiki.libsdl.org/SDL3_image/IMG_SaveCUR) | [:question:]() | [:question:]() |
| [IMG_SaveCUR_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveCUR_IO) | [:question:]() | [:question:]() |
| [IMG_SaveGIF](https://wiki.libsdl.org/SDL3_image/IMG_SaveGIF) | [:question:]() | [:question:]() |
| [IMG_SaveGIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveGIF_IO) | [:question:]() | [:question:]() |
| [IMG_SaveICO](https://wiki.libsdl.org/SDL3_image/IMG_SaveICO) | [:question:]() | [:question:]() |
| [IMG_SaveICO_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveICO_IO) | [:question:]() | [:question:]() |
| [IMG_SaveJPG](https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG) | [:heavy_check_mark:](img/functions.go#L475) | [:heavy_check_mark:](img/img_functions_js.go#L373) |
| [IMG_SaveJPG_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG_IO) | [:heavy_check_mark:](img/functions.go#L485) | [:heavy_check_mark:](img/img_functions_js.go#L388) |
| [IMG_SavePNG](https://wiki.libsdl.org/SDL3_image/IMG_SavePNG) | [:heavy_check_mark:](img/functions.go#L455) | [:heavy_check_mark:](img/img_functions_js.go#L347) |
| [IMG_SavePNG_IO](https://wiki.libsdl.org/SDL3_image/IMG_SavePNG_IO) | [:heavy_check_mark:](img/functions.go#L465) | [:heavy_check_mark:](img/img_functions_js.go#L360) |
| [IMG_SaveTGA](https://wiki.libsdl.org/SDL3_image/IMG_SaveTGA) | [:question:]() | [:question:]() |
| [IMG_SaveTGA_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveTGA_IO) | [:question:]() | [:question:]() |
| [IMG_SaveWEBP](https://wiki.libsdl.org/SDL3_image/IMG_SaveWEBP) | [:question:]() | [:question:]() |
| [IMG_SaveWEBP_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveWEBP_IO) | [:question:]() | [:question:]() |
| [IMG_LoadAnimation](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation) | [:heavy_check_mark:](img/functions.go#L495) | [:heavy_check_mark:](img/img_functions_js.go#L403) |
| [IMG_LoadAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation_IO) | [:heavy_check_mark:](img/functions.go#L506) | [:heavy_check_mark:](img/img_functions_js.go#L414) |
| [IMG_LoadAnimationTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimationTyped_IO) | [:heavy_check_mark:](img/functions.go#L517) | [:heavy_check_mark:](img/img_functions_js.go#L425) |
| [IMG_LoadANIAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadANIAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_LoadAPNGAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAPNGAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_LoadAVIFAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAVIFAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_LoadGIFAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadGIFAnimation_IO) | [:heavy_check_mark:](img/functions.go#L528) | [:heavy_check_mark:](img/img_functions_js.go#L448) |
| [IMG_LoadWEBPAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadWEBPAnimation_IO) | [:heavy_check_mark:](img/functions.go#L539) | [:heavy_check_mark:](img/img_functions_js.go#L457) |
| [IMG_SaveAnimation](https://wiki.libsdl.org/SDL3_image/IMG_SaveAnimation) | [:question:]() | [:question:]() |
| [IMG_SaveAnimationTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveAnimationTyped_IO) | [:question:]() | [:question:]() |
| [IMG_SaveANIAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveANIAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_SaveAPNGAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveAPNGAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_SaveAVIFAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIFAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_SaveGIFAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveGIFAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_SaveWEBPAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveWEBPAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_CreateAnimatedCursor](https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimatedCursor) | [:question:]() | [:question:]() |
| [IMG_FreeAnimation](https://wiki.libsdl.org/SDL3_image/IMG_FreeAnimation) | [:heavy_check_mark:](img/methods.go#L6) | [:heavy_check_mark:](img/img_functions_js.go#L440) |
| [IMG_CreateAnimationEncoder](https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimationEncoder) | [:question:]() | [:question:]() |
| [IMG_CreateAnimationEncoder_IO](https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimationEncoder_IO) | [:question:]() | [:question:]() |
| [IMG_CreateAnimationEncoderWithProperties](https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimationEncoderWithProperties) | [:question:]() | [:question:]() |
| [IMG_AddAnimationEncoderFrame](https://wiki.libsdl.org/SDL3_image/IMG_AddAnimationEncoderFrame) | [:question:]() | [:question:]() |
| [IMG_CloseAnimationEncoder](https://wiki.libsdl.org/SDL3_image/IMG_CloseAnimationEncoder) | [:question:]() | [:question:]() |
| [IMG_CreateAnimationDecoder](https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimationDecoder) | [:question:]() | [:question:]() |
| [IMG_CreateAnimationDecoder_IO](https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimationDecoder_IO) | [:question:]() | [:question:]() |
| [IMG_CreateAnimationDecoderWithProperties](https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimationDecoderWithProperties) | [:question:]() | [:question:]() |
| [IMG_GetAnimationDecoderProperties](https://wiki.libsdl.org/SDL3_image/IMG_GetAnimationDecoderProperties) | [:question:]() | [:question:]() |
| [IMG_GetAnimationDecoderFrame](https://wiki.libsdl.org/SDL3_image/IMG_GetAnimationDecoderFrame) | [:question:]() | [:question:]() |
| [IMG_GetAnimationDecoderStatus](https://wiki.libsdl.org/SDL3_image/IMG_GetAnimationDecoderStatus) | [:question:]() | [:question:]() |
| [IMG_ResetAnimationDecoder](https://wiki.libsdl.org/SDL3_image/IMG_ResetAnimationDecoder) | [:question:]() | [:question:]() |
| [IMG_CloseAnimationDecoder](https://wiki.libsdl.org/SDL3_image/IMG_CloseAnimationDecoder) | [:question:]() | [:question:]() |
</details>
</details>
<details open>
<summary><h2>MIXER</h2></summary>
<details open>
<summary><h3>Mixer</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [MIX_Version](https://wiki.libsdl.org/SDL3_mixer/MIX_Version) | [:heavy_check_mark:](mixer/functions.go#L9) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L13) |
| [MIX_Init](https://wiki.libsdl.org/SDL3_mixer/MIX_Init) | [:heavy_check_mark:](mixer/functions.go#L15) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L20) |
| [MIX_Quit](https://wiki.libsdl.org/SDL3_mixer/MIX_Quit) | [:heavy_check_mark:](mixer/functions.go#L25) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L27) |
| [MIX_GetNumAudioDecoders](https://wiki.libsdl.org/SDL3_mixer/MIX_GetNumAudioDecoders) | [:heavy_check_mark:](mixer/functions.go#L31) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L33) |
| [MIX_GetAudioDecoder](https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioDecoder) | [:heavy_check_mark:](mixer/functions.go#L37) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L40) |
| [MIX_CreateMixerDevice](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateMixerDevice) | [:heavy_check_mark:](mixer/functions.go#L43) | [:x:](mixer/mixer_functions_js.go#L49) |
| [MIX_CreateMixer](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateMixer) | [:heavy_check_mark:](mixer/functions.go#L54) | [:x:](mixer/mixer_functions_js.go#L63) |
| [MIX_DestroyMixer](https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyMixer) | [:heavy_check_mark:](mixer/methods.go#L90) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L75) |
| [MIX_GetMixerProperties](https://wiki.libsdl.org/SDL3_mixer/MIX_GetMixerProperties) | [:heavy_check_mark:](mixer/methods.go#L96) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L84) |
| [MIX_GetMixerFormat](https://wiki.libsdl.org/SDL3_mixer/MIX_GetMixerFormat) | [:heavy_check_mark:](mixer/methods.go#L102) | [:x:](mixer/mixer_functions_js.go#L93) |
| [MIX_LockMixer](https://wiki.libsdl.org/SDL3_mixer/MIX_LockMixer) | [:heavy_check_mark:](mixer/methods.go#L112) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L927) |
| [MIX_UnlockMixer](https://wiki.libsdl.org/SDL3_mixer/MIX_UnlockMixer) | [:heavy_check_mark:](mixer/methods.go#L118) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L935) |
| [MIX_LoadAudio_IO](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadAudio_IO) | [:heavy_check_mark:](mixer/methods.go#L124) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L107) |
| [MIX_LoadAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadAudio) | [:heavy_check_mark:](mixer/methods.go#L135) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L122) |
| [MIX_LoadAudioNoCopy](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadAudioNoCopy) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L943) |
| [MIX_LoadAudioWithProperties](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadAudioWithProperties) | [:heavy_check_mark:](mixer/functions.go#L65) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L137) |
| [MIX_LoadRawAudio_IO](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadRawAudio_IO) | [:heavy_check_mark:](mixer/methods.go#L146) | [:x:](mixer/mixer_functions_js.go#L146) |
| [MIX_LoadRawAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadRawAudio) | [:heavy_check_mark:](mixer/methods.go#L157) | [:x:](mixer/mixer_functions_js.go#L164) |
| [MIX_LoadRawAudioNoCopy](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadRawAudioNoCopy) | [:heavy_check_mark:](mixer/methods.go#L169) | [:x:](mixer/mixer_functions_js.go#L182) |
| [MIX_CreateSineWaveAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateSineWaveAudio) | [:heavy_check_mark:](mixer/methods.go#L181) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L202) |
| [MIX_GetAudioProperties](https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioProperties) | [:heavy_check_mark:](mixer/methods.go#L347) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L217) |
| [MIX_GetAudioDuration](https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioDuration) | [:heavy_check_mark:](mixer/methods.go#L358) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L226) |
| [MIX_GetAudioFormat](https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioFormat) | [:heavy_check_mark:](mixer/methods.go#L364) | [:x:](mixer/mixer_functions_js.go#L235) |
| [MIX_DestroyAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyAudio) | [:heavy_check_mark:](mixer/methods.go#L374) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L249) |
| [MIX_CreateTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateTrack) | [:heavy_check_mark:](mixer/methods.go#L192) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L258) |
| [MIX_DestroyTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyTrack) | [:heavy_check_mark:](mixer/methods.go#L394) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L267) |
| [MIX_GetTrackProperties](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackProperties) | [:heavy_check_mark:](mixer/methods.go#L400) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L276) |
| [MIX_GetTrackMixer](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackMixer) | [:heavy_check_mark:](mixer/methods.go#L411) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L285) |
| [MIX_SetTrackAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackAudio) | [:heavy_check_mark:](mixer/methods.go#L422) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L294) |
| [MIX_SetTrackAudioStream](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackAudioStream) | [:heavy_check_mark:](mixer/methods.go#L432) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L305) |
| [MIX_SetTrackIOStream](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackIOStream) | [:heavy_check_mark:](mixer/methods.go#L442) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L316) |
| [MIX_SetTrackRawIOStream](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackRawIOStream) | [:heavy_check_mark:](mixer/methods.go#L452) | [:x:](mixer/mixer_functions_js.go#L959) |
| [MIX_TagTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_TagTrack) | [:heavy_check_mark:](mixer/methods.go#L462) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L329) |
| [MIX_UntagTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_UntagTrack) | [:heavy_check_mark:](mixer/methods.go#L472) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L342) |
| [MIX_GetTrackTags](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackTags) | [:heavy_check_mark:](mixer/methods.go#L478) | [:x:](mixer/mixer_functions_js.go#L977) |
| [MIX_GetTaggedTracks](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTaggedTracks) | [:heavy_check_mark:](mixer/methods.go#L326) | [:x:](mixer/mixer_functions_js.go#L991) |
| [MIX_SetTrackPlaybackPosition](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackPlaybackPosition) | [:heavy_check_mark:](mixer/methods.go#L499) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L354) |
| [MIX_GetTrackPlaybackPosition](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackPlaybackPosition) | [:heavy_check_mark:](mixer/methods.go#L509) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L365) |
| [MIX_GetTrackFadeFrames](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackFadeFrames) | [:heavy_check_mark:](mixer/methods.go#L520) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L1007) |
| [MIX_GetTrackLoops](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackLoops) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1016) |
| [MIX_SetTrackLoops](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackLoops) | [:heavy_check_mark:](mixer/methods.go#L526) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L1025) |
| [MIX_GetTrackAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackAudio) | [:heavy_check_mark:](mixer/methods.go#L536) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L374) |
| [MIX_GetTrackAudioStream](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackAudioStream) | [:heavy_check_mark:](mixer/methods.go#L542) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L383) |
| [MIX_GetTrackRemaining](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackRemaining) | [:heavy_check_mark:](mixer/methods.go#L548) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L392) |
| [MIX_TrackMSToFrames](https://wiki.libsdl.org/SDL3_mixer/MIX_TrackMSToFrames) | [:heavy_check_mark:](mixer/methods.go#L554) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L401) |
| [MIX_TrackFramesToMS](https://wiki.libsdl.org/SDL3_mixer/MIX_TrackFramesToMS) | [:heavy_check_mark:](mixer/methods.go#L560) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L412) |
| [MIX_AudioMSToFrames](https://wiki.libsdl.org/SDL3_mixer/MIX_AudioMSToFrames) | [:heavy_check_mark:](mixer/methods.go#L380) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L423) |
| [MIX_AudioFramesToMS](https://wiki.libsdl.org/SDL3_mixer/MIX_AudioFramesToMS) | [:heavy_check_mark:](mixer/methods.go#L386) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L434) |
| [MIX_MSToFrames](https://wiki.libsdl.org/SDL3_mixer/MIX_MSToFrames) | [:heavy_check_mark:](mixer/functions.go#L76) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L445) |
| [MIX_FramesToMS](https://wiki.libsdl.org/SDL3_mixer/MIX_FramesToMS) | [:heavy_check_mark:](mixer/functions.go#L82) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L456) |
| [MIX_PlayTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_PlayTrack) | [:heavy_check_mark:](mixer/methods.go#L566) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L467) |
| [MIX_PlayTag](https://wiki.libsdl.org/SDL3_mixer/MIX_PlayTag) | [:heavy_check_mark:](mixer/methods.go#L203) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L478) |
| [MIX_PlayAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_PlayAudio) | [:heavy_check_mark:](mixer/methods.go#L213) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L493) |
| [MIX_StopTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_StopTrack) | [:heavy_check_mark:](mixer/methods.go#L576) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L504) |
| [MIX_StopAllTracks](https://wiki.libsdl.org/SDL3_mixer/MIX_StopAllTracks) | [:heavy_check_mark:](mixer/methods.go#L223) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L515) |
| [MIX_StopTag](https://wiki.libsdl.org/SDL3_mixer/MIX_StopTag) | [:heavy_check_mark:](mixer/methods.go#L233) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L526) |
| [MIX_PauseTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_PauseTrack) | [:heavy_check_mark:](mixer/methods.go#L586) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L541) |
| [MIX_PauseAllTracks](https://wiki.libsdl.org/SDL3_mixer/MIX_PauseAllTracks) | [:heavy_check_mark:](mixer/methods.go#L243) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L550) |
| [MIX_PauseTag](https://wiki.libsdl.org/SDL3_mixer/MIX_PauseTag) | [:heavy_check_mark:](mixer/methods.go#L253) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L559) |
| [MIX_ResumeTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_ResumeTrack) | [:heavy_check_mark:](mixer/methods.go#L596) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L572) |
| [MIX_ResumeAllTracks](https://wiki.libsdl.org/SDL3_mixer/MIX_ResumeAllTracks) | [:heavy_check_mark:](mixer/methods.go#L263) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L581) |
| [MIX_ResumeTag](https://wiki.libsdl.org/SDL3_mixer/MIX_ResumeTag) | [:heavy_check_mark:](mixer/methods.go#L273) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L590) |
| [MIX_TrackPlaying](https://wiki.libsdl.org/SDL3_mixer/MIX_TrackPlaying) | [:heavy_check_mark:](mixer/methods.go#L606) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L603) |
| [MIX_TrackPaused](https://wiki.libsdl.org/SDL3_mixer/MIX_TrackPaused) | [:heavy_check_mark:](mixer/methods.go#L612) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L612) |
| [MIX_SetMixerGain](https://wiki.libsdl.org/SDL3_mixer/MIX_SetMixerGain) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1036) |
| [MIX_GetMixerGain](https://wiki.libsdl.org/SDL3_mixer/MIX_GetMixerGain) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1047) |
| [MIX_SetTrackGain](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackGain) | [:heavy_check_mark:](mixer/methods.go#L618) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L621) |
| [MIX_GetTrackGain](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackGain) | [:heavy_check_mark:](mixer/methods.go#L628) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L632) |
| [MIX_SetTagGain](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTagGain) | [:heavy_check_mark:](mixer/methods.go#L283) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L641) |
| [MIX_SetMixerFrequencyRatio](https://wiki.libsdl.org/SDL3_mixer/MIX_SetMixerFrequencyRatio) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1056) |
| [MIX_GetMixerFrequencyRatio](https://wiki.libsdl.org/SDL3_mixer/MIX_GetMixerFrequencyRatio) | [:question:]() | [:question:](mixer/mixer_functions_js.go#L1067) |
| [MIX_SetTrackFrequencyRatio](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackFrequencyRatio) | [:heavy_check_mark:](mixer/methods.go#L634) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L656) |
| [MIX_GetTrackFrequencyRatio](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackFrequencyRatio) | [:heavy_check_mark:](mixer/methods.go#L644) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L667) |
| [MIX_SetTrackOutputChannelMap](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackOutputChannelMap) | [:heavy_check_mark:](mixer/methods.go#L650) | [:x:](mixer/mixer_functions_js.go#L676) |
| [MIX_SetTrackStereo](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackStereo) | [:heavy_check_mark:](mixer/methods.go#L661) | [:x:](mixer/mixer_functions_js.go#L692) |
| [MIX_SetTrack3DPosition](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrack3DPosition) | [:heavy_check_mark:](mixer/methods.go#L672) | [:x:](mixer/mixer_functions_js.go#L706) |
| [MIX_GetTrack3DPosition](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrack3DPosition) | [:heavy_check_mark:](mixer/methods.go#L682) | [:x:](mixer/mixer_functions_js.go#L720) |
| [MIX_CreateGroup](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateGroup) | [:heavy_check_mark:](mixer/methods.go#L293) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L734) |
| [MIX_DestroyGroup](https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyGroup) | [:heavy_check_mark:](mixer/methods.go#L14) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L743) |
| [MIX_GetGroupProperties](https://wiki.libsdl.org/SDL3_mixer/MIX_GetGroupProperties) | [:heavy_check_mark:](mixer/methods.go#L20) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L752) |
| [MIX_GetGroupMixer](https://wiki.libsdl.org/SDL3_mixer/MIX_GetGroupMixer) | [:heavy_check_mark:](mixer/methods.go#L31) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L761) |
| [MIX_SetTrackGroup](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackGroup) | [:heavy_check_mark:](mixer/methods.go#L694) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L770) |
| [MIX_SetTrackStoppedCallback](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackStoppedCallback) | [:heavy_check_mark:](mixer/methods.go#L704) | [:x:](mixer/mixer_functions_js.go#L781) |
| [MIX_SetTrackRawCallback](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackRawCallback) | [:heavy_check_mark:](mixer/methods.go#L714) | [:x:](mixer/mixer_functions_js.go#L795) |
| [MIX_SetTrackCookedCallback](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackCookedCallback) | [:heavy_check_mark:](mixer/methods.go#L724) | [:x:](mixer/mixer_functions_js.go#L809) |
| [MIX_SetGroupPostMixCallback](https://wiki.libsdl.org/SDL3_mixer/MIX_SetGroupPostMixCallback) | [:heavy_check_mark:](mixer/methods.go#L37) | [:x:](mixer/mixer_functions_js.go#L823) |
| [MIX_SetPostMixCallback](https://wiki.libsdl.org/SDL3_mixer/MIX_SetPostMixCallback) | [:heavy_check_mark:](mixer/methods.go#L304) | [:x:](mixer/mixer_functions_js.go#L837) |
| [MIX_Generate](https://wiki.libsdl.org/SDL3_mixer/MIX_Generate) | [:heavy_check_mark:](mixer/methods.go#L314) | [:x:](mixer/mixer_functions_js.go#L1076) |
| [MIX_CreateAudioDecoder](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateAudioDecoder) | [:heavy_check_mark:](mixer/functions.go#L88) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L851) |
| [MIX_CreateAudioDecoder_IO](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateAudioDecoder_IO) | [:heavy_check_mark:](mixer/functions.go#L99) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L864) |
| [MIX_DestroyAudioDecoder](https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyAudioDecoder) | [:heavy_check_mark:](mixer/methods.go#L49) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L877) |
| [MIX_GetAudioDecoderProperties](https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioDecoderProperties) | [:heavy_check_mark:](mixer/methods.go#L55) | [:heavy_check_mark:](mixer/mixer_functions_js.go#L886) |
| [MIX_GetAudioDecoderFormat](https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioDecoderFormat) | [:heavy_check_mark:](mixer/methods.go#L66) | [:x:](mixer/mixer_functions_js.go#L895) |
| [MIX_DecodeAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_DecodeAudio) | [:heavy_check_mark:](mixer/methods.go#L76) | [:x:](mixer/mixer_functions_js.go#L909) |
</details>
</details>
<details open>
<summary><h2>MIDI</h2></summary>
<details open>
<summary><h3>NativeMIDI</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| NativeMidi_Init | [:heavy_check_mark:](midi/functions.go#L29) | [:question:]() |
| NativeMidi_Quit | [:heavy_check_mark:](midi/functions.go#L38) | [:question:]() |
| NativeMidi_LoadSong_IO | [:heavy_check_mark:](midi/functions.go#L53) | [:question:]() |
| NativeMidi_LoadSong | [:heavy_check_mark:](midi/functions.go#L43) | [:question:]() |
| NativeMidi_DestroySong | [:heavy_check_mark:](midi/methods.go#L3) | [:question:]() |
| NativeMidi_Start | [:heavy_check_mark:](midi/methods.go#L11) | [:question:]() |
| NativeMidi_Pause | [:heavy_check_mark:](midi/functions.go#L63) | [:question:]() |
| NativeMidi_Resume | [:heavy_check_mark:](midi/functions.go#L68) | [:question:]() |
| NativeMidi_Stop | [:heavy_check_mark:](midi/functions.go#L73) | [:question:]() |
| NativeMidi_Active | [:heavy_check_mark:](midi/functions.go#L78) | [:question:]() |
| NativeMidi_SetVolume | [:heavy_check_mark:](midi/functions.go#L83) | [:question:]() |
</details>
</details>

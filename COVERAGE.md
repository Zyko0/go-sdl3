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
| [SDL_Init](https://wiki.libsdl.org/SDL3/SDL_Init) | [:heavy_check_mark:](sdl/functions.go#L14) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L13677) |
| [SDL_InitSubSystem](https://wiki.libsdl.org/SDL3/SDL_InitSubSystem) | [:heavy_check_mark:](sdl/functions.go#L24) | [:x:](sdl/sdl_functions_js.go#L13687) |
| [SDL_QuitSubSystem](https://wiki.libsdl.org/SDL3/SDL_QuitSubSystem) | [:heavy_check_mark:](sdl/functions.go#L34) | [:x:](sdl/sdl_functions_js.go#L13700) |
| [SDL_WasInit](https://wiki.libsdl.org/SDL3/SDL_WasInit) | [:heavy_check_mark:](sdl/functions.go#L40) | [:x:](sdl/sdl_functions_js.go#L13711) |
| [SDL_Quit](https://wiki.libsdl.org/SDL3/SDL_Quit) | [:heavy_check_mark:](sdl/functions.go#L46) | [:question:]() |
| [SDL_IsMainThread](https://wiki.libsdl.org/SDL3/SDL_IsMainThread) | [:heavy_check_mark:](sdl/functions.go#L52) | [:x:](sdl/sdl_functions_js.go#L13728) |
| [SDL_RunOnMainThread](https://wiki.libsdl.org/SDL3/SDL_RunOnMainThread) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13739) |
| [SDL_SetAppMetadata](https://wiki.libsdl.org/SDL3/SDL_SetAppMetadata) | [:heavy_check_mark:](sdl/functions.go#L58) | [:x:](sdl/sdl_functions_js.go#L13756) |
| [SDL_SetAppMetadataProperty](https://wiki.libsdl.org/SDL3/SDL_SetAppMetadataProperty) | [:heavy_check_mark:](sdl/functions.go#L68) | [:x:](sdl/sdl_functions_js.go#L13773) |
| [SDL_GetAppMetadataProperty](https://wiki.libsdl.org/SDL3/SDL_GetAppMetadataProperty) | [:heavy_check_mark:](sdl/functions.go#L78) | [:x:](sdl/sdl_functions_js.go#L13788) |
</details>
<details open>
<summary><h3>Hints</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetHintWithPriority](https://wiki.libsdl.org/SDL3/SDL_SetHintWithPriority) | [:heavy_check_mark:](sdl/functions.go#L86) | [:x:](sdl/sdl_functions_js.go#L13563) |
| [SDL_SetHint](https://wiki.libsdl.org/SDL3/SDL_SetHint) | [:heavy_check_mark:](sdl/functions.go#L96) | [:x:](sdl/sdl_functions_js.go#L13580) |
| [SDL_ResetHint](https://wiki.libsdl.org/SDL3/SDL_ResetHint) | [:heavy_check_mark:](sdl/functions.go#L106) | [:x:](sdl/sdl_functions_js.go#L13595) |
| [SDL_ResetHints](https://wiki.libsdl.org/SDL3/SDL_ResetHints) | [:heavy_check_mark:](sdl/functions.go#L116) | [:x:](sdl/sdl_functions_js.go#L13608) |
| [SDL_GetHint](https://wiki.libsdl.org/SDL3/SDL_GetHint) | [:heavy_check_mark:](sdl/functions.go#L122) | [:x:](sdl/sdl_functions_js.go#L13617) |
| [SDL_GetHintBoolean](https://wiki.libsdl.org/SDL3/SDL_GetHintBoolean) | [:heavy_check_mark:](sdl/functions.go#L128) | [:x:](sdl/sdl_functions_js.go#L13630) |
| [SDL_AddHintCallback](https://wiki.libsdl.org/SDL3/SDL_AddHintCallback) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13645) |
| [SDL_RemoveHintCallback](https://wiki.libsdl.org/SDL3/SDL_RemoveHintCallback) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13662) |
</details>
<details>
<summary><h3>Error</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetError](https://wiki.libsdl.org/SDL3/SDL_SetError) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L496) |
| [SDL_SetErrorV](https://wiki.libsdl.org/SDL3/SDL_SetErrorV) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L509) |
| [SDL_OutOfMemory](https://wiki.libsdl.org/SDL3/SDL_OutOfMemory) | [:heavy_check_mark:](sdl/functions.go#L139) | [:x:](sdl/sdl_functions_js.go#L524) |
| [SDL_GetError](https://wiki.libsdl.org/SDL3/SDL_GetError) | [:heavy_check_mark:](sdl/init_notjs.go#L32) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L535) |
| [SDL_ClearError](https://wiki.libsdl.org/SDL3/SDL_ClearError) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L543) |
</details>
<details>
<summary><h3>Version</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetVersion](https://wiki.libsdl.org/SDL3/SDL_GetVersion) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L24) |
| [SDL_GetRevision](https://wiki.libsdl.org/SDL3/SDL_GetRevision) | [:question:]() | [:question:]() |
</details>
<details open>
<summary><h3>Properties</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetGlobalProperties](https://wiki.libsdl.org/SDL3/SDL_GetGlobalProperties) | [:heavy_check_mark:](sdl/functions.go#L147) | [:x:](sdl/sdl_functions_js.go#L554) |
| [SDL_CreateProperties](https://wiki.libsdl.org/SDL3/SDL_CreateProperties) | [:heavy_check_mark:](sdl/functions.go#L158) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L565) |
| [SDL_CopyProperties](https://wiki.libsdl.org/SDL3/SDL_CopyProperties) | [:heavy_check_mark:](sdl/functions.go#L169) | [:x:](sdl/sdl_functions_js.go#L573) |
| [SDL_LockProperties](https://wiki.libsdl.org/SDL3/SDL_LockProperties) | [:heavy_check_mark:](sdl/methods.go#L5779) | [:x:](sdl/sdl_functions_js.go#L588) |
| [SDL_UnlockProperties](https://wiki.libsdl.org/SDL3/SDL_UnlockProperties) | [:heavy_check_mark:](sdl/methods.go#L5789) | [:x:](sdl/sdl_functions_js.go#L601) |
| [SDL_SetPointerPropertyWithCleanup](https://wiki.libsdl.org/SDL3/SDL_SetPointerPropertyWithCleanup) | [:x:](sdl/methods.go#L5795) | [:x:](sdl/sdl_functions_js.go#L612) |
| [SDL_SetPointerProperty](https://wiki.libsdl.org/SDL3/SDL_SetPointerProperty) | [:x:](sdl/methods.go#L5802) | [:x:](sdl/sdl_functions_js.go#L633) |
| [SDL_SetStringProperty](https://wiki.libsdl.org/SDL3/SDL_SetStringProperty) | [:heavy_check_mark:](sdl/methods.go#L5809) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L650) |
| [SDL_SetNumberProperty](https://wiki.libsdl.org/SDL3/SDL_SetNumberProperty) | [:heavy_check_mark:](sdl/methods.go#L5819) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L666) |
| [SDL_SetFloatProperty](https://wiki.libsdl.org/SDL3/SDL_SetFloatProperty) | [:heavy_check_mark:](sdl/methods.go#L5829) | [:x:](sdl/sdl_functions_js.go#L682) |
| [SDL_SetBooleanProperty](https://wiki.libsdl.org/SDL3/SDL_SetBooleanProperty) | [:heavy_check_mark:](sdl/methods.go#L5839) | [:x:](sdl/sdl_functions_js.go#L699) |
| [SDL_HasProperty](https://wiki.libsdl.org/SDL3/SDL_HasProperty) | [:heavy_check_mark:](sdl/methods.go#L5849) | [:x:](sdl/sdl_functions_js.go#L716) |
| [SDL_GetPropertyType](https://wiki.libsdl.org/SDL3/SDL_GetPropertyType) | [:heavy_check_mark:](sdl/methods.go#L5855) | [:x:](sdl/sdl_functions_js.go#L731) |
| [SDL_GetPointerProperty](https://wiki.libsdl.org/SDL3/SDL_GetPointerProperty) | [:x:](sdl/methods.go#L5861) | [:x:](sdl/sdl_functions_js.go#L746) |
| [SDL_GetStringProperty](https://wiki.libsdl.org/SDL3/SDL_GetStringProperty) | [:heavy_check_mark:](sdl/methods.go#L5868) | [:x:](sdl/sdl_functions_js.go#L763) |
| [SDL_GetNumberProperty](https://wiki.libsdl.org/SDL3/SDL_GetNumberProperty) | [:heavy_check_mark:](sdl/methods.go#L5874) | [:x:](sdl/sdl_functions_js.go#L780) |
| [SDL_GetFloatProperty](https://wiki.libsdl.org/SDL3/SDL_GetFloatProperty) | [:heavy_check_mark:](sdl/methods.go#L5880) | [:x:](sdl/sdl_functions_js.go#L797) |
| [SDL_GetBooleanProperty](https://wiki.libsdl.org/SDL3/SDL_GetBooleanProperty) | [:heavy_check_mark:](sdl/methods.go#L5886) | [:x:](sdl/sdl_functions_js.go#L814) |
| [SDL_ClearProperty](https://wiki.libsdl.org/SDL3/SDL_ClearProperty) | [:heavy_check_mark:](sdl/methods.go#L5892) | [:x:](sdl/sdl_functions_js.go#L831) |
| [SDL_EnumerateProperties](https://wiki.libsdl.org/SDL3/SDL_EnumerateProperties) | [:heavy_check_mark:](sdl/methods.go#L5902) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L846) |
| [SDL_DestroyProperties](https://wiki.libsdl.org/SDL3/SDL_DestroyProperties) | [:heavy_check_mark:](sdl/methods.go#L5908) | [:x:](sdl/sdl_functions_js.go#L861) |
</details>
<details>
<summary><h3>Log</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetLogPriorities](https://wiki.libsdl.org/SDL3/SDL_SetLogPriorities) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13865) |
| [SDL_SetLogPriority](https://wiki.libsdl.org/SDL3/SDL_SetLogPriority) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13876) |
| [SDL_GetLogPriority](https://wiki.libsdl.org/SDL3/SDL_GetLogPriority) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13889) |
| [SDL_ResetLogPriorities](https://wiki.libsdl.org/SDL3/SDL_ResetLogPriorities) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13902) |
| [SDL_SetLogPriorityPrefix](https://wiki.libsdl.org/SDL3/SDL_SetLogPriorityPrefix) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13911) |
| [SDL_Log](https://wiki.libsdl.org/SDL3/SDL_Log) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13926) |
| [SDL_LogTrace](https://wiki.libsdl.org/SDL3/SDL_LogTrace) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13937) |
| [SDL_LogVerbose](https://wiki.libsdl.org/SDL3/SDL_LogVerbose) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13950) |
| [SDL_LogDebug](https://wiki.libsdl.org/SDL3/SDL_LogDebug) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13963) |
| [SDL_LogInfo](https://wiki.libsdl.org/SDL3/SDL_LogInfo) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13976) |
| [SDL_LogWarn](https://wiki.libsdl.org/SDL3/SDL_LogWarn) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13989) |
| [SDL_LogError](https://wiki.libsdl.org/SDL3/SDL_LogError) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14002) |
| [SDL_LogCritical](https://wiki.libsdl.org/SDL3/SDL_LogCritical) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14015) |
| [SDL_LogMessage](https://wiki.libsdl.org/SDL3/SDL_LogMessage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14028) |
| [SDL_LogMessageV](https://wiki.libsdl.org/SDL3/SDL_LogMessageV) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14043) |
| [SDL_GetDefaultLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_GetDefaultLogOutputFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14060) |
| [SDL_GetLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_GetLogOutputFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14071) |
| [SDL_SetLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_SetLogOutputFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14090) |
</details>
<details open>
<summary><h3>Video</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumVideoDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumVideoDrivers) | [:heavy_check_mark:](sdl/functions.go#L614) | [:x:](sdl/sdl_functions_js.go#L5604) |
| [SDL_GetVideoDriver](https://wiki.libsdl.org/SDL3/SDL_GetVideoDriver) | [:heavy_check_mark:](sdl/functions.go#L620) | [:x:](sdl/sdl_functions_js.go#L5615) |
| [SDL_GetCurrentVideoDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentVideoDriver) | [:heavy_check_mark:](sdl/functions.go#L626) | [:x:](sdl/sdl_functions_js.go#L5628) |
| [SDL_GetSystemTheme](https://wiki.libsdl.org/SDL3/SDL_GetSystemTheme) | [:heavy_check_mark:](sdl/functions.go#L632) | [:x:](sdl/sdl_functions_js.go#L5639) |
| [SDL_GetDisplays](https://wiki.libsdl.org/SDL3/SDL_GetDisplays) | [:heavy_check_mark:](sdl/functions.go#L638) | [:x:](sdl/sdl_functions_js.go#L5650) |
| [SDL_GetPrimaryDisplay](https://wiki.libsdl.org/SDL3/SDL_GetPrimaryDisplay) | [:heavy_check_mark:](sdl/functions.go#L652) | [:x:](sdl/sdl_functions_js.go#L5666) |
| [SDL_GetDisplayProperties](https://wiki.libsdl.org/SDL3/SDL_GetDisplayProperties) | [:heavy_check_mark:](sdl/methods.go#L3696) | [:x:](sdl/sdl_functions_js.go#L5677) |
| [SDL_GetDisplayName](https://wiki.libsdl.org/SDL3/SDL_GetDisplayName) | [:heavy_check_mark:](sdl/methods.go#L3707) | [:x:](sdl/sdl_functions_js.go#L5690) |
| [SDL_GetDisplayBounds](https://wiki.libsdl.org/SDL3/SDL_GetDisplayBounds) | [:heavy_check_mark:](sdl/methods.go#L3718) | [:x:](sdl/sdl_functions_js.go#L5703) |
| [SDL_GetDisplayUsableBounds](https://wiki.libsdl.org/SDL3/SDL_GetDisplayUsableBounds) | [:heavy_check_mark:](sdl/methods.go#L3730) | [:x:](sdl/sdl_functions_js.go#L5721) |
| [SDL_GetNaturalDisplayOrientation](https://wiki.libsdl.org/SDL3/SDL_GetNaturalDisplayOrientation) | [:heavy_check_mark:](sdl/methods.go#L3742) | [:x:](sdl/sdl_functions_js.go#L5739) |
| [SDL_GetCurrentDisplayOrientation](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayOrientation) | [:heavy_check_mark:](sdl/methods.go#L3748) | [:x:](sdl/sdl_functions_js.go#L5752) |
| [SDL_GetDisplayContentScale](https://wiki.libsdl.org/SDL3/SDL_GetDisplayContentScale) | [:heavy_check_mark:](sdl/methods.go#L3754) | [:x:](sdl/sdl_functions_js.go#L5765) |
| [SDL_GetFullscreenDisplayModes](https://wiki.libsdl.org/SDL3/SDL_GetFullscreenDisplayModes) | [:heavy_check_mark:](sdl/methods.go#L3765) | [:x:](sdl/sdl_functions_js.go#L5778) |
| [SDL_GetClosestFullscreenDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetClosestFullscreenDisplayMode) | [:heavy_check_mark:](sdl/methods.go#L3779) | [:x:](sdl/sdl_functions_js.go#L5796) |
| [SDL_GetDesktopDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetDesktopDisplayMode) | [:heavy_check_mark:](sdl/methods.go#L3791) | [:x:](sdl/sdl_functions_js.go#L5822) |
| [SDL_GetCurrentDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayMode) | [:heavy_check_mark:](sdl/methods.go#L3802) | [:x:](sdl/sdl_functions_js.go#L5838) |
| [SDL_GetDisplayForPoint](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForPoint) | [:heavy_check_mark:](sdl/functions.go#L658) | [:x:](sdl/sdl_functions_js.go#L5854) |
| [SDL_GetDisplayForRect](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForRect) | [:heavy_check_mark:](sdl/functions.go#L664) | [:x:](sdl/sdl_functions_js.go#L5870) |
| [SDL_GetDisplayForWindow](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForWindow) | [:heavy_check_mark:](sdl/functions.go#L670) | [:x:](sdl/sdl_functions_js.go#L5886) |
| [SDL_GetWindowPixelDensity](https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelDensity) | [:heavy_check_mark:](sdl/methods.go#L4136) | [:x:](sdl/sdl_functions_js.go#L5902) |
| [SDL_GetWindowDisplayScale](https://wiki.libsdl.org/SDL3/SDL_GetWindowDisplayScale) | [:heavy_check_mark:](sdl/methods.go#L4147) | [:x:](sdl/sdl_functions_js.go#L5918) |
| [SDL_SetWindowFullscreenMode](https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreenMode) | [:heavy_check_mark:](sdl/methods.go#L4158) | [:x:](sdl/sdl_functions_js.go#L5934) |
| [SDL_GetWindowFullscreenMode](https://wiki.libsdl.org/SDL3/SDL_GetWindowFullscreenMode) | [:heavy_check_mark:](sdl/methods.go#L4168) | [:x:](sdl/sdl_functions_js.go#L5955) |
| [SDL_GetWindowICCProfile](https://wiki.libsdl.org/SDL3/SDL_GetWindowICCProfile) | [:heavy_check_mark:](sdl/methods.go#L4174) | [:x:](sdl/sdl_functions_js.go#L5974) |
| [SDL_GetWindowPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelFormat) | [:heavy_check_mark:](sdl/methods.go#L4188) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5995) |
| [SDL_GetWindows](https://wiki.libsdl.org/SDL3/SDL_GetWindows) | [:heavy_check_mark:](sdl/functions.go#L676) | [:x:](sdl/sdl_functions_js.go#L6008) |
| [SDL_CreateWindow](https://wiki.libsdl.org/SDL3/SDL_CreateWindow) | [:heavy_check_mark:](sdl/functions.go#L690) | [:x:](sdl/sdl_functions_js.go#L6024) |
| [SDL_CreatePopupWindow](https://wiki.libsdl.org/SDL3/SDL_CreatePopupWindow) | [:heavy_check_mark:](sdl/methods.go#L4199) | [:x:](sdl/sdl_functions_js.go#L6046) |
| [SDL_CreateWindowWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateWindowWithProperties) | [:heavy_check_mark:](sdl/functions.go#L701) | [:x:](sdl/sdl_functions_js.go#L6075) |
| [SDL_GetWindowID](https://wiki.libsdl.org/SDL3/SDL_GetWindowID) | [:heavy_check_mark:](sdl/methods.go#L4210) | [:x:](sdl/sdl_functions_js.go#L6091) |
| [SDL_GetWindowFromID](https://wiki.libsdl.org/SDL3/SDL_GetWindowFromID) | [:heavy_check_mark:](sdl/methods.go#L32) | [:x:](sdl/sdl_functions_js.go#L6107) |
| [SDL_GetWindowParent](https://wiki.libsdl.org/SDL3/SDL_GetWindowParent) | [:heavy_check_mark:](sdl/methods.go#L4221) | [:x:](sdl/sdl_functions_js.go#L6123) |
| [SDL_GetWindowProperties](https://wiki.libsdl.org/SDL3/SDL_GetWindowProperties) | [:heavy_check_mark:](sdl/methods.go#L4227) | [:x:](sdl/sdl_functions_js.go#L6142) |
| [SDL_GetWindowFlags](https://wiki.libsdl.org/SDL3/SDL_GetWindowFlags) | [:heavy_check_mark:](sdl/methods.go#L4238) | [:x:](sdl/sdl_functions_js.go#L6158) |
| [SDL_SetWindowTitle](https://wiki.libsdl.org/SDL3/SDL_SetWindowTitle) | [:heavy_check_mark:](sdl/methods.go#L4244) | [:x:](sdl/sdl_functions_js.go#L6174) |
| [SDL_GetWindowTitle](https://wiki.libsdl.org/SDL3/SDL_GetWindowTitle) | [:heavy_check_mark:](sdl/methods.go#L4254) | [:x:](sdl/sdl_functions_js.go#L6192) |
| [SDL_SetWindowIcon](https://wiki.libsdl.org/SDL3/SDL_SetWindowIcon) | [:heavy_check_mark:](sdl/methods.go#L4260) | [:x:](sdl/sdl_functions_js.go#L6208) |
| [SDL_SetWindowPosition](https://wiki.libsdl.org/SDL3/SDL_SetWindowPosition) | [:heavy_check_mark:](sdl/methods.go#L4270) | [:x:](sdl/sdl_functions_js.go#L6229) |
| [SDL_GetWindowPosition](https://wiki.libsdl.org/SDL3/SDL_GetWindowPosition) | [:heavy_check_mark:](sdl/methods.go#L4280) | [:x:](sdl/sdl_functions_js.go#L6249) |
| [SDL_SetWindowSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowSize) | [:heavy_check_mark:](sdl/methods.go#L4292) | [:x:](sdl/sdl_functions_js.go#L6275) |
| [SDL_GetWindowSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowSize) | [:heavy_check_mark:](sdl/methods.go#L4302) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6295) |
| [SDL_GetWindowSafeArea](https://wiki.libsdl.org/SDL3/SDL_GetWindowSafeArea) | [:heavy_check_mark:](sdl/methods.go#L4314) | [:x:](sdl/sdl_functions_js.go#L6317) |
| [SDL_SetWindowAspectRatio](https://wiki.libsdl.org/SDL3/SDL_SetWindowAspectRatio) | [:heavy_check_mark:](sdl/methods.go#L4326) | [:x:](sdl/sdl_functions_js.go#L6338) |
| [SDL_GetWindowAspectRatio](https://wiki.libsdl.org/SDL3/SDL_GetWindowAspectRatio) | [:heavy_check_mark:](sdl/methods.go#L4336) | [:x:](sdl/sdl_functions_js.go#L6358) |
| [SDL_GetWindowBordersSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowBordersSize) | [:heavy_check_mark:](sdl/methods.go#L4348) | [:x:](sdl/sdl_functions_js.go#L6384) |
| [SDL_GetWindowSizeInPixels](https://wiki.libsdl.org/SDL3/SDL_GetWindowSizeInPixels) | [:heavy_check_mark:](sdl/methods.go#L4360) | [:x:](sdl/sdl_functions_js.go#L6420) |
| [SDL_SetWindowMinimumSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowMinimumSize) | [:heavy_check_mark:](sdl/methods.go#L4372) | [:x:](sdl/sdl_functions_js.go#L6446) |
| [SDL_GetWindowMinimumSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowMinimumSize) | [:heavy_check_mark:](sdl/methods.go#L4382) | [:x:](sdl/sdl_functions_js.go#L6466) |
| [SDL_SetWindowMaximumSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowMaximumSize) | [:heavy_check_mark:](sdl/methods.go#L4394) | [:x:](sdl/sdl_functions_js.go#L6492) |
| [SDL_GetWindowMaximumSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowMaximumSize) | [:heavy_check_mark:](sdl/methods.go#L4404) | [:x:](sdl/sdl_functions_js.go#L6512) |
| [SDL_SetWindowBordered](https://wiki.libsdl.org/SDL3/SDL_SetWindowBordered) | [:heavy_check_mark:](sdl/methods.go#L4416) | [:x:](sdl/sdl_functions_js.go#L6538) |
| [SDL_SetWindowResizable](https://wiki.libsdl.org/SDL3/SDL_SetWindowResizable) | [:heavy_check_mark:](sdl/methods.go#L4426) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6556) |
| [SDL_SetWindowAlwaysOnTop](https://wiki.libsdl.org/SDL3/SDL_SetWindowAlwaysOnTop) | [:heavy_check_mark:](sdl/methods.go#L4436) | [:x:](sdl/sdl_functions_js.go#L6571) |
| [SDL_SetWindowFillDocument](https://wiki.libsdl.org/SDL3/SDL_SetWindowFillDocument) | [:heavy_check_mark:](sdl/methods.go#L4446) | [:question:]() |
| [SDL_ShowWindow](https://wiki.libsdl.org/SDL3/SDL_ShowWindow) | [:heavy_check_mark:](sdl/methods.go#L4456) | [:x:](sdl/sdl_functions_js.go#L6589) |
| [SDL_HideWindow](https://wiki.libsdl.org/SDL3/SDL_HideWindow) | [:heavy_check_mark:](sdl/methods.go#L4466) | [:x:](sdl/sdl_functions_js.go#L6605) |
| [SDL_RaiseWindow](https://wiki.libsdl.org/SDL3/SDL_RaiseWindow) | [:heavy_check_mark:](sdl/methods.go#L4476) | [:x:](sdl/sdl_functions_js.go#L6621) |
| [SDL_MaximizeWindow](https://wiki.libsdl.org/SDL3/SDL_MaximizeWindow) | [:heavy_check_mark:](sdl/methods.go#L4486) | [:x:](sdl/sdl_functions_js.go#L6637) |
| [SDL_MinimizeWindow](https://wiki.libsdl.org/SDL3/SDL_MinimizeWindow) | [:heavy_check_mark:](sdl/methods.go#L4496) | [:x:](sdl/sdl_functions_js.go#L6653) |
| [SDL_RestoreWindow](https://wiki.libsdl.org/SDL3/SDL_RestoreWindow) | [:heavy_check_mark:](sdl/methods.go#L4506) | [:x:](sdl/sdl_functions_js.go#L6669) |
| [SDL_SetWindowFullscreen](https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreen) | [:heavy_check_mark:](sdl/methods.go#L4516) | [:x:](sdl/sdl_functions_js.go#L6685) |
| [SDL_SyncWindow](https://wiki.libsdl.org/SDL3/SDL_SyncWindow) | [:heavy_check_mark:](sdl/methods.go#L4526) | [:x:](sdl/sdl_functions_js.go#L6703) |
| [SDL_WindowHasSurface](https://wiki.libsdl.org/SDL3/SDL_WindowHasSurface) | [:heavy_check_mark:](sdl/methods.go#L4536) | [:x:](sdl/sdl_functions_js.go#L6719) |
| [SDL_GetWindowSurface](https://wiki.libsdl.org/SDL3/SDL_GetWindowSurface) | [:heavy_check_mark:](sdl/methods.go#L4542) | [:x:](sdl/sdl_functions_js.go#L6735) |
| [SDL_SetWindowSurfaceVSync](https://wiki.libsdl.org/SDL3/SDL_SetWindowSurfaceVSync) | [:heavy_check_mark:](sdl/methods.go#L4553) | [:x:](sdl/sdl_functions_js.go#L6754) |
| [SDL_GetWindowSurfaceVSync](https://wiki.libsdl.org/SDL3/SDL_GetWindowSurfaceVSync) | [:heavy_check_mark:](sdl/methods.go#L4563) | [:x:](sdl/sdl_functions_js.go#L6772) |
| [SDL_UpdateWindowSurface](https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurface) | [:heavy_check_mark:](sdl/methods.go#L4575) | [:x:](sdl/sdl_functions_js.go#L6793) |
| [SDL_UpdateWindowSurfaceRects](https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurfaceRects) | [:heavy_check_mark:](sdl/methods.go#L4585) | [:x:](sdl/sdl_functions_js.go#L6809) |
| [SDL_DestroyWindowSurface](https://wiki.libsdl.org/SDL3/SDL_DestroyWindowSurface) | [:heavy_check_mark:](sdl/methods.go#L4595) | [:x:](sdl/sdl_functions_js.go#L6832) |
| [SDL_SetWindowKeyboardGrab](https://wiki.libsdl.org/SDL3/SDL_SetWindowKeyboardGrab) | [:heavy_check_mark:](sdl/methods.go#L4605) | [:x:](sdl/sdl_functions_js.go#L6848) |
| [SDL_SetWindowMouseGrab](https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseGrab) | [:heavy_check_mark:](sdl/methods.go#L4615) | [:x:](sdl/sdl_functions_js.go#L6866) |
| [SDL_GetWindowKeyboardGrab](https://wiki.libsdl.org/SDL3/SDL_GetWindowKeyboardGrab) | [:heavy_check_mark:](sdl/methods.go#L4625) | [:x:](sdl/sdl_functions_js.go#L6884) |
| [SDL_GetWindowMouseGrab](https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseGrab) | [:heavy_check_mark:](sdl/methods.go#L4631) | [:x:](sdl/sdl_functions_js.go#L6900) |
| [SDL_GetGrabbedWindow](https://wiki.libsdl.org/SDL3/SDL_GetGrabbedWindow) | [:heavy_check_mark:](sdl/functions.go#L712) | [:x:](sdl/sdl_functions_js.go#L6916) |
| [SDL_SetWindowMouseRect](https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseRect) | [:heavy_check_mark:](sdl/methods.go#L4637) | [:x:](sdl/sdl_functions_js.go#L6930) |
| [SDL_GetWindowMouseRect](https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseRect) | [:heavy_check_mark:](sdl/methods.go#L4647) | [:x:](sdl/sdl_functions_js.go#L6951) |
| [SDL_SetWindowOpacity](https://wiki.libsdl.org/SDL3/SDL_SetWindowOpacity) | [:heavy_check_mark:](sdl/methods.go#L4653) | [:x:](sdl/sdl_functions_js.go#L6970) |
| [SDL_GetWindowOpacity](https://wiki.libsdl.org/SDL3/SDL_GetWindowOpacity) | [:heavy_check_mark:](sdl/methods.go#L4663) | [:x:](sdl/sdl_functions_js.go#L6988) |
| [SDL_SetWindowParent](https://wiki.libsdl.org/SDL3/SDL_SetWindowParent) | [:heavy_check_mark:](sdl/methods.go#L4669) | [:x:](sdl/sdl_functions_js.go#L7004) |
| [SDL_SetWindowModal](https://wiki.libsdl.org/SDL3/SDL_SetWindowModal) | [:heavy_check_mark:](sdl/methods.go#L4679) | [:x:](sdl/sdl_functions_js.go#L7025) |
| [SDL_SetWindowFocusable](https://wiki.libsdl.org/SDL3/SDL_SetWindowFocusable) | [:heavy_check_mark:](sdl/methods.go#L4689) | [:x:](sdl/sdl_functions_js.go#L7043) |
| [SDL_ShowWindowSystemMenu](https://wiki.libsdl.org/SDL3/SDL_ShowWindowSystemMenu) | [:heavy_check_mark:](sdl/methods.go#L4699) | [:x:](sdl/sdl_functions_js.go#L7061) |
| [SDL_SetWindowHitTest](https://wiki.libsdl.org/SDL3/SDL_SetWindowHitTest) | [:heavy_check_mark:](sdl/methods.go#L4709) | [:x:](sdl/sdl_functions_js.go#L7081) |
| [SDL_SetWindowShape](https://wiki.libsdl.org/SDL3/SDL_SetWindowShape) | [:heavy_check_mark:](sdl/methods.go#L4719) | [:x:](sdl/sdl_functions_js.go#L7101) |
| [SDL_FlashWindow](https://wiki.libsdl.org/SDL3/SDL_FlashWindow) | [:heavy_check_mark:](sdl/methods.go#L4729) | [:x:](sdl/sdl_functions_js.go#L7122) |
| [SDL_SetWindowProgressState](https://wiki.libsdl.org/SDL3/SDL_SetWindowProgressState) | [:heavy_check_mark:](sdl/methods.go#L4739) | [:question:]() |
| [SDL_GetWindowProgressState](https://wiki.libsdl.org/SDL3/SDL_GetWindowProgressState) | [:heavy_check_mark:](sdl/methods.go#L4749) | [:question:]() |
| [SDL_SetWindowProgressValue](https://wiki.libsdl.org/SDL3/SDL_SetWindowProgressValue) | [:heavy_check_mark:](sdl/methods.go#L4760) | [:question:]() |
| [SDL_GetWindowProgressValue](https://wiki.libsdl.org/SDL3/SDL_GetWindowProgressValue) | [:heavy_check_mark:](sdl/methods.go#L4770) | [:question:]() |
| [SDL_DestroyWindow](https://wiki.libsdl.org/SDL3/SDL_DestroyWindow) | [:heavy_check_mark:](sdl/methods.go#L4781) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7140) |
| [SDL_ScreenSaverEnabled](https://wiki.libsdl.org/SDL3/SDL_ScreenSaverEnabled) | [:heavy_check_mark:](sdl/functions.go#L718) | [:x:](sdl/sdl_functions_js.go#L7152) |
| [SDL_EnableScreenSaver](https://wiki.libsdl.org/SDL3/SDL_EnableScreenSaver) | [:heavy_check_mark:](sdl/functions.go#L724) | [:x:](sdl/sdl_functions_js.go#L7163) |
| [SDL_DisableScreenSaver](https://wiki.libsdl.org/SDL3/SDL_DisableScreenSaver) | [:heavy_check_mark:](sdl/functions.go#L734) | [:x:](sdl/sdl_functions_js.go#L7174) |
| [SDL_GL_LoadLibrary](https://wiki.libsdl.org/SDL3/SDL_GL_LoadLibrary) | [:heavy_check_mark:](sdl/functions.go#L744) | [:x:](sdl/sdl_functions_js.go#L7185) |
| [SDL_GL_GetProcAddress](https://wiki.libsdl.org/SDL3/SDL_GL_GetProcAddress) | [:heavy_check_mark:](sdl/functions.go#L754) | [:x:](sdl/sdl_functions_js.go#L7198) |
| [SDL_EGL_GetProcAddress](https://wiki.libsdl.org/SDL3/SDL_EGL_GetProcAddress) | [:heavy_check_mark:](sdl/functions.go#L760) | [:x:](sdl/sdl_functions_js.go#L7211) |
| [SDL_GL_UnloadLibrary](https://wiki.libsdl.org/SDL3/SDL_GL_UnloadLibrary) | [:heavy_check_mark:](sdl/functions.go#L772) | [:x:](sdl/sdl_functions_js.go#L7224) |
| [SDL_GL_ExtensionSupported](https://wiki.libsdl.org/SDL3/SDL_GL_ExtensionSupported) | [:heavy_check_mark:](sdl/functions.go#L766) | [:x:](sdl/sdl_functions_js.go#L7233) |
| [SDL_GL_ResetAttributes](https://wiki.libsdl.org/SDL3/SDL_GL_ResetAttributes) | [:heavy_check_mark:](sdl/functions.go#L821) | [:x:](sdl/sdl_functions_js.go#L7246) |
| [SDL_GL_SetAttribute](https://wiki.libsdl.org/SDL3/SDL_GL_SetAttribute) | [:heavy_check_mark:](sdl/functions.go#L827) | [:x:](sdl/sdl_functions_js.go#L7255) |
| [SDL_GL_GetAttribute](https://wiki.libsdl.org/SDL3/SDL_GL_GetAttribute) | [:heavy_check_mark:](sdl/functions.go#L837) | [:x:](sdl/sdl_functions_js.go#L7270) |
| [SDL_GL_CreateContext](https://wiki.libsdl.org/SDL3/SDL_GL_CreateContext) | [:heavy_check_mark:](sdl/functions.go#L778) | [:x:](sdl/sdl_functions_js.go#L7288) |
| [SDL_GL_MakeCurrent](https://wiki.libsdl.org/SDL3/SDL_GL_MakeCurrent) | [:heavy_check_mark:](sdl/functions.go#L789) | [:x:](sdl/sdl_functions_js.go#L7304) |
| [SDL_GL_GetCurrentWindow](https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentWindow) | [:heavy_check_mark:](sdl/functions.go#L849) | [:x:](sdl/sdl_functions_js.go#L7322) |
| [SDL_GL_GetCurrentContext](https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentContext) | [:heavy_check_mark:](sdl/functions.go#L860) | [:x:](sdl/sdl_functions_js.go#L7336) |
| [SDL_EGL_GetCurrentDisplay](https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentDisplay) | [:heavy_check_mark:](sdl/functions.go#L871) | [:x:](sdl/sdl_functions_js.go#L7347) |
| [SDL_EGL_GetCurrentConfig](https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentConfig) | [:heavy_check_mark:](sdl/functions.go#L882) | [:x:](sdl/sdl_functions_js.go#L7358) |
| [SDL_EGL_GetWindowSurface](https://wiki.libsdl.org/SDL3/SDL_EGL_GetWindowSurface) | [:heavy_check_mark:](sdl/functions.go#L799) | [:x:](sdl/sdl_functions_js.go#L7369) |
| [SDL_EGL_SetAttributeCallbacks](https://wiki.libsdl.org/SDL3/SDL_EGL_SetAttributeCallbacks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7385) |
| [SDL_GL_SetSwapInterval](https://wiki.libsdl.org/SDL3/SDL_GL_SetSwapInterval) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7402) |
| [SDL_GL_GetSwapInterval](https://wiki.libsdl.org/SDL3/SDL_GL_GetSwapInterval) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7415) |
| [SDL_GL_SwapWindow](https://wiki.libsdl.org/SDL3/SDL_GL_SwapWindow) | [:heavy_check_mark:](sdl/functions.go#L805) | [:x:](sdl/sdl_functions_js.go#L7431) |
| [SDL_GL_DestroyContext](https://wiki.libsdl.org/SDL3/SDL_GL_DestroyContext) | [:heavy_check_mark:](sdl/functions.go#L815) | [:x:](sdl/sdl_functions_js.go#L7447) |
</details>
<details open>
<summary><h3>Events</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_PumpEvents](https://wiki.libsdl.org/SDL3/SDL_PumpEvents) | [:heavy_check_mark:](sdl/functions.go#L185) | [:x:](sdl/sdl_functions_js.go#L10707) |
| [SDL_PeepEvents](https://wiki.libsdl.org/SDL3/SDL_PeepEvents) | [:heavy_check_mark:](sdl/functions.go#L213) | [:x:](sdl/sdl_functions_js.go#L10716) |
| [SDL_HasEvent](https://wiki.libsdl.org/SDL3/SDL_HasEvent) | [:heavy_check_mark:](sdl/functions.go#L239) | [:x:](sdl/sdl_functions_js.go#L10740) |
| [SDL_HasEvents](https://wiki.libsdl.org/SDL3/SDL_HasEvents) | [:heavy_check_mark:](sdl/functions.go#L245) | [:x:](sdl/sdl_functions_js.go#L10753) |
| [SDL_FlushEvent](https://wiki.libsdl.org/SDL3/SDL_FlushEvent) | [:heavy_check_mark:](sdl/functions.go#L251) | [:x:](sdl/sdl_functions_js.go#L10768) |
| [SDL_FlushEvents](https://wiki.libsdl.org/SDL3/SDL_FlushEvents) | [:heavy_check_mark:](sdl/functions.go#L257) | [:x:](sdl/sdl_functions_js.go#L10779) |
| [SDL_PollEvent](https://wiki.libsdl.org/SDL3/SDL_PollEvent) | [:heavy_check_mark:](sdl/functions.go#L263) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10792) |
| [SDL_WaitEvent](https://wiki.libsdl.org/SDL3/SDL_WaitEvent) | [:heavy_check_mark:](sdl/functions.go#L269) | [:x:](sdl/sdl_functions_js.go#L10805) |
| [SDL_WaitEventTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitEventTimeout) | [:heavy_check_mark:](sdl/functions.go#L279) | [:x:](sdl/sdl_functions_js.go#L10821) |
| [SDL_PushEvent](https://wiki.libsdl.org/SDL3/SDL_PushEvent) | [:heavy_check_mark:](sdl/functions.go#L285) | [:x:](sdl/sdl_functions_js.go#L10839) |
| [SDL_SetEventFilter](https://wiki.libsdl.org/SDL3/SDL_SetEventFilter) | [:heavy_check_mark:](sdl/functions.go#L295) | [:x:](sdl/sdl_functions_js.go#L10855) |
| [SDL_GetEventFilter](https://wiki.libsdl.org/SDL3/SDL_GetEventFilter) | [:heavy_check_mark:](sdl/functions.go#L301) | [:x:](sdl/sdl_functions_js.go#L10868) |
| [SDL_AddEventWatch](https://wiki.libsdl.org/SDL3/SDL_AddEventWatch) | [:heavy_check_mark:](sdl/functions.go#L314) | [:x:](sdl/sdl_functions_js.go#L10889) |
| [SDL_RemoveEventWatch](https://wiki.libsdl.org/SDL3/SDL_RemoveEventWatch) | [:heavy_check_mark:](sdl/functions.go#L324) | [:x:](sdl/sdl_functions_js.go#L10904) |
| [SDL_FilterEvents](https://wiki.libsdl.org/SDL3/SDL_FilterEvents) | [:heavy_check_mark:](sdl/functions.go#L330) | [:x:](sdl/sdl_functions_js.go#L10917) |
| [SDL_SetEventEnabled](https://wiki.libsdl.org/SDL3/SDL_SetEventEnabled) | [:heavy_check_mark:](sdl/functions.go#L336) | [:x:](sdl/sdl_functions_js.go#L10930) |
| [SDL_EventEnabled](https://wiki.libsdl.org/SDL3/SDL_EventEnabled) | [:heavy_check_mark:](sdl/functions.go#L342) | [:x:](sdl/sdl_functions_js.go#L10943) |
| [SDL_RegisterEvents](https://wiki.libsdl.org/SDL3/SDL_RegisterEvents) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10956) |
| [SDL_GetWindowFromEvent](https://wiki.libsdl.org/SDL3/SDL_GetWindowFromEvent) | [:heavy_check_mark:](sdl/methods.go#L1964) | [:x:](sdl/sdl_functions_js.go#L10969) |
| [SDL_GetEventDescription](https://wiki.libsdl.org/SDL3/SDL_GetEventDescription) | [:heavy_check_mark:](sdl/methods.go#L1970) | [:question:]() |
</details>
<details open>
<summary><h3>Keyboard</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_HasKeyboard](https://wiki.libsdl.org/SDL3/SDL_HasKeyboard) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9933) |
| [SDL_GetKeyboards](https://wiki.libsdl.org/SDL3/SDL_GetKeyboards) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9944) |
| [SDL_GetKeyboardNameForID](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardNameForID) | [:heavy_check_mark:](sdl/methods.go#L3815) | [:x:](sdl/sdl_functions_js.go#L9960) |
| [SDL_GetKeyboardFocus](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardFocus) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9973) |
| [SDL_GetKeyboardState](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9987) |
| [SDL_ResetKeyboard](https://wiki.libsdl.org/SDL3/SDL_ResetKeyboard) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10005) |
| [SDL_GetModState](https://wiki.libsdl.org/SDL3/SDL_GetModState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10014) |
| [SDL_SetModState](https://wiki.libsdl.org/SDL3/SDL_SetModState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10025) |
| [SDL_GetKeyFromScancode](https://wiki.libsdl.org/SDL3/SDL_GetKeyFromScancode) | [:heavy_check_mark:](sdl/methods.go#L4910) | [:x:](sdl/sdl_functions_js.go#L10036) |
| [SDL_GetScancodeFromKey](https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromKey) | [:heavy_check_mark:](sdl/methods.go#L6001) | [:x:](sdl/sdl_functions_js.go#L10053) |
| [SDL_SetScancodeName](https://wiki.libsdl.org/SDL3/SDL_SetScancodeName) | [:heavy_check_mark:](sdl/methods.go#L4916) | [:x:](sdl/sdl_functions_js.go#L10071) |
| [SDL_GetScancodeName](https://wiki.libsdl.org/SDL3/SDL_GetScancodeName) | [:heavy_check_mark:](sdl/methods.go#L4926) | [:x:](sdl/sdl_functions_js.go#L10086) |
| [SDL_GetScancodeFromName](https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromName) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10099) |
| [SDL_GetKeyName](https://wiki.libsdl.org/SDL3/SDL_GetKeyName) | [:heavy_check_mark:](sdl/methods.go#L6007) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10112) |
| [SDL_GetKeyFromName](https://wiki.libsdl.org/SDL3/SDL_GetKeyFromName) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10122) |
| [SDL_StartTextInput](https://wiki.libsdl.org/SDL3/SDL_StartTextInput) | [:heavy_check_mark:](sdl/methods.go#L4787) | [:x:](sdl/sdl_functions_js.go#L10135) |
| [SDL_StartTextInputWithProperties](https://wiki.libsdl.org/SDL3/SDL_StartTextInputWithProperties) | [:heavy_check_mark:](sdl/methods.go#L4797) | [:x:](sdl/sdl_functions_js.go#L10151) |
| [SDL_TextInputActive](https://wiki.libsdl.org/SDL3/SDL_TextInputActive) | [:heavy_check_mark:](sdl/methods.go#L4807) | [:x:](sdl/sdl_functions_js.go#L10169) |
| [SDL_StopTextInput](https://wiki.libsdl.org/SDL3/SDL_StopTextInput) | [:heavy_check_mark:](sdl/methods.go#L4813) | [:x:](sdl/sdl_functions_js.go#L10185) |
| [SDL_ClearComposition](https://wiki.libsdl.org/SDL3/SDL_ClearComposition) | [:heavy_check_mark:](sdl/methods.go#L4823) | [:x:](sdl/sdl_functions_js.go#L10201) |
| [SDL_SetTextInputArea](https://wiki.libsdl.org/SDL3/SDL_SetTextInputArea) | [:heavy_check_mark:](sdl/methods.go#L4833) | [:x:](sdl/sdl_functions_js.go#L10217) |
| [SDL_GetTextInputArea](https://wiki.libsdl.org/SDL3/SDL_GetTextInputArea) | [:heavy_check_mark:](sdl/methods.go#L4843) | [:x:](sdl/sdl_functions_js.go#L10240) |
| [SDL_HasScreenKeyboardSupport](https://wiki.libsdl.org/SDL3/SDL_HasScreenKeyboardSupport) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10266) |
| [SDL_ScreenKeyboardShown](https://wiki.libsdl.org/SDL3/SDL_ScreenKeyboardShown) | [:heavy_check_mark:](sdl/methods.go#L4856) | [:x:](sdl/sdl_functions_js.go#L10277) |
</details>
<details open>
<summary><h3>Mouse</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_HasMouse](https://wiki.libsdl.org/SDL3/SDL_HasMouse) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10293) |
| [SDL_GetMice](https://wiki.libsdl.org/SDL3/SDL_GetMice) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10304) |
| [SDL_GetMouseNameForID](https://wiki.libsdl.org/SDL3/SDL_GetMouseNameForID) | [:heavy_check_mark:](sdl/methods.go#L3828) | [:x:](sdl/sdl_functions_js.go#L10320) |
| [SDL_GetMouseFocus](https://wiki.libsdl.org/SDL3/SDL_GetMouseFocus) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10333) |
| [SDL_GetMouseState](https://wiki.libsdl.org/SDL3/SDL_GetMouseState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10347) |
| [SDL_GetGlobalMouseState](https://wiki.libsdl.org/SDL3/SDL_GetGlobalMouseState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10363) |
| [SDL_GetRelativeMouseState](https://wiki.libsdl.org/SDL3/SDL_GetRelativeMouseState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10384) |
| [SDL_WarpMouseInWindow](https://wiki.libsdl.org/SDL3/SDL_WarpMouseInWindow) | [:heavy_check_mark:](sdl/methods.go#L4862) | [:x:](sdl/sdl_functions_js.go#L10405) |
| [SDL_WarpMouseGlobal](https://wiki.libsdl.org/SDL3/SDL_WarpMouseGlobal) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10423) |
| [SDL_SetRelativeMouseTransform](https://wiki.libsdl.org/SDL3/SDL_SetRelativeMouseTransform) | [:heavy_check_mark:](sdl/functions.go#L893) | [:question:]() |
| [SDL_SetWindowRelativeMouseMode](https://wiki.libsdl.org/SDL3/SDL_SetWindowRelativeMouseMode) | [:heavy_check_mark:](sdl/methods.go#L4868) | [:x:](sdl/sdl_functions_js.go#L10438) |
| [SDL_GetWindowRelativeMouseMode](https://wiki.libsdl.org/SDL3/SDL_GetWindowRelativeMouseMode) | [:heavy_check_mark:](sdl/methods.go#L4878) | [:x:](sdl/sdl_functions_js.go#L10456) |
| [SDL_CaptureMouse](https://wiki.libsdl.org/SDL3/SDL_CaptureMouse) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10472) |
| [SDL_CreateCursor](https://wiki.libsdl.org/SDL3/SDL_CreateCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10485) |
| [SDL_CreateColorCursor](https://wiki.libsdl.org/SDL3/SDL_CreateColorCursor) | [:heavy_check_mark:](sdl/methods.go#L1940) | [:x:](sdl/sdl_functions_js.go#L10517) |
| [SDL_CreateAnimatedCursor](https://wiki.libsdl.org/SDL3/SDL_CreateAnimatedCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:question:]() |
| [SDL_CreateSystemCursor](https://wiki.libsdl.org/SDL3/SDL_CreateSystemCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10540) |
| [SDL_SetCursor](https://wiki.libsdl.org/SDL3/SDL_SetCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10556) |
| [SDL_GetCursor](https://wiki.libsdl.org/SDL3/SDL_GetCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10572) |
| [SDL_GetDefaultCursor](https://wiki.libsdl.org/SDL3/SDL_GetDefaultCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10586) |
| [SDL_DestroyCursor](https://wiki.libsdl.org/SDL3/SDL_DestroyCursor) | [:heavy_check_mark:](sdl/methods.go#L663) | [:x:](sdl/sdl_functions_js.go#L10600) |
| [SDL_ShowCursor](https://wiki.libsdl.org/SDL3/SDL_ShowCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10614) |
| [SDL_HideCursor](https://wiki.libsdl.org/SDL3/SDL_HideCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10625) |
| [SDL_CursorVisible](https://wiki.libsdl.org/SDL3/SDL_CursorVisible) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10636) |
</details>
<details open>
<summary><h3>Touch</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetTouchDevices](https://wiki.libsdl.org/SDL3/SDL_GetTouchDevices) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L10647) |
| [SDL_GetTouchDeviceName](https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceName) | [:heavy_check_mark:](sdl/methods.go#L45) | [:x:](sdl/sdl_functions_js.go#L10663) |
| [SDL_GetTouchDeviceType](https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceType) | [:heavy_check_mark:](sdl/methods.go#L56) | [:x:](sdl/sdl_functions_js.go#L10676) |
| [SDL_GetTouchFingers](https://wiki.libsdl.org/SDL3/SDL_GetTouchFingers) | [:heavy_check_mark:](sdl/methods.go#L62) | [:x:](sdl/sdl_functions_js.go#L10689) |
</details>
<details open>
<summary><h3>Gamepad</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AddGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMapping) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L8756) |
| [SDL_AddGamepadMappingsFromIO](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromIO) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8769) |
| [SDL_AddGamepadMappingsFromFile](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromFile) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L8787) |
| [SDL_ReloadGamepadMappings](https://wiki.libsdl.org/SDL3/SDL_ReloadGamepadMappings) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L8800) |
| [SDL_GetGamepadMappings](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappings) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L8811) |
| [SDL_GetGamepadMappingForGUID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForGUID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8827) |
| [SDL_GetGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMapping) | [:heavy_check_mark:](sdl/methods.go#L2503) | [:x:](sdl/sdl_functions_js.go#L8840) |
| [SDL_SetGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_SetGamepadMapping) | [:heavy_check_mark:](sdl/methods.go#L800) | [:x:](sdl/sdl_functions_js.go#L8856) |
| [SDL_HasGamepad](https://wiki.libsdl.org/SDL3/SDL_HasGamepad) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L8873) |
| [SDL_GetGamepads](https://wiki.libsdl.org/SDL3/SDL_GetGamepads) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L8884) |
| [SDL_IsGamepad](https://wiki.libsdl.org/SDL3/SDL_IsGamepad) | [:heavy_check_mark:](sdl/methods.go#L812) | [:x:](sdl/sdl_functions_js.go#L8900) |
| [SDL_GetGamepadNameForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadNameForID) | [:heavy_check_mark:](sdl/methods.go#L818) | [:x:](sdl/sdl_functions_js.go#L8913) |
| [SDL_GetGamepadPathForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPathForID) | [:heavy_check_mark:](sdl/methods.go#L829) | [:x:](sdl/sdl_functions_js.go#L8926) |
| [SDL_GetGamepadPlayerIndexForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndexForID) | [:heavy_check_mark:](sdl/methods.go#L840) | [:x:](sdl/sdl_functions_js.go#L8939) |
| [SDL_GetGamepadGUIDForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadGUIDForID) | [:x:](sdl/methods.go#L846) | [:x:](sdl/sdl_functions_js.go#L8952) |
| [SDL_GetGamepadVendorForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendorForID) | [:heavy_check_mark:](sdl/methods.go#L853) | [:x:](sdl/sdl_functions_js.go#L8965) |
| [SDL_GetGamepadProductForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductForID) | [:heavy_check_mark:](sdl/methods.go#L859) | [:x:](sdl/sdl_functions_js.go#L8978) |
| [SDL_GetGamepadProductVersionForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersionForID) | [:heavy_check_mark:](sdl/methods.go#L865) | [:x:](sdl/sdl_functions_js.go#L8991) |
| [SDL_GetGamepadTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeForID) | [:heavy_check_mark:](sdl/methods.go#L871) | [:x:](sdl/sdl_functions_js.go#L9004) |
| [SDL_GetRealGamepadTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadTypeForID) | [:heavy_check_mark:](sdl/methods.go#L877) | [:x:](sdl/sdl_functions_js.go#L9017) |
| [SDL_GetGamepadMappingForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForID) | [:heavy_check_mark:](sdl/methods.go#L883) | [:x:](sdl/sdl_functions_js.go#L9030) |
| [SDL_OpenGamepad](https://wiki.libsdl.org/SDL3/SDL_OpenGamepad) | [:heavy_check_mark:](sdl/methods.go#L895) | [:x:](sdl/sdl_functions_js.go#L9043) |
| [SDL_GetGamepadFromID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromID) | [:heavy_check_mark:](sdl/methods.go#L906) | [:x:](sdl/sdl_functions_js.go#L9059) |
| [SDL_GetGamepadFromPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromPlayerIndex) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9075) |
| [SDL_GetGamepadProperties](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProperties) | [:heavy_check_mark:](sdl/methods.go#L2515) | [:x:](sdl/sdl_functions_js.go#L9091) |
| [SDL_GetGamepadID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadID) | [:heavy_check_mark:](sdl/methods.go#L2526) | [:x:](sdl/sdl_functions_js.go#L9107) |
| [SDL_GetGamepadName](https://wiki.libsdl.org/SDL3/SDL_GetGamepadName) | [:heavy_check_mark:](sdl/methods.go#L2537) | [:x:](sdl/sdl_functions_js.go#L9123) |
| [SDL_GetGamepadPath](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPath) | [:heavy_check_mark:](sdl/methods.go#L2543) | [:x:](sdl/sdl_functions_js.go#L9139) |
| [SDL_GetGamepadType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadType) | [:heavy_check_mark:](sdl/methods.go#L2549) | [:x:](sdl/sdl_functions_js.go#L9155) |
| [SDL_GetRealGamepadType](https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadType) | [:heavy_check_mark:](sdl/methods.go#L2555) | [:x:](sdl/sdl_functions_js.go#L9171) |
| [SDL_GetGamepadPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L2561) | [:x:](sdl/sdl_functions_js.go#L9187) |
| [SDL_SetGamepadPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_SetGamepadPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L2567) | [:x:](sdl/sdl_functions_js.go#L9203) |
| [SDL_GetGamepadVendor](https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendor) | [:heavy_check_mark:](sdl/methods.go#L2577) | [:x:](sdl/sdl_functions_js.go#L9221) |
| [SDL_GetGamepadProduct](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProduct) | [:heavy_check_mark:](sdl/methods.go#L2583) | [:x:](sdl/sdl_functions_js.go#L9237) |
| [SDL_GetGamepadProductVersion](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersion) | [:heavy_check_mark:](sdl/methods.go#L2589) | [:x:](sdl/sdl_functions_js.go#L9253) |
| [SDL_GetGamepadFirmwareVersion](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFirmwareVersion) | [:heavy_check_mark:](sdl/methods.go#L2595) | [:x:](sdl/sdl_functions_js.go#L9269) |
| [SDL_GetGamepadSerial](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSerial) | [:heavy_check_mark:](sdl/methods.go#L2601) | [:x:](sdl/sdl_functions_js.go#L9285) |
| [SDL_GetGamepadSteamHandle](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSteamHandle) | [:heavy_check_mark:](sdl/methods.go#L2607) | [:x:](sdl/sdl_functions_js.go#L9301) |
| [SDL_GetGamepadConnectionState](https://wiki.libsdl.org/SDL3/SDL_GetGamepadConnectionState) | [:heavy_check_mark:](sdl/methods.go#L2613) | [:x:](sdl/sdl_functions_js.go#L9317) |
| [SDL_GetGamepadPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPowerInfo) | [:heavy_check_mark:](sdl/methods.go#L2625) | [:x:](sdl/sdl_functions_js.go#L9333) |
| [SDL_GamepadConnected](https://wiki.libsdl.org/SDL3/SDL_GamepadConnected) | [:heavy_check_mark:](sdl/methods.go#L2635) | [:x:](sdl/sdl_functions_js.go#L9354) |
| [SDL_GetGamepadJoystick](https://wiki.libsdl.org/SDL3/SDL_GetGamepadJoystick) | [:heavy_check_mark:](sdl/methods.go#L2641) | [:x:](sdl/sdl_functions_js.go#L9370) |
| [SDL_SetGamepadEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_SetGamepadEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9389) |
| [SDL_GamepadEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_GamepadEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9400) |
| [SDL_GetGamepadBindings](https://wiki.libsdl.org/SDL3/SDL_GetGamepadBindings) | [:heavy_check_mark:](sdl/methods.go#L2652) | [:x:](sdl/sdl_functions_js.go#L9411) |
| [SDL_UpdateGamepads](https://wiki.libsdl.org/SDL3/SDL_UpdateGamepads) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9432) |
| [SDL_GetGamepadTypeFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeFromString) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9441) |
| [SDL_GetGamepadStringForType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForType) | [:heavy_check_mark:](sdl/methods.go#L5730) | [:x:](sdl/sdl_functions_js.go#L9454) |
| [SDL_GetGamepadAxisFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxisFromString) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9467) |
| [SDL_GetGamepadStringForAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForAxis) | [:heavy_check_mark:](sdl/methods.go#L655) | [:x:](sdl/sdl_functions_js.go#L9480) |
| [SDL_GamepadHasAxis](https://wiki.libsdl.org/SDL3/SDL_GamepadHasAxis) | [:heavy_check_mark:](sdl/methods.go#L2666) | [:x:](sdl/sdl_functions_js.go#L9493) |
| [SDL_GetGamepadAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxis) | [:heavy_check_mark:](sdl/methods.go#L2672) | [:x:](sdl/sdl_functions_js.go#L9511) |
| [SDL_GetGamepadButtonFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonFromString) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L9529) |
| [SDL_GetGamepadStringForButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForButton) | [:heavy_check_mark:](sdl/methods.go#L550) | [:x:](sdl/sdl_functions_js.go#L9542) |
| [SDL_GamepadHasButton](https://wiki.libsdl.org/SDL3/SDL_GamepadHasButton) | [:heavy_check_mark:](sdl/methods.go#L2678) | [:x:](sdl/sdl_functions_js.go#L9555) |
| [SDL_GetGamepadButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButton) | [:heavy_check_mark:](sdl/methods.go#L2684) | [:x:](sdl/sdl_functions_js.go#L9573) |
| [SDL_GetGamepadButtonLabelForType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabelForType) | [:heavy_check_mark:](sdl/methods.go#L5736) | [:x:](sdl/sdl_functions_js.go#L9591) |
| [SDL_GetGamepadButtonLabel](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabel) | [:heavy_check_mark:](sdl/methods.go#L2690) | [:x:](sdl/sdl_functions_js.go#L9606) |
| [SDL_GetNumGamepadTouchpads](https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpads) | [:heavy_check_mark:](sdl/methods.go#L2696) | [:x:](sdl/sdl_functions_js.go#L9624) |
| [SDL_GetNumGamepadTouchpadFingers](https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpadFingers) | [:heavy_check_mark:](sdl/methods.go#L2702) | [:x:](sdl/sdl_functions_js.go#L9640) |
| [SDL_GetGamepadTouchpadFinger](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTouchpadFinger) | [:x:](sdl/methods.go#L2708) | [:x:](sdl/sdl_functions_js.go#L9658) |
| [SDL_GamepadHasSensor](https://wiki.libsdl.org/SDL3/SDL_GamepadHasSensor) | [:heavy_check_mark:](sdl/methods.go#L2715) | [:x:](sdl/sdl_functions_js.go#L9698) |
| [SDL_SetGamepadSensorEnabled](https://wiki.libsdl.org/SDL3/SDL_SetGamepadSensorEnabled) | [:heavy_check_mark:](sdl/methods.go#L2721) | [:x:](sdl/sdl_functions_js.go#L9716) |
| [SDL_GamepadSensorEnabled](https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEnabled) | [:heavy_check_mark:](sdl/methods.go#L2731) | [:x:](sdl/sdl_functions_js.go#L9736) |
| [SDL_GetGamepadSensorDataRate](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorDataRate) | [:heavy_check_mark:](sdl/methods.go#L2737) | [:x:](sdl/sdl_functions_js.go#L9754) |
| [SDL_GetGamepadSensorData](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorData) | [:x:](sdl/methods.go#L2743) | [:x:](sdl/sdl_functions_js.go#L9772) |
| [SDL_RumbleGamepad](https://wiki.libsdl.org/SDL3/SDL_RumbleGamepad) | [:heavy_check_mark:](sdl/methods.go#L2750) | [:x:](sdl/sdl_functions_js.go#L9797) |
| [SDL_RumbleGamepadTriggers](https://wiki.libsdl.org/SDL3/SDL_RumbleGamepadTriggers) | [:heavy_check_mark:](sdl/methods.go#L2760) | [:x:](sdl/sdl_functions_js.go#L9819) |
| [SDL_SetGamepadLED](https://wiki.libsdl.org/SDL3/SDL_SetGamepadLED) | [:heavy_check_mark:](sdl/methods.go#L2770) | [:x:](sdl/sdl_functions_js.go#L9841) |
| [SDL_SendGamepadEffect](https://wiki.libsdl.org/SDL3/SDL_SendGamepadEffect) | [:heavy_check_mark:](sdl/methods.go#L2780) | [:x:](sdl/sdl_functions_js.go#L9863) |
| [SDL_CloseGamepad](https://wiki.libsdl.org/SDL3/SDL_CloseGamepad) | [:heavy_check_mark:](sdl/methods.go#L2791) | [:x:](sdl/sdl_functions_js.go#L9883) |
| [SDL_GetGamepadAppleSFSymbolsNameForButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForButton) | [:heavy_check_mark:](sdl/methods.go#L2797) | [:x:](sdl/sdl_functions_js.go#L9897) |
| [SDL_GetGamepadAppleSFSymbolsNameForAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForAxis) | [:heavy_check_mark:](sdl/methods.go#L2803) | [:x:](sdl/sdl_functions_js.go#L9915) |
</details>
<details open>
<summary><h3>Joystick</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_LockJoysticks](https://wiki.libsdl.org/SDL3/SDL_LockJoysticks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7817) |
| [SDL_UnlockJoysticks](https://wiki.libsdl.org/SDL3/SDL_UnlockJoysticks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7826) |
| [SDL_HasJoystick](https://wiki.libsdl.org/SDL3/SDL_HasJoystick) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7835) |
| [SDL_GetJoysticks](https://wiki.libsdl.org/SDL3/SDL_GetJoysticks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7846) |
| [SDL_GetJoystickNameForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickNameForID) | [:heavy_check_mark:](sdl/methods.go#L703) | [:x:](sdl/sdl_functions_js.go#L7862) |
| [SDL_GetJoystickPathForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPathForID) | [:heavy_check_mark:](sdl/methods.go#L714) | [:x:](sdl/sdl_functions_js.go#L7875) |
| [SDL_GetJoystickPlayerIndexForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndexForID) | [:heavy_check_mark:](sdl/methods.go#L725) | [:x:](sdl/sdl_functions_js.go#L7888) |
| [SDL_GetJoystickGUIDForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDForID) | [:x:](sdl/methods.go#L731) | [:x:](sdl/sdl_functions_js.go#L7901) |
| [SDL_GetJoystickVendorForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendorForID) | [:heavy_check_mark:](sdl/methods.go#L738) | [:x:](sdl/sdl_functions_js.go#L7914) |
| [SDL_GetJoystickProductForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductForID) | [:heavy_check_mark:](sdl/methods.go#L744) | [:x:](sdl/sdl_functions_js.go#L7927) |
| [SDL_GetJoystickProductVersionForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersionForID) | [:heavy_check_mark:](sdl/methods.go#L750) | [:x:](sdl/sdl_functions_js.go#L7940) |
| [SDL_GetJoystickTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickTypeForID) | [:heavy_check_mark:](sdl/methods.go#L756) | [:x:](sdl/sdl_functions_js.go#L7953) |
| [SDL_OpenJoystick](https://wiki.libsdl.org/SDL3/SDL_OpenJoystick) | [:heavy_check_mark:](sdl/methods.go#L762) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7966) |
| [SDL_GetJoystickFromID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromID) | [:heavy_check_mark:](sdl/methods.go#L773) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7978) |
| [SDL_GetJoystickFromPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromPlayerIndex) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7992) |
| [SDL_AttachVirtualJoystick](https://wiki.libsdl.org/SDL3/SDL_AttachVirtualJoystick) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L8008) |
| [SDL_DetachVirtualJoystick](https://wiki.libsdl.org/SDL3/SDL_DetachVirtualJoystick) | [:heavy_check_mark:](sdl/methods.go#L784) | [:x:](sdl/sdl_functions_js.go#L8024) |
| [SDL_IsJoystickVirtual](https://wiki.libsdl.org/SDL3/SDL_IsJoystickVirtual) | [:heavy_check_mark:](sdl/methods.go#L794) | [:x:](sdl/sdl_functions_js.go#L8037) |
| [SDL_SetJoystickVirtualAxis](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualAxis) | [:heavy_check_mark:](sdl/methods.go#L5388) | [:x:](sdl/sdl_functions_js.go#L8050) |
| [SDL_SetJoystickVirtualBall](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualBall) | [:heavy_check_mark:](sdl/methods.go#L5398) | [:x:](sdl/sdl_functions_js.go#L8070) |
| [SDL_SetJoystickVirtualButton](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualButton) | [:heavy_check_mark:](sdl/methods.go#L5408) | [:x:](sdl/sdl_functions_js.go#L8092) |
| [SDL_SetJoystickVirtualHat](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualHat) | [:heavy_check_mark:](sdl/methods.go#L5418) | [:x:](sdl/sdl_functions_js.go#L8112) |
| [SDL_SetJoystickVirtualTouchpad](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualTouchpad) | [:heavy_check_mark:](sdl/methods.go#L5428) | [:x:](sdl/sdl_functions_js.go#L8132) |
| [SDL_SendJoystickVirtualSensorData](https://wiki.libsdl.org/SDL3/SDL_SendJoystickVirtualSensorData) | [:heavy_check_mark:](sdl/methods.go#L5438) | [:x:](sdl/sdl_functions_js.go#L8160) |
| [SDL_GetJoystickProperties](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProperties) | [:heavy_check_mark:](sdl/methods.go#L5448) | [:x:](sdl/sdl_functions_js.go#L8187) |
| [SDL_GetJoystickName](https://wiki.libsdl.org/SDL3/SDL_GetJoystickName) | [:heavy_check_mark:](sdl/methods.go#L5459) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8203) |
| [SDL_GetJoystickPath](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPath) | [:heavy_check_mark:](sdl/methods.go#L5470) | [:x:](sdl/sdl_functions_js.go#L8216) |
| [SDL_GetJoystickPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L5481) | [:x:](sdl/sdl_functions_js.go#L8232) |
| [SDL_SetJoystickPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_SetJoystickPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L5487) | [:x:](sdl/sdl_functions_js.go#L8248) |
| [SDL_GetJoystickGUID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUID) | [:x:](sdl/methods.go#L5497) | [:x:](sdl/sdl_functions_js.go#L8266) |
| [SDL_GetJoystickVendor](https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendor) | [:heavy_check_mark:](sdl/methods.go#L5504) | [:x:](sdl/sdl_functions_js.go#L8282) |
| [SDL_GetJoystickProduct](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProduct) | [:heavy_check_mark:](sdl/methods.go#L5510) | [:x:](sdl/sdl_functions_js.go#L8298) |
| [SDL_GetJoystickProductVersion](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersion) | [:heavy_check_mark:](sdl/methods.go#L5516) | [:x:](sdl/sdl_functions_js.go#L8314) |
| [SDL_GetJoystickFirmwareVersion](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFirmwareVersion) | [:heavy_check_mark:](sdl/methods.go#L5522) | [:x:](sdl/sdl_functions_js.go#L8330) |
| [SDL_GetJoystickSerial](https://wiki.libsdl.org/SDL3/SDL_GetJoystickSerial) | [:heavy_check_mark:](sdl/methods.go#L5528) | [:x:](sdl/sdl_functions_js.go#L8346) |
| [SDL_GetJoystickType](https://wiki.libsdl.org/SDL3/SDL_GetJoystickType) | [:heavy_check_mark:](sdl/methods.go#L5534) | [:x:](sdl/sdl_functions_js.go#L8362) |
| [SDL_GetJoystickGUIDInfo](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDInfo) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8378) |
| [SDL_JoystickConnected](https://wiki.libsdl.org/SDL3/SDL_JoystickConnected) | [:heavy_check_mark:](sdl/methods.go#L5540) | [:x:](sdl/sdl_functions_js.go#L8409) |
| [SDL_GetJoystickID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickID) | [:heavy_check_mark:](sdl/methods.go#L5546) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8425) |
| [SDL_GetNumJoystickAxes](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickAxes) | [:heavy_check_mark:](sdl/methods.go#L5557) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8438) |
| [SDL_GetNumJoystickBalls](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickBalls) | [:heavy_check_mark:](sdl/methods.go#L5568) | [:x:](sdl/sdl_functions_js.go#L8451) |
| [SDL_GetNumJoystickHats](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickHats) | [:heavy_check_mark:](sdl/methods.go#L5579) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8467) |
| [SDL_GetNumJoystickButtons](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickButtons) | [:heavy_check_mark:](sdl/methods.go#L5590) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8480) |
| [SDL_SetJoystickEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_SetJoystickEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L8493) |
| [SDL_JoystickEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_JoystickEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L8504) |
| [SDL_UpdateJoysticks](https://wiki.libsdl.org/SDL3/SDL_UpdateJoysticks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L8515) |
| [SDL_GetJoystickAxis](https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxis) | [:heavy_check_mark:](sdl/methods.go#L5601) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8524) |
| [SDL_GetJoystickAxisInitialState](https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxisInitialState) | [:heavy_check_mark:](sdl/methods.go#L5612) | [:x:](sdl/sdl_functions_js.go#L8539) |
| [SDL_GetJoystickBall](https://wiki.libsdl.org/SDL3/SDL_GetJoystickBall) | [:heavy_check_mark:](sdl/methods.go#L5621) | [:x:](sdl/sdl_functions_js.go#L8562) |
| [SDL_GetJoystickHat](https://wiki.libsdl.org/SDL3/SDL_GetJoystickHat) | [:heavy_check_mark:](sdl/methods.go#L5633) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8590) |
| [SDL_GetJoystickButton](https://wiki.libsdl.org/SDL3/SDL_GetJoystickButton) | [:heavy_check_mark:](sdl/methods.go#L5639) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8605) |
| [SDL_RumbleJoystick](https://wiki.libsdl.org/SDL3/SDL_RumbleJoystick) | [:heavy_check_mark:](sdl/methods.go#L5645) | [:x:](sdl/sdl_functions_js.go#L8620) |
| [SDL_RumbleJoystickTriggers](https://wiki.libsdl.org/SDL3/SDL_RumbleJoystickTriggers) | [:heavy_check_mark:](sdl/methods.go#L5651) | [:x:](sdl/sdl_functions_js.go#L8642) |
| [SDL_SetJoystickLED](https://wiki.libsdl.org/SDL3/SDL_SetJoystickLED) | [:heavy_check_mark:](sdl/methods.go#L5661) | [:x:](sdl/sdl_functions_js.go#L8664) |
| [SDL_SendJoystickEffect](https://wiki.libsdl.org/SDL3/SDL_SendJoystickEffect) | [:heavy_check_mark:](sdl/methods.go#L5671) | [:x:](sdl/sdl_functions_js.go#L8686) |
| [SDL_CloseJoystick](https://wiki.libsdl.org/SDL3/SDL_CloseJoystick) | [:heavy_check_mark:](sdl/methods.go#L5681) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8706) |
| [SDL_GetJoystickConnectionState](https://wiki.libsdl.org/SDL3/SDL_GetJoystickConnectionState) | [:heavy_check_mark:](sdl/methods.go#L5687) | [:x:](sdl/sdl_functions_js.go#L8719) |
| [SDL_GetJoystickPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPowerInfo) | [:heavy_check_mark:](sdl/methods.go#L5698) | [:x:](sdl/sdl_functions_js.go#L8735) |
</details>
<details open>
<summary><h3>Haptic</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetHaptics](https://wiki.libsdl.org/SDL3/SDL_GetHaptics) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L13043) |
| [SDL_GetHapticNameForID](https://wiki.libsdl.org/SDL3/SDL_GetHapticNameForID) | [:heavy_check_mark:](sdl/methods.go#L2855) | [:x:](sdl/sdl_functions_js.go#L13059) |
| [SDL_OpenHaptic](https://wiki.libsdl.org/SDL3/SDL_OpenHaptic) | [:heavy_check_mark:](sdl/methods.go#L2866) | [:x:](sdl/sdl_functions_js.go#L13072) |
| [SDL_GetHapticFromID](https://wiki.libsdl.org/SDL3/SDL_GetHapticFromID) | [:heavy_check_mark:](sdl/methods.go#L2877) | [:x:](sdl/sdl_functions_js.go#L13088) |
| [SDL_GetHapticID](https://wiki.libsdl.org/SDL3/SDL_GetHapticID) | [:heavy_check_mark:](sdl/methods.go#L2284) | [:x:](sdl/sdl_functions_js.go#L13104) |
| [SDL_GetHapticName](https://wiki.libsdl.org/SDL3/SDL_GetHapticName) | [:heavy_check_mark:](sdl/methods.go#L2295) | [:x:](sdl/sdl_functions_js.go#L13120) |
| [SDL_IsMouseHaptic](https://wiki.libsdl.org/SDL3/SDL_IsMouseHaptic) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L13136) |
| [SDL_OpenHapticFromMouse](https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromMouse) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L13147) |
| [SDL_IsJoystickHaptic](https://wiki.libsdl.org/SDL3/SDL_IsJoystickHaptic) | [:heavy_check_mark:](sdl/methods.go#L5711) | [:x:](sdl/sdl_functions_js.go#L13161) |
| [SDL_OpenHapticFromJoystick](https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromJoystick) | [:heavy_check_mark:](sdl/methods.go#L5717) | [:x:](sdl/sdl_functions_js.go#L13177) |
| [SDL_CloseHaptic](https://wiki.libsdl.org/SDL3/SDL_CloseHaptic) | [:heavy_check_mark:](sdl/methods.go#L2306) | [:x:](sdl/sdl_functions_js.go#L13196) |
| [SDL_GetMaxHapticEffects](https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffects) | [:heavy_check_mark:](sdl/methods.go#L2312) | [:x:](sdl/sdl_functions_js.go#L13210) |
| [SDL_GetMaxHapticEffectsPlaying](https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffectsPlaying) | [:heavy_check_mark:](sdl/methods.go#L2323) | [:x:](sdl/sdl_functions_js.go#L13226) |
| [SDL_GetHapticFeatures](https://wiki.libsdl.org/SDL3/SDL_GetHapticFeatures) | [:heavy_check_mark:](sdl/methods.go#L2334) | [:x:](sdl/sdl_functions_js.go#L13242) |
| [SDL_GetNumHapticAxes](https://wiki.libsdl.org/SDL3/SDL_GetNumHapticAxes) | [:heavy_check_mark:](sdl/methods.go#L2345) | [:x:](sdl/sdl_functions_js.go#L13258) |
| [SDL_HapticEffectSupported](https://wiki.libsdl.org/SDL3/SDL_HapticEffectSupported) | [:heavy_check_mark:](sdl/methods.go#L2356) | [:x:](sdl/sdl_functions_js.go#L13274) |
| [SDL_CreateHapticEffect](https://wiki.libsdl.org/SDL3/SDL_CreateHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2362) | [:x:](sdl/sdl_functions_js.go#L13295) |
| [SDL_UpdateHapticEffect](https://wiki.libsdl.org/SDL3/SDL_UpdateHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2373) | [:x:](sdl/sdl_functions_js.go#L13316) |
| [SDL_RunHapticEffect](https://wiki.libsdl.org/SDL3/SDL_RunHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2383) | [:x:](sdl/sdl_functions_js.go#L13339) |
| [SDL_StopHapticEffect](https://wiki.libsdl.org/SDL3/SDL_StopHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2393) | [:x:](sdl/sdl_functions_js.go#L13359) |
| [SDL_DestroyHapticEffect](https://wiki.libsdl.org/SDL3/SDL_DestroyHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2403) | [:x:](sdl/sdl_functions_js.go#L13377) |
| [SDL_GetHapticEffectStatus](https://wiki.libsdl.org/SDL3/SDL_GetHapticEffectStatus) | [:heavy_check_mark:](sdl/methods.go#L2409) | [:x:](sdl/sdl_functions_js.go#L13393) |
| [SDL_SetHapticGain](https://wiki.libsdl.org/SDL3/SDL_SetHapticGain) | [:heavy_check_mark:](sdl/methods.go#L2415) | [:x:](sdl/sdl_functions_js.go#L13411) |
| [SDL_SetHapticAutocenter](https://wiki.libsdl.org/SDL3/SDL_SetHapticAutocenter) | [:heavy_check_mark:](sdl/methods.go#L2425) | [:x:](sdl/sdl_functions_js.go#L13429) |
| [SDL_PauseHaptic](https://wiki.libsdl.org/SDL3/SDL_PauseHaptic) | [:heavy_check_mark:](sdl/methods.go#L2435) | [:x:](sdl/sdl_functions_js.go#L13447) |
| [SDL_ResumeHaptic](https://wiki.libsdl.org/SDL3/SDL_ResumeHaptic) | [:heavy_check_mark:](sdl/methods.go#L2445) | [:x:](sdl/sdl_functions_js.go#L13463) |
| [SDL_StopHapticEffects](https://wiki.libsdl.org/SDL3/SDL_StopHapticEffects) | [:heavy_check_mark:](sdl/methods.go#L2455) | [:x:](sdl/sdl_functions_js.go#L13479) |
| [SDL_HapticRumbleSupported](https://wiki.libsdl.org/SDL3/SDL_HapticRumbleSupported) | [:heavy_check_mark:](sdl/methods.go#L2465) | [:x:](sdl/sdl_functions_js.go#L13495) |
| [SDL_InitHapticRumble](https://wiki.libsdl.org/SDL3/SDL_InitHapticRumble) | [:heavy_check_mark:](sdl/methods.go#L2471) | [:x:](sdl/sdl_functions_js.go#L13511) |
| [SDL_PlayHapticRumble](https://wiki.libsdl.org/SDL3/SDL_PlayHapticRumble) | [:heavy_check_mark:](sdl/methods.go#L2481) | [:x:](sdl/sdl_functions_js.go#L13527) |
| [SDL_StopHapticRumble](https://wiki.libsdl.org/SDL3/SDL_StopHapticRumble) | [:heavy_check_mark:](sdl/methods.go#L2491) | [:x:](sdl/sdl_functions_js.go#L13547) |
</details>
<details open>
<summary><h3>Audio</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumAudioDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumAudioDrivers) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L2358) |
| [SDL_GetAudioDriver](https://wiki.libsdl.org/SDL3/SDL_GetAudioDriver) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L2369) |
| [SDL_GetCurrentAudioDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentAudioDriver) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L2382) |
| [SDL_GetAudioPlaybackDevices](https://wiki.libsdl.org/SDL3/SDL_GetAudioPlaybackDevices) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L2393) |
| [SDL_GetAudioRecordingDevices](https://wiki.libsdl.org/SDL3/SDL_GetAudioRecordingDevices) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L2409) |
| [SDL_GetAudioDeviceName](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceName) | [:heavy_check_mark:](sdl/methods.go#L211) | [:x:](sdl/sdl_functions_js.go#L2425) |
| [SDL_GetAudioDeviceFormat](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceFormat) | [:heavy_check_mark:](sdl/methods.go#L222) | [:x:](sdl/sdl_functions_js.go#L2438) |
| [SDL_GetAudioDeviceChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceChannelMap) | [:heavy_check_mark:](sdl/methods.go#L235) | [:x:](sdl/sdl_functions_js.go#L2461) |
| [SDL_OpenAudioDevice](https://wiki.libsdl.org/SDL3/SDL_OpenAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L249) | [:x:](sdl/sdl_functions_js.go#L2479) |
| [SDL_IsAudioDevicePhysical](https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePhysical) | [:heavy_check_mark:](sdl/methods.go#L260) | [:x:](sdl/sdl_functions_js.go#L2497) |
| [SDL_IsAudioDevicePlayback](https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePlayback) | [:heavy_check_mark:](sdl/methods.go#L266) | [:x:](sdl/sdl_functions_js.go#L2510) |
| [SDL_PauseAudioDevice](https://wiki.libsdl.org/SDL3/SDL_PauseAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L272) | [:x:](sdl/sdl_functions_js.go#L2523) |
| [SDL_ResumeAudioDevice](https://wiki.libsdl.org/SDL3/SDL_ResumeAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L282) | [:x:](sdl/sdl_functions_js.go#L2536) |
| [SDL_AudioDevicePaused](https://wiki.libsdl.org/SDL3/SDL_AudioDevicePaused) | [:heavy_check_mark:](sdl/methods.go#L292) | [:x:](sdl/sdl_functions_js.go#L2549) |
| [SDL_GetAudioDeviceGain](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceGain) | [:heavy_check_mark:](sdl/methods.go#L298) | [:x:](sdl/sdl_functions_js.go#L2562) |
| [SDL_SetAudioDeviceGain](https://wiki.libsdl.org/SDL3/SDL_SetAudioDeviceGain) | [:heavy_check_mark:](sdl/methods.go#L309) | [:x:](sdl/sdl_functions_js.go#L2575) |
| [SDL_CloseAudioDevice](https://wiki.libsdl.org/SDL3/SDL_CloseAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L319) | [:x:](sdl/sdl_functions_js.go#L2590) |
| [SDL_BindAudioStreams](https://wiki.libsdl.org/SDL3/SDL_BindAudioStreams) | [:heavy_check_mark:](sdl/methods.go#L325) | [:x:](sdl/sdl_functions_js.go#L2601) |
| [SDL_BindAudioStream](https://wiki.libsdl.org/SDL3/SDL_BindAudioStream) | [:heavy_check_mark:](sdl/methods.go#L335) | [:x:](sdl/sdl_functions_js.go#L2621) |
| [SDL_UnbindAudioStreams](https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStreams) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L2639) |
| [SDL_UnbindAudioStream](https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStream) | [:heavy_check_mark:](sdl/methods.go#L3843) | [:x:](sdl/sdl_functions_js.go#L2655) |
| [SDL_GetAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamDevice) | [:heavy_check_mark:](sdl/methods.go#L3849) | [:x:](sdl/sdl_functions_js.go#L2669) |
| [SDL_CreateAudioStream](https://wiki.libsdl.org/SDL3/SDL_CreateAudioStream) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L2685) |
| [SDL_GetAudioStreamProperties](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamProperties) | [:heavy_check_mark:](sdl/methods.go#L3855) | [:x:](sdl/sdl_functions_js.go#L2709) |
| [SDL_GetAudioStreamFormat](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFormat) | [:heavy_check_mark:](sdl/methods.go#L3866) | [:x:](sdl/sdl_functions_js.go#L2725) |
| [SDL_SetAudioStreamFormat](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFormat) | [:heavy_check_mark:](sdl/methods.go#L3876) | [:x:](sdl/sdl_functions_js.go#L2751) |
| [SDL_GetAudioStreamFrequencyRatio](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFrequencyRatio) | [:heavy_check_mark:](sdl/methods.go#L3886) | [:x:](sdl/sdl_functions_js.go#L2777) |
| [SDL_SetAudioStreamFrequencyRatio](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFrequencyRatio) | [:heavy_check_mark:](sdl/methods.go#L3897) | [:x:](sdl/sdl_functions_js.go#L2793) |
| [SDL_GetAudioStreamGain](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamGain) | [:heavy_check_mark:](sdl/methods.go#L3907) | [:x:](sdl/sdl_functions_js.go#L2811) |
| [SDL_SetAudioStreamGain](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGain) | [:heavy_check_mark:](sdl/methods.go#L3918) | [:x:](sdl/sdl_functions_js.go#L2827) |
| [SDL_GetAudioStreamInputChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamInputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3928) | [:x:](sdl/sdl_functions_js.go#L2845) |
| [SDL_GetAudioStreamOutputChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamOutputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3942) | [:x:](sdl/sdl_functions_js.go#L2866) |
| [SDL_SetAudioStreamInputChannelMap](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamInputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3956) | [:x:](sdl/sdl_functions_js.go#L2887) |
| [SDL_SetAudioStreamOutputChannelMap](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamOutputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3966) | [:x:](sdl/sdl_functions_js.go#L2910) |
| [SDL_PutAudioStreamData](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData) | [:heavy_check_mark:](sdl/methods.go#L3976) | [:x:](sdl/sdl_functions_js.go#L2933) |
| [SDL_PutAudioStreamDataNoCopy](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamDataNoCopy) | [:question:]() | [:question:]() |
| [SDL_PutAudioStreamPlanarData](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamPlanarData) | [:question:]() | [:question:]() |
| [SDL_GetAudioStreamData](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData) | [:heavy_check_mark:](sdl/methods.go#L3988) | [:x:](sdl/sdl_functions_js.go#L2953) |
| [SDL_GetAudioStreamAvailable](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamAvailable) | [:heavy_check_mark:](sdl/methods.go#L3999) | [:x:](sdl/sdl_functions_js.go#L2973) |
| [SDL_GetAudioStreamQueued](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamQueued) | [:heavy_check_mark:](sdl/methods.go#L4010) | [:x:](sdl/sdl_functions_js.go#L2989) |
| [SDL_FlushAudioStream](https://wiki.libsdl.org/SDL3/SDL_FlushAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4021) | [:x:](sdl/sdl_functions_js.go#L3005) |
| [SDL_ClearAudioStream](https://wiki.libsdl.org/SDL3/SDL_ClearAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4031) | [:x:](sdl/sdl_functions_js.go#L3021) |
| [SDL_PauseAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_PauseAudioStreamDevice) | [:heavy_check_mark:](sdl/methods.go#L4041) | [:x:](sdl/sdl_functions_js.go#L3037) |
| [SDL_ResumeAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_ResumeAudioStreamDevice) | [:heavy_check_mark:](sdl/methods.go#L4051) | [:x:](sdl/sdl_functions_js.go#L3053) |
| [SDL_AudioStreamDevicePaused](https://wiki.libsdl.org/SDL3/SDL_AudioStreamDevicePaused) | [:heavy_check_mark:](sdl/methods.go#L4061) | [:x:](sdl/sdl_functions_js.go#L3069) |
| [SDL_LockAudioStream](https://wiki.libsdl.org/SDL3/SDL_LockAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4067) | [:x:](sdl/sdl_functions_js.go#L3085) |
| [SDL_UnlockAudioStream](https://wiki.libsdl.org/SDL3/SDL_UnlockAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4077) | [:x:](sdl/sdl_functions_js.go#L3101) |
| [SDL_SetAudioStreamGetCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGetCallback) | [:heavy_check_mark:](sdl/methods.go#L4087) | [:x:](sdl/sdl_functions_js.go#L3117) |
| [SDL_SetAudioStreamPutCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamPutCallback) | [:heavy_check_mark:](sdl/methods.go#L4093) | [:x:](sdl/sdl_functions_js.go#L3137) |
| [SDL_DestroyAudioStream](https://wiki.libsdl.org/SDL3/SDL_DestroyAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4099) | [:x:](sdl/sdl_functions_js.go#L3157) |
| [SDL_OpenAudioDeviceStream](https://wiki.libsdl.org/SDL3/SDL_OpenAudioDeviceStream) | [:heavy_check_mark:](sdl/methods.go#L345) | [:x:](sdl/sdl_functions_js.go#L3171) |
| [SDL_SetAudioPostmixCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioPostmixCallback) | [:heavy_check_mark:](sdl/methods.go#L351) | [:x:](sdl/sdl_functions_js.go#L3195) |
| [SDL_LoadWAV_IO](https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L3212) |
| [SDL_LoadWAV](https://wiki.libsdl.org/SDL3/SDL_LoadWAV) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L3245) |
| [SDL_MixAudio](https://wiki.libsdl.org/SDL3/SDL_MixAudio) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L3273) |
| [SDL_ConvertAudioSamples](https://wiki.libsdl.org/SDL3/SDL_ConvertAudioSamples) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L3300) |
| [SDL_GetAudioFormatName](https://wiki.libsdl.org/SDL3/SDL_GetAudioFormatName) | [:heavy_check_mark:](sdl/methods.go#L1285) | [:x:](sdl/sdl_functions_js.go#L3338) |
| [SDL_GetSilenceValueForFormat](https://wiki.libsdl.org/SDL3/SDL_GetSilenceValueForFormat) | [:heavy_check_mark:](sdl/methods.go#L1291) | [:x:](sdl/sdl_functions_js.go#L3351) |
</details>
<details>
<summary><h3>Time</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetDateTimeLocalePreferences](https://wiki.libsdl.org/SDL3/SDL_GetDateTimeLocalePreferences) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16727) |
| [SDL_GetCurrentTime](https://wiki.libsdl.org/SDL3/SDL_GetCurrentTime) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16748) |
| [SDL_TimeToDateTime](https://wiki.libsdl.org/SDL3/SDL_TimeToDateTime) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16764) |
| [SDL_DateTimeToTime](https://wiki.libsdl.org/SDL3/SDL_DateTimeToTime) | [:heavy_check_mark:](sdl/methods.go#L5987) | [:x:](sdl/sdl_functions_js.go#L16784) |
| [SDL_TimeToWindows](https://wiki.libsdl.org/SDL3/SDL_TimeToWindows) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16805) |
| [SDL_TimeFromWindows](https://wiki.libsdl.org/SDL3/SDL_TimeFromWindows) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16826) |
| [SDL_GetDaysInMonth](https://wiki.libsdl.org/SDL3/SDL_GetDaysInMonth) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16841) |
| [SDL_GetDayOfYear](https://wiki.libsdl.org/SDL3/SDL_GetDayOfYear) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16856) |
| [SDL_GetDayOfWeek](https://wiki.libsdl.org/SDL3/SDL_GetDayOfWeek) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16873) |
</details>
<details open>
<summary><h3>Timer</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetTicks](https://wiki.libsdl.org/SDL3/SDL_GetTicks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16890) |
| [SDL_GetTicksNS](https://wiki.libsdl.org/SDL3/SDL_GetTicksNS) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L16898) |
| [SDL_GetPerformanceCounter](https://wiki.libsdl.org/SDL3/SDL_GetPerformanceCounter) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L16909) |
| [SDL_GetPerformanceFrequency](https://wiki.libsdl.org/SDL3/SDL_GetPerformanceFrequency) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L16920) |
| [SDL_Delay](https://wiki.libsdl.org/SDL3/SDL_Delay) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L16931) |
| [SDL_DelayNS](https://wiki.libsdl.org/SDL3/SDL_DelayNS) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L16942) |
| [SDL_DelayPrecise](https://wiki.libsdl.org/SDL3/SDL_DelayPrecise) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L16953) |
| [SDL_AddTimer](https://wiki.libsdl.org/SDL3/SDL_AddTimer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16964) |
| [SDL_AddTimerNS](https://wiki.libsdl.org/SDL3/SDL_AddTimerNS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16981) |
| [SDL_RemoveTimer](https://wiki.libsdl.org/SDL3/SDL_RemoveTimer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16998) |
</details>
<details open>
<summary><h3>Render</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumRenderDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumRenderDrivers) | [:heavy_check_mark:](sdl/functions.go#L420) | [:x:](sdl/sdl_functions_js.go#L14364) |
| [SDL_GetRenderDriver](https://wiki.libsdl.org/SDL3/SDL_GetRenderDriver) | [:heavy_check_mark:](sdl/functions.go#L426) | [:x:](sdl/sdl_functions_js.go#L14375) |
| [SDL_CreateWindowAndRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateWindowAndRenderer) | [:heavy_check_mark:](sdl/functions.go#L432) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L14388) |
| [SDL_CreateRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateRenderer) | [:heavy_check_mark:](sdl/methods.go#L4891) | [:x:](sdl/sdl_functions_js.go#L14416) |
| [SDL_CreateRendererWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateRendererWithProperties) | [:heavy_check_mark:](sdl/functions.go#L445) | [:x:](sdl/sdl_functions_js.go#L14437) |
| [SDL_CreateGPURenderer](https://wiki.libsdl.org/SDL3/SDL_CreateGPURenderer) | [:heavy_check_mark:](sdl/methods.go#L2002) | [:question:]() |
| [SDL_GetGPURendererDevice](https://wiki.libsdl.org/SDL3/SDL_GetGPURendererDevice) | [:heavy_check_mark:](sdl/methods.go#L2901) | [:question:]() |
| [SDL_CreateSoftwareRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateSoftwareRenderer) | [:heavy_check_mark:](sdl/methods.go#L1951) | [:x:](sdl/sdl_functions_js.go#L14453) |
| [SDL_GetRenderer](https://wiki.libsdl.org/SDL3/SDL_GetRenderer) | [:heavy_check_mark:](sdl/methods.go#L4902) | [:x:](sdl/sdl_functions_js.go#L14472) |
| [SDL_GetRenderWindow](https://wiki.libsdl.org/SDL3/SDL_GetRenderWindow) | [:heavy_check_mark:](sdl/methods.go#L2890) | [:x:](sdl/sdl_functions_js.go#L14491) |
| [SDL_GetRendererName](https://wiki.libsdl.org/SDL3/SDL_GetRendererName) | [:heavy_check_mark:](sdl/methods.go#L2912) | [:x:](sdl/sdl_functions_js.go#L14510) |
| [SDL_GetRendererProperties](https://wiki.libsdl.org/SDL3/SDL_GetRendererProperties) | [:heavy_check_mark:](sdl/methods.go#L2923) | [:x:](sdl/sdl_functions_js.go#L14526) |
| [SDL_GetRenderOutputSize](https://wiki.libsdl.org/SDL3/SDL_GetRenderOutputSize) | [:heavy_check_mark:](sdl/methods.go#L2929) | [:x:](sdl/sdl_functions_js.go#L14542) |
| [SDL_GetCurrentRenderOutputSize](https://wiki.libsdl.org/SDL3/SDL_GetCurrentRenderOutputSize) | [:heavy_check_mark:](sdl/methods.go#L2940) | [:x:](sdl/sdl_functions_js.go#L14568) |
| [SDL_CreateTexture](https://wiki.libsdl.org/SDL3/SDL_CreateTexture) | [:heavy_check_mark:](sdl/methods.go#L2951) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L14594) |
| [SDL_CreateTextureFromSurface](https://wiki.libsdl.org/SDL3/SDL_CreateTextureFromSurface) | [:heavy_check_mark:](sdl/methods.go#L2962) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L14617) |
| [SDL_CreateTextureWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateTextureWithProperties) | [:heavy_check_mark:](sdl/methods.go#L2973) | [:x:](sdl/sdl_functions_js.go#L14637) |
| [SDL_GetTextureProperties](https://wiki.libsdl.org/SDL3/SDL_GetTextureProperties) | [:heavy_check_mark:](sdl/methods.go#L992) | [:x:](sdl/sdl_functions_js.go#L14658) |
| [SDL_GetRendererFromTexture](https://wiki.libsdl.org/SDL3/SDL_GetRendererFromTexture) | [:heavy_check_mark:](sdl/methods.go#L1003) | [:x:](sdl/sdl_functions_js.go#L14674) |
| [SDL_GetTextureSize](https://wiki.libsdl.org/SDL3/SDL_GetTextureSize) | [:heavy_check_mark:](sdl/methods.go#L1014) | [:x:](sdl/sdl_functions_js.go#L14693) |
| [SDL_SetTexturePalette](https://wiki.libsdl.org/SDL3/SDL_SetTexturePalette) | [:heavy_check_mark:](sdl/methods.go#L1025) | [:question:]() |
| [SDL_GetTexturePalette](https://wiki.libsdl.org/SDL3/SDL_GetTexturePalette) | [:heavy_check_mark:](sdl/methods.go#L1035) | [:question:]() |
| [SDL_SetTextureColorMod](https://wiki.libsdl.org/SDL3/SDL_SetTextureColorMod) | [:heavy_check_mark:](sdl/methods.go#L1041) | [:x:](sdl/sdl_functions_js.go#L14719) |
| [SDL_SetTextureColorModFloat](https://wiki.libsdl.org/SDL3/SDL_SetTextureColorModFloat) | [:heavy_check_mark:](sdl/methods.go#L1051) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L14741) |
| [SDL_GetTextureColorMod](https://wiki.libsdl.org/SDL3/SDL_GetTextureColorMod) | [:heavy_check_mark:](sdl/methods.go#L1061) | [:x:](sdl/sdl_functions_js.go#L14757) |
| [SDL_GetTextureColorModFloat](https://wiki.libsdl.org/SDL3/SDL_GetTextureColorModFloat) | [:heavy_check_mark:](sdl/methods.go#L1072) | [:x:](sdl/sdl_functions_js.go#L14788) |
| [SDL_SetTextureAlphaMod](https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1083) | [:x:](sdl/sdl_functions_js.go#L14819) |
| [SDL_SetTextureAlphaModFloat](https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaModFloat) | [:heavy_check_mark:](sdl/methods.go#L1093) | [:x:](sdl/sdl_functions_js.go#L14837) |
| [SDL_GetTextureAlphaMod](https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1103) | [:x:](sdl/sdl_functions_js.go#L14855) |
| [SDL_GetTextureAlphaModFloat](https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaModFloat) | [:heavy_check_mark:](sdl/methods.go#L1114) | [:x:](sdl/sdl_functions_js.go#L14876) |
| [SDL_SetTextureBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetTextureBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1125) | [:x:](sdl/sdl_functions_js.go#L14897) |
| [SDL_GetTextureBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetTextureBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1135) | [:x:](sdl/sdl_functions_js.go#L14915) |
| [SDL_SetTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_SetTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L1146) | [:x:](sdl/sdl_functions_js.go#L14936) |
| [SDL_GetTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_GetTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L1156) | [:x:](sdl/sdl_functions_js.go#L14954) |
| [SDL_UpdateTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateTexture) | [:heavy_check_mark:](sdl/methods.go#L1167) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L14975) |
| [SDL_UpdateYUVTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateYUVTexture) | [:heavy_check_mark:](sdl/methods.go#L1179) | [:x:](sdl/sdl_functions_js.go#L14999) |
| [SDL_UpdateNVTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateNVTexture) | [:heavy_check_mark:](sdl/methods.go#L1193) | [:x:](sdl/sdl_functions_js.go#L15041) |
| [SDL_LockTexture](https://wiki.libsdl.org/SDL3/SDL_LockTexture) | [:heavy_check_mark:](sdl/methods.go#L1207) | [:x:](sdl/sdl_functions_js.go#L15076) |
| [SDL_LockTextureToSurface](https://wiki.libsdl.org/SDL3/SDL_LockTextureToSurface) | [:heavy_check_mark:](sdl/methods.go#L1220) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15107) |
| [SDL_UnlockTexture](https://wiki.libsdl.org/SDL3/SDL_UnlockTexture) | [:heavy_check_mark:](sdl/methods.go#L1230) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15129) |
| [SDL_SetRenderTarget](https://wiki.libsdl.org/SDL3/SDL_SetRenderTarget) | [:heavy_check_mark:](sdl/methods.go#L2984) | [:x:](sdl/sdl_functions_js.go#L15140) |
| [SDL_GetRenderTarget](https://wiki.libsdl.org/SDL3/SDL_GetRenderTarget) | [:heavy_check_mark:](sdl/methods.go#L2994) | [:x:](sdl/sdl_functions_js.go#L15161) |
| [SDL_SetRenderLogicalPresentation](https://wiki.libsdl.org/SDL3/SDL_SetRenderLogicalPresentation) | [:heavy_check_mark:](sdl/methods.go#L3000) | [:x:](sdl/sdl_functions_js.go#L15180) |
| [SDL_GetRenderLogicalPresentation](https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentation) | [:heavy_check_mark:](sdl/methods.go#L3010) | [:x:](sdl/sdl_functions_js.go#L15202) |
| [SDL_GetRenderLogicalPresentationRect](https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentationRect) | [:heavy_check_mark:](sdl/methods.go#L3022) | [:x:](sdl/sdl_functions_js.go#L15233) |
| [SDL_RenderCoordinatesFromWindow](https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesFromWindow) | [:heavy_check_mark:](sdl/methods.go#L3034) | [:x:](sdl/sdl_functions_js.go#L15254) |
| [SDL_RenderCoordinatesToWindow](https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesToWindow) | [:heavy_check_mark:](sdl/methods.go#L3046) | [:x:](sdl/sdl_functions_js.go#L15284) |
| [SDL_ConvertEventToRenderCoordinates](https://wiki.libsdl.org/SDL3/SDL_ConvertEventToRenderCoordinates) | [:heavy_check_mark:](sdl/methods.go#L3058) | [:x:](sdl/sdl_functions_js.go#L15314) |
| [SDL_SetRenderViewport](https://wiki.libsdl.org/SDL3/SDL_SetRenderViewport) | [:heavy_check_mark:](sdl/methods.go#L3068) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15335) |
| [SDL_GetRenderViewport](https://wiki.libsdl.org/SDL3/SDL_GetRenderViewport) | [:heavy_check_mark:](sdl/methods.go#L3078) | [:x:](sdl/sdl_functions_js.go#L15352) |
| [SDL_RenderViewportSet](https://wiki.libsdl.org/SDL3/SDL_RenderViewportSet) | [:heavy_check_mark:](sdl/methods.go#L3090) | [:x:](sdl/sdl_functions_js.go#L15373) |
| [SDL_GetRenderSafeArea](https://wiki.libsdl.org/SDL3/SDL_GetRenderSafeArea) | [:heavy_check_mark:](sdl/methods.go#L3096) | [:x:](sdl/sdl_functions_js.go#L15389) |
| [SDL_SetRenderClipRect](https://wiki.libsdl.org/SDL3/SDL_SetRenderClipRect) | [:heavy_check_mark:](sdl/methods.go#L3108) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15410) |
| [SDL_GetRenderClipRect](https://wiki.libsdl.org/SDL3/SDL_GetRenderClipRect) | [:heavy_check_mark:](sdl/methods.go#L3118) | [:x:](sdl/sdl_functions_js.go#L15427) |
| [SDL_RenderClipEnabled](https://wiki.libsdl.org/SDL3/SDL_RenderClipEnabled) | [:heavy_check_mark:](sdl/methods.go#L3130) | [:x:](sdl/sdl_functions_js.go#L15448) |
| [SDL_SetRenderScale](https://wiki.libsdl.org/SDL3/SDL_SetRenderScale) | [:heavy_check_mark:](sdl/methods.go#L3140) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15464) |
| [SDL_GetRenderScale](https://wiki.libsdl.org/SDL3/SDL_GetRenderScale) | [:heavy_check_mark:](sdl/methods.go#L3150) | [:x:](sdl/sdl_functions_js.go#L15481) |
| [SDL_SetRenderDrawColor](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColor) | [:heavy_check_mark:](sdl/methods.go#L3162) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15507) |
| [SDL_SetRenderDrawColorFloat](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColorFloat) | [:heavy_check_mark:](sdl/methods.go#L3172) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15529) |
| [SDL_GetRenderDrawColor](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColor) | [:heavy_check_mark:](sdl/methods.go#L3182) | [:x:](sdl/sdl_functions_js.go#L15546) |
| [SDL_GetRenderDrawColorFloat](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColorFloat) | [:heavy_check_mark:](sdl/methods.go#L3194) | [:x:](sdl/sdl_functions_js.go#L15582) |
| [SDL_SetRenderColorScale](https://wiki.libsdl.org/SDL3/SDL_SetRenderColorScale) | [:heavy_check_mark:](sdl/methods.go#L3206) | [:x:](sdl/sdl_functions_js.go#L15618) |
| [SDL_GetRenderColorScale](https://wiki.libsdl.org/SDL3/SDL_GetRenderColorScale) | [:heavy_check_mark:](sdl/methods.go#L3216) | [:x:](sdl/sdl_functions_js.go#L15636) |
| [SDL_SetRenderDrawBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawBlendMode) | [:heavy_check_mark:](sdl/methods.go#L3228) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15657) |
| [SDL_GetRenderDrawBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawBlendMode) | [:heavy_check_mark:](sdl/methods.go#L3238) | [:x:](sdl/sdl_functions_js.go#L15672) |
| [SDL_RenderClear](https://wiki.libsdl.org/SDL3/SDL_RenderClear) | [:heavy_check_mark:](sdl/methods.go#L3250) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15693) |
| [SDL_RenderPoint](https://wiki.libsdl.org/SDL3/SDL_RenderPoint) | [:heavy_check_mark:](sdl/methods.go#L3259) | [:x:](sdl/sdl_functions_js.go#L15706) |
| [SDL_RenderPoints](https://wiki.libsdl.org/SDL3/SDL_RenderPoints) | [:heavy_check_mark:](sdl/methods.go#L3269) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15726) |
| [SDL_RenderLine](https://wiki.libsdl.org/SDL3/SDL_RenderLine) | [:heavy_check_mark:](sdl/methods.go#L3279) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15744) |
| [SDL_RenderLines](https://wiki.libsdl.org/SDL3/SDL_RenderLines) | [:heavy_check_mark:](sdl/methods.go#L3289) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15761) |
| [SDL_RenderRect](https://wiki.libsdl.org/SDL3/SDL_RenderRect) | [:heavy_check_mark:](sdl/methods.go#L3299) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15779) |
| [SDL_RenderRects](https://wiki.libsdl.org/SDL3/SDL_RenderRects) | [:heavy_check_mark:](sdl/methods.go#L3309) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15796) |
| [SDL_RenderFillRect](https://wiki.libsdl.org/SDL3/SDL_RenderFillRect) | [:heavy_check_mark:](sdl/methods.go#L3319) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15814) |
| [SDL_RenderFillRects](https://wiki.libsdl.org/SDL3/SDL_RenderFillRects) | [:heavy_check_mark:](sdl/methods.go#L3329) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15831) |
| [SDL_RenderTexture](https://wiki.libsdl.org/SDL3/SDL_RenderTexture) | [:heavy_check_mark:](sdl/methods.go#L3339) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15849) |
| [SDL_RenderTextureRotated](https://wiki.libsdl.org/SDL3/SDL_RenderTextureRotated) | [:heavy_check_mark:](sdl/methods.go#L3349) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15873) |
| [SDL_RenderTextureAffine](https://wiki.libsdl.org/SDL3/SDL_RenderTextureAffine) | [:heavy_check_mark:](sdl/methods.go#L3359) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L15903) |
| [SDL_RenderTextureTiled](https://wiki.libsdl.org/SDL3/SDL_RenderTextureTiled) | [:heavy_check_mark:](sdl/methods.go#L3369) | [:x:](sdl/sdl_functions_js.go#L15931) |
| [SDL_RenderTexture9Grid](https://wiki.libsdl.org/SDL3/SDL_RenderTexture9Grid) | [:heavy_check_mark:](sdl/methods.go#L3379) | [:x:](sdl/sdl_functions_js.go#L15964) |
| [SDL_RenderTexture9GridTiled](https://wiki.libsdl.org/SDL3/SDL_RenderTexture9GridTiled) | [:question:]() | [:question:]() |
| [SDL_RenderGeometry](https://wiki.libsdl.org/SDL3/SDL_RenderGeometry) | [:heavy_check_mark:](sdl/methods.go#L3389) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16005) |
| [SDL_RenderGeometryRaw](https://wiki.libsdl.org/SDL3/SDL_RenderGeometryRaw) | [:heavy_check_mark:](sdl/methods.go#L3399) | [:x:](sdl/sdl_functions_js.go#L16032) |
| [SDL_SetRenderTextureAddressMode](https://wiki.libsdl.org/SDL3/SDL_SetRenderTextureAddressMode) | [:heavy_check_mark:](sdl/methods.go#L3420) | [:question:]() |
| [SDL_GetRenderTextureAddressMode](https://wiki.libsdl.org/SDL3/SDL_GetRenderTextureAddressMode) | [:heavy_check_mark:](sdl/methods.go#L3430) | [:question:]() |
| [SDL_RenderReadPixels](https://wiki.libsdl.org/SDL3/SDL_RenderReadPixels) | [:heavy_check_mark:](sdl/methods.go#L3467) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16082) |
| [SDL_RenderPresent](https://wiki.libsdl.org/SDL3/SDL_RenderPresent) | [:heavy_check_mark:](sdl/methods.go#L3478) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16101) |
| [SDL_DestroyTexture](https://wiki.libsdl.org/SDL3/SDL_DestroyTexture) | [:heavy_check_mark:](sdl/methods.go#L1236) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16114) |
| [SDL_DestroyRenderer](https://wiki.libsdl.org/SDL3/SDL_DestroyRenderer) | [:heavy_check_mark:](sdl/methods.go#L3488) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16126) |
| [SDL_FlushRenderer](https://wiki.libsdl.org/SDL3/SDL_FlushRenderer) | [:heavy_check_mark:](sdl/methods.go#L3494) | [:x:](sdl/sdl_functions_js.go#L16138) |
| [SDL_GetRenderMetalLayer](https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalLayer) | [:x:](sdl/methods.go#L3504) | [:x:](sdl/sdl_functions_js.go#L16154) |
| [SDL_GetRenderMetalCommandEncoder](https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalCommandEncoder) | [:x:](sdl/methods.go#L3511) | [:x:](sdl/sdl_functions_js.go#L16170) |
| [SDL_AddVulkanRenderSemaphores](https://wiki.libsdl.org/SDL3/SDL_AddVulkanRenderSemaphores) | [:heavy_check_mark:](sdl/methods.go#L3518) | [:x:](sdl/sdl_functions_js.go#L16186) |
| [SDL_SetRenderVSync](https://wiki.libsdl.org/SDL3/SDL_SetRenderVSync) | [:heavy_check_mark:](sdl/methods.go#L3528) | [:x:](sdl/sdl_functions_js.go#L16208) |
| [SDL_GetRenderVSync](https://wiki.libsdl.org/SDL3/SDL_GetRenderVSync) | [:heavy_check_mark:](sdl/methods.go#L3538) | [:x:](sdl/sdl_functions_js.go#L16226) |
| [SDL_RenderDebugText](https://wiki.libsdl.org/SDL3/SDL_RenderDebugText) | [:heavy_check_mark:](sdl/methods.go#L3560) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L16247) |
| [SDL_RenderDebugTextFormat](https://wiki.libsdl.org/SDL3/SDL_RenderDebugTextFormat) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16267) |
| [SDL_SetDefaultTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_SetDefaultTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L3566) | [:question:]() |
| [SDL_GetDefaultTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_GetDefaultTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L3576) | [:question:]() |
| [SDL_CreateGPURenderState](https://wiki.libsdl.org/SDL3/SDL_CreateGPURenderState) | [:heavy_check_mark:](sdl/methods.go#L3442) | [:question:]() |
| [SDL_SetGPURenderStateFragmentUniforms](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderStateFragmentUniforms) | [:question:]() | [:question:]() |
| [SDL_SetGPURenderState](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderState) | [:heavy_check_mark:](sdl/methods.go#L3455) | [:question:]() |
| [SDL_DestroyGPURenderState](https://wiki.libsdl.org/SDL3/SDL_DestroyGPURenderState) | [:heavy_check_mark:](sdl/methods.go#L3590) | [:question:]() |
</details>
<details>
<summary><h3>SharedObject</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_LoadObject](https://wiki.libsdl.org/SDL3/SDL_LoadObject) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13801) |
| [SDL_LoadFunction](https://wiki.libsdl.org/SDL3/SDL_LoadFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13817) |
| [SDL_UnloadObject](https://wiki.libsdl.org/SDL3/SDL_UnloadObject) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L13835) |
</details>
<details>
<summary><h3>Thread</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateThread](https://wiki.libsdl.org/SDL3/SDL_CreateThread) | [:question:]() | [:question:]() |
| [SDL_CreateThreadWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateThreadWithProperties) | [:question:]() | [:question:]() |
| [SDL_GetThreadName](https://wiki.libsdl.org/SDL3/SDL_GetThreadName) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L872) |
| [SDL_GetCurrentThreadID](https://wiki.libsdl.org/SDL3/SDL_GetCurrentThreadID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L888) |
| [SDL_GetThreadID](https://wiki.libsdl.org/SDL3/SDL_GetThreadID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L899) |
| [SDL_SetCurrentThreadPriority](https://wiki.libsdl.org/SDL3/SDL_SetCurrentThreadPriority) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L915) |
| [SDL_WaitThread](https://wiki.libsdl.org/SDL3/SDL_WaitThread) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L928) |
| [SDL_GetThreadState](https://wiki.libsdl.org/SDL3/SDL_GetThreadState) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L947) |
| [SDL_DetachThread](https://wiki.libsdl.org/SDL3/SDL_DetachThread) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L963) |
| [SDL_GetTLS](https://wiki.libsdl.org/SDL3/SDL_GetTLS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L977) |
| [SDL_SetTLS](https://wiki.libsdl.org/SDL3/SDL_SetTLS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L993) |
| [SDL_CleanupTLS](https://wiki.libsdl.org/SDL3/SDL_CleanupTLS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1013) |
</details>
<details>
<summary><h3>Mutex</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateMutex](https://wiki.libsdl.org/SDL3/SDL_CreateMutex) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1022) |
| [SDL_LockMutex](https://wiki.libsdl.org/SDL3/SDL_LockMutex) | [:heavy_check_mark:](sdl/methods.go#L5753) | [:x:](sdl/sdl_functions_js.go#L1036) |
| [SDL_TryLockMutex](https://wiki.libsdl.org/SDL3/SDL_TryLockMutex) | [:heavy_check_mark:](sdl/methods.go#L5759) | [:x:](sdl/sdl_functions_js.go#L1050) |
| [SDL_UnlockMutex](https://wiki.libsdl.org/SDL3/SDL_UnlockMutex) | [:heavy_check_mark:](sdl/methods.go#L5765) | [:x:](sdl/sdl_functions_js.go#L1066) |
| [SDL_DestroyMutex](https://wiki.libsdl.org/SDL3/SDL_DestroyMutex) | [:heavy_check_mark:](sdl/methods.go#L5771) | [:x:](sdl/sdl_functions_js.go#L1080) |
| [SDL_CreateRWLock](https://wiki.libsdl.org/SDL3/SDL_CreateRWLock) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1094) |
| [SDL_LockRWLockForReading](https://wiki.libsdl.org/SDL3/SDL_LockRWLockForReading) | [:heavy_check_mark:](sdl/methods.go#L1247) | [:x:](sdl/sdl_functions_js.go#L1108) |
| [SDL_LockRWLockForWriting](https://wiki.libsdl.org/SDL3/SDL_LockRWLockForWriting) | [:heavy_check_mark:](sdl/methods.go#L1253) | [:x:](sdl/sdl_functions_js.go#L1122) |
| [SDL_TryLockRWLockForReading](https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForReading) | [:heavy_check_mark:](sdl/methods.go#L1259) | [:x:](sdl/sdl_functions_js.go#L1136) |
| [SDL_TryLockRWLockForWriting](https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForWriting) | [:heavy_check_mark:](sdl/methods.go#L1265) | [:x:](sdl/sdl_functions_js.go#L1152) |
| [SDL_UnlockRWLock](https://wiki.libsdl.org/SDL3/SDL_UnlockRWLock) | [:heavy_check_mark:](sdl/methods.go#L1271) | [:x:](sdl/sdl_functions_js.go#L1168) |
| [SDL_DestroyRWLock](https://wiki.libsdl.org/SDL3/SDL_DestroyRWLock) | [:heavy_check_mark:](sdl/methods.go#L1277) | [:x:](sdl/sdl_functions_js.go#L1182) |
| [SDL_CreateSemaphore](https://wiki.libsdl.org/SDL3/SDL_CreateSemaphore) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1196) |
| [SDL_DestroySemaphore](https://wiki.libsdl.org/SDL3/SDL_DestroySemaphore) | [:heavy_check_mark:](sdl/methods.go#L5922) | [:x:](sdl/sdl_functions_js.go#L1212) |
| [SDL_WaitSemaphore](https://wiki.libsdl.org/SDL3/SDL_WaitSemaphore) | [:heavy_check_mark:](sdl/methods.go#L5928) | [:x:](sdl/sdl_functions_js.go#L1226) |
| [SDL_TryWaitSemaphore](https://wiki.libsdl.org/SDL3/SDL_TryWaitSemaphore) | [:heavy_check_mark:](sdl/methods.go#L5934) | [:x:](sdl/sdl_functions_js.go#L1240) |
| [SDL_WaitSemaphoreTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitSemaphoreTimeout) | [:heavy_check_mark:](sdl/methods.go#L5940) | [:x:](sdl/sdl_functions_js.go#L1256) |
| [SDL_SignalSemaphore](https://wiki.libsdl.org/SDL3/SDL_SignalSemaphore) | [:heavy_check_mark:](sdl/methods.go#L5946) | [:x:](sdl/sdl_functions_js.go#L1274) |
| [SDL_GetSemaphoreValue](https://wiki.libsdl.org/SDL3/SDL_GetSemaphoreValue) | [:heavy_check_mark:](sdl/methods.go#L5952) | [:x:](sdl/sdl_functions_js.go#L1288) |
| [SDL_CreateCondition](https://wiki.libsdl.org/SDL3/SDL_CreateCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1304) |
| [SDL_DestroyCondition](https://wiki.libsdl.org/SDL3/SDL_DestroyCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1318) |
| [SDL_SignalCondition](https://wiki.libsdl.org/SDL3/SDL_SignalCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1332) |
| [SDL_BroadcastCondition](https://wiki.libsdl.org/SDL3/SDL_BroadcastCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1346) |
| [SDL_WaitCondition](https://wiki.libsdl.org/SDL3/SDL_WaitCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1360) |
| [SDL_WaitConditionTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitConditionTimeout) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1379) |
| [SDL_ShouldInit](https://wiki.libsdl.org/SDL3/SDL_ShouldInit) | [:heavy_check_mark:](sdl/methods.go#L3635) | [:x:](sdl/sdl_functions_js.go#L1402) |
| [SDL_ShouldQuit](https://wiki.libsdl.org/SDL3/SDL_ShouldQuit) | [:heavy_check_mark:](sdl/methods.go#L3641) | [:x:](sdl/sdl_functions_js.go#L1418) |
| [SDL_SetInitialized](https://wiki.libsdl.org/SDL3/SDL_SetInitialized) | [:heavy_check_mark:](sdl/methods.go#L3647) | [:x:](sdl/sdl_functions_js.go#L1434) |
</details>
<details>
<summary><h3>Atomic</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_TryLockSpinlock](https://wiki.libsdl.org/SDL3/SDL_TryLockSpinlock) | [:heavy_check_mark:](sdl/methods.go#L572) | [:x:](sdl/sdl_functions_js.go#L254) |
| [SDL_LockSpinlock](https://wiki.libsdl.org/SDL3/SDL_LockSpinlock) | [:heavy_check_mark:](sdl/methods.go#L578) | [:x:](sdl/sdl_functions_js.go#L270) |
| [SDL_UnlockSpinlock](https://wiki.libsdl.org/SDL3/SDL_UnlockSpinlock) | [:heavy_check_mark:](sdl/methods.go#L584) | [:x:](sdl/sdl_functions_js.go#L284) |
| [SDL_MemoryBarrierReleaseFunction](https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierReleaseFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L298) |
| [SDL_MemoryBarrierAcquireFunction](https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierAcquireFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L307) |
| [SDL_CompareAndSwapAtomicInt](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L919) | [:x:](sdl/sdl_functions_js.go#L316) |
| [SDL_SetAtomicInt](https://wiki.libsdl.org/SDL3/SDL_SetAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L925) | [:x:](sdl/sdl_functions_js.go#L336) |
| [SDL_GetAtomicInt](https://wiki.libsdl.org/SDL3/SDL_GetAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L931) | [:x:](sdl/sdl_functions_js.go#L354) |
| [SDL_AddAtomicInt](https://wiki.libsdl.org/SDL3/SDL_AddAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L937) | [:x:](sdl/sdl_functions_js.go#L370) |
| [SDL_CompareAndSwapAtomicU32](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicU32) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L388) |
| [SDL_SetAtomicU32](https://wiki.libsdl.org/SDL3/SDL_SetAtomicU32) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L408) |
| [SDL_GetAtomicU32](https://wiki.libsdl.org/SDL3/SDL_GetAtomicU32) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L426) |
| [SDL_AddAtomicU32](https://wiki.libsdl.org/SDL3/SDL_AddAtomicU32) | [:question:]() | [:question:]() |
| [SDL_CompareAndSwapAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicPointer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L442) |
| [SDL_SetAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_SetAtomicPointer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L462) |
| [SDL_GetAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_GetAtomicPointer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L480) |
</details>
<details>
<summary><h3>Filesystem</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetBasePath](https://wiki.libsdl.org/SDL3/SDL_GetBasePath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10988) |
| [SDL_GetPrefPath](https://wiki.libsdl.org/SDL3/SDL_GetPrefPath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10999) |
| [SDL_GetUserFolder](https://wiki.libsdl.org/SDL3/SDL_GetUserFolder) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11014) |
| [SDL_CreateDirectory](https://wiki.libsdl.org/SDL3/SDL_CreateDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11027) |
| [SDL_EnumerateDirectory](https://wiki.libsdl.org/SDL3/SDL_EnumerateDirectory) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L11040) |
| [SDL_RemovePath](https://wiki.libsdl.org/SDL3/SDL_RemovePath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11057) |
| [SDL_RenamePath](https://wiki.libsdl.org/SDL3/SDL_RenamePath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11070) |
| [SDL_CopyFile](https://wiki.libsdl.org/SDL3/SDL_CopyFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11085) |
| [SDL_GetPathInfo](https://wiki.libsdl.org/SDL3/SDL_GetPathInfo) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11100) |
| [SDL_GlobDirectory](https://wiki.libsdl.org/SDL3/SDL_GlobDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11118) |
| [SDL_GetCurrentDirectory](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11140) |
</details>
<details open>
<summary><h3>IOStream</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_IOFromFile](https://wiki.libsdl.org/SDL3/SDL_IOFromFile) | [:heavy_check_mark:](sdl/functions.go#L370) | [:x:](sdl/sdl_functions_js.go#L1450) |
| [SDL_IOFromMem](https://wiki.libsdl.org/SDL3/SDL_IOFromMem) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1468) |
| [SDL_IOFromConstMem](https://wiki.libsdl.org/SDL3/SDL_IOFromConstMem) | [:heavy_check_mark:](sdl/functions.go#L386) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1486) |
| [SDL_IOFromDynamicMem](https://wiki.libsdl.org/SDL3/SDL_IOFromDynamicMem) | [:heavy_check_mark:](sdl/functions.go#L402) | [:x:](sdl/sdl_functions_js.go#L1506) |
| [SDL_OpenIO](https://wiki.libsdl.org/SDL3/SDL_OpenIO) | [:x:](sdl/methods.go#L23) | [:x:](sdl/sdl_functions_js.go#L1520) |
| [SDL_CloseIO](https://wiki.libsdl.org/SDL3/SDL_CloseIO) | [:heavy_check_mark:](sdl/methods.go#L4938) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1541) |
| [SDL_GetIOProperties](https://wiki.libsdl.org/SDL3/SDL_GetIOProperties) | [:heavy_check_mark:](sdl/methods.go#L4948) | [:x:](sdl/sdl_functions_js.go#L1554) |
| [SDL_GetIOStatus](https://wiki.libsdl.org/SDL3/SDL_GetIOStatus) | [:heavy_check_mark:](sdl/methods.go#L4959) | [:x:](sdl/sdl_functions_js.go#L1570) |
| [SDL_GetIOSize](https://wiki.libsdl.org/SDL3/SDL_GetIOSize) | [:heavy_check_mark:](sdl/methods.go#L4965) | [:x:](sdl/sdl_functions_js.go#L1586) |
| [SDL_SeekIO](https://wiki.libsdl.org/SDL3/SDL_SeekIO) | [:heavy_check_mark:](sdl/methods.go#L4976) | [:x:](sdl/sdl_functions_js.go#L1602) |
| [SDL_TellIO](https://wiki.libsdl.org/SDL3/SDL_TellIO) | [:heavy_check_mark:](sdl/methods.go#L4987) | [:x:](sdl/sdl_functions_js.go#L1622) |
| [SDL_ReadIO](https://wiki.libsdl.org/SDL3/SDL_ReadIO) | [:heavy_check_mark:](sdl/methods.go#L4993) | [:x:](sdl/sdl_functions_js.go#L1638) |
| [SDL_WriteIO](https://wiki.libsdl.org/SDL3/SDL_WriteIO) | [:heavy_check_mark:](sdl/methods.go#L5004) | [:x:](sdl/sdl_functions_js.go#L1658) |
| [SDL_IOprintf](https://wiki.libsdl.org/SDL3/SDL_IOprintf) | [:heavy_check_mark:](sdl/methods.go#L5015) | [:x:](sdl/sdl_functions_js.go#L1678) |
| [SDL_IOvprintf](https://wiki.libsdl.org/SDL3/SDL_IOvprintf) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1696) |
| [SDL_FlushIO](https://wiki.libsdl.org/SDL3/SDL_FlushIO) | [:heavy_check_mark:](sdl/methods.go#L5026) | [:x:](sdl/sdl_functions_js.go#L1716) |
| [SDL_LoadFile_IO](https://wiki.libsdl.org/SDL3/SDL_LoadFile_IO) | [:heavy_check_mark:](sdl/methods.go#L5036) | [:x:](sdl/sdl_functions_js.go#L1732) |
| [SDL_LoadFile](https://wiki.libsdl.org/SDL3/SDL_LoadFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1755) |
| [SDL_SaveFile_IO](https://wiki.libsdl.org/SDL3/SDL_SaveFile_IO) | [:heavy_check_mark:](sdl/methods.go#L5049) | [:x:](sdl/sdl_functions_js.go#L1773) |
| [SDL_SaveFile](https://wiki.libsdl.org/SDL3/SDL_SaveFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1795) |
| [SDL_ReadU8](https://wiki.libsdl.org/SDL3/SDL_ReadU8) | [:heavy_check_mark:](sdl/methods.go#L5060) | [:x:](sdl/sdl_functions_js.go#L1812) |
| [SDL_ReadS8](https://wiki.libsdl.org/SDL3/SDL_ReadS8) | [:heavy_check_mark:](sdl/methods.go#L5072) | [:x:](sdl/sdl_functions_js.go#L1833) |
| [SDL_ReadU16LE](https://wiki.libsdl.org/SDL3/SDL_ReadU16LE) | [:heavy_check_mark:](sdl/methods.go#L5084) | [:x:](sdl/sdl_functions_js.go#L1854) |
| [SDL_ReadS16LE](https://wiki.libsdl.org/SDL3/SDL_ReadS16LE) | [:heavy_check_mark:](sdl/methods.go#L5096) | [:x:](sdl/sdl_functions_js.go#L1875) |
| [SDL_ReadU16BE](https://wiki.libsdl.org/SDL3/SDL_ReadU16BE) | [:heavy_check_mark:](sdl/methods.go#L5108) | [:x:](sdl/sdl_functions_js.go#L1896) |
| [SDL_ReadS16BE](https://wiki.libsdl.org/SDL3/SDL_ReadS16BE) | [:heavy_check_mark:](sdl/methods.go#L5120) | [:x:](sdl/sdl_functions_js.go#L1917) |
| [SDL_ReadU32LE](https://wiki.libsdl.org/SDL3/SDL_ReadU32LE) | [:heavy_check_mark:](sdl/methods.go#L5132) | [:x:](sdl/sdl_functions_js.go#L1938) |
| [SDL_ReadS32LE](https://wiki.libsdl.org/SDL3/SDL_ReadS32LE) | [:heavy_check_mark:](sdl/methods.go#L5144) | [:x:](sdl/sdl_functions_js.go#L1959) |
| [SDL_ReadU32BE](https://wiki.libsdl.org/SDL3/SDL_ReadU32BE) | [:heavy_check_mark:](sdl/methods.go#L5156) | [:x:](sdl/sdl_functions_js.go#L1980) |
| [SDL_ReadS32BE](https://wiki.libsdl.org/SDL3/SDL_ReadS32BE) | [:heavy_check_mark:](sdl/methods.go#L5168) | [:x:](sdl/sdl_functions_js.go#L2001) |
| [SDL_ReadU64LE](https://wiki.libsdl.org/SDL3/SDL_ReadU64LE) | [:heavy_check_mark:](sdl/methods.go#L5180) | [:x:](sdl/sdl_functions_js.go#L2022) |
| [SDL_ReadS64LE](https://wiki.libsdl.org/SDL3/SDL_ReadS64LE) | [:heavy_check_mark:](sdl/methods.go#L5192) | [:x:](sdl/sdl_functions_js.go#L2043) |
| [SDL_ReadU64BE](https://wiki.libsdl.org/SDL3/SDL_ReadU64BE) | [:heavy_check_mark:](sdl/methods.go#L5204) | [:x:](sdl/sdl_functions_js.go#L2064) |
| [SDL_ReadS64BE](https://wiki.libsdl.org/SDL3/SDL_ReadS64BE) | [:heavy_check_mark:](sdl/methods.go#L5216) | [:x:](sdl/sdl_functions_js.go#L2085) |
| [SDL_WriteU8](https://wiki.libsdl.org/SDL3/SDL_WriteU8) | [:heavy_check_mark:](sdl/methods.go#L5228) | [:x:](sdl/sdl_functions_js.go#L2106) |
| [SDL_WriteS8](https://wiki.libsdl.org/SDL3/SDL_WriteS8) | [:heavy_check_mark:](sdl/methods.go#L5238) | [:x:](sdl/sdl_functions_js.go#L2124) |
| [SDL_WriteU16LE](https://wiki.libsdl.org/SDL3/SDL_WriteU16LE) | [:heavy_check_mark:](sdl/methods.go#L5248) | [:x:](sdl/sdl_functions_js.go#L2142) |
| [SDL_WriteS16LE](https://wiki.libsdl.org/SDL3/SDL_WriteS16LE) | [:heavy_check_mark:](sdl/methods.go#L5258) | [:x:](sdl/sdl_functions_js.go#L2160) |
| [SDL_WriteU16BE](https://wiki.libsdl.org/SDL3/SDL_WriteU16BE) | [:heavy_check_mark:](sdl/methods.go#L5268) | [:x:](sdl/sdl_functions_js.go#L2178) |
| [SDL_WriteS16BE](https://wiki.libsdl.org/SDL3/SDL_WriteS16BE) | [:heavy_check_mark:](sdl/methods.go#L5278) | [:x:](sdl/sdl_functions_js.go#L2196) |
| [SDL_WriteU32LE](https://wiki.libsdl.org/SDL3/SDL_WriteU32LE) | [:heavy_check_mark:](sdl/methods.go#L5288) | [:x:](sdl/sdl_functions_js.go#L2214) |
| [SDL_WriteS32LE](https://wiki.libsdl.org/SDL3/SDL_WriteS32LE) | [:heavy_check_mark:](sdl/methods.go#L5298) | [:x:](sdl/sdl_functions_js.go#L2232) |
| [SDL_WriteU32BE](https://wiki.libsdl.org/SDL3/SDL_WriteU32BE) | [:heavy_check_mark:](sdl/methods.go#L5308) | [:x:](sdl/sdl_functions_js.go#L2250) |
| [SDL_WriteS32BE](https://wiki.libsdl.org/SDL3/SDL_WriteS32BE) | [:heavy_check_mark:](sdl/methods.go#L5318) | [:x:](sdl/sdl_functions_js.go#L2268) |
| [SDL_WriteU64LE](https://wiki.libsdl.org/SDL3/SDL_WriteU64LE) | [:heavy_check_mark:](sdl/methods.go#L5328) | [:x:](sdl/sdl_functions_js.go#L2286) |
| [SDL_WriteS64LE](https://wiki.libsdl.org/SDL3/SDL_WriteS64LE) | [:heavy_check_mark:](sdl/methods.go#L5338) | [:x:](sdl/sdl_functions_js.go#L2304) |
| [SDL_WriteU64BE](https://wiki.libsdl.org/SDL3/SDL_WriteU64BE) | [:heavy_check_mark:](sdl/methods.go#L5348) | [:x:](sdl/sdl_functions_js.go#L2322) |
| [SDL_WriteS64BE](https://wiki.libsdl.org/SDL3/SDL_WriteS64BE) | [:heavy_check_mark:](sdl/methods.go#L5358) | [:x:](sdl/sdl_functions_js.go#L2340) |
</details>
<details open>
<summary><h3>AsyncIO</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AsyncIOFromFile](https://wiki.libsdl.org/SDL3/SDL_AsyncIOFromFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L31) |
| [SDL_GetAsyncIOSize](https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOSize) | [:heavy_check_mark:](sdl/methods.go#L3598) | [:x:](sdl/sdl_functions_js.go#L49) |
| [SDL_ReadAsyncIO](https://wiki.libsdl.org/SDL3/SDL_ReadAsyncIO) | [:x:](sdl/methods.go#L3609) | [:x:](sdl/sdl_functions_js.go#L65) |
| [SDL_WriteAsyncIO](https://wiki.libsdl.org/SDL3/SDL_WriteAsyncIO) | [:x:](sdl/methods.go#L3616) | [:x:](sdl/sdl_functions_js.go#L94) |
| [SDL_CloseAsyncIO](https://wiki.libsdl.org/SDL3/SDL_CloseAsyncIO) | [:heavy_check_mark:](sdl/methods.go#L3623) | [:x:](sdl/sdl_functions_js.go#L123) |
| [SDL_CreateAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_CreateAsyncIOQueue) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L148) |
| [SDL_DestroyAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_DestroyAsyncIOQueue) | [:heavy_check_mark:](sdl/methods.go#L1410) | [:x:](sdl/sdl_functions_js.go#L162) |
| [SDL_GetAsyncIOResult](https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOResult) | [:heavy_check_mark:](sdl/methods.go#L1416) | [:x:](sdl/sdl_functions_js.go#L176) |
| [SDL_WaitAsyncIOResult](https://wiki.libsdl.org/SDL3/SDL_WaitAsyncIOResult) | [:heavy_check_mark:](sdl/methods.go#L1426) | [:x:](sdl/sdl_functions_js.go#L197) |
| [SDL_SignalAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_SignalAsyncIOQueue) | [:heavy_check_mark:](sdl/methods.go#L1436) | [:x:](sdl/sdl_functions_js.go#L220) |
| [SDL_LoadFileAsync](https://wiki.libsdl.org/SDL3/SDL_LoadFileAsync) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L234) |
</details>
<details open>
<summary><h3>Storage</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_OpenTitleStorage](https://wiki.libsdl.org/SDL3/SDL_OpenTitleStorage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16289) |
| [SDL_OpenUserStorage](https://wiki.libsdl.org/SDL3/SDL_OpenUserStorage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16307) |
| [SDL_OpenFileStorage](https://wiki.libsdl.org/SDL3/SDL_OpenFileStorage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16327) |
| [SDL_OpenStorage](https://wiki.libsdl.org/SDL3/SDL_OpenStorage) | [:x:](sdl/methods.go#L5744) | [:x:](sdl/sdl_functions_js.go#L16343) |
| [SDL_CloseStorage](https://wiki.libsdl.org/SDL3/SDL_CloseStorage) | [:heavy_check_mark:](sdl/methods.go#L78) | [:x:](sdl/sdl_functions_js.go#L16364) |
| [SDL_StorageReady](https://wiki.libsdl.org/SDL3/SDL_StorageReady) | [:heavy_check_mark:](sdl/methods.go#L88) | [:x:](sdl/sdl_functions_js.go#L16380) |
| [SDL_GetStorageFileSize](https://wiki.libsdl.org/SDL3/SDL_GetStorageFileSize) | [:heavy_check_mark:](sdl/methods.go#L94) | [:x:](sdl/sdl_functions_js.go#L16396) |
| [SDL_ReadStorageFile](https://wiki.libsdl.org/SDL3/SDL_ReadStorageFile) | [:heavy_check_mark:](sdl/methods.go#L106) | [:x:](sdl/sdl_functions_js.go#L16419) |
| [SDL_WriteStorageFile](https://wiki.libsdl.org/SDL3/SDL_WriteStorageFile) | [:heavy_check_mark:](sdl/methods.go#L118) | [:x:](sdl/sdl_functions_js.go#L16441) |
| [SDL_CreateStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_CreateStorageDirectory) | [:heavy_check_mark:](sdl/methods.go#L128) | [:x:](sdl/sdl_functions_js.go#L16463) |
| [SDL_EnumerateStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_EnumerateStorageDirectory) | [:heavy_check_mark:](sdl/methods.go#L138) | [:x:](sdl/sdl_functions_js.go#L16481) |
| [SDL_RemoveStoragePath](https://wiki.libsdl.org/SDL3/SDL_RemoveStoragePath) | [:heavy_check_mark:](sdl/methods.go#L148) | [:x:](sdl/sdl_functions_js.go#L16503) |
| [SDL_RenameStoragePath](https://wiki.libsdl.org/SDL3/SDL_RenameStoragePath) | [:heavy_check_mark:](sdl/methods.go#L158) | [:x:](sdl/sdl_functions_js.go#L16521) |
| [SDL_CopyStorageFile](https://wiki.libsdl.org/SDL3/SDL_CopyStorageFile) | [:heavy_check_mark:](sdl/methods.go#L168) | [:x:](sdl/sdl_functions_js.go#L16541) |
| [SDL_GetStoragePathInfo](https://wiki.libsdl.org/SDL3/SDL_GetStoragePathInfo) | [:heavy_check_mark:](sdl/methods.go#L178) | [:x:](sdl/sdl_functions_js.go#L16561) |
| [SDL_GetStorageSpaceRemaining](https://wiki.libsdl.org/SDL3/SDL_GetStorageSpaceRemaining) | [:heavy_check_mark:](sdl/methods.go#L190) | [:x:](sdl/sdl_functions_js.go#L16584) |
| [SDL_GlobStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_GlobStorageDirectory) | [:heavy_check_mark:](sdl/methods.go#L196) | [:x:](sdl/sdl_functions_js.go#L16600) |
</details>
<details open>
<summary><h3>Pixels</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPixelFormatName](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatName) | [:heavy_check_mark:](sdl/methods.go#L5960) | [:x:](sdl/sdl_functions_js.go#L3387) |
| [SDL_GetMasksForPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetMasksForPixelFormat) | [:x:](sdl/methods.go#L5966) | [:x:](sdl/sdl_functions_js.go#L3400) |
| [SDL_GetPixelFormatForMasks](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatForMasks) | [:heavy_check_mark:](sdl/functions.go#L458) | [:x:](sdl/sdl_functions_js.go#L3438) |
| [SDL_GetPixelFormatDetails](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatDetails) | [:heavy_check_mark:](sdl/methods.go#L5974) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3459) |
| [SDL_CreatePalette](https://wiki.libsdl.org/SDL3/SDL_CreatePalette) | [:heavy_check_mark:](sdl/functions.go#L464) | [:x:](sdl/sdl_functions_js.go#L3471) |
| [SDL_SetPaletteColors](https://wiki.libsdl.org/SDL3/SDL_SetPaletteColors) | [:heavy_check_mark:](sdl/methods.go#L5370) | [:x:](sdl/sdl_functions_js.go#L3487) |
| [SDL_DestroyPalette](https://wiki.libsdl.org/SDL3/SDL_DestroyPalette) | [:heavy_check_mark:](sdl/methods.go#L5380) | [:x:](sdl/sdl_functions_js.go#L3512) |
| [SDL_MapRGB](https://wiki.libsdl.org/SDL3/SDL_MapRGB) | [:heavy_check_mark:](sdl/methods.go#L1444) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3526) |
| [SDL_MapRGBA](https://wiki.libsdl.org/SDL3/SDL_MapRGBA) | [:heavy_check_mark:](sdl/methods.go#L1450) | [:x:](sdl/sdl_functions_js.go#L3553) |
| [SDL_GetRGB](https://wiki.libsdl.org/SDL3/SDL_GetRGB) | [:heavy_check_mark:](sdl/functions.go#L487) | [:x:](sdl/sdl_functions_js.go#L3582) |
| [SDL_GetRGBA](https://wiki.libsdl.org/SDL3/SDL_GetRGBA) | [:heavy_check_mark:](sdl/functions.go#L495) | [:x:](sdl/sdl_functions_js.go#L3618) |
</details>
<details open>
<summary><h3>Surface</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateSurface](https://wiki.libsdl.org/SDL3/SDL_CreateSurface) | [:heavy_check_mark:](sdl/functions.go#L505) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3933) |
| [SDL_CreateSurfaceFrom](https://wiki.libsdl.org/SDL3/SDL_CreateSurfaceFrom) | [:heavy_check_mark:](sdl/functions.go#L516) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3944) |
| [SDL_DestroySurface](https://wiki.libsdl.org/SDL3/SDL_DestroySurface) | [:heavy_check_mark:](sdl/methods.go#L1458) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3958) |
| [SDL_GetSurfaceProperties](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceProperties) | [:heavy_check_mark:](sdl/methods.go#L1467) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3970) |
| [SDL_SetSurfaceColorspace](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorspace) | [:heavy_check_mark:](sdl/methods.go#L1478) | [:x:](sdl/sdl_functions_js.go#L3983) |
| [SDL_GetSurfaceColorspace](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorspace) | [:heavy_check_mark:](sdl/methods.go#L1488) | [:x:](sdl/sdl_functions_js.go#L4001) |
| [SDL_CreateSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_CreateSurfacePalette) | [:heavy_check_mark:](sdl/methods.go#L1494) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4017) |
| [SDL_SetSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_SetSurfacePalette) | [:heavy_check_mark:](sdl/methods.go#L1505) | [:x:](sdl/sdl_functions_js.go#L4030) |
| [SDL_GetSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_GetSurfacePalette) | [:heavy_check_mark:](sdl/methods.go#L1515) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4051) |
| [SDL_AddSurfaceAlternateImage](https://wiki.libsdl.org/SDL3/SDL_AddSurfaceAlternateImage) | [:heavy_check_mark:](sdl/methods.go#L1521) | [:x:](sdl/sdl_functions_js.go#L4064) |
| [SDL_SurfaceHasAlternateImages](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasAlternateImages) | [:heavy_check_mark:](sdl/methods.go#L1531) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4085) |
| [SDL_GetSurfaceImages](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceImages) | [:heavy_check_mark:](sdl/methods.go#L1537) | [:x:](sdl/sdl_functions_js.go#L4098) |
| [SDL_RemoveSurfaceAlternateImages](https://wiki.libsdl.org/SDL3/SDL_RemoveSurfaceAlternateImages) | [:heavy_check_mark:](sdl/methods.go#L1551) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4119) |
| [SDL_LockSurface](https://wiki.libsdl.org/SDL3/SDL_LockSurface) | [:heavy_check_mark:](sdl/methods.go#L1557) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4130) |
| [SDL_UnlockSurface](https://wiki.libsdl.org/SDL3/SDL_UnlockSurface) | [:heavy_check_mark:](sdl/methods.go#L1567) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4143) |
| [SDL_LoadSurface_IO](https://wiki.libsdl.org/SDL3/SDL_LoadSurface_IO) | [:question:]() | [:question:]() |
| [SDL_LoadSurface](https://wiki.libsdl.org/SDL3/SDL_LoadSurface) | [:question:]() | [:question:]() |
| [SDL_LoadBMP_IO](https://wiki.libsdl.org/SDL3/SDL_LoadBMP_IO) | [:heavy_check_mark:](sdl/functions.go#L528) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4154) |
| [SDL_LoadBMP](https://wiki.libsdl.org/SDL3/SDL_LoadBMP) | [:heavy_check_mark:](sdl/functions.go#L539) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4174) |
| [SDL_SaveBMP_IO](https://wiki.libsdl.org/SDL3/SDL_SaveBMP_IO) | [:heavy_check_mark:](sdl/methods.go#L1573) | [:x:](sdl/sdl_functions_js.go#L4186) |
| [SDL_SaveBMP](https://wiki.libsdl.org/SDL3/SDL_SaveBMP) | [:heavy_check_mark:](sdl/methods.go#L1583) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4209) |
| [SDL_LoadPNG_IO](https://wiki.libsdl.org/SDL3/SDL_LoadPNG_IO) | [:question:]() | [:question:]() |
| [SDL_LoadPNG](https://wiki.libsdl.org/SDL3/SDL_LoadPNG) | [:question:]() | [:question:]() |
| [SDL_SavePNG_IO](https://wiki.libsdl.org/SDL3/SDL_SavePNG_IO) | [:question:]() | [:question:]() |
| [SDL_SavePNG](https://wiki.libsdl.org/SDL3/SDL_SavePNG) | [:question:]() | [:question:]() |
| [SDL_SetSurfaceRLE](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceRLE) | [:heavy_check_mark:](sdl/methods.go#L1593) | [:x:](sdl/sdl_functions_js.go#L4226) |
| [SDL_SurfaceHasRLE](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasRLE) | [:heavy_check_mark:](sdl/methods.go#L1603) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4244) |
| [SDL_SetSurfaceColorKey](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorKey) | [:heavy_check_mark:](sdl/methods.go#L1609) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4257) |
| [SDL_SurfaceHasColorKey](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasColorKey) | [:heavy_check_mark:](sdl/methods.go#L1619) | [:x:](sdl/sdl_functions_js.go#L4273) |
| [SDL_GetSurfaceColorKey](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorKey) | [:heavy_check_mark:](sdl/methods.go#L1625) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4289) |
| [SDL_SetSurfaceColorMod](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorMod) | [:heavy_check_mark:](sdl/methods.go#L1636) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4307) |
| [SDL_GetSurfaceColorMod](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorMod) | [:heavy_check_mark:](sdl/methods.go#L1646) | [:x:](sdl/sdl_functions_js.go#L4323) |
| [SDL_SetSurfaceAlphaMod](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1658) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4354) |
| [SDL_GetSurfaceAlphaMod](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1668) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4368) |
| [SDL_SetSurfaceBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1680) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4386) |
| [SDL_GetSurfaceBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1690) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4400) |
| [SDL_SetSurfaceClipRect](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceClipRect) | [:heavy_check_mark:](sdl/methods.go#L1702) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4418) |
| [SDL_GetSurfaceClipRect](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceClipRect) | [:heavy_check_mark:](sdl/methods.go#L1708) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4435) |
| [SDL_FlipSurface](https://wiki.libsdl.org/SDL3/SDL_FlipSurface) | [:heavy_check_mark:](sdl/methods.go#L1720) | [:x:](sdl/sdl_functions_js.go#L4453) |
| [SDL_RotateSurface](https://wiki.libsdl.org/SDL3/SDL_RotateSurface) | [:question:]() | [:question:]() |
| [SDL_DuplicateSurface](https://wiki.libsdl.org/SDL3/SDL_DuplicateSurface) | [:heavy_check_mark:](sdl/methods.go#L1730) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4471) |
| [SDL_ScaleSurface](https://wiki.libsdl.org/SDL3/SDL_ScaleSurface) | [:heavy_check_mark:](sdl/methods.go#L1741) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4484) |
| [SDL_ConvertSurface](https://wiki.libsdl.org/SDL3/SDL_ConvertSurface) | [:heavy_check_mark:](sdl/methods.go#L1752) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4500) |
| [SDL_ConvertSurfaceAndColorspace](https://wiki.libsdl.org/SDL3/SDL_ConvertSurfaceAndColorspace) | [:heavy_check_mark:](sdl/methods.go#L1763) | [:x:](sdl/sdl_functions_js.go#L4517) |
| [SDL_ConvertPixels](https://wiki.libsdl.org/SDL3/SDL_ConvertPixels) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4547) |
| [SDL_ConvertPixelsAndColorspace](https://wiki.libsdl.org/SDL3/SDL_ConvertPixelsAndColorspace) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4574) |
| [SDL_PremultiplyAlpha](https://wiki.libsdl.org/SDL3/SDL_PremultiplyAlpha) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4609) |
| [SDL_PremultiplySurfaceAlpha](https://wiki.libsdl.org/SDL3/SDL_PremultiplySurfaceAlpha) | [:heavy_check_mark:](sdl/methods.go#L1774) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4638) |
| [SDL_ClearSurface](https://wiki.libsdl.org/SDL3/SDL_ClearSurface) | [:heavy_check_mark:](sdl/methods.go#L1784) | [:x:](sdl/sdl_functions_js.go#L4655) |
| [SDL_FillSurfaceRect](https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRect) | [:heavy_check_mark:](sdl/methods.go#L1794) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4679) |
| [SDL_FillSurfaceRects](https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRects) | [:heavy_check_mark:](sdl/methods.go#L1804) | [:x:](sdl/sdl_functions_js.go#L4698) |
| [SDL_BlitSurface](https://wiki.libsdl.org/SDL3/SDL_BlitSurface) | [:heavy_check_mark:](sdl/methods.go#L1814) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4723) |
| [SDL_BlitSurfaceUnchecked](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUnchecked) | [:heavy_check_mark:](sdl/methods.go#L1824) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4747) |
| [SDL_BlitSurfaceScaled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceScaled) | [:heavy_check_mark:](sdl/methods.go#L1834) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4771) |
| [SDL_BlitSurfaceUncheckedScaled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUncheckedScaled) | [:heavy_check_mark:](sdl/methods.go#L1844) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4796) |
| [SDL_StretchSurface](https://wiki.libsdl.org/SDL3/SDL_StretchSurface) | [:question:]() | [:question:]() |
| [SDL_BlitSurfaceTiled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiled) | [:heavy_check_mark:](sdl/methods.go#L1854) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4821) |
| [SDL_BlitSurfaceTiledWithScale](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiledWithScale) | [:heavy_check_mark:](sdl/methods.go#L1864) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4845) |
| [SDL_BlitSurface9Grid](https://wiki.libsdl.org/SDL3/SDL_BlitSurface9Grid) | [:heavy_check_mark:](sdl/methods.go#L1874) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4873) |
| [SDL_MapSurfaceRGB](https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGB) | [:heavy_check_mark:](sdl/methods.go#L1884) | [:x:](sdl/sdl_functions_js.go#L4909) |
| [SDL_MapSurfaceRGBA](https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGBA) | [:heavy_check_mark:](sdl/methods.go#L1890) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4931) |
| [SDL_ReadSurfacePixel](https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixel) | [:heavy_check_mark:](sdl/methods.go#L1896) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4948) |
| [SDL_ReadSurfacePixelFloat](https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixelFloat) | [:heavy_check_mark:](sdl/methods.go#L1908) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4979) |
| [SDL_WriteSurfacePixel](https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixel) | [:heavy_check_mark:](sdl/methods.go#L1920) | [:x:](sdl/sdl_functions_js.go#L5008) |
| [SDL_WriteSurfacePixelFloat](https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixelFloat) | [:heavy_check_mark:](sdl/methods.go#L1930) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5036) |
</details>
<details open>
<summary><h3>BlendMode</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ComposeCustomBlendMode](https://wiki.libsdl.org/SDL3/SDL_ComposeCustomBlendMode) | [:heavy_check_mark:](sdl/functions.go#L556) | [:x:](sdl/sdl_functions_js.go#L3364) |
</details>
<details open>
<summary><h3>Rect</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_RectToFRect](https://wiki.libsdl.org/SDL3/SDL_RectToFRect) | [:question:]() | [:question:]() |
| [SDL_PointInRect](https://wiki.libsdl.org/SDL3/SDL_PointInRect) | [:question:]() | [:question:]() |
| [SDL_RectEmpty](https://wiki.libsdl.org/SDL3/SDL_RectEmpty) | [:question:]() | [:question:]() |
| [SDL_RectsEqual](https://wiki.libsdl.org/SDL3/SDL_RectsEqual) | [:question:]() | [:question:]() |
| [SDL_HasRectIntersection](https://wiki.libsdl.org/SDL3/SDL_HasRectIntersection) | [:heavy_check_mark:](sdl/methods.go#L671) | [:x:](sdl/sdl_functions_js.go#L3659) |
| [SDL_GetRectIntersection](https://wiki.libsdl.org/SDL3/SDL_GetRectIntersection) | [:heavy_check_mark:](sdl/methods.go#L677) | [:x:](sdl/sdl_functions_js.go#L3680) |
| [SDL_GetRectUnion](https://wiki.libsdl.org/SDL3/SDL_GetRectUnion) | [:heavy_check_mark:](sdl/methods.go#L689) | [:x:](sdl/sdl_functions_js.go#L3706) |
| [SDL_GetRectEnclosingPoints](https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPoints) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L3732) |
| [SDL_GetRectAndLineIntersection](https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersection) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3760) |
| [SDL_PointInRectFloat](https://wiki.libsdl.org/SDL3/SDL_PointInRectFloat) | [:question:]() | [:question:]() |
| [SDL_RectEmptyFloat](https://wiki.libsdl.org/SDL3/SDL_RectEmptyFloat) | [:question:]() | [:question:]() |
| [SDL_RectsEqualEpsilon](https://wiki.libsdl.org/SDL3/SDL_RectsEqualEpsilon) | [:question:]() | [:question:]() |
| [SDL_RectsEqualFloat](https://wiki.libsdl.org/SDL3/SDL_RectsEqualFloat) | [:question:]() | [:question:]() |
| [SDL_HasRectIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_HasRectIntersectionFloat) | [:heavy_check_mark:](sdl/methods.go#L4107) | [:x:](sdl/sdl_functions_js.go#L3796) |
| [SDL_GetRectIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectIntersectionFloat) | [:x:](sdl/methods.go#L4113) | [:x:](sdl/sdl_functions_js.go#L3817) |
| [SDL_GetRectUnionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectUnionFloat) | [:x:](sdl/methods.go#L4120) | [:x:](sdl/sdl_functions_js.go#L3843) |
| [SDL_GetRectEnclosingPointsFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPointsFloat) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3869) |
| [SDL_GetRectAndLineIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersectionFloat) | [:x:](sdl/methods.go#L4127) | [:x:](sdl/sdl_functions_js.go#L3897) |
</details>
<details open>
<summary><h3>Camera</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumCameraDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumCameraDrivers) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5055) |
| [SDL_GetCameraDriver](https://wiki.libsdl.org/SDL3/SDL_GetCameraDriver) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5063) |
| [SDL_GetCurrentCameraDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentCameraDriver) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5072) |
| [SDL_GetCameras](https://wiki.libsdl.org/SDL3/SDL_GetCameras) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L5080) |
| [SDL_GetCameraSupportedFormats](https://wiki.libsdl.org/SDL3/SDL_GetCameraSupportedFormats) | [:heavy_check_mark:](sdl/methods.go#L3655) | [:x:](sdl/sdl_functions_js.go#L5096) |
| [SDL_GetCameraName](https://wiki.libsdl.org/SDL3/SDL_GetCameraName) | [:x:](sdl/methods.go#L3670) | [:x:](sdl/sdl_functions_js.go#L5114) |
| [SDL_GetCameraPosition](https://wiki.libsdl.org/SDL3/SDL_GetCameraPosition) | [:heavy_check_mark:](sdl/methods.go#L3677) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5127) |
| [SDL_OpenCamera](https://wiki.libsdl.org/SDL3/SDL_OpenCamera) | [:heavy_check_mark:](sdl/methods.go#L3683) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5136) |
| [SDL_GetCameraPermissionState](https://wiki.libsdl.org/SDL3/SDL_GetCameraPermissionState) | [:heavy_check_mark:](sdl/methods.go#L359) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5150) |
| [SDL_GetCameraID](https://wiki.libsdl.org/SDL3/SDL_GetCameraID) | [:heavy_check_mark:](sdl/methods.go#L365) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5163) |
| [SDL_GetCameraProperties](https://wiki.libsdl.org/SDL3/SDL_GetCameraProperties) | [:heavy_check_mark:](sdl/methods.go#L376) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5176) |
| [SDL_GetCameraFormat](https://wiki.libsdl.org/SDL3/SDL_GetCameraFormat) | [:heavy_check_mark:](sdl/methods.go#L387) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5189) |
| [SDL_AcquireCameraFrame](https://wiki.libsdl.org/SDL3/SDL_AcquireCameraFrame) | [:heavy_check_mark:](sdl/methods.go#L399) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5207) |
| [SDL_ReleaseCameraFrame](https://wiki.libsdl.org/SDL3/SDL_ReleaseCameraFrame) | [:heavy_check_mark:](sdl/methods.go#L412) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5227) |
| [SDL_CloseCamera](https://wiki.libsdl.org/SDL3/SDL_CloseCamera) | [:heavy_check_mark:](sdl/methods.go#L418) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5243) |
</details>
<details open>
<summary><h3>Clipboard</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetClipboardText](https://wiki.libsdl.org/SDL3/SDL_SetClipboardText) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L5254) |
| [SDL_GetClipboardText](https://wiki.libsdl.org/SDL3/SDL_GetClipboardText) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L5267) |
| [SDL_HasClipboardText](https://wiki.libsdl.org/SDL3/SDL_HasClipboardText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5278) |
| [SDL_SetPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_SetPrimarySelectionText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5289) |
| [SDL_GetPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_GetPrimarySelectionText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5302) |
| [SDL_HasPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_HasPrimarySelectionText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5313) |
| [SDL_SetClipboardData](https://wiki.libsdl.org/SDL3/SDL_SetClipboardData) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L5324) |
| [SDL_ClearClipboardData](https://wiki.libsdl.org/SDL3/SDL_ClearClipboardData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5348) |
| [SDL_GetClipboardData](https://wiki.libsdl.org/SDL3/SDL_GetClipboardData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5359) |
| [SDL_HasClipboardData](https://wiki.libsdl.org/SDL3/SDL_HasClipboardData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5377) |
| [SDL_GetClipboardMimeTypes](https://wiki.libsdl.org/SDL3/SDL_GetClipboardMimeTypes) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5390) |
</details>
<details open>
<summary><h3>Dialog</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ShowOpenFileDialog](https://wiki.libsdl.org/SDL3/SDL_ShowOpenFileDialog) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7460) |
| [SDL_ShowSaveFileDialog](https://wiki.libsdl.org/SDL3/SDL_ShowSaveFileDialog) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7489) |
| [SDL_ShowOpenFolderDialog](https://wiki.libsdl.org/SDL3/SDL_ShowOpenFolderDialog) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7516) |
| [SDL_ShowFileDialogWithProperties](https://wiki.libsdl.org/SDL3/SDL_ShowFileDialogWithProperties) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7538) |
</details>
<details open>
<summary><h3>Tray</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateTray](https://wiki.libsdl.org/SDL3/SDL_CreateTray) | [:heavy_check_mark:](sdl/functions.go#L893) | [:question:]() |
| [SDL_SetTrayIcon](https://wiki.libsdl.org/SDL3/SDL_SetTrayIcon) | [:heavy_check_mark:](sdl/methods.go#L426) | [:question:]() |
| [SDL_SetTrayTooltip](https://wiki.libsdl.org/SDL3/SDL_SetTrayTooltip) | [:heavy_check_mark:](sdl/methods.go#L432) | [:question:]() |
| [SDL_CreateTrayMenu](https://wiki.libsdl.org/SDL3/SDL_CreateTrayMenu) | [:heavy_check_mark:](sdl/methods.go#L444) | [:question:]() |
| [SDL_CreateTraySubmenu](https://wiki.libsdl.org/SDL3/SDL_CreateTraySubmenu) | [:heavy_check_mark:](sdl/functions.go#L893) | [:question:]() |
| [SDL_GetTrayMenu](https://wiki.libsdl.org/SDL3/SDL_GetTrayMenu) | [:heavy_check_mark:](sdl/methods.go#L438) | [:question:]() |
| [SDL_GetTraySubmenu](https://wiki.libsdl.org/SDL3/SDL_GetTraySubmenu) | [:heavy_check_mark:](sdl/methods.go#L488) | [:question:]() |
| [SDL_GetTrayEntries](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntries) | [:heavy_check_mark:](sdl/methods.go#L458) | [:question:]() |
| [SDL_RemoveTrayEntry](https://wiki.libsdl.org/SDL3/SDL_RemoveTrayEntry) | [:heavy_check_mark:](sdl/functions.go#L893) | [:question:]() |
| [SDL_InsertTrayEntryAt](https://wiki.libsdl.org/SDL3/SDL_InsertTrayEntryAt) | [:heavy_check_mark:](sdl/methods.go#L468) | [:question:]() |
| [SDL_SetTrayEntryLabel](https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryLabel) | [:heavy_check_mark:](sdl/methods.go#L494) | [:question:]() |
| [SDL_GetTrayEntryLabel](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryLabel) | [:heavy_check_mark:](sdl/methods.go#L500) | [:question:]() |
| [SDL_SetTrayEntryChecked](https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryChecked) | [:heavy_check_mark:](sdl/methods.go#L506) | [:question:]() |
| [SDL_GetTrayEntryChecked](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryChecked) | [:heavy_check_mark:](sdl/methods.go#L512) | [:question:]() |
| [SDL_SetTrayEntryEnabled](https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryEnabled) | [:heavy_check_mark:](sdl/methods.go#L518) | [:question:]() |
| [SDL_GetTrayEntryEnabled](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryEnabled) | [:heavy_check_mark:](sdl/methods.go#L524) | [:question:]() |
| [SDL_SetTrayEntryCallback](https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryCallback) | [:heavy_check_mark:](sdl/methods.go#L530) | [:question:]() |
| [SDL_ClickTrayEntry](https://wiki.libsdl.org/SDL3/SDL_ClickTrayEntry) | [:heavy_check_mark:](sdl/methods.go#L536) | [:question:]() |
| [SDL_DestroyTray](https://wiki.libsdl.org/SDL3/SDL_DestroyTray) | [:heavy_check_mark:](sdl/methods.go#L450) | [:question:]() |
| [SDL_GetTrayEntryParent](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryParent) | [:heavy_check_mark:](sdl/methods.go#L542) | [:question:]() |
| [SDL_GetTrayMenuParentEntry](https://wiki.libsdl.org/SDL3/SDL_GetTrayMenuParentEntry) | [:heavy_check_mark:](sdl/methods.go#L474) | [:question:]() |
| [SDL_GetTrayMenuParentTray](https://wiki.libsdl.org/SDL3/SDL_GetTrayMenuParentTray) | [:heavy_check_mark:](sdl/methods.go#L480) | [:question:]() |
| [SDL_UpdateTrays](https://wiki.libsdl.org/SDL3/SDL_UpdateTrays) | [:heavy_check_mark:](sdl/functions.go#L893) | [:question:]() |
</details>
<details open>
<summary><h3>MessageBox</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ShowMessageBox](https://wiki.libsdl.org/SDL3/SDL_ShowMessageBox) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L14103) |
| [SDL_ShowSimpleMessageBox](https://wiki.libsdl.org/SDL3/SDL_ShowSimpleMessageBox) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L14118) |
</details>
<details open>
<summary><h3>GPU</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GPUSupportsShaderFormats](https://wiki.libsdl.org/SDL3/SDL_GPUSupportsShaderFormats) | [:heavy_check_mark:](sdl/functions.go#L564) | [:x:](sdl/sdl_functions_js.go#L11151) |
| [SDL_GPUSupportsProperties](https://wiki.libsdl.org/SDL3/SDL_GPUSupportsProperties) | [:heavy_check_mark:](sdl/methods.go#L5914) | [:x:](sdl/sdl_functions_js.go#L11166) |
| [SDL_CreateGPUDevice](https://wiki.libsdl.org/SDL3/SDL_CreateGPUDevice) | [:heavy_check_mark:](sdl/functions.go#L574) | [:x:](sdl/sdl_functions_js.go#L11179) |
| [SDL_CreateGPUDeviceWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateGPUDeviceWithProperties) | [:heavy_check_mark:](sdl/functions.go#L589) | [:x:](sdl/sdl_functions_js.go#L11199) |
| [SDL_DestroyGPUDevice](https://wiki.libsdl.org/SDL3/SDL_DestroyGPUDevice) | [:heavy_check_mark:](sdl/methods.go#L1984) | [:x:](sdl/sdl_functions_js.go#L11215) |
| [SDL_GetNumGPUDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumGPUDrivers) | [:heavy_check_mark:](sdl/functions.go#L600) | [:x:](sdl/sdl_functions_js.go#L11229) |
| [SDL_GetGPUDriver](https://wiki.libsdl.org/SDL3/SDL_GetGPUDriver) | [:heavy_check_mark:](sdl/functions.go#L606) | [:x:](sdl/sdl_functions_js.go#L11240) |
| [SDL_GetGPUDeviceDriver](https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceDriver) | [:heavy_check_mark:](sdl/methods.go#L1990) | [:x:](sdl/sdl_functions_js.go#L11253) |
| [SDL_GetGPUShaderFormats](https://wiki.libsdl.org/SDL3/SDL_GetGPUShaderFormats) | [:heavy_check_mark:](sdl/methods.go#L1996) | [:x:](sdl/sdl_functions_js.go#L11269) |
| [SDL_GetGPUDeviceProperties](https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceProperties) | [:question:]() | [:question:]() |
| [SDL_CreateGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_CreateGPUComputePipeline) | [:heavy_check_mark:](sdl/methods.go#L2013) | [:x:](sdl/sdl_functions_js.go#L11285) |
| [SDL_CreateGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_CreateGPUGraphicsPipeline) | [:heavy_check_mark:](sdl/methods.go#L2025) | [:x:](sdl/sdl_functions_js.go#L11309) |
| [SDL_CreateGPUSampler](https://wiki.libsdl.org/SDL3/SDL_CreateGPUSampler) | [:heavy_check_mark:](sdl/methods.go#L2036) | [:x:](sdl/sdl_functions_js.go#L11333) |
| [SDL_CreateGPUShader](https://wiki.libsdl.org/SDL3/SDL_CreateGPUShader) | [:heavy_check_mark:](sdl/methods.go#L2047) | [:x:](sdl/sdl_functions_js.go#L11357) |
| [SDL_CreateGPUTexture](https://wiki.libsdl.org/SDL3/SDL_CreateGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2059) | [:x:](sdl/sdl_functions_js.go#L11381) |
| [SDL_CreateGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_CreateGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2070) | [:x:](sdl/sdl_functions_js.go#L11405) |
| [SDL_CreateGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_CreateGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2081) | [:x:](sdl/sdl_functions_js.go#L11429) |
| [SDL_SetGPUBufferName](https://wiki.libsdl.org/SDL3/SDL_SetGPUBufferName) | [:heavy_check_mark:](sdl/methods.go#L2092) | [:x:](sdl/sdl_functions_js.go#L11453) |
| [SDL_SetGPUTextureName](https://wiki.libsdl.org/SDL3/SDL_SetGPUTextureName) | [:heavy_check_mark:](sdl/methods.go#L2098) | [:x:](sdl/sdl_functions_js.go#L11474) |
| [SDL_InsertGPUDebugLabel](https://wiki.libsdl.org/SDL3/SDL_InsertGPUDebugLabel) | [:heavy_check_mark:](sdl/methods.go#L6015) | [:x:](sdl/sdl_functions_js.go#L11495) |
| [SDL_PushGPUDebugGroup](https://wiki.libsdl.org/SDL3/SDL_PushGPUDebugGroup) | [:heavy_check_mark:](sdl/methods.go#L6021) | [:x:](sdl/sdl_functions_js.go#L11511) |
| [SDL_PopGPUDebugGroup](https://wiki.libsdl.org/SDL3/SDL_PopGPUDebugGroup) | [:heavy_check_mark:](sdl/methods.go#L6027) | [:x:](sdl/sdl_functions_js.go#L11527) |
| [SDL_ReleaseGPUTexture](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2104) | [:x:](sdl/sdl_functions_js.go#L11541) |
| [SDL_ReleaseGPUSampler](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUSampler) | [:heavy_check_mark:](sdl/methods.go#L2110) | [:x:](sdl/sdl_functions_js.go#L11560) |
| [SDL_ReleaseGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2116) | [:x:](sdl/sdl_functions_js.go#L11579) |
| [SDL_ReleaseGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2122) | [:x:](sdl/sdl_functions_js.go#L11598) |
| [SDL_ReleaseGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUComputePipeline) | [:heavy_check_mark:](sdl/methods.go#L2128) | [:x:](sdl/sdl_functions_js.go#L11617) |
| [SDL_ReleaseGPUShader](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUShader) | [:heavy_check_mark:](sdl/methods.go#L2134) | [:x:](sdl/sdl_functions_js.go#L11636) |
| [SDL_ReleaseGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUGraphicsPipeline) | [:heavy_check_mark:](sdl/methods.go#L2140) | [:x:](sdl/sdl_functions_js.go#L11655) |
| [SDL_AcquireGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_AcquireGPUCommandBuffer) | [:heavy_check_mark:](sdl/methods.go#L2146) | [:x:](sdl/sdl_functions_js.go#L11674) |
| [SDL_PushGPUVertexUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUVertexUniformData) | [:heavy_check_mark:](sdl/methods.go#L6033) | [:x:](sdl/sdl_functions_js.go#L11693) |
| [SDL_PushGPUFragmentUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUFragmentUniformData) | [:heavy_check_mark:](sdl/methods.go#L6039) | [:x:](sdl/sdl_functions_js.go#L11713) |
| [SDL_PushGPUComputeUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUComputeUniformData) | [:heavy_check_mark:](sdl/methods.go#L6045) | [:x:](sdl/sdl_functions_js.go#L11733) |
| [SDL_BeginGPURenderPass](https://wiki.libsdl.org/SDL3/SDL_BeginGPURenderPass) | [:heavy_check_mark:](sdl/methods.go#L6051) | [:x:](sdl/sdl_functions_js.go#L11753) |
| [SDL_BindGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_BindGPUGraphicsPipeline) | [:heavy_check_mark:](sdl/methods.go#L1299) | [:x:](sdl/sdl_functions_js.go#L11784) |
| [SDL_SetGPUViewport](https://wiki.libsdl.org/SDL3/SDL_SetGPUViewport) | [:heavy_check_mark:](sdl/methods.go#L1305) | [:x:](sdl/sdl_functions_js.go#L11803) |
| [SDL_SetGPUScissor](https://wiki.libsdl.org/SDL3/SDL_SetGPUScissor) | [:heavy_check_mark:](sdl/methods.go#L1311) | [:x:](sdl/sdl_functions_js.go#L11822) |
| [SDL_SetGPUBlendConstants](https://wiki.libsdl.org/SDL3/SDL_SetGPUBlendConstants) | [:question:]() | [:question:]() |
| [SDL_SetGPUStencilReference](https://wiki.libsdl.org/SDL3/SDL_SetGPUStencilReference) | [:heavy_check_mark:](sdl/methods.go#L1317) | [:x:](sdl/sdl_functions_js.go#L11841) |
| [SDL_BindGPUVertexBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexBuffers) | [:heavy_check_mark:](sdl/methods.go#L1323) | [:x:](sdl/sdl_functions_js.go#L11857) |
| [SDL_BindGPUIndexBuffer](https://wiki.libsdl.org/SDL3/SDL_BindGPUIndexBuffer) | [:heavy_check_mark:](sdl/methods.go#L1330) | [:x:](sdl/sdl_functions_js.go#L11880) |
| [SDL_BindGPUVertexSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexSamplers) | [:heavy_check_mark:](sdl/methods.go#L1336) | [:x:](sdl/sdl_functions_js.go#L11901) |
| [SDL_BindGPUVertexStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageTextures) | [:heavy_check_mark:](sdl/methods.go#L1343) | [:x:](sdl/sdl_functions_js.go#L11924) |
| [SDL_BindGPUVertexStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageBuffers) | [:heavy_check_mark:](sdl/methods.go#L1350) | [:x:](sdl/sdl_functions_js.go#L11947) |
| [SDL_BindGPUFragmentSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentSamplers) | [:heavy_check_mark:](sdl/methods.go#L1357) | [:x:](sdl/sdl_functions_js.go#L11970) |
| [SDL_BindGPUFragmentStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageTextures) | [:heavy_check_mark:](sdl/methods.go#L1364) | [:x:](sdl/sdl_functions_js.go#L11993) |
| [SDL_BindGPUFragmentStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageBuffers) | [:heavy_check_mark:](sdl/methods.go#L1371) | [:x:](sdl/sdl_functions_js.go#L12016) |
| [SDL_DrawGPUIndexedPrimitives](https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitives) | [:heavy_check_mark:](sdl/methods.go#L1378) | [:x:](sdl/sdl_functions_js.go#L12039) |
| [SDL_DrawGPUPrimitives](https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitives) | [:heavy_check_mark:](sdl/methods.go#L1384) | [:x:](sdl/sdl_functions_js.go#L12063) |
| [SDL_DrawGPUPrimitivesIndirect](https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitivesIndirect) | [:heavy_check_mark:](sdl/methods.go#L1390) | [:x:](sdl/sdl_functions_js.go#L12085) |
| [SDL_DrawGPUIndexedPrimitivesIndirect](https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitivesIndirect) | [:heavy_check_mark:](sdl/methods.go#L1396) | [:x:](sdl/sdl_functions_js.go#L12108) |
| [SDL_EndGPURenderPass](https://wiki.libsdl.org/SDL3/SDL_EndGPURenderPass) | [:heavy_check_mark:](sdl/methods.go#L1402) | [:x:](sdl/sdl_functions_js.go#L12131) |
| [SDL_BeginGPUComputePass](https://wiki.libsdl.org/SDL3/SDL_BeginGPUComputePass) | [:heavy_check_mark:](sdl/methods.go#L6057) | [:x:](sdl/sdl_functions_js.go#L12145) |
| [SDL_BindGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputePipeline) | [:heavy_check_mark:](sdl/methods.go#L945) | [:x:](sdl/sdl_functions_js.go#L12178) |
| [SDL_BindGPUComputeSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeSamplers) | [:heavy_check_mark:](sdl/methods.go#L951) | [:x:](sdl/sdl_functions_js.go#L12197) |
| [SDL_BindGPUComputeStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageTextures) | [:heavy_check_mark:](sdl/methods.go#L958) | [:x:](sdl/sdl_functions_js.go#L12220) |
| [SDL_BindGPUComputeStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageBuffers) | [:heavy_check_mark:](sdl/methods.go#L965) | [:x:](sdl/sdl_functions_js.go#L12243) |
| [SDL_DispatchGPUCompute](https://wiki.libsdl.org/SDL3/SDL_DispatchGPUCompute) | [:heavy_check_mark:](sdl/methods.go#L972) | [:x:](sdl/sdl_functions_js.go#L12266) |
| [SDL_DispatchGPUComputeIndirect](https://wiki.libsdl.org/SDL3/SDL_DispatchGPUComputeIndirect) | [:heavy_check_mark:](sdl/methods.go#L978) | [:x:](sdl/sdl_functions_js.go#L12286) |
| [SDL_EndGPUComputePass](https://wiki.libsdl.org/SDL3/SDL_EndGPUComputePass) | [:heavy_check_mark:](sdl/methods.go#L984) | [:x:](sdl/sdl_functions_js.go#L12307) |
| [SDL_MapGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_MapGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2157) | [:x:](sdl/sdl_functions_js.go#L12321) |
| [SDL_UnmapGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_UnmapGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2168) | [:x:](sdl/sdl_functions_js.go#L12344) |
| [SDL_BeginGPUCopyPass](https://wiki.libsdl.org/SDL3/SDL_BeginGPUCopyPass) | [:heavy_check_mark:](sdl/methods.go#L6063) | [:x:](sdl/sdl_functions_js.go#L12363) |
| [SDL_UploadToGPUTexture](https://wiki.libsdl.org/SDL3/SDL_UploadToGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2811) | [:x:](sdl/sdl_functions_js.go#L12382) |
| [SDL_UploadToGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_UploadToGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2817) | [:x:](sdl/sdl_functions_js.go#L12408) |
| [SDL_CopyGPUTextureToTexture](https://wiki.libsdl.org/SDL3/SDL_CopyGPUTextureToTexture) | [:heavy_check_mark:](sdl/methods.go#L2823) | [:x:](sdl/sdl_functions_js.go#L12434) |
| [SDL_CopyGPUBufferToBuffer](https://wiki.libsdl.org/SDL3/SDL_CopyGPUBufferToBuffer) | [:heavy_check_mark:](sdl/methods.go#L2829) | [:x:](sdl/sdl_functions_js.go#L12466) |
| [SDL_DownloadFromGPUTexture](https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2835) | [:x:](sdl/sdl_functions_js.go#L12494) |
| [SDL_DownloadFromGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2841) | [:x:](sdl/sdl_functions_js.go#L12518) |
| [SDL_EndGPUCopyPass](https://wiki.libsdl.org/SDL3/SDL_EndGPUCopyPass) | [:heavy_check_mark:](sdl/methods.go#L2847) | [:x:](sdl/sdl_functions_js.go#L12542) |
| [SDL_GenerateMipmapsForGPUTexture](https://wiki.libsdl.org/SDL3/SDL_GenerateMipmapsForGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L6069) | [:x:](sdl/sdl_functions_js.go#L12556) |
| [SDL_BlitGPUTexture](https://wiki.libsdl.org/SDL3/SDL_BlitGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L6075) | [:x:](sdl/sdl_functions_js.go#L12575) |
| [SDL_WindowSupportsGPUSwapchainComposition](https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUSwapchainComposition) | [:heavy_check_mark:](sdl/methods.go#L2174) | [:x:](sdl/sdl_functions_js.go#L12594) |
| [SDL_WindowSupportsGPUPresentMode](https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUPresentMode) | [:heavy_check_mark:](sdl/methods.go#L2180) | [:x:](sdl/sdl_functions_js.go#L12617) |
| [SDL_ClaimWindowForGPUDevice](https://wiki.libsdl.org/SDL3/SDL_ClaimWindowForGPUDevice) | [:heavy_check_mark:](sdl/methods.go#L2186) | [:x:](sdl/sdl_functions_js.go#L12640) |
| [SDL_ReleaseWindowFromGPUDevice](https://wiki.libsdl.org/SDL3/SDL_ReleaseWindowFromGPUDevice) | [:heavy_check_mark:](sdl/methods.go#L2196) | [:x:](sdl/sdl_functions_js.go#L12661) |
| [SDL_SetGPUSwapchainParameters](https://wiki.libsdl.org/SDL3/SDL_SetGPUSwapchainParameters) | [:heavy_check_mark:](sdl/methods.go#L2202) | [:x:](sdl/sdl_functions_js.go#L12680) |
| [SDL_SetGPUAllowedFramesInFlight](https://wiki.libsdl.org/SDL3/SDL_SetGPUAllowedFramesInFlight) | [:heavy_check_mark:](sdl/methods.go#L2212) | [:x:](sdl/sdl_functions_js.go#L12705) |
| [SDL_GetGPUSwapchainTextureFormat](https://wiki.libsdl.org/SDL3/SDL_GetGPUSwapchainTextureFormat) | [:heavy_check_mark:](sdl/methods.go#L2222) | [:x:](sdl/sdl_functions_js.go#L12723) |
| [SDL_AcquireGPUSwapchainTexture](https://wiki.libsdl.org/SDL3/SDL_AcquireGPUSwapchainTexture) | [:heavy_check_mark:](sdl/methods.go#L6081) | [:x:](sdl/sdl_functions_js.go#L12744) |
| [SDL_WaitForGPUSwapchain](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUSwapchain) | [:heavy_check_mark:](sdl/methods.go#L2228) | [:x:](sdl/sdl_functions_js.go#L12780) |
| [SDL_WaitAndAcquireGPUSwapchainTexture](https://wiki.libsdl.org/SDL3/SDL_WaitAndAcquireGPUSwapchainTexture) | [:heavy_check_mark:](sdl/methods.go#L6093) | [:x:](sdl/sdl_functions_js.go#L12801) |
| [SDL_SubmitGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBuffer) | [:heavy_check_mark:](sdl/methods.go#L6105) | [:x:](sdl/sdl_functions_js.go#L12837) |
| [SDL_SubmitGPUCommandBufferAndAcquireFence](https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBufferAndAcquireFence) | [:heavy_check_mark:](sdl/methods.go#L6115) | [:x:](sdl/sdl_functions_js.go#L12853) |
| [SDL_CancelGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_CancelGPUCommandBuffer) | [:heavy_check_mark:](sdl/methods.go#L6126) | [:x:](sdl/sdl_functions_js.go#L12872) |
| [SDL_WaitForGPUIdle](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUIdle) | [:heavy_check_mark:](sdl/methods.go#L2238) | [:x:](sdl/sdl_functions_js.go#L12888) |
| [SDL_WaitForGPUFences](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUFences) | [:heavy_check_mark:](sdl/methods.go#L2248) | [:x:](sdl/sdl_functions_js.go#L12904) |
| [SDL_QueryGPUFence](https://wiki.libsdl.org/SDL3/SDL_QueryGPUFence) | [:heavy_check_mark:](sdl/methods.go#L2258) | [:x:](sdl/sdl_functions_js.go#L12929) |
| [SDL_ReleaseGPUFence](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUFence) | [:heavy_check_mark:](sdl/methods.go#L2264) | [:x:](sdl/sdl_functions_js.go#L12950) |
| [SDL_GPUTextureFormatTexelBlockSize](https://wiki.libsdl.org/SDL3/SDL_GPUTextureFormatTexelBlockSize) | [:heavy_check_mark:](sdl/methods.go#L558) | [:x:](sdl/sdl_functions_js.go#L12969) |
| [SDL_GPUTextureSupportsFormat](https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsFormat) | [:heavy_check_mark:](sdl/methods.go#L2270) | [:x:](sdl/sdl_functions_js.go#L12982) |
| [SDL_GPUTextureSupportsSampleCount](https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsSampleCount) | [:heavy_check_mark:](sdl/methods.go#L2276) | [:x:](sdl/sdl_functions_js.go#L13004) |
| [SDL_CalculateGPUTextureFormatSize](https://wiki.libsdl.org/SDL3/SDL_CalculateGPUTextureFormatSize) | [:heavy_check_mark:](sdl/methods.go#L564) | [:x:](sdl/sdl_functions_js.go#L13024) |
| [SDL_GetPixelFormatFromGPUTextureFormat](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatFromGPUTextureFormat) | [:question:]() | [:question:]() |
| [SDL_GetGPUTextureFormatFromPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetGPUTextureFormatFromPixelFormat) | [:question:]() | [:question:]() |
| [SDL_GDKSuspendGPU](https://wiki.libsdl.org/SDL3/SDL_GDKSuspendGPU) | [:question:]() | [:question:]() |
| [SDL_GDKResumeGPU](https://wiki.libsdl.org/SDL3/SDL_GDKResumeGPU) | [:question:]() | [:question:]() |
</details>
<details>
<summary><h3>Vulkan</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_Vulkan_LoadLibrary](https://wiki.libsdl.org/SDL3/SDL_Vulkan_LoadLibrary) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L17059) |
| [SDL_Vulkan_GetVkGetInstanceProcAddr](https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetVkGetInstanceProcAddr) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L17067) |
| [SDL_Vulkan_UnloadLibrary](https://wiki.libsdl.org/SDL3/SDL_Vulkan_UnloadLibrary) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L17075) |
| [SDL_Vulkan_GetInstanceExtensions](https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetInstanceExtensions) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L17082) |
| [SDL_Vulkan_CreateSurface](https://wiki.libsdl.org/SDL3/SDL_Vulkan_CreateSurface) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L17090) |
| [SDL_Vulkan_DestroySurface](https://wiki.libsdl.org/SDL3/SDL_Vulkan_DestroySurface) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L17098) |
| [SDL_Vulkan_GetPresentationSupport](https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetPresentationSupport) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L17105) |
</details>
<details>
<summary><h3>Metal</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_Metal_CreateView](https://wiki.libsdl.org/SDL3/SDL_Metal_CreateView) | [:x:](sdl/methods.go#L4884) | [:x:](sdl/sdl_functions_js.go#L14140) |
| [SDL_Metal_DestroyView](https://wiki.libsdl.org/SDL3/SDL_Metal_DestroyView) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14156) |
| [SDL_Metal_GetLayer](https://wiki.libsdl.org/SDL3/SDL_Metal_GetLayer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14167) |
</details>
<details open>
<summary><h3>Power</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetPowerInfo) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7583) |
</details>
<details open>
<summary><h3>Sensor</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetSensors](https://wiki.libsdl.org/SDL3/SDL_GetSensors) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7604) |
| [SDL_GetSensorNameForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorNameForID) | [:heavy_check_mark:](sdl/methods.go#L6217) | [:x:](sdl/sdl_functions_js.go#L7620) |
| [SDL_GetSensorTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorTypeForID) | [:heavy_check_mark:](sdl/methods.go#L6223) | [:x:](sdl/sdl_functions_js.go#L7633) |
| [SDL_GetSensorNonPortableTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableTypeForID) | [:heavy_check_mark:](sdl/methods.go#L6229) | [:x:](sdl/sdl_functions_js.go#L7646) |
| [SDL_OpenSensor](https://wiki.libsdl.org/SDL3/SDL_OpenSensor) | [:heavy_check_mark:](sdl/methods.go#L6235) | [:x:](sdl/sdl_functions_js.go#L7659) |
| [SDL_GetSensorFromID](https://wiki.libsdl.org/SDL3/SDL_GetSensorFromID) | [:heavy_check_mark:](sdl/methods.go#L6246) | [:x:](sdl/sdl_functions_js.go#L7675) |
| [SDL_GetSensorProperties](https://wiki.libsdl.org/SDL3/SDL_GetSensorProperties) | [:heavy_check_mark:](sdl/methods.go#L592) | [:x:](sdl/sdl_functions_js.go#L7691) |
| [SDL_GetSensorName](https://wiki.libsdl.org/SDL3/SDL_GetSensorName) | [:heavy_check_mark:](sdl/methods.go#L603) | [:x:](sdl/sdl_functions_js.go#L7707) |
| [SDL_GetSensorType](https://wiki.libsdl.org/SDL3/SDL_GetSensorType) | [:heavy_check_mark:](sdl/methods.go#L614) | [:x:](sdl/sdl_functions_js.go#L7723) |
| [SDL_GetSensorNonPortableType](https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableType) | [:heavy_check_mark:](sdl/methods.go#L620) | [:x:](sdl/sdl_functions_js.go#L7739) |
| [SDL_GetSensorID](https://wiki.libsdl.org/SDL3/SDL_GetSensorID) | [:heavy_check_mark:](sdl/methods.go#L626) | [:x:](sdl/sdl_functions_js.go#L7755) |
| [SDL_GetSensorData](https://wiki.libsdl.org/SDL3/SDL_GetSensorData) | [:heavy_check_mark:](sdl/methods.go#L637) | [:x:](sdl/sdl_functions_js.go#L7771) |
| [SDL_CloseSensor](https://wiki.libsdl.org/SDL3/SDL_CloseSensor) | [:heavy_check_mark:](sdl/methods.go#L647) | [:x:](sdl/sdl_functions_js.go#L7794) |
| [SDL_UpdateSensors](https://wiki.libsdl.org/SDL3/SDL_UpdateSensors) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L7808) |
</details>
<details>
<summary><h3>Process</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateProcess](https://wiki.libsdl.org/SDL3/SDL_CreateProcess) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14193) |
| [SDL_CreateProcessWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateProcessWithProperties) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L14213) |
| [SDL_GetProcessProperties](https://wiki.libsdl.org/SDL3/SDL_GetProcessProperties) | [:heavy_check_mark:](sdl/methods.go#L6138) | [:x:](sdl/sdl_functions_js.go#L14229) |
| [SDL_ReadProcess](https://wiki.libsdl.org/SDL3/SDL_ReadProcess) | [:heavy_check_mark:](sdl/methods.go#L6149) | [:x:](sdl/sdl_functions_js.go#L14245) |
| [SDL_GetProcessInput](https://wiki.libsdl.org/SDL3/SDL_GetProcessInput) | [:heavy_check_mark:](sdl/methods.go#L6167) | [:x:](sdl/sdl_functions_js.go#L14271) |
| [SDL_GetProcessOutput](https://wiki.libsdl.org/SDL3/SDL_GetProcessOutput) | [:heavy_check_mark:](sdl/methods.go#L6178) | [:x:](sdl/sdl_functions_js.go#L14290) |
| [SDL_KillProcess](https://wiki.libsdl.org/SDL3/SDL_KillProcess) | [:heavy_check_mark:](sdl/methods.go#L6189) | [:x:](sdl/sdl_functions_js.go#L14309) |
| [SDL_WaitProcess](https://wiki.libsdl.org/SDL3/SDL_WaitProcess) | [:heavy_check_mark:](sdl/methods.go#L6199) | [:x:](sdl/sdl_functions_js.go#L14327) |
| [SDL_DestroyProcess](https://wiki.libsdl.org/SDL3/SDL_DestroyProcess) | [:heavy_check_mark:](sdl/methods.go#L6209) | [:x:](sdl/sdl_functions_js.go#L14350) |
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
| [SDL_GetNumLogicalCPUCores](https://wiki.libsdl.org/SDL3/SDL_GetNumLogicalCPUCores) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5406) |
| [SDL_GetCPUCacheLineSize](https://wiki.libsdl.org/SDL3/SDL_GetCPUCacheLineSize) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5417) |
| [SDL_HasAltiVec](https://wiki.libsdl.org/SDL3/SDL_HasAltiVec) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5428) |
| [SDL_HasMMX](https://wiki.libsdl.org/SDL3/SDL_HasMMX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5439) |
| [SDL_HasSSE](https://wiki.libsdl.org/SDL3/SDL_HasSSE) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5450) |
| [SDL_HasSSE2](https://wiki.libsdl.org/SDL3/SDL_HasSSE2) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5461) |
| [SDL_HasSSE3](https://wiki.libsdl.org/SDL3/SDL_HasSSE3) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5472) |
| [SDL_HasSSE41](https://wiki.libsdl.org/SDL3/SDL_HasSSE41) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5483) |
| [SDL_HasSSE42](https://wiki.libsdl.org/SDL3/SDL_HasSSE42) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5494) |
| [SDL_HasAVX](https://wiki.libsdl.org/SDL3/SDL_HasAVX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5505) |
| [SDL_HasAVX2](https://wiki.libsdl.org/SDL3/SDL_HasAVX2) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5516) |
| [SDL_HasAVX512F](https://wiki.libsdl.org/SDL3/SDL_HasAVX512F) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5527) |
| [SDL_HasARMSIMD](https://wiki.libsdl.org/SDL3/SDL_HasARMSIMD) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5538) |
| [SDL_HasNEON](https://wiki.libsdl.org/SDL3/SDL_HasNEON) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5549) |
| [SDL_HasLSX](https://wiki.libsdl.org/SDL3/SDL_HasLSX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5560) |
| [SDL_HasLASX](https://wiki.libsdl.org/SDL3/SDL_HasLASX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5571) |
| [SDL_GetSystemRAM](https://wiki.libsdl.org/SDL3/SDL_GetSystemRAM) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5582) |
| [SDL_GetSIMDAlignment](https://wiki.libsdl.org/SDL3/SDL_GetSIMDAlignment) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5593) |
| [SDL_GetSystemPageSize](https://wiki.libsdl.org/SDL3/SDL_GetSystemPageSize) | [:question:]() | [:question:]() |
</details>
<details>
<summary><h3>Locale</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPreferredLocales](https://wiki.libsdl.org/SDL3/SDL_GetPreferredLocales) | [:heavy_check_mark:](sdl/functions.go#L893) | [:x:](sdl/sdl_functions_js.go#L13849) |
</details>
<details>
<summary><h3>System</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetWindowsMessageHook](https://wiki.libsdl.org/SDL3/SDL_SetWindowsMessageHook) | [:question:]() | [:question:]() |
| [SDL_GetDirect3D9AdapterIndex](https://wiki.libsdl.org/SDL3/SDL_GetDirect3D9AdapterIndex) | [:question:]() | [:question:]() |
| [SDL_GetDXGIOutputInfo](https://wiki.libsdl.org/SDL3/SDL_GetDXGIOutputInfo) | [:question:]() | [:question:]() |
| [SDL_SetX11EventHook](https://wiki.libsdl.org/SDL3/SDL_SetX11EventHook) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16627) |
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
| [SDL_IsTablet](https://wiki.libsdl.org/SDL3/SDL_IsTablet) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16640) |
| [SDL_IsTV](https://wiki.libsdl.org/SDL3/SDL_IsTV) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16651) |
| [SDL_GetSandbox](https://wiki.libsdl.org/SDL3/SDL_GetSandbox) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16662) |
| [SDL_OnApplicationWillTerminate](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillTerminate) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16673) |
| [SDL_OnApplicationDidReceiveMemoryWarning](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidReceiveMemoryWarning) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16682) |
| [SDL_OnApplicationWillEnterBackground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillEnterBackground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16691) |
| [SDL_OnApplicationDidEnterBackground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidEnterBackground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16700) |
| [SDL_OnApplicationWillEnterForeground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillEnterForeground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16709) |
| [SDL_OnApplicationDidEnterForeground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidEnterForeground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L16718) |
| [SDL_OnApplicationDidChangeStatusBarOrientation](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidChangeStatusBarOrientation) | [:question:]() | [:question:]() |
| [SDL_GetGDKTaskQueue](https://wiki.libsdl.org/SDL3/SDL_GetGDKTaskQueue) | [:question:]() | [:question:]() |
| [SDL_GetGDKDefaultUser](https://wiki.libsdl.org/SDL3/SDL_GetGDKDefaultUser) | [:question:]() | [:question:]() |
</details>
<details>
<summary><h3>Misc</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_OpenURL](https://wiki.libsdl.org/SDL3/SDL_OpenURL) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L14180) |
</details>
<details>
<summary><h3>GUID</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GUIDToString](https://wiki.libsdl.org/SDL3/SDL_GUIDToString) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7555) |
| [SDL_StringToGUID](https://wiki.libsdl.org/SDL3/SDL_StringToGUID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7570) |
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
</details>
</details>
<details open>
<summary><h2>TTF</h2></summary>
<details open>
<summary><h3>TTF</h3></summary>

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
| [TTF_SetFontCharSpacing](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontCharSpacing) | [:question:]() | [:question:]() |
| [TTF_GetFontCharSpacing](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontCharSpacing) | [:question:]() | [:question:]() |
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
| [TTF_CreateGPUTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngine) | [:heavy_check_mark:](ttf/functions.go#L144) | [:x:](ttf/ttf_functions_js.go#L1120) |
| [TTF_CreateGPUTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngineWithProperties) | [:heavy_check_mark:](ttf/functions.go#L154) | [:x:](ttf/ttf_functions_js.go#L1137) |
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
| [TTF_Quit](https://wiki.libsdl.org/SDL3_ttf/TTF_Quit) | [:heavy_check_mark:](ttf/functions.go#L164) | [:x:](ttf/ttf_functions_js.go#L1938) |
| [TTF_WasInit](https://wiki.libsdl.org/SDL3_ttf/TTF_WasInit) | [:heavy_check_mark:](ttf/functions.go#L170) | [:x:](ttf/ttf_functions_js.go#L1947) |
</details>
</details>
<details open>
<summary><h2>IMG</h2></summary>
<details open>
<summary><h3>Image</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [IMG_Version](https://wiki.libsdl.org/SDL3_image/IMG_Version) | [:heavy_check_mark:](img/functions.go#L11) | [:x:](img/img_functions_js.go#L13) |
| [IMG_Load](https://wiki.libsdl.org/SDL3_image/IMG_Load) | [:heavy_check_mark:](img/functions.go#L28) | [:x:](img/img_functions_js.go#L45) |
| [IMG_Load_IO](https://wiki.libsdl.org/SDL3_image/IMG_Load_IO) | [:heavy_check_mark:](img/functions.go#L39) | [:x:](img/img_functions_js.go#L59) |
| [IMG_LoadTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTyped_IO) | [:heavy_check_mark:](img/functions.go#L17) | [:x:](img/img_functions_js.go#L24) |
| [IMG_LoadTexture](https://wiki.libsdl.org/SDL3_image/IMG_LoadTexture) | [:heavy_check_mark:](img/functions.go#L50) | [:x:](img/img_functions_js.go#L78) |
| [IMG_LoadTexture_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTexture_IO) | [:heavy_check_mark:](img/functions.go#L61) | [:x:](img/img_functions_js.go#L97) |
| [IMG_LoadTextureTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTextureTyped_IO) | [:heavy_check_mark:](img/functions.go#L72) | [:x:](img/img_functions_js.go#L121) |
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
| [IMG_LoadAVIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAVIF_IO) | [:heavy_check_mark:](img/functions.go#L191) | [:x:](img/img_functions_js.go#L147) |
| [IMG_LoadBMP_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadBMP_IO) | [:heavy_check_mark:](img/functions.go#L224) | [:x:](img/img_functions_js.go#L198) |
| [IMG_LoadCUR_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadCUR_IO) | [:heavy_check_mark:](img/functions.go#L213) | [:x:](img/img_functions_js.go#L181) |
| [IMG_LoadGIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadGIF_IO) | [:heavy_check_mark:](img/functions.go#L235) | [:x:](img/img_functions_js.go#L215) |
| [IMG_LoadICO_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadICO_IO) | [:heavy_check_mark:](img/functions.go#L202) | [:x:](img/img_functions_js.go#L164) |
| [IMG_LoadJPG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadJPG_IO) | [:heavy_check_mark:](img/functions.go#L246) | [:x:](img/img_functions_js.go#L232) |
| [IMG_LoadJXL_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadJXL_IO) | [:heavy_check_mark:](img/functions.go#L257) | [:x:](img/img_functions_js.go#L249) |
| [IMG_LoadLBM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadLBM_IO) | [:heavy_check_mark:](img/functions.go#L268) | [:x:](img/img_functions_js.go#L266) |
| [IMG_LoadPCX_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPCX_IO) | [:heavy_check_mark:](img/functions.go#L279) | [:x:](img/img_functions_js.go#L283) |
| [IMG_LoadPNG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPNG_IO) | [:heavy_check_mark:](img/functions.go#L290) | [:x:](img/img_functions_js.go#L300) |
| [IMG_LoadPNM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadPNM_IO) | [:heavy_check_mark:](img/functions.go#L301) | [:x:](img/img_functions_js.go#L317) |
| [IMG_LoadSVG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadSVG_IO) | [:heavy_check_mark:](img/functions.go#L312) | [:x:](img/img_functions_js.go#L334) |
| [IMG_LoadSizedSVG_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadSizedSVG_IO) | [:heavy_check_mark:](img/functions.go#L400) | [:x:](img/img_functions_js.go#L470) |
| [IMG_LoadQOI_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadQOI_IO) | [:heavy_check_mark:](img/functions.go#L323) | [:x:](img/img_functions_js.go#L351) |
| [IMG_LoadTGA_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTGA_IO) | [:heavy_check_mark:](img/functions.go#L334) | [:x:](img/img_functions_js.go#L368) |
| [IMG_LoadTIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadTIF_IO) | [:heavy_check_mark:](img/functions.go#L345) | [:x:](img/img_functions_js.go#L385) |
| [IMG_LoadWEBP_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadWEBP_IO) | [:heavy_check_mark:](img/functions.go#L389) | [:x:](img/img_functions_js.go#L453) |
| [IMG_LoadXCF_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXCF_IO) | [:heavy_check_mark:](img/functions.go#L356) | [:x:](img/img_functions_js.go#L402) |
| [IMG_LoadXPM_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXPM_IO) | [:heavy_check_mark:](img/functions.go#L367) | [:x:](img/img_functions_js.go#L419) |
| [IMG_LoadXV_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadXV_IO) | [:heavy_check_mark:](img/functions.go#L378) | [:x:](img/img_functions_js.go#L436) |
| [IMG_ReadXPMFromArray](https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArray) | [:heavy_check_mark:](img/functions.go#L411) | [:x:](img/img_functions_js.go#L491) |
| [IMG_ReadXPMFromArrayToRGB888](https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArrayToRGB888) | [:heavy_check_mark:](img/functions.go#L423) | [:x:](img/img_functions_js.go#L508) |
| [IMG_Save](https://wiki.libsdl.org/SDL3_image/IMG_Save) | [:question:]() | [:question:]() |
| [IMG_SaveTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveTyped_IO) | [:question:]() | [:question:]() |
| [IMG_SaveAVIF](https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF) | [:heavy_check_mark:](img/functions.go#L435) | [:x:](img/img_functions_js.go#L525) |
| [IMG_SaveAVIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF_IO) | [:heavy_check_mark:](img/functions.go#L445) | [:x:](img/img_functions_js.go#L545) |
| [IMG_SaveBMP](https://wiki.libsdl.org/SDL3_image/IMG_SaveBMP) | [:question:]() | [:question:]() |
| [IMG_SaveBMP_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveBMP_IO) | [:question:]() | [:question:]() |
| [IMG_SaveCUR](https://wiki.libsdl.org/SDL3_image/IMG_SaveCUR) | [:question:]() | [:question:]() |
| [IMG_SaveCUR_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveCUR_IO) | [:question:]() | [:question:]() |
| [IMG_SaveGIF](https://wiki.libsdl.org/SDL3_image/IMG_SaveGIF) | [:question:]() | [:question:]() |
| [IMG_SaveGIF_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveGIF_IO) | [:question:]() | [:question:]() |
| [IMG_SaveICO](https://wiki.libsdl.org/SDL3_image/IMG_SaveICO) | [:question:]() | [:question:]() |
| [IMG_SaveICO_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveICO_IO) | [:question:]() | [:question:]() |
| [IMG_SaveJPG](https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG) | [:heavy_check_mark:](img/functions.go#L475) | [:x:](img/img_functions_js.go#L611) |
| [IMG_SaveJPG_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG_IO) | [:heavy_check_mark:](img/functions.go#L485) | [:x:](img/img_functions_js.go#L631) |
| [IMG_SavePNG](https://wiki.libsdl.org/SDL3_image/IMG_SavePNG) | [:heavy_check_mark:](img/functions.go#L455) | [:x:](img/img_functions_js.go#L570) |
| [IMG_SavePNG_IO](https://wiki.libsdl.org/SDL3_image/IMG_SavePNG_IO) | [:heavy_check_mark:](img/functions.go#L465) | [:x:](img/img_functions_js.go#L588) |
| [IMG_SaveTGA](https://wiki.libsdl.org/SDL3_image/IMG_SaveTGA) | [:question:]() | [:question:]() |
| [IMG_SaveTGA_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveTGA_IO) | [:question:]() | [:question:]() |
| [IMG_SaveWEBP](https://wiki.libsdl.org/SDL3_image/IMG_SaveWEBP) | [:question:]() | [:question:]() |
| [IMG_SaveWEBP_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveWEBP_IO) | [:question:]() | [:question:]() |
| [IMG_LoadAnimation](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation) | [:heavy_check_mark:](img/functions.go#L495) | [:x:](img/img_functions_js.go#L656) |
| [IMG_LoadAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation_IO) | [:heavy_check_mark:](img/functions.go#L506) | [:x:](img/img_functions_js.go#L670) |
| [IMG_LoadAnimationTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimationTyped_IO) | [:heavy_check_mark:](img/functions.go#L517) | [:x:](img/img_functions_js.go#L689) |
| [IMG_LoadANIAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadANIAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_LoadAPNGAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAPNGAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_LoadAVIFAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadAVIFAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_LoadGIFAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadGIFAnimation_IO) | [:heavy_check_mark:](img/functions.go#L528) | [:x:](img/img_functions_js.go#L724) |
| [IMG_LoadWEBPAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_LoadWEBPAnimation_IO) | [:heavy_check_mark:](img/functions.go#L539) | [:x:](img/img_functions_js.go#L741) |
| [IMG_SaveAnimation](https://wiki.libsdl.org/SDL3_image/IMG_SaveAnimation) | [:question:]() | [:question:]() |
| [IMG_SaveAnimationTyped_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveAnimationTyped_IO) | [:question:]() | [:question:]() |
| [IMG_SaveANIAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveANIAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_SaveAPNGAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveAPNGAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_SaveAVIFAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIFAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_SaveGIFAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveGIFAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_SaveWEBPAnimation_IO](https://wiki.libsdl.org/SDL3_image/IMG_SaveWEBPAnimation_IO) | [:question:]() | [:question:]() |
| [IMG_CreateAnimatedCursor](https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimatedCursor) | [:question:]() | [:question:]() |
| [IMG_FreeAnimation](https://wiki.libsdl.org/SDL3_image/IMG_FreeAnimation) | [:heavy_check_mark:](img/methods.go#L6) | [:x:](img/img_functions_js.go#L710) |
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
| [MIX_Version](https://wiki.libsdl.org/SDL3_mixer/MIX_Version) | [:heavy_check_mark:](mixer/functions.go#L9) | [:x:](mixer/mixer_functions_js.go#L13) |
| [MIX_Init](https://wiki.libsdl.org/SDL3_mixer/MIX_Init) | [:heavy_check_mark:](mixer/functions.go#L15) | [:x:](mixer/mixer_functions_js.go#L24) |
| [MIX_Quit](https://wiki.libsdl.org/SDL3_mixer/MIX_Quit) | [:heavy_check_mark:](mixer/functions.go#L25) | [:x:](mixer/mixer_functions_js.go#L35) |
| [MIX_GetNumAudioDecoders](https://wiki.libsdl.org/SDL3_mixer/MIX_GetNumAudioDecoders) | [:heavy_check_mark:](mixer/functions.go#L31) | [:x:](mixer/mixer_functions_js.go#L44) |
| [MIX_GetAudioDecoder](https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioDecoder) | [:heavy_check_mark:](mixer/functions.go#L37) | [:x:](mixer/mixer_functions_js.go#L55) |
| [MIX_CreateMixerDevice](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateMixerDevice) | [:heavy_check_mark:](mixer/functions.go#L43) | [:x:](mixer/mixer_functions_js.go#L68) |
| [MIX_CreateMixer](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateMixer) | [:heavy_check_mark:](mixer/functions.go#L54) | [:x:](mixer/mixer_functions_js.go#L87) |
| [MIX_DestroyMixer](https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyMixer) | [:heavy_check_mark:](mixer/methods.go#L90) | [:x:](mixer/mixer_functions_js.go#L104) |
| [MIX_GetMixerProperties](https://wiki.libsdl.org/SDL3_mixer/MIX_GetMixerProperties) | [:heavy_check_mark:](mixer/methods.go#L96) | [:x:](mixer/mixer_functions_js.go#L118) |
| [MIX_GetMixerFormat](https://wiki.libsdl.org/SDL3_mixer/MIX_GetMixerFormat) | [:heavy_check_mark:](mixer/methods.go#L102) | [:x:](mixer/mixer_functions_js.go#L134) |
| [MIX_LoadAudio_IO](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadAudio_IO) | [:heavy_check_mark:](mixer/methods.go#L112) | [:x:](mixer/mixer_functions_js.go#L155) |
| [MIX_LoadAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadAudio) | [:heavy_check_mark:](mixer/methods.go#L123) | [:x:](mixer/mixer_functions_js.go#L181) |
| [MIX_LoadAudioWithProperties](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadAudioWithProperties) | [:heavy_check_mark:](mixer/functions.go#L65) | [:x:](mixer/mixer_functions_js.go#L202) |
| [MIX_LoadRawAudio_IO](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadRawAudio_IO) | [:heavy_check_mark:](mixer/methods.go#L134) | [:x:](mixer/mixer_functions_js.go#L216) |
| [MIX_LoadRawAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadRawAudio) | [:heavy_check_mark:](mixer/methods.go#L145) | [:x:](mixer/mixer_functions_js.go#L245) |
| [MIX_LoadRawAudioNoCopy](https://wiki.libsdl.org/SDL3_mixer/MIX_LoadRawAudioNoCopy) | [:heavy_check_mark:](mixer/methods.go#L157) | [:x:](mixer/mixer_functions_js.go#L271) |
| [MIX_CreateSineWaveAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateSineWaveAudio) | [:heavy_check_mark:](mixer/methods.go#L169) | [:x:](mixer/mixer_functions_js.go#L299) |
| [MIX_GetAudioProperties](https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioProperties) | [:heavy_check_mark:](mixer/methods.go#L350) | [:x:](mixer/mixer_functions_js.go#L320) |
| [MIX_GetAudioDuration](https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioDuration) | [:heavy_check_mark:](mixer/methods.go#L361) | [:x:](mixer/mixer_functions_js.go#L336) |
| [MIX_GetAudioFormat](https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioFormat) | [:heavy_check_mark:](mixer/methods.go#L367) | [:x:](mixer/mixer_functions_js.go#L352) |
| [MIX_DestroyAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyAudio) | [:heavy_check_mark:](mixer/methods.go#L377) | [:x:](mixer/mixer_functions_js.go#L373) |
| [MIX_CreateTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateTrack) | [:heavy_check_mark:](mixer/methods.go#L180) | [:x:](mixer/mixer_functions_js.go#L387) |
| [MIX_DestroyTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyTrack) | [:heavy_check_mark:](mixer/methods.go#L397) | [:x:](mixer/mixer_functions_js.go#L404) |
| [MIX_GetTrackProperties](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackProperties) | [:heavy_check_mark:](mixer/methods.go#L403) | [:x:](mixer/mixer_functions_js.go#L418) |
| [MIX_GetTrackMixer](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackMixer) | [:heavy_check_mark:](mixer/methods.go#L414) | [:x:](mixer/mixer_functions_js.go#L434) |
| [MIX_SetTrackAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackAudio) | [:heavy_check_mark:](mixer/methods.go#L425) | [:x:](mixer/mixer_functions_js.go#L451) |
| [MIX_SetTrackAudioStream](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackAudioStream) | [:heavy_check_mark:](mixer/methods.go#L435) | [:x:](mixer/mixer_functions_js.go#L472) |
| [MIX_SetTrackIOStream](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackIOStream) | [:heavy_check_mark:](mixer/methods.go#L445) | [:x:](mixer/mixer_functions_js.go#L493) |
| [MIX_SetTrackRawIOStream](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackRawIOStream) | [:heavy_check_mark:](mixer/methods.go#L455) | [:question:]() |
| [MIX_TagTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_TagTrack) | [:heavy_check_mark:](mixer/methods.go#L465) | [:x:](mixer/mixer_functions_js.go#L516) |
| [MIX_UntagTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_UntagTrack) | [:heavy_check_mark:](mixer/methods.go#L475) | [:x:](mixer/mixer_functions_js.go#L534) |
| [MIX_GetTrackTags](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackTags) | [:heavy_check_mark:](mixer/methods.go#L481) | [:question:]() |
| [MIX_GetTaggedTracks](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTaggedTracks) | [:heavy_check_mark:](mixer/methods.go#L329) | [:question:]() |
| [MIX_SetTrackPlaybackPosition](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackPlaybackPosition) | [:heavy_check_mark:](mixer/methods.go#L502) | [:x:](mixer/mixer_functions_js.go#L550) |
| [MIX_GetTrackPlaybackPosition](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackPlaybackPosition) | [:heavy_check_mark:](mixer/methods.go#L512) | [:x:](mixer/mixer_functions_js.go#L568) |
| [MIX_GetTrackFadeFrames](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackFadeFrames) | [:heavy_check_mark:](mixer/methods.go#L523) | [:question:]() |
| [MIX_GetTrackLoops](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackLoops) | [:question:]() | [:question:]() |
| [MIX_SetTrackLoops](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackLoops) | [:heavy_check_mark:](mixer/methods.go#L535) | [:question:]() |
| [MIX_GetTrackAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackAudio) | [:heavy_check_mark:](mixer/methods.go#L545) | [:x:](mixer/mixer_functions_js.go#L600) |
| [MIX_GetTrackAudioStream](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackAudioStream) | [:heavy_check_mark:](mixer/methods.go#L551) | [:x:](mixer/mixer_functions_js.go#L617) |
| [MIX_GetTrackRemaining](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackRemaining) | [:heavy_check_mark:](mixer/methods.go#L557) | [:x:](mixer/mixer_functions_js.go#L634) |
| [MIX_TrackMSToFrames](https://wiki.libsdl.org/SDL3_mixer/MIX_TrackMSToFrames) | [:heavy_check_mark:](mixer/methods.go#L563) | [:x:](mixer/mixer_functions_js.go#L650) |
| [MIX_TrackFramesToMS](https://wiki.libsdl.org/SDL3_mixer/MIX_TrackFramesToMS) | [:heavy_check_mark:](mixer/methods.go#L569) | [:x:](mixer/mixer_functions_js.go#L668) |
| [MIX_AudioMSToFrames](https://wiki.libsdl.org/SDL3_mixer/MIX_AudioMSToFrames) | [:heavy_check_mark:](mixer/methods.go#L383) | [:x:](mixer/mixer_functions_js.go#L686) |
| [MIX_AudioFramesToMS](https://wiki.libsdl.org/SDL3_mixer/MIX_AudioFramesToMS) | [:heavy_check_mark:](mixer/methods.go#L389) | [:x:](mixer/mixer_functions_js.go#L704) |
| [MIX_MSToFrames](https://wiki.libsdl.org/SDL3_mixer/MIX_MSToFrames) | [:heavy_check_mark:](mixer/functions.go#L76) | [:x:](mixer/mixer_functions_js.go#L722) |
| [MIX_FramesToMS](https://wiki.libsdl.org/SDL3_mixer/MIX_FramesToMS) | [:heavy_check_mark:](mixer/functions.go#L82) | [:x:](mixer/mixer_functions_js.go#L737) |
| [MIX_PlayTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_PlayTrack) | [:heavy_check_mark:](mixer/methods.go#L575) | [:x:](mixer/mixer_functions_js.go#L752) |
| [MIX_PlayTag](https://wiki.libsdl.org/SDL3_mixer/MIX_PlayTag) | [:heavy_check_mark:](mixer/methods.go#L191) | [:x:](mixer/mixer_functions_js.go#L770) |
| [MIX_PlayAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_PlayAudio) | [:heavy_check_mark:](mixer/methods.go#L201) | [:x:](mixer/mixer_functions_js.go#L790) |
| [MIX_StopTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_StopTrack) | [:heavy_check_mark:](mixer/methods.go#L585) | [:x:](mixer/mixer_functions_js.go#L811) |
| [MIX_StopAllTracks](https://wiki.libsdl.org/SDL3_mixer/MIX_StopAllTracks) | [:heavy_check_mark:](mixer/methods.go#L211) | [:x:](mixer/mixer_functions_js.go#L829) |
| [MIX_StopTag](https://wiki.libsdl.org/SDL3_mixer/MIX_StopTag) | [:heavy_check_mark:](mixer/methods.go#L221) | [:x:](mixer/mixer_functions_js.go#L847) |
| [MIX_PauseTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_PauseTrack) | [:heavy_check_mark:](mixer/methods.go#L595) | [:x:](mixer/mixer_functions_js.go#L867) |
| [MIX_PauseAllTracks](https://wiki.libsdl.org/SDL3_mixer/MIX_PauseAllTracks) | [:heavy_check_mark:](mixer/methods.go#L231) | [:x:](mixer/mixer_functions_js.go#L883) |
| [MIX_PauseTag](https://wiki.libsdl.org/SDL3_mixer/MIX_PauseTag) | [:heavy_check_mark:](mixer/methods.go#L241) | [:x:](mixer/mixer_functions_js.go#L899) |
| [MIX_ResumeTrack](https://wiki.libsdl.org/SDL3_mixer/MIX_ResumeTrack) | [:heavy_check_mark:](mixer/methods.go#L605) | [:x:](mixer/mixer_functions_js.go#L917) |
| [MIX_ResumeAllTracks](https://wiki.libsdl.org/SDL3_mixer/MIX_ResumeAllTracks) | [:heavy_check_mark:](mixer/methods.go#L251) | [:x:](mixer/mixer_functions_js.go#L933) |
| [MIX_ResumeTag](https://wiki.libsdl.org/SDL3_mixer/MIX_ResumeTag) | [:heavy_check_mark:](mixer/methods.go#L261) | [:x:](mixer/mixer_functions_js.go#L949) |
| [MIX_TrackPlaying](https://wiki.libsdl.org/SDL3_mixer/MIX_TrackPlaying) | [:heavy_check_mark:](mixer/methods.go#L615) | [:x:](mixer/mixer_functions_js.go#L967) |
| [MIX_TrackPaused](https://wiki.libsdl.org/SDL3_mixer/MIX_TrackPaused) | [:heavy_check_mark:](mixer/methods.go#L621) | [:x:](mixer/mixer_functions_js.go#L983) |
| [MIX_SetMasterGain](https://wiki.libsdl.org/SDL3_mixer/MIX_SetMasterGain) | [:heavy_check_mark:](mixer/methods.go#L271) | [:x:](mixer/mixer_functions_js.go#L999) |
| [MIX_GetMasterGain](https://wiki.libsdl.org/SDL3_mixer/MIX_GetMasterGain) | [:heavy_check_mark:](mixer/methods.go#L281) | [:x:](mixer/mixer_functions_js.go#L1017) |
| [MIX_SetTrackGain](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackGain) | [:heavy_check_mark:](mixer/methods.go#L627) | [:x:](mixer/mixer_functions_js.go#L1033) |
| [MIX_GetTrackGain](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackGain) | [:heavy_check_mark:](mixer/methods.go#L637) | [:x:](mixer/mixer_functions_js.go#L1051) |
| [MIX_SetTagGain](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTagGain) | [:heavy_check_mark:](mixer/methods.go#L287) | [:x:](mixer/mixer_functions_js.go#L1067) |
| [MIX_SetTrackFrequencyRatio](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackFrequencyRatio) | [:heavy_check_mark:](mixer/methods.go#L643) | [:x:](mixer/mixer_functions_js.go#L1087) |
| [MIX_GetTrackFrequencyRatio](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackFrequencyRatio) | [:heavy_check_mark:](mixer/methods.go#L653) | [:x:](mixer/mixer_functions_js.go#L1105) |
| [MIX_SetTrackOutputChannelMap](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackOutputChannelMap) | [:heavy_check_mark:](mixer/methods.go#L659) | [:x:](mixer/mixer_functions_js.go#L1121) |
| [MIX_SetTrackStereo](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackStereo) | [:heavy_check_mark:](mixer/methods.go#L670) | [:x:](mixer/mixer_functions_js.go#L1144) |
| [MIX_SetTrack3DPosition](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrack3DPosition) | [:heavy_check_mark:](mixer/methods.go#L681) | [:x:](mixer/mixer_functions_js.go#L1165) |
| [MIX_GetTrack3DPosition](https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrack3DPosition) | [:heavy_check_mark:](mixer/methods.go#L691) | [:x:](mixer/mixer_functions_js.go#L1186) |
| [MIX_CreateGroup](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateGroup) | [:heavy_check_mark:](mixer/methods.go#L297) | [:x:](mixer/mixer_functions_js.go#L1207) |
| [MIX_DestroyGroup](https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyGroup) | [:heavy_check_mark:](mixer/methods.go#L14) | [:x:](mixer/mixer_functions_js.go#L1224) |
| [MIX_GetGroupProperties](https://wiki.libsdl.org/SDL3_mixer/MIX_GetGroupProperties) | [:heavy_check_mark:](mixer/methods.go#L20) | [:x:](mixer/mixer_functions_js.go#L1238) |
| [MIX_GetGroupMixer](https://wiki.libsdl.org/SDL3_mixer/MIX_GetGroupMixer) | [:heavy_check_mark:](mixer/methods.go#L31) | [:x:](mixer/mixer_functions_js.go#L1254) |
| [MIX_SetTrackGroup](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackGroup) | [:heavy_check_mark:](mixer/methods.go#L703) | [:x:](mixer/mixer_functions_js.go#L1271) |
| [MIX_SetTrackStoppedCallback](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackStoppedCallback) | [:heavy_check_mark:](mixer/methods.go#L713) | [:x:](mixer/mixer_functions_js.go#L1292) |
| [MIX_SetTrackRawCallback](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackRawCallback) | [:heavy_check_mark:](mixer/methods.go#L723) | [:x:](mixer/mixer_functions_js.go#L1312) |
| [MIX_SetTrackCookedCallback](https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackCookedCallback) | [:heavy_check_mark:](mixer/methods.go#L733) | [:x:](mixer/mixer_functions_js.go#L1332) |
| [MIX_SetGroupPostMixCallback](https://wiki.libsdl.org/SDL3_mixer/MIX_SetGroupPostMixCallback) | [:heavy_check_mark:](mixer/methods.go#L37) | [:x:](mixer/mixer_functions_js.go#L1352) |
| [MIX_SetPostMixCallback](https://wiki.libsdl.org/SDL3_mixer/MIX_SetPostMixCallback) | [:heavy_check_mark:](mixer/methods.go#L308) | [:x:](mixer/mixer_functions_js.go#L1372) |
| [MIX_Generate](https://wiki.libsdl.org/SDL3_mixer/MIX_Generate) | [:heavy_check_mark:](mixer/methods.go#L318) | [:x:](mixer/mixer_functions_js.go#L1392) |
| [MIX_CreateAudioDecoder](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateAudioDecoder) | [:heavy_check_mark:](mixer/functions.go#L88) | [:x:](mixer/mixer_functions_js.go#L1412) |
| [MIX_CreateAudioDecoder_IO](https://wiki.libsdl.org/SDL3_mixer/MIX_CreateAudioDecoder_IO) | [:heavy_check_mark:](mixer/functions.go#L99) | [:x:](mixer/mixer_functions_js.go#L1428) |
| [MIX_DestroyAudioDecoder](https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyAudioDecoder) | [:heavy_check_mark:](mixer/methods.go#L49) | [:x:](mixer/mixer_functions_js.go#L1449) |
| [MIX_GetAudioDecoderProperties](https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioDecoderProperties) | [:heavy_check_mark:](mixer/methods.go#L55) | [:x:](mixer/mixer_functions_js.go#L1463) |
| [MIX_GetAudioDecoderFormat](https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioDecoderFormat) | [:heavy_check_mark:](mixer/methods.go#L66) | [:x:](mixer/mixer_functions_js.go#L1479) |
| [MIX_DecodeAudio](https://wiki.libsdl.org/SDL3_mixer/MIX_DecodeAudio) | [:heavy_check_mark:](mixer/methods.go#L76) | [:x:](mixer/mixer_functions_js.go#L1500) |
</details>
</details>

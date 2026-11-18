# API Coverage


This file tracks the functions that have been wrapped.<br>
The following emojis mean (they are clickable and should link to the code implementation):
- :heavy_check_mark: = implemented
- :x: = not implemented yet
- :question: = not planned / don't know about integrating it or not
- :no_entry_sign: = not supported by the web build of SDL (WASM/js only)
<details open>
<summary><h2>SDL</h2></summary>
<details open>
<summary><h3>Init</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_Init](https://wiki.libsdl.org/SDL3/SDL_Init) | [:heavy_check_mark:](sdl/functions.go#L14) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8588) |
| [SDL_InitSubSystem](https://wiki.libsdl.org/SDL3/SDL_InitSubSystem) | [:heavy_check_mark:](sdl/functions.go#L24) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8598) |
| [SDL_QuitSubSystem](https://wiki.libsdl.org/SDL3/SDL_QuitSubSystem) | [:heavy_check_mark:](sdl/functions.go#L34) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8607) |
| [SDL_WasInit](https://wiki.libsdl.org/SDL3/SDL_WasInit) | [:heavy_check_mark:](sdl/functions.go#L40) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8615) |
| [SDL_Quit](https://wiki.libsdl.org/SDL3/SDL_Quit) | [:heavy_check_mark:](sdl/functions.go#L46) | [:question:]() |
| [SDL_IsMainThread](https://wiki.libsdl.org/SDL3/SDL_IsMainThread) | [:heavy_check_mark:](sdl/functions.go#L52) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8628) |
| [SDL_RunOnMainThread](https://wiki.libsdl.org/SDL3/SDL_RunOnMainThread) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11538) |
| [SDL_SetAppMetadata](https://wiki.libsdl.org/SDL3/SDL_SetAppMetadata) | [:heavy_check_mark:](sdl/functions.go#L58) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8635) |
| [SDL_SetAppMetadataProperty](https://wiki.libsdl.org/SDL3/SDL_SetAppMetadataProperty) | [:heavy_check_mark:](sdl/functions.go#L68) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8650) |
| [SDL_GetAppMetadataProperty](https://wiki.libsdl.org/SDL3/SDL_GetAppMetadataProperty) | [:heavy_check_mark:](sdl/functions.go#L78) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8663) |
</details>
<details open>
<summary><h3>Hints</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetHintWithPriority](https://wiki.libsdl.org/SDL3/SDL_SetHintWithPriority) | [:heavy_check_mark:](sdl/functions.go#L86) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8519) |
| [SDL_SetHint](https://wiki.libsdl.org/SDL3/SDL_SetHint) | [:heavy_check_mark:](sdl/functions.go#L96) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8534) |
| [SDL_ResetHint](https://wiki.libsdl.org/SDL3/SDL_ResetHint) | [:heavy_check_mark:](sdl/functions.go#L106) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8547) |
| [SDL_ResetHints](https://wiki.libsdl.org/SDL3/SDL_ResetHints) | [:heavy_check_mark:](sdl/functions.go#L116) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8558) |
| [SDL_GetHint](https://wiki.libsdl.org/SDL3/SDL_GetHint) | [:heavy_check_mark:](sdl/functions.go#L122) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8564) |
| [SDL_GetHintBoolean](https://wiki.libsdl.org/SDL3/SDL_GetHintBoolean) | [:heavy_check_mark:](sdl/functions.go#L128) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8575) |
| [SDL_AddHintCallback](https://wiki.libsdl.org/SDL3/SDL_AddHintCallback) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11507) |
| [SDL_RemoveHintCallback](https://wiki.libsdl.org/SDL3/SDL_RemoveHintCallback) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11523) |
</details>
<details>
<summary><h3>Error</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetError](https://wiki.libsdl.org/SDL3/SDL_SetError) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L363) |
| [SDL_SetErrorV](https://wiki.libsdl.org/SDL3/SDL_SetErrorV) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L374) |
| [SDL_OutOfMemory](https://wiki.libsdl.org/SDL3/SDL_OutOfMemory) | [:heavy_check_mark:](sdl/functions.go#L139) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L388) |
| [SDL_GetError](https://wiki.libsdl.org/SDL3/SDL_GetError) | [:heavy_check_mark:](sdl/init_notjs.go#L32) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L395) |
| [SDL_ClearError](https://wiki.libsdl.org/SDL3/SDL_ClearError) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L403) |
</details>
<details>
<summary><h3>Version</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetVersion](https://wiki.libsdl.org/SDL3/SDL_GetVersion) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L23) |
| [SDL_GetRevision](https://wiki.libsdl.org/SDL3/SDL_GetRevision) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11959) |
</details>
<details open>
<summary><h3>Properties</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetGlobalProperties](https://wiki.libsdl.org/SDL3/SDL_GetGlobalProperties) | [:heavy_check_mark:](sdl/functions.go#L147) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L410) |
| [SDL_CreateProperties](https://wiki.libsdl.org/SDL3/SDL_CreateProperties) | [:heavy_check_mark:](sdl/functions.go#L158) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L417) |
| [SDL_CopyProperties](https://wiki.libsdl.org/SDL3/SDL_CopyProperties) | [:heavy_check_mark:](sdl/functions.go#L169) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L425) |
| [SDL_LockProperties](https://wiki.libsdl.org/SDL3/SDL_LockProperties) | [:heavy_check_mark:](sdl/methods.go#L5789) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L436) |
| [SDL_UnlockProperties](https://wiki.libsdl.org/SDL3/SDL_UnlockProperties) | [:heavy_check_mark:](sdl/methods.go#L5799) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L445) |
| [SDL_SetPointerPropertyWithCleanup](https://wiki.libsdl.org/SDL3/SDL_SetPointerPropertyWithCleanup) | [:x:](sdl/methods.go#L5805) | [:x:](sdl/sdl_functions_js.go#L10932) |
| [SDL_SetPointerProperty](https://wiki.libsdl.org/SDL3/SDL_SetPointerProperty) | [:heavy_check_mark:](sdl/methods.go#L5812) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L453) |
| [SDL_SetStringProperty](https://wiki.libsdl.org/SDL3/SDL_SetStringProperty) | [:heavy_check_mark:](sdl/methods.go#L5822) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L469) |
| [SDL_SetNumberProperty](https://wiki.libsdl.org/SDL3/SDL_SetNumberProperty) | [:heavy_check_mark:](sdl/methods.go#L5832) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L485) |
| [SDL_SetFloatProperty](https://wiki.libsdl.org/SDL3/SDL_SetFloatProperty) | [:heavy_check_mark:](sdl/methods.go#L5842) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L501) |
| [SDL_SetBooleanProperty](https://wiki.libsdl.org/SDL3/SDL_SetBooleanProperty) | [:heavy_check_mark:](sdl/methods.go#L5852) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L516) |
| [SDL_HasProperty](https://wiki.libsdl.org/SDL3/SDL_HasProperty) | [:heavy_check_mark:](sdl/methods.go#L5862) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L531) |
| [SDL_GetPropertyType](https://wiki.libsdl.org/SDL3/SDL_GetPropertyType) | [:heavy_check_mark:](sdl/methods.go#L5868) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L544) |
| [SDL_GetPointerProperty](https://wiki.libsdl.org/SDL3/SDL_GetPointerProperty) | [:heavy_check_mark:](sdl/methods.go#L5874) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L557) |
| [SDL_GetStringProperty](https://wiki.libsdl.org/SDL3/SDL_GetStringProperty) | [:heavy_check_mark:](sdl/methods.go#L5880) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L573) |
| [SDL_GetNumberProperty](https://wiki.libsdl.org/SDL3/SDL_GetNumberProperty) | [:heavy_check_mark:](sdl/methods.go#L5886) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L588) |
| [SDL_GetFloatProperty](https://wiki.libsdl.org/SDL3/SDL_GetFloatProperty) | [:heavy_check_mark:](sdl/methods.go#L5892) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L603) |
| [SDL_GetBooleanProperty](https://wiki.libsdl.org/SDL3/SDL_GetBooleanProperty) | [:heavy_check_mark:](sdl/methods.go#L5898) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L618) |
| [SDL_ClearProperty](https://wiki.libsdl.org/SDL3/SDL_ClearProperty) | [:heavy_check_mark:](sdl/methods.go#L5904) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L633) |
| [SDL_GetNumProperties](https://wiki.libsdl.org/SDL3/SDL_GetNumProperties) | [:question:]() | [:question:]() |
| [SDL_EnumerateProperties](https://wiki.libsdl.org/SDL3/SDL_EnumerateProperties) | [:heavy_check_mark:](sdl/methods.go#L5914) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L646) |
| [SDL_DestroyProperties](https://wiki.libsdl.org/SDL3/SDL_DestroyProperties) | [:heavy_check_mark:](sdl/methods.go#L5920) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L661) |
</details>
<details>
<summary><h3>Log</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetLogPriorities](https://wiki.libsdl.org/SDL3/SDL_SetLogPriorities) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8705) |
| [SDL_SetLogPriority](https://wiki.libsdl.org/SDL3/SDL_SetLogPriority) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8713) |
| [SDL_GetLogPriority](https://wiki.libsdl.org/SDL3/SDL_GetLogPriority) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8723) |
| [SDL_ResetLogPriorities](https://wiki.libsdl.org/SDL3/SDL_ResetLogPriorities) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8732) |
| [SDL_SetLogPriorityPrefix](https://wiki.libsdl.org/SDL3/SDL_SetLogPriorityPrefix) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8738) |
| [SDL_Log](https://wiki.libsdl.org/SDL3/SDL_Log) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8751) |
| [SDL_LogTrace](https://wiki.libsdl.org/SDL3/SDL_LogTrace) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8761) |
| [SDL_LogVerbose](https://wiki.libsdl.org/SDL3/SDL_LogVerbose) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8773) |
| [SDL_LogDebug](https://wiki.libsdl.org/SDL3/SDL_LogDebug) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8785) |
| [SDL_LogInfo](https://wiki.libsdl.org/SDL3/SDL_LogInfo) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8797) |
| [SDL_LogWarn](https://wiki.libsdl.org/SDL3/SDL_LogWarn) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8809) |
| [SDL_LogError](https://wiki.libsdl.org/SDL3/SDL_LogError) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8821) |
| [SDL_LogCritical](https://wiki.libsdl.org/SDL3/SDL_LogCritical) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8833) |
| [SDL_LogMessage](https://wiki.libsdl.org/SDL3/SDL_LogMessage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8845) |
| [SDL_LogMessageV](https://wiki.libsdl.org/SDL3/SDL_LogMessageV) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8859) |
| [SDL_GetDefaultLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_GetDefaultLogOutputFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11566) |
| [SDL_GetLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_GetLogOutputFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8876) |
| [SDL_SetLogOutputFunction](https://wiki.libsdl.org/SDL3/SDL_SetLogOutputFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11574) |
</details>
<details open>
<summary><h3>Video</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumVideoDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumVideoDrivers) | [:heavy_check_mark:](sdl/functions.go#L614) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4045) |
| [SDL_GetVideoDriver](https://wiki.libsdl.org/SDL3/SDL_GetVideoDriver) | [:heavy_check_mark:](sdl/functions.go#L620) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4052) |
| [SDL_GetCurrentVideoDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentVideoDriver) | [:heavy_check_mark:](sdl/functions.go#L626) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4061) |
| [SDL_GetSystemTheme](https://wiki.libsdl.org/SDL3/SDL_GetSystemTheme) | [:heavy_check_mark:](sdl/functions.go#L632) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4068) |
| [SDL_GetDisplays](https://wiki.libsdl.org/SDL3/SDL_GetDisplays) | [:heavy_check_mark:](sdl/functions.go#L638) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4075) |
| [SDL_GetPrimaryDisplay](https://wiki.libsdl.org/SDL3/SDL_GetPrimaryDisplay) | [:heavy_check_mark:](sdl/functions.go#L652) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4087) |
| [SDL_GetDisplayProperties](https://wiki.libsdl.org/SDL3/SDL_GetDisplayProperties) | [:heavy_check_mark:](sdl/methods.go#L3703) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4094) |
| [SDL_GetDisplayName](https://wiki.libsdl.org/SDL3/SDL_GetDisplayName) | [:heavy_check_mark:](sdl/methods.go#L3714) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4103) |
| [SDL_GetDisplayBounds](https://wiki.libsdl.org/SDL3/SDL_GetDisplayBounds) | [:heavy_check_mark:](sdl/methods.go#L3725) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4112) |
| [SDL_GetDisplayUsableBounds](https://wiki.libsdl.org/SDL3/SDL_GetDisplayUsableBounds) | [:heavy_check_mark:](sdl/methods.go#L3737) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4126) |
| [SDL_GetNaturalDisplayOrientation](https://wiki.libsdl.org/SDL3/SDL_GetNaturalDisplayOrientation) | [:heavy_check_mark:](sdl/methods.go#L3749) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4140) |
| [SDL_GetCurrentDisplayOrientation](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayOrientation) | [:heavy_check_mark:](sdl/methods.go#L3755) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4149) |
| [SDL_GetDisplayContentScale](https://wiki.libsdl.org/SDL3/SDL_GetDisplayContentScale) | [:heavy_check_mark:](sdl/methods.go#L3761) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4158) |
| [SDL_GetFullscreenDisplayModes](https://wiki.libsdl.org/SDL3/SDL_GetFullscreenDisplayModes) | [:heavy_check_mark:](sdl/methods.go#L3772) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4167) |
| [SDL_GetClosestFullscreenDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetClosestFullscreenDisplayMode) | [:heavy_check_mark:](sdl/methods.go#L3786) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4181) |
| [SDL_GetDesktopDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetDesktopDisplayMode) | [:heavy_check_mark:](sdl/methods.go#L3798) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4203) |
| [SDL_GetCurrentDisplayMode](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayMode) | [:heavy_check_mark:](sdl/methods.go#L3809) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4212) |
| [SDL_GetDisplayForPoint](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForPoint) | [:heavy_check_mark:](sdl/functions.go#L658) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4221) |
| [SDL_GetDisplayForRect](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForRect) | [:heavy_check_mark:](sdl/functions.go#L664) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4232) |
| [SDL_GetDisplayForWindow](https://wiki.libsdl.org/SDL3/SDL_GetDisplayForWindow) | [:heavy_check_mark:](sdl/functions.go#L670) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4243) |
| [SDL_GetWindowPixelDensity](https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelDensity) | [:heavy_check_mark:](sdl/methods.go#L4143) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4252) |
| [SDL_GetWindowDisplayScale](https://wiki.libsdl.org/SDL3/SDL_GetWindowDisplayScale) | [:heavy_check_mark:](sdl/methods.go#L4154) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4261) |
| [SDL_SetWindowFullscreenMode](https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreenMode) | [:heavy_check_mark:](sdl/methods.go#L4165) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4270) |
| [SDL_GetWindowFullscreenMode](https://wiki.libsdl.org/SDL3/SDL_GetWindowFullscreenMode) | [:heavy_check_mark:](sdl/methods.go#L4175) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4283) |
| [SDL_GetWindowICCProfile](https://wiki.libsdl.org/SDL3/SDL_GetWindowICCProfile) | [:heavy_check_mark:](sdl/methods.go#L4181) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4292) |
| [SDL_GetWindowPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelFormat) | [:heavy_check_mark:](sdl/methods.go#L4195) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4306) |
| [SDL_GetWindows](https://wiki.libsdl.org/SDL3/SDL_GetWindows) | [:heavy_check_mark:](sdl/functions.go#L676) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4319) |
| [SDL_CreateWindow](https://wiki.libsdl.org/SDL3/SDL_CreateWindow) | [:heavy_check_mark:](sdl/functions.go#L690) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4331) |
| [SDL_CreatePopupWindow](https://wiki.libsdl.org/SDL3/SDL_CreatePopupWindow) | [:heavy_check_mark:](sdl/methods.go#L4206) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4348) |
| [SDL_CreateWindowWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateWindowWithProperties) | [:heavy_check_mark:](sdl/functions.go#L701) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4367) |
| [SDL_GetWindowID](https://wiki.libsdl.org/SDL3/SDL_GetWindowID) | [:heavy_check_mark:](sdl/methods.go#L4217) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4376) |
| [SDL_GetWindowFromID](https://wiki.libsdl.org/SDL3/SDL_GetWindowFromID) | [:heavy_check_mark:](sdl/methods.go#L33) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4385) |
| [SDL_GetWindowParent](https://wiki.libsdl.org/SDL3/SDL_GetWindowParent) | [:heavy_check_mark:](sdl/methods.go#L4228) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4394) |
| [SDL_GetWindowProperties](https://wiki.libsdl.org/SDL3/SDL_GetWindowProperties) | [:heavy_check_mark:](sdl/methods.go#L4234) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4403) |
| [SDL_GetWindowFlags](https://wiki.libsdl.org/SDL3/SDL_GetWindowFlags) | [:heavy_check_mark:](sdl/methods.go#L4245) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4412) |
| [SDL_SetWindowTitle](https://wiki.libsdl.org/SDL3/SDL_SetWindowTitle) | [:heavy_check_mark:](sdl/methods.go#L4251) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4421) |
| [SDL_GetWindowTitle](https://wiki.libsdl.org/SDL3/SDL_GetWindowTitle) | [:heavy_check_mark:](sdl/methods.go#L4261) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4434) |
| [SDL_SetWindowIcon](https://wiki.libsdl.org/SDL3/SDL_SetWindowIcon) | [:heavy_check_mark:](sdl/methods.go#L4267) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4443) |
| [SDL_SetWindowPosition](https://wiki.libsdl.org/SDL3/SDL_SetWindowPosition) | [:heavy_check_mark:](sdl/methods.go#L4277) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4454) |
| [SDL_GetWindowPosition](https://wiki.libsdl.org/SDL3/SDL_GetWindowPosition) | [:heavy_check_mark:](sdl/methods.go#L4287) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4467) |
| [SDL_SetWindowSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowSize) | [:heavy_check_mark:](sdl/methods.go#L4299) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4484) |
| [SDL_GetWindowSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowSize) | [:heavy_check_mark:](sdl/methods.go#L4309) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4497) |
| [SDL_GetWindowSafeArea](https://wiki.libsdl.org/SDL3/SDL_GetWindowSafeArea) | [:heavy_check_mark:](sdl/methods.go#L4321) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4519) |
| [SDL_SetWindowAspectRatio](https://wiki.libsdl.org/SDL3/SDL_SetWindowAspectRatio) | [:heavy_check_mark:](sdl/methods.go#L4333) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4533) |
| [SDL_GetWindowAspectRatio](https://wiki.libsdl.org/SDL3/SDL_GetWindowAspectRatio) | [:heavy_check_mark:](sdl/methods.go#L4343) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4546) |
| [SDL_GetWindowBordersSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowBordersSize) | [:heavy_check_mark:](sdl/methods.go#L4355) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4563) |
| [SDL_GetWindowSizeInPixels](https://wiki.libsdl.org/SDL3/SDL_GetWindowSizeInPixels) | [:heavy_check_mark:](sdl/methods.go#L4367) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4586) |
| [SDL_SetWindowMinimumSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowMinimumSize) | [:heavy_check_mark:](sdl/methods.go#L4379) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4603) |
| [SDL_GetWindowMinimumSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowMinimumSize) | [:heavy_check_mark:](sdl/methods.go#L4389) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4616) |
| [SDL_SetWindowMaximumSize](https://wiki.libsdl.org/SDL3/SDL_SetWindowMaximumSize) | [:heavy_check_mark:](sdl/methods.go#L4401) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4633) |
| [SDL_GetWindowMaximumSize](https://wiki.libsdl.org/SDL3/SDL_GetWindowMaximumSize) | [:heavy_check_mark:](sdl/methods.go#L4411) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4646) |
| [SDL_SetWindowBordered](https://wiki.libsdl.org/SDL3/SDL_SetWindowBordered) | [:heavy_check_mark:](sdl/methods.go#L4423) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4663) |
| [SDL_SetWindowResizable](https://wiki.libsdl.org/SDL3/SDL_SetWindowResizable) | [:heavy_check_mark:](sdl/methods.go#L4433) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4674) |
| [SDL_SetWindowAlwaysOnTop](https://wiki.libsdl.org/SDL3/SDL_SetWindowAlwaysOnTop) | [:heavy_check_mark:](sdl/methods.go#L4443) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4689) |
| [SDL_SetWindowFillDocument](https://wiki.libsdl.org/SDL3/SDL_SetWindowFillDocument) | [:heavy_check_mark:](sdl/methods.go#L4453) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11157) |
| [SDL_ShowWindow](https://wiki.libsdl.org/SDL3/SDL_ShowWindow) | [:heavy_check_mark:](sdl/methods.go#L4463) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4700) |
| [SDL_HideWindow](https://wiki.libsdl.org/SDL3/SDL_HideWindow) | [:heavy_check_mark:](sdl/methods.go#L4473) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4709) |
| [SDL_RaiseWindow](https://wiki.libsdl.org/SDL3/SDL_RaiseWindow) | [:heavy_check_mark:](sdl/methods.go#L4483) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4718) |
| [SDL_MaximizeWindow](https://wiki.libsdl.org/SDL3/SDL_MaximizeWindow) | [:heavy_check_mark:](sdl/methods.go#L4493) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4727) |
| [SDL_MinimizeWindow](https://wiki.libsdl.org/SDL3/SDL_MinimizeWindow) | [:heavy_check_mark:](sdl/methods.go#L4503) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4736) |
| [SDL_RestoreWindow](https://wiki.libsdl.org/SDL3/SDL_RestoreWindow) | [:heavy_check_mark:](sdl/methods.go#L4513) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4745) |
| [SDL_SetWindowFullscreen](https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreen) | [:heavy_check_mark:](sdl/methods.go#L4523) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4754) |
| [SDL_SyncWindow](https://wiki.libsdl.org/SDL3/SDL_SyncWindow) | [:heavy_check_mark:](sdl/methods.go#L4533) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4765) |
| [SDL_WindowHasSurface](https://wiki.libsdl.org/SDL3/SDL_WindowHasSurface) | [:heavy_check_mark:](sdl/methods.go#L4543) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4774) |
| [SDL_GetWindowSurface](https://wiki.libsdl.org/SDL3/SDL_GetWindowSurface) | [:heavy_check_mark:](sdl/methods.go#L4549) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4783) |
| [SDL_SetWindowSurfaceVSync](https://wiki.libsdl.org/SDL3/SDL_SetWindowSurfaceVSync) | [:heavy_check_mark:](sdl/methods.go#L4560) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4792) |
| [SDL_GetWindowSurfaceVSync](https://wiki.libsdl.org/SDL3/SDL_GetWindowSurfaceVSync) | [:heavy_check_mark:](sdl/methods.go#L4570) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4803) |
| [SDL_UpdateWindowSurface](https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurface) | [:heavy_check_mark:](sdl/methods.go#L4582) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4817) |
| [SDL_UpdateWindowSurfaceRects](https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurfaceRects) | [:heavy_check_mark:](sdl/methods.go#L4592) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4826) |
| [SDL_DestroyWindowSurface](https://wiki.libsdl.org/SDL3/SDL_DestroyWindowSurface) | [:heavy_check_mark:](sdl/methods.go#L4602) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4839) |
| [SDL_SetWindowKeyboardGrab](https://wiki.libsdl.org/SDL3/SDL_SetWindowKeyboardGrab) | [:heavy_check_mark:](sdl/methods.go#L4612) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4849) |
| [SDL_SetWindowMouseGrab](https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseGrab) | [:heavy_check_mark:](sdl/methods.go#L4622) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4860) |
| [SDL_GetWindowKeyboardGrab](https://wiki.libsdl.org/SDL3/SDL_GetWindowKeyboardGrab) | [:heavy_check_mark:](sdl/methods.go#L4632) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4871) |
| [SDL_GetWindowMouseGrab](https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseGrab) | [:heavy_check_mark:](sdl/methods.go#L4638) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4880) |
| [SDL_GetGrabbedWindow](https://wiki.libsdl.org/SDL3/SDL_GetGrabbedWindow) | [:heavy_check_mark:](sdl/functions.go#L712) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4889) |
| [SDL_SetWindowMouseRect](https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseRect) | [:heavy_check_mark:](sdl/methods.go#L4644) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4896) |
| [SDL_GetWindowMouseRect](https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseRect) | [:heavy_check_mark:](sdl/methods.go#L4654) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4909) |
| [SDL_SetWindowOpacity](https://wiki.libsdl.org/SDL3/SDL_SetWindowOpacity) | [:heavy_check_mark:](sdl/methods.go#L4660) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4918) |
| [SDL_GetWindowOpacity](https://wiki.libsdl.org/SDL3/SDL_GetWindowOpacity) | [:heavy_check_mark:](sdl/methods.go#L4670) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4929) |
| [SDL_SetWindowParent](https://wiki.libsdl.org/SDL3/SDL_SetWindowParent) | [:heavy_check_mark:](sdl/methods.go#L4676) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4938) |
| [SDL_SetWindowModal](https://wiki.libsdl.org/SDL3/SDL_SetWindowModal) | [:heavy_check_mark:](sdl/methods.go#L4686) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4949) |
| [SDL_SetWindowFocusable](https://wiki.libsdl.org/SDL3/SDL_SetWindowFocusable) | [:heavy_check_mark:](sdl/methods.go#L4696) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4960) |
| [SDL_ShowWindowSystemMenu](https://wiki.libsdl.org/SDL3/SDL_ShowWindowSystemMenu) | [:heavy_check_mark:](sdl/methods.go#L4706) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4971) |
| [SDL_SetWindowHitTest](https://wiki.libsdl.org/SDL3/SDL_SetWindowHitTest) | [:heavy_check_mark:](sdl/methods.go#L4716) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4984) |
| [SDL_SetWindowShape](https://wiki.libsdl.org/SDL3/SDL_SetWindowShape) | [:heavy_check_mark:](sdl/methods.go#L4726) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L4995) |
| [SDL_FlashWindow](https://wiki.libsdl.org/SDL3/SDL_FlashWindow) | [:heavy_check_mark:](sdl/methods.go#L4736) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5006) |
| [SDL_SetWindowProgressState](https://wiki.libsdl.org/SDL3/SDL_SetWindowProgressState) | [:heavy_check_mark:](sdl/methods.go#L4746) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11168) |
| [SDL_GetWindowProgressState](https://wiki.libsdl.org/SDL3/SDL_GetWindowProgressState) | [:heavy_check_mark:](sdl/methods.go#L4756) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11179) |
| [SDL_SetWindowProgressValue](https://wiki.libsdl.org/SDL3/SDL_SetWindowProgressValue) | [:heavy_check_mark:](sdl/methods.go#L4767) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11188) |
| [SDL_GetWindowProgressValue](https://wiki.libsdl.org/SDL3/SDL_GetWindowProgressValue) | [:heavy_check_mark:](sdl/methods.go#L4777) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11199) |
| [SDL_DestroyWindow](https://wiki.libsdl.org/SDL3/SDL_DestroyWindow) | [:heavy_check_mark:](sdl/methods.go#L4788) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5017) |
| [SDL_ScreenSaverEnabled](https://wiki.libsdl.org/SDL3/SDL_ScreenSaverEnabled) | [:heavy_check_mark:](sdl/functions.go#L718) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5029) |
| [SDL_EnableScreenSaver](https://wiki.libsdl.org/SDL3/SDL_EnableScreenSaver) | [:heavy_check_mark:](sdl/functions.go#L724) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5036) |
| [SDL_DisableScreenSaver](https://wiki.libsdl.org/SDL3/SDL_DisableScreenSaver) | [:heavy_check_mark:](sdl/functions.go#L734) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5043) |
| [SDL_GL_LoadLibrary](https://wiki.libsdl.org/SDL3/SDL_GL_LoadLibrary) | [:heavy_check_mark:](sdl/functions.go#L744) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5050) |
| [SDL_GL_GetProcAddress](https://wiki.libsdl.org/SDL3/SDL_GL_GetProcAddress) | [:heavy_check_mark:](sdl/functions.go#L754) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11208) |
| [SDL_EGL_GetProcAddress](https://wiki.libsdl.org/SDL3/SDL_EGL_GetProcAddress) | [:heavy_check_mark:](sdl/functions.go#L760) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11221) |
| [SDL_GL_UnloadLibrary](https://wiki.libsdl.org/SDL3/SDL_GL_UnloadLibrary) | [:heavy_check_mark:](sdl/functions.go#L772) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5061) |
| [SDL_GL_ExtensionSupported](https://wiki.libsdl.org/SDL3/SDL_GL_ExtensionSupported) | [:heavy_check_mark:](sdl/functions.go#L766) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5067) |
| [SDL_GL_ResetAttributes](https://wiki.libsdl.org/SDL3/SDL_GL_ResetAttributes) | [:heavy_check_mark:](sdl/functions.go#L821) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5078) |
| [SDL_GL_SetAttribute](https://wiki.libsdl.org/SDL3/SDL_GL_SetAttribute) | [:heavy_check_mark:](sdl/functions.go#L827) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5084) |
| [SDL_GL_GetAttribute](https://wiki.libsdl.org/SDL3/SDL_GL_GetAttribute) | [:heavy_check_mark:](sdl/functions.go#L837) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5095) |
| [SDL_GL_CreateContext](https://wiki.libsdl.org/SDL3/SDL_GL_CreateContext) | [:heavy_check_mark:](sdl/functions.go#L778) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11234) |
| [SDL_GL_MakeCurrent](https://wiki.libsdl.org/SDL3/SDL_GL_MakeCurrent) | [:heavy_check_mark:](sdl/functions.go#L789) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11243) |
| [SDL_GL_GetCurrentWindow](https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentWindow) | [:heavy_check_mark:](sdl/functions.go#L849) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5109) |
| [SDL_GL_GetCurrentContext](https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentContext) | [:heavy_check_mark:](sdl/functions.go#L860) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11254) |
| [SDL_EGL_GetCurrentDisplay](https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentDisplay) | [:heavy_check_mark:](sdl/functions.go#L871) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5116) |
| [SDL_EGL_GetCurrentConfig](https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentConfig) | [:heavy_check_mark:](sdl/functions.go#L882) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5123) |
| [SDL_EGL_GetWindowSurface](https://wiki.libsdl.org/SDL3/SDL_EGL_GetWindowSurface) | [:heavy_check_mark:](sdl/functions.go#L799) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5130) |
| [SDL_EGL_SetAttributeCallbacks](https://wiki.libsdl.org/SDL3/SDL_EGL_SetAttributeCallbacks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11261) |
| [SDL_GL_SetSwapInterval](https://wiki.libsdl.org/SDL3/SDL_GL_SetSwapInterval) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5139) |
| [SDL_GL_GetSwapInterval](https://wiki.libsdl.org/SDL3/SDL_GL_GetSwapInterval) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5148) |
| [SDL_GL_SwapWindow](https://wiki.libsdl.org/SDL3/SDL_GL_SwapWindow) | [:heavy_check_mark:](sdl/functions.go#L805) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5160) |
| [SDL_GL_DestroyContext](https://wiki.libsdl.org/SDL3/SDL_GL_DestroyContext) | [:heavy_check_mark:](sdl/functions.go#L815) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11271) |
</details>
<details open>
<summary><h3>Events</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_PumpEvents](https://wiki.libsdl.org/SDL3/SDL_PumpEvents) | [:heavy_check_mark:](sdl/functions.go#L185) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7216) |
| [SDL_PeepEvents](https://wiki.libsdl.org/SDL3/SDL_PeepEvents) | [:heavy_check_mark:](sdl/functions.go#L213) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7222) |
| [SDL_HasEvent](https://wiki.libsdl.org/SDL3/SDL_HasEvent) | [:heavy_check_mark:](sdl/functions.go#L239) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7252) |
| [SDL_HasEvents](https://wiki.libsdl.org/SDL3/SDL_HasEvents) | [:heavy_check_mark:](sdl/functions.go#L245) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7261) |
| [SDL_FlushEvent](https://wiki.libsdl.org/SDL3/SDL_FlushEvent) | [:heavy_check_mark:](sdl/functions.go#L251) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7272) |
| [SDL_FlushEvents](https://wiki.libsdl.org/SDL3/SDL_FlushEvents) | [:heavy_check_mark:](sdl/functions.go#L257) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7280) |
| [SDL_PollEvent](https://wiki.libsdl.org/SDL3/SDL_PollEvent) | [:heavy_check_mark:](sdl/functions.go#L263) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7290) |
| [SDL_WaitEvent](https://wiki.libsdl.org/SDL3/SDL_WaitEvent) | [:heavy_check_mark:](sdl/functions.go#L269) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7303) |
| [SDL_WaitEventTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitEventTimeout) | [:heavy_check_mark:](sdl/functions.go#L279) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7317) |
| [SDL_PushEvent](https://wiki.libsdl.org/SDL3/SDL_PushEvent) | [:heavy_check_mark:](sdl/functions.go#L285) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7332) |
| [SDL_SetEventFilter](https://wiki.libsdl.org/SDL3/SDL_SetEventFilter) | [:heavy_check_mark:](sdl/functions.go#L295) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11418) |
| [SDL_GetEventFilter](https://wiki.libsdl.org/SDL3/SDL_GetEventFilter) | [:heavy_check_mark:](sdl/functions.go#L301) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7343) |
| [SDL_AddEventWatch](https://wiki.libsdl.org/SDL3/SDL_AddEventWatch) | [:heavy_check_mark:](sdl/functions.go#L314) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11426) |
| [SDL_RemoveEventWatch](https://wiki.libsdl.org/SDL3/SDL_RemoveEventWatch) | [:heavy_check_mark:](sdl/functions.go#L324) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11435) |
| [SDL_FilterEvents](https://wiki.libsdl.org/SDL3/SDL_FilterEvents) | [:heavy_check_mark:](sdl/functions.go#L330) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11443) |
| [SDL_SetEventEnabled](https://wiki.libsdl.org/SDL3/SDL_SetEventEnabled) | [:heavy_check_mark:](sdl/functions.go#L336) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7358) |
| [SDL_EventEnabled](https://wiki.libsdl.org/SDL3/SDL_EventEnabled) | [:heavy_check_mark:](sdl/functions.go#L342) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7368) |
| [SDL_RegisterEvents](https://wiki.libsdl.org/SDL3/SDL_RegisterEvents) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7377) |
| [SDL_GetWindowFromEvent](https://wiki.libsdl.org/SDL3/SDL_GetWindowFromEvent) | [:heavy_check_mark:](sdl/methods.go#L1965) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7386) |
| [SDL_GetEventDescription](https://wiki.libsdl.org/SDL3/SDL_GetEventDescription) | [:heavy_check_mark:](sdl/methods.go#L1971) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11451) |
</details>
<details open>
<summary><h3>Keyboard</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_HasKeyboard](https://wiki.libsdl.org/SDL3/SDL_HasKeyboard) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6674) |
| [SDL_GetKeyboards](https://wiki.libsdl.org/SDL3/SDL_GetKeyboards) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6682) |
| [SDL_GetKeyboardNameForID](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardNameForID) | [:heavy_check_mark:](sdl/methods.go#L3822) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6694) |
| [SDL_GetKeyboardFocus](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardFocus) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6703) |
| [SDL_GetKeyboardState](https://wiki.libsdl.org/SDL3/SDL_GetKeyboardState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11379) |
| [SDL_ResetKeyboard](https://wiki.libsdl.org/SDL3/SDL_ResetKeyboard) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6711) |
| [SDL_GetModState](https://wiki.libsdl.org/SDL3/SDL_GetModState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6717) |
| [SDL_SetModState](https://wiki.libsdl.org/SDL3/SDL_SetModState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6725) |
| [SDL_GetKeyFromScancode](https://wiki.libsdl.org/SDL3/SDL_GetKeyFromScancode) | [:heavy_check_mark:](sdl/methods.go#L4917) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6732) |
| [SDL_GetScancodeFromKey](https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromKey) | [:heavy_check_mark:](sdl/methods.go#L6013) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6745) |
| [SDL_SetScancodeName](https://wiki.libsdl.org/SDL3/SDL_SetScancodeName) | [:heavy_check_mark:](sdl/methods.go#L4923) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6763) |
| [SDL_GetScancodeName](https://wiki.libsdl.org/SDL3/SDL_GetScancodeName) | [:heavy_check_mark:](sdl/methods.go#L4933) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6776) |
| [SDL_GetScancodeFromName](https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromName) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6785) |
| [SDL_GetKeyName](https://wiki.libsdl.org/SDL3/SDL_GetKeyName) | [:heavy_check_mark:](sdl/methods.go#L6019) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6797) |
| [SDL_GetKeyFromName](https://wiki.libsdl.org/SDL3/SDL_GetKeyFromName) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6807) |
| [SDL_StartTextInput](https://wiki.libsdl.org/SDL3/SDL_StartTextInput) | [:heavy_check_mark:](sdl/methods.go#L4794) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6819) |
| [SDL_StartTextInputWithProperties](https://wiki.libsdl.org/SDL3/SDL_StartTextInputWithProperties) | [:heavy_check_mark:](sdl/methods.go#L4804) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6832) |
| [SDL_TextInputActive](https://wiki.libsdl.org/SDL3/SDL_TextInputActive) | [:heavy_check_mark:](sdl/methods.go#L4814) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6846) |
| [SDL_StopTextInput](https://wiki.libsdl.org/SDL3/SDL_StopTextInput) | [:heavy_check_mark:](sdl/methods.go#L4820) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6859) |
| [SDL_ClearComposition](https://wiki.libsdl.org/SDL3/SDL_ClearComposition) | [:heavy_check_mark:](sdl/methods.go#L4830) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6872) |
| [SDL_SetTextInputArea](https://wiki.libsdl.org/SDL3/SDL_SetTextInputArea) | [:heavy_check_mark:](sdl/methods.go#L4840) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6885) |
| [SDL_GetTextInputArea](https://wiki.libsdl.org/SDL3/SDL_GetTextInputArea) | [:heavy_check_mark:](sdl/methods.go#L4850) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6903) |
| [SDL_HasScreenKeyboardSupport](https://wiki.libsdl.org/SDL3/SDL_HasScreenKeyboardSupport) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6920) |
| [SDL_ScreenKeyboardShown](https://wiki.libsdl.org/SDL3/SDL_ScreenKeyboardShown) | [:heavy_check_mark:](sdl/methods.go#L4863) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6928) |
</details>
<details open>
<summary><h3>Mouse</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_HasMouse](https://wiki.libsdl.org/SDL3/SDL_HasMouse) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6941) |
| [SDL_GetMice](https://wiki.libsdl.org/SDL3/SDL_GetMice) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6948) |
| [SDL_GetMouseNameForID](https://wiki.libsdl.org/SDL3/SDL_GetMouseNameForID) | [:heavy_check_mark:](sdl/methods.go#L3835) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6960) |
| [SDL_GetMouseFocus](https://wiki.libsdl.org/SDL3/SDL_GetMouseFocus) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6969) |
| [SDL_GetMouseState](https://wiki.libsdl.org/SDL3/SDL_GetMouseState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6976) |
| [SDL_GetGlobalMouseState](https://wiki.libsdl.org/SDL3/SDL_GetGlobalMouseState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6992) |
| [SDL_GetRelativeMouseState](https://wiki.libsdl.org/SDL3/SDL_GetRelativeMouseState) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7007) |
| [SDL_WarpMouseInWindow](https://wiki.libsdl.org/SDL3/SDL_WarpMouseInWindow) | [:heavy_check_mark:](sdl/methods.go#L4869) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7022) |
| [SDL_WarpMouseGlobal](https://wiki.libsdl.org/SDL3/SDL_WarpMouseGlobal) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7034) |
| [SDL_SetRelativeMouseTransform](https://wiki.libsdl.org/SDL3/SDL_SetRelativeMouseTransform) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11393) |
| [SDL_SetWindowRelativeMouseMode](https://wiki.libsdl.org/SDL3/SDL_SetWindowRelativeMouseMode) | [:heavy_check_mark:](sdl/methods.go#L4875) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7045) |
| [SDL_GetWindowRelativeMouseMode](https://wiki.libsdl.org/SDL3/SDL_GetWindowRelativeMouseMode) | [:heavy_check_mark:](sdl/methods.go#L4885) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7056) |
| [SDL_CaptureMouse](https://wiki.libsdl.org/SDL3/SDL_CaptureMouse) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7065) |
| [SDL_CreateCursor](https://wiki.libsdl.org/SDL3/SDL_CreateCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7074) |
| [SDL_CreateColorCursor](https://wiki.libsdl.org/SDL3/SDL_CreateColorCursor) | [:heavy_check_mark:](sdl/methods.go#L1941) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7097) |
| [SDL_CreateAnimatedCursor](https://wiki.libsdl.org/SDL3/SDL_CreateAnimatedCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11402) |
| [SDL_CreateSystemCursor](https://wiki.libsdl.org/SDL3/SDL_CreateSystemCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7110) |
| [SDL_SetCursor](https://wiki.libsdl.org/SDL3/SDL_SetCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7119) |
| [SDL_GetCursor](https://wiki.libsdl.org/SDL3/SDL_GetCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7128) |
| [SDL_GetDefaultCursor](https://wiki.libsdl.org/SDL3/SDL_GetDefaultCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7135) |
| [SDL_DestroyCursor](https://wiki.libsdl.org/SDL3/SDL_DestroyCursor) | [:heavy_check_mark:](sdl/methods.go#L664) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7142) |
| [SDL_ShowCursor](https://wiki.libsdl.org/SDL3/SDL_ShowCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7151) |
| [SDL_HideCursor](https://wiki.libsdl.org/SDL3/SDL_HideCursor) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7158) |
| [SDL_CursorVisible](https://wiki.libsdl.org/SDL3/SDL_CursorVisible) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7165) |
</details>
<details open>
<summary><h3>Touch</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetTouchDevices](https://wiki.libsdl.org/SDL3/SDL_GetTouchDevices) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7172) |
| [SDL_GetTouchDeviceName](https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceName) | [:heavy_check_mark:](sdl/methods.go#L46) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7184) |
| [SDL_GetTouchDeviceType](https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceType) | [:heavy_check_mark:](sdl/methods.go#L57) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7193) |
| [SDL_GetTouchFingers](https://wiki.libsdl.org/SDL3/SDL_GetTouchFingers) | [:heavy_check_mark:](sdl/methods.go#L63) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7202) |
</details>
<details open>
<summary><h3>Gamepad</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AddGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMapping) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5939) |
| [SDL_AddGamepadMappingsFromIO](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromIO) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L5950) |
| [SDL_AddGamepadMappingsFromFile](https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromFile) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5961) |
| [SDL_ReloadGamepadMappings](https://wiki.libsdl.org/SDL3/SDL_ReloadGamepadMappings) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5972) |
| [SDL_GetGamepadMappings](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappings) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5979) |
| [SDL_GetGamepadMappingForGUID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForGUID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11358) |
| [SDL_GetGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMapping) | [:heavy_check_mark:](sdl/methods.go#L2510) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5991) |
| [SDL_SetGamepadMapping](https://wiki.libsdl.org/SDL3/SDL_SetGamepadMapping) | [:heavy_check_mark:](sdl/methods.go#L801) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6000) |
| [SDL_HasGamepad](https://wiki.libsdl.org/SDL3/SDL_HasGamepad) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6012) |
| [SDL_GetGamepads](https://wiki.libsdl.org/SDL3/SDL_GetGamepads) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6019) |
| [SDL_IsGamepad](https://wiki.libsdl.org/SDL3/SDL_IsGamepad) | [:heavy_check_mark:](sdl/methods.go#L813) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6031) |
| [SDL_GetGamepadNameForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadNameForID) | [:heavy_check_mark:](sdl/methods.go#L819) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6040) |
| [SDL_GetGamepadPathForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPathForID) | [:heavy_check_mark:](sdl/methods.go#L830) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6049) |
| [SDL_GetGamepadPlayerIndexForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndexForID) | [:heavy_check_mark:](sdl/methods.go#L841) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6058) |
| [SDL_GetGamepadGUIDForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadGUIDForID) | [:x:](sdl/methods.go#L847) | [:x:](sdl/sdl_functions_js.go#L11368) |
| [SDL_GetGamepadVendorForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendorForID) | [:heavy_check_mark:](sdl/methods.go#L854) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6067) |
| [SDL_GetGamepadProductForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductForID) | [:heavy_check_mark:](sdl/methods.go#L860) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6076) |
| [SDL_GetGamepadProductVersionForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersionForID) | [:heavy_check_mark:](sdl/methods.go#L866) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6085) |
| [SDL_GetGamepadTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeForID) | [:heavy_check_mark:](sdl/methods.go#L872) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6094) |
| [SDL_GetRealGamepadTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadTypeForID) | [:heavy_check_mark:](sdl/methods.go#L878) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6103) |
| [SDL_GetGamepadMappingForID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForID) | [:heavy_check_mark:](sdl/methods.go#L884) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6112) |
| [SDL_OpenGamepad](https://wiki.libsdl.org/SDL3/SDL_OpenGamepad) | [:heavy_check_mark:](sdl/methods.go#L896) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6120) |
| [SDL_GetGamepadFromID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromID) | [:heavy_check_mark:](sdl/methods.go#L907) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6129) |
| [SDL_GetGamepadFromPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromPlayerIndex) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6138) |
| [SDL_GetGamepadProperties](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProperties) | [:heavy_check_mark:](sdl/methods.go#L2522) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6147) |
| [SDL_GetGamepadID](https://wiki.libsdl.org/SDL3/SDL_GetGamepadID) | [:heavy_check_mark:](sdl/methods.go#L2533) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6156) |
| [SDL_GetGamepadName](https://wiki.libsdl.org/SDL3/SDL_GetGamepadName) | [:heavy_check_mark:](sdl/methods.go#L2544) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6165) |
| [SDL_GetGamepadPath](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPath) | [:heavy_check_mark:](sdl/methods.go#L2550) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6174) |
| [SDL_GetGamepadType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadType) | [:heavy_check_mark:](sdl/methods.go#L2556) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6183) |
| [SDL_GetRealGamepadType](https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadType) | [:heavy_check_mark:](sdl/methods.go#L2562) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6192) |
| [SDL_GetGamepadPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L2568) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6201) |
| [SDL_SetGamepadPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_SetGamepadPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L2574) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6210) |
| [SDL_GetGamepadVendor](https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendor) | [:heavy_check_mark:](sdl/methods.go#L2584) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6221) |
| [SDL_GetGamepadProduct](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProduct) | [:heavy_check_mark:](sdl/methods.go#L2590) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6230) |
| [SDL_GetGamepadProductVersion](https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersion) | [:heavy_check_mark:](sdl/methods.go#L2596) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6239) |
| [SDL_GetGamepadFirmwareVersion](https://wiki.libsdl.org/SDL3/SDL_GetGamepadFirmwareVersion) | [:heavy_check_mark:](sdl/methods.go#L2602) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6248) |
| [SDL_GetGamepadSerial](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSerial) | [:heavy_check_mark:](sdl/methods.go#L2608) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6257) |
| [SDL_GetGamepadSteamHandle](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSteamHandle) | [:heavy_check_mark:](sdl/methods.go#L2614) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6266) |
| [SDL_GetGamepadConnectionState](https://wiki.libsdl.org/SDL3/SDL_GetGamepadConnectionState) | [:heavy_check_mark:](sdl/methods.go#L2620) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6275) |
| [SDL_GetGamepadPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetGamepadPowerInfo) | [:heavy_check_mark:](sdl/methods.go#L2632) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6284) |
| [SDL_GamepadConnected](https://wiki.libsdl.org/SDL3/SDL_GamepadConnected) | [:heavy_check_mark:](sdl/methods.go#L2642) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6298) |
| [SDL_GetGamepadJoystick](https://wiki.libsdl.org/SDL3/SDL_GetGamepadJoystick) | [:heavy_check_mark:](sdl/methods.go#L2648) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6307) |
| [SDL_SetGamepadEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_SetGamepadEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6316) |
| [SDL_GamepadEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_GamepadEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6324) |
| [SDL_GetGamepadBindings](https://wiki.libsdl.org/SDL3/SDL_GetGamepadBindings) | [:heavy_check_mark:](sdl/methods.go#L2659) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6331) |
| [SDL_UpdateGamepads](https://wiki.libsdl.org/SDL3/SDL_UpdateGamepads) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6345) |
| [SDL_GetGamepadTypeFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeFromString) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6351) |
| [SDL_GetGamepadStringForType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForType) | [:heavy_check_mark:](sdl/methods.go#L5740) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6362) |
| [SDL_GetGamepadAxisFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxisFromString) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6371) |
| [SDL_GetGamepadStringForAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForAxis) | [:heavy_check_mark:](sdl/methods.go#L656) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6382) |
| [SDL_GamepadHasAxis](https://wiki.libsdl.org/SDL3/SDL_GamepadHasAxis) | [:heavy_check_mark:](sdl/methods.go#L2673) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6391) |
| [SDL_GetGamepadAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxis) | [:heavy_check_mark:](sdl/methods.go#L2679) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6402) |
| [SDL_GetGamepadButtonFromString](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonFromString) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6413) |
| [SDL_GetGamepadStringForButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForButton) | [:heavy_check_mark:](sdl/methods.go#L551) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6424) |
| [SDL_GamepadHasButton](https://wiki.libsdl.org/SDL3/SDL_GamepadHasButton) | [:heavy_check_mark:](sdl/methods.go#L2685) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6433) |
| [SDL_GetGamepadButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButton) | [:heavy_check_mark:](sdl/methods.go#L2691) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6444) |
| [SDL_GetGamepadButtonLabelForType](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabelForType) | [:heavy_check_mark:](sdl/methods.go#L5746) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6455) |
| [SDL_GetGamepadButtonLabel](https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabel) | [:heavy_check_mark:](sdl/methods.go#L2697) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6466) |
| [SDL_GetNumGamepadTouchpads](https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpads) | [:heavy_check_mark:](sdl/methods.go#L2703) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6477) |
| [SDL_GetNumGamepadTouchpadFingers](https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpadFingers) | [:heavy_check_mark:](sdl/methods.go#L2709) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6486) |
| [SDL_GetGamepadTouchpadFinger](https://wiki.libsdl.org/SDL3/SDL_GetGamepadTouchpadFinger) | [:x:](sdl/methods.go#L2715) | [:x:](sdl/sdl_functions_js.go#L6497) |
| [SDL_GamepadHasSensor](https://wiki.libsdl.org/SDL3/SDL_GamepadHasSensor) | [:heavy_check_mark:](sdl/methods.go#L2722) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6521) |
| [SDL_SetGamepadSensorEnabled](https://wiki.libsdl.org/SDL3/SDL_SetGamepadSensorEnabled) | [:heavy_check_mark:](sdl/methods.go#L2728) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6532) |
| [SDL_GamepadSensorEnabled](https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEnabled) | [:heavy_check_mark:](sdl/methods.go#L2738) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6545) |
| [SDL_GetGamepadSensorDataRate](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorDataRate) | [:heavy_check_mark:](sdl/methods.go#L2744) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6556) |
| [SDL_GetGamepadSensorData](https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorData) | [:x:](sdl/methods.go#L2750) | [:x:](sdl/sdl_functions_js.go#L6567) |
| [SDL_GamepadHasCapSense](https://wiki.libsdl.org/SDL3/SDL_GamepadHasCapSense) | [:question:]() | [:question:]() |
| [SDL_GetGamepadCapSense](https://wiki.libsdl.org/SDL3/SDL_GetGamepadCapSense) | [:question:]() | [:question:]() |
| [SDL_RumbleGamepad](https://wiki.libsdl.org/SDL3/SDL_RumbleGamepad) | [:heavy_check_mark:](sdl/methods.go#L2757) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6585) |
| [SDL_RumbleGamepadTriggers](https://wiki.libsdl.org/SDL3/SDL_RumbleGamepadTriggers) | [:heavy_check_mark:](sdl/methods.go#L2767) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6600) |
| [SDL_SetGamepadLED](https://wiki.libsdl.org/SDL3/SDL_SetGamepadLED) | [:heavy_check_mark:](sdl/methods.go#L2777) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6615) |
| [SDL_SendGamepadEffect](https://wiki.libsdl.org/SDL3/SDL_SendGamepadEffect) | [:heavy_check_mark:](sdl/methods.go#L2787) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6630) |
| [SDL_CloseGamepad](https://wiki.libsdl.org/SDL3/SDL_CloseGamepad) | [:heavy_check_mark:](sdl/methods.go#L2798) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6643) |
| [SDL_GetGamepadAppleSFSymbolsNameForButton](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForButton) | [:heavy_check_mark:](sdl/methods.go#L2804) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6652) |
| [SDL_GetGamepadAppleSFSymbolsNameForAxis](https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForAxis) | [:heavy_check_mark:](sdl/methods.go#L2810) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L6663) |
</details>
<details open>
<summary><h3>Joystick</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_LockJoysticks](https://wiki.libsdl.org/SDL3/SDL_LockJoysticks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5323) |
| [SDL_TryLockJoysticks](https://wiki.libsdl.org/SDL3/SDL_TryLockJoysticks) | [:question:]() | [:question:]() |
| [SDL_UnlockJoysticks](https://wiki.libsdl.org/SDL3/SDL_UnlockJoysticks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5329) |
| [SDL_HasJoystick](https://wiki.libsdl.org/SDL3/SDL_HasJoystick) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5335) |
| [SDL_GetJoysticks](https://wiki.libsdl.org/SDL3/SDL_GetJoysticks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5342) |
| [SDL_GetJoystickNameForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickNameForID) | [:heavy_check_mark:](sdl/methods.go#L704) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5354) |
| [SDL_GetJoystickPathForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPathForID) | [:heavy_check_mark:](sdl/methods.go#L715) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5363) |
| [SDL_GetJoystickPlayerIndexForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndexForID) | [:heavy_check_mark:](sdl/methods.go#L726) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5372) |
| [SDL_GetJoystickGUIDForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDForID) | [:x:](sdl/methods.go#L732) | [:x:](sdl/sdl_functions_js.go#L11317) |
| [SDL_GetJoystickVendorForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendorForID) | [:heavy_check_mark:](sdl/methods.go#L739) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5381) |
| [SDL_GetJoystickProductForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductForID) | [:heavy_check_mark:](sdl/methods.go#L745) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5390) |
| [SDL_GetJoystickProductVersionForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersionForID) | [:heavy_check_mark:](sdl/methods.go#L751) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5399) |
| [SDL_GetJoystickTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickTypeForID) | [:heavy_check_mark:](sdl/methods.go#L757) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5408) |
| [SDL_OpenJoystick](https://wiki.libsdl.org/SDL3/SDL_OpenJoystick) | [:heavy_check_mark:](sdl/methods.go#L763) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5417) |
| [SDL_GetJoystickFromID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromID) | [:heavy_check_mark:](sdl/methods.go#L774) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5429) |
| [SDL_GetJoystickFromPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromPlayerIndex) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5443) |
| [SDL_AttachVirtualJoystick](https://wiki.libsdl.org/SDL3/SDL_AttachVirtualJoystick) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5452) |
| [SDL_DetachVirtualJoystick](https://wiki.libsdl.org/SDL3/SDL_DetachVirtualJoystick) | [:heavy_check_mark:](sdl/methods.go#L785) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5463) |
| [SDL_IsJoystickVirtual](https://wiki.libsdl.org/SDL3/SDL_IsJoystickVirtual) | [:heavy_check_mark:](sdl/methods.go#L795) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5472) |
| [SDL_SetJoystickVirtualAxis](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualAxis) | [:heavy_check_mark:](sdl/methods.go#L5398) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5481) |
| [SDL_SetJoystickVirtualBall](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualBall) | [:heavy_check_mark:](sdl/methods.go#L5408) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5494) |
| [SDL_SetJoystickVirtualButton](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualButton) | [:heavy_check_mark:](sdl/methods.go#L5418) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5509) |
| [SDL_SetJoystickVirtualHat](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualHat) | [:heavy_check_mark:](sdl/methods.go#L5428) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5522) |
| [SDL_SetJoystickVirtualTouchpad](https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualTouchpad) | [:heavy_check_mark:](sdl/methods.go#L5438) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5535) |
| [SDL_SendJoystickVirtualSensorData](https://wiki.libsdl.org/SDL3/SDL_SendJoystickVirtualSensorData) | [:heavy_check_mark:](sdl/methods.go#L5448) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5556) |
| [SDL_GetJoystickProperties](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProperties) | [:heavy_check_mark:](sdl/methods.go#L5458) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5571) |
| [SDL_GetJoystickName](https://wiki.libsdl.org/SDL3/SDL_GetJoystickName) | [:heavy_check_mark:](sdl/methods.go#L5469) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5580) |
| [SDL_GetJoystickPath](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPath) | [:heavy_check_mark:](sdl/methods.go#L5480) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5593) |
| [SDL_GetJoystickPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L5491) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5602) |
| [SDL_SetJoystickPlayerIndex](https://wiki.libsdl.org/SDL3/SDL_SetJoystickPlayerIndex) | [:heavy_check_mark:](sdl/methods.go#L5497) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5611) |
| [SDL_GetJoystickGUID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUID) | [:x:](sdl/methods.go#L5507) | [:x:](sdl/sdl_functions_js.go#L11328) |
| [SDL_GetJoystickVendor](https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendor) | [:heavy_check_mark:](sdl/methods.go#L5514) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5622) |
| [SDL_GetJoystickProduct](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProduct) | [:heavy_check_mark:](sdl/methods.go#L5520) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5631) |
| [SDL_GetJoystickProductVersion](https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersion) | [:heavy_check_mark:](sdl/methods.go#L5526) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5640) |
| [SDL_GetJoystickFirmwareVersion](https://wiki.libsdl.org/SDL3/SDL_GetJoystickFirmwareVersion) | [:heavy_check_mark:](sdl/methods.go#L5532) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5649) |
| [SDL_GetJoystickSerial](https://wiki.libsdl.org/SDL3/SDL_GetJoystickSerial) | [:heavy_check_mark:](sdl/methods.go#L5538) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5658) |
| [SDL_GetJoystickType](https://wiki.libsdl.org/SDL3/SDL_GetJoystickType) | [:heavy_check_mark:](sdl/methods.go#L5544) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5667) |
| [SDL_GetJoystickGUIDInfo](https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDInfo) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11339) |
| [SDL_JoystickConnected](https://wiki.libsdl.org/SDL3/SDL_JoystickConnected) | [:heavy_check_mark:](sdl/methods.go#L5550) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5676) |
| [SDL_GetJoystickID](https://wiki.libsdl.org/SDL3/SDL_GetJoystickID) | [:heavy_check_mark:](sdl/methods.go#L5556) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5685) |
| [SDL_GetNumJoystickAxes](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickAxes) | [:heavy_check_mark:](sdl/methods.go#L5567) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5698) |
| [SDL_GetNumJoystickBalls](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickBalls) | [:heavy_check_mark:](sdl/methods.go#L5578) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5711) |
| [SDL_GetNumJoystickHats](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickHats) | [:heavy_check_mark:](sdl/methods.go#L5589) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5720) |
| [SDL_GetNumJoystickButtons](https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickButtons) | [:heavy_check_mark:](sdl/methods.go#L5600) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5733) |
| [SDL_SetJoystickEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_SetJoystickEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5746) |
| [SDL_JoystickEventsEnabled](https://wiki.libsdl.org/SDL3/SDL_JoystickEventsEnabled) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5754) |
| [SDL_UpdateJoysticks](https://wiki.libsdl.org/SDL3/SDL_UpdateJoysticks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5761) |
| [SDL_GetJoystickAxis](https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxis) | [:heavy_check_mark:](sdl/methods.go#L5611) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5767) |
| [SDL_GetJoystickAxisInitialState](https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxisInitialState) | [:heavy_check_mark:](sdl/methods.go#L5622) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5782) |
| [SDL_GetJoystickBall](https://wiki.libsdl.org/SDL3/SDL_GetJoystickBall) | [:heavy_check_mark:](sdl/methods.go#L5631) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5797) |
| [SDL_GetJoystickHat](https://wiki.libsdl.org/SDL3/SDL_GetJoystickHat) | [:heavy_check_mark:](sdl/methods.go#L5643) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5815) |
| [SDL_GetJoystickButton](https://wiki.libsdl.org/SDL3/SDL_GetJoystickButton) | [:heavy_check_mark:](sdl/methods.go#L5649) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5830) |
| [SDL_RumbleJoystick](https://wiki.libsdl.org/SDL3/SDL_RumbleJoystick) | [:heavy_check_mark:](sdl/methods.go#L5655) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5845) |
| [SDL_RumbleJoystickTriggers](https://wiki.libsdl.org/SDL3/SDL_RumbleJoystickTriggers) | [:heavy_check_mark:](sdl/methods.go#L5661) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5860) |
| [SDL_SetJoystickLED](https://wiki.libsdl.org/SDL3/SDL_SetJoystickLED) | [:heavy_check_mark:](sdl/methods.go#L5671) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5875) |
| [SDL_SendJoystickEffect](https://wiki.libsdl.org/SDL3/SDL_SendJoystickEffect) | [:heavy_check_mark:](sdl/methods.go#L5681) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5890) |
| [SDL_CloseJoystick](https://wiki.libsdl.org/SDL3/SDL_CloseJoystick) | [:heavy_check_mark:](sdl/methods.go#L5691) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5903) |
| [SDL_GetJoystickConnectionState](https://wiki.libsdl.org/SDL3/SDL_GetJoystickConnectionState) | [:heavy_check_mark:](sdl/methods.go#L5697) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5916) |
| [SDL_GetJoystickPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetJoystickPowerInfo) | [:heavy_check_mark:](sdl/methods.go#L5708) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5925) |
</details>
<details open>
<summary><h3>Haptic</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetHaptics](https://wiki.libsdl.org/SDL3/SDL_GetHaptics) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8238) |
| [SDL_GetHapticNameForID](https://wiki.libsdl.org/SDL3/SDL_GetHapticNameForID) | [:heavy_check_mark:](sdl/methods.go#L2862) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8250) |
| [SDL_OpenHaptic](https://wiki.libsdl.org/SDL3/SDL_OpenHaptic) | [:heavy_check_mark:](sdl/methods.go#L2873) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8259) |
| [SDL_GetHapticFromID](https://wiki.libsdl.org/SDL3/SDL_GetHapticFromID) | [:heavy_check_mark:](sdl/methods.go#L2884) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8268) |
| [SDL_GetHapticID](https://wiki.libsdl.org/SDL3/SDL_GetHapticID) | [:heavy_check_mark:](sdl/methods.go#L2291) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8277) |
| [SDL_GetHapticName](https://wiki.libsdl.org/SDL3/SDL_GetHapticName) | [:heavy_check_mark:](sdl/methods.go#L2302) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8286) |
| [SDL_IsMouseHaptic](https://wiki.libsdl.org/SDL3/SDL_IsMouseHaptic) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8295) |
| [SDL_OpenHapticFromMouse](https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromMouse) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8302) |
| [SDL_IsJoystickHaptic](https://wiki.libsdl.org/SDL3/SDL_IsJoystickHaptic) | [:heavy_check_mark:](sdl/methods.go#L5721) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8309) |
| [SDL_OpenHapticFromJoystick](https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromJoystick) | [:heavy_check_mark:](sdl/methods.go#L5727) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8318) |
| [SDL_CloseHaptic](https://wiki.libsdl.org/SDL3/SDL_CloseHaptic) | [:heavy_check_mark:](sdl/methods.go#L2313) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8327) |
| [SDL_GetMaxHapticEffects](https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffects) | [:heavy_check_mark:](sdl/methods.go#L2319) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8336) |
| [SDL_GetMaxHapticEffectsPlaying](https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffectsPlaying) | [:heavy_check_mark:](sdl/methods.go#L2330) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8345) |
| [SDL_GetHapticFeatures](https://wiki.libsdl.org/SDL3/SDL_GetHapticFeatures) | [:heavy_check_mark:](sdl/methods.go#L2341) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8354) |
| [SDL_GetNumHapticAxes](https://wiki.libsdl.org/SDL3/SDL_GetNumHapticAxes) | [:heavy_check_mark:](sdl/methods.go#L2352) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8363) |
| [SDL_HapticEffectSupported](https://wiki.libsdl.org/SDL3/SDL_HapticEffectSupported) | [:heavy_check_mark:](sdl/methods.go#L2363) | [:no_entry_sign:](sdl/sdl_functions_js.go#L8372) |
| [SDL_CreateHapticEffect](https://wiki.libsdl.org/SDL3/SDL_CreateHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2369) | [:no_entry_sign:](sdl/sdl_functions_js.go#L8376) |
| [SDL_UpdateHapticEffect](https://wiki.libsdl.org/SDL3/SDL_UpdateHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2380) | [:no_entry_sign:](sdl/sdl_functions_js.go#L8380) |
| [SDL_RunHapticEffect](https://wiki.libsdl.org/SDL3/SDL_RunHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2390) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8384) |
| [SDL_StopHapticEffect](https://wiki.libsdl.org/SDL3/SDL_StopHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2400) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8397) |
| [SDL_DestroyHapticEffect](https://wiki.libsdl.org/SDL3/SDL_DestroyHapticEffect) | [:heavy_check_mark:](sdl/methods.go#L2410) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8408) |
| [SDL_GetHapticEffectStatus](https://wiki.libsdl.org/SDL3/SDL_GetHapticEffectStatus) | [:heavy_check_mark:](sdl/methods.go#L2416) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8419) |
| [SDL_SetHapticGain](https://wiki.libsdl.org/SDL3/SDL_SetHapticGain) | [:heavy_check_mark:](sdl/methods.go#L2422) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8430) |
| [SDL_SetHapticAutocenter](https://wiki.libsdl.org/SDL3/SDL_SetHapticAutocenter) | [:heavy_check_mark:](sdl/methods.go#L2432) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8441) |
| [SDL_PauseHaptic](https://wiki.libsdl.org/SDL3/SDL_PauseHaptic) | [:heavy_check_mark:](sdl/methods.go#L2442) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8452) |
| [SDL_ResumeHaptic](https://wiki.libsdl.org/SDL3/SDL_ResumeHaptic) | [:heavy_check_mark:](sdl/methods.go#L2452) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8461) |
| [SDL_StopHapticEffects](https://wiki.libsdl.org/SDL3/SDL_StopHapticEffects) | [:heavy_check_mark:](sdl/methods.go#L2462) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8470) |
| [SDL_HapticRumbleSupported](https://wiki.libsdl.org/SDL3/SDL_HapticRumbleSupported) | [:heavy_check_mark:](sdl/methods.go#L2472) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8479) |
| [SDL_InitHapticRumble](https://wiki.libsdl.org/SDL3/SDL_InitHapticRumble) | [:heavy_check_mark:](sdl/methods.go#L2478) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8488) |
| [SDL_PlayHapticRumble](https://wiki.libsdl.org/SDL3/SDL_PlayHapticRumble) | [:heavy_check_mark:](sdl/methods.go#L2488) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8497) |
| [SDL_StopHapticRumble](https://wiki.libsdl.org/SDL3/SDL_StopHapticRumble) | [:heavy_check_mark:](sdl/methods.go#L2498) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8510) |
</details>
<details open>
<summary><h3>Audio</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumAudioDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumAudioDrivers) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1601) |
| [SDL_GetAudioDriver](https://wiki.libsdl.org/SDL3/SDL_GetAudioDriver) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1608) |
| [SDL_GetCurrentAudioDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentAudioDriver) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1617) |
| [SDL_GetAudioPlaybackDevices](https://wiki.libsdl.org/SDL3/SDL_GetAudioPlaybackDevices) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1624) |
| [SDL_GetAudioRecordingDevices](https://wiki.libsdl.org/SDL3/SDL_GetAudioRecordingDevices) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1636) |
| [SDL_GetAudioDeviceName](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceName) | [:heavy_check_mark:](sdl/methods.go#L212) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1648) |
| [SDL_GetAudioDeviceFormat](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceFormat) | [:heavy_check_mark:](sdl/methods.go#L223) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1657) |
| [SDL_GetAudioDeviceChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceChannelMap) | [:heavy_check_mark:](sdl/methods.go#L236) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1673) |
| [SDL_OpenAudioDevice](https://wiki.libsdl.org/SDL3/SDL_OpenAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L250) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1686) |
| [SDL_IsAudioDevicePhysical](https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePhysical) | [:heavy_check_mark:](sdl/methods.go#L261) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1697) |
| [SDL_IsAudioDevicePlayback](https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePlayback) | [:heavy_check_mark:](sdl/methods.go#L267) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1706) |
| [SDL_PauseAudioDevice](https://wiki.libsdl.org/SDL3/SDL_PauseAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L273) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1715) |
| [SDL_ResumeAudioDevice](https://wiki.libsdl.org/SDL3/SDL_ResumeAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L283) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1724) |
| [SDL_AudioDevicePaused](https://wiki.libsdl.org/SDL3/SDL_AudioDevicePaused) | [:heavy_check_mark:](sdl/methods.go#L293) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1733) |
| [SDL_GetAudioDeviceGain](https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceGain) | [:heavy_check_mark:](sdl/methods.go#L299) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1742) |
| [SDL_SetAudioDeviceGain](https://wiki.libsdl.org/SDL3/SDL_SetAudioDeviceGain) | [:heavy_check_mark:](sdl/methods.go#L310) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1751) |
| [SDL_CloseAudioDevice](https://wiki.libsdl.org/SDL3/SDL_CloseAudioDevice) | [:heavy_check_mark:](sdl/methods.go#L320) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1762) |
| [SDL_BindAudioStreams](https://wiki.libsdl.org/SDL3/SDL_BindAudioStreams) | [:heavy_check_mark:](sdl/methods.go#L326) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1770) |
| [SDL_BindAudioStream](https://wiki.libsdl.org/SDL3/SDL_BindAudioStream) | [:heavy_check_mark:](sdl/methods.go#L336) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1782) |
| [SDL_UnbindAudioStreams](https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStreams) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1793) |
| [SDL_UnbindAudioStream](https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStream) | [:heavy_check_mark:](sdl/methods.go#L3850) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1803) |
| [SDL_GetAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamDevice) | [:heavy_check_mark:](sdl/methods.go#L3856) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1811) |
| [SDL_CreateAudioStream](https://wiki.libsdl.org/SDL3/SDL_CreateAudioStream) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1820) |
| [SDL_GetAudioStreamProperties](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamProperties) | [:heavy_check_mark:](sdl/methods.go#L3862) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1831) |
| [SDL_GetAudioStreamFormat](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFormat) | [:heavy_check_mark:](sdl/methods.go#L3873) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1840) |
| [SDL_SetAudioStreamFormat](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFormat) | [:heavy_check_mark:](sdl/methods.go#L3883) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1866) |
| [SDL_GetAudioStreamFrequencyRatio](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFrequencyRatio) | [:heavy_check_mark:](sdl/methods.go#L3893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1879) |
| [SDL_SetAudioStreamFrequencyRatio](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFrequencyRatio) | [:heavy_check_mark:](sdl/methods.go#L3904) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1888) |
| [SDL_GetAudioStreamGain](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamGain) | [:heavy_check_mark:](sdl/methods.go#L3914) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1899) |
| [SDL_SetAudioStreamGain](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGain) | [:heavy_check_mark:](sdl/methods.go#L3925) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1908) |
| [SDL_GetAudioStreamInputChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamInputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3935) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1919) |
| [SDL_GetAudioStreamOutputChannelMap](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamOutputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3949) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1933) |
| [SDL_SetAudioStreamInputChannelMap](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamInputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3963) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1947) |
| [SDL_SetAudioStreamOutputChannelMap](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamOutputChannelMap) | [:heavy_check_mark:](sdl/methods.go#L3973) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1960) |
| [SDL_PutAudioStreamData](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData) | [:heavy_check_mark:](sdl/methods.go#L3983) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1973) |
| [SDL_PutAudioStreamDataNoCopy](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamDataNoCopy) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10968) |
| [SDL_PutAudioStreamPlanarData](https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamPlanarData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10986) |
| [SDL_GetAudioStreamData](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData) | [:heavy_check_mark:](sdl/methods.go#L3995) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1986) |
| [SDL_GetAudioStreamAvailable](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamAvailable) | [:heavy_check_mark:](sdl/methods.go#L4006) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2003) |
| [SDL_GetAudioStreamQueued](https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamQueued) | [:heavy_check_mark:](sdl/methods.go#L4017) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2012) |
| [SDL_FlushAudioStream](https://wiki.libsdl.org/SDL3/SDL_FlushAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4028) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2021) |
| [SDL_ClearAudioStream](https://wiki.libsdl.org/SDL3/SDL_ClearAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4038) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2030) |
| [SDL_PauseAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_PauseAudioStreamDevice) | [:heavy_check_mark:](sdl/methods.go#L4048) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2039) |
| [SDL_ResumeAudioStreamDevice](https://wiki.libsdl.org/SDL3/SDL_ResumeAudioStreamDevice) | [:heavy_check_mark:](sdl/methods.go#L4058) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2048) |
| [SDL_AudioStreamDevicePaused](https://wiki.libsdl.org/SDL3/SDL_AudioStreamDevicePaused) | [:heavy_check_mark:](sdl/methods.go#L4068) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2057) |
| [SDL_LockAudioStream](https://wiki.libsdl.org/SDL3/SDL_LockAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4074) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2066) |
| [SDL_UnlockAudioStream](https://wiki.libsdl.org/SDL3/SDL_UnlockAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4084) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2075) |
| [SDL_SetAudioStreamGetCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGetCallback) | [:heavy_check_mark:](sdl/methods.go#L4094) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11004) |
| [SDL_SetAudioStreamPutCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamPutCallback) | [:heavy_check_mark:](sdl/methods.go#L4100) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11015) |
| [SDL_DestroyAudioStream](https://wiki.libsdl.org/SDL3/SDL_DestroyAudioStream) | [:heavy_check_mark:](sdl/methods.go#L4106) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2084) |
| [SDL_OpenAudioDeviceStream](https://wiki.libsdl.org/SDL3/SDL_OpenAudioDeviceStream) | [:heavy_check_mark:](sdl/methods.go#L346) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11026) |
| [SDL_SetAudioPostmixCallback](https://wiki.libsdl.org/SDL3/SDL_SetAudioPostmixCallback) | [:heavy_check_mark:](sdl/methods.go#L352) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11039) |
| [SDL_LoadWAV_IO](https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2093) |
| [SDL_LoadWAV](https://wiki.libsdl.org/SDL3/SDL_LoadWAV) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2117) |
| [SDL_MixAudio](https://wiki.libsdl.org/SDL3/SDL_MixAudio) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2139) |
| [SDL_ConvertAudioSamples](https://wiki.libsdl.org/SDL3/SDL_ConvertAudioSamples) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2158) |
| [SDL_GetAudioFormatName](https://wiki.libsdl.org/SDL3/SDL_GetAudioFormatName) | [:heavy_check_mark:](sdl/methods.go#L1286) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2182) |
| [SDL_GetSilenceValueForFormat](https://wiki.libsdl.org/SDL3/SDL_GetSilenceValueForFormat) | [:heavy_check_mark:](sdl/methods.go#L1292) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2191) |
</details>
<details>
<summary><h3>Time</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetDateTimeLocalePreferences](https://wiki.libsdl.org/SDL3/SDL_GetDateTimeLocalePreferences) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10701) |
| [SDL_GetCurrentTime](https://wiki.libsdl.org/SDL3/SDL_GetCurrentTime) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10715) |
| [SDL_TimeToDateTime](https://wiki.libsdl.org/SDL3/SDL_TimeToDateTime) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10727) |
| [SDL_DateTimeToTime](https://wiki.libsdl.org/SDL3/SDL_DateTimeToTime) | [:heavy_check_mark:](sdl/methods.go#L5999) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10743) |
| [SDL_TimeToWindows](https://wiki.libsdl.org/SDL3/SDL_TimeToWindows) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10756) |
| [SDL_TimeFromWindows](https://wiki.libsdl.org/SDL3/SDL_TimeFromWindows) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10771) |
| [SDL_GetDaysInMonth](https://wiki.libsdl.org/SDL3/SDL_GetDaysInMonth) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10782) |
| [SDL_GetDayOfYear](https://wiki.libsdl.org/SDL3/SDL_GetDayOfYear) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10793) |
| [SDL_GetDayOfWeek](https://wiki.libsdl.org/SDL3/SDL_GetDayOfWeek) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10806) |
</details>
<details open>
<summary><h3>Timer</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetTicks](https://wiki.libsdl.org/SDL3/SDL_GetTicks) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10819) |
| [SDL_GetTicksNS](https://wiki.libsdl.org/SDL3/SDL_GetTicksNS) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10827) |
| [SDL_GetPerformanceCounter](https://wiki.libsdl.org/SDL3/SDL_GetPerformanceCounter) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10834) |
| [SDL_GetPerformanceFrequency](https://wiki.libsdl.org/SDL3/SDL_GetPerformanceFrequency) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10841) |
| [SDL_Delay](https://wiki.libsdl.org/SDL3/SDL_Delay) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10848) |
| [SDL_DelayNS](https://wiki.libsdl.org/SDL3/SDL_DelayNS) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10856) |
| [SDL_DelayPrecise](https://wiki.libsdl.org/SDL3/SDL_DelayPrecise) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10864) |
| [SDL_AddTimer](https://wiki.libsdl.org/SDL3/SDL_AddTimer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11756) |
| [SDL_AddTimerNS](https://wiki.libsdl.org/SDL3/SDL_AddTimerNS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11770) |
| [SDL_RemoveTimer](https://wiki.libsdl.org/SDL3/SDL_RemoveTimer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10872) |
</details>
<details open>
<summary><h3>Render</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumRenderDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumRenderDrivers) | [:heavy_check_mark:](sdl/functions.go#L420) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9008) |
| [SDL_GetRenderDriver](https://wiki.libsdl.org/SDL3/SDL_GetRenderDriver) | [:heavy_check_mark:](sdl/functions.go#L426) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9015) |
| [SDL_CreateWindowAndRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateWindowAndRenderer) | [:heavy_check_mark:](sdl/functions.go#L432) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9024) |
| [SDL_CreateRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateRenderer) | [:heavy_check_mark:](sdl/methods.go#L4898) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9052) |
| [SDL_CreateRendererWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateRendererWithProperties) | [:heavy_check_mark:](sdl/functions.go#L445) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9065) |
| [SDL_CreateGPURenderer](https://wiki.libsdl.org/SDL3/SDL_CreateGPURenderer) | [:heavy_check_mark:](sdl/methods.go#L2009) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11589) |
| [SDL_GetGPURendererDevice](https://wiki.libsdl.org/SDL3/SDL_GetGPURendererDevice) | [:heavy_check_mark:](sdl/methods.go#L2908) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11600) |
| [SDL_CreateSoftwareRenderer](https://wiki.libsdl.org/SDL3/SDL_CreateSoftwareRenderer) | [:heavy_check_mark:](sdl/methods.go#L1952) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9074) |
| [SDL_GetRenderer](https://wiki.libsdl.org/SDL3/SDL_GetRenderer) | [:heavy_check_mark:](sdl/methods.go#L4909) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9083) |
| [SDL_GetRenderWindow](https://wiki.libsdl.org/SDL3/SDL_GetRenderWindow) | [:heavy_check_mark:](sdl/methods.go#L2897) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9092) |
| [SDL_GetRendererName](https://wiki.libsdl.org/SDL3/SDL_GetRendererName) | [:heavy_check_mark:](sdl/methods.go#L2919) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9101) |
| [SDL_GetRendererProperties](https://wiki.libsdl.org/SDL3/SDL_GetRendererProperties) | [:heavy_check_mark:](sdl/methods.go#L2930) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9110) |
| [SDL_GetRenderOutputSize](https://wiki.libsdl.org/SDL3/SDL_GetRenderOutputSize) | [:heavy_check_mark:](sdl/methods.go#L2936) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9119) |
| [SDL_GetCurrentRenderOutputSize](https://wiki.libsdl.org/SDL3/SDL_GetCurrentRenderOutputSize) | [:heavy_check_mark:](sdl/methods.go#L2947) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9136) |
| [SDL_CreateTexture](https://wiki.libsdl.org/SDL3/SDL_CreateTexture) | [:heavy_check_mark:](sdl/methods.go#L2958) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9153) |
| [SDL_CreateTextureFromSurface](https://wiki.libsdl.org/SDL3/SDL_CreateTextureFromSurface) | [:heavy_check_mark:](sdl/methods.go#L2969) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9176) |
| [SDL_CreateTextureWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateTextureWithProperties) | [:heavy_check_mark:](sdl/methods.go#L2980) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9196) |
| [SDL_GetTextureProperties](https://wiki.libsdl.org/SDL3/SDL_GetTextureProperties) | [:heavy_check_mark:](sdl/methods.go#L993) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9207) |
| [SDL_GetRendererFromTexture](https://wiki.libsdl.org/SDL3/SDL_GetRendererFromTexture) | [:heavy_check_mark:](sdl/methods.go#L1004) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9216) |
| [SDL_GetTextureSize](https://wiki.libsdl.org/SDL3/SDL_GetTextureSize) | [:heavy_check_mark:](sdl/methods.go#L1015) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9225) |
| [SDL_SetTexturePalette](https://wiki.libsdl.org/SDL3/SDL_SetTexturePalette) | [:heavy_check_mark:](sdl/methods.go#L1026) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11609) |
| [SDL_GetTexturePalette](https://wiki.libsdl.org/SDL3/SDL_GetTexturePalette) | [:heavy_check_mark:](sdl/methods.go#L1036) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11620) |
| [SDL_SetTextureColorMod](https://wiki.libsdl.org/SDL3/SDL_SetTextureColorMod) | [:heavy_check_mark:](sdl/methods.go#L1042) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9242) |
| [SDL_SetTextureColorModFloat](https://wiki.libsdl.org/SDL3/SDL_SetTextureColorModFloat) | [:heavy_check_mark:](sdl/methods.go#L1052) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9257) |
| [SDL_GetTextureColorMod](https://wiki.libsdl.org/SDL3/SDL_GetTextureColorMod) | [:heavy_check_mark:](sdl/methods.go#L1062) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9273) |
| [SDL_GetTextureColorModFloat](https://wiki.libsdl.org/SDL3/SDL_GetTextureColorModFloat) | [:heavy_check_mark:](sdl/methods.go#L1073) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9293) |
| [SDL_SetTextureAlphaMod](https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1084) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9313) |
| [SDL_SetTextureAlphaModFloat](https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaModFloat) | [:heavy_check_mark:](sdl/methods.go#L1094) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9324) |
| [SDL_GetTextureAlphaMod](https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1104) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9335) |
| [SDL_GetTextureAlphaModFloat](https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaModFloat) | [:heavy_check_mark:](sdl/methods.go#L1115) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9349) |
| [SDL_SetTextureBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetTextureBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1126) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9363) |
| [SDL_GetTextureBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetTextureBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1136) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9374) |
| [SDL_SetTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_SetTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L1147) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9388) |
| [SDL_GetTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_GetTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L1157) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9399) |
| [SDL_UpdateTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateTexture) | [:heavy_check_mark:](sdl/methods.go#L1168) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9413) |
| [SDL_UpdateYUVTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateYUVTexture) | [:heavy_check_mark:](sdl/methods.go#L1180) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9437) |
| [SDL_UpdateNVTexture](https://wiki.libsdl.org/SDL3/SDL_UpdateNVTexture) | [:heavy_check_mark:](sdl/methods.go#L1194) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9466) |
| [SDL_LockTexture](https://wiki.libsdl.org/SDL3/SDL_LockTexture) | [:heavy_check_mark:](sdl/methods.go#L1208) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9491) |
| [SDL_LockTextureToSurface](https://wiki.libsdl.org/SDL3/SDL_LockTextureToSurface) | [:heavy_check_mark:](sdl/methods.go#L1221) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9513) |
| [SDL_UnlockTexture](https://wiki.libsdl.org/SDL3/SDL_UnlockTexture) | [:heavy_check_mark:](sdl/methods.go#L1231) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9535) |
| [SDL_SetRenderTarget](https://wiki.libsdl.org/SDL3/SDL_SetRenderTarget) | [:heavy_check_mark:](sdl/methods.go#L2991) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9547) |
| [SDL_GetRenderTarget](https://wiki.libsdl.org/SDL3/SDL_GetRenderTarget) | [:heavy_check_mark:](sdl/methods.go#L3001) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9558) |
| [SDL_SetRenderLogicalPresentation](https://wiki.libsdl.org/SDL3/SDL_SetRenderLogicalPresentation) | [:heavy_check_mark:](sdl/methods.go#L3007) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9567) |
| [SDL_GetRenderLogicalPresentation](https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentation) | [:heavy_check_mark:](sdl/methods.go#L3017) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9582) |
| [SDL_GetRenderLogicalPresentationRect](https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentationRect) | [:heavy_check_mark:](sdl/methods.go#L3029) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9602) |
| [SDL_RenderCoordinatesFromWindow](https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesFromWindow) | [:heavy_check_mark:](sdl/methods.go#L3041) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9616) |
| [SDL_RenderCoordinatesToWindow](https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesToWindow) | [:heavy_check_mark:](sdl/methods.go#L3053) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9635) |
| [SDL_ConvertEventToRenderCoordinates](https://wiki.libsdl.org/SDL3/SDL_ConvertEventToRenderCoordinates) | [:heavy_check_mark:](sdl/methods.go#L3065) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9654) |
| [SDL_SetRenderViewport](https://wiki.libsdl.org/SDL3/SDL_SetRenderViewport) | [:heavy_check_mark:](sdl/methods.go#L3075) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9668) |
| [SDL_GetRenderViewport](https://wiki.libsdl.org/SDL3/SDL_GetRenderViewport) | [:heavy_check_mark:](sdl/methods.go#L3085) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9685) |
| [SDL_RenderViewportSet](https://wiki.libsdl.org/SDL3/SDL_RenderViewportSet) | [:heavy_check_mark:](sdl/methods.go#L3097) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9699) |
| [SDL_GetRenderSafeArea](https://wiki.libsdl.org/SDL3/SDL_GetRenderSafeArea) | [:heavy_check_mark:](sdl/methods.go#L3103) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9708) |
| [SDL_SetRenderClipRect](https://wiki.libsdl.org/SDL3/SDL_SetRenderClipRect) | [:heavy_check_mark:](sdl/methods.go#L3115) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9722) |
| [SDL_GetRenderClipRect](https://wiki.libsdl.org/SDL3/SDL_GetRenderClipRect) | [:heavy_check_mark:](sdl/methods.go#L3125) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9739) |
| [SDL_RenderClipEnabled](https://wiki.libsdl.org/SDL3/SDL_RenderClipEnabled) | [:heavy_check_mark:](sdl/methods.go#L3137) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9753) |
| [SDL_SetRenderScale](https://wiki.libsdl.org/SDL3/SDL_SetRenderScale) | [:heavy_check_mark:](sdl/methods.go#L3147) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9762) |
| [SDL_GetRenderScale](https://wiki.libsdl.org/SDL3/SDL_GetRenderScale) | [:heavy_check_mark:](sdl/methods.go#L3157) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9779) |
| [SDL_SetRenderDrawColor](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColor) | [:heavy_check_mark:](sdl/methods.go#L3169) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9796) |
| [SDL_SetRenderDrawColorFloat](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColorFloat) | [:heavy_check_mark:](sdl/methods.go#L3179) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9818) |
| [SDL_GetRenderDrawColor](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColor) | [:heavy_check_mark:](sdl/methods.go#L3189) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9835) |
| [SDL_GetRenderDrawColorFloat](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColorFloat) | [:heavy_check_mark:](sdl/methods.go#L3201) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9858) |
| [SDL_SetRenderColorScale](https://wiki.libsdl.org/SDL3/SDL_SetRenderColorScale) | [:heavy_check_mark:](sdl/methods.go#L3213) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9881) |
| [SDL_GetRenderColorScale](https://wiki.libsdl.org/SDL3/SDL_GetRenderColorScale) | [:heavy_check_mark:](sdl/methods.go#L3223) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9892) |
| [SDL_SetRenderDrawBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawBlendMode) | [:heavy_check_mark:](sdl/methods.go#L3235) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9906) |
| [SDL_GetRenderDrawBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawBlendMode) | [:heavy_check_mark:](sdl/methods.go#L3245) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9921) |
| [SDL_RenderClear](https://wiki.libsdl.org/SDL3/SDL_RenderClear) | [:heavy_check_mark:](sdl/methods.go#L3257) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9935) |
| [SDL_RenderPoint](https://wiki.libsdl.org/SDL3/SDL_RenderPoint) | [:heavy_check_mark:](sdl/methods.go#L3266) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9948) |
| [SDL_RenderPoints](https://wiki.libsdl.org/SDL3/SDL_RenderPoints) | [:heavy_check_mark:](sdl/methods.go#L3276) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9961) |
| [SDL_RenderLine](https://wiki.libsdl.org/SDL3/SDL_RenderLine) | [:heavy_check_mark:](sdl/methods.go#L3286) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9979) |
| [SDL_RenderLines](https://wiki.libsdl.org/SDL3/SDL_RenderLines) | [:heavy_check_mark:](sdl/methods.go#L3296) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L9996) |
| [SDL_RenderRect](https://wiki.libsdl.org/SDL3/SDL_RenderRect) | [:heavy_check_mark:](sdl/methods.go#L3306) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10014) |
| [SDL_RenderRects](https://wiki.libsdl.org/SDL3/SDL_RenderRects) | [:heavy_check_mark:](sdl/methods.go#L3316) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10031) |
| [SDL_RenderFillRect](https://wiki.libsdl.org/SDL3/SDL_RenderFillRect) | [:heavy_check_mark:](sdl/methods.go#L3326) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10049) |
| [SDL_RenderFillRects](https://wiki.libsdl.org/SDL3/SDL_RenderFillRects) | [:heavy_check_mark:](sdl/methods.go#L3336) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10066) |
| [SDL_RenderTexture](https://wiki.libsdl.org/SDL3/SDL_RenderTexture) | [:heavy_check_mark:](sdl/methods.go#L3346) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10084) |
| [SDL_RenderTextureRotated](https://wiki.libsdl.org/SDL3/SDL_RenderTextureRotated) | [:heavy_check_mark:](sdl/methods.go#L3356) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10108) |
| [SDL_RenderTextureAffine](https://wiki.libsdl.org/SDL3/SDL_RenderTextureAffine) | [:heavy_check_mark:](sdl/methods.go#L3366) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10138) |
| [SDL_RenderTextureTiled](https://wiki.libsdl.org/SDL3/SDL_RenderTextureTiled) | [:heavy_check_mark:](sdl/methods.go#L3376) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10166) |
| [SDL_RenderTexture9Grid](https://wiki.libsdl.org/SDL3/SDL_RenderTexture9Grid) | [:heavy_check_mark:](sdl/methods.go#L3386) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10184) |
| [SDL_RenderTexture9GridTiled](https://wiki.libsdl.org/SDL3/SDL_RenderTexture9GridTiled) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11629) |
| [SDL_RenderGeometry](https://wiki.libsdl.org/SDL3/SDL_RenderGeometry) | [:heavy_check_mark:](sdl/methods.go#L3396) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10206) |
| [SDL_RenderGeometryRaw](https://wiki.libsdl.org/SDL3/SDL_RenderGeometryRaw) | [:heavy_check_mark:](sdl/methods.go#L3406) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10233) |
| [SDL_SetRenderTextureAddressMode](https://wiki.libsdl.org/SDL3/SDL_SetRenderTextureAddressMode) | [:heavy_check_mark:](sdl/methods.go#L3427) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11652) |
| [SDL_GetRenderTextureAddressMode](https://wiki.libsdl.org/SDL3/SDL_GetRenderTextureAddressMode) | [:heavy_check_mark:](sdl/methods.go#L3437) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11665) |
| [SDL_RenderReadPixels](https://wiki.libsdl.org/SDL3/SDL_RenderReadPixels) | [:heavy_check_mark:](sdl/methods.go#L3474) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10265) |
| [SDL_RenderPresent](https://wiki.libsdl.org/SDL3/SDL_RenderPresent) | [:heavy_check_mark:](sdl/methods.go#L3485) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10284) |
| [SDL_DestroyTexture](https://wiki.libsdl.org/SDL3/SDL_DestroyTexture) | [:heavy_check_mark:](sdl/methods.go#L1237) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10297) |
| [SDL_DestroyRenderer](https://wiki.libsdl.org/SDL3/SDL_DestroyRenderer) | [:heavy_check_mark:](sdl/methods.go#L3495) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10310) |
| [SDL_FlushRenderer](https://wiki.libsdl.org/SDL3/SDL_FlushRenderer) | [:heavy_check_mark:](sdl/methods.go#L3501) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10322) |
| [SDL_GetRenderMetalLayer](https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalLayer) | [:x:](sdl/methods.go#L3511) | [:x:](sdl/sdl_functions_js.go#L10331) |
| [SDL_GetRenderMetalCommandEncoder](https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalCommandEncoder) | [:x:](sdl/methods.go#L3518) | [:x:](sdl/sdl_functions_js.go#L10335) |
| [SDL_AddVulkanRenderSemaphores](https://wiki.libsdl.org/SDL3/SDL_AddVulkanRenderSemaphores) | [:heavy_check_mark:](sdl/methods.go#L3525) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10339) |
| [SDL_SetRenderVSync](https://wiki.libsdl.org/SDL3/SDL_SetRenderVSync) | [:heavy_check_mark:](sdl/methods.go#L3535) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10354) |
| [SDL_GetRenderVSync](https://wiki.libsdl.org/SDL3/SDL_GetRenderVSync) | [:heavy_check_mark:](sdl/methods.go#L3545) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10365) |
| [SDL_RenderDebugText](https://wiki.libsdl.org/SDL3/SDL_RenderDebugText) | [:heavy_check_mark:](sdl/methods.go#L3567) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10379) |
| [SDL_RenderDebugTextFormat](https://wiki.libsdl.org/SDL3/SDL_RenderDebugTextFormat) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10399) |
| [SDL_SetDefaultTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_SetDefaultTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L3573) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11682) |
| [SDL_GetDefaultTextureScaleMode](https://wiki.libsdl.org/SDL3/SDL_GetDefaultTextureScaleMode) | [:heavy_check_mark:](sdl/methods.go#L3583) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11693) |
| [SDL_CreateGPURenderState](https://wiki.libsdl.org/SDL3/SDL_CreateGPURenderState) | [:heavy_check_mark:](sdl/methods.go#L3449) | [:no_entry_sign:](sdl/sdl_functions_js.go#L11993) |
| [SDL_SetGPURenderStateSamplerBindings](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderStateSamplerBindings) | [:question:]() | [:question:]() |
| [SDL_SetGPURenderStateStorageTextures](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderStateStorageTextures) | [:question:]() | [:question:]() |
| [SDL_SetGPURenderStateStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderStateStorageBuffers) | [:question:]() | [:question:]() |
| [SDL_SetGPURenderStateFragmentUniforms](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderStateFragmentUniforms) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11707) |
| [SDL_SetGPURenderState](https://wiki.libsdl.org/SDL3/SDL_SetGPURenderState) | [:heavy_check_mark:](sdl/methods.go#L3462) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11711) |
| [SDL_DestroyGPURenderState](https://wiki.libsdl.org/SDL3/SDL_DestroyGPURenderState) | [:heavy_check_mark:](sdl/methods.go#L3597) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11722) |
| [SDL_GDKSuspendRenderer](https://wiki.libsdl.org/SDL3/SDL_GDKSuspendRenderer) | [:question:]() | [:question:]() |
| [SDL_GDKResumeRenderer](https://wiki.libsdl.org/SDL3/SDL_GDKResumeRenderer) | [:question:]() | [:question:]() |
</details>
<details>
<summary><h3>SharedObject</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_LoadObject](https://wiki.libsdl.org/SDL3/SDL_LoadObject) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8674) |
| [SDL_LoadFunction](https://wiki.libsdl.org/SDL3/SDL_LoadFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11552) |
| [SDL_UnloadObject](https://wiki.libsdl.org/SDL3/SDL_UnloadObject) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8685) |
</details>
<details>
<summary><h3>Thread</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateThread](https://wiki.libsdl.org/SDL3/SDL_CreateThread) | [:question:]() | [:question:]() |
| [SDL_CreateThreadWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateThreadWithProperties) | [:question:]() | [:question:]() |
| [SDL_GetThreadName](https://wiki.libsdl.org/SDL3/SDL_GetThreadName) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L669) |
| [SDL_GetCurrentThreadID](https://wiki.libsdl.org/SDL3/SDL_GetCurrentThreadID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L678) |
| [SDL_GetThreadID](https://wiki.libsdl.org/SDL3/SDL_GetThreadID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L685) |
| [SDL_SetCurrentThreadPriority](https://wiki.libsdl.org/SDL3/SDL_SetCurrentThreadPriority) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L694) |
| [SDL_WaitThread](https://wiki.libsdl.org/SDL3/SDL_WaitThread) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L703) |
| [SDL_GetThreadState](https://wiki.libsdl.org/SDL3/SDL_GetThreadState) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L707) |
| [SDL_DetachThread](https://wiki.libsdl.org/SDL3/SDL_DetachThread) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L716) |
| [SDL_GetTLS](https://wiki.libsdl.org/SDL3/SDL_GetTLS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L724) |
| [SDL_SetTLS](https://wiki.libsdl.org/SDL3/SDL_SetTLS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10952) |
| [SDL_CleanupTLS](https://wiki.libsdl.org/SDL3/SDL_CleanupTLS) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L736) |
</details>
<details>
<summary><h3>Mutex</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateMutex](https://wiki.libsdl.org/SDL3/SDL_CreateMutex) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L742) |
| [SDL_LockMutex](https://wiki.libsdl.org/SDL3/SDL_LockMutex) | [:heavy_check_mark:](sdl/methods.go#L5763) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L749) |
| [SDL_TryLockMutex](https://wiki.libsdl.org/SDL3/SDL_TryLockMutex) | [:heavy_check_mark:](sdl/methods.go#L5769) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L757) |
| [SDL_UnlockMutex](https://wiki.libsdl.org/SDL3/SDL_UnlockMutex) | [:heavy_check_mark:](sdl/methods.go#L5775) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L766) |
| [SDL_DestroyMutex](https://wiki.libsdl.org/SDL3/SDL_DestroyMutex) | [:heavy_check_mark:](sdl/methods.go#L5781) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L774) |
| [SDL_CreateRWLock](https://wiki.libsdl.org/SDL3/SDL_CreateRWLock) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L783) |
| [SDL_LockRWLockForReading](https://wiki.libsdl.org/SDL3/SDL_LockRWLockForReading) | [:heavy_check_mark:](sdl/methods.go#L1248) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L790) |
| [SDL_LockRWLockForWriting](https://wiki.libsdl.org/SDL3/SDL_LockRWLockForWriting) | [:heavy_check_mark:](sdl/methods.go#L1254) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L798) |
| [SDL_TryLockRWLockForReading](https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForReading) | [:heavy_check_mark:](sdl/methods.go#L1260) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L806) |
| [SDL_TryLockRWLockForWriting](https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForWriting) | [:heavy_check_mark:](sdl/methods.go#L1266) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L815) |
| [SDL_UnlockRWLock](https://wiki.libsdl.org/SDL3/SDL_UnlockRWLock) | [:heavy_check_mark:](sdl/methods.go#L1272) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L824) |
| [SDL_DestroyRWLock](https://wiki.libsdl.org/SDL3/SDL_DestroyRWLock) | [:heavy_check_mark:](sdl/methods.go#L1278) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L832) |
| [SDL_CreateSemaphore](https://wiki.libsdl.org/SDL3/SDL_CreateSemaphore) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L841) |
| [SDL_DestroySemaphore](https://wiki.libsdl.org/SDL3/SDL_DestroySemaphore) | [:heavy_check_mark:](sdl/methods.go#L5934) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L850) |
| [SDL_WaitSemaphore](https://wiki.libsdl.org/SDL3/SDL_WaitSemaphore) | [:heavy_check_mark:](sdl/methods.go#L5940) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L859) |
| [SDL_TryWaitSemaphore](https://wiki.libsdl.org/SDL3/SDL_TryWaitSemaphore) | [:heavy_check_mark:](sdl/methods.go#L5946) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L867) |
| [SDL_WaitSemaphoreTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitSemaphoreTimeout) | [:heavy_check_mark:](sdl/methods.go#L5952) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L876) |
| [SDL_SignalSemaphore](https://wiki.libsdl.org/SDL3/SDL_SignalSemaphore) | [:heavy_check_mark:](sdl/methods.go#L5958) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L887) |
| [SDL_GetSemaphoreValue](https://wiki.libsdl.org/SDL3/SDL_GetSemaphoreValue) | [:heavy_check_mark:](sdl/methods.go#L5964) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L895) |
| [SDL_CreateCondition](https://wiki.libsdl.org/SDL3/SDL_CreateCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L904) |
| [SDL_DestroyCondition](https://wiki.libsdl.org/SDL3/SDL_DestroyCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L911) |
| [SDL_SignalCondition](https://wiki.libsdl.org/SDL3/SDL_SignalCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L920) |
| [SDL_BroadcastCondition](https://wiki.libsdl.org/SDL3/SDL_BroadcastCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L928) |
| [SDL_WaitCondition](https://wiki.libsdl.org/SDL3/SDL_WaitCondition) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L936) |
| [SDL_WaitConditionTimeout](https://wiki.libsdl.org/SDL3/SDL_WaitConditionTimeout) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L946) |
| [SDL_ShouldInit](https://wiki.libsdl.org/SDL3/SDL_ShouldInit) | [:heavy_check_mark:](sdl/methods.go#L3642) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L959) |
| [SDL_ShouldQuit](https://wiki.libsdl.org/SDL3/SDL_ShouldQuit) | [:heavy_check_mark:](sdl/methods.go#L3648) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L971) |
| [SDL_SetInitialized](https://wiki.libsdl.org/SDL3/SDL_SetInitialized) | [:heavy_check_mark:](sdl/methods.go#L3654) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L983) |
</details>
<details>
<summary><h3>Atomic</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_TryLockSpinlock](https://wiki.libsdl.org/SDL3/SDL_TryLockSpinlock) | [:heavy_check_mark:](sdl/methods.go#L573) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L181) |
| [SDL_LockSpinlock](https://wiki.libsdl.org/SDL3/SDL_LockSpinlock) | [:heavy_check_mark:](sdl/methods.go#L579) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L193) |
| [SDL_UnlockSpinlock](https://wiki.libsdl.org/SDL3/SDL_UnlockSpinlock) | [:heavy_check_mark:](sdl/methods.go#L585) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L204) |
| [SDL_MemoryBarrierReleaseFunction](https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierReleaseFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L215) |
| [SDL_MemoryBarrierAcquireFunction](https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierAcquireFunction) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L221) |
| [SDL_CompareAndSwapAtomicInt](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L920) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L227) |
| [SDL_SetAtomicInt](https://wiki.libsdl.org/SDL3/SDL_SetAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L926) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L241) |
| [SDL_GetAtomicInt](https://wiki.libsdl.org/SDL3/SDL_GetAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L932) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L254) |
| [SDL_AddAtomicInt](https://wiki.libsdl.org/SDL3/SDL_AddAtomicInt) | [:heavy_check_mark:](sdl/methods.go#L938) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L266) |
| [SDL_CompareAndSwapAtomicU32](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicU32) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L279) |
| [SDL_SetAtomicU32](https://wiki.libsdl.org/SDL3/SDL_SetAtomicU32) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L295) |
| [SDL_GetAtomicU32](https://wiki.libsdl.org/SDL3/SDL_GetAtomicU32) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L309) |
| [SDL_AddAtomicU32](https://wiki.libsdl.org/SDL3/SDL_AddAtomicU32) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10918) |
| [SDL_CompareAndSwapAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicPointer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L321) |
| [SDL_SetAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_SetAtomicPointer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L337) |
| [SDL_GetAtomicPointer](https://wiki.libsdl.org/SDL3/SDL_GetAtomicPointer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L351) |
</details>
<details>
<summary><h3>Filesystem</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetBasePath](https://wiki.libsdl.org/SDL3/SDL_GetBasePath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7397) |
| [SDL_GetPrefPath](https://wiki.libsdl.org/SDL3/SDL_GetPrefPath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7404) |
| [SDL_GetUserFolder](https://wiki.libsdl.org/SDL3/SDL_GetUserFolder) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7418) |
| [SDL_CreateDirectory](https://wiki.libsdl.org/SDL3/SDL_CreateDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7427) |
| [SDL_EnumerateDirectory](https://wiki.libsdl.org/SDL3/SDL_EnumerateDirectory) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11468) |
| [SDL_RemovePath](https://wiki.libsdl.org/SDL3/SDL_RemovePath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7438) |
| [SDL_RenamePath](https://wiki.libsdl.org/SDL3/SDL_RenamePath) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7449) |
| [SDL_CopyFile](https://wiki.libsdl.org/SDL3/SDL_CopyFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7462) |
| [SDL_GetPathInfo](https://wiki.libsdl.org/SDL3/SDL_GetPathInfo) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7475) |
| [SDL_GlobDirectory](https://wiki.libsdl.org/SDL3/SDL_GlobDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7489) |
| [SDL_GetCurrentDirectory](https://wiki.libsdl.org/SDL3/SDL_GetCurrentDirectory) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L7507) |
</details>
<details open>
<summary><h3>IOStream</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_IOFromFile](https://wiki.libsdl.org/SDL3/SDL_IOFromFile) | [:heavy_check_mark:](sdl/functions.go#L370) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L995) |
| [SDL_IOFromMem](https://wiki.libsdl.org/SDL3/SDL_IOFromMem) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1008) |
| [SDL_IOFromConstMem](https://wiki.libsdl.org/SDL3/SDL_IOFromConstMem) | [:heavy_check_mark:](sdl/functions.go#L386) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1020) |
| [SDL_IOFromDynamicMem](https://wiki.libsdl.org/SDL3/SDL_IOFromDynamicMem) | [:heavy_check_mark:](sdl/functions.go#L402) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1040) |
| [SDL_OpenIO](https://wiki.libsdl.org/SDL3/SDL_OpenIO) | [:x:](sdl/methods.go#L24) | [:x:](sdl/sdl_functions_js.go#L1047) |
| [SDL_CloseIO](https://wiki.libsdl.org/SDL3/SDL_CloseIO) | [:heavy_check_mark:](sdl/methods.go#L4945) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1061) |
| [SDL_GetIOProperties](https://wiki.libsdl.org/SDL3/SDL_GetIOProperties) | [:heavy_check_mark:](sdl/methods.go#L4955) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1075) |
| [SDL_GetIOStatus](https://wiki.libsdl.org/SDL3/SDL_GetIOStatus) | [:heavy_check_mark:](sdl/methods.go#L4966) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1084) |
| [SDL_GetIOSize](https://wiki.libsdl.org/SDL3/SDL_GetIOSize) | [:heavy_check_mark:](sdl/methods.go#L4972) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1093) |
| [SDL_SeekIO](https://wiki.libsdl.org/SDL3/SDL_SeekIO) | [:heavy_check_mark:](sdl/methods.go#L4983) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1102) |
| [SDL_TellIO](https://wiki.libsdl.org/SDL3/SDL_TellIO) | [:heavy_check_mark:](sdl/methods.go#L4994) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1115) |
| [SDL_ReadIO](https://wiki.libsdl.org/SDL3/SDL_ReadIO) | [:heavy_check_mark:](sdl/methods.go#L5000) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1124) |
| [SDL_WriteIO](https://wiki.libsdl.org/SDL3/SDL_WriteIO) | [:heavy_check_mark:](sdl/methods.go#L5011) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1142) |
| [SDL_IOprintf](https://wiki.libsdl.org/SDL3/SDL_IOprintf) | [:heavy_check_mark:](sdl/methods.go#L5022) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1155) |
| [SDL_IOvprintf](https://wiki.libsdl.org/SDL3/SDL_IOvprintf) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1167) |
| [SDL_FlushIO](https://wiki.libsdl.org/SDL3/SDL_FlushIO) | [:heavy_check_mark:](sdl/methods.go#L5036) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1183) |
| [SDL_LoadFile_IO](https://wiki.libsdl.org/SDL3/SDL_LoadFile_IO) | [:heavy_check_mark:](sdl/methods.go#L5046) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1192) |
| [SDL_LoadFile](https://wiki.libsdl.org/SDL3/SDL_LoadFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1207) |
| [SDL_SaveFile_IO](https://wiki.libsdl.org/SDL3/SDL_SaveFile_IO) | [:heavy_check_mark:](sdl/methods.go#L5059) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1221) |
| [SDL_SaveFile](https://wiki.libsdl.org/SDL3/SDL_SaveFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L1235) |
| [SDL_ReadU8](https://wiki.libsdl.org/SDL3/SDL_ReadU8) | [:heavy_check_mark:](sdl/methods.go#L5070) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1251) |
| [SDL_ReadS8](https://wiki.libsdl.org/SDL3/SDL_ReadS8) | [:heavy_check_mark:](sdl/methods.go#L5082) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1265) |
| [SDL_ReadU16LE](https://wiki.libsdl.org/SDL3/SDL_ReadU16LE) | [:heavy_check_mark:](sdl/methods.go#L5094) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1279) |
| [SDL_ReadS16LE](https://wiki.libsdl.org/SDL3/SDL_ReadS16LE) | [:heavy_check_mark:](sdl/methods.go#L5106) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1293) |
| [SDL_ReadU16BE](https://wiki.libsdl.org/SDL3/SDL_ReadU16BE) | [:heavy_check_mark:](sdl/methods.go#L5118) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1307) |
| [SDL_ReadS16BE](https://wiki.libsdl.org/SDL3/SDL_ReadS16BE) | [:heavy_check_mark:](sdl/methods.go#L5130) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1321) |
| [SDL_ReadU32LE](https://wiki.libsdl.org/SDL3/SDL_ReadU32LE) | [:heavy_check_mark:](sdl/methods.go#L5142) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1335) |
| [SDL_ReadS32LE](https://wiki.libsdl.org/SDL3/SDL_ReadS32LE) | [:heavy_check_mark:](sdl/methods.go#L5154) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1349) |
| [SDL_ReadU32BE](https://wiki.libsdl.org/SDL3/SDL_ReadU32BE) | [:heavy_check_mark:](sdl/methods.go#L5166) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1363) |
| [SDL_ReadS32BE](https://wiki.libsdl.org/SDL3/SDL_ReadS32BE) | [:heavy_check_mark:](sdl/methods.go#L5178) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1377) |
| [SDL_ReadU64LE](https://wiki.libsdl.org/SDL3/SDL_ReadU64LE) | [:heavy_check_mark:](sdl/methods.go#L5190) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1391) |
| [SDL_ReadS64LE](https://wiki.libsdl.org/SDL3/SDL_ReadS64LE) | [:heavy_check_mark:](sdl/methods.go#L5202) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1405) |
| [SDL_ReadU64BE](https://wiki.libsdl.org/SDL3/SDL_ReadU64BE) | [:heavy_check_mark:](sdl/methods.go#L5214) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1419) |
| [SDL_ReadS64BE](https://wiki.libsdl.org/SDL3/SDL_ReadS64BE) | [:heavy_check_mark:](sdl/methods.go#L5226) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1433) |
| [SDL_WriteU8](https://wiki.libsdl.org/SDL3/SDL_WriteU8) | [:heavy_check_mark:](sdl/methods.go#L5238) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1447) |
| [SDL_WriteS8](https://wiki.libsdl.org/SDL3/SDL_WriteS8) | [:heavy_check_mark:](sdl/methods.go#L5248) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1458) |
| [SDL_WriteU16LE](https://wiki.libsdl.org/SDL3/SDL_WriteU16LE) | [:heavy_check_mark:](sdl/methods.go#L5258) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1469) |
| [SDL_WriteS16LE](https://wiki.libsdl.org/SDL3/SDL_WriteS16LE) | [:heavy_check_mark:](sdl/methods.go#L5268) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1480) |
| [SDL_WriteU16BE](https://wiki.libsdl.org/SDL3/SDL_WriteU16BE) | [:heavy_check_mark:](sdl/methods.go#L5278) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1491) |
| [SDL_WriteS16BE](https://wiki.libsdl.org/SDL3/SDL_WriteS16BE) | [:heavy_check_mark:](sdl/methods.go#L5288) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1502) |
| [SDL_WriteU32LE](https://wiki.libsdl.org/SDL3/SDL_WriteU32LE) | [:heavy_check_mark:](sdl/methods.go#L5298) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1513) |
| [SDL_WriteS32LE](https://wiki.libsdl.org/SDL3/SDL_WriteS32LE) | [:heavy_check_mark:](sdl/methods.go#L5308) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1524) |
| [SDL_WriteU32BE](https://wiki.libsdl.org/SDL3/SDL_WriteU32BE) | [:heavy_check_mark:](sdl/methods.go#L5318) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1535) |
| [SDL_WriteS32BE](https://wiki.libsdl.org/SDL3/SDL_WriteS32BE) | [:heavy_check_mark:](sdl/methods.go#L5328) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1546) |
| [SDL_WriteU64LE](https://wiki.libsdl.org/SDL3/SDL_WriteU64LE) | [:heavy_check_mark:](sdl/methods.go#L5338) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1557) |
| [SDL_WriteS64LE](https://wiki.libsdl.org/SDL3/SDL_WriteS64LE) | [:heavy_check_mark:](sdl/methods.go#L5348) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1568) |
| [SDL_WriteU64BE](https://wiki.libsdl.org/SDL3/SDL_WriteU64BE) | [:heavy_check_mark:](sdl/methods.go#L5358) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1579) |
| [SDL_WriteS64BE](https://wiki.libsdl.org/SDL3/SDL_WriteS64BE) | [:heavy_check_mark:](sdl/methods.go#L5368) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L1590) |
</details>
<details open>
<summary><h3>AsyncIO</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_AsyncIOFromFile](https://wiki.libsdl.org/SDL3/SDL_AsyncIOFromFile) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L30) |
| [SDL_GetAsyncIOSize](https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOSize) | [:heavy_check_mark:](sdl/methods.go#L3605) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L43) |
| [SDL_ReadAsyncIO](https://wiki.libsdl.org/SDL3/SDL_ReadAsyncIO) | [:x:](sdl/methods.go#L3616) | [:x:](sdl/sdl_functions_js.go#L52) |
| [SDL_WriteAsyncIO](https://wiki.libsdl.org/SDL3/SDL_WriteAsyncIO) | [:x:](sdl/methods.go#L3623) | [:x:](sdl/sdl_functions_js.go#L72) |
| [SDL_CloseAsyncIO](https://wiki.libsdl.org/SDL3/SDL_CloseAsyncIO) | [:heavy_check_mark:](sdl/methods.go#L3630) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L92) |
| [SDL_CreateAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_CreateAsyncIOQueue) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L106) |
| [SDL_DestroyAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_DestroyAsyncIOQueue) | [:heavy_check_mark:](sdl/methods.go#L1411) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L113) |
| [SDL_GetAsyncIOResult](https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOResult) | [:heavy_check_mark:](sdl/methods.go#L1417) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L122) |
| [SDL_WaitAsyncIOResult](https://wiki.libsdl.org/SDL3/SDL_WaitAsyncIOResult) | [:heavy_check_mark:](sdl/methods.go#L1427) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L139) |
| [SDL_SignalAsyncIOQueue](https://wiki.libsdl.org/SDL3/SDL_SignalAsyncIOQueue) | [:heavy_check_mark:](sdl/methods.go#L1437) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L157) |
| [SDL_LoadFileAsync](https://wiki.libsdl.org/SDL3/SDL_LoadFileAsync) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L165) |
</details>
<details open>
<summary><h3>Storage</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_OpenTitleStorage](https://wiki.libsdl.org/SDL3/SDL_OpenTitleStorage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10416) |
| [SDL_OpenUserStorage](https://wiki.libsdl.org/SDL3/SDL_OpenUserStorage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10429) |
| [SDL_OpenFileStorage](https://wiki.libsdl.org/SDL3/SDL_OpenFileStorage) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10444) |
| [SDL_OpenStorage](https://wiki.libsdl.org/SDL3/SDL_OpenStorage) | [:x:](sdl/methods.go#L5754) | [:x:](sdl/sdl_functions_js.go#L10455) |
| [SDL_CloseStorage](https://wiki.libsdl.org/SDL3/SDL_CloseStorage) | [:heavy_check_mark:](sdl/methods.go#L79) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10469) |
| [SDL_StorageReady](https://wiki.libsdl.org/SDL3/SDL_StorageReady) | [:heavy_check_mark:](sdl/methods.go#L89) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10479) |
| [SDL_GetStorageFileSize](https://wiki.libsdl.org/SDL3/SDL_GetStorageFileSize) | [:heavy_check_mark:](sdl/methods.go#L95) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10488) |
| [SDL_ReadStorageFile](https://wiki.libsdl.org/SDL3/SDL_ReadStorageFile) | [:heavy_check_mark:](sdl/methods.go#L107) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10503) |
| [SDL_WriteStorageFile](https://wiki.libsdl.org/SDL3/SDL_WriteStorageFile) | [:heavy_check_mark:](sdl/methods.go#L119) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10526) |
| [SDL_CreateStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_CreateStorageDirectory) | [:heavy_check_mark:](sdl/methods.go#L129) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10542) |
| [SDL_EnumerateStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_EnumerateStorageDirectory) | [:heavy_check_mark:](sdl/methods.go#L139) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11731) |
| [SDL_RemoveStoragePath](https://wiki.libsdl.org/SDL3/SDL_RemoveStoragePath) | [:heavy_check_mark:](sdl/methods.go#L149) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10555) |
| [SDL_RenameStoragePath](https://wiki.libsdl.org/SDL3/SDL_RenameStoragePath) | [:heavy_check_mark:](sdl/methods.go#L159) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10568) |
| [SDL_CopyStorageFile](https://wiki.libsdl.org/SDL3/SDL_CopyStorageFile) | [:heavy_check_mark:](sdl/methods.go#L169) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10583) |
| [SDL_GetStoragePathInfo](https://wiki.libsdl.org/SDL3/SDL_GetStoragePathInfo) | [:heavy_check_mark:](sdl/methods.go#L179) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10598) |
| [SDL_GetStorageSpaceRemaining](https://wiki.libsdl.org/SDL3/SDL_GetStorageSpaceRemaining) | [:heavy_check_mark:](sdl/methods.go#L191) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10618) |
| [SDL_GlobStorageDirectory](https://wiki.libsdl.org/SDL3/SDL_GlobStorageDirectory) | [:heavy_check_mark:](sdl/methods.go#L197) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10627) |
</details>
<details open>
<summary><h3>Pixels</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPixelFormatName](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatName) | [:heavy_check_mark:](sdl/methods.go#L5972) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2219) |
| [SDL_GetMasksForPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetMasksForPixelFormat) | [:x:](sdl/methods.go#L5978) | [:x:](sdl/sdl_functions_js.go#L2228) |
| [SDL_GetPixelFormatForMasks](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatForMasks) | [:heavy_check_mark:](sdl/functions.go#L458) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2250) |
| [SDL_GetPixelFormatDetails](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatDetails) | [:heavy_check_mark:](sdl/methods.go#L5986) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2263) |
| [SDL_CreatePalette](https://wiki.libsdl.org/SDL3/SDL_CreatePalette) | [:heavy_check_mark:](sdl/functions.go#L464) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2275) |
| [SDL_SetPaletteColors](https://wiki.libsdl.org/SDL3/SDL_SetPaletteColors) | [:heavy_check_mark:](sdl/methods.go#L5380) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2284) |
| [SDL_DestroyPalette](https://wiki.libsdl.org/SDL3/SDL_DestroyPalette) | [:heavy_check_mark:](sdl/methods.go#L5390) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2304) |
| [SDL_MapRGB](https://wiki.libsdl.org/SDL3/SDL_MapRGB) | [:heavy_check_mark:](sdl/methods.go#L1445) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2316) |
| [SDL_MapRGBA](https://wiki.libsdl.org/SDL3/SDL_MapRGBA) | [:heavy_check_mark:](sdl/methods.go#L1451) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2343) |
| [SDL_GetRGB](https://wiki.libsdl.org/SDL3/SDL_GetRGB) | [:heavy_check_mark:](sdl/functions.go#L487) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2364) |
| [SDL_GetRGBA](https://wiki.libsdl.org/SDL3/SDL_GetRGBA) | [:heavy_check_mark:](sdl/functions.go#L495) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2395) |
</details>
<details open>
<summary><h3>Surface</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateSurface](https://wiki.libsdl.org/SDL3/SDL_CreateSurface) | [:heavy_check_mark:](sdl/functions.go#L505) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2642) |
| [SDL_CreateSurfaceFrom](https://wiki.libsdl.org/SDL3/SDL_CreateSurfaceFrom) | [:heavy_check_mark:](sdl/functions.go#L516) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2642) |
| [SDL_DestroySurface](https://wiki.libsdl.org/SDL3/SDL_DestroySurface) | [:heavy_check_mark:](sdl/methods.go#L1459) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2675) |
| [SDL_GetSurfaceProperties](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceProperties) | [:heavy_check_mark:](sdl/methods.go#L1468) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2687) |
| [SDL_SetSurfaceColorspace](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorspace) | [:heavy_check_mark:](sdl/methods.go#L1479) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2700) |
| [SDL_GetSurfaceColorspace](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorspace) | [:heavy_check_mark:](sdl/methods.go#L1489) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2711) |
| [SDL_CreateSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_CreateSurfacePalette) | [:heavy_check_mark:](sdl/methods.go#L1495) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2720) |
| [SDL_SetSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_SetSurfacePalette) | [:heavy_check_mark:](sdl/methods.go#L1506) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2733) |
| [SDL_GetSurfacePalette](https://wiki.libsdl.org/SDL3/SDL_GetSurfacePalette) | [:heavy_check_mark:](sdl/methods.go#L1516) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2744) |
| [SDL_AddSurfaceAlternateImage](https://wiki.libsdl.org/SDL3/SDL_AddSurfaceAlternateImage) | [:heavy_check_mark:](sdl/methods.go#L1522) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2757) |
| [SDL_SurfaceHasAlternateImages](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasAlternateImages) | [:heavy_check_mark:](sdl/methods.go#L1532) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2768) |
| [SDL_GetSurfaceImages](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceImages) | [:heavy_check_mark:](sdl/methods.go#L1538) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2781) |
| [SDL_RemoveSurfaceAlternateImages](https://wiki.libsdl.org/SDL3/SDL_RemoveSurfaceAlternateImages) | [:heavy_check_mark:](sdl/methods.go#L1552) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2795) |
| [SDL_LockSurface](https://wiki.libsdl.org/SDL3/SDL_LockSurface) | [:heavy_check_mark:](sdl/methods.go#L1558) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2806) |
| [SDL_UnlockSurface](https://wiki.libsdl.org/SDL3/SDL_UnlockSurface) | [:heavy_check_mark:](sdl/methods.go#L1568) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2819) |
| [SDL_LoadSurface_IO](https://wiki.libsdl.org/SDL3/SDL_LoadSurface_IO) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11049) |
| [SDL_LoadSurface](https://wiki.libsdl.org/SDL3/SDL_LoadSurface) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11060) |
| [SDL_LoadBMP_IO](https://wiki.libsdl.org/SDL3/SDL_LoadBMP_IO) | [:heavy_check_mark:](sdl/functions.go#L528) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2830) |
| [SDL_LoadBMP](https://wiki.libsdl.org/SDL3/SDL_LoadBMP) | [:heavy_check_mark:](sdl/functions.go#L539) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2850) |
| [SDL_SaveBMP_IO](https://wiki.libsdl.org/SDL3/SDL_SaveBMP_IO) | [:heavy_check_mark:](sdl/methods.go#L1574) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2862) |
| [SDL_SaveBMP](https://wiki.libsdl.org/SDL3/SDL_SaveBMP) | [:heavy_check_mark:](sdl/methods.go#L1584) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2875) |
| [SDL_LoadPNG_IO](https://wiki.libsdl.org/SDL3/SDL_LoadPNG_IO) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11071) |
| [SDL_LoadPNG](https://wiki.libsdl.org/SDL3/SDL_LoadPNG) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11082) |
| [SDL_SavePNG_IO](https://wiki.libsdl.org/SDL3/SDL_SavePNG_IO) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11093) |
| [SDL_SavePNG](https://wiki.libsdl.org/SDL3/SDL_SavePNG) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11106) |
| [SDL_LoadJPG_IO](https://wiki.libsdl.org/SDL3/SDL_LoadJPG_IO) | [:question:]() | [:question:]() |
| [SDL_LoadJPG](https://wiki.libsdl.org/SDL3/SDL_LoadJPG) | [:question:]() | [:question:]() |
| [SDL_SetSurfaceRLE](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceRLE) | [:heavy_check_mark:](sdl/methods.go#L1594) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2892) |
| [SDL_SurfaceHasRLE](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasRLE) | [:heavy_check_mark:](sdl/methods.go#L1604) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2903) |
| [SDL_SetSurfaceColorKey](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorKey) | [:heavy_check_mark:](sdl/methods.go#L1610) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2916) |
| [SDL_SurfaceHasColorKey](https://wiki.libsdl.org/SDL3/SDL_SurfaceHasColorKey) | [:heavy_check_mark:](sdl/methods.go#L1620) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2932) |
| [SDL_GetSurfaceColorKey](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorKey) | [:heavy_check_mark:](sdl/methods.go#L1626) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2941) |
| [SDL_SetSurfaceColorMod](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorMod) | [:heavy_check_mark:](sdl/methods.go#L1637) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2959) |
| [SDL_GetSurfaceColorMod](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorMod) | [:heavy_check_mark:](sdl/methods.go#L1647) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2975) |
| [SDL_SetSurfaceAlphaMod](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1659) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2995) |
| [SDL_GetSurfaceAlphaMod](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceAlphaMod) | [:heavy_check_mark:](sdl/methods.go#L1669) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3009) |
| [SDL_SetSurfaceBlendMode](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1681) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3027) |
| [SDL_GetSurfaceBlendMode](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceBlendMode) | [:heavy_check_mark:](sdl/methods.go#L1691) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3041) |
| [SDL_SetSurfaceClipRect](https://wiki.libsdl.org/SDL3/SDL_SetSurfaceClipRect) | [:heavy_check_mark:](sdl/methods.go#L1703) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3059) |
| [SDL_GetSurfaceClipRect](https://wiki.libsdl.org/SDL3/SDL_GetSurfaceClipRect) | [:heavy_check_mark:](sdl/methods.go#L1709) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3076) |
| [SDL_FlipSurface](https://wiki.libsdl.org/SDL3/SDL_FlipSurface) | [:heavy_check_mark:](sdl/methods.go#L1721) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3094) |
| [SDL_RotateSurface](https://wiki.libsdl.org/SDL3/SDL_RotateSurface) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11119) |
| [SDL_DuplicateSurface](https://wiki.libsdl.org/SDL3/SDL_DuplicateSurface) | [:heavy_check_mark:](sdl/methods.go#L1731) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3105) |
| [SDL_ScaleSurface](https://wiki.libsdl.org/SDL3/SDL_ScaleSurface) | [:heavy_check_mark:](sdl/methods.go#L1742) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3118) |
| [SDL_ConvertSurface](https://wiki.libsdl.org/SDL3/SDL_ConvertSurface) | [:heavy_check_mark:](sdl/methods.go#L1753) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3134) |
| [SDL_ConvertSurfaceAndColorspace](https://wiki.libsdl.org/SDL3/SDL_ConvertSurfaceAndColorspace) | [:heavy_check_mark:](sdl/methods.go#L1764) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3151) |
| [SDL_ConvertPixels](https://wiki.libsdl.org/SDL3/SDL_ConvertPixels) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3168) |
| [SDL_ConvertPixelsAndColorspace](https://wiki.libsdl.org/SDL3/SDL_ConvertPixelsAndColorspace) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3192) |
| [SDL_PremultiplyAlpha](https://wiki.libsdl.org/SDL3/SDL_PremultiplyAlpha) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3224) |
| [SDL_PremultiplySurfaceAlpha](https://wiki.libsdl.org/SDL3/SDL_PremultiplySurfaceAlpha) | [:heavy_check_mark:](sdl/methods.go#L1775) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3250) |
| [SDL_ClearSurface](https://wiki.libsdl.org/SDL3/SDL_ClearSurface) | [:heavy_check_mark:](sdl/methods.go#L1785) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3267) |
| [SDL_FillSurfaceRect](https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRect) | [:heavy_check_mark:](sdl/methods.go#L1795) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3284) |
| [SDL_FillSurfaceRects](https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRects) | [:heavy_check_mark:](sdl/methods.go#L1805) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3303) |
| [SDL_BlitSurface](https://wiki.libsdl.org/SDL3/SDL_BlitSurface) | [:heavy_check_mark:](sdl/methods.go#L1815) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3317) |
| [SDL_BlitSurfaceUnchecked](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUnchecked) | [:heavy_check_mark:](sdl/methods.go#L1825) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3341) |
| [SDL_BlitSurfaceScaled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceScaled) | [:heavy_check_mark:](sdl/methods.go#L1835) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3365) |
| [SDL_BlitSurfaceUncheckedScaled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUncheckedScaled) | [:heavy_check_mark:](sdl/methods.go#L1845) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3390) |
| [SDL_StretchSurface](https://wiki.libsdl.org/SDL3/SDL_StretchSurface) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11130) |
| [SDL_BlitSurfaceTiled](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiled) | [:heavy_check_mark:](sdl/methods.go#L1855) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3415) |
| [SDL_BlitSurfaceTiledWithScale](https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiledWithScale) | [:heavy_check_mark:](sdl/methods.go#L1865) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3439) |
| [SDL_BlitSurface9Grid](https://wiki.libsdl.org/SDL3/SDL_BlitSurface9Grid) | [:heavy_check_mark:](sdl/methods.go#L1875) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3467) |
| [SDL_MapSurfaceRGB](https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGB) | [:heavy_check_mark:](sdl/methods.go#L1885) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3503) |
| [SDL_MapSurfaceRGBA](https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGBA) | [:heavy_check_mark:](sdl/methods.go#L1891) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3518) |
| [SDL_ReadSurfacePixel](https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixel) | [:heavy_check_mark:](sdl/methods.go#L1897) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3535) |
| [SDL_ReadSurfacePixelFloat](https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixelFloat) | [:heavy_check_mark:](sdl/methods.go#L1909) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3566) |
| [SDL_WriteSurfacePixel](https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixel) | [:heavy_check_mark:](sdl/methods.go#L1921) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3595) |
| [SDL_WriteSurfacePixelFloat](https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixelFloat) | [:heavy_check_mark:](sdl/methods.go#L1931) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3616) |
</details>
<details open>
<summary><h3>BlendMode</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ComposeCustomBlendMode](https://wiki.libsdl.org/SDL3/SDL_ComposeCustomBlendMode) | [:heavy_check_mark:](sdl/functions.go#L556) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2200) |
</details>
<details open>
<summary><h3>Rect</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_RectToFRect](https://wiki.libsdl.org/SDL3/SDL_RectToFRect) | [:question:]() | [:question:]() |
| [SDL_PointInRect](https://wiki.libsdl.org/SDL3/SDL_PointInRect) | [:question:]() | [:question:]() |
| [SDL_RectEmpty](https://wiki.libsdl.org/SDL3/SDL_RectEmpty) | [:question:]() | [:question:]() |
| [SDL_RectsEqual](https://wiki.libsdl.org/SDL3/SDL_RectsEqual) | [:question:]() | [:question:]() |
| [SDL_HasRectIntersection](https://wiki.libsdl.org/SDL3/SDL_HasRectIntersection) | [:heavy_check_mark:](sdl/methods.go#L672) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2431) |
| [SDL_GetRectIntersection](https://wiki.libsdl.org/SDL3/SDL_GetRectIntersection) | [:heavy_check_mark:](sdl/methods.go#L678) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2444) |
| [SDL_GetRectUnion](https://wiki.libsdl.org/SDL3/SDL_GetRectUnion) | [:heavy_check_mark:](sdl/methods.go#L690) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2462) |
| [SDL_GetRectEnclosingPoints](https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPoints) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2480) |
| [SDL_GetRectAndLineIntersection](https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersection) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L2500) |
| [SDL_PointInRectFloat](https://wiki.libsdl.org/SDL3/SDL_PointInRectFloat) | [:question:]() | [:question:]() |
| [SDL_RectEmptyFloat](https://wiki.libsdl.org/SDL3/SDL_RectEmptyFloat) | [:question:]() | [:question:]() |
| [SDL_RectsEqualEpsilon](https://wiki.libsdl.org/SDL3/SDL_RectsEqualEpsilon) | [:question:]() | [:question:]() |
| [SDL_RectsEqualFloat](https://wiki.libsdl.org/SDL3/SDL_RectsEqualFloat) | [:question:]() | [:question:]() |
| [SDL_HasRectIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_HasRectIntersectionFloat) | [:heavy_check_mark:](sdl/methods.go#L4114) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L2531) |
| [SDL_GetRectIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectIntersectionFloat) | [:x:](sdl/methods.go#L4120) | [:x:](sdl/sdl_functions_js.go#L2544) |
| [SDL_GetRectUnionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectUnionFloat) | [:x:](sdl/methods.go#L4127) | [:x:](sdl/sdl_functions_js.go#L2562) |
| [SDL_GetRectEnclosingPointsFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPointsFloat) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L2580) |
| [SDL_GetRectAndLineIntersectionFloat](https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersectionFloat) | [:x:](sdl/methods.go#L4134) | [:x:](sdl/sdl_functions_js.go#L2600) |
</details>
<details open>
<summary><h3>Camera</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetNumCameraDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumCameraDrivers) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3635) |
| [SDL_GetCameraDriver](https://wiki.libsdl.org/SDL3/SDL_GetCameraDriver) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3643) |
| [SDL_GetCurrentCameraDriver](https://wiki.libsdl.org/SDL3/SDL_GetCurrentCameraDriver) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3652) |
| [SDL_GetCameras](https://wiki.libsdl.org/SDL3/SDL_GetCameras) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3660) |
| [SDL_GetCameraSupportedFormats](https://wiki.libsdl.org/SDL3/SDL_GetCameraSupportedFormats) | [:heavy_check_mark:](sdl/methods.go#L3662) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3672) |
| [SDL_GetCameraName](https://wiki.libsdl.org/SDL3/SDL_GetCameraName) | [:x:](sdl/methods.go#L3677) | [:x:](sdl/sdl_functions_js.go#L3685) |
| [SDL_GetCameraPosition](https://wiki.libsdl.org/SDL3/SDL_GetCameraPosition) | [:heavy_check_mark:](sdl/methods.go#L3684) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3694) |
| [SDL_OpenCamera](https://wiki.libsdl.org/SDL3/SDL_OpenCamera) | [:heavy_check_mark:](sdl/methods.go#L3690) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3703) |
| [SDL_GetCameraPermissionState](https://wiki.libsdl.org/SDL3/SDL_GetCameraPermissionState) | [:heavy_check_mark:](sdl/methods.go#L360) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3717) |
| [SDL_GetCameraID](https://wiki.libsdl.org/SDL3/SDL_GetCameraID) | [:heavy_check_mark:](sdl/methods.go#L366) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3730) |
| [SDL_GetCameraProperties](https://wiki.libsdl.org/SDL3/SDL_GetCameraProperties) | [:heavy_check_mark:](sdl/methods.go#L377) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3743) |
| [SDL_GetCameraFormat](https://wiki.libsdl.org/SDL3/SDL_GetCameraFormat) | [:heavy_check_mark:](sdl/methods.go#L388) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3756) |
| [SDL_AcquireCameraFrame](https://wiki.libsdl.org/SDL3/SDL_AcquireCameraFrame) | [:heavy_check_mark:](sdl/methods.go#L400) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3774) |
| [SDL_ReleaseCameraFrame](https://wiki.libsdl.org/SDL3/SDL_ReleaseCameraFrame) | [:heavy_check_mark:](sdl/methods.go#L413) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3794) |
| [SDL_CloseCamera](https://wiki.libsdl.org/SDL3/SDL_CloseCamera) | [:heavy_check_mark:](sdl/methods.go#L419) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3810) |
</details>
<details open>
<summary><h3>MessageBox</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ShowMessageBox](https://wiki.libsdl.org/SDL3/SDL_ShowMessageBox) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8889) |
| [SDL_ShowSimpleMessageBox](https://wiki.libsdl.org/SDL3/SDL_ShowSimpleMessageBox) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8904) |
</details>
<details open>
<summary><h3>Clipboard</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetClipboardText](https://wiki.libsdl.org/SDL3/SDL_SetClipboardText) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3822) |
| [SDL_GetClipboardText](https://wiki.libsdl.org/SDL3/SDL_GetClipboardText) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L3834) |
| [SDL_HasClipboardText](https://wiki.libsdl.org/SDL3/SDL_HasClipboardText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3841) |
| [SDL_SetPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_SetPrimarySelectionText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3848) |
| [SDL_GetPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_GetPrimarySelectionText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3859) |
| [SDL_HasPrimarySelectionText](https://wiki.libsdl.org/SDL3/SDL_HasPrimarySelectionText) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3867) |
| [SDL_SetClipboardData](https://wiki.libsdl.org/SDL3/SDL_SetClipboardData) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11970) |
| [SDL_ClearClipboardData](https://wiki.libsdl.org/SDL3/SDL_ClearClipboardData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3874) |
| [SDL_GetClipboardData](https://wiki.libsdl.org/SDL3/SDL_GetClipboardData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3881) |
| [SDL_HasClipboardData](https://wiki.libsdl.org/SDL3/SDL_HasClipboardData) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3895) |
| [SDL_GetClipboardMimeTypes](https://wiki.libsdl.org/SDL3/SDL_GetClipboardMimeTypes) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3906) |
</details>
<details open>
<summary><h3>Dialog</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_ShowOpenFileDialog](https://wiki.libsdl.org/SDL3/SDL_ShowOpenFileDialog) | [:heavy_check_mark:](sdl/functions.go#L893) | [:no_entry_sign:](sdl/sdl_functions_js.go#L11985) |
| [SDL_ShowSaveFileDialog](https://wiki.libsdl.org/SDL3/SDL_ShowSaveFileDialog) | [:heavy_check_mark:](sdl/functions.go#L893) | [:no_entry_sign:](sdl/sdl_functions_js.go#L11989) |
| [SDL_ShowOpenFolderDialog](https://wiki.libsdl.org/SDL3/SDL_ShowOpenFolderDialog) | [:heavy_check_mark:](sdl/functions.go#L893) | [:no_entry_sign:](sdl/sdl_functions_js.go#L11281) |
| [SDL_ShowFileDialogWithProperties](https://wiki.libsdl.org/SDL3/SDL_ShowFileDialogWithProperties) | [:heavy_check_mark:](sdl/functions.go#L893) | [:no_entry_sign:](sdl/sdl_functions_js.go#L11285) |
</details>
<details open>
<summary><h3>Tray</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateTray](https://wiki.libsdl.org/SDL3/SDL_CreateTray) | [:heavy_check_mark:](sdl/functions.go#L893) | [:no_entry_sign:](sdl/sdl_functions_js.go#L11784) |
| [SDL_CreateTrayWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateTrayWithProperties) | [:question:]() | [:question:]() |
| [SDL_SetTrayIcon](https://wiki.libsdl.org/SDL3/SDL_SetTrayIcon) | [:heavy_check_mark:](sdl/methods.go#L427) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11788) |
| [SDL_SetTrayTooltip](https://wiki.libsdl.org/SDL3/SDL_SetTrayTooltip) | [:heavy_check_mark:](sdl/methods.go#L433) | [:no_entry_sign:](sdl/sdl_functions_js.go#L11798) |
| [SDL_CreateTrayMenu](https://wiki.libsdl.org/SDL3/SDL_CreateTrayMenu) | [:heavy_check_mark:](sdl/methods.go#L445) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11802) |
| [SDL_CreateTraySubmenu](https://wiki.libsdl.org/SDL3/SDL_CreateTraySubmenu) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11811) |
| [SDL_GetTrayMenu](https://wiki.libsdl.org/SDL3/SDL_GetTrayMenu) | [:heavy_check_mark:](sdl/methods.go#L439) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11820) |
| [SDL_GetTraySubmenu](https://wiki.libsdl.org/SDL3/SDL_GetTraySubmenu) | [:heavy_check_mark:](sdl/methods.go#L489) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11829) |
| [SDL_GetTrayEntries](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntries) | [:heavy_check_mark:](sdl/methods.go#L459) | [:no_entry_sign:](sdl/sdl_functions_js.go#L11838) |
| [SDL_RemoveTrayEntry](https://wiki.libsdl.org/SDL3/SDL_RemoveTrayEntry) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11842) |
| [SDL_InsertTrayEntryAt](https://wiki.libsdl.org/SDL3/SDL_InsertTrayEntryAt) | [:heavy_check_mark:](sdl/methods.go#L469) | [:no_entry_sign:](sdl/sdl_functions_js.go#L11850) |
| [SDL_SetTrayEntryLabel](https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryLabel) | [:heavy_check_mark:](sdl/methods.go#L495) | [:no_entry_sign:](sdl/sdl_functions_js.go#L11854) |
| [SDL_GetTrayEntryLabel](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryLabel) | [:heavy_check_mark:](sdl/methods.go#L501) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11858) |
| [SDL_SetTrayEntryChecked](https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryChecked) | [:heavy_check_mark:](sdl/methods.go#L507) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11867) |
| [SDL_GetTrayEntryChecked](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryChecked) | [:heavy_check_mark:](sdl/methods.go#L513) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11877) |
| [SDL_SetTrayEntryEnabled](https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryEnabled) | [:heavy_check_mark:](sdl/methods.go#L519) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11886) |
| [SDL_GetTrayEntryEnabled](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryEnabled) | [:heavy_check_mark:](sdl/methods.go#L525) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11896) |
| [SDL_SetTrayEntryCallback](https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryCallback) | [:heavy_check_mark:](sdl/methods.go#L531) | [:no_entry_sign:](sdl/sdl_functions_js.go#L11905) |
| [SDL_ClickTrayEntry](https://wiki.libsdl.org/SDL3/SDL_ClickTrayEntry) | [:heavy_check_mark:](sdl/methods.go#L537) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11909) |
| [SDL_DestroyTray](https://wiki.libsdl.org/SDL3/SDL_DestroyTray) | [:heavy_check_mark:](sdl/methods.go#L451) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11917) |
| [SDL_GetTrayEntryParent](https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryParent) | [:heavy_check_mark:](sdl/methods.go#L543) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11926) |
| [SDL_GetTrayMenuParentEntry](https://wiki.libsdl.org/SDL3/SDL_GetTrayMenuParentEntry) | [:heavy_check_mark:](sdl/methods.go#L475) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11935) |
| [SDL_GetTrayMenuParentTray](https://wiki.libsdl.org/SDL3/SDL_GetTrayMenuParentTray) | [:heavy_check_mark:](sdl/methods.go#L481) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11944) |
| [SDL_UpdateTrays](https://wiki.libsdl.org/SDL3/SDL_UpdateTrays) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11953) |
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
| [SDL_GPUSupportsShaderFormats](https://wiki.libsdl.org/SDL3/SDL_GPUSupportsShaderFormats) | [:heavy_check_mark:](sdl/functions.go#L564) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7515) |
| [SDL_GPUSupportsProperties](https://wiki.libsdl.org/SDL3/SDL_GPUSupportsProperties) | [:heavy_check_mark:](sdl/methods.go#L5926) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7519) |
| [SDL_CreateGPUDevice](https://wiki.libsdl.org/SDL3/SDL_CreateGPUDevice) | [:heavy_check_mark:](sdl/functions.go#L574) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7528) |
| [SDL_CreateGPUDeviceWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateGPUDeviceWithProperties) | [:heavy_check_mark:](sdl/functions.go#L589) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7532) |
| [SDL_DestroyGPUDevice](https://wiki.libsdl.org/SDL3/SDL_DestroyGPUDevice) | [:heavy_check_mark:](sdl/methods.go#L1985) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7541) |
| [SDL_GetNumGPUDrivers](https://wiki.libsdl.org/SDL3/SDL_GetNumGPUDrivers) | [:heavy_check_mark:](sdl/functions.go#L600) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7550) |
| [SDL_GetGPUDriver](https://wiki.libsdl.org/SDL3/SDL_GetGPUDriver) | [:heavy_check_mark:](sdl/functions.go#L606) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7557) |
| [SDL_GetGPUDeviceDriver](https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceDriver) | [:heavy_check_mark:](sdl/methods.go#L1991) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7566) |
| [SDL_GetGPUShaderFormats](https://wiki.libsdl.org/SDL3/SDL_GetGPUShaderFormats) | [:heavy_check_mark:](sdl/methods.go#L2003) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7575) |
| [SDL_GetGPUDeviceProperties](https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceProperties) | [:heavy_check_mark:](sdl/methods.go#L1997) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L11480) |
| [SDL_CreateGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_CreateGPUComputePipeline) | [:heavy_check_mark:](sdl/methods.go#L2020) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7584) |
| [SDL_CreateGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_CreateGPUGraphicsPipeline) | [:heavy_check_mark:](sdl/methods.go#L2032) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7588) |
| [SDL_CreateGPUSampler](https://wiki.libsdl.org/SDL3/SDL_CreateGPUSampler) | [:heavy_check_mark:](sdl/methods.go#L2043) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7592) |
| [SDL_CreateGPUShader](https://wiki.libsdl.org/SDL3/SDL_CreateGPUShader) | [:heavy_check_mark:](sdl/methods.go#L2054) | [:no_entry_sign:](sdl/sdl_functions_js.go#L11966) |
| [SDL_CreateGPUTexture](https://wiki.libsdl.org/SDL3/SDL_CreateGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2066) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7596) |
| [SDL_CreateGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_CreateGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2077) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7600) |
| [SDL_CreateGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_CreateGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2088) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7604) |
| [SDL_SetGPUBufferName](https://wiki.libsdl.org/SDL3/SDL_SetGPUBufferName) | [:heavy_check_mark:](sdl/methods.go#L2099) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7608) |
| [SDL_SetGPUTextureName](https://wiki.libsdl.org/SDL3/SDL_SetGPUTextureName) | [:heavy_check_mark:](sdl/methods.go#L2105) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7622) |
| [SDL_InsertGPUDebugLabel](https://wiki.libsdl.org/SDL3/SDL_InsertGPUDebugLabel) | [:heavy_check_mark:](sdl/methods.go#L6027) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7636) |
| [SDL_PushGPUDebugGroup](https://wiki.libsdl.org/SDL3/SDL_PushGPUDebugGroup) | [:heavy_check_mark:](sdl/methods.go#L6033) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7648) |
| [SDL_PopGPUDebugGroup](https://wiki.libsdl.org/SDL3/SDL_PopGPUDebugGroup) | [:heavy_check_mark:](sdl/methods.go#L6039) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7660) |
| [SDL_ReleaseGPUTexture](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2111) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7668) |
| [SDL_ReleaseGPUSampler](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUSampler) | [:heavy_check_mark:](sdl/methods.go#L2117) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7678) |
| [SDL_ReleaseGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2123) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7688) |
| [SDL_ReleaseGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2129) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7698) |
| [SDL_ReleaseGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUComputePipeline) | [:heavy_check_mark:](sdl/methods.go#L2135) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7708) |
| [SDL_ReleaseGPUShader](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUShader) | [:heavy_check_mark:](sdl/methods.go#L2141) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7718) |
| [SDL_ReleaseGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUGraphicsPipeline) | [:heavy_check_mark:](sdl/methods.go#L2147) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7728) |
| [SDL_AcquireGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_AcquireGPUCommandBuffer) | [:heavy_check_mark:](sdl/methods.go#L2153) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7738) |
| [SDL_PushGPUVertexUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUVertexUniformData) | [:heavy_check_mark:](sdl/methods.go#L6045) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7747) |
| [SDL_PushGPUFragmentUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUFragmentUniformData) | [:heavy_check_mark:](sdl/methods.go#L6051) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7751) |
| [SDL_PushGPUComputeUniformData](https://wiki.libsdl.org/SDL3/SDL_PushGPUComputeUniformData) | [:heavy_check_mark:](sdl/methods.go#L6057) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7755) |
| [SDL_BeginGPURenderPass](https://wiki.libsdl.org/SDL3/SDL_BeginGPURenderPass) | [:heavy_check_mark:](sdl/methods.go#L6063) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7759) |
| [SDL_BindGPUGraphicsPipeline](https://wiki.libsdl.org/SDL3/SDL_BindGPUGraphicsPipeline) | [:heavy_check_mark:](sdl/methods.go#L1300) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7763) |
| [SDL_SetGPUViewport](https://wiki.libsdl.org/SDL3/SDL_SetGPUViewport) | [:heavy_check_mark:](sdl/methods.go#L1306) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7773) |
| [SDL_SetGPUScissor](https://wiki.libsdl.org/SDL3/SDL_SetGPUScissor) | [:heavy_check_mark:](sdl/methods.go#L1312) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7777) |
| [SDL_SetGPUBlendConstants](https://wiki.libsdl.org/SDL3/SDL_SetGPUBlendConstants) | [:question:]() | [:question:]() |
| [SDL_SetGPUStencilReference](https://wiki.libsdl.org/SDL3/SDL_SetGPUStencilReference) | [:heavy_check_mark:](sdl/methods.go#L1318) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7781) |
| [SDL_BindGPUVertexBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexBuffers) | [:heavy_check_mark:](sdl/methods.go#L1324) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7791) |
| [SDL_BindGPUIndexBuffer](https://wiki.libsdl.org/SDL3/SDL_BindGPUIndexBuffer) | [:heavy_check_mark:](sdl/methods.go#L1331) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7795) |
| [SDL_BindGPUVertexSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexSamplers) | [:heavy_check_mark:](sdl/methods.go#L1337) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7799) |
| [SDL_BindGPUVertexStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageTextures) | [:heavy_check_mark:](sdl/methods.go#L1344) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7803) |
| [SDL_BindGPUVertexStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageBuffers) | [:heavy_check_mark:](sdl/methods.go#L1351) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7807) |
| [SDL_BindGPUFragmentSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentSamplers) | [:heavy_check_mark:](sdl/methods.go#L1358) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7811) |
| [SDL_BindGPUFragmentStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageTextures) | [:heavy_check_mark:](sdl/methods.go#L1365) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7815) |
| [SDL_BindGPUFragmentStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageBuffers) | [:heavy_check_mark:](sdl/methods.go#L1372) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7819) |
| [SDL_DrawGPUIndexedPrimitives](https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitives) | [:heavy_check_mark:](sdl/methods.go#L1379) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7823) |
| [SDL_DrawGPUPrimitives](https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitives) | [:heavy_check_mark:](sdl/methods.go#L1385) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7841) |
| [SDL_DrawGPUPrimitivesIndirect](https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitivesIndirect) | [:heavy_check_mark:](sdl/methods.go#L1391) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7857) |
| [SDL_DrawGPUIndexedPrimitivesIndirect](https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitivesIndirect) | [:heavy_check_mark:](sdl/methods.go#L1397) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7871) |
| [SDL_EndGPURenderPass](https://wiki.libsdl.org/SDL3/SDL_EndGPURenderPass) | [:heavy_check_mark:](sdl/methods.go#L1403) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7885) |
| [SDL_BeginGPUComputePass](https://wiki.libsdl.org/SDL3/SDL_BeginGPUComputePass) | [:heavy_check_mark:](sdl/methods.go#L6069) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7893) |
| [SDL_BindGPUComputePipeline](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputePipeline) | [:heavy_check_mark:](sdl/methods.go#L946) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7897) |
| [SDL_BindGPUComputeSamplers](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeSamplers) | [:heavy_check_mark:](sdl/methods.go#L952) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7907) |
| [SDL_BindGPUComputeStorageTextures](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageTextures) | [:heavy_check_mark:](sdl/methods.go#L959) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7911) |
| [SDL_BindGPUComputeStorageBuffers](https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageBuffers) | [:heavy_check_mark:](sdl/methods.go#L966) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7915) |
| [SDL_DispatchGPUCompute](https://wiki.libsdl.org/SDL3/SDL_DispatchGPUCompute) | [:heavy_check_mark:](sdl/methods.go#L973) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7919) |
| [SDL_DispatchGPUComputeIndirect](https://wiki.libsdl.org/SDL3/SDL_DispatchGPUComputeIndirect) | [:heavy_check_mark:](sdl/methods.go#L979) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7933) |
| [SDL_EndGPUComputePass](https://wiki.libsdl.org/SDL3/SDL_EndGPUComputePass) | [:heavy_check_mark:](sdl/methods.go#L985) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7945) |
| [SDL_MapGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_MapGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2164) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7953) |
| [SDL_UnmapGPUTransferBuffer](https://wiki.libsdl.org/SDL3/SDL_UnmapGPUTransferBuffer) | [:heavy_check_mark:](sdl/methods.go#L2175) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7957) |
| [SDL_BeginGPUCopyPass](https://wiki.libsdl.org/SDL3/SDL_BeginGPUCopyPass) | [:heavy_check_mark:](sdl/methods.go#L6075) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L7967) |
| [SDL_UploadToGPUTexture](https://wiki.libsdl.org/SDL3/SDL_UploadToGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2818) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7976) |
| [SDL_UploadToGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_UploadToGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2824) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7980) |
| [SDL_CopyGPUTextureToTexture](https://wiki.libsdl.org/SDL3/SDL_CopyGPUTextureToTexture) | [:heavy_check_mark:](sdl/methods.go#L2830) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7984) |
| [SDL_CopyGPUBufferToBuffer](https://wiki.libsdl.org/SDL3/SDL_CopyGPUBufferToBuffer) | [:heavy_check_mark:](sdl/methods.go#L2836) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7988) |
| [SDL_DownloadFromGPUTexture](https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L2842) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7992) |
| [SDL_DownloadFromGPUBuffer](https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUBuffer) | [:heavy_check_mark:](sdl/methods.go#L2848) | [:no_entry_sign:](sdl/sdl_functions_js.go#L7996) |
| [SDL_EndGPUCopyPass](https://wiki.libsdl.org/SDL3/SDL_EndGPUCopyPass) | [:heavy_check_mark:](sdl/methods.go#L2854) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8000) |
| [SDL_GenerateMipmapsForGPUTexture](https://wiki.libsdl.org/SDL3/SDL_GenerateMipmapsForGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L6081) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8008) |
| [SDL_BlitGPUTexture](https://wiki.libsdl.org/SDL3/SDL_BlitGPUTexture) | [:heavy_check_mark:](sdl/methods.go#L6087) | [:no_entry_sign:](sdl/sdl_functions_js.go#L8018) |
| [SDL_WindowSupportsGPUSwapchainComposition](https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUSwapchainComposition) | [:heavy_check_mark:](sdl/methods.go#L2181) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8022) |
| [SDL_WindowSupportsGPUPresentMode](https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUPresentMode) | [:heavy_check_mark:](sdl/methods.go#L2187) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8035) |
| [SDL_ClaimWindowForGPUDevice](https://wiki.libsdl.org/SDL3/SDL_ClaimWindowForGPUDevice) | [:heavy_check_mark:](sdl/methods.go#L2193) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8048) |
| [SDL_ReleaseWindowFromGPUDevice](https://wiki.libsdl.org/SDL3/SDL_ReleaseWindowFromGPUDevice) | [:heavy_check_mark:](sdl/methods.go#L2203) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8059) |
| [SDL_SetGPUSwapchainParameters](https://wiki.libsdl.org/SDL3/SDL_SetGPUSwapchainParameters) | [:heavy_check_mark:](sdl/methods.go#L2209) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8069) |
| [SDL_SetGPUAllowedFramesInFlight](https://wiki.libsdl.org/SDL3/SDL_SetGPUAllowedFramesInFlight) | [:heavy_check_mark:](sdl/methods.go#L2219) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8084) |
| [SDL_GetGPUSwapchainTextureFormat](https://wiki.libsdl.org/SDL3/SDL_GetGPUSwapchainTextureFormat) | [:heavy_check_mark:](sdl/methods.go#L2229) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8095) |
| [SDL_AcquireGPUSwapchainTexture](https://wiki.libsdl.org/SDL3/SDL_AcquireGPUSwapchainTexture) | [:heavy_check_mark:](sdl/methods.go#L6093) | [:no_entry_sign:](sdl/sdl_functions_js.go#L8106) |
| [SDL_WaitForGPUSwapchain](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUSwapchain) | [:heavy_check_mark:](sdl/methods.go#L2235) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8110) |
| [SDL_WaitAndAcquireGPUSwapchainTexture](https://wiki.libsdl.org/SDL3/SDL_WaitAndAcquireGPUSwapchainTexture) | [:heavy_check_mark:](sdl/methods.go#L6105) | [:no_entry_sign:](sdl/sdl_functions_js.go#L8121) |
| [SDL_SubmitGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBuffer) | [:heavy_check_mark:](sdl/methods.go#L6117) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8125) |
| [SDL_SubmitGPUCommandBufferAndAcquireFence](https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBufferAndAcquireFence) | [:heavy_check_mark:](sdl/methods.go#L6127) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8134) |
| [SDL_CancelGPUCommandBuffer](https://wiki.libsdl.org/SDL3/SDL_CancelGPUCommandBuffer) | [:heavy_check_mark:](sdl/methods.go#L6138) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8143) |
| [SDL_WaitForGPUIdle](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUIdle) | [:heavy_check_mark:](sdl/methods.go#L2245) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8152) |
| [SDL_WaitForGPUFences](https://wiki.libsdl.org/SDL3/SDL_WaitForGPUFences) | [:heavy_check_mark:](sdl/methods.go#L2255) | [:no_entry_sign:](sdl/sdl_functions_js.go#L8161) |
| [SDL_QueryGPUFence](https://wiki.libsdl.org/SDL3/SDL_QueryGPUFence) | [:heavy_check_mark:](sdl/methods.go#L2265) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8165) |
| [SDL_ReleaseGPUFence](https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUFence) | [:heavy_check_mark:](sdl/methods.go#L2271) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8176) |
| [SDL_GPUTextureFormatTexelBlockSize](https://wiki.libsdl.org/SDL3/SDL_GPUTextureFormatTexelBlockSize) | [:heavy_check_mark:](sdl/methods.go#L559) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8186) |
| [SDL_GPUTextureSupportsFormat](https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsFormat) | [:heavy_check_mark:](sdl/methods.go#L2277) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8195) |
| [SDL_GPUTextureSupportsSampleCount](https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsSampleCount) | [:heavy_check_mark:](sdl/methods.go#L2283) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8210) |
| [SDL_CalculateGPUTextureFormatSize](https://wiki.libsdl.org/SDL3/SDL_CalculateGPUTextureFormatSize) | [:heavy_check_mark:](sdl/methods.go#L565) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8223) |
| [SDL_GetPixelFormatFromGPUTextureFormat](https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatFromGPUTextureFormat) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11489) |
| [SDL_GetGPUTextureFormatFromPixelFormat](https://wiki.libsdl.org/SDL3/SDL_GetGPUTextureFormatFromPixelFormat) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11498) |
| [SDL_GDKSuspendGPU](https://wiki.libsdl.org/SDL3/SDL_GDKSuspendGPU) | [:question:]() | [:question:]() |
| [SDL_GDKResumeGPU](https://wiki.libsdl.org/SDL3/SDL_GDKResumeGPU) | [:question:]() | [:question:]() |
</details>
<details>
<summary><h3>Vulkan</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_Vulkan_LoadLibrary](https://wiki.libsdl.org/SDL3/SDL_Vulkan_LoadLibrary) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10881) |
| [SDL_Vulkan_GetVkGetInstanceProcAddr](https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetVkGetInstanceProcAddr) | [:heavy_check_mark:](sdl/functions.go#L893) | [:no_entry_sign:](sdl/sdl_functions_js.go#L10892) |
| [SDL_Vulkan_UnloadLibrary](https://wiki.libsdl.org/SDL3/SDL_Vulkan_UnloadLibrary) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L10896) |
| [SDL_Vulkan_GetInstanceExtensions](https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetInstanceExtensions) | [:heavy_check_mark:](sdl/functions.go#L893) | [:no_entry_sign:](sdl/sdl_functions_js.go#L10902) |
| [SDL_Vulkan_CreateSurface](https://wiki.libsdl.org/SDL3/SDL_Vulkan_CreateSurface) | [:heavy_check_mark:](sdl/functions.go#L893) | [:no_entry_sign:](sdl/sdl_functions_js.go#L10906) |
| [SDL_Vulkan_DestroySurface](https://wiki.libsdl.org/SDL3/SDL_Vulkan_DestroySurface) | [:heavy_check_mark:](sdl/functions.go#L893) | [:no_entry_sign:](sdl/sdl_functions_js.go#L10910) |
| [SDL_Vulkan_GetPresentationSupport](https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetPresentationSupport) | [:heavy_check_mark:](sdl/functions.go#L893) | [:no_entry_sign:](sdl/sdl_functions_js.go#L10914) |
</details>
<details>
<summary><h3>Metal</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_Metal_CreateView](https://wiki.libsdl.org/SDL3/SDL_Metal_CreateView) | [:x:](sdl/methods.go#L4891) | [:x:](sdl/sdl_functions_js.go#L8921) |
| [SDL_Metal_DestroyView](https://wiki.libsdl.org/SDL3/SDL_Metal_DestroyView) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8925) |
| [SDL_Metal_GetLayer](https://wiki.libsdl.org/SDL3/SDL_Metal_GetLayer) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8929) |
</details>
<details open>
<summary><h3>Power</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPowerInfo](https://wiki.libsdl.org/SDL3/SDL_GetPowerInfo) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5169) |
</details>
<details open>
<summary><h3>Sensor</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetSensors](https://wiki.libsdl.org/SDL3/SDL_GetSensors) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5188) |
| [SDL_GetSensorNameForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorNameForID) | [:heavy_check_mark:](sdl/methods.go#L6229) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5200) |
| [SDL_GetSensorTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorTypeForID) | [:heavy_check_mark:](sdl/methods.go#L6235) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5209) |
| [SDL_GetSensorNonPortableTypeForID](https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableTypeForID) | [:heavy_check_mark:](sdl/methods.go#L6241) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5218) |
| [SDL_OpenSensor](https://wiki.libsdl.org/SDL3/SDL_OpenSensor) | [:heavy_check_mark:](sdl/methods.go#L6247) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5227) |
| [SDL_GetSensorFromID](https://wiki.libsdl.org/SDL3/SDL_GetSensorFromID) | [:heavy_check_mark:](sdl/methods.go#L6258) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5236) |
| [SDL_GetSensorProperties](https://wiki.libsdl.org/SDL3/SDL_GetSensorProperties) | [:heavy_check_mark:](sdl/methods.go#L593) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5245) |
| [SDL_GetSensorName](https://wiki.libsdl.org/SDL3/SDL_GetSensorName) | [:heavy_check_mark:](sdl/methods.go#L604) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5254) |
| [SDL_GetSensorType](https://wiki.libsdl.org/SDL3/SDL_GetSensorType) | [:heavy_check_mark:](sdl/methods.go#L615) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5263) |
| [SDL_GetSensorNonPortableType](https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableType) | [:heavy_check_mark:](sdl/methods.go#L621) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5272) |
| [SDL_GetSensorID](https://wiki.libsdl.org/SDL3/SDL_GetSensorID) | [:heavy_check_mark:](sdl/methods.go#L627) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5281) |
| [SDL_GetSensorData](https://wiki.libsdl.org/SDL3/SDL_GetSensorData) | [:heavy_check_mark:](sdl/methods.go#L638) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5290) |
| [SDL_CloseSensor](https://wiki.libsdl.org/SDL3/SDL_CloseSensor) | [:heavy_check_mark:](sdl/methods.go#L648) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5308) |
| [SDL_UpdateSensors](https://wiki.libsdl.org/SDL3/SDL_UpdateSensors) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L5317) |
</details>
<details>
<summary><h3>Process</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_CreateProcess](https://wiki.libsdl.org/SDL3/SDL_CreateProcess) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11585) |
| [SDL_CreateProcessWithProperties](https://wiki.libsdl.org/SDL3/SDL_CreateProcessWithProperties) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8944) |
| [SDL_GetProcessProperties](https://wiki.libsdl.org/SDL3/SDL_GetProcessProperties) | [:heavy_check_mark:](sdl/methods.go#L6150) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8953) |
| [SDL_ReadProcess](https://wiki.libsdl.org/SDL3/SDL_ReadProcess) | [:heavy_check_mark:](sdl/methods.go#L6161) | [:no_entry_sign:](sdl/sdl_functions_js.go#L8962) |
| [SDL_GetProcessInput](https://wiki.libsdl.org/SDL3/SDL_GetProcessInput) | [:heavy_check_mark:](sdl/methods.go#L6179) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8966) |
| [SDL_GetProcessOutput](https://wiki.libsdl.org/SDL3/SDL_GetProcessOutput) | [:heavy_check_mark:](sdl/methods.go#L6190) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8975) |
| [SDL_KillProcess](https://wiki.libsdl.org/SDL3/SDL_KillProcess) | [:heavy_check_mark:](sdl/methods.go#L6201) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8984) |
| [SDL_WaitProcess](https://wiki.libsdl.org/SDL3/SDL_WaitProcess) | [:heavy_check_mark:](sdl/methods.go#L6211) | [:no_entry_sign:](sdl/sdl_functions_js.go#L8995) |
| [SDL_DestroyProcess](https://wiki.libsdl.org/SDL3/SDL_DestroyProcess) | [:heavy_check_mark:](sdl/methods.go#L6221) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8999) |
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
| [SDL_GetNumLogicalCPUCores](https://wiki.libsdl.org/SDL3/SDL_GetNumLogicalCPUCores) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3918) |
| [SDL_GetCPUCacheLineSize](https://wiki.libsdl.org/SDL3/SDL_GetCPUCacheLineSize) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3925) |
| [SDL_HasAltiVec](https://wiki.libsdl.org/SDL3/SDL_HasAltiVec) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3932) |
| [SDL_HasMMX](https://wiki.libsdl.org/SDL3/SDL_HasMMX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3939) |
| [SDL_HasSSE](https://wiki.libsdl.org/SDL3/SDL_HasSSE) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3946) |
| [SDL_HasSSE2](https://wiki.libsdl.org/SDL3/SDL_HasSSE2) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3953) |
| [SDL_HasSSE3](https://wiki.libsdl.org/SDL3/SDL_HasSSE3) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3960) |
| [SDL_HasSSE41](https://wiki.libsdl.org/SDL3/SDL_HasSSE41) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3967) |
| [SDL_HasSSE42](https://wiki.libsdl.org/SDL3/SDL_HasSSE42) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3974) |
| [SDL_HasAVX](https://wiki.libsdl.org/SDL3/SDL_HasAVX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3981) |
| [SDL_HasAVX2](https://wiki.libsdl.org/SDL3/SDL_HasAVX2) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3988) |
| [SDL_HasAVX512F](https://wiki.libsdl.org/SDL3/SDL_HasAVX512F) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L3995) |
| [SDL_HasARMSIMD](https://wiki.libsdl.org/SDL3/SDL_HasARMSIMD) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4002) |
| [SDL_HasNEON](https://wiki.libsdl.org/SDL3/SDL_HasNEON) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4009) |
| [SDL_HasSVE2](https://wiki.libsdl.org/SDL3/SDL_HasSVE2) | [:question:]() | [:question:]() |
| [SDL_HasLSX](https://wiki.libsdl.org/SDL3/SDL_HasLSX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4016) |
| [SDL_HasLASX](https://wiki.libsdl.org/SDL3/SDL_HasLASX) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4023) |
| [SDL_GetSystemRAM](https://wiki.libsdl.org/SDL3/SDL_GetSystemRAM) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4030) |
| [SDL_GetSIMDAlignment](https://wiki.libsdl.org/SDL3/SDL_GetSIMDAlignment) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L4037) |
| [SDL_GetSystemPageSize](https://wiki.libsdl.org/SDL3/SDL_GetSystemPageSize) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11150) |
</details>
<details>
<summary><h3>Locale</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GetPreferredLocales](https://wiki.libsdl.org/SDL3/SDL_GetPreferredLocales) | [:heavy_check_mark:](sdl/functions.go#L893) | [:heavy_check_mark:](sdl/sdl_functions_js.go#L8693) |
</details>
<details>
<summary><h3>System</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_SetWindowsMessageHook](https://wiki.libsdl.org/SDL3/SDL_SetWindowsMessageHook) | [:question:]() | [:question:]() |
| [SDL_GetDirect3D9AdapterIndex](https://wiki.libsdl.org/SDL3/SDL_GetDirect3D9AdapterIndex) | [:question:]() | [:question:]() |
| [SDL_GetDXGIOutputInfo](https://wiki.libsdl.org/SDL3/SDL_GetDXGIOutputInfo) | [:question:]() | [:question:]() |
| [SDL_SetX11EventHook](https://wiki.libsdl.org/SDL3/SDL_SetX11EventHook) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11745) |
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
| [SDL_IsTablet](https://wiki.libsdl.org/SDL3/SDL_IsTablet) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10644) |
| [SDL_IsTV](https://wiki.libsdl.org/SDL3/SDL_IsTV) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10651) |
| [SDL_GetDeviceFormFactor](https://wiki.libsdl.org/SDL3/SDL_GetDeviceFormFactor) | [:question:]() | [:question:]() |
| [SDL_GetDeviceFormFactorName](https://wiki.libsdl.org/SDL3/SDL_GetDeviceFormFactorName) | [:question:]() | [:question:]() |
| [SDL_GetSandbox](https://wiki.libsdl.org/SDL3/SDL_GetSandbox) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10658) |
| [SDL_OnApplicationWillTerminate](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillTerminate) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10665) |
| [SDL_OnApplicationDidReceiveMemoryWarning](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidReceiveMemoryWarning) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10671) |
| [SDL_OnApplicationWillEnterBackground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillEnterBackground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10677) |
| [SDL_OnApplicationDidEnterBackground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidEnterBackground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10683) |
| [SDL_OnApplicationWillEnterForeground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillEnterForeground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10689) |
| [SDL_OnApplicationDidEnterForeground](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidEnterForeground) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L10695) |
| [SDL_OnApplicationDidChangeStatusBarOrientation](https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidChangeStatusBarOrientation) | [:question:]() | [:question:]() |
| [SDL_GetGDKTaskQueue](https://wiki.libsdl.org/SDL3/SDL_GetGDKTaskQueue) | [:question:]() | [:question:]() |
| [SDL_GetGDKDefaultUser](https://wiki.libsdl.org/SDL3/SDL_GetGDKDefaultUser) | [:question:]() | [:question:]() |
| [SDL_IsUbuntuTouch](https://wiki.libsdl.org/SDL3/SDL_IsUbuntuTouch) | [:question:]() | [:question:]() |
</details>
<details>
<summary><h3>Misc</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_OpenURL](https://wiki.libsdl.org/SDL3/SDL_OpenURL) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L8933) |
</details>
<details>
<summary><h3>GUID</h3></summary>

|Function|Desktop|WASM/js|
|:--|:--:|:--:|
| [SDL_GUIDToString](https://wiki.libsdl.org/SDL3/SDL_GUIDToString) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11289) |
| [SDL_StringToGUID](https://wiki.libsdl.org/SDL3/SDL_StringToGUID) | [:question:]() | [:question:](sdl/sdl_functions_js.go#L11304) |
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
| [TTF_GetFreeTypeVersion](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFreeTypeVersion) | [:heavy_check_mark:](ttf/functions.go#L17) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L20) |
| [TTF_GetHarfBuzzVersion](https://wiki.libsdl.org/SDL3_ttf/TTF_GetHarfBuzzVersion) | [:heavy_check_mark:](ttf/functions.go#L27) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L37) |
| [TTF_Init](https://wiki.libsdl.org/SDL3_ttf/TTF_Init) | [:heavy_check_mark:](ttf/functions.go#L37) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L54) |
| [TTF_OpenFont](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFont) | [:heavy_check_mark:](ttf/functions.go#L47) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L62) |
| [TTF_OpenFontIO](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontIO) | [:heavy_check_mark:](ttf/functions.go#L58) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L75) |
| [TTF_OpenFontWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontWithProperties) | [:heavy_check_mark:](ttf/functions.go#L69) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1328) |
| [TTF_CopyFont](https://wiki.libsdl.org/SDL3_ttf/TTF_CopyFont) | [:heavy_check_mark:](ttf/methods.go#L409) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L93) |
| [TTF_GetFontProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontProperties) | [:heavy_check_mark:](ttf/methods.go#L420) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L102) |
| [TTF_GetFontGeneration](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontGeneration) | [:heavy_check_mark:](ttf/methods.go#L426) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L111) |
| [TTF_AddFallbackFont](https://wiki.libsdl.org/SDL3_ttf/TTF_AddFallbackFont) | [:heavy_check_mark:](ttf/methods.go#L432) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L120) |
| [TTF_RemoveFallbackFont](https://wiki.libsdl.org/SDL3_ttf/TTF_RemoveFallbackFont) | [:heavy_check_mark:](ttf/methods.go#L442) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L131) |
| [TTF_ClearFallbackFonts](https://wiki.libsdl.org/SDL3_ttf/TTF_ClearFallbackFonts) | [:heavy_check_mark:](ttf/methods.go#L448) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L141) |
| [TTF_SetFontSize](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSize) | [:heavy_check_mark:](ttf/methods.go#L454) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L149) |
| [TTF_SetFontSizeDPI](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSizeDPI) | [:heavy_check_mark:](ttf/methods.go#L464) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L160) |
| [TTF_GetFontSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSize) | [:heavy_check_mark:](ttf/methods.go#L474) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L175) |
| [TTF_GetFontDPI](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDPI) | [:heavy_check_mark:](ttf/methods.go#L485) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L184) |
| [TTF_SetFontStyle](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontStyle) | [:heavy_check_mark:](ttf/methods.go#L497) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L201) |
| [TTF_GetFontStyle](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyle) | [:heavy_check_mark:](ttf/methods.go#L503) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L211) |
| [TTF_SetFontOutline](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontOutline) | [:heavy_check_mark:](ttf/methods.go#L509) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L220) |
| [TTF_GetFontOutline](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontOutline) | [:heavy_check_mark:](ttf/methods.go#L519) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L231) |
| [TTF_SetFontHinting](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontHinting) | [:heavy_check_mark:](ttf/methods.go#L525) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L240) |
| [TTF_GetNumFontFaces](https://wiki.libsdl.org/SDL3_ttf/TTF_GetNumFontFaces) | [:heavy_check_mark:](ttf/methods.go#L531) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L250) |
| [TTF_GetFontHinting](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHinting) | [:heavy_check_mark:](ttf/methods.go#L537) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L259) |
| [TTF_SetFontSDF](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSDF) | [:heavy_check_mark:](ttf/methods.go#L543) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L268) |
| [TTF_GetFontSDF](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSDF) | [:heavy_check_mark:](ttf/methods.go#L553) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L279) |
| [TTF_GetFontWeight](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontWeight) | [:question:]() | [:question:](ttf/ttf_functions_js.go#L1337) |
| [TTF_SetFontWrapAlignment](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontWrapAlignment) | [:heavy_check_mark:](ttf/methods.go#L559) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L288) |
| [TTF_GetFontWrapAlignment](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontWrapAlignment) | [:heavy_check_mark:](ttf/methods.go#L565) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L298) |
| [TTF_GetFontHeight](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHeight) | [:heavy_check_mark:](ttf/methods.go#L571) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L307) |
| [TTF_GetFontAscent](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontAscent) | [:heavy_check_mark:](ttf/methods.go#L577) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L316) |
| [TTF_GetFontDescent](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDescent) | [:heavy_check_mark:](ttf/methods.go#L583) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L325) |
| [TTF_SetFontLineSkip](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLineSkip) | [:heavy_check_mark:](ttf/methods.go#L589) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L334) |
| [TTF_GetFontLineSkip](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontLineSkip) | [:heavy_check_mark:](ttf/methods.go#L595) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L344) |
| [TTF_SetFontKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontKerning) | [:heavy_check_mark:](ttf/methods.go#L601) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L353) |
| [TTF_GetFontKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontKerning) | [:heavy_check_mark:](ttf/methods.go#L607) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L363) |
| [TTF_FontIsFixedWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsFixedWidth) | [:heavy_check_mark:](ttf/methods.go#L613) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L372) |
| [TTF_FontIsScalable](https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsScalable) | [:heavy_check_mark:](ttf/methods.go#L619) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L381) |
| [TTF_GetFontFamilyName](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontFamilyName) | [:heavy_check_mark:](ttf/methods.go#L625) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L390) |
| [TTF_GetFontStyleName](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyleName) | [:heavy_check_mark:](ttf/methods.go#L631) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L399) |
| [TTF_SetFontDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontDirection) | [:heavy_check_mark:](ttf/methods.go#L637) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L408) |
| [TTF_GetFontDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDirection) | [:heavy_check_mark:](ttf/methods.go#L647) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L419) |
| [TTF_SetFontCharSpacing](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontCharSpacing) | [:question:]() | [:question:]() |
| [TTF_GetFontCharSpacing](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontCharSpacing) | [:question:]() | [:question:]() |
| [TTF_StringToTag](https://wiki.libsdl.org/SDL3_ttf/TTF_StringToTag) | [:heavy_check_mark:](ttf/functions.go#L80) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L428) |
| [TTF_TagToString](https://wiki.libsdl.org/SDL3_ttf/TTF_TagToString) | [:heavy_check_mark:](ttf/functions.go#L86) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1227) |
| [TTF_SetFontScript](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontScript) | [:heavy_check_mark:](ttf/methods.go#L653) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L439) |
| [TTF_GetFontScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontScript) | [:heavy_check_mark:](ttf/methods.go#L663) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L450) |
| [TTF_GetGlyphScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphScript) | [:heavy_check_mark:](ttf/functions.go#L100) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L459) |
| [TTF_SetFontLanguage](https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLanguage) | [:heavy_check_mark:](ttf/methods.go#L669) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L468) |
| [TTF_FontHasGlyph](https://wiki.libsdl.org/SDL3_ttf/TTF_FontHasGlyph) | [:heavy_check_mark:](ttf/methods.go#L679) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L481) |
| [TTF_GetGlyphImage](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImage) | [:heavy_check_mark:](ttf/methods.go#L685) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L492) |
| [TTF_GetGlyphImageForIndex](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImageForIndex) | [:heavy_check_mark:](ttf/methods.go#L698) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L508) |
| [TTF_GetGlyphMetrics](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphMetrics) | [:heavy_check_mark:](ttf/methods.go#L711) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L524) |
| [TTF_GetGlyphKerning](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphKerning) | [:heavy_check_mark:](ttf/methods.go#L723) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L552) |
| [TTF_GetStringSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSize) | [:heavy_check_mark:](ttf/methods.go#L735) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L570) |
| [TTF_GetStringSizeWrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSizeWrapped) | [:heavy_check_mark:](ttf/methods.go#L747) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L595) |
| [TTF_MeasureString](https://wiki.libsdl.org/SDL3_ttf/TTF_MeasureString) | [:heavy_check_mark:](ttf/methods.go#L760) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L618) |
| [TTF_RenderText_Solid](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid) | [:heavy_check_mark:](ttf/methods.go#L773) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1242) |
| [TTF_RenderText_Solid_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid_Wrapped) | [:heavy_check_mark:](ttf/methods.go#L784) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1265) |
| [TTF_RenderGlyph_Solid](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Solid) | [:heavy_check_mark:](ttf/methods.go#L795) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1384) |
| [TTF_RenderText_Shaded](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded) | [:heavy_check_mark:](ttf/methods.go#L806) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1400) |
| [TTF_RenderText_Shaded_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded_Wrapped) | [:heavy_check_mark:](ttf/methods.go#L817) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1420) |
| [TTF_RenderGlyph_Shaded](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Shaded) | [:heavy_check_mark:](ttf/methods.go#L828) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1442) |
| [TTF_RenderText_Blended](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended) | [:heavy_check_mark:](ttf/methods.go#L839) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1285) |
| [TTF_RenderText_Blended_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended_Wrapped) | [:heavy_check_mark:](ttf/methods.go#L850) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1308) |
| [TTF_RenderGlyph_Blended](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Blended) | [:heavy_check_mark:](ttf/methods.go#L861) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1460) |
| [TTF_RenderText_LCD](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD) | [:heavy_check_mark:](ttf/methods.go#L872) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1476) |
| [TTF_RenderText_LCD_Wrapped](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD_Wrapped) | [:heavy_check_mark:](ttf/methods.go#L883) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1496) |
| [TTF_RenderGlyph_LCD](https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_LCD) | [:heavy_check_mark:](ttf/methods.go#L894) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1518) |
| [TTF_CreateSurfaceTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateSurfaceTextEngine) | [:heavy_check_mark:](ttf/functions.go#L111) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L641) |
| [TTF_DrawSurfaceText](https://wiki.libsdl.org/SDL3_ttf/TTF_DrawSurfaceText) | [:heavy_check_mark:](ttf/methods.go#L13) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L648) |
| [TTF_DestroySurfaceTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroySurfaceTextEngine) | [:heavy_check_mark:](ttf/methods.go#L366) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L663) |
| [TTF_CreateRendererTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngine) | [:heavy_check_mark:](ttf/functions.go#L122) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L672) |
| [TTF_CreateRendererTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngineWithProperties) | [:heavy_check_mark:](ttf/functions.go#L133) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1346) |
| [TTF_DrawRendererText](https://wiki.libsdl.org/SDL3_ttf/TTF_DrawRendererText) | [:heavy_check_mark:](ttf/methods.go#L23) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L686) |
| [TTF_DestroyRendererTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyRendererTextEngine) | [:heavy_check_mark:](ttf/methods.go#L372) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L701) |
| [TTF_CreateGPUTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngine) | [:heavy_check_mark:](ttf/functions.go#L144) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L710) |
| [TTF_CreateGPUTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngineWithProperties) | [:heavy_check_mark:](ttf/functions.go#L154) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1355) |
| [TTF_GetGPUTextDrawData](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextDrawData) | [:heavy_check_mark:](ttf/methods.go#L33) | [:no_entry_sign:](ttf/ttf_functions_js.go#L719) |
| [TTF_DestroyGPUTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyGPUTextEngine) | [:heavy_check_mark:](ttf/methods.go#L378) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L723) |
| [TTF_SetGPUTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_SetGPUTextEngineWinding) | [:heavy_check_mark:](ttf/methods.go#L384) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L732) |
| [TTF_GetGPUTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextEngineWinding) | [:heavy_check_mark:](ttf/methods.go#L390) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L742) |
| [TTF_CreateGLTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGLTextEngine) | [:question:]() | [:question:]() |
| [TTF_CreateGLTextEngineWithProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGLTextEngineWithProperties) | [:question:]() | [:question:]() |
| [TTF_GetGLTextDrawData](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGLTextDrawData) | [:question:]() | [:question:]() |
| [TTF_DestroyGLTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyGLTextEngine) | [:question:]() | [:question:]() |
| [TTF_SetGLTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_SetGLTextEngineWinding) | [:question:]() | [:question:]() |
| [TTF_GetGLTextEngineWinding](https://wiki.libsdl.org/SDL3_ttf/TTF_GetGLTextEngineWinding) | [:question:]() | [:question:]() |
| [TTF_CreateText](https://wiki.libsdl.org/SDL3_ttf/TTF_CreateText) | [:heavy_check_mark:](ttf/methods.go#L396) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L751) |
| [TTF_GetTextProperties](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextProperties) | [:heavy_check_mark:](ttf/methods.go#L44) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L777) |
| [TTF_SetTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextEngine) | [:heavy_check_mark:](ttf/methods.go#L50) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L786) |
| [TTF_GetTextEngine](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextEngine) | [:heavy_check_mark:](ttf/methods.go#L60) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L797) |
| [TTF_SetTextFont](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextFont) | [:heavy_check_mark:](ttf/methods.go#L71) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L806) |
| [TTF_GetTextFont](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextFont) | [:heavy_check_mark:](ttf/methods.go#L81) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L817) |
| [TTF_SetTextDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextDirection) | [:heavy_check_mark:](ttf/methods.go#L92) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L826) |
| [TTF_GetTextDirection](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextDirection) | [:heavy_check_mark:](ttf/methods.go#L102) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L837) |
| [TTF_SetTextScript](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextScript) | [:heavy_check_mark:](ttf/methods.go#L108) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L846) |
| [TTF_GetTextScript](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextScript) | [:heavy_check_mark:](ttf/methods.go#L118) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L857) |
| [TTF_SetTextColor](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColor) | [:heavy_check_mark:](ttf/methods.go#L124) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L866) |
| [TTF_SetTextColorFloat](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColorFloat) | [:heavy_check_mark:](ttf/methods.go#L134) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L887) |
| [TTF_GetTextColor](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColor) | [:heavy_check_mark:](ttf/methods.go#L144) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L904) |
| [TTF_GetTextColorFloat](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColorFloat) | [:heavy_check_mark:](ttf/methods.go#L156) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L927) |
| [TTF_SetTextPosition](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextPosition) | [:heavy_check_mark:](ttf/methods.go#L168) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L950) |
| [TTF_GetTextPosition](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextPosition) | [:heavy_check_mark:](ttf/methods.go#L174) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L963) |
| [TTF_SetTextWrapWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWidth) | [:heavy_check_mark:](ttf/methods.go#L184) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L980) |
| [TTF_GetTextWrapWidth](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextWrapWidth) | [:heavy_check_mark:](ttf/methods.go#L194) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L991) |
| [TTF_SetTextWrapWhitespaceVisible](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWhitespaceVisible) | [:heavy_check_mark:](ttf/methods.go#L206) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1005) |
| [TTF_TextWrapWhitespaceVisible](https://wiki.libsdl.org/SDL3_ttf/TTF_TextWrapWhitespaceVisible) | [:heavy_check_mark:](ttf/methods.go#L216) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1016) |
| [TTF_SetTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextString) | [:heavy_check_mark:](ttf/methods.go#L222) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1025) |
| [TTF_InsertTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_InsertTextString) | [:heavy_check_mark:](ttf/methods.go#L232) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1040) |
| [TTF_AppendTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_AppendTextString) | [:heavy_check_mark:](ttf/methods.go#L242) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1057) |
| [TTF_DeleteTextString](https://wiki.libsdl.org/SDL3_ttf/TTF_DeleteTextString) | [:heavy_check_mark:](ttf/methods.go#L252) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1072) |
| [TTF_GetTextSize](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSize) | [:heavy_check_mark:](ttf/methods.go#L262) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1085) |
| [TTF_GetTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubString) | [:heavy_check_mark:](ttf/methods.go#L274) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1102) |
| [TTF_GetTextSubStringForLine](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForLine) | [:heavy_check_mark:](ttf/methods.go#L286) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1118) |
| [TTF_GetTextSubStringsForRange](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringsForRange) | [:heavy_check_mark:](ttf/methods.go#L298) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1364) |
| [TTF_GetTextSubStringForPoint](https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForPoint) | [:heavy_check_mark:](ttf/methods.go#L312) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1134) |
| [TTF_GetPreviousTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetPreviousTextSubString) | [:heavy_check_mark:](ttf/methods.go#L324) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1152) |
| [TTF_GetNextTextSubString](https://wiki.libsdl.org/SDL3_ttf/TTF_GetNextTextSubString) | [:heavy_check_mark:](ttf/methods.go#L336) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1168) |
| [TTF_UpdateText](https://wiki.libsdl.org/SDL3_ttf/TTF_UpdateText) | [:heavy_check_mark:](ttf/methods.go#L348) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1184) |
| [TTF_DestroyText](https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyText) | [:heavy_check_mark:](ttf/methods.go#L358) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1193) |
| [TTF_CloseFont](https://wiki.libsdl.org/SDL3_ttf/TTF_CloseFont) | [:heavy_check_mark:](ttf/methods.go#L905) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1205) |
| [TTF_Quit](https://wiki.libsdl.org/SDL3_ttf/TTF_Quit) | [:heavy_check_mark:](ttf/functions.go#L164) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1214) |
| [TTF_WasInit](https://wiki.libsdl.org/SDL3_ttf/TTF_WasInit) | [:heavy_check_mark:](ttf/functions.go#L170) | [:heavy_check_mark:](ttf/ttf_functions_js.go#L1220) |
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

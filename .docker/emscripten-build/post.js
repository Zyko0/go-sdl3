// Optimized builds drop Module.calledRun, which the Go side polls to know
// the runtime is initialized (see internal/memory_js.go).
{
  const onRuntimeInitialized = Module['onRuntimeInitialized'];
  Module['onRuntimeInitialized'] = () => {
    Module['calledRun'] = true;
    onRuntimeInitialized?.();
  };
}

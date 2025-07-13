// include: shell.js
// The Module object: Our interface to the outside world. We import
// and export values on it. There are various ways Module can be used:
// 1. Not defined. We create it here
// 2. A function parameter, function(moduleArg) => Promise<Module>
// 3. pre-run appended it, var Module = {}; ..generated code..
// 4. External script tag defines var Module.
// We need to check if Module already exists (e.g. case 3 above).
// Substitution will be replaced with actual code on later stage of the build,
// this way Closure Compiler will not mangle it (e.g. case 4. above).
// Note that if you want to run closure, and also to use Module
// after the generated code, you will need to define   var Module = {};
// before the code. Then that object will be used in the code, and you
// can continue to use Module afterwards as well.
var Module = typeof Module != 'undefined' ? Module : {};

// Determine the runtime environment we are in. You can customize this by
// setting the ENVIRONMENT setting at compile time (see settings.js).

// Attempt to auto-detect the environment
var ENVIRONMENT_IS_WEB = typeof window == 'object';
var ENVIRONMENT_IS_WORKER = typeof WorkerGlobalScope != 'undefined';
// N.b. Electron.js environment is simultaneously a NODE-environment, but
// also a web environment.
var ENVIRONMENT_IS_NODE = typeof process == 'object' && process.versions?.node && process.type != 'renderer';
var ENVIRONMENT_IS_SHELL = !ENVIRONMENT_IS_WEB && !ENVIRONMENT_IS_NODE && !ENVIRONMENT_IS_WORKER;

// --pre-jses are emitted after the Module integration code, so that they can
// refer to Module (if they choose; they can also define Module)


var arguments_ = [];
var thisProgram = './this.program';
var quit_ = (status, toThrow) => {
  throw toThrow;
};

// In MODULARIZE mode _scriptName needs to be captured already at the very top of the page immediately when the page is parsed, so it is generated there
// before the page load. In non-MODULARIZE modes generate it here.
var _scriptName = typeof document != 'undefined' ? document.currentScript?.src : undefined;

if (typeof __filename != 'undefined') { // Node
  _scriptName = __filename;
} else
if (ENVIRONMENT_IS_WORKER) {
  _scriptName = self.location.href;
}

// `/` should be present at the end if `scriptDirectory` is not empty
var scriptDirectory = '';
function locateFile(path) {
  if (Module['locateFile']) {
    return Module['locateFile'](path, scriptDirectory);
  }
  return scriptDirectory + path;
}

// Hooks that are implemented differently in different runtime environments.
var readAsync, readBinary;

if (ENVIRONMENT_IS_NODE) {
  const isNode = typeof process == 'object' && process.versions?.node && process.type != 'renderer';
  if (!isNode) throw new Error('not compiled for this environment (did you build to HTML and try to run it not on the web, or set ENVIRONMENT to something - like node - and run it someplace else - like on the web?)');

  var nodeVersion = process.versions.node;
  var numericVersion = nodeVersion.split('.').slice(0, 3);
  numericVersion = (numericVersion[0] * 10000) + (numericVersion[1] * 100) + (numericVersion[2].split('-')[0] * 1);
  if (numericVersion < 160000) {
    throw new Error('This emscripten-generated code requires node v16.0.0 (detected v' + nodeVersion + ')');
  }

  // These modules will usually be used on Node.js. Load them eagerly to avoid
  // the complexity of lazy-loading.
  var fs = require('fs');

  scriptDirectory = __dirname + '/';

// include: node_shell_read.js
readBinary = (filename) => {
  // We need to re-wrap `file://` strings to URLs.
  filename = isFileURI(filename) ? new URL(filename) : filename;
  var ret = fs.readFileSync(filename);
  assert(Buffer.isBuffer(ret));
  return ret;
};

readAsync = async (filename, binary = true) => {
  // See the comment in the `readBinary` function.
  filename = isFileURI(filename) ? new URL(filename) : filename;
  var ret = fs.readFileSync(filename, binary ? undefined : 'utf8');
  assert(binary ? Buffer.isBuffer(ret) : typeof ret == 'string');
  return ret;
};
// end include: node_shell_read.js
  if (process.argv.length > 1) {
    thisProgram = process.argv[1].replace(/\\/g, '/');
  }

  arguments_ = process.argv.slice(2);

  // MODULARIZE will export the module in the proper place outside, we don't need to export here
  if (typeof module != 'undefined') {
    module['exports'] = Module;
  }

  quit_ = (status, toThrow) => {
    process.exitCode = status;
    throw toThrow;
  };

} else
if (ENVIRONMENT_IS_SHELL) {

  const isNode = typeof process == 'object' && process.versions?.node && process.type != 'renderer';
  if (isNode || typeof window == 'object' || typeof WorkerGlobalScope != 'undefined') throw new Error('not compiled for this environment (did you build to HTML and try to run it not on the web, or set ENVIRONMENT to something - like node - and run it someplace else - like on the web?)');

} else

// Note that this includes Node.js workers when relevant (pthreads is enabled).
// Node.js workers are detected as a combination of ENVIRONMENT_IS_WORKER and
// ENVIRONMENT_IS_NODE.
if (ENVIRONMENT_IS_WEB || ENVIRONMENT_IS_WORKER) {
  try {
    scriptDirectory = new URL('.', _scriptName).href; // includes trailing slash
  } catch {
    // Must be a `blob:` or `data:` URL (e.g. `blob:http://site.com/etc/etc`), we cannot
    // infer anything from them.
  }

  if (!(typeof window == 'object' || typeof WorkerGlobalScope != 'undefined')) throw new Error('not compiled for this environment (did you build to HTML and try to run it not on the web, or set ENVIRONMENT to something - like node - and run it someplace else - like on the web?)');

  {
// include: web_or_worker_shell_read.js
if (ENVIRONMENT_IS_WORKER) {
    readBinary = (url) => {
      var xhr = new XMLHttpRequest();
      xhr.open('GET', url, false);
      xhr.responseType = 'arraybuffer';
      xhr.send(null);
      return new Uint8Array(/** @type{!ArrayBuffer} */(xhr.response));
    };
  }

  readAsync = async (url) => {
    // Fetch has some additional restrictions over XHR, like it can't be used on a file:// url.
    // See https://github.com/github/fetch/pull/92#issuecomment-140665932
    // Cordova or Electron apps are typically loaded from a file:// url.
    // So use XHR on webview if URL is a file URL.
    if (isFileURI(url)) {
      return new Promise((resolve, reject) => {
        var xhr = new XMLHttpRequest();
        xhr.open('GET', url, true);
        xhr.responseType = 'arraybuffer';
        xhr.onload = () => {
          if (xhr.status == 200 || (xhr.status == 0 && xhr.response)) { // file URLs can return 0
            resolve(xhr.response);
            return;
          }
          reject(xhr.status);
        };
        xhr.onerror = reject;
        xhr.send(null);
      });
    }
    var response = await fetch(url, { credentials: 'same-origin' });
    if (response.ok) {
      return response.arrayBuffer();
    }
    throw new Error(response.status + ' : ' + response.url);
  };
// end include: web_or_worker_shell_read.js
  }
} else
{
  throw new Error('environment detection error');
}

var out = console.log.bind(console);
var err = console.error.bind(console);

var IDBFS = 'IDBFS is no longer included by default; build with -lidbfs.js';
var PROXYFS = 'PROXYFS is no longer included by default; build with -lproxyfs.js';
var WORKERFS = 'WORKERFS is no longer included by default; build with -lworkerfs.js';
var FETCHFS = 'FETCHFS is no longer included by default; build with -lfetchfs.js';
var ICASEFS = 'ICASEFS is no longer included by default; build with -licasefs.js';
var JSFILEFS = 'JSFILEFS is no longer included by default; build with -ljsfilefs.js';
var OPFS = 'OPFS is no longer included by default; build with -lopfs.js';

var NODEFS = 'NODEFS is no longer included by default; build with -lnodefs.js';

// perform assertions in shell.js after we set up out() and err(), as otherwise
// if an assertion fails it cannot print the message

assert(!ENVIRONMENT_IS_SHELL, 'shell environment detected but not enabled at build time.  Add `shell` to `-sENVIRONMENT` to enable.');

// end include: shell.js

// include: preamble.js
// === Preamble library stuff ===

// Documentation for the public APIs defined in this file must be updated in:
//    site/source/docs/api_reference/preamble.js.rst
// A prebuilt local version of the documentation is available at:
//    site/build/text/docs/api_reference/preamble.js.txt
// You can also build docs locally as HTML or other formats in site/
// An online HTML version (which may be of a different version of Emscripten)
//    is up at http://kripken.github.io/emscripten-site/docs/api_reference/preamble.js.html

var wasmBinary;

if (typeof WebAssembly != 'object') {
  err('no native wasm support detected');
}

// Wasm globals

//========================================
// Runtime essentials
//========================================

// whether we are quitting the application. no code should run after this.
// set in exit() and abort()
var ABORT = false;

// set by exit() and abort().  Passed to 'onExit' handler.
// NOTE: This is also used as the process return code code in shell environments
// but only when noExitRuntime is false.
var EXITSTATUS;

// In STRICT mode, we only define assert() when ASSERTIONS is set.  i.e. we
// don't define it at all in release modes.  This matches the behaviour of
// MINIMAL_RUNTIME.
// TODO(sbc): Make this the default even without STRICT enabled.
/** @type {function(*, string=)} */
function assert(condition, text) {
  if (!condition) {
    abort('Assertion failed' + (text ? ': ' + text : ''));
  }
}

// We used to include malloc/free by default in the past. Show a helpful error in
// builds with assertions.

/**
 * Indicates whether filename is delivered via file protocol (as opposed to http/https)
 * @noinline
 */
var isFileURI = (filename) => filename.startsWith('file://');

// include: runtime_common.js
// include: runtime_stack_check.js
// Initializes the stack cookie. Called at the startup of main and at the startup of each thread in pthreads mode.
function writeStackCookie() {
  var max = _emscripten_stack_get_end();
  assert((max & 3) == 0);
  // If the stack ends at address zero we write our cookies 4 bytes into the
  // stack.  This prevents interference with SAFE_HEAP and ASAN which also
  // monitor writes to address zero.
  if (max == 0) {
    max += 4;
  }
  // The stack grow downwards towards _emscripten_stack_get_end.
  // We write cookies to the final two words in the stack and detect if they are
  // ever overwritten.
  HEAPU32[((max)>>2)] = 0x02135467;
  HEAPU32[(((max)+(4))>>2)] = 0x89BACDFE;
  // Also test the global address 0 for integrity.
  HEAPU32[((0)>>2)] = 1668509029;
}

function checkStackCookie() {
  if (ABORT) return;
  var max = _emscripten_stack_get_end();
  // See writeStackCookie().
  if (max == 0) {
    max += 4;
  }
  var cookie1 = HEAPU32[((max)>>2)];
  var cookie2 = HEAPU32[(((max)+(4))>>2)];
  if (cookie1 != 0x02135467 || cookie2 != 0x89BACDFE) {
    abort(`Stack overflow! Stack cookie has been overwritten at ${ptrToString(max)}, expected hex dwords 0x89BACDFE and 0x2135467, but received ${ptrToString(cookie2)} ${ptrToString(cookie1)}`);
  }
  // Also test the global address 0 for integrity.
  if (HEAPU32[((0)>>2)] != 0x63736d65 /* 'emsc' */) {
    abort('Runtime error: The application has corrupted its heap memory area (address zero)!');
  }
}
// end include: runtime_stack_check.js
// include: runtime_exceptions.js
// end include: runtime_exceptions.js
// include: runtime_debug.js
var runtimeDebug = true; // Switch to false at runtime to disable logging at the right times

// Used by XXXXX_DEBUG settings to output debug messages.
function dbg(...args) {
  if (!runtimeDebug && typeof runtimeDebug != 'undefined') return;
  // TODO(sbc): Make this configurable somehow.  Its not always convenient for
  // logging to show up as warnings.
  console.warn(...args);
}

// Endianness check
(() => {
  var h16 = new Int16Array(1);
  var h8 = new Int8Array(h16.buffer);
  h16[0] = 0x6373;
  if (h8[0] !== 0x73 || h8[1] !== 0x63) throw 'Runtime error: expected the system to be little-endian! (Run with -sSUPPORT_BIG_ENDIAN to bypass)';
})();

function consumedModuleProp(prop) {
  if (!Object.getOwnPropertyDescriptor(Module, prop)) {
    Object.defineProperty(Module, prop, {
      configurable: true,
      set() {
        abort(`Attempt to set \`Module.${prop}\` after it has already been processed.  This can happen, for example, when code is injected via '--post-js' rather than '--pre-js'`);

      }
    });
  }
}

function makeInvalidEarlyAccess(name) {
  return () => assert(false, `call to '${name}' via reference taken before Wasm module initialization`);

}

function ignoredModuleProp(prop) {
  if (Object.getOwnPropertyDescriptor(Module, prop)) {
    abort(`\`Module.${prop}\` was supplied but \`${prop}\` not included in INCOMING_MODULE_JS_API`);
  }
}

// forcing the filesystem exports a few things by default
function isExportedByForceFilesystem(name) {
  return name === 'FS_createPath' ||
         name === 'FS_createDataFile' ||
         name === 'FS_createPreloadedFile' ||
         name === 'FS_unlink' ||
         name === 'addRunDependency' ||
         // The old FS has some functionality that WasmFS lacks.
         name === 'FS_createLazyFile' ||
         name === 'FS_createDevice' ||
         name === 'removeRunDependency';
}

/**
 * Intercept access to a global symbol.  This enables us to give informative
 * warnings/errors when folks attempt to use symbols they did not include in
 * their build, or no symbols that no longer exist.
 */
function hookGlobalSymbolAccess(sym, func) {
  if (typeof globalThis != 'undefined' && !Object.getOwnPropertyDescriptor(globalThis, sym)) {
    Object.defineProperty(globalThis, sym, {
      configurable: true,
      get() {
        func();
        return undefined;
      }
    });
  }
}

function missingGlobal(sym, msg) {
  hookGlobalSymbolAccess(sym, () => {
    warnOnce(`\`${sym}\` is not longer defined by emscripten. ${msg}`);
  });
}

missingGlobal('buffer', 'Please use HEAP8.buffer or wasmMemory.buffer');
missingGlobal('asm', 'Please use wasmExports instead');

function missingLibrarySymbol(sym) {
  hookGlobalSymbolAccess(sym, () => {
    // Can't `abort()` here because it would break code that does runtime
    // checks.  e.g. `if (typeof SDL === 'undefined')`.
    var msg = `\`${sym}\` is a library symbol and not included by default; add it to your library.js __deps or to DEFAULT_LIBRARY_FUNCS_TO_INCLUDE on the command line`;
    // DEFAULT_LIBRARY_FUNCS_TO_INCLUDE requires the name as it appears in
    // library.js, which means $name for a JS name with no prefix, or name
    // for a JS name like _name.
    var librarySymbol = sym;
    if (!librarySymbol.startsWith('_')) {
      librarySymbol = '$' + sym;
    }
    msg += ` (e.g. -sDEFAULT_LIBRARY_FUNCS_TO_INCLUDE='${librarySymbol}')`;
    if (isExportedByForceFilesystem(sym)) {
      msg += '. Alternatively, forcing filesystem support (-sFORCE_FILESYSTEM) can export this for you';
    }
    warnOnce(msg);
  });

  // Any symbol that is not included from the JS library is also (by definition)
  // not exported on the Module object.
  unexportedRuntimeSymbol(sym);
}

function unexportedRuntimeSymbol(sym) {
  if (!Object.getOwnPropertyDescriptor(Module, sym)) {
    Object.defineProperty(Module, sym, {
      configurable: true,
      get() {
        var msg = `'${sym}' was not exported. add it to EXPORTED_RUNTIME_METHODS (see the Emscripten FAQ)`;
        if (isExportedByForceFilesystem(sym)) {
          msg += '. Alternatively, forcing filesystem support (-sFORCE_FILESYSTEM) can export this for you';
        }
        abort(msg);
      }
    });
  }
}

// end include: runtime_debug.js
// Memory management

var wasmMemory;

var
/** @type {!Int8Array} */
  HEAP8,
/** @type {!Uint8Array} */
  HEAPU8,
/** @type {!Int16Array} */
  HEAP16,
/** @type {!Uint16Array} */
  HEAPU16,
/** @type {!Int32Array} */
  HEAP32,
/** @type {!Uint32Array} */
  HEAPU32,
/** @type {!Float32Array} */
  HEAPF32,
/** @type {!Float64Array} */
  HEAPF64;

// BigInt64Array type is not correctly defined in closure
var
/** not-@type {!BigInt64Array} */
  HEAP64,
/* BigUint64Array type is not correctly defined in closure
/** not-@type {!BigUint64Array} */
  HEAPU64;

var runtimeInitialized = false;



function updateMemoryViews() {
  var b = wasmMemory.buffer;
  HEAP8 = new Int8Array(b);
  HEAP16 = new Int16Array(b);
  HEAPU8 = new Uint8Array(b);
  HEAPU16 = new Uint16Array(b);
  HEAP32 = new Int32Array(b);
  HEAPU32 = new Uint32Array(b);
  HEAPF32 = new Float32Array(b);
  HEAPF64 = new Float64Array(b);
  HEAP64 = new BigInt64Array(b);
  HEAPU64 = new BigUint64Array(b);
}

// include: memoryprofiler.js
// end include: memoryprofiler.js
// end include: runtime_common.js
assert(typeof Int32Array != 'undefined' && typeof Float64Array !== 'undefined' && Int32Array.prototype.subarray != undefined && Int32Array.prototype.set != undefined,
       'JS engine does not provide full typed array support');

function preRun() {
  if (Module['preRun']) {
    if (typeof Module['preRun'] == 'function') Module['preRun'] = [Module['preRun']];
    while (Module['preRun'].length) {
      addOnPreRun(Module['preRun'].shift());
    }
  }
  consumedModuleProp('preRun');
  // Begin ATPRERUNS hooks
  callRuntimeCallbacks(onPreRuns);
  // End ATPRERUNS hooks
}

function initRuntime() {
  assert(!runtimeInitialized);
  runtimeInitialized = true;

  checkStackCookie();

  // Begin ATINITS hooks
  if (!Module['noFSInit'] && !FS.initialized) FS.init();
TTY.init();
  // End ATINITS hooks

  wasmExports['__wasm_call_ctors']();

  // Begin ATPOSTCTORS hooks
  FS.ignorePermissions = false;
  // End ATPOSTCTORS hooks
}

function postRun() {
  checkStackCookie();
   // PThreads reuse the runtime from the main thread.

  if (Module['postRun']) {
    if (typeof Module['postRun'] == 'function') Module['postRun'] = [Module['postRun']];
    while (Module['postRun'].length) {
      addOnPostRun(Module['postRun'].shift());
    }
  }
  consumedModuleProp('postRun');

  // Begin ATPOSTRUNS hooks
  callRuntimeCallbacks(onPostRuns);
  // End ATPOSTRUNS hooks
}

// A counter of dependencies for calling run(). If we need to
// do asynchronous work before running, increment this and
// decrement it. Incrementing must happen in a place like
// Module.preRun (used by emcc to add file preloading).
// Note that you can add dependencies in preRun, even though
// it happens right before run - run will be postponed until
// the dependencies are met.
var runDependencies = 0;
var dependenciesFulfilled = null; // overridden to take different actions when all run dependencies are fulfilled
var runDependencyTracking = {};
var runDependencyWatcher = null;

function addRunDependency(id) {
  runDependencies++;

  Module['monitorRunDependencies']?.(runDependencies);

  if (id) {
    assert(!runDependencyTracking[id]);
    runDependencyTracking[id] = 1;
    if (runDependencyWatcher === null && typeof setInterval != 'undefined') {
      // Check for missing dependencies every few seconds
      runDependencyWatcher = setInterval(() => {
        if (ABORT) {
          clearInterval(runDependencyWatcher);
          runDependencyWatcher = null;
          return;
        }
        var shown = false;
        for (var dep in runDependencyTracking) {
          if (!shown) {
            shown = true;
            err('still waiting on run dependencies:');
          }
          err(`dependency: ${dep}`);
        }
        if (shown) {
          err('(end of list)');
        }
      }, 10000);
    }
  } else {
    err('warning: run dependency added without ID');
  }
}

function removeRunDependency(id) {
  runDependencies--;

  Module['monitorRunDependencies']?.(runDependencies);

  if (id) {
    assert(runDependencyTracking[id]);
    delete runDependencyTracking[id];
  } else {
    err('warning: run dependency removed without ID');
  }
  if (runDependencies == 0) {
    if (runDependencyWatcher !== null) {
      clearInterval(runDependencyWatcher);
      runDependencyWatcher = null;
    }
    if (dependenciesFulfilled) {
      var callback = dependenciesFulfilled;
      dependenciesFulfilled = null;
      callback(); // can add another dependenciesFulfilled
    }
  }
}

/** @param {string|number=} what */
function abort(what) {
  Module['onAbort']?.(what);

  what = 'Aborted(' + what + ')';
  // TODO(sbc): Should we remove printing and leave it up to whoever
  // catches the exception?
  err(what);

  ABORT = true;

  // Use a wasm runtime error, because a JS error might be seen as a foreign
  // exception, which means we'd run destructors on it. We need the error to
  // simply make the program stop.
  // FIXME This approach does not work in Wasm EH because it currently does not assume
  // all RuntimeErrors are from traps; it decides whether a RuntimeError is from
  // a trap or not based on a hidden field within the object. So at the moment
  // we don't have a way of throwing a wasm trap from JS. TODO Make a JS API that
  // allows this in the wasm spec.

  // Suppress closure compiler warning here. Closure compiler's builtin extern
  // definition for WebAssembly.RuntimeError claims it takes no arguments even
  // though it can.
  // TODO(https://github.com/google/closure-compiler/pull/3913): Remove if/when upstream closure gets fixed.
  /** @suppress {checkTypes} */
  var e = new WebAssembly.RuntimeError(what);

  // Throw the error whether or not MODULARIZE is set because abort is used
  // in code paths apart from instantiation where an exception is expected
  // to be thrown when abort is called.
  throw e;
}

function createExportWrapper(name, nargs) {
  return (...args) => {
    assert(runtimeInitialized, `native function \`${name}\` called before runtime initialization`);
    var f = wasmExports[name];
    assert(f, `exported native function \`${name}\` not found`);
    // Only assert for too many arguments. Too few can be valid since the missing arguments will be zero filled.
    assert(args.length <= nargs, `native function \`${name}\` called with ${args.length} args but expects ${nargs}`);
    return f(...args);
  };
}

var wasmBinaryFile;

function findWasmBinary() {
    return locateFile('sdl.wasm');
}

function getBinarySync(file) {
  if (file == wasmBinaryFile && wasmBinary) {
    return new Uint8Array(wasmBinary);
  }
  if (readBinary) {
    return readBinary(file);
  }
  throw 'both async and sync fetching of the wasm failed';
}

async function getWasmBinary(binaryFile) {
  // If we don't have the binary yet, load it asynchronously using readAsync.
  if (!wasmBinary) {
    // Fetch the binary using readAsync
    try {
      var response = await readAsync(binaryFile);
      return new Uint8Array(response);
    } catch {
      // Fall back to getBinarySync below;
    }
  }

  // Otherwise, getBinarySync should be able to get it synchronously
  return getBinarySync(binaryFile);
}

async function instantiateArrayBuffer(binaryFile, imports) {
  try {
    var binary = await getWasmBinary(binaryFile);
    var instance = await WebAssembly.instantiate(binary, imports);
    return instance;
  } catch (reason) {
    err(`failed to asynchronously prepare wasm: ${reason}`);

    // Warn on some common problems.
    if (isFileURI(wasmBinaryFile)) {
      err(`warning: Loading from a file URI (${wasmBinaryFile}) is not supported in most browsers. See https://emscripten.org/docs/getting_started/FAQ.html#how-do-i-run-a-local-webserver-for-testing-why-does-my-program-stall-in-downloading-or-preparing`);
    }
    abort(reason);
  }
}

async function instantiateAsync(binary, binaryFile, imports) {
  if (!binary && typeof WebAssembly.instantiateStreaming == 'function'
      // Don't use streaming for file:// delivered objects in a webview, fetch them synchronously.
      && !isFileURI(binaryFile)
      // Avoid instantiateStreaming() on Node.js environment for now, as while
      // Node.js v18.1.0 implements it, it does not have a full fetch()
      // implementation yet.
      //
      // Reference:
      //   https://github.com/emscripten-core/emscripten/pull/16917
      && !ENVIRONMENT_IS_NODE
     ) {
    try {
      var response = fetch(binaryFile, { credentials: 'same-origin' });
      var instantiationResult = await WebAssembly.instantiateStreaming(response, imports);
      return instantiationResult;
    } catch (reason) {
      // We expect the most common failure cause to be a bad MIME type for the binary,
      // in which case falling back to ArrayBuffer instantiation should work.
      err(`wasm streaming compile failed: ${reason}`);
      err('falling back to ArrayBuffer instantiation');
      // fall back of instantiateArrayBuffer below
    };
  }
  return instantiateArrayBuffer(binaryFile, imports);
}

function getWasmImports() {
  // prepare imports
  return {
    'env': wasmImports,
    'wasi_snapshot_preview1': wasmImports,
  }
}

// Create the wasm instance.
// Receives the wasm imports, returns the exports.
async function createWasm() {
  // Load the wasm module and create an instance of using native support in the JS engine.
  // handle a generated wasm instance, receiving its exports and
  // performing other necessary setup
  /** @param {WebAssembly.Module=} module*/
  function receiveInstance(instance, module) {
    wasmExports = instance.exports;

    

    wasmMemory = wasmExports['memory'];
    
    assert(wasmMemory, 'memory not found in wasm exports');
    updateMemoryViews();

    wasmTable = wasmExports['__indirect_function_table'];
    
    assert(wasmTable, 'table not found in wasm exports');

    assignWasmExports(wasmExports);
    removeRunDependency('wasm-instantiate');
    return wasmExports;
  }
  // wait for the pthread pool (if any)
  addRunDependency('wasm-instantiate');

  // Prefer streaming instantiation if available.
  // Async compilation can be confusing when an error on the page overwrites Module
  // (for example, if the order of elements is wrong, and the one defining Module is
  // later), so we save Module and check it later.
  var trueModule = Module;
  function receiveInstantiationResult(result) {
    // 'result' is a ResultObject object which has both the module and instance.
    // receiveInstance() will swap in the exports (to Module.asm) so they can be called
    assert(Module === trueModule, 'the Module object should not be replaced during async compilation - perhaps the order of HTML elements is wrong?');
    trueModule = null;
    // TODO: Due to Closure regression https://github.com/google/closure-compiler/issues/3193, the above line no longer optimizes out down to the following line.
    // When the regression is fixed, can restore the above PTHREADS-enabled path.
    return receiveInstance(result['instance']);
  }

  var info = getWasmImports();

  // User shell pages can write their own Module.instantiateWasm = function(imports, successCallback) callback
  // to manually instantiate the Wasm module themselves. This allows pages to
  // run the instantiation parallel to any other async startup actions they are
  // performing.
  // Also pthreads and wasm workers initialize the wasm instance through this
  // path.
  if (Module['instantiateWasm']) {
    return new Promise((resolve, reject) => {
      try {
        Module['instantiateWasm'](info, (mod, inst) => {
          resolve(receiveInstance(mod, inst));
        });
      } catch(e) {
        err(`Module.instantiateWasm callback failed with error: ${e}`);
        reject(e);
      }
    });
  }

  wasmBinaryFile ??= findWasmBinary();
  var result = await instantiateAsync(wasmBinary, wasmBinaryFile, info);
  var exports = receiveInstantiationResult(result);
  return exports;
}

// end include: preamble.js

// Begin JS library code


  class ExitStatus {
      name = 'ExitStatus';
      constructor(status) {
        this.message = `Program terminated with exit(${status})`;
        this.status = status;
      }
    }

  var uleb128EncodeWithLen = (arr) => {
      const n = arr.length;
      assert(n < 16384);
      // Note: this LEB128 length encoding produces extra byte for n < 128,
      // but we don't care as it's only used in a temporary representation.
      return [(n % 128) | 128, n >> 7, ...arr];
    };
  
  
  var wasmTypeCodes = {
      'i': 0x7f, // i32
      'p': 0x7f, // i32
      'j': 0x7e, // i64
      'f': 0x7d, // f32
      'd': 0x7c, // f64
      'e': 0x6f, // externref
    };
  var generateTypePack = (types) => uleb128EncodeWithLen(Array.from(types, (type) => {
      var code = wasmTypeCodes[type];
      assert(code, `invalid signature char: ${type}`);
      return code;
    }));
  var convertJsFunctionToWasm = (func, sig) => {
  
      // Rest of the module is static
      var bytes = Uint8Array.of(
        0x00, 0x61, 0x73, 0x6d, // magic ("\0asm")
        0x01, 0x00, 0x00, 0x00, // version: 1
        0x01, // Type section code
          // The module is static, with the exception of the type section, which is
          // generated based on the signature passed in.
          ...uleb128EncodeWithLen([
            0x01, // count: 1
            0x60 /* form: func */,
            // param types
            ...generateTypePack(sig.slice(1)),
            // return types (for now only supporting [] if `void` and single [T] otherwise)
            ...generateTypePack(sig[0] === 'v' ? '' : sig[0])
          ]),
        // The rest of the module is static
        0x02, 0x07, // import section
          // (import "e" "f" (func 0 (type 0)))
          0x01, 0x01, 0x65, 0x01, 0x66, 0x00, 0x00,
        0x07, 0x05, // export section
          // (export "f" (func 0 (type 0)))
          0x01, 0x01, 0x66, 0x00, 0x00,
      );
  
      // We can compile this wasm module synchronously because it is very small.
      // This accepts an import (at "e.f"), that it reroutes to an export (at "f")
      var module = new WebAssembly.Module(bytes);
      var instance = new WebAssembly.Instance(module, { 'e': { 'f': func } });
      var wrappedFunc = instance.exports['f'];
      return wrappedFunc;
    };
  
  var wasmTableMirror = [];
  
  /** @type {WebAssembly.Table} */
  var wasmTable;
  var getWasmTableEntry = (funcPtr) => {
      var func = wasmTableMirror[funcPtr];
      if (!func) {
        /** @suppress {checkTypes} */
        wasmTableMirror[funcPtr] = func = wasmTable.get(funcPtr);
      }
      /** @suppress {checkTypes} */
      assert(wasmTable.get(funcPtr) == func, 'JavaScript-side Wasm function table mirror is out of date!');
      return func;
    };
  
  var updateTableMap = (offset, count) => {
      if (functionsInTableMap) {
        for (var i = offset; i < offset + count; i++) {
          var item = getWasmTableEntry(i);
          // Ignore null values.
          if (item) {
            functionsInTableMap.set(item, i);
          }
        }
      }
    };
  
  var functionsInTableMap;
  
  var getFunctionAddress = (func) => {
      // First, create the map if this is the first use.
      if (!functionsInTableMap) {
        functionsInTableMap = new WeakMap();
        updateTableMap(0, wasmTable.length);
      }
      return functionsInTableMap.get(func) || 0;
    };
  
  
  var freeTableIndexes = [];
  
  var getEmptyTableSlot = () => {
      // Reuse a free index if there is one, otherwise grow.
      if (freeTableIndexes.length) {
        return freeTableIndexes.pop();
      }
      try {
        // Grow the table
        return wasmTable['grow'](1);
      } catch (err) {
        if (!(err instanceof RangeError)) {
          throw err;
        }
        throw 'Unable to grow wasm table. Set ALLOW_TABLE_GROWTH.';
      }
    };
  
  
  var setWasmTableEntry = (idx, func) => {
      /** @suppress {checkTypes} */
      wasmTable.set(idx, func);
      // With ABORT_ON_WASM_EXCEPTIONS wasmTable.get is overridden to return wrapped
      // functions so we need to call it here to retrieve the potential wrapper correctly
      // instead of just storing 'func' directly into wasmTableMirror
      /** @suppress {checkTypes} */
      wasmTableMirror[idx] = wasmTable.get(idx);
    };
  /** @param {string=} sig */
  var addFunction = (func, sig) => {
      assert(typeof func != 'undefined');
      // Check if the function is already in the table, to ensure each function
      // gets a unique index.
      var rtn = getFunctionAddress(func);
      if (rtn) {
        return rtn;
      }
  
      // It's not in the table, add it now.
  
      var ret = getEmptyTableSlot();
  
      // Set the new value.
      try {
        // Attempting to call this with JS function will cause of table.set() to fail
        setWasmTableEntry(ret, func);
      } catch (err) {
        if (!(err instanceof TypeError)) {
          throw err;
        }
        assert(typeof sig != 'undefined', 'Missing signature argument to addFunction: ' + func);
        var wrapped = convertJsFunctionToWasm(func, sig);
        setWasmTableEntry(ret, wrapped);
      }
  
      functionsInTableMap.set(func, ret);
  
      return ret;
    };

  var callRuntimeCallbacks = (callbacks) => {
      while (callbacks.length > 0) {
        // Pass the module as the first argument.
        callbacks.shift()(Module);
      }
    };
  var onPostRuns = [];
  var addOnPostRun = (cb) => onPostRuns.push(cb);

  var onPreRuns = [];
  var addOnPreRun = (cb) => onPreRuns.push(cb);


  
    /**
     * @param {number} ptr
     * @param {string} type
     */
  function getValue(ptr, type = 'i8') {
    if (type.endsWith('*')) type = '*';
    switch (type) {
      case 'i1': return HEAP8[ptr];
      case 'i8': return HEAP8[ptr];
      case 'i16': return HEAP16[((ptr)>>1)];
      case 'i32': return HEAP32[((ptr)>>2)];
      case 'i64': return HEAP64[((ptr)>>3)];
      case 'float': return HEAPF32[((ptr)>>2)];
      case 'double': return HEAPF64[((ptr)>>3)];
      case '*': return HEAPU32[((ptr)>>2)];
      default: abort(`invalid type for getValue: ${type}`);
    }
  }

  var noExitRuntime = true;

  var ptrToString = (ptr) => {
      assert(typeof ptr === 'number');
      // With CAN_ADDRESS_2GB or MEMORY64, pointers are already unsigned.
      ptr >>>= 0;
      return '0x' + ptr.toString(16).padStart(8, '0');
    };

  
    /**
     * @param {number} ptr
     * @param {number} value
     * @param {string} type
     */
  function setValue(ptr, value, type = 'i8') {
    if (type.endsWith('*')) type = '*';
    switch (type) {
      case 'i1': HEAP8[ptr] = value; break;
      case 'i8': HEAP8[ptr] = value; break;
      case 'i16': HEAP16[((ptr)>>1)] = value; break;
      case 'i32': HEAP32[((ptr)>>2)] = value; break;
      case 'i64': HEAP64[((ptr)>>3)] = BigInt(value); break;
      case 'float': HEAPF32[((ptr)>>2)] = value; break;
      case 'double': HEAPF64[((ptr)>>3)] = value; break;
      case '*': HEAPU32[((ptr)>>2)] = value; break;
      default: abort(`invalid type for setValue: ${type}`);
    }
  }

  var stackRestore = (val) => __emscripten_stack_restore(val);

  var stackSave = () => _emscripten_stack_get_current();

  var warnOnce = (text) => {
      warnOnce.shown ||= {};
      if (!warnOnce.shown[text]) {
        warnOnce.shown[text] = 1;
        if (ENVIRONMENT_IS_NODE) text = 'warning: ' + text;
        err(text);
      }
    };

  var ___call_sighandler = (fp, sig) => getWasmTableEntry(fp)(sig);

  /** @suppress {duplicate } */
  var syscallGetVarargI = () => {
      assert(SYSCALLS.varargs != undefined);
      // the `+` prepended here is necessary to convince the JSCompiler that varargs is indeed a number.
      var ret = HEAP32[((+SYSCALLS.varargs)>>2)];
      SYSCALLS.varargs += 4;
      return ret;
    };
  var syscallGetVarargP = syscallGetVarargI;
  
  
  var PATH = {
  isAbs:(path) => path.charAt(0) === '/',
  splitPath:(filename) => {
        var splitPathRe = /^(\/?|)([\s\S]*?)((?:\.{1,2}|[^\/]+?|)(\.[^.\/]*|))(?:[\/]*)$/;
        return splitPathRe.exec(filename).slice(1);
      },
  normalizeArray:(parts, allowAboveRoot) => {
        // if the path tries to go above the root, `up` ends up > 0
        var up = 0;
        for (var i = parts.length - 1; i >= 0; i--) {
          var last = parts[i];
          if (last === '.') {
            parts.splice(i, 1);
          } else if (last === '..') {
            parts.splice(i, 1);
            up++;
          } else if (up) {
            parts.splice(i, 1);
            up--;
          }
        }
        // if the path is allowed to go above the root, restore leading ..s
        if (allowAboveRoot) {
          for (; up; up--) {
            parts.unshift('..');
          }
        }
        return parts;
      },
  normalize:(path) => {
        var isAbsolute = PATH.isAbs(path),
            trailingSlash = path.slice(-1) === '/';
        // Normalize the path
        path = PATH.normalizeArray(path.split('/').filter((p) => !!p), !isAbsolute).join('/');
        if (!path && !isAbsolute) {
          path = '.';
        }
        if (path && trailingSlash) {
          path += '/';
        }
        return (isAbsolute ? '/' : '') + path;
      },
  dirname:(path) => {
        var result = PATH.splitPath(path),
            root = result[0],
            dir = result[1];
        if (!root && !dir) {
          // No dirname whatsoever
          return '.';
        }
        if (dir) {
          // It has a dirname, strip trailing slash
          dir = dir.slice(0, -1);
        }
        return root + dir;
      },
  basename:(path) => path && path.match(/([^\/]+|\/)\/*$/)[1],
  join:(...paths) => PATH.normalize(paths.join('/')),
  join2:(l, r) => PATH.normalize(l + '/' + r),
  };
  
  var initRandomFill = () => {
      // This block is not needed on v19+ since crypto.getRandomValues is builtin
      if (ENVIRONMENT_IS_NODE) {
        var nodeCrypto = require('crypto');
        return (view) => nodeCrypto.randomFillSync(view);
      }
  
      return (view) => crypto.getRandomValues(view);
    };
  var randomFill = (view) => {
      // Lazily init on the first invocation.
      (randomFill = initRandomFill())(view);
    };
  
  
  
  var PATH_FS = {
  resolve:(...args) => {
        var resolvedPath = '',
          resolvedAbsolute = false;
        for (var i = args.length - 1; i >= -1 && !resolvedAbsolute; i--) {
          var path = (i >= 0) ? args[i] : FS.cwd();
          // Skip empty and invalid entries
          if (typeof path != 'string') {
            throw new TypeError('Arguments to path.resolve must be strings');
          } else if (!path) {
            return ''; // an invalid portion invalidates the whole thing
          }
          resolvedPath = path + '/' + resolvedPath;
          resolvedAbsolute = PATH.isAbs(path);
        }
        // At this point the path should be resolved to a full absolute path, but
        // handle relative paths to be safe (might happen when process.cwd() fails)
        resolvedPath = PATH.normalizeArray(resolvedPath.split('/').filter((p) => !!p), !resolvedAbsolute).join('/');
        return ((resolvedAbsolute ? '/' : '') + resolvedPath) || '.';
      },
  relative:(from, to) => {
        from = PATH_FS.resolve(from).slice(1);
        to = PATH_FS.resolve(to).slice(1);
        function trim(arr) {
          var start = 0;
          for (; start < arr.length; start++) {
            if (arr[start] !== '') break;
          }
          var end = arr.length - 1;
          for (; end >= 0; end--) {
            if (arr[end] !== '') break;
          }
          if (start > end) return [];
          return arr.slice(start, end - start + 1);
        }
        var fromParts = trim(from.split('/'));
        var toParts = trim(to.split('/'));
        var length = Math.min(fromParts.length, toParts.length);
        var samePartsLength = length;
        for (var i = 0; i < length; i++) {
          if (fromParts[i] !== toParts[i]) {
            samePartsLength = i;
            break;
          }
        }
        var outputParts = [];
        for (var i = samePartsLength; i < fromParts.length; i++) {
          outputParts.push('..');
        }
        outputParts = outputParts.concat(toParts.slice(samePartsLength));
        return outputParts.join('/');
      },
  };
  
  
  var UTF8Decoder = typeof TextDecoder != 'undefined' ? new TextDecoder() : undefined;
  
  var findStringEnd = (heapOrArray, idx, maxBytesToRead, ignoreNul) => {
      var maxIdx = idx + maxBytesToRead;
      if (ignoreNul) return maxIdx;
      // TextDecoder needs to know the byte length in advance, it doesn't stop on
      // null terminator by itself.
      // As a tiny code save trick, compare idx against maxIdx using a negation,
      // so that maxBytesToRead=undefined/NaN means Infinity.
      while (heapOrArray[idx] && !(idx >= maxIdx)) ++idx;
      return idx;
    };
  
    /**
     * Given a pointer 'idx' to a null-terminated UTF8-encoded string in the given
     * array that contains uint8 values, returns a copy of that string as a
     * Javascript String object.
     * heapOrArray is either a regular array, or a JavaScript typed array view.
     * @param {number=} idx
     * @param {number=} maxBytesToRead
     * @param {boolean=} ignoreNul - If true, the function will not stop on a NUL character.
     * @return {string}
     */
  var UTF8ArrayToString = (heapOrArray, idx = 0, maxBytesToRead, ignoreNul) => {
  
      var endPtr = findStringEnd(heapOrArray, idx, maxBytesToRead, ignoreNul);
  
      // When using conditional TextDecoder, skip it for short strings as the overhead of the native call is not worth it.
      if (endPtr - idx > 16 && heapOrArray.buffer && UTF8Decoder) {
        return UTF8Decoder.decode(heapOrArray.subarray(idx, endPtr));
      }
      var str = '';
      // If building with TextDecoder, we have already computed the string length
      // above, so test loop end condition against that
      while (idx < endPtr) {
        // For UTF8 byte structure, see:
        // http://en.wikipedia.org/wiki/UTF-8#Description
        // https://www.ietf.org/rfc/rfc2279.txt
        // https://tools.ietf.org/html/rfc3629
        var u0 = heapOrArray[idx++];
        if (!(u0 & 0x80)) { str += String.fromCharCode(u0); continue; }
        var u1 = heapOrArray[idx++] & 63;
        if ((u0 & 0xE0) == 0xC0) { str += String.fromCharCode(((u0 & 31) << 6) | u1); continue; }
        var u2 = heapOrArray[idx++] & 63;
        if ((u0 & 0xF0) == 0xE0) {
          u0 = ((u0 & 15) << 12) | (u1 << 6) | u2;
        } else {
          if ((u0 & 0xF8) != 0xF0) warnOnce('Invalid UTF-8 leading byte ' + ptrToString(u0) + ' encountered when deserializing a UTF-8 string in wasm memory to a JS string!');
          u0 = ((u0 & 7) << 18) | (u1 << 12) | (u2 << 6) | (heapOrArray[idx++] & 63);
        }
  
        if (u0 < 0x10000) {
          str += String.fromCharCode(u0);
        } else {
          var ch = u0 - 0x10000;
          str += String.fromCharCode(0xD800 | (ch >> 10), 0xDC00 | (ch & 0x3FF));
        }
      }
      return str;
    };
  
  var FS_stdin_getChar_buffer = [];
  
  var lengthBytesUTF8 = (str) => {
      var len = 0;
      for (var i = 0; i < str.length; ++i) {
        // Gotcha: charCodeAt returns a 16-bit word that is a UTF-16 encoded code
        // unit, not a Unicode code point of the character! So decode
        // UTF16->UTF32->UTF8.
        // See http://unicode.org/faq/utf_bom.html#utf16-3
        var c = str.charCodeAt(i); // possibly a lead surrogate
        if (c <= 0x7F) {
          len++;
        } else if (c <= 0x7FF) {
          len += 2;
        } else if (c >= 0xD800 && c <= 0xDFFF) {
          len += 4; ++i;
        } else {
          len += 3;
        }
      }
      return len;
    };
  
  var stringToUTF8Array = (str, heap, outIdx, maxBytesToWrite) => {
      assert(typeof str === 'string', `stringToUTF8Array expects a string (got ${typeof str})`);
      // Parameter maxBytesToWrite is not optional. Negative values, 0, null,
      // undefined and false each don't write out any bytes.
      if (!(maxBytesToWrite > 0))
        return 0;
  
      var startIdx = outIdx;
      var endIdx = outIdx + maxBytesToWrite - 1; // -1 for string null terminator.
      for (var i = 0; i < str.length; ++i) {
        // For UTF8 byte structure, see http://en.wikipedia.org/wiki/UTF-8#Description
        // and https://www.ietf.org/rfc/rfc2279.txt
        // and https://tools.ietf.org/html/rfc3629
        var u = str.codePointAt(i);
        if (u <= 0x7F) {
          if (outIdx >= endIdx) break;
          heap[outIdx++] = u;
        } else if (u <= 0x7FF) {
          if (outIdx + 1 >= endIdx) break;
          heap[outIdx++] = 0xC0 | (u >> 6);
          heap[outIdx++] = 0x80 | (u & 63);
        } else if (u <= 0xFFFF) {
          if (outIdx + 2 >= endIdx) break;
          heap[outIdx++] = 0xE0 | (u >> 12);
          heap[outIdx++] = 0x80 | ((u >> 6) & 63);
          heap[outIdx++] = 0x80 | (u & 63);
        } else {
          if (outIdx + 3 >= endIdx) break;
          if (u > 0x10FFFF) warnOnce('Invalid Unicode code point ' + ptrToString(u) + ' encountered when serializing a JS string to a UTF-8 string in wasm memory! (Valid unicode code points should be in range 0-0x10FFFF).');
          heap[outIdx++] = 0xF0 | (u >> 18);
          heap[outIdx++] = 0x80 | ((u >> 12) & 63);
          heap[outIdx++] = 0x80 | ((u >> 6) & 63);
          heap[outIdx++] = 0x80 | (u & 63);
          // Gotcha: if codePoint is over 0xFFFF, it is represented as a surrogate pair in UTF-16.
          // We need to manually skip over the second code unit for correct iteration.
          i++;
        }
      }
      // Null-terminate the pointer to the buffer.
      heap[outIdx] = 0;
      return outIdx - startIdx;
    };
  /** @type {function(string, boolean=, number=)} */
  var intArrayFromString = (stringy, dontAddNull, length) => {
      var len = length > 0 ? length : lengthBytesUTF8(stringy)+1;
      var u8array = new Array(len);
      var numBytesWritten = stringToUTF8Array(stringy, u8array, 0, u8array.length);
      if (dontAddNull) u8array.length = numBytesWritten;
      return u8array;
    };
  var FS_stdin_getChar = () => {
      if (!FS_stdin_getChar_buffer.length) {
        var result = null;
        if (ENVIRONMENT_IS_NODE) {
          // we will read data by chunks of BUFSIZE
          var BUFSIZE = 256;
          var buf = Buffer.alloc(BUFSIZE);
          var bytesRead = 0;
  
          // For some reason we must suppress a closure warning here, even though
          // fd definitely exists on process.stdin, and is even the proper way to
          // get the fd of stdin,
          // https://github.com/nodejs/help/issues/2136#issuecomment-523649904
          // This started to happen after moving this logic out of library_tty.js,
          // so it is related to the surrounding code in some unclear manner.
          /** @suppress {missingProperties} */
          var fd = process.stdin.fd;
  
          try {
            bytesRead = fs.readSync(fd, buf, 0, BUFSIZE);
          } catch(e) {
            // Cross-platform differences: on Windows, reading EOF throws an
            // exception, but on other OSes, reading EOF returns 0. Uniformize
            // behavior by treating the EOF exception to return 0.
            if (e.toString().includes('EOF')) bytesRead = 0;
            else throw e;
          }
  
          if (bytesRead > 0) {
            result = buf.slice(0, bytesRead).toString('utf-8');
          }
        } else
        if (typeof window != 'undefined' &&
          typeof window.prompt == 'function') {
          // Browser.
          result = window.prompt('Input: ');  // returns null on cancel
          if (result !== null) {
            result += '\n';
          }
        } else
        {}
        if (!result) {
          return null;
        }
        FS_stdin_getChar_buffer = intArrayFromString(result, true);
      }
      return FS_stdin_getChar_buffer.shift();
    };
  var TTY = {
  ttys:[],
  init() {
        // https://github.com/emscripten-core/emscripten/pull/1555
        // if (ENVIRONMENT_IS_NODE) {
        //   // currently, FS.init does not distinguish if process.stdin is a file or TTY
        //   // device, it always assumes it's a TTY device. because of this, we're forcing
        //   // process.stdin to UTF8 encoding to at least make stdin reading compatible
        //   // with text files until FS.init can be refactored.
        //   process.stdin.setEncoding('utf8');
        // }
      },
  shutdown() {
        // https://github.com/emscripten-core/emscripten/pull/1555
        // if (ENVIRONMENT_IS_NODE) {
        //   // inolen: any idea as to why node -e 'process.stdin.read()' wouldn't exit immediately (with process.stdin being a tty)?
        //   // isaacs: because now it's reading from the stream, you've expressed interest in it, so that read() kicks off a _read() which creates a ReadReq operation
        //   // inolen: I thought read() in that case was a synchronous operation that just grabbed some amount of buffered data if it exists?
        //   // isaacs: it is. but it also triggers a _read() call, which calls readStart() on the handle
        //   // isaacs: do process.stdin.pause() and i'd think it'd probably close the pending call
        //   process.stdin.pause();
        // }
      },
  register(dev, ops) {
        TTY.ttys[dev] = { input: [], output: [], ops: ops };
        FS.registerDevice(dev, TTY.stream_ops);
      },
  stream_ops:{
  open(stream) {
          var tty = TTY.ttys[stream.node.rdev];
          if (!tty) {
            throw new FS.ErrnoError(43);
          }
          stream.tty = tty;
          stream.seekable = false;
        },
  close(stream) {
          // flush any pending line data
          stream.tty.ops.fsync(stream.tty);
        },
  fsync(stream) {
          stream.tty.ops.fsync(stream.tty);
        },
  read(stream, buffer, offset, length, pos /* ignored */) {
          if (!stream.tty || !stream.tty.ops.get_char) {
            throw new FS.ErrnoError(60);
          }
          var bytesRead = 0;
          for (var i = 0; i < length; i++) {
            var result;
            try {
              result = stream.tty.ops.get_char(stream.tty);
            } catch (e) {
              throw new FS.ErrnoError(29);
            }
            if (result === undefined && bytesRead === 0) {
              throw new FS.ErrnoError(6);
            }
            if (result === null || result === undefined) break;
            bytesRead++;
            buffer[offset+i] = result;
          }
          if (bytesRead) {
            stream.node.atime = Date.now();
          }
          return bytesRead;
        },
  write(stream, buffer, offset, length, pos) {
          if (!stream.tty || !stream.tty.ops.put_char) {
            throw new FS.ErrnoError(60);
          }
          try {
            for (var i = 0; i < length; i++) {
              stream.tty.ops.put_char(stream.tty, buffer[offset+i]);
            }
          } catch (e) {
            throw new FS.ErrnoError(29);
          }
          if (length) {
            stream.node.mtime = stream.node.ctime = Date.now();
          }
          return i;
        },
  },
  default_tty_ops:{
  get_char(tty) {
          return FS_stdin_getChar();
        },
  put_char(tty, val) {
          if (val === null || val === 10) {
            out(UTF8ArrayToString(tty.output));
            tty.output = [];
          } else {
            if (val != 0) tty.output.push(val); // val == 0 would cut text output off in the middle.
          }
        },
  fsync(tty) {
          if (tty.output?.length > 0) {
            out(UTF8ArrayToString(tty.output));
            tty.output = [];
          }
        },
  ioctl_tcgets(tty) {
          // typical setting
          return {
            c_iflag: 25856,
            c_oflag: 5,
            c_cflag: 191,
            c_lflag: 35387,
            c_cc: [
              0x03, 0x1c, 0x7f, 0x15, 0x04, 0x00, 0x01, 0x00, 0x11, 0x13, 0x1a, 0x00,
              0x12, 0x0f, 0x17, 0x16, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
              0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
            ]
          };
        },
  ioctl_tcsets(tty, optional_actions, data) {
          // currently just ignore
          return 0;
        },
  ioctl_tiocgwinsz(tty) {
          return [24, 80];
        },
  },
  default_tty1_ops:{
  put_char(tty, val) {
          if (val === null || val === 10) {
            err(UTF8ArrayToString(tty.output));
            tty.output = [];
          } else {
            if (val != 0) tty.output.push(val);
          }
        },
  fsync(tty) {
          if (tty.output?.length > 0) {
            err(UTF8ArrayToString(tty.output));
            tty.output = [];
          }
        },
  },
  };
  
  
  var zeroMemory = (ptr, size) => HEAPU8.fill(0, ptr, ptr + size);
  
  var alignMemory = (size, alignment) => {
      assert(alignment, "alignment argument is required");
      return Math.ceil(size / alignment) * alignment;
    };
  var mmapAlloc = (size) => {
      size = alignMemory(size, 65536);
      var ptr = _emscripten_builtin_memalign(65536, size);
      if (ptr) zeroMemory(ptr, size);
      return ptr;
    };
  var MEMFS = {
  ops_table:null,
  mount(mount) {
        return MEMFS.createNode(null, '/', 16895, 0);
      },
  createNode(parent, name, mode, dev) {
        if (FS.isBlkdev(mode) || FS.isFIFO(mode)) {
          // no supported
          throw new FS.ErrnoError(63);
        }
        MEMFS.ops_table ||= {
          dir: {
            node: {
              getattr: MEMFS.node_ops.getattr,
              setattr: MEMFS.node_ops.setattr,
              lookup: MEMFS.node_ops.lookup,
              mknod: MEMFS.node_ops.mknod,
              rename: MEMFS.node_ops.rename,
              unlink: MEMFS.node_ops.unlink,
              rmdir: MEMFS.node_ops.rmdir,
              readdir: MEMFS.node_ops.readdir,
              symlink: MEMFS.node_ops.symlink
            },
            stream: {
              llseek: MEMFS.stream_ops.llseek
            }
          },
          file: {
            node: {
              getattr: MEMFS.node_ops.getattr,
              setattr: MEMFS.node_ops.setattr
            },
            stream: {
              llseek: MEMFS.stream_ops.llseek,
              read: MEMFS.stream_ops.read,
              write: MEMFS.stream_ops.write,
              mmap: MEMFS.stream_ops.mmap,
              msync: MEMFS.stream_ops.msync
            }
          },
          link: {
            node: {
              getattr: MEMFS.node_ops.getattr,
              setattr: MEMFS.node_ops.setattr,
              readlink: MEMFS.node_ops.readlink
            },
            stream: {}
          },
          chrdev: {
            node: {
              getattr: MEMFS.node_ops.getattr,
              setattr: MEMFS.node_ops.setattr
            },
            stream: FS.chrdev_stream_ops
          }
        };
        var node = FS.createNode(parent, name, mode, dev);
        if (FS.isDir(node.mode)) {
          node.node_ops = MEMFS.ops_table.dir.node;
          node.stream_ops = MEMFS.ops_table.dir.stream;
          node.contents = {};
        } else if (FS.isFile(node.mode)) {
          node.node_ops = MEMFS.ops_table.file.node;
          node.stream_ops = MEMFS.ops_table.file.stream;
          node.usedBytes = 0; // The actual number of bytes used in the typed array, as opposed to contents.length which gives the whole capacity.
          // When the byte data of the file is populated, this will point to either a typed array, or a normal JS array. Typed arrays are preferred
          // for performance, and used by default. However, typed arrays are not resizable like normal JS arrays are, so there is a small disk size
          // penalty involved for appending file writes that continuously grow a file similar to std::vector capacity vs used -scheme.
          node.contents = null; 
        } else if (FS.isLink(node.mode)) {
          node.node_ops = MEMFS.ops_table.link.node;
          node.stream_ops = MEMFS.ops_table.link.stream;
        } else if (FS.isChrdev(node.mode)) {
          node.node_ops = MEMFS.ops_table.chrdev.node;
          node.stream_ops = MEMFS.ops_table.chrdev.stream;
        }
        node.atime = node.mtime = node.ctime = Date.now();
        // add the new node to the parent
        if (parent) {
          parent.contents[name] = node;
          parent.atime = parent.mtime = parent.ctime = node.atime;
        }
        return node;
      },
  getFileDataAsTypedArray(node) {
        if (!node.contents) return new Uint8Array(0);
        if (node.contents.subarray) return node.contents.subarray(0, node.usedBytes); // Make sure to not return excess unused bytes.
        return new Uint8Array(node.contents);
      },
  expandFileStorage(node, newCapacity) {
        var prevCapacity = node.contents ? node.contents.length : 0;
        if (prevCapacity >= newCapacity) return; // No need to expand, the storage was already large enough.
        // Don't expand strictly to the given requested limit if it's only a very small increase, but instead geometrically grow capacity.
        // For small filesizes (<1MB), perform size*2 geometric increase, but for large sizes, do a much more conservative size*1.125 increase to
        // avoid overshooting the allocation cap by a very large margin.
        var CAPACITY_DOUBLING_MAX = 1024 * 1024;
        newCapacity = Math.max(newCapacity, (prevCapacity * (prevCapacity < CAPACITY_DOUBLING_MAX ? 2.0 : 1.125)) >>> 0);
        if (prevCapacity != 0) newCapacity = Math.max(newCapacity, 256); // At minimum allocate 256b for each file when expanding.
        var oldContents = node.contents;
        node.contents = new Uint8Array(newCapacity); // Allocate new storage.
        if (node.usedBytes > 0) node.contents.set(oldContents.subarray(0, node.usedBytes), 0); // Copy old data over to the new storage.
      },
  resizeFileStorage(node, newSize) {
        if (node.usedBytes == newSize) return;
        if (newSize == 0) {
          node.contents = null; // Fully decommit when requesting a resize to zero.
          node.usedBytes = 0;
        } else {
          var oldContents = node.contents;
          node.contents = new Uint8Array(newSize); // Allocate new storage.
          if (oldContents) {
            node.contents.set(oldContents.subarray(0, Math.min(newSize, node.usedBytes))); // Copy old data over to the new storage.
          }
          node.usedBytes = newSize;
        }
      },
  node_ops:{
  getattr(node) {
          var attr = {};
          // device numbers reuse inode numbers.
          attr.dev = FS.isChrdev(node.mode) ? node.id : 1;
          attr.ino = node.id;
          attr.mode = node.mode;
          attr.nlink = 1;
          attr.uid = 0;
          attr.gid = 0;
          attr.rdev = node.rdev;
          if (FS.isDir(node.mode)) {
            attr.size = 4096;
          } else if (FS.isFile(node.mode)) {
            attr.size = node.usedBytes;
          } else if (FS.isLink(node.mode)) {
            attr.size = node.link.length;
          } else {
            attr.size = 0;
          }
          attr.atime = new Date(node.atime);
          attr.mtime = new Date(node.mtime);
          attr.ctime = new Date(node.ctime);
          // NOTE: In our implementation, st_blocks = Math.ceil(st_size/st_blksize),
          //       but this is not required by the standard.
          attr.blksize = 4096;
          attr.blocks = Math.ceil(attr.size / attr.blksize);
          return attr;
        },
  setattr(node, attr) {
          for (const key of ["mode", "atime", "mtime", "ctime"]) {
            if (attr[key] != null) {
              node[key] = attr[key];
            }
          }
          if (attr.size !== undefined) {
            MEMFS.resizeFileStorage(node, attr.size);
          }
        },
  lookup(parent, name) {
          throw new FS.ErrnoError(44);
        },
  mknod(parent, name, mode, dev) {
          return MEMFS.createNode(parent, name, mode, dev);
        },
  rename(old_node, new_dir, new_name) {
          var new_node;
          try {
            new_node = FS.lookupNode(new_dir, new_name);
          } catch (e) {}
          if (new_node) {
            if (FS.isDir(old_node.mode)) {
              // if we're overwriting a directory at new_name, make sure it's empty.
              for (var i in new_node.contents) {
                throw new FS.ErrnoError(55);
              }
            }
            FS.hashRemoveNode(new_node);
          }
          // do the internal rewiring
          delete old_node.parent.contents[old_node.name];
          new_dir.contents[new_name] = old_node;
          old_node.name = new_name;
          new_dir.ctime = new_dir.mtime = old_node.parent.ctime = old_node.parent.mtime = Date.now();
        },
  unlink(parent, name) {
          delete parent.contents[name];
          parent.ctime = parent.mtime = Date.now();
        },
  rmdir(parent, name) {
          var node = FS.lookupNode(parent, name);
          for (var i in node.contents) {
            throw new FS.ErrnoError(55);
          }
          delete parent.contents[name];
          parent.ctime = parent.mtime = Date.now();
        },
  readdir(node) {
          return ['.', '..', ...Object.keys(node.contents)];
        },
  symlink(parent, newname, oldpath) {
          var node = MEMFS.createNode(parent, newname, 0o777 | 40960, 0);
          node.link = oldpath;
          return node;
        },
  readlink(node) {
          if (!FS.isLink(node.mode)) {
            throw new FS.ErrnoError(28);
          }
          return node.link;
        },
  },
  stream_ops:{
  read(stream, buffer, offset, length, position) {
          var contents = stream.node.contents;
          if (position >= stream.node.usedBytes) return 0;
          var size = Math.min(stream.node.usedBytes - position, length);
          assert(size >= 0);
          if (size > 8 && contents.subarray) { // non-trivial, and typed array
            buffer.set(contents.subarray(position, position + size), offset);
          } else {
            for (var i = 0; i < size; i++) buffer[offset + i] = contents[position + i];
          }
          return size;
        },
  write(stream, buffer, offset, length, position, canOwn) {
          // The data buffer should be a typed array view
          assert(!(buffer instanceof ArrayBuffer));
          // If the buffer is located in main memory (HEAP), and if
          // memory can grow, we can't hold on to references of the
          // memory buffer, as they may get invalidated. That means we
          // need to do copy its contents.
          if (buffer.buffer === HEAP8.buffer) {
            canOwn = false;
          }
  
          if (!length) return 0;
          var node = stream.node;
          node.mtime = node.ctime = Date.now();
  
          if (buffer.subarray && (!node.contents || node.contents.subarray)) { // This write is from a typed array to a typed array?
            if (canOwn) {
              assert(position === 0, 'canOwn must imply no weird position inside the file');
              node.contents = buffer.subarray(offset, offset + length);
              node.usedBytes = length;
              return length;
            } else if (node.usedBytes === 0 && position === 0) { // If this is a simple first write to an empty file, do a fast set since we don't need to care about old data.
              node.contents = buffer.slice(offset, offset + length);
              node.usedBytes = length;
              return length;
            } else if (position + length <= node.usedBytes) { // Writing to an already allocated and used subrange of the file?
              node.contents.set(buffer.subarray(offset, offset + length), position);
              return length;
            }
          }
  
          // Appending to an existing file and we need to reallocate, or source data did not come as a typed array.
          MEMFS.expandFileStorage(node, position+length);
          if (node.contents.subarray && buffer.subarray) {
            // Use typed array write which is available.
            node.contents.set(buffer.subarray(offset, offset + length), position);
          } else {
            for (var i = 0; i < length; i++) {
             node.contents[position + i] = buffer[offset + i]; // Or fall back to manual write if not.
            }
          }
          node.usedBytes = Math.max(node.usedBytes, position + length);
          return length;
        },
  llseek(stream, offset, whence) {
          var position = offset;
          if (whence === 1) {
            position += stream.position;
          } else if (whence === 2) {
            if (FS.isFile(stream.node.mode)) {
              position += stream.node.usedBytes;
            }
          }
          if (position < 0) {
            throw new FS.ErrnoError(28);
          }
          return position;
        },
  mmap(stream, length, position, prot, flags) {
          if (!FS.isFile(stream.node.mode)) {
            throw new FS.ErrnoError(43);
          }
          var ptr;
          var allocated;
          var contents = stream.node.contents;
          // Only make a new copy when MAP_PRIVATE is specified.
          if (!(flags & 2) && contents && contents.buffer === HEAP8.buffer) {
            // We can't emulate MAP_SHARED when the file is not backed by the
            // buffer we're mapping to (e.g. the HEAP buffer).
            allocated = false;
            ptr = contents.byteOffset;
          } else {
            allocated = true;
            ptr = mmapAlloc(length);
            if (!ptr) {
              throw new FS.ErrnoError(48);
            }
            if (contents) {
              // Try to avoid unnecessary slices.
              if (position > 0 || position + length < contents.length) {
                if (contents.subarray) {
                  contents = contents.subarray(position, position + length);
                } else {
                  contents = Array.prototype.slice.call(contents, position, position + length);
                }
              }
              HEAP8.set(contents, ptr);
            }
          }
          return { ptr, allocated };
        },
  msync(stream, buffer, offset, length, mmapFlags) {
          MEMFS.stream_ops.write(stream, buffer, 0, length, offset, false);
          // should we check if bytesWritten and length are the same?
          return 0;
        },
  },
  };
  
  var asyncLoad = async (url) => {
      var arrayBuffer = await readAsync(url);
      assert(arrayBuffer, `Loading data file "${url}" failed (no arrayBuffer).`);
      return new Uint8Array(arrayBuffer);
    };
  
  
  var FS_createDataFile = (...args) => FS.createDataFile(...args);
  
  var getUniqueRunDependency = (id) => {
      var orig = id;
      while (1) {
        if (!runDependencyTracking[id]) return id;
        id = orig + Math.random();
      }
    };
  
  var preloadPlugins = [];
  var FS_handledByPreloadPlugin = (byteArray, fullname, finish, onerror) => {
      // Ensure plugins are ready.
      if (typeof Browser != 'undefined') Browser.init();
  
      var handled = false;
      preloadPlugins.forEach((plugin) => {
        if (handled) return;
        if (plugin['canHandle'](fullname)) {
          plugin['handle'](byteArray, fullname, finish, onerror);
          handled = true;
        }
      });
      return handled;
    };
  var FS_createPreloadedFile = (parent, name, url, canRead, canWrite, onload, onerror, dontCreateFile, canOwn, preFinish) => {
      // TODO we should allow people to just pass in a complete filename instead
      // of parent and name being that we just join them anyways
      var fullname = name ? PATH_FS.resolve(PATH.join2(parent, name)) : parent;
      var dep = getUniqueRunDependency(`cp ${fullname}`); // might have several active requests for the same fullname
      function processData(byteArray) {
        function finish(byteArray) {
          preFinish?.();
          if (!dontCreateFile) {
            FS_createDataFile(parent, name, byteArray, canRead, canWrite, canOwn);
          }
          onload?.();
          removeRunDependency(dep);
        }
        if (FS_handledByPreloadPlugin(byteArray, fullname, finish, () => {
          onerror?.();
          removeRunDependency(dep);
        })) {
          return;
        }
        finish(byteArray);
      }
      addRunDependency(dep);
      if (typeof url == 'string') {
        asyncLoad(url).then(processData, onerror);
      } else {
        processData(url);
      }
    };
  
  var FS_modeStringToFlags = (str) => {
      var flagModes = {
        'r': 0,
        'r+': 2,
        'w': 512 | 64 | 1,
        'w+': 512 | 64 | 2,
        'a': 1024 | 64 | 1,
        'a+': 1024 | 64 | 2,
      };
      var flags = flagModes[str];
      if (typeof flags == 'undefined') {
        throw new Error(`Unknown file open mode: ${str}`);
      }
      return flags;
    };
  
  var FS_getMode = (canRead, canWrite) => {
      var mode = 0;
      if (canRead) mode |= 292 | 73;
      if (canWrite) mode |= 146;
      return mode;
    };
  
  
  
  
    /**
     * Given a pointer 'ptr' to a null-terminated UTF8-encoded string in the
     * emscripten HEAP, returns a copy of that string as a Javascript String object.
     *
     * @param {number} ptr
     * @param {number=} maxBytesToRead - An optional length that specifies the
     *   maximum number of bytes to read. You can omit this parameter to scan the
     *   string until the first 0 byte. If maxBytesToRead is passed, and the string
     *   at [ptr, ptr+maxBytesToReadr[ contains a null byte in the middle, then the
     *   string will cut short at that byte index.
     * @param {boolean=} ignoreNul - If true, the function will not stop on a NUL character.
     * @return {string}
     */
  var UTF8ToString = (ptr, maxBytesToRead, ignoreNul) => {
      assert(typeof ptr == 'number', `UTF8ToString expects a number (got ${typeof ptr})`);
      return ptr ? UTF8ArrayToString(HEAPU8, ptr, maxBytesToRead, ignoreNul) : '';
    };
  
  var strError = (errno) => UTF8ToString(_strerror(errno));
  
  var ERRNO_CODES = {
      'EPERM': 63,
      'ENOENT': 44,
      'ESRCH': 71,
      'EINTR': 27,
      'EIO': 29,
      'ENXIO': 60,
      'E2BIG': 1,
      'ENOEXEC': 45,
      'EBADF': 8,
      'ECHILD': 12,
      'EAGAIN': 6,
      'EWOULDBLOCK': 6,
      'ENOMEM': 48,
      'EACCES': 2,
      'EFAULT': 21,
      'ENOTBLK': 105,
      'EBUSY': 10,
      'EEXIST': 20,
      'EXDEV': 75,
      'ENODEV': 43,
      'ENOTDIR': 54,
      'EISDIR': 31,
      'EINVAL': 28,
      'ENFILE': 41,
      'EMFILE': 33,
      'ENOTTY': 59,
      'ETXTBSY': 74,
      'EFBIG': 22,
      'ENOSPC': 51,
      'ESPIPE': 70,
      'EROFS': 69,
      'EMLINK': 34,
      'EPIPE': 64,
      'EDOM': 18,
      'ERANGE': 68,
      'ENOMSG': 49,
      'EIDRM': 24,
      'ECHRNG': 106,
      'EL2NSYNC': 156,
      'EL3HLT': 107,
      'EL3RST': 108,
      'ELNRNG': 109,
      'EUNATCH': 110,
      'ENOCSI': 111,
      'EL2HLT': 112,
      'EDEADLK': 16,
      'ENOLCK': 46,
      'EBADE': 113,
      'EBADR': 114,
      'EXFULL': 115,
      'ENOANO': 104,
      'EBADRQC': 103,
      'EBADSLT': 102,
      'EDEADLOCK': 16,
      'EBFONT': 101,
      'ENOSTR': 100,
      'ENODATA': 116,
      'ETIME': 117,
      'ENOSR': 118,
      'ENONET': 119,
      'ENOPKG': 120,
      'EREMOTE': 121,
      'ENOLINK': 47,
      'EADV': 122,
      'ESRMNT': 123,
      'ECOMM': 124,
      'EPROTO': 65,
      'EMULTIHOP': 36,
      'EDOTDOT': 125,
      'EBADMSG': 9,
      'ENOTUNIQ': 126,
      'EBADFD': 127,
      'EREMCHG': 128,
      'ELIBACC': 129,
      'ELIBBAD': 130,
      'ELIBSCN': 131,
      'ELIBMAX': 132,
      'ELIBEXEC': 133,
      'ENOSYS': 52,
      'ENOTEMPTY': 55,
      'ENAMETOOLONG': 37,
      'ELOOP': 32,
      'EOPNOTSUPP': 138,
      'EPFNOSUPPORT': 139,
      'ECONNRESET': 15,
      'ENOBUFS': 42,
      'EAFNOSUPPORT': 5,
      'EPROTOTYPE': 67,
      'ENOTSOCK': 57,
      'ENOPROTOOPT': 50,
      'ESHUTDOWN': 140,
      'ECONNREFUSED': 14,
      'EADDRINUSE': 3,
      'ECONNABORTED': 13,
      'ENETUNREACH': 40,
      'ENETDOWN': 38,
      'ETIMEDOUT': 73,
      'EHOSTDOWN': 142,
      'EHOSTUNREACH': 23,
      'EINPROGRESS': 26,
      'EALREADY': 7,
      'EDESTADDRREQ': 17,
      'EMSGSIZE': 35,
      'EPROTONOSUPPORT': 66,
      'ESOCKTNOSUPPORT': 137,
      'EADDRNOTAVAIL': 4,
      'ENETRESET': 39,
      'EISCONN': 30,
      'ENOTCONN': 53,
      'ETOOMANYREFS': 141,
      'EUSERS': 136,
      'EDQUOT': 19,
      'ESTALE': 72,
      'ENOTSUP': 138,
      'ENOMEDIUM': 148,
      'EILSEQ': 25,
      'EOVERFLOW': 61,
      'ECANCELED': 11,
      'ENOTRECOVERABLE': 56,
      'EOWNERDEAD': 62,
      'ESTRPIPE': 135,
    };
  var FS = {
  root:null,
  mounts:[],
  devices:{
  },
  streams:[],
  nextInode:1,
  nameTable:null,
  currentPath:"/",
  initialized:false,
  ignorePermissions:true,
  filesystems:null,
  syncFSRequests:0,
  readFiles:{
  },
  ErrnoError:class extends Error {
        name = 'ErrnoError';
        // We set the `name` property to be able to identify `FS.ErrnoError`
        // - the `name` is a standard ECMA-262 property of error objects. Kind of good to have it anyway.
        // - when using PROXYFS, an error can come from an underlying FS
        // as different FS objects have their own FS.ErrnoError each,
        // the test `err instanceof FS.ErrnoError` won't detect an error coming from another filesystem, causing bugs.
        // we'll use the reliable test `err.name == "ErrnoError"` instead
        constructor(errno) {
          super(runtimeInitialized ? strError(errno) : '');
          this.errno = errno;
          for (var key in ERRNO_CODES) {
            if (ERRNO_CODES[key] === errno) {
              this.code = key;
              break;
            }
          }
        }
      },
  FSStream:class {
        shared = {};
        get object() {
          return this.node;
        }
        set object(val) {
          this.node = val;
        }
        get isRead() {
          return (this.flags & 2097155) !== 1;
        }
        get isWrite() {
          return (this.flags & 2097155) !== 0;
        }
        get isAppend() {
          return (this.flags & 1024);
        }
        get flags() {
          return this.shared.flags;
        }
        set flags(val) {
          this.shared.flags = val;
        }
        get position() {
          return this.shared.position;
        }
        set position(val) {
          this.shared.position = val;
        }
      },
  FSNode:class {
        node_ops = {};
        stream_ops = {};
        readMode = 292 | 73;
        writeMode = 146;
        mounted = null;
        constructor(parent, name, mode, rdev) {
          if (!parent) {
            parent = this;  // root node sets parent to itself
          }
          this.parent = parent;
          this.mount = parent.mount;
          this.id = FS.nextInode++;
          this.name = name;
          this.mode = mode;
          this.rdev = rdev;
          this.atime = this.mtime = this.ctime = Date.now();
        }
        get read() {
          return (this.mode & this.readMode) === this.readMode;
        }
        set read(val) {
          val ? this.mode |= this.readMode : this.mode &= ~this.readMode;
        }
        get write() {
          return (this.mode & this.writeMode) === this.writeMode;
        }
        set write(val) {
          val ? this.mode |= this.writeMode : this.mode &= ~this.writeMode;
        }
        get isFolder() {
          return FS.isDir(this.mode);
        }
        get isDevice() {
          return FS.isChrdev(this.mode);
        }
      },
  lookupPath(path, opts = {}) {
        if (!path) {
          throw new FS.ErrnoError(44);
        }
        opts.follow_mount ??= true
  
        if (!PATH.isAbs(path)) {
          path = FS.cwd() + '/' + path;
        }
  
        // limit max consecutive symlinks to 40 (SYMLOOP_MAX).
        linkloop: for (var nlinks = 0; nlinks < 40; nlinks++) {
          // split the absolute path
          var parts = path.split('/').filter((p) => !!p);
  
          // start at the root
          var current = FS.root;
          var current_path = '/';
  
          for (var i = 0; i < parts.length; i++) {
            var islast = (i === parts.length-1);
            if (islast && opts.parent) {
              // stop resolving
              break;
            }
  
            if (parts[i] === '.') {
              continue;
            }
  
            if (parts[i] === '..') {
              current_path = PATH.dirname(current_path);
              if (FS.isRoot(current)) {
                path = current_path + '/' + parts.slice(i + 1).join('/');
                // We're making progress here, don't let many consecutive ..'s
                // lead to ELOOP
                nlinks--;
                continue linkloop;
              } else {
                current = current.parent;
              }
              continue;
            }
  
            current_path = PATH.join2(current_path, parts[i]);
            try {
              current = FS.lookupNode(current, parts[i]);
            } catch (e) {
              // if noent_okay is true, suppress a ENOENT in the last component
              // and return an object with an undefined node. This is needed for
              // resolving symlinks in the path when creating a file.
              if ((e?.errno === 44) && islast && opts.noent_okay) {
                return { path: current_path };
              }
              throw e;
            }
  
            // jump to the mount's root node if this is a mountpoint
            if (FS.isMountpoint(current) && (!islast || opts.follow_mount)) {
              current = current.mounted.root;
            }
  
            // by default, lookupPath will not follow a symlink if it is the final path component.
            // setting opts.follow = true will override this behavior.
            if (FS.isLink(current.mode) && (!islast || opts.follow)) {
              if (!current.node_ops.readlink) {
                throw new FS.ErrnoError(52);
              }
              var link = current.node_ops.readlink(current);
              if (!PATH.isAbs(link)) {
                link = PATH.dirname(current_path) + '/' + link;
              }
              path = link + '/' + parts.slice(i + 1).join('/');
              continue linkloop;
            }
          }
          return { path: current_path, node: current };
        }
        throw new FS.ErrnoError(32);
      },
  getPath(node) {
        var path;
        while (true) {
          if (FS.isRoot(node)) {
            var mount = node.mount.mountpoint;
            if (!path) return mount;
            return mount[mount.length-1] !== '/' ? `${mount}/${path}` : mount + path;
          }
          path = path ? `${node.name}/${path}` : node.name;
          node = node.parent;
        }
      },
  hashName(parentid, name) {
        var hash = 0;
  
        for (var i = 0; i < name.length; i++) {
          hash = ((hash << 5) - hash + name.charCodeAt(i)) | 0;
        }
        return ((parentid + hash) >>> 0) % FS.nameTable.length;
      },
  hashAddNode(node) {
        var hash = FS.hashName(node.parent.id, node.name);
        node.name_next = FS.nameTable[hash];
        FS.nameTable[hash] = node;
      },
  hashRemoveNode(node) {
        var hash = FS.hashName(node.parent.id, node.name);
        if (FS.nameTable[hash] === node) {
          FS.nameTable[hash] = node.name_next;
        } else {
          var current = FS.nameTable[hash];
          while (current) {
            if (current.name_next === node) {
              current.name_next = node.name_next;
              break;
            }
            current = current.name_next;
          }
        }
      },
  lookupNode(parent, name) {
        var errCode = FS.mayLookup(parent);
        if (errCode) {
          throw new FS.ErrnoError(errCode);
        }
        var hash = FS.hashName(parent.id, name);
        for (var node = FS.nameTable[hash]; node; node = node.name_next) {
          var nodeName = node.name;
          if (node.parent.id === parent.id && nodeName === name) {
            return node;
          }
        }
        // if we failed to find it in the cache, call into the VFS
        return FS.lookup(parent, name);
      },
  createNode(parent, name, mode, rdev) {
        assert(typeof parent == 'object')
        var node = new FS.FSNode(parent, name, mode, rdev);
  
        FS.hashAddNode(node);
  
        return node;
      },
  destroyNode(node) {
        FS.hashRemoveNode(node);
      },
  isRoot(node) {
        return node === node.parent;
      },
  isMountpoint(node) {
        return !!node.mounted;
      },
  isFile(mode) {
        return (mode & 61440) === 32768;
      },
  isDir(mode) {
        return (mode & 61440) === 16384;
      },
  isLink(mode) {
        return (mode & 61440) === 40960;
      },
  isChrdev(mode) {
        return (mode & 61440) === 8192;
      },
  isBlkdev(mode) {
        return (mode & 61440) === 24576;
      },
  isFIFO(mode) {
        return (mode & 61440) === 4096;
      },
  isSocket(mode) {
        return (mode & 49152) === 49152;
      },
  flagsToPermissionString(flag) {
        var perms = ['r', 'w', 'rw'][flag & 3];
        if ((flag & 512)) {
          perms += 'w';
        }
        return perms;
      },
  nodePermissions(node, perms) {
        if (FS.ignorePermissions) {
          return 0;
        }
        // return 0 if any user, group or owner bits are set.
        if (perms.includes('r') && !(node.mode & 292)) {
          return 2;
        } else if (perms.includes('w') && !(node.mode & 146)) {
          return 2;
        } else if (perms.includes('x') && !(node.mode & 73)) {
          return 2;
        }
        return 0;
      },
  mayLookup(dir) {
        if (!FS.isDir(dir.mode)) return 54;
        var errCode = FS.nodePermissions(dir, 'x');
        if (errCode) return errCode;
        if (!dir.node_ops.lookup) return 2;
        return 0;
      },
  mayCreate(dir, name) {
        if (!FS.isDir(dir.mode)) {
          return 54;
        }
        try {
          var node = FS.lookupNode(dir, name);
          return 20;
        } catch (e) {
        }
        return FS.nodePermissions(dir, 'wx');
      },
  mayDelete(dir, name, isdir) {
        var node;
        try {
          node = FS.lookupNode(dir, name);
        } catch (e) {
          return e.errno;
        }
        var errCode = FS.nodePermissions(dir, 'wx');
        if (errCode) {
          return errCode;
        }
        if (isdir) {
          if (!FS.isDir(node.mode)) {
            return 54;
          }
          if (FS.isRoot(node) || FS.getPath(node) === FS.cwd()) {
            return 10;
          }
        } else {
          if (FS.isDir(node.mode)) {
            return 31;
          }
        }
        return 0;
      },
  mayOpen(node, flags) {
        if (!node) {
          return 44;
        }
        if (FS.isLink(node.mode)) {
          return 32;
        } else if (FS.isDir(node.mode)) {
          if (FS.flagsToPermissionString(flags) !== 'r' // opening for write
              || (flags & (512 | 64))) { // TODO: check for O_SEARCH? (== search for dir only)
            return 31;
          }
        }
        return FS.nodePermissions(node, FS.flagsToPermissionString(flags));
      },
  checkOpExists(op, err) {
        if (!op) {
          throw new FS.ErrnoError(err);
        }
        return op;
      },
  MAX_OPEN_FDS:4096,
  nextfd() {
        for (var fd = 0; fd <= FS.MAX_OPEN_FDS; fd++) {
          if (!FS.streams[fd]) {
            return fd;
          }
        }
        throw new FS.ErrnoError(33);
      },
  getStreamChecked(fd) {
        var stream = FS.getStream(fd);
        if (!stream) {
          throw new FS.ErrnoError(8);
        }
        return stream;
      },
  getStream:(fd) => FS.streams[fd],
  createStream(stream, fd = -1) {
        assert(fd >= -1);
  
        // clone it, so we can return an instance of FSStream
        stream = Object.assign(new FS.FSStream(), stream);
        if (fd == -1) {
          fd = FS.nextfd();
        }
        stream.fd = fd;
        FS.streams[fd] = stream;
        return stream;
      },
  closeStream(fd) {
        FS.streams[fd] = null;
      },
  dupStream(origStream, fd = -1) {
        var stream = FS.createStream(origStream, fd);
        stream.stream_ops?.dup?.(stream);
        return stream;
      },
  doSetAttr(stream, node, attr) {
        var setattr = stream?.stream_ops.setattr;
        var arg = setattr ? stream : node;
        setattr ??= node.node_ops.setattr;
        FS.checkOpExists(setattr, 63)
        setattr(arg, attr);
      },
  chrdev_stream_ops:{
  open(stream) {
          var device = FS.getDevice(stream.node.rdev);
          // override node's stream ops with the device's
          stream.stream_ops = device.stream_ops;
          // forward the open call
          stream.stream_ops.open?.(stream);
        },
  llseek() {
          throw new FS.ErrnoError(70);
        },
  },
  major:(dev) => ((dev) >> 8),
  minor:(dev) => ((dev) & 0xff),
  makedev:(ma, mi) => ((ma) << 8 | (mi)),
  registerDevice(dev, ops) {
        FS.devices[dev] = { stream_ops: ops };
      },
  getDevice:(dev) => FS.devices[dev],
  getMounts(mount) {
        var mounts = [];
        var check = [mount];
  
        while (check.length) {
          var m = check.pop();
  
          mounts.push(m);
  
          check.push(...m.mounts);
        }
  
        return mounts;
      },
  syncfs(populate, callback) {
        if (typeof populate == 'function') {
          callback = populate;
          populate = false;
        }
  
        FS.syncFSRequests++;
  
        if (FS.syncFSRequests > 1) {
          err(`warning: ${FS.syncFSRequests} FS.syncfs operations in flight at once, probably just doing extra work`);
        }
  
        var mounts = FS.getMounts(FS.root.mount);
        var completed = 0;
  
        function doCallback(errCode) {
          assert(FS.syncFSRequests > 0);
          FS.syncFSRequests--;
          return callback(errCode);
        }
  
        function done(errCode) {
          if (errCode) {
            if (!done.errored) {
              done.errored = true;
              return doCallback(errCode);
            }
            return;
          }
          if (++completed >= mounts.length) {
            doCallback(null);
          }
        };
  
        // sync all mounts
        mounts.forEach((mount) => {
          if (!mount.type.syncfs) {
            return done(null);
          }
          mount.type.syncfs(mount, populate, done);
        });
      },
  mount(type, opts, mountpoint) {
        if (typeof type == 'string') {
          // The filesystem was not included, and instead we have an error
          // message stored in the variable.
          throw type;
        }
        var root = mountpoint === '/';
        var pseudo = !mountpoint;
        var node;
  
        if (root && FS.root) {
          throw new FS.ErrnoError(10);
        } else if (!root && !pseudo) {
          var lookup = FS.lookupPath(mountpoint, { follow_mount: false });
  
          mountpoint = lookup.path;  // use the absolute path
          node = lookup.node;
  
          if (FS.isMountpoint(node)) {
            throw new FS.ErrnoError(10);
          }
  
          if (!FS.isDir(node.mode)) {
            throw new FS.ErrnoError(54);
          }
        }
  
        var mount = {
          type,
          opts,
          mountpoint,
          mounts: []
        };
  
        // create a root node for the fs
        var mountRoot = type.mount(mount);
        mountRoot.mount = mount;
        mount.root = mountRoot;
  
        if (root) {
          FS.root = mountRoot;
        } else if (node) {
          // set as a mountpoint
          node.mounted = mount;
  
          // add the new mount to the current mount's children
          if (node.mount) {
            node.mount.mounts.push(mount);
          }
        }
  
        return mountRoot;
      },
  unmount(mountpoint) {
        var lookup = FS.lookupPath(mountpoint, { follow_mount: false });
  
        if (!FS.isMountpoint(lookup.node)) {
          throw new FS.ErrnoError(28);
        }
  
        // destroy the nodes for this mount, and all its child mounts
        var node = lookup.node;
        var mount = node.mounted;
        var mounts = FS.getMounts(mount);
  
        Object.keys(FS.nameTable).forEach((hash) => {
          var current = FS.nameTable[hash];
  
          while (current) {
            var next = current.name_next;
  
            if (mounts.includes(current.mount)) {
              FS.destroyNode(current);
            }
  
            current = next;
          }
        });
  
        // no longer a mountpoint
        node.mounted = null;
  
        // remove this mount from the child mounts
        var idx = node.mount.mounts.indexOf(mount);
        assert(idx !== -1);
        node.mount.mounts.splice(idx, 1);
      },
  lookup(parent, name) {
        return parent.node_ops.lookup(parent, name);
      },
  mknod(path, mode, dev) {
        var lookup = FS.lookupPath(path, { parent: true });
        var parent = lookup.node;
        var name = PATH.basename(path);
        if (!name) {
          throw new FS.ErrnoError(28);
        }
        if (name === '.' || name === '..') {
          throw new FS.ErrnoError(20);
        }
        var errCode = FS.mayCreate(parent, name);
        if (errCode) {
          throw new FS.ErrnoError(errCode);
        }
        if (!parent.node_ops.mknod) {
          throw new FS.ErrnoError(63);
        }
        return parent.node_ops.mknod(parent, name, mode, dev);
      },
  statfs(path) {
        return FS.statfsNode(FS.lookupPath(path, {follow: true}).node);
      },
  statfsStream(stream) {
        // We keep a separate statfsStream function because noderawfs overrides
        // it. In noderawfs, stream.node is sometimes null. Instead, we need to
        // look at stream.path.
        return FS.statfsNode(stream.node);
      },
  statfsNode(node) {
        // NOTE: None of the defaults here are true. We're just returning safe and
        //       sane values. Currently nodefs and rawfs replace these defaults,
        //       other file systems leave them alone.
        var rtn = {
          bsize: 4096,
          frsize: 4096,
          blocks: 1e6,
          bfree: 5e5,
          bavail: 5e5,
          files: FS.nextInode,
          ffree: FS.nextInode - 1,
          fsid: 42,
          flags: 2,
          namelen: 255,
        };
  
        if (node.node_ops.statfs) {
          Object.assign(rtn, node.node_ops.statfs(node.mount.opts.root));
        }
        return rtn;
      },
  create(path, mode = 0o666) {
        mode &= 4095;
        mode |= 32768;
        return FS.mknod(path, mode, 0);
      },
  mkdir(path, mode = 0o777) {
        mode &= 511 | 512;
        mode |= 16384;
        return FS.mknod(path, mode, 0);
      },
  mkdirTree(path, mode) {
        var dirs = path.split('/');
        var d = '';
        for (var dir of dirs) {
          if (!dir) continue;
          if (d || PATH.isAbs(path)) d += '/';
          d += dir;
          try {
            FS.mkdir(d, mode);
          } catch(e) {
            if (e.errno != 20) throw e;
          }
        }
      },
  mkdev(path, mode, dev) {
        if (typeof dev == 'undefined') {
          dev = mode;
          mode = 0o666;
        }
        mode |= 8192;
        return FS.mknod(path, mode, dev);
      },
  symlink(oldpath, newpath) {
        if (!PATH_FS.resolve(oldpath)) {
          throw new FS.ErrnoError(44);
        }
        var lookup = FS.lookupPath(newpath, { parent: true });
        var parent = lookup.node;
        if (!parent) {
          throw new FS.ErrnoError(44);
        }
        var newname = PATH.basename(newpath);
        var errCode = FS.mayCreate(parent, newname);
        if (errCode) {
          throw new FS.ErrnoError(errCode);
        }
        if (!parent.node_ops.symlink) {
          throw new FS.ErrnoError(63);
        }
        return parent.node_ops.symlink(parent, newname, oldpath);
      },
  rename(old_path, new_path) {
        var old_dirname = PATH.dirname(old_path);
        var new_dirname = PATH.dirname(new_path);
        var old_name = PATH.basename(old_path);
        var new_name = PATH.basename(new_path);
        // parents must exist
        var lookup, old_dir, new_dir;
  
        // let the errors from non existent directories percolate up
        lookup = FS.lookupPath(old_path, { parent: true });
        old_dir = lookup.node;
        lookup = FS.lookupPath(new_path, { parent: true });
        new_dir = lookup.node;
  
        if (!old_dir || !new_dir) throw new FS.ErrnoError(44);
        // need to be part of the same mount
        if (old_dir.mount !== new_dir.mount) {
          throw new FS.ErrnoError(75);
        }
        // source must exist
        var old_node = FS.lookupNode(old_dir, old_name);
        // old path should not be an ancestor of the new path
        var relative = PATH_FS.relative(old_path, new_dirname);
        if (relative.charAt(0) !== '.') {
          throw new FS.ErrnoError(28);
        }
        // new path should not be an ancestor of the old path
        relative = PATH_FS.relative(new_path, old_dirname);
        if (relative.charAt(0) !== '.') {
          throw new FS.ErrnoError(55);
        }
        // see if the new path already exists
        var new_node;
        try {
          new_node = FS.lookupNode(new_dir, new_name);
        } catch (e) {
          // not fatal
        }
        // early out if nothing needs to change
        if (old_node === new_node) {
          return;
        }
        // we'll need to delete the old entry
        var isdir = FS.isDir(old_node.mode);
        var errCode = FS.mayDelete(old_dir, old_name, isdir);
        if (errCode) {
          throw new FS.ErrnoError(errCode);
        }
        // need delete permissions if we'll be overwriting.
        // need create permissions if new doesn't already exist.
        errCode = new_node ?
          FS.mayDelete(new_dir, new_name, isdir) :
          FS.mayCreate(new_dir, new_name);
        if (errCode) {
          throw new FS.ErrnoError(errCode);
        }
        if (!old_dir.node_ops.rename) {
          throw new FS.ErrnoError(63);
        }
        if (FS.isMountpoint(old_node) || (new_node && FS.isMountpoint(new_node))) {
          throw new FS.ErrnoError(10);
        }
        // if we are going to change the parent, check write permissions
        if (new_dir !== old_dir) {
          errCode = FS.nodePermissions(old_dir, 'w');
          if (errCode) {
            throw new FS.ErrnoError(errCode);
          }
        }
        // remove the node from the lookup hash
        FS.hashRemoveNode(old_node);
        // do the underlying fs rename
        try {
          old_dir.node_ops.rename(old_node, new_dir, new_name);
          // update old node (we do this here to avoid each backend
          // needing to)
          old_node.parent = new_dir;
        } catch (e) {
          throw e;
        } finally {
          // add the node back to the hash (in case node_ops.rename
          // changed its name)
          FS.hashAddNode(old_node);
        }
      },
  rmdir(path) {
        var lookup = FS.lookupPath(path, { parent: true });
        var parent = lookup.node;
        var name = PATH.basename(path);
        var node = FS.lookupNode(parent, name);
        var errCode = FS.mayDelete(parent, name, true);
        if (errCode) {
          throw new FS.ErrnoError(errCode);
        }
        if (!parent.node_ops.rmdir) {
          throw new FS.ErrnoError(63);
        }
        if (FS.isMountpoint(node)) {
          throw new FS.ErrnoError(10);
        }
        parent.node_ops.rmdir(parent, name);
        FS.destroyNode(node);
      },
  readdir(path) {
        var lookup = FS.lookupPath(path, { follow: true });
        var node = lookup.node;
        var readdir = FS.checkOpExists(node.node_ops.readdir, 54);
        return readdir(node);
      },
  unlink(path) {
        var lookup = FS.lookupPath(path, { parent: true });
        var parent = lookup.node;
        if (!parent) {
          throw new FS.ErrnoError(44);
        }
        var name = PATH.basename(path);
        var node = FS.lookupNode(parent, name);
        var errCode = FS.mayDelete(parent, name, false);
        if (errCode) {
          // According to POSIX, we should map EISDIR to EPERM, but
          // we instead do what Linux does (and we must, as we use
          // the musl linux libc).
          throw new FS.ErrnoError(errCode);
        }
        if (!parent.node_ops.unlink) {
          throw new FS.ErrnoError(63);
        }
        if (FS.isMountpoint(node)) {
          throw new FS.ErrnoError(10);
        }
        parent.node_ops.unlink(parent, name);
        FS.destroyNode(node);
      },
  readlink(path) {
        var lookup = FS.lookupPath(path);
        var link = lookup.node;
        if (!link) {
          throw new FS.ErrnoError(44);
        }
        if (!link.node_ops.readlink) {
          throw new FS.ErrnoError(28);
        }
        return link.node_ops.readlink(link);
      },
  stat(path, dontFollow) {
        var lookup = FS.lookupPath(path, { follow: !dontFollow });
        var node = lookup.node;
        var getattr = FS.checkOpExists(node.node_ops.getattr, 63);
        return getattr(node);
      },
  fstat(fd) {
        var stream = FS.getStreamChecked(fd);
        var node = stream.node;
        var getattr = stream.stream_ops.getattr;
        var arg = getattr ? stream : node;
        getattr ??= node.node_ops.getattr;
        FS.checkOpExists(getattr, 63)
        return getattr(arg);
      },
  lstat(path) {
        return FS.stat(path, true);
      },
  doChmod(stream, node, mode, dontFollow) {
        FS.doSetAttr(stream, node, {
          mode: (mode & 4095) | (node.mode & ~4095),
          ctime: Date.now(),
          dontFollow
        });
      },
  chmod(path, mode, dontFollow) {
        var node;
        if (typeof path == 'string') {
          var lookup = FS.lookupPath(path, { follow: !dontFollow });
          node = lookup.node;
        } else {
          node = path;
        }
        FS.doChmod(null, node, mode, dontFollow);
      },
  lchmod(path, mode) {
        FS.chmod(path, mode, true);
      },
  fchmod(fd, mode) {
        var stream = FS.getStreamChecked(fd);
        FS.doChmod(stream, stream.node, mode, false);
      },
  doChown(stream, node, dontFollow) {
        FS.doSetAttr(stream, node, {
          timestamp: Date.now(),
          dontFollow
          // we ignore the uid / gid for now
        });
      },
  chown(path, uid, gid, dontFollow) {
        var node;
        if (typeof path == 'string') {
          var lookup = FS.lookupPath(path, { follow: !dontFollow });
          node = lookup.node;
        } else {
          node = path;
        }
        FS.doChown(null, node, dontFollow);
      },
  lchown(path, uid, gid) {
        FS.chown(path, uid, gid, true);
      },
  fchown(fd, uid, gid) {
        var stream = FS.getStreamChecked(fd);
        FS.doChown(stream, stream.node, false);
      },
  doTruncate(stream, node, len) {
        if (FS.isDir(node.mode)) {
          throw new FS.ErrnoError(31);
        }
        if (!FS.isFile(node.mode)) {
          throw new FS.ErrnoError(28);
        }
        var errCode = FS.nodePermissions(node, 'w');
        if (errCode) {
          throw new FS.ErrnoError(errCode);
        }
        FS.doSetAttr(stream, node, {
          size: len,
          timestamp: Date.now()
        });
      },
  truncate(path, len) {
        if (len < 0) {
          throw new FS.ErrnoError(28);
        }
        var node;
        if (typeof path == 'string') {
          var lookup = FS.lookupPath(path, { follow: true });
          node = lookup.node;
        } else {
          node = path;
        }
        FS.doTruncate(null, node, len);
      },
  ftruncate(fd, len) {
        var stream = FS.getStreamChecked(fd);
        if (len < 0 || (stream.flags & 2097155) === 0) {
          throw new FS.ErrnoError(28);
        }
        FS.doTruncate(stream, stream.node, len);
      },
  utime(path, atime, mtime) {
        var lookup = FS.lookupPath(path, { follow: true });
        var node = lookup.node;
        var setattr = FS.checkOpExists(node.node_ops.setattr, 63);
        setattr(node, {
          atime: atime,
          mtime: mtime
        });
      },
  open(path, flags, mode = 0o666) {
        if (path === "") {
          throw new FS.ErrnoError(44);
        }
        flags = typeof flags == 'string' ? FS_modeStringToFlags(flags) : flags;
        if ((flags & 64)) {
          mode = (mode & 4095) | 32768;
        } else {
          mode = 0;
        }
        var node;
        var isDirPath;
        if (typeof path == 'object') {
          node = path;
        } else {
          isDirPath = path.endsWith("/");
          // noent_okay makes it so that if the final component of the path
          // doesn't exist, lookupPath returns `node: undefined`. `path` will be
          // updated to point to the target of all symlinks.
          var lookup = FS.lookupPath(path, {
            follow: !(flags & 131072),
            noent_okay: true
          });
          node = lookup.node;
          path = lookup.path;
        }
        // perhaps we need to create the node
        var created = false;
        if ((flags & 64)) {
          if (node) {
            // if O_CREAT and O_EXCL are set, error out if the node already exists
            if ((flags & 128)) {
              throw new FS.ErrnoError(20);
            }
          } else if (isDirPath) {
            throw new FS.ErrnoError(31);
          } else {
            // node doesn't exist, try to create it
            // Ignore the permission bits here to ensure we can `open` this new
            // file below. We use chmod below the apply the permissions once the
            // file is open.
            node = FS.mknod(path, mode | 0o777, 0);
            created = true;
          }
        }
        if (!node) {
          throw new FS.ErrnoError(44);
        }
        // can't truncate a device
        if (FS.isChrdev(node.mode)) {
          flags &= ~512;
        }
        // if asked only for a directory, then this must be one
        if ((flags & 65536) && !FS.isDir(node.mode)) {
          throw new FS.ErrnoError(54);
        }
        // check permissions, if this is not a file we just created now (it is ok to
        // create and write to a file with read-only permissions; it is read-only
        // for later use)
        if (!created) {
          var errCode = FS.mayOpen(node, flags);
          if (errCode) {
            throw new FS.ErrnoError(errCode);
          }
        }
        // do truncation if necessary
        if ((flags & 512) && !created) {
          FS.truncate(node, 0);
        }
        // we've already handled these, don't pass down to the underlying vfs
        flags &= ~(128 | 512 | 131072);
  
        // register the stream with the filesystem
        var stream = FS.createStream({
          node,
          path: FS.getPath(node),  // we want the absolute path to the node
          flags,
          seekable: true,
          position: 0,
          stream_ops: node.stream_ops,
          // used by the file family libc calls (fopen, fwrite, ferror, etc.)
          ungotten: [],
          error: false
        });
        // call the new stream's open function
        if (stream.stream_ops.open) {
          stream.stream_ops.open(stream);
        }
        if (created) {
          FS.chmod(node, mode & 0o777);
        }
        if (Module['logReadFiles'] && !(flags & 1)) {
          if (!(path in FS.readFiles)) {
            FS.readFiles[path] = 1;
          }
        }
        return stream;
      },
  close(stream) {
        if (FS.isClosed(stream)) {
          throw new FS.ErrnoError(8);
        }
        if (stream.getdents) stream.getdents = null; // free readdir state
        try {
          if (stream.stream_ops.close) {
            stream.stream_ops.close(stream);
          }
        } catch (e) {
          throw e;
        } finally {
          FS.closeStream(stream.fd);
        }
        stream.fd = null;
      },
  isClosed(stream) {
        return stream.fd === null;
      },
  llseek(stream, offset, whence) {
        if (FS.isClosed(stream)) {
          throw new FS.ErrnoError(8);
        }
        if (!stream.seekable || !stream.stream_ops.llseek) {
          throw new FS.ErrnoError(70);
        }
        if (whence != 0 && whence != 1 && whence != 2) {
          throw new FS.ErrnoError(28);
        }
        stream.position = stream.stream_ops.llseek(stream, offset, whence);
        stream.ungotten = [];
        return stream.position;
      },
  read(stream, buffer, offset, length, position) {
        assert(offset >= 0);
        if (length < 0 || position < 0) {
          throw new FS.ErrnoError(28);
        }
        if (FS.isClosed(stream)) {
          throw new FS.ErrnoError(8);
        }
        if ((stream.flags & 2097155) === 1) {
          throw new FS.ErrnoError(8);
        }
        if (FS.isDir(stream.node.mode)) {
          throw new FS.ErrnoError(31);
        }
        if (!stream.stream_ops.read) {
          throw new FS.ErrnoError(28);
        }
        var seeking = typeof position != 'undefined';
        if (!seeking) {
          position = stream.position;
        } else if (!stream.seekable) {
          throw new FS.ErrnoError(70);
        }
        var bytesRead = stream.stream_ops.read(stream, buffer, offset, length, position);
        if (!seeking) stream.position += bytesRead;
        return bytesRead;
      },
  write(stream, buffer, offset, length, position, canOwn) {
        assert(offset >= 0);
        if (length < 0 || position < 0) {
          throw new FS.ErrnoError(28);
        }
        if (FS.isClosed(stream)) {
          throw new FS.ErrnoError(8);
        }
        if ((stream.flags & 2097155) === 0) {
          throw new FS.ErrnoError(8);
        }
        if (FS.isDir(stream.node.mode)) {
          throw new FS.ErrnoError(31);
        }
        if (!stream.stream_ops.write) {
          throw new FS.ErrnoError(28);
        }
        if (stream.seekable && stream.flags & 1024) {
          // seek to the end before writing in append mode
          FS.llseek(stream, 0, 2);
        }
        var seeking = typeof position != 'undefined';
        if (!seeking) {
          position = stream.position;
        } else if (!stream.seekable) {
          throw new FS.ErrnoError(70);
        }
        var bytesWritten = stream.stream_ops.write(stream, buffer, offset, length, position, canOwn);
        if (!seeking) stream.position += bytesWritten;
        return bytesWritten;
      },
  mmap(stream, length, position, prot, flags) {
        // User requests writing to file (prot & PROT_WRITE != 0).
        // Checking if we have permissions to write to the file unless
        // MAP_PRIVATE flag is set. According to POSIX spec it is possible
        // to write to file opened in read-only mode with MAP_PRIVATE flag,
        // as all modifications will be visible only in the memory of
        // the current process.
        if ((prot & 2) !== 0
            && (flags & 2) === 0
            && (stream.flags & 2097155) !== 2) {
          throw new FS.ErrnoError(2);
        }
        if ((stream.flags & 2097155) === 1) {
          throw new FS.ErrnoError(2);
        }
        if (!stream.stream_ops.mmap) {
          throw new FS.ErrnoError(43);
        }
        if (!length) {
          throw new FS.ErrnoError(28);
        }
        return stream.stream_ops.mmap(stream, length, position, prot, flags);
      },
  msync(stream, buffer, offset, length, mmapFlags) {
        assert(offset >= 0);
        if (!stream.stream_ops.msync) {
          return 0;
        }
        return stream.stream_ops.msync(stream, buffer, offset, length, mmapFlags);
      },
  ioctl(stream, cmd, arg) {
        if (!stream.stream_ops.ioctl) {
          throw new FS.ErrnoError(59);
        }
        return stream.stream_ops.ioctl(stream, cmd, arg);
      },
  readFile(path, opts = {}) {
        opts.flags = opts.flags || 0;
        opts.encoding = opts.encoding || 'binary';
        if (opts.encoding !== 'utf8' && opts.encoding !== 'binary') {
          throw new Error(`Invalid encoding type "${opts.encoding}"`);
        }
        var stream = FS.open(path, opts.flags);
        var stat = FS.stat(path);
        var length = stat.size;
        var buf = new Uint8Array(length);
        FS.read(stream, buf, 0, length, 0);
        if (opts.encoding === 'utf8') {
          buf = UTF8ArrayToString(buf);
        }
        FS.close(stream);
        return buf;
      },
  writeFile(path, data, opts = {}) {
        opts.flags = opts.flags || 577;
        var stream = FS.open(path, opts.flags, opts.mode);
        if (typeof data == 'string') {
          data = new Uint8Array(intArrayFromString(data, true));
        }
        if (ArrayBuffer.isView(data)) {
          FS.write(stream, data, 0, data.byteLength, undefined, opts.canOwn);
        } else {
          throw new Error('Unsupported data type');
        }
        FS.close(stream);
      },
  cwd:() => FS.currentPath,
  chdir(path) {
        var lookup = FS.lookupPath(path, { follow: true });
        if (lookup.node === null) {
          throw new FS.ErrnoError(44);
        }
        if (!FS.isDir(lookup.node.mode)) {
          throw new FS.ErrnoError(54);
        }
        var errCode = FS.nodePermissions(lookup.node, 'x');
        if (errCode) {
          throw new FS.ErrnoError(errCode);
        }
        FS.currentPath = lookup.path;
      },
  createDefaultDirectories() {
        FS.mkdir('/tmp');
        FS.mkdir('/home');
        FS.mkdir('/home/web_user');
      },
  createDefaultDevices() {
        // create /dev
        FS.mkdir('/dev');
        // setup /dev/null
        FS.registerDevice(FS.makedev(1, 3), {
          read: () => 0,
          write: (stream, buffer, offset, length, pos) => length,
          llseek: () => 0,
        });
        FS.mkdev('/dev/null', FS.makedev(1, 3));
        // setup /dev/tty and /dev/tty1
        // stderr needs to print output using err() rather than out()
        // so we register a second tty just for it.
        TTY.register(FS.makedev(5, 0), TTY.default_tty_ops);
        TTY.register(FS.makedev(6, 0), TTY.default_tty1_ops);
        FS.mkdev('/dev/tty', FS.makedev(5, 0));
        FS.mkdev('/dev/tty1', FS.makedev(6, 0));
        // setup /dev/[u]random
        // use a buffer to avoid overhead of individual crypto calls per byte
        var randomBuffer = new Uint8Array(1024), randomLeft = 0;
        var randomByte = () => {
          if (randomLeft === 0) {
            randomFill(randomBuffer);
            randomLeft = randomBuffer.byteLength;
          }
          return randomBuffer[--randomLeft];
        };
        FS.createDevice('/dev', 'random', randomByte);
        FS.createDevice('/dev', 'urandom', randomByte);
        // we're not going to emulate the actual shm device,
        // just create the tmp dirs that reside in it commonly
        FS.mkdir('/dev/shm');
        FS.mkdir('/dev/shm/tmp');
      },
  createSpecialDirectories() {
        // create /proc/self/fd which allows /proc/self/fd/6 => readlink gives the
        // name of the stream for fd 6 (see test_unistd_ttyname)
        FS.mkdir('/proc');
        var proc_self = FS.mkdir('/proc/self');
        FS.mkdir('/proc/self/fd');
        FS.mount({
          mount() {
            var node = FS.createNode(proc_self, 'fd', 16895, 73);
            node.stream_ops = {
              llseek: MEMFS.stream_ops.llseek,
            };
            node.node_ops = {
              lookup(parent, name) {
                var fd = +name;
                var stream = FS.getStreamChecked(fd);
                var ret = {
                  parent: null,
                  mount: { mountpoint: 'fake' },
                  node_ops: { readlink: () => stream.path },
                  id: fd + 1,
                };
                ret.parent = ret; // make it look like a simple root node
                return ret;
              },
              readdir() {
                return Array.from(FS.streams.entries())
                  .filter(([k, v]) => v)
                  .map(([k, v]) => k.toString());
              }
            };
            return node;
          }
        }, {}, '/proc/self/fd');
      },
  createStandardStreams(input, output, error) {
        // TODO deprecate the old functionality of a single
        // input / output callback and that utilizes FS.createDevice
        // and instead require a unique set of stream ops
  
        // by default, we symlink the standard streams to the
        // default tty devices. however, if the standard streams
        // have been overwritten we create a unique device for
        // them instead.
        if (input) {
          FS.createDevice('/dev', 'stdin', input);
        } else {
          FS.symlink('/dev/tty', '/dev/stdin');
        }
        if (output) {
          FS.createDevice('/dev', 'stdout', null, output);
        } else {
          FS.symlink('/dev/tty', '/dev/stdout');
        }
        if (error) {
          FS.createDevice('/dev', 'stderr', null, error);
        } else {
          FS.symlink('/dev/tty1', '/dev/stderr');
        }
  
        // open default streams for the stdin, stdout and stderr devices
        var stdin = FS.open('/dev/stdin', 0);
        var stdout = FS.open('/dev/stdout', 1);
        var stderr = FS.open('/dev/stderr', 1);
        assert(stdin.fd === 0, `invalid handle for stdin (${stdin.fd})`);
        assert(stdout.fd === 1, `invalid handle for stdout (${stdout.fd})`);
        assert(stderr.fd === 2, `invalid handle for stderr (${stderr.fd})`);
      },
  staticInit() {
        FS.nameTable = new Array(4096);
  
        FS.mount(MEMFS, {}, '/');
  
        FS.createDefaultDirectories();
        FS.createDefaultDevices();
        FS.createSpecialDirectories();
  
        FS.filesystems = {
          'MEMFS': MEMFS,
        };
      },
  init(input, output, error) {
        assert(!FS.initialized, 'FS.init was previously called. If you want to initialize later with custom parameters, remove any earlier calls (note that one is automatically added to the generated code)');
        FS.initialized = true;
  
        // Allow Module.stdin etc. to provide defaults, if none explicitly passed to us here
        input ??= Module['stdin'];
        output ??= Module['stdout'];
        error ??= Module['stderr'];
  
        FS.createStandardStreams(input, output, error);
      },
  quit() {
        FS.initialized = false;
        // force-flush all streams, so we get musl std streams printed out
        _fflush(0);
        // close all of our streams
        for (var stream of FS.streams) {
          if (stream) {
            FS.close(stream);
          }
        }
      },
  findObject(path, dontResolveLastLink) {
        var ret = FS.analyzePath(path, dontResolveLastLink);
        if (!ret.exists) {
          return null;
        }
        return ret.object;
      },
  analyzePath(path, dontResolveLastLink) {
        // operate from within the context of the symlink's target
        try {
          var lookup = FS.lookupPath(path, { follow: !dontResolveLastLink });
          path = lookup.path;
        } catch (e) {
        }
        var ret = {
          isRoot: false, exists: false, error: 0, name: null, path: null, object: null,
          parentExists: false, parentPath: null, parentObject: null
        };
        try {
          var lookup = FS.lookupPath(path, { parent: true });
          ret.parentExists = true;
          ret.parentPath = lookup.path;
          ret.parentObject = lookup.node;
          ret.name = PATH.basename(path);
          lookup = FS.lookupPath(path, { follow: !dontResolveLastLink });
          ret.exists = true;
          ret.path = lookup.path;
          ret.object = lookup.node;
          ret.name = lookup.node.name;
          ret.isRoot = lookup.path === '/';
        } catch (e) {
          ret.error = e.errno;
        };
        return ret;
      },
  createPath(parent, path, canRead, canWrite) {
        parent = typeof parent == 'string' ? parent : FS.getPath(parent);
        var parts = path.split('/').reverse();
        while (parts.length) {
          var part = parts.pop();
          if (!part) continue;
          var current = PATH.join2(parent, part);
          try {
            FS.mkdir(current);
          } catch (e) {
            if (e.errno != 20) throw e;
          }
          parent = current;
        }
        return current;
      },
  createFile(parent, name, properties, canRead, canWrite) {
        var path = PATH.join2(typeof parent == 'string' ? parent : FS.getPath(parent), name);
        var mode = FS_getMode(canRead, canWrite);
        return FS.create(path, mode);
      },
  createDataFile(parent, name, data, canRead, canWrite, canOwn) {
        var path = name;
        if (parent) {
          parent = typeof parent == 'string' ? parent : FS.getPath(parent);
          path = name ? PATH.join2(parent, name) : parent;
        }
        var mode = FS_getMode(canRead, canWrite);
        var node = FS.create(path, mode);
        if (data) {
          if (typeof data == 'string') {
            var arr = new Array(data.length);
            for (var i = 0, len = data.length; i < len; ++i) arr[i] = data.charCodeAt(i);
            data = arr;
          }
          // make sure we can write to the file
          FS.chmod(node, mode | 146);
          var stream = FS.open(node, 577);
          FS.write(stream, data, 0, data.length, 0, canOwn);
          FS.close(stream);
          FS.chmod(node, mode);
        }
      },
  createDevice(parent, name, input, output) {
        var path = PATH.join2(typeof parent == 'string' ? parent : FS.getPath(parent), name);
        var mode = FS_getMode(!!input, !!output);
        FS.createDevice.major ??= 64;
        var dev = FS.makedev(FS.createDevice.major++, 0);
        // Create a fake device that a set of stream ops to emulate
        // the old behavior.
        FS.registerDevice(dev, {
          open(stream) {
            stream.seekable = false;
          },
          close(stream) {
            // flush any pending line data
            if (output?.buffer?.length) {
              output(10);
            }
          },
          read(stream, buffer, offset, length, pos /* ignored */) {
            var bytesRead = 0;
            for (var i = 0; i < length; i++) {
              var result;
              try {
                result = input();
              } catch (e) {
                throw new FS.ErrnoError(29);
              }
              if (result === undefined && bytesRead === 0) {
                throw new FS.ErrnoError(6);
              }
              if (result === null || result === undefined) break;
              bytesRead++;
              buffer[offset+i] = result;
            }
            if (bytesRead) {
              stream.node.atime = Date.now();
            }
            return bytesRead;
          },
          write(stream, buffer, offset, length, pos) {
            for (var i = 0; i < length; i++) {
              try {
                output(buffer[offset+i]);
              } catch (e) {
                throw new FS.ErrnoError(29);
              }
            }
            if (length) {
              stream.node.mtime = stream.node.ctime = Date.now();
            }
            return i;
          }
        });
        return FS.mkdev(path, mode, dev);
      },
  forceLoadFile(obj) {
        if (obj.isDevice || obj.isFolder || obj.link || obj.contents) return true;
        if (typeof XMLHttpRequest != 'undefined') {
          throw new Error("Lazy loading should have been performed (contents set) in createLazyFile, but it was not. Lazy loading only works in web workers. Use --embed-file or --preload-file in emcc on the main thread.");
        } else { // Command-line.
          try {
            obj.contents = readBinary(obj.url);
            obj.usedBytes = obj.contents.length;
          } catch (e) {
            throw new FS.ErrnoError(29);
          }
        }
      },
  createLazyFile(parent, name, url, canRead, canWrite) {
        // Lazy chunked Uint8Array (implements get and length from Uint8Array).
        // Actual getting is abstracted away for eventual reuse.
        class LazyUint8Array {
          lengthKnown = false;
          chunks = []; // Loaded chunks. Index is the chunk number
          get(idx) {
            if (idx > this.length-1 || idx < 0) {
              return undefined;
            }
            var chunkOffset = idx % this.chunkSize;
            var chunkNum = (idx / this.chunkSize)|0;
            return this.getter(chunkNum)[chunkOffset];
          }
          setDataGetter(getter) {
            this.getter = getter;
          }
          cacheLength() {
            // Find length
            var xhr = new XMLHttpRequest();
            xhr.open('HEAD', url, false);
            xhr.send(null);
            if (!(xhr.status >= 200 && xhr.status < 300 || xhr.status === 304)) throw new Error("Couldn't load " + url + ". Status: " + xhr.status);
            var datalength = Number(xhr.getResponseHeader("Content-length"));
            var header;
            var hasByteServing = (header = xhr.getResponseHeader("Accept-Ranges")) && header === "bytes";
            var usesGzip = (header = xhr.getResponseHeader("Content-Encoding")) && header === "gzip";
  
            var chunkSize = 1024*1024; // Chunk size in bytes
  
            if (!hasByteServing) chunkSize = datalength;
  
            // Function to get a range from the remote URL.
            var doXHR = (from, to) => {
              if (from > to) throw new Error("invalid range (" + from + ", " + to + ") or no bytes requested!");
              if (to > datalength-1) throw new Error("only " + datalength + " bytes available! programmer error!");
  
              // TODO: Use mozResponseArrayBuffer, responseStream, etc. if available.
              var xhr = new XMLHttpRequest();
              xhr.open('GET', url, false);
              if (datalength !== chunkSize) xhr.setRequestHeader("Range", "bytes=" + from + "-" + to);
  
              // Some hints to the browser that we want binary data.
              xhr.responseType = 'arraybuffer';
              if (xhr.overrideMimeType) {
                xhr.overrideMimeType('text/plain; charset=x-user-defined');
              }
  
              xhr.send(null);
              if (!(xhr.status >= 200 && xhr.status < 300 || xhr.status === 304)) throw new Error("Couldn't load " + url + ". Status: " + xhr.status);
              if (xhr.response !== undefined) {
                return new Uint8Array(/** @type{Array<number>} */(xhr.response || []));
              }
              return intArrayFromString(xhr.responseText || '', true);
            };
            var lazyArray = this;
            lazyArray.setDataGetter((chunkNum) => {
              var start = chunkNum * chunkSize;
              var end = (chunkNum+1) * chunkSize - 1; // including this byte
              end = Math.min(end, datalength-1); // if datalength-1 is selected, this is the last block
              if (typeof lazyArray.chunks[chunkNum] == 'undefined') {
                lazyArray.chunks[chunkNum] = doXHR(start, end);
              }
              if (typeof lazyArray.chunks[chunkNum] == 'undefined') throw new Error('doXHR failed!');
              return lazyArray.chunks[chunkNum];
            });
  
            if (usesGzip || !datalength) {
              // if the server uses gzip or doesn't supply the length, we have to download the whole file to get the (uncompressed) length
              chunkSize = datalength = 1; // this will force getter(0)/doXHR do download the whole file
              datalength = this.getter(0).length;
              chunkSize = datalength;
              out("LazyFiles on gzip forces download of the whole file when length is accessed");
            }
  
            this._length = datalength;
            this._chunkSize = chunkSize;
            this.lengthKnown = true;
          }
          get length() {
            if (!this.lengthKnown) {
              this.cacheLength();
            }
            return this._length;
          }
          get chunkSize() {
            if (!this.lengthKnown) {
              this.cacheLength();
            }
            return this._chunkSize;
          }
        }
  
        if (typeof XMLHttpRequest != 'undefined') {
          if (!ENVIRONMENT_IS_WORKER) throw 'Cannot do synchronous binary XHRs outside webworkers in modern browsers. Use --embed-file or --preload-file in emcc';
          var lazyArray = new LazyUint8Array();
          var properties = { isDevice: false, contents: lazyArray };
        } else {
          var properties = { isDevice: false, url: url };
        }
  
        var node = FS.createFile(parent, name, properties, canRead, canWrite);
        // This is a total hack, but I want to get this lazy file code out of the
        // core of MEMFS. If we want to keep this lazy file concept I feel it should
        // be its own thin LAZYFS proxying calls to MEMFS.
        if (properties.contents) {
          node.contents = properties.contents;
        } else if (properties.url) {
          node.contents = null;
          node.url = properties.url;
        }
        // Add a function that defers querying the file size until it is asked the first time.
        Object.defineProperties(node, {
          usedBytes: {
            get: function() { return this.contents.length; }
          }
        });
        // override each stream op with one that tries to force load the lazy file first
        var stream_ops = {};
        var keys = Object.keys(node.stream_ops);
        keys.forEach((key) => {
          var fn = node.stream_ops[key];
          stream_ops[key] = (...args) => {
            FS.forceLoadFile(node);
            return fn(...args);
          };
        });
        function writeChunks(stream, buffer, offset, length, position) {
          var contents = stream.node.contents;
          if (position >= contents.length)
            return 0;
          var size = Math.min(contents.length - position, length);
          assert(size >= 0);
          if (contents.slice) { // normal array
            for (var i = 0; i < size; i++) {
              buffer[offset + i] = contents[position + i];
            }
          } else {
            for (var i = 0; i < size; i++) { // LazyUint8Array from sync binary XHR
              buffer[offset + i] = contents.get(position + i);
            }
          }
          return size;
        }
        // use a custom read function
        stream_ops.read = (stream, buffer, offset, length, position) => {
          FS.forceLoadFile(node);
          return writeChunks(stream, buffer, offset, length, position)
        };
        // use a custom mmap function
        stream_ops.mmap = (stream, length, position, prot, flags) => {
          FS.forceLoadFile(node);
          var ptr = mmapAlloc(length);
          if (!ptr) {
            throw new FS.ErrnoError(48);
          }
          writeChunks(stream, HEAP8, ptr, length, position);
          return { ptr, allocated: true };
        };
        node.stream_ops = stream_ops;
        return node;
      },
  absolutePath() {
        abort('FS.absolutePath has been removed; use PATH_FS.resolve instead');
      },
  createFolder() {
        abort('FS.createFolder has been removed; use FS.mkdir instead');
      },
  createLink() {
        abort('FS.createLink has been removed; use FS.symlink instead');
      },
  joinPath() {
        abort('FS.joinPath has been removed; use PATH.join instead');
      },
  mmapAlloc() {
        abort('FS.mmapAlloc has been replaced by the top level function mmapAlloc');
      },
  standardizePath() {
        abort('FS.standardizePath has been removed; use PATH.normalize instead');
      },
  };
  
  var SYSCALLS = {
  DEFAULT_POLLMASK:5,
  calculateAt(dirfd, path, allowEmpty) {
        if (PATH.isAbs(path)) {
          return path;
        }
        // relative path
        var dir;
        if (dirfd === -100) {
          dir = FS.cwd();
        } else {
          var dirstream = SYSCALLS.getStreamFromFD(dirfd);
          dir = dirstream.path;
        }
        if (path.length == 0) {
          if (!allowEmpty) {
            throw new FS.ErrnoError(44);;
          }
          return dir;
        }
        return dir + '/' + path;
      },
  writeStat(buf, stat) {
        HEAP32[((buf)>>2)] = stat.dev;
        HEAP32[(((buf)+(4))>>2)] = stat.mode;
        HEAPU32[(((buf)+(8))>>2)] = stat.nlink;
        HEAP32[(((buf)+(12))>>2)] = stat.uid;
        HEAP32[(((buf)+(16))>>2)] = stat.gid;
        HEAP32[(((buf)+(20))>>2)] = stat.rdev;
        HEAP64[(((buf)+(24))>>3)] = BigInt(stat.size);
        HEAP32[(((buf)+(32))>>2)] = 4096;
        HEAP32[(((buf)+(36))>>2)] = stat.blocks;
        var atime = stat.atime.getTime();
        var mtime = stat.mtime.getTime();
        var ctime = stat.ctime.getTime();
        HEAP64[(((buf)+(40))>>3)] = BigInt(Math.floor(atime / 1000));
        HEAPU32[(((buf)+(48))>>2)] = (atime % 1000) * 1000 * 1000;
        HEAP64[(((buf)+(56))>>3)] = BigInt(Math.floor(mtime / 1000));
        HEAPU32[(((buf)+(64))>>2)] = (mtime % 1000) * 1000 * 1000;
        HEAP64[(((buf)+(72))>>3)] = BigInt(Math.floor(ctime / 1000));
        HEAPU32[(((buf)+(80))>>2)] = (ctime % 1000) * 1000 * 1000;
        HEAP64[(((buf)+(88))>>3)] = BigInt(stat.ino);
        return 0;
      },
  writeStatFs(buf, stats) {
        HEAP32[(((buf)+(4))>>2)] = stats.bsize;
        HEAP32[(((buf)+(40))>>2)] = stats.bsize;
        HEAP32[(((buf)+(8))>>2)] = stats.blocks;
        HEAP32[(((buf)+(12))>>2)] = stats.bfree;
        HEAP32[(((buf)+(16))>>2)] = stats.bavail;
        HEAP32[(((buf)+(20))>>2)] = stats.files;
        HEAP32[(((buf)+(24))>>2)] = stats.ffree;
        HEAP32[(((buf)+(28))>>2)] = stats.fsid;
        HEAP32[(((buf)+(44))>>2)] = stats.flags;  // ST_NOSUID
        HEAP32[(((buf)+(36))>>2)] = stats.namelen;
      },
  doMsync(addr, stream, len, flags, offset) {
        if (!FS.isFile(stream.node.mode)) {
          throw new FS.ErrnoError(43);
        }
        if (flags & 2) {
          // MAP_PRIVATE calls need not to be synced back to underlying fs
          return 0;
        }
        var buffer = HEAPU8.slice(addr, addr + len);
        FS.msync(stream, buffer, offset, len, flags);
      },
  getStreamFromFD(fd) {
        var stream = FS.getStreamChecked(fd);
        return stream;
      },
  varargs:undefined,
  getStr(ptr) {
        var ret = UTF8ToString(ptr);
        return ret;
      },
  };
  function ___syscall_fcntl64(fd, cmd, varargs) {
  SYSCALLS.varargs = varargs;
  try {
  
      var stream = SYSCALLS.getStreamFromFD(fd);
      switch (cmd) {
        case 0: {
          var arg = syscallGetVarargI();
          if (arg < 0) {
            return -28;
          }
          while (FS.streams[arg]) {
            arg++;
          }
          var newStream;
          newStream = FS.dupStream(stream, arg);
          return newStream.fd;
        }
        case 1:
        case 2:
          return 0;  // FD_CLOEXEC makes no sense for a single process.
        case 3:
          return stream.flags;
        case 4: {
          var arg = syscallGetVarargI();
          stream.flags |= arg;
          return 0;
        }
        case 12: {
          var arg = syscallGetVarargP();
          var offset = 0;
          // We're always unlocked.
          HEAP16[(((arg)+(offset))>>1)] = 2;
          return 0;
        }
        case 13:
        case 14:
          // Pretend that the locking is successful. These are process-level locks,
          // and Emscripten programs are a single process. If we supported linking a
          // filesystem between programs, we'd need to do more here.
          // See https://github.com/emscripten-core/emscripten/issues/23697
          return 0;
      }
      return -28;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  function ___syscall_fdatasync(fd) {
  try {
  
      var stream = SYSCALLS.getStreamFromFD(fd);
      return 0; // we can't do anything synchronously; the in-memory FS is already synced to
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  function ___syscall_fstat64(fd, buf) {
  try {
  
      return SYSCALLS.writeStat(buf, FS.fstat(fd));
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  var INT53_MAX = 9007199254740992;
  
  var INT53_MIN = -9007199254740992;
  var bigintToI53Checked = (num) => (num < INT53_MIN || num > INT53_MAX) ? NaN : Number(num);
  function ___syscall_ftruncate64(fd, length) {
    length = bigintToI53Checked(length);
  
  
  try {
  
      if (isNaN(length)) return -61;
      FS.ftruncate(fd, length);
      return 0;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  ;
  }

  
  var stringToUTF8 = (str, outPtr, maxBytesToWrite) => {
      assert(typeof maxBytesToWrite == 'number', 'stringToUTF8(str, outPtr, maxBytesToWrite) is missing the third parameter that specifies the length of the output buffer!');
      return stringToUTF8Array(str, HEAPU8, outPtr, maxBytesToWrite);
    };
  function ___syscall_getcwd(buf, size) {
  try {
  
      if (size === 0) return -28;
      var cwd = FS.cwd();
      var cwdLengthInBytes = lengthBytesUTF8(cwd) + 1;
      if (size < cwdLengthInBytes) return -68;
      stringToUTF8(cwd, buf, size);
      return cwdLengthInBytes;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  
  function ___syscall_getdents64(fd, dirp, count) {
  try {
  
      var stream = SYSCALLS.getStreamFromFD(fd)
      stream.getdents ||= FS.readdir(stream.path);
  
      var struct_size = 280;
      var pos = 0;
      var off = FS.llseek(stream, 0, 1);
  
      var startIdx = Math.floor(off / struct_size);
      var endIdx = Math.min(stream.getdents.length, startIdx + Math.floor(count/struct_size))
      for (var idx = startIdx; idx < endIdx; idx++) {
        var id;
        var type;
        var name = stream.getdents[idx];
        if (name === '.') {
          id = stream.node.id;
          type = 4; // DT_DIR
        }
        else if (name === '..') {
          var lookup = FS.lookupPath(stream.path, { parent: true });
          id = lookup.node.id;
          type = 4; // DT_DIR
        }
        else {
          var child;
          try {
            child = FS.lookupNode(stream.node, name);
          } catch (e) {
            // If the entry is not a directory, file, or symlink, nodefs
            // lookupNode will raise EINVAL. Skip these and continue.
            if (e?.errno === 28) {
              continue;
            }
            throw e;
          }
          id = child.id;
          type = FS.isChrdev(child.mode) ? 2 :  // DT_CHR, character device.
                 FS.isDir(child.mode) ? 4 :     // DT_DIR, directory.
                 FS.isLink(child.mode) ? 10 :   // DT_LNK, symbolic link.
                 8;                             // DT_REG, regular file.
        }
        assert(id);
        HEAP64[((dirp + pos)>>3)] = BigInt(id);
        HEAP64[(((dirp + pos)+(8))>>3)] = BigInt((idx + 1) * struct_size);
        HEAP16[(((dirp + pos)+(16))>>1)] = 280;
        HEAP8[(dirp + pos)+(18)] = type;
        stringToUTF8(name, dirp + pos + 19, 256);
        pos += struct_size;
      }
      FS.llseek(stream, idx * struct_size, 0);
      return pos;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  
  function ___syscall_ioctl(fd, op, varargs) {
  SYSCALLS.varargs = varargs;
  try {
  
      var stream = SYSCALLS.getStreamFromFD(fd);
      switch (op) {
        case 21509: {
          if (!stream.tty) return -59;
          return 0;
        }
        case 21505: {
          if (!stream.tty) return -59;
          if (stream.tty.ops.ioctl_tcgets) {
            var termios = stream.tty.ops.ioctl_tcgets(stream);
            var argp = syscallGetVarargP();
            HEAP32[((argp)>>2)] = termios.c_iflag || 0;
            HEAP32[(((argp)+(4))>>2)] = termios.c_oflag || 0;
            HEAP32[(((argp)+(8))>>2)] = termios.c_cflag || 0;
            HEAP32[(((argp)+(12))>>2)] = termios.c_lflag || 0;
            for (var i = 0; i < 32; i++) {
              HEAP8[(argp + i)+(17)] = termios.c_cc[i] || 0;
            }
            return 0;
          }
          return 0;
        }
        case 21510:
        case 21511:
        case 21512: {
          if (!stream.tty) return -59;
          return 0; // no-op, not actually adjusting terminal settings
        }
        case 21506:
        case 21507:
        case 21508: {
          if (!stream.tty) return -59;
          if (stream.tty.ops.ioctl_tcsets) {
            var argp = syscallGetVarargP();
            var c_iflag = HEAP32[((argp)>>2)];
            var c_oflag = HEAP32[(((argp)+(4))>>2)];
            var c_cflag = HEAP32[(((argp)+(8))>>2)];
            var c_lflag = HEAP32[(((argp)+(12))>>2)];
            var c_cc = []
            for (var i = 0; i < 32; i++) {
              c_cc.push(HEAP8[(argp + i)+(17)]);
            }
            return stream.tty.ops.ioctl_tcsets(stream.tty, op, { c_iflag, c_oflag, c_cflag, c_lflag, c_cc });
          }
          return 0; // no-op, not actually adjusting terminal settings
        }
        case 21519: {
          if (!stream.tty) return -59;
          var argp = syscallGetVarargP();
          HEAP32[((argp)>>2)] = 0;
          return 0;
        }
        case 21520: {
          if (!stream.tty) return -59;
          return -28; // not supported
        }
        case 21531: {
          var argp = syscallGetVarargP();
          return FS.ioctl(stream, op, argp);
        }
        case 21523: {
          // TODO: in theory we should write to the winsize struct that gets
          // passed in, but for now musl doesn't read anything on it
          if (!stream.tty) return -59;
          if (stream.tty.ops.ioctl_tiocgwinsz) {
            var winsize = stream.tty.ops.ioctl_tiocgwinsz(stream.tty);
            var argp = syscallGetVarargP();
            HEAP16[((argp)>>1)] = winsize[0];
            HEAP16[(((argp)+(2))>>1)] = winsize[1];
          }
          return 0;
        }
        case 21524: {
          // TODO: technically, this ioctl call should change the window size.
          // but, since emscripten doesn't have any concept of a terminal window
          // yet, we'll just silently throw it away as we do TIOCGWINSZ
          if (!stream.tty) return -59;
          return 0;
        }
        case 21515: {
          if (!stream.tty) return -59;
          return 0;
        }
        default: return -28; // not supported
      }
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  function ___syscall_lstat64(path, buf) {
  try {
  
      path = SYSCALLS.getStr(path);
      return SYSCALLS.writeStat(buf, FS.lstat(path));
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  function ___syscall_mkdirat(dirfd, path, mode) {
  try {
  
      path = SYSCALLS.getStr(path);
      path = SYSCALLS.calculateAt(dirfd, path);
      FS.mkdir(path, mode, 0);
      return 0;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  function ___syscall_newfstatat(dirfd, path, buf, flags) {
  try {
  
      path = SYSCALLS.getStr(path);
      var nofollow = flags & 256;
      var allowEmpty = flags & 4096;
      flags = flags & (~6400);
      assert(!flags, `unknown flags in __syscall_newfstatat: ${flags}`);
      path = SYSCALLS.calculateAt(dirfd, path, allowEmpty);
      return SYSCALLS.writeStat(buf, nofollow ? FS.lstat(path) : FS.stat(path));
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  
  function ___syscall_openat(dirfd, path, flags, varargs) {
  SYSCALLS.varargs = varargs;
  try {
  
      path = SYSCALLS.getStr(path);
      path = SYSCALLS.calculateAt(dirfd, path);
      var mode = varargs ? syscallGetVarargI() : 0;
      return FS.open(path, flags, mode).fd;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  function ___syscall_renameat(olddirfd, oldpath, newdirfd, newpath) {
  try {
  
      oldpath = SYSCALLS.getStr(oldpath);
      newpath = SYSCALLS.getStr(newpath);
      oldpath = SYSCALLS.calculateAt(olddirfd, oldpath);
      newpath = SYSCALLS.calculateAt(newdirfd, newpath);
      FS.rename(oldpath, newpath);
      return 0;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  function ___syscall_rmdir(path) {
  try {
  
      path = SYSCALLS.getStr(path);
      FS.rmdir(path);
      return 0;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  function ___syscall_stat64(path, buf) {
  try {
  
      path = SYSCALLS.getStr(path);
      return SYSCALLS.writeStat(buf, FS.stat(path));
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  function ___syscall_unlinkat(dirfd, path, flags) {
  try {
  
      path = SYSCALLS.getStr(path);
      path = SYSCALLS.calculateAt(dirfd, path);
      if (!flags) {
        FS.unlink(path);
      } else if (flags === 512) {
        FS.rmdir(path);
      } else {
        return -28;
      }
      return 0;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  }

  var __abort_js = () =>
      abort('native code called abort()');

  var getExecutableName = () => thisProgram || './this.program';
  
  var __emscripten_get_progname = (str, len) => stringToUTF8(getExecutableName(), str, len);

  var runtimeKeepaliveCounter = 0;
  var __emscripten_runtime_keepalive_clear = () => {
      noExitRuntime = false;
      runtimeKeepaliveCounter = 0;
    };

  var __emscripten_throw_longjmp = () => {
      throw Infinity;
    };

  function __gmtime_js(time, tmPtr) {
    time = bigintToI53Checked(time);
  
  
      var date = new Date(time * 1000);
      HEAP32[((tmPtr)>>2)] = date.getUTCSeconds();
      HEAP32[(((tmPtr)+(4))>>2)] = date.getUTCMinutes();
      HEAP32[(((tmPtr)+(8))>>2)] = date.getUTCHours();
      HEAP32[(((tmPtr)+(12))>>2)] = date.getUTCDate();
      HEAP32[(((tmPtr)+(16))>>2)] = date.getUTCMonth();
      HEAP32[(((tmPtr)+(20))>>2)] = date.getUTCFullYear()-1900;
      HEAP32[(((tmPtr)+(24))>>2)] = date.getUTCDay();
      var start = Date.UTC(date.getUTCFullYear(), 0, 1, 0, 0, 0, 0);
      var yday = ((date.getTime() - start) / (1000 * 60 * 60 * 24))|0;
      HEAP32[(((tmPtr)+(28))>>2)] = yday;
    ;
  }

  var isLeapYear = (year) => year%4 === 0 && (year%100 !== 0 || year%400 === 0);
  
  var MONTH_DAYS_LEAP_CUMULATIVE = [0,31,60,91,121,152,182,213,244,274,305,335];
  
  var MONTH_DAYS_REGULAR_CUMULATIVE = [0,31,59,90,120,151,181,212,243,273,304,334];
  var ydayFromDate = (date) => {
      var leap = isLeapYear(date.getFullYear());
      var monthDaysCumulative = (leap ? MONTH_DAYS_LEAP_CUMULATIVE : MONTH_DAYS_REGULAR_CUMULATIVE);
      var yday = monthDaysCumulative[date.getMonth()] + date.getDate() - 1; // -1 since it's days since Jan 1
  
      return yday;
    };
  
  function __localtime_js(time, tmPtr) {
    time = bigintToI53Checked(time);
  
  
      var date = new Date(time*1000);
      HEAP32[((tmPtr)>>2)] = date.getSeconds();
      HEAP32[(((tmPtr)+(4))>>2)] = date.getMinutes();
      HEAP32[(((tmPtr)+(8))>>2)] = date.getHours();
      HEAP32[(((tmPtr)+(12))>>2)] = date.getDate();
      HEAP32[(((tmPtr)+(16))>>2)] = date.getMonth();
      HEAP32[(((tmPtr)+(20))>>2)] = date.getFullYear()-1900;
      HEAP32[(((tmPtr)+(24))>>2)] = date.getDay();
  
      var yday = ydayFromDate(date)|0;
      HEAP32[(((tmPtr)+(28))>>2)] = yday;
      HEAP32[(((tmPtr)+(36))>>2)] = -(date.getTimezoneOffset() * 60);
  
      // Attention: DST is in December in South, and some regions don't have DST at all.
      var start = new Date(date.getFullYear(), 0, 1);
      var summerOffset = new Date(date.getFullYear(), 6, 1).getTimezoneOffset();
      var winterOffset = start.getTimezoneOffset();
      var dst = (summerOffset != winterOffset && date.getTimezoneOffset() == Math.min(winterOffset, summerOffset))|0;
      HEAP32[(((tmPtr)+(32))>>2)] = dst;
    ;
  }

  
  
  
  
  
  function __mmap_js(len, prot, flags, fd, offset, allocated, addr) {
    offset = bigintToI53Checked(offset);
  
  
  try {
  
      // musl's mmap doesn't allow values over a certain limit
      // see OFF_MASK in mmap.c.
      assert(!isNaN(offset));
      var stream = SYSCALLS.getStreamFromFD(fd);
      var res = FS.mmap(stream, len, offset, prot, flags);
      var ptr = res.ptr;
      HEAP32[((allocated)>>2)] = res.allocated;
      HEAPU32[((addr)>>2)] = ptr;
      return 0;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  ;
  }

  
  function __munmap_js(addr, len, prot, flags, fd, offset) {
    offset = bigintToI53Checked(offset);
  
  
  try {
  
      var stream = SYSCALLS.getStreamFromFD(fd);
      if (prot & 2) {
        SYSCALLS.doMsync(addr, stream, len, flags, offset);
      }
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return -e.errno;
  }
  ;
  }

  var timers = {
  };
  
  var handleException = (e) => {
      // Certain exception types we do not treat as errors since they are used for
      // internal control flow.
      // 1. ExitStatus, which is thrown by exit()
      // 2. "unwind", which is thrown by emscripten_unwind_to_js_event_loop() and others
      //    that wish to return to JS event loop.
      if (e instanceof ExitStatus || e == 'unwind') {
        return EXITSTATUS;
      }
      checkStackCookie();
      if (e instanceof WebAssembly.RuntimeError) {
        if (_emscripten_stack_get_current() <= 0) {
          err('Stack overflow detected.  You can try increasing -sSTACK_SIZE (currently set to 65536)');
        }
      }
      quit_(1, e);
    };
  
  
  var keepRuntimeAlive = () => noExitRuntime || runtimeKeepaliveCounter > 0;
  var _proc_exit = (code) => {
      EXITSTATUS = code;
      if (!keepRuntimeAlive()) {
        Module['onExit']?.(code);
        ABORT = true;
      }
      quit_(code, new ExitStatus(code));
    };
  
  
  /** @suppress {duplicate } */
  /** @param {boolean|number=} implicit */
  var exitJS = (status, implicit) => {
      EXITSTATUS = status;
  
      checkUnflushedContent();
  
      // if exit() was called explicitly, warn the user if the runtime isn't actually being shut down
      if (keepRuntimeAlive() && !implicit) {
        var msg = `program exited (with status: ${status}), but keepRuntimeAlive() is set (counter=${runtimeKeepaliveCounter}) due to an async operation, so halting execution but not exiting the runtime or preventing further async execution (you can use emscripten_force_exit, if you want to force a true shutdown)`;
        err(msg);
      }
  
      _proc_exit(status);
    };
  var _exit = exitJS;
  
  
  var maybeExit = () => {
      if (!keepRuntimeAlive()) {
        try {
          _exit(EXITSTATUS);
        } catch (e) {
          handleException(e);
        }
      }
    };
  var callUserCallback = (func) => {
      if (ABORT) {
        err('user callback triggered after runtime exited or application aborted.  Ignoring.');
        return;
      }
      try {
        func();
        maybeExit();
      } catch (e) {
        handleException(e);
      }
    };
  
  
  var _emscripten_get_now = () => performance.now();
  var __setitimer_js = (which, timeout_ms) => {
      // First, clear any existing timer.
      if (timers[which]) {
        clearTimeout(timers[which].id);
        delete timers[which];
      }
  
      // A timeout of zero simply cancels the current timeout so we have nothing
      // more to do.
      if (!timeout_ms) return 0;
  
      var id = setTimeout(() => {
        assert(which in timers);
        delete timers[which];
        callUserCallback(() => __emscripten_timeout(which, _emscripten_get_now()));
      }, timeout_ms);
      timers[which] = { id, timeout_ms };
      return 0;
    };

  
  var __tzset_js = (timezone, daylight, std_name, dst_name) => {
      // TODO: Use (malleable) environment variables instead of system settings.
      var currentYear = new Date().getFullYear();
      var winter = new Date(currentYear, 0, 1);
      var summer = new Date(currentYear, 6, 1);
      var winterOffset = winter.getTimezoneOffset();
      var summerOffset = summer.getTimezoneOffset();
  
      // Local standard timezone offset. Local standard time is not adjusted for
      // daylight savings.  This code uses the fact that getTimezoneOffset returns
      // a greater value during Standard Time versus Daylight Saving Time (DST).
      // Thus it determines the expected output during Standard Time, and it
      // compares whether the output of the given date the same (Standard) or less
      // (DST).
      var stdTimezoneOffset = Math.max(winterOffset, summerOffset);
  
      // timezone is specified as seconds west of UTC ("The external variable
      // `timezone` shall be set to the difference, in seconds, between
      // Coordinated Universal Time (UTC) and local standard time."), the same
      // as returned by stdTimezoneOffset.
      // See http://pubs.opengroup.org/onlinepubs/009695399/functions/tzset.html
      HEAPU32[((timezone)>>2)] = stdTimezoneOffset * 60;
  
      HEAP32[((daylight)>>2)] = Number(winterOffset != summerOffset);
  
      var extractZone = (timezoneOffset) => {
        // Why inverse sign?
        // Read here https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/getTimezoneOffset
        var sign = timezoneOffset >= 0 ? "-" : "+";
  
        var absOffset = Math.abs(timezoneOffset)
        var hours = String(Math.floor(absOffset / 60)).padStart(2, "0");
        var minutes = String(absOffset % 60).padStart(2, "0");
  
        return `UTC${sign}${hours}${minutes}`;
      }
  
      var winterName = extractZone(winterOffset);
      var summerName = extractZone(summerOffset);
      assert(winterName);
      assert(summerName);
      assert(lengthBytesUTF8(winterName) <= 16, `timezone name truncated to fit in TZNAME_MAX (${winterName})`);
      assert(lengthBytesUTF8(summerName) <= 16, `timezone name truncated to fit in TZNAME_MAX (${summerName})`);
      if (summerOffset < winterOffset) {
        // Northern hemisphere
        stringToUTF8(winterName, std_name, 17);
        stringToUTF8(summerName, dst_name, 17);
      } else {
        stringToUTF8(winterName, dst_name, 17);
        stringToUTF8(summerName, std_name, 17);
      }
    };

  
  var _emscripten_date_now = () => Date.now();
  
  var nowIsMonotonic = 1;
  
  var checkWasiClock = (clock_id) => clock_id >= 0 && clock_id <= 3;
  
  function _clock_time_get(clk_id, ignored_precision, ptime) {
    ignored_precision = bigintToI53Checked(ignored_precision);
  
  
      if (!checkWasiClock(clk_id)) {
        return 28;
      }
      var now;
      // all wasi clocks but realtime are monotonic
      if (clk_id === 0) {
        now = _emscripten_date_now();
      } else if (nowIsMonotonic) {
        now = _emscripten_get_now();
      } else {
        return 52;
      }
      // "now" is in ms, and wasi times are in ns.
      var nsec = Math.round(now * 1000 * 1000);
      HEAP64[((ptime)>>3)] = BigInt(nsec);
      return 0;
    ;
  }

  var readEmAsmArgsArray = [];
  var readEmAsmArgs = (sigPtr, buf) => {
      // Nobody should have mutated _readEmAsmArgsArray underneath us to be something else than an array.
      assert(Array.isArray(readEmAsmArgsArray));
      // The input buffer is allocated on the stack, so it must be stack-aligned.
      assert(buf % 16 == 0);
      readEmAsmArgsArray.length = 0;
      var ch;
      // Most arguments are i32s, so shift the buffer pointer so it is a plain
      // index into HEAP32.
      while (ch = HEAPU8[sigPtr++]) {
        var chr = String.fromCharCode(ch);
        var validChars = ['d', 'f', 'i', 'p'];
        // In WASM_BIGINT mode we support passing i64 values as bigint.
        validChars.push('j');
        assert(validChars.includes(chr), `Invalid character ${ch}("${chr}") in readEmAsmArgs! Use only [${validChars}], and do not specify "v" for void return argument.`);
        // Floats are always passed as doubles, so all types except for 'i'
        // are 8 bytes and require alignment.
        var wide = (ch != 105);
        wide &= (ch != 112);
        buf += wide && (buf % 8) ? 4 : 0;
        readEmAsmArgsArray.push(
          // Special case for pointers under wasm64 or CAN_ADDRESS_2GB mode.
          ch == 112 ? HEAPU32[((buf)>>2)] :
          ch == 106 ? HEAP64[((buf)>>3)] :
          ch == 105 ?
            HEAP32[((buf)>>2)] :
            HEAPF64[((buf)>>3)]
        );
        buf += wide ? 8 : 4;
      }
      return readEmAsmArgsArray;
    };
  var runEmAsmFunction = (code, sigPtr, argbuf) => {
      var args = readEmAsmArgs(sigPtr, argbuf);
      assert(ASM_CONSTS.hasOwnProperty(code), `No EM_ASM constant found at address ${code}.  The loaded WebAssembly file is likely out of sync with the generated JavaScript.`);
      return ASM_CONSTS[code](...args);
    };
  var _emscripten_asm_const_int = (code, sigPtr, argbuf) => {
      return runEmAsmFunction(code, sigPtr, argbuf);
    };

  var runMainThreadEmAsm = (emAsmAddr, sigPtr, argbuf, sync) => {
      var args = readEmAsmArgs(sigPtr, argbuf);
      assert(ASM_CONSTS.hasOwnProperty(emAsmAddr), `No EM_ASM constant found at address ${emAsmAddr}.  The loaded WebAssembly file is likely out of sync with the generated JavaScript.`);
      return ASM_CONSTS[emAsmAddr](...args);
    };
  var _emscripten_asm_const_int_sync_on_main_thread = (emAsmAddr, sigPtr, argbuf) => runMainThreadEmAsm(emAsmAddr, sigPtr, argbuf, 1);

  var _emscripten_asm_const_ptr_sync_on_main_thread = (emAsmAddr, sigPtr, argbuf) => runMainThreadEmAsm(emAsmAddr, sigPtr, argbuf, 1);

  
  var _emscripten_set_main_loop_timing = (mode, value) => {
      MainLoop.timingMode = mode;
      MainLoop.timingValue = value;
  
      if (!MainLoop.func) {
        err('emscripten_set_main_loop_timing: Cannot set timing mode for main loop since a main loop does not exist! Call emscripten_set_main_loop first to set one up.');
        return 1; // Return non-zero on failure, can't set timing mode when there is no main loop.
      }
  
      if (!MainLoop.running) {
        
        MainLoop.running = true;
      }
      if (mode == 0) {
        MainLoop.scheduler = function MainLoop_scheduler_setTimeout() {
          var timeUntilNextTick = Math.max(0, MainLoop.tickStartTime + value - _emscripten_get_now())|0;
          setTimeout(MainLoop.runner, timeUntilNextTick); // doing this each time means that on exception, we stop
        };
        MainLoop.method = 'timeout';
      } else if (mode == 1) {
        MainLoop.scheduler = function MainLoop_scheduler_rAF() {
          MainLoop.requestAnimationFrame(MainLoop.runner);
        };
        MainLoop.method = 'rAF';
      } else if (mode == 2) {
        if (typeof MainLoop.setImmediate == 'undefined') {
          if (typeof setImmediate == 'undefined') {
            // Emulate setImmediate. (note: not a complete polyfill, we don't emulate clearImmediate() to keep code size to minimum, since not needed)
            var setImmediates = [];
            var emscriptenMainLoopMessageId = 'setimmediate';
            /** @param {Event} event */
            var MainLoop_setImmediate_messageHandler = (event) => {
              // When called in current thread or Worker, the main loop ID is structured slightly different to accommodate for --proxy-to-worker runtime listening to Worker events,
              // so check for both cases.
              if (event.data === emscriptenMainLoopMessageId || event.data.target === emscriptenMainLoopMessageId) {
                event.stopPropagation();
                setImmediates.shift()();
              }
            };
            addEventListener("message", MainLoop_setImmediate_messageHandler, true);
            MainLoop.setImmediate = /** @type{function(function(): ?, ...?): number} */((func) => {
              setImmediates.push(func);
              if (ENVIRONMENT_IS_WORKER) {
                Module['setImmediates'] ??= [];
                Module['setImmediates'].push(func);
                postMessage({target: emscriptenMainLoopMessageId}); // In --proxy-to-worker, route the message via proxyClient.js
              } else postMessage(emscriptenMainLoopMessageId, "*"); // On the main thread, can just send the message to itself.
            });
          } else {
            MainLoop.setImmediate = setImmediate;
          }
        }
        MainLoop.scheduler = function MainLoop_scheduler_setImmediate() {
          MainLoop.setImmediate(MainLoop.runner);
        };
        MainLoop.method = 'immediate';
      }
      return 0;
    };
  
  
  
    /**
     * @param {number=} arg
     * @param {boolean=} noSetTiming
     */
  var setMainLoop = (iterFunc, fps, simulateInfiniteLoop, arg, noSetTiming) => {
      assert(!MainLoop.func, 'emscripten_set_main_loop: there can only be one main loop function at once: call emscripten_cancel_main_loop to cancel the previous one before setting a new one with different parameters.');
      MainLoop.func = iterFunc;
      MainLoop.arg = arg;
  
      var thisMainLoopId = MainLoop.currentlyRunningMainloop;
      function checkIsRunning() {
        if (thisMainLoopId < MainLoop.currentlyRunningMainloop) {
          
          maybeExit();
          return false;
        }
        return true;
      }
  
      // We create the loop runner here but it is not actually running until
      // _emscripten_set_main_loop_timing is called (which might happen a
      // later time).  This member signifies that the current runner has not
      // yet been started so that we can call runtimeKeepalivePush when it
      // gets it timing set for the first time.
      MainLoop.running = false;
      MainLoop.runner = function MainLoop_runner() {
        if (ABORT) return;
        if (MainLoop.queue.length > 0) {
          var start = Date.now();
          var blocker = MainLoop.queue.shift();
          blocker.func(blocker.arg);
          if (MainLoop.remainingBlockers) {
            var remaining = MainLoop.remainingBlockers;
            var next = remaining%1 == 0 ? remaining-1 : Math.floor(remaining);
            if (blocker.counted) {
              MainLoop.remainingBlockers = next;
            } else {
              // not counted, but move the progress along a tiny bit
              next = next + 0.5; // do not steal all the next one's progress
              MainLoop.remainingBlockers = (8*remaining + next)/9;
            }
          }
          MainLoop.updateStatus();
  
          // catches pause/resume main loop from blocker execution
          if (!checkIsRunning()) return;
  
          setTimeout(MainLoop.runner, 0);
          return;
        }
  
        // catch pauses from non-main loop sources
        if (!checkIsRunning()) return;
  
        // Implement very basic swap interval control
        MainLoop.currentFrameNumber = MainLoop.currentFrameNumber + 1 | 0;
        if (MainLoop.timingMode == 1 && MainLoop.timingValue > 1 && MainLoop.currentFrameNumber % MainLoop.timingValue != 0) {
          // Not the scheduled time to render this frame - skip.
          MainLoop.scheduler();
          return;
        } else if (MainLoop.timingMode == 0) {
          MainLoop.tickStartTime = _emscripten_get_now();
        }
  
        if (MainLoop.method === 'timeout' && Module['ctx']) {
          warnOnce('Looks like you are rendering without using requestAnimationFrame for the main loop. You should use 0 for the frame rate in emscripten_set_main_loop in order to use requestAnimationFrame, as that can greatly improve your frame rates!');
          MainLoop.method = ''; // just warn once per call to set main loop
        }
  
        MainLoop.runIter(iterFunc);
  
        // catch pauses from the main loop itself
        if (!checkIsRunning()) return;
  
        MainLoop.scheduler();
      }
  
      if (!noSetTiming) {
        if (fps > 0) {
          _emscripten_set_main_loop_timing(0, 1000.0 / fps);
        } else {
          // Do rAF by rendering each frame (no decimating)
          _emscripten_set_main_loop_timing(1, 1);
        }
  
        MainLoop.scheduler();
      }
  
      if (simulateInfiniteLoop) {
        throw 'unwind';
      }
    };
  
  
  var MainLoop = {
  running:false,
  scheduler:null,
  method:"",
  currentlyRunningMainloop:0,
  func:null,
  arg:0,
  timingMode:0,
  timingValue:0,
  currentFrameNumber:0,
  queue:[],
  preMainLoop:[],
  postMainLoop:[],
  pause() {
        MainLoop.scheduler = null;
        // Incrementing this signals the previous main loop that it's now become old, and it must return.
        MainLoop.currentlyRunningMainloop++;
      },
  resume() {
        MainLoop.currentlyRunningMainloop++;
        var timingMode = MainLoop.timingMode;
        var timingValue = MainLoop.timingValue;
        var func = MainLoop.func;
        MainLoop.func = null;
        // do not set timing and call scheduler, we will do it on the next lines
        setMainLoop(func, 0, false, MainLoop.arg, true);
        _emscripten_set_main_loop_timing(timingMode, timingValue);
        MainLoop.scheduler();
      },
  updateStatus() {
        if (Module['setStatus']) {
          var message = Module['statusMessage'] || 'Please wait...';
          var remaining = MainLoop.remainingBlockers ?? 0;
          var expected = MainLoop.expectedBlockers ?? 0;
          if (remaining) {
            if (remaining < expected) {
              Module['setStatus'](`{message} ({expected - remaining}/{expected})`);
            } else {
              Module['setStatus'](message);
            }
          } else {
            Module['setStatus']('');
          }
        }
      },
  init() {
        Module['preMainLoop'] && MainLoop.preMainLoop.push(Module['preMainLoop']);
        Module['postMainLoop'] && MainLoop.postMainLoop.push(Module['postMainLoop']);
      },
  runIter(func) {
        if (ABORT) return;
        for (var pre of MainLoop.preMainLoop) {
          if (pre() === false) {
            return; // |return false| skips a frame
          }
        }
        callUserCallback(func);
        for (var post of MainLoop.postMainLoop) {
          post();
        }
        checkStackCookie();
      },
  nextRAF:0,
  fakeRequestAnimationFrame(func) {
        // try to keep 60fps between calls to here
        var now = Date.now();
        if (MainLoop.nextRAF === 0) {
          MainLoop.nextRAF = now + 1000/60;
        } else {
          while (now + 2 >= MainLoop.nextRAF) { // fudge a little, to avoid timer jitter causing us to do lots of delay:0
            MainLoop.nextRAF += 1000/60;
          }
        }
        var delay = Math.max(MainLoop.nextRAF - now, 0);
        setTimeout(func, delay);
      },
  requestAnimationFrame(func) {
        if (typeof requestAnimationFrame == 'function') {
          requestAnimationFrame(func);
          return;
        }
        var RAF = MainLoop.fakeRequestAnimationFrame;
        RAF(func);
      },
  };
  var _emscripten_cancel_main_loop = () => {
      MainLoop.pause();
      MainLoop.func = null;
    };

  var _emscripten_clear_timeout = clearTimeout;


  var _emscripten_err = (str) => err(UTF8ToString(str));

  var onExits = [];
  var addOnExit = (cb) => onExits.push(cb);
  var JSEvents = {
  memcpy(target, src, size) {
        HEAP8.set(HEAP8.subarray(src, src + size), target);
      },
  removeAllEventListeners() {
        while (JSEvents.eventHandlers.length) {
          JSEvents._removeHandler(JSEvents.eventHandlers.length - 1);
        }
        JSEvents.deferredCalls = [];
      },
  inEventHandler:0,
  deferredCalls:[],
  deferCall(targetFunction, precedence, argsList) {
        function arraysHaveEqualContent(arrA, arrB) {
          if (arrA.length != arrB.length) return false;
  
          for (var i in arrA) {
            if (arrA[i] != arrB[i]) return false;
          }
          return true;
        }
        // Test if the given call was already queued, and if so, don't add it again.
        for (var call of JSEvents.deferredCalls) {
          if (call.targetFunction == targetFunction && arraysHaveEqualContent(call.argsList, argsList)) {
            return;
          }
        }
        JSEvents.deferredCalls.push({
          targetFunction,
          precedence,
          argsList
        });
  
        JSEvents.deferredCalls.sort((x,y) => x.precedence < y.precedence);
      },
  removeDeferredCalls(targetFunction) {
        JSEvents.deferredCalls = JSEvents.deferredCalls.filter((call) => call.targetFunction != targetFunction);
      },
  canPerformEventHandlerRequests() {
        if (navigator.userActivation) {
          // Verify against transient activation status from UserActivation API
          // whether it is possible to perform a request here without needing to defer. See
          // https://developer.mozilla.org/en-US/docs/Web/Security/User_activation#transient_activation
          // and https://caniuse.com/mdn-api_useractivation
          // At the time of writing, Firefox does not support this API: https://bugzilla.mozilla.org/show_bug.cgi?id=1791079
          return navigator.userActivation.isActive;
        }
  
        return JSEvents.inEventHandler && JSEvents.currentEventHandler.allowsDeferredCalls;
      },
  runDeferredCalls() {
        if (!JSEvents.canPerformEventHandlerRequests()) {
          return;
        }
        var deferredCalls = JSEvents.deferredCalls;
        JSEvents.deferredCalls = [];
        for (var call of deferredCalls) {
          call.targetFunction(...call.argsList);
        }
      },
  eventHandlers:[],
  removeAllHandlersOnTarget:(target, eventTypeString) => {
        for (var i = 0; i < JSEvents.eventHandlers.length; ++i) {
          if (JSEvents.eventHandlers[i].target == target &&
            (!eventTypeString || eventTypeString == JSEvents.eventHandlers[i].eventTypeString)) {
             JSEvents._removeHandler(i--);
           }
        }
      },
  _removeHandler(i) {
        var h = JSEvents.eventHandlers[i];
        h.target.removeEventListener(h.eventTypeString, h.eventListenerFunc, h.useCapture);
        JSEvents.eventHandlers.splice(i, 1);
      },
  registerOrRemoveHandler(eventHandler) {
        if (!eventHandler.target) {
          err('registerOrRemoveHandler: the target element for event handler registration does not exist, when processing the following event handler registration:');
          console.dir(eventHandler);
          return -4;
        }
        if (eventHandler.callbackfunc) {
          eventHandler.eventListenerFunc = function(event) {
            // Increment nesting count for the event handler.
            ++JSEvents.inEventHandler;
            JSEvents.currentEventHandler = eventHandler;
            // Process any old deferred calls the user has placed.
            JSEvents.runDeferredCalls();
            // Process the actual event, calls back to user C code handler.
            eventHandler.handlerFunc(event);
            // Process any new deferred calls that were placed right now from this event handler.
            JSEvents.runDeferredCalls();
            // Out of event handler - restore nesting count.
            --JSEvents.inEventHandler;
          };
  
          eventHandler.target.addEventListener(eventHandler.eventTypeString,
                                               eventHandler.eventListenerFunc,
                                               eventHandler.useCapture);
          JSEvents.eventHandlers.push(eventHandler);
        } else {
          for (var i = 0; i < JSEvents.eventHandlers.length; ++i) {
            if (JSEvents.eventHandlers[i].target == eventHandler.target
             && JSEvents.eventHandlers[i].eventTypeString == eventHandler.eventTypeString) {
               JSEvents._removeHandler(i--);
             }
          }
        }
        return 0;
      },
  getNodeNameForTarget(target) {
        if (!target) return '';
        if (target == window) return '#window';
        if (target == screen) return '#screen';
        return target?.nodeName || '';
      },
  fullscreenEnabled() {
        return document.fullscreenEnabled
        // Safari 13.0.3 on macOS Catalina 10.15.1 still ships with prefixed webkitFullscreenEnabled.
        // TODO: If Safari at some point ships with unprefixed version, update the version check above.
        || document.webkitFullscreenEnabled
         ;
      },
  };
  
  /** @type {Object} */
  var specialHTMLTargets = [0, typeof document != 'undefined' ? document : 0, typeof window != 'undefined' ? window : 0];
  
  
  var maybeCStringToJsString = (cString) => {
      // "cString > 2" checks if the input is a number, and isn't of the special
      // values we accept here, EMSCRIPTEN_EVENT_TARGET_* (which map to 0, 1, 2).
      // In other words, if cString > 2 then it's a pointer to a valid place in
      // memory, and points to a C string.
      return cString > 2 ? UTF8ToString(cString) : cString;
    };
  
  /** @suppress {duplicate } */
  var findEventTarget = (target) => {
      target = maybeCStringToJsString(target);
      var domElement = specialHTMLTargets[target] || (typeof document != 'undefined' ? document.querySelector(target) : null);
      return domElement;
    };
  var findCanvasEventTarget = findEventTarget;
  var _emscripten_get_canvas_element_size = (target, width, height) => {
      var canvas = findCanvasEventTarget(target);
      if (!canvas) return -4;
      HEAP32[((width)>>2)] = canvas.width;
      HEAP32[((height)>>2)] = canvas.height;
    };
  
  
  
  
  
  var stackAlloc = (sz) => __emscripten_stack_alloc(sz);
  var stringToUTF8OnStack = (str) => {
      var size = lengthBytesUTF8(str) + 1;
      var ret = stackAlloc(size);
      stringToUTF8(str, ret, size);
      return ret;
    };
  var getCanvasElementSize = (target) => {
      var sp = stackSave();
      var w = stackAlloc(8);
      var h = w + 4;
  
      var targetInt = stringToUTF8OnStack(target.id);
      var ret = _emscripten_get_canvas_element_size(targetInt, w, h);
      var size = [HEAP32[((w)>>2)], HEAP32[((h)>>2)]];
      stackRestore(sp);
      return size;
    };
  
  var _emscripten_set_canvas_element_size = (target, width, height) => {
      var canvas = findCanvasEventTarget(target);
      if (!canvas) return -4;
      canvas.width = width;
      canvas.height = height;
      return 0;
    };
  
  
  
  var setCanvasElementSize = (target, width, height) => {
      if (!target.controlTransferredOffscreen) {
        target.width = width;
        target.height = height;
      } else {
        // This function is being called from high-level JavaScript code instead of asm.js/Wasm,
        // and it needs to synchronously proxy over to another thread, so marshal the string onto the heap to do the call.
        var sp = stackSave();
        var targetInt = stringToUTF8OnStack(target.id);
        _emscripten_set_canvas_element_size(targetInt, width, height);
        stackRestore(sp);
      }
    };
  
  var currentFullscreenStrategy = {
  };
  
  var registerRestoreOldStyle = (canvas) => {
      var canvasSize = getCanvasElementSize(canvas);
      var oldWidth = canvasSize[0];
      var oldHeight = canvasSize[1];
      var oldCssWidth = canvas.style.width;
      var oldCssHeight = canvas.style.height;
      var oldBackgroundColor = canvas.style.backgroundColor; // Chrome reads color from here.
      var oldDocumentBackgroundColor = document.body.style.backgroundColor; // IE11 reads color from here.
      // Firefox always has black background color.
      var oldPaddingLeft = canvas.style.paddingLeft; // Chrome, FF, Safari
      var oldPaddingRight = canvas.style.paddingRight;
      var oldPaddingTop = canvas.style.paddingTop;
      var oldPaddingBottom = canvas.style.paddingBottom;
      var oldMarginLeft = canvas.style.marginLeft; // IE11
      var oldMarginRight = canvas.style.marginRight;
      var oldMarginTop = canvas.style.marginTop;
      var oldMarginBottom = canvas.style.marginBottom;
      var oldDocumentBodyMargin = document.body.style.margin;
      var oldDocumentOverflow = document.documentElement.style.overflow; // Chrome, Firefox
      var oldDocumentScroll = document.body.scroll; // IE
      var oldImageRendering = canvas.style.imageRendering;
  
      function restoreOldStyle() {
        var fullscreenElement = document.fullscreenElement
          || document.webkitFullscreenElement
          ;
        if (!fullscreenElement) {
          document.removeEventListener('fullscreenchange', restoreOldStyle);
  
          // Unprefixed Fullscreen API shipped in Chromium 71 (https://bugs.chromium.org/p/chromium/issues/detail?id=383813)
          // As of Safari 13.0.3 on macOS Catalina 10.15.1 still ships with prefixed webkitfullscreenchange. TODO: revisit this check once Safari ships unprefixed version.
          document.removeEventListener('webkitfullscreenchange', restoreOldStyle);
  
          setCanvasElementSize(canvas, oldWidth, oldHeight);
  
          canvas.style.width = oldCssWidth;
          canvas.style.height = oldCssHeight;
          canvas.style.backgroundColor = oldBackgroundColor; // Chrome
          // IE11 hack: assigning 'undefined' or an empty string to document.body.style.backgroundColor has no effect, so first assign back the default color
          // before setting the undefined value. Setting undefined value is also important, or otherwise we would later treat that as something that the user
          // had explicitly set so subsequent fullscreen transitions would not set background color properly.
          if (!oldDocumentBackgroundColor) document.body.style.backgroundColor = 'white';
          document.body.style.backgroundColor = oldDocumentBackgroundColor; // IE11
          canvas.style.paddingLeft = oldPaddingLeft; // Chrome, FF, Safari
          canvas.style.paddingRight = oldPaddingRight;
          canvas.style.paddingTop = oldPaddingTop;
          canvas.style.paddingBottom = oldPaddingBottom;
          canvas.style.marginLeft = oldMarginLeft; // IE11
          canvas.style.marginRight = oldMarginRight;
          canvas.style.marginTop = oldMarginTop;
          canvas.style.marginBottom = oldMarginBottom;
          document.body.style.margin = oldDocumentBodyMargin;
          document.documentElement.style.overflow = oldDocumentOverflow; // Chrome, Firefox
          document.body.scroll = oldDocumentScroll; // IE
          canvas.style.imageRendering = oldImageRendering;
          if (canvas.GLctxObject) canvas.GLctxObject.GLctx.viewport(0, 0, oldWidth, oldHeight);
  
          if (currentFullscreenStrategy.canvasResizedCallback) {
            getWasmTableEntry(currentFullscreenStrategy.canvasResizedCallback)(37, 0, currentFullscreenStrategy.canvasResizedCallbackUserData);
          }
        }
      }
      document.addEventListener('fullscreenchange', restoreOldStyle);
      // Unprefixed Fullscreen API shipped in Chromium 71 (https://bugs.chromium.org/p/chromium/issues/detail?id=383813)
      // As of Safari 13.0.3 on macOS Catalina 10.15.1 still ships with prefixed webkitfullscreenchange. TODO: revisit this check once Safari ships unprefixed version.
      document.addEventListener('webkitfullscreenchange', restoreOldStyle);
      return restoreOldStyle;
    };
  
  
  var setLetterbox = (element, topBottom, leftRight) => {
      // Cannot use margin to specify letterboxes in FF or Chrome, since those ignore margins in fullscreen mode.
      element.style.paddingLeft = element.style.paddingRight = leftRight + 'px';
      element.style.paddingTop = element.style.paddingBottom = topBottom + 'px';
    };
  
  
  var getBoundingClientRect = (e) => specialHTMLTargets.indexOf(e) < 0 ? e.getBoundingClientRect() : {'left':0,'top':0};
  var JSEvents_resizeCanvasForFullscreen = (target, strategy) => {
      var restoreOldStyle = registerRestoreOldStyle(target);
      var cssWidth = strategy.softFullscreen ? innerWidth : screen.width;
      var cssHeight = strategy.softFullscreen ? innerHeight : screen.height;
      var rect = getBoundingClientRect(target);
      var windowedCssWidth = rect.width;
      var windowedCssHeight = rect.height;
      var canvasSize = getCanvasElementSize(target);
      var windowedRttWidth = canvasSize[0];
      var windowedRttHeight = canvasSize[1];
  
      if (strategy.scaleMode == 3) {
        setLetterbox(target, (cssHeight - windowedCssHeight) / 2, (cssWidth - windowedCssWidth) / 2);
        cssWidth = windowedCssWidth;
        cssHeight = windowedCssHeight;
      } else if (strategy.scaleMode == 2) {
        if (cssWidth*windowedRttHeight < windowedRttWidth*cssHeight) {
          var desiredCssHeight = windowedRttHeight * cssWidth / windowedRttWidth;
          setLetterbox(target, (cssHeight - desiredCssHeight) / 2, 0);
          cssHeight = desiredCssHeight;
        } else {
          var desiredCssWidth = windowedRttWidth * cssHeight / windowedRttHeight;
          setLetterbox(target, 0, (cssWidth - desiredCssWidth) / 2);
          cssWidth = desiredCssWidth;
        }
      }
  
      // If we are adding padding, must choose a background color or otherwise Chrome will give the
      // padding a default white color. Do it only if user has not customized their own background color.
      target.style.backgroundColor ||= 'black';
      // IE11 does the same, but requires the color to be set in the document body.
      document.body.style.backgroundColor ||= 'black'; // IE11
      // Firefox always shows black letterboxes independent of style color.
  
      target.style.width = cssWidth + 'px';
      target.style.height = cssHeight + 'px';
  
      if (strategy.filteringMode == 1) {
        target.style.imageRendering = 'optimizeSpeed';
        target.style.imageRendering = '-moz-crisp-edges';
        target.style.imageRendering = '-o-crisp-edges';
        target.style.imageRendering = '-webkit-optimize-contrast';
        target.style.imageRendering = 'optimize-contrast';
        target.style.imageRendering = 'crisp-edges';
        target.style.imageRendering = 'pixelated';
      }
  
      var dpiScale = (strategy.canvasResolutionScaleMode == 2) ? devicePixelRatio : 1;
      if (strategy.canvasResolutionScaleMode != 0) {
        var newWidth = (cssWidth * dpiScale)|0;
        var newHeight = (cssHeight * dpiScale)|0;
        setCanvasElementSize(target, newWidth, newHeight);
        if (target.GLctxObject) target.GLctxObject.GLctx.viewport(0, 0, newWidth, newHeight);
      }
      return restoreOldStyle;
    };
  
  var JSEvents_requestFullscreen = (target, strategy) => {
      // EMSCRIPTEN_FULLSCREEN_SCALE_DEFAULT + EMSCRIPTEN_FULLSCREEN_CANVAS_SCALE_NONE is a mode where no extra logic is performed to the DOM elements.
      if (strategy.scaleMode != 0 || strategy.canvasResolutionScaleMode != 0) {
        JSEvents_resizeCanvasForFullscreen(target, strategy);
      }
  
      if (target.requestFullscreen) {
        target.requestFullscreen();
      } else if (target.webkitRequestFullscreen) {
        target.webkitRequestFullscreen(Element.ALLOW_KEYBOARD_INPUT);
      } else {
        return JSEvents.fullscreenEnabled() ? -3 : -1;
      }
  
      currentFullscreenStrategy = strategy;
  
      if (strategy.canvasResizedCallback) {
        getWasmTableEntry(strategy.canvasResizedCallback)(37, 0, strategy.canvasResizedCallbackUserData);
      }
  
      return 0;
    };
  var _emscripten_exit_fullscreen = () => {
      if (!JSEvents.fullscreenEnabled()) return -1;
      // Make sure no queued up calls will fire after this.
      JSEvents.removeDeferredCalls(JSEvents_requestFullscreen);
  
      var d = specialHTMLTargets[1];
      if (d.exitFullscreen) {
        d.fullscreenElement && d.exitFullscreen();
      } else if (d.webkitExitFullscreen) {
        d.webkitFullscreenElement && d.webkitExitFullscreen();
      } else {
        return -1;
      }
  
      return 0;
    };

  
  var requestPointerLock = (target) => {
      if (target.requestPointerLock) {
        target.requestPointerLock();
      } else {
        // document.body is known to accept pointer lock, so use that to differentiate if the user passed a bad element,
        // or if the whole browser just doesn't support the feature.
        if (document.body.requestPointerLock) {
          return -3;
        }
        return -1;
      }
      return 0;
    };
  var _emscripten_exit_pointerlock = () => {
      // Make sure no queued up calls will fire after this.
      JSEvents.removeDeferredCalls(requestPointerLock);
      if (!document.exitPointerLock) return -1;
      document.exitPointerLock();
      return 0;
    };

  
  
  var _emscripten_force_exit = (status) => {
      warnOnce('emscripten_force_exit cannot actually shut down the runtime, as the build does not have EXIT_RUNTIME set');
      __emscripten_runtime_keepalive_clear();
      _exit(status);
    };

  var fillBatteryEventData = (eventStruct, e) => {
      HEAPF64[((eventStruct)>>3)] = e.chargingTime;
      HEAPF64[(((eventStruct)+(8))>>3)] = e.dischargingTime;
      HEAPF64[(((eventStruct)+(16))>>3)] = e.level;
      HEAP8[(eventStruct)+(24)] = e.charging;
    };
  
  var battery = () => navigator.battery || navigator.mozBattery || navigator.webkitBattery;
  var _emscripten_get_battery_status = (batteryState) => {
      if (!battery()) return -1;
      fillBatteryEventData(batteryState, battery());
      return 0;
    };

  var _emscripten_get_device_pixel_ratio = () => {
      return (typeof devicePixelRatio == 'number' && devicePixelRatio) || 1.0;
    };

  
  var _emscripten_get_element_css_size = (target, width, height) => {
      target = findEventTarget(target);
      if (!target) return -4;
  
      var rect = getBoundingClientRect(target);
      HEAPF64[((width)>>3)] = rect.width;
      HEAPF64[((height)>>3)] = rect.height;
  
      return 0;
    };

  
  var fillGamepadEventData = (eventStruct, e) => {
      HEAPF64[((eventStruct)>>3)] = e.timestamp;
      for (var i = 0; i < e.axes.length; ++i) {
        HEAPF64[(((eventStruct+i*8)+(16))>>3)] = e.axes[i];
      }
      for (var i = 0; i < e.buttons.length; ++i) {
        if (typeof e.buttons[i] == 'object') {
          HEAPF64[(((eventStruct+i*8)+(528))>>3)] = e.buttons[i].value;
        } else {
          HEAPF64[(((eventStruct+i*8)+(528))>>3)] = e.buttons[i];
        }
      }
      for (var i = 0; i < e.buttons.length; ++i) {
        if (typeof e.buttons[i] == 'object') {
          HEAP8[(eventStruct+i)+(1040)] = e.buttons[i].pressed;
        } else {
          // Assigning a boolean to HEAP32, that's ok, but Closure would like to warn about it:
          /** @suppress {checkTypes} */
          HEAP8[(eventStruct+i)+(1040)] = e.buttons[i] == 1;
        }
      }
      HEAP8[(eventStruct)+(1104)] = e.connected;
      HEAP32[(((eventStruct)+(1108))>>2)] = e.index;
      HEAP32[(((eventStruct)+(8))>>2)] = e.axes.length;
      HEAP32[(((eventStruct)+(12))>>2)] = e.buttons.length;
      stringToUTF8(e.id, eventStruct + 1112, 64);
      stringToUTF8(e.mapping, eventStruct + 1176, 64);
    };
  var _emscripten_get_gamepad_status = (index, gamepadState) => {
      if (!JSEvents.lastGamepadState) throw 'emscripten_get_gamepad_status() can only be called after having first called emscripten_sample_gamepad_data() and that function has returned EMSCRIPTEN_RESULT_SUCCESS!';
      // INVALID_PARAM is returned on a Gamepad index that never was there.
      if (index < 0 || index >= JSEvents.lastGamepadState.length) return -5;
  
      // NO_DATA is returned on a Gamepad index that was removed.
      // For previously disconnected gamepads there should be an empty slot (null/undefined/false) at the index.
      // This is because gamepads must keep their original position in the array.
      // For example, removing the first of two gamepads produces [null/undefined/false, gamepad].
      if (!JSEvents.lastGamepadState[index]) return -7;
  
      fillGamepadEventData(gamepadState, JSEvents.lastGamepadState[index]);
      return 0;
    };

  var getHeapMax = () =>
      // Stay one Wasm page short of 4GB: while e.g. Chrome is able to allocate
      // full 4GB Wasm memories, the size will wrap back to 0 bytes in Wasm side
      // for any code that deals with heap sizes, which would require special
      // casing all heap size related code to treat 0 specially.
      2147483648;
  var _emscripten_get_heap_max = () => getHeapMax();

  var _emscripten_get_main_loop_timing = (mode, value) => {
      if (mode) HEAP32[((mode)>>2)] = MainLoop.timingMode;
      if (value) HEAP32[((value)>>2)] = MainLoop.timingValue;
    };


  var _emscripten_get_num_gamepads = () => {
      if (!JSEvents.lastGamepadState) throw 'emscripten_get_num_gamepads() can only be called after having first called emscripten_sample_gamepad_data() and that function has returned EMSCRIPTEN_RESULT_SUCCESS!';
      // N.B. Do not call emscripten_get_num_gamepads() unless having first called emscripten_sample_gamepad_data(), and that has returned EMSCRIPTEN_RESULT_SUCCESS.
      // Otherwise the following line will throw an exception.
      return JSEvents.lastGamepadState.length;
    };

  
  /** @param {number=} timeout */
  var safeSetTimeout = (func, timeout) => {
      
      return setTimeout(() => {
        
        callUserCallback(func);
      }, timeout);
    };
  
  
  
  var Browser = {
  useWebGL:false,
  isFullscreen:false,
  pointerLock:false,
  moduleContextCreatedCallbacks:[],
  workers:[],
  preloadedImages:{
  },
  preloadedAudios:{
  },
  getCanvas:() => Module['canvas'],
  init() {
        if (Browser.initted) return;
        Browser.initted = true;
  
        // Support for plugins that can process preloaded files. You can add more of these to
        // your app by creating and appending to preloadPlugins.
        //
        // Each plugin is asked if it can handle a file based on the file's name. If it can,
        // it is given the file's raw data. When it is done, it calls a callback with the file's
        // (possibly modified) data. For example, a plugin might decompress a file, or it
        // might create some side data structure for use later (like an Image element, etc.).
  
        var imagePlugin = {};
        imagePlugin['canHandle'] = function imagePlugin_canHandle(name) {
          return !Module['noImageDecoding'] && /\.(jpg|jpeg|png|bmp|webp)$/i.test(name);
        };
        imagePlugin['handle'] = function imagePlugin_handle(byteArray, name, onload, onerror) {
          var b = new Blob([byteArray], { type: Browser.getMimetype(name) });
          if (b.size !== byteArray.length) { // Safari bug #118630
            // Safari's Blob can only take an ArrayBuffer
            b = new Blob([(new Uint8Array(byteArray)).buffer], { type: Browser.getMimetype(name) });
          }
          var url = URL.createObjectURL(b);
          assert(typeof url == 'string', 'createObjectURL must return a url as a string');
          var img = new Image();
          img.onload = () => {
            assert(img.complete, `Image ${name} could not be decoded`);
            var canvas = /** @type {!HTMLCanvasElement} */ (document.createElement('canvas'));
            canvas.width = img.width;
            canvas.height = img.height;
            var ctx = canvas.getContext('2d');
            ctx.drawImage(img, 0, 0);
            Browser.preloadedImages[name] = canvas;
            URL.revokeObjectURL(url);
            onload?.(byteArray);
          };
          img.onerror = (event) => {
            err(`Image ${url} could not be decoded`);
            onerror?.();
          };
          img.src = url;
        };
        preloadPlugins.push(imagePlugin);
  
        var audioPlugin = {};
        audioPlugin['canHandle'] = function audioPlugin_canHandle(name) {
          return !Module['noAudioDecoding'] && name.slice(-4) in { '.ogg': 1, '.wav': 1, '.mp3': 1 };
        };
        audioPlugin['handle'] = function audioPlugin_handle(byteArray, name, onload, onerror) {
          var done = false;
          function finish(audio) {
            if (done) return;
            done = true;
            Browser.preloadedAudios[name] = audio;
            onload?.(byteArray);
          }
          function fail() {
            if (done) return;
            done = true;
            Browser.preloadedAudios[name] = new Audio(); // empty shim
            onerror?.();
          }
          var b = new Blob([byteArray], { type: Browser.getMimetype(name) });
          var url = URL.createObjectURL(b); // XXX we never revoke this!
          assert(typeof url == 'string', 'createObjectURL must return a url as a string');
          var audio = new Audio();
          audio.addEventListener('canplaythrough', () => finish(audio), false); // use addEventListener due to chromium bug 124926
          audio.onerror = function audio_onerror(event) {
            if (done) return;
            err(`warning: browser could not fully decode audio ${name}, trying slower base64 approach`);
            function encode64(data) {
              var BASE = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/';
              var PAD = '=';
              var ret = '';
              var leftchar = 0;
              var leftbits = 0;
              for (var i = 0; i < data.length; i++) {
                leftchar = (leftchar << 8) | data[i];
                leftbits += 8;
                while (leftbits >= 6) {
                  var curr = (leftchar >> (leftbits-6)) & 0x3f;
                  leftbits -= 6;
                  ret += BASE[curr];
                }
              }
              if (leftbits == 2) {
                ret += BASE[(leftchar&3) << 4];
                ret += PAD + PAD;
              } else if (leftbits == 4) {
                ret += BASE[(leftchar&0xf) << 2];
                ret += PAD;
              }
              return ret;
            }
            audio.src = 'data:audio/x-' + name.slice(-3) + ';base64,' + encode64(byteArray);
            finish(audio); // we don't wait for confirmation this worked - but it's worth trying
          };
          audio.src = url;
          // workaround for chrome bug 124926 - we do not always get oncanplaythrough or onerror
          safeSetTimeout(() => {
            finish(audio); // try to use it even though it is not necessarily ready to play
          }, 10000);
        };
        preloadPlugins.push(audioPlugin);
  
        // Canvas event setup
  
        function pointerLockChange() {
          var canvas = Browser.getCanvas();
          Browser.pointerLock = document['pointerLockElement'] === canvas ||
                                document['mozPointerLockElement'] === canvas ||
                                document['webkitPointerLockElement'] === canvas ||
                                document['msPointerLockElement'] === canvas;
        }
        var canvas = Browser.getCanvas();
        if (canvas) {
          // forced aspect ratio can be enabled by defining 'forcedAspectRatio' on Module
          // Module['forcedAspectRatio'] = 4 / 3;
  
          canvas.requestPointerLock = canvas['requestPointerLock'] ||
                                      canvas['mozRequestPointerLock'] ||
                                      canvas['webkitRequestPointerLock'] ||
                                      canvas['msRequestPointerLock'] ||
                                      (() => {});
          canvas.exitPointerLock = document['exitPointerLock'] ||
                                   document['mozExitPointerLock'] ||
                                   document['webkitExitPointerLock'] ||
                                   document['msExitPointerLock'] ||
                                   (() => {}); // no-op if function does not exist
          canvas.exitPointerLock = canvas.exitPointerLock.bind(document);
  
          document.addEventListener('pointerlockchange', pointerLockChange, false);
          document.addEventListener('mozpointerlockchange', pointerLockChange, false);
          document.addEventListener('webkitpointerlockchange', pointerLockChange, false);
          document.addEventListener('mspointerlockchange', pointerLockChange, false);
  
          if (Module['elementPointerLock']) {
            canvas.addEventListener("click", (ev) => {
              if (!Browser.pointerLock && Browser.getCanvas().requestPointerLock) {
                Browser.getCanvas().requestPointerLock();
                ev.preventDefault();
              }
            }, false);
          }
        }
      },
  createContext(/** @type {HTMLCanvasElement} */ canvas, useWebGL, setInModule, webGLContextAttributes) {
        if (useWebGL && Module['ctx'] && canvas == Browser.getCanvas()) return Module['ctx']; // no need to recreate GL context if it's already been created for this canvas.
  
        var ctx;
        var contextHandle;
        if (useWebGL) {
          // For GLES2/desktop GL compatibility, adjust a few defaults to be different to WebGL defaults, so that they align better with the desktop defaults.
          var contextAttributes = {
            antialias: false,
            alpha: false,
            majorVersion: 1,
          };
  
          if (webGLContextAttributes) {
            for (var attribute in webGLContextAttributes) {
              contextAttributes[attribute] = webGLContextAttributes[attribute];
            }
          }
  
          // This check of existence of GL is here to satisfy Closure compiler, which yells if variable GL is referenced below but GL object is not
          // actually compiled in because application is not doing any GL operations. TODO: Ideally if GL is not being used, this function
          // Browser.createContext() should not even be emitted.
          if (typeof GL != 'undefined') {
            contextHandle = GL.createContext(canvas, contextAttributes);
            if (contextHandle) {
              ctx = GL.getContext(contextHandle).GLctx;
            }
          }
        } else {
          ctx = canvas.getContext('2d');
        }
  
        if (!ctx) return null;
  
        if (setInModule) {
          if (!useWebGL) assert(typeof GLctx == 'undefined', 'cannot set in module if GLctx is used, but we are a non-GL context that would replace it');
          Module['ctx'] = ctx;
          if (useWebGL) GL.makeContextCurrent(contextHandle);
          Browser.useWebGL = useWebGL;
          Browser.moduleContextCreatedCallbacks.forEach((callback) => callback());
          Browser.init();
        }
        return ctx;
      },
  fullscreenHandlersInstalled:false,
  lockPointer:undefined,
  resizeCanvas:undefined,
  requestFullscreen(lockPointer, resizeCanvas) {
        Browser.lockPointer = lockPointer;
        Browser.resizeCanvas = resizeCanvas;
        if (typeof Browser.lockPointer == 'undefined') Browser.lockPointer = true;
        if (typeof Browser.resizeCanvas == 'undefined') Browser.resizeCanvas = false;
  
        var canvas = Browser.getCanvas();
        function fullscreenChange() {
          Browser.isFullscreen = false;
          var canvasContainer = canvas.parentNode;
          if ((document['fullscreenElement'] || document['mozFullScreenElement'] ||
               document['msFullscreenElement'] || document['webkitFullscreenElement'] ||
               document['webkitCurrentFullScreenElement']) === canvasContainer) {
            canvas.exitFullscreen = Browser.exitFullscreen;
            if (Browser.lockPointer) canvas.requestPointerLock();
            Browser.isFullscreen = true;
            if (Browser.resizeCanvas) {
              Browser.setFullscreenCanvasSize();
            } else {
              Browser.updateCanvasDimensions(canvas);
            }
          } else {
            // remove the full screen specific parent of the canvas again to restore the HTML structure from before going full screen
            canvasContainer.parentNode.insertBefore(canvas, canvasContainer);
            canvasContainer.parentNode.removeChild(canvasContainer);
  
            if (Browser.resizeCanvas) {
              Browser.setWindowedCanvasSize();
            } else {
              Browser.updateCanvasDimensions(canvas);
            }
          }
          Module['onFullScreen']?.(Browser.isFullscreen);
          Module['onFullscreen']?.(Browser.isFullscreen);
        }
  
        if (!Browser.fullscreenHandlersInstalled) {
          Browser.fullscreenHandlersInstalled = true;
          document.addEventListener('fullscreenchange', fullscreenChange, false);
          document.addEventListener('mozfullscreenchange', fullscreenChange, false);
          document.addEventListener('webkitfullscreenchange', fullscreenChange, false);
          document.addEventListener('MSFullscreenChange', fullscreenChange, false);
        }
  
        // create a new parent to ensure the canvas has no siblings. this allows browsers to optimize full screen performance when its parent is the full screen root
        var canvasContainer = document.createElement("div");
        canvas.parentNode.insertBefore(canvasContainer, canvas);
        canvasContainer.appendChild(canvas);
  
        // use parent of canvas as full screen root to allow aspect ratio correction (Firefox stretches the root to screen size)
        canvasContainer.requestFullscreen = canvasContainer['requestFullscreen'] ||
                                            canvasContainer['mozRequestFullScreen'] ||
                                            canvasContainer['msRequestFullscreen'] ||
                                           (canvasContainer['webkitRequestFullscreen'] ? () => canvasContainer['webkitRequestFullscreen'](Element['ALLOW_KEYBOARD_INPUT']) : null) ||
                                           (canvasContainer['webkitRequestFullScreen'] ? () => canvasContainer['webkitRequestFullScreen'](Element['ALLOW_KEYBOARD_INPUT']) : null);
  
        canvasContainer.requestFullscreen();
      },
  requestFullScreen() {
        abort('Module.requestFullScreen has been replaced by Module.requestFullscreen (without a capital S)');
      },
  exitFullscreen() {
        // This is workaround for chrome. Trying to exit from fullscreen
        // not in fullscreen state will cause "TypeError: Document not active"
        // in chrome. See https://github.com/emscripten-core/emscripten/pull/8236
        if (!Browser.isFullscreen) {
          return false;
        }
  
        var CFS = document['exitFullscreen'] ||
                  document['cancelFullScreen'] ||
                  document['mozCancelFullScreen'] ||
                  document['msExitFullscreen'] ||
                  document['webkitCancelFullScreen'] ||
            (() => {});
        CFS.apply(document, []);
        return true;
      },
  safeSetTimeout(func, timeout) {
        // Legacy function, this is used by the SDL2 port so we need to keep it
        // around at least until that is updated.
        // See https://github.com/libsdl-org/SDL/pull/6304
        return safeSetTimeout(func, timeout);
      },
  getMimetype(name) {
        return {
          'jpg': 'image/jpeg',
          'jpeg': 'image/jpeg',
          'png': 'image/png',
          'bmp': 'image/bmp',
          'ogg': 'audio/ogg',
          'wav': 'audio/wav',
          'mp3': 'audio/mpeg'
        }[name.slice(name.lastIndexOf('.')+1)];
      },
  getUserMedia(func) {
        window.getUserMedia ||= navigator['getUserMedia'] ||
                                navigator['mozGetUserMedia'];
        window.getUserMedia(func);
      },
  getMovementX(event) {
        return event['movementX'] ||
               event['mozMovementX'] ||
               event['webkitMovementX'] ||
               0;
      },
  getMovementY(event) {
        return event['movementY'] ||
               event['mozMovementY'] ||
               event['webkitMovementY'] ||
               0;
      },
  getMouseWheelDelta(event) {
        var delta = 0;
        switch (event.type) {
          case 'DOMMouseScroll':
            // 3 lines make up a step
            delta = event.detail / 3;
            break;
          case 'mousewheel':
            // 120 units make up a step
            delta = event.wheelDelta / 120;
            break;
          case 'wheel':
            delta = event.deltaY
            switch (event.deltaMode) {
              case 0:
                // DOM_DELTA_PIXEL: 100 pixels make up a step
                delta /= 100;
                break;
              case 1:
                // DOM_DELTA_LINE: 3 lines make up a step
                delta /= 3;
                break;
              case 2:
                // DOM_DELTA_PAGE: A page makes up 80 steps
                delta *= 80;
                break;
              default:
                throw 'unrecognized mouse wheel delta mode: ' + event.deltaMode;
            }
            break;
          default:
            throw 'unrecognized mouse wheel event: ' + event.type;
        }
        return delta;
      },
  mouseX:0,
  mouseY:0,
  mouseMovementX:0,
  mouseMovementY:0,
  touches:{
  },
  lastTouches:{
  },
  calculateMouseCoords(pageX, pageY) {
        // Calculate the movement based on the changes
        // in the coordinates.
        var canvas = Browser.getCanvas();
        var rect = canvas.getBoundingClientRect();
  
        // Neither .scrollX or .pageXOffset are defined in a spec, but
        // we prefer .scrollX because it is currently in a spec draft.
        // (see: http://www.w3.org/TR/2013/WD-cssom-view-20131217/)
        var scrollX = ((typeof window.scrollX != 'undefined') ? window.scrollX : window.pageXOffset);
        var scrollY = ((typeof window.scrollY != 'undefined') ? window.scrollY : window.pageYOffset);
        // If this assert lands, it's likely because the browser doesn't support scrollX or pageXOffset
        // and we have no viable fallback.
        assert((typeof scrollX != 'undefined') && (typeof scrollY != 'undefined'), 'Unable to retrieve scroll position, mouse positions likely broken.');
        var adjustedX = pageX - (scrollX + rect.left);
        var adjustedY = pageY - (scrollY + rect.top);
  
        // the canvas might be CSS-scaled compared to its backbuffer;
        // SDL-using content will want mouse coordinates in terms
        // of backbuffer units.
        adjustedX = adjustedX * (canvas.width / rect.width);
        adjustedY = adjustedY * (canvas.height / rect.height);
  
        return { x: adjustedX, y: adjustedY };
      },
  setMouseCoords(pageX, pageY) {
        const {x, y} = Browser.calculateMouseCoords(pageX, pageY);
        Browser.mouseMovementX = x - Browser.mouseX;
        Browser.mouseMovementY = y - Browser.mouseY;
        Browser.mouseX = x;
        Browser.mouseY = y;
      },
  calculateMouseEvent(event) { // event should be mousemove, mousedown or mouseup
        if (Browser.pointerLock) {
          // When the pointer is locked, calculate the coordinates
          // based on the movement of the mouse.
          // Workaround for Firefox bug 764498
          if (event.type != 'mousemove' &&
              ('mozMovementX' in event)) {
            Browser.mouseMovementX = Browser.mouseMovementY = 0;
          } else {
            Browser.mouseMovementX = Browser.getMovementX(event);
            Browser.mouseMovementY = Browser.getMovementY(event);
          }
  
          // add the mouse delta to the current absolute mouse position
          Browser.mouseX += Browser.mouseMovementX;
          Browser.mouseY += Browser.mouseMovementY;
        } else {
          if (event.type === 'touchstart' || event.type === 'touchend' || event.type === 'touchmove') {
            var touch = event.touch;
            if (touch === undefined) {
              return; // the "touch" property is only defined in SDL
  
            }
            var coords = Browser.calculateMouseCoords(touch.pageX, touch.pageY);
  
            if (event.type === 'touchstart') {
              Browser.lastTouches[touch.identifier] = coords;
              Browser.touches[touch.identifier] = coords;
            } else if (event.type === 'touchend' || event.type === 'touchmove') {
              var last = Browser.touches[touch.identifier];
              last ||= coords;
              Browser.lastTouches[touch.identifier] = last;
              Browser.touches[touch.identifier] = coords;
            }
            return;
          }
  
          Browser.setMouseCoords(event.pageX, event.pageY);
        }
      },
  resizeListeners:[],
  updateResizeListeners() {
        var canvas = Browser.getCanvas();
        Browser.resizeListeners.forEach((listener) => listener(canvas.width, canvas.height));
      },
  setCanvasSize(width, height, noUpdates) {
        var canvas = Browser.getCanvas();
        Browser.updateCanvasDimensions(canvas, width, height);
        if (!noUpdates) Browser.updateResizeListeners();
      },
  windowedWidth:0,
  windowedHeight:0,
  setFullscreenCanvasSize() {
        // check if SDL is available
        if (typeof SDL != "undefined") {
          var flags = HEAPU32[((SDL.screen)>>2)];
          flags = flags | 0x00800000; // set SDL_FULLSCREEN flag
          HEAP32[((SDL.screen)>>2)] = flags;
        }
        Browser.updateCanvasDimensions(Browser.getCanvas());
        Browser.updateResizeListeners();
      },
  setWindowedCanvasSize() {
        // check if SDL is available
        if (typeof SDL != "undefined") {
          var flags = HEAPU32[((SDL.screen)>>2)];
          flags = flags & ~0x00800000; // clear SDL_FULLSCREEN flag
          HEAP32[((SDL.screen)>>2)] = flags;
        }
        Browser.updateCanvasDimensions(Browser.getCanvas());
        Browser.updateResizeListeners();
      },
  updateCanvasDimensions(canvas, wNative, hNative) {
        if (wNative && hNative) {
          canvas.widthNative = wNative;
          canvas.heightNative = hNative;
        } else {
          wNative = canvas.widthNative;
          hNative = canvas.heightNative;
        }
        var w = wNative;
        var h = hNative;
        if (Module['forcedAspectRatio'] > 0) {
          if (w/h < Module['forcedAspectRatio']) {
            w = Math.round(h * Module['forcedAspectRatio']);
          } else {
            h = Math.round(w / Module['forcedAspectRatio']);
          }
        }
        if (((document['fullscreenElement'] || document['mozFullScreenElement'] ||
             document['msFullscreenElement'] || document['webkitFullscreenElement'] ||
             document['webkitCurrentFullScreenElement']) === canvas.parentNode) && (typeof screen != 'undefined')) {
           var factor = Math.min(screen.width / w, screen.height / h);
           w = Math.round(w * factor);
           h = Math.round(h * factor);
        }
        if (Browser.resizeCanvas) {
          if (canvas.width  != w) canvas.width  = w;
          if (canvas.height != h) canvas.height = h;
          if (typeof canvas.style != 'undefined') {
            canvas.style.removeProperty( "width");
            canvas.style.removeProperty("height");
          }
        } else {
          if (canvas.width  != wNative) canvas.width  = wNative;
          if (canvas.height != hNative) canvas.height = hNative;
          if (typeof canvas.style != 'undefined') {
            if (w != wNative || h != hNative) {
              canvas.style.setProperty( "width", w + "px", "important");
              canvas.style.setProperty("height", h + "px", "important");
            } else {
              canvas.style.removeProperty( "width");
              canvas.style.removeProperty("height");
            }
          }
        }
      },
  };
  var getPreloadedImageData = (path, w, h) => {
      path = PATH_FS.resolve(path);
  
      var canvas = /** @type {HTMLCanvasElement} */(Browser.preloadedImages[path]);
      if (!canvas) return 0;
  
      var ctx = canvas.getContext("2d");
      var image = ctx.getImageData(0, 0, canvas.width, canvas.height);
      var buf = _malloc(canvas.width * canvas.height * 4);
  
      HEAPU8.set(image.data, buf);
  
      HEAP32[((w)>>2)] = canvas.width;
      HEAP32[((h)>>2)] = canvas.height;
      return buf;
    };
  
  
  
  var _emscripten_get_preloaded_image_data = (path, w, h) => getPreloadedImageData(UTF8ToString(path), w, h);

  
  
  var _emscripten_get_preloaded_image_data_from_FILE = (file, w, h) => {
      var fd = _fileno(file);
      var stream = FS.getStream(fd);
      if (stream) {
        return getPreloadedImageData(stream.path, w, h);
      }
  
      return 0;
    };

  var _emscripten_get_screen_size = (width, height) => {
      HEAP32[((width)>>2)] = screen.width;
      HEAP32[((height)>>2)] = screen.height;
    };

  var GLctx;
  
  var webgl_enable_ANGLE_instanced_arrays = (ctx) => {
      // Extension available in WebGL 1 from Firefox 26 and Google Chrome 30 onwards. Core feature in WebGL 2.
      var ext = ctx.getExtension('ANGLE_instanced_arrays');
      // Because this extension is a core function in WebGL 2, assign the extension entry points in place of
      // where the core functions will reside in WebGL 2. This way the calling code can call these without
      // having to dynamically branch depending if running against WebGL 1 or WebGL 2.
      if (ext) {
        ctx['vertexAttribDivisor'] = (index, divisor) => ext['vertexAttribDivisorANGLE'](index, divisor);
        ctx['drawArraysInstanced'] = (mode, first, count, primcount) => ext['drawArraysInstancedANGLE'](mode, first, count, primcount);
        ctx['drawElementsInstanced'] = (mode, count, type, indices, primcount) => ext['drawElementsInstancedANGLE'](mode, count, type, indices, primcount);
        return 1;
      }
    };
  
  var webgl_enable_OES_vertex_array_object = (ctx) => {
      // Extension available in WebGL 1 from Firefox 25 and WebKit 536.28/desktop Safari 6.0.3 onwards. Core feature in WebGL 2.
      var ext = ctx.getExtension('OES_vertex_array_object');
      if (ext) {
        ctx['createVertexArray'] = () => ext['createVertexArrayOES']();
        ctx['deleteVertexArray'] = (vao) => ext['deleteVertexArrayOES'](vao);
        ctx['bindVertexArray'] = (vao) => ext['bindVertexArrayOES'](vao);
        ctx['isVertexArray'] = (vao) => ext['isVertexArrayOES'](vao);
        return 1;
      }
    };
  
  var webgl_enable_WEBGL_draw_buffers = (ctx) => {
      // Extension available in WebGL 1 from Firefox 28 onwards. Core feature in WebGL 2.
      var ext = ctx.getExtension('WEBGL_draw_buffers');
      if (ext) {
        ctx['drawBuffers'] = (n, bufs) => ext['drawBuffersWEBGL'](n, bufs);
        return 1;
      }
    };
  
  var webgl_enable_EXT_polygon_offset_clamp = (ctx) =>
      !!(ctx.extPolygonOffsetClamp = ctx.getExtension('EXT_polygon_offset_clamp'));
  
  var webgl_enable_EXT_clip_control = (ctx) =>
      !!(ctx.extClipControl = ctx.getExtension('EXT_clip_control'));
  
  var webgl_enable_WEBGL_polygon_mode = (ctx) =>
      !!(ctx.webglPolygonMode = ctx.getExtension('WEBGL_polygon_mode'));
  
  var webgl_enable_WEBGL_multi_draw = (ctx) =>
      // Closure is expected to be allowed to minify the '.multiDrawWebgl' property, so not accessing it quoted.
      !!(ctx.multiDrawWebgl = ctx.getExtension('WEBGL_multi_draw'));
  
  var getEmscriptenSupportedExtensions = (ctx) => {
      // Restrict the list of advertised extensions to those that we actually
      // support.
      var supportedExtensions = [
        // WebGL 1 extensions
        'ANGLE_instanced_arrays',
        'EXT_blend_minmax',
        'EXT_disjoint_timer_query',
        'EXT_frag_depth',
        'EXT_shader_texture_lod',
        'EXT_sRGB',
        'OES_element_index_uint',
        'OES_fbo_render_mipmap',
        'OES_standard_derivatives',
        'OES_texture_float',
        'OES_texture_half_float',
        'OES_texture_half_float_linear',
        'OES_vertex_array_object',
        'WEBGL_color_buffer_float',
        'WEBGL_depth_texture',
        'WEBGL_draw_buffers',
        // WebGL 1 and WebGL 2 extensions
        'EXT_clip_control',
        'EXT_color_buffer_half_float',
        'EXT_depth_clamp',
        'EXT_float_blend',
        'EXT_polygon_offset_clamp',
        'EXT_texture_compression_bptc',
        'EXT_texture_compression_rgtc',
        'EXT_texture_filter_anisotropic',
        'KHR_parallel_shader_compile',
        'OES_texture_float_linear',
        'WEBGL_blend_func_extended',
        'WEBGL_compressed_texture_astc',
        'WEBGL_compressed_texture_etc',
        'WEBGL_compressed_texture_etc1',
        'WEBGL_compressed_texture_s3tc',
        'WEBGL_compressed_texture_s3tc_srgb',
        'WEBGL_debug_renderer_info',
        'WEBGL_debug_shaders',
        'WEBGL_lose_context',
        'WEBGL_multi_draw',
        'WEBGL_polygon_mode'
      ];
      // .getSupportedExtensions() can return null if context is lost, so coerce to empty array.
      return (ctx.getSupportedExtensions() || []).filter(ext => supportedExtensions.includes(ext));
    };
  
  
  var GL = {
  counter:1,
  buffers:[],
  programs:[],
  framebuffers:[],
  renderbuffers:[],
  textures:[],
  shaders:[],
  vaos:[],
  contexts:[],
  offscreenCanvases:{
  },
  queries:[],
  stringCache:{
  },
  unpackAlignment:4,
  unpackRowLength:0,
  recordError:(errorCode) => {
        if (!GL.lastError) {
          GL.lastError = errorCode;
        }
      },
  getNewId:(table) => {
        var ret = GL.counter++;
        for (var i = table.length; i < ret; i++) {
          table[i] = null;
        }
        return ret;
      },
  genObject:(n, buffers, createFunction, objectTable
        ) => {
        for (var i = 0; i < n; i++) {
          var buffer = GLctx[createFunction]();
          var id = buffer && GL.getNewId(objectTable);
          if (buffer) {
            buffer.name = id;
            objectTable[id] = buffer;
          } else {
            GL.recordError(0x502 /* GL_INVALID_OPERATION */);
          }
          HEAP32[(((buffers)+(i*4))>>2)] = id;
        }
      },
  getSource:(shader, count, string, length) => {
        var source = '';
        for (var i = 0; i < count; ++i) {
          var len = length ? HEAPU32[(((length)+(i*4))>>2)] : undefined;
          source += UTF8ToString(HEAPU32[(((string)+(i*4))>>2)], len);
        }
        return source;
      },
  createContext:(/** @type {HTMLCanvasElement} */ canvas, webGLContextAttributes) => {
  
        // BUG: Workaround Safari WebGL issue: After successfully acquiring WebGL
        // context on a canvas, calling .getContext() will always return that
        // context independent of which 'webgl' or 'webgl2'
        // context version was passed. See:
        //   https://bugs.webkit.org/show_bug.cgi?id=222758
        // and:
        //   https://github.com/emscripten-core/emscripten/issues/13295.
        // TODO: Once the bug is fixed and shipped in Safari, adjust the Safari
        // version field in above check.
        if (!canvas.getContextSafariWebGL2Fixed) {
          canvas.getContextSafariWebGL2Fixed = canvas.getContext;
          /** @type {function(this:HTMLCanvasElement, string, (Object|null)=): (Object|null)} */
          function fixedGetContext(ver, attrs) {
            var gl = canvas.getContextSafariWebGL2Fixed(ver, attrs);
            return ((ver == 'webgl') == (gl instanceof WebGLRenderingContext)) ? gl : null;
          }
          canvas.getContext = fixedGetContext;
        }
  
        var ctx =
          canvas.getContext("webgl", webGLContextAttributes);
  
        if (!ctx) return 0;
  
        var handle = GL.registerContext(ctx, webGLContextAttributes);
  
        return handle;
      },
  registerContext:(ctx, webGLContextAttributes) => {
        // without pthreads a context is just an integer ID
        var handle = GL.getNewId(GL.contexts);
  
        var context = {
          handle,
          attributes: webGLContextAttributes,
          version: webGLContextAttributes.majorVersion,
          GLctx: ctx
        };
  
        // Store the created context object so that we can access the context
        // given a canvas without having to pass the parameters again.
        if (ctx.canvas) ctx.canvas.GLctxObject = context;
        GL.contexts[handle] = context;
        if (typeof webGLContextAttributes.enableExtensionsByDefault == 'undefined' || webGLContextAttributes.enableExtensionsByDefault) {
          GL.initExtensions(context);
        }
  
        return handle;
      },
  makeContextCurrent:(contextHandle) => {
  
        // Active Emscripten GL layer context object.
        GL.currentContext = GL.contexts[contextHandle];
        // Active WebGL context object.
        Module['ctx'] = GLctx = GL.currentContext?.GLctx;
        return !(contextHandle && !GLctx);
      },
  getContext:(contextHandle) => {
        return GL.contexts[contextHandle];
      },
  deleteContext:(contextHandle) => {
        if (GL.currentContext === GL.contexts[contextHandle]) {
          GL.currentContext = null;
        }
        if (typeof JSEvents == 'object') {
          // Release all JS event handlers on the DOM element that the GL context is
          // associated with since the context is now deleted.
          JSEvents.removeAllHandlersOnTarget(GL.contexts[contextHandle].GLctx.canvas);
        }
        // Make sure the canvas object no longer refers to the context object so
        // there are no GC surprises.
        if (GL.contexts[contextHandle]?.GLctx.canvas) {
          GL.contexts[contextHandle].GLctx.canvas.GLctxObject = undefined;
        }
        GL.contexts[contextHandle] = null;
      },
  initExtensions:(context) => {
        // If this function is called without a specific context object, init the
        // extensions of the currently active context.
        context ||= GL.currentContext;
  
        if (context.initExtensionsDone) return;
        context.initExtensionsDone = true;
  
        var GLctx = context.GLctx;
  
        // Detect the presence of a few extensions manually, ction GL interop
        // layer itself will need to know if they exist.
  
        // Extensions that are available in both WebGL 1 and WebGL 2
        webgl_enable_WEBGL_multi_draw(GLctx);
        webgl_enable_EXT_polygon_offset_clamp(GLctx);
        webgl_enable_EXT_clip_control(GLctx);
        webgl_enable_WEBGL_polygon_mode(GLctx);
        // Extensions that are only available in WebGL 1 (the calls will be no-ops
        // if called on a WebGL 2 context active)
        webgl_enable_ANGLE_instanced_arrays(GLctx);
        webgl_enable_OES_vertex_array_object(GLctx);
        webgl_enable_WEBGL_draw_buffers(GLctx);
        {
          GLctx.disjointTimerQueryExt = GLctx.getExtension("EXT_disjoint_timer_query");
        }
  
        getEmscriptenSupportedExtensions(GLctx).forEach((ext) => {
          // WEBGL_lose_context, WEBGL_debug_renderer_info and WEBGL_debug_shaders
          // are not enabled by default.
          if (!ext.includes('lose_context') && !ext.includes('debug')) {
            // Call .getExtension() to enable that extension permanently.
            GLctx.getExtension(ext);
          }
        });
      },
  };
  /** @suppress {duplicate } */
  var _glActiveTexture = (x0) => GLctx.activeTexture(x0);
  var _emscripten_glActiveTexture = _glActiveTexture;

  /** @suppress {duplicate } */
  var _glAttachShader = (program, shader) => {
      GLctx.attachShader(GL.programs[program], GL.shaders[shader]);
    };
  var _emscripten_glAttachShader = _glAttachShader;

  /** @suppress {duplicate } */
  var _glBeginQueryEXT = (target, id) => {
      GLctx.disjointTimerQueryExt['beginQueryEXT'](target, GL.queries[id]);
    };
  var _emscripten_glBeginQueryEXT = _glBeginQueryEXT;

  
  /** @suppress {duplicate } */
  var _glBindAttribLocation = (program, index, name) => {
      GLctx.bindAttribLocation(GL.programs[program], index, UTF8ToString(name));
    };
  var _emscripten_glBindAttribLocation = _glBindAttribLocation;

  /** @suppress {duplicate } */
  var _glBindBuffer = (target, buffer) => {
  
      GLctx.bindBuffer(target, GL.buffers[buffer]);
    };
  var _emscripten_glBindBuffer = _glBindBuffer;

  /** @suppress {duplicate } */
  var _glBindFramebuffer = (target, framebuffer) => {
  
      GLctx.bindFramebuffer(target, GL.framebuffers[framebuffer]);
  
    };
  var _emscripten_glBindFramebuffer = _glBindFramebuffer;

  /** @suppress {duplicate } */
  var _glBindRenderbuffer = (target, renderbuffer) => {
      GLctx.bindRenderbuffer(target, GL.renderbuffers[renderbuffer]);
    };
  var _emscripten_glBindRenderbuffer = _glBindRenderbuffer;

  /** @suppress {duplicate } */
  var _glBindTexture = (target, texture) => {
      GLctx.bindTexture(target, GL.textures[texture]);
    };
  var _emscripten_glBindTexture = _glBindTexture;

  
  /** @suppress {duplicate } */
  var _glBindVertexArray = (vao) => {
      GLctx.bindVertexArray(GL.vaos[vao]);
    };
  /** @suppress {duplicate } */
  var _glBindVertexArrayOES = _glBindVertexArray;
  var _emscripten_glBindVertexArrayOES = _glBindVertexArrayOES;

  /** @suppress {duplicate } */
  var _glBlendColor = (x0, x1, x2, x3) => GLctx.blendColor(x0, x1, x2, x3);
  var _emscripten_glBlendColor = _glBlendColor;

  /** @suppress {duplicate } */
  var _glBlendEquation = (x0) => GLctx.blendEquation(x0);
  var _emscripten_glBlendEquation = _glBlendEquation;

  /** @suppress {duplicate } */
  var _glBlendEquationSeparate = (x0, x1) => GLctx.blendEquationSeparate(x0, x1);
  var _emscripten_glBlendEquationSeparate = _glBlendEquationSeparate;

  /** @suppress {duplicate } */
  var _glBlendFunc = (x0, x1) => GLctx.blendFunc(x0, x1);
  var _emscripten_glBlendFunc = _glBlendFunc;

  /** @suppress {duplicate } */
  var _glBlendFuncSeparate = (x0, x1, x2, x3) => GLctx.blendFuncSeparate(x0, x1, x2, x3);
  var _emscripten_glBlendFuncSeparate = _glBlendFuncSeparate;

  /** @suppress {duplicate } */
  var _glBufferData = (target, size, data, usage) => {
  
      // N.b. here first form specifies a heap subarray, second form an integer
      // size, so the ?: code here is polymorphic. It is advised to avoid
      // randomly mixing both uses in calling code, to avoid any potential JS
      // engine JIT issues.
      GLctx.bufferData(target, data ? HEAPU8.subarray(data, data+size) : size, usage);
    };
  var _emscripten_glBufferData = _glBufferData;

  /** @suppress {duplicate } */
  var _glBufferSubData = (target, offset, size, data) => {
      GLctx.bufferSubData(target, offset, HEAPU8.subarray(data, data+size));
    };
  var _emscripten_glBufferSubData = _glBufferSubData;

  /** @suppress {duplicate } */
  var _glCheckFramebufferStatus = (x0) => GLctx.checkFramebufferStatus(x0);
  var _emscripten_glCheckFramebufferStatus = _glCheckFramebufferStatus;

  /** @suppress {duplicate } */
  var _glClear = (x0) => GLctx.clear(x0);
  var _emscripten_glClear = _glClear;

  /** @suppress {duplicate } */
  var _glClearColor = (x0, x1, x2, x3) => GLctx.clearColor(x0, x1, x2, x3);
  var _emscripten_glClearColor = _glClearColor;

  /** @suppress {duplicate } */
  var _glClearDepthf = (x0) => GLctx.clearDepth(x0);
  var _emscripten_glClearDepthf = _glClearDepthf;

  /** @suppress {duplicate } */
  var _glClearStencil = (x0) => GLctx.clearStencil(x0);
  var _emscripten_glClearStencil = _glClearStencil;

  /** @suppress {duplicate } */
  var _glClipControlEXT = (origin, depth) => {
      GLctx.extClipControl['clipControlEXT'](origin, depth);
    };
  var _emscripten_glClipControlEXT = _glClipControlEXT;

  /** @suppress {duplicate } */
  var _glColorMask = (red, green, blue, alpha) => {
      GLctx.colorMask(!!red, !!green, !!blue, !!alpha);
    };
  var _emscripten_glColorMask = _glColorMask;

  /** @suppress {duplicate } */
  var _glCompileShader = (shader) => {
      GLctx.compileShader(GL.shaders[shader]);
    };
  var _emscripten_glCompileShader = _glCompileShader;

  /** @suppress {duplicate } */
  var _glCompressedTexImage2D = (target, level, internalFormat, width, height, border, imageSize, data) => {
      // `data` may be null here, which means "allocate uniniitalized space but
      // don't upload" in GLES parlance, but `compressedTexImage2D` requires the
      // final data parameter, so we simply pass a heap view starting at zero
      // effectively uploading whatever happens to be near address zero.  See
      // https://github.com/emscripten-core/emscripten/issues/19300.
      GLctx.compressedTexImage2D(target, level, internalFormat, width, height, border, HEAPU8.subarray((data), data+imageSize));
    };
  var _emscripten_glCompressedTexImage2D = _glCompressedTexImage2D;

  /** @suppress {duplicate } */
  var _glCompressedTexSubImage2D = (target, level, xoffset, yoffset, width, height, format, imageSize, data) => {
      GLctx.compressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, HEAPU8.subarray((data), data+imageSize));
    };
  var _emscripten_glCompressedTexSubImage2D = _glCompressedTexSubImage2D;

  /** @suppress {duplicate } */
  var _glCopyTexImage2D = (x0, x1, x2, x3, x4, x5, x6, x7) => GLctx.copyTexImage2D(x0, x1, x2, x3, x4, x5, x6, x7);
  var _emscripten_glCopyTexImage2D = _glCopyTexImage2D;

  /** @suppress {duplicate } */
  var _glCopyTexSubImage2D = (x0, x1, x2, x3, x4, x5, x6, x7) => GLctx.copyTexSubImage2D(x0, x1, x2, x3, x4, x5, x6, x7);
  var _emscripten_glCopyTexSubImage2D = _glCopyTexSubImage2D;

  /** @suppress {duplicate } */
  var _glCreateProgram = () => {
      var id = GL.getNewId(GL.programs);
      var program = GLctx.createProgram();
      // Store additional information needed for each shader program:
      program.name = id;
      // Lazy cache results of
      // glGetProgramiv(GL_ACTIVE_UNIFORM_MAX_LENGTH/GL_ACTIVE_ATTRIBUTE_MAX_LENGTH/GL_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH)
      program.maxUniformLength = program.maxAttributeLength = program.maxUniformBlockNameLength = 0;
      program.uniformIdCounter = 1;
      GL.programs[id] = program;
      return id;
    };
  var _emscripten_glCreateProgram = _glCreateProgram;

  /** @suppress {duplicate } */
  var _glCreateShader = (shaderType) => {
      var id = GL.getNewId(GL.shaders);
      GL.shaders[id] = GLctx.createShader(shaderType);
  
      return id;
    };
  var _emscripten_glCreateShader = _glCreateShader;

  /** @suppress {duplicate } */
  var _glCullFace = (x0) => GLctx.cullFace(x0);
  var _emscripten_glCullFace = _glCullFace;

  /** @suppress {duplicate } */
  var _glDeleteBuffers = (n, buffers) => {
      for (var i = 0; i < n; i++) {
        var id = HEAP32[(((buffers)+(i*4))>>2)];
        var buffer = GL.buffers[id];
  
        // From spec: "glDeleteBuffers silently ignores 0's and names that do not
        // correspond to existing buffer objects."
        if (!buffer) continue;
  
        GLctx.deleteBuffer(buffer);
        buffer.name = 0;
        GL.buffers[id] = null;
  
      }
    };
  var _emscripten_glDeleteBuffers = _glDeleteBuffers;

  /** @suppress {duplicate } */
  var _glDeleteFramebuffers = (n, framebuffers) => {
      for (var i = 0; i < n; ++i) {
        var id = HEAP32[(((framebuffers)+(i*4))>>2)];
        var framebuffer = GL.framebuffers[id];
        if (!framebuffer) continue; // GL spec: "glDeleteFramebuffers silently ignores 0s and names that do not correspond to existing framebuffer objects".
        GLctx.deleteFramebuffer(framebuffer);
        framebuffer.name = 0;
        GL.framebuffers[id] = null;
      }
    };
  var _emscripten_glDeleteFramebuffers = _glDeleteFramebuffers;

  /** @suppress {duplicate } */
  var _glDeleteProgram = (id) => {
      if (!id) return;
      var program = GL.programs[id];
      if (!program) {
        // glDeleteProgram actually signals an error when deleting a nonexisting
        // object, unlike some other GL delete functions.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      GLctx.deleteProgram(program);
      program.name = 0;
      GL.programs[id] = null;
    };
  var _emscripten_glDeleteProgram = _glDeleteProgram;

  /** @suppress {duplicate } */
  var _glDeleteQueriesEXT = (n, ids) => {
      for (var i = 0; i < n; i++) {
        var id = HEAP32[(((ids)+(i*4))>>2)];
        var query = GL.queries[id];
        if (!query) continue; // GL spec: "unused names in ids are ignored, as is the name zero."
        GLctx.disjointTimerQueryExt['deleteQueryEXT'](query);
        GL.queries[id] = null;
      }
    };
  var _emscripten_glDeleteQueriesEXT = _glDeleteQueriesEXT;

  /** @suppress {duplicate } */
  var _glDeleteRenderbuffers = (n, renderbuffers) => {
      for (var i = 0; i < n; i++) {
        var id = HEAP32[(((renderbuffers)+(i*4))>>2)];
        var renderbuffer = GL.renderbuffers[id];
        if (!renderbuffer) continue; // GL spec: "glDeleteRenderbuffers silently ignores 0s and names that do not correspond to existing renderbuffer objects".
        GLctx.deleteRenderbuffer(renderbuffer);
        renderbuffer.name = 0;
        GL.renderbuffers[id] = null;
      }
    };
  var _emscripten_glDeleteRenderbuffers = _glDeleteRenderbuffers;

  /** @suppress {duplicate } */
  var _glDeleteShader = (id) => {
      if (!id) return;
      var shader = GL.shaders[id];
      if (!shader) {
        // glDeleteShader actually signals an error when deleting a nonexisting
        // object, unlike some other GL delete functions.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      GLctx.deleteShader(shader);
      GL.shaders[id] = null;
    };
  var _emscripten_glDeleteShader = _glDeleteShader;

  /** @suppress {duplicate } */
  var _glDeleteTextures = (n, textures) => {
      for (var i = 0; i < n; i++) {
        var id = HEAP32[(((textures)+(i*4))>>2)];
        var texture = GL.textures[id];
        // GL spec: "glDeleteTextures silently ignores 0s and names that do not
        // correspond to existing textures".
        if (!texture) continue;
        GLctx.deleteTexture(texture);
        texture.name = 0;
        GL.textures[id] = null;
      }
    };
  var _emscripten_glDeleteTextures = _glDeleteTextures;

  
  /** @suppress {duplicate } */
  var _glDeleteVertexArrays = (n, vaos) => {
      for (var i = 0; i < n; i++) {
        var id = HEAP32[(((vaos)+(i*4))>>2)];
        GLctx.deleteVertexArray(GL.vaos[id]);
        GL.vaos[id] = null;
      }
    };
  /** @suppress {duplicate } */
  var _glDeleteVertexArraysOES = _glDeleteVertexArrays;
  var _emscripten_glDeleteVertexArraysOES = _glDeleteVertexArraysOES;

  /** @suppress {duplicate } */
  var _glDepthFunc = (x0) => GLctx.depthFunc(x0);
  var _emscripten_glDepthFunc = _glDepthFunc;

  /** @suppress {duplicate } */
  var _glDepthMask = (flag) => {
      GLctx.depthMask(!!flag);
    };
  var _emscripten_glDepthMask = _glDepthMask;

  /** @suppress {duplicate } */
  var _glDepthRangef = (x0, x1) => GLctx.depthRange(x0, x1);
  var _emscripten_glDepthRangef = _glDepthRangef;

  /** @suppress {duplicate } */
  var _glDetachShader = (program, shader) => {
      GLctx.detachShader(GL.programs[program], GL.shaders[shader]);
    };
  var _emscripten_glDetachShader = _glDetachShader;

  /** @suppress {duplicate } */
  var _glDisable = (x0) => GLctx.disable(x0);
  var _emscripten_glDisable = _glDisable;

  /** @suppress {duplicate } */
  var _glDisableVertexAttribArray = (index) => {
      GLctx.disableVertexAttribArray(index);
    };
  var _emscripten_glDisableVertexAttribArray = _glDisableVertexAttribArray;

  /** @suppress {duplicate } */
  var _glDrawArrays = (mode, first, count) => {
  
      GLctx.drawArrays(mode, first, count);
  
    };
  var _emscripten_glDrawArrays = _glDrawArrays;

  
  /** @suppress {duplicate } */
  var _glDrawArraysInstanced = (mode, first, count, primcount) => {
      GLctx.drawArraysInstanced(mode, first, count, primcount);
    };
  /** @suppress {duplicate } */
  var _glDrawArraysInstancedANGLE = _glDrawArraysInstanced;
  var _emscripten_glDrawArraysInstancedANGLE = _glDrawArraysInstancedANGLE;

  
  var tempFixedLengthArray = [];
  
  /** @suppress {duplicate } */
  var _glDrawBuffers = (n, bufs) => {
  
      var bufArray = tempFixedLengthArray[n];
      for (var i = 0; i < n; i++) {
        bufArray[i] = HEAP32[(((bufs)+(i*4))>>2)];
      }
  
      GLctx.drawBuffers(bufArray);
    };
  /** @suppress {duplicate } */
  var _glDrawBuffersWEBGL = _glDrawBuffers;
  var _emscripten_glDrawBuffersWEBGL = _glDrawBuffersWEBGL;

  /** @suppress {duplicate } */
  var _glDrawElements = (mode, count, type, indices) => {
  
      GLctx.drawElements(mode, count, type, indices);
  
    };
  var _emscripten_glDrawElements = _glDrawElements;

  
  /** @suppress {duplicate } */
  var _glDrawElementsInstanced = (mode, count, type, indices, primcount) => {
      GLctx.drawElementsInstanced(mode, count, type, indices, primcount);
    };
  /** @suppress {duplicate } */
  var _glDrawElementsInstancedANGLE = _glDrawElementsInstanced;
  var _emscripten_glDrawElementsInstancedANGLE = _glDrawElementsInstancedANGLE;

  /** @suppress {duplicate } */
  var _glEnable = (x0) => GLctx.enable(x0);
  var _emscripten_glEnable = _glEnable;

  /** @suppress {duplicate } */
  var _glEnableVertexAttribArray = (index) => {
      GLctx.enableVertexAttribArray(index);
    };
  var _emscripten_glEnableVertexAttribArray = _glEnableVertexAttribArray;

  /** @suppress {duplicate } */
  var _glEndQueryEXT = (target) => {
      GLctx.disjointTimerQueryExt['endQueryEXT'](target);
    };
  var _emscripten_glEndQueryEXT = _glEndQueryEXT;

  /** @suppress {duplicate } */
  var _glFinish = () => GLctx.finish();
  var _emscripten_glFinish = _glFinish;

  /** @suppress {duplicate } */
  var _glFlush = () => GLctx.flush();
  var _emscripten_glFlush = _glFlush;

  /** @suppress {duplicate } */
  var _glFramebufferRenderbuffer = (target, attachment, renderbuffertarget, renderbuffer) => {
      GLctx.framebufferRenderbuffer(target, attachment, renderbuffertarget,
                                         GL.renderbuffers[renderbuffer]);
    };
  var _emscripten_glFramebufferRenderbuffer = _glFramebufferRenderbuffer;

  /** @suppress {duplicate } */
  var _glFramebufferTexture2D = (target, attachment, textarget, texture, level) => {
      GLctx.framebufferTexture2D(target, attachment, textarget,
                                      GL.textures[texture], level);
    };
  var _emscripten_glFramebufferTexture2D = _glFramebufferTexture2D;

  /** @suppress {duplicate } */
  var _glFrontFace = (x0) => GLctx.frontFace(x0);
  var _emscripten_glFrontFace = _glFrontFace;

  /** @suppress {duplicate } */
  var _glGenBuffers = (n, buffers) => {
      GL.genObject(n, buffers, 'createBuffer', GL.buffers
        );
    };
  var _emscripten_glGenBuffers = _glGenBuffers;

  /** @suppress {duplicate } */
  var _glGenFramebuffers = (n, ids) => {
      GL.genObject(n, ids, 'createFramebuffer', GL.framebuffers
        );
    };
  var _emscripten_glGenFramebuffers = _glGenFramebuffers;

  /** @suppress {duplicate } */
  var _glGenQueriesEXT = (n, ids) => {
      for (var i = 0; i < n; i++) {
        var query = GLctx.disjointTimerQueryExt['createQueryEXT']();
        if (!query) {
          GL.recordError(0x502 /* GL_INVALID_OPERATION */);
          while (i < n) HEAP32[(((ids)+(i++*4))>>2)] = 0;
          return;
        }
        var id = GL.getNewId(GL.queries);
        query.name = id;
        GL.queries[id] = query;
        HEAP32[(((ids)+(i*4))>>2)] = id;
      }
    };
  var _emscripten_glGenQueriesEXT = _glGenQueriesEXT;

  /** @suppress {duplicate } */
  var _glGenRenderbuffers = (n, renderbuffers) => {
      GL.genObject(n, renderbuffers, 'createRenderbuffer', GL.renderbuffers
        );
    };
  var _emscripten_glGenRenderbuffers = _glGenRenderbuffers;

  /** @suppress {duplicate } */
  var _glGenTextures = (n, textures) => {
      GL.genObject(n, textures, 'createTexture', GL.textures
        );
    };
  var _emscripten_glGenTextures = _glGenTextures;

  
  /** @suppress {duplicate } */
  var _glGenVertexArrays = (n, arrays) => {
      GL.genObject(n, arrays, 'createVertexArray', GL.vaos
        );
    };
  /** @suppress {duplicate } */
  var _glGenVertexArraysOES = _glGenVertexArrays;
  var _emscripten_glGenVertexArraysOES = _glGenVertexArraysOES;

  /** @suppress {duplicate } */
  var _glGenerateMipmap = (x0) => GLctx.generateMipmap(x0);
  var _emscripten_glGenerateMipmap = _glGenerateMipmap;

  
  var __glGetActiveAttribOrUniform = (funcName, program, index, bufSize, length, size, type, name) => {
      program = GL.programs[program];
      var info = GLctx[funcName](program, index);
      if (info) {
        // If an error occurs, nothing will be written to length, size and type and name.
        var numBytesWrittenExclNull = name && stringToUTF8(info.name, name, bufSize);
        if (length) HEAP32[((length)>>2)] = numBytesWrittenExclNull;
        if (size) HEAP32[((size)>>2)] = info.size;
        if (type) HEAP32[((type)>>2)] = info.type;
      }
    };
  
  /** @suppress {duplicate } */
  var _glGetActiveAttrib = (program, index, bufSize, length, size, type, name) =>
      __glGetActiveAttribOrUniform('getActiveAttrib', program, index, bufSize, length, size, type, name);
  var _emscripten_glGetActiveAttrib = _glGetActiveAttrib;

  
  /** @suppress {duplicate } */
  var _glGetActiveUniform = (program, index, bufSize, length, size, type, name) =>
      __glGetActiveAttribOrUniform('getActiveUniform', program, index, bufSize, length, size, type, name);
  var _emscripten_glGetActiveUniform = _glGetActiveUniform;

  /** @suppress {duplicate } */
  var _glGetAttachedShaders = (program, maxCount, count, shaders) => {
      var result = GLctx.getAttachedShaders(GL.programs[program]);
      var len = result.length;
      if (len > maxCount) {
        len = maxCount;
      }
      HEAP32[((count)>>2)] = len;
      for (var i = 0; i < len; ++i) {
        var id = GL.shaders.indexOf(result[i]);
        HEAP32[(((shaders)+(i*4))>>2)] = id;
      }
    };
  var _emscripten_glGetAttachedShaders = _glGetAttachedShaders;

  
  /** @suppress {duplicate } */
  var _glGetAttribLocation = (program, name) =>
      GLctx.getAttribLocation(GL.programs[program], UTF8ToString(name));
  var _emscripten_glGetAttribLocation = _glGetAttribLocation;

  var readI53FromI64 = (ptr) => {
      return HEAPU32[((ptr)>>2)] + HEAP32[(((ptr)+(4))>>2)] * 4294967296;
    };
  
  var readI53FromU64 = (ptr) => {
      return HEAPU32[((ptr)>>2)] + HEAPU32[(((ptr)+(4))>>2)] * 4294967296;
    };
  var writeI53ToI64 = (ptr, num) => {
      HEAPU32[((ptr)>>2)] = num;
      var lower = HEAPU32[((ptr)>>2)];
      HEAPU32[(((ptr)+(4))>>2)] = (num - lower)/4294967296;
      var deserialized = (num >= 0) ? readI53FromU64(ptr) : readI53FromI64(ptr);
      var offset = ((ptr)>>2);
      if (deserialized != num) warnOnce(`writeI53ToI64() out of range: serialized JS Number ${num} to Wasm heap as bytes lo=${ptrToString(HEAPU32[offset])}, hi=${ptrToString(HEAPU32[offset+1])}, which deserializes back to ${deserialized} instead!`);
    };
  
  var emscriptenWebGLGet = (name_, p, type) => {
      // Guard against user passing a null pointer.
      // Note that GLES2 spec does not say anything about how passing a null
      // pointer should be treated.  Testing on desktop core GL 3, the application
      // crashes on glGetIntegerv to a null pointer, but better to report an error
      // instead of doing anything random.
      if (!p) {
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      var ret = undefined;
      switch (name_) { // Handle a few trivial GLES values
        case 0x8DFA: // GL_SHADER_COMPILER
          ret = 1;
          break;
        case 0x8DF8: // GL_SHADER_BINARY_FORMATS
          if (type != 0 && type != 1) {
            GL.recordError(0x500); // GL_INVALID_ENUM
          }
          // Do not write anything to the out pointer, since no binary formats are
          // supported.
          return;
        case 0x8DF9: // GL_NUM_SHADER_BINARY_FORMATS
          ret = 0;
          break;
        case 0x86A2: // GL_NUM_COMPRESSED_TEXTURE_FORMATS
          // WebGL doesn't have GL_NUM_COMPRESSED_TEXTURE_FORMATS (it's obsolete
          // since GL_COMPRESSED_TEXTURE_FORMATS returns a JS array that can be
          // queried for length), so implement it ourselves to allow C++ GLES2
          // code get the length.
          var formats = GLctx.getParameter(0x86A3 /*GL_COMPRESSED_TEXTURE_FORMATS*/);
          ret = formats ? formats.length : 0;
          break;
  
      }
  
      if (ret === undefined) {
        var result = GLctx.getParameter(name_);
        switch (typeof result) {
          case "number":
            ret = result;
            break;
          case "boolean":
            ret = result ? 1 : 0;
            break;
          case "string":
            GL.recordError(0x500); // GL_INVALID_ENUM
            return;
          case "object":
            if (result === null) {
              // null is a valid result for some (e.g., which buffer is bound -
              // perhaps nothing is bound), but otherwise can mean an invalid
              // name_, which we need to report as an error
              switch (name_) {
                case 0x8894: // ARRAY_BUFFER_BINDING
                case 0x8B8D: // CURRENT_PROGRAM
                case 0x8895: // ELEMENT_ARRAY_BUFFER_BINDING
                case 0x8CA6: // FRAMEBUFFER_BINDING or DRAW_FRAMEBUFFER_BINDING
                case 0x8CA7: // RENDERBUFFER_BINDING
                case 0x8069: // TEXTURE_BINDING_2D
                case 0x85B5: // WebGL 2 GL_VERTEX_ARRAY_BINDING, or WebGL 1 extension OES_vertex_array_object GL_VERTEX_ARRAY_BINDING_OES
                case 0x8514: { // TEXTURE_BINDING_CUBE_MAP
                  ret = 0;
                  break;
                }
                default: {
                  GL.recordError(0x500); // GL_INVALID_ENUM
                  return;
                }
              }
            } else if (result instanceof Float32Array ||
                       result instanceof Uint32Array ||
                       result instanceof Int32Array ||
                       result instanceof Array) {
              for (var i = 0; i < result.length; ++i) {
                switch (type) {
                  case 0: HEAP32[(((p)+(i*4))>>2)] = result[i]; break;
                  case 2: HEAPF32[(((p)+(i*4))>>2)] = result[i]; break;
                  case 4: HEAP8[(p)+(i)] = result[i] ? 1 : 0; break;
                }
              }
              return;
            } else {
              try {
                ret = result.name | 0;
              } catch(e) {
                GL.recordError(0x500); // GL_INVALID_ENUM
                err(`GL_INVALID_ENUM in glGet${type}v: Unknown object returned from WebGL getParameter(${name_})! (error: ${e})`);
                return;
              }
            }
            break;
          default:
            GL.recordError(0x500); // GL_INVALID_ENUM
            err(`GL_INVALID_ENUM in glGet${type}v: Native code calling glGet${type}v(${name_}) and it returns ${result} of type ${typeof(result)}!`);
            return;
        }
      }
  
      switch (type) {
        case 1: writeI53ToI64(p, ret); break;
        case 0: HEAP32[((p)>>2)] = ret; break;
        case 2:   HEAPF32[((p)>>2)] = ret; break;
        case 4: HEAP8[p] = ret ? 1 : 0; break;
      }
    };
  
  /** @suppress {duplicate } */
  var _glGetBooleanv = (name_, p) => emscriptenWebGLGet(name_, p, 4);
  var _emscripten_glGetBooleanv = _glGetBooleanv;

  /** @suppress {duplicate } */
  var _glGetBufferParameteriv = (target, value, data) => {
      if (!data) {
        // GLES2 specification does not specify how to behave if data is a null
        // pointer. Since calling this function does not make sense if data ==
        // null, issue a GL error to notify user about it.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      HEAP32[((data)>>2)] = GLctx.getBufferParameter(target, value);
    };
  var _emscripten_glGetBufferParameteriv = _glGetBufferParameteriv;

  /** @suppress {duplicate } */
  var _glGetError = () => {
      var error = GLctx.getError() || GL.lastError;
      GL.lastError = 0/*GL_NO_ERROR*/;
      return error;
    };
  var _emscripten_glGetError = _glGetError;

  
  /** @suppress {duplicate } */
  var _glGetFloatv = (name_, p) => emscriptenWebGLGet(name_, p, 2);
  var _emscripten_glGetFloatv = _glGetFloatv;

  /** @suppress {duplicate } */
  var _glGetFramebufferAttachmentParameteriv = (target, attachment, pname, params) => {
      var result = GLctx.getFramebufferAttachmentParameter(target, attachment, pname);
      if (result instanceof WebGLRenderbuffer ||
          result instanceof WebGLTexture) {
        result = result.name | 0;
      }
      HEAP32[((params)>>2)] = result;
    };
  var _emscripten_glGetFramebufferAttachmentParameteriv = _glGetFramebufferAttachmentParameteriv;

  
  /** @suppress {duplicate } */
  var _glGetIntegerv = (name_, p) => emscriptenWebGLGet(name_, p, 0);
  var _emscripten_glGetIntegerv = _glGetIntegerv;

  /** @suppress {duplicate } */
  var _glGetProgramInfoLog = (program, maxLength, length, infoLog) => {
      var log = GLctx.getProgramInfoLog(GL.programs[program]);
      if (log === null) log = '(unknown error)';
      var numBytesWrittenExclNull = (maxLength > 0 && infoLog) ? stringToUTF8(log, infoLog, maxLength) : 0;
      if (length) HEAP32[((length)>>2)] = numBytesWrittenExclNull;
    };
  var _emscripten_glGetProgramInfoLog = _glGetProgramInfoLog;

  /** @suppress {duplicate } */
  var _glGetProgramiv = (program, pname, p) => {
      if (!p) {
        // GLES2 specification does not specify how to behave if p is a null
        // pointer. Since calling this function does not make sense if p == null,
        // issue a GL error to notify user about it.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
  
      if (program >= GL.counter) {
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
  
      program = GL.programs[program];
  
      if (pname == 0x8B84) { // GL_INFO_LOG_LENGTH
        var log = GLctx.getProgramInfoLog(program);
        if (log === null) log = '(unknown error)';
        HEAP32[((p)>>2)] = log.length + 1;
      } else if (pname == 0x8B87 /* GL_ACTIVE_UNIFORM_MAX_LENGTH */) {
        if (!program.maxUniformLength) {
          var numActiveUniforms = GLctx.getProgramParameter(program, 0x8B86/*GL_ACTIVE_UNIFORMS*/);
          for (var i = 0; i < numActiveUniforms; ++i) {
            program.maxUniformLength = Math.max(program.maxUniformLength, GLctx.getActiveUniform(program, i).name.length+1);
          }
        }
        HEAP32[((p)>>2)] = program.maxUniformLength;
      } else if (pname == 0x8B8A /* GL_ACTIVE_ATTRIBUTE_MAX_LENGTH */) {
        if (!program.maxAttributeLength) {
          var numActiveAttributes = GLctx.getProgramParameter(program, 0x8B89/*GL_ACTIVE_ATTRIBUTES*/);
          for (var i = 0; i < numActiveAttributes; ++i) {
            program.maxAttributeLength = Math.max(program.maxAttributeLength, GLctx.getActiveAttrib(program, i).name.length+1);
          }
        }
        HEAP32[((p)>>2)] = program.maxAttributeLength;
      } else if (pname == 0x8A35 /* GL_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH */) {
        if (!program.maxUniformBlockNameLength) {
          var numActiveUniformBlocks = GLctx.getProgramParameter(program, 0x8A36/*GL_ACTIVE_UNIFORM_BLOCKS*/);
          for (var i = 0; i < numActiveUniformBlocks; ++i) {
            program.maxUniformBlockNameLength = Math.max(program.maxUniformBlockNameLength, GLctx.getActiveUniformBlockName(program, i).length+1);
          }
        }
        HEAP32[((p)>>2)] = program.maxUniformBlockNameLength;
      } else {
        HEAP32[((p)>>2)] = GLctx.getProgramParameter(program, pname);
      }
    };
  var _emscripten_glGetProgramiv = _glGetProgramiv;

  
  /** @suppress {duplicate } */
  var _glGetQueryObjecti64vEXT = (id, pname, params) => {
      if (!params) {
        // GLES2 specification does not specify how to behave if params is a null pointer. Since calling this function does not make sense
        // if p == null, issue a GL error to notify user about it.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      var query = GL.queries[id];
      var param;
      {
        param = GLctx.disjointTimerQueryExt['getQueryObjectEXT'](query, pname);
      }
      var ret;
      if (typeof param == 'boolean') {
        ret = param ? 1 : 0;
      } else {
        ret = param;
      }
      writeI53ToI64(params, ret);
    };
  var _emscripten_glGetQueryObjecti64vEXT = _glGetQueryObjecti64vEXT;

  /** @suppress {duplicate } */
  var _glGetQueryObjectivEXT = (id, pname, params) => {
      if (!params) {
        // GLES2 specification does not specify how to behave if params is a null pointer. Since calling this function does not make sense
        // if p == null, issue a GL error to notify user about it.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      var query = GL.queries[id];
      var param = GLctx.disjointTimerQueryExt['getQueryObjectEXT'](query, pname);
      var ret;
      if (typeof param == 'boolean') {
        ret = param ? 1 : 0;
      } else {
        ret = param;
      }
      HEAP32[((params)>>2)] = ret;
    };
  var _emscripten_glGetQueryObjectivEXT = _glGetQueryObjectivEXT;

  
  /** @suppress {duplicate } */
  var _glGetQueryObjectui64vEXT = _glGetQueryObjecti64vEXT;
  var _emscripten_glGetQueryObjectui64vEXT = _glGetQueryObjectui64vEXT;

  
  /** @suppress {duplicate } */
  var _glGetQueryObjectuivEXT = _glGetQueryObjectivEXT;
  var _emscripten_glGetQueryObjectuivEXT = _glGetQueryObjectuivEXT;

  /** @suppress {duplicate } */
  var _glGetQueryivEXT = (target, pname, params) => {
      if (!params) {
        // GLES2 specification does not specify how to behave if params is a null pointer. Since calling this function does not make sense
        // if p == null, issue a GL error to notify user about it.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      HEAP32[((params)>>2)] = GLctx.disjointTimerQueryExt['getQueryEXT'](target, pname);
    };
  var _emscripten_glGetQueryivEXT = _glGetQueryivEXT;

  /** @suppress {duplicate } */
  var _glGetRenderbufferParameteriv = (target, pname, params) => {
      if (!params) {
        // GLES2 specification does not specify how to behave if params is a null pointer. Since calling this function does not make sense
        // if params == null, issue a GL error to notify user about it.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      HEAP32[((params)>>2)] = GLctx.getRenderbufferParameter(target, pname);
    };
  var _emscripten_glGetRenderbufferParameteriv = _glGetRenderbufferParameteriv;

  
  /** @suppress {duplicate } */
  var _glGetShaderInfoLog = (shader, maxLength, length, infoLog) => {
      var log = GLctx.getShaderInfoLog(GL.shaders[shader]);
      if (log === null) log = '(unknown error)';
      var numBytesWrittenExclNull = (maxLength > 0 && infoLog) ? stringToUTF8(log, infoLog, maxLength) : 0;
      if (length) HEAP32[((length)>>2)] = numBytesWrittenExclNull;
    };
  var _emscripten_glGetShaderInfoLog = _glGetShaderInfoLog;

  /** @suppress {duplicate } */
  var _glGetShaderPrecisionFormat = (shaderType, precisionType, range, precision) => {
      var result = GLctx.getShaderPrecisionFormat(shaderType, precisionType);
      HEAP32[((range)>>2)] = result.rangeMin;
      HEAP32[(((range)+(4))>>2)] = result.rangeMax;
      HEAP32[((precision)>>2)] = result.precision;
    };
  var _emscripten_glGetShaderPrecisionFormat = _glGetShaderPrecisionFormat;

  /** @suppress {duplicate } */
  var _glGetShaderSource = (shader, bufSize, length, source) => {
      var result = GLctx.getShaderSource(GL.shaders[shader]);
      if (!result) return; // If an error occurs, nothing will be written to length or source.
      var numBytesWrittenExclNull = (bufSize > 0 && source) ? stringToUTF8(result, source, bufSize) : 0;
      if (length) HEAP32[((length)>>2)] = numBytesWrittenExclNull;
    };
  var _emscripten_glGetShaderSource = _glGetShaderSource;

  /** @suppress {duplicate } */
  var _glGetShaderiv = (shader, pname, p) => {
      if (!p) {
        // GLES2 specification does not specify how to behave if p is a null
        // pointer. Since calling this function does not make sense if p == null,
        // issue a GL error to notify user about it.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      if (pname == 0x8B84) { // GL_INFO_LOG_LENGTH
        var log = GLctx.getShaderInfoLog(GL.shaders[shader]);
        if (log === null) log = '(unknown error)';
        // The GLES2 specification says that if the shader has an empty info log,
        // a value of 0 is returned. Otherwise the log has a null char appended.
        // (An empty string is falsey, so we can just check that instead of
        // looking at log.length.)
        var logLength = log ? log.length + 1 : 0;
        HEAP32[((p)>>2)] = logLength;
      } else if (pname == 0x8B88) { // GL_SHADER_SOURCE_LENGTH
        var source = GLctx.getShaderSource(GL.shaders[shader]);
        // source may be a null, or the empty string, both of which are falsey
        // values that we report a 0 length for.
        var sourceLength = source ? source.length + 1 : 0;
        HEAP32[((p)>>2)] = sourceLength;
      } else {
        HEAP32[((p)>>2)] = GLctx.getShaderParameter(GL.shaders[shader], pname);
      }
    };
  var _emscripten_glGetShaderiv = _glGetShaderiv;

  
  
  var stringToNewUTF8 = (str) => {
      var size = lengthBytesUTF8(str) + 1;
      var ret = _malloc(size);
      if (ret) stringToUTF8(str, ret, size);
      return ret;
    };
  
  
  var webglGetExtensions = () => {
      var exts = getEmscriptenSupportedExtensions(GLctx);
      exts = exts.concat(exts.map((e) => "GL_" + e));
      return exts;
    };
  
  /** @suppress {duplicate } */
  var _glGetString = (name_) => {
      var ret = GL.stringCache[name_];
      if (!ret) {
        switch (name_) {
          case 0x1F03 /* GL_EXTENSIONS */:
            ret = stringToNewUTF8(webglGetExtensions().join(' '));
            break;
          case 0x1F00 /* GL_VENDOR */:
          case 0x1F01 /* GL_RENDERER */:
          case 0x9245 /* UNMASKED_VENDOR_WEBGL */:
          case 0x9246 /* UNMASKED_RENDERER_WEBGL */:
            var s = GLctx.getParameter(name_);
            if (!s) {
              GL.recordError(0x500/*GL_INVALID_ENUM*/);
            }
            ret = s ? stringToNewUTF8(s) : 0;
            break;
  
          case 0x1F02 /* GL_VERSION */:
            var webGLVersion = GLctx.getParameter(0x1F02 /*GL_VERSION*/);
            // return GLES version string corresponding to the version of the WebGL context
            var glVersion = `OpenGL ES 2.0 (${webGLVersion})`;
            ret = stringToNewUTF8(glVersion);
            break;
          case 0x8B8C /* GL_SHADING_LANGUAGE_VERSION */:
            var glslVersion = GLctx.getParameter(0x8B8C /*GL_SHADING_LANGUAGE_VERSION*/);
            // extract the version number 'N.M' from the string 'WebGL GLSL ES N.M ...'
            var ver_re = /^WebGL GLSL ES ([0-9]\.[0-9][0-9]?)(?:$| .*)/;
            var ver_num = glslVersion.match(ver_re);
            if (ver_num !== null) {
              if (ver_num[1].length == 3) ver_num[1] = ver_num[1] + '0'; // ensure minor version has 2 digits
              glslVersion = `OpenGL ES GLSL ES ${ver_num[1]} (${glslVersion})`;
            }
            ret = stringToNewUTF8(glslVersion);
            break;
          default:
            GL.recordError(0x500/*GL_INVALID_ENUM*/);
            // fall through
        }
        GL.stringCache[name_] = ret;
      }
      return ret;
    };
  var _emscripten_glGetString = _glGetString;

  /** @suppress {duplicate } */
  var _glGetTexParameterfv = (target, pname, params) => {
      if (!params) {
        // GLES2 specification does not specify how to behave if params is a null
        // pointer. Since calling this function does not make sense if p == null,
        // issue a GL error to notify user about it.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      HEAPF32[((params)>>2)] = GLctx.getTexParameter(target, pname);
    };
  var _emscripten_glGetTexParameterfv = _glGetTexParameterfv;

  /** @suppress {duplicate } */
  var _glGetTexParameteriv = (target, pname, params) => {
      if (!params) {
        // GLES2 specification does not specify how to behave if params is a null
        // pointer. Since calling this function does not make sense if p == null,
        // issue a GL error to notify user about it.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      HEAP32[((params)>>2)] = GLctx.getTexParameter(target, pname);
    };
  var _emscripten_glGetTexParameteriv = _glGetTexParameteriv;

  /** @suppress {checkTypes} */
  var jstoi_q = (str) => parseInt(str);
  
  /** @noinline */
  var webglGetLeftBracePos = (name) => name.slice(-1) == ']' && name.lastIndexOf('[');
  
  var webglPrepareUniformLocationsBeforeFirstUse = (program) => {
      var uniformLocsById = program.uniformLocsById, // Maps GLuint -> WebGLUniformLocation
        uniformSizeAndIdsByName = program.uniformSizeAndIdsByName, // Maps name -> [uniform array length, GLuint]
        i, j;
  
      // On the first time invocation of glGetUniformLocation on this shader program:
      // initialize cache data structures and discover which uniforms are arrays.
      if (!uniformLocsById) {
        // maps GLint integer locations to WebGLUniformLocations
        program.uniformLocsById = uniformLocsById = {};
        // maps integer locations back to uniform name strings, so that we can lazily fetch uniform array locations
        program.uniformArrayNamesById = {};
  
        var numActiveUniforms = GLctx.getProgramParameter(program, 0x8B86/*GL_ACTIVE_UNIFORMS*/);
        for (i = 0; i < numActiveUniforms; ++i) {
          var u = GLctx.getActiveUniform(program, i);
          var nm = u.name;
          var sz = u.size;
          var lb = webglGetLeftBracePos(nm);
          var arrayName = lb > 0 ? nm.slice(0, lb) : nm;
  
          // Assign a new location.
          var id = program.uniformIdCounter;
          program.uniformIdCounter += sz;
          // Eagerly get the location of the uniformArray[0] base element.
          // The remaining indices >0 will be left for lazy evaluation to
          // improve performance. Those may never be needed to fetch, if the
          // application fills arrays always in full starting from the first
          // element of the array.
          uniformSizeAndIdsByName[arrayName] = [sz, id];
  
          // Store placeholder integers in place that highlight that these
          // >0 index locations are array indices pending population.
          for (j = 0; j < sz; ++j) {
            uniformLocsById[id] = j;
            program.uniformArrayNamesById[id++] = arrayName;
          }
        }
      }
    };
  
  
  
  /** @suppress {duplicate } */
  var _glGetUniformLocation = (program, name) => {
  
      name = UTF8ToString(name);
  
      if (program = GL.programs[program]) {
        webglPrepareUniformLocationsBeforeFirstUse(program);
        var uniformLocsById = program.uniformLocsById; // Maps GLuint -> WebGLUniformLocation
        var arrayIndex = 0;
        var uniformBaseName = name;
  
        // Invariant: when populating integer IDs for uniform locations, we must
        // maintain the precondition that arrays reside in contiguous addresses,
        // i.e. for a 'vec4 colors[10];', colors[4] must be at location
        // colors[0]+4.  However, user might call glGetUniformLocation(program,
        // "colors") for an array, so we cannot discover based on the user input
        // arguments whether the uniform we are dealing with is an array. The only
        // way to discover which uniforms are arrays is to enumerate over all the
        // active uniforms in the program.
        var leftBrace = webglGetLeftBracePos(name);
  
        // If user passed an array accessor "[index]", parse the array index off the accessor.
        if (leftBrace > 0) {
          arrayIndex = jstoi_q(name.slice(leftBrace + 1)) >>> 0; // "index]", coerce parseInt(']') with >>>0 to treat "foo[]" as "foo[0]" and foo[-1] as unsigned out-of-bounds.
          uniformBaseName = name.slice(0, leftBrace);
        }
  
        // Have we cached the location of this uniform before?
        // A pair [array length, GLint of the uniform location]
        var sizeAndId = program.uniformSizeAndIdsByName[uniformBaseName];
  
        // If an uniform with this name exists, and if its index is within the
        // array limits (if it's even an array), query the WebGLlocation, or
        // return an existing cached location.
        if (sizeAndId && arrayIndex < sizeAndId[0]) {
          arrayIndex += sizeAndId[1]; // Add the base location of the uniform to the array index offset.
          if ((uniformLocsById[arrayIndex] = uniformLocsById[arrayIndex] || GLctx.getUniformLocation(program, name))) {
            return arrayIndex;
          }
        }
      }
      else {
        // N.b. we are currently unable to distinguish between GL program IDs that
        // never existed vs GL program IDs that have been deleted, so report
        // GL_INVALID_VALUE in both cases.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
      }
      return -1;
    };
  var _emscripten_glGetUniformLocation = _glGetUniformLocation;

  var webglGetUniformLocation = (location) => {
      var p = GLctx.currentProgram;
  
      if (p) {
        var webglLoc = p.uniformLocsById[location];
        // p.uniformLocsById[location] stores either an integer, or a
        // WebGLUniformLocation.
        // If an integer, we have not yet bound the location, so do it now. The
        // integer value specifies the array index we should bind to.
        if (typeof webglLoc == 'number') {
          p.uniformLocsById[location] = webglLoc = GLctx.getUniformLocation(p, p.uniformArrayNamesById[location] + (webglLoc > 0 ? `[${webglLoc}]` : ''));
        }
        // Else an already cached WebGLUniformLocation, return it.
        return webglLoc;
      } else {
        GL.recordError(0x502/*GL_INVALID_OPERATION*/);
      }
    };
  
  
  /** @suppress{checkTypes} */
  var emscriptenWebGLGetUniform = (program, location, params, type) => {
      if (!params) {
        // GLES2 specification does not specify how to behave if params is a null
        // pointer. Since calling this function does not make sense if params ==
        // null, issue a GL error to notify user about it.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      program = GL.programs[program];
      webglPrepareUniformLocationsBeforeFirstUse(program);
      var data = GLctx.getUniform(program, webglGetUniformLocation(location));
      if (typeof data == 'number' || typeof data == 'boolean') {
        switch (type) {
          case 0: HEAP32[((params)>>2)] = data; break;
          case 2: HEAPF32[((params)>>2)] = data; break;
        }
      } else {
        for (var i = 0; i < data.length; i++) {
          switch (type) {
            case 0: HEAP32[(((params)+(i*4))>>2)] = data[i]; break;
            case 2: HEAPF32[(((params)+(i*4))>>2)] = data[i]; break;
          }
        }
      }
    };
  
  /** @suppress {duplicate } */
  var _glGetUniformfv = (program, location, params) => {
      emscriptenWebGLGetUniform(program, location, params, 2);
    };
  var _emscripten_glGetUniformfv = _glGetUniformfv;

  
  /** @suppress {duplicate } */
  var _glGetUniformiv = (program, location, params) => {
      emscriptenWebGLGetUniform(program, location, params, 0);
    };
  var _emscripten_glGetUniformiv = _glGetUniformiv;

  /** @suppress {duplicate } */
  var _glGetVertexAttribPointerv = (index, pname, pointer) => {
      if (!pointer) {
        // GLES2 specification does not specify how to behave if pointer is a null
        // pointer. Since calling this function does not make sense if pointer ==
        // null, issue a GL error to notify user about it.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      HEAP32[((pointer)>>2)] = GLctx.getVertexAttribOffset(index, pname);
    };
  var _emscripten_glGetVertexAttribPointerv = _glGetVertexAttribPointerv;

  /** @suppress{checkTypes} */
  var emscriptenWebGLGetVertexAttrib = (index, pname, params, type) => {
      if (!params) {
        // GLES2 specification does not specify how to behave if params is a null
        // pointer. Since calling this function does not make sense if params ==
        // null, issue a GL error to notify user about it.
        GL.recordError(0x501 /* GL_INVALID_VALUE */);
        return;
      }
      var data = GLctx.getVertexAttrib(index, pname);
      if (pname == 0x889F/*VERTEX_ATTRIB_ARRAY_BUFFER_BINDING*/) {
        HEAP32[((params)>>2)] = data && data["name"];
      } else if (typeof data == 'number' || typeof data == 'boolean') {
        switch (type) {
          case 0: HEAP32[((params)>>2)] = data; break;
          case 2: HEAPF32[((params)>>2)] = data; break;
          case 5: HEAP32[((params)>>2)] = Math.fround(data); break;
        }
      } else {
        for (var i = 0; i < data.length; i++) {
          switch (type) {
            case 0: HEAP32[(((params)+(i*4))>>2)] = data[i]; break;
            case 2: HEAPF32[(((params)+(i*4))>>2)] = data[i]; break;
            case 5: HEAP32[(((params)+(i*4))>>2)] = Math.fround(data[i]); break;
          }
        }
      }
    };
  
  /** @suppress {duplicate } */
  var _glGetVertexAttribfv = (index, pname, params) => {
      // N.B. This function may only be called if the vertex attribute was
      // specified using the function glVertexAttrib*f(), otherwise the results
      // are undefined. (GLES3 spec 6.1.12)
      emscriptenWebGLGetVertexAttrib(index, pname, params, 2);
    };
  var _emscripten_glGetVertexAttribfv = _glGetVertexAttribfv;

  
  /** @suppress {duplicate } */
  var _glGetVertexAttribiv = (index, pname, params) => {
      // N.B. This function may only be called if the vertex attribute was
      // specified using the function glVertexAttrib*f(), otherwise the results
      // are undefined. (GLES3 spec 6.1.12)
      emscriptenWebGLGetVertexAttrib(index, pname, params, 5);
    };
  var _emscripten_glGetVertexAttribiv = _glGetVertexAttribiv;

  /** @suppress {duplicate } */
  var _glHint = (x0, x1) => GLctx.hint(x0, x1);
  var _emscripten_glHint = _glHint;

  /** @suppress {duplicate } */
  var _glIsBuffer = (buffer) => {
      var b = GL.buffers[buffer];
      if (!b) return 0;
      return GLctx.isBuffer(b);
    };
  var _emscripten_glIsBuffer = _glIsBuffer;

  /** @suppress {duplicate } */
  var _glIsEnabled = (x0) => GLctx.isEnabled(x0);
  var _emscripten_glIsEnabled = _glIsEnabled;

  /** @suppress {duplicate } */
  var _glIsFramebuffer = (framebuffer) => {
      var fb = GL.framebuffers[framebuffer];
      if (!fb) return 0;
      return GLctx.isFramebuffer(fb);
    };
  var _emscripten_glIsFramebuffer = _glIsFramebuffer;

  /** @suppress {duplicate } */
  var _glIsProgram = (program) => {
      program = GL.programs[program];
      if (!program) return 0;
      return GLctx.isProgram(program);
    };
  var _emscripten_glIsProgram = _glIsProgram;

  /** @suppress {duplicate } */
  var _glIsQueryEXT = (id) => {
      var query = GL.queries[id];
      if (!query) return 0;
      return GLctx.disjointTimerQueryExt['isQueryEXT'](query);
    };
  var _emscripten_glIsQueryEXT = _glIsQueryEXT;

  /** @suppress {duplicate } */
  var _glIsRenderbuffer = (renderbuffer) => {
      var rb = GL.renderbuffers[renderbuffer];
      if (!rb) return 0;
      return GLctx.isRenderbuffer(rb);
    };
  var _emscripten_glIsRenderbuffer = _glIsRenderbuffer;

  /** @suppress {duplicate } */
  var _glIsShader = (shader) => {
      var s = GL.shaders[shader];
      if (!s) return 0;
      return GLctx.isShader(s);
    };
  var _emscripten_glIsShader = _glIsShader;

  /** @suppress {duplicate } */
  var _glIsTexture = (id) => {
      var texture = GL.textures[id];
      if (!texture) return 0;
      return GLctx.isTexture(texture);
    };
  var _emscripten_glIsTexture = _glIsTexture;

  
  /** @suppress {duplicate } */
  var _glIsVertexArray = (array) => {
  
      var vao = GL.vaos[array];
      if (!vao) return 0;
      return GLctx.isVertexArray(vao);
    };
  /** @suppress {duplicate } */
  var _glIsVertexArrayOES = _glIsVertexArray;
  var _emscripten_glIsVertexArrayOES = _glIsVertexArrayOES;

  /** @suppress {duplicate } */
  var _glLineWidth = (x0) => GLctx.lineWidth(x0);
  var _emscripten_glLineWidth = _glLineWidth;

  /** @suppress {duplicate } */
  var _glLinkProgram = (program) => {
      program = GL.programs[program];
      GLctx.linkProgram(program);
      // Invalidate earlier computed uniform->ID mappings, those have now become stale
      program.uniformLocsById = 0; // Mark as null-like so that glGetUniformLocation() knows to populate this again.
      program.uniformSizeAndIdsByName = {};
  
    };
  var _emscripten_glLinkProgram = _glLinkProgram;

  /** @suppress {duplicate } */
  var _glPixelStorei = (pname, param) => {
      if (pname == 3317) {
        GL.unpackAlignment = param;
      } else if (pname == 3314) {
        GL.unpackRowLength = param;
      }
      GLctx.pixelStorei(pname, param);
    };
  var _emscripten_glPixelStorei = _glPixelStorei;

  /** @suppress {duplicate } */
  var _glPolygonModeWEBGL = (face, mode) => {
      GLctx.webglPolygonMode['polygonModeWEBGL'](face, mode);
    };
  var _emscripten_glPolygonModeWEBGL = _glPolygonModeWEBGL;

  /** @suppress {duplicate } */
  var _glPolygonOffset = (x0, x1) => GLctx.polygonOffset(x0, x1);
  var _emscripten_glPolygonOffset = _glPolygonOffset;

  /** @suppress {duplicate } */
  var _glPolygonOffsetClampEXT = (factor, units, clamp) => {
      GLctx.extPolygonOffsetClamp['polygonOffsetClampEXT'](factor, units, clamp);
    };
  var _emscripten_glPolygonOffsetClampEXT = _glPolygonOffsetClampEXT;

  /** @suppress {duplicate } */
  var _glQueryCounterEXT = (id, target) => {
      GLctx.disjointTimerQueryExt['queryCounterEXT'](GL.queries[id], target);
    };
  var _emscripten_glQueryCounterEXT = _glQueryCounterEXT;

  var computeUnpackAlignedImageSize = (width, height, sizePerPixel) => {
      function roundedToNextMultipleOf(x, y) {
        return (x + y - 1) & -y;
      }
      var plainRowSize = (GL.unpackRowLength || width) * sizePerPixel;
      var alignedRowSize = roundedToNextMultipleOf(plainRowSize, GL.unpackAlignment);
      return height * alignedRowSize;
    };
  
  var colorChannelsInGlTextureFormat = (format) => {
      // Micro-optimizations for size: map format to size by subtracting smallest
      // enum value (0x1902) from all values first.  Also omit the most common
      // size value (1) from the list, which is assumed by formats not on the
      // list.
      var colorChannels = {
        // 0x1902 /* GL_DEPTH_COMPONENT */ - 0x1902: 1,
        // 0x1906 /* GL_ALPHA */ - 0x1902: 1,
        5: 3,
        6: 4,
        // 0x1909 /* GL_LUMINANCE */ - 0x1902: 1,
        8: 2,
        29502: 3,
        29504: 4,
      };
      return colorChannels[format - 0x1902]||1;
    };
  
  var heapObjectForWebGLType = (type) => {
      // Micro-optimization for size: Subtract lowest GL enum number (0x1400/* GL_BYTE */) from type to compare
      // smaller values for the heap, for shorter generated code size.
      // Also the type HEAPU16 is not tested for explicitly, but any unrecognized type will return out HEAPU16.
      // (since most types are HEAPU16)
      type -= 0x1400;
  
      if (type == 1) return HEAPU8;
  
      if (type == 4) return HEAP32;
  
      if (type == 6) return HEAPF32;
  
      if (type == 5
        || type == 28922
        )
        return HEAPU32;
  
      return HEAPU16;
    };
  
  var toTypedArrayIndex = (pointer, heap) =>
      pointer >>> (31 - Math.clz32(heap.BYTES_PER_ELEMENT));
  
  var emscriptenWebGLGetTexPixelData = (type, format, width, height, pixels, internalFormat) => {
      var heap = heapObjectForWebGLType(type);
      var sizePerPixel = colorChannelsInGlTextureFormat(format) * heap.BYTES_PER_ELEMENT;
      var bytes = computeUnpackAlignedImageSize(width, height, sizePerPixel);
      return heap.subarray(toTypedArrayIndex(pixels, heap), toTypedArrayIndex(pixels + bytes, heap));
    };
  
  /** @suppress {duplicate } */
  var _glReadPixels = (x, y, width, height, format, type, pixels) => {
      var pixelData = emscriptenWebGLGetTexPixelData(type, format, width, height, pixels, format);
      if (!pixelData) {
        GL.recordError(0x500/*GL_INVALID_ENUM*/);
        return;
      }
      GLctx.readPixels(x, y, width, height, format, type, pixelData);
    };
  var _emscripten_glReadPixels = _glReadPixels;

  /** @suppress {duplicate } */
  var _glReleaseShaderCompiler = () => {
      // NOP (as allowed by GLES 2.0 spec)
    };
  var _emscripten_glReleaseShaderCompiler = _glReleaseShaderCompiler;

  /** @suppress {duplicate } */
  var _glRenderbufferStorage = (x0, x1, x2, x3) => GLctx.renderbufferStorage(x0, x1, x2, x3);
  var _emscripten_glRenderbufferStorage = _glRenderbufferStorage;

  /** @suppress {duplicate } */
  var _glSampleCoverage = (value, invert) => {
      GLctx.sampleCoverage(value, !!invert);
    };
  var _emscripten_glSampleCoverage = _glSampleCoverage;

  /** @suppress {duplicate } */
  var _glScissor = (x0, x1, x2, x3) => GLctx.scissor(x0, x1, x2, x3);
  var _emscripten_glScissor = _glScissor;

  /** @suppress {duplicate } */
  var _glShaderBinary = (count, shaders, binaryformat, binary, length) => {
      GL.recordError(0x500/*GL_INVALID_ENUM*/);
    };
  var _emscripten_glShaderBinary = _glShaderBinary;

  /** @suppress {duplicate } */
  var _glShaderSource = (shader, count, string, length) => {
      var source = GL.getSource(shader, count, string, length);
  
      GLctx.shaderSource(GL.shaders[shader], source);
    };
  var _emscripten_glShaderSource = _glShaderSource;

  /** @suppress {duplicate } */
  var _glStencilFunc = (x0, x1, x2) => GLctx.stencilFunc(x0, x1, x2);
  var _emscripten_glStencilFunc = _glStencilFunc;

  /** @suppress {duplicate } */
  var _glStencilFuncSeparate = (x0, x1, x2, x3) => GLctx.stencilFuncSeparate(x0, x1, x2, x3);
  var _emscripten_glStencilFuncSeparate = _glStencilFuncSeparate;

  /** @suppress {duplicate } */
  var _glStencilMask = (x0) => GLctx.stencilMask(x0);
  var _emscripten_glStencilMask = _glStencilMask;

  /** @suppress {duplicate } */
  var _glStencilMaskSeparate = (x0, x1) => GLctx.stencilMaskSeparate(x0, x1);
  var _emscripten_glStencilMaskSeparate = _glStencilMaskSeparate;

  /** @suppress {duplicate } */
  var _glStencilOp = (x0, x1, x2) => GLctx.stencilOp(x0, x1, x2);
  var _emscripten_glStencilOp = _glStencilOp;

  /** @suppress {duplicate } */
  var _glStencilOpSeparate = (x0, x1, x2, x3) => GLctx.stencilOpSeparate(x0, x1, x2, x3);
  var _emscripten_glStencilOpSeparate = _glStencilOpSeparate;

  
  /** @suppress {duplicate } */
  var _glTexImage2D = (target, level, internalFormat, width, height, border, format, type, pixels) => {
      var pixelData = pixels ? emscriptenWebGLGetTexPixelData(type, format, width, height, pixels, internalFormat) : null;
      GLctx.texImage2D(target, level, internalFormat, width, height, border, format, type, pixelData);
    };
  var _emscripten_glTexImage2D = _glTexImage2D;

  /** @suppress {duplicate } */
  var _glTexParameterf = (x0, x1, x2) => GLctx.texParameterf(x0, x1, x2);
  var _emscripten_glTexParameterf = _glTexParameterf;

  /** @suppress {duplicate } */
  var _glTexParameterfv = (target, pname, params) => {
      var param = HEAPF32[((params)>>2)];
      GLctx.texParameterf(target, pname, param);
    };
  var _emscripten_glTexParameterfv = _glTexParameterfv;

  /** @suppress {duplicate } */
  var _glTexParameteri = (x0, x1, x2) => GLctx.texParameteri(x0, x1, x2);
  var _emscripten_glTexParameteri = _glTexParameteri;

  /** @suppress {duplicate } */
  var _glTexParameteriv = (target, pname, params) => {
      var param = HEAP32[((params)>>2)];
      GLctx.texParameteri(target, pname, param);
    };
  var _emscripten_glTexParameteriv = _glTexParameteriv;

  
  /** @suppress {duplicate } */
  var _glTexSubImage2D = (target, level, xoffset, yoffset, width, height, format, type, pixels) => {
      var pixelData = pixels ? emscriptenWebGLGetTexPixelData(type, format, width, height, pixels, 0) : null;
      GLctx.texSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixelData);
    };
  var _emscripten_glTexSubImage2D = _glTexSubImage2D;

  
  /** @suppress {duplicate } */
  var _glUniform1f = (location, v0) => {
      GLctx.uniform1f(webglGetUniformLocation(location), v0);
    };
  var _emscripten_glUniform1f = _glUniform1f;

  
  var miniTempWebGLFloatBuffers = [];
  
  /** @suppress {duplicate } */
  var _glUniform1fv = (location, count, value) => {
  
      if (count <= 288) {
        // avoid allocation when uploading few enough uniforms
        var view = miniTempWebGLFloatBuffers[count];
        for (var i = 0; i < count; ++i) {
          view[i] = HEAPF32[(((value)+(4*i))>>2)];
        }
      } else
      {
        var view = HEAPF32.subarray((((value)>>2)), ((value+count*4)>>2));
      }
      GLctx.uniform1fv(webglGetUniformLocation(location), view);
    };
  var _emscripten_glUniform1fv = _glUniform1fv;

  
  /** @suppress {duplicate } */
  var _glUniform1i = (location, v0) => {
      GLctx.uniform1i(webglGetUniformLocation(location), v0);
    };
  var _emscripten_glUniform1i = _glUniform1i;

  
  var miniTempWebGLIntBuffers = [];
  
  /** @suppress {duplicate } */
  var _glUniform1iv = (location, count, value) => {
  
      if (count <= 288) {
        // avoid allocation when uploading few enough uniforms
        var view = miniTempWebGLIntBuffers[count];
        for (var i = 0; i < count; ++i) {
          view[i] = HEAP32[(((value)+(4*i))>>2)];
        }
      } else
      {
        var view = HEAP32.subarray((((value)>>2)), ((value+count*4)>>2));
      }
      GLctx.uniform1iv(webglGetUniformLocation(location), view);
    };
  var _emscripten_glUniform1iv = _glUniform1iv;

  
  /** @suppress {duplicate } */
  var _glUniform2f = (location, v0, v1) => {
      GLctx.uniform2f(webglGetUniformLocation(location), v0, v1);
    };
  var _emscripten_glUniform2f = _glUniform2f;

  
  
  /** @suppress {duplicate } */
  var _glUniform2fv = (location, count, value) => {
  
      if (count <= 144) {
        // avoid allocation when uploading few enough uniforms
        count *= 2;
        var view = miniTempWebGLFloatBuffers[count];
        for (var i = 0; i < count; i += 2) {
          view[i] = HEAPF32[(((value)+(4*i))>>2)];
          view[i+1] = HEAPF32[(((value)+(4*i+4))>>2)];
        }
      } else
      {
        var view = HEAPF32.subarray((((value)>>2)), ((value+count*8)>>2));
      }
      GLctx.uniform2fv(webglGetUniformLocation(location), view);
    };
  var _emscripten_glUniform2fv = _glUniform2fv;

  
  /** @suppress {duplicate } */
  var _glUniform2i = (location, v0, v1) => {
      GLctx.uniform2i(webglGetUniformLocation(location), v0, v1);
    };
  var _emscripten_glUniform2i = _glUniform2i;

  
  
  /** @suppress {duplicate } */
  var _glUniform2iv = (location, count, value) => {
  
      if (count <= 144) {
        // avoid allocation when uploading few enough uniforms
        count *= 2;
        var view = miniTempWebGLIntBuffers[count];
        for (var i = 0; i < count; i += 2) {
          view[i] = HEAP32[(((value)+(4*i))>>2)];
          view[i+1] = HEAP32[(((value)+(4*i+4))>>2)];
        }
      } else
      {
        var view = HEAP32.subarray((((value)>>2)), ((value+count*8)>>2));
      }
      GLctx.uniform2iv(webglGetUniformLocation(location), view);
    };
  var _emscripten_glUniform2iv = _glUniform2iv;

  
  /** @suppress {duplicate } */
  var _glUniform3f = (location, v0, v1, v2) => {
      GLctx.uniform3f(webglGetUniformLocation(location), v0, v1, v2);
    };
  var _emscripten_glUniform3f = _glUniform3f;

  
  
  /** @suppress {duplicate } */
  var _glUniform3fv = (location, count, value) => {
  
      if (count <= 96) {
        // avoid allocation when uploading few enough uniforms
        count *= 3;
        var view = miniTempWebGLFloatBuffers[count];
        for (var i = 0; i < count; i += 3) {
          view[i] = HEAPF32[(((value)+(4*i))>>2)];
          view[i+1] = HEAPF32[(((value)+(4*i+4))>>2)];
          view[i+2] = HEAPF32[(((value)+(4*i+8))>>2)];
        }
      } else
      {
        var view = HEAPF32.subarray((((value)>>2)), ((value+count*12)>>2));
      }
      GLctx.uniform3fv(webglGetUniformLocation(location), view);
    };
  var _emscripten_glUniform3fv = _glUniform3fv;

  
  /** @suppress {duplicate } */
  var _glUniform3i = (location, v0, v1, v2) => {
      GLctx.uniform3i(webglGetUniformLocation(location), v0, v1, v2);
    };
  var _emscripten_glUniform3i = _glUniform3i;

  
  
  /** @suppress {duplicate } */
  var _glUniform3iv = (location, count, value) => {
  
      if (count <= 96) {
        // avoid allocation when uploading few enough uniforms
        count *= 3;
        var view = miniTempWebGLIntBuffers[count];
        for (var i = 0; i < count; i += 3) {
          view[i] = HEAP32[(((value)+(4*i))>>2)];
          view[i+1] = HEAP32[(((value)+(4*i+4))>>2)];
          view[i+2] = HEAP32[(((value)+(4*i+8))>>2)];
        }
      } else
      {
        var view = HEAP32.subarray((((value)>>2)), ((value+count*12)>>2));
      }
      GLctx.uniform3iv(webglGetUniformLocation(location), view);
    };
  var _emscripten_glUniform3iv = _glUniform3iv;

  
  /** @suppress {duplicate } */
  var _glUniform4f = (location, v0, v1, v2, v3) => {
      GLctx.uniform4f(webglGetUniformLocation(location), v0, v1, v2, v3);
    };
  var _emscripten_glUniform4f = _glUniform4f;

  
  
  /** @suppress {duplicate } */
  var _glUniform4fv = (location, count, value) => {
  
      if (count <= 72) {
        // avoid allocation when uploading few enough uniforms
        var view = miniTempWebGLFloatBuffers[4*count];
        // hoist the heap out of the loop for size and for pthreads+growth.
        var heap = HEAPF32;
        value = ((value)>>2);
        count *= 4;
        for (var i = 0; i < count; i += 4) {
          var dst = value + i;
          view[i] = heap[dst];
          view[i + 1] = heap[dst + 1];
          view[i + 2] = heap[dst + 2];
          view[i + 3] = heap[dst + 3];
        }
      } else
      {
        var view = HEAPF32.subarray((((value)>>2)), ((value+count*16)>>2));
      }
      GLctx.uniform4fv(webglGetUniformLocation(location), view);
    };
  var _emscripten_glUniform4fv = _glUniform4fv;

  
  /** @suppress {duplicate } */
  var _glUniform4i = (location, v0, v1, v2, v3) => {
      GLctx.uniform4i(webglGetUniformLocation(location), v0, v1, v2, v3);
    };
  var _emscripten_glUniform4i = _glUniform4i;

  
  
  /** @suppress {duplicate } */
  var _glUniform4iv = (location, count, value) => {
  
      if (count <= 72) {
        // avoid allocation when uploading few enough uniforms
        count *= 4;
        var view = miniTempWebGLIntBuffers[count];
        for (var i = 0; i < count; i += 4) {
          view[i] = HEAP32[(((value)+(4*i))>>2)];
          view[i+1] = HEAP32[(((value)+(4*i+4))>>2)];
          view[i+2] = HEAP32[(((value)+(4*i+8))>>2)];
          view[i+3] = HEAP32[(((value)+(4*i+12))>>2)];
        }
      } else
      {
        var view = HEAP32.subarray((((value)>>2)), ((value+count*16)>>2));
      }
      GLctx.uniform4iv(webglGetUniformLocation(location), view);
    };
  var _emscripten_glUniform4iv = _glUniform4iv;

  
  
  /** @suppress {duplicate } */
  var _glUniformMatrix2fv = (location, count, transpose, value) => {
  
      if (count <= 72) {
        // avoid allocation when uploading few enough uniforms
        count *= 4;
        var view = miniTempWebGLFloatBuffers[count];
        for (var i = 0; i < count; i += 4) {
          view[i] = HEAPF32[(((value)+(4*i))>>2)];
          view[i+1] = HEAPF32[(((value)+(4*i+4))>>2)];
          view[i+2] = HEAPF32[(((value)+(4*i+8))>>2)];
          view[i+3] = HEAPF32[(((value)+(4*i+12))>>2)];
        }
      } else
      {
        var view = HEAPF32.subarray((((value)>>2)), ((value+count*16)>>2));
      }
      GLctx.uniformMatrix2fv(webglGetUniformLocation(location), !!transpose, view);
    };
  var _emscripten_glUniformMatrix2fv = _glUniformMatrix2fv;

  
  
  /** @suppress {duplicate } */
  var _glUniformMatrix3fv = (location, count, transpose, value) => {
  
      if (count <= 32) {
        // avoid allocation when uploading few enough uniforms
        count *= 9;
        var view = miniTempWebGLFloatBuffers[count];
        for (var i = 0; i < count; i += 9) {
          view[i] = HEAPF32[(((value)+(4*i))>>2)];
          view[i+1] = HEAPF32[(((value)+(4*i+4))>>2)];
          view[i+2] = HEAPF32[(((value)+(4*i+8))>>2)];
          view[i+3] = HEAPF32[(((value)+(4*i+12))>>2)];
          view[i+4] = HEAPF32[(((value)+(4*i+16))>>2)];
          view[i+5] = HEAPF32[(((value)+(4*i+20))>>2)];
          view[i+6] = HEAPF32[(((value)+(4*i+24))>>2)];
          view[i+7] = HEAPF32[(((value)+(4*i+28))>>2)];
          view[i+8] = HEAPF32[(((value)+(4*i+32))>>2)];
        }
      } else
      {
        var view = HEAPF32.subarray((((value)>>2)), ((value+count*36)>>2));
      }
      GLctx.uniformMatrix3fv(webglGetUniformLocation(location), !!transpose, view);
    };
  var _emscripten_glUniformMatrix3fv = _glUniformMatrix3fv;

  
  
  /** @suppress {duplicate } */
  var _glUniformMatrix4fv = (location, count, transpose, value) => {
  
      if (count <= 18) {
        // avoid allocation when uploading few enough uniforms
        var view = miniTempWebGLFloatBuffers[16*count];
        // hoist the heap out of the loop for size and for pthreads+growth.
        var heap = HEAPF32;
        value = ((value)>>2);
        count *= 16;
        for (var i = 0; i < count; i += 16) {
          var dst = value + i;
          view[i] = heap[dst];
          view[i + 1] = heap[dst + 1];
          view[i + 2] = heap[dst + 2];
          view[i + 3] = heap[dst + 3];
          view[i + 4] = heap[dst + 4];
          view[i + 5] = heap[dst + 5];
          view[i + 6] = heap[dst + 6];
          view[i + 7] = heap[dst + 7];
          view[i + 8] = heap[dst + 8];
          view[i + 9] = heap[dst + 9];
          view[i + 10] = heap[dst + 10];
          view[i + 11] = heap[dst + 11];
          view[i + 12] = heap[dst + 12];
          view[i + 13] = heap[dst + 13];
          view[i + 14] = heap[dst + 14];
          view[i + 15] = heap[dst + 15];
        }
      } else
      {
        var view = HEAPF32.subarray((((value)>>2)), ((value+count*64)>>2));
      }
      GLctx.uniformMatrix4fv(webglGetUniformLocation(location), !!transpose, view);
    };
  var _emscripten_glUniformMatrix4fv = _glUniformMatrix4fv;

  /** @suppress {duplicate } */
  var _glUseProgram = (program) => {
      program = GL.programs[program];
      GLctx.useProgram(program);
      // Record the currently active program so that we can access the uniform
      // mapping table of that program.
      GLctx.currentProgram = program;
    };
  var _emscripten_glUseProgram = _glUseProgram;

  /** @suppress {duplicate } */
  var _glValidateProgram = (program) => {
      GLctx.validateProgram(GL.programs[program]);
    };
  var _emscripten_glValidateProgram = _glValidateProgram;

  /** @suppress {duplicate } */
  var _glVertexAttrib1f = (x0, x1) => GLctx.vertexAttrib1f(x0, x1);
  var _emscripten_glVertexAttrib1f = _glVertexAttrib1f;

  /** @suppress {duplicate } */
  var _glVertexAttrib1fv = (index, v) => {
  
      GLctx.vertexAttrib1f(index, HEAPF32[v>>2]);
    };
  var _emscripten_glVertexAttrib1fv = _glVertexAttrib1fv;

  /** @suppress {duplicate } */
  var _glVertexAttrib2f = (x0, x1, x2) => GLctx.vertexAttrib2f(x0, x1, x2);
  var _emscripten_glVertexAttrib2f = _glVertexAttrib2f;

  /** @suppress {duplicate } */
  var _glVertexAttrib2fv = (index, v) => {
  
      GLctx.vertexAttrib2f(index, HEAPF32[v>>2], HEAPF32[v+4>>2]);
    };
  var _emscripten_glVertexAttrib2fv = _glVertexAttrib2fv;

  /** @suppress {duplicate } */
  var _glVertexAttrib3f = (x0, x1, x2, x3) => GLctx.vertexAttrib3f(x0, x1, x2, x3);
  var _emscripten_glVertexAttrib3f = _glVertexAttrib3f;

  /** @suppress {duplicate } */
  var _glVertexAttrib3fv = (index, v) => {
  
      GLctx.vertexAttrib3f(index, HEAPF32[v>>2], HEAPF32[v+4>>2], HEAPF32[v+8>>2]);
    };
  var _emscripten_glVertexAttrib3fv = _glVertexAttrib3fv;

  /** @suppress {duplicate } */
  var _glVertexAttrib4f = (x0, x1, x2, x3, x4) => GLctx.vertexAttrib4f(x0, x1, x2, x3, x4);
  var _emscripten_glVertexAttrib4f = _glVertexAttrib4f;

  /** @suppress {duplicate } */
  var _glVertexAttrib4fv = (index, v) => {
  
      GLctx.vertexAttrib4f(index, HEAPF32[v>>2], HEAPF32[v+4>>2], HEAPF32[v+8>>2], HEAPF32[v+12>>2]);
    };
  var _emscripten_glVertexAttrib4fv = _glVertexAttrib4fv;

  
  /** @suppress {duplicate } */
  var _glVertexAttribDivisor = (index, divisor) => {
      GLctx.vertexAttribDivisor(index, divisor);
    };
  /** @suppress {duplicate } */
  var _glVertexAttribDivisorANGLE = _glVertexAttribDivisor;
  var _emscripten_glVertexAttribDivisorANGLE = _glVertexAttribDivisorANGLE;

  /** @suppress {duplicate } */
  var _glVertexAttribPointer = (index, size, type, normalized, stride, ptr) => {
      GLctx.vertexAttribPointer(index, size, type, !!normalized, stride, ptr);
    };
  var _emscripten_glVertexAttribPointer = _glVertexAttribPointer;

  /** @suppress {duplicate } */
  var _glViewport = (x0, x1, x2, x3) => GLctx.viewport(x0, x1, x2, x3);
  var _emscripten_glViewport = _glViewport;

  var _emscripten_has_asyncify = () => 0;

  
  
  var doRequestFullscreen = (target, strategy) => {
      if (!JSEvents.fullscreenEnabled()) return -1;
      target = findEventTarget(target);
      if (!target) return -4;
  
      if (!target.requestFullscreen
        && !target.webkitRequestFullscreen
        ) {
        return -3;
      }
  
      // Queue this function call if we're not currently in an event handler and
      // the user saw it appropriate to do so.
      if (!JSEvents.canPerformEventHandlerRequests()) {
        if (strategy.deferUntilInEventHandler) {
          JSEvents.deferCall(JSEvents_requestFullscreen, 1 /* priority over pointer lock */, [target, strategy]);
          return 1;
        }
        return -2;
      }
  
      return JSEvents_requestFullscreen(target, strategy);
    };
  var _emscripten_request_fullscreen_strategy = (target, deferUntilInEventHandler, fullscreenStrategy) => {
      var strategy = {
        scaleMode: HEAP32[((fullscreenStrategy)>>2)],
        canvasResolutionScaleMode: HEAP32[(((fullscreenStrategy)+(4))>>2)],
        filteringMode: HEAP32[(((fullscreenStrategy)+(8))>>2)],
        deferUntilInEventHandler,
        canvasResizedCallback: HEAP32[(((fullscreenStrategy)+(12))>>2)],
        canvasResizedCallbackUserData: HEAP32[(((fullscreenStrategy)+(16))>>2)]
      };
  
      return doRequestFullscreen(target, strategy);
    };

  
  
  var _emscripten_request_pointerlock = (target, deferUntilInEventHandler) => {
      target = findEventTarget(target);
      if (!target) return -4;
      if (!target.requestPointerLock) {
        return -1;
      }
  
      // Queue this function call if we're not currently in an event handler and
      // the user saw it appropriate to do so.
      if (!JSEvents.canPerformEventHandlerRequests()) {
        if (deferUntilInEventHandler) {
          JSEvents.deferCall(requestPointerLock, 2 /* priority below fullscreen */, [target]);
          return 1;
        }
        return -2;
      }
  
      return requestPointerLock(target);
    };

  
  
  var growMemory = (size) => {
      var oldHeapSize = wasmMemory.buffer.byteLength;
      var pages = ((size - oldHeapSize + 65535) / 65536) | 0;
      try {
        // round size grow request up to wasm page size (fixed 64KB per spec)
        wasmMemory.grow(pages); // .grow() takes a delta compared to the previous size
        updateMemoryViews();
        return 1 /*success*/;
      } catch(e) {
        err(`growMemory: Attempted to grow heap from ${oldHeapSize} bytes to ${size} bytes, but got error: ${e}`);
      }
      // implicit 0 return to save code size (caller will cast "undefined" into 0
      // anyhow)
    };
  var _emscripten_resize_heap = (requestedSize) => {
      var oldSize = HEAPU8.length;
      // With CAN_ADDRESS_2GB or MEMORY64, pointers are already unsigned.
      requestedSize >>>= 0;
      // With multithreaded builds, races can happen (another thread might increase the size
      // in between), so return a failure, and let the caller retry.
      assert(requestedSize > oldSize);
  
      // Memory resize rules:
      // 1.  Always increase heap size to at least the requested size, rounded up
      //     to next page multiple.
      // 2a. If MEMORY_GROWTH_LINEAR_STEP == -1, excessively resize the heap
      //     geometrically: increase the heap size according to
      //     MEMORY_GROWTH_GEOMETRIC_STEP factor (default +20%), At most
      //     overreserve by MEMORY_GROWTH_GEOMETRIC_CAP bytes (default 96MB).
      // 2b. If MEMORY_GROWTH_LINEAR_STEP != -1, excessively resize the heap
      //     linearly: increase the heap size by at least
      //     MEMORY_GROWTH_LINEAR_STEP bytes.
      // 3.  Max size for the heap is capped at 2048MB-WASM_PAGE_SIZE, or by
      //     MAXIMUM_MEMORY, or by ASAN limit, depending on which is smallest
      // 4.  If we were unable to allocate as much memory, it may be due to
      //     over-eager decision to excessively reserve due to (3) above.
      //     Hence if an allocation fails, cut down on the amount of excess
      //     growth, in an attempt to succeed to perform a smaller allocation.
  
      // A limit is set for how much we can grow. We should not exceed that
      // (the wasm binary specifies it, so if we tried, we'd fail anyhow).
      var maxHeapSize = getHeapMax();
      if (requestedSize > maxHeapSize) {
        err(`Cannot enlarge memory, requested ${requestedSize} bytes, but the limit is ${maxHeapSize} bytes!`);
        return false;
      }
  
      // Loop through potential heap size increases. If we attempt a too eager
      // reservation that fails, cut down on the attempted size and reserve a
      // smaller bump instead. (max 3 times, chosen somewhat arbitrarily)
      for (var cutDown = 1; cutDown <= 4; cutDown *= 2) {
        var overGrownHeapSize = oldSize * (1 + 0.2 / cutDown); // ensure geometric growth
        // but limit overreserving (default to capping at +96MB overgrowth at most)
        overGrownHeapSize = Math.min(overGrownHeapSize, requestedSize + 100663296 );
  
        var newSize = Math.min(maxHeapSize, alignMemory(Math.max(requestedSize, overGrownHeapSize), 65536));
  
        var replacement = growMemory(newSize);
        if (replacement) {
  
          return true;
        }
      }
      err(`Failed to grow the heap from ${oldSize} bytes to ${newSize} bytes, not enough memory!`);
      return false;
    };

  /** @suppress {checkTypes} */
  var _emscripten_sample_gamepad_data = () => {
      try {
        if (navigator.getGamepads) return (JSEvents.lastGamepadState = navigator.getGamepads())
          ? 0 : -1;
      } catch(e) {
        err(`navigator.getGamepads() exists, but failed to execute with exception ${e}. Disabling Gamepad access.`);
        navigator.getGamepads = null; // Disable getGamepads() so that it won't be attempted to be used again.
      }
      return -1;
    };

  
  
  
  var registerBeforeUnloadEventCallback = (target, userData, useCapture, callbackfunc, eventTypeId, eventTypeString) => {
      var beforeUnloadEventHandlerFunc = (e = event) => {
        // Note: This is always called on the main browser thread, since it needs synchronously return a value!
        var confirmationMessage = getWasmTableEntry(callbackfunc)(eventTypeId, 0, userData);
  
        if (confirmationMessage) {
          confirmationMessage = UTF8ToString(confirmationMessage);
        }
        if (confirmationMessage) {
          e.preventDefault();
          e.returnValue = confirmationMessage;
          return confirmationMessage;
        }
      };
  
      var eventHandler = {
        target: findEventTarget(target),
        eventTypeString,
        callbackfunc,
        handlerFunc: beforeUnloadEventHandlerFunc,
        useCapture
      };
      return JSEvents.registerOrRemoveHandler(eventHandler);
    };
  var _emscripten_set_beforeunload_callback_on_thread = (userData, callbackfunc, targetThread) => {
      if (typeof onbeforeunload == 'undefined') return -1;
      // beforeunload callback can only be registered on the main browser thread, because the page will go away immediately after returning from the handler,
      // and there is no time to start proxying it anywhere.
      if (targetThread !== 1) return -5;
      return registerBeforeUnloadEventCallback(2, userData, true, callbackfunc, 28, "beforeunload");
    };

  
  
  
  
  var registerFocusEventCallback = (target, userData, useCapture, callbackfunc, eventTypeId, eventTypeString, targetThread) => {
      JSEvents.focusEvent ||= _malloc(256);
  
      var focusEventHandlerFunc = (e = event) => {
        var nodeName = JSEvents.getNodeNameForTarget(e.target);
        var id = e.target.id ? e.target.id : '';
  
        var focusEvent = JSEvents.focusEvent;
        stringToUTF8(nodeName, focusEvent + 0, 128);
        stringToUTF8(id, focusEvent + 128, 128);
  
        if (getWasmTableEntry(callbackfunc)(eventTypeId, focusEvent, userData)) e.preventDefault();
      };
  
      var eventHandler = {
        target: findEventTarget(target),
        eventTypeString,
        callbackfunc,
        handlerFunc: focusEventHandlerFunc,
        useCapture
      };
      return JSEvents.registerOrRemoveHandler(eventHandler);
    };
  var _emscripten_set_blur_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerFocusEventCallback(target, userData, useCapture, callbackfunc, 12, "blur", targetThread);


  var _emscripten_set_element_css_size = (target, width, height) => {
      target = findEventTarget(target);
      if (!target) return -4;
  
      target.style.width = width + "px";
      target.style.height = height + "px";
  
      return 0;
    };

  var _emscripten_set_focus_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerFocusEventCallback(target, userData, useCapture, callbackfunc, 13, "focus", targetThread);

  
  
  
  var fillFullscreenChangeEventData = (eventStruct) => {
      var fullscreenElement = document.fullscreenElement || document.mozFullScreenElement || document.webkitFullscreenElement || document.msFullscreenElement;
      var isFullscreen = !!fullscreenElement;
      // Assigning a boolean to HEAP32 with expected type coercion.
      /** @suppress{checkTypes} */
      HEAP8[eventStruct] = isFullscreen;
      HEAP8[(eventStruct)+(1)] = JSEvents.fullscreenEnabled();
      // If transitioning to fullscreen, report info about the element that is now fullscreen.
      // If transitioning to windowed mode, report info about the element that just was fullscreen.
      var reportedElement = isFullscreen ? fullscreenElement : JSEvents.previousFullscreenElement;
      var nodeName = JSEvents.getNodeNameForTarget(reportedElement);
      var id = reportedElement?.id || '';
      stringToUTF8(nodeName, eventStruct + 2, 128);
      stringToUTF8(id, eventStruct + 130, 128);
      HEAP32[(((eventStruct)+(260))>>2)] = reportedElement ? reportedElement.clientWidth : 0;
      HEAP32[(((eventStruct)+(264))>>2)] = reportedElement ? reportedElement.clientHeight : 0;
      HEAP32[(((eventStruct)+(268))>>2)] = screen.width;
      HEAP32[(((eventStruct)+(272))>>2)] = screen.height;
      if (isFullscreen) {
        JSEvents.previousFullscreenElement = fullscreenElement;
      }
    };
  
  
  var registerFullscreenChangeEventCallback = (target, userData, useCapture, callbackfunc, eventTypeId, eventTypeString, targetThread) => {
      JSEvents.fullscreenChangeEvent ||= _malloc(276);
  
      var fullscreenChangeEventhandlerFunc = (e = event) => {
        var fullscreenChangeEvent = JSEvents.fullscreenChangeEvent;
  
        fillFullscreenChangeEventData(fullscreenChangeEvent);
  
        if (getWasmTableEntry(callbackfunc)(eventTypeId, fullscreenChangeEvent, userData)) e.preventDefault();
      };
  
      var eventHandler = {
        target,
        eventTypeString,
        callbackfunc,
        handlerFunc: fullscreenChangeEventhandlerFunc,
        useCapture
      };
      return JSEvents.registerOrRemoveHandler(eventHandler);
    };
  
  var _emscripten_set_fullscreenchange_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) => {
      if (!JSEvents.fullscreenEnabled()) return -1;
      target = findEventTarget(target);
      if (!target) return -4;
  
      // Unprefixed Fullscreen API shipped in Chromium 71 (https://bugs.chromium.org/p/chromium/issues/detail?id=383813)
      // As of Safari 13.0.3 on macOS Catalina 10.15.1 still ships with prefixed webkitfullscreenchange. TODO: revisit this check once Safari ships unprefixed version.
      registerFullscreenChangeEventCallback(target, userData, useCapture, callbackfunc, 19, "webkitfullscreenchange", targetThread);
  
      return registerFullscreenChangeEventCallback(target, userData, useCapture, callbackfunc, 19, "fullscreenchange", targetThread);
    };

  
  
  
  
  var registerGamepadEventCallback = (target, userData, useCapture, callbackfunc, eventTypeId, eventTypeString, targetThread) => {
      JSEvents.gamepadEvent ||= _malloc(1240);
  
      var gamepadEventHandlerFunc = (e = event) => {
        var gamepadEvent = JSEvents.gamepadEvent;
        fillGamepadEventData(gamepadEvent, e["gamepad"]);
  
        if (getWasmTableEntry(callbackfunc)(eventTypeId, gamepadEvent, userData)) e.preventDefault();
      };
  
      var eventHandler = {
        target: findEventTarget(target),
        allowsDeferredCalls: true,
        eventTypeString,
        callbackfunc,
        handlerFunc: gamepadEventHandlerFunc,
        useCapture
      };
      return JSEvents.registerOrRemoveHandler(eventHandler);
    };
  
  var _emscripten_set_gamepadconnected_callback_on_thread = (userData, useCapture, callbackfunc, targetThread) => {
      if (_emscripten_sample_gamepad_data()) return -1;
      return registerGamepadEventCallback(2, userData, useCapture, callbackfunc, 26, "gamepadconnected", targetThread);
    };

  
  var _emscripten_set_gamepaddisconnected_callback_on_thread = (userData, useCapture, callbackfunc, targetThread) => {
      if (_emscripten_sample_gamepad_data()) return -1;
      return registerGamepadEventCallback(2, userData, useCapture, callbackfunc, 27, "gamepaddisconnected", targetThread);
    };

  
  
  
  
  var registerKeyEventCallback = (target, userData, useCapture, callbackfunc, eventTypeId, eventTypeString, targetThread) => {
      JSEvents.keyEvent ||= _malloc(160);
  
      var keyEventHandlerFunc = (e) => {
        assert(e);
  
        var keyEventData = JSEvents.keyEvent;
        HEAPF64[((keyEventData)>>3)] = e.timeStamp;
  
        var idx = ((keyEventData)>>2);
  
        HEAP32[idx + 2] = e.location;
        HEAP8[keyEventData + 12] = e.ctrlKey;
        HEAP8[keyEventData + 13] = e.shiftKey;
        HEAP8[keyEventData + 14] = e.altKey;
        HEAP8[keyEventData + 15] = e.metaKey;
        HEAP8[keyEventData + 16] = e.repeat;
        HEAP32[idx + 5] = e.charCode;
        HEAP32[idx + 6] = e.keyCode;
        HEAP32[idx + 7] = e.which;
        stringToUTF8(e.key || '', keyEventData + 32, 32);
        stringToUTF8(e.code || '', keyEventData + 64, 32);
        stringToUTF8(e.char || '', keyEventData + 96, 32);
        stringToUTF8(e.locale || '', keyEventData + 128, 32);
  
        if (getWasmTableEntry(callbackfunc)(eventTypeId, keyEventData, userData)) e.preventDefault();
      };
  
      var eventHandler = {
        target: findEventTarget(target),
        eventTypeString,
        callbackfunc,
        handlerFunc: keyEventHandlerFunc,
        useCapture
      };
      return JSEvents.registerOrRemoveHandler(eventHandler);
    };
  var _emscripten_set_keydown_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerKeyEventCallback(target, userData, useCapture, callbackfunc, 2, "keydown", targetThread);

  var _emscripten_set_keypress_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerKeyEventCallback(target, userData, useCapture, callbackfunc, 1, "keypress", targetThread);

  var _emscripten_set_keyup_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerKeyEventCallback(target, userData, useCapture, callbackfunc, 3, "keyup", targetThread);

  
  var _emscripten_set_main_loop = (func, fps, simulateInfiniteLoop) => {
      var iterFunc = getWasmTableEntry(func);
      setMainLoop(iterFunc, fps, simulateInfiniteLoop);
    };


  
  var fillMouseEventData = (eventStruct, e, target) => {
      assert(eventStruct % 4 == 0);
      HEAPF64[((eventStruct)>>3)] = e.timeStamp;
      var idx = ((eventStruct)>>2);
      HEAP32[idx + 2] = e.screenX;
      HEAP32[idx + 3] = e.screenY;
      HEAP32[idx + 4] = e.clientX;
      HEAP32[idx + 5] = e.clientY;
      HEAP8[eventStruct + 24] = e.ctrlKey;
      HEAP8[eventStruct + 25] = e.shiftKey;
      HEAP8[eventStruct + 26] = e.altKey;
      HEAP8[eventStruct + 27] = e.metaKey;
      HEAP16[idx*2 + 14] = e.button;
      HEAP16[idx*2 + 15] = e.buttons;
  
      HEAP32[idx + 8] = e["movementX"];
  
      HEAP32[idx + 9] = e["movementY"];
  
      // Note: rect contains doubles (truncated to placate SAFE_HEAP, which is the same behaviour when writing to HEAP32 anyway)
      var rect = getBoundingClientRect(target);
      HEAP32[idx + 10] = e.clientX - (rect.left | 0);
      HEAP32[idx + 11] = e.clientY - (rect.top  | 0);
    };
  
  
  
  var registerMouseEventCallback = (target, userData, useCapture, callbackfunc, eventTypeId, eventTypeString, targetThread) => {
      JSEvents.mouseEvent ||= _malloc(64);
      target = findEventTarget(target);
  
      var mouseEventHandlerFunc = (e = event) => {
        // TODO: Make this access thread safe, or this could update live while app is reading it.
        fillMouseEventData(JSEvents.mouseEvent, e, target);
  
        if (getWasmTableEntry(callbackfunc)(eventTypeId, JSEvents.mouseEvent, userData)) e.preventDefault();
      };
  
      var eventHandler = {
        target,
        allowsDeferredCalls: eventTypeString != 'mousemove' && eventTypeString != 'mouseenter' && eventTypeString != 'mouseleave', // Mouse move events do not allow fullscreen/pointer lock requests to be handled in them!
        eventTypeString,
        callbackfunc,
        handlerFunc: mouseEventHandlerFunc,
        useCapture
      };
      return JSEvents.registerOrRemoveHandler(eventHandler);
    };
  var _emscripten_set_mousedown_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerMouseEventCallback(target, userData, useCapture, callbackfunc, 5, "mousedown", targetThread);

  var _emscripten_set_mouseenter_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerMouseEventCallback(target, userData, useCapture, callbackfunc, 33, "mouseenter", targetThread);

  var _emscripten_set_mouseleave_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerMouseEventCallback(target, userData, useCapture, callbackfunc, 34, "mouseleave", targetThread);

  var _emscripten_set_mousemove_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerMouseEventCallback(target, userData, useCapture, callbackfunc, 8, "mousemove", targetThread);

  var _emscripten_set_mouseup_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerMouseEventCallback(target, userData, useCapture, callbackfunc, 6, "mouseup", targetThread);

  
  var screenOrientation = () => {
      if (!window.screen) return undefined;
      return screen.orientation || screen['mozOrientation'] || screen['webkitOrientation'];
    };
  var fillOrientationChangeEventData = (eventStruct) => {
      // OrientationType enum
      var orientationsType1 = ['portrait-primary', 'portrait-secondary', 'landscape-primary', 'landscape-secondary'];
      // alternative selection from OrientationLockType enum
      var orientationsType2 = ['portrait',         'portrait',           'landscape',         'landscape'];
  
      var orientationIndex = 0;
      var orientationAngle = 0;
      var screenOrientObj  = screenOrientation();
      if (typeof screenOrientObj === 'object') {
        orientationIndex = orientationsType1.indexOf(screenOrientObj.type);
        if (orientationIndex < 0) {
          orientationIndex = orientationsType2.indexOf(screenOrientObj.type);
        }
        if (orientationIndex >= 0) {
          orientationIndex = 1 << orientationIndex;
        }
        orientationAngle = screenOrientObj.angle;
      }
      else {
        // fallback for Safari earlier than 16.4 (March 2023)
        orientationAngle = window.orientation;
      }
  
      HEAP32[((eventStruct)>>2)] = orientationIndex;
      HEAP32[(((eventStruct)+(4))>>2)] = orientationAngle;
    };
  
  
  var registerOrientationChangeEventCallback = (target, userData, useCapture, callbackfunc, eventTypeId, eventTypeString, targetThread) => {
      JSEvents.orientationChangeEvent ||= _malloc(8);
  
      var orientationChangeEventHandlerFunc = (e = event) => {
        var orientationChangeEvent = JSEvents.orientationChangeEvent;
  
        fillOrientationChangeEventData(orientationChangeEvent);
  
        if (getWasmTableEntry(callbackfunc)(eventTypeId, orientationChangeEvent, userData)) e.preventDefault();
      };
  
      var eventHandler = {
        target,
        eventTypeString,
        callbackfunc,
        handlerFunc: orientationChangeEventHandlerFunc,
        useCapture
      };
      return JSEvents.registerOrRemoveHandler(eventHandler);
    };
  var _emscripten_set_orientationchange_callback_on_thread = (userData, useCapture, callbackfunc, targetThread) => {
      if (!window.screen || !screen.orientation) return -1;
      return registerOrientationChangeEventCallback(screen.orientation, userData, useCapture, callbackfunc, 18, 'change', targetThread);
    };

  
  
  var fillPointerlockChangeEventData = (eventStruct) => {
      var pointerLockElement = document.pointerLockElement || document.mozPointerLockElement || document.webkitPointerLockElement || document.msPointerLockElement;
      var isPointerlocked = !!pointerLockElement;
      // Assigning a boolean to HEAP32 with expected type coercion.
      /** @suppress{checkTypes} */
      HEAP8[eventStruct] = isPointerlocked;
      var nodeName = JSEvents.getNodeNameForTarget(pointerLockElement);
      var id = pointerLockElement?.id || '';
      stringToUTF8(nodeName, eventStruct + 1, 128);
      stringToUTF8(id, eventStruct + 129, 128);
    };
  
  
  var registerPointerlockChangeEventCallback = (target, userData, useCapture, callbackfunc, eventTypeId, eventTypeString, targetThread) => {
      JSEvents.pointerlockChangeEvent ||= _malloc(257);
  
      var pointerlockChangeEventHandlerFunc = (e = event) => {
        var pointerlockChangeEvent = JSEvents.pointerlockChangeEvent;
        fillPointerlockChangeEventData(pointerlockChangeEvent);
  
        if (getWasmTableEntry(callbackfunc)(eventTypeId, pointerlockChangeEvent, userData)) e.preventDefault();
      };
  
      var eventHandler = {
        target,
        eventTypeString,
        callbackfunc,
        handlerFunc: pointerlockChangeEventHandlerFunc,
        useCapture
      };
      return JSEvents.registerOrRemoveHandler(eventHandler);
    };
  
  /** @suppress {missingProperties} */
  var _emscripten_set_pointerlockchange_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) => {
      // TODO: Currently not supported in pthreads or in --proxy-to-worker mode. (In pthreads mode, document object is not defined)
      if (!document || !document.body || (!document.body.requestPointerLock && !document.body.mozRequestPointerLock && !document.body.webkitRequestPointerLock && !document.body.msRequestPointerLock)) {
        return -1;
      }
  
      target = findEventTarget(target);
      if (!target) return -4;
      registerPointerlockChangeEventCallback(target, userData, useCapture, callbackfunc, 20, "mozpointerlockchange", targetThread);
      registerPointerlockChangeEventCallback(target, userData, useCapture, callbackfunc, 20, "webkitpointerlockchange", targetThread);
      registerPointerlockChangeEventCallback(target, userData, useCapture, callbackfunc, 20, "mspointerlockchange", targetThread);
      return registerPointerlockChangeEventCallback(target, userData, useCapture, callbackfunc, 20, "pointerlockchange", targetThread);
    };

  
  
  
  var registerUiEventCallback = (target, userData, useCapture, callbackfunc, eventTypeId, eventTypeString, targetThread) => {
      JSEvents.uiEvent ||= _malloc(36);
  
      target = findEventTarget(target);
  
      var uiEventHandlerFunc = (e = event) => {
        if (e.target != target) {
          // Never take ui events such as scroll via a 'bubbled' route, but always from the direct element that
          // was targeted. Otherwise e.g. if app logs a message in response to a page scroll, the Emscripten log
          // message box could cause to scroll, generating a new (bubbled) scroll message, causing a new log print,
          // causing a new scroll, etc..
          return;
        }
        var b = document.body; // Take document.body to a variable, Closure compiler does not outline access to it on its own.
        if (!b) {
          // During a page unload 'body' can be null, with "Cannot read property 'clientWidth' of null" being thrown
          return;
        }
        var uiEvent = JSEvents.uiEvent;
        HEAP32[((uiEvent)>>2)] = 0; // always zero for resize and scroll
        HEAP32[(((uiEvent)+(4))>>2)] = b.clientWidth;
        HEAP32[(((uiEvent)+(8))>>2)] = b.clientHeight;
        HEAP32[(((uiEvent)+(12))>>2)] = innerWidth;
        HEAP32[(((uiEvent)+(16))>>2)] = innerHeight;
        HEAP32[(((uiEvent)+(20))>>2)] = outerWidth;
        HEAP32[(((uiEvent)+(24))>>2)] = outerHeight;
        HEAP32[(((uiEvent)+(28))>>2)] = pageXOffset | 0; // scroll offsets are float
        HEAP32[(((uiEvent)+(32))>>2)] = pageYOffset | 0;
        if (getWasmTableEntry(callbackfunc)(eventTypeId, uiEvent, userData)) e.preventDefault();
      };
  
      var eventHandler = {
        target,
        eventTypeString,
        callbackfunc,
        handlerFunc: uiEventHandlerFunc,
        useCapture
      };
      return JSEvents.registerOrRemoveHandler(eventHandler);
    };
  var _emscripten_set_resize_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerUiEventCallback(target, userData, useCapture, callbackfunc, 10, "resize", targetThread);

  
  var _emscripten_set_timeout = (cb, msecs, userData) =>
      safeSetTimeout(() => getWasmTableEntry(cb)(userData), msecs);

  
  
  
  
  var registerTouchEventCallback = (target, userData, useCapture, callbackfunc, eventTypeId, eventTypeString, targetThread) => {
      JSEvents.touchEvent ||= _malloc(1552);
  
      target = findEventTarget(target);
  
      var touchEventHandlerFunc = (e) => {
        assert(e);
        var t, touches = {}, et = e.touches;
        // To ease marshalling different kinds of touches that browser reports (all touches are listed in e.touches,
        // only changed touches in e.changedTouches, and touches on target at a.targetTouches), mark a boolean in
        // each Touch object so that we can later loop only once over all touches we see to marshall over to Wasm.
  
        for (let t of et) {
          // Browser might recycle the generated Touch objects between each frame (Firefox on Android), so reset any
          // changed/target states we may have set from previous frame.
          t.isChanged = t.onTarget = 0;
          touches[t.identifier] = t;
        }
        // Mark which touches are part of the changedTouches list.
        for (let t of e.changedTouches) {
          t.isChanged = 1;
          touches[t.identifier] = t;
        }
        // Mark which touches are part of the targetTouches list.
        for (let t of e.targetTouches) {
          touches[t.identifier].onTarget = 1;
        }
  
        var touchEvent = JSEvents.touchEvent;
        HEAPF64[((touchEvent)>>3)] = e.timeStamp;
        HEAP8[touchEvent + 12] = e.ctrlKey;
        HEAP8[touchEvent + 13] = e.shiftKey;
        HEAP8[touchEvent + 14] = e.altKey;
        HEAP8[touchEvent + 15] = e.metaKey;
        var idx = touchEvent + 16;
        var targetRect = getBoundingClientRect(target);
        var numTouches = 0;
        for (let t of Object.values(touches)) {
          var idx32 = ((idx)>>2); // Pre-shift the ptr to index to HEAP32 to save code size
          HEAP32[idx32 + 0] = t.identifier;
          HEAP32[idx32 + 1] = t.screenX;
          HEAP32[idx32 + 2] = t.screenY;
          HEAP32[idx32 + 3] = t.clientX;
          HEAP32[idx32 + 4] = t.clientY;
          HEAP32[idx32 + 5] = t.pageX;
          HEAP32[idx32 + 6] = t.pageY;
          HEAP8[idx + 28] = t.isChanged;
          HEAP8[idx + 29] = t.onTarget;
          HEAP32[idx32 + 8] = t.clientX - (targetRect.left | 0);
          HEAP32[idx32 + 9] = t.clientY - (targetRect.top  | 0);
  
          idx += 48;
  
          if (++numTouches > 31) {
            break;
          }
        }
        HEAP32[(((touchEvent)+(8))>>2)] = numTouches;
  
        if (getWasmTableEntry(callbackfunc)(eventTypeId, touchEvent, userData)) e.preventDefault();
      };
  
      var eventHandler = {
        target,
        allowsDeferredCalls: eventTypeString == 'touchstart' || eventTypeString == 'touchend',
        eventTypeString,
        callbackfunc,
        handlerFunc: touchEventHandlerFunc,
        useCapture
      };
      return JSEvents.registerOrRemoveHandler(eventHandler);
    };
  var _emscripten_set_touchcancel_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerTouchEventCallback(target, userData, useCapture, callbackfunc, 25, "touchcancel", targetThread);

  var _emscripten_set_touchend_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerTouchEventCallback(target, userData, useCapture, callbackfunc, 23, "touchend", targetThread);

  var _emscripten_set_touchmove_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerTouchEventCallback(target, userData, useCapture, callbackfunc, 24, "touchmove", targetThread);

  var _emscripten_set_touchstart_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) =>
      registerTouchEventCallback(target, userData, useCapture, callbackfunc, 22, "touchstart", targetThread);

  
  var fillVisibilityChangeEventData = (eventStruct) => {
      var visibilityStates = [ "hidden", "visible", "prerender", "unloaded" ];
      var visibilityState = visibilityStates.indexOf(document.visibilityState);
  
      // Assigning a boolean to HEAP32 with expected type coercion.
      /** @suppress{checkTypes} */
      HEAP8[eventStruct] = document.hidden;
      HEAP32[(((eventStruct)+(4))>>2)] = visibilityState;
    };
  
  
  var registerVisibilityChangeEventCallback = (target, userData, useCapture, callbackfunc, eventTypeId, eventTypeString, targetThread) => {
      JSEvents.visibilityChangeEvent ||= _malloc(8);
  
      var visibilityChangeEventHandlerFunc = (e = event) => {
        var visibilityChangeEvent = JSEvents.visibilityChangeEvent;
  
        fillVisibilityChangeEventData(visibilityChangeEvent);
  
        if (getWasmTableEntry(callbackfunc)(eventTypeId, visibilityChangeEvent, userData)) e.preventDefault();
      };
  
      var eventHandler = {
        target,
        eventTypeString,
        callbackfunc,
        handlerFunc: visibilityChangeEventHandlerFunc,
        useCapture
      };
      return JSEvents.registerOrRemoveHandler(eventHandler);
    };
  
  var _emscripten_set_visibilitychange_callback_on_thread = (userData, useCapture, callbackfunc, targetThread) => {
    if (!specialHTMLTargets[1]) {
      return -4;
    }
      return registerVisibilityChangeEventCallback(specialHTMLTargets[1], userData, useCapture, callbackfunc, 21, "visibilitychange", targetThread);
    };

  
  
  
  var registerWheelEventCallback = (target, userData, useCapture, callbackfunc, eventTypeId, eventTypeString, targetThread) => {
      JSEvents.wheelEvent ||= _malloc(96);
  
      // The DOM Level 3 events spec event 'wheel'
      var wheelHandlerFunc = (e = event) => {
        var wheelEvent = JSEvents.wheelEvent;
        fillMouseEventData(wheelEvent, e, target);
        HEAPF64[(((wheelEvent)+(64))>>3)] = e["deltaX"];
        HEAPF64[(((wheelEvent)+(72))>>3)] = e["deltaY"];
        HEAPF64[(((wheelEvent)+(80))>>3)] = e["deltaZ"];
        HEAP32[(((wheelEvent)+(88))>>2)] = e["deltaMode"];
        if (getWasmTableEntry(callbackfunc)(eventTypeId, wheelEvent, userData)) e.preventDefault();
      };
  
      var eventHandler = {
        target,
        allowsDeferredCalls: true,
        eventTypeString,
        callbackfunc,
        handlerFunc: wheelHandlerFunc,
        useCapture
      };
      return JSEvents.registerOrRemoveHandler(eventHandler);
    };
  
  var _emscripten_set_wheel_callback_on_thread = (target, userData, useCapture, callbackfunc, targetThread) => {
      target = findEventTarget(target);
      if (!target) return -4;
      if (typeof target.onwheel != 'undefined') {
        return registerWheelEventCallback(target, userData, useCapture, callbackfunc, 9, "wheel", targetThread);
      } else {
        return -1;
      }
    };

  
  var _emscripten_set_window_title = (title) => document.title = UTF8ToString(title);

  var _emscripten_sleep = () => {
      throw 'Please compile your program with async support in order to use asynchronous operations like emscripten_sleep';
    };

  
  var webglPowerPreferences = ["default","low-power","high-performance"];
  
  
  /** @suppress {duplicate } */
  var _emscripten_webgl_do_create_context = (target, attributes) => {
      assert(attributes);
      var attr32 = ((attributes)>>2);
      var powerPreference = HEAP32[attr32 + (8>>2)];
      var contextAttributes = {
        'alpha': !!HEAP8[attributes + 0],
        'depth': !!HEAP8[attributes + 1],
        'stencil': !!HEAP8[attributes + 2],
        'antialias': !!HEAP8[attributes + 3],
        'premultipliedAlpha': !!HEAP8[attributes + 4],
        'preserveDrawingBuffer': !!HEAP8[attributes + 5],
        'powerPreference': webglPowerPreferences[powerPreference],
        'failIfMajorPerformanceCaveat': !!HEAP8[attributes + 12],
        // The following are not predefined WebGL context attributes in the WebGL specification, so the property names can be minified by Closure.
        majorVersion: HEAP32[attr32 + (16>>2)],
        minorVersion: HEAP32[attr32 + (20>>2)],
        enableExtensionsByDefault: HEAP8[attributes + 24],
        explicitSwapControl: HEAP8[attributes + 25],
        proxyContextToMainThread: HEAP32[attr32 + (28>>2)],
        renderViaOffscreenBackBuffer: HEAP8[attributes + 32]
      };
  
      //  TODO: Make these into hard errors at some point in the future
      if (contextAttributes.majorVersion !== 1 && contextAttributes.majorVersion !== 2) {
        err(`Invalid WebGL version requested: ${contextAttributes.majorVersion}`);
      }
      if (contextAttributes.majorVersion !== 1) {
        err('WebGL 2 requested but only WebGL 1 is supported (set -sMAX_WEBGL_VERSION=2 to fix the problem)');
      }
  
      var canvas = findCanvasEventTarget(target);
  
      if (!canvas) {
        return 0;
      }
  
      if (contextAttributes.explicitSwapControl) {
        return 0;
      }
  
      var contextHandle = GL.createContext(canvas, contextAttributes);
      return contextHandle;
    };
  var _emscripten_webgl_create_context = _emscripten_webgl_do_create_context;

  var _emscripten_webgl_destroy_context = (contextHandle) => {
      if (GL.currentContext == contextHandle) GL.currentContext = 0;
      GL.deleteContext(contextHandle);
    };

  var _emscripten_webgl_make_context_current = (contextHandle) => {
      var success = GL.makeContextCurrent(contextHandle);
      return success ? 0 : -5;
    };

  var ENV = {
  };
  
  var getEnvStrings = () => {
      if (!getEnvStrings.strings) {
        // Default values.
        // Browser language detection #8751
        var lang = ((typeof navigator == 'object' && navigator.language) || 'C').replace('-', '_') + '.UTF-8';
        var env = {
          'USER': 'web_user',
          'LOGNAME': 'web_user',
          'PATH': '/',
          'PWD': '/',
          'HOME': '/home/web_user',
          'LANG': lang,
          '_': getExecutableName()
        };
        // Apply the user-provided values, if any.
        for (var x in ENV) {
          // x is a key in ENV; if ENV[x] is undefined, that means it was
          // explicitly set to be so. We allow user code to do that to
          // force variables with default values to remain unset.
          if (ENV[x] === undefined) delete env[x];
          else env[x] = ENV[x];
        }
        var strings = [];
        for (var x in env) {
          strings.push(`${x}=${env[x]}`);
        }
        getEnvStrings.strings = strings;
      }
      return getEnvStrings.strings;
    };
  
  var _environ_get = (__environ, environ_buf) => {
      var bufSize = 0;
      var envp = 0;
      for (var string of getEnvStrings()) {
        var ptr = environ_buf + bufSize;
        HEAPU32[(((__environ)+(envp))>>2)] = ptr;
        bufSize += stringToUTF8(string, ptr, Infinity) + 1;
        envp += 4;
      }
      return 0;
    };

  
  var _environ_sizes_get = (penviron_count, penviron_buf_size) => {
      var strings = getEnvStrings();
      HEAPU32[((penviron_count)>>2)] = strings.length;
      var bufSize = 0;
      for (var string of strings) {
        bufSize += lengthBytesUTF8(string) + 1;
      }
      HEAPU32[((penviron_buf_size)>>2)] = bufSize;
      return 0;
    };


  function _fd_close(fd) {
  try {
  
      var stream = SYSCALLS.getStreamFromFD(fd);
      FS.close(stream);
      return 0;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return e.errno;
  }
  }

  /** @param {number=} offset */
  var doReadv = (stream, iov, iovcnt, offset) => {
      var ret = 0;
      for (var i = 0; i < iovcnt; i++) {
        var ptr = HEAPU32[((iov)>>2)];
        var len = HEAPU32[(((iov)+(4))>>2)];
        iov += 8;
        var curr = FS.read(stream, HEAP8, ptr, len, offset);
        if (curr < 0) return -1;
        ret += curr;
        if (curr < len) break; // nothing more to read
        if (typeof offset != 'undefined') {
          offset += curr;
        }
      }
      return ret;
    };
  
  function _fd_read(fd, iov, iovcnt, pnum) {
  try {
  
      var stream = SYSCALLS.getStreamFromFD(fd);
      var num = doReadv(stream, iov, iovcnt);
      HEAPU32[((pnum)>>2)] = num;
      return 0;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return e.errno;
  }
  }

  
  function _fd_seek(fd, offset, whence, newOffset) {
    offset = bigintToI53Checked(offset);
  
  
  try {
  
      if (isNaN(offset)) return 61;
      var stream = SYSCALLS.getStreamFromFD(fd);
      FS.llseek(stream, offset, whence);
      HEAP64[((newOffset)>>3)] = BigInt(stream.position);
      if (stream.getdents && offset === 0 && whence === 0) stream.getdents = null; // reset readdir state
      return 0;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return e.errno;
  }
  ;
  }

  /** @param {number=} offset */
  var doWritev = (stream, iov, iovcnt, offset) => {
      var ret = 0;
      for (var i = 0; i < iovcnt; i++) {
        var ptr = HEAPU32[((iov)>>2)];
        var len = HEAPU32[(((iov)+(4))>>2)];
        iov += 8;
        var curr = FS.write(stream, HEAP8, ptr, len, offset);
        if (curr < 0) return -1;
        ret += curr;
        if (curr < len) {
          // No more space to write.
          break;
        }
        if (typeof offset != 'undefined') {
          offset += curr;
        }
      }
      return ret;
    };
  
  function _fd_write(fd, iov, iovcnt, pnum) {
  try {
  
      var stream = SYSCALLS.getStreamFromFD(fd);
      var num = doWritev(stream, iov, iovcnt);
      HEAPU32[((pnum)>>2)] = num;
      return 0;
    } catch (e) {
    if (typeof FS == 'undefined' || !(e.name === 'ErrnoError')) throw e;
    return e.errno;
  }
  }


  var dynCall = (sig, ptr, args = [], promising = false) => {
      assert(!promising, 'async dynCall is not supported in this mode')
      assert(getWasmTableEntry(ptr), `missing table entry in dynCall: ${ptr}`);
      var func = getWasmTableEntry(ptr);
      var rtn = func(...args);
  
      function convert(rtn) {
        return rtn;
      }
  
      return convert(rtn);
    };



  /** @param {Object=} elements */
  var autoResumeAudioContext = (ctx, elements) => {
      if (!elements) {
        elements = [document, document.getElementById('canvas')];
      }
      ['keydown', 'mousedown', 'touchstart'].forEach((event) => {
        elements.forEach((element) => {
          element?.addEventListener(event, () => {
            if (ctx.state === 'suspended') ctx.resume();
          }, { 'once': true });
        });
      });
    };



  var writeArrayToMemory = (array, buffer) => {
      assert(array.length >= 0, 'writeArrayToMemory array must have a length (should be an array or typed array)')
      HEAP8.set(array, buffer);
    };





  
  var allocateUTF8 = stringToNewUTF8;














































































































  FS.createPreloadedFile = FS_createPreloadedFile;
  FS.staticInit();;

      Module['requestAnimationFrame'] = MainLoop.requestAnimationFrame;
      Module['pauseMainLoop'] = MainLoop.pause;
      Module['resumeMainLoop'] = MainLoop.resume;
      MainLoop.init();;
for (let i = 0; i < 32; ++i) tempFixedLengthArray.push(new Array(i));;
var miniTempWebGLFloatBuffersStorage = new Float32Array(288);
  // Create GL_POOL_TEMP_BUFFERS_SIZE+1 temporary buffers, for uploads of size 0 through GL_POOL_TEMP_BUFFERS_SIZE inclusive
  for (/**@suppress{duplicate}*/var i = 0; i <= 288; ++i) {
    miniTempWebGLFloatBuffers[i] = miniTempWebGLFloatBuffersStorage.subarray(0, i);
  };
var miniTempWebGLIntBuffersStorage = new Int32Array(288);
  // Create GL_POOL_TEMP_BUFFERS_SIZE+1 temporary buffers, for uploads of size 0 through GL_POOL_TEMP_BUFFERS_SIZE inclusive
  for (/**@suppress{duplicate}*/var i = 0; i <= 288; ++i) {
    miniTempWebGLIntBuffers[i] = miniTempWebGLIntBuffersStorage.subarray(0, i);
  };
// End JS library code

// include: postlibrary.js
// This file is included after the automatically-generated JS library code
// but before the wasm module is created.

{

  // Begin ATMODULES hooks
  if (Module['noExitRuntime']) noExitRuntime = Module['noExitRuntime'];
if (Module['preloadPlugins']) preloadPlugins = Module['preloadPlugins'];
if (Module['print']) out = Module['print'];
if (Module['printErr']) err = Module['printErr'];
if (Module['wasmBinary']) wasmBinary = Module['wasmBinary'];
  // End ATMODULES hooks

  checkIncomingModuleAPI();

  if (Module['arguments']) arguments_ = Module['arguments'];
  if (Module['thisProgram']) thisProgram = Module['thisProgram'];

  // Assertions on removed incoming Module JS APIs.
  assert(typeof Module['memoryInitializerPrefixURL'] == 'undefined', 'Module.memoryInitializerPrefixURL option was removed, use Module.locateFile instead');
  assert(typeof Module['pthreadMainPrefixURL'] == 'undefined', 'Module.pthreadMainPrefixURL option was removed, use Module.locateFile instead');
  assert(typeof Module['cdInitializerPrefixURL'] == 'undefined', 'Module.cdInitializerPrefixURL option was removed, use Module.locateFile instead');
  assert(typeof Module['filePackagePrefixURL'] == 'undefined', 'Module.filePackagePrefixURL option was removed, use Module.locateFile instead');
  assert(typeof Module['read'] == 'undefined', 'Module.read option was removed');
  assert(typeof Module['readAsync'] == 'undefined', 'Module.readAsync option was removed (modify readAsync in JS)');
  assert(typeof Module['readBinary'] == 'undefined', 'Module.readBinary option was removed (modify readBinary in JS)');
  assert(typeof Module['setWindowTitle'] == 'undefined', 'Module.setWindowTitle option was removed (modify emscripten_set_window_title in JS)');
  assert(typeof Module['TOTAL_MEMORY'] == 'undefined', 'Module.TOTAL_MEMORY has been renamed Module.INITIAL_MEMORY');
  assert(typeof Module['ENVIRONMENT'] == 'undefined', 'Module.ENVIRONMENT has been deprecated. To force the environment, use the ENVIRONMENT compile-time option (for example, -sENVIRONMENT=web or -sENVIRONMENT=node)');
  assert(typeof Module['STACK_SIZE'] == 'undefined', 'STACK_SIZE can no longer be set at runtime.  Use -sSTACK_SIZE at link time')
  // If memory is defined in wasm, the user can't provide it, or set INITIAL_MEMORY
  assert(typeof Module['wasmMemory'] == 'undefined', 'Use of `wasmMemory` detected.  Use -sIMPORTED_MEMORY to define wasmMemory externally');
  assert(typeof Module['INITIAL_MEMORY'] == 'undefined', 'Detected runtime INITIAL_MEMORY setting.  Use -sIMPORTED_MEMORY to define wasmMemory dynamically');

}

// Begin runtime exports
  Module['addFunction'] = addFunction;
  Module['setValue'] = setValue;
  Module['getValue'] = getValue;
  Module['UTF8ToString'] = UTF8ToString;
  Module['stringToUTF8'] = stringToUTF8;
  Module['allocateUTF8'] = allocateUTF8;
  // End runtime exports
  // Begin JS library exports
  Module['ExitStatus'] = ExitStatus;
  Module['addFunction'] = addFunction;
  Module['convertJsFunctionToWasm'] = convertJsFunctionToWasm;
  Module['uleb128EncodeWithLen'] = uleb128EncodeWithLen;
  Module['generateTypePack'] = generateTypePack;
  Module['wasmTypeCodes'] = wasmTypeCodes;
  Module['getFunctionAddress'] = getFunctionAddress;
  Module['updateTableMap'] = updateTableMap;
  Module['getWasmTableEntry'] = getWasmTableEntry;
  Module['wasmTableMirror'] = wasmTableMirror;
  Module['wasmTable'] = wasmTable;
  Module['functionsInTableMap'] = functionsInTableMap;
  Module['getEmptyTableSlot'] = getEmptyTableSlot;
  Module['freeTableIndexes'] = freeTableIndexes;
  Module['setWasmTableEntry'] = setWasmTableEntry;
  Module['addOnPostRun'] = addOnPostRun;
  Module['onPostRuns'] = onPostRuns;
  Module['callRuntimeCallbacks'] = callRuntimeCallbacks;
  Module['addOnPreRun'] = addOnPreRun;
  Module['onPreRuns'] = onPreRuns;
  Module['getValue'] = getValue;
  Module['noExitRuntime'] = noExitRuntime;
  Module['ptrToString'] = ptrToString;
  Module['setValue'] = setValue;
  Module['stackRestore'] = stackRestore;
  Module['stackSave'] = stackSave;
  Module['warnOnce'] = warnOnce;
  Module['___call_sighandler'] = ___call_sighandler;
  Module['___syscall_fcntl64'] = ___syscall_fcntl64;
  Module['syscallGetVarargP'] = syscallGetVarargP;
  Module['syscallGetVarargI'] = syscallGetVarargI;
  Module['SYSCALLS'] = SYSCALLS;
  Module['PATH'] = PATH;
  Module['FS'] = FS;
  Module['randomFill'] = randomFill;
  Module['initRandomFill'] = initRandomFill;
  Module['PATH_FS'] = PATH_FS;
  Module['TTY'] = TTY;
  Module['UTF8ArrayToString'] = UTF8ArrayToString;
  Module['UTF8Decoder'] = UTF8Decoder;
  Module['findStringEnd'] = findStringEnd;
  Module['FS_stdin_getChar'] = FS_stdin_getChar;
  Module['FS_stdin_getChar_buffer'] = FS_stdin_getChar_buffer;
  Module['intArrayFromString'] = intArrayFromString;
  Module['lengthBytesUTF8'] = lengthBytesUTF8;
  Module['stringToUTF8Array'] = stringToUTF8Array;
  Module['MEMFS'] = MEMFS;
  Module['mmapAlloc'] = mmapAlloc;
  Module['zeroMemory'] = zeroMemory;
  Module['alignMemory'] = alignMemory;
  Module['FS_createPreloadedFile'] = FS_createPreloadedFile;
  Module['asyncLoad'] = asyncLoad;
  Module['FS_createDataFile'] = FS_createDataFile;
  Module['getUniqueRunDependency'] = getUniqueRunDependency;
  Module['FS_handledByPreloadPlugin'] = FS_handledByPreloadPlugin;
  Module['preloadPlugins'] = preloadPlugins;
  Module['FS_modeStringToFlags'] = FS_modeStringToFlags;
  Module['FS_getMode'] = FS_getMode;
  Module['strError'] = strError;
  Module['UTF8ToString'] = UTF8ToString;
  Module['ERRNO_CODES'] = ERRNO_CODES;
  Module['___syscall_fdatasync'] = ___syscall_fdatasync;
  Module['___syscall_fstat64'] = ___syscall_fstat64;
  Module['___syscall_ftruncate64'] = ___syscall_ftruncate64;
  Module['bigintToI53Checked'] = bigintToI53Checked;
  Module['INT53_MAX'] = INT53_MAX;
  Module['INT53_MIN'] = INT53_MIN;
  Module['___syscall_getcwd'] = ___syscall_getcwd;
  Module['stringToUTF8'] = stringToUTF8;
  Module['___syscall_getdents64'] = ___syscall_getdents64;
  Module['___syscall_ioctl'] = ___syscall_ioctl;
  Module['___syscall_lstat64'] = ___syscall_lstat64;
  Module['___syscall_mkdirat'] = ___syscall_mkdirat;
  Module['___syscall_newfstatat'] = ___syscall_newfstatat;
  Module['___syscall_openat'] = ___syscall_openat;
  Module['___syscall_renameat'] = ___syscall_renameat;
  Module['___syscall_rmdir'] = ___syscall_rmdir;
  Module['___syscall_stat64'] = ___syscall_stat64;
  Module['___syscall_unlinkat'] = ___syscall_unlinkat;
  Module['__abort_js'] = __abort_js;
  Module['__emscripten_get_progname'] = __emscripten_get_progname;
  Module['getExecutableName'] = getExecutableName;
  Module['__emscripten_runtime_keepalive_clear'] = __emscripten_runtime_keepalive_clear;
  Module['runtimeKeepaliveCounter'] = runtimeKeepaliveCounter;
  Module['__emscripten_throw_longjmp'] = __emscripten_throw_longjmp;
  Module['__gmtime_js'] = __gmtime_js;
  Module['__localtime_js'] = __localtime_js;
  Module['ydayFromDate'] = ydayFromDate;
  Module['isLeapYear'] = isLeapYear;
  Module['MONTH_DAYS_LEAP_CUMULATIVE'] = MONTH_DAYS_LEAP_CUMULATIVE;
  Module['MONTH_DAYS_REGULAR_CUMULATIVE'] = MONTH_DAYS_REGULAR_CUMULATIVE;
  Module['__mmap_js'] = __mmap_js;
  Module['__munmap_js'] = __munmap_js;
  Module['__setitimer_js'] = __setitimer_js;
  Module['timers'] = timers;
  Module['callUserCallback'] = callUserCallback;
  Module['handleException'] = handleException;
  Module['maybeExit'] = maybeExit;
  Module['_exit'] = _exit;
  Module['exitJS'] = exitJS;
  Module['_proc_exit'] = _proc_exit;
  Module['keepRuntimeAlive'] = keepRuntimeAlive;
  Module['_emscripten_get_now'] = _emscripten_get_now;
  Module['__tzset_js'] = __tzset_js;
  Module['_clock_time_get'] = _clock_time_get;
  Module['_emscripten_date_now'] = _emscripten_date_now;
  Module['nowIsMonotonic'] = nowIsMonotonic;
  Module['checkWasiClock'] = checkWasiClock;
  Module['_emscripten_asm_const_int'] = _emscripten_asm_const_int;
  Module['runEmAsmFunction'] = runEmAsmFunction;
  Module['readEmAsmArgs'] = readEmAsmArgs;
  Module['readEmAsmArgsArray'] = readEmAsmArgsArray;
  Module['_emscripten_asm_const_int_sync_on_main_thread'] = _emscripten_asm_const_int_sync_on_main_thread;
  Module['runMainThreadEmAsm'] = runMainThreadEmAsm;
  Module['_emscripten_asm_const_ptr_sync_on_main_thread'] = _emscripten_asm_const_ptr_sync_on_main_thread;
  Module['_emscripten_cancel_main_loop'] = _emscripten_cancel_main_loop;
  Module['MainLoop'] = MainLoop;
  Module['setMainLoop'] = setMainLoop;
  Module['_emscripten_set_main_loop_timing'] = _emscripten_set_main_loop_timing;
  Module['_emscripten_clear_timeout'] = _emscripten_clear_timeout;
  Module['_emscripten_err'] = _emscripten_err;
  Module['_emscripten_exit_fullscreen'] = _emscripten_exit_fullscreen;
  Module['JSEvents'] = JSEvents;
  Module['addOnExit'] = addOnExit;
  Module['onExits'] = onExits;
  Module['specialHTMLTargets'] = specialHTMLTargets;
  Module['JSEvents_requestFullscreen'] = JSEvents_requestFullscreen;
  Module['JSEvents_resizeCanvasForFullscreen'] = JSEvents_resizeCanvasForFullscreen;
  Module['registerRestoreOldStyle'] = registerRestoreOldStyle;
  Module['getCanvasElementSize'] = getCanvasElementSize;
  Module['_emscripten_get_canvas_element_size'] = _emscripten_get_canvas_element_size;
  Module['findCanvasEventTarget'] = findCanvasEventTarget;
  Module['findEventTarget'] = findEventTarget;
  Module['maybeCStringToJsString'] = maybeCStringToJsString;
  Module['stringToUTF8OnStack'] = stringToUTF8OnStack;
  Module['stackAlloc'] = stackAlloc;
  Module['setCanvasElementSize'] = setCanvasElementSize;
  Module['_emscripten_set_canvas_element_size'] = _emscripten_set_canvas_element_size;
  Module['currentFullscreenStrategy'] = currentFullscreenStrategy;
  Module['setLetterbox'] = setLetterbox;
  Module['getBoundingClientRect'] = getBoundingClientRect;
  Module['_emscripten_exit_pointerlock'] = _emscripten_exit_pointerlock;
  Module['requestPointerLock'] = requestPointerLock;
  Module['_emscripten_force_exit'] = _emscripten_force_exit;
  Module['_emscripten_get_battery_status'] = _emscripten_get_battery_status;
  Module['fillBatteryEventData'] = fillBatteryEventData;
  Module['battery'] = battery;
  Module['_emscripten_get_device_pixel_ratio'] = _emscripten_get_device_pixel_ratio;
  Module['_emscripten_get_element_css_size'] = _emscripten_get_element_css_size;
  Module['_emscripten_get_gamepad_status'] = _emscripten_get_gamepad_status;
  Module['fillGamepadEventData'] = fillGamepadEventData;
  Module['_emscripten_get_heap_max'] = _emscripten_get_heap_max;
  Module['getHeapMax'] = getHeapMax;
  Module['_emscripten_get_main_loop_timing'] = _emscripten_get_main_loop_timing;
  Module['_emscripten_get_num_gamepads'] = _emscripten_get_num_gamepads;
  Module['_emscripten_get_preloaded_image_data'] = _emscripten_get_preloaded_image_data;
  Module['getPreloadedImageData'] = getPreloadedImageData;
  Module['Browser'] = Browser;
  Module['safeSetTimeout'] = safeSetTimeout;
  Module['_emscripten_get_preloaded_image_data_from_FILE'] = _emscripten_get_preloaded_image_data_from_FILE;
  Module['_emscripten_get_screen_size'] = _emscripten_get_screen_size;
  Module['_emscripten_glActiveTexture'] = _emscripten_glActiveTexture;
  Module['_glActiveTexture'] = _glActiveTexture;
  Module['GL'] = GL;
  Module['GLctx'] = GLctx;
  Module['webgl_enable_ANGLE_instanced_arrays'] = webgl_enable_ANGLE_instanced_arrays;
  Module['webgl_enable_OES_vertex_array_object'] = webgl_enable_OES_vertex_array_object;
  Module['webgl_enable_WEBGL_draw_buffers'] = webgl_enable_WEBGL_draw_buffers;
  Module['webgl_enable_EXT_polygon_offset_clamp'] = webgl_enable_EXT_polygon_offset_clamp;
  Module['webgl_enable_EXT_clip_control'] = webgl_enable_EXT_clip_control;
  Module['webgl_enable_WEBGL_polygon_mode'] = webgl_enable_WEBGL_polygon_mode;
  Module['webgl_enable_WEBGL_multi_draw'] = webgl_enable_WEBGL_multi_draw;
  Module['getEmscriptenSupportedExtensions'] = getEmscriptenSupportedExtensions;
  Module['_emscripten_glAttachShader'] = _emscripten_glAttachShader;
  Module['_glAttachShader'] = _glAttachShader;
  Module['_emscripten_glBeginQueryEXT'] = _emscripten_glBeginQueryEXT;
  Module['_glBeginQueryEXT'] = _glBeginQueryEXT;
  Module['_emscripten_glBindAttribLocation'] = _emscripten_glBindAttribLocation;
  Module['_glBindAttribLocation'] = _glBindAttribLocation;
  Module['_emscripten_glBindBuffer'] = _emscripten_glBindBuffer;
  Module['_glBindBuffer'] = _glBindBuffer;
  Module['_emscripten_glBindFramebuffer'] = _emscripten_glBindFramebuffer;
  Module['_glBindFramebuffer'] = _glBindFramebuffer;
  Module['_emscripten_glBindRenderbuffer'] = _emscripten_glBindRenderbuffer;
  Module['_glBindRenderbuffer'] = _glBindRenderbuffer;
  Module['_emscripten_glBindTexture'] = _emscripten_glBindTexture;
  Module['_glBindTexture'] = _glBindTexture;
  Module['_emscripten_glBindVertexArrayOES'] = _emscripten_glBindVertexArrayOES;
  Module['_glBindVertexArrayOES'] = _glBindVertexArrayOES;
  Module['_glBindVertexArray'] = _glBindVertexArray;
  Module['_emscripten_glBlendColor'] = _emscripten_glBlendColor;
  Module['_glBlendColor'] = _glBlendColor;
  Module['_emscripten_glBlendEquation'] = _emscripten_glBlendEquation;
  Module['_glBlendEquation'] = _glBlendEquation;
  Module['_emscripten_glBlendEquationSeparate'] = _emscripten_glBlendEquationSeparate;
  Module['_glBlendEquationSeparate'] = _glBlendEquationSeparate;
  Module['_emscripten_glBlendFunc'] = _emscripten_glBlendFunc;
  Module['_glBlendFunc'] = _glBlendFunc;
  Module['_emscripten_glBlendFuncSeparate'] = _emscripten_glBlendFuncSeparate;
  Module['_glBlendFuncSeparate'] = _glBlendFuncSeparate;
  Module['_emscripten_glBufferData'] = _emscripten_glBufferData;
  Module['_glBufferData'] = _glBufferData;
  Module['_emscripten_glBufferSubData'] = _emscripten_glBufferSubData;
  Module['_glBufferSubData'] = _glBufferSubData;
  Module['_emscripten_glCheckFramebufferStatus'] = _emscripten_glCheckFramebufferStatus;
  Module['_glCheckFramebufferStatus'] = _glCheckFramebufferStatus;
  Module['_emscripten_glClear'] = _emscripten_glClear;
  Module['_glClear'] = _glClear;
  Module['_emscripten_glClearColor'] = _emscripten_glClearColor;
  Module['_glClearColor'] = _glClearColor;
  Module['_emscripten_glClearDepthf'] = _emscripten_glClearDepthf;
  Module['_glClearDepthf'] = _glClearDepthf;
  Module['_emscripten_glClearStencil'] = _emscripten_glClearStencil;
  Module['_glClearStencil'] = _glClearStencil;
  Module['_emscripten_glClipControlEXT'] = _emscripten_glClipControlEXT;
  Module['_glClipControlEXT'] = _glClipControlEXT;
  Module['_emscripten_glColorMask'] = _emscripten_glColorMask;
  Module['_glColorMask'] = _glColorMask;
  Module['_emscripten_glCompileShader'] = _emscripten_glCompileShader;
  Module['_glCompileShader'] = _glCompileShader;
  Module['_emscripten_glCompressedTexImage2D'] = _emscripten_glCompressedTexImage2D;
  Module['_glCompressedTexImage2D'] = _glCompressedTexImage2D;
  Module['_emscripten_glCompressedTexSubImage2D'] = _emscripten_glCompressedTexSubImage2D;
  Module['_glCompressedTexSubImage2D'] = _glCompressedTexSubImage2D;
  Module['_emscripten_glCopyTexImage2D'] = _emscripten_glCopyTexImage2D;
  Module['_glCopyTexImage2D'] = _glCopyTexImage2D;
  Module['_emscripten_glCopyTexSubImage2D'] = _emscripten_glCopyTexSubImage2D;
  Module['_glCopyTexSubImage2D'] = _glCopyTexSubImage2D;
  Module['_emscripten_glCreateProgram'] = _emscripten_glCreateProgram;
  Module['_glCreateProgram'] = _glCreateProgram;
  Module['_emscripten_glCreateShader'] = _emscripten_glCreateShader;
  Module['_glCreateShader'] = _glCreateShader;
  Module['_emscripten_glCullFace'] = _emscripten_glCullFace;
  Module['_glCullFace'] = _glCullFace;
  Module['_emscripten_glDeleteBuffers'] = _emscripten_glDeleteBuffers;
  Module['_glDeleteBuffers'] = _glDeleteBuffers;
  Module['_emscripten_glDeleteFramebuffers'] = _emscripten_glDeleteFramebuffers;
  Module['_glDeleteFramebuffers'] = _glDeleteFramebuffers;
  Module['_emscripten_glDeleteProgram'] = _emscripten_glDeleteProgram;
  Module['_glDeleteProgram'] = _glDeleteProgram;
  Module['_emscripten_glDeleteQueriesEXT'] = _emscripten_glDeleteQueriesEXT;
  Module['_glDeleteQueriesEXT'] = _glDeleteQueriesEXT;
  Module['_emscripten_glDeleteRenderbuffers'] = _emscripten_glDeleteRenderbuffers;
  Module['_glDeleteRenderbuffers'] = _glDeleteRenderbuffers;
  Module['_emscripten_glDeleteShader'] = _emscripten_glDeleteShader;
  Module['_glDeleteShader'] = _glDeleteShader;
  Module['_emscripten_glDeleteTextures'] = _emscripten_glDeleteTextures;
  Module['_glDeleteTextures'] = _glDeleteTextures;
  Module['_emscripten_glDeleteVertexArraysOES'] = _emscripten_glDeleteVertexArraysOES;
  Module['_glDeleteVertexArraysOES'] = _glDeleteVertexArraysOES;
  Module['_glDeleteVertexArrays'] = _glDeleteVertexArrays;
  Module['_emscripten_glDepthFunc'] = _emscripten_glDepthFunc;
  Module['_glDepthFunc'] = _glDepthFunc;
  Module['_emscripten_glDepthMask'] = _emscripten_glDepthMask;
  Module['_glDepthMask'] = _glDepthMask;
  Module['_emscripten_glDepthRangef'] = _emscripten_glDepthRangef;
  Module['_glDepthRangef'] = _glDepthRangef;
  Module['_emscripten_glDetachShader'] = _emscripten_glDetachShader;
  Module['_glDetachShader'] = _glDetachShader;
  Module['_emscripten_glDisable'] = _emscripten_glDisable;
  Module['_glDisable'] = _glDisable;
  Module['_emscripten_glDisableVertexAttribArray'] = _emscripten_glDisableVertexAttribArray;
  Module['_glDisableVertexAttribArray'] = _glDisableVertexAttribArray;
  Module['_emscripten_glDrawArrays'] = _emscripten_glDrawArrays;
  Module['_glDrawArrays'] = _glDrawArrays;
  Module['_emscripten_glDrawArraysInstancedANGLE'] = _emscripten_glDrawArraysInstancedANGLE;
  Module['_glDrawArraysInstancedANGLE'] = _glDrawArraysInstancedANGLE;
  Module['_glDrawArraysInstanced'] = _glDrawArraysInstanced;
  Module['_emscripten_glDrawBuffersWEBGL'] = _emscripten_glDrawBuffersWEBGL;
  Module['_glDrawBuffersWEBGL'] = _glDrawBuffersWEBGL;
  Module['_glDrawBuffers'] = _glDrawBuffers;
  Module['tempFixedLengthArray'] = tempFixedLengthArray;
  Module['_emscripten_glDrawElements'] = _emscripten_glDrawElements;
  Module['_glDrawElements'] = _glDrawElements;
  Module['_emscripten_glDrawElementsInstancedANGLE'] = _emscripten_glDrawElementsInstancedANGLE;
  Module['_glDrawElementsInstancedANGLE'] = _glDrawElementsInstancedANGLE;
  Module['_glDrawElementsInstanced'] = _glDrawElementsInstanced;
  Module['_emscripten_glEnable'] = _emscripten_glEnable;
  Module['_glEnable'] = _glEnable;
  Module['_emscripten_glEnableVertexAttribArray'] = _emscripten_glEnableVertexAttribArray;
  Module['_glEnableVertexAttribArray'] = _glEnableVertexAttribArray;
  Module['_emscripten_glEndQueryEXT'] = _emscripten_glEndQueryEXT;
  Module['_glEndQueryEXT'] = _glEndQueryEXT;
  Module['_emscripten_glFinish'] = _emscripten_glFinish;
  Module['_glFinish'] = _glFinish;
  Module['_emscripten_glFlush'] = _emscripten_glFlush;
  Module['_glFlush'] = _glFlush;
  Module['_emscripten_glFramebufferRenderbuffer'] = _emscripten_glFramebufferRenderbuffer;
  Module['_glFramebufferRenderbuffer'] = _glFramebufferRenderbuffer;
  Module['_emscripten_glFramebufferTexture2D'] = _emscripten_glFramebufferTexture2D;
  Module['_glFramebufferTexture2D'] = _glFramebufferTexture2D;
  Module['_emscripten_glFrontFace'] = _emscripten_glFrontFace;
  Module['_glFrontFace'] = _glFrontFace;
  Module['_emscripten_glGenBuffers'] = _emscripten_glGenBuffers;
  Module['_glGenBuffers'] = _glGenBuffers;
  Module['_emscripten_glGenFramebuffers'] = _emscripten_glGenFramebuffers;
  Module['_glGenFramebuffers'] = _glGenFramebuffers;
  Module['_emscripten_glGenQueriesEXT'] = _emscripten_glGenQueriesEXT;
  Module['_glGenQueriesEXT'] = _glGenQueriesEXT;
  Module['_emscripten_glGenRenderbuffers'] = _emscripten_glGenRenderbuffers;
  Module['_glGenRenderbuffers'] = _glGenRenderbuffers;
  Module['_emscripten_glGenTextures'] = _emscripten_glGenTextures;
  Module['_glGenTextures'] = _glGenTextures;
  Module['_emscripten_glGenVertexArraysOES'] = _emscripten_glGenVertexArraysOES;
  Module['_glGenVertexArraysOES'] = _glGenVertexArraysOES;
  Module['_glGenVertexArrays'] = _glGenVertexArrays;
  Module['_emscripten_glGenerateMipmap'] = _emscripten_glGenerateMipmap;
  Module['_glGenerateMipmap'] = _glGenerateMipmap;
  Module['_emscripten_glGetActiveAttrib'] = _emscripten_glGetActiveAttrib;
  Module['_glGetActiveAttrib'] = _glGetActiveAttrib;
  Module['__glGetActiveAttribOrUniform'] = __glGetActiveAttribOrUniform;
  Module['_emscripten_glGetActiveUniform'] = _emscripten_glGetActiveUniform;
  Module['_glGetActiveUniform'] = _glGetActiveUniform;
  Module['_emscripten_glGetAttachedShaders'] = _emscripten_glGetAttachedShaders;
  Module['_glGetAttachedShaders'] = _glGetAttachedShaders;
  Module['_emscripten_glGetAttribLocation'] = _emscripten_glGetAttribLocation;
  Module['_glGetAttribLocation'] = _glGetAttribLocation;
  Module['_emscripten_glGetBooleanv'] = _emscripten_glGetBooleanv;
  Module['_glGetBooleanv'] = _glGetBooleanv;
  Module['emscriptenWebGLGet'] = emscriptenWebGLGet;
  Module['writeI53ToI64'] = writeI53ToI64;
  Module['readI53FromI64'] = readI53FromI64;
  Module['readI53FromU64'] = readI53FromU64;
  Module['_emscripten_glGetBufferParameteriv'] = _emscripten_glGetBufferParameteriv;
  Module['_glGetBufferParameteriv'] = _glGetBufferParameteriv;
  Module['_emscripten_glGetError'] = _emscripten_glGetError;
  Module['_glGetError'] = _glGetError;
  Module['_emscripten_glGetFloatv'] = _emscripten_glGetFloatv;
  Module['_glGetFloatv'] = _glGetFloatv;
  Module['_emscripten_glGetFramebufferAttachmentParameteriv'] = _emscripten_glGetFramebufferAttachmentParameteriv;
  Module['_glGetFramebufferAttachmentParameteriv'] = _glGetFramebufferAttachmentParameteriv;
  Module['_emscripten_glGetIntegerv'] = _emscripten_glGetIntegerv;
  Module['_glGetIntegerv'] = _glGetIntegerv;
  Module['_emscripten_glGetProgramInfoLog'] = _emscripten_glGetProgramInfoLog;
  Module['_glGetProgramInfoLog'] = _glGetProgramInfoLog;
  Module['_emscripten_glGetProgramiv'] = _emscripten_glGetProgramiv;
  Module['_glGetProgramiv'] = _glGetProgramiv;
  Module['_emscripten_glGetQueryObjecti64vEXT'] = _emscripten_glGetQueryObjecti64vEXT;
  Module['_glGetQueryObjecti64vEXT'] = _glGetQueryObjecti64vEXT;
  Module['_emscripten_glGetQueryObjectivEXT'] = _emscripten_glGetQueryObjectivEXT;
  Module['_glGetQueryObjectivEXT'] = _glGetQueryObjectivEXT;
  Module['_emscripten_glGetQueryObjectui64vEXT'] = _emscripten_glGetQueryObjectui64vEXT;
  Module['_glGetQueryObjectui64vEXT'] = _glGetQueryObjectui64vEXT;
  Module['_emscripten_glGetQueryObjectuivEXT'] = _emscripten_glGetQueryObjectuivEXT;
  Module['_glGetQueryObjectuivEXT'] = _glGetQueryObjectuivEXT;
  Module['_emscripten_glGetQueryivEXT'] = _emscripten_glGetQueryivEXT;
  Module['_glGetQueryivEXT'] = _glGetQueryivEXT;
  Module['_emscripten_glGetRenderbufferParameteriv'] = _emscripten_glGetRenderbufferParameteriv;
  Module['_glGetRenderbufferParameteriv'] = _glGetRenderbufferParameteriv;
  Module['_emscripten_glGetShaderInfoLog'] = _emscripten_glGetShaderInfoLog;
  Module['_glGetShaderInfoLog'] = _glGetShaderInfoLog;
  Module['_emscripten_glGetShaderPrecisionFormat'] = _emscripten_glGetShaderPrecisionFormat;
  Module['_glGetShaderPrecisionFormat'] = _glGetShaderPrecisionFormat;
  Module['_emscripten_glGetShaderSource'] = _emscripten_glGetShaderSource;
  Module['_glGetShaderSource'] = _glGetShaderSource;
  Module['_emscripten_glGetShaderiv'] = _emscripten_glGetShaderiv;
  Module['_glGetShaderiv'] = _glGetShaderiv;
  Module['_emscripten_glGetString'] = _emscripten_glGetString;
  Module['_glGetString'] = _glGetString;
  Module['stringToNewUTF8'] = stringToNewUTF8;
  Module['webglGetExtensions'] = webglGetExtensions;
  Module['_emscripten_glGetTexParameterfv'] = _emscripten_glGetTexParameterfv;
  Module['_glGetTexParameterfv'] = _glGetTexParameterfv;
  Module['_emscripten_glGetTexParameteriv'] = _emscripten_glGetTexParameteriv;
  Module['_glGetTexParameteriv'] = _glGetTexParameteriv;
  Module['_emscripten_glGetUniformLocation'] = _emscripten_glGetUniformLocation;
  Module['_glGetUniformLocation'] = _glGetUniformLocation;
  Module['jstoi_q'] = jstoi_q;
  Module['webglPrepareUniformLocationsBeforeFirstUse'] = webglPrepareUniformLocationsBeforeFirstUse;
  Module['webglGetLeftBracePos'] = webglGetLeftBracePos;
  Module['_emscripten_glGetUniformfv'] = _emscripten_glGetUniformfv;
  Module['_glGetUniformfv'] = _glGetUniformfv;
  Module['emscriptenWebGLGetUniform'] = emscriptenWebGLGetUniform;
  Module['webglGetUniformLocation'] = webglGetUniformLocation;
  Module['_emscripten_glGetUniformiv'] = _emscripten_glGetUniformiv;
  Module['_glGetUniformiv'] = _glGetUniformiv;
  Module['_emscripten_glGetVertexAttribPointerv'] = _emscripten_glGetVertexAttribPointerv;
  Module['_glGetVertexAttribPointerv'] = _glGetVertexAttribPointerv;
  Module['_emscripten_glGetVertexAttribfv'] = _emscripten_glGetVertexAttribfv;
  Module['_glGetVertexAttribfv'] = _glGetVertexAttribfv;
  Module['emscriptenWebGLGetVertexAttrib'] = emscriptenWebGLGetVertexAttrib;
  Module['_emscripten_glGetVertexAttribiv'] = _emscripten_glGetVertexAttribiv;
  Module['_glGetVertexAttribiv'] = _glGetVertexAttribiv;
  Module['_emscripten_glHint'] = _emscripten_glHint;
  Module['_glHint'] = _glHint;
  Module['_emscripten_glIsBuffer'] = _emscripten_glIsBuffer;
  Module['_glIsBuffer'] = _glIsBuffer;
  Module['_emscripten_glIsEnabled'] = _emscripten_glIsEnabled;
  Module['_glIsEnabled'] = _glIsEnabled;
  Module['_emscripten_glIsFramebuffer'] = _emscripten_glIsFramebuffer;
  Module['_glIsFramebuffer'] = _glIsFramebuffer;
  Module['_emscripten_glIsProgram'] = _emscripten_glIsProgram;
  Module['_glIsProgram'] = _glIsProgram;
  Module['_emscripten_glIsQueryEXT'] = _emscripten_glIsQueryEXT;
  Module['_glIsQueryEXT'] = _glIsQueryEXT;
  Module['_emscripten_glIsRenderbuffer'] = _emscripten_glIsRenderbuffer;
  Module['_glIsRenderbuffer'] = _glIsRenderbuffer;
  Module['_emscripten_glIsShader'] = _emscripten_glIsShader;
  Module['_glIsShader'] = _glIsShader;
  Module['_emscripten_glIsTexture'] = _emscripten_glIsTexture;
  Module['_glIsTexture'] = _glIsTexture;
  Module['_emscripten_glIsVertexArrayOES'] = _emscripten_glIsVertexArrayOES;
  Module['_glIsVertexArrayOES'] = _glIsVertexArrayOES;
  Module['_glIsVertexArray'] = _glIsVertexArray;
  Module['_emscripten_glLineWidth'] = _emscripten_glLineWidth;
  Module['_glLineWidth'] = _glLineWidth;
  Module['_emscripten_glLinkProgram'] = _emscripten_glLinkProgram;
  Module['_glLinkProgram'] = _glLinkProgram;
  Module['_emscripten_glPixelStorei'] = _emscripten_glPixelStorei;
  Module['_glPixelStorei'] = _glPixelStorei;
  Module['_emscripten_glPolygonModeWEBGL'] = _emscripten_glPolygonModeWEBGL;
  Module['_glPolygonModeWEBGL'] = _glPolygonModeWEBGL;
  Module['_emscripten_glPolygonOffset'] = _emscripten_glPolygonOffset;
  Module['_glPolygonOffset'] = _glPolygonOffset;
  Module['_emscripten_glPolygonOffsetClampEXT'] = _emscripten_glPolygonOffsetClampEXT;
  Module['_glPolygonOffsetClampEXT'] = _glPolygonOffsetClampEXT;
  Module['_emscripten_glQueryCounterEXT'] = _emscripten_glQueryCounterEXT;
  Module['_glQueryCounterEXT'] = _glQueryCounterEXT;
  Module['_emscripten_glReadPixels'] = _emscripten_glReadPixels;
  Module['_glReadPixels'] = _glReadPixels;
  Module['emscriptenWebGLGetTexPixelData'] = emscriptenWebGLGetTexPixelData;
  Module['computeUnpackAlignedImageSize'] = computeUnpackAlignedImageSize;
  Module['colorChannelsInGlTextureFormat'] = colorChannelsInGlTextureFormat;
  Module['heapObjectForWebGLType'] = heapObjectForWebGLType;
  Module['toTypedArrayIndex'] = toTypedArrayIndex;
  Module['_emscripten_glReleaseShaderCompiler'] = _emscripten_glReleaseShaderCompiler;
  Module['_glReleaseShaderCompiler'] = _glReleaseShaderCompiler;
  Module['_emscripten_glRenderbufferStorage'] = _emscripten_glRenderbufferStorage;
  Module['_glRenderbufferStorage'] = _glRenderbufferStorage;
  Module['_emscripten_glSampleCoverage'] = _emscripten_glSampleCoverage;
  Module['_glSampleCoverage'] = _glSampleCoverage;
  Module['_emscripten_glScissor'] = _emscripten_glScissor;
  Module['_glScissor'] = _glScissor;
  Module['_emscripten_glShaderBinary'] = _emscripten_glShaderBinary;
  Module['_glShaderBinary'] = _glShaderBinary;
  Module['_emscripten_glShaderSource'] = _emscripten_glShaderSource;
  Module['_glShaderSource'] = _glShaderSource;
  Module['_emscripten_glStencilFunc'] = _emscripten_glStencilFunc;
  Module['_glStencilFunc'] = _glStencilFunc;
  Module['_emscripten_glStencilFuncSeparate'] = _emscripten_glStencilFuncSeparate;
  Module['_glStencilFuncSeparate'] = _glStencilFuncSeparate;
  Module['_emscripten_glStencilMask'] = _emscripten_glStencilMask;
  Module['_glStencilMask'] = _glStencilMask;
  Module['_emscripten_glStencilMaskSeparate'] = _emscripten_glStencilMaskSeparate;
  Module['_glStencilMaskSeparate'] = _glStencilMaskSeparate;
  Module['_emscripten_glStencilOp'] = _emscripten_glStencilOp;
  Module['_glStencilOp'] = _glStencilOp;
  Module['_emscripten_glStencilOpSeparate'] = _emscripten_glStencilOpSeparate;
  Module['_glStencilOpSeparate'] = _glStencilOpSeparate;
  Module['_emscripten_glTexImage2D'] = _emscripten_glTexImage2D;
  Module['_glTexImage2D'] = _glTexImage2D;
  Module['_emscripten_glTexParameterf'] = _emscripten_glTexParameterf;
  Module['_glTexParameterf'] = _glTexParameterf;
  Module['_emscripten_glTexParameterfv'] = _emscripten_glTexParameterfv;
  Module['_glTexParameterfv'] = _glTexParameterfv;
  Module['_emscripten_glTexParameteri'] = _emscripten_glTexParameteri;
  Module['_glTexParameteri'] = _glTexParameteri;
  Module['_emscripten_glTexParameteriv'] = _emscripten_glTexParameteriv;
  Module['_glTexParameteriv'] = _glTexParameteriv;
  Module['_emscripten_glTexSubImage2D'] = _emscripten_glTexSubImage2D;
  Module['_glTexSubImage2D'] = _glTexSubImage2D;
  Module['_emscripten_glUniform1f'] = _emscripten_glUniform1f;
  Module['_glUniform1f'] = _glUniform1f;
  Module['_emscripten_glUniform1fv'] = _emscripten_glUniform1fv;
  Module['_glUniform1fv'] = _glUniform1fv;
  Module['miniTempWebGLFloatBuffers'] = miniTempWebGLFloatBuffers;
  Module['_emscripten_glUniform1i'] = _emscripten_glUniform1i;
  Module['_glUniform1i'] = _glUniform1i;
  Module['_emscripten_glUniform1iv'] = _emscripten_glUniform1iv;
  Module['_glUniform1iv'] = _glUniform1iv;
  Module['miniTempWebGLIntBuffers'] = miniTempWebGLIntBuffers;
  Module['_emscripten_glUniform2f'] = _emscripten_glUniform2f;
  Module['_glUniform2f'] = _glUniform2f;
  Module['_emscripten_glUniform2fv'] = _emscripten_glUniform2fv;
  Module['_glUniform2fv'] = _glUniform2fv;
  Module['_emscripten_glUniform2i'] = _emscripten_glUniform2i;
  Module['_glUniform2i'] = _glUniform2i;
  Module['_emscripten_glUniform2iv'] = _emscripten_glUniform2iv;
  Module['_glUniform2iv'] = _glUniform2iv;
  Module['_emscripten_glUniform3f'] = _emscripten_glUniform3f;
  Module['_glUniform3f'] = _glUniform3f;
  Module['_emscripten_glUniform3fv'] = _emscripten_glUniform3fv;
  Module['_glUniform3fv'] = _glUniform3fv;
  Module['_emscripten_glUniform3i'] = _emscripten_glUniform3i;
  Module['_glUniform3i'] = _glUniform3i;
  Module['_emscripten_glUniform3iv'] = _emscripten_glUniform3iv;
  Module['_glUniform3iv'] = _glUniform3iv;
  Module['_emscripten_glUniform4f'] = _emscripten_glUniform4f;
  Module['_glUniform4f'] = _glUniform4f;
  Module['_emscripten_glUniform4fv'] = _emscripten_glUniform4fv;
  Module['_glUniform4fv'] = _glUniform4fv;
  Module['_emscripten_glUniform4i'] = _emscripten_glUniform4i;
  Module['_glUniform4i'] = _glUniform4i;
  Module['_emscripten_glUniform4iv'] = _emscripten_glUniform4iv;
  Module['_glUniform4iv'] = _glUniform4iv;
  Module['_emscripten_glUniformMatrix2fv'] = _emscripten_glUniformMatrix2fv;
  Module['_glUniformMatrix2fv'] = _glUniformMatrix2fv;
  Module['_emscripten_glUniformMatrix3fv'] = _emscripten_glUniformMatrix3fv;
  Module['_glUniformMatrix3fv'] = _glUniformMatrix3fv;
  Module['_emscripten_glUniformMatrix4fv'] = _emscripten_glUniformMatrix4fv;
  Module['_glUniformMatrix4fv'] = _glUniformMatrix4fv;
  Module['_emscripten_glUseProgram'] = _emscripten_glUseProgram;
  Module['_glUseProgram'] = _glUseProgram;
  Module['_emscripten_glValidateProgram'] = _emscripten_glValidateProgram;
  Module['_glValidateProgram'] = _glValidateProgram;
  Module['_emscripten_glVertexAttrib1f'] = _emscripten_glVertexAttrib1f;
  Module['_glVertexAttrib1f'] = _glVertexAttrib1f;
  Module['_emscripten_glVertexAttrib1fv'] = _emscripten_glVertexAttrib1fv;
  Module['_glVertexAttrib1fv'] = _glVertexAttrib1fv;
  Module['_emscripten_glVertexAttrib2f'] = _emscripten_glVertexAttrib2f;
  Module['_glVertexAttrib2f'] = _glVertexAttrib2f;
  Module['_emscripten_glVertexAttrib2fv'] = _emscripten_glVertexAttrib2fv;
  Module['_glVertexAttrib2fv'] = _glVertexAttrib2fv;
  Module['_emscripten_glVertexAttrib3f'] = _emscripten_glVertexAttrib3f;
  Module['_glVertexAttrib3f'] = _glVertexAttrib3f;
  Module['_emscripten_glVertexAttrib3fv'] = _emscripten_glVertexAttrib3fv;
  Module['_glVertexAttrib3fv'] = _glVertexAttrib3fv;
  Module['_emscripten_glVertexAttrib4f'] = _emscripten_glVertexAttrib4f;
  Module['_glVertexAttrib4f'] = _glVertexAttrib4f;
  Module['_emscripten_glVertexAttrib4fv'] = _emscripten_glVertexAttrib4fv;
  Module['_glVertexAttrib4fv'] = _glVertexAttrib4fv;
  Module['_emscripten_glVertexAttribDivisorANGLE'] = _emscripten_glVertexAttribDivisorANGLE;
  Module['_glVertexAttribDivisorANGLE'] = _glVertexAttribDivisorANGLE;
  Module['_glVertexAttribDivisor'] = _glVertexAttribDivisor;
  Module['_emscripten_glVertexAttribPointer'] = _emscripten_glVertexAttribPointer;
  Module['_glVertexAttribPointer'] = _glVertexAttribPointer;
  Module['_emscripten_glViewport'] = _emscripten_glViewport;
  Module['_glViewport'] = _glViewport;
  Module['_emscripten_has_asyncify'] = _emscripten_has_asyncify;
  Module['_emscripten_request_fullscreen_strategy'] = _emscripten_request_fullscreen_strategy;
  Module['doRequestFullscreen'] = doRequestFullscreen;
  Module['_emscripten_request_pointerlock'] = _emscripten_request_pointerlock;
  Module['_emscripten_resize_heap'] = _emscripten_resize_heap;
  Module['growMemory'] = growMemory;
  Module['_emscripten_sample_gamepad_data'] = _emscripten_sample_gamepad_data;
  Module['_emscripten_set_beforeunload_callback_on_thread'] = _emscripten_set_beforeunload_callback_on_thread;
  Module['registerBeforeUnloadEventCallback'] = registerBeforeUnloadEventCallback;
  Module['_emscripten_set_blur_callback_on_thread'] = _emscripten_set_blur_callback_on_thread;
  Module['registerFocusEventCallback'] = registerFocusEventCallback;
  Module['_emscripten_set_element_css_size'] = _emscripten_set_element_css_size;
  Module['_emscripten_set_focus_callback_on_thread'] = _emscripten_set_focus_callback_on_thread;
  Module['_emscripten_set_fullscreenchange_callback_on_thread'] = _emscripten_set_fullscreenchange_callback_on_thread;
  Module['registerFullscreenChangeEventCallback'] = registerFullscreenChangeEventCallback;
  Module['fillFullscreenChangeEventData'] = fillFullscreenChangeEventData;
  Module['_emscripten_set_gamepadconnected_callback_on_thread'] = _emscripten_set_gamepadconnected_callback_on_thread;
  Module['registerGamepadEventCallback'] = registerGamepadEventCallback;
  Module['_emscripten_set_gamepaddisconnected_callback_on_thread'] = _emscripten_set_gamepaddisconnected_callback_on_thread;
  Module['_emscripten_set_keydown_callback_on_thread'] = _emscripten_set_keydown_callback_on_thread;
  Module['registerKeyEventCallback'] = registerKeyEventCallback;
  Module['_emscripten_set_keypress_callback_on_thread'] = _emscripten_set_keypress_callback_on_thread;
  Module['_emscripten_set_keyup_callback_on_thread'] = _emscripten_set_keyup_callback_on_thread;
  Module['_emscripten_set_main_loop'] = _emscripten_set_main_loop;
  Module['_emscripten_set_mousedown_callback_on_thread'] = _emscripten_set_mousedown_callback_on_thread;
  Module['registerMouseEventCallback'] = registerMouseEventCallback;
  Module['fillMouseEventData'] = fillMouseEventData;
  Module['_emscripten_set_mouseenter_callback_on_thread'] = _emscripten_set_mouseenter_callback_on_thread;
  Module['_emscripten_set_mouseleave_callback_on_thread'] = _emscripten_set_mouseleave_callback_on_thread;
  Module['_emscripten_set_mousemove_callback_on_thread'] = _emscripten_set_mousemove_callback_on_thread;
  Module['_emscripten_set_mouseup_callback_on_thread'] = _emscripten_set_mouseup_callback_on_thread;
  Module['_emscripten_set_orientationchange_callback_on_thread'] = _emscripten_set_orientationchange_callback_on_thread;
  Module['registerOrientationChangeEventCallback'] = registerOrientationChangeEventCallback;
  Module['fillOrientationChangeEventData'] = fillOrientationChangeEventData;
  Module['screenOrientation'] = screenOrientation;
  Module['_emscripten_set_pointerlockchange_callback_on_thread'] = _emscripten_set_pointerlockchange_callback_on_thread;
  Module['registerPointerlockChangeEventCallback'] = registerPointerlockChangeEventCallback;
  Module['fillPointerlockChangeEventData'] = fillPointerlockChangeEventData;
  Module['_emscripten_set_resize_callback_on_thread'] = _emscripten_set_resize_callback_on_thread;
  Module['registerUiEventCallback'] = registerUiEventCallback;
  Module['_emscripten_set_timeout'] = _emscripten_set_timeout;
  Module['_emscripten_set_touchcancel_callback_on_thread'] = _emscripten_set_touchcancel_callback_on_thread;
  Module['registerTouchEventCallback'] = registerTouchEventCallback;
  Module['_emscripten_set_touchend_callback_on_thread'] = _emscripten_set_touchend_callback_on_thread;
  Module['_emscripten_set_touchmove_callback_on_thread'] = _emscripten_set_touchmove_callback_on_thread;
  Module['_emscripten_set_touchstart_callback_on_thread'] = _emscripten_set_touchstart_callback_on_thread;
  Module['_emscripten_set_visibilitychange_callback_on_thread'] = _emscripten_set_visibilitychange_callback_on_thread;
  Module['registerVisibilityChangeEventCallback'] = registerVisibilityChangeEventCallback;
  Module['fillVisibilityChangeEventData'] = fillVisibilityChangeEventData;
  Module['_emscripten_set_wheel_callback_on_thread'] = _emscripten_set_wheel_callback_on_thread;
  Module['registerWheelEventCallback'] = registerWheelEventCallback;
  Module['_emscripten_set_window_title'] = _emscripten_set_window_title;
  Module['_emscripten_sleep'] = _emscripten_sleep;
  Module['_emscripten_webgl_create_context'] = _emscripten_webgl_create_context;
  Module['_emscripten_webgl_do_create_context'] = _emscripten_webgl_do_create_context;
  Module['webglPowerPreferences'] = webglPowerPreferences;
  Module['_emscripten_webgl_destroy_context'] = _emscripten_webgl_destroy_context;
  Module['_emscripten_webgl_make_context_current'] = _emscripten_webgl_make_context_current;
  Module['_environ_get'] = _environ_get;
  Module['getEnvStrings'] = getEnvStrings;
  Module['ENV'] = ENV;
  Module['_environ_sizes_get'] = _environ_sizes_get;
  Module['_fd_close'] = _fd_close;
  Module['_fd_read'] = _fd_read;
  Module['doReadv'] = doReadv;
  Module['_fd_seek'] = _fd_seek;
  Module['_fd_write'] = _fd_write;
  Module['doWritev'] = doWritev;
  Module['dynCall'] = dynCall;
  Module['autoResumeAudioContext'] = autoResumeAudioContext;
  Module['writeArrayToMemory'] = writeArrayToMemory;
  Module['allocateUTF8'] = allocateUTF8;
  // End JS library exports

// end include: postlibrary.js

function checkIncomingModuleAPI() {
  ignoredModuleProp('fetchSettings');
}
var ASM_CONSTS = {
  622568: ($0) => { var str = UTF8ToString($0) + '\n\n' + 'Abort/Retry/Ignore/AlwaysIgnore? [ariA] :'; var reply = window.prompt(str, "i"); if (reply === null) { reply = "i"; } return reply.length === 1 ? reply.charCodeAt(0) : -1; },  
 622783: ($0, $1) => { alert(UTF8ToString($0) + "\n\n" + UTF8ToString($1)); },  
 622840: () => { if (typeof(Module['SDL3']) === 'undefined') { Module['SDL3'] = {}; } Module['SDL3'].dummy_audio = {}; Module['SDL3'].dummy_audio.timers = []; Module['SDL3'].dummy_audio.timers[0] = undefined; Module['SDL3'].dummy_audio.timers[1] = undefined; },  
 623086: ($0, $1, $2, $3, $4) => { var a = Module['SDL3'].dummy_audio; if (a.timers[$0] !== undefined) { clearInterval(a.timers[$0]); } a.timers[$0] = setInterval(function() { dynCall('vi', $3, [$4]); }, ($1 / $2) * 1000); },  
 623278: ($0) => { var a = Module['SDL3'].dummy_audio; if (a.timers[$0] !== undefined) { clearInterval(a.timers[$0]); } a.timers[$0] = undefined; },  
 623409: ($0) => { var parms = new URLSearchParams(window.location.search); for (const [key, value] of parms) { if (key.startsWith("SDL_")) { var ckey = stringToNewUTF8(key); var cvalue = stringToNewUTF8(value); if ((ckey != 0) && (cvalue != 0)) { dynCall('iiii', $0, [ckey, cvalue, 1]); } _free(ckey); _free(cvalue); } } },  
 623716: ($0) => { window.open(UTF8ToString($0), "_blank") },  
 623756: () => { if (typeof(AudioContext) !== 'undefined') { return true; } else if (typeof(webkitAudioContext) !== 'undefined') { return true; } return false; },  
 623903: () => { if ((typeof(navigator.mediaDevices) !== 'undefined') && (typeof(navigator.mediaDevices.getUserMedia) !== 'undefined')) { return true; } else if (typeof(navigator.webkitGetUserMedia) !== 'undefined') { return true; } return false; },  
 624137: ($0) => { if (typeof(Module['SDL3']) === 'undefined') { Module['SDL3'] = {}; } var SDL3 = Module['SDL3']; if (!$0) { SDL3.audio_playback = {}; } else { SDL3.audio_recording = {}; } if (!SDL3.audioContext) { if (typeof(AudioContext) !== 'undefined') { SDL3.audioContext = new AudioContext(); } else if (typeof(webkitAudioContext) !== 'undefined') { SDL3.audioContext = new webkitAudioContext(); } if (SDL3.audioContext) { if ((typeof navigator.userActivation) === 'undefined') { autoResumeAudioContext(SDL3.audioContext); } } } return (SDL3.audioContext !== undefined); },  
 624700: () => { return Module['SDL3'].audioContext.sampleRate; },  
 624751: ($0, $1, $2, $3) => { var SDL3 = Module['SDL3']; var have_microphone = function(stream) { if (SDL3.audio_recording.silenceTimer !== undefined) { clearInterval(SDL3.audio_recording.silenceTimer); SDL3.audio_recording.silenceTimer = undefined; SDL3.audio_recording.silenceBuffer = undefined } SDL3.audio_recording.mediaStreamNode = SDL3.audioContext.createMediaStreamSource(stream); SDL3.audio_recording.scriptProcessorNode = SDL3.audioContext.createScriptProcessor($1, $0, 1); SDL3.audio_recording.scriptProcessorNode.onaudioprocess = function(audioProcessingEvent) { if ((SDL3 === undefined) || (SDL3.audio_recording === undefined)) { return; } audioProcessingEvent.outputBuffer.getChannelData(0).fill(0.0); SDL3.audio_recording.currentRecordingBuffer = audioProcessingEvent.inputBuffer; dynCall('ip', $2, [$3]); }; SDL3.audio_recording.mediaStreamNode.connect(SDL3.audio_recording.scriptProcessorNode); SDL3.audio_recording.scriptProcessorNode.connect(SDL3.audioContext.destination); SDL3.audio_recording.stream = stream; }; var no_microphone = function(error) { }; SDL3.audio_recording.silenceBuffer = SDL3.audioContext.createBuffer($0, $1, SDL3.audioContext.sampleRate); SDL3.audio_recording.silenceBuffer.getChannelData(0).fill(0.0); var silence_callback = function() { SDL3.audio_recording.currentRecordingBuffer = SDL3.audio_recording.silenceBuffer; dynCall('ip', $2, [$3]); }; SDL3.audio_recording.silenceTimer = setInterval(silence_callback, ($1 / SDL3.audioContext.sampleRate) * 1000); if ((navigator.mediaDevices !== undefined) && (navigator.mediaDevices.getUserMedia !== undefined)) { navigator.mediaDevices.getUserMedia({ audio: true, video: false }).then(have_microphone).catch(no_microphone); } else if (navigator.webkitGetUserMedia !== undefined) { navigator.webkitGetUserMedia({ audio: true, video: false }, have_microphone, no_microphone); } },  
 626592: ($0, $1, $2, $3) => { var SDL3 = Module['SDL3']; SDL3.audio_playback.scriptProcessorNode = SDL3.audioContext['createScriptProcessor']($1, 0, $0); SDL3.audio_playback.scriptProcessorNode['onaudioprocess'] = function (e) { if ((SDL3 === undefined) || (SDL3.audio_playback === undefined)) { return; } if (SDL3.audio_playback.silenceTimer !== undefined) { clearInterval(SDL3.audio_playback.silenceTimer); SDL3.audio_playback.silenceTimer = undefined; SDL3.audio_playback.silenceBuffer = undefined; } SDL3.audio_playback.currentPlaybackBuffer = e['outputBuffer']; dynCall('ip', $2, [$3]); }; SDL3.audio_playback.scriptProcessorNode['connect'](SDL3.audioContext['destination']); if (SDL3.audioContext.state === 'suspended') { SDL3.audio_playback.silenceBuffer = SDL3.audioContext.createBuffer($0, $1, SDL3.audioContext.sampleRate); SDL3.audio_playback.silenceBuffer.getChannelData(0).fill(0.0); var silence_callback = function() { if ((typeof navigator.userActivation) !== 'undefined') { if (navigator.userActivation.hasBeenActive) { SDL3.audioContext.resume(); } } SDL3.audio_playback.currentPlaybackBuffer = SDL3.audio_playback.silenceBuffer; dynCall('ip', $2, [$3]); SDL3.audio_playback.currentPlaybackBuffer = undefined; }; SDL3.audio_playback.silenceTimer = setInterval(silence_callback, ($1 / SDL3.audioContext.sampleRate) * 1000); } },  
 627908: ($0) => { var SDL3 = Module['SDL3']; if ($0) { if (SDL3.audio_recording.silenceTimer !== undefined) { clearInterval(SDL3.audio_recording.silenceTimer); } if (SDL3.audio_recording.stream !== undefined) { var tracks = SDL3.audio_recording.stream.getAudioTracks(); for (var i = 0; i < tracks.length; i++) { SDL3.audio_recording.stream.removeTrack(tracks[i]); } } if (SDL3.audio_recording.scriptProcessorNode !== undefined) { SDL3.audio_recording.scriptProcessorNode.onaudioprocess = function(audioProcessingEvent) {}; SDL3.audio_recording.scriptProcessorNode.disconnect(); } if (SDL3.audio_recording.mediaStreamNode !== undefined) { SDL3.audio_recording.mediaStreamNode.disconnect(); } SDL3.audio_recording = undefined; } else { if (SDL3.audio_playback.scriptProcessorNode != undefined) { SDL3.audio_playback.scriptProcessorNode.disconnect(); } if (SDL3.audio_playback.silenceTimer !== undefined) { clearInterval(SDL3.audio_playback.silenceTimer); } SDL3.audio_playback = undefined; } if ((SDL3.audioContext !== undefined) && (SDL3.audio_playback === undefined) && (SDL3.audio_recording === undefined)) { SDL3.audioContext.close(); SDL3.audioContext = undefined; } },  
 629064: ($0, $1) => { var buf = $0 >>> 2; var SDL3 = Module['SDL3']; var numChannels = SDL3.audio_playback.currentPlaybackBuffer['numberOfChannels']; for (var c = 0; c < numChannels; ++c) { var channelData = SDL3.audio_playback.currentPlaybackBuffer['getChannelData'](c); if (channelData.length != $1) { throw 'Web Audio playback buffer length mismatch! Destination size: ' + channelData.length + ' samples vs expected ' + $1 + ' samples!'; } for (var j = 0; j < $1; ++j) { channelData[j] = HEAPF32[buf + (j*numChannels + c)]; } } },  
 629577: ($0, $1) => { var SDL3 = Module['SDL3']; var numChannels = SDL3.audio_recording.currentRecordingBuffer.numberOfChannels; for (var c = 0; c < numChannels; ++c) { var channelData = SDL3.audio_recording.currentRecordingBuffer.getChannelData(c); if (channelData.length != $1) { throw 'Web Audio recording buffer length mismatch! Destination size: ' + channelData.length + ' samples vs expected ' + $1 + ' samples!'; } if (numChannels == 1) { for (var j = 0; j < $1; ++j) { setValue($0 + (j * 4), channelData[j], 'float'); } } else { for (var j = 0; j < $1; ++j) { setValue($0 + (((j * numChannels) + c) * 4), channelData[j], 'float'); } } } },  
 630204: () => { if (typeof(Module['SDL3']) === 'undefined') { Module['SDL3'] = {}; } Module['SDL3'].camera = {}; },  
 630305: () => { return (navigator.mediaDevices === undefined) ? 0 : 1; },  
 630364: ($0, $1, $2, $3, $4, $5, $6) => { const device = $0; const w = $1; const h = $2; const framerate_numerator = $3; const framerate_denominator = $4; const outcome = $5; const iterate = $6; const constraints = {}; if ((w <= 0) || (h <= 0)) { constraints.video = true; } else { constraints.video = {}; constraints.video.width = w; constraints.video.height = h; } if ((framerate_numerator > 0) && (framerate_denominator > 0)) { var fps = framerate_numerator / framerate_denominator; constraints.video.frameRate = { ideal: fps }; } function grabNextCameraFrame() { const SDL3 = Module['SDL3']; if ((typeof(SDL3) === 'undefined') || (typeof(SDL3.camera) === 'undefined') || (typeof(SDL3.camera.stream) === 'undefined')) { return; } const nextframems = SDL3.camera.next_frame_time; const now = performance.now(); if (now >= nextframems) { dynCall('vi', iterate, [device]); while (SDL3.camera.next_frame_time < now) { SDL3.camera.next_frame_time += SDL3.camera.fpsincrms; } } requestAnimationFrame(grabNextCameraFrame); } navigator.mediaDevices.getUserMedia(constraints) .then((stream) => { const settings = stream.getVideoTracks()[0].getSettings(); const actualw = settings.width; const actualh = settings.height; const actualfps = settings.frameRate; console.log("Camera is opened! Actual spec: (" + actualw + "x" + actualh + "), fps=" + actualfps); if (dynCall('iiiiii', outcome, [device, 1, actualw, actualh, actualfps])) { const video = document.createElement("video"); video.width = actualw; video.height = actualh; video.style.display = 'none'; video.srcObject = stream; const canvas = document.createElement("canvas"); canvas.width = actualw; canvas.height = actualh; canvas.style.display = 'none'; const ctx2d = canvas.getContext('2d'); const SDL3 = Module['SDL3']; SDL3.camera.width = actualw; SDL3.camera.height = actualh; SDL3.camera.fps = actualfps; SDL3.camera.fpsincrms = 1000.0 / actualfps; SDL3.camera.stream = stream; SDL3.camera.video = video; SDL3.camera.canvas = canvas; SDL3.camera.ctx2d = ctx2d; SDL3.camera.next_frame_time = performance.now(); video.play(); video.addEventListener('loadedmetadata', () => { grabNextCameraFrame(); }); } }) .catch((err) => { console.error("Tried to open camera but it threw an error! " + err.name + ": " + err.message); dynCall('iiiiii', outcome, [device, 0, 0, 0, 0]); }); },  
 632655: () => { const SDL3 = Module['SDL3']; if ((typeof(SDL3) === 'undefined') || (typeof(SDL3.camera) === 'undefined') || (typeof(SDL3.camera.stream) === 'undefined')) { return; } SDL3.camera.stream.getTracks().forEach(track => track.stop()); SDL3.camera = {}; },  
 632906: ($0, $1, $2) => { const w = $0; const h = $1; const rgba = $2; const SDL3 = Module['SDL3']; if ((typeof(SDL3) === 'undefined') || (typeof(SDL3.camera) === 'undefined') || (typeof(SDL3.camera.ctx2d) === 'undefined')) { return 0; } SDL3.camera.ctx2d.drawImage(SDL3.camera.video, 0, 0, w, h); const imgrgba = SDL3.camera.ctx2d.getImageData(0, 0, w, h).data; HEAPU8.set(imgrgba, rgba); return 1; },  
 633284: () => { if (typeof(Module['SDL3']) !== 'undefined') { Module['SDL3'].camera = undefined; } },  
 633371: ($0, $1) => { var buf = $0; var buflen = $1; var list = undefined; if (navigator.languages && navigator.languages.length) { list = navigator.languages; } else { var oneOfThese = navigator.userLanguage || navigator.language || navigator.browserLanguage || navigator.systemLanguage; if (oneOfThese !== undefined) { list = [ oneOfThese ]; } } if (list === undefined) { return; } var str = ""; for (var i = 0; i < list.length; i++) { var item = list[i]; if ((str.length + item.length + 1) > buflen) { break; } if (str.length > 0) { str += ","; } str += item; } str = str.replace(/-/g, "_"); if (buflen > str.length) { buflen = str.length; } for (var i = 0; i < buflen; i++) { setValue(buf + i, str.charCodeAt(i), "i8"); } },  
 634079: ($0, $1, $2) => { var target = document.querySelector(UTF8ToString($1)); if (target) { var data = $0; if (typeof(Module['SDL3']) === 'undefined') { Module['SDL3'] = {}; } var SDL3 = Module['SDL3']; var makePointerEventCStruct = function(event) { var ptr = 0; if (event.pointerType == "pen") { ptr = _SDL_malloc($2); if (ptr != 0) { var rect = target.getBoundingClientRect(); var idx = ptr >> 2; HEAP32[idx++] = event.pointerId; HEAP32[idx++] = (typeof(event.button) !== "undefined") ? event.button : -1; HEAP32[idx++] = event.buttons; HEAPF32[idx++] = event.movementX; HEAPF32[idx++] = event.movementY; HEAPF32[idx++] = event.clientX - rect.left; HEAPF32[idx++] = event.clientY - rect.top; HEAPF32[idx++] = event.pressure; HEAPF32[idx++] = event.tangentialPressure; HEAPF32[idx++] = event.tiltX; HEAPF32[idx++] = event.tiltY; HEAPF32[idx++] = event.twist; } } return ptr; }; SDL3.eventHandlerPointerEnter = function(event) { var d = makePointerEventCStruct(event); if (d != 0) { _Emscripten_HandlePointerEnter(data, d); _SDL_free(d); } }; target.addEventListener("pointerenter", SDL3.eventHandlerPointerEnter); SDL3.eventHandlerPointerLeave = function(event) { var d = makePointerEventCStruct(event); if (d != 0) { _Emscripten_HandlePointerLeave(data, d); _SDL_free(d); } }; target.addEventListener("pointerleave", SDL3.eventHandlerPointerLeave); target.addEventListener("pointercancel", SDL3.eventHandlerPointerLeave); SDL3.eventHandlerPointerGeneric = function(event) { var d = makePointerEventCStruct(event); if (d != 0) { _Emscripten_HandlePointerGeneric(data, d); _SDL_free(d); } }; target.addEventListener("pointerdown", SDL3.eventHandlerPointerGeneric); target.addEventListener("pointerup", SDL3.eventHandlerPointerGeneric); target.addEventListener("pointermove", SDL3.eventHandlerPointerGeneric); } },  
 635872: ($0, $1, $2) => { var target = document.querySelector(UTF8ToString($1)); if (target) { var data = $0; if (typeof(Module['SDL3']) === 'undefined') { Module['SDL3'] = {}; } var SDL3 = Module['SDL3']; var makeDropEventCStruct = function(event) { var ptr = 0; ptr = _SDL_malloc($2); if (ptr != 0) { var idx = ptr >> 2; var rect = target.getBoundingClientRect(); HEAP32[idx++] = event.clientX - rect.left; HEAP32[idx++] = event.clientY - rect.top; } return ptr; }; SDL3.eventHandlerDropDragover = function(event) { event.preventDefault(); var d = makeDropEventCStruct(event); if (d != 0) { _Emscripten_SendDragEvent(data, d); _SDL_free(d); } }; target.addEventListener("dragover", SDL3.eventHandlerDropDragover); SDL3.drop_count = 0; FS.mkdir("/tmp/filedrop"); SDL3.eventHandlerDropDrop = function(event) { event.preventDefault(); if (event.dataTransfer.types.includes("text/plain")) { let plain_text = stringToNewUTF8(event.dataTransfer.getData("text/plain")); _Emscripten_SendDragTextEvent(data, plain_text); _free(plain_text); } else if (event.dataTransfer.types.includes("Files")) { for (let i = 0; i < event.dataTransfer.files.length; i++) { const file = event.dataTransfer.files.item(i); const file_reader = new FileReader(); file_reader.readAsArrayBuffer(file); file_reader.onload = function(event) { const fs_dropdir = `/tmp/filedrop/${SDL3.drop_count}`; SDL3.drop_count += 1; const fs_filepath = `${fs_dropdir}/${file.name}`; const c_fs_filepath = stringToNewUTF8(fs_filepath); const contents_array8 = new Uint8Array(event.target.result); FS.mkdir(fs_dropdir); var stream = FS.open(fs_filepath, "w"); FS.write(stream, contents_array8, 0, contents_array8.length, 0); FS.close(stream); _Emscripten_SendDragFileEvent(data, c_fs_filepath); _free(c_fs_filepath); _Emscripten_SendDragCompleteEvent(data); }; } } _Emscripten_SendDragCompleteEvent(data); }; target.addEventListener("drop", SDL3.eventHandlerDropDrop); SDL3.eventHandlerDropDragend = function(event) { event.preventDefault(); _Emscripten_SendDragCompleteEvent(data); }; target.addEventListener("dragend", SDL3.eventHandlerDropDragend); target.addEventListener("dragleave", SDL3.eventHandlerDropDragend); } },  
 638025: ($0) => { var target = document.querySelector(UTF8ToString($0)); if (target) { var SDL3 = Module['SDL3']; target.removeEventListener("dragleave", SDL3.eventHandlerDropDragend); target.removeEventListener("dragend", SDL3.eventHandlerDropDragend); target.removeEventListener("drop", SDL3.eventHandlerDropDrop); SDL3.drop_count = undefined; function recursive_remove(dirpath) { FS.readdir(dirpath).forEach((filename) => { const p = `${dirpath}/${filename}`; const p_s = FS.stat(p); if (FS.isFile(p_s.mode)) { FS.unlink(p); } else if (FS.isDir(p)) { recursive_remove(p); } }); FS.rmdir(dirpath); }("/tmp/filedrop"); FS.rmdir("/tmp/filedrop"); target.removeEventListener("dragover", SDL3.eventHandlerDropDragover); SDL3.eventHandlerDropDragover = undefined; SDL3.eventHandlerDropDrop = undefined; SDL3.eventHandlerDropDragend = undefined; } },  
 638855: ($0) => { var target = document.querySelector(UTF8ToString($0)); if (target) { var SDL3 = Module['SDL3']; target.removeEventListener("pointerenter", SDL3.eventHandlerPointerEnter); target.removeEventListener("pointerleave", SDL3.eventHandlerPointerLeave); target.removeEventListener("pointercancel", SDL3.eventHandlerPointerLeave); target.removeEventListener("pointerdown", SDL3.eventHandlerPointerGeneric); target.removeEventListener("pointerup", SDL3.eventHandlerPointerGeneric); target.removeEventListener("pointermove", SDL3.eventHandlerPointerGeneric); SDL3.eventHandlerPointerEnter = undefined; SDL3.eventHandlerPointerLeave = undefined; SDL3.eventHandlerPointerGeneric = undefined; } },  
 639540: ($0, $1, $2, $3) => { var w = $0; var h = $1; var pixels = $2; var canvasId = UTF8ToString($3); var canvas = document.querySelector(canvasId); if (!Module['SDL3']) Module['SDL3'] = {}; var SDL3 = Module['SDL3']; if (SDL3.ctxCanvas !== canvas) { SDL3.ctx = Browser.createContext(canvas, false, true); SDL3.ctxCanvas = canvas; } if (SDL3.w !== w || SDL3.h !== h || SDL3.imageCtx !== SDL3.ctx) { SDL3.image = SDL3.ctx.createImageData(w, h); SDL3.w = w; SDL3.h = h; SDL3.imageCtx = SDL3.ctx; } var data = SDL3.image.data; var src = pixels / 4; var dst = 0; var num; if (SDL3.data32Data !== data) { SDL3.data32 = new Int32Array(data.buffer); SDL3.data8 = new Uint8Array(data.buffer); SDL3.data32Data = data; } var data32 = SDL3.data32; num = data32.length; data32.set(HEAP32.subarray(src, src + num)); var data8 = SDL3.data8; var i = 3; var j = i + 4*num; if (num % 8 == 0) { while (i < j) { data8[i] = 0xff; i = i + 4 | 0; data8[i] = 0xff; i = i + 4 | 0; data8[i] = 0xff; i = i + 4 | 0; data8[i] = 0xff; i = i + 4 | 0; data8[i] = 0xff; i = i + 4 | 0; data8[i] = 0xff; i = i + 4 | 0; data8[i] = 0xff; i = i + 4 | 0; data8[i] = 0xff; i = i + 4 | 0; } } else { while (i < j) { data8[i] = 0xff; i = i + 4 | 0; } } SDL3.ctx.putImageData(SDL3.image, 0, 0); },  
 640769: ($0, $1, $2, $3, $4) => { var w = $0; var h = $1; var hot_x = $2; var hot_y = $3; var pixels = $4; var canvas = document.createElement("canvas"); canvas.width = w; canvas.height = h; var ctx = canvas.getContext("2d"); var image = ctx.createImageData(w, h); var data = image.data; var src = pixels / 4; var data32 = new Int32Array(data.buffer); data32.set(HEAP32.subarray(src, src + data32.length)); ctx.putImageData(image, 0, 0); var url = hot_x === 0 && hot_y === 0 ? "url(" + canvas.toDataURL() + "), auto" : "url(" + canvas.toDataURL() + ") " + hot_x + " " + hot_y + ", auto"; var urlBuf = _SDL_malloc(url.length + 1); stringToUTF8(url, urlBuf, url.length + 1); return urlBuf; },  
 641427: ($0) => { if (Module['canvas']) { Module['canvas'].style['cursor'] = UTF8ToString($0); } },  
 641510: () => { if (Module['canvas']) { Module['canvas'].style['cursor'] = 'none'; } },  
 641579: () => { if (!window.matchMedia) { return -1; } if (window.matchMedia('(prefers-color-scheme: light)').matches) { return 0; } if (window.matchMedia('(prefers-color-scheme: dark)').matches) { return 1; } return -1; },  
 641788: () => { if (typeof(Module['SDL3']) !== 'undefined') { var SDL3 = Module['SDL3']; SDL3.themeChangedMatchMedia.removeEventListener('change', SDL3.eventHandlerThemeChanged); SDL3.themeChangedMatchMedia = undefined; SDL3.eventHandlerThemeChanged = undefined; } },  
 642041: () => { return window.innerWidth; },  
 642071: () => { return window.innerHeight; },  
 642102: ($0) => { Module['requestFullscreen'] = function(lockPointer, resizeCanvas) { _requestFullscreenThroughSDL($0); }; },  
 642211: () => { Module['requestFullscreen'] = function(lockPointer, resizeCanvas) {}; },  
 642285: () => { if (window.matchMedia) { if (typeof(Module['SDL3']) === 'undefined') { Module['SDL3'] = {}; } var SDL3 = Module['SDL3']; SDL3.eventHandlerThemeChanged = function(event) { _Emscripten_SendSystemThemeChangedEvent(); }; SDL3.themeChangedMatchMedia = window.matchMedia('(prefers-color-scheme: dark)'); SDL3.themeChangedMatchMedia.addEventListener('change', SDL3.eventHandlerThemeChanged); } }
};

// Imports from the Wasm binary.
var _IMG_Version = Module['_IMG_Version'] = makeInvalidEarlyAccess('_IMG_Version');
var _IMG_Load = Module['_IMG_Load'] = makeInvalidEarlyAccess('_IMG_Load');
var _SDL_CreateSurface = Module['_SDL_CreateSurface'] = makeInvalidEarlyAccess('_SDL_CreateSurface');
var _free = Module['_free'] = makeInvalidEarlyAccess('_free');
var _SDL_IOFromFile = Module['_SDL_IOFromFile'] = makeInvalidEarlyAccess('_SDL_IOFromFile');
var _SDL_strrchr = Module['_SDL_strrchr'] = makeInvalidEarlyAccess('_SDL_strrchr');
var _IMG_LoadTyped_IO = Module['_IMG_LoadTyped_IO'] = makeInvalidEarlyAccess('_IMG_LoadTyped_IO');
var _SDL_SetError = Module['_SDL_SetError'] = makeInvalidEarlyAccess('_SDL_SetError');
var _SDL_SeekIO = Module['_SDL_SeekIO'] = makeInvalidEarlyAccess('_SDL_SeekIO');
var _SDL_CloseIO = Module['_SDL_CloseIO'] = makeInvalidEarlyAccess('_SDL_CloseIO');
var _SDL_GetIOProperties = Module['_SDL_GetIOProperties'] = makeInvalidEarlyAccess('_SDL_GetIOProperties');
var _SDL_GetPointerProperty = Module['_SDL_GetPointerProperty'] = makeInvalidEarlyAccess('_SDL_GetPointerProperty');
var _IMG_isAVIF = Module['_IMG_isAVIF'] = makeInvalidEarlyAccess('_IMG_isAVIF');
var _IMG_isCUR = Module['_IMG_isCUR'] = makeInvalidEarlyAccess('_IMG_isCUR');
var _IMG_isICO = Module['_IMG_isICO'] = makeInvalidEarlyAccess('_IMG_isICO');
var _IMG_isBMP = Module['_IMG_isBMP'] = makeInvalidEarlyAccess('_IMG_isBMP');
var _IMG_isGIF = Module['_IMG_isGIF'] = makeInvalidEarlyAccess('_IMG_isGIF');
var _IMG_isJPG = Module['_IMG_isJPG'] = makeInvalidEarlyAccess('_IMG_isJPG');
var _IMG_isJXL = Module['_IMG_isJXL'] = makeInvalidEarlyAccess('_IMG_isJXL');
var _IMG_isLBM = Module['_IMG_isLBM'] = makeInvalidEarlyAccess('_IMG_isLBM');
var _IMG_isPCX = Module['_IMG_isPCX'] = makeInvalidEarlyAccess('_IMG_isPCX');
var _IMG_isPNG = Module['_IMG_isPNG'] = makeInvalidEarlyAccess('_IMG_isPNG');
var _IMG_isPNM = Module['_IMG_isPNM'] = makeInvalidEarlyAccess('_IMG_isPNM');
var _IMG_isSVG = Module['_IMG_isSVG'] = makeInvalidEarlyAccess('_IMG_isSVG');
var _IMG_isTIF = Module['_IMG_isTIF'] = makeInvalidEarlyAccess('_IMG_isTIF');
var _IMG_isXCF = Module['_IMG_isXCF'] = makeInvalidEarlyAccess('_IMG_isXCF');
var _IMG_isXPM = Module['_IMG_isXPM'] = makeInvalidEarlyAccess('_IMG_isXPM');
var _IMG_isXV = Module['_IMG_isXV'] = makeInvalidEarlyAccess('_IMG_isXV');
var _IMG_isWEBP = Module['_IMG_isWEBP'] = makeInvalidEarlyAccess('_IMG_isWEBP');
var _IMG_isQOI = Module['_IMG_isQOI'] = makeInvalidEarlyAccess('_IMG_isQOI');
var _SDL_strcasecmp = Module['_SDL_strcasecmp'] = makeInvalidEarlyAccess('_SDL_strcasecmp');
var _IMG_Load_IO = Module['_IMG_Load_IO'] = makeInvalidEarlyAccess('_IMG_Load_IO');
var _IMG_LoadTexture = Module['_IMG_LoadTexture'] = makeInvalidEarlyAccess('_IMG_LoadTexture');
var _SDL_CreateTextureFromSurface = Module['_SDL_CreateTextureFromSurface'] = makeInvalidEarlyAccess('_SDL_CreateTextureFromSurface');
var _SDL_DestroySurface = Module['_SDL_DestroySurface'] = makeInvalidEarlyAccess('_SDL_DestroySurface');
var _IMG_LoadTexture_IO = Module['_IMG_LoadTexture_IO'] = makeInvalidEarlyAccess('_IMG_LoadTexture_IO');
var _IMG_LoadTextureTyped_IO = Module['_IMG_LoadTextureTyped_IO'] = makeInvalidEarlyAccess('_IMG_LoadTextureTyped_IO');
var _IMG_LoadAnimation = Module['_IMG_LoadAnimation'] = makeInvalidEarlyAccess('_IMG_LoadAnimation');
var _IMG_LoadAnimationTyped_IO = Module['_IMG_LoadAnimationTyped_IO'] = makeInvalidEarlyAccess('_IMG_LoadAnimationTyped_IO');
var _SDL_malloc = Module['_SDL_malloc'] = makeInvalidEarlyAccess('_SDL_malloc');
var _SDL_calloc = Module['_SDL_calloc'] = makeInvalidEarlyAccess('_SDL_calloc');
var _SDL_free = Module['_SDL_free'] = makeInvalidEarlyAccess('_SDL_free');
var _IMG_LoadAnimation_IO = Module['_IMG_LoadAnimation_IO'] = makeInvalidEarlyAccess('_IMG_LoadAnimation_IO');
var _IMG_FreeAnimation = Module['_IMG_FreeAnimation'] = makeInvalidEarlyAccess('_IMG_FreeAnimation');
var _IMG_LoadTGA_IO = Module['_IMG_LoadTGA_IO'] = makeInvalidEarlyAccess('_IMG_LoadTGA_IO');
var _IMG_LoadAVIF_IO = Module['_IMG_LoadAVIF_IO'] = makeInvalidEarlyAccess('_IMG_LoadAVIF_IO');
var _IMG_LoadCUR_IO = Module['_IMG_LoadCUR_IO'] = makeInvalidEarlyAccess('_IMG_LoadCUR_IO');
var _IMG_LoadICO_IO = Module['_IMG_LoadICO_IO'] = makeInvalidEarlyAccess('_IMG_LoadICO_IO');
var _IMG_LoadBMP_IO = Module['_IMG_LoadBMP_IO'] = makeInvalidEarlyAccess('_IMG_LoadBMP_IO');
var _IMG_LoadGIF_IO = Module['_IMG_LoadGIF_IO'] = makeInvalidEarlyAccess('_IMG_LoadGIF_IO');
var _IMG_LoadJPG_IO = Module['_IMG_LoadJPG_IO'] = makeInvalidEarlyAccess('_IMG_LoadJPG_IO');
var _IMG_LoadJXL_IO = Module['_IMG_LoadJXL_IO'] = makeInvalidEarlyAccess('_IMG_LoadJXL_IO');
var _IMG_LoadLBM_IO = Module['_IMG_LoadLBM_IO'] = makeInvalidEarlyAccess('_IMG_LoadLBM_IO');
var _IMG_LoadPCX_IO = Module['_IMG_LoadPCX_IO'] = makeInvalidEarlyAccess('_IMG_LoadPCX_IO');
var _IMG_LoadPNG_IO = Module['_IMG_LoadPNG_IO'] = makeInvalidEarlyAccess('_IMG_LoadPNG_IO');
var _IMG_LoadPNM_IO = Module['_IMG_LoadPNM_IO'] = makeInvalidEarlyAccess('_IMG_LoadPNM_IO');
var _IMG_LoadSVG_IO = Module['_IMG_LoadSVG_IO'] = makeInvalidEarlyAccess('_IMG_LoadSVG_IO');
var _IMG_LoadTIF_IO = Module['_IMG_LoadTIF_IO'] = makeInvalidEarlyAccess('_IMG_LoadTIF_IO');
var _IMG_LoadXCF_IO = Module['_IMG_LoadXCF_IO'] = makeInvalidEarlyAccess('_IMG_LoadXCF_IO');
var _IMG_LoadXPM_IO = Module['_IMG_LoadXPM_IO'] = makeInvalidEarlyAccess('_IMG_LoadXPM_IO');
var _IMG_LoadXV_IO = Module['_IMG_LoadXV_IO'] = makeInvalidEarlyAccess('_IMG_LoadXV_IO');
var _IMG_LoadWEBP_IO = Module['_IMG_LoadWEBP_IO'] = makeInvalidEarlyAccess('_IMG_LoadWEBP_IO');
var _IMG_LoadQOI_IO = Module['_IMG_LoadQOI_IO'] = makeInvalidEarlyAccess('_IMG_LoadQOI_IO');
var _IMG_LoadGIFAnimation_IO = Module['_IMG_LoadGIFAnimation_IO'] = makeInvalidEarlyAccess('_IMG_LoadGIFAnimation_IO');
var _IMG_LoadWEBPAnimation_IO = Module['_IMG_LoadWEBPAnimation_IO'] = makeInvalidEarlyAccess('_IMG_LoadWEBPAnimation_IO');
var _IMG_SaveAVIF = Module['_IMG_SaveAVIF'] = makeInvalidEarlyAccess('_IMG_SaveAVIF');
var _IMG_SaveAVIF_IO = Module['_IMG_SaveAVIF_IO'] = makeInvalidEarlyAccess('_IMG_SaveAVIF_IO');
var _SDL_TellIO = Module['_SDL_TellIO'] = makeInvalidEarlyAccess('_SDL_TellIO');
var _SDL_ReadIO = Module['_SDL_ReadIO'] = makeInvalidEarlyAccess('_SDL_ReadIO');
var _SDL_strncmp = Module['_SDL_strncmp'] = makeInvalidEarlyAccess('_SDL_strncmp');
var _SDL_ReadU16LE = Module['_SDL_ReadU16LE'] = makeInvalidEarlyAccess('_SDL_ReadU16LE');
var _SDL_LoadBMP_IO = Module['_SDL_LoadBMP_IO'] = makeInvalidEarlyAccess('_SDL_LoadBMP_IO');
var _SDL_ReadU8 = Module['_SDL_ReadU8'] = makeInvalidEarlyAccess('_SDL_ReadU8');
var _SDL_ReadU32LE = Module['_SDL_ReadU32LE'] = makeInvalidEarlyAccess('_SDL_ReadU32LE');
var _SDL_ReadS32LE = Module['_SDL_ReadS32LE'] = makeInvalidEarlyAccess('_SDL_ReadS32LE');
var _SDL_GetSurfaceProperties = Module['_SDL_GetSurfaceProperties'] = makeInvalidEarlyAccess('_SDL_GetSurfaceProperties');
var _SDL_SetNumberProperty = Module['_SDL_SetNumberProperty'] = makeInvalidEarlyAccess('_SDL_SetNumberProperty');
var _SDL_strcmp = Module['_SDL_strcmp'] = makeInvalidEarlyAccess('_SDL_strcmp');
var _SDL_SetSurfaceColorKey = Module['_SDL_SetSurfaceColorKey'] = makeInvalidEarlyAccess('_SDL_SetSurfaceColorKey');
var _SDL_realloc = Module['_SDL_realloc'] = makeInvalidEarlyAccess('_SDL_realloc');
var _SDL_SurfaceHasColorKey = Module['_SDL_SurfaceHasColorKey'] = makeInvalidEarlyAccess('_SDL_SurfaceHasColorKey');
var _SDL_ConvertSurface = Module['_SDL_ConvertSurface'] = makeInvalidEarlyAccess('_SDL_ConvertSurface');
var _SDL_MapSurfaceRGBA = Module['_SDL_MapSurfaceRGBA'] = makeInvalidEarlyAccess('_SDL_MapSurfaceRGBA');
var _SDL_FillSurfaceRect = Module['_SDL_FillSurfaceRect'] = makeInvalidEarlyAccess('_SDL_FillSurfaceRect');
var _SDL_BlitSurface = Module['_SDL_BlitSurface'] = makeInvalidEarlyAccess('_SDL_BlitSurface');
var _SDL_DuplicateSurface = Module['_SDL_DuplicateSurface'] = makeInvalidEarlyAccess('_SDL_DuplicateSurface');
var _SDL_memcmp = Module['_SDL_memcmp'] = makeInvalidEarlyAccess('_SDL_memcmp');
var _SDL_CreateSurfacePalette = Module['_SDL_CreateSurfacePalette'] = makeInvalidEarlyAccess('_SDL_CreateSurfacePalette');
var _SDL_Log = Module['_SDL_Log'] = makeInvalidEarlyAccess('_SDL_Log');
var _IMG_SaveJPG = Module['_IMG_SaveJPG'] = makeInvalidEarlyAccess('_IMG_SaveJPG');
var _IMG_SaveJPG_IO = Module['_IMG_SaveJPG_IO'] = makeInvalidEarlyAccess('_IMG_SaveJPG_IO');
var _SDL_floorf = Module['_SDL_floorf'] = makeInvalidEarlyAccess('_SDL_floorf');
var _SDL_WriteIO = Module['_SDL_WriteIO'] = makeInvalidEarlyAccess('_SDL_WriteIO');
var _IMG_SavePNG = Module['_IMG_SavePNG'] = makeInvalidEarlyAccess('_IMG_SavePNG');
var _IMG_SavePNG_IO = Module['_IMG_SavePNG_IO'] = makeInvalidEarlyAccess('_IMG_SavePNG_IO');
var _SDL_CreatePalette = Module['_SDL_CreatePalette'] = makeInvalidEarlyAccess('_SDL_CreatePalette');
var _SDL_SetSurfacePalette = Module['_SDL_SetSurfacePalette'] = makeInvalidEarlyAccess('_SDL_SetSurfacePalette');
var _SDL_DestroyPalette = Module['_SDL_DestroyPalette'] = makeInvalidEarlyAccess('_SDL_DestroyPalette');
var _SDL_isspace = Module['_SDL_isspace'] = makeInvalidEarlyAccess('_SDL_isspace');
var _SDL_isdigit = Module['_SDL_isdigit'] = makeInvalidEarlyAccess('_SDL_isdigit');
var _SDL_LoadFile_IO = Module['_SDL_LoadFile_IO'] = makeInvalidEarlyAccess('_SDL_LoadFile_IO');
var _SDL_CreateSurfaceFrom = Module['_SDL_CreateSurfaceFrom'] = makeInvalidEarlyAccess('_SDL_CreateSurfaceFrom');
var _SDL_memset = Module['_SDL_memset'] = makeInvalidEarlyAccess('_SDL_memset');
var _SDL_SetSurfaceBlendMode = Module['_SDL_SetSurfaceBlendMode'] = makeInvalidEarlyAccess('_SDL_SetSurfaceBlendMode');
var _SDL_GetIOStatus = Module['_SDL_GetIOStatus'] = makeInvalidEarlyAccess('_SDL_GetIOStatus');
var _SDL_strstr = Module['_SDL_strstr'] = makeInvalidEarlyAccess('_SDL_strstr');
var _IMG_LoadSizedSVG_IO = Module['_IMG_LoadSizedSVG_IO'] = makeInvalidEarlyAccess('_IMG_LoadSizedSVG_IO');
var _SDL_strchr = Module['_SDL_strchr'] = makeInvalidEarlyAccess('_SDL_strchr');
var _SDL_strlen = Module['_SDL_strlen'] = makeInvalidEarlyAccess('_SDL_strlen');
var _SDL_ceilf = Module['_SDL_ceilf'] = makeInvalidEarlyAccess('_SDL_ceilf');
var _SDL_qsort = Module['_SDL_qsort'] = makeInvalidEarlyAccess('_SDL_qsort');
var _SDL_sqrtf = Module['_SDL_sqrtf'] = makeInvalidEarlyAccess('_SDL_sqrtf');
var _SDL_fmodf = Module['_SDL_fmodf'] = makeInvalidEarlyAccess('_SDL_fmodf');
var _SDL_fabsf = Module['_SDL_fabsf'] = makeInvalidEarlyAccess('_SDL_fabsf');
var _SDL_sinf = Module['_SDL_sinf'] = makeInvalidEarlyAccess('_SDL_sinf');
var _SDL_cosf = Module['_SDL_cosf'] = makeInvalidEarlyAccess('_SDL_cosf');
var _SDL_acosf = Module['_SDL_acosf'] = makeInvalidEarlyAccess('_SDL_acosf');
var _SDL_strlcpy = Module['_SDL_strlcpy'] = makeInvalidEarlyAccess('_SDL_strlcpy');
var _SDL_strtoll = Module['_SDL_strtoll'] = makeInvalidEarlyAccess('_SDL_strtoll');
var _SDL_pow = Module['_SDL_pow'] = makeInvalidEarlyAccess('_SDL_pow');
var _SDL_strtol = Module['_SDL_strtol'] = makeInvalidEarlyAccess('_SDL_strtol');
var _SDL_tanf = Module['_SDL_tanf'] = makeInvalidEarlyAccess('_SDL_tanf');
var _SDL_sscanf = Module['_SDL_sscanf'] = makeInvalidEarlyAccess('_SDL_sscanf');
var _SDL_roundf = Module['_SDL_roundf'] = makeInvalidEarlyAccess('_SDL_roundf');
var _SDL_fabs = Module['_SDL_fabs'] = makeInvalidEarlyAccess('_SDL_fabs');
var _SDL_sqrt = Module['_SDL_sqrt'] = makeInvalidEarlyAccess('_SDL_sqrt');
var _SDL_atan2f = Module['_SDL_atan2f'] = makeInvalidEarlyAccess('_SDL_atan2f');
var _SDL_ReadU32BE = Module['_SDL_ReadU32BE'] = makeInvalidEarlyAccess('_SDL_ReadU32BE');
var _SDL_ReadS32BE = Module['_SDL_ReadS32BE'] = makeInvalidEarlyAccess('_SDL_ReadS32BE');
var _SDL_GetIOSize = Module['_SDL_GetIOSize'] = makeInvalidEarlyAccess('_SDL_GetIOSize');
var _SDL_strncasecmp = Module['_SDL_strncasecmp'] = makeInvalidEarlyAccess('_SDL_strncasecmp');
var _IMG_ReadXPMFromArray = Module['_IMG_ReadXPMFromArray'] = makeInvalidEarlyAccess('_IMG_ReadXPMFromArray');
var _IMG_ReadXPMFromArrayToRGB888 = Module['_IMG_ReadXPMFromArrayToRGB888'] = makeInvalidEarlyAccess('_IMG_ReadXPMFromArrayToRGB888');
var _SDL_CreateRWLock = Module['_SDL_CreateRWLock'] = makeInvalidEarlyAccess('_SDL_CreateRWLock');
var _SDL_DestroyRWLock = Module['_SDL_DestroyRWLock'] = makeInvalidEarlyAccess('_SDL_DestroyRWLock');
var _SDL_LockRWLockForWriting = Module['_SDL_LockRWLockForWriting'] = makeInvalidEarlyAccess('_SDL_LockRWLockForWriting');
var _SDL_UnlockRWLock = Module['_SDL_UnlockRWLock'] = makeInvalidEarlyAccess('_SDL_UnlockRWLock');
var _SDL_LockRWLockForReading = Module['_SDL_LockRWLockForReading'] = makeInvalidEarlyAccess('_SDL_LockRWLockForReading');
var _SDL_murmur3_32 = Module['_SDL_murmur3_32'] = makeInvalidEarlyAccess('_SDL_murmur3_32');
var _TTF_DestroyGPUTextEngine = Module['_TTF_DestroyGPUTextEngine'] = makeInvalidEarlyAccess('_TTF_DestroyGPUTextEngine');
var _TTF_GetFontGeneration = Module['_TTF_GetFontGeneration'] = makeInvalidEarlyAccess('_TTF_GetFontGeneration');
var _TTF_GetGlyphImageForIndex = Module['_TTF_GetGlyphImageForIndex'] = makeInvalidEarlyAccess('_TTF_GetGlyphImageForIndex');
var _SDL_qsort_r = Module['_SDL_qsort_r'] = makeInvalidEarlyAccess('_SDL_qsort_r');
var _SDL_ReleaseGPUTexture = Module['_SDL_ReleaseGPUTexture'] = makeInvalidEarlyAccess('_SDL_ReleaseGPUTexture');
var _TTF_CreateGPUTextEngine = Module['_TTF_CreateGPUTextEngine'] = makeInvalidEarlyAccess('_TTF_CreateGPUTextEngine');
var _SDL_CreateProperties = Module['_SDL_CreateProperties'] = makeInvalidEarlyAccess('_SDL_CreateProperties');
var _SDL_SetPointerProperty = Module['_SDL_SetPointerProperty'] = makeInvalidEarlyAccess('_SDL_SetPointerProperty');
var _TTF_CreateGPUTextEngineWithProperties = Module['_TTF_CreateGPUTextEngineWithProperties'] = makeInvalidEarlyAccess('_TTF_CreateGPUTextEngineWithProperties');
var _SDL_GetNumberProperty = Module['_SDL_GetNumberProperty'] = makeInvalidEarlyAccess('_SDL_GetNumberProperty');
var _TTF_GetGPUTextDrawData = Module['_TTF_GetGPUTextDrawData'] = makeInvalidEarlyAccess('_TTF_GetGPUTextDrawData');
var _TTF_UpdateText = Module['_TTF_UpdateText'] = makeInvalidEarlyAccess('_TTF_UpdateText');
var _TTF_SetGPUTextEngineWinding = Module['_TTF_SetGPUTextEngineWinding'] = makeInvalidEarlyAccess('_TTF_SetGPUTextEngineWinding');
var _TTF_GetGPUTextEngineWinding = Module['_TTF_GetGPUTextEngineWinding'] = makeInvalidEarlyAccess('_TTF_GetGPUTextEngineWinding');
var _SDL_CreateGPUTexture = Module['_SDL_CreateGPUTexture'] = makeInvalidEarlyAccess('_SDL_CreateGPUTexture');
var _SDL_AcquireGPUCommandBuffer = Module['_SDL_AcquireGPUCommandBuffer'] = makeInvalidEarlyAccess('_SDL_AcquireGPUCommandBuffer');
var _SDL_BeginGPURenderPass = Module['_SDL_BeginGPURenderPass'] = makeInvalidEarlyAccess('_SDL_BeginGPURenderPass');
var _SDL_EndGPURenderPass = Module['_SDL_EndGPURenderPass'] = makeInvalidEarlyAccess('_SDL_EndGPURenderPass');
var _SDL_SubmitGPUCommandBuffer = Module['_SDL_SubmitGPUCommandBuffer'] = makeInvalidEarlyAccess('_SDL_SubmitGPUCommandBuffer');
var _SDL_CreateGPUTransferBuffer = Module['_SDL_CreateGPUTransferBuffer'] = makeInvalidEarlyAccess('_SDL_CreateGPUTransferBuffer');
var _SDL_MapGPUTransferBuffer = Module['_SDL_MapGPUTransferBuffer'] = makeInvalidEarlyAccess('_SDL_MapGPUTransferBuffer');
var _SDL_UnmapGPUTransferBuffer = Module['_SDL_UnmapGPUTransferBuffer'] = makeInvalidEarlyAccess('_SDL_UnmapGPUTransferBuffer');
var _SDL_BeginGPUCopyPass = Module['_SDL_BeginGPUCopyPass'] = makeInvalidEarlyAccess('_SDL_BeginGPUCopyPass');
var _SDL_UploadToGPUTexture = Module['_SDL_UploadToGPUTexture'] = makeInvalidEarlyAccess('_SDL_UploadToGPUTexture');
var _SDL_EndGPUCopyPass = Module['_SDL_EndGPUCopyPass'] = makeInvalidEarlyAccess('_SDL_EndGPUCopyPass');
var _SDL_ReleaseGPUTransferBuffer = Module['_SDL_ReleaseGPUTransferBuffer'] = makeInvalidEarlyAccess('_SDL_ReleaseGPUTransferBuffer');
var _TTF_CreateRendererTextEngine = Module['_TTF_CreateRendererTextEngine'] = makeInvalidEarlyAccess('_TTF_CreateRendererTextEngine');
var _TTF_CreateRendererTextEngineWithProperties = Module['_TTF_CreateRendererTextEngineWithProperties'] = makeInvalidEarlyAccess('_TTF_CreateRendererTextEngineWithProperties');
var _TTF_DestroyRendererTextEngine = Module['_TTF_DestroyRendererTextEngine'] = makeInvalidEarlyAccess('_TTF_DestroyRendererTextEngine');
var _TTF_DrawRendererText = Module['_TTF_DrawRendererText'] = makeInvalidEarlyAccess('_TTF_DrawRendererText');
var _SDL_RenderGeometryRaw = Module['_SDL_RenderGeometryRaw'] = makeInvalidEarlyAccess('_SDL_RenderGeometryRaw');
var _SDL_DestroyTexture = Module['_SDL_DestroyTexture'] = makeInvalidEarlyAccess('_SDL_DestroyTexture');
var _SDL_CreateTexture = Module['_SDL_CreateTexture'] = makeInvalidEarlyAccess('_SDL_CreateTexture');
var _SDL_SetTextureScaleMode = Module['_SDL_SetTextureScaleMode'] = makeInvalidEarlyAccess('_SDL_SetTextureScaleMode');
var _SDL_LockTexture = Module['_SDL_LockTexture'] = makeInvalidEarlyAccess('_SDL_LockTexture');
var _SDL_UnlockTexture = Module['_SDL_UnlockTexture'] = makeInvalidEarlyAccess('_SDL_UnlockTexture');
var _TTF_CreateSurfaceTextEngine = Module['_TTF_CreateSurfaceTextEngine'] = makeInvalidEarlyAccess('_TTF_CreateSurfaceTextEngine');
var _TTF_DestroySurfaceTextEngine = Module['_TTF_DestroySurfaceTextEngine'] = makeInvalidEarlyAccess('_TTF_DestroySurfaceTextEngine');
var _TTF_DrawSurfaceText = Module['_TTF_DrawSurfaceText'] = makeInvalidEarlyAccess('_TTF_DrawSurfaceText');
var _SDL_SetSurfaceColorMod = Module['_SDL_SetSurfaceColorMod'] = makeInvalidEarlyAccess('_SDL_SetSurfaceColorMod');
var _SDL_SetSurfaceAlphaMod = Module['_SDL_SetSurfaceAlphaMod'] = makeInvalidEarlyAccess('_SDL_SetSurfaceAlphaMod');
var _TTF_Version = Module['_TTF_Version'] = makeInvalidEarlyAccess('_TTF_Version');
var _TTF_Init = Module['_TTF_Init'] = makeInvalidEarlyAccess('_TTF_Init');
var _SDL_AddAtomicInt = Module['_SDL_AddAtomicInt'] = makeInvalidEarlyAccess('_SDL_AddAtomicInt');
var _SDL_ShouldInit = Module['_SDL_ShouldInit'] = makeInvalidEarlyAccess('_SDL_ShouldInit');
var _plutosvg_ft_svg_hooks = Module['_plutosvg_ft_svg_hooks'] = makeInvalidEarlyAccess('_plutosvg_ft_svg_hooks');
var _SDL_CreateMutex = Module['_SDL_CreateMutex'] = makeInvalidEarlyAccess('_SDL_CreateMutex');
var _SDL_SetInitialized = Module['_SDL_SetInitialized'] = makeInvalidEarlyAccess('_SDL_SetInitialized');
var _TTF_GetFreeTypeVersion = Module['_TTF_GetFreeTypeVersion'] = makeInvalidEarlyAccess('_TTF_GetFreeTypeVersion');
var _SDL_LockMutex = Module['_SDL_LockMutex'] = makeInvalidEarlyAccess('_SDL_LockMutex');
var _SDL_UnlockMutex = Module['_SDL_UnlockMutex'] = makeInvalidEarlyAccess('_SDL_UnlockMutex');
var _TTF_GetHarfBuzzVersion = Module['_TTF_GetHarfBuzzVersion'] = makeInvalidEarlyAccess('_TTF_GetHarfBuzzVersion');
var _TTF_OpenFontWithProperties = Module['_TTF_OpenFontWithProperties'] = makeInvalidEarlyAccess('_TTF_OpenFontWithProperties');
var _SDL_GetStringProperty = Module['_SDL_GetStringProperty'] = makeInvalidEarlyAccess('_SDL_GetStringProperty');
var _SDL_GetBooleanProperty = Module['_SDL_GetBooleanProperty'] = makeInvalidEarlyAccess('_SDL_GetBooleanProperty');
var _SDL_GetFloatProperty = Module['_SDL_GetFloatProperty'] = makeInvalidEarlyAccess('_SDL_GetFloatProperty');
var _SDL_strdup = Module['_SDL_strdup'] = makeInvalidEarlyAccess('_SDL_strdup');
var _TTF_SetFontKerning = Module['_TTF_SetFontKerning'] = makeInvalidEarlyAccess('_TTF_SetFontKerning');
var _TTF_SetFontSizeDPI = Module['_TTF_SetFontSizeDPI'] = makeInvalidEarlyAccess('_TTF_SetFontSizeDPI');
var _TTF_CloseFont = Module['_TTF_CloseFont'] = makeInvalidEarlyAccess('_TTF_CloseFont');
var _SDL_DestroyProperties = Module['_SDL_DestroyProperties'] = makeInvalidEarlyAccess('_SDL_DestroyProperties');
var _TTF_OpenFont = Module['_TTF_OpenFont'] = makeInvalidEarlyAccess('_TTF_OpenFont');
var _SDL_SetStringProperty = Module['_SDL_SetStringProperty'] = makeInvalidEarlyAccess('_SDL_SetStringProperty');
var _SDL_SetFloatProperty = Module['_SDL_SetFloatProperty'] = makeInvalidEarlyAccess('_SDL_SetFloatProperty');
var _TTF_OpenFontIO = Module['_TTF_OpenFontIO'] = makeInvalidEarlyAccess('_TTF_OpenFontIO');
var _SDL_SetBooleanProperty = Module['_SDL_SetBooleanProperty'] = makeInvalidEarlyAccess('_SDL_SetBooleanProperty');
var _TTF_CopyFont = Module['_TTF_CopyFont'] = makeInvalidEarlyAccess('_TTF_CopyFont');
var _TTF_GetFontProperties = Module['_TTF_GetFontProperties'] = makeInvalidEarlyAccess('_TTF_GetFontProperties');
var _TTF_AddFallbackFont = Module['_TTF_AddFallbackFont'] = makeInvalidEarlyAccess('_TTF_AddFallbackFont');
var _TTF_RemoveFallbackFont = Module['_TTF_RemoveFallbackFont'] = makeInvalidEarlyAccess('_TTF_RemoveFallbackFont');
var _TTF_ClearFallbackFonts = Module['_TTF_ClearFallbackFonts'] = makeInvalidEarlyAccess('_TTF_ClearFallbackFonts');
var _TTF_FontHasGlyph = Module['_TTF_FontHasGlyph'] = makeInvalidEarlyAccess('_TTF_FontHasGlyph');
var _TTF_GetGlyphImage = Module['_TTF_GetGlyphImage'] = makeInvalidEarlyAccess('_TTF_GetGlyphImage');
var _TTF_GetGlyphMetrics = Module['_TTF_GetGlyphMetrics'] = makeInvalidEarlyAccess('_TTF_GetGlyphMetrics');
var _TTF_GetGlyphKerning = Module['_TTF_GetGlyphKerning'] = makeInvalidEarlyAccess('_TTF_GetGlyphKerning');
var _TTF_GetStringSize = Module['_TTF_GetStringSize'] = makeInvalidEarlyAccess('_TTF_GetStringSize');
var _TTF_MeasureString = Module['_TTF_MeasureString'] = makeInvalidEarlyAccess('_TTF_MeasureString');
var _TTF_RenderText_Solid = Module['_TTF_RenderText_Solid'] = makeInvalidEarlyAccess('_TTF_RenderText_Solid');
var _SDL_GetSurfacePalette = Module['_SDL_GetSurfacePalette'] = makeInvalidEarlyAccess('_SDL_GetSurfacePalette');
var _TTF_RenderGlyph_Solid = Module['_TTF_RenderGlyph_Solid'] = makeInvalidEarlyAccess('_TTF_RenderGlyph_Solid');
var _SDL_UCS4ToUTF8 = Module['_SDL_UCS4ToUTF8'] = makeInvalidEarlyAccess('_SDL_UCS4ToUTF8');
var _TTF_RenderText_Shaded = Module['_TTF_RenderText_Shaded'] = makeInvalidEarlyAccess('_TTF_RenderText_Shaded');
var _TTF_RenderGlyph_Shaded = Module['_TTF_RenderGlyph_Shaded'] = makeInvalidEarlyAccess('_TTF_RenderGlyph_Shaded');
var _TTF_RenderText_Blended = Module['_TTF_RenderText_Blended'] = makeInvalidEarlyAccess('_TTF_RenderText_Blended');
var _TTF_RenderGlyph_Blended = Module['_TTF_RenderGlyph_Blended'] = makeInvalidEarlyAccess('_TTF_RenderGlyph_Blended');
var _TTF_RenderText_LCD = Module['_TTF_RenderText_LCD'] = makeInvalidEarlyAccess('_TTF_RenderText_LCD');
var _TTF_RenderGlyph_LCD = Module['_TTF_RenderGlyph_LCD'] = makeInvalidEarlyAccess('_TTF_RenderGlyph_LCD');
var _TTF_GetStringSizeWrapped = Module['_TTF_GetStringSizeWrapped'] = makeInvalidEarlyAccess('_TTF_GetStringSizeWrapped');
var _SDL_StepUTF8 = Module['_SDL_StepUTF8'] = makeInvalidEarlyAccess('_SDL_StepUTF8');
var _TTF_RenderText_Solid_Wrapped = Module['_TTF_RenderText_Solid_Wrapped'] = makeInvalidEarlyAccess('_TTF_RenderText_Solid_Wrapped');
var _SDL_memset4 = Module['_SDL_memset4'] = makeInvalidEarlyAccess('_SDL_memset4');
var _TTF_RenderText_Shaded_Wrapped = Module['_TTF_RenderText_Shaded_Wrapped'] = makeInvalidEarlyAccess('_TTF_RenderText_Shaded_Wrapped');
var _TTF_RenderText_Blended_Wrapped = Module['_TTF_RenderText_Blended_Wrapped'] = makeInvalidEarlyAccess('_TTF_RenderText_Blended_Wrapped');
var _TTF_RenderText_LCD_Wrapped = Module['_TTF_RenderText_LCD_Wrapped'] = makeInvalidEarlyAccess('_TTF_RenderText_LCD_Wrapped');
var _TTF_CreateText = Module['_TTF_CreateText'] = makeInvalidEarlyAccess('_TTF_CreateText');
var _TTF_GetTextProperties = Module['_TTF_GetTextProperties'] = makeInvalidEarlyAccess('_TTF_GetTextProperties');
var _TTF_SetTextEngine = Module['_TTF_SetTextEngine'] = makeInvalidEarlyAccess('_TTF_SetTextEngine');
var _TTF_GetTextEngine = Module['_TTF_GetTextEngine'] = makeInvalidEarlyAccess('_TTF_GetTextEngine');
var _TTF_SetTextFont = Module['_TTF_SetTextFont'] = makeInvalidEarlyAccess('_TTF_SetTextFont');
var _TTF_GetTextFont = Module['_TTF_GetTextFont'] = makeInvalidEarlyAccess('_TTF_GetTextFont');
var _TTF_SetTextDirection = Module['_TTF_SetTextDirection'] = makeInvalidEarlyAccess('_TTF_SetTextDirection');
var _TTF_GetTextDirection = Module['_TTF_GetTextDirection'] = makeInvalidEarlyAccess('_TTF_GetTextDirection');
var _TTF_GetFontDirection = Module['_TTF_GetFontDirection'] = makeInvalidEarlyAccess('_TTF_GetFontDirection');
var _TTF_SetTextScript = Module['_TTF_SetTextScript'] = makeInvalidEarlyAccess('_TTF_SetTextScript');
var _TTF_GetTextScript = Module['_TTF_GetTextScript'] = makeInvalidEarlyAccess('_TTF_GetTextScript');
var _TTF_GetFontScript = Module['_TTF_GetFontScript'] = makeInvalidEarlyAccess('_TTF_GetFontScript');
var _TTF_SetTextColor = Module['_TTF_SetTextColor'] = makeInvalidEarlyAccess('_TTF_SetTextColor');
var _TTF_SetTextColorFloat = Module['_TTF_SetTextColorFloat'] = makeInvalidEarlyAccess('_TTF_SetTextColorFloat');
var _TTF_GetTextColor = Module['_TTF_GetTextColor'] = makeInvalidEarlyAccess('_TTF_GetTextColor');
var _TTF_GetTextColorFloat = Module['_TTF_GetTextColorFloat'] = makeInvalidEarlyAccess('_TTF_GetTextColorFloat');
var _TTF_SetTextPosition = Module['_TTF_SetTextPosition'] = makeInvalidEarlyAccess('_TTF_SetTextPosition');
var _TTF_GetTextPosition = Module['_TTF_GetTextPosition'] = makeInvalidEarlyAccess('_TTF_GetTextPosition');
var _TTF_SetTextWrapWidth = Module['_TTF_SetTextWrapWidth'] = makeInvalidEarlyAccess('_TTF_SetTextWrapWidth');
var _TTF_GetTextWrapWidth = Module['_TTF_GetTextWrapWidth'] = makeInvalidEarlyAccess('_TTF_GetTextWrapWidth');
var _TTF_SetTextWrapWhitespaceVisible = Module['_TTF_SetTextWrapWhitespaceVisible'] = makeInvalidEarlyAccess('_TTF_SetTextWrapWhitespaceVisible');
var _TTF_TextWrapWhitespaceVisible = Module['_TTF_TextWrapWhitespaceVisible'] = makeInvalidEarlyAccess('_TTF_TextWrapWhitespaceVisible');
var _TTF_SetTextString = Module['_TTF_SetTextString'] = makeInvalidEarlyAccess('_TTF_SetTextString');
var _TTF_InsertTextString = Module['_TTF_InsertTextString'] = makeInvalidEarlyAccess('_TTF_InsertTextString');
var _TTF_AppendTextString = Module['_TTF_AppendTextString'] = makeInvalidEarlyAccess('_TTF_AppendTextString');
var _TTF_DeleteTextString = Module['_TTF_DeleteTextString'] = makeInvalidEarlyAccess('_TTF_DeleteTextString');
var _TTF_GetTextSize = Module['_TTF_GetTextSize'] = makeInvalidEarlyAccess('_TTF_GetTextSize');
var _SDL_GetRectUnion = Module['_SDL_GetRectUnion'] = makeInvalidEarlyAccess('_SDL_GetRectUnion');
var _TTF_GetTextSubString = Module['_TTF_GetTextSubString'] = makeInvalidEarlyAccess('_TTF_GetTextSubString');
var _TTF_GetTextSubStringForLine = Module['_TTF_GetTextSubStringForLine'] = makeInvalidEarlyAccess('_TTF_GetTextSubStringForLine');
var _TTF_GetTextSubStringsForRange = Module['_TTF_GetTextSubStringsForRange'] = makeInvalidEarlyAccess('_TTF_GetTextSubStringsForRange');
var _TTF_GetPreviousTextSubString = Module['_TTF_GetPreviousTextSubString'] = makeInvalidEarlyAccess('_TTF_GetPreviousTextSubString');
var _TTF_GetTextSubStringForPoint = Module['_TTF_GetTextSubStringForPoint'] = makeInvalidEarlyAccess('_TTF_GetTextSubStringForPoint');
var _SDL_abs = Module['_SDL_abs'] = makeInvalidEarlyAccess('_SDL_abs');
var _TTF_GetNextTextSubString = Module['_TTF_GetNextTextSubString'] = makeInvalidEarlyAccess('_TTF_GetNextTextSubString');
var _TTF_DestroyText = Module['_TTF_DestroyText'] = makeInvalidEarlyAccess('_TTF_DestroyText');
var _TTF_SetFontSize = Module['_TTF_SetFontSize'] = makeInvalidEarlyAccess('_TTF_SetFontSize');
var _TTF_GetFontSize = Module['_TTF_GetFontSize'] = makeInvalidEarlyAccess('_TTF_GetFontSize');
var _TTF_GetFontDPI = Module['_TTF_GetFontDPI'] = makeInvalidEarlyAccess('_TTF_GetFontDPI');
var _TTF_SetFontStyle = Module['_TTF_SetFontStyle'] = makeInvalidEarlyAccess('_TTF_SetFontStyle');
var _TTF_GetFontStyle = Module['_TTF_GetFontStyle'] = makeInvalidEarlyAccess('_TTF_GetFontStyle');
var _TTF_SetFontOutline = Module['_TTF_SetFontOutline'] = makeInvalidEarlyAccess('_TTF_SetFontOutline');
var _TTF_GetFontOutline = Module['_TTF_GetFontOutline'] = makeInvalidEarlyAccess('_TTF_GetFontOutline');
var _TTF_SetFontHinting = Module['_TTF_SetFontHinting'] = makeInvalidEarlyAccess('_TTF_SetFontHinting');
var _TTF_GetFontHinting = Module['_TTF_GetFontHinting'] = makeInvalidEarlyAccess('_TTF_GetFontHinting');
var _TTF_SetFontSDF = Module['_TTF_SetFontSDF'] = makeInvalidEarlyAccess('_TTF_SetFontSDF');
var _TTF_GetFontSDF = Module['_TTF_GetFontSDF'] = makeInvalidEarlyAccess('_TTF_GetFontSDF');
var _TTF_SetFontWrapAlignment = Module['_TTF_SetFontWrapAlignment'] = makeInvalidEarlyAccess('_TTF_SetFontWrapAlignment');
var _TTF_GetFontWrapAlignment = Module['_TTF_GetFontWrapAlignment'] = makeInvalidEarlyAccess('_TTF_GetFontWrapAlignment');
var _TTF_GetFontHeight = Module['_TTF_GetFontHeight'] = makeInvalidEarlyAccess('_TTF_GetFontHeight');
var _TTF_GetFontAscent = Module['_TTF_GetFontAscent'] = makeInvalidEarlyAccess('_TTF_GetFontAscent');
var _TTF_GetFontDescent = Module['_TTF_GetFontDescent'] = makeInvalidEarlyAccess('_TTF_GetFontDescent');
var _TTF_SetFontLineSkip = Module['_TTF_SetFontLineSkip'] = makeInvalidEarlyAccess('_TTF_SetFontLineSkip');
var _TTF_GetFontLineSkip = Module['_TTF_GetFontLineSkip'] = makeInvalidEarlyAccess('_TTF_GetFontLineSkip');
var _TTF_GetFontKerning = Module['_TTF_GetFontKerning'] = makeInvalidEarlyAccess('_TTF_GetFontKerning');
var _TTF_GetNumFontFaces = Module['_TTF_GetNumFontFaces'] = makeInvalidEarlyAccess('_TTF_GetNumFontFaces');
var _TTF_FontIsFixedWidth = Module['_TTF_FontIsFixedWidth'] = makeInvalidEarlyAccess('_TTF_FontIsFixedWidth');
var _TTF_FontIsScalable = Module['_TTF_FontIsScalable'] = makeInvalidEarlyAccess('_TTF_FontIsScalable');
var _TTF_GetFontFamilyName = Module['_TTF_GetFontFamilyName'] = makeInvalidEarlyAccess('_TTF_GetFontFamilyName');
var _TTF_GetFontStyleName = Module['_TTF_GetFontStyleName'] = makeInvalidEarlyAccess('_TTF_GetFontStyleName');
var _TTF_SetFontDirection = Module['_TTF_SetFontDirection'] = makeInvalidEarlyAccess('_TTF_SetFontDirection');
var _TTF_StringToTag = Module['_TTF_StringToTag'] = makeInvalidEarlyAccess('_TTF_StringToTag');
var _TTF_TagToString = Module['_TTF_TagToString'] = makeInvalidEarlyAccess('_TTF_TagToString');
var _TTF_SetFontScript = Module['_TTF_SetFontScript'] = makeInvalidEarlyAccess('_TTF_SetFontScript');
var _TTF_GetGlyphScript = Module['_TTF_GetGlyphScript'] = makeInvalidEarlyAccess('_TTF_GetGlyphScript');
var _TTF_SetFontLanguage = Module['_TTF_SetFontLanguage'] = makeInvalidEarlyAccess('_TTF_SetFontLanguage');
var _TTF_Quit = Module['_TTF_Quit'] = makeInvalidEarlyAccess('_TTF_Quit');
var _SDL_ShouldQuit = Module['_SDL_ShouldQuit'] = makeInvalidEarlyAccess('_SDL_ShouldQuit');
var _SDL_DestroyMutex = Module['_SDL_DestroyMutex'] = makeInvalidEarlyAccess('_SDL_DestroyMutex');
var _TTF_WasInit = Module['_TTF_WasInit'] = makeInvalidEarlyAccess('_TTF_WasInit');
var _SDL_GetAtomicInt = Module['_SDL_GetAtomicInt'] = makeInvalidEarlyAccess('_SDL_GetAtomicInt');
var _SDL_aligned_alloc = Module['_SDL_aligned_alloc'] = makeInvalidEarlyAccess('_SDL_aligned_alloc');
var _SDL_aligned_free = Module['_SDL_aligned_free'] = makeInvalidEarlyAccess('_SDL_aligned_free');
var _realloc = Module['_realloc'] = makeInvalidEarlyAccess('_realloc');
var _calloc = Module['_calloc'] = makeInvalidEarlyAccess('_calloc');
var _malloc = Module['_malloc'] = makeInvalidEarlyAccess('_malloc');
var _strerror = Module['_strerror'] = makeInvalidEarlyAccess('_strerror');
var _memcmp = Module['_memcmp'] = makeInvalidEarlyAccess('_memcmp');
var _setTempRet0 = Module['_setTempRet0'] = makeInvalidEarlyAccess('_setTempRet0');
var _getTempRet0 = Module['_getTempRet0'] = makeInvalidEarlyAccess('_getTempRet0');
var _plutosvg_version = Module['_plutosvg_version'] = makeInvalidEarlyAccess('_plutosvg_version');
var _plutosvg_version_string = Module['_plutosvg_version_string'] = makeInvalidEarlyAccess('_plutosvg_version_string');
var _plutosvg_document_destroy = Module['_plutosvg_document_destroy'] = makeInvalidEarlyAccess('_plutosvg_document_destroy');
var _plutovg_path_destroy = Module['_plutovg_path_destroy'] = makeInvalidEarlyAccess('_plutovg_path_destroy');
var _plutosvg_document_load_from_data = Module['_plutosvg_document_load_from_data'] = makeInvalidEarlyAccess('_plutosvg_document_load_from_data');
var _plutovg_path_create = Module['_plutovg_path_create'] = makeInvalidEarlyAccess('_plutovg_path_create');
var _plutosvg_document_load_from_file = Module['_plutosvg_document_load_from_file'] = makeInvalidEarlyAccess('_plutosvg_document_load_from_file');
var _plutosvg_document_render = Module['_plutosvg_document_render'] = makeInvalidEarlyAccess('_plutosvg_document_render');
var _plutovg_canvas_get_matrix = Module['_plutovg_canvas_get_matrix'] = makeInvalidEarlyAccess('_plutovg_canvas_get_matrix');
var _plutovg_matrix_translate = Module['_plutovg_matrix_translate'] = makeInvalidEarlyAccess('_plutovg_matrix_translate');
var _plutovg_path_reset = Module['_plutovg_path_reset'] = makeInvalidEarlyAccess('_plutovg_path_reset');
var _plutovg_path_move_to = Module['_plutovg_path_move_to'] = makeInvalidEarlyAccess('_plutovg_path_move_to');
var _plutovg_path_line_to = Module['_plutovg_path_line_to'] = makeInvalidEarlyAccess('_plutovg_path_line_to');
var _plutovg_path_add_ellipse = Module['_plutovg_path_add_ellipse'] = makeInvalidEarlyAccess('_plutovg_path_add_ellipse');
var _plutovg_path_add_circle = Module['_plutovg_path_add_circle'] = makeInvalidEarlyAccess('_plutovg_path_add_circle');
var _plutovg_path_add_round_rect = Module['_plutovg_path_add_round_rect'] = makeInvalidEarlyAccess('_plutovg_path_add_round_rect');
var _plutovg_path_close = Module['_plutovg_path_close'] = makeInvalidEarlyAccess('_plutovg_path_close');
var _plutovg_path_extents = Module['_plutovg_path_extents'] = makeInvalidEarlyAccess('_plutovg_path_extents');
var _plutovg_path_parse = Module['_plutovg_path_parse'] = makeInvalidEarlyAccess('_plutovg_path_parse');
var _plutovg_surface_load_from_image_base64 = Module['_plutovg_surface_load_from_image_base64'] = makeInvalidEarlyAccess('_plutovg_surface_load_from_image_base64');
var _plutovg_surface_get_width = Module['_plutovg_surface_get_width'] = makeInvalidEarlyAccess('_plutovg_surface_get_width');
var _plutovg_surface_get_height = Module['_plutovg_surface_get_height'] = makeInvalidEarlyAccess('_plutovg_surface_get_height');
var _plutovg_canvas_set_fill_rule = Module['_plutovg_canvas_set_fill_rule'] = makeInvalidEarlyAccess('_plutovg_canvas_set_fill_rule');
var _plutovg_canvas_set_opacity = Module['_plutovg_canvas_set_opacity'] = makeInvalidEarlyAccess('_plutovg_canvas_set_opacity');
var _plutovg_canvas_set_matrix = Module['_plutovg_canvas_set_matrix'] = makeInvalidEarlyAccess('_plutovg_canvas_set_matrix');
var _plutovg_canvas_translate = Module['_plutovg_canvas_translate'] = makeInvalidEarlyAccess('_plutovg_canvas_translate');
var _plutovg_canvas_set_texture = Module['_plutovg_canvas_set_texture'] = makeInvalidEarlyAccess('_plutovg_canvas_set_texture');
var _plutovg_canvas_fill_rect = Module['_plutovg_canvas_fill_rect'] = makeInvalidEarlyAccess('_plutovg_canvas_fill_rect');
var _plutovg_surface_destroy = Module['_plutovg_surface_destroy'] = makeInvalidEarlyAccess('_plutovg_surface_destroy');
var _plutosvg_document_render_to_surface = Module['_plutosvg_document_render_to_surface'] = makeInvalidEarlyAccess('_plutosvg_document_render_to_surface');
var _plutosvg_document_extents = Module['_plutosvg_document_extents'] = makeInvalidEarlyAccess('_plutosvg_document_extents');
var _plutovg_surface_create = Module['_plutovg_surface_create'] = makeInvalidEarlyAccess('_plutovg_surface_create');
var _plutovg_canvas_create = Module['_plutovg_canvas_create'] = makeInvalidEarlyAccess('_plutovg_canvas_create');
var _plutovg_canvas_scale = Module['_plutovg_canvas_scale'] = makeInvalidEarlyAccess('_plutovg_canvas_scale');
var _plutovg_canvas_destroy = Module['_plutovg_canvas_destroy'] = makeInvalidEarlyAccess('_plutovg_canvas_destroy');
var _plutovg_matrix_init_identity = Module['_plutovg_matrix_init_identity'] = makeInvalidEarlyAccess('_plutovg_matrix_init_identity');
var _plutosvg_document_get_width = Module['_plutosvg_document_get_width'] = makeInvalidEarlyAccess('_plutosvg_document_get_width');
var _plutosvg_document_get_height = Module['_plutosvg_document_get_height'] = makeInvalidEarlyAccess('_plutosvg_document_get_height');
var _plutovg_matrix_parse = Module['_plutovg_matrix_parse'] = makeInvalidEarlyAccess('_plutovg_matrix_parse');
var _plutovg_matrix_multiply = Module['_plutovg_matrix_multiply'] = makeInvalidEarlyAccess('_plutovg_matrix_multiply');
var _plutovg_matrix_scale = Module['_plutovg_matrix_scale'] = makeInvalidEarlyAccess('_plutovg_matrix_scale');
var _plutovg_matrix_invert = Module['_plutovg_matrix_invert'] = makeInvalidEarlyAccess('_plutovg_matrix_invert');
var _plutovg_matrix_map_rect = Module['_plutovg_matrix_map_rect'] = makeInvalidEarlyAccess('_plutovg_matrix_map_rect');
var _plutovg_canvas_fill_path = Module['_plutovg_canvas_fill_path'] = makeInvalidEarlyAccess('_plutovg_canvas_fill_path');
var _plutovg_canvas_set_dash_offset = Module['_plutovg_canvas_set_dash_offset'] = makeInvalidEarlyAccess('_plutovg_canvas_set_dash_offset');
var _plutovg_canvas_set_dash_array = Module['_plutovg_canvas_set_dash_array'] = makeInvalidEarlyAccess('_plutovg_canvas_set_dash_array');
var _plutovg_canvas_set_line_width = Module['_plutovg_canvas_set_line_width'] = makeInvalidEarlyAccess('_plutovg_canvas_set_line_width');
var _plutovg_canvas_set_line_cap = Module['_plutovg_canvas_set_line_cap'] = makeInvalidEarlyAccess('_plutovg_canvas_set_line_cap');
var _plutovg_canvas_set_line_join = Module['_plutovg_canvas_set_line_join'] = makeInvalidEarlyAccess('_plutovg_canvas_set_line_join');
var _plutovg_canvas_set_miter_limit = Module['_plutovg_canvas_set_miter_limit'] = makeInvalidEarlyAccess('_plutovg_canvas_set_miter_limit');
var _plutovg_canvas_stroke_path = Module['_plutovg_canvas_stroke_path'] = makeInvalidEarlyAccess('_plutovg_canvas_stroke_path');
var _plutovg_color_init_argb32 = Module['_plutovg_color_init_argb32'] = makeInvalidEarlyAccess('_plutovg_color_init_argb32');
var _plutovg_canvas_set_color = Module['_plutovg_canvas_set_color'] = makeInvalidEarlyAccess('_plutovg_canvas_set_color');
var _plutovg_canvas_set_linear_gradient = Module['_plutovg_canvas_set_linear_gradient'] = makeInvalidEarlyAccess('_plutovg_canvas_set_linear_gradient');
var _plutovg_canvas_set_radial_gradient = Module['_plutovg_canvas_set_radial_gradient'] = makeInvalidEarlyAccess('_plutovg_canvas_set_radial_gradient');
var _plutovg_color_parse = Module['_plutovg_color_parse'] = makeInvalidEarlyAccess('_plutovg_color_parse');
var _plutovg_color_to_argb32 = Module['_plutovg_color_to_argb32'] = makeInvalidEarlyAccess('_plutovg_color_to_argb32');
var _plutovg_matrix_init_translate = Module['_plutovg_matrix_init_translate'] = makeInvalidEarlyAccess('_plutovg_matrix_init_translate');
var _plutovg_surface_create_for_data = Module['_plutovg_surface_create_for_data'] = makeInvalidEarlyAccess('_plutovg_surface_create_for_data');
var _plutovg_canvas_transform = Module['_plutovg_canvas_transform'] = makeInvalidEarlyAccess('_plutovg_canvas_transform');
var _plutovg_matrix_init_scale = Module['_plutovg_matrix_init_scale'] = makeInvalidEarlyAccess('_plutovg_matrix_init_scale');
var _plutovg_version = Module['_plutovg_version'] = makeInvalidEarlyAccess('_plutovg_version');
var _plutovg_version_string = Module['_plutovg_version_string'] = makeInvalidEarlyAccess('_plutovg_version_string');
var _plutovg_surface_reference = Module['_plutovg_surface_reference'] = makeInvalidEarlyAccess('_plutovg_surface_reference');
var _plutovg_canvas_reference = Module['_plutovg_canvas_reference'] = makeInvalidEarlyAccess('_plutovg_canvas_reference');
var _plutovg_paint_destroy = Module['_plutovg_paint_destroy'] = makeInvalidEarlyAccess('_plutovg_paint_destroy');
var _plutovg_font_face_destroy = Module['_plutovg_font_face_destroy'] = makeInvalidEarlyAccess('_plutovg_font_face_destroy');
var _plutovg_canvas_get_reference_count = Module['_plutovg_canvas_get_reference_count'] = makeInvalidEarlyAccess('_plutovg_canvas_get_reference_count');
var _plutovg_canvas_get_surface = Module['_plutovg_canvas_get_surface'] = makeInvalidEarlyAccess('_plutovg_canvas_get_surface');
var _plutovg_canvas_save = Module['_plutovg_canvas_save'] = makeInvalidEarlyAccess('_plutovg_canvas_save');
var _plutovg_paint_reference = Module['_plutovg_paint_reference'] = makeInvalidEarlyAccess('_plutovg_paint_reference');
var _plutovg_font_face_reference = Module['_plutovg_font_face_reference'] = makeInvalidEarlyAccess('_plutovg_font_face_reference');
var _plutovg_canvas_restore = Module['_plutovg_canvas_restore'] = makeInvalidEarlyAccess('_plutovg_canvas_restore');
var _plutovg_canvas_set_rgb = Module['_plutovg_canvas_set_rgb'] = makeInvalidEarlyAccess('_plutovg_canvas_set_rgb');
var _plutovg_color_init_rgba = Module['_plutovg_color_init_rgba'] = makeInvalidEarlyAccess('_plutovg_color_init_rgba');
var _plutovg_canvas_set_rgba = Module['_plutovg_canvas_set_rgba'] = makeInvalidEarlyAccess('_plutovg_canvas_set_rgba');
var _plutovg_canvas_set_paint = Module['_plutovg_canvas_set_paint'] = makeInvalidEarlyAccess('_plutovg_canvas_set_paint');
var _plutovg_paint_create_linear_gradient = Module['_plutovg_paint_create_linear_gradient'] = makeInvalidEarlyAccess('_plutovg_paint_create_linear_gradient');
var _plutovg_paint_create_radial_gradient = Module['_plutovg_paint_create_radial_gradient'] = makeInvalidEarlyAccess('_plutovg_paint_create_radial_gradient');
var _plutovg_paint_create_texture = Module['_plutovg_paint_create_texture'] = makeInvalidEarlyAccess('_plutovg_paint_create_texture');
var _plutovg_canvas_get_paint = Module['_plutovg_canvas_get_paint'] = makeInvalidEarlyAccess('_plutovg_canvas_get_paint');
var _plutovg_canvas_set_font = Module['_plutovg_canvas_set_font'] = makeInvalidEarlyAccess('_plutovg_canvas_set_font');
var _plutovg_canvas_set_font_face = Module['_plutovg_canvas_set_font_face'] = makeInvalidEarlyAccess('_plutovg_canvas_set_font_face');
var _plutovg_canvas_set_font_size = Module['_plutovg_canvas_set_font_size'] = makeInvalidEarlyAccess('_plutovg_canvas_set_font_size');
var _plutovg_canvas_get_font_face = Module['_plutovg_canvas_get_font_face'] = makeInvalidEarlyAccess('_plutovg_canvas_get_font_face');
var _plutovg_canvas_get_font_size = Module['_plutovg_canvas_get_font_size'] = makeInvalidEarlyAccess('_plutovg_canvas_get_font_size');
var _plutovg_canvas_get_fill_rule = Module['_plutovg_canvas_get_fill_rule'] = makeInvalidEarlyAccess('_plutovg_canvas_get_fill_rule');
var _plutovg_canvas_set_operator = Module['_plutovg_canvas_set_operator'] = makeInvalidEarlyAccess('_plutovg_canvas_set_operator');
var _plutovg_canvas_get_operator = Module['_plutovg_canvas_get_operator'] = makeInvalidEarlyAccess('_plutovg_canvas_get_operator');
var _plutovg_canvas_get_opacity = Module['_plutovg_canvas_get_opacity'] = makeInvalidEarlyAccess('_plutovg_canvas_get_opacity');
var _plutovg_canvas_get_line_width = Module['_plutovg_canvas_get_line_width'] = makeInvalidEarlyAccess('_plutovg_canvas_get_line_width');
var _plutovg_canvas_get_line_cap = Module['_plutovg_canvas_get_line_cap'] = makeInvalidEarlyAccess('_plutovg_canvas_get_line_cap');
var _plutovg_canvas_get_line_join = Module['_plutovg_canvas_get_line_join'] = makeInvalidEarlyAccess('_plutovg_canvas_get_line_join');
var _plutovg_canvas_get_miter_limit = Module['_plutovg_canvas_get_miter_limit'] = makeInvalidEarlyAccess('_plutovg_canvas_get_miter_limit');
var _plutovg_canvas_set_dash = Module['_plutovg_canvas_set_dash'] = makeInvalidEarlyAccess('_plutovg_canvas_set_dash');
var _plutovg_canvas_get_dash_offset = Module['_plutovg_canvas_get_dash_offset'] = makeInvalidEarlyAccess('_plutovg_canvas_get_dash_offset');
var _plutovg_canvas_get_dash_array = Module['_plutovg_canvas_get_dash_array'] = makeInvalidEarlyAccess('_plutovg_canvas_get_dash_array');
var _plutovg_canvas_shear = Module['_plutovg_canvas_shear'] = makeInvalidEarlyAccess('_plutovg_canvas_shear');
var _plutovg_matrix_shear = Module['_plutovg_matrix_shear'] = makeInvalidEarlyAccess('_plutovg_matrix_shear');
var _plutovg_canvas_rotate = Module['_plutovg_canvas_rotate'] = makeInvalidEarlyAccess('_plutovg_canvas_rotate');
var _plutovg_matrix_rotate = Module['_plutovg_matrix_rotate'] = makeInvalidEarlyAccess('_plutovg_matrix_rotate');
var _plutovg_canvas_reset_matrix = Module['_plutovg_canvas_reset_matrix'] = makeInvalidEarlyAccess('_plutovg_canvas_reset_matrix');
var _plutovg_canvas_map = Module['_plutovg_canvas_map'] = makeInvalidEarlyAccess('_plutovg_canvas_map');
var _plutovg_matrix_map = Module['_plutovg_matrix_map'] = makeInvalidEarlyAccess('_plutovg_matrix_map');
var _plutovg_canvas_map_point = Module['_plutovg_canvas_map_point'] = makeInvalidEarlyAccess('_plutovg_canvas_map_point');
var _plutovg_matrix_map_point = Module['_plutovg_matrix_map_point'] = makeInvalidEarlyAccess('_plutovg_matrix_map_point');
var _plutovg_canvas_map_rect = Module['_plutovg_canvas_map_rect'] = makeInvalidEarlyAccess('_plutovg_canvas_map_rect');
var _plutovg_canvas_move_to = Module['_plutovg_canvas_move_to'] = makeInvalidEarlyAccess('_plutovg_canvas_move_to');
var _plutovg_canvas_line_to = Module['_plutovg_canvas_line_to'] = makeInvalidEarlyAccess('_plutovg_canvas_line_to');
var _plutovg_canvas_quad_to = Module['_plutovg_canvas_quad_to'] = makeInvalidEarlyAccess('_plutovg_canvas_quad_to');
var _plutovg_path_quad_to = Module['_plutovg_path_quad_to'] = makeInvalidEarlyAccess('_plutovg_path_quad_to');
var _plutovg_canvas_cubic_to = Module['_plutovg_canvas_cubic_to'] = makeInvalidEarlyAccess('_plutovg_canvas_cubic_to');
var _plutovg_path_cubic_to = Module['_plutovg_path_cubic_to'] = makeInvalidEarlyAccess('_plutovg_path_cubic_to');
var _plutovg_canvas_arc_to = Module['_plutovg_canvas_arc_to'] = makeInvalidEarlyAccess('_plutovg_canvas_arc_to');
var _plutovg_path_arc_to = Module['_plutovg_path_arc_to'] = makeInvalidEarlyAccess('_plutovg_path_arc_to');
var _plutovg_canvas_rect = Module['_plutovg_canvas_rect'] = makeInvalidEarlyAccess('_plutovg_canvas_rect');
var _plutovg_path_add_rect = Module['_plutovg_path_add_rect'] = makeInvalidEarlyAccess('_plutovg_path_add_rect');
var _plutovg_canvas_round_rect = Module['_plutovg_canvas_round_rect'] = makeInvalidEarlyAccess('_plutovg_canvas_round_rect');
var _plutovg_canvas_ellipse = Module['_plutovg_canvas_ellipse'] = makeInvalidEarlyAccess('_plutovg_canvas_ellipse');
var _plutovg_canvas_circle = Module['_plutovg_canvas_circle'] = makeInvalidEarlyAccess('_plutovg_canvas_circle');
var _plutovg_canvas_arc = Module['_plutovg_canvas_arc'] = makeInvalidEarlyAccess('_plutovg_canvas_arc');
var _plutovg_path_add_arc = Module['_plutovg_path_add_arc'] = makeInvalidEarlyAccess('_plutovg_path_add_arc');
var _plutovg_canvas_add_path = Module['_plutovg_canvas_add_path'] = makeInvalidEarlyAccess('_plutovg_canvas_add_path');
var _plutovg_path_add_path = Module['_plutovg_path_add_path'] = makeInvalidEarlyAccess('_plutovg_path_add_path');
var _plutovg_canvas_new_path = Module['_plutovg_canvas_new_path'] = makeInvalidEarlyAccess('_plutovg_canvas_new_path');
var _plutovg_canvas_close_path = Module['_plutovg_canvas_close_path'] = makeInvalidEarlyAccess('_plutovg_canvas_close_path');
var _plutovg_canvas_get_current_point = Module['_plutovg_canvas_get_current_point'] = makeInvalidEarlyAccess('_plutovg_canvas_get_current_point');
var _plutovg_path_get_current_point = Module['_plutovg_path_get_current_point'] = makeInvalidEarlyAccess('_plutovg_path_get_current_point');
var _plutovg_canvas_get_path = Module['_plutovg_canvas_get_path'] = makeInvalidEarlyAccess('_plutovg_canvas_get_path');
var _plutovg_canvas_fill_extents = Module['_plutovg_canvas_fill_extents'] = makeInvalidEarlyAccess('_plutovg_canvas_fill_extents');
var _plutovg_canvas_stroke_extents = Module['_plutovg_canvas_stroke_extents'] = makeInvalidEarlyAccess('_plutovg_canvas_stroke_extents');
var _plutovg_canvas_clip_extents = Module['_plutovg_canvas_clip_extents'] = makeInvalidEarlyAccess('_plutovg_canvas_clip_extents');
var _plutovg_canvas_fill = Module['_plutovg_canvas_fill'] = makeInvalidEarlyAccess('_plutovg_canvas_fill');
var _plutovg_canvas_fill_preserve = Module['_plutovg_canvas_fill_preserve'] = makeInvalidEarlyAccess('_plutovg_canvas_fill_preserve');
var _plutovg_canvas_stroke = Module['_plutovg_canvas_stroke'] = makeInvalidEarlyAccess('_plutovg_canvas_stroke');
var _plutovg_canvas_stroke_preserve = Module['_plutovg_canvas_stroke_preserve'] = makeInvalidEarlyAccess('_plutovg_canvas_stroke_preserve');
var _plutovg_canvas_clip = Module['_plutovg_canvas_clip'] = makeInvalidEarlyAccess('_plutovg_canvas_clip');
var _plutovg_canvas_clip_preserve = Module['_plutovg_canvas_clip_preserve'] = makeInvalidEarlyAccess('_plutovg_canvas_clip_preserve');
var _plutovg_canvas_paint = Module['_plutovg_canvas_paint'] = makeInvalidEarlyAccess('_plutovg_canvas_paint');
var _plutovg_canvas_stroke_rect = Module['_plutovg_canvas_stroke_rect'] = makeInvalidEarlyAccess('_plutovg_canvas_stroke_rect');
var _plutovg_canvas_clip_rect = Module['_plutovg_canvas_clip_rect'] = makeInvalidEarlyAccess('_plutovg_canvas_clip_rect');
var _plutovg_canvas_clip_path = Module['_plutovg_canvas_clip_path'] = makeInvalidEarlyAccess('_plutovg_canvas_clip_path');
var _plutovg_canvas_add_glyph = Module['_plutovg_canvas_add_glyph'] = makeInvalidEarlyAccess('_plutovg_canvas_add_glyph');
var _plutovg_font_face_get_glyph_path = Module['_plutovg_font_face_get_glyph_path'] = makeInvalidEarlyAccess('_plutovg_font_face_get_glyph_path');
var _plutovg_canvas_add_text = Module['_plutovg_canvas_add_text'] = makeInvalidEarlyAccess('_plutovg_canvas_add_text');
var _plutovg_text_iterator_init = Module['_plutovg_text_iterator_init'] = makeInvalidEarlyAccess('_plutovg_text_iterator_init');
var _plutovg_text_iterator_has_next = Module['_plutovg_text_iterator_has_next'] = makeInvalidEarlyAccess('_plutovg_text_iterator_has_next');
var _plutovg_text_iterator_next = Module['_plutovg_text_iterator_next'] = makeInvalidEarlyAccess('_plutovg_text_iterator_next');
var _plutovg_canvas_fill_text = Module['_plutovg_canvas_fill_text'] = makeInvalidEarlyAccess('_plutovg_canvas_fill_text');
var _plutovg_canvas_stroke_text = Module['_plutovg_canvas_stroke_text'] = makeInvalidEarlyAccess('_plutovg_canvas_stroke_text');
var _plutovg_canvas_clip_text = Module['_plutovg_canvas_clip_text'] = makeInvalidEarlyAccess('_plutovg_canvas_clip_text');
var _plutovg_canvas_font_metrics = Module['_plutovg_canvas_font_metrics'] = makeInvalidEarlyAccess('_plutovg_canvas_font_metrics');
var _plutovg_font_face_get_metrics = Module['_plutovg_font_face_get_metrics'] = makeInvalidEarlyAccess('_plutovg_font_face_get_metrics');
var _plutovg_canvas_glyph_metrics = Module['_plutovg_canvas_glyph_metrics'] = makeInvalidEarlyAccess('_plutovg_canvas_glyph_metrics');
var _plutovg_font_face_get_glyph_metrics = Module['_plutovg_font_face_get_glyph_metrics'] = makeInvalidEarlyAccess('_plutovg_font_face_get_glyph_metrics');
var _plutovg_canvas_text_extents = Module['_plutovg_canvas_text_extents'] = makeInvalidEarlyAccess('_plutovg_canvas_text_extents');
var _plutovg_font_face_text_extents = Module['_plutovg_font_face_text_extents'] = makeInvalidEarlyAccess('_plutovg_font_face_text_extents');
var _plutovg_font_face_load_from_file = Module['_plutovg_font_face_load_from_file'] = makeInvalidEarlyAccess('_plutovg_font_face_load_from_file');
var _plutovg_font_face_load_from_data = Module['_plutovg_font_face_load_from_data'] = makeInvalidEarlyAccess('_plutovg_font_face_load_from_data');
var _plutovg_font_face_get_reference_count = Module['_plutovg_font_face_get_reference_count'] = makeInvalidEarlyAccess('_plutovg_font_face_get_reference_count');
var _plutovg_font_face_traverse_glyph_path = Module['_plutovg_font_face_traverse_glyph_path'] = makeInvalidEarlyAccess('_plutovg_font_face_traverse_glyph_path');
var _plutovg_matrix_map_points = Module['_plutovg_matrix_map_points'] = makeInvalidEarlyAccess('_plutovg_matrix_map_points');
var _plutovg_matrix_init = Module['_plutovg_matrix_init'] = makeInvalidEarlyAccess('_plutovg_matrix_init');
var _plutovg_matrix_init_rotate = Module['_plutovg_matrix_init_rotate'] = makeInvalidEarlyAccess('_plutovg_matrix_init_rotate');
var _plutovg_matrix_init_shear = Module['_plutovg_matrix_init_shear'] = makeInvalidEarlyAccess('_plutovg_matrix_init_shear');
var _plutovg_color_init_rgb = Module['_plutovg_color_init_rgb'] = makeInvalidEarlyAccess('_plutovg_color_init_rgb');
var _plutovg_color_init_rgb8 = Module['_plutovg_color_init_rgb8'] = makeInvalidEarlyAccess('_plutovg_color_init_rgb8');
var _plutovg_color_init_rgba8 = Module['_plutovg_color_init_rgba8'] = makeInvalidEarlyAccess('_plutovg_color_init_rgba8');
var _plutovg_color_init_rgba32 = Module['_plutovg_color_init_rgba32'] = makeInvalidEarlyAccess('_plutovg_color_init_rgba32');
var _plutovg_color_to_rgba32 = Module['_plutovg_color_to_rgba32'] = makeInvalidEarlyAccess('_plutovg_color_to_rgba32');
var _plutovg_paint_create_rgb = Module['_plutovg_paint_create_rgb'] = makeInvalidEarlyAccess('_plutovg_paint_create_rgb');
var _plutovg_paint_create_rgba = Module['_plutovg_paint_create_rgba'] = makeInvalidEarlyAccess('_plutovg_paint_create_rgba');
var _plutovg_paint_create_color = Module['_plutovg_paint_create_color'] = makeInvalidEarlyAccess('_plutovg_paint_create_color');
var _plutovg_paint_get_reference_count = Module['_plutovg_paint_get_reference_count'] = makeInvalidEarlyAccess('_plutovg_paint_get_reference_count');
var _plutovg_path_iterator_init = Module['_plutovg_path_iterator_init'] = makeInvalidEarlyAccess('_plutovg_path_iterator_init');
var _plutovg_path_iterator_has_next = Module['_plutovg_path_iterator_has_next'] = makeInvalidEarlyAccess('_plutovg_path_iterator_has_next');
var _plutovg_path_iterator_next = Module['_plutovg_path_iterator_next'] = makeInvalidEarlyAccess('_plutovg_path_iterator_next');
var _plutovg_path_reference = Module['_plutovg_path_reference'] = makeInvalidEarlyAccess('_plutovg_path_reference');
var _plutovg_path_get_reference_count = Module['_plutovg_path_get_reference_count'] = makeInvalidEarlyAccess('_plutovg_path_get_reference_count');
var _plutovg_path_get_elements = Module['_plutovg_path_get_elements'] = makeInvalidEarlyAccess('_plutovg_path_get_elements');
var _plutovg_path_reserve = Module['_plutovg_path_reserve'] = makeInvalidEarlyAccess('_plutovg_path_reserve');
var _plutovg_path_transform = Module['_plutovg_path_transform'] = makeInvalidEarlyAccess('_plutovg_path_transform');
var _plutovg_path_traverse = Module['_plutovg_path_traverse'] = makeInvalidEarlyAccess('_plutovg_path_traverse');
var _plutovg_path_traverse_flatten = Module['_plutovg_path_traverse_flatten'] = makeInvalidEarlyAccess('_plutovg_path_traverse_flatten');
var _plutovg_path_traverse_dashed = Module['_plutovg_path_traverse_dashed'] = makeInvalidEarlyAccess('_plutovg_path_traverse_dashed');
var _plutovg_path_clone = Module['_plutovg_path_clone'] = makeInvalidEarlyAccess('_plutovg_path_clone');
var _plutovg_path_clone_flatten = Module['_plutovg_path_clone_flatten'] = makeInvalidEarlyAccess('_plutovg_path_clone_flatten');
var _plutovg_path_clone_dashed = Module['_plutovg_path_clone_dashed'] = makeInvalidEarlyAccess('_plutovg_path_clone_dashed');
var _plutovg_path_length = Module['_plutovg_path_length'] = makeInvalidEarlyAccess('_plutovg_path_length');
var _plutovg_surface_load_from_image_file = Module['_plutovg_surface_load_from_image_file'] = makeInvalidEarlyAccess('_plutovg_surface_load_from_image_file');
var _plutovg_surface_load_from_image_data = Module['_plutovg_surface_load_from_image_data'] = makeInvalidEarlyAccess('_plutovg_surface_load_from_image_data');
var _plutovg_surface_get_reference_count = Module['_plutovg_surface_get_reference_count'] = makeInvalidEarlyAccess('_plutovg_surface_get_reference_count');
var _plutovg_surface_get_data = Module['_plutovg_surface_get_data'] = makeInvalidEarlyAccess('_plutovg_surface_get_data');
var _plutovg_surface_get_stride = Module['_plutovg_surface_get_stride'] = makeInvalidEarlyAccess('_plutovg_surface_get_stride');
var _plutovg_surface_clear = Module['_plutovg_surface_clear'] = makeInvalidEarlyAccess('_plutovg_surface_clear');
var _plutovg_surface_write_to_png = Module['_plutovg_surface_write_to_png'] = makeInvalidEarlyAccess('_plutovg_surface_write_to_png');
var _plutovg_surface_write_to_jpg = Module['_plutovg_surface_write_to_jpg'] = makeInvalidEarlyAccess('_plutovg_surface_write_to_jpg');
var _plutovg_surface_write_to_png_stream = Module['_plutovg_surface_write_to_png_stream'] = makeInvalidEarlyAccess('_plutovg_surface_write_to_png_stream');
var _plutovg_surface_write_to_jpg_stream = Module['_plutovg_surface_write_to_jpg_stream'] = makeInvalidEarlyAccess('_plutovg_surface_write_to_jpg_stream');
var _plutovg_convert_argb_to_rgba = Module['_plutovg_convert_argb_to_rgba'] = makeInvalidEarlyAccess('_plutovg_convert_argb_to_rgba');
var _plutovg_convert_rgba_to_argb = Module['_plutovg_convert_rgba_to_argb'] = makeInvalidEarlyAccess('_plutovg_convert_rgba_to_argb');
var _SDL_ReadU16BE = Module['_SDL_ReadU16BE'] = makeInvalidEarlyAccess('_SDL_ReadU16BE');
var _SDL_iconv_string = Module['_SDL_iconv_string'] = makeInvalidEarlyAccess('_SDL_iconv_string');
var _SDL_IOFromConstMem = Module['_SDL_IOFromConstMem'] = makeInvalidEarlyAccess('_SDL_IOFromConstMem');
var _SDL_CreateAudioStream = Module['_SDL_CreateAudioStream'] = makeInvalidEarlyAccess('_SDL_CreateAudioStream');
var _SDL_ClearAudioStream = Module['_SDL_ClearAudioStream'] = makeInvalidEarlyAccess('_SDL_ClearAudioStream');
var _SDL_DestroyAudioStream = Module['_SDL_DestroyAudioStream'] = makeInvalidEarlyAccess('_SDL_DestroyAudioStream');
var _SDL_GetAudioStreamData = Module['_SDL_GetAudioStreamData'] = makeInvalidEarlyAccess('_SDL_GetAudioStreamData');
var _SDL_PutAudioStreamData = Module['_SDL_PutAudioStreamData'] = makeInvalidEarlyAccess('_SDL_PutAudioStreamData');
var _SDL_FlushAudioStream = Module['_SDL_FlushAudioStream'] = makeInvalidEarlyAccess('_SDL_FlushAudioStream');
var _SDL_scalbn = Module['_SDL_scalbn'] = makeInvalidEarlyAccess('_SDL_scalbn');
var _SDL_log = Module['_SDL_log'] = makeInvalidEarlyAccess('_SDL_log');
var _SDL_exp = Module['_SDL_exp'] = makeInvalidEarlyAccess('_SDL_exp');
var _SDL_floor = Module['_SDL_floor'] = makeInvalidEarlyAccess('_SDL_floor');
var _SDL_cos = Module['_SDL_cos'] = makeInvalidEarlyAccess('_SDL_cos');
var _SDL_sin = Module['_SDL_sin'] = makeInvalidEarlyAccess('_SDL_sin');
var _SDL_memcpy = Module['_SDL_memcpy'] = makeInvalidEarlyAccess('_SDL_memcpy');
var _SDL_getenv = Module['_SDL_getenv'] = makeInvalidEarlyAccess('_SDL_getenv');
var _Mix_GetTimidityCfg = Module['_Mix_GetTimidityCfg'] = makeInvalidEarlyAccess('_Mix_GetTimidityCfg');
var _Mix_SetPanning = Module['_Mix_SetPanning'] = makeInvalidEarlyAccess('_Mix_SetPanning');
var _Mix_QuerySpec = Module['_Mix_QuerySpec'] = makeInvalidEarlyAccess('_Mix_QuerySpec');
var _Mix_SetPosition = Module['_Mix_SetPosition'] = makeInvalidEarlyAccess('_Mix_SetPosition');
var _Mix_SetDistance = Module['_Mix_SetDistance'] = makeInvalidEarlyAccess('_Mix_SetDistance');
var _Mix_SetReverseStereo = Module['_Mix_SetReverseStereo'] = makeInvalidEarlyAccess('_Mix_SetReverseStereo');
var _Mix_UnregisterEffect = Module['_Mix_UnregisterEffect'] = makeInvalidEarlyAccess('_Mix_UnregisterEffect');
var _Mix_RegisterEffect = Module['_Mix_RegisterEffect'] = makeInvalidEarlyAccess('_Mix_RegisterEffect');
var _Mix_GetNumChunkDecoders = Module['_Mix_GetNumChunkDecoders'] = makeInvalidEarlyAccess('_Mix_GetNumChunkDecoders');
var _Mix_GetChunkDecoder = Module['_Mix_GetChunkDecoder'] = makeInvalidEarlyAccess('_Mix_GetChunkDecoder');
var _Mix_HasChunkDecoder = Module['_Mix_HasChunkDecoder'] = makeInvalidEarlyAccess('_Mix_HasChunkDecoder');
var _Mix_Version = Module['_Mix_Version'] = makeInvalidEarlyAccess('_Mix_Version');
var _Mix_Init = Module['_Mix_Init'] = makeInvalidEarlyAccess('_Mix_Init');
var _Mix_Quit = Module['_Mix_Quit'] = makeInvalidEarlyAccess('_Mix_Quit');
var _Mix_OpenAudio = Module['_Mix_OpenAudio'] = makeInvalidEarlyAccess('_Mix_OpenAudio');
var _SDL_WasInit = Module['_SDL_WasInit'] = makeInvalidEarlyAccess('_SDL_WasInit');
var _SDL_InitSubSystem = Module['_SDL_InitSubSystem'] = makeInvalidEarlyAccess('_SDL_InitSubSystem');
var _Mix_CloseAudio = Module['_Mix_CloseAudio'] = makeInvalidEarlyAccess('_Mix_CloseAudio');
var _SDL_OpenAudioDevice = Module['_SDL_OpenAudioDevice'] = makeInvalidEarlyAccess('_SDL_OpenAudioDevice');
var _SDL_GetAudioDeviceFormat = Module['_SDL_GetAudioDeviceFormat'] = makeInvalidEarlyAccess('_SDL_GetAudioDeviceFormat');
var _SDL_CloseAudioDevice = Module['_SDL_CloseAudioDevice'] = makeInvalidEarlyAccess('_SDL_CloseAudioDevice');
var _Mix_VolumeMusic = Module['_Mix_VolumeMusic'] = makeInvalidEarlyAccess('_Mix_VolumeMusic');
var _SDL_BindAudioStream = Module['_SDL_BindAudioStream'] = makeInvalidEarlyAccess('_SDL_BindAudioStream');
var _SDL_SetAudioStreamGetCallback = Module['_SDL_SetAudioStreamGetCallback'] = makeInvalidEarlyAccess('_SDL_SetAudioStreamGetCallback');
var _SDL_LockAudioStream = Module['_SDL_LockAudioStream'] = makeInvalidEarlyAccess('_SDL_LockAudioStream');
var _SDL_UnlockAudioStream = Module['_SDL_UnlockAudioStream'] = makeInvalidEarlyAccess('_SDL_UnlockAudioStream');
var _SDL_GetSIMDAlignment = Module['_SDL_GetSIMDAlignment'] = makeInvalidEarlyAccess('_SDL_GetSIMDAlignment');
var _SDL_GetSilenceValueForFormat = Module['_SDL_GetSilenceValueForFormat'] = makeInvalidEarlyAccess('_SDL_GetSilenceValueForFormat');
var _SDL_GetTicks = Module['_SDL_GetTicks'] = makeInvalidEarlyAccess('_SDL_GetTicks');
var _Mix_Volume = Module['_Mix_Volume'] = makeInvalidEarlyAccess('_Mix_Volume');
var _SDL_MixAudio = Module['_SDL_MixAudio'] = makeInvalidEarlyAccess('_SDL_MixAudio');
var _Mix_PauseAudio = Module['_Mix_PauseAudio'] = makeInvalidEarlyAccess('_Mix_PauseAudio');
var _SDL_PauseAudioDevice = Module['_SDL_PauseAudioDevice'] = makeInvalidEarlyAccess('_SDL_PauseAudioDevice');
var _SDL_ResumeAudioDevice = Module['_SDL_ResumeAudioDevice'] = makeInvalidEarlyAccess('_SDL_ResumeAudioDevice');
var _Mix_AllocateChannels = Module['_Mix_AllocateChannels'] = makeInvalidEarlyAccess('_Mix_AllocateChannels');
var _Mix_UnregisterAllEffects = Module['_Mix_UnregisterAllEffects'] = makeInvalidEarlyAccess('_Mix_UnregisterAllEffects');
var _Mix_HaltChannel = Module['_Mix_HaltChannel'] = makeInvalidEarlyAccess('_Mix_HaltChannel');
var _Mix_LoadWAV_IO = Module['_Mix_LoadWAV_IO'] = makeInvalidEarlyAccess('_Mix_LoadWAV_IO');
var _SDL_LoadWAV_IO = Module['_SDL_LoadWAV_IO'] = makeInvalidEarlyAccess('_SDL_LoadWAV_IO');
var _SDL_ConvertAudioSamples = Module['_SDL_ConvertAudioSamples'] = makeInvalidEarlyAccess('_SDL_ConvertAudioSamples');
var _Mix_LoadWAV = Module['_Mix_LoadWAV'] = makeInvalidEarlyAccess('_Mix_LoadWAV');
var _Mix_QuickLoad_WAV = Module['_Mix_QuickLoad_WAV'] = makeInvalidEarlyAccess('_Mix_QuickLoad_WAV');
var _Mix_QuickLoad_RAW = Module['_Mix_QuickLoad_RAW'] = makeInvalidEarlyAccess('_Mix_QuickLoad_RAW');
var _Mix_FreeChunk = Module['_Mix_FreeChunk'] = makeInvalidEarlyAccess('_Mix_FreeChunk');
var _Mix_SetPostMix = Module['_Mix_SetPostMix'] = makeInvalidEarlyAccess('_Mix_SetPostMix');
var _Mix_HookMusic = Module['_Mix_HookMusic'] = makeInvalidEarlyAccess('_Mix_HookMusic');
var _Mix_GetMusicHookData = Module['_Mix_GetMusicHookData'] = makeInvalidEarlyAccess('_Mix_GetMusicHookData');
var _Mix_ChannelFinished = Module['_Mix_ChannelFinished'] = makeInvalidEarlyAccess('_Mix_ChannelFinished');
var _Mix_ReserveChannels = Module['_Mix_ReserveChannels'] = makeInvalidEarlyAccess('_Mix_ReserveChannels');
var _Mix_PlayChannelTimed = Module['_Mix_PlayChannelTimed'] = makeInvalidEarlyAccess('_Mix_PlayChannelTimed');
var _Mix_Playing = Module['_Mix_Playing'] = makeInvalidEarlyAccess('_Mix_Playing');
var _Mix_PlayChannel = Module['_Mix_PlayChannel'] = makeInvalidEarlyAccess('_Mix_PlayChannel');
var _Mix_ExpireChannel = Module['_Mix_ExpireChannel'] = makeInvalidEarlyAccess('_Mix_ExpireChannel');
var _Mix_FadeInChannelTimed = Module['_Mix_FadeInChannelTimed'] = makeInvalidEarlyAccess('_Mix_FadeInChannelTimed');
var _Mix_FadeInChannel = Module['_Mix_FadeInChannel'] = makeInvalidEarlyAccess('_Mix_FadeInChannel');
var _Mix_VolumeChunk = Module['_Mix_VolumeChunk'] = makeInvalidEarlyAccess('_Mix_VolumeChunk');
var _Mix_HaltGroup = Module['_Mix_HaltGroup'] = makeInvalidEarlyAccess('_Mix_HaltGroup');
var _Mix_FadeOutChannel = Module['_Mix_FadeOutChannel'] = makeInvalidEarlyAccess('_Mix_FadeOutChannel');
var _Mix_FadeOutGroup = Module['_Mix_FadeOutGroup'] = makeInvalidEarlyAccess('_Mix_FadeOutGroup');
var _Mix_FadingChannel = Module['_Mix_FadingChannel'] = makeInvalidEarlyAccess('_Mix_FadingChannel');
var _Mix_GetChunk = Module['_Mix_GetChunk'] = makeInvalidEarlyAccess('_Mix_GetChunk');
var _Mix_Pause = Module['_Mix_Pause'] = makeInvalidEarlyAccess('_Mix_Pause');
var _Mix_PauseGroup = Module['_Mix_PauseGroup'] = makeInvalidEarlyAccess('_Mix_PauseGroup');
var _Mix_Resume = Module['_Mix_Resume'] = makeInvalidEarlyAccess('_Mix_Resume');
var _Mix_ResumeGroup = Module['_Mix_ResumeGroup'] = makeInvalidEarlyAccess('_Mix_ResumeGroup');
var _Mix_Paused = Module['_Mix_Paused'] = makeInvalidEarlyAccess('_Mix_Paused');
var _Mix_GroupChannel = Module['_Mix_GroupChannel'] = makeInvalidEarlyAccess('_Mix_GroupChannel');
var _Mix_GroupChannels = Module['_Mix_GroupChannels'] = makeInvalidEarlyAccess('_Mix_GroupChannels');
var _Mix_GroupAvailable = Module['_Mix_GroupAvailable'] = makeInvalidEarlyAccess('_Mix_GroupAvailable');
var _Mix_GroupCount = Module['_Mix_GroupCount'] = makeInvalidEarlyAccess('_Mix_GroupCount');
var _Mix_GroupOldest = Module['_Mix_GroupOldest'] = makeInvalidEarlyAccess('_Mix_GroupOldest');
var _Mix_GroupNewer = Module['_Mix_GroupNewer'] = makeInvalidEarlyAccess('_Mix_GroupNewer');
var _Mix_MasterVolume = Module['_Mix_MasterVolume'] = makeInvalidEarlyAccess('_Mix_MasterVolume');
var _SDL_SetAtomicInt = Module['_SDL_SetAtomicInt'] = makeInvalidEarlyAccess('_SDL_SetAtomicInt');
var _Mix_GetNumMusicDecoders = Module['_Mix_GetNumMusicDecoders'] = makeInvalidEarlyAccess('_Mix_GetNumMusicDecoders');
var _Mix_GetMusicDecoder = Module['_Mix_GetMusicDecoder'] = makeInvalidEarlyAccess('_Mix_GetMusicDecoder');
var _Mix_HasMusicDecoder = Module['_Mix_HasMusicDecoder'] = makeInvalidEarlyAccess('_Mix_HasMusicDecoder');
var _Mix_HookMusicFinished = Module['_Mix_HookMusicFinished'] = makeInvalidEarlyAccess('_Mix_HookMusicFinished');
var _SDL_snprintf = Module['_SDL_snprintf'] = makeInvalidEarlyAccess('_SDL_snprintf');
var _SDL_GetHintBoolean = Module['_SDL_GetHintBoolean'] = makeInvalidEarlyAccess('_SDL_GetHintBoolean');
var _SDL_GetError = Module['_SDL_GetError'] = makeInvalidEarlyAccess('_SDL_GetError');
var _Mix_LoadMUS = Module['_Mix_LoadMUS'] = makeInvalidEarlyAccess('_Mix_LoadMUS');
var _Mix_LoadMUSType_IO = Module['_Mix_LoadMUSType_IO'] = makeInvalidEarlyAccess('_Mix_LoadMUSType_IO');
var _SDL_ClearError = Module['_SDL_ClearError'] = makeInvalidEarlyAccess('_SDL_ClearError');
var _Mix_LoadMUS_IO = Module['_Mix_LoadMUS_IO'] = makeInvalidEarlyAccess('_Mix_LoadMUS_IO');
var _Mix_FreeMusic = Module['_Mix_FreeMusic'] = makeInvalidEarlyAccess('_Mix_FreeMusic');
var _SDL_Delay = Module['_SDL_Delay'] = makeInvalidEarlyAccess('_SDL_Delay');
var _Mix_GetMusicType = Module['_Mix_GetMusicType'] = makeInvalidEarlyAccess('_Mix_GetMusicType');
var _Mix_GetMusicTitleTag = Module['_Mix_GetMusicTitleTag'] = makeInvalidEarlyAccess('_Mix_GetMusicTitleTag');
var _Mix_GetMusicTitle = Module['_Mix_GetMusicTitle'] = makeInvalidEarlyAccess('_Mix_GetMusicTitle');
var _Mix_GetMusicArtistTag = Module['_Mix_GetMusicArtistTag'] = makeInvalidEarlyAccess('_Mix_GetMusicArtistTag');
var _Mix_GetMusicAlbumTag = Module['_Mix_GetMusicAlbumTag'] = makeInvalidEarlyAccess('_Mix_GetMusicAlbumTag');
var _Mix_GetMusicCopyrightTag = Module['_Mix_GetMusicCopyrightTag'] = makeInvalidEarlyAccess('_Mix_GetMusicCopyrightTag');
var _Mix_FadeInMusicPos = Module['_Mix_FadeInMusicPos'] = makeInvalidEarlyAccess('_Mix_FadeInMusicPos');
var _Mix_FadeInMusic = Module['_Mix_FadeInMusic'] = makeInvalidEarlyAccess('_Mix_FadeInMusic');
var _Mix_PlayMusic = Module['_Mix_PlayMusic'] = makeInvalidEarlyAccess('_Mix_PlayMusic');
var _Mix_ModMusicJumpToOrder = Module['_Mix_ModMusicJumpToOrder'] = makeInvalidEarlyAccess('_Mix_ModMusicJumpToOrder');
var _Mix_SetMusicPosition = Module['_Mix_SetMusicPosition'] = makeInvalidEarlyAccess('_Mix_SetMusicPosition');
var _Mix_GetMusicPosition = Module['_Mix_GetMusicPosition'] = makeInvalidEarlyAccess('_Mix_GetMusicPosition');
var _Mix_MusicDuration = Module['_Mix_MusicDuration'] = makeInvalidEarlyAccess('_Mix_MusicDuration');
var _Mix_GetMusicLoopStartTime = Module['_Mix_GetMusicLoopStartTime'] = makeInvalidEarlyAccess('_Mix_GetMusicLoopStartTime');
var _Mix_GetMusicLoopEndTime = Module['_Mix_GetMusicLoopEndTime'] = makeInvalidEarlyAccess('_Mix_GetMusicLoopEndTime');
var _Mix_GetMusicLoopLengthTime = Module['_Mix_GetMusicLoopLengthTime'] = makeInvalidEarlyAccess('_Mix_GetMusicLoopLengthTime');
var _Mix_GetMusicVolume = Module['_Mix_GetMusicVolume'] = makeInvalidEarlyAccess('_Mix_GetMusicVolume');
var _Mix_HaltMusic = Module['_Mix_HaltMusic'] = makeInvalidEarlyAccess('_Mix_HaltMusic');
var _Mix_FadeOutMusic = Module['_Mix_FadeOutMusic'] = makeInvalidEarlyAccess('_Mix_FadeOutMusic');
var _Mix_FadingMusic = Module['_Mix_FadingMusic'] = makeInvalidEarlyAccess('_Mix_FadingMusic');
var _Mix_PauseMusic = Module['_Mix_PauseMusic'] = makeInvalidEarlyAccess('_Mix_PauseMusic');
var _Mix_ResumeMusic = Module['_Mix_ResumeMusic'] = makeInvalidEarlyAccess('_Mix_ResumeMusic');
var _Mix_RewindMusic = Module['_Mix_RewindMusic'] = makeInvalidEarlyAccess('_Mix_RewindMusic');
var _Mix_PausedMusic = Module['_Mix_PausedMusic'] = makeInvalidEarlyAccess('_Mix_PausedMusic');
var _Mix_StartTrack = Module['_Mix_StartTrack'] = makeInvalidEarlyAccess('_Mix_StartTrack');
var _Mix_GetNumTracks = Module['_Mix_GetNumTracks'] = makeInvalidEarlyAccess('_Mix_GetNumTracks');
var _Mix_PlayingMusic = Module['_Mix_PlayingMusic'] = makeInvalidEarlyAccess('_Mix_PlayingMusic');
var _Mix_SetTimidityCfg = Module['_Mix_SetTimidityCfg'] = makeInvalidEarlyAccess('_Mix_SetTimidityCfg');
var _Mix_SetSoundFonts = Module['_Mix_SetSoundFonts'] = makeInvalidEarlyAccess('_Mix_SetSoundFonts');
var _Mix_GetSoundFonts = Module['_Mix_GetSoundFonts'] = makeInvalidEarlyAccess('_Mix_GetSoundFonts');
var _Mix_EachSoundFont = Module['_Mix_EachSoundFont'] = makeInvalidEarlyAccess('_Mix_EachSoundFont');
var _SDL_strtok_r = Module['_SDL_strtok_r'] = makeInvalidEarlyAccess('_SDL_strtok_r');
var _SDL_atoi = Module['_SDL_atoi'] = makeInvalidEarlyAccess('_SDL_atoi');
var _SDL_atof = Module['_SDL_atof'] = makeInvalidEarlyAccess('_SDL_atof');
var _SDL_ReadS16BE = Module['_SDL_ReadS16BE'] = makeInvalidEarlyAccess('_SDL_ReadS16BE');
var _SDL_SetAppMetadata = Module['_SDL_SetAppMetadata'] = makeInvalidEarlyAccess('_SDL_SetAppMetadata');
var _SDL_GetGlobalProperties = Module['_SDL_GetGlobalProperties'] = makeInvalidEarlyAccess('_SDL_GetGlobalProperties');
var _SDL_SetAppMetadataProperty = Module['_SDL_SetAppMetadataProperty'] = makeInvalidEarlyAccess('_SDL_SetAppMetadataProperty');
var _SDL_GetAppMetadataProperty = Module['_SDL_GetAppMetadataProperty'] = makeInvalidEarlyAccess('_SDL_GetAppMetadataProperty');
var _SDL_GetHint = Module['_SDL_GetHint'] = makeInvalidEarlyAccess('_SDL_GetHint');
var _SDL_SetMainReady = Module['_SDL_SetMainReady'] = makeInvalidEarlyAccess('_SDL_SetMainReady');
var _SDL_GetCurrentThreadID = Module['_SDL_GetCurrentThreadID'] = makeInvalidEarlyAccess('_SDL_GetCurrentThreadID');
var _SDL_IsMainThread = Module['_SDL_IsMainThread'] = makeInvalidEarlyAccess('_SDL_IsMainThread');
var _SDL_LogInfo = Module['_SDL_LogInfo'] = makeInvalidEarlyAccess('_SDL_LogInfo');
var _SDL_QuitSubSystem = Module['_SDL_QuitSubSystem'] = makeInvalidEarlyAccess('_SDL_QuitSubSystem');
var _SDL_Init = Module['_SDL_Init'] = makeInvalidEarlyAccess('_SDL_Init');
var _SDL_Quit = Module['_SDL_Quit'] = makeInvalidEarlyAccess('_SDL_Quit');
var _SDL_GetVersion = Module['_SDL_GetVersion'] = makeInvalidEarlyAccess('_SDL_GetVersion');
var _SDL_GetRevision = Module['_SDL_GetRevision'] = makeInvalidEarlyAccess('_SDL_GetRevision');
var _SDL_GetPlatform = Module['_SDL_GetPlatform'] = makeInvalidEarlyAccess('_SDL_GetPlatform');
var _SDL_IsTablet = Module['_SDL_IsTablet'] = makeInvalidEarlyAccess('_SDL_IsTablet');
var _SDL_IsTV = Module['_SDL_IsTV'] = makeInvalidEarlyAccess('_SDL_IsTV');
var _SDL_GetSandbox = Module['_SDL_GetSandbox'] = makeInvalidEarlyAccess('_SDL_GetSandbox');
var _SDL_ReportAssertion = Module['_SDL_ReportAssertion'] = makeInvalidEarlyAccess('_SDL_ReportAssertion');
var _SDL_SetAssertionHandler = Module['_SDL_SetAssertionHandler'] = makeInvalidEarlyAccess('_SDL_SetAssertionHandler');
var _SDL_MinimizeWindow = Module['_SDL_MinimizeWindow'] = makeInvalidEarlyAccess('_SDL_MinimizeWindow');
var _SDL_ShowMessageBox = Module['_SDL_ShowMessageBox'] = makeInvalidEarlyAccess('_SDL_ShowMessageBox');
var _SDL_RestoreWindow = Module['_SDL_RestoreWindow'] = makeInvalidEarlyAccess('_SDL_RestoreWindow');
var _SDL_GetAssertionReport = Module['_SDL_GetAssertionReport'] = makeInvalidEarlyAccess('_SDL_GetAssertionReport');
var _SDL_ResetAssertionReport = Module['_SDL_ResetAssertionReport'] = makeInvalidEarlyAccess('_SDL_ResetAssertionReport');
var _SDL_GetDefaultAssertionHandler = Module['_SDL_GetDefaultAssertionHandler'] = makeInvalidEarlyAccess('_SDL_GetDefaultAssertionHandler');
var _SDL_GetAssertionHandler = Module['_SDL_GetAssertionHandler'] = makeInvalidEarlyAccess('_SDL_GetAssertionHandler');
var _SDL_LogMessageV = Module['_SDL_LogMessageV'] = makeInvalidEarlyAccess('_SDL_LogMessageV');
var _SDL_SetErrorV = Module['_SDL_SetErrorV'] = makeInvalidEarlyAccess('_SDL_SetErrorV');
var _SDL_vsnprintf = Module['_SDL_vsnprintf'] = makeInvalidEarlyAccess('_SDL_vsnprintf');
var _SDL_OutOfMemory = Module['_SDL_OutOfMemory'] = makeInvalidEarlyAccess('_SDL_OutOfMemory');
var _SDL_GUIDToString = Module['_SDL_GUIDToString'] = makeInvalidEarlyAccess('_SDL_GUIDToString');
var _SDL_StringToGUID = Module['_SDL_StringToGUID'] = makeInvalidEarlyAccess('_SDL_StringToGUID');
var _SDL_GetAtomicU32 = Module['_SDL_GetAtomicU32'] = makeInvalidEarlyAccess('_SDL_GetAtomicU32');
var _SDL_CompareAndSwapAtomicU32 = Module['_SDL_CompareAndSwapAtomicU32'] = makeInvalidEarlyAccess('_SDL_CompareAndSwapAtomicU32');
var _SDL_SetHintWithPriority = Module['_SDL_SetHintWithPriority'] = makeInvalidEarlyAccess('_SDL_SetHintWithPriority');
var _SDL_LockProperties = Module['_SDL_LockProperties'] = makeInvalidEarlyAccess('_SDL_LockProperties');
var _SDL_SetPointerPropertyWithCleanup = Module['_SDL_SetPointerPropertyWithCleanup'] = makeInvalidEarlyAccess('_SDL_SetPointerPropertyWithCleanup');
var _SDL_UnlockProperties = Module['_SDL_UnlockProperties'] = makeInvalidEarlyAccess('_SDL_UnlockProperties');
var _SDL_ResetHint = Module['_SDL_ResetHint'] = makeInvalidEarlyAccess('_SDL_ResetHint');
var _SDL_ResetHints = Module['_SDL_ResetHints'] = makeInvalidEarlyAccess('_SDL_ResetHints');
var _SDL_EnumerateProperties = Module['_SDL_EnumerateProperties'] = makeInvalidEarlyAccess('_SDL_EnumerateProperties');
var _SDL_SetHint = Module['_SDL_SetHint'] = makeInvalidEarlyAccess('_SDL_SetHint');
var _SDL_AddHintCallback = Module['_SDL_AddHintCallback'] = makeInvalidEarlyAccess('_SDL_AddHintCallback');
var _SDL_RemoveHintCallback = Module['_SDL_RemoveHintCallback'] = makeInvalidEarlyAccess('_SDL_RemoveHintCallback');
var _SDL_ResetLogPriorities = Module['_SDL_ResetLogPriorities'] = makeInvalidEarlyAccess('_SDL_ResetLogPriorities');
var _SDL_SetLogPriorities = Module['_SDL_SetLogPriorities'] = makeInvalidEarlyAccess('_SDL_SetLogPriorities');
var _SDL_SetLogPriority = Module['_SDL_SetLogPriority'] = makeInvalidEarlyAccess('_SDL_SetLogPriority');
var _SDL_GetLogPriority = Module['_SDL_GetLogPriority'] = makeInvalidEarlyAccess('_SDL_GetLogPriority');
var _SDL_SetLogPriorityPrefix = Module['_SDL_SetLogPriorityPrefix'] = makeInvalidEarlyAccess('_SDL_SetLogPriorityPrefix');
var _SDL_LogTrace = Module['_SDL_LogTrace'] = makeInvalidEarlyAccess('_SDL_LogTrace');
var _SDL_LogVerbose = Module['_SDL_LogVerbose'] = makeInvalidEarlyAccess('_SDL_LogVerbose');
var _SDL_LogDebug = Module['_SDL_LogDebug'] = makeInvalidEarlyAccess('_SDL_LogDebug');
var _SDL_LogWarn = Module['_SDL_LogWarn'] = makeInvalidEarlyAccess('_SDL_LogWarn');
var _SDL_LogError = Module['_SDL_LogError'] = makeInvalidEarlyAccess('_SDL_LogError');
var _SDL_LogCritical = Module['_SDL_LogCritical'] = makeInvalidEarlyAccess('_SDL_LogCritical');
var _SDL_LogMessage = Module['_SDL_LogMessage'] = makeInvalidEarlyAccess('_SDL_LogMessage');
var _SDL_GetDefaultLogOutputFunction = Module['_SDL_GetDefaultLogOutputFunction'] = makeInvalidEarlyAccess('_SDL_GetDefaultLogOutputFunction');
var _SDL_GetLogOutputFunction = Module['_SDL_GetLogOutputFunction'] = makeInvalidEarlyAccess('_SDL_GetLogOutputFunction');
var _SDL_SetLogOutputFunction = Module['_SDL_SetLogOutputFunction'] = makeInvalidEarlyAccess('_SDL_SetLogOutputFunction');
var _SDL_CopyProperties = Module['_SDL_CopyProperties'] = makeInvalidEarlyAccess('_SDL_CopyProperties');
var _SDL_ClearProperty = Module['_SDL_ClearProperty'] = makeInvalidEarlyAccess('_SDL_ClearProperty');
var _SDL_HasProperty = Module['_SDL_HasProperty'] = makeInvalidEarlyAccess('_SDL_HasProperty');
var _SDL_GetPropertyType = Module['_SDL_GetPropertyType'] = makeInvalidEarlyAccess('_SDL_GetPropertyType');
var _SDL_asprintf = Module['_SDL_asprintf'] = makeInvalidEarlyAccess('_SDL_asprintf');
var _SDL_round = Module['_SDL_round'] = makeInvalidEarlyAccess('_SDL_round');
var _SDL_GetTLS = Module['_SDL_GetTLS'] = makeInvalidEarlyAccess('_SDL_GetTLS');
var _SDL_SetTLS = Module['_SDL_SetTLS'] = makeInvalidEarlyAccess('_SDL_SetTLS');
var _SDL_tolower = Module['_SDL_tolower'] = makeInvalidEarlyAccess('_SDL_tolower');
var _SDL_CompareAndSwapAtomicInt = Module['_SDL_CompareAndSwapAtomicInt'] = makeInvalidEarlyAccess('_SDL_CompareAndSwapAtomicInt');
var _SDL_CompareAndSwapAtomicPointer = Module['_SDL_CompareAndSwapAtomicPointer'] = makeInvalidEarlyAccess('_SDL_CompareAndSwapAtomicPointer');
var _SDL_SetAtomicU32 = Module['_SDL_SetAtomicU32'] = makeInvalidEarlyAccess('_SDL_SetAtomicU32');
var _SDL_SetAtomicPointer = Module['_SDL_SetAtomicPointer'] = makeInvalidEarlyAccess('_SDL_SetAtomicPointer');
var _SDL_GetAtomicPointer = Module['_SDL_GetAtomicPointer'] = makeInvalidEarlyAccess('_SDL_GetAtomicPointer');
var _SDL_MemoryBarrierReleaseFunction = Module['_SDL_MemoryBarrierReleaseFunction'] = makeInvalidEarlyAccess('_SDL_MemoryBarrierReleaseFunction');
var _SDL_LockSpinlock = Module['_SDL_LockSpinlock'] = makeInvalidEarlyAccess('_SDL_LockSpinlock');
var _SDL_UnlockSpinlock = Module['_SDL_UnlockSpinlock'] = makeInvalidEarlyAccess('_SDL_UnlockSpinlock');
var _SDL_MemoryBarrierAcquireFunction = Module['_SDL_MemoryBarrierAcquireFunction'] = makeInvalidEarlyAccess('_SDL_MemoryBarrierAcquireFunction');
var _SDL_TryLockSpinlock = Module['_SDL_TryLockSpinlock'] = makeInvalidEarlyAccess('_SDL_TryLockSpinlock');
var _SDL_GetNumAudioDrivers = Module['_SDL_GetNumAudioDrivers'] = makeInvalidEarlyAccess('_SDL_GetNumAudioDrivers');
var _SDL_GetAudioDriver = Module['_SDL_GetAudioDriver'] = makeInvalidEarlyAccess('_SDL_GetAudioDriver');
var _SDL_GetCurrentAudioDriver = Module['_SDL_GetCurrentAudioDriver'] = makeInvalidEarlyAccess('_SDL_GetCurrentAudioDriver');
var _SDL_IsAudioDevicePhysical = Module['_SDL_IsAudioDevicePhysical'] = makeInvalidEarlyAccess('_SDL_IsAudioDevicePhysical');
var _SDL_IsAudioDevicePlayback = Module['_SDL_IsAudioDevicePlayback'] = makeInvalidEarlyAccess('_SDL_IsAudioDevicePlayback');
var _SDL_DestroyCondition = Module['_SDL_DestroyCondition'] = makeInvalidEarlyAccess('_SDL_DestroyCondition');
var _SDL_GetAudioPlaybackDevices = Module['_SDL_GetAudioPlaybackDevices'] = makeInvalidEarlyAccess('_SDL_GetAudioPlaybackDevices');
var _SDL_GetAudioRecordingDevices = Module['_SDL_GetAudioRecordingDevices'] = makeInvalidEarlyAccess('_SDL_GetAudioRecordingDevices');
var _SDL_GetAudioDeviceName = Module['_SDL_GetAudioDeviceName'] = makeInvalidEarlyAccess('_SDL_GetAudioDeviceName');
var _SDL_GetAudioDeviceChannelMap = Module['_SDL_GetAudioDeviceChannelMap'] = makeInvalidEarlyAccess('_SDL_GetAudioDeviceChannelMap');
var _SDL_WaitCondition = Module['_SDL_WaitCondition'] = makeInvalidEarlyAccess('_SDL_WaitCondition');
var _SDL_WaitThread = Module['_SDL_WaitThread'] = makeInvalidEarlyAccess('_SDL_WaitThread');
var _SDL_BroadcastCondition = Module['_SDL_BroadcastCondition'] = makeInvalidEarlyAccess('_SDL_BroadcastCondition');
var _SDL_CreateThreadRuntime = Module['_SDL_CreateThreadRuntime'] = makeInvalidEarlyAccess('_SDL_CreateThreadRuntime');
var _SDL_AudioDevicePaused = Module['_SDL_AudioDevicePaused'] = makeInvalidEarlyAccess('_SDL_AudioDevicePaused');
var _SDL_GetAudioDeviceGain = Module['_SDL_GetAudioDeviceGain'] = makeInvalidEarlyAccess('_SDL_GetAudioDeviceGain');
var _SDL_SetAudioDeviceGain = Module['_SDL_SetAudioDeviceGain'] = makeInvalidEarlyAccess('_SDL_SetAudioDeviceGain');
var _SDL_SetAudioPostmixCallback = Module['_SDL_SetAudioPostmixCallback'] = makeInvalidEarlyAccess('_SDL_SetAudioPostmixCallback');
var _SDL_BindAudioStreams = Module['_SDL_BindAudioStreams'] = makeInvalidEarlyAccess('_SDL_BindAudioStreams');
var _SDL_UnbindAudioStreams = Module['_SDL_UnbindAudioStreams'] = makeInvalidEarlyAccess('_SDL_UnbindAudioStreams');
var _SDL_UnbindAudioStream = Module['_SDL_UnbindAudioStream'] = makeInvalidEarlyAccess('_SDL_UnbindAudioStream');
var _SDL_GetAudioStreamDevice = Module['_SDL_GetAudioStreamDevice'] = makeInvalidEarlyAccess('_SDL_GetAudioStreamDevice');
var _SDL_OpenAudioDeviceStream = Module['_SDL_OpenAudioDeviceStream'] = makeInvalidEarlyAccess('_SDL_OpenAudioDeviceStream');
var _SDL_SetAudioStreamPutCallback = Module['_SDL_SetAudioStreamPutCallback'] = makeInvalidEarlyAccess('_SDL_SetAudioStreamPutCallback');
var _SDL_PauseAudioStreamDevice = Module['_SDL_PauseAudioStreamDevice'] = makeInvalidEarlyAccess('_SDL_PauseAudioStreamDevice');
var _SDL_ResumeAudioStreamDevice = Module['_SDL_ResumeAudioStreamDevice'] = makeInvalidEarlyAccess('_SDL_ResumeAudioStreamDevice');
var _SDL_AudioStreamDevicePaused = Module['_SDL_AudioStreamDevicePaused'] = makeInvalidEarlyAccess('_SDL_AudioStreamDevicePaused');
var _SDL_GetAudioFormatName = Module['_SDL_GetAudioFormatName'] = makeInvalidEarlyAccess('_SDL_GetAudioFormatName');
var _SDL_EventEnabled = Module['_SDL_EventEnabled'] = makeInvalidEarlyAccess('_SDL_EventEnabled');
var _SDL_PushEvent = Module['_SDL_PushEvent'] = makeInvalidEarlyAccess('_SDL_PushEvent');
var _SDL_CreateCondition = Module['_SDL_CreateCondition'] = makeInvalidEarlyAccess('_SDL_CreateCondition');
var _SDL_SetCurrentThreadPriority = Module['_SDL_SetCurrentThreadPriority'] = makeInvalidEarlyAccess('_SDL_SetCurrentThreadPriority');
var _SDL_SetAudioStreamFormat = Module['_SDL_SetAudioStreamFormat'] = makeInvalidEarlyAccess('_SDL_SetAudioStreamFormat');
var _SDL_GetAudioStreamProperties = Module['_SDL_GetAudioStreamProperties'] = makeInvalidEarlyAccess('_SDL_GetAudioStreamProperties');
var _SDL_GetAudioStreamFormat = Module['_SDL_GetAudioStreamFormat'] = makeInvalidEarlyAccess('_SDL_GetAudioStreamFormat');
var _SDL_SetAudioStreamInputChannelMap = Module['_SDL_SetAudioStreamInputChannelMap'] = makeInvalidEarlyAccess('_SDL_SetAudioStreamInputChannelMap');
var _SDL_SetAudioStreamOutputChannelMap = Module['_SDL_SetAudioStreamOutputChannelMap'] = makeInvalidEarlyAccess('_SDL_SetAudioStreamOutputChannelMap');
var _SDL_GetAudioStreamInputChannelMap = Module['_SDL_GetAudioStreamInputChannelMap'] = makeInvalidEarlyAccess('_SDL_GetAudioStreamInputChannelMap');
var _SDL_GetAudioStreamOutputChannelMap = Module['_SDL_GetAudioStreamOutputChannelMap'] = makeInvalidEarlyAccess('_SDL_GetAudioStreamOutputChannelMap');
var _SDL_GetAudioStreamFrequencyRatio = Module['_SDL_GetAudioStreamFrequencyRatio'] = makeInvalidEarlyAccess('_SDL_GetAudioStreamFrequencyRatio');
var _SDL_SetAudioStreamFrequencyRatio = Module['_SDL_SetAudioStreamFrequencyRatio'] = makeInvalidEarlyAccess('_SDL_SetAudioStreamFrequencyRatio');
var _SDL_GetAudioStreamGain = Module['_SDL_GetAudioStreamGain'] = makeInvalidEarlyAccess('_SDL_GetAudioStreamGain');
var _SDL_SetAudioStreamGain = Module['_SDL_SetAudioStreamGain'] = makeInvalidEarlyAccess('_SDL_SetAudioStreamGain');
var _SDL_GetAudioStreamAvailable = Module['_SDL_GetAudioStreamAvailable'] = makeInvalidEarlyAccess('_SDL_GetAudioStreamAvailable');
var _SDL_GetAudioStreamQueued = Module['_SDL_GetAudioStreamQueued'] = makeInvalidEarlyAccess('_SDL_GetAudioStreamQueued');
var _SDL_LoadWAV = Module['_SDL_LoadWAV'] = makeInvalidEarlyAccess('_SDL_LoadWAV');
var _SDL_GetNumCameraDrivers = Module['_SDL_GetNumCameraDrivers'] = makeInvalidEarlyAccess('_SDL_GetNumCameraDrivers');
var _SDL_GetCameraDriver = Module['_SDL_GetCameraDriver'] = makeInvalidEarlyAccess('_SDL_GetCameraDriver');
var _SDL_GetCurrentCameraDriver = Module['_SDL_GetCurrentCameraDriver'] = makeInvalidEarlyAccess('_SDL_GetCurrentCameraDriver');
var _SDL_GetTicksNS = Module['_SDL_GetTicksNS'] = makeInvalidEarlyAccess('_SDL_GetTicksNS');
var _SDL_CloseCamera = Module['_SDL_CloseCamera'] = makeInvalidEarlyAccess('_SDL_CloseCamera');
var _SDL_GetCameraFormat = Module['_SDL_GetCameraFormat'] = makeInvalidEarlyAccess('_SDL_GetCameraFormat');
var _SDL_GetCameraName = Module['_SDL_GetCameraName'] = makeInvalidEarlyAccess('_SDL_GetCameraName');
var _SDL_GetCameraPosition = Module['_SDL_GetCameraPosition'] = makeInvalidEarlyAccess('_SDL_GetCameraPosition');
var _SDL_GetCameras = Module['_SDL_GetCameras'] = makeInvalidEarlyAccess('_SDL_GetCameras');
var _SDL_GetCameraSupportedFormats = Module['_SDL_GetCameraSupportedFormats'] = makeInvalidEarlyAccess('_SDL_GetCameraSupportedFormats');
var _SDL_StretchSurface = Module['_SDL_StretchSurface'] = makeInvalidEarlyAccess('_SDL_StretchSurface');
var _SDL_ConvertPixels = Module['_SDL_ConvertPixels'] = makeInvalidEarlyAccess('_SDL_ConvertPixels');
var _SDL_SetSurfaceColorspace = Module['_SDL_SetSurfaceColorspace'] = makeInvalidEarlyAccess('_SDL_SetSurfaceColorspace');
var _SDL_OpenCamera = Module['_SDL_OpenCamera'] = makeInvalidEarlyAccess('_SDL_OpenCamera');
var _SDL_AcquireCameraFrame = Module['_SDL_AcquireCameraFrame'] = makeInvalidEarlyAccess('_SDL_AcquireCameraFrame');
var _SDL_ReleaseCameraFrame = Module['_SDL_ReleaseCameraFrame'] = makeInvalidEarlyAccess('_SDL_ReleaseCameraFrame');
var _SDL_GetCameraID = Module['_SDL_GetCameraID'] = makeInvalidEarlyAccess('_SDL_GetCameraID');
var _SDL_GetCameraProperties = Module['_SDL_GetCameraProperties'] = makeInvalidEarlyAccess('_SDL_GetCameraProperties');
var _SDL_GetCameraPermissionState = Module['_SDL_GetCameraPermissionState'] = makeInvalidEarlyAccess('_SDL_GetCameraPermissionState');
var _SDL_SetX11EventHook = Module['_SDL_SetX11EventHook'] = makeInvalidEarlyAccess('_SDL_SetX11EventHook');
var _SDL_SetLinuxThreadPriority = Module['_SDL_SetLinuxThreadPriority'] = makeInvalidEarlyAccess('_SDL_SetLinuxThreadPriority');
var _SDL_SetLinuxThreadPriorityAndPolicy = Module['_SDL_SetLinuxThreadPriorityAndPolicy'] = makeInvalidEarlyAccess('_SDL_SetLinuxThreadPriorityAndPolicy');
var _SDL_GDKSuspendComplete = Module['_SDL_GDKSuspendComplete'] = makeInvalidEarlyAccess('_SDL_GDKSuspendComplete');
var _SDL_GetGDKDefaultUser = Module['_SDL_GetGDKDefaultUser'] = makeInvalidEarlyAccess('_SDL_GetGDKDefaultUser');
var _SDL_GDKSuspendGPU = Module['_SDL_GDKSuspendGPU'] = makeInvalidEarlyAccess('_SDL_GDKSuspendGPU');
var _SDL_GDKResumeGPU = Module['_SDL_GDKResumeGPU'] = makeInvalidEarlyAccess('_SDL_GDKResumeGPU');
var _SDL_RegisterApp = Module['_SDL_RegisterApp'] = makeInvalidEarlyAccess('_SDL_RegisterApp');
var _SDL_SetWindowsMessageHook = Module['_SDL_SetWindowsMessageHook'] = makeInvalidEarlyAccess('_SDL_SetWindowsMessageHook');
var _SDL_UnregisterApp = Module['_SDL_UnregisterApp'] = makeInvalidEarlyAccess('_SDL_UnregisterApp');
var _SDL_SendAndroidBackButton = Module['_SDL_SendAndroidBackButton'] = makeInvalidEarlyAccess('_SDL_SendAndroidBackButton');
var _SDL_GetAndroidActivity = Module['_SDL_GetAndroidActivity'] = makeInvalidEarlyAccess('_SDL_GetAndroidActivity');
var _SDL_GetAndroidCachePath = Module['_SDL_GetAndroidCachePath'] = makeInvalidEarlyAccess('_SDL_GetAndroidCachePath');
var _SDL_GetAndroidExternalStoragePath = Module['_SDL_GetAndroidExternalStoragePath'] = makeInvalidEarlyAccess('_SDL_GetAndroidExternalStoragePath');
var _SDL_GetAndroidExternalStorageState = Module['_SDL_GetAndroidExternalStorageState'] = makeInvalidEarlyAccess('_SDL_GetAndroidExternalStorageState');
var _SDL_GetAndroidInternalStoragePath = Module['_SDL_GetAndroidInternalStoragePath'] = makeInvalidEarlyAccess('_SDL_GetAndroidInternalStoragePath');
var _SDL_GetAndroidJNIEnv = Module['_SDL_GetAndroidJNIEnv'] = makeInvalidEarlyAccess('_SDL_GetAndroidJNIEnv');
var _SDL_RequestAndroidPermission = Module['_SDL_RequestAndroidPermission'] = makeInvalidEarlyAccess('_SDL_RequestAndroidPermission');
var _SDL_SendAndroidMessage = Module['_SDL_SendAndroidMessage'] = makeInvalidEarlyAccess('_SDL_SendAndroidMessage');
var _SDL_ShowAndroidToast = Module['_SDL_ShowAndroidToast'] = makeInvalidEarlyAccess('_SDL_ShowAndroidToast');
var _SDL_GetAndroidSDKVersion = Module['_SDL_GetAndroidSDKVersion'] = makeInvalidEarlyAccess('_SDL_GetAndroidSDKVersion');
var _SDL_IsChromebook = Module['_SDL_IsChromebook'] = makeInvalidEarlyAccess('_SDL_IsChromebook');
var _SDL_IsDeXMode = Module['_SDL_IsDeXMode'] = makeInvalidEarlyAccess('_SDL_IsDeXMode');
var _JNI_OnLoad = Module['_JNI_OnLoad'] = makeInvalidEarlyAccess('_JNI_OnLoad');
var _SDL_GetNumLogicalCPUCores = Module['_SDL_GetNumLogicalCPUCores'] = makeInvalidEarlyAccess('_SDL_GetNumLogicalCPUCores');
var _SDL_GetCPUCacheLineSize = Module['_SDL_GetCPUCacheLineSize'] = makeInvalidEarlyAccess('_SDL_GetCPUCacheLineSize');
var _SDL_HasAltiVec = Module['_SDL_HasAltiVec'] = makeInvalidEarlyAccess('_SDL_HasAltiVec');
var _SDL_HasMMX = Module['_SDL_HasMMX'] = makeInvalidEarlyAccess('_SDL_HasMMX');
var _SDL_HasSSE = Module['_SDL_HasSSE'] = makeInvalidEarlyAccess('_SDL_HasSSE');
var _SDL_HasSSE2 = Module['_SDL_HasSSE2'] = makeInvalidEarlyAccess('_SDL_HasSSE2');
var _SDL_HasSSE3 = Module['_SDL_HasSSE3'] = makeInvalidEarlyAccess('_SDL_HasSSE3');
var _SDL_HasSSE41 = Module['_SDL_HasSSE41'] = makeInvalidEarlyAccess('_SDL_HasSSE41');
var _SDL_HasSSE42 = Module['_SDL_HasSSE42'] = makeInvalidEarlyAccess('_SDL_HasSSE42');
var _SDL_HasAVX = Module['_SDL_HasAVX'] = makeInvalidEarlyAccess('_SDL_HasAVX');
var _SDL_HasAVX2 = Module['_SDL_HasAVX2'] = makeInvalidEarlyAccess('_SDL_HasAVX2');
var _SDL_HasAVX512F = Module['_SDL_HasAVX512F'] = makeInvalidEarlyAccess('_SDL_HasAVX512F');
var _SDL_HasARMSIMD = Module['_SDL_HasARMSIMD'] = makeInvalidEarlyAccess('_SDL_HasARMSIMD');
var _SDL_HasNEON = Module['_SDL_HasNEON'] = makeInvalidEarlyAccess('_SDL_HasNEON');
var _SDL_HasLSX = Module['_SDL_HasLSX'] = makeInvalidEarlyAccess('_SDL_HasLSX');
var _SDL_HasLASX = Module['_SDL_HasLASX'] = makeInvalidEarlyAccess('_SDL_HasLASX');
var _SDL_GetSystemRAM = Module['_SDL_GetSystemRAM'] = makeInvalidEarlyAccess('_SDL_GetSystemRAM');
var _SDL_GetWindowFromEvent = Module['_SDL_GetWindowFromEvent'] = makeInvalidEarlyAccess('_SDL_GetWindowFromEvent');
var _SDL_GetWindowFromID = Module['_SDL_GetWindowFromID'] = makeInvalidEarlyAccess('_SDL_GetWindowFromID');
var _SDL_GetCurrentVideoDriver = Module['_SDL_GetCurrentVideoDriver'] = makeInvalidEarlyAccess('_SDL_GetCurrentVideoDriver');
var _SDL_PeepEvents = Module['_SDL_PeepEvents'] = makeInvalidEarlyAccess('_SDL_PeepEvents');
var _SDL_HasEvent = Module['_SDL_HasEvent'] = makeInvalidEarlyAccess('_SDL_HasEvent');
var _SDL_HasEvents = Module['_SDL_HasEvents'] = makeInvalidEarlyAccess('_SDL_HasEvents');
var _SDL_FlushEvent = Module['_SDL_FlushEvent'] = makeInvalidEarlyAccess('_SDL_FlushEvent');
var _SDL_FlushEvents = Module['_SDL_FlushEvents'] = makeInvalidEarlyAccess('_SDL_FlushEvents');
var _SDL_RunOnMainThread = Module['_SDL_RunOnMainThread'] = makeInvalidEarlyAccess('_SDL_RunOnMainThread');
var _SDL_CreateSemaphore = Module['_SDL_CreateSemaphore'] = makeInvalidEarlyAccess('_SDL_CreateSemaphore');
var _SDL_WaitSemaphore = Module['_SDL_WaitSemaphore'] = makeInvalidEarlyAccess('_SDL_WaitSemaphore');
var _SDL_DestroySemaphore = Module['_SDL_DestroySemaphore'] = makeInvalidEarlyAccess('_SDL_DestroySemaphore');
var _SDL_UpdateSensors = Module['_SDL_UpdateSensors'] = makeInvalidEarlyAccess('_SDL_UpdateSensors');
var _SDL_UpdateJoysticks = Module['_SDL_UpdateJoysticks'] = makeInvalidEarlyAccess('_SDL_UpdateJoysticks');
var _SDL_UpdateTrays = Module['_SDL_UpdateTrays'] = makeInvalidEarlyAccess('_SDL_UpdateTrays');
var _SDL_PumpEvents = Module['_SDL_PumpEvents'] = makeInvalidEarlyAccess('_SDL_PumpEvents');
var _SDL_SignalSemaphore = Module['_SDL_SignalSemaphore'] = makeInvalidEarlyAccess('_SDL_SignalSemaphore');
var _SDL_PollEvent = Module['_SDL_PollEvent'] = makeInvalidEarlyAccess('_SDL_PollEvent');
var _SDL_DelayNS = Module['_SDL_DelayNS'] = makeInvalidEarlyAccess('_SDL_DelayNS');
var _SDL_WaitEvent = Module['_SDL_WaitEvent'] = makeInvalidEarlyAccess('_SDL_WaitEvent');
var _SDL_WaitEventTimeout = Module['_SDL_WaitEventTimeout'] = makeInvalidEarlyAccess('_SDL_WaitEventTimeout');
var _SDL_SetEventFilter = Module['_SDL_SetEventFilter'] = makeInvalidEarlyAccess('_SDL_SetEventFilter');
var _SDL_GetEventFilter = Module['_SDL_GetEventFilter'] = makeInvalidEarlyAccess('_SDL_GetEventFilter');
var _SDL_AddEventWatch = Module['_SDL_AddEventWatch'] = makeInvalidEarlyAccess('_SDL_AddEventWatch');
var _SDL_RemoveEventWatch = Module['_SDL_RemoveEventWatch'] = makeInvalidEarlyAccess('_SDL_RemoveEventWatch');
var _SDL_FilterEvents = Module['_SDL_FilterEvents'] = makeInvalidEarlyAccess('_SDL_FilterEvents');
var _SDL_SetEventEnabled = Module['_SDL_SetEventEnabled'] = makeInvalidEarlyAccess('_SDL_SetEventEnabled');
var _SDL_RegisterEvents = Module['_SDL_RegisterEvents'] = makeInvalidEarlyAccess('_SDL_RegisterEvents');
var _SDL_HasKeyboard = Module['_SDL_HasKeyboard'] = makeInvalidEarlyAccess('_SDL_HasKeyboard');
var _SDL_GetKeyboards = Module['_SDL_GetKeyboards'] = makeInvalidEarlyAccess('_SDL_GetKeyboards');
var _SDL_GetKeyboardNameForID = Module['_SDL_GetKeyboardNameForID'] = makeInvalidEarlyAccess('_SDL_GetKeyboardNameForID');
var _SDL_ResetKeyboard = Module['_SDL_ResetKeyboard'] = makeInvalidEarlyAccess('_SDL_ResetKeyboard');
var _SDL_GetKeyboardFocus = Module['_SDL_GetKeyboardFocus'] = makeInvalidEarlyAccess('_SDL_GetKeyboardFocus');
var _SDL_WarpMouseGlobal = Module['_SDL_WarpMouseGlobal'] = makeInvalidEarlyAccess('_SDL_WarpMouseGlobal');
var _SDL_TextInputActive = Module['_SDL_TextInputActive'] = makeInvalidEarlyAccess('_SDL_TextInputActive');
var _SDL_GetKeyFromScancode = Module['_SDL_GetKeyFromScancode'] = makeInvalidEarlyAccess('_SDL_GetKeyFromScancode');
var _SDL_GetScancodeFromKey = Module['_SDL_GetScancodeFromKey'] = makeInvalidEarlyAccess('_SDL_GetScancodeFromKey');
var _SDL_GetModState = Module['_SDL_GetModState'] = makeInvalidEarlyAccess('_SDL_GetModState');
var _SDL_iscntrl = Module['_SDL_iscntrl'] = makeInvalidEarlyAccess('_SDL_iscntrl');
var _SDL_GetKeyboardState = Module['_SDL_GetKeyboardState'] = makeInvalidEarlyAccess('_SDL_GetKeyboardState');
var _SDL_SetModState = Module['_SDL_SetModState'] = makeInvalidEarlyAccess('_SDL_SetModState');
var _SDL_SetScancodeName = Module['_SDL_SetScancodeName'] = makeInvalidEarlyAccess('_SDL_SetScancodeName');
var _SDL_GetScancodeName = Module['_SDL_GetScancodeName'] = makeInvalidEarlyAccess('_SDL_GetScancodeName');
var _SDL_GetScancodeFromName = Module['_SDL_GetScancodeFromName'] = makeInvalidEarlyAccess('_SDL_GetScancodeFromName');
var _SDL_GetKeyName = Module['_SDL_GetKeyName'] = makeInvalidEarlyAccess('_SDL_GetKeyName');
var _SDL_GetKeyFromName = Module['_SDL_GetKeyFromName'] = makeInvalidEarlyAccess('_SDL_GetKeyFromName');
var _SDL_CreateColorCursor = Module['_SDL_CreateColorCursor'] = makeInvalidEarlyAccess('_SDL_CreateColorCursor');
var _SDL_HasMouse = Module['_SDL_HasMouse'] = makeInvalidEarlyAccess('_SDL_HasMouse');
var _SDL_GetMice = Module['_SDL_GetMice'] = makeInvalidEarlyAccess('_SDL_GetMice');
var _SDL_GetMouseNameForID = Module['_SDL_GetMouseNameForID'] = makeInvalidEarlyAccess('_SDL_GetMouseNameForID');
var _SDL_SetCursor = Module['_SDL_SetCursor'] = makeInvalidEarlyAccess('_SDL_SetCursor');
var _SDL_GetMouseFocus = Module['_SDL_GetMouseFocus'] = makeInvalidEarlyAccess('_SDL_GetMouseFocus');
var _SDL_modff = Module['_SDL_modff'] = makeInvalidEarlyAccess('_SDL_modff');
var _SDL_truncf = Module['_SDL_truncf'] = makeInvalidEarlyAccess('_SDL_truncf');
var _SDL_DestroyCursor = Module['_SDL_DestroyCursor'] = makeInvalidEarlyAccess('_SDL_DestroyCursor');
var _SDL_CaptureMouse = Module['_SDL_CaptureMouse'] = makeInvalidEarlyAccess('_SDL_CaptureMouse');
var _SDL_ShowCursor = Module['_SDL_ShowCursor'] = makeInvalidEarlyAccess('_SDL_ShowCursor');
var _SDL_GetMouseState = Module['_SDL_GetMouseState'] = makeInvalidEarlyAccess('_SDL_GetMouseState');
var _SDL_GetRelativeMouseState = Module['_SDL_GetRelativeMouseState'] = makeInvalidEarlyAccess('_SDL_GetRelativeMouseState');
var _SDL_GetGlobalMouseState = Module['_SDL_GetGlobalMouseState'] = makeInvalidEarlyAccess('_SDL_GetGlobalMouseState');
var _SDL_WarpMouseInWindow = Module['_SDL_WarpMouseInWindow'] = makeInvalidEarlyAccess('_SDL_WarpMouseInWindow');
var _SDL_CreateCursor = Module['_SDL_CreateCursor'] = makeInvalidEarlyAccess('_SDL_CreateCursor');
var _SDL_CreateSystemCursor = Module['_SDL_CreateSystemCursor'] = makeInvalidEarlyAccess('_SDL_CreateSystemCursor');
var _SDL_GetCursor = Module['_SDL_GetCursor'] = makeInvalidEarlyAccess('_SDL_GetCursor');
var _SDL_GetDefaultCursor = Module['_SDL_GetDefaultCursor'] = makeInvalidEarlyAccess('_SDL_GetDefaultCursor');
var _SDL_HideCursor = Module['_SDL_HideCursor'] = makeInvalidEarlyAccess('_SDL_HideCursor');
var _SDL_CursorVisible = Module['_SDL_CursorVisible'] = makeInvalidEarlyAccess('_SDL_CursorVisible');
var _SDL_GetWindowMouseRect = Module['_SDL_GetWindowMouseRect'] = makeInvalidEarlyAccess('_SDL_GetWindowMouseRect');
var _SDL_GetRectIntersection = Module['_SDL_GetRectIntersection'] = makeInvalidEarlyAccess('_SDL_GetRectIntersection');
var _SDL_GetTouchDevices = Module['_SDL_GetTouchDevices'] = makeInvalidEarlyAccess('_SDL_GetTouchDevices');
var _SDL_GetTouchDeviceName = Module['_SDL_GetTouchDeviceName'] = makeInvalidEarlyAccess('_SDL_GetTouchDeviceName');
var _SDL_GetTouchDeviceType = Module['_SDL_GetTouchDeviceType'] = makeInvalidEarlyAccess('_SDL_GetTouchDeviceType');
var _SDL_GetTouchFingers = Module['_SDL_GetTouchFingers'] = makeInvalidEarlyAccess('_SDL_GetTouchFingers');
var _SDL_GetWindowID = Module['_SDL_GetWindowID'] = makeInvalidEarlyAccess('_SDL_GetWindowID');
var _SDL_RemovePath = Module['_SDL_RemovePath'] = makeInvalidEarlyAccess('_SDL_RemovePath');
var _SDL_RenamePath = Module['_SDL_RenamePath'] = makeInvalidEarlyAccess('_SDL_RenamePath');
var _SDL_CopyFile = Module['_SDL_CopyFile'] = makeInvalidEarlyAccess('_SDL_CopyFile');
var _SDL_CreateDirectory = Module['_SDL_CreateDirectory'] = makeInvalidEarlyAccess('_SDL_CreateDirectory');
var _SDL_EnumerateDirectory = Module['_SDL_EnumerateDirectory'] = makeInvalidEarlyAccess('_SDL_EnumerateDirectory');
var _SDL_GetPathInfo = Module['_SDL_GetPathInfo'] = makeInvalidEarlyAccess('_SDL_GetPathInfo');
var _SDL_IOFromDynamicMem = Module['_SDL_IOFromDynamicMem'] = makeInvalidEarlyAccess('_SDL_IOFromDynamicMem');
var _SDL_GlobDirectory = Module['_SDL_GlobDirectory'] = makeInvalidEarlyAccess('_SDL_GlobDirectory');
var _SDL_GetBasePath = Module['_SDL_GetBasePath'] = makeInvalidEarlyAccess('_SDL_GetBasePath');
var _SDL_GetUserFolder = Module['_SDL_GetUserFolder'] = makeInvalidEarlyAccess('_SDL_GetUserFolder');
var _SDL_GetPrefPath = Module['_SDL_GetPrefPath'] = makeInvalidEarlyAccess('_SDL_GetPrefPath');
var _SDL_GetCurrentDirectory = Module['_SDL_GetCurrentDirectory'] = makeInvalidEarlyAccess('_SDL_GetCurrentDirectory');
var _SDL_CreateGPUGraphicsPipeline = Module['_SDL_CreateGPUGraphicsPipeline'] = makeInvalidEarlyAccess('_SDL_CreateGPUGraphicsPipeline');
var _SDL_GPUTextureSupportsFormat = Module['_SDL_GPUTextureSupportsFormat'] = makeInvalidEarlyAccess('_SDL_GPUTextureSupportsFormat');
var _SDL_BindGPUFragmentSamplers = Module['_SDL_BindGPUFragmentSamplers'] = makeInvalidEarlyAccess('_SDL_BindGPUFragmentSamplers');
var _SDL_DrawGPUPrimitives = Module['_SDL_DrawGPUPrimitives'] = makeInvalidEarlyAccess('_SDL_DrawGPUPrimitives');
var _SDL_SetGPUViewport = Module['_SDL_SetGPUViewport'] = makeInvalidEarlyAccess('_SDL_SetGPUViewport');
var _SDL_BindGPUGraphicsPipeline = Module['_SDL_BindGPUGraphicsPipeline'] = makeInvalidEarlyAccess('_SDL_BindGPUGraphicsPipeline');
var _SDL_PushGPUFragmentUniformData = Module['_SDL_PushGPUFragmentUniformData'] = makeInvalidEarlyAccess('_SDL_PushGPUFragmentUniformData');
var _SDL_GPUSupportsShaderFormats = Module['_SDL_GPUSupportsShaderFormats'] = makeInvalidEarlyAccess('_SDL_GPUSupportsShaderFormats');
var _SDL_GPUSupportsProperties = Module['_SDL_GPUSupportsProperties'] = makeInvalidEarlyAccess('_SDL_GPUSupportsProperties');
var _SDL_CreateGPUDevice = Module['_SDL_CreateGPUDevice'] = makeInvalidEarlyAccess('_SDL_CreateGPUDevice');
var _SDL_CreateGPUDeviceWithProperties = Module['_SDL_CreateGPUDeviceWithProperties'] = makeInvalidEarlyAccess('_SDL_CreateGPUDeviceWithProperties');
var _SDL_DestroyGPUDevice = Module['_SDL_DestroyGPUDevice'] = makeInvalidEarlyAccess('_SDL_DestroyGPUDevice');
var _SDL_GetNumGPUDrivers = Module['_SDL_GetNumGPUDrivers'] = makeInvalidEarlyAccess('_SDL_GetNumGPUDrivers');
var _SDL_GetGPUDriver = Module['_SDL_GetGPUDriver'] = makeInvalidEarlyAccess('_SDL_GetGPUDriver');
var _SDL_GetGPUDeviceDriver = Module['_SDL_GetGPUDeviceDriver'] = makeInvalidEarlyAccess('_SDL_GetGPUDeviceDriver');
var _SDL_GetGPUShaderFormats = Module['_SDL_GetGPUShaderFormats'] = makeInvalidEarlyAccess('_SDL_GetGPUShaderFormats');
var _SDL_GPUTextureFormatTexelBlockSize = Module['_SDL_GPUTextureFormatTexelBlockSize'] = makeInvalidEarlyAccess('_SDL_GPUTextureFormatTexelBlockSize');
var _SDL_GPUTextureSupportsSampleCount = Module['_SDL_GPUTextureSupportsSampleCount'] = makeInvalidEarlyAccess('_SDL_GPUTextureSupportsSampleCount');
var _SDL_CreateGPUComputePipeline = Module['_SDL_CreateGPUComputePipeline'] = makeInvalidEarlyAccess('_SDL_CreateGPUComputePipeline');
var _SDL_CreateGPUSampler = Module['_SDL_CreateGPUSampler'] = makeInvalidEarlyAccess('_SDL_CreateGPUSampler');
var _SDL_CreateGPUShader = Module['_SDL_CreateGPUShader'] = makeInvalidEarlyAccess('_SDL_CreateGPUShader');
var _SDL_CreateGPUBuffer = Module['_SDL_CreateGPUBuffer'] = makeInvalidEarlyAccess('_SDL_CreateGPUBuffer');
var _SDL_SetGPUBufferName = Module['_SDL_SetGPUBufferName'] = makeInvalidEarlyAccess('_SDL_SetGPUBufferName');
var _SDL_SetGPUTextureName = Module['_SDL_SetGPUTextureName'] = makeInvalidEarlyAccess('_SDL_SetGPUTextureName');
var _SDL_InsertGPUDebugLabel = Module['_SDL_InsertGPUDebugLabel'] = makeInvalidEarlyAccess('_SDL_InsertGPUDebugLabel');
var _SDL_PushGPUDebugGroup = Module['_SDL_PushGPUDebugGroup'] = makeInvalidEarlyAccess('_SDL_PushGPUDebugGroup');
var _SDL_PopGPUDebugGroup = Module['_SDL_PopGPUDebugGroup'] = makeInvalidEarlyAccess('_SDL_PopGPUDebugGroup');
var _SDL_ReleaseGPUSampler = Module['_SDL_ReleaseGPUSampler'] = makeInvalidEarlyAccess('_SDL_ReleaseGPUSampler');
var _SDL_ReleaseGPUBuffer = Module['_SDL_ReleaseGPUBuffer'] = makeInvalidEarlyAccess('_SDL_ReleaseGPUBuffer');
var _SDL_ReleaseGPUShader = Module['_SDL_ReleaseGPUShader'] = makeInvalidEarlyAccess('_SDL_ReleaseGPUShader');
var _SDL_ReleaseGPUComputePipeline = Module['_SDL_ReleaseGPUComputePipeline'] = makeInvalidEarlyAccess('_SDL_ReleaseGPUComputePipeline');
var _SDL_ReleaseGPUGraphicsPipeline = Module['_SDL_ReleaseGPUGraphicsPipeline'] = makeInvalidEarlyAccess('_SDL_ReleaseGPUGraphicsPipeline');
var _SDL_PushGPUVertexUniformData = Module['_SDL_PushGPUVertexUniformData'] = makeInvalidEarlyAccess('_SDL_PushGPUVertexUniformData');
var _SDL_PushGPUComputeUniformData = Module['_SDL_PushGPUComputeUniformData'] = makeInvalidEarlyAccess('_SDL_PushGPUComputeUniformData');
var _SDL_SetGPUScissor = Module['_SDL_SetGPUScissor'] = makeInvalidEarlyAccess('_SDL_SetGPUScissor');
var _SDL_SetGPUBlendConstants = Module['_SDL_SetGPUBlendConstants'] = makeInvalidEarlyAccess('_SDL_SetGPUBlendConstants');
var _SDL_SetGPUStencilReference = Module['_SDL_SetGPUStencilReference'] = makeInvalidEarlyAccess('_SDL_SetGPUStencilReference');
var _SDL_BindGPUVertexBuffers = Module['_SDL_BindGPUVertexBuffers'] = makeInvalidEarlyAccess('_SDL_BindGPUVertexBuffers');
var _SDL_BindGPUIndexBuffer = Module['_SDL_BindGPUIndexBuffer'] = makeInvalidEarlyAccess('_SDL_BindGPUIndexBuffer');
var _SDL_BindGPUVertexSamplers = Module['_SDL_BindGPUVertexSamplers'] = makeInvalidEarlyAccess('_SDL_BindGPUVertexSamplers');
var _SDL_BindGPUVertexStorageTextures = Module['_SDL_BindGPUVertexStorageTextures'] = makeInvalidEarlyAccess('_SDL_BindGPUVertexStorageTextures');
var _SDL_BindGPUVertexStorageBuffers = Module['_SDL_BindGPUVertexStorageBuffers'] = makeInvalidEarlyAccess('_SDL_BindGPUVertexStorageBuffers');
var _SDL_BindGPUFragmentStorageTextures = Module['_SDL_BindGPUFragmentStorageTextures'] = makeInvalidEarlyAccess('_SDL_BindGPUFragmentStorageTextures');
var _SDL_BindGPUFragmentStorageBuffers = Module['_SDL_BindGPUFragmentStorageBuffers'] = makeInvalidEarlyAccess('_SDL_BindGPUFragmentStorageBuffers');
var _SDL_DrawGPUIndexedPrimitives = Module['_SDL_DrawGPUIndexedPrimitives'] = makeInvalidEarlyAccess('_SDL_DrawGPUIndexedPrimitives');
var _SDL_DrawGPUPrimitivesIndirect = Module['_SDL_DrawGPUPrimitivesIndirect'] = makeInvalidEarlyAccess('_SDL_DrawGPUPrimitivesIndirect');
var _SDL_DrawGPUIndexedPrimitivesIndirect = Module['_SDL_DrawGPUIndexedPrimitivesIndirect'] = makeInvalidEarlyAccess('_SDL_DrawGPUIndexedPrimitivesIndirect');
var _SDL_BeginGPUComputePass = Module['_SDL_BeginGPUComputePass'] = makeInvalidEarlyAccess('_SDL_BeginGPUComputePass');
var _SDL_BindGPUComputePipeline = Module['_SDL_BindGPUComputePipeline'] = makeInvalidEarlyAccess('_SDL_BindGPUComputePipeline');
var _SDL_BindGPUComputeSamplers = Module['_SDL_BindGPUComputeSamplers'] = makeInvalidEarlyAccess('_SDL_BindGPUComputeSamplers');
var _SDL_BindGPUComputeStorageTextures = Module['_SDL_BindGPUComputeStorageTextures'] = makeInvalidEarlyAccess('_SDL_BindGPUComputeStorageTextures');
var _SDL_BindGPUComputeStorageBuffers = Module['_SDL_BindGPUComputeStorageBuffers'] = makeInvalidEarlyAccess('_SDL_BindGPUComputeStorageBuffers');
var _SDL_DispatchGPUCompute = Module['_SDL_DispatchGPUCompute'] = makeInvalidEarlyAccess('_SDL_DispatchGPUCompute');
var _SDL_DispatchGPUComputeIndirect = Module['_SDL_DispatchGPUComputeIndirect'] = makeInvalidEarlyAccess('_SDL_DispatchGPUComputeIndirect');
var _SDL_EndGPUComputePass = Module['_SDL_EndGPUComputePass'] = makeInvalidEarlyAccess('_SDL_EndGPUComputePass');
var _SDL_UploadToGPUBuffer = Module['_SDL_UploadToGPUBuffer'] = makeInvalidEarlyAccess('_SDL_UploadToGPUBuffer');
var _SDL_CopyGPUTextureToTexture = Module['_SDL_CopyGPUTextureToTexture'] = makeInvalidEarlyAccess('_SDL_CopyGPUTextureToTexture');
var _SDL_CopyGPUBufferToBuffer = Module['_SDL_CopyGPUBufferToBuffer'] = makeInvalidEarlyAccess('_SDL_CopyGPUBufferToBuffer');
var _SDL_DownloadFromGPUTexture = Module['_SDL_DownloadFromGPUTexture'] = makeInvalidEarlyAccess('_SDL_DownloadFromGPUTexture');
var _SDL_DownloadFromGPUBuffer = Module['_SDL_DownloadFromGPUBuffer'] = makeInvalidEarlyAccess('_SDL_DownloadFromGPUBuffer');
var _SDL_GenerateMipmapsForGPUTexture = Module['_SDL_GenerateMipmapsForGPUTexture'] = makeInvalidEarlyAccess('_SDL_GenerateMipmapsForGPUTexture');
var _SDL_BlitGPUTexture = Module['_SDL_BlitGPUTexture'] = makeInvalidEarlyAccess('_SDL_BlitGPUTexture');
var _SDL_WindowSupportsGPUSwapchainComposition = Module['_SDL_WindowSupportsGPUSwapchainComposition'] = makeInvalidEarlyAccess('_SDL_WindowSupportsGPUSwapchainComposition');
var _SDL_WindowSupportsGPUPresentMode = Module['_SDL_WindowSupportsGPUPresentMode'] = makeInvalidEarlyAccess('_SDL_WindowSupportsGPUPresentMode');
var _SDL_ClaimWindowForGPUDevice = Module['_SDL_ClaimWindowForGPUDevice'] = makeInvalidEarlyAccess('_SDL_ClaimWindowForGPUDevice');
var _SDL_ReleaseWindowFromGPUDevice = Module['_SDL_ReleaseWindowFromGPUDevice'] = makeInvalidEarlyAccess('_SDL_ReleaseWindowFromGPUDevice');
var _SDL_SetGPUSwapchainParameters = Module['_SDL_SetGPUSwapchainParameters'] = makeInvalidEarlyAccess('_SDL_SetGPUSwapchainParameters');
var _SDL_SetGPUAllowedFramesInFlight = Module['_SDL_SetGPUAllowedFramesInFlight'] = makeInvalidEarlyAccess('_SDL_SetGPUAllowedFramesInFlight');
var _SDL_GetGPUSwapchainTextureFormat = Module['_SDL_GetGPUSwapchainTextureFormat'] = makeInvalidEarlyAccess('_SDL_GetGPUSwapchainTextureFormat');
var _SDL_AcquireGPUSwapchainTexture = Module['_SDL_AcquireGPUSwapchainTexture'] = makeInvalidEarlyAccess('_SDL_AcquireGPUSwapchainTexture');
var _SDL_WaitForGPUSwapchain = Module['_SDL_WaitForGPUSwapchain'] = makeInvalidEarlyAccess('_SDL_WaitForGPUSwapchain');
var _SDL_WaitAndAcquireGPUSwapchainTexture = Module['_SDL_WaitAndAcquireGPUSwapchainTexture'] = makeInvalidEarlyAccess('_SDL_WaitAndAcquireGPUSwapchainTexture');
var _SDL_SubmitGPUCommandBufferAndAcquireFence = Module['_SDL_SubmitGPUCommandBufferAndAcquireFence'] = makeInvalidEarlyAccess('_SDL_SubmitGPUCommandBufferAndAcquireFence');
var _SDL_CancelGPUCommandBuffer = Module['_SDL_CancelGPUCommandBuffer'] = makeInvalidEarlyAccess('_SDL_CancelGPUCommandBuffer');
var _SDL_WaitForGPUIdle = Module['_SDL_WaitForGPUIdle'] = makeInvalidEarlyAccess('_SDL_WaitForGPUIdle');
var _SDL_WaitForGPUFences = Module['_SDL_WaitForGPUFences'] = makeInvalidEarlyAccess('_SDL_WaitForGPUFences');
var _SDL_QueryGPUFence = Module['_SDL_QueryGPUFence'] = makeInvalidEarlyAccess('_SDL_QueryGPUFence');
var _SDL_ReleaseGPUFence = Module['_SDL_ReleaseGPUFence'] = makeInvalidEarlyAccess('_SDL_ReleaseGPUFence');
var _SDL_CalculateGPUTextureFormatSize = Module['_SDL_CalculateGPUTextureFormatSize'] = makeInvalidEarlyAccess('_SDL_CalculateGPUTextureFormatSize');
var _SDL_GetHaptics = Module['_SDL_GetHaptics'] = makeInvalidEarlyAccess('_SDL_GetHaptics');
var _SDL_GetHapticNameForID = Module['_SDL_GetHapticNameForID'] = makeInvalidEarlyAccess('_SDL_GetHapticNameForID');
var _SDL_OpenHaptic = Module['_SDL_OpenHaptic'] = makeInvalidEarlyAccess('_SDL_OpenHaptic');
var _SDL_SetHapticGain = Module['_SDL_SetHapticGain'] = makeInvalidEarlyAccess('_SDL_SetHapticGain');
var _SDL_SetHapticAutocenter = Module['_SDL_SetHapticAutocenter'] = makeInvalidEarlyAccess('_SDL_SetHapticAutocenter');
var _SDL_GetHapticFromID = Module['_SDL_GetHapticFromID'] = makeInvalidEarlyAccess('_SDL_GetHapticFromID');
var _SDL_GetHapticID = Module['_SDL_GetHapticID'] = makeInvalidEarlyAccess('_SDL_GetHapticID');
var _SDL_GetHapticName = Module['_SDL_GetHapticName'] = makeInvalidEarlyAccess('_SDL_GetHapticName');
var _SDL_IsMouseHaptic = Module['_SDL_IsMouseHaptic'] = makeInvalidEarlyAccess('_SDL_IsMouseHaptic');
var _SDL_OpenHapticFromMouse = Module['_SDL_OpenHapticFromMouse'] = makeInvalidEarlyAccess('_SDL_OpenHapticFromMouse');
var _SDL_IsJoystickHaptic = Module['_SDL_IsJoystickHaptic'] = makeInvalidEarlyAccess('_SDL_IsJoystickHaptic');
var _SDL_LockJoysticks = Module['_SDL_LockJoysticks'] = makeInvalidEarlyAccess('_SDL_LockJoysticks');
var _SDL_GetJoystickID = Module['_SDL_GetJoystickID'] = makeInvalidEarlyAccess('_SDL_GetJoystickID');
var _SDL_IsGamepad = Module['_SDL_IsGamepad'] = makeInvalidEarlyAccess('_SDL_IsGamepad');
var _SDL_UnlockJoysticks = Module['_SDL_UnlockJoysticks'] = makeInvalidEarlyAccess('_SDL_UnlockJoysticks');
var _SDL_OpenHapticFromJoystick = Module['_SDL_OpenHapticFromJoystick'] = makeInvalidEarlyAccess('_SDL_OpenHapticFromJoystick');
var _SDL_GetJoystickVendor = Module['_SDL_GetJoystickVendor'] = makeInvalidEarlyAccess('_SDL_GetJoystickVendor');
var _SDL_GetJoystickProduct = Module['_SDL_GetJoystickProduct'] = makeInvalidEarlyAccess('_SDL_GetJoystickProduct');
var _SDL_GetNumJoystickAxes = Module['_SDL_GetNumJoystickAxes'] = makeInvalidEarlyAccess('_SDL_GetNumJoystickAxes');
var _SDL_CloseHaptic = Module['_SDL_CloseHaptic'] = makeInvalidEarlyAccess('_SDL_CloseHaptic');
var _SDL_DestroyHapticEffect = Module['_SDL_DestroyHapticEffect'] = makeInvalidEarlyAccess('_SDL_DestroyHapticEffect');
var _SDL_GetMaxHapticEffects = Module['_SDL_GetMaxHapticEffects'] = makeInvalidEarlyAccess('_SDL_GetMaxHapticEffects');
var _SDL_GetMaxHapticEffectsPlaying = Module['_SDL_GetMaxHapticEffectsPlaying'] = makeInvalidEarlyAccess('_SDL_GetMaxHapticEffectsPlaying');
var _SDL_GetHapticFeatures = Module['_SDL_GetHapticFeatures'] = makeInvalidEarlyAccess('_SDL_GetHapticFeatures');
var _SDL_GetNumHapticAxes = Module['_SDL_GetNumHapticAxes'] = makeInvalidEarlyAccess('_SDL_GetNumHapticAxes');
var _SDL_HapticEffectSupported = Module['_SDL_HapticEffectSupported'] = makeInvalidEarlyAccess('_SDL_HapticEffectSupported');
var _SDL_CreateHapticEffect = Module['_SDL_CreateHapticEffect'] = makeInvalidEarlyAccess('_SDL_CreateHapticEffect');
var _SDL_UpdateHapticEffect = Module['_SDL_UpdateHapticEffect'] = makeInvalidEarlyAccess('_SDL_UpdateHapticEffect');
var _SDL_RunHapticEffect = Module['_SDL_RunHapticEffect'] = makeInvalidEarlyAccess('_SDL_RunHapticEffect');
var _SDL_StopHapticEffect = Module['_SDL_StopHapticEffect'] = makeInvalidEarlyAccess('_SDL_StopHapticEffect');
var _SDL_GetHapticEffectStatus = Module['_SDL_GetHapticEffectStatus'] = makeInvalidEarlyAccess('_SDL_GetHapticEffectStatus');
var _SDL_PauseHaptic = Module['_SDL_PauseHaptic'] = makeInvalidEarlyAccess('_SDL_PauseHaptic');
var _SDL_ResumeHaptic = Module['_SDL_ResumeHaptic'] = makeInvalidEarlyAccess('_SDL_ResumeHaptic');
var _SDL_StopHapticEffects = Module['_SDL_StopHapticEffects'] = makeInvalidEarlyAccess('_SDL_StopHapticEffects');
var _SDL_HapticRumbleSupported = Module['_SDL_HapticRumbleSupported'] = makeInvalidEarlyAccess('_SDL_HapticRumbleSupported');
var _SDL_InitHapticRumble = Module['_SDL_InitHapticRumble'] = makeInvalidEarlyAccess('_SDL_InitHapticRumble');
var _SDL_PlayHapticRumble = Module['_SDL_PlayHapticRumble'] = makeInvalidEarlyAccess('_SDL_PlayHapticRumble');
var _SDL_StopHapticRumble = Module['_SDL_StopHapticRumble'] = makeInvalidEarlyAccess('_SDL_StopHapticRumble');
var _SDL_strcasestr = Module['_SDL_strcasestr'] = makeInvalidEarlyAccess('_SDL_strcasestr');
var _SDL_hid_init = Module['_SDL_hid_init'] = makeInvalidEarlyAccess('_SDL_hid_init');
var _SDL_hid_exit = Module['_SDL_hid_exit'] = makeInvalidEarlyAccess('_SDL_hid_exit');
var _SDL_hid_device_change_count = Module['_SDL_hid_device_change_count'] = makeInvalidEarlyAccess('_SDL_hid_device_change_count');
var _SDL_hid_enumerate = Module['_SDL_hid_enumerate'] = makeInvalidEarlyAccess('_SDL_hid_enumerate');
var _SDL_hid_free_enumeration = Module['_SDL_hid_free_enumeration'] = makeInvalidEarlyAccess('_SDL_hid_free_enumeration');
var _SDL_hid_open = Module['_SDL_hid_open'] = makeInvalidEarlyAccess('_SDL_hid_open');
var _SDL_hid_open_path = Module['_SDL_hid_open_path'] = makeInvalidEarlyAccess('_SDL_hid_open_path');
var _SDL_hid_write = Module['_SDL_hid_write'] = makeInvalidEarlyAccess('_SDL_hid_write');
var _SDL_hid_read_timeout = Module['_SDL_hid_read_timeout'] = makeInvalidEarlyAccess('_SDL_hid_read_timeout');
var _SDL_hid_read = Module['_SDL_hid_read'] = makeInvalidEarlyAccess('_SDL_hid_read');
var _SDL_hid_set_nonblocking = Module['_SDL_hid_set_nonblocking'] = makeInvalidEarlyAccess('_SDL_hid_set_nonblocking');
var _SDL_hid_send_feature_report = Module['_SDL_hid_send_feature_report'] = makeInvalidEarlyAccess('_SDL_hid_send_feature_report');
var _SDL_hid_get_feature_report = Module['_SDL_hid_get_feature_report'] = makeInvalidEarlyAccess('_SDL_hid_get_feature_report');
var _SDL_hid_get_input_report = Module['_SDL_hid_get_input_report'] = makeInvalidEarlyAccess('_SDL_hid_get_input_report');
var _SDL_hid_close = Module['_SDL_hid_close'] = makeInvalidEarlyAccess('_SDL_hid_close');
var _SDL_hid_get_manufacturer_string = Module['_SDL_hid_get_manufacturer_string'] = makeInvalidEarlyAccess('_SDL_hid_get_manufacturer_string');
var _SDL_hid_get_product_string = Module['_SDL_hid_get_product_string'] = makeInvalidEarlyAccess('_SDL_hid_get_product_string');
var _SDL_hid_get_serial_number_string = Module['_SDL_hid_get_serial_number_string'] = makeInvalidEarlyAccess('_SDL_hid_get_serial_number_string');
var _SDL_hid_get_indexed_string = Module['_SDL_hid_get_indexed_string'] = makeInvalidEarlyAccess('_SDL_hid_get_indexed_string');
var _SDL_hid_get_device_info = Module['_SDL_hid_get_device_info'] = makeInvalidEarlyAccess('_SDL_hid_get_device_info');
var _SDL_wcsdup = Module['_SDL_wcsdup'] = makeInvalidEarlyAccess('_SDL_wcsdup');
var _SDL_hid_get_report_descriptor = Module['_SDL_hid_get_report_descriptor'] = makeInvalidEarlyAccess('_SDL_hid_get_report_descriptor');
var _SDL_hid_ble_scan = Module['_SDL_hid_ble_scan'] = makeInvalidEarlyAccess('_SDL_hid_ble_scan');
var _SDL_AsyncIOFromFile = Module['_SDL_AsyncIOFromFile'] = makeInvalidEarlyAccess('_SDL_AsyncIOFromFile');
var _SDL_GetAsyncIOSize = Module['_SDL_GetAsyncIOSize'] = makeInvalidEarlyAccess('_SDL_GetAsyncIOSize');
var _SDL_ReadAsyncIO = Module['_SDL_ReadAsyncIO'] = makeInvalidEarlyAccess('_SDL_ReadAsyncIO');
var _SDL_WriteAsyncIO = Module['_SDL_WriteAsyncIO'] = makeInvalidEarlyAccess('_SDL_WriteAsyncIO');
var _SDL_CloseAsyncIO = Module['_SDL_CloseAsyncIO'] = makeInvalidEarlyAccess('_SDL_CloseAsyncIO');
var _SDL_CreateAsyncIOQueue = Module['_SDL_CreateAsyncIOQueue'] = makeInvalidEarlyAccess('_SDL_CreateAsyncIOQueue');
var _SDL_GetAsyncIOResult = Module['_SDL_GetAsyncIOResult'] = makeInvalidEarlyAccess('_SDL_GetAsyncIOResult');
var _SDL_WaitAsyncIOResult = Module['_SDL_WaitAsyncIOResult'] = makeInvalidEarlyAccess('_SDL_WaitAsyncIOResult');
var _SDL_SignalAsyncIOQueue = Module['_SDL_SignalAsyncIOQueue'] = makeInvalidEarlyAccess('_SDL_SignalAsyncIOQueue');
var _SDL_DestroyAsyncIOQueue = Module['_SDL_DestroyAsyncIOQueue'] = makeInvalidEarlyAccess('_SDL_DestroyAsyncIOQueue');
var _SDL_LoadFileAsync = Module['_SDL_LoadFileAsync'] = makeInvalidEarlyAccess('_SDL_LoadFileAsync');
var _SDL_OpenIO = Module['_SDL_OpenIO'] = makeInvalidEarlyAccess('_SDL_OpenIO');
var _fileno = Module['_fileno'] = makeInvalidEarlyAccess('_fileno');
var _fflush = Module['_fflush'] = makeInvalidEarlyAccess('_fflush');
var _SDL_IOFromMem = Module['_SDL_IOFromMem'] = makeInvalidEarlyAccess('_SDL_IOFromMem');
var _SDL_LoadFile = Module['_SDL_LoadFile'] = makeInvalidEarlyAccess('_SDL_LoadFile');
var _SDL_SaveFile_IO = Module['_SDL_SaveFile_IO'] = makeInvalidEarlyAccess('_SDL_SaveFile_IO');
var _SDL_SaveFile = Module['_SDL_SaveFile'] = makeInvalidEarlyAccess('_SDL_SaveFile');
var _SDL_IOprintf = Module['_SDL_IOprintf'] = makeInvalidEarlyAccess('_SDL_IOprintf');
var _SDL_vasprintf = Module['_SDL_vasprintf'] = makeInvalidEarlyAccess('_SDL_vasprintf');
var _SDL_IOvprintf = Module['_SDL_IOvprintf'] = makeInvalidEarlyAccess('_SDL_IOvprintf');
var _SDL_FlushIO = Module['_SDL_FlushIO'] = makeInvalidEarlyAccess('_SDL_FlushIO');
var _SDL_ReadS8 = Module['_SDL_ReadS8'] = makeInvalidEarlyAccess('_SDL_ReadS8');
var _SDL_ReadS16LE = Module['_SDL_ReadS16LE'] = makeInvalidEarlyAccess('_SDL_ReadS16LE');
var _SDL_ReadU64LE = Module['_SDL_ReadU64LE'] = makeInvalidEarlyAccess('_SDL_ReadU64LE');
var _SDL_ReadS64LE = Module['_SDL_ReadS64LE'] = makeInvalidEarlyAccess('_SDL_ReadS64LE');
var _SDL_ReadU64BE = Module['_SDL_ReadU64BE'] = makeInvalidEarlyAccess('_SDL_ReadU64BE');
var _SDL_ReadS64BE = Module['_SDL_ReadS64BE'] = makeInvalidEarlyAccess('_SDL_ReadS64BE');
var _SDL_WriteU8 = Module['_SDL_WriteU8'] = makeInvalidEarlyAccess('_SDL_WriteU8');
var _SDL_WriteS8 = Module['_SDL_WriteS8'] = makeInvalidEarlyAccess('_SDL_WriteS8');
var _SDL_WriteU16LE = Module['_SDL_WriteU16LE'] = makeInvalidEarlyAccess('_SDL_WriteU16LE');
var _SDL_WriteS16LE = Module['_SDL_WriteS16LE'] = makeInvalidEarlyAccess('_SDL_WriteS16LE');
var _SDL_WriteU16BE = Module['_SDL_WriteU16BE'] = makeInvalidEarlyAccess('_SDL_WriteU16BE');
var _SDL_WriteS16BE = Module['_SDL_WriteS16BE'] = makeInvalidEarlyAccess('_SDL_WriteS16BE');
var _SDL_WriteU32LE = Module['_SDL_WriteU32LE'] = makeInvalidEarlyAccess('_SDL_WriteU32LE');
var _SDL_WriteS32LE = Module['_SDL_WriteS32LE'] = makeInvalidEarlyAccess('_SDL_WriteS32LE');
var _SDL_WriteU32BE = Module['_SDL_WriteU32BE'] = makeInvalidEarlyAccess('_SDL_WriteU32BE');
var _SDL_WriteS32BE = Module['_SDL_WriteS32BE'] = makeInvalidEarlyAccess('_SDL_WriteS32BE');
var _SDL_WriteU64LE = Module['_SDL_WriteU64LE'] = makeInvalidEarlyAccess('_SDL_WriteU64LE');
var _SDL_WriteS64LE = Module['_SDL_WriteS64LE'] = makeInvalidEarlyAccess('_SDL_WriteS64LE');
var _SDL_WriteU64BE = Module['_SDL_WriteU64BE'] = makeInvalidEarlyAccess('_SDL_WriteU64BE');
var _SDL_WriteS64BE = Module['_SDL_WriteS64BE'] = makeInvalidEarlyAccess('_SDL_WriteS64BE');
var _SDL_SignalCondition = Module['_SDL_SignalCondition'] = makeInvalidEarlyAccess('_SDL_SignalCondition');
var _SDL_WaitConditionTimeout = Module['_SDL_WaitConditionTimeout'] = makeInvalidEarlyAccess('_SDL_WaitConditionTimeout');
var _SDL_GetGamepadButton = Module['_SDL_GetGamepadButton'] = makeInvalidEarlyAccess('_SDL_GetGamepadButton');
var _SDL_GetGamepadAxis = Module['_SDL_GetGamepadAxis'] = makeInvalidEarlyAccess('_SDL_GetGamepadAxis');
var _SDL_GetGamepadTypeFromString = Module['_SDL_GetGamepadTypeFromString'] = makeInvalidEarlyAccess('_SDL_GetGamepadTypeFromString');
var _SDL_GetGamepadStringForType = Module['_SDL_GetGamepadStringForType'] = makeInvalidEarlyAccess('_SDL_GetGamepadStringForType');
var _SDL_GetGamepadAxisFromString = Module['_SDL_GetGamepadAxisFromString'] = makeInvalidEarlyAccess('_SDL_GetGamepadAxisFromString');
var _SDL_GetGamepadStringForAxis = Module['_SDL_GetGamepadStringForAxis'] = makeInvalidEarlyAccess('_SDL_GetGamepadStringForAxis');
var _SDL_GetGamepadButtonFromString = Module['_SDL_GetGamepadButtonFromString'] = makeInvalidEarlyAccess('_SDL_GetGamepadButtonFromString');
var _SDL_GetGamepadStringForButton = Module['_SDL_GetGamepadStringForButton'] = makeInvalidEarlyAccess('_SDL_GetGamepadStringForButton');
var _SDL_AddGamepadMappingsFromIO = Module['_SDL_AddGamepadMappingsFromIO'] = makeInvalidEarlyAccess('_SDL_AddGamepadMappingsFromIO');
var _SDL_GetJoysticks = Module['_SDL_GetJoysticks'] = makeInvalidEarlyAccess('_SDL_GetJoysticks');
var _SDL_GetJoystickNameForID = Module['_SDL_GetJoystickNameForID'] = makeInvalidEarlyAccess('_SDL_GetJoystickNameForID');
var _SDL_GetJoystickGUIDForID = Module['_SDL_GetJoystickGUIDForID'] = makeInvalidEarlyAccess('_SDL_GetJoystickGUIDForID');
var _SDL_AddGamepadMapping = Module['_SDL_AddGamepadMapping'] = makeInvalidEarlyAccess('_SDL_AddGamepadMapping');
var _SDL_AddGamepadMappingsFromFile = Module['_SDL_AddGamepadMappingsFromFile'] = makeInvalidEarlyAccess('_SDL_AddGamepadMappingsFromFile');
var _SDL_ReloadGamepadMappings = Module['_SDL_ReloadGamepadMappings'] = makeInvalidEarlyAccess('_SDL_ReloadGamepadMappings');
var _SDL_GetGamepadMappings = Module['_SDL_GetGamepadMappings'] = makeInvalidEarlyAccess('_SDL_GetGamepadMappings');
var _SDL_strlcat = Module['_SDL_strlcat'] = makeInvalidEarlyAccess('_SDL_strlcat');
var _SDL_GetGamepadMappingForGUID = Module['_SDL_GetGamepadMappingForGUID'] = makeInvalidEarlyAccess('_SDL_GetGamepadMappingForGUID');
var _SDL_GetJoystickGUIDInfo = Module['_SDL_GetJoystickGUIDInfo'] = makeInvalidEarlyAccess('_SDL_GetJoystickGUIDInfo');
var _SDL_GetGamepadMapping = Module['_SDL_GetGamepadMapping'] = makeInvalidEarlyAccess('_SDL_GetGamepadMapping');
var _SDL_SetGamepadMapping = Module['_SDL_SetGamepadMapping'] = makeInvalidEarlyAccess('_SDL_SetGamepadMapping');
var _SDL_HasGamepad = Module['_SDL_HasGamepad'] = makeInvalidEarlyAccess('_SDL_HasGamepad');
var _SDL_GetGamepads = Module['_SDL_GetGamepads'] = makeInvalidEarlyAccess('_SDL_GetGamepads');
var _SDL_GetGamepadNameForID = Module['_SDL_GetGamepadNameForID'] = makeInvalidEarlyAccess('_SDL_GetGamepadNameForID');
var _SDL_GetGamepadPathForID = Module['_SDL_GetGamepadPathForID'] = makeInvalidEarlyAccess('_SDL_GetGamepadPathForID');
var _SDL_GetJoystickPathForID = Module['_SDL_GetJoystickPathForID'] = makeInvalidEarlyAccess('_SDL_GetJoystickPathForID');
var _SDL_GetGamepadPlayerIndexForID = Module['_SDL_GetGamepadPlayerIndexForID'] = makeInvalidEarlyAccess('_SDL_GetGamepadPlayerIndexForID');
var _SDL_GetJoystickPlayerIndexForID = Module['_SDL_GetJoystickPlayerIndexForID'] = makeInvalidEarlyAccess('_SDL_GetJoystickPlayerIndexForID');
var _SDL_GetGamepadGUIDForID = Module['_SDL_GetGamepadGUIDForID'] = makeInvalidEarlyAccess('_SDL_GetGamepadGUIDForID');
var _SDL_GetGamepadVendorForID = Module['_SDL_GetGamepadVendorForID'] = makeInvalidEarlyAccess('_SDL_GetGamepadVendorForID');
var _SDL_GetJoystickVendorForID = Module['_SDL_GetJoystickVendorForID'] = makeInvalidEarlyAccess('_SDL_GetJoystickVendorForID');
var _SDL_GetGamepadProductForID = Module['_SDL_GetGamepadProductForID'] = makeInvalidEarlyAccess('_SDL_GetGamepadProductForID');
var _SDL_GetJoystickProductForID = Module['_SDL_GetJoystickProductForID'] = makeInvalidEarlyAccess('_SDL_GetJoystickProductForID');
var _SDL_GetGamepadProductVersionForID = Module['_SDL_GetGamepadProductVersionForID'] = makeInvalidEarlyAccess('_SDL_GetGamepadProductVersionForID');
var _SDL_GetJoystickProductVersionForID = Module['_SDL_GetJoystickProductVersionForID'] = makeInvalidEarlyAccess('_SDL_GetJoystickProductVersionForID');
var _SDL_GetGamepadTypeForID = Module['_SDL_GetGamepadTypeForID'] = makeInvalidEarlyAccess('_SDL_GetGamepadTypeForID');
var _SDL_GetRealGamepadTypeForID = Module['_SDL_GetRealGamepadTypeForID'] = makeInvalidEarlyAccess('_SDL_GetRealGamepadTypeForID');
var _SDL_GetGamepadMappingForID = Module['_SDL_GetGamepadMappingForID'] = makeInvalidEarlyAccess('_SDL_GetGamepadMappingForID');
var _SDL_OpenGamepad = Module['_SDL_OpenGamepad'] = makeInvalidEarlyAccess('_SDL_OpenGamepad');
var _SDL_OpenJoystick = Module['_SDL_OpenJoystick'] = makeInvalidEarlyAccess('_SDL_OpenJoystick');
var _SDL_CloseJoystick = Module['_SDL_CloseJoystick'] = makeInvalidEarlyAccess('_SDL_CloseJoystick');
var _SDL_UpdateGamepads = Module['_SDL_UpdateGamepads'] = makeInvalidEarlyAccess('_SDL_UpdateGamepads');
var _SDL_GamepadHasAxis = Module['_SDL_GamepadHasAxis'] = makeInvalidEarlyAccess('_SDL_GamepadHasAxis');
var _SDL_GetJoystickAxis = Module['_SDL_GetJoystickAxis'] = makeInvalidEarlyAccess('_SDL_GetJoystickAxis');
var _SDL_GetJoystickButton = Module['_SDL_GetJoystickButton'] = makeInvalidEarlyAccess('_SDL_GetJoystickButton');
var _SDL_GetJoystickHat = Module['_SDL_GetJoystickHat'] = makeInvalidEarlyAccess('_SDL_GetJoystickHat');
var _SDL_GamepadHasButton = Module['_SDL_GamepadHasButton'] = makeInvalidEarlyAccess('_SDL_GamepadHasButton');
var _SDL_GetGamepadButtonLabelForType = Module['_SDL_GetGamepadButtonLabelForType'] = makeInvalidEarlyAccess('_SDL_GetGamepadButtonLabelForType');
var _SDL_GetGamepadButtonLabel = Module['_SDL_GetGamepadButtonLabel'] = makeInvalidEarlyAccess('_SDL_GetGamepadButtonLabel');
var _SDL_GetNumGamepadTouchpads = Module['_SDL_GetNumGamepadTouchpads'] = makeInvalidEarlyAccess('_SDL_GetNumGamepadTouchpads');
var _SDL_GetGamepadJoystick = Module['_SDL_GetGamepadJoystick'] = makeInvalidEarlyAccess('_SDL_GetGamepadJoystick');
var _SDL_GetNumGamepadTouchpadFingers = Module['_SDL_GetNumGamepadTouchpadFingers'] = makeInvalidEarlyAccess('_SDL_GetNumGamepadTouchpadFingers');
var _SDL_GetGamepadTouchpadFinger = Module['_SDL_GetGamepadTouchpadFinger'] = makeInvalidEarlyAccess('_SDL_GetGamepadTouchpadFinger');
var _SDL_GamepadHasSensor = Module['_SDL_GamepadHasSensor'] = makeInvalidEarlyAccess('_SDL_GamepadHasSensor');
var _SDL_SetGamepadSensorEnabled = Module['_SDL_SetGamepadSensorEnabled'] = makeInvalidEarlyAccess('_SDL_SetGamepadSensorEnabled');
var _SDL_OpenSensor = Module['_SDL_OpenSensor'] = makeInvalidEarlyAccess('_SDL_OpenSensor');
var _SDL_CloseSensor = Module['_SDL_CloseSensor'] = makeInvalidEarlyAccess('_SDL_CloseSensor');
var _SDL_GamepadSensorEnabled = Module['_SDL_GamepadSensorEnabled'] = makeInvalidEarlyAccess('_SDL_GamepadSensorEnabled');
var _SDL_GetGamepadSensorDataRate = Module['_SDL_GetGamepadSensorDataRate'] = makeInvalidEarlyAccess('_SDL_GetGamepadSensorDataRate');
var _SDL_GetGamepadSensorData = Module['_SDL_GetGamepadSensorData'] = makeInvalidEarlyAccess('_SDL_GetGamepadSensorData');
var _SDL_GetGamepadID = Module['_SDL_GetGamepadID'] = makeInvalidEarlyAccess('_SDL_GetGamepadID');
var _SDL_GetGamepadProperties = Module['_SDL_GetGamepadProperties'] = makeInvalidEarlyAccess('_SDL_GetGamepadProperties');
var _SDL_GetJoystickProperties = Module['_SDL_GetJoystickProperties'] = makeInvalidEarlyAccess('_SDL_GetJoystickProperties');
var _SDL_GetGamepadName = Module['_SDL_GetGamepadName'] = makeInvalidEarlyAccess('_SDL_GetGamepadName');
var _SDL_GetJoystickName = Module['_SDL_GetJoystickName'] = makeInvalidEarlyAccess('_SDL_GetJoystickName');
var _SDL_GetGamepadPath = Module['_SDL_GetGamepadPath'] = makeInvalidEarlyAccess('_SDL_GetGamepadPath');
var _SDL_GetJoystickPath = Module['_SDL_GetJoystickPath'] = makeInvalidEarlyAccess('_SDL_GetJoystickPath');
var _SDL_GetGamepadType = Module['_SDL_GetGamepadType'] = makeInvalidEarlyAccess('_SDL_GetGamepadType');
var _SDL_GetRealGamepadType = Module['_SDL_GetRealGamepadType'] = makeInvalidEarlyAccess('_SDL_GetRealGamepadType');
var _SDL_GetJoystickGUID = Module['_SDL_GetJoystickGUID'] = makeInvalidEarlyAccess('_SDL_GetJoystickGUID');
var _SDL_GetGamepadPlayerIndex = Module['_SDL_GetGamepadPlayerIndex'] = makeInvalidEarlyAccess('_SDL_GetGamepadPlayerIndex');
var _SDL_GetJoystickPlayerIndex = Module['_SDL_GetJoystickPlayerIndex'] = makeInvalidEarlyAccess('_SDL_GetJoystickPlayerIndex');
var _SDL_SetGamepadPlayerIndex = Module['_SDL_SetGamepadPlayerIndex'] = makeInvalidEarlyAccess('_SDL_SetGamepadPlayerIndex');
var _SDL_SetJoystickPlayerIndex = Module['_SDL_SetJoystickPlayerIndex'] = makeInvalidEarlyAccess('_SDL_SetJoystickPlayerIndex');
var _SDL_GetGamepadVendor = Module['_SDL_GetGamepadVendor'] = makeInvalidEarlyAccess('_SDL_GetGamepadVendor');
var _SDL_GetGamepadProduct = Module['_SDL_GetGamepadProduct'] = makeInvalidEarlyAccess('_SDL_GetGamepadProduct');
var _SDL_GetGamepadProductVersion = Module['_SDL_GetGamepadProductVersion'] = makeInvalidEarlyAccess('_SDL_GetGamepadProductVersion');
var _SDL_GetJoystickProductVersion = Module['_SDL_GetJoystickProductVersion'] = makeInvalidEarlyAccess('_SDL_GetJoystickProductVersion');
var _SDL_GetGamepadFirmwareVersion = Module['_SDL_GetGamepadFirmwareVersion'] = makeInvalidEarlyAccess('_SDL_GetGamepadFirmwareVersion');
var _SDL_GetJoystickFirmwareVersion = Module['_SDL_GetJoystickFirmwareVersion'] = makeInvalidEarlyAccess('_SDL_GetJoystickFirmwareVersion');
var _SDL_GetGamepadSerial = Module['_SDL_GetGamepadSerial'] = makeInvalidEarlyAccess('_SDL_GetGamepadSerial');
var _SDL_GetJoystickSerial = Module['_SDL_GetJoystickSerial'] = makeInvalidEarlyAccess('_SDL_GetJoystickSerial');
var _SDL_GetGamepadSteamHandle = Module['_SDL_GetGamepadSteamHandle'] = makeInvalidEarlyAccess('_SDL_GetGamepadSteamHandle');
var _SDL_GetGamepadConnectionState = Module['_SDL_GetGamepadConnectionState'] = makeInvalidEarlyAccess('_SDL_GetGamepadConnectionState');
var _SDL_GetJoystickConnectionState = Module['_SDL_GetJoystickConnectionState'] = makeInvalidEarlyAccess('_SDL_GetJoystickConnectionState');
var _SDL_GetGamepadPowerInfo = Module['_SDL_GetGamepadPowerInfo'] = makeInvalidEarlyAccess('_SDL_GetGamepadPowerInfo');
var _SDL_GetJoystickPowerInfo = Module['_SDL_GetJoystickPowerInfo'] = makeInvalidEarlyAccess('_SDL_GetJoystickPowerInfo');
var _SDL_GamepadConnected = Module['_SDL_GamepadConnected'] = makeInvalidEarlyAccess('_SDL_GamepadConnected');
var _SDL_JoystickConnected = Module['_SDL_JoystickConnected'] = makeInvalidEarlyAccess('_SDL_JoystickConnected');
var _SDL_GetGamepadFromID = Module['_SDL_GetGamepadFromID'] = makeInvalidEarlyAccess('_SDL_GetGamepadFromID');
var _SDL_GetGamepadFromPlayerIndex = Module['_SDL_GetGamepadFromPlayerIndex'] = makeInvalidEarlyAccess('_SDL_GetGamepadFromPlayerIndex');
var _SDL_GetJoystickFromPlayerIndex = Module['_SDL_GetJoystickFromPlayerIndex'] = makeInvalidEarlyAccess('_SDL_GetJoystickFromPlayerIndex');
var _SDL_GetGamepadBindings = Module['_SDL_GetGamepadBindings'] = makeInvalidEarlyAccess('_SDL_GetGamepadBindings');
var _SDL_RumbleGamepad = Module['_SDL_RumbleGamepad'] = makeInvalidEarlyAccess('_SDL_RumbleGamepad');
var _SDL_RumbleJoystick = Module['_SDL_RumbleJoystick'] = makeInvalidEarlyAccess('_SDL_RumbleJoystick');
var _SDL_RumbleGamepadTriggers = Module['_SDL_RumbleGamepadTriggers'] = makeInvalidEarlyAccess('_SDL_RumbleGamepadTriggers');
var _SDL_RumbleJoystickTriggers = Module['_SDL_RumbleJoystickTriggers'] = makeInvalidEarlyAccess('_SDL_RumbleJoystickTriggers');
var _SDL_SetGamepadLED = Module['_SDL_SetGamepadLED'] = makeInvalidEarlyAccess('_SDL_SetGamepadLED');
var _SDL_SetJoystickLED = Module['_SDL_SetJoystickLED'] = makeInvalidEarlyAccess('_SDL_SetJoystickLED');
var _SDL_SendGamepadEffect = Module['_SDL_SendGamepadEffect'] = makeInvalidEarlyAccess('_SDL_SendGamepadEffect');
var _SDL_SendJoystickEffect = Module['_SDL_SendJoystickEffect'] = makeInvalidEarlyAccess('_SDL_SendJoystickEffect');
var _SDL_CloseGamepad = Module['_SDL_CloseGamepad'] = makeInvalidEarlyAccess('_SDL_CloseGamepad');
var _SDL_SetGamepadEventsEnabled = Module['_SDL_SetGamepadEventsEnabled'] = makeInvalidEarlyAccess('_SDL_SetGamepadEventsEnabled');
var _SDL_GamepadEventsEnabled = Module['_SDL_GamepadEventsEnabled'] = makeInvalidEarlyAccess('_SDL_GamepadEventsEnabled');
var _SDL_GetGamepadAppleSFSymbolsNameForButton = Module['_SDL_GetGamepadAppleSFSymbolsNameForButton'] = makeInvalidEarlyAccess('_SDL_GetGamepadAppleSFSymbolsNameForButton');
var _SDL_GetGamepadAppleSFSymbolsNameForAxis = Module['_SDL_GetGamepadAppleSFSymbolsNameForAxis'] = makeInvalidEarlyAccess('_SDL_GetGamepadAppleSFSymbolsNameForAxis');
var _SDL_HasJoystick = Module['_SDL_HasJoystick'] = makeInvalidEarlyAccess('_SDL_HasJoystick');
var _SDL_GetSensors = Module['_SDL_GetSensors'] = makeInvalidEarlyAccess('_SDL_GetSensors');
var _SDL_GetSensorTypeForID = Module['_SDL_GetSensorTypeForID'] = makeInvalidEarlyAccess('_SDL_GetSensorTypeForID');
var _SDL_GetSensorNameForID = Module['_SDL_GetSensorNameForID'] = makeInvalidEarlyAccess('_SDL_GetSensorNameForID');
var _SDL_GetPrimaryDisplay = Module['_SDL_GetPrimaryDisplay'] = makeInvalidEarlyAccess('_SDL_GetPrimaryDisplay');
var _SDL_GetNaturalDisplayOrientation = Module['_SDL_GetNaturalDisplayOrientation'] = makeInvalidEarlyAccess('_SDL_GetNaturalDisplayOrientation');
var _SDL_AttachVirtualJoystick = Module['_SDL_AttachVirtualJoystick'] = makeInvalidEarlyAccess('_SDL_AttachVirtualJoystick');
var _SDL_DetachVirtualJoystick = Module['_SDL_DetachVirtualJoystick'] = makeInvalidEarlyAccess('_SDL_DetachVirtualJoystick');
var _SDL_IsJoystickVirtual = Module['_SDL_IsJoystickVirtual'] = makeInvalidEarlyAccess('_SDL_IsJoystickVirtual');
var _SDL_SetJoystickVirtualAxis = Module['_SDL_SetJoystickVirtualAxis'] = makeInvalidEarlyAccess('_SDL_SetJoystickVirtualAxis');
var _SDL_SetJoystickVirtualBall = Module['_SDL_SetJoystickVirtualBall'] = makeInvalidEarlyAccess('_SDL_SetJoystickVirtualBall');
var _SDL_SetJoystickVirtualButton = Module['_SDL_SetJoystickVirtualButton'] = makeInvalidEarlyAccess('_SDL_SetJoystickVirtualButton');
var _SDL_SetJoystickVirtualHat = Module['_SDL_SetJoystickVirtualHat'] = makeInvalidEarlyAccess('_SDL_SetJoystickVirtualHat');
var _SDL_SetJoystickVirtualTouchpad = Module['_SDL_SetJoystickVirtualTouchpad'] = makeInvalidEarlyAccess('_SDL_SetJoystickVirtualTouchpad');
var _SDL_SendJoystickVirtualSensorData = Module['_SDL_SendJoystickVirtualSensorData'] = makeInvalidEarlyAccess('_SDL_SendJoystickVirtualSensorData');
var _SDL_GetNumJoystickHats = Module['_SDL_GetNumJoystickHats'] = makeInvalidEarlyAccess('_SDL_GetNumJoystickHats');
var _SDL_GetNumJoystickBalls = Module['_SDL_GetNumJoystickBalls'] = makeInvalidEarlyAccess('_SDL_GetNumJoystickBalls');
var _SDL_GetNumJoystickButtons = Module['_SDL_GetNumJoystickButtons'] = makeInvalidEarlyAccess('_SDL_GetNumJoystickButtons');
var _SDL_GetJoystickAxisInitialState = Module['_SDL_GetJoystickAxisInitialState'] = makeInvalidEarlyAccess('_SDL_GetJoystickAxisInitialState');
var _SDL_GetJoystickBall = Module['_SDL_GetJoystickBall'] = makeInvalidEarlyAccess('_SDL_GetJoystickBall');
var _SDL_GetJoystickFromID = Module['_SDL_GetJoystickFromID'] = makeInvalidEarlyAccess('_SDL_GetJoystickFromID');
var _SDL_SetJoystickEventsEnabled = Module['_SDL_SetJoystickEventsEnabled'] = makeInvalidEarlyAccess('_SDL_SetJoystickEventsEnabled');
var _SDL_JoystickEventsEnabled = Module['_SDL_JoystickEventsEnabled'] = makeInvalidEarlyAccess('_SDL_JoystickEventsEnabled');
var _SDL_crc16 = Module['_SDL_crc16'] = makeInvalidEarlyAccess('_SDL_crc16');
var _SDL_GetJoystickTypeForID = Module['_SDL_GetJoystickTypeForID'] = makeInvalidEarlyAccess('_SDL_GetJoystickTypeForID');
var _SDL_GetJoystickType = Module['_SDL_GetJoystickType'] = makeInvalidEarlyAccess('_SDL_GetJoystickType');
var _SDL_strtoul = Module['_SDL_strtoul'] = makeInvalidEarlyAccess('_SDL_strtoul');
var _SDL_strtoull = Module['_SDL_strtoull'] = makeInvalidEarlyAccess('_SDL_strtoull');
var _SDL_GetPreferredLocales = Module['_SDL_GetPreferredLocales'] = makeInvalidEarlyAccess('_SDL_GetPreferredLocales');
var _SDL_OpenURL = Module['_SDL_OpenURL'] = makeInvalidEarlyAccess('_SDL_OpenURL');
var _SDL_GetPowerInfo = Module['_SDL_GetPowerInfo'] = makeInvalidEarlyAccess('_SDL_GetPowerInfo');
var _SDL_DestroyRenderer = Module['_SDL_DestroyRenderer'] = makeInvalidEarlyAccess('_SDL_DestroyRenderer');
var _SDL_GetRendererProperties = Module['_SDL_GetRendererProperties'] = makeInvalidEarlyAccess('_SDL_GetRendererProperties');
var _SDL_FlushRenderer = Module['_SDL_FlushRenderer'] = makeInvalidEarlyAccess('_SDL_FlushRenderer');
var _SDL_GetNumRenderDrivers = Module['_SDL_GetNumRenderDrivers'] = makeInvalidEarlyAccess('_SDL_GetNumRenderDrivers');
var _SDL_GetRenderDriver = Module['_SDL_GetRenderDriver'] = makeInvalidEarlyAccess('_SDL_GetRenderDriver');
var _SDL_CreateWindowAndRenderer = Module['_SDL_CreateWindowAndRenderer'] = makeInvalidEarlyAccess('_SDL_CreateWindowAndRenderer');
var _SDL_CreateWindow = Module['_SDL_CreateWindow'] = makeInvalidEarlyAccess('_SDL_CreateWindow');
var _SDL_CreateRendererWithProperties = Module['_SDL_CreateRendererWithProperties'] = makeInvalidEarlyAccess('_SDL_CreateRendererWithProperties');
var _SDL_DestroyWindow = Module['_SDL_DestroyWindow'] = makeInvalidEarlyAccess('_SDL_DestroyWindow');
var _SDL_ShowWindow = Module['_SDL_ShowWindow'] = makeInvalidEarlyAccess('_SDL_ShowWindow');
var _SDL_CreateRenderer = Module['_SDL_CreateRenderer'] = makeInvalidEarlyAccess('_SDL_CreateRenderer');
var _SDL_WindowHasSurface = Module['_SDL_WindowHasSurface'] = makeInvalidEarlyAccess('_SDL_WindowHasSurface');
var _SDL_GetWindowProperties = Module['_SDL_GetWindowProperties'] = makeInvalidEarlyAccess('_SDL_GetWindowProperties');
var _SDL_GetWindowFlags = Module['_SDL_GetWindowFlags'] = makeInvalidEarlyAccess('_SDL_GetWindowFlags');
var _SDL_SetRenderViewport = Module['_SDL_SetRenderViewport'] = makeInvalidEarlyAccess('_SDL_SetRenderViewport');
var _SDL_SetRenderVSync = Module['_SDL_SetRenderVSync'] = makeInvalidEarlyAccess('_SDL_SetRenderVSync');
var _SDL_GetDisplayForWindow = Module['_SDL_GetDisplayForWindow'] = makeInvalidEarlyAccess('_SDL_GetDisplayForWindow');
var _SDL_GetDesktopDisplayMode = Module['_SDL_GetDesktopDisplayMode'] = makeInvalidEarlyAccess('_SDL_GetDesktopDisplayMode');
var _SDL_GetRenderer = Module['_SDL_GetRenderer'] = makeInvalidEarlyAccess('_SDL_GetRenderer');
var _SDL_GetWindowSize = Module['_SDL_GetWindowSize'] = makeInvalidEarlyAccess('_SDL_GetWindowSize');
var _SDL_GetWindowSizeInPixels = Module['_SDL_GetWindowSizeInPixels'] = makeInvalidEarlyAccess('_SDL_GetWindowSizeInPixels');
var _SDL_CreateSoftwareRenderer = Module['_SDL_CreateSoftwareRenderer'] = makeInvalidEarlyAccess('_SDL_CreateSoftwareRenderer');
var _SDL_GetRenderWindow = Module['_SDL_GetRenderWindow'] = makeInvalidEarlyAccess('_SDL_GetRenderWindow');
var _SDL_GetRendererName = Module['_SDL_GetRendererName'] = makeInvalidEarlyAccess('_SDL_GetRendererName');
var _SDL_GetRenderOutputSize = Module['_SDL_GetRenderOutputSize'] = makeInvalidEarlyAccess('_SDL_GetRenderOutputSize');
var _SDL_GetCurrentRenderOutputSize = Module['_SDL_GetCurrentRenderOutputSize'] = makeInvalidEarlyAccess('_SDL_GetCurrentRenderOutputSize');
var _SDL_CreateTextureWithProperties = Module['_SDL_CreateTextureWithProperties'] = makeInvalidEarlyAccess('_SDL_CreateTextureWithProperties');
var _SDL_GetTextureProperties = Module['_SDL_GetTextureProperties'] = makeInvalidEarlyAccess('_SDL_GetTextureProperties');
var _SDL_GetSurfaceColorspace = Module['_SDL_GetSurfaceColorspace'] = makeInvalidEarlyAccess('_SDL_GetSurfaceColorspace');
var _SDL_LockSurface = Module['_SDL_LockSurface'] = makeInvalidEarlyAccess('_SDL_LockSurface');
var _SDL_UpdateTexture = Module['_SDL_UpdateTexture'] = makeInvalidEarlyAccess('_SDL_UpdateTexture');
var _SDL_UnlockSurface = Module['_SDL_UnlockSurface'] = makeInvalidEarlyAccess('_SDL_UnlockSurface');
var _SDL_ConvertSurfaceAndColorspace = Module['_SDL_ConvertSurfaceAndColorspace'] = makeInvalidEarlyAccess('_SDL_ConvertSurfaceAndColorspace');
var _SDL_GetSurfaceColorMod = Module['_SDL_GetSurfaceColorMod'] = makeInvalidEarlyAccess('_SDL_GetSurfaceColorMod');
var _SDL_GetSurfaceAlphaMod = Module['_SDL_GetSurfaceAlphaMod'] = makeInvalidEarlyAccess('_SDL_GetSurfaceAlphaMod');
var _SDL_GetSurfaceBlendMode = Module['_SDL_GetSurfaceBlendMode'] = makeInvalidEarlyAccess('_SDL_GetSurfaceBlendMode');
var _SDL_SetTextureBlendMode = Module['_SDL_SetTextureBlendMode'] = makeInvalidEarlyAccess('_SDL_SetTextureBlendMode');
var _SDL_GetRendererFromTexture = Module['_SDL_GetRendererFromTexture'] = makeInvalidEarlyAccess('_SDL_GetRendererFromTexture');
var _SDL_GetTextureSize = Module['_SDL_GetTextureSize'] = makeInvalidEarlyAccess('_SDL_GetTextureSize');
var _SDL_SetTextureColorMod = Module['_SDL_SetTextureColorMod'] = makeInvalidEarlyAccess('_SDL_SetTextureColorMod');
var _SDL_SetTextureColorModFloat = Module['_SDL_SetTextureColorModFloat'] = makeInvalidEarlyAccess('_SDL_SetTextureColorModFloat');
var _SDL_GetTextureColorMod = Module['_SDL_GetTextureColorMod'] = makeInvalidEarlyAccess('_SDL_GetTextureColorMod');
var _SDL_GetTextureColorModFloat = Module['_SDL_GetTextureColorModFloat'] = makeInvalidEarlyAccess('_SDL_GetTextureColorModFloat');
var _SDL_SetTextureAlphaMod = Module['_SDL_SetTextureAlphaMod'] = makeInvalidEarlyAccess('_SDL_SetTextureAlphaMod');
var _SDL_SetTextureAlphaModFloat = Module['_SDL_SetTextureAlphaModFloat'] = makeInvalidEarlyAccess('_SDL_SetTextureAlphaModFloat');
var _SDL_GetTextureAlphaMod = Module['_SDL_GetTextureAlphaMod'] = makeInvalidEarlyAccess('_SDL_GetTextureAlphaMod');
var _SDL_GetTextureAlphaModFloat = Module['_SDL_GetTextureAlphaModFloat'] = makeInvalidEarlyAccess('_SDL_GetTextureAlphaModFloat');
var _SDL_GetTextureBlendMode = Module['_SDL_GetTextureBlendMode'] = makeInvalidEarlyAccess('_SDL_GetTextureBlendMode');
var _SDL_GetTextureScaleMode = Module['_SDL_GetTextureScaleMode'] = makeInvalidEarlyAccess('_SDL_GetTextureScaleMode');
var _SDL_ConvertPixelsAndColorspace = Module['_SDL_ConvertPixelsAndColorspace'] = makeInvalidEarlyAccess('_SDL_ConvertPixelsAndColorspace');
var _SDL_UpdateYUVTexture = Module['_SDL_UpdateYUVTexture'] = makeInvalidEarlyAccess('_SDL_UpdateYUVTexture');
var _SDL_UpdateNVTexture = Module['_SDL_UpdateNVTexture'] = makeInvalidEarlyAccess('_SDL_UpdateNVTexture');
var _SDL_LockTextureToSurface = Module['_SDL_LockTextureToSurface'] = makeInvalidEarlyAccess('_SDL_LockTextureToSurface');
var _SDL_SetRenderTarget = Module['_SDL_SetRenderTarget'] = makeInvalidEarlyAccess('_SDL_SetRenderTarget');
var _SDL_GetRenderTarget = Module['_SDL_GetRenderTarget'] = makeInvalidEarlyAccess('_SDL_GetRenderTarget');
var _SDL_SetRenderLogicalPresentation = Module['_SDL_SetRenderLogicalPresentation'] = makeInvalidEarlyAccess('_SDL_SetRenderLogicalPresentation');
var _SDL_GetRenderLogicalPresentation = Module['_SDL_GetRenderLogicalPresentation'] = makeInvalidEarlyAccess('_SDL_GetRenderLogicalPresentation');
var _SDL_GetRenderLogicalPresentationRect = Module['_SDL_GetRenderLogicalPresentationRect'] = makeInvalidEarlyAccess('_SDL_GetRenderLogicalPresentationRect');
var _SDL_RenderCoordinatesFromWindow = Module['_SDL_RenderCoordinatesFromWindow'] = makeInvalidEarlyAccess('_SDL_RenderCoordinatesFromWindow');
var _SDL_RenderCoordinatesToWindow = Module['_SDL_RenderCoordinatesToWindow'] = makeInvalidEarlyAccess('_SDL_RenderCoordinatesToWindow');
var _SDL_ConvertEventToRenderCoordinates = Module['_SDL_ConvertEventToRenderCoordinates'] = makeInvalidEarlyAccess('_SDL_ConvertEventToRenderCoordinates');
var _SDL_GetRenderViewport = Module['_SDL_GetRenderViewport'] = makeInvalidEarlyAccess('_SDL_GetRenderViewport');
var _SDL_RenderViewportSet = Module['_SDL_RenderViewportSet'] = makeInvalidEarlyAccess('_SDL_RenderViewportSet');
var _SDL_GetRenderSafeArea = Module['_SDL_GetRenderSafeArea'] = makeInvalidEarlyAccess('_SDL_GetRenderSafeArea');
var _SDL_GetWindowSafeArea = Module['_SDL_GetWindowSafeArea'] = makeInvalidEarlyAccess('_SDL_GetWindowSafeArea');
var _SDL_SetRenderClipRect = Module['_SDL_SetRenderClipRect'] = makeInvalidEarlyAccess('_SDL_SetRenderClipRect');
var _SDL_GetRenderClipRect = Module['_SDL_GetRenderClipRect'] = makeInvalidEarlyAccess('_SDL_GetRenderClipRect');
var _SDL_RenderClipEnabled = Module['_SDL_RenderClipEnabled'] = makeInvalidEarlyAccess('_SDL_RenderClipEnabled');
var _SDL_SetRenderScale = Module['_SDL_SetRenderScale'] = makeInvalidEarlyAccess('_SDL_SetRenderScale');
var _SDL_GetRenderScale = Module['_SDL_GetRenderScale'] = makeInvalidEarlyAccess('_SDL_GetRenderScale');
var _SDL_SetRenderDrawColor = Module['_SDL_SetRenderDrawColor'] = makeInvalidEarlyAccess('_SDL_SetRenderDrawColor');
var _SDL_SetRenderDrawColorFloat = Module['_SDL_SetRenderDrawColorFloat'] = makeInvalidEarlyAccess('_SDL_SetRenderDrawColorFloat');
var _SDL_GetRenderDrawColor = Module['_SDL_GetRenderDrawColor'] = makeInvalidEarlyAccess('_SDL_GetRenderDrawColor');
var _SDL_GetRenderDrawColorFloat = Module['_SDL_GetRenderDrawColorFloat'] = makeInvalidEarlyAccess('_SDL_GetRenderDrawColorFloat');
var _SDL_SetRenderColorScale = Module['_SDL_SetRenderColorScale'] = makeInvalidEarlyAccess('_SDL_SetRenderColorScale');
var _SDL_GetRenderColorScale = Module['_SDL_GetRenderColorScale'] = makeInvalidEarlyAccess('_SDL_GetRenderColorScale');
var _SDL_SetRenderDrawBlendMode = Module['_SDL_SetRenderDrawBlendMode'] = makeInvalidEarlyAccess('_SDL_SetRenderDrawBlendMode');
var _SDL_GetRenderDrawBlendMode = Module['_SDL_GetRenderDrawBlendMode'] = makeInvalidEarlyAccess('_SDL_GetRenderDrawBlendMode');
var _SDL_RenderClear = Module['_SDL_RenderClear'] = makeInvalidEarlyAccess('_SDL_RenderClear');
var _SDL_RenderPoint = Module['_SDL_RenderPoint'] = makeInvalidEarlyAccess('_SDL_RenderPoint');
var _SDL_RenderPoints = Module['_SDL_RenderPoints'] = makeInvalidEarlyAccess('_SDL_RenderPoints');
var _SDL_RenderLine = Module['_SDL_RenderLine'] = makeInvalidEarlyAccess('_SDL_RenderLine');
var _SDL_RenderLines = Module['_SDL_RenderLines'] = makeInvalidEarlyAccess('_SDL_RenderLines');
var _SDL_RenderRect = Module['_SDL_RenderRect'] = makeInvalidEarlyAccess('_SDL_RenderRect');
var _SDL_RenderRects = Module['_SDL_RenderRects'] = makeInvalidEarlyAccess('_SDL_RenderRects');
var _SDL_RenderFillRect = Module['_SDL_RenderFillRect'] = makeInvalidEarlyAccess('_SDL_RenderFillRect');
var _SDL_RenderFillRects = Module['_SDL_RenderFillRects'] = makeInvalidEarlyAccess('_SDL_RenderFillRects');
var _SDL_RenderTexture = Module['_SDL_RenderTexture'] = makeInvalidEarlyAccess('_SDL_RenderTexture');
var _SDL_GetRectIntersectionFloat = Module['_SDL_GetRectIntersectionFloat'] = makeInvalidEarlyAccess('_SDL_GetRectIntersectionFloat');
var _SDL_RenderTextureAffine = Module['_SDL_RenderTextureAffine'] = makeInvalidEarlyAccess('_SDL_RenderTextureAffine');
var _SDL_RenderTextureRotated = Module['_SDL_RenderTextureRotated'] = makeInvalidEarlyAccess('_SDL_RenderTextureRotated');
var _SDL_RenderTextureTiled = Module['_SDL_RenderTextureTiled'] = makeInvalidEarlyAccess('_SDL_RenderTextureTiled');
var _SDL_RenderTexture9Grid = Module['_SDL_RenderTexture9Grid'] = makeInvalidEarlyAccess('_SDL_RenderTexture9Grid');
var _SDL_RenderGeometry = Module['_SDL_RenderGeometry'] = makeInvalidEarlyAccess('_SDL_RenderGeometry');
var _SDL_RenderReadPixels = Module['_SDL_RenderReadPixels'] = makeInvalidEarlyAccess('_SDL_RenderReadPixels');
var _SDL_GetPixelFormatDetails = Module['_SDL_GetPixelFormatDetails'] = makeInvalidEarlyAccess('_SDL_GetPixelFormatDetails');
var _SDL_RenderPresent = Module['_SDL_RenderPresent'] = makeInvalidEarlyAccess('_SDL_RenderPresent');
var _SDL_DelayPrecise = Module['_SDL_DelayPrecise'] = makeInvalidEarlyAccess('_SDL_DelayPrecise');
var _SDL_GetRenderMetalLayer = Module['_SDL_GetRenderMetalLayer'] = makeInvalidEarlyAccess('_SDL_GetRenderMetalLayer');
var _SDL_GetRenderMetalCommandEncoder = Module['_SDL_GetRenderMetalCommandEncoder'] = makeInvalidEarlyAccess('_SDL_GetRenderMetalCommandEncoder');
var _SDL_AddVulkanRenderSemaphores = Module['_SDL_AddVulkanRenderSemaphores'] = makeInvalidEarlyAccess('_SDL_AddVulkanRenderSemaphores');
var _SDL_ComposeCustomBlendMode = Module['_SDL_ComposeCustomBlendMode'] = makeInvalidEarlyAccess('_SDL_ComposeCustomBlendMode');
var _SDL_GetRenderVSync = Module['_SDL_GetRenderVSync'] = makeInvalidEarlyAccess('_SDL_GetRenderVSync');
var _SDL_RenderDebugText = Module['_SDL_RenderDebugText'] = makeInvalidEarlyAccess('_SDL_RenderDebugText');
var _SDL_RenderDebugTextFormat = Module['_SDL_RenderDebugTextFormat'] = makeInvalidEarlyAccess('_SDL_RenderDebugTextFormat');
var _SDL_GetRectAndLineIntersection = Module['_SDL_GetRectAndLineIntersection'] = makeInvalidEarlyAccess('_SDL_GetRectAndLineIntersection');
var _SDL_GetPixelFormatName = Module['_SDL_GetPixelFormatName'] = makeInvalidEarlyAccess('_SDL_GetPixelFormatName');
var _SDL_GL_GetAttribute = Module['_SDL_GL_GetAttribute'] = makeInvalidEarlyAccess('_SDL_GL_GetAttribute');
var _SDL_SyncWindow = Module['_SDL_SyncWindow'] = makeInvalidEarlyAccess('_SDL_SyncWindow');
var _SDL_GL_SetAttribute = Module['_SDL_GL_SetAttribute'] = makeInvalidEarlyAccess('_SDL_GL_SetAttribute');
var _SDL_GL_CreateContext = Module['_SDL_GL_CreateContext'] = makeInvalidEarlyAccess('_SDL_GL_CreateContext');
var _SDL_GL_MakeCurrent = Module['_SDL_GL_MakeCurrent'] = makeInvalidEarlyAccess('_SDL_GL_MakeCurrent');
var _SDL_GL_ExtensionSupported = Module['_SDL_GL_ExtensionSupported'] = makeInvalidEarlyAccess('_SDL_GL_ExtensionSupported');
var _SDL_GL_GetProcAddress = Module['_SDL_GL_GetProcAddress'] = makeInvalidEarlyAccess('_SDL_GL_GetProcAddress');
var _SDL_GL_GetCurrentContext = Module['_SDL_GL_GetCurrentContext'] = makeInvalidEarlyAccess('_SDL_GL_GetCurrentContext');
var _SDL_FlipSurface = Module['_SDL_FlipSurface'] = makeInvalidEarlyAccess('_SDL_FlipSurface');
var _SDL_GL_SwapWindow = Module['_SDL_GL_SwapWindow'] = makeInvalidEarlyAccess('_SDL_GL_SwapWindow');
var _SDL_GL_DestroyContext = Module['_SDL_GL_DestroyContext'] = makeInvalidEarlyAccess('_SDL_GL_DestroyContext');
var _SDL_GL_SetSwapInterval = Module['_SDL_GL_SetSwapInterval'] = makeInvalidEarlyAccess('_SDL_GL_SetSwapInterval');
var _SDL_GL_GetSwapInterval = Module['_SDL_GL_GetSwapInterval'] = makeInvalidEarlyAccess('_SDL_GL_GetSwapInterval');
var _SDL_GetRGBA = Module['_SDL_GetRGBA'] = makeInvalidEarlyAccess('_SDL_GetRGBA');
var _SDL_SetSurfaceRLE = Module['_SDL_SetSurfaceRLE'] = makeInvalidEarlyAccess('_SDL_SetSurfaceRLE');
var _SDL_GetWindowSurface = Module['_SDL_GetWindowSurface'] = makeInvalidEarlyAccess('_SDL_GetWindowSurface');
var _SDL_SetSurfaceClipRect = Module['_SDL_SetSurfaceClipRect'] = makeInvalidEarlyAccess('_SDL_SetSurfaceClipRect');
var _SDL_FillSurfaceRects = Module['_SDL_FillSurfaceRects'] = makeInvalidEarlyAccess('_SDL_FillSurfaceRects');
var _SDL_BlitSurfaceScaled = Module['_SDL_BlitSurfaceScaled'] = makeInvalidEarlyAccess('_SDL_BlitSurfaceScaled');
var _SDL_UpdateWindowSurface = Module['_SDL_UpdateWindowSurface'] = makeInvalidEarlyAccess('_SDL_UpdateWindowSurface');
var _SDL_DestroyWindowSurface = Module['_SDL_DestroyWindowSurface'] = makeInvalidEarlyAccess('_SDL_DestroyWindowSurface');
var _SDL_ceil = Module['_SDL_ceil'] = makeInvalidEarlyAccess('_SDL_ceil');
var _SDL_GetSurfaceColorKey = Module['_SDL_GetSurfaceColorKey'] = makeInvalidEarlyAccess('_SDL_GetSurfaceColorKey');
var _SDL_GetSurfaceClipRect = Module['_SDL_GetSurfaceClipRect'] = makeInvalidEarlyAccess('_SDL_GetSurfaceClipRect');
var _SDL_MapRGBA = Module['_SDL_MapRGBA'] = makeInvalidEarlyAccess('_SDL_MapRGBA');
var _SDL_GetSensorNonPortableTypeForID = Module['_SDL_GetSensorNonPortableTypeForID'] = makeInvalidEarlyAccess('_SDL_GetSensorNonPortableTypeForID');
var _SDL_GetSensorFromID = Module['_SDL_GetSensorFromID'] = makeInvalidEarlyAccess('_SDL_GetSensorFromID');
var _SDL_GetSensorProperties = Module['_SDL_GetSensorProperties'] = makeInvalidEarlyAccess('_SDL_GetSensorProperties');
var _SDL_GetSensorName = Module['_SDL_GetSensorName'] = makeInvalidEarlyAccess('_SDL_GetSensorName');
var _SDL_GetSensorType = Module['_SDL_GetSensorType'] = makeInvalidEarlyAccess('_SDL_GetSensorType');
var _SDL_GetSensorNonPortableType = Module['_SDL_GetSensorNonPortableType'] = makeInvalidEarlyAccess('_SDL_GetSensorNonPortableType');
var _SDL_GetSensorID = Module['_SDL_GetSensorID'] = makeInvalidEarlyAccess('_SDL_GetSensorID');
var _SDL_GetSensorData = Module['_SDL_GetSensorData'] = makeInvalidEarlyAccess('_SDL_GetSensorData');
var _SDL_crc32 = Module['_SDL_crc32'] = makeInvalidEarlyAccess('_SDL_crc32');
var _SDL_GetEnvironment = Module['_SDL_GetEnvironment'] = makeInvalidEarlyAccess('_SDL_GetEnvironment');
var _SDL_CreateEnvironment = Module['_SDL_CreateEnvironment'] = makeInvalidEarlyAccess('_SDL_CreateEnvironment');
var _SDL_DestroyEnvironment = Module['_SDL_DestroyEnvironment'] = makeInvalidEarlyAccess('_SDL_DestroyEnvironment');
var _SDL_GetEnvironmentVariable = Module['_SDL_GetEnvironmentVariable'] = makeInvalidEarlyAccess('_SDL_GetEnvironmentVariable');
var _SDL_GetEnvironmentVariables = Module['_SDL_GetEnvironmentVariables'] = makeInvalidEarlyAccess('_SDL_GetEnvironmentVariables');
var _SDL_SetEnvironmentVariable = Module['_SDL_SetEnvironmentVariable'] = makeInvalidEarlyAccess('_SDL_SetEnvironmentVariable');
var _SDL_UnsetEnvironmentVariable = Module['_SDL_UnsetEnvironmentVariable'] = makeInvalidEarlyAccess('_SDL_UnsetEnvironmentVariable');
var _SDL_setenv_unsafe = Module['_SDL_setenv_unsafe'] = makeInvalidEarlyAccess('_SDL_setenv_unsafe');
var _SDL_unsetenv_unsafe = Module['_SDL_unsetenv_unsafe'] = makeInvalidEarlyAccess('_SDL_unsetenv_unsafe');
var _SDL_getenv_unsafe = Module['_SDL_getenv_unsafe'] = makeInvalidEarlyAccess('_SDL_getenv_unsafe');
var _SDL_iconv_open = Module['_SDL_iconv_open'] = makeInvalidEarlyAccess('_SDL_iconv_open');
var _SDL_iconv_close = Module['_SDL_iconv_close'] = makeInvalidEarlyAccess('_SDL_iconv_close');
var _SDL_iconv = Module['_SDL_iconv'] = makeInvalidEarlyAccess('_SDL_iconv');
var _SDL_GetOriginalMemoryFunctions = Module['_SDL_GetOriginalMemoryFunctions'] = makeInvalidEarlyAccess('_SDL_GetOriginalMemoryFunctions');
var _SDL_GetMemoryFunctions = Module['_SDL_GetMemoryFunctions'] = makeInvalidEarlyAccess('_SDL_GetMemoryFunctions');
var _SDL_SetMemoryFunctions = Module['_SDL_SetMemoryFunctions'] = makeInvalidEarlyAccess('_SDL_SetMemoryFunctions');
var _SDL_GetNumAllocations = Module['_SDL_GetNumAllocations'] = makeInvalidEarlyAccess('_SDL_GetNumAllocations');
var _SDL_memmove = Module['_SDL_memmove'] = makeInvalidEarlyAccess('_SDL_memmove');
var _SDL_bsearch_r = Module['_SDL_bsearch_r'] = makeInvalidEarlyAccess('_SDL_bsearch_r');
var _SDL_bsearch = Module['_SDL_bsearch'] = makeInvalidEarlyAccess('_SDL_bsearch');
var _SDL_srand = Module['_SDL_srand'] = makeInvalidEarlyAccess('_SDL_srand');
var _SDL_GetPerformanceCounter = Module['_SDL_GetPerformanceCounter'] = makeInvalidEarlyAccess('_SDL_GetPerformanceCounter');
var _SDL_rand = Module['_SDL_rand'] = makeInvalidEarlyAccess('_SDL_rand');
var _SDL_rand_r = Module['_SDL_rand_r'] = makeInvalidEarlyAccess('_SDL_rand_r');
var _SDL_randf = Module['_SDL_randf'] = makeInvalidEarlyAccess('_SDL_randf');
var _SDL_randf_r = Module['_SDL_randf_r'] = makeInvalidEarlyAccess('_SDL_randf_r');
var _SDL_rand_bits = Module['_SDL_rand_bits'] = makeInvalidEarlyAccess('_SDL_rand_bits');
var _SDL_rand_bits_r = Module['_SDL_rand_bits_r'] = makeInvalidEarlyAccess('_SDL_rand_bits_r');
var _SDL_atan = Module['_SDL_atan'] = makeInvalidEarlyAccess('_SDL_atan');
var _SDL_atanf = Module['_SDL_atanf'] = makeInvalidEarlyAccess('_SDL_atanf');
var _SDL_atan2 = Module['_SDL_atan2'] = makeInvalidEarlyAccess('_SDL_atan2');
var _SDL_acos = Module['_SDL_acos'] = makeInvalidEarlyAccess('_SDL_acos');
var _SDL_asin = Module['_SDL_asin'] = makeInvalidEarlyAccess('_SDL_asin');
var _SDL_asinf = Module['_SDL_asinf'] = makeInvalidEarlyAccess('_SDL_asinf');
var _SDL_copysign = Module['_SDL_copysign'] = makeInvalidEarlyAccess('_SDL_copysign');
var _SDL_copysignf = Module['_SDL_copysignf'] = makeInvalidEarlyAccess('_SDL_copysignf');
var _SDL_expf = Module['_SDL_expf'] = makeInvalidEarlyAccess('_SDL_expf');
var _SDL_trunc = Module['_SDL_trunc'] = makeInvalidEarlyAccess('_SDL_trunc');
var _SDL_fmod = Module['_SDL_fmod'] = makeInvalidEarlyAccess('_SDL_fmod');
var _SDL_isinf = Module['_SDL_isinf'] = makeInvalidEarlyAccess('_SDL_isinf');
var _SDL_isinff = Module['_SDL_isinff'] = makeInvalidEarlyAccess('_SDL_isinff');
var _SDL_isnan = Module['_SDL_isnan'] = makeInvalidEarlyAccess('_SDL_isnan');
var _SDL_isnanf = Module['_SDL_isnanf'] = makeInvalidEarlyAccess('_SDL_isnanf');
var _SDL_logf = Module['_SDL_logf'] = makeInvalidEarlyAccess('_SDL_logf');
var _SDL_log10 = Module['_SDL_log10'] = makeInvalidEarlyAccess('_SDL_log10');
var _SDL_log10f = Module['_SDL_log10f'] = makeInvalidEarlyAccess('_SDL_log10f');
var _SDL_modf = Module['_SDL_modf'] = makeInvalidEarlyAccess('_SDL_modf');
var _SDL_powf = Module['_SDL_powf'] = makeInvalidEarlyAccess('_SDL_powf');
var _SDL_lround = Module['_SDL_lround'] = makeInvalidEarlyAccess('_SDL_lround');
var _SDL_lroundf = Module['_SDL_lroundf'] = makeInvalidEarlyAccess('_SDL_lroundf');
var _SDL_scalbnf = Module['_SDL_scalbnf'] = makeInvalidEarlyAccess('_SDL_scalbnf');
var _SDL_tan = Module['_SDL_tan'] = makeInvalidEarlyAccess('_SDL_tan');
var _SDL_isalpha = Module['_SDL_isalpha'] = makeInvalidEarlyAccess('_SDL_isalpha');
var _SDL_isupper = Module['_SDL_isupper'] = makeInvalidEarlyAccess('_SDL_isupper');
var _SDL_islower = Module['_SDL_islower'] = makeInvalidEarlyAccess('_SDL_islower');
var _SDL_isalnum = Module['_SDL_isalnum'] = makeInvalidEarlyAccess('_SDL_isalnum');
var _SDL_isxdigit = Module['_SDL_isxdigit'] = makeInvalidEarlyAccess('_SDL_isxdigit');
var _SDL_ispunct = Module['_SDL_ispunct'] = makeInvalidEarlyAccess('_SDL_ispunct');
var _SDL_isgraph = Module['_SDL_isgraph'] = makeInvalidEarlyAccess('_SDL_isgraph');
var _SDL_isprint = Module['_SDL_isprint'] = makeInvalidEarlyAccess('_SDL_isprint');
var _SDL_toupper = Module['_SDL_toupper'] = makeInvalidEarlyAccess('_SDL_toupper');
var _SDL_isblank = Module['_SDL_isblank'] = makeInvalidEarlyAccess('_SDL_isblank');
var _SDL_StepBackUTF8 = Module['_SDL_StepBackUTF8'] = makeInvalidEarlyAccess('_SDL_StepBackUTF8');
var _SDL_strnlen = Module['_SDL_strnlen'] = makeInvalidEarlyAccess('_SDL_strnlen');
var _SDL_wcslen = Module['_SDL_wcslen'] = makeInvalidEarlyAccess('_SDL_wcslen');
var _SDL_wcsnlen = Module['_SDL_wcsnlen'] = makeInvalidEarlyAccess('_SDL_wcsnlen');
var _SDL_wcslcpy = Module['_SDL_wcslcpy'] = makeInvalidEarlyAccess('_SDL_wcslcpy');
var _SDL_wcslcat = Module['_SDL_wcslcat'] = makeInvalidEarlyAccess('_SDL_wcslcat');
var _SDL_wcsnstr = Module['_SDL_wcsnstr'] = makeInvalidEarlyAccess('_SDL_wcsnstr');
var _SDL_wcsncmp = Module['_SDL_wcsncmp'] = makeInvalidEarlyAccess('_SDL_wcsncmp');
var _SDL_wcsstr = Module['_SDL_wcsstr'] = makeInvalidEarlyAccess('_SDL_wcsstr');
var _SDL_wcscmp = Module['_SDL_wcscmp'] = makeInvalidEarlyAccess('_SDL_wcscmp');
var _SDL_wcscasecmp = Module['_SDL_wcscasecmp'] = makeInvalidEarlyAccess('_SDL_wcscasecmp');
var _SDL_wcsncasecmp = Module['_SDL_wcsncasecmp'] = makeInvalidEarlyAccess('_SDL_wcsncasecmp');
var _SDL_wcstol = Module['_SDL_wcstol'] = makeInvalidEarlyAccess('_SDL_wcstol');
var _SDL_utf8strlcpy = Module['_SDL_utf8strlcpy'] = makeInvalidEarlyAccess('_SDL_utf8strlcpy');
var _SDL_utf8strlen = Module['_SDL_utf8strlen'] = makeInvalidEarlyAccess('_SDL_utf8strlen');
var _SDL_utf8strnlen = Module['_SDL_utf8strnlen'] = makeInvalidEarlyAccess('_SDL_utf8strnlen');
var _SDL_strndup = Module['_SDL_strndup'] = makeInvalidEarlyAccess('_SDL_strndup');
var _SDL_strrev = Module['_SDL_strrev'] = makeInvalidEarlyAccess('_SDL_strrev');
var _SDL_strupr = Module['_SDL_strupr'] = makeInvalidEarlyAccess('_SDL_strupr');
var _SDL_strlwr = Module['_SDL_strlwr'] = makeInvalidEarlyAccess('_SDL_strlwr');
var _SDL_strnstr = Module['_SDL_strnstr'] = makeInvalidEarlyAccess('_SDL_strnstr');
var _SDL_itoa = Module['_SDL_itoa'] = makeInvalidEarlyAccess('_SDL_itoa');
var _SDL_ltoa = Module['_SDL_ltoa'] = makeInvalidEarlyAccess('_SDL_ltoa');
var _SDL_uitoa = Module['_SDL_uitoa'] = makeInvalidEarlyAccess('_SDL_uitoa');
var _SDL_ultoa = Module['_SDL_ultoa'] = makeInvalidEarlyAccess('_SDL_ultoa');
var _SDL_lltoa = Module['_SDL_lltoa'] = makeInvalidEarlyAccess('_SDL_lltoa');
var _SDL_ulltoa = Module['_SDL_ulltoa'] = makeInvalidEarlyAccess('_SDL_ulltoa');
var _SDL_strtod = Module['_SDL_strtod'] = makeInvalidEarlyAccess('_SDL_strtod');
var _SDL_vsscanf = Module['_SDL_vsscanf'] = makeInvalidEarlyAccess('_SDL_vsscanf');
var _SDL_swprintf = Module['_SDL_swprintf'] = makeInvalidEarlyAccess('_SDL_swprintf');
var _SDL_vswprintf = Module['_SDL_vswprintf'] = makeInvalidEarlyAccess('_SDL_vswprintf');
var _SDL_strpbrk = Module['_SDL_strpbrk'] = makeInvalidEarlyAccess('_SDL_strpbrk');
var _SDL_OpenTitleStorage = Module['_SDL_OpenTitleStorage'] = makeInvalidEarlyAccess('_SDL_OpenTitleStorage');
var _SDL_OpenUserStorage = Module['_SDL_OpenUserStorage'] = makeInvalidEarlyAccess('_SDL_OpenUserStorage');
var _SDL_OpenFileStorage = Module['_SDL_OpenFileStorage'] = makeInvalidEarlyAccess('_SDL_OpenFileStorage');
var _SDL_OpenStorage = Module['_SDL_OpenStorage'] = makeInvalidEarlyAccess('_SDL_OpenStorage');
var _SDL_CloseStorage = Module['_SDL_CloseStorage'] = makeInvalidEarlyAccess('_SDL_CloseStorage');
var _SDL_StorageReady = Module['_SDL_StorageReady'] = makeInvalidEarlyAccess('_SDL_StorageReady');
var _SDL_GetStorageFileSize = Module['_SDL_GetStorageFileSize'] = makeInvalidEarlyAccess('_SDL_GetStorageFileSize');
var _SDL_GetStoragePathInfo = Module['_SDL_GetStoragePathInfo'] = makeInvalidEarlyAccess('_SDL_GetStoragePathInfo');
var _SDL_ReadStorageFile = Module['_SDL_ReadStorageFile'] = makeInvalidEarlyAccess('_SDL_ReadStorageFile');
var _SDL_WriteStorageFile = Module['_SDL_WriteStorageFile'] = makeInvalidEarlyAccess('_SDL_WriteStorageFile');
var _SDL_CreateStorageDirectory = Module['_SDL_CreateStorageDirectory'] = makeInvalidEarlyAccess('_SDL_CreateStorageDirectory');
var _SDL_EnumerateStorageDirectory = Module['_SDL_EnumerateStorageDirectory'] = makeInvalidEarlyAccess('_SDL_EnumerateStorageDirectory');
var _SDL_RemoveStoragePath = Module['_SDL_RemoveStoragePath'] = makeInvalidEarlyAccess('_SDL_RemoveStoragePath');
var _SDL_RenameStoragePath = Module['_SDL_RenameStoragePath'] = makeInvalidEarlyAccess('_SDL_RenameStoragePath');
var _SDL_CopyStorageFile = Module['_SDL_CopyStorageFile'] = makeInvalidEarlyAccess('_SDL_CopyStorageFile');
var _SDL_GetStorageSpaceRemaining = Module['_SDL_GetStorageSpaceRemaining'] = makeInvalidEarlyAccess('_SDL_GetStorageSpaceRemaining');
var _SDL_GlobStorageDirectory = Module['_SDL_GlobStorageDirectory'] = makeInvalidEarlyAccess('_SDL_GlobStorageDirectory');
var _SDL_CleanupTLS = Module['_SDL_CleanupTLS'] = makeInvalidEarlyAccess('_SDL_CleanupTLS');
var _SDL_GetThreadState = Module['_SDL_GetThreadState'] = makeInvalidEarlyAccess('_SDL_GetThreadState');
var _SDL_CreateThreadWithPropertiesRuntime = Module['_SDL_CreateThreadWithPropertiesRuntime'] = makeInvalidEarlyAccess('_SDL_CreateThreadWithPropertiesRuntime');
var _SDL_GetThreadID = Module['_SDL_GetThreadID'] = makeInvalidEarlyAccess('_SDL_GetThreadID');
var _SDL_GetThreadName = Module['_SDL_GetThreadName'] = makeInvalidEarlyAccess('_SDL_GetThreadName');
var _SDL_DetachThread = Module['_SDL_DetachThread'] = makeInvalidEarlyAccess('_SDL_DetachThread');
var _SDL_TryWaitSemaphore = Module['_SDL_TryWaitSemaphore'] = makeInvalidEarlyAccess('_SDL_TryWaitSemaphore');
var _SDL_WaitSemaphoreTimeout = Module['_SDL_WaitSemaphoreTimeout'] = makeInvalidEarlyAccess('_SDL_WaitSemaphoreTimeout');
var _SDL_GetDateTimeLocalePreferences = Module['_SDL_GetDateTimeLocalePreferences'] = makeInvalidEarlyAccess('_SDL_GetDateTimeLocalePreferences');
var _SDL_GetDaysInMonth = Module['_SDL_GetDaysInMonth'] = makeInvalidEarlyAccess('_SDL_GetDaysInMonth');
var _SDL_GetDayOfYear = Module['_SDL_GetDayOfYear'] = makeInvalidEarlyAccess('_SDL_GetDayOfYear');
var _SDL_GetDayOfWeek = Module['_SDL_GetDayOfWeek'] = makeInvalidEarlyAccess('_SDL_GetDayOfWeek');
var _SDL_DateTimeToTime = Module['_SDL_DateTimeToTime'] = makeInvalidEarlyAccess('_SDL_DateTimeToTime');
var _SDL_TimeToWindows = Module['_SDL_TimeToWindows'] = makeInvalidEarlyAccess('_SDL_TimeToWindows');
var _SDL_TimeFromWindows = Module['_SDL_TimeFromWindows'] = makeInvalidEarlyAccess('_SDL_TimeFromWindows');
var _SDL_AddTimer = Module['_SDL_AddTimer'] = makeInvalidEarlyAccess('_SDL_AddTimer');
var _SDL_AddTimerNS = Module['_SDL_AddTimerNS'] = makeInvalidEarlyAccess('_SDL_AddTimerNS');
var _SDL_RemoveTimer = Module['_SDL_RemoveTimer'] = makeInvalidEarlyAccess('_SDL_RemoveTimer');
var _SDL_GetPerformanceFrequency = Module['_SDL_GetPerformanceFrequency'] = makeInvalidEarlyAccess('_SDL_GetPerformanceFrequency');
var _SDL_GetPixelFormatForMasks = Module['_SDL_GetPixelFormatForMasks'] = makeInvalidEarlyAccess('_SDL_GetPixelFormatForMasks');
var _SDL_LoadBMP = Module['_SDL_LoadBMP'] = makeInvalidEarlyAccess('_SDL_LoadBMP');
var _SDL_SaveBMP_IO = Module['_SDL_SaveBMP_IO'] = makeInvalidEarlyAccess('_SDL_SaveBMP_IO');
var _SDL_SaveBMP = Module['_SDL_SaveBMP'] = makeInvalidEarlyAccess('_SDL_SaveBMP');
var _SDL_SetClipboardData = Module['_SDL_SetClipboardData'] = makeInvalidEarlyAccess('_SDL_SetClipboardData');
var _SDL_ClearClipboardData = Module['_SDL_ClearClipboardData'] = makeInvalidEarlyAccess('_SDL_ClearClipboardData');
var _SDL_GetClipboardData = Module['_SDL_GetClipboardData'] = makeInvalidEarlyAccess('_SDL_GetClipboardData');
var _SDL_HasClipboardData = Module['_SDL_HasClipboardData'] = makeInvalidEarlyAccess('_SDL_HasClipboardData');
var _SDL_GetClipboardMimeTypes = Module['_SDL_GetClipboardMimeTypes'] = makeInvalidEarlyAccess('_SDL_GetClipboardMimeTypes');
var _SDL_SetClipboardText = Module['_SDL_SetClipboardText'] = makeInvalidEarlyAccess('_SDL_SetClipboardText');
var _SDL_GetClipboardText = Module['_SDL_GetClipboardText'] = makeInvalidEarlyAccess('_SDL_GetClipboardText');
var _SDL_HasClipboardText = Module['_SDL_HasClipboardText'] = makeInvalidEarlyAccess('_SDL_HasClipboardText');
var _SDL_SetPrimarySelectionText = Module['_SDL_SetPrimarySelectionText'] = makeInvalidEarlyAccess('_SDL_SetPrimarySelectionText');
var _SDL_GetPrimarySelectionText = Module['_SDL_GetPrimarySelectionText'] = makeInvalidEarlyAccess('_SDL_GetPrimarySelectionText');
var _SDL_HasPrimarySelectionText = Module['_SDL_HasPrimarySelectionText'] = makeInvalidEarlyAccess('_SDL_HasPrimarySelectionText');
var _SDL_UnloadObject = Module['_SDL_UnloadObject'] = makeInvalidEarlyAccess('_SDL_UnloadObject');
var _SDL_LoadObject = Module['_SDL_LoadObject'] = makeInvalidEarlyAccess('_SDL_LoadObject');
var _SDL_LoadFunction = Module['_SDL_LoadFunction'] = makeInvalidEarlyAccess('_SDL_LoadFunction');
var _SDL_GetMasksForPixelFormat = Module['_SDL_GetMasksForPixelFormat'] = makeInvalidEarlyAccess('_SDL_GetMasksForPixelFormat');
var _SDL_SetPaletteColors = Module['_SDL_SetPaletteColors'] = makeInvalidEarlyAccess('_SDL_SetPaletteColors');
var _SDL_MapRGB = Module['_SDL_MapRGB'] = makeInvalidEarlyAccess('_SDL_MapRGB');
var _SDL_GetRGB = Module['_SDL_GetRGB'] = makeInvalidEarlyAccess('_SDL_GetRGB');
var _SDL_HasRectIntersection = Module['_SDL_HasRectIntersection'] = makeInvalidEarlyAccess('_SDL_HasRectIntersection');
var _SDL_GetRectEnclosingPoints = Module['_SDL_GetRectEnclosingPoints'] = makeInvalidEarlyAccess('_SDL_GetRectEnclosingPoints');
var _SDL_HasRectIntersectionFloat = Module['_SDL_HasRectIntersectionFloat'] = makeInvalidEarlyAccess('_SDL_HasRectIntersectionFloat');
var _SDL_GetRectUnionFloat = Module['_SDL_GetRectUnionFloat'] = makeInvalidEarlyAccess('_SDL_GetRectUnionFloat');
var _SDL_GetRectEnclosingPointsFloat = Module['_SDL_GetRectEnclosingPointsFloat'] = makeInvalidEarlyAccess('_SDL_GetRectEnclosingPointsFloat');
var _SDL_GetRectAndLineIntersectionFloat = Module['_SDL_GetRectAndLineIntersectionFloat'] = makeInvalidEarlyAccess('_SDL_GetRectAndLineIntersectionFloat');
var _SDL_SurfaceHasRLE = Module['_SDL_SurfaceHasRLE'] = makeInvalidEarlyAccess('_SDL_SurfaceHasRLE');
var _SDL_AddSurfaceAlternateImage = Module['_SDL_AddSurfaceAlternateImage'] = makeInvalidEarlyAccess('_SDL_AddSurfaceAlternateImage');
var _SDL_SurfaceHasAlternateImages = Module['_SDL_SurfaceHasAlternateImages'] = makeInvalidEarlyAccess('_SDL_SurfaceHasAlternateImages');
var _SDL_GetSurfaceImages = Module['_SDL_GetSurfaceImages'] = makeInvalidEarlyAccess('_SDL_GetSurfaceImages');
var _SDL_ScaleSurface = Module['_SDL_ScaleSurface'] = makeInvalidEarlyAccess('_SDL_ScaleSurface');
var _SDL_RemoveSurfaceAlternateImages = Module['_SDL_RemoveSurfaceAlternateImages'] = makeInvalidEarlyAccess('_SDL_RemoveSurfaceAlternateImages');
var _SDL_BlitSurfaceUnchecked = Module['_SDL_BlitSurfaceUnchecked'] = makeInvalidEarlyAccess('_SDL_BlitSurfaceUnchecked');
var _SDL_BlitSurfaceUncheckedScaled = Module['_SDL_BlitSurfaceUncheckedScaled'] = makeInvalidEarlyAccess('_SDL_BlitSurfaceUncheckedScaled');
var _SDL_BlitSurfaceTiled = Module['_SDL_BlitSurfaceTiled'] = makeInvalidEarlyAccess('_SDL_BlitSurfaceTiled');
var _SDL_BlitSurfaceTiledWithScale = Module['_SDL_BlitSurfaceTiledWithScale'] = makeInvalidEarlyAccess('_SDL_BlitSurfaceTiledWithScale');
var _SDL_BlitSurface9Grid = Module['_SDL_BlitSurface9Grid'] = makeInvalidEarlyAccess('_SDL_BlitSurface9Grid');
var _SDL_PremultiplyAlpha = Module['_SDL_PremultiplyAlpha'] = makeInvalidEarlyAccess('_SDL_PremultiplyAlpha');
var _SDL_PremultiplySurfaceAlpha = Module['_SDL_PremultiplySurfaceAlpha'] = makeInvalidEarlyAccess('_SDL_PremultiplySurfaceAlpha');
var _SDL_ClearSurface = Module['_SDL_ClearSurface'] = makeInvalidEarlyAccess('_SDL_ClearSurface');
var _SDL_MapSurfaceRGB = Module['_SDL_MapSurfaceRGB'] = makeInvalidEarlyAccess('_SDL_MapSurfaceRGB');
var _SDL_ReadSurfacePixel = Module['_SDL_ReadSurfacePixel'] = makeInvalidEarlyAccess('_SDL_ReadSurfacePixel');
var _SDL_ReadSurfacePixelFloat = Module['_SDL_ReadSurfacePixelFloat'] = makeInvalidEarlyAccess('_SDL_ReadSurfacePixelFloat');
var _SDL_WriteSurfacePixel = Module['_SDL_WriteSurfacePixel'] = makeInvalidEarlyAccess('_SDL_WriteSurfacePixel');
var _SDL_WriteSurfacePixelFloat = Module['_SDL_WriteSurfacePixelFloat'] = makeInvalidEarlyAccess('_SDL_WriteSurfacePixelFloat');
var _SDL_GetNumVideoDrivers = Module['_SDL_GetNumVideoDrivers'] = makeInvalidEarlyAccess('_SDL_GetNumVideoDrivers');
var _SDL_GetVideoDriver = Module['_SDL_GetVideoDriver'] = makeInvalidEarlyAccess('_SDL_GetVideoDriver');
var _SDL_GL_ResetAttributes = Module['_SDL_GL_ResetAttributes'] = makeInvalidEarlyAccess('_SDL_GL_ResetAttributes');
var _SDL_DisableScreenSaver = Module['_SDL_DisableScreenSaver'] = makeInvalidEarlyAccess('_SDL_DisableScreenSaver');
var _SDL_GetSystemTheme = Module['_SDL_GetSystemTheme'] = makeInvalidEarlyAccess('_SDL_GetSystemTheme');
var _SDL_GetDisplayBounds = Module['_SDL_GetDisplayBounds'] = makeInvalidEarlyAccess('_SDL_GetDisplayBounds');
var _SDL_GetDisplays = Module['_SDL_GetDisplays'] = makeInvalidEarlyAccess('_SDL_GetDisplays');
var _SDL_GetDisplayProperties = Module['_SDL_GetDisplayProperties'] = makeInvalidEarlyAccess('_SDL_GetDisplayProperties');
var _SDL_GetDisplayName = Module['_SDL_GetDisplayName'] = makeInvalidEarlyAccess('_SDL_GetDisplayName');
var _SDL_GetDisplayUsableBounds = Module['_SDL_GetDisplayUsableBounds'] = makeInvalidEarlyAccess('_SDL_GetDisplayUsableBounds');
var _SDL_GetCurrentDisplayOrientation = Module['_SDL_GetCurrentDisplayOrientation'] = makeInvalidEarlyAccess('_SDL_GetCurrentDisplayOrientation');
var _SDL_GetWindowPixelDensity = Module['_SDL_GetWindowPixelDensity'] = makeInvalidEarlyAccess('_SDL_GetWindowPixelDensity');
var _SDL_GetDisplayContentScale = Module['_SDL_GetDisplayContentScale'] = makeInvalidEarlyAccess('_SDL_GetDisplayContentScale');
var _SDL_GetFullscreenDisplayModes = Module['_SDL_GetFullscreenDisplayModes'] = makeInvalidEarlyAccess('_SDL_GetFullscreenDisplayModes');
var _SDL_GetClosestFullscreenDisplayMode = Module['_SDL_GetClosestFullscreenDisplayMode'] = makeInvalidEarlyAccess('_SDL_GetClosestFullscreenDisplayMode');
var _SDL_GetCurrentDisplayMode = Module['_SDL_GetCurrentDisplayMode'] = makeInvalidEarlyAccess('_SDL_GetCurrentDisplayMode');
var _SDL_GetDisplayForPoint = Module['_SDL_GetDisplayForPoint'] = makeInvalidEarlyAccess('_SDL_GetDisplayForPoint');
var _SDL_GetDisplayForRect = Module['_SDL_GetDisplayForRect'] = makeInvalidEarlyAccess('_SDL_GetDisplayForRect');
var _SDL_GetWindowDisplayScale = Module['_SDL_GetWindowDisplayScale'] = makeInvalidEarlyAccess('_SDL_GetWindowDisplayScale');
var _SDL_GetWindowFullscreenMode = Module['_SDL_GetWindowFullscreenMode'] = makeInvalidEarlyAccess('_SDL_GetWindowFullscreenMode');
var _SDL_SetWindowFullscreenMode = Module['_SDL_SetWindowFullscreenMode'] = makeInvalidEarlyAccess('_SDL_SetWindowFullscreenMode');
var _SDL_GetWindowICCProfile = Module['_SDL_GetWindowICCProfile'] = makeInvalidEarlyAccess('_SDL_GetWindowICCProfile');
var _SDL_GetWindowPixelFormat = Module['_SDL_GetWindowPixelFormat'] = makeInvalidEarlyAccess('_SDL_GetWindowPixelFormat');
var _SDL_GetWindows = Module['_SDL_GetWindows'] = makeInvalidEarlyAccess('_SDL_GetWindows');
var _SDL_CreateWindowWithProperties = Module['_SDL_CreateWindowWithProperties'] = makeInvalidEarlyAccess('_SDL_CreateWindowWithProperties');
var _SDL_Vulkan_LoadLibrary = Module['_SDL_Vulkan_LoadLibrary'] = makeInvalidEarlyAccess('_SDL_Vulkan_LoadLibrary');
var _SDL_SetWindowTitle = Module['_SDL_SetWindowTitle'] = makeInvalidEarlyAccess('_SDL_SetWindowTitle');
var _SDL_GL_LoadLibrary = Module['_SDL_GL_LoadLibrary'] = makeInvalidEarlyAccess('_SDL_GL_LoadLibrary');
var _SDL_HideWindow = Module['_SDL_HideWindow'] = makeInvalidEarlyAccess('_SDL_HideWindow');
var _SDL_CreatePopupWindow = Module['_SDL_CreatePopupWindow'] = makeInvalidEarlyAccess('_SDL_CreatePopupWindow');
var _SDL_SetWindowModal = Module['_SDL_SetWindowModal'] = makeInvalidEarlyAccess('_SDL_SetWindowModal');
var _SDL_GL_UnloadLibrary = Module['_SDL_GL_UnloadLibrary'] = makeInvalidEarlyAccess('_SDL_GL_UnloadLibrary');
var _SDL_Vulkan_UnloadLibrary = Module['_SDL_Vulkan_UnloadLibrary'] = makeInvalidEarlyAccess('_SDL_Vulkan_UnloadLibrary');
var _SDL_GetWindowParent = Module['_SDL_GetWindowParent'] = makeInvalidEarlyAccess('_SDL_GetWindowParent');
var _SDL_GetWindowTitle = Module['_SDL_GetWindowTitle'] = makeInvalidEarlyAccess('_SDL_GetWindowTitle');
var _SDL_SetWindowIcon = Module['_SDL_SetWindowIcon'] = makeInvalidEarlyAccess('_SDL_SetWindowIcon');
var _SDL_SetWindowPosition = Module['_SDL_SetWindowPosition'] = makeInvalidEarlyAccess('_SDL_SetWindowPosition');
var _SDL_GetWindowPosition = Module['_SDL_GetWindowPosition'] = makeInvalidEarlyAccess('_SDL_GetWindowPosition');
var _SDL_SetWindowBordered = Module['_SDL_SetWindowBordered'] = makeInvalidEarlyAccess('_SDL_SetWindowBordered');
var _SDL_SetWindowResizable = Module['_SDL_SetWindowResizable'] = makeInvalidEarlyAccess('_SDL_SetWindowResizable');
var _SDL_SetWindowAlwaysOnTop = Module['_SDL_SetWindowAlwaysOnTop'] = makeInvalidEarlyAccess('_SDL_SetWindowAlwaysOnTop');
var _SDL_SetWindowSize = Module['_SDL_SetWindowSize'] = makeInvalidEarlyAccess('_SDL_SetWindowSize');
var _SDL_SetWindowAspectRatio = Module['_SDL_SetWindowAspectRatio'] = makeInvalidEarlyAccess('_SDL_SetWindowAspectRatio');
var _SDL_GetWindowAspectRatio = Module['_SDL_GetWindowAspectRatio'] = makeInvalidEarlyAccess('_SDL_GetWindowAspectRatio');
var _SDL_GetWindowBordersSize = Module['_SDL_GetWindowBordersSize'] = makeInvalidEarlyAccess('_SDL_GetWindowBordersSize');
var _SDL_SetWindowMinimumSize = Module['_SDL_SetWindowMinimumSize'] = makeInvalidEarlyAccess('_SDL_SetWindowMinimumSize');
var _SDL_GetWindowMinimumSize = Module['_SDL_GetWindowMinimumSize'] = makeInvalidEarlyAccess('_SDL_GetWindowMinimumSize');
var _SDL_SetWindowMaximumSize = Module['_SDL_SetWindowMaximumSize'] = makeInvalidEarlyAccess('_SDL_SetWindowMaximumSize');
var _SDL_GetWindowMaximumSize = Module['_SDL_GetWindowMaximumSize'] = makeInvalidEarlyAccess('_SDL_GetWindowMaximumSize');
var _SDL_RaiseWindow = Module['_SDL_RaiseWindow'] = makeInvalidEarlyAccess('_SDL_RaiseWindow');
var _SDL_MaximizeWindow = Module['_SDL_MaximizeWindow'] = makeInvalidEarlyAccess('_SDL_MaximizeWindow');
var _SDL_SetWindowFullscreen = Module['_SDL_SetWindowFullscreen'] = makeInvalidEarlyAccess('_SDL_SetWindowFullscreen');
var _SDL_SetWindowSurfaceVSync = Module['_SDL_SetWindowSurfaceVSync'] = makeInvalidEarlyAccess('_SDL_SetWindowSurfaceVSync');
var _SDL_GetWindowSurfaceVSync = Module['_SDL_GetWindowSurfaceVSync'] = makeInvalidEarlyAccess('_SDL_GetWindowSurfaceVSync');
var _SDL_UpdateWindowSurfaceRects = Module['_SDL_UpdateWindowSurfaceRects'] = makeInvalidEarlyAccess('_SDL_UpdateWindowSurfaceRects');
var _SDL_SetWindowOpacity = Module['_SDL_SetWindowOpacity'] = makeInvalidEarlyAccess('_SDL_SetWindowOpacity');
var _SDL_GetWindowOpacity = Module['_SDL_GetWindowOpacity'] = makeInvalidEarlyAccess('_SDL_GetWindowOpacity');
var _SDL_SetWindowParent = Module['_SDL_SetWindowParent'] = makeInvalidEarlyAccess('_SDL_SetWindowParent');
var _SDL_SetWindowFocusable = Module['_SDL_SetWindowFocusable'] = makeInvalidEarlyAccess('_SDL_SetWindowFocusable');
var _SDL_SetWindowKeyboardGrab = Module['_SDL_SetWindowKeyboardGrab'] = makeInvalidEarlyAccess('_SDL_SetWindowKeyboardGrab');
var _SDL_SetWindowMouseGrab = Module['_SDL_SetWindowMouseGrab'] = makeInvalidEarlyAccess('_SDL_SetWindowMouseGrab');
var _SDL_GetWindowKeyboardGrab = Module['_SDL_GetWindowKeyboardGrab'] = makeInvalidEarlyAccess('_SDL_GetWindowKeyboardGrab');
var _SDL_GetWindowMouseGrab = Module['_SDL_GetWindowMouseGrab'] = makeInvalidEarlyAccess('_SDL_GetWindowMouseGrab');
var _SDL_GetGrabbedWindow = Module['_SDL_GetGrabbedWindow'] = makeInvalidEarlyAccess('_SDL_GetGrabbedWindow');
var _SDL_SetWindowMouseRect = Module['_SDL_SetWindowMouseRect'] = makeInvalidEarlyAccess('_SDL_SetWindowMouseRect');
var _SDL_SetWindowRelativeMouseMode = Module['_SDL_SetWindowRelativeMouseMode'] = makeInvalidEarlyAccess('_SDL_SetWindowRelativeMouseMode');
var _SDL_GetWindowRelativeMouseMode = Module['_SDL_GetWindowRelativeMouseMode'] = makeInvalidEarlyAccess('_SDL_GetWindowRelativeMouseMode');
var _SDL_FlashWindow = Module['_SDL_FlashWindow'] = makeInvalidEarlyAccess('_SDL_FlashWindow');
var _SDL_ScreenSaverEnabled = Module['_SDL_ScreenSaverEnabled'] = makeInvalidEarlyAccess('_SDL_ScreenSaverEnabled');
var _SDL_EnableScreenSaver = Module['_SDL_EnableScreenSaver'] = makeInvalidEarlyAccess('_SDL_EnableScreenSaver');
var _SDL_EGL_GetProcAddress = Module['_SDL_EGL_GetProcAddress'] = makeInvalidEarlyAccess('_SDL_EGL_GetProcAddress');
var _SDL_EGL_SetAttributeCallbacks = Module['_SDL_EGL_SetAttributeCallbacks'] = makeInvalidEarlyAccess('_SDL_EGL_SetAttributeCallbacks');
var _SDL_GL_GetCurrentWindow = Module['_SDL_GL_GetCurrentWindow'] = makeInvalidEarlyAccess('_SDL_GL_GetCurrentWindow');
var _SDL_EGL_GetCurrentDisplay = Module['_SDL_EGL_GetCurrentDisplay'] = makeInvalidEarlyAccess('_SDL_EGL_GetCurrentDisplay');
var _SDL_EGL_GetCurrentConfig = Module['_SDL_EGL_GetCurrentConfig'] = makeInvalidEarlyAccess('_SDL_EGL_GetCurrentConfig');
var _SDL_EGL_GetWindowSurface = Module['_SDL_EGL_GetWindowSurface'] = makeInvalidEarlyAccess('_SDL_EGL_GetWindowSurface');
var _SDL_StartTextInput = Module['_SDL_StartTextInput'] = makeInvalidEarlyAccess('_SDL_StartTextInput');
var _SDL_StartTextInputWithProperties = Module['_SDL_StartTextInputWithProperties'] = makeInvalidEarlyAccess('_SDL_StartTextInputWithProperties');
var _SDL_ScreenKeyboardShown = Module['_SDL_ScreenKeyboardShown'] = makeInvalidEarlyAccess('_SDL_ScreenKeyboardShown');
var _SDL_StopTextInput = Module['_SDL_StopTextInput'] = makeInvalidEarlyAccess('_SDL_StopTextInput');
var _SDL_SetTextInputArea = Module['_SDL_SetTextInputArea'] = makeInvalidEarlyAccess('_SDL_SetTextInputArea');
var _SDL_GetTextInputArea = Module['_SDL_GetTextInputArea'] = makeInvalidEarlyAccess('_SDL_GetTextInputArea');
var _SDL_ClearComposition = Module['_SDL_ClearComposition'] = makeInvalidEarlyAccess('_SDL_ClearComposition');
var _SDL_HasScreenKeyboardSupport = Module['_SDL_HasScreenKeyboardSupport'] = makeInvalidEarlyAccess('_SDL_HasScreenKeyboardSupport');
var _SDL_ShowSimpleMessageBox = Module['_SDL_ShowSimpleMessageBox'] = makeInvalidEarlyAccess('_SDL_ShowSimpleMessageBox');
var _SDL_ShowWindowSystemMenu = Module['_SDL_ShowWindowSystemMenu'] = makeInvalidEarlyAccess('_SDL_ShowWindowSystemMenu');
var _SDL_SetWindowHitTest = Module['_SDL_SetWindowHitTest'] = makeInvalidEarlyAccess('_SDL_SetWindowHitTest');
var _SDL_SetWindowShape = Module['_SDL_SetWindowShape'] = makeInvalidEarlyAccess('_SDL_SetWindowShape');
var _SDL_OnApplicationWillTerminate = Module['_SDL_OnApplicationWillTerminate'] = makeInvalidEarlyAccess('_SDL_OnApplicationWillTerminate');
var _SDL_OnApplicationDidReceiveMemoryWarning = Module['_SDL_OnApplicationDidReceiveMemoryWarning'] = makeInvalidEarlyAccess('_SDL_OnApplicationDidReceiveMemoryWarning');
var _SDL_OnApplicationWillEnterBackground = Module['_SDL_OnApplicationWillEnterBackground'] = makeInvalidEarlyAccess('_SDL_OnApplicationWillEnterBackground');
var _SDL_OnApplicationDidEnterBackground = Module['_SDL_OnApplicationDidEnterBackground'] = makeInvalidEarlyAccess('_SDL_OnApplicationDidEnterBackground');
var _SDL_OnApplicationWillEnterForeground = Module['_SDL_OnApplicationWillEnterForeground'] = makeInvalidEarlyAccess('_SDL_OnApplicationWillEnterForeground');
var _SDL_OnApplicationDidEnterForeground = Module['_SDL_OnApplicationDidEnterForeground'] = makeInvalidEarlyAccess('_SDL_OnApplicationDidEnterForeground');
var _SDL_Vulkan_GetVkGetInstanceProcAddr = Module['_SDL_Vulkan_GetVkGetInstanceProcAddr'] = makeInvalidEarlyAccess('_SDL_Vulkan_GetVkGetInstanceProcAddr');
var _SDL_Vulkan_GetInstanceExtensions = Module['_SDL_Vulkan_GetInstanceExtensions'] = makeInvalidEarlyAccess('_SDL_Vulkan_GetInstanceExtensions');
var _SDL_Vulkan_CreateSurface = Module['_SDL_Vulkan_CreateSurface'] = makeInvalidEarlyAccess('_SDL_Vulkan_CreateSurface');
var _SDL_Vulkan_DestroySurface = Module['_SDL_Vulkan_DestroySurface'] = makeInvalidEarlyAccess('_SDL_Vulkan_DestroySurface');
var _SDL_Vulkan_GetPresentationSupport = Module['_SDL_Vulkan_GetPresentationSupport'] = makeInvalidEarlyAccess('_SDL_Vulkan_GetPresentationSupport');
var _SDL_Metal_CreateView = Module['_SDL_Metal_CreateView'] = makeInvalidEarlyAccess('_SDL_Metal_CreateView');
var _SDL_Metal_DestroyView = Module['_SDL_Metal_DestroyView'] = makeInvalidEarlyAccess('_SDL_Metal_DestroyView');
var _SDL_Metal_GetLayer = Module['_SDL_Metal_GetLayer'] = makeInvalidEarlyAccess('_SDL_Metal_GetLayer');
var _SDL_GetDXGIOutputInfo = Module['_SDL_GetDXGIOutputInfo'] = makeInvalidEarlyAccess('_SDL_GetDXGIOutputInfo');
var _SDL_GetDirect3D9AdapterIndex = Module['_SDL_GetDirect3D9AdapterIndex'] = makeInvalidEarlyAccess('_SDL_GetDirect3D9AdapterIndex');
var _SDL_GetGDKTaskQueue = Module['_SDL_GetGDKTaskQueue'] = makeInvalidEarlyAccess('_SDL_GetGDKTaskQueue');
var _SDL_OnApplicationDidChangeStatusBarOrientation = Module['_SDL_OnApplicationDidChangeStatusBarOrientation'] = makeInvalidEarlyAccess('_SDL_OnApplicationDidChangeStatusBarOrientation');
var _SDL_SetiOSAnimationCallback = Module['_SDL_SetiOSAnimationCallback'] = makeInvalidEarlyAccess('_SDL_SetiOSAnimationCallback');
var _SDL_SetiOSEventPump = Module['_SDL_SetiOSEventPump'] = makeInvalidEarlyAccess('_SDL_SetiOSEventPump');
var _SDL_EnterAppMainCallbacks = Module['_SDL_EnterAppMainCallbacks'] = makeInvalidEarlyAccess('_SDL_EnterAppMainCallbacks');
var _SDL_RunApp = Module['_SDL_RunApp'] = makeInvalidEarlyAccess('_SDL_RunApp');
var _SDL_GetCurrentTime = Module['_SDL_GetCurrentTime'] = makeInvalidEarlyAccess('_SDL_GetCurrentTime');
var _SDL_TimeToDateTime = Module['_SDL_TimeToDateTime'] = makeInvalidEarlyAccess('_SDL_TimeToDateTime');
var _Emscripten_HandlePointerEnter = Module['_Emscripten_HandlePointerEnter'] = makeInvalidEarlyAccess('_Emscripten_HandlePointerEnter');
var _Emscripten_HandlePointerLeave = Module['_Emscripten_HandlePointerLeave'] = makeInvalidEarlyAccess('_Emscripten_HandlePointerLeave');
var _Emscripten_HandlePointerGeneric = Module['_Emscripten_HandlePointerGeneric'] = makeInvalidEarlyAccess('_Emscripten_HandlePointerGeneric');
var _Emscripten_SendDragEvent = Module['_Emscripten_SendDragEvent'] = makeInvalidEarlyAccess('_Emscripten_SendDragEvent');
var _Emscripten_SendDragCompleteEvent = Module['_Emscripten_SendDragCompleteEvent'] = makeInvalidEarlyAccess('_Emscripten_SendDragCompleteEvent');
var _Emscripten_SendDragTextEvent = Module['_Emscripten_SendDragTextEvent'] = makeInvalidEarlyAccess('_Emscripten_SendDragTextEvent');
var _Emscripten_SendDragFileEvent = Module['_Emscripten_SendDragFileEvent'] = makeInvalidEarlyAccess('_Emscripten_SendDragFileEvent');
var _Emscripten_SendSystemThemeChangedEvent = Module['_Emscripten_SendSystemThemeChangedEvent'] = makeInvalidEarlyAccess('_Emscripten_SendSystemThemeChangedEvent');
var _requestFullscreenThroughSDL = Module['_requestFullscreenThroughSDL'] = makeInvalidEarlyAccess('_requestFullscreenThroughSDL');
var _SDL_ShowFileDialogWithProperties = Module['_SDL_ShowFileDialogWithProperties'] = makeInvalidEarlyAccess('_SDL_ShowFileDialogWithProperties');
var _SDL_ShowOpenFileDialog = Module['_SDL_ShowOpenFileDialog'] = makeInvalidEarlyAccess('_SDL_ShowOpenFileDialog');
var _SDL_ShowSaveFileDialog = Module['_SDL_ShowSaveFileDialog'] = makeInvalidEarlyAccess('_SDL_ShowSaveFileDialog');
var _SDL_ShowOpenFolderDialog = Module['_SDL_ShowOpenFolderDialog'] = makeInvalidEarlyAccess('_SDL_ShowOpenFolderDialog');
var _SDL_CreateProcessWithProperties = Module['_SDL_CreateProcessWithProperties'] = makeInvalidEarlyAccess('_SDL_CreateProcessWithProperties');
var _SDL_ReadProcess = Module['_SDL_ReadProcess'] = makeInvalidEarlyAccess('_SDL_ReadProcess');
var _SDL_DestroyProcess = Module['_SDL_DestroyProcess'] = makeInvalidEarlyAccess('_SDL_DestroyProcess');
var _SDL_WaitProcess = Module['_SDL_WaitProcess'] = makeInvalidEarlyAccess('_SDL_WaitProcess');
var _SDL_CreateProcess = Module['_SDL_CreateProcess'] = makeInvalidEarlyAccess('_SDL_CreateProcess');
var _SDL_GetProcessProperties = Module['_SDL_GetProcessProperties'] = makeInvalidEarlyAccess('_SDL_GetProcessProperties');
var _SDL_GetProcessInput = Module['_SDL_GetProcessInput'] = makeInvalidEarlyAccess('_SDL_GetProcessInput');
var _SDL_GetProcessOutput = Module['_SDL_GetProcessOutput'] = makeInvalidEarlyAccess('_SDL_GetProcessOutput');
var _SDL_KillProcess = Module['_SDL_KillProcess'] = makeInvalidEarlyAccess('_SDL_KillProcess');
var _SDL_DestroyTray = Module['_SDL_DestroyTray'] = makeInvalidEarlyAccess('_SDL_DestroyTray');
var _SDL_CreateTray = Module['_SDL_CreateTray'] = makeInvalidEarlyAccess('_SDL_CreateTray');
var _SDL_SetTrayIcon = Module['_SDL_SetTrayIcon'] = makeInvalidEarlyAccess('_SDL_SetTrayIcon');
var _SDL_SetTrayTooltip = Module['_SDL_SetTrayTooltip'] = makeInvalidEarlyAccess('_SDL_SetTrayTooltip');
var _SDL_CreateTrayMenu = Module['_SDL_CreateTrayMenu'] = makeInvalidEarlyAccess('_SDL_CreateTrayMenu');
var _SDL_GetTrayMenu = Module['_SDL_GetTrayMenu'] = makeInvalidEarlyAccess('_SDL_GetTrayMenu');
var _SDL_CreateTraySubmenu = Module['_SDL_CreateTraySubmenu'] = makeInvalidEarlyAccess('_SDL_CreateTraySubmenu');
var _SDL_GetTraySubmenu = Module['_SDL_GetTraySubmenu'] = makeInvalidEarlyAccess('_SDL_GetTraySubmenu');
var _SDL_GetTrayEntries = Module['_SDL_GetTrayEntries'] = makeInvalidEarlyAccess('_SDL_GetTrayEntries');
var _SDL_RemoveTrayEntry = Module['_SDL_RemoveTrayEntry'] = makeInvalidEarlyAccess('_SDL_RemoveTrayEntry');
var _SDL_InsertTrayEntryAt = Module['_SDL_InsertTrayEntryAt'] = makeInvalidEarlyAccess('_SDL_InsertTrayEntryAt');
var _SDL_SetTrayEntryLabel = Module['_SDL_SetTrayEntryLabel'] = makeInvalidEarlyAccess('_SDL_SetTrayEntryLabel');
var _SDL_GetTrayEntryLabel = Module['_SDL_GetTrayEntryLabel'] = makeInvalidEarlyAccess('_SDL_GetTrayEntryLabel');
var _SDL_SetTrayEntryChecked = Module['_SDL_SetTrayEntryChecked'] = makeInvalidEarlyAccess('_SDL_SetTrayEntryChecked');
var _SDL_GetTrayEntryChecked = Module['_SDL_GetTrayEntryChecked'] = makeInvalidEarlyAccess('_SDL_GetTrayEntryChecked');
var _SDL_SetTrayEntryEnabled = Module['_SDL_SetTrayEntryEnabled'] = makeInvalidEarlyAccess('_SDL_SetTrayEntryEnabled');
var _SDL_GetTrayEntryEnabled = Module['_SDL_GetTrayEntryEnabled'] = makeInvalidEarlyAccess('_SDL_GetTrayEntryEnabled');
var _SDL_SetTrayEntryCallback = Module['_SDL_SetTrayEntryCallback'] = makeInvalidEarlyAccess('_SDL_SetTrayEntryCallback');
var _SDL_ClickTrayEntry = Module['_SDL_ClickTrayEntry'] = makeInvalidEarlyAccess('_SDL_ClickTrayEntry');
var _SDL_GetTrayEntryParent = Module['_SDL_GetTrayEntryParent'] = makeInvalidEarlyAccess('_SDL_GetTrayEntryParent');
var _SDL_GetTrayMenuParentEntry = Module['_SDL_GetTrayMenuParentEntry'] = makeInvalidEarlyAccess('_SDL_GetTrayMenuParentEntry');
var _SDL_GetTrayMenuParentTray = Module['_SDL_GetTrayMenuParentTray'] = makeInvalidEarlyAccess('_SDL_GetTrayMenuParentTray');
var _SDL_TryLockMutex = Module['_SDL_TryLockMutex'] = makeInvalidEarlyAccess('_SDL_TryLockMutex');
var _SDL_TryLockRWLockForReading = Module['_SDL_TryLockRWLockForReading'] = makeInvalidEarlyAccess('_SDL_TryLockRWLockForReading');
var _SDL_TryLockRWLockForWriting = Module['_SDL_TryLockRWLockForWriting'] = makeInvalidEarlyAccess('_SDL_TryLockRWLockForWriting');
var _SDL_GetSemaphoreValue = Module['_SDL_GetSemaphoreValue'] = makeInvalidEarlyAccess('_SDL_GetSemaphoreValue');
var _emscripten_stack_get_end = Module['_emscripten_stack_get_end'] = makeInvalidEarlyAccess('_emscripten_stack_get_end');
var _emscripten_stack_get_base = Module['_emscripten_stack_get_base'] = makeInvalidEarlyAccess('_emscripten_stack_get_base');
var _memcpy = Module['_memcpy'] = makeInvalidEarlyAccess('_memcpy');
var __emscripten_memcpy_bulkmem = Module['__emscripten_memcpy_bulkmem'] = makeInvalidEarlyAccess('__emscripten_memcpy_bulkmem');
var __emscripten_memset_bulkmem = Module['__emscripten_memset_bulkmem'] = makeInvalidEarlyAccess('__emscripten_memset_bulkmem');
var _emscripten_builtin_memalign = Module['_emscripten_builtin_memalign'] = makeInvalidEarlyAccess('_emscripten_builtin_memalign');
var _emscripten_stack_get_current = Module['_emscripten_stack_get_current'] = makeInvalidEarlyAccess('_emscripten_stack_get_current');
var _htons = Module['_htons'] = makeInvalidEarlyAccess('_htons');
var _ntohs = Module['_ntohs'] = makeInvalidEarlyAccess('_ntohs');
var _htonl = Module['_htonl'] = makeInvalidEarlyAccess('_htonl');
var __emscripten_timeout = Module['__emscripten_timeout'] = makeInvalidEarlyAccess('__emscripten_timeout');
var _setThrew = Module['_setThrew'] = makeInvalidEarlyAccess('_setThrew');
var __emscripten_tempret_set = Module['__emscripten_tempret_set'] = makeInvalidEarlyAccess('__emscripten_tempret_set');
var __emscripten_tempret_get = Module['__emscripten_tempret_get'] = makeInvalidEarlyAccess('__emscripten_tempret_get');
var ___get_temp_ret = Module['___get_temp_ret'] = makeInvalidEarlyAccess('___get_temp_ret');
var ___set_temp_ret = Module['___set_temp_ret'] = makeInvalidEarlyAccess('___set_temp_ret');
var ___emutls_get_address = Module['___emutls_get_address'] = makeInvalidEarlyAccess('___emutls_get_address');
var _emscripten_stack_init = Module['_emscripten_stack_init'] = makeInvalidEarlyAccess('_emscripten_stack_init');
var _emscripten_stack_set_limits = Module['_emscripten_stack_set_limits'] = makeInvalidEarlyAccess('_emscripten_stack_set_limits');
var _emscripten_stack_get_free = Module['_emscripten_stack_get_free'] = makeInvalidEarlyAccess('_emscripten_stack_get_free');
var __emscripten_stack_restore = Module['__emscripten_stack_restore'] = makeInvalidEarlyAccess('__emscripten_stack_restore');
var __emscripten_stack_alloc = Module['__emscripten_stack_alloc'] = makeInvalidEarlyAccess('__emscripten_stack_alloc');
var __ZNSt8bad_castD2Ev = Module['__ZNSt8bad_castD2Ev'] = makeInvalidEarlyAccess('__ZNSt8bad_castD2Ev');
var __ZdlPvm = Module['__ZdlPvm'] = makeInvalidEarlyAccess('__ZdlPvm');
var __Znwm = Module['__Znwm'] = makeInvalidEarlyAccess('__Znwm');
var __ZnamSt11align_val_t = Module['__ZnamSt11align_val_t'] = makeInvalidEarlyAccess('__ZnamSt11align_val_t');
var __ZdaPvSt11align_val_t = Module['__ZdaPvSt11align_val_t'] = makeInvalidEarlyAccess('__ZdaPvSt11align_val_t');
var __ZNSt13runtime_errorD2Ev = Module['__ZNSt13runtime_errorD2Ev'] = makeInvalidEarlyAccess('__ZNSt13runtime_errorD2Ev');
var __ZNKSt13runtime_error4whatEv = Module['__ZNKSt13runtime_error4whatEv'] = makeInvalidEarlyAccess('__ZNKSt13runtime_error4whatEv');
var __ZnwmSt11align_val_t = Module['__ZnwmSt11align_val_t'] = makeInvalidEarlyAccess('__ZnwmSt11align_val_t');
var __ZdlPvmSt11align_val_t = Module['__ZdlPvmSt11align_val_t'] = makeInvalidEarlyAccess('__ZdlPvmSt11align_val_t');
var ___cxa_pure_virtual = Module['___cxa_pure_virtual'] = makeInvalidEarlyAccess('___cxa_pure_virtual');
var ___cxa_uncaught_exceptions = Module['___cxa_uncaught_exceptions'] = makeInvalidEarlyAccess('___cxa_uncaught_exceptions');
var ___cxa_decrement_exception_refcount = Module['___cxa_decrement_exception_refcount'] = makeInvalidEarlyAccess('___cxa_decrement_exception_refcount');
var ___cxa_increment_exception_refcount = Module['___cxa_increment_exception_refcount'] = makeInvalidEarlyAccess('___cxa_increment_exception_refcount');
var ___cxa_current_primary_exception = Module['___cxa_current_primary_exception'] = makeInvalidEarlyAccess('___cxa_current_primary_exception');
var __ZSt9terminatev = Module['__ZSt9terminatev'] = makeInvalidEarlyAccess('__ZSt9terminatev');
var ___cxa_rethrow_primary_exception = Module['___cxa_rethrow_primary_exception'] = makeInvalidEarlyAccess('___cxa_rethrow_primary_exception');
var __ZNSt9exceptionD2Ev = Module['__ZNSt9exceptionD2Ev'] = makeInvalidEarlyAccess('__ZNSt9exceptionD2Ev');
var __ZNSt11logic_errorD2Ev = Module['__ZNSt11logic_errorD2Ev'] = makeInvalidEarlyAccess('__ZNSt11logic_errorD2Ev');
var __ZNKSt11logic_error4whatEv = Module['__ZNKSt11logic_error4whatEv'] = makeInvalidEarlyAccess('__ZNKSt11logic_error4whatEv');
var __ZdaPv = Module['__ZdaPv'] = makeInvalidEarlyAccess('__ZdaPv');
var __Znam = Module['__Znam'] = makeInvalidEarlyAccess('__Znam');
var __ZSt15get_new_handlerv = Module['__ZSt15get_new_handlerv'] = makeInvalidEarlyAccess('__ZSt15get_new_handlerv');
var __ZdlPv = Module['__ZdlPv'] = makeInvalidEarlyAccess('__ZdlPv');
var __ZdaPvm = Module['__ZdaPvm'] = makeInvalidEarlyAccess('__ZdaPvm');
var __ZdlPvSt11align_val_t = Module['__ZdlPvSt11align_val_t'] = makeInvalidEarlyAccess('__ZdlPvSt11align_val_t');
var __ZdaPvmSt11align_val_t = Module['__ZdaPvmSt11align_val_t'] = makeInvalidEarlyAccess('__ZdaPvmSt11align_val_t');
var ___dynamic_cast = Module['___dynamic_cast'] = makeInvalidEarlyAccess('___dynamic_cast');
var ___cxa_bad_cast = Module['___cxa_bad_cast'] = makeInvalidEarlyAccess('___cxa_bad_cast');
var ___cxa_bad_typeid = Module['___cxa_bad_typeid'] = makeInvalidEarlyAccess('___cxa_bad_typeid');
var ___cxa_throw_bad_array_new_length = Module['___cxa_throw_bad_array_new_length'] = makeInvalidEarlyAccess('___cxa_throw_bad_array_new_length');
var __ZSt14set_unexpectedPFvvE = Module['__ZSt14set_unexpectedPFvvE'] = makeInvalidEarlyAccess('__ZSt14set_unexpectedPFvvE');
var __ZSt13set_terminatePFvvE = Module['__ZSt13set_terminatePFvvE'] = makeInvalidEarlyAccess('__ZSt13set_terminatePFvvE');
var __ZSt15set_new_handlerPFvvE = Module['__ZSt15set_new_handlerPFvvE'] = makeInvalidEarlyAccess('__ZSt15set_new_handlerPFvvE');
var ___cxa_demangle = Module['___cxa_demangle'] = makeInvalidEarlyAccess('___cxa_demangle');
var ___cxa_guard_acquire = Module['___cxa_guard_acquire'] = makeInvalidEarlyAccess('___cxa_guard_acquire');
var ___cxa_guard_release = Module['___cxa_guard_release'] = makeInvalidEarlyAccess('___cxa_guard_release');
var ___cxa_guard_abort = Module['___cxa_guard_abort'] = makeInvalidEarlyAccess('___cxa_guard_abort');
var __ZSt14get_unexpectedv = Module['__ZSt14get_unexpectedv'] = makeInvalidEarlyAccess('__ZSt14get_unexpectedv');
var __ZSt10unexpectedv = Module['__ZSt10unexpectedv'] = makeInvalidEarlyAccess('__ZSt10unexpectedv');
var __ZSt13get_terminatev = Module['__ZSt13get_terminatev'] = makeInvalidEarlyAccess('__ZSt13get_terminatev');
var ___cxa_uncaught_exception = Module['___cxa_uncaught_exception'] = makeInvalidEarlyAccess('___cxa_uncaught_exception');
var ___cxa_allocate_exception = Module['___cxa_allocate_exception'] = makeInvalidEarlyAccess('___cxa_allocate_exception');
var ___cxa_free_exception = Module['___cxa_free_exception'] = makeInvalidEarlyAccess('___cxa_free_exception');
var ___cxa_init_primary_exception = Module['___cxa_init_primary_exception'] = makeInvalidEarlyAccess('___cxa_init_primary_exception');
var ___cxa_thread_atexit = Module['___cxa_thread_atexit'] = makeInvalidEarlyAccess('___cxa_thread_atexit');
var ___cxa_deleted_virtual = Module['___cxa_deleted_virtual'] = makeInvalidEarlyAccess('___cxa_deleted_virtual');
var __ZNSt9type_infoD2Ev = Module['__ZNSt9type_infoD2Ev'] = makeInvalidEarlyAccess('__ZNSt9type_infoD2Ev');
var ___cxa_can_catch = Module['___cxa_can_catch'] = makeInvalidEarlyAccess('___cxa_can_catch');
var ___cxa_get_exception_ptr = Module['___cxa_get_exception_ptr'] = makeInvalidEarlyAccess('___cxa_get_exception_ptr');
var __ZNSt9exceptionD0Ev = Module['__ZNSt9exceptionD0Ev'] = makeInvalidEarlyAccess('__ZNSt9exceptionD0Ev');
var __ZNSt9exceptionD1Ev = Module['__ZNSt9exceptionD1Ev'] = makeInvalidEarlyAccess('__ZNSt9exceptionD1Ev');
var __ZNKSt9exception4whatEv = Module['__ZNKSt9exception4whatEv'] = makeInvalidEarlyAccess('__ZNKSt9exception4whatEv');
var __ZNSt13bad_exceptionD0Ev = Module['__ZNSt13bad_exceptionD0Ev'] = makeInvalidEarlyAccess('__ZNSt13bad_exceptionD0Ev');
var __ZNSt13bad_exceptionD1Ev = Module['__ZNSt13bad_exceptionD1Ev'] = makeInvalidEarlyAccess('__ZNSt13bad_exceptionD1Ev');
var __ZNKSt13bad_exception4whatEv = Module['__ZNKSt13bad_exception4whatEv'] = makeInvalidEarlyAccess('__ZNKSt13bad_exception4whatEv');
var __ZNSt9bad_allocC2Ev = Module['__ZNSt9bad_allocC2Ev'] = makeInvalidEarlyAccess('__ZNSt9bad_allocC2Ev');
var __ZNSt9bad_allocD0Ev = Module['__ZNSt9bad_allocD0Ev'] = makeInvalidEarlyAccess('__ZNSt9bad_allocD0Ev');
var __ZNSt9bad_allocD1Ev = Module['__ZNSt9bad_allocD1Ev'] = makeInvalidEarlyAccess('__ZNSt9bad_allocD1Ev');
var __ZNKSt9bad_alloc4whatEv = Module['__ZNKSt9bad_alloc4whatEv'] = makeInvalidEarlyAccess('__ZNKSt9bad_alloc4whatEv');
var __ZNSt20bad_array_new_lengthC2Ev = Module['__ZNSt20bad_array_new_lengthC2Ev'] = makeInvalidEarlyAccess('__ZNSt20bad_array_new_lengthC2Ev');
var __ZNSt20bad_array_new_lengthD0Ev = Module['__ZNSt20bad_array_new_lengthD0Ev'] = makeInvalidEarlyAccess('__ZNSt20bad_array_new_lengthD0Ev');
var __ZNSt20bad_array_new_lengthD1Ev = Module['__ZNSt20bad_array_new_lengthD1Ev'] = makeInvalidEarlyAccess('__ZNSt20bad_array_new_lengthD1Ev');
var __ZNKSt20bad_array_new_length4whatEv = Module['__ZNKSt20bad_array_new_length4whatEv'] = makeInvalidEarlyAccess('__ZNKSt20bad_array_new_length4whatEv');
var __ZNSt13bad_exceptionD2Ev = Module['__ZNSt13bad_exceptionD2Ev'] = makeInvalidEarlyAccess('__ZNSt13bad_exceptionD2Ev');
var __ZNSt9bad_allocC1Ev = Module['__ZNSt9bad_allocC1Ev'] = makeInvalidEarlyAccess('__ZNSt9bad_allocC1Ev');
var __ZNSt9bad_allocD2Ev = Module['__ZNSt9bad_allocD2Ev'] = makeInvalidEarlyAccess('__ZNSt9bad_allocD2Ev');
var __ZNSt20bad_array_new_lengthC1Ev = Module['__ZNSt20bad_array_new_lengthC1Ev'] = makeInvalidEarlyAccess('__ZNSt20bad_array_new_lengthC1Ev');
var __ZNSt20bad_array_new_lengthD2Ev = Module['__ZNSt20bad_array_new_lengthD2Ev'] = makeInvalidEarlyAccess('__ZNSt20bad_array_new_lengthD2Ev');
var __ZNSt11logic_errorD0Ev = Module['__ZNSt11logic_errorD0Ev'] = makeInvalidEarlyAccess('__ZNSt11logic_errorD0Ev');
var __ZNSt11logic_errorD1Ev = Module['__ZNSt11logic_errorD1Ev'] = makeInvalidEarlyAccess('__ZNSt11logic_errorD1Ev');
var __ZNSt13runtime_errorD0Ev = Module['__ZNSt13runtime_errorD0Ev'] = makeInvalidEarlyAccess('__ZNSt13runtime_errorD0Ev');
var __ZNSt13runtime_errorD1Ev = Module['__ZNSt13runtime_errorD1Ev'] = makeInvalidEarlyAccess('__ZNSt13runtime_errorD1Ev');
var __ZNSt12domain_errorD0Ev = Module['__ZNSt12domain_errorD0Ev'] = makeInvalidEarlyAccess('__ZNSt12domain_errorD0Ev');
var __ZNSt12domain_errorD1Ev = Module['__ZNSt12domain_errorD1Ev'] = makeInvalidEarlyAccess('__ZNSt12domain_errorD1Ev');
var __ZNSt16invalid_argumentD0Ev = Module['__ZNSt16invalid_argumentD0Ev'] = makeInvalidEarlyAccess('__ZNSt16invalid_argumentD0Ev');
var __ZNSt16invalid_argumentD1Ev = Module['__ZNSt16invalid_argumentD1Ev'] = makeInvalidEarlyAccess('__ZNSt16invalid_argumentD1Ev');
var __ZNSt12length_errorD0Ev = Module['__ZNSt12length_errorD0Ev'] = makeInvalidEarlyAccess('__ZNSt12length_errorD0Ev');
var __ZNSt12length_errorD1Ev = Module['__ZNSt12length_errorD1Ev'] = makeInvalidEarlyAccess('__ZNSt12length_errorD1Ev');
var __ZNSt12out_of_rangeD0Ev = Module['__ZNSt12out_of_rangeD0Ev'] = makeInvalidEarlyAccess('__ZNSt12out_of_rangeD0Ev');
var __ZNSt12out_of_rangeD1Ev = Module['__ZNSt12out_of_rangeD1Ev'] = makeInvalidEarlyAccess('__ZNSt12out_of_rangeD1Ev');
var __ZNSt11range_errorD0Ev = Module['__ZNSt11range_errorD0Ev'] = makeInvalidEarlyAccess('__ZNSt11range_errorD0Ev');
var __ZNSt11range_errorD1Ev = Module['__ZNSt11range_errorD1Ev'] = makeInvalidEarlyAccess('__ZNSt11range_errorD1Ev');
var __ZNSt14overflow_errorD0Ev = Module['__ZNSt14overflow_errorD0Ev'] = makeInvalidEarlyAccess('__ZNSt14overflow_errorD0Ev');
var __ZNSt14overflow_errorD1Ev = Module['__ZNSt14overflow_errorD1Ev'] = makeInvalidEarlyAccess('__ZNSt14overflow_errorD1Ev');
var __ZNSt15underflow_errorD0Ev = Module['__ZNSt15underflow_errorD0Ev'] = makeInvalidEarlyAccess('__ZNSt15underflow_errorD0Ev');
var __ZNSt15underflow_errorD1Ev = Module['__ZNSt15underflow_errorD1Ev'] = makeInvalidEarlyAccess('__ZNSt15underflow_errorD1Ev');
var __ZNSt12domain_errorD2Ev = Module['__ZNSt12domain_errorD2Ev'] = makeInvalidEarlyAccess('__ZNSt12domain_errorD2Ev');
var __ZNSt16invalid_argumentD2Ev = Module['__ZNSt16invalid_argumentD2Ev'] = makeInvalidEarlyAccess('__ZNSt16invalid_argumentD2Ev');
var __ZNSt12length_errorD2Ev = Module['__ZNSt12length_errorD2Ev'] = makeInvalidEarlyAccess('__ZNSt12length_errorD2Ev');
var __ZNSt12out_of_rangeD2Ev = Module['__ZNSt12out_of_rangeD2Ev'] = makeInvalidEarlyAccess('__ZNSt12out_of_rangeD2Ev');
var __ZNSt11range_errorD2Ev = Module['__ZNSt11range_errorD2Ev'] = makeInvalidEarlyAccess('__ZNSt11range_errorD2Ev');
var __ZNSt14overflow_errorD2Ev = Module['__ZNSt14overflow_errorD2Ev'] = makeInvalidEarlyAccess('__ZNSt14overflow_errorD2Ev');
var __ZNSt15underflow_errorD2Ev = Module['__ZNSt15underflow_errorD2Ev'] = makeInvalidEarlyAccess('__ZNSt15underflow_errorD2Ev');
var __ZNSt9type_infoD0Ev = Module['__ZNSt9type_infoD0Ev'] = makeInvalidEarlyAccess('__ZNSt9type_infoD0Ev');
var __ZNSt9type_infoD1Ev = Module['__ZNSt9type_infoD1Ev'] = makeInvalidEarlyAccess('__ZNSt9type_infoD1Ev');
var __ZNSt8bad_castC2Ev = Module['__ZNSt8bad_castC2Ev'] = makeInvalidEarlyAccess('__ZNSt8bad_castC2Ev');
var __ZNSt8bad_castD0Ev = Module['__ZNSt8bad_castD0Ev'] = makeInvalidEarlyAccess('__ZNSt8bad_castD0Ev');
var __ZNSt8bad_castD1Ev = Module['__ZNSt8bad_castD1Ev'] = makeInvalidEarlyAccess('__ZNSt8bad_castD1Ev');
var __ZNKSt8bad_cast4whatEv = Module['__ZNKSt8bad_cast4whatEv'] = makeInvalidEarlyAccess('__ZNKSt8bad_cast4whatEv');
var __ZNSt10bad_typeidC2Ev = Module['__ZNSt10bad_typeidC2Ev'] = makeInvalidEarlyAccess('__ZNSt10bad_typeidC2Ev');
var __ZNSt10bad_typeidD2Ev = Module['__ZNSt10bad_typeidD2Ev'] = makeInvalidEarlyAccess('__ZNSt10bad_typeidD2Ev');
var __ZNSt10bad_typeidD0Ev = Module['__ZNSt10bad_typeidD0Ev'] = makeInvalidEarlyAccess('__ZNSt10bad_typeidD0Ev');
var __ZNSt10bad_typeidD1Ev = Module['__ZNSt10bad_typeidD1Ev'] = makeInvalidEarlyAccess('__ZNSt10bad_typeidD1Ev');
var __ZNKSt10bad_typeid4whatEv = Module['__ZNKSt10bad_typeid4whatEv'] = makeInvalidEarlyAccess('__ZNKSt10bad_typeid4whatEv');
var __ZNSt8bad_castC1Ev = Module['__ZNSt8bad_castC1Ev'] = makeInvalidEarlyAccess('__ZNSt8bad_castC1Ev');
var __ZNSt10bad_typeidC1Ev = Module['__ZNSt10bad_typeidC1Ev'] = makeInvalidEarlyAccess('__ZNSt10bad_typeidC1Ev');

function assignWasmExports(wasmExports) {
  Module['_IMG_Version'] = _IMG_Version = createExportWrapper('IMG_Version', 0);
  Module['_IMG_Load'] = _IMG_Load = createExportWrapper('IMG_Load', 1);
  Module['_SDL_CreateSurface'] = _SDL_CreateSurface = createExportWrapper('SDL_CreateSurface', 3);
  Module['_free'] = _free = createExportWrapper('free', 1);
  Module['_SDL_IOFromFile'] = _SDL_IOFromFile = createExportWrapper('SDL_IOFromFile', 2);
  Module['_SDL_strrchr'] = _SDL_strrchr = createExportWrapper('SDL_strrchr', 2);
  Module['_IMG_LoadTyped_IO'] = _IMG_LoadTyped_IO = createExportWrapper('IMG_LoadTyped_IO', 3);
  Module['_SDL_SetError'] = _SDL_SetError = createExportWrapper('SDL_SetError', 2);
  Module['_SDL_SeekIO'] = _SDL_SeekIO = createExportWrapper('SDL_SeekIO', 3);
  Module['_SDL_CloseIO'] = _SDL_CloseIO = createExportWrapper('SDL_CloseIO', 1);
  Module['_SDL_GetIOProperties'] = _SDL_GetIOProperties = createExportWrapper('SDL_GetIOProperties', 1);
  Module['_SDL_GetPointerProperty'] = _SDL_GetPointerProperty = createExportWrapper('SDL_GetPointerProperty', 3);
  Module['_IMG_isAVIF'] = _IMG_isAVIF = createExportWrapper('IMG_isAVIF', 1);
  Module['_IMG_isCUR'] = _IMG_isCUR = createExportWrapper('IMG_isCUR', 1);
  Module['_IMG_isICO'] = _IMG_isICO = createExportWrapper('IMG_isICO', 1);
  Module['_IMG_isBMP'] = _IMG_isBMP = createExportWrapper('IMG_isBMP', 1);
  Module['_IMG_isGIF'] = _IMG_isGIF = createExportWrapper('IMG_isGIF', 1);
  Module['_IMG_isJPG'] = _IMG_isJPG = createExportWrapper('IMG_isJPG', 1);
  Module['_IMG_isJXL'] = _IMG_isJXL = createExportWrapper('IMG_isJXL', 1);
  Module['_IMG_isLBM'] = _IMG_isLBM = createExportWrapper('IMG_isLBM', 1);
  Module['_IMG_isPCX'] = _IMG_isPCX = createExportWrapper('IMG_isPCX', 1);
  Module['_IMG_isPNG'] = _IMG_isPNG = createExportWrapper('IMG_isPNG', 1);
  Module['_IMG_isPNM'] = _IMG_isPNM = createExportWrapper('IMG_isPNM', 1);
  Module['_IMG_isSVG'] = _IMG_isSVG = createExportWrapper('IMG_isSVG', 1);
  Module['_IMG_isTIF'] = _IMG_isTIF = createExportWrapper('IMG_isTIF', 1);
  Module['_IMG_isXCF'] = _IMG_isXCF = createExportWrapper('IMG_isXCF', 1);
  Module['_IMG_isXPM'] = _IMG_isXPM = createExportWrapper('IMG_isXPM', 1);
  Module['_IMG_isXV'] = _IMG_isXV = createExportWrapper('IMG_isXV', 1);
  Module['_IMG_isWEBP'] = _IMG_isWEBP = createExportWrapper('IMG_isWEBP', 1);
  Module['_IMG_isQOI'] = _IMG_isQOI = createExportWrapper('IMG_isQOI', 1);
  Module['_SDL_strcasecmp'] = _SDL_strcasecmp = createExportWrapper('SDL_strcasecmp', 2);
  Module['_IMG_Load_IO'] = _IMG_Load_IO = createExportWrapper('IMG_Load_IO', 2);
  Module['_IMG_LoadTexture'] = _IMG_LoadTexture = createExportWrapper('IMG_LoadTexture', 2);
  Module['_SDL_CreateTextureFromSurface'] = _SDL_CreateTextureFromSurface = createExportWrapper('SDL_CreateTextureFromSurface', 2);
  Module['_SDL_DestroySurface'] = _SDL_DestroySurface = createExportWrapper('SDL_DestroySurface', 1);
  Module['_IMG_LoadTexture_IO'] = _IMG_LoadTexture_IO = createExportWrapper('IMG_LoadTexture_IO', 3);
  Module['_IMG_LoadTextureTyped_IO'] = _IMG_LoadTextureTyped_IO = createExportWrapper('IMG_LoadTextureTyped_IO', 4);
  Module['_IMG_LoadAnimation'] = _IMG_LoadAnimation = createExportWrapper('IMG_LoadAnimation', 1);
  Module['_IMG_LoadAnimationTyped_IO'] = _IMG_LoadAnimationTyped_IO = createExportWrapper('IMG_LoadAnimationTyped_IO', 3);
  Module['_SDL_malloc'] = _SDL_malloc = createExportWrapper('SDL_malloc', 1);
  Module['_SDL_calloc'] = _SDL_calloc = createExportWrapper('SDL_calloc', 2);
  Module['_SDL_free'] = _SDL_free = createExportWrapper('SDL_free', 1);
  Module['_IMG_LoadAnimation_IO'] = _IMG_LoadAnimation_IO = createExportWrapper('IMG_LoadAnimation_IO', 2);
  Module['_IMG_FreeAnimation'] = _IMG_FreeAnimation = createExportWrapper('IMG_FreeAnimation', 1);
  Module['_IMG_LoadTGA_IO'] = _IMG_LoadTGA_IO = createExportWrapper('IMG_LoadTGA_IO', 1);
  Module['_IMG_LoadAVIF_IO'] = _IMG_LoadAVIF_IO = createExportWrapper('IMG_LoadAVIF_IO', 1);
  Module['_IMG_LoadCUR_IO'] = _IMG_LoadCUR_IO = createExportWrapper('IMG_LoadCUR_IO', 1);
  Module['_IMG_LoadICO_IO'] = _IMG_LoadICO_IO = createExportWrapper('IMG_LoadICO_IO', 1);
  Module['_IMG_LoadBMP_IO'] = _IMG_LoadBMP_IO = createExportWrapper('IMG_LoadBMP_IO', 1);
  Module['_IMG_LoadGIF_IO'] = _IMG_LoadGIF_IO = createExportWrapper('IMG_LoadGIF_IO', 1);
  Module['_IMG_LoadJPG_IO'] = _IMG_LoadJPG_IO = createExportWrapper('IMG_LoadJPG_IO', 1);
  Module['_IMG_LoadJXL_IO'] = _IMG_LoadJXL_IO = createExportWrapper('IMG_LoadJXL_IO', 1);
  Module['_IMG_LoadLBM_IO'] = _IMG_LoadLBM_IO = createExportWrapper('IMG_LoadLBM_IO', 1);
  Module['_IMG_LoadPCX_IO'] = _IMG_LoadPCX_IO = createExportWrapper('IMG_LoadPCX_IO', 1);
  Module['_IMG_LoadPNG_IO'] = _IMG_LoadPNG_IO = createExportWrapper('IMG_LoadPNG_IO', 1);
  Module['_IMG_LoadPNM_IO'] = _IMG_LoadPNM_IO = createExportWrapper('IMG_LoadPNM_IO', 1);
  Module['_IMG_LoadSVG_IO'] = _IMG_LoadSVG_IO = createExportWrapper('IMG_LoadSVG_IO', 1);
  Module['_IMG_LoadTIF_IO'] = _IMG_LoadTIF_IO = createExportWrapper('IMG_LoadTIF_IO', 1);
  Module['_IMG_LoadXCF_IO'] = _IMG_LoadXCF_IO = createExportWrapper('IMG_LoadXCF_IO', 1);
  Module['_IMG_LoadXPM_IO'] = _IMG_LoadXPM_IO = createExportWrapper('IMG_LoadXPM_IO', 1);
  Module['_IMG_LoadXV_IO'] = _IMG_LoadXV_IO = createExportWrapper('IMG_LoadXV_IO', 1);
  Module['_IMG_LoadWEBP_IO'] = _IMG_LoadWEBP_IO = createExportWrapper('IMG_LoadWEBP_IO', 1);
  Module['_IMG_LoadQOI_IO'] = _IMG_LoadQOI_IO = createExportWrapper('IMG_LoadQOI_IO', 1);
  Module['_IMG_LoadGIFAnimation_IO'] = _IMG_LoadGIFAnimation_IO = createExportWrapper('IMG_LoadGIFAnimation_IO', 1);
  Module['_IMG_LoadWEBPAnimation_IO'] = _IMG_LoadWEBPAnimation_IO = createExportWrapper('IMG_LoadWEBPAnimation_IO', 1);
  Module['_IMG_SaveAVIF'] = _IMG_SaveAVIF = createExportWrapper('IMG_SaveAVIF', 3);
  Module['_IMG_SaveAVIF_IO'] = _IMG_SaveAVIF_IO = createExportWrapper('IMG_SaveAVIF_IO', 4);
  Module['_SDL_TellIO'] = _SDL_TellIO = createExportWrapper('SDL_TellIO', 1);
  Module['_SDL_ReadIO'] = _SDL_ReadIO = createExportWrapper('SDL_ReadIO', 3);
  Module['_SDL_strncmp'] = _SDL_strncmp = createExportWrapper('SDL_strncmp', 3);
  Module['_SDL_ReadU16LE'] = _SDL_ReadU16LE = createExportWrapper('SDL_ReadU16LE', 2);
  Module['_SDL_LoadBMP_IO'] = _SDL_LoadBMP_IO = createExportWrapper('SDL_LoadBMP_IO', 2);
  Module['_SDL_ReadU8'] = _SDL_ReadU8 = createExportWrapper('SDL_ReadU8', 2);
  Module['_SDL_ReadU32LE'] = _SDL_ReadU32LE = createExportWrapper('SDL_ReadU32LE', 2);
  Module['_SDL_ReadS32LE'] = _SDL_ReadS32LE = createExportWrapper('SDL_ReadS32LE', 2);
  Module['_SDL_GetSurfaceProperties'] = _SDL_GetSurfaceProperties = createExportWrapper('SDL_GetSurfaceProperties', 1);
  Module['_SDL_SetNumberProperty'] = _SDL_SetNumberProperty = createExportWrapper('SDL_SetNumberProperty', 3);
  Module['_SDL_strcmp'] = _SDL_strcmp = createExportWrapper('SDL_strcmp', 2);
  Module['_SDL_SetSurfaceColorKey'] = _SDL_SetSurfaceColorKey = createExportWrapper('SDL_SetSurfaceColorKey', 3);
  Module['_SDL_realloc'] = _SDL_realloc = createExportWrapper('SDL_realloc', 2);
  Module['_SDL_SurfaceHasColorKey'] = _SDL_SurfaceHasColorKey = createExportWrapper('SDL_SurfaceHasColorKey', 1);
  Module['_SDL_ConvertSurface'] = _SDL_ConvertSurface = createExportWrapper('SDL_ConvertSurface', 2);
  Module['_SDL_MapSurfaceRGBA'] = _SDL_MapSurfaceRGBA = createExportWrapper('SDL_MapSurfaceRGBA', 5);
  Module['_SDL_FillSurfaceRect'] = _SDL_FillSurfaceRect = createExportWrapper('SDL_FillSurfaceRect', 3);
  Module['_SDL_BlitSurface'] = _SDL_BlitSurface = createExportWrapper('SDL_BlitSurface', 4);
  Module['_SDL_DuplicateSurface'] = _SDL_DuplicateSurface = createExportWrapper('SDL_DuplicateSurface', 1);
  Module['_SDL_memcmp'] = _SDL_memcmp = createExportWrapper('SDL_memcmp', 3);
  Module['_SDL_CreateSurfacePalette'] = _SDL_CreateSurfacePalette = createExportWrapper('SDL_CreateSurfacePalette', 1);
  Module['_SDL_Log'] = _SDL_Log = createExportWrapper('SDL_Log', 2);
  Module['_IMG_SaveJPG'] = _IMG_SaveJPG = createExportWrapper('IMG_SaveJPG', 3);
  Module['_IMG_SaveJPG_IO'] = _IMG_SaveJPG_IO = createExportWrapper('IMG_SaveJPG_IO', 4);
  Module['_SDL_floorf'] = _SDL_floorf = createExportWrapper('SDL_floorf', 1);
  Module['_SDL_WriteIO'] = _SDL_WriteIO = createExportWrapper('SDL_WriteIO', 3);
  Module['_IMG_SavePNG'] = _IMG_SavePNG = createExportWrapper('IMG_SavePNG', 2);
  Module['_IMG_SavePNG_IO'] = _IMG_SavePNG_IO = createExportWrapper('IMG_SavePNG_IO', 3);
  Module['_SDL_CreatePalette'] = _SDL_CreatePalette = createExportWrapper('SDL_CreatePalette', 1);
  Module['_SDL_SetSurfacePalette'] = _SDL_SetSurfacePalette = createExportWrapper('SDL_SetSurfacePalette', 2);
  Module['_SDL_DestroyPalette'] = _SDL_DestroyPalette = createExportWrapper('SDL_DestroyPalette', 1);
  Module['_SDL_isspace'] = _SDL_isspace = createExportWrapper('SDL_isspace', 1);
  Module['_SDL_isdigit'] = _SDL_isdigit = createExportWrapper('SDL_isdigit', 1);
  Module['_SDL_LoadFile_IO'] = _SDL_LoadFile_IO = createExportWrapper('SDL_LoadFile_IO', 3);
  Module['_SDL_CreateSurfaceFrom'] = _SDL_CreateSurfaceFrom = createExportWrapper('SDL_CreateSurfaceFrom', 5);
  Module['_SDL_memset'] = _SDL_memset = createExportWrapper('SDL_memset', 3);
  Module['_SDL_SetSurfaceBlendMode'] = _SDL_SetSurfaceBlendMode = createExportWrapper('SDL_SetSurfaceBlendMode', 2);
  Module['_SDL_GetIOStatus'] = _SDL_GetIOStatus = createExportWrapper('SDL_GetIOStatus', 1);
  Module['_SDL_strstr'] = _SDL_strstr = createExportWrapper('SDL_strstr', 2);
  Module['_IMG_LoadSizedSVG_IO'] = _IMG_LoadSizedSVG_IO = createExportWrapper('IMG_LoadSizedSVG_IO', 3);
  Module['_SDL_strchr'] = _SDL_strchr = createExportWrapper('SDL_strchr', 2);
  Module['_SDL_strlen'] = _SDL_strlen = createExportWrapper('SDL_strlen', 1);
  Module['_SDL_ceilf'] = _SDL_ceilf = createExportWrapper('SDL_ceilf', 1);
  Module['_SDL_qsort'] = _SDL_qsort = createExportWrapper('SDL_qsort', 4);
  Module['_SDL_sqrtf'] = _SDL_sqrtf = createExportWrapper('SDL_sqrtf', 1);
  Module['_SDL_fmodf'] = _SDL_fmodf = createExportWrapper('SDL_fmodf', 2);
  Module['_SDL_fabsf'] = _SDL_fabsf = createExportWrapper('SDL_fabsf', 1);
  Module['_SDL_sinf'] = _SDL_sinf = createExportWrapper('SDL_sinf', 1);
  Module['_SDL_cosf'] = _SDL_cosf = createExportWrapper('SDL_cosf', 1);
  Module['_SDL_acosf'] = _SDL_acosf = createExportWrapper('SDL_acosf', 1);
  Module['_SDL_strlcpy'] = _SDL_strlcpy = createExportWrapper('SDL_strlcpy', 3);
  Module['_SDL_strtoll'] = _SDL_strtoll = createExportWrapper('SDL_strtoll', 3);
  Module['_SDL_pow'] = _SDL_pow = createExportWrapper('SDL_pow', 2);
  Module['_SDL_strtol'] = _SDL_strtol = createExportWrapper('SDL_strtol', 3);
  Module['_SDL_tanf'] = _SDL_tanf = createExportWrapper('SDL_tanf', 1);
  Module['_SDL_sscanf'] = _SDL_sscanf = createExportWrapper('SDL_sscanf', 3);
  Module['_SDL_roundf'] = _SDL_roundf = createExportWrapper('SDL_roundf', 1);
  Module['_SDL_fabs'] = _SDL_fabs = createExportWrapper('SDL_fabs', 1);
  Module['_SDL_sqrt'] = _SDL_sqrt = createExportWrapper('SDL_sqrt', 1);
  Module['_SDL_atan2f'] = _SDL_atan2f = createExportWrapper('SDL_atan2f', 2);
  Module['_SDL_ReadU32BE'] = _SDL_ReadU32BE = createExportWrapper('SDL_ReadU32BE', 2);
  Module['_SDL_ReadS32BE'] = _SDL_ReadS32BE = createExportWrapper('SDL_ReadS32BE', 2);
  Module['_SDL_GetIOSize'] = _SDL_GetIOSize = createExportWrapper('SDL_GetIOSize', 1);
  Module['_SDL_strncasecmp'] = _SDL_strncasecmp = createExportWrapper('SDL_strncasecmp', 3);
  Module['_IMG_ReadXPMFromArray'] = _IMG_ReadXPMFromArray = createExportWrapper('IMG_ReadXPMFromArray', 1);
  Module['_IMG_ReadXPMFromArrayToRGB888'] = _IMG_ReadXPMFromArrayToRGB888 = createExportWrapper('IMG_ReadXPMFromArrayToRGB888', 1);
  Module['_SDL_CreateRWLock'] = _SDL_CreateRWLock = createExportWrapper('SDL_CreateRWLock', 0);
  Module['_SDL_DestroyRWLock'] = _SDL_DestroyRWLock = createExportWrapper('SDL_DestroyRWLock', 1);
  Module['_SDL_LockRWLockForWriting'] = _SDL_LockRWLockForWriting = createExportWrapper('SDL_LockRWLockForWriting', 1);
  Module['_SDL_UnlockRWLock'] = _SDL_UnlockRWLock = createExportWrapper('SDL_UnlockRWLock', 1);
  Module['_SDL_LockRWLockForReading'] = _SDL_LockRWLockForReading = createExportWrapper('SDL_LockRWLockForReading', 1);
  Module['_SDL_murmur3_32'] = _SDL_murmur3_32 = createExportWrapper('SDL_murmur3_32', 3);
  Module['_TTF_DestroyGPUTextEngine'] = _TTF_DestroyGPUTextEngine = createExportWrapper('TTF_DestroyGPUTextEngine', 1);
  Module['_TTF_GetFontGeneration'] = _TTF_GetFontGeneration = createExportWrapper('TTF_GetFontGeneration', 1);
  Module['_TTF_GetGlyphImageForIndex'] = _TTF_GetGlyphImageForIndex = createExportWrapper('TTF_GetGlyphImageForIndex', 3);
  Module['_SDL_qsort_r'] = _SDL_qsort_r = createExportWrapper('SDL_qsort_r', 5);
  Module['_SDL_ReleaseGPUTexture'] = _SDL_ReleaseGPUTexture = createExportWrapper('SDL_ReleaseGPUTexture', 2);
  Module['_TTF_CreateGPUTextEngine'] = _TTF_CreateGPUTextEngine = createExportWrapper('TTF_CreateGPUTextEngine', 1);
  Module['_SDL_CreateProperties'] = _SDL_CreateProperties = createExportWrapper('SDL_CreateProperties', 0);
  Module['_SDL_SetPointerProperty'] = _SDL_SetPointerProperty = createExportWrapper('SDL_SetPointerProperty', 3);
  Module['_TTF_CreateGPUTextEngineWithProperties'] = _TTF_CreateGPUTextEngineWithProperties = createExportWrapper('TTF_CreateGPUTextEngineWithProperties', 1);
  Module['_SDL_GetNumberProperty'] = _SDL_GetNumberProperty = createExportWrapper('SDL_GetNumberProperty', 3);
  Module['_TTF_GetGPUTextDrawData'] = _TTF_GetGPUTextDrawData = createExportWrapper('TTF_GetGPUTextDrawData', 1);
  Module['_TTF_UpdateText'] = _TTF_UpdateText = createExportWrapper('TTF_UpdateText', 1);
  Module['_TTF_SetGPUTextEngineWinding'] = _TTF_SetGPUTextEngineWinding = createExportWrapper('TTF_SetGPUTextEngineWinding', 2);
  Module['_TTF_GetGPUTextEngineWinding'] = _TTF_GetGPUTextEngineWinding = createExportWrapper('TTF_GetGPUTextEngineWinding', 1);
  Module['_SDL_CreateGPUTexture'] = _SDL_CreateGPUTexture = createExportWrapper('SDL_CreateGPUTexture', 2);
  Module['_SDL_AcquireGPUCommandBuffer'] = _SDL_AcquireGPUCommandBuffer = createExportWrapper('SDL_AcquireGPUCommandBuffer', 1);
  Module['_SDL_BeginGPURenderPass'] = _SDL_BeginGPURenderPass = createExportWrapper('SDL_BeginGPURenderPass', 4);
  Module['_SDL_EndGPURenderPass'] = _SDL_EndGPURenderPass = createExportWrapper('SDL_EndGPURenderPass', 1);
  Module['_SDL_SubmitGPUCommandBuffer'] = _SDL_SubmitGPUCommandBuffer = createExportWrapper('SDL_SubmitGPUCommandBuffer', 1);
  Module['_SDL_CreateGPUTransferBuffer'] = _SDL_CreateGPUTransferBuffer = createExportWrapper('SDL_CreateGPUTransferBuffer', 2);
  Module['_SDL_MapGPUTransferBuffer'] = _SDL_MapGPUTransferBuffer = createExportWrapper('SDL_MapGPUTransferBuffer', 3);
  Module['_SDL_UnmapGPUTransferBuffer'] = _SDL_UnmapGPUTransferBuffer = createExportWrapper('SDL_UnmapGPUTransferBuffer', 2);
  Module['_SDL_BeginGPUCopyPass'] = _SDL_BeginGPUCopyPass = createExportWrapper('SDL_BeginGPUCopyPass', 1);
  Module['_SDL_UploadToGPUTexture'] = _SDL_UploadToGPUTexture = createExportWrapper('SDL_UploadToGPUTexture', 4);
  Module['_SDL_EndGPUCopyPass'] = _SDL_EndGPUCopyPass = createExportWrapper('SDL_EndGPUCopyPass', 1);
  Module['_SDL_ReleaseGPUTransferBuffer'] = _SDL_ReleaseGPUTransferBuffer = createExportWrapper('SDL_ReleaseGPUTransferBuffer', 2);
  Module['_TTF_CreateRendererTextEngine'] = _TTF_CreateRendererTextEngine = createExportWrapper('TTF_CreateRendererTextEngine', 1);
  Module['_TTF_CreateRendererTextEngineWithProperties'] = _TTF_CreateRendererTextEngineWithProperties = createExportWrapper('TTF_CreateRendererTextEngineWithProperties', 1);
  Module['_TTF_DestroyRendererTextEngine'] = _TTF_DestroyRendererTextEngine = createExportWrapper('TTF_DestroyRendererTextEngine', 1);
  Module['_TTF_DrawRendererText'] = _TTF_DrawRendererText = createExportWrapper('TTF_DrawRendererText', 3);
  Module['_SDL_RenderGeometryRaw'] = _SDL_RenderGeometryRaw = createExportWrapper('SDL_RenderGeometryRaw', 12);
  Module['_SDL_DestroyTexture'] = _SDL_DestroyTexture = createExportWrapper('SDL_DestroyTexture', 1);
  Module['_SDL_CreateTexture'] = _SDL_CreateTexture = createExportWrapper('SDL_CreateTexture', 5);
  Module['_SDL_SetTextureScaleMode'] = _SDL_SetTextureScaleMode = createExportWrapper('SDL_SetTextureScaleMode', 2);
  Module['_SDL_LockTexture'] = _SDL_LockTexture = createExportWrapper('SDL_LockTexture', 4);
  Module['_SDL_UnlockTexture'] = _SDL_UnlockTexture = createExportWrapper('SDL_UnlockTexture', 1);
  Module['_TTF_CreateSurfaceTextEngine'] = _TTF_CreateSurfaceTextEngine = createExportWrapper('TTF_CreateSurfaceTextEngine', 0);
  Module['_TTF_DestroySurfaceTextEngine'] = _TTF_DestroySurfaceTextEngine = createExportWrapper('TTF_DestroySurfaceTextEngine', 1);
  Module['_TTF_DrawSurfaceText'] = _TTF_DrawSurfaceText = createExportWrapper('TTF_DrawSurfaceText', 4);
  Module['_SDL_SetSurfaceColorMod'] = _SDL_SetSurfaceColorMod = createExportWrapper('SDL_SetSurfaceColorMod', 4);
  Module['_SDL_SetSurfaceAlphaMod'] = _SDL_SetSurfaceAlphaMod = createExportWrapper('SDL_SetSurfaceAlphaMod', 2);
  Module['_TTF_Version'] = _TTF_Version = createExportWrapper('TTF_Version', 0);
  Module['_TTF_Init'] = _TTF_Init = createExportWrapper('TTF_Init', 0);
  Module['_SDL_AddAtomicInt'] = _SDL_AddAtomicInt = createExportWrapper('SDL_AddAtomicInt', 2);
  Module['_SDL_ShouldInit'] = _SDL_ShouldInit = createExportWrapper('SDL_ShouldInit', 1);
  Module['_plutosvg_ft_svg_hooks'] = _plutosvg_ft_svg_hooks = createExportWrapper('plutosvg_ft_svg_hooks', 0);
  Module['_SDL_CreateMutex'] = _SDL_CreateMutex = createExportWrapper('SDL_CreateMutex', 0);
  Module['_SDL_SetInitialized'] = _SDL_SetInitialized = createExportWrapper('SDL_SetInitialized', 2);
  Module['_TTF_GetFreeTypeVersion'] = _TTF_GetFreeTypeVersion = createExportWrapper('TTF_GetFreeTypeVersion', 3);
  Module['_SDL_LockMutex'] = _SDL_LockMutex = createExportWrapper('SDL_LockMutex', 1);
  Module['_SDL_UnlockMutex'] = _SDL_UnlockMutex = createExportWrapper('SDL_UnlockMutex', 1);
  Module['_TTF_GetHarfBuzzVersion'] = _TTF_GetHarfBuzzVersion = createExportWrapper('TTF_GetHarfBuzzVersion', 3);
  Module['_TTF_OpenFontWithProperties'] = _TTF_OpenFontWithProperties = createExportWrapper('TTF_OpenFontWithProperties', 1);
  Module['_SDL_GetStringProperty'] = _SDL_GetStringProperty = createExportWrapper('SDL_GetStringProperty', 3);
  Module['_SDL_GetBooleanProperty'] = _SDL_GetBooleanProperty = createExportWrapper('SDL_GetBooleanProperty', 3);
  Module['_SDL_GetFloatProperty'] = _SDL_GetFloatProperty = createExportWrapper('SDL_GetFloatProperty', 3);
  Module['_SDL_strdup'] = _SDL_strdup = createExportWrapper('SDL_strdup', 1);
  Module['_TTF_SetFontKerning'] = _TTF_SetFontKerning = createExportWrapper('TTF_SetFontKerning', 2);
  Module['_TTF_SetFontSizeDPI'] = _TTF_SetFontSizeDPI = createExportWrapper('TTF_SetFontSizeDPI', 4);
  Module['_TTF_CloseFont'] = _TTF_CloseFont = createExportWrapper('TTF_CloseFont', 1);
  Module['_SDL_DestroyProperties'] = _SDL_DestroyProperties = createExportWrapper('SDL_DestroyProperties', 1);
  Module['_TTF_OpenFont'] = _TTF_OpenFont = createExportWrapper('TTF_OpenFont', 2);
  Module['_SDL_SetStringProperty'] = _SDL_SetStringProperty = createExportWrapper('SDL_SetStringProperty', 3);
  Module['_SDL_SetFloatProperty'] = _SDL_SetFloatProperty = createExportWrapper('SDL_SetFloatProperty', 3);
  Module['_TTF_OpenFontIO'] = _TTF_OpenFontIO = createExportWrapper('TTF_OpenFontIO', 3);
  Module['_SDL_SetBooleanProperty'] = _SDL_SetBooleanProperty = createExportWrapper('SDL_SetBooleanProperty', 3);
  Module['_TTF_CopyFont'] = _TTF_CopyFont = createExportWrapper('TTF_CopyFont', 1);
  Module['_TTF_GetFontProperties'] = _TTF_GetFontProperties = createExportWrapper('TTF_GetFontProperties', 1);
  Module['_TTF_AddFallbackFont'] = _TTF_AddFallbackFont = createExportWrapper('TTF_AddFallbackFont', 2);
  Module['_TTF_RemoveFallbackFont'] = _TTF_RemoveFallbackFont = createExportWrapper('TTF_RemoveFallbackFont', 2);
  Module['_TTF_ClearFallbackFonts'] = _TTF_ClearFallbackFonts = createExportWrapper('TTF_ClearFallbackFonts', 1);
  Module['_TTF_FontHasGlyph'] = _TTF_FontHasGlyph = createExportWrapper('TTF_FontHasGlyph', 2);
  Module['_TTF_GetGlyphImage'] = _TTF_GetGlyphImage = createExportWrapper('TTF_GetGlyphImage', 3);
  Module['_TTF_GetGlyphMetrics'] = _TTF_GetGlyphMetrics = createExportWrapper('TTF_GetGlyphMetrics', 7);
  Module['_TTF_GetGlyphKerning'] = _TTF_GetGlyphKerning = createExportWrapper('TTF_GetGlyphKerning', 4);
  Module['_TTF_GetStringSize'] = _TTF_GetStringSize = createExportWrapper('TTF_GetStringSize', 5);
  Module['_TTF_MeasureString'] = _TTF_MeasureString = createExportWrapper('TTF_MeasureString', 6);
  Module['_TTF_RenderText_Solid'] = _TTF_RenderText_Solid = createExportWrapper('TTF_RenderText_Solid', 4);
  Module['_SDL_GetSurfacePalette'] = _SDL_GetSurfacePalette = createExportWrapper('SDL_GetSurfacePalette', 1);
  Module['_TTF_RenderGlyph_Solid'] = _TTF_RenderGlyph_Solid = createExportWrapper('TTF_RenderGlyph_Solid', 3);
  Module['_SDL_UCS4ToUTF8'] = _SDL_UCS4ToUTF8 = createExportWrapper('SDL_UCS4ToUTF8', 2);
  Module['_TTF_RenderText_Shaded'] = _TTF_RenderText_Shaded = createExportWrapper('TTF_RenderText_Shaded', 5);
  Module['_TTF_RenderGlyph_Shaded'] = _TTF_RenderGlyph_Shaded = createExportWrapper('TTF_RenderGlyph_Shaded', 4);
  Module['_TTF_RenderText_Blended'] = _TTF_RenderText_Blended = createExportWrapper('TTF_RenderText_Blended', 4);
  Module['_TTF_RenderGlyph_Blended'] = _TTF_RenderGlyph_Blended = createExportWrapper('TTF_RenderGlyph_Blended', 3);
  Module['_TTF_RenderText_LCD'] = _TTF_RenderText_LCD = createExportWrapper('TTF_RenderText_LCD', 5);
  Module['_TTF_RenderGlyph_LCD'] = _TTF_RenderGlyph_LCD = createExportWrapper('TTF_RenderGlyph_LCD', 4);
  Module['_TTF_GetStringSizeWrapped'] = _TTF_GetStringSizeWrapped = createExportWrapper('TTF_GetStringSizeWrapped', 6);
  Module['_SDL_StepUTF8'] = _SDL_StepUTF8 = createExportWrapper('SDL_StepUTF8', 2);
  Module['_TTF_RenderText_Solid_Wrapped'] = _TTF_RenderText_Solid_Wrapped = createExportWrapper('TTF_RenderText_Solid_Wrapped', 5);
  Module['_SDL_memset4'] = _SDL_memset4 = createExportWrapper('SDL_memset4', 3);
  Module['_TTF_RenderText_Shaded_Wrapped'] = _TTF_RenderText_Shaded_Wrapped = createExportWrapper('TTF_RenderText_Shaded_Wrapped', 6);
  Module['_TTF_RenderText_Blended_Wrapped'] = _TTF_RenderText_Blended_Wrapped = createExportWrapper('TTF_RenderText_Blended_Wrapped', 5);
  Module['_TTF_RenderText_LCD_Wrapped'] = _TTF_RenderText_LCD_Wrapped = createExportWrapper('TTF_RenderText_LCD_Wrapped', 6);
  Module['_TTF_CreateText'] = _TTF_CreateText = createExportWrapper('TTF_CreateText', 4);
  Module['_TTF_GetTextProperties'] = _TTF_GetTextProperties = createExportWrapper('TTF_GetTextProperties', 1);
  Module['_TTF_SetTextEngine'] = _TTF_SetTextEngine = createExportWrapper('TTF_SetTextEngine', 2);
  Module['_TTF_GetTextEngine'] = _TTF_GetTextEngine = createExportWrapper('TTF_GetTextEngine', 1);
  Module['_TTF_SetTextFont'] = _TTF_SetTextFont = createExportWrapper('TTF_SetTextFont', 2);
  Module['_TTF_GetTextFont'] = _TTF_GetTextFont = createExportWrapper('TTF_GetTextFont', 1);
  Module['_TTF_SetTextDirection'] = _TTF_SetTextDirection = createExportWrapper('TTF_SetTextDirection', 2);
  Module['_TTF_GetTextDirection'] = _TTF_GetTextDirection = createExportWrapper('TTF_GetTextDirection', 1);
  Module['_TTF_GetFontDirection'] = _TTF_GetFontDirection = createExportWrapper('TTF_GetFontDirection', 1);
  Module['_TTF_SetTextScript'] = _TTF_SetTextScript = createExportWrapper('TTF_SetTextScript', 2);
  Module['_TTF_GetTextScript'] = _TTF_GetTextScript = createExportWrapper('TTF_GetTextScript', 1);
  Module['_TTF_GetFontScript'] = _TTF_GetFontScript = createExportWrapper('TTF_GetFontScript', 1);
  Module['_TTF_SetTextColor'] = _TTF_SetTextColor = createExportWrapper('TTF_SetTextColor', 5);
  Module['_TTF_SetTextColorFloat'] = _TTF_SetTextColorFloat = createExportWrapper('TTF_SetTextColorFloat', 5);
  Module['_TTF_GetTextColor'] = _TTF_GetTextColor = createExportWrapper('TTF_GetTextColor', 5);
  Module['_TTF_GetTextColorFloat'] = _TTF_GetTextColorFloat = createExportWrapper('TTF_GetTextColorFloat', 5);
  Module['_TTF_SetTextPosition'] = _TTF_SetTextPosition = createExportWrapper('TTF_SetTextPosition', 3);
  Module['_TTF_GetTextPosition'] = _TTF_GetTextPosition = createExportWrapper('TTF_GetTextPosition', 3);
  Module['_TTF_SetTextWrapWidth'] = _TTF_SetTextWrapWidth = createExportWrapper('TTF_SetTextWrapWidth', 2);
  Module['_TTF_GetTextWrapWidth'] = _TTF_GetTextWrapWidth = createExportWrapper('TTF_GetTextWrapWidth', 2);
  Module['_TTF_SetTextWrapWhitespaceVisible'] = _TTF_SetTextWrapWhitespaceVisible = createExportWrapper('TTF_SetTextWrapWhitespaceVisible', 2);
  Module['_TTF_TextWrapWhitespaceVisible'] = _TTF_TextWrapWhitespaceVisible = createExportWrapper('TTF_TextWrapWhitespaceVisible', 1);
  Module['_TTF_SetTextString'] = _TTF_SetTextString = createExportWrapper('TTF_SetTextString', 3);
  Module['_TTF_InsertTextString'] = _TTF_InsertTextString = createExportWrapper('TTF_InsertTextString', 4);
  Module['_TTF_AppendTextString'] = _TTF_AppendTextString = createExportWrapper('TTF_AppendTextString', 3);
  Module['_TTF_DeleteTextString'] = _TTF_DeleteTextString = createExportWrapper('TTF_DeleteTextString', 3);
  Module['_TTF_GetTextSize'] = _TTF_GetTextSize = createExportWrapper('TTF_GetTextSize', 3);
  Module['_SDL_GetRectUnion'] = _SDL_GetRectUnion = createExportWrapper('SDL_GetRectUnion', 3);
  Module['_TTF_GetTextSubString'] = _TTF_GetTextSubString = createExportWrapper('TTF_GetTextSubString', 3);
  Module['_TTF_GetTextSubStringForLine'] = _TTF_GetTextSubStringForLine = createExportWrapper('TTF_GetTextSubStringForLine', 3);
  Module['_TTF_GetTextSubStringsForRange'] = _TTF_GetTextSubStringsForRange = createExportWrapper('TTF_GetTextSubStringsForRange', 4);
  Module['_TTF_GetPreviousTextSubString'] = _TTF_GetPreviousTextSubString = createExportWrapper('TTF_GetPreviousTextSubString', 3);
  Module['_TTF_GetTextSubStringForPoint'] = _TTF_GetTextSubStringForPoint = createExportWrapper('TTF_GetTextSubStringForPoint', 4);
  Module['_SDL_abs'] = _SDL_abs = createExportWrapper('SDL_abs', 1);
  Module['_TTF_GetNextTextSubString'] = _TTF_GetNextTextSubString = createExportWrapper('TTF_GetNextTextSubString', 3);
  Module['_TTF_DestroyText'] = _TTF_DestroyText = createExportWrapper('TTF_DestroyText', 1);
  Module['_TTF_SetFontSize'] = _TTF_SetFontSize = createExportWrapper('TTF_SetFontSize', 2);
  Module['_TTF_GetFontSize'] = _TTF_GetFontSize = createExportWrapper('TTF_GetFontSize', 1);
  Module['_TTF_GetFontDPI'] = _TTF_GetFontDPI = createExportWrapper('TTF_GetFontDPI', 3);
  Module['_TTF_SetFontStyle'] = _TTF_SetFontStyle = createExportWrapper('TTF_SetFontStyle', 2);
  Module['_TTF_GetFontStyle'] = _TTF_GetFontStyle = createExportWrapper('TTF_GetFontStyle', 1);
  Module['_TTF_SetFontOutline'] = _TTF_SetFontOutline = createExportWrapper('TTF_SetFontOutline', 2);
  Module['_TTF_GetFontOutline'] = _TTF_GetFontOutline = createExportWrapper('TTF_GetFontOutline', 1);
  Module['_TTF_SetFontHinting'] = _TTF_SetFontHinting = createExportWrapper('TTF_SetFontHinting', 2);
  Module['_TTF_GetFontHinting'] = _TTF_GetFontHinting = createExportWrapper('TTF_GetFontHinting', 1);
  Module['_TTF_SetFontSDF'] = _TTF_SetFontSDF = createExportWrapper('TTF_SetFontSDF', 2);
  Module['_TTF_GetFontSDF'] = _TTF_GetFontSDF = createExportWrapper('TTF_GetFontSDF', 1);
  Module['_TTF_SetFontWrapAlignment'] = _TTF_SetFontWrapAlignment = createExportWrapper('TTF_SetFontWrapAlignment', 2);
  Module['_TTF_GetFontWrapAlignment'] = _TTF_GetFontWrapAlignment = createExportWrapper('TTF_GetFontWrapAlignment', 1);
  Module['_TTF_GetFontHeight'] = _TTF_GetFontHeight = createExportWrapper('TTF_GetFontHeight', 1);
  Module['_TTF_GetFontAscent'] = _TTF_GetFontAscent = createExportWrapper('TTF_GetFontAscent', 1);
  Module['_TTF_GetFontDescent'] = _TTF_GetFontDescent = createExportWrapper('TTF_GetFontDescent', 1);
  Module['_TTF_SetFontLineSkip'] = _TTF_SetFontLineSkip = createExportWrapper('TTF_SetFontLineSkip', 2);
  Module['_TTF_GetFontLineSkip'] = _TTF_GetFontLineSkip = createExportWrapper('TTF_GetFontLineSkip', 1);
  Module['_TTF_GetFontKerning'] = _TTF_GetFontKerning = createExportWrapper('TTF_GetFontKerning', 1);
  Module['_TTF_GetNumFontFaces'] = _TTF_GetNumFontFaces = createExportWrapper('TTF_GetNumFontFaces', 1);
  Module['_TTF_FontIsFixedWidth'] = _TTF_FontIsFixedWidth = createExportWrapper('TTF_FontIsFixedWidth', 1);
  Module['_TTF_FontIsScalable'] = _TTF_FontIsScalable = createExportWrapper('TTF_FontIsScalable', 1);
  Module['_TTF_GetFontFamilyName'] = _TTF_GetFontFamilyName = createExportWrapper('TTF_GetFontFamilyName', 1);
  Module['_TTF_GetFontStyleName'] = _TTF_GetFontStyleName = createExportWrapper('TTF_GetFontStyleName', 1);
  Module['_TTF_SetFontDirection'] = _TTF_SetFontDirection = createExportWrapper('TTF_SetFontDirection', 2);
  Module['_TTF_StringToTag'] = _TTF_StringToTag = createExportWrapper('TTF_StringToTag', 1);
  Module['_TTF_TagToString'] = _TTF_TagToString = createExportWrapper('TTF_TagToString', 3);
  Module['_TTF_SetFontScript'] = _TTF_SetFontScript = createExportWrapper('TTF_SetFontScript', 2);
  Module['_TTF_GetGlyphScript'] = _TTF_GetGlyphScript = createExportWrapper('TTF_GetGlyphScript', 1);
  Module['_TTF_SetFontLanguage'] = _TTF_SetFontLanguage = createExportWrapper('TTF_SetFontLanguage', 2);
  Module['_TTF_Quit'] = _TTF_Quit = createExportWrapper('TTF_Quit', 0);
  Module['_SDL_ShouldQuit'] = _SDL_ShouldQuit = createExportWrapper('SDL_ShouldQuit', 1);
  Module['_SDL_DestroyMutex'] = _SDL_DestroyMutex = createExportWrapper('SDL_DestroyMutex', 1);
  Module['_TTF_WasInit'] = _TTF_WasInit = createExportWrapper('TTF_WasInit', 0);
  Module['_SDL_GetAtomicInt'] = _SDL_GetAtomicInt = createExportWrapper('SDL_GetAtomicInt', 1);
  Module['_SDL_aligned_alloc'] = _SDL_aligned_alloc = createExportWrapper('SDL_aligned_alloc', 2);
  Module['_SDL_aligned_free'] = _SDL_aligned_free = createExportWrapper('SDL_aligned_free', 1);
  Module['_realloc'] = _realloc = createExportWrapper('realloc', 2);
  Module['_calloc'] = _calloc = createExportWrapper('calloc', 2);
  Module['_malloc'] = _malloc = createExportWrapper('malloc', 1);
  Module['_strerror'] = _strerror = createExportWrapper('strerror', 1);
  Module['_memcmp'] = _memcmp = createExportWrapper('memcmp', 3);
  Module['_setTempRet0'] = _setTempRet0 = createExportWrapper('setTempRet0', 1);
  Module['_getTempRet0'] = _getTempRet0 = createExportWrapper('getTempRet0', 0);
  Module['_plutosvg_version'] = _plutosvg_version = createExportWrapper('plutosvg_version', 0);
  Module['_plutosvg_version_string'] = _plutosvg_version_string = createExportWrapper('plutosvg_version_string', 0);
  Module['_plutosvg_document_destroy'] = _plutosvg_document_destroy = createExportWrapper('plutosvg_document_destroy', 1);
  Module['_plutovg_path_destroy'] = _plutovg_path_destroy = createExportWrapper('plutovg_path_destroy', 1);
  Module['_plutosvg_document_load_from_data'] = _plutosvg_document_load_from_data = createExportWrapper('plutosvg_document_load_from_data', 6);
  Module['_plutovg_path_create'] = _plutovg_path_create = createExportWrapper('plutovg_path_create', 0);
  Module['_plutosvg_document_load_from_file'] = _plutosvg_document_load_from_file = createExportWrapper('plutosvg_document_load_from_file', 3);
  Module['_plutosvg_document_render'] = _plutosvg_document_render = createExportWrapper('plutosvg_document_render', 6);
  Module['_plutovg_canvas_get_matrix'] = _plutovg_canvas_get_matrix = createExportWrapper('plutovg_canvas_get_matrix', 2);
  Module['_plutovg_matrix_translate'] = _plutovg_matrix_translate = createExportWrapper('plutovg_matrix_translate', 3);
  Module['_plutovg_path_reset'] = _plutovg_path_reset = createExportWrapper('plutovg_path_reset', 1);
  Module['_plutovg_path_move_to'] = _plutovg_path_move_to = createExportWrapper('plutovg_path_move_to', 3);
  Module['_plutovg_path_line_to'] = _plutovg_path_line_to = createExportWrapper('plutovg_path_line_to', 3);
  Module['_plutovg_path_add_ellipse'] = _plutovg_path_add_ellipse = createExportWrapper('plutovg_path_add_ellipse', 5);
  Module['_plutovg_path_add_circle'] = _plutovg_path_add_circle = createExportWrapper('plutovg_path_add_circle', 4);
  Module['_plutovg_path_add_round_rect'] = _plutovg_path_add_round_rect = createExportWrapper('plutovg_path_add_round_rect', 7);
  Module['_plutovg_path_close'] = _plutovg_path_close = createExportWrapper('plutovg_path_close', 1);
  Module['_plutovg_path_extents'] = _plutovg_path_extents = createExportWrapper('plutovg_path_extents', 3);
  Module['_plutovg_path_parse'] = _plutovg_path_parse = createExportWrapper('plutovg_path_parse', 3);
  Module['_plutovg_surface_load_from_image_base64'] = _plutovg_surface_load_from_image_base64 = createExportWrapper('plutovg_surface_load_from_image_base64', 2);
  Module['_plutovg_surface_get_width'] = _plutovg_surface_get_width = createExportWrapper('plutovg_surface_get_width', 1);
  Module['_plutovg_surface_get_height'] = _plutovg_surface_get_height = createExportWrapper('plutovg_surface_get_height', 1);
  Module['_plutovg_canvas_set_fill_rule'] = _plutovg_canvas_set_fill_rule = createExportWrapper('plutovg_canvas_set_fill_rule', 2);
  Module['_plutovg_canvas_set_opacity'] = _plutovg_canvas_set_opacity = createExportWrapper('plutovg_canvas_set_opacity', 2);
  Module['_plutovg_canvas_set_matrix'] = _plutovg_canvas_set_matrix = createExportWrapper('plutovg_canvas_set_matrix', 2);
  Module['_plutovg_canvas_translate'] = _plutovg_canvas_translate = createExportWrapper('plutovg_canvas_translate', 3);
  Module['_plutovg_canvas_set_texture'] = _plutovg_canvas_set_texture = createExportWrapper('plutovg_canvas_set_texture', 5);
  Module['_plutovg_canvas_fill_rect'] = _plutovg_canvas_fill_rect = createExportWrapper('plutovg_canvas_fill_rect', 5);
  Module['_plutovg_surface_destroy'] = _plutovg_surface_destroy = createExportWrapper('plutovg_surface_destroy', 1);
  Module['_plutosvg_document_render_to_surface'] = _plutosvg_document_render_to_surface = createExportWrapper('plutosvg_document_render_to_surface', 7);
  Module['_plutosvg_document_extents'] = _plutosvg_document_extents = createExportWrapper('plutosvg_document_extents', 3);
  Module['_plutovg_surface_create'] = _plutovg_surface_create = createExportWrapper('plutovg_surface_create', 2);
  Module['_plutovg_canvas_create'] = _plutovg_canvas_create = createExportWrapper('plutovg_canvas_create', 1);
  Module['_plutovg_canvas_scale'] = _plutovg_canvas_scale = createExportWrapper('plutovg_canvas_scale', 3);
  Module['_plutovg_canvas_destroy'] = _plutovg_canvas_destroy = createExportWrapper('plutovg_canvas_destroy', 1);
  Module['_plutovg_matrix_init_identity'] = _plutovg_matrix_init_identity = createExportWrapper('plutovg_matrix_init_identity', 1);
  Module['_plutosvg_document_get_width'] = _plutosvg_document_get_width = createExportWrapper('plutosvg_document_get_width', 1);
  Module['_plutosvg_document_get_height'] = _plutosvg_document_get_height = createExportWrapper('plutosvg_document_get_height', 1);
  Module['_plutovg_matrix_parse'] = _plutovg_matrix_parse = createExportWrapper('plutovg_matrix_parse', 3);
  Module['_plutovg_matrix_multiply'] = _plutovg_matrix_multiply = createExportWrapper('plutovg_matrix_multiply', 3);
  Module['_plutovg_matrix_scale'] = _plutovg_matrix_scale = createExportWrapper('plutovg_matrix_scale', 3);
  Module['_plutovg_matrix_invert'] = _plutovg_matrix_invert = createExportWrapper('plutovg_matrix_invert', 2);
  Module['_plutovg_matrix_map_rect'] = _plutovg_matrix_map_rect = createExportWrapper('plutovg_matrix_map_rect', 3);
  Module['_plutovg_canvas_fill_path'] = _plutovg_canvas_fill_path = createExportWrapper('plutovg_canvas_fill_path', 2);
  Module['_plutovg_canvas_set_dash_offset'] = _plutovg_canvas_set_dash_offset = createExportWrapper('plutovg_canvas_set_dash_offset', 2);
  Module['_plutovg_canvas_set_dash_array'] = _plutovg_canvas_set_dash_array = createExportWrapper('plutovg_canvas_set_dash_array', 3);
  Module['_plutovg_canvas_set_line_width'] = _plutovg_canvas_set_line_width = createExportWrapper('plutovg_canvas_set_line_width', 2);
  Module['_plutovg_canvas_set_line_cap'] = _plutovg_canvas_set_line_cap = createExportWrapper('plutovg_canvas_set_line_cap', 2);
  Module['_plutovg_canvas_set_line_join'] = _plutovg_canvas_set_line_join = createExportWrapper('plutovg_canvas_set_line_join', 2);
  Module['_plutovg_canvas_set_miter_limit'] = _plutovg_canvas_set_miter_limit = createExportWrapper('plutovg_canvas_set_miter_limit', 2);
  Module['_plutovg_canvas_stroke_path'] = _plutovg_canvas_stroke_path = createExportWrapper('plutovg_canvas_stroke_path', 2);
  Module['_plutovg_color_init_argb32'] = _plutovg_color_init_argb32 = createExportWrapper('plutovg_color_init_argb32', 2);
  Module['_plutovg_canvas_set_color'] = _plutovg_canvas_set_color = createExportWrapper('plutovg_canvas_set_color', 2);
  Module['_plutovg_canvas_set_linear_gradient'] = _plutovg_canvas_set_linear_gradient = createExportWrapper('plutovg_canvas_set_linear_gradient', 9);
  Module['_plutovg_canvas_set_radial_gradient'] = _plutovg_canvas_set_radial_gradient = createExportWrapper('plutovg_canvas_set_radial_gradient', 11);
  Module['_plutovg_color_parse'] = _plutovg_color_parse = createExportWrapper('plutovg_color_parse', 3);
  Module['_plutovg_color_to_argb32'] = _plutovg_color_to_argb32 = createExportWrapper('plutovg_color_to_argb32', 1);
  Module['_plutovg_matrix_init_translate'] = _plutovg_matrix_init_translate = createExportWrapper('plutovg_matrix_init_translate', 3);
  Module['_plutovg_surface_create_for_data'] = _plutovg_surface_create_for_data = createExportWrapper('plutovg_surface_create_for_data', 4);
  Module['_plutovg_canvas_transform'] = _plutovg_canvas_transform = createExportWrapper('plutovg_canvas_transform', 2);
  Module['_plutovg_matrix_init_scale'] = _plutovg_matrix_init_scale = createExportWrapper('plutovg_matrix_init_scale', 3);
  Module['_plutovg_version'] = _plutovg_version = createExportWrapper('plutovg_version', 0);
  Module['_plutovg_version_string'] = _plutovg_version_string = createExportWrapper('plutovg_version_string', 0);
  Module['_plutovg_surface_reference'] = _plutovg_surface_reference = createExportWrapper('plutovg_surface_reference', 1);
  Module['_plutovg_canvas_reference'] = _plutovg_canvas_reference = createExportWrapper('plutovg_canvas_reference', 1);
  Module['_plutovg_paint_destroy'] = _plutovg_paint_destroy = createExportWrapper('plutovg_paint_destroy', 1);
  Module['_plutovg_font_face_destroy'] = _plutovg_font_face_destroy = createExportWrapper('plutovg_font_face_destroy', 1);
  Module['_plutovg_canvas_get_reference_count'] = _plutovg_canvas_get_reference_count = createExportWrapper('plutovg_canvas_get_reference_count', 1);
  Module['_plutovg_canvas_get_surface'] = _plutovg_canvas_get_surface = createExportWrapper('plutovg_canvas_get_surface', 1);
  Module['_plutovg_canvas_save'] = _plutovg_canvas_save = createExportWrapper('plutovg_canvas_save', 1);
  Module['_plutovg_paint_reference'] = _plutovg_paint_reference = createExportWrapper('plutovg_paint_reference', 1);
  Module['_plutovg_font_face_reference'] = _plutovg_font_face_reference = createExportWrapper('plutovg_font_face_reference', 1);
  Module['_plutovg_canvas_restore'] = _plutovg_canvas_restore = createExportWrapper('plutovg_canvas_restore', 1);
  Module['_plutovg_canvas_set_rgb'] = _plutovg_canvas_set_rgb = createExportWrapper('plutovg_canvas_set_rgb', 4);
  Module['_plutovg_color_init_rgba'] = _plutovg_color_init_rgba = createExportWrapper('plutovg_color_init_rgba', 5);
  Module['_plutovg_canvas_set_rgba'] = _plutovg_canvas_set_rgba = createExportWrapper('plutovg_canvas_set_rgba', 5);
  Module['_plutovg_canvas_set_paint'] = _plutovg_canvas_set_paint = createExportWrapper('plutovg_canvas_set_paint', 2);
  Module['_plutovg_paint_create_linear_gradient'] = _plutovg_paint_create_linear_gradient = createExportWrapper('plutovg_paint_create_linear_gradient', 8);
  Module['_plutovg_paint_create_radial_gradient'] = _plutovg_paint_create_radial_gradient = createExportWrapper('plutovg_paint_create_radial_gradient', 10);
  Module['_plutovg_paint_create_texture'] = _plutovg_paint_create_texture = createExportWrapper('plutovg_paint_create_texture', 4);
  Module['_plutovg_canvas_get_paint'] = _plutovg_canvas_get_paint = createExportWrapper('plutovg_canvas_get_paint', 2);
  Module['_plutovg_canvas_set_font'] = _plutovg_canvas_set_font = createExportWrapper('plutovg_canvas_set_font', 3);
  Module['_plutovg_canvas_set_font_face'] = _plutovg_canvas_set_font_face = createExportWrapper('plutovg_canvas_set_font_face', 2);
  Module['_plutovg_canvas_set_font_size'] = _plutovg_canvas_set_font_size = createExportWrapper('plutovg_canvas_set_font_size', 2);
  Module['_plutovg_canvas_get_font_face'] = _plutovg_canvas_get_font_face = createExportWrapper('plutovg_canvas_get_font_face', 1);
  Module['_plutovg_canvas_get_font_size'] = _plutovg_canvas_get_font_size = createExportWrapper('plutovg_canvas_get_font_size', 1);
  Module['_plutovg_canvas_get_fill_rule'] = _plutovg_canvas_get_fill_rule = createExportWrapper('plutovg_canvas_get_fill_rule', 1);
  Module['_plutovg_canvas_set_operator'] = _plutovg_canvas_set_operator = createExportWrapper('plutovg_canvas_set_operator', 2);
  Module['_plutovg_canvas_get_operator'] = _plutovg_canvas_get_operator = createExportWrapper('plutovg_canvas_get_operator', 1);
  Module['_plutovg_canvas_get_opacity'] = _plutovg_canvas_get_opacity = createExportWrapper('plutovg_canvas_get_opacity', 1);
  Module['_plutovg_canvas_get_line_width'] = _plutovg_canvas_get_line_width = createExportWrapper('plutovg_canvas_get_line_width', 1);
  Module['_plutovg_canvas_get_line_cap'] = _plutovg_canvas_get_line_cap = createExportWrapper('plutovg_canvas_get_line_cap', 1);
  Module['_plutovg_canvas_get_line_join'] = _plutovg_canvas_get_line_join = createExportWrapper('plutovg_canvas_get_line_join', 1);
  Module['_plutovg_canvas_get_miter_limit'] = _plutovg_canvas_get_miter_limit = createExportWrapper('plutovg_canvas_get_miter_limit', 1);
  Module['_plutovg_canvas_set_dash'] = _plutovg_canvas_set_dash = createExportWrapper('plutovg_canvas_set_dash', 4);
  Module['_plutovg_canvas_get_dash_offset'] = _plutovg_canvas_get_dash_offset = createExportWrapper('plutovg_canvas_get_dash_offset', 1);
  Module['_plutovg_canvas_get_dash_array'] = _plutovg_canvas_get_dash_array = createExportWrapper('plutovg_canvas_get_dash_array', 2);
  Module['_plutovg_canvas_shear'] = _plutovg_canvas_shear = createExportWrapper('plutovg_canvas_shear', 3);
  Module['_plutovg_matrix_shear'] = _plutovg_matrix_shear = createExportWrapper('plutovg_matrix_shear', 3);
  Module['_plutovg_canvas_rotate'] = _plutovg_canvas_rotate = createExportWrapper('plutovg_canvas_rotate', 2);
  Module['_plutovg_matrix_rotate'] = _plutovg_matrix_rotate = createExportWrapper('plutovg_matrix_rotate', 2);
  Module['_plutovg_canvas_reset_matrix'] = _plutovg_canvas_reset_matrix = createExportWrapper('plutovg_canvas_reset_matrix', 1);
  Module['_plutovg_canvas_map'] = _plutovg_canvas_map = createExportWrapper('plutovg_canvas_map', 5);
  Module['_plutovg_matrix_map'] = _plutovg_matrix_map = createExportWrapper('plutovg_matrix_map', 5);
  Module['_plutovg_canvas_map_point'] = _plutovg_canvas_map_point = createExportWrapper('plutovg_canvas_map_point', 3);
  Module['_plutovg_matrix_map_point'] = _plutovg_matrix_map_point = createExportWrapper('plutovg_matrix_map_point', 3);
  Module['_plutovg_canvas_map_rect'] = _plutovg_canvas_map_rect = createExportWrapper('plutovg_canvas_map_rect', 3);
  Module['_plutovg_canvas_move_to'] = _plutovg_canvas_move_to = createExportWrapper('plutovg_canvas_move_to', 3);
  Module['_plutovg_canvas_line_to'] = _plutovg_canvas_line_to = createExportWrapper('plutovg_canvas_line_to', 3);
  Module['_plutovg_canvas_quad_to'] = _plutovg_canvas_quad_to = createExportWrapper('plutovg_canvas_quad_to', 5);
  Module['_plutovg_path_quad_to'] = _plutovg_path_quad_to = createExportWrapper('plutovg_path_quad_to', 5);
  Module['_plutovg_canvas_cubic_to'] = _plutovg_canvas_cubic_to = createExportWrapper('plutovg_canvas_cubic_to', 7);
  Module['_plutovg_path_cubic_to'] = _plutovg_path_cubic_to = createExportWrapper('plutovg_path_cubic_to', 7);
  Module['_plutovg_canvas_arc_to'] = _plutovg_canvas_arc_to = createExportWrapper('plutovg_canvas_arc_to', 8);
  Module['_plutovg_path_arc_to'] = _plutovg_path_arc_to = createExportWrapper('plutovg_path_arc_to', 8);
  Module['_plutovg_canvas_rect'] = _plutovg_canvas_rect = createExportWrapper('plutovg_canvas_rect', 5);
  Module['_plutovg_path_add_rect'] = _plutovg_path_add_rect = createExportWrapper('plutovg_path_add_rect', 5);
  Module['_plutovg_canvas_round_rect'] = _plutovg_canvas_round_rect = createExportWrapper('plutovg_canvas_round_rect', 7);
  Module['_plutovg_canvas_ellipse'] = _plutovg_canvas_ellipse = createExportWrapper('plutovg_canvas_ellipse', 5);
  Module['_plutovg_canvas_circle'] = _plutovg_canvas_circle = createExportWrapper('plutovg_canvas_circle', 4);
  Module['_plutovg_canvas_arc'] = _plutovg_canvas_arc = createExportWrapper('plutovg_canvas_arc', 7);
  Module['_plutovg_path_add_arc'] = _plutovg_path_add_arc = createExportWrapper('plutovg_path_add_arc', 7);
  Module['_plutovg_canvas_add_path'] = _plutovg_canvas_add_path = createExportWrapper('plutovg_canvas_add_path', 2);
  Module['_plutovg_path_add_path'] = _plutovg_path_add_path = createExportWrapper('plutovg_path_add_path', 3);
  Module['_plutovg_canvas_new_path'] = _plutovg_canvas_new_path = createExportWrapper('plutovg_canvas_new_path', 1);
  Module['_plutovg_canvas_close_path'] = _plutovg_canvas_close_path = createExportWrapper('plutovg_canvas_close_path', 1);
  Module['_plutovg_canvas_get_current_point'] = _plutovg_canvas_get_current_point = createExportWrapper('plutovg_canvas_get_current_point', 3);
  Module['_plutovg_path_get_current_point'] = _plutovg_path_get_current_point = createExportWrapper('plutovg_path_get_current_point', 3);
  Module['_plutovg_canvas_get_path'] = _plutovg_canvas_get_path = createExportWrapper('plutovg_canvas_get_path', 1);
  Module['_plutovg_canvas_fill_extents'] = _plutovg_canvas_fill_extents = createExportWrapper('plutovg_canvas_fill_extents', 2);
  Module['_plutovg_canvas_stroke_extents'] = _plutovg_canvas_stroke_extents = createExportWrapper('plutovg_canvas_stroke_extents', 2);
  Module['_plutovg_canvas_clip_extents'] = _plutovg_canvas_clip_extents = createExportWrapper('plutovg_canvas_clip_extents', 2);
  Module['_plutovg_canvas_fill'] = _plutovg_canvas_fill = createExportWrapper('plutovg_canvas_fill', 1);
  Module['_plutovg_canvas_fill_preserve'] = _plutovg_canvas_fill_preserve = createExportWrapper('plutovg_canvas_fill_preserve', 1);
  Module['_plutovg_canvas_stroke'] = _plutovg_canvas_stroke = createExportWrapper('plutovg_canvas_stroke', 1);
  Module['_plutovg_canvas_stroke_preserve'] = _plutovg_canvas_stroke_preserve = createExportWrapper('plutovg_canvas_stroke_preserve', 1);
  Module['_plutovg_canvas_clip'] = _plutovg_canvas_clip = createExportWrapper('plutovg_canvas_clip', 1);
  Module['_plutovg_canvas_clip_preserve'] = _plutovg_canvas_clip_preserve = createExportWrapper('plutovg_canvas_clip_preserve', 1);
  Module['_plutovg_canvas_paint'] = _plutovg_canvas_paint = createExportWrapper('plutovg_canvas_paint', 1);
  Module['_plutovg_canvas_stroke_rect'] = _plutovg_canvas_stroke_rect = createExportWrapper('plutovg_canvas_stroke_rect', 5);
  Module['_plutovg_canvas_clip_rect'] = _plutovg_canvas_clip_rect = createExportWrapper('plutovg_canvas_clip_rect', 5);
  Module['_plutovg_canvas_clip_path'] = _plutovg_canvas_clip_path = createExportWrapper('plutovg_canvas_clip_path', 2);
  Module['_plutovg_canvas_add_glyph'] = _plutovg_canvas_add_glyph = createExportWrapper('plutovg_canvas_add_glyph', 4);
  Module['_plutovg_font_face_get_glyph_path'] = _plutovg_font_face_get_glyph_path = createExportWrapper('plutovg_font_face_get_glyph_path', 6);
  Module['_plutovg_canvas_add_text'] = _plutovg_canvas_add_text = createExportWrapper('plutovg_canvas_add_text', 6);
  Module['_plutovg_text_iterator_init'] = _plutovg_text_iterator_init = createExportWrapper('plutovg_text_iterator_init', 4);
  Module['_plutovg_text_iterator_has_next'] = _plutovg_text_iterator_has_next = createExportWrapper('plutovg_text_iterator_has_next', 1);
  Module['_plutovg_text_iterator_next'] = _plutovg_text_iterator_next = createExportWrapper('plutovg_text_iterator_next', 1);
  Module['_plutovg_canvas_fill_text'] = _plutovg_canvas_fill_text = createExportWrapper('plutovg_canvas_fill_text', 6);
  Module['_plutovg_canvas_stroke_text'] = _plutovg_canvas_stroke_text = createExportWrapper('plutovg_canvas_stroke_text', 6);
  Module['_plutovg_canvas_clip_text'] = _plutovg_canvas_clip_text = createExportWrapper('plutovg_canvas_clip_text', 6);
  Module['_plutovg_canvas_font_metrics'] = _plutovg_canvas_font_metrics = createExportWrapper('plutovg_canvas_font_metrics', 5);
  Module['_plutovg_font_face_get_metrics'] = _plutovg_font_face_get_metrics = createExportWrapper('plutovg_font_face_get_metrics', 6);
  Module['_plutovg_canvas_glyph_metrics'] = _plutovg_canvas_glyph_metrics = createExportWrapper('plutovg_canvas_glyph_metrics', 5);
  Module['_plutovg_font_face_get_glyph_metrics'] = _plutovg_font_face_get_glyph_metrics = createExportWrapper('plutovg_font_face_get_glyph_metrics', 6);
  Module['_plutovg_canvas_text_extents'] = _plutovg_canvas_text_extents = createExportWrapper('plutovg_canvas_text_extents', 5);
  Module['_plutovg_font_face_text_extents'] = _plutovg_font_face_text_extents = createExportWrapper('plutovg_font_face_text_extents', 6);
  Module['_plutovg_font_face_load_from_file'] = _plutovg_font_face_load_from_file = createExportWrapper('plutovg_font_face_load_from_file', 2);
  Module['_plutovg_font_face_load_from_data'] = _plutovg_font_face_load_from_data = createExportWrapper('plutovg_font_face_load_from_data', 5);
  Module['_plutovg_font_face_get_reference_count'] = _plutovg_font_face_get_reference_count = createExportWrapper('plutovg_font_face_get_reference_count', 1);
  Module['_plutovg_font_face_traverse_glyph_path'] = _plutovg_font_face_traverse_glyph_path = createExportWrapper('plutovg_font_face_traverse_glyph_path', 7);
  Module['_plutovg_matrix_map_points'] = _plutovg_matrix_map_points = createExportWrapper('plutovg_matrix_map_points', 4);
  Module['_plutovg_matrix_init'] = _plutovg_matrix_init = createExportWrapper('plutovg_matrix_init', 7);
  Module['_plutovg_matrix_init_rotate'] = _plutovg_matrix_init_rotate = createExportWrapper('plutovg_matrix_init_rotate', 2);
  Module['_plutovg_matrix_init_shear'] = _plutovg_matrix_init_shear = createExportWrapper('plutovg_matrix_init_shear', 3);
  Module['_plutovg_color_init_rgb'] = _plutovg_color_init_rgb = createExportWrapper('plutovg_color_init_rgb', 4);
  Module['_plutovg_color_init_rgb8'] = _plutovg_color_init_rgb8 = createExportWrapper('plutovg_color_init_rgb8', 4);
  Module['_plutovg_color_init_rgba8'] = _plutovg_color_init_rgba8 = createExportWrapper('plutovg_color_init_rgba8', 5);
  Module['_plutovg_color_init_rgba32'] = _plutovg_color_init_rgba32 = createExportWrapper('plutovg_color_init_rgba32', 2);
  Module['_plutovg_color_to_rgba32'] = _plutovg_color_to_rgba32 = createExportWrapper('plutovg_color_to_rgba32', 1);
  Module['_plutovg_paint_create_rgb'] = _plutovg_paint_create_rgb = createExportWrapper('plutovg_paint_create_rgb', 3);
  Module['_plutovg_paint_create_rgba'] = _plutovg_paint_create_rgba = createExportWrapper('plutovg_paint_create_rgba', 4);
  Module['_plutovg_paint_create_color'] = _plutovg_paint_create_color = createExportWrapper('plutovg_paint_create_color', 1);
  Module['_plutovg_paint_get_reference_count'] = _plutovg_paint_get_reference_count = createExportWrapper('plutovg_paint_get_reference_count', 1);
  Module['_plutovg_path_iterator_init'] = _plutovg_path_iterator_init = createExportWrapper('plutovg_path_iterator_init', 2);
  Module['_plutovg_path_iterator_has_next'] = _plutovg_path_iterator_has_next = createExportWrapper('plutovg_path_iterator_has_next', 1);
  Module['_plutovg_path_iterator_next'] = _plutovg_path_iterator_next = createExportWrapper('plutovg_path_iterator_next', 2);
  Module['_plutovg_path_reference'] = _plutovg_path_reference = createExportWrapper('plutovg_path_reference', 1);
  Module['_plutovg_path_get_reference_count'] = _plutovg_path_get_reference_count = createExportWrapper('plutovg_path_get_reference_count', 1);
  Module['_plutovg_path_get_elements'] = _plutovg_path_get_elements = createExportWrapper('plutovg_path_get_elements', 2);
  Module['_plutovg_path_reserve'] = _plutovg_path_reserve = createExportWrapper('plutovg_path_reserve', 2);
  Module['_plutovg_path_transform'] = _plutovg_path_transform = createExportWrapper('plutovg_path_transform', 2);
  Module['_plutovg_path_traverse'] = _plutovg_path_traverse = createExportWrapper('plutovg_path_traverse', 3);
  Module['_plutovg_path_traverse_flatten'] = _plutovg_path_traverse_flatten = createExportWrapper('plutovg_path_traverse_flatten', 3);
  Module['_plutovg_path_traverse_dashed'] = _plutovg_path_traverse_dashed = createExportWrapper('plutovg_path_traverse_dashed', 6);
  Module['_plutovg_path_clone'] = _plutovg_path_clone = createExportWrapper('plutovg_path_clone', 1);
  Module['_plutovg_path_clone_flatten'] = _plutovg_path_clone_flatten = createExportWrapper('plutovg_path_clone_flatten', 1);
  Module['_plutovg_path_clone_dashed'] = _plutovg_path_clone_dashed = createExportWrapper('plutovg_path_clone_dashed', 4);
  Module['_plutovg_path_length'] = _plutovg_path_length = createExportWrapper('plutovg_path_length', 1);
  Module['_plutovg_surface_load_from_image_file'] = _plutovg_surface_load_from_image_file = createExportWrapper('plutovg_surface_load_from_image_file', 1);
  Module['_plutovg_surface_load_from_image_data'] = _plutovg_surface_load_from_image_data = createExportWrapper('plutovg_surface_load_from_image_data', 2);
  Module['_plutovg_surface_get_reference_count'] = _plutovg_surface_get_reference_count = createExportWrapper('plutovg_surface_get_reference_count', 1);
  Module['_plutovg_surface_get_data'] = _plutovg_surface_get_data = createExportWrapper('plutovg_surface_get_data', 1);
  Module['_plutovg_surface_get_stride'] = _plutovg_surface_get_stride = createExportWrapper('plutovg_surface_get_stride', 1);
  Module['_plutovg_surface_clear'] = _plutovg_surface_clear = createExportWrapper('plutovg_surface_clear', 2);
  Module['_plutovg_surface_write_to_png'] = _plutovg_surface_write_to_png = createExportWrapper('plutovg_surface_write_to_png', 2);
  Module['_plutovg_surface_write_to_jpg'] = _plutovg_surface_write_to_jpg = createExportWrapper('plutovg_surface_write_to_jpg', 3);
  Module['_plutovg_surface_write_to_png_stream'] = _plutovg_surface_write_to_png_stream = createExportWrapper('plutovg_surface_write_to_png_stream', 3);
  Module['_plutovg_surface_write_to_jpg_stream'] = _plutovg_surface_write_to_jpg_stream = createExportWrapper('plutovg_surface_write_to_jpg_stream', 4);
  Module['_plutovg_convert_argb_to_rgba'] = _plutovg_convert_argb_to_rgba = createExportWrapper('plutovg_convert_argb_to_rgba', 5);
  Module['_plutovg_convert_rgba_to_argb'] = _plutovg_convert_rgba_to_argb = createExportWrapper('plutovg_convert_rgba_to_argb', 5);
  Module['_SDL_ReadU16BE'] = _SDL_ReadU16BE = createExportWrapper('SDL_ReadU16BE', 2);
  Module['_SDL_iconv_string'] = _SDL_iconv_string = createExportWrapper('SDL_iconv_string', 4);
  Module['_SDL_IOFromConstMem'] = _SDL_IOFromConstMem = createExportWrapper('SDL_IOFromConstMem', 2);
  Module['_SDL_CreateAudioStream'] = _SDL_CreateAudioStream = createExportWrapper('SDL_CreateAudioStream', 2);
  Module['_SDL_ClearAudioStream'] = _SDL_ClearAudioStream = createExportWrapper('SDL_ClearAudioStream', 1);
  Module['_SDL_DestroyAudioStream'] = _SDL_DestroyAudioStream = createExportWrapper('SDL_DestroyAudioStream', 1);
  Module['_SDL_GetAudioStreamData'] = _SDL_GetAudioStreamData = createExportWrapper('SDL_GetAudioStreamData', 3);
  Module['_SDL_PutAudioStreamData'] = _SDL_PutAudioStreamData = createExportWrapper('SDL_PutAudioStreamData', 3);
  Module['_SDL_FlushAudioStream'] = _SDL_FlushAudioStream = createExportWrapper('SDL_FlushAudioStream', 1);
  Module['_SDL_scalbn'] = _SDL_scalbn = createExportWrapper('SDL_scalbn', 2);
  Module['_SDL_log'] = _SDL_log = createExportWrapper('SDL_log', 1);
  Module['_SDL_exp'] = _SDL_exp = createExportWrapper('SDL_exp', 1);
  Module['_SDL_floor'] = _SDL_floor = createExportWrapper('SDL_floor', 1);
  Module['_SDL_cos'] = _SDL_cos = createExportWrapper('SDL_cos', 1);
  Module['_SDL_sin'] = _SDL_sin = createExportWrapper('SDL_sin', 1);
  Module['_SDL_memcpy'] = _SDL_memcpy = createExportWrapper('SDL_memcpy', 3);
  Module['_SDL_getenv'] = _SDL_getenv = createExportWrapper('SDL_getenv', 1);
  Module['_Mix_GetTimidityCfg'] = _Mix_GetTimidityCfg = createExportWrapper('Mix_GetTimidityCfg', 0);
  Module['_Mix_SetPanning'] = _Mix_SetPanning = createExportWrapper('Mix_SetPanning', 3);
  Module['_Mix_QuerySpec'] = _Mix_QuerySpec = createExportWrapper('Mix_QuerySpec', 3);
  Module['_Mix_SetPosition'] = _Mix_SetPosition = createExportWrapper('Mix_SetPosition', 3);
  Module['_Mix_SetDistance'] = _Mix_SetDistance = createExportWrapper('Mix_SetDistance', 2);
  Module['_Mix_SetReverseStereo'] = _Mix_SetReverseStereo = createExportWrapper('Mix_SetReverseStereo', 2);
  Module['_Mix_UnregisterEffect'] = _Mix_UnregisterEffect = createExportWrapper('Mix_UnregisterEffect', 2);
  Module['_Mix_RegisterEffect'] = _Mix_RegisterEffect = createExportWrapper('Mix_RegisterEffect', 4);
  Module['_Mix_GetNumChunkDecoders'] = _Mix_GetNumChunkDecoders = createExportWrapper('Mix_GetNumChunkDecoders', 0);
  Module['_Mix_GetChunkDecoder'] = _Mix_GetChunkDecoder = createExportWrapper('Mix_GetChunkDecoder', 1);
  Module['_Mix_HasChunkDecoder'] = _Mix_HasChunkDecoder = createExportWrapper('Mix_HasChunkDecoder', 1);
  Module['_Mix_Version'] = _Mix_Version = createExportWrapper('Mix_Version', 0);
  Module['_Mix_Init'] = _Mix_Init = createExportWrapper('Mix_Init', 1);
  Module['_Mix_Quit'] = _Mix_Quit = createExportWrapper('Mix_Quit', 0);
  Module['_Mix_OpenAudio'] = _Mix_OpenAudio = createExportWrapper('Mix_OpenAudio', 2);
  Module['_SDL_WasInit'] = _SDL_WasInit = createExportWrapper('SDL_WasInit', 1);
  Module['_SDL_InitSubSystem'] = _SDL_InitSubSystem = createExportWrapper('SDL_InitSubSystem', 1);
  Module['_Mix_CloseAudio'] = _Mix_CloseAudio = createExportWrapper('Mix_CloseAudio', 0);
  Module['_SDL_OpenAudioDevice'] = _SDL_OpenAudioDevice = createExportWrapper('SDL_OpenAudioDevice', 2);
  Module['_SDL_GetAudioDeviceFormat'] = _SDL_GetAudioDeviceFormat = createExportWrapper('SDL_GetAudioDeviceFormat', 3);
  Module['_SDL_CloseAudioDevice'] = _SDL_CloseAudioDevice = createExportWrapper('SDL_CloseAudioDevice', 1);
  Module['_Mix_VolumeMusic'] = _Mix_VolumeMusic = createExportWrapper('Mix_VolumeMusic', 1);
  Module['_SDL_BindAudioStream'] = _SDL_BindAudioStream = createExportWrapper('SDL_BindAudioStream', 2);
  Module['_SDL_SetAudioStreamGetCallback'] = _SDL_SetAudioStreamGetCallback = createExportWrapper('SDL_SetAudioStreamGetCallback', 3);
  Module['_SDL_LockAudioStream'] = _SDL_LockAudioStream = createExportWrapper('SDL_LockAudioStream', 1);
  Module['_SDL_UnlockAudioStream'] = _SDL_UnlockAudioStream = createExportWrapper('SDL_UnlockAudioStream', 1);
  Module['_SDL_GetSIMDAlignment'] = _SDL_GetSIMDAlignment = createExportWrapper('SDL_GetSIMDAlignment', 0);
  Module['_SDL_GetSilenceValueForFormat'] = _SDL_GetSilenceValueForFormat = createExportWrapper('SDL_GetSilenceValueForFormat', 1);
  Module['_SDL_GetTicks'] = _SDL_GetTicks = createExportWrapper('SDL_GetTicks', 0);
  Module['_Mix_Volume'] = _Mix_Volume = createExportWrapper('Mix_Volume', 2);
  Module['_SDL_MixAudio'] = _SDL_MixAudio = createExportWrapper('SDL_MixAudio', 5);
  Module['_Mix_PauseAudio'] = _Mix_PauseAudio = createExportWrapper('Mix_PauseAudio', 1);
  Module['_SDL_PauseAudioDevice'] = _SDL_PauseAudioDevice = createExportWrapper('SDL_PauseAudioDevice', 1);
  Module['_SDL_ResumeAudioDevice'] = _SDL_ResumeAudioDevice = createExportWrapper('SDL_ResumeAudioDevice', 1);
  Module['_Mix_AllocateChannels'] = _Mix_AllocateChannels = createExportWrapper('Mix_AllocateChannels', 1);
  Module['_Mix_UnregisterAllEffects'] = _Mix_UnregisterAllEffects = createExportWrapper('Mix_UnregisterAllEffects', 1);
  Module['_Mix_HaltChannel'] = _Mix_HaltChannel = createExportWrapper('Mix_HaltChannel', 1);
  Module['_Mix_LoadWAV_IO'] = _Mix_LoadWAV_IO = createExportWrapper('Mix_LoadWAV_IO', 2);
  Module['_SDL_LoadWAV_IO'] = _SDL_LoadWAV_IO = createExportWrapper('SDL_LoadWAV_IO', 5);
  Module['_SDL_ConvertAudioSamples'] = _SDL_ConvertAudioSamples = createExportWrapper('SDL_ConvertAudioSamples', 6);
  Module['_Mix_LoadWAV'] = _Mix_LoadWAV = createExportWrapper('Mix_LoadWAV', 1);
  Module['_Mix_QuickLoad_WAV'] = _Mix_QuickLoad_WAV = createExportWrapper('Mix_QuickLoad_WAV', 1);
  Module['_Mix_QuickLoad_RAW'] = _Mix_QuickLoad_RAW = createExportWrapper('Mix_QuickLoad_RAW', 2);
  Module['_Mix_FreeChunk'] = _Mix_FreeChunk = createExportWrapper('Mix_FreeChunk', 1);
  Module['_Mix_SetPostMix'] = _Mix_SetPostMix = createExportWrapper('Mix_SetPostMix', 2);
  Module['_Mix_HookMusic'] = _Mix_HookMusic = createExportWrapper('Mix_HookMusic', 2);
  Module['_Mix_GetMusicHookData'] = _Mix_GetMusicHookData = createExportWrapper('Mix_GetMusicHookData', 0);
  Module['_Mix_ChannelFinished'] = _Mix_ChannelFinished = createExportWrapper('Mix_ChannelFinished', 1);
  Module['_Mix_ReserveChannels'] = _Mix_ReserveChannels = createExportWrapper('Mix_ReserveChannels', 1);
  Module['_Mix_PlayChannelTimed'] = _Mix_PlayChannelTimed = createExportWrapper('Mix_PlayChannelTimed', 4);
  Module['_Mix_Playing'] = _Mix_Playing = createExportWrapper('Mix_Playing', 1);
  Module['_Mix_PlayChannel'] = _Mix_PlayChannel = createExportWrapper('Mix_PlayChannel', 3);
  Module['_Mix_ExpireChannel'] = _Mix_ExpireChannel = createExportWrapper('Mix_ExpireChannel', 2);
  Module['_Mix_FadeInChannelTimed'] = _Mix_FadeInChannelTimed = createExportWrapper('Mix_FadeInChannelTimed', 5);
  Module['_Mix_FadeInChannel'] = _Mix_FadeInChannel = createExportWrapper('Mix_FadeInChannel', 4);
  Module['_Mix_VolumeChunk'] = _Mix_VolumeChunk = createExportWrapper('Mix_VolumeChunk', 2);
  Module['_Mix_HaltGroup'] = _Mix_HaltGroup = createExportWrapper('Mix_HaltGroup', 1);
  Module['_Mix_FadeOutChannel'] = _Mix_FadeOutChannel = createExportWrapper('Mix_FadeOutChannel', 2);
  Module['_Mix_FadeOutGroup'] = _Mix_FadeOutGroup = createExportWrapper('Mix_FadeOutGroup', 2);
  Module['_Mix_FadingChannel'] = _Mix_FadingChannel = createExportWrapper('Mix_FadingChannel', 1);
  Module['_Mix_GetChunk'] = _Mix_GetChunk = createExportWrapper('Mix_GetChunk', 1);
  Module['_Mix_Pause'] = _Mix_Pause = createExportWrapper('Mix_Pause', 1);
  Module['_Mix_PauseGroup'] = _Mix_PauseGroup = createExportWrapper('Mix_PauseGroup', 1);
  Module['_Mix_Resume'] = _Mix_Resume = createExportWrapper('Mix_Resume', 1);
  Module['_Mix_ResumeGroup'] = _Mix_ResumeGroup = createExportWrapper('Mix_ResumeGroup', 1);
  Module['_Mix_Paused'] = _Mix_Paused = createExportWrapper('Mix_Paused', 1);
  Module['_Mix_GroupChannel'] = _Mix_GroupChannel = createExportWrapper('Mix_GroupChannel', 2);
  Module['_Mix_GroupChannels'] = _Mix_GroupChannels = createExportWrapper('Mix_GroupChannels', 3);
  Module['_Mix_GroupAvailable'] = _Mix_GroupAvailable = createExportWrapper('Mix_GroupAvailable', 1);
  Module['_Mix_GroupCount'] = _Mix_GroupCount = createExportWrapper('Mix_GroupCount', 1);
  Module['_Mix_GroupOldest'] = _Mix_GroupOldest = createExportWrapper('Mix_GroupOldest', 1);
  Module['_Mix_GroupNewer'] = _Mix_GroupNewer = createExportWrapper('Mix_GroupNewer', 1);
  Module['_Mix_MasterVolume'] = _Mix_MasterVolume = createExportWrapper('Mix_MasterVolume', 1);
  Module['_SDL_SetAtomicInt'] = _SDL_SetAtomicInt = createExportWrapper('SDL_SetAtomicInt', 2);
  Module['_Mix_GetNumMusicDecoders'] = _Mix_GetNumMusicDecoders = createExportWrapper('Mix_GetNumMusicDecoders', 0);
  Module['_Mix_GetMusicDecoder'] = _Mix_GetMusicDecoder = createExportWrapper('Mix_GetMusicDecoder', 1);
  Module['_Mix_HasMusicDecoder'] = _Mix_HasMusicDecoder = createExportWrapper('Mix_HasMusicDecoder', 1);
  Module['_Mix_HookMusicFinished'] = _Mix_HookMusicFinished = createExportWrapper('Mix_HookMusicFinished', 1);
  Module['_SDL_snprintf'] = _SDL_snprintf = createExportWrapper('SDL_snprintf', 4);
  Module['_SDL_GetHintBoolean'] = _SDL_GetHintBoolean = createExportWrapper('SDL_GetHintBoolean', 2);
  Module['_SDL_GetError'] = _SDL_GetError = createExportWrapper('SDL_GetError', 0);
  Module['_Mix_LoadMUS'] = _Mix_LoadMUS = createExportWrapper('Mix_LoadMUS', 1);
  Module['_Mix_LoadMUSType_IO'] = _Mix_LoadMUSType_IO = createExportWrapper('Mix_LoadMUSType_IO', 3);
  Module['_SDL_ClearError'] = _SDL_ClearError = createExportWrapper('SDL_ClearError', 0);
  Module['_Mix_LoadMUS_IO'] = _Mix_LoadMUS_IO = createExportWrapper('Mix_LoadMUS_IO', 2);
  Module['_Mix_FreeMusic'] = _Mix_FreeMusic = createExportWrapper('Mix_FreeMusic', 1);
  Module['_SDL_Delay'] = _SDL_Delay = createExportWrapper('SDL_Delay', 1);
  Module['_Mix_GetMusicType'] = _Mix_GetMusicType = createExportWrapper('Mix_GetMusicType', 1);
  Module['_Mix_GetMusicTitleTag'] = _Mix_GetMusicTitleTag = createExportWrapper('Mix_GetMusicTitleTag', 1);
  Module['_Mix_GetMusicTitle'] = _Mix_GetMusicTitle = createExportWrapper('Mix_GetMusicTitle', 1);
  Module['_Mix_GetMusicArtistTag'] = _Mix_GetMusicArtistTag = createExportWrapper('Mix_GetMusicArtistTag', 1);
  Module['_Mix_GetMusicAlbumTag'] = _Mix_GetMusicAlbumTag = createExportWrapper('Mix_GetMusicAlbumTag', 1);
  Module['_Mix_GetMusicCopyrightTag'] = _Mix_GetMusicCopyrightTag = createExportWrapper('Mix_GetMusicCopyrightTag', 1);
  Module['_Mix_FadeInMusicPos'] = _Mix_FadeInMusicPos = createExportWrapper('Mix_FadeInMusicPos', 4);
  Module['_Mix_FadeInMusic'] = _Mix_FadeInMusic = createExportWrapper('Mix_FadeInMusic', 3);
  Module['_Mix_PlayMusic'] = _Mix_PlayMusic = createExportWrapper('Mix_PlayMusic', 2);
  Module['_Mix_ModMusicJumpToOrder'] = _Mix_ModMusicJumpToOrder = createExportWrapper('Mix_ModMusicJumpToOrder', 1);
  Module['_Mix_SetMusicPosition'] = _Mix_SetMusicPosition = createExportWrapper('Mix_SetMusicPosition', 1);
  Module['_Mix_GetMusicPosition'] = _Mix_GetMusicPosition = createExportWrapper('Mix_GetMusicPosition', 1);
  Module['_Mix_MusicDuration'] = _Mix_MusicDuration = createExportWrapper('Mix_MusicDuration', 1);
  Module['_Mix_GetMusicLoopStartTime'] = _Mix_GetMusicLoopStartTime = createExportWrapper('Mix_GetMusicLoopStartTime', 1);
  Module['_Mix_GetMusicLoopEndTime'] = _Mix_GetMusicLoopEndTime = createExportWrapper('Mix_GetMusicLoopEndTime', 1);
  Module['_Mix_GetMusicLoopLengthTime'] = _Mix_GetMusicLoopLengthTime = createExportWrapper('Mix_GetMusicLoopLengthTime', 1);
  Module['_Mix_GetMusicVolume'] = _Mix_GetMusicVolume = createExportWrapper('Mix_GetMusicVolume', 1);
  Module['_Mix_HaltMusic'] = _Mix_HaltMusic = createExportWrapper('Mix_HaltMusic', 0);
  Module['_Mix_FadeOutMusic'] = _Mix_FadeOutMusic = createExportWrapper('Mix_FadeOutMusic', 1);
  Module['_Mix_FadingMusic'] = _Mix_FadingMusic = createExportWrapper('Mix_FadingMusic', 0);
  Module['_Mix_PauseMusic'] = _Mix_PauseMusic = createExportWrapper('Mix_PauseMusic', 0);
  Module['_Mix_ResumeMusic'] = _Mix_ResumeMusic = createExportWrapper('Mix_ResumeMusic', 0);
  Module['_Mix_RewindMusic'] = _Mix_RewindMusic = createExportWrapper('Mix_RewindMusic', 0);
  Module['_Mix_PausedMusic'] = _Mix_PausedMusic = createExportWrapper('Mix_PausedMusic', 0);
  Module['_Mix_StartTrack'] = _Mix_StartTrack = createExportWrapper('Mix_StartTrack', 2);
  Module['_Mix_GetNumTracks'] = _Mix_GetNumTracks = createExportWrapper('Mix_GetNumTracks', 1);
  Module['_Mix_PlayingMusic'] = _Mix_PlayingMusic = createExportWrapper('Mix_PlayingMusic', 0);
  Module['_Mix_SetTimidityCfg'] = _Mix_SetTimidityCfg = createExportWrapper('Mix_SetTimidityCfg', 1);
  Module['_Mix_SetSoundFonts'] = _Mix_SetSoundFonts = createExportWrapper('Mix_SetSoundFonts', 1);
  Module['_Mix_GetSoundFonts'] = _Mix_GetSoundFonts = createExportWrapper('Mix_GetSoundFonts', 0);
  Module['_Mix_EachSoundFont'] = _Mix_EachSoundFont = createExportWrapper('Mix_EachSoundFont', 2);
  Module['_SDL_strtok_r'] = _SDL_strtok_r = createExportWrapper('SDL_strtok_r', 3);
  Module['_SDL_atoi'] = _SDL_atoi = createExportWrapper('SDL_atoi', 1);
  Module['_SDL_atof'] = _SDL_atof = createExportWrapper('SDL_atof', 1);
  Module['_SDL_ReadS16BE'] = _SDL_ReadS16BE = createExportWrapper('SDL_ReadS16BE', 2);
  Module['_SDL_SetAppMetadata'] = _SDL_SetAppMetadata = createExportWrapper('SDL_SetAppMetadata', 3);
  Module['_SDL_GetGlobalProperties'] = _SDL_GetGlobalProperties = createExportWrapper('SDL_GetGlobalProperties', 0);
  Module['_SDL_SetAppMetadataProperty'] = _SDL_SetAppMetadataProperty = createExportWrapper('SDL_SetAppMetadataProperty', 2);
  Module['_SDL_GetAppMetadataProperty'] = _SDL_GetAppMetadataProperty = createExportWrapper('SDL_GetAppMetadataProperty', 1);
  Module['_SDL_GetHint'] = _SDL_GetHint = createExportWrapper('SDL_GetHint', 1);
  Module['_SDL_SetMainReady'] = _SDL_SetMainReady = createExportWrapper('SDL_SetMainReady', 0);
  Module['_SDL_GetCurrentThreadID'] = _SDL_GetCurrentThreadID = createExportWrapper('SDL_GetCurrentThreadID', 0);
  Module['_SDL_IsMainThread'] = _SDL_IsMainThread = createExportWrapper('SDL_IsMainThread', 0);
  Module['_SDL_LogInfo'] = _SDL_LogInfo = createExportWrapper('SDL_LogInfo', 3);
  Module['_SDL_QuitSubSystem'] = _SDL_QuitSubSystem = createExportWrapper('SDL_QuitSubSystem', 1);
  Module['_SDL_Init'] = _SDL_Init = createExportWrapper('SDL_Init', 1);
  Module['_SDL_Quit'] = _SDL_Quit = createExportWrapper('SDL_Quit', 0);
  Module['_SDL_GetVersion'] = _SDL_GetVersion = createExportWrapper('SDL_GetVersion', 0);
  Module['_SDL_GetRevision'] = _SDL_GetRevision = createExportWrapper('SDL_GetRevision', 0);
  Module['_SDL_GetPlatform'] = _SDL_GetPlatform = createExportWrapper('SDL_GetPlatform', 0);
  Module['_SDL_IsTablet'] = _SDL_IsTablet = createExportWrapper('SDL_IsTablet', 0);
  Module['_SDL_IsTV'] = _SDL_IsTV = createExportWrapper('SDL_IsTV', 0);
  Module['_SDL_GetSandbox'] = _SDL_GetSandbox = createExportWrapper('SDL_GetSandbox', 0);
  Module['_SDL_ReportAssertion'] = _SDL_ReportAssertion = createExportWrapper('SDL_ReportAssertion', 4);
  Module['_SDL_SetAssertionHandler'] = _SDL_SetAssertionHandler = createExportWrapper('SDL_SetAssertionHandler', 2);
  Module['_SDL_MinimizeWindow'] = _SDL_MinimizeWindow = createExportWrapper('SDL_MinimizeWindow', 1);
  Module['_SDL_ShowMessageBox'] = _SDL_ShowMessageBox = createExportWrapper('SDL_ShowMessageBox', 2);
  Module['_SDL_RestoreWindow'] = _SDL_RestoreWindow = createExportWrapper('SDL_RestoreWindow', 1);
  Module['_SDL_GetAssertionReport'] = _SDL_GetAssertionReport = createExportWrapper('SDL_GetAssertionReport', 0);
  Module['_SDL_ResetAssertionReport'] = _SDL_ResetAssertionReport = createExportWrapper('SDL_ResetAssertionReport', 0);
  Module['_SDL_GetDefaultAssertionHandler'] = _SDL_GetDefaultAssertionHandler = createExportWrapper('SDL_GetDefaultAssertionHandler', 0);
  Module['_SDL_GetAssertionHandler'] = _SDL_GetAssertionHandler = createExportWrapper('SDL_GetAssertionHandler', 1);
  Module['_SDL_LogMessageV'] = _SDL_LogMessageV = createExportWrapper('SDL_LogMessageV', 4);
  Module['_SDL_SetErrorV'] = _SDL_SetErrorV = createExportWrapper('SDL_SetErrorV', 2);
  Module['_SDL_vsnprintf'] = _SDL_vsnprintf = createExportWrapper('SDL_vsnprintf', 4);
  Module['_SDL_OutOfMemory'] = _SDL_OutOfMemory = createExportWrapper('SDL_OutOfMemory', 0);
  Module['_SDL_GUIDToString'] = _SDL_GUIDToString = createExportWrapper('SDL_GUIDToString', 3);
  Module['_SDL_StringToGUID'] = _SDL_StringToGUID = createExportWrapper('SDL_StringToGUID', 2);
  Module['_SDL_GetAtomicU32'] = _SDL_GetAtomicU32 = createExportWrapper('SDL_GetAtomicU32', 1);
  Module['_SDL_CompareAndSwapAtomicU32'] = _SDL_CompareAndSwapAtomicU32 = createExportWrapper('SDL_CompareAndSwapAtomicU32', 3);
  Module['_SDL_SetHintWithPriority'] = _SDL_SetHintWithPriority = createExportWrapper('SDL_SetHintWithPriority', 3);
  Module['_SDL_LockProperties'] = _SDL_LockProperties = createExportWrapper('SDL_LockProperties', 1);
  Module['_SDL_SetPointerPropertyWithCleanup'] = _SDL_SetPointerPropertyWithCleanup = createExportWrapper('SDL_SetPointerPropertyWithCleanup', 5);
  Module['_SDL_UnlockProperties'] = _SDL_UnlockProperties = createExportWrapper('SDL_UnlockProperties', 1);
  Module['_SDL_ResetHint'] = _SDL_ResetHint = createExportWrapper('SDL_ResetHint', 1);
  Module['_SDL_ResetHints'] = _SDL_ResetHints = createExportWrapper('SDL_ResetHints', 0);
  Module['_SDL_EnumerateProperties'] = _SDL_EnumerateProperties = createExportWrapper('SDL_EnumerateProperties', 3);
  Module['_SDL_SetHint'] = _SDL_SetHint = createExportWrapper('SDL_SetHint', 2);
  Module['_SDL_AddHintCallback'] = _SDL_AddHintCallback = createExportWrapper('SDL_AddHintCallback', 3);
  Module['_SDL_RemoveHintCallback'] = _SDL_RemoveHintCallback = createExportWrapper('SDL_RemoveHintCallback', 3);
  Module['_SDL_ResetLogPriorities'] = _SDL_ResetLogPriorities = createExportWrapper('SDL_ResetLogPriorities', 0);
  Module['_SDL_SetLogPriorities'] = _SDL_SetLogPriorities = createExportWrapper('SDL_SetLogPriorities', 1);
  Module['_SDL_SetLogPriority'] = _SDL_SetLogPriority = createExportWrapper('SDL_SetLogPriority', 2);
  Module['_SDL_GetLogPriority'] = _SDL_GetLogPriority = createExportWrapper('SDL_GetLogPriority', 1);
  Module['_SDL_SetLogPriorityPrefix'] = _SDL_SetLogPriorityPrefix = createExportWrapper('SDL_SetLogPriorityPrefix', 2);
  Module['_SDL_LogTrace'] = _SDL_LogTrace = createExportWrapper('SDL_LogTrace', 3);
  Module['_SDL_LogVerbose'] = _SDL_LogVerbose = createExportWrapper('SDL_LogVerbose', 3);
  Module['_SDL_LogDebug'] = _SDL_LogDebug = createExportWrapper('SDL_LogDebug', 3);
  Module['_SDL_LogWarn'] = _SDL_LogWarn = createExportWrapper('SDL_LogWarn', 3);
  Module['_SDL_LogError'] = _SDL_LogError = createExportWrapper('SDL_LogError', 3);
  Module['_SDL_LogCritical'] = _SDL_LogCritical = createExportWrapper('SDL_LogCritical', 3);
  Module['_SDL_LogMessage'] = _SDL_LogMessage = createExportWrapper('SDL_LogMessage', 4);
  Module['_SDL_GetDefaultLogOutputFunction'] = _SDL_GetDefaultLogOutputFunction = createExportWrapper('SDL_GetDefaultLogOutputFunction', 0);
  Module['_SDL_GetLogOutputFunction'] = _SDL_GetLogOutputFunction = createExportWrapper('SDL_GetLogOutputFunction', 2);
  Module['_SDL_SetLogOutputFunction'] = _SDL_SetLogOutputFunction = createExportWrapper('SDL_SetLogOutputFunction', 2);
  Module['_SDL_CopyProperties'] = _SDL_CopyProperties = createExportWrapper('SDL_CopyProperties', 2);
  Module['_SDL_ClearProperty'] = _SDL_ClearProperty = createExportWrapper('SDL_ClearProperty', 2);
  Module['_SDL_HasProperty'] = _SDL_HasProperty = createExportWrapper('SDL_HasProperty', 2);
  Module['_SDL_GetPropertyType'] = _SDL_GetPropertyType = createExportWrapper('SDL_GetPropertyType', 2);
  Module['_SDL_asprintf'] = _SDL_asprintf = createExportWrapper('SDL_asprintf', 3);
  Module['_SDL_round'] = _SDL_round = createExportWrapper('SDL_round', 1);
  Module['_SDL_GetTLS'] = _SDL_GetTLS = createExportWrapper('SDL_GetTLS', 1);
  Module['_SDL_SetTLS'] = _SDL_SetTLS = createExportWrapper('SDL_SetTLS', 3);
  Module['_SDL_tolower'] = _SDL_tolower = createExportWrapper('SDL_tolower', 1);
  Module['_SDL_CompareAndSwapAtomicInt'] = _SDL_CompareAndSwapAtomicInt = createExportWrapper('SDL_CompareAndSwapAtomicInt', 3);
  Module['_SDL_CompareAndSwapAtomicPointer'] = _SDL_CompareAndSwapAtomicPointer = createExportWrapper('SDL_CompareAndSwapAtomicPointer', 3);
  Module['_SDL_SetAtomicU32'] = _SDL_SetAtomicU32 = createExportWrapper('SDL_SetAtomicU32', 2);
  Module['_SDL_SetAtomicPointer'] = _SDL_SetAtomicPointer = createExportWrapper('SDL_SetAtomicPointer', 2);
  Module['_SDL_GetAtomicPointer'] = _SDL_GetAtomicPointer = createExportWrapper('SDL_GetAtomicPointer', 1);
  Module['_SDL_MemoryBarrierReleaseFunction'] = _SDL_MemoryBarrierReleaseFunction = createExportWrapper('SDL_MemoryBarrierReleaseFunction', 0);
  Module['_SDL_LockSpinlock'] = _SDL_LockSpinlock = createExportWrapper('SDL_LockSpinlock', 1);
  Module['_SDL_UnlockSpinlock'] = _SDL_UnlockSpinlock = createExportWrapper('SDL_UnlockSpinlock', 1);
  Module['_SDL_MemoryBarrierAcquireFunction'] = _SDL_MemoryBarrierAcquireFunction = createExportWrapper('SDL_MemoryBarrierAcquireFunction', 0);
  Module['_SDL_TryLockSpinlock'] = _SDL_TryLockSpinlock = createExportWrapper('SDL_TryLockSpinlock', 1);
  Module['_SDL_GetNumAudioDrivers'] = _SDL_GetNumAudioDrivers = createExportWrapper('SDL_GetNumAudioDrivers', 0);
  Module['_SDL_GetAudioDriver'] = _SDL_GetAudioDriver = createExportWrapper('SDL_GetAudioDriver', 1);
  Module['_SDL_GetCurrentAudioDriver'] = _SDL_GetCurrentAudioDriver = createExportWrapper('SDL_GetCurrentAudioDriver', 0);
  Module['_SDL_IsAudioDevicePhysical'] = _SDL_IsAudioDevicePhysical = createExportWrapper('SDL_IsAudioDevicePhysical', 1);
  Module['_SDL_IsAudioDevicePlayback'] = _SDL_IsAudioDevicePlayback = createExportWrapper('SDL_IsAudioDevicePlayback', 1);
  Module['_SDL_DestroyCondition'] = _SDL_DestroyCondition = createExportWrapper('SDL_DestroyCondition', 1);
  Module['_SDL_GetAudioPlaybackDevices'] = _SDL_GetAudioPlaybackDevices = createExportWrapper('SDL_GetAudioPlaybackDevices', 1);
  Module['_SDL_GetAudioRecordingDevices'] = _SDL_GetAudioRecordingDevices = createExportWrapper('SDL_GetAudioRecordingDevices', 1);
  Module['_SDL_GetAudioDeviceName'] = _SDL_GetAudioDeviceName = createExportWrapper('SDL_GetAudioDeviceName', 1);
  Module['_SDL_GetAudioDeviceChannelMap'] = _SDL_GetAudioDeviceChannelMap = createExportWrapper('SDL_GetAudioDeviceChannelMap', 2);
  Module['_SDL_WaitCondition'] = _SDL_WaitCondition = createExportWrapper('SDL_WaitCondition', 2);
  Module['_SDL_WaitThread'] = _SDL_WaitThread = createExportWrapper('SDL_WaitThread', 2);
  Module['_SDL_BroadcastCondition'] = _SDL_BroadcastCondition = createExportWrapper('SDL_BroadcastCondition', 1);
  Module['_SDL_CreateThreadRuntime'] = _SDL_CreateThreadRuntime = createExportWrapper('SDL_CreateThreadRuntime', 5);
  Module['_SDL_AudioDevicePaused'] = _SDL_AudioDevicePaused = createExportWrapper('SDL_AudioDevicePaused', 1);
  Module['_SDL_GetAudioDeviceGain'] = _SDL_GetAudioDeviceGain = createExportWrapper('SDL_GetAudioDeviceGain', 1);
  Module['_SDL_SetAudioDeviceGain'] = _SDL_SetAudioDeviceGain = createExportWrapper('SDL_SetAudioDeviceGain', 2);
  Module['_SDL_SetAudioPostmixCallback'] = _SDL_SetAudioPostmixCallback = createExportWrapper('SDL_SetAudioPostmixCallback', 3);
  Module['_SDL_BindAudioStreams'] = _SDL_BindAudioStreams = createExportWrapper('SDL_BindAudioStreams', 3);
  Module['_SDL_UnbindAudioStreams'] = _SDL_UnbindAudioStreams = createExportWrapper('SDL_UnbindAudioStreams', 2);
  Module['_SDL_UnbindAudioStream'] = _SDL_UnbindAudioStream = createExportWrapper('SDL_UnbindAudioStream', 1);
  Module['_SDL_GetAudioStreamDevice'] = _SDL_GetAudioStreamDevice = createExportWrapper('SDL_GetAudioStreamDevice', 1);
  Module['_SDL_OpenAudioDeviceStream'] = _SDL_OpenAudioDeviceStream = createExportWrapper('SDL_OpenAudioDeviceStream', 4);
  Module['_SDL_SetAudioStreamPutCallback'] = _SDL_SetAudioStreamPutCallback = createExportWrapper('SDL_SetAudioStreamPutCallback', 3);
  Module['_SDL_PauseAudioStreamDevice'] = _SDL_PauseAudioStreamDevice = createExportWrapper('SDL_PauseAudioStreamDevice', 1);
  Module['_SDL_ResumeAudioStreamDevice'] = _SDL_ResumeAudioStreamDevice = createExportWrapper('SDL_ResumeAudioStreamDevice', 1);
  Module['_SDL_AudioStreamDevicePaused'] = _SDL_AudioStreamDevicePaused = createExportWrapper('SDL_AudioStreamDevicePaused', 1);
  Module['_SDL_GetAudioFormatName'] = _SDL_GetAudioFormatName = createExportWrapper('SDL_GetAudioFormatName', 1);
  Module['_SDL_EventEnabled'] = _SDL_EventEnabled = createExportWrapper('SDL_EventEnabled', 1);
  Module['_SDL_PushEvent'] = _SDL_PushEvent = createExportWrapper('SDL_PushEvent', 1);
  Module['_SDL_CreateCondition'] = _SDL_CreateCondition = createExportWrapper('SDL_CreateCondition', 0);
  Module['_SDL_SetCurrentThreadPriority'] = _SDL_SetCurrentThreadPriority = createExportWrapper('SDL_SetCurrentThreadPriority', 1);
  Module['_SDL_SetAudioStreamFormat'] = _SDL_SetAudioStreamFormat = createExportWrapper('SDL_SetAudioStreamFormat', 3);
  Module['_SDL_GetAudioStreamProperties'] = _SDL_GetAudioStreamProperties = createExportWrapper('SDL_GetAudioStreamProperties', 1);
  Module['_SDL_GetAudioStreamFormat'] = _SDL_GetAudioStreamFormat = createExportWrapper('SDL_GetAudioStreamFormat', 3);
  Module['_SDL_SetAudioStreamInputChannelMap'] = _SDL_SetAudioStreamInputChannelMap = createExportWrapper('SDL_SetAudioStreamInputChannelMap', 3);
  Module['_SDL_SetAudioStreamOutputChannelMap'] = _SDL_SetAudioStreamOutputChannelMap = createExportWrapper('SDL_SetAudioStreamOutputChannelMap', 3);
  Module['_SDL_GetAudioStreamInputChannelMap'] = _SDL_GetAudioStreamInputChannelMap = createExportWrapper('SDL_GetAudioStreamInputChannelMap', 2);
  Module['_SDL_GetAudioStreamOutputChannelMap'] = _SDL_GetAudioStreamOutputChannelMap = createExportWrapper('SDL_GetAudioStreamOutputChannelMap', 2);
  Module['_SDL_GetAudioStreamFrequencyRatio'] = _SDL_GetAudioStreamFrequencyRatio = createExportWrapper('SDL_GetAudioStreamFrequencyRatio', 1);
  Module['_SDL_SetAudioStreamFrequencyRatio'] = _SDL_SetAudioStreamFrequencyRatio = createExportWrapper('SDL_SetAudioStreamFrequencyRatio', 2);
  Module['_SDL_GetAudioStreamGain'] = _SDL_GetAudioStreamGain = createExportWrapper('SDL_GetAudioStreamGain', 1);
  Module['_SDL_SetAudioStreamGain'] = _SDL_SetAudioStreamGain = createExportWrapper('SDL_SetAudioStreamGain', 2);
  Module['_SDL_GetAudioStreamAvailable'] = _SDL_GetAudioStreamAvailable = createExportWrapper('SDL_GetAudioStreamAvailable', 1);
  Module['_SDL_GetAudioStreamQueued'] = _SDL_GetAudioStreamQueued = createExportWrapper('SDL_GetAudioStreamQueued', 1);
  Module['_SDL_LoadWAV'] = _SDL_LoadWAV = createExportWrapper('SDL_LoadWAV', 4);
  Module['_SDL_GetNumCameraDrivers'] = _SDL_GetNumCameraDrivers = createExportWrapper('SDL_GetNumCameraDrivers', 0);
  Module['_SDL_GetCameraDriver'] = _SDL_GetCameraDriver = createExportWrapper('SDL_GetCameraDriver', 1);
  Module['_SDL_GetCurrentCameraDriver'] = _SDL_GetCurrentCameraDriver = createExportWrapper('SDL_GetCurrentCameraDriver', 0);
  Module['_SDL_GetTicksNS'] = _SDL_GetTicksNS = createExportWrapper('SDL_GetTicksNS', 0);
  Module['_SDL_CloseCamera'] = _SDL_CloseCamera = createExportWrapper('SDL_CloseCamera', 1);
  Module['_SDL_GetCameraFormat'] = _SDL_GetCameraFormat = createExportWrapper('SDL_GetCameraFormat', 2);
  Module['_SDL_GetCameraName'] = _SDL_GetCameraName = createExportWrapper('SDL_GetCameraName', 1);
  Module['_SDL_GetCameraPosition'] = _SDL_GetCameraPosition = createExportWrapper('SDL_GetCameraPosition', 1);
  Module['_SDL_GetCameras'] = _SDL_GetCameras = createExportWrapper('SDL_GetCameras', 1);
  Module['_SDL_GetCameraSupportedFormats'] = _SDL_GetCameraSupportedFormats = createExportWrapper('SDL_GetCameraSupportedFormats', 2);
  Module['_SDL_StretchSurface'] = _SDL_StretchSurface = createExportWrapper('SDL_StretchSurface', 5);
  Module['_SDL_ConvertPixels'] = _SDL_ConvertPixels = createExportWrapper('SDL_ConvertPixels', 8);
  Module['_SDL_SetSurfaceColorspace'] = _SDL_SetSurfaceColorspace = createExportWrapper('SDL_SetSurfaceColorspace', 2);
  Module['_SDL_OpenCamera'] = _SDL_OpenCamera = createExportWrapper('SDL_OpenCamera', 2);
  Module['_SDL_AcquireCameraFrame'] = _SDL_AcquireCameraFrame = createExportWrapper('SDL_AcquireCameraFrame', 2);
  Module['_SDL_ReleaseCameraFrame'] = _SDL_ReleaseCameraFrame = createExportWrapper('SDL_ReleaseCameraFrame', 2);
  Module['_SDL_GetCameraID'] = _SDL_GetCameraID = createExportWrapper('SDL_GetCameraID', 1);
  Module['_SDL_GetCameraProperties'] = _SDL_GetCameraProperties = createExportWrapper('SDL_GetCameraProperties', 1);
  Module['_SDL_GetCameraPermissionState'] = _SDL_GetCameraPermissionState = createExportWrapper('SDL_GetCameraPermissionState', 1);
  Module['_SDL_SetX11EventHook'] = _SDL_SetX11EventHook = createExportWrapper('SDL_SetX11EventHook', 2);
  Module['_SDL_SetLinuxThreadPriority'] = _SDL_SetLinuxThreadPriority = createExportWrapper('SDL_SetLinuxThreadPriority', 2);
  Module['_SDL_SetLinuxThreadPriorityAndPolicy'] = _SDL_SetLinuxThreadPriorityAndPolicy = createExportWrapper('SDL_SetLinuxThreadPriorityAndPolicy', 3);
  Module['_SDL_GDKSuspendComplete'] = _SDL_GDKSuspendComplete = createExportWrapper('SDL_GDKSuspendComplete', 0);
  Module['_SDL_GetGDKDefaultUser'] = _SDL_GetGDKDefaultUser = createExportWrapper('SDL_GetGDKDefaultUser', 1);
  Module['_SDL_GDKSuspendGPU'] = _SDL_GDKSuspendGPU = createExportWrapper('SDL_GDKSuspendGPU', 1);
  Module['_SDL_GDKResumeGPU'] = _SDL_GDKResumeGPU = createExportWrapper('SDL_GDKResumeGPU', 1);
  Module['_SDL_RegisterApp'] = _SDL_RegisterApp = createExportWrapper('SDL_RegisterApp', 3);
  Module['_SDL_SetWindowsMessageHook'] = _SDL_SetWindowsMessageHook = createExportWrapper('SDL_SetWindowsMessageHook', 2);
  Module['_SDL_UnregisterApp'] = _SDL_UnregisterApp = createExportWrapper('SDL_UnregisterApp', 0);
  Module['_SDL_SendAndroidBackButton'] = _SDL_SendAndroidBackButton = createExportWrapper('SDL_SendAndroidBackButton', 0);
  Module['_SDL_GetAndroidActivity'] = _SDL_GetAndroidActivity = createExportWrapper('SDL_GetAndroidActivity', 0);
  Module['_SDL_GetAndroidCachePath'] = _SDL_GetAndroidCachePath = createExportWrapper('SDL_GetAndroidCachePath', 0);
  Module['_SDL_GetAndroidExternalStoragePath'] = _SDL_GetAndroidExternalStoragePath = createExportWrapper('SDL_GetAndroidExternalStoragePath', 0);
  Module['_SDL_GetAndroidExternalStorageState'] = _SDL_GetAndroidExternalStorageState = createExportWrapper('SDL_GetAndroidExternalStorageState', 0);
  Module['_SDL_GetAndroidInternalStoragePath'] = _SDL_GetAndroidInternalStoragePath = createExportWrapper('SDL_GetAndroidInternalStoragePath', 0);
  Module['_SDL_GetAndroidJNIEnv'] = _SDL_GetAndroidJNIEnv = createExportWrapper('SDL_GetAndroidJNIEnv', 0);
  Module['_SDL_RequestAndroidPermission'] = _SDL_RequestAndroidPermission = createExportWrapper('SDL_RequestAndroidPermission', 3);
  Module['_SDL_SendAndroidMessage'] = _SDL_SendAndroidMessage = createExportWrapper('SDL_SendAndroidMessage', 2);
  Module['_SDL_ShowAndroidToast'] = _SDL_ShowAndroidToast = createExportWrapper('SDL_ShowAndroidToast', 5);
  Module['_SDL_GetAndroidSDKVersion'] = _SDL_GetAndroidSDKVersion = createExportWrapper('SDL_GetAndroidSDKVersion', 0);
  Module['_SDL_IsChromebook'] = _SDL_IsChromebook = createExportWrapper('SDL_IsChromebook', 0);
  Module['_SDL_IsDeXMode'] = _SDL_IsDeXMode = createExportWrapper('SDL_IsDeXMode', 0);
  Module['_JNI_OnLoad'] = _JNI_OnLoad = createExportWrapper('JNI_OnLoad', 2);
  Module['_SDL_GetNumLogicalCPUCores'] = _SDL_GetNumLogicalCPUCores = createExportWrapper('SDL_GetNumLogicalCPUCores', 0);
  Module['_SDL_GetCPUCacheLineSize'] = _SDL_GetCPUCacheLineSize = createExportWrapper('SDL_GetCPUCacheLineSize', 0);
  Module['_SDL_HasAltiVec'] = _SDL_HasAltiVec = createExportWrapper('SDL_HasAltiVec', 0);
  Module['_SDL_HasMMX'] = _SDL_HasMMX = createExportWrapper('SDL_HasMMX', 0);
  Module['_SDL_HasSSE'] = _SDL_HasSSE = createExportWrapper('SDL_HasSSE', 0);
  Module['_SDL_HasSSE2'] = _SDL_HasSSE2 = createExportWrapper('SDL_HasSSE2', 0);
  Module['_SDL_HasSSE3'] = _SDL_HasSSE3 = createExportWrapper('SDL_HasSSE3', 0);
  Module['_SDL_HasSSE41'] = _SDL_HasSSE41 = createExportWrapper('SDL_HasSSE41', 0);
  Module['_SDL_HasSSE42'] = _SDL_HasSSE42 = createExportWrapper('SDL_HasSSE42', 0);
  Module['_SDL_HasAVX'] = _SDL_HasAVX = createExportWrapper('SDL_HasAVX', 0);
  Module['_SDL_HasAVX2'] = _SDL_HasAVX2 = createExportWrapper('SDL_HasAVX2', 0);
  Module['_SDL_HasAVX512F'] = _SDL_HasAVX512F = createExportWrapper('SDL_HasAVX512F', 0);
  Module['_SDL_HasARMSIMD'] = _SDL_HasARMSIMD = createExportWrapper('SDL_HasARMSIMD', 0);
  Module['_SDL_HasNEON'] = _SDL_HasNEON = createExportWrapper('SDL_HasNEON', 0);
  Module['_SDL_HasLSX'] = _SDL_HasLSX = createExportWrapper('SDL_HasLSX', 0);
  Module['_SDL_HasLASX'] = _SDL_HasLASX = createExportWrapper('SDL_HasLASX', 0);
  Module['_SDL_GetSystemRAM'] = _SDL_GetSystemRAM = createExportWrapper('SDL_GetSystemRAM', 0);
  Module['_SDL_GetWindowFromEvent'] = _SDL_GetWindowFromEvent = createExportWrapper('SDL_GetWindowFromEvent', 1);
  Module['_SDL_GetWindowFromID'] = _SDL_GetWindowFromID = createExportWrapper('SDL_GetWindowFromID', 1);
  Module['_SDL_GetCurrentVideoDriver'] = _SDL_GetCurrentVideoDriver = createExportWrapper('SDL_GetCurrentVideoDriver', 0);
  Module['_SDL_PeepEvents'] = _SDL_PeepEvents = createExportWrapper('SDL_PeepEvents', 5);
  Module['_SDL_HasEvent'] = _SDL_HasEvent = createExportWrapper('SDL_HasEvent', 1);
  Module['_SDL_HasEvents'] = _SDL_HasEvents = createExportWrapper('SDL_HasEvents', 2);
  Module['_SDL_FlushEvent'] = _SDL_FlushEvent = createExportWrapper('SDL_FlushEvent', 1);
  Module['_SDL_FlushEvents'] = _SDL_FlushEvents = createExportWrapper('SDL_FlushEvents', 2);
  Module['_SDL_RunOnMainThread'] = _SDL_RunOnMainThread = createExportWrapper('SDL_RunOnMainThread', 3);
  Module['_SDL_CreateSemaphore'] = _SDL_CreateSemaphore = createExportWrapper('SDL_CreateSemaphore', 1);
  Module['_SDL_WaitSemaphore'] = _SDL_WaitSemaphore = createExportWrapper('SDL_WaitSemaphore', 1);
  Module['_SDL_DestroySemaphore'] = _SDL_DestroySemaphore = createExportWrapper('SDL_DestroySemaphore', 1);
  Module['_SDL_UpdateSensors'] = _SDL_UpdateSensors = createExportWrapper('SDL_UpdateSensors', 0);
  Module['_SDL_UpdateJoysticks'] = _SDL_UpdateJoysticks = createExportWrapper('SDL_UpdateJoysticks', 0);
  Module['_SDL_UpdateTrays'] = _SDL_UpdateTrays = createExportWrapper('SDL_UpdateTrays', 0);
  Module['_SDL_PumpEvents'] = _SDL_PumpEvents = createExportWrapper('SDL_PumpEvents', 0);
  Module['_SDL_SignalSemaphore'] = _SDL_SignalSemaphore = createExportWrapper('SDL_SignalSemaphore', 1);
  Module['_SDL_PollEvent'] = _SDL_PollEvent = createExportWrapper('SDL_PollEvent', 1);
  Module['_SDL_DelayNS'] = _SDL_DelayNS = createExportWrapper('SDL_DelayNS', 1);
  Module['_SDL_WaitEvent'] = _SDL_WaitEvent = createExportWrapper('SDL_WaitEvent', 1);
  Module['_SDL_WaitEventTimeout'] = _SDL_WaitEventTimeout = createExportWrapper('SDL_WaitEventTimeout', 2);
  Module['_SDL_SetEventFilter'] = _SDL_SetEventFilter = createExportWrapper('SDL_SetEventFilter', 2);
  Module['_SDL_GetEventFilter'] = _SDL_GetEventFilter = createExportWrapper('SDL_GetEventFilter', 2);
  Module['_SDL_AddEventWatch'] = _SDL_AddEventWatch = createExportWrapper('SDL_AddEventWatch', 2);
  Module['_SDL_RemoveEventWatch'] = _SDL_RemoveEventWatch = createExportWrapper('SDL_RemoveEventWatch', 2);
  Module['_SDL_FilterEvents'] = _SDL_FilterEvents = createExportWrapper('SDL_FilterEvents', 2);
  Module['_SDL_SetEventEnabled'] = _SDL_SetEventEnabled = createExportWrapper('SDL_SetEventEnabled', 2);
  Module['_SDL_RegisterEvents'] = _SDL_RegisterEvents = createExportWrapper('SDL_RegisterEvents', 1);
  Module['_SDL_HasKeyboard'] = _SDL_HasKeyboard = createExportWrapper('SDL_HasKeyboard', 0);
  Module['_SDL_GetKeyboards'] = _SDL_GetKeyboards = createExportWrapper('SDL_GetKeyboards', 1);
  Module['_SDL_GetKeyboardNameForID'] = _SDL_GetKeyboardNameForID = createExportWrapper('SDL_GetKeyboardNameForID', 1);
  Module['_SDL_ResetKeyboard'] = _SDL_ResetKeyboard = createExportWrapper('SDL_ResetKeyboard', 0);
  Module['_SDL_GetKeyboardFocus'] = _SDL_GetKeyboardFocus = createExportWrapper('SDL_GetKeyboardFocus', 0);
  Module['_SDL_WarpMouseGlobal'] = _SDL_WarpMouseGlobal = createExportWrapper('SDL_WarpMouseGlobal', 2);
  Module['_SDL_TextInputActive'] = _SDL_TextInputActive = createExportWrapper('SDL_TextInputActive', 1);
  Module['_SDL_GetKeyFromScancode'] = _SDL_GetKeyFromScancode = createExportWrapper('SDL_GetKeyFromScancode', 3);
  Module['_SDL_GetScancodeFromKey'] = _SDL_GetScancodeFromKey = createExportWrapper('SDL_GetScancodeFromKey', 2);
  Module['_SDL_GetModState'] = _SDL_GetModState = createExportWrapper('SDL_GetModState', 0);
  Module['_SDL_iscntrl'] = _SDL_iscntrl = createExportWrapper('SDL_iscntrl', 1);
  Module['_SDL_GetKeyboardState'] = _SDL_GetKeyboardState = createExportWrapper('SDL_GetKeyboardState', 1);
  Module['_SDL_SetModState'] = _SDL_SetModState = createExportWrapper('SDL_SetModState', 1);
  Module['_SDL_SetScancodeName'] = _SDL_SetScancodeName = createExportWrapper('SDL_SetScancodeName', 2);
  Module['_SDL_GetScancodeName'] = _SDL_GetScancodeName = createExportWrapper('SDL_GetScancodeName', 1);
  Module['_SDL_GetScancodeFromName'] = _SDL_GetScancodeFromName = createExportWrapper('SDL_GetScancodeFromName', 1);
  Module['_SDL_GetKeyName'] = _SDL_GetKeyName = createExportWrapper('SDL_GetKeyName', 1);
  Module['_SDL_GetKeyFromName'] = _SDL_GetKeyFromName = createExportWrapper('SDL_GetKeyFromName', 1);
  Module['_SDL_CreateColorCursor'] = _SDL_CreateColorCursor = createExportWrapper('SDL_CreateColorCursor', 3);
  Module['_SDL_HasMouse'] = _SDL_HasMouse = createExportWrapper('SDL_HasMouse', 0);
  Module['_SDL_GetMice'] = _SDL_GetMice = createExportWrapper('SDL_GetMice', 1);
  Module['_SDL_GetMouseNameForID'] = _SDL_GetMouseNameForID = createExportWrapper('SDL_GetMouseNameForID', 1);
  Module['_SDL_SetCursor'] = _SDL_SetCursor = createExportWrapper('SDL_SetCursor', 1);
  Module['_SDL_GetMouseFocus'] = _SDL_GetMouseFocus = createExportWrapper('SDL_GetMouseFocus', 0);
  Module['_SDL_modff'] = _SDL_modff = createExportWrapper('SDL_modff', 2);
  Module['_SDL_truncf'] = _SDL_truncf = createExportWrapper('SDL_truncf', 1);
  Module['_SDL_DestroyCursor'] = _SDL_DestroyCursor = createExportWrapper('SDL_DestroyCursor', 1);
  Module['_SDL_CaptureMouse'] = _SDL_CaptureMouse = createExportWrapper('SDL_CaptureMouse', 1);
  Module['_SDL_ShowCursor'] = _SDL_ShowCursor = createExportWrapper('SDL_ShowCursor', 0);
  Module['_SDL_GetMouseState'] = _SDL_GetMouseState = createExportWrapper('SDL_GetMouseState', 2);
  Module['_SDL_GetRelativeMouseState'] = _SDL_GetRelativeMouseState = createExportWrapper('SDL_GetRelativeMouseState', 2);
  Module['_SDL_GetGlobalMouseState'] = _SDL_GetGlobalMouseState = createExportWrapper('SDL_GetGlobalMouseState', 2);
  Module['_SDL_WarpMouseInWindow'] = _SDL_WarpMouseInWindow = createExportWrapper('SDL_WarpMouseInWindow', 3);
  Module['_SDL_CreateCursor'] = _SDL_CreateCursor = createExportWrapper('SDL_CreateCursor', 6);
  Module['_SDL_CreateSystemCursor'] = _SDL_CreateSystemCursor = createExportWrapper('SDL_CreateSystemCursor', 1);
  Module['_SDL_GetCursor'] = _SDL_GetCursor = createExportWrapper('SDL_GetCursor', 0);
  Module['_SDL_GetDefaultCursor'] = _SDL_GetDefaultCursor = createExportWrapper('SDL_GetDefaultCursor', 0);
  Module['_SDL_HideCursor'] = _SDL_HideCursor = createExportWrapper('SDL_HideCursor', 0);
  Module['_SDL_CursorVisible'] = _SDL_CursorVisible = createExportWrapper('SDL_CursorVisible', 0);
  Module['_SDL_GetWindowMouseRect'] = _SDL_GetWindowMouseRect = createExportWrapper('SDL_GetWindowMouseRect', 1);
  Module['_SDL_GetRectIntersection'] = _SDL_GetRectIntersection = createExportWrapper('SDL_GetRectIntersection', 3);
  Module['_SDL_GetTouchDevices'] = _SDL_GetTouchDevices = createExportWrapper('SDL_GetTouchDevices', 1);
  Module['_SDL_GetTouchDeviceName'] = _SDL_GetTouchDeviceName = createExportWrapper('SDL_GetTouchDeviceName', 1);
  Module['_SDL_GetTouchDeviceType'] = _SDL_GetTouchDeviceType = createExportWrapper('SDL_GetTouchDeviceType', 1);
  Module['_SDL_GetTouchFingers'] = _SDL_GetTouchFingers = createExportWrapper('SDL_GetTouchFingers', 2);
  Module['_SDL_GetWindowID'] = _SDL_GetWindowID = createExportWrapper('SDL_GetWindowID', 1);
  Module['_SDL_RemovePath'] = _SDL_RemovePath = createExportWrapper('SDL_RemovePath', 1);
  Module['_SDL_RenamePath'] = _SDL_RenamePath = createExportWrapper('SDL_RenamePath', 2);
  Module['_SDL_CopyFile'] = _SDL_CopyFile = createExportWrapper('SDL_CopyFile', 2);
  Module['_SDL_CreateDirectory'] = _SDL_CreateDirectory = createExportWrapper('SDL_CreateDirectory', 1);
  Module['_SDL_EnumerateDirectory'] = _SDL_EnumerateDirectory = createExportWrapper('SDL_EnumerateDirectory', 3);
  Module['_SDL_GetPathInfo'] = _SDL_GetPathInfo = createExportWrapper('SDL_GetPathInfo', 2);
  Module['_SDL_IOFromDynamicMem'] = _SDL_IOFromDynamicMem = createExportWrapper('SDL_IOFromDynamicMem', 0);
  Module['_SDL_GlobDirectory'] = _SDL_GlobDirectory = createExportWrapper('SDL_GlobDirectory', 4);
  Module['_SDL_GetBasePath'] = _SDL_GetBasePath = createExportWrapper('SDL_GetBasePath', 0);
  Module['_SDL_GetUserFolder'] = _SDL_GetUserFolder = createExportWrapper('SDL_GetUserFolder', 1);
  Module['_SDL_GetPrefPath'] = _SDL_GetPrefPath = createExportWrapper('SDL_GetPrefPath', 2);
  Module['_SDL_GetCurrentDirectory'] = _SDL_GetCurrentDirectory = createExportWrapper('SDL_GetCurrentDirectory', 0);
  Module['_SDL_CreateGPUGraphicsPipeline'] = _SDL_CreateGPUGraphicsPipeline = createExportWrapper('SDL_CreateGPUGraphicsPipeline', 2);
  Module['_SDL_GPUTextureSupportsFormat'] = _SDL_GPUTextureSupportsFormat = createExportWrapper('SDL_GPUTextureSupportsFormat', 4);
  Module['_SDL_BindGPUFragmentSamplers'] = _SDL_BindGPUFragmentSamplers = createExportWrapper('SDL_BindGPUFragmentSamplers', 4);
  Module['_SDL_DrawGPUPrimitives'] = _SDL_DrawGPUPrimitives = createExportWrapper('SDL_DrawGPUPrimitives', 5);
  Module['_SDL_SetGPUViewport'] = _SDL_SetGPUViewport = createExportWrapper('SDL_SetGPUViewport', 2);
  Module['_SDL_BindGPUGraphicsPipeline'] = _SDL_BindGPUGraphicsPipeline = createExportWrapper('SDL_BindGPUGraphicsPipeline', 2);
  Module['_SDL_PushGPUFragmentUniformData'] = _SDL_PushGPUFragmentUniformData = createExportWrapper('SDL_PushGPUFragmentUniformData', 4);
  Module['_SDL_GPUSupportsShaderFormats'] = _SDL_GPUSupportsShaderFormats = createExportWrapper('SDL_GPUSupportsShaderFormats', 2);
  Module['_SDL_GPUSupportsProperties'] = _SDL_GPUSupportsProperties = createExportWrapper('SDL_GPUSupportsProperties', 1);
  Module['_SDL_CreateGPUDevice'] = _SDL_CreateGPUDevice = createExportWrapper('SDL_CreateGPUDevice', 3);
  Module['_SDL_CreateGPUDeviceWithProperties'] = _SDL_CreateGPUDeviceWithProperties = createExportWrapper('SDL_CreateGPUDeviceWithProperties', 1);
  Module['_SDL_DestroyGPUDevice'] = _SDL_DestroyGPUDevice = createExportWrapper('SDL_DestroyGPUDevice', 1);
  Module['_SDL_GetNumGPUDrivers'] = _SDL_GetNumGPUDrivers = createExportWrapper('SDL_GetNumGPUDrivers', 0);
  Module['_SDL_GetGPUDriver'] = _SDL_GetGPUDriver = createExportWrapper('SDL_GetGPUDriver', 1);
  Module['_SDL_GetGPUDeviceDriver'] = _SDL_GetGPUDeviceDriver = createExportWrapper('SDL_GetGPUDeviceDriver', 1);
  Module['_SDL_GetGPUShaderFormats'] = _SDL_GetGPUShaderFormats = createExportWrapper('SDL_GetGPUShaderFormats', 1);
  Module['_SDL_GPUTextureFormatTexelBlockSize'] = _SDL_GPUTextureFormatTexelBlockSize = createExportWrapper('SDL_GPUTextureFormatTexelBlockSize', 1);
  Module['_SDL_GPUTextureSupportsSampleCount'] = _SDL_GPUTextureSupportsSampleCount = createExportWrapper('SDL_GPUTextureSupportsSampleCount', 3);
  Module['_SDL_CreateGPUComputePipeline'] = _SDL_CreateGPUComputePipeline = createExportWrapper('SDL_CreateGPUComputePipeline', 2);
  Module['_SDL_CreateGPUSampler'] = _SDL_CreateGPUSampler = createExportWrapper('SDL_CreateGPUSampler', 2);
  Module['_SDL_CreateGPUShader'] = _SDL_CreateGPUShader = createExportWrapper('SDL_CreateGPUShader', 2);
  Module['_SDL_CreateGPUBuffer'] = _SDL_CreateGPUBuffer = createExportWrapper('SDL_CreateGPUBuffer', 2);
  Module['_SDL_SetGPUBufferName'] = _SDL_SetGPUBufferName = createExportWrapper('SDL_SetGPUBufferName', 3);
  Module['_SDL_SetGPUTextureName'] = _SDL_SetGPUTextureName = createExportWrapper('SDL_SetGPUTextureName', 3);
  Module['_SDL_InsertGPUDebugLabel'] = _SDL_InsertGPUDebugLabel = createExportWrapper('SDL_InsertGPUDebugLabel', 2);
  Module['_SDL_PushGPUDebugGroup'] = _SDL_PushGPUDebugGroup = createExportWrapper('SDL_PushGPUDebugGroup', 2);
  Module['_SDL_PopGPUDebugGroup'] = _SDL_PopGPUDebugGroup = createExportWrapper('SDL_PopGPUDebugGroup', 1);
  Module['_SDL_ReleaseGPUSampler'] = _SDL_ReleaseGPUSampler = createExportWrapper('SDL_ReleaseGPUSampler', 2);
  Module['_SDL_ReleaseGPUBuffer'] = _SDL_ReleaseGPUBuffer = createExportWrapper('SDL_ReleaseGPUBuffer', 2);
  Module['_SDL_ReleaseGPUShader'] = _SDL_ReleaseGPUShader = createExportWrapper('SDL_ReleaseGPUShader', 2);
  Module['_SDL_ReleaseGPUComputePipeline'] = _SDL_ReleaseGPUComputePipeline = createExportWrapper('SDL_ReleaseGPUComputePipeline', 2);
  Module['_SDL_ReleaseGPUGraphicsPipeline'] = _SDL_ReleaseGPUGraphicsPipeline = createExportWrapper('SDL_ReleaseGPUGraphicsPipeline', 2);
  Module['_SDL_PushGPUVertexUniformData'] = _SDL_PushGPUVertexUniformData = createExportWrapper('SDL_PushGPUVertexUniformData', 4);
  Module['_SDL_PushGPUComputeUniformData'] = _SDL_PushGPUComputeUniformData = createExportWrapper('SDL_PushGPUComputeUniformData', 4);
  Module['_SDL_SetGPUScissor'] = _SDL_SetGPUScissor = createExportWrapper('SDL_SetGPUScissor', 2);
  Module['_SDL_SetGPUBlendConstants'] = _SDL_SetGPUBlendConstants = createExportWrapper('SDL_SetGPUBlendConstants', 2);
  Module['_SDL_SetGPUStencilReference'] = _SDL_SetGPUStencilReference = createExportWrapper('SDL_SetGPUStencilReference', 2);
  Module['_SDL_BindGPUVertexBuffers'] = _SDL_BindGPUVertexBuffers = createExportWrapper('SDL_BindGPUVertexBuffers', 4);
  Module['_SDL_BindGPUIndexBuffer'] = _SDL_BindGPUIndexBuffer = createExportWrapper('SDL_BindGPUIndexBuffer', 3);
  Module['_SDL_BindGPUVertexSamplers'] = _SDL_BindGPUVertexSamplers = createExportWrapper('SDL_BindGPUVertexSamplers', 4);
  Module['_SDL_BindGPUVertexStorageTextures'] = _SDL_BindGPUVertexStorageTextures = createExportWrapper('SDL_BindGPUVertexStorageTextures', 4);
  Module['_SDL_BindGPUVertexStorageBuffers'] = _SDL_BindGPUVertexStorageBuffers = createExportWrapper('SDL_BindGPUVertexStorageBuffers', 4);
  Module['_SDL_BindGPUFragmentStorageTextures'] = _SDL_BindGPUFragmentStorageTextures = createExportWrapper('SDL_BindGPUFragmentStorageTextures', 4);
  Module['_SDL_BindGPUFragmentStorageBuffers'] = _SDL_BindGPUFragmentStorageBuffers = createExportWrapper('SDL_BindGPUFragmentStorageBuffers', 4);
  Module['_SDL_DrawGPUIndexedPrimitives'] = _SDL_DrawGPUIndexedPrimitives = createExportWrapper('SDL_DrawGPUIndexedPrimitives', 6);
  Module['_SDL_DrawGPUPrimitivesIndirect'] = _SDL_DrawGPUPrimitivesIndirect = createExportWrapper('SDL_DrawGPUPrimitivesIndirect', 4);
  Module['_SDL_DrawGPUIndexedPrimitivesIndirect'] = _SDL_DrawGPUIndexedPrimitivesIndirect = createExportWrapper('SDL_DrawGPUIndexedPrimitivesIndirect', 4);
  Module['_SDL_BeginGPUComputePass'] = _SDL_BeginGPUComputePass = createExportWrapper('SDL_BeginGPUComputePass', 5);
  Module['_SDL_BindGPUComputePipeline'] = _SDL_BindGPUComputePipeline = createExportWrapper('SDL_BindGPUComputePipeline', 2);
  Module['_SDL_BindGPUComputeSamplers'] = _SDL_BindGPUComputeSamplers = createExportWrapper('SDL_BindGPUComputeSamplers', 4);
  Module['_SDL_BindGPUComputeStorageTextures'] = _SDL_BindGPUComputeStorageTextures = createExportWrapper('SDL_BindGPUComputeStorageTextures', 4);
  Module['_SDL_BindGPUComputeStorageBuffers'] = _SDL_BindGPUComputeStorageBuffers = createExportWrapper('SDL_BindGPUComputeStorageBuffers', 4);
  Module['_SDL_DispatchGPUCompute'] = _SDL_DispatchGPUCompute = createExportWrapper('SDL_DispatchGPUCompute', 4);
  Module['_SDL_DispatchGPUComputeIndirect'] = _SDL_DispatchGPUComputeIndirect = createExportWrapper('SDL_DispatchGPUComputeIndirect', 3);
  Module['_SDL_EndGPUComputePass'] = _SDL_EndGPUComputePass = createExportWrapper('SDL_EndGPUComputePass', 1);
  Module['_SDL_UploadToGPUBuffer'] = _SDL_UploadToGPUBuffer = createExportWrapper('SDL_UploadToGPUBuffer', 4);
  Module['_SDL_CopyGPUTextureToTexture'] = _SDL_CopyGPUTextureToTexture = createExportWrapper('SDL_CopyGPUTextureToTexture', 7);
  Module['_SDL_CopyGPUBufferToBuffer'] = _SDL_CopyGPUBufferToBuffer = createExportWrapper('SDL_CopyGPUBufferToBuffer', 5);
  Module['_SDL_DownloadFromGPUTexture'] = _SDL_DownloadFromGPUTexture = createExportWrapper('SDL_DownloadFromGPUTexture', 3);
  Module['_SDL_DownloadFromGPUBuffer'] = _SDL_DownloadFromGPUBuffer = createExportWrapper('SDL_DownloadFromGPUBuffer', 3);
  Module['_SDL_GenerateMipmapsForGPUTexture'] = _SDL_GenerateMipmapsForGPUTexture = createExportWrapper('SDL_GenerateMipmapsForGPUTexture', 2);
  Module['_SDL_BlitGPUTexture'] = _SDL_BlitGPUTexture = createExportWrapper('SDL_BlitGPUTexture', 2);
  Module['_SDL_WindowSupportsGPUSwapchainComposition'] = _SDL_WindowSupportsGPUSwapchainComposition = createExportWrapper('SDL_WindowSupportsGPUSwapchainComposition', 3);
  Module['_SDL_WindowSupportsGPUPresentMode'] = _SDL_WindowSupportsGPUPresentMode = createExportWrapper('SDL_WindowSupportsGPUPresentMode', 3);
  Module['_SDL_ClaimWindowForGPUDevice'] = _SDL_ClaimWindowForGPUDevice = createExportWrapper('SDL_ClaimWindowForGPUDevice', 2);
  Module['_SDL_ReleaseWindowFromGPUDevice'] = _SDL_ReleaseWindowFromGPUDevice = createExportWrapper('SDL_ReleaseWindowFromGPUDevice', 2);
  Module['_SDL_SetGPUSwapchainParameters'] = _SDL_SetGPUSwapchainParameters = createExportWrapper('SDL_SetGPUSwapchainParameters', 4);
  Module['_SDL_SetGPUAllowedFramesInFlight'] = _SDL_SetGPUAllowedFramesInFlight = createExportWrapper('SDL_SetGPUAllowedFramesInFlight', 2);
  Module['_SDL_GetGPUSwapchainTextureFormat'] = _SDL_GetGPUSwapchainTextureFormat = createExportWrapper('SDL_GetGPUSwapchainTextureFormat', 2);
  Module['_SDL_AcquireGPUSwapchainTexture'] = _SDL_AcquireGPUSwapchainTexture = createExportWrapper('SDL_AcquireGPUSwapchainTexture', 5);
  Module['_SDL_WaitForGPUSwapchain'] = _SDL_WaitForGPUSwapchain = createExportWrapper('SDL_WaitForGPUSwapchain', 2);
  Module['_SDL_WaitAndAcquireGPUSwapchainTexture'] = _SDL_WaitAndAcquireGPUSwapchainTexture = createExportWrapper('SDL_WaitAndAcquireGPUSwapchainTexture', 5);
  Module['_SDL_SubmitGPUCommandBufferAndAcquireFence'] = _SDL_SubmitGPUCommandBufferAndAcquireFence = createExportWrapper('SDL_SubmitGPUCommandBufferAndAcquireFence', 1);
  Module['_SDL_CancelGPUCommandBuffer'] = _SDL_CancelGPUCommandBuffer = createExportWrapper('SDL_CancelGPUCommandBuffer', 1);
  Module['_SDL_WaitForGPUIdle'] = _SDL_WaitForGPUIdle = createExportWrapper('SDL_WaitForGPUIdle', 1);
  Module['_SDL_WaitForGPUFences'] = _SDL_WaitForGPUFences = createExportWrapper('SDL_WaitForGPUFences', 4);
  Module['_SDL_QueryGPUFence'] = _SDL_QueryGPUFence = createExportWrapper('SDL_QueryGPUFence', 2);
  Module['_SDL_ReleaseGPUFence'] = _SDL_ReleaseGPUFence = createExportWrapper('SDL_ReleaseGPUFence', 2);
  Module['_SDL_CalculateGPUTextureFormatSize'] = _SDL_CalculateGPUTextureFormatSize = createExportWrapper('SDL_CalculateGPUTextureFormatSize', 4);
  Module['_SDL_GetHaptics'] = _SDL_GetHaptics = createExportWrapper('SDL_GetHaptics', 1);
  Module['_SDL_GetHapticNameForID'] = _SDL_GetHapticNameForID = createExportWrapper('SDL_GetHapticNameForID', 1);
  Module['_SDL_OpenHaptic'] = _SDL_OpenHaptic = createExportWrapper('SDL_OpenHaptic', 1);
  Module['_SDL_SetHapticGain'] = _SDL_SetHapticGain = createExportWrapper('SDL_SetHapticGain', 2);
  Module['_SDL_SetHapticAutocenter'] = _SDL_SetHapticAutocenter = createExportWrapper('SDL_SetHapticAutocenter', 2);
  Module['_SDL_GetHapticFromID'] = _SDL_GetHapticFromID = createExportWrapper('SDL_GetHapticFromID', 1);
  Module['_SDL_GetHapticID'] = _SDL_GetHapticID = createExportWrapper('SDL_GetHapticID', 1);
  Module['_SDL_GetHapticName'] = _SDL_GetHapticName = createExportWrapper('SDL_GetHapticName', 1);
  Module['_SDL_IsMouseHaptic'] = _SDL_IsMouseHaptic = createExportWrapper('SDL_IsMouseHaptic', 0);
  Module['_SDL_OpenHapticFromMouse'] = _SDL_OpenHapticFromMouse = createExportWrapper('SDL_OpenHapticFromMouse', 0);
  Module['_SDL_IsJoystickHaptic'] = _SDL_IsJoystickHaptic = createExportWrapper('SDL_IsJoystickHaptic', 1);
  Module['_SDL_LockJoysticks'] = _SDL_LockJoysticks = createExportWrapper('SDL_LockJoysticks', 0);
  Module['_SDL_GetJoystickID'] = _SDL_GetJoystickID = createExportWrapper('SDL_GetJoystickID', 1);
  Module['_SDL_IsGamepad'] = _SDL_IsGamepad = createExportWrapper('SDL_IsGamepad', 1);
  Module['_SDL_UnlockJoysticks'] = _SDL_UnlockJoysticks = createExportWrapper('SDL_UnlockJoysticks', 0);
  Module['_SDL_OpenHapticFromJoystick'] = _SDL_OpenHapticFromJoystick = createExportWrapper('SDL_OpenHapticFromJoystick', 1);
  Module['_SDL_GetJoystickVendor'] = _SDL_GetJoystickVendor = createExportWrapper('SDL_GetJoystickVendor', 1);
  Module['_SDL_GetJoystickProduct'] = _SDL_GetJoystickProduct = createExportWrapper('SDL_GetJoystickProduct', 1);
  Module['_SDL_GetNumJoystickAxes'] = _SDL_GetNumJoystickAxes = createExportWrapper('SDL_GetNumJoystickAxes', 1);
  Module['_SDL_CloseHaptic'] = _SDL_CloseHaptic = createExportWrapper('SDL_CloseHaptic', 1);
  Module['_SDL_DestroyHapticEffect'] = _SDL_DestroyHapticEffect = createExportWrapper('SDL_DestroyHapticEffect', 2);
  Module['_SDL_GetMaxHapticEffects'] = _SDL_GetMaxHapticEffects = createExportWrapper('SDL_GetMaxHapticEffects', 1);
  Module['_SDL_GetMaxHapticEffectsPlaying'] = _SDL_GetMaxHapticEffectsPlaying = createExportWrapper('SDL_GetMaxHapticEffectsPlaying', 1);
  Module['_SDL_GetHapticFeatures'] = _SDL_GetHapticFeatures = createExportWrapper('SDL_GetHapticFeatures', 1);
  Module['_SDL_GetNumHapticAxes'] = _SDL_GetNumHapticAxes = createExportWrapper('SDL_GetNumHapticAxes', 1);
  Module['_SDL_HapticEffectSupported'] = _SDL_HapticEffectSupported = createExportWrapper('SDL_HapticEffectSupported', 2);
  Module['_SDL_CreateHapticEffect'] = _SDL_CreateHapticEffect = createExportWrapper('SDL_CreateHapticEffect', 2);
  Module['_SDL_UpdateHapticEffect'] = _SDL_UpdateHapticEffect = createExportWrapper('SDL_UpdateHapticEffect', 3);
  Module['_SDL_RunHapticEffect'] = _SDL_RunHapticEffect = createExportWrapper('SDL_RunHapticEffect', 3);
  Module['_SDL_StopHapticEffect'] = _SDL_StopHapticEffect = createExportWrapper('SDL_StopHapticEffect', 2);
  Module['_SDL_GetHapticEffectStatus'] = _SDL_GetHapticEffectStatus = createExportWrapper('SDL_GetHapticEffectStatus', 2);
  Module['_SDL_PauseHaptic'] = _SDL_PauseHaptic = createExportWrapper('SDL_PauseHaptic', 1);
  Module['_SDL_ResumeHaptic'] = _SDL_ResumeHaptic = createExportWrapper('SDL_ResumeHaptic', 1);
  Module['_SDL_StopHapticEffects'] = _SDL_StopHapticEffects = createExportWrapper('SDL_StopHapticEffects', 1);
  Module['_SDL_HapticRumbleSupported'] = _SDL_HapticRumbleSupported = createExportWrapper('SDL_HapticRumbleSupported', 1);
  Module['_SDL_InitHapticRumble'] = _SDL_InitHapticRumble = createExportWrapper('SDL_InitHapticRumble', 1);
  Module['_SDL_PlayHapticRumble'] = _SDL_PlayHapticRumble = createExportWrapper('SDL_PlayHapticRumble', 3);
  Module['_SDL_StopHapticRumble'] = _SDL_StopHapticRumble = createExportWrapper('SDL_StopHapticRumble', 1);
  Module['_SDL_strcasestr'] = _SDL_strcasestr = createExportWrapper('SDL_strcasestr', 2);
  Module['_SDL_hid_init'] = _SDL_hid_init = createExportWrapper('SDL_hid_init', 0);
  Module['_SDL_hid_exit'] = _SDL_hid_exit = createExportWrapper('SDL_hid_exit', 0);
  Module['_SDL_hid_device_change_count'] = _SDL_hid_device_change_count = createExportWrapper('SDL_hid_device_change_count', 0);
  Module['_SDL_hid_enumerate'] = _SDL_hid_enumerate = createExportWrapper('SDL_hid_enumerate', 2);
  Module['_SDL_hid_free_enumeration'] = _SDL_hid_free_enumeration = createExportWrapper('SDL_hid_free_enumeration', 1);
  Module['_SDL_hid_open'] = _SDL_hid_open = createExportWrapper('SDL_hid_open', 3);
  Module['_SDL_hid_open_path'] = _SDL_hid_open_path = createExportWrapper('SDL_hid_open_path', 1);
  Module['_SDL_hid_write'] = _SDL_hid_write = createExportWrapper('SDL_hid_write', 3);
  Module['_SDL_hid_read_timeout'] = _SDL_hid_read_timeout = createExportWrapper('SDL_hid_read_timeout', 4);
  Module['_SDL_hid_read'] = _SDL_hid_read = createExportWrapper('SDL_hid_read', 3);
  Module['_SDL_hid_set_nonblocking'] = _SDL_hid_set_nonblocking = createExportWrapper('SDL_hid_set_nonblocking', 2);
  Module['_SDL_hid_send_feature_report'] = _SDL_hid_send_feature_report = createExportWrapper('SDL_hid_send_feature_report', 3);
  Module['_SDL_hid_get_feature_report'] = _SDL_hid_get_feature_report = createExportWrapper('SDL_hid_get_feature_report', 3);
  Module['_SDL_hid_get_input_report'] = _SDL_hid_get_input_report = createExportWrapper('SDL_hid_get_input_report', 3);
  Module['_SDL_hid_close'] = _SDL_hid_close = createExportWrapper('SDL_hid_close', 1);
  Module['_SDL_hid_get_manufacturer_string'] = _SDL_hid_get_manufacturer_string = createExportWrapper('SDL_hid_get_manufacturer_string', 3);
  Module['_SDL_hid_get_product_string'] = _SDL_hid_get_product_string = createExportWrapper('SDL_hid_get_product_string', 3);
  Module['_SDL_hid_get_serial_number_string'] = _SDL_hid_get_serial_number_string = createExportWrapper('SDL_hid_get_serial_number_string', 3);
  Module['_SDL_hid_get_indexed_string'] = _SDL_hid_get_indexed_string = createExportWrapper('SDL_hid_get_indexed_string', 4);
  Module['_SDL_hid_get_device_info'] = _SDL_hid_get_device_info = createExportWrapper('SDL_hid_get_device_info', 1);
  Module['_SDL_wcsdup'] = _SDL_wcsdup = createExportWrapper('SDL_wcsdup', 1);
  Module['_SDL_hid_get_report_descriptor'] = _SDL_hid_get_report_descriptor = createExportWrapper('SDL_hid_get_report_descriptor', 3);
  Module['_SDL_hid_ble_scan'] = _SDL_hid_ble_scan = createExportWrapper('SDL_hid_ble_scan', 1);
  Module['_SDL_AsyncIOFromFile'] = _SDL_AsyncIOFromFile = createExportWrapper('SDL_AsyncIOFromFile', 2);
  Module['_SDL_GetAsyncIOSize'] = _SDL_GetAsyncIOSize = createExportWrapper('SDL_GetAsyncIOSize', 1);
  Module['_SDL_ReadAsyncIO'] = _SDL_ReadAsyncIO = createExportWrapper('SDL_ReadAsyncIO', 6);
  Module['_SDL_WriteAsyncIO'] = _SDL_WriteAsyncIO = createExportWrapper('SDL_WriteAsyncIO', 6);
  Module['_SDL_CloseAsyncIO'] = _SDL_CloseAsyncIO = createExportWrapper('SDL_CloseAsyncIO', 4);
  Module['_SDL_CreateAsyncIOQueue'] = _SDL_CreateAsyncIOQueue = createExportWrapper('SDL_CreateAsyncIOQueue', 0);
  Module['_SDL_GetAsyncIOResult'] = _SDL_GetAsyncIOResult = createExportWrapper('SDL_GetAsyncIOResult', 2);
  Module['_SDL_WaitAsyncIOResult'] = _SDL_WaitAsyncIOResult = createExportWrapper('SDL_WaitAsyncIOResult', 3);
  Module['_SDL_SignalAsyncIOQueue'] = _SDL_SignalAsyncIOQueue = createExportWrapper('SDL_SignalAsyncIOQueue', 1);
  Module['_SDL_DestroyAsyncIOQueue'] = _SDL_DestroyAsyncIOQueue = createExportWrapper('SDL_DestroyAsyncIOQueue', 1);
  Module['_SDL_LoadFileAsync'] = _SDL_LoadFileAsync = createExportWrapper('SDL_LoadFileAsync', 3);
  Module['_SDL_OpenIO'] = _SDL_OpenIO = createExportWrapper('SDL_OpenIO', 2);
  Module['_fileno'] = _fileno = createExportWrapper('fileno', 1);
  Module['_fflush'] = _fflush = createExportWrapper('fflush', 1);
  Module['_SDL_IOFromMem'] = _SDL_IOFromMem = createExportWrapper('SDL_IOFromMem', 2);
  Module['_SDL_LoadFile'] = _SDL_LoadFile = createExportWrapper('SDL_LoadFile', 2);
  Module['_SDL_SaveFile_IO'] = _SDL_SaveFile_IO = createExportWrapper('SDL_SaveFile_IO', 4);
  Module['_SDL_SaveFile'] = _SDL_SaveFile = createExportWrapper('SDL_SaveFile', 3);
  Module['_SDL_IOprintf'] = _SDL_IOprintf = createExportWrapper('SDL_IOprintf', 3);
  Module['_SDL_vasprintf'] = _SDL_vasprintf = createExportWrapper('SDL_vasprintf', 3);
  Module['_SDL_IOvprintf'] = _SDL_IOvprintf = createExportWrapper('SDL_IOvprintf', 3);
  Module['_SDL_FlushIO'] = _SDL_FlushIO = createExportWrapper('SDL_FlushIO', 1);
  Module['_SDL_ReadS8'] = _SDL_ReadS8 = createExportWrapper('SDL_ReadS8', 2);
  Module['_SDL_ReadS16LE'] = _SDL_ReadS16LE = createExportWrapper('SDL_ReadS16LE', 2);
  Module['_SDL_ReadU64LE'] = _SDL_ReadU64LE = createExportWrapper('SDL_ReadU64LE', 2);
  Module['_SDL_ReadS64LE'] = _SDL_ReadS64LE = createExportWrapper('SDL_ReadS64LE', 2);
  Module['_SDL_ReadU64BE'] = _SDL_ReadU64BE = createExportWrapper('SDL_ReadU64BE', 2);
  Module['_SDL_ReadS64BE'] = _SDL_ReadS64BE = createExportWrapper('SDL_ReadS64BE', 2);
  Module['_SDL_WriteU8'] = _SDL_WriteU8 = createExportWrapper('SDL_WriteU8', 2);
  Module['_SDL_WriteS8'] = _SDL_WriteS8 = createExportWrapper('SDL_WriteS8', 2);
  Module['_SDL_WriteU16LE'] = _SDL_WriteU16LE = createExportWrapper('SDL_WriteU16LE', 2);
  Module['_SDL_WriteS16LE'] = _SDL_WriteS16LE = createExportWrapper('SDL_WriteS16LE', 2);
  Module['_SDL_WriteU16BE'] = _SDL_WriteU16BE = createExportWrapper('SDL_WriteU16BE', 2);
  Module['_SDL_WriteS16BE'] = _SDL_WriteS16BE = createExportWrapper('SDL_WriteS16BE', 2);
  Module['_SDL_WriteU32LE'] = _SDL_WriteU32LE = createExportWrapper('SDL_WriteU32LE', 2);
  Module['_SDL_WriteS32LE'] = _SDL_WriteS32LE = createExportWrapper('SDL_WriteS32LE', 2);
  Module['_SDL_WriteU32BE'] = _SDL_WriteU32BE = createExportWrapper('SDL_WriteU32BE', 2);
  Module['_SDL_WriteS32BE'] = _SDL_WriteS32BE = createExportWrapper('SDL_WriteS32BE', 2);
  Module['_SDL_WriteU64LE'] = _SDL_WriteU64LE = createExportWrapper('SDL_WriteU64LE', 2);
  Module['_SDL_WriteS64LE'] = _SDL_WriteS64LE = createExportWrapper('SDL_WriteS64LE', 2);
  Module['_SDL_WriteU64BE'] = _SDL_WriteU64BE = createExportWrapper('SDL_WriteU64BE', 2);
  Module['_SDL_WriteS64BE'] = _SDL_WriteS64BE = createExportWrapper('SDL_WriteS64BE', 2);
  Module['_SDL_SignalCondition'] = _SDL_SignalCondition = createExportWrapper('SDL_SignalCondition', 1);
  Module['_SDL_WaitConditionTimeout'] = _SDL_WaitConditionTimeout = createExportWrapper('SDL_WaitConditionTimeout', 3);
  Module['_SDL_GetGamepadButton'] = _SDL_GetGamepadButton = createExportWrapper('SDL_GetGamepadButton', 2);
  Module['_SDL_GetGamepadAxis'] = _SDL_GetGamepadAxis = createExportWrapper('SDL_GetGamepadAxis', 2);
  Module['_SDL_GetGamepadTypeFromString'] = _SDL_GetGamepadTypeFromString = createExportWrapper('SDL_GetGamepadTypeFromString', 1);
  Module['_SDL_GetGamepadStringForType'] = _SDL_GetGamepadStringForType = createExportWrapper('SDL_GetGamepadStringForType', 1);
  Module['_SDL_GetGamepadAxisFromString'] = _SDL_GetGamepadAxisFromString = createExportWrapper('SDL_GetGamepadAxisFromString', 1);
  Module['_SDL_GetGamepadStringForAxis'] = _SDL_GetGamepadStringForAxis = createExportWrapper('SDL_GetGamepadStringForAxis', 1);
  Module['_SDL_GetGamepadButtonFromString'] = _SDL_GetGamepadButtonFromString = createExportWrapper('SDL_GetGamepadButtonFromString', 1);
  Module['_SDL_GetGamepadStringForButton'] = _SDL_GetGamepadStringForButton = createExportWrapper('SDL_GetGamepadStringForButton', 1);
  Module['_SDL_AddGamepadMappingsFromIO'] = _SDL_AddGamepadMappingsFromIO = createExportWrapper('SDL_AddGamepadMappingsFromIO', 2);
  Module['_SDL_GetJoysticks'] = _SDL_GetJoysticks = createExportWrapper('SDL_GetJoysticks', 1);
  Module['_SDL_GetJoystickNameForID'] = _SDL_GetJoystickNameForID = createExportWrapper('SDL_GetJoystickNameForID', 1);
  Module['_SDL_GetJoystickGUIDForID'] = _SDL_GetJoystickGUIDForID = createExportWrapper('SDL_GetJoystickGUIDForID', 2);
  Module['_SDL_AddGamepadMapping'] = _SDL_AddGamepadMapping = createExportWrapper('SDL_AddGamepadMapping', 1);
  Module['_SDL_AddGamepadMappingsFromFile'] = _SDL_AddGamepadMappingsFromFile = createExportWrapper('SDL_AddGamepadMappingsFromFile', 1);
  Module['_SDL_ReloadGamepadMappings'] = _SDL_ReloadGamepadMappings = createExportWrapper('SDL_ReloadGamepadMappings', 0);
  Module['_SDL_GetGamepadMappings'] = _SDL_GetGamepadMappings = createExportWrapper('SDL_GetGamepadMappings', 1);
  Module['_SDL_strlcat'] = _SDL_strlcat = createExportWrapper('SDL_strlcat', 3);
  Module['_SDL_GetGamepadMappingForGUID'] = _SDL_GetGamepadMappingForGUID = createExportWrapper('SDL_GetGamepadMappingForGUID', 1);
  Module['_SDL_GetJoystickGUIDInfo'] = _SDL_GetJoystickGUIDInfo = createExportWrapper('SDL_GetJoystickGUIDInfo', 5);
  Module['_SDL_GetGamepadMapping'] = _SDL_GetGamepadMapping = createExportWrapper('SDL_GetGamepadMapping', 1);
  Module['_SDL_SetGamepadMapping'] = _SDL_SetGamepadMapping = createExportWrapper('SDL_SetGamepadMapping', 2);
  Module['_SDL_HasGamepad'] = _SDL_HasGamepad = createExportWrapper('SDL_HasGamepad', 0);
  Module['_SDL_GetGamepads'] = _SDL_GetGamepads = createExportWrapper('SDL_GetGamepads', 1);
  Module['_SDL_GetGamepadNameForID'] = _SDL_GetGamepadNameForID = createExportWrapper('SDL_GetGamepadNameForID', 1);
  Module['_SDL_GetGamepadPathForID'] = _SDL_GetGamepadPathForID = createExportWrapper('SDL_GetGamepadPathForID', 1);
  Module['_SDL_GetJoystickPathForID'] = _SDL_GetJoystickPathForID = createExportWrapper('SDL_GetJoystickPathForID', 1);
  Module['_SDL_GetGamepadPlayerIndexForID'] = _SDL_GetGamepadPlayerIndexForID = createExportWrapper('SDL_GetGamepadPlayerIndexForID', 1);
  Module['_SDL_GetJoystickPlayerIndexForID'] = _SDL_GetJoystickPlayerIndexForID = createExportWrapper('SDL_GetJoystickPlayerIndexForID', 1);
  Module['_SDL_GetGamepadGUIDForID'] = _SDL_GetGamepadGUIDForID = createExportWrapper('SDL_GetGamepadGUIDForID', 2);
  Module['_SDL_GetGamepadVendorForID'] = _SDL_GetGamepadVendorForID = createExportWrapper('SDL_GetGamepadVendorForID', 1);
  Module['_SDL_GetJoystickVendorForID'] = _SDL_GetJoystickVendorForID = createExportWrapper('SDL_GetJoystickVendorForID', 1);
  Module['_SDL_GetGamepadProductForID'] = _SDL_GetGamepadProductForID = createExportWrapper('SDL_GetGamepadProductForID', 1);
  Module['_SDL_GetJoystickProductForID'] = _SDL_GetJoystickProductForID = createExportWrapper('SDL_GetJoystickProductForID', 1);
  Module['_SDL_GetGamepadProductVersionForID'] = _SDL_GetGamepadProductVersionForID = createExportWrapper('SDL_GetGamepadProductVersionForID', 1);
  Module['_SDL_GetJoystickProductVersionForID'] = _SDL_GetJoystickProductVersionForID = createExportWrapper('SDL_GetJoystickProductVersionForID', 1);
  Module['_SDL_GetGamepadTypeForID'] = _SDL_GetGamepadTypeForID = createExportWrapper('SDL_GetGamepadTypeForID', 1);
  Module['_SDL_GetRealGamepadTypeForID'] = _SDL_GetRealGamepadTypeForID = createExportWrapper('SDL_GetRealGamepadTypeForID', 1);
  Module['_SDL_GetGamepadMappingForID'] = _SDL_GetGamepadMappingForID = createExportWrapper('SDL_GetGamepadMappingForID', 1);
  Module['_SDL_OpenGamepad'] = _SDL_OpenGamepad = createExportWrapper('SDL_OpenGamepad', 1);
  Module['_SDL_OpenJoystick'] = _SDL_OpenJoystick = createExportWrapper('SDL_OpenJoystick', 1);
  Module['_SDL_CloseJoystick'] = _SDL_CloseJoystick = createExportWrapper('SDL_CloseJoystick', 1);
  Module['_SDL_UpdateGamepads'] = _SDL_UpdateGamepads = createExportWrapper('SDL_UpdateGamepads', 0);
  Module['_SDL_GamepadHasAxis'] = _SDL_GamepadHasAxis = createExportWrapper('SDL_GamepadHasAxis', 2);
  Module['_SDL_GetJoystickAxis'] = _SDL_GetJoystickAxis = createExportWrapper('SDL_GetJoystickAxis', 2);
  Module['_SDL_GetJoystickButton'] = _SDL_GetJoystickButton = createExportWrapper('SDL_GetJoystickButton', 2);
  Module['_SDL_GetJoystickHat'] = _SDL_GetJoystickHat = createExportWrapper('SDL_GetJoystickHat', 2);
  Module['_SDL_GamepadHasButton'] = _SDL_GamepadHasButton = createExportWrapper('SDL_GamepadHasButton', 2);
  Module['_SDL_GetGamepadButtonLabelForType'] = _SDL_GetGamepadButtonLabelForType = createExportWrapper('SDL_GetGamepadButtonLabelForType', 2);
  Module['_SDL_GetGamepadButtonLabel'] = _SDL_GetGamepadButtonLabel = createExportWrapper('SDL_GetGamepadButtonLabel', 2);
  Module['_SDL_GetNumGamepadTouchpads'] = _SDL_GetNumGamepadTouchpads = createExportWrapper('SDL_GetNumGamepadTouchpads', 1);
  Module['_SDL_GetGamepadJoystick'] = _SDL_GetGamepadJoystick = createExportWrapper('SDL_GetGamepadJoystick', 1);
  Module['_SDL_GetNumGamepadTouchpadFingers'] = _SDL_GetNumGamepadTouchpadFingers = createExportWrapper('SDL_GetNumGamepadTouchpadFingers', 2);
  Module['_SDL_GetGamepadTouchpadFinger'] = _SDL_GetGamepadTouchpadFinger = createExportWrapper('SDL_GetGamepadTouchpadFinger', 7);
  Module['_SDL_GamepadHasSensor'] = _SDL_GamepadHasSensor = createExportWrapper('SDL_GamepadHasSensor', 2);
  Module['_SDL_SetGamepadSensorEnabled'] = _SDL_SetGamepadSensorEnabled = createExportWrapper('SDL_SetGamepadSensorEnabled', 3);
  Module['_SDL_OpenSensor'] = _SDL_OpenSensor = createExportWrapper('SDL_OpenSensor', 1);
  Module['_SDL_CloseSensor'] = _SDL_CloseSensor = createExportWrapper('SDL_CloseSensor', 1);
  Module['_SDL_GamepadSensorEnabled'] = _SDL_GamepadSensorEnabled = createExportWrapper('SDL_GamepadSensorEnabled', 2);
  Module['_SDL_GetGamepadSensorDataRate'] = _SDL_GetGamepadSensorDataRate = createExportWrapper('SDL_GetGamepadSensorDataRate', 2);
  Module['_SDL_GetGamepadSensorData'] = _SDL_GetGamepadSensorData = createExportWrapper('SDL_GetGamepadSensorData', 4);
  Module['_SDL_GetGamepadID'] = _SDL_GetGamepadID = createExportWrapper('SDL_GetGamepadID', 1);
  Module['_SDL_GetGamepadProperties'] = _SDL_GetGamepadProperties = createExportWrapper('SDL_GetGamepadProperties', 1);
  Module['_SDL_GetJoystickProperties'] = _SDL_GetJoystickProperties = createExportWrapper('SDL_GetJoystickProperties', 1);
  Module['_SDL_GetGamepadName'] = _SDL_GetGamepadName = createExportWrapper('SDL_GetGamepadName', 1);
  Module['_SDL_GetJoystickName'] = _SDL_GetJoystickName = createExportWrapper('SDL_GetJoystickName', 1);
  Module['_SDL_GetGamepadPath'] = _SDL_GetGamepadPath = createExportWrapper('SDL_GetGamepadPath', 1);
  Module['_SDL_GetJoystickPath'] = _SDL_GetJoystickPath = createExportWrapper('SDL_GetJoystickPath', 1);
  Module['_SDL_GetGamepadType'] = _SDL_GetGamepadType = createExportWrapper('SDL_GetGamepadType', 1);
  Module['_SDL_GetRealGamepadType'] = _SDL_GetRealGamepadType = createExportWrapper('SDL_GetRealGamepadType', 1);
  Module['_SDL_GetJoystickGUID'] = _SDL_GetJoystickGUID = createExportWrapper('SDL_GetJoystickGUID', 2);
  Module['_SDL_GetGamepadPlayerIndex'] = _SDL_GetGamepadPlayerIndex = createExportWrapper('SDL_GetGamepadPlayerIndex', 1);
  Module['_SDL_GetJoystickPlayerIndex'] = _SDL_GetJoystickPlayerIndex = createExportWrapper('SDL_GetJoystickPlayerIndex', 1);
  Module['_SDL_SetGamepadPlayerIndex'] = _SDL_SetGamepadPlayerIndex = createExportWrapper('SDL_SetGamepadPlayerIndex', 2);
  Module['_SDL_SetJoystickPlayerIndex'] = _SDL_SetJoystickPlayerIndex = createExportWrapper('SDL_SetJoystickPlayerIndex', 2);
  Module['_SDL_GetGamepadVendor'] = _SDL_GetGamepadVendor = createExportWrapper('SDL_GetGamepadVendor', 1);
  Module['_SDL_GetGamepadProduct'] = _SDL_GetGamepadProduct = createExportWrapper('SDL_GetGamepadProduct', 1);
  Module['_SDL_GetGamepadProductVersion'] = _SDL_GetGamepadProductVersion = createExportWrapper('SDL_GetGamepadProductVersion', 1);
  Module['_SDL_GetJoystickProductVersion'] = _SDL_GetJoystickProductVersion = createExportWrapper('SDL_GetJoystickProductVersion', 1);
  Module['_SDL_GetGamepadFirmwareVersion'] = _SDL_GetGamepadFirmwareVersion = createExportWrapper('SDL_GetGamepadFirmwareVersion', 1);
  Module['_SDL_GetJoystickFirmwareVersion'] = _SDL_GetJoystickFirmwareVersion = createExportWrapper('SDL_GetJoystickFirmwareVersion', 1);
  Module['_SDL_GetGamepadSerial'] = _SDL_GetGamepadSerial = createExportWrapper('SDL_GetGamepadSerial', 1);
  Module['_SDL_GetJoystickSerial'] = _SDL_GetJoystickSerial = createExportWrapper('SDL_GetJoystickSerial', 1);
  Module['_SDL_GetGamepadSteamHandle'] = _SDL_GetGamepadSteamHandle = createExportWrapper('SDL_GetGamepadSteamHandle', 1);
  Module['_SDL_GetGamepadConnectionState'] = _SDL_GetGamepadConnectionState = createExportWrapper('SDL_GetGamepadConnectionState', 1);
  Module['_SDL_GetJoystickConnectionState'] = _SDL_GetJoystickConnectionState = createExportWrapper('SDL_GetJoystickConnectionState', 1);
  Module['_SDL_GetGamepadPowerInfo'] = _SDL_GetGamepadPowerInfo = createExportWrapper('SDL_GetGamepadPowerInfo', 2);
  Module['_SDL_GetJoystickPowerInfo'] = _SDL_GetJoystickPowerInfo = createExportWrapper('SDL_GetJoystickPowerInfo', 2);
  Module['_SDL_GamepadConnected'] = _SDL_GamepadConnected = createExportWrapper('SDL_GamepadConnected', 1);
  Module['_SDL_JoystickConnected'] = _SDL_JoystickConnected = createExportWrapper('SDL_JoystickConnected', 1);
  Module['_SDL_GetGamepadFromID'] = _SDL_GetGamepadFromID = createExportWrapper('SDL_GetGamepadFromID', 1);
  Module['_SDL_GetGamepadFromPlayerIndex'] = _SDL_GetGamepadFromPlayerIndex = createExportWrapper('SDL_GetGamepadFromPlayerIndex', 1);
  Module['_SDL_GetJoystickFromPlayerIndex'] = _SDL_GetJoystickFromPlayerIndex = createExportWrapper('SDL_GetJoystickFromPlayerIndex', 1);
  Module['_SDL_GetGamepadBindings'] = _SDL_GetGamepadBindings = createExportWrapper('SDL_GetGamepadBindings', 2);
  Module['_SDL_RumbleGamepad'] = _SDL_RumbleGamepad = createExportWrapper('SDL_RumbleGamepad', 4);
  Module['_SDL_RumbleJoystick'] = _SDL_RumbleJoystick = createExportWrapper('SDL_RumbleJoystick', 4);
  Module['_SDL_RumbleGamepadTriggers'] = _SDL_RumbleGamepadTriggers = createExportWrapper('SDL_RumbleGamepadTriggers', 4);
  Module['_SDL_RumbleJoystickTriggers'] = _SDL_RumbleJoystickTriggers = createExportWrapper('SDL_RumbleJoystickTriggers', 4);
  Module['_SDL_SetGamepadLED'] = _SDL_SetGamepadLED = createExportWrapper('SDL_SetGamepadLED', 4);
  Module['_SDL_SetJoystickLED'] = _SDL_SetJoystickLED = createExportWrapper('SDL_SetJoystickLED', 4);
  Module['_SDL_SendGamepadEffect'] = _SDL_SendGamepadEffect = createExportWrapper('SDL_SendGamepadEffect', 3);
  Module['_SDL_SendJoystickEffect'] = _SDL_SendJoystickEffect = createExportWrapper('SDL_SendJoystickEffect', 3);
  Module['_SDL_CloseGamepad'] = _SDL_CloseGamepad = createExportWrapper('SDL_CloseGamepad', 1);
  Module['_SDL_SetGamepadEventsEnabled'] = _SDL_SetGamepadEventsEnabled = createExportWrapper('SDL_SetGamepadEventsEnabled', 1);
  Module['_SDL_GamepadEventsEnabled'] = _SDL_GamepadEventsEnabled = createExportWrapper('SDL_GamepadEventsEnabled', 0);
  Module['_SDL_GetGamepadAppleSFSymbolsNameForButton'] = _SDL_GetGamepadAppleSFSymbolsNameForButton = createExportWrapper('SDL_GetGamepadAppleSFSymbolsNameForButton', 2);
  Module['_SDL_GetGamepadAppleSFSymbolsNameForAxis'] = _SDL_GetGamepadAppleSFSymbolsNameForAxis = createExportWrapper('SDL_GetGamepadAppleSFSymbolsNameForAxis', 2);
  Module['_SDL_HasJoystick'] = _SDL_HasJoystick = createExportWrapper('SDL_HasJoystick', 0);
  Module['_SDL_GetSensors'] = _SDL_GetSensors = createExportWrapper('SDL_GetSensors', 1);
  Module['_SDL_GetSensorTypeForID'] = _SDL_GetSensorTypeForID = createExportWrapper('SDL_GetSensorTypeForID', 1);
  Module['_SDL_GetSensorNameForID'] = _SDL_GetSensorNameForID = createExportWrapper('SDL_GetSensorNameForID', 1);
  Module['_SDL_GetPrimaryDisplay'] = _SDL_GetPrimaryDisplay = createExportWrapper('SDL_GetPrimaryDisplay', 0);
  Module['_SDL_GetNaturalDisplayOrientation'] = _SDL_GetNaturalDisplayOrientation = createExportWrapper('SDL_GetNaturalDisplayOrientation', 1);
  Module['_SDL_AttachVirtualJoystick'] = _SDL_AttachVirtualJoystick = createExportWrapper('SDL_AttachVirtualJoystick', 1);
  Module['_SDL_DetachVirtualJoystick'] = _SDL_DetachVirtualJoystick = createExportWrapper('SDL_DetachVirtualJoystick', 1);
  Module['_SDL_IsJoystickVirtual'] = _SDL_IsJoystickVirtual = createExportWrapper('SDL_IsJoystickVirtual', 1);
  Module['_SDL_SetJoystickVirtualAxis'] = _SDL_SetJoystickVirtualAxis = createExportWrapper('SDL_SetJoystickVirtualAxis', 3);
  Module['_SDL_SetJoystickVirtualBall'] = _SDL_SetJoystickVirtualBall = createExportWrapper('SDL_SetJoystickVirtualBall', 4);
  Module['_SDL_SetJoystickVirtualButton'] = _SDL_SetJoystickVirtualButton = createExportWrapper('SDL_SetJoystickVirtualButton', 3);
  Module['_SDL_SetJoystickVirtualHat'] = _SDL_SetJoystickVirtualHat = createExportWrapper('SDL_SetJoystickVirtualHat', 3);
  Module['_SDL_SetJoystickVirtualTouchpad'] = _SDL_SetJoystickVirtualTouchpad = createExportWrapper('SDL_SetJoystickVirtualTouchpad', 7);
  Module['_SDL_SendJoystickVirtualSensorData'] = _SDL_SendJoystickVirtualSensorData = createExportWrapper('SDL_SendJoystickVirtualSensorData', 5);
  Module['_SDL_GetNumJoystickHats'] = _SDL_GetNumJoystickHats = createExportWrapper('SDL_GetNumJoystickHats', 1);
  Module['_SDL_GetNumJoystickBalls'] = _SDL_GetNumJoystickBalls = createExportWrapper('SDL_GetNumJoystickBalls', 1);
  Module['_SDL_GetNumJoystickButtons'] = _SDL_GetNumJoystickButtons = createExportWrapper('SDL_GetNumJoystickButtons', 1);
  Module['_SDL_GetJoystickAxisInitialState'] = _SDL_GetJoystickAxisInitialState = createExportWrapper('SDL_GetJoystickAxisInitialState', 3);
  Module['_SDL_GetJoystickBall'] = _SDL_GetJoystickBall = createExportWrapper('SDL_GetJoystickBall', 4);
  Module['_SDL_GetJoystickFromID'] = _SDL_GetJoystickFromID = createExportWrapper('SDL_GetJoystickFromID', 1);
  Module['_SDL_SetJoystickEventsEnabled'] = _SDL_SetJoystickEventsEnabled = createExportWrapper('SDL_SetJoystickEventsEnabled', 1);
  Module['_SDL_JoystickEventsEnabled'] = _SDL_JoystickEventsEnabled = createExportWrapper('SDL_JoystickEventsEnabled', 0);
  Module['_SDL_crc16'] = _SDL_crc16 = createExportWrapper('SDL_crc16', 3);
  Module['_SDL_GetJoystickTypeForID'] = _SDL_GetJoystickTypeForID = createExportWrapper('SDL_GetJoystickTypeForID', 1);
  Module['_SDL_GetJoystickType'] = _SDL_GetJoystickType = createExportWrapper('SDL_GetJoystickType', 1);
  Module['_SDL_strtoul'] = _SDL_strtoul = createExportWrapper('SDL_strtoul', 3);
  Module['_SDL_strtoull'] = _SDL_strtoull = createExportWrapper('SDL_strtoull', 3);
  Module['_SDL_GetPreferredLocales'] = _SDL_GetPreferredLocales = createExportWrapper('SDL_GetPreferredLocales', 1);
  Module['_SDL_OpenURL'] = _SDL_OpenURL = createExportWrapper('SDL_OpenURL', 1);
  Module['_SDL_GetPowerInfo'] = _SDL_GetPowerInfo = createExportWrapper('SDL_GetPowerInfo', 2);
  Module['_SDL_DestroyRenderer'] = _SDL_DestroyRenderer = createExportWrapper('SDL_DestroyRenderer', 1);
  Module['_SDL_GetRendererProperties'] = _SDL_GetRendererProperties = createExportWrapper('SDL_GetRendererProperties', 1);
  Module['_SDL_FlushRenderer'] = _SDL_FlushRenderer = createExportWrapper('SDL_FlushRenderer', 1);
  Module['_SDL_GetNumRenderDrivers'] = _SDL_GetNumRenderDrivers = createExportWrapper('SDL_GetNumRenderDrivers', 0);
  Module['_SDL_GetRenderDriver'] = _SDL_GetRenderDriver = createExportWrapper('SDL_GetRenderDriver', 1);
  Module['_SDL_CreateWindowAndRenderer'] = _SDL_CreateWindowAndRenderer = createExportWrapper('SDL_CreateWindowAndRenderer', 6);
  Module['_SDL_CreateWindow'] = _SDL_CreateWindow = createExportWrapper('SDL_CreateWindow', 4);
  Module['_SDL_CreateRendererWithProperties'] = _SDL_CreateRendererWithProperties = createExportWrapper('SDL_CreateRendererWithProperties', 1);
  Module['_SDL_DestroyWindow'] = _SDL_DestroyWindow = createExportWrapper('SDL_DestroyWindow', 1);
  Module['_SDL_ShowWindow'] = _SDL_ShowWindow = createExportWrapper('SDL_ShowWindow', 1);
  Module['_SDL_CreateRenderer'] = _SDL_CreateRenderer = createExportWrapper('SDL_CreateRenderer', 2);
  Module['_SDL_WindowHasSurface'] = _SDL_WindowHasSurface = createExportWrapper('SDL_WindowHasSurface', 1);
  Module['_SDL_GetWindowProperties'] = _SDL_GetWindowProperties = createExportWrapper('SDL_GetWindowProperties', 1);
  Module['_SDL_GetWindowFlags'] = _SDL_GetWindowFlags = createExportWrapper('SDL_GetWindowFlags', 1);
  Module['_SDL_SetRenderViewport'] = _SDL_SetRenderViewport = createExportWrapper('SDL_SetRenderViewport', 2);
  Module['_SDL_SetRenderVSync'] = _SDL_SetRenderVSync = createExportWrapper('SDL_SetRenderVSync', 2);
  Module['_SDL_GetDisplayForWindow'] = _SDL_GetDisplayForWindow = createExportWrapper('SDL_GetDisplayForWindow', 1);
  Module['_SDL_GetDesktopDisplayMode'] = _SDL_GetDesktopDisplayMode = createExportWrapper('SDL_GetDesktopDisplayMode', 1);
  Module['_SDL_GetRenderer'] = _SDL_GetRenderer = createExportWrapper('SDL_GetRenderer', 1);
  Module['_SDL_GetWindowSize'] = _SDL_GetWindowSize = createExportWrapper('SDL_GetWindowSize', 3);
  Module['_SDL_GetWindowSizeInPixels'] = _SDL_GetWindowSizeInPixels = createExportWrapper('SDL_GetWindowSizeInPixels', 3);
  Module['_SDL_CreateSoftwareRenderer'] = _SDL_CreateSoftwareRenderer = createExportWrapper('SDL_CreateSoftwareRenderer', 1);
  Module['_SDL_GetRenderWindow'] = _SDL_GetRenderWindow = createExportWrapper('SDL_GetRenderWindow', 1);
  Module['_SDL_GetRendererName'] = _SDL_GetRendererName = createExportWrapper('SDL_GetRendererName', 1);
  Module['_SDL_GetRenderOutputSize'] = _SDL_GetRenderOutputSize = createExportWrapper('SDL_GetRenderOutputSize', 3);
  Module['_SDL_GetCurrentRenderOutputSize'] = _SDL_GetCurrentRenderOutputSize = createExportWrapper('SDL_GetCurrentRenderOutputSize', 3);
  Module['_SDL_CreateTextureWithProperties'] = _SDL_CreateTextureWithProperties = createExportWrapper('SDL_CreateTextureWithProperties', 2);
  Module['_SDL_GetTextureProperties'] = _SDL_GetTextureProperties = createExportWrapper('SDL_GetTextureProperties', 1);
  Module['_SDL_GetSurfaceColorspace'] = _SDL_GetSurfaceColorspace = createExportWrapper('SDL_GetSurfaceColorspace', 1);
  Module['_SDL_LockSurface'] = _SDL_LockSurface = createExportWrapper('SDL_LockSurface', 1);
  Module['_SDL_UpdateTexture'] = _SDL_UpdateTexture = createExportWrapper('SDL_UpdateTexture', 4);
  Module['_SDL_UnlockSurface'] = _SDL_UnlockSurface = createExportWrapper('SDL_UnlockSurface', 1);
  Module['_SDL_ConvertSurfaceAndColorspace'] = _SDL_ConvertSurfaceAndColorspace = createExportWrapper('SDL_ConvertSurfaceAndColorspace', 5);
  Module['_SDL_GetSurfaceColorMod'] = _SDL_GetSurfaceColorMod = createExportWrapper('SDL_GetSurfaceColorMod', 4);
  Module['_SDL_GetSurfaceAlphaMod'] = _SDL_GetSurfaceAlphaMod = createExportWrapper('SDL_GetSurfaceAlphaMod', 2);
  Module['_SDL_GetSurfaceBlendMode'] = _SDL_GetSurfaceBlendMode = createExportWrapper('SDL_GetSurfaceBlendMode', 2);
  Module['_SDL_SetTextureBlendMode'] = _SDL_SetTextureBlendMode = createExportWrapper('SDL_SetTextureBlendMode', 2);
  Module['_SDL_GetRendererFromTexture'] = _SDL_GetRendererFromTexture = createExportWrapper('SDL_GetRendererFromTexture', 1);
  Module['_SDL_GetTextureSize'] = _SDL_GetTextureSize = createExportWrapper('SDL_GetTextureSize', 3);
  Module['_SDL_SetTextureColorMod'] = _SDL_SetTextureColorMod = createExportWrapper('SDL_SetTextureColorMod', 4);
  Module['_SDL_SetTextureColorModFloat'] = _SDL_SetTextureColorModFloat = createExportWrapper('SDL_SetTextureColorModFloat', 4);
  Module['_SDL_GetTextureColorMod'] = _SDL_GetTextureColorMod = createExportWrapper('SDL_GetTextureColorMod', 4);
  Module['_SDL_GetTextureColorModFloat'] = _SDL_GetTextureColorModFloat = createExportWrapper('SDL_GetTextureColorModFloat', 4);
  Module['_SDL_SetTextureAlphaMod'] = _SDL_SetTextureAlphaMod = createExportWrapper('SDL_SetTextureAlphaMod', 2);
  Module['_SDL_SetTextureAlphaModFloat'] = _SDL_SetTextureAlphaModFloat = createExportWrapper('SDL_SetTextureAlphaModFloat', 2);
  Module['_SDL_GetTextureAlphaMod'] = _SDL_GetTextureAlphaMod = createExportWrapper('SDL_GetTextureAlphaMod', 2);
  Module['_SDL_GetTextureAlphaModFloat'] = _SDL_GetTextureAlphaModFloat = createExportWrapper('SDL_GetTextureAlphaModFloat', 2);
  Module['_SDL_GetTextureBlendMode'] = _SDL_GetTextureBlendMode = createExportWrapper('SDL_GetTextureBlendMode', 2);
  Module['_SDL_GetTextureScaleMode'] = _SDL_GetTextureScaleMode = createExportWrapper('SDL_GetTextureScaleMode', 2);
  Module['_SDL_ConvertPixelsAndColorspace'] = _SDL_ConvertPixelsAndColorspace = createExportWrapper('SDL_ConvertPixelsAndColorspace', 12);
  Module['_SDL_UpdateYUVTexture'] = _SDL_UpdateYUVTexture = createExportWrapper('SDL_UpdateYUVTexture', 8);
  Module['_SDL_UpdateNVTexture'] = _SDL_UpdateNVTexture = createExportWrapper('SDL_UpdateNVTexture', 6);
  Module['_SDL_LockTextureToSurface'] = _SDL_LockTextureToSurface = createExportWrapper('SDL_LockTextureToSurface', 3);
  Module['_SDL_SetRenderTarget'] = _SDL_SetRenderTarget = createExportWrapper('SDL_SetRenderTarget', 2);
  Module['_SDL_GetRenderTarget'] = _SDL_GetRenderTarget = createExportWrapper('SDL_GetRenderTarget', 1);
  Module['_SDL_SetRenderLogicalPresentation'] = _SDL_SetRenderLogicalPresentation = createExportWrapper('SDL_SetRenderLogicalPresentation', 4);
  Module['_SDL_GetRenderLogicalPresentation'] = _SDL_GetRenderLogicalPresentation = createExportWrapper('SDL_GetRenderLogicalPresentation', 4);
  Module['_SDL_GetRenderLogicalPresentationRect'] = _SDL_GetRenderLogicalPresentationRect = createExportWrapper('SDL_GetRenderLogicalPresentationRect', 2);
  Module['_SDL_RenderCoordinatesFromWindow'] = _SDL_RenderCoordinatesFromWindow = createExportWrapper('SDL_RenderCoordinatesFromWindow', 5);
  Module['_SDL_RenderCoordinatesToWindow'] = _SDL_RenderCoordinatesToWindow = createExportWrapper('SDL_RenderCoordinatesToWindow', 5);
  Module['_SDL_ConvertEventToRenderCoordinates'] = _SDL_ConvertEventToRenderCoordinates = createExportWrapper('SDL_ConvertEventToRenderCoordinates', 2);
  Module['_SDL_GetRenderViewport'] = _SDL_GetRenderViewport = createExportWrapper('SDL_GetRenderViewport', 2);
  Module['_SDL_RenderViewportSet'] = _SDL_RenderViewportSet = createExportWrapper('SDL_RenderViewportSet', 1);
  Module['_SDL_GetRenderSafeArea'] = _SDL_GetRenderSafeArea = createExportWrapper('SDL_GetRenderSafeArea', 2);
  Module['_SDL_GetWindowSafeArea'] = _SDL_GetWindowSafeArea = createExportWrapper('SDL_GetWindowSafeArea', 2);
  Module['_SDL_SetRenderClipRect'] = _SDL_SetRenderClipRect = createExportWrapper('SDL_SetRenderClipRect', 2);
  Module['_SDL_GetRenderClipRect'] = _SDL_GetRenderClipRect = createExportWrapper('SDL_GetRenderClipRect', 2);
  Module['_SDL_RenderClipEnabled'] = _SDL_RenderClipEnabled = createExportWrapper('SDL_RenderClipEnabled', 1);
  Module['_SDL_SetRenderScale'] = _SDL_SetRenderScale = createExportWrapper('SDL_SetRenderScale', 3);
  Module['_SDL_GetRenderScale'] = _SDL_GetRenderScale = createExportWrapper('SDL_GetRenderScale', 3);
  Module['_SDL_SetRenderDrawColor'] = _SDL_SetRenderDrawColor = createExportWrapper('SDL_SetRenderDrawColor', 5);
  Module['_SDL_SetRenderDrawColorFloat'] = _SDL_SetRenderDrawColorFloat = createExportWrapper('SDL_SetRenderDrawColorFloat', 5);
  Module['_SDL_GetRenderDrawColor'] = _SDL_GetRenderDrawColor = createExportWrapper('SDL_GetRenderDrawColor', 5);
  Module['_SDL_GetRenderDrawColorFloat'] = _SDL_GetRenderDrawColorFloat = createExportWrapper('SDL_GetRenderDrawColorFloat', 5);
  Module['_SDL_SetRenderColorScale'] = _SDL_SetRenderColorScale = createExportWrapper('SDL_SetRenderColorScale', 2);
  Module['_SDL_GetRenderColorScale'] = _SDL_GetRenderColorScale = createExportWrapper('SDL_GetRenderColorScale', 2);
  Module['_SDL_SetRenderDrawBlendMode'] = _SDL_SetRenderDrawBlendMode = createExportWrapper('SDL_SetRenderDrawBlendMode', 2);
  Module['_SDL_GetRenderDrawBlendMode'] = _SDL_GetRenderDrawBlendMode = createExportWrapper('SDL_GetRenderDrawBlendMode', 2);
  Module['_SDL_RenderClear'] = _SDL_RenderClear = createExportWrapper('SDL_RenderClear', 1);
  Module['_SDL_RenderPoint'] = _SDL_RenderPoint = createExportWrapper('SDL_RenderPoint', 3);
  Module['_SDL_RenderPoints'] = _SDL_RenderPoints = createExportWrapper('SDL_RenderPoints', 3);
  Module['_SDL_RenderLine'] = _SDL_RenderLine = createExportWrapper('SDL_RenderLine', 5);
  Module['_SDL_RenderLines'] = _SDL_RenderLines = createExportWrapper('SDL_RenderLines', 3);
  Module['_SDL_RenderRect'] = _SDL_RenderRect = createExportWrapper('SDL_RenderRect', 2);
  Module['_SDL_RenderRects'] = _SDL_RenderRects = createExportWrapper('SDL_RenderRects', 3);
  Module['_SDL_RenderFillRect'] = _SDL_RenderFillRect = createExportWrapper('SDL_RenderFillRect', 2);
  Module['_SDL_RenderFillRects'] = _SDL_RenderFillRects = createExportWrapper('SDL_RenderFillRects', 3);
  Module['_SDL_RenderTexture'] = _SDL_RenderTexture = createExportWrapper('SDL_RenderTexture', 4);
  Module['_SDL_GetRectIntersectionFloat'] = _SDL_GetRectIntersectionFloat = createExportWrapper('SDL_GetRectIntersectionFloat', 3);
  Module['_SDL_RenderTextureAffine'] = _SDL_RenderTextureAffine = createExportWrapper('SDL_RenderTextureAffine', 6);
  Module['_SDL_RenderTextureRotated'] = _SDL_RenderTextureRotated = createExportWrapper('SDL_RenderTextureRotated', 7);
  Module['_SDL_RenderTextureTiled'] = _SDL_RenderTextureTiled = createExportWrapper('SDL_RenderTextureTiled', 5);
  Module['_SDL_RenderTexture9Grid'] = _SDL_RenderTexture9Grid = createExportWrapper('SDL_RenderTexture9Grid', 9);
  Module['_SDL_RenderGeometry'] = _SDL_RenderGeometry = createExportWrapper('SDL_RenderGeometry', 6);
  Module['_SDL_RenderReadPixels'] = _SDL_RenderReadPixels = createExportWrapper('SDL_RenderReadPixels', 2);
  Module['_SDL_GetPixelFormatDetails'] = _SDL_GetPixelFormatDetails = createExportWrapper('SDL_GetPixelFormatDetails', 1);
  Module['_SDL_RenderPresent'] = _SDL_RenderPresent = createExportWrapper('SDL_RenderPresent', 1);
  Module['_SDL_DelayPrecise'] = _SDL_DelayPrecise = createExportWrapper('SDL_DelayPrecise', 1);
  Module['_SDL_GetRenderMetalLayer'] = _SDL_GetRenderMetalLayer = createExportWrapper('SDL_GetRenderMetalLayer', 1);
  Module['_SDL_GetRenderMetalCommandEncoder'] = _SDL_GetRenderMetalCommandEncoder = createExportWrapper('SDL_GetRenderMetalCommandEncoder', 1);
  Module['_SDL_AddVulkanRenderSemaphores'] = _SDL_AddVulkanRenderSemaphores = createExportWrapper('SDL_AddVulkanRenderSemaphores', 4);
  Module['_SDL_ComposeCustomBlendMode'] = _SDL_ComposeCustomBlendMode = createExportWrapper('SDL_ComposeCustomBlendMode', 6);
  Module['_SDL_GetRenderVSync'] = _SDL_GetRenderVSync = createExportWrapper('SDL_GetRenderVSync', 2);
  Module['_SDL_RenderDebugText'] = _SDL_RenderDebugText = createExportWrapper('SDL_RenderDebugText', 4);
  Module['_SDL_RenderDebugTextFormat'] = _SDL_RenderDebugTextFormat = createExportWrapper('SDL_RenderDebugTextFormat', 5);
  Module['_SDL_GetRectAndLineIntersection'] = _SDL_GetRectAndLineIntersection = createExportWrapper('SDL_GetRectAndLineIntersection', 5);
  Module['_SDL_GetPixelFormatName'] = _SDL_GetPixelFormatName = createExportWrapper('SDL_GetPixelFormatName', 1);
  Module['_SDL_GL_GetAttribute'] = _SDL_GL_GetAttribute = createExportWrapper('SDL_GL_GetAttribute', 2);
  Module['_SDL_SyncWindow'] = _SDL_SyncWindow = createExportWrapper('SDL_SyncWindow', 1);
  Module['_SDL_GL_SetAttribute'] = _SDL_GL_SetAttribute = createExportWrapper('SDL_GL_SetAttribute', 2);
  Module['_SDL_GL_CreateContext'] = _SDL_GL_CreateContext = createExportWrapper('SDL_GL_CreateContext', 1);
  Module['_SDL_GL_MakeCurrent'] = _SDL_GL_MakeCurrent = createExportWrapper('SDL_GL_MakeCurrent', 2);
  Module['_SDL_GL_ExtensionSupported'] = _SDL_GL_ExtensionSupported = createExportWrapper('SDL_GL_ExtensionSupported', 1);
  Module['_SDL_GL_GetProcAddress'] = _SDL_GL_GetProcAddress = createExportWrapper('SDL_GL_GetProcAddress', 1);
  Module['_SDL_GL_GetCurrentContext'] = _SDL_GL_GetCurrentContext = createExportWrapper('SDL_GL_GetCurrentContext', 0);
  Module['_SDL_FlipSurface'] = _SDL_FlipSurface = createExportWrapper('SDL_FlipSurface', 2);
  Module['_SDL_GL_SwapWindow'] = _SDL_GL_SwapWindow = createExportWrapper('SDL_GL_SwapWindow', 1);
  Module['_SDL_GL_DestroyContext'] = _SDL_GL_DestroyContext = createExportWrapper('SDL_GL_DestroyContext', 1);
  Module['_SDL_GL_SetSwapInterval'] = _SDL_GL_SetSwapInterval = createExportWrapper('SDL_GL_SetSwapInterval', 1);
  Module['_SDL_GL_GetSwapInterval'] = _SDL_GL_GetSwapInterval = createExportWrapper('SDL_GL_GetSwapInterval', 1);
  Module['_SDL_GetRGBA'] = _SDL_GetRGBA = createExportWrapper('SDL_GetRGBA', 7);
  Module['_SDL_SetSurfaceRLE'] = _SDL_SetSurfaceRLE = createExportWrapper('SDL_SetSurfaceRLE', 2);
  Module['_SDL_GetWindowSurface'] = _SDL_GetWindowSurface = createExportWrapper('SDL_GetWindowSurface', 1);
  Module['_SDL_SetSurfaceClipRect'] = _SDL_SetSurfaceClipRect = createExportWrapper('SDL_SetSurfaceClipRect', 2);
  Module['_SDL_FillSurfaceRects'] = _SDL_FillSurfaceRects = createExportWrapper('SDL_FillSurfaceRects', 4);
  Module['_SDL_BlitSurfaceScaled'] = _SDL_BlitSurfaceScaled = createExportWrapper('SDL_BlitSurfaceScaled', 5);
  Module['_SDL_UpdateWindowSurface'] = _SDL_UpdateWindowSurface = createExportWrapper('SDL_UpdateWindowSurface', 1);
  Module['_SDL_DestroyWindowSurface'] = _SDL_DestroyWindowSurface = createExportWrapper('SDL_DestroyWindowSurface', 1);
  Module['_SDL_ceil'] = _SDL_ceil = createExportWrapper('SDL_ceil', 1);
  Module['_SDL_GetSurfaceColorKey'] = _SDL_GetSurfaceColorKey = createExportWrapper('SDL_GetSurfaceColorKey', 2);
  Module['_SDL_GetSurfaceClipRect'] = _SDL_GetSurfaceClipRect = createExportWrapper('SDL_GetSurfaceClipRect', 2);
  Module['_SDL_MapRGBA'] = _SDL_MapRGBA = createExportWrapper('SDL_MapRGBA', 6);
  Module['_SDL_GetSensorNonPortableTypeForID'] = _SDL_GetSensorNonPortableTypeForID = createExportWrapper('SDL_GetSensorNonPortableTypeForID', 1);
  Module['_SDL_GetSensorFromID'] = _SDL_GetSensorFromID = createExportWrapper('SDL_GetSensorFromID', 1);
  Module['_SDL_GetSensorProperties'] = _SDL_GetSensorProperties = createExportWrapper('SDL_GetSensorProperties', 1);
  Module['_SDL_GetSensorName'] = _SDL_GetSensorName = createExportWrapper('SDL_GetSensorName', 1);
  Module['_SDL_GetSensorType'] = _SDL_GetSensorType = createExportWrapper('SDL_GetSensorType', 1);
  Module['_SDL_GetSensorNonPortableType'] = _SDL_GetSensorNonPortableType = createExportWrapper('SDL_GetSensorNonPortableType', 1);
  Module['_SDL_GetSensorID'] = _SDL_GetSensorID = createExportWrapper('SDL_GetSensorID', 1);
  Module['_SDL_GetSensorData'] = _SDL_GetSensorData = createExportWrapper('SDL_GetSensorData', 3);
  Module['_SDL_crc32'] = _SDL_crc32 = createExportWrapper('SDL_crc32', 3);
  Module['_SDL_GetEnvironment'] = _SDL_GetEnvironment = createExportWrapper('SDL_GetEnvironment', 0);
  Module['_SDL_CreateEnvironment'] = _SDL_CreateEnvironment = createExportWrapper('SDL_CreateEnvironment', 1);
  Module['_SDL_DestroyEnvironment'] = _SDL_DestroyEnvironment = createExportWrapper('SDL_DestroyEnvironment', 1);
  Module['_SDL_GetEnvironmentVariable'] = _SDL_GetEnvironmentVariable = createExportWrapper('SDL_GetEnvironmentVariable', 2);
  Module['_SDL_GetEnvironmentVariables'] = _SDL_GetEnvironmentVariables = createExportWrapper('SDL_GetEnvironmentVariables', 1);
  Module['_SDL_SetEnvironmentVariable'] = _SDL_SetEnvironmentVariable = createExportWrapper('SDL_SetEnvironmentVariable', 4);
  Module['_SDL_UnsetEnvironmentVariable'] = _SDL_UnsetEnvironmentVariable = createExportWrapper('SDL_UnsetEnvironmentVariable', 2);
  Module['_SDL_setenv_unsafe'] = _SDL_setenv_unsafe = createExportWrapper('SDL_setenv_unsafe', 3);
  Module['_SDL_unsetenv_unsafe'] = _SDL_unsetenv_unsafe = createExportWrapper('SDL_unsetenv_unsafe', 1);
  Module['_SDL_getenv_unsafe'] = _SDL_getenv_unsafe = createExportWrapper('SDL_getenv_unsafe', 1);
  Module['_SDL_iconv_open'] = _SDL_iconv_open = createExportWrapper('SDL_iconv_open', 2);
  Module['_SDL_iconv_close'] = _SDL_iconv_close = createExportWrapper('SDL_iconv_close', 1);
  Module['_SDL_iconv'] = _SDL_iconv = createExportWrapper('SDL_iconv', 5);
  Module['_SDL_GetOriginalMemoryFunctions'] = _SDL_GetOriginalMemoryFunctions = createExportWrapper('SDL_GetOriginalMemoryFunctions', 4);
  Module['_SDL_GetMemoryFunctions'] = _SDL_GetMemoryFunctions = createExportWrapper('SDL_GetMemoryFunctions', 4);
  Module['_SDL_SetMemoryFunctions'] = _SDL_SetMemoryFunctions = createExportWrapper('SDL_SetMemoryFunctions', 4);
  Module['_SDL_GetNumAllocations'] = _SDL_GetNumAllocations = createExportWrapper('SDL_GetNumAllocations', 0);
  Module['_SDL_memmove'] = _SDL_memmove = createExportWrapper('SDL_memmove', 3);
  Module['_SDL_bsearch_r'] = _SDL_bsearch_r = createExportWrapper('SDL_bsearch_r', 6);
  Module['_SDL_bsearch'] = _SDL_bsearch = createExportWrapper('SDL_bsearch', 5);
  Module['_SDL_srand'] = _SDL_srand = createExportWrapper('SDL_srand', 1);
  Module['_SDL_GetPerformanceCounter'] = _SDL_GetPerformanceCounter = createExportWrapper('SDL_GetPerformanceCounter', 0);
  Module['_SDL_rand'] = _SDL_rand = createExportWrapper('SDL_rand', 1);
  Module['_SDL_rand_r'] = _SDL_rand_r = createExportWrapper('SDL_rand_r', 2);
  Module['_SDL_randf'] = _SDL_randf = createExportWrapper('SDL_randf', 0);
  Module['_SDL_randf_r'] = _SDL_randf_r = createExportWrapper('SDL_randf_r', 1);
  Module['_SDL_rand_bits'] = _SDL_rand_bits = createExportWrapper('SDL_rand_bits', 0);
  Module['_SDL_rand_bits_r'] = _SDL_rand_bits_r = createExportWrapper('SDL_rand_bits_r', 1);
  Module['_SDL_atan'] = _SDL_atan = createExportWrapper('SDL_atan', 1);
  Module['_SDL_atanf'] = _SDL_atanf = createExportWrapper('SDL_atanf', 1);
  Module['_SDL_atan2'] = _SDL_atan2 = createExportWrapper('SDL_atan2', 2);
  Module['_SDL_acos'] = _SDL_acos = createExportWrapper('SDL_acos', 1);
  Module['_SDL_asin'] = _SDL_asin = createExportWrapper('SDL_asin', 1);
  Module['_SDL_asinf'] = _SDL_asinf = createExportWrapper('SDL_asinf', 1);
  Module['_SDL_copysign'] = _SDL_copysign = createExportWrapper('SDL_copysign', 2);
  Module['_SDL_copysignf'] = _SDL_copysignf = createExportWrapper('SDL_copysignf', 2);
  Module['_SDL_expf'] = _SDL_expf = createExportWrapper('SDL_expf', 1);
  Module['_SDL_trunc'] = _SDL_trunc = createExportWrapper('SDL_trunc', 1);
  Module['_SDL_fmod'] = _SDL_fmod = createExportWrapper('SDL_fmod', 2);
  Module['_SDL_isinf'] = _SDL_isinf = createExportWrapper('SDL_isinf', 1);
  Module['_SDL_isinff'] = _SDL_isinff = createExportWrapper('SDL_isinff', 1);
  Module['_SDL_isnan'] = _SDL_isnan = createExportWrapper('SDL_isnan', 1);
  Module['_SDL_isnanf'] = _SDL_isnanf = createExportWrapper('SDL_isnanf', 1);
  Module['_SDL_logf'] = _SDL_logf = createExportWrapper('SDL_logf', 1);
  Module['_SDL_log10'] = _SDL_log10 = createExportWrapper('SDL_log10', 1);
  Module['_SDL_log10f'] = _SDL_log10f = createExportWrapper('SDL_log10f', 1);
  Module['_SDL_modf'] = _SDL_modf = createExportWrapper('SDL_modf', 2);
  Module['_SDL_powf'] = _SDL_powf = createExportWrapper('SDL_powf', 2);
  Module['_SDL_lround'] = _SDL_lround = createExportWrapper('SDL_lround', 1);
  Module['_SDL_lroundf'] = _SDL_lroundf = createExportWrapper('SDL_lroundf', 1);
  Module['_SDL_scalbnf'] = _SDL_scalbnf = createExportWrapper('SDL_scalbnf', 2);
  Module['_SDL_tan'] = _SDL_tan = createExportWrapper('SDL_tan', 1);
  Module['_SDL_isalpha'] = _SDL_isalpha = createExportWrapper('SDL_isalpha', 1);
  Module['_SDL_isupper'] = _SDL_isupper = createExportWrapper('SDL_isupper', 1);
  Module['_SDL_islower'] = _SDL_islower = createExportWrapper('SDL_islower', 1);
  Module['_SDL_isalnum'] = _SDL_isalnum = createExportWrapper('SDL_isalnum', 1);
  Module['_SDL_isxdigit'] = _SDL_isxdigit = createExportWrapper('SDL_isxdigit', 1);
  Module['_SDL_ispunct'] = _SDL_ispunct = createExportWrapper('SDL_ispunct', 1);
  Module['_SDL_isgraph'] = _SDL_isgraph = createExportWrapper('SDL_isgraph', 1);
  Module['_SDL_isprint'] = _SDL_isprint = createExportWrapper('SDL_isprint', 1);
  Module['_SDL_toupper'] = _SDL_toupper = createExportWrapper('SDL_toupper', 1);
  Module['_SDL_isblank'] = _SDL_isblank = createExportWrapper('SDL_isblank', 1);
  Module['_SDL_StepBackUTF8'] = _SDL_StepBackUTF8 = createExportWrapper('SDL_StepBackUTF8', 2);
  Module['_SDL_strnlen'] = _SDL_strnlen = createExportWrapper('SDL_strnlen', 2);
  Module['_SDL_wcslen'] = _SDL_wcslen = createExportWrapper('SDL_wcslen', 1);
  Module['_SDL_wcsnlen'] = _SDL_wcsnlen = createExportWrapper('SDL_wcsnlen', 2);
  Module['_SDL_wcslcpy'] = _SDL_wcslcpy = createExportWrapper('SDL_wcslcpy', 3);
  Module['_SDL_wcslcat'] = _SDL_wcslcat = createExportWrapper('SDL_wcslcat', 3);
  Module['_SDL_wcsnstr'] = _SDL_wcsnstr = createExportWrapper('SDL_wcsnstr', 3);
  Module['_SDL_wcsncmp'] = _SDL_wcsncmp = createExportWrapper('SDL_wcsncmp', 3);
  Module['_SDL_wcsstr'] = _SDL_wcsstr = createExportWrapper('SDL_wcsstr', 2);
  Module['_SDL_wcscmp'] = _SDL_wcscmp = createExportWrapper('SDL_wcscmp', 2);
  Module['_SDL_wcscasecmp'] = _SDL_wcscasecmp = createExportWrapper('SDL_wcscasecmp', 2);
  Module['_SDL_wcsncasecmp'] = _SDL_wcsncasecmp = createExportWrapper('SDL_wcsncasecmp', 3);
  Module['_SDL_wcstol'] = _SDL_wcstol = createExportWrapper('SDL_wcstol', 3);
  Module['_SDL_utf8strlcpy'] = _SDL_utf8strlcpy = createExportWrapper('SDL_utf8strlcpy', 3);
  Module['_SDL_utf8strlen'] = _SDL_utf8strlen = createExportWrapper('SDL_utf8strlen', 1);
  Module['_SDL_utf8strnlen'] = _SDL_utf8strnlen = createExportWrapper('SDL_utf8strnlen', 2);
  Module['_SDL_strndup'] = _SDL_strndup = createExportWrapper('SDL_strndup', 2);
  Module['_SDL_strrev'] = _SDL_strrev = createExportWrapper('SDL_strrev', 1);
  Module['_SDL_strupr'] = _SDL_strupr = createExportWrapper('SDL_strupr', 1);
  Module['_SDL_strlwr'] = _SDL_strlwr = createExportWrapper('SDL_strlwr', 1);
  Module['_SDL_strnstr'] = _SDL_strnstr = createExportWrapper('SDL_strnstr', 3);
  Module['_SDL_itoa'] = _SDL_itoa = createExportWrapper('SDL_itoa', 3);
  Module['_SDL_ltoa'] = _SDL_ltoa = createExportWrapper('SDL_ltoa', 3);
  Module['_SDL_uitoa'] = _SDL_uitoa = createExportWrapper('SDL_uitoa', 3);
  Module['_SDL_ultoa'] = _SDL_ultoa = createExportWrapper('SDL_ultoa', 3);
  Module['_SDL_lltoa'] = _SDL_lltoa = createExportWrapper('SDL_lltoa', 3);
  Module['_SDL_ulltoa'] = _SDL_ulltoa = createExportWrapper('SDL_ulltoa', 3);
  Module['_SDL_strtod'] = _SDL_strtod = createExportWrapper('SDL_strtod', 2);
  Module['_SDL_vsscanf'] = _SDL_vsscanf = createExportWrapper('SDL_vsscanf', 3);
  Module['_SDL_swprintf'] = _SDL_swprintf = createExportWrapper('SDL_swprintf', 4);
  Module['_SDL_vswprintf'] = _SDL_vswprintf = createExportWrapper('SDL_vswprintf', 4);
  Module['_SDL_strpbrk'] = _SDL_strpbrk = createExportWrapper('SDL_strpbrk', 2);
  Module['_SDL_OpenTitleStorage'] = _SDL_OpenTitleStorage = createExportWrapper('SDL_OpenTitleStorage', 2);
  Module['_SDL_OpenUserStorage'] = _SDL_OpenUserStorage = createExportWrapper('SDL_OpenUserStorage', 3);
  Module['_SDL_OpenFileStorage'] = _SDL_OpenFileStorage = createExportWrapper('SDL_OpenFileStorage', 1);
  Module['_SDL_OpenStorage'] = _SDL_OpenStorage = createExportWrapper('SDL_OpenStorage', 2);
  Module['_SDL_CloseStorage'] = _SDL_CloseStorage = createExportWrapper('SDL_CloseStorage', 1);
  Module['_SDL_StorageReady'] = _SDL_StorageReady = createExportWrapper('SDL_StorageReady', 1);
  Module['_SDL_GetStorageFileSize'] = _SDL_GetStorageFileSize = createExportWrapper('SDL_GetStorageFileSize', 3);
  Module['_SDL_GetStoragePathInfo'] = _SDL_GetStoragePathInfo = createExportWrapper('SDL_GetStoragePathInfo', 3);
  Module['_SDL_ReadStorageFile'] = _SDL_ReadStorageFile = createExportWrapper('SDL_ReadStorageFile', 4);
  Module['_SDL_WriteStorageFile'] = _SDL_WriteStorageFile = createExportWrapper('SDL_WriteStorageFile', 4);
  Module['_SDL_CreateStorageDirectory'] = _SDL_CreateStorageDirectory = createExportWrapper('SDL_CreateStorageDirectory', 2);
  Module['_SDL_EnumerateStorageDirectory'] = _SDL_EnumerateStorageDirectory = createExportWrapper('SDL_EnumerateStorageDirectory', 4);
  Module['_SDL_RemoveStoragePath'] = _SDL_RemoveStoragePath = createExportWrapper('SDL_RemoveStoragePath', 2);
  Module['_SDL_RenameStoragePath'] = _SDL_RenameStoragePath = createExportWrapper('SDL_RenameStoragePath', 3);
  Module['_SDL_CopyStorageFile'] = _SDL_CopyStorageFile = createExportWrapper('SDL_CopyStorageFile', 3);
  Module['_SDL_GetStorageSpaceRemaining'] = _SDL_GetStorageSpaceRemaining = createExportWrapper('SDL_GetStorageSpaceRemaining', 1);
  Module['_SDL_GlobStorageDirectory'] = _SDL_GlobStorageDirectory = createExportWrapper('SDL_GlobStorageDirectory', 5);
  Module['_SDL_CleanupTLS'] = _SDL_CleanupTLS = createExportWrapper('SDL_CleanupTLS', 0);
  Module['_SDL_GetThreadState'] = _SDL_GetThreadState = createExportWrapper('SDL_GetThreadState', 1);
  Module['_SDL_CreateThreadWithPropertiesRuntime'] = _SDL_CreateThreadWithPropertiesRuntime = createExportWrapper('SDL_CreateThreadWithPropertiesRuntime', 3);
  Module['_SDL_GetThreadID'] = _SDL_GetThreadID = createExportWrapper('SDL_GetThreadID', 1);
  Module['_SDL_GetThreadName'] = _SDL_GetThreadName = createExportWrapper('SDL_GetThreadName', 1);
  Module['_SDL_DetachThread'] = _SDL_DetachThread = createExportWrapper('SDL_DetachThread', 1);
  Module['_SDL_TryWaitSemaphore'] = _SDL_TryWaitSemaphore = createExportWrapper('SDL_TryWaitSemaphore', 1);
  Module['_SDL_WaitSemaphoreTimeout'] = _SDL_WaitSemaphoreTimeout = createExportWrapper('SDL_WaitSemaphoreTimeout', 2);
  Module['_SDL_GetDateTimeLocalePreferences'] = _SDL_GetDateTimeLocalePreferences = createExportWrapper('SDL_GetDateTimeLocalePreferences', 2);
  Module['_SDL_GetDaysInMonth'] = _SDL_GetDaysInMonth = createExportWrapper('SDL_GetDaysInMonth', 2);
  Module['_SDL_GetDayOfYear'] = _SDL_GetDayOfYear = createExportWrapper('SDL_GetDayOfYear', 3);
  Module['_SDL_GetDayOfWeek'] = _SDL_GetDayOfWeek = createExportWrapper('SDL_GetDayOfWeek', 3);
  Module['_SDL_DateTimeToTime'] = _SDL_DateTimeToTime = createExportWrapper('SDL_DateTimeToTime', 2);
  Module['_SDL_TimeToWindows'] = _SDL_TimeToWindows = createExportWrapper('SDL_TimeToWindows', 3);
  Module['_SDL_TimeFromWindows'] = _SDL_TimeFromWindows = createExportWrapper('SDL_TimeFromWindows', 2);
  Module['_SDL_AddTimer'] = _SDL_AddTimer = createExportWrapper('SDL_AddTimer', 3);
  Module['_SDL_AddTimerNS'] = _SDL_AddTimerNS = createExportWrapper('SDL_AddTimerNS', 3);
  Module['_SDL_RemoveTimer'] = _SDL_RemoveTimer = createExportWrapper('SDL_RemoveTimer', 1);
  Module['_SDL_GetPerformanceFrequency'] = _SDL_GetPerformanceFrequency = createExportWrapper('SDL_GetPerformanceFrequency', 0);
  Module['_SDL_GetPixelFormatForMasks'] = _SDL_GetPixelFormatForMasks = createExportWrapper('SDL_GetPixelFormatForMasks', 5);
  Module['_SDL_LoadBMP'] = _SDL_LoadBMP = createExportWrapper('SDL_LoadBMP', 1);
  Module['_SDL_SaveBMP_IO'] = _SDL_SaveBMP_IO = createExportWrapper('SDL_SaveBMP_IO', 3);
  Module['_SDL_SaveBMP'] = _SDL_SaveBMP = createExportWrapper('SDL_SaveBMP', 2);
  Module['_SDL_SetClipboardData'] = _SDL_SetClipboardData = createExportWrapper('SDL_SetClipboardData', 5);
  Module['_SDL_ClearClipboardData'] = _SDL_ClearClipboardData = createExportWrapper('SDL_ClearClipboardData', 0);
  Module['_SDL_GetClipboardData'] = _SDL_GetClipboardData = createExportWrapper('SDL_GetClipboardData', 2);
  Module['_SDL_HasClipboardData'] = _SDL_HasClipboardData = createExportWrapper('SDL_HasClipboardData', 1);
  Module['_SDL_GetClipboardMimeTypes'] = _SDL_GetClipboardMimeTypes = createExportWrapper('SDL_GetClipboardMimeTypes', 1);
  Module['_SDL_SetClipboardText'] = _SDL_SetClipboardText = createExportWrapper('SDL_SetClipboardText', 1);
  Module['_SDL_GetClipboardText'] = _SDL_GetClipboardText = createExportWrapper('SDL_GetClipboardText', 0);
  Module['_SDL_HasClipboardText'] = _SDL_HasClipboardText = createExportWrapper('SDL_HasClipboardText', 0);
  Module['_SDL_SetPrimarySelectionText'] = _SDL_SetPrimarySelectionText = createExportWrapper('SDL_SetPrimarySelectionText', 1);
  Module['_SDL_GetPrimarySelectionText'] = _SDL_GetPrimarySelectionText = createExportWrapper('SDL_GetPrimarySelectionText', 0);
  Module['_SDL_HasPrimarySelectionText'] = _SDL_HasPrimarySelectionText = createExportWrapper('SDL_HasPrimarySelectionText', 0);
  Module['_SDL_UnloadObject'] = _SDL_UnloadObject = createExportWrapper('SDL_UnloadObject', 1);
  Module['_SDL_LoadObject'] = _SDL_LoadObject = createExportWrapper('SDL_LoadObject', 1);
  Module['_SDL_LoadFunction'] = _SDL_LoadFunction = createExportWrapper('SDL_LoadFunction', 2);
  Module['_SDL_GetMasksForPixelFormat'] = _SDL_GetMasksForPixelFormat = createExportWrapper('SDL_GetMasksForPixelFormat', 6);
  Module['_SDL_SetPaletteColors'] = _SDL_SetPaletteColors = createExportWrapper('SDL_SetPaletteColors', 4);
  Module['_SDL_MapRGB'] = _SDL_MapRGB = createExportWrapper('SDL_MapRGB', 5);
  Module['_SDL_GetRGB'] = _SDL_GetRGB = createExportWrapper('SDL_GetRGB', 6);
  Module['_SDL_HasRectIntersection'] = _SDL_HasRectIntersection = createExportWrapper('SDL_HasRectIntersection', 2);
  Module['_SDL_GetRectEnclosingPoints'] = _SDL_GetRectEnclosingPoints = createExportWrapper('SDL_GetRectEnclosingPoints', 4);
  Module['_SDL_HasRectIntersectionFloat'] = _SDL_HasRectIntersectionFloat = createExportWrapper('SDL_HasRectIntersectionFloat', 2);
  Module['_SDL_GetRectUnionFloat'] = _SDL_GetRectUnionFloat = createExportWrapper('SDL_GetRectUnionFloat', 3);
  Module['_SDL_GetRectEnclosingPointsFloat'] = _SDL_GetRectEnclosingPointsFloat = createExportWrapper('SDL_GetRectEnclosingPointsFloat', 4);
  Module['_SDL_GetRectAndLineIntersectionFloat'] = _SDL_GetRectAndLineIntersectionFloat = createExportWrapper('SDL_GetRectAndLineIntersectionFloat', 5);
  Module['_SDL_SurfaceHasRLE'] = _SDL_SurfaceHasRLE = createExportWrapper('SDL_SurfaceHasRLE', 1);
  Module['_SDL_AddSurfaceAlternateImage'] = _SDL_AddSurfaceAlternateImage = createExportWrapper('SDL_AddSurfaceAlternateImage', 2);
  Module['_SDL_SurfaceHasAlternateImages'] = _SDL_SurfaceHasAlternateImages = createExportWrapper('SDL_SurfaceHasAlternateImages', 1);
  Module['_SDL_GetSurfaceImages'] = _SDL_GetSurfaceImages = createExportWrapper('SDL_GetSurfaceImages', 2);
  Module['_SDL_ScaleSurface'] = _SDL_ScaleSurface = createExportWrapper('SDL_ScaleSurface', 4);
  Module['_SDL_RemoveSurfaceAlternateImages'] = _SDL_RemoveSurfaceAlternateImages = createExportWrapper('SDL_RemoveSurfaceAlternateImages', 1);
  Module['_SDL_BlitSurfaceUnchecked'] = _SDL_BlitSurfaceUnchecked = createExportWrapper('SDL_BlitSurfaceUnchecked', 4);
  Module['_SDL_BlitSurfaceUncheckedScaled'] = _SDL_BlitSurfaceUncheckedScaled = createExportWrapper('SDL_BlitSurfaceUncheckedScaled', 5);
  Module['_SDL_BlitSurfaceTiled'] = _SDL_BlitSurfaceTiled = createExportWrapper('SDL_BlitSurfaceTiled', 4);
  Module['_SDL_BlitSurfaceTiledWithScale'] = _SDL_BlitSurfaceTiledWithScale = createExportWrapper('SDL_BlitSurfaceTiledWithScale', 6);
  Module['_SDL_BlitSurface9Grid'] = _SDL_BlitSurface9Grid = createExportWrapper('SDL_BlitSurface9Grid', 10);
  Module['_SDL_PremultiplyAlpha'] = _SDL_PremultiplyAlpha = createExportWrapper('SDL_PremultiplyAlpha', 9);
  Module['_SDL_PremultiplySurfaceAlpha'] = _SDL_PremultiplySurfaceAlpha = createExportWrapper('SDL_PremultiplySurfaceAlpha', 2);
  Module['_SDL_ClearSurface'] = _SDL_ClearSurface = createExportWrapper('SDL_ClearSurface', 5);
  Module['_SDL_MapSurfaceRGB'] = _SDL_MapSurfaceRGB = createExportWrapper('SDL_MapSurfaceRGB', 4);
  Module['_SDL_ReadSurfacePixel'] = _SDL_ReadSurfacePixel = createExportWrapper('SDL_ReadSurfacePixel', 7);
  Module['_SDL_ReadSurfacePixelFloat'] = _SDL_ReadSurfacePixelFloat = createExportWrapper('SDL_ReadSurfacePixelFloat', 7);
  Module['_SDL_WriteSurfacePixel'] = _SDL_WriteSurfacePixel = createExportWrapper('SDL_WriteSurfacePixel', 7);
  Module['_SDL_WriteSurfacePixelFloat'] = _SDL_WriteSurfacePixelFloat = createExportWrapper('SDL_WriteSurfacePixelFloat', 7);
  Module['_SDL_GetNumVideoDrivers'] = _SDL_GetNumVideoDrivers = createExportWrapper('SDL_GetNumVideoDrivers', 0);
  Module['_SDL_GetVideoDriver'] = _SDL_GetVideoDriver = createExportWrapper('SDL_GetVideoDriver', 1);
  Module['_SDL_GL_ResetAttributes'] = _SDL_GL_ResetAttributes = createExportWrapper('SDL_GL_ResetAttributes', 0);
  Module['_SDL_DisableScreenSaver'] = _SDL_DisableScreenSaver = createExportWrapper('SDL_DisableScreenSaver', 0);
  Module['_SDL_GetSystemTheme'] = _SDL_GetSystemTheme = createExportWrapper('SDL_GetSystemTheme', 0);
  Module['_SDL_GetDisplayBounds'] = _SDL_GetDisplayBounds = createExportWrapper('SDL_GetDisplayBounds', 2);
  Module['_SDL_GetDisplays'] = _SDL_GetDisplays = createExportWrapper('SDL_GetDisplays', 1);
  Module['_SDL_GetDisplayProperties'] = _SDL_GetDisplayProperties = createExportWrapper('SDL_GetDisplayProperties', 1);
  Module['_SDL_GetDisplayName'] = _SDL_GetDisplayName = createExportWrapper('SDL_GetDisplayName', 1);
  Module['_SDL_GetDisplayUsableBounds'] = _SDL_GetDisplayUsableBounds = createExportWrapper('SDL_GetDisplayUsableBounds', 2);
  Module['_SDL_GetCurrentDisplayOrientation'] = _SDL_GetCurrentDisplayOrientation = createExportWrapper('SDL_GetCurrentDisplayOrientation', 1);
  Module['_SDL_GetWindowPixelDensity'] = _SDL_GetWindowPixelDensity = createExportWrapper('SDL_GetWindowPixelDensity', 1);
  Module['_SDL_GetDisplayContentScale'] = _SDL_GetDisplayContentScale = createExportWrapper('SDL_GetDisplayContentScale', 1);
  Module['_SDL_GetFullscreenDisplayModes'] = _SDL_GetFullscreenDisplayModes = createExportWrapper('SDL_GetFullscreenDisplayModes', 2);
  Module['_SDL_GetClosestFullscreenDisplayMode'] = _SDL_GetClosestFullscreenDisplayMode = createExportWrapper('SDL_GetClosestFullscreenDisplayMode', 6);
  Module['_SDL_GetCurrentDisplayMode'] = _SDL_GetCurrentDisplayMode = createExportWrapper('SDL_GetCurrentDisplayMode', 1);
  Module['_SDL_GetDisplayForPoint'] = _SDL_GetDisplayForPoint = createExportWrapper('SDL_GetDisplayForPoint', 1);
  Module['_SDL_GetDisplayForRect'] = _SDL_GetDisplayForRect = createExportWrapper('SDL_GetDisplayForRect', 1);
  Module['_SDL_GetWindowDisplayScale'] = _SDL_GetWindowDisplayScale = createExportWrapper('SDL_GetWindowDisplayScale', 1);
  Module['_SDL_GetWindowFullscreenMode'] = _SDL_GetWindowFullscreenMode = createExportWrapper('SDL_GetWindowFullscreenMode', 1);
  Module['_SDL_SetWindowFullscreenMode'] = _SDL_SetWindowFullscreenMode = createExportWrapper('SDL_SetWindowFullscreenMode', 2);
  Module['_SDL_GetWindowICCProfile'] = _SDL_GetWindowICCProfile = createExportWrapper('SDL_GetWindowICCProfile', 2);
  Module['_SDL_GetWindowPixelFormat'] = _SDL_GetWindowPixelFormat = createExportWrapper('SDL_GetWindowPixelFormat', 1);
  Module['_SDL_GetWindows'] = _SDL_GetWindows = createExportWrapper('SDL_GetWindows', 1);
  Module['_SDL_CreateWindowWithProperties'] = _SDL_CreateWindowWithProperties = createExportWrapper('SDL_CreateWindowWithProperties', 1);
  Module['_SDL_Vulkan_LoadLibrary'] = _SDL_Vulkan_LoadLibrary = createExportWrapper('SDL_Vulkan_LoadLibrary', 1);
  Module['_SDL_SetWindowTitle'] = _SDL_SetWindowTitle = createExportWrapper('SDL_SetWindowTitle', 2);
  Module['_SDL_GL_LoadLibrary'] = _SDL_GL_LoadLibrary = createExportWrapper('SDL_GL_LoadLibrary', 1);
  Module['_SDL_HideWindow'] = _SDL_HideWindow = createExportWrapper('SDL_HideWindow', 1);
  Module['_SDL_CreatePopupWindow'] = _SDL_CreatePopupWindow = createExportWrapper('SDL_CreatePopupWindow', 6);
  Module['_SDL_SetWindowModal'] = _SDL_SetWindowModal = createExportWrapper('SDL_SetWindowModal', 2);
  Module['_SDL_GL_UnloadLibrary'] = _SDL_GL_UnloadLibrary = createExportWrapper('SDL_GL_UnloadLibrary', 0);
  Module['_SDL_Vulkan_UnloadLibrary'] = _SDL_Vulkan_UnloadLibrary = createExportWrapper('SDL_Vulkan_UnloadLibrary', 0);
  Module['_SDL_GetWindowParent'] = _SDL_GetWindowParent = createExportWrapper('SDL_GetWindowParent', 1);
  Module['_SDL_GetWindowTitle'] = _SDL_GetWindowTitle = createExportWrapper('SDL_GetWindowTitle', 1);
  Module['_SDL_SetWindowIcon'] = _SDL_SetWindowIcon = createExportWrapper('SDL_SetWindowIcon', 2);
  Module['_SDL_SetWindowPosition'] = _SDL_SetWindowPosition = createExportWrapper('SDL_SetWindowPosition', 3);
  Module['_SDL_GetWindowPosition'] = _SDL_GetWindowPosition = createExportWrapper('SDL_GetWindowPosition', 3);
  Module['_SDL_SetWindowBordered'] = _SDL_SetWindowBordered = createExportWrapper('SDL_SetWindowBordered', 2);
  Module['_SDL_SetWindowResizable'] = _SDL_SetWindowResizable = createExportWrapper('SDL_SetWindowResizable', 2);
  Module['_SDL_SetWindowAlwaysOnTop'] = _SDL_SetWindowAlwaysOnTop = createExportWrapper('SDL_SetWindowAlwaysOnTop', 2);
  Module['_SDL_SetWindowSize'] = _SDL_SetWindowSize = createExportWrapper('SDL_SetWindowSize', 3);
  Module['_SDL_SetWindowAspectRatio'] = _SDL_SetWindowAspectRatio = createExportWrapper('SDL_SetWindowAspectRatio', 3);
  Module['_SDL_GetWindowAspectRatio'] = _SDL_GetWindowAspectRatio = createExportWrapper('SDL_GetWindowAspectRatio', 3);
  Module['_SDL_GetWindowBordersSize'] = _SDL_GetWindowBordersSize = createExportWrapper('SDL_GetWindowBordersSize', 5);
  Module['_SDL_SetWindowMinimumSize'] = _SDL_SetWindowMinimumSize = createExportWrapper('SDL_SetWindowMinimumSize', 3);
  Module['_SDL_GetWindowMinimumSize'] = _SDL_GetWindowMinimumSize = createExportWrapper('SDL_GetWindowMinimumSize', 3);
  Module['_SDL_SetWindowMaximumSize'] = _SDL_SetWindowMaximumSize = createExportWrapper('SDL_SetWindowMaximumSize', 3);
  Module['_SDL_GetWindowMaximumSize'] = _SDL_GetWindowMaximumSize = createExportWrapper('SDL_GetWindowMaximumSize', 3);
  Module['_SDL_RaiseWindow'] = _SDL_RaiseWindow = createExportWrapper('SDL_RaiseWindow', 1);
  Module['_SDL_MaximizeWindow'] = _SDL_MaximizeWindow = createExportWrapper('SDL_MaximizeWindow', 1);
  Module['_SDL_SetWindowFullscreen'] = _SDL_SetWindowFullscreen = createExportWrapper('SDL_SetWindowFullscreen', 2);
  Module['_SDL_SetWindowSurfaceVSync'] = _SDL_SetWindowSurfaceVSync = createExportWrapper('SDL_SetWindowSurfaceVSync', 2);
  Module['_SDL_GetWindowSurfaceVSync'] = _SDL_GetWindowSurfaceVSync = createExportWrapper('SDL_GetWindowSurfaceVSync', 2);
  Module['_SDL_UpdateWindowSurfaceRects'] = _SDL_UpdateWindowSurfaceRects = createExportWrapper('SDL_UpdateWindowSurfaceRects', 3);
  Module['_SDL_SetWindowOpacity'] = _SDL_SetWindowOpacity = createExportWrapper('SDL_SetWindowOpacity', 2);
  Module['_SDL_GetWindowOpacity'] = _SDL_GetWindowOpacity = createExportWrapper('SDL_GetWindowOpacity', 1);
  Module['_SDL_SetWindowParent'] = _SDL_SetWindowParent = createExportWrapper('SDL_SetWindowParent', 2);
  Module['_SDL_SetWindowFocusable'] = _SDL_SetWindowFocusable = createExportWrapper('SDL_SetWindowFocusable', 2);
  Module['_SDL_SetWindowKeyboardGrab'] = _SDL_SetWindowKeyboardGrab = createExportWrapper('SDL_SetWindowKeyboardGrab', 2);
  Module['_SDL_SetWindowMouseGrab'] = _SDL_SetWindowMouseGrab = createExportWrapper('SDL_SetWindowMouseGrab', 2);
  Module['_SDL_GetWindowKeyboardGrab'] = _SDL_GetWindowKeyboardGrab = createExportWrapper('SDL_GetWindowKeyboardGrab', 1);
  Module['_SDL_GetWindowMouseGrab'] = _SDL_GetWindowMouseGrab = createExportWrapper('SDL_GetWindowMouseGrab', 1);
  Module['_SDL_GetGrabbedWindow'] = _SDL_GetGrabbedWindow = createExportWrapper('SDL_GetGrabbedWindow', 0);
  Module['_SDL_SetWindowMouseRect'] = _SDL_SetWindowMouseRect = createExportWrapper('SDL_SetWindowMouseRect', 2);
  Module['_SDL_SetWindowRelativeMouseMode'] = _SDL_SetWindowRelativeMouseMode = createExportWrapper('SDL_SetWindowRelativeMouseMode', 2);
  Module['_SDL_GetWindowRelativeMouseMode'] = _SDL_GetWindowRelativeMouseMode = createExportWrapper('SDL_GetWindowRelativeMouseMode', 1);
  Module['_SDL_FlashWindow'] = _SDL_FlashWindow = createExportWrapper('SDL_FlashWindow', 2);
  Module['_SDL_ScreenSaverEnabled'] = _SDL_ScreenSaverEnabled = createExportWrapper('SDL_ScreenSaverEnabled', 0);
  Module['_SDL_EnableScreenSaver'] = _SDL_EnableScreenSaver = createExportWrapper('SDL_EnableScreenSaver', 0);
  Module['_SDL_EGL_GetProcAddress'] = _SDL_EGL_GetProcAddress = createExportWrapper('SDL_EGL_GetProcAddress', 1);
  Module['_SDL_EGL_SetAttributeCallbacks'] = _SDL_EGL_SetAttributeCallbacks = createExportWrapper('SDL_EGL_SetAttributeCallbacks', 4);
  Module['_SDL_GL_GetCurrentWindow'] = _SDL_GL_GetCurrentWindow = createExportWrapper('SDL_GL_GetCurrentWindow', 0);
  Module['_SDL_EGL_GetCurrentDisplay'] = _SDL_EGL_GetCurrentDisplay = createExportWrapper('SDL_EGL_GetCurrentDisplay', 0);
  Module['_SDL_EGL_GetCurrentConfig'] = _SDL_EGL_GetCurrentConfig = createExportWrapper('SDL_EGL_GetCurrentConfig', 0);
  Module['_SDL_EGL_GetWindowSurface'] = _SDL_EGL_GetWindowSurface = createExportWrapper('SDL_EGL_GetWindowSurface', 1);
  Module['_SDL_StartTextInput'] = _SDL_StartTextInput = createExportWrapper('SDL_StartTextInput', 1);
  Module['_SDL_StartTextInputWithProperties'] = _SDL_StartTextInputWithProperties = createExportWrapper('SDL_StartTextInputWithProperties', 2);
  Module['_SDL_ScreenKeyboardShown'] = _SDL_ScreenKeyboardShown = createExportWrapper('SDL_ScreenKeyboardShown', 1);
  Module['_SDL_StopTextInput'] = _SDL_StopTextInput = createExportWrapper('SDL_StopTextInput', 1);
  Module['_SDL_SetTextInputArea'] = _SDL_SetTextInputArea = createExportWrapper('SDL_SetTextInputArea', 3);
  Module['_SDL_GetTextInputArea'] = _SDL_GetTextInputArea = createExportWrapper('SDL_GetTextInputArea', 3);
  Module['_SDL_ClearComposition'] = _SDL_ClearComposition = createExportWrapper('SDL_ClearComposition', 1);
  Module['_SDL_HasScreenKeyboardSupport'] = _SDL_HasScreenKeyboardSupport = createExportWrapper('SDL_HasScreenKeyboardSupport', 0);
  Module['_SDL_ShowSimpleMessageBox'] = _SDL_ShowSimpleMessageBox = createExportWrapper('SDL_ShowSimpleMessageBox', 4);
  Module['_SDL_ShowWindowSystemMenu'] = _SDL_ShowWindowSystemMenu = createExportWrapper('SDL_ShowWindowSystemMenu', 3);
  Module['_SDL_SetWindowHitTest'] = _SDL_SetWindowHitTest = createExportWrapper('SDL_SetWindowHitTest', 3);
  Module['_SDL_SetWindowShape'] = _SDL_SetWindowShape = createExportWrapper('SDL_SetWindowShape', 2);
  Module['_SDL_OnApplicationWillTerminate'] = _SDL_OnApplicationWillTerminate = createExportWrapper('SDL_OnApplicationWillTerminate', 0);
  Module['_SDL_OnApplicationDidReceiveMemoryWarning'] = _SDL_OnApplicationDidReceiveMemoryWarning = createExportWrapper('SDL_OnApplicationDidReceiveMemoryWarning', 0);
  Module['_SDL_OnApplicationWillEnterBackground'] = _SDL_OnApplicationWillEnterBackground = createExportWrapper('SDL_OnApplicationWillEnterBackground', 0);
  Module['_SDL_OnApplicationDidEnterBackground'] = _SDL_OnApplicationDidEnterBackground = createExportWrapper('SDL_OnApplicationDidEnterBackground', 0);
  Module['_SDL_OnApplicationWillEnterForeground'] = _SDL_OnApplicationWillEnterForeground = createExportWrapper('SDL_OnApplicationWillEnterForeground', 0);
  Module['_SDL_OnApplicationDidEnterForeground'] = _SDL_OnApplicationDidEnterForeground = createExportWrapper('SDL_OnApplicationDidEnterForeground', 0);
  Module['_SDL_Vulkan_GetVkGetInstanceProcAddr'] = _SDL_Vulkan_GetVkGetInstanceProcAddr = createExportWrapper('SDL_Vulkan_GetVkGetInstanceProcAddr', 0);
  Module['_SDL_Vulkan_GetInstanceExtensions'] = _SDL_Vulkan_GetInstanceExtensions = createExportWrapper('SDL_Vulkan_GetInstanceExtensions', 1);
  Module['_SDL_Vulkan_CreateSurface'] = _SDL_Vulkan_CreateSurface = createExportWrapper('SDL_Vulkan_CreateSurface', 4);
  Module['_SDL_Vulkan_DestroySurface'] = _SDL_Vulkan_DestroySurface = createExportWrapper('SDL_Vulkan_DestroySurface', 3);
  Module['_SDL_Vulkan_GetPresentationSupport'] = _SDL_Vulkan_GetPresentationSupport = createExportWrapper('SDL_Vulkan_GetPresentationSupport', 3);
  Module['_SDL_Metal_CreateView'] = _SDL_Metal_CreateView = createExportWrapper('SDL_Metal_CreateView', 1);
  Module['_SDL_Metal_DestroyView'] = _SDL_Metal_DestroyView = createExportWrapper('SDL_Metal_DestroyView', 1);
  Module['_SDL_Metal_GetLayer'] = _SDL_Metal_GetLayer = createExportWrapper('SDL_Metal_GetLayer', 1);
  Module['_SDL_GetDXGIOutputInfo'] = _SDL_GetDXGIOutputInfo = createExportWrapper('SDL_GetDXGIOutputInfo', 3);
  Module['_SDL_GetDirect3D9AdapterIndex'] = _SDL_GetDirect3D9AdapterIndex = createExportWrapper('SDL_GetDirect3D9AdapterIndex', 1);
  Module['_SDL_GetGDKTaskQueue'] = _SDL_GetGDKTaskQueue = createExportWrapper('SDL_GetGDKTaskQueue', 1);
  Module['_SDL_OnApplicationDidChangeStatusBarOrientation'] = _SDL_OnApplicationDidChangeStatusBarOrientation = createExportWrapper('SDL_OnApplicationDidChangeStatusBarOrientation', 0);
  Module['_SDL_SetiOSAnimationCallback'] = _SDL_SetiOSAnimationCallback = createExportWrapper('SDL_SetiOSAnimationCallback', 4);
  Module['_SDL_SetiOSEventPump'] = _SDL_SetiOSEventPump = createExportWrapper('SDL_SetiOSEventPump', 1);
  Module['_SDL_EnterAppMainCallbacks'] = _SDL_EnterAppMainCallbacks = createExportWrapper('SDL_EnterAppMainCallbacks', 6);
  Module['_SDL_RunApp'] = _SDL_RunApp = createExportWrapper('SDL_RunApp', 4);
  Module['_SDL_GetCurrentTime'] = _SDL_GetCurrentTime = createExportWrapper('SDL_GetCurrentTime', 1);
  Module['_SDL_TimeToDateTime'] = _SDL_TimeToDateTime = createExportWrapper('SDL_TimeToDateTime', 3);
  Module['_Emscripten_HandlePointerEnter'] = _Emscripten_HandlePointerEnter = createExportWrapper('Emscripten_HandlePointerEnter', 2);
  Module['_Emscripten_HandlePointerLeave'] = _Emscripten_HandlePointerLeave = createExportWrapper('Emscripten_HandlePointerLeave', 2);
  Module['_Emscripten_HandlePointerGeneric'] = _Emscripten_HandlePointerGeneric = createExportWrapper('Emscripten_HandlePointerGeneric', 2);
  Module['_Emscripten_SendDragEvent'] = _Emscripten_SendDragEvent = createExportWrapper('Emscripten_SendDragEvent', 2);
  Module['_Emscripten_SendDragCompleteEvent'] = _Emscripten_SendDragCompleteEvent = createExportWrapper('Emscripten_SendDragCompleteEvent', 1);
  Module['_Emscripten_SendDragTextEvent'] = _Emscripten_SendDragTextEvent = createExportWrapper('Emscripten_SendDragTextEvent', 2);
  Module['_Emscripten_SendDragFileEvent'] = _Emscripten_SendDragFileEvent = createExportWrapper('Emscripten_SendDragFileEvent', 2);
  Module['_Emscripten_SendSystemThemeChangedEvent'] = _Emscripten_SendSystemThemeChangedEvent = createExportWrapper('Emscripten_SendSystemThemeChangedEvent', 0);
  Module['_requestFullscreenThroughSDL'] = _requestFullscreenThroughSDL = createExportWrapper('requestFullscreenThroughSDL', 1);
  Module['_SDL_ShowFileDialogWithProperties'] = _SDL_ShowFileDialogWithProperties = createExportWrapper('SDL_ShowFileDialogWithProperties', 4);
  Module['_SDL_ShowOpenFileDialog'] = _SDL_ShowOpenFileDialog = createExportWrapper('SDL_ShowOpenFileDialog', 7);
  Module['_SDL_ShowSaveFileDialog'] = _SDL_ShowSaveFileDialog = createExportWrapper('SDL_ShowSaveFileDialog', 6);
  Module['_SDL_ShowOpenFolderDialog'] = _SDL_ShowOpenFolderDialog = createExportWrapper('SDL_ShowOpenFolderDialog', 5);
  Module['_SDL_CreateProcessWithProperties'] = _SDL_CreateProcessWithProperties = createExportWrapper('SDL_CreateProcessWithProperties', 1);
  Module['_SDL_ReadProcess'] = _SDL_ReadProcess = createExportWrapper('SDL_ReadProcess', 3);
  Module['_SDL_DestroyProcess'] = _SDL_DestroyProcess = createExportWrapper('SDL_DestroyProcess', 1);
  Module['_SDL_WaitProcess'] = _SDL_WaitProcess = createExportWrapper('SDL_WaitProcess', 3);
  Module['_SDL_CreateProcess'] = _SDL_CreateProcess = createExportWrapper('SDL_CreateProcess', 2);
  Module['_SDL_GetProcessProperties'] = _SDL_GetProcessProperties = createExportWrapper('SDL_GetProcessProperties', 1);
  Module['_SDL_GetProcessInput'] = _SDL_GetProcessInput = createExportWrapper('SDL_GetProcessInput', 1);
  Module['_SDL_GetProcessOutput'] = _SDL_GetProcessOutput = createExportWrapper('SDL_GetProcessOutput', 1);
  Module['_SDL_KillProcess'] = _SDL_KillProcess = createExportWrapper('SDL_KillProcess', 2);
  Module['_SDL_DestroyTray'] = _SDL_DestroyTray = createExportWrapper('SDL_DestroyTray', 1);
  Module['_SDL_CreateTray'] = _SDL_CreateTray = createExportWrapper('SDL_CreateTray', 2);
  Module['_SDL_SetTrayIcon'] = _SDL_SetTrayIcon = createExportWrapper('SDL_SetTrayIcon', 2);
  Module['_SDL_SetTrayTooltip'] = _SDL_SetTrayTooltip = createExportWrapper('SDL_SetTrayTooltip', 2);
  Module['_SDL_CreateTrayMenu'] = _SDL_CreateTrayMenu = createExportWrapper('SDL_CreateTrayMenu', 1);
  Module['_SDL_GetTrayMenu'] = _SDL_GetTrayMenu = createExportWrapper('SDL_GetTrayMenu', 1);
  Module['_SDL_CreateTraySubmenu'] = _SDL_CreateTraySubmenu = createExportWrapper('SDL_CreateTraySubmenu', 1);
  Module['_SDL_GetTraySubmenu'] = _SDL_GetTraySubmenu = createExportWrapper('SDL_GetTraySubmenu', 1);
  Module['_SDL_GetTrayEntries'] = _SDL_GetTrayEntries = createExportWrapper('SDL_GetTrayEntries', 2);
  Module['_SDL_RemoveTrayEntry'] = _SDL_RemoveTrayEntry = createExportWrapper('SDL_RemoveTrayEntry', 1);
  Module['_SDL_InsertTrayEntryAt'] = _SDL_InsertTrayEntryAt = createExportWrapper('SDL_InsertTrayEntryAt', 4);
  Module['_SDL_SetTrayEntryLabel'] = _SDL_SetTrayEntryLabel = createExportWrapper('SDL_SetTrayEntryLabel', 2);
  Module['_SDL_GetTrayEntryLabel'] = _SDL_GetTrayEntryLabel = createExportWrapper('SDL_GetTrayEntryLabel', 1);
  Module['_SDL_SetTrayEntryChecked'] = _SDL_SetTrayEntryChecked = createExportWrapper('SDL_SetTrayEntryChecked', 2);
  Module['_SDL_GetTrayEntryChecked'] = _SDL_GetTrayEntryChecked = createExportWrapper('SDL_GetTrayEntryChecked', 1);
  Module['_SDL_SetTrayEntryEnabled'] = _SDL_SetTrayEntryEnabled = createExportWrapper('SDL_SetTrayEntryEnabled', 2);
  Module['_SDL_GetTrayEntryEnabled'] = _SDL_GetTrayEntryEnabled = createExportWrapper('SDL_GetTrayEntryEnabled', 1);
  Module['_SDL_SetTrayEntryCallback'] = _SDL_SetTrayEntryCallback = createExportWrapper('SDL_SetTrayEntryCallback', 3);
  Module['_SDL_ClickTrayEntry'] = _SDL_ClickTrayEntry = createExportWrapper('SDL_ClickTrayEntry', 1);
  Module['_SDL_GetTrayEntryParent'] = _SDL_GetTrayEntryParent = createExportWrapper('SDL_GetTrayEntryParent', 1);
  Module['_SDL_GetTrayMenuParentEntry'] = _SDL_GetTrayMenuParentEntry = createExportWrapper('SDL_GetTrayMenuParentEntry', 1);
  Module['_SDL_GetTrayMenuParentTray'] = _SDL_GetTrayMenuParentTray = createExportWrapper('SDL_GetTrayMenuParentTray', 1);
  Module['_SDL_TryLockMutex'] = _SDL_TryLockMutex = createExportWrapper('SDL_TryLockMutex', 1);
  Module['_SDL_TryLockRWLockForReading'] = _SDL_TryLockRWLockForReading = createExportWrapper('SDL_TryLockRWLockForReading', 1);
  Module['_SDL_TryLockRWLockForWriting'] = _SDL_TryLockRWLockForWriting = createExportWrapper('SDL_TryLockRWLockForWriting', 1);
  Module['_SDL_GetSemaphoreValue'] = _SDL_GetSemaphoreValue = createExportWrapper('SDL_GetSemaphoreValue', 1);
  Module['_emscripten_stack_get_end'] = _emscripten_stack_get_end = wasmExports['emscripten_stack_get_end'];
  Module['_emscripten_stack_get_base'] = _emscripten_stack_get_base = wasmExports['emscripten_stack_get_base'];
  Module['_memcpy'] = _memcpy = createExportWrapper('memcpy', 3);
  Module['__emscripten_memcpy_bulkmem'] = __emscripten_memcpy_bulkmem = createExportWrapper('_emscripten_memcpy_bulkmem', 3);
  Module['__emscripten_memset_bulkmem'] = __emscripten_memset_bulkmem = createExportWrapper('_emscripten_memset_bulkmem', 3);
  Module['_emscripten_builtin_memalign'] = _emscripten_builtin_memalign = createExportWrapper('emscripten_builtin_memalign', 2);
  Module['_emscripten_stack_get_current'] = _emscripten_stack_get_current = wasmExports['emscripten_stack_get_current'];
  Module['_htons'] = _htons = createExportWrapper('htons', 1);
  Module['_ntohs'] = _ntohs = createExportWrapper('ntohs', 1);
  Module['_htonl'] = _htonl = createExportWrapper('htonl', 1);
  Module['__emscripten_timeout'] = __emscripten_timeout = createExportWrapper('_emscripten_timeout', 2);
  Module['_setThrew'] = _setThrew = createExportWrapper('setThrew', 2);
  Module['__emscripten_tempret_set'] = __emscripten_tempret_set = createExportWrapper('_emscripten_tempret_set', 1);
  Module['__emscripten_tempret_get'] = __emscripten_tempret_get = createExportWrapper('_emscripten_tempret_get', 0);
  Module['___get_temp_ret'] = ___get_temp_ret = createExportWrapper('__get_temp_ret', 0);
  Module['___set_temp_ret'] = ___set_temp_ret = createExportWrapper('__set_temp_ret', 1);
  Module['___emutls_get_address'] = ___emutls_get_address = createExportWrapper('__emutls_get_address', 1);
  Module['_emscripten_stack_init'] = _emscripten_stack_init = wasmExports['emscripten_stack_init'];
  Module['_emscripten_stack_set_limits'] = _emscripten_stack_set_limits = wasmExports['emscripten_stack_set_limits'];
  Module['_emscripten_stack_get_free'] = _emscripten_stack_get_free = wasmExports['emscripten_stack_get_free'];
  Module['__emscripten_stack_restore'] = __emscripten_stack_restore = wasmExports['_emscripten_stack_restore'];
  Module['__emscripten_stack_alloc'] = __emscripten_stack_alloc = wasmExports['_emscripten_stack_alloc'];
  Module['__ZNSt8bad_castD2Ev'] = __ZNSt8bad_castD2Ev = createExportWrapper('_ZNSt8bad_castD2Ev', 1);
  Module['__ZdlPvm'] = __ZdlPvm = createExportWrapper('_ZdlPvm', 2);
  Module['__Znwm'] = __Znwm = createExportWrapper('_Znwm', 1);
  Module['__ZnamSt11align_val_t'] = __ZnamSt11align_val_t = createExportWrapper('_ZnamSt11align_val_t', 2);
  Module['__ZdaPvSt11align_val_t'] = __ZdaPvSt11align_val_t = createExportWrapper('_ZdaPvSt11align_val_t', 2);
  Module['__ZNSt13runtime_errorD2Ev'] = __ZNSt13runtime_errorD2Ev = createExportWrapper('_ZNSt13runtime_errorD2Ev', 1);
  Module['__ZNKSt13runtime_error4whatEv'] = __ZNKSt13runtime_error4whatEv = createExportWrapper('_ZNKSt13runtime_error4whatEv', 1);
  Module['__ZnwmSt11align_val_t'] = __ZnwmSt11align_val_t = createExportWrapper('_ZnwmSt11align_val_t', 2);
  Module['__ZdlPvmSt11align_val_t'] = __ZdlPvmSt11align_val_t = createExportWrapper('_ZdlPvmSt11align_val_t', 3);
  Module['___cxa_pure_virtual'] = ___cxa_pure_virtual = createExportWrapper('__cxa_pure_virtual', 0);
  Module['___cxa_uncaught_exceptions'] = ___cxa_uncaught_exceptions = createExportWrapper('__cxa_uncaught_exceptions', 0);
  Module['___cxa_decrement_exception_refcount'] = ___cxa_decrement_exception_refcount = createExportWrapper('__cxa_decrement_exception_refcount', 1);
  Module['___cxa_increment_exception_refcount'] = ___cxa_increment_exception_refcount = createExportWrapper('__cxa_increment_exception_refcount', 1);
  Module['___cxa_current_primary_exception'] = ___cxa_current_primary_exception = createExportWrapper('__cxa_current_primary_exception', 0);
  Module['__ZSt9terminatev'] = __ZSt9terminatev = createExportWrapper('_ZSt9terminatev', 0);
  Module['___cxa_rethrow_primary_exception'] = ___cxa_rethrow_primary_exception = createExportWrapper('__cxa_rethrow_primary_exception', 1);
  Module['__ZNSt9exceptionD2Ev'] = __ZNSt9exceptionD2Ev = createExportWrapper('_ZNSt9exceptionD2Ev', 1);
  Module['__ZNSt11logic_errorD2Ev'] = __ZNSt11logic_errorD2Ev = createExportWrapper('_ZNSt11logic_errorD2Ev', 1);
  Module['__ZNKSt11logic_error4whatEv'] = __ZNKSt11logic_error4whatEv = createExportWrapper('_ZNKSt11logic_error4whatEv', 1);
  Module['__ZdaPv'] = __ZdaPv = createExportWrapper('_ZdaPv', 1);
  Module['__Znam'] = __Znam = createExportWrapper('_Znam', 1);
  Module['__ZSt15get_new_handlerv'] = __ZSt15get_new_handlerv = createExportWrapper('_ZSt15get_new_handlerv', 0);
  Module['__ZdlPv'] = __ZdlPv = createExportWrapper('_ZdlPv', 1);
  Module['__ZdaPvm'] = __ZdaPvm = createExportWrapper('_ZdaPvm', 2);
  Module['__ZdlPvSt11align_val_t'] = __ZdlPvSt11align_val_t = createExportWrapper('_ZdlPvSt11align_val_t', 2);
  Module['__ZdaPvmSt11align_val_t'] = __ZdaPvmSt11align_val_t = createExportWrapper('_ZdaPvmSt11align_val_t', 3);
  Module['___dynamic_cast'] = ___dynamic_cast = createExportWrapper('__dynamic_cast', 4);
  Module['___cxa_bad_cast'] = ___cxa_bad_cast = createExportWrapper('__cxa_bad_cast', 0);
  Module['___cxa_bad_typeid'] = ___cxa_bad_typeid = createExportWrapper('__cxa_bad_typeid', 0);
  Module['___cxa_throw_bad_array_new_length'] = ___cxa_throw_bad_array_new_length = createExportWrapper('__cxa_throw_bad_array_new_length', 0);
  Module['__ZSt14set_unexpectedPFvvE'] = __ZSt14set_unexpectedPFvvE = createExportWrapper('_ZSt14set_unexpectedPFvvE', 1);
  Module['__ZSt13set_terminatePFvvE'] = __ZSt13set_terminatePFvvE = createExportWrapper('_ZSt13set_terminatePFvvE', 1);
  Module['__ZSt15set_new_handlerPFvvE'] = __ZSt15set_new_handlerPFvvE = createExportWrapper('_ZSt15set_new_handlerPFvvE', 1);
  Module['___cxa_demangle'] = ___cxa_demangle = createExportWrapper('__cxa_demangle', 4);
  Module['___cxa_guard_acquire'] = ___cxa_guard_acquire = createExportWrapper('__cxa_guard_acquire', 1);
  Module['___cxa_guard_release'] = ___cxa_guard_release = createExportWrapper('__cxa_guard_release', 1);
  Module['___cxa_guard_abort'] = ___cxa_guard_abort = createExportWrapper('__cxa_guard_abort', 1);
  Module['__ZSt14get_unexpectedv'] = __ZSt14get_unexpectedv = createExportWrapper('_ZSt14get_unexpectedv', 0);
  Module['__ZSt10unexpectedv'] = __ZSt10unexpectedv = createExportWrapper('_ZSt10unexpectedv', 0);
  Module['__ZSt13get_terminatev'] = __ZSt13get_terminatev = createExportWrapper('_ZSt13get_terminatev', 0);
  Module['___cxa_uncaught_exception'] = ___cxa_uncaught_exception = createExportWrapper('__cxa_uncaught_exception', 0);
  Module['___cxa_allocate_exception'] = ___cxa_allocate_exception = createExportWrapper('__cxa_allocate_exception', 1);
  Module['___cxa_free_exception'] = ___cxa_free_exception = createExportWrapper('__cxa_free_exception', 1);
  Module['___cxa_init_primary_exception'] = ___cxa_init_primary_exception = createExportWrapper('__cxa_init_primary_exception', 3);
  Module['___cxa_thread_atexit'] = ___cxa_thread_atexit = createExportWrapper('__cxa_thread_atexit', 3);
  Module['___cxa_deleted_virtual'] = ___cxa_deleted_virtual = createExportWrapper('__cxa_deleted_virtual', 0);
  Module['__ZNSt9type_infoD2Ev'] = __ZNSt9type_infoD2Ev = createExportWrapper('_ZNSt9type_infoD2Ev', 1);
  Module['___cxa_can_catch'] = ___cxa_can_catch = createExportWrapper('__cxa_can_catch', 3);
  Module['___cxa_get_exception_ptr'] = ___cxa_get_exception_ptr = createExportWrapper('__cxa_get_exception_ptr', 1);
  Module['__ZNSt9exceptionD0Ev'] = __ZNSt9exceptionD0Ev = createExportWrapper('_ZNSt9exceptionD0Ev', 1);
  Module['__ZNSt9exceptionD1Ev'] = __ZNSt9exceptionD1Ev = createExportWrapper('_ZNSt9exceptionD1Ev', 1);
  Module['__ZNKSt9exception4whatEv'] = __ZNKSt9exception4whatEv = createExportWrapper('_ZNKSt9exception4whatEv', 1);
  Module['__ZNSt13bad_exceptionD0Ev'] = __ZNSt13bad_exceptionD0Ev = createExportWrapper('_ZNSt13bad_exceptionD0Ev', 1);
  Module['__ZNSt13bad_exceptionD1Ev'] = __ZNSt13bad_exceptionD1Ev = createExportWrapper('_ZNSt13bad_exceptionD1Ev', 1);
  Module['__ZNKSt13bad_exception4whatEv'] = __ZNKSt13bad_exception4whatEv = createExportWrapper('_ZNKSt13bad_exception4whatEv', 1);
  Module['__ZNSt9bad_allocC2Ev'] = __ZNSt9bad_allocC2Ev = createExportWrapper('_ZNSt9bad_allocC2Ev', 1);
  Module['__ZNSt9bad_allocD0Ev'] = __ZNSt9bad_allocD0Ev = createExportWrapper('_ZNSt9bad_allocD0Ev', 1);
  Module['__ZNSt9bad_allocD1Ev'] = __ZNSt9bad_allocD1Ev = createExportWrapper('_ZNSt9bad_allocD1Ev', 1);
  Module['__ZNKSt9bad_alloc4whatEv'] = __ZNKSt9bad_alloc4whatEv = createExportWrapper('_ZNKSt9bad_alloc4whatEv', 1);
  Module['__ZNSt20bad_array_new_lengthC2Ev'] = __ZNSt20bad_array_new_lengthC2Ev = createExportWrapper('_ZNSt20bad_array_new_lengthC2Ev', 1);
  Module['__ZNSt20bad_array_new_lengthD0Ev'] = __ZNSt20bad_array_new_lengthD0Ev = createExportWrapper('_ZNSt20bad_array_new_lengthD0Ev', 1);
  Module['__ZNSt20bad_array_new_lengthD1Ev'] = __ZNSt20bad_array_new_lengthD1Ev = createExportWrapper('_ZNSt20bad_array_new_lengthD1Ev', 1);
  Module['__ZNKSt20bad_array_new_length4whatEv'] = __ZNKSt20bad_array_new_length4whatEv = createExportWrapper('_ZNKSt20bad_array_new_length4whatEv', 1);
  Module['__ZNSt13bad_exceptionD2Ev'] = __ZNSt13bad_exceptionD2Ev = createExportWrapper('_ZNSt13bad_exceptionD2Ev', 1);
  Module['__ZNSt9bad_allocC1Ev'] = __ZNSt9bad_allocC1Ev = createExportWrapper('_ZNSt9bad_allocC1Ev', 1);
  Module['__ZNSt9bad_allocD2Ev'] = __ZNSt9bad_allocD2Ev = createExportWrapper('_ZNSt9bad_allocD2Ev', 1);
  Module['__ZNSt20bad_array_new_lengthC1Ev'] = __ZNSt20bad_array_new_lengthC1Ev = createExportWrapper('_ZNSt20bad_array_new_lengthC1Ev', 1);
  Module['__ZNSt20bad_array_new_lengthD2Ev'] = __ZNSt20bad_array_new_lengthD2Ev = createExportWrapper('_ZNSt20bad_array_new_lengthD2Ev', 1);
  Module['__ZNSt11logic_errorD0Ev'] = __ZNSt11logic_errorD0Ev = createExportWrapper('_ZNSt11logic_errorD0Ev', 1);
  Module['__ZNSt11logic_errorD1Ev'] = __ZNSt11logic_errorD1Ev = createExportWrapper('_ZNSt11logic_errorD1Ev', 1);
  Module['__ZNSt13runtime_errorD0Ev'] = __ZNSt13runtime_errorD0Ev = createExportWrapper('_ZNSt13runtime_errorD0Ev', 1);
  Module['__ZNSt13runtime_errorD1Ev'] = __ZNSt13runtime_errorD1Ev = createExportWrapper('_ZNSt13runtime_errorD1Ev', 1);
  Module['__ZNSt12domain_errorD0Ev'] = __ZNSt12domain_errorD0Ev = createExportWrapper('_ZNSt12domain_errorD0Ev', 1);
  Module['__ZNSt12domain_errorD1Ev'] = __ZNSt12domain_errorD1Ev = createExportWrapper('_ZNSt12domain_errorD1Ev', 1);
  Module['__ZNSt16invalid_argumentD0Ev'] = __ZNSt16invalid_argumentD0Ev = createExportWrapper('_ZNSt16invalid_argumentD0Ev', 1);
  Module['__ZNSt16invalid_argumentD1Ev'] = __ZNSt16invalid_argumentD1Ev = createExportWrapper('_ZNSt16invalid_argumentD1Ev', 1);
  Module['__ZNSt12length_errorD0Ev'] = __ZNSt12length_errorD0Ev = createExportWrapper('_ZNSt12length_errorD0Ev', 1);
  Module['__ZNSt12length_errorD1Ev'] = __ZNSt12length_errorD1Ev = createExportWrapper('_ZNSt12length_errorD1Ev', 1);
  Module['__ZNSt12out_of_rangeD0Ev'] = __ZNSt12out_of_rangeD0Ev = createExportWrapper('_ZNSt12out_of_rangeD0Ev', 1);
  Module['__ZNSt12out_of_rangeD1Ev'] = __ZNSt12out_of_rangeD1Ev = createExportWrapper('_ZNSt12out_of_rangeD1Ev', 1);
  Module['__ZNSt11range_errorD0Ev'] = __ZNSt11range_errorD0Ev = createExportWrapper('_ZNSt11range_errorD0Ev', 1);
  Module['__ZNSt11range_errorD1Ev'] = __ZNSt11range_errorD1Ev = createExportWrapper('_ZNSt11range_errorD1Ev', 1);
  Module['__ZNSt14overflow_errorD0Ev'] = __ZNSt14overflow_errorD0Ev = createExportWrapper('_ZNSt14overflow_errorD0Ev', 1);
  Module['__ZNSt14overflow_errorD1Ev'] = __ZNSt14overflow_errorD1Ev = createExportWrapper('_ZNSt14overflow_errorD1Ev', 1);
  Module['__ZNSt15underflow_errorD0Ev'] = __ZNSt15underflow_errorD0Ev = createExportWrapper('_ZNSt15underflow_errorD0Ev', 1);
  Module['__ZNSt15underflow_errorD1Ev'] = __ZNSt15underflow_errorD1Ev = createExportWrapper('_ZNSt15underflow_errorD1Ev', 1);
  Module['__ZNSt12domain_errorD2Ev'] = __ZNSt12domain_errorD2Ev = createExportWrapper('_ZNSt12domain_errorD2Ev', 1);
  Module['__ZNSt16invalid_argumentD2Ev'] = __ZNSt16invalid_argumentD2Ev = createExportWrapper('_ZNSt16invalid_argumentD2Ev', 1);
  Module['__ZNSt12length_errorD2Ev'] = __ZNSt12length_errorD2Ev = createExportWrapper('_ZNSt12length_errorD2Ev', 1);
  Module['__ZNSt12out_of_rangeD2Ev'] = __ZNSt12out_of_rangeD2Ev = createExportWrapper('_ZNSt12out_of_rangeD2Ev', 1);
  Module['__ZNSt11range_errorD2Ev'] = __ZNSt11range_errorD2Ev = createExportWrapper('_ZNSt11range_errorD2Ev', 1);
  Module['__ZNSt14overflow_errorD2Ev'] = __ZNSt14overflow_errorD2Ev = createExportWrapper('_ZNSt14overflow_errorD2Ev', 1);
  Module['__ZNSt15underflow_errorD2Ev'] = __ZNSt15underflow_errorD2Ev = createExportWrapper('_ZNSt15underflow_errorD2Ev', 1);
  Module['__ZNSt9type_infoD0Ev'] = __ZNSt9type_infoD0Ev = createExportWrapper('_ZNSt9type_infoD0Ev', 1);
  Module['__ZNSt9type_infoD1Ev'] = __ZNSt9type_infoD1Ev = createExportWrapper('_ZNSt9type_infoD1Ev', 1);
  Module['__ZNSt8bad_castC2Ev'] = __ZNSt8bad_castC2Ev = createExportWrapper('_ZNSt8bad_castC2Ev', 1);
  Module['__ZNSt8bad_castD0Ev'] = __ZNSt8bad_castD0Ev = createExportWrapper('_ZNSt8bad_castD0Ev', 1);
  Module['__ZNSt8bad_castD1Ev'] = __ZNSt8bad_castD1Ev = createExportWrapper('_ZNSt8bad_castD1Ev', 1);
  Module['__ZNKSt8bad_cast4whatEv'] = __ZNKSt8bad_cast4whatEv = createExportWrapper('_ZNKSt8bad_cast4whatEv', 1);
  Module['__ZNSt10bad_typeidC2Ev'] = __ZNSt10bad_typeidC2Ev = createExportWrapper('_ZNSt10bad_typeidC2Ev', 1);
  Module['__ZNSt10bad_typeidD2Ev'] = __ZNSt10bad_typeidD2Ev = createExportWrapper('_ZNSt10bad_typeidD2Ev', 1);
  Module['__ZNSt10bad_typeidD0Ev'] = __ZNSt10bad_typeidD0Ev = createExportWrapper('_ZNSt10bad_typeidD0Ev', 1);
  Module['__ZNSt10bad_typeidD1Ev'] = __ZNSt10bad_typeidD1Ev = createExportWrapper('_ZNSt10bad_typeidD1Ev', 1);
  Module['__ZNKSt10bad_typeid4whatEv'] = __ZNKSt10bad_typeid4whatEv = createExportWrapper('_ZNKSt10bad_typeid4whatEv', 1);
  Module['__ZNSt8bad_castC1Ev'] = __ZNSt8bad_castC1Ev = createExportWrapper('_ZNSt8bad_castC1Ev', 1);
  Module['__ZNSt10bad_typeidC1Ev'] = __ZNSt10bad_typeidC1Ev = createExportWrapper('_ZNSt10bad_typeidC1Ev', 1);
}
var __ZTVN10__cxxabiv120__si_class_type_infoE = Module['__ZTVN10__cxxabiv120__si_class_type_infoE'] = 606672;
var __ZTISt8bad_cast = Module['__ZTISt8bad_cast'] = 607688;
var __ZTISt13runtime_error = Module['__ZTISt13runtime_error'] = 607472;
var __ZTVN10__cxxabiv117__class_type_infoE = Module['__ZTVN10__cxxabiv117__class_type_infoE'] = 606632;
var __ZTISt9exception = Module['__ZTISt9exception'] = 607000;
var __ZTISt11logic_error = Module['__ZTISt11logic_error'] = 607236;
var __ZTVN10__cxxabiv121__vmi_class_type_infoE = Module['__ZTVN10__cxxabiv121__vmi_class_type_infoE'] = 606764;
var __ZTVSt11logic_error = Module['__ZTVSt11logic_error'] = 607144;
var __ZTVSt9exception = Module['__ZTVSt9exception'] = 606980;
var __ZTVSt13runtime_error = Module['__ZTVSt13runtime_error'] = 607164;
var ___cxa_unexpected_handler = Module['___cxa_unexpected_handler'] = 622564;
var ___cxa_terminate_handler = Module['___cxa_terminate_handler'] = 622560;
var ___cxa_new_handler = Module['___cxa_new_handler'] = 680028;
var __ZTIN10__cxxabiv116__shim_type_infoE = Module['__ZTIN10__cxxabiv116__shim_type_infoE'] = 604708;
var __ZTIN10__cxxabiv117__class_type_infoE = Module['__ZTIN10__cxxabiv117__class_type_infoE'] = 604756;
var __ZTIN10__cxxabiv117__pbase_type_infoE = Module['__ZTIN10__cxxabiv117__pbase_type_infoE'] = 604804;
var __ZTIDn = Module['__ZTIDn'] = 605184;
var __ZTIN10__cxxabiv119__pointer_type_infoE = Module['__ZTIN10__cxxabiv119__pointer_type_infoE'] = 604852;
var __ZTIv = Module['__ZTIv'] = 605132;
var __ZTIN10__cxxabiv120__function_type_infoE = Module['__ZTIN10__cxxabiv120__function_type_infoE'] = 604900;
var __ZTIN10__cxxabiv129__pointer_to_member_type_infoE = Module['__ZTIN10__cxxabiv129__pointer_to_member_type_infoE'] = 604952;
var __ZTISt9type_info = Module['__ZTISt9type_info'] = 607664;
var __ZTSN10__cxxabiv116__shim_type_infoE = Module['__ZTSN10__cxxabiv116__shim_type_infoE'] = 604720;
var __ZTSN10__cxxabiv117__class_type_infoE = Module['__ZTSN10__cxxabiv117__class_type_infoE'] = 604768;
var __ZTSN10__cxxabiv117__pbase_type_infoE = Module['__ZTSN10__cxxabiv117__pbase_type_infoE'] = 604816;
var __ZTSN10__cxxabiv119__pointer_type_infoE = Module['__ZTSN10__cxxabiv119__pointer_type_infoE'] = 604864;
var __ZTSN10__cxxabiv120__function_type_infoE = Module['__ZTSN10__cxxabiv120__function_type_infoE'] = 604912;
var __ZTSN10__cxxabiv129__pointer_to_member_type_infoE = Module['__ZTSN10__cxxabiv129__pointer_to_member_type_infoE'] = 604964;
var __ZTVN10__cxxabiv116__shim_type_infoE = Module['__ZTVN10__cxxabiv116__shim_type_infoE'] = 605024;
var __ZTVN10__cxxabiv123__fundamental_type_infoE = Module['__ZTVN10__cxxabiv123__fundamental_type_infoE'] = 605052;
var __ZTIN10__cxxabiv123__fundamental_type_infoE = Module['__ZTIN10__cxxabiv123__fundamental_type_infoE'] = 605080;
var __ZTSN10__cxxabiv123__fundamental_type_infoE = Module['__ZTSN10__cxxabiv123__fundamental_type_infoE'] = 605092;
var __ZTSv = Module['__ZTSv'] = 605140;
var __ZTIPv = Module['__ZTIPv'] = 605144;
var __ZTVN10__cxxabiv119__pointer_type_infoE = Module['__ZTVN10__cxxabiv119__pointer_type_infoE'] = 606884;
var __ZTSPv = Module['__ZTSPv'] = 605160;
var __ZTIPKv = Module['__ZTIPKv'] = 605164;
var __ZTSPKv = Module['__ZTSPKv'] = 605180;
var __ZTSDn = Module['__ZTSDn'] = 605192;
var __ZTIPDn = Module['__ZTIPDn'] = 605196;
var __ZTSPDn = Module['__ZTSPDn'] = 605212;
var __ZTIPKDn = Module['__ZTIPKDn'] = 605216;
var __ZTSPKDn = Module['__ZTSPKDn'] = 605232;
var __ZTIb = Module['__ZTIb'] = 605240;
var __ZTSb = Module['__ZTSb'] = 605248;
var __ZTIPb = Module['__ZTIPb'] = 605252;
var __ZTSPb = Module['__ZTSPb'] = 605268;
var __ZTIPKb = Module['__ZTIPKb'] = 605272;
var __ZTSPKb = Module['__ZTSPKb'] = 605288;
var __ZTIw = Module['__ZTIw'] = 605292;
var __ZTSw = Module['__ZTSw'] = 605300;
var __ZTIPw = Module['__ZTIPw'] = 605304;
var __ZTSPw = Module['__ZTSPw'] = 605320;
var __ZTIPKw = Module['__ZTIPKw'] = 605324;
var __ZTSPKw = Module['__ZTSPKw'] = 605340;
var __ZTIc = Module['__ZTIc'] = 605344;
var __ZTSc = Module['__ZTSc'] = 605352;
var __ZTIPc = Module['__ZTIPc'] = 605356;
var __ZTSPc = Module['__ZTSPc'] = 605372;
var __ZTIPKc = Module['__ZTIPKc'] = 605376;
var __ZTSPKc = Module['__ZTSPKc'] = 605392;
var __ZTIh = Module['__ZTIh'] = 605396;
var __ZTSh = Module['__ZTSh'] = 605404;
var __ZTIPh = Module['__ZTIPh'] = 605408;
var __ZTSPh = Module['__ZTSPh'] = 605424;
var __ZTIPKh = Module['__ZTIPKh'] = 605428;
var __ZTSPKh = Module['__ZTSPKh'] = 605444;
var __ZTIa = Module['__ZTIa'] = 605448;
var __ZTSa = Module['__ZTSa'] = 605456;
var __ZTIPa = Module['__ZTIPa'] = 605460;
var __ZTSPa = Module['__ZTSPa'] = 605476;
var __ZTIPKa = Module['__ZTIPKa'] = 605480;
var __ZTSPKa = Module['__ZTSPKa'] = 605496;
var __ZTIs = Module['__ZTIs'] = 605500;
var __ZTSs = Module['__ZTSs'] = 605508;
var __ZTIPs = Module['__ZTIPs'] = 605512;
var __ZTSPs = Module['__ZTSPs'] = 605528;
var __ZTIPKs = Module['__ZTIPKs'] = 605532;
var __ZTSPKs = Module['__ZTSPKs'] = 605548;
var __ZTIt = Module['__ZTIt'] = 605552;
var __ZTSt = Module['__ZTSt'] = 605560;
var __ZTIPt = Module['__ZTIPt'] = 605564;
var __ZTSPt = Module['__ZTSPt'] = 605580;
var __ZTIPKt = Module['__ZTIPKt'] = 605584;
var __ZTSPKt = Module['__ZTSPKt'] = 605600;
var __ZTIi = Module['__ZTIi'] = 605604;
var __ZTSi = Module['__ZTSi'] = 605612;
var __ZTIPi = Module['__ZTIPi'] = 605616;
var __ZTSPi = Module['__ZTSPi'] = 605632;
var __ZTIPKi = Module['__ZTIPKi'] = 605636;
var __ZTSPKi = Module['__ZTSPKi'] = 605652;
var __ZTIj = Module['__ZTIj'] = 605656;
var __ZTSj = Module['__ZTSj'] = 605664;
var __ZTIPj = Module['__ZTIPj'] = 605668;
var __ZTSPj = Module['__ZTSPj'] = 605684;
var __ZTIPKj = Module['__ZTIPKj'] = 605688;
var __ZTSPKj = Module['__ZTSPKj'] = 605704;
var __ZTIl = Module['__ZTIl'] = 605708;
var __ZTSl = Module['__ZTSl'] = 605716;
var __ZTIPl = Module['__ZTIPl'] = 605720;
var __ZTSPl = Module['__ZTSPl'] = 605736;
var __ZTIPKl = Module['__ZTIPKl'] = 605740;
var __ZTSPKl = Module['__ZTSPKl'] = 605756;
var __ZTIm = Module['__ZTIm'] = 605760;
var __ZTSm = Module['__ZTSm'] = 605768;
var __ZTIPm = Module['__ZTIPm'] = 605772;
var __ZTSPm = Module['__ZTSPm'] = 605788;
var __ZTIPKm = Module['__ZTIPKm'] = 605792;
var __ZTSPKm = Module['__ZTSPKm'] = 605808;
var __ZTIx = Module['__ZTIx'] = 605812;
var __ZTSx = Module['__ZTSx'] = 605820;
var __ZTIPx = Module['__ZTIPx'] = 605824;
var __ZTSPx = Module['__ZTSPx'] = 605840;
var __ZTIPKx = Module['__ZTIPKx'] = 605844;
var __ZTSPKx = Module['__ZTSPKx'] = 605860;
var __ZTIy = Module['__ZTIy'] = 605864;
var __ZTSy = Module['__ZTSy'] = 605872;
var __ZTIPy = Module['__ZTIPy'] = 605876;
var __ZTSPy = Module['__ZTSPy'] = 605892;
var __ZTIPKy = Module['__ZTIPKy'] = 605896;
var __ZTSPKy = Module['__ZTSPKy'] = 605912;
var __ZTIn = Module['__ZTIn'] = 605916;
var __ZTSn = Module['__ZTSn'] = 605924;
var __ZTIPn = Module['__ZTIPn'] = 605928;
var __ZTSPn = Module['__ZTSPn'] = 605944;
var __ZTIPKn = Module['__ZTIPKn'] = 605948;
var __ZTSPKn = Module['__ZTSPKn'] = 605964;
var __ZTIo = Module['__ZTIo'] = 605968;
var __ZTSo = Module['__ZTSo'] = 605976;
var __ZTIPo = Module['__ZTIPo'] = 605980;
var __ZTSPo = Module['__ZTSPo'] = 605996;
var __ZTIPKo = Module['__ZTIPKo'] = 606000;
var __ZTSPKo = Module['__ZTSPKo'] = 606016;
var __ZTIDh = Module['__ZTIDh'] = 606020;
var __ZTSDh = Module['__ZTSDh'] = 606028;
var __ZTIPDh = Module['__ZTIPDh'] = 606032;
var __ZTSPDh = Module['__ZTSPDh'] = 606048;
var __ZTIPKDh = Module['__ZTIPKDh'] = 606052;
var __ZTSPKDh = Module['__ZTSPKDh'] = 606068;
var __ZTIf = Module['__ZTIf'] = 606076;
var __ZTSf = Module['__ZTSf'] = 606084;
var __ZTIPf = Module['__ZTIPf'] = 606088;
var __ZTSPf = Module['__ZTSPf'] = 606104;
var __ZTIPKf = Module['__ZTIPKf'] = 606108;
var __ZTSPKf = Module['__ZTSPKf'] = 606124;
var __ZTId = Module['__ZTId'] = 606128;
var __ZTSd = Module['__ZTSd'] = 606136;
var __ZTIPd = Module['__ZTIPd'] = 606140;
var __ZTSPd = Module['__ZTSPd'] = 606156;
var __ZTIPKd = Module['__ZTIPKd'] = 606160;
var __ZTSPKd = Module['__ZTSPKd'] = 606176;
var __ZTIe = Module['__ZTIe'] = 606180;
var __ZTSe = Module['__ZTSe'] = 606188;
var __ZTIPe = Module['__ZTIPe'] = 606192;
var __ZTSPe = Module['__ZTSPe'] = 606208;
var __ZTIPKe = Module['__ZTIPKe'] = 606212;
var __ZTSPKe = Module['__ZTSPKe'] = 606228;
var __ZTIg = Module['__ZTIg'] = 606232;
var __ZTSg = Module['__ZTSg'] = 606240;
var __ZTIPg = Module['__ZTIPg'] = 606244;
var __ZTSPg = Module['__ZTSPg'] = 606260;
var __ZTIPKg = Module['__ZTIPKg'] = 606264;
var __ZTSPKg = Module['__ZTSPKg'] = 606280;
var __ZTIDu = Module['__ZTIDu'] = 606284;
var __ZTSDu = Module['__ZTSDu'] = 606292;
var __ZTIPDu = Module['__ZTIPDu'] = 606296;
var __ZTSPDu = Module['__ZTSPDu'] = 606312;
var __ZTIPKDu = Module['__ZTIPKDu'] = 606316;
var __ZTSPKDu = Module['__ZTSPKDu'] = 606332;
var __ZTIDs = Module['__ZTIDs'] = 606340;
var __ZTSDs = Module['__ZTSDs'] = 606348;
var __ZTIPDs = Module['__ZTIPDs'] = 606352;
var __ZTSPDs = Module['__ZTSPDs'] = 606368;
var __ZTIPKDs = Module['__ZTIPKDs'] = 606372;
var __ZTSPKDs = Module['__ZTSPKDs'] = 606388;
var __ZTIDi = Module['__ZTIDi'] = 606396;
var __ZTSDi = Module['__ZTSDi'] = 606404;
var __ZTIPDi = Module['__ZTIPDi'] = 606408;
var __ZTSPDi = Module['__ZTSPDi'] = 606424;
var __ZTIPKDi = Module['__ZTIPKDi'] = 606428;
var __ZTSPKDi = Module['__ZTSPKDi'] = 606444;
var __ZTVN10__cxxabiv117__array_type_infoE = Module['__ZTVN10__cxxabiv117__array_type_infoE'] = 606452;
var __ZTIN10__cxxabiv117__array_type_infoE = Module['__ZTIN10__cxxabiv117__array_type_infoE'] = 606480;
var __ZTSN10__cxxabiv117__array_type_infoE = Module['__ZTSN10__cxxabiv117__array_type_infoE'] = 606492;
var __ZTVN10__cxxabiv120__function_type_infoE = Module['__ZTVN10__cxxabiv120__function_type_infoE'] = 606528;
var __ZTVN10__cxxabiv116__enum_type_infoE = Module['__ZTVN10__cxxabiv116__enum_type_infoE'] = 606556;
var __ZTIN10__cxxabiv116__enum_type_infoE = Module['__ZTIN10__cxxabiv116__enum_type_infoE'] = 606584;
var __ZTSN10__cxxabiv116__enum_type_infoE = Module['__ZTSN10__cxxabiv116__enum_type_infoE'] = 606596;
var __ZTIN10__cxxabiv120__si_class_type_infoE = Module['__ZTIN10__cxxabiv120__si_class_type_infoE'] = 606712;
var __ZTSN10__cxxabiv120__si_class_type_infoE = Module['__ZTSN10__cxxabiv120__si_class_type_infoE'] = 606724;
var __ZTIN10__cxxabiv121__vmi_class_type_infoE = Module['__ZTIN10__cxxabiv121__vmi_class_type_infoE'] = 606804;
var __ZTSN10__cxxabiv121__vmi_class_type_infoE = Module['__ZTSN10__cxxabiv121__vmi_class_type_infoE'] = 606816;
var __ZTVN10__cxxabiv117__pbase_type_infoE = Module['__ZTVN10__cxxabiv117__pbase_type_infoE'] = 606856;
var __ZTVN10__cxxabiv129__pointer_to_member_type_infoE = Module['__ZTVN10__cxxabiv129__pointer_to_member_type_infoE'] = 606912;
var __ZTVSt9bad_alloc = Module['__ZTVSt9bad_alloc'] = 606940;
var __ZTVSt20bad_array_new_length = Module['__ZTVSt20bad_array_new_length'] = 606960;
var __ZTISt9bad_alloc = Module['__ZTISt9bad_alloc'] = 607076;
var __ZTISt20bad_array_new_length = Module['__ZTISt20bad_array_new_length'] = 607104;
var __ZTSSt9exception = Module['__ZTSSt9exception'] = 607008;
var __ZTVSt13bad_exception = Module['__ZTVSt13bad_exception'] = 607024;
var __ZTISt13bad_exception = Module['__ZTISt13bad_exception'] = 607044;
var __ZTSSt13bad_exception = Module['__ZTSSt13bad_exception'] = 607056;
var __ZTSSt9bad_alloc = Module['__ZTSSt9bad_alloc'] = 607088;
var __ZTSSt20bad_array_new_length = Module['__ZTSSt20bad_array_new_length'] = 607116;
var __ZTVSt12domain_error = Module['__ZTVSt12domain_error'] = 607184;
var __ZTISt12domain_error = Module['__ZTISt12domain_error'] = 607204;
var __ZTSSt12domain_error = Module['__ZTSSt12domain_error'] = 607216;
var __ZTSSt11logic_error = Module['__ZTSSt11logic_error'] = 607248;
var __ZTVSt16invalid_argument = Module['__ZTVSt16invalid_argument'] = 607264;
var __ZTISt16invalid_argument = Module['__ZTISt16invalid_argument'] = 607284;
var __ZTSSt16invalid_argument = Module['__ZTSSt16invalid_argument'] = 607296;
var __ZTVSt12length_error = Module['__ZTVSt12length_error'] = 607320;
var __ZTISt12length_error = Module['__ZTISt12length_error'] = 607340;
var __ZTSSt12length_error = Module['__ZTSSt12length_error'] = 607352;
var __ZTVSt12out_of_range = Module['__ZTVSt12out_of_range'] = 607372;
var __ZTISt12out_of_range = Module['__ZTISt12out_of_range'] = 607392;
var __ZTSSt12out_of_range = Module['__ZTSSt12out_of_range'] = 607404;
var __ZTVSt11range_error = Module['__ZTVSt11range_error'] = 607424;
var __ZTISt11range_error = Module['__ZTISt11range_error'] = 607444;
var __ZTSSt11range_error = Module['__ZTSSt11range_error'] = 607456;
var __ZTSSt13runtime_error = Module['__ZTSSt13runtime_error'] = 607484;
var __ZTVSt14overflow_error = Module['__ZTVSt14overflow_error'] = 607504;
var __ZTISt14overflow_error = Module['__ZTISt14overflow_error'] = 607524;
var __ZTSSt14overflow_error = Module['__ZTSSt14overflow_error'] = 607536;
var __ZTVSt15underflow_error = Module['__ZTVSt15underflow_error'] = 607556;
var __ZTISt15underflow_error = Module['__ZTISt15underflow_error'] = 607576;
var __ZTSSt15underflow_error = Module['__ZTSSt15underflow_error'] = 607588;
var __ZTVSt8bad_cast = Module['__ZTVSt8bad_cast'] = 607608;
var __ZTVSt10bad_typeid = Module['__ZTVSt10bad_typeid'] = 607628;
var __ZTISt10bad_typeid = Module['__ZTISt10bad_typeid'] = 607712;
var __ZTVSt9type_info = Module['__ZTVSt9type_info'] = 607648;
var __ZTSSt9type_info = Module['__ZTSSt9type_info'] = 607672;
var __ZTSSt8bad_cast = Module['__ZTSSt8bad_cast'] = 607700;
var __ZTSSt10bad_typeid = Module['__ZTSSt10bad_typeid'] = 607724;var wasmImports = {
  /** @export */
  __call_sighandler: ___call_sighandler,
  /** @export */
  __syscall_fcntl64: ___syscall_fcntl64,
  /** @export */
  __syscall_fdatasync: ___syscall_fdatasync,
  /** @export */
  __syscall_fstat64: ___syscall_fstat64,
  /** @export */
  __syscall_ftruncate64: ___syscall_ftruncate64,
  /** @export */
  __syscall_getcwd: ___syscall_getcwd,
  /** @export */
  __syscall_getdents64: ___syscall_getdents64,
  /** @export */
  __syscall_ioctl: ___syscall_ioctl,
  /** @export */
  __syscall_lstat64: ___syscall_lstat64,
  /** @export */
  __syscall_mkdirat: ___syscall_mkdirat,
  /** @export */
  __syscall_newfstatat: ___syscall_newfstatat,
  /** @export */
  __syscall_openat: ___syscall_openat,
  /** @export */
  __syscall_renameat: ___syscall_renameat,
  /** @export */
  __syscall_rmdir: ___syscall_rmdir,
  /** @export */
  __syscall_stat64: ___syscall_stat64,
  /** @export */
  __syscall_unlinkat: ___syscall_unlinkat,
  /** @export */
  _abort_js: __abort_js,
  /** @export */
  _emscripten_get_progname: __emscripten_get_progname,
  /** @export */
  _emscripten_runtime_keepalive_clear: __emscripten_runtime_keepalive_clear,
  /** @export */
  _emscripten_throw_longjmp: __emscripten_throw_longjmp,
  /** @export */
  _gmtime_js: __gmtime_js,
  /** @export */
  _localtime_js: __localtime_js,
  /** @export */
  _mmap_js: __mmap_js,
  /** @export */
  _munmap_js: __munmap_js,
  /** @export */
  _setitimer_js: __setitimer_js,
  /** @export */
  _tzset_js: __tzset_js,
  /** @export */
  clock_time_get: _clock_time_get,
  /** @export */
  emscripten_asm_const_int: _emscripten_asm_const_int,
  /** @export */
  emscripten_asm_const_int_sync_on_main_thread: _emscripten_asm_const_int_sync_on_main_thread,
  /** @export */
  emscripten_asm_const_ptr_sync_on_main_thread: _emscripten_asm_const_ptr_sync_on_main_thread,
  /** @export */
  emscripten_cancel_main_loop: _emscripten_cancel_main_loop,
  /** @export */
  emscripten_clear_timeout: _emscripten_clear_timeout,
  /** @export */
  emscripten_date_now: _emscripten_date_now,
  /** @export */
  emscripten_err: _emscripten_err,
  /** @export */
  emscripten_exit_fullscreen: _emscripten_exit_fullscreen,
  /** @export */
  emscripten_exit_pointerlock: _emscripten_exit_pointerlock,
  /** @export */
  emscripten_force_exit: _emscripten_force_exit,
  /** @export */
  emscripten_get_battery_status: _emscripten_get_battery_status,
  /** @export */
  emscripten_get_device_pixel_ratio: _emscripten_get_device_pixel_ratio,
  /** @export */
  emscripten_get_element_css_size: _emscripten_get_element_css_size,
  /** @export */
  emscripten_get_gamepad_status: _emscripten_get_gamepad_status,
  /** @export */
  emscripten_get_heap_max: _emscripten_get_heap_max,
  /** @export */
  emscripten_get_main_loop_timing: _emscripten_get_main_loop_timing,
  /** @export */
  emscripten_get_now: _emscripten_get_now,
  /** @export */
  emscripten_get_num_gamepads: _emscripten_get_num_gamepads,
  /** @export */
  emscripten_get_preloaded_image_data: _emscripten_get_preloaded_image_data,
  /** @export */
  emscripten_get_preloaded_image_data_from_FILE: _emscripten_get_preloaded_image_data_from_FILE,
  /** @export */
  emscripten_get_screen_size: _emscripten_get_screen_size,
  /** @export */
  emscripten_glActiveTexture: _emscripten_glActiveTexture,
  /** @export */
  emscripten_glAttachShader: _emscripten_glAttachShader,
  /** @export */
  emscripten_glBeginQueryEXT: _emscripten_glBeginQueryEXT,
  /** @export */
  emscripten_glBindAttribLocation: _emscripten_glBindAttribLocation,
  /** @export */
  emscripten_glBindBuffer: _emscripten_glBindBuffer,
  /** @export */
  emscripten_glBindFramebuffer: _emscripten_glBindFramebuffer,
  /** @export */
  emscripten_glBindRenderbuffer: _emscripten_glBindRenderbuffer,
  /** @export */
  emscripten_glBindTexture: _emscripten_glBindTexture,
  /** @export */
  emscripten_glBindVertexArrayOES: _emscripten_glBindVertexArrayOES,
  /** @export */
  emscripten_glBlendColor: _emscripten_glBlendColor,
  /** @export */
  emscripten_glBlendEquation: _emscripten_glBlendEquation,
  /** @export */
  emscripten_glBlendEquationSeparate: _emscripten_glBlendEquationSeparate,
  /** @export */
  emscripten_glBlendFunc: _emscripten_glBlendFunc,
  /** @export */
  emscripten_glBlendFuncSeparate: _emscripten_glBlendFuncSeparate,
  /** @export */
  emscripten_glBufferData: _emscripten_glBufferData,
  /** @export */
  emscripten_glBufferSubData: _emscripten_glBufferSubData,
  /** @export */
  emscripten_glCheckFramebufferStatus: _emscripten_glCheckFramebufferStatus,
  /** @export */
  emscripten_glClear: _emscripten_glClear,
  /** @export */
  emscripten_glClearColor: _emscripten_glClearColor,
  /** @export */
  emscripten_glClearDepthf: _emscripten_glClearDepthf,
  /** @export */
  emscripten_glClearStencil: _emscripten_glClearStencil,
  /** @export */
  emscripten_glClipControlEXT: _emscripten_glClipControlEXT,
  /** @export */
  emscripten_glColorMask: _emscripten_glColorMask,
  /** @export */
  emscripten_glCompileShader: _emscripten_glCompileShader,
  /** @export */
  emscripten_glCompressedTexImage2D: _emscripten_glCompressedTexImage2D,
  /** @export */
  emscripten_glCompressedTexSubImage2D: _emscripten_glCompressedTexSubImage2D,
  /** @export */
  emscripten_glCopyTexImage2D: _emscripten_glCopyTexImage2D,
  /** @export */
  emscripten_glCopyTexSubImage2D: _emscripten_glCopyTexSubImage2D,
  /** @export */
  emscripten_glCreateProgram: _emscripten_glCreateProgram,
  /** @export */
  emscripten_glCreateShader: _emscripten_glCreateShader,
  /** @export */
  emscripten_glCullFace: _emscripten_glCullFace,
  /** @export */
  emscripten_glDeleteBuffers: _emscripten_glDeleteBuffers,
  /** @export */
  emscripten_glDeleteFramebuffers: _emscripten_glDeleteFramebuffers,
  /** @export */
  emscripten_glDeleteProgram: _emscripten_glDeleteProgram,
  /** @export */
  emscripten_glDeleteQueriesEXT: _emscripten_glDeleteQueriesEXT,
  /** @export */
  emscripten_glDeleteRenderbuffers: _emscripten_glDeleteRenderbuffers,
  /** @export */
  emscripten_glDeleteShader: _emscripten_glDeleteShader,
  /** @export */
  emscripten_glDeleteTextures: _emscripten_glDeleteTextures,
  /** @export */
  emscripten_glDeleteVertexArraysOES: _emscripten_glDeleteVertexArraysOES,
  /** @export */
  emscripten_glDepthFunc: _emscripten_glDepthFunc,
  /** @export */
  emscripten_glDepthMask: _emscripten_glDepthMask,
  /** @export */
  emscripten_glDepthRangef: _emscripten_glDepthRangef,
  /** @export */
  emscripten_glDetachShader: _emscripten_glDetachShader,
  /** @export */
  emscripten_glDisable: _emscripten_glDisable,
  /** @export */
  emscripten_glDisableVertexAttribArray: _emscripten_glDisableVertexAttribArray,
  /** @export */
  emscripten_glDrawArrays: _emscripten_glDrawArrays,
  /** @export */
  emscripten_glDrawArraysInstancedANGLE: _emscripten_glDrawArraysInstancedANGLE,
  /** @export */
  emscripten_glDrawBuffersWEBGL: _emscripten_glDrawBuffersWEBGL,
  /** @export */
  emscripten_glDrawElements: _emscripten_glDrawElements,
  /** @export */
  emscripten_glDrawElementsInstancedANGLE: _emscripten_glDrawElementsInstancedANGLE,
  /** @export */
  emscripten_glEnable: _emscripten_glEnable,
  /** @export */
  emscripten_glEnableVertexAttribArray: _emscripten_glEnableVertexAttribArray,
  /** @export */
  emscripten_glEndQueryEXT: _emscripten_glEndQueryEXT,
  /** @export */
  emscripten_glFinish: _emscripten_glFinish,
  /** @export */
  emscripten_glFlush: _emscripten_glFlush,
  /** @export */
  emscripten_glFramebufferRenderbuffer: _emscripten_glFramebufferRenderbuffer,
  /** @export */
  emscripten_glFramebufferTexture2D: _emscripten_glFramebufferTexture2D,
  /** @export */
  emscripten_glFrontFace: _emscripten_glFrontFace,
  /** @export */
  emscripten_glGenBuffers: _emscripten_glGenBuffers,
  /** @export */
  emscripten_glGenFramebuffers: _emscripten_glGenFramebuffers,
  /** @export */
  emscripten_glGenQueriesEXT: _emscripten_glGenQueriesEXT,
  /** @export */
  emscripten_glGenRenderbuffers: _emscripten_glGenRenderbuffers,
  /** @export */
  emscripten_glGenTextures: _emscripten_glGenTextures,
  /** @export */
  emscripten_glGenVertexArraysOES: _emscripten_glGenVertexArraysOES,
  /** @export */
  emscripten_glGenerateMipmap: _emscripten_glGenerateMipmap,
  /** @export */
  emscripten_glGetActiveAttrib: _emscripten_glGetActiveAttrib,
  /** @export */
  emscripten_glGetActiveUniform: _emscripten_glGetActiveUniform,
  /** @export */
  emscripten_glGetAttachedShaders: _emscripten_glGetAttachedShaders,
  /** @export */
  emscripten_glGetAttribLocation: _emscripten_glGetAttribLocation,
  /** @export */
  emscripten_glGetBooleanv: _emscripten_glGetBooleanv,
  /** @export */
  emscripten_glGetBufferParameteriv: _emscripten_glGetBufferParameteriv,
  /** @export */
  emscripten_glGetError: _emscripten_glGetError,
  /** @export */
  emscripten_glGetFloatv: _emscripten_glGetFloatv,
  /** @export */
  emscripten_glGetFramebufferAttachmentParameteriv: _emscripten_glGetFramebufferAttachmentParameteriv,
  /** @export */
  emscripten_glGetIntegerv: _emscripten_glGetIntegerv,
  /** @export */
  emscripten_glGetProgramInfoLog: _emscripten_glGetProgramInfoLog,
  /** @export */
  emscripten_glGetProgramiv: _emscripten_glGetProgramiv,
  /** @export */
  emscripten_glGetQueryObjecti64vEXT: _emscripten_glGetQueryObjecti64vEXT,
  /** @export */
  emscripten_glGetQueryObjectivEXT: _emscripten_glGetQueryObjectivEXT,
  /** @export */
  emscripten_glGetQueryObjectui64vEXT: _emscripten_glGetQueryObjectui64vEXT,
  /** @export */
  emscripten_glGetQueryObjectuivEXT: _emscripten_glGetQueryObjectuivEXT,
  /** @export */
  emscripten_glGetQueryivEXT: _emscripten_glGetQueryivEXT,
  /** @export */
  emscripten_glGetRenderbufferParameteriv: _emscripten_glGetRenderbufferParameteriv,
  /** @export */
  emscripten_glGetShaderInfoLog: _emscripten_glGetShaderInfoLog,
  /** @export */
  emscripten_glGetShaderPrecisionFormat: _emscripten_glGetShaderPrecisionFormat,
  /** @export */
  emscripten_glGetShaderSource: _emscripten_glGetShaderSource,
  /** @export */
  emscripten_glGetShaderiv: _emscripten_glGetShaderiv,
  /** @export */
  emscripten_glGetString: _emscripten_glGetString,
  /** @export */
  emscripten_glGetTexParameterfv: _emscripten_glGetTexParameterfv,
  /** @export */
  emscripten_glGetTexParameteriv: _emscripten_glGetTexParameteriv,
  /** @export */
  emscripten_glGetUniformLocation: _emscripten_glGetUniformLocation,
  /** @export */
  emscripten_glGetUniformfv: _emscripten_glGetUniformfv,
  /** @export */
  emscripten_glGetUniformiv: _emscripten_glGetUniformiv,
  /** @export */
  emscripten_glGetVertexAttribPointerv: _emscripten_glGetVertexAttribPointerv,
  /** @export */
  emscripten_glGetVertexAttribfv: _emscripten_glGetVertexAttribfv,
  /** @export */
  emscripten_glGetVertexAttribiv: _emscripten_glGetVertexAttribiv,
  /** @export */
  emscripten_glHint: _emscripten_glHint,
  /** @export */
  emscripten_glIsBuffer: _emscripten_glIsBuffer,
  /** @export */
  emscripten_glIsEnabled: _emscripten_glIsEnabled,
  /** @export */
  emscripten_glIsFramebuffer: _emscripten_glIsFramebuffer,
  /** @export */
  emscripten_glIsProgram: _emscripten_glIsProgram,
  /** @export */
  emscripten_glIsQueryEXT: _emscripten_glIsQueryEXT,
  /** @export */
  emscripten_glIsRenderbuffer: _emscripten_glIsRenderbuffer,
  /** @export */
  emscripten_glIsShader: _emscripten_glIsShader,
  /** @export */
  emscripten_glIsTexture: _emscripten_glIsTexture,
  /** @export */
  emscripten_glIsVertexArrayOES: _emscripten_glIsVertexArrayOES,
  /** @export */
  emscripten_glLineWidth: _emscripten_glLineWidth,
  /** @export */
  emscripten_glLinkProgram: _emscripten_glLinkProgram,
  /** @export */
  emscripten_glPixelStorei: _emscripten_glPixelStorei,
  /** @export */
  emscripten_glPolygonModeWEBGL: _emscripten_glPolygonModeWEBGL,
  /** @export */
  emscripten_glPolygonOffset: _emscripten_glPolygonOffset,
  /** @export */
  emscripten_glPolygonOffsetClampEXT: _emscripten_glPolygonOffsetClampEXT,
  /** @export */
  emscripten_glQueryCounterEXT: _emscripten_glQueryCounterEXT,
  /** @export */
  emscripten_glReadPixels: _emscripten_glReadPixels,
  /** @export */
  emscripten_glReleaseShaderCompiler: _emscripten_glReleaseShaderCompiler,
  /** @export */
  emscripten_glRenderbufferStorage: _emscripten_glRenderbufferStorage,
  /** @export */
  emscripten_glSampleCoverage: _emscripten_glSampleCoverage,
  /** @export */
  emscripten_glScissor: _emscripten_glScissor,
  /** @export */
  emscripten_glShaderBinary: _emscripten_glShaderBinary,
  /** @export */
  emscripten_glShaderSource: _emscripten_glShaderSource,
  /** @export */
  emscripten_glStencilFunc: _emscripten_glStencilFunc,
  /** @export */
  emscripten_glStencilFuncSeparate: _emscripten_glStencilFuncSeparate,
  /** @export */
  emscripten_glStencilMask: _emscripten_glStencilMask,
  /** @export */
  emscripten_glStencilMaskSeparate: _emscripten_glStencilMaskSeparate,
  /** @export */
  emscripten_glStencilOp: _emscripten_glStencilOp,
  /** @export */
  emscripten_glStencilOpSeparate: _emscripten_glStencilOpSeparate,
  /** @export */
  emscripten_glTexImage2D: _emscripten_glTexImage2D,
  /** @export */
  emscripten_glTexParameterf: _emscripten_glTexParameterf,
  /** @export */
  emscripten_glTexParameterfv: _emscripten_glTexParameterfv,
  /** @export */
  emscripten_glTexParameteri: _emscripten_glTexParameteri,
  /** @export */
  emscripten_glTexParameteriv: _emscripten_glTexParameteriv,
  /** @export */
  emscripten_glTexSubImage2D: _emscripten_glTexSubImage2D,
  /** @export */
  emscripten_glUniform1f: _emscripten_glUniform1f,
  /** @export */
  emscripten_glUniform1fv: _emscripten_glUniform1fv,
  /** @export */
  emscripten_glUniform1i: _emscripten_glUniform1i,
  /** @export */
  emscripten_glUniform1iv: _emscripten_glUniform1iv,
  /** @export */
  emscripten_glUniform2f: _emscripten_glUniform2f,
  /** @export */
  emscripten_glUniform2fv: _emscripten_glUniform2fv,
  /** @export */
  emscripten_glUniform2i: _emscripten_glUniform2i,
  /** @export */
  emscripten_glUniform2iv: _emscripten_glUniform2iv,
  /** @export */
  emscripten_glUniform3f: _emscripten_glUniform3f,
  /** @export */
  emscripten_glUniform3fv: _emscripten_glUniform3fv,
  /** @export */
  emscripten_glUniform3i: _emscripten_glUniform3i,
  /** @export */
  emscripten_glUniform3iv: _emscripten_glUniform3iv,
  /** @export */
  emscripten_glUniform4f: _emscripten_glUniform4f,
  /** @export */
  emscripten_glUniform4fv: _emscripten_glUniform4fv,
  /** @export */
  emscripten_glUniform4i: _emscripten_glUniform4i,
  /** @export */
  emscripten_glUniform4iv: _emscripten_glUniform4iv,
  /** @export */
  emscripten_glUniformMatrix2fv: _emscripten_glUniformMatrix2fv,
  /** @export */
  emscripten_glUniformMatrix3fv: _emscripten_glUniformMatrix3fv,
  /** @export */
  emscripten_glUniformMatrix4fv: _emscripten_glUniformMatrix4fv,
  /** @export */
  emscripten_glUseProgram: _emscripten_glUseProgram,
  /** @export */
  emscripten_glValidateProgram: _emscripten_glValidateProgram,
  /** @export */
  emscripten_glVertexAttrib1f: _emscripten_glVertexAttrib1f,
  /** @export */
  emscripten_glVertexAttrib1fv: _emscripten_glVertexAttrib1fv,
  /** @export */
  emscripten_glVertexAttrib2f: _emscripten_glVertexAttrib2f,
  /** @export */
  emscripten_glVertexAttrib2fv: _emscripten_glVertexAttrib2fv,
  /** @export */
  emscripten_glVertexAttrib3f: _emscripten_glVertexAttrib3f,
  /** @export */
  emscripten_glVertexAttrib3fv: _emscripten_glVertexAttrib3fv,
  /** @export */
  emscripten_glVertexAttrib4f: _emscripten_glVertexAttrib4f,
  /** @export */
  emscripten_glVertexAttrib4fv: _emscripten_glVertexAttrib4fv,
  /** @export */
  emscripten_glVertexAttribDivisorANGLE: _emscripten_glVertexAttribDivisorANGLE,
  /** @export */
  emscripten_glVertexAttribPointer: _emscripten_glVertexAttribPointer,
  /** @export */
  emscripten_glViewport: _emscripten_glViewport,
  /** @export */
  emscripten_has_asyncify: _emscripten_has_asyncify,
  /** @export */
  emscripten_request_fullscreen_strategy: _emscripten_request_fullscreen_strategy,
  /** @export */
  emscripten_request_pointerlock: _emscripten_request_pointerlock,
  /** @export */
  emscripten_resize_heap: _emscripten_resize_heap,
  /** @export */
  emscripten_sample_gamepad_data: _emscripten_sample_gamepad_data,
  /** @export */
  emscripten_set_beforeunload_callback_on_thread: _emscripten_set_beforeunload_callback_on_thread,
  /** @export */
  emscripten_set_blur_callback_on_thread: _emscripten_set_blur_callback_on_thread,
  /** @export */
  emscripten_set_canvas_element_size: _emscripten_set_canvas_element_size,
  /** @export */
  emscripten_set_element_css_size: _emscripten_set_element_css_size,
  /** @export */
  emscripten_set_focus_callback_on_thread: _emscripten_set_focus_callback_on_thread,
  /** @export */
  emscripten_set_fullscreenchange_callback_on_thread: _emscripten_set_fullscreenchange_callback_on_thread,
  /** @export */
  emscripten_set_gamepadconnected_callback_on_thread: _emscripten_set_gamepadconnected_callback_on_thread,
  /** @export */
  emscripten_set_gamepaddisconnected_callback_on_thread: _emscripten_set_gamepaddisconnected_callback_on_thread,
  /** @export */
  emscripten_set_keydown_callback_on_thread: _emscripten_set_keydown_callback_on_thread,
  /** @export */
  emscripten_set_keypress_callback_on_thread: _emscripten_set_keypress_callback_on_thread,
  /** @export */
  emscripten_set_keyup_callback_on_thread: _emscripten_set_keyup_callback_on_thread,
  /** @export */
  emscripten_set_main_loop: _emscripten_set_main_loop,
  /** @export */
  emscripten_set_main_loop_timing: _emscripten_set_main_loop_timing,
  /** @export */
  emscripten_set_mousedown_callback_on_thread: _emscripten_set_mousedown_callback_on_thread,
  /** @export */
  emscripten_set_mouseenter_callback_on_thread: _emscripten_set_mouseenter_callback_on_thread,
  /** @export */
  emscripten_set_mouseleave_callback_on_thread: _emscripten_set_mouseleave_callback_on_thread,
  /** @export */
  emscripten_set_mousemove_callback_on_thread: _emscripten_set_mousemove_callback_on_thread,
  /** @export */
  emscripten_set_mouseup_callback_on_thread: _emscripten_set_mouseup_callback_on_thread,
  /** @export */
  emscripten_set_orientationchange_callback_on_thread: _emscripten_set_orientationchange_callback_on_thread,
  /** @export */
  emscripten_set_pointerlockchange_callback_on_thread: _emscripten_set_pointerlockchange_callback_on_thread,
  /** @export */
  emscripten_set_resize_callback_on_thread: _emscripten_set_resize_callback_on_thread,
  /** @export */
  emscripten_set_timeout: _emscripten_set_timeout,
  /** @export */
  emscripten_set_touchcancel_callback_on_thread: _emscripten_set_touchcancel_callback_on_thread,
  /** @export */
  emscripten_set_touchend_callback_on_thread: _emscripten_set_touchend_callback_on_thread,
  /** @export */
  emscripten_set_touchmove_callback_on_thread: _emscripten_set_touchmove_callback_on_thread,
  /** @export */
  emscripten_set_touchstart_callback_on_thread: _emscripten_set_touchstart_callback_on_thread,
  /** @export */
  emscripten_set_visibilitychange_callback_on_thread: _emscripten_set_visibilitychange_callback_on_thread,
  /** @export */
  emscripten_set_wheel_callback_on_thread: _emscripten_set_wheel_callback_on_thread,
  /** @export */
  emscripten_set_window_title: _emscripten_set_window_title,
  /** @export */
  emscripten_sleep: _emscripten_sleep,
  /** @export */
  emscripten_webgl_create_context: _emscripten_webgl_create_context,
  /** @export */
  emscripten_webgl_destroy_context: _emscripten_webgl_destroy_context,
  /** @export */
  emscripten_webgl_make_context_current: _emscripten_webgl_make_context_current,
  /** @export */
  environ_get: _environ_get,
  /** @export */
  environ_sizes_get: _environ_sizes_get,
  /** @export */
  exit: _exit,
  /** @export */
  fd_close: _fd_close,
  /** @export */
  fd_read: _fd_read,
  /** @export */
  fd_seek: _fd_seek,
  /** @export */
  fd_write: _fd_write,
  /** @export */
  invoke_iii,
  /** @export */
  invoke_iiii,
  /** @export */
  invoke_iiiii,
  /** @export */
  invoke_v,
  /** @export */
  invoke_vii,
  /** @export */
  invoke_viii,
  /** @export */
  invoke_viiii,
  /** @export */
  invoke_viiiii,
  /** @export */
  invoke_viiiiiii,
  /** @export */
  proc_exit: _proc_exit
};
var wasmExports;
createWasm();

function invoke_viiii(index,a1,a2,a3,a4) {
  var sp = stackSave();
  try {
    getWasmTableEntry(index)(a1,a2,a3,a4);
  } catch(e) {
    stackRestore(sp);
    if (e !== e+0) throw e;
    _setThrew(1, 0);
  }
}

function invoke_iii(index,a1,a2) {
  var sp = stackSave();
  try {
    return getWasmTableEntry(index)(a1,a2);
  } catch(e) {
    stackRestore(sp);
    if (e !== e+0) throw e;
    _setThrew(1, 0);
  }
}

function invoke_iiiii(index,a1,a2,a3,a4) {
  var sp = stackSave();
  try {
    return getWasmTableEntry(index)(a1,a2,a3,a4);
  } catch(e) {
    stackRestore(sp);
    if (e !== e+0) throw e;
    _setThrew(1, 0);
  }
}

function invoke_v(index) {
  var sp = stackSave();
  try {
    getWasmTableEntry(index)();
  } catch(e) {
    stackRestore(sp);
    if (e !== e+0) throw e;
    _setThrew(1, 0);
  }
}

function invoke_iiii(index,a1,a2,a3) {
  var sp = stackSave();
  try {
    return getWasmTableEntry(index)(a1,a2,a3);
  } catch(e) {
    stackRestore(sp);
    if (e !== e+0) throw e;
    _setThrew(1, 0);
  }
}

function invoke_vii(index,a1,a2) {
  var sp = stackSave();
  try {
    getWasmTableEntry(index)(a1,a2);
  } catch(e) {
    stackRestore(sp);
    if (e !== e+0) throw e;
    _setThrew(1, 0);
  }
}

function invoke_viii(index,a1,a2,a3) {
  var sp = stackSave();
  try {
    getWasmTableEntry(index)(a1,a2,a3);
  } catch(e) {
    stackRestore(sp);
    if (e !== e+0) throw e;
    _setThrew(1, 0);
  }
}

function invoke_viiiii(index,a1,a2,a3,a4,a5) {
  var sp = stackSave();
  try {
    getWasmTableEntry(index)(a1,a2,a3,a4,a5);
  } catch(e) {
    stackRestore(sp);
    if (e !== e+0) throw e;
    _setThrew(1, 0);
  }
}

function invoke_viiiiiii(index,a1,a2,a3,a4,a5,a6,a7) {
  var sp = stackSave();
  try {
    getWasmTableEntry(index)(a1,a2,a3,a4,a5,a6,a7);
  } catch(e) {
    stackRestore(sp);
    if (e !== e+0) throw e;
    _setThrew(1, 0);
  }
}


// include: postamble.js
// === Auto-generated postamble setup entry stuff ===

var calledRun;

function stackCheckInit() {
  // This is normally called automatically during __wasm_call_ctors but need to
  // get these values before even running any of the ctors so we call it redundantly
  // here.
  _emscripten_stack_init();
  // TODO(sbc): Move writeStackCookie to native to to avoid this.
  writeStackCookie();
}

function run() {

  if (runDependencies > 0) {
    dependenciesFulfilled = run;
    return;
  }

  stackCheckInit();

  preRun();

  // a preRun added a dependency, run will be called later
  if (runDependencies > 0) {
    dependenciesFulfilled = run;
    return;
  }

  function doRun() {
    // run may have just been called through dependencies being fulfilled just in this very frame,
    // or while the async setStatus time below was happening
    assert(!calledRun);
    calledRun = true;
    Module['calledRun'] = true;

    if (ABORT) return;

    initRuntime();

    Module['onRuntimeInitialized']?.();
    consumedModuleProp('onRuntimeInitialized');

    assert(!Module['_main'], 'compiled without a main, but one is present. if you added it from JS, use Module["onRuntimeInitialized"]');

    postRun();
  }

  if (Module['setStatus']) {
    Module['setStatus']('Running...');
    setTimeout(() => {
      setTimeout(() => Module['setStatus'](''), 1);
      doRun();
    }, 1);
  } else
  {
    doRun();
  }
  checkStackCookie();
}

function checkUnflushedContent() {
  // Compiler settings do not allow exiting the runtime, so flushing
  // the streams is not possible. but in ASSERTIONS mode we check
  // if there was something to flush, and if so tell the user they
  // should request that the runtime be exitable.
  // Normally we would not even include flush() at all, but in ASSERTIONS
  // builds we do so just for this check, and here we see if there is any
  // content to flush, that is, we check if there would have been
  // something a non-ASSERTIONS build would have not seen.
  // How we flush the streams depends on whether we are in SYSCALLS_REQUIRE_FILESYSTEM=0
  // mode (which has its own special function for this; otherwise, all
  // the code is inside libc)
  var oldOut = out;
  var oldErr = err;
  var has = false;
  out = err = (x) => {
    has = true;
  }
  try { // it doesn't matter if it fails
    _fflush(0);
    // also flush in the JS FS layer
    ['stdout', 'stderr'].forEach((name) => {
      var info = FS.analyzePath('/dev/' + name);
      if (!info) return;
      var stream = info.object;
      var rdev = stream.rdev;
      var tty = TTY.ttys[rdev];
      if (tty?.output?.length) {
        has = true;
      }
    });
  } catch(e) {}
  out = oldOut;
  err = oldErr;
  if (has) {
    warnOnce('stdio streams had content in them that was not flushed. you should set EXIT_RUNTIME to 1 (see the Emscripten FAQ), or make sure to emit a newline when you printf etc.');
  }
}

function preInit() {
  if (Module['preInit']) {
    if (typeof Module['preInit'] == 'function') Module['preInit'] = [Module['preInit']];
    while (Module['preInit'].length > 0) {
      Module['preInit'].shift()();
    }
  }
  consumedModuleProp('preInit');
}

preInit();
run();

// end include: postamble.js


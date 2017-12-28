package protocol

import (
	"encoding/json"

	debugger "github.com/mkenney/go-chrome/protocol/debugger"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Debugger is a struct that provides a namespace for the Chrome Debugger protocol methods. The
Debugger protocol exposes JavaScript debugging capabilities. It allows setting and removing
breakpoints, stepping through execution, exploring stack traces, etc.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/
*/
var Debugger = _debugger{}

type _debugger struct{}

/*
ContinueToLocation continues execution until specific location is reached.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-continueToLocation
*/
func (_debugger) ContinueToLocation(
	socket sock.Socketer,
	params *debugger.ContinueToLocationParams,
) error {
	command := sock.NewCommand("Debugger.continueToLocation", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables debugger for given page.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-disable
*/
func (_debugger) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Debugger.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables debugger for the given page. Clients should not assume that the debugging has been
enabled until the result for this command is received.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-enable
*/
func (_debugger) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Debugger.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
EvaluateOnCallFrame evaluates expression on a given call frame.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-evaluateOnCallFrame
*/
func (_debugger) EvaluateOnCallFrame(
	socket sock.Socketer,
	params *debugger.EvaluateOnCallFrameParams,
) (debugger.EvaluateOnCallFrameResult, error) {
	command := sock.NewCommand("Debugger.evaluateOnCallFrame", params)
	result := debugger.EvaluateOnCallFrameResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetPossibleBreakpoints returns possible locations for breakpoint. scriptId in start and end range
locations should be the same.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getPossibleBreakpoints
*/
func (_debugger) GetPossibleBreakpoints(
	socket sock.Socketer,
	params *debugger.GetPossibleBreakpointsParams,
) (debugger.GetPossibleBreakpointsResult, error) {
	command := sock.NewCommand("Debugger.getPossibleBreakpoints", params)
	result := debugger.GetPossibleBreakpointsResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetScriptSource returns source for the script with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getScriptSource
*/
func (_debugger) GetScriptSource(
	socket sock.Socketer,
	params *debugger.GetScriptSourceParams,
) (debugger.GetScriptSourceResult, error) {
	command := sock.NewCommand("Debugger.getScriptSource", params)
	result := debugger.GetScriptSourceResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetStackTrace returns stack trace with given stackTraceId. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getStackTrace
*/
func (_debugger) GetStackTrace(
	socket sock.Socketer,
	params *debugger.GetStackTraceParams,
) (debugger.GetStackTraceResult, error) {
	command := sock.NewCommand("Debugger.getStackTrace", params)
	result := debugger.GetStackTraceResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
Pause stops on the next JavaScript statement.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pause
*/
func (_debugger) Pause(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Debugger.pause", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
PauseOnAsyncCall EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pauseOnAsyncCall
*/
func (_debugger) PauseOnAsyncCall(
	socket sock.Socketer,
	params *debugger.PauseOnAsyncCallParams,
) error {
	command := sock.NewCommand("Debugger.pauseOnAsyncCall", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
RemoveBreakpoint removes JavaScript breakpoint.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-removeBreakpoint
*/
func (_debugger) RemoveBreakpoint(
	socket sock.Socketer,
	params *debugger.RemoveBreakpointParams,
) error {
	command := sock.NewCommand("Debugger.removeBreakpoint", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
RestartFrame restarts particular call frame from the beginning.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-restartFrame
*/
func (_debugger) RestartFrame(
	socket sock.Socketer,
	params *debugger.RestartFrameParams,
) (debugger.RestartFrameResult, error) {
	command := sock.NewCommand("Debugger.restartFrame", params)
	result := debugger.RestartFrameResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
Resume resumes JavaScript execution.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-resume
*/
func (_debugger) Resume(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Debugger.resume", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
ScheduleStepIntoAsync is deprecated - use Debugger.stepInto with breakOnAsyncCall and
Debugger.pauseOnAsyncTask instead. Steps into next scheduled async task if any is scheduled before
next pause. Returns success when async task is actually scheduled, returns error if no task were
scheduled or another scheduleStepIntoAsync was called. EXPERIMENTAL. DEPRECATED.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-scheduleStepIntoAsync
*/
func (_debugger) ScheduleStepIntoAsync(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Debugger.scheduleStepIntoAsync", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
SearchInContent searches for given string in script content.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-searchInContent
*/
func (_debugger) SearchInContent(
	socket sock.Socketer,
	params *debugger.SearchInContentParams,
) (debugger.SearchInContentResult, error) {
	command := sock.NewCommand("Debugger.searchInContent", params)
	result := debugger.SearchInContentResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
SetAsyncCallStackDepth enables or disables async call stacks tracking.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setAsyncCallStackDepth
*/
func (_debugger) SetAsyncCallStackDepth(
	socket sock.Socketer,
	params *debugger.SetAsyncCallStackDepthParams,
) error {
	command := sock.NewCommand("Debugger.setAsyncCallStackDepth", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetBlackboxPatterns replaces previous blackbox patterns with passed ones. Forces backend to skip
stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed
script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxPatterns
*/
func (_debugger) SetBlackboxPatterns(
	socket sock.Socketer,
	params *debugger.SetBlackboxPatternsParams,
) error {
	command := sock.NewCommand("Debugger.setBlackboxPatterns", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetBlackboxedRanges makes backend skip steps in the script in blackboxed ranges. VM will try leave
blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if
unsuccessful. Positions array contains positions where blackbox state is changed. First interval
isn't blackboxed. Array should be sorted.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxedRanges
*/
func (_debugger) SetBlackboxedRanges(
	socket sock.Socketer,
	params *debugger.SetBlackboxedRangesParams,
) error {
	command := sock.NewCommand("Debugger.setBlackboxedRanges", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetBreakpoint sets JavaScript breakpoint at a given location.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpoint
*/
func (_debugger) SetBreakpoint(
	socket sock.Socketer,
	params *debugger.SetBreakpointParams,
) (debugger.SetBreakpointResult, error) {
	command := sock.NewCommand("Debugger.setBreakpoint", params)
	result := debugger.SetBreakpointResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
SetBreakpointByURL sets JavaScript breakpoint at given location specified either by URL or URL
regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and
returned in locations property. Further matching script parsing will result in subsequent
breakpointResolved events issued. This logical breakpoint will survive page reloads.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointByUrl
*/
func (_debugger) SetBreakpointByURL(
	socket sock.Socketer,
	params *debugger.SetBreakpointByURLParams,
) (debugger.SetBreakpointByURLResult, error) {
	command := sock.NewCommand("Debugger.setBreakpointByUrl", params)
	result := debugger.SetBreakpointByURLResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
SetBreakpointsActive activates / deactivates all breakpoints on the page.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointsActive
*/
func (_debugger) SetBreakpointsActive(
	socket sock.Socketer,
	params *debugger.SetBreakpointsActiveParams,
) error {
	command := sock.NewCommand("Debugger.setBreakpointsActive", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetPauseOnExceptions defines the pause on exceptions state. Can be set to stop on all exceptions,
uncaught exceptions or no exceptions. Initial pause on exceptions state is none.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setPauseOnExceptions
*/
func (_debugger) SetPauseOnExceptions(
	socket sock.Socketer,
	params *debugger.SetPauseOnExceptionsParams,
) error {
	command := sock.NewCommand("Debugger.setPauseOnExceptions", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetReturnValue changes return value in top frame. Available only at return break position.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setReturnValue
*/
func (_debugger) SetReturnValue(
	socket sock.Socketer,
	params *debugger.SetReturnValueParams,
) error {
	command := sock.NewCommand("Debugger.setReturnValue", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetScriptSource edits JavaScript source live.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setScriptSource
*/
func (_debugger) SetScriptSource(
	socket sock.Socketer,
	params *debugger.SetScriptSourceParams,
) (debugger.SetScriptSourceResult, error) {
	command := sock.NewCommand("Debugger.setScriptSource", params)
	result := debugger.SetScriptSourceResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
SetSkipAllPauses makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setSkipAllPauses
*/
func (_debugger) SetSkipAllPauses(
	socket sock.Socketer,
	params *debugger.SetSkipAllPausesParams,
) error {
	command := sock.NewCommand("Debugger.setSkipAllPauses", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetVariableValue changes value of variable in a callframe. Object-based scopes are not supported and
must be mutated manually.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setVariableValue
*/
func (_debugger) SetVariableValue(
	socket sock.Socketer,
	params *debugger.SetVariableValueParams,
) error {
	command := sock.NewCommand("Debugger.setVariableValue", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
StepInto steps into the function call.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepInto
*/
func (_debugger) StepInto(
	socket sock.Socketer,
	params *debugger.StepIntoParams,
) error {
	command := sock.NewCommand("Debugger.stepInto", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
StepOut steps out of the function call.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOut
*/
func (_debugger) StepOut(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Debugger.stepOut", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
StepOver steps over the statement.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOver
*/
func (_debugger) StepOver(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Debugger.stepOver", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnBreakpointResolved adds a handler to the Debugger.breakpointResolved event.
Debugger.breakpointResolved fires when breakpoint is resolved to an actual script and location.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-breakpointResolved
*/
func (_debugger) OnBreakpointResolved(
	socket sock.Socketer,
	callback func(event *debugger.BreakpointResolvedEvent),
) {
	handler := sock.NewEventHandler(
		"Debugger.breakpointResolved",
		func(response *sock.Response) {
			event := &debugger.BreakpointResolvedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnPaused adds a handler to the Debugger.paused event. Debugger.paused fires when the virtual machine
stopped on breakpoint or exception or any other stop criteria.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-paused
*/
func (_debugger) OnPaused(
	socket sock.Socketer,
	callback func(event *debugger.PausedEvent),
) {
	handler := sock.NewEventHandler(
		"Debugger.paused",
		func(response *sock.Response) {
			event := &debugger.PausedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnResumed adds a handler to the Debugger.resumed event. Debugger.resumed fires when the virtual
machine resumes execution.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-resumed
*/
func (_debugger) OnResumed(
	socket sock.Socketer,
	callback func(event *debugger.ResumedEvent),
) {
	handler := sock.NewEventHandler(
		"Debugger.resumed",
		func(response *sock.Response) {
			event := &debugger.ResumedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnScriptFailedToParse adds a handler to the Debugger.scriptFailedToParse event.
Debugger.scriptFailedToParse fires when the virtual machine fails to parse the script.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-scriptFailedToParse
*/
func (_debugger) OnScriptFailedToParse(
	socket sock.Socketer,
	callback func(event *debugger.ScriptFailedToParseEvent),
) {
	handler := sock.NewEventHandler(
		"Debugger.scriptFailedToParse",
		func(response *sock.Response) {
			event := &debugger.ScriptFailedToParseEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnScriptParsed adds a handler to the Debugger.ScriptParsed event. Debugger.ScriptParsed fires when
virtual machine parses script. This event is also fired for all known and uncollected scripts upon
enabling debugger.

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-scriptParsed
*/
func (_debugger) OnScriptParsed(
	socket sock.Socketer,
	callback func(event *debugger.ScriptParsedEvent),
) {
	handler := sock.NewEventHandler(
		"Debugger.scriptParsed",
		func(response *sock.Response) {
			event := &debugger.ScriptParsedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
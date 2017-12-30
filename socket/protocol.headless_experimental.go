package socket

import (
	"encoding/json"

	headlessExperimental "github.com/mkenney/go-chrome/protocol/headless_experimental"
	log "github.com/sirupsen/logrus"
)

/*
HeadlessExperimentalProtocol provides a namespace for the Chrome HeadlessExperimental protocol
methods. The HeadlessExperimental protocol provides experimental commands only supported in headless
mode.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/ EXPERIMENTAL.
*/
type HeadlessExperimentalProtocol struct {
	Socket Socketer
}

/*
BeginFrame sends a BeginFrame to the target and returns when the frame was completed. Optionally
captures a screenshot from the resulting frame. Requires that the target was created with enabled
BeginFrameControl.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-beginFrame
*/
func (protocol *HeadlessExperimentalProtocol) BeginFrame(
	params *headlessExperimental.BeginFrameParams,
) (*headlessExperimental.BeginFrameResult, error) {
	command := NewCommand("HeadlessExperimental.beginFrame", params)
	result := &headlessExperimental.BeginFrameResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
Disable disables headless events for the target.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-disable
*/
func (protocol *HeadlessExperimentalProtocol) Disable() error {
	command := NewCommand("HeadlessExperimental.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables headless events for the target.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-enable
*/
func (protocol *HeadlessExperimentalProtocol) Enable() error {
	command := NewCommand("HeadlessExperimental.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
OnMainFrameReadyForScreenshots adds a handler to the
HeadlessExperimental.mainFrameReadyForScreenshots event.
HeadlessExperimental.mainFrameReadyForScreenshots fires when the main frame has first submitted a
frame to the browser. May only be fired while a BeginFrame is in flight. Before this event,
screenshotting requests may fail.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-mainFrameReadyForScreenshots
*/
func (protocol *HeadlessExperimentalProtocol) OnMainFrameReadyForScreenshots(
	callback func(event *headlessExperimental.MainFrameReadyForScreenshotsEvent),
) {
	handler := NewEventHandler(
		"HeadlessExperimental.mainFrameReadyForScreenshots",
		func(response *Response) {
			event := &headlessExperimental.MainFrameReadyForScreenshotsEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnNeedsBeginFramesChanged adds a handler to the HeadlessExperimental.needsBeginFramesChanged event.
HeadlessExperimental.needsBeginFramesChanged fires when the target starts or stops needing
BeginFrames.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-needsBeginFramesChanged
*/
func (protocol *HeadlessExperimentalProtocol) OnNeedsBeginFramesChanged(
	callback func(event *headlessExperimental.NeedsBeginFramesChangedEvent),
) {
	handler := NewEventHandler(
		"HeadlessExperimental.needsBeginFramesChanged",
		func(response *Response) {
			event := &headlessExperimental.NeedsBeginFramesChangedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
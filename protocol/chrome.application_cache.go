package protocol

import (
	"encoding/json"

	applicationCache "github.com/mkenney/go-chrome/protocol/application_cache"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
ApplicationCache is a struct that provides a namespace for the Chrome Animation protocol methods.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/
*/
var ApplicationCache = _applicationcache{}

type _applicationcache struct{}

/*
Enable enables application cache domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-enable
*/
func (_applicationcache) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("ApplicationCache.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetForFrame returns relevant application cache data for the document
in given frame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
*/
func (_applicationcache) GetForFrame(
	socket sock.Socketer,
	params *applicationCache.GetForFrameParams,
) (applicationCache.GetForFrameResult, error) {
	command := sock.NewCommand("ApplicationCache.getApplicationCacheForFrame", params)
	result := applicationCache.GetForFrameResult{}
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
GetFramesWithManifests returns array of frame identifiers with manifest urls for each frame
containing a document associated with some application cache.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getFramesWithManifests
*/
func (_applicationcache) GetFramesWithManifests(
	socket sock.Socketer,
) (applicationCache.GetFramesWithManifestsResult, error) {
	command := sock.NewCommand("ApplicationCache.getFramesWithManifests", nil)
	result := applicationCache.GetFramesWithManifestsResult{}
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
GetManifestForFrame returns manifest URL for document in the given frame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getManifestForFrame
*/
func (_applicationcache) GetManifestForFrame(
	socket sock.Socketer,
	params *applicationCache.GetManifestForFrameParams,
) (applicationCache.GetManifestForFrameResult, error) {
	command := sock.NewCommand("ApplicationCache.getManifestForFrame", params)
	result := applicationCache.GetManifestForFrameResult{}
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
OnApplicationCacheStatusUpdated adds a handler to the ApplicationCache.StatusUpdated event.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-applicationCacheStatusUpdated
*/
func (_applicationcache) OnApplicationCacheStatusUpdated(
	socket sock.Socketer,
	callback func(event *applicationCache.StatusUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"ApplicationCache.applicationCacheStatusUpdated",
		func(response *sock.Response) {
			event := &applicationCache.StatusUpdatedEvent{}
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
OnNetworkStateUpdated adds a handler to the ApplicationCache.StatusUpdated event.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-networkStateUpdated
*/
func (_applicationcache) OnNetworkStateUpdated(
	socket sock.Socketer,
	callback func(event *applicationCache.StatusUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"ApplicationCache.networkStateUpdated",
		func(response *sock.Response) {
			event := &applicationCache.StatusUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
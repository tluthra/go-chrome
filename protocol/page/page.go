/*
Package page provides type definitions for use with the Chrome Page protocol

https://chromedevtools.github.io/devtools-protocol/tot/Page/
*/
package page

import (
	"github.com/mkenney/go-chrome/protocol/debugger"
	"github.com/mkenney/go-chrome/protocol/runtime"
)

/*
AddScriptToEvaluateOnLoadParams represents Page.addScriptToEvaluateOnLoad parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnLoad
*/
type AddScriptToEvaluateOnLoadParams struct {
	ScriptSource string `json:"scriptSource"`
}

/*
AddScriptToEvaluateOnLoadResult represents the result of calls to Page.addScriptToEvaluateOnLoad.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnLoad
*/
type AddScriptToEvaluateOnLoadResult struct {
	// Identifier of the added script.
	Identifier ScriptIdentifier `json:"identifier"`
}

/*
AddScriptToEvaluateOnNewDocumentParams represents Page.addScriptToEvaluateOnNewDocument
parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnNewDocument
*/
type AddScriptToEvaluateOnNewDocumentParams struct {
	Source string `json:"source"`
}

/*
AddScriptToEvaluateOnNewDocumentResult represents the result of calls to
Page.addScriptToEvaluateOnNewDocument.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnNewDocument
*/
type AddScriptToEvaluateOnNewDocumentResult struct {
	// Identifier of the added script.
	Identifier ScriptIdentifier `json:"identifier"`
}

/*
CaptureScreenshotParams represents Page.captureScreenshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot
*/
type CaptureScreenshotParams struct {
	// Optional. Image compression format (defaults to png). Allowed values: jpeg, png.
	Format string `json:"format,omitempty"`

	// Optional. Compression quality from range [0..100] (jpeg only).
	Quality int `json:"quality,omitempty"`

	// Optional. Capture the screenshot of a given region only.
	//
	// This expects an instance of Viewport, but that doesn't omitempty correctly so it must be
	// added manually.
	Clip interface{} `json:"clip,omitempty"`

	// Optional. Capture the screenshot from the surface, rather than the view. Defaults to true.
	// EXPERIMENTAL
	FromSurface bool `json:"fromSurface,omitempty"`
}

/*
CaptureScreenshotResult represents the result of calls to Page.captureScreenshot.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot
*/
type CaptureScreenshotResult struct {
	// Base64-encoded image data.
	Data string `json:"data"`
}

/*
CreateIsolatedWorldParams represents Page.createIsolatedWorld parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-createIsolatedWorld
*/
type CreateIsolatedWorldParams struct {
	// ID of the frame in which the isolated world should be created.
	FrameID FrameID `json:"frameId"`

	// Optional. An optional name which is reported in the Execution Context.
	WorldName string `json:"worldName,omitempty"`

	// Optional. Whether or not universal access should be granted to the isolated world. This is a
	// powerful option, use with caution.
	GrantUniveralAccess bool `json:"grantUniveralAccess,omitempty"`
}

/*
CreateIsolatedWorldResult represents the result of calls to Page.createIsolatedWorld.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-createIsolatedWorld
*/
type CreateIsolatedWorldResult struct {
	// Execution context of the isolated world.
	ExecutionContextID runtime.ExecutionContextID `json:"executionContextId"`
}

/*
GetAppManifestParams represents Page.getAppManifest parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getAppManifest
*/
type GetAppManifestParams struct {
	// Manifest location.
	URL string `json:"url"`

	// Errors.
	Errors []AppManifestError `json:"errors"`

	// Optional. Manifest content.
	Data string `json:"data,omitempty"`
}

/*
GetFrameTreeResult represents the result of calls to Page.getFrameTree.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getFrameTree
*/
type GetFrameTreeResult struct {
	// Present frame tree structure.
	FrameTree FrameTree `json:"frameTree"`
}

/*
GetLayoutMetricsResult represents the result of calls to Page.getLayoutMetrics.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getLayoutMetrics
*/
type GetLayoutMetricsResult struct {
	// Metrics relating to the layout viewport.
	LayoutViewport LayoutViewport `json:"layoutViewport"`

	// Metrics relating to the visual viewport.
	VisualViewport VisualViewport `json:"visualViewport"`

	// Size of scrollable area. Rect is a local implementation of DOM.Rect
	ContentSize Rect `json:"contentSize"`
}

/*
GetNavigationHistoryResult represents the result of calls to Page.getNavigationHistory.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getNavigationHistory
*/
type GetNavigationHistoryResult struct {
	// Index of the current navigation history entry.
	CurrentIndex int `json:"currentIndex"`

	// Array of navigation history entries.
	Entries []NavigationEntry `json:"entries"`
}

/*
GetResourceContentParams represents Page.getResourceContent parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceContent
*/
type GetResourceContentParams struct {
	// Frame ID to get resource for.
	FrameID FrameID `json:"frameId"`

	// URL of the resource to get content for.
	URL string `json:"url"`
}

/*
GetResourceContentResult represents the result of calls to Page.getResourceContent.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceContent
*/
type GetResourceContentResult struct {
	// Resource content.
	Content string `json:"content"`

	// True, if content was served as base64.
	Base64Encoded bool `json:"base64Encoded"`
}

/*
GetResourceTreeResult represents the result of calls to Page.getResourceTree.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceTree
*/
type GetResourceTreeResult struct {
	// Present frame / resource tree structure.
	FrameTree FrameResourceTree `json:"frameTree"`
}

/*
HandleJavaScriptDialogParams represents Page.handleJavaScriptDialog parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-handleJavaScriptDialog
*/
type HandleJavaScriptDialogParams struct {
	// Whether to accept or dismiss the dialog.
	Accept bool `json:"accept"`

	// Optional. The text to enter into the dialog prompt before accepting. Used only if this is a
	// prompt dialog.
	PromptText string `json:"promptText,omitempty"`
}

/*
NavigateParams represents Page.navigate parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/
type NavigateParams struct {
	// URL to navigate the page to.
	URL string `json:"url"`

	// Optional. Referrer URL.
	Referrer string `json:"referrer,omitempty"`

	// Optional. Intended transition type.
	TransitionType TransitionType `json:"transitionType,omitempty"`
}

/*
NavigateResult represents the result of calls to Page.navigate.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/
type NavigateResult struct {
	// Frame ID that has navigated (or failed to navigate).
	FrameID FrameID `json:"frameId"`

	// Loader identifier.
	LoaderID LoaderID `json:"loaderId"`

	// User friendly error message, present if and only if navigation has failed.
	ErrorText string `json:"errorText"`
}

/*
NavigateToHistoryEntryParams represents Page.navigateToHistoryEntry parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigateToHistoryEntry
*/
type NavigateToHistoryEntryParams struct {
	// Unique ID of the entry to navigate to.
	EntryID int `json:"entryId"`
}

/*
PrintToPDFParams represents Page.printToPDF parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-printToPDF
*/
type PrintToPDFParams struct {
	// Optional. Paper orientation. Defaults to false.
	Landscape bool `json:"landscape,omitempty"`

	// Optional. Display header and footer. Defaults to false.
	DisplayHeaderFooter bool `json:"displayHeaderFooter,omitempty"`

	// Optional. Print background graphics. Defaults to false.
	PrintBackground bool `json:"printBackground,omitempty"`

	// Optional. Scale of the webpage rendering. Defaults to 1.
	Scale float64 `json:"scale,omitempty"`

	// Optional. Paper width in inches. Defaults to 8.5 inches.
	PaperWidth float64 `json:"paperWidth,omitempty"`

	// Optional. Paper height in inches. Defaults to 11 inches.
	PaperHeight float64 `json:"paperHeight,omitempty"`

	// Optional. Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginTop float64 `json:"marginTop,omitempty"`

	// Optional. Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom float64 `json:"marginBottom,omitempty"`

	// Optional. Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft float64 `json:"marginLeft,omitempty"`

	// Optional. Right margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight float64 `json:"marginRight,omitempty"`

	// Optional. Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which
	// means print all pages.
	PageRanges string `json:"pageRanges,omitempty"`

	// Optional. Whether to silently ignore invalid but successfully parsed page ranges, such as
	// '3-2'. Defaults to false.
	IgnoreInvalidPageRanges bool `json:"ignoreInvalidPageRanges,omitempty"`
}

/*
PrintToPDFResult represents the result of calls to Page.printToPDF.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-printToPDF
*/
type PrintToPDFResult struct {
	// Base64-encoded pdf data.
	Data string `json:"data"`
}

/*
ReloadParams represents Page.reload parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload
*/
type ReloadParams struct {
	// Optional. If true, browser cache is ignored (as if the user pressed Shift+refresh).
	IgnoreCache bool `json:"ignoreCache,omitempty"`

	// Optional. If set, the script will be injected into all frames of the inspected page after
	// reload.
	ScriptToEvaluateOnLoad string `json:"scriptToEvaluateOnLoad,omitempty"`
}

/*
RemoveScriptToEvaluateOnLoadParams represents Page.removeScriptToEvaluateOnLoad parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnLoad
*/
type RemoveScriptToEvaluateOnLoadParams struct {
	Identifier ScriptIdentifier `json:"identifier"`
}

/*
RemoveScriptToEvaluateOnNewDocumentParams represents Page.removeScriptToEvaluateOnNewDocument
parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnNewDocument
*/
type RemoveScriptToEvaluateOnNewDocumentParams struct {
	Identifier ScriptIdentifier `json:"identifier"`
}

/*
ScreencastFrameAckParams represents Page.screencastFrameAck parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-screencastFrameAck
*/
type ScreencastFrameAckParams struct {
	// Frame number.
	SessionID int `json:"sessionId"`
}

/*
SearchInResourceParams represents Page.searchInResource parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-searchInResource
*/
type SearchInResourceParams struct {
	// Frame ID for resource to search in.
	FrameID FrameID `json:"frameId"`

	// URL of the resource to search in.
	URL string `json:"url"`

	// String to search for.
	Query string `json:"query"`

	// Optional. If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`

	// Optional. If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

/*
SearchInResourceResult represents the result of calls to Page.searchInResource.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-searchInResource
*/
type SearchInResourceResult struct {
	// List of search matches.
	Result []debugger.SearchMatch `json:"result"`
}

/*
SetAdBlockingEnabledParams represents Page.setAdBlockingEnabled parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAdBlockingEnabled
*/
type SetAdBlockingEnabledParams struct {
	// Whether to block ads.
	Enabled bool `json:"enabled"`
}

/*
SetAutoAttachToCreatedPagesParams represents Page.setAutoAttachToCreatedPages parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAutoAttachToCreatedPages
*/
type SetAutoAttachToCreatedPagesParams struct {
	// If true, browser will open a new inspector window for every page created from this one.
	AutoAttach bool `json:"autoAttach"`
}

/*
SetDocumentContentParams represents Page.setDocumentContent parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDocumentContent
*/
type SetDocumentContentParams struct {
	// Frame ID to set HTML for.
	FrameID FrameID `json:"frameId"`

	// HTML content to set.
	HTML string `json:"html"`
}

/*
SetDownloadBehaviorParams represents Page.setDownloadBehavior parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDownloadBehavior
*/
type SetDownloadBehaviorParams struct {
	// Whether to allow all or deny all download requests, or use default Chrome behavior if
	// available (otherwise deny). Allowed values: deny, allow, default.
	Behavior string `json:"behavior"`

	// Optional. The default path to save downloaded files to. This is required if behavior is set
	// to 'allow'.
	DownloadPath string `json:"downloadPath,omitempty"`
}

/*
SetLifecycleEventsEnabledParams represents Page.setLifecycleEventsEnabled parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setLifecycleEventsEnabled
*/
type SetLifecycleEventsEnabledParams struct {
	// If true, starts emitting lifecycle events.
	Enabled bool `json:"enabled"`
}

/*
StartScreencastParams represents Page.startScreencast parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-startScreencast
*/
type StartScreencastParams struct {
	// Optional. Image compression format. Allowed values: jpeg, png.
	Format string `json:"format,omitempty"`

	// Optional. Compression quality from range [0..100].
	Quality int `json:"quality,omitempty"`

	// Optional. Maximum screenshot width.
	MaxWidth int `json:"maxWidth,omitempty"`

	// Optional. Maximum screenshot height.
	MaxHeight int `json:"maxHeight,omitempty"`

	// Optional. Send every n-th frame.
	EveryNthFrame int `json:"everyNthFrame,omitempty"`
}

/*
DOMContentEventFiredEvent represents Page.domContentEventFired event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-domContentEventFired
*/
type DOMContentEventFiredEvent struct {
	Timestamp MonotonicTime `json:"timestamp"`
}

/*
FrameAttachedEvent represents Page.frameAttached event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameAttached
*/
type FrameAttachedEvent struct {
	// ID of the frame that has been attached.
	FrameID FrameID `json:"frameId"`

	// Parent frame identifier.
	ParentFrameID FrameID `json:"parentFrameId"`

	// Optional. JavaScript stack trace of when frame was attached, only set if frame initiated from
	// script.
	Stack runtime.StackTrace `json:"stack,omitempty"`
}

/*
FrameClearedScheduledNavigationEvent represents Page.frameClearedScheduledNavigation event
data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameClearedScheduledNavigation
*/
type FrameClearedScheduledNavigationEvent struct {
	// ID of the frame that has cleared its scheduled navigation.
	FrameID FrameID `json:"frameId"`
}

/*
FrameDetachedEvent represents Page.frameDetached event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameDetached
*/
type FrameDetachedEvent struct {
	// ID of the frame that has been detached.
	FrameID FrameID `json:"frameId"`
}

/*
FrameNavigatedEvent represents Page.frameNavigated event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameNavigated
*/
type FrameNavigatedEvent struct {
	// Frame object.
	Frame Frame `json:"frame"`
}

/*
FrameResizedEvent represents Page.frameResized event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameResized
*/
type FrameResizedEvent struct{}

/*
FrameScheduledNavigationEvent represents Page.frameScheduledNavigation event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameScheduledNavigation
*/
type FrameScheduledNavigationEvent struct {
	// ID of the frame that has scheduled a navigation.
	FrameID FrameID `json:"frameId"`

	// Delay (in seconds) until the navigation is scheduled to begin. The navigation is not
	// guaranteed to start.
	Delay float64 `json:"delay"`

	// The reason for the navigation. Allowed values: formSubmissionGet, formSubmissionPost,
	// httpHeaderRefresh, scriptInitiated, metaTagRefresh, pageBlockInterstitial, reload.
	Reason string `json:"reason"`

	// The destination URL for the scheduled navigation.
	URL string `json:"url"`
}

/*
FrameStartedLoadingEvent represents Page.frameStartedLoading event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStartedLoading
*/
type FrameStartedLoadingEvent struct {
	// ID of the frame that has started loading.
	FrameID FrameID `json:"frameId"`
}

/*
FrameStoppedLoadingEvent represents Page.frameStoppedLoading event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStoppedLoading
*/
type FrameStoppedLoadingEvent struct {
	// ID of the frame that has stopped loading.
	FrameID FrameID `json:"frameId"`
}

/*
InterstitialHiddenEvent represents Page.interstitialHidden event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-interstitialHidden
*/
type InterstitialHiddenEvent struct{}

/*
InterstitialShownEvent represents Page.interstitialShown event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-interstitialShown
*/
type InterstitialShownEvent struct{}

/*
JavascriptDialogClosedEvent represents Page.javascriptDialogClosed event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-javascriptDialogClosed
*/
type JavascriptDialogClosedEvent struct {
	// Whether dialog was confirmed.
	Result bool `json:"result"`

	// User input in case of prompt.
	UserInput string `json:"userInput"`
}

/*
JavascriptDialogOpeningEvent represents Page.javascriptDialogOpening event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-javascriptDialogOpening
*/
type JavascriptDialogOpeningEvent struct {
	// Frame url.
	URL string `json:"url"`

	// Message that will be displayed by the dialog.
	Message string `json:"message"`

	// Dialog type.
	Type DialogType `json:"type"`

	// Optional. Default dialog prompt.
	DefaultPrompt string `json:"defaultPrompt,omitempty"`
}

/*
LifecycleEventEvent represents Page.lifecycleEvent event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-lifecycleEvent
*/
type LifecycleEventEvent struct {
	// ID of the frame.
	FrameID FrameID `json:"frameId"`

	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderID LoaderID `json:"loaderId"`

	// name.
	Name string `json:"name"`

	// timestamp.
	Timestamp MonotonicTime `json:"timestamp"`
}

/*
LoadEventFiredEvent represents Page.loadEventFired event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-loadEventFired
*/
type LoadEventFiredEvent struct {
	// timestamp.
	Timestamp MonotonicTime `json:"timestamp"`
}

/*
ScreencastFrameEvent represents Page.screencastFrame event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastFrame
*/
type ScreencastFrameEvent struct {
	// Base64-encoded compressed image.
	Data string `json:"data"`

	// Screencast frame metadata.
	Metadata ScreencastFrameMetadata `json:"metadata"`

	// Frame number.
	SessionID int `json:"sessionId"`
}

/*
ScreencastVisibilityChangedEvent represents Page.screencastVisibilityChanged event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastVisibilityChanged
*/
type ScreencastVisibilityChangedEvent struct {
	// True if the page is visible.
	Visible bool `json:"visible"`
}

/*
WindowOpenEvent represents Page.windowOpen event data.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-windowOpen
*/
type WindowOpenEvent struct {
	// The URL for the new window.
	URL string `json:"url"`

	// Window name.
	WindowName string `json:"windowName"`

	// An array of enabled window features.
	WindowFeatures []string `json:"windowFeatures"`

	// Whether or not it was triggered by user gesture.
	UserGesture bool `json:"userGesture"`
}

/*
LoaderID is the Unique loader identifier.
This is a duplicate of Network.LoaderID to avoid an invalid import cycle

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-LoaderId
*/
type LoaderID string

/*
MonotonicTime is the monotonically increasing time in seconds since an arbitrary point in the past.
This is a duplicate of Network.MonotonicTime to avoid an invalid import cycle

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-MonotonicTime
*/
type MonotonicTime int

/*
Rect defines a rectangle.
This is a duplicate of DOM.Rect to avoid an invalid import cycle

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-Rect
*/
type Rect struct {
	// X coordinate.
	X float64 `json:"x"`

	// Y coordinate.
	Y float64 `json:"y"`

	// Rectangle width.
	Width float64 `json:"width"`

	// Rectangle height.
	Height float64 `json:"height"`
}

/*
TimeSinceEpoch represents UTC time in seconds, counted from January 1, 1970.
This is a duplicate of DOM.Rect to avoid an invalid import cycle

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-TimeSinceEpoch
*/
type TimeSinceEpoch int

/*
AppManifestError defines an error that occurs while parsing an app manifest.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-AppManifestError
*/
type AppManifestError struct {
	// Error message.
	Message string `json:"message"`

	// If criticial, this is a non-recoverable parse error.
	Critical int `json:"critical"`

	// Error line.
	Line int `json:"line"`

	// Error column.
	Column int `json:"column"`
}

/*
DialogType defines the Javascript dialog type.

ALLOWED VALUES
	- alert
	- confirm
	- prompt
	- beforeunload

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-DialogType
*/
type DialogType string

/*
Frame details information about the Frame on the page.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-Frame
*/
type Frame struct {
	// Frame unique identifier.
	ID string `json:"id"`

	// Optional. Parent frame identifier.
	ParentID string `json:"parentId,omitempty"`

	// Identifier of the loader associated with this frame.
	LoaderID LoaderID `json:"loaderId"`

	// Optional. Frame's name as specified in the tag.
	Name string `json:"name,omitempty"`

	// Frame document's URL.
	URL string `json:"url"`

	// Frame document's security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Frame document's mimeType as determined by the browser.
	MimeType string `json:"mimeType"`

	// Optional. If the frame failed to load, this contains the URL that could not be loaded.
	// EXPERIMENTAL
	UnreachableURL string `json:"unreachableUrl,omitempty"`
}

/*
FrameID is a unique frame identifier

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameId
*/
type FrameID string

/*
FrameResource provides information about the Resource on the page. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameResource
*/
type FrameResource struct {
	// Resource URL.
	URL string `json:"url"`

	// Type of this resource.
	Type ResourceType `json:"type"`

	// Resource mimeType as determined by the browser.
	MimeType string `json:"mimeType"`

	// Optional. last-modified timestamp as reported by server.
	LastModified TimeSinceEpoch `json:"lastModified,omitempty"`

	// Optional. Resource content size.
	ContentSize int `json:"contentSize,omitempty"`

	// Optional. True if the resource failed to load.
	Failed bool `json:"failed,omitempty"`

	// Optional. True if the resource was canceled during loading.
	Canceled bool `json:"canceled,omitempty"`
}

/*
FrameResourceTree provides information about the Frame hierarchy along with their cached resources.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameResourceTree
*/
type FrameResourceTree struct {
	// Frame information for this tree item.
	Frame Frame `json:"frame"`

	// Optional. Child frames.
	ChildFrames []*FrameResourceTree `json:"childFrames,omitempty"`

	// Information about frame resources.
	Resources []*FrameResource `json:"resources"`
}

/*
FrameTree provides information about the Frame hierarchy.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameTree
*/
type FrameTree struct {
	// Frame information for this tree item.
	Frame Frame `json:"frame"`

	// Optional. Child frames.
	ChildFrames []*FrameTree `json:"childFrames,omitempty"`
}

/*
LayoutViewport defines layout viewport position and dimensions.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-LayoutViewport
*/
type LayoutViewport struct {

	// Horizontal offset relative to the document (CSS pixels).
	PageX int `json:"pageX"`

	// Vertical offset relative to the document (CSS pixels).
	PageY int `json:"pageY"`

	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth int `json:"clientWidth"`

	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight int `json:"clientHeight"`
}

/*
NavigationEntry defines a navigation history entry.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-NavigationEntry
*/
type NavigationEntry struct {
	// Unique id of the navigation history entry.
	ID int `json:"id"`

	// URL of the navigation history entry.
	URL string `json:"url"`

	// URL that the user typed in the url bar.
	UserTypedURL string `json:"userTypedURL"`

	// Title of the navigation history entry.
	Title string `json:"title"`

	// Transition type.
	TransitionType TransitionType `json:"transitionType"`
}

/*
ResourceType is the resource type as it was perceived by the rendering engine.

ALLOWED VALUES
	- Document
	- Stylesheet
	- Image
	- Media
	- Font
	- Script
	- TextTrack
	- XHR
	- Fetch
	- EventSource
	- WebSocket
	- Manifest
	- Other

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ResourceType
*/
type ResourceType string

/*
ScreencastFrameMetadata provides screencast frame metadata. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ScreencastFrameMetadata
*/
type ScreencastFrameMetadata struct {

	// Top offset in DIP.
	OffsetTop int `json:"offsetTop"`

	// Page scale factor.
	PageScaleFactor int `json:"pageScaleFactor"`

	// Device screen width in DIP.
	DeviceWidth int `json:"deviceWidth"`

	// Device screen height in DIP.
	DeviceHeight int `json:"deviceHeight"`

	// Position of horizontal scroll in CSS pixels.
	ScrollOffsetX int `json:"scrollOffsetX"`

	// Position of vertical scroll in CSS pixels.
	ScrollOffsetY int `json:"scrollOffsetY"`

	// Optional. Frame swap timestamp.
	Timestamp TimeSinceEpoch `json:"timestamp,omitempty"`
}

/*
ScriptIdentifier is the unique script identifier.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ScriptIdentifier
*/
type ScriptIdentifier string

/*
TransitionType is the transition type.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-TransitionType
*/
type TransitionType string

/*
Viewport defines the viewport for capturing screenshot.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-Viewport
*/
type Viewport struct {
	// Required. X offset in CSS pixels.
	X int `json:"x"`

	// Required. Y offset in CSS pixels.
	Y int `json:"y"`

	// Required. Rectangle width in CSS pixels
	Width int `json:"width"`

	// Required. Rectangle height in CSS pixels
	Height int `json:"height"`

	// Required. Page scale factor.
	Scale int `json:"scale"`
}

/*
VisualViewport defines visual viewport position, dimensions, and scale.

https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-VisualViewport
*/
type VisualViewport struct {
	// Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetX int `json:"offsetX"`

	// Vertical offset relative to the layout viewport (CSS pixels).
	OffsetY int `json:"offsetY"`

	// Horizontal offset relative to the document (CSS pixels).
	PageX int `json:"pageX"`

	// Vertical offset relative to the document (CSS pixels).
	PageY int `json:"pageY"`

	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth int `json:"clientWidth"`

	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight int `json:"clientHeight"`

	// Scale relative to the ideal viewport (size at width=device-width).
	Scale int `json:"scale"`
}
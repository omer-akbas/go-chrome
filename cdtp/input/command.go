package input

/*
DispatchKeyEventParams represents Input.dispatchKeyEvent parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchKeyEvent
*/
type DispatchKeyEventParams struct {
	// Type of the key event. Allowed values: keyDown, keyUp, rawKeyDown, char.
	Type string `json:"type"`

	// Optional. Bit field representing pressed modifier keys. Alt=1, Ctrl=2,
	// Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`

	// Optional. Time at which the event occurred.
	Timestamp TimeSinceEpoch `json:"timestamp,omitempty"`

	// Optional. Text as generated by processing a virtual key code with a
	// keyboard layout. Not needed for for `keyUp` and `rawKeyDown` events
	// (default: "").
	Text string `json:"text,omitempty"`

	// Optional. Text that would have been generated by the keyboard if no
	// modifiers were pressed (except for shift). Useful for shortcut
	// (accelerator) key handling (default: "").
	UnmodifiedText string `json:"unmodifiedText,omitempty"`

	// Optional. Unique key identifier (e.g., 'U+0041') (default: "").
	KeyIdentifier string `json:"keyIdentifier,omitempty"`

	// Optional. Unique DOM defined string value for each physical key (e.g.,
	// 'KeyA') (default: "").
	Code string `json:"code,omitempty"`

	// Optional. Unique DOM defined string value describing the meaning of the
	// key in the context of active modifiers, keyboard layout, etc (e.g.,
	// 'AltGr') (default: "").
	Key string `json:"key,omitempty"`

	// Optional. Windows virtual key code (default: 0).
	WindowsVirtualKeyCode int `json:"windowsVirtualKeyCode,omitempty"`

	// Optional. Native virtual key code (default: 0).
	NativeVirtualKeyCode int `json:"nativeVirtualKeyCode,omitempty"`

	// Optional. Whether the event was generated from auto repeat (default:
	// false).
	AutoRepeat bool `json:"autoRepeat,omitempty"`

	// Optional. Whether the event was generated from the keypad (default:
	// false).
	IsKeypad bool `json:"isKeypad,omitempty"`

	// Optional. Whether the event was a system key event (default: false).
	IsSystemKey bool `json:"isSystemKey,omitempty"`

	// Optional. Whether the event was from the left or right side of the
	// keyboard. 1=Left, 2=Right (default: 0).
	Location int `json:"location,omitempty"`
}

/*
DispatchKeyEventResult represents the result of calls to Input.dispatchKeyEvent.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchKeyEvent
*/
type DispatchKeyEventResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DispatchMouseEventParams represents Input.dispatchMouseEvent parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent
*/
type DispatchMouseEventParams struct {
	// Type of the mouse event.
	//
	// Allowed values:
	//	- mousePressed
	//	- mouseReleased
	//	- mouseMoved
	//	- mouseWheel
	Type string `json:"type"`

	// X coordinate of the event relative to the main frame's viewport in CSS
	// pixels.
	X int `json:"x"`

	// Y coordinate of the event relative to the main frame's viewport in CSS
	// pixels. 0 refers to the top of the viewport and Y increases as it
	// proceeds towards the bottom of the viewport.
	Y int `json:"y"`

	// Optional. Bit field representing pressed modifier keys. Alt=1, Ctrl=2,
	// Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`

	// Optional. Time at which the event occurred.
	Timestamp TimeSinceEpoch `json:"timestamp,omitempty"`

	// Optional. Mouse button (default: "none"). Allowed values: none, left,
	// middle, right.
	Button string `json:"button,omitempty"`

	// Optional. Number of times the mouse button was clicked (default: 0).
	ClickCount int `json:"clickCount,omitempty"`

	// Optional. X delta in CSS pixels for mouse wheel event (default: 0).
	DeltaX int `json:"deltaX,omitempty"`

	// Optional. Y delta in CSS pixels for mouse wheel event (default: 0).
	DeltaY int `json:"deltaY,omitempty"`
}

/*
DispatchMouseEventResult represents the result of calls to Input.dispatchMouseEvent.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent
*/
type DispatchMouseEventResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DispatchTouchEventParams represents Input.dispatchTouchEvent parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchTouchEvent
*/
type DispatchTouchEventParams struct {
	// Type of the touch event. TouchEnd and TouchCancel must not contain any
	// touch points, while TouchStart and TouchMove must contains at least one.
	// Allowed values:
	//	- touchStart
	//	- touchEnd
	//	- touchMove
	//	- touchCancel
	Type string `json:"type"`

	// Active touch points on the touch device. One event per any changed point
	// (compared to previous touch event in a sequence) is generated, emulating
	// pressing/moving/releasing points one by one.
	TouchPoints []*TouchPoint `json:"touchPoints"`

	// Optional. Bit field representing pressed modifier keys. Alt=1, Ctrl=2,
	// Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`

	// Optional. Time at which the event occurred.
	Timestamp TimeSinceEpoch `json:"timestamp,omitempty"`
}

/*
DispatchTouchEventResult represents the result of calls to Input.dispatchTouchEvent.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchTouchEvent
*/
type DispatchTouchEventResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EmulateTouchFromMouseEventParams represents Input.emulateTouchFromMouseEvent parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-emulateTouchFromMouseEvent
*/
type EmulateTouchFromMouseEventParams struct {
	// Type of the mouse event.
	//
	// Allowed values:
	//	- mousePressed
	//	- mouseReleased
	//	- mouseMoved
	//	- mouseWheel
	Type string `json:"type"`

	// X coordinate of the mouse pointer in DIP.
	X int `json:"x"`

	// Y coordinate of the mouse pointer in DIP.
	Y int `json:"y"`

	// Time at which the event occurred.
	Timestamp TimeSinceEpoch `json:"timestamp"`

	// Mouse button. Allowed values: none, left, middle, right.
	Button string `json:"button"`

	// Optional. X delta in DIP for mouse wheel event (default: 0).
	DeltaX int `json:"deltaX,omitempty"`

	// Optional. Y delta in DIP for mouse wheel event (default: 0).
	DeltaY int `json:"deltaY,omitempty"`

	// Optional. Bit field representing pressed modifier keys. Alt=1, Ctrl=2,
	// Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`

	// Optional. Number of times the mouse button was clicked (default: 0).
	ClickCount int `json:"clickCount,omitempty"`
}

/*
EmulateTouchFromMouseEventResult represents the result of calls to
Input.emulateTouchFromMouseEvent.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-emulateTouchFromMouseEvent
*/
type EmulateTouchFromMouseEventResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetIgnoreEventsParams represents Input.setIgnoreInputEvents parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-setIgnoreInputEvents
*/
type SetIgnoreEventsParams struct {
	// Ignores input events processing when set to true.
	Ignore bool `json:"ignore"`
}

/*
SetIgnoreEventsResult represents the result of calls to Input.setIgnoreEvents.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-setIgnoreEvents
*/
type SetIgnoreEventsResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SynthesizePinchGestureParams represents Input.synthesizePinchGesture parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizePinchGesture
*/
type SynthesizePinchGestureParams struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X float64 `json:"x"`

	// Y coordinate of the start of the gesture in CSS pixels.
	Y float64 `json:"y"`

	// Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
	ScaleFactor int `json:"scaleFactor"`

	// Optional. Relative pointer speed in pixels per second (default: 800).
	RelativeSpeed int `json:"relativeSpeed,omitempty"`

	// Optional. Which type of input events to be generated (default: 'default',
	// which queries the platform for the preferred input type).
	GestureSourceType *GestureSourceType `json:"gestureSourceType,omitempty"`
}

/*
SynthesizePinchGestureResult represents the result of calls to Input.synthesizePinchGesture.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizePinchGesture
*/
type SynthesizePinchGestureResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SynthesizeScrollGestureParams represents Input.synthesizeScrollGesture parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizeScrollGesture
*/
type SynthesizeScrollGestureParams struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X int `json:"x"`

	// Y coordinate of the start of the gesture in CSS pixels.
	Y int `json:"y"`

	// Optional. The distance to scroll along the X axis (positive to scroll
	// left).
	XDistance int `json:"xDistance,omitempty"`

	// Optional. The distance to scroll along the Y axis (positive to scroll up).
	YDistance int `json:"yDistance,omitempty"`

	// Optional. The number of additional pixels to scroll back along the X axis,
	// in addition to the given distance.
	XOverscroll int `json:"xOverscroll,omitempty"`

	// Optional. The number of additional pixels to scroll back along the Y axis,
	// in addition to the given distance.
	YOverscroll int `json:"yOverscroll,omitempty"`

	// Optional. Prevent fling (default: true).
	PreventFling bool `json:"preventFling,omitempty"`

	// Optional. Swipe speed in pixels per second (default: 800).
	Speed int `json:"speed,omitempty"`

	// Optional. Which type of input events to be generated (default: 'default',
	// which queries the platform for the preferred input type).
	GestureSourceType *GestureSourceType `json:"gestureSourceType,omitempty"`

	// Optional. The number of times to repeat the gesture (default: 0).
	RepeatCount int `json:"repeatCount,omitempty"`

	// Optional. The number of milliseconds delay between each repeat. (default:
	// 250).
	RepeatDelayMs int `json:"repeatDelayMs,omitempty"`

	// Optional. The name of the interaction markers to generate, if not empty
	// (default: "").
	InteractionMarkerName string `json:"interactionMarkerName,omitempty"`
}

/*
SynthesizeScrollGestureResult represents the result of calls to Input.SynthesizeScrollGesture.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-SynthesizeScrollGesture
*/
type SynthesizeScrollGestureResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SynthesizeTapGestureParams represents Input.synthesizeTapGesture parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-synthesizeTapGesture
*/
type SynthesizeTapGestureParams struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X int `json:"x"`

	// Y coordinate of the start of the gesture in CSS pixels.
	Y int `json:"y"`

	// Optional. Duration between touchdown and touchup events in ms (default:
	// 50).
	Duration int `json:"duration,omitempty"`

	// Optional. Number of times to perform the tap (e.g. 2 for double tap,
	// default: 1).
	TapCount int `json:"tapCount,omitempty"`

	// Optional. Which type of input events to be generated (default: 'default',
	// which queries the platform for the preferred input type).
	GestureSourceType *GestureSourceType `json:"gestureSourceType,omitempty"`
}

/*
SynthesizeTapGestureResult represents the result of calls to Input.SynthesizeTapGesture.

https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-SynthesizeTapGesture
*/
type SynthesizeTapGestureResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

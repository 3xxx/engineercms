import macro from 'vtk.js/Sources/macros';
import * as vtkMath from 'vtk.js/Sources/Common/Core/Math';

import Constants from 'vtk.js/Sources/Rendering/Core/RenderWindowInteractor/Constants';

const { Device, Input } = Constants;
const { vtkWarningMacro, vtkErrorMacro, normalizeWheel, vtkOnceErrorMacro } =
  macro;

// ----------------------------------------------------------------------------
// Global methods
// ----------------------------------------------------------------------------

const deviceInputMap = {
  'xr-standard': [
    Input.Trigger,
    Input.Grip,
    Input.TrackPad,
    Input.Thumbstick,
    Input.A,
    Input.B,
  ],
};

const handledEvents = [
  'StartAnimation',
  'Animation',
  'EndAnimation',
  'MouseEnter',
  'MouseLeave',
  'StartMouseMove',
  'MouseMove',
  'EndMouseMove',
  'LeftButtonPress',
  'LeftButtonRelease',
  'MiddleButtonPress',
  'MiddleButtonRelease',
  'RightButtonPress',
  'RightButtonRelease',
  'KeyPress',
  'KeyDown',
  'KeyUp',
  'StartMouseWheel',
  'MouseWheel',
  'EndMouseWheel',
  'StartPinch',
  'Pinch',
  'EndPinch',
  'StartPan',
  'Pan',
  'EndPan',
  'StartRotate',
  'Rotate',
  'EndRotate',
  'Button3D',
  'Move3D',
  'StartPointerLock',
  'EndPointerLock',
  'StartInteraction',
  'Interaction',
  'EndInteraction',
];

function preventDefault(event) {
  if (event.cancelable) {
    event.stopPropagation();
    event.preventDefault();
  }

  return false;
}

// ----------------------------------------------------------------------------
// vtkRenderWindowInteractor methods
// ----------------------------------------------------------------------------

function vtkRenderWindowInteractor(publicAPI, model) {
  // Set our className
  model.classHierarchy.push('vtkRenderWindowInteractor');

  // Initialize list of requesters
  const animationRequesters = new Set();

  // track active event listeners to handle simultaneous button tracking
  let activeListenerCount = 0;

  // Public API methods

  //----------------------------------------------------------------------
  publicAPI.start = () => {
    // Let the compositing handle the event loop if it wants to.
    // if (publicAPI.HasObserver(vtkCommand::StartEvent) && !publicAPI.HandleEventLoop) {
    //   publicAPI.invokeEvent({ type: 'StartEvent' });
    //   return;
    // }

    // As a convenience, initialize if we aren't initialized yet.
    if (!model.initialized) {
      publicAPI.initialize();
      if (!model.initialized) {
        return;
      }
    }
    // Pass execution to the subclass which will run the event loop,
    // this will not return until TerminateApp is called.
    publicAPI.startEventLoop();
  };

  //----------------------------------------------------------------------
  publicAPI.setRenderWindow = (aren) => {
    vtkErrorMacro(
      'you want to call setView(view) instead of setRenderWindow on a vtk.js interactor'
    );
  };

  //----------------------------------------------------------------------
  publicAPI.setInteractorStyle = (style) => {
    if (model.interactorStyle !== style) {
      if (model.interactorStyle != null) {
        model.interactorStyle.setInteractor(null);
      }
      model.interactorStyle = style;
      if (model.interactorStyle != null) {
        if (model.interactorStyle.getInteractor() !== publicAPI) {
          model.interactorStyle.setInteractor(publicAPI);
        }
      }
    }
  };

  //---------------------------------------------------------------------
  publicAPI.initialize = () => {
    model.initialized = true;
    publicAPI.enable();
    publicAPI.render();
  };

  publicAPI.enable = () => publicAPI.setEnabled(true);

  publicAPI.disable = () => publicAPI.setEnabled(false);

  publicAPI.startEventLoop = () => vtkWarningMacro('empty event loop');

  function updateCurrentRenderer(x, y) {
    if (!model._forcedRenderer) {
      model.currentRenderer = publicAPI.findPokedRenderer(x, y);
    }
  }

  publicAPI.getCurrentRenderer = () => {
    if (model.currentRenderer) {
      return model.currentRenderer;
    }
    updateCurrentRenderer(0, 0);
    return model.currentRenderer;
  };

  function getScreenEventPositionFor(source) {
    const bounds = model.container.getBoundingClientRect();
    const canvas = model.view.getCanvas();
    const scaleX = canvas.width / bounds.width;
    const scaleY = canvas.height / bounds.height;
    const position = {
      x: scaleX * (source.clientX - bounds.left),
      y: scaleY * (bounds.height - source.clientY + bounds.top),
      z: 0,
    };
    updateCurrentRenderer(position.x, position.y);
    return position;
  }

  function getTouchEventPositionsFor(touches) {
    const positions = {};
    for (let i = 0; i < touches.length; i++) {
      const touch = touches[i];
      positions[touch.identifier] = getScreenEventPositionFor(touch);
    }
    return positions;
  }

  function getModifierKeysFor(event) {
    return {
      controlKey: event.ctrlKey,
      altKey: event.altKey,
      shiftKey: event.shiftKey,
    };
  }

  function getKeysFor(event) {
    const modifierKeys = getModifierKeysFor(event);
    const keys = {
      key: event.key,
      keyCode: event.charCode,
      ...modifierKeys,
    };
    return keys;
  }

  function interactionRegistration(addListeners, force = false) {
    const rootElm = document;
    const method = addListeners ? 'addEventListener' : 'removeEventListener';
    const invMethod = addListeners ? 'removeEventListener' : 'addEventListener';

    if (!force && !addListeners && activeListenerCount > 0) {
      --activeListenerCount;
    }

    // only add/remove listeners when there are no registered listeners
    if (!activeListenerCount || force) {
      activeListenerCount = 0;

      if (model.container) {
        model.container[invMethod]('mousemove', publicAPI.handleMouseMove);
      }

      rootElm[method]('mouseup', publicAPI.handleMouseUp);
      rootElm[method]('mousemove', publicAPI.handleMouseMove);
      rootElm[method]('touchend', publicAPI.handleTouchEnd, false);
      rootElm[method]('touchcancel', publicAPI.handleTouchEnd, false);
      rootElm[method]('touchmove', publicAPI.handleTouchMove, false);
    }

    if (!force && addListeners) {
      ++activeListenerCount;
    }
  }

  publicAPI.bindEvents = (container) => {
    model.container = container;
    container.addEventListener('contextmenu', preventDefault);
    // container.addEventListener('click', preventDefault); // Avoid stopping event propagation
    container.addEventListener('wheel', publicAPI.handleWheel);
    container.addEventListener('DOMMouseScroll', publicAPI.handleWheel);
    container.addEventListener('mouseenter', publicAPI.handleMouseEnter);
    container.addEventListener('mouseleave', publicAPI.handleMouseLeave);
    container.addEventListener('mousemove', publicAPI.handleMouseMove);
    container.addEventListener('mousedown', publicAPI.handleMouseDown);
    document.addEventListener('keypress', publicAPI.handleKeyPress);
    document.addEventListener('keydown', publicAPI.handleKeyDown);
    document.addEventListener('keyup', publicAPI.handleKeyUp);

    document.addEventListener(
      'pointerlockchange',
      publicAPI.handlePointerLockChange
    );

    container.addEventListener('touchstart', publicAPI.handleTouchStart, false);
  };

  publicAPI.unbindEvents = () => {
    // force unbinding listeners
    interactionRegistration(false, true);
    model.container.removeEventListener('contextmenu', preventDefault);
    // model.container.removeEventListener('click', preventDefault); // Avoid stopping event propagation
    model.container.removeEventListener('wheel', publicAPI.handleWheel);
    model.container.removeEventListener(
      'DOMMouseScroll',
      publicAPI.handleWheel
    );
    model.container.removeEventListener(
      'mouseenter',
      publicAPI.handleMouseEnter
    );
    model.container.removeEventListener(
      'mouseleave',
      publicAPI.handleMouseLeave
    );
    model.container.removeEventListener('mousemove', publicAPI.handleMouseMove);
    model.container.removeEventListener('mousedown', publicAPI.handleMouseDown);
    document.removeEventListener('keypress', publicAPI.handleKeyPress);
    document.removeEventListener('keydown', publicAPI.handleKeyDown);
    document.removeEventListener('keyup', publicAPI.handleKeyUp);
    document.removeEventListener(
      'pointerlockchange',
      publicAPI.handlePointerLockChange
    );
    model.container.removeEventListener(
      'touchstart',
      publicAPI.handleTouchStart
    );
    model.container = null;
  };

  publicAPI.handleKeyPress = (event) => {
    const data = getKeysFor(event);
    publicAPI.keyPressEvent(data);
  };

  publicAPI.handleKeyDown = (event) => {
    const data = getKeysFor(event);
    publicAPI.keyDownEvent(data);
  };

  publicAPI.handleKeyUp = (event) => {
    const data = getKeysFor(event);
    publicAPI.keyUpEvent(data);
  };

  publicAPI.handleMouseDown = (event) => {
    if (event.button > 2) {
      // ignore events from extra mouse buttons such as `back` and `forward`
      return;
    }

    interactionRegistration(true);
    preventDefault(event);

    const callData = {
      ...getModifierKeysFor(event),
      position: getScreenEventPositionFor(event),
    };
    switch (event.button) {
      case 0:
        publicAPI.leftButtonPressEvent(callData);
        break;
      case 1:
        publicAPI.middleButtonPressEvent(callData);
        break;
      case 2:
        publicAPI.rightButtonPressEvent(callData);
        break;
      default:
        vtkErrorMacro(`Unknown mouse button pressed: ${event.button}`);
        break;
    }
  };

  //----------------------------------------------------------------------
  publicAPI.requestPointerLock = () => {
    const canvas = publicAPI.getView().getCanvas();
    canvas.requestPointerLock();
  };

  //----------------------------------------------------------------------
  publicAPI.exitPointerLock = () => document.exitPointerLock();

  //----------------------------------------------------------------------
  publicAPI.isPointerLocked = () => !!document.pointerLockElement;

  //----------------------------------------------------------------------
  publicAPI.handlePointerLockChange = () => {
    if (publicAPI.isPointerLocked()) {
      publicAPI.startPointerLockEvent();
    } else {
      publicAPI.endPointerLockEvent();
    }
  };

  //----------------------------------------------------------------------
  function forceRender() {
    if (model.view && model.enabled && model.enableRender) {
      model.inRender = true;
      model.view.traverseAllPasses();
      model.inRender = false;
    }
    // outside the above test so that third-party code can redirect
    // the render to the appropriate class
    publicAPI.invokeRenderEvent();
  }

  publicAPI.requestAnimation = (requestor) => {
    if (requestor === undefined) {
      vtkErrorMacro(`undefined requester, can not start animating`);
      return;
    }
    if (animationRequesters.has(requestor)) {
      vtkWarningMacro(`requester is already registered for animating`);
      return;
    }
    animationRequesters.add(requestor);
    if (animationRequesters.size === 1 && !model.xrAnimation) {
      model.lastFrameTime = 0.1;
      model.lastFrameStart = Date.now();
      model.animationRequest = requestAnimationFrame(publicAPI.handleAnimation);
      publicAPI.startAnimationEvent();
    }
  };

  publicAPI.isAnimating = () =>
    model.xrAnimation || model.animationRequest !== null;

  publicAPI.cancelAnimation = (requestor, skipWarning = false) => {
    if (!animationRequesters.has(requestor)) {
      if (!skipWarning) {
        const requestStr =
          requestor && requestor.getClassName
            ? requestor.getClassName()
            : requestor;
        vtkWarningMacro(`${requestStr} did not request an animation`);
      }

      return;
    }
    animationRequesters.delete(requestor);
    if (model.animationRequest && animationRequesters.size === 0) {
      cancelAnimationFrame(model.animationRequest);
      model.animationRequest = null;
      publicAPI.endAnimationEvent();
      publicAPI.render();
    }
  };

  publicAPI.switchToXRAnimation = () => {
    // cancel existing animation if any
    if (model.animationRequest) {
      cancelAnimationFrame(model.animationRequest);
      model.animationRequest = null;
    }
    model.xrAnimation = true;
  };

  publicAPI.returnFromXRAnimation = () => {
    model.xrAnimation = false;
    if (animationRequesters.size !== 0) {
      model.FrameTime = -1;
      model.animationRequest = requestAnimationFrame(publicAPI.handleAnimation);
    }
  };

  publicAPI.updateXRGamepads = (xrSession, xrFrame, xrRefSpace) => {
    // watch for when buttons change state and fire events
    xrSession.inputSources.forEach((inputSource) => {
      const pose = xrFrame.getPose(inputSource.gripSpace, xrRefSpace);
      const gp = inputSource.gamepad;
      const hand = inputSource.handedness;
      if (gp) {
        if (!(gp.index in model.lastGamepadValues)) {
          model.lastGamepadValues[gp.index] = {
            left: { buttons: {} },
            right: { buttons: {} },
          };
        }
        for (let b = 0; b < gp.buttons.length; ++b) {
          if (!(b in model.lastGamepadValues[gp.index][hand].buttons)) {
            model.lastGamepadValues[gp.index][hand].buttons[b] = false;
          }
          if (
            model.lastGamepadValues[gp.index][hand].buttons[b] !==
            gp.buttons[b].pressed
          ) {
            publicAPI.button3DEvent({
              gamepad: gp,
              position: pose.transform.position,
              orientation: pose.transform.orientation,
              pressed: gp.buttons[b].pressed,
              device:
                inputSource.handedness === 'left'
                  ? Device.LeftController
                  : Device.RightController,
              input:
                deviceInputMap[gp.mapping] && deviceInputMap[gp.mapping][b]
                  ? deviceInputMap[gp.mapping][b]
                  : Input.Trigger,
            });
            model.lastGamepadValues[gp.index][hand].buttons[b] =
              gp.buttons[b].pressed;
          }
          if (model.lastGamepadValues[gp.index][hand].buttons[b]) {
            publicAPI.move3DEvent({
              gamepad: gp,
              position: pose.transform.position,
              orientation: pose.transform.orientation,
              device:
                inputSource.handedness === 'left'
                  ? Device.LeftController
                  : Device.RightController,
            });
          }
        }
      }
    });
  };

  publicAPI.handleMouseMove = (event) => {
    // Do not consume event for move
    // preventDefault(event);

    const callData = {
      ...getModifierKeysFor(event),
      position: getScreenEventPositionFor(event),
    };

    if (model.moveTimeoutID === 0) {
      publicAPI.startMouseMoveEvent(callData);
    } else {
      publicAPI.mouseMoveEvent(callData);
      clearTimeout(model.moveTimeoutID);
    }

    // start a timer to keep us animating while we get mouse move events
    model.moveTimeoutID = setTimeout(() => {
      publicAPI.endMouseMoveEvent();
      model.moveTimeoutID = 0;
    }, 200);
  };

  publicAPI.handleAnimation = () => {
    const currTime = Date.now();
    if (model.FrameTime === -1.0) {
      model.lastFrameTime = 0.1;
    } else {
      model.lastFrameTime = (currTime - model.lastFrameStart) / 1000.0;
    }
    model.lastFrameTime = Math.max(0.01, model.lastFrameTime);
    model.lastFrameStart = currTime;
    publicAPI.animationEvent();
    forceRender();
    model.animationRequest = requestAnimationFrame(publicAPI.handleAnimation);
  };

  publicAPI.handleWheel = (event) => {
    preventDefault(event);

    /**
     * wheel event values can vary significantly across browsers, platforms
     * and devices [1]. `normalizeWheel` uses facebook's solution from their
     * fixed-data-table repository [2].
     *
     * [1] https://developer.mozilla.org/en-US/docs/Web/Events/mousewheel
     * [2] https://github.com/facebookarchive/fixed-data-table/blob/master/src/vendor_upstream/dom/normalizeWheel.js
     *
     * This code will return an object with properties:
     *
     *   spinX   -- normalized spin speed (use for zoom) - x plane
     *   spinY   -- " - y plane
     *   pixelX  -- normalized distance (to pixels) - x plane
     *   pixelY  -- " - y plane
     *
     */
    const callData = {
      ...normalizeWheel(event),
      ...getModifierKeysFor(event),
      position: getScreenEventPositionFor(event),
    };

    if (model.wheelTimeoutID === 0) {
      publicAPI.startMouseWheelEvent(callData);
    } else {
      publicAPI.mouseWheelEvent(callData);
      clearTimeout(model.wheelTimeoutID);
    }

    // start a timer to keep us animating while we get wheel events
    model.wheelTimeoutID = setTimeout(() => {
      publicAPI.endMouseWheelEvent();
      model.wheelTimeoutID = 0;
    }, 200);
  };

  publicAPI.handleMouseEnter = (event) => {
    const callData = {
      ...getModifierKeysFor(event),
      position: getScreenEventPositionFor(event),
    };
    publicAPI.mouseEnterEvent(callData);
  };

  publicAPI.handleMouseLeave = (event) => {
    const callData = {
      ...getModifierKeysFor(event),
      position: getScreenEventPositionFor(event),
    };

    publicAPI.mouseLeaveEvent(callData);
  };

  publicAPI.handleMouseUp = (event) => {
    interactionRegistration(false);
    preventDefault(event);

    const callData = {
      ...getModifierKeysFor(event),
      position: getScreenEventPositionFor(event),
    };
    switch (event.button) {
      case 0:
        publicAPI.leftButtonReleaseEvent(callData);
        break;
      case 1:
        publicAPI.middleButtonReleaseEvent(callData);
        break;
      case 2:
        publicAPI.rightButtonReleaseEvent(callData);
        break;
      default:
        vtkErrorMacro(`Unknown mouse button released: ${event.button}`);
        break;
    }
  };

  publicAPI.handleTouchStart = (event) => {
    interactionRegistration(true);
    preventDefault(event);

    // If multitouch
    if (model.recognizeGestures && event.touches.length > 1) {
      const positions = getTouchEventPositionsFor(event.touches);
      // did we just transition to multitouch?
      if (event.touches.length === 2) {
        const touch = event.touches[0];
        const callData = {
          position: getScreenEventPositionFor(touch),
          shiftKey: false,
          altKey: false,
          controlKey: false,
        };
        publicAPI.leftButtonReleaseEvent(callData);
      }
      // handle the gesture
      publicAPI.recognizeGesture('TouchStart', positions);
    } else {
      const touch = event.touches[0];
      const callData = {
        position: getScreenEventPositionFor(touch),
        shiftKey: false,
        altKey: false,
        controlKey: false,
      };
      publicAPI.leftButtonPressEvent(callData);
    }
  };

  publicAPI.handleTouchMove = (event) => {
    preventDefault(event);

    if (model.recognizeGestures && event.touches.length > 1) {
      const positions = getTouchEventPositionsFor(event.touches);
      publicAPI.recognizeGesture('TouchMove', positions);
    } else {
      const touch = event.touches[0];
      const callData = {
        position: getScreenEventPositionFor(touch),
        shiftKey: false,
        altKey: false,
        controlKey: false,
      };
      publicAPI.mouseMoveEvent(callData);
    }
  };

  publicAPI.handleTouchEnd = (event) => {
    preventDefault(event);

    if (model.recognizeGestures) {
      // No more fingers down
      if (event.touches.length === 0) {
        // If just one finger released, consider as left button
        if (event.changedTouches.length === 1) {
          const touch = event.changedTouches[0];
          const callData = {
            position: getScreenEventPositionFor(touch),
            shiftKey: false,
            altKey: false,
            controlKey: false,
          };
          publicAPI.leftButtonReleaseEvent(callData);
          interactionRegistration(false);
        } else {
          // If more than one finger released, recognize touchend
          const positions = getTouchEventPositionsFor(event.changedTouches);
          publicAPI.recognizeGesture('TouchEnd', positions);
          interactionRegistration(false);
        }
      } else if (event.touches.length === 1) {
        // If one finger left, end touch and start button press
        const positions = getTouchEventPositionsFor(event.changedTouches);
        publicAPI.recognizeGesture('TouchEnd', positions);
        const touch = event.touches[0];
        const callData = {
          position: getScreenEventPositionFor(touch),
          shiftKey: false,
          altKey: false,
          controlKey: false,
        };
        publicAPI.leftButtonPressEvent(callData);
      } else {
        // If more than one finger left, keep touch move
        const positions = getTouchEventPositionsFor(event.touches);
        publicAPI.recognizeGesture('TouchMove', positions);
      }
    } else {
      const touch = event.changedTouches[0];
      const callData = {
        position: getScreenEventPositionFor(touch),
        shiftKey: false,
        altKey: false,
        controlKey: false,
      };
      publicAPI.leftButtonReleaseEvent(callData);
      interactionRegistration(false);
    }
  };

  publicAPI.setView = (val) => {
    if (model.view === val) {
      return;
    }
    model.view = val;
    model.view.getRenderable().setInteractor(publicAPI);
    publicAPI.modified();
  };

  publicAPI.getFirstRenderer = () =>
    model.view.getRenderable().getRenderersByReference()[0];

  publicAPI.findPokedRenderer = (x = 0, y = 0) => {
    if (!model.view) {
      return null;
    }
    // The original order of renderers needs to remain as
    // the first one is the one we want to manipulate the camera on.
    const rc = model.view.getRenderable().getRenderers();
    rc.sort((a, b) => a.getLayer() - b.getLayer());
    let interactiveren = null;
    let viewportren = null;
    let currentRenderer = null;

    let count = rc.length;
    while (count--) {
      const aren = rc[count];
      if (model.view.isInViewport(x, y, aren) && aren.getInteractive()) {
        currentRenderer = aren;
        break;
      }

      if (interactiveren === null && aren.getInteractive()) {
        // Save this renderer in case we can't find one in the viewport that
        // is interactive.
        interactiveren = aren;
      }
      if (viewportren === null && model.view.isInViewport(x, y, aren)) {
        // Save this renderer in case we can't find one in the viewport that
        // is interactive.
        viewportren = aren;
      }
    }

    // We must have a value.  If we found an interactive renderer before, that's
    // better than a non-interactive renderer.
    if (currentRenderer === null) {
      currentRenderer = interactiveren;
    }

    // We must have a value.  If we found a renderer that is in the viewport,
    // that is better than any old viewport (but not as good as an interactive
    // one).
    if (currentRenderer === null) {
      currentRenderer = viewportren;
    }

    // We must have a value - take anything.
    if (currentRenderer == null) {
      currentRenderer = rc[0];
    }

    return currentRenderer;
  };

  // only render if we are not animating. If we are animating
  // then renders will happen naturally anyhow and we definitely
  // do not want extra renders as the make the apparent interaction
  // rate slower.
  publicAPI.render = () => {
    if (!publicAPI.isAnimating() && !model.inRender) {
      forceRender();
    }
  };

  // create the generic Event methods
  handledEvents.forEach((eventName) => {
    const lowerFirst = eventName.charAt(0).toLowerCase() + eventName.slice(1);
    publicAPI[`${lowerFirst}Event`] = (arg) => {
      // Check that interactor enabled
      if (!model.enabled) {
        return;
      }

      // Check that a poked renderer exists
      const renderer = publicAPI.getCurrentRenderer();
      if (!renderer) {
        vtkOnceErrorMacro(`
          Can not forward events without a current renderer on the interactor.
        `);
        return;
      }

      // Pass the eventName and the poked renderer
      const callData = {
        type: eventName,
        pokedRenderer: model.currentRenderer,
        firstRenderer: publicAPI.getFirstRenderer(),
        // Add the arguments to the call data
        ...arg,
      };

      // Call invoke
      publicAPI[`invoke${eventName}`](callData);
    };
  });

  // we know we are in multitouch now, so start recognizing
  publicAPI.recognizeGesture = (event, positions) => {
    // more than two pointers we ignore
    if (Object.keys(positions).length > 2) {
      return;
    }

    if (!model.startingEventPositions) {
      model.startingEventPositions = {};
    }

    // store the initial positions
    if (event === 'TouchStart') {
      Object.keys(positions).forEach((key) => {
        model.startingEventPositions[key] = positions[key];
      });
      // we do not know what the gesture is yet
      model.currentGesture = 'Start';
      return;
    }

    // end the gesture if needed
    if (event === 'TouchEnd') {
      if (model.currentGesture === 'Pinch') {
        publicAPI.render();
        publicAPI.endPinchEvent();
      }
      if (model.currentGesture === 'Rotate') {
        publicAPI.render();
        publicAPI.endRotateEvent();
      }
      if (model.currentGesture === 'Pan') {
        publicAPI.render();
        publicAPI.endPanEvent();
      }
      model.currentGesture = 'Start';
      model.startingEventPositions = {};
      return;
    }

    // what are the two pointers we are working with
    let count = 0;
    const posVals = [];
    const startVals = [];
    Object.keys(positions).forEach((key) => {
      posVals[count] = positions[key];
      startVals[count] = model.startingEventPositions[key];
      count++;
    });

    // The meat of the algorithm
    // on move events we analyze them to determine what type
    // of movement it is and then deal with it.
    // calculate the distances
    const originalDistance = Math.sqrt(
      (startVals[0].x - startVals[1].x) * (startVals[0].x - startVals[1].x) +
        (startVals[0].y - startVals[1].y) * (startVals[0].y - startVals[1].y)
    );
    const newDistance = Math.sqrt(
      (posVals[0].x - posVals[1].x) * (posVals[0].x - posVals[1].x) +
        (posVals[0].y - posVals[1].y) * (posVals[0].y - posVals[1].y)
    );

    // calculate rotations
    let originalAngle = vtkMath.degreesFromRadians(
      Math.atan2(
        startVals[1].y - startVals[0].y,
        startVals[1].x - startVals[0].x
      )
    );
    let newAngle = vtkMath.degreesFromRadians(
      Math.atan2(posVals[1].y - posVals[0].y, posVals[1].x - posVals[0].x)
    );

    // angles are cyclic so watch for that, 1 and 359 are only 2 apart :)
    let angleDeviation = newAngle - originalAngle;
    newAngle = newAngle + 180.0 >= 360.0 ? newAngle - 180.0 : newAngle + 180.0;
    originalAngle =
      originalAngle + 180.0 >= 360.0
        ? originalAngle - 180.0
        : originalAngle + 180.0;
    if (Math.abs(newAngle - originalAngle) < Math.abs(angleDeviation)) {
      angleDeviation = newAngle - originalAngle;
    }

    // calculate the translations
    const trans = [];
    trans[0] =
      (posVals[0].x - startVals[0].x + posVals[1].x - startVals[1].x) / 2.0;
    trans[1] =
      (posVals[0].y - startVals[0].y + posVals[1].y - startVals[1].y) / 2.0;

    if (event === 'TouchMove') {
      // OK we want to
      // - immediately respond to the user
      // - allow the user to zoom without panning (saves focal point)
      // - allow the user to rotate without panning (saves focal point)

      // do we know what gesture we are doing yet? If not
      // see if we can figure it out
      if (model.currentGesture === 'Start') {
        // pinch is a move to/from the center point
        // rotate is a move along the circumference
        // pan is a move of the center point
        // compute the distance along each of these axes in pixels
        // the first to break thresh wins
        let thresh =
          0.01 *
          Math.sqrt(
            model.container.clientWidth * model.container.clientWidth +
              model.container.clientHeight * model.container.clientHeight
          );
        if (thresh < 15.0) {
          thresh = 15.0;
        }
        const pinchDistance = Math.abs(newDistance - originalDistance);
        const rotateDistance =
          (newDistance * 3.1415926 * Math.abs(angleDeviation)) / 360.0;
        const panDistance = Math.sqrt(
          trans[0] * trans[0] + trans[1] * trans[1]
        );
        if (
          pinchDistance > thresh &&
          pinchDistance > rotateDistance &&
          pinchDistance > panDistance
        ) {
          model.currentGesture = 'Pinch';
          const callData = {
            scale: 1.0,
            touches: positions,
          };
          publicAPI.startPinchEvent(callData);
        } else if (rotateDistance > thresh && rotateDistance > panDistance) {
          model.currentGesture = 'Rotate';
          const callData = {
            rotation: 0.0,
            touches: positions,
          };
          publicAPI.startRotateEvent(callData);
        } else if (panDistance > thresh) {
          model.currentGesture = 'Pan';
          const callData = {
            translation: [0, 0],
            touches: positions,
          };
          publicAPI.startPanEvent(callData);
        }
      } else {
        // if we have found a specific type of movement then
        // handle it
        if (model.currentGesture === 'Rotate') {
          const callData = {
            rotation: angleDeviation,
            touches: positions,
          };
          publicAPI.rotateEvent(callData);
        }

        if (model.currentGesture === 'Pinch') {
          const callData = {
            scale: newDistance / originalDistance,
            touches: positions,
          };
          publicAPI.pinchEvent(callData);
        }

        if (model.currentGesture === 'Pan') {
          const callData = {
            translation: trans,
            touches: positions,
          };
          publicAPI.panEvent(callData);
        }
      }
    }
  };

  publicAPI.handleVisibilityChange = () => {
    model.lastFrameStart = Date.now();
  };

  publicAPI.setCurrentRenderer = (r) => {
    model._forcedRenderer = !!r;
    model.currentRenderer = r;
  };

  // Stop animating if the renderWindowInteractor is deleted.
  const superDelete = publicAPI.delete;
  publicAPI.delete = () => {
    while (animationRequesters.size) {
      publicAPI.cancelAnimation(animationRequesters.values().next().value);
    }
    if (typeof document.hidden !== 'undefined') {
      document.removeEventListener(
        'visibilitychange',
        publicAPI.handleVisibilityChange
      );
    }
    superDelete();
  };

  // Use the Page Visibility API to detect when we switch away from or back to
  // this tab, and reset the lastFrameStart. When tabs are not active, browsers
  // will stop calling requestAnimationFrame callbacks.
  if (typeof document.hidden !== 'undefined') {
    document.addEventListener(
      'visibilitychange',
      publicAPI.handleVisibilityChange,
      false
    );
  }
}

// ----------------------------------------------------------------------------
// Object factory
// ----------------------------------------------------------------------------

const DEFAULT_VALUES = {
  renderWindow: null,
  interactorStyle: null,
  picker: null,
  pickingManager: null,
  initialized: false,
  enabled: false,
  enableRender: true,
  currentRenderer: null,
  lightFollowCamera: true,
  desiredUpdateRate: 30.0,
  stillUpdateRate: 2.0,
  container: null,
  view: null,
  recognizeGestures: true,
  currentGesture: 'Start',
  animationRequest: null,
  lastFrameTime: 0.1,
  wheelTimeoutID: 0,
  moveTimeoutID: 0,
  lastGamepadValues: {},
};

// ----------------------------------------------------------------------------

export function extend(publicAPI, model, initialValues = {}) {
  Object.assign(model, DEFAULT_VALUES, initialValues);

  // Object methods
  macro.obj(publicAPI, model);

  macro.event(publicAPI, model, 'RenderEvent');
  handledEvents.forEach((eventName) =>
    macro.event(publicAPI, model, eventName)
  );

  // Create get-only macros
  macro.get(publicAPI, model, [
    'initialized',
    'container',
    'interactorStyle',
    'lastFrameTime',
    'view',
  ]);

  // Create get-set macros
  macro.setGet(publicAPI, model, [
    'lightFollowCamera',
    'enabled',
    'enableRender',
    'recognizeGestures',
    'desiredUpdateRate',
    'stillUpdateRate',
    'picker',
  ]);

  // For more macro methods, see "Sources/macros.js"

  // Object specific methods
  vtkRenderWindowInteractor(publicAPI, model);
}

// ----------------------------------------------------------------------------

export const newInstance = macro.newInstance(
  extend,
  'vtkRenderWindowInteractor'
);

// ----------------------------------------------------------------------------

export default { newInstance, extend, handledEvents, ...Constants };

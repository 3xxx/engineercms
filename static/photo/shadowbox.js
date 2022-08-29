/*!
 * Shadowbox, version 4.0.0 <http://shadowbox-js.com/>
 * Copyright 2014 Michael Jackson
 */

(function (global) {

  var documentElement = document.documentElement;

  // Detect support for opacity.
  var supportsOpacity = "opacity" in documentElement.style && typeof documentElement.style.opacity === "string";

  // Detect support for fixed positioning.
  var fixedDiv = document.createElement("div");
  fixedDiv.style.position = "fixed";
  fixedDiv.style.margin = 0;
  fixedDiv.style.top = "20px";
  documentElement.appendChild(fixedDiv, documentElement.firstChild);
  var supportsFixed = (fixedDiv.offsetTop == 20);
  documentElement.removeChild(fixedDiv);

  // Detect touch-based devices.
  var supportsTouch = ("createTouch" in document);

  /**
   * The current version of Shadowbox.
   */
  shadowbox.version = "4.0.0";

  shadowbox.guid = 1;

  shadowbox.K = function () {
    return this;
  };

  /**
   * The default set of options.
   */
  shadowbox.options = {

    // Animate height/width transitions.
    animate: true,

    // Automatically close when done playing movies.
    autoClose: false,

    // Able to navigate from one end of a gallery to the other (i.e. from
    // last item to first or vice versa) by choosing next/previous?
    continuous: false,

    // Easing function for animations. Based on a cubic polynomial.
    ease: function (state) {
      return 1 + Math.pow(state - 1, 3);
    },

    // Enable control of Shadowbox via the keyboard?
    enableKeys: !supportsTouch,

    // The space to maintain around the edge of Shadowbox at all times.
    margin: 40,

    // A hook function that is called when closing.
    onClose: shadowbox.K,

    // A hook function that is called when a player is finished loading and
    // all display transitions are complete. Receives the player object as
    // its only argument.
    onDone: shadowbox.K,

    // A hook function that is called when opening.
    onOpen: shadowbox.K,

    // A hook function that is called when a player is ready to be
    // displayed. Receives the player object as its only argument.
    onShow: shadowbox.K,

    // Background color for the overlay.
    overlayColor: "black",

    // Opacity for the overlay.
    overlayOpacity: 0.6,

    // The index in the current gallery at which to start when first opening.
    startIndex: 0

  };

  /**
   * A map of file extensions to the player class that should be used to play
   * files with that extension.
   */
  shadowbox.players = {};

  /**
   * Registers the given player class to be used with the given file
   * extensions.
   *
   *   shadowbox.registerPlayer(shadowbox.VideoPlayer, "mov");
   *   shadowbox.registerPlayer(shadowbox.PhotoPlayer, [ "jpg", "jpeg" ]);
   */
  shadowbox.registerPlayer = function (playerClass, extensions) {
    extensions = extensions || [];

    if (!isArray(extensions))
      extensions = [ extensions ];

    forEach(extensions, function (extension) {
      shadowbox.players[extension] = playerClass;
    });
  };

  // Cache references to oft-used DOM elements.
  var containerElement, overlayElement, wrapperElement, bodyElement, contentElement, coverElement;

  /**
   * Appends Shadowbox to the DOM and initializes DOM references.
   */
  function initialize() {
    if (containerElement)
      return; // Don't initialize twice!

    // The Shadowbox markup:
    //
    // <div id="shadowbox">
    //   <div id="sb-overlay"></div>
    //   <div id="sb-wrapper">
    //     <div id="sb-body">
    //       <div id="sb-content"></div>
    //       <div id="sb-cover"></div>
    //     </div>
    //     <div id="sb-close"></div>
    //     <div id="sb-next"></div>
    //     <div id="sb-prev"></div>
    //   </div>
    // </div>

    containerElement = makeDom("div", { id: "shadowbox" });
    overlayElement = makeDom("div", { id: "sb-overlay" });
    wrapperElement = makeDom("div", { id: "sb-wrapper" });
    bodyElement = makeDom("div", { id: "sb-body" });
    contentElement = makeDom("div", { id: "sb-content" });
    coverElement = makeDom("div", { id: "sb-cover" });
    var closeElement = makeDom("div", { id: "sb-close" });
    var nextElement = makeDom("div", { id: "sb-next" });
    var previousElement = makeDom("div", { id: "sb-prev" });

    // Append #shadowbox to the DOM.
    makeDom(document.body, [
      makeDom(containerElement, [
        overlayElement,
        makeDom(wrapperElement, [
          makeDom(bodyElement, [ contentElement, coverElement ]),
          closeElement,
          nextElement,
          previousElement
        ])
      ])
    ]);

    // Use an absolutely positioned container in browsers that don't
    // support fixed positioning.
    if (!supportsFixed)
      setStyle(containerElement, "position", "absolute");

    // Setup a click listener on the overlay to close Shadowbox.
    addEvent(overlayElement, "click", shadowbox.close);

    // Setup callbacks on navigation elements.
    addEvent(closeElement, "click", cancel(shadowbox.close));
    addEvent(nextElement, "click", cancel(shadowbox.showNext));
    addEvent(previousElement, "click", cancel(shadowbox.showPrevious));
  }

  var currentIndex = -1,
      currentGallery,
      currentPlayer,
      currentOptions;

  /**
   * Opens an object (or an array of objects) in Shadowbox. Takes options as
   * the second argument.
   *
   *   shadowbox("myphoto.jpg");
   *   shadowbox([ "myphoto1.jpg", "myphoto2.jpg" ]);
   *   shadowbox([ "myphoto1.jpg", "myphoto2.jpg" ], {
   *     animate:         false,
   *     overlayColor:    "white",
   *     overlayOpacity:  0.8
   *   });
   *
   * Options may be any of shadowbox.options. Returns the number of objects
   * that were able to be opened.
   */
  function shadowbox(objects, options) {
    if (typeof options === 'number')
      options = { startIndex: options };

    if (!isArray(objects))
      objects = [ objects ];

    currentOptions = mergeProperties({}, shadowbox.options);

    if (options)
      mergeProperties(currentOptions, options);

    currentGallery = [];

    // Normalize into player objects and append them to the gallery.
    var startIndex = currentOptions.startIndex;
    forEach(objects, function (object, index) {
      var player = shadowbox.makePlayer(object);

      if (player) {
        currentGallery.push(player);
      } else {
        if (index < startIndex) {
          startIndex -= 1;
        } else if (index === startIndex) {
          startIndex = 0;
        }
      }
    });

    // Display the first item in the gallery, if there's anything left.
    if (currentGallery.length > 0) {
      if (currentIndex == -1) {
        initialize();

        if (isFunction(currentOptions.onOpen))
          currentOptions.onOpen();

        setStyle(containerElement, "display", "block");
        setContainerPosition();
        setContainerSize();
        toggleTroubleElements(0);
        setStyle(overlayElement, "backgroundColor", currentOptions.overlayColor);
        setStyle(overlayElement, "opacity", 0);
        setStyle(containerElement, "visibility", "visible");

        animateStyle(overlayElement, "opacity", currentOptions.overlayOpacity, 0.35, function () {
          setWrapperSize({ width: 340, height: 200 });
          setStyle(wrapperElement, "visibility", "visible");
          shadowbox.show(startIndex);
        });
      } else {
        shadowbox.show(startIndex);
      }
    }

    return currentGallery.length;
  }

  // Alias.
  shadowbox.open = shadowbox;

  /**
   * Displays the gallery item at the given index in Shadowbox. Assumes that
   * Shadowbox is already initialized and open.
   */
  shadowbox.show = function (index) {
    // Guard against invalid indices and no-ops.
    if (index < 0 || !currentGallery[index] || currentIndex === index)
      return;

    toggleControls(0);
    toggleWindowHandlers(0);
    toggleMouseMoveHandler(0);
    toggleKeyDownHandler(0);

    setStyle(coverElement, "display", "block");
    setStyle(coverElement, "opacity", 1);

    if (currentPlayer)
      currentPlayer.remove();

    currentIndex = index;
    currentPlayer = currentGallery[currentIndex];

    function playerIsReady() {
      return !currentPlayer || currentPlayer.isReady !== false;
    }

    waitUntil(playerIsReady, function () {
      if (!currentPlayer)
        return; // Shadowbox was closed.

      if (isFunction(currentOptions.onShow))
        currentOptions.onShow(currentPlayer);

      var size = getWrapperSize();
      var fromWidth = parseInt(getStyle(wrapperElement, "width")) || 0,
          fromHeight = parseInt(getStyle(wrapperElement, "height")) || 0,
          toWidth = size.width,
          toHeight = size.height,
          changeWidth = toWidth - fromWidth,
          changeHeight = toHeight - fromHeight;

      function frameHandler(value) {
        if (!currentPlayer)
          return false; // Shadowbox was closed, cancel the animation.

        setWrapperSize({
          width: fromWidth + (changeWidth * value),
          height: fromHeight + (changeHeight * value)
        });
      }

      // Open to the correct dimensions. Use the low-level animation
      // primitive to make this transition as smooth as possible.
      animate(0, 1, 0.5, frameHandler, function () {
        if (currentPlayer) {
          currentPlayer.injectInto(contentElement);

          if (currentPlayer.fadeCover) {
            animateStyle(coverElement, "opacity", 0, 0.5, finishShow);
          } else {
            finishShow();
          }
        }
      });
    });
  };

  function finishShow() {
    if (currentPlayer) {
      setStyle(coverElement, "display", "none");

      toggleWindowHandlers(1);
      toggleMouseMoveHandler(1);
      toggleKeyDownHandler(1);

      if (isFunction(currentOptions.onDone))
        currentOptions.onDone(currentPlayer);
    }
  }

  /**
   * Opens the previous item in the gallery.
   */
  shadowbox.showPrevious = function () {
    shadowbox.show(getPreviousIndex());
  };

  /**
   * Gets the index of the previous item in the gallery, -1 if there is none.
   */
  function getPreviousIndex() {
    if (currentIndex === 0)
      return currentOptions.continuous ? (currentGallery.length - 1) : -1;

    return currentIndex - 1;
  }

  /**
   * Opens the next item in the gallery.
   */
  shadowbox.showNext = function () {
    shadowbox.show(getNextIndex());
  };

  /**
   * Gets the index of the next item in the gallery, -1 if there is none.
   */
  function getNextIndex() {
    if (currentIndex === currentGallery.length - 1)
      return (currentOptions.continuous && currentIndex !== 0) ? 0 : -1;

    return currentIndex + 1;
  }

  /**
   * Closes Shadowbox immediately.
   */
  shadowbox.close = function () {
    if (shadowbox.isOpen()) {
      currentIndex = -1;
      currentPlayer = null;

      setStyle(wrapperElement, "visibility", "hidden");
      setStyle(coverElement, "opacity", 1);
      contentElement.innerHTML = "";

      toggleControls(0);
      toggleWindowHandlers(0);
      toggleMouseMoveHandler(0);
      toggleKeyDownHandler(0);

      animateStyle(overlayElement, "opacity", 0, 0.5, function () {
        setStyle(containerElement, "visibility", "hidden");
        setStyle(containerElement, "display", "none");
        toggleTroubleElements(1);

        if (isFunction(currentOptions.onClose))
          currentOptions.onClose();
      });
    }
  };

  /**
   * Returns true if Shadowbox is currently open.
   */
  shadowbox.isOpen = function () {
    return currentIndex !== -1;
  };

  /**
   * Gets the current player instance.
   */
  shadowbox.getPlayer = function () {
    return currentPlayer;
  };

  /**
   * Gets the size that should be used for the wrapper element. Should be
   * called when Shadowbox is open and has a player that is ready.
   */
  function getWrapperSize() {
    var margin = Math.max(currentOptions.margin, 20); // Minimum 20px margin.

    return constrainSize(currentPlayer.width, currentPlayer.height,
      overlayElement.offsetWidth, overlayElement.offsetHeight, margin);
  }

  /**
   * Sets the size and position of the wrapper.
   */
  function setWrapperSize(size) {
    setStyle(wrapperElement, "width", size.width + "px");
    setStyle(wrapperElement, "marginLeft", (-size.width / 2) + "px");
    setStyle(wrapperElement, "height", size.height + "px");
    setStyle(wrapperElement, "marginTop", (-size.height / 2) + "px");
  }

  /**
   * Scales the given width and height to be within the bounds of the given
   * maximum width and height, allowing for margin. Returns an array of the
   * constrained [width, height].
   */
  function constrainSize(width, height, maxWidth, maxHeight, margin) {
    var originalWidth = width, originalHeight = height;

    // Constrain width/height to max.
    var marginWidth = 2 * margin;
    if (width + marginWidth > maxWidth)
      width = maxWidth - marginWidth;

    var marginHeight = 2 * margin;
    if (height + marginHeight > maxHeight)
      height = maxHeight - marginHeight;

    var changeWidth = (originalWidth - width) / originalWidth;
    var changeHeight = (originalHeight - height) / originalHeight;

    // Adjust width/height if oversized.
    if (changeWidth > 0 || changeHeight > 0) {
      // Preserve original aspect ratio according to greatest change.
      if (changeWidth > changeHeight) {
        height = Math.round((originalHeight / originalWidth) * width);
      } else if (changeHeight > changeWidth) {
        width = Math.round((originalWidth / originalHeight) * height);
      }
    }

    return { width: width, height: height };
  }

  /**
   * Sets the size of the container element to the size of the window.
   */
  function setContainerSize() {
    setStyle(containerElement, "width", documentElement.clientWidth + "px");
    setStyle(containerElement, "height", documentElement.clientHeight + "px");

    if (currentPlayer)
      setWrapperSize(getWrapperSize());
  }

  /**
   * Sets the position of the container element to the top left corner of
   * the window. Necessary when using absolute positioning instead of fixed.
   */
  function setContainerPosition() {
    setStyle(containerElement, "left", documentElement.scrollLeft + "px");
    setStyle(containerElement, "top", "0px");
  }

  var troubleElementTagNames = [ "select", "object", "embed", "canvas" ];
  var troubleVisibilityCache = [];

  /**
   * Toggles the visibility of elements that are troublesome for overlays.
   */
  function toggleTroubleElements(on) {
    if (on) {
      forEach(troubleVisibilityCache, function (item) {
        setStyle(item.element, "visibility", item.visibility || "");
      });
    } else {
      troubleVisibilityCache = [];

      forEach(troubleElementTagNames, function (tagName) {
        forEach(document.getElementsByTagName(tagName), function (element) {
          troubleVisibilityCache.push({
            element: element,
            visibility: getStyle(element, "visibility")
          });

          setStyle(element, "visibility", "hidden");
        });
      });
    }
  }

  /**
   * Creates a new player object based on the properties of the given object.
   * Valid properties include:
   *
   *   - url          The URL of the content to display
   *   - width        (optional) The width of the content
   *   - height       (optional) The height of the content
   *   - playerClass  (optional) The player class to use to play the content.
   *                  Can be guessed in most cases from the URL
   *   - encodings    (video only) Encoding name/URL pairs of alternate URL's
   *                  for the video. Possible encoding names are "h264", "ogg"
   *                  "webm", and "flv"
   *   - posterUrl    (video only) The URL to a poster image of the video
   *   - flashParams  (flash only) Name/value pairs of <param>'s to use for
   *                  the Flash <object>
   *   - flashVars    (flash only) Name/value pairs of variables to pass to
   *                  the Flash object as variables
   *
   * If a string is given, it will be used as the value of the URL. If a DOM
   * element is given, it should have an href property (i.e. either an <a> or
   * an <area> element) which will be used as the URL. It may also contain
   * a data-shadowbox attribute that has any of the other options formatted
   * in a JSON string.
   *
   * If no player is specified, it will be guessed using the registered player
   * for the URL's file extension (see shadowbox.registerPlayer).
   *
   * Returns null if no player is able to be created, or this browser does
   * not have proper support for that content.
   */
  shadowbox.makePlayer = function (object) {
    if (typeof object === "string") {
      object = { url: object };
    } else if (isElement(object) && object.href) {
      // The object is a DOM element. Should be an <a> or <area>. The
      // data-shadowbox attribute may contain a string specifying
      // options for the player object (see parseData).
      var data = object.getAttribute("data-shadowbox");

      object = { url: object.href };

      if (data)
        mergeProperties(object, parseData(data));
    }

    if (object && typeof object.url === "string") {
      var playerClass;
      if (object.playerClass) {
        playerClass = object.playerClass;
      } else {
        // Guess the player class using the URL's file extension.
        var match = object.url.match(/\.([0-9a-z]+)(\?.*)?$/i);

        if (match) {
          var extension = match[1].toLowerCase();
          playerClass = shadowbox.players[extension];
        }
      }

      playerClass = playerClass || FramePlayer;

      var player = new playerClass(object, "sb-player-" + String(shadowbox.guid++));

      if (player.isSupported())
        return player;
    }

    return null;
  };

  // Toggles the visibility of clickable controls.
  function toggleControls(on) {
    var name = "";

    if (on) {
      name += "active";

      if (getNextIndex() !== -1)
        name += " has-next";

      if (getPreviousIndex() !== -1)
        name += " has-prev";
    }

    containerElement.className = name;
  }

  var resizeTimer, scrollTimer, mouseMoveTimer;

  // Toggles window resize and scroll event handlers.
  function toggleWindowHandlers(on) {
    var addOrRemoveEvent;
    if (on) {
      addOrRemoveEvent = addEvent;
    } else {
      addOrRemoveEvent = removeEvent;

      // Clear cached timers.
      if (resizeTimer) {
        clearTimeout(resizeTimer);
        resizeTimer = null;
      }

      if (scrollTimer) {
        clearTimeout(scrollTimer);
        scrollTimer = null;
      }
    }

    addOrRemoveEvent(window, "resize", handleWindowResize);

    if (!supportsFixed)
      addOrRemoveEvent(window, "scroll", handleWindowScroll);
  }

  // Updates the size of the container when the window size changes.
  function handleWindowResize() {
    if (resizeTimer) {
      clearTimeout(resizeTimer);
      resizeTimer = null;
    }

    resizeTimer = setTimeout(function () {
      resizeTimer = null;
      setContainerSize();
    }, 10);
  }

  // Updates the position of the container when the window scrolls.
  function handleWindowScroll() {
    if (scrollTimer) {
      clearTimeout(scrollTimer);
      scrollTimer = null;
    }

    scrollTimer = setTimeout(function () {
      scrollTimer = null;
      setContainerPosition();
    }, 10);
  }

  // Toggles document mouse move handler on/off.
  function toggleMouseMoveHandler(on) {
    if (supportsTouch) {
      toggleControls(on);
      return;
    }

    var addOrRemoveEvent;
    if (on) {
      addOrRemoveEvent = addEvent;
    } else {
      addOrRemoveEvent = removeEvent;

      // Clear cached timers.
      if (mouseMoveTimer) {
        clearTimeout(mouseMoveTimer);
        mouseMoveTimer = null;
      }
    }

    addOrRemoveEvent(document, "mousemove", handleMouseMove);
  }

  var lastMouseX, lastMouseY;

  // Shows clickable controls when the mouse moves.
  function handleMouseMove(event) {
    // Ignore consecutive mousemove events from the same location.
    if (lastMouseX !== event.clientX || lastMouseY !== event.clientY) {
      lastMouseX = event.clientX;
      lastMouseY = event.clientY;

      if (mouseMoveTimer) {
        clearTimeout(mouseMoveTimer);
        mouseMoveTimer = null;
      } else {
        toggleControls(1);
      }

      mouseMoveTimer = setTimeout(function () {
        mouseMoveTimer = null;
        toggleControls(0);
      }, 1500);
    }
  }

  function toggleKeyDownHandler(on) {
    if (currentOptions.enableKeys)
      (on ? addEvent : removeEvent)(document, "keydown", handleDocumentKeyDown);
  }

  var KEY_ESCAPE  = 27;
  var KEY_SPACE   = 32;
  var KEY_LEFT    = 37;
  var KEY_RIGHT   = 39;
  var KEY_Q       = 81;
  var KEY_X       = 88;

  function handleDocumentKeyDown(event) {
    if (eventHasModifierKey(event))
      return;

    switch (event.keyCode) {
    case KEY_ESCAPE:
    case KEY_Q:
    case KEY_X:
      event.preventDefault();
      shadowbox.close();
      break;
    case KEY_LEFT:
      event.preventDefault();
      shadowbox.showPrevious();
      break;
    case KEY_RIGHT:
      event.preventDefault();
      shadowbox.showNext();
      break;
    case KEY_SPACE:
      if (currentPlayer && isFunction(currentPlayer.togglePlay)) {
        event.preventDefault();
        currentPlayer.togglePlay();
      }
      break;
    }
  }

  function eventHasModifierKey(event) {
    return event.ctrlKey || event.metaKey;
  }

  function toggleClickHandler(on) {
    (on ? addEvent : removeEvent)(document, 'click', handleDocumentClick);
  }

  function handleDocumentClick(event) {
    var target = event.target;

    if (isElement(target)) {
      var matcher = /^(?:shadow|light)box(?:\[(\w+)\])?$/i,
          links = [],
          index = 0,
          match;

      // Find an ancestor node with rel="shadowbox" attribute.
      while (target) {
        match = (target.rel || "").match(matcher);

        if (match) {
          var galleryName = match[1];

          // Look for other <a> elements in the document that also have
          // rel="shadowbox" attribute with the same gallery.
          if (galleryName) {
            var galleryMatcher = new RegExp("^(shadow|light)box\\[" + galleryName + "\\]$", "i");

            forEach(document.getElementsByTagName('a'), function (link) {
              if (link.rel && galleryMatcher.test(link.rel)) {
                if (link == target)
                  index = links.length;

                links.push(link);
              }
            });
          } else {
            links.push(target);
          }

          break;
        }

        target = target.parentNode;
      }

      // Good for debugging.
      // event.preventDefault();

      if (links.length > 0 && shadowbox.open(links, index) > 0)
        event.preventDefault(); // Prevent the browser from following the link.
    }
  }


  //// PLAYERS ////


  shadowbox.FramePlayer = FramePlayer;

  /**
   * A player that displays its content inside an <iframe>. This is the default
   * player for Shadowbox that is used when no other player is suitable for a
   * piece of content.
   */
  function FramePlayer(object, id) {
    this.url = object.url;
    this.width = object.width ? parseInt(object.width, 10) : documentElement.clientWidth;
    this.height = object.height ? parseInt(object.height, 10) : documentElement.clientHeight;
    this.id = id;

    // Preload the iframe so it's ready when needed.
    this.isReady = false;
    this._preload();
  }

  mergeProperties(FramePlayer.prototype, {

    _preload: function () {
      var iframe = makeDom("iframe");

      iframe.id = this.id;
      iframe.name = this.id;
      iframe.width = "0px";
      iframe.height = "0px";
      iframe.frameBorder = "0";
      iframe.marginWidth = "0";
      iframe.marginHeight = "0";
      iframe.scrolling = "auto";
      iframe.allowTransparency = "true";
      iframe.src = this.url;

      var self = this;

      if (iframe.attachEvent) {
        iframe.attachEvent("onload", function () {
          self.isReady = true;
        });
      } else {
        iframe.onload = function () {
          self.isReady = true;
        };
      }

      // Starts the actual loading of the iframe.
      makeDom(document.body, iframe);

      this.element = iframe;
    },

    /**
     * Returns true if this player is supported on this browser.
     */
    isSupported: function () {
      return true;
    },

    /**
     * Inserts this object as the only child of the given DOM element.
     */
    injectInto: function (element) {
      removeChildren(element);

      var iframe = this.element;
      iframe.style.visibility = "hidden";
      iframe.width = "100%";
      iframe.height = "100%";

      element.appendChild(iframe);

      iframe.style.visibility = "";
    },

    /**
     * Removes this object from the DOM.
     */
    remove: function () {
      if (this.element) {
        removeElement(this.element);
        delete this.element;

        // Needed for Firefox, IE <= 8 throws error.
        try {
          delete window.frames[this.id];
        } catch (error) {}
      }
    }

  });

  shadowbox.PhotoPlayer = PhotoPlayer;

  /**
   * A player that is used for displaying images.
   */
  function PhotoPlayer(object, id) {
    this.url = object.url;
    this.width = parseInt(object.width, 10);
    this.height = parseInt(object.height, 10);
    this.id = id;

    // Preload the image so it's ready when needed.
    this.isReady = false;
    this._preload();
  }

  mergeProperties(PhotoPlayer.prototype, {

    fadeCover: true,

    _preload: function () {
      var preloader = new Image;

      var self = this;
      preloader.onload = function () {
        // Width and height default to image dimensions.
        self.width = self.width || preloader.width;
        self.height = self.height || preloader.height;

        // Ready to go.
        self.isReady = true;

        // Clean up to prevent memory leak in IE.
        preloader.onload = preloader = null;
      };

      // Start loading the image.
      preloader.src = this.url;
    },

    /**
     * Returns true if this player is supported on this browser.
     */
    isSupported: function () {
      return true;
    },

    /**
     * Inserts this object as the only child of the given DOM element.
     */
    injectInto: function (element) {
      element.innerHTML = '<img id="' + this.id + '" src="' + this.url + '" width="100%" height="100%">';
      this.element = element.firstChild;
    },

    /**
     * Removes this object from the DOM.
     */
    remove: function () {
      if (this.element) {
        removeElement(this.element);
        delete this.element;
      }
    }

  });


  //// JAVASCRIPT UTILITIES ////


  function isFunction(object) {
    return typeof object === "function";
  }

  function isArray(object) {
    if (isFunction(Array.isArray))
      return Array.isArray(object);

    return Object.prototype.toString.call(object) === "[object Array]";
  }

  /**
   * Calls the given callback function for each element in the given object,
   * which must be an array-like object. Return false from any callback to
   * stop execution.
   */
  function forEach(object, callback) {
    var length = object.length, index = 0, item;
    for (item = object[0]; index < length && callback.call(object, item, index, object) !== false; item = object[++index]) {}
  }

  /**
   * Merges all properties of extension into the given object.
   */
  function mergeProperties(object, extension) {
    for (var property in extension) {
      if (extension.hasOwnProperty(property))
        object[property] = extension[property];
    }

    return object;
  }

  /**
   * Gets the current time in milliseconds.
   */
  function getTime() {
    return (new Date).getTime();
  }

  /**
   * Parses the value of the data-shadowbox attribute which is a string
   * of key=value pairs separated by commas, e.g.:
   *
   * "margin=40,url=http://example.com" => { margin: 40, url: "http://example.com" }
   */
  function parseData(data) {
    var hash = {};
    var pairs = data.split(/\s*,\s*/);

    forEach(pairs, function (pair) {
      var split = pair.split(/\s*=\s*/);
      if (split.length === 2) {
        hash[split[0]] = parseValue(split[1]);
      } else {
        throw new Error('Invalid data: ' + pair);
      }
    });

    return hash;
  }

  var numericRe = /^(\d+)?\.?\d+$/;

  function parseValue(value) {
    return numericRe.test(value) ? parseFloat(value, 10) : value;
  }

  /**
   * Animates from one numeric value to another over the given duration,
   * calling the given callback for each frame with the eased value. Return
   * false from the callback at any time to cancel the animation.
   */
  function animate(from, to, duration, frameHandler, callback) {
    var delta = to - from;

    if (delta === 0 || duration === 0 || !currentOptions.animate) {
      frameHandler(to);

      if (isFunction(callback))
        callback();

      return; // Don't animate!
    }

    // Convert duration to milliseconds.
    duration = (duration || 0.35) * 1000;

    var ease = currentOptions.ease,
        begin = getTime(),
        end = begin + duration,
        time;

    var timer = setInterval(function () {
      time = getTime();

      if (time >= end) {
        clearInterval(timer);
        timer = null;

        frameHandler(to);

        if (isFunction(callback))
          callback();
      } else if (frameHandler(from + ease((time - begin) / duration) * delta) === false) {
        clearInterval(timer);
        timer = null;
      }
    }, 10);
  }

  function waitUntil(check, callback) {
    var timer = setInterval(function () {
      if (check()) {
        clearInterval(timer);
        timer = null;
        callback();
      }
    }, 10);
  }


  //// DOM UTILITIES ////


  function isElement(object) {
    return object && object.nodeType === 1; // Node.ELEMENT_NODE
  }

  /**
   * Multipurpose utility function for creating DOM elements, assigning
   * attributes, and appending child nodes.
   */
  function makeDom(element, properties, children) {
    if (typeof element === "string")
      element = document.createElement(element);

    if (isArray(properties)) {
      children = properties;
      properties = null;
    } else if (properties && properties.nodeType) {
      children = [ properties ];
      properties = null;
    }

    if (properties)
      mergeProperties(element, properties);

    if (isArray(children)) {
      forEach(children, function (child) {
        element.appendChild(child);
      });
    }

    return element;
  }

  /**
   * Removes the given element from the DOM.
   */
  function removeElement(element) {
    return element.parentNode.removeChild(element);
  }

  /**
   * Removes all child nodes from the given element.
   */
  function removeChildren(element) {
    var child = element.firstChild;

    while (child) {
      element.removeChild(child);
      child = element.firstChild;
    }
  }

  /**
   * Wraps an event handler to cancel the default behavior and prevent
   * event bubbling when called.
   */
  function cancel(callback) {
    return function (event) {
      event.preventDefault();
      event.stopPropagation();
      callback(event);
    };
  }

  /**
   * Animates the style of an element from its current value to another over
   * the given duration. Calls the given callback when complete.
   */
  function animateStyle(element, style, to, duration, callback) {
    var from = parseFloat(getStyle(element, style)) || 0;

    var frameHandler;
    if (style === "opacity") {
      frameHandler = function (value) {
        setStyle(element, style, value);
      };
    } else {
      // Assume pixel values for all styles besides opacity.
      frameHandler = function (value) {
        setStyle(element, style, Math.round(value) + "px");
      };
    }

    animate(from, to, duration, frameHandler, callback);
  }

  var opacityRe = /opacity=([^)]*)/i;
  var getComputedStyle = document.defaultView && document.defaultView.getComputedStyle;

  /**
   * Gets the current value of the given style on the given element. The style
   * name should be camel-cased.
   *
   * Note: This function is not safe for retrieving float or non-pixel values
   * in Internet Explorer.
   */
  function getStyle(element, style) {
    var value = "";

    if (!supportsOpacity && style == "opacity" && element.currentStyle) {
      if (opacityRe.test(element.currentStyle.filter || ""))
        value = (parseFloat(RegExp.$1) / 100) + "";

      return value == "" ? "1" : value;
    }

    if (getComputedStyle) {
      var computedStyle = getComputedStyle(element, null);

      if (computedStyle)
        value = computedStyle[style];

      if (style == "opacity" && value == "")
        value = "1";
    } else {
      value = element.currentStyle[style];
    }

    return value;
  }

  /**
   * Sets the style on the given element to the given value. The style name
   * should be camel-cased.
   *
   * Note: This function is not safe for setting float values.
   */
  function setStyle(element, styleName, value) {
    var elementStyle = element.style;

    if (styleName == "opacity") {
      if (value == 1) {
        value = "";
      } else {
        value = (value < 0.00001 ? 0 : value);
      }

      if (!supportsOpacity) {
        elementStyle.zoom = 1; // Trigger hasLayout.

        if (value == 1) {
          if (typeof elementStyle.filter === "string" && (/alpha/i).test(elementStyle.filter))
            elementStyle.filter = elementStyle.filter.replace(/\s*[\w\.]*alpha\([^\)]*\);?/gi, "");
        } else {
          elementStyle.filter = (elementStyle.filter || "").replace(/\s*[\w\.]*alpha\([^\)]*\)/gi, "") +
            " alpha(opacity=" + (value * 100) + ")";
        }

        return;
      }
    }

    elementStyle[styleName] = value;
  }

  // Event handling functions modified from originals by Dean Edwards.
  // http://dean.edwards.name/my/events.js

  /**
   * Adds an event handler to the given element. The handler will be called
   * in the context of the element with the event object as its only argument.
   */
  function addEvent(element, type, handler) {
    if (element.addEventListener) {
      element.addEventListener(type, handler, false);
    } else {
      if (element.nodeType === 3 || element.nodeType === 8) // Node.TEXT_NODE or Node.COMMENT_NODE
        return;

      if (element.setInterval && (element !== window && !element.frameElement))
        element = window;

      if (!handler.__guid)
        handler.__guid = shadowbox.guid++;

      if (!element.events)
        element.events = {};

      var handlers = element.events[type];

      if (!handlers) {
        handlers = element.events[type] = {};

        if (element["on" + type])
          handlers[0] = element["on" + type];
      }

      handlers[handler.__guid] = handler;

      element["on" + type] = handleEvent;
    }
  }

  function handleEvent(event) {
    event = event || fixEvent(((this.ownerDocument || this.document || this).parentWindow || window).event);

    var handlers = this.events[event.type], result = true;

    for (var id in handlers) {
      if (handlers[id].call(this, event) === false)
        result = false;
    }

    return result;
  }

  function fixEvent(event) {
    event.preventDefault = preventDefault;
    event.stopPropagation = stopPropagation;
    event.target = event.srcElement;
    event.keyCode = event.which;
    return event;
  }

  function preventDefault() {
    this.returnValue = false;
  }

  function stopPropagation() {
    this.cancelBubble = true;
  }

  /**
   * Removes an event handler from the given element.
   */
  function removeEvent(element, type, handler) {
    if (element.removeEventListener) {
      element.removeEventListener(type, handler, false);
    } else {
      if (element.events && element.events[type] && handler.__guid)
        delete element.events[type][handler.__guid];
    }
  }

  // Setup document click handler.
  toggleClickHandler(1);

  // Register the photo player for common image extensions.
  shadowbox.registerPlayer(shadowbox.PhotoPlayer, [ "gif", "jpg", "jpeg", "png", "bmp" ]);

  // Expose for the sake of shadowbox-flash.js and shadowbox-video.js.
  shadowbox.forEach = forEach;
  shadowbox.mergeProperties = mergeProperties;
  shadowbox.makeDom = makeDom;
  shadowbox.removeElement = removeElement;
  shadowbox.removeChildren = removeChildren;
  shadowbox.addEvent = addEvent;
  shadowbox.removeEvent = removeEvent;

  // Expose.
  global.shadowbox = shadowbox;

}(this));

/* Copyright 2012 Mozilla Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import {
  addLinkAttributes, PDFJS, removeNullCharacters
} from 'pdfjs-lib';

const DEFAULT_TITLE = '\u2013';

/**
 * @typedef {Object} PDFOutlineViewerOptions
 * @property {HTMLDivElement} container - The viewer element.
 * @property {IPDFLinkService} linkService - The navigation/linking service.
 * @property {EventBus} eventBus - The application event bus.
 */

/**
 * @typedef {Object} PDFOutlineViewerRenderParameters
 * @property {Array|null} outline - An array of outline objects.
 */

class PDFOutlineViewer {
  /**
   * @param {PDFOutlineViewerOptions} options
   */
  constructor({ container, linkService, eventBus, }) {
    this.container = container;
    this.linkService = linkService;
    this.eventBus = eventBus;

    this.reset();
  }

  reset() {
    this.outline = null;
    this.lastToggleIsShow = true;

    // Remove the outline from the DOM.
    this.container.textContent = '';

    // Ensure that the left (right in RTL locales) margin is always reset,
    // to prevent incorrect outline alignment if a new document is opened.
    this.container.classList.remove('outlineWithDeepNesting');
  }

  /**
   * @private
   */
  _dispatchEvent(outlineCount) {
    this.eventBus.dispatch('outlineloaded', {
      source: this,
      outlineCount,
    });
  }

  /**
   * @private
   */
  _bindLink(element, item) {
    if (item.url) {
      addLinkAttributes(element, {
        url: item.url,
        target: (item.newWindow ? PDFJS.LinkTarget.BLANK : undefined),
      });
      return;
    }
    let destination = item.dest;

    element.href = this.linkService.getDestinationHash(destination);
    element.onclick = () => {
      if (destination) {
        this.linkService.navigateTo(destination);
      }
      return false;
    };
  }

  /**
   * @private
   */
  _setStyles(element, item) {
    let styleStr = '';
    if (item.bold) {
      styleStr += 'font-weight: bold;';
    }
    if (item.italic) {
      styleStr += 'font-style: italic;';
    }

    if (styleStr) {
      element.setAttribute('style', styleStr);
    }
  }

  /**
   * Prepend a button before an outline item which allows the user to toggle
   * the visibility of all outline items at that level.
   *
   * @private
   */
  _addToggleButton(div) {
    let toggler = document.createElement('div');
    toggler.className = 'outlineItemToggler';
    toggler.onclick = (evt) => {
      evt.stopPropagation();
      toggler.classList.toggle('outlineItemsHidden');

      if (evt.shiftKey) {
        let shouldShowAll = !toggler.classList.contains('outlineItemsHidden');
        this._toggleOutlineItem(div, shouldShowAll);
      }
    };
    div.insertBefore(toggler, div.firstChild);
  }

  /**
   * Toggle the visibility of the subtree of an outline item.
   *
   * @param {Element} root - the root of the outline (sub)tree.
   * @param {boolean} show - whether to show the outline (sub)tree. If false,
   *   the outline subtree rooted at |root| will be collapsed.
   *
   * @private
   */
  _toggleOutlineItem(root, show) {
    this.lastToggleIsShow = show;
    let togglers = root.querySelectorAll('.outlineItemToggler');
    for (let i = 0, ii = togglers.length; i < ii; ++i) {
      togglers[i].classList[show ? 'remove' : 'add']('outlineItemsHidden');
    }
  }

  /**
   * Collapse or expand all subtrees of the outline.
   */
  toggleOutlineTree() {
    if (!this.outline) {
      return;
    }
    this._toggleOutlineItem(this.container, !this.lastToggleIsShow);
  }

  /**
   * @param {PDFOutlineViewerRenderParameters} params
   */
  render({ outline, }) {
    let outlineCount = 0;

    if (this.outline) {
      this.reset();
    }
    this.outline = outline || null;

    if (!outline) {
      this._dispatchEvent(outlineCount);
      return;
    }

    let fragment = document.createDocumentFragment();
    let queue = [{ parent: fragment, items: this.outline, }];
    let hasAnyNesting = false;
    while (queue.length > 0) {
      let levelData = queue.shift();
      for (let i = 0, len = levelData.items.length; i < len; i++) {
        let item = levelData.items[i];

        let div = document.createElement('div');
        div.className = 'outlineItem';

        let element = document.createElement('a');
        this._bindLink(element, item);
        this._setStyles(element, item);
        element.textContent =
          removeNullCharacters(item.title) || DEFAULT_TITLE;

        div.appendChild(element);

        if (item.items.length > 0) {
          hasAnyNesting = true;
          this._addToggleButton(div);

          let itemsDiv = document.createElement('div');
          itemsDiv.className = 'outlineItems';
          div.appendChild(itemsDiv);
          queue.push({ parent: itemsDiv, items: item.items, });
        }

        levelData.parent.appendChild(div);
        outlineCount++;
      }
    }
    if (hasAnyNesting) {
      this.container.classList.add('outlineWithDeepNesting');
    }

    this.container.appendChild(fragment);

    this._dispatchEvent(outlineCount);
  }
}

export {
  PDFOutlineViewer,
};

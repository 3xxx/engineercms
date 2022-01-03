# Table options []({{ site.repo }}/blob/develop/docs/_i18n/{{ site.lang }}/documentation/table-options.md)

---

The table options are defined in `jQuery.fn.bootstrapTable.defaults`.

<table class="table"
       id="t"
       data-search="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-mobile-responsive="true">
    <thead>
    <tr>
        <th>Name</th>
        <th>Attribute</th>
        <th>Type</th>
        <th>Default</th>
        <th>Description</th>
    </tr>
    </thead>
    <tbody>
    <tr>
        <td>-</td>
        <td>data-toggle</td>
        <td>String</td>
        <td>'table'</td>
        <td>Activate bootstrap table without writing JavaScript.</td>
    </tr>
    <tr>
        <td>classes</td>
        <td>data-classes</td>
        <td>String</td>
        <td>'table table-hover'</td>
        <td>The class name of table. By default, the table is bordered, you can add 'table-no-bordered' to remove table-bordered style.</td>
    </tr>
    <tr>
        <td>sortClass</td>
        <td>data-sort-class</td>
        <td>String</td>
        <td>undefined</td>
        <td>The class name of the td elements which are sorted.</td>
    </tr>
    <tr>
        <td>height</td>
        <td>data-height</td>
        <td>Number</td>
        <td>undefined</td>
        <td>The height of table.</td>
    </tr>
    <tr>
        <td>undefinedText</td>
        <td>data-undefined-text</td>
        <td>String</td>
        <td>'-'</td>
        <td>Defines the default undefined text.</td>
    </tr>
    <tr>
        <td>striped</td>
        <td>data-striped</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to stripe the rows.</td>
    </tr>
    <tr>
        <td>sortName</td>
        <td>data-sort-name</td>
        <td>String</td>
        <td>undefined</td>
        <td>Defines which column will be sorted.</td>
    </tr>
    <tr>
        <td>sortOrder</td>
        <td>data-sort-order</td>
        <td>String</td>
        <td>'asc'</td>
        <td>Defines the column sort order, can only be 'asc' or 'desc'.</td>
    </tr>
    <tr>
        <td>sortStable</td>
        <td>data-sort-stable</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to get a stable sorting. We will add <code>_position</code> property to the row.</td>
    </tr>
    <tr>
        <td>rememberOrder</td>
        <td>data-remember-order</td>
        <td>Boolean</td>
        <td>false</td>
        <td>Set <code>true</code> remember the order for each column.</td>
    </tr>
    <tr>
        <td>iconsPrefix</td>
        <td>data-icons-prefix</td>
        <td>String</td>
        <td>'glyphicon'</td>
        <td>Defines icon set name ('glyphicon' or 'fa' for FontAwesome). By default 'glyphicon' is used. </td>
    </tr>
    <tr>
        <td>iconSize</td>
        <td>data-icon-size</td>
        <td>String</td>
        <td>undefined</td>
        <td>Defines icon size: <ul><li>undefined => btn</li><li>xs => btn-xs</li><li>sm => btn-sm</li><li>lg => btn-lg</li></ul>
        </td>
    </tr>
    <tr>
        <td>buttonsClass</td>
        <td>data-buttons-class</td>
        <td>String</td>
        <td>'default'</td>
        <td>Defines the Bootstrap class (added after 'btn-') of table buttons: EX: 'primary', 'danger', 'warning'...</td>
    </tr>
    <tr>
        <td>icons</td>
        <td>data-icons</td>
        <td>Object</td>
        <td>{<br/>
        &nbsp;&nbsp;<i>paginationSwitchDown:</i> 'glyphicon-collapse-down icon-chevron-down',<br/>
        &nbsp;&nbsp;<i>paginationSwitchUp:</i> 'glyphicon-collapse-up icon-chevron-up',<br/>
        &nbsp;&nbsp;<i>refresh:</i> 'glyphicon-refresh icon-refresh',<br/>
        &nbsp;&nbsp;<i>toggle:</i> 'glyphicon-list-alt icon-list-alt',<br/>
        &nbsp;&nbsp;<i>columns:</i> 'glyphicon-th icon-th',<br/>
        &nbsp;&nbsp;<i>detailOpen:</i> 'glyphicon-plus icon-plus',<br/>
        &nbsp;&nbsp;<i>detailClose:</i> 'glyphicon-minus icon-minus'<br/>
        }
        </td>
        <td>Defines icons used in the toolbar, pagination, and details view.</td>
    </tr>
    <tr>
        <td>columns</td>
        <td>-</td>
        <td>Array</td>
        <td>[]</td>
        <td>The table columns config object, see column properties for more details.
        </td>
    </tr>
    <tr>
        <td>data</td>
        <td>-</td>
        <td>Array</td>
        <td>[]</td>
        <td>The data to be loaded.</td>
    </tr>
    <tr>
        <td>dataField</td>
        <td>data-data-field</td>
        <td>String</td>
        <td>'rows'</td>
        <td>Key in incoming json containing rows data list.</td>
    </tr>
    <tr>
        <td>totalField</td>
        <td>data-total-field</td>
        <td>String</td>
        <td>'total'</td>
        <td>Key in incoming json containing  "total" data .</td>
    </tr>
    <tr>
        <td>ajax</td>
        <td>data-ajax</td>
        <td>Function</td>
        <td>undefined</td>
        <td>A method to replace ajax call. Should implement the same API as jQuery ajax method.</td>
    </tr>
    <tr>
        <td>method</td>
        <td>data-method</td>
        <td>String</td>
        <td>'get'</td>
        <td>The method type to request remote data.</td>
    </tr>
    <tr>
        <td>url</td>
        <td>data-url</td>
        <td>String</td>
        <td>undefined</td>
        <td>
        	A URL to request data from remote site.
        	<br/>Note that the required server response format is different depending on whether the 'sidePagination'
        	option is specified. See the following examples:
        	<ul>
        		<li><a href="https://github.com/wenzhixin/bootstrap-table-examples/blob/master/json/data1.json">Without server-side pagination</a></li>
        		<li><a href="https://github.com/wenzhixin/bootstrap-table-examples/blob/master/json/data2.json">With server-side pagination</a></li>
        	</ul>
        </td>
    </tr>
    <tr>
        <td>cache</td>
        <td>data-cache</td>
        <td>Boolean</td>
        <td>true</td>
        <td>False to disable caching of AJAX requests.</td>
    </tr>
    <tr>
        <td>contentType</td>
        <td>data-content-type</td>
        <td>String</td>
        <td>'application/json'</td>
        <td>The contentType of request remote data.</td>
    </tr>
    <tr>
        <td>dataType</td>
        <td>data-data-type</td>
        <td>String</td>
        <td>'json'</td>
        <td>The type of data that you are expecting back from the server.</td>
    </tr>
    <tr>
        <td>ajaxOptions</td>
        <td>data-ajax-options</td>
        <td>Object</td>
        <td>{}</td>
        <td>Additional options for submit ajax request. List of values: <a href="http://api.jquery.com/jQuery.ajax">http://api.jquery.com/jQuery.ajax</a>.</td>
    </tr>
    <tr>
        <td>queryParams</td>
        <td>data-query-params</td>
        <td>Function</td>
        <td>function(params) {<br>return params;<br>}</td>
        <td>
        When requesting remote data, you can send additional parameters by modifying queryParams.
        If queryParamsType = 'limit', the params object contains: <br>
        limit, offset, search, sort, order
        Else, it contains: <br>
        pageSize, pageNumber, searchText, sortName, sortOrder. <br>
        Return false to stop request.
        </td>
    </tr>
    <tr>
        <td>queryParamsType</td>
        <td>data-query-params-type</td>
        <td>String</td>
        <td>'limit'</td>
        <td>Set 'limit' to send query params width RESTFul type.</td>
    </tr>
    <tr>
        <td>responseHandler</td>
        <td>data-response-handler</td>
        <td>Function</td>
        <td>function(res) {<br>return res;<br>}</td>
        <td>
        Before load remote data, handler the response data format, the parameters object contains: <br>
        res: the response data.
        </td>
    </tr>
    <tr>
        <td>pagination</td>
        <td>data-pagination</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to show a pagination toolbar on table bottom.</td>
    </tr>
    <tr>
        <td>paginationLoop</td>
        <td>data-pagination-loop</td>
        <td>Boolean</td>
        <td>true</td>
        <td>True to enable pagination continuous loop mode.</td>
    </tr>
    <tr>
        <td>onlyInfoPagination</td>
        <td>data-only-info-pagination</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to show only the quantity of the data that is showing in the table. It needs the pagination table options is set to true.</td>
    </tr>
    <tr>
        <td>sidePagination</td>
        <td>data-side-pagination</td>
        <td>String</td>
        <td>'client'</td>
        <td>
        	Defines the side pagination of table, can only be 'client' or 'server'.
        	Using 'server' side requires either setting the 'url' or 'ajax' option.
        	<br/>Note that the required server response format is different depending on whether
        	the 'client' or 'server' option is specified. See the following examples:
        	<ul>
        		<li><a href="https://github.com/wenzhixin/bootstrap-table-examples/blob/master/json/data1.json">Without server-side pagination</a></li>
        		<li><a href="https://github.com/wenzhixin/bootstrap-table-examples/blob/master/json/data2.json">With server-side pagination</a></li>
        	</ul>
        </td>
    </tr>
    <tr>
        <td>pageNumber</td>
        <td>data-page-number</td>
        <td>Number</td>
        <td>1</td>
        <td>When set pagination property, initialize the page number.</td>
    </tr>
    <tr>
        <td>pageSize</td>
        <td>data-page-size</td>
        <td>Number</td>
        <td>10</td>
        <td>When set pagination property, initialize the page size.</td>
    </tr>
    <tr>
        <td>pageList</td>
        <td>data-page-list</td>
        <td>Array</td>
        <td>[10, 25, 50, 100]</td>
        <td>When set pagination property, initialize the page size selecting list. If you include the 'All' or 'Unlimited' option, all the records will be shown in your table.</td>
    </tr>
    <tr>
        <td>selectItemName</td>
        <td>data-select-item-name</td>
        <td>String</td>
        <td>'btSelectItem'</td>
        <td>The name of radio or checkbox input.</td>
    </tr>
    <tr>
        <td>smartDisplay</td>
        <td>data-smart-display</td>
        <td>Boolean</td>
        <td>true</td>
        <td>True to display pagination or card view smartly.</td>
    </tr>
    <tr>
        <td>escape</td>
        <td>data-escape</td>
        <td>Boolean</td>
        <td>false</td>
        <td>Escapes a string for insertion into HTML,
        replacing <code>&</code>, <code><</code>, <code>></code>,
        <code>"</code>, <code>`</code>, and <code>'</code> characters.</td>
    </tr>
    <tr>
        <td>search</td>
        <td>data-search</td>
        <td>Boolean</td>
        <td>false</td>
        <td>Enable the search input.</td>
    </tr>
    <tr>
        <td>searchOnEnterKey</td>
        <td>data-search-on-enter-key</td>
        <td>Boolean</td>
        <td>false</td>
        <td>The search method will be executed until the Enter key is pressed.</td>
    </tr>
    <tr>
        <td>strictSearch</td>
        <td>data-strict-search</td>
        <td>Boolean</td>
        <td>false</td>
        <td>Enable the strict search.</td>
    </tr>
	<tr>
        <td>searchText</td>
        <td>data-search-text</td>
        <td>String</td>
        <td>''</td>
        <td>When set search property, initialize the search text.</td>
    </tr>
    <tr>
        <td>searchTimeOut</td>
        <td>data-search-time-out</td>
        <td>Number</td>
        <td>500</td>
        <td>Set timeout for search fire.</td>
    </tr>
    <tr>
        <td>trimOnSearch</td>
        <td>data-trim-on-search</td>
        <td>Boolean</td>
        <td>true</td>
        <td>True to trim spaces in search field.</td>
    </tr
    <tr>
        <td>showHeader</td>
        <td>data-show-header</td>
        <td>Boolean</td>
        <td>true</td>
        <td>False to hide the table header.</td>
    </tr>
    <tr>
        <td>showFooter</td>
        <td>data-show-footer</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to show the summary footer row.</td>
    </tr>
    <tr>
        <td>showColumns</td>
        <td>data-show-columns</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to show the columns drop down list.</td>
    </tr>
    <tr>
        <td>showRefresh</td>
        <td>data-show-refresh</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to show the refresh button.</td>
    </tr>
    <tr>
        <td>showToggle</td>
        <td>data-show-toggle</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to show the toggle button to toggle table / card view.
        </td>
    </tr>
    <tr>
        <td>showPaginationSwitch</td>
        <td>data-show-pagination-switch</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to show the pagination switch button.</td>
    </tr>
    <tr>
        <td>showFullscreen</td>
        <td>data-show-fullscreen</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to show the fullscreen button.</td>
    </tr>
    <tr>
        <td>minimumCountColumns</td>
        <td>data-minimum-count-columns</td>
        <td>Number</td>
        <td>1</td>
        <td>The minimum number of columns to hide from the columns drop down list.
        </td>
    </tr>
    <tr>
        <td>idField</td>
        <td>data-id-field</td>
        <td>String</td>
        <td>undefined</td>
        <td>Indicate which field is an identity field.</td>
    </tr>
    <tr>
        <td>uniqueId</td>
        <td>data-unique-id</td>
        <td>String</td>
        <td>undefined</td>
        <td>Indicate an unique identifier for each row.</td>
    </tr>
    <tr>
        <td>cardView</td>
        <td>data-card-view</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to show card view table, for example mobile view.</td>
    </tr>
    <tr>
        <td>detailView</td>
        <td>data-detail-view</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to show detail view table.</td>
    </tr>
    <tr>
        <td>detailFormatter</td>
        <td>data-detail-formatter</td>
        <td>Function</td>
        <td>function(index, row, element) {<br>return '';<br>}</td>
        <td>Format your detail view when <code>detailView</code> is set to <code>true</code>. Return a String and it will be appended into the detail view cell, optionally render the element directly using the third parameter which is a jQuery element of the target cell.</td>
    </tr>
    <tr>
        <td>detailFilter</td>
        <td>data-detail-filter</td>
        <td>Function</td>
        <td>function(index, row) {<br>return true;<br>}</td>
        <td>Enable expansion per row when <code>detailView</code> is set to <code>true</code>. Return <code>true</code> and the row will be enabled for expansion, return <code>false</code> and expansion for the row will be disabled. Default function returns <code>true</code> to enable expansion for all rows.</td>
    </tr>
    <tr>
        <td>searchAlign</td>
        <td>data-search-align</td>
        <td>String</td>
        <td>'right'</td>
        <td>Indicate how to align the search input. 'left', 'right' can be used.</td>
    </tr>
    <tr>
        <td>buttonsAlign</td>
        <td>data-buttons-align</td>
        <td>String</td>
        <td>'right'</td>
        <td>Indicate how to align the toolbar buttons. 'left', 'right' can be used.</td>
    </tr>
    <tr>
        <td>toolbarAlign</td>
        <td>data-toolbar-align</td>
        <td>String</td>
        <td>'left'</td>
        <td>Indicate how to align the custom toolbar. 'left', 'right' can be used.</td>
    </tr>
    <tr>
        <td>paginationVAlign</td>
        <td>data-pagination-v-align</td>
        <td>String</td>
        <td>'bottom'</td>
        <td>Indicate how to align the pagination. 'top', 'bottom', 'both' (put the pagination on top and bottom)  can be used.</td>
    </tr>
    <tr>
        <td>paginationHAlign</td>
        <td>data-pagination-h-align</td>
        <td>String</td>
        <td>'right'</td>
        <td>Indicate how to align the pagination. 'left', 'right' can be used.</td>
    </tr>
    <tr>
        <td>paginationDetailHAlign</td>
        <td>data-pagination-detail-h-align</td>
        <td>String</td>
        <td>'left'</td>
        <td>Indicate how to align the pagination detail. 'left', 'right' can be used.</td>
    </tr>
    <tr>
        <td>paginationPreText</td>
        <td>data-pagination-pre-text</td>
        <td>String</td>
        <td>'&lsaquo;'</td>
        <td>Indicate the icon or text to be shown in the pagination detail, the previous button.</td>
    </tr>
    <tr>
        <td>paginationNextText</td>
        <td>data-pagination-next-text</td>
        <td>String</td>
        <td>'&rsaquo;'</td>
        <td>Indicate the icon or text to be shown in the pagination detail, the next button.</td>
    </tr>
    <tr>
        <td>clickToSelect</td>
        <td>data-click-to-select</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to select checkbox or radiobox when clicking rows.</td>
    </tr>
    <tr>
        <td>ignoreClickToSelectOn</td>
        <td>data-ignore-click-to-select-on</td>
        <td>Function</td>
        <td><code>{ return $.inArray(element.tagName, ['A', 'BUTTON']); }</code></td>
        <td>
        Takes one parameters:<br>
        element: the element clicked on.<br>
        Return true if the click should be ignored, false if the click should cause the row to be selected. This option is only relevant if clickToSelect is true.
        </td>
    </tr>
    <tr>
        <td>singleSelect</td>
        <td>data-single-select</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to allow checkbox selecting only one row.</td>
    </tr>
    <tr>
        <td>toolbar</td>
        <td>data-toolbar</td>
        <td>String | Node</td>
        <td>undefined</td>
        <td>
        A jQuery selector that indicates the toolbar, for example:<br>
        #toolbar, .toolbar, or a DOM node.
        </td>
    </tr>
    <tr>
        <td>buttonsToolbar</td>
        <td>data-buttons-toolbar</td>
        <td>String | Node</td>
        <td>undefined</td>
        <td>
        A jQuery selector that indicates the buttons toolbar, for example:<br>
        #buttons-toolbar, .buttons-toolbar, or a DOM node.
        </td>
    </tr>
    <tr>
        <td>checkboxHeader</td>
        <td>data-checkbox-header</td>
        <td>Boolean</td>
        <td>true</td>
        <td>False to hide check-all checkbox in header row.</td>
    </tr>
    <tr>
        <td>maintainSelected</td>
        <td>data-maintain-selected</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to maintain selected rows on change page and search.</td>
    </tr>
    <tr>
        <td>sortable</td>
        <td>data-sortable</td>
        <td>Boolean</td>
        <td>true</td>
        <td>False to disable sortable of all columns.</td>
    </tr>
    <tr>
        <td>silentSort</td>
        <td>data-silent-sort</td>
        <td>Boolean</td>
        <td>true</td>
        <td>Set <code>false</code> to sort the data silently. This options works when the sidePagination option is set to <code>server</code>.</td>
    </tr>
    <tr>
        <td>rowStyle</td>
        <td>data-row-style</td>
        <td>Function</td>
        <td>{}</td>
        <td>
        The row style formatter function, takes two parameters: <br>
        row: the row record data.<br>
        index: the row index.<br>
        Support classes or css. Example usage:<br>
<pre>
function rowStyle(row, index) {
  return {
    classes: 'text-nowrap another-class',
    css: {"color": "blue", "font-size": "50px"}
  };
}
</pre>
        </td>
    </tr>
    <tr>
        <td>rowAttributes</td>
        <td>data-row-attributes</td>
        <td>Function</td>
        <td>{}</td>
        <td>
        The row attribute formatter function, takes two parameters: <br>
        row: the row record data.<br>
        index: the row index.<br>
        Support all custom attributes.
        </td>
    </tr>
    <tr>
        <td>customSearch</td>
        <td>data-custom-search</td>
        <td>Function</td>
        <td>$.noop</td>
        <td>
        The custom search function is executed instead of built-in search function, takes one parameters: <br>
        text: the search text.<br>
        Example usage:<br>
        <pre>
        function customSearch(text) {
            //Search logic here.
            //You must use `this.data` array in order to filter the data. NO use `this.options.data`.
        }
        </pre>
        </td>
    </tr>
    <tr>
        <td>customSort</td>
        <td>data-custom-sort</td>
        <td>Function</td>
        <td>$.noop</td>
        <td>
        The custom sort function is executed instead of built-in sort function, takes two parameters: <br>
        sortName: the sort name.<br>
        sortOrder: the sort order.<br>
        Example usage:<br>
        <pre>
        function customSort(sortName, sortOrder) {
            //Sort logic here.
            //You must use `this.data` array in order to sort the data. NO use `this.options.data`.
        }
        </pre>
        </td>
    </tr>
     <tr>
        <td>locale</td>
        <td>data-locale</td>
        <td>String</td>
        <td>undefined</td>
        <td>
        Sets the locale to use (i.e. <code>'fr-CA'</code>). Locale files must be pre-loaded.
        Allows for fallback locales, if loaded, in the following order:<br>
        <ol>
        <li>First tries for the locale as specified,</li>
        <li>Then tries the locale with <code>'_'</code> translated to
        <code>'-'</code> and the region code upper cased,</li>
        <li>Then tries the the short locale code (i.e. <code>'fr'</code> instead of <code>'fr-CA'</code>),</li>
        <li>And finally will use the last locale file loaded (or the default locale if no locales loaded).</li>
        </ol>
        If left undfined or an empty string, uses the last locale loaded (or <code>'en-US'</code>
        if no locale files loaded).
        </td>
    </tr>
    <tr>
	<td>footerStyle</td>
	<td>data-footer-style</td>
	<td>Function</td>
	<td>{}</td>
	<td>
	        The footer style formatter function, takes two parameters: <br>
	        row: the row record data.<br>
	        index: the row index.<br>
	        Support classes or css. Example usage:<br>
		<pre>
		function footerStyle(value, row, index) {
		  return {
		    css: { "font-weight": "bold" }
		  };
		}
		</pre>
	</td>
    </tr>
   </tbody>
</table>

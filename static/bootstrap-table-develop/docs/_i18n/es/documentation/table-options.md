# Table options []({{ site.repo }}/blob/develop/docs/_i18n/{{ site.lang }}/documentation/table-options.md)

---

Las opciones de la tabla están definidas en `jQuery.fn.bootstrapTable.defaults`.

<table class="table"
       id="t"
       data-search="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-mobile-responsive="true">
    <thead>
    <tr>
        <th>Nombre</th>
        <th>Atributo</th>
        <th>Tipo</th>
        <th>Valor por defecto</th>
        <th>Descripción</th>
    </tr>
    </thead>
    <tbody>
    <tr>
        <td>-</td>
        <td>data-toggle</td>
        <td>String</td>
        <td>'table'</td>
        <td>Activa bootstrap table sin escribir código JavaScript.</td>
    </tr>
    <tr>
        <td>classes</td>
        <td>data-classes</td>
        <td>String</td>
        <td>'table table-hover'</td>
        <td>El nombre de la clase de la tabla.</td>
    </tr>
    <tr>
        <td>sortClass</td>
        <td>data-sort-class</td>
        <td>String</td>
        <td>undefined</td>
        <td>El nombre de la clase de los elementos td que están ordenados.</td>
    </tr>
    <tr>
        <td>height</td>
        <td>data-height</td>
        <td>Number</td>
        <td>undefined</td>
        <td>El alto de la tabla.</td>
    </tr>
    <tr>
        <td>undefinedText</td>
        <td>data-undefined-text</td>
        <td>String</td>
        <td>'-'</td>
        <td>Define el texto por defecto.</td>
    </tr>
    <tr>
        <td>striped</td>
        <td>data-striped</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True para stripe las filas.</td>
    </tr>
    <tr>
        <td>sortName</td>
        <td>data-sort-name</td>
        <td>String</td>
        <td>undefined</td>
        <td>Define cuales columnas pueden ser ordenadas.</td>
    </tr>
    <tr>
        <td>sortOrder</td>
        <td>data-sort-order</td>
        <td>String</td>
        <td>'asc'</td>
        <td>Define el método de ordenamiento, solo puede ser 'asc' o 'desc'.</td>
    </tr>
    <tr>
        <td>sortStable</td>
        <td>data-sort-stable</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True to get a stable sorting. We will add <code>_position</code> property to the row.</td>
    </tr>
    <tr>
        <td>iconsPrefix</td>
        <td>data-icons-prefix</td>
        <td>String</td>
        <td>'glyphicon'</td>
        <td>Define el nombre del icono ('glyphicon' o 'fa' para FontAwesome). Por defecto se usa 'glyphicon'. </td>
    </tr>
    <tr>
        <td>icons</td>
        <td>data-icons</td>
        <td>Object</td>
        <td>{<br/>
        &nbsp;&nbsp;refresh: 'glyphicon-refresh icon-refresh',<br/>
        &nbsp;&nbsp;toggle: 'glyphicon-list-alt icon-list-alt',<br/>
        &nbsp;&nbsp;columns: 'glyphicon-th icon-th'<br/>
        }</td>
        <td>Define los iconos que son usados para los botones de refresh, toggle y columnas.</td>
    </tr>
    <tr>
        <td>columns</td>
        <td>-</td>
        <td>Array</td>
        <td>[]</td>
        <td>El array de columnas de la tabla, vea las propiedades de las columnas para más información.</td>
    </tr>
    <tr>
        <td>data</td>
        <td>-</td>
        <td>Array</td>
        <td>[]</td>
        <td>Los datos que serán cargados.</td>
    </tr>
    <tr>
        <td>method</td>
        <td>data-method</td>
        <td>String</td>
        <td>'get'</td>
        <td>El tipo de método para hacer request de los datos remotos.</td>
    </tr>
    <tr>
        <td>url</td>
        <td>data-url</td>
        <td>String</td>
        <td>undefined</td>
        <td>Una URL para hacer request de datos en un sitio remoto.</td>
    </tr>
    <tr>
        <td>cache</td>
        <td>data-cache</td>
        <td>Boolean</td>
        <td>true</td>
        <td>False para deshabilitar los AJAX requests.</td>
    </tr>
    <tr>
        <td>contentType</td>
        <td>data-content-type</td>
        <td>String</td>
        <td>'application/json'</td>
        <td>EL contentType para hacer request de los datos.</td>
    </tr>
    <tr>
        <td>dataType</td>
        <td>data-data-type</td>
        <td>String</td>
        <td>'json'</td>
        <td>El tipo de datos que se esperan del servidor.</td>
    </tr>
    <tr>
        <td>ajaxOptions</td>
        <td>data-ajax-options</td>
        <td>Object</td>
        <td>{}</td>
        <td>Opciones adicionales para enviar ajax request. Lista de valores: <a href="http://api.jquery.com/jQuery.ajax">http://api.jquery.com/jQuery.ajax</a>.</td>
    </tr>
    <tr>
        <td>queryParams</td>
        <td>data-query-params</td>
        <td>Function</td>
        <td>function(params) {<br>return params;<br>}</td>
        <td>
        Cuando se solicita datos remotos, se debe enviar parámetros adicionales para modificar los queryParams.
        Si queryParamsType = 'limit', el objecto params contiene: <br>
        limit, offset, search, sort, order
        Sino, el objeoto contiene: <br>
        pageSize, pageNumber, searchText, sortName, sortOrder. <br>
        Retorna false para parar el request.
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
        Antes de cargar los datos remotos, manejar el formato de respuesta de los datos, el objecto parameters contiene: <br>
        res: los datos devueltos.
        </td>
    </tr>
    <tr>
        <td>pagination</td>
        <td>data-pagination</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True para mostrar la paginación al final de la tabla.</td>
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
        <td>True para mostrar solo la cantidad de los registros que se están mostrando en la tabla. Esta opción necesita que la opción pagination este en true.</td>
    </tr>
    <tr>
        <td>sidePagination</td>
        <td>data-side-pagination</td>
        <td>String</td>
        <td>'client'</td>
        <td>Define el tipo de paginación de la tabla, puede ser solo 'client' o 'server'.</td>
    </tr>
    <tr>
        <td>pageNumber</td>
        <td>data-page-number</td>
        <td>Number</td>
        <td>1</td>
        <td>Cuando se habilita la paginación, inicializa el número de página.</td>
    </tr>
    <tr>
        <td>pageSize</td>
        <td>data-page-size</td>
        <td>Number</td>
        <td>10</td>
        <td>Cuando se habilita la paginación, inicializa el tamaño de la página.</td>
    </tr>
    <tr>
        <td>pageList</td>
        <td>data-page-list</td>
        <td>Array</td>
        <td>[10, 25, 50, 100, All]</td>
        <td>Cuando se habilita la paginación, inicializa la lista de cantidad de registros por página.</td>
    </tr>
    <tr>
        <td>selectItemName</td>
        <td>data-select-item-name</td>
        <td>String</td>
        <td>'btSelectItem'</td>
        <td>El nombre del radio o del checkbox.</td>
    </tr>
    <tr>
        <td>smartDisplay</td>
        <td>data-smart-display</td>
        <td>Boolean</td>
        <td>true</td>
        <td>True para mostrar la páginación o la vista de tarjeta inteligentemente.</td>
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
        <td>Habilita el campo para búsqueda.</td>
    </tr>
    <tr>
        <td>searchOnEnterKey</td>
        <td>data-search-on-enter-key</td>
        <td>Boolean</td>
        <td>false</td>
        <td>El método será ejecutado hasta que la tecla Enter sea presionada.</td>
    </tr>
    <tr>
        <td>strictSearch</td>
        <td>data-strict-search</td>
        <td>Boolean</td>
        <td>false</td>
        <td>Habilita la busqueda exacta.</td>
    </tr>
	<tr>
        <td>searchText</td>
        <td>data-search-text</td>
        <td>String</td>
        <td>''</td>
        <td>Inicializa el campo de búsqueda con el texto especificado.</td>
    </tr>
    <tr>
        <td>searchTimeOut</td>
        <td>data-search-time-out</td>
        <td>Number</td>
        <td>500</td>
        <td>Setea el tiempo de espera para iniciar la búsqueda.</td>
    </tr>
    <tr>
        <td>showHeader</td>
        <td>data-show-header</td>
        <td>Boolean</td>
        <td>true</td>
        <td>False para ocultar el encabezado de la tabla.</td>
    </tr>
    <tr>
        <td>showFooter</td>
        <td>data-show-footer</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True para mostrar el footer.</td>
    </tr>
	<tr>
        <td>showColumns</td>
        <td>data-show-columns</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True para mostrar las columnas en una lista.</td>
    </tr>
    <tr>
        <td>showRefresh</td>
        <td>data-show-refresh</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True para mostrar el botón de refresh.</td>
    </tr>
    <tr>
        <td>showToggle</td>
        <td>data-show-toggle</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True para mostrar el botón de vista cambiada entre vista de tabla y vista de tarjeta.</td>
    </tr>
    <tr>
        <td>showPaginationSwitch</td>
        <td>data-show-pagination-switch</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True para mostrar el botón de mostrar/ocultar la paginación.</td>
    </tr>
    <tr>
        <td>showFullscreen</td>
        <td>data-show-fullscreen</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True para mostrar botón de fullscreen.</td>
    </tr>
    <tr>
        <td>minimumCountColumns</td>
        <td>data-minimum-count-columns</td>
        <td>Number</td>
        <td>1</td>
        <td>la cantidad mínima de columnas que se deben mostrar.</td>
    </tr>
    <tr>
        <td>idField</td>
        <td>data-id-field</td>
        <td>String</td>
        <td>undefined</td>
        <td>Indica cuál campo es el identificador.</td>
    </tr>
    <tr>
        <td>uniqueId</td>
        <td>data-unique-id</td>
        <td>String</td>
        <td>undefined</td>
        <td>Indica un único id para cada fila.</td>
    </tr>
    <tr>
        <td>cardView</td>
        <td>data-card-view</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True para mostrar la vista de tarjeta, por ejemplo en móviles.</td>
    </tr>
    <tr>
        <td>detailView</td>
        <td>data-detail-view</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True para mostrar la vista detalle en la tabla.</td>
    </tr>
    <tr>
        <td>detailFormatter</td>
        <td>data-detail-formatter</td>
        <td>Function</td>
        <td>function(index, row) {<br>return '';<br>}</td>
        <td>Formatee su vista detalle cuando <code>detailView</code> está seteada en <code>true</code>.</td>
    </tr>
    <tr>
        <td>searchAlign</td>
        <td>data-search-align</td>
        <td>String</td>
        <td>'right'</td>
        <td>Indica cómo alinear el campo de búsqueda. Se puede usar 'left', 'right'.</td>
    </tr>
    <tr>
        <td>buttonsAlign</td>
        <td>data-buttons-align</td>
        <td>String</td>
        <td>'right'</td>
        <td>Indica cómo alinear los botones de la barra de herramientas. Se puede usar 'left', 'right'.</td>
    </tr>
    <tr>
        <td>toolbarAlign</td>
        <td>data-toolbar-align</td>
        <td>String</td>
        <td>'left'</td>
        <td>Indica cómo alinear la barra de herramientas customizable. Se puede usar 'left', 'right'.</td>
    </tr>
    <tr>
        <td>paginationVAlign</td>
        <td>data-pagination-v-align</td>
        <td>String</td>
        <td>'bottom'</td>
        <td>Indica cómo alinear la paginación. Se puede usar: 'top', 'bottom', 'both' (coloca la paginación arriba y abajo de la tabla.</td>
    </tr>
	<tr>
        <td>paginationHAlign</td>
        <td>data-pagination-h-align</td>
        <td>String</td>
        <td>'right'</td>
        <td>Indica cómo alinear la paginación. Se puede usar: 'left', 'right'.</td>
    </tr>
    <tr>
        <td>paginationDetailHAlign</td>
        <td>data-pagination-detail-h-align</td>
        <td>String</td>
        <td>'left'</td>
        <td>Indica cómo alinear el detalle de la paginación. Se puede usar: 'left', 'right'.</td>
    </tr>
    <tr>
        <td>paginationPreText</td>
        <td>data-pagination-pre-text</td>
        <td>String</td>
        <td>'&lt;'</td>
        <td>Indica el icono o el texto a mostrar en la paginación, el botón previous del detalle de la paginación.</td>
    </tr>
    <tr>
        <td>paginationNextText</td>
        <td>data-pagination-next-text</td>
        <td>String</td>
        <td>'&gt;'</td>
        <td>Indica el icono o el texto a mostrar en la paginación, el botón next del detalle de la paginación.</td>
    </tr>
    <tr>
        <td>clickToSelect</td>
        <td>data-click-to-select</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True para seleccionar el checkbox o el radiobox cuando se da click sobre las filas.</td>
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
        <td>True para permirir solo un checkbox seleccionado.</td>
    </tr>
    <tr>
        <td>toolbar</td>
        <td>data-toolbar</td>
        <td>String</td>
        <td>undefined</td>
        <td>Un selector jQuery que indica la barra de herramientas, por ejemplo:<br> #toolbar, .toolbar.</td>
    </tr>
    <tr>
        <td>buttonsToolbar</td>
        <td>data-buttons-toolbar</td>
        <td>String | Node</td>
        <td>undefined</td>
        <td>
        Un selector jQuery que indica la barra de herramientas de botones, por ejemplo: <br>
        #buttons-toolbar, .buttons-toolbar, o un nodo DOM.
        </td>
    </tr>
    <tr>
        <td>checkboxHeader</td>
        <td>data-checkbox-header</td>
        <td>Boolean</td>
        <td>true</td>
        <td>False para ocular el checkbox check-all en el encabezado de la fila.</td>
    </tr>
    <tr>
        <td>maintainSelected</td>
        <td>data-maintain-selected</td>
        <td>Boolean</td>
        <td>false</td>
        <td>True para mantener las columnas después de seleccionar o cambiar entre páginas.</td>
    </tr>
    <tr>
        <td>sortable</td>
        <td>data-sortable</td>
        <td>Boolean</td>
        <td>true</td>
        <td>False para deshabilitar el ordenamiento en todas las columnas.</td>
    </tr>
    <tr>
        <td>silentSort</td>
        <td>data-silent-sort</td>
        <td>Boolean</td>
        <td>true</td>
        <td>Setear a <code>false</code> para ordenar los datos silenciosamente. Esta opción funciona cuando la opción sidePagination es seteada a <code>server</code>.</td>
    </tr>
    <tr>
        <td>rowStyle</td>
        <td>data-row-style</td>
        <td>Function</td>
        <td>{}</td>
        <td>
        La función formatter para aplicar el estilo de la fila, toma dos parámetros: <br>
        row: los datos de la fila.<br>
        index: el índice de la fila.<br>
        Soporta clases y CSS.
        </td>
    </tr>
    <tr>
        <td>rowAttributes</td>
        <td>data-row-attributes</td>
        <td>Function</td>
        <td>{}</td>
        <td>
        La función formatter para los atributos de la fiila, toma dos parámetros: <br>
        row: los datos de la fila.<br>
        index: el índice de la fila.<br>
        Soporta cualquier atributo customizable.
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
    </tbody>
</table>

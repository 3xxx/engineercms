 $(document).ready(function(){
  // Coordinates to center the map. Could let the user choose when creating a room & persist when sharing a link (via GET params)
  // const lat = 51.52;
  // const lon = -0.09;
  const lat = 23.125178;
  const lon = 113.280637;

  // Initialize the Leaflet map
  var map = L.map('mapDiv', {
    crs: L.CRS.Baidu,
    // minZoom: 3,
    maxZoom: 18,
    // attributionControl: false,
    // center: [31.834912, 117.220102],
    // zoom: 12,
    // zoomControl: false, //关闭左上角放大缩小
    doubleClickZoom: false, // 关闭双击放大

    renderer: L.canvas({ tolerance: 10 })
  }).setView([lat, lon], 12);
  L.PM.setOptIn(true);

  // var db = firebase.database();
  var mapname = "";
  var oldname = "";
  var mapdescription = "";
  var olddescription = "";
  var editingname = false;
  var editingdescription = false;
  var dragging = false;
  var enteringdata = false;
  var cursorcoords = [0,0];
  var session = 0;
  var drawing = false;
  var erasing = false;
  var markerson = false;
  var lineon = false;
  var linelastcoord = [0,0];
  var observing = {status:false, id:0};
  var linedistance = 0;
  var mousedown = false;
  var objects = [];
  var currentid = 0;
  var color = "#634FF1";
  var cursors = [];
  var userlocation = "";
  var places = [];
  var place_ids = [];
  var room = "";
  // var loginUser = new Object();
  var user = new Object();

  // Available cursor colors
  var colors = ["#EC1D43", "#EC811D", "#ECBE1D", "#B6EC1D", "#1DA2EC", "#781DEC", "#CF1DEC", "#222222"];

  // Get URL params
  var params = new URLSearchParams(window.location.search);

  // Check if URL has the file GET parameter, use it to set the room. Could rewrite URL to be more fancy
  if (params.has('file')) {
    room = params.get('file');
    $("#share-url").val(window.location.href);
  }

  // Oddly enough Firebase auth doesn't initialize right on startup. It needs a slight delay?
  window.setTimeout(async function(){
    if (await checkAuth() && params.has('file')) {
      checkData();
    } else {
      if (await checkAuth() && !params.has('file')) {
        // Prompt the user to create a map
        $("#popup").find(".header-text").html("Create a map");
        $("#popup").find(".subheader-text").html("Maps can be shared with friends to collaborate in real-time.");
        $("#google-signin").attr("id", "create-map");
        $("#create-map").html("Create a map");
      }
      // Show popup & overlay
      $("#overlay").addClass("signin");
      $("#popup").addClass("signin");
    }
  }, 500)

  function initMap() {
    // Makimum bounds for zooming and panning
    // map.setMaxBounds([[84.67351256610522, -174.0234375], [-58.995311187950925, 223.2421875]]);

    // Set the tile layer. Could use Mapbox, OpenStreetMap, etc.
    // L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    //   maxZoom: 18,
    //   zoomControl: false,
    //   minZoom:3,
    //   noWrap: true
    // }).addTo(map);

    //控制地图底图
    L.control
      .layers(
        {
          百度地图: L.tileLayer.baidu({ layer: "vec"}).addTo(map),
          百度卫星: L.tileLayer.baidu({ layer: "img" }),
          "百度地图-大字体": L.tileLayer.baidu({
              layer: "vec",
              bigfont: true,
          }),
          "百度卫星-大字体": L.tileLayer.baidu({
              layer: "img",
              bigfont: true,
          }),
          "自定义样式-黑色地图": L.tileLayer.baidu({
              layer: "custom",
              customid: "dark",
          }),
          "自定义样式-蓝色地图": L.tileLayer.baidu({
              layer: "custom",
              customid: "midnight",
          }), //自定义样式地图，customid可选值：dark,midnight,grayscale,hardedge,light,redalert,googlelite,grassgreen,pink,darkgreen,bluish
        },
        {
            实时交通信息: L.tileLayer.baidu({ layer: "time" }),
        },
        { position: "topright" }
      )
    .addTo(map);
    // test
    // new L.marker([31.839177, 117.232039]).addTo(map);

    // Hide the default zoom control. I want a custom one!
    map.removeControl(map.zoomControl);

    // No idea why but Leaflet seems to place default markers on startup...
    $("img[src='https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon-2x.png']").remove();
  }

  // Hints for drawing lines or polygons
  var followcursor = L.marker([0, 0], {pane: "overlayPane", interactive:false}).addTo(map);
  followcursor.setOpacity(0);
  var tooltip = followcursor.bindTooltip("", { permanent: true, offset:[5,25], sticky: true, className: "hints", direction:"right"}).addTo(map);
  followcursor.closeTooltip();

  // Show live location
  function liveLocation() {
    if (navigator.geolocation) {
      // Get initial location
      navigator.geolocation.getCurrentPosition(function(position){
        var icon = L.icon({
          iconUrl: 'assets/liveLocation.svg',
          iconSize:     [24, 24],
          iconAnchor:   [12, 12],
        });
        // Create a marker to show the user location
        userlocation = L.marker([position.coords.latitude, position.coords.longitude], {icon:icon, pane: "overlayPane"});
        userlocation.addTo(map);
      });
    }
  }

  function targetLiveLocation() {
    stopObserving();

    // Check if user has geolocation enabled一定要在https下才能有效果
    if (navigator.geolocation) {
      if (userlocation != "") {
        // If current location is already set, fly there
        navigator.geolocation.getCurrentPosition(function(position){
          userlocation.setLatLng([position.coords.latitude, position.coords.longitude]);

          // Flies to the location (more fancy)
          map.flyTo(userlocation.getLatLng(), 18)
        });
      } else {
        // If the location is unknown, set it and fly there
        liveLocation();
        targetLiveLocation();
      }
    }
  }

  // Tooltips for UI elements
  function showTooltip() {
    if ($(this).attr("id") == "cursor-tool") {
      $(this).append('<div id="tooltip">Move (V)</div>');
    } else if ($(this).attr("id") == "pen-tool") {
      $(this).append('<div id="tooltip">Pencil (P)</div>');
    } else if ($(this).attr("id") == "eraser-tool") {
      $(this).append('<div id="tooltip">Eraser (E)</div>');
    } else if ($(this).attr("id") == "marker-tool") {
      $(this).append('<div id="tooltip">Marker (M)</div>');
    } else if ($(this).attr("id") == "area-tool") {
      $(this).append('<div id="tooltip">Area (A)</div>');
    } else if ($(this).attr("id") == "path-tool") {
      $(this).append('<div id="tooltip">Line (L)</div>');
    }
  }
  function hideTooltip() {
    $(this).find("#tooltip").remove();
  }

  // Reset tools (when switching tools)
  function resetTools() {
    drawing = false;
    erasing = false;
    markerson = false;
    lineon = false;
    map.pm.disableDraw();
    map.pm.disableGlobalRemovalMode();
    map.pm.disableGlobalDragMode();
  }

  // Enable cursor tool (default)
  function cursorTool() {
    resetTools();
    map.dragging.enable();
    $(".tool-active").removeClass("tool-active");
    $("#cursor-tool").addClass("tool-active");
  }

  // Enable pen tool
  function penTool() {
    resetTools();
    drawing = true;
    map.dragging.disable();
    $(".tool-active").removeClass("tool-active");
    $("#pen-tool").addClass("tool-active");
    showAnnotations();
  }

  // Enable eraser tool
  function eraserTool() {
    resetTools();
    erasing = true;
    $(".tool-active").removeClass("tool-active");
    $("#eraser-tool").addClass("tool-active");
    map.pm.enableGlobalRemovalMode();
    showAnnotations();
  }

  // Enable marker tool
  function markerTool() {
    resetTools();
    markerson = true;
    $(".tool-active").removeClass("tool-active");
    $("#marker-tool").addClass("tool-active");
    showAnnotations();
  }

  // Enable area tool
  function areaTool() {
    resetTools();
    $(".tool-active").removeClass("tool-active");
    $("#area-tool").addClass("tool-active");

    // Start creating an area
    map.pm.setGlobalOptions({ pinning: true, snappable: true });
    map.pm.setPathOptions({
      color: color,
      fillColor: color,
      fillOpacity: 0.4,
    });
    map.pm.enableDraw('Polygon', {
      tooltips: false,
      snappable: true,
      templineStyle: {color: color},
      hintlineStyle: {color: color, dashArray: [5, 5]},
      pmIgnore: false
    });
    showAnnotations();
  }

  // Enable line tool
  function pathTool() {
    resetTools();
    $(".tool-active").removeClass("tool-active");
    $("#path-tool").addClass("tool-active");

    // Start creating a line
    map.pm.setGlobalOptions({ pinning: true, snappable: true });
    map.pm.setPathOptions({
      color: color,
      fillColor: color,
      fillOpacity: 0.4,
    });
    map.pm.enableDraw('Line', {
      tooltips: false,
      snappable: true,
      templineStyle: {color: color},
      hintlineStyle: {color: color, dashArray: [5, 5]},                                                                                                                                                                    
      pmIgnore: false,
      finishOn: 'dblclick',
    });
    showAnnotations();
  }

  // Show/hide color picker
  function toggleColor() {
    $("#color-list").toggleClass("color-enabled");
  }

  // Switch color (color picker)
  function switchColor(e) {
    e.stopPropagation();
    color = $(this).attr("data-color");
    $("#inner-color").css({background:color});
    toggleColor();
  }

  // Sanitizing input strings
  function sanitize(string) {
    const map = {
        '&': '&amp;',
        '<': '&lt;',
        '>': '&gt;',
        '"': '&quot;',
        "'": '&#x27;',
        "/": '&#x2F;',
    };
    const reg = /[&<>"'/]/ig;
    return string.replace(reg, (match)=>(map[match]));
  }

  // Perform a search query. No autocomplete due to API rules :(
  function search() {
    $.get('https://nominatim.openstreetmap.org/search?q='+sanitize($("#search-box input").val())+'&format=json', function(data) {
      // Navigate to the first result of the search query
      map.panTo(new L.LatLng(data[0].lat, data[0].lon));
    })
  }

  // Find nearby
  async function findNearby() {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      var locationtype = $(this).attr("data-type");
      var markercolor = $(this).attr("data-color");
      var coordinates = map.getBounds().getNorthWest().lng+','+map.getBounds().getNorthWest().lat+','+map.getBounds().getSouthEast().lng+','+map.getBounds().getSouthEast().lat;

      // Call Nominatim API to get places nearby the current view, of the amenity that the user has selected
      $.get('https://nominatim.openstreetmap.org/search?viewbox='+coordinates+'&format=geocodejson&limit=20&bounded=1&amenity='+locationtype+'&exclude_place_ids='+JSON.stringify(place_ids), function(data) {
        // Custom marker icon depending on the amenity
        var marker_icon = L.icon({
          iconUrl: 'assets/'+locationtype+'-marker.svg',
          iconSize:     [30, 30],
          iconAnchor:   [15, 30],
          shadowAnchor: [4, 62],
          popupAnchor:  [-3, -76]
        });
        data.features.forEach(function(place){
          // Create a marker for the place
          var marker = L.marker([place.geometry.coordinates[1], place.geometry.coordinates[0]], {icon:marker_icon, pane:"overlayPane", interactive:true}).addTo(map);

          // Create a popup with information about the place, and the options to save it or delete it (it's only local for now)
          marker.bindTooltip('<h1>'+place.properties.geocoding.name+'</h1><div class="shape-data"><h3><img src="assets/marker-small-icon.svg">'+place.geometry.coordinates[1].toFixed(5)+', '+place.geometry.coordinates[0].toFixed(5)+'</h3></div><br><div id="buttons2"><button class="cancel-button-place" data-id='+place.properties.geocoding.place_id+'>Remove</button><button class="save-button-place" data-id='+place.properties.geocoding.place_id+'>Save</button></div><div class="arrow-down"></div>', {permanent: false, direction:"top", interactive:false, bubblingMouseEvents:false, className:"create-shape-flow", offset: L.point({x: 0, y: -35})});
          places.push({id: "", place_id:place.properties.geocoding.place_id, user:user.id, name:place.properties.geocoding.name, desc:"", lat:place.geometry.coordinates[1], lng:place.geometry.coordinates[0], trigger:marker, completed:true, marker:marker, m_type:locationtype, type:"marker", session:session, color:markercolor});
          place_ids.push(place.properties.geocoding.place_id);
        });
      });
    // }
  }

  // Save nearby location (share with other users, they all will see it)
  async function saveNearby() {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      var inst = places.find(x => x.place_id == $(this).attr("data-id"));
      const { data,error } = await _supabase
        .from('objects')
        .insert({
          color: inst.color,
          place_id: inst.place_id,
          lat: inst.lat,
          lng: inst.lng,
          user: user.id,
          type: "marker",
          m_type: inst.m_type,
          session: session,
          name: inst.name,
          desc: ""
        }).select('id');
      // console.log('data', data)
      // console.log('data', data[0].id)
      currentid=data[0].id

      var key = currentid;
      inst.id = currentid;
      objects.push(inst);

      // Create a popup with information about the place
      inst.marker.bindTooltip('<h1>'+inst.name+'</h1><div class="shape-data"><h3><img src="assets/marker-small-icon.svg">'+inst.lat.toFixed(5)+', '+inst.lng.toFixed(5)+'</h3></div><br><div class="arrow-down"></div>', {permanent: false, direction:"top", interactive:false, bubblingMouseEvents:false, className:"create-shape-flow", offset: L.point({x: 0, y: -35})});
    // }
  }

  // Remove nearby location
  function cancelNearby() {
    var inst = places.find(x => x.place_id == $(this).attr("data-id"));
    inst.marker.remove();
    places = $.grep(places, function(e){
         return e.place_id != inst.id;
    });
    place_ids = $.grep(place_ids, function(e){
         return e != inst.id;
    });
  }

  // Enable observation mode
  async function observationMode() {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      var otheruser = $(this).attr("data-id");
      if (otheruser != user.id) {
        if (observing.id == otheruser) {
          // When clicking on the avatar of the current user you're observing, stop observing them
          stopObserving();
        } else {
          // Start observing the selected user
          observing.status = true;
          observing.id = otheruser;

          // Show that observation mode is enabled
          $("#outline").css({"border": "6px solid "+cursors.find(x => x.id === otheruser).color});
          $("#observing-name").html("Observing "+cursors.find(x => x.id === otheruser).name);
          $("#observing-name").css({"background": cursors.find(x => x.id === otheruser).color});
          $("#outline").addClass("observing");
        }
      }
    // }
  }

  // Disable observation mode
  function stopObserving() {
    observing.status = false;
    $("#outline").css({"border": "none"});
    $("#outline").removeClass("observing");
  }

  async function saveForm(e) {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      enteringdata = false;
      var inst = objects.filter(function(result){
        return result.id === currentid && result.user === user.id;
      })[0];

      // Get input values for the name and description and sanitize them
      inst.name = sanitize($("#shape-name").val());
      inst.desc = sanitize($("#shape-desc").val());
      inst.completed = true;

      // Remove existing popup (for inputting data)
      inst.trigger.unbindTooltip();
      if (inst.type == "area") {
        // Create a popup showing info about the area
        inst.trigger.bindTooltip('<h1>'+inst.name+'</h1><h2>'+inst.desc+'</h2><div class="shape-data"><h3><img src="assets/area-icon.svg">'+inst.area+' km&sup2;</h3></div><div class="shape-data"><h3><img src="assets/perimeter-icon.svg">'+inst.distance+' km</h3></div><div class="arrow-down"></div>', {permanent: false, direction:"top", interactive:false, bubblingMouseEvents:false, className:"create-shape-flow", offset: L.point({x: -15, y: 18})});
        // db.ref("rooms/"+room+"/objects/"+currentid).update({
        //   area:inst.area,
        //   distance: inst.distance,
        //   name: inst.name,
        //   desc: inst.desc,
        //   completed: true
        // })
        const { data, error } = await _supabase
          .from('objects')
          .update({ area:inst.area,distance: inst.distance,name: inst.name,desc: inst.desc,completed: true })
          // .eq('room', room)
          .eq('id', currentid)
          .select()
      } else if (inst.type == "line") {
        // Create a popup showing info about the line
        inst.trigger.bindTooltip('<h1>'+inst.name+'</h1><h2>'+inst.desc+'</h2><div class="shape-data"><h3><img src="assets/distance-icon.svg">'+inst.distance+' km</h3></div><div class="arrow-down"></div>', {permanent: false, direction:"top", interactive:false, bubblingMouseEvents:false, className:"create-shape-flow", offset: L.point({x: -15, y: 18})});
        // db.ref("rooms/"+room+"/objects/"+currentid).update({
        //   distance: inst.distance,
        //   name: inst.name,
        //   desc: inst.desc,
        //   completed: true
        // })
        const { data, error } = await _supabase
          .from('objects')
          .update({ distance: inst.distance,name: inst.name,desc: inst.desc,completed: true })
          // .eq('room', room)
          .eq('id', currentid)
          .select()
      } else if (inst.type == "marker") {
        // Create a popup showing info about the marker
        inst.trigger.bindTooltip('<h1>'+inst.name+'</h1><h2>'+inst.desc+'</h2><div class="shape-data"><h3><img src="assets/marker-small-icon.svg">'+inst.lat.toFixed(5)+', '+inst.lng.toFixed(5)+'</h3></div><div class="arrow-down"></div>', {permanent: false, direction:"top", interactive:false, bubblingMouseEvents:false, className:"create-shape-flow", offset: L.point({x: 0, y: -35})});
        // db.ref("rooms/"+room+"/objects/"+currentid).update({
        //   name: inst.name,
        //   desc: inst.desc,
        //   completed: true
        // })
        const { data, error } = await _supabase
          .from('objects')
          .update({ name: inst.name,desc: inst.desc,completed: true })
          // .eq('room', room)
          .eq('id', currentid)
          .select()
      }

      // Render the shape in the sidebar list and focus it
      renderObjectLayer(inst);
      $(".annotation-item[data-id='"+inst.id+"']").find(".annotation-name span").addClass("annotation-focus");

      // Automatically open the new popup with data about the shape
      window.setTimeout(function(){
        inst.trigger.openTooltip();
      }, 200)
    // }
  }

  async function cancelForm() {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      enteringdata = false;
      var inst = objects.filter(function(result){
        return result.id === currentid && result.user === user.id;
      })[0];

      // Delete existing popup (for inputting data)
      inst.trigger.unbindTooltip();
      inst.completed = true;
      if (inst.type == "area") {
        // Create a popup showing info about the area
        inst.trigger.bindTooltip('<h1>'+inst.name+'</h1><h2>'+inst.desc+'</h2><div class="shape-data"><h3><img src="assets/area-icon.svg">'+inst.area+' km&sup2;</h3></div><div class="shape-data"><h3><img src="assets/perimeter-icon.svg">'+inst.distance+' km</h3></div><div class="arrow-down"></div>', {permanent: false, direction:"top", interactive:false, bubblingMouseEvents:false, className:"create-shape-flow", offset: L.point({x: -15, y: 18})});
        // db.ref("rooms/"+room+"/objects/"+currentid).update({
        //   area:inst.area,
        //   distance: inst.distance,
        //   name: inst.name,
        //   desc: inst.desc,
        //   completed: true
        // })
        const { data, error } = await _supabase
          .from('objects')
          .update({ area:inst.area,distance: inst.distance,name: inst.name,desc: inst.desc,completed: true })
          // .eq('room', room)
          .eq('id', currentid)
          .select()
      } else if (inst.type == "line") {
        // Create a popup showing info about the line
        inst.trigger.bindTooltip('<h1>'+inst.name+'</h1><h2>'+inst.desc+'</h2><div class="shape-data"><h3><img src="assets/distance-icon.svg">'+inst.distance+' km</h3></div><div class="arrow-down"></div>', {permanent: false, direction:"top", interactive:false, bubblingMouseEvents:false, className:"create-shape-flow", offset: L.point({x: -15, y: 18})});
        // db.ref("rooms/"+room+"/objects/"+currentid).update({
        //   distance: inst.distance,
        //   name: inst.name,
        //   desc: inst.desc,
        //   completed: true
        // })
        const { data, error } = await _supabase
          .from('objects')
          .update({ distance: inst.distance,name: inst.name,desc: inst.desc,completed: true })
          // .eq('room', room)
          .eq('id', currentid)
          .select()
      } else if (inst.type == "marker") {
        // Create a popup showing info about the marker
        inst.trigger.bindTooltip('<h1>'+inst.name+'</h1><h2>'+inst.desc+'</h2><div class="shape-data"><h3><img src="assets/marker-small-icon.svg">'+inst.lat.toFixed(5)+', '+inst.lng.toFixed(5)+'</h3></div><div class="arrow-down"></div>', {permanent: false, direction:"top", interactive:false, bubblingMouseEvents:false, className:"create-shape-flow", offset: L.point({x: 0, y: -35})});
        // db.ref("rooms/"+room+"/objects/"+currentid).update({
        //   name: inst.name,
        //   desc: inst.desc,
        //   completed: true
        // })
        const { data, error } = await _supabase
          .from('objects')
          .update({ name: inst.name,desc: inst.desc,completed: true })
          // .eq('room', room)
          .eq('id', currentid)
          .select()
      }

      // Render shape in the sidebar list and focus it
      renderObjectLayer(inst);
      $(".annotation-item[data-id='"+inst.id+"']").find(".annotation-name span").addClass("annotation-focus");

      // Automatically open the new popup with data about the shape
      window.setTimeout(function(){
        inst.trigger.openTooltip();
      }, 200)
    // }
  }

  // Start drawing lines/areas
  map.on('pm:drawstart',async ({ workingLayer }) => {
    // var user = await checkAuth();
    // console.log(user)
    // if (await checkAuth() != false) {
      // Show hints for drawing lines/areas
      followcursor.openTooltip();
      followcursor.setTooltipContent("Click to place first vertex");

      // Detect when a vertex is added to a line or area
      workingLayer.on('pm:vertexadded', async e => {
        if (e.shape == "Polygon") {
          // Update hints
          followcursor.setTooltipContent("Click on first vertex to finish");
          linelastcoord = e.layer._latlngs[e.layer._latlngs.length-1];
          if (e.layer._latlngs.length == 1) {
            // If this is the first vertex, get a key and add the new shape in the database
            const { data,error } = await _supabase
              .from('objects')
              .insert({
                room: room,
                // id: currentid,
                color: color,
                initlat: e.layer._latlngs[0].lat,
                initlng: e.layer._latlngs[0].lng,
                user: user.id,
                type: "area",
                session: session,
                name: "Area",
                desc: "",
                distance: 0,
                area: 0,
                completed: false,
                path: []
              }).select('id');
            // console.log('data', data)
            // console.log('data', data[0].id)
            currentid=data[0].id

            const { error:error1 } = await _supabase
            .from('coords')
            .insert({
              objects_id: data[0].id,
              // id: coordsid
              lat: linelastcoord.lat,
              lng: linelastcoord.lng
            });

            objects.push({id:currentid, local:true, color:color, user:user.id, name:"Area", desc:"", trigger:"", distance:0, area:0, layer:"", type:"area", session:session, completed:false});
          } else {
            // If this is not the first vertex, update the data in the database with the latest coordinates
            // db.ref("rooms/"+room+"/objects/"+currentid+"/coords/").push({
            //   set:[linelastcoord.lat,linelastcoord.lng]
            // })
            const { error } = await _supabase
            .from('coords')
            .insert({
              objects_id: currentid,
              // id: coordsid
              lat: linelastcoord.lat,
              lng: linelastcoord.lng
            });
          }
        } else if (e.shape == "Line") {
          lineon = true;
          linedistance = 0;
          linelastcoord = e.layer._latlngs[e.layer._latlngs.length-1];
          if (e.layer._latlngs.length == 1) {
            // If this is the first vertex, get a key and add the new shape in the database
            const { data,error } = await _supabase
            .from('objects')
            .insert({
              room: room,
              // id: currentid,
              color: color,
              initlat: e.layer._latlngs[0].lat,
              initlng: e.layer._latlngs[0].lng,
              user: user.id,
              type: "line",
              session: session,
              name: "Line",
              desc: "",
              distance: 0,
              completed: false,
              path: []
            }).select('id');
            // console.log('data', data)
            // console.log('data', data[0].id)
            currentid=data[0].id

            const { error:error1 } = await _supabase
            .from('coords')
            .insert({
              objects_id: data[0].id,
              // id: coordsid
              lat: linelastcoord.lat,
              lng: linelastcoord.lng
            });
            objects.push({id:currentid, local:true, color:color, user:user.id, name:"Line", desc:"", trigger:"", distance:0, layer:"", type:"line", session:session, completed:false});
          } else {
            // If this is not the first vertex, update hints to show total distance drawn
            e.layer._latlngs.forEach(function(coordinate, index){
              if (index != 0) {
                linedistance += e.layer._latlngs[index-1].distanceTo(coordinate);
              }
            });
            followcursor.setTooltipContent((linedistance/1000)+"km | Double click to finish");

            // Save new vertext in the database
            // db.ref("rooms/"+room+"/objects/"+currentid+"/coords/").push({
            //   set:[linelastcoord.lat,linelastcoord.lng]
            // })
            const { error } = await _supabase
            .from('coords')
            .insert({
              objects_id: currentid,
              // id: coordsid
              lat: linelastcoord.lat,
              lng: linelastcoord.lng
            });
          }
        }
      });
    // }
  });

  // Stop drawing lines / polygons
  map.on('pm:drawend', e => {
    lineon = false;
    followcursor.closeTooltip();
    cursorTool();
  });

  // Add tooltip to lines and polygons
  map.on('pm:create', async e => {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      enteringdata = true;
      var inst = objects.filter(function(result){return result.id === currentid && result.user === user.id;})[0];
      // Calculate total distance / perimeter
      inst.distance = parseFloat(turf.length(e.layer.toGeoJSON()).toFixed(2));//注释这里
      inst.layer = e.layer;//注释这里
      if (inst.type == "area") {
        // Calculate area
        inst.area = parseFloat((turf.area(e.layer.toGeoJSON())*0.000001).toFixed(2));

        // Save all the area coordinates
        var temppath = [];
        Object.values(e.layer.getLatLngs()[0]).forEach(function(a){
          temppath.push([Object.values(a)[0], Object.values(a)[1]]);
        })

        // Update the data in the database
        const { data, error } = await _supabase
          .from('objects')
          .update({ path: temppath, area:inst.area })
          // .eq('room', room)
          .eq('id', currentid)
          .select()

      } else if (inst.type == "line") {
        // Save all the line coordinates
        var temppath = [];
        Object.values(e.layer.getLatLngs()).forEach(function(a){
          temppath.push([Object.values(a)[0], Object.values(a)[1]]);
        })

        // Update the data in the database
        const { data, error } = await _supabase
          .from('objects')
          .update({ path: temppath })
          // .eq('room', room)
          .eq('id', currentid)
          .select()
      }

      // Create a marker so it can trigger a popup when clicking on a line, or area
      var centermarker = L.marker(e.layer.getBounds().getCenter(), {zIndexOffset:9999, interactive:false, pane:"overlayPane"});

      // Create a popup so users can name and give a description to the shape
      centermarker.bindTooltip('<label for="shape-name">Name</label><input value="'+inst.name+'" id="shape-name" name="shape-name" /><label for="shape-desc">Description</label><textarea id="shape-desc" name="description"></textarea><br><div id="buttons"><button class="cancel-button">Cancel</button><button class="save-button">Save</button></div><div class="arrow-down"></div>', {permanent: true, direction:"top", interactive:true, bubblingMouseEvents:false, className:"create-shape-flow create-form", offset: L.point({x: -15, y: 18})});

      // The marker is supposed to be hidden, it's just for placing the tooltip on the map and triggering it
      centermarker.setOpacity(0);
      centermarker.addTo(map);
      centermarker.openTooltip();

      // Automatically select the name so it's faster to edit
      $("#shape-name").focus();
      $("#shape-name").select();
      // console.log(objects)
      inst.trigger = centermarker;// 为什么这步后，objects里的当前对象的trigger也自动有了值？？
      // console.log(objects)
      // Detect when clicking on the shape
      e.layer.on("click", async function(e){
        if (!erasing) {
          // Set the popup to the mouse coordinates and open it
          centermarker.setLatLng(cursorcoords);
          centermarker.openTooltip();
        } else {
          // If erasing, delete the shape
          inst.trigger.remove();
          e.layer.remove();
          // db.ref("rooms/"+room+"/objects/"+inst.id).remove();
          const { error } = await _supabase
              .from('coords')
              .delete()
              .eq('objects_id', inst.id)

          const { error:error2 } = await _supabase
            .from('objects')
            .delete()
            .eq('id', inst.id)
          objects = $.grep(objects, function(e){
               return e.id != inst.id;
          });
          $(".annotation-item[data-id='"+inst.id+"']").remove();
        }
      });

      // Detect when closing the popup (e.g. when clicking outside of it)
      centermarker.on('tooltipclose', function(e){
        if (enteringdata) {
          // If closing the popup before a name and description has been set, revert to defaults
          cancelForm();
        }

        // De-select the object from the sidebar list
        $(".annotation-item[data-id="+inst.id+"]").find(".annotation-name span").removeClass("annotation-focus");
      });
    // }
  });

  // Start free drawing
  async function startDrawing(lat,lng,user) {
    var line = L.polyline([[lat,lng]], {color: color});
    // Create a new key for the line object, and set initial data in the database
    const { data,error } = await _supabase
      .from('objects')
      .insert({
        room: room,
        color: color,
        initlat: lat,
        initlng: lng,
        user: user.id,
        type: "draw",
        session: session,
        completed: true
      }).select('id');
      // console.log('data', data)
      // console.log('data', data[0].id)
      currentid=data[0].id

      const { error:error1 } = await _supabase
      .from('coords')
      .insert({
        objects_id: data[0].id,
        // id: coordsid
        lat: lat,
        lng: lng
      });

    // Save an object with all the defaults
    objects.push({id:currentid, user:user.id, line:line, session:session, local:true, completed:true, type:"draw"});
    line.addTo(map);

    // Event handling for lines
    objects.forEach(function(inst){
      inst.line.on("click", async function(event){
        if (erasing) {
          inst.line.remove();
          // db.ref("rooms/"+room+"/objects/"+inst.id).remove();
          const { error } = await _supabase
              .from('coords')
              .delete()
              .eq('objects_id', inst.id)

          const { error:error2 } = await _supabase
            .from('objects')
            .delete()
            .eq('id', inst.id)
          objects = $.grep(objects, function(e){
               return e.id != inst.id;
          });
          $(".annotation-item[data-id='"+inst.id+"']").remove();
        }
      });
      inst.line.on("mouseover", function(event){
        if (erasing) {
          inst.line.setStyle({opacity: .3});
        }
      });
      inst.line.on("mouseout", function(event){
        inst.line.setStyle({opacity: 1});
      });
    });
  }

  // Create a new marker
  async function createMarker(lat, lng, user) {
    if (markerson) {
      // Go back to cursor tool after creating a marker
      cursorTool();

      // Set custom marker icon
      var marker_icon = L.divIcon({
        html: '<svg width="30" height="30" viewBox="0 0 46 46" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M23 44.0833C23 44.0833 40.25 32.5833 40.25 19.1666C40.25 14.5916 38.4326 10.204 35.1976 6.96903C31.9626 3.73403 27.575 1.91663 23 1.91663C18.425 1.91663 14.0374 3.73403 10.8024 6.96903C7.56741 10.204 5.75 14.5916 5.75 19.1666C5.75 32.5833 23 44.0833 23 44.0833ZM28.75 19.1666C28.75 22.3423 26.1756 24.9166 23 24.9166C19.8244 24.9166 17.25 22.3423 17.25 19.1666C17.25 15.991 19.8244 13.4166 23 13.4166C26.1756 13.4166 28.75 15.991 28.75 19.1666Z" fill="'+color+'"/>/svg>',
        iconSize:     [30, 30], // size of the icon
        iconAnchor:   [15, 30], // point of the icon which will correspond to marker's location
        shadowAnchor: [4, 62],  // the same for the shadow
        popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
      });
      var marker = L.marker([lat, lng], {icon:marker_icon, direction:"top", interactive:true, pane:"overlayPane"});

      // Create a popup to set the name and description of the marker
      marker.bindTooltip('<label for="shape-name">Name</label><input value="Marker" id="shape-name" name="shape-name" /><label for="shape-desc">Description</label><textarea id="shape-desc" name="description"></textarea><br><div id="buttons"><button class="cancel-button">Cancel</button><button class="save-button">Save</button></div><div class="arrow-down"></div>', {permanent: true, direction:"top", interactive:false, bubblingMouseEvents:false, className:"create-shape-flow create-form", offset: L.point({x: 0, y: -35})});
      marker.addTo(map);
      marker.openTooltip();

      // Create a new key for the marker, and add it to the database
      const { data,error } = await _supabase
        .from('objects')
        .insert({
          room:room,
          color: color,
          lat: lat,
          lng: lng,
          user: user.id,
          type: "marker",
          m_type: "none",
          session: session,
          name: "Marker",
          desc: ""
        }).select('id');
        // console.log('data', data)
        // console.log('data', data[0].id)
        currentid=data[0].id
        var key = currentid;

      objects.push({id:currentid, user:user.id, color:color, name:"Marker", m_type:"none",  desc:"", lat:lat, lng:lng, marker:marker, trigger:marker, session:session, completed:true, type:"marker"});

      // Detect when the tooltip is closed
      marker.on('tooltipclose', function(e){
        if (enteringdata) {
          // If closing the tooltip but the name and description haven't been set yet, revert to defaults
          cancelForm();
        } else {
          // De-select object from sidebar
          $(".annotation-item[data-id="+key+"]").find(".annotation-name span").removeClass("annotation-focus");
        }
      });

      // Detect when the marker is clicked
      marker.on('click', async function(e){
        if (!erasing) {
          // Open tooltip when the marker is clicked
          marker.openTooltip();
        } else {
          // If erasing, delete the marker
          marker.remove();
          // db.ref("rooms/"+room+"/objects/"+inst.id).remove();
          // const { error } = await _supabase
          //   .from('coords')
          //   .delete()
          //   .eq('objects_id', inst.id)
          const { error:error2 } = await _supabase
            .from('objects')
            .delete()
            .eq('id', inst.id)

          objects = $.grep(objects, function(e){
               return e.id != key;
          });
        }
      })
    }
  }

  // Map events
  map.addEventListener('mousedown', async (event) => {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      mousedown = true;
      // Get mouse coordinates and save them locally
      let lat = Math.round(event.latlng.lat * 100000) / 100000;
      let lng = Math.round(event.latlng.lng * 100000) / 100000;
      cursorcoords = [lat,lng];
      if (drawing) {
        // If the pencil tool is enabled, start drawing
        startDrawing(lat,lng,user);
      }
    // }
  });

  map.addEventListener('click', async (event) => {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      // Get mouse coordinates and save them locally
      let lat = Math.round(event.latlng.lat * 100000) / 100000;
      let lng = Math.round(event.latlng.lng * 100000) / 100000;
      cursorcoords = [lat,lng];
      // Create a marker if the marker tool is enabled
      createMarker(lat,lng,user);
      if (drawing) {
        // If the pencil tool is enabled, start drawing
        startDrawing(lat,lng,user);
      }
    // }
  });

  map.addEventListener('mouseup', (event) => {
    mousedown = false;
  })

  map.addEventListener('mousemove', async (event) => {
    // var user = await checkAuth();
    // console.log(user)
    // if (await checkAuth() != false) {
      // Get cursor coordinates and save them locally
      let lat = Math.round(event.latlng.lat * 100000) / 100000;
      let lng = Math.round(event.latlng.lng * 100000) / 100000;
      cursorcoords = [lat,lng];

      // Make tooltip for line and area hints follow the cursor
      followcursor.setLatLng([lat,lng]);
      if (mousedown && drawing) {
        // If the pencil tool is enabled, draw to the mouse coordinates
        objects.filter(function(result){
          return result.id === currentid && result.user === user.id;
        })[0].line.addLatLng([lat,lng]);

        // Update drawn path in the database
        const { error:error1 } = await _supabase
          .from('coords')
          .insert({
            objects_id: currentid,
            // id: coordsid
            lat: lat,
            lng: lng
          });
      }

      // If drawing a line, show the distance of drawn line in the tooltip
      if (lineon) {
        followcursor.setTooltipContent(((linedistance+linelastcoord.distanceTo([lat,lng]))/1000).toFixed(2)+"km | Double click to finish");
      }
      if (typeof lat != undefined && typeof lng != undefined) {
        if (!dragging) {
          // Save mouse coordinates in the database for real-time cursors, plus current view for observation mode
          // db.ref('rooms/'+room+'/users/'+user.uid).update({
          //     lat:lat,
          //     lng:lng,
          //     view: [map.getBounds().getCenter().lat, map.getBounds().getCenter().lng]
          // });

          // const { data, error } = await _supabase
          // .from('users')
          // .update({ lat:lat,lng:lng,view: [map.getBounds().getCenter().lat, map.getBounds().getCenter().lng] })
          // // .eq('room', room)
          // .eq('uid', user.id)
          // .select()

        } else {
          // Save current view for observation mode
          // db.ref('rooms/'+room+'/users/'+user.uid).update({
          //     view: [map.getBounds().getCenter().lat, map.getBounds().getCenter().lng]
          // });

          // const { data, error } = await _supabase
          // .from('users')
          // .update({ view: [map.getBounds().getCenter().lat, map.getBounds().getCenter().lng] })
          // // .eq('room', room)
          // .eq('uid', user.id)
          // .select()
        }
      }
    // }
  });

  map.addEventListener('zoom', async (event) => {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      // Save current view and zoom for observation mode
      // db.ref('rooms/'+room+'/users/'+user.uid).update({
      //     view: [map.getBounds().getCenter().lat, map.getBounds().getCenter().lng],
      //     zoom: map.getZoom()
      // });
      const { data, error } = await _supabase
          .from('users')
          .update({ view: [map.getBounds().getCenter().lat, map.getBounds().getCenter().lng], zoom: map.getZoom() })
          // .eq('room', room)
          .eq('uid', user.id)
          .select()
      stopObserving();
    // }
  });
  map.addEventListener('movestart', (event) => {
    dragging = true;
  });
  map.addEventListener('moveend', (event) => {
    dragging = false;
  });

  // Server code
  async function checkAuth() {
      // var user = firebase.auth().currentUser;
      const { data: { user:user2 } } = await _supabase.auth.getUser()
      // console.log('data', user)
      user = user2
      // user.uid = 1//Math.floor((Math.random()*10)+1);
      if (user == null) {
          return false;
      } else {
          return user;
      }
  }

  // Sign in use supabase
  async function signIn() {
    var email = sanitize($("#email").val())
    var password = sanitize($("#password").val())
    const { data, error } = await _supabase.auth.signInWithPassword({
      email: email,
      password: password,
    })
    if (error==null){
      closeSignPopup()
    }
    // Sign in using Google
    const { data: { user:user2 } } = await _supabase.auth.getUser()
    user = user2
    // console.log(user)
    // Check if user is inside a file
    if (params.has('file')) {
      $(".signin").removeClass("signin")
      // var user = result.user;
      var usercolor = colors[Math.floor(Math.random()*colors.length)];
      // 先查询有没有，有则update，无则新建
      // 用查询
      const { data:data2, error:error2 } = await _supabase
        .from('users')
        .select()
        .eq('uid',user.id)
        // console.log(data2)

      if (data2.length == 0) {
        const { data:data3,error:error3} = await _supabase
          .from('users')
          .insert({
            room:room,
            uid:user.id,
            lat:0,
            lng:0,
            active: true,
            color: usercolor,
            // session: firebase.database.ServerValue.increment(1),
            name: user.email,
            // imgsrc: profile.photoURL,
            view: [map.getBounds().getCenter().lat, map.getBounds().getCenter().lng],
            zoom: map.getZoom()
          }).select('id')
          // console.log(data3)
      }else{
        const { error:error3 } = await _supabase
          .from('users')
          .update({
            lat:0,
            lng:0,
            active: true,
            color: usercolor,
            // session: firebase.database.ServerValue.increment(1),
            name: user.email,
            // imgsrc: profile.photoURL,
            view: [map.getBounds().getCenter().lat, map.getBounds().getCenter().lng],
            zoom: map.getZoom()
          })
          // .eq('room', room)
          .eq('uid', user.id)
      }
      // Sometimes the session doesn't set properly, might need some time for the last call to go through?
      window.setTimeout(async function(){
        // 用查询
        const { data:data3, error:error3 } = await _supabase
          .from('users')
          .select()
          .eq('uid',user.id)
        session = data3.session
      }, 100);

      // Get data from database
      checkData();
    } else {
      // Prompt the user with a popup to create a map新建room房间
      $("#popup").find(".header-text").html("Create a map");
      $("#popup").find(".subheader-text").html("Maps can be shared with friends to collaborate in real-time.");
      $("#google-signin").attr("id", "create-map");
      $("#create-map").html("Create a map");
    }
  }

  // Log out
  async function logOut() {
    // firebase.auth().signOut().then(function() {
    //   $("#popup").addClass("signin");
    //   $("#overlay").addClass("signin");
    // });
    const { error } = await _supabase.auth.signOut().then(function() {
      $("#popup").addClass("signin");
      $("#overlay").addClass("signin");
    });
  }

  // Create a map
  async function createMap() {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      // var key = db.ref('rooms').push().key;// 新建room id
      const { data,error } = await _supabase
      .from('rooms')
      .insert({
        name: "New map",
        description: "Map description"
      }).select('id')
      // console.log('data', data[0].id)
      var key = data[0].id
      // db.ref("rooms/"+key+"/details").set({// room的detail里的name和description
      //   name: "New map",
      //   description: "Map description"
      // });
      window.location.replace(window.location.href+"?file="+key);
    // }
  }

  // Collapse/expand objects in the sidebar
  function toggleLayer(e) {
    e.preventDefault();
    e.stopPropagation();
    if ($(this).hasClass("arrow-open")) {
      $(this).removeClass("arrow-open");
      $(this).parent().parent().find(".annotation-details").addClass("annotation-closed");
    } else {
      $(this).addClass("arrow-open");
      $(this).parent().parent().find(".annotation-details").removeClass("annotation-closed");
    }
  }

  // Highlight an object in the sidebar
  function focusLayer() {
    showAnnotations();
    if (!$(this).find(".annotation-name span").hasClass("annotation-focus")) {
      const id = $(this).attr("data-id");// 这里取出的是带双引号的字符型id
      // const inst = objects.find(x => x.id === id);// 而objects的id是数值型
      const inst = $.grep(objects, function(e){ return e.id == id; })[0];
      // De-select any previously selected objects
      $(".annotation-focus").removeClass("annotation-focus");

      // Close any opened tooltips
      map.eachLayer(function(layer){
        if (layer.options.pane != "markerPane") {
          layer.closeTooltip();
        }
      });

      // Make layer name bold to show that it has been selected
      $(this).find(".annotation-name span").addClass("annotation-focus");

      // Pan to the annotation and trigger the associated popup
      if (inst.type == "line" || inst.type == "area") {
        map.panTo(inst.trigger.getLatLng());
        $(inst.trigger.getTooltip()._container).removeClass('tooltip-off');
        inst.trigger.openTooltip();
      } else if (inst.type == "marker") {
        map.panTo(inst.marker.getLatLng());
        $(inst.marker.getTooltip()._container).removeClass('tooltip-off');
        inst.marker.openTooltip();
      }
    }
  }

  // Render object in the sidebar
  function renderObjectLayer(object) {
    // Check that the object isn't already rendered in the list
    if ($(".annotation-item[data-id='"+object.id+"']").length == 0) {
      // Render the object in the list depending on the type (different data for each)
      if (object.type == "line") {
        const icon = '<svg class="annotation-icon" width="23" height="23" viewBox="0 0 23 23" fill="none" xmlns="http://www.w3.org/2000/svg"><rect width="23" height="23" rx="5" fill="'+object.color+'"/><path d="M14.5 8.5L8.5 14.5" stroke="white" stroke-width="1.5" stroke-linecap="square"/><path d="M15.8108 8.53378C16.7176 8.53378 17.4527 7.79868 17.4527 6.89189C17.4527 5.9851 16.7176 5.25 15.8108 5.25C14.904 5.25 14.1689 5.9851 14.1689 6.89189C14.1689 7.79868 14.904 8.53378 15.8108 8.53378Z" stroke="white" stroke-width="1.5"/><circle cx="6.89189" cy="15.8108" r="1.64189" stroke="white" stroke-width="1.5"/></svg>'
        $("#annotations-list").prepend('<div class="annotation-item" data-id="'+object.id+'"><div class="annotation-name"><img class="annotation-arrow" src="assets/arrow.svg">'+icon+'<span>'+object.name+'</span><img class="delete-layer" src="assets/delete.svg"></div><div class="annotation-details annotation-closed"><div class="annotation-description">'+object.desc+'</div><div class="annotation-data"><div class="annotation-data-field"><img src="assets/distance-icon.svg">'+object.distance+' km</div></div></div></div>');
      } else if (object.type == "area") {
        const icon = '<svg class="annotation-icon" width="23" height="23" viewBox="0 0 23 23" fill="none" xmlns="http://www.w3.org/2000/svg"><rect width="23" height="23" rx="5" fill="'+object.color+'"/><path d="M15.3652 8.5V13.5" stroke="white" stroke-width="1.5" stroke-linecap="round"/><path d="M8.5 15.3649H13.5" stroke="white" stroke-width="1.5" stroke-linecap="round"/><path d="M14.5303 9.03033C14.8232 8.73744 14.8232 8.26256 14.5303 7.96967C14.2374 7.67678 13.7626 7.67678 13.4697 7.96967L14.5303 9.03033ZM7.96967 13.4697C7.67678 13.7626 7.67678 14.2374 7.96967 14.5303C8.26256 14.8232 8.73744 14.8232 9.03033 14.5303L7.96967 13.4697ZM13.4697 7.96967L7.96967 13.4697L9.03033 14.5303L14.5303 9.03033L13.4697 7.96967Z" fill="white"/><circle cx="15.365" cy="6.85135" r="1.60135" stroke="white" stroke-width="1.5"/><circle cx="15.365" cy="15.3649" r="1.60135" stroke="white" stroke-width="1.5"/><circle cx="6.85135" cy="15.3649" r="1.60135" stroke="white" stroke-width="1.5"/></svg>';
        $("#annotations-list").prepend('<div class="annotation-item" data-id="'+object.id+'"><div class="annotation-name"><img class="annotation-arrow" src="assets/arrow.svg">'+icon+'<span>'+object.name+'</span><img class="delete-layer" src="assets/delete.svg"></div><div class="annotation-details annotation-closed"><div class="annotation-description">'+object.desc+'</div><div class="annotation-data"><div class="annotation-data-field"><img src="assets/area-icon.svg">'+object.area+' km&sup2;</div><div class="annotation-data-field"><img src="assets/perimeter-icon.svg">'+object.distance+' km</div></div></div></div>');
      } else if (object.type == "marker") {
        const icon = '<svg class="annotation-icon" width="23" height="23" viewBox="0 0 23 23" fill="none" xmlns="http://www.w3.org/2000/svg"><rect width="23" height="23" rx="5" fill="'+object.color+'"/><path d="M16.0252 11.2709C16.0252 14.8438 11.3002 17.9063 11.3002 17.9063C11.3002 17.9063 6.5752 14.8438 6.5752 11.2709C6.5752 10.0525 7.07301 8.8841 7.95912 8.0226C8.84522 7.16111 10.047 6.67712 11.3002 6.67712C12.5533 6.67712 13.7552 7.16111 14.6413 8.0226C15.5274 8.8841 16.0252 10.0525 16.0252 11.2709Z" stroke="white" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/><path d="M11.2996 12.8021C12.1695 12.8021 12.8746 12.1166 12.8746 11.2709C12.8746 10.4252 12.1695 9.73962 11.2996 9.73962C10.4298 9.73962 9.72461 10.4252 9.72461 11.2709C9.72461 12.1166 10.4298 12.8021 11.2996 12.8021Z" stroke="white" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>';
        $("#annotations-list").prepend('<div class="annotation-item" data-id="'+object.id+'"><div class="annotation-name"><img class="annotation-arrow" src="assets/arrow.svg">'+icon+'<span>'+object.name+'</span><img class="delete-layer" src="assets/delete.svg"></div><div class="annotation-details annotation-closed"><div class="annotation-description">'+object.desc+'</div><div class="annotation-data"><div class="annotation-data-field"><img src="assets/marker-small-icon.svg">'+object.lat.toFixed(5)+', '+object.lng.toFixed(5)+'</div></div></div></div>');
      }
    } else {
      // If the object already exists, update existing data
      const layer = $(".annotation-item[data-id='"+object.id+"']");
      if (object.type == "line") {
        layer.find(".annotation-name span").html(object.name);
        layer.find(".annotation-description").html(object.desc);
        layer.find(".annotation-data").html('<div class="annotation-data-field"><img src="assets/distance-icon.svg">'+object.distance+' km</div>');
      } else if (object.type == "area") {
        layer.find(".annotation-name span").html(object.name);
        layer.find(".annotation-description").html(object.desc);
        layer.find(".annotation-data").html('<div class="annotation-data-field"><img src="assets/area-icon.svg">'+object.area+' km&sup2;</div><div class="annotation-data-field"><img src="assets/perimeter-icon.svg">'+object.distance+' km</div>');
      } else if (object.type == "marker") {
        layer.find(".annotation-name span").html(object.name);
        layer.find(".annotation-description").html(object.desc);
        layer.find(".annotation-data").html('<div class="annotation-data-field"><img src="assets/marker-small-icon.svg">'+object.lat.toFixed(5)+', '+object.lng.toFixed(5)+'</div>');
      }
    }
  }

  // Delete an object from the sidebar
  async function deleteLayer(e) {
    e.preventDefault();
    e.stopPropagation();
    const id = $(this).parent().parent().attr("data-id");
    // const inst = objects.find(x => x.id === id);
    const inst = $.grep(objects, function(e){ return e.id == id; })[0];
    $(".annotation-item[data-id='"+id+"']").remove();
    if (inst.type != "marker") {
      inst.trigger.remove();
      inst.line.remove();
      // db.ref("rooms/"+room+"/objects/"+inst.id).remove();
      const { error } = await _supabase
        .from('coords')
        .delete()
        .eq('objects_id', inst.id)
      const { error:error2 } = await _supabase
        .from('objects')
        .delete()
        .eq('id', inst.id)

      objects = $.grep(objects, function(e){
           return e.id != inst.id;
      });
    } else {
      inst.marker.remove();
      // db.ref("rooms/"+room+"/objects/"+inst.id).remove();
      // const { error } = await _supabase
      //   .from('coords')
      //   .delete()
      //   .eq('objects_id', inst.id)
      const { error:error2 } = await _supabase
        .from('objects')
        .delete()
        .eq('id', inst.id)
      objects = $.grep(objects, function(e){
           return e.id != inst.id;
      });
    }
  }

  // Editing the name of the map
  async function editMapName(e) {
    if (e.which != 3) {
      return;
    }
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      if (!editingname) {
        oldname = mapname;
        editingname = true;
        $("#map-name").prop("disabled", false);
        $("#map-name").addClass("map-editing");
      }
    // }
  }
  function focusMapName() {
    $("#map-name").select();
    $("#map-name").addClass("map-editing");
    $("#map-name").prop("disabled", false);
  }
  async function stopEditingMapName() {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      editingname = false;
      $("#map-name").prop("disabled", true);
      $("#map-name").removeClass("map-editing");
      var name = sanitize($("#map-name").val());
      if (name.length == 0) {
        // Revert to the old name if its length is 0
        $("#map-name").val(oldname);
      } else {
        // Otherwise, update the name in the database
        // db.ref("rooms/"+room+"/details").update({
        //   name: name
        // })
        const { error } = await _supabase
              .from('rooms')
              .update({
                name: name,
              })
              .eq('id', room)
      }
    // }
  }

  // Editing the description of the map
  async function editMapDescription() {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      if (!editingdescription) {
        olddescription = mapdescription;
        editingdescription = true;
        $("#map-description").prop("disabled", false);
        $("#map-description").addClass("map-editing");
      }
    // }
  }
  function focusMapDescription() {
    $("#map-description").select();
    $("#map-description").addClass("map-editing");
  }
  async function stopEditingMapDescription() {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      editingdescription = false;
      $("#map-description").prop("disabled", true);
      $("#map-description").removeClass("map-editing");
      var name = sanitize($("#map-description").val());
      if (name.length == 0) {
        // Revert to the old description if its length is 0
        $("#map-description").val(olddescription);
      } else {
        // Otherwise, update the description in the database
        // db.ref("rooms/"+room+"/details").update({
        //   description: name
        // })
        const { error } = await _supabase
              .from('rooms')
              .update({
                description: name
              })
              .eq('id', room)
      }
    // }
  }

  // Toggle annotation visibility
  function toggleAnnotations() {
    if (!$("#hide-annotations").hasClass("hidden-annotations")) {
      $(".leaflet-overlay-pane").css({"visibility": "hidden", "pointer-events":"none"});
      $(".leaflet-tooltip-pane").css({"visibility": "hidden", "pointer-events":"none"});
      $("#hide-annotations").addClass("hidden-annotations");
      $("#hide-annotations").html("Show all");
    } else {
      showAnnotations();
    }
  }
  function showAnnotations() {
    $(".leaflet-overlay-pane").css({"visibility": "visible", "pointer-events":"all"});
    $(".leaflet-tooltip-pane").css({"visibility": "visible", "pointer-events":"all"});
    $("#hide-annotations").removeClass("hidden-annotations");
    $("#hide-annotations").html("Hide all");
  }

  // Toggle dots menu
  function toggleMoreMenu() {
    if ($("#more-menu").hasClass("menu-show")) {
      $("#more-menu").removeClass("menu-show");
    } else {
      $("#more-menu").addClass("menu-show");
    }
  }

  // Show share popup
  function showSharePopup() {
    $("#share").addClass("share-show");
    $("#overlay").addClass("share-show");
  }

  // Close share popup
  function closeSharePopup() {
    if ($("#overlay").hasClass("share-show")) {
      $(".share-show").removeClass("share-show");
    }
  }

  // Copy share link
  function copyShareLink() {
    $("#share-url").focus();
    $("#share-url").select();
    document.execCommand('copy');
  }

  // Zoom in
  function zoomIn() {
    map.zoomIn();
  }

  // Zoom out
  function zoomOut() {
    map.zoomOut();
  }

  // Global click handler
  function handleGlobalClicks(e) {
    if ($("#more-menu").hasClass("menu-show") && $(e.target).attr("id") != "more-vertical" && $(e.target).parent().attr("id") != "more-vertical") {
      $("#more-menu").removeClass("menu-show");
    }
  }

  // Export GeoJSON
  function exportGeoJSON() {
    var tempgroup = new L.FeatureGroup();
    map.addLayer(tempgroup);
    map.eachLayer(function(layer) {
      if (layer instanceof L.Marker || layer instanceof L.Polyline || layer instanceof L.Polygon) {
        layer.addTo(tempgroup);
      }
    });

    // Download GeoJSON locally
    var a = document.createElement("a");
    var file = new Blob([JSON.stringify(tempgroup.toGeoJSON())], {type: "application/json"});
    a.href = URL.createObjectURL(file);
    a.download = "geojson";
    a.click();
  }

  // Render user cursors
  async function renderCursors(snapshot, key) {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      if (key != user.id) {
        if (snapshot.active) {
          if (!cursors.find(x => x.id === key)) {
            // Custom cursor icon
            var cursor_icon = L.divIcon({
              html: '<svg width="18" height="18" style="z-index:9999!important" viewBox="0 0 18 18" fill="none" style="background:none;" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M5.51169 15.8783L1.08855 3.64956C0.511984 2.05552 2.05554 0.511969 3.64957 1.08853L15.8783 5.51168C17.5843 6.12877 17.6534 8.51606 15.9858 9.23072L11.2573 11.2573L9.23074 15.9858C8.51607 17.6534 6.12878 17.5843 5.51169 15.8783Z" fill="'+snapshot.color+'"/></svg>',
              iconSize:     [22, 22], // size of the icon
              iconAnchor:   [0, 0], // point of the icon which will correspond to marker's location
              shadowAnchor: [4, 62],  // the same for the shadow
              popupAnchor:  [-3, -76], // point from which the popup should open relative to the iconAnchor
              className: "cursoricon"
            });

            // Create a marker for the cursor
            var cursor_instance = L.marker([snapshot.lat, snapshot.lng], {icon: cursor_icon, pane:"markerPane"});

            // The "tooltip" is just the name of the user that's displayed in the cursor
            cursor_instance.bindTooltip(snapshot.name, { permanent: true, offset: [14, 32], className: "cursor-label color"+snapshot.color.replace("#", ""), direction:"right"});
            cursor_instance.addTo(map)
            cursors.push({id:key, cursor:cursor_instance, color:snapshot.color, name:snapshot.name});

            // Show user avatar on the top right. If they don't have a picture, just put the initial
            var avatar = snapshot.imgsrc;
            if (avatar == null) {
              // avatar = snapshot.name.charAt(0).toUpperCase();
              avatar = snapshot.name;
              $("#right-nav").prepend('<div id="profile" style="background:'+snapshot.color+'!important" class="avatars" data-id="'+key+'">'+avatar+'</div>');
            } else {
              $("#right-nav").prepend('<div id="profile" style="background:'+snapshot.color+'!important" class="avatars" data-id="'+key+'"><img src="'+avatar+'"></div>');
            }
          } else {
            cursors.find(x => x.id === key).cursor.setLatLng([snapshot.lat, snapshot.lng]);
            // Observation mode
            if (observing.status == true && key == observing.id) {
              map.setZoom(snapshot.zoom);
              map.panTo(new L.LatLng(snapshot.view[0], snapshot.view[1]));
            }
          }
        } else if (!snapshot.active && cursors.find(x => x.id === key)) {
          // If the user has disconnected, stop observing them
          if (observing.status == true && key == observing.id) {
            stopObserving();
          }

          // Remove the avatar from the top right
          $(".avatars[data-id="+key+"]").remove();

          // Remove the cursor
          cursors.find(x => x.id === key).cursor.remove();
          cursors = $.grep(cursors, function(e){
               return e.id != key;
          });
        }
      }
    // }
  }

  // Add tooltips for shapes
  async function addShapeInfo(snapshot, key) {
    if (snapshot.type != "draw" && snapshot.type != "marker") {
      var user = await checkAuth();
      if (await checkAuth() != false) {
        var inst = objects.filter(function(result){
          return result.id === key && result.user === snapshot.user
        })[0];
        inst.completed = true;

        // Create a marker so a popup can be shown for the shape
        var centermarker = L.marker(inst.line.getBounds().getCenter(), {zIndexOffset:9999, interactive:false, pane:"overlayPane"});

        // If a marker was already set for the object, simply update the previous variable to reflect that
        if (inst.trigger != "") {
          centermarker = inst.trigger;
          centermarker.unbindTooltip();
        }

        // Create popups that show the name, description, and data of the shapes
        if (inst.type == "line") {
          centermarker.bindTooltip('<h1>'+inst.name+'</h1><h2>'+inst.desc+'</h2><div class="shape-data"><h3><img src="assets/distance-icon.svg">'+inst.distance+' km</h3></div><div class="arrow-down"></div>', {permanent: false, direction:"top", interactive:false, closeOnClick:false, autoclose:false, bubblingMouseEvents:true, className:"create-shape-flow tooltip-off", offset: L.point({x: -15, y: 18})});
        } else if (inst.type == "area") {
          centermarker.bindTooltip('<h1>'+inst.name+'</h1><h2>'+inst.desc+'</h2><div class="shape-data"><h3><img src="assets/area-icon.svg">'+inst.area+' km&sup2;</h3></div><div class="shape-data"><h3><img src="assets/perimeter-icon.svg">'+inst.distance+' km</h3></div><div class="arrow-down"></div>', {permanent: false, direction:"top", interactive:false, bubblingMouseEvents:false, className:"create-shape-flow tooltip-off" ,offset: L.point({x: -15, y: 18})});

          // Areas are lines until they are completed, when they become a polygon
          inst.line.remove();
          var polygon = L.polygon(inst.path, {color:inst.color}).addTo(map);
          inst.line = polygon;
        }

        // Hide the marker (since it's only used for positioning the popup)
        centermarker.setOpacity(0);
        centermarker.addTo(map);
        inst.trigger = centermarker;
        centermarker.getTooltip().update();

        // Detect when clicking on the shape————这里是删除矩形！
        inst.line.on("click", async function(){
          if (!erasing) {
            // Set the marker to the cursor coordinates and show the popup
            inst.trigger.setLatLng(cursorcoords);
            inst.trigger.openTooltip();
            $(inst.trigger.getTooltip()._container).removeClass('tooltip-off');
          } else {
            // If erasing, delete the shape
            inst.trigger.remove();
            inst.line.remove();
            // db.ref("rooms/"+room+"/objects/"+inst.id).remove();
            // 先删除外键关联的coords
            const { error } = await _supabase
              .from('coords')
              .delete()
              .eq('objects_id', inst.id)

            const { error:error2 } = await _supabase
              .from('objects')
              .delete()
              .eq('id', inst.id)

            objects = $.grep(objects, function(e){
                 return e.id != inst.id;
            });
            $(".annotation-item[data-id='"+inst.id+"']").remove();
          }
        });

        // Detect when closing the popup
        inst.line.on("tooltipclose", function(){
          // De-select the object from the sidebar list
          $(".annotation-item[data-id="+inst.id+"]").find(".annotation-name span").removeClass("annotation-focus");
        });

        // Render the object in the sidebar list
        renderObjectLayer(inst);
      }
    }
  }

  // Render a new object
  async function renderShape(snapshot, key) {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      if (snapshot.type == "draw" || snapshot.type == "line" || snapshot.type == "area") {
        if (objects.filter(function(result){
          return result.id === key && result.user === snapshot.user
        }).length == 0) {
          // If the shape doesn't exist locally, create it
          var line = L.polyline([[snapshot.initlat, snapshot.initlng]], {color: snapshot.color});
          if (snapshot.completed && (snapshot.type == "line" || snapshot.type == "area")) {
            // If the shape is already finished, give it all its coordinates
            line = L.polyline(snapshot.path, {color:snapshot.color});
          }
          line.addTo(map);

          // Save shape locally
          if (snapshot.type == "area") {
              objects.push({id:key, local:false, user:snapshot.user, color: snapshot.color, line:line, name:snapshot.name, desc:snapshot.desc, distance:snapshot.distance, area:snapshot.area, path:snapshot.path, completed:snapshot.completed, type:snapshot.type, trigger:"", session:snapshot.session});
          } else {
              objects.push({id:key, local:false, user:snapshot.user, color: snapshot.color, line:line, name:snapshot.name, desc:snapshot.desc, distance:snapshot.distance, path:snapshot.path, completed:snapshot.completed, type:snapshot.type, trigger:"", session:snapshot.session});
          }

          // Detect when clicking on the shape (for freedrawing only)
          line.on("click", async function(e){
            if (snapshot.completed && snapshot.type == "draw" && erasing) {
              // If erasing, delete the line
              line.remove();
              // db.ref("rooms/"+room+"/objects/"+key).remove();
              const { error } = await _supabase
              .from('coords')
              .delete()
              .eq('objects_id', inst.id)

              const { error:error2 } = await _supabase
                .from('objects')
                .delete()
                .eq('id', key)
              objects = $.grep(objects, function(e){
                   return e.id != key;
              });
            }
          });

          // Detect when hovering over a line
          line.on("mouseover", function(event){
            if (erasing && snapshot.type == "draw") {
              line.setStyle({opacity: .3});
            }
          });

          // Detect mouseout on a line
          line.on("mouseout", function(event){
            if (snapshot.type == "draw") {
              line.setStyle({opacity: 1});
            }
          });
        } else {
          // If the object already exists (drawing in progress or already completed), update the coordinates of the object
          var coords = [];
          Object.values(snapshot.coords).forEach(function(coord){
            // coords.push([coord.set[0], coord.set[1]]);
            coords.push([coord.lat, coord.lng]);
          });
          objects.filter(function(result){
            return result.id === key && result.user === snapshot.user
          })[0].line.setLatLngs(coords);
        }
      } else if (snapshot.type == "marker") {
        if (objects.filter(function(result){
          return result.id === key && result.user === snapshot.user
        }).length == 0) {
          // If the marker doesn't exist locally, create it
          var marker_icon;
          if (snapshot.m_type == "none") {
            // Set custom marker icon
            marker_icon = L.divIcon({
              html: '<svg width="30" height="30" viewBox="0 0 46 46" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M23 44.0833C23 44.0833 40.25 32.5833 40.25 19.1666C40.25 14.5916 38.4326 10.204 35.1976 6.96903C31.9626 3.73403 27.575 1.91663 23 1.91663C18.425 1.91663 14.0374 3.73403 10.8024 6.96903C7.56741 10.204 5.75 14.5916 5.75 19.1666C5.75 32.5833 23 44.0833 23 44.0833ZM28.75 19.1666C28.75 22.3423 26.1756 24.9166 23 24.9166C19.8244 24.9166 17.25 22.3423 17.25 19.1666C17.25 15.991 19.8244 13.4166 23 13.4166C26.1756 13.4166 28.75 15.991 28.75 19.1666Z" fill="'+snapshot.color+'"/>/svg>',
              iconSize:     [30, 30], // size of the icon
              iconAnchor:   [15, 30], // point of the icon which will correspond to marker's location
              shadowAnchor: [4, 62],  // the same for the shadow
              popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
            });
          } else {
            // If the marker is of a place found using "find nearby", use a different icon
            var marker_icon = L.icon({
              iconUrl: 'assets/'+snapshot.m_type+'-marker.svg',
              iconSize:     [30, 30],
              iconAnchor:   [15, 30],
              shadowAnchor: [4, 62],
              popupAnchor:  [-3, -76]
            });
          }
          var marker = L.marker([snapshot.lat, snapshot.lng], {icon:marker_icon, interactive:true, direction:"top", pane:"overlayPane"});

          // Create the popup that shows data about the marker
          marker.bindTooltip('<h1>'+snapshot.name+'</h1><h2>'+snapshot.desc+'</h2><div class="shape-data"><h3><img src="assets/marker-small-icon.svg">'+snapshot.lat.toFixed(5)+', '+snapshot.lng.toFixed(5)+'</h3></div><div class="arrow-down"></div>', {permanent: false, direction:"top", className:"create-shape-flow tooltip-off", interactive:false, bubblingMouseEvents:false, offset: L.point({x: 0, y: -35})});
          marker.addTo(map);
          marker.openTooltip();

          // Save the marker locally
          objects.push({id:key, user:snapshot.user, marker:marker, color:snapshot.color, name:snapshot.name, desc:snapshot.desc, session:snapshot.session, local:false, lat:snapshot.lat, lng:snapshot.lng, completed:true, type:"marker"});

          // Detect when clicking on the marker
          marker.on("click", async function(e){
            if (!erasing) {
              // Show the popup
              $(marker.getTooltip()._container).removeClass('tooltip-off');
              marker.openTooltip();
            } else {
              // If erasing, delete the marker
              marker.remove();
              // db.ref("rooms/"+room+"/objects/"+key).remove();
              // 删除mark，mark没有coords
              // const { error } = await _supabase
              // .from('coords')
              // .delete()
              // .eq('objects_id', inst.id)

              const { error:error2 } = await _supabase
                .from('objects')
                .delete()
                .eq('id', key)
              objects = $.grep(objects, function(e){
                   return e.id != key;
              });
              $(".annotation-item[data-id='"+key+"']").remove();
            }
          })

          // Detect when closing the popup
          marker.on("tooltipclose", function(){
            // De-select the marker from the sidebar list
            $(".annotation-item[data-id="+key+"]").find(".annotation-name span").removeClass("annotation-focus");
          });
          marker.closeTooltip();

          // Render the marker in the sidebar list
          renderObjectLayer(objects.find(x => x.id == key));
        } else {
          // If the marker already exists locally, just update its info in the sidebar list
          renderObjectLayer(objects.filter(function(result){
            return result.id === key && result.user === snapshot.user
          }));
        }
      }
      if (snapshot.completed) {
        if (objects.filter(function(result){
          return result.id === key && result.user === snapshot.user
        }).length > 0) {
          // If the shape is completed and it exists locally, update its data
          var inst = objects.filter(function(result){
            return result.id === key && result.user === snapshot.user
          })[0];
          inst.name = snapshot.name;
          inst.desc = snapshot.desc;
          if (snapshot.type == "area") {
            inst.area = snapshot.area;
            inst.distance = snapshot.distance;
            inst.path = snapshot.path;
          } else {
            inst.distance = snapshot.distance;
          }
          addShapeInfo(snapshot, key);
        }
      }
    // }
  }

  // Update object coordinates
  async function updateShapeCoords(snapshot,key) {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      if (snapshot.type == "draw" || snapshot.type == "line" || snapshot.type == "area") {
        if (objects.filter(function(result){
          return result.id === key && result.user === snapshot.user
        }).length > 0) {
          var coords = [];
          Object.values(snapshot.coords).forEach(function(coord){
            // coords.push([coord.set[0], coord.set[1]]);
            coords.push([coord.lat, coord.lng]);
          });
          objects.filter(function(result){
            return result.id === key && result.user === snapshot.user
          })[0].line.setLatLngs(coords);
        }
      }
    // }
  }

  // Interact with the database关键！
  async function checkData() {
    // var user = await checkAuth();
    // if (await checkAuth() != false) {
      // 用查询——这里机理不对，源代码是用.once对一个用户只是查第一次，不会查第二次。
      const { data, error } = await _supabase
        .from('rooms')
        .select()
        .eq('id',room)
        // console.log("room",data)
        mapname = data[0].name;
        mapdescription = data[0].description;
        $("#map-name").val(mapname);
        $("#map-description").val(mapdescription);
        $("#share-nav span").html("Share "+mapname);

      // Check current users on startup源代码是用.once
      const { data:data2, error:error2 } = await _supabase // 用户要写入数据库
        .from('users')
        .select()
        .eq('room',room)
        // console.log("data2",data2)
        if (data2.length != 0) {
          // Object.values(data2).forEach(function(object, index){
          //   renderCursors(cursors, Object.keys(data2)[index]);
          // });调试看看
          data2.forEach(function(object, index){
            renderCursors(cursors, object.id);
          });
        }

      // Check current objects on startup源代码是用.once
      // 用查询
      const { data:data3, error:error3 } = await _supabase
        .from('objects')
        .select(`
            *,
            coords (
              lat,lng
            )
          `)
        .eq('room',room)
        // console.log("object",data2)
        if (data3.length != 0) {
          Object.values(data3).forEach(function(object, index){
            // renderShape(object, Object.keys(data3)[index]);
            renderShape(object, data3[index].id);
          });
        }

      // Update name and description when a change is detected
      _supabase
      .channel('public:rooms')
      .on('postgres_changes', { event: '*', schema: 'public', table: 'rooms' }, payload => {
        console.log('rooms Change received!', payload)
        if (payload.new.id == room) {
          mapname = payload.new.name;
          mapdescription = payload.new.description;
          $("#map-name").val(mapname);
          $("#map-description").val(mapdescription);
        }
      })
      .subscribe()

      // Detect when a user joins the room
      _supabase
      .channel('public:users')
      .on('postgres_changes', { event: 'INSERT', schema: 'public', table: 'users' }, payload => {
        console.log('insert users Change received!', payload)
        if (payload.new.room == room) {
          renderCursors(payload.new, payload.new.id);
        }
      })
      .subscribe()

      // Detect when a user moves their cursor or interacts with the map
      _supabase
      .channel('public:users')
      .on('postgres_changes', { event: 'UPDATE', schema: 'public', table: 'users' }, payload => {
        console.log('update users Change received!', payload)
        if (payload.new.room == room) {
          renderCursors(payload.new, payload.new.id);
        }
      })
      .subscribe()

      // Detect when new objects are added or modified
      _supabase
      .channel('public:objects')
      .on('postgres_changes', { event: '*', schema: 'public', table: 'objects' }, async payload => {
        console.log('objects Change received!', payload)
        if (payload.new != null) {
          // 这里要将coords查询出来哦！！！其他在rendershape前都要关联查出。
          // 这个查询会使得图元消失，奇怪。
          const { data:data4, error:error4 } = await _supabase
            .from('objects')
            .select(`
                *,
                coords (
                  lat,lng
                )
              `)
            .eq('room',room)
          // Check for deleted objects
          objects.forEach(async function(inst){
            if (inst.completed) {
              // if ($.inArray(inst.id, payload.new.id——这个不是数组，所以永远返回-1！！！) == -1) {
              // $.inArray()函数用于在数组中搜索指定的值,并返回其索引值。如果数组中不存在该值,则返回-1;
              // $.inArray(value,array)    --value是要查找的值，array是被查找的数组。
              // console.log(Object.keys(data4))
                // 这里加个循环将id取出来！！！
              var objectsid = [];  
              data4.forEach(function(object, index){
                objectsid.push(object.id);
              })  
              if ($.inArray(inst.id, objectsid) == -1) {
                if (inst.type == "draw") {
                  inst.line.remove();
                } else if (inst.type == "marker") {
                  inst.marker.remove();
                  $(".annotation-item[data-id='"+inst.id+"']").remove();
                } else {
                  inst.trigger.remove();// 注释掉这里了就显示了。
                  if (!inst.local) {
                    inst.line.remove();
                  } else {
                    inst.layer.remove();
                  }
                  $(".annotation-item[data-id='"+inst.id+"']").remove();
                }
                objects = $.grep(objects, function(e){
                  return e.id != inst.id;
                });
              }
            }
          });
          
            // console.log("data4:",data4)
          // Check for new or modified objects，，，将session处理好哦。
          if (Object.keys(data4).length != 0){
            data4.forEach(function(object, index){
              // if (object.user != user.id || object.session != session) {
              if (object.user != user.id) {
                renderShape(object, data4[index].id);
                updateShapeCoords(object, data4[index].id);
              }
            })
          }
        }
      })
      .subscribe()

      // Update user status when disconnected
      // db.ref("rooms/"+room+"/users/"+user.uid).onDisconnect().update({
      //   active: false
      // })
    // }
  }

  // Show sign popup
  function showSignPopup() {
    $("#sign").addClass("sign-show");
    $("#overlay").addClass("sign-show");
  }

  // Close sign popup
  function closeSignPopup() {
    if ($("#overlay").hasClass("sign-show")) {
      $(".sign-show").removeClass("sign-show");
    }
  }


  function popFileSelector() {
    $(`<input type="file" value="选择文件"></input>`)
      .click()
      .on('change', event => {
        let file = event.target.files[0];
        let file_reader = new FileReader();
        file_reader.onload = () => {
            let fc = file_reader.result;
            console.log(fc); // 打印文件文本内容
            // var obj = fc.parseJSON(); //由JSON字符串转换为JSON对象
            var obj = JSON.parse(fc); //由JSON字符串转换为JSON对象
            loadData(obj);
        };
      file_reader.readAsText(file, 'UTF-8');
    });
  }

  // Keyboard shortcuts & more
  $(document).keyup(function(e) {
    if ($(e.target).is("input") || $(e.target).is("textarea")) {
      return;
    }
    if (e.key === "Escape") {
      normalMode();
    } else if (e.key === "Enter") {
      if (editingname) {
        stopEditingMapName();
      } else if (editingdescription) {
        stopEditingMapDescription();
      }
    } else if (e.which == 86) {
      cursorTool();
    } else if (e.which == 80) {
      penTool();
    } else if (e.which == 69) {
      eraserTool();
    } else if (e.which == 77) {
      markerTool();
    } else if (e.which == 76) {
      pathTool();
    } else if (e.which == 65) {
      areaTool();
    }
  });

  // Event handlers
  $(document).on("click", handleGlobalClicks);
  $(document).on("click", "#pen-tool", penTool);
  $(document).on("click", "#cursor-tool", cursorTool);
  $(document).on("click", "#eraser-tool", eraserTool);
  $(document).on("click", "#marker-tool", markerTool);
  $(document).on("click", "#area-tool", areaTool);
  $(document).on("click", "#path-tool", pathTool);
  $(document).on("click", ".color", switchColor);
  $(document).on("click", "#inner-color", toggleColor);
  $(document).on("mouseover", ".tool", showTooltip);
  $(document).on("mouseout", ".tool", hideTooltip);
  $(document).on("click", ".save-button", saveForm);
  $(document).on("click", ".cancel-button", cancelForm);
  $(document).on("click", ".avatars", observationMode);
  $(document).on("click", ".annotation-arrow", toggleLayer);
  $(document).on("click", ".annotation-item", focusLayer);
  $(document).on("click", ".delete-layer", deleteLayer);
  $(document).on("mousedown", "#map-name", editMapName);
  $(document).on("mouseup", "#map-name", focusMapName);
  $(document).on("focusout", "#map-name", stopEditingMapName);
  $(document).on("mousedown", "#map-description", editMapDescription);
  $(document).on("mouseup", "#map-description", focusMapDescription);
  $(document).on("focusout", "#map-description", stopEditingMapDescription);
  $(document).on("click", "#hide-annotations", toggleAnnotations);
  $(document).on("click", "#location-control", targetLiveLocation);
  $(document).on("click", ".find-nearby", findNearby);
  $(document).on("click", ".save-button-place", saveNearby);
  $(document).on("click", ".cancel-button-place", cancelNearby);
  $(document).on("click", "#more-vertical", toggleMoreMenu);
  $(document).on("click", "#geojson", exportGeoJSON);
  $(document).on("click", "#search-box img", search);
  $(document).on("click", "#google-signin", signIn);
  $(document).on("click", "#create-map", createMap);
  $(document).on("click", "#logout", logOut);
  $(document).on("click", "#share-button", showSharePopup);
  $(document).on("click", "#overlay", closeSharePopup);
  $(document).on("click", "#close-share", closeSharePopup);
  $(document).on("click", "#share-copy", copyShareLink);
  $(document).on("click", "#zoom-in", zoomIn);
  $(document).on("click", "#zoom-out", zoomOut);

  $(document).on("click", "#sign-button", showSignPopup);
  // $(document).on("click", "#overlay", closeSignPopup);
  $(document).on("click", "#close-sign", closeSignPopup);
  // $(document).on("click", "#sign-copy", copySignLink);
  $(document).on("click", "#importgeojson", popFileSelector);
  // $(document).on("click", "#login", logIn);

  // Search automatically when focused & pressing enter
  $(document).on("keydown", "#search-input", function(e){
    if (e.key === "Enter") {
      search();
    }
  });

  // Iniialize the map. Could also be done after signing in (but it's less pretty)
  initMap();

  // Get live location of the current user. Only if Geolocation is activated (local only)
  liveLocation();

  // 编辑多边形
  function editMapLayer(layer) {
      layer.pm.enable({
        allowSelfIntersection: true,
        preventMarkerRemoval: false,  // 禁止右键删除点
      })
      let that = this
      // 监听编辑事件
      layer.on('pm:edit', e => {
        // 拖动后的坐标
        that.currentGridData.rangePoints = e.target._latlngs[0]
      })
      layer.on('pm:vertexadded', e =>{
        // 添加顶点
        console.log(e, '添加顶点')
      })
  };


  // Load data
  async function loadData(geojsonFeature){
    // console.log(geojsonFeature)
    // $.ajax("data/geojson.geojson", {
      // dataType: "json",
      // success: function(response){
      // geojsonFeature = response;
    var geojsonMarkerOptions = {
      radius: 8,
      fillColor: "#ff7800",
      color: "#000",
      weight: 1,
      opacity: 1,
      fillOpacity: 0.8
    };
    // L.geoJSON(response.features,{
    L.geoJSON(geojsonFeature.features,{
      pointToLayer: function (feature, latlng){
          return L.circleMarker(latlng, geojsonMarkerOptions);
      },
      // onEachFeature: onEachFeature
    }).addTo(map);

    geojsonFeature.features.forEach(async function(geometry, index){
      if (geometry.geometry.type=="LineString") {
        var temppath = [];
        linedistance = 0;
        var coordlength=geometry.geometry.coordinates.length
        geometry.geometry.coordinates.forEach(function(coordinate, index){
          if (index != 0) {
            // 第一个点
            let latlng = L.latLng(geometry.geometry.coordinates[index-1][1], geometry.geometry.coordinates[index-1][0]);
            // 第二个点
            let latlng2 = L.latLng(coordinate[1], coordinate[0]);
            // 计算结果为：29215.753202443742，注意单位为米
            linedistance += latlng.distanceTo(latlng2);
          }
              
          // Object.values(e.layer.getLatLngs()[0]).forEach(function(a){
          // temppath.push([Object.values(a)[0], Object.values(a)[1]])
          temppath.push([coordinate[1], coordinate[0]]);
          // })
        });            
        // 线路存入数据库
        // If this is the first vertex, get a key and add the new shape in the database
        const { data,error } = await _supabase
          .from('objects')
          .insert({
            room: room,
            // id: currentid,
            color: color,
            // 起点坐标
            initlat: geometry.geometry.coordinates[0][1],
            initlng: geometry.geometry.coordinates[0][0],
            user: user.id,
            type: "line",
            session: session,
            name: "Line",
            desc: "",
            distance: linedistance,
            completed: true,
            path: temppath
          }).select('id');
        // console.log('data', data)
        // console.log('data', data[0].id)
        currentid=data[0].id

        const { error:error1 } = await _supabase
          .from('coords')
          .insert({
            objects_id: data[0].id,
            // id: coordsid 终点坐标
            lat: geometry.geometry.coordinates[coordlength-1][1],
            lng: geometry.geometry.coordinates[coordlength-1][0]
          });
        objects.push({id:currentid, local:true, color:color, user:user.id, name:"Line", desc:"", trigger:"", distance:linedistance, layer:"", type:"line", session:session, completed:true});

        // followcursor.setTooltipContent((linedistance/1000)+"km | Double click to finish");
        // Create a marker so it can trigger a popup when clicking on a line, or area
        // var centermarker = L.marker(map.getBounds().getCenter(), {zIndexOffset:9999, interactive:false, pane:"overlayPane"});
        // Create a popup so users can name and give a description to the shape
        // centermarker.bindTooltip('<label for="shape-name">Name</label><input value="Line" id="shape-name" name="shape-name" /><label for="shape-desc">Description</label><textarea id="shape-desc" name="description"></textarea><br><div id="buttons"><button class="cancel-button">Cancel</button><button class="save-button">Save</button></div><div class="arrow-down"></div>', {permanent: true, direction:"top", interactive:true, bubblingMouseEvents:false, className:"create-shape-flow create-form", offset: L.point({x: -15, y: 18})});
        enteringdata = true;
        var inst = objects.filter(function(result){return result.id === currentid && result.user === user.id;})[0];
        // Calculate total distance / perimeter
        // Create a marker so it can trigger a popup when clicking on a line, or area
        var centermarker = L.marker(map.getBounds().getCenter(), {zIndexOffset:9999, interactive:false, pane:"overlayPane"});
        // Create a popup so users can name and give a description to the shape
        centermarker.bindTooltip('<label for="shape-name">Name</label><input value="'+inst.name+'" id="shape-name" name="shape-name" /><label for="shape-desc">Description</label><textarea id="shape-desc" name="description"></textarea><br><div id="buttons"><button class="cancel-button">Cancel</button><button class="save-button">Save</button></div><div class="arrow-down"></div>', {permanent: true, direction:"top", interactive:true, bubblingMouseEvents:false, className:"create-shape-flow create-form", offset: L.point({x: -15, y: 18})});
        // The marker is supposed to be hidden, it's just for placing the tooltip on the map and triggering it
        centermarker.setOpacity(0);
        centermarker.addTo(map);
        centermarker.openTooltip();
        // Automatically select the name so it's faster to edit
        $("#shape-name").focus();
        $("#shape-name").select();
        // console.log(objects)
        inst.trigger = centermarker;// 为什么这步后，objects里的当前对象的trigger也自动有了值？？
        // console.log(objects)
        // Detect when clicking on the shape

        if (!erasing) {
          // Set the popup to the mouse coordinates and open it
          centermarker.setLatLng(cursorcoords);
          centermarker.openTooltip();
        }

        // Detect when closing the popup (e.g. when clicking outside of it)
        centermarker.on('tooltipclose', function(e){
          if (enteringdata) {
            // If closing the popup before a name and description has been set, revert to defaults
            cancelForm();
          }
          // De-select the object from the sidebar list
          $(".annotation-item[data-id="+currentid+"]").find(".annotation-name span").removeClass("annotation-focus");
        });
      }
    });
    // L.geoJSON(response, {
    //     pointToLayer: function (feature, latlng){
    //         return L.circleMarker(latlng, geojsonMarkerOptions);
    //     },
    //     onEachFeature: onEachFeature
    // }).addTo(mymap);
    // }
    // });
  }
  // Define function for each feature
  function onEachFeature(feature, layer) {
    var city = feature.properties.name;
    var pop_1985 = feature.properties.Pop_1985;
    var pop_1990 = feature.properties.Pop_1990;
    var pop_1995 = feature.properties.Pop_1995;
    var pop_2000 = feature.properties.Pop_2000;
    var pop_2005 = feature.properties.Pop_2005;
    var pop_2010 = feature.properties.Pop_2010;
    var pop_2015 = feature.properties.Pop_2015;

    // Set the context in the popup
    var popupContent = "<h2>"+city+"</h2>"+"<table><tr><th>Year</th><th>Population</th></tr>"
    +"<tr><th>Pop 1985</th><th>"+pop_1985+"</th></tr>"
    +"<tr><th>Pop 1990</th><th>"+pop_1990+"</th></tr>"
    +"<tr><th>Pop 1995</th><th>"+pop_1995+"</th></tr>"
    +"<tr><th>Pop 2000</th><th>"+pop_2000+"</th></tr>"
    +"<tr><th>Pop 2005</th><th>"+pop_2005+"</th></tr>"
    +"<tr><th>Pop 2010</th><th>"+pop_2010+"</th></tr>"
    +"<tr><th>Pop 2015</th><th>"+pop_2015+"</th></tr></table>";

    if (feature.properties && feature.properties.popupContent) {
        popupContent += feature.properties.popupContent;
    }

    layer.bindPopup(popupContent); 
  }

// Load geoJSON data
// loadData(geojsonFeature);

// mymap2.on("click", function (e) {
//     //单击事件
//     addBoat(e.latlng.lat, e.latlng.lng)
// })
/*定义标记*/
var boatIcon = L.icon({
    iconUrl: "img/boat.png",
    iconSize: [23, 30],   //icon的大小
    iconAnchor: [15, 15]  //icon指定点与实际坐标点的偏移量，此处为底部中心点
});

var boatMarker = L.marker([], {   //坐标留空，方便后续添加
    icon: boatIcon,
    draggable: false,  //不允许拖动
})

var boatLineGroup = L.layerGroup([])  //轨迹线图层组，方便管理
boatLineGroup.addTo(map) 
var boatLines = L.polyline([], {color: '#3fe522', weight: 2,}) //轨迹线的样式设定

var boatLineList = []   //绘制轨迹用的

function addBoat(lat_value, long_value) {
    boatMarker.setLatLng([lat_value, long_value])
    boatMarker.addTo(map)

    // boatLineList.push([lat_value, long_value]) //添加至轨迹数组，绘制轨迹用

    //根据marker增加轨迹线
    // const latlngs_length = boatLineList.push([lat_value, long_value])  //push返回新的长度
    // boatLineList_cache.push([lat_value, long_value])
    // if (latlngs_length >= 3){
    //     redrawLine(boatLines, boatLineList)
    // } else if (latlngs_length < 3 && latlngs_length > 1) {
    //     boatLines.setLatLngs(boatLineList)
    //     boatLineGroup.addLayer(boatLines)
    // }
}
function moveBoat(lat_value, long_value, yaw) {
    //parse: 经纬度、角度
    boatMarker.setLatLng([lat_value, long_value]) //设定船的新坐标
    boatMarker.setRotationAngle(yaw);   //旋转船标,1~360，需要leaflet.rotatedMarker.js

    const boat_line_line = boatLineList.push([lat_value, long_value])

    boatLineGroup.addLayer(boatLines)    //向图层组添加轨迹线
    if (boat_line_line > 1) { 
        redrawLine(boatLines, boatLineList)
    }
}
function redrawLine(line_obj, line_list) {
    //接收一个线对象、坐标数组，重绘
    line_obj.setLatLngs(line_list) //给轨迹线设定坐标值，参数为数组
    line_obj.redraw()
}


/*双击添加标记，绘制连接线，拖动标记重绘连接线*/
var marker_group = L.layerGroup([]) //图层组，存放marker
marker_group.addTo(map)

var line_group = L.layerGroup([])  //图层组，存放绘线为marker计数，不可放置同一组
line_group.addTo(map)

var polylinelist = []    //拖动期间该值将被修改
var polylinelist_cache = []   //拖动结束后该值将被修改

var marker_lines = L.polyline([], {color: '#25db71', weight: 1, dashArray: [5, 5],})  //定义线的样式


map.on("dblclick", function (e) {
    //双击事件
    addMarker(e.latlng.lat, e.latlng.lng)
})

// function redrawLine(line_obj, line_list) {
//     //接收一个线对象、坐标数组，重绘
//     line_obj.setLatLngs(line_list)
//     line_obj.redraw()
// }

function IniDivIcon(marker_id) {
    //初始化marker，使其具有编号
    return L.divIcon({
        className: 'my-div-icon',
        html: "<p class='text-center' style='margin-bottom:0;font-size:22px;color:#00ff3a;'>" +
        marker_id +
        "</p>" +
        "<img src='img/position.png' height=30 width=22>",
        iconSize: [23, 30],
        iconAnchor: [11, 63],
    });
}

function addMarker(lat, long) {
    //增加新的marker
    const marker_id = Object.keys(marker_group._layers).length + 1
    const marker = L.marker([lat, long], {
        icon: IniDivIcon(marker_id),
        draggable: true,  //允许拖动
    })
    console.log(lat, long,"marker_id",marker_id)
    marker_group.addLayer(marker)
    //绑定弹出框

    //为每个marker绑定事件
    marker.on("drag", function (e) {
        //拖动事件
        IconOnMove([e.latlng.lat, e.latlng.lng], [e.oldLatLng.lat, e.oldLatLng.lng])
    })
     marker.on("dragend", function (e) {
         //拖动停止事件
         polylinelist_cache = Array.from(polylinelist)  //深copy
     })

     // marker.on("contextmenu", function (e) {
     //     //右键事件
     //     alert("删除标记？")
     //     console.log(e)
     // })

     //根据marker增加轨迹线
     const latlngs_length = polylinelist.push([lat, long])  //push返回新的长度
     polylinelist_cache.push([lat, long])
     if (latlngs_length >= 3){
         redrawLine(marker_lines, polylinelist)
     } else if (latlngs_length < 3 && latlngs_length > 1) {
         marker_lines.setLatLngs(polylinelist)
         line_group.addLayer(marker_lines)
     }
}

function IconOnMove(new_latlng, old_latlng) {
    //移动marker，更新轨迹线
    if (polylinelist_cache.length < 2) {
        return
    }
    polylinelist_cache.forEach(function (value, index) {
        // 遍历连接线的数据列表，与 old_latlng对比，根据index赋予新值
        const result = value.toString() === old_latlng.toString()

        if (result) {
            polylinelist[index] = new_latlng
            redrawLine(marker_lines, polylinelist)
        }
    })
}

function clearAllIcon() {
    //清除所有图层组
    marker_group.clearLayers()
    line_group.clearLayers()
}


});


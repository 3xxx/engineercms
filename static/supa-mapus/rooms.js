rooms
    key

	details
		name: "New map",
    	description: "Map description"
    	active: false

    objects
    	color: inst.color,
        place_id: inst.place_id,
        lat: inst.lat,
        lng: inst.lng,


        user: user.uid,

        type: "marker",
        m_type: inst.m_type,
        session: session,
        name: inst.name,
        desc: ""
        
        area:inst.area,
        distance: inst.distance,
        completed: true
    
        initlat: e.layer._latlngs[0].lat,
        initlng: e.layer._latlngs[0].lng,
        type: "area",
        path: ""
    
        type: "line",
    
        coords ——没有这个？e.shape == "Line"
        set:[linelastcoord.lat,linelastcoord.lng]
        var coords = [];
          Object.values(snapshot.coords).forEach(function(coord){
            coords.push([coord.set[0], coord.set[1]]);
          });
    
        m_type: "none",

    users
    	lat:lat,
        lng:lng,
        view: [map.getBounds().getCenter().lat, map.getBounds().getCenter().lng]
        zoom: map.getZoom()
        active: true,
        color: usercolor,
        session: firebase.database.ServerValue.increment(1),
        name: user.displayName,
        imgsrc: profile.photoURL,

用户登录新建room
箭头函数
还差一个实时怎实现的checkdata 里的once和on
snapshot怎么代替？

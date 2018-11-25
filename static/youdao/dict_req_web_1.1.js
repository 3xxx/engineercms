(function() {
	function a(b, c) {
		var a = "function" == typeof encodeURIComponent ? encodeURIComponent(c) : escape(c);
		d = d.concat("&" + b + "=" + a)
	}
	var d = "";
	window.ADSupporter = {};
	window.ADSupporter.getAdText = function(b, c, e, f, g, h, i) {
		youdao_doctype = b;
		youdao_syndid = c;
		youdao_posid = e;
		youdao_member = f;
		youdao_template = g;
		youdao_width = h;
		youdao_height = i;
		a("req", document.location);
		a("rnd", Math.floor(1E3 * Math.random()));
		"undefined" != typeof youdao_doctype && a("doctype", youdao_doctype);
		"undefined" != typeof youdao_syndid && a("syndid", youdao_syndid);
		"undefined" != typeof youdao_posid && a("posid", youdao_posid);
		"undefined" != typeof youdao_member && a("memberid", youdao_member);
		"undefined" != typeof youdao_template && a("tn", youdao_template);
		b = '<div style=display:none>-</div><iframe align="center,center"';
		"undefined" != typeof youdao_width && (a("width", youdao_width), b += " width=" + youdao_width);
		"undefined" != typeof youdao_height && (a("height", youdao_height), b += " height=" + youdao_height);
                "undefined" != typeof abtest && a("abtest", abtest);
		d = d.substring(0, 2E3);
		d = d.replace(/%\w?$/, "");
		c = document.referrer;
		if (!c) try {
			c = window.opener.location.href
		} catch(j) {
			c = null
		}
		c && (d += "&ref2=" + c);
		b += ' src="http://impservice.dictapp.youdao.com/imp/request.s?' + d.slice(1) + '" marginwidth=0 marginheight=0 scrolling=no frameborder=0 allowtransparency></iframe>';
		d = "";
		return b
	}
})();

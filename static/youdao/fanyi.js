! function(e, t) {
	function o(e) {
		return "[object Function]" === Object.prototype.toString.call(e)
	}

	function r(e) {
		if (!d[e]) throw new Error("Module " + e + " is not defined.");
		var t = d[e];
		return t.module.status !== a.INITIALIZED && n(e), t.module.exports
	}

	function n(e) {
		var n = d[e],
			i = n.module,
			f = i.exports,
			s = i.factory;
		if (i.parent = u, u = n, o(s)) {
			var c = s(r, f, i);
			c !== t && (i.exports = c)
		} else i.exports = d[e].factory;
		i.status = a.INITIALIZED
	}

	function i(e, r, i) {
		if (d[e]) throw new Error("Module " + e + " has been defined already.");
		if ("undefined" == typeof i && (i = r), !o(i) && i !== Object(i)) throw new Error("factory of module " + e + " must be an object or a function.");
		d[e] = {
			module: {
				id: e,
				exports: {},
				uri: "",
				dependencies: [],
				parent: t,
				factory: i,
				status: a.DEFINED
			}
		}, e === f && n(e)
	}
	var a = {
		DEFINED: "The module is just DEFINED",
		INITIALIZED: "The module is compiled and module.exports is available."
	};
	if (!e.define) {
		for (var u, d = {}, f = null, s = document.getElementsByTagName("script"), c = 0, l = s.length; l > c && !f; c++) f = s[c].getAttribute("data-main");
		if (!f) throw new Error("No data-main attribute in script tag.");
		e.define = i
	}
}(window);
define("/lib/jquery-1.6.2", function(e, t) {
	return function(e, t) {
		function n(e, n, r) {
			if (r === t && 1 === e.nodeType) {
				var i = "data-" + n.replace(q, "$1-$2").toLowerCase();
				if (r = e.getAttribute(i), "string" == typeof r) {
					try {
						r = "true" === r ? !0 : "false" === r ? !1 : "null" === r ? null : H.isNaN(r) ? I.test(r) ? H.parseJSON(r) : r : parseFloat(r)
					} catch (o) {}
					H.data(e, n, r)
				} else r = t
			}
			return r
		}

		function r(e) {
			for (var t in e)
				if ("toJSON" !== t) return !1;
			return !0
		}

		function i(e, n, r) {
			var i = n + "defer",
				o = n + "queue",
				a = n + "mark",
				s = H.data(e, i, t, !0);
			!s || "queue" !== r && H.data(e, o, t, !0) || "mark" !== r && H.data(e, a, t, !0) || setTimeout(function() {
				H.data(e, o, t, !0) || H.data(e, a, t, !0) || (H.removeData(e, i, !0), s.resolve())
			}, 0)
		}

		function o() {
			return !1
		}

		function a() {
			return !0
		}

		function s(e, n, r) {
			var i = H.extend({}, r[0]);
			i.type = e, i.originalEvent = {}, i.liveFired = t, H.event.handle.call(n, i), i.isDefaultPrevented() && r[0].preventDefault()
		}

		function l(e) {
			var t, n, r, i, o, a, s, l, u, c, f, d, p = [],
				h = [],
				m = H._data(this, "events");
			if (e.liveFired !== this && m && m.live && !e.target.disabled && (!e.button || "click" !== e.type)) {
				e.namespace && (f = new RegExp("(^|\\.)" + e.namespace.split(".").join("\\.(?:.*\\.)?") + "(\\.|$)")), e.liveFired = this;
				var g = m.live.slice(0);
				for (s = 0; s < g.length; s++) o = g[s], o.origType.replace(G, "") === e.type ? h.push(o.selector) : g.splice(s--, 1);
				for (i = H(e.target).closest(h, e.currentTarget), l = 0, u = i.length; u > l; l++)
					for (c = i[l], s = 0; s < g.length; s++) o = g[s], c.selector !== o.selector || f && !f.test(o.namespace) || c.elem.disabled || (a = c.elem, r = null, ("mouseenter" === o.preType || "mouseleave" === o.preType) && (e.type = o.preType, r = H(e.relatedTarget).closest(o.selector)[0], r && H.contains(a, r) && (r = a)), r && r === a || p.push({
						elem: a,
						handleObj: o,
						level: c.level
					}));
				for (l = 0, u = p.length; u > l && (i = p[l], !(n && i.level > n)) && (e.currentTarget = i.elem, e.data = i.handleObj.data, e.handleObj = i.handleObj, d = i.handleObj.origHandler.apply(i.elem, arguments), d !== !1 && !e.isPropagationStopped() || (n = i.level, d === !1 && (t = !1), !e.isImmediatePropagationStopped())); l++);
				return t
			}
		}

		function u(e, t) {
			return (e && "*" !== e ? e + "." : "") + t.replace(Q, "`").replace(Z, "&")
		}

		function c(e) {
			return !e || !e.parentNode || 11 === e.parentNode.nodeType
		}

		function f(e, t, n) {
			if (t = t || 0, H.isFunction(t)) return H.grep(e, function(e, r) {
				var i = !!t.call(e, r, e);
				return i === n
			});
			if (t.nodeType) return H.grep(e, function(e, r) {
				return e === t === n
			});
			if ("string" == typeof t) {
				var r = H.grep(e, function(e) {
					return 1 === e.nodeType
				});
				if (fe.test(t)) return H.filter(t, r, !n);
				t = H.filter(t, r)
			}
			return H.grep(e, function(e, r) {
				return H.inArray(e, t) >= 0 === n
			})
		}

		function d(e, t) {
			return H.nodeName(e, "table") ? e.getElementsByTagName("tbody")[0] || e.appendChild(e.ownerDocument.createElement("tbody")) : e
		}

		function p(e, t) {
			if (1 === t.nodeType && H.hasData(e)) {
				var n = H.expando,
					r = H.data(e),
					i = H.data(t, r);
				if (r = r[n]) {
					var o = r.events;
					if (i = i[n] = H.extend({}, r), o) {
						delete i.handle, i.events = {};
						for (var a in o)
							for (var s = 0, l = o[a].length; l > s; s++) H.event.add(t, a + (o[a][s].namespace ? "." : "") + o[a][s].namespace, o[a][s], o[a][s].data)
					}
				}
			}
		}

		function h(e, t) {
			var n;
			1 === t.nodeType && (t.clearAttributes && t.clearAttributes(), t.mergeAttributes && t.mergeAttributes(e), n = t.nodeName.toLowerCase(), "object" === n ? t.outerHTML = e.outerHTML : "input" !== n || "checkbox" !== e.type && "radio" !== e.type ? "option" === n ? t.selected = e.defaultSelected : ("input" === n || "textarea" === n) && (t.defaultValue = e.defaultValue) : (e.checked && (t.defaultChecked = t.checked = e.checked), t.value !== e.value && (t.value = e.value)), t.removeAttribute(H.expando))
		}

		function m(e) {
			return "getElementsByTagName" in e ? e.getElementsByTagName("*") : "querySelectorAll" in e ? e.querySelectorAll("*") : []
		}

		function g(e) {
			("checkbox" === e.type || "radio" === e.type) && (e.defaultChecked = e.checked)
		}

		function y(e) {
			H.nodeName(e, "input") ? g(e) : "getElementsByTagName" in e && H.grep(e.getElementsByTagName("input"), g)
		}

		function v(e, t) {
			t.src ? H.ajax({
				url: t.src,
				async: !1,
				dataType: "script"
			}) : H.globalEval((t.text || t.textContent || t.innerHTML || "").replace(Ce, "/*$0*/")), t.parentNode && t.parentNode.removeChild(t)
		}

		function b(e, t, n) {
			var r = "width" === t ? e.offsetWidth : e.offsetHeight,
				i = "width" === t ? Pe : Ie;
			return r > 0 ? ("border" !== n && H.each(i, function() {
				n || (r -= parseFloat(H.css(e, "padding" + this)) || 0), "margin" === n ? r += parseFloat(H.css(e, n + this)) || 0 : r -= parseFloat(H.css(e, "border" + this + "Width")) || 0
			}), r + "px") : (r = Se(e, t, t), (0 > r || null == r) && (r = e.style[t] || 0), r = parseFloat(r) || 0, n && H.each(i, function() {
				r += parseFloat(H.css(e, "padding" + this)) || 0, "padding" !== n && (r += parseFloat(H.css(e, "border" + this + "Width")) || 0), "margin" === n && (r += parseFloat(H.css(e, n + this)) || 0)
			}), r + "px")
		}

		function x(e) {
			return function(t, n) {
				if ("string" != typeof t && (n = t, t = "*"), H.isFunction(n))
					for (var r, i, o, a = t.toLowerCase().split(Ze), s = 0, l = a.length; l > s; s++) r = a[s], o = /^\+/.test(r), o && (r = r.substr(1) || "*"), i = e[r] = e[r] || [], i[o ? "unshift" : "push"](n)
			}
		}

		function T(e, n, r, i, o, a) {
			o = o || n.dataTypes[0], a = a || {}, a[o] = !0;
			for (var s, l = e[o], u = 0, c = l ? l.length : 0, f = e === rt; c > u && (f || !s); u++) s = l[u](n, r, i), "string" == typeof s && (!f || a[s] ? s = t : (n.dataTypes.unshift(s), s = T(e, n, r, i, s, a)));
			return !f && s || a["*"] || (s = T(e, n, r, i, "*", a)), s
		}

		function w(e, t, n, r) {
			if (H.isArray(t)) H.each(t, function(t, i) {
				n || We.test(e) ? r(e, i) : w(e + "[" + ("object" == typeof i || H.isArray(i) ? t : "") + "]", i, n, r)
			});
			else if (n || null == t || "object" != typeof t) r(e, t);
			else
				for (var i in t) w(e + "[" + i + "]", t[i], n, r)
		}

		function N(e, n, r) {
			var i, o, a, s, l = e.contents,
				u = e.dataTypes,
				c = e.responseFields;
			for (o in c) o in r && (n[c[o]] = r[o]);
			for (;
				"*" === u[0];) u.shift(), i === t && (i = e.mimeType || n.getResponseHeader("content-type"));
			if (i)
				for (o in l)
					if (l[o] && l[o].test(i)) {
						u.unshift(o);
						break
					}
			if (u[0] in r) a = u[0];
			else {
				for (o in r) {
					if (!u[0] || e.converters[o + " " + u[0]]) {
						a = o;
						break
					}
					s || (s = o)
				}
				a = a || s
			}
			return a ? (a !== u[0] && u.unshift(a), r[a]) : void 0
		}

		function C(e, n) {
			e.dataFilter && (n = e.dataFilter(n, e.dataType));
			var r, i, o, a, s, l, u, c, f = e.dataTypes,
				d = {},
				p = f.length,
				h = f[0];
			for (r = 1; p > r; r++) {
				if (1 === r)
					for (i in e.converters) "string" == typeof i && (d[i.toLowerCase()] = e.converters[i]);
				if (a = h, h = f[r], "*" === h) h = a;
				else if ("*" !== a && a !== h) {
					if (s = a + " " + h, l = d[s] || d["* " + h], !l) {
						c = t;
						for (u in d)
							if (o = u.split(" "), (o[0] === a || "*" === o[0]) && (c = d[o[1] + " " + h])) {
								u = d[u], u === !0 ? l = c : c === !0 && (l = u);
								break
							}
					}
					l || c || H.error("No conversion from " + s.replace(" ", " to ")), l !== !0 && (n = l ? l(n) : c(u(n)))
				}
			}
			return n
		}

		function E() {
			try {
				return new e.XMLHttpRequest
			} catch (t) {}
		}

		function S() {
			try {
				return new e.ActiveXObject("Microsoft.XMLHTTP")
			} catch (t) {}
		}

		function A() {
			return setTimeout(k, 0), ht = H.now()
		}

		function k() {
			ht = t
		}

		function D(e, t) {
			var n = {};
			return H.each(vt.concat.apply([], vt.slice(0, t)), function() {
				n[this] = e
			}), n
		}

		function F(e) {
			if (!mt[e]) {
				var t = L.body,
					n = H("<" + e + ">").appendTo(t),
					r = n.css("display");
				n.remove(), ("none" === r || "" === r) && (ft || (ft = L.createElement("iframe"), ft.frameBorder = ft.width = ft.height = 0), t.appendChild(ft), dt && ft.createElement || (dt = (ft.contentWindow || ft.contentDocument).document, dt.write(("CSS1Compat" === L.compatMode ? "<!doctype html>" : "") + "<html><body>"), dt.close()), n = dt.createElement(e), dt.body.appendChild(n), r = H.css(n, "display"), t.removeChild(ft)), mt[e] = r
			}
			return mt[e]
		}

		function j(e) {
			return H.isWindow(e) ? e : 9 === e.nodeType ? e.defaultView || e.parentWindow : !1
		}
		var L = e.document,
			M = e.navigator,
			O = e.location,
			H = function() {
				function n() {
					if (!s.isReady) {
						try {
							L.documentElement.doScroll("left")
						} catch (e) {
							return void setTimeout(n, 1)
						}
						s.ready()
					}
				}
				var r, i, o, a, s = function(e, t) {
						return new s.fn.init(e, t, r)
					},
					l = e.jQuery,
					u = e.$,
					c = /^(?:[^<]*(<[\w\W]+>)[^>]*$|#([\w\-]*)$)/,
					f = /\S/,
					d = /^\s+/,
					p = /\s+$/,
					h = /\d/,
					m = /^<(\w+)\s*\/?>(?:<\/\1>)?$/,
					g = /^[\],:{}\s]*$/,
					y = /\\(?:["\\\/bfnrt]|u[0-9a-fA-F]{4})/g,
					v = /"[^"\\\n\r]*"|true|false|null|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?/g,
					b = /(?:^|:|,)(?:\s*\[)+/g,
					x = /(webkit)[ \/]([\w.]+)/,
					T = /(opera)(?:.*version)?[ \/]([\w.]+)/,
					w = /(msie) ([\w.]+)/,
					N = /(mozilla)(?:.*? rv:([\w.]+))?/,
					C = /-([a-z])/gi,
					E = function(e, t) {
						return t.toUpperCase()
					},
					S = M.userAgent,
					A = Object.prototype.toString,
					k = Object.prototype.hasOwnProperty,
					D = Array.prototype.push,
					F = Array.prototype.slice,
					j = String.prototype.trim,
					O = Array.prototype.indexOf,
					H = {};
				return s.fn = s.prototype = {
					constructor: s,
					init: function(e, n, r) {
						var i, o, a, l;
						if (!e) return this;
						if (e.nodeType) return this.context = this[0] = e, this.length = 1, this;
						if ("body" === e && !n && L.body) return this.context = L, this[0] = L.body, this.selector = e, this.length = 1, this;
						if ("string" == typeof e) {
							if (i = "<" === e.charAt(0) && ">" === e.charAt(e.length - 1) && e.length >= 3 ? [null, e, null] : c.exec(e), !i || !i[1] && n) return !n || n.jquery ? (n || r).find(e) : this.constructor(n).find(e);
							if (i[1]) return n = n instanceof s ? n[0] : n, l = n ? n.ownerDocument || n : L, a = m.exec(e), a ? s.isPlainObject(n) ? (e = [L.createElement(a[1])], s.fn.attr.call(e, n, !0)) : e = [l.createElement(a[1])] : (a = s.buildFragment([i[1]], [l]), e = (a.cacheable ? s.clone(a.fragment) : a.fragment).childNodes), s.merge(this, e);
							if (o = L.getElementById(i[2]), o && o.parentNode) {
								if (o.id !== i[2]) return r.find(e);
								this.length = 1, this[0] = o
							}
							return this.context = L, this.selector = e, this
						}
						return s.isFunction(e) ? r.ready(e) : (e.selector !== t && (this.selector = e.selector, this.context = e.context), s.makeArray(e, this))
					},
					selector: "",
					jquery: "1.6.2",
					length: 0,
					size: function() {
						return this.length
					},
					toArray: function() {
						return F.call(this, 0)
					},
					get: function(e) {
						return null == e ? this.toArray() : 0 > e ? this[this.length + e] : this[e]
					},
					pushStack: function(e, t, n) {
						var r = this.constructor();
						return s.isArray(e) ? D.apply(r, e) : s.merge(r, e), r.prevObject = this, r.context = this.context, "find" === t ? r.selector = this.selector + (this.selector ? " " : "") + n : t && (r.selector = this.selector + "." + t + "(" + n + ")"), r
					},
					each: function(e, t) {
						return s.each(this, e, t)
					},
					ready: function(e) {
						return s.bindReady(), o.done(e), this
					},
					eq: function(e) {
						return -1 === e ? this.slice(e) : this.slice(e, +e + 1)
					},
					first: function() {
						return this.eq(0)
					},
					last: function() {
						return this.eq(-1)
					},
					slice: function() {
						return this.pushStack(F.apply(this, arguments), "slice", F.call(arguments).join(","))
					},
					map: function(e) {
						return this.pushStack(s.map(this, function(t, n) {
							return e.call(t, n, t)
						}))
					},
					end: function() {
						return this.prevObject || this.constructor(null)
					},
					push: D,
					sort: [].sort,
					splice: [].splice
				}, s.fn.init.prototype = s.fn, s.extend = s.fn.extend = function() {
					var e, n, r, i, o, a, l = arguments[0] || {},
						u = 1,
						c = arguments.length,
						f = !1;
					for ("boolean" == typeof l && (f = l, l = arguments[1] || {}, u = 2), "object" == typeof l || s.isFunction(l) || (l = {}), c === u && (l = this, --u); c > u; u++)
						if (null != (e = arguments[u]))
							for (n in e) r = l[n], i = e[n], l !== i && (f && i && (s.isPlainObject(i) || (o = s.isArray(i))) ? (o ? (o = !1, a = r && s.isArray(r) ? r : []) : a = r && s.isPlainObject(r) ? r : {}, l[n] = s.extend(f, a, i)) : i !== t && (l[n] = i));
					return l
				}, s.extend({
					noConflict: function(t) {
						return e.$ === s && (e.$ = u), t && e.jQuery === s && (e.jQuery = l), s
					},
					isReady: !1,
					readyWait: 1,
					holdReady: function(e) {
						e ? s.readyWait++ : s.ready(!0)
					},
					ready: function(e) {
						if (e === !0 && !--s.readyWait || e !== !0 && !s.isReady) {
							if (!L.body) return setTimeout(s.ready, 1);
							if (s.isReady = !0, e !== !0 && --s.readyWait > 0) return;
							o.resolveWith(L, [s]), s.fn.trigger && s(L).trigger("ready").unbind("ready")
						}
					},
					bindReady: function() {
						if (!o) {
							if (o = s._Deferred(), "complete" === L.readyState) return setTimeout(s.ready, 1);
							if (L.addEventListener) L.addEventListener("DOMContentLoaded", a, !1), e.addEventListener("load", s.ready, !1);
							else if (L.attachEvent) {
								L.attachEvent("onreadystatechange", a), e.attachEvent("onload", s.ready);
								var t = !1;
								try {
									t = null == e.frameElement
								} catch (r) {}
								L.documentElement.doScroll && t && n()
							}
						}
					},
					isFunction: function(e) {
						return "function" === s.type(e)
					},
					isArray: Array.isArray || function(e) {
						return "array" === s.type(e)
					},
					isWindow: function(e) {
						return e && "object" == typeof e && "setInterval" in e
					},
					isNaN: function(e) {
						return null == e || !h.test(e) || isNaN(e)
					},
					type: function(e) {
						return null == e ? String(e) : H[A.call(e)] || "object"
					},
					isPlainObject: function(e) {
						if (!e || "object" !== s.type(e) || e.nodeType || s.isWindow(e)) return !1;
						if (e.constructor && !k.call(e, "constructor") && !k.call(e.constructor.prototype, "isPrototypeOf")) return !1;
						var n;
						for (n in e);
						return n === t || k.call(e, n)
					},
					isEmptyObject: function(e) {
						for (var t in e) return !1;
						return !0
					},
					error: function(e) {
						throw e
					},
					parseJSON: function(t) {
						return "string" == typeof t && t ? (t = s.trim(t), e.JSON && e.JSON.parse ? e.JSON.parse(t) : g.test(t.replace(y, "@").replace(v, "]").replace(b, "")) ? new Function("return " + t)() : void s.error("Invalid JSON: " + t)) : null
					},
					parseXML: function(t, n, r) {
						return e.DOMParser ? (r = new DOMParser, n = r.parseFromString(t, "text/xml")) : (n = new ActiveXObject("Microsoft.XMLDOM"), n.async = "false", n.loadXML(t)), r = n.documentElement, r && r.nodeName && "parsererror" !== r.nodeName || s.error("Invalid XML: " + t), n
					},
					noop: function() {},
					globalEval: function(t) {
						t && f.test(t) && (e.execScript || function(t) {
							e.eval.call(e, t)
						})(t)
					},
					camelCase: function(e) {
						return e.replace(C, E)
					},
					nodeName: function(e, t) {
						return e.nodeName && e.nodeName.toUpperCase() === t.toUpperCase()
					},
					each: function(e, n, r) {
						var i, o = 0,
							a = e.length,
							l = a === t || s.isFunction(e);
						if (r)
							if (l) {
								for (i in e)
									if (n.apply(e[i], r) === !1) break
							} else
								for (; a > o && n.apply(e[o++], r) !== !1;);
						else if (l) {
							for (i in e)
								if (n.call(e[i], i, e[i]) === !1) break
						} else
							for (; a > o && n.call(e[o], o, e[o++]) !== !1;);
						return e
					},
					trim: j ? function(e) {
						return null == e ? "" : j.call(e)
					} : function(e) {
						return null == e ? "" : e.toString().replace(d, "").replace(p, "")
					},
					makeArray: function(e, t) {
						var n = t || [];
						if (null != e) {
							var r = s.type(e);
							null == e.length || "string" === r || "function" === r || "regexp" === r || s.isWindow(e) ? D.call(n, e) : s.merge(n, e)
						}
						return n
					},
					inArray: function(e, t) {
						if (O) return O.call(t, e);
						for (var n = 0, r = t.length; r > n; n++)
							if (t[n] === e) return n;
						return -1
					},
					merge: function(e, n) {
						var r = e.length,
							i = 0;
						if ("number" == typeof n.length)
							for (var o = n.length; o > i; i++) e[r++] = n[i];
						else
							for (; n[i] !== t;) e[r++] = n[i++];
						return e.length = r, e
					},
					grep: function(e, t, n) {
						var r, i = [];
						n = !!n;
						for (var o = 0, a = e.length; a > o; o++) r = !!t(e[o], o), n !== r && i.push(e[o]);
						return i
					},
					map: function(e, n, r) {
						var i, o, a = [],
							l = 0,
							u = e.length,
							c = e instanceof s || u !== t && "number" == typeof u && (u > 0 && e[0] && e[u - 1] || 0 === u || s.isArray(e));
						if (c)
							for (; u > l; l++) i = n(e[l], l, r), null != i && (a[a.length] = i);
						else
							for (o in e) i = n(e[o], o, r), null != i && (a[a.length] = i);
						return a.concat.apply([], a)
					},
					guid: 1,
					proxy: function(e, n) {
						if ("string" == typeof n) {
							var r = e[n];
							n = e, e = r
						}
						if (!s.isFunction(e)) return t;
						var i = F.call(arguments, 2),
							o = function() {
								return e.apply(n, i.concat(F.call(arguments)))
							};
						return o.guid = e.guid = e.guid || o.guid || s.guid++, o
					},
					access: function(e, n, r, i, o, a) {
						var l = e.length;
						if ("object" == typeof n) {
							for (var u in n) s.access(e, u, n[u], i, o, r);
							return e
						}
						if (r !== t) {
							i = !a && i && s.isFunction(r);
							for (var c = 0; l > c; c++) o(e[c], n, i ? r.call(e[c], c, o(e[c], n)) : r, a);
							return e
						}
						return l ? o(e[0], n) : t
					},
					now: function() {
						return (new Date).getTime()
					},
					uaMatch: function(e) {
						e = e.toLowerCase();
						var t = x.exec(e) || T.exec(e) || w.exec(e) || e.indexOf("compatible") < 0 && N.exec(e) || [];
						return {
							browser: t[1] || "",
							version: t[2] || "0"
						}
					},
					sub: function() {
						function e(t, n) {
							return new e.fn.init(t, n)
						}
						s.extend(!0, e, this), e.superclass = this, e.fn = e.prototype = this(), e.fn.constructor = e, e.sub = this.sub, e.fn.init = function(n, r) {
							return r && r instanceof s && !(r instanceof e) && (r = e(r)), s.fn.init.call(this, n, r, t)
						}, e.fn.init.prototype = e.fn;
						var t = e(L);
						return e
					},
					browser: {}
				}), s.each("Boolean Number String Function Array Date RegExp Object".split(" "), function(e, t) {
					H["[object " + t + "]"] = t.toLowerCase()
				}), i = s.uaMatch(S), i.browser && (s.browser[i.browser] = !0, s.browser.version = i.version), s.browser.webkit && (s.browser.safari = !0), f.test(" ") && (d = /^[\s\xA0]+/, p = /[\s\xA0]+$/), r = s(L), L.addEventListener ? a = function() {
					L.removeEventListener("DOMContentLoaded", a, !1), s.ready()
				} : L.attachEvent && (a = function() {
					"complete" === L.readyState && (L.detachEvent("onreadystatechange", a), s.ready())
				}), s
			}(),
			B = "done fail isResolved isRejected promise then always pipe".split(" "),
			P = [].slice;
		H.extend({
			_Deferred: function() {
				var e, t, n, r = [],
					i = {
						done: function() {
							if (!n) {
								var t, o, a, s, l, u = arguments;
								for (e && (l = e, e = 0), t = 0, o = u.length; o > t; t++) a = u[t], s = H.type(a), "array" === s ? i.done.apply(i, a) : "function" === s && r.push(a);
								l && i.resolveWith(l[0], l[1])
							}
							return this
						},
						resolveWith: function(i, o) {
							if (!n && !e && !t) {
								o = o || [], t = 1;
								try {
									for (; r[0];) r.shift().apply(i, o)
								} finally {
									e = [i, o], t = 0
								}
							}
							return this
						},
						resolve: function() {
							return i.resolveWith(this, arguments), this
						},
						isResolved: function() {
							return !(!t && !e)
						},
						cancel: function() {
							return n = 1, r = [], this
						}
					};
				return i
			},
			Deferred: function(e) {
				var t, n = H._Deferred(),
					r = H._Deferred();
				return H.extend(n, {
					then: function(e, t) {
						return n.done(e).fail(t), this
					},
					always: function() {
						return n.done.apply(n, arguments).fail.apply(this, arguments)
					},
					fail: r.done,
					rejectWith: r.resolveWith,
					reject: r.resolve,
					isRejected: r.isResolved,
					pipe: function(e, t) {
						return H.Deferred(function(r) {
							H.each({
								done: [e, "resolve"],
								fail: [t, "reject"]
							}, function(e, t) {
								var i, o = t[0],
									a = t[1];
								n[e](H.isFunction(o) ? function() {
									i = o.apply(this, arguments), i && H.isFunction(i.promise) ? i.promise().then(r.resolve, r.reject) : r[a](i)
								} : r[a])
							})
						}).promise()
					},
					promise: function(e) {
						if (null == e) {
							if (t) return t;
							t = e = {}
						}
						for (var r = B.length; r--;) e[B[r]] = n[B[r]];
						return e
					}
				}), n.done(r.cancel).fail(n.cancel), delete n.cancel, e && e.call(n, n), n
			},
			when: function(e) {
				function t(e) {
					return function(t) {
						n[e] = arguments.length > 1 ? P.call(arguments, 0) : t, --o || a.resolveWith(a, P.call(n, 0))
					}
				}
				var n = arguments,
					r = 0,
					i = n.length,
					o = i,
					a = 1 >= i && e && H.isFunction(e.promise) ? e : H.Deferred();
				if (i > 1) {
					for (; i > r; r++) n[r] && H.isFunction(n[r].promise) ? n[r].promise().then(t(r), a.reject) : --o;
					o || a.resolveWith(a, n)
				} else a !== e && a.resolveWith(a, i ? [e] : []);
				return a.promise()
			}
		}), H.support = function() {
			var e, t, n, r, i, o, a, s, l, u, c, f, d, p, h, m, g = L.createElement("div"),
				y = L.documentElement;
			if (g.setAttribute("className", "t"), g.innerHTML = "   <link/><table></table><a href='/a' style='top:1px;float:left;opacity:.55;'>a</a><input type='checkbox'/>", e = g.getElementsByTagName("*"), t = g.getElementsByTagName("a")[0], !e || !e.length || !t) return {};
			n = L.createElement("select"), r = n.appendChild(L.createElement("option")), i = g.getElementsByTagName("input")[0], a = {
				leadingWhitespace: 3 === g.firstChild.nodeType,
				tbody: !g.getElementsByTagName("tbody").length,
				htmlSerialize: !!g.getElementsByTagName("link").length,
				style: /top/.test(t.getAttribute("style")),
				hrefNormalized: "/a" === t.getAttribute("href"),
				opacity: /^0.55$/.test(t.style.opacity),
				cssFloat: !!t.style.cssFloat,
				checkOn: "on" === i.value,
				optSelected: r.selected,
				getSetAttribute: "t" !== g.className,
				submitBubbles: !0,
				changeBubbles: !0,
				focusinBubbles: !1,
				deleteExpando: !0,
				noCloneEvent: !0,
				inlineBlockNeedsLayout: !1,
				shrinkWrapBlocks: !1,
				reliableMarginRight: !0
			}, i.checked = !0, a.noCloneChecked = i.cloneNode(!0).checked, n.disabled = !0, a.optDisabled = !r.disabled;
			try {
				delete g.test
			} catch (v) {
				a.deleteExpando = !1
			}!g.addEventListener && g.attachEvent && g.fireEvent && (g.attachEvent("onclick", function() {
				a.noCloneEvent = !1
			}), g.cloneNode(!0).fireEvent("onclick")), i = L.createElement("input"), i.value = "t", i.setAttribute("type", "radio"), a.radioValue = "t" === i.value, i.setAttribute("checked", "checked"), g.appendChild(i), s = L.createDocumentFragment(), s.appendChild(g.firstChild), a.checkClone = s.cloneNode(!0).cloneNode(!0).lastChild.checked, g.innerHTML = "", g.style.width = g.style.paddingLeft = "1px", l = L.getElementsByTagName("body")[0], c = L.createElement(l ? "div" : "body"), f = {
				visibility: "hidden",
				width: 0,
				height: 0,
				border: 0,
				margin: 0
			}, l && H.extend(f, {
				position: "absolute",
				left: -1e3,
				top: -1e3
			});
			for (h in f) c.style[h] = f[h];
			if (c.appendChild(g), u = l || y, u.insertBefore(c, u.firstChild), a.appendChecked = i.checked, a.boxModel = 2 === g.offsetWidth, "zoom" in g.style && (g.style.display = "inline", g.style.zoom = 1, a.inlineBlockNeedsLayout = 2 === g.offsetWidth, g.style.display = "", g.innerHTML = "<div style='width:4px;'></div>", a.shrinkWrapBlocks = 2 !== g.offsetWidth), g.innerHTML = "<table><tr><td style='padding:0;border:0;display:none'></td><td>t</td></tr></table>", d = g.getElementsByTagName("td"), m = 0 === d[0].offsetHeight, d[0].style.display = "", d[1].style.display = "none", a.reliableHiddenOffsets = m && 0 === d[0].offsetHeight, g.innerHTML = "", L.defaultView && L.defaultView.getComputedStyle && (o = L.createElement("div"), o.style.width = "0", o.style.marginRight = "0", g.appendChild(o), a.reliableMarginRight = 0 === (parseInt((L.defaultView.getComputedStyle(o, null) || {
					marginRight: 0
				}).marginRight, 10) || 0)), c.innerHTML = "", u.removeChild(c), g.attachEvent)
				for (h in {
						submit: 1,
						change: 1,
						focusin: 1
					}) p = "on" + h, m = p in g, m || (g.setAttribute(p, "return;"), m = "function" == typeof g[p]), a[h + "Bubbles"] = m;
			return c = s = n = r = l = o = g = i = null, a
		}(), H.boxModel = H.support.boxModel;
		var I = /^(?:\{.*\}|\[.*\])$/,
			q = /([a-z])([A-Z])/g;
		H.extend({
			cache: {},
			uuid: 0,
			expando: "jQuery" + (H.fn.jquery + Math.random()).replace(/\D/g, ""),
			noData: {
				embed: !0,
				object: "clsid:D27CDB6E-AE6D-11cf-96B8-444553540000",
				applet: !0
			},
			hasData: function(e) {
				return e = e.nodeType ? H.cache[e[H.expando]] : e[H.expando], !!e && !r(e)
			},
			data: function(e, n, r, i) {
				if (H.acceptData(e)) {
					var o, a = H.expando,
						s = "string" == typeof n,
						l = e.nodeType,
						u = l ? H.cache : e,
						c = l ? e[H.expando] : e[H.expando] && H.expando;
					if (c && (!i || !c || u[c][a]) || !s || r !== t) return c || (l ? e[H.expando] = c = ++H.uuid : c = H.expando), u[c] || (u[c] = {}, l || (u[c].toJSON = H.noop)), ("object" == typeof n || "function" == typeof n) && (i ? u[c][a] = H.extend(u[c][a], n) : u[c] = H.extend(u[c], n)), o = u[c], i && (o[a] || (o[a] = {}), o = o[a]), r !== t && (o[H.camelCase(n)] = r), "events" !== n || o[n] ? s ? o[H.camelCase(n)] || o[n] : o : o[a] && o[a].events
				}
			},
			removeData: function(t, n, i) {
				if (H.acceptData(t)) {
					var o = H.expando,
						a = t.nodeType,
						s = a ? H.cache : t,
						l = a ? t[H.expando] : H.expando;
					if (s[l]) {
						if (n) {
							var u = i ? s[l][o] : s[l];
							if (u && (delete u[n], !r(u))) return
						}
						if (!i || (delete s[l][o], r(s[l]))) {
							var c = s[l][o];
							H.support.deleteExpando || s != e ? delete s[l] : s[l] = null, c ? (s[l] = {}, a || (s[l].toJSON = H.noop), s[l][o] = c) : a && (H.support.deleteExpando ? delete t[H.expando] : t.removeAttribute ? t.removeAttribute(H.expando) : t[H.expando] = null)
						}
					}
				}
			},
			_data: function(e, t, n) {
				return H.data(e, t, n, !0)
			},
			acceptData: function(e) {
				if (e.nodeName) {
					var t = H.noData[e.nodeName.toLowerCase()];
					if (t) return !(t === !0 || e.getAttribute("classid") !== t)
				}
				return !0
			}
		}), H.fn.extend({
			data: function(e, r) {
				var i = null;
				if ("undefined" == typeof e) {
					if (this.length && (i = H.data(this[0]), 1 === this[0].nodeType))
						for (var o, a = this[0].attributes, s = 0, l = a.length; l > s; s++) o = a[s].name, 0 === o.indexOf("data-") && (o = H.camelCase(o.substring(5)), n(this[0], o, i[o]));
					return i
				}
				if ("object" == typeof e) return this.each(function() {
					H.data(this, e)
				});
				var u = e.split(".");
				return u[1] = u[1] ? "." + u[1] : "", r === t ? (i = this.triggerHandler("getData" + u[1] + "!", [u[0]]), i === t && this.length && (i = H.data(this[0], e), i = n(this[0], e, i)), i === t && u[1] ? this.data(u[0]) : i) : this.each(function() {
					var t = H(this),
						n = [u[0], r];
					t.triggerHandler("setData" + u[1] + "!", n), H.data(this, e, r), t.triggerHandler("changeData" + u[1] + "!", n)
				})
			},
			removeData: function(e) {
				return this.each(function() {
					H.removeData(this, e)
				})
			}
		}), H.extend({
			_mark: function(e, n) {
				e && (n = (n || "fx") + "mark", H.data(e, n, (H.data(e, n, t, !0) || 0) + 1, !0))
			},
			_unmark: function(e, n, r) {
				if (e !== !0 && (r = n, n = e, e = !1), n) {
					r = r || "fx";
					var o = r + "mark",
						a = e ? 0 : (H.data(n, o, t, !0) || 1) - 1;
					a ? H.data(n, o, a, !0) : (H.removeData(n, o, !0), i(n, r, "mark"))
				}
			},
			queue: function(e, n, r) {
				if (e) {
					n = (n || "fx") + "queue";
					var i = H.data(e, n, t, !0);
					return r && (!i || H.isArray(r) ? i = H.data(e, n, H.makeArray(r), !0) : i.push(r)), i || []
				}
			},
			dequeue: function(e, t) {
				t = t || "fx";
				var n = H.queue(e, t),
					r = n.shift();
				"inprogress" === r && (r = n.shift()), r && ("fx" === t && n.unshift("inprogress"), r.call(e, function() {
					H.dequeue(e, t)
				})), n.length || (H.removeData(e, t + "queue", !0), i(e, t, "queue"))
			}
		}), H.fn.extend({
			queue: function(e, n) {
				return "string" != typeof e && (n = e, e = "fx"), n === t ? H.queue(this[0], e) : this.each(function() {
					var t = H.queue(this, e, n);
					"fx" === e && "inprogress" !== t[0] && H.dequeue(this, e)
				})
			},
			dequeue: function(e) {
				return this.each(function() {
					H.dequeue(this, e)
				})
			},
			delay: function(e, t) {
				return e = H.fx ? H.fx.speeds[e] || e : e, t = t || "fx", this.queue(t, function() {
					var n = this;
					setTimeout(function() {
						H.dequeue(n, t)
					}, e)
				})
			},
			clearQueue: function(e) {
				return this.queue(e || "fx", [])
			},
			promise: function(e, n) {
				function r() {
					--l || o.resolveWith(a, [a])
				}
				"string" != typeof e && (n = e, e = t), e = e || "fx";
				for (var i, o = H.Deferred(), a = this, s = a.length, l = 1, u = e + "defer", c = e + "queue", f = e + "mark"; s--;)(i = H.data(a[s], u, t, !0) || (H.data(a[s], c, t, !0) || H.data(a[s], f, t, !0)) && H.data(a[s], u, H._Deferred(), !0)) && (l++, i.done(r));
				return r(), o.promise()
			}
		});
		var _, R, W = /[\n\t\r]/g,
			$ = /\s+/,
			z = /\r/g,
			X = /^(?:button|input)$/i,
			V = /^(?:button|input|object|select|textarea)$/i,
			U = /^a(?:rea)?$/i,
			J = /^(?:autofocus|autoplay|async|checked|controls|defer|disabled|hidden|loop|multiple|open|readonly|required|scoped|selected)$/i,
			Y = /\:|^on/;
		H.fn.extend({
			attr: function(e, t) {
				return H.access(this, e, t, !0, H.attr)
			},
			removeAttr: function(e) {
				return this.each(function() {
					H.removeAttr(this, e)
				})
			},
			prop: function(e, t) {
				return H.access(this, e, t, !0, H.prop)
			},
			removeProp: function(e) {
				return e = H.propFix[e] || e, this.each(function() {
					try {
						this[e] = t, delete this[e]
					} catch (n) {}
				})
			},
			addClass: function(e) {
				var t, n, r, i, o, a, s;
				if (H.isFunction(e)) return this.each(function(t) {
					H(this).addClass(e.call(this, t, this.className))
				});
				if (e && "string" == typeof e)
					for (t = e.split($), n = 0, r = this.length; r > n; n++)
						if (i = this[n], 1 === i.nodeType)
							if (i.className || 1 !== t.length) {
								for (o = " " + i.className + " ", a = 0, s = t.length; s > a; a++) ~o.indexOf(" " + t[a] + " ") || (o += t[a] + " ");
								i.className = H.trim(o)
							} else i.className = e;
				return this
			},
			removeClass: function(e) {
				var n, r, i, o, a, s, l;
				if (H.isFunction(e)) return this.each(function(t) {
					H(this).removeClass(e.call(this, t, this.className))
				});
				if (e && "string" == typeof e || e === t)
					for (n = (e || "").split($), r = 0, i = this.length; i > r; r++)
						if (o = this[r], 1 === o.nodeType && o.className)
							if (e) {
								for (a = (" " + o.className + " ").replace(W, " "), s = 0, l = n.length; l > s; s++) a = a.replace(" " + n[s] + " ", " ");
								o.className = H.trim(a)
							} else o.className = "";
				return this
			},
			toggleClass: function(e, t) {
				var n = typeof e,
					r = "boolean" == typeof t;
				return this.each(H.isFunction(e) ? function(n) {
					H(this).toggleClass(e.call(this, n, this.className, t), t)
				} : function() {
					if ("string" === n)
						for (var i, o = 0, a = H(this), s = t, l = e.split($); i = l[o++];) s = r ? s : !a.hasClass(i), a[s ? "addClass" : "removeClass"](i);
					else("undefined" === n || "boolean" === n) && (this.className && H._data(this, "__className__", this.className), this.className = this.className || e === !1 ? "" : H._data(this, "__className__") || "")
				})
			},
			hasClass: function(e) {
				for (var t = " " + e + " ", n = 0, r = this.length; r > n; n++)
					if ((" " + this[n].className + " ").replace(W, " ").indexOf(t) > -1) return !0;
				return !1
			},
			val: function(e) {
				var n, r, i = this[0];
				if (!arguments.length) return i ? (n = H.valHooks[i.nodeName.toLowerCase()] || H.valHooks[i.type], n && "get" in n && (r = n.get(i, "value")) !== t ? r : (r = i.value, "string" == typeof r ? r.replace(z, "") : null == r ? "" : r)) : t;
				var o = H.isFunction(e);
				return this.each(function(r) {
					var i, a = H(this);
					1 === this.nodeType && (i = o ? e.call(this, r, a.val()) : e, null == i ? i = "" : "number" == typeof i ? i += "" : H.isArray(i) && (i = H.map(i, function(e) {
						return null == e ? "" : e + ""
					})), n = H.valHooks[this.nodeName.toLowerCase()] || H.valHooks[this.type], n && "set" in n && n.set(this, i, "value") !== t || (this.value = i))
				})
			}
		}), H.extend({
			valHooks: {
				option: {
					get: function(e) {
						var t = e.attributes.value;
						return !t || t.specified ? e.value : e.text
					}
				},
				select: {
					get: function(e) {
						var t, n = e.selectedIndex,
							r = [],
							i = e.options,
							o = "select-one" === e.type;
						if (0 > n) return null;
						for (var a = o ? n : 0, s = o ? n + 1 : i.length; s > a; a++) {
							var l = i[a];
							if (!(!l.selected || (H.support.optDisabled ? l.disabled : null !== l.getAttribute("disabled")) || l.parentNode.disabled && H.nodeName(l.parentNode, "optgroup"))) {
								if (t = H(l).val(), o) return t;
								r.push(t)
							}
						}
						return o && !r.length && i.length ? H(i[n]).val() : r
					},
					set: function(e, t) {
						var n = H.makeArray(t);
						return H(e).find("option").each(function() {
							this.selected = H.inArray(H(this).val(), n) >= 0
						}), n.length || (e.selectedIndex = -1), n
					}
				}
			},
			attrFn: {
				val: !0,
				css: !0,
				html: !0,
				text: !0,
				data: !0,
				width: !0,
				height: !0,
				offset: !0
			},
			attrFix: {
				tabindex: "tabIndex"
			},
			attr: function(e, n, r, i) {
				var o = e.nodeType;
				if (!e || 3 === o || 8 === o || 2 === o) return t;
				if (i && n in H.attrFn) return H(e)[n](r);
				if (!("getAttribute" in e)) return H.prop(e, n, r);
				var a, s, l = 1 !== o || !H.isXMLDoc(e);
				return l && (n = H.attrFix[n] || n, s = H.attrHooks[n], s || (J.test(n) ? s = R : _ && "className" !== n && (H.nodeName(e, "form") || Y.test(n)) && (s = _))), r !== t ? null === r ? (H.removeAttr(e, n), t) : s && "set" in s && l && (a = s.set(e, r, n)) !== t ? a : (e.setAttribute(n, "" + r), r) : s && "get" in s && l && null !== (a = s.get(e, n)) ? a : (a = e.getAttribute(n), null === a ? t : a)
			},
			removeAttr: function(e, t) {
				var n;
				1 === e.nodeType && (t = H.attrFix[t] || t, H.support.getSetAttribute ? e.removeAttribute(t) : (H.attr(e, t, ""), e.removeAttributeNode(e.getAttributeNode(t))), J.test(t) && (n = H.propFix[t] || t) in e && (e[n] = !1))
			},
			attrHooks: {
				type: {
					set: function(e, t) {
						if (X.test(e.nodeName) && e.parentNode) H.error("type property can't be changed");
						else if (!H.support.radioValue && "radio" === t && H.nodeName(e, "input")) {
							var n = e.value;
							return e.setAttribute("type", t), n && (e.value = n), t
						}
					}
				},
				tabIndex: {
					get: function(e) {
						var n = e.getAttributeNode("tabIndex");
						return n && n.specified ? parseInt(n.value, 10) : V.test(e.nodeName) || U.test(e.nodeName) && e.href ? 0 : t
					}
				},
				value: {
					get: function(e, t) {
						return _ && H.nodeName(e, "button") ? _.get(e, t) : t in e ? e.value : null
					},
					set: function(e, t, n) {
						return _ && H.nodeName(e, "button") ? _.set(e, t, n) : void(e.value = t)
					}
				}
			},
			propFix: {
				tabindex: "tabIndex",
				readonly: "readOnly",
				"for": "htmlFor",
				"class": "className",
				maxlength: "maxLength",
				cellspacing: "cellSpacing",
				cellpadding: "cellPadding",
				rowspan: "rowSpan",
				colspan: "colSpan",
				usemap: "useMap",
				frameborder: "frameBorder",
				contenteditable: "contentEditable"
			},
			prop: function(e, n, r) {
				var i = e.nodeType;
				if (!e || 3 === i || 8 === i || 2 === i) return t;
				var o, a, s = 1 !== i || !H.isXMLDoc(e);
				return s && (n = H.propFix[n] || n, a = H.propHooks[n]), r !== t ? a && "set" in a && (o = a.set(e, r, n)) !== t ? o : e[n] = r : a && "get" in a && (o = a.get(e, n)) !== t ? o : e[n]
			},
			propHooks: {}
		}), R = {
			get: function(e, n) {
				return H.prop(e, n) ? n.toLowerCase() : t
			},
			set: function(e, t, n) {
				var r;
				return t === !1 ? H.removeAttr(e, n) : (r = H.propFix[n] || n, r in e && (e[r] = !0), e.setAttribute(n, n.toLowerCase())), n
			}
		}, H.support.getSetAttribute || (H.attrFix = H.propFix, _ = H.attrHooks.name = H.attrHooks.title = H.valHooks.button = {
			get: function(e, n) {
				var r;
				return r = e.getAttributeNode(n), r && "" !== r.nodeValue ? r.nodeValue : t
			},
			set: function(e, t, n) {
				var r = e.getAttributeNode(n);
				return r ? (r.nodeValue = t, t) : void 0
			}
		}, H.each(["width", "height"], function(e, t) {
			H.attrHooks[t] = H.extend(H.attrHooks[t], {
				set: function(e, n) {
					return "" === n ? (e.setAttribute(t, "auto"), n) : void 0
				}
			})
		})), H.support.hrefNormalized || H.each(["href", "src", "width", "height"], function(e, n) {
			H.attrHooks[n] = H.extend(H.attrHooks[n], {
				get: function(e) {
					var r = e.getAttribute(n, 2);
					return null === r ? t : r
				}
			})
		}), H.support.style || (H.attrHooks.style = {
			get: function(e) {
				return e.style.cssText.toLowerCase() || t
			},
			set: function(e, t) {
				return e.style.cssText = "" + t
			}
		}), H.support.optSelected || (H.propHooks.selected = H.extend(H.propHooks.selected, {
			get: function(e) {
				var t = e.parentNode;
				t && (t.selectedIndex, t.parentNode && t.parentNode.selectedIndex)
			}
		})), H.support.checkOn || H.each(["radio", "checkbox"], function() {
			H.valHooks[this] = {
				get: function(e) {
					return null === e.getAttribute("value") ? "on" : e.value
				}
			}
		}), H.each(["radio", "checkbox"], function() {
			H.valHooks[this] = H.extend(H.valHooks[this], {
				set: function(e, t) {
					return H.isArray(t) ? e.checked = H.inArray(H(e).val(), t) >= 0 : void 0
				}
			})
		});
		var G = /\.(.*)$/,
			K = /^(?:textarea|input|select)$/i,
			Q = /\./g,
			Z = / /g,
			ee = /[^\w\s.|`]/g,
			te = function(e) {
				return e.replace(ee, "\\$&")
			};
		H.event = {
			add: function(e, n, r, i) {
				if (3 !== e.nodeType && 8 !== e.nodeType) {
					if (r === !1) r = o;
					else if (!r) return;
					var a, s;
					r.handler && (a = r, r = a.handler), r.guid || (r.guid = H.guid++);
					var l = H._data(e);
					if (l) {
						var u = l.events,
							c = l.handle;
						u || (l.events = u = {}), c || (l.handle = c = function(e) {
							return "undefined" == typeof H || e && H.event.triggered === e.type ? t : H.event.handle.apply(c.elem, arguments)
						}), c.elem = e, n = n.split(" ");
						for (var f, d, p = 0; f = n[p++];) {
							s = a ? H.extend({}, a) : {
								handler: r,
								data: i
							}, f.indexOf(".") > -1 ? (d = f.split("."), f = d.shift(), s.namespace = d.slice(0).sort().join(".")) : (d = [], s.namespace = ""), s.type = f, s.guid || (s.guid = r.guid);
							var h = u[f],
								m = H.event.special[f] || {};
							h || (h = u[f] = [], m.setup && m.setup.call(e, i, d, c) !== !1 || (e.addEventListener ? e.addEventListener(f, c, !1) : e.attachEvent && e.attachEvent("on" + f, c))), m.add && (m.add.call(e, s), s.handler.guid || (s.handler.guid = r.guid)), h.push(s), H.event.global[f] = !0
						}
						e = null
					}
				}
			},
			global: {},
			remove: function(e, n, r, i) {
				if (3 !== e.nodeType && 8 !== e.nodeType) {
					r === !1 && (r = o);
					var a, s, l, u, c, f, d, p, h, m, g = 0,
						y = H.hasData(e) && H._data(e),
						v = y && y.events;
					if (y && v)
						if (n && n.type && (r = n.handler, n = n.type), !n || "string" == typeof n && "." === n.charAt(0)) {
							n = n || "";
							for (s in v) H.event.remove(e, s + n)
						} else {
							for (n = n.split(" "); s = n[g++];)
								if (m = s, h = null, u = s.indexOf(".") < 0, c = [], u || (c = s.split("."), s = c.shift(), f = new RegExp("(^|\\.)" + H.map(c.slice(0).sort(), te).join("\\.(?:.*\\.)?") + "(\\.|$)")), p = v[s])
									if (r) {
										for (d = H.event.special[s] || {}, l = i || 0; l < p.length && (h = p[l], r.guid !== h.guid || ((u || f.test(h.namespace)) && (null == i && p.splice(l--, 1), d.remove && d.remove.call(e, h)), null == i)); l++);
										(0 === p.length || null != i && 1 === p.length) && (d.teardown && d.teardown.call(e, c) !== !1 || H.removeEvent(e, s, y.handle), a = null, delete v[s])
									} else
										for (l = 0; l < p.length; l++) h = p[l], (u || f.test(h.namespace)) && (H.event.remove(e, m, h.handler, l), p.splice(l--, 1));
							if (H.isEmptyObject(v)) {
								var b = y.handle;
								b && (b.elem = null), delete y.events, delete y.handle, H.isEmptyObject(y) && H.removeData(e, t, !0)
							}
						}
				}
			},
			customEvent: {
				getData: !0,
				setData: !0,
				changeData: !0
			},
			trigger: function(n, r, i, o) {
				var a, s = n.type || n,
					l = [];
				if (s.indexOf("!") >= 0 && (s = s.slice(0, -1), a = !0), s.indexOf(".") >= 0 && (l = s.split("."), s = l.shift(), l.sort()), i && !H.event.customEvent[s] || H.event.global[s]) {
					if (n = "object" == typeof n ? n[H.expando] ? n : new H.Event(s, n) : new H.Event(s), n.type = s, n.exclusive = a, n.namespace = l.join("."), n.namespace_re = new RegExp("(^|\\.)" + l.join("\\.(?:.*\\.)?") + "(\\.|$)"), (o || !i) && (n.preventDefault(), n.stopPropagation()), !i) return void H.each(H.cache, function() {
						var e = H.expando,
							t = this[e];
						t && t.events && t.events[s] && H.event.trigger(n, r, t.handle.elem)
					});
					if (3 !== i.nodeType && 8 !== i.nodeType) {
						n.result = t, n.target = i, r = null != r ? H.makeArray(r) : [], r.unshift(n);
						var u = i,
							c = s.indexOf(":") < 0 ? "on" + s : "";
						do {
							var f = H._data(u, "handle");
							n.currentTarget = u, f && f.apply(u, r), c && H.acceptData(u) && u[c] && u[c].apply(u, r) === !1 && (n.result = !1, n.preventDefault()), u = u.parentNode || u.ownerDocument || u === n.target.ownerDocument && e
						} while (u && !n.isPropagationStopped());
						if (!n.isDefaultPrevented()) {
							var d, p = H.event.special[s] || {};
							if (!(p._default && p._default.call(i.ownerDocument, n) !== !1 || "click" === s && H.nodeName(i, "a") || !H.acceptData(i))) {
								try {
									c && i[s] && (d = i[c], d && (i[c] = null), H.event.triggered = s, i[s]())
								} catch (h) {}
								d && (i[c] = d), H.event.triggered = t
							}
						}
						return n.result
					}
				}
			},
			handle: function(n) {
				n = H.event.fix(n || e.event);
				var r = ((H._data(this, "events") || {})[n.type] || []).slice(0),
					i = !n.exclusive && !n.namespace,
					o = Array.prototype.slice.call(arguments, 0);
				o[0] = n, n.currentTarget = this;
				for (var a = 0, s = r.length; s > a; a++) {
					var l = r[a];
					if (i || n.namespace_re.test(l.namespace)) {
						n.handler = l.handler, n.data = l.data, n.handleObj = l;
						var u = l.handler.apply(this, o);
						if (u !== t && (n.result = u, u === !1 && (n.preventDefault(), n.stopPropagation())), n.isImmediatePropagationStopped()) break
					}
				}
				return n.result
			},
			props: "altKey attrChange attrName bubbles button cancelable charCode clientX clientY ctrlKey currentTarget data detail eventPhase fromElement handler keyCode layerX layerY metaKey newValue offsetX offsetY pageX pageY prevValue relatedNode relatedTarget screenX screenY shiftKey srcElement target toElement view wheelDelta which".split(" "),
			fix: function(e) {
				if (e[H.expando]) return e;
				var n = e;
				e = H.Event(n);
				for (var r, i = this.props.length; i;) r = this.props[--i], e[r] = n[r];
				if (e.target || (e.target = e.srcElement || L), 3 === e.target.nodeType && (e.target = e.target.parentNode), !e.relatedTarget && e.fromElement && (e.relatedTarget = e.fromElement === e.target ? e.toElement : e.fromElement), null == e.pageX && null != e.clientX) {
					var o = e.target.ownerDocument || L,
						a = o.documentElement,
						s = o.body;
					e.pageX = e.clientX + (a && a.scrollLeft || s && s.scrollLeft || 0) - (a && a.clientLeft || s && s.clientLeft || 0), e.pageY = e.clientY + (a && a.scrollTop || s && s.scrollTop || 0) - (a && a.clientTop || s && s.clientTop || 0)
				}
				return null != e.which || null == e.charCode && null == e.keyCode || (e.which = null != e.charCode ? e.charCode : e.keyCode), !e.metaKey && e.ctrlKey && (e.metaKey = e.ctrlKey), e.which || e.button === t || (e.which = 1 & e.button ? 1 : 2 & e.button ? 3 : 4 & e.button ? 2 : 0), e
			},
			guid: 1e8,
			proxy: H.proxy,
			special: {
				ready: {
					setup: H.bindReady,
					teardown: H.noop
				},
				live: {
					add: function(e) {
						H.event.add(this, u(e.origType, e.selector), H.extend({}, e, {
							handler: l,
							guid: e.handler.guid
						}))
					},
					remove: function(e) {
						H.event.remove(this, u(e.origType, e.selector), e)
					}
				},
				beforeunload: {
					setup: function(e, t, n) {
						H.isWindow(this) && (this.onbeforeunload = n)
					},
					teardown: function(e, t) {
						this.onbeforeunload === t && (this.onbeforeunload = null)
					}
				}
			}
		}, H.removeEvent = L.removeEventListener ? function(e, t, n) {
			e.removeEventListener && e.removeEventListener(t, n, !1)
		} : function(e, t, n) {
			e.detachEvent && e.detachEvent("on" + t, n)
		}, H.Event = function(e, t) {
			return this.preventDefault ? (e && e.type ? (this.originalEvent = e, this.type = e.type, this.isDefaultPrevented = e.defaultPrevented || e.returnValue === !1 || e.getPreventDefault && e.getPreventDefault() ? a : o) : this.type = e, t && H.extend(this, t), this.timeStamp = H.now(), void(this[H.expando] = !0)) : new H.Event(e, t)
		}, H.Event.prototype = {
			preventDefault: function() {
				this.isDefaultPrevented = a;
				var e = this.originalEvent;
				e && (e.preventDefault ? e.preventDefault() : e.returnValue = !1)
			},
			stopPropagation: function() {
				this.isPropagationStopped = a;
				var e = this.originalEvent;
				e && (e.stopPropagation && e.stopPropagation(), e.cancelBubble = !0)
			},
			stopImmediatePropagation: function() {
				this.isImmediatePropagationStopped = a, this.stopPropagation()
			},
			isDefaultPrevented: o,
			isPropagationStopped: o,
			isImmediatePropagationStopped: o
		};
		var ne = function(e) {
				var t = e.relatedTarget,
					n = !1,
					r = e.type;
				e.type = e.data, t !== this && (t && (n = H.contains(this, t)), n || (H.event.handle.apply(this, arguments), e.type = r))
			},
			re = function(e) {
				e.type = e.data, H.event.handle.apply(this, arguments)
			};
		if (H.each({
				mouseenter: "mouseover",
				mouseleave: "mouseout"
			}, function(e, t) {
				H.event.special[e] = {
					setup: function(n) {
						H.event.add(this, t, n && n.selector ? re : ne, e)
					},
					teardown: function(e) {
						H.event.remove(this, t, e && e.selector ? re : ne)
					}
				}
			}), H.support.submitBubbles || (H.event.special.submit = {
				setup: function(e, t) {
					return H.nodeName(this, "form") ? !1 : (H.event.add(this, "click.specialSubmit", function(e) {
						var t = e.target,
							n = t.type;
						"submit" !== n && "image" !== n || !H(t).closest("form").length || s("submit", this, arguments)
					}), void H.event.add(this, "keypress.specialSubmit", function(e) {
						var t = e.target,
							n = t.type;
						"text" !== n && "password" !== n || !H(t).closest("form").length || 13 !== e.keyCode || s("submit", this, arguments)
					}))
				},
				teardown: function(e) {
					H.event.remove(this, ".specialSubmit")
				}
			}), !H.support.changeBubbles) {
			var ie, oe = function(e) {
					var t = e.type,
						n = e.value;
					return "radio" === t || "checkbox" === t ? n = e.checked : "select-multiple" === t ? n = e.selectedIndex > -1 ? H.map(e.options, function(e) {
						return e.selected
					}).join("-") : "" : H.nodeName(e, "select") && (n = e.selectedIndex), n
				},
				ae = function(e) {
					var n, r, i = e.target;
					K.test(i.nodeName) && !i.readOnly && (n = H._data(i, "_change_data"), r = oe(i), ("focusout" !== e.type || "radio" !== i.type) && H._data(i, "_change_data", r), n !== t && r !== n && (null != n || r) && (e.type = "change", e.liveFired = t, H.event.trigger(e, arguments[1], i)))
				};
			H.event.special.change = {
				filters: {
					focusout: ae,
					beforedeactivate: ae,
					click: function(e) {
						var t = e.target,
							n = H.nodeName(t, "input") ? t.type : "";
						("radio" === n || "checkbox" === n || H.nodeName(t, "select")) && ae.call(this, e)
					},
					keydown: function(e) {
						var t = e.target,
							n = H.nodeName(t, "input") ? t.type : "";
						(13 === e.keyCode && !H.nodeName(t, "textarea") || 32 === e.keyCode && ("checkbox" === n || "radio" === n) || "select-multiple" === n) && ae.call(this, e)
					},
					beforeactivate: function(e) {
						var t = e.target;
						H._data(t, "_change_data", oe(t))
					}
				},
				setup: function(e, t) {
					if ("file" === this.type) return !1;
					for (var n in ie) H.event.add(this, n + ".specialChange", ie[n]);
					return K.test(this.nodeName)
				},
				teardown: function(e) {
					return H.event.remove(this, ".specialChange"), K.test(this.nodeName)
				}
			}, ie = H.event.special.change.filters, ie.focus = ie.beforeactivate
		}
		H.support.focusinBubbles || H.each({
			focus: "focusin",
			blur: "focusout"
		}, function(e, t) {
			function n(e) {
				var n = H.event.fix(e);
				n.type = t, n.originalEvent = {}, H.event.trigger(n, null, n.target), n.isDefaultPrevented() && e.preventDefault()
			}
			var r = 0;
			H.event.special[t] = {
				setup: function() {
					0 === r++ && L.addEventListener(e, n, !0)
				},
				teardown: function() {
					0 === --r && L.removeEventListener(e, n, !0)
				}
			}
		}), H.each(["bind", "one"], function(e, n) {
			H.fn[n] = function(e, r, i) {
				var o;
				if ("object" == typeof e) {
					for (var a in e) this[n](a, r, e[a], i);
					return this
				}
				if ((2 === arguments.length || r === !1) && (i = r, r = t), "one" === n ? (o = function(e) {
						return H(this).unbind(e, o), i.apply(this, arguments)
					}, o.guid = i.guid || H.guid++) : o = i, "unload" === e && "one" !== n) this.one(e, r, i);
				else
					for (var s = 0, l = this.length; l > s; s++) H.event.add(this[s], e, o, r);
				return this
			}
		}), H.fn.extend({
			unbind: function(e, t) {
				if ("object" != typeof e || e.preventDefault)
					for (var n = 0, r = this.length; r > n; n++) H.event.remove(this[n], e, t);
				else
					for (var i in e) this.unbind(i, e[i]);
				return this
			},
			delegate: function(e, t, n, r) {
				return this.live(t, n, r, e)
			},
			undelegate: function(e, t, n) {
				return 0 === arguments.length ? this.unbind("live") : this.die(t, null, n, e)
			},
			trigger: function(e, t) {
				return this.each(function() {
					H.event.trigger(e, t, this)
				})
			},
			triggerHandler: function(e, t) {
				return this[0] ? H.event.trigger(e, t, this[0], !0) : void 0
			},
			toggle: function(e) {
				var t = arguments,
					n = e.guid || H.guid++,
					r = 0,
					i = function(n) {
						var i = (H.data(this, "lastToggle" + e.guid) || 0) % r;
						return H.data(this, "lastToggle" + e.guid, i + 1), n.preventDefault(), t[i].apply(this, arguments) || !1
					};
				for (i.guid = n; r < t.length;) t[r++].guid = n;
				return this.click(i)
			},
			hover: function(e, t) {
				return this.mouseenter(e).mouseleave(t || e)
			}
		});
		var se = {
			focus: "focusin",
			blur: "focusout",
			mouseenter: "mouseover",
			mouseleave: "mouseout"
		};
		H.each(["live", "die"], function(e, n) {
				H.fn[n] = function(e, r, i, a) {
					var s, l, c, f, d = 0,
						p = a || this.selector,
						h = a ? this : H(this.context);
					if ("object" == typeof e && !e.preventDefault) {
						for (var m in e) h[n](m, r, e[m], p);
						return this
					}
					if ("die" === n && !e && a && "." === a.charAt(0)) return h.unbind(a), this;
					for ((r === !1 || H.isFunction(r)) && (i = r || o, r = t), e = (e || "").split(" "); null != (s = e[d++]);)
						if (l = G.exec(s), c = "", l && (c = l[0], s = s.replace(G, "")), "hover" !== s)
							if (f = s, se[s] ? (e.push(se[s] + c), s += c) : s = (se[s] || s) + c, "live" === n)
								for (var g = 0, y = h.length; y > g; g++) H.event.add(h[g], "live." + u(s, p), {
									data: r,
									selector: p,
									handler: i,
									origType: s,
									origHandler: i,
									preType: f
								});
							else h.unbind("live." + u(s, p), i);
					else e.push("mouseenter" + c, "mouseleave" + c);
					return this
				}
			}), H.each("blur focus focusin focusout load resize scroll unload click dblclick mousedown mouseup mousemove mouseover mouseout mouseenter mouseleave change select submit keydown keypress keyup error".split(" "), function(e, t) {
				H.fn[t] = function(e, n) {
					return null == n && (n = e, e = null), arguments.length > 0 ? this.bind(t, e, n) : this.trigger(t)
				}, H.attrFn && (H.attrFn[t] = !0)
			}),
			function() {
				function e(e, t, n, r, i, o) {
					for (var a = 0, s = r.length; s > a; a++) {
						var l = r[a];
						if (l) {
							var u = !1;
							for (l = l[e]; l;) {
								if (l.sizcache === n) {
									u = r[l.sizset];
									break
								}
								if (1 !== l.nodeType || o || (l.sizcache = n, l.sizset = a), l.nodeName.toLowerCase() === t) {
									u = l;
									break
								}
								l = l[e]
							}
							r[a] = u
						}
					}
				}

				function n(e, t, n, r, i, o) {
					for (var a = 0, s = r.length; s > a; a++) {
						var l = r[a];
						if (l) {
							var u = !1;
							for (l = l[e]; l;) {
								if (l.sizcache === n) {
									u = r[l.sizset];
									break
								}
								if (1 === l.nodeType)
									if (o || (l.sizcache = n, l.sizset = a), "string" != typeof t) {
										if (l === t) {
											u = !0;
											break
										}
									} else if (c.filter(t, [l]).length > 0) {
									u = l;
									break
								}
								l = l[e]
							}
							r[a] = u
						}
					}
				}
				var r = /((?:\((?:\([^()]+\)|[^()]+)+\)|\[(?:\[[^\[\]]*\]|['"][^'"]*['"]|[^\[\]'"]+)+\]|\\.|[^ >+~,(\[\\]+)+|[>+~])(\s*,\s*)?((?:.|\r|\n)*)/g,
					i = 0,
					o = Object.prototype.toString,
					a = !1,
					s = !0,
					l = /\\/g,
					u = /\W/;
				[0, 0].sort(function() {
					return s = !1, 0
				});
				var c = function(e, t, n, i) {
					n = n || [], t = t || L;
					var a = t;
					if (1 !== t.nodeType && 9 !== t.nodeType) return [];
					if (!e || "string" != typeof e) return n;
					var s, l, u, p, h, g, y, v, x = !0,
						T = c.isXML(t),
						w = [],
						N = e;
					do
						if (r.exec(""), s = r.exec(N), s && (N = s[3], w.push(s[1]), s[2])) {
							p = s[3];
							break
						}
					while (s);
					if (w.length > 1 && d.exec(e))
						if (2 === w.length && f.relative[w[0]]) l = b(w[0] + w[1], t);
						else
							for (l = f.relative[w[0]] ? [t] : c(w.shift(), t); w.length;) e = w.shift(), f.relative[e] && (e += w.shift()), l = b(e, l);
					else if (!i && w.length > 1 && 9 === t.nodeType && !T && f.match.ID.test(w[0]) && !f.match.ID.test(w[w.length - 1]) && (h = c.find(w.shift(), t, T), t = h.expr ? c.filter(h.expr, h.set)[0] : h.set[0]), t)
						for (h = i ? {
								expr: w.pop(),
								set: m(i)
							} : c.find(w.pop(), 1 !== w.length || "~" !== w[0] && "+" !== w[0] || !t.parentNode ? t : t.parentNode, T), l = h.expr ? c.filter(h.expr, h.set) : h.set, w.length > 0 ? u = m(l) : x = !1; w.length;) g = w.pop(), y = g, f.relative[g] ? y = w.pop() : g = "", null == y && (y = t), f.relative[g](u, y, T);
					else u = w = [];
					if (u || (u = l), u || c.error(g || e), "[object Array]" === o.call(u))
						if (x)
							if (t && 1 === t.nodeType)
								for (v = 0; null != u[v]; v++) u[v] && (u[v] === !0 || 1 === u[v].nodeType && c.contains(t, u[v])) && n.push(l[v]);
							else
								for (v = 0; null != u[v]; v++) u[v] && 1 === u[v].nodeType && n.push(l[v]);
					else n.push.apply(n, u);
					else m(u, n);
					return p && (c(p, a, n, i), c.uniqueSort(n)), n
				};
				c.uniqueSort = function(e) {
					if (y && (a = s, e.sort(y), a))
						for (var t = 1; t < e.length; t++) e[t] === e[t - 1] && e.splice(t--, 1);
					return e
				}, c.matches = function(e, t) {
					return c(e, null, null, t)
				}, c.matchesSelector = function(e, t) {
					return c(t, null, null, [e]).length > 0
				}, c.find = function(e, t, n) {
					var r;
					if (!e) return [];
					for (var i = 0, o = f.order.length; o > i; i++) {
						var a, s = f.order[i];
						if (a = f.leftMatch[s].exec(e)) {
							var u = a[1];
							if (a.splice(1, 1), "\\" !== u.substr(u.length - 1) && (a[1] = (a[1] || "").replace(l, ""), r = f.find[s](a, t, n), null != r)) {
								e = e.replace(f.match[s], "");
								break
							}
						}
					}
					return r || (r = "undefined" != typeof t.getElementsByTagName ? t.getElementsByTagName("*") : []), {
						set: r,
						expr: e
					}
				}, c.filter = function(e, n, r, i) {
					for (var o, a, s = e, l = [], u = n, d = n && n[0] && c.isXML(n[0]); e && n.length;) {
						for (var p in f.filter)
							if (null != (o = f.leftMatch[p].exec(e)) && o[2]) {
								var h, m, g = f.filter[p],
									y = o[1];
								if (a = !1, o.splice(1, 1), "\\" === y.substr(y.length - 1)) continue;
								if (u === l && (l = []), f.preFilter[p])
									if (o = f.preFilter[p](o, u, r, l, i, d)) {
										if (o === !0) continue
									} else a = h = !0;
								if (o)
									for (var v = 0; null != (m = u[v]); v++)
										if (m) {
											h = g(m, o, v, u);
											var b = i ^ !!h;
											r && null != h ? b ? a = !0 : u[v] = !1 : b && (l.push(m), a = !0)
										}
								if (h !== t) {
									if (r || (u = l), e = e.replace(f.match[p], ""), !a) return [];
									break
								}
							}
						if (e === s) {
							if (null != a) break;
							c.error(e)
						}
						s = e
					}
					return u
				}, c.error = function(e) {
					throw "Syntax error, unrecognized expression: " + e
				};
				var f = c.selectors = {
						order: ["ID", "NAME", "TAG"],
						match: {
							ID: /#((?:[\w\u00c0-\uFFFF\-]|\\.)+)/,
							CLASS: /\.((?:[\w\u00c0-\uFFFF\-]|\\.)+)/,
							NAME: /\[name=['"]*((?:[\w\u00c0-\uFFFF\-]|\\.)+)['"]*\]/,
							ATTR: /\[\s*((?:[\w\u00c0-\uFFFF\-]|\\.)+)\s*(?:(\S?=)\s*(?:(['"])(.*?)\3|(#?(?:[\w\u00c0-\uFFFF\-]|\\.)*)|)|)\s*\]/,
							TAG: /^((?:[\w\u00c0-\uFFFF\*\-]|\\.)+)/,
							CHILD: /:(only|nth|last|first)-child(?:\(\s*(even|odd|(?:[+\-]?\d+|(?:[+\-]?\d*)?n\s*(?:[+\-]\s*\d+)?))\s*\))?/,
							POS: /:(nth|eq|gt|lt|first|last|even|odd)(?:\((\d*)\))?(?=[^\-]|$)/,
							PSEUDO: /:((?:[\w\u00c0-\uFFFF\-]|\\.)+)(?:\((['"]?)((?:\([^\)]+\)|[^\(\)]*)+)\2\))?/
						},
						leftMatch: {},
						attrMap: {
							"class": "className",
							"for": "htmlFor"
						},
						attrHandle: {
							href: function(e) {
								return e.getAttribute("href")
							},
							type: function(e) {
								return e.getAttribute("type")
							}
						},
						relative: {
							"+": function(e, t) {
								var n = "string" == typeof t,
									r = n && !u.test(t),
									i = n && !r;
								r && (t = t.toLowerCase());
								for (var o, a = 0, s = e.length; s > a; a++)
									if (o = e[a]) {
										for (;
											(o = o.previousSibling) && 1 !== o.nodeType;);
										e[a] = i || o && o.nodeName.toLowerCase() === t ? o || !1 : o === t
									}
								i && c.filter(t, e, !0)
							},
							">": function(e, t) {
								var n, r = "string" == typeof t,
									i = 0,
									o = e.length;
								if (r && !u.test(t)) {
									for (t = t.toLowerCase(); o > i; i++)
										if (n = e[i]) {
											var a = n.parentNode;
											e[i] = a.nodeName.toLowerCase() === t ? a : !1
										}
								} else {
									for (; o > i; i++) n = e[i], n && (e[i] = r ? n.parentNode : n.parentNode === t);
									r && c.filter(t, e, !0)
								}
							},
							"": function(t, r, o) {
								var a, s = i++,
									l = n;
								"string" != typeof r || u.test(r) || (r = r.toLowerCase(), a = r, l = e), l("parentNode", r, s, t, a, o)
							},
							"~": function(t, r, o) {
								var a, s = i++,
									l = n;
								"string" != typeof r || u.test(r) || (r = r.toLowerCase(), a = r, l = e), l("previousSibling", r, s, t, a, o)
							}
						},
						find: {
							ID: function(e, t, n) {
								if ("undefined" != typeof t.getElementById && !n) {
									var r = t.getElementById(e[1]);
									return r && r.parentNode ? [r] : []
								}
							},
							NAME: function(e, t) {
								if ("undefined" != typeof t.getElementsByName) {
									for (var n = [], r = t.getElementsByName(e[1]), i = 0, o = r.length; o > i; i++) r[i].getAttribute("name") === e[1] && n.push(r[i]);
									return 0 === n.length ? null : n
								}
							},
							TAG: function(e, t) {
								return "undefined" != typeof t.getElementsByTagName ? t.getElementsByTagName(e[1]) : void 0
							}
						},
						preFilter: {
							CLASS: function(e, t, n, r, i, o) {
								if (e = " " + e[1].replace(l, "") + " ", o) return e;
								for (var a, s = 0; null != (a = t[s]); s++) a && (i ^ (a.className && (" " + a.className + " ").replace(/[\t\n\r]/g, " ").indexOf(e) >= 0) ? n || r.push(a) : n && (t[s] = !1));
								return !1
							},
							ID: function(e) {
								return e[1].replace(l, "")
							},
							TAG: function(e, t) {
								return e[1].replace(l, "").toLowerCase()
							},
							CHILD: function(e) {
								if ("nth" === e[1]) {
									e[2] || c.error(e[0]), e[2] = e[2].replace(/^\+|\s*/g, "");
									var t = /(-?)(\d*)(?:n([+\-]?\d*))?/.exec("even" === e[2] && "2n" || "odd" === e[2] && "2n+1" || !/\D/.test(e[2]) && "0n+" + e[2] || e[2]);
									e[2] = t[1] + (t[2] || 1) - 0, e[3] = t[3] - 0
								} else e[2] && c.error(e[0]);
								return e[0] = i++, e
							},
							ATTR: function(e, t, n, r, i, o) {
								var a = e[1] = e[1].replace(l, "");
								return !o && f.attrMap[a] && (e[1] = f.attrMap[a]), e[4] = (e[4] || e[5] || "").replace(l, ""), "~=" === e[2] && (e[4] = " " + e[4] + " "), e
							},
							PSEUDO: function(e, t, n, i, o) {
								if ("not" === e[1]) {
									if (!((r.exec(e[3]) || "").length > 1 || /^\w/.test(e[3]))) {
										var a = c.filter(e[3], t, n, !0 ^ o);
										return n || i.push.apply(i, a), !1
									}
									e[3] = c(e[3], null, null, t)
								} else if (f.match.POS.test(e[0]) || f.match.CHILD.test(e[0])) return !0;
								return e
							},
							POS: function(e) {
								return e.unshift(!0), e
							}
						},
						filters: {
							enabled: function(e) {
								return e.disabled === !1 && "hidden" !== e.type
							},
							disabled: function(e) {
								return e.disabled === !0
							},
							checked: function(e) {
								return e.checked === !0
							},
							selected: function(e) {
								return e.parentNode && e.parentNode.selectedIndex, e.selected === !0
							},
							parent: function(e) {
								return !!e.firstChild
							},
							empty: function(e) {
								return !e.firstChild
							},
							has: function(e, t, n) {
								return !!c(n[3], e).length
							},
							header: function(e) {
								return /h\d/i.test(e.nodeName)
							},
							text: function(e) {
								var t = e.getAttribute("type"),
									n = e.type;
								return "input" === e.nodeName.toLowerCase() && "text" === n && (t === n || null === t)
							},
							radio: function(e) {
								return "input" === e.nodeName.toLowerCase() && "radio" === e.type
							},
							checkbox: function(e) {
								return "input" === e.nodeName.toLowerCase() && "checkbox" === e.type
							},
							file: function(e) {
								return "input" === e.nodeName.toLowerCase() && "file" === e.type
							},
							password: function(e) {
								return "input" === e.nodeName.toLowerCase() && "password" === e.type
							},
							submit: function(e) {
								var t = e.nodeName.toLowerCase();
								return ("input" === t || "button" === t) && "submit" === e.type
							},
							image: function(e) {
								return "input" === e.nodeName.toLowerCase() && "image" === e.type
							},
							reset: function(e) {
								var t = e.nodeName.toLowerCase();
								return ("input" === t || "button" === t) && "reset" === e.type
							},
							button: function(e) {
								var t = e.nodeName.toLowerCase();
								return "input" === t && "button" === e.type || "button" === t
							},
							input: function(e) {
								return /input|select|textarea|button/i.test(e.nodeName)
							},
							focus: function(e) {
								return e === e.ownerDocument.activeElement
							}
						},
						setFilters: {
							first: function(e, t) {
								return 0 === t
							},
							last: function(e, t, n, r) {
								return t === r.length - 1
							},
							even: function(e, t) {
								return t % 2 === 0
							},
							odd: function(e, t) {
								return t % 2 === 1
							},
							lt: function(e, t, n) {
								return t < n[3] - 0
							},
							gt: function(e, t, n) {
								return t > n[3] - 0
							},
							nth: function(e, t, n) {
								return n[3] - 0 === t
							},
							eq: function(e, t, n) {
								return n[3] - 0 === t
							}
						},
						filter: {
							PSEUDO: function(e, t, n, r) {
								var i = t[1],
									o = f.filters[i];
								if (o) return o(e, n, t, r);
								if ("contains" === i) return (e.textContent || e.innerText || c.getText([e]) || "").indexOf(t[3]) >= 0;
								if ("not" === i) {
									for (var a = t[3], s = 0, l = a.length; l > s; s++)
										if (a[s] === e) return !1;
									return !0
								}
								c.error(i)
							},
							CHILD: function(e, t) {
								var n = t[1],
									r = e;
								switch (n) {
									case "only":
									case "first":
										for (; r = r.previousSibling;)
											if (1 === r.nodeType) return !1;
										if ("first" === n) return !0;
										r = e;
									case "last":
										for (; r = r.nextSibling;)
											if (1 === r.nodeType) return !1;
										return !0;
									case "nth":
										var i = t[2],
											o = t[3];
										if (1 === i && 0 === o) return !0;
										var a = t[0],
											s = e.parentNode;
										if (s && (s.sizcache !== a || !e.nodeIndex)) {
											var l = 0;
											for (r = s.firstChild; r; r = r.nextSibling) 1 === r.nodeType && (r.nodeIndex = ++l);
											s.sizcache = a
										}
										var u = e.nodeIndex - o;
										return 0 === i ? 0 === u : u % i === 0 && u / i >= 0
								}
							},
							ID: function(e, t) {
								return 1 === e.nodeType && e.getAttribute("id") === t
							},
							TAG: function(e, t) {
								return "*" === t && 1 === e.nodeType || e.nodeName.toLowerCase() === t
							},
							CLASS: function(e, t) {
								return (" " + (e.className || e.getAttribute("class")) + " ").indexOf(t) > -1
							},
							ATTR: function(e, t) {
								var n = t[1],
									r = f.attrHandle[n] ? f.attrHandle[n](e) : null != e[n] ? e[n] : e.getAttribute(n),
									i = r + "",
									o = t[2],
									a = t[4];
								return null == r ? "!=" === o : "=" === o ? i === a : "*=" === o ? i.indexOf(a) >= 0 : "~=" === o ? (" " + i + " ").indexOf(a) >= 0 : a ? "!=" === o ? i !== a : "^=" === o ? 0 === i.indexOf(a) : "$=" === o ? i.substr(i.length - a.length) === a : "|=" === o ? i === a || i.substr(0, a.length + 1) === a + "-" : !1 : i && r !== !1
							},
							POS: function(e, t, n, r) {
								var i = t[2],
									o = f.setFilters[i];
								return o ? o(e, n, t, r) : void 0
							}
						}
					},
					d = f.match.POS,
					p = function(e, t) {
						return "\\" + (t - 0 + 1)
					};
				for (var h in f.match) f.match[h] = new RegExp(f.match[h].source + /(?![^\[]*\])(?![^\(]*\))/.source), f.leftMatch[h] = new RegExp(/(^(?:.|\r|\n)*?)/.source + f.match[h].source.replace(/\\(\d+)/g, p));
				var m = function(e, t) {
					return e = Array.prototype.slice.call(e, 0), t ? (t.push.apply(t, e), t) : e
				};
				try {
					Array.prototype.slice.call(L.documentElement.childNodes, 0)[0].nodeType
				} catch (g) {
					m = function(e, t) {
						var n = 0,
							r = t || [];
						if ("[object Array]" === o.call(e)) Array.prototype.push.apply(r, e);
						else if ("number" == typeof e.length)
							for (var i = e.length; i > n; n++) r.push(e[n]);
						else
							for (; e[n]; n++) r.push(e[n]);
						return r
					}
				}
				var y, v;
				L.documentElement.compareDocumentPosition ? y = function(e, t) {
						return e === t ? (a = !0, 0) : e.compareDocumentPosition && t.compareDocumentPosition ? 4 & e.compareDocumentPosition(t) ? -1 : 1 : e.compareDocumentPosition ? -1 : 1
					} : (y = function(e, t) {
						if (e === t) return a = !0, 0;
						if (e.sourceIndex && t.sourceIndex) return e.sourceIndex - t.sourceIndex;
						var n, r, i = [],
							o = [],
							s = e.parentNode,
							l = t.parentNode,
							u = s;
						if (s === l) return v(e, t);
						if (!s) return -1;
						if (!l) return 1;
						for (; u;) i.unshift(u), u = u.parentNode;
						for (u = l; u;) o.unshift(u), u = u.parentNode;
						n = i.length, r = o.length;
						for (var c = 0; n > c && r > c; c++)
							if (i[c] !== o[c]) return v(i[c], o[c]);
						return c === n ? v(e, o[c], -1) : v(i[c], t, 1)
					}, v = function(e, t, n) {
						if (e === t) return n;
						for (var r = e.nextSibling; r;) {
							if (r === t) return -1;
							r = r.nextSibling
						}
						return 1
					}), c.getText = function(e) {
						for (var t, n = "", r = 0; e[r]; r++) t = e[r], 3 === t.nodeType || 4 === t.nodeType ? n += t.nodeValue : 8 !== t.nodeType && (n += c.getText(t.childNodes));
						return n
					},
					function() {
						var e = L.createElement("div"),
							n = "script" + (new Date).getTime(),
							r = L.documentElement;
						e.innerHTML = "<a name='" + n + "'/>", r.insertBefore(e, r.firstChild), L.getElementById(n) && (f.find.ID = function(e, n, r) {
							if ("undefined" != typeof n.getElementById && !r) {
								var i = n.getElementById(e[1]);
								return i ? i.id === e[1] || "undefined" != typeof i.getAttributeNode && i.getAttributeNode("id").nodeValue === e[1] ? [i] : t : []
							}
						}, f.filter.ID = function(e, t) {
							var n = "undefined" != typeof e.getAttributeNode && e.getAttributeNode("id");
							return 1 === e.nodeType && n && n.nodeValue === t
						}), r.removeChild(e), r = e = null
					}(),
					function() {
						var e = L.createElement("div");
						e.appendChild(L.createComment("")), e.getElementsByTagName("*").length > 0 && (f.find.TAG = function(e, t) {
							var n = t.getElementsByTagName(e[1]);
							if ("*" === e[1]) {
								for (var r = [], i = 0; n[i]; i++) 1 === n[i].nodeType && r.push(n[i]);
								n = r
							}
							return n
						}), e.innerHTML = "<a href='#'></a>", e.firstChild && "undefined" != typeof e.firstChild.getAttribute && "#" !== e.firstChild.getAttribute("href") && (f.attrHandle.href = function(e) {
							return e.getAttribute("href", 2)
						}), e = null
					}(), L.querySelectorAll && ! function() {
						var e = c,
							t = L.createElement("div"),
							n = "__sizzle__";
						if (t.innerHTML = "<p class='TEST'></p>", !t.querySelectorAll || 0 !== t.querySelectorAll(".TEST").length) {
							c = function(t, r, i, o) {
								if (r = r || L, !o && !c.isXML(r)) {
									var a = /^(\w+$)|^\.([\w\-]+$)|^#([\w\-]+$)/.exec(t);
									if (a && (1 === r.nodeType || 9 === r.nodeType)) {
										if (a[1]) return m(r.getElementsByTagName(t), i);
										if (a[2] && f.find.CLASS && r.getElementsByClassName) return m(r.getElementsByClassName(a[2]), i)
									}
									if (9 === r.nodeType) {
										if ("body" === t && r.body) return m([r.body], i);
										if (a && a[3]) {
											var s = r.getElementById(a[3]);
											if (!s || !s.parentNode) return m([], i);
											if (s.id === a[3]) return m([s], i)
										}
										try {
											return m(r.querySelectorAll(t), i)
										} catch (l) {}
									} else if (1 === r.nodeType && "object" !== r.nodeName.toLowerCase()) {
										var u = r,
											d = r.getAttribute("id"),
											p = d || n,
											h = r.parentNode,
											g = /^\s*[+~]/.test(t);
										d ? p = p.replace(/'/g, "\\$&") : r.setAttribute("id", p), g && h && (r = r.parentNode);
										try {
											if (!g || h) return m(r.querySelectorAll("[id='" + p + "'] " + t), i)
										} catch (y) {} finally {
											d || u.removeAttribute("id")
										}
									}
								}
								return e(t, r, i, o)
							};
							for (var r in e) c[r] = e[r];
							t = null
						}
					}(),
					function() {
						var e = L.documentElement,
							t = e.matchesSelector || e.mozMatchesSelector || e.webkitMatchesSelector || e.msMatchesSelector;
						if (t) {
							var n = !t.call(L.createElement("div"), "div"),
								r = !1;
							try {
								t.call(L.documentElement, "[test!='']:sizzle")
							} catch (i) {
								r = !0
							}
							c.matchesSelector = function(e, i) {
								if (i = i.replace(/\=\s*([^'"\]]*)\s*\]/g, "='$1']"), !c.isXML(e)) try {
									if (r || !f.match.PSEUDO.test(i) && !/!=/.test(i)) {
										var o = t.call(e, i);
										if (o || !n || e.document && 11 !== e.document.nodeType) return o
									}
								} catch (a) {}
								return c(i, null, null, [e]).length > 0
							}
						}
					}(),
					function() {
						var e = L.createElement("div");
						e.innerHTML = "<div class='test e'></div><div class='test'></div>", e.getElementsByClassName && 0 !== e.getElementsByClassName("e").length && (e.lastChild.className = "e", 1 !== e.getElementsByClassName("e").length && (f.order.splice(1, 0, "CLASS"), f.find.CLASS = function(e, t, n) {
							return "undefined" == typeof t.getElementsByClassName || n ? void 0 : t.getElementsByClassName(e[1])
						}, e = null))
					}(), L.documentElement.contains ? c.contains = function(e, t) {
						return e !== t && (e.contains ? e.contains(t) : !0)
					} : L.documentElement.compareDocumentPosition ? c.contains = function(e, t) {
						return !!(16 & e.compareDocumentPosition(t))
					} : c.contains = function() {
						return !1
					}, c.isXML = function(e) {
						var t = (e ? e.ownerDocument || e : 0).documentElement;
						return t ? "HTML" !== t.nodeName : !1
					};
				var b = function(e, t) {
					for (var n, r = [], i = "", o = t.nodeType ? [t] : t; n = f.match.PSEUDO.exec(e);) i += n[0], e = e.replace(f.match.PSEUDO, "");
					e = f.relative[e] ? e + "*" : e;
					for (var a = 0, s = o.length; s > a; a++) c(e, o[a], r);
					return c.filter(i, r)
				};
				H.find = c, H.expr = c.selectors, H.expr[":"] = H.expr.filters, H.unique = c.uniqueSort, H.text = c.getText, H.isXMLDoc = c.isXML, H.contains = c.contains
			}();
		var le = /Until$/,
			ue = /^(?:parents|prevUntil|prevAll)/,
			ce = /,/,
			fe = /^.[^:#\[\.,]*$/,
			de = Array.prototype.slice,
			pe = H.expr.match.POS,
			he = {
				children: !0,
				contents: !0,
				next: !0,
				prev: !0
			};
		H.fn.extend({
			find: function(e) {
				var t, n, r = this;
				if ("string" != typeof e) return H(e).filter(function() {
					for (t = 0, n = r.length; n > t; t++)
						if (H.contains(r[t], this)) return !0
				});
				var i, o, a, s = this.pushStack("", "find", e);
				for (t = 0, n = this.length; n > t; t++)
					if (i = s.length, H.find(e, this[t], s), t > 0)
						for (o = i; o < s.length; o++)
							for (a = 0; i > a; a++)
								if (s[a] === s[o]) {
									s.splice(o--, 1);
									break
								}
				return s
			},
			has: function(e) {
				var t = H(e);
				return this.filter(function() {
					for (var e = 0, n = t.length; n > e; e++)
						if (H.contains(this, t[e])) return !0
				})
			},
			not: function(e) {
				return this.pushStack(f(this, e, !1), "not", e)
			},
			filter: function(e) {
				return this.pushStack(f(this, e, !0), "filter", e)
			},
			is: function(e) {
				return !!e && ("string" == typeof e ? H.filter(e, this).length > 0 : this.filter(e).length > 0)
			},
			closest: function(e, t) {
				var n, r, i = [],
					o = this[0];
				if (H.isArray(e)) {
					var a, s, l = {},
						u = 1;
					if (o && e.length) {
						for (n = 0, r = e.length; r > n; n++) s = e[n], l[s] || (l[s] = pe.test(s) ? H(s, t || this.context) : s);
						for (; o && o.ownerDocument && o !== t;) {
							for (s in l) a = l[s], (a.jquery ? a.index(o) > -1 : H(o).is(a)) && i.push({
								selector: s,
								elem: o,
								level: u
							});
							o = o.parentNode, u++
						}
					}
					return i
				}
				var c = pe.test(e) || "string" != typeof e ? H(e, t || this.context) : 0;
				for (n = 0, r = this.length; r > n; n++)
					for (o = this[n]; o;) {
						if (c ? c.index(o) > -1 : H.find.matchesSelector(o, e)) {
							i.push(o);
							break
						}
						if (o = o.parentNode, !o || !o.ownerDocument || o === t || 11 === o.nodeType) break
					}
				return i = i.length > 1 ? H.unique(i) : i, this.pushStack(i, "closest", e)
			},
			index: function(e) {
				return e && "string" != typeof e ? H.inArray(e.jquery ? e[0] : e, this) : H.inArray(this[0], e ? H(e) : this.parent().children())
			},
			add: function(e, t) {
				var n = "string" == typeof e ? H(e, t) : H.makeArray(e && e.nodeType ? [e] : e),
					r = H.merge(this.get(), n);
				return this.pushStack(c(n[0]) || c(r[0]) ? r : H.unique(r))
			},
			andSelf: function() {
				return this.add(this.prevObject)
			}
		}), H.each({
			parent: function(e) {
				var t = e.parentNode;
				return t && 11 !== t.nodeType ? t : null
			},
			parents: function(e) {
				return H.dir(e, "parentNode")
			},
			parentsUntil: function(e, t, n) {
				return H.dir(e, "parentNode", n)
			},
			next: function(e) {
				return H.nth(e, 2, "nextSibling")
			},
			prev: function(e) {
				return H.nth(e, 2, "previousSibling")
			},
			nextAll: function(e) {
				return H.dir(e, "nextSibling")
			},
			prevAll: function(e) {
				return H.dir(e, "previousSibling")
			},
			nextUntil: function(e, t, n) {
				return H.dir(e, "nextSibling", n)
			},
			prevUntil: function(e, t, n) {
				return H.dir(e, "previousSibling", n)
			},
			siblings: function(e) {
				return H.sibling(e.parentNode.firstChild, e)
			},
			children: function(e) {
				return H.sibling(e.firstChild)
			},
			contents: function(e) {
				return H.nodeName(e, "iframe") ? e.contentDocument || e.contentWindow.document : H.makeArray(e.childNodes)
			}
		}, function(e, t) {
			H.fn[e] = function(n, r) {
				var i = H.map(this, t, n),
					o = de.call(arguments);
				return le.test(e) || (r = n), r && "string" == typeof r && (i = H.filter(r, i)), i = this.length > 1 && !he[e] ? H.unique(i) : i, (this.length > 1 || ce.test(r)) && ue.test(e) && (i = i.reverse()), this.pushStack(i, e, o.join(","))
			}
		}), H.extend({
			filter: function(e, t, n) {
				return n && (e = ":not(" + e + ")"), 1 === t.length ? H.find.matchesSelector(t[0], e) ? [t[0]] : [] : H.find.matches(e, t)
			},
			dir: function(e, n, r) {
				for (var i = [], o = e[n]; o && 9 !== o.nodeType && (r === t || 1 !== o.nodeType || !H(o).is(r));) 1 === o.nodeType && i.push(o), o = o[n];
				return i
			},
			nth: function(e, t, n, r) {
				t = t || 1;
				for (var i = 0; e && (1 !== e.nodeType || ++i !== t); e = e[n]);
				return e
			},
			sibling: function(e, t) {
				for (var n = []; e; e = e.nextSibling) 1 === e.nodeType && e !== t && n.push(e);
				return n
			}
		});
		var me = / jQuery\d+="(?:\d+|null)"/g,
			ge = /^\s+/,
			ye = /<(?!area|br|col|embed|hr|img|input|link|meta|param)(([\w:]+)[^>]*)\/>/gi,
			ve = /<([\w:]+)/,
			be = /<tbody/i,
			xe = /<|&#?\w+;/,
			Te = /<(?:script|object|embed|option|style)/i,
			we = /checked\s*(?:[^=]|=\s*.checked.)/i,
			Ne = /\/(java|ecma)script/i,
			Ce = /^\s*<!(?:\[CDATA\[|\-\-)/,
			Ee = {
				option: [1, "<select multiple='multiple'>", "</select>"],
				legend: [1, "<fieldset>", "</fieldset>"],
				thead: [1, "<table>", "</table>"],
				tr: [2, "<table><tbody>", "</tbody></table>"],
				td: [3, "<table><tbody><tr>", "</tr></tbody></table>"],
				col: [2, "<table><tbody></tbody><colgroup>", "</colgroup></table>"],
				area: [1, "<map>", "</map>"],
				_default: [0, "", ""]
			};
		Ee.optgroup = Ee.option, Ee.tbody = Ee.tfoot = Ee.colgroup = Ee.caption = Ee.thead, Ee.th = Ee.td, H.support.htmlSerialize || (Ee._default = [1, "div<div>", "</div>"]), H.fn.extend({
			text: function(e) {
				return H.isFunction(e) ? this.each(function(t) {
					var n = H(this);
					n.text(e.call(this, t, n.text()))
				}) : "object" != typeof e && e !== t ? this.empty().append((this[0] && this[0].ownerDocument || L).createTextNode(e)) : H.text(this)
			},
			wrapAll: function(e) {
				if (H.isFunction(e)) return this.each(function(t) {
					H(this).wrapAll(e.call(this, t))
				});
				if (this[0]) {
					var t = H(e, this[0].ownerDocument).eq(0).clone(!0);
					this[0].parentNode && t.insertBefore(this[0]), t.map(function() {
						for (var e = this; e.firstChild && 1 === e.firstChild.nodeType;) e = e.firstChild;
						return e
					}).append(this)
				}
				return this
			},
			wrapInner: function(e) {
				return this.each(H.isFunction(e) ? function(t) {
					H(this).wrapInner(e.call(this, t))
				} : function() {
					var t = H(this),
						n = t.contents();
					n.length ? n.wrapAll(e) : t.append(e)
				})
			},
			wrap: function(e) {
				return this.each(function() {
					H(this).wrapAll(e)
				})
			},
			unwrap: function() {
				return this.parent().each(function() {
					H.nodeName(this, "body") || H(this).replaceWith(this.childNodes)
				}).end()
			},
			append: function() {
				return this.domManip(arguments, !0, function(e) {
					1 === this.nodeType && this.appendChild(e)
				})
			},
			prepend: function() {
				return this.domManip(arguments, !0, function(e) {
					1 === this.nodeType && this.insertBefore(e, this.firstChild)
				})
			},
			before: function() {
				if (this[0] && this[0].parentNode) return this.domManip(arguments, !1, function(e) {
					this.parentNode.insertBefore(e, this)
				});
				if (arguments.length) {
					var e = H(arguments[0]);
					return e.push.apply(e, this.toArray()), this.pushStack(e, "before", arguments)
				}
			},
			after: function() {
				if (this[0] && this[0].parentNode) return this.domManip(arguments, !1, function(e) {
					this.parentNode.insertBefore(e, this.nextSibling)
				});
				if (arguments.length) {
					var e = this.pushStack(this, "after", arguments);
					return e.push.apply(e, H(arguments[0]).toArray()), e
				}
			},
			remove: function(e, t) {
				for (var n, r = 0; null != (n = this[r]); r++)(!e || H.filter(e, [n]).length) && (t || 1 !== n.nodeType || (H.cleanData(n.getElementsByTagName("*")), H.cleanData([n])), n.parentNode && n.parentNode.removeChild(n));
				return this
			},
			empty: function() {
				for (var e, t = 0; null != (e = this[t]); t++)
					for (1 === e.nodeType && H.cleanData(e.getElementsByTagName("*")); e.firstChild;) e.removeChild(e.firstChild);
				return this
			},
			clone: function(e, t) {
				return e = null == e ? !1 : e, t = null == t ? e : t, this.map(function() {
					return H.clone(this, e, t)
				})
			},
			html: function(e) {
				if (e === t) return this[0] && 1 === this[0].nodeType ? this[0].innerHTML.replace(me, "") : null;
				if ("string" != typeof e || Te.test(e) || !H.support.leadingWhitespace && ge.test(e) || Ee[(ve.exec(e) || ["", ""])[1].toLowerCase()]) H.isFunction(e) ? this.each(function(t) {
					var n = H(this);
					n.html(e.call(this, t, n.html()))
				}) : this.empty().append(e);
				else {
					e = e.replace(ye, "<$1></$2>");
					try {
						for (var n = 0, r = this.length; r > n; n++) 1 === this[n].nodeType && (H.cleanData(this[n].getElementsByTagName("*")), this[n].innerHTML = e)
					} catch (i) {
						this.empty().append(e)
					}
				}
				return this
			},
			replaceWith: function(e) {
				return this[0] && this[0].parentNode ? H.isFunction(e) ? this.each(function(t) {
					var n = H(this),
						r = n.html();
					n.replaceWith(e.call(this, t, r))
				}) : ("string" != typeof e && (e = H(e).detach()), this.each(function() {
					var t = this.nextSibling,
						n = this.parentNode;
					H(this).remove(), t ? H(t).before(e) : H(n).append(e)
				})) : this.length ? this.pushStack(H(H.isFunction(e) ? e() : e), "replaceWith", e) : this
			},
			detach: function(e) {
				return this.remove(e, !0)
			},
			domManip: function(e, n, r) {
				var i, o, a, s, l = e[0],
					u = [];
				if (!H.support.checkClone && 3 === arguments.length && "string" == typeof l && we.test(l)) return this.each(function() {
					H(this).domManip(e, n, r, !0)
				});
				if (H.isFunction(l)) return this.each(function(i) {
					var o = H(this);
					e[0] = l.call(this, i, n ? o.html() : t), o.domManip(e, n, r)
				});
				if (this[0]) {
					if (s = l && l.parentNode, i = H.support.parentNode && s && 11 === s.nodeType && s.childNodes.length === this.length ? {
							fragment: s
						} : H.buildFragment(e, this, u), a = i.fragment, o = 1 === a.childNodes.length ? a = a.firstChild : a.firstChild) {
						n = n && H.nodeName(o, "tr");
						for (var c = 0, f = this.length, p = f - 1; f > c; c++) r.call(n ? d(this[c], o) : this[c], i.cacheable || f > 1 && p > c ? H.clone(a, !0, !0) : a)
					}
					u.length && H.each(u, v)
				}
				return this
			}
		}), H.buildFragment = function(e, t, n) {
			var r, i, o, a;
			return t && t[0] && (a = t[0].ownerDocument || t[0]), a.createDocumentFragment || (a = L), 1 === e.length && "string" == typeof e[0] && e[0].length < 512 && a === L && "<" === e[0].charAt(0) && !Te.test(e[0]) && (H.support.checkClone || !we.test(e[0])) && (i = !0,
				o = H.fragments[e[0]], o && 1 !== o && (r = o)), r || (r = a.createDocumentFragment(), H.clean(e, a, r, n)), i && (H.fragments[e[0]] = o ? r : 1), {
				fragment: r,
				cacheable: i
			}
		}, H.fragments = {}, H.each({
			appendTo: "append",
			prependTo: "prepend",
			insertBefore: "before",
			insertAfter: "after",
			replaceAll: "replaceWith"
		}, function(e, t) {
			H.fn[e] = function(n) {
				var r = [],
					i = H(n),
					o = 1 === this.length && this[0].parentNode;
				if (o && 11 === o.nodeType && 1 === o.childNodes.length && 1 === i.length) return i[t](this[0]), this;
				for (var a = 0, s = i.length; s > a; a++) {
					var l = (a > 0 ? this.clone(!0) : this).get();
					H(i[a])[t](l), r = r.concat(l)
				}
				return this.pushStack(r, e, i.selector)
			}
		}), H.extend({
			clone: function(e, t, n) {
				var r, i, o, a = e.cloneNode(!0);
				if (!(H.support.noCloneEvent && H.support.noCloneChecked || 1 !== e.nodeType && 11 !== e.nodeType || H.isXMLDoc(e)))
					for (h(e, a), r = m(e), i = m(a), o = 0; r[o]; ++o) h(r[o], i[o]);
				if (t && (p(e, a), n))
					for (r = m(e), i = m(a), o = 0; r[o]; ++o) p(r[o], i[o]);
				return r = i = null, a
			},
			clean: function(e, t, n, r) {
				var i;
				t = t || L, "undefined" == typeof t.createElement && (t = t.ownerDocument || t[0] && t[0].ownerDocument || L);
				for (var o, a, s = [], l = 0; null != (a = e[l]); l++)
					if ("number" == typeof a && (a += ""), a) {
						if ("string" == typeof a)
							if (xe.test(a)) {
								a = a.replace(ye, "<$1></$2>");
								var u = (ve.exec(a) || ["", ""])[1].toLowerCase(),
									c = Ee[u] || Ee._default,
									f = c[0],
									d = t.createElement("div");
								for (d.innerHTML = c[1] + a + c[2]; f--;) d = d.lastChild;
								if (!H.support.tbody) {
									var p = be.test(a),
										h = "table" !== u || p ? "<table>" !== c[1] || p ? [] : d.childNodes : d.firstChild && d.firstChild.childNodes;
									for (o = h.length - 1; o >= 0; --o) H.nodeName(h[o], "tbody") && !h[o].childNodes.length && h[o].parentNode.removeChild(h[o])
								}!H.support.leadingWhitespace && ge.test(a) && d.insertBefore(t.createTextNode(ge.exec(a)[0]), d.firstChild), a = d.childNodes
							} else a = t.createTextNode(a);
						var m;
						if (!H.support.appendChecked)
							if (a[0] && "number" == typeof(m = a.length))
								for (o = 0; m > o; o++) y(a[o]);
							else y(a);
						a.nodeType ? s.push(a) : s = H.merge(s, a)
					}
				if (n)
					for (i = function(e) {
							return !e.type || Ne.test(e.type)
						}, l = 0; s[l]; l++)
						if (!r || !H.nodeName(s[l], "script") || s[l].type && "text/javascript" !== s[l].type.toLowerCase()) {
							if (1 === s[l].nodeType) {
								var g = H.grep(s[l].getElementsByTagName("script"), i);
								s.splice.apply(s, [l + 1, 0].concat(g))
							}
							n.appendChild(s[l])
						} else r.push(s[l].parentNode ? s[l].parentNode.removeChild(s[l]) : s[l]);
				return s
			},
			cleanData: function(e) {
				for (var t, n, r, i = H.cache, o = H.expando, a = H.event.special, s = H.support.deleteExpando, l = 0; null != (r = e[l]); l++)
					if ((!r.nodeName || !H.noData[r.nodeName.toLowerCase()]) && (n = r[H.expando])) {
						if (t = i[n] && i[n][o], t && t.events) {
							for (var u in t.events) a[u] ? H.event.remove(r, u) : H.removeEvent(r, u, t.handle);
							t.handle && (t.handle.elem = null)
						}
						s ? delete r[H.expando] : r.removeAttribute && r.removeAttribute(H.expando), delete i[n]
					}
			}
		});
		var Se, Ae, ke, De = /alpha\([^)]*\)/i,
			Fe = /opacity=([^)]*)/,
			je = /([A-Z]|^ms)/g,
			Le = /^-?\d+(?:px)?$/i,
			Me = /^-?\d/,
			Oe = /^[+\-]=/,
			He = /[^+\-\.\de]+/g,
			Be = {
				position: "absolute",
				visibility: "hidden",
				display: "block"
			},
			Pe = ["Left", "Right"],
			Ie = ["Top", "Bottom"];
		H.fn.css = function(e, n) {
			return 2 === arguments.length && n === t ? this : H.access(this, e, n, !0, function(e, n, r) {
				return r !== t ? H.style(e, n, r) : H.css(e, n)
			})
		}, H.extend({
			cssHooks: {
				opacity: {
					get: function(e, t) {
						if (t) {
							var n = Se(e, "opacity", "opacity");
							return "" === n ? "1" : n
						}
						return e.style.opacity
					}
				}
			},
			cssNumber: {
				fillOpacity: !0,
				fontWeight: !0,
				lineHeight: !0,
				opacity: !0,
				orphans: !0,
				widows: !0,
				zIndex: !0,
				zoom: !0
			},
			cssProps: {
				"float": H.support.cssFloat ? "cssFloat" : "styleFloat"
			},
			style: function(e, n, r, i) {
				if (e && 3 !== e.nodeType && 8 !== e.nodeType && e.style) {
					var o, a, s = H.camelCase(n),
						l = e.style,
						u = H.cssHooks[s];
					if (n = H.cssProps[s] || s, r === t) return u && "get" in u && (o = u.get(e, !1, i)) !== t ? o : l[n];
					if (a = typeof r, !("number" === a && isNaN(r) || null == r || ("string" === a && Oe.test(r) && (r = +r.replace(He, "") + parseFloat(H.css(e, n)), a = "number"), "number" !== a || H.cssNumber[s] || (r += "px"), u && "set" in u && (r = u.set(e, r)) === t))) try {
						l[n] = r
					} catch (c) {}
				}
			},
			css: function(e, n, r) {
				var i, o;
				return n = H.camelCase(n), o = H.cssHooks[n], n = H.cssProps[n] || n, "cssFloat" === n && (n = "float"), o && "get" in o && (i = o.get(e, !0, r)) !== t ? i : Se ? Se(e, n) : void 0
			},
			swap: function(e, t, n) {
				var r = {};
				for (var i in t) r[i] = e.style[i], e.style[i] = t[i];
				n.call(e);
				for (i in t) e.style[i] = r[i]
			}
		}), H.curCSS = H.css, H.each(["height", "width"], function(e, t) {
			H.cssHooks[t] = {
				get: function(e, n, r) {
					var i;
					return n ? 0 !== e.offsetWidth ? b(e, t, r) : (H.swap(e, Be, function() {
						i = b(e, t, r)
					}), i) : void 0
				},
				set: function(e, t) {
					return Le.test(t) ? (t = parseFloat(t), t >= 0 ? t + "px" : void 0) : t
				}
			}
		}), H.support.opacity || (H.cssHooks.opacity = {
			get: function(e, t) {
				return Fe.test((t && e.currentStyle ? e.currentStyle.filter : e.style.filter) || "") ? parseFloat(RegExp.$1) / 100 + "" : t ? "1" : ""
			},
			set: function(e, t) {
				var n = e.style,
					r = e.currentStyle;
				n.zoom = 1;
				var i = H.isNaN(t) ? "" : "alpha(opacity=" + 100 * t + ")",
					o = r && r.filter || n.filter || "";
				n.filter = De.test(o) ? o.replace(De, i) : o + " " + i
			}
		}), H(function() {
			H.support.reliableMarginRight || (H.cssHooks.marginRight = {
				get: function(e, t) {
					var n;
					return H.swap(e, {
						display: "inline-block"
					}, function() {
						n = t ? Se(e, "margin-right", "marginRight") : e.style.marginRight
					}), n
				}
			})
		}), L.defaultView && L.defaultView.getComputedStyle && (Ae = function(e, n) {
			var r, i, o;
			return n = n.replace(je, "-$1").toLowerCase(), (i = e.ownerDocument.defaultView) ? ((o = i.getComputedStyle(e, null)) && (r = o.getPropertyValue(n), "" !== r || H.contains(e.ownerDocument.documentElement, e) || (r = H.style(e, n))), r) : t
		}), L.documentElement.currentStyle && (ke = function(e, t) {
			var n, r = e.currentStyle && e.currentStyle[t],
				i = e.runtimeStyle && e.runtimeStyle[t],
				o = e.style;
			return !Le.test(r) && Me.test(r) && (n = o.left, i && (e.runtimeStyle.left = e.currentStyle.left), o.left = "fontSize" === t ? "1em" : r || 0, r = o.pixelLeft + "px", o.left = n, i && (e.runtimeStyle.left = i)), "" === r ? "auto" : r
		}), Se = Ae || ke, H.expr && H.expr.filters && (H.expr.filters.hidden = function(e) {
			var t = e.offsetWidth,
				n = e.offsetHeight;
			return 0 === t && 0 === n || !H.support.reliableHiddenOffsets && "none" === (e.style.display || H.css(e, "display"))
		}, H.expr.filters.visible = function(e) {
			return !H.expr.filters.hidden(e)
		});
		var qe, _e, Re = /%20/g,
			We = /\[\]$/,
			$e = /\r?\n/g,
			ze = /#.*$/,
			Xe = /^(.*?):[ \t]*([^\r\n]*)\r?$/gm,
			Ve = /^(?:color|date|datetime|email|hidden|month|number|password|range|search|tel|text|time|url|week)$/i,
			Ue = /^(?:about|app|app\-storage|.+\-extension|file|widget):$/,
			Je = /^(?:GET|HEAD)$/,
			Ye = /^\/\//,
			Ge = /\?/,
			Ke = /<script\b[^<]*(?:(?!<\/script>)<[^<]*)*<\/script>/gi,
			Qe = /^(?:select|textarea)/i,
			Ze = /\s+/,
			et = /([?&])_=[^&]*/,
			tt = /^([\w\+\.\-]+:)(?:\/\/([^\/?#:]*)(?::(\d+))?)?/,
			nt = H.fn.load,
			rt = {},
			it = {};
		try {
			qe = O.href
		} catch (ot) {
			qe = L.createElement("a"), qe.href = "", qe = qe.href
		}
		_e = tt.exec(qe.toLowerCase()) || [], H.fn.extend({
			load: function(e, n, r) {
				if ("string" != typeof e && nt) return nt.apply(this, arguments);
				if (!this.length) return this;
				var i = e.indexOf(" ");
				if (i >= 0) {
					var o = e.slice(i, e.length);
					e = e.slice(0, i)
				}
				var a = "GET";
				n && (H.isFunction(n) ? (r = n, n = t) : "object" == typeof n && (n = H.param(n, H.ajaxSettings.traditional), a = "POST"));
				var s = this;
				return H.ajax({
					url: e,
					type: a,
					dataType: "html",
					data: n,
					complete: function(e, t, n) {
						n = e.responseText, e.isResolved() && (e.done(function(e) {
							n = e
						}), s.html(o ? H("<div>").append(n.replace(Ke, "")).find(o) : n)), r && s.each(r, [n, t, e])
					}
				}), this
			},
			serialize: function() {
				return H.param(this.serializeArray())
			},
			serializeArray: function() {
				return this.map(function() {
					return this.elements ? H.makeArray(this.elements) : this
				}).filter(function() {
					return this.name && !this.disabled && (this.checked || Qe.test(this.nodeName) || Ve.test(this.type))
				}).map(function(e, t) {
					var n = H(this).val();
					return null == n ? null : H.isArray(n) ? H.map(n, function(e, n) {
						return {
							name: t.name,
							value: e.replace($e, "\r\n")
						}
					}) : {
						name: t.name,
						value: n.replace($e, "\r\n")
					}
				}).get()
			}
		}), H.each("ajaxStart ajaxStop ajaxComplete ajaxError ajaxSuccess ajaxSend".split(" "), function(e, t) {
			H.fn[t] = function(e) {
				return this.bind(t, e)
			}
		}), H.each(["get", "post"], function(e, n) {
			H[n] = function(e, r, i, o) {
				return H.isFunction(r) && (o = o || i, i = r, r = t), H.ajax({
					type: n,
					url: e,
					data: r,
					success: i,
					dataType: o
				})
			}
		}), H.extend({
			getScript: function(e, n) {
				return H.get(e, t, n, "script")
			},
			getJSON: function(e, t, n) {
				return H.get(e, t, n, "json")
			},
			ajaxSetup: function(e, t) {
				t ? H.extend(!0, e, H.ajaxSettings, t) : (t = e, e = H.extend(!0, H.ajaxSettings, t));
				for (var n in {
						context: 1,
						url: 1
					}) n in t ? e[n] = t[n] : n in H.ajaxSettings && (e[n] = H.ajaxSettings[n]);
				return e
			},
			ajaxSettings: {
				url: qe,
				isLocal: Ue.test(_e[1]),
				global: !0,
				type: "GET",
				contentType: "application/x-www-form-urlencoded",
				processData: !0,
				async: !0,
				accepts: {
					xml: "application/xml, text/xml",
					html: "text/html",
					text: "text/plain",
					json: "application/json, text/javascript",
					"*": "*/*"
				},
				contents: {
					xml: /xml/,
					html: /html/,
					json: /json/
				},
				responseFields: {
					xml: "responseXML",
					text: "responseText"
				},
				converters: {
					"* text": e.String,
					"text html": !0,
					"text json": H.parseJSON,
					"text xml": H.parseXML
				}
			},
			ajaxPrefilter: x(rt),
			ajaxTransport: x(it),
			ajax: function(e, n) {
				function r(e, n, r, a) {
					if (2 !== x) {
						x = 2, l && clearTimeout(l), s = t, o = a || "", w.readyState = e ? 4 : 0;
						var u, f, v, b, T, E = r ? N(d, w, r) : t;
						if (e >= 200 && 300 > e || 304 === e)
							if (d.ifModified && ((b = w.getResponseHeader("Last-Modified")) && (H.lastModified[i] = b), (T = w.getResponseHeader("Etag")) && (H.etag[i] = T)), 304 === e) n = "notmodified", u = !0;
							else try {
								f = C(d, E), n = "success", u = !0
							} catch (S) {
								n = "parsererror", v = S
							} else v = n, (!n || e) && (n = "error", 0 > e && (e = 0));
						w.status = e, w.statusText = n, u ? m.resolveWith(p, [f, n, w]) : m.rejectWith(p, [w, n, v]), w.statusCode(y), y = t, c && h.trigger("ajax" + (u ? "Success" : "Error"), [w, d, u ? f : v]), g.resolveWith(p, [w, n]), c && (h.trigger("ajaxComplete", [w, d]), --H.active || H.event.trigger("ajaxStop"))
					}
				}
				"object" == typeof e && (n = e, e = t), n = n || {};
				var i, o, a, s, l, u, c, f, d = H.ajaxSetup({}, n),
					p = d.context || d,
					h = p !== d && (p.nodeType || p instanceof H) ? H(p) : H.event,
					m = H.Deferred(),
					g = H._Deferred(),
					y = d.statusCode || {},
					v = {},
					b = {},
					x = 0,
					w = {
						readyState: 0,
						setRequestHeader: function(e, t) {
							if (!x) {
								var n = e.toLowerCase();
								e = b[n] = b[n] || e, v[e] = t
							}
							return this
						},
						getAllResponseHeaders: function() {
							return 2 === x ? o : null
						},
						getResponseHeader: function(e) {
							var n;
							if (2 === x) {
								if (!a)
									for (a = {}; n = Xe.exec(o);) a[n[1].toLowerCase()] = n[2];
								n = a[e.toLowerCase()]
							}
							return n === t ? null : n
						},
						overrideMimeType: function(e) {
							return x || (d.mimeType = e), this
						},
						abort: function(e) {
							return e = e || "abort", s && s.abort(e), r(0, e), this
						}
					};
				if (m.promise(w), w.success = w.done, w.error = w.fail, w.complete = g.done, w.statusCode = function(e) {
						if (e) {
							var t;
							if (2 > x)
								for (t in e) y[t] = [y[t], e[t]];
							else t = e[w.status], w.then(t, t)
						}
						return this
					}, d.url = ((e || d.url) + "").replace(ze, "").replace(Ye, _e[1] + "//"), d.dataTypes = H.trim(d.dataType || "*").toLowerCase().split(Ze), null == d.crossDomain && (u = tt.exec(d.url.toLowerCase()), d.crossDomain = !(!u || u[1] == _e[1] && u[2] == _e[2] && (u[3] || ("http:" === u[1] ? 80 : 443)) == (_e[3] || ("http:" === _e[1] ? 80 : 443)))), d.data && d.processData && "string" != typeof d.data && (d.data = H.param(d.data, d.traditional)), T(rt, d, n, w), 2 === x) return !1;
				if (c = d.global, d.type = d.type.toUpperCase(), d.hasContent = !Je.test(d.type), c && 0 === H.active++ && H.event.trigger("ajaxStart"), !d.hasContent && (d.data && (d.url += (Ge.test(d.url) ? "&" : "?") + d.data), i = d.url, d.cache === !1)) {
					var E = H.now(),
						S = d.url.replace(et, "$1_=" + E);
					d.url = S + (S === d.url ? (Ge.test(d.url) ? "&" : "?") + "_=" + E : "")
				}(d.data && d.hasContent && d.contentType !== !1 || n.contentType) && w.setRequestHeader("Content-Type", d.contentType), d.ifModified && (i = i || d.url, H.lastModified[i] && w.setRequestHeader("If-Modified-Since", H.lastModified[i]), H.etag[i] && w.setRequestHeader("If-None-Match", H.etag[i])), w.setRequestHeader("Accept", d.dataTypes[0] && d.accepts[d.dataTypes[0]] ? d.accepts[d.dataTypes[0]] + ("*" !== d.dataTypes[0] ? ", */*; q=0.01" : "") : d.accepts["*"]);
				for (f in d.headers) w.setRequestHeader(f, d.headers[f]);
				if (d.beforeSend && (d.beforeSend.call(p, w, d) === !1 || 2 === x)) return w.abort(), !1;
				for (f in {
						success: 1,
						error: 1,
						complete: 1
					}) w[f](d[f]);
				if (s = T(it, d, n, w)) {
					w.readyState = 1, c && h.trigger("ajaxSend", [w, d]), d.async && d.timeout > 0 && (l = setTimeout(function() {
						w.abort("timeout")
					}, d.timeout));
					try {
						x = 1, s.send(v, r)
					} catch (A) {
						status < 2 ? r(-1, A) : H.error(A)
					}
				} else r(-1, "No Transport");
				return w
			},
			param: function(e, n) {
				var r = [],
					i = function(e, t) {
						t = H.isFunction(t) ? t() : t, r[r.length] = encodeURIComponent(e) + "=" + encodeURIComponent(t)
					};
				if (n === t && (n = H.ajaxSettings.traditional), H.isArray(e) || e.jquery && !H.isPlainObject(e)) H.each(e, function() {
					i(this.name, this.value)
				});
				else
					for (var o in e) w(o, e[o], n, i);
				return r.join("&").replace(Re, "+")
			}
		}), H.extend({
			active: 0,
			lastModified: {},
			etag: {}
		});
		var at = H.now(),
			st = /(\=)\?(&|$)|\?\?/i;
		H.ajaxSetup({
			jsonp: "callback",
			jsonpCallback: function() {
				return H.expando + "_" + at++
			}
		}), H.ajaxPrefilter("json jsonp", function(t, n, r) {
			var i = "application/x-www-form-urlencoded" === t.contentType && "string" == typeof t.data;
			if ("jsonp" === t.dataTypes[0] || t.jsonp !== !1 && (st.test(t.url) || i && st.test(t.data))) {
				var o, a = t.jsonpCallback = H.isFunction(t.jsonpCallback) ? t.jsonpCallback() : t.jsonpCallback,
					s = e[a],
					l = t.url,
					u = t.data,
					c = "$1" + a + "$2";
				return t.jsonp !== !1 && (l = l.replace(st, c), t.url === l && (i && (u = u.replace(st, c)), t.data === u && (l += (/\?/.test(l) ? "&" : "?") + t.jsonp + "=" + a))), t.url = l, t.data = u, e[a] = function(e) {
					o = [e]
				}, r.always(function() {
					e[a] = s, o && H.isFunction(s) && e[a](o[0])
				}), t.converters["script json"] = function() {
					return o || H.error(a + " was not called"), o[0]
				}, t.dataTypes[0] = "json", "script"
			}
		}), H.ajaxSetup({
			accepts: {
				script: "text/javascript, application/javascript, application/ecmascript, application/x-ecmascript"
			},
			contents: {
				script: /javascript|ecmascript/
			},
			converters: {
				"text script": function(e) {
					return H.globalEval(e), e
				}
			}
		}), H.ajaxPrefilter("script", function(e) {
			e.cache === t && (e.cache = !1), e.crossDomain && (e.type = "GET", e.global = !1)
		}), H.ajaxTransport("script", function(e) {
			if (e.crossDomain) {
				var n, r = L.head || L.getElementsByTagName("head")[0] || L.documentElement;
				return {
					send: function(i, o) {
						n = L.createElement("script"), n.async = "async", e.scriptCharset && (n.charset = e.scriptCharset), n.src = e.url, n.onload = n.onreadystatechange = function(e, i) {
							(i || !n.readyState || /loaded|complete/.test(n.readyState)) && (n.onload = n.onreadystatechange = null, r && n.parentNode && r.removeChild(n), n = t, i || o(200, "success"))
						}, r.insertBefore(n, r.firstChild)
					},
					abort: function() {
						n && n.onload(0, 1)
					}
				}
			}
		});
		var lt, ut = e.ActiveXObject ? function() {
				for (var e in lt) lt[e](0, 1)
			} : !1,
			ct = 0;
		H.ajaxSettings.xhr = e.ActiveXObject ? function() {
				return !this.isLocal && E() || S()
			} : E,
			function(e) {
				H.extend(H.support, {
					ajax: !!e,
					cors: !!e && "withCredentials" in e
				})
			}(H.ajaxSettings.xhr()), H.support.ajax && H.ajaxTransport(function(n) {
				if (!n.crossDomain || H.support.cors) {
					var r;
					return {
						send: function(i, o) {
							var a, s, l = n.xhr();
							if (n.username ? l.open(n.type, n.url, n.async, n.username, n.password) : l.open(n.type, n.url, n.async), n.xhrFields)
								for (s in n.xhrFields) l[s] = n.xhrFields[s];
							n.mimeType && l.overrideMimeType && l.overrideMimeType(n.mimeType), n.crossDomain || i["X-Requested-With"] || (i["X-Requested-With"] = "XMLHttpRequest");
							try {
								for (s in i) l.setRequestHeader(s, i[s])
							} catch (u) {}
							l.send(n.hasContent && n.data || null), r = function(e, i) {
								var s, u, c, f, d;
								try {
									if (r && (i || 4 === l.readyState))
										if (r = t, a && (l.onreadystatechange = H.noop, ut && delete lt[a]), i) 4 !== l.readyState && l.abort();
										else {
											s = l.status, c = l.getAllResponseHeaders(), f = {}, d = l.responseXML, d && d.documentElement && (f.xml = d), f.text = l.responseText;
											try {
												u = l.statusText
											} catch (p) {
												u = ""
											}
											s || !n.isLocal || n.crossDomain ? 1223 === s && (s = 204) : s = f.text ? 200 : 404
										}
								} catch (h) {
									i || o(-1, h)
								}
								f && o(s, u, f, c)
							}, n.async && 4 !== l.readyState ? (a = ++ct, ut && (lt || (lt = {}, H(e).unload(ut)), lt[a] = r), l.onreadystatechange = r) : r()
						},
						abort: function() {
							r && r(0, 1)
						}
					}
				}
			});
		var ft, dt, pt, ht, mt = {},
			gt = /^(?:toggle|show|hide)$/,
			yt = /^([+\-]=)?([\d+.\-]+)([a-z%]*)$/i,
			vt = [
				["height", "marginTop", "marginBottom", "paddingTop", "paddingBottom"],
				["width", "marginLeft", "marginRight", "paddingLeft", "paddingRight"],
				["opacity"]
			],
			bt = e.webkitRequestAnimationFrame || e.mozRequestAnimationFrame || e.oRequestAnimationFrame;
		H.fn.extend({
			show: function(e, t, n) {
				var r, i;
				if (e || 0 === e) return this.animate(D("show", 3), e, t, n);
				for (var o = 0, a = this.length; a > o; o++) r = this[o], r.style && (i = r.style.display, H._data(r, "olddisplay") || "none" !== i || (i = r.style.display = ""), "" === i && "none" === H.css(r, "display") && H._data(r, "olddisplay", F(r.nodeName)));
				for (o = 0; a > o; o++) r = this[o], r.style && (i = r.style.display, ("" === i || "none" === i) && (r.style.display = H._data(r, "olddisplay") || ""));
				return this
			},
			hide: function(e, t, n) {
				if (e || 0 === e) return this.animate(D("hide", 3), e, t, n);
				for (var r = 0, i = this.length; i > r; r++)
					if (this[r].style) {
						var o = H.css(this[r], "display");
						"none" === o || H._data(this[r], "olddisplay") || H._data(this[r], "olddisplay", o)
					}
				for (r = 0; i > r; r++) this[r].style && (this[r].style.display = "none");
				return this
			},
			_toggle: H.fn.toggle,
			toggle: function(e, t, n) {
				var r = "boolean" == typeof e;
				return H.isFunction(e) && H.isFunction(t) ? this._toggle.apply(this, arguments) : null == e || r ? this.each(function() {
					var t = r ? e : H(this).is(":hidden");
					H(this)[t ? "show" : "hide"]()
				}) : this.animate(D("toggle", 3), e, t, n), this
			},
			fadeTo: function(e, t, n, r) {
				return this.filter(":hidden").css("opacity", 0).show().end().animate({
					opacity: t
				}, e, n, r)
			},
			animate: function(e, t, n, r) {
				var i = H.speed(t, n, r);
				return H.isEmptyObject(e) ? this.each(i.complete, [!1]) : (e = H.extend({}, e), this[i.queue === !1 ? "each" : "queue"](function() {
					i.queue === !1 && H._mark(this);
					var t, n, r, o, a, s, l, u, c, f = H.extend({}, i),
						d = 1 === this.nodeType,
						p = d && H(this).is(":hidden");
					f.animatedProperties = {};
					for (r in e) {
						if (t = H.camelCase(r), r !== t && (e[t] = e[r], delete e[r]), n = e[t], H.isArray(n) ? (f.animatedProperties[t] = n[1], n = e[t] = n[0]) : f.animatedProperties[t] = f.specialEasing && f.specialEasing[t] || f.easing || "swing", "hide" === n && p || "show" === n && !p) return f.complete.call(this);
						!d || "height" !== t && "width" !== t || (f.overflow = [this.style.overflow, this.style.overflowX, this.style.overflowY], "inline" === H.css(this, "display") && "none" === H.css(this, "float") && (H.support.inlineBlockNeedsLayout ? (o = F(this.nodeName), "inline" === o ? this.style.display = "inline-block" : (this.style.display = "inline", this.style.zoom = 1)) : this.style.display = "inline-block"))
					}
					null != f.overflow && (this.style.overflow = "hidden");
					for (r in e) a = new H.fx(this, f, r), n = e[r], gt.test(n) ? a["toggle" === n ? p ? "show" : "hide" : n]() : (s = yt.exec(n), l = a.cur(), s ? (u = parseFloat(s[2]), c = s[3] || (H.cssNumber[r] ? "" : "px"), "px" !== c && (H.style(this, r, (u || 1) + c), l = (u || 1) / a.cur() * l, H.style(this, r, l + c)), s[1] && (u = ("-=" === s[1] ? -1 : 1) * u + l), a.custom(l, u, c)) : a.custom(l, n, ""));
					return !0
				}))
			},
			stop: function(e, t) {
				return e && this.queue([]), this.each(function() {
					var e = H.timers,
						n = e.length;
					for (t || H._unmark(!0, this); n--;) e[n].elem === this && (t && e[n](!0), e.splice(n, 1))
				}), t || this.dequeue(), this
			}
		}), H.each({
			slideDown: D("show", 1),
			slideUp: D("hide", 1),
			slideToggle: D("toggle", 1),
			fadeIn: {
				opacity: "show"
			},
			fadeOut: {
				opacity: "hide"
			},
			fadeToggle: {
				opacity: "toggle"
			}
		}, function(e, t) {
			H.fn[e] = function(e, n, r) {
				return this.animate(t, e, n, r)
			}
		}), H.extend({
			speed: function(e, t, n) {
				var r = e && "object" == typeof e ? H.extend({}, e) : {
					complete: n || !n && t || H.isFunction(e) && e,
					duration: e,
					easing: n && t || t && !H.isFunction(t) && t
				};
				return r.duration = H.fx.off ? 0 : "number" == typeof r.duration ? r.duration : r.duration in H.fx.speeds ? H.fx.speeds[r.duration] : H.fx.speeds._default, r.old = r.complete, r.complete = function(e) {
					H.isFunction(r.old) && r.old.call(this), r.queue !== !1 ? H.dequeue(this) : e !== !1 && H._unmark(this)
				}, r
			},
			easing: {
				linear: function(e, t, n, r) {
					return n + r * e
				},
				swing: function(e, t, n, r) {
					return (-Math.cos(e * Math.PI) / 2 + .5) * r + n
				}
			},
			timers: [],
			fx: function(e, t, n) {
				this.options = t, this.elem = e, this.prop = n, t.orig = t.orig || {}
			}
		}), H.fx.prototype = {
			update: function() {
				this.options.step && this.options.step.call(this.elem, this.now, this), (H.fx.step[this.prop] || H.fx.step._default)(this)
			},
			cur: function() {
				if (null != this.elem[this.prop] && (!this.elem.style || null == this.elem.style[this.prop])) return this.elem[this.prop];
				var e, t = H.css(this.elem, this.prop);
				return isNaN(e = parseFloat(t)) ? t && "auto" !== t ? t : 0 : e
			},
			custom: function(e, t, n) {
				function r(e) {
					return o.step(e)
				}
				var i, o = this,
					a = H.fx;
				this.startTime = ht || A(), this.start = e, this.end = t, this.unit = n || this.unit || (H.cssNumber[this.prop] ? "" : "px"), this.now = this.start, this.pos = this.state = 0, r.elem = this.elem, r() && H.timers.push(r) && !pt && (bt ? (pt = !0, i = function() {
					pt && (bt(i), a.tick())
				}, bt(i)) : pt = setInterval(a.tick, a.interval))
			},
			show: function() {
				this.options.orig[this.prop] = H.style(this.elem, this.prop), this.options.show = !0, this.custom("width" === this.prop || "height" === this.prop ? 1 : 0, this.cur()), H(this.elem).show()
			},
			hide: function() {
				this.options.orig[this.prop] = H.style(this.elem, this.prop), this.options.hide = !0, this.custom(this.cur(), 0)
			},
			step: function(e) {
				var t, n, r = ht || A(),
					i = !0,
					o = this.elem,
					a = this.options;
				if (e || r >= a.duration + this.startTime) {
					this.now = this.end, this.pos = this.state = 1, this.update(), a.animatedProperties[this.prop] = !0;
					for (t in a.animatedProperties) a.animatedProperties[t] !== !0 && (i = !1);
					if (i) {
						if (null == a.overflow || H.support.shrinkWrapBlocks || H.each(["", "X", "Y"], function(e, t) {
								o.style["overflow" + t] = a.overflow[e]
							}), a.hide && H(o).hide(), a.hide || a.show)
							for (var s in a.animatedProperties) H.style(o, s, a.orig[s]);
						a.complete.call(o)
					}
					return !1
				}
				return a.duration == 1 / 0 ? this.now = r : (n = r - this.startTime, this.state = n / a.duration, this.pos = H.easing[a.animatedProperties[this.prop]](this.state, n, 0, 1, a.duration), this.now = this.start + (this.end - this.start) * this.pos), this.update(), !0
			}
		}, H.extend(H.fx, {
			tick: function() {
				for (var e = H.timers, t = 0; t < e.length; ++t) e[t]() || e.splice(t--, 1);
				e.length || H.fx.stop()
			},
			interval: 13,
			stop: function() {
				clearInterval(pt), pt = null
			},
			speeds: {
				slow: 600,
				fast: 200,
				_default: 400
			},
			step: {
				opacity: function(e) {
					H.style(e.elem, "opacity", e.now)
				},
				_default: function(e) {
					e.elem.style && null != e.elem.style[e.prop] ? e.elem.style[e.prop] = ("width" === e.prop || "height" === e.prop ? Math.max(0, e.now) : e.now) + e.unit : e.elem[e.prop] = e.now
				}
			}
		}), H.expr && H.expr.filters && (H.expr.filters.animated = function(e) {
			return H.grep(H.timers, function(t) {
				return e === t.elem
			}).length
		});
		var xt = /^t(?:able|d|h)$/i,
			Tt = /^(?:body|html)$/i;
		"getBoundingClientRect" in L.documentElement ? H.fn.offset = function(e) {
			var t, n = this[0];
			if (e) return this.each(function(t) {
				H.offset.setOffset(this, e, t)
			});
			if (!n || !n.ownerDocument) return null;
			if (n === n.ownerDocument.body) return H.offset.bodyOffset(n);
			try {
				t = n.getBoundingClientRect()
			} catch (r) {}
			var i = n.ownerDocument,
				o = i.documentElement;
			if (!t || !H.contains(o, n)) return t ? {
				top: t.top,
				left: t.left
			} : {
				top: 0,
				left: 0
			};
			var a = i.body,
				s = j(i),
				l = o.clientTop || a.clientTop || 0,
				u = o.clientLeft || a.clientLeft || 0,
				c = s.pageYOffset || H.support.boxModel && o.scrollTop || a.scrollTop,
				f = s.pageXOffset || H.support.boxModel && o.scrollLeft || a.scrollLeft,
				d = t.top + c - l,
				p = t.left + f - u;
			return {
				top: d,
				left: p
			}
		} : H.fn.offset = function(e) {
			var t = this[0];
			if (e) return this.each(function(t) {
				H.offset.setOffset(this, e, t)
			});
			if (!t || !t.ownerDocument) return null;
			if (t === t.ownerDocument.body) return H.offset.bodyOffset(t);
			H.offset.initialize();
			for (var n, r = t.offsetParent, i = t, o = t.ownerDocument, a = o.documentElement, s = o.body, l = o.defaultView, u = l ? l.getComputedStyle(t, null) : t.currentStyle, c = t.offsetTop, f = t.offsetLeft;
				(t = t.parentNode) && t !== s && t !== a && (!H.offset.supportsFixedPosition || "fixed" !== u.position);) n = l ? l.getComputedStyle(t, null) : t.currentStyle, c -= t.scrollTop, f -= t.scrollLeft, t === r && (c += t.offsetTop, f += t.offsetLeft, !H.offset.doesNotAddBorder || H.offset.doesAddBorderForTableAndCells && xt.test(t.nodeName) || (c += parseFloat(n.borderTopWidth) || 0, f += parseFloat(n.borderLeftWidth) || 0), i = r, r = t.offsetParent), H.offset.subtractsBorderForOverflowNotVisible && "visible" !== n.overflow && (c += parseFloat(n.borderTopWidth) || 0, f += parseFloat(n.borderLeftWidth) || 0), u = n;
			return ("relative" === u.position || "static" === u.position) && (c += s.offsetTop, f += s.offsetLeft), H.offset.supportsFixedPosition && "fixed" === u.position && (c += Math.max(a.scrollTop, s.scrollTop), f += Math.max(a.scrollLeft, s.scrollLeft)), {
				top: c,
				left: f
			}
		}, H.offset = {
			initialize: function() {
				var e, t, n, r = L.body,
					i = L.createElement("div"),
					o = parseFloat(H.css(r, "marginTop")) || 0,
					a = "<div style='position:absolute;top:0;left:0;margin:0;border:5px solid #000;padding:0;width:1px;height:1px;'><div></div></div><table style='position:absolute;top:0;left:0;margin:0;border:5px solid #000;padding:0;width:1px;height:1px;' cellpadding='0' cellspacing='0'><tr><td></td></tr></table>";
				H.extend(i.style, {
					position: "absolute",
					top: 0,
					left: 0,
					margin: 0,
					border: 0,
					width: "1px",
					height: "1px",
					visibility: "hidden"
				}), i.innerHTML = a, r.insertBefore(i, r.firstChild), e = i.firstChild, t = e.firstChild, n = e.nextSibling.firstChild.firstChild, this.doesNotAddBorder = 5 !== t.offsetTop, this.doesAddBorderForTableAndCells = 5 === n.offsetTop, t.style.position = "fixed", t.style.top = "20px", this.supportsFixedPosition = 20 === t.offsetTop || 15 === t.offsetTop, t.style.position = t.style.top = "", e.style.overflow = "hidden", e.style.position = "relative", this.subtractsBorderForOverflowNotVisible = -5 === t.offsetTop, this.doesNotIncludeMarginInBodyOffset = r.offsetTop !== o, r.removeChild(i), H.offset.initialize = H.noop
			},
			bodyOffset: function(e) {
				var t = e.offsetTop,
					n = e.offsetLeft;
				return H.offset.initialize(), H.offset.doesNotIncludeMarginInBodyOffset && (t += parseFloat(H.css(e, "marginTop")) || 0, n += parseFloat(H.css(e, "marginLeft")) || 0), {
					top: t,
					left: n
				}
			},
			setOffset: function(e, t, n) {
				var r = H.css(e, "position");
				"static" === r && (e.style.position = "relative");
				var i, o, a = H(e),
					s = a.offset(),
					l = H.css(e, "top"),
					u = H.css(e, "left"),
					c = ("absolute" === r || "fixed" === r) && H.inArray("auto", [l, u]) > -1,
					f = {},
					d = {};
				c ? (d = a.position(), i = d.top, o = d.left) : (i = parseFloat(l) || 0, o = parseFloat(u) || 0), H.isFunction(t) && (t = t.call(e, n, s)), null != t.top && (f.top = t.top - s.top + i), null != t.left && (f.left = t.left - s.left + o), "using" in t ? t.using.call(e, f) : a.css(f)
			}
		}, H.fn.extend({
			position: function() {
				if (!this[0]) return null;
				var e = this[0],
					t = this.offsetParent(),
					n = this.offset(),
					r = Tt.test(t[0].nodeName) ? {
						top: 0,
						left: 0
					} : t.offset();
				return n.top -= parseFloat(H.css(e, "marginTop")) || 0, n.left -= parseFloat(H.css(e, "marginLeft")) || 0, r.top += parseFloat(H.css(t[0], "borderTopWidth")) || 0, r.left += parseFloat(H.css(t[0], "borderLeftWidth")) || 0, {
					top: n.top - r.top,
					left: n.left - r.left
				}
			},
			offsetParent: function() {
				return this.map(function() {
					for (var e = this.offsetParent || L.body; e && !Tt.test(e.nodeName) && "static" === H.css(e, "position");) e = e.offsetParent;
					return e
				})
			}
		}), H.each(["Left", "Top"], function(e, n) {
			var r = "scroll" + n;
			H.fn[r] = function(n) {
				var i, o;
				return n === t ? (i = this[0]) ? (o = j(i), o ? "pageXOffset" in o ? o[e ? "pageYOffset" : "pageXOffset"] : H.support.boxModel && o.document.documentElement[r] || o.document.body[r] : i[r]) : null : this.each(function() {
					o = j(this), o ? o.scrollTo(e ? H(o).scrollLeft() : n, e ? n : H(o).scrollTop()) : this[r] = n
				})
			}
		}), H.each(["Height", "Width"], function(e, n) {
			var r = n.toLowerCase();
			H.fn["inner" + n] = function() {
				var e = this[0];
				return e && e.style ? parseFloat(H.css(e, r, "padding")) : null
			}, H.fn["outer" + n] = function(e) {
				var t = this[0];
				return t && t.style ? parseFloat(H.css(t, r, e ? "margin" : "border")) : null
			}, H.fn[r] = function(e) {
				var i = this[0];
				if (!i) return null == e ? null : this;
				if (H.isFunction(e)) return this.each(function(t) {
					var n = H(this);
					n[r](e.call(this, t, n[r]()))
				});
				if (H.isWindow(i)) {
					var o = i.document.documentElement["client" + n];
					return "CSS1Compat" === i.document.compatMode && o || i.document.body["client" + n] || o
				}
				if (9 === i.nodeType) return Math.max(i.documentElement["client" + n], i.body["scroll" + n], i.documentElement["scroll" + n], i.body["offset" + n], i.documentElement["offset" + n]);
				if (e === t) {
					var a = H.css(i, r),
						s = parseFloat(a);
					return H.isNaN(s) ? a : s
				}
				return this.css(r, "string" == typeof e ? e : e + "px")
			}
		}), e.jQuery = e.$ = H
	}(window), $.noConflict(!0)
});
define("/select-plugin", function(e, t, n) {
	var r = e("/lib/jquery-1.6.2");
	r.fn.sel = function(e) {
		var t = r.extend({}, r.fn.sel.defaults, e),
			n = r(this),
			i = r('<ul id="' + t.containerId + '"></ul>').appendTo(n).addClass("sel");
		if ("number" == typeof t.number)
			for (var o = t.number <= t.hintList.length ? t.number : t.hintList.length, a = 0; o > a; a++) i.append('<li class="' + t.hintList[a].liClass + '"><a val="' + t.hintList[a].val + '" href="#">' + t.hintList[a].liText + "</a></li>");
		else if ("number" != typeof t.number) throw "number选项必须为数字";
		return n.click(function() {
			return i.toggle().show(), r("#customSelectBtn").addClass("selectbtn"), !1
		}), r("body").click(function() {
			i.hide(), r("#customSelectBtn").removeClass("focus"), r("#customSelectBtn").removeClass("selectbtn")
		}), r("#customSelectVal").val("AUTO"), i.find("li a").click(function() {
			var e = r(this).attr("val");
			return t.callback(e), n.toggleClass("focus"), r("#customSelectVal").val(e), n.find(".btn_text").text(r(this).text()), r("#outputLang").html(""), i.hide(), n.removeClass("focus"), i.find(".on").removeClass("on"), r(this).parent().addClass("on"), !1
		}), n
	}, r.fn.sel.defaults = {
		number: 14,
		hintList: [{
			val: "AUTO",
			liText: "自动检测语言",
			liClass: "normal"
		}, {
			val: "SPACE",
			liText: "",
			liClass: "space"
		}, {
			val: "ZH_CN2EN",
			liText: "中文&nbsp; » &nbsp;英语",
			liClass: "isfl topBorder1"
		}, {
			val: "EN2ZH_CN",
			liText: "英语&nbsp; » &nbsp;中文",
			liClass: "isfl rightBorder topBorder"
		}, {
			val: "ZH_CN2JA",
			liText: "中文&nbsp; » &nbsp;日语",
			liClass: "isfl"
		}, {
			val: "JA2ZH_CN",
			liText: "日语&nbsp; » &nbsp;中文",
			liClass: "isfl rightBorder"
		}, {
			val: "ZH_CN2KR",
			liText: "中文&nbsp; » &nbsp;韩语",
			liClass: "isfl"
		}, {
			val: "KR2ZH_CN",
			liText: "韩语&nbsp; » &nbsp;中文",
			liClass: "isfl rightBorder"
		}, {
			val: "ZH_CN2FR",
			liText: "中文&nbsp; » &nbsp;法语",
			liClass: "isfl"
		}, {
			val: "FR2ZH_CN",
			liText: "法语&nbsp; » &nbsp;中文",
			liClass: "isfl rightBorder"
		}, {
			val: "ZH_CN2RU",
			liText: "中文&nbsp; » &nbsp;俄语",
			liClass: "isfl"
		}, {
			val: "RU2ZH_CN",
			liText: "俄语&nbsp; » &nbsp;中文",
			liClass: "isfl rightBorder"
		}, {
			val: "ZH_CN2SP",
			liText: "中文&nbsp; » &nbsp;西班牙语",
			liClass: "isfl"
		}, {
			val: "SP2ZH_CN",
			liText: "西班牙语&nbsp; » &nbsp;中文",
			liClass: "isfl rightBorder"
		}],
		containerId: "customSelectOption",
		init: function() {},
		callback: function() {}
	}
});
define("/atEntrance", function(e, t) {
	var n = e("/lib/jquery-1.6.2");
	t.showAtButton = function(e, t) {
		function i(e) {
			"enhance" === e ? (o.removeClass(o.className), o.addClass("button atEnhanceBtn")) : "disable" === e ? (o.removeClass(o.className), o.addClass("button atDisableBtn"), n("#atBtn").hover(function() {
				n("#atBtnhover").css("display", "none")
			})) : (o.removeClass(o.className), o.addClass("button atNormalBtn"))
		}
		var o = n("#atBtn");
		if ("AUTO" === t || "EN2ZH_CN" === t || "ZH_CN2EN" === t)
			if ("" !== e)
				if ("AUTO" === t) {
					var a = r(e);
					i("en-zh" === a || "zh-en" === a ? "enhance" : "disable")
				} else i("enhance");
		else i("normal");
		else i("disable")
	}, t.timero = function(e, t) {
		var r = function(e, t) {
			this.func = e, this.time = t
		};
		return r.prototype = {
			run: function() {
				var e = this;
				this.clear(), this.timeout = setTimeout(function() {
					n.isFunction(e.func) && e.func()
				}, e.time)
			},
			clear: function() {
				clearTimeout(this.timeout)
			}
		}, new r(e, t)
	};
	var r = function(e) {
		function t(e) {
			var t = e.charCodeAt(0);
			return "　" === e ? !0 : 32 >= t || 255 === t ? !0 : !1
		}

		function r(e) {
			return v[e.charCodeAt(0)] === !0
		}

		function i(e) {
			var t = e.charCodeAt(0);
			return t >= x[0] && t <= x[1] || t >= T[0] && t <= T[1] ? !0 : !1
		}

		function o(e) {
			var t = e.charCodeAt(0);
			return t >= N[0] && t <= N[1] || t >= w[0] && t <= w[1] || t >= C[0] && t <= C[1] || t >= E[0] && t <= E[1] ? !0 : !1
		}

		function a(e) {
			var t = e.charCodeAt(0);
			return t >= S[0] && t <= S[1] || t >= A[0] && t <= A[1] || t >= k[0] && t <= k[1] ? !0 : !1
		}

		function s(e) {
			if ("" === n.trim(e)) return [c, 0];
			for (var s = [], l = [], u = 0, p = 0, h = 0, m = 0, g = 0, v = !1, y = !1, b = 0, x = e.length; x > b; b++) {
				var T = e.charAt(b),
					N = !1,
					w = !1;
				if (t(T) || r(T) ? (N = w = !0, u++) : o(T) ? (v || (v = !0), w = !0) : i(T) ? (y || (y = !0), N = !0) : a(T) ? (N = w = !0, p++) : (N = w = !0, h++), v && (N || b == x - 1) || y && (w || b == x - 1)) {
					var C = b;
					C == x - 1 && (C = b + 1);
					var E = e.substring(g, C);
					if (E.length > 24) return [c, 0];
					E.length > 10 && m++, v ? s[s.length] = E : y && (l[l.length] = E), (v || y) && (g = b), v = !1, y = !1
				}
				N && w && (g = b + 1, v = !1, y = !1)
			}
			var S = l.length,
				A = s.length,
				k = A + p,
				F = k + S;
			0 == k && (A = S, k = S, F = S);
			var D = F + h,
				L = h / D;
			if (L >= .4) return [c, 0];
			var j = A / k;
			if (j >= .8) return m / A >= .4 && m >= 2 ? [c, 0] : [f, D];
			var O = p / k;
			return O >= .8 ? [d, D] : [c, 0]
		}
		for (var l = null, u = null, c = 0, f = 1, d = 2, p = "en", h = "zh", m = [
				["'", "‘", "’"],
				['"', "“", "”"],
				[",", "，"],
				[".", "。"],
				[";", "；"],
				[":", "："],
				["、"],
				["·"],
				["/", "／"],
				["?", "？"],
				["\\", "＼"],
				["|", "｜"],
				["`"],
				["~", "～"],
				["!", "！"],
				["@", "＠"],
				["#", "＃"],
				["$", "＄"],
				["¥", "￥"],
				["%", "％"],
				["^", "＾", "…"],
				["&", "＆"],
				["_"],
				["(", ")"],
				["（", "）"],
				["[", "]"],
				["［", "］"],
				["{", "}"],
				["<", ">"],
				["+", "＋"],
				["-", "－"],
				["*", "＊", "×"],
				["÷"],
				["=", "＝"]
			], g = [
				[10078, 12301, 12303, 12318, 65379],
				[12305, 12309, 12311, 12313, 12315],
				[8250, 12297, 12299]
			], v = {}, y = 0; y < m.length; y++)
			for (var b = 0; b < m[y].length; b++) v[m[y][b].charCodeAt(0)] = !0;
		for (var y = 0; y < g.length; y++)
			for (var b = 0; b < g[y].length; b++) v[g[y][b] - 1] = !0, v[g[y][b]] = !0;
		var x = [48, 57],
			T = [65296, 65305],
			N = [65, 90],
			w = [97, 122],
			C = [65313, 65338],
			E = [65345, 65370],
			S = [19968, 40959],
			A = [13312, 19903],
			k = [63744, 64255],
			F = s(e),
			D = F[0];
		if (D === c) return null;
		if (D == f) l = p, u = h;
		else {
			if (D != d) return null;
			l = h, u = p
		}
		return l + "-" + u
	}
});
define("/sel", function(e, t) {
	var n = e("/lib/jquery-1.6.2");
	e("/select-plugin");
	var r = e("/atEntrance"),
		i = function() {
			n(window).height() < 610 ? setTimeout(function() {
				n(".sponsor-content").addClass("hide-mode-for-ads")
			}, 100) : setTimeout(function() {
				n(".sponsor-content").removeClass("hide-mode-for-ads")
			}, 100)
		};
	i(), n(window).resize(i), t.initSel = function() {
		n("#customSelectBtn").sel({
			callback: function(e) {
				var t = n("#inputText").val();
				r.showAtButton(t, e)
			}
		})
	}
});
define("/ZeroClipboard", function(e, t) {
	! function(e, t) {
		"use strict";
		Array.prototype.map || (Array.prototype.map = function(e, t) {
			var n, r, i;
			if (null == this) throw new TypeError(" this is null or not defined");
			var a = Object(this),
				o = a.length >>> 0;
			if ("function" != typeof e) throw new TypeError(e + " is not a function");
			for (arguments.length > 1 && (n = t), r = new Array(o), i = 0; o > i;) {
				var s, l;
				i in a && (s = a[i], l = e.call(n, s, i, a), r[i] = l), i++
			}
			return r
		});
		var n, r, i, a = e,
			o = a.document,
			s = a.navigator,
			l = a.setTimeout,
			u = a.clearTimeout,
			c = a.setInterval,
			f = a.clearInterval,
			d = a.getComputedStyle,
			p = a.encodeURIComponent,
			h = a.ActiveXObject,
			m = a.Error,
			g = a.Number.parseInt || a.parseInt,
			v = a.Number.parseFloat || a.parseFloat,
			y = a.Number.isNaN || a.isNaN,
			b = a.Date.now,
			x = a.Object.keys,
			w = (a.Object.defineProperty, a.Object.prototype.hasOwnProperty),
			T = a.Array.prototype.slice,
			C = function() {
				var e = function(e) {
					return e
				};
				if ("function" == typeof a.wrap && "function" == typeof a.unwrap) try {
					var t = o.createElement("div"),
						n = a.unwrap(t);
					1 === t.nodeType && n && 1 === n.nodeType && (e = a.unwrap)
				} catch (r) {}
				return e
			}(),
			N = function(e) {
				return T.call(e, 0)
			},
			E = function() {
				var e, n, r, i, a, o, s = N(arguments),
					l = s[0] || {};
				for (e = 1, n = s.length; n > e; e++)
					if (null != (r = s[e]))
						for (i in r) w.call(r, i) && (a = l[i], o = r[i], l !== o && o !== t && (l[i] = o));
				return l
			},
			S = function(e) {
				var t, n, r, i;
				if ("object" != typeof e || null == e || "number" == typeof e.nodeType) t = e;
				else if ("number" == typeof e.length)
					for (t = [], n = 0, r = e.length; r > n; n++) w.call(e, n) && (t[n] = S(e[n]));
				else {
					t = {};
					for (i in e) w.call(e, i) && (t[i] = S(e[i]))
				}
				return t
			},
			A = function(e, t) {
				for (var n = {}, r = 0, i = t.length; i > r; r++) t[r] in e && (n[t[r]] = e[t[r]]);
				return n
			},
			k = function(e, t) {
				var n = {};
				for (var r in e) - 1 === t.indexOf(r) && (n[r] = e[r]);
				return n
			},
			j = function(e) {
				if (e)
					for (var t in e) w.call(e, t) && delete e[t];
				return e
			},
			D = function(e, t) {
				if (e && 1 === e.nodeType && e.ownerDocument && t && (1 === t.nodeType && t.ownerDocument && t.ownerDocument === e.ownerDocument || 9 === t.nodeType && !t.ownerDocument && t === e.ownerDocument))
					do {
						if (e === t) return !0;
						e = e.parentNode
					} while (e);
				return !1
			},
			F = function(e) {
				var t;
				return "string" == typeof e && e && (t = e.split("#")[0].split("?")[0], t = e.slice(0, e.lastIndexOf("/") + 1)), t
			},
			L = function(e) {
				var t, n;
				return "string" == typeof e && e && (n = e.match(/^(?:|[^:@]*@|.+\)@(?=http[s]?|file)|.+?\s+(?: at |@)(?:[^:\(]+ )*[\(]?)((?:http[s]?|file):\/\/[\/]?.+?\/[^:\)]*?)(?::\d+)(?::\d+)?/), n && n[1] ? t = n[1] : (n = e.match(/\)@((?:http[s]?|file):\/\/[\/]?.+?\/[^:\)]*?)(?::\d+)(?::\d+)?/), n && n[1] && (t = n[1]))), t
			},
			O = function() {
				var e, t;
				try {
					throw new m
				} catch (n) {
					t = n
				}
				return t && (e = t.sourceURL || t.fileName || L(t.stack)), e
			},
			I = function() {
				var e, n, r;
				if (o.currentScript && (e = o.currentScript.src)) return e;
				if (n = o.getElementsByTagName("script"), 1 === n.length) return n[0].src || t;
				if ("readyState" in n[0])
					for (r = n.length; r--;)
						if ("interactive" === n[r].readyState && (e = n[r].src)) return e;
				return "loading" === o.readyState && (e = n[n.length - 1].src) ? e : (e = O()) ? e : t
			},
			B = function() {
				var e, n, r, i = o.getElementsByTagName("script");
				for (e = i.length; e--;) {
					if (!(r = i[e].src)) {
						n = null;
						break
					}
					if (r = F(r), null == n) n = r;
					else if (n !== r) {
						n = null;
						break
					}
				}
				return n || t
			},
			M = function() {
				var e = F(I()) || B() || "";
				return e + "ZeroClipboard.swf"
			},
			H = function() {
				return null == e.opener && (!!e.top && e != e.top || !!e.parent && e != e.parent)
			}(),
			_ = {
				bridge: null,
				version: "0.0.0",
				pluginType: "unknown",
				disabled: null,
				outdated: null,
				sandboxed: null,
				unavailable: null,
				degraded: null,
				deactivated: null,
				overdue: null,
				ready: null
			},
			P = "11.0.0",
			R = {},
			$ = {},
			q = null,
			z = 0,
			W = 0,
			X = {
				ready: "Flash communication is established",
				error: {
					"flash-disabled": "Flash is disabled or not installed. May also be attempting to run Flash in a sandboxed iframe, which is impossible.",
					"flash-outdated": "Flash is too outdated to support ZeroClipboard",
					"flash-sandboxed": "Attempting to run Flash in a sandboxed iframe, which is impossible",
					"flash-unavailable": "Flash is unable to communicate bidirectionally with JavaScript",
					"flash-degraded": "Flash is unable to preserve data fidelity when communicating with JavaScript",
					"flash-deactivated": "Flash is too outdated for your browser and/or is configured as click-to-activate.\nThis may also mean that the ZeroClipboard SWF object could not be loaded, so please check your `swfPath` configuration and/or network connectivity.\nMay also be attempting to run Flash in a sandboxed iframe, which is impossible.",
					"flash-overdue": "Flash communication was established but NOT within the acceptable time limit",
					"version-mismatch": "ZeroClipboard JS version number does not match ZeroClipboard SWF version number",
					"clipboard-error": "At least one error was thrown while ZeroClipboard was attempting to inject your data into the clipboard",
					"config-mismatch": "ZeroClipboard configuration does not match Flash's reality",
					"swf-not-found": "The ZeroClipboard SWF object could not be loaded, so please check your `swfPath` configuration and/or network connectivity"
				}
			},
			V = ["flash-unavailable", "flash-degraded", "flash-overdue", "version-mismatch", "config-mismatch", "clipboard-error"],
			Z = ["flash-disabled", "flash-outdated", "flash-sandboxed", "flash-unavailable", "flash-degraded", "flash-deactivated", "flash-overdue"],
			U = new RegExp("^flash-(" + Z.map(function(e) {
				return e.replace(/^flash-/, "")
			}).join("|") + ")$"),
			Y = new RegExp("^flash-(" + Z.slice(1).map(function(e) {
				return e.replace(/^flash-/, "")
			}).join("|") + ")$"),
			J = {
				swfPath: M(),
				trustedDomains: e.location.host ? [e.location.host] : [],
				cacheBust: !0,
				forceEnhancedClipboard: !1,
				flashLoadTimeout: 3e4,
				autoActivate: !0,
				bubbleEvents: !0,
				containerId: "global-zeroclipboard-html-bridge",
				containerClass: "global-zeroclipboard-container",
				swfObjectId: "global-zeroclipboard-flash-bridge",
				hoverClass: "zeroclipboard-is-hover",
				activeClass: "zeroclipboard-is-active",
				forceHandCursor: !1,
				title: null,
				zIndex: 999999999
			},
			K = function(e) {
				if ("object" == typeof e && null !== e)
					for (var t in e)
						if (w.call(e, t))
							if (/^(?:forceHandCursor|title|zIndex|bubbleEvents)$/.test(t)) J[t] = e[t];
							else if (null == _.bridge)
					if ("containerId" === t || "swfObjectId" === t) {
						if (!de(e[t])) throw new Error("The specified `" + t + "` value is not valid as an HTML4 Element ID");
						J[t] = e[t]
					} else J[t] = e[t]; {
						if ("string" != typeof e || !e) return S(J);
						if (w.call(J, e)) return J[e]
					}
			},
			G = function() {
				return We(), {
					browser: A(s, ["userAgent", "platform", "appName"]),
					flash: k(_, ["bridge"]),
					zeroclipboard: {
						version: Ve.version,
						config: Ve.config()
					}
				}
			},
			Q = function() {
				return !!(_.disabled || _.outdated || _.sandboxed || _.unavailable || _.degraded || _.deactivated)
			},
			ee = function(e, r) {
				var i, a, o, s = {};
				if ("string" == typeof e && e) o = e.toLowerCase().split(/\s+/);
				else if ("object" == typeof e && e && "undefined" == typeof r)
					for (i in e) w.call(e, i) && "string" == typeof i && i && "function" == typeof e[i] && Ve.on(i, e[i]);
				if (o && o.length) {
					for (i = 0, a = o.length; a > i; i++) e = o[i].replace(/^on/, ""), s[e] = !0, R[e] || (R[e] = []), R[e].push(r);
					if (s.ready && _.ready && Ve.emit({
							type: "ready"
						}), s.error) {
						for (i = 0, a = Z.length; a > i; i++)
							if (_[Z[i].replace(/^flash-/, "")] === !0) {
								Ve.emit({
									type: "error",
									name: Z[i]
								});
								break
							}
						n !== t && Ve.version !== n && Ve.emit({
							type: "error",
							name: "version-mismatch",
							jsVersion: Ve.version,
							swfVersion: n
						})
					}
				}
				return Ve
			},
			te = function(e, t) {
				var n, r, i, a, o;
				if (0 === arguments.length) a = x(R);
				else if ("string" == typeof e && e) a = e.split(/\s+/);
				else if ("object" == typeof e && e && "undefined" == typeof t)
					for (n in e) w.call(e, n) && "string" == typeof n && n && "function" == typeof e[n] && Ve.off(n, e[n]);
				if (a && a.length)
					for (n = 0, r = a.length; r > n; n++)
						if (e = a[n].toLowerCase().replace(/^on/, ""), o = R[e], o && o.length)
							if (t)
								for (i = o.indexOf(t); - 1 !== i;) o.splice(i, 1), i = o.indexOf(t, i);
							else o.length = 0;
				return Ve
			},
			ne = function(e) {
				var t;
				return t = "string" == typeof e && e ? S(R[e]) || null : S(R)
			},
			re = function(e) {
				var t, n, r;
				return e = pe(e), e && !xe(e) ? "ready" === e.type && _.overdue === !0 ? Ve.emit({
					type: "error",
					name: "flash-overdue"
				}) : (t = E({}, e), ye.call(this, t), "copy" === e.type && (r = ke($), n = r.data, q = r.formatMap), n) : void 0
			},
			ie = function() {
				var e = _.sandboxed;
				if (We(), "boolean" != typeof _.ready && (_.ready = !1), _.sandboxed !== e && _.sandboxed === !0) _.ready = !1, Ve.emit({
					type: "error",
					name: "flash-sandboxed"
				});
				else if (!Ve.isFlashUnusable() && null === _.bridge) {
					var t = J.flashLoadTimeout;
					"number" == typeof t && t >= 0 && (z = l(function() {
						"boolean" != typeof _.deactivated && (_.deactivated = !0), _.deactivated === !0 && Ve.emit({
							type: "error",
							name: "flash-deactivated"
						})
					}, t)), _.overdue = !1, Se()
				}
			},
			ae = function() {
				Ve.clearData(), Ve.blur(), Ve.emit("destroy"), Ae(), Ve.off()
			},
			oe = function(e, t) {
				var n;
				if ("object" == typeof e && e && "undefined" == typeof t) n = e, Ve.clearData();
				else {
					if ("string" != typeof e || !e) return;
					n = {}, n[e] = t
				}
				for (var r in n) "string" == typeof r && r && w.call(n, r) && "string" == typeof n[r] && n[r] && ($[r] = n[r])
			},
			se = function(e) {
				"undefined" == typeof e ? (j($), q = null) : "string" == typeof e && w.call($, e) && delete $[e]
			},
			le = function(e) {
				return "undefined" == typeof e ? S($) : "string" == typeof e && w.call($, e) ? $[e] : void 0
			},
			ue = function(e) {
				if (e && 1 === e.nodeType) {
					r && (Me(r, J.activeClass), r !== e && Me(r, J.hoverClass)), r = e, Be(e, J.hoverClass);
					var t = e.getAttribute("title") || J.title;
					if ("string" == typeof t && t) {
						var n = Ee(_.bridge);
						n && n.setAttribute("title", t)
					}
					var i = J.forceHandCursor === !0 || "pointer" === He(e, "cursor");
					qe(i), $e()
				}
			},
			ce = function() {
				var e = Ee(_.bridge);
				e && (e.removeAttribute("title"), e.style.left = "0px", e.style.top = "-9999px", e.style.width = "1px", e.style.height = "1px"), r && (Me(r, J.hoverClass), Me(r, J.activeClass), r = null)
			},
			fe = function() {
				return r || null
			},
			de = function(e) {
				return "string" == typeof e && e && /^[A-Za-z][A-Za-z0-9_:\-\.]*$/.test(e)
			},
			pe = function(e) {
				var t;
				if ("string" == typeof e && e ? (t = e, e = {}) : "object" == typeof e && e && "string" == typeof e.type && e.type && (t = e.type), t) {
					t = t.toLowerCase(), !e.target && (/^(copy|aftercopy|_click)$/.test(t) || "error" === t && "clipboard-error" === e.name) && (e.target = i), E(e, {
						type: t,
						target: e.target || r || null,
						relatedTarget: e.relatedTarget || null,
						currentTarget: _ && _.bridge || null,
						timeStamp: e.timeStamp || b() || null
					});
					var n = X[e.type];
					return "error" === e.type && e.name && n && (n = n[e.name]), n && (e.message = n), "ready" === e.type && E(e, {
						target: null,
						version: _.version
					}), "error" === e.type && (U.test(e.name) && E(e, {
						target: null,
						minimumVersion: P
					}), Y.test(e.name) && E(e, {
						version: _.version
					})), "copy" === e.type && (e.clipboardData = {
						setData: Ve.setData,
						clearData: Ve.clearData
					}), "aftercopy" === e.type && (e = je(e, q)), e.target && !e.relatedTarget && (e.relatedTarget = he(e.target)), me(e)
				}
			},
			he = function(e) {
				var t = e && e.getAttribute && e.getAttribute("data-clipboard-target");
				return t ? o.getElementById(t) : null
			},
			me = function(e) {
				if (e && /^_(?:click|mouse(?:over|out|down|up|move))$/.test(e.type)) {
					var n = e.target,
						r = "_mouseover" === e.type && e.relatedTarget ? e.relatedTarget : t,
						i = "_mouseout" === e.type && e.relatedTarget ? e.relatedTarget : t,
						s = _e(n),
						l = a.screenLeft || a.screenX || 0,
						u = a.screenTop || a.screenY || 0,
						c = o.body.scrollLeft + o.documentElement.scrollLeft,
						f = o.body.scrollTop + o.documentElement.scrollTop,
						d = s.left + ("number" == typeof e._stageX ? e._stageX : 0),
						p = s.top + ("number" == typeof e._stageY ? e._stageY : 0),
						h = d - c,
						m = p - f,
						g = l + h,
						v = u + m,
						y = "number" == typeof e.movementX ? e.movementX : 0,
						b = "number" == typeof e.movementY ? e.movementY : 0;
					delete e._stageX, delete e._stageY, E(e, {
						srcElement: n,
						fromElement: r,
						toElement: i,
						screenX: g,
						screenY: v,
						pageX: d,
						pageY: p,
						clientX: h,
						clientY: m,
						x: h,
						y: m,
						movementX: y,
						movementY: b,
						offsetX: 0,
						offsetY: 0,
						layerX: 0,
						layerY: 0
					})
				}
				return e
			},
			ge = function(e) {
				var t = e && "string" == typeof e.type && e.type || "";
				return !/^(?:(?:before)?copy|destroy)$/.test(t)
			},
			ve = function(e, t, n, r) {
				r ? l(function() {
					e.apply(t, n)
				}, 0) : e.apply(t, n)
			},
			ye = function(e) {
				if ("object" == typeof e && e && e.type) {
					var t = ge(e),
						n = R["*"] || [],
						r = R[e.type] || [],
						i = n.concat(r);
					if (i && i.length) {
						var o, s, l, u, c, f = this;
						for (o = 0, s = i.length; s > o; o++) l = i[o], u = f, "string" == typeof l && "function" == typeof a[l] && (l = a[l]), "object" == typeof l && l && "function" == typeof l.handleEvent && (u = l, l = l.handleEvent), "function" == typeof l && (c = E({}, e), ve(l, u, [c], t))
					}
					return this
				}
			},
			be = function(e) {
				var t = null;
				return (H === !1 || e && "error" === e.type && e.name && -1 !== V.indexOf(e.name)) && (t = !1), t
			},
			xe = function(e) {
				var t = e.target || r || null,
					a = "swf" === e._source;
				switch (delete e._source, e.type) {
					case "error":
						var o = "flash-sandboxed" === e.name || be(e);
						"boolean" == typeof o && (_.sandboxed = o), -1 !== Z.indexOf(e.name) ? E(_, {
							disabled: "flash-disabled" === e.name,
							outdated: "flash-outdated" === e.name,
							unavailable: "flash-unavailable" === e.name,
							degraded: "flash-degraded" === e.name,
							deactivated: "flash-deactivated" === e.name,
							overdue: "flash-overdue" === e.name,
							ready: !1
						}) : "version-mismatch" === e.name && (n = e.swfVersion, E(_, {
							disabled: !1,
							outdated: !1,
							unavailable: !1,
							degraded: !1,
							deactivated: !1,
							overdue: !1,
							ready: !1
						})), Re();
						break;
					case "ready":
						n = e.swfVersion;
						var s = _.deactivated === !0;
						E(_, {
							disabled: !1,
							outdated: !1,
							sandboxed: !1,
							unavailable: !1,
							degraded: !1,
							deactivated: !1,
							overdue: s,
							ready: !s
						}), Re();
						break;
					case "beforecopy":
						i = t;
						break;
					case "copy":
						var l, u, c = e.relatedTarget;
						!$["text/html"] && !$["text/plain"] && c && (u = c.value || c.outerHTML || c.innerHTML) && (l = c.value || c.textContent || c.innerText) ? (e.clipboardData.clearData(), e.clipboardData.setData("text/plain", l), u !== l && e.clipboardData.setData("text/html", u)) : !$["text/plain"] && e.target && (l = e.target.getAttribute("data-clipboard-text")) && (e.clipboardData.clearData(), e.clipboardData.setData("text/plain", l));
						break;
					case "aftercopy":
						we(e), Ve.clearData(), t && t !== Ie() && t.focus && t.focus();
						break;
					case "_mouseover":
						Ve.focus(t), J.bubbleEvents === !0 && a && (t && t !== e.relatedTarget && !D(e.relatedTarget, t) && Te(E({}, e, {
							type: "mouseenter",
							bubbles: !1,
							cancelable: !1
						})), Te(E({}, e, {
							type: "mouseover"
						})));
						break;
					case "_mouseout":
						Ve.blur(), J.bubbleEvents === !0 && a && (t && t !== e.relatedTarget && !D(e.relatedTarget, t) && Te(E({}, e, {
							type: "mouseleave",
							bubbles: !1,
							cancelable: !1
						})), Te(E({}, e, {
							type: "mouseout"
						})));
						break;
					case "_mousedown":
						Be(t, J.activeClass), J.bubbleEvents === !0 && a && Te(E({}, e, {
							type: e.type.slice(1)
						}));
						break;
					case "_mouseup":
						Me(t, J.activeClass), J.bubbleEvents === !0 && a && Te(E({}, e, {
							type: e.type.slice(1)
						}));
						break;
					case "_click":
						i = null, J.bubbleEvents === !0 && a && Te(E({}, e, {
							type: e.type.slice(1)
						}));
						break;
					case "_mousemove":
						J.bubbleEvents === !0 && a && Te(E({}, e, {
							type: e.type.slice(1)
						}))
				}
				return /^_(?:click|mouse(?:over|out|down|up|move))$/.test(e.type) ? !0 : void 0
			},
			we = function(e) {
				if (e.errors && e.errors.length > 0) {
					var t = S(e);
					E(t, {
						type: "error",
						name: "clipboard-error"
					}), delete t.success, l(function() {
						Ve.emit(t)
					}, 0)
				}
			},
			Te = function(e) {
				if (e && "string" == typeof e.type && e) {
					var t, n = e.target || null,
						r = n && n.ownerDocument || o,
						i = {
							view: r.defaultView || a,
							canBubble: !0,
							cancelable: !0,
							detail: "click" === e.type ? 1 : 0,
							button: "number" == typeof e.which ? e.which - 1 : "number" == typeof e.button ? e.button : r.createEvent ? 0 : 1
						},
						s = E(i, e);
					n && r.createEvent && n.dispatchEvent && (s = [s.type, s.canBubble, s.cancelable, s.view, s.detail, s.screenX, s.screenY, s.clientX, s.clientY, s.ctrlKey, s.altKey, s.shiftKey, s.metaKey, s.button, s.relatedTarget], t = r.createEvent("MouseEvents"), t.initMouseEvent && (t.initMouseEvent.apply(t, s), t._source = "js", n.dispatchEvent(t)))
				}
			},
			Ce = function() {
				var e = J.flashLoadTimeout;
				if ("number" == typeof e && e >= 0) {
					var t = Math.min(1e3, e / 10),
						n = J.swfObjectId + "_fallbackContent";
					W = c(function() {
						var e = o.getElementById(n);
						Pe(e) && (Re(), _.deactivated = null, Ve.emit({
							type: "error",
							name: "swf-not-found"
						}))
					}, t)
				}
			},
			Ne = function() {
				var e = o.createElement("div");
				return e.id = J.containerId, e.className = J.containerClass, e.style.position = "absolute", e.style.left = "0px", e.style.top = "-9999px", e.style.width = "1px", e.style.height = "1px", e.style.zIndex = "" + ze(J.zIndex), e
			},
			Ee = function(e) {
				for (var t = e && e.parentNode; t && "OBJECT" === t.nodeName && t.parentNode;) t = t.parentNode;
				return t || null
			},
			Se = function() {
				var e, t = _.bridge,
					n = Ee(t);
				if (!t) {
					var r = Oe(a.location.host, J),
						i = "never" === r ? "none" : "all",
						s = Fe(E({
							jsVersion: Ve.version
						}, J)),
						l = J.swfPath + De(J.swfPath, J);
					n = Ne();
					var u = o.createElement("div");
					n.appendChild(u), o.body.appendChild(n);
					var c = o.createElement("div"),
						f = "activex" === _.pluginType;
					c.innerHTML = '<object id="' + J.swfObjectId + '" name="' + J.swfObjectId + '" width="100%" height="100%" ' + (f ? 'classid="clsid:d27cdb6e-ae6d-11cf-96b8-444553540000"' : 'type="application/x-shockwave-flash" data="' + l + '"') + ">" + (f ? '<param name="movie" value="' + l + '"/>' : "") + '<param name="allowScriptAccess" value="' + r + '"/><param name="allowNetworking" value="' + i + '"/><param name="menu" value="false"/><param name="wmode" value="transparent"/><param name="flashvars" value="' + s + '"/><div id="' + J.swfObjectId + '_fallbackContent">&nbsp;</div></object>', t = c.firstChild, c = null, C(t).ZeroClipboard = Ve, n.replaceChild(t, u), Ce()
				}
				return t || (t = o[J.swfObjectId], t && (e = t.length) && (t = t[e - 1]), !t && n && (t = n.firstChild)), _.bridge = t || null, t
			},
			Ae = function() {
				var e = _.bridge;
				if (e) {
					var r = Ee(e);
					r && ("activex" === _.pluginType && "readyState" in e ? (e.style.display = "none", function i() {
						if (4 === e.readyState) {
							for (var t in e) "function" == typeof e[t] && (e[t] = null);
							e.parentNode && e.parentNode.removeChild(e), r.parentNode && r.parentNode.removeChild(r)
						} else l(i, 10)
					}()) : (e.parentNode && e.parentNode.removeChild(e), r.parentNode && r.parentNode.removeChild(r))), Re(), _.ready = null, _.bridge = null, _.deactivated = null, n = t
				}
			},
			ke = function(e) {
				var t = {},
					n = {};
				if ("object" == typeof e && e) {
					for (var r in e)
						if (r && w.call(e, r) && "string" == typeof e[r] && e[r]) switch (r.toLowerCase()) {
							case "text/plain":
							case "text":
							case "air:text":
							case "flash:text":
								t.text = e[r], n.text = r;
								break;
							case "text/html":
							case "html":
							case "air:html":
							case "flash:html":
								t.html = e[r], n.html = r;
								break;
							case "application/rtf":
							case "text/rtf":
							case "rtf":
							case "richtext":
							case "air:rtf":
							case "flash:rtf":
								t.rtf = e[r], n.rtf = r
						}
						return {
							data: t,
							formatMap: n
						}
				}
			},
			je = function(e, t) {
				if ("object" != typeof e || !e || "object" != typeof t || !t) return e;
				var n = {};
				for (var r in e)
					if (w.call(e, r))
						if ("errors" === r) {
							n[r] = e[r] ? e[r].slice() : [];
							for (var i = 0, a = n[r].length; a > i; i++) n[r][i].format = t[n[r][i].format]
						} else if ("success" !== r && "data" !== r) n[r] = e[r];
				else {
					n[r] = {};
					var o = e[r];
					for (var s in o) s && w.call(o, s) && w.call(t, s) && (n[r][t[s]] = o[s])
				}
				return n
			},
			De = function(e, t) {
				var n = null == t || t && t.cacheBust === !0;
				return n ? (-1 === e.indexOf("?") ? "?" : "&") + "noCache=" + b() : ""
			},
			Fe = function(e) {
				var t, n, r, i, o = "",
					s = [];
				if (e.trustedDomains && ("string" == typeof e.trustedDomains ? i = [e.trustedDomains] : "object" == typeof e.trustedDomains && "length" in e.trustedDomains && (i = e.trustedDomains)), i && i.length)
					for (t = 0, n = i.length; n > t; t++)
						if (w.call(i, t) && i[t] && "string" == typeof i[t]) {
							if (r = Le(i[t]), !r) continue;
							if ("*" === r) {
								s.length = 0, s.push(r);
								break
							}
							s.push.apply(s, [r, "//" + r, a.location.protocol + "//" + r])
						}
				return s.length && (o += "trustedOrigins=" + p(s.join(","))), e.forceEnhancedClipboard === !0 && (o += (o ? "&" : "") + "forceEnhancedClipboard=true"), "string" == typeof e.swfObjectId && e.swfObjectId && (o += (o ? "&" : "") + "swfObjectId=" + p(e.swfObjectId)), "string" == typeof e.jsVersion && e.jsVersion && (o += (o ? "&" : "") + "jsVersion=" + p(e.jsVersion)), o
			},
			Le = function(e) {
				if (null == e || "" === e) return null;
				if (e = e.replace(/^\s+|\s+$/g, ""), "" === e) return null;
				var t = e.indexOf("//");
				e = -1 === t ? e : e.slice(t + 2);
				var n = e.indexOf("/");
				return e = -1 === n ? e : -1 === t || 0 === n ? null : e.slice(0, n), e && ".swf" === e.slice(-4).toLowerCase() ? null : e || null
			},
			Oe = function() {
				var e = function(e) {
					var t, n, r, i = [];
					if ("string" == typeof e && (e = [e]), "object" != typeof e || !e || "number" != typeof e.length) return i;
					for (t = 0, n = e.length; n > t; t++)
						if (w.call(e, t) && (r = Le(e[t]))) {
							if ("*" === r) {
								i.length = 0, i.push("*");
								break
							} - 1 === i.indexOf(r) && i.push(r)
						}
					return i
				};
				return function(t, n) {
					var r = Le(n.swfPath);
					null === r && (r = t);
					var i = e(n.trustedDomains),
						a = i.length;
					if (a > 0) {
						if (1 === a && "*" === i[0]) return "always";
						if (-1 !== i.indexOf(t)) return 1 === a && t === r ? "sameDomain" : "always"
					}
					return "never"
				}
			}(),
			Ie = function() {
				try {
					return o.activeElement
				} catch (e) {
					return null
				}
			},
			Be = function(e, t) {
				var n, r, i, a = [];
				if ("string" == typeof t && t && (a = t.split(/\s+/)), e && 1 === e.nodeType && a.length > 0)
					if (e.classList)
						for (n = 0, r = a.length; r > n; n++) e.classList.add(a[n]);
					else if (e.hasOwnProperty("className")) {
					for (i = " " + e.className + " ", n = 0, r = a.length; r > n; n++) - 1 === i.indexOf(" " + a[n] + " ") && (i += a[n] + " ");
					e.className = i.replace(/^\s+|\s+$/g, "")
				}
				return e
			},
			Me = function(e, t) {
				var n, r, i, a = [];
				if ("string" == typeof t && t && (a = t.split(/\s+/)), e && 1 === e.nodeType && a.length > 0)
					if (e.classList && e.classList.length > 0)
						for (n = 0, r = a.length; r > n; n++) e.classList.remove(a[n]);
					else if (e.className) {
					for (i = (" " + e.className + " ").replace(/[\r\n\t]/g, " "), n = 0, r = a.length; r > n; n++) i = i.replace(" " + a[n] + " ", " ");
					e.className = i.replace(/^\s+|\s+$/g, "")
				}
				return e
			},
			He = function(e, t) {
				var n = d(e, null).getPropertyValue(t);
				return "cursor" !== t || n && "auto" !== n || "A" !== e.nodeName ? n : "pointer"
			},
			_e = function(e) {
				var t = {
					left: 0,
					top: 0,
					width: 0,
					height: 0
				};
				if (e.getBoundingClientRect) {
					var n = e.getBoundingClientRect(),
						r = a.pageXOffset,
						i = a.pageYOffset,
						s = o.documentElement.clientLeft || 0,
						l = o.documentElement.clientTop || 0,
						u = 0,
						c = 0;
					if ("relative" === He(o.body, "position")) {
						var f = o.body.getBoundingClientRect(),
							d = o.documentElement.getBoundingClientRect();
						u = f.left - d.left || 0, c = f.top - d.top || 0
					}
					t.left = n.left + r - s - u, t.top = n.top + i - l - c, t.width = "width" in n ? n.width : n.right - n.left, t.height = "height" in n ? n.height : n.bottom - n.top
				}
				return t
			},
			Pe = function(e) {
				if (!e) return !1;
				var t = d(e, null),
					n = v(t.height) > 0,
					r = v(t.width) > 0,
					i = v(t.top) >= 0,
					a = v(t.left) >= 0,
					o = n && r && i && a,
					s = o ? null : _e(e),
					l = "none" !== t.display && "collapse" !== t.visibility && (o || !!s && (n || s.height > 0) && (r || s.width > 0) && (i || s.top >= 0) && (a || s.left >= 0));
				return l
			},
			Re = function() {
				u(z), z = 0, f(W), W = 0
			},
			$e = function() {
				var e;
				if (r && (e = Ee(_.bridge))) {
					var t = _e(r);
					E(e.style, {
						width: t.width + "px",
						height: t.height + "px",
						top: t.top + "px",
						left: t.left + "px",
						zIndex: "" + ze(J.zIndex)
					})
				}
			},
			qe = function(e) {
				_.ready === !0 && (_.bridge && "function" == typeof _.bridge.setHandCursor ? _.bridge.setHandCursor(e) : _.ready = !1)
			},
			ze = function(e) {
				if (/^(?:auto|inherit)$/.test(e)) return e;
				var t;
				return "number" != typeof e || y(e) ? "string" == typeof e && (t = ze(g(e, 10))) : t = e, "number" == typeof t ? t : "auto"
			},
			We = function(t) {
				var n, r, i, a = _.sandboxed,
					o = null;
				if (t = t === !0, H === !1) o = !1;
				else {
					try {
						r = e.frameElement || null
					} catch (s) {
						i = {
							name: s.name,
							message: s.message
						}
					}
					if (r && 1 === r.nodeType && "IFRAME" === r.nodeName) try {
						o = r.hasAttribute("sandbox")
					} catch (s) {
						o = null
					} else {
						try {
							n = document.domain || null
						} catch (s) {
							n = null
						}(null === n || i && "SecurityError" === i.name && /(^|[\s\(\[@])sandbox(es|ed|ing|[\s\.,!\)\]@]|$)/.test(i.message.toLowerCase())) && (o = !0)
					}
				}
				return _.sandboxed = o, a === o || t || Xe(h), o
			},
			Xe = function(e) {
				function t(e) {
					var t = e.match(/[\d]+/g);
					return t.length = 3, t.join(".")
				}

				function n(e) {
					return !!e && (e = e.toLowerCase()) && (/^(pepflashplayer\.dll|libpepflashplayer\.so|pepperflashplayer\.plugin)$/.test(e) || "chrome.plugin" === e.slice(-13))
				}

				function r(e) {
					e && (l = !0, e.version && (f = t(e.version)), !f && e.description && (f = t(e.description)), e.filename && (c = n(e.filename)))
				}
				var i, a, o, l = !1,
					u = !1,
					c = !1,
					f = "";
				if (s.plugins && s.plugins.length) i = s.plugins["Shockwave Flash"], r(i), s.plugins["Shockwave Flash 2.0"] && (l = !0, f = "2.0.0.11");
				else if (s.mimeTypes && s.mimeTypes.length) o = s.mimeTypes["application/x-shockwave-flash"], i = o && o.enabledPlugin, r(i);
				else if ("undefined" != typeof e) {
					u = !0;
					try {
						a = new e("ShockwaveFlash.ShockwaveFlash.7"), l = !0, f = t(a.GetVariable("$version"))
					} catch (d) {
						try {
							a = new e("ShockwaveFlash.ShockwaveFlash.6"), l = !0, f = "6.0.21"
						} catch (p) {
							try {
								a = new e("ShockwaveFlash.ShockwaveFlash"), l = !0, f = t(a.GetVariable("$version"))
							} catch (h) {
								u = !1
							}
						}
					}
				}
				_.disabled = l !== !0, _.outdated = f && v(f) < v(P), _.version = f || "0.0.0", _.pluginType = c ? "pepper" : u ? "activex" : l ? "netscape" : "unknown"
			};
		Xe(h), We(!0);
		var Ve = function() {
			return this instanceof Ve ? void("function" == typeof Ve._createClient && Ve._createClient.apply(this, N(arguments))) : new Ve
		};
		Ve.version = "2.2.0", Ve.config = function() {
			return K.apply(this, N(arguments))
		}, Ve.state = function() {
			return G.apply(this, N(arguments))
		}, Ve.isFlashUnusable = function() {
			return Q.apply(this, N(arguments))
		}, Ve.on = function() {
			return ee.apply(this, N(arguments))
		}, Ve.off = function() {
			return te.apply(this, N(arguments))
		}, Ve.handlers = function() {
			return ne.apply(this, N(arguments))
		}, Ve.emit = function() {
			return re.apply(this, N(arguments))
		}, Ve.create = function() {
			return ie.apply(this, N(arguments))
		}, Ve.destroy = function() {
			return ae.apply(this, N(arguments))
		}, Ve.setData = function() {
			return oe.apply(this, N(arguments))
		}, Ve.clearData = function() {
			return se.apply(this, N(arguments))
		}, Ve.getData = function() {
			return le.apply(this, N(arguments))
		}, Ve.focus = Ve.activate = function() {
			return ue.apply(this, N(arguments))
		}, Ve.blur = Ve.deactivate = function() {
			return ce.apply(this, N(arguments))
		}, Ve.activeElement = function() {
			return fe.apply(this, N(arguments))
		};
		var Ze = 0,
			Ue = {},
			Ye = 0,
			Je = {},
			Ke = {};
		E(J, {
			autoActivate: !0
		});
		var Ge = function(e) {
				var t = this;
				t.id = "" + Ze++, Ue[t.id] = {
					instance: t,
					elements: [],
					handlers: {}
				}, e && t.clip(e), Ve.on("*", function(e) {
					return t.emit(e)
				}), Ve.on("destroy", function() {
					t.destroy()
				}), Ve.create()
			},
			Qe = function(e, r) {
				var i, a, o, s = {},
					l = Ue[this.id],
					u = l && l.handlers;
				if (!l) throw new Error("Attempted to add new listener(s) to a destroyed ZeroClipboard client instance");
				if ("string" == typeof e && e) o = e.toLowerCase().split(/\s+/);
				else if ("object" == typeof e && e && "undefined" == typeof r)
					for (i in e) w.call(e, i) && "string" == typeof i && i && "function" == typeof e[i] && this.on(i, e[i]);
				if (o && o.length) {
					for (i = 0, a = o.length; a > i; i++) e = o[i].replace(/^on/, ""), s[e] = !0, u[e] || (u[e] = []), u[e].push(r);
					if (s.ready && _.ready && this.emit({
							type: "ready",
							client: this
						}), s.error) {
						for (i = 0, a = Z.length; a > i; i++)
							if (_[Z[i].replace(/^flash-/, "")]) {
								this.emit({
									type: "error",
									name: Z[i],
									client: this
								});
								break
							}
						n !== t && Ve.version !== n && this.emit({
							type: "error",
							name: "version-mismatch",
							jsVersion: Ve.version,
							swfVersion: n
						})
					}
				}
				return this
			},
			et = function(e, t) {
				var n, r, i, a, o, s = Ue[this.id],
					l = s && s.handlers;
				if (!l) return this;
				if (0 === arguments.length) a = x(l);
				else if ("string" == typeof e && e) a = e.split(/\s+/);
				else if ("object" == typeof e && e && "undefined" == typeof t)
					for (n in e) w.call(e, n) && "string" == typeof n && n && "function" == typeof e[n] && this.off(n, e[n]);
				if (a && a.length)
					for (n = 0, r = a.length; r > n; n++)
						if (e = a[n].toLowerCase().replace(/^on/, ""), o = l[e], o && o.length)
							if (t)
								for (i = o.indexOf(t); - 1 !== i;) o.splice(i, 1), i = o.indexOf(t, i);
							else o.length = 0;
				return this
			},
			tt = function(e) {
				var t = null,
					n = Ue[this.id] && Ue[this.id].handlers;
				return n && (t = "string" == typeof e && e ? n[e] ? n[e].slice(0) : [] : S(n)), t
			},
			nt = function(e) {
				if (st.call(this, e)) {
					"object" == typeof e && e && "string" == typeof e.type && e.type && (e = E({}, e));
					var t = E({}, pe(e), {
						client: this
					});
					lt.call(this, t)
				}
				return this
			},
			rt = function(e) {
				if (!Ue[this.id]) throw new Error("Attempted to clip element(s) to a destroyed ZeroClipboard client instance");
				e = ut(e);
				for (var t = 0; t < e.length; t++)
					if (w.call(e, t) && e[t] && 1 === e[t].nodeType) {
						e[t].zcClippingId ? -1 === Je[e[t].zcClippingId].indexOf(this.id) && Je[e[t].zcClippingId].push(this.id) : (e[t].zcClippingId = "zcClippingId_" + Ye++, Je[e[t].zcClippingId] = [this.id], J.autoActivate === !0 && ct(e[t]));
						var n = Ue[this.id] && Ue[this.id].elements; - 1 === n.indexOf(e[t]) && n.push(e[t])
					}
				return this
			},
			it = function(e) {
				var t = Ue[this.id];
				if (!t) return this;
				var n, r = t.elements;
				e = "undefined" == typeof e ? r.slice(0) : ut(e);
				for (var i = e.length; i--;)
					if (w.call(e, i) && e[i] && 1 === e[i].nodeType) {
						for (n = 0; - 1 !== (n = r.indexOf(e[i], n));) r.splice(n, 1);
						var a = Je[e[i].zcClippingId];
						if (a) {
							for (n = 0; - 1 !== (n = a.indexOf(this.id, n));) a.splice(n, 1);
							0 === a.length && (J.autoActivate === !0 && ft(e[i]), delete e[i].zcClippingId)
						}
					}
				return this
			},
			at = function() {
				var e = Ue[this.id];
				return e && e.elements ? e.elements.slice(0) : []
			},
			ot = function() {
				Ue[this.id] && (this.unclip(), this.off(), delete Ue[this.id])
			},
			st = function(e) {
				if (!e || !e.type) return !1;
				if (e.client && e.client !== this) return !1;
				var t = Ue[this.id],
					n = t && t.elements,
					r = !!n && n.length > 0,
					i = !e.target || r && -1 !== n.indexOf(e.target),
					a = e.relatedTarget && r && -1 !== n.indexOf(e.relatedTarget),
					o = e.client && e.client === this;
				return t && (i || a || o) ? !0 : !1
			},
			lt = function(e) {
				var t = Ue[this.id];
				if ("object" == typeof e && e && e.type && t) {
					var n = ge(e),
						r = t && t.handlers["*"] || [],
						i = t && t.handlers[e.type] || [],
						o = r.concat(i);
					if (o && o.length) {
						var s, l, u, c, f, d = this;
						for (s = 0, l = o.length; l > s; s++) u = o[s], c = d, "string" == typeof u && "function" == typeof a[u] && (u = a[u]), "object" == typeof u && u && "function" == typeof u.handleEvent && (c = u, u = u.handleEvent), "function" == typeof u && (f = E({}, e), ve(u, c, [f], n))
					}
				}
			},
			ut = function(e) {
				return "string" == typeof e && (e = []), "number" != typeof e.length ? [e] : e
			},
			ct = function(e) {
				if (e && 1 === e.nodeType) {
					var t = function(e) {
							(e || (e = a.event)) && ("js" !== e._source && (e.stopImmediatePropagation(), e.preventDefault()), delete e._source)
						},
						n = function(n) {
							(n || (n = a.event)) && (t(n), Ve.focus(e))
						};
					e.addEventListener("mouseover", n, !1), e.addEventListener("mouseout", t, !1), e.addEventListener("mouseenter", t, !1), e.addEventListener("mouseleave", t, !1), e.addEventListener("mousemove", t, !1), Ke[e.zcClippingId] = {
						mouseover: n,
						mouseout: t,
						mouseenter: t,
						mouseleave: t,
						mousemove: t
					}
				}
			},
			ft = function(e) {
				if (e && 1 === e.nodeType) {
					var t = Ke[e.zcClippingId];
					if ("object" == typeof t && t) {
						for (var n, r, i = ["move", "leave", "enter", "out", "over"], a = 0, o = i.length; o > a; a++) n = "mouse" + i[a], r = t[n], "function" == typeof r && e.removeEventListener(n, r, !1);
						delete Ke[e.zcClippingId]
					}
				}
			};
		Ve._createClient = function() {
			Ge.apply(this, N(arguments))
		}, Ve.prototype.on = function() {
			return Qe.apply(this, N(arguments))
		}, Ve.prototype.off = function() {
			return et.apply(this, N(arguments))
		}, Ve.prototype.handlers = function() {
			return tt.apply(this, N(arguments))
		}, Ve.prototype.emit = function() {
			return nt.apply(this, N(arguments))
		}, Ve.prototype.clip = function() {
			return rt.apply(this, N(arguments))
		}, Ve.prototype.unclip = function() {
			return it.apply(this, N(arguments))
		}, Ve.prototype.elements = function() {
			return at.apply(this, N(arguments))
		}, Ve.prototype.destroy = function() {
			return ot.apply(this, N(arguments))
		}, Ve.prototype.setText = function(e) {
			if (!Ue[this.id]) throw new Error("Attempted to set pending clipboard data from a destroyed ZeroClipboard client instance");
			return Ve.setData("text/plain", e), this
		}, Ve.prototype.setHtml = function(e) {
			if (!Ue[this.id]) throw new Error("Attempted to set pending clipboard data from a destroyed ZeroClipboard client instance");
			return Ve.setData("text/html", e), this
		}, Ve.prototype.setRichText = function(e) {
			if (!Ue[this.id]) throw new Error("Attempted to set pending clipboard data from a destroyed ZeroClipboard client instance");
			return Ve.setData("application/rtf", e), this
		}, Ve.prototype.setData = function() {
			if (!Ue[this.id]) throw new Error("Attempted to set pending clipboard data from a destroyed ZeroClipboard client instance");
			return Ve.setData.apply(this, N(arguments)), this
		}, Ve.prototype.clearData = function() {
			if (!Ue[this.id]) throw new Error("Attempted to clear pending clipboard data from a destroyed ZeroClipboard client instance");
			return Ve.clearData.apply(this, N(arguments)), this
		}, Ve.prototype.getData = function() {
			if (!Ue[this.id]) throw new Error("Attempted to get pending clipboard data from a destroyed ZeroClipboard client instance");
			return Ve.getData.apply(this, N(arguments))
		}, "function" == typeof define && define.amd ? define("/ZeroClipboard", function() {
			return Ve
		}) : "object" == typeof module && module && "object" == typeof module.exports && module.exports ? module.exports = Ve : e.ZeroClipboard = Ve
	}(function() {
		return this || window
	}())
});
define("/css3-mediaqueries", function(e, t) {
	"function" != typeof Object.create && (Object.create = function(e) {
		function t() {}
		return t.prototype = e, new t
	});
	var n = {
		toString: function() {
			return navigator.userAgent
		},
		test: function(e) {
			return this.toString().toLowerCase().indexOf(e.toLowerCase()) > -1
		}
	};
	n.version = (n.toString().toLowerCase().match(/[\s\S]+(?:rv|it|ra|ie)[\/: ]([\d.]+)/) || [])[1], n.webkit = n.test("webkit"), n.gecko = n.test("gecko") && !n.webkit, n.opera = n.test("opera"), n.ie = n.test("msie") && !n.opera, n.ie6 = n.ie && document.compatMode && "undefined" == typeof document.documentElement.style.maxHeight, n.ie7 = n.ie && document.documentElement && "undefined" != typeof document.documentElement.style.maxHeight && "undefined" == typeof XDomainRequest, n.ie8 = n.ie && "undefined" != typeof XDomainRequest;
	var r = function() {
			var e = [],
				t = function() {
					if (!arguments.callee.done) {
						arguments.callee.done = !0;
						for (var t = 0; t < e.length; t++) e[t]()
					}
				};
			return document.addEventListener && document.addEventListener("DOMContentLoaded", t, !1), n.ie && (! function() {
					try {
						document.documentElement.doScroll("left"), document.body.length
					} catch (e) {
						return void setTimeout(arguments.callee, 50)
					}
					t()
				}(), document.onreadystatechange = function() {
					"complete" === document.readyState && (document.onreadystatechange = null, t())
				}), n.webkit && document.readyState && ! function() {
					"loading" !== document.readyState ? t() : setTimeout(arguments.callee, 10)
				}(), window.onload = t,
				function(n) {
					return "function" == typeof n && (t.done ? n() : e[e.length] = n), n
				}
		}(),
		i = function() {
			var e, t = {
					BLOCKS: /[^\s{][^{]*\{(?:[^{}]*\{[^{}]*\}[^{}]*|[^{}]*)*\}/g,
					BLOCKS_INSIDE: /[^\s{][^{]*\{[^{}]*\}/g,
					DECLARATIONS: /[a-zA-Z\-]+[^;]*:[^;]+;/g,
					RELATIVE_URLS: /url\(['"]?([^\/\)'"][^:\)'"]+)['"]?\)/g,
					REDUNDANT_COMPONENTS: /(?:\/\*([^*\\\\]|\*(?!\/))+\*\/|@import[^;]+;|@-moz-document\s*url-prefix\(\)\s*{(([^{}])+{([^{}])+}([^{}])+)+})/g,
					REDUNDANT_WHITESPACE: /\s*(,|:|;|\{|\})\s*/g,
					MORE_WHITESPACE: /\s{2,}/g,
					FINAL_SEMICOLONS: /;\}/g,
					NOT_WHITESPACE: /\S+/g
				},
				o = !1,
				a = [],
				s = function(e) {
					"function" == typeof e && (a[a.length] = e)
				},
				l = function() {
					for (var t = 0; t < a.length; t++) a[t](e)
				},
				u = {},
				c = function(e, t) {
					if (u[e]) {
						var n = u[e].listeners;
						if (n)
							for (var r = 0; r < n.length; r++) n[r](t)
					}
				},
				f = function(e, t, r) {
					if (n.ie && !window.XMLHttpRequest && (window.XMLHttpRequest = function() {
							return new ActiveXObject("Microsoft.XMLHTTP")
						}), !XMLHttpRequest) return "";
					var i = new XMLHttpRequest;
					try {
						i.open("get", e, !0), i.setRequestHeader("X_REQUESTED_WITH", "XMLHttpRequest")
					} catch (o) {
						return void r()
					}
					var a = !1;
					setTimeout(function() {
						a = !0
					}, 5e3), document.documentElement.style.cursor = "progress", i.onreadystatechange = function() {
						4 !== i.readyState || a || (!i.status && "file:" === location.protocol || i.status >= 200 && i.status < 300 || 304 === i.status || navigator.userAgent.indexOf("Safari") > -1 && "undefined" == typeof i.status ? t(i.responseText) : r(), document.documentElement.style.cursor = "", i = null)
					}, i.send("")
				},
				d = function(e) {
					return e = e.replace(t.REDUNDANT_COMPONENTS, ""), e = e.replace(t.REDUNDANT_WHITESPACE, "$1"), e = e.replace(t.MORE_WHITESPACE, " "), e = e.replace(t.FINAL_SEMICOLONS, "}")
				},
				p = {
					mediaQueryList: function(e) {
						var n = {},
							r = e.indexOf("{"),
							i = e.substring(0, r);
						e = e.substring(r + 1, e.length - 1);
						for (var o = [], a = [], s = i.toLowerCase().substring(7).split(","), l = 0; l < s.length; l++) o[o.length] = p.mediaQuery(s[l], n);
						var u = e.match(t.BLOCKS_INSIDE);
						if (null !== u)
							for (l = 0; l < u.length; l++) a[a.length] = p.rule(u[l], n);
						return n.getMediaQueries = function() {
							return o
						}, n.getRules = function() {
							return a
						}, n.getListText = function() {
							return i
						}, n.getCssText = function() {
							return e
						}, n
					},
					mediaQuery: function(e, n) {
						e = e || "";
						for (var r, i = !1, o = [], a = !0, s = e.match(t.NOT_WHITESPACE), l = 0; l < s.length; l++) {
							var u = s[l];
							if (r || "not" !== u && "only" !== u)
								if (r) {
									if ("(" === u.charAt(0)) {
										var c = u.substring(1, u.length - 1).split(":");
										o[o.length] = {
											mediaFeature: c[0],
											value: c[1] || null
										}
									}
								} else r = u;
							else "not" === u && (i = !0)
						}
						return {
							getList: function() {
								return n || null
							},
							getValid: function() {
								return a
							},
							getNot: function() {
								return i
							},
							getMediaType: function() {
								return r
							},
							getExpressions: function() {
								return o
							}
						}
					},
					rule: function(e, t) {
						for (var n = {}, r = e.indexOf("{"), i = e.substring(0, r), o = i.split(","), a = [], s = e.substring(r + 1, e.length - 1).split(";"), l = 0; l < s.length; l++) a[a.length] = p.declaration(s[l], n);
						return n.getMediaQueryList = function() {
							return t || null
						}, n.getSelectors = function() {
							return o
						}, n.getSelectorText = function() {
							return i
						}, n.getDeclarations = function() {
							return a
						}, n.getPropertyValue = function(e) {
							for (var t = 0; t < a.length; t++)
								if (a[t].getProperty() === e) return a[t].getValue();
							return null
						}, n
					},
					declaration: function(e, t) {
						var n = e.indexOf(":"),
							r = e.substring(0, n),
							i = e.substring(n + 1);
						return {
							getRule: function() {
								return t || null
							},
							getProperty: function() {
								return r
							},
							getValue: function() {
								return i
							}
						}
					}
				},
				h = function(n) {
					if ("string" == typeof n.cssHelperText) {
						var r = {
								mediaQueryLists: [],
								rules: [],
								selectors: {},
								declarations: [],
								properties: {}
							},
							i = r.mediaQueryLists,
							o = r.rules,
							a = n.cssHelperText.match(t.BLOCKS);
						if (null !== a)
							for (var s = 0; s < a.length; s++) "@media " === a[s].substring(0, 7) ? (i[i.length] = p.mediaQueryList(a[s]), o = r.rules = o.concat(i[i.length - 1].getRules())) : o[o.length] = p.rule(a[s]);
						var l = r.selectors,
							u = function(e) {
								for (var t = e.getSelectors(), n = 0; n < t.length; n++) {
									var r = t[n];
									l[r] || (l[r] = []), l[r][l[r].length] = e
								}
							};
						for (s = 0; s < o.length; s++) u(o[s]);
						var c = r.declarations;
						for (s = 0; s < o.length; s++) c = r.declarations = c.concat(o[s].getDeclarations());
						var f = r.properties;
						for (s = 0; s < c.length; s++) {
							var d = c[s].getProperty();
							f[d] || (f[d] = []), f[d][f[d].length] = c[s]
						}
						return n.cssHelperParsed = r, e[e.length] = n, r
					}
				},
				m = function(e, t) {
					return e.cssHelperText = d(t || e.innerHTML), h(e)
				},
				g = function() {
					o = !0, e = [];
					for (var n = [], r = function() {
							for (var e = 0; e < n.length; e++) h(n[e]);
							var t = document.getElementsByTagName("style");
							for (e = 0; e < t.length; e++) m(t[e]);
							o = !1, l()
						}, i = document.getElementsByTagName("link"), a = 0; a < i.length; a++) {
						var s = i[a];
						s.getAttribute("rel").indexOf("style") > -1 && s.href && 0 !== s.href.length && !s.disabled && (n[n.length] = s)
					}
					if (n.length > 0) {
						var u = 0,
							c = function() {
								u++, u === n.length && r()
							},
							p = function(e) {
								var n = e.href;
								f(n, function(r) {
									r = d(r).replace(t.RELATIVE_URLS, "url(" + n.substring(0, n.lastIndexOf("/")) + "/$1)"), e.cssHelperText = r, c()
								}, c)
							};
						for (a = 0; a < n.length; a++) p(n[a])
					} else r()
				},
				y = {
					mediaQueryLists: "array",
					rules: "array",
					selectors: "object",
					declarations: "array",
					properties: "object"
				},
				v = {
					mediaQueryLists: null,
					rules: null,
					selectors: null,
					declarations: null,
					properties: null
				},
				b = function(e, t) {
					if (null !== v[e]) {
						if ("array" === y[e]) return v[e] = v[e].concat(t);
						var n = v[e];
						for (var r in t) t.hasOwnProperty(r) && (n[r] ? n[r] = n[r].concat(t[r]) : n[r] = t[r]);
						return n
					}
				},
				x = function(t) {
					v[t] = "array" === y[t] ? [] : {};
					for (var n = 0; n < e.length; n++) b(t, e[n].cssHelperParsed[t]);
					return v[t]
				};
			r(function() {
				for (var e = document.body.getElementsByTagName("*"), t = 0; t < e.length; t++) e[t].checkedByCssHelper = !0;
				document.implementation.hasFeature("MutationEvents", "2.0") || window.MutationEvent ? document.body.addEventListener("DOMNodeInserted", function(e) {
					var t = e.target;
					1 === t.nodeType && (c("DOMElementInserted", t), t.checkedByCssHelper = !0)
				}, !1) : setInterval(function() {
					for (var e = document.body.getElementsByTagName("*"), t = 0; t < e.length; t++) e[t].checkedByCssHelper || (c("DOMElementInserted", e[t]), e[t].checkedByCssHelper = !0)
				}, 1e3)
			});
			var w = function(e) {
				return "undefined" != typeof window.innerWidth ? window["inner" + e] : "undefined" != typeof document.documentElement && "undefined" != typeof document.documentElement.clientWidth && 0 != document.documentElement.clientWidth ? document.documentElement["client" + e] : void 0
			};
			return {
				addStyle: function(e, t) {
					var n;
					return null !== document.getElementById("css-mediaqueries-js") ? n = document.getElementById("css-mediaqueries-js") : (n = document.createElement("style"), n.setAttribute("type", "text/css"), n.setAttribute("id", "css-mediaqueries-js"), document.getElementsByTagName("head")[0].appendChild(n)), n.styleSheet ? n.styleSheet.cssText += e : n.appendChild(document.createTextNode(e)), n.addedWithCssHelper = !0, "undefined" == typeof t || t === !0 ? i.parsed(function(t) {
						var r = m(n, e);
						for (var i in r) r.hasOwnProperty(i) && b(i, r[i]);
						c("newStyleParsed", n)
					}) : n.parsingDisallowed = !0, n
				},
				removeStyle: function(e) {
					return e.parentNode ? e.parentNode.removeChild(e) : void 0
				},
				parsed: function(t) {
					o ? s(t) : "undefined" != typeof e ? "function" == typeof t && t(e) : (s(t), g())
				},
				mediaQueryLists: function(e) {
					i.parsed(function(t) {
						e(v.mediaQueryLists || x("mediaQueryLists"))
					})
				},
				rules: function(e) {
					i.parsed(function(t) {
						e(v.rules || x("rules"))
					})
				},
				selectors: function(e) {
					i.parsed(function(t) {
						e(v.selectors || x("selectors"))
					})
				},
				declarations: function(e) {
					i.parsed(function(t) {
						e(v.declarations || x("declarations"))
					})
				},
				properties: function(e) {
					i.parsed(function(t) {
						e(v.properties || x("properties"))
					})
				},
				broadcast: c,
				addListener: function(e, t) {
					"function" == typeof t && (u[e] || (u[e] = {
						listeners: []
					}), u[e].listeners[u[e].listeners.length] = t)
				},
				removeListener: function(e, t) {
					if ("function" == typeof t && u[e])
						for (var n = u[e].listeners, r = 0; r < n.length; r++) n[r] === t && (n.splice(r, 1), r -= 1)
				},
				getViewportWidth: function() {
					return w("Width")
				},
				getViewportHeight: function() {
					return w("Height")
				}
			}
		}();
	r(function() {
		var e, t = {
				LENGTH_UNIT: /[0-9]+(em|ex|px|in|cm|mm|pt|pc)$/,
				RESOLUTION_UNIT: /[0-9]+(dpi|dpcm)$/,
				ASPECT_RATIO: /^[0-9]+\/[0-9]+$/,
				ABSOLUTE_VALUE: /^[0-9]*(\.[0-9]+)*$/
			},
			r = [],
			o = function() {
				var e = "css3-mediaqueries-test",
					t = document.createElement("div");
				t.id = e;
				var n = i.addStyle("@media all and (width) { #" + e + " { width: 1px !important; } }", !1);
				document.body.appendChild(t);
				var r = 1 === t.offsetWidth;
				return n.parentNode.removeChild(n), t.parentNode.removeChild(t), o = function() {
					return r
				}, r
			},
			a = function() {
				e = document.createElement("div"), e.style.cssText = "position:absolute;top:-9999em;left:-9999em;margin:0;border:none;padding:0;width:1em;font-size:1em;", document.body.appendChild(e), 16 !== e.offsetWidth && (e.style.fontSize = 16 / e.offsetWidth + "em"), e.style.width = ""
			},
			s = function(t) {
				e.style.width = t;
				var n = e.offsetWidth;
				return e.style.width = "", n
			},
			l = function(e, n) {
				var r = e.length,
					i = "min-" === e.substring(0, 4),
					o = !i && "max-" === e.substring(0, 4);
				if (null !== n) {
					var a, l;
					if (t.LENGTH_UNIT.exec(n)) a = "length", l = s(n);
					else if (t.RESOLUTION_UNIT.exec(n)) {
						a = "resolution", l = parseInt(n, 10);
						var u = n.substring((l + "").length)
					} else t.ASPECT_RATIO.exec(n) ? (a = "aspect-ratio", l = n.split("/")) : t.ABSOLUTE_VALUE ? (a = "absolute", l = n) : a = "unknown"
				}
				var c, f;
				if ("device-width" === e.substring(r - 12, r)) return c = screen.width, null !== n ? "length" === a ? i && c >= l || o && l > c || !i && !o && c === l : !1 : c > 0;
				if ("device-height" === e.substring(r - 13, r)) return f = screen.height, null !== n ? "length" === a ? i && f >= l || o && l > f || !i && !o && f === l : !1 : f > 0;
				if ("width" === e.substring(r - 5, r)) return c = document.documentElement.clientWidth || document.body.clientWidth, null !== n ? "length" === a ? i && c >= l || o && l > c || !i && !o && c === l : !1 : c > 0;
				if ("height" === e.substring(r - 6, r)) return f = document.documentElement.clientHeight || document.body.clientHeight, null !== n ? "length" === a ? i && f >= l || o && l > f || !i && !o && f === l : !1 : f > 0;
				if ("orientation" === e.substring(r - 11, r)) return c = document.documentElement.clientWidth || document.body.clientWidth, f = document.documentElement.clientHeight || document.body.clientHeight, "absolute" === a ? "portrait" === l ? f >= c : c > f : !1;
				if ("aspect-ratio" === e.substring(r - 12, r)) {
					c = document.documentElement.clientWidth || document.body.clientWidth, f = document.documentElement.clientHeight || document.body.clientHeight;
					var d = c / f,
						p = l[1] / l[0];
					return "aspect-ratio" === a ? i && d >= p || o && p > d || !i && !o && d === p : !1
				}
				if ("device-aspect-ratio" === e.substring(r - 19, r)) return "aspect-ratio" === a && screen.width * l[1] === screen.height * l[0];
				if ("color-index" === e.substring(r - 11, r)) {
					var h = Math.pow(2, screen.colorDepth);
					return null !== n ? "absolute" === a ? i && h >= l || o && l > h || !i && !o && h === l : !1 : h > 0
				}
				if ("color" === e.substring(r - 5, r)) {
					var m = screen.colorDepth;
					return null !== n ? "absolute" === a ? i && m >= l || o && l > m || !i && !o && m === l : !1 : m > 0
				}
				if ("resolution" === e.substring(r - 10, r)) {
					var g;
					return g = s("dpcm" === u ? "1cm" : "1in"), null !== n ? "resolution" === a ? i && g >= l || o && l > g || !i && !o && g === l : !1 : g > 0
				}
				return !1
			},
			u = function(e) {
				var t = e.getValid(),
					n = e.getExpressions(),
					r = n.length;
				if (r > 0) {
					for (var i = 0; r > i && t; i++) t = l(n[i].mediaFeature, n[i].value);
					var o = e.getNot();
					return t && !o || o && !t
				}
			},
			c = function(e) {
				for (var t = e.getMediaQueries(), n = {}, o = 0; o < t.length; o++) u(t[o]) && (n[t[o].getMediaType()] = !0);
				var a = [],
					s = 0;
				for (var l in n) n.hasOwnProperty(l) && (s > 0 && (a[s++] = ","), a[s++] = l);
				a.length > 0 && (r[r.length] = i.addStyle("@media " + a.join("") + "{" + e.getCssText() + "}", !1))
			},
			f = function(e) {
				for (var t = 0; t < e.length; t++) c(e[t]);
				n.ie ? (document.documentElement.style.display = "block", setTimeout(function() {
					document.documentElement.style.display = ""
				}, 0), setTimeout(function() {
					i.broadcast("cssMediaQueriesTested")
				}, 100)) : i.broadcast("cssMediaQueriesTested")
			},
			d = function() {
				for (var e = 0; e < r.length; e++) i.removeStyle(r[e]);
				r = [], i.mediaQueryLists(f)
			},
			p = 0,
			h = function() {
				var e = i.getViewportWidth(),
					t = i.getViewportHeight();
				if (n.ie) {
					var r = document.createElement("div");
					r.style.width = "100px", r.style.height = "100px", r.style.position = "absolute", r.style.top = "-9999em", r.style.overflow = "scroll", document.body.appendChild(r), p = r.offsetWidth - r.clientWidth, document.body.removeChild(r)
				}
				var a, s = function() {
					var n = i.getViewportWidth(),
						r = i.getViewportHeight();
					(Math.abs(n - e) > p || Math.abs(r - t) > p) && (e = n, t = r, clearTimeout(a), a = setTimeout(function() {
						o() ? i.broadcast("cssMediaQueriesTested") : d()
					}, 500))
				};
				window.onresize = function() {
					var e = window.onresize || function() {};
					return function() {
						e(), s()
					}
				}()
			},
			m = document.documentElement;
		return m.style.marginLeft = "-32767px", setTimeout(function() {
				m.style.marginTop = ""
			}, 2e4),
			function() {
				o() ? m.style.marginLeft = "" : (i.addListener("newStyleParsed", function(e) {
					f(e.cssHelperParsed.mediaQueryLists)
				}), i.addListener("cssMediaQueriesTested", function() {
					n.ie && (m.style.width = "1px"), setTimeout(function() {
						m.style.width = "", m.style.marginLeft = ""
					}, 0), i.removeListener("cssMediaQueriesTested", arguments.callee)
				}), a(), d()), h()
			}
	}());
	try {
		document.execCommand("BackgroundImageCache", !1, !0)
	} catch (o) {}
});
define("/log", function(e, t) {
	var n = e("/lib/jquery-1.6.2");
	t.clog = function(e, t, n) {
		n ? (new Image).src = "/ctlog?pos=" + t + "&action=" + e + "&check=" + n : (new Image).src = "/ctlog?pos=" + t + "&action=" + e
	}, t.init = function() {
		n(".clog-js").live("click", function() {
			var e = n(this),
				r = e.data("clog"),
				i = e.data("pos"),
				o = e.prop("checked");
			t.clog(r, i, o)
		})
	}
});
define("/utils", function(e) {
	var t = e("/lib/jquery-1.6.2");
	return {
		utf8_decode: function(e) {
			return unescape(encodeURIComponent(e))
		},
		isWeb: function(e) {
			var t = !1,
				n = /^((http|https):\/\/)?([\w\d\.]+)(\.)(com|edu|gov|int|mil|net|org|biz|info|pro|name|museum|coop|aero|xxx|idv|al|dz|af|ar|ae|aw|om|az|eg|et|ie|ee|ad|ao|ai|ag|at|au|mo|bb|pg|bs|pk|py|ps|bh|pa|br|by|bm|bg|mp|bj|be|is|pr|ba|pl|bo|bz|bw|bt|bf|bi|bv|kp|gq|dk|de|tl|tp|tg|dm|do|ru|ec|er|fr|fo|pf|gf|tf|va|ph|fj|fi|cv|fk|gm|cg|cd|co|cr|gg|gd|gl|ge|cu|gp|gu|gy|kz|ht|kr|nl|an|hm|hn|ki|dj|kg|gn|gw|ca|gh|ga|kh|cz|zw|cm|qa|ky|km|ci|kw|cc|hr|ke|ck|lv|ls|la|lb|lt|lr|ly|li|re|lu|rw|ro|mg|im|mv|mt|mw|my|ml|mk|mh|mq|yt|mu|mr|us|um|as|vi|mn|ms|bd|pe|fm|mm|md|ma|mc|mz|mx|nr|np|ni|ne|ng|nu|no|nf|na|za|zq|aq|gs|eu|pw|pn|pt|jp|se|ch|sv|ws|yu|sl|sn|cy|sc|sa|cx|st|sh|kn|lc|sm|pm|vc|lk|sk|si|sj|sz|sd|sr|sb|so|tj|tw|th|tz|to|tc|tt|tn|tv|tr|tm|tk|wf|vu|gt|ve|bn|ug|ua|uy|uz|es|eh|gr|hk|sg|nc|nz|hu|sy|jm|am|ac|ye|iq|ir|il|it|in|id|uk|vg|io|jo|vn|zm|je|td|gi|cl|cf|cn)(\/[^\s]+)?(\/)?$/,
				r = /^((http|https):\/\/)?([1-9]|([1-9][0-9])|((1[0-9]{2})|(2(([0-4][0-9])|(5[0-5])))))\.([0-9]|([1-9][0-9])|((1[0-9]{2})|(2(([0-4][0-9])|(5[0-5])))))\.([0-9]|([1-9][0-9])|((1[0-9]{2})|(2(([0-4][0-9])|(5[0-5])))))\.([0-9]|([1-9][0-9])|((1[0-9]{2})|(2(([0-4][0-9])|(5[0-5])))))(\:\d+)?(\/[^\s]+)?(\/)?$/;
			return t = !!n.exec(e.toLowerCase()) || !!r.exec(e.toLowerCase())
		},
		addToFav: function() {
			t.browser.msie ? window.external.addFavorite("http://fanyi.youdao.com/?keyfrom=favorite", "有道翻译 - 免费在线翻译") : alert("请使用Ctrl + D 手动加入收藏")
		},
		cookie: function(e, t, n) {
			function r(e, t, n) {
				n = n || 30;
				var r = new Date;
				r.setTime(r.getTime() + 24 * n * 60 * 60 * 1e3), document.cookie = e + "=" + encodeURIComponent(t) + ";expires=" + r.toGMTString()
			}

			function i(e) {
				var t = document.cookie.match(new RegExp("(^| )" + e + "=([^;]*)(;|$)"));
				return t ? decodeURIComponent(t[2]) : null
			}
			return void 0 === t ? i(e) : void r(e, t, n)
		},
		timerProxy: function() {
			var e = function() {
				window.timerProxyTimeout && window.clearTimeout(window.timerProxyTimeout)
			};
			return function n(r, i) {
				this.timerProxy.clearProxy = n.clearProxy = e, e(), window.timerProxyTimeout = window.setTimeout(function() {
					t.isFunction(r) && r()
				}, i)
			}
		}()
	}
});
define("/constData", function(e, t) {
	t.LANGUAGETYPES = {
		AUTO: "自动检测语言",
		EN2ZH_CN: "英语 &raquo; 中文",
		ZH_CN2EN: "中文 &raquo; 英语",
		JA2ZH_CN: "日语 &raquo; 中文",
		ZH_CN2JA: "中文 &raquo; 日语",
		FR2ZH_CN: "法语 &raquo; 中文",
		ZH_CN2FR: "中文 &raquo; 法语",
		KR2ZH_CN: "韩语 &raquo; 中文",
		ZH_CN2KR: "中文 &raquo; 韩语",
		RU2ZH_CN: "俄语 &raquo; 中文",
		ZH_CN2RU: "中文 &raquo; 俄语",
		SP2ZH_CN: "西语 &raquo; 中文",
		ZH_CN2SP: "中文 &raquo; 西语"
	}, t.LANGUAGEDESCRIPTIONS = {
		"英语 » 中文": "EN2ZH_CN",
		"中文 » 英语": "ZH_CN2EN",
		"日语 » 中文": "JA2ZH_CN",
		"中文 » 日语": "ZH_CN2JA",
		"法语 » 中文": "FR2ZH_CN",
		"中文 » 法语": "ZH_CN2FR",
		"韩语 » 中文": "KR2ZH_CN",
		"中文 » 韩语": "ZH_CN2KR",
		"俄语 » 中文": "RU2ZH_CN",
		"中文 » 俄语": "ZH_CN2RU",
		"西语 » 中文": "SP2ZH_CN",
		"中文 » 西语": "ZH_CN2SP"
	}, t.errorMsg = {
		10: "抱歉，个别句子太长啦，我读不懂",
		20: "抱歉，超过2万字实在太长啦，让我喘口气",
		30: "抱歉，我绞尽脑汁了",
		40: "抱歉，我还在学习该语种中",
		transRequestError: "翻译出错，请检查网络后重试"
	}, t.smartSRC = {
		1: "有道词典结果",
		2: "有道智能翻译结果",
		3: "微博粉丝犀利翻译"
	}, t.smartText = {
		1: "查看详细结果",
		2: "",
		3: '猛击参加"有道翻译题" show出你的犀利'
	}
});
define("/tips", function(e, t) {
	var n, r = e("/lib/jquery-1.6.2");
	t.hide = function(e) {
		switch (e) {
			case "all":
				n.fadeOut();
				break;
			case "welcome":
				r("#errorHolder").find(".error_text:contains('有道翻译')").html() && n.fadeOut();
				break;
			case "allExceptWelcome":
				n.find(".error_text:contains('有道翻译')").html() || n.fadeOut()
		}
	}, t.show = function(e) {
		n.hide(), n.find(".error_text").html(e), "请输入内容" === e ? n.addClass("nullError") : n.removeClass("nullError"), n.fadeIn()
	}, t.init = function() {
		n = r("#errorHolder"), r("body").click(function() {
			t.hide("allExceptWelcome")
		})
	}, t.transError = function(n) {
		n > 0 && t.show(e("/constData").errorMsg[n])
	}, t.transRequestError = function(n) {
		var r = n.msg || e("/constData").errorMsg.transRequestError;
		t.show(r)
	}
});
define("/addFav", function(e, t) {
	var n = e("/lib/jquery-1.6.2");
	t.create = function(t) {
		var r = e("/utils"),
			i = n.browser.msie ? "添加到收藏夹" : "添加书签",
			o = n.browser.msie ? "CLICK_ADDFAV" : "CLICK_BOOKMARK",
			a = '<a href="#" class="add-fav clog-js" data-clog="' + o + '" data-pos="' + t + '"><strong>' + i + "</strong></a>";
		return n(a).click(function(t) {
			"ADD" !== r.cookie("YOUDAO_FANYI_FAV") && r.cookie("YOUDAO_FANYI_FAV", "ADD"), r.addToFav(), e("/tips").hide("welcome"), t.preventDefault()
		})
	}
});
define("/navigateBar", function(e) {
	function t() {
		n(window).width() <= 900 ? setTimeout(function() {
			n("html").addClass("win900")
		}, 100) : setTimeout(function() {
			n("html").removeClass("win900")
		}, 100)
	}
	var n = e("/lib/jquery-1.6.2"),
		r = function() {
			n(".c-sust").append(n('<span class="c-sp">|</span>'), e("/addFav").create("web.righttop"))
		},
		i = function() {
			function e() {
				var e = Number(n("body").css("margin-top").replace("px", ""));
				n(".un_op").css({
					left: t.offset().left - n(".user-info").offset().left,
					top: t.position().top + t.height() - 1 - e
				})
			}
			var t = n(".un_box");
			t.size() > 0 && (n(window).resize(e), n(document).click(function() {
				t.removeClass("un_box_on"), n("body").removeClass("header-namelist-open"), n(".un_op").hide()
			}))
		},
		o = function(e, t, n) {
			e.href = "http://" + t + ".youdao.com/", window.RegExp && window.encodeURIComponent && (e.href = e.href + "?keyfrom=" + n)
		};
	n(function() {
		function e() {
			l ? (u.removeClass("sidebarButton-over").attr("title", "关闭左栏导航"), t.animate(navigator.userAgent.indexOf("MSIE 6") > -1 ? {
				left: 170
			} : {
				left: 0
			}), a.animate({
				paddingLeft: 170
			}, function() {
				a.toggleClass(s)
			}), l = !1, document.cookie = "resultSideCollapse=" + l + ";expires=" + d.toGMTString()) : (u.removeClass("sidebarButton-over").attr("title", "打开左栏导航"), t.animate({
				left: -150
			}), a.animate({
				paddingLeft: 20
			}, function() {
				a.toggleClass(s)
			}), l = !0, document.cookie = "resultSideCollapse=" + l + ";expires=" + d.toGMTString())
		}
		r(), i(), n(".uri-js").click(function() {
			return o(this, n(this).data("product"), n(this).data("pos")), !0
		}), n("#mn").hasMenu("#nm", ".pm", 0, 17), n("#deleteHistory").hasMenu("#historyConfirm", ".pm", 0, 15), n("#recommenderWidget li a,#historyWidget li a").each(function() {
			var e = n(this);
			e.height() > e.parent("li").height() && e.attr("title", e.text())
		});
		var t = n("#s"),
			a = n("#w"),
			s = "side-collapsed",
			l = a.hasClass(s),
			u = n("#sideBarButton"),
			c = n('<div style="position:absolute;left:0;top:10px;width:13px;height:20px;cursor:pointer">&nbsp;</div>').css("cursor", "pointer").click(e).hover(function() {
				u.addClass("sidebarButton-over")
			}, function() {
				u.removeClass("sidebarButton-over")
			});
		u.append(c);
		var f = new Date,
			d = new Date;
		d.setTime(f.getTime() + 31536e7)
	}), n.fn.extend({
		hasMenu: function(e, t, r, i) {
			var o = n(e),
				a = n(t);
			return n("body").click(function() {
				a.hide()
			}), this.click(function() {
				var t = n(this).offset();
				return a.each(function() {
					"#" + n(this).attr("id") != e && n(this).hide()
				}), n(this).blur(), o.css({
					position: "absolute",
					"z-index": 1e3,
					left: t.left + r,
					top: t.top + i
				}).toggle(), !1
			}), this
		}
	}), n(window).resize(function() {
		t()
	}), t()
});
define("/example", function(e, t) {
	var n = e("/lib/jquery-1.6.2"),
		r = n("#translateBtn"),
		i = !1,
		o = function l(e, t, r) {
			return r = r || 1, r == e.length + 1 ? void t() : (n("#inputText").focus().val(e.substring(0, r++)), void setTimeout(function() {
				l(e, t, r)
			}, 150))
		},
		a = function() {
			var e = n("#transBtnTip");
			return e.hide().fadeIn(), e
		},
		s = function(e) {
			o(e, function() {
				setTimeout(function() {
					var e = a();
					n("#transBtnTipOK").click(function(t) {
						e.fadeOut(), t.preventDefault()
					}), r.click(function() {
						e.hide()
					}), i = !1
				}, 500)
			})
		};
	t.hide = function() {
		n("#entryList").hide()
	}, t.init = function() {
		n(".ext_web").click(function() {
			return i || s("http://www.chinadaily.com.cn"), !1
		})
	}
});
define("/lib/jquery-extension", function(e, t) {
	var n = e("/lib/jquery-1.6.2");
	n.fn.extend({
			selectNodeText: function() {
				if (n.browser.msie) {
					var e = document.body.createTextRange();
					e.moveToElementText(this.get(0)), e.select()
				} else {
					var t = window.getSelection(),
						r = document.createRange();
					r.setStartBefore(this.get(0)), r.setEndAfter(this.get(0)), t.removeAllRanges(), t.addRange(r)
				}
			}
		}),
		function(e) {
			e.event.special.textchange = {
				setup: function(t, n) {
					e(this).data("lastValue", "true" === this.contentEditable ? e(this).html() : e(this).val()), e(this).bind("keyup.textchange", e.event.special.textchange.handler), e(this).bind("cut.textchange paste.textchange input.textchange", e.event.special.textchange.delayedHandler)
				},
				teardown: function(t) {
					e(this).unbind(".textchange")
				},
				handler: function(t) {
					e.event.special.textchange.triggerIfChanged(e(this))
				},
				delayedHandler: function(t) {
					var n = e(this);
					setTimeout(function() {
						e.event.special.textchange.triggerIfChanged(n)
					}, 25)
				},
				triggerIfChanged: function(e) {
					var t = "true" === e[0].contentEditable ? e.html() : e.val();
					t !== e.data("lastValue") && (e.trigger("textchange", [e.data("lastValue")]), e.data("lastValue", t))
				}
			}, e.event.special.hastext = {
				setup: function(t, n) {
					e(this).bind("textchange", e.event.special.hastext.handler)
				},
				teardown: function(t) {
					e(this).unbind("textchange", e.event.special.hastext.handler)
				},
				handler: function(t, n) {
					"" === n && n !== e(this).val() && e(this).trigger("hastext")
				}
			}, e.event.special.notext = {
				setup: function(t, n) {
					e(this).bind("textchange", e.event.special.notext.handler)
				},
				teardown: function(t) {
					e(this).unbind("textchange", e.event.special.notext.handler)
				},
				handler: function(t, n) {
					"" === e(this).val() && e(this).val() !== n && e(this).trigger("notext")
				}
			}
		}(n)
});
define("/output/bilingualControl", function(e, t) {
	var n, r = e("/lib/jquery-1.6.2");
	t.isMutiMode = function() {
		return n
	}, t.changeMutiMode = function() {
		r("#compare").click()
	}, t.init = function() {
		r("#compare").click(function() {
			n = !n, e("/output/result").show(), r("#inputText").blur()
		});
		var t = r("#compare"),
			i = r(".compare-mode");
		t.bind("mouseover", function() {
			i.addClass("hover")
		}), t.bind("mouseout", function() {
			i.removeClass("hover")
		}), t.bind("click", function() {
			i.hasClass("compare-enable") ? (i.removeClass("compare-enable"), i.addClass("compare-disable")) : i.hasClass("compare-disable") && (i.removeClass("compare-disable"), i.addClass("compare-enable"))
		})
	}
});
define("/output/voice", function(e, t) {
	var n = e("/lib/jquery-1.6.2"),
		r = n("#outputText"),
		i = n("#speech"),
		o = function(e) {
			return -1 != navigator.appName.indexOf("Microsoft") ? window[e] : document[e]
		},
		a = function() {
			n('<object classid="clsid:D27CDB6E-AE6D-11cf-96B8-444553540000" id="fanyivoice" width="0" height="0" codebase="http://download.macromedia.com/pub/shockwave/cabs/flash/swflash.cab"><param name="movie" value="scripts/fanyivoice.swf?onload=swfLoad&time=' + (new Date).getTime() + '"/><param name="quality" value="high"/><param name="bgcolor" value="#869ca7"/><param name="allowScriptAccess" value="sameDomain"/><embed src="scripts/fanyivoice.swf?onload=swfLoad&time=' + (new Date).getTime() + '" quality="high" bgcolor="#869ca7" width="0" height="0" name="fanyivoice" align="middle" play="true" loop="false" quality="high" allowScriptAccess="sameDomain" type="application/x-shockwave-flash" pluginspage="http://www.macromedia.com/go/getflashplayer"></embed></object>').appendTo("body"), n("#speech").click(function() {
				var t = e("/output/output").getTransResult().type,
					i = r.find(".translated_result .tgt").text();
				return u(l(i, t)), n(this).blur(), !1
			})
		},
		s = {
			ZH_CN2EN: "eng",
			ZH_CN2JA: "jap",
			ZH_CN2KR: "ko",
			ZH_CN2FR: "fr"
		},
		l = function(e, t) {
			var n = "http://tts.youdao.com/fanyivoice?word=",
				r = encodeURIComponent(e);
			return n + r + "&le=" + s[t] + "&keyfrom=fanyi.web.index"
		},
		u = function(e) {
			e && o("fanyivoice") && jsReady === !0 && o("fanyivoice").playVoice(e)
		},
		c = function() {
			var e = 0;
			return r.find(".translated_result .tgt").each(function() {
				e += n(this).text().length
			}), e
		};
	t.stopVoice = function() {
		o("fanyivoice") && o("fanyivoice").stopVoice && o("fanyivoice").stopVoice()
	}, t.showVoice = function(e) {
		n("#fanyivoice").length <= 0 && a(), ("ZH_CN2EN" == e || "ZH_CN2JA" == e || "ZH_CN2KR" == e || "ZH_CN2FR" == e) && c() <= 300 ? i.show() : i.hide()
	}, t.init = function() {
		window.swfReady = !1, window.jsReady = !1, window.isContainerReady = function() {
			return jsReady
		}, window.setSWFIsReady = function() {
			swfReady = !0
		}, window.setJSReady = function() {
			jsReady = !0
		}, setTimeout(function() {
			setJSReady()
		}, 200)
	}
});
define("/output/copy", function(e, t) {
	var n = e("/lib/jquery-1.6.2");
	e("/ZeroClipboard"), window.swfIsReady = function() {
		return "getTranslatedResult"
	}, window.getTranslatedResult = function() {
		return e("/log").clog("COPY_CLICK", "web.o.righttop"), n("#outputText .translated_result").html().replace(/<br\s+.*?>/gi, "\n").replace(/<\/p>/gi, "\r\n").replace(/<p.*?>/gi, "").replace(/&amp;/g, "&").replace(/&lt;/g, "<").replace(/&gt;/g, ">").replace(/&copy;/g, "©")
	};
	var r = function() {
		window.clipboardData && (window.clipboardData.clearData(), window.clipboardData.setData("Text", getTranslatedResult()) || alert("您的浏览器不支持直接复制功能，请手工复制或者查看帮助中的常见问题进行设置。"))
	};
	t.init = function() {
		if (!(n("#copyit").length > 0))
			if (n.browser.msie) {
				var e = n('<a id="copyit" href="#" title=""></a>').appendTo("#copyOutput").show();
				e.click(function() {
					return r(), n(".copy-notice").text("已复制"), n(this).text("").blur(), setTimeout(function() {
						n(".copy-notice").text("复制")
					}, 500), !1
				})
			} else {
				var t = new ZeroClipboard(document.getElementById("copyOutput"));
				t.on("ready", function(e) {
					t.on("copy", function(e) {
						e.clipboardData.setData("text/plain", getTranslatedResult()), n(".copy-notice").text("已复制"), setTimeout(function() {
							n(".copy-notice").text("复制")
						}, 500)
					}), t.on("_mouseover", function(e) {
						n(".copynotice").css("display", "block")
					})
				})
			}
	}
});
define("/output/smartResult", function(e, t) {
	var n = e("/lib/jquery-1.6.2"),
		r = n("#outputText"),
		i = function(e, t) {
			t = t ? t : function(e) {
				return e
			};
			var r = "";
			return n.each(e, function(e, n) {
				t(n).length > 0 && (r += "<p>" + t(n) + "</p>")
			}), r
		},
		o = {
			1: function(t, n) {
				return i(t) + '<a class="smartresult_more" id="resultMore" href="http://dict.youdao.com/search?q=' + encodeURI(n.replace("&", "%26")) + '&keyfrom=fanyi.smartResult" target="_blank">' + e("/constData").smartText[1] + "&raquo;</a>"
			},
			2: function(e) {
				return i(e)
			},
			3: function(t, n, r) {
				return i(t, function(e) {
					return ' <a href="' + e.url + '" target="_blank">' + e.from + "</a><span>" + e.content + "</span>"
				}) + '<a class="smartresult_more" id="resultMore" href="' + r + '" target="_blank">' + e("/constData").smartText[3] + "&raquo;</a>"
			}
		},
		a = function(t, n, r, i) {
			var a = '<div class="smart_result">';
			return a += '<div class="smart_src_title">' + e("/constData").smartSRC[t] + "</div>", a += o[t](n, r, i), a += "</div>"
		};
	t.show = function(t) {
		return t = e("/output/output").getTransResult(), t.smartResult ? (n(".read-mode").hide(), void r.append(a(t.smartResult.type, t.smartResult.entries, t.translateResult[0][0].src, t.smartResult.aurl))) : void n(".read-mode").show()
	}
});
define("/output/typo", function(e, t) {
	var n = e("/lib/jquery-1.6.2"),
		r = n("#transForm .typo-suggest"),
		i = function(e) {
			return e.replace(/&gt;/g, ">").replace(/&lt;/g, "<").replace(/&amp;/g, "&").replace(/&quot;/g, '"').replace(/&apos;/g, "&#39;").replace(/&#39;/g, "'")
		},
		o = function(t) {
			for (var n = "", i = t.length, o = 0; i > o; o++) n += (0 !== o ? " " : "") + t[o];
			r.find(".spell-corrected").html(n.replace(/&apos;/g, "&#39;").replace(/&#39;/g, "'")), window.setTimeout(function() {
				r.show(function() {
					e("/log").clog("SHOW_TYPO", "web.i.leftbottom")
				})
			}, 1e3)
		},
		a = function() {
			for (var t = e("/output/output").getTransResult().typoResult, o = [], a = t.length, s = 0; a > s; s++) o.push(t[s].replace(/<b\>|<\/b\>/g, ""));
			r.hide(), n("#inputText").val(i(o.join("\n"))), n("#translateBtn").click()
		};
	t.hide = function() {
		r.hide()
	}, t.show = function(e) {
		e ? o(e) : t.hide()
	}, t.init = function() {
		n("#transForm .spell-corrected").bind("click.fillAndSearch", function() {
			return e("/log").clog("CLICK_TYPO", "web.i.leftbottom"), a(), !1
		})
	}
});
define("/output/readingMode", function(e, t) {
	var n = e("/lib/jquery-1.6.2"),
		r = n("#inputText"),
		i = n("#outputText"),
		o = n("#main"),
		a = n(".sponsor-content"),
		s = (n("#compare"), n(".compare-mode")),
		l = function(n) {
			return o.removeClass("reading-mode"), o.css("width", "auto"), a.removeClass("reading-mode"), r.scrollTop(n.inputScrollTop), i.scrollTop(n.outputScrollTop), s.hasClass("compare-enable") != n.revertStatus ? (e("/output/bilingualControl").changeMutiMode(), n.revertStatus = !1) : t.initReadingModeBtn(), n
		},
		u = function(t) {
			return n("body").addClass("full-mode"), t.inputScrollTop = r.scrollTop(), t.outputScrollTop = i.scrollTop(), o.addClass("reading-mode"), o.css("width", "1000px"), a.addClass("reading-mode"), t.revertStatus = s.hasClass("compare-enable"), s.hasClass("compare-enable") || e("/output/bilingualControl").changeMutiMode(), t
		};
	t.initReadingMode = function() {
		var t = {
			inputScrollTop: 0,
			outputScrollTop: 0,
			revertStatus: !1
		};
		n(".read-mode .open-reading-mode").click(function() {
			return e("/log").clog("CLICK_READING_MODE_BTN", "web.o.righttop"), t = u(t), !1
		}), n(".read-mode .close-reading-mode").click(function(e) {
			return n("body").removeClass("full-mode"), t = l(t), !1
		})
	}, t.initReadingModeBtn = function() {
		n("#outputText .translated_result").height() > n("#outputText").height() ? (e("/log").clog("SHOW_READING_MODE_BTN", "web.o.righttop"), n("#translated .read-mode").addClass("show-reading-mode")) : n("#translated .read-mode").removeClass("show-reading-mode")
	}
});
define("/output/queryClassifyAd", function(e, t) {
	var n = e("/lib/jquery-1.6.2"),
		r = function(e) {
			var t = i(e),
				r = "";
			if (1 == t) {
				var a = "fanyi.output.bottom",
					s = o(e.url, a);
				r = '<a target="_blank" style="color: red" href="' + s + '"><font color="red">' + e.slogan + "</font></a>", n(".suggest").css("display", "block")
			} else n(".suggest").css("display", "none");
			n("#suggestYou").html(r)
		},
		i = function(e) {
			return "undefined" == typeof e || null == e ? !1 : "ARTICLE" == e.clas || "ABSTRACT" == e.clas
		},
		o = function(e, t) {
			var n = null;
			return n = "undefined" == typeof e || null == e || 0 == e.length ? "http://f.youdao.com/" : e, n = "/" == n.charAt(n.length - 1) ? n + "?vendor=" + t : -1 == n.indexOf("?") && -1 == n.indexOf("&") ? n + "/?vendor=" + t : n + "&vendor=" + t
		};
	t.show = function(e) {
		r(e)
	}
});
define("/output/result", function(e, t) {
	var n, r = e("/lib/jquery-1.6.2"),
		i = r("#outputText"),
		o = function() {
			var e = i.find(".tgt");
			i.removeClass("small_font"), (e.size() >= 2 || e.height() > 60) && i.addClass("small_font")
		},
		a = function(e) {
			return "EN" === e || "FR" === e || "RU" === e || "SP" === e
		},
		s = function(e, t, n) {
			var i = "";
			return r.each(e, function(e, r) {
				i += r[t].replace(/</g, "&lt;").replace(/>/g, "&gt;") + n
			}), i
		},
		l = function(t, n, i) {
			var o = i.split("2")[0],
				l = i.split("2")[1],
				u = a(l) ? " " : "",
				c = a(o) ? " " : "",
				d = '<div class="translated_result">';
			return r.each(t, function(t, r) {
				e("/output/bilingualControl").isMutiMode() && n && (d += '<p class="src">' + s(r, "src", c) + "</p>"), d += '<p class="tgt">' + s(r, "tgt", u) + "</p>"
			}), d += "</div>"
		},
		u = function() {
			var t = e("/constData").LANGUAGETYPES[n.type];
			t ? "自动检测语言" == r(".btn_text").text() || -1 != r("#outputLang").text().indexOf("检测到") ? (r(".btn_text").html(""), r("#outputLang").html("检测到：" + t)) : (r(".btn_text").html(""), r("#outputLang").html(t)) : (r(".btn_text").html("检测到：暂不支持的语言"), r("#outputLang").html("")), r("#main").addClass("show-translate"), e("/example").hide(), i.html(l(n.translateResult, !n.smartResult, n.type)), r("#translated").show(), e("/output/voice").showVoice(n.type)
		};
	t.show = function() {
		n = e("/output/output").getTransResult(), e("/output/voice").stopVoice(), e("/output/copy").init(), u(), e("/output/smartResult").show(), o(), e("/output/typo").show(n.typoResult), e("/output/readingMode").initReadingModeBtn(), e("/tips").transError(n.errorCode), e("/output/queryClassifyAd").show(n.qClfResult), r("#inputText").focus()
	}
});
define("/output/output", function(e, t) {
	e("/lib/jquery-extension");
	var n, r, i, o = e("/lib/jquery-1.6.2"),
		a = e("/utils"),
		s = function() {
			var e = !1;
			o(document).bind("keydown", function(t) {
				e && t.ctrlKey && 65 == t.keyCode && (o("#outputText .translated_result").selectNodeText(), t.preventDefault())
			}), o("#outputText").click(function(t) {
				e = !0, t.stopPropagation()
			}), o(document).click(function() {
				e = !1
			})
		},
		l = function(t, r) {
			n = r;
			d().type;
			e("/output/result").show()
		},
		u = function(e) {
			l(e, i), i = null
		},
		c = function(t) {
			__rl_event("translate_text");
			var n = "translate?smartresult=dict&smartresult=rule&smartresult=ugc&sessionFrom=" + r;
			o.ajax({
				type: "POST",
				contentType: "application/x-www-form-urlencoded; charset=UTF-8",
				url: n,
				data: t,
				dataType: "json",
				success: function(e) {
					l(t, e)
				},
				error: function(t) {
					e("/tips").transRequestError(t)
				}
			})
		},
		d = function() {
			return n
		};
	t.getTransResult = d, t.translate = function(e) {
		i ? u(e) : c(e)
	}, t.clear = function() {
		o("#translated").fadeOut(), o("#main").removeClass("show-translate"), e("/output/voice").stopVoice()
	}, t.init = function(t, n) {
		i = t, r = n, e("/output/bilingualControl").init(), s(), e("/output/typo").init(), e("/output/readingMode").initReadingMode(), "true" === a.cookie("YOUDAO_FANYI_REALTIME_CLOSE") && (o(".close-realtime-mode").html("开启即时翻译"), global.realTime = !1), o(".close-realtime-mode").click(function() {
			return global.realTime ? (o(this).html("开启即时翻译"), a.cookie("YOUDAO_FANYI_REALTIME_CLOSE", "true"), e("/log").clog("FY_REALTIME_CLOSE", ""), global.realTime = !1) : (o(this).html("关闭即时翻译"), e("/log").clog("FY_REALTIME_OPEN", ""), a.cookie("YOUDAO_FANYI_REALTIME_CLOSE", "false"), global.realTime = !0), !1
		}), void 0 !== window.YoudaoSelector && f();
		var l = o("#addons");
		o("#resultScore").hover(function() {
			l.css("display", "block")
		}, function() {
			l.css("display", "none")
		}), l.bind("mouseover", function() {
			l.css("display", "block")
		}), l.bind("mouseout", function() {
			l.css("display", "none")
		}), o("#speech").hover(function() {
			o(".speechnotice").css("display", "block")
		}, function() {
			o(".speechnotice").css("display", "none")
		}), o(".actions").hover(function() {
			o(".copynotice").css("display", "block")
		}, function() {
			o(".copynotice").css("display", "none")
		}), o(".open-reading-mode").hover(function() {
			o(".opennotice").css("display", "block")
		}, function() {
			o(".opennotice").css("display", "none")
		}), o(".close-reading-mode").hover(function() {
			o(".closenotice").css("display", "block")
		}, function() {
			o(".closenotice").css("display", "none")
		})
	};
	var f = function() {
		o("#outputText").bind("click", YoudaoSelector.UI.hide);
		var t = o("#selectorSwitcher");
		t.bind("mouseover", function() {
			t.addClass("hover")
		}), t.bind("mouseout", function() {
			t.removeClass("hover")
		});
		var n = e("/log"),
			r = e("/utils");
		t.bind("click", function() {
			t.hasClass("selector-enable") ? (t.removeClass("selector-enable"), t.addClass("selector-disable"), YoudaoSelector.Config.select = "off", r.cookie("YOUDAO_FANYI_SELECTOR", "OFF"), n.clog("SELECT_OFF", "web.o.righttop")) : t.hasClass("selector-disable") && (t.removeClass("selector-disable"), t.addClass("selector-enable"), YoudaoSelector.Config.select = "on", r.cookie("YOUDAO_FANYI_SELECTOR", "ON"), n.clog("SELECT_ON", "web.o.righttop"))
		});
		var i = r.cookie("YOUDAO_FANYI_SELECTOR");
		"ON" === i ? (t.addClass("selector-enable"), YoudaoSelector.Config.select = "on") : (t.addClass("selector-disable"), YoudaoSelector.Config.select = "off"), t.css({
			display: "block"
		}), n.clog("SELECT_INIT_" + i, "web.o.righttop")
	}
});
define("/output/rankScore", function(e, t) {
	var n = e("/lib/jquery-1.6.2"),
		r = e("/utils"),
		i = function() {
			n("a.rank").unbind("click.rate"), n("a.rank").bind("click.rate", function() {
				return o(n(this)), !1
			})
		},
		o = function(t) {
			t.addClass("rateOn"), n("#rating_1").addClass("cover_rate"), n("#star .star_show").css("display", "none"), n("#star .star_result").css("display", "block"), n("#resultScore").css("background-position", "-1px -1px");
			var i = u();
			a({
				src: r.utf8_decode(s(n("#inputText").val())),
				tgt: r.utf8_decode(s(i)),
				score: t.data("score"),
				isSmartResult: l(),
				transType: n("#customSelectVal").val(),
				guessType: e("/constData").LANGUAGEDESCRIPTIONS[n("#outputLang").text()]
			})
		},
		a = function(e) {
			var t = "/score";
			n.ajax({
				type: "POST",
				contentType: "application/x-www-form-urlencoded; charset=UTF-8",
				url: t,
				data: e,
				dataType: "json"
			})
		},
		s = function(e) {
			var t = n.trim(e),
				r = t.replace(/\r\n/g, "\n").replace(/\n/g, "\\n");
			return r
		},
		l = function() {
			return n("#outputText .smart_result").length > 0 ? "true" : "false"
		},
		u = function() {
			return e("/output/bilingualControl").isMutiMode() ? n("#outputText .tgt").text() : n("#outputText .translated_result").text()
		},
		c = function() {
			n("#rating_1").removeClass("cover_rate"), n("#star .star_show").css("display", "inline"), n("#star .star_result").css("display", "none"), n("#resultScore").css("background-position", "0px -27px"), n("a.rank").removeClass("rateOn")
		};
	t.init = function() {
		c(), i()
	}
});
define("/input/advert", function(e, t) {
	var n = e("/lib/jquery-1.6.2"),
		r = function(e, t) {
			var r = document.head || document.getElementsByTagName("head")[0] || document.documentElement,
				i = document.createElement("script");
			i.type = "text/javascript", i.async = !0, i.src = e, i.charset = "utf-8";
			var o = !1;
			i.onload = i.onreadystatechange = function() {
				o || this.readyState && "loaded" !== this.readyState && "complete" !== this.readyState || (n.isFunction(t.successFunc) && t.successFunc(), i.onload = i.onreadystatechange = null, r && i.parentNode && r.removeChild(i), i = null, o = !0)
			}, i.onerror = function() {
				n.isFunction(t.errorFunc) && t.errorFunc()
			}, r.insertBefore(i, r.firstChild)
		};
	t.init = function() {
		var e = "http://shared.ydstatic.com/ead/js/dict_req_web_1.1.js";
		r(e, {
			successFunc: function() {
				"undefined" != typeof ADSupporter && (n("#advertisement").html(ADSupporter.getAdText("dw", "58", "301", "107861", "text_960_75", "960", "75")), n("#advertText").html(ADSupporter.getAdText("dws", "58", "0", "107861", "text_960_25", "960", "25")))
			}
		})
	}, t.reloadSP = function() {
		n("#advertisement iframe").attr("src", n("#advertisement iframe").attr("src") + "&clearCache=" + +new Date), n("#advertText iframe").attr("src", n("#advertText iframe").attr("src") + "&clearCache=" + +new Date)
	}
});
define("/input/input", function(e, t) {
	var n = e("/lib/jquery-1.6.2"),
		r = e("/tips"),
		i = e("/utils"),
		o = e("/atEntrance");
	e("/css3-mediaqueries");
	var a, s, l, u, c, d = 500,
		f = function() {
			var t = function(e) {
				return document.referrer.indexOf(e) >= 0
			};
			if (t("baidu.com") || t("google.com") || t("sogou.com") || t("soso.com") || t("so.com") || t("360.cn") || t("haosou.com")) {
				var o = '<strong>有道翻译</strong>&mdash;&mdash;中国最大最稳定的免费在线翻译 <img id="closeit" class="tip-close" src="http://shared.ydstatic.com/r/1.0/p/x.png" />';
				r.show(o), "ADD" !== i.cookie("YOUDAO_FANYI_FAV") && n("#closeit").before(e("/addFav").create("web.middle")), n("#closeit").click(function() {
					r.hide("all")
				})
			}
		},
		p = function() {
			return n.trim(a.val())
		},
		h = function() {
			a.val(""), s.fadeOut(function() {
				n("#entryList").show()
			}), e("/output/output").clear(), r.hide("allExceptWelcome"), e("/output/typo").hide();
			var t = n("#customSelectVal").val();
			o.showAtButton("", t), -1 != n("#outputLang").text().indexOf("检测到") && (n(".btn_text").html("自动检测语言"), n("#outputLang").html(""))
		},
		m = function(e) {
			a.val(p().replace(/\n/g, "").replace(/\r\n/g, ""));
			var t = a.val();
			location.href = "http://fanyi.youdao.com/WebpageTranslate?keyfrom=fanyi.web.index&url=" + encodeURIComponent(t) + "&type=" + n("#customSelectVal").val() + "&action=" + encodeURIComponent(e)
		},
		g = function(t) {
			e("/output/output").translate({
				type: n("#customSelectVal").val(),
				i: i.utf8_decode(a.val()),
				doctype: "json",
				xmlVersion: "1.8",
				keyfrom: "fanyi.web",
				ue: "UTF-8",
				action: t,
				typoResult: !0
			})
		},
		v = "",
		y = function(t, n) {
			e("/output/rankScore").init();
			var o = p();
			v = o, o.length > 0 ? (s.fadeIn(), i.isWeb(o) ? m(n) : g(n), r.hide("allExceptWelcome")) : (a.focus(), r.show("请输入内容"))
		},
		b = function() {
			var e = new Date;
			return e.getHours() < 8 || e.getHours() > 22 ? !0 : !1
		},
		x = function(e, n) {
			var r = t.getExpectCost(e),
				i = "0",
				o = 5,
				a = 0;
			if ("en" === n) {
				var s = Math.ceil(1.8 * r);
				s > 15 ? o = Math.ceil(1.4 * s) : o += s, a = s, o = parseInt(o), i = o
			} else r > 15 ? o = Math.ceil(1.4 * r) : o += r, a = r, i = o;
			if (!b()) return i;
			var l = new Date,
				u = l.getHours(),
				c = l.getMinutes();
			return e > 700 ? (23 == u ? i = 60 - c + o + 480 : 8 > u && (i = 60 * (7 - u) + 60 - c + o), i) : (23 == u ? i = parseInt(3 * o) : 0 == u || 1 == u ? i = parseInt(5 * o) : 2 == u ? i = parseInt(6 * o) : 3 == u || 4 == u || 5 == u ? i = parseInt(3 * o) : (6 == u || 7 == u) && (i = parseInt(2 * o)), i)
		};
	t.getExpectCost = function(e) {
		return Math.ceil(e > 1e3 ? (e - 1e3) / 8 + 100 : e / 10)
	}, t.init = function() {
		if (a = n("#inputText"), s = n("#clearInput"), u = n("#atBtn"), atNormalBtn = n(".atNormalBtn"), c = n("#atLink"), l = e("/sel").initSel(), "" !== p()) {
			var t = a.val();
			i.isWeb(t.toLowerCase()) || y(!1, "FY_BY_WITHQUERY")
		}
		a.focus(function() {
			"" !== p() ? s.show() : s.hide(), o.showAtButton(p(), n("#customSelectVal").val())
		}), f(), a.bind("textchange", function(t, r) {
			e("/output/typo").hide(), o.timero(function() {
				o.showAtButton(p(), n("#customSelectVal").val())
			}, 100).run()
		}), a.keyup(function(t) {
			a.val().length > 0 ? s.show() : h(), t.ctrlKey && 13 == t.keyCode || 13 == t.keyCode ? (t.ctrlKey && 13 == t.keyCode && y(!1, "FY_BY_CTRL_ENTER"), 13 == t.keyCode && y(!1, "FY_BY_ENTER"), e("/input/advert").reloadSP()) : global.realTime && v !== p() && i.timerProxy(function() {
				y(!0, "FY_BY_REALTIME")
			}, d), t.stopPropagation()
		}), s.click(function() {
			return e("/log").clog("CLEAN_CLICK", "web.i.clean"), h(), a.focus(), u.hover(function() {
				n("#atBtnhover").css("display", "none")
			}), !1
		}), n("#clearInput").hover(function() {
			n(".clearnotice").css("display", "block")
		}, function() {
			n(".clearnotice").css("display", "none")
		}), n("#translateBtn").click(function() {
			return y(!1, "FY_BY_CLICKBUTTON"), e("/input/advert").reloadSP(), !1
		}), u.click(function() {
			if ("button atDisableBtn" !== this.className) {
				var e = document.getElementById("mapForm"),
					t = document.getElementById("mapInput");
				t ? t.value = p() : (t = document.createElement("input"), t.id = "mapInput", t.type = "hidden", t.name = "text", t.value = p(), e.appendChild(t)), e.submit(), c.click()
			}
			return !1
		}), a.bind("textchange", function(e, t) {
			if ("" !== p()) {
				var r = /[\u4E00-\u9FA5]/i,
					i = r.test(p()) ? "other" : "en",
					o = x(p().length, i);
				u.hover(function() {
					n(".time-num").html(o + "分钟"), n("#atBtnhover").css("display", "block")
				}, function() {
					n("#atBtnhover").css("display", "none")
				})
			} else u.hover(u.hasClass("atDisableBtn") ? function() {
				n("#atBtnhover").css("display", "none")
			} : function() {
				n("#atBtnhover").css("display", "none")
			})
		}), a.bind("textchange", function(e, t) {
			"" == p() && n(".suggest").css("display", "none")
		}), n(".qrCode").hover(function() {
			n(".qrcodeHover").css("display", "block")
		}, function() {
			n(".qrcodeHover").css("display", "none")
		}), n(".dict-client").hover(function() {
			n(".dict-notice").css("display", "block")
		}, function() {
			n(".dict-notice").css("display", "none")
		}), e("/input/advert").init(), a.focus()
	}
});
define("/imageTest", function(e, t) {
	var n = e("/lib/jquery-1.6.2"),
		r = function() {
			var e = !1,
				t = n("#advertisement iframe").get(0);
			try {
				e = !!t.contentWindow.document.body;
				var r = n(".user-research a").attr("href");
				e && n(".user-research a").attr("href", r + "@image")
			} catch (i) {}
		};
	t.init = function() {
		setTimeout(function() {
			r()
		}, 1e3)
	}
});
define("/macapp_download", function(e, t) {
	function n() {
		var e, t, n, i, o;
		r(function() {
			function a(e, t, n) {
				var r = new Date;
				r.setDate(r.getDate() + n), document.cookie = e + "=" + t + "; path=" + window.location.pathname + ";expires=" + r
			}

			function s(e) {
				for (var t = document.cookie.split("; "), n = 0; n < t.length; n++) {
					var r = t[n].split("=");
					if (r[0] == e) return r[1]
				}
				return ""
			}

			function l(e) {
				return e / 1e3 / 60 / 60 / 24
			}
			var c = (function() {
				var e, t = location.search.replace(/^\?/, "").split("&"),
					n = {},
					r = 0;
				for (r = 0; r < t.length; r++) e = t[r], e && (e = e.split("="), n[e[0]] = void 0 === e[1] ? null : decodeURIComponent(e[1]));
				return n
			}(), '<div class="cpm-wrap"><div class="cpm"><div class="dict-cpm-mask"></div><div class="cpm-click dict-inline-block dict-vam"><div class="cpm-img"><a class="download-btn download-btn-now clog-js" data-clog="click-mac-download" href="http://c.youdao.com/dict/download.html?url=http://codown.youdao.com/cidian/download/MacDict_unsilent2.dmg target="_blank"></a><a class="download-btn download-btn-appstore clog-js" data-clog="click-mac-appstore" href="http://c.youdao.com/dict/download.html?url=https://itunes.apple.com/cn/app/you-dao-ci-dian/id491854842" target="_blank"></a></div><a class="cpm-close clog-js" data-clog="click-mac-close" href="javascript:;"></a></div><div class="cpm-height dict-inline-block dict-vam"></div></div></div>');
			r("body").append(c);
			var u, d = navigator.platform,
				f = r(".cpm-wrap"),
				p = 0,
				h = document.referrer,
				m = 1;
			switch (m) {
				case 1:
					u = "dict_mac_cpm_show1", e = "dict-cpm1", t = "http://c.youdao.com/dict/download.html?url=http://codown.youdao.com/cidian/download/MacDict_unsilent5.dmg&vendor=websuggest.fy", n = "http://c.youdao.com/dict/download.html?url=https://itunes.apple.com/cn/app/you-dao-ci-dian/id491854842&vendor=websuggest.fy", i = "dictweb_mac_download1", o = "dictweb_mac_appstore1";
					break;
				case 2:
					u = "dict_mac_cpm_show2", e = "dict-cpm2", t = "http://c.youdao.com/dict/download.html?url=http://codown.youdao.com/cidian/download/MacDict_unsilent5.dmg&vendor=websuggest.fy", n = "http://c.youdao.com/dict/download.html?url=https://itunes.apple.com/cn/app/you-dao-ci-dian/id491854842&vendor=websuggest.fy", i = "dictweb_mac_download2", o = "dictweb_mac_appstore2"
			}
			r(".cpm-click .cpm-img").addClass(e), r(".cpm-click .cpm-img .download-btn-now").attr("href", t), r(".cpm-click .cpm-img .download-btn-now").attr("data-rlog", i), r(".cpm-click .cpm-img .download-btn-appstore").attr("href", n), r(".cpm-click .cpm-img .download-btn-appstore").attr("data-rlog", o), (h.indexOf("so.com") > -1 || h.indexOf("baidu.com") > -1 || h.indexOf("sogou.com") > -1 || h.indexOf("youdao.com") > -1 || h.indexOf("haosou.com") > -1) && d.indexOf("Mac") > -1 && (p = parseInt(s("_dict_cpm_show") || 0, 10), -1 != p && l((new Date).getTime() - p) > 7 ? (f.show(), p = +new Date, a("_dict_cpm_show", p, 7), r(".download-btn").click(function() {
				a("_dict_cpm_show", -1, 365)
			}), r(".cpm-close").click(function() {
				var e = parseInt(s("_dict_cpm_close") || 0, 10);
				e++, e >= 3 ? a("_dict_cpm_show", -1, 365) : a("_dict_cpm_close", e, 365), f.hide()
			})) : f.hide())
		})
	}
	var r = e("/lib/jquery-1.6.2");
	t.init = n
});
define("/fanyi", function(e) {
	e("/sel"), e("/ZeroClipboard"), e("/css3-mediaqueries"), e("/log").init(), e("/navigateBar"), e("/tips").init(), e("/example").init(), e("/output/output").init(global.translatedJson, global.sessionFrom), e("/input/input").init(), e("/imageTest").init(), e("/output/voice").init(), e("/macapp_download").init()
});
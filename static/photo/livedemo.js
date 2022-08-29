/*var _0x1003 = ["domain", "w1.com", "indexOf", "wuaisheji.com", "href", "http://w1.com"];
$ym = document[_0x1003[0]];
if ($ym[_0x1003[2]](_0x1003[1]) < 0 && $ym[_0x1003[2]](_0x1003[3]) < 0) {
	location[_0x1003[4]] = _0x1003[5]
};*/
(function(a, b, c) {
	function g(f) {
		var c = b.console;
		h[f] || (h[f] = !0, a.migrateWarnings.push(f), c && (c.warn && !a.migrateMute) && (c.warn("JQMIGRATE: " + f), a.migrateTrace && c.trace && c.trace()))
	}
	var h = {};
	a.migrateWarnings = [];
	a.fn.toggle = function(b, f) {
		if (!a.isFunction(b) || !a.isFunction(f)) return M.apply(this, arguments);
		g("jQuery.fn.toggle(handler, handler...) is deprecated");
		var c = arguments,
			h = b.guid || a.guid++,
			j = 0,
			k = function(f) {
				var g = (a._data(this, "lastToggle" + b.guid) || 0) % j;
				a._data(this, "lastToggle" + b.guid, g + 1);
				f.preventDefault();
				return c[g].apply(this, arguments) || !1
			};
		for (k.guid = h; j < c.length;) c[j++].guid = h;
		return this.click(k)
	}
})(jQuery, window);
if (0 < $("#demo-page").length) {
	$(function() {
		var a = document.querySelector && document.querySelector('meta[name="viewport"]'),
			b = navigator.userAgent,
			c = function() {
				a.content = "width=device-width, minimum-scale=0.25, maximum-scale=1.6"
			};
		a && (/iPhone|iPad/.test(b) && !/Opera Mini/.test(b)) && (a.content = "width=device-width, minimum-scale=1.0, maximum-scale=1.0", document.addEventListener("gesturestart", c, !1))
	});
	$(".hide-price").removeClass("hide-price");
	var calcHeight = function() {
			var a = $("#headerlivedemo").height();
			$("#advanced").hasClass("closed") ? $("#iframelive").height($(window).height() + 3) : $("#iframelive").height($(window).height() - a)
		};
	$(document).ready(function() {
		calcHeight()
	});
	$(window).resize(function() {
		calcHeight()
	}).load(function() {
		calcHeight()
	});
	$(function() {
		mobileCss = function() {
			if (navigator.userAgent.match(/iPhone/i) || navigator.userAgent.match(/iPod/i) || navigator.userAgent.match(/iPad/i) || navigator.userAgent.match(/Android/i) || navigator.userAgent.match(/Opera Mini/)) $("#dropdown.box-drop ul.list-drop li .popover").css("top", "-1000px").css("display", "none !important"), $("#qr").hide(), $("body").css({
				height: "auto"
			}), $("html").css({
				height: "auto"
			}), $("#headerlivedemo").addClass("mobile")
		};
		mobileCss()
	});
	0 < $(".is-responsive").length ? $(document).ready(function() {
		if (!navigator.userAgent.match(/iPhone/i) && !navigator.userAgent.match(/iPod/i) && !navigator.userAgent.match(/iPad/i) && !navigator.userAgent.match(/Android/i) && !navigator.userAgent.match(/Opera Mini/)) {
			var a = document.getElementById("frame");
			$("#responsivator").show();
			$("#responsivator .response").click(function() {
				$("#responsivator .response").removeClass("active");
				$(this).addClass("active");
				var b = $(this).attr("id");
				$("#iframelive").removeAttr("class").addClass(b);
				a.src = a.src
			});
			$(".responsive-block").show()
		}
	}) : $(".js-demo-upper-menu").addClass("js-no-responsive");
	$(function() {
		function a() {
			var a = $(".js-popupi-3").width(),
				b = $(".js-price").width(),
				c = $(".price-title").width(),
				b = b + c + 39;
			b > a && 500 < window.width && $(".js-popupi-3").width(b)
		}
		a();
		$(".js-offer-trigger").on("click", function() {
			a()
		});
		var b = 0 < $(".livemonstroid").length ? "-60px" : "-53px";
		$("#advanced").css({
			marginTop: "0px"
		}).removeClass("closed");
		$("#headerlivedemo .trigger").toggle(function() {
			$(".js-popup-content").hide();
			$(this).removeClass("icon-arrow-small-up").addClass("icon-arrow-small-down");
			$("#advanced").addClass("closed").animate({
				marginTop: b
			}, "fast", function() {
				calcHeight()
			});
			strCookies2 = $.cookie("panel2", null);
			strCookies = $.cookie("panel", "boo")
		}, function() {
			$(this).removeClass("icon-arrow-small-down").addClass("icon-arrow-small-up");
			$("#advanced").removeClass("closed").animate({
				marginTop: "0px"
			}, "fast", function() {
				calcHeight()
			});
			strCookies2 = $.cookie("panel2", "opened");
			strCookies = $.cookie("panel", null)
		});
		$(".js-popup-open").on("click ", function() {
			var a = $(this).attr("data-popupi"),
				a = $(".js-popup-content." + a);
			if (a.is(":visible")) return !1;
			if ("js-popupi-4" == $(this).data("popupi")) {
				$(".js-demo-flexslider").css("display", "block").animate({
					top: 60
				}, 400);
				var b = $("#headerlivedemo").data("template_id");
				$(".thumb-wr").each(function() {
					$(this).data("tid") == b && $(this).addClass("active")
				});
				$(this).addClass("active");
				return !1
			}
			$(this).addClass("active");
			a.fadeIn(100)
		});
		jQuery(function(a) {
			a(document).on("mouseup", function(b) {
				var c = a(".js-popup-content");
				if (!c.is(b.target) && 0 === c.has(b.target).length) {
					if (c.prev(".theme-choose").hasClass("active")) return a(".js-demo-flexslider").animate({
						top: -260
					}, 400, function() {
						a(this).css("display", "none")
					}), a(".js-popup-open").removeClass("active"), !1;
					c.fadeOut(100);
					a(".js-popup-open").removeClass("active")
				}
			})
		});
		$(".icon-close").mouseup(function() {
			$(".js-popup-content").fadeOut(100);
			$(".js-popup-open").removeClass("active")
		});
		$(".js-demo-list-li .js-drop").on("click", function() {
			var a = $(this).parent().attr("data-variant"),
				b, c = $(this).parent(),
				f = $(this).parent().data("price") || $(".js-price", c).data("price"),
				c = $(".js-price", c).data("price");
			b = $("#advanced").hasClass("livemonstroid") ? $(this).children(".recommended-wrap").text() : tm.stringUtils.bareText($(this));
			$(".js-popup-content").fadeOut(100);
			$(".js-popup-open").removeClass("active");
			$(".js-offer-trigger .price-title").text(b);
			b = appDic.require("language");
			var g = '<span class="new-price">' + b.formatMoney(f) + "</span> ";
			c != f && (g = '<span class="discount">' + b.formatMoney(c) + "</span> " + g);
			$(".js-offer-trigger .js-price").html(g);
			$(".js-demo-buy-button[data-variant=" + a + "]").removeClass("js-none").siblings().addClass("js-none");
			$(".js-popupi-3").css("width", "auto")
		});
		$(".js-drop").hover(function() {
			$(this).siblings(".popover").css("opacity", "1").css("z-index", "999")
		}, function() {
			$(".list-drop li .popover").css("opacity", "0").css("z-index", "-999")
		});
		$.browser.msie && $("body").addClass("ie");
		var c = navigator.userAgent.toLowerCase(),
			g = function(a) {
				return -1 != c.indexOf(a)
			},
			f = document.getElementsByTagName("html")[0],
			g = [!/opera|webtv/i.test(c) && /msie (\d)/.test(c) ? "ie ie" + RegExp.$1 : g("firefox/2") ? "gecko ff2" : g("firefox/3") ? "gecko ff3" : g("gecko/") ? "gecko" : g("opera/9") ? "opera opera9" : /opera (\d)/.test(c) ? "opera opera" + RegExp.$1 : g("konqueror") ? "konqueror" : g("safari/") ? "webkit safari" : g("mozilla/") ? "gecko" : "", g("x11") || g("linux") ? " linux" : g("mac") ? " mac" : g("win") ? " win" : ""].join(" ");
		if (f.className) {
			var h = f.className;
			f.className = h + (" " + g)
		} else f.className = g
	})
};
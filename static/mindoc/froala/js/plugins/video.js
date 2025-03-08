/*!
 * froala_editor v2.8.4 (https://www.froala.com/wysiwyg-editor)
 * License https://froala.com/wysiwyg-editor/terms/
 * Copyright 2014-2018 Froala Labs
 */

!
function(i) {
    "function" == typeof define && define.amd ? define(["jquery"], i) : "object" == typeof module && module.exports ? module.exports = function(e, t) {
        return t === undefined && (t = "undefined" != typeof window ? require("jquery") : require("jquery")(e)),
        i(t)
    }: i(window.jQuery)
} (function(Z) {
    Z.extend(Z.FE.POPUP_TEMPLATES, {
        "video.insert": "[_BUTTONS_][_BY_URL_LAYER_][_EMBED_LAYER_][_UPLOAD_LAYER_][_PROGRESS_BAR_]",
        "video.edit": "[_BUTTONS_]",
        "video.size": "[_BUTTONS_][_SIZE_LAYER_]"
    }),
    Z.extend(Z.FE.DEFAULTS, {
        videoAllowedTypes: ["mp4", "webm", "ogg"],
        videoAllowedProviders: [".*"],
        videoDefaultAlign: "center",
        videoDefaultDisplay: "block",
        videoDefaultWidth: 600,
        videoEditButtons: ["videoReplace", "videoRemove", "|", "videoDisplay", "videoAlign", "videoSize"],
        videoInsertButtons: ["videoBack", "|", "videoByURL", "videoEmbed", "videoUpload"],
        videoMaxSize: 52428800,
        videoMove: !0,
        videoResize: !0,
        videoSizeButtons: ["videoBack", "|"],
        videoSplitHTML: !1,
        videoTextNear: !0,
        videoUpload: !0,
        videoUploadMethod: "POST",
        videoUploadParam: "file",
        videoUploadParams: {},
        videoUploadToS3: !1,
        videoUploadURL: null
    }),
    Z.FE.VIDEO_PROVIDERS = [{
        test_regex: /^.*((youtu.be)|(youtube.com))\/((v\/)|(\/u\/\w\/)|(embed\/)|(watch\?))?\??v?=?([^#\&\?]*).*/,
        url_regex: /(?:https?:\/\/)?(?:www\.)?(?:m\.)?(?:youtube\.com|youtu\.be)\/(?:watch\?v=|embed\/)?([0-9a-zA-Z_\-]+)(.+)?/g,
        url_text: "https://www.youtube.com/embed/$1",
        html: '<iframe width="640" height="360" src="{url}?wmode=opaque" frameborder="0" allowfullscreen></iframe>',
        provider: "youtube"
    },
    {
        test_regex: /^.*(?:vimeo.com)\/(?:channels(\/\w+\/)?|groups\/*\/videos\/\u200b\d+\/|video\/|)(\d+)(?:$|\/|\?)/,
        url_regex: /(?:https?:\/\/)?(?:www\.|player\.)?vimeo.com\/(?:channels\/(?:\w+\/)?|groups\/(?:[^\/]*)\/videos\/|album\/(?:\d+)\/video\/|video\/|)(\d+)(?:[a-zA-Z0-9_\-]+)?(\/[a-zA-Z0-9_\-]+)?/i,
        url_text: "https://player.vimeo.com/video/$1",
        html: '<iframe width="640" height="360" src="{url}" frameborder="0" allowfullscreen></iframe>',
        provider: "vimeo"
    },
    {
        test_regex: /^.+(dailymotion.com|dai.ly)\/(video|hub)?\/?([^_]+)[^#]*(#video=([^_&]+))?/,
        url_regex: /(?:https?:\/\/)?(?:www\.)?(?:dailymotion\.com|dai\.ly)\/(?:video|hub)?\/?(.+)/g,
        url_text: "https://www.dailymotion.com/embed/video/$1",
        html: '<iframe width="640" height="360" src="{url}" frameborder="0" allowfullscreen></iframe>',
        provider: "dailymotion"
    },
    {
        test_regex: /^.+(screen.yahoo.com)\/[^_&]+/,
        url_regex: "",
        url_text: "",
        html: '<iframe width="640" height="360" src="{url}?format=embed" frameborder="0" allowfullscreen="true" mozallowfullscreen="true" webkitallowfullscreen="true" allowtransparency="true"></iframe>',
        provider: "yahoo"
    },
    {
        test_regex: /^.+(rutube.ru)\/[^_&]+/,
        url_regex: /(?:https?:\/\/)?(?:www\.)?(?:rutube\.ru)\/(?:video)?\/?(.+)/g,
        url_text: "https://rutube.ru/play/embed/$1",
        html: '<iframe width="640" height="360" src="{url}" frameborder="0" allowfullscreen="true" mozallowfullscreen="true" webkitallowfullscreen="true" allowtransparency="true"></iframe>',
        provider: "rutube"
    },
    {
        test_regex: /^(?:.+)vidyard.com\/(?:watch)?\/?([^.&/] + )\ / ?( ? :[ ^ _. & ] + ) ? /,url_regex:/ ^ ( ? :. + ) vidyard.com\ / ( ? :watch) ? \ / ?([ ^ . & /]+)\/?(?:[^_.&]+)?/g, url_text: "https://play.vidyard.com/$1", html: '<iframe width="640" height="360" src="{url}" frameborder="0" allowfullscreen></iframe>', provider: "vidyard"
    }],
    Z.FE.VIDEO_EMBED_REGEX = /^\W*((<iframe.*><\/iframe>)|(<embed.*>))\W*$/i,
    Z.FE.PLUGINS.video = function(v) {
        var a, f, p, u, o, i, l = "https://i.froala.com/upload",
        d = 2,
        c = 3,
        h = 4,
        g = 5,
        m = 6,
        r = {};
        function b() {
            var e = v.popups.get("video.insert");
            e.find(".fr-video-by-url-layer input").val("").trigger("change");
            var t = e.find(".fr-video-embed-layer textarea");
            t.val("").trigger("change"),
            (t = e.find(".fr-video-upload-layer input")).val("").trigger("change")
        }
        function s() {
            var e = v.popups.get("video.edit");
            if (e || (e = function() {
                var e = "";
                if (0 < v.opts.videoEditButtons.length) {
                    e += '<div class="fr-buttons">',
                    e += v.button.buildList(v.opts.videoEditButtons);
                    var t = {
                        buttons: e += "</div>"
                    },
                    i = v.popups.create("video.edit", t);
                    return v.events.$on(v.$wp, "scroll.video-edit",
                    function() {
                        u && v.popups.isVisible("video.edit") && (v.events.disableBlur(), x(u))
                    }),
                    i
                }
                return ! 1
            } ()), e) {
                v.popups.setContainer("video.edit", v.$sc),
                v.popups.refresh("video.edit");
                var t = u.find("iframe, embed, video"),
                i = t.offset().left + t.outerWidth() / 2,
                o = t.offset().top + t.outerHeight();
                v.popups.show("video.edit", i, o, t.outerHeight())
            }
        }
        function n(e) {
            if (e) return v.popups.onRefresh("video.insert", b),
            v.popups.onHide("image.insert", j),
            !0;
            var t = "";
            v.opts.videoUpload || v.opts.videoInsertButtons.splice(v.opts.videoInsertButtons.indexOf("videoUpload"), 1),
            1 < v.opts.videoInsertButtons.length && (t = '<div class="fr-buttons">' + v.button.buildList(v.opts.videoInsertButtons) + "</div>");
            var i, o = "",
            r = v.opts.videoInsertButtons.indexOf("videoUpload"),
            s = v.opts.videoInsertButtons.indexOf("videoByURL"),
            n = v.opts.videoInsertButtons.indexOf("videoEmbed");
            0 <= s && (i = " fr-active", (r < s && 0 <= r || n < s && 0 <= n) && (i = ""), o = '<div class="fr-video-by-url-layer fr-layer' + i + '" id="fr-video-by-url-layer-' + v.id + '"><div class="fr-input-line"><input id="fr-video-by-url-layer-text-' + v.id + '" type="text" placeholder="' + v.language.translate("Paste in a video URL") + '" tabIndex="1" aria-required="true"></div><div class="fr-action-buttons"><button type="button" class="fr-command fr-submit" data-cmd="videoInsertByURL" tabIndex="2" role="button">' + v.language.translate("Insert") + "</button></div></div>");
            var a = "";
            0 <= n && (i = " fr-active", (r < n && 0 <= r || s < n && 0 <= s) && (i = ""), a = '<div class="fr-video-embed-layer fr-layer' + i + '" id="fr-video-embed-layer-' + v.id + '"><div class="fr-input-line"><textarea id="fr-video-embed-layer-text' + v.id + '" type="text" placeholder="' + v.language.translate("Embedded Code") + '" tabIndex="1" aria-required="true" rows="5"></textarea></div><div class="fr-action-buttons"><button type="button" class="fr-command fr-submit" data-cmd="videoInsertEmbed" tabIndex="2" role="button">' + v.language.translate("Insert") + "</button></div></div>");
            var d = "";
            0 <= r && (i = " fr-active", (n < r && 0 <= n || s < r && 0 <= s) && (i = ""), d = '<div class="fr-video-upload-layer fr-layer' + i + '" id="fr-video-upload-layer-' + v.id + '"><strong>' + v.language.translate("Drop video") + "</strong><br>(" + v.language.translate("or click") + ')<div class="fr-form"><input type="file" accept="video/' + v.opts.videoAllowedTypes.join(", video/").toLowerCase() + '" tabIndex="-1" aria-labelledby="fr-video-upload-layer-' + v.id + '" role="button"></div></div>');
            var l = {
                buttons: t,
                by_url_layer: o,
                embed_layer: a,
                upload_layer: d,
                progress_bar: '<div class="fr-video-progress-bar-layer fr-layer"><h3 tabIndex="-1" class="fr-message">Uploading</h3><div class="fr-loader"><span class="fr-progress"></span></div><div class="fr-action-buttons"><button type="button" class="fr-command fr-dismiss" data-cmd="videoDismissError" tabIndex="2" role="button">OK</button></div></div>'
            },
            f = v.popups.create("video.insert", l);
            return function(o) {
                v.events.$on(o, "dragover dragenter", ".fr-video-upload-layer",
                function() {
                    return Z(this).addClass("fr-drop"),
                    !1
                },
                !0),
                v.events.$on(o, "dragleave dragend", ".fr-video-upload-layer",
                function() {
                    return Z(this).removeClass("fr-drop"),
                    !1
                },
                !0),
                v.events.$on(o, "drop", ".fr-video-upload-layer",
                function(e) {
                    e.preventDefault(),
                    e.stopPropagation(),
                    Z(this).removeClass("fr-drop");
                    var t = e.originalEvent.dataTransfer;
                    if (t && t.files) {
                        var i = o.data("instance") || v;
                        i.events.disableBlur(),
                        i.video.upload(t.files),
                        i.events.enableBlur()
                    }
                },
                !0),
                v.helpers.isIOS() && v.events.$on(o, "touchstart", '.fr-video-upload-layer input[type="file"]',
                function() {
                    Z(this).trigger("click")
                },
                !0);
                v.events.$on(o, "change", '.fr-video-upload-layer input[type="file"]',
                function() {
                    if (this.files) {
                        var e = o.data("instance") || v;
                        e.events.disableBlur(),
                        o.find("input:focus").blur(),
                        e.events.enableBlur(),
                        e.video.upload(this.files)
                    }
                    Z(this).val("")
                },
                !0)
            } (f),
            f
        }
        function y(e) {
            v.events.focus(!0),
            v.selection.restore();
            var t = !1;
            u && (K(), t = !0),
            v.html.insert('<span contenteditable="false" draggable="true" class="fr-jiv fr-video">' + e + "</span>", !1, v.opts.videoSplitHTML),
            v.popups.hide("video.insert");
            var i = v.$el.find(".fr-jiv");
            i.removeClass("fr-jiv"),
            W(i, v.opts.videoDefaultDisplay, v.opts.videoDefaultAlign),
            i.toggleClass("fr-draggable", v.opts.videoMove),
            v.events.trigger(t ? "video.replaced": "video.inserted", [i])
        }
        function w() {
            var e = Z(this);
            v.popups.hide("video.insert"),
            e.removeClass("fr-uploading"),
            e.parent().next().is("br") && e.parent().next().remove(),
            x(e.parent()),
            v.events.trigger("video.loaded", [e.parent()])
        }
        function E(a, e, d, l, f) {
            v.edit.off(),
            _("Loading video"),
            e && (a = v.helpers.sanitizeURL(a));
            C("Loading video"),
            function() {
                var e, t;
                if (l) {
                    v.undo.canDo() || l.find("video").hasClass("fr-uploading") || v.undo.saveStep();
                    var i = l.find("video").data("fr-old-src"),
                    o = l.data("fr-replaced");
                    l.data("fr-replaced", !1),
                    v.$wp ? ((e = l.clone()).find("video").removeData("fr-old-src").removeClass("fr-uploading"), e.find("video").off("canplay"), i && l.find("video").attr("src", i), l.replaceWith(e)) : e = l;
                    for (var r = e.find("video").get(0).attributes, s = 0; s < r.length; s++) {
                        var n = r[s];
                        0 === n.nodeName.indexOf("data-") && e.find("video").removeAttr(n.nodeName)
                    }
                    if (void 0 !== d) for (t in d) d.hasOwnProperty(t) && "link" != t && e.find("video").attr("data-" + t, d[t]);
                    e.find("video").on("canplay", w),
                    e.find("video").attr("src", a),
                    v.edit.on(),
                    F(),
                    v.undo.saveStep(),
                    v.$el.blur(),
                    v.events.trigger(o ? "video.replaced": "video.inserted", [e, f])
                } else e = function(e, t, i) {
                    var o, r = "";
                    if (t && void 0 !== t) for (o in t) t.hasOwnProperty(o) && "link" != o && (r += " data-" + o + '="' + t[o] + '"');
                    var s = v.opts.videoDefaultWidth;
                    s && "auto" != s && (s += "px");
                    var n = Z('<span contenteditable="false" draggable="true" class="fr-video fr-dv' + v.opts.videoDefaultDisplay[0] + ("center" != v.opts.videoDefaultAlign ? " fr-fv" + v.opts.videoDefaultAlign[0] : "") + '"><video src="' + e + '" ' + r + (s ? ' style="width: ' + s + ';" ': "") + " controls>" + v.language.translate("Your browser does not support HTML5 video.") + "</video></span>");
                    n.toggleClass("fr-draggable", v.opts.videoMove),
                    v.edit.on(),
                    v.events.focus(!0),
                    v.selection.restore(),
                    v.undo.saveStep(),
                    v.opts.videoSplitHTML ? v.markers.split() : v.markers.insert(),
                    v.html.wrap();
                    var a = v.$el.find(".fr-marker");
                    return v.node.isLastSibling(a) && a.parent().hasClass("fr-deletable") && a.insertAfter(a.parent()),
                    a.replaceWith(n),
                    v.selection.clear(),
                    n.find("video").get(0).readyState > n.find("video").get(0).HAVE_FUTURE_DATA || v.helpers.isIOS() ? i.call(n.find("video").get(0)) : n.find("video").on("canplaythrough load", i),
                    n
                } (a, d, w),
                F(),
                v.undo.saveStep(),
                v.events.trigger("video.inserted", [e, f])
            } ()
        }
        function C(e) {
            var t = v.popups.get("video.insert");
            if (t || (t = n()), t.find(".fr-layer.fr-active").removeClass("fr-active").addClass("fr-pactive"), t.find(".fr-video-progress-bar-layer").addClass("fr-active"), t.find(".fr-buttons").hide(), u) {
                var i = u.find("video");
                v.popups.setContainer("video.insert", v.$sc);
                var o = i.offset().left + i.width() / 2,
                r = i.offset().top + i.height();
                v.popups.show("video.insert", o, r, i.outerHeight())
            }
            void 0 === e && _(v.language.translate("Uploading"), 0)
        }
        function A(e) {
            var t = v.popups.get("video.insert");
            if (t && (t.find(".fr-layer.fr-pactive").addClass("fr-active").removeClass("fr-pactive"), t.find(".fr-video-progress-bar-layer").removeClass("fr-active"), t.find(".fr-buttons").show(), e || v.$el.find("video.fr-error").length)) {
                if (v.events.focus(), v.$el.find("video.fr-error").length && (v.$el.find("video.fr-error").parent().remove(), v.undo.saveStep(), v.undo.run(), v.undo.dropRedo()), !v.$wp && u) {
                    var i = u;
                    z(!0),
                    v.selection.setAfter(i.find("video").get(0)),
                    v.selection.restore()
                }
                v.popups.hide("video.insert")
            }
        }
        function _(e, t) {
            var i = v.popups.get("video.insert");
            if (i) {
                var o = i.find(".fr-video-progress-bar-layer");
                o.find("h3").text(e + (t ? " " + t + "%": "")),
                o.removeClass("fr-error"),
                t ? (o.find("div").removeClass("fr-indeterminate"), o.find("div > span").css("width", t + "%")) : o.find("div").addClass("fr-indeterminate")
            }
        }
        function x(e) {
            O.call(e.get(0))
        }
        function R(e) {
            _("Loading video");
            var t = this.status,
            i = this.response,
            o = this.responseXML,
            r = this.responseText;
            try {
                if (v.opts.videoUploadToS3) if (201 == t) {
                    var s = function(e) {
                        try {
                            var t = Z(e).find("Location").text(),
                            i = Z(e).find("Key").text();
                            return ! 1 === v.events.trigger("video.uploadedToS3", [t, i, e], !0) ? (v.edit.on(), !1) : t
                        } catch(o) {
                            return N(h, e),
                            !1
                        }
                    } (o);
                    s && E(s, !1, [], e, i || o)
                } else N(h, i || o);
                else if (200 <= t && t < 300) {
                    var n = function(e) {
                        try {
                            if (!1 === v.events.trigger("video.uploaded", [e], !0)) return v.edit.on(),
                            !1;
                            var t = JSON.parse(e);
                            return t.link ? t: (N(d, e), !1)
                        } catch(i) {
                            return N(h, e),
                            !1
                        }
                    } (r);
                    n && E(n.link, !1, n, e, i || r)
                } else N(c, i || r)
            } catch(a) {
                N(h, i || r)
            }
        }
        function S() {
            N(h, this.response || this.responseText || this.responseXML)
        }
        function D(e) {
            if (e.lengthComputable) {
                var t = e.loaded / e.total * 100 | 0;
                _(v.language.translate("Uploading"), t)
            }
        }
        function U() {
            v.edit.on(),
            A(!0)
        }
        function I(e) {
            if (!v.core.sameInstance(p)) return ! 0;
            e.preventDefault(),
            e.stopPropagation();
            var t = e.pageX || (e.originalEvent.touches ? e.originalEvent.touches[0].pageX: null),
            i = e.pageY || (e.originalEvent.touches ? e.originalEvent.touches[0].pageY: null);
            if (!t || !i) return ! 1;
            if ("mousedown" == e.type) {
                var o = v.$oel.get(0).ownerDocument,
                r = o.defaultView || o.parentWindow,
                s = !1;
                try {
                    s = r.location != r.parent.location && !(r.$ && r.$.FE)
                } catch(n) {}
                s && r.frameElement && (t += v.helpers.getPX(Z(r.frameElement).offset().left) + r.frameElement.clientLeft, i = e.clientY + v.helpers.getPX(Z(r.frameElement).offset().top) + r.frameElement.clientTop)
            }
            v.undo.canDo() || v.undo.saveStep(),
            (f = Z(this)).data("start-x", t),
            f.data("start-y", i),
            a.show(),
            v.popups.hideAll(),
            P()
        }
        function B(e) {
            if (!v.core.sameInstance(p)) return ! 0;
            if (f) {
                e.preventDefault();
                var t = e.pageX || (e.originalEvent.touches ? e.originalEvent.touches[0].pageX: null),
                i = e.pageY || (e.originalEvent.touches ? e.originalEvent.touches[0].pageY: null);
                if (!t || !i) return ! 1;
                var o = f.data("start-x"),
                r = f.data("start-y");
                f.data("start-x", t),
                f.data("start-y", i);
                var s = t - o,
                n = i - r,
                a = u.find("iframe, embed, video"),
                d = a.width(),
                l = a.height(); (f.hasClass("fr-hnw") || f.hasClass("fr-hsw")) && (s = 0 - s),
                (f.hasClass("fr-hnw") || f.hasClass("fr-hne")) && (n = 0 - n),
                a.css("width", d + s),
                a.css("height", l + n),
                a.removeAttr("width"),
                a.removeAttr("height"),
                L()
            }
        }
        function $(e) {
            if (!v.core.sameInstance(p)) return ! 0;
            f && u && (e && e.stopPropagation(), f = null, a.hide(), L(), s(), v.undo.saveStep())
        }
        function t(e) {
            return '<div class="fr-handler fr-h' + e + '"></div>'
        }
        function k(e, t, i, o) {
            return e.pageX = t,
            e.pageY = t,
            I.call(this, e),
            e.pageX = e.pageX + i * Math.floor(Math.pow(1.1, o)),
            e.pageY = e.pageY + i * Math.floor(Math.pow(1.1, o)),
            B.call(this, e),
            $.call(this, e),
            ++o
        }
        function F() {
            var e, t = Array.prototype.slice.call(v.el.querySelectorAll("video, .fr-video > *")),
            i = [];
            for (e = 0; e < t.length; e++) i.push(t[e].getAttribute("src")),
            Z(t[e]).toggleClass("fr-draggable", v.opts.videoMove),
            "" === t[e].getAttribute("class") && t[e].removeAttribute("class"),
            "" === t[e].getAttribute("style") && t[e].removeAttribute("style");
            if (o) for (e = 0; e < o.length; e++) i.indexOf(o[e].getAttribute("src")) < 0 && v.events.trigger("video.removed", [Z(o[e])]);
            o = t
        }
        function L() {
            p ||
            function() {
                var e;
                if (v.shared.$video_resizer ? (p = v.shared.$video_resizer, a = v.shared.$vid_overlay, v.events.on("destroy",
                function() {
                    p.removeClass("fr-active").appendTo(Z("body:first"))
                },
                !0)) : (v.shared.$video_resizer = Z('<div class="fr-video-resizer"></div>'), p = v.shared.$video_resizer, v.events.$on(p, "mousedown",
                function(e) {
                    e.stopPropagation()
                },
                !0), v.opts.videoResize && (p.append(t("nw") + t("ne") + t("sw") + t("se")), v.shared.$vid_overlay = Z('<div class="fr-video-overlay"></div>'), a = v.shared.$vid_overlay, e = p.get(0).ownerDocument, Z(e).find("body:first").append(a))), v.events.on("shared.destroy",
                function() {
                    p.html("").removeData().remove(),
                    p = null,
                    v.opts.videoResize && (a.remove(), a = null)
                },
                !0), v.helpers.isMobile() || v.events.$on(Z(v.o_win), "resize.video",
                function() {
                    z(!0)
                }), v.opts.videoResize) {
                    e = p.get(0).ownerDocument,
                    v.events.$on(p, v._mousedown, ".fr-handler", I),
                    v.events.$on(Z(e), v._mousemove, B),
                    v.events.$on(Z(e.defaultView || e.parentWindow), v._mouseup, $),
                    v.events.$on(a, "mouseleave", $);
                    var o = 1,
                    r = null,
                    s = 0;
                    v.events.on("keydown",
                    function(e) {
                        if (u) {
                            var t = -1 != navigator.userAgent.indexOf("Mac OS X") ? e.metaKey: e.ctrlKey,
                            i = e.which; (i !== r || 200 < e.timeStamp - s) && (o = 1),
                            (i == Z.FE.KEYCODE.EQUALS || v.browser.mozilla && i == Z.FE.KEYCODE.FF_EQUALS) && t && !e.altKey ? o = k.call(this, e, 1, 1, o) : (i == Z.FE.KEYCODE.HYPHEN || v.browser.mozilla && i == Z.FE.KEYCODE.FF_HYPHEN) && t && !e.altKey && (o = k.call(this, e, 2, -1, o)),
                            r = i,
                            s = e.timeStamp
                        }
                    }),
                    v.events.on("keyup",
                    function() {
                        o = 1
                    })
                }
            } (),
            (v.$wp || v.$sc).append(p),
            p.data("instance", v);
            var e = u.find("iframe, embed, video");
            p.css("top", (v.opts.iframe ? e.offset().top - 1 : e.offset().top - v.$wp.offset().top - 1) + v.$wp.scrollTop()).css("left", (v.opts.iframe ? e.offset().left - 1 : e.offset().left - v.$wp.offset().left - 1) + v.$wp.scrollLeft()).css("width", e.get(0).getBoundingClientRect().width).css("height", e.get(0).getBoundingClientRect().height).addClass("fr-active")
        }
        function O(e) {
            if (e && "touchend" == e.type && i) return ! 0;
            if (e && v.edit.isDisabled()) return e.stopPropagation(),
            e.preventDefault(),
            !1;
            if (v.edit.isDisabled()) return ! 1;
            for (var t = 0; t < Z.FE.INSTANCES.length; t++) Z.FE.INSTANCES[t] != v && Z.FE.INSTANCES[t].events.trigger("video.hideResizer");
            v.toolbar.disable(),
            v.helpers.isMobile() && (v.events.disableBlur(), v.$el.blur(), v.events.enableBlur()),
            v.$el.find(".fr-video.fr-active").removeClass("fr-active"),
            (u = Z(this)).addClass("fr-active"),
            v.opts.iframe && v.size.syncIframe(),
            G(),
            L(),
            s(),
            v.selection.clear(),
            v.button.bulkRefresh(),
            v.events.trigger("image.hideResizer")
        }
        function z(e) {
            u && (v.shared.vid_exit_flag || !0 === e) && (p.removeClass("fr-active"), v.toolbar.enable(), u.removeClass("fr-active"), u = null, P())
        }
        function e() {
            v.shared.vid_exit_flag = !0
        }
        function P() {
            v.shared.vid_exit_flag = !1
        }
        function T(e) {
            var t = e.originalEvent.dataTransfer;
            if (t && t.files && t.files.length) {
                var i = t.files[0];
                if (i && i.type && -1 !== i.type.indexOf("video")) {
                    if (!v.opts.videoUpload) return e.preventDefault(),
                    e.stopPropagation(),
                    !1;
                    v.markers.remove(),
                    v.markers.insertAtPoint(e.originalEvent),
                    v.$el.find(".fr-marker").replaceWith(Z.FE.MARKERS),
                    v.popups.hideAll();
                    var o = v.popups.get("video.insert");
                    return o || (o = n()),
                    v.popups.setContainer("video.insert", v.$sc),
                    v.popups.show("video.insert", e.originalEvent.pageX, e.originalEvent.pageY),
                    C(),
                    0 <= v.opts.videoAllowedTypes.indexOf(i.type.replace(/video\//g, "")) ? M(t.files) : N(m),
                    e.preventDefault(),
                    e.stopPropagation(),
                    !1
                }
            }
        }
        function M(e) {
            if (void 0 !== e && 0 < e.length) {
                if (!1 === v.events.trigger("video.beforeUpload", [e])) return ! 1;
                var t, i = e[0];
                if (null === v.opts.videoUploadURL || v.opts.videoUploadURL == l) return a = i,
                (d = new FileReader).addEventListener("load",
                function() {
                    d.result;
                    for (var e = atob(d.result.split(",")[1]), t = [], i = 0; i < e.length; i++) t.push(e.charCodeAt(i));
                    E(window.URL.createObjectURL(new Blob([new Uint8Array(t)], {
                        type: a.type
                    })), !1, null, u)
                },
                !1),
                C(),
                d.readAsDataURL(a),
                !1;
                if (i.size > v.opts.videoMaxSize) return N(g),
                !1;
                if (v.opts.videoAllowedTypes.indexOf(i.type.replace(/video\//g, "")) < 0) return N(m),
                !1;
                if (v.drag_support.formdata && (t = v.drag_support.formdata ? new FormData: null), t) {
                    var o;
                    if (!1 !== v.opts.videoUploadToS3) for (o in t.append("key", v.opts.videoUploadToS3.keyStart + (new Date).getTime() + "-" + (i.name || "untitled")), t.append("success_action_status", "201"), t.append("X-Requested-With", "xhr"), t.append("Content-Type", i.type), v.opts.videoUploadToS3.params) v.opts.videoUploadToS3.params.hasOwnProperty(o) && t.append(o, v.opts.videoUploadToS3.params[o]);
                    for (o in v.opts.videoUploadParams) v.opts.videoUploadParams.hasOwnProperty(o) && t.append(o, v.opts.videoUploadParams[o]);
                    t.append(v.opts.videoUploadParam, i);
                    var r = v.opts.videoUploadURL;
                    v.opts.videoUploadToS3 && (r = v.opts.videoUploadToS3.uploadURL ? v.opts.videoUploadToS3.uploadURL: "https://" + v.opts.videoUploadToS3.region + ".amazonaws.com/" + v.opts.videoUploadToS3.bucket);
                    var s = v.core.getXHR(r, v.opts.videoUploadMethod);
                    s.onload = function() {
                        R.call(s, u)
                    },
                    s.onerror = S,
                    s.upload.onprogress = D,
                    s.onabort = U,
                    C(),
                    v.events.disableBlur(),
                    v.edit.off(),
                    v.events.enableBlur();
                    var n = v.popups.get("video.insert");
                    n && n.off("abortUpload").on("abortUpload",
                    function() {
                        4 != s.readyState && s.abort()
                    }),
                    s.send(t)
                }
            }
            var a, d
        }
        function N(e, t) {
            v.edit.on(),
            u && u.find("video").addClass("fr-error"),
            function(e) {
                C();
                var t = v.popups.get("video.insert").find(".fr-video-progress-bar-layer");
                t.addClass("fr-error");
                var i = t.find("h3");
                i.text(e),
                v.events.disableBlur(),
                i.focus()
            } (v.language.translate("Something went wrong. Please try again.")),
            v.events.trigger("video.error", [{
                code: e,
                message: r[e]
            },
            t])
        }
        function V() {
            if (u) {
                var e = v.popups.get("video.size"),
                t = u.find("iframe, embed, video");
                e.find('input[name="width"]').val(t.get(0).style.width || t.attr("width")).trigger("change"),
                e.find('input[name="height"]').val(t.get(0).style.height || t.attr("height")).trigger("change")
            }
        }
        function Y(e) {
            if (e) return v.popups.onRefresh("video.size", V),
            !0;
            var t = {
                buttons: '<div class="fr-buttons">' + v.button.buildList(v.opts.videoSizeButtons) + "</div>",
                size_layer: '<div class="fr-video-size-layer fr-layer fr-active" id="fr-video-size-layer-' + v.id + '"><div class="fr-video-group"><div class="fr-input-line"><input id="fr-video-size-layer-width-' + v.id + '" type="text" name="width" placeholder="' + v.language.translate("Width") + '" tabIndex="1"></div><div class="fr-input-line"><input id="fr-video-size-layer-height-' + v.id + '" type="text" name="height" placeholder="' + v.language.translate("Height") + '" tabIndex="1"></div></div><div class="fr-action-buttons"><button type="button" class="fr-command fr-submit" data-cmd="videoSetSize" tabIndex="2" role="button">' + v.language.translate("Update") + "</button></div></div>"
            },
            i = v.popups.create("video.size", t);
            return v.events.$on(v.$wp, "scroll",
            function() {
                u && v.popups.isVisible("video.size") && (v.events.disableBlur(), x(u))
            }),
            i
        }
        function H(e) {
            if (void 0 === e && (e = u), e) {
                if (e.hasClass("fr-fvl")) return "left";
                if (e.hasClass("fr-fvr")) return "right";
                if (e.hasClass("fr-dvb") || e.hasClass("fr-dvi")) return "center";
                if ("block" == e.css("display")) {
                    if ("left" == e.css("text-algin")) return "left";
                    if ("right" == e.css("text-align")) return "right"
                } else {
                    if ("left" == e.css("float")) return "left";
                    if ("right" == e.css("float")) return "right"
                }
            }
            return "center"
        }
        function X(e) {
            void 0 === e && (e = u);
            var t = e.css("float");
            return e.css("float", "none"),
            "block" == e.css("display") ? (e.css("float", ""), e.css("float") != t && e.css("float", t), "block") : (e.css("float", ""), e.css("float") != t && e.css("float", t), "inline")
        }
        function K() {
            if (u && !1 !== v.events.trigger("video.beforeRemove", [u])) {
                var e = u;
                v.popups.hideAll(),
                z(!0),
                v.selection.setBefore(e.get(0)) || v.selection.setAfter(e.get(0)),
                e.remove(),
                v.selection.restore(),
                v.html.fillEmptyBlocks(),
                v.events.trigger("video.removed", [e])
            }
        }
        function j() {
            A()
        }
        function W(e, t, i) { ! v.opts.htmlUntouched && v.opts.useClasses ? (e.removeClass("fr-fvl fr-fvr fr-dvb fr-dvi"), e.addClass("fr-fv" + i[0] + " fr-dv" + t[0])) : "inline" == t ? (e.css({
                display: "inline-block"
            }), "center" == i ? e.css({
                "float": "none"
            }) : "left" == i ? e.css({
                "float": "left"
            }) : e.css({
                "float": "right"
            })) : (e.css({
                display: "block",
                clear: "both"
            }), "left" == i ? e.css({
                textAlign: "left"
            }) : "right" == i ? e.css({
                textAlign: "right"
            }) : e.css({
                textAlign: "center"
            }))
        }
        function q() {
            v.$el.find("video").filter(function() {
                return 0 === Z(this).parents("span.fr-video").length
            }).wrap('<span class="fr-video" contenteditable="false"></span>'),
            v.$el.find("embed, iframe").filter(function() {
                if (v.browser.safari && this.getAttribute("src") && this.setAttribute("src", this.src), 0 < Z(this).parents("span.fr-video").length) return ! 1;
                for (var e = Z(this).attr("src"), t = 0; t < Z.FE.VIDEO_PROVIDERS.length; t++) {
                    var i = Z.FE.VIDEO_PROVIDERS[t];
                    if (i.test_regex.test(e) && new RegExp(v.opts.videoAllowedProviders.join("|")).test(i.provider)) return ! 0
                }
                return ! 1
            }).map(function() {
                return 0 === Z(this).parents("object").length ? this: Z(this).parents("object").get(0)
            }).wrap('<span class="fr-video" contenteditable="false"></span>');
            for (var e, t, i = v.$el.find("span.fr-video, video"), o = 0; o < i.length; o++) {
                var r = Z(i[o]); ! v.opts.htmlUntouched && v.opts.useClasses ? ((t = r).hasClass("fr-dvi") || t.hasClass("fr-dvb") || (t.addClass("fr-fv" + H(t)[0]), t.addClass("fr-dv" + X(t)[0])), v.opts.videoTextNear || r.removeClass("fr-dvi").addClass("fr-dvb")) : v.opts.htmlUntouched || v.opts.useClasses || (W(e = r, e.hasClass("fr-dvb") ? "block": e.hasClass("fr-dvi") ? "inline": null, e.hasClass("fr-fvl") ? "left": e.hasClass("fr-fvr") ? "right": H(e)), e.removeClass("fr-dvb fr-dvi fr-fvr fr-fvl"))
            }
            i.toggleClass("fr-draggable", v.opts.videoMove)
        }
        function G() {
            if (u) {
                v.selection.clear();
                var e = v.doc.createRange();
                e.selectNode(u.get(0)),
                v.selection.get().addRange(e)
            }
        }
        return r[1] = "Video cannot be loaded from the passed link.",
        r[d] = "No link in upload response.",
        r[c] = "Error during file upload.",
        r[h] = "Parsing response failed.",
        r[g] = "File is too large.",
        r[m] = "Video file type is invalid.",
        r[7] = "Files can be uploaded only to same domain in IE 8 and IE 9.",
        v.shared.vid_exit_flag = !1,
        {
            _init: function() {
                v.events.on("drop", T, !0),
                v.events.on("mousedown window.mousedown", e),
                v.events.on("window.touchmove", P),
                v.events.on("mouseup window.mouseup", z),
                v.events.on("commands.mousedown",
                function(e) {
                    0 < e.parents(".fr-toolbar").length && z()
                }),
                v.events.on("video.hideResizer commands.undo commands.redo element.dropped",
                function() {
                    z(!0)
                }),
                v.helpers.isMobile() && (v.events.$on(v.$el, "touchstart", "span.fr-video",
                function() {
                    i = !1
                }), v.events.$on(v.$el, "touchmove",
                function() {
                    i = !0
                })),
                v.events.on("html.set", q),
                q(),
                v.events.$on(v.$el, "mousedown", "span.fr-video",
                function(e) {
                    e.stopPropagation()
                }),
                v.events.$on(v.$el, "click touchend", "span.fr-video",
                function(e) {
                    if ("false" == Z(this).parents("[contenteditable]:not(.fr-element):not(.fr-img-caption):not(body):first").attr("contenteditable")) return ! 0;
                    O.call(this, e)
                }),
                v.events.on("keydown",
                function(e) {
                    var t = e.which;
                    return ! u || t != Z.FE.KEYCODE.BACKSPACE && t != Z.FE.KEYCODE.DELETE ? u && t == Z.FE.KEYCODE.ESC ? (z(!0), e.preventDefault(), !1) : u && t != Z.FE.KEYCODE.F10 && !v.keys.isBrowserAction(e) ? (e.preventDefault(), !1) : void 0 : (e.preventDefault(), K(), v.undo.saveStep(), !1)
                },
                !0),
                v.events.on("toolbar.esc",
                function() {
                    if (u) return v.events.disableBlur(),
                    v.events.focus(),
                    !1
                },
                !0),
                v.events.on("toolbar.focusEditor",
                function() {
                    if (u) return ! 1
                },
                !0),
                v.events.on("keydown",
                function() {
                    v.$el.find("span.fr-video:empty").remove()
                }),
                v.$wp && (F(), v.events.on("contentChanged", F)),
                n(!0),
                Y(!0)
            },
            showInsertPopup: function() {
                var e = v.$tb.find('.fr-command[data-cmd="insertVideo"]'),
                t = v.popups.get("video.insert");
                if (t || (t = n()), A(), !t.hasClass("fr-active")) if (v.popups.refresh("video.insert"), v.popups.setContainer("video.insert", v.$tb), e.is(":visible")) {
                    var i = e.offset().left + e.outerWidth() / 2,
                    o = e.offset().top + (v.opts.toolbarBottom ? 10 : e.outerHeight() - 10);
                    v.popups.show("video.insert", i, o, e.outerHeight())
                } else v.position.forSelection(t),
                v.popups.show("video.insert")
            },
            showLayer: function(e) {
                var t, i, o = v.popups.get("video.insert");
                if (!u && !v.opts.toolbarInline) {
                    var r = v.$tb.find('.fr-command[data-cmd="insertVideo"]');
                    t = r.offset().left + r.outerWidth() / 2,
                    i = r.offset().top + (v.opts.toolbarBottom ? 10 : r.outerHeight() - 10)
                }
                v.opts.toolbarInline && (i = o.offset().top - v.helpers.getPX(o.css("margin-top")), o.hasClass("fr-above") && (i += o.outerHeight())),
                o.find(".fr-layer").removeClass("fr-active"),
                o.find(".fr-" + e + "-layer").addClass("fr-active"),
                v.popups.show("video.insert", t, i, 0),
                v.accessibility.focusPopup(o)
            },
            refreshByURLButton: function(e) {
                v.popups.get("video.insert").find(".fr-video-by-url-layer").hasClass("fr-active") && e.addClass("fr-active").attr("aria-pressed", !0)
            },
            refreshEmbedButton: function(e) {
                v.popups.get("video.insert").find(".fr-video-embed-layer").hasClass("fr-active") && e.addClass("fr-active").attr("aria-pressed", !0)
            },
            refreshUploadButton: function(e) {
                v.popups.get("video.insert").find(".fr-video-upload-layer").hasClass("fr-active") && e.addClass("fr-active").attr("aria-pressed", !0)
            },
            upload: M,
            insertByURL: function(e) {
                void 0 === e && (e = (v.popups.get("video.insert").find('.fr-video-by-url-layer input[type="text"]').val() || "").trim());
                var t = null;
                if (/^http/.test(e) || (e = "https://" + e), v.helpers.isURL(e)) for (var i = 0; i < Z.FE.VIDEO_PROVIDERS.length; i++) {
                    var o = Z.FE.VIDEO_PROVIDERS[i];
                    if (o.test_regex.test(e) && new RegExp(v.opts.videoAllowedProviders.join("|")).test(o.provider)) {
                        t = e.replace(o.url_regex, o.url_text),
                        t = o.html.replace(/\{url\}/, t);
                        break
                    }
                }
                t ? y(t) : v.events.trigger("video.linkError", [e])
            },
            insertEmbed: function(e) {
                void 0 === e && (e = v.popups.get("video.insert").find(".fr-video-embed-layer textarea").val() || ""),
                0 !== e.length && Z.FE.VIDEO_EMBED_REGEX.test(e) ? y(e) : v.events.trigger("video.codeError", [e])
            },
            insert: y,
            align: function(e) {
                u.removeClass("fr-fvr fr-fvl"),
                !v.opts.htmlUntouched && v.opts.useClasses ? "left" == e ? u.addClass("fr-fvl") : "right" == e && u.addClass("fr-fvr") : W(u, X(), e),
                G(),
                L(),
                s(),
                v.selection.clear()
            },
            refreshAlign: function(e) {
                if (!u) return ! 1;
                e.find("> *:first").replaceWith(v.icon.create("video-align-" + H()))
            },
            refreshAlignOnShow: function(e, t) {
                u && t.find('.fr-command[data-param1="' + H() + '"]').addClass("fr-active").attr("aria-selected", !0)
            },
            display: function(e) {
                u.removeClass("fr-dvi fr-dvb"),
                !v.opts.htmlUntouched && v.opts.useClasses ? "inline" == e ? u.addClass("fr-dvi") : "block" == e && u.addClass("fr-dvb") : W(u, e, H()),
                G(),
                L(),
                s(),
                v.selection.clear()
            },
            refreshDisplayOnShow: function(e, t) {
                u && t.find('.fr-command[data-param1="' + X() + '"]').addClass("fr-active").attr("aria-selected", !0)
            },
            remove: K,
            hideProgressBar: A,
            showSizePopup: function() {
                var e = v.popups.get("video.size");
                e || (e = Y()),
                A(),
                v.popups.refresh("video.size"),
                v.popups.setContainer("video.size", v.$sc);
                var t = u.find("iframe, embed, video"),
                i = t.offset().left + t.width() / 2,
                o = t.offset().top + t.height();
                v.popups.show("video.size", i, o, t.height())
            },
            replace: function() {
                var e = v.popups.get("video.insert");
                e || (e = n()),
                v.popups.isVisible("video.insert") || (A(), v.popups.refresh("video.insert"), v.popups.setContainer("video.insert", v.$sc));
                var t = u.offset().left + u.width() / 2,
                i = u.offset().top + u.height();
                v.popups.show("video.insert", t, i, u.outerHeight())
            },
            back: function() {
                u ? (v.events.disableBlur(), u.trigger("click")) : (v.events.disableBlur(), v.selection.restore(), v.events.enableBlur(), v.popups.hide("video.insert"), v.toolbar.showInline())
            },
            setSize: function(e, t) {
                if (u) {
                    var i = v.popups.get("video.size"),
                    o = u.find("iframe, embed, video");
                    o.css("width", e || i.find('input[name="width"]').val()),
                    o.css("height", t || i.find('input[name="height"]').val()),
                    o.get(0).style.width && o.removeAttr("width"),
                    o.get(0).style.height && o.removeAttr("height"),
                    i.find("input:focus").blur(),
                    setTimeout(function() {
                        u.trigger("click")
                    },
                    v.helpers.isAndroid() ? 50 : 0)
                }
            },
            get: function() {
                return u
            }
        }
    },
    Z.FE.RegisterCommand("insertVideo", {
        title: "Insert Video",
        undo: !1,
        focus: !0,
        refreshAfterCallback: !1,
        popup: !0,
        callback: function() {
            this.popups.isVisible("video.insert") ? (this.$el.find(".fr-marker").length && (this.events.disableBlur(), this.selection.restore()), this.popups.hide("video.insert")) : this.video.showInsertPopup()
        },
        plugin: "video"
    }),
    Z.FE.DefineIcon("insertVideo", {
        NAME: "video-camera",
        FA5NAME: "camera"
    }),
    Z.FE.DefineIcon("videoByURL", {
        NAME: "link"
    }),
    Z.FE.RegisterCommand("videoByURL", {
        title: "By URL",
        undo: !1,
        focus: !1,
        toggle: !0,
        callback: function() {
            this.video.showLayer("video-by-url")
        },
        refresh: function(e) {
            this.video.refreshByURLButton(e)
        }
    }),
    Z.FE.DefineIcon("videoEmbed", {
        NAME: "code"
    }),
    Z.FE.RegisterCommand("videoEmbed", {
        title: "Embedded Code",
        undo: !1,
        focus: !1,
        toggle: !0,
        callback: function() {
            this.video.showLayer("video-embed")
        },
        refresh: function(e) {
            this.video.refreshEmbedButton(e)
        }
    }),
    Z.FE.DefineIcon("videoUpload", {
        NAME: "upload"
    }),
    Z.FE.RegisterCommand("videoUpload", {
        title: "Upload Video",
        undo: !1,
        focus: !1,
        toggle: !0,
        callback: function() {
            this.video.showLayer("video-upload")
        },
        refresh: function(e) {
            this.video.refreshUploadButton(e)
        }
    }),
    Z.FE.RegisterCommand("videoInsertByURL", {
        undo: !0,
        focus: !0,
        callback: function() {
            this.video.insertByURL()
        }
    }),
    Z.FE.RegisterCommand("videoInsertEmbed", {
        undo: !0,
        focus: !0,
        callback: function() {
            this.video.insertEmbed()
        }
    }),
    Z.FE.DefineIcon("videoDisplay", {
        NAME: "star"
    }),
    Z.FE.RegisterCommand("videoDisplay", {
        title: "Display",
        type: "dropdown",
        options: {
            inline: "Inline",
            block: "Break Text"
        },
        callback: function(e, t) {
            this.video.display(t)
        },
        refresh: function(e) {
            this.opts.videoTextNear || e.addClass("fr-hidden")
        },
        refreshOnShow: function(e, t) {
            this.video.refreshDisplayOnShow(e, t)
        }
    }),
    Z.FE.DefineIcon("video-align", {
        NAME: "align-left"
    }),
    Z.FE.DefineIcon("video-align-left", {
        NAME: "align-left"
    }),
    Z.FE.DefineIcon("video-align-right", {
        NAME: "align-right"
    }),
    Z.FE.DefineIcon("video-align-center", {
        NAME: "align-justify"
    }),
    Z.FE.DefineIcon("videoAlign", {
        NAME: "align-center"
    }),
    Z.FE.RegisterCommand("videoAlign", {
        type: "dropdown",
        title: "Align",
        options: {
            left: "Align Left",
            center: "None",
            right: "Align Right"
        },
        html: function() {
            var e = '<ul class="fr-dropdown-list" role="presentation">',
            t = Z.FE.COMMANDS.videoAlign.options;
            for (var i in t) t.hasOwnProperty(i) && (e += '<li role="presentation"><a class="fr-command fr-title" tabIndex="-1" role="option" data-cmd="videoAlign" data-param1="' + i + '" title="' + this.language.translate(t[i]) + '">' + this.icon.create("video-align-" + i) + '<span class="fr-sr-only">' + this.language.translate(t[i]) + "</span></a></li>");
            return e += "</ul>"
        },
        callback: function(e, t) {
            this.video.align(t)
        },
        refresh: function(e) {
            this.video.refreshAlign(e)
        },
        refreshOnShow: function(e, t) {
            this.video.refreshAlignOnShow(e, t)
        }
    }),
    Z.FE.DefineIcon("videoReplace", {
        NAME: "exchange",
        FA5NAME: "exchange-alt"
    }),
    Z.FE.RegisterCommand("videoReplace", {
        title: "Replace",
        undo: !1,
        focus: !1,
        popup: !0,
        refreshAfterCallback: !1,
        callback: function() {
            this.video.replace()
        }
    }),
    Z.FE.DefineIcon("videoRemove", {
        NAME: "trash"
    }),
    Z.FE.RegisterCommand("videoRemove", {
        title: "Remove",
        callback: function() {
            this.video.remove()
        }
    }),
    Z.FE.DefineIcon("videoSize", {
        NAME: "arrows-alt"
    }),
    Z.FE.RegisterCommand("videoSize", {
        undo: !1,
        focus: !1,
        popup: !0,
        title: "Change Size",
        callback: function() {
            this.video.showSizePopup()
        }
    }),
    Z.FE.DefineIcon("videoBack", {
        NAME: "arrow-left"
    }),
    Z.FE.RegisterCommand("videoBack", {
        title: "Back",
        undo: !1,
        focus: !1,
        back: !0,
        callback: function() {
            this.video.back()
        },
        refresh: function(e) {
            this.video.get() || this.opts.toolbarInline ? (e.removeClass("fr-hidden"), e.next(".fr-separator").removeClass("fr-hidden")) : (e.addClass("fr-hidden"), e.next(".fr-separator").addClass("fr-hidden"))
        }
    }),
    Z.FE.RegisterCommand("videoDismissError", {
        title: "OK",
        undo: !1,
        callback: function() {
            this.video.hideProgressBar(!0)
        }
    }),
    Z.FE.RegisterCommand("videoSetSize", {
        undo: !0,
        focus: !1,
        title: "Update",
        refreshAfterCallback: !1,
        callback: function() {
            this.video.setSize()
        }
    })
});
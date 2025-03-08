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
} (function(C) {
    C.extend(C.FE.POPUP_TEMPLATES, {
        "file.insert": "[_BUTTONS_][_UPLOAD_LAYER_][_PROGRESS_BAR_]"
    }),
    C.extend(C.FE.DEFAULTS, {
        fileUpload: !0,
        fileUploadURL: null,
        fileUploadParam: "file",
        fileUploadParams: {},
        fileUploadToS3: !1,
        fileUploadMethod: "POST",
        fileMaxSize: 10485760,
        fileAllowedTypes: ["*"],
        fileInsertButtons: ["fileBack", "|"],
        fileUseSelectedText: !1
    }),
    C.FE.PLUGINS.file = function(f) {
        var r, p = "https://i.froala.com/upload",
        l = 2,
        d = 3,
        u = 4,
        c = 5,
        v = 6,
        i = {};
        function g() {
            var e = f.popups.get("file.insert");
            e || (e = E()),
            e.find(".fr-layer.fr-active").removeClass("fr-active").addClass("fr-pactive"),
            e.find(".fr-file-progress-bar-layer").addClass("fr-active"),
            e.find(".fr-buttons").hide(),
            o(f.language.translate("Uploading"), 0)
        }
        function n(e) {
            var t = f.popups.get("file.insert");
            t && (t.find(".fr-layer.fr-pactive").addClass("fr-active").removeClass("fr-pactive"), t.find(".fr-file-progress-bar-layer").removeClass("fr-active"), t.find(".fr-buttons").show(), e && (f.events.focus(), f.popups.hide("file.insert")))
        }
        function o(e, t) {
            var i = f.popups.get("file.insert");
            if (i) {
                var r = i.find(".fr-file-progress-bar-layer");
                r.find("h3").text(e + (t ? " " + t + "%": "")),
                r.removeClass("fr-error"),
                t ? (r.find("div").removeClass("fr-indeterminate"), r.find("div > span").css("width", t + "%")) : r.find("div").addClass("fr-indeterminate")
            }
        }
        function h(e, t, i) {
            f.edit.on(),
            f.events.focus(!0),
            f.selection.restore(),
            f.opts.fileUseSelectedText && f.selection.text().length && (t = f.selection.text()),
            f.html.insert('<a href="' + e + '" target="_blank" id="fr-inserted-file" class="fr-file">' + t + "</a>");
            var r = f.$el.find("#fr-inserted-file");
            r.removeAttr("id"),
            f.popups.hide("file.insert"),
            f.undo.saveStep(),
            w(),
            f.events.trigger("file.inserted", [r, i])
        }
        function m(e) {
            var t = this.status,
            i = this.response,
            r = this.responseXML,
            n = this.responseText;
            try {
                if (f.opts.fileUploadToS3) if (201 == t) {
                    var o = function(e) {
                        try {
                            var t = C(e).find("Location").text(),
                            i = C(e).find("Key").text();
                            return ! 1 === f.events.trigger("file.uploadedToS3", [t, i, e], !0) ? (f.edit.on(), !1) : t
                        } catch(r) {
                            return U(u, e),
                            !1
                        }
                    } (r);
                    o && h(o, e, i || r)
                } else U(u, i || r);
                else if (200 <= t && t < 300) {
                    var a = function(e) {
                        try {
                            if (!1 === f.events.trigger("file.uploaded", [e], !0)) return f.edit.on(),
                            !1;
                            var t = JSON.parse(e);
                            return t.link ? t: (U(l, e), !1)
                        } catch(i) {
                            return U(u, e),
                            !1
                        }
                    } (n);
                    a && h(a.link, e, i || n)
                } else U(d, i || n)
            } catch(s) {
                U(u, i || n)
            }
        }
        function b() {
            U(u, this.response || this.responseText || this.responseXML)
        }
        function y(e) {
            if (e.lengthComputable) {
                var t = e.loaded / e.total * 100 | 0;
                o(f.language.translate("Uploading"), t)
            }
        }
        function U(e, t) {
            f.edit.on(),
            function(e) {
                g();
                var t = f.popups.get("file.insert").find(".fr-file-progress-bar-layer");
                t.addClass("fr-error");
                var i = t.find("h3");
                i.text(e),
                f.events.disableBlur(),
                i.focus()
            } (f.language.translate("Something went wrong. Please try again.")),
            f.events.trigger("file.error", [{
                code: e,
                message: i[e]
            },
            t])
        }
        function S() {
            f.edit.on(),
            n(!0)
        }
        function a(e) {
            if (void 0 !== e && 0 < e.length) {
                if (!1 === f.events.trigger("file.beforeUpload", [e])) return ! 1;
                var t, i = e[0];
                if (null === f.opts.fileUploadURL || f.opts.fileUploadURL == p) return s = i,
                (l = new FileReader).addEventListener("load",
                function() {
                    for (var e = l.result,
                    t = atob(l.result.split(",")[1]), i = [], r = 0; r < t.length; r++) i.push(t.charCodeAt(r));
                    e = window.URL.createObjectURL(new Blob([new Uint8Array(i)], {
                        type: s.type
                    })),
                    f.file.insert(e, s.name, null)
                },
                !1),
                g(),
                l.readAsDataURL(s),
                !1;
                if (i.size > f.opts.fileMaxSize) return U(c),
                !1;
                if (f.opts.fileAllowedTypes.indexOf("*") < 0 && f.opts.fileAllowedTypes.indexOf(i.type.replace(/file\//g, "")) < 0) return U(v),
                !1;
                if (f.drag_support.formdata && (t = f.drag_support.formdata ? new FormData: null), t) {
                    var r;
                    if (!1 !== f.opts.fileUploadToS3) for (r in t.append("key", f.opts.fileUploadToS3.keyStart + (new Date).getTime() + "-" + (i.name || "untitled")), t.append("success_action_status", "201"), t.append("X-Requested-With", "xhr"), t.append("Content-Type", i.type), f.opts.fileUploadToS3.params) f.opts.fileUploadToS3.params.hasOwnProperty(r) && t.append(r, f.opts.fileUploadToS3.params[r]);
                    for (r in f.opts.fileUploadParams) f.opts.fileUploadParams.hasOwnProperty(r) && t.append(r, f.opts.fileUploadParams[r]);
                    t.append(f.opts.fileUploadParam, i);
                    var n = f.opts.fileUploadURL;
                    f.opts.fileUploadToS3 && (n = f.opts.fileUploadToS3.uploadURL ? f.opts.fileUploadToS3.uploadURL: "https://" + f.opts.fileUploadToS3.region + ".amazonaws.com/" + f.opts.fileUploadToS3.bucket);
                    var o = f.core.getXHR(n, f.opts.fileUploadMethod);
                    o.onload = function() {
                        m.call(o, i.name)
                    },
                    o.onerror = b,
                    o.upload.onprogress = y,
                    o.onabort = S,
                    g();
                    var a = f.popups.get("file.insert");
                    a && a.off("abortUpload").on("abortUpload",
                    function() {
                        4 != o.readyState && o.abort()
                    }),
                    o.send(t)
                }
            }
            var s, l
        }
        function s() {
            n()
        }
        function E(e) {
            if (e) return f.popups.onHide("file.insert", s),
            !0;
            var t;
            f.opts.fileUpload || f.opts.fileInsertButtons.splice(f.opts.fileInsertButtons.indexOf("fileUpload"), 1),
            t = '<div class="fr-buttons">' + f.button.buildList(f.opts.fileInsertButtons) + "</div>";
            var i = "";
            f.opts.fileUpload && (i = '<div class="fr-file-upload-layer fr-layer fr-active" id="fr-file-upload-layer-' + f.id + '"><strong>' + f.language.translate("Drop file") + "</strong><br>(" + f.language.translate("or click") + ')<div class="fr-form"><input type="file" name="' + f.opts.fileUploadParam + '" accept="/*" tabIndex="-1" aria-labelledby="fr-file-upload-layer-' + f.id + '" role="button"></div></div>');
            var r, n = {
                buttons: t,
                upload_layer: i,
                progress_bar: '<div class="fr-file-progress-bar-layer fr-layer"><h3 tabIndex="-1" class="fr-message">Uploading</h3><div class="fr-loader"><span class="fr-progress"></span></div><div class="fr-action-buttons"><button type="button" class="fr-command fr-dismiss" data-cmd="fileDismissError" tabIndex="2" role="button">OK</button></div></div>'
            },
            o = f.popups.create("file.insert", n);
            return r = o,
            f.events.$on(r, "dragover dragenter", ".fr-file-upload-layer",
            function() {
                return C(this).addClass("fr-drop"),
                !1
            },
            !0),
            f.events.$on(r, "dragleave dragend", ".fr-file-upload-layer",
            function() {
                return C(this).removeClass("fr-drop"),
                !1
            },
            !0),
            f.events.$on(r, "drop", ".fr-file-upload-layer",
            function(e) {
                e.preventDefault(),
                e.stopPropagation(),
                C(this).removeClass("fr-drop");
                var t = e.originalEvent.dataTransfer;
                t && t.files && (r.data("instance") || f).file.upload(t.files)
            },
            !0),
            f.helpers.isIOS() && f.events.$on(r, "touchstart", '.fr-file-upload-layer input[type="file"]',
            function() {
                C(this).trigger("click")
            }),
            f.events.$on(r, "change", '.fr-file-upload-layer input[type="file"]',
            function() {
                if (this.files) {
                    var e = r.data("instance") || f;
                    e.events.disableBlur(),
                    r.find("input:focus").blur(),
                    e.events.enableBlur(),
                    e.file.upload(this.files)
                }
                C(this).val("")
            },
            !0),
            o
        }
        function e(e) {
            f.node.hasClass(e, "fr-file")
        }
        function t(e) {
            var t = e.originalEvent.dataTransfer;
            if (t && t.files && t.files.length) {
                var i = t.files[0];
                if (i && "undefined" != typeof i.type) {
                    if (i.type.indexOf("image") < 0) {
                        if (!f.opts.fileUpload) return e.preventDefault(),
                        e.stopPropagation(),
                        !1;
                        f.markers.remove(),
                        f.markers.insertAtPoint(e.originalEvent),
                        f.$el.find(".fr-marker").replaceWith(C.FE.MARKERS),
                        f.popups.hideAll();
                        var r = f.popups.get("file.insert");
                        return r || (r = E()),
                        f.popups.setContainer("file.insert", f.$sc),
                        f.popups.show("file.insert", e.originalEvent.pageX, e.originalEvent.pageY),
                        g(),
                        a(t.files),
                        e.preventDefault(),
                        e.stopPropagation(),
                        !1
                    }
                } else i.type.indexOf("image") < 0 && (e.preventDefault(), e.stopPropagation())
            }
        }
        function w() {
            var e, t = Array.prototype.slice.call(f.el.querySelectorAll("a.fr-file")),
            i = [];
            for (e = 0; e < t.length; e++) i.push(t[e].getAttribute("href"));
            if (r) for (e = 0; e < r.length; e++) i.indexOf(r[e].getAttribute("href")) < 0 && f.events.trigger("file.unlink", [r[e]]);
            r = t
        }
        return i[1] = "File cannot be loaded from the passed link.",
        i[l] = "No link in upload response.",
        i[d] = "Error during file upload.",
        i[u] = "Parsing response failed.",
        i[c] = "File is too large.",
        i[v] = "File file type is invalid.",
        i[7] = "Files can be uploaded only to same domain in IE 8 and IE 9.",
        {
            _init: function() {
                f.events.on("drop", t),
                f.events.$on(f.$win, "keydown",
                function(e) {
                    var t = e.which,
                    i = f.popups.get("file.insert");
                    i && t == C.FE.KEYCODE.ESC && i.trigger("abortUpload")
                }),
                f.events.on("destroy",
                function() {
                    var e = f.popups.get("file.insert");
                    e && e.trigger("abortUpload")
                }),
                f.events.on("link.beforeRemove", e),
                f.$wp && (w(), f.events.on("contentChanged", w)),
                E(!0)
            },
            showInsertPopup: function() {
                var e = f.$tb.find('.fr-command[data-cmd="insertFile"]'),
                t = f.popups.get("file.insert");
                if (t || (t = E()), n(), !t.hasClass("fr-active")) if (f.popups.refresh("file.insert"), f.popups.setContainer("file.insert", f.$tb), e.is(":visible")) {
                    var i = e.offset().left + e.outerWidth() / 2,
                    r = e.offset().top + (f.opts.toolbarBottom ? 10 : e.outerHeight() - 10);
                    f.popups.show("file.insert", i, r, e.outerHeight())
                } else f.position.forSelection(t),
                f.popups.show("file.insert")
            },
            upload: a,
            insert: h,
            back: function() {
                f.events.disableBlur(),
                f.selection.restore(),
                f.events.enableBlur(),
                f.popups.hide("file.insert"),
                f.toolbar.showInline()
            },
            hideProgressBar: n
        }
    },
    C.FE.DefineIcon("insertFile", {
        NAME: "file-o",
        FA5NAME: "file"
    }),
    C.FE.RegisterCommand("insertFile", {
        title: "Upload File",
        undo: !1,
        focus: !0,
        refreshAfterCallback: !1,
        popup: !0,
        callback: function() {
            this.popups.isVisible("file.insert") ? (this.$el.find(".fr-marker").length && (this.events.disableBlur(), this.selection.restore()), this.popups.hide("file.insert")) : this.file.showInsertPopup()
        },
        plugin: "file"
    }),
    C.FE.DefineIcon("fileBack", {
        NAME: "arrow-left"
    }),
    C.FE.RegisterCommand("fileBack", {
        title: "Back",
        undo: !1,
        focus: !1,
        back: !0,
        refreshAfterCallback: !1,
        callback: function() {
            this.file.back()
        },
        refresh: function(e) {
            this.opts.toolbarInline ? (e.removeClass("fr-hidden"), e.next(".fr-separator").removeClass("fr-hidden")) : (e.addClass("fr-hidden"), e.next(".fr-separator").addClass("fr-hidden"))
        }
    }),
    C.FE.RegisterCommand("fileDismissError", {
        title: "OK",
        callback: function() {
            this.file.hideProgressBar(!0)
        }
    })
});
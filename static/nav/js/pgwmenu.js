/**
 * PgwMenu - Version 2.0
 *
 * Copyright 2014, Jonathan M. Piat
 * http://pgwjs.com - http://pagawa.com
 * 
 * Released under the GNU GPLv3 license - http://opensource.org/licenses/gpl-3.0
 */
;(function($){
    $.fn.pgwMenu = function(options) {

        var defaults = {
            mainClassName : 'pgwMenu',
            dropDownLabel : '<span class="icon"></span>',
            viewMoreEnabled : true,
            viewMoreLabel : 'View more <span class="icon"></span>',
            viewMoreMaxWidth : 480
        };

        if (this.length == 0) {
            return this;
        } else if(this.length > 1) {
            this.each(function() {
                $(this).pgwMenu(options);
            });
            return this;
        }

        var pgwMenu = this;
        pgwMenu.plugin = this;
        pgwMenu.config = {};
        pgwMenu.resizeEvent = null;
        pgwMenu.window = $(window);

        // Init function
        var init = function() {

            // Merge user options with default configuration
            pgwMenu.config = $.extend({}, defaults, options);

            // Setup
            setup();
            pgwMenu.checkMenu();

            // Resize trigger
            pgwMenu.window.resize(function() {
                pgwMenu.plugin.css('overflow', 'hidden');

                clearTimeout(pgwMenu.resizeEvent);
                pgwMenu.resizeEvent = setTimeout(function() {
                    pgwMenu.checkMenu();
                }, 100);
            });

            pgwMenu.plugin.find('.pm-dropDown').click(function(e) {
                pgwMenu.enableMobileDropDown();
                e.stopPropagation();
            });

            pgwMenu.plugin.find('.pm-viewMore').click(function(e) {
                pgwMenu.enableViewMoreDropDown();
                e.stopPropagation();
            });

            $(document).click(function() {
                pgwMenu.disableMobileDropDown();
                pgwMenu.disableViewMoreDropDown();
            });
        };

        // Setup
        var setup = function() {
            var wrapClass = pgwMenu.config.mainClassName;
            var defaultClass = pgwMenu.plugin.attr('class');

            if (defaultClass && defaultClass.indexOf('light') > -1) {
                wrapClass += ' light';
            }

            pgwMenu.plugin.removeClass().addClass('pm-links');
            pgwMenu.plugin.wrap('<div class="' + wrapClass + '"></div>');
            pgwMenu.plugin = pgwMenu.plugin.parent();
            pgwMenu.plugin.prepend('<div class="pm-dropDown"><a href="javascript:void(0)">' + pgwMenu.config.dropDownLabel + '</a></div>');

            if (pgwMenu.config.viewMoreEnabled) {
                pgwMenu.plugin.append('<div class="pm-viewMore" style="display:inline-block"><a href="javascript:void(0)">' + pgwMenu.config.viewMoreLabel + '</a><ul></ul></div>');
            }
        };

        // Check menu
        pgwMenu.checkMenu = function() {
            var pluginMaxWidth = pgwMenu.plugin.width();

            if (pgwMenu.config.viewMoreEnabled) {
                var viewMoreLinkWidth = pgwMenu.plugin.find('.pm-viewMore').outerWidth(true);
            }

            function getContentWidth() {
                var menuContentWidth = 0;
                pgwMenu.plugin.find('.pm-links').removeClass('mobile').show();
                pgwMenu.plugin.find('.pm-links > li').each(function() {
                    menuContentWidth += $(this).outerWidth(true);
                });
                return menuContentWidth;
            }

            function switchMenu(type) {
                if (type == 'viewmore') {
                    var viewMoreMenuWidth = viewMoreLinkWidth;

                    pgwMenu.plugin.find('.pm-links').removeClass('mobile').show();
                    pgwMenu.plugin.find('.pm-viewMore > ul > li').remove();
                    pgwMenu.plugin.find('.pm-links > li').show().each(function() {
                        if (viewMoreMenuWidth + $(this).outerWidth(true) < pluginMaxWidth) {
                            viewMoreMenuWidth += $(this).outerWidth(true);
                        } else {
                            pgwMenu.plugin.find('.pm-viewMore > ul').append($(this).clone().show());
                            $(this).hide();
                        }
                    });

                    pgwMenu.plugin.find('.pm-dropDown, .pm-viewMore > ul').hide();
                    pgwMenu.plugin.find('.pm-viewMore').show();

                } else if (type == 'dropdown') {
                    pgwMenu.plugin.find('.pm-links > li').show();
                    pgwMenu.plugin.find('.pm-links').addClass('mobile').hide();
                    pgwMenu.plugin.find('.pm-viewMore, .pm-viewMore > ul').hide();
                    pgwMenu.plugin.find('.pm-viewMore > ul > li').remove();
                    pgwMenu.plugin.find('.pm-dropDown').show();

                } else {
                    pgwMenu.plugin.find('.pm-links > li').show();
                    pgwMenu.plugin.find('.pm-links').removeClass('mobile').show();
                    pgwMenu.plugin.find('.pm-dropDown, .pm-viewMore, .pm-viewMore > ul').hide();
                    pgwMenu.plugin.find('.pm-viewMore > ul > li').remove();
                }

                pgwMenu.plugin.find('.pm-dropDown > a, .pm-viewMore > a').removeClass('active');
            }

            if (getContentWidth() > pluginMaxWidth) {
                if (pgwMenu.config.viewMoreEnabled && (pluginMaxWidth > pgwMenu.config.viewMoreMaxWidth)) {
                    switchMenu('viewmore');
                } else {
                    switchMenu('dropdown');
                }
            } else {
                switchMenu('normal');
            }

            pgwMenu.plugin.css('overflow', '');
        };

        // Enable view more drop down
        pgwMenu.enableViewMoreDropDown = function() {
            if (pgwMenu.plugin.find('.pm-viewMore > a').hasClass('active')) {
                pgwMenu.disableViewMoreDropDown();
                return false;
            }

            pgwMenu.plugin.find('.pm-viewMore > a').addClass('active');
            pgwMenu.plugin.find('.pm-viewMore > ul').show();
        };

        // Disable view more drop down
        pgwMenu.disableViewMoreDropDown = function() {
            if (pgwMenu.plugin.find('.pm-viewMore > a').hasClass('active')) {
                pgwMenu.plugin.find('.pm-viewMore > a').removeClass('active');
                pgwMenu.plugin.find('.pm-viewMore > ul').hide();
            }
        };

        // Enable mobile drop down
        pgwMenu.enableMobileDropDown = function() {
            if (pgwMenu.plugin.find('.pm-dropDown > a').hasClass('active')) {
                pgwMenu.disableMobileDropDown();
                return false;
            }

            pgwMenu.plugin.find('.pm-dropDown > a').addClass('active');
            pgwMenu.plugin.find('.pm-links').show();
        };

        // Disable mobile drop down
        pgwMenu.disableMobileDropDown = function() {
            if (pgwMenu.plugin.find('.pm-dropDown > a').hasClass('active')) {
                pgwMenu.plugin.find('.pm-dropDown > a').removeClass('active');
                pgwMenu.plugin.find('.pm-links').hide();
            }
        };

        // Menu initialization
        init();

        return this;
    }
})(window.Zepto || window.jQuery);

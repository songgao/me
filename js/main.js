Date.prototype.toISO8601 = function() {
    var pad = function (amount, width) {
        var padding = "";
        while (padding.length < width - 1 && amount < Math.pow(10, width - padding.length - 1))
            padding += "0";
        return padding + amount.toString();
    }
    date = this;
    var offset = date.getTimezoneOffset();
    return pad(date.getFullYear(), 4)
    + "-" + pad(date.getMonth() + 1, 2)
    + "-" + pad(date.getDate(), 2)
    + "T" + pad(date.getHours(), 2)
    + ":" + pad(date.getMinutes(), 2)
    + ":" + pad(date.getUTCSeconds(), 2)
    + (offset > 0 ? "-" : "+")
    + pad(Math.floor(Math.abs(offset) / 60), 2)
    + ":" + pad(Math.abs(offset) % 60, 2);
};

var Main = function() { };

Main.prototype.init = function() {
    var self = this;
    var list = [
        {
            service: "github",
            user: "songgao"
        },
        {
            service: "twitter",
            user: "__songgao__"
        },
        {
            service: "stackoverflow",
            user: "218439"
        },
        {
            service: "atom",
            user: "http://blog.song.gao.io/atom.xml"
        }
    ];
    var count = 0;
    $("#lifestream").lifestream({
        limit: 1024,
        list:list,
        feedloaded: function() {
            count++;
            // Check if all the feeds have been loaded
            if( count === list.length ){
                self.feed_ready();
            }
        }
    });
};

Main.prototype.feed_ready = function() {

    var append_time = function(element){
        date = new Date(element.data("time"));
        element.append(' | <a class="timeago" href="' + element.data('url') + '" title="' + date.toISO8601() + '">' + date + "</a>");
    };

    $('li.lifestream-github').each(function() {
        var element = $(this);
        append_time(element);
        element.wrapInner('<div class="lifestream-li-content"/>');
        element.append('<i class="fa fa-github fa-3x"></i>');
    });

    $('li.lifestream-twitter').each(function() {
        var element = $(this);
        console.log(element);
        append_time(element);
        element.wrapInner('<div class="lifestream-li-content"/>');
        element.append('<i class="fa fa-twitter fa-3x twitter"></i>');
    });

    $('li.lifestream-atom').each(function() {
        var element = $(this);
        append_time(element);
        element.wrapInner('<div class="lifestream-li-content"/>');
        element.append('<i class="fa fa-rss fa-3x blog"></i>');
    });

    $('li.lifestream-stackoverflow').each(function() {
        var element = $(this);
        append_time(element);
        element.wrapInner('<div class="lifestream-li-content"/>');
        element.append('<i class="fa fa-stack-overflow fa-3x stack-overflow"></i>');
    });

    $("#lifestream .timeago").timeago();
}

$(window).load(function() {
    var main = new Main();
    main.init();
});

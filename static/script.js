jQuery.fn.flash = function () {
    var item = this;

    $({alpha:1}).animate({alpha:0}, {
        duration: 1000,
        step: function(){
            item.css('box-shadow','0 0 ' + this.alpha * 40 + 'px ' + '#38963c');
        },
        complete: function () {
            item.css('box-shadow', '0 0 0px #fff');
        }
    });
};

//noinspection JSUnresolvedFunction
new EventSource('/status').addEventListener('test', function(event) {
    var data = JSON.parse(event.data);

    var node = $("#" + data.name);

    //console.log($('.status', $("#" + data.name))[0].innerHTML);
    $('.status', node)[0].innerHTML = data.message;

    switch(data.status) {
        case "WORKING":
            node.css('backgroundColor', '#ffc372');
            break;
        case "WAITING":
            node.css('backgroundColor', '#ff353d');
            break;
        case "FINISHED":
            node.flash();
            break;
    }
});

$( "#customer-plus" ).click(function() {
    $.post( "/layout/customers", function( data ) {
        console.log(data);
    });
});
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
var source = new EventSource('/status');

source.addEventListener('nodeStatus', function(event) {
    var data = JSON.parse(event.data);
    var node = $("#" + data.name);

    var color;
    switch(data.status) {
        case "WORKING":
            color = '#ffc372';
            break;
        case "WAITING":
            color = '#ff353d';
            break;
        case "FINISHED":
            color = '#1eff1a';
            node.flash();
            break;
    }

    node.css('backgroundColor', color);
    $('.status', node).html(data.message);
});

source.addEventListener('rowUpdate', function(event) {
    var data = JSON.parse(event.data);

    switch(data.action) {
        case "REMOVE":
            setTimeout(function() {
                $("#" + data.name).remove();
            }, 1000);
            break;
        case "ADD":
            $("#row-" + data.row).append(

'<div class="box customer" id="' + data.name + '"> \
    <h2>' + data.name + '</h2> \
    <div class="status">Hello!</div> \
</div>'

            );
            break;
    }

});

$( "#customer-plus" ).click(function() {
    $.post( "/layout/customers", function( data ) {});
});
$(function() {
    var conn;


    function pageLoad() {
        if (window["WebSocket"]) {
            conn = new WebSocket("ws://localhost:9000/ws/alert");
            conn.onclose = function(evt) {
                console.log("WS Connection closed");
            };
            conn.onmessage = function(evt) {
                var messageAlert = evt.data;
                console.log(messageAlert)

                var alert = JSON.parse(messageAlert);
                console.log(alert)
                Messenger().post({
                message: JSON.stringify(alert),
                type: 'error',
                showCloseButton: true
                });


            };
        } else {
            console.log("Your Browser does not support WebSockets")
        }
    }

    pageLoad();
    SingApp.onPageLoad(pageLoad);
});
// $(function() {
var conn;
var file = "";

function download(filename, text) {
    var element = document.createElement('a');
    element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
    element.setAttribute('download', filename);

    element.style.display = 'none';
    document.body.appendChild(element);

    element.click();

    document.body.removeChild(element);
}

if (window["WebSocket"]) {
    conn = new WebSocket(ws_server + "/ws/alert");
    conn.onclose = function(evt) {
        console.log("WS Connection closed");
    };
    conn.onmessage = function(evt) {
        var messageAlert = evt.data;
        console.log(messageAlert)
        var alert = JSON.parse(messageAlert);
        console.log(alert)
    };
} else {
    console.log("Your Browser does not support WebSockets")
}

// pageLoad();
//     SingApp.onPageLoad(pageLoad);
// });

/*

HOST=127.0.0.1
HOSTNAME=beegons
PORT=9000
DB_USERNAME=
DB_PASSWORD=
DB_HOST=mongo-db
DB_PORT=27017
DB_DATABASE=test
ORION_CB_HOST=orion
ORION_CB_PORT=1026
CYGNUS_HOST=cygnus
CYGNUS_PORT=5050
CYGNUS_DATABASE=sth_default
FLINK_HOST=taskmanager
FLINK_PORT=9001*/
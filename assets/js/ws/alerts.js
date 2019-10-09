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
        conn = new WebSocket("ws://localhost:9000/ws/alert");
        conn.onclose = function(evt) {
            console.log("WS Connection closed");
        };
        conn.onmessage = function(evt) {
            var messageAlert = evt.data;
            console.log(messageAlert)
            var alert = JSON.parse(messageAlert);
            console.log(alert)


            let dateObserved = new Date(alert.dateObserved);
        

            let dateReceived = new Date();
            let timeDifference = dateReceived.getTime() - dateObserved.getTime();
            let counter = alert.counter;

            if(dateObserved == null || counter == null ) {
                return
            }

            let appendToFile = counter + "\t" + timeDifference + "\n";

            file += appendToFile;

            let max = 1000
            if(counte === max) {
                download(max+"_alerts.txt", file);
                file = "";
            }
            
            // Messenger().post({
            //     message: JSON.stringify(alert),
            //     type: 'error',
            //     showCloseButton: true
            // });


        };
    } else {
        console.log("Your Browser does not support WebSockets")
    }

    // pageLoad();
//     SingApp.onPageLoad(pageLoad);
// });
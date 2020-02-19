$(function() {

    var charts = [];
    var gauges = [];
    var dataId; 

    function pageLoad() {
        $('#moduleid').val("none");
        $("#parameter").val("none")

        $("#moduleid").on("change", function() {
            let selectedModule = $("#moduleid").val()
            dataId = selectedModule.replace("Module", "DataObserved");
            $.getJSON(url_server + `/api/module/${selectedModule}/realtime`, function(res) {
                if (res.status == 200) {
                    charts = [];
                    $("#charts").empty();

                    $("#gauges").empty();
                    $("#gaugesMet").empty();

                    $("#gaugesDiv").removeClass("hidden");
                    $("#gaugesDivMet").removeClass("hidden");

                    let sensors = res.content;
                    if (sensors == null) return;
                    sensors.sort(function(a,b) {
                        var x = a.parameter.toLowerCase();
                        var y = b.parameter.toLowerCase();
                        return x < y ? -1 : x > y ? 1 : 0;
                    });

                    sensors.forEach(sensor => {

                        gauges[sensor.parameter] = {};
                        charts[sensor.parameter] = {};

                        let nc = getNameAndColor(sensor.parameter);
                        appendChartDiv(sensor.parameter, nc.name);
                        let gauge = appendGauge(sensor.parameter, nc.name, nc.color);
                        if (sensor.data!= null) {
                            let chart = getChart(sensor.data, nc, sensor.parameter);
                            charts[sensor.parameter] = chart;
                            gauges[sensor.parameter] = gauge;
                            gauges[sensor.parameter].update(sensor.data[sensor.data.length - 1].attrValue)
                            chart.render();
                        }
                    });

                    console.log()
                }
            }).fail(function() {
                $("#charts").empty();

                $("#gauges").empty();
                $("#gaugesMet").empty();

                $("#gaugesDiv").addClass("hidden");
            });
        });

        if (window["WebSocket"]) {
            console.log("Connecting WS...")
            conn = new WebSocket(ws_server+"/ws/data");

            conn.onclose = function(evt) {
                console.log("WS Connection closed");
            };
            conn.onmessage = function(evt) {
                var messageData = evt.data;
                // console.log(messageData)
                var data = JSON.parse(messageData);
                // console.log(data)

                if(dataId === data.id) {
                    let timestamp = getTime(data['dateObserved'])
                    for (key in data) {
                        if (key !== "id") {
                            if(gauges[key] != undefined) {
                                gauges[key].update(data[key])
                                charts[key].series[0].data.shift()
                                charts[key].series[0].data.push({x: timestamp, y: data[key], y0: 0})
                                charts[key].update()
                            }
                        }
                    }
                }
            };
        } else {
            console.log("Your Browser does not support WebSockets")
        }
    }

    pageLoad();
    SingApp.onPageLoad(pageLoad);
});
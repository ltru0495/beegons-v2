$(function() {

    var charts = [];

    function pageLoad() {
        $('#moduleid').val("none");
        $("#parameter").val("none")
        $("#parameter").attr("disabled", "disabled")

        $("#moduleid").on("change", function() {
            let selectedModule = $("#moduleid").val()
            $.getJSON(url_server + `/api/module/${selectedModule}/realtime`, function(res) {
                if (res.status == 200) {
                    charts = [];
                    $("#charts").empty();
                    $("#parameter").removeAttr("disabled")
                    let sensors = res.content;
                    if (sensors == null) return;
                    sensors.forEach(sensor => {
                        console.log(sensor.parameter)
                        let nc = getNameAndColor(sensor.parameter);
                        appendChartDiv(sensor.parameter, nc.name);
                        let chart = getChart(sensor.data, nc.color, sensor.parameter);
                        charts.push(chart);
                        chart.render();
                    });
                }
            });
        });


    }

    pageLoad();
    SingApp.onPageLoad(pageLoad);
});
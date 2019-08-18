$(function() {

    var dataaa;
    var selectedModule;

    $('#moduleid').val("none");

    $('#moduleid').on('change', function() {
        selectedModule = this.value;
        moduleid = selectedModule;
        // const moduleString = "Module"
        // moduleid = selectedModule.substring(selectedModule
        // .indexOf(moduleString) + moduleString.length + 1, selectedModule.length)
        if (selectedModule === 'none') {
            return
        }
        console.log(selectedModule)

        $.getJSON(url_server + '/api/modulewdata/' + moduleid, function(res) {
            let moduleInfo = res.data.Module;
            let sensors = res.data.Sensors;
            let data = res.data.Data;
            appendDivs();
            initAirQualityDiv(sensors, data.parameters)
            initStationDiv(sensors, data.parameters)

            $.getJSON(url_server + '/api/lastdata/' + moduleid + "/airQualityIndex", function(r) {
                appendCard($("#charts"), "chartCard", 12, "Última Hora");
                appendChart("Calidad de Aire");
                initChart("airQualityIndex", r.data);
            });
        });


    });


    function appendDivs() {
        $("#content").append(`
            <div id="info" class="row"></div>
            <div id="charts" class="row">
            </div>`);


        appendCard($("#info"), "airQualityDiv", 6, "Calidad de Aire");
        appendCard($("#info"), "stationDiv", 5, "Estación Meterológica");
    }

    function initAirQualityDiv(sensors, dataValues) {
        $("#airQualityDiv").append(`
            <div class="col-md-4" id="airQualityIndexDiv" style="text-align:center;">
                <p id="airQualityIndex" style="display:inline-block;text-align:center">
                </p>
                <p id="airQualityLevel" style="display:inline-block;text-align:center">

                </p>
            </div>
            <div class="col-md-8 divDataContent" id="airQualityTable" style="text-align:center">
            </div>
        `)
            // $("#airQualityTable")
            // .append(`<table class="table"><tbody id="airQualityTableBody"></tbody></table>`)
        sensors.forEach(function(sensor) {
            if (sensor.model === "SPEC") {
                $("#airQualityTable")
                    .append('<div class="col-md-4" style="height:80px;">' +
                        '<p class="dataNumber">' + dataValues[sensor.parameter] + '</p>' +
                        '<p>' + sensor.name + '(' + sensor.unit + ')</p>' +
                        "</div>");
            }
        });

        let aqi = dataValues["airQualityIndex"]
        $("#airQualityIndex").append(`
          <p id="airQualityNumber">` + aqi + `</p>
          <p id="airQualityNumber">Moderado</p>
        `)
        $("#airQualityNumber").addClass("numberCircle");
        $("#airQualityNumber").css("color", colorAQI(aqi));
        $("#airQualityNumber").css("border", "4px solid" + colorAQI(aqi));
        $("#airQualityLevel").css("color", "white")
        $("#airQualityLevel")
            .append(`
                <p class="aqiLevel" style="background-color:green">Buena 0-50</p>
                <p class="aqiLevel" style="background-color:#dada25">Regular 51-100</p>
                <p class="aqiLevel" style="background-color:orange">Mala 101-150</p>
                <p class="aqiLevel" style="background-color:red">Muy mala 151-200</p>
                <p class="aqiLevel" style="background-color:purple">Peligrosa 201-250</p>)`);
    }


    function initStationDiv(sensors, dataValues) {
        $("#stationDiv").css("text-align", "center")
        $("#stationDiv").addClass("divDataContent")

        $("#stationDiv").append(`
        `)
        sensors.forEach(function(sensor) {
            if (sensor.model === "station") {
                $("#stationDiv")
                    .append('<div class="col-md-4" style="height:80px;">' +
                        '<p class="dataNumber">' + dataValues[sensor.parameter] + '</p>' +
                        '<p>' + sensor.name + '(' + sensor.unit + ')</p>' +
                        "</div>");
            }
        });
    }

    function initChart(parameter, rawData) {
        let data = []
        for (i = 0; i < rawData.length; i++) {
            data.push({
                x: (new Date(rawData[i]["dateObserved"])).getTime(),
                y: rawData[i][parameter]
            })
        }

        let graph = new Rickshaw.Graph({
            element: document.querySelector('#chart'),
            width: 1000,
            height: 350,
            renderer: 'line',
            interpolation: 'linear',
            stroke: true,
            series: [{
                data: data,
                color: "blue",
                name: "asdasd"
            }]
        });

        document.getElementById('chart').style = 'position:relative; left:40px;';
        document.getElementById('x_axis').style = 'position: relative; left: 40px; height: 40px;'
        document.getElementById('y_axis').style = 'position: absolute;  width: 40px; height: 500px;';

        ticksTreatment = "glow";
        var xAxis = new Rickshaw.Graph.Axis.X({
            graph: graph,
            orientation: 'bottom',
            element: document.getElementById("x_axis"),
            tickFormat: function(x) {
                var d = new Date(x);
                var h = d.getUTCHours();
                var m = d.getUTCMinutes();
                var s = d.getUTCSeconds();

                if (h < 10) h = "0" + h;
                if (m < 10) m = "0" + m;
                if (s < 10) s = "0" + s;

                return h + ":" + m + ":" + s;
            }
        });
        xAxis.render();
        var yAxis = new Rickshaw.Graph.Axis.Y({
            graph: graph,
            orientation: 'left',
            element: document.getElementById('y_axis'),
            ticksTreatment: ticksTreatment,
            tickFormat: Rickshaw.Fixtures.Number.formatKMBT,
        });
        yAxis.render();

        var hoverDetail = new Rickshaw.Graph.HoverDetail({
            graph: graph
        });

        graph.render();

    }



    function appendChart(title) {
        $('#chartCard').append(
            '<div class="col-md-9">' + ///
            '<div id="y_axis"></div>' +
            '<div id="chart" class="rickshaw_graph"></div>' +
            '<div id="x_axis"></div>' +
            '</div>'
        );
    }

    function colorAQI(aqi) {
        if (aqi > 0 & aqi <= 50) {
            return "green"
        }
        if (aqi > 50 & aqi <= 100) {
            return "#dada25"
        }
        if (aqi > 100 & aqi <= 150) {
            return "orange"
        }
        if (aqi > 150 & aqi <= 200) {
            return "red"
        }
        if (aqi > 200 & aqi <= 250) {
            return "purple"
        }
        return "red"
    }

    function appendCard(element, id, size, title) {
        element
            .append(`<div class="col-md-` + size + `"><div class="card"><div class="card-header">` +
                title + `</div><div class="card-body" id="` + id + `"></div></div></div>`)
    }

    function appendTable(parameters, values) {
        $('#dataTable').append(`<table class="table"><tbody id="table-body"></tbody></table>`)
        console.log(parameters)
        for (var i = 0; i < parameters.length; i++) {
            $("#table-body").append("<tr><td>" + parameters[i] + "</td><td>" + values[parameters[i]] + "</td></tr>")
        }
    }



});
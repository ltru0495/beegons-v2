function appendChartDiv(idgraph, title = "No Title") {
    $('#charts').append(
        '<div class="row">' +
        '<div class="col-md-12">' +
        '<section class="widget">' +
        '<header><h4>' + title + '</h4></header>' +
        '<div class="widget-body">' +
        '<div class="mt mb">' +
        '<div class="row">' + //
        '<div class="col-md-9">' + ///
        '<div id="y_axis' + idgraph + '"></div>' +
        '<div id="chart' + idgraph + '" class="rickshaw_graph"></div>' +
        '<div id="x_axis' + idgraph + '"></div>' +
        '</div>' +
        '</div>' +
        '</div>' +
        '</div>' +
        '</section>' +
        '</div>' +
        '</div>'
    );
}

function getTime(dateFormat) {
    let d = new Date(dateFormat);

    return d.getTime()
}



function getNameAndColor(parameter) {
    switch (parameter) {
        case "temperature":
            return {
                name: "Temperatura",
                color: "#387aa3"
            };
        case "humidity":
            return {
                name: "Humedad",
                color: "#77ab59"
            };
    }
    return {
        name: "No Title",
        color: "#7d5836"
    };
}

function sort(data) {
    for(k = 0 ; k < data.length; k ++ ) {
        for(i= k+1 ; i < data.length; i ++) {
            
        }
    }
}

function getChart(data, color, id) {
    var dataChart = [];

    data.forEach(singleDataObject => {
        let x = getTime(singleDataObject["recvTime"]);
        let y = Number(singleDataObject["attrValue"]);

        dataChart.unshift({
            x: x,
            y: y
        });
    });
    console.log(dataChart)
    // dataChart = sort(dataChart);
    let chart = new Rickshaw.Graph({
        element: document.querySelector("#chart" + id),
        width: 800,
        height: 350,
        renderer: 'area',
        interpolation: 'linear',
        stroke: true,

        series: [{
            data: dataChart,
            color: color
        }]
    });

    document.getElementById('chart' + id).style = 'position:relative; left:40px;';
    document.getElementById('x_axis' + id).style = 'position: relative; left: 40px; height: 40px;'
    document.getElementById('y_axis' + id).style = 'position: absolute;  width: 40px; ';

    ticksTreatment = "glow";
    var xAxis = new Rickshaw.Graph.Axis.X({
        graph: chart,
        orientation: 'bottom',
        element: document.getElementById("x_axis" + id),
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
        graph: chart,
        orientation: 'left',
        element: document.getElementById('y_axis' + id),
        ticksTreatment: ticksTreatment,
        tickFormat: Rickshaw.Fixtures.Number.formatKMBT,
    });
    yAxis.render();
    var hoverDetail = new Rickshaw.Graph.HoverDetail({
        graph: chart,
        formatter: function(series, x, y) {
            var date = '<span class="date">' + new Date(x).toUTCString() + '</span>';
            var swatch = '<span class="detail_swatch" style="background-color: ' + series.color + '"></span>';
            var content = swatch + series.name + ": " + Math.round(y * 1000) / 1000 + '<br>' + date;
            return content;
        }
    });

    var legend = document.querySelector('#legend' + id);


    return chart;

}
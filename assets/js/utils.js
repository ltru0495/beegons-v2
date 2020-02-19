function appendChartDiv(idgraph, title = "No Title") {
    $('#charts').append(
        '<div class="row">' +
        '<div class="col-md-12">' +
        '<section class="widget">' +
        '<header><h4>' + title + '</h4></header>' +
        '<div class="widget-body">' +
        '<div class="mt mb">' +
        '<div class="row">' + //
        '<div class="col-md-2"></div>' +
        '<div class="col-md-8">' + ///
        '<div id="y_axis' + idgraph + '"></div>' +
        '<div id="chart' + idgraph + '" class="rickshaw_graph"></div>' +
        '<div id="x_axis' + idgraph + '"></div>' +
        '</div>' +
        '<div id="spaceAfter" class="col-md-2"></div>' +
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


// Acalar u obscurecer un color (tomado de How to Calculate Lighter or Darker Hex Colors in JavaScript)
function ColorLuminance(hex, lum) {

        // validate hex string
        hex = String(hex).replace(/[^0-9a-f]/gi, '');
        if (hex.length < 6) {
                hex = hex[0]+hex[0]+hex[1]+hex[1]+hex[2]+hex[2];
        }
        lum = lum || 0;

        // convert to decimal and change luminosity
        var rgb = "#", c, i;
        for (i = 0; i < 3; i++) {
                c = parseInt(hex.substr(i*2,2), 16);
                c = Math.round(Math.min(Math.max(0, c + (c * lum)), 255)).toString(16);
                rgb += ("00"+c).substr(c.length);
        }

        return rgb;
}



function appendGauge(parameter, title, color) {
    if(parameter.indexOf("pm") != -1 || parameter.indexOf("gas") != -1){ //Calidad de aire
        $('#gauges').append('<div id="divfillgauge_' + parameter + '"><header class="gauge-label" style="text-align:center;">' + title + '</header></div>');
    } else{
        $('#gaugesMet').append('<div id="divfillgauge_' + parameter + '"><header class="gauge-label" style="text-align:center;">' + title + '</header></div>');
    }
    $('#divfillgauge_' + parameter).addClass("gauge");
    $('#divfillgauge_' + parameter).append('<svg id="fillgauge_' + parameter + '" width="100%" height="120"></svg>');

    var config = liquidFillGaugeDefaultSettings();
    config.circleColor = color;
    config.textColor = "#FF4444";
    config.waveTextColor = "#FFAAAA";
    config.waveColor = ColorLuminance(color,0.9);//"#FFDDDD";
    config.circleThickness = 0.2;
    config.textVertPosition = 0.52;

    config.waveHeight = 0.1;
    config.waveCount = 2;
    config.waveAnimateTime = 1000;
    config.displayPercent = false;
    config.textSize = 1.0;
    // .liquidFillGaugeText { font-family: Helvetica; font-weight: bold; font-size: 30px; }
    // $(".liquidFillGaugeText").css("font-size", 22);
    $(".liquidFillGaugeText").css("font-weight", 400);
    let gauge = loadLiquidFillGauge("fillgauge_" + parameter, 0, config);

    // No color over text
    //$('circle').css('fill', '#fff');
    $(".liquidFillGaugeText").css("fill", 'red');
    return gauge;
}



function getNameAndColor(parameter) {
    if (parameter.indexOf("temp") != -1){
        return {
                name: "Temperatura (°C)",
                color: "#387aa3"
        };
    }	    
    if (parameter.indexOf("humidity") != -1){
        return {
                name: "Humedad (%)",
                color: "#598042"
        };
    }	    
    if (parameter.indexOf("pressure") != -1){
        return {
                name: "Presión (hPa)",
                color: "#963b20"
        };
    }	    
    if (parameter.indexOf("o3") != -1){
        return gasColor("O"+"3".sub())
    }
    if (parameter.indexOf("co") != -1){
        return gasColor("CO")
    }
    if (parameter.indexOf("co2") != -1){
        return gasColor("CO"+"2".sub())
    }
    if (parameter.indexOf("so2") != -1){
        return gasColor("SO"+"2".sub())
    }
    if (parameter.indexOf("h2s") != -1){
        return gasColor("H"+"2".sub()+"S")
    }
    if (parameter.indexOf("no2") != -1){
        return gasColor("NO"+"2".sub())
    }
    if (parameter.indexOf("altitude") != -1){
        return {
                name: "Altitud (m)",
                color: "#808015"
        };
    }	    
    if (parameter.indexOf("UV") != -1){
        return {
                name: "Índice UV",
                color: "#660066", 
        };
    }
    if (parameter.indexOf("luminosity") != -1){
        return {
                name: "Luminosidad (lx)",
                color: "#f54242",
        };
    }
    if (parameter.indexOf("sonido") != -1){
        return {
                name: "Sonido (V)",
                color: "#9403fc",
        };
    }
    if (parameter.indexOf("pm") != -1) {
        if (parameter.indexOf("_") != -1){
            return {
                name: parameter.replace("_",".") + " (µg/m"+"3".sup()+")",
                color: "#fc5a03"
            };
	}		
	return {
            name: parameter + " (µg/m"+"3".sup()+")",
            color: "#fc5a03"
        };
    }
    return {
        name: parameter,
        color: "#598042"
    };
}

function gasColor(name) {
    return {
        name: name+" (ppm)",
        color: "#ddcb53"
    }
}



function getChart(data, nameAndColor, id) {
    var dataChart = [];

    data.forEach(singleDataObject => {
        let x = getTime(singleDataObject["recvTime"]);
        let y = Number(singleDataObject["attrValue"]);
        dataChart.unshift({
            x: x,
            y: y
        });
    });
    let chart = new Rickshaw.Graph({
        element: document.querySelector("#chart" + id),
        width: 1000,
        height: 350,
        renderer: 'area',
        interpolation: 'linear',
        stroke: true,

        series: [{
            data: dataChart,
            color: nameAndColor.color,
            name: nameAndColor.name
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

var opts = {
    lines: 8, // The number of lines to draw
    length: 38, // The length of each line
    width: 17, // The line thickness
    radius: 45, // The radius of the inner circle
    scale: 0.85, // Scales overall size of the spinner
    corners: 1, // Corner roundness (0..1)
    color: '#ffffff', // CSS color or array of colors
    fadeColor: 'transparent', // CSS color or array of colors
    speed: 1, // Rounds per second
    rotate: 0, // The rotation offset
    animation: 'spinner-line-fade-default', // The CSS animation name for the lines
    direction: 1, // 1: clockwise, -1: counterclockwise
    zIndex: 2e9, // The z-index (defaults to 2000000000)
    className: 'spinner', // The CSS class to assign to the spinner
    top: '50%', // Top position relative to parent
    left: '50%', // Left position relative to parent
    shadow: '0 0 1px transparent', // Box-shadow for the lines
    position: 'absolute' // Element positioning
};




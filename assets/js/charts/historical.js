$(function() {

    function formatDate(dateString) {
    	console.log(dateString)
        var aux = dateString.split(" ");
        var h = aux[1];
        var d = aux[0];
        var dmy = d.split("/");
        // YYYY-MM-DD
        var ydm = dmy[2] + "-" + dmy[1] + "-" + dmy[0];

        return ydm + "T" + h + "Z";
    }


    function initDatePicker() {
        $('#start_date').datetimepicker({});
        $('#start_date').on("dp.change", function(e) {
            $('#end_date').data("DateTimePicker").minDate(e.date);
        });

        $('#end_date').datetimepicker({
            useCurrent: false
        });
        $('#end_date').on("dp.change", function(e) {
            $('#start_date').data("DateTimePicker").maxDate(e.date);
        });
    }

    function pageLoad() {
        initDatePicker();
        $('#moduleid').val("none");
        $("#parameter").val("none")
        $("#parameter").attr("disabled", "disabled")

        $("#moduleid").on("change", function() {

        	let selectedModule = $("#moduleid").val()
            $.getJSON(url_server + `/api/module/${selectedModule}/parameters`, function(res) {
                if (res.status == 200) {
                    $("#parameter").removeAttr("disabled")
                    $("#parameter").empty()
                    $("#parameter").append(`<option value="none">Seleccionar parametro</option>`)
                    let params = res.content
                    if (params == null) return;
                    params.forEach(function(p) {
                        $("#parameter").append(`<option value="${p}">${p}</option>`)
                    })
                }
            });
        });

        $("#chart").on("click", function() {
            $("#charts").empty();
            let moduleid = $("#moduleid").val();
            let parameter = $("#parameter").val()
            let start = formatDate($("#start_date").val())
            let end = formatDate($("#end_date").val())
            if (moduleid != "none" && parameter != "none" && start != "" && end  != "") {

                $.getJSON(url_server + `/api/data/${moduleid}/${parameter}/${start}/${end}`, function(res) {
                    $("#charts").empty();
                    let content = res.content;
                    let p = content.parameter;
                    let nc = getNameAndColor(p);
                    appendChartDiv(p, nc.name);
                    let chart = getChart(content.data.reverse(), nc.color, p);
                    chart.render();
                	console.log(res)
                });

            }
        })
    }

    pageLoad();
    SingApp.onPageLoad(pageLoad);
});
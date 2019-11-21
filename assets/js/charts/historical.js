$(function() {

    function formatDate(dateString) {
        var aux = dateString.split(" ");
        let offset = 0 ;
        if (aux[2] == "PM") {
            offset = 12;
        }

        var h = aux[1];
        var d = aux[0];
        var dmy = d.split("/");
        // YYYY-MM-DD
        var ydm = dmy[2] + "-" + dmy[1] + "-" + dmy[0];

        return ydm + "T" + h + "Z";
    }


    function initDatePicker() {
        $('#start_date').datetimepicker({
            format: "MM/DD/Y HH:mm"
        });

        $('#start_date').on("dp.change", function(e) {
            $('#end_date').data("DateTimePicker").minDate(e.date);
        });

        $('#end_date').datetimepicker({
            format: "MM/DD/Y HH:mm",
            useCurrent: false
        });
        $('#end_date').on("dp.change", function(e) {
            $('#start_date').data("DateTimePicker").maxDate(e.date);
        });
    }

     function appendLinks(moduleid, parameter, start, end) {
        $("#spaceAfter").append(
            `<div class="row container">
                <div class="col-md-1"></div>
                <div><a target="_blank" class="btn btn-info btn-sm" href="${url_server}/api/file/${moduleid}/${parameter}/${start}/${end}/json">JSON</a>
            `
        )
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
                    // console.log(res)
                    $("#charts").empty();
                    let content = res.content;
                    let p = content.parameter;
                    let nc = getNameAndColor(p);
                    appendChartDiv(p, nc.name);
                    appendLinks(moduleid, parameter, start, end);
                    let chart = getChart(content.data.reverse(), nc, p);
                    chart.render();
                });

            }
        })
    }

    pageLoad();
    SingApp.onPageLoad(pageLoad);
});
$(function() {


    function pageLoad() {
        $('#moduleid').val("none");
        $("#parameter").val("none")
        // $("#parameter").attr("disabled", "disabled")

        $("#moduleid").on("change", function() {
        	let selectedModule = $("#moduleid").val()
            $.getJSON(url_server + `/api/module/${selectedModule}/realtime`, function(res) {
                if (res.status == 200) {
                    $("#parameter").removeAttr("disabled")
                    let params = res.content;
                    if (params == null) return;
                        
                    console.log(params)
                }
            });
        });

      
    }

    pageLoad();
    SingApp.onPageLoad(pageLoad);
});
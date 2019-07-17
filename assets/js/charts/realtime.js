$(function() {
    $('#moduleid').val("none");

    $('#moduleid').on('change', function() {
        selectedModule = this.value;
        const moduleString = "Module"
        moduleid = selectedModule.substring(selectedModule.indexOf(moduleString) + moduleString.length + 1, selectedModule.length)
        if (selectedModule === 'none') {
            return
        }
        $.getJSON(url_server + '/api/module/' + moduleid, function(mod) {
            console.log(mod.data)
        });

        $.getJSON(url_server + '/api/aqo/' + moduleid, function(aqo) {
            console.log(aqo.data)
        });
    });
});
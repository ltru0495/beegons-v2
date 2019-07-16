function post(url, data, method) {
    $.ajaxSettings.traditional = true;

    $.ajax({
        url: url,
        type: method,
        data: data,
        success: function() {
            console.log('OK');
        },
        error: function(jqXHR, textStatus, errorThrown) {
            console.log('ERROR');
            console.log(jqXHR.status);
            console.log('textStatus: '+textStatus);
            console.log('errorThrown: '+ errorThrown);

        }
    }).done(function(res){
    	console.log(res)
    });
}
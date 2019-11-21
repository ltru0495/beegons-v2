function post(url, data, method, callback) {
    $.ajaxSettings.traditional = true;

    $.ajax({
        url: url,
        type: method,
        data: data,
        success: function() {
            console.log('OK');
        },
        error: function(jqXHR, textStatus, errorThrown) {
            callback(textStatus)
            console.log('ERROR');
            console.log(jqXHR.status);
            console.log('textStatus: '+textStatus);
            console.log('errorThrown: '+ errorThrown);

        }
    }).done(function(res){
    	callback(res);
    });
}
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
            console.log('jqXHR:');
            console.log(jqXHR);
            console.log('textStatus:');
            console.log(textStatus);
            console.log('errorThrown:');
            console.log(errorThrown);

        }
    }).done(function(res){
    	console.log(res)
    });
}
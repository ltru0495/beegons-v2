$(function(){

    function pageLoad(){
    	console.log($('#infotable'));
		$('#infotable').DataTable();
    }
    pageLoad();
    SingApp.onPageLoad(pageLoad);

    
});
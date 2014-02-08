$(document).ready(function() {
	$('#station-names').change(function() {
		LoadStation();
    });
	
	$('#load-station-button').click(function() {
		LoadStation();
	});
	
	$('#station-names-json').change(function() {
		LoadStationJson();
    });
	
	$('#load-station-button-json').click(function() {
		LoadStationJsonOwnTab();
	});
	
	LoadStation();
	LoadStationJson();
});

function Standard_Callback() {
    try {
        alert(this.ResultString);
    }

    catch (e) {   
        alert(e);
    }
}

function Standard_ValidationCallback() {
    try {
        alert(this.ResultString);
    }

    catch (e) {   
        alert(e);
    }
}

function Standard_ErrorCallback() {
    try {
        alert(this.ResultString);
    }

    catch (e) {   
        alert(e);
    }
}

function LoadStation() {
	try {
		$('#stations-view').html('Loading View, Please Wait...');
		
		var postData = {};
		postData["stationId"] = $('#station-names').val();
		
        var service = new ServiceResult();
        service.getJSONData("/buoy/retrievestation",
                            postData,
                            LoadStation_Callback,
                            Standard_ValidationCallback,
                            Standard_ErrorCallback
                            );
    }

    catch (e) {
        alert(e);
    }
}

function LoadStation_Callback() {
	try {
		$('#stations-view').html(this.ResultObject);
	}
	
	catch (e) {
        alert(e);
    }
}

function LoadStationJson() {
	try {
		$('#stations-view').html('Loading View, Please Wait...');
		
		url = "/buoy/station/" + $('#station-names-json').val();
		
		var postData = {};

		var service = new ServiceResult();
        service.getJSONDataRaw(url,
                            postData,
                            LoadStationJson_Callback
                            );
    }

    catch (e) {
        alert(e);
    }
}

function LoadStationJson_Callback() {
	try {
		$('#stations-view-json').html(JSON.stringify(this.Data));
	}
	
	catch (e) {
        alert(e);
    }
}

function LoadStationJsonOwnTab() {
	url = "/buoy/station/" + $('#station-names-json').val();
	window.open(url);
}
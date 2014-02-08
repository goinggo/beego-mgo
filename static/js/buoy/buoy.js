$(document).ready(function() {
	$('.detail').click(function(e) {
		e.preventDefault();
		ShowDetail(this);
	});
	
	$('#station-names-json').change(function() {
		LoadStationJson();
    });
	
	$('#load-station-button-json').click(function() {
		LoadStationJsonOwnTab();
	});
	
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

function ShowDetail(result) {
	try {
		var postData = {};
		postData["stationId"] = $(result).attr('data');
		
        var service = new ServiceResult();
        service.getJSONData("/buoy/retrievestation",
                            postData,
                            ShowDetail_Callback,
                            Standard_ValidationCallback,
                            Standard_ErrorCallback
                            );
    }

    catch (e) {
        alert(e);
    }
}

function ShowDetail_Callback() {
	try {
		$('#system-modal-title').html("Buoy Details");
		$('#system-modal-content').html(this.ResultObject);
		$("#systemModal").modal('show');
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

$(document).ready(function() {
	$('#testajax').click(function() {
		TestAjax();
    });
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

function TestAjax() {
	try {
		var postData = {};
        var service = new ServiceResult();
        service.getJSONData("/testajax",
                            postData,
                            TestAjax_Callback,
                            Standard_ValidationCallback,
                            Standard_ErrorCallback
                            );
    }

    catch (e) {
        alert(e);
    }
}

function TestAjax_Callback() {
	alert('Name: ' + this.ResultObject.Name + ' - Test: ' + this.ResultObject.Test)
}
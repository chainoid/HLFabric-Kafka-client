// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	$("#error_sender").hide();
	$("#error_history").hide();
	$("#all_parsels").hide();
	$("#query_parsel").hide();
	$("#history_parsel").hide();
	$("#sender_parsels").hide();

	$("#error_parsel_id").hide();
	$("#error_delivered").hide();
	$("#success_delivery").hide();
	
			
	$scope.queryAllParsels = function(){

		appFactory.queryAllParsels(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				//parseInt(data[i].Key);
				data[i].Record.Key = data[i].Key;
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return a.senderTS.localeCompare(b.senderTS);
			});
			$scope.all_parsels = array;
		});

		$("#history_parsel").hide();
		$("#query_parsel").hide();
		$("#sender_parsels").hide();

		$("#all_parsels").show();
	}

	$scope.queryParsel = function(){

		var id = $scope.parsel_id;

		appFactory.queryParsel(id, function(data){
			$scope.query_parsel = data;

			if ($scope.query_parsel == "Could not locate parsel"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.showParselParams = function(parsel){
		
		$("#history_parsel").hide();
		$("#sender_parsels").hide();
				
		var id = parsel.Key;

		appFactory.queryParsel(id, function(data){
			$scope.query_parsel = data;

			if ($scope.query_parsel == "Could not locate parsel"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});

		$("#query_parsel").show();
	}

	$scope.querySender = function(){

		var name = $scope.sender_name;

		appFactory.querySender(name, function(data){
						
			var array = [];
			for (var i = 0; i < data.length; i++){
				//parseInt(data[i].Key);
				data[i].Record.Key = data[i].Key;
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return a.senderTS.localeCompare(b.senderTS);
			});

			$scope.sender_parsels = array;
			
			if (data  == "No parsels for sender"){
				console.log()
				$("#error_sender").show();
			} else{
				$("#error_sender").hide();
			}
		});

		$("#sender_parsels").show();
	}

    $scope.showSenderParsels = function(parsel){

		$("#query_parsel").hide();
		$("#history_parsel").hide();

		//alert(" ShowSenderParsels " + parsel.sender );

		var name = parsel.sender;

		appFactory.querySender(name, function(data){
						
			var array = [];
			for (var i = 0; i < data.length; i++){
				//parseInt(data[i].Key);
				data[i].Record.Key = data[i].Key;
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return a.senderTS.localeCompare(b.senderTS);
			});

			$scope.sender_parsels = array;
			
			if (data  == "No parsels for sender"){
				console.log()
				$("#error_sender").show();
			} else{
				$("#error_sender").hide();
			}
		});

		$("#sender_parsels").show();
	}
	
	$scope.historyParsel = function(){
		
		var historyId = $scope.historyId;

		appFactory.historyParsel(historyId, function(data){
						           
			var array = [];
			for (var i = 0; i < data.length; i++){
				data[i].Record.Key = parseInt(i);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return a.senderTS.localeCompare(b.senderTS);
			});
			$scope.history_parsel = array;
			
			if (data  == "No history for parsel"){
				console.log()
				$("#error_history").show();
			} else{alert(" HistoryParsel!!!!");
				$("#error_history").hide();
			}
		});
	}

	$scope.showHistoryParsel = function(parsel){

		$("#query_parsel").hide();
		$("#sender_parsels").hide();

		//alert(" ShowHistoryParsel for ID=" + parsel.Key);

		var historyId = parsel.Key;

		appFactory.historyParsel(historyId, function(data){
						           
			var array = [];
			for (var i = 0; i < data.length; i++){
				
				data[i].Record.TxId = data[i].TxId;
				data[i].Record.TxTS = data[i].TxTS;
				data[i].Record.IsDelete = data[i].IsDelete;
			
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return a.senderTS.localeCompare(b.senderTS);
			});
			$scope.history_parsel = array;
			
			if (data  == "No history for parsel"){
				console.log()
				$("#error_history").show();
			} else{
				$("#error_history").hide();
			}
		});

		$("#history_parsel").show();
		$("#history_parsel_id").show();
	}


	$scope.acceptParsel = function(){

		appFactory.acceptParsel($scope.parsel, function(data){
			$scope.accepted_parsel_id = data;
			$("#success_create").show();
		});
	}

	$scope.deliveryParsel = function(){

		appFactory.deliveryParsel($scope.parsel, function(data){
			$scope.delivery_parsel = data;
			
			$("#error_parsel_id").hide();
			$("#error_delivered").hide();
			$("#success_delivery").show();

			if ($scope.delivery_parsel == "Error: Parsel not found") {
				$("#error_parsel_id").show();
				$("#success_delivery").hide();
				$("#error_delivered").hide();
			} 
			else if ($scope.delivery_parsel == "Error: Already delivered") {
				$("#error_parsel_id").hide();
				$("#success_delivery").hide();
		    	$("#error_delivered").show();
        		} 
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllParsels = function(callback){
	
    	$http.get('/get_all_parsels/').success(function(output){
			callback(output)
		});
	}

	factory.queryParsel = function(id, callback){
    	$http.get('/get_parsel/'+id).success(function(output){
			callback(output)
		});
	}

	factory.acceptParsel = function(data, callback){
		
		var parsel = data.sender + "-" + data.senderBranch + "-" + data.receiver+ "-" + data.receiverBranch;

    	$http.get('/add_parsel/'+parsel).success(function(output){
			callback(output)
		});
	}

	factory.deliveryParsel = function(data, callback){

		var parsel = data.id;

    	$http.get('/delivery_parsel/'+parsel).success(function(output){
			callback(output)
		});
	}

	factory.historyParsel = function(historyId, callback){
    	$http.get('/history_parsel/'+historyId).success(function(output){
			callback(output)
		});
	}

    factory.querySender = function(name, callback){
    	$http.get('/get_sender/'+name).success(function(output){
			callback(output)
		});
	}
		
	

	return factory;
});



//SPDX-License-Identifier: Apache-2.0

var parsel = require('./controller.js');

module.exports = function(app){

  app.get('/get_parsel/:id', function(req, res){
    parsel.get_parsel(req, res);
  });
  app.get('/add_parsel/:parsel', function(req, res){
    parsel.add_parsel(req, res);
  });
  app.get('/get_all_parsels', function(req, res){
    parsel.get_all_parsels(req, res);
  });
  app.get('/delivery_parsel/:holder', function(req, res){
    parsel.delivery_parsel(req, res);
  });
  app.get('/get_sender/:name', function(req, res){
    parsel.get_sender(req, res);
  })
}

var request = require('request');
module.exports = { main: function (event, context) {
        console.log("data is" + JSON.stringify(event.data))
        request({
            url: "http://event-publish-service.kyma-system.svc.cluster.local:8080/v1/events",
            method: "POST",
            json: true,
            body: event.data
        }, function (error, response, body){
            console.log(response);
        });
    } }
const axios = require("axios");
module.exports = { main: async function (event, context) {
        console.log("data is" + JSON.stringify(event.data))
        const urlEvent = "http://event-publish-service.kyma-system.svc.cluster.local:8080/v1/events"
        var postEventResponse = await axios({
            method: 'post',
            url: urlEvent,
            headers: {
                'Content-Type': 'application/json'
            },
            data:event.data
        });
        console.log("response" + JSON.stringify(postEventResponse))
        return postEventResponse;
    } }
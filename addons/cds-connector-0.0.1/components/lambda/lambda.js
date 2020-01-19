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
        const Url = process.env.GATEWAY_URL + "/profile/htp741143850/profiles/44a14fe9-adb9-4fbd-ad2f-4c9d415fbcff";
        console.log( "sending get to:", Url);
        const { data: result } = await axios.get(Url, {});
        console.log(result);
        return result;
    } }
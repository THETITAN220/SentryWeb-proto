chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {

    if (request.type === "analyze") {

        fetch("http://localhost:8000/analyze", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                data: request.data
            })
        })
            .then(res => res.json())
            .then(data => sendResponse(data))
            .catch(err => sendResponse({ error: err.toString() }));

        return true;
    }
});

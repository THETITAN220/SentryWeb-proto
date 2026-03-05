function extractSecurityFeatures() {

    const visibleText = document.body.innerText.slice(0, 3000);

    const links = Array.from(document.querySelectorAll("a"))
        .map(a => a.href)
        .slice(0, 20);

    const forms = Array.from(document.querySelectorAll("form"))
        .map(f => f.action)
        .slice(0, 10);

    const scripts = Array.from(document.querySelectorAll("script"))
        .map(s => s.src || "inline-script")
        .slice(0, 10);

    const hiddenElements = Array.from(document.querySelectorAll("*"))
        .filter(el => {
            const style = window.getComputedStyle(el);
            return style.display === "none" || style.visibility === "hidden";
        })
        .map(el => el.innerText)
        .filter(text => text.length > 0)
        .slice(0, 10);

    return {
        visible_text: visibleText,
        links: links,
        forms: forms,
        scripts: scripts,
        hidden_elements: hiddenElements
    };
}

chrome.runtime.sendMessage(
    {
        type: "analyze",
        data: extractSecurityFeatures()
    },
    (response) => {

        if (!response) {
            console.error("No response from security agent");
            return;
        }

        if (response.error) {
            console.error("Security agent error:", response.error);
            return;
        }

        console.log("AI Security Analysis:", response.analysis);

    });

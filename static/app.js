const shortenButton = document.querySelector("#shorten");
const inputField = document.querySelector("#input-field");
const shortUrlElement = document.querySelector("#new-url");
const resultDiv = document.querySelector("#output-div");
const errorDiv = document.querySelector("#error-div");
const errorText = document.querySelector("#error-text");
const clearButton = document.querySelector("#clear-btn");
const copyButton = document.querySelector("#copy-btn");

/* ===== Helper functions ===== */
const showResult = () => (shortUrlElement.style.display = "flex");
const hideResult = () => (shortUrlElement.style.display = "none");
const showError = (msg) => {
    errorText.textContent = msg;
    errorDiv.style.display = "block";
};
const hideError = () => (errorDiv.style.display = "none");

/* ===== Clear fields ===== */
const clearFields = () => {
    inputField.value = "";
    hideResult();
    hideError();
};

clearButton.addEventListener("click", (e) => {
    e.preventDefault();
    clearFields();
});

/* ===== Shorten URL ===== */
async function shortenURL(longURL) {
    try {
        const res = await fetch("/shorten", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ url: longURL }),
        });

        if (!res.ok) {
            showError("Failed to shorten URL. Please try again.");
            hideResult();
            return;
        }

        const data = await res.json();
        shortUrlElement.textContent = data.short_url;
        shortUrlElement.style.cursor = "pointer";

        shortUrlElement.onclick = () => {
            window.open(data.short_url, "_blank");
        };

        showResult();
        hideError();
    } catch (err) {
        console.error(err);
        showError("Network error. Please check your connection.");
        hideResult();
    }
}

/* ===== Handle shorten button click ===== */
shortenButton.addEventListener("click", (e) => {
    e.preventDefault();
    const longURL = inputField.value.trim();

    if (!longURL) {
        showError("Please enter a URL.");
        hideResult();
        return;
    }

    shortenURL(longURL);
});

/* ===== Clipboard copy ===== */
const clipboard = new ClipboardJS("#copy-btn");
clipboard.on("success", (e) => {
    console.log("Copied:", e.text);
    e.clearSelection();
});
clipboard.on("error", (e) => {
    console.error("Copy failed:", e);
});

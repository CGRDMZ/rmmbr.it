let copyables = document.querySelectorAll('.copyable');

if (copyables.length > 0) {
    for (let i = 0; i < copyables.length; i++) {
        const copyable = copyables[i];

        let url = ""
        try {
            url = new URL(copyable.innerText).toString();
        } catch (error) {
            url = window.location.host + '/' + encodeURIComponent(decodeURIComponent(copyable.innerText))
        }
        copyable.addEventListener('click', (e) => copyToClipboard(url));
    }
}

let copyableUrl = document.querySelectorAll('.copyable-url')[0];
if (copyableUrl) {
    let url = ""
    try {
        url = new URL(copyableUrl.innerText).toString();
    } catch (error) {
        url = window.location.host + '/' + encodeURIComponent(decodeURIComponent(copyableUrl.innerText))
    }
    copyableUrl.addEventListener('click', (e) => copyToClipboard(url));
}

async function copyToClipboard(text) {


    if (!text) return;

    // set clipboard text
    await navigator.clipboard.writeText(text);

    alert('Copied to clipboard!');

}
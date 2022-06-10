let copyables = document.querySelectorAll('.copyable');

if (copyables.length > 0) {
    for (let i = 0; i < copyables.length; i++) {
        const copyable = copyables[i];
        copyable.addEventListener('click', (e) => copyToClipboard(e.target.innerText));
    }
}

let copyableUrl = document.querySelectorAll('.copyable-url')[0];
if (copyableUrl) {
    copyableUrl.addEventListener('click', (e) => copyToClipboard(window.location.host + '/' + e.target.innerText));
}

async function copyToClipboard(text) {


    if (!text) return;

    // set clipboard text
    await navigator.clipboard.writeText(text);

    alert('Copied to clipboard!');

}
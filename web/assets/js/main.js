let copyables = document.querySelectorAll('.copyable');

if (copyables.length > 0) {
    for (let i = 0; i < copyables.length; i++) {
        const copyable = copyables[i];
        copyable.addEventListener('click', copyToClipboard);
    }
}


async function copyToClipboard(e) {

    const url = e.target.innerText;

    if (!url) return;

    // set clipboard text
    await navigator.clipboard.writeText(url);

    alert('Copied to clipboard!');

}
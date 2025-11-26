const shortenForm = document.getElementById('shorten-form');
const urlInput = document.getElementById('url-input');
const resultBox = document.getElementById('result-box');
const shortLinkDisplay = document.getElementById('shortLinkDisplay');
const copyButton = document.querySelector('.copy-btn');

shortenForm.addEventListener('submit', async (event) => {
    event.preventDefault();

    const longUrl = urlInput.value;
    if (!longUrl) {
        alert('Please enter a URL.');
        return;
    }

    try {
        const response = await fetch('/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ long_url: longUrl }),
        });

        if (!response.ok) {
            // Gracefully handle non-JSON error responses
            let errorMessage = `HTTP error ${response.status}: ${response.statusText}`;
            try {
                const errorData = await response.json();
                errorMessage = errorData.error || errorMessage;
            } catch (e) {
                // The error response was not JSON. We can ignore this error
                // and stick with the default HTTP error message.
                console.warn("Could not parse error response as JSON.");
            }
            throw new Error(errorMessage);
        }

        // Check if the response body is empty before trying to parse it
        const contentLength = response.headers.get('Content-Length');
        if (contentLength === '0') {
            throw new Error('Received an empty response from the server.');
        }
        const data = await response.json();

        // Cập nhật link rút gọn và hiển thị kết quả
        shortLinkDisplay.href = data.short_url;
        shortLinkDisplay.textContent = data.short_url;
        resultBox.style.display = 'flex'; // Hiển thị hộp kết quả

    } catch (error) {
        alert('Error: ' + error.message);
    }
});

// Hàm sao chép link
function copyLink() {
    const linkToCopy = shortLinkDisplay.href;
    navigator.clipboard.writeText(linkToCopy).then(() => {
        // Thay đổi text của nút để thông báo đã copy thành công
        copyButton.textContent = 'Copied!';
        setTimeout(() => { copyButton.textContent = 'Copy'; }, 2000); // Reset text sau 2 giây
    }).catch(err => {
        alert('Failed to copy link: ' + err);
    });
}
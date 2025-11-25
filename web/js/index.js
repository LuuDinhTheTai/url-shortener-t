// Lấy các element từ DOM
const shortenForm = document.getElementById('shorten-form');
const urlInput = document.getElementById('url-input');
const resultBox = document.getElementById('result-box');
const shortLinkDisplay = document.getElementById('shortLinkDisplay');
const copyButton = document.querySelector('.copy-btn');

// Xử lý sự kiện submit form
shortenForm.addEventListener('submit', async (event) => {
    event.preventDefault(); // Ngăn form gửi theo cách truyền thống

    const longUrl = urlInput.value;
    if (!longUrl) {
        alert('Please enter a URL.');
        return;
    }

    try {
        // Gửi yêu cầu POST đến server bằng Fetch API
        const response = await fetch('/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ long_url: longUrl }),
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Something went wrong');
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
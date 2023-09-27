document.getElementById('uploadButton').addEventListener('click', function () {
    const fileInput = document.createElement('input');
    fileInput.type = 'file';
    fileInput.accept = '*/*'; // Определите принимаемые типы файлов
    fileInput.click();

    fileInput.addEventListener('change', function () {
        const selectedFile = fileInput.files[0];
        if (selectedFile) {
            const formData = new FormData();
            formData.append('file', selectedFile);

            fetch('/upload', {
                method: 'POST',
                body: formData,
            })
            .then(response => response.json())
            .then(data => {
                // Обработайте ответ от сервера, если необходимо
                console.log('Server response:', data);
            })
            .catch(error => {
                // Обработайте ошибку, если необходимо
                console.error('Error:', error);
            });
        }
    });
});
function auth() {
    // создать объект для формы
    let formData = new FormData(document.forms.authForm);

    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/auth/login");
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            if (this.status === 200) {
                window.location.href = JSON.parse(xhr.responseText);
            } else {
                document.getElementById('passwordInput').classList.add('is-invalid');
            }
        }
    }
    xhr.send(formData);
}
function createUser() {
    // создать объект для формы
    let formData = new FormData(document.forms.userCreateForm);
    let jsonObj = {}
    formData.forEach((value, key) => {
        jsonObj[key] = value;
    });
    jsonObj['departments'] = [document.getElementById('departments').value]
    jsonObj['specializations'] = [document.getElementById('specializations').value]
    jsonObj['positions'] = [document.getElementById('positions').value]
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/user/");
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            if (this.status === 200) {
                // Находим Toast и инициализируем его
                var successToast = new bootstrap.Toast(document.getElementById('successToast'));

                // Вызываем метод show() для отображения уведомления
                successToast.show();

                // Через 5 секунд скрываем уведомление (можете настроить время по своему усмотрению)
                setTimeout(function() {
                    successToast.hide();
                }, 5000);            } else {
                console.log("bad");
                // document.getElementById('passwordInput').classList.add('is-invalid');
            }
        }
    }
    xhr.send(JSON.stringify(jsonObj));
}
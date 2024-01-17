function createUser() {
    // создать объект для формы
    let formData = new FormData(document.forms.userCreateForm);
    formData.append('departments', [document.getElementById('departments').value]);
    formData.append('specializations', [document.getElementById('specializations').value]);
    formData.append('positions',[document.getElementById('positions').value]);
    let jsonObj = {}
    formData.forEach((value, key) => {
        jsonObj[key] = value;
    });
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/user/");
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            if (this.status === 200) {
                console.log("good");
                // window.location.href = JSON.parse(xhr.responseText);
            } else {
                console.log("bad");
                // document.getElementById('passwordInput').classList.add('is-invalid');
            }
        }
    }
    xhr.send(JSON.stringify(jsonObj));
}
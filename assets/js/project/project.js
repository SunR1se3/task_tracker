function createProject() {
    // создать объект для формы
    let formData = new FormData(document.forms.projectCreateForm);
    let jsonObj = {}
    formData.forEach((value, key) => {
        jsonObj[key] = value;
    });
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/project/");
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            let resp = JSON.parse(xhr.responseText)
            if (resp.status) {
                console.log("good");
            }
        }
    }
    xhr.send(JSON.stringify(jsonObj));
}
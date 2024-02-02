function createUser() {
    // создать объект для формы
    let formData = new FormData(document.forms.userCreateForm);
    let jsonObj = {}
    formData.forEach((value, key) => {
        jsonObj[key] = value;
    });
    let selectData = prepareData('departments', 'specializations', 'positions');
    jsonObj['departments'] = selectData[0]
    jsonObj['specializations'] = selectData[1]
    jsonObj['positions'] = selectData[2]
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/user/");
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            let resp = JSON.parse(xhr.responseText)
            if (resp.status) {
                let alertUserAdd = document.getElementById('alertUserAdd')
                alertUserAdd.hidden = false
                setTimeout(function() {
                    alertUserAdd.hidden = true
                }, 3000)
                updateTable();
            }
        }
    }
    xhr.send(JSON.stringify(jsonObj));
}

function updateUser(userId) {
    // создать объект для формы
    let formData = new FormData(document.forms.userEditForm);
    let jsonObj = {}
    formData.forEach((value, key) => {
        jsonObj[key] = value;
    });
    let selectData = prepareData('departmentsEdit', 'specializationsEdit', 'positionsEdit');
    jsonObj['departments'] = selectData[0];
    jsonObj['specializations'] = selectData[1];
    jsonObj['positions'] = selectData[2];
    jsonObj['systemRole'] = parseInt(document.getElementById('systemRoleEdit').value, 10);

    let xhr = new XMLHttpRequest();
    xhr.open("PUT", `/api/user/${userId}`);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            let resp = JSON.parse(xhr.responseText)
            if (resp.status) {
                let alertUserAdd = document.getElementById('alertUserSave')
                alertUserAdd.hidden = false
                setTimeout(function() {
                    alertUserAdd.hidden = true
                }, 3000)
                updateTable();
            }
        }
    }
    xhr.send(JSON.stringify(jsonObj));
}

function disableUser(e) {
    console.log(e);
    let userId = e.getAttribute('data-user-id');
    let disable =  e.checked;
    let xhr = new XMLHttpRequest();
    xhr.open("PUT", `/api/user/${userId}/activation?disabled=${disable}`);
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            let resp = JSON.parse(xhr.responseText)
            if (resp.status) {
                updateTable();
            }
        }
    }
    xhr.send();
}

function getUserData(e) {
    // создать объект для формы
    let xhr = new XMLHttpRequest();
    let userId = e.getAttribute('data-user-id');
    xhr.open("GET", `/admin/users/edit/${userId}`);
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            let resp = JSON.parse(xhr.responseText)
            if (resp.status) {
                // Создаем новый элемент div с помощью парсинга HTML из AJAX-ответа
                var newElement = document.createElement('div');
                newElement.innerHTML = JSON.parse(this.responseText).data;
                let domEditModal = document.getElementById('userEditModal')
                // Заменяем текущий элемент на новый
                domEditModal.outerHTML = newElement.innerHTML;

                // Пересоздаем объект bootstrap.Modal
                let editModal = new bootstrap.Modal(document.getElementById('userEditModal'));
                editModal.show();
            }
        }
    }
    xhr.send();
}

function updateTable() {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/admin/users/update_table");
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            let resp = JSON.parse(xhr.responseText)
            if (resp.status) {
                document.getElementById("tableUser").innerHTML = JSON.parse(this.responseText).data
            }
        }
    }
    xhr.send();
}

function prepareData(depFieldId, specFieldId, posFieldId) {
    let deps = document.getElementById(depFieldId).value
    deps = deps === "" ? null : [deps]
    let spec = document.getElementById(specFieldId).value
    spec = spec === "" ? null : [spec]
    let pos = document.getElementById(posFieldId).value
    pos = pos === "" ? null : [pos]
    return [deps, spec, pos];
}
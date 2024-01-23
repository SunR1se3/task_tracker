function createUser() {
    // создать объект для формы
    let formData = new FormData(document.forms.userCreateForm);
    let jsonObj = {}
    formData.forEach((value, key) => {
        jsonObj[key] = value;
    });
    jsonObj['departments'] = prepareData()[0]
    jsonObj['specializations'] = prepareData()[1]
    jsonObj['positions'] = prepareData()[2]
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/user/");
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            if (this.status === 200) {
                let alertUserAdd = document.getElementById('alertUserAdd')
                alertUserAdd.hidden = false
                setTimeout(function() {
                    alertUserAdd.hidden = true
                }, 3000)
                getUserById(JSON.parse(this.responseText).id);
            }
        }
    }
    xhr.send(JSON.stringify(jsonObj));
}

function getUserById(id) {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", `/api/user/${id}`);
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            if (this.status === 200) {
                insertUserIntoTable(JSON.parse(this.responseText).data)
            }
        }
    }
    xhr.send();
}

function insertUserIntoTable(user) {
    let table = document.getElementById("tableUser");
    let newRow = table.insertRow();
    let tableHeader = table.rows[0].cells

    for (let i = 0; i < tableHeader.length; i++) {
        let cell = newRow.insertCell();
        switch (tableHeader[i].dataset.ident) {
            case "fio":
                cell.innerHTML = user.lastname + " " + user.firstname + " " + user.middlename;
                break;
            case "actions":
                cell.innerHTML = `<a href="${user.id}">Редактировать</a> <a href="#">Удалить</a>`;
                break;
            case "positions":
                let pos = user.positions
                for (let j = 0; j < pos.length; j++) {
                    cell.innerHTML = pos[j].title
                }
                break;
            case "specializations":
                let spec = user.specializations
                for (let j = 0; j < spec.length; j++) {
                    cell.innerHTML = spec[j].title
                }
                break;
            case "departments":
                let dep = user.departments
                for (let j = 0; j < dep.length; j++) {
                    cell.innerHTML = dep[j].title
                }
                break;
            default:
                cell.innerHTML = user[tableHeader[i].dataset.ident]
                break;
        }
    }
    // for (let key in user) {
    //     if (user.hasOwnProperty(key)) {
    //         let cell = newRow.insertCell();
    //         cell.innerHTML = user[key];
    //     }
    // }
}

function prepareData() {
    let deps = document.getElementById('departments').value
    deps = deps === "" ? null : [deps]
    let spec = document.getElementById('specializations').value
    spec = spec === "" ? null : [spec]
    let pos = document.getElementById('positions').value
    pos = pos === "" ? null : [pos]
    return [deps, spec, pos];
}
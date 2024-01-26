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
                updateTable();
                // getUserById(JSON.parse(this.responseText).id);
            }
        }
    }
    xhr.send(JSON.stringify(jsonObj));
}

function updateTable() {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/admin/users/update_table");
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            if (this.status === 200) {
                document.getElementById("tableUser").innerHTML = JSON.parse(this.responseText).data
            }
        }
    }
    xhr.send();
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
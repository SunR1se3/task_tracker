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
            }
        }
    }
    xhr.send(JSON.stringify(jsonObj));
}

function setProjectRole(userId, select) {
    let jsonObj = {}
    jsonObj['userId'] = userId;
    jsonObj['projectId'] = document.getElementById('projectContent').dataset.projectId;
    let projectRole = select.value;
    if (projectRole !== "") {
        jsonObj['projectRoleId'] = projectRole;
    }

    console.log(jsonObj);
    console.log(document.getElementById('projectContent'));
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/project/set_role");
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            let resp = JSON.parse(xhr.responseText)
            if (resp.status) {
            }
        }
    }
    xhr.send(JSON.stringify(jsonObj));
}

function addUserToTeam() {
    let userId = document.getElementById('autocompleteInput').dataset.selectedId
    let projectId = document.getElementById('projectContent').dataset.projectId
    let jsonObj = {
        "userId": userId,
        "projectId": projectId
    }

    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/project/add_to_team");
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            let resp = JSON.parse(xhr.responseText)
            if (resp.status) {
            }
        }
    }
    xhr.send(JSON.stringify(jsonObj));
}
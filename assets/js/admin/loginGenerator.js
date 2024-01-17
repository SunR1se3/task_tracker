
function generateLogin() {
    const lastNameInput = document.getElementById('lastname');
    const firstNameInput = document.getElementById('firstname');
    const middleNameInput = document.getElementById('middlename');
    const loginInput = document.getElementById('login');

    lastNameInput.addEventListener('input', function() {
        updateLogin(lastNameInput, firstNameInput, middleNameInput, loginInput);
    });
    firstNameInput.addEventListener('input',  function() {
        updateLogin(lastNameInput, firstNameInput, middleNameInput, loginInput);
    });
    middleNameInput.addEventListener('input', function() {
        updateLogin(lastNameInput, firstNameInput, middleNameInput, loginInput);
    });
}

function updateLogin(lastNameInput, firstNameInput, middleNameInput, loginInput) {
    let lastname = translateLiterate(lastNameInput.value);
    lastname = lastname !== "" ? lastname + " " : "";
    let firstname = translateLiterate(firstNameInput.value.charAt(0));
    firstname = firstname !== "" ? firstname + ". " : "";
    let middlename = translateLiterate(middleNameInput.value.charAt(0));
    middlename = middlename !== "" ? middlename + "." : "";
    loginInput.value = lastname + firstname + middlename;
}

function translateLiterate(s) {
    let cyr = [
        'ж', 'ч', 'щ', 'ш', 'ю', 'а', 'б', 'в', 'г', 'д', 'е', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п',
        'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ъ', 'ь', 'я', 'ы', 'э', 'ё',
        'Ж', 'Ч', 'Щ', 'Ш', 'Ю', 'А', 'Б', 'В', 'Г', 'Д', 'Е', 'З', 'И', 'Й', 'К', 'Л', 'М', 'Н', 'О', 'П', 'Р',
        'С', 'Т', 'У', 'Ф', 'Х', 'Ц', 'Ъ', 'Ь', 'Я', 'Ы', 'Э', 'Ё'
    ];

    let lat = [
        'zh', 'ch', 'sht', 'sh', 'yu', 'a', 'b', 'v', 'g', 'd', 'e', 'z', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
        'r', 's', 't', 'u', 'f', 'h', 'c', '', '', 'ya', 'yi', 'e', 'yo',
        'Zh', 'Ch', 'Sht', 'Sh', 'Yu', 'A', 'B', 'V', 'G', 'D', 'E', 'Z', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
        'R', 'S', 'T', 'U', 'F', 'H', 'c', '', '', 'Ya', 'Yi', 'E', 'Yo'
    ];

    return s.replace(/./g, function (char) {
        let index = cyr.indexOf(char);
        return index !== -1 ? lat[index] : char;
    }).toLowerCase();
}
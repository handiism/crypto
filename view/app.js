function affine() {
  document.getElementById("choice").value = "affine";
  document.getElementById("choicet").innerHTML =
    "Current chosen is Affine Chiper algorithm";

  document.getElementById("keyLabelNumberOne").hidden = false;
  document.getElementById("keyLabelNumberOne").innerHTML = "Key A";
  document.getElementById("keyNumberOne").hidden = false;

  document.getElementById("keyLabelNumberTwo").hidden = false;
  document.getElementById("keyLabelNumberTwo").innerHTML = "Key B";
  document.getElementById("keyNumberTwo").hidden = false;

  document.getElementById("keyLabelText").hidden = true;
  document.getElementById("keyLabelText").innerHTML = "Key";
  document.getElementById("keyText").hidden = true;
}
function caesar() {
  document.getElementById("choice").value = "caesar";
  document.getElementById("choicet").innerHTML =
    "Current chosen is Caesar Chiper algorithm";

  document.getElementById("keyLabelNumberOne").hidden = false;
  document.getElementById("keyLabelNumberOne").innerHTML = "Key";
  document.getElementById("keyNumberOne").hidden = false;

  document.getElementById("keyLabelNumberTwo").hidden = true;
  document.getElementById("keyLabelNumberTwo").innerHTML = "";
  document.getElementById("keyNumberTwo").hidden = true;

  document.getElementById("keyLabelText").hidden = true;
  document.getElementById("keyLabelText").innerHTML = "";
  document.getElementById("keyText").hidden = true;
}
function railfence() {
  document.getElementById("choice").value = "railfence";
  document.getElementById("choicet").innerHTML =
    "Current chosen is Railfence Chiper algorithm";

  document.getElementById("keyLabelNumberOne").hidden = false;
  document.getElementById("keyLabelNumberOne").innerHTML = "Key";
  document.getElementById("keyNumberOne").hidden = false;

  document.getElementById("keyLabelNumberTwo").hidden = true;
  document.getElementById("keyNumberTwo").hidden = true;
  document.getElementById("keyLabelNumberTwo").innerHTML = "";

  document.getElementById("keyLabelText").hidden = true;
  document.getElementById("keyText").hidden = true;
  document.getElementById("keyLabelText").innerHTML = "";
}
function superEncryption() {
  document.getElementById("choice").value = "super";
  document.getElementById("choicet").innerHTML =
    "Current chosen is Super Encryption algorithm";

  document.getElementById("keyLabelNumberOne").hidden = false;
  document.getElementById("keyLabelNumberOne").innerHTML = "Key A";
  document.getElementById("keyNumberOne").hidden = false;

  document.getElementById("keyLabelNumberTwo").hidden = false;
  document.getElementById("keyNumberTwo").hidden = false;
  document.getElementById("keyLabelNumberTwo").innerHTML = "Key B";

  document.getElementById("keyLabelText").hidden = true;
  document.getElementById("keyText").hidden = true;
  document.getElementById("keyLabelText").innerHTML = "";
}

async function encipher() {
  let alg = document.getElementById("choice").value;
  let key = {
    numbers: [
      parseInt(document.getElementById("keyNumberOne").value),
      parseInt(document.getElementById("keyNumberTwo").value),
    ],
    text: document.getElementById("keyText").value,
  };

  var plaintext = document.getElementById("plaintext").value;

  await fetch("/encipher", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      algorithm: alg,
      plaintext: plaintext,
      key: key,
    }),
  })
    .then((res) => res.json())
    .then((json) => {
      if (json["status"] == "success") {
        document.getElementById(
          "result"
        ).innerHTML = `Chipertext: ${json["data"]["ciphertext"]}`;
      } else {
        alert(json["data"]["message"]);
      }
    })
    .catch((err) => {
      alert(err);
    });
}

async function decipher() {
  let alg = document.getElementById("choice").value;
  let key = {
    numbers: [
      parseInt(document.getElementById("keyNumberOne").value),
      parseInt(document.getElementById("keyNumberTwo").value),
    ],
    text: document.getElementById("keyText").value,
  };

  var ciphertext = document.getElementById("plaintext").value;

  await fetch("/decipher", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      algorithm: alg,
      ciphertext: ciphertext,
      key: key,
    }),
  })
    .then((res) => res.json())
    .then((json) => {
      if (json["status"] == "success") {
        document.getElementById(
          "result"
        ).innerHTML = `Plaintext: ${json["data"]["plaintext"]}`;
      } else {
        alert(json["data"]["message"]);
      }
    })
    .catch((err) => {
      alert(err);
    });
}

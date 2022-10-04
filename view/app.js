function affine() {
  document.getElementById("choice").value = "affine";
  document.getElementById("choicet").innerHTML =
    "Current chosen is Affine Cipher algorithm";
  document.getElementById("textLabel").innerHTML = "Plaintext/Ciphertext";

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
    "Current chosen is Caesar Cipher algorithm";
  document.getElementById("textLabel").innerHTML = "Plaintext/Ciphertext";

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

function otp() {
  document.getElementById("choice").value = "otp";
  document.getElementById("choicet").innerHTML =
    "Current chosen is One Time Pad algorithm";
  document.getElementById("textLabel").innerHTML = "Plaintext/Ciphertext";

  document.getElementById("keyLabelNumberOne").hidden = true;
  document.getElementById("keyLabelNumberOne").innerHTML = "Key";
  document.getElementById("keyNumberOne").hidden = true;

  document.getElementById("keyLabelNumberTwo").hidden = true;
  document.getElementById("keyNumberTwo").hidden = true;
  document.getElementById("keyLabelNumberTwo").innerHTML = "";

  document.getElementById("keyLabelText").hidden = false;
  document.getElementById("keyText").hidden = false;
  document.getElementById("keyLabelText").innerHTML = "Key";
}

function superEncryption() {
  document.getElementById("choice").value = "super";
  document.getElementById("choicet").innerHTML =
    "Current chosen is Super Encryption algorithm";
  document.getElementById("textLabel").innerHTML = "Plaintext/Ciphertext";

  document.getElementById("keyLabelNumberOne").hidden = false;
  document.getElementById("keyLabelNumberOne").innerHTML = "Key A";
  document.getElementById("keyNumberOne").hidden = false;

  document.getElementById("keyLabelNumberTwo").hidden = false;
  document.getElementById("keyNumberTwo").hidden = false;
  document.getElementById("keyLabelNumberTwo").innerHTML = "Key B";

  document.getElementById("keyLabelText").hidden = false;
  document.getElementById("keyText").hidden = false;
  document.getElementById("keyLabelText").innerHTML = "Key Text";
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

  document.getElementById("textLabel").innerHTML = "Plaintext";
  let plaintext = document.getElementById("text").value;

  if (alg && plaintext && (alg == "otp" || alg == "super")) {
    if (plaintext.length < key.text.length) {
      key.text = key.text.slice(0, plaintext.length);
    } else if (plaintext.length > key.text.length) {
      let multiplication = Math.ceil(plaintext.length / key.text.length);
      let fixedKey = "";
      for (let i = 0; i < multiplication; i++) {
        fixedKey += key.text;
      }
      key.text = fixedKey.slice(0, plaintext.length);
    }
  }

  document.getElementById("keyText").value = key.text;

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
        ).innerHTML = `Ciphertext: ${json["data"]["ciphertext"]}`;
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

  document.getElementById("textLabel").innerHTML = "Ciphertext";
  let ciphertext = document.getElementById("text").value;

  if (alg && ciphertext && (alg == "otp" || alg == "super")) {
    if (ciphertext.length < key.text.length) {
      key.text = key.text.slice(0, ciphertext.length);
    } else if (ciphertext.length > key.text.length) {
      let multiplication = Math.ceil(ciphertext.length / key.text.length);
      let fixedKey = "";
      for (let i = 0; i < multiplication; i++) {
        fixedKey += key.text;
      }
      key.text = fixedKey.slice(0, ciphertext.length);
    }
  }

  document.getElementById("keyText").value = key.text;

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

document.getElementById("text").addEventListener("keydown", function (e) {
  let key = e.key.charCodeAt();
  let alg = document.getElementById("choice").value;
  if (alg == "railfence" || alg == "super") {
    if (key < 65 || (key > 90 && key < 97) || key > 122) {
      e.preventDefault();
    }
  } else {
    if (
      key < 32 ||
      (key > 32 && key < 65) ||
      (key > 90 && key < 97) ||
      key > 122
    ) {
      e.preventDefault();
    }
  }

  e.target.value = e.target.value.toUpperCase();
});

document.getElementById("text").addEventListener("input", function (e) {
  e.target.value = e.target.value.toUpperCase();
});

document.getElementById("keyText").addEventListener("keydown", function (e) {
  let key = e.key.charCodeAt();
  if (key < 65 || (key > 90 && key < 97) || key > 122) {
    e.preventDefault();
  }
});

document.getElementById("keyText").addEventListener("input", function (e) {
  e.target.value = e.target.value.toUpperCase();
});

// icons
let clipboardIcon = `<svg
    class="fill-current w-4 h-4 mr-2"
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 24 24"
    >
    <path
      d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"
    ></path>
    <rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect>
    </svg>`;
let tickIcon = `<svg
    class="fill-current w-4 h-4 mr-2"
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 20 20"
    >
    <path
    fill-rule="evenodd"
    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
    clip-rule="evenodd"
    />
    </svg>`;

document.getElementById("clipboard").innerHTML =
  clipboardIcon + "<span>Copy</span>";

// clipboard toggle
let toggleClipboard = false;
// copyToClipboard copies the text to the clipboard
function copyToClipboard(elementId) {
  // Get the text field
  var copyText = document.getElementById(elementId);
  // Select the text field
  copyText.select();
  copyText.setSelectionRange(0, 99999); /* For mobile devices */

  // Copy the text inside the text field
  navigator.clipboard.writeText(copyText.value);

  // Alert the copied text
  //   alert("Copied the text: " + copyText.value);

  // Toggle Icon
  clipButton = document.getElementById("clipboard");
  if (toggleClipboard == false) {
    clipButton.innerHTML = tickIcon + "<span>Copied</span>";
    toggleClipboard = true;
  } else {
    clipButton.innerHTML = clipboardIcon + "<span>Copy</span>";
    toggleClipboard = false;
  }
}

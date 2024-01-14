function readURL(input) {
  if (input.files && input.files[0]) {
    var reader = new FileReader();

    reader.onload = function (e) {
      document.querySelector('.profile-pic').setAttribute('src', e.target.result);
      showDeleteButton(); // Show the delete button when an image is selected
    }

    reader.readAsDataURL(input.files[0]);
  }
}

function clickFileUpload() {
  document.querySelector('.file-upload').click();
}

function deleteAvatar() {
  // Reset the file input and remove the currently displayed avatar image
  document.querySelector('.file-upload').value = null;
  document.querySelector('.profile-pic').setAttribute('src', './assets/images/user.png');
  hideDeleteButton(); // Hide the delete button when the image is deleted
}

function showDeleteButton() {
  document.querySelector('.delete-button').style.display = 'block';
}

function hideDeleteButton() {
  document.querySelector('.delete-button').style.display = 'none';
}

document.querySelector('.file-upload').addEventListener('change', function () {
  readURL(this);
});

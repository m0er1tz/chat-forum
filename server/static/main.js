function submitForms() {
      const textForm = document.getElementById('text-form');
      const usernameForm = document.getElementById('username-form');

      const textFormData = new FormData(textForm);
      const usernameFormData = new FormData(usernameForm);

      // Combine form data
      const combinedFormData = new FormData();
      for (let pair of textFormData.entries()) {
        combinedFormData.append(pair[0], pair[1]);
      }
      for (let pair of usernameFormData.entries()) {
        combinedFormData.append(pair[0], pair[1]);
      }

      // Send the POST request
      fetch('/submit', {
        method: 'POST',
        body: combinedFormData
      })
      .then(response => {
        if (response.ok) {
          // Request successful, do something with the response
          console.log('POST request sent successfully!');
        } else {
          // Request failed, handle the error
          console.error('Error:', response.status);
        }
      })
      .catch(error => {
        // An error occurred during the request
        console.error('Error:', error);
      });
	  document.getElementById('text-input').value = '';
      //document.getElementById('username-input').value = '';
    }
function fetchMessages() {
      fetch('/get-messages')
        .then(response => response.json())
        .then(data => {
		  console.log("json", data)
          const messageList = document.getElementById('message-list');
          messageList.innerHTML = '';

          data.messages.forEach(message => {
			console.log(message)
            const li = document.createElement('li');
            li.textContent = message.text;
            messageList.appendChild(li);
          });
        })
        .catch(error => {
          console.error('Error:', error);
        });
    }
fetchMessages();

setInterval(fetchMessages, 100);



